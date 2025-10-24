// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class PDFView */


/* debug [class_header]: Header for PDFView */
// The class instance for the [PDFView] class.
var (
	PDFViewClass     _PDFViewClass
	PDFViewClassOnce sync.Once
)

func getPDFViewClass() _PDFViewClass {
	PDFViewClassOnce.Do(func() {
		PDFViewClass = _PDFViewClass{objc.GetClass("PDFView")}
	})
	return PDFViewClass
}

type _PDFViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFView */
// An interface definition for the [PDFView] class.
type IPDFView interface {
	appkit.IView
	
/* debug [class_interface_properties]: Properties for PDFView */
	// properties:
	AcceptsDraggedFiles() bool
	SetAcceptsDraggedFiles(value bool)
	AllowsDragging() bool
	SetAllowsDragging(value bool)
	AutoScales() bool
	SetAutoScales(value bool)
	BackgroundColor() appkit.Color
	SetBackgroundColor(value appkit.Color)
	CanGoBack() bool
	CanGoForward() bool
	CanGoToFirstPage() bool
	CanGoToLastPage() bool
	CanGoToNextPage() bool
	CanGoToPreviousPage() bool
	CanZoomIn() bool
	CanZoomOut() bool
	CurrentDestination() IPDFDestination
	CurrentPage() IPDFPage
	CurrentSelection() IPDFSelection
	SetCurrentSelection(value IPDFSelection)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DisplayBox() PDFDisplayBox
	SetDisplayBox(value PDFDisplayBox)
	DisplayDirection() PDFDisplayDirection
	SetDisplayDirection(value PDFDisplayDirection)
	DisplayMode() PDFDisplayMode
	SetDisplayMode(value PDFDisplayMode)
	DisplaysAsBook() bool
	SetDisplaysAsBook(value bool)
	DisplaysPageBreaks() bool
	SetDisplaysPageBreaks(value bool)
	DisplaysRTL() bool
	SetDisplaysRTL(value bool)
	Document() IPDFDocument
	SetDocument(value IPDFDocument)
	DocumentView() appkit.View
	EnableDataDetectors() bool
	SetEnableDataDetectors(value bool)
	GreekingThreshold() float64
	SetGreekingThreshold(value float64)
	HighlightedSelections() []PDFSelection
	SetHighlightedSelections(value []PDFSelection)
	InterpolationQuality() PDFInterpolationQuality
	SetInterpolationQuality(value PDFInterpolationQuality)
	InMarkupMode() bool
	SetInMarkupMode(value bool)
	MaxScaleFactor() float64
	SetMaxScaleFactor(value float64)
	MinScaleFactor() float64
	SetMinScaleFactor(value float64)
	PageBreakMargins() foundation.EdgeInsets
	SetPageBreakMargins(value foundation.EdgeInsets)
	PageOverlayViewProvider() unsafe.Pointer
	SetPageOverlayViewProvider(value unsafe.Pointer)
	PageShadowsEnabled() bool
	SetPageShadowsEnabled(value bool)
	ScaleFactor() float64
	SetScaleFactor(value float64)
	ScaleFactorForSizeToFit() float64
	ShouldAntiAlias() bool
	SetShouldAntiAlias(value bool)
	VisiblePages() []PDFPage
	IsFindInteractionEnabled() bool
	SetIsFindInteractionEnabled(value bool)
	IsInMarkupMode() bool
	SetIsInMarkupMode(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFView */
	// methods:
	AnnotationsChangedOnPage(page IPDFPage)
	AreaOfInterestForPoint(cursorLocation corefoundation.CGPoint) PDFAreaOfInterest
	AreaOfInterestForMouse(event appkit.Event) PDFAreaOfInterest
	ClearSelection()
	ConvertPointFromPage(point corefoundation.CGPoint, page IPDFPage) corefoundation.CGPoint
	ConvertRectFromPage(rect Rect /* not a class type */, page IPDFPage) Rect /* not a class type */
	ConvertRectToPage(rect Rect /* not a class type */, page IPDFPage) Rect /* not a class type */
	ConvertPointToPage(point corefoundation.CGPoint, page IPDFPage) corefoundation.CGPoint
	Copy(sender objc.IObject)
	DrawPageToContext(page IPDFPage, context ContextRef /* not a class type */)
	DrawPagePostToContext(page IPDFPage, context ContextRef /* not a class type */)
	GoToSelection(selection IPDFSelection)
	GoToDestination(destination IPDFDestination)
	GoToPage(page IPDFPage)
	GoToRectOnPage(rect Rect /* not a class type */, page IPDFPage)
	GoBack(sender objc.IObject)
	GoForward(sender objc.IObject)
	GoToFirstPage(sender objc.IObject)
	GoToLastPage(sender objc.IObject)
	GoToNextPage(sender objc.IObject)
	GoToPreviousPage(sender objc.IObject)
	LayoutDocumentView()
	PageForPointNearest(point vision.Point, nearest bool) IPDFPage
	PerformAction(action IPDFAction)
	PrintWithInfoAutoRotate(printInfo appkit.PrintInfo, doRotate bool)
	PrintWithInfoAutoRotatePageScaling(printInfo appkit.PrintInfo, doRotate bool, scale PDFPrintScalingMode)
	RowSizeForPage(page IPDFPage) corefoundation.CGSize
	ScrollSelectionToVisible(sender objc.IObject)
	SelectAll(sender objc.IObject)
	SetCurrentSelectionAnimate(selection IPDFSelection, animate bool)
	SetCursorForAreaOfInterest(area PDFAreaOfInterest)
	ZoomIn(sender objc.IObject)
	ZoomOut(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFView */
// Alloc allocates a new instance without initialization.
func (pc _PDFViewClass) Alloc() PDFView {
	rv := objc.Send[PDFView](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFViewClass) New() PDFView {
	rv := objc.Send[PDFView](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFView) Init() PDFView {
	rv := objc.Send[PDFView](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFView) Autorelease() PDFView {
	rv := objc.Send[PDFView](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFView creates a new PDFView instance.
func NewPDFView() PDFView {
	return getPDFViewClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFView */
// An object that encapsulates the functionality of PDF Kit into a single widget that you can add to your application using Interface Builder.
//
// may be the only class you need to deal with for adding PDF functionality to your application. It lets you display PDF data and allows users to select content, navigate through a document, set zoom level, and copy textual content to the Pasteboard. also keeps track of page history. You can subclass to create a custom PDF viewer. You can also create a custom PDF viewer by using the PDF Kit utility classes directly and not using at all.


// An object that encapsulates the functionality of PDF Kit into a single widget that you can add to your application using Interface Builder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView
type PDFView struct {
	appkit.View
}

// PDFViewFrom constructs a [PDFView] from an unsafe.Pointer.
//
// An object that encapsulates the functionality of PDF Kit into a single widget that you can add to your application using Interface Builder.
func PDFViewFrom(ptr unsafe.Pointer) PDFView {
	return PDFView{
		View: appkit.ViewFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFView *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFView */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFView */

// Tells the PDF view that an annotation on the specified page has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/annotationsChanged(on:)
func (p_ PDFView) AnnotationsChangedOnPage(page IPDFPage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("annotationsChangedOnPage:"), page)
}/* debug [instance_methods/method]: AnnotationsChangedOnPage */


// Returns the type of area for a specific cursor location point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/areaOfInterest(for:)
func (p_ PDFView) AreaOfInterestForPoint(cursorLocation corefoundation.CGPoint) PDFAreaOfInterest {
	rv := objc.Send[PDFAreaOfInterest](p_.ID, objc.Sel("areaOfInterestForPoint:"), cursorLocation)
	return rv
}/* debug [instance_methods/method]: AreaOfInterestForPoint */


// Returns the type of area the mouse cursor is over.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/areaOfInterest(forMouse:)
func (p_ PDFView) AreaOfInterestForMouse(event appkit.Event) PDFAreaOfInterest {
	rv := objc.Send[PDFAreaOfInterest](p_.ID, objc.Sel("areaOfInterestForMouse:"), event)
	return rv
}/* debug [instance_methods/method]: AreaOfInterestForMouse */


// Clears the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/clearSelection()
func (p_ PDFView) ClearSelection() {
	objc.Send[objc.ID](p_.ID, objc.Sel("clearSelection"))
}/* debug [instance_methods/method]: ClearSelection */


// Converts a point from page space to view space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/convert(_:from:)-4evlx
func (p_ PDFView) ConvertPointFromPage(point corefoundation.CGPoint, page IPDFPage) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p_.ID, objc.Sel("convertPoint:fromPage:"), point, page)
	return rv
}/* debug [instance_methods/method]: ConvertPointFromPage */


// Converts a rectangle from page space to view space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/convert(_:from:)-9xv1z
func (p_ PDFView) ConvertRectFromPage(rect Rect /* not a class type */, page IPDFPage) Rect /* not a class type */ {
	rv := objc.Send[Rect](p_.ID, objc.Sel("convertRect:fromPage:"), rect, page)
	return rv
}/* debug [instance_methods/method]: ConvertRectFromPage */


// Converts a rectangle from view space to page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/convert(_:to:)-8cp0c
func (p_ PDFView) ConvertRectToPage(rect Rect /* not a class type */, page IPDFPage) Rect /* not a class type */ {
	rv := objc.Send[Rect](p_.ID, objc.Sel("convertRect:toPage:"), rect, page)
	return rv
}/* debug [instance_methods/method]: ConvertRectToPage */


// Converts a point from view space to page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/convert(_:to:)-9twqk
func (p_ PDFView) ConvertPointToPage(point corefoundation.CGPoint, page IPDFPage) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p_.ID, objc.Sel("convertPoint:toPage:"), point, page)
	return rv
}/* debug [instance_methods/method]: ConvertPointToPage */


// Copies the text in the selection, if any, to the Pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/copy(_:)
func (p_ PDFView) Copy(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("copy:"), sender)
}/* debug [instance_methods/method]: Copy */


// Draw and render a visible page to a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/draw(_:to:)
func (p_ PDFView) DrawPageToContext(page IPDFPage, context ContextRef /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawPage:toContext:"), page, context)
}/* debug [instance_methods/method]: DrawPageToContext */


// Perform post-page rendering for a page rendered to a context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/drawPagePost(_:to:)
func (p_ PDFView) DrawPagePostToContext(page IPDFPage, context ContextRef /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("drawPagePost:toContext:"), page, context)
}/* debug [instance_methods/method]: DrawPagePostToContext */


// Scrolls to the first character of the specified selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/go(to:)-3t5go
func (p_ PDFView) GoToSelection(selection IPDFSelection) {
	objc.Send[objc.ID](p_.ID, objc.Sel("goToSelection:"), selection)
}/* debug [instance_methods/method]: GoToSelection */


// Navigates to the specified destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/go(to:)-5lh5d
func (p_ PDFView) GoToDestination(destination IPDFDestination) {
	objc.Send[objc.ID](p_.ID, objc.Sel("goToDestination:"), destination)
}/* debug [instance_methods/method]: GoToDestination */


// Scrolls to the specified page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/go(to:)-6x8y2
func (p_ PDFView) GoToPage(page IPDFPage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("goToPage:"), page)
}/* debug [instance_methods/method]: GoToPage */


// Navigates to the specified rectangle on the specified page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/go(to:on:)
func (p_ PDFView) GoToRectOnPage(rect Rect /* not a class type */, page IPDFPage) {
	objc.Send[objc.ID](p_.ID, objc.Sel("goToRect:onPage:"), rect, page)
}/* debug [instance_methods/method]: GoToRectOnPage */


// Navigates back one step in the page history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/goBack(_:)
func (p_ PDFView) GoBack(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("goBack:"), sender)
}/* debug [instance_methods/method]: GoBack */


// Navigates forward one step in the page history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/goForward(_:)
func (p_ PDFView) GoForward(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("goForward:"), sender)
}/* debug [instance_methods/method]: GoForward */


// Navigates to the first page of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/goToFirstPage(_:)
func (p_ PDFView) GoToFirstPage(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("goToFirstPage:"), sender)
}/* debug [instance_methods/method]: GoToFirstPage */


