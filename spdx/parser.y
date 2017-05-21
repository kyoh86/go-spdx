%{
package spdx

import(
  "io"
  "strings"
)

type Expression interface {
  String() string
  Expression() string
  bracedIfNeeded() string
}

type Operation struct {
  Left Expression
  Right Expression
}

type Or Operation

func (o Or) Expression() string {
  return strings.Join([]string{o.Left.Expression(), "OR", o.Right.Expression()}, " ")
}

func (o Or) String() string {
  return o.Expression()
}

func (o Or) bracedIfNeeded() string {
  return strings.Join([]string{"(", o.Left.Expression(), "OR", o.Right.Expression(), ")"}, " ")
}

type And Operation

func (a And) Expression() string {
  return strings.Join([]string{a.Left.bracedIfNeeded(), "AND", a.Right.bracedIfNeeded()}, " ")
}

func (a And) String() string {
  return a.Expression()
}

func (a And) bracedIfNeeded() string {
  return a.Expression()
}

type License string

func (l License) Expression() string {
  return string(l)
}

func (l License) String() string {
  return l.Expression()
}

func (l License) bracedIfNeeded() string {
  return string(l)
}

%}

%union{
  exp Expression
  operator string
  license License
}

%type<exp> spdx operation operand braced
%token<license> LICENSE
%token<operator> OR AND
%left OR
%left AND

%%

spdx:
  LICENSE
  {
    yylex.(*spdxLexer).result = License($1)
  }
  | operation
  {
    yylex.(*spdxLexer).result = $1
  }

operation:
  operand OR operand
  {
    $$ = Or{ $1, $3 }
  }
  | operand AND operand
  {
    $$ = And{ $1, $3 }
  }

operand:
  LICENSE
  {
    $$ = License($1)
  }
  | operation
  {
    $$ = $1
  }
  | braced
  {
    $$ = $1
  }

braced:
  '(' operation ')'
  {
    $$ = $2
  }

%%

type spdxLexer struct {
  Lexer
  err error
  syntaxErr string
  result Expression
}

func (s *spdxLexer) Lex(lval *yySymType) int {
  if s.err != nil {
    return -1
  }
  token, err := s.NextToken()
  switch err {
  case nil:
    // noop
  case io.EOF:
    return -1
  default:
    println(err)
    s.err = err
    return -1
  }
  switch token.Type {
    case TokenTypeRaw:
      return int(token.Raw)
    case TokenTypeOperatorOr:
      return OR
    case TokenTypeOperatorAnd:
      return AND
    case TokenTypeLicense:
      lval.license = License(token.License)
      return LICENSE
  }
  panic("invalid operation")
}

func (s *spdxLexer) Error(e string) {
  s.syntaxErr = e
}

func Parse(expression string) (Expression, error) {
  lex := &spdxLexer{ Lexer: Lexer{expression: []rune(expression)} }
  yyParse(lex)
  if lex.err != nil {
    return lex.result, lex.err
  }
  if lex.syntaxErr != "" {
    return lex.result, lex.errorf(lex.syntaxErr)
  }
  return lex.result, nil
}
