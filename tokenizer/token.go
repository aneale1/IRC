package tokenizer

type Token struct {
	Kind             string
	Content          string
	StartPos, EndPos int
}
