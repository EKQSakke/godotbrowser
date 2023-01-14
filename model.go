package main

import (
	"fmt"
	"strings"
)

type model struct {
	currentUrl string
	desc       string
	choices    []row // items on the to-do list
	cursor     int   // which to-do list item our cursor is pointing at
	errorMsg   string
}

func InitialModel(input string, currentUrl string, errorMsg string) model {
	splitInput := strings.Split(input, "\n")
	var validInputs []row

	for _, inputRow := range splitInput {
		if strings.Contains(inputRow, "<tr>") {
			validInputs = append(validInputs, toRow(inputRow))
		}
	}

	return model{
		desc:       input,
		choices:    validInputs,
		currentUrl: currentUrl,
		errorMsg:   errorMsg,
	}
}

func ResetModelsTo(m *model, input string, currentUrl string, errorMsg string) {
	*m = InitialModel(input, currentUrl, errorMsg)
}

func toRow(input string) row {
	splitInput := strings.Split(input, ">")
	var formattedInput []string

	for _, inputRow := range splitInput {
		line := strings.Split(inputRow, "<")[0]
		formattedInput = append(formattedInput, line)
	}

	return row{
		href:         formattedInput[3] + "/",
		text:         formattedInput[3],
		dateModified: formattedInput[6],
		itemType:     formattedInput[10],
	}
}

type row struct {
	href         string
	text         string
	dateModified string
	itemType     string
}

func (r row) ToString() string {
	return fmt.Sprintf("%s %s %s %s", r.href, r.text, r.dateModified, r.itemType)
}
