
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ColorList] class.
var ColorListClass _ColorListClass

func init() {
	ColorListClass = _ColorListClass{objc.GetClass("NSColorList")}
}

type _ColorListClass struct {
	objc.Class
}

// An interface definition for the [ColorList] class.
type IColorList interface {
	ID() objc.ID
}

type ColorList struct {
	id objc.ID
}

func ColorListFrom(ptr unsafe.Pointer) ColorList {
	return ColorList{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ ColorList) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _ColorListClass) Alloc() ColorList {
	rv := objc.Send[ColorList](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _ColorListClass) New() ColorList {
	rv := objc.Send[ColorList](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewColorList creates and returns a new initialized instance.
func NewColorList() ColorList {
	return ColorListClass.New()
}

// Init initializes the instance.
func (c_ ColorList) Init() ColorList {
	rv := objc.Send[ColorList](c_.ID(), selInit)
	return rv
}
