// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPrintOperation */


/* debug [class_header]: Header for NSPrintOperation */
// The class instance for the [PrintOperation] class.
var (
	PrintOperationClass     _PrintOperationClass
	PrintOperationClassOnce sync.Once
)

func getPrintOperationClass() _PrintOperationClass {
	PrintOperationClassOnce.Do(func() {
		PrintOperationClass = _PrintOperationClass{objc.GetClass("NSPrintOperation")}
	})
	return PrintOperationClass
}

type _PrintOperationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PrintOperation */
// An interface definition for the [PrintOperation] class.
type IPrintOperation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PrintOperation */
	// properties:
	CanSpawnSeparateThread() bool
	SetCanSpawnSeparateThread(value bool)
	Context() IGraphicsContext
	CurrentPage() int
	CopyingOperation() bool
	JobTitle() objc.IObject /* cross-framework: NSString */
	SetJobTitle(value objc.IObject /* cross-framework: NSString */)
	PageOrder() PrintingPageOrder
	SetPageOrder(value PrintingPageOrder)
	PageRange() corefoundation.Range
	PDFPanel() IPDFPanel
	SetPDFPanel(value IPDFPanel)
	PreferredRenderingQuality() PrintRenderingQuality
	PrintInfo() IPrintInfo
	SetPrintInfo(value IPrintInfo)
	PrintPanel() IPrintPanel
	SetPrintPanel(value IPrintPanel)
	ShowsPrintPanel() bool
	SetShowsPrintPanel(value bool)
	ShowsProgressPanel() bool
	SetShowsProgressPanel(value bool)
	View() IView
	IsCopyingOperation() bool
	SetIsCopyingOperation(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PrintOperation */
	// methods:
	CleanUpOperation()
	CreateContext() IGraphicsContext
	DeliverResult() bool
	DestroyContext()
	RunOperation() bool
	RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow IWindow, delegate objc.IObject, didRunSelector objc.SEL, contextInfo objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PrintOperation */
// Alloc allocates a new instance without initialization.
func (pc _PrintOperationClass) Alloc() PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PrintOperation */
// An object that controls operations that generate Encapsulated PostScript (EPS) code, Portable Document Format (PDF) code, or print jobs.
//
// An object works in conjunction with two other objects: an object, which specifies how the code should be generated, and an object, which generates the actual code. It is important to note that the majority of methods in copy the instance of passed into them. Future changes to that print info are not reflected in the print info retained by the current object. All changes should be made to the print info before passing to the methods of this class. The only method in which does not copy the instance is .


// An object that controls operations that generate Encapsulated PostScript (EPS) code, Portable Document Format (PDF) code, or print jobs.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PrintOperation */

// Creates and returns an print operation object ready to control the printing of the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/init(view:)
func NewPrintOperationWithView(view IView) PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(getPrintOperationClass().class), objc.Sel("printOperationWithView:"), view)
	return rv
}/* debug [class_init_methods/constructor]: NewPrintOperationWithView */


// Creates and returns an print operation object ready to control the printing of the specified view using custom print settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/init(view:printInfo:)
func NewPrintOperationWithViewPrintInfo(view IView, printInfo IPrintInfo) PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(getPrintOperationClass().class), objc.Sel("printOperationWithView:printInfo:"), view, printInfo)
	return rv
}/* debug [class_init_methods/constructor]: NewPrintOperationWithViewPrintInfo */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PrintOperation */

// Creates and returns a new print operation object ready to control the copying of EPS graphics from the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/epsOperation(with:inside:to:)
func (pc _PrintOperationClass) EPSOperationWithViewInsideRectToData(view IView, rect Rect /* not a class type */, data foundation.MutableData) IPrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("EPSOperationWithView:insideRect:toData:"), view, rect, data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EPSOperationWithViewInsideRectToData) */


// Creates and returns a new print operation object ready to control the copying of EPS graphics from the specified view using the specified print settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/epsOperation(with:inside:to:printInfo:)
func (pc _PrintOperationClass) EPSOperationWithViewInsideRectToDataPrintInfo(view IView, rect Rect /* not a class type */, data foundation.MutableData, printInfo IPrintInfo) IPrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("EPSOperationWithView:insideRect:toData:printInfo:"), view, rect, data, printInfo)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EPSOperationWithViewInsideRectToDataPrintInfo) */


// Creates and returns a new print operation object ready to control the copying of EPS graphics from the specified view and write the resulting data to the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/epsOperation(with:inside:toPath:printInfo:)
func (pc _PrintOperationClass) EPSOperationWithViewInsideRectToPathPrintInfo(view IView, rect Rect /* not a class type */, path objc.IObject /* cross-framework: NSString */, printInfo IPrintInfo) IPrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("EPSOperationWithView:insideRect:toPath:printInfo:"), view, rect, path, printInfo)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EPSOperationWithViewInsideRectToPathPrintInfo) */


