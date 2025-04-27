package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/dan-frohlich/bubbleup"
)

func main() {
	var items []bubbleup.BubbleUpItem[string]
	for _, arg := range os.Args[1:] {
		items = append(items, bubbleup.NewItem[string](arg, arg))
	}
	bu := bubbleup.New[string]().
		WithTitle("Reorder the list").
		WithItems(items...).
		WithTheme(huh.ThemeBase16()).
		WithHelp(true)

	e := ExampleApp{subModel: bu}

	p := tea.NewProgram(e, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		//TODO should we panic?
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}

}

type ExampleApp struct {
	subModel bubbleup.BubbleUp[string]
}

// Init implements tea.Model.
func (e ExampleApp) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (e ExampleApp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "f9", "ctrl+c": //quit
			return e, tea.Quit
		}
	}

	m, cmd := e.subModel.Update(msg)
	if m != nil {
		if bu, ok := m.(bubbleup.BubbleUp[string]); ok {
			e.subModel = bu
		}
	}

	return e, cmd
}

// View implements tea.Model.
func (e ExampleApp) View() string {
	if e.subModel.IsSubmitted() {
		return fmt.Sprintf("list submitted: %s", strings.Join(e.subModel.Values(), ", "))
	}
	return e.subModel.View()
}

var _ tea.Model = ExampleApp{}
