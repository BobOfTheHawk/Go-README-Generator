package main

import "github.com/AlecAivazis/survey/v2"

// Answers holds all the data collected from the user's prompts.
// The `survey` tag is used to match the question name to the struct field.
type Answers struct {
	ProjectName string `survey:"name"`
	Version     string `survey:"version"`
	Description string `survey:"description"`
	Homepage    string `survey:"homepage"`
	AuthorName  string `survey:"authorName"`
	GithubUser  string `survey:"githubUser"`
	TwitterUser string `survey:"twitterUser"`
	InstallCmd  string `survey:"install"`
	UsageCmd    string `survey:"usage"`
	TestCmd     string `survey:"test"`
	License     string `survey:"license"`
}

// The set of questions to ask the user.
var questions = []*survey.Question{
	{
		Name:      "name",
		Prompt:    &survey.Input{Message: "👤 What is the name of your project?"},
		Validate:  survey.Required,
		Transform: SanitizeInput,
	},
	{
		Name:      "version",
		Prompt:    &survey.Input{Message: "📦 What is the project version? (e.g., 1.0.0)"},
		Transform: SanitizeInput,
	},
	{
		Name:      "description",
		Prompt:    &survey.Input{Message: "📝 Please provide a short description of your project:"},
		Validate:  survey.Required,
		Transform: SanitizeInput,
	},
	{
		Name:      "homepage",
		Prompt:    &survey.Input{Message: "🏠 What is the project's homepage URL?"},
		Transform: SanitizeInput,
	},
	{
		Name:      "authorName",
		Prompt:    &survey.Input{Message: "✍️  What is the author's name?"},
		Transform: SanitizeInput,
	},
	{
		Name:      "githubUser",
		Prompt:    &survey.Input{Message: "🐙 What is the author's GitHub username?"},
		Transform: SanitizeInput,
	},
	{
		Name:      "twitterUser",
		Prompt:    &survey.Input{Message: "🐦 What is the author's Twitter username (without the @)?"},
		Transform: SanitizeInput,
	},
	{
		Name: "license",
		Prompt: &survey.Select{
			Message: "📜 Choose a license for your project:",
			Options: []string{"MIT", "Apache 2.0", "GPLv3", "BSD 3-Clause", "Unlicense", "None"},
			Default: "MIT",
		},
	},
	{
		Name:      "install",
		Prompt:    &survey.Input{Message: "🛠️  What command should be run to install dependencies?", Default: "npm install"},
		Transform: SanitizeInput,
	},
	{
		Name:      "usage",
		Prompt:    &survey.Input{Message: "⚙️  What command should be run to use the project?", Default: "npm start"},
		Transform: SanitizeInput,
	},
	{
		Name:      "test",
		Prompt:    &survey.Input{Message: "🧪 What command should be run to run tests?", Default: "npm test"},
		Transform: SanitizeInput,
	},
}

// RunPrompts executes the survey and returns the collected answers.
func RunPrompts() (*Answers, error) {
	answers := &Answers{}
	err := survey.Ask(questions, answers)
	if err != nil {
		return nil, err
	}
	return answers, nil
}
