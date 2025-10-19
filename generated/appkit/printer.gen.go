// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Printer] class.
var (
	printerClass     _PrinterClass
	printerClassOnce sync.Once
)

func getPrinterClass() _PrinterClass {
	printerClassOnce.Do(func() {
		printerClass = _PrinterClass{objc.GetClass("NSPrinter")}
	})
	return printerClass
}

type _PrinterClass struct {
	class objc.Class
}

// An interface definition for the [Printer] class.
type IPrinter interface {
	objectivec.IObject
}

// An object that describes a printer’s capabilities. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter

type Printer struct {
	objectivec.Object
}

// PrinterFrom constructs a [Printer] from an unsafe.Pointer.
//
// An object that describes a printer’s capabilities.
func PrinterFrom(ptr unsafe.Pointer) Printer {
	return Printer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (pc _PrinterClass) Alloc() Printer {
	rv := objc.Send[Printer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PrinterClass) New() Printer {
	rv := objc.Send[Printer](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Printer) Init() Printer {
	rv := objc.Send[Printer](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Printer) Autorelease() Printer {
	rv := objc.Send[Printer](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPrinter creates a new Printer instance.
func NewPrinter() Printer {
	return getPrinterClass().New()
}




