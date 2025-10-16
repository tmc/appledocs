
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [cornerRadius] class.
var cornerRadiusClass _cornerRadiusClass

func init() {
	cornerRadiusClass = _cornerRadiusClass{objc.GetClass("cornerRadius")}
}

type _cornerRadiusClass struct {
	objc.Class
}

// An interface definition for the [cornerRadius] class.
type IcornerRadius interface {
	ID() objc.ID
}

type cornerRadius struct {
	id objc.ID
}

func cornerRadiusFrom(ptr unsafe.Pointer) cornerRadius {
	return cornerRadius{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ cornerRadius) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _cornerRadiusClass) Alloc() cornerRadius {
	rv := objc.Send[cornerRadius](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _cornerRadiusClass) New() cornerRadius {
	rv := objc.Send[cornerRadius](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcornerRadius creates and returns a new initialized instance.
func NewcornerRadius() cornerRadius {
	return cornerRadiusClass.New()
}

// Init initializes the instance.
func (c_ cornerRadius) Init() cornerRadius {
	rv := objc.Send[cornerRadius](c_.ID(), selInit)
	return rv
}
