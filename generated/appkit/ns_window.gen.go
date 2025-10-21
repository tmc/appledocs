// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [Window] class.
var (
	WindowClass     _WindowClass
	WindowClassOnce sync.Once
)

func getWindowClass() _WindowClass {
	WindowClassOnce.Do(func() {
		WindowClass = _WindowClass{objc.GetClass("NSWindow")}
	})
	return WindowClass
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
	AnimationResizeTime(newFrame coregraphics.CGRect) float64
	AutorecalculatesContentBorderThicknessForEdge(edge int) bool
	BackingAlignedRectOptions(rect coregraphics.CGRect, options unsafe.Pointer) coregraphics.CGRect
	BecomeKeyWindow()
	BecomeMainWindow()
	BeginCriticalSheetCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer)
	BeginDraggingSessionWithItemsEventSource(items unsafe.Pointer, event unsafe.Pointer, source objc.ID) unsafe.Pointer
	BeginSheetCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer)
	CacheImageInRect(rect coregraphics.CGRect)
	CanRepresentDisplayGamut(displayGamut unsafe.Pointer) bool
	CanStoreColor() bool
	CascadeTopLeftFromPoint(topLeftPoint coregraphics.CGPoint) coregraphics.CGPoint
	Center()
	Close()
	ConstrainFrameRectToScreen(frameRect coregraphics.CGRect, screen unsafe.Pointer) coregraphics.CGRect
	ContentBorderThicknessForEdge(edge int) float64
	ContentRectForFrameRect(frameRect coregraphics.CGRect) coregraphics.CGRect
	ConvertBaseToScreen(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertRectFromBacking(rect coregraphics.CGRect) coregraphics.CGRect
	ConvertRectFromScreen(rect coregraphics.CGRect) coregraphics.CGRect
	ConvertPointFromScreen(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertPointToScreen(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertPointFromBacking(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertPointToBacking(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertScreenToBase(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertRectToBacking(rect coregraphics.CGRect) coregraphics.CGRect
	ConvertRectToScreen(rect coregraphics.CGRect) coregraphics.CGRect
	DataWithEPSInsideRect(rect coregraphics.CGRect) unsafe.Pointer
	DataWithPDFInsideRect(rect coregraphics.CGRect) unsafe.Pointer
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
	DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, baseLocation coregraphics.CGPoint, initialOffset coregraphics.CGSize, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool)
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
	FrameRectForContentRect(contentRect coregraphics.CGRect) coregraphics.CGRect
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
	SetContentSize(size coregraphics.CGSize)
	SetDynamicDepthLimit(flag bool)
	SetFrameDisplay(frameRect coregraphics.CGRect, flag bool)
	SetFrameDisplayAnimate(frameRect coregraphics.CGRect, displayFlag bool, animateFlag bool)
	SetFrameFromString(string_ unsafe.Pointer)
	SetFrameOrigin(point coregraphics.CGPoint)
	SetFrameTopLeftPoint(point coregraphics.CGPoint)
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
// A single object corresponds to, at most, one on-screen window. Windows perform two principal functions: To place views in a provided area To accept and distribute mouse and keyboard events the user generates to the appropriate views
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
func NewWindowWithContentRectStyleMaskBackingDefer(contentRect coregraphics.CGRect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	instance := getWindowClass().Alloc()
	rv := objc.Send[Window](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	rv.Autorelease()
	return rv
}

// Initializes an allocated window with the specified values.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:screen:)
func NewWindowWithContentRectStyleMaskBackingDeferScreen(contentRect coregraphics.CGRect, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen unsafe.Pointer) Window {
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
func (wc _WindowClass) ContentRectForFrameRectStyleMask(fRect coregraphics.CGRect, style WindowStyleMask) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](objc.ID(wc.class), objc.Sel("contentRectForFrameRect:styleMask:"), fRect, style)
	return rv
}

// Returns the frame rectangle used by a window with a given content rectangle and window style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frameRect(forContentRect:styleMask:)
func (wc _WindowClass) FrameRectForContentRectStyleMask(cRect coregraphics.CGRect, style WindowStyleMask) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](objc.ID(wc.class), objc.Sel("frameRectForContentRect:styleMask:"), cRect, style)
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
func (wc _WindowClass) WindowNumberAtPointBelowWindowWithWindowNumber(point coregraphics.CGPoint, windowNumber int) int {
	rv := objc.Send[int](objc.ID(wc.class), objc.Sel("windowNumberAtPoint:belowWindowWithWindowNumber:"), point, windowNumber)
	return rv
}

// Returns the window numbers for all visible windows satisfying the specified options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumbers(options:)
func (wc _WindowClass) WindowNumbersWithOptions(options unsafe.Pointer) []NSNumber {
	rv := objc.Send[[]NSNumber](objc.ID(wc.class), objc.Sel("windowNumbersWithOptions:"), options)
	return rv
}

// A Boolean value that indicates whether the app can automatically organize windows into tabs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsAutomaticWindowTabbing
func (wc _WindowClass) AllowsAutomaticWindowTabbing() bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("allowsAutomaticWindowTabbing"))
	return rv
}
// Returns the default depth limit for instances of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/defaultDepthLimit
func (wc _WindowClass) DefaultDepthLimit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("defaultDepthLimit"))
	return rv
}
// A value that indicates the user’s preference for window tabbing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/userTabbingPreference-swift.type.property
func (wc _WindowClass) UserTabbingPreference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.class), objc.Sel("userTabbingPreference"))
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
func (w_ Window) AnimationResizeTime(newFrame coregraphics.CGRect) float64 {
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
func (w_ Window) BackingAlignedRectOptions(rect coregraphics.CGRect, options unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
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
func (w_ Window) BeginDraggingSessionWithItemsEventSource(items unsafe.Pointer, event unsafe.Pointer, source objc.ID) unsafe.Pointer {
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
func (w_ Window) CacheImageInRect(rect coregraphics.CGRect) {
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
func (w_ Window) CascadeTopLeftFromPoint(topLeftPoint coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](w_.ID, objc.Sel("cascadeTopLeftFromPoint:"), topLeftPoint)
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
func (w_ Window) ConstrainFrameRectToScreen(frameRect coregraphics.CGRect, screen unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("constrainFrameRect:toScreen:"), frameRect, screen)
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
func (w_ Window) ContentRectForFrameRect(frameRect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("contentRectForFrameRect:"), frameRect)
	return rv
}

// Converts a given point from the window’s base coordinate system to the screen coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertBaseToScreen:
func (w_ Window) ConvertBaseToScreen(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](w_.ID, objc.Sel("convertBaseToScreen:"), point)
	return rv
}

// Converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertFromBacking(_:)
func (w_ Window) ConvertRectFromBacking(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}

// Converts a rectangle from the screen coordinate system to the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertFromScreen(_:)
func (w_ Window) ConvertRectFromScreen(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("convertRectFromScreen:"), rect)
	return rv
}

// Converts a point from the screen coordinate system to the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPoint(fromScreen:)
func (w_ Window) ConvertPointFromScreen(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](w_.ID, objc.Sel("convertPointFromScreen:"), point)
	return rv
}

