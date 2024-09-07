package localtea

// A simple program demonstrating the text input component from the Bubbles
// component library.

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

var GlobalInputWizardModel InputWizardModel

type InputWizardModel struct {
	TextInput textinput.Model
	Err       error
}

func InitialInputModel() InputWizardModel {
	ti := textinput.New()
	ti.Placeholder = "Private Key"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 200

	return InputWizardModel{
		TextInput: ti,
		Err:       nil,
	}
}

func (m InputWizardModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m InputWizardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.Err = msg
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	GlobalInputWizardModel = m
	return m, cmd
}

func (m InputWizardModel) View() string {
	return fmt.Sprintf(
		"Please Enter your private key?\n\n%s\n\n%s",
		m.TextInput.View(),
		"(esc to quit)",
	) + "\n"
}
