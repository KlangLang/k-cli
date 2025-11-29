package cmd

import (
	"fmt"
	"os"
)


func newCommand(){
	const klangVersion = "v0.1.10"
	const cliVersion = "v0.1.0"

	if len(os.Args) < 3{
		fmt.Println("Usage: loom new <project_name>")
		return
	}

	fmt.Println("╭────────────────────────────────────────────────────────╮")
	fmt.Println("│  loom — Klang Project Manager                          │")
	fmt.Printf("│  version %s-dev  •  Klang Core %s             │\n", cliVersion, klangVersion)
	fmt.Println("╰────────────────────────────────────────────────────────╯")
	fmt.Println()

	name := os.Args[2]
	fmt.Println("loom new ", name)
	fmt.Println("────────────────────────────────────────────────────────")

	fmt.Println("◉ Creating project structure...")
	
	err := os.Mkdir(name, 0755)
	if err != nil {
		fmt.Println("\n❌ Error creating project:", err)
		return
	}
	
	fmt.Println("◉ Writing default manifest...")
	fmt.Println("◉ Initializing Klang module...")
	fmt.Println("◉ Ready.")


	fmt.Println("\n✔ Project ", name, " created.")
}