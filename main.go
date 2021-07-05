package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/TomSuzuki/mdtowikigo/mdtowiki"
)

func main() {
	flag.Parse()
	inputFile := flag.Args()[0]
	outputFile := inputFile + ".txt"
	fmt.Println("running...")

	md, _ := ioutil.ReadFile(inputFile)
	wiki := mdtowiki.MdToWiki(string(md))

	ioutil.WriteFile(outputFile, []byte(wiki), 0666)

	fmt.Printf("%s -> %s\n", inputFile, outputFile)

}
