package main

import "github.com/russross/blackfriday"

type Markdown struct {
	text []byte
}

func NewMarkdown(text []byte) *Markdown {
	return &Markdown{text}
}

func (md *Markdown) ToHtml(title string, enableExtensions bool) []byte {
	htmlFlags := 0
	htmlFlags |= blackfriday.HTML_COMPLETE_PAGE
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_FRACTIONS
	htmlFlags |= blackfriday.HTML_SMARTYPANTS_LATEX_DASHES
	htmlFlags |= blackfriday.HTML_USE_SMARTYPANTS
	htmlFlags |= blackfriday.HTML_USE_XHTML
	renderer := blackfriday.HtmlRenderer(htmlFlags, title, "")

	extensions := 0
	if enableExtensions {
		extensions |= blackfriday.EXTENSION_AUTOLINK
		extensions |= blackfriday.EXTENSION_FENCED_CODE
		extensions |= blackfriday.EXTENSION_HEADER_IDS
		extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
		extensions |= blackfriday.EXTENSION_SPACE_HEADERS
		extensions |= blackfriday.EXTENSION_STRIKETHROUGH
		extensions |= blackfriday.EXTENSION_TABLES
	}
	return blackfriday.Markdown(md.text, renderer, extensions)
}

func (md *Markdown) ToPdf(outfile, title string, opts ...string) error {
	html := md.ToHtml(title, true)
	return HtmlToPdf(html, outfile, opts...)
}
