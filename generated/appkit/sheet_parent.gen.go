
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [sheetParent] class.
var sheetParentClass _sheetParentClass

func init() {
	sheetParentClass = _sheetParentClass{objc.GetClass("sheetParent")}
}

type _sheetParentClass struct {
	objc.Class
}

// An interface definition for the [sheetParent] class.
type IsheetParent interface {
	ID() objc.ID
}

type sheetParent struct {
	id objc.ID
}

func sheetParentFrom(ptr unsafe.Pointer) sheetParent {
	return sheetParent{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ sheetParent) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _sheetParentClass) Alloc() sheetParent {
	rv := objc.Send[sheetParent](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _sheetParentClass) New() sheetParent {
	rv := objc.Send[sheetParent](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewsheetParent creates and returns a new initialized instance.
func NewsheetParent() sheetParent {
	return sheetParentClass.New()
}

// Init initializes the instance.
func (s_ sheetParent) Init() sheetParent {
	rv := objc.Send[sheetParent](s_.ID(), selInit)
	return rv
}
