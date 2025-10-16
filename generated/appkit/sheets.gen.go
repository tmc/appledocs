
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [sheets] class.
var sheetsClass _sheetsClass

func init() {
	sheetsClass = _sheetsClass{objc.GetClass("sheets")}
}

type _sheetsClass struct {
	objc.Class
}

// An interface definition for the [sheets] class.
type Isheets interface {
	ID() objc.ID
}

type sheets struct {
	id objc.ID
}

func sheetsFrom(ptr unsafe.Pointer) sheets {
	return sheets{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ sheets) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _sheetsClass) Alloc() sheets {
	rv := objc.Send[sheets](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _sheetsClass) New() sheets {
	rv := objc.Send[sheets](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newsheets creates and returns a new initialized instance.
func Newsheets() sheets {
	return sheetsClass.New()
}

// Init initializes the instance.
func (s_ sheets) Init() sheets {
	rv := objc.Send[sheets](s_.ID(), selInit)
	return rv
}