// Creates and returns an print operation object ready to control the printing of the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/init(view:)
func (pc _PrintOperationClass) PrintOperationWithView(view IView) IPrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("printOperationWithView:"), view)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrintOperationWithView) */


// Creates and returns an print operation object ready to control the printing of the specified view using custom print settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/init(view:printInfo:)
func (pc _PrintOperationClass) PrintOperationWithViewPrintInfo(view IView, printInfo IPrintInfo) IPrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("printOperationWithView:printInfo:"), view, printInfo)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrintOperationWithViewPrintInfo) */


// Creates and returns a new print operation object ready to control the copying of PDF graphics from the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfOperation(with:inside:to:)
func (pc _PrintOperationClass) PDFOperationWithViewInsideRectToData(view IView, rect Rect /* not a class type */, data foundation.MutableData) IPrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("PDFOperationWithView:insideRect:toData:"), view, rect, data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PDFOperationWithViewInsideRectToData) */


// Creates and returns a new print operation object ready to control the copying of PDF graphics from the specified view using the specified print settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfOperation(with:inside:to:printInfo:)
func (pc _PrintOperationClass) PDFOperationWithViewInsideRectToDataPrintInfo(view IView, rect Rect /* not a class type */, data foundation.MutableData, printInfo IPrintInfo) IPrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("PDFOperationWithView:insideRect:toData:printInfo:"), view, rect, data, printInfo)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PDFOperationWithViewInsideRectToDataPrintInfo) */


// Creates and returns a new print operation object ready to control the copying of PDF graphics from the specified view and write the resulting data to the specified file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfOperation(with:inside:toPath:printInfo:)
func (pc _PrintOperationClass) PDFOperationWithViewInsideRectToPathPrintInfo(view IView, rect Rect /* not a class type */, path objc.IObject /* cross-framework: NSString */, printInfo IPrintInfo) IPrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("PDFOperationWithView:insideRect:toPath:printInfo:"), view, rect, path, printInfo)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PDFOperationWithViewInsideRectToPathPrintInfo) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PrintOperation */

// The current print operation for this thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/current
func (pc _PrintOperationClass) CurrentOperation() PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("currentOperation"))
	return rv
}/* debug [class_properties_class/property]: currentOperation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PrintOperation */

// Called at the end of a print operation to remove the print operation as the current operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/cleanUp()
func (p_ PrintOperation) CleanUpOperation() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cleanUpOperation"))
}/* debug [instance_methods/method]: CleanUpOperation */


// Creates the graphics context object used for drawing during the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/createContext()
func (p_ PrintOperation) CreateContext() IGraphicsContext {
	rv := objc.Send[GraphicsContext](p_.ID, objc.Sel("createContext"))
	return rv
}/* debug [instance_methods/method]: CreateContext */


// Delivers the results of the print operation to the intended destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/deliverResult()
func (p_ PrintOperation) DeliverResult() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("deliverResult"))
	return rv
}/* debug [instance_methods/method]: DeliverResult */


// Destroys the print operation’s graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/destroyContext()
func (p_ PrintOperation) DestroyContext() {
	objc.Send[objc.ID](p_.ID, objc.Sel("destroyContext"))
}/* debug [instance_methods/method]: DestroyContext */


// Runs the print operation on the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/run()
func (p_ PrintOperation) RunOperation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("runOperation"))
	return rv
}/* debug [instance_methods/method]: RunOperation */


// Runs the print operation, calling your custom delegate method upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/runModal(for:delegate:didRun:contextInfo:)
func (p_ PrintOperation) RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow IWindow, delegate objc.IObject, didRunSelector objc.SEL, contextInfo objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("runOperationModalForWindow:delegate:didRunSelector:contextInfo:"), docWindow, delegate, didRunSelector, contextInfo)
}/* debug [instance_methods/method]: RunOperationModalForWindowDelegateDidRunSelectorContextInfo */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PrintOperation */

// A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/canSpawnSeparateThread
func (p_ PrintOperation) CanSpawnSeparateThread() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canSpawnSeparateThread"))
	return rv
}/* debug [instance_properties/getter]: canSpawnSeparateThread */


// A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/canSpawnSeparateThread
func (p_ PrintOperation) SetCanSpawnSeparateThread(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanSpawnSeparateThread:"), value)
}/* debug [instance_properties/setter]: canSpawnSeparateThread */


// The graphics context object used for generating output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/context
func (p_ PrintOperation) Context() IGraphicsContext {
	rv := objc.Send[GraphicsContext](p_.ID, objc.Sel("context"))
	return rv
}/* debug [instance_properties/getter]: context */


// The current print operation for this thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/current
func (p_ PrintOperation) CurrentOperation() IPrintOperation {
	rv := objc.Send[PrintOperation](p_.ID, objc.Sel("currentOperation"))
	return rv
}/* debug [instance_properties/getter]: currentOperation */


// The current print operation for this thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/current
func (p_ PrintOperation) SetCurrentOperation(value IPrintOperation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentOperation:"), value)
}/* debug [instance_properties/setter]: currentOperation */


