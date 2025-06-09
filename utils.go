package main

import "regexp"

// stripAnsi is a regex to strip ANSI escape codes from pasted text.
var stripAnsi = regexp.MustCompile("[\u001B\u009B][[\\]()#;?]*.{0,2}(?:(?:;*.{0,2})*)[a-zA-Z\\d]")

// SanitizeInput is a survey.Transformer that strips ANSI escape codes.
// This prevents the program from freezing if the user pastes text with control characters.
func SanitizeInput(ans interface{}) interface{} {
	if str, ok := ans.(string); ok {
		return stripAnsi.ReplaceAllString(str, "")
	}
	return ans
}
