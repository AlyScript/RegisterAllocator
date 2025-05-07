#!/bin/bash

if [ -f output.txt ] && [ -f expected.txt ]; then
    go run main.go input.txt output.txt
    diff output.txt expected.txt > /dev/null
    if [ $? -eq 0 ]; then
        echo "Test passed"
    else
        echo "Test failed"
    fi
else
    echo "Missing input file or expected output file"
fi
