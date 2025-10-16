
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [subtitle] class.
var subtitleClass _subtitleClass

func init() {
	subtitleClass = _subtitleClass{objc.GetClass("subtitle")}
}

type _subtitleClass struct {
	objc.Class
}

// An interface definition for the [subtitle] class.
type Isubtitle interface {
	ID() objc.ID
}

type subtitle struct {
	id objc.ID
}

func subtitleFrom(ptr unsafe.Pointer) subtitle {
	return subtitle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ subtitle) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _subtitleClass) Alloc() subtitle {
	rv := objc.Send[subtitle](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _subtitleClass) New() subtitle {
	rv := objc.Send[subtitle](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newsubtitle creates and returns a new initialized instance.
func Newsubtitle() subtitle {
	return subtitleClass.New()
}

// Init initializes the instance.
func (s_ subtitle) Init() subtitle {
	rv := objc.Send[subtitle](s_.ID(), selInit)
	return rv
}
