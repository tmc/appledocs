// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	PrintInfo() IPrintInfo
	SetPrintInfo(value IPrintInfo)
	CanSpawnSeparateThread() bool /* primitive/slice/pointer. */
	SetCanSpawnSeparateThread(value bool /* primitive/slice/pointer. */)
	Context() IGraphicsContext
	SetContext(value IGraphicsContext)
	CurrentPage() int /* primitive/slice/pointer. */
	SetCurrentPage(value int /* primitive/slice/pointer. */)
	IsCopyingOperation() bool /* primitive/slice/pointer. */
	SetIsCopyingOperation(value bool /* primitive/slice/pointer. */)
	JobTitle() string /* primitive/slice/pointer. */
	SetJobTitle(value string /* primitive/slice/pointer. */)
	PageOrder() unsafe.Pointer
	SetPageOrder(value unsafe.Pointer)
	PageRange() foundation.objc.IObject /* cross-framework: Range */
	SetPageRange(value foundation.objc.IObject /* cross-framework: Range */)
	PdfPanel() objc.IObject /* cross-framework: PDFPanel */
	SetPdfPanel(value objc.IObject /* cross-framework: PDFPanel */)
	PreferredRenderingQuality() unsafe.Pointer
	SetPreferredRenderingQuality(value unsafe.Pointer)
	PrintPanel() IPrintPanel
	SetPrintPanel(value IPrintPanel)
	ShowsPrintPanel() bool /* primitive/slice/pointer. */
	SetShowsPrintPanel(value bool /* primitive/slice/pointer. */)
	ShowsProgressPanel() bool /* primitive/slice/pointer. */
	SetShowsProgressPanel(value bool /* primitive/slice/pointer. */)
	View() IView
	SetView(value IView)
	// methods:
}

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



// The printing information associated with the print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printInfo
func (p_ PrintOperation) PrintInfo() IPrintInfo {
	rv := objc.Send[PrintInfo](p_.ID, objc.Sel("printInfo"))
	return rv
}


// The printing information associated with the print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printInfo
func (p_ PrintOperation) SetPrintInfo(value IPrintInfo) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrintInfo:"), value)
}


// A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/canspawnseparatethread
func (p_ PrintOperation) CanSpawnSeparateThread() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("canSpawnSeparateThread"))
	return rv
}


// A Boolean value that determines whether the print operation is allowed to spawn a separate printing thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/canspawnseparatethread
func (p_ PrintOperation) SetCanSpawnSeparateThread(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCanSpawnSeparateThread:"), value)
}


// The graphics context object used for generating output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/context
func (p_ PrintOperation) Context() IGraphicsContext {
	rv := objc.Send[GraphicsContext](p_.ID, objc.Sel("context"))
	return rv
}


// The graphics context object used for generating output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/context
func (p_ PrintOperation) SetContext(value IGraphicsContext) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContext:"), value)
}


// The current page number being printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/currentpage
func (p_ PrintOperation) CurrentPage() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](p_.ID, objc.Sel("currentPage"))
	return rv
}


// The current page number being printed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/currentpage
func (p_ PrintOperation) SetCurrentPage(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentPage:"), value)
}


// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/iscopyingoperation
func (p_ PrintOperation) IsCopyingOperation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("isCopyingOperation"))
	return rv
}


// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/iscopyingoperation
func (p_ PrintOperation) SetIsCopyingOperation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsCopyingOperation:"), value)
}


// The custom title of the print job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/jobtitle
func (p_ PrintOperation) JobTitle() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](p_.ID, objc.Sel("jobTitle"))
	return rv
}


// The custom title of the print job.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/jobtitle
func (p_ PrintOperation) SetJobTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJobTitle:"), objc.String(value))
}


// The print order for the pages of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/pageorder-swift.property
func (p_ PrintOperation) PageOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pageOrder"))
	return rv
}


// The print order for the pages of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/pageorder-swift.property
func (p_ PrintOperation) SetPageOrder(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageOrder:"), value)
}


// The range of pages associated with the print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/pagerange
func (p_ PrintOperation) PageRange() foundation.objc.IObject /* cross-framework: Range */ {
	rv := objc.Send[foundation.Range](p_.ID, objc.Sel("pageRange"))
	return rv
}


// The range of pages associated with the print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/pagerange
func (p_ PrintOperation) SetPageRange(value foundation.objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageRange:"), value)
}


// The PDF panel object to use during the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/pdfpanel
func (p_ PrintOperation) PdfPanel() objc.IObject /* cross-framework: PDFPanel */ {
	rv := objc.Send[PDFPanel](p_.ID, objc.Sel("pdfPanel"))
	return rv
}


// The PDF panel object to use during the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/pdfpanel
func (p_ PrintOperation) SetPdfPanel(value objc.IObject /* cross-framework: PDFPanel */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPdfPanel:"), value)
}


// The printing quality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/preferredrenderingquality
func (p_ PrintOperation) PreferredRenderingQuality() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("preferredRenderingQuality"))
	return rv
}


// The printing quality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/preferredrenderingquality
func (p_ PrintOperation) SetPreferredRenderingQuality(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferredRenderingQuality:"), value)
}


// The print panel object to use during the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/printpanel
func (p_ PrintOperation) PrintPanel() IPrintPanel {
	rv := objc.Send[PrintPanel](p_.ID, objc.Sel("printPanel"))
	return rv
}


// The print panel object to use during the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/printpanel
func (p_ PrintOperation) SetPrintPanel(value IPrintPanel) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrintPanel:"), value)
}


// A Boolean value that determines whether the print operation displays a print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/showsprintpanel
func (p_ PrintOperation) ShowsPrintPanel() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsPrintPanel"))
	return rv
}


// A Boolean value that determines whether the print operation displays a print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/showsprintpanel
func (p_ PrintOperation) SetShowsPrintPanel(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsPrintPanel:"), value)
}


// A Boolean value that determines whether the print operation displays a progress panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/showsprogresspanel
func (p_ PrintOperation) ShowsProgressPanel() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](p_.ID, objc.Sel("showsProgressPanel"))
	return rv
}


// A Boolean value that determines whether the print operation displays a progress panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/showsprogresspanel
func (p_ PrintOperation) SetShowsProgressPanel(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShowsProgressPanel:"), value)
}


// The view object that generates the actual data for the print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/view
func (p_ PrintOperation) View() IView {
	rv := objc.Send[View](p_.ID, objc.Sel("view"))
	return rv
}


// The view object that generates the actual data for the print operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsprintoperation/view
func (p_ PrintOperation) SetView(value IView) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setView:"), value)
}