// Navigates to the last page of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/goToLastPage(_:)
func (p_ PDFView) GoToLastPage(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("goToLastPage:"), sender)
}/* debug [instance_methods/method]: GoToLastPage */


// Navigates to the next page of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/goToNextPage(_:)
func (p_ PDFView) GoToNextPage(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("goToNextPage:"), sender)
}/* debug [instance_methods/method]: GoToNextPage */


// Navigates to the previous page of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/goToPreviousPage(_:)
func (p_ PDFView) GoToPreviousPage(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("goToPreviousPage:"), sender)
}/* debug [instance_methods/method]: GoToPreviousPage */


// Performs layout of the inner views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/layoutDocumentView()
func (p_ PDFView) LayoutDocumentView() {
	objc.Send[objc.ID](p_.ID, objc.Sel("layoutDocumentView"))
}/* debug [instance_methods/method]: LayoutDocumentView */


// Returns the page containing a point specified in view coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/page(for:nearest:)
func (p_ PDFView) PageForPointNearest(point vision.Point, nearest bool) IPDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("pageForPoint:nearest:"), point, nearest)
	return rv
}/* debug [instance_methods/method]: PageForPointNearest */


// Performs the specified action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/perform(_:)
func (p_ PDFView) PerformAction(action IPDFAction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("performAction:"), action)
}/* debug [instance_methods/method]: PerformAction */


