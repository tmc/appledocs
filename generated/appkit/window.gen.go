// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Window] class.
var (
	windowClass     _WindowClass
	windowClassOnce sync.Once
)

func getWindowClass() _WindowClass {
	windowClassOnce.Do(func() {
		windowClass = _WindowClass{objc.GetClass("NSWindow")}
	})
	return windowClass
}

type _WindowClass struct {
	class objc.Class
}

// An interface definition for the [Window] class.
type IWindow interface {
	IResponder
	AddChildWindowOrdered(childWin unsafe.Pointer, place WindowOrderingMode)
	AddTabbedWindowOrdered(window unsafe.Pointer, ordered WindowOrderingMode)
	AddTitlebarAccessoryViewController(childViewController unsafe.Pointer)
	AnchorAttributeForOrientation(orientation unsafe.Pointer) unsafe.Pointer
	AnimationResizeTime(newFrame unsafe.Pointer) float64
	AutorecalculatesContentBorderThicknessForEdge(edge int) bool
	BackingAlignedRectOptions(rect unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer
	BecomeKeyWindow()
	BecomeMainWindow()
	BeginCriticalSheetCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer)
	BeginDraggingSessionWithItemsEventSource(items unsafe.Pointer, event unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer
	BeginSheetCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer)
	CacheImageInRect(rect unsafe.Pointer)
	CanRepresentDisplayGamut(displayGamut unsafe.Pointer) bool
	CanStoreColor() bool
	CascadeTopLeftFromPoint(topLeftPoint unsafe.Pointer) unsafe.Pointer
	Center()
	Close()
	ConstrainFrameRectToScreen(frameRect unsafe.Pointer, screen unsafe.Pointer) unsafe.Pointer
	ContentBorderThicknessForEdge(edge int) float64
	ContentRectForFrameRect(frameRect unsafe.Pointer) unsafe.Pointer
	ConvertBaseToScreen(point unsafe.Pointer) unsafe.Pointer
	ConvertRectFromBacking(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectFromScreen(rect unsafe.Pointer) unsafe.Pointer
	ConvertPointFromScreen(point unsafe.Pointer) unsafe.Pointer
	ConvertPointToScreen(point unsafe.Pointer) unsafe.Pointer
	ConvertPointFromBacking(point unsafe.Pointer) unsafe.Pointer
	ConvertPointToBacking(point unsafe.Pointer) unsafe.Pointer
	ConvertScreenToBase(point unsafe.Pointer) unsafe.Pointer
	ConvertRectToBacking(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectToScreen(rect unsafe.Pointer) unsafe.Pointer
	DataWithEPSInsideRect(rect unsafe.Pointer) unsafe.Pointer
	DataWithPDFInsideRect(rect unsafe.Pointer) unsafe.Pointer
	Deminiaturize(sender objc.ID)
	DisableCursorRects()
	DisableFlushWindow()
	DisableKeyEquivalentForDefaultButtonCell()
	DisableScreenUpdatesUntilFlush()
	DisableSnapshotRestoration()
	DiscardCachedImage()
	DiscardCursorRects()
	DiscardEventsMatchingMaskBeforeEvent(mask unsafe.Pointer, lastEvent unsafe.Pointer)
	Display()
	DisplayIfNeeded()
	DisplayLinkWithTargetSelector(target objc.ID, selector objc.SEL) unsafe.Pointer
	DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, baseLocation unsafe.Pointer, initialOffset unsafe.Pointer, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool)
	EnableCursorRects()
	EnableFlushWindow()
	EnableKeyEquivalentForDefaultButtonCell()
	EnableSnapshotRestoration()
	EndEditingFor(object objc.ID)
	EndSheet(sheetWindow unsafe.Pointer)
	EndSheetReturnCode(sheetWindow unsafe.Pointer, returnCode unsafe.Pointer)
	FieldEditorForObject(createFlag bool, object objc.ID) unsafe.Pointer
	FlushWindow()
	FlushWindowIfNeeded()
	FrameRectForContentRect(contentRect unsafe.Pointer) unsafe.Pointer
	GState() int
	HandleCloseScriptCommand(command unsafe.Pointer) objc.ID
	HandlePrintScriptCommand(command unsafe.Pointer) objc.ID
	HandleSaveScriptCommand(command unsafe.Pointer) objc.ID
	InsertTitlebarAccessoryViewControllerAtIndex(childViewController unsafe.Pointer, index int)
	InvalidateCursorRectsForView(view unsafe.Pointer)
	InvalidateShadow()
	LayoutIfNeeded()
	MakeFirstResponder(responder unsafe.Pointer) bool
	MakeKeyWindow()
	MakeKeyAndOrderFront(sender objc.ID)
	MakeMainWindow()
	MergeAllWindows(sender objc.ID)
	Miniaturize(sender objc.ID)
	MoveTabToNewWindow(sender objc.ID)
	NextEventMatchingMask(mask unsafe.Pointer) unsafe.Pointer
	NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer
	OrderWindowRelativeTo(place WindowOrderingMode, otherWin int)
	OrderBack(sender objc.ID)
	OrderFront(sender objc.ID)
	OrderFrontRegardless()
	OrderOut(sender objc.ID)
	PerformClose(sender objc.ID)
	PerformWindowDragWithEvent(event unsafe.Pointer)
	PerformMiniaturize(sender objc.ID)
	PerformZoom(sender objc.ID)
	PostEventAtStart(event unsafe.Pointer, flag bool)
	Print(sender objc.ID)
	RecalculateKeyViewLoop()
	RegisterForDraggedTypes(newTypes unsafe.Pointer)
	RemoveChildWindow(childWin unsafe.Pointer)
	RemoveTitlebarAccessoryViewControllerAtIndex(index int)
	RequestSharingOfWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer)
	RequestSharingOfWindowUsingPreviewTitleCompletionHandler(image unsafe.Pointer, title string, completionHandler unsafe.Pointer)
	ResetCursorRects()
	ResignKeyWindow()
	ResignMainWindow()
	RestoreCachedImage()
	RunToolbarCustomizationPalette(sender objc.ID)
	SaveFrameUsingName(name unsafe.Pointer)
	SelectKeyViewFollowingView(view unsafe.Pointer)
	SelectKeyViewPrecedingView(view unsafe.Pointer)
	SelectNextKeyView(sender objc.ID)
	SelectNextTab(sender objc.ID)
	SelectPreviousKeyView(sender objc.ID)
	SelectPreviousTab(sender objc.ID)
	SendEvent(event unsafe.Pointer)
	SetAnchorAttributeForOrientation(attr unsafe.Pointer, orientation unsafe.Pointer)
	SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge int)
	SetContentBorderThicknessForEdge(thickness float64, edge int)
	SetContentSize(size unsafe.Pointer)
	SetDynamicDepthLimit(flag bool)
	SetFrameDisplay(frameRect unsafe.Pointer, flag bool)
	SetFrameDisplayAnimate(frameRect unsafe.Pointer, displayFlag bool, animateFlag bool)
	SetFrameFromString(string unsafe.Pointer)
	SetFrameAutosaveName(name unsafe.Pointer) bool
	SetFrameOrigin(point unsafe.Pointer)
	SetFrameTopLeftPoint(point unsafe.Pointer)
	SetFrameUsingName(name unsafe.Pointer) bool
	SetFrameUsingNameForce(name unsafe.Pointer, force bool) bool
	SetIsMiniaturized(flag bool)
	SetIsVisible(flag bool)
	SetIsZoomed(flag bool)
	SetTitleWithRepresentedFilename(filename string)
	StandardWindowButton(b unsafe.Pointer) unsafe.Pointer
	ToggleFullScreen(sender objc.ID)
	ToggleTabBar(sender objc.ID)
	ToggleTabOverview(sender objc.ID)
	ToggleToolbarShown(sender objc.ID)
	TrackEventsMatchingMaskTimeoutModeHandler(mask unsafe.Pointer, timeout float64, mode unsafe.Pointer, trackingHandler unsafe.Pointer)
	TransferWindowSharingToWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer)
	TryToPerformWith(action objc.SEL, object objc.ID) bool
	UnregisterDraggedTypes()
	Update()
	UpdateConstraintsIfNeeded()
	UseOptimizedDrawing(flag bool)
	UserSpaceScaleFactor() float64
	ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID
	VisualizeConstraints(constraints unsafe.Pointer)
	Zoom(sender objc.ID)
}

