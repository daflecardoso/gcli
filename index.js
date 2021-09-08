#!/usr/bin/env node

const chalk = require("chalk");
const clear = require("clear");
const figlet = require("figlet");
const inquirer = require("./lib/inquire");
const { exec } = require("child_process");
const fs = require('fs');
const path = require("path");

clear();

const configFileName = "config.json"

const env = process.argv[process.argv.length - 1]

if (env == '--update') {
  console.log("Updating gcli...")
  shell("cd ~/gcli/gcli git pull");
  const FgGreen = "\x1b[32m"
  console.log(FgGreen, `\n🎉 Successful updated`);
  return;
}

if (process.argv.length == 2) {
  const FgGreen = "\x1b[32m"
  console.log(FgGreen, `\n🎉 Welcome to gcli, to configure your first env type:`);
  console.log('$ gcli <env>')
  return
}

if (process.argv.includes('-d')) {
  if (env != '-d') {
    const FgGreen = "\x1b[32m"

    const file = path.resolve(__dirname, configFileName);
    let rawdata = fs.readFileSync(file);
    let json = JSON.parse(rawdata);
    json[env] = undefined;
    saveConfigFile(json);

    console.log(FgGreen, `\n🎉 ${env} env were removed`)
    return
  } else {
    const FgRed = "\x1b[31m"
    console.log(FgRed, '\n🚨 Please you must tell the env ex: $ gcli -d <enviroment>');
    return
  }
}

const file = path.resolve(__dirname, configFileName);
if (!fs.existsSync(file)) {
  const FgGreen = "\x1b[32m"
  console.log(FgGreen, '\n🎉 Welcome to gcli')
  registerNewEnv(false);
  return
}

let rawdata = fs.readFileSync(file);
let json = JSON.parse(rawdata);
let config = json[env];

if (!config) {
  registerNewEnv(true)
  return;
}

function registerNewEnv(confExists) {
  const FgRed = "\x1b[31m"
  console.log(FgRed, `\n🚨 Enviroment ${env} isn't configured\n`)

  const run = async () => {
    const configure = await inquirer.askIfWantConfigure();

    if (configure) {

      const hexColor = await inquirer.askLogoColor();
      
      const contextApp = await inquirer.askContextApp();

      const newEnv = {
        cliName: env,
        cliNameColor: hexColor.message, 
        showTutorial: false,
        contextChoices: contextApp.message.split(',').map((s) => s.trim()),
      };
     
      if (confExists) {
        json[env] = newEnv;
        saveConfigFile(json);
      } else {
        let envs = {};
        envs[env] = newEnv
        saveConfigFile(envs);
      }
      const FgGreen = "\x1b[32m"
      console.log(FgGreen, `\n🎉 Success! now you can use: $ gcli ${env}`)
    }
  };
  run();
}

function saveConfigFile(envs) {
  const file = path.resolve(__dirname, configFileName);
  let data = JSON.stringify(envs, null, 2);
  fs.writeFileSync(file, data);
}

console.log(
  chalk.hex(config.cliNameColor)(figlet.textSync(config.cliName, { horizontalLayout: "full" }))
);

if (config.showTutorial) {
  console.log(`
feat: (new feature for the user, not a new feature for build script)
fix: (bug fix for the user, not a fix to a build script)
docs: (changes to the documentation)
style: (formatting, missing semi colons, etc; no production code change)
refactor: (refactoring production code, eg. renaming a variable)
test: (adding missing tests, refactoring tests; no production code change)
chore: (updating grunt tasks etc; no production code change)
`);
}

function shell(command) {
  return new Promise((resolve, reject) => {
    exec(command, (error, stdout, stderr) => {
      if (error) {
        console.log(error)
          reject(error)
          return;
      }
      if (stderr) {
          //console.log(`stderr: ${stderr}`);
          return;
      }
      console.log(`stdout: ${stdout}`);
      
     // console.log(`stdout: ${stdout}`);
    });
  })
}

const run = async () => {
  const commitType = await inquirer.askCommitType();
  
  const contextType = await inquirer.askContextType(config.contextChoices);

  const commit = await inquirer.askCommitMessage();
  
  const commitMessage = `${commitType.ignore}(${contextType.ignore}): ${commit.message}`
  console.log("\n")
  console.log(commitMessage);
  console.log("\n")

  const isGood = await inquirer.askLooksLikeGood();

  if (isGood) {
      const gitAdd = "git add ."
      console.log(`✅ ${gitAdd}`)
      await shell(gitAdd);

      const gitCommit = `git commit -m "${commitMessage}"`;
      console.log(`✅ ${gitCommit}`)
      await shell(gitCommit);
      
      const gitPush = `git push`;
      console.log(`✅ ${gitPush}`);
      const res = await shell(gitPush);
      console.log(res)
  }
};

run();
