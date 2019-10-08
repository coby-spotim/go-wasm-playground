package document

import "syscall/js"

type document struct {
	value js.Value
}

// Document returns a struct representing the Document interface in the Browser
func Document() *document {
	return &document{
		value: js.Global().Get("document"),
	}
}

func (d *document) All() js.Value {
	return d.value.Get("all")
}

func (d *document) Anchors() js.Value {
	return d.value.Get("anchors")
}

func (d *document) Body() js.Value {
	return d.value.Get("body")
}

func (d *document) CharacterSet() string {
	return d.value.Get("characterSet").String()
}

func (d *document) CompatMode() string {
	return d.value.Get("compatMode").String()
}

func (d *document) ContentType() string {
	return d.value.Get("contentType").String()
}

func (d *document) DocType() string {
	return d.value.Get("docType").String()
}

func (d *document) DocumentElement() js.Value {
	return d.value.Get("documentElement")
}

func (d *document) DocumentURI() string {
	return d.value.Get("documentURI").String()
}

func (d *document) Embeds() js.Value {
	return d.value.Get("embeds")
}

func (d *document) Fonts() js.Value {
	return d.value.Get("fonts")
}

func (d *document) Forms() js.Value {
	return d.value.Get("forms")
}

func (d *document) Head() js.Value {
	return d.value.Get("head")
}

func (d *document) Hidden() bool {
	return d.value.Get("hidden").Bool()
}

func (d *document) Images() js.Value {
	return d.value.Get("images")
}

func (d *document) Implementation() js.Value {
	return d.value.Get("implementation")
}

func (d *document) LastStyleSheetSet() js.Value {
	return d.value.Get("lastStyleSheetSet")
}

func (d *document) Links() js.Value {
	return d.value.Get("links")
}

func (d *document) MozSyntheticDocument() bool {
	return d.value.Get("mozSyntheticDocument").Bool()
}

func (d *document) Plugins() js.Value {
	return d.value.Get("plugins")
}

func (d *document) FeaturePolicy() js.Value {
	return d.value.Get("featurePolicy")
}

func (d *document) PreferredStyleSheetSet() js.Value {
	return d.value.Get("preferredStyleSheetSet")
}

func (d *document) Scripts() js.Value {
	return d.value.Get("scripts")
}

func (d *document) ScrollingElement() js.Value {
	return d.value.Get("scrollingElement")
}

func (d *document) SelectedStyleSheetSet() js.Value {
	return d.value.Get("selectedStyleSheetSet")
}

func (d *document) StyleSheetSets() js.Value {
	return d.value.Get("styleSheetSets")
}

func (d *document) Timeline() js.Value {
	return d.value.Get("timeline")
}

func (d *document) VisibilityState() js.Value {
	return d.value.Get("visibilityState")
}