// A window that an app displays on the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow
type Window struct {
	Responder
}

// WindowFrom constructs a [Window] from an unsafe.Pointer.
//
// A window that an app displays on the screen.
func WindowFrom(ptr unsafe.Pointer) Window {
	return Window{
		Responder: ResponderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (wc _WindowClass) Alloc() Window {
	rv := objc.Send[Window](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WindowClass) New() Window {
	rv := objc.Send[Window](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ Window) Init() Window {
	rv := objc.Send[Window](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ Window) Autorelease() Window {
	rv := objc.Send[Window](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWindow creates a new Window instance.
func NewWindow() Window {
	return getWindowClass().New()
}


// Initializes the window with the specified values.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:)
func NewWindowWithContentRectStyleMaskBackingDefer(contentRect unsafe.Pointer, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	instance := getWindowClass().Alloc()
	rv := objc.Send[Window](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	rv.Autorelease()
	return rv
}

// Initializes an allocated window with the specified values.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:screen:)
func NewWindowWithContentRectStyleMaskBackingDeferScreen(contentRect unsafe.Pointer, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen unsafe.Pointer) Window {
	instance := getWindowClass().Alloc()
	rv := objc.Send[Window](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	rv.Autorelease()
	return rv
}

// Creates a titled window that contains the specified content view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentViewController:)
func NewWindowWithContentViewController(contentViewController unsafe.Pointer) Window {
	rv := objc.Send[Window](objc.ID(getWindowClass().class), objc.Sel("windowWithContentViewController:"), contentViewController)
	return rv
}

// Returns a Cocoa window created from a Carbon window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(windowRef:)
func NewWindowWithWindowRef(windowRef unsafe.Pointer) Window {
	instance := getWindowClass().Alloc()
	rv := objc.Send[Window](instance.ID, objc.Sel("initWithWindowRef:"), windowRef)
	rv.Autorelease()
	return rv
}


// Returns the content rectangle used by a window with a given frame rectangle and window style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentRect(forFrameRect:styleMask:)
func (wc _WindowClass) ContentRectForFrameRectStyleMask(fRect unsafe.Pointer, style WindowStyleMask) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("contentRectForFrameRect:styleMask:"), fRect, style)
	return rv
}

// Returns the frame rectangle used by a window with a given content rectangle and window style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frameRect(forContentRect:styleMask:)
func (wc _WindowClass) FrameRectForContentRectStyleMask(cRect unsafe.Pointer, style WindowStyleMask) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("frameRectForContentRect:styleMask:"), cRect, style)
	return rv
}

// Creates a titled window that contains the specified content view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentViewController:)
func (wc _WindowClass) WindowWithContentViewController(contentViewController unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("windowWithContentViewController:"), contentViewController)
	return rv
}

// This method does nothing; it is here for backward compatibility.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/menuChanged(_:)
func (wc _WindowClass) MenuChanged(menu unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("menuChanged:"), menu)
}

// Returns the minimum width a window’s frame rectangle must have for it to display a title, with a given window style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/minFrameWidth(withTitle:styleMask:)
func (wc _WindowClass) MinFrameWidthWithTitleStyleMask(title string, style WindowStyleMask) float64 {
	rv := objc.Send[float64](objc.ID(wc.class), objc.Sel("minFrameWidthWithTitle:styleMask:"), objc.String(title), style)
	return rv
}

// Removes the frame data stored under a given name from the application’s user defaults.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/removeFrame(usingName:)
func (wc _WindowClass) RemoveFrameUsingName(name unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("removeFrameUsingName:"), name)
}

// Returns a new instance of a given standard window button, sized appropriately for a given window style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/standardWindowButton(_:for:)
func (wc _WindowClass) StandardWindowButtonForStyleMask(b unsafe.Pointer, styleMask WindowStyleMask) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("standardWindowButton:forStyleMask:"), b, styleMask)
	return rv
}

