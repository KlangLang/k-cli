package cmd

import (
	"fmt"
)

func showHelp() {
	const klangVersion = "v0.1.10"
	const cliVersion = "v0.1.0"

	fmt.Println("╭────────────────────────────────────────────────────────╮")
	fmt.Println("│  loom — Klang Project Manager                          │")
	fmt.Printf("│  version %s-dev  •  Klang Core %s             │\n", cliVersion, klangVersion)
	fmt.Println("╰────────────────────────────────────────────────────────╯")
	fmt.Println()

	fmt.Println("loom --help")
	fmt.Println("────────────────────────────────────────────────────────")
	fmt.Println()

	fmt.Println("Commands")
	fmt.Println("	new <project_name>		Create a new Klang project")
	fmt.Println("	lex <.k file>	     		Lexicalize a .k file")

	fmt.Println()

	fmt.Println("Options")
	fmt.Println("	-V, --version   		Show versions")
	fmt.Println("	-h, --help       		Show this help log")

	fmt.Println()

	fmt.Println("Versions")
	fmt.Printf("	loom: 				%s-dev\n", cliVersion)
	fmt.Printf("	Klang core: 			%s-dev\n", klangVersion)

}