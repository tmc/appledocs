// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPrintInfo */


/* debug [class_header]: Header for NSPrintInfo */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PrintInfo */
// An interface definition for the [PrintInfo] class.
type IPrintInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PrintInfo */
	// properties:
	BottomMargin() float64
	SetBottomMargin(value float64)
	HorizontalPagination() PrintingPaginationMode
	SetHorizontalPagination(value PrintingPaginationMode)
	ImageablePageBounds() Rect /* not a class type */
	HorizontallyCentered() bool
	SetHorizontallyCentered(value bool)
	SelectionOnly() bool
	SetSelectionOnly(value bool)
	VerticallyCentered() bool
	SetVerticallyCentered(value bool)
	JobDisposition() PrintJobDispositionValue /* typedef */
	SetJobDisposition(value PrintJobDispositionValue /* typedef */)
	LeftMargin() float64
	SetLeftMargin(value float64)
	LocalizedPaperName() objc.IObject /* cross-framework: NSString */
	Orientation() PaperOrientation
	SetOrientation(value PaperOrientation)
	PaperName() PrinterPaperName /* typedef */
	SetPaperName(value PrinterPaperName /* typedef */)
	PaperSize() Size /* not a class type */
	SetPaperSize(value Size /* not a class type */)
	PrintSettings() unsafe.Pointer
	Printer() IPrinter
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PrintInfo */
	// methods:
	Dictionary() unsafe.Pointer
	PMPageFormat()
	PMPrintSession()
	PMPrintSettings()
	SetUpPrintOperationDefaultValues()
	TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo)
	UpdateFromPMPageFormat()
	UpdateFromPMPrintSettings()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PrintInfo */
