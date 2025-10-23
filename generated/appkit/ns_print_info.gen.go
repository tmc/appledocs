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
	PMPageFormat()
	PMPrintSession()
	PMPrintSettings()
	SetUpPrintOperationDefaultValues()
	TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo)
	UpdateFromPMPageFormat()
	UpdateFromPMPrintSettings()
	BottomMargin() float64
	SetBottomMargin(value float64)
	HorizontalPagination() PrintingPaginationMode
	SetHorizontalPagination(value PrintingPaginationMode)
	ImageablePageBounds() coregraphics.CGRect
	HorizontallyCentered() bool
	SetHorizontallyCentered(value bool)
	SelectionOnly() bool
	SetSelectionOnly(value bool)
	VerticallyCentered() bool
	SetVerticallyCentered(value bool)
	JobDisposition() PrintJobDispositionValue
	SetJobDisposition(value IPrintJobDispositionValue)
	LeftMargin() float64
	SetLeftMargin(value float64)
	LocalizedPaperName() string
	Orientation() PaperOrientation
	SetOrientation(value IPaperOrientation)
	PaperName() PrinterPaperName
	SetPaperName(value IPrinterPaperName)
	PaperSize() coregraphics.CGSize
	SetPaperSize(value coregraphics.CGSize)
	PrintSettings() unsafe.Pointer
	Printer() NSPrinter
	SetPrinter(value IPrinter)
	RightMargin() float64
	SetRightMargin(value float64)
	ScalingFactor() float64
	SetScalingFactor(value float64)
	TopMargin() float64
	SetTopMargin(value float64)
	VerticalPagination() PrintingPaginationMode
	SetVerticalPagination(value PrintingPaginationMode)
	IsHorizontallyCentered() bool
	SetIsHorizontallyCentered(value bool)
	IsSelectionOnly() bool
	SetIsSelectionOnly(value bool)
	IsVerticallyCentered() bool
	SetIsVerticallyCentered(value bool)
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



// Creates a printing information object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/init(coder:)
func NewPrintInfoWithCoder(coder foundation.ICoder) PrintInfo {
	instance := getPrintInfoClass().Alloc()
	rv := objc.Send[PrintInfo](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Returns a printing information object initialized with the parameters in the specified dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/init(dictionary:)
func NewPrintInfoWithDictionary(attributes unsafe.Pointer) PrintInfo {
	instance := getPrintInfoClass().Alloc()
	rv := objc.Send[PrintInfo](instance.ID, objc.Sel("initWithDictionary:"), attributes)
	rv.Autorelease()
	return rv
}



// Deprecated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/setDefaultPrinter:
func (pc _PrintInfoClass) SetDefaultPrinter(printer IPrinter) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("setDefaultPrinter:"), printer)
}


// Deprecated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/sizeForPaperName:
func (pc _PrintInfoClass) SizeForPaperName(name IPrinterPaperName) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](objc.ID(pc.class), objc.Sel("sizeForPaperName:"), name)
	return rv
}


// Deprecated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/defaultPrinter
func (pc _PrintInfoClass) DefaultPrinter() NSPrinter {
	rv := objc.Send[NSPrinter](objc.ID(pc.class), objc.Sel("defaultPrinter"))
	return rv
}

// The shared printing information object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/shared
func (pc _PrintInfoClass) SharedPrintInfo() PrintInfo {
	rv := objc.Send[NSPrintInfo](objc.ID(pc.class), objc.Sel("sharedPrintInfo"))
	return rv
}

// Returns the print info’s dictionary that contains the printing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/dictionary()
func (p_ PrintInfo) Dictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dictionary"))
	return rv
}


// Returns a Core Printing object configured with the print info’s page format information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/pmPageFormat()
func (p_ PrintInfo) PMPageFormat() {
	objc.Send[objc.ID](p_.ID, objc.Sel("PMPageFormat"))
}


// Returns a Core Printing object configured with the print info’s session information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/pmPrintSession()
func (p_ PrintInfo) PMPrintSession() {
	objc.Send[objc.ID](p_.ID, objc.Sel("PMPrintSession"))
}


// Returns a Core Printing object configured with the print info’s print settings information
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/pmPrintSettings()
func (p_ PrintInfo) PMPrintSettings() {
	objc.Send[objc.ID](p_.ID, objc.Sel("PMPrintSettings"))
}


// Validates the attributes encapsulated by the print info.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/setUpPrintOperationDefaultValues()
func (p_ PrintInfo) SetUpPrintOperationDefaultValues() {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUpPrintOperationDefaultValues"))
}


// Updates the print info with all the settings and attributes in the specified PDF info object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/takeSettings(from:)
func (p_ PrintInfo) TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo) {
	objc.Send[objc.ID](p_.ID, objc.Sel("takeSettingsFromPDFInfo:"), inPDFInfo)
}


// Synchronizes the print info’s page format information with information from its associated page format object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/updateFromPMPageFormat()
func (p_ PrintInfo) UpdateFromPMPageFormat() {
	objc.Send[objc.ID](p_.ID, objc.Sel("updateFromPMPageFormat"))
}


// Synchronizes the print info’s print settings information with information from its associated print settings object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/updateFromPMPrintSettings()
func (p_ PrintInfo) UpdateFromPMPrintSettings() {
	objc.Send[objc.ID](p_.ID, objc.Sel("updateFromPMPrintSettings"))
}


