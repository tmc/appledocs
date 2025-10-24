// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPrinter */


/* debug [class_header]: Header for NSPrinter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Printer */
// An interface definition for the [Printer] class.
type IPrinter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Printer */
	// properties:
	DeviceDescription() foundation.IDictionary
	LanguageLevel() int
	Name() objc.IObject /* cross-framework: NSString */
	Type() PrinterTypeName /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Printer */
	// methods:
	PageSizeForPaper(paperName PrinterPaperName /* typedef */) Size /* not a class type */
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Printer */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Printer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Printer */

// Creates and returns a printer object initialized with the specified printer name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/init(name:)
func NewPrinterWithName(name objc.IObject /* cross-framework: NSString */) Printer {
	rv := objc.Send[Printer](objc.ID(getPrinterClass().class), objc.Sel("printerWithName:"), name)
	return rv
}/* debug [class_init_methods/constructor]: NewPrinterWithName */


// Creates and returns a printer object initialized to the first available printer with the specified make and model information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/init(type:)
func NewPrinterWithType(type_ PrinterTypeName /* typedef */) Printer {
	rv := objc.Send[Printer](objc.ID(getPrinterClass().class), objc.Sel("printerWithType:"), type_)
	return rv
}/* debug [class_init_methods/constructor]: NewPrinterWithType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Printer */

// Creates and returns a printer object initialized with the specified printer name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/init(name:)
func (pc _PrinterClass) PrinterWithName(name objc.IObject /* cross-framework: NSString */) IPrinter {
	rv := objc.Send[Printer](objc.ID(pc.class), objc.Sel("printerWithName:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrinterWithName) */


// Creates and returns a printer object initialized to the first available printer with the specified make and model information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/init(type:)
func (pc _PrinterClass) PrinterWithType(type_ PrinterTypeName /* typedef */) IPrinter {
	rv := objc.Send[Printer](objc.ID(pc.class), objc.Sel("printerWithType:"), type_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrinterWithType) */


// Deprecated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/printerWithName:domain:includeUnavailable:
func (pc _PrinterClass) PrinterWithNameDomainIncludeUnavailable(name objc.IObject /* cross-framework: NSString */, domain objc.IObject /* cross-framework: NSString */, flag bool) IPrinter {
	rv := objc.Send[Printer](objc.ID(pc.class), objc.Sel("printerWithName:domain:includeUnavailable:"), name, domain, flag)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrinterWithNameDomainIncludeUnavailable) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Printer */

// Returns the names of all available printers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/printerNames
func (pc _PrinterClass) PrinterNames() []string {
	rv := objc.Send[[]string](objc.ID(pc.class), objc.Sel("printerNames"))
	return rv
}/* debug [class_properties_class/property]: printerNames */

// Returns descriptions of the makes and models of all available printers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/printerTypes
func (pc _PrinterClass) PrinterTypes() []string {
	rv := objc.Send[[]string](objc.ID(pc.class), objc.Sel("printerTypes"))
	return rv
}/* debug [class_properties_class/property]: printerTypes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Printer */

// Returns the size of the page for the specified paper type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/pageSize(forPaper:)
func (p_ Printer) PageSizeForPaper(paperName PrinterPaperName /* typedef */) Size /* not a class type */ {
	rv := objc.Send[Size](p_.ID, objc.Sel("pageSizeForPaper:"), paperName)
	return rv
}/* debug [instance_methods/method]: PageSizeForPaper */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Printer */

// A dictionary of keys and values that describe the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/deviceDescription
func (p_ Printer) DeviceDescription() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](p_.ID, objc.Sel("deviceDescription"))
	return rv
}/* debug [instance_properties/getter]: deviceDescription */


// The PostScript language level recognized by the printer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/languageLevel
func (p_ Printer) LanguageLevel() int {
	rv := objc.Send[int](p_.ID, objc.Sel("languageLevel"))
	return rv
}/* debug [instance_properties/getter]: languageLevel */


// The printer’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/name
func (p_ Printer) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// Returns the names of all available printers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/printerNames
func (p_ Printer) PrinterNames() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("printerNames"))
	return rv
}/* debug [instance_properties/getter]: printerNames */


// Returns descriptions of the makes and models of all available printers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/printerTypes
func (p_ Printer) PrinterTypes() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("printerTypes"))
	return rv
}/* debug [instance_properties/getter]: printerTypes */


// A description of the printer’s make and model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/type
func (p_ Printer) Type() PrinterTypeName /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPrinter */


