
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SliderTouchBarItem] class.
var SliderTouchBarItemClass _SliderTouchBarItemClass

func init() {
	SliderTouchBarItemClass = _SliderTouchBarItemClass{objc.GetClass("NSSliderTouchBarItem")}
}

type _SliderTouchBarItemClass struct {
	objc.Class
}

// An interface definition for the [SliderTouchBarItem] class.
type ISliderTouchBarItem interface {
	ID() objc.ID
}

type SliderTouchBarItem struct {
	id objc.ID
}

func SliderTouchBarItemFrom(ptr unsafe.Pointer) SliderTouchBarItem {
	return SliderTouchBarItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SliderTouchBarItem) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SliderTouchBarItemClass) Alloc() SliderTouchBarItem {
	rv := objc.Send[SliderTouchBarItem](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SliderTouchBarItemClass) New() SliderTouchBarItem {
	rv := objc.Send[SliderTouchBarItem](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSliderTouchBarItem creates and returns a new initialized instance.
func NewSliderTouchBarItem() SliderTouchBarItem {
	return SliderTouchBarItemClass.New()
}

// Init initializes the instance.
func (s_ SliderTouchBarItem) Init() SliderTouchBarItem {
	rv := objc.Send[SliderTouchBarItem](s_.ID(), selInit)
	return rv
}