// Alloc allocates a new instance without initialization.
func (pc _PrintInfoClass) Alloc() PrintInfo {
	rv := objc.Send[PrintInfo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PrintInfo */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PrintInfo */

// Creates a printing information object from data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/init(coder:)
func NewPrintInfoWithCoder(coder foundation.Coder) PrintInfo {
	instance := getPrintInfoClass().Alloc()
	rv := objc.Send[PrintInfo](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPrintInfoWithCoder */


// Returns a printing information object initialized with the parameters in the specified dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/init(dictionary:)
func NewPrintInfoWithDictionary(attributes foundation.IDictionary) PrintInfo {
	instance := getPrintInfoClass().Alloc()
	rv := objc.Send[PrintInfo](instance.ID, objc.Sel("initWithDictionary:"), attributes)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPrintInfoWithDictionary */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PrintInfo */

// Deprecated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/setDefaultPrinter:
func (pc _PrintInfoClass) SetDefaultPrinter(printer IPrinter) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("setDefaultPrinter:"), printer)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetDefaultPrinter) */


// Deprecated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/sizeForPaperName:
func (pc _PrintInfoClass) SizeForPaperName(name PrinterPaperName /* typedef */) Size /* not a class type */ {
	rv := objc.Send[Size](objc.ID(pc.class), objc.Sel("sizeForPaperName:"), name)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SizeForPaperName) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PrintInfo */

// Deprecated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/defaultPrinter
func (pc _PrintInfoClass) DefaultPrinter() IPrinter {
	rv := objc.Send[Printer](objc.ID(pc.class), objc.Sel("defaultPrinter"))
	return rv
}/* debug [class_properties_class/property]: defaultPrinter */

// The shared printing information object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/shared
func (pc _PrintInfoClass) SharedPrintInfo() PrintInfo {
	rv := objc.Send[PrintInfo](objc.ID(pc.class), objc.Sel("sharedPrintInfo"))
	return rv
}/* debug [class_properties_class/property]: sharedPrintInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PrintInfo */

// Returns the print info’s dictionary that contains the printing attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/dictionary()
func (p_ PrintInfo) Dictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dictionary"))
	return rv
}/* debug [instance_methods/method]: Dictionary */


// Returns a Core Printing object configured with the print info’s page format information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/pmPageFormat()
func (p_ PrintInfo) PMPageFormat() {
	objc.Send[objc.ID](p_.ID, objc.Sel("PMPageFormat"))
}/* debug [instance_methods/method]: PMPageFormat */


// Returns a Core Printing object configured with the print info’s session information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/pmPrintSession()
func (p_ PrintInfo) PMPrintSession() {
	objc.Send[objc.ID](p_.ID, objc.Sel("PMPrintSession"))
}/* debug [instance_methods/method]: PMPrintSession */


// Returns a Core Printing object configured with the print info’s print settings information
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/pmPrintSettings()
func (p_ PrintInfo) PMPrintSettings() {
	objc.Send[objc.ID](p_.ID, objc.Sel("PMPrintSettings"))
}/* debug [instance_methods/method]: PMPrintSettings */


// Validates the attributes encapsulated by the print info.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/setUpPrintOperationDefaultValues()
func (p_ PrintInfo) SetUpPrintOperationDefaultValues() {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUpPrintOperationDefaultValues"))
}/* debug [instance_methods/method]: SetUpPrintOperationDefaultValues */


// Updates the print info with all the settings and attributes in the specified PDF info object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/takeSettings(from:)
func (p_ PrintInfo) TakeSettingsFromPDFInfo(inPDFInfo IPDFInfo) {
	objc.Send[objc.ID](p_.ID, objc.Sel("takeSettingsFromPDFInfo:"), inPDFInfo)
}/* debug [instance_methods/method]: TakeSettingsFromPDFInfo */


// Synchronizes the print info’s page format information with information from its associated page format object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/updateFromPMPageFormat()
func (p_ PrintInfo) UpdateFromPMPageFormat() {
	objc.Send[objc.ID](p_.ID, objc.Sel("updateFromPMPageFormat"))
}/* debug [instance_methods/method]: UpdateFromPMPageFormat */


// Synchronizes the print info’s print settings information with information from its associated print settings object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/updateFromPMPrintSettings()
func (p_ PrintInfo) UpdateFromPMPrintSettings() {
	objc.Send[objc.ID](p_.ID, objc.Sel("updateFromPMPrintSettings"))
}/* debug [instance_methods/method]: UpdateFromPMPrintSettings */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PrintInfo */

// The height of the bottom margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/bottomMargin
func (p_ PrintInfo) BottomMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("bottomMargin"))
	return rv
}/* debug [instance_properties/getter]: bottomMargin */


// The height of the bottom margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/bottomMargin
func (p_ PrintInfo) SetBottomMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBottomMargin:"), value)
}/* debug [instance_properties/setter]: bottomMargin */


// Deprecated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/defaultPrinter
func (p_ PrintInfo) DefaultPrinter() IPrinter {
	rv := objc.Send[Printer](p_.ID, objc.Sel("defaultPrinter"))
	return rv
}/* debug [instance_properties/getter]: defaultPrinter */


// The horizontal pagination mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/horizontalPagination
func (p_ PrintInfo) HorizontalPagination() PrintingPaginationMode {
	rv := objc.Send[PrintingPaginationMode](p_.ID, objc.Sel("horizontalPagination"))
	return rv
}/* debug [instance_properties/getter]: horizontalPagination */


// The horizontal pagination mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/horizontalPagination
func (p_ PrintInfo) SetHorizontalPagination(value PrintingPaginationMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHorizontalPagination:"), value)
}/* debug [instance_properties/setter]: horizontalPagination */


// The imageable area of a sheet of paper specified by the print info.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/imageablePageBounds
func (p_ PrintInfo) ImageablePageBounds() Rect /* not a class type */ {
	rv := objc.Send[Rect](p_.ID, objc.Sel("imageablePageBounds"))
	return rv
}/* debug [instance_properties/getter]: imageablePageBounds */


// A Boolean value that indicates whether the image is centered horizontally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isHorizontallyCentered
func (p_ PrintInfo) HorizontallyCentered() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("horizontallyCentered"))
	return rv
}/* debug [instance_properties/getter]: horizontallyCentered */


