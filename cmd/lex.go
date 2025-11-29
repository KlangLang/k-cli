package cmd

import (
	"fmt"
	"os"
	"strings"
)

func lexCommand(){
	if len(os.Args) < 3{
		fmt.Println("Usage: loom lex <.k file>")
		return
	}

	
	log := NewLog()

	log.Header()
	fmt.Println()

	file := os.Args[2]
	if !strings.HasSuffix(file, ".k"){
		fmt.Printf("%s✖ Error:%s Not a .k file\n", log.ERROR_COLOR, log.RESET_COLOR)
		log.Line()
		fmt.Printf("%sFile:%s %s\n", log.PRIMARY_COLOR, log.RESET_COLOR, file)
		fmt.Printf("%sHint:%s Did you pass the wrong file?\n", log.PRIMARY_COLOR, log.RESET_COLOR)
		
		return
	}

    data, err := os.ReadFile(file)

	if err != nil {
		fmt.Println("\n❌ Error when opening ", file, ": ", err)
		return
	}

	fmt.Printf("%s◉%s Running lexer in simulation...\n", log.PRIMARY_LIGHT, log.RESET_COLOR)
	fmt.Printf("%s✔%s %s successfully lexicalized!\n", log.SUCESS_COLOR, log.RESET_COLOR, file)

	log.Line()
	fmt.Printf("%sPreview (50 bytes):%s\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Println(string(data[:50]))

}