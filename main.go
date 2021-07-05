package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/TomSuzuki/gomarkdown"
)

func main() {
	flag.Parse()
	inputFile := flag.Args()[0]
	outputFile := inputFile + ".txt"
	fmt.Println("running...")

	md, _ := ioutil.ReadFile(inputFile)
	wiki := gomarkdown.MarkdownToHTML(string(md))

	ioutil.WriteFile("write.txt", []byte(wiki), 0666)

	fmt.Printf("%s -> %s\n", inputFile, outputFile)

}
