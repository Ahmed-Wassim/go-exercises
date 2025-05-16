package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	// "fmt"
	// "strings"

	"golang.org/x/net/html"
)

var raw = `<!DOCTYPE html>
<html>
    <body>
    <h1>My First Heading</h1>
    <p>My first paragraph.</p>
    <p>HTML images are defined with the img tag:</p>
    <img src="xxx.jpg" width="104" height="142">
    </body>
</html>`

func main() {

	doc, err := html.Parse(bytes.NewReader([]byte(raw)))

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
	}

	words, pics := countWordsAndImages(doc)
	fmt.Printf("words: %d\nimages: %d\n", words, pics)
}

func countWordsAndImages(node *html.Node) (int, int) {
	var words, pics int

	visit(node, &words, &pics)

	return words, pics
}

func visit(n *html.Node, words, pics *int) {
	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}

}
