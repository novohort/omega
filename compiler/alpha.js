const fs = require("fs");
const { exec } = require("child_process");
const tokenize = require("./tokenizer");
const parse = require("./parser");
const codegen = require("./codegen");
const rustFilename = "hello_world.rs";
const omega_input_filepattern = /^.*\.o$/i;
const omega_output_file = process.argv[3];
const omega_input_filename = process.argv[2];

if (!omega_input_filepattern.test(omega_input_filename)) {
  console.error("Omega input filename unrecognized");
  return;
}

if (!omega_output_file) {
  console.error("Omega output filename unrecognized");
  return;
}

fs.readFile(omega_input_filename, "utf8", (err, data) => {
  if (err) {
    console.error(`Error reading file ${omega_input_filename}: ${err}`);
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

      exec(
        `rustc ${rustFilename} -o ${omega_output_file}`,
        (err, stdout, stderr) => {
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
          });

          const pdbFilename = omega_output_file.replace(".exe", ".pdb");
          fs.unlink(pdbFilename, (err) => {
            if (!err) console.log(`${pdbFilename} deleted successfully.`);
          });

          console.log(`Compiled successfully to ${omega_output_file}`);
        }
      );
    });
  } catch (error) {
    console.error(error);
  }
});
