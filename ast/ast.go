package ast

import (
	"fmt"
	"monkey/token"
)

// Nodeはノードに実装されるインターフェース
type Node interface {
	// TokenLiteralはそのノードが関連付けられているトークンの
	// リテラル値を返す
	TokenLiteral() string
}

// Statementは文に実装されるインターフェース
type Statement interface {
	Node

	// ダミーメソッド
	statementNode()
}

// Expressionは式に実装されるインターフェース
type Expression interface {
	Node

	// ダミーメソッド
	expressionNode()
}

// Identifierは識別子のASTノード
type Identifier struct {
	// token.IDENTトークン
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ProgramはASTのルートノード
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatementは代入文のASTノード
type LetStatement struct {
	// token.LETトークン
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) String() string {
	return fmt.Sprintf("<LetStatement Name=%s Value=%s>",
		ls.Name.Value, ls.Value)
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
