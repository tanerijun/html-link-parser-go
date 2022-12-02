package parser

import (
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	// test1.html: Parse link correctly
	f, err := os.Open("testdata/test1.html")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	links, err := Parse(f)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(links) != 1 {
		t.Errorf("expected len: 1, got: %d", len(links))
	}

	expected := "/other-page"
	if href := links[0].Href; href != expected {
		t.Errorf("expected href: %s, got: %s", expected, href)
	}

	expected = "A link to another page"
	if text := links[0].Text; text != expected {
		t.Errorf(`expected text: "%s"\ngot: "%s"`, expected, text)
	}

	f.Close()

	// test2.html: Parse multiple links
	f2, err := os.Open("testdata/test2.html")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	links, err = Parse(f2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(links) != 2 {
		t.Errorf("expected len: 2, got: %d", len(links))
	}

	// Check if the texts are parsed correctly
	expected = "Check me out on twitter"
	if text := links[0].Text; text != expected {
		t.Errorf("expected text: %s\ngot: %s", expected, text)
	}

	expected = "Gophercises is on Github!"
	if text := links[1].Text; text != expected {
		t.Errorf("expected text: %s\ngot: %s", expected, text)
	}

	f2.Close()

	// test3.html: Parse nested links
	f3, err := os.Open("testdata/test3.html")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	links, err = Parse(f3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(links) != 3 {
		t.Errorf("expected len: 3, got: %d", len(links))
	}

	f3.Close()

	// test4.html: Parse link with comment
	f4, err := os.Open("testdata/test4.html")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	links, err = Parse(f4)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(links) != 1 {
		t.Errorf("expected len: 1, got: %d", len(links))
	}

	// Should ignore comments
	expected = "dog cat"
	if text := links[0].Text; text != expected {
		t.Errorf("expected text: %s\ngot: %s", expected, text)
	}

	f4.Close()
}
