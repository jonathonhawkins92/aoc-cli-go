/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package generate

import (
	"log"
	"os"

	"github.com/jonathonhawkins92/aoc-cli-go/cmd/generate/generators"
	"github.com/jonathonhawkins92/aoc-cli-go/cmd/generate/generators/generator"
	"github.com/jonathonhawkins92/aoc-cli-go/components/dir"
	"github.com/spf13/cobra"
)

func Run(repository string, targetDirectory string) {
	config := generator.Config{
		Repository:      repository,
		TargetDirectory: targetDirectory,
		PackageName:     "2023",
	}

	generators.Root(config)
	generators.Days(config)

	log.Println(config.PackageName, "created in", config.TargetDirectory)
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
			log.Fatalf("Unable to create directory: directory - %v\n", directory)
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
