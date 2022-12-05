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

ipynb = `{
    "cells": [
     {
      "cell_type": "code",
      "execution_count": null,
      "metadata": {
       "dotnet_interactive": {
        "language": "csharp"
       },
       "vscode": {
        "languageId": "dotnet-interactive.csharp"
       }
      },
      "outputs": [],
      "source": [
       "using System.IO;"
      ]
     },
     {
      "cell_type": "code",
      "execution_count": null,
      "metadata": {
       "dotnet_interactive": {
        "language": "csharp"
       },
       "vscode": {
        "languageId": "dotnet-interactive.csharp"
       }
      },
      "outputs": [],
      "source": [
       "var lines = File.ReadAllLines(\\"input.txt\\");"
      ]
     },
     {
      "cell_type": "code",
      "execution_count": null,
      "metadata": {
       "dotnet_interactive": {
        "language": "csharp"
       },
       "vscode": {
        "languageId": "dotnet-interactive.csharp"
       }
      },
      "outputs": [],
      "source": []
     }
    ],
    "metadata": {
     "kernelspec": {
      "display_name": ".NET (C#)",
      "language": "C#",
      "name": ".net-csharp"
     },
     "language_info": {
      "file_extension": ".cs",
      "mimetype": "text/x-csharp",
      "name": "C#",
      "pygments_lexer": "csharp",
      "version": "9.0"
     },
     "orig_nbformat": 4
    },
    "nbformat": 4,
    "nbformat_minor": 2
   }`

   fs.writeFile(dir + "input.txt", "placeholder", function(err) {
    if(err) {
        return console.log(err);
    }
    console.log("The input.txt was saved!");
}); 

fs.writeFile(dir + "p"+n+".ipynb", ipynb, function(err) {
    if(err) {
        return console.log(err);
    }
    console.log("The notebook was saved!");
}); 