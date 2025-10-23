// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Printer] class.
var (
	PrinterClass     _PrinterClass
	PrinterClassOnce sync.Once
)

func getPrinterClass() _PrinterClass {
	PrinterClassOnce.Do(func() {
		PrinterClass = _PrinterClass{objc.GetClass("NSPrinter")}
	})
	return PrinterClass
}

type _PrinterClass struct {
	class objc.Class
}

// An interface definition for the [Printer] class.
type IPrinter interface {
	objectivec.IObject
	// properties:
	DeviceDescription() objc.IObject /* cross-framework: DeviceDescriptionKey */
	SetDeviceDescription(value objc.IObject /* cross-framework: DeviceDescriptionKey */)
	LanguageLevel() int /* primitive/slice/pointer. */
	SetLanguageLevel(value int /* primitive/slice/pointer. */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
	// methods:
}

// An object that describes a printer’s capabilities.
//
// provides information about a printer; it does not modify printer attributes or control a printing job. A printer object can be constructed by specifying either the printer name or the make and model of an available printer. Typically, Cocoa apps don’t create objects; instead, the printing system uses these objects to support the printing jobs and when it shows users a list of printers.


// An object that describes a printer’s capabilities.
//
// [Full Topic]
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



// A dictionary of keys and values that describe the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/devicedescription
func (p_ Printer) DeviceDescription() objc.IObject /* cross-framework: DeviceDescriptionKey */ {
	rv := objc.Send[DeviceDescriptionKey](p_.ID, objc.Sel("deviceDescription"))
	return rv
}


// A dictionary of keys and values that describe the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/devicedescription
func (p_ Printer) SetDeviceDescription(value objc.IObject /* cross-framework: DeviceDescriptionKey */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDeviceDescription:"), value)
}


// The PostScript language level recognized by the printer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/languagelevel
func (p_ Printer) LanguageLevel() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("languageLevel"))
	return rv
}


// The PostScript language level recognized by the printer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/languagelevel
func (p_ Printer) SetLanguageLevel(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLanguageLevel:"), value)
}


// The printer’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/name
func (p_ Printer) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("name"))
	return rv
}


// The printer’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/name
func (p_ Printer) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setName:"), value)
}


// A description of the printer’s make and model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/type
func (p_ Printer) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("type"))
	return rv
}


// A description of the printer’s make and model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprinter/type
func (p_ Printer) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), value)
}



