// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PrintInfo] class.
var (
	PrintInfoClass     _PrintInfoClass
	PrintInfoClassOnce sync.Once
)

func getPrintInfoClass() _PrintInfoClass {
	PrintInfoClassOnce.Do(func() {
		PrintInfoClass = _PrintInfoClass{objc.GetClass("NSPrintInfo")}
	})
	return PrintInfoClass
}

type _PrintInfoClass struct {
	class objc.Class
}

// An interface definition for the [PrintInfo] class.
type IPrintInfo interface {
	objectivec.IObject
	Dictionary() unsafe.Pointer
	HorizontallyCentered() bool
	SetHorizontallyCentered(value bool)
	SelectionOnly() bool
	SetSelectionOnly(value bool)
	JobDisposition() unsafe.Pointer
	SetJobDisposition(value unsafe.Pointer)
	BottomMargin() float64
	SetBottomMargin(value float64)
	HorizontalPagination() unsafe.Pointer
	SetHorizontalPagination(value unsafe.Pointer)
	ImageablePageBounds() coregraphics.CGRect
	SetImageablePageBounds(value coregraphics.CGRect)
	IsHorizontallyCentered() bool
	SetIsHorizontallyCentered(value bool)
	IsSelectionOnly() bool
	SetIsSelectionOnly(value bool)
	IsVerticallyCentered() bool
	SetIsVerticallyCentered(value bool)
	LeftMargin() float64
	SetLeftMargin(value float64)
	LocalizedPaperName() string
	SetLocalizedPaperName(value string)
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
	PaperName() unsafe.Pointer
	SetPaperName(value unsafe.Pointer)
	PaperSize() coregraphics.CGSize
	SetPaperSize(value coregraphics.CGSize)
	PrintSettings() foundation.MutableDictionary
	SetPrintSettings(value foundation.IMutableDictionary)
	Printer() NSPrinter
	SetPrinter(value IPrinter)
	RightMargin() float64
	SetRightMargin(value float64)
	ScalingFactor() float64
	SetScalingFactor(value float64)
	TopMargin() float64
	SetTopMargin(value float64)
	VerticalPagination() unsafe.Pointer
	SetVerticalPagination(value unsafe.Pointer)
}

// An object that stores information that’s used to generate printed output.
//
// A shared object is automatically created for an app and is used by default for all printing jobs for that app. The printing information in an object is stored in a dictionary. To access the standard attributes in the dictionary directly, this class defines a set of keys and provides the method. You can also initialize an instance of this class using the method. You can use this dictionary to store custom information associated with a print job. Any non-object values should be stored as or objects in the dictionary. See for a list of types which should be stored as numbers. For other non-object values, use the class. To store custom information that belongs in printing presets you should use the dictionary returned by the method.


// An object that stores information that’s used to generate printed output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo

type PrintInfo struct {
	objectivec.Object
}

// PrintInfoFrom constructs a [PrintInfo] from an unsafe.Pointer.
//
// An object that stores information that’s used to generate printed output.
func PrintInfoFrom(ptr unsafe.Pointer) PrintInfo {
	return PrintInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PrintInfoClass) Alloc() PrintInfo {
	rv := objc.Send[PrintInfo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PrintInfoClass) New() PrintInfo {
	rv := objc.Send[PrintInfo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PrintInfo) Init() PrintInfo {
	rv := objc.Send[PrintInfo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PrintInfo) Autorelease() PrintInfo {
	rv := objc.Send[PrintInfo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPrintInfo creates a new PrintInfo instance.
func NewPrintInfo() PrintInfo {
	return getPrintInfoClass().New()
}




// Returns the print info’s dictionary that contains the printing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/dictionary()

func (p_ PrintInfo) Dictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dictionary"))
	return rv
}


// A Boolean value that indicates whether the image is centered horizontally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isHorizontallyCentered

func (p_ PrintInfo) HorizontallyCentered() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("horizontallyCentered"))
	return rv
}


// A Boolean value that indicates whether the image is centered horizontally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isHorizontallyCentered

func (p_ PrintInfo) SetHorizontallyCentered(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHorizontallyCentered:"), value)
}


// A Boolean value that indicates whether only the currently selected contents should be printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isSelectionOnly

func (p_ PrintInfo) SelectionOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("selectionOnly"))
	return rv
}


// A Boolean value that indicates whether only the currently selected contents should be printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isSelectionOnly

func (p_ PrintInfo) SetSelectionOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectionOnly:"), value)
}


// The action specified for the job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property

func (p_ PrintInfo) JobDisposition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("jobDisposition"))
	return rv
}


// The action specified for the job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property

func (p_ PrintInfo) SetJobDisposition(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJobDisposition:"), value)
}


// The height of the bottom margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/bottommargin

func (p_ PrintInfo) BottomMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("bottomMargin"))
	return rv
}


// The height of the bottom margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/bottommargin

func (p_ PrintInfo) SetBottomMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBottomMargin:"), value)
}


// The horizontal pagination mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/horizontalpagination

