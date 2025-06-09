package main

import "fmt"

func main() {
	// Greet the user and explain what's happening.
	fmt.Println("🎉 Welcome to the Go README Generator!")
	fmt.Println("Let's create a beautiful README for your project. Please answer the following questions.")
	fmt.Println("------------------------------------------------------------------------------------")

	// Run the interactive prompts to gather project information.
	// This function is defined in prompts.go
	answers, err := RunPrompts()
	if err != nil {
		// This can happen if the user hits Ctrl+C.
		fmt.Println("\nAborted by user. No README.md generated.")
		return
	}

	fmt.Println("------------------------------------------------------------------------------------")
	fmt.Println("👍 Great! I have all the information I need.")

	// Generate the README.md file using the collected answers.
	// This function is defined in template.go
	err = GenerateReadme(answers)
	if err != nil {
		fmt.Printf("❌ Oh no! An error occurred: %v\n", err)
		return
	}

	fmt.Println("✅ Successfully generated README.md! ✨ You can find it in the current directory.")
}
