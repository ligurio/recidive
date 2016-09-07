// https://help.github.com/articles/searching-issues/
// "cat format:junit status:pass created:2006-08-10"
// go/scanner
// text/scanner

// FIXME: add ability to show line and position

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Token int

const (
	IN Token = iota
	FORMAT
	STATUS
	CREATED

	ILLEGAL
	EOF
	WS

	IDENT

	EQUAL
	LESS
	EQLESS
	LESSEQ
	MORE
	EQMORE
	MOREEQ
	COLON
)

var (
	TokenKeyword = map[Token]string{
		IN:      "IN",
		FORMAT:  "FORMAT",
		STATUS:  "STATUS",
		CREATED: "CREATED",
	}

	TokenOperation = map[Token]string{
		EQUAL:  "=",
		LESS:   "<",
		EQLESS: "=<",
		LESSEQ: "<=",
		MORE:   ">",
		EQMORE: "=>",
		MOREEQ: ">=",
		COLON:  ":",
	}

	eof = rune(0)
)

type SearchString struct {
	SearchKeyword []string
	Params        []Expression
}

type Expression struct {
	Keyword   string
	Operation string
	Value     string
}

type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

func (s *Scanner) Scan() (tok Token, lit string) {
	count := 0
	ch := s.read()

	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isOperation(ch) {
		s.unread()
		return s.scanOperation()
	} else if isLetter(ch) {
		s.unread()
		return s.scanIdent()
	}

	switch ch {
	case eof:
		return EOF, ""
	}

	return ILLEGAL, string(ch)
}

func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			buf.WriteRune(ch)
		}
	}

	return WS, buf.String()
}

func (s *Scanner) scanOperation() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isOperation(ch) {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	switch strings.ToUpper(buf.String()) {
	case TokenKeyword[EQUAL]:
		return EQUAL, buf.String()
	case TokenKeyword[MORE]:
		return MORE, buf.String()
	case TokenKeyword[LESS]:
		return LESS, buf.String()
	case TokenKeyword[EQLESS]:
		return EQLESS, buf.String()
	case TokenKeyword[LESSEQ]:
		return EQLESS, buf.String()
	case TokenKeyword[EQMORE]:
		return EQMORE, buf.String()
	case TokenKeyword[MOREEQ]:
		return EQMORE, buf.String()
	}

	return IDENT, buf.String()
}

func (s *Scanner) scanIdent() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	for {
		if ch := s.read(); ch == eof {
			break
		} else if !isLetter(ch) && !isDigit(ch) && ch != '_' {
			s.unread()
			break
		} else {
			_, _ = buf.WriteRune(ch)
		}
	}

	switch strings.ToUpper(buf.String()) {
	case TokenKeyword[IN]:
		return IN, buf.String()
	case TokenKeyword[FORMAT]:
		return FORMAT, buf.String()
	case TokenKeyword[STATUS]:
		return STATUS, buf.String()
	case TokenKeyword[CREATED]:
		return CREATED, buf.String()
	}

	return IDENT, buf.String()
}

func (s *Scanner) read() rune {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return eof
	}
	return ch
}

func (s *Scanner) unread() { _ = s.r.UnreadRune() }

func isWhitespace(ch rune) bool { return ch == ' ' || ch == '\t' || ch == '\n' }

func isLetter(ch rune) bool { return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') }

func isDigit(ch rune) bool { return (ch >= '0' && ch <= '9') }

func isOperation(ch rune) bool { return ch == '>' || ch == '<' || ch == '=' || ch == ':' }

func main() {
	const src = `cat format:junit status:pass created<=2009 created>=2008`

	s := NewScanner(strings.NewReader(src))
	var tok Token
	var l string
	for tok != EOF {
		tok, l = s.Scan()
		fmt.Println(l)
	}

}
