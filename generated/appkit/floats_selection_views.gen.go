
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [floatsSelectionViews] class.
var floatsSelectionViewsClass _floatsSelectionViewsClass

func init() {
	floatsSelectionViewsClass = _floatsSelectionViewsClass{objc.GetClass("floatsSelectionViews")}
}

type _floatsSelectionViewsClass struct {
	objc.Class
}

// An interface definition for the [floatsSelectionViews] class.
type IfloatsSelectionViews interface {
	ID() objc.ID
}

type floatsSelectionViews struct {
	id objc.ID
}

func floatsSelectionViewsFrom(ptr unsafe.Pointer) floatsSelectionViews {
	return floatsSelectionViews{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (f_ floatsSelectionViews) ID() objc.ID {
	return f_.id
}

// Alloc allocates a new instance without initialization.
func (fc _floatsSelectionViewsClass) Alloc() floatsSelectionViews {
	rv := objc.Send[floatsSelectionViews](objc.ID(fc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (fc _floatsSelectionViewsClass) New() floatsSelectionViews {
	rv := objc.Send[floatsSelectionViews](objc.ID(fc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewfloatsSelectionViews creates and returns a new initialized instance.
func NewfloatsSelectionViews() floatsSelectionViews {
	return floatsSelectionViewsClass.New()
}

// Init initializes the instance.
func (f_ floatsSelectionViews) Init() floatsSelectionViews {
	rv := objc.Send[floatsSelectionViews](f_.ID(), selInit)
	return rv
}
