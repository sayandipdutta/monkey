package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/sayandipdutta/monkey/lexer"
	"github.com/sayandipdutta/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}
		lexer := lexer.New(line)
		parser := parser.New(lexer)
		program := parser.ParseProgram()
		fmt.Println(program.String())
	}
}
