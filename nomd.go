package nomd

import (
	"fmt"
	"strings"
)

// ============================================================================
// Document

type document struct {
	head []element
	body []element
}

func NewDocument() document {
	return document{
		head: make([]element, 0),
		body: make([]element, 0),
	}
}

func (d *document) AddHeadElement(e element) {
	d.head = append(d.head, e)
}

func (d *document) AddHeadElements(e []element) {
	d.head = append(d.head, e...)
}

func (d *document) AddBodyElement(e element) {
	d.body = append(d.body, e)
}

func (d *document) AddBodyElements(e []element) {
	d.body = append(d.body, e...)
}

func (d *document) build() strings.Builder {
	ret := strings.Builder{}

	// Build opening tag of document
	ret.WriteString("<!DOCTYPE html>")
	ret.WriteString(`<html lang="eng">`)

	// Build head of document
	ret.WriteString("<head>")
	for _, e := range d.head {
		ret.WriteString(e.Render())
	}
	ret.WriteString("</head>")

	// Build body of document
	ret.WriteString("<body>")
	for _, e := range d.body {
		ret.WriteString(e.Render())
	}
	ret.WriteString("</body>")

	// Build closing tag of document
	ret.WriteString("</html>")

	return ret
}

func (d *document) Render() string {
	ret := d.build()
	return ret.String()
}

// ============================================================================
// Element

type Renderer interface {
	Render() string
}

type Attributes = map[string]string

type Elements = []element

type element struct {
	tag        string
	attributes Attributes
	children   []Renderer
}

func NewElement(tag string, attributes map[string]string, children ...Renderer) element {
	el := element{
		tag:        tag,
		attributes: make(map[string]string),
		children:   make([]Renderer, 0),
	}

	if attributes != nil {
		el.attributes = attributes
	}

	if len(children) > 0 {
		el.children = children
	}

	return el
}

func (e *element) AddChild(c Renderer) {
	e.children = append(e.children, c)
}

func (e *element) AddChildren(c ...Renderer) {
	e.children = append(e.children, c...)
}

func (e *element) AddChildrenNested(c ...element) {
	for i := len(c) - 1; i >= 0; i-- {
		if i == 0 {
			e.children = append(e.children, c[i])
			continue
		}

		c[i-1].children = append(c[i-1].children, c[i])
	}

}

func (e element) build() strings.Builder {
	ret := strings.Builder{}

	// Build opening tag of element
	ret.WriteString("<")
	ret.WriteString(e.tag)

	for key, value := range e.attributes {
		ret.WriteString(" ")
		ret.WriteString(key)
		ret.WriteString(`="`)
		ret.WriteString(value)
		ret.WriteString(`"`)
	}
	ret.WriteString(">")

	// Build innerHTML of element
	for _, child := range e.children {
		ret.WriteString(child.Render())
	}

	// Build closing tag of element
	ret.WriteString("</")
	ret.WriteString(e.tag)
	ret.WriteString(">")

	return ret
}

func (e element) Render() string {
	ret := e.build()
	return ret.String()
}

// ============================================================================
// Text

type Text string

func (t Text) Render() string {
	return fmt.Sprint(t)
}
