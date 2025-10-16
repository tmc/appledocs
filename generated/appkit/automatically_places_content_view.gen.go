
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [automaticallyPlacesContentView] class.
var automaticallyPlacesContentViewClass _automaticallyPlacesContentViewClass

func init() {
	automaticallyPlacesContentViewClass = _automaticallyPlacesContentViewClass{objc.GetClass("automaticallyPlacesContentView")}
}

type _automaticallyPlacesContentViewClass struct {
	objc.Class
}

// An interface definition for the [automaticallyPlacesContentView] class.
type IautomaticallyPlacesContentView interface {
	ID() objc.ID
}

type automaticallyPlacesContentView struct {
	id objc.ID
}

func automaticallyPlacesContentViewFrom(ptr unsafe.Pointer) automaticallyPlacesContentView {
	return automaticallyPlacesContentView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ automaticallyPlacesContentView) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _automaticallyPlacesContentViewClass) Alloc() automaticallyPlacesContentView {
	rv := objc.Send[automaticallyPlacesContentView](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _automaticallyPlacesContentViewClass) New() automaticallyPlacesContentView {
	rv := objc.Send[automaticallyPlacesContentView](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewautomaticallyPlacesContentView creates and returns a new initialized instance.
func NewautomaticallyPlacesContentView() automaticallyPlacesContentView {
	return automaticallyPlacesContentViewClass.New()
}

// Init initializes the instance.
func (a_ automaticallyPlacesContentView) Init() automaticallyPlacesContentView {
	rv := objc.Send[automaticallyPlacesContentView](a_.ID(), selInit)
	return rv
}
