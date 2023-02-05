package main

import (
	"github.com/aoman-n/slinter"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(slinter.Analyzer) }
