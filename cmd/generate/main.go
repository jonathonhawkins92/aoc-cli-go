/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package generate

import (
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"github.com/jonathonhawkins92/aoc-cli-go/cmd/generate/dir"
	"github.com/jonathonhawkins92/aoc-cli-go/components/dir"
	"github.com/spf13/cobra"
)

func createMainRoot(repository string, targetDirectory string, packageName string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	parts := strings.Split(filename, string(os.PathSeparator))
	allByLast := parts[:len(parts)-1]
	currentDir := strings.Join(allByLast, string(os.PathSeparator))

	path := filepath.Join(currentDir, "templates", "main.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(targetDirectory, "main.go")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	days := make([]string, 0)
	packages := make([]string, 0)
	for i := 1; i <= 25; i++ {
		index := strconv.Itoa(i)
		days = append(days, index)
		packages = append(packages, filepath.Join(repository, packageName, "day"+index))
	}

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, struct {
		PackageName string
		Packages    []string
		Days        []string
		Repository  string
	}{
		PackageName: packageName,
		Packages:    packages,
		Days:        days,
		Repository:  repository,
	})
	if err != nil {
		panic(err)
	}
}

func createMod(targetDirectory string, packageName string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	parts := strings.Split(filename, string(os.PathSeparator))
	allByLast := parts[:len(parts)-1]
	currentDir := strings.Join(allByLast, string(os.PathSeparator))

	path := filepath.Join(currentDir, "templates", "go.mod.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(targetDirectory, "go.mod")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, struct {
		PackageName string
	}{
		PackageName: packageName,
	})
	if err != nil {
		panic(err)
	}
}

func createDayMainInput(targetDirectory string) {
	mainOutputPath := filepath.Join(targetDirectory, "input.txt")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
}

func createDayTestInput(targetDirectory string) {
	mainOutputPath := filepath.Join(targetDirectory, "test.txt")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()
}

func createDayRoot(targetDirectory string, repository string, day string, packageName string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	parts := strings.Split(filename, string(os.PathSeparator))
	allByLast := parts[:len(parts)-1]
	currentDir := strings.Join(allByLast, string(os.PathSeparator))

	path := filepath.Join(currentDir, "templates", "day", "main.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(targetDirectory, "main.go")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, struct {
		PackageName string
		Day         string
		Repository  string
	}{
		PackageName: packageName,
		Day:         day,
		Repository:  repository,
	})
	if err != nil {
		panic(err)
	}
}

func createDay1Root(targetDirectory string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	parts := strings.Split(filename, string(os.PathSeparator))
	allByLast := parts[:len(parts)-1]
	currentDir := strings.Join(allByLast, string(os.PathSeparator))

	path := filepath.Join(currentDir, "templates", "day", "part1", "main.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(targetDirectory, "part1", "main.go")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, struct{}{})
	if err != nil {
		panic(err)
	}
}

func createDay1Test(targetDirectory string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	parts := strings.Split(filename, string(os.PathSeparator))
	allByLast := parts[:len(parts)-1]
	currentDir := strings.Join(allByLast, string(os.PathSeparator))

	path := filepath.Join(currentDir, "templates", "day", "part1", "main_test.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(targetDirectory, "part1", "main_test.go")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, struct{}{})
	if err != nil {
		panic(err)
	}
}

func createDay2Root(targetDirectory string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	parts := strings.Split(filename, string(os.PathSeparator))
	allByLast := parts[:len(parts)-1]
	currentDir := strings.Join(allByLast, string(os.PathSeparator))

	path := filepath.Join(currentDir, "templates", "day", "part2", "main.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(targetDirectory, "part2", "main.go")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, struct{}{})
	if err != nil {
		panic(err)
	}
}

func createDay2Test(targetDirectory string) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("No file information")
		return
	}

	parts := strings.Split(filename, string(os.PathSeparator))
	allByLast := parts[:len(parts)-1]
	currentDir := strings.Join(allByLast, string(os.PathSeparator))

	path := filepath.Join(currentDir, "templates", "day", "part2", "main_test.go.tmpl")
	// Open the template file
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		panic(err)
	}

	mainOutputPath := filepath.Join(targetDirectory, "part2", "main_test.go")
	// Create a file to hold the generated code
	outputFile, err := os.Create(mainOutputPath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Execute the template and write to the output file
	err = tmpl.Execute(outputFile, struct{}{})
	if err != nil {
		panic(err)
	}
}

func createDay(repository string, targetDirectory string, packageName string) {
	for i := 1; i <= 25; i++ {
		day := strconv.Itoa(i)
		dayPath := filepath.Join(targetDirectory, "day"+day)

		err := os.Mkdir(dayPath, 0755)
		if err != nil {
			log.Fatalf("Unable to create directory: %v\n", dayPath)
		}

		createDayMainInput(dayPath)
		createDayTestInput(dayPath)
		createDayRoot(dayPath, repository+"/"+targetDirectory, day, packageName)

		part1Path := filepath.Join(dayPath, "part1")
		err = os.Mkdir(part1Path, 0755)
		if err != nil {
			log.Fatalf("Unable to create directory: %v\n", dayPath)
		}
		createDay1Root(dayPath)
		createDay1Test(dayPath)

		part2Path := filepath.Join(dayPath, "part2")
		err = os.Mkdir(part2Path, 0755)
		if err != nil {
			log.Fatalf("Unable to create directory: %v\n", dayPath)
		}
		createDay2Root(dayPath)
		createDay2Test(dayPath)
	}
}

func Run(repository string, targetDirectory string) {
	packageName := "2023"

	createMainRoot(repository, targetDirectory, packageName)
	createMod(targetDirectory, packageName)
	createDay(repository, targetDirectory, packageName)

	log.Println(packageName, "created in", targetDirectory)
}

var GenereateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate an AoC project",
	Long:    `This command will generate an AoC project with day 1 -> 25`,
	Aliases: []string{"gen"},
	Run: func(cmd *cobra.Command, args []string) {
		directory := cmd.Flag("directory").Value.String()

		if directory == "" {
			directory = dir.Main()
		}

		if doesDirectoryExistAndIsNotEmpty(directory) {
			log.Fatalf("Directory exisits and is not empty: %v\n", directory)
		}

		err := os.MkdirAll(directory, 0755)
		if err != nil {
			log.Fatalf("Unable to create directory: %v\n", directory)
		}
		repository := cmd.Flag("repository").Value.String()

		Run(repository, directory)
	},
}

func init() {
	GenereateCmd.Flags().StringP("repository", "r", "", "Which repository do you want your project to be in?")
	GenereateCmd.Flags().StringP("directory", "d", "", "Which directory do you want your project to be in?")
}

func doesDirectoryExistAndIsNotEmpty(name string) bool {
	if _, err := os.Stat(name); err == nil {
		dirEntries, err := os.ReadDir(name)
		if err != nil {
			log.Panicf("could not read directory: %v\n", err)
		}
		if len(dirEntries) > 0 {
			return true
		}
	}
	return false
}