// Converts a point to the screen coordinate system from the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPoint(toScreen:)
func (w_ Window) ConvertPointToScreen(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](w_.ID, objc.Sel("convertPointToScreen:"), point)
	return rv
}

// Converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPointFromBacking(_:)
func (w_ Window) ConvertPointFromBacking(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](w_.ID, objc.Sel("convertPointFromBacking:"), point)
	return rv
}

// Converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPointToBacking(_:)
func (w_ Window) ConvertPointToBacking(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](w_.ID, objc.Sel("convertPointToBacking:"), point)
	return rv
}

// Converts a given point from the screen coordinate system to the window’s base coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertScreenToBase:
func (w_ Window) ConvertScreenToBase(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](w_.ID, objc.Sel("convertScreenToBase:"), point)
	return rv
}

// Converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertToBacking(_:)
func (w_ Window) ConvertRectToBacking(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("convertRectToBacking:"), rect)
	return rv
}

// Converts a rectangle to the screen coordinate system from the window’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertToScreen(_:)
func (w_ Window) ConvertRectToScreen(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("convertRectToScreen:"), rect)
	return rv
}

// Returns EPS data that draws the region of the window within a given rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/dataWithEPS(inside:)
func (w_ Window) DataWithEPSInsideRect(rect coregraphics.CGRect) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("dataWithEPSInsideRect:"), rect)
	return rv
}

// Returns PDF data that draws the region of the window within a given rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/dataWithPDF(inside:)
func (w_ Window) DataWithPDFInsideRect(rect coregraphics.CGRect) unsafe.Pointer {
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
func (w_ Window) DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, baseLocation coregraphics.CGPoint, initialOffset coregraphics.CGSize, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool) {
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
func (w_ Window) FrameRectForContentRect(contentRect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("frameRectForContentRect:"), contentRect)
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
func (w_ Window) SetContentSize(size coregraphics.CGSize) {
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
func (w_ Window) SetFrameDisplay(frameRect coregraphics.CGRect, flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrame:display:"), frameRect, flag)
}

// Sets the origin and size of the window’s frame rectangle, with optional animation, according to a given frame rectangle, thereby setting its position and size onscreen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(_:display:animate:)
func (w_ Window) SetFrameDisplayAnimate(frameRect coregraphics.CGRect, displayFlag bool, animateFlag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrame:display:animate:"), frameRect, displayFlag, animateFlag)
}

// Sets the window’s frame rectangle from a given string representation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(from:)
func (w_ Window) SetFrameFromString(string_ unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameFromString:"), string_)
}

// Positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameOrigin(_:)
func (w_ Window) SetFrameOrigin(point coregraphics.CGPoint) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameOrigin:"), point)
}

// Positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameTopLeftPoint(_:)
func (w_ Window) SetFrameTopLeftPoint(point coregraphics.CGPoint) {
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

// A Boolean value that indicates whether the window accepts mouse-moved events.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/acceptsMouseMovedEvents
func (w_ Window) AcceptsMouseMovedEvents() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("acceptsMouseMovedEvents"))
	return rv
}


// SetAcceptsMouseMovedEvents sets the value of the acceptsMouseMovedEvents property.
// A Boolean value that indicates whether the window accepts mouse-moved events.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/acceptsMouseMovedEvents
func (w_ Window) SetAcceptsMouseMovedEvents(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAcceptsMouseMovedEvents:"), value)
}
// A Boolean value that indicates whether the app can automatically organize windows into tabs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsAutomaticWindowTabbing
func (w_ Window) AllowsAutomaticWindowTabbing() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsAutomaticWindowTabbing"))
	return rv
}


// SetAllowsAutomaticWindowTabbing sets the value of the allowsAutomaticWindowTabbing property.
// A Boolean value that indicates whether the app can automatically organize windows into tabs.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsAutomaticWindowTabbing
func (w_ Window) SetAllowsAutomaticWindowTabbing(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsAutomaticWindowTabbing:"), value)
}
// A Boolean value that indicates whether the window allows multithreaded view drawing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsConcurrentViewDrawing
func (w_ Window) AllowsConcurrentViewDrawing() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsConcurrentViewDrawing"))
	return rv
}