// Prints the document with the specified printer information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/print(with:autoRotate:)
func (p_ PDFView) PrintWithInfoAutoRotate(printInfo appkit.PrintInfo, doRotate bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("printWithInfo:autoRotate:"), printInfo, doRotate)
}/* debug [instance_methods/method]: PrintWithInfoAutoRotate */


// Prints the document with the specified printer and page-scaling information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/print(with:autoRotate:pageScaling:)
func (p_ PDFView) PrintWithInfoAutoRotatePageScaling(printInfo appkit.PrintInfo, doRotate bool, scale PDFPrintScalingMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("printWithInfo:autoRotate:pageScaling:"), printInfo, doRotate, scale)
}/* debug [instance_methods/method]: PrintWithInfoAutoRotatePageScaling */


// Returns the size needed to display a row of the current document page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/rowSize(for:)
func (p_ PDFView) RowSizeForPage(page IPDFPage) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p_.ID, objc.Sel("rowSizeForPage:"), page)
	return rv
}/* debug [instance_methods/method]: RowSizeForPage */


// Scrolls the view until the selection is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/scrollSelectionToVisible(_:)
func (p_ PDFView) ScrollSelectionToVisible(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("scrollSelectionToVisible:"), sender)
}/* debug [instance_methods/method]: ScrollSelectionToVisible */