// Returns the number of the frontmost window that would be hit by a mouse-down at the specified screen location.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumber(at:belowWindowWithWindowNumber:)
func (wc _WindowClass) WindowNumberAtPointBelowWindowWithWindowNumber(point unsafe.Pointer, windowNumber int) int {
	rv := objc.Send[int](objc.ID(wc.class), objc.Sel("windowNumberAtPoint:belowWindowWithWindowNumber:"), point, windowNumber)
	return rv
}

// Returns the window numbers for all visible windows satisfying the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumbers(options:)
func (wc _WindowClass) WindowNumbersWithOptions(options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("windowNumbersWithOptions:"), options)
	return rv
}

// Adds a given window as a child window of the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/addChildWindow(_:ordered:)
func (w_ Window) AddChildWindowOrdered(childWin unsafe.Pointer, place WindowOrderingMode) {
	objc.Send[objc.ID](w_.ID, objc.Sel("addChildWindow:ordered:"), childWin, place)
}

// Adds the provided window as a new tab in a tabbed window using the specified ordering instruction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/addTabbedWindow(_:ordered:)
func (w_ Window) AddTabbedWindowOrdered(window unsafe.Pointer, ordered WindowOrderingMode) {
	objc.Send[objc.ID](w_.ID, objc.Sel("addTabbedWindow:ordered:"), window, ordered)
}

// Adds the specified title bar accessory view controller to the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/addTitlebarAccessoryViewController(_:)
func (w_ Window) AddTitlebarAccessoryViewController(childViewController unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("addTitlebarAccessoryViewController:"), childViewController)
}

// Returns the part of the window that stays stationary during constraint-based layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/anchorAttribute(for:)
func (w_ Window) AnchorAttributeForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("anchorAttributeForOrientation:"), orientation)
	return rv
}

// Specifies the duration of a smooth frame-size change.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/animationResizeTime(_:)
func (w_ Window) AnimationResizeTime(newFrame unsafe.Pointer) float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("animationResizeTime:"), newFrame)
	return rv
}

// Indicates whether the window calculates the thickness of a given border automatically.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/autorecalculatesContentBorderThickness(for:)
func (w_ Window) AutorecalculatesContentBorderThicknessForEdge(edge int) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("autorecalculatesContentBorderThicknessForEdge:"), edge)
	return rv
}

// Returns a backing store pixel-aligned rectangle in window coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backingAlignedRect(_:options:)
func (w_ Window) BackingAlignedRectOptions(rect unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
	return rv
}

// Informs the window that it has become the key window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/becomeKey()
func (w_ Window) BecomeKeyWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("becomeKeyWindow"))
}

// Informs the window that it has become the main window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/becomeMain()
func (w_ Window) BecomeMainWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("becomeMainWindow"))
}

// Starts a document-modal session and presents the specified critical sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/beginCriticalSheet(_:completionHandler:)
func (w_ Window) BeginCriticalSheetCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("beginCriticalSheet:completionHandler:"), sheetWindow, handler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/beginDraggingSession(items:event:source:)
func (w_ Window) BeginDraggingSessionWithItemsEventSource(items unsafe.Pointer, event unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("beginDraggingSessionWithItems:event:source:"), items, event, source)
	return rv
}

