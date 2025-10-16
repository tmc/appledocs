
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [attributedTitle] class.
var attributedTitleClass _attributedTitleClass

func init() {
	attributedTitleClass = _attributedTitleClass{objc.GetClass("attributedTitle")}
}

type _attributedTitleClass struct {
	objc.Class
}

// An interface definition for the [attributedTitle] class.
type IattributedTitle interface {
	ID() objc.ID
}

type attributedTitle struct {
	id objc.ID
}

func attributedTitleFrom(ptr unsafe.Pointer) attributedTitle {
	return attributedTitle{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ attributedTitle) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _attributedTitleClass) Alloc() attributedTitle {
	rv := objc.Send[attributedTitle](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _attributedTitleClass) New() attributedTitle {
	rv := objc.Send[attributedTitle](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewattributedTitle creates and returns a new initialized instance.
func NewattributedTitle() attributedTitle {
	return attributedTitleClass.New()
}

// Init initializes the instance.
func (a_ attributedTitle) Init() attributedTitle {
	rv := objc.Send[attributedTitle](a_.ID(), selInit)
	return rv
}
