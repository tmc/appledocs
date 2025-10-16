
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PrintOperation] class.
var PrintOperationClass _PrintOperationClass

func init() {
	PrintOperationClass = _PrintOperationClass{objc.GetClass("NSPrintOperation")}
}

type _PrintOperationClass struct {
	objc.Class
}

// An interface definition for the [PrintOperation] class.
type IPrintOperation interface {
	ID() objc.ID
}

type PrintOperation struct {
	id objc.ID
}

func PrintOperationFrom(ptr unsafe.Pointer) PrintOperation {
	return PrintOperation{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PrintOperation) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PrintOperationClass) Alloc() PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PrintOperationClass) New() PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPrintOperation creates and returns a new initialized instance.
func NewPrintOperation() PrintOperation {
	return PrintOperationClass.New()
}

// Init initializes the instance.
func (p_ PrintOperation) Init() PrintOperation {
	rv := objc.Send[PrintOperation](p_.ID(), selInit)
	return rv
}