// The current page number being printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/currentPage
func (p_ PrintOperation) CurrentPage() int {
	rv := objc.Send[int](p_.ID, objc.Sel("currentPage"))
	return rv
}/* debug [instance_properties/getter]: currentPage */


// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/isCopyingOperation
func (p_ PrintOperation) CopyingOperation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("copyingOperation"))
	return rv
}/* debug [instance_properties/getter]: copyingOperation */


// The custom title of the print job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/jobTitle
func (p_ PrintOperation) JobTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("jobTitle"))
	return rv
}/* debug [instance_properties/getter]: jobTitle */


// The custom title of the print job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/jobTitle
func (p_ PrintOperation) SetJobTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJobTitle:"), value)
}/* debug [instance_properties/setter]: jobTitle */


// The print order for the pages of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pageOrder-swift.property
func (p_ PrintOperation) PageOrder() PrintingPageOrder {
	rv := objc.Send[PrintingPageOrder](p_.ID, objc.Sel("pageOrder"))
	return rv
}/* debug [instance_properties/getter]: pageOrder */


// The print order for the pages of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pageOrder-swift.property
func (p_ PrintOperation) SetPageOrder(value PrintingPageOrder) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageOrder:"), value)
}/* debug [instance_properties/setter]: pageOrder */


// The range of pages associated with the print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pageRange
func (p_ PrintOperation) PageRange() corefoundation.Range {
	rv := objc.Send[corefoundation.Range](p_.ID, objc.Sel("pageRange"))
	return rv
}/* debug [instance_properties/getter]: pageRange */


// The PDF panel object to use during the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfPanel
func (p_ PrintOperation) PDFPanel() IPDFPanel {
	rv := objc.Send[PDFPanel](p_.ID, objc.Sel("PDFPanel"))
	return rv
}/* debug [instance_properties/getter]: PDFPanel */


// The PDF panel object to use during the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfPanel
func (p_ PrintOperation) SetPDFPanel(value IPDFPanel) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPDFPanel:"), value)
}/* debug [instance_properties/setter]: PDFPanel */


// The printing quality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/preferredRenderingQuality
func (p_ PrintOperation) PreferredRenderingQuality() PrintRenderingQuality {
	rv := objc.Send[PrintRenderingQuality](p_.ID, objc.Sel("preferredRenderingQuality"))
	return rv
}/* debug [instance_properties/getter]: preferredRenderingQuality */


// The printing information associated with the print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printInfo
func (p_ PrintOperation) PrintInfo() IPrintInfo {
	rv := objc.Send[PrintInfo](p_.ID, objc.Sel("printInfo"))
	return rv
}/* debug [instance_properties/getter]: printInfo */


// The printing information associated with the print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printInfo
func (p_ PrintOperation) SetPrintInfo(value IPrintInfo) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrintInfo:"), value)
}/* debug [instance_properties/setter]: printInfo */


// The print panel object to use during the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printPanel
func (p_ PrintOperation) PrintPanel() IPrintPanel {
	rv := objc.Send[PrintPanel](p_.ID, objc.Sel("printPanel"))
	return rv
}/* debug [instance_properties/getter]: printPanel */


// The print panel object to use during the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printPanel
func (p_ PrintOperation) SetPrintPanel(value IPrintPanel) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrintPanel:"), value)
}/* debug [instance_properties/setter]: printPanel */


// A Boolean value that determines whether the print operation displays a print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/showsPrintPanel
func (p_ PrintOperation) ShowsPrintPanel() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsPrintPanel"))
	return rv
}/* debug [instance_properties/getter]: showsPrintPanel */


// A Boolean value that determines whether the print operation displays a print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/showsPrintPanel
func (p_ PrintOperation) SetShowsPrintPanel(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsPrintPanel:"), value)
}/* debug [instance_properties/setter]: showsPrintPanel */


// A Boolean value that determines whether the print operation displays a progress panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/showsProgressPanel
func (p_ PrintOperation) ShowsProgressPanel() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsProgressPanel"))
	return rv
}/* debug [instance_properties/getter]: showsProgressPanel */


// A Boolean value that determines whether the print operation displays a progress panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/showsProgressPanel
func (p_ PrintOperation) SetShowsProgressPanel(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsProgressPanel:"), value)
}/* debug [instance_properties/setter]: showsProgressPanel */


// The view object that generates the actual data for the print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/view
func (p_ PrintOperation) View() IView {
	rv := objc.Send[View](p_.ID, objc.Sel("view"))
	return rv
}/* debug [instance_properties/getter]: view */


// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/iscopyingoperation
func (p_ PrintOperation) IsCopyingOperation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCopyingOperation"))
	return rv
}/* debug [instance_properties/getter]: isCopyingOperation */


// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/iscopyingoperation
func (p_ PrintOperation) SetIsCopyingOperation(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCopyingOperation:"), value)
}/* debug [instance_properties/setter]: isCopyingOperation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPrintOperation */


