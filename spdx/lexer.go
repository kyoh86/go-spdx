package spdx

import (
	"fmt"
	"io"
)

// Lexer will parse expressions and get all tokens from it.
type Lexer struct {
	expression []rune
	isOperator bool
	index      int
}

// TokenType represents that which type a token is.
type TokenType int

const (
	// TokenTypeOperatorAnd is "AND" operator
	TokenTypeOperatorAnd = iota
	// TokenTypeOperatorOr is "OR" operator
	TokenTypeOperatorOr
	// TokenTypeLicense is a part of the license name
	TokenTypeLicense
	// TokenTypeRaw is the raw rune
	TokenTypeRaw
)

// Token holds a token and a type of it
type Token struct {
	Type    TokenType
	License string
	Raw     rune
}

// LexerError represents an error in parsing expression
type LexerError struct {
	message    string
	expression []rune
	index      int
}

func (s LexerError) Error() string {
	head := s.index - 5
	if head < 0 {
		head = 0
	}
	tail := s.index + 5
	if tail > len(s.expression) {
		tail = len(s.expression)
	}
	return fmt.Sprintf("%s at %d (%q)", s.message, s.index, string(s.expression[head:tail]))
}

// InvalidOperatorError represents that expression contains invalid operator
type InvalidOperatorError LexerError

func (e InvalidOperatorError) Error() string { return LexerError(e).Error() }

// InvalidLicenseError represents that expression is not valid license spdx notation
type InvalidLicenseError LexerError

func (e InvalidLicenseError) Error() string { return LexerError(e).Error() }

func (s *Lexer) skipSpaces() {
	for ; s.index < len(s.expression); s.index++ {
		if s.expression[s.index] == ' ' {
			continue
		}
		break
	}
}

func (s *Lexer) skipWord() {
	for ; s.index < len(s.expression); s.index++ {
		switch s.expression[s.index] {
		case '(', ')', ' ':
		default:
			continue
		}
		break
	}
}

func (s *Lexer) genRawRuneToken() *Token {
	t := s.expression[s.index]
	s.index++
	return &Token{Type: TokenTypeRaw, Raw: t}
}

func (s *Lexer) genOperatorTokenFrom(head int) (*Token, error) {
	operator := string(s.expression[head:s.index])
	s.isOperator = false
	switch operator {
	case "AND":
		return &Token{Type: TokenTypeOperatorAnd}, nil
	case "OR":
		return &Token{Type: TokenTypeOperatorOr}, nil
	default:
		s.index = head
		return nil, InvalidOperatorError(s.errorf("invalid operator %q", operator))
	}
}

func (s *Lexer) genLicenseTokenFrom(head int) (*Token, error) {
	license := string(s.expression[head:s.index])
	s.isOperator = true
	if _, ok := validLicenses[license]; !ok {
		s.index = head
		return nil, InvalidLicenseError(s.errorf("invalid license %q", license))
	}
	return &Token{Type: TokenTypeLicense, License: license}, nil
}

// NextToken searches next token from expression
func (s *Lexer) NextToken() (*Token, error) {
	s.skipSpaces()
	if len(s.expression) <= s.index {
		return nil, io.EOF
	}
	switch s.expression[s.index] {
	case '(', ')':
		return s.genRawRuneToken(), nil
	default:
		head := s.index
		s.skipWord()

		if s.isOperator {
			return s.genOperatorTokenFrom(head)
		}
		return s.genLicenseTokenFrom(head)
	}
}

func (s *Lexer) errorf(message string, args ...interface{}) LexerError {
	return LexerError{message: fmt.Sprintf(message, args...), index: s.index, expression: s.expression}
}