// Starts a document-modal session and presents—or queues for presentation—a sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/beginSheet(_:completionHandler:)
func (w_ Window) BeginSheetCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("beginSheet:completionHandler:"), sheetWindow, handler)
}

// Stores the window’s raster image from a given rectangle expressed in the window’s base coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/cacheImage(in:)
func (w_ Window) CacheImageInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("cacheImageInRect:"), rect)
}

// A Boolean value that indicates if the window and its screen use a color space that can represent the specified display gamut.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canRepresent(_:)
func (w_ Window) CanRepresentDisplayGamut(displayGamut unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canRepresentDisplayGamut:"), displayGamut)
	return rv
}

// Indicates whether the window has a depth limit that allows it to store color values.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canStoreColor()
func (w_ Window) CanStoreColor() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canStoreColor"))
	return rv
}

// Positions the window’s top-left to a given point.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/cascadeTopLeft(from:)
func (w_ Window) CascadeTopLeftFromPoint(topLeftPoint unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("cascadeTopLeftFromPoint:"), topLeftPoint)
	return rv
}

// Sets the window’s location to the center of the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/center()
func (w_ Window) Center() {
	objc.Send[objc.ID](w_.ID, objc.Sel("center"))
}

// Removes the window from the screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/close()
func (w_ Window) Close() {
	objc.Send[objc.ID](w_.ID, objc.Sel("close"))
}

// Modifies and returns a frame rectangle so that its top edge lies on a specific screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/constrainFrameRect(_:to:)
func (w_ Window) ConstrainFrameRectToScreen(frameRect unsafe.Pointer, screen unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("constrainFrameRect:toScreen:"), frameRect, screen)
	return rv
}

// Indicates the thickness of a given border of the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentBorderThickness(for:)
func (w_ Window) ContentBorderThicknessForEdge(edge int) float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("contentBorderThicknessForEdge:"), edge)
	return rv
}

// Returns the window’s content rectangle with a given frame rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentRect(forFrameRect:)
func (w_ Window) ContentRectForFrameRect(frameRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("contentRectForFrameRect:"), frameRect)
	return rv
}

// Converts a given point from the window’s base coordinate system to the screen coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertBaseToScreen:
func (w_ Window) ConvertBaseToScreen(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("convertBaseToScreen:"), point)
	return rv
}

// Converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertFromBacking(_:)
func (w_ Window) ConvertRectFromBacking(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}

// Converts a rectangle from the screen coordinate system to the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertFromScreen(_:)
func (w_ Window) ConvertRectFromScreen(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("convertRectFromScreen:"), rect)
	return rv
}

// Converts a point from the screen coordinate system to the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPoint(fromScreen:)
func (w_ Window) ConvertPointFromScreen(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("convertPointFromScreen:"), point)
	return rv
}

// Converts a point to the screen coordinate system from the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPoint(toScreen:)
func (w_ Window) ConvertPointToScreen(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("convertPointToScreen:"), point)
	return rv
}

// Converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPointFromBacking(_:)
func (w_ Window) ConvertPointFromBacking(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("convertPointFromBacking:"), point)
	return rv
}

// Converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPointToBacking(_:)
func (w_ Window) ConvertPointToBacking(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("convertPointToBacking:"), point)
	return rv
}

// Converts a given point from the screen coordinate system to the window’s base coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertScreenToBase:
func (w_ Window) ConvertScreenToBase(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("convertScreenToBase:"), point)
	return rv
}

// Converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertToBacking(_:)
func (w_ Window) ConvertRectToBacking(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("convertRectToBacking:"), rect)
	return rv
}

// Converts a rectangle to the screen coordinate system from the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertToScreen(_:)
func (w_ Window) ConvertRectToScreen(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("convertRectToScreen:"), rect)
	return rv
}

// Returns EPS data that draws the region of the window within a given rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/dataWithEPS(inside:)
func (w_ Window) DataWithEPSInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("dataWithEPSInsideRect:"), rect)
	return rv
}

// Returns PDF data that draws the region of the window within a given rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/dataWithPDF(inside:)
func (w_ Window) DataWithPDFInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("dataWithPDFInsideRect:"), rect)
	return rv
}

// De-minimizes the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/deminiaturize(_:)
func (w_ Window) Deminiaturize(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("deminiaturize:"), sender)
}

// Disables all cursor rectangle management within the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/disableCursorRects()
func (w_ Window) DisableCursorRects() {
	objc.Send[objc.ID](w_.ID, objc.Sel("disableCursorRects"))
}

// Disables the method for the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/disableFlushing()
func (w_ Window) DisableFlushWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("disableFlushWindow"))
}

// Disables the default button cell’s key equivalent, so it doesn’t perform a click when the user presses Return (or Enter).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/disableKeyEquivalentForDefaultButtonCell()
func (w_ Window) DisableKeyEquivalentForDefaultButtonCell() {
	objc.Send[objc.ID](w_.ID, objc.Sel("disableKeyEquivalentForDefaultButtonCell"))
}