// SetAllowsConcurrentViewDrawing sets the value of the allowsConcurrentViewDrawing property.
// A Boolean value that indicates whether the window allows multithreaded view drawing.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsConcurrentViewDrawing
func (w_ Window) SetAllowsConcurrentViewDrawing(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsConcurrentViewDrawing:"), value)
}
// A Boolean value that indicates whether the window can display tooltips even when the application is in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsToolTipsWhenApplicationIsInactive
func (w_ Window) AllowsToolTipsWhenApplicationIsInactive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsToolTipsWhenApplicationIsInactive"))
	return rv
}


// SetAllowsToolTipsWhenApplicationIsInactive sets the value of the allowsToolTipsWhenApplicationIsInactive property.
// A Boolean value that indicates whether the window can display tooltips even when the application is in the background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsToolTipsWhenApplicationIsInactive
func (w_ Window) SetAllowsToolTipsWhenApplicationIsInactive(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsToolTipsWhenApplicationIsInactive:"), value)
}
// The window’s alpha value.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/alphaValue
func (w_ Window) AlphaValue() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("alphaValue"))
	return rv
}


// SetAlphaValue sets the value of the alphaValue property.
// The window’s alpha value.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/alphaValue
func (w_ Window) SetAlphaValue(value float64) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAlphaValue:"), value)
}
// The window’s automatic animation behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/animationBehavior-swift.property
func (w_ Window) AnimationBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("animationBehavior"))
	return rv
}


// SetAnimationBehavior sets the value of the animationBehavior property.
// The window’s automatic animation behavior.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/animationBehavior-swift.property
func (w_ Window) SetAnimationBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAnimationBehavior:"), value)
}
// An object that the window inherits its appearance from.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/appearanceSource
func (w_ Window) AppearanceSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("appearanceSource"))
	return rv
}


// SetAppearanceSource sets the value of the appearanceSource property.
// An object that the window inherits its appearance from.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/appearanceSource
func (w_ Window) SetAppearanceSource(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAppearanceSource:"), value)
}
// A Boolean value that indicates whether the window’s cursor rectangles are enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/areCursorRectsEnabled
func (w_ Window) AreCursorRectsEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("areCursorRectsEnabled"))
	return rv
}

// The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/aspectRatio
func (w_ Window) AspectRatio() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](w_.ID, objc.Sel("aspectRatio"))
	return rv
}


// SetAspectRatio sets the value of the aspectRatio property.
// The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/aspectRatio
func (w_ Window) SetAspectRatio(value coregraphics.CGSize) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAspectRatio:"), value)
}
// The sheet attached to the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/attachedSheet
func (w_ Window) AttachedSheet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("attachedSheet"))
	return rv
}

// A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/autorecalculatesKeyViewLoop
func (w_ Window) AutorecalculatesKeyViewLoop() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("autorecalculatesKeyViewLoop"))
	return rv
}


// SetAutorecalculatesKeyViewLoop sets the value of the autorecalculatesKeyViewLoop property.
// A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/autorecalculatesKeyViewLoop
func (w_ Window) SetAutorecalculatesKeyViewLoop(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAutorecalculatesKeyViewLoop:"), value)
}
// The color of the window’s background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backgroundColor
func (w_ Window) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The color of the window’s background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backgroundColor
func (w_ Window) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setBackgroundColor:"), value)
}
// The location of the window’s backing store.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backingLocation-swift.property
func (w_ Window) BackingLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("backingLocation"))
	return rv
}

// The backing scale factor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backingScaleFactor
func (w_ Window) BackingScaleFactor() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("backingScaleFactor"))
	return rv
}

// The window’s backing store type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backingType
func (w_ Window) BackingType() BackingStoreType {
	rv := objc.Send[BackingStoreType](w_.ID, objc.Sel("backingType"))
	return rv
}


// SetBackingType sets the value of the backingType property.
// The window’s backing store type.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backingType
func (w_ Window) SetBackingType(value BackingStoreType) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setBackingType:"), value)
}
// A Boolean value that indicates whether the window can become the key window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeKey
func (w_ Window) CanBecomeKeyWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canBecomeKeyWindow"))
	return rv
}

// A Boolean value that indicates whether the window can become the application’s main window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeMain
func (w_ Window) CanBecomeMainWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canBecomeMainWindow"))
	return rv
}

// A Boolean value that indicates whether the window can be displayed at the login window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeVisibleWithoutLogin
func (w_ Window) CanBecomeVisibleWithoutLogin() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canBecomeVisibleWithoutLogin"))
	return rv
}


// SetCanBecomeVisibleWithoutLogin sets the value of the canBecomeVisibleWithoutLogin property.
// A Boolean value that indicates whether the window can be displayed at the login window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeVisibleWithoutLogin
func (w_ Window) SetCanBecomeVisibleWithoutLogin(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanBecomeVisibleWithoutLogin:"), value)
}
// A Boolean value that indicates whether the window can hide when its application becomes hidden.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canHide
func (w_ Window) CanHide() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canHide"))
	return rv
}


// SetCanHide sets the value of the canHide property.
// A Boolean value that indicates whether the window can hide when its application becomes hidden.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canHide
func (w_ Window) SetCanHide(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanHide:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/cascadingReferenceFrame
func (w_ Window) CascadingReferenceFrame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("cascadingReferenceFrame"))
	return rv
}

// An array of the window’s attached child windows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/childWindows
func (w_ Window) ChildWindows() []Window {
	rv := objc.Send[[]Window](w_.ID, objc.Sel("childWindows"))
	return rv
}

// A value that identifies the window’s behavior in window collections.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/collectionBehavior-swift.property
func (w_ Window) CollectionBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("collectionBehavior"))
	return rv
}


// SetCollectionBehavior sets the value of the collectionBehavior property.
// A value that identifies the window’s behavior in window collections.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/collectionBehavior-swift.property
func (w_ Window) SetCollectionBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCollectionBehavior:"), value)
}
// The window’s color space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/colorSpace
func (w_ Window) ColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("colorSpace"))
	return rv
}


