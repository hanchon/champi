#!/bin/bash
temp=$PWD
echo $temp
cd /tmp
mkdir mud
cd /tmp/mud
wget https://raw.githubusercontent.com/latticexyz/mud/main/packages/store/src/StoreCore.sol
abigen --abi storecore.abi.json --pkg main --out generated.go