// Disables the window’s screen updates until the window is flushed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/disableScreenUpdatesUntilFlush()
func (w_ Window) DisableScreenUpdatesUntilFlush() {
	objc.Send[objc.ID](w_.ID, objc.Sel("disableScreenUpdatesUntilFlush"))
}

// Disables snapshot restoration.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/disableSnapshotRestoration()
func (w_ Window) DisableSnapshotRestoration() {
	objc.Send[objc.ID](w_.ID, objc.Sel("disableSnapshotRestoration"))
}

// Discards all of the window’s cached image rectangles.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/discardCachedImage()
func (w_ Window) DiscardCachedImage() {
	objc.Send[objc.ID](w_.ID, objc.Sel("discardCachedImage"))
}

// Invalidates all cursor rectangles in the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/discardCursorRects()
func (w_ Window) DiscardCursorRects() {
	objc.Send[objc.ID](w_.ID, objc.Sel("discardCursorRects"))
}

// Forwards the message to the global application object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/discardEvents(matching:before:)
func (w_ Window) DiscardEventsMatchingMaskBeforeEvent(mask unsafe.Pointer, lastEvent unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("discardEventsMatchingMask:beforeEvent:"), mask, lastEvent)
}

// Passes a display message down the window’s view hierarchy, thus redrawing all views within the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/display()
func (w_ Window) Display() {
	objc.Send[objc.ID](w_.ID, objc.Sel("display"))
}

// Passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/displayIfNeeded()
func (w_ Window) DisplayIfNeeded() {
	objc.Send[objc.ID](w_.ID, objc.Sel("displayIfNeeded"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/displayLink(target:selector:)
func (w_ Window) DisplayLinkWithTargetSelector(target objc.ID, selector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return rv
}

// Begins a dragging session.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/drag(_:at:offset:event:pasteboard:source:slideBack:)
func (w_ Window) DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, baseLocation unsafe.Pointer, initialOffset unsafe.Pointer, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("dragImage:at:offset:event:pasteboard:source:slideBack:"), image, baseLocation, initialOffset, event, pboard, sourceObj, slideFlag)
}

// Reenables cursor rectangle management within the window after a message.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/enableCursorRects()
func (w_ Window) EnableCursorRects() {
	objc.Send[objc.ID](w_.ID, objc.Sel("enableCursorRects"))
}

// Reenables the method for the window after it was disabled through a previous message.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/enableFlushing()
func (w_ Window) EnableFlushWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("enableFlushWindow"))
}

// Reenables the default button cell’s key equivalent, so it performs a click when the user presses Return (or Enter).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/enableKeyEquivalentForDefaultButtonCell()
func (w_ Window) EnableKeyEquivalentForDefaultButtonCell() {
	objc.Send[objc.ID](w_.ID, objc.Sel("enableKeyEquivalentForDefaultButtonCell"))
}

// Enables snapshot restoration.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/enableSnapshotRestoration()
func (w_ Window) EnableSnapshotRestoration() {
	objc.Send[objc.ID](w_.ID, objc.Sel("enableSnapshotRestoration"))
}

// Forces the field editor to give up its first responder status and prepares it for its next assignment.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/endEditing(for:)
func (w_ Window) EndEditingFor(object objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("endEditingFor:"), object)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/endSheet(_:)-4dmmq
func (w_ Window) EndSheet(sheetWindow unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("endSheet:"), sheetWindow)
}

// Ends a document-modal session and dismisses the specified sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/endSheet(_:returnCode:)
func (w_ Window) EndSheetReturnCode(sheetWindow unsafe.Pointer, returnCode unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("endSheet:returnCode:"), sheetWindow, returnCode)
}

// Returns the window’s field editor, creating it if requested.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/fieldEditor(_:for:)
func (w_ Window) FieldEditorForObject(createFlag bool, object objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("fieldEditor:forObject:"), createFlag, object)
	return rv
}

// Flushes the window’s offscreen buffer to the screen if the window is buffered and flushing is enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/flush()
func (w_ Window) FlushWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("flushWindow"))
}

// Flushes the window’s offscreen buffer to the screen if flushing is enabled and if the last message had no effect because flushing was disabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/flushIfNeeded()
func (w_ Window) FlushWindowIfNeeded() {
	objc.Send[objc.ID](w_.ID, objc.Sel("flushWindowIfNeeded"))
}

// Returns the window’s frame rectangle with a given content rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frameRect(forContentRect:)
func (w_ Window) FrameRectForContentRect(contentRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("frameRectForContentRect:"), contentRect)
	return rv
}

// Returns the window’s graphics state object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/gState()
func (w_ Window) GState() int {
	rv := objc.Send[int](w_.ID, objc.Sel("gState"))
	return rv
}

// Handles the AppleScript command to close the window (and its associated document, if any).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/handleClose(_:)
func (w_ Window) HandleCloseScriptCommand(command unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("handleCloseScriptCommand:"), command)
	return rv
}

// Handles the AppleScript command to print the contents of the window (or its associated document, if any).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/handlePrint(_:)
func (w_ Window) HandlePrintScriptCommand(command unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("handlePrintScriptCommand:"), command)
	return rv
}

