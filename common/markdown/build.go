package markdown

import (
	"bytes"
	"errors"
	"fmt"
)

type Builder struct {
	Mapping map[string]TextType
	Rows    []RowContent
}

func CreateBuilder(mapping map[string]TextType) *Builder {
	return &Builder{Mapping: mapping}
}

func (b *Builder) AddContent(opts ...Option) *Builder {
	//default
	rowContent := &RowContent{
		Content: "",
		Label:   "",
	}
	for _, opt := range opts {
		opt(rowContent)
	}
	b.Rows = append(b.Rows, *rowContent)
	return b
}

func (b *Builder) Build() (bytes.Buffer, error) {
	var buffer bytes.Buffer
	if b.Rows == nil {
		return buffer, errors.New("rows is not defined")
	}
	for _, row := range b.Rows {
		buffer.WriteString(b.write(row))
	}
	return buffer, nil
}

func (b *Builder) write(row RowContent) string {
	var buffer bytes.Buffer
	tab := ""
	if row.TabCount != nil {
		for i := 0; i < *row.TabCount; i++ {
			tab += "\t" //Tab key
		}
	}
	textType := b.getTextType(row)
	switch textType {
	case Text:
		buffer.WriteString(fmt.Sprintf("%s%s\n", tab, row.Content))
	case H1:
		buffer.WriteString(fmt.Sprintf("%s# %s\n\n", tab, row.Content))
	case H2:
		buffer.WriteString(fmt.Sprintf("\n%s## %s\n\n", tab, row.Content))
	case H3:
		buffer.WriteString(fmt.Sprintf("\n%s### %s\n\n", tab, row.Content))
	case H4:
		buffer.WriteString(fmt.Sprintf("\n%s#### %s\n\n", tab, row.Content))
	case Italic:
		buffer.WriteString(fmt.Sprintf("%s*%s*", tab, row.Content))
	case Bold:
		buffer.WriteString(fmt.Sprintf("%s**%s**", tab, row.Content))
	case Monospace:
		buffer.WriteString(fmt.Sprintf("%s`%s`", tab, row.Content))
	case Itemize:
		buffer.WriteString(fmt.Sprintf("\n%s* %s\n", tab, row.Content))
	case BlockQuotes:
		buffer.WriteString(fmt.Sprintf("\n%s> %s\n", tab, row.Content))
	case NumberList:
		number := 0
		if row.Index != nil {
			number = *row.Index
		}
		buffer.WriteString(fmt.Sprintf("%d. %s\n", number, row.Content))
	case CodeBlockQuotes:
		buffer.WriteString("```\n")
		buffer.WriteString(fmt.Sprintf("%s\n", row.Content))
		buffer.WriteString("```\n")
	default:
		buffer.WriteString(fmt.Sprintf("%s%s", tab, row.Content))
	}
	return buffer.String()
}

func (b *Builder) getTextType(row RowContent) TextType {
	if row.LabelType != nil {
		return *row.LabelType
	}
	if b.Mapping == nil {
		return Unknown
	}
	result, ok := b.Mapping[row.Label]
	if ok {
		return result
	}
	return Unknown
}
