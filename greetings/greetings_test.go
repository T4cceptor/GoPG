package greetings

import (
	"regexp"
	"strings"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Gladys"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := PrintMessage(name)
	if !want.MatchString(message) || err != nil {
		t.Fatalf(
			`Hello("Gladys") = %q, %v, want match 
			for %#q, nil`,
			message,
			err,
			want,
		)
	}
}

func TestSimpleHelloName(t *testing.T) {
	// Given a Name as string
	name := "Draven"
	// When: executing the PrintMessage function
	message, err := PrintMessage(name)
	// Then: the message contains the given name
	if !strings.Contains(message, name) || err != nil {
		t.Fatalf("test didnt work")
	}
}