// Handles the AppleScript command to save the window (and its associated document, if any).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/handleSave(_:)
func (w_ Window) HandleSaveScriptCommand(command unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("handleSaveScriptCommand:"), command)
	return rv
}

// Inserts the view controller into the window’s array of title bar accessory view controllers at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/insertTitlebarAccessoryViewController(_:at:)
func (w_ Window) InsertTitlebarAccessoryViewControllerAtIndex(childViewController unsafe.Pointer, index int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("insertTitlebarAccessoryViewController:atIndex:"), childViewController, index)
}

// Marks as invalid the cursor rectangles of a given view object in the window, so they’ll be set up again when the window becomes key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/invalidateCursorRects(for:)
func (w_ Window) InvalidateCursorRectsForView(view unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("invalidateCursorRectsForView:"), view)
}

// Invalidates the window shadow so that it is recomputed based on the current window shape.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/invalidateShadow()
func (w_ Window) InvalidateShadow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("invalidateShadow"))
}

// Updates the layout of views in the window based on the current views and constraints.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/layoutIfNeeded()
func (w_ Window) LayoutIfNeeded() {
	objc.Send[objc.ID](w_.ID, objc.Sel("layoutIfNeeded"))
}

// Attempts to make a given responder the first responder for the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/makeFirstResponder(_:)
func (w_ Window) MakeFirstResponder(responder unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("makeFirstResponder:"), responder)
	return rv
}

// Makes the window the key window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/makeKey()
func (w_ Window) MakeKeyWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("makeKeyWindow"))
}

// Moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/makeKeyAndOrderFront(_:)
func (w_ Window) MakeKeyAndOrderFront(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("makeKeyAndOrderFront:"), sender)
}

// Makes the window the main window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/makeMain()
func (w_ Window) MakeMainWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("makeMainWindow"))
}

// Merges all open windows into a single tabbed window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/mergeAllWindows(_:)
func (w_ Window) MergeAllWindows(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("mergeAllWindows:"), sender)
}

// Removes the window from the screen list and displays the minimized window in the Dock.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/miniaturize(_:)
func (w_ Window) Miniaturize(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("miniaturize:"), sender)
}

// Moves the tab to a new containing window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/moveTabToNewWindow(_:)
func (w_ Window) MoveTabToNewWindow(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("moveTabToNewWindow:"), sender)
}

// Returns the next event matching a given mask.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/nextEvent(matching:)
func (w_ Window) NextEventMatchingMask(mask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("nextEventMatchingMask:"), mask)
	return rv
}

// Forwards the message to the global application object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/nextEvent(matching:until:inMode:dequeue:)
func (w_ Window) NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, mode, deqFlag)
	return rv
}

// Repositions the window’s window device in the window server’s screen list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/order(_:relativeTo:)
func (w_ Window) OrderWindowRelativeTo(place WindowOrderingMode, otherWin int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("orderWindow:relativeTo:"), place, otherWin)
}

// Moves the window to the back of its level in the screen list, without changing either the key window or the main window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/orderBack(_:)
func (w_ Window) OrderBack(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("orderBack:"), sender)
}

// Moves the window to the front of its level in the screen list, without changing either the key window or the main window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/orderFront(_:)
func (w_ Window) OrderFront(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("orderFront:"), sender)
}

// Moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/orderFrontRegardless()
func (w_ Window) OrderFrontRegardless() {
	objc.Send[objc.ID](w_.ID, objc.Sel("orderFrontRegardless"))
}

// Removes the window from the screen list, which hides the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/orderOut(_:)
func (w_ Window) OrderOut(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("orderOut:"), sender)
}

// Simulates the user clicking the close button by momentarily highlighting the button and then closing the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/performClose(_:)
func (w_ Window) PerformClose(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performClose:"), sender)
}

// Starts a window drag based on the specified mouse-down event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/performDrag(with:)
func (w_ Window) PerformWindowDragWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performWindowDragWithEvent:"), event)
}

// Simulates the user clicking the minimize button by momentarily highlighting the button, then minimizing the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/performMiniaturize(_:)
func (w_ Window) PerformMiniaturize(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performMiniaturize:"), sender)
}

// This action method simulates the user clicking the zoom box by momentarily highlighting the button and then zooming the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/performZoom(_:)
func (w_ Window) PerformZoom(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performZoom:"), sender)
}

// Forwards the message to the global application object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/postEvent(_:atStart:)
func (w_ Window) PostEventAtStart(event unsafe.Pointer, flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("postEvent:atStart:"), event, flag)
}

// Runs the Print panel, and if the user chooses an option other than canceling, prints the window (its frame view and all subviews).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/printWindow(_:)
func (w_ Window) Print(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("print:"), sender)
}

// Marks the key view loop as “dirty” and in need of recalculation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/recalculateKeyViewLoop()
func (w_ Window) RecalculateKeyViewLoop() {
	objc.Send[objc.ID](w_.ID, objc.Sel("recalculateKeyViewLoop"))
}

// Registers a set of pasteboard types that the window accepts as the destination of an image-dragging session.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/registerForDraggedTypes(_:)
func (w_ Window) RegisterForDraggedTypes(newTypes unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("registerForDraggedTypes:"), newTypes)
}

