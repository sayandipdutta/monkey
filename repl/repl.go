package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/sayandipdutta/monkey/lexer"
	"github.com/sayandipdutta/monkey/token"
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
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
