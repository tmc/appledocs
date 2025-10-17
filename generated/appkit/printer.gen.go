
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Printer] class.
var PrinterClass _PrinterClass

func init() {
	PrinterClass = _PrinterClass{objc.GetClass("NSPrinter")}
}

type _PrinterClass struct {
	objc.Class
}

// An interface definition for the [Printer] class.
type IPrinter interface {
	ID() objc.ID
}

type Printer struct {
	id objc.ID
}

func PrinterFrom(ptr unsafe.Pointer) Printer {
	return Printer{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ Printer) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PrinterClass) Alloc() Printer {
	rv := objc.Send[Printer](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PrinterClass) New() Printer {
	rv := objc.Send[Printer](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPrinter creates and returns a new initialized instance.
func NewPrinter() Printer {
	return PrinterClass.New()
}

// Init initializes the instance.
func (p_ Printer) Init() Printer {
	rv := objc.Send[Printer](p_.ID(), selInit)
	return rv
}
