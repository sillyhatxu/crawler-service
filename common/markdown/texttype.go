package markdown

type TextType int

const (
	Text TextType = iota
	H1
	H2
	H3
	H4
	Italic
	Bold
	Monospace
	Itemize
	BlockQuotes
	NumberList
	CodeBlockQuotes
	Unknown
)
