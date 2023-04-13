#!/usr/bin/bash

./triangulate.out --lengthAlongAxisX=$1 --lengthAlongAxisY=$2 --numSplitsAxisX=$3 --numSplitsAxisY=$4 > in.txt
./integrate.out < in.txt