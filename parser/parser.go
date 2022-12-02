package parser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link contains the "href" attribute and
// the texts inside a HTML anchor tag <a>.
type Link struct {
	Href string
	Text string
}

// Parse takes an HTML document as an argument
// and return all the valid links inside.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	links := getLinks(doc)

	return links, nil
}

func getLinks(n *html.Node) []Link {
	if n.Type == html.ElementNode && n.Data == "a" {
		href := getHref(n)
		text := getText(n)

		return []Link{{href, text}}
	}

	var links []Link

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, getLinks(c)...)
	}

	return links
}

func getHref(n *html.Node) string {
	var href string

	for _, a := range n.Attr {
		if a.Key == "href" {
			href = a.Val
		}
	}

	return href
}

func getText(n *html.Node) string {
	if n.Type == html.CommentNode {
		return ""
	}

	if n.Type == html.TextNode {
		return n.Data
	}

	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getText(c)
	}

	return strings.Join(strings.Fields(text), " ")
}
