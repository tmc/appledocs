// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	DeviceDescription() foundation.IDictionary
	LanguageLevel() int
	Name() foundation.foundation.INSString
	Type() PrinterTypeName


	

	// methods:
	PageSizeForPaper(paperName PrinterPaperName) corefoundation.CGSize


}





// Alloc allocates a new instance without initialization.
func (pc _PrinterClass) Alloc() Printer {
	rv := objc.Send[Printer](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates and returns a printer object initialized with the specified printer name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/init(name:)
func NewPrinterWithName(name foundation.foundation.INSString) Printer {
	rv := objc.Send[Printer](objc.ID(getPrinterClass().class), objc.Sel("printerWithName:"), name)
	return rv
}


// Creates and returns a printer object initialized to the first available printer with the specified make and model information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/init(type:)
func NewPrinterWithType(type_ PrinterTypeName) Printer {
	rv := objc.Send[Printer](objc.ID(getPrinterClass().class), objc.Sel("printerWithType:"), type_)
	return rv
}







// Creates and returns a printer object initialized with the specified printer name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/init(name:)
func (pc _PrinterClass) PrinterWithName(name foundation.foundation.INSString) IPrinter {
	rv := objc.Send[Printer](objc.ID(pc.class), objc.Sel("printerWithName:"), name)
	return rv
}


// Creates and returns a printer object initialized to the first available printer with the specified make and model information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/init(type:)
func (pc _PrinterClass) PrinterWithType(type_ PrinterTypeName) IPrinter {
	rv := objc.Send[Printer](objc.ID(pc.class), objc.Sel("printerWithType:"), type_)
	return rv
}


// Deprecated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/printerWithName:domain:includeUnavailable:
func (pc _PrinterClass) PrinterWithNameDomainIncludeUnavailable(name foundation.foundation.INSString, domain foundation.foundation.INSString, flag bool) IPrinter {
	rv := objc.Send[Printer](objc.ID(pc.class), objc.Sel("printerWithName:domain:includeUnavailable:"), name, domain, flag)
	return rv
}







// Returns the names of all available printers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/printerNames
func (pc _PrinterClass) PrinterNames() []string {
	rv := objc.Send[[]string](objc.ID(pc.class), objc.Sel("printerNames"))
	return rv
}

// Returns descriptions of the makes and models of all available printers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/printerTypes
func (pc _PrinterClass) PrinterTypes() []string {
	rv := objc.Send[[]string](objc.ID(pc.class), objc.Sel("printerTypes"))
	return rv
}






// Returns the size of the page for the specified paper type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/pageSize(forPaper:)
func (p_ Printer) PageSizeForPaper(paperName PrinterPaperName) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p_.ID, objc.Sel("pageSizeForPaper:"), paperName)
	return rv
}







// A dictionary of keys and values that describe the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/deviceDescription
func (p_ Printer) DeviceDescription() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("deviceDescription"))
	return rv
}


// The PostScript language level recognized by the printer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/languageLevel
func (p_ Printer) LanguageLevel() int {
	rv := objc.Send[int](p_.ID, objc.Sel("languageLevel"))
	return rv
}


// The printer’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/name
func (p_ Printer) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("name"))
	return rv
}


// Returns the names of all available printers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/printerNames
func (p_ Printer) PrinterNames() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("printerNames"))
	return rv
}


// Returns descriptions of the makes and models of all available printers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/printerTypes
func (p_ Printer) PrinterTypes() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("printerTypes"))
	return rv
}


// A description of the printer’s make and model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/type
func (p_ Printer) Type() PrinterTypeName {
	rv := objc.Send[PrinterTypeName](p_.ID, objc.Sel("type"))
	return rv
}







