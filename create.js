var fs = require('fs');

// console.log(process.argv)

if (process.argv.length !== 3 ) {
    console.error("Wrong number of parameters. "+ process.argv.length)
    process.exit(-1);
}
const n = process.argv[2] 
var dir = './p' + n +"/";
console.log(`Creating ${dir}`);
if (!fs.existsSync(dir)){
    fs.mkdirSync(dir);
}

gomod = `module github.com/stivio00/aoc2022/p${n}

go 1.19
`

gofile = `package main

func main() {

}
`

fs.writeFile(dir + "input.txt", "placeholder", function(err) {
    if(err) {
        return console.log(err);
    }
    console.log("The input.txt was saved!");
}); 

fs.writeFile(dir + "go.mod", gomod, function(err) {
    if(err) {
        return console.log(err);
    }
    console.log("The notebook was saved!");
}); 

fs.writeFile(dir + "main.go", gofile, function(err) {
    if(err) {
        return console.log(err);
    }
    console.log("The notebook was saved!");
}); 