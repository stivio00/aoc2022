#!/bin/bash
n=$1
echo 🟢Init problem $n folder...
node create.js $n
echo 🟢Fetching input file
node input.js $n > p$n/input.txt
echo 🟢Done