// Detaches a given child window from the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/removeChildWindow(_:)
func (w_ Window) RemoveChildWindow(childWin unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("removeChildWindow:"), childWin)
}

// Removes the view controller at the specified index from the window’s array of title bar accessory view controllers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/removeTitlebarAccessoryViewController(at:)
func (w_ Window) RemoveTitlebarAccessoryViewControllerAtIndex(index int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("removeTitlebarAccessoryViewControllerAtIndex:"), index)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/requestSharingOfWindow(_:completionHandler:)
func (w_ Window) RequestSharingOfWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("requestSharingOfWindow:completionHandler:"), window, completionHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/requestSharingOfWindow(usingPreview:title:completionHandler:)
func (w_ Window) RequestSharingOfWindowUsingPreviewTitleCompletionHandler(image unsafe.Pointer, title string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("requestSharingOfWindowUsingPreview:title:completionHandler:"), image, objc.String(title), completionHandler)
}

// Clears the window’s cursor rectangles and the cursor rectangles of the objects in its view hierarchy.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resetCursorRects()
func (w_ Window) ResetCursorRects() {
	objc.Send[objc.ID](w_.ID, objc.Sel("resetCursorRects"))
}

// Resigns the window’s key window status.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resignKey()
func (w_ Window) ResignKeyWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("resignKeyWindow"))
}

// Resigns the window’s main window status.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resignMain()
func (w_ Window) ResignMainWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("resignMainWindow"))
}

// Splices the window’s cached image rectangles, if any, back into its raster image (and buffer if it has one), undoing the effect of any drawing performed within those areas since they were established using .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/restoreCachedImage()
func (w_ Window) RestoreCachedImage() {
	objc.Send[objc.ID](w_.ID, objc.Sel("restoreCachedImage"))
}

// Presents the toolbar customization user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/runToolbarCustomizationPalette(_:)
func (w_ Window) RunToolbarCustomizationPalette(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("runToolbarCustomizationPalette:"), sender)
}

// Saves the window’s frame rectangle in the user defaults system under a given name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/saveFrame(usingName:)
func (w_ Window) SaveFrameUsingName(name unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("saveFrameUsingName:"), name)
}

// Gives key view status to the view that follows the given view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectKeyView(following:)
func (w_ Window) SelectKeyViewFollowingView(view unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectKeyViewFollowingView:"), view)
}

// Gives key view status to the view that precedes the given view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectKeyView(preceding:)
func (w_ Window) SelectKeyViewPrecedingView(view unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectKeyViewPrecedingView:"), view)
}

// Searches for a candidate next key view and, if it finds one, tries to make it the first responder.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectNextKeyView(_:)
func (w_ Window) SelectNextKeyView(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectNextKeyView:"), sender)
}

// Selects the next tab in the tab group in the trailing direction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectNextTab(_:)
func (w_ Window) SelectNextTab(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectNextTab:"), sender)
}

// Searches for a candidate previous key view and, if it finds one, tries to make it the first responder.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectPreviousKeyView(_:)
func (w_ Window) SelectPreviousKeyView(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectPreviousKeyView:"), sender)
}

// Selects the previous tab in the tab group in the leading direction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectPreviousTab(_:)
func (w_ Window) SelectPreviousTab(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectPreviousTab:"), sender)
}

// This action method dispatches mouse and keyboard events the global application object sends to the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/sendEvent(_:)
func (w_ Window) SendEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("sendEvent:"), event)
}

// Sets the part of the window that stays stationary during constraint-based layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setAnchorAttribute(_:for:)
func (w_ Window) SetAnchorAttributeForOrientation(attr unsafe.Pointer, orientation unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAnchorAttribute:forOrientation:"), attr, orientation)
}

// Specifies whether the window calculates the thickness of a given border automatically.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setAutorecalculatesContentBorderThickness(_:for:)
func (w_ Window) SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAutorecalculatesContentBorderThickness:forEdge:"), flag, edge)
}

// Specifies the thickness of a given border of the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setContentBorderThickness(_:for:)
func (w_ Window) SetContentBorderThicknessForEdge(thickness float64, edge int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentBorderThickness:forEdge:"), thickness, edge)
}

// Sets the size of the window’s content view to a given size, which is expressed in the window’s base coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setContentSize(_:)
func (w_ Window) SetContentSize(size unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentSize:"), size)
}

// Sets a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setDynamicDepthLimit(_:)
func (w_ Window) SetDynamicDepthLimit(flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDynamicDepthLimit:"), flag)
}

// Sets the origin and size of the window’s frame rectangle according to a given frame rectangle, thereby setting its position and size onscreen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(_:display:)
func (w_ Window) SetFrameDisplay(frameRect unsafe.Pointer, flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrame:display:"), frameRect, flag)
}

// Sets the origin and size of the window’s frame rectangle, with optional animation, according to a given frame rectangle, thereby setting its position and size onscreen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(_:display:animate:)
func (w_ Window) SetFrameDisplayAnimate(frameRect unsafe.Pointer, displayFlag bool, animateFlag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrame:display:animate:"), frameRect, displayFlag, animateFlag)
}

