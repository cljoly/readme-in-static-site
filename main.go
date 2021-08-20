package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"github.com/yuin/goldmark/text"
)

var mdText []byte = []byte(`<!-- C2 -->
# Title
<!-- COMMENT -->
Some text
- list
- of
- stuff`)

// commentValue extract the content of an html comment starting with “<!--” and ending with “-->”
func commentValue(h *ast.HTMLBlock) (string, error) {
	if h.HTMLBlockType != 2 {
		return "", errors.New("Not an HTML comment")
	}

	lines := h.Lines()
	firstSeg := lines.At(0)
	lastSeg := lines.At(lines.Len() - 1)
	innerComment := firstSeg.WithStop(lastSeg.Stop - 4)
	innerComment = innerComment.WithStart(firstSeg.Start + 4)
	return string(innerComment.Value(mdText)), nil
}

func walker(n ast.Node, entering bool) (ast.WalkStatus, error) {
	// fmt.Println("--------------", entering)

	// n.Dump(mdText, 2)
	// fmt.Printf("%+v\n", n.Kind() == ast.KindHTMLBlock)
	if n.Kind() == ast.KindHTMLBlock {
		h := n.(*ast.HTMLBlock)
		h.Text(mdText)

		h.Dump(mdText, 2)
		text := string(h.Text(mdText))
		fmt.Println(">>", text, n.IsRaw(), n.ChildCount())
		nx := h.NextSibling()
		nx.Dump(mdText, 2)
		s, err := commentValue(h)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf(">>>>--- “%v”\n", s)
		// t := ast.NewTextSegment(innerComment)
		// h.InsertAfter(h, h.Parent(), t)
	}
	// if n.Kind() == ast.KindParagraph {
	// text := string(n.Text(mdText))
	// fmt.Println("@@@", text, n.IsRaw(), n.ChildCount())
	// }
	// if n.Kind() == ast.KindListItem {
	// n.Dump(mdText, 2)
	// text := string(n.Text(mdText))
	// fmt.Println("**", text)
	// }
	return ast.WalkContinue, nil
}

func main() {
	r := text.NewReader(mdText)

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)
	parser := md.Parser()

	doc := parser.Parse(r)
	fmt.Printf("%+v\n", doc)

	// doc.Dump(mdText, 2)
	ast.Walk(doc, walker)
	// doc.Dump(mdText, 3)

	var b bytes.Buffer
	md.Renderer().Render(&b, mdText, doc)

	fmt.Println()
	fmt.Println(b.String())
	/*
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(buf)
	*/
}
