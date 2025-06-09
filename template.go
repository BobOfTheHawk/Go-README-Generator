package main

import (
	"os"
	"strings"
	"text/template"
)

// The new, more detailed README template.
// It uses conditionals to only show sections if the user provided an answer.
const readmeTemplate = `
# {{ .ProjectName }} {{if .Version}} (v{{.Version}}){{end}}

{{if .License | ne "None"}}
![License: {{.License}}](https://img.shields.io/badge/License-{{.License | urlencode}}-blue.svg)
{{end}}
{{if .TwitterUser}}
[![Twitter: {{.TwitterUser}}](https://img.shields.io/twitter/follow/{{.TwitterUser}}?style=social)](https://twitter.com/{{.TwitterUser}})
{{end}}

> {{ .Description }}

{{if .Homepage}}
### 🏠 [Homepage]({{.Homepage}})
{{end}}

## Prerequisites

* Prerequisite 1
* Prerequisite 2

## Install

` + "```sh" + `
{{.InstallCmd}}
` + "```" + `

## Usage

` + "```sh" + `
{{.UsageCmd}}
` + "```" + `

{{if .TestCmd}}
## Run tests

` + "```sh" + `
{{.TestCmd}}
` + "```" + `
{{end}}

{{if .AuthorName}}
## Author

👤 **{{.AuthorName}}**
{{if .GithubUser}}
* Github: [@{{.GithubUser}}](https://github.com/{{.GithubUser}})
{{end}}
{{if .TwitterUser}}
* Twitter: [@{{.TwitterUser}}](https://twitter.com/{{.TwitterUser}})
{{end}}
{{end}}

## 🤝 Contributing

Contributions, issues and feature requests are welcome!
Feel free to check the issues page.

## Show your support

Give a ⭐️ if this project helped you!

{{if .License | ne "None"}}
## 📝 License

Copyright © {{ .AuthorName }}.
This project is [{{.License}}](https://choosealicense.com/licenses/{{.License | lower }}/) licensed.
{{end}}

***
_This README was generated with ❤️ by [Go-README-Generator](https://github.com/your-repo/go-readme-generator)_
`

// GenerateReadme creates the README.md file from the template and answers.
func GenerateReadme(answers *Answers) error {
	// Create the README.md file.
	file, err := os.Create("README.md")
	if err != nil {
		return err
	}
	defer file.Close()

	// FuncMap provides utility functions to the template.
	funcMap := template.FuncMap{
		// urlencode replaces spaces with '--' for badge URLs.
		"urlencode": func(s string) string {
			return strings.ReplaceAll(s, " ", "--")
		},
		"lower": strings.ToLower,
	}

	// Create a new template with our FuncMap and parse the template string.
	tmpl, err := template.New("readme").Funcs(funcMap).Parse(readmeTemplate)
	if err != nil {
		return err
	}

	// Execute the template, passing in the answers, and write to the file.
	err = tmpl.Execute(file, answers)
	if err != nil {
		return err
	}

	return nil
}
