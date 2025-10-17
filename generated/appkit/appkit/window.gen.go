// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Window] class.
var WindowClass objc.Class

func init() {
	WindowClass = objc.GetClass("NSWindow")
}

type Window struct {
	objc.ID
}

func WindowFrom(ptr unsafe.Pointer) Window {
	return Window{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (wc Window) Alloc() Window {
	ret := objc.ID(WindowClass).Send(objc.RegisterName("alloc"))
	return Window{ret}
}

// Init initializes the instance.
func (w_ Window) Init() Window {
	ret := w_.ID.Send(objc.RegisterName("init"))
	return Window{ret}
}
// Initializes the window with the specified values. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:)
func NewWindowWithContentRectStyleMaskBackingDefer(contentRect unsafe.Pointer, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	instance := Window{}.Alloc()
	sel := objc.RegisterName("initWithContentRect:styleMask:backing:defer:")
	ret := instance.ID.Send(sel, contentRect, style, backingStoreType, flag)
	instance = Window{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes an allocated window with the specified values. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:screen:)
func NewWindowWithContentRectStyleMaskBackingDeferScreen(contentRect unsafe.Pointer, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen unsafe.Pointer) Window {
	instance := Window{}.Alloc()
	sel := objc.RegisterName("initWithContentRect:styleMask:backing:defer:screen:")
	ret := instance.ID.Send(sel, contentRect, style, backingStoreType, flag, screen)
	instance = Window{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Returns a Cocoa window created from a Carbon window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/init(windowRef:)
func NewWindowWithWindowRef(windowRef unsafe.Pointer) Window {
	instance := Window{}.Alloc()
	sel := objc.RegisterName("initWithWindowRef:")
	ret := instance.ID.Send(sel, windowRef)
	instance = Window{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Returns the content rectangle used by a window with a given frame rectangle and window style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentRect(forFrameRect:styleMask:)
func (wc Window) ContentRectForFrameRectStyleMask(fRect unsafe.Pointer, style WindowStyleMask) unsafe.Pointer {
	sel := objc.RegisterName("contentRectForFrameRect:styleMask:")
	ret := objc.ID(WindowClass).Send(sel, fRect, style)
	return unsafe.Pointer(ret)
}
// Returns the frame rectangle used by a window with a given content rectangle and window style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/frameRect(forContentRect:styleMask:)
func (wc Window) FrameRectForContentRectStyleMask(cRect unsafe.Pointer, style WindowStyleMask) unsafe.Pointer {
	sel := objc.RegisterName("frameRectForContentRect:styleMask:")
	ret := objc.ID(WindowClass).Send(sel, cRect, style)
	return unsafe.Pointer(ret)
}
// Creates a titled window that contains the specified content view controller. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/init(contentViewController:)
func (wc Window) WindowWithContentViewController(contentViewController unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("windowWithContentViewController:")
	ret := objc.ID(WindowClass).Send(sel, contentViewController)
	return unsafe.Pointer(ret)
}
// This method does nothing; it is here for backward compatibility. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/menuChanged(_:)
func (wc Window) MenuChanged(menu unsafe.Pointer) {
	sel := objc.RegisterName("menuChanged:")
	objc.ID(WindowClass).Send(sel, menu)
}
// Returns the minimum width a window’s frame rectangle must have for it to display a title, with a given window style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/minFrameWidth(withTitle:styleMask:)
func (wc Window) MinFrameWidthWithTitleStyleMask(title string, style WindowStyleMask) float64 {
	sel := objc.RegisterName("minFrameWidthWithTitle:styleMask:")
	ret := objc.ID(WindowClass).Send(sel, title, style)
	return float64(ret)
}
// Removes the frame data stored under a given name from the application’s user defaults. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/removeFrame(usingName:)
func (wc Window) RemoveFrameUsingName(name unsafe.Pointer) {
	sel := objc.RegisterName("removeFrameUsingName:")
	objc.ID(WindowClass).Send(sel, name)
}
// Returns a new instance of a given standard window button, sized appropriately for a given window style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/standardWindowButton(_:for:)
func (wc Window) StandardWindowButtonForStyleMask(b unsafe.Pointer, styleMask WindowStyleMask) unsafe.Pointer {
	sel := objc.RegisterName("standardWindowButton:forStyleMask:")
	ret := objc.ID(WindowClass).Send(sel, b, styleMask)
	return unsafe.Pointer(ret)
}
// Returns the number of the frontmost window that would be hit by a mouse-down at the specified screen location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/windowNumber(at:belowWindowWithWindowNumber:)
func (wc Window) WindowNumberAtPointBelowWindowWithWindowNumber(point unsafe.Pointer, windowNumber int) int {
	sel := objc.RegisterName("windowNumberAtPoint:belowWindowWithWindowNumber:")
	ret := objc.ID(WindowClass).Send(sel, point, windowNumber)
	return int(ret)
}
// Returns the window numbers for all visible windows satisfying the specified options. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/windowNumbers(options:)
func (wc Window) WindowNumbersWithOptions(options unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("windowNumbersWithOptions:")
	ret := objc.ID(WindowClass).Send(sel, options)
	return unsafe.Pointer(ret)
}
// Adds a given window as a child window of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/addChildWindow(_:ordered:)
func (w_ Window) AddChildWindowOrdered(childWin unsafe.Pointer, place WindowOrderingMode) {
	sel := objc.RegisterName("addChildWindow:ordered:")
	w_.ID.Send(sel, childWin, place)
}
// Adds the provided window as a new tab in a tabbed window using the specified ordering instruction. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/addTabbedWindow(_:ordered:)
func (w_ Window) AddTabbedWindowOrdered(window unsafe.Pointer, ordered WindowOrderingMode) {
	sel := objc.RegisterName("addTabbedWindow:ordered:")
	w_.ID.Send(sel, window, ordered)
}
// Adds the specified title bar accessory view controller to the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/addTitlebarAccessoryViewController(_:)
func (w_ Window) AddTitlebarAccessoryViewController(childViewController unsafe.Pointer) {
	sel := objc.RegisterName("addTitlebarAccessoryViewController:")
	w_.ID.Send(sel, childViewController)
}
// Returns the part of the window that stays stationary during constraint-based layout. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/anchorAttribute(for:)
func (w_ Window) AnchorAttributeForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("anchorAttributeForOrientation:")
	ret := w_.ID.Send(sel, orientation)
	return unsafe.Pointer(ret)
}
// Specifies the duration of a smooth frame-size change. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/animationResizeTime(_:)
func (w_ Window) AnimationResizeTime(newFrame unsafe.Pointer) float64 {
	sel := objc.RegisterName("animationResizeTime:")
	ret := w_.ID.Send(sel, newFrame)
	return float64(ret)
}
// Indicates whether the window calculates the thickness of a given border automatically. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/autorecalculatesContentBorderThickness(for:)
func (w_ Window) AutorecalculatesContentBorderThicknessForEdge(edge int) bool {
	sel := objc.RegisterName("autorecalculatesContentBorderThicknessForEdge:")
	ret := w_.ID.Send(sel, edge)
	return ret != 0
}
// Returns a backing store pixel-aligned rectangle in window coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/backingAlignedRect(_:options:)
func (w_ Window) BackingAlignedRectOptions(rect unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("backingAlignedRect:options:")
	ret := w_.ID.Send(sel, rect, options)
	return unsafe.Pointer(ret)
}
// Informs the window that it has become the key window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/becomeKey()
func (w_ Window) BecomeKeyWindow() {
	sel := objc.RegisterName("becomeKeyWindow")
	w_.ID.Send(sel)
}
// Informs the window that it has become the main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/becomeMain()
func (w_ Window) BecomeMainWindow() {
	sel := objc.RegisterName("becomeMainWindow")
	w_.ID.Send(sel)
}
// Starts a document-modal session and presents the specified critical sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/beginCriticalSheet(_:completionHandler:)
func (w_ Window) BeginCriticalSheetCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer) {
	sel := objc.RegisterName("beginCriticalSheet:completionHandler:")
	w_.ID.Send(sel, sheetWindow, handler)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/beginDraggingSession(items:event:source:)
func (w_ Window) BeginDraggingSessionWithItemsEventSource(items unsafe.Pointer, event unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("beginDraggingSessionWithItems:event:source:")
	ret := w_.ID.Send(sel, items, event, source)
	return unsafe.Pointer(ret)
}
// Starts a document-modal session and presents—or queues for presentation—a sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/beginSheet(_:completionHandler:)
func (w_ Window) BeginSheetCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer) {
	sel := objc.RegisterName("beginSheet:completionHandler:")
	w_.ID.Send(sel, sheetWindow, handler)
}
// Stores the window’s raster image from a given rectangle expressed in the window’s base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/cacheImage(in:)
func (w_ Window) CacheImageInRect(rect unsafe.Pointer) {
	sel := objc.RegisterName("cacheImageInRect:")
	w_.ID.Send(sel, rect)
}
// A Boolean value that indicates if the window and its screen use a color space that can represent the specified display gamut. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/canRepresent(_:)
func (w_ Window) CanRepresentDisplayGamut(displayGamut unsafe.Pointer) bool {
	sel := objc.RegisterName("canRepresentDisplayGamut:")
	ret := w_.ID.Send(sel, displayGamut)
	return ret != 0
}
// Indicates whether the window has a depth limit that allows it to store color values. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/canStoreColor()
func (w_ Window) CanStoreColor() bool {
	sel := objc.RegisterName("canStoreColor")
	ret := w_.ID.Send(sel)
	return ret != 0
}
// Positions the window’s top-left to a given point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/cascadeTopLeft(from:)
func (w_ Window) CascadeTopLeftFromPoint(topLeftPoint unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("cascadeTopLeftFromPoint:")
	ret := w_.ID.Send(sel, topLeftPoint)
	return unsafe.Pointer(ret)
}
// Sets the window’s location to the center of the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/center()
func (w_ Window) Center() {
	sel := objc.RegisterName("center")
	w_.ID.Send(sel)
}
// Removes the window from the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/close()
func (w_ Window) Close() {
	sel := objc.RegisterName("close")
	w_.ID.Send(sel)
}
// Modifies and returns a frame rectangle so that its top edge lies on a specific screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/constrainFrameRect(_:to:)
func (w_ Window) ConstrainFrameRectToScreen(frameRect unsafe.Pointer, screen unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("constrainFrameRect:toScreen:")
	ret := w_.ID.Send(sel, frameRect, screen)
	return unsafe.Pointer(ret)
}
// Indicates the thickness of a given border of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentBorderThickness(for:)
func (w_ Window) ContentBorderThicknessForEdge(edge int) float64 {
	sel := objc.RegisterName("contentBorderThicknessForEdge:")
	ret := w_.ID.Send(sel, edge)
	return float64(ret)
}
// Returns the window’s content rectangle with a given frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentRect(forFrameRect:)
func (w_ Window) ContentRectForFrameRect(frameRect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("contentRectForFrameRect:")
	ret := w_.ID.Send(sel, frameRect)
	return unsafe.Pointer(ret)
}
// Converts a given point from the window’s base coordinate system to the screen coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertBaseToScreen:
func (w_ Window) ConvertBaseToScreen(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertBaseToScreen:")
	ret := w_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertFromBacking(_:)
func (w_ Window) ConvertRectFromBacking(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRectFromBacking:")
	ret := w_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Converts a rectangle from the screen coordinate system to the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertFromScreen(_:)
func (w_ Window) ConvertRectFromScreen(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRectFromScreen:")
	ret := w_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Converts a point from the screen coordinate system to the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertPoint(fromScreen:)
func (w_ Window) ConvertPointFromScreen(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPointFromScreen:")
	ret := w_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Converts a point to the screen coordinate system from the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertPoint(toScreen:)
func (w_ Window) ConvertPointToScreen(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPointToScreen:")
	ret := w_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertPointFromBacking(_:)
func (w_ Window) ConvertPointFromBacking(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPointFromBacking:")
	ret := w_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertPointToBacking(_:)
func (w_ Window) ConvertPointToBacking(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPointToBacking:")
	ret := w_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Converts a given point from the screen coordinate system to the window’s base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertScreenToBase:
func (w_ Window) ConvertScreenToBase(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertScreenToBase:")
	ret := w_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertToBacking(_:)
func (w_ Window) ConvertRectToBacking(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRectToBacking:")
	ret := w_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Converts a rectangle to the screen coordinate system from the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertToScreen(_:)
func (w_ Window) ConvertRectToScreen(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRectToScreen:")
	ret := w_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Returns EPS data that draws the region of the window within a given rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/dataWithEPS(inside:)
func (w_ Window) DataWithEPSInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataWithEPSInsideRect:")
	ret := w_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Returns PDF data that draws the region of the window within a given rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/dataWithPDF(inside:)
func (w_ Window) DataWithPDFInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataWithPDFInsideRect:")
	ret := w_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// De-minimizes the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/deminiaturize(_:)
func (w_ Window) Deminiaturize(sender objc.ID) {
	sel := objc.RegisterName("deminiaturize:")
	w_.ID.Send(sel, sender)
}
// Disables all cursor rectangle management within the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/disableCursorRects()
func (w_ Window) DisableCursorRects() {
	sel := objc.RegisterName("disableCursorRects")
	w_.ID.Send(sel)
}
// Disables the   method for the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/disableFlushing()
func (w_ Window) DisableFlushWindow() {
	sel := objc.RegisterName("disableFlushWindow")
	w_.ID.Send(sel)
}
// Disables the default button cell’s key equivalent, so it doesn’t perform a click when the user presses Return (or Enter). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/disableKeyEquivalentForDefaultButtonCell()
func (w_ Window) DisableKeyEquivalentForDefaultButtonCell() {
	sel := objc.RegisterName("disableKeyEquivalentForDefaultButtonCell")
	w_.ID.Send(sel)
}
// Disables the window’s screen updates until the window is flushed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/disableScreenUpdatesUntilFlush()
func (w_ Window) DisableScreenUpdatesUntilFlush() {
	sel := objc.RegisterName("disableScreenUpdatesUntilFlush")
	w_.ID.Send(sel)
}
// Disables snapshot restoration. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/disableSnapshotRestoration()
func (w_ Window) DisableSnapshotRestoration() {
	sel := objc.RegisterName("disableSnapshotRestoration")
	w_.ID.Send(sel)
}
// Discards all of the window’s cached image rectangles. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/discardCachedImage()
func (w_ Window) DiscardCachedImage() {
	sel := objc.RegisterName("discardCachedImage")
	w_.ID.Send(sel)
}
// Invalidates all cursor rectangles in the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/discardCursorRects()
func (w_ Window) DiscardCursorRects() {
	sel := objc.RegisterName("discardCursorRects")
	w_.ID.Send(sel)
}
// Forwards the message to the global application object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/discardEvents(matching:before:)
func (w_ Window) DiscardEventsMatchingMaskBeforeEvent(mask unsafe.Pointer, lastEvent unsafe.Pointer) {
	sel := objc.RegisterName("discardEventsMatchingMask:beforeEvent:")
	w_.ID.Send(sel, mask, lastEvent)
}
// Passes a display message down the window’s view hierarchy, thus redrawing all views within the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/display()
func (w_ Window) Display() {
	sel := objc.RegisterName("display")
	w_.ID.Send(sel)
}
// Passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/displayIfNeeded()
func (w_ Window) DisplayIfNeeded() {
	sel := objc.RegisterName("displayIfNeeded")
	w_.ID.Send(sel)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/displayLink(target:selector:)
func (w_ Window) DisplayLinkWithTargetSelector(target objc.ID, selector objc.SEL) unsafe.Pointer {
	sel := objc.RegisterName("displayLinkWithTarget:selector:")
	ret := w_.ID.Send(sel, target, selector)
	return unsafe.Pointer(ret)
}
// Begins a dragging session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/drag(_:at:offset:event:pasteboard:source:slideBack:)
func (w_ Window) DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, baseLocation unsafe.Pointer, initialOffset unsafe.Pointer, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool) {
	sel := objc.RegisterName("dragImage:at:offset:event:pasteboard:source:slideBack:")
	w_.ID.Send(sel, image, baseLocation, initialOffset, event, pboard, sourceObj, slideFlag)
}
// Reenables cursor rectangle management within the window after a   message. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/enableCursorRects()
func (w_ Window) EnableCursorRects() {
	sel := objc.RegisterName("enableCursorRects")
	w_.ID.Send(sel)
}
// Reenables the   method for the window after it was disabled through a previous   message. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/enableFlushing()
func (w_ Window) EnableFlushWindow() {
	sel := objc.RegisterName("enableFlushWindow")
	w_.ID.Send(sel)
}
// Reenables the default button cell’s key equivalent, so it performs a click when the user presses Return (or Enter). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/enableKeyEquivalentForDefaultButtonCell()
func (w_ Window) EnableKeyEquivalentForDefaultButtonCell() {
	sel := objc.RegisterName("enableKeyEquivalentForDefaultButtonCell")
	w_.ID.Send(sel)
}
// Enables snapshot restoration. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/enableSnapshotRestoration()
func (w_ Window) EnableSnapshotRestoration() {
	sel := objc.RegisterName("enableSnapshotRestoration")
	w_.ID.Send(sel)
}
// Forces the field editor to give up its first responder status and prepares it for its next assignment. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/endEditing(for:)
func (w_ Window) EndEditingFor(object objc.ID) {
	sel := objc.RegisterName("endEditingFor:")
	w_.ID.Send(sel, object)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/endSheet(_:)-4dmmq
func (w_ Window) EndSheet(sheetWindow unsafe.Pointer) {
	sel := objc.RegisterName("endSheet:")
	w_.ID.Send(sel, sheetWindow)
}
// Ends a document-modal session and dismisses the specified sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/endSheet(_:returnCode:)
func (w_ Window) EndSheetReturnCode(sheetWindow unsafe.Pointer, returnCode unsafe.Pointer) {
	sel := objc.RegisterName("endSheet:returnCode:")
	w_.ID.Send(sel, sheetWindow, returnCode)
}
// Returns the window’s field editor, creating it if requested. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/fieldEditor(_:for:)
func (w_ Window) FieldEditorForObject(createFlag bool, object objc.ID) unsafe.Pointer {
	sel := objc.RegisterName("fieldEditor:forObject:")
	ret := w_.ID.Send(sel, createFlag, object)
	return unsafe.Pointer(ret)
}
// Flushes the window’s offscreen buffer to the screen if the window is buffered and flushing is enabled. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/flush()
func (w_ Window) FlushWindow() {
	sel := objc.RegisterName("flushWindow")
	w_.ID.Send(sel)
}
// Flushes the window’s offscreen buffer to the screen if flushing is enabled and if the last   message had no effect because flushing was disabled. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/flushIfNeeded()
func (w_ Window) FlushWindowIfNeeded() {
	sel := objc.RegisterName("flushWindowIfNeeded")
	w_.ID.Send(sel)
}
// Returns the window’s frame rectangle with a given content rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/frameRect(forContentRect:)
func (w_ Window) FrameRectForContentRect(contentRect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("frameRectForContentRect:")
	ret := w_.ID.Send(sel, contentRect)
	return unsafe.Pointer(ret)
}
// Returns the window’s graphics state object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/gState()
func (w_ Window) GState() int {
	sel := objc.RegisterName("gState")
	ret := w_.ID.Send(sel)
	return int(ret)
}
// Handles the AppleScript command to close the window (and its associated document, if any). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/handleClose(_:)
func (w_ Window) HandleCloseScriptCommand(command unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("handleCloseScriptCommand:")
	ret := w_.ID.Send(sel, command)
	return ret
}
// Handles the AppleScript command to print the contents of the window (or its associated document, if any). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/handlePrint(_:)
func (w_ Window) HandlePrintScriptCommand(command unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("handlePrintScriptCommand:")
	ret := w_.ID.Send(sel, command)
	return ret
}
// Handles the AppleScript command to save the window (and its associated document, if any). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/handleSave(_:)
func (w_ Window) HandleSaveScriptCommand(command unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("handleSaveScriptCommand:")
	ret := w_.ID.Send(sel, command)
	return ret
}
// Inserts the view controller into the window’s array of title bar accessory view controllers at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/insertTitlebarAccessoryViewController(_:at:)
func (w_ Window) InsertTitlebarAccessoryViewControllerAtIndex(childViewController unsafe.Pointer, index int) {
	sel := objc.RegisterName("insertTitlebarAccessoryViewController:atIndex:")
	w_.ID.Send(sel, childViewController, index)
}
// Marks as invalid the cursor rectangles of a given view object in the window, so they’ll be set up again when the window becomes key. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/invalidateCursorRects(for:)
func (w_ Window) InvalidateCursorRectsForView(view unsafe.Pointer) {
	sel := objc.RegisterName("invalidateCursorRectsForView:")
	w_.ID.Send(sel, view)
}
// Invalidates the window shadow so that it is recomputed based on the current window shape. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/invalidateShadow()
func (w_ Window) InvalidateShadow() {
	sel := objc.RegisterName("invalidateShadow")
	w_.ID.Send(sel)
}
// Updates the layout of views in the window based on the current views and constraints. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/layoutIfNeeded()
func (w_ Window) LayoutIfNeeded() {
	sel := objc.RegisterName("layoutIfNeeded")
	w_.ID.Send(sel)
}
// Attempts to make a given responder the first responder for the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/makeFirstResponder(_:)
func (w_ Window) MakeFirstResponder(responder unsafe.Pointer) bool {
	sel := objc.RegisterName("makeFirstResponder:")
	ret := w_.ID.Send(sel, responder)
	return ret != 0
}
// Makes the window the key window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/makeKey()
func (w_ Window) MakeKeyWindow() {
	sel := objc.RegisterName("makeKeyWindow")
	w_.ID.Send(sel)
}
// Moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/makeKeyAndOrderFront(_:)
func (w_ Window) MakeKeyAndOrderFront(sender objc.ID) {
	sel := objc.RegisterName("makeKeyAndOrderFront:")
	w_.ID.Send(sel, sender)
}
// Makes the window the main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/makeMain()
func (w_ Window) MakeMainWindow() {
	sel := objc.RegisterName("makeMainWindow")
	w_.ID.Send(sel)
}
// Merges all open windows into a single tabbed window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/mergeAllWindows(_:)
func (w_ Window) MergeAllWindows(sender objc.ID) {
	sel := objc.RegisterName("mergeAllWindows:")
	w_.ID.Send(sel, sender)
}
// Removes the window from the screen list and displays the minimized window in the Dock. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/miniaturize(_:)
func (w_ Window) Miniaturize(sender objc.ID) {
	sel := objc.RegisterName("miniaturize:")
	w_.ID.Send(sel, sender)
}
// Moves the tab to a new containing window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/moveTabToNewWindow(_:)
func (w_ Window) MoveTabToNewWindow(sender objc.ID) {
	sel := objc.RegisterName("moveTabToNewWindow:")
	w_.ID.Send(sel, sender)
}
// Returns the next event matching a given mask. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/nextEvent(matching:)
func (w_ Window) NextEventMatchingMask(mask unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("nextEventMatchingMask:")
	ret := w_.ID.Send(sel, mask)
	return unsafe.Pointer(ret)
}
// Forwards the message to the global application object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/nextEvent(matching:until:inMode:dequeue:)
func (w_ Window) NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer {
	sel := objc.RegisterName("nextEventMatchingMask:untilDate:inMode:dequeue:")
	ret := w_.ID.Send(sel, mask, expiration, mode, deqFlag)
	return unsafe.Pointer(ret)
}
// Repositions the window’s window device in the window server’s screen list. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/order(_:relativeTo:)
func (w_ Window) OrderWindowRelativeTo(place WindowOrderingMode, otherWin int) {
	sel := objc.RegisterName("orderWindow:relativeTo:")
	w_.ID.Send(sel, place, otherWin)
}
// Moves the window to the back of its level in the screen list, without changing either the key window or the main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/orderBack(_:)
func (w_ Window) OrderBack(sender objc.ID) {
	sel := objc.RegisterName("orderBack:")
	w_.ID.Send(sel, sender)
}
// Moves the window to the front of its level in the screen list, without changing either the key window or the main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/orderFront(_:)
func (w_ Window) OrderFront(sender objc.ID) {
	sel := objc.RegisterName("orderFront:")
	w_.ID.Send(sel, sender)
}
// Moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/orderFrontRegardless()
func (w_ Window) OrderFrontRegardless() {
	sel := objc.RegisterName("orderFrontRegardless")
	w_.ID.Send(sel)
}
// Removes the window from the screen list, which hides the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/orderOut(_:)
func (w_ Window) OrderOut(sender objc.ID) {
	sel := objc.RegisterName("orderOut:")
	w_.ID.Send(sel, sender)
}
// Simulates the user clicking the close button by momentarily highlighting the button and then closing the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/performClose(_:)
func (w_ Window) PerformClose(sender objc.ID) {
	sel := objc.RegisterName("performClose:")
	w_.ID.Send(sel, sender)
}
// Starts a window drag based on the specified mouse-down event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/performDrag(with:)
func (w_ Window) PerformWindowDragWithEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("performWindowDragWithEvent:")
	w_.ID.Send(sel, event)
}
// Simulates the user clicking the minimize button by momentarily highlighting the button, then minimizing the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/performMiniaturize(_:)
func (w_ Window) PerformMiniaturize(sender objc.ID) {
	sel := objc.RegisterName("performMiniaturize:")
	w_.ID.Send(sel, sender)
}
// This action method simulates the user clicking the zoom box by momentarily highlighting the button and then zooming the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/performZoom(_:)
func (w_ Window) PerformZoom(sender objc.ID) {
	sel := objc.RegisterName("performZoom:")
	w_.ID.Send(sel, sender)
}
// Forwards the message to the global application object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/postEvent(_:atStart:)
func (w_ Window) PostEventAtStart(event unsafe.Pointer, flag bool) {
	sel := objc.RegisterName("postEvent:atStart:")
	w_.ID.Send(sel, event, flag)
}
// Runs the Print panel, and if the user chooses an option other than canceling, prints the window (its frame view and all subviews). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/printWindow(_:)
func (w_ Window) Print(sender objc.ID) {
	sel := objc.RegisterName("print:")
	w_.ID.Send(sel, sender)
}
// Marks the key view loop as “dirty” and in need of recalculation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/recalculateKeyViewLoop()
func (w_ Window) RecalculateKeyViewLoop() {
	sel := objc.RegisterName("recalculateKeyViewLoop")
	w_.ID.Send(sel)
}
// Registers a set of pasteboard types that the window accepts as the destination of an image-dragging session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/registerForDraggedTypes(_:)
func (w_ Window) RegisterForDraggedTypes(newTypes unsafe.Pointer) {
	sel := objc.RegisterName("registerForDraggedTypes:")
	w_.ID.Send(sel, newTypes)
}
// Detaches a given child window from the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/removeChildWindow(_:)
func (w_ Window) RemoveChildWindow(childWin unsafe.Pointer) {
	sel := objc.RegisterName("removeChildWindow:")
	w_.ID.Send(sel, childWin)
}
// Removes the view controller at the specified index from the window’s array of title bar accessory view controllers. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/removeTitlebarAccessoryViewController(at:)
func (w_ Window) RemoveTitlebarAccessoryViewControllerAtIndex(index int) {
	sel := objc.RegisterName("removeTitlebarAccessoryViewControllerAtIndex:")
	w_.ID.Send(sel, index)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/requestSharingOfWindow(_:completionHandler:)
func (w_ Window) RequestSharingOfWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("requestSharingOfWindow:completionHandler:")
	w_.ID.Send(sel, window, completionHandler)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/requestSharingOfWindow(usingPreview:title:completionHandler:)
func (w_ Window) RequestSharingOfWindowUsingPreviewTitleCompletionHandler(image unsafe.Pointer, title string, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("requestSharingOfWindowUsingPreview:title:completionHandler:")
	w_.ID.Send(sel, image, title, completionHandler)
}
// Clears the window’s cursor rectangles and the cursor rectangles of the   objects in its view hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/resetCursorRects()
func (w_ Window) ResetCursorRects() {
	sel := objc.RegisterName("resetCursorRects")
	w_.ID.Send(sel)
}
// Resigns the window’s key window status. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/resignKey()
func (w_ Window) ResignKeyWindow() {
	sel := objc.RegisterName("resignKeyWindow")
	w_.ID.Send(sel)
}
// Resigns the window’s main window status. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/resignMain()
func (w_ Window) ResignMainWindow() {
	sel := objc.RegisterName("resignMainWindow")
	w_.ID.Send(sel)
}
// Splices the window’s cached image rectangles, if any, back into its raster image (and buffer if it has one), undoing the effect of any drawing performed within those areas since they were established using  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/restoreCachedImage()
func (w_ Window) RestoreCachedImage() {
	sel := objc.RegisterName("restoreCachedImage")
	w_.ID.Send(sel)
}
// Presents the toolbar customization user interface. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/runToolbarCustomizationPalette(_:)
func (w_ Window) RunToolbarCustomizationPalette(sender objc.ID) {
	sel := objc.RegisterName("runToolbarCustomizationPalette:")
	w_.ID.Send(sel, sender)
}
// Saves the window’s frame rectangle in the user defaults system under a given name. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/saveFrame(usingName:)
func (w_ Window) SaveFrameUsingName(name unsafe.Pointer) {
	sel := objc.RegisterName("saveFrameUsingName:")
	w_.ID.Send(sel, name)
}
// Gives key view status to the view that follows the given view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectKeyView(following:)
func (w_ Window) SelectKeyViewFollowingView(view unsafe.Pointer) {
	sel := objc.RegisterName("selectKeyViewFollowingView:")
	w_.ID.Send(sel, view)
}
// Gives key view status to the view that precedes the given view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectKeyView(preceding:)
func (w_ Window) SelectKeyViewPrecedingView(view unsafe.Pointer) {
	sel := objc.RegisterName("selectKeyViewPrecedingView:")
	w_.ID.Send(sel, view)
}
// Searches for a candidate next key view and, if it finds one, tries to make it the first responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectNextKeyView(_:)
func (w_ Window) SelectNextKeyView(sender objc.ID) {
	sel := objc.RegisterName("selectNextKeyView:")
	w_.ID.Send(sel, sender)
}
// Selects the next tab in the tab group in the trailing direction. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectNextTab(_:)
func (w_ Window) SelectNextTab(sender objc.ID) {
	sel := objc.RegisterName("selectNextTab:")
	w_.ID.Send(sel, sender)
}
// Searches for a candidate previous key view and, if it finds one, tries to make it the first responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectPreviousKeyView(_:)
func (w_ Window) SelectPreviousKeyView(sender objc.ID) {
	sel := objc.RegisterName("selectPreviousKeyView:")
	w_.ID.Send(sel, sender)
}
// Selects the previous tab in the tab group in the leading direction. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectPreviousTab(_:)
func (w_ Window) SelectPreviousTab(sender objc.ID) {
	sel := objc.RegisterName("selectPreviousTab:")
	w_.ID.Send(sel, sender)
}
// This action method dispatches mouse and keyboard events the global application object sends to the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/sendEvent(_:)
func (w_ Window) SendEvent(event unsafe.Pointer) {
	sel := objc.RegisterName("sendEvent:")
	w_.ID.Send(sel, event)
}
// Sets the part of the window that stays stationary during constraint-based layout. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setAnchorAttribute(_:for:)
func (w_ Window) SetAnchorAttributeForOrientation(attr unsafe.Pointer, orientation unsafe.Pointer) {
	sel := objc.RegisterName("setAnchorAttribute:forOrientation:")
	w_.ID.Send(sel, attr, orientation)
}
// Specifies whether the window calculates the thickness of a given border automatically. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setAutorecalculatesContentBorderThickness(_:for:)
func (w_ Window) SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge int) {
	sel := objc.RegisterName("setAutorecalculatesContentBorderThickness:forEdge:")
	w_.ID.Send(sel, flag, edge)
}
// Specifies the thickness of a given border of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setContentBorderThickness(_:for:)
func (w_ Window) SetContentBorderThicknessForEdge(thickness float64, edge int) {
	sel := objc.RegisterName("setContentBorderThickness:forEdge:")
	w_.ID.Send(sel, thickness, edge)
}
// Sets the size of the window’s content view to a given size, which is expressed in the window’s base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setContentSize(_:)
func (w_ Window) SetContentSize(size unsafe.Pointer) {
	sel := objc.RegisterName("setContentSize:")
	w_.ID.Send(sel, size)
}
// Sets a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setDynamicDepthLimit(_:)
func (w_ Window) SetDynamicDepthLimit(flag bool) {
	sel := objc.RegisterName("setDynamicDepthLimit:")
	w_.ID.Send(sel, flag)
}
// Sets the origin and size of the window’s frame rectangle according to a given frame rectangle, thereby setting its position and size onscreen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrame(_:display:)
func (w_ Window) SetFrameDisplay(frameRect unsafe.Pointer, flag bool) {
	sel := objc.RegisterName("setFrame:display:")
	w_.ID.Send(sel, frameRect, flag)
}
// Sets the origin and size of the window’s frame rectangle, with optional animation, according to a given frame rectangle, thereby setting its position and size onscreen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrame(_:display:animate:)
func (w_ Window) SetFrameDisplayAnimate(frameRect unsafe.Pointer, displayFlag bool, animateFlag bool) {
	sel := objc.RegisterName("setFrame:display:animate:")
	w_.ID.Send(sel, frameRect, displayFlag, animateFlag)
}
// Sets the window’s frame rectangle from a given string representation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrame(from:)
func (w_ Window) SetFrameFromString(string unsafe.Pointer) {
	sel := objc.RegisterName("setFrameFromString:")
	w_.ID.Send(sel, string)
}
// Sets the name AppKit uses to automatically save the window’s frame rectangle data in the defaults system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrameAutosaveName(_:)
func (w_ Window) SetFrameAutosaveName(name unsafe.Pointer) bool {
	sel := objc.RegisterName("setFrameAutosaveName:")
	ret := w_.ID.Send(sel, name)
	return ret != 0
}
// Positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrameOrigin(_:)
func (w_ Window) SetFrameOrigin(point unsafe.Pointer) {
	sel := objc.RegisterName("setFrameOrigin:")
	w_.ID.Send(sel, point)
}
// Positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrameTopLeftPoint(_:)
func (w_ Window) SetFrameTopLeftPoint(point unsafe.Pointer) {
	sel := objc.RegisterName("setFrameTopLeftPoint:")
	w_.ID.Send(sel, point)
}
// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrameUsingName(_:)
func (w_ Window) SetFrameUsingName(name unsafe.Pointer) bool {
	sel := objc.RegisterName("setFrameUsingName:")
	ret := w_.ID.Send(sel, name)
	return ret != 0
}
// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system. Can operate on non-resizable windows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrameUsingName(_:force:)
func (w_ Window) SetFrameUsingNameForce(name unsafe.Pointer, force bool) bool {
	sel := objc.RegisterName("setFrameUsingName:force:")
	ret := w_.ID.Send(sel, name, force)
	return ret != 0
}
// Sets the window’s miniaturized state to the value you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setIsMiniaturized(_:)
func (w_ Window) SetIsMiniaturized(flag bool) {
	sel := objc.RegisterName("setIsMiniaturized:")
	w_.ID.Send(sel, flag)
}
// Sets the window’s visible state to the value you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setIsVisible(_:)
func (w_ Window) SetIsVisible(flag bool) {
	sel := objc.RegisterName("setIsVisible:")
	w_.ID.Send(sel, flag)
}
// Sets the window’s zoomed state to the value you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setIsZoomed(_:)
func (w_ Window) SetIsZoomed(flag bool) {
	sel := objc.RegisterName("setIsZoomed:")
	w_.ID.Send(sel, flag)
}
// Sets a given path as the window’s title, formatting it as a file-system path, and records this path as the window’s associated file. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setTitleWithRepresentedFilename(_:)
func (w_ Window) SetTitleWithRepresentedFilename(filename string) {
	sel := objc.RegisterName("setTitleWithRepresentedFilename:")
	w_.ID.Send(sel, filename)
}
// Returns the window button of a given window button kind in the window’s view hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/standardWindowButton(_:)
func (w_ Window) StandardWindowButton(b unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("standardWindowButton:")
	ret := w_.ID.Send(sel, b)
	return unsafe.Pointer(ret)
}
// Takes the window into or out of fullscreen mode, [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toggleFullScreen(_:)
func (w_ Window) ToggleFullScreen(sender objc.ID) {
	sel := objc.RegisterName("toggleFullScreen:")
	w_.ID.Send(sel, sender)
}
// Shows or hides the tab bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toggleTabBar(_:)
func (w_ Window) ToggleTabBar(sender objc.ID) {
	sel := objc.RegisterName("toggleTabBar:")
	w_.ID.Send(sel, sender)
}
// Shows or hides the tab overview. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toggleTabOverview(_:)
func (w_ Window) ToggleTabOverview(sender objc.ID) {
	sel := objc.RegisterName("toggleTabOverview:")
	w_.ID.Send(sel, sender)
}
// Toggles the visibility of the window’s toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toggleToolbarShown(_:)
func (w_ Window) ToggleToolbarShown(sender objc.ID) {
	sel := objc.RegisterName("toggleToolbarShown:")
	w_.ID.Send(sel, sender)
}
// Tracks events that match the specified mask using the specified tracking handler until the tracking handler explicitly terminates tracking. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/trackEvents(matching:timeout:mode:handler:)
func (w_ Window) TrackEventsMatchingMaskTimeoutModeHandler(mask unsafe.Pointer, timeout float64, mode unsafe.Pointer, trackingHandler unsafe.Pointer) {
	sel := objc.RegisterName("trackEventsMatchingMask:timeout:mode:handler:")
	w_.ID.Send(sel, mask, timeout, mode, trackingHandler)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/transferWindowSharing(to:completionHandler:)
func (w_ Window) TransferWindowSharingToWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer) {
	sel := objc.RegisterName("transferWindowSharingToWindow:completionHandler:")
	w_.ID.Send(sel, window, completionHandler)
}
// Dispatches action messages with a given argument. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/tryToPerform(_:with:)
func (w_ Window) TryToPerformWith(action objc.SEL, object objc.ID) bool {
	sel := objc.RegisterName("tryToPerform:with:")
	ret := w_.ID.Send(sel, action, object)
	return ret != 0
}
// Unregisters the window as a possible destination for dragging operations. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/unregisterDraggedTypes()
func (w_ Window) UnregisterDraggedTypes() {
	sel := objc.RegisterName("unregisterDraggedTypes")
	w_.ID.Send(sel)
}
// Updates the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/update()
func (w_ Window) Update() {
	sel := objc.RegisterName("update")
	w_.ID.Send(sel)
}
// Updates the constraints based on changes to views in the window since the last layout. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/updateConstraintsIfNeeded()
func (w_ Window) UpdateConstraintsIfNeeded() {
	sel := objc.RegisterName("updateConstraintsIfNeeded")
	w_.ID.Send(sel)
}
// Specifies whether the window is to optimize focusing and drawing when displaying its views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/useOptimizedDrawing(_:)
func (w_ Window) UseOptimizedDrawing(flag bool) {
	sel := objc.RegisterName("useOptimizedDrawing:")
	w_.ID.Send(sel, flag)
}
// Returns the scale factor applied to the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/userSpaceScaleFactor
func (w_ Window) UserSpaceScaleFactor() float64 {
	sel := objc.RegisterName("userSpaceScaleFactor")
	ret := w_.ID.Send(sel)
	return float64(ret)
}
// Searches for an object that responds to a Services request. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/validRequestor(forSendType:returnType:)
func (w_ Window) ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID {
	sel := objc.RegisterName("validRequestorForSendType:returnType:")
	ret := w_.ID.Send(sel, sendType, returnType)
	return ret
}
// Displays a visual representation of the supplied constraints in the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/visualizeConstraints(_:)
func (w_ Window) VisualizeConstraints(constraints unsafe.Pointer) {
	sel := objc.RegisterName("visualizeConstraints:")
	w_.ID.Send(sel, constraints)
}
// Toggles the size and location of the window between its standard state (which the application provides as the best size to display the window’s data) and its user state (a new size and location the user may have set by moving or resizing the window). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/zoom(_:)
func (w_ Window) Zoom(sender objc.ID) {
	sel := objc.RegisterName("zoom:")
	w_.ID.Send(sel, sender)
}

