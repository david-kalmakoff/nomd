package nomd_test

import (
	"testing"

	"gitlab.com/david-kalmakoff/nomd"
)

func TestAddElements(t *testing.T) {
	doc := nomd.NewDocument()

	els := nomd.Elements{
		nomd.NewElement("h1", nil, nomd.Text("testing")),
		nomd.NewElement("h2", nil, nomd.Text("test")),
	}
	doc.AddBodyElements(els)

	els = nomd.Elements{
		nomd.NewElement("title", nil, nomd.Text("Title")),
		nomd.NewElement("script", nil),
	}
	doc.AddHeadElements(els)

	got := doc.Render()
	want := `<!DOCTYPE html><html lang="eng"><head><title>Title</title><script></script></head><body><h1>testing</h1><h2>test</h2></body></html>`
	if got != want {
		t.Fatalf("\tF\tshould get expected doc:\nwant: %s\n got: %s", want, got)
	}
	t.Log("\tP\tshould get expected doc")
}

func TestAddElement(t *testing.T) {
	doc := nomd.NewDocument()

	el := nomd.NewElement("h1", nomd.Attributes{"class": "heading"}, nomd.Text("testing"))
	doc.AddBodyElement(el)

	el = nomd.NewElement("title", nil, nomd.Text("Title"))
	doc.AddHeadElement(el)

	got := doc.Render()
	want := `<!DOCTYPE html><html lang="eng"><head><title>Title</title></head><body><h1 class="heading">testing</h1></body></html>`
	if got != want {
		t.Fatalf("\tF\tshould get expected doc:\nwant: %s\n got: %s", want, got)
	}
	t.Log("\tP\tshould get expected doc")
}

func TestAddChildren(t *testing.T) {
	doc := nomd.NewElement("ul", nomd.Attributes{"class": "list"})

	els := []nomd.Renderer{
		nomd.NewElement("li", nil, nomd.Text("First")),
		nomd.NewElement("li", nil, nomd.Text("Second")),
	}
	doc.AddChildren(els...)

	got := doc.Render()
	want := `<ul class="list"><li>First</li><li>Second</li></ul>`
	if got != want {
		t.Fatalf("\tF\tshould get expected doc:\nwant: %s\n got: %s", want, got)
	}
	t.Log("\tP\tshould get expected doc")
}

func TestAddChild(t *testing.T) {
	doc := nomd.NewElement("div", nomd.Attributes{"class": "container"})

	el := nomd.NewElement("div", nomd.Attributes{"class": "row"})
	el1 := nomd.NewElement("div", nomd.Attributes{"class": "col"}, nomd.Text("Welcome"))
	el.AddChild(el1)
	doc.AddChild(el)

	got := doc.Render()
	want := `<div class="container"><div class="row"><div class="col">Welcome</div></div></div>`
	if got != want {
		t.Fatalf("\tF\tshould get expected doc:\nwant: %s\n got: %s", want, got)
	}
	t.Log("\tP\tshould get expected doc")
}

func TestAddChildrenNested(t *testing.T) {
	doc := nomd.NewElement("div", nomd.Attributes{"class": "container"})

	els := nomd.Elements{
		nomd.NewElement("div", nomd.Attributes{"class": "row"}),
		nomd.NewElement("div", nomd.Attributes{"class": "col"}, nomd.Text("Here")),
	}
	doc.AddChildrenNested(els...)

	got := doc.Render()
	want := `<div class="container"><div class="row"><div class="col">Here</div></div></div>`
	if got != want {
		t.Fatalf("\tF\tshould get expected doc:\nwant: %s\n got: %s", want, got)
	}
	t.Log("\tP\tshould get expected doc")
}