// SetColorSpace sets the value of the colorSpace property.
// The window’s color space.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/colorSpace
func (w_ Window) SetColorSpace(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setColorSpace:"), value)
}
// The window’s content aspect ratio.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentAspectRatio
func (w_ Window) ContentAspectRatio() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](w_.ID, objc.Sel("contentAspectRatio"))
	return rv
}


// SetContentAspectRatio sets the value of the contentAspectRatio property.
// The window’s content aspect ratio.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentAspectRatio
func (w_ Window) SetContentAspectRatio(value coregraphics.CGSize) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentAspectRatio:"), value)
}
// A value used by Auto Layout constraints to automatically bind to the value of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentLayoutGuide
func (w_ Window) ContentLayoutGuide() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("contentLayoutGuide"))
	return rv
}

// The area inside the window that is for non-obscured content, in window coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentLayoutRect
func (w_ Window) ContentLayoutRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("contentLayoutRect"))
	return rv
}

// The maximum size of the window’s content view in the window’s base coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentMaxSize
func (w_ Window) ContentMaxSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](w_.ID, objc.Sel("contentMaxSize"))
	return rv
}


// SetContentMaxSize sets the value of the contentMaxSize property.
// The maximum size of the window’s content view in the window’s base coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentMaxSize
func (w_ Window) SetContentMaxSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentMaxSize:"), value)
}
// The minimum size of the window’s content view in the window’s base coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentMinSize
func (w_ Window) ContentMinSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](w_.ID, objc.Sel("contentMinSize"))
	return rv
}


// SetContentMinSize sets the value of the contentMinSize property.
// The minimum size of the window’s content view in the window’s base coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentMinSize
func (w_ Window) SetContentMinSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentMinSize:"), value)
}
// The window’s content-view resizing increments.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentResizeIncrements
func (w_ Window) ContentResizeIncrements() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](w_.ID, objc.Sel("contentResizeIncrements"))
	return rv
}


// SetContentResizeIncrements sets the value of the contentResizeIncrements property.
// The window’s content-view resizing increments.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentResizeIncrements
func (w_ Window) SetContentResizeIncrements(value coregraphics.CGSize) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentResizeIncrements:"), value)
}
// The window’s content view, the highest accessible view object in the window’s view hierarchy.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentView
func (w_ Window) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("contentView"))
	return rv
}


// SetContentView sets the value of the contentView property.
// The window’s content view, the highest accessible view object in the window’s view hierarchy.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentView
func (w_ Window) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentView:"), value)
}
// The main content view controller for the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentViewController
func (w_ Window) ContentViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("contentViewController"))
	return rv
}


// SetContentViewController sets the value of the contentViewController property.
// The main content view controller for the window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentViewController
func (w_ Window) SetContentViewController(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentViewController:"), value)
}
// The event currently being processed by the application.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/currentEvent
func (w_ Window) CurrentEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("currentEvent"))
	return rv
}

// The deepest screen the window is on (it may be split over several screens).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/deepestScreen
func (w_ Window) DeepestScreen() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("deepestScreen"))
	return rv
}

// The button cell that performs as if clicked when the window receives a Return (or Enter) key event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/defaultButtonCell
func (w_ Window) DefaultButtonCell() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("defaultButtonCell"))
	return rv
}


// SetDefaultButtonCell sets the value of the defaultButtonCell property.
// The button cell that performs as if clicked when the window receives a Return (or Enter) key event.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/defaultButtonCell
func (w_ Window) SetDefaultButtonCell(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultButtonCell:"), value)
}
// Returns the default depth limit for instances of .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/defaultDepthLimit
func (w_ Window) DefaultDepthLimit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("defaultDepthLimit"))
	return rv
}

// The window’s delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/delegate
func (w_ Window) Delegate() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The window’s delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/delegate
func (w_ Window) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDelegate:"), value)
}
// The depth limit of the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/depthLimit
func (w_ Window) DepthLimit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("depthLimit"))
	return rv
}


// SetDepthLimit sets the value of the depthLimit property.
// The depth limit of the window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/depthLimit
func (w_ Window) SetDepthLimit(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDepthLimit:"), value)
}
// A dictionary containing information about the window’s resolution, such as color, depth, and so on.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/deviceDescription
func (w_ Window) DeviceDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("deviceDescription"))
	return rv
}

// A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/displaysWhenScreenProfileChanges
func (w_ Window) DisplaysWhenScreenProfileChanges() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("displaysWhenScreenProfileChanges"))
	return rv
}


// SetDisplaysWhenScreenProfileChanges sets the value of the displaysWhenScreenProfileChanges property.
// A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/displaysWhenScreenProfileChanges
func (w_ Window) SetDisplaysWhenScreenProfileChanges(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplaysWhenScreenProfileChanges:"), value)
}
// The application’s Dock tile.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/dockTile
func (w_ Window) DockTile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("dockTile"))
	return rv
}

// The collection of drawers associated with the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/drawers
func (w_ Window) Drawers() []Drawer {
	rv := objc.Send[[]Drawer](w_.ID, objc.Sel("drawers"))
	return rv
}

// The window’s first responder.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/firstResponder
func (w_ Window) FirstResponder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("firstResponder"))
	return rv
}

// The window’s frame rectangle in screen coordinates, including the title bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frame
func (w_ Window) Frame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](w_.ID, objc.Sel("frame"))
	return rv
}

// The name used to automatically save the window’s frame rectangle data in the defaults system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frameAutosaveName-swift.property
func (w_ Window) FrameAutosaveName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("frameAutosaveName"))
	return rv
}

// A string representation of the window’s frame rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frameDescriptor
func (w_ Window) StringWithSavedFrame() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("stringWithSavedFrame"))
	return rv
}