// Selects all text in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/selectAll(_:)
func (p_ PDFView) SelectAll(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("selectAll:"), sender)
}/* debug [instance_methods/method]: SelectAll */


// Sets the current selection, in an animated way, if desired.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/setCurrentSelection(_:animate:)
func (p_ PDFView) SetCurrentSelectionAnimate(selection IPDFSelection, animate bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentSelection:animate:"), selection, animate)
}/* debug [instance_methods/method]: SetCurrentSelectionAnimate */


// Sets the type of mouse cursor according to the type of area the mouse cursor is over.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/setCursorFor(_:)
func (p_ PDFView) SetCursorForAreaOfInterest(area PDFAreaOfInterest) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCursorForAreaOfInterest:"), area)
}/* debug [instance_methods/method]: SetCursorForAreaOfInterest */


// Zooms in by increasing the scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/zoomIn(_:)
func (p_ PDFView) ZoomIn(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("zoomIn:"), sender)
}/* debug [instance_methods/method]: ZoomIn */


// Zooms out by decreasing the scaling factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/zoomOut(_:)
func (p_ PDFView) ZoomOut(sender objc.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("zoomOut:"), sender)
}/* debug [instance_methods/method]: ZoomOut */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFView */

// A Boolean value indicating whether you can drag a file into the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/acceptsDraggedFiles
func (p_ PDFView) AcceptsDraggedFiles() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("acceptsDraggedFiles"))
	return rv
}/* debug [instance_properties/getter]: acceptsDraggedFiles */


