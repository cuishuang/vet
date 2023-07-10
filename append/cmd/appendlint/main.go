package main

import (
	"vettt/append/appendcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	//singlechecker.Main(appendcheck.Analyzer)
	singlechecker.Main(appendcheck.Analyzer2)
}
