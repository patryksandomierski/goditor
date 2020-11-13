package ui

import (
	"log"
	// "reflect"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

// TextArea custom code area
type TextArea struct {
	widget.Entry
}

// TypedKey override
func (ta *TextArea) TypedKey(key *fyne.KeyEvent) {
    if key.Name == fyne.KeyTab {
		log.Println("TAB?") // check fyne/window.go KeyTab
    } else {
		log.Println("pressed: " + key.Name)
		ta.Entry.TypedKey(key)
    }
}

// NewTextArea creates new text area
func NewTextArea() *TextArea {
	e := &TextArea{}
	e.ExtendBaseWidget(e)
	e.MultiLine = true
	// v := reflect.ValueOf(e.Entry)
	// ptr := v.FieldByName("ReadOnly")
	//  v.MethodByName("textStyle").Call([]reflect.Value{})
	// v = reflect.Indirect(ptr)
	// // v2 := reflect.Indirect(ptr2)
	// log.Println(v)
	// log.Println(v2)
	return e
}