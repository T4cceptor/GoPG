package greetings

import (
	"errors"
	"fmt"
	"strings"
	"text/template"
)

func main() {
	fmt.Println("Greetings!")
}

// Returns a nice greeting based on the provided string.
func PrintMessage(str string) (string, error) {
	if str == "" {
		return "", errors.New("Empty string")
	}
	var result string = fmt.Sprintf(
		RandomFormat(),
		str,
	) // initialize variable
	return result, nil // some result
}

func GetMessages(names []string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, errors.New("Empty names!")
	}

	result := make(map[string]string)
	for _, name := range names {
		// name := names[index]
		// fmt.Printf("Test: %v, %v\n", index, name)
		new_message := fmt.Sprintf(
			RandomFormat(),
			name,
		)
		result[name] = new_message
	}
	return result, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func RandomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Example to use string template to create format complex output and store it as variable
	tmpl, _ := template.New("test").Parse("{{.test1}} items are made of %v")
	b := new(strings.Builder)
	_ = tmpl.Execute(b, map[string]any{"test1": "Whatever test 123"})
	txt := b.String()

	// Append new format string here
	// fmt.Println(txt)
	formats = append(
		formats,
		txt,
	)

	// Checking how to call a function from a sub-directory
	// greetings.TestExport1()

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return randomEntryTest(formats)
}
