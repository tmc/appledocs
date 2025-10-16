
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [nextKeyView] class.
var nextKeyViewClass _nextKeyViewClass

func init() {
	nextKeyViewClass = _nextKeyViewClass{objc.GetClass("nextKeyView")}
}

type _nextKeyViewClass struct {
	objc.Class
}

// An interface definition for the [nextKeyView] class.
type InextKeyView interface {
	ID() objc.ID
}

type nextKeyView struct {
	id objc.ID
}

func nextKeyViewFrom(ptr unsafe.Pointer) nextKeyView {
	return nextKeyView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ nextKeyView) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _nextKeyViewClass) Alloc() nextKeyView {
	rv := objc.Send[nextKeyView](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _nextKeyViewClass) New() nextKeyView {
	rv := objc.Send[nextKeyView](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnextKeyView creates and returns a new initialized instance.
func NewnextKeyView() nextKeyView {
	return nextKeyViewClass.New()
}

// Init initializes the instance.
func (n_ nextKeyView) Init() nextKeyView {
	rv := objc.Send[nextKeyView](n_.ID(), selInit)
	return rv
}