// A Boolean value that indicates whether the image is centered horizontally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isHorizontallyCentered
func (p_ PrintInfo) SetHorizontallyCentered(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHorizontallyCentered:"), value)
}/* debug [instance_properties/setter]: horizontallyCentered */


// A Boolean value that indicates whether only the currently selected contents should be printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isSelectionOnly
func (p_ PrintInfo) SelectionOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("selectionOnly"))
	return rv
}/* debug [instance_properties/getter]: selectionOnly */


// A Boolean value that indicates whether only the currently selected contents should be printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isSelectionOnly
func (p_ PrintInfo) SetSelectionOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectionOnly:"), value)
}/* debug [instance_properties/setter]: selectionOnly */


// A Boolean value that indicates whether the image is centered vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isVerticallyCentered
func (p_ PrintInfo) VerticallyCentered() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("verticallyCentered"))
	return rv
}/* debug [instance_properties/getter]: verticallyCentered */


// A Boolean value that indicates whether the image is centered vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isVerticallyCentered
func (p_ PrintInfo) SetVerticallyCentered(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVerticallyCentered:"), value)
}/* debug [instance_properties/setter]: verticallyCentered */


// The action specified for the job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property
func (p_ PrintInfo) JobDisposition() PrintJobDispositionValue /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("jobDisposition"))
	return rv
}/* debug [instance_properties/getter]: jobDisposition */


// The action specified for the job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property
func (p_ PrintInfo) SetJobDisposition(value PrintJobDispositionValue /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJobDisposition:"), value)
}/* debug [instance_properties/setter]: jobDisposition */


// The width of the left margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/leftMargin
func (p_ PrintInfo) LeftMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("leftMargin"))
	return rv
}/* debug [instance_properties/getter]: leftMargin */


// The width of the left margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/leftMargin
func (p_ PrintInfo) SetLeftMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLeftMargin:"), value)
}/* debug [instance_properties/setter]: leftMargin */


// The human-readable name of the currently selected paper size, suitable for presentation in user interfaces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/localizedPaperName
func (p_ PrintInfo) LocalizedPaperName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localizedPaperName"))
	return rv
}/* debug [instance_properties/getter]: localizedPaperName */


// The orientation attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/orientation-swift.property
func (p_ PrintInfo) Orientation() PaperOrientation {
	rv := objc.Send[PaperOrientation](p_.ID, objc.Sel("orientation"))
	return rv
}/* debug [instance_properties/getter]: orientation */


// The orientation attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/orientation-swift.property
func (p_ PrintInfo) SetOrientation(value PaperOrientation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOrientation:"), value)
}/* debug [instance_properties/setter]: orientation */


// The name of the currently selected paper size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/paperName
func (p_ PrintInfo) PaperName() PrinterPaperName /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("paperName"))
	return rv
}/* debug [instance_properties/getter]: paperName */


// The name of the currently selected paper size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/paperName
func (p_ PrintInfo) SetPaperName(value PrinterPaperName /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaperName:"), value)
}/* debug [instance_properties/setter]: paperName */


// The size of the paper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/paperSize
func (p_ PrintInfo) PaperSize() Size /* not a class type */ {
	rv := objc.Send[Size](p_.ID, objc.Sel("paperSize"))
	return rv
}/* debug [instance_properties/getter]: paperSize */


// The size of the paper.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/paperSize
func (p_ PrintInfo) SetPaperSize(value Size /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaperSize:"), value)
}/* debug [instance_properties/setter]: paperSize */


// A mutable dictionary containing the print settings from Core Printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/printSettings
func (p_ PrintInfo) PrintSettings() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("printSettings"))
	return rv
}/* debug [instance_properties/getter]: printSettings */


// The printer object to be used for printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/printer
func (p_ PrintInfo) Printer() IPrinter {
	rv := objc.Send[Printer](p_.ID, objc.Sel("printer"))
	return rv
}/* debug [instance_properties/getter]: printer */


