/*
Copyright © 2024 Jonathon Hawkins <garbage@thedumpster.com>
*/
package dir

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func Main() string {
	p := tea.NewProgram(initialModel())
	result, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
	return result.(model).value
}

type model struct {
	value       string
	suggestions []string
	selected    int
}

func initialModel() model {
	return model{
		value:       "",
		suggestions: make([]string, 0),
		selected:    0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRunes:
			m.value += string(msg.Runes)
			m.suggestions = nil
		case tea.KeyBackspace:
			if len(m.value) > 0 {
				m.value = m.value[:len(m.value)-1]
			}
			m.suggestions = nil
		case tea.KeyTab:
			m.suggestions = getSuggestions(m.value)
			m.selected = 0
		case tea.KeyUp, tea.KeyDown:
			if len(m.suggestions) <= 0 {
				return m, nil
			}
			if msg.Type == tea.KeyUp {
				m.selected--
				if m.selected < 0 {
					m.selected = len(m.suggestions) - 1
				}
			} else {
				m.selected++
				if m.selected >= len(m.suggestions) {
					m.selected = 0
				}
			}
		case tea.KeyEnter:
			if len(m.suggestions) <= 0 {

				wd, err := os.Getwd()
				if err != nil {
					log.Panicln("Error:", err)
				}
				m.value = filepath.Join(wd, m.value)

				return m, tea.Quit
			}
			if len(m.value) > 0 {
				runes := []rune(m.value)
				lastRune := runes[len(runes)-1]
				if string(lastRune) != "/" {
					m.value += string(os.PathSeparator)
				}
			}
			m.value += m.suggestions[m.selected] + string(os.PathSeparator)
			m.suggestions = nil
		}
	}
	return m, nil
}

func getSuggestions(input string) []string {
	var suggestions []string

	dir, file := filepath.Split(input)

	wd, err := os.Getwd()
	if err != nil {
		log.Panicln("Error:", err)
	}

	// Read the directory contents
	objects, err := os.ReadDir(filepath.Join(wd, dir, file))
	if err != nil {
		return nil
	}

	for _, object := range objects {
		if !object.IsDir() {
			continue
		}
		name := object.Name()
		suggestions = append(suggestions, name)
	}

	return suggestions
}

func (m model) View() string {
	s := "Which day did you want to run?\n"
	s += "Directory: " + fmt.Sprintln(m.value)
	if len(m.suggestions) > 0 {
		for index, suggestion := range m.suggestions {
			checked := " "
			if m.selected == index {
				checked = "x"
			}

			s += fmt.Sprintf("[%s] %s\n", checked, suggestion)
		}
	}
	s += "\n"

	return s
}
