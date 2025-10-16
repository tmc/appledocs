
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [wraps] class.
var wrapsClass _wrapsClass

func init() {
	wrapsClass = _wrapsClass{objc.GetClass("wraps")}
}

type _wrapsClass struct {
	objc.Class
}

// An interface definition for the [wraps] class.
type Iwraps interface {
	ID() objc.ID
}

type wraps struct {
	id objc.ID
}

func wrapsFrom(ptr unsafe.Pointer) wraps {
	return wraps{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ wraps) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _wrapsClass) Alloc() wraps {
	rv := objc.Send[wraps](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _wrapsClass) New() wraps {
	rv := objc.Send[wraps](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newwraps creates and returns a new initialized instance.
func Newwraps() wraps {
	return wrapsClass.New()
}

// Init initializes the instance.
func (w_ wraps) Init() wraps {
	rv := objc.Send[wraps](w_.ID(), selInit)
	return rv
}
