const { exec } = require("child_process");

module.exports = {
  shell(command) {
    return new Promise((resolve, reject) => {
      exec(command, (error, stdout, stderr) => {
        if (error) {
          console.log(error);
          reject(error);
          return;
        }
        if (stderr) {
          //console.log(`stderr: ${stderr}`);
          return;
        }
        //console.log(`stdout: ${stdout}`);
        resolve(stdout);
      });
    });
  },
};