// The printer object to be used for printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/printer
func (p_ PrintInfo) SetPrinter(value IPrinter) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrinter:"), value)
}/* debug [instance_properties/setter]: printer */


// The width of the right margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/rightMargin
func (p_ PrintInfo) RightMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("rightMargin"))
	return rv
}/* debug [instance_properties/getter]: rightMargin */


// The width of the right margin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/rightMargin
func (p_ PrintInfo) SetRightMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRightMargin:"), value)
}/* debug [instance_properties/setter]: rightMargin */


// The current scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/scalingFactor
func (p_ PrintInfo) ScalingFactor() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("scalingFactor"))
	return rv
}/* debug [instance_properties/getter]: scalingFactor */


// The current scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/scalingFactor
func (p_ PrintInfo) SetScalingFactor(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setScalingFactor:"), value)
}/* debug [instance_properties/setter]: scalingFactor */


// The shared printing information object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/shared
func (p_ PrintInfo) SharedPrintInfo() IPrintInfo {
	rv := objc.Send[PrintInfo](p_.ID, objc.Sel("sharedPrintInfo"))
	return rv
}/* debug [instance_properties/getter]: sharedPrintInfo */


// The shared printing information object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/shared
func (p_ PrintInfo) SetSharedPrintInfo(value IPrintInfo) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSharedPrintInfo:"), value)
}/* debug [instance_properties/setter]: sharedPrintInfo */


// The top margin to the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/topMargin
func (p_ PrintInfo) TopMargin() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("topMargin"))
	return rv
}/* debug [instance_properties/getter]: topMargin */


// The top margin to the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/topMargin
func (p_ PrintInfo) SetTopMargin(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTopMargin:"), value)
}/* debug [instance_properties/setter]: topMargin */


// The vertical pagination to the specified mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/verticalPagination
func (p_ PrintInfo) VerticalPagination() PrintingPaginationMode {
	rv := objc.Send[PrintingPaginationMode](p_.ID, objc.Sel("verticalPagination"))
	return rv
}/* debug [instance_properties/getter]: verticalPagination */


// The vertical pagination to the specified mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/verticalPagination
func (p_ PrintInfo) SetVerticalPagination(value PrintingPaginationMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setVerticalPagination:"), value)
}/* debug [instance_properties/setter]: verticalPagination */


// A Boolean value that indicates whether the image is centered horizontally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/ishorizontallycentered
func (p_ PrintInfo) IsHorizontallyCentered() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isHorizontallyCentered"))
	return rv
}/* debug [instance_properties/getter]: isHorizontallyCentered */


// A Boolean value that indicates whether the image is centered horizontally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/ishorizontallycentered
func (p_ PrintInfo) SetIsHorizontallyCentered(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsHorizontallyCentered:"), value)
}/* debug [instance_properties/setter]: isHorizontallyCentered */


// A Boolean value that indicates whether only the currently selected contents should be printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/isselectiononly
func (p_ PrintInfo) IsSelectionOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSelectionOnly"))
	return rv
}/* debug [instance_properties/getter]: isSelectionOnly */


// A Boolean value that indicates whether only the currently selected contents should be printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/isselectiononly
func (p_ PrintInfo) SetIsSelectionOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSelectionOnly:"), value)
}/* debug [instance_properties/setter]: isSelectionOnly */


// A Boolean value that indicates whether the image is centered vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/isverticallycentered
func (p_ PrintInfo) IsVerticallyCentered() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isVerticallyCentered"))
	return rv
}/* debug [instance_properties/getter]: isVerticallyCentered */


// A Boolean value that indicates whether the image is centered vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintinfo/isverticallycentered
func (p_ PrintInfo) SetIsVerticallyCentered(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsVerticallyCentered:"), value)
}/* debug [instance_properties/setter]: isVerticallyCentered */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPrintInfo */


