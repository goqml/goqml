package main

import (
	"flag"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
)

func main() {
	genCmd := flag.NewFlagSet("gen", flag.ExitOnError)
	outputFile := genCmd.String("o", "", "Output file (defaults to <input-file>_goqml.go/s)")
	force := genCmd.Bool("f", false, "Force overwrite output file if it exists")
	insert := genCmd.Bool("i", false, "Allow input and output to be the same file")

	if len(os.Args) < 3 {
		fmt.Println("Expected 'gen' subcommand and input file")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "gen":
		if err := genCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if genCmd.NArg() != 1 {
			fmt.Println("Expected one input file")
			os.Exit(1)
		}
		inputFile := genCmd.Arg(0)
		baseName := inputFile[:len(inputFile)-len(filepath.Ext(inputFile))]
		if *outputFile == "" {
			*outputFile = baseName + "_goqml"
		}
		if baseName == *outputFile && !*insert {
			fmt.Printf("Input file and output file cannot be the same: %s\n", inputFile)
			os.Exit(1)
		}
		processFile(inputFile, *outputFile, *force)
	default:
		fmt.Printf("Unknown command %q\n", os.Args[1])
		os.Exit(1)
	}
}

func processFile(input, output string, force bool) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, input, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	parseResult := parseStructs(node)
	generateCode(node.Name.Name, parseResult, output, force)
}
