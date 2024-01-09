/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package day

import (
	"fmt"
	"log"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func Main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

var DayCmd = &cobra.Command{
	Use:   "day",
	Short: "Select day 1 - 25",
	Long:  `Select which advent of code day you want to run, days 1 through 25.`,
	Run: func(cmd *cobra.Command, args []string) {
		Main()
	},
}

type model struct {
	selection int
	value     string
}

func initialModel() model {
	return model{
		selection: -1,
		value:     "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:
		prev := m.value

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			return m, tea.Quit
		}

		m.value = msg.String()
		day, err := strconv.Atoi(m.value)
		if err == nil {
			if prev == "1" {
				m.selection = 10 + day
			} else if prev == "2" && day <= 5 {
				m.selection = 20 + day
			} else {
				m.selection = day
			}
		}
	}

	// strconv.Atoi(string(tea.KeyMsg.Runes))

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	s := "Which day did you want to run?\n\n"

	if m.selection != -1 {
		s += fmt.Sprintln(m.selection)
	}

	s += "\nPress q to quit.\n"

	return s
}
