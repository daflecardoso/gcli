const inquirer = require("inquirer");

module.exports = {
  askIfWantConfigure: () => {
    const questions = [
      {
        name: "configure",
        type: "confirm",
        message: "Would you like to register like a new env?",
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
        message: "What is the contexts app? (type words comma separated)",
        validate: function (value) {
          if (value.length) {
            return true;
          } else {
            return "Please write a contexts names, ex: home, products, settings";
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
        message: "Qual o tipo do commit?",
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
        message: "Qual o contexto do app?",
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
        message: "Agora descreva o que foi feito:",
        validate: function (value) {
          if (value.length) {
            return true;
          } else {
            return "Por favor a mensagem não pode ser vazia";
          }
        },
      },
    ];
    return inquirer.prompt(questions);
  },
  askLooksLikeGood: () => {
    const questions = [
      {
        name: "isGood",
        type: "confirm",
        message: "Está tudo certo? ☝🏼",
        default: true
      },
    ];
    return inquirer.prompt(questions);
  }
};