// The graphics context associated with the window for the current thread.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/graphicsContext
func (w_ Window) GraphicsContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("graphicsContext"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hasActiveWindowSharingSession
func (w_ Window) HasActiveWindowSharingSession() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasActiveWindowSharingSession"))
	return rv
}

// A Boolean value that indicates if the window has a close box.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hasCloseBox
func (w_ Window) HasCloseBox() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasCloseBox"))
	return rv
}

// A Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hasDynamicDepthLimit
func (w_ Window) HasDynamicDepthLimit() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasDynamicDepthLimit"))
	return rv
}

// A Boolean value that indicates whether the window has a shadow.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hasShadow
func (w_ Window) HasShadow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasShadow"))
	return rv
}


// SetHasShadow sets the value of the hasShadow property.
// A Boolean value that indicates whether the window has a shadow.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hasShadow
func (w_ Window) SetHasShadow(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasShadow:"), value)
}
// A Boolean value that indicates if the window has a title bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hasTitleBar
func (w_ Window) HasTitleBar() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasTitleBar"))
	return rv
}

// A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hidesOnDeactivate
func (w_ Window) HidesOnDeactivate() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hidesOnDeactivate"))
	return rv
}


// SetHidesOnDeactivate sets the value of the hidesOnDeactivate property.
// A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hidesOnDeactivate
func (w_ Window) SetHidesOnDeactivate(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHidesOnDeactivate:"), value)
}
// A Boolean value that indicates whether the window is transparent to mouse events.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ignoresMouseEvents
func (w_ Window) IgnoresMouseEvents() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("ignoresMouseEvents"))
	return rv
}


// SetIgnoresMouseEvents sets the value of the ignoresMouseEvents property.
// A Boolean value that indicates whether the window is transparent to mouse events.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ignoresMouseEvents
func (w_ Window) SetIgnoresMouseEvents(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIgnoresMouseEvents:"), value)
}
// A Boolean value that indicates whether the window is being resized by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/inLiveResize
func (w_ Window) InLiveResize() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("inLiveResize"))
	return rv
}

// The view that’s made first responder (also called the key view) the first time the window is placed onscreen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/initialFirstResponder
func (w_ Window) InitialFirstResponder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("initialFirstResponder"))
	return rv
}


// SetInitialFirstResponder sets the value of the initialFirstResponder property.
// The view that’s made first responder (also called the key view) the first time the window is placed onscreen.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/initialFirstResponder
func (w_ Window) SetInitialFirstResponder(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInitialFirstResponder:"), value)
}
// A Boolean value that indicates whether the window automatically displays views that need to be displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isAutodisplay
func (w_ Window) Autodisplay() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("autodisplay"))
	return rv
}


// SetAutodisplay sets the value of the autodisplay property.
// A Boolean value that indicates whether the window automatically displays views that need to be displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isAutodisplay
func (w_ Window) SetAutodisplay(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAutodisplay:"), value)
}
// A Boolean value that indicates whether the window’s document has been edited.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isDocumentEdited
func (w_ Window) DocumentEdited() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("documentEdited"))
	return rv
}


// SetDocumentEdited sets the value of the documentEdited property.
// A Boolean value that indicates whether the window’s document has been edited.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isDocumentEdited
func (w_ Window) SetDocumentEdited(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDocumentEdited:"), value)
}
// A Boolean value that indicates whether the window is excluded from the application’s Windows menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isExcludedFromWindowsMenu
func (w_ Window) ExcludedFromWindowsMenu() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("excludedFromWindowsMenu"))
	return rv
}


// SetExcludedFromWindowsMenu sets the value of the excludedFromWindowsMenu property.
// A Boolean value that indicates whether the window is excluded from the application’s Windows menu.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isExcludedFromWindowsMenu
func (w_ Window) SetExcludedFromWindowsMenu(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setExcludedFromWindowsMenu:"), value)
}
// A Boolean value that indicates whether the window is a floating panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isFloatingPanel
func (w_ Window) FloatingPanel() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("floatingPanel"))
	return rv
}

// A Boolean value that indicates whether the window’s flushing ability is disabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isFlushWindowDisabled
func (w_ Window) FlushWindowDisabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("flushWindowDisabled"))
	return rv
}

// A Boolean value that indicates whether the window is the key window for the application.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isKeyWindow
func (w_ Window) KeyWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("keyWindow"))
	return rv
}

// A Boolean value that indicates whether the window is the application’s main window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMainWindow
func (w_ Window) MainWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("mainWindow"))
	return rv
}

// A Boolean value that indicates whether the window can minimize.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMiniaturizable
func (w_ Window) Miniaturizable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("miniaturizable"))
	return rv
}

// A Boolean value that indicates whether the window is minimized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMiniaturized
func (w_ Window) Miniaturized() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("miniaturized"))
	return rv
}

// A Boolean value that indicates whether the window is a modal panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isModalPanel
func (w_ Window) ModalPanel() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("modalPanel"))
	return rv
}

// A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMovable
func (w_ Window) Movable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("movable"))
	return rv
}


// SetMovable sets the value of the movable property.
// A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMovable
func (w_ Window) SetMovable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMovable:"), value)
}
// A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMovableByWindowBackground
func (w_ Window) MovableByWindowBackground() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("movableByWindowBackground"))
	return rv
}


// SetMovableByWindowBackground sets the value of the movableByWindowBackground property.
// A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMovableByWindowBackground
func (w_ Window) SetMovableByWindowBackground(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMovableByWindowBackground:"), value)
}
// A Boolean value that indicates whether the window is on the currently active space.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isOnActiveSpace
func (w_ Window) OnActiveSpace() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("onActiveSpace"))
	return rv
}

// A Boolean value that indicates whether the window device the window manages is freed when it’s removed from the screen list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isOneShot
func (w_ Window) OneShot() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("oneShot"))
	return rv
}


