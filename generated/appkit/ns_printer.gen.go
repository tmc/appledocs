// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objc"
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
	BooleanForKeyInTable(key string, table string) bool
	ImageRectForPaper(paperName string) coregraphics.CGRect
	IsColor() bool
	IsKeyInTable(key string, table string) bool
	PageSizeForPaper(paperName unsafe.Pointer) coregraphics.CGSize
	StatusForTable(tableName string) unsafe.Pointer
	StringForKeyInTable(key string, table string) unsafe.Pointer
	StringListForKeyInTable(key string, table string) unsafe.Pointer
}

// An object that describes a printer’s capabilities.
//
// provides information about a printer; it does not modify printer attributes or control a printing job. A printer object can be constructed by specifying either the printer name or the make and model of an available printer. Typically, Cocoa apps don’t create objects; instead, the printing system uses these objects to support the printing jobs and when it shows users a list of printers.
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

// Returns the Boolean value associated with the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/booleanForKey:inTable:
func (p_ Printer) BooleanForKeyInTable(key string, table string) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("booleanForKey:inTable:"), objc.String(key), objc.String(table))
	return rv
}

// Deprecated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/imageRectForPaper:
func (p_ Printer) ImageRectForPaper(paperName string) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("imageRectForPaper:"), objc.String(paperName))
	return rv
}

// Deprecated.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/isColor
func (p_ Printer) IsColor() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isColor"))
	return rv
}

// Returns a Boolean value that indicates whether the specified key is in the specified table.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/isKey:inTable:
func (p_ Printer) IsKeyInTable(key string, table string) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isKey:inTable:"), objc.String(key), objc.String(table))
	return rv
}

// Returns the size of the page for the specified paper type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/pageSize(forPaper:)
func (p_ Printer) PageSizeForPaper(paperName unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](p_.ID, objc.Sel("pageSizeForPaper:"), paperName)
	return rv
}

// Returns the status of the specified table.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/statusForTable:
func (p_ Printer) StatusForTable(tableName string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("statusForTable:"), objc.String(tableName))
	return rv
}

// Returns the first occurrence of a value associated with specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/stringForKey:inTable:
func (p_ Printer) StringForKeyInTable(key string, table string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("stringForKey:inTable:"), objc.String(key), objc.String(table))
	return rv
}

// Returns an array of strings, one for each occurrence, associated with specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/stringListForKey:inTable:
func (p_ Printer) StringListForKeyInTable(key string, table string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("stringListForKey:inTable:"), objc.String(key), objc.String(table))
	return rv
}

// A dictionary of keys and values that describe the device.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/deviceDescription
func (p_ Printer) DeviceDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("deviceDescription"))
	return rv
}

// The PostScript language level recognized by the printer.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/languageLevel
func (p_ Printer) LanguageLevel() int {
	rv := objc.Send[int](p_.ID, objc.Sel("languageLevel"))
	return rv
}

// A description of the printer’s make and model.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/type
func (p_ Printer) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("type"))
	return rv
}
