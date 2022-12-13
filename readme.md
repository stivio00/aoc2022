# Advent of Code 2022

By Stephen Krol

## Create a new problem folder
To create a new problem folder use
```` bash
node create.js {n}
````

where __{n}__ is the problem number.

## Set your session token
To set your session TOKEN
set the file session.env with the cookie content
```text
session=a235a33e57c...
```
This session token is used by the input.js utility script.

### Where to find the session token
By using Chrome developer tools you can inspect into secure cookies

## Get input File

To get the input file for the problem _3_
```bash
node input.js 3 > p3/input.txt
```

## Use init.sh 
An alternative to executing the two commands (create, input), we can init a new problem and fetch the problem data in one step.
```bash
init.sh 4
```
this will automate the creation and fetching of the problem 4 in one step.


## How to open this Repo
Use VScode with the Polyglot Notebooks extension to open the *.ipynb files.
The .NET Extension Pack v1.0.9 is recommended with .NET 6 or 7

### Deps
 * Node JS 18.x
 * DotNet 5, 6 or 7 
 * Go 1.19.4
 * Bash Shell (or ZSH)
 * VS Code
   * .NET Extension Pack
   * Go v0.36.0