// SetOneShot sets the value of the oneShot property.
// A Boolean value that indicates whether the window device the window manages is freed when it’s removed from the screen list.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isOneShot
func (w_ Window) SetOneShot(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOneShot:"), value)
}
// A Boolean value that indicates whether the window is opaque.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isOpaque
func (w_ Window) Opaque() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("opaque"))
	return rv
}


// SetOpaque sets the value of the opaque property.
// A Boolean value that indicates whether the window is opaque.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isOpaque
func (w_ Window) SetOpaque(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOpaque:"), value)
}
// A Boolean value that indicates whether the window is released when it receives the message.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isReleasedWhenClosed
func (w_ Window) ReleasedWhenClosed() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("releasedWhenClosed"))
	return rv
}


// SetReleasedWhenClosed sets the value of the releasedWhenClosed property.
// A Boolean value that indicates whether the window is released when it receives the message.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isReleasedWhenClosed
func (w_ Window) SetReleasedWhenClosed(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setReleasedWhenClosed:"), value)
}
// A Boolean value that indicates if the user can resize the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isResizable
func (w_ Window) Resizable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("resizable"))
	return rv
}

// A Boolean value indicating whether the window configuration is preserved between application launches.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isRestorable
func (w_ Window) Restorable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("restorable"))
	return rv
}


// SetRestorable sets the value of the restorable property.
// A Boolean value indicating whether the window configuration is preserved between application launches.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isRestorable
func (w_ Window) SetRestorable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRestorable:"), value)
}
// A Boolean value that indicates whether the window has ever run as a modal sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isSheet
func (w_ Window) Sheet() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("sheet"))
	return rv
}

// A Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isVisible
func (w_ Window) Visible() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("visible"))
	return rv
}

// A Boolean value that indicates whether the window allows zooming.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isZoomable
func (w_ Window) Zoomable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("zoomable"))
	return rv
}

// A Boolean value that indicates whether the window is in a zoomed state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isZoomed
func (w_ Window) Zoomed() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("zoomed"))
	return rv
}

// The direction the window is currently using to change the key view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/keyViewSelectionDirection
func (w_ Window) KeyViewSelectionDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("keyViewSelectionDirection"))
	return rv
}

// The window level of the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/level-swift.property
func (w_ Window) Level() WindowLevel {
	rv := objc.Send[WindowLevel](w_.ID, objc.Sel("level"))
	return rv
}


// SetLevel sets the value of the level property.
// The window level of the window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/level-swift.property
func (w_ Window) SetLevel(value WindowLevel) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setLevel:"), value)
}
// A maximum size that is used to determine if a window can fit when it is in full screen in a tile.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/maxFullScreenContentSize
func (w_ Window) MaxFullScreenContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](w_.ID, objc.Sel("maxFullScreenContentSize"))
	return rv
}


// SetMaxFullScreenContentSize sets the value of the maxFullScreenContentSize property.
// A maximum size that is used to determine if a window can fit when it is in full screen in a tile.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/maxFullScreenContentSize
func (w_ Window) SetMaxFullScreenContentSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMaxFullScreenContentSize:"), value)
}
// The maximum size to which the window’s frame (including its title bar) can be sized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/maxSize
func (w_ Window) MaxSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](w_.ID, objc.Sel("maxSize"))
	return rv
}


// SetMaxSize sets the value of the maxSize property.
// The maximum size to which the window’s frame (including its title bar) can be sized.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/maxSize
func (w_ Window) SetMaxSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMaxSize:"), value)
}
// A minimum size that is used to determine if a window can fit when it is in full screen in a tile.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/minFullScreenContentSize
func (w_ Window) MinFullScreenContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](w_.ID, objc.Sel("minFullScreenContentSize"))
	return rv
}


// SetMinFullScreenContentSize sets the value of the minFullScreenContentSize property.
// A minimum size that is used to determine if a window can fit when it is in full screen in a tile.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/minFullScreenContentSize
func (w_ Window) SetMinFullScreenContentSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMinFullScreenContentSize:"), value)
}
// The minimum size to which the window’s frame (including its title bar) can be sized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/minSize
func (w_ Window) MinSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](w_.ID, objc.Sel("minSize"))
	return rv
}


// SetMinSize sets the value of the minSize property.
// The minimum size to which the window’s frame (including its title bar) can be sized.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/minSize
func (w_ Window) SetMinSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMinSize:"), value)
}
// The custom miniaturized window image of the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowImage
func (w_ Window) MiniwindowImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("miniwindowImage"))
	return rv
}


// SetMiniwindowImage sets the value of the miniwindowImage property.
// The custom miniaturized window image of the window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowImage
func (w_ Window) SetMiniwindowImage(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMiniwindowImage:"), value)
}
// The title displayed in the window’s minimized window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowTitle
func (w_ Window) MiniwindowTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("miniwindowTitle"))
	return rv
}


// SetMiniwindowTitle sets the value of the miniwindowTitle property.
// The title displayed in the window’s minimized window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowTitle
func (w_ Window) SetMiniwindowTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMiniwindowTitle:"), value)
}
// The current location of the pointer reckoned in the window’s base coordinate system, regardless of the current event being handled or of any events pending.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/mouseLocationOutsideOfEventStream
func (w_ Window) MouseLocationOutsideOfEventStream() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](w_.ID, objc.Sel("mouseLocationOutsideOfEventStream"))
	return rv
}

// The occlusion state of the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/occlusionState-swift.property
func (w_ Window) OcclusionState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("occlusionState"))
	return rv
}

// The zero-based position of the window, based on its order from front to back among all visible application windows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/orderedIndex
func (w_ Window) OrderedIndex() int {
	rv := objc.Send[int](w_.ID, objc.Sel("orderedIndex"))
	return rv
}


