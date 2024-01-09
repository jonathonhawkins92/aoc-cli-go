/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package root

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"text/template"

	"github.com/jonathonhawkins92/aoc-cli-go/cmd/generate/generators/generator"
)

type mainVars struct {
	generator.Config
	Packages []string
	Days     []string
}

func createMain(config generator.Config) {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	currentDir := generator.RightRemoveSegment(filename)

	path := filepath.Join(currentDir, "..", "..", "..", "templates", "main.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(config.TargetDirectory, "main.go")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	days := make([]string, 0)
	packages := make([]string, 0)
	for i := 1; i <= 25; i++ {
		dayDigit := strconv.Itoa(i)
		days = append(days, dayDigit)
		packages = append(packages, filepath.Join(config.Repository, config.PackageName, "day"+dayDigit))
	}

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, mainVars{
		Config:   config,
		Packages: packages,
		Days:     days,
	})
	if err != nil {
		panic(err)
	}
}

func createMod(config generator.Config) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	currentDir := generator.RightRemoveSegment(filename)

	path := filepath.Join(currentDir, "..", "..", "..", "templates", "go.mod.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(config.TargetDirectory, "go.mod")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, config)
	if err != nil {
		panic(err)
	}
}

func Create(config generator.Config) {
	createMod(config)
	createMain(config)
}
