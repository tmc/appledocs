
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [scrollsDynamically] class.
var scrollsDynamicallyClass _scrollsDynamicallyClass

func init() {
	scrollsDynamicallyClass = _scrollsDynamicallyClass{objc.GetClass("scrollsDynamically")}
}

type _scrollsDynamicallyClass struct {
	objc.Class
}

// An interface definition for the [scrollsDynamically] class.
type IscrollsDynamically interface {
	ID() objc.ID
}

type scrollsDynamically struct {
	id objc.ID
}

func scrollsDynamicallyFrom(ptr unsafe.Pointer) scrollsDynamically {
	return scrollsDynamically{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ scrollsDynamically) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _scrollsDynamicallyClass) Alloc() scrollsDynamically {
	rv := objc.Send[scrollsDynamically](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _scrollsDynamicallyClass) New() scrollsDynamically {
	rv := objc.Send[scrollsDynamically](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewscrollsDynamically creates and returns a new initialized instance.
func NewscrollsDynamically() scrollsDynamically {
	return scrollsDynamicallyClass.New()
}

// Init initializes the instance.
func (s_ scrollsDynamically) Init() scrollsDynamically {
	rv := objc.Send[scrollsDynamically](s_.ID(), selInit)
	return rv
}