// The height of the bottom margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/bottomMargin
func (p_ PrintInfo) BottomMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("bottomMargin"))
	return rv
}


// The height of the bottom margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/bottomMargin
func (p_ PrintInfo) SetBottomMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBottomMargin:"), value)
}


// Deprecated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/defaultPrinter
func (p_ PrintInfo) DefaultPrinter() NSPrinter {
	rv := objc.Send[NSPrinter](p_.ID, objc.Sel("defaultPrinter"))
	return rv
}


// The horizontal pagination mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/horizontalPagination
func (p_ PrintInfo) HorizontalPagination() PrintingPaginationMode {
	rv := objc.Send[PrintingPaginationMode](p_.ID, objc.Sel("horizontalPagination"))
	return rv
}


// The horizontal pagination mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/horizontalPagination
func (p_ PrintInfo) SetHorizontalPagination(value PrintingPaginationMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHorizontalPagination:"), value)
}


// The imageable area of a sheet of paper specified by the print info.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/imageablePageBounds
func (p_ PrintInfo) ImageablePageBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("imageablePageBounds"))
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


// A Boolean value that indicates whether the image is centered vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isVerticallyCentered
func (p_ PrintInfo) VerticallyCentered() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("verticallyCentered"))
	return rv
}


// A Boolean value that indicates whether the image is centered vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isVerticallyCentered
func (p_ PrintInfo) SetVerticallyCentered(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVerticallyCentered:"), value)
}


// The action specified for the job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property
func (p_ PrintInfo) JobDisposition() PrintJobDispositionValue {
	rv := objc.Send[PrintJobDispositionValue](p_.ID, objc.Sel("jobDisposition"))
	return rv
}


// The action specified for the job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property
func (p_ PrintInfo) SetJobDisposition(value IPrintJobDispositionValue) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJobDisposition:"), value)
}


// The width of the left margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/leftMargin
func (p_ PrintInfo) LeftMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("leftMargin"))
	return rv
}


// The width of the left margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/leftMargin
func (p_ PrintInfo) SetLeftMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLeftMargin:"), value)
}


// The human-readable name of the currently selected paper size, suitable for presentation in user interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/localizedPaperName
func (p_ PrintInfo) LocalizedPaperName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedPaperName"))
	return rv
}


// The orientation attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/orientation-swift.property
func (p_ PrintInfo) Orientation() PaperOrientation {
	rv := objc.Send[PaperOrientation](p_.ID, objc.Sel("orientation"))
	return rv
}


// The orientation attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/orientation-swift.property
func (p_ PrintInfo) SetOrientation(value IPaperOrientation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOrientation:"), value)
}


// The name of the currently selected paper size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/paperName
func (p_ PrintInfo) PaperName() PrinterPaperName {
	rv := objc.Send[PrinterPaperName](p_.ID, objc.Sel("paperName"))
	return rv
}


// The name of the currently selected paper size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/paperName
func (p_ PrintInfo) SetPaperName(value IPrinterPaperName) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaperName:"), value)
}


// The size of the paper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/paperSize
func (p_ PrintInfo) PaperSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](p_.ID, objc.Sel("paperSize"))
	return rv
}


// The size of the paper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/paperSize
func (p_ PrintInfo) SetPaperSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaperSize:"), value)
}


// A mutable dictionary containing the print settings from Core Printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/printSettings
func (p_ PrintInfo) PrintSettings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("printSettings"))
	return rv
}


// The printer object to be used for printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/printer
func (p_ PrintInfo) Printer() NSPrinter {
	rv := objc.Send[NSPrinter](p_.ID, objc.Sel("printer"))
	return rv
}


// The printer object to be used for printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/printer
func (p_ PrintInfo) SetPrinter(value IPrinter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrinter:"), value)
}


// The width of the right margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/rightMargin
func (p_ PrintInfo) RightMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("rightMargin"))
	return rv
}


// The width of the right margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/rightMargin
func (p_ PrintInfo) SetRightMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRightMargin:"), value)
}


// The current scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/scalingFactor
func (p_ PrintInfo) ScalingFactor() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("scalingFactor"))
	return rv
}


// The current scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/scalingFactor
func (p_ PrintInfo) SetScalingFactor(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setScalingFactor:"), value)
}


// The shared printing information object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/shared
func (p_ PrintInfo) SharedPrintInfo() NSPrintInfo {
	rv := objc.Send[NSPrintInfo](p_.ID, objc.Sel("sharedPrintInfo"))
	return rv
}


// The shared printing information object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/shared
func (p_ PrintInfo) SetSharedPrintInfo(value IPrintInfo) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSharedPrintInfo:"), value)
}


// The top margin to the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/topMargin
func (p_ PrintInfo) TopMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("topMargin"))
	return rv
}


// The top margin to the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/topMargin
func (p_ PrintInfo) SetTopMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTopMargin:"), value)
}


// The vertical pagination to the specified mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/verticalPagination
func (p_ PrintInfo) VerticalPagination() PrintingPaginationMode {
	rv := objc.Send[PrintingPaginationMode](p_.ID, objc.Sel("verticalPagination"))
	return rv
}


// The vertical pagination to the specified mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/verticalPagination
func (p_ PrintInfo) SetVerticalPagination(value PrintingPaginationMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVerticalPagination:"), value)
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


