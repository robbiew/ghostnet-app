#!/bin/bash

# This script is an example of how to build the application

echo "Chekcking bin/ drirectory exists..."
if [ ! -d "bin" ]; then
    echo "making bin/ directory..."
    mkdir bin
fi

echo "building binary..."
go build .

echo "moving binary to bin/ directory..."
mv ./ghostnet-app bin/ 

echo "copyting binary to Talisman door directory"
cp bin/ghostnet-app /Users/whiting/talisman/doors/ghostnet-app/

echo "> Build completed successfully!"
echo "> Executable is located in bin/ directory"
echo "> Executable is also copied to Talisman door directory"