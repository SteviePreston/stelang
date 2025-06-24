package repl

import (
	"bufio"
	"fmt"
	lexer "github.com/steviepreston/stelang/interpreter/lexer"
	parser "github.com/steviepreston/stelang/interpreter/parser"
	"io"
)

const PROMPT = ">> "
const STELANG_LOGO = `
     ___  _            _                        
    / __|| |_    ___  | |  __ _   _ _    __ _   
    \__ \|  _|  / -_) | | / _` + "`" + ` | | ' \  / _` + "`" + ` |  
    |___/ \__|  \___| |_| \__,_| |_||_| \__, |  
                                       |___/
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, STELANG_LOGO)
	io.WriteString(out, "Whoops, you messed up! D:\n")
	io.WriteString(out, "Parser Errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
