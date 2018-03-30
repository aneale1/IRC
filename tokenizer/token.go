package tokenizer

// Token is a single element of the Tokenizer.
// It contains one of 4 Kinds: chstring, numstring, special, control,
// each of which corresponds to a kind of input readable by the server.
// The Content of the Token corresponds to the byte(s) read by the
// server.
type Token struct {
	Kind             string
	Content          string
	StartPos, EndPos int
}