// SetOrderedIndex sets the value of the orderedIndex property.
// The zero-based position of the window, based on its order from front to back among all visible application windows.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/orderedIndex
func (w_ Window) SetOrderedIndex(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOrderedIndex:"), value)
}
// The parent window to which the window is attached as a child.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/parent
func (w_ Window) ParentWindow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("parentWindow"))
	return rv
}


// SetParentWindow sets the value of the parentWindow property.
// The parent window to which the window is attached as a child.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/parent
func (w_ Window) SetParentWindow(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setParentWindow:"), value)
}
// A Boolean value that indicates the preferred location for the window’s backing store.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/preferredBackingLocation
func (w_ Window) PreferredBackingLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("preferredBackingLocation"))
	return rv
}


// SetPreferredBackingLocation sets the value of the preferredBackingLocation property.
// A Boolean value that indicates the preferred location for the window’s backing store.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/preferredBackingLocation
func (w_ Window) SetPreferredBackingLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferredBackingLocation:"), value)
}
// A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/preservesContentDuringLiveResize
func (w_ Window) PreservesContentDuringLiveResize() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("preservesContentDuringLiveResize"))
	return rv
}


// SetPreservesContentDuringLiveResize sets the value of the preservesContentDuringLiveResize property.
// A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/preservesContentDuringLiveResize
func (w_ Window) SetPreservesContentDuringLiveResize(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreservesContentDuringLiveResize:"), value)
}
// A Boolean value that indicates whether the window prevents application termination when modal.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/preventsApplicationTerminationWhenModal
func (w_ Window) PreventsApplicationTerminationWhenModal() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("preventsApplicationTerminationWhenModal"))
	return rv
}


// SetPreventsApplicationTerminationWhenModal sets the value of the preventsApplicationTerminationWhenModal property.
// A Boolean value that indicates whether the window prevents application termination when modal.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/preventsApplicationTerminationWhenModal
func (w_ Window) SetPreventsApplicationTerminationWhenModal(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreventsApplicationTerminationWhenModal:"), value)
}
// The path to the file of the window’s represented file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/representedFilename
func (w_ Window) RepresentedFilename() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("representedFilename"))
	return rv
}


// SetRepresentedFilename sets the value of the representedFilename property.
// The path to the file of the window’s represented file.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/representedFilename
func (w_ Window) SetRepresentedFilename(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRepresentedFilename:"), value)
}
// The URL of the file the window represents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/representedURL
func (w_ Window) RepresentedURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("representedURL"))
	return rv
}


// SetRepresentedURL sets the value of the representedURL property.
// The URL of the file the window represents.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/representedURL
func (w_ Window) SetRepresentedURL(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRepresentedURL:"), value)
}
// The flags field of the event record for the mouse-down event that initiated the resizing session.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resizeFlags
func (w_ Window) ResizeFlags() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](w_.ID, objc.Sel("resizeFlags"))
	return rv
}

// The window’s resizing increments.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resizeIncrements
func (w_ Window) ResizeIncrements() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](w_.ID, objc.Sel("resizeIncrements"))
	return rv
}


// SetResizeIncrements sets the value of the resizeIncrements property.
// The window’s resizing increments.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resizeIncrements
func (w_ Window) SetResizeIncrements(value coregraphics.CGSize) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setResizeIncrements:"), value)
}
// The restoration class associated with the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/restorationClass
func (w_ Window) RestorationClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("restorationClass"))
	return rv
}


// SetRestorationClass sets the value of the restorationClass property.
// The restoration class associated with the window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/restorationClass
func (w_ Window) SetRestorationClass(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRestorationClass:"), value)
}
// The screen the window is on.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/screen
func (w_ Window) Screen() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("screen"))
	return rv
}

// A Boolean value that indicates the level of access other processes have to the window’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/sharingType-swift.property
func (w_ Window) SharingType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("sharingType"))
	return rv
}


// SetSharingType sets the value of the sharingType property.
// A Boolean value that indicates the level of access other processes have to the window’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/sharingType-swift.property
func (w_ Window) SetSharingType(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSharingType:"), value)
}
// The window to which the sheet is attached.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/sheetParent
func (w_ Window) SheetParent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("sheetParent"))
	return rv
}

// An array of the sheets currently attached to the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/sheets
func (w_ Window) Sheets() []Window {
	rv := objc.Send[[]Window](w_.ID, objc.Sel("sheets"))
	return rv
}

// A Boolean value that indicates whether the window’s resize indicator is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/showsResizeIndicator
func (w_ Window) ShowsResizeIndicator() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("showsResizeIndicator"))
	return rv
}


// SetShowsResizeIndicator sets the value of the showsResizeIndicator property.
// A Boolean value that indicates whether the window’s resize indicator is visible.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/showsResizeIndicator
func (w_ Window) SetShowsResizeIndicator(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShowsResizeIndicator:"), value)
}
// A Boolean value that indicates whether the toolbar control button is currently displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/showsToolbarButton
func (w_ Window) ShowsToolbarButton() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("showsToolbarButton"))
	return rv
}


// SetShowsToolbarButton sets the value of the showsToolbarButton property.
// A Boolean value that indicates whether the toolbar control button is currently displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/showsToolbarButton
func (w_ Window) SetShowsToolbarButton(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShowsToolbarButton:"), value)
}
// Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/styleMask-swift.property
func (w_ Window) StyleMask() WindowStyleMask {
	rv := objc.Send[WindowStyleMask](w_.ID, objc.Sel("styleMask"))
	return rv
}


// SetStyleMask sets the value of the styleMask property.
// Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/styleMask-swift.property
func (w_ Window) SetStyleMask(value WindowStyleMask) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setStyleMask:"), value)
}
// A secondary line of text that appears in the title bar of the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/subtitle
func (w_ Window) Subtitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("subtitle"))
	return rv
}


