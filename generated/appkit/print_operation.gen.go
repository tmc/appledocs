// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PrintOperation] class.
var (
	printOperationClass     _PrintOperationClass
	printOperationClassOnce sync.Once
)

func getPrintOperationClass() _PrintOperationClass {
	printOperationClassOnce.Do(func() {
		printOperationClass = _PrintOperationClass{objc.GetClass("NSPrintOperation")}
	})
	return printOperationClass
}

type _PrintOperationClass struct {
	class objc.Class
}

// An interface definition for the [PrintOperation] class.
type IPrintOperation interface {
	objectivec.IObject
}

// An object that controls operations that generate Encapsulated PostScript (EPS) code, Portable Document Format (PDF) code, or print jobs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation
type PrintOperation struct {
	objectivec.Object
}

// PrintOperationFrom constructs a [PrintOperation] from an unsafe.Pointer.
//
// An object that controls operations that generate Encapsulated PostScript (EPS) code, Portable Document Format (PDF) code, or print jobs.
func PrintOperationFrom(ptr unsafe.Pointer) PrintOperation {
	return PrintOperation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PrintOperationClass) Alloc() PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PrintOperationClass) New() PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PrintOperation) Init() PrintOperation {
	rv := objc.Send[PrintOperation](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PrintOperation) Autorelease() PrintOperation {
	rv := objc.Send[PrintOperation](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPrintOperation creates a new PrintOperation instance.
func NewPrintOperation() PrintOperation {
	return getPrintOperationClass().New()
}




