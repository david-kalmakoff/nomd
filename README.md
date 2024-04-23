# nomd
Build HTML strings without the need of templates.

# Install
```bash
go get github.com/david-kalmakoff/nomd
```

# Example
```go
doc := nomd.NewDocument()

heading := nomd.NewElement("h1", nomd.Attributes{"class": "heading"}, nomd.Text("testing"))
doc.AddBodyElements(heading)

title := nomd.NewElement("title", nil, nomd.Text("Title"))
doc.AddHeadElements(els)

str := doc.Render()
fmt.Println(str)
```
