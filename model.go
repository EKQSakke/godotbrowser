package main

import "strings"

type model struct {
	desc    string
	choices []row // items on the to-do list
	cursor  int      // which to-do list item our cursor is pointing at
}

func InitialModel(input string) model {
	splitInput := strings.Split(input, "\n")
	var validInputs []row

	for _, inputRow := range splitInput {
		if strings.Contains(inputRow, "<tr>") {
			validInputs = append(validInputs, toRow(inputRow))
		}
	}

	return model{
		desc: input,
		choices: validInputs,
	}
}

func toRow(input string) row {
	return row{

	}
}

type row struct {
	href string
	text string
	dateModified string
	itemType string
}