// A Boolean value indicating whether you can drag a file into the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/acceptsDraggedFiles
func (p_ PDFView) SetAcceptsDraggedFiles(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAcceptsDraggedFiles:"), value)
}/* debug [instance_properties/setter]: acceptsDraggedFiles */


// A Boolean value indicating whether the view can accept new PDF documents dragged into it by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/allowsDragging
func (p_ PDFView) AllowsDragging() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsDragging"))
	return rv
}/* debug [instance_properties/getter]: allowsDragging */


// A Boolean value indicating whether the view can accept new PDF documents dragged into it by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/allowsDragging
func (p_ PDFView) SetAllowsDragging(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAllowsDragging:"), value)
}/* debug [instance_properties/setter]: allowsDragging */


// A Boolean value indicating whether autoscaling is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/autoScales
func (p_ PDFView) AutoScales() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("autoScales"))
	return rv
}/* debug [instance_properties/getter]: autoScales */


// A Boolean value indicating whether autoscaling is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/autoScales
func (p_ PDFView) SetAutoScales(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAutoScales:"), value)
}/* debug [instance_properties/setter]: autoScales */


// The view’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/backgroundColor
func (p_ PDFView) BackgroundColor() appkit.Color {
	rv := objc.Send[appkit.Color](p_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The view’s background color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/backgroundColor
func (p_ PDFView) SetBackgroundColor(value appkit.Color) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// Returns a Boolean value indicating whether the user can navigate to the previous page in the page history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/canGoBack
func (p_ PDFView) CanGoBack() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canGoBack"))
	return rv
}/* debug [instance_properties/getter]: canGoBack */


// Returns a Boolean value indicating whether the user can navigate to the next page in the page history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/canGoForward
func (p_ PDFView) CanGoForward() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canGoForward"))
	return rv
}/* debug [instance_properties/getter]: canGoForward */


// Returns a Boolean value indicating whether the user can navigate to the first page of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/canGoToFirstPage
func (p_ PDFView) CanGoToFirstPage() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canGoToFirstPage"))
	return rv
}/* debug [instance_properties/getter]: canGoToFirstPage */


// Returns a Boolean value indicating whether the user can navigate to the last page of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/canGoToLastPage
func (p_ PDFView) CanGoToLastPage() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canGoToLastPage"))
	return rv
}/* debug [instance_properties/getter]: canGoToLastPage */


// Returns a Boolean value indicating whether the user can navigate to the next page of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/canGoToNextPage
func (p_ PDFView) CanGoToNextPage() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canGoToNextPage"))
	return rv
}/* debug [instance_properties/getter]: canGoToNextPage */


// Returns a Boolean value indicating whether the user can navigate to the previous page of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/canGoToPreviousPage
func (p_ PDFView) CanGoToPreviousPage() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canGoToPreviousPage"))
	return rv
}/* debug [instance_properties/getter]: canGoToPreviousPage */


// Returns a Boolean value indicating whether the user can magnify the view and zoom in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/canZoomIn
func (p_ PDFView) CanZoomIn() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canZoomIn"))
	return rv
}/* debug [instance_properties/getter]: canZoomIn */


