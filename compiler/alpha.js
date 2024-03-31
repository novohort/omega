const fs = require('fs');
const { exec } = require('child_process');
const tokenize = require('./tokenizer');
const parse = require('./parser');
const codegen = require('./codegen');

const omegaFilename = "hello_world.o";
const rustFilename = "hello_world.rs";
const executableName = "hello_world.exe";

fs.readFile(omegaFilename, "utf8", (err, data) => {
  if (err) {
    console.error(`Error reading file ${omegaFilename}: ${err}`);
    return;
  }

  try {
    const tokens = tokenize(data);
    const ast = parse(tokens);
    const rustCode = codegen(ast);

    fs.writeFile(rustFilename, rustCode, (err) => {
      if (err) {
        console.error(`Error writing Rust file: ${err}`);
        return;
      }

      exec(`rustc ${rustFilename} -o ${executableName}`, (err, stdout, stderr) => {
        if (err) {
          console.error(`Compilation error: ${err}`);
          return;
        }
        if (stderr) {
          console.error(`Compiler stderr: ${stderr}`);
          return;
        }
        
        fs.unlink(rustFilename, (err) => {
          if (err) console.error(`Error deleting Rust source file: ${err}`);
          else console.log(`${rustFilename} deleted successfully.`);
        });
        
        const pdbFilename = executableName.replace('.exe', '.pdb');
        fs.unlink(pdbFilename, (err) => {
          if (!err) console.log(`${pdbFilename} deleted successfully.`);
        });

        console.log(`Compiled successfully to ${executableName}`);
      });
    });
  } catch (error) {
    console.error(error);
  }
});
