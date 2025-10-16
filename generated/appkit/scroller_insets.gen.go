
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [scrollerInsets] class.
var scrollerInsetsClass _scrollerInsetsClass

func init() {
	scrollerInsetsClass = _scrollerInsetsClass{objc.GetClass("scrollerInsets")}
}

type _scrollerInsetsClass struct {
	objc.Class
}

// An interface definition for the [scrollerInsets] class.
type IscrollerInsets interface {
	ID() objc.ID
}

type scrollerInsets struct {
	id objc.ID
}

func scrollerInsetsFrom(ptr unsafe.Pointer) scrollerInsets {
	return scrollerInsets{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ scrollerInsets) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _scrollerInsetsClass) Alloc() scrollerInsets {
	rv := objc.Send[scrollerInsets](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _scrollerInsetsClass) New() scrollerInsets {
	rv := objc.Send[scrollerInsets](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewscrollerInsets creates and returns a new initialized instance.
func NewscrollerInsets() scrollerInsets {
	return scrollerInsetsClass.New()
}

// Init initializes the instance.
func (s_ scrollerInsets) Init() scrollerInsets {
	rv := objc.Send[scrollerInsets](s_.ID(), selInit)
	return rv
}
