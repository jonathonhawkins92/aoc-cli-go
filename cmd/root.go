/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/jonathonhawkins92/aoc-cli-go/cmd/generate"
	"github.com/jonathonhawkins92/aoc-cli-go/cmd/run"
	"github.com/spf13/cobra"
)

// https://patorjk.com/software/taag/#p=testall&f=Isometric3&t=AoCGo
const logo = `

      ___           ___           ___           ___           ___     
     /  /\         /  /\         /  /\         /  /\         /  /\    
    /  /::\       /  /::\       /  /:/        /  /:/_       /  /::\   
   /  /:/\:\     /  /:/\:\     /  /:/        /  /:/ /\     /  /:/\:\  
  /  /:/~/::\   /  /:/  \:\   /  /:/  ___   /  /:/_/::\   /  /:/  \:\ 
 /__/:/ /:/\:\ /__/:/ \__\:\ /__/:/  /  /\ /__/:/__\/\:\ /__/:/ \__\:\
 \  \:\/:/__\/ \  \:\ /  /:/ \  \:\ /  /:/ \  \:\ /~~/:/ \  \:\ /  /:/
  \  \::/       \  \:\  /:/   \  \:\  /:/   \  \:\  /:/   \  \:\  /:/ 
   \  \:\        \  \:\/:/     \  \:\/:/     \  \:\/:/     \  \:\/:/  
    \  \:\        \  \::/       \  \::/       \  \::/       \  \::/   
     \__\/         \__\/         \__\/         \__\/         \__\/    

`

var (
	logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
)

func init() {
	fmt.Printf("%s\n", logoStyle.Render(logo))
}

var rootCmd = &cobra.Command{
	Version: "v0.0.1",
	Use:     "AoCGo",
	Short:   "Advent of Code Go Lang CLI",
	Long:    `An Advent of Code Command Line Interface written in Go Lang`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(generate.GenereateCmd)
	rootCmd.AddCommand(run.RunCmd)
}
