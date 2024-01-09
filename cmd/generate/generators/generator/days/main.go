/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package days

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"text/template"

	"github.com/jonathonhawkins92/aoc-cli-go/cmd/generate/generators/generator"
)

type dayConfig struct {
	generator.Config
	Day  string
	path string
}

func createMainInput(config dayConfig) {
	mainOutputPath := filepath.Join(config.path, "input.txt")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
}

func createTestInput(config dayConfig) {
	mainOutputPath := filepath.Join(config.path, "test.txt")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
}

func createMain(config dayConfig) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	currentDir := generator.RightRemoveSegment(filename)

	path := filepath.Join(currentDir, "..", "..", "..", "templates", "day", "main.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(config.path, "main.go")
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

func create1Main(config dayConfig) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	currentDir := generator.RightRemoveSegment(filename)

	path := filepath.Join(currentDir, "..", "..", "..", "templates", "day", "part1", "main.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(config.path, "part1", "main.go")
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

func create1Test(config dayConfig) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	currentDir := generator.RightRemoveSegment(filename)

	path := filepath.Join(currentDir, "..", "..", "..", "templates", "day", "part1", "main_test.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(config.path, "part1", "main_test.go")
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

func create2Main(config dayConfig) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	currentDir := generator.RightRemoveSegment(filename)

	path := filepath.Join(currentDir, "..", "..", "..", "templates", "day", "part2", "main.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(config.path, "part2", "main.go")
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

func create2Test(config dayConfig) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	currentDir := generator.RightRemoveSegment(filename)

	path := filepath.Join(currentDir, "..", "..", "..", "templates", "day", "part2", "main_test.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(config.path, "part2", "main_test.go")
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

func createDay(generatorConfig generator.Config, dayDigit int) {
	day := "day" + strconv.Itoa(dayDigit)
	config := dayConfig{Config: generatorConfig, Day: day, path: filepath.Join(generatorConfig.TargetDirectory, day)}

	err := os.Mkdir(config.path, 0755)
	if err != nil {
		log.Fatalf("Unable to create directory: day - %v\n", config.path)
	}

	createMainInput(config)
	createTestInput(config)
	createMain(config)

	part1Path := filepath.Join(config.path, "part1")
	err = os.Mkdir(part1Path, 0755)
	if err != nil {
		log.Fatalf("Unable to create directory: part1Path - %v\n", part1Path)
	}
	create1Main(config)
	create1Test(config)

	part2Path := filepath.Join(config.path, "part2")
	err = os.Mkdir(part2Path, 0755)
	if err != nil {
		log.Fatalf("Unable to create directory: part2Path - %v\n", part2Path)
	}
	create2Main(config)
	create2Test(config)
}

func Create(config generator.Config) {
	for i := 1; i <= 25; i++ {
		createDay(config, i)
	}
}