func (p_ PrintInfo) HorizontalPagination() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("horizontalPagination"))
	return rv
}


// The horizontal pagination mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/horizontalpagination

func (p_ PrintInfo) SetHorizontalPagination(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHorizontalPagination:"), value)
}


// The imageable area of a sheet of paper specified by the print info.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/imageablepagebounds

func (p_ PrintInfo) ImageablePageBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("imageablePageBounds"))
	return rv
}


// The imageable area of a sheet of paper specified by the print info.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/imageablepagebounds

func (p_ PrintInfo) SetImageablePageBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setImageablePageBounds:"), value)
}


// A Boolean value that indicates whether the image is centered horizontally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/ishorizontallycentered

func (p_ PrintInfo) IsHorizontallyCentered() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isHorizontallyCentered"))
	return rv
}


// A Boolean value that indicates whether the image is centered horizontally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/ishorizontallycentered

func (p_ PrintInfo) SetIsHorizontallyCentered(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsHorizontallyCentered:"), value)
}


// A Boolean value that indicates whether only the currently selected contents should be printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/isselectiononly

func (p_ PrintInfo) IsSelectionOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSelectionOnly"))
	return rv
}


// A Boolean value that indicates whether only the currently selected contents should be printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/isselectiononly

func (p_ PrintInfo) SetIsSelectionOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSelectionOnly:"), value)
}


// A Boolean value that indicates whether the image is centered vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/isverticallycentered

func (p_ PrintInfo) IsVerticallyCentered() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isVerticallyCentered"))
	return rv
}


// A Boolean value that indicates whether the image is centered vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/isverticallycentered

func (p_ PrintInfo) SetIsVerticallyCentered(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsVerticallyCentered:"), value)
}


// The width of the left margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/leftmargin

func (p_ PrintInfo) LeftMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("leftMargin"))
	return rv
}


// The width of the left margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/leftmargin

func (p_ PrintInfo) SetLeftMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLeftMargin:"), value)
}


// The human-readable name of the currently selected paper size, suitable for presentation in user interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/localizedpapername

func (p_ PrintInfo) LocalizedPaperName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedPaperName"))
	return rv
}


// The human-readable name of the currently selected paper size, suitable for presentation in user interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/localizedpapername

func (p_ PrintInfo) SetLocalizedPaperName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedPaperName:"), objc.String(value))
}


// The orientation attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/orientation-swift.property

func (p_ PrintInfo) Orientation() PaperOrientation {
	rv := objc.Send[PaperOrientation](p_.ID, objc.Sel("orientation"))
	return rv
}


// The orientation attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/orientation-swift.property

func (p_ PrintInfo) SetOrientation(value PaperOrientation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOrientation:"), value)
}


// The name of the currently selected paper size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/papername

func (p_ PrintInfo) PaperName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("paperName"))
	return rv
}


// The name of the currently selected paper size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/papername

func (p_ PrintInfo) SetPaperName(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaperName:"), value)
}


// The size of the paper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/papersize

func (p_ PrintInfo) PaperSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](p_.ID, objc.Sel("paperSize"))
	return rv
}


// The size of the paper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/papersize

func (p_ PrintInfo) SetPaperSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaperSize:"), value)
}


// A mutable dictionary containing the print settings from Core Printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/printsettings

func (p_ PrintInfo) PrintSettings() foundation.MutableDictionary {
	rv := objc.Send[foundation.MutableDictionary](p_.ID, objc.Sel("printSettings"))
	return rv
}


// A mutable dictionary containing the print settings from Core Printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/printsettings

func (p_ PrintInfo) SetPrintSettings(value foundation.IMutableDictionary) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrintSettings:"), value)
}


// The printer object to be used for printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/printer

func (p_ PrintInfo) Printer() NSPrinter {
	rv := objc.Send[NSPrinter](p_.ID, objc.Sel("printer"))
	return rv
}


// The printer object to be used for printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/printer

func (p_ PrintInfo) SetPrinter(value IPrinter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrinter:"), value)
}


// The width of the right margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/rightmargin

func (p_ PrintInfo) RightMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("rightMargin"))
	return rv
}


// The width of the right margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/rightmargin

func (p_ PrintInfo) SetRightMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRightMargin:"), value)
}


// The current scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/scalingfactor

func (p_ PrintInfo) ScalingFactor() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("scalingFactor"))
	return rv
}


// The current scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/scalingfactor

func (p_ PrintInfo) SetScalingFactor(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setScalingFactor:"), value)
}


// The top margin to the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/topmargin

func (p_ PrintInfo) TopMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("topMargin"))
	return rv
}


// The top margin to the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/topmargin

func (p_ PrintInfo) SetTopMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTopMargin:"), value)
}


// The vertical pagination to the specified mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/verticalpagination

func (p_ PrintInfo) VerticalPagination() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("verticalPagination"))
	return rv
}


// The vertical pagination to the specified mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/verticalpagination

func (p_ PrintInfo) SetVerticalPagination(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVerticalPagination:"), value)
}



