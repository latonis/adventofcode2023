#!/bin/bash

if [[ $1 =~ [0-9]{1,2} ]]; then
    mkdir -p $1
    cp -r ./template/* $1/
    echo "Directory for day $1 created!"
    echo "Changing to the newly created directory..."
    echo "{\"rust-analyzer.linkedProjects\": [\"$1/rust/Cargo.toml\"]}" > .vscode/settings.json
    cd ./$1/
    find . -type f -name "Cargo.toml" -exec sed -i "s/day/day$1/g" {} +
 
else 
    echo "Directory not created, invalid input!"
fi