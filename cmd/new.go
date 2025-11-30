package cmd

import (
	"fmt"
	"os"
)


func newCommand(){
	if len(os.Args) < 3{
		fmt.Println("Usage: loom new <project_name>")
		return
	}

	log := NewLog()
	name := os.Args[2]
	
	log.Header()
	log.finalizeBottomheader()
	fmt.Println()
	
	log.Line()
	fmt.Printf("%s◉%s Creating project structure...\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	
	err := os.Mkdir(name, 0755)
	if err != nil {
	    fmt.Printf("%s✖ Error:%s %s\n", log.ERROR_COLOR, log.RESET_COLOR, err)
	    return
	}

	fmt.Printf("%s◉%s Writing manifest...\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("%s◉%s Initializing Klang module...\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("%s◈%s Finalizing...\n", log.ACCENT_COLOR, log.RESET_COLOR)

	fmt.Printf("\n%s✔%s Project '%s' created.\n",
	    log.SUCESS_COLOR, log.RESET_COLOR, name)
	log.Line()

}