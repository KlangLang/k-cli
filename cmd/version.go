package cmd

import (
	"fmt"
	"os"
)


func showVersion(){
	const cliVersion = "v0.1.0"
	const klangVersion = "v0.1.10"

	if len(os.Args) > 3 {
		if os.Args[2] == "--short"{
			fmt.Println("loom --version")
			fmt.Println("────────────────────────────────────────────────────────")

			fmt.Printf("loom %s-dev\n", cliVersion)
			fmt.Printf("Klang core %s-dev\n", klangVersion)
			return
		} 
		
		fmt.Println("Unexpected command:'", os.Args[3], "'")
		return
		
	}

	fmt.Println("loom --version")
	fmt.Println("────────────────────────────────────────────────────────")
	fmt.Printf("loom %s-dev\n", cliVersion)
	fmt.Printf("Klang core %s-dev\n", klangVersion)
	fmt.Println("Target: JVM", )
	fmt.Println("Build: debug")
	fmt.Println("────────────────────────────────────────────────────────")
}
