const inquirer = require("inquirer");

module.exports = {
  askIfWantConfigure: () => {
    const questions = [
      {
        name: "configure",
        type: "confirm",
        message: "Would you like to register this env?",
        default: true
      },
    ];
    return inquirer.prompt(questions);
  },
  askLogoColor: () => {
    const questions = [
      {
        name: "message",
        type: "input",
        message: "What is the logo color? (example: #FFFFFF)",
        validate: function (value) {
          if (value.length) {
            return true;
          } else {
            return "Please write the hex color";
          }
        },
      },
    ];
    return inquirer.prompt(questions);
  },
  askContextApp: () => {
    const questions = [
      {
        name: "message",
        type: "input",
        message: "What is the scope? (type words comma separated)",
        validate: function (value) {
          if (value.length) {
            return true;
          } else {
            return "Please write a scope names, ex: home, products, settings";
          }
        },
      },
    ];
    return inquirer.prompt(questions);
  },
  askCommitType: () => {
    const questions = [
      {
        type: "list",
        name: "ignore",
        message: "What is commit type?",
        choices: ["feat", "fix", "docs", "style", "refactor", "test", "chore"],
        default: [],
      },
    ];
    return inquirer.prompt(questions);
  },
  askContextType: (choices) => {
    const questions = [
      {
        type: "list",
        name: "ignore",
        message: "What is scope?",
        choices: choices,
        default: [],
      },
    ];
    return inquirer.prompt(questions);
  },
  askCommitMessage: () => {
    const questions = [
      {
        name: "message",
        type: "input",
        message: "Now type what did you do:",
        validate: function (value) {
          if (value.length) {
            return true;
          } else {
            return "Hey the message cannot be empty";
          }
        },
      },
    ];
    return inquirer.prompt(questions);
  },
  askBreakingChanges: () => {
    const questions = [
      {
        name: "message",
        type: "input",
        message: "Breaking change:",
      },
    ];
    return inquirer.prompt(questions);
  },
  askLooksLikeGood: () => {
    const questions = [
      {
        name: "isGood",
        type: "confirm",
        message: "All right? ???????",
        default: true
      },
    ];
    return inquirer.prompt(questions);
  }
};
