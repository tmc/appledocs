
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [needsLayout] class.
var needsLayoutClass _needsLayoutClass

func init() {
	needsLayoutClass = _needsLayoutClass{objc.GetClass("needsLayout")}
}

type _needsLayoutClass struct {
	objc.Class
}

// An interface definition for the [needsLayout] class.
type IneedsLayout interface {
	ID() objc.ID
}

type needsLayout struct {
	id objc.ID
}

func needsLayoutFrom(ptr unsafe.Pointer) needsLayout {
	return needsLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ needsLayout) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _needsLayoutClass) Alloc() needsLayout {
	rv := objc.Send[needsLayout](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _needsLayoutClass) New() needsLayout {
	rv := objc.Send[needsLayout](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewneedsLayout creates and returns a new initialized instance.
func NewneedsLayout() needsLayout {
	return needsLayoutClass.New()
}

// Init initializes the instance.
func (n_ needsLayout) Init() needsLayout {
	rv := objc.Send[needsLayout](n_.ID(), selInit)
	return rv
}
