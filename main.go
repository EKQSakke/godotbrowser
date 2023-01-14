package main

import (
	"fmt"
	"os"
	"strings"

	"io"
	"log"
	"net/http"

	tea "github.com/charmbracelet/bubbletea"
)

const baseUrl = "https://downloads.tuxfamily.org/godotengine/4.0/";

func main() {
	desc := request(baseUrl)
	p := tea.NewProgram(InitialModel(desc, baseUrl, ""))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func request(input string) string {
	print(input)
	resp, err := http.Get(input)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			if strings.Contains(m.choices[m.cursor].href, "Name") || strings.Contains(m.choices[m.cursor].href, "Parent") {
				lastHref := strings.Split(m.currentUrl, "/")
				m.currentUrl = strings.TrimSuffix(m.currentUrl, lastHref[len(lastHref)-2] + "/");
				ResetModelsTo(&m, request(m.currentUrl), m.currentUrl, ":" + lastHref[len(lastHref)-2] + "/")
			} else {
				ResetModelsTo(&m, request(m.currentUrl + m.choices[m.cursor].href), m.currentUrl + m.choices[m.cursor].href, m.choices[m.cursor].href)
			}
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// The header
	s := fmt.Sprintf("%s \n", m.currentUrl)
	s += fmt.Sprintf("%s \n\n", m.errorMsg)

	// Iterate over our choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice.ToString())
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}
