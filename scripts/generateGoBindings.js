const fs = require('fs')
const path = require('path')
const { exec } = require("child_process")

const abis = fs.readdirSync(path.resolve(__dirname, '../abigenBindings/abi'))

abis.forEach((abi) => {
  const source = path.resolve(__dirname, '../abigenBindings/abi', abi)
  const goName = abi.replace(".abi", "")
  const destination = path.resolve(__dirname, '../abigo', goName)

  const cmd = `abigen --abi ${source} --pkg abigo --type ${goName} --out ${destination}.go`
  exec(cmd, (error, stdout, stderr) => {
    if (error) {
        console.log(`error: ${error.message}`);
        return;
    }
    if (stderr) {
        console.log(`stderr: ${stderr}`);
        return;
    }
    console.log(`${abi} -> ${goName}.go`);
});
})
