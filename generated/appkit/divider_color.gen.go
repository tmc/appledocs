
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [dividerColor] class.
var dividerColorClass _dividerColorClass

func init() {
	dividerColorClass = _dividerColorClass{objc.GetClass("dividerColor")}
}

type _dividerColorClass struct {
	objc.Class
}

// An interface definition for the [dividerColor] class.
type IdividerColor interface {
	ID() objc.ID
}

type dividerColor struct {
	id objc.ID
}

func dividerColorFrom(ptr unsafe.Pointer) dividerColor {
	return dividerColor{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ dividerColor) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _dividerColorClass) Alloc() dividerColor {
	rv := objc.Send[dividerColor](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _dividerColorClass) New() dividerColor {
	rv := objc.Send[dividerColor](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdividerColor creates and returns a new initialized instance.
func NewdividerColor() dividerColor {
	return dividerColorClass.New()
}

// Init initializes the instance.
func (d_ dividerColor) Init() dividerColor {
	rv := objc.Send[dividerColor](d_.ID(), selInit)
	return rv
}
