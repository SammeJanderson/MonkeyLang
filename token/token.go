package token


type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // indentifiers + literals
    IDENT = "IDENT"
    INT = "INT"

    /* operators */
    ASSING = "="
    PLUS = "+"

    // demiliters
    COMMA = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // keywords 
    FUNCTION = "FUNCTION"
    LET = "LET"
)

var keywords = map[string]TokenType{
    "fn" : FUNCTION,
    "let" : LET,
}

func LookUpIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
