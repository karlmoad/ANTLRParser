package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
)

func ParseString(input string) (antlr.ParseTree, error) {
	return parse(antlr.NewInputStream(input))
}

func parse(stream *antlr.InputStream) (antlr.ParseTree, error) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(ParsingErrorInterface); ok {
				log.Println(err)
			} else {
				panic(r)
			}
		}
	}()

	lexer := NewPlSqlLexer(stream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(NewErrorListener())

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := NewPlSqlParser(tokens)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(NewErrorListener())

	return parser.Sql_script(), nil
}

type ParsingErrorInterface interface {
	Where() (int, int)
}

type ParsingError struct {
	Line   int
	Column int
	Msg    string
}

func (p *ParsingError) Where() (int, int) {
	return p.Line, p.Column
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("Parsing error on line %d column %d: %s", e.Line, e.Column, e.Msg)
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	err error
}

func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

func (e *ErrorListener) ParsingError(recognizer antlr.Recognizer,
	offendingSymbol any,
	line int,
	column int,
	msg string,
	ex antlr.RecognitionException) {
	panic(&ParsingError{line, column, msg})
}
