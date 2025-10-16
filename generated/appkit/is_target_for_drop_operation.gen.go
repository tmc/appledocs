
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isTargetForDropOperation] class.
var isTargetForDropOperationClass _isTargetForDropOperationClass

func init() {
	isTargetForDropOperationClass = _isTargetForDropOperationClass{objc.GetClass("isTargetForDropOperation")}
}

type _isTargetForDropOperationClass struct {
	objc.Class
}

// An interface definition for the [isTargetForDropOperation] class.
type IisTargetForDropOperation interface {
	ID() objc.ID
}

type isTargetForDropOperation struct {
	id objc.ID
}

func isTargetForDropOperationFrom(ptr unsafe.Pointer) isTargetForDropOperation {
	return isTargetForDropOperation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isTargetForDropOperation) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isTargetForDropOperationClass) Alloc() isTargetForDropOperation {
	rv := objc.Send[isTargetForDropOperation](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isTargetForDropOperationClass) New() isTargetForDropOperation {
	rv := objc.Send[isTargetForDropOperation](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisTargetForDropOperation creates and returns a new initialized instance.
func NewisTargetForDropOperation() isTargetForDropOperation {
	return isTargetForDropOperationClass.New()
}

// Init initializes the instance.
func (i_ isTargetForDropOperation) Init() isTargetForDropOperation {
	rv := objc.Send[isTargetForDropOperation](i_.ID(), selInit)
	return rv
}
