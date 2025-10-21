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

// An interface definition for the [PrintOperation] class.
type IPrintOperation interface {
	objectivec.IObject
	CleanUpOperation()
	RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow IWindow, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo unsafe.Pointer)
}

// An object that controls operations that generate Encapsulated PostScript (EPS) code, Portable Document Format (PDF) code, or print jobs.
//
// An object works in conjunction with two other objects: an object, which specifies how the code should be generated, and an object, which generates the actual code. It is important to note that the majority of methods in copy the instance of passed into them. Future changes to that print info are not reflected in the print info retained by the current object. All changes should be made to the print info before passing to the methods of this class. The only method in which does not copy the instance is .
//
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

// Alloc allocates a new instance without initialization.
func (pc _PrintOperationClass) Alloc() PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates and returns a new print operation object ready to control the copying of EPS graphics from the specified view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/epsOperation(with:inside:to:)
func (pc _PrintOperationClass) EPSOperationWithViewInsideRectToData(view IView, rect coregraphics.CGRect, data IMutableData) PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("EPSOperationWithView:insideRect:toData:"), view, rect, data)
	return rv
}

// Called at the end of a print operation to remove the print operation as the current operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/cleanUp()
func (p_ PrintOperation) CleanUpOperation() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cleanUpOperation"))
}

// Runs the print operation, calling your custom delegate method upon completion.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/runModal(for:delegate:didRun:contextInfo:)
func (p_ PrintOperation) RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow IWindow, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("runOperationModalForWindow:delegate:didRunSelector:contextInfo:"), docWindow, delegate, didRunSelector, contextInfo)
}

// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/isCopyingOperation
func (p_ PrintOperation) CopyingOperation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("copyingOperation"))
	return rv
}

// The PDF panel object to use during the operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfPanel
func (p_ PrintOperation) PDFPanel() NSPDFPanel {
	rv := objc.Send[NSPDFPanel](p_.ID, objc.Sel("PDFPanel"))
	return rv
}


// SetPDFPanel sets the value of the PDFPanel property.
// The PDF panel object to use during the operation.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfPanel
func (p_ PrintOperation) SetPDFPanel(value IPDFPanel) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPDFPanel:"), value)
}

// The printing information associated with the print operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printInfo
func (p_ PrintOperation) PrintInfo() NSPrintInfo {
	rv := objc.Send[NSPrintInfo](p_.ID, objc.Sel("printInfo"))
	return rv
}


// SetPrintInfo sets the value of the printInfo property.
// The printing information associated with the print operation.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printInfo
func (p_ PrintOperation) SetPrintInfo(value IPrintInfo) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrintInfo:"), value)
}

// The view object that generates the actual data for the print operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/view
func (p_ PrintOperation) View() NSView {
	rv := objc.Send[NSView](p_.ID, objc.Sel("view"))
	return rv
}

// A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/canspawnseparatethread
func (p_ PrintOperation) CanSpawnSeparateThread() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canSpawnSeparateThread"))
	return rv
}


// SetCanSpawnSeparateThread sets the value of the canSpawnSeparateThread property.
// A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/canspawnseparatethread
func (p_ PrintOperation) SetCanSpawnSeparateThread(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanSpawnSeparateThread:"), value)
}

// The graphics context object used for generating output.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/context
func (p_ PrintOperation) Context() NSGraphicsContext {
	rv := objc.Send[NSGraphicsContext](p_.ID, objc.Sel("context"))
	return rv
}


// SetContext sets the value of the context property.
// The graphics context object used for generating output.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/context
func (p_ PrintOperation) SetContext(value IGraphicsContext) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContext:"), value)
}

// The current page number being printed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/currentpage
func (p_ PrintOperation) CurrentPage() int {
	rv := objc.Send[int](p_.ID, objc.Sel("currentPage"))
	return rv
}


// SetCurrentPage sets the value of the currentPage property.
// The current page number being printed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/currentpage
func (p_ PrintOperation) SetCurrentPage(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentPage:"), value)
}

// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/iscopyingoperation
func (p_ PrintOperation) IsCopyingOperation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCopyingOperation"))
	return rv
}


// SetIsCopyingOperation sets the value of the isCopyingOperation property.
// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/iscopyingoperation
func (p_ PrintOperation) SetIsCopyingOperation(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCopyingOperation:"), value)
}

// The custom title of the print job.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/jobtitle
func (p_ PrintOperation) JobTitle() string {
	rv := objc.Send[string](p_.ID, objc.Sel("jobTitle"))
	return rv
}


// SetJobTitle sets the value of the jobTitle property.
// The custom title of the print job.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/jobtitle
func (p_ PrintOperation) SetJobTitle(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJobTitle:"), objc.String(value))
}

// The print order for the pages of the operation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/pageorder-swift.property
func (p_ PrintOperation) PageOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pageOrder"))
	return rv
}


// SetPageOrder sets the value of the pageOrder property.
// The print order for the pages of the operation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/pageorder-swift.property
func (p_ PrintOperation) SetPageOrder(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageOrder:"), value)
}

// The range of pages associated with the print operation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/pagerange
func (p_ PrintOperation) PageRange() foundation.Range {
	rv := objc.Send[foundation.Range](p_.ID, objc.Sel("pageRange"))
	return rv
}


// SetPageRange sets the value of the pageRange property.
// The range of pages associated with the print operation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/pagerange
func (p_ PrintOperation) SetPageRange(value foundation.IRange) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageRange:"), value)
}

// The printing quality.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/preferredrenderingquality
func (p_ PrintOperation) PreferredRenderingQuality() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("preferredRenderingQuality"))
	return rv
}


// SetPreferredRenderingQuality sets the value of the preferredRenderingQuality property.
// The printing quality.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/preferredrenderingquality
func (p_ PrintOperation) SetPreferredRenderingQuality(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredRenderingQuality:"), value)
}

// The print panel object to use during the operation.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/printpanel
func (p_ PrintOperation) PrintPanel() NSPrintPanel {
	rv := objc.Send[NSPrintPanel](p_.ID, objc.Sel("printPanel"))
	return rv
}


// SetPrintPanel sets the value of the printPanel property.
// The print panel object to use during the operation.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/printpanel
func (p_ PrintOperation) SetPrintPanel(value IPrintPanel) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrintPanel:"), value)
}

// A Boolean value that determines whether the print operation displays a print panel.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/showsprintpanel
func (p_ PrintOperation) ShowsPrintPanel() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsPrintPanel"))
	return rv
}


// SetShowsPrintPanel sets the value of the showsPrintPanel property.
// A Boolean value that determines whether the print operation displays a print panel.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/showsprintpanel
func (p_ PrintOperation) SetShowsPrintPanel(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsPrintPanel:"), value)
}

// A Boolean value that determines whether the print operation displays a progress panel.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/showsprogresspanel
func (p_ PrintOperation) ShowsProgressPanel() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsProgressPanel"))
	return rv
}


// SetShowsProgressPanel sets the value of the showsProgressPanel property.
// A Boolean value that determines whether the print operation displays a progress panel.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/showsprogresspanel
func (p_ PrintOperation) SetShowsProgressPanel(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsProgressPanel:"), value)
}



