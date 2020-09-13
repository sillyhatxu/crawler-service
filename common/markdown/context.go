package markdown

type RowContent struct {
	Content   string
	LabelType *TextType
	Label     string
	Index     *int // just NumberList
	TabCount  *int
}

type Option func(*RowContent)

func Content(content string) Option {
	return func(c *RowContent) {
		c.Content = content
	}
}

func LabelType(labelType TextType) Option {
	return func(c *RowContent) {
		c.LabelType = &labelType
	}
}

func Label(label string) Option {
	return func(c *RowContent) {
		c.Label = label
	}
}

func Index(index int) Option {
	return func(c *RowContent) {
		c.Index = &index
	}
}

func TabCount(tabCount int) Option {
	return func(c *RowContent) {
		c.TabCount = &tabCount
	}
}
