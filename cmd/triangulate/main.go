package main

import (
	"flag"
	"fmt"
	"integrate2d/pkg/triangulation"
	"os"
)

func usage() {
	fmt.Println("usage:")
	flag.PrintDefaults()
	os.Exit(2)
}

var lengthAlongAxisX float64
var lengthAlongAxisY float64
var numSplitsAxisX uint64
var numSplitsAxisY uint64

func init() {
	flag.Float64Var(&lengthAlongAxisX, "lengthAlongAxisX", 1.0, "rectangle length along X axis")
	flag.Float64Var(&lengthAlongAxisY, "lengthAlongAxisY", 1.0, "rectangle length along Y axis")
	flag.Uint64Var(&numSplitsAxisX, "numSplitsAxisX", 100, "number of splits along X axis")
	flag.Uint64Var(&numSplitsAxisY, "numSplitsAxisY", 100, "number of splits along Y axis")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	out := make(chan string)

	go triangulation.GenerateTriangulation(
		lengthAlongAxisX,
		lengthAlongAxisY,
		numSplitsAxisX,
		numSplitsAxisY,
		out,
	)

	for line := range out {
		fmt.Print(line)
	}
}
