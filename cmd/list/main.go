/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package list

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func Main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "A list example",
	Long:  `A List example using cobra, bubbletea and lipgloss.`,
	Run: func(cmd *cobra.Command, args []string) {
		Main()
	},
}

type model struct {
	choices  []string
	cursor   int
	selected int
	input    string
}

func initialModel() model {
	choices := make([]string, 0)
	for i := 0; i < 25; i++ {
		day := strconv.Itoa(i + 1)
		choices = append(choices, "Day "+day)
	}

	return model{
		choices:  choices,
		selected: -1,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		lastInput := m.input
		m.input = msg.String()

		switch m.input {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.selected = m.cursor
		}
		day, err := strconv.Atoi(m.input)
		if err == nil {
			if lastInput == "1" {
				m.cursor = 10 + day - 1
			} else if lastInput == "2" {
				m.cursor = 20 + day - 1
			} else {
				m.cursor = day - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What should we buy at the market?\n\n"

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selected == i {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}
