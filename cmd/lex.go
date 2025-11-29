package cmd

import (
	"fmt"
	"os"
	"strings"
)

func lexCommand(){
	const klangVersion = "v0.1.10"
	const cliVersion = "v0.1.0"

	if len(os.Args) < 3{
		fmt.Println("Usage: loom lex <.k file>")
		return
	}

	fmt.Println("╭────────────────────────────────────────────────────────╮")
	fmt.Println("│  loom — Klang Project Manager                          │")
	fmt.Printf("│  version %s-dev  •  Klang Core %s             │\n", cliVersion, klangVersion)
	fmt.Println("╰────────────────────────────────────────────────────────╯")
	fmt.Println()

	file := os.Args[2]
	if !strings.HasSuffix(file, ".k"){
		indexOfDot := strings.LastIndex(file, ".")

		if indexOfDot == -1 {
			fmt.Println("No suffix/extension found in: ", file)
			return
		}

		sufixo := file[indexOfDot:]

		fmt.Println("Erro: Not a .k file")
		fmt.Println("────────────────────────────────────────────────────────")

		fmt.Println("File: ", file, " is a ", sufixo, " file")
		fmt.Println("did you provide the wrong file?")
		
		return
	}

    data, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("\n❌ Error when opening ", file, ": ", err)
		return
	}




	fmt.Println("loom lex ", file)
	fmt.Println("────────────────────────────────────────────────────────")

	fmt.Println("Running lexer in (simulation)...")
	fmt.Println("\n✔ ", file, " successfully lexicalized!")

	fmt.Println("Conteúdo de ", file, " (primeiros 50 bytes):")
	fmt.Println("────────────────────────────────────────────────────────")
	fmt.Println(string(data[:50])) 

	fmt.Println()

	fmt.Println("Tokens:")
	fmt.Println("────────────────────────────────────────────────────────")
	fmt.Println("Tokens here...")

}