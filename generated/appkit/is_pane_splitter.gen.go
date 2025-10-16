
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isPaneSplitter] class.
var isPaneSplitterClass _isPaneSplitterClass

func init() {
	isPaneSplitterClass = _isPaneSplitterClass{objc.GetClass("isPaneSplitter")}
}

type _isPaneSplitterClass struct {
	objc.Class
}

// An interface definition for the [isPaneSplitter] class.
type IisPaneSplitter interface {
	ID() objc.ID
}

type isPaneSplitter struct {
	id objc.ID
}

func isPaneSplitterFrom(ptr unsafe.Pointer) isPaneSplitter {
	return isPaneSplitter{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isPaneSplitter) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isPaneSplitterClass) Alloc() isPaneSplitter {
	rv := objc.Send[isPaneSplitter](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isPaneSplitterClass) New() isPaneSplitter {
	rv := objc.Send[isPaneSplitter](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisPaneSplitter creates and returns a new initialized instance.
func NewisPaneSplitter() isPaneSplitter {
	return isPaneSplitterClass.New()
}

// Init initializes the instance.
func (i_ isPaneSplitter) Init() isPaneSplitter {
	rv := objc.Send[isPaneSplitter](i_.ID(), selInit)
	return rv
}
