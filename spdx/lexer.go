package spdx

import (
	"fmt"
	"io"
)

type Lexer struct {
	expression []rune
	isOperator bool
	index      int
}

type TokenType int

const (
	TokenTypeOperatorAnd = iota
	TokenTypeOperatorOr
	TokenTypeLicense
	TokenTypeRaw
)

type Token struct {
	Type    TokenType
	License string
	Raw     rune
}

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

type InvalidOperatorError LexerError

func (e InvalidOperatorError) Error() string { return LexerError(e).Error() }

type InvalidLicenseError LexerError

func (e InvalidLicenseError) Error() string { return LexerError(e).Error() }

func (s *Lexer) NextToken() (*Token, error) {
	for ; s.index < len(s.expression); s.index++ {
		if s.expression[s.index] == ' ' {
			continue
		}
		break
	}
	if len(s.expression) <= s.index {
		return nil, io.EOF
	}
	switch s.expression[s.index] {
	case '(', ')':
		t := s.expression[s.index]
		s.index++
		return &Token{Type: TokenTypeRaw, Raw: t}, nil
	default:
		head := s.index

		for ; s.index < len(s.expression); s.index++ {
			switch s.expression[s.index] {
			case '(', ')', ' ':
			default:
				continue
			}
			break
		}
		if s.isOperator {
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
		} else {
			license := string(s.expression[head:s.index])
			s.isOperator = true
			if _, ok := validLicenses[license]; !ok {
				s.index = head
				return nil, InvalidLicenseError(s.errorf("invalid license %q", license))
			}
			return &Token{Type: TokenTypeLicense, License: license}, nil
		}
	}
}

func (s *Lexer) errorf(message string, args ...interface{}) LexerError {
	return LexerError{message: fmt.Sprintf(message, args...), index: s.index, expression: s.expression}
}
