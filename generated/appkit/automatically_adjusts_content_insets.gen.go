
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [automaticallyAdjustsContentInsets] class.
var automaticallyAdjustsContentInsetsClass _automaticallyAdjustsContentInsetsClass

func init() {
	automaticallyAdjustsContentInsetsClass = _automaticallyAdjustsContentInsetsClass{objc.GetClass("automaticallyAdjustsContentInsets")}
}

type _automaticallyAdjustsContentInsetsClass struct {
	objc.Class
}

// An interface definition for the [automaticallyAdjustsContentInsets] class.
type IautomaticallyAdjustsContentInsets interface {
	ID() objc.ID
}

type automaticallyAdjustsContentInsets struct {
	id objc.ID
}

func automaticallyAdjustsContentInsetsFrom(ptr unsafe.Pointer) automaticallyAdjustsContentInsets {
	return automaticallyAdjustsContentInsets{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ automaticallyAdjustsContentInsets) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _automaticallyAdjustsContentInsetsClass) Alloc() automaticallyAdjustsContentInsets {
	rv := objc.Send[automaticallyAdjustsContentInsets](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _automaticallyAdjustsContentInsetsClass) New() automaticallyAdjustsContentInsets {
	rv := objc.Send[automaticallyAdjustsContentInsets](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewautomaticallyAdjustsContentInsets creates and returns a new initialized instance.
func NewautomaticallyAdjustsContentInsets() automaticallyAdjustsContentInsets {
	return automaticallyAdjustsContentInsetsClass.New()
}

// Init initializes the instance.
func (a_ automaticallyAdjustsContentInsets) Init() automaticallyAdjustsContentInsets {
	rv := objc.Send[automaticallyAdjustsContentInsets](a_.ID(), selInit)
	return rv
}