// Returns a Boolean value indicating whether the user can view an expanded area and zoom out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/canZoomOut
func (p_ PDFView) CanZoomOut() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canZoomOut"))
	return rv
}/* debug [instance_properties/getter]: canZoomOut */


// Returns a object representing the current page and the current point in the view specified in page space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/currentDestination
func (p_ PDFView) CurrentDestination() IPDFDestination {
	rv := objc.Send[PDFDestination](p_.ID, objc.Sel("currentDestination"))
	return rv
}/* debug [instance_properties/getter]: currentDestination */


// Returns the current page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/currentPage
func (p_ PDFView) CurrentPage() IPDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("currentPage"))
	return rv
}/* debug [instance_properties/getter]: currentPage */


// The current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/currentSelection
func (p_ PDFView) CurrentSelection() IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("currentSelection"))
	return rv
}/* debug [instance_properties/getter]: currentSelection */


// The current selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/currentSelection
func (p_ PDFView) SetCurrentSelection(value IPDFSelection) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCurrentSelection:"), value)
}/* debug [instance_properties/setter]: currentSelection */


// Returns the view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/delegate
func (p_ PDFView) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Returns the view’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/delegate
func (p_ PDFView) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The current style of display box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displayBox
func (p_ PDFView) DisplayBox() PDFDisplayBox {
	rv := objc.Send[PDFDisplayBox](p_.ID, objc.Sel("displayBox"))
	return rv
}/* debug [instance_properties/getter]: displayBox */


// The current style of display box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displayBox
func (p_ PDFView) SetDisplayBox(value PDFDisplayBox) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayBox:"), value)
}/* debug [instance_properties/setter]: displayBox */


// The layout direction, either vertical or horizontal, for the given display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displayDirection
func (p_ PDFView) DisplayDirection() PDFDisplayDirection {
	rv := objc.Send[PDFDisplayDirection](p_.ID, objc.Sel("displayDirection"))
	return rv
}/* debug [instance_properties/getter]: displayDirection */


// The layout direction, either vertical or horizontal, for the given display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displayDirection
func (p_ PDFView) SetDisplayDirection(value PDFDisplayDirection) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayDirection:"), value)
}/* debug [instance_properties/setter]: displayDirection */


// The current display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displayMode
func (p_ PDFView) DisplayMode() PDFDisplayMode {
	rv := objc.Send[PDFDisplayMode](p_.ID, objc.Sel("displayMode"))
	return rv
}/* debug [instance_properties/getter]: displayMode */


// The current display mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displayMode
func (p_ PDFView) SetDisplayMode(value PDFDisplayMode) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayMode:"), value)
}/* debug [instance_properties/setter]: displayMode */


// A Boolean value indicating whether the view will display the first page as a book cover (meaningful only when the document is in two-up or two-up continuous display mode).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displaysAsBook
func (p_ PDFView) DisplaysAsBook() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("displaysAsBook"))
	return rv
}/* debug [instance_properties/getter]: displaysAsBook */


// A Boolean value indicating whether the view will display the first page as a book cover (meaningful only when the document is in two-up or two-up continuous display mode).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displaysAsBook
func (p_ PDFView) SetDisplaysAsBook(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplaysAsBook:"), value)
}/* debug [instance_properties/setter]: displaysAsBook */


// A Boolean value indicating whether the view is displaying page breaks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displaysPageBreaks
func (p_ PDFView) DisplaysPageBreaks() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("displaysPageBreaks"))
	return rv
}/* debug [instance_properties/getter]: displaysPageBreaks */


// A Boolean value indicating whether the view is displaying page breaks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displaysPageBreaks
func (p_ PDFView) SetDisplaysPageBreaks(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplaysPageBreaks:"), value)
}/* debug [instance_properties/setter]: displaysPageBreaks */


