/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package run

import (
	"log"
	"strconv"

	runDay "github.com/jonathonhawkins92/aoc-cli-go/components/day"
	"github.com/spf13/cobra"
)

func Run() {
	runDay.Main()
}

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run an AoC solution",
	Run: func(cmd *cobra.Command, args []string) {
		day_raw := cmd.Flag("day").Value.String()
		day, err := strconv.Atoi(day_raw)
		if err != nil || day < 1 || day > 25 {
			log.Println("bad day")
			runDay.Main()
		}

		part_raw := cmd.Flag("part").Value.String()
		part, err := strconv.Atoi(part_raw)
		if err != nil || part < 1 || part > 2 {
			log.Println("bad part")
			runDay.Main()
		}

		input := cmd.Flag("input").Value.String()
		if input != "test" && input != "input" {
			log.Println("bad input")
			runDay.Main()
		}
		log.Println(day, part, input)
	},
}

func init() {
	RunCmd.Flags().Int16P("day", "d", -1, "Which day do you want to run? (min 1) (max 25)")
	RunCmd.Flags().Int16P("part", "p", -1, "Which part do you want to run? (min 1) (max 2)")
	RunCmd.Flags().StringP("input", "i", "", "Which input do you want to use? (input | test)")
}
