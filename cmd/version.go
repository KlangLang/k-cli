package cmd

import (
	"fmt"
	"os"
)


func showVersion(){
	log := NewLog()

	if len(os.Args) - 1 >= 2 {
		if os.Args[2] == "--short" || os.Args[2] == "-st"{
			fmt.Println("loom --version")
			log.Line()

			fmt.Printf("loom %s-dev\n", log.LoomVersion)
			fmt.Printf("Klang core %s-dev\n", log.KlangVersion)
			return
		} 
		
		fmt.Println("Unexpected command:'", os.Args[3], "'")
		return
		
	}

	fmt.Println("loom --version")
	log.Line()
	fmt.Println()
	fmt.Printf("%sloom%s %s-dev\n", log.PRIMARY_COLOR, log.RESET_COLOR, log.LoomVersion)
	fmt.Printf("%sKlang Core:%s %s-dev\n", log.PRIMARY_COLOR, log.RESET_COLOR, log.KlangVersion)
	fmt.Printf("%sTarget:%s JVM\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("%sBuild:%s  debug\n", log.PRIMARY_COLOR, log.RESET_COLOR)

}