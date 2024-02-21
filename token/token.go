package token

type TokenType string

type Token struct {
  Type      TokenType
  Literal  string
}

var keywords = map[string]TokenType{
  "fn": FUNCTION,
  "let": LET,
  "true": TRUE,
  "false": FALSE,
  "if": IF,
  "else": ELSE,
  "return": RETURN,
}

func LookupIdent(ident string) TokenType {
  if tok, ok := keywords[ident]; ok {
    return tok
  }
  return IDENT
}

const (
  ILLEGAL   = "ILLEGAL"
  EOF       = "EOD"
  IDENT     = "IDENT"
  INT       = "INT"
  STRING    = "STRING"
  ASSIGN    = "="
  PLUS      = "PLUS"
  MINUS     = "MINUS"
  BANG      = "!"
  ASTERISK  = "*"
  SLASH     = "/"
  EQ        = "=="
  NOT_EQ    = "!="
  LT        = "<"
  GT        = ">"
  COMMA     = ","
  COLON     = ":"
  SEMICOLON = ";"
  LPAREN    = "("
  RPAREN    = ")"
  LBRACE    = "{"
  RBRACE    = "}"
  LBRACKET  = "["
  RBRACKET  = "]"
  FUNCTION  = "FUNCTION"
  LET       = "LET"
  TRUE      = "TRUE"
  FALSE     = "FALSE"
  IF        = "IF"
  ELSE      = "ELSE"
  RETURN    = "RETURN"
)
