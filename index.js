#!/usr/bin/env node

const chalk = require("chalk");
const clear = require("clear");
const figlet = require("figlet");
const inquirer = require("./lib/inquire");
const { shell } = require("./lib/utils");
const fs = require('fs');
const path = require("path");

clear();

const configFileName = "gcli.json"

const env = process.argv[process.argv.length - 1]

if (env == '--update') {
  console.log("Updating gcli...")
  shell(`
  cd ~/.gcli
  git pull
  `);
  console.log("\x1b[32m", `\n🎉 Successful updated`);
  return;
}

const file = path.resolve(configFileName);
if (!fs.existsSync(file)) { 
  console.log("\x1b[32m", `\nPlease create a ${configFileName} file`)
  return;
}

let config = JSON.parse(fs.readFileSync(file));

console.log(
  chalk.hex(config.color)(figlet.textSync(config.name, { horizontalLayout: "full" }))
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

const run = async () => {
  
  const commitType = await inquirer.askCommitType();
  
  const contextType = await inquirer.askContextType(config.scopes);

  const commit = await inquirer.askCommitMessage();

  const breakingChange = await inquirer.askBreakingChanges();
  
  var breakingMessage = breakingChange.message;

  if (breakingMessage) {
    breakingMessage = `\n\nBREAKING CHANGE: ${breakingMessage}`
  }

  const commitMessage = `${commitType.ignore}(${contextType.ignore}): ${commit.message}${breakingMessage}`
  console.log("\n")
  const FgGreen = "\x1b[33m"
  console.log(FgGreen, commitMessage);
  console.log("\n")

  const askIsGood = await inquirer.askLooksLikeGood();

  if (askIsGood.isGood) {
     
      const gitAdd = "git add ."
      await shell(gitAdd);
      console.log(`✅ ${gitAdd}`)

      const gitCommit = `git commit -m "${commitMessage}"`;
      await shell(gitCommit);
      console.log(`✅ ${gitCommit}`)
      
      const gitPush = `git push`;
      try {
        console.log(`✅ ${gitPush}`);
        await shell(gitPush);
      } catch (err) {
        console.log(`🔴 ${gitPush}\n`);
        console.log(err.message)
      }
  }
};

run();
