
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [indentationForDropOperation] class.
var indentationForDropOperationClass _indentationForDropOperationClass

func init() {
	indentationForDropOperationClass = _indentationForDropOperationClass{objc.GetClass("indentationForDropOperation")}
}

type _indentationForDropOperationClass struct {
	objc.Class
}

// An interface definition for the [indentationForDropOperation] class.
type IindentationForDropOperation interface {
	ID() objc.ID
}

type indentationForDropOperation struct {
	id objc.ID
}

func indentationForDropOperationFrom(ptr unsafe.Pointer) indentationForDropOperation {
	return indentationForDropOperation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ indentationForDropOperation) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _indentationForDropOperationClass) Alloc() indentationForDropOperation {
	rv := objc.Send[indentationForDropOperation](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _indentationForDropOperationClass) New() indentationForDropOperation {
	rv := objc.Send[indentationForDropOperation](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewindentationForDropOperation creates and returns a new initialized instance.
func NewindentationForDropOperation() indentationForDropOperation {
	return indentationForDropOperationClass.New()
}

// Init initializes the instance.
func (i_ indentationForDropOperation) Init() indentationForDropOperation {
	rv := objc.Send[indentationForDropOperation](i_.ID(), selInit)
	return rv
}
