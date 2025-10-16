
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Browser] class.
var BrowserClass _BrowserClass

func init() {
	BrowserClass = _BrowserClass{objc.GetClass("NSBrowser")}
}

type _BrowserClass struct {
	objc.Class
}

// An interface definition for the [Browser] class.
type IBrowser interface {
	ID() objc.ID
}

type Browser struct {
	id objc.ID
}

func BrowserFrom(ptr unsafe.Pointer) Browser {
	return Browser{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ Browser) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _BrowserClass) Alloc() Browser {
	rv := objc.Send[Browser](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _BrowserClass) New() Browser {
	rv := objc.Send[Browser](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewBrowser creates and returns a new initialized instance.
func NewBrowser() Browser {
	return BrowserClass.New()
}

// Init initializes the instance.
func (b_ Browser) Init() Browser {
	rv := objc.Send[Browser](b_.ID(), selInit)
	return rv
}
