package tokenizer

import (
	"errors"
)

type Tokenizer struct {
	tokens   []*Token
	position int
}

func (t *Tokenizer) GetNext() (*Token, error) {
	if t.position >= len(tokens) {
		return nil, errors.New("empty tokenizer")
	}

	token := t.tokens[t.position]
	t.position++
	return token, nil
}

func (t *Tokenizer) GetAll() []*Token {
	return t.tokens, nil
}

func Tokenize(input string) *Tokenizer {
	return nil
}
