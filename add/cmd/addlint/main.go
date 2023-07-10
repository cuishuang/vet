package main

import (
	"vettt/add/addcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	//singlechecker.Main(addcheck.Analyzer)
	singlechecker.Main(addcheck.Analyzer2)
}
