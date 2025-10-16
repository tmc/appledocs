
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [rulerViewClass] class.
var rulerViewClassClass _rulerViewClassClass

func init() {
	rulerViewClassClass = _rulerViewClassClass{objc.GetClass("rulerViewClass")}
}

type _rulerViewClassClass struct {
	objc.Class
}

// An interface definition for the [rulerViewClass] class.
type IrulerViewClass interface {
	ID() objc.ID
}

type rulerViewClass struct {
	id objc.ID
}

func rulerViewClassFrom(ptr unsafe.Pointer) rulerViewClass {
	return rulerViewClass{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ rulerViewClass) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _rulerViewClassClass) Alloc() rulerViewClass {
	rv := objc.Send[rulerViewClass](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _rulerViewClassClass) New() rulerViewClass {
	rv := objc.Send[rulerViewClass](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewrulerViewClass creates and returns a new initialized instance.
func NewrulerViewClass() rulerViewClass {
	return rulerViewClassClass.New()
}

// Init initializes the instance.
func (r_ rulerViewClass) Init() rulerViewClass {
	rv := objc.Send[rulerViewClass](r_.ID(), selInit)
	return rv
}
