
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [truncatesLastVisibleLine] class.
var truncatesLastVisibleLineClass _truncatesLastVisibleLineClass

func init() {
	truncatesLastVisibleLineClass = _truncatesLastVisibleLineClass{objc.GetClass("truncatesLastVisibleLine")}
}

type _truncatesLastVisibleLineClass struct {
	objc.Class
}

// An interface definition for the [truncatesLastVisibleLine] class.
type ItruncatesLastVisibleLine interface {
	ID() objc.ID
}

type truncatesLastVisibleLine struct {
	id objc.ID
}

func truncatesLastVisibleLineFrom(ptr unsafe.Pointer) truncatesLastVisibleLine {
	return truncatesLastVisibleLine{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ truncatesLastVisibleLine) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _truncatesLastVisibleLineClass) Alloc() truncatesLastVisibleLine {
	rv := objc.Send[truncatesLastVisibleLine](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _truncatesLastVisibleLineClass) New() truncatesLastVisibleLine {
	rv := objc.Send[truncatesLastVisibleLine](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtruncatesLastVisibleLine creates and returns a new initialized instance.
func NewtruncatesLastVisibleLine() truncatesLastVisibleLine {
	return truncatesLastVisibleLineClass.New()
}

// Init initializes the instance.
func (t_ truncatesLastVisibleLine) Init() truncatesLastVisibleLine {
	rv := objc.Send[truncatesLastVisibleLine](t_.ID(), selInit)
	return rv
}