// The presentation of pages from right-to-left.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displaysRTL
func (p_ PDFView) DisplaysRTL() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("displaysRTL"))
	return rv
}/* debug [instance_properties/getter]: displaysRTL */


// The presentation of pages from right-to-left.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/displaysRTL
func (p_ PDFView) SetDisplaysRTL(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplaysRTL:"), value)
}/* debug [instance_properties/setter]: displaysRTL */


// Returns the document associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/document
func (p_ PDFView) Document() IPDFDocument {
	rv := objc.Send[PDFDocument](p_.ID, objc.Sel("document"))
	return rv
}/* debug [instance_properties/getter]: document */


// Returns the document associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/document
func (p_ PDFView) SetDocument(value IPDFDocument) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDocument:"), value)
}/* debug [instance_properties/setter]: document */


// The innermost view used by or by your subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/documentView
func (p_ PDFView) DocumentView() appkit.View {
	rv := objc.Send[appkit.View](p_.ID, objc.Sel("documentView"))
	return rv
}/* debug [instance_properties/getter]: documentView */


// A Boolean value indicating whether to turns on or off data detection, which adds annotations for detected URLs in a page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/enableDataDetectors
func (p_ PDFView) EnableDataDetectors() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("enableDataDetectors"))
	return rv
}/* debug [instance_properties/getter]: enableDataDetectors */


// A Boolean value indicating whether to turns on or off data detection, which adds annotations for detected URLs in a page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/enableDataDetectors
func (p_ PDFView) SetEnableDataDetectors(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEnableDataDetectors:"), value)
}/* debug [instance_properties/setter]: enableDataDetectors */


// Returns the current greeking threshold for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/greekingThreshold
func (p_ PDFView) GreekingThreshold() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("greekingThreshold"))
	return rv
}/* debug [instance_properties/getter]: greekingThreshold */


// Returns the current greeking threshold for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/greekingThreshold
func (p_ PDFView) SetGreekingThreshold(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGreekingThreshold:"), value)
}/* debug [instance_properties/setter]: greekingThreshold */


// Returns the array of selections that are highlighted using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/highlightedSelections
func (p_ PDFView) HighlightedSelections() []PDFSelection {
	rv := objc.Send[[]PDFSelection](p_.ID, objc.Sel("highlightedSelections"))
	return rv
}/* debug [instance_properties/getter]: highlightedSelections */


// Returns the array of selections that are highlighted using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/highlightedSelections
func (p_ PDFView) SetHighlightedSelections(value []PDFSelection) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setHighlightedSelections:"), nsArray)
}/* debug [instance_properties/setter]: highlightedSelections */


// The interpolation quality for images drawn into the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/interpolationQuality
func (p_ PDFView) InterpolationQuality() PDFInterpolationQuality {
	rv := objc.Send[PDFInterpolationQuality](p_.ID, objc.Sel("interpolationQuality"))
	return rv
}/* debug [instance_properties/getter]: interpolationQuality */


// The interpolation quality for images drawn into the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/interpolationQuality
func (p_ PDFView) SetInterpolationQuality(value PDFInterpolationQuality) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInterpolationQuality:"), value)
}/* debug [instance_properties/setter]: interpolationQuality */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/isInMarkupMode
func (p_ PDFView) InMarkupMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("inMarkupMode"))
	return rv
}/* debug [instance_properties/getter]: inMarkupMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/isInMarkupMode
func (p_ PDFView) SetInMarkupMode(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInMarkupMode:"), value)
}/* debug [instance_properties/setter]: inMarkupMode */


// The maximum scaling factor for the PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/maxScaleFactor
func (p_ PDFView) MaxScaleFactor() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("maxScaleFactor"))
	return rv
}/* debug [instance_properties/getter]: maxScaleFactor */


// The maximum scaling factor for the PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/maxScaleFactor
func (p_ PDFView) SetMaxScaleFactor(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMaxScaleFactor:"), value)
}/* debug [instance_properties/setter]: maxScaleFactor */


