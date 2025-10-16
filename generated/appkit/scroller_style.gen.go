
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [scrollerStyle] class.
var scrollerStyleClass _scrollerStyleClass

func init() {
	scrollerStyleClass = _scrollerStyleClass{objc.GetClass("scrollerStyle")}
}

type _scrollerStyleClass struct {
	objc.Class
}

// An interface definition for the [scrollerStyle] class.
type IscrollerStyle interface {
	ID() objc.ID
}

type scrollerStyle struct {
	id objc.ID
}

func scrollerStyleFrom(ptr unsafe.Pointer) scrollerStyle {
	return scrollerStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ scrollerStyle) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _scrollerStyleClass) Alloc() scrollerStyle {
	rv := objc.Send[scrollerStyle](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _scrollerStyleClass) New() scrollerStyle {
	rv := objc.Send[scrollerStyle](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewscrollerStyle creates and returns a new initialized instance.
func NewscrollerStyle() scrollerStyle {
	return scrollerStyleClass.New()
}

// Init initializes the instance.
func (s_ scrollerStyle) Init() scrollerStyle {
	rv := objc.Send[scrollerStyle](s_.ID(), selInit)
	return rv
}
