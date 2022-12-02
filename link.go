package link

import "io"

// Link contains the "href" attribute and
// the texts inside a HTML anchor tag <a>.
type Link struct {
	Href string
	Text string
}

// Parse takes an HTML document as an argument
// and return all the valid links inside.
func Parse(r io.Reader) ([]Link, error) {
	return nil, nil
}
