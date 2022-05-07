package repl

import (
	"bufio"
	"fmt"
	"interpreter-in-go/lexer"
	"interpreter-in-go/token"
	"io"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, ">>>")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for _token := l.NextToken(); _token.Type != token.EOF; _token = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", _token)
		}
	}
}
