
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [screens] class.
var screensClass _screensClass

func init() {
	screensClass = _screensClass{objc.GetClass("screens")}
}

type _screensClass struct {
	objc.Class
}

// An interface definition for the [screens] class.
type Iscreens interface {
	ID() objc.ID
}

type screens struct {
	id objc.ID
}

func screensFrom(ptr unsafe.Pointer) screens {
	return screens{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ screens) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _screensClass) Alloc() screens {
	rv := objc.Send[screens](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _screensClass) New() screens {
	rv := objc.Send[screens](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newscreens creates and returns a new initialized instance.
func Newscreens() screens {
	return screensClass.New()
}

// Init initializes the instance.
func (s_ screens) Init() screens {
	rv := objc.Send[screens](s_.ID(), selInit)
	return rv
}
