
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [enclosingScrollView] class.
var enclosingScrollViewClass _enclosingScrollViewClass

func init() {
	enclosingScrollViewClass = _enclosingScrollViewClass{objc.GetClass("enclosingScrollView")}
}

type _enclosingScrollViewClass struct {
	objc.Class
}

// An interface definition for the [enclosingScrollView] class.
type IenclosingScrollView interface {
	ID() objc.ID
}

type enclosingScrollView struct {
	id objc.ID
}

func enclosingScrollViewFrom(ptr unsafe.Pointer) enclosingScrollView {
	return enclosingScrollView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (e_ enclosingScrollView) ID() objc.ID {
	return e_.id
}

// Alloc allocates a new instance without initialization.
func (ec _enclosingScrollViewClass) Alloc() enclosingScrollView {
	rv := objc.Send[enclosingScrollView](objc.ID(ec.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ec _enclosingScrollViewClass) New() enclosingScrollView {
	rv := objc.Send[enclosingScrollView](objc.ID(ec.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewenclosingScrollView creates and returns a new initialized instance.
func NewenclosingScrollView() enclosingScrollView {
	return enclosingScrollViewClass.New()
}

// Init initializes the instance.
func (e_ enclosingScrollView) Init() enclosingScrollView {
	rv := objc.Send[enclosingScrollView](e_.ID(), selInit)
	return rv
}
