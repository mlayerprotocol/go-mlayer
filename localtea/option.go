package localtea

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

var GlobalWizardModel WizardModel

type WizardModel struct {
	Choices []string // items on the to-do list
	Cursor  int      // which to-do list item our Cursor is pointing at
	// selected int      // which to-do items are selected
}

func InitialOptionModel(options []string) WizardModel {
	return WizardModel{
		// Our to-do list is a grocery list
		Choices: options,

		// A map which indicates which Choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `Choices` slice, above.
		// selected: 0,
	}
}

func (m WizardModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m WizardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	GlobalWizardModel = m
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the Cursor up
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}

		// The "down" and "j" keys move the Cursor down
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the Cursor is pointing at.
		case "enter", " ":

			return m, tea.Quit

			// _, ok := m.selected[m.Cursor]
			// if ok {
			// 	delete(m.selected, m.Cursor)
			// } else {
			// 	m.selected[m.Cursor] = struct{}{}
			// }
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m WizardModel) View() string {
	// The header
	s := "Start application with::\n\n"

	// Iterate over our Choices
	for i, choice := range m.Choices {

		// Is the Cursor pointing at this choice?
		Cursor := " " // no Cursor
		if m.Cursor == i {
			Cursor = ">" // Cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if i == m.Cursor {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", Cursor, checked, choice)
	}

	// The footer
	s += fmt.Sprintf("\n >>>> %v \n", m.Cursor)
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
