package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	glamoureverforest "github.com/minimalshenist/glamour-everforest/internal/generator"
)

func main() {
	var generate bool

	flag.BoolVar(&generate, "generate", false, "Generates glamour styles")

	flag.Parse()

	if generate {
		glamoureverforest.Generate()
		return
	}

	in, err := os.ReadFile("tests/example.md")

	if err != nil {
		log.Fatal(err)
	}

	renderer, err := glamour.NewTermRenderer(
		glamour.WithStylePath("styles/everforest.json"),
	)

	if err != nil {
		log.Fatal(err)
	}

	out, err := renderer.RenderBytes(in)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}
