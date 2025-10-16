
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [nextValidKeyView] class.
var nextValidKeyViewClass _nextValidKeyViewClass

func init() {
	nextValidKeyViewClass = _nextValidKeyViewClass{objc.GetClass("nextValidKeyView")}
}

type _nextValidKeyViewClass struct {
	objc.Class
}

// An interface definition for the [nextValidKeyView] class.
type InextValidKeyView interface {
	ID() objc.ID
}

type nextValidKeyView struct {
	id objc.ID
}

func nextValidKeyViewFrom(ptr unsafe.Pointer) nextValidKeyView {
	return nextValidKeyView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ nextValidKeyView) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _nextValidKeyViewClass) Alloc() nextValidKeyView {
	rv := objc.Send[nextValidKeyView](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _nextValidKeyViewClass) New() nextValidKeyView {
	rv := objc.Send[nextValidKeyView](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnextValidKeyView creates and returns a new initialized instance.
func NewnextValidKeyView() nextValidKeyView {
	return nextValidKeyViewClass.New()
}

// Init initializes the instance.
func (n_ nextValidKeyView) Init() nextValidKeyView {
	rv := objc.Send[nextValidKeyView](n_.ID(), selInit)
	return rv
}
