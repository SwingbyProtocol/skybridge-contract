const fs = require('fs')
const path = require('path')
const { exec } = require("child_process")

const abis = fs.readdirSync(path.resolve(__dirname, '../abigenBindings/abi'))

try {
  if (!fs.existsSync(`${process.env.GOROOT}/bin/abigen`)) {
    console.error(`ERR: ${process.env.GOROOT}/bin/abigen not found`)
    console.error(`Please install go-ethereum abigen to ${process.env.GOROOT} \n`)
    console.log('cd $GOPATH/src/github.com/ethereum/go-ethereum')
    console.log('go install ./cmd/abigen')
    process.exit()
  }
} catch(err) {
  process.exit()
}

abis.forEach((abi) => {
  const source = path.resolve(__dirname, '../abigenBindings/abi', abi)
  const goName = abi.replace(".abi", "")
  const destination = path.resolve(__dirname, '../abigo', goName)

  const cmd = `${process.env.GOROOT}/bin/abigen --abi ${source} --pkg abigo --type ${goName} --out ${destination}.go`
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