// Sets the window’s frame rectangle from a given string representation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(from:)
func (w_ Window) SetFrameFromString(string unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameFromString:"), string)
}

// Sets the name AppKit uses to automatically save the window’s frame rectangle data in the defaults system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameAutosaveName(_:)
func (w_ Window) SetFrameAutosaveName(name unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("setFrameAutosaveName:"), name)
	return rv
}

// Positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameOrigin(_:)
func (w_ Window) SetFrameOrigin(point unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameOrigin:"), point)
}

// Positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameTopLeftPoint(_:)
func (w_ Window) SetFrameTopLeftPoint(point unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameTopLeftPoint:"), point)
}

// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameUsingName(_:)
func (w_ Window) SetFrameUsingName(name unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("setFrameUsingName:"), name)
	return rv
}

// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system. Can operate on non-resizable windows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameUsingName(_:force:)
func (w_ Window) SetFrameUsingNameForce(name unsafe.Pointer, force bool) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("setFrameUsingName:force:"), name, force)
	return rv
}

// Sets the window’s miniaturized state to the value you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setIsMiniaturized(_:)
func (w_ Window) SetIsMiniaturized(flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsMiniaturized:"), flag)
}

// Sets the window’s visible state to the value you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setIsVisible(_:)
func (w_ Window) SetIsVisible(flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsVisible:"), flag)
}

// Sets the window’s zoomed state to the value you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setIsZoomed(_:)
func (w_ Window) SetIsZoomed(flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsZoomed:"), flag)
}

// Sets a given path as the window’s title, formatting it as a file-system path, and records this path as the window’s associated file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setTitleWithRepresentedFilename(_:)
func (w_ Window) SetTitleWithRepresentedFilename(filename string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitleWithRepresentedFilename:"), objc.String(filename))
}

// Returns the window button of a given window button kind in the window’s view hierarchy.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/standardWindowButton(_:)
func (w_ Window) StandardWindowButton(b unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("standardWindowButton:"), b)
	return rv
}

// Takes the window into or out of fullscreen mode,
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toggleFullScreen(_:)
func (w_ Window) ToggleFullScreen(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("toggleFullScreen:"), sender)
}

// Shows or hides the tab bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toggleTabBar(_:)
func (w_ Window) ToggleTabBar(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("toggleTabBar:"), sender)
}

// Shows or hides the tab overview.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toggleTabOverview(_:)
func (w_ Window) ToggleTabOverview(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("toggleTabOverview:"), sender)
}

// Toggles the visibility of the window’s toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toggleToolbarShown(_:)
func (w_ Window) ToggleToolbarShown(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("toggleToolbarShown:"), sender)
}

// Tracks events that match the specified mask using the specified tracking handler until the tracking handler explicitly terminates tracking.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/trackEvents(matching:timeout:mode:handler:)
func (w_ Window) TrackEventsMatchingMaskTimeoutModeHandler(mask unsafe.Pointer, timeout float64, mode unsafe.Pointer, trackingHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("trackEventsMatchingMask:timeout:mode:handler:"), mask, timeout, mode, trackingHandler)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/transferWindowSharing(to:completionHandler:)
func (w_ Window) TransferWindowSharingToWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("transferWindowSharingToWindow:completionHandler:"), window, completionHandler)
}

// Dispatches action messages with a given argument.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tryToPerform(_:with:)
func (w_ Window) TryToPerformWith(action objc.SEL, object objc.ID) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("tryToPerform:with:"), action, object)
	return rv
}

// Unregisters the window as a possible destination for dragging operations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/unregisterDraggedTypes()
func (w_ Window) UnregisterDraggedTypes() {
	objc.Send[objc.ID](w_.ID, objc.Sel("unregisterDraggedTypes"))
}

// Updates the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/update()
func (w_ Window) Update() {
	objc.Send[objc.ID](w_.ID, objc.Sel("update"))
}

// Updates the constraints based on changes to views in the window since the last layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/updateConstraintsIfNeeded()
func (w_ Window) UpdateConstraintsIfNeeded() {
	objc.Send[objc.ID](w_.ID, objc.Sel("updateConstraintsIfNeeded"))
}

// Specifies whether the window is to optimize focusing and drawing when displaying its views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/useOptimizedDrawing(_:)
func (w_ Window) UseOptimizedDrawing(flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("useOptimizedDrawing:"), flag)
}

// Returns the scale factor applied to the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/userSpaceScaleFactor
func (w_ Window) UserSpaceScaleFactor() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("userSpaceScaleFactor"))
	return rv
}

// Searches for an object that responds to a Services request.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/validRequestor(forSendType:returnType:)
func (w_ Window) ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}

// Displays a visual representation of the supplied constraints in the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/visualizeConstraints(_:)
func (w_ Window) VisualizeConstraints(constraints unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("visualizeConstraints:"), constraints)
}

// Toggles the size and location of the window between its standard state (which the application provides as the best size to display the window’s data) and its user state (a new size and location the user may have set by moving or resizing the window).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/zoom(_:)
func (w_ Window) Zoom(sender objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("zoom:"), sender)
}


