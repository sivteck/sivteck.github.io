package main

import (
    "slices"
	"os"
    "html/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Post struct {
	Content template.HTML
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func main() {
	files, _ := os.ReadDir("./posts")

    var allPosts []Post

	for _, file := range files {
		postMDContent, _ := os.ReadFile("./posts/" + file.Name())
		md := []byte(postMDContent)
		postHTMLContent := string(mdToHTML(md))

        allPosts = append(allPosts, Post{template.HTML(postHTMLContent)})
	}

    template, _ := template.ParseFiles("the_blog.tmpl")

    indexFile, _ := os.Create("index.html")

    slices.Reverse(allPosts)
    template.Execute(indexFile, allPosts)
}
