package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

    // 識別子
	IDENT = "IDENT" // 変数名
	INT   = "INT" // 数字

    // 演算子
	ASSIGN = "="
	PLUS   = "+"
    MINUS  = "-"
    BANG   = "!"
    ASTERISK = "*"
    SLASH  = "/"
    EQ     = "=="
    NOT_EQ = "!="

    LT = "<"
    GT = ">"

    //　デリミタ
	COMMA     = ","
	SEMICOLON = ";"

    // 括弧
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

    // キーワード
	FUNCTION = "FUNCTION"
	LET      = "LET"
    TRUE     = "TRUE"
    FALSE    = "FALSE"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"
)

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