// The minimum scaling factor for the PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/minScaleFactor
func (p_ PDFView) MinScaleFactor() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minScaleFactor"))
	return rv
}/* debug [instance_properties/getter]: minScaleFactor */


// The minimum scaling factor for the PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/minScaleFactor
func (p_ PDFView) SetMinScaleFactor(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMinScaleFactor:"), value)
}/* debug [instance_properties/setter]: minScaleFactor */


// The spacing between pages as defined by the top, bottom, left, and right margins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageBreakMargins
func (p_ PDFView) PageBreakMargins() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](p_.ID, objc.Sel("pageBreakMargins"))
	return rv
}/* debug [instance_properties/getter]: pageBreakMargins */


// The spacing between pages as defined by the top, bottom, left, and right margins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageBreakMargins
func (p_ PDFView) SetPageBreakMargins(value foundation.EdgeInsets) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageBreakMargins:"), value)
}/* debug [instance_properties/setter]: pageBreakMargins */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageOverlayViewProvider
func (p_ PDFView) PageOverlayViewProvider() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("pageOverlayViewProvider"))
	return rv
}/* debug [instance_properties/getter]: pageOverlayViewProvider */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageOverlayViewProvider
func (p_ PDFView) SetPageOverlayViewProvider(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageOverlayViewProvider:"), value)
}/* debug [instance_properties/setter]: pageOverlayViewProvider */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageShadowsEnabled
func (p_ PDFView) PageShadowsEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("pageShadowsEnabled"))
	return rv
}/* debug [instance_properties/getter]: pageShadowsEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/pageShadowsEnabled
func (p_ PDFView) SetPageShadowsEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPageShadowsEnabled:"), value)
}/* debug [instance_properties/setter]: pageShadowsEnabled */


// The current scale factor for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/scaleFactor
func (p_ PDFView) ScaleFactor() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("scaleFactor"))
	return rv
}/* debug [instance_properties/getter]: scaleFactor */


// The current scale factor for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/scaleFactor
func (p_ PDFView) SetScaleFactor(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setScaleFactor:"), value)
}/* debug [instance_properties/setter]: scaleFactor */


// The “size to fit” scale factor that would use for scaling the current document and layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/scaleFactorForSizeToFit
func (p_ PDFView) ScaleFactorForSizeToFit() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("scaleFactorForSizeToFit"))
	return rv
}/* debug [instance_properties/getter]: scaleFactorForSizeToFit */


// A Boolean value indicating whether the view is antialiased.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/shouldAntiAlias
func (p_ PDFView) ShouldAntiAlias() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldAntiAlias"))
	return rv
}/* debug [instance_properties/getter]: shouldAntiAlias */


// A Boolean value indicating whether the view is antialiased.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/shouldAntiAlias
func (p_ PDFView) SetShouldAntiAlias(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldAntiAlias:"), value)
}/* debug [instance_properties/setter]: shouldAntiAlias */


// Returns an array of objects that represent the currently visible pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFView/visiblePages
func (p_ PDFView) VisiblePages() []PDFPage {
	rv := objc.Send[[]PDFPage](p_.ID, objc.Sel("visiblePages"))
	return rv
}/* debug [instance_properties/getter]: visiblePages */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isfindinteractionenabled
func (p_ PDFView) IsFindInteractionEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFindInteractionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isFindInteractionEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isfindinteractionenabled
func (p_ PDFView) SetIsFindInteractionEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFindInteractionEnabled:"), value)
}/* debug [instance_properties/setter]: isFindInteractionEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isinmarkupmode
func (p_ PDFView) IsInMarkupMode() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isInMarkupMode"))
	return rv
}/* debug [instance_properties/getter]: isInMarkupMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/pdfkit/pdfview/isinmarkupmode
func (p_ PDFView) SetIsInMarkupMode(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsInMarkupMode:"), value)
}/* debug [instance_properties/setter]: isInMarkupMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFView */


