
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [scrollerKnobStyle] class.
var scrollerKnobStyleClass _scrollerKnobStyleClass

func init() {
	scrollerKnobStyleClass = _scrollerKnobStyleClass{objc.GetClass("scrollerKnobStyle")}
}

type _scrollerKnobStyleClass struct {
	objc.Class
}

// An interface definition for the [scrollerKnobStyle] class.
type IscrollerKnobStyle interface {
	ID() objc.ID
}

type scrollerKnobStyle struct {
	id objc.ID
}

func scrollerKnobStyleFrom(ptr unsafe.Pointer) scrollerKnobStyle {
	return scrollerKnobStyle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ scrollerKnobStyle) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _scrollerKnobStyleClass) Alloc() scrollerKnobStyle {
	rv := objc.Send[scrollerKnobStyle](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _scrollerKnobStyleClass) New() scrollerKnobStyle {
	rv := objc.Send[scrollerKnobStyle](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewscrollerKnobStyle creates and returns a new initialized instance.
func NewscrollerKnobStyle() scrollerKnobStyle {
	return scrollerKnobStyleClass.New()
}

// Init initializes the instance.
func (s_ scrollerKnobStyle) Init() scrollerKnobStyle {
	rv := objc.Send[scrollerKnobStyle](s_.ID(), selInit)
	return rv
}
