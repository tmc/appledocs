
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [fullScreenAccessoryView] class.
var fullScreenAccessoryViewClass _fullScreenAccessoryViewClass

func init() {
	fullScreenAccessoryViewClass = _fullScreenAccessoryViewClass{objc.GetClass("fullScreenAccessoryView")}
}

type _fullScreenAccessoryViewClass struct {
	objc.Class
}

// An interface definition for the [fullScreenAccessoryView] class.
type IfullScreenAccessoryView interface {
	ID() objc.ID
}

type fullScreenAccessoryView struct {
	id objc.ID
}

func fullScreenAccessoryViewFrom(ptr unsafe.Pointer) fullScreenAccessoryView {
	return fullScreenAccessoryView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ fullScreenAccessoryView) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _fullScreenAccessoryViewClass) Alloc() fullScreenAccessoryView {
	rv := objc.Send[fullScreenAccessoryView](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _fullScreenAccessoryViewClass) New() fullScreenAccessoryView {
	rv := objc.Send[fullScreenAccessoryView](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfullScreenAccessoryView creates and returns a new initialized instance.
func NewfullScreenAccessoryView() fullScreenAccessoryView {
	return fullScreenAccessoryViewClass.New()
}

// Init initializes the instance.
func (f_ fullScreenAccessoryView) Init() fullScreenAccessoryView {
	rv := objc.Send[fullScreenAccessoryView](f_.ID(), selInit)
	return rv
}
