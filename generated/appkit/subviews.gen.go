
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [subviews] class.
var subviewsClass _subviewsClass

func init() {
	subviewsClass = _subviewsClass{objc.GetClass("subviews")}
}

type _subviewsClass struct {
	objc.Class
}

// An interface definition for the [subviews] class.
type Isubviews interface {
	ID() objc.ID
}

type subviews struct {
	id objc.ID
}

func subviewsFrom(ptr unsafe.Pointer) subviews {
	return subviews{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ subviews) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _subviewsClass) Alloc() subviews {
	rv := objc.Send[subviews](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _subviewsClass) New() subviews {
	rv := objc.Send[subviews](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newsubviews creates and returns a new initialized instance.
func Newsubviews() subviews {
	return subviewsClass.New()
}

// Init initializes the instance.
func (s_ subviews) Init() subviews {
	rv := objc.Send[subviews](s_.ID(), selInit)
	return rv
}