// SetSubtitle sets the value of the subtitle property.
// A secondary line of text that appears in the title bar of the window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/subtitle
func (w_ Window) SetSubtitle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSubtitle:"), value)
}
// An object that represents information about a window when it displays as a tab.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tab
func (w_ Window) Tab() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("tab"))
	return rv
}

// A group of windows that display together as a tab group.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabGroup
func (w_ Window) TabGroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("tabGroup"))
	return rv
}

// An array of windows that display as tabs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabbedWindows
func (w_ Window) TabbedWindows() []Window {
	rv := objc.Send[[]Window](w_.ID, objc.Sel("tabbedWindows"))
	return rv
}

// A value that allows a group of related windows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingIdentifier-swift.property
func (w_ Window) TabbingIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("tabbingIdentifier"))
	return rv
}


// SetTabbingIdentifier sets the value of the tabbingIdentifier property.
// A value that allows a group of related windows.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingIdentifier-swift.property
func (w_ Window) SetTabbingIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTabbingIdentifier:"), value)
}
// A value that indicates when a window displays tabs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingMode-swift.property
func (w_ Window) TabbingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("tabbingMode"))
	return rv
}


// SetTabbingMode sets the value of the tabbingMode property.
// A value that indicates when a window displays tabs.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingMode-swift.property
func (w_ Window) SetTabbingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTabbingMode:"), value)
}
// The string that appears in the title bar of the window or the path to the represented file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/title
func (w_ Window) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The string that appears in the title bar of the window or the path to the represented file.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/title
func (w_ Window) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitle:"), value)
}
// A value that indicates the visibility of the window’s title and title bar buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/titleVisibility-swift.property
func (w_ Window) TitleVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("titleVisibility"))
	return rv
}


// SetTitleVisibility sets the value of the titleVisibility property.
// A value that indicates the visibility of the window’s title and title bar buttons.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/titleVisibility-swift.property
func (w_ Window) SetTitleVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitleVisibility:"), value)
}
// An array of title bar accessory view controllers that are currently added to the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/titlebarAccessoryViewControllers
func (w_ Window) TitlebarAccessoryViewControllers() []TitlebarAccessoryViewController {
	rv := objc.Send[[]TitlebarAccessoryViewController](w_.ID, objc.Sel("titlebarAccessoryViewControllers"))
	return rv
}


// SetTitlebarAccessoryViewControllers sets the value of the titlebarAccessoryViewControllers property.
// An array of title bar accessory view controllers that are currently added to the window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/titlebarAccessoryViewControllers
func (w_ Window) SetTitlebarAccessoryViewControllers(value []TitlebarAccessoryViewController) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitlebarAccessoryViewControllers:"), nsArray)
}
// A Boolean value that indicates whether the title bar draws its background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/titlebarAppearsTransparent
func (w_ Window) TitlebarAppearsTransparent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("titlebarAppearsTransparent"))
	return rv
}


// SetTitlebarAppearsTransparent sets the value of the titlebarAppearsTransparent property.
// A Boolean value that indicates whether the title bar draws its background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/titlebarAppearsTransparent
func (w_ Window) SetTitlebarAppearsTransparent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitlebarAppearsTransparent:"), value)
}
// The type of separator that the app displays between the title bar and content of a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/titlebarSeparatorStyle
func (w_ Window) TitlebarSeparatorStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("titlebarSeparatorStyle"))
	return rv
}


// SetTitlebarSeparatorStyle sets the value of the titlebarSeparatorStyle property.
// The type of separator that the app displays between the title bar and content of a window.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/titlebarSeparatorStyle
func (w_ Window) SetTitlebarSeparatorStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitlebarSeparatorStyle:"), value)
}
// The window’s toolbar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toolbar
func (w_ Window) Toolbar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("toolbar"))
	return rv
}


// SetToolbar sets the value of the toolbar property.
// The window’s toolbar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toolbar
func (w_ Window) SetToolbar(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setToolbar:"), value)
}
// The style that determines the appearance and location of the toolbar in relation to the title bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toolbarStyle-swift.property
func (w_ Window) ToolbarStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("toolbarStyle"))
	return rv
}


// SetToolbarStyle sets the value of the toolbarStyle property.
// The style that determines the appearance and location of the toolbar in relation to the title bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toolbarStyle-swift.property
func (w_ Window) SetToolbarStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setToolbarStyle:"), value)
}
// A value that indicates the user’s preference for window tabbing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/userTabbingPreference-swift.type.property
func (w_ Window) UserTabbingPreference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("userTabbingPreference"))
	return rv
}

// A Boolean value that indicates whether any of the window’s views need to be displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/viewsNeedDisplay
func (w_ Window) ViewsNeedDisplay() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("viewsNeedDisplay"))
	return rv
}


// SetViewsNeedDisplay sets the value of the viewsNeedDisplay property.
// A Boolean value that indicates whether any of the window’s views need to be displayed.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/viewsNeedDisplay
func (w_ Window) SetViewsNeedDisplay(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setViewsNeedDisplay:"), value)
}
// The window’s window controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowController
func (w_ Window) WindowController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("windowController"))
	return rv
}


// SetWindowController sets the value of the windowController property.
// The window’s window controller.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowController
func (w_ Window) SetWindowController(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowController:"), value)
}
// The window number of the window’s window device.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumber
func (w_ Window) WindowNumber() int {
	rv := objc.Send[int](w_.ID, objc.Sel("windowNumber"))
	return rv
}

// The direction the window’s title bar lays text out, either left to right or right to left.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowTitlebarLayoutDirection
func (w_ Window) WindowTitlebarLayoutDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("windowTitlebarLayoutDirection"))
	return rv
}

// A Boolean value that indicates whether the window is able to receive keyboard and mouse events even when some other window is being run modally.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/worksWhenModal
func (w_ Window) WorksWhenModal() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("worksWhenModal"))
	return rv
}


