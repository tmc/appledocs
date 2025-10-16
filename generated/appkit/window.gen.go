
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Window] class.
var WindowClass _WindowClass

func init() {
	WindowClass = _WindowClass{objc.GetClass("NSWindow")}
}

type _WindowClass struct {
	objc.Class
}

// An interface definition for the [Window] class.
type IWindow interface {
	ID() objc.ID
	AddChildWindowOrdered(childWin unsafe.Pointer, place unsafe.Pointer)
	AddTabbedWindowOrdered(window unsafe.Pointer, ordered unsafe.Pointer)
	AddTitlebarAccessoryViewController(childViewController unsafe.Pointer)
	AnchorAttributeForOrientation(orientation unsafe.Pointer) unsafe.Pointer
	AnimationResizeTime(newFrame unsafe.Pointer) unsafe.Pointer
	AutorecalculatesContentBorderThicknessForEdge(edge unsafe.Pointer) bool
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
	ContentBorderThicknessForEdge(edge unsafe.Pointer) float64
	ContentRectForFrameRect(frameRect unsafe.Pointer) unsafe.Pointer
	ConvertBaseToScreen(point unsafe.Pointer) unsafe.Pointer
	ConvertPointFromBacking(point unsafe.Pointer) unsafe.Pointer
	ConvertPointFromScreen(point unsafe.Pointer) unsafe.Pointer
	ConvertPointToBacking(point unsafe.Pointer) unsafe.Pointer
	ConvertPointToScreen(point unsafe.Pointer) unsafe.Pointer
	ConvertRectFromBacking(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectFromScreen(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectToBacking(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectToScreen(rect unsafe.Pointer) unsafe.Pointer
	ConvertScreenToBase(point unsafe.Pointer) unsafe.Pointer
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
	InitWithContentRectStyleMaskBackingDefer(contentRect unsafe.Pointer, style unsafe.Pointer, backingStoreType unsafe.Pointer, flag bool) unsafe.Pointer
	InitWithContentRectStyleMaskBackingDeferScreen(contentRect unsafe.Pointer, style unsafe.Pointer, backingStoreType unsafe.Pointer, flag bool, screen unsafe.Pointer) unsafe.Pointer
	InitWithWindowRef(windowRef unsafe.Pointer) unsafe.Pointer
	InsertTitlebarAccessoryViewControllerAtIndex(childViewController unsafe.Pointer, index int)
	InvalidateCursorRectsForView(view unsafe.Pointer)
	InvalidateShadow()
	LayoutIfNeeded()
	MakeFirstResponder(responder unsafe.Pointer) bool
	MakeKeyAndOrderFront(sender objc.ID)
	MakeKeyWindow()
	MakeMainWindow()
	MergeAllWindows(sender objc.ID)
	Miniaturize(sender objc.ID)
	MoveTabToNewWindow(sender objc.ID)
	NextEventMatchingMask(mask unsafe.Pointer) unsafe.Pointer
	NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer
	OrderBack(sender objc.ID)
	OrderFront(sender objc.ID)
	OrderFrontRegardless()
	OrderOut(sender objc.ID)
	OrderWindowRelativeTo(place unsafe.Pointer, otherWin int)
	PerformClose(sender objc.ID)
	PerformMiniaturize(sender objc.ID)
	PerformWindowDragWithEvent(event unsafe.Pointer)
	PerformZoom(sender objc.ID)
	PostEventAtStart(event unsafe.Pointer, flag bool)
	Print(sender objc.ID)
	RecalculateKeyViewLoop()
	RegisterForDraggedTypes(newTypes unsafe.Pointer)
	RemoveChildWindow(childWin unsafe.Pointer)
	RemoveTitlebarAccessoryViewControllerAtIndex(index int)
	RequestSharingOfWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer)
	RequestSharingOfWindowUsingPreviewTitleCompletionHandler(image unsafe.Pointer, title unsafe.Pointer, completionHandler unsafe.Pointer)
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
	SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge unsafe.Pointer)
	SetContentBorderThicknessForEdge(thickness float64, edge unsafe.Pointer)
	SetContentSize(size unsafe.Pointer)
	SetDynamicDepthLimit(flag bool)
	SetFrameAutosaveName(name unsafe.Pointer) bool
	SetFrameDisplay(frameRect unsafe.Pointer, flag bool)
	SetFrameDisplayAnimate(frameRect unsafe.Pointer, displayFlag bool, animateFlag bool)
	SetFrameFromString(string unsafe.Pointer)
	SetFrameOrigin(point unsafe.Pointer)
	SetFrameTopLeftPoint(point unsafe.Pointer)
	SetFrameUsingName(name unsafe.Pointer) bool
	SetFrameUsingNameForce(name unsafe.Pointer, force bool) bool
	SetIsMiniaturized(flag bool)
	SetIsVisible(flag bool)
	SetIsZoomed(flag bool)
	SetTitleWithRepresentedFilename(filename unsafe.Pointer)
	StandardWindowButton(b unsafe.Pointer) unsafe.Pointer
	ToggleFullScreen(sender objc.ID)
	ToggleTabBar(sender objc.ID)
	ToggleTabOverview(sender objc.ID)
	ToggleToolbarShown(sender objc.ID)
	TrackEventsMatchingMaskTimeoutModeHandler(mask unsafe.Pointer, timeout unsafe.Pointer, mode unsafe.Pointer, trackingHandler unsafe.Pointer)
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

type Window struct {
	id objc.ID
}

func WindowFrom(ptr unsafe.Pointer) Window {
	return Window{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (w_ Window) ID() objc.ID {
	return w_.id
}

// Alloc allocates a new instance without initialization.
func (wc _WindowClass) Alloc() Window {
	rv := objc.Send[Window](objc.ID(wc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (wc _WindowClass) New() Window {
	rv := objc.Send[Window](objc.ID(wc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewWindow creates and returns a new initialized instance.
func NewWindow() Window {
	return WindowClass.New()
}

// Init initializes the instance.
func (w_ Window) Init() Window {
	rv := objc.Send[Window](w_.ID(), selInit)
	return rv
}
// Returns the content rectangle used by a window with a given frame rectangle and window style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentRect(forFrameRect:styleMask:)
func (wc _WindowClass) ContentRectForFrameRectStyleMask(fRect unsafe.Pointer, style unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.Class), objc.RegisterName("contentRectForFrameRect:styleMask:"), fRect, style)
	return rv
}
// Returns the frame rectangle used by a window with a given content rectangle and window style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/frameRect(forContentRect:styleMask:)
func (wc _WindowClass) FrameRectForContentRectStyleMask(cRect unsafe.Pointer, style unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.Class), objc.RegisterName("frameRectForContentRect:styleMask:"), cRect, style)
	return rv
}
// Creates a titled window that contains the specified content view controller. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/init(contentViewController:)
func (wc _WindowClass) WindowWithContentViewController(contentViewController unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.Class), objc.RegisterName("windowWithContentViewController:"), contentViewController)
	return rv
}

// Window_WindowWithContentViewController creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/init(contentViewController:)
func Window_WindowWithContentViewController(contentViewController unsafe.Pointer) unsafe.Pointer {
	return WindowClass.WindowWithContentViewController(contentViewController)
}
// This method does nothing; it is here for backward compatibility. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/menuChanged(_:)
func (wc _WindowClass) MenuChanged(menu unsafe.Pointer)  {
	objc.Send[objc.ID](objc.ID(wc.Class), objc.RegisterName("menuChanged:"), menu)
}
// Returns the minimum width a window’s frame rectangle must have for it to display a title, with a given window style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/minFrameWidth(withTitle:styleMask:)
func (wc _WindowClass) MinFrameWidthWithTitleStyleMask(title unsafe.Pointer, style unsafe.Pointer) float64 {
	rv := objc.Send[float64](objc.ID(wc.Class), objc.RegisterName("minFrameWidthWithTitle:styleMask:"), title, style)
	return rv
}
// Removes the frame data stored under a given name from the application’s user defaults. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/removeFrame(usingName:)
func (wc _WindowClass) RemoveFrameUsingName(name unsafe.Pointer)  {
	objc.Send[objc.ID](objc.ID(wc.Class), objc.RegisterName("removeFrameUsingName:"), name)
}
// Returns a new instance of a given standard window button, sized appropriately for a given window style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/standardWindowButton(_:for:)
func (wc _WindowClass) StandardWindowButtonForStyleMask(b unsafe.Pointer, styleMask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.Class), objc.RegisterName("standardWindowButton:forStyleMask:"), b, styleMask)
	return rv
}
// Returns the number of the frontmost window that would be hit by a mouse-down at the specified screen location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/windowNumber(at:belowWindowWithWindowNumber:)
func (wc _WindowClass) WindowNumberAtPointBelowWindowWithWindowNumber(point unsafe.Pointer, windowNumber int) int {
	rv := objc.Send[int](objc.ID(wc.Class), objc.RegisterName("windowNumberAtPoint:belowWindowWithWindowNumber:"), point, windowNumber)
	return rv
}

// Window_WindowNumberAtPointBelowWindowWithWindowNumber creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/windowNumber(at:belowWindowWithWindowNumber:)
func Window_WindowNumberAtPointBelowWindowWithWindowNumber(point unsafe.Pointer, windowNumber int) int {
	return WindowClass.WindowNumberAtPointBelowWindowWithWindowNumber(point, windowNumber)
}
// Returns the window numbers for all visible windows satisfying the specified options. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/windowNumbers(options:)
func (wc _WindowClass) WindowNumbersWithOptions(options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(wc.Class), objc.RegisterName("windowNumbersWithOptions:"), options)
	return rv
}

// Window_WindowNumbersWithOptions creates a new instance via class method. [Full Topic]
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/windowNumbers(options:)
func Window_WindowNumbersWithOptions(options unsafe.Pointer) unsafe.Pointer {
	return WindowClass.WindowNumbersWithOptions(options)
}
// Adds a given window as a child window of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/addChildWindow(_:ordered:)
func (w_ Window) AddChildWindowOrdered(childWin unsafe.Pointer, place unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("addChildWindow:ordered:"), childWin, place)
}
// Adds the provided window as a new tab in a tabbed window using the specified ordering instruction. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/addTabbedWindow(_:ordered:)
func (w_ Window) AddTabbedWindowOrdered(window unsafe.Pointer, ordered unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("addTabbedWindow:ordered:"), window, ordered)
}
// Adds the specified title bar accessory view controller to the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/addTitlebarAccessoryViewController(_:)
func (w_ Window) AddTitlebarAccessoryViewController(childViewController unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("addTitlebarAccessoryViewController:"), childViewController)
}
// Returns the part of the window that stays stationary during constraint-based layout. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/anchorAttribute(for:)
func (w_ Window) AnchorAttributeForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("anchorAttributeForOrientation:"), orientation)
	return rv
}
// Specifies the duration of a smooth frame-size change. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/animationResizeTime(_:)
func (w_ Window) AnimationResizeTime(newFrame unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("animationResizeTime:"), newFrame)
	return rv
}
// Indicates whether the window calculates the thickness of a given border automatically. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/autorecalculatesContentBorderThickness(for:)
func (w_ Window) AutorecalculatesContentBorderThicknessForEdge(edge unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("autorecalculatesContentBorderThicknessForEdge:"), edge)
	return rv
}
// Returns a backing store pixel-aligned rectangle in window coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/backingAlignedRect(_:options:)
func (w_ Window) BackingAlignedRectOptions(rect unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("backingAlignedRect:options:"), rect, options)
	return rv
}
// Informs the window that it has become the key window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/becomeKey()
func (w_ Window) BecomeKeyWindow() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("becomeKeyWindow"))
}
// Informs the window that it has become the main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/becomeMain()
func (w_ Window) BecomeMainWindow() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("becomeMainWindow"))
}
// Starts a document-modal session and presents the specified critical sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/beginCriticalSheet(_:completionHandler:)
func (w_ Window) BeginCriticalSheetCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("beginCriticalSheet:completionHandler:"), sheetWindow, handler)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/beginDraggingSession(items:event:source:)
func (w_ Window) BeginDraggingSessionWithItemsEventSource(items unsafe.Pointer, event unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("beginDraggingSessionWithItems:event:source:"), items, event, source)
	return rv
}
// Starts a document-modal session and presents—or queues for presentation—a sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/beginSheet(_:completionHandler:)
func (w_ Window) BeginSheetCompletionHandler(sheetWindow unsafe.Pointer, handler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("beginSheet:completionHandler:"), sheetWindow, handler)
}
// Stores the window’s raster image from a given rectangle expressed in the window’s base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/cacheImage(in:)
func (w_ Window) CacheImageInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("cacheImageInRect:"), rect)
}
// A Boolean value that indicates if the window and its screen use a color space that can represent the specified display gamut. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/canRepresent(_:)
func (w_ Window) CanRepresentDisplayGamut(displayGamut unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("canRepresentDisplayGamut:"), displayGamut)
	return rv
}
// Indicates whether the window has a depth limit that allows it to store color values. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/canStoreColor()
func (w_ Window) CanStoreColor() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("canStoreColor"))
	return rv
}
// Positions the window’s top-left to a given point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/cascadeTopLeft(from:)
func (w_ Window) CascadeTopLeftFromPoint(topLeftPoint unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("cascadeTopLeftFromPoint:"), topLeftPoint)
	return rv
}
// Sets the window’s location to the center of the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/center()
func (w_ Window) Center() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("center"))
}
// Removes the window from the screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/close()
func (w_ Window) Close() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("close"))
}
// Modifies and returns a frame rectangle so that its top edge lies on a specific screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/constrainFrameRect(_:to:)
func (w_ Window) ConstrainFrameRectToScreen(frameRect unsafe.Pointer, screen unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("constrainFrameRect:toScreen:"), frameRect, screen)
	return rv
}
// Indicates the thickness of a given border of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentBorderThickness(for:)
func (w_ Window) ContentBorderThicknessForEdge(edge unsafe.Pointer) float64 {
	rv := objc.Send[float64](w_.ID(), objc.RegisterName("contentBorderThicknessForEdge:"), edge)
	return rv
}
// Returns the window’s content rectangle with a given frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentRect(forFrameRect:)
func (w_ Window) ContentRectForFrameRect(frameRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("contentRectForFrameRect:"), frameRect)
	return rv
}
// Converts a given point from the window’s base coordinate system to the screen coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertBaseToScreen:
func (w_ Window) ConvertBaseToScreen(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("convertBaseToScreen:"), point)
	return rv
}
// Converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertFromBacking(_:)
func (w_ Window) ConvertRectFromBacking(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("convertRectFromBacking:"), rect)
	return rv
}
// Converts a rectangle from the screen coordinate system to the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertFromScreen(_:)
func (w_ Window) ConvertRectFromScreen(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("convertRectFromScreen:"), rect)
	return rv
}
// Converts a point from the screen coordinate system to the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertPoint(fromScreen:)
func (w_ Window) ConvertPointFromScreen(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("convertPointFromScreen:"), point)
	return rv
}
// Converts a point to the screen coordinate system from the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertPoint(toScreen:)
func (w_ Window) ConvertPointToScreen(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("convertPointToScreen:"), point)
	return rv
}
// Converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertPointFromBacking(_:)
func (w_ Window) ConvertPointFromBacking(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("convertPointFromBacking:"), point)
	return rv
}
// Converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertPointToBacking(_:)
func (w_ Window) ConvertPointToBacking(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("convertPointToBacking:"), point)
	return rv
}
// Converts a given point from the screen coordinate system to the window’s base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertScreenToBase:
func (w_ Window) ConvertScreenToBase(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("convertScreenToBase:"), point)
	return rv
}
// Converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertToBacking(_:)
func (w_ Window) ConvertRectToBacking(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("convertRectToBacking:"), rect)
	return rv
}
// Converts a rectangle to the screen coordinate system from the window’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/convertToScreen(_:)
func (w_ Window) ConvertRectToScreen(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("convertRectToScreen:"), rect)
	return rv
}
// Returns EPS data that draws the region of the window within a given rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/dataWithEPS(inside:)
func (w_ Window) DataWithEPSInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("dataWithEPSInsideRect:"), rect)
	return rv
}
// Returns PDF data that draws the region of the window within a given rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/dataWithPDF(inside:)
func (w_ Window) DataWithPDFInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("dataWithPDFInsideRect:"), rect)
	return rv
}
// De-minimizes the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/deminiaturize(_:)
func (w_ Window) Deminiaturize(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("deminiaturize:"), sender)
}
// Disables all cursor rectangle management within the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/disableCursorRects()
func (w_ Window) DisableCursorRects() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("disableCursorRects"))
}
// Disables the   method for the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/disableFlushing()
func (w_ Window) DisableFlushWindow() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("disableFlushWindow"))
}
// Disables the default button cell’s key equivalent, so it doesn’t perform a click when the user presses Return (or Enter). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/disableKeyEquivalentForDefaultButtonCell()
func (w_ Window) DisableKeyEquivalentForDefaultButtonCell() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("disableKeyEquivalentForDefaultButtonCell"))
}
// Disables the window’s screen updates until the window is flushed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/disableScreenUpdatesUntilFlush()
func (w_ Window) DisableScreenUpdatesUntilFlush() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("disableScreenUpdatesUntilFlush"))
}
// Disables snapshot restoration. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/disableSnapshotRestoration()
func (w_ Window) DisableSnapshotRestoration() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("disableSnapshotRestoration"))
}
// Discards all of the window’s cached image rectangles. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/discardCachedImage()
func (w_ Window) DiscardCachedImage() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("discardCachedImage"))
}
// Invalidates all cursor rectangles in the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/discardCursorRects()
func (w_ Window) DiscardCursorRects() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("discardCursorRects"))
}
// Forwards the message to the global application object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/discardEvents(matching:before:)
func (w_ Window) DiscardEventsMatchingMaskBeforeEvent(mask unsafe.Pointer, lastEvent unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("discardEventsMatchingMask:beforeEvent:"), mask, lastEvent)
}
// Passes a display message down the window’s view hierarchy, thus redrawing all views within the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/display()
func (w_ Window) Display() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("display"))
}
// Passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/displayIfNeeded()
func (w_ Window) DisplayIfNeeded() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("displayIfNeeded"))
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/displayLink(target:selector:)
func (w_ Window) DisplayLinkWithTargetSelector(target objc.ID, selector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("displayLinkWithTarget:selector:"), target, selector)
	return rv
}
// Begins a dragging session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/drag(_:at:offset:event:pasteboard:source:slideBack:)
func (w_ Window) DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, baseLocation unsafe.Pointer, initialOffset unsafe.Pointer, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("dragImage:at:offset:event:pasteboard:source:slideBack:"), image, baseLocation, initialOffset, event, pboard, sourceObj, slideFlag)
}
// Reenables cursor rectangle management within the window after a   message. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/enableCursorRects()
func (w_ Window) EnableCursorRects() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("enableCursorRects"))
}
// Reenables the   method for the window after it was disabled through a previous   message. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/enableFlushing()
func (w_ Window) EnableFlushWindow() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("enableFlushWindow"))
}
// Reenables the default button cell’s key equivalent, so it performs a click when the user presses Return (or Enter). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/enableKeyEquivalentForDefaultButtonCell()
func (w_ Window) EnableKeyEquivalentForDefaultButtonCell() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("enableKeyEquivalentForDefaultButtonCell"))
}
// Enables snapshot restoration. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/enableSnapshotRestoration()
func (w_ Window) EnableSnapshotRestoration() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("enableSnapshotRestoration"))
}
// Forces the field editor to give up its first responder status and prepares it for its next assignment. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/endEditing(for:)
func (w_ Window) EndEditingFor(object objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("endEditingFor:"), object)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/endSheet(_:)-4dmmq
func (w_ Window) EndSheet(sheetWindow unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("endSheet:"), sheetWindow)
}
// Ends a document-modal session and dismisses the specified sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/endSheet(_:returnCode:)
func (w_ Window) EndSheetReturnCode(sheetWindow unsafe.Pointer, returnCode unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("endSheet:returnCode:"), sheetWindow, returnCode)
}
// Returns the window’s field editor, creating it if requested. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/fieldEditor(_:for:)
func (w_ Window) FieldEditorForObject(createFlag bool, object objc.ID) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("fieldEditor:forObject:"), createFlag, object)
	return rv
}
// Flushes the window’s offscreen buffer to the screen if the window is buffered and flushing is enabled. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/flush()
func (w_ Window) FlushWindow() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("flushWindow"))
}
// Flushes the window’s offscreen buffer to the screen if flushing is enabled and if the last   message had no effect because flushing was disabled. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/flushIfNeeded()
func (w_ Window) FlushWindowIfNeeded() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("flushWindowIfNeeded"))
}
// Returns the window’s frame rectangle with a given content rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/frameRect(forContentRect:)
func (w_ Window) FrameRectForContentRect(contentRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("frameRectForContentRect:"), contentRect)
	return rv
}
// Returns the window’s graphics state object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/gState()
func (w_ Window) GState() int {
	rv := objc.Send[int](w_.ID(), objc.RegisterName("gState"))
	return rv
}
// Handles the AppleScript command to close the window (and its associated document, if any). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/handleClose(_:)
func (w_ Window) HandleCloseScriptCommand(command unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](w_.ID(), objc.RegisterName("handleCloseScriptCommand:"), command)
	return rv
}
// Handles the AppleScript command to print the contents of the window (or its associated document, if any). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/handlePrint(_:)
func (w_ Window) HandlePrintScriptCommand(command unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](w_.ID(), objc.RegisterName("handlePrintScriptCommand:"), command)
	return rv
}
// Handles the AppleScript command to save the window (and its associated document, if any). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/handleSave(_:)
func (w_ Window) HandleSaveScriptCommand(command unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](w_.ID(), objc.RegisterName("handleSaveScriptCommand:"), command)
	return rv
}
// Initializes the window with the specified values. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:)
func (w_ Window) InitWithContentRectStyleMaskBackingDefer(contentRect unsafe.Pointer, style unsafe.Pointer, backingStoreType unsafe.Pointer, flag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return rv
}
// Initializes an allocated window with the specified values. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:screen:)
func (w_ Window) InitWithContentRectStyleMaskBackingDeferScreen(contentRect unsafe.Pointer, style unsafe.Pointer, backingStoreType unsafe.Pointer, flag bool, screen unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	return rv
}
// Returns a Cocoa window created from a Carbon window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/init(windowRef:)
func (w_ Window) InitWithWindowRef(windowRef unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("initWithWindowRef:"), windowRef)
	return rv
}
// Inserts the view controller into the window’s array of title bar accessory view controllers at the specified index. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/insertTitlebarAccessoryViewController(_:at:)
func (w_ Window) InsertTitlebarAccessoryViewControllerAtIndex(childViewController unsafe.Pointer, index int) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("insertTitlebarAccessoryViewController:atIndex:"), childViewController, index)
}
// Marks as invalid the cursor rectangles of a given view object in the window, so they’ll be set up again when the window becomes key. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/invalidateCursorRects(for:)
func (w_ Window) InvalidateCursorRectsForView(view unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("invalidateCursorRectsForView:"), view)
}
// Invalidates the window shadow so that it is recomputed based on the current window shape. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/invalidateShadow()
func (w_ Window) InvalidateShadow() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("invalidateShadow"))
}
// Updates the layout of views in the window based on the current views and constraints. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/layoutIfNeeded()
func (w_ Window) LayoutIfNeeded() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("layoutIfNeeded"))
}
// Attempts to make a given responder the first responder for the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/makeFirstResponder(_:)
func (w_ Window) MakeFirstResponder(responder unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("makeFirstResponder:"), responder)
	return rv
}
// Makes the window the key window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/makeKey()
func (w_ Window) MakeKeyWindow() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("makeKeyWindow"))
}
// Moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/makeKeyAndOrderFront(_:)
func (w_ Window) MakeKeyAndOrderFront(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("makeKeyAndOrderFront:"), sender)
}
// Makes the window the main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/makeMain()
func (w_ Window) MakeMainWindow() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("makeMainWindow"))
}
// Merges all open windows into a single tabbed window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/mergeAllWindows(_:)
func (w_ Window) MergeAllWindows(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("mergeAllWindows:"), sender)
}
// Removes the window from the screen list and displays the minimized window in the Dock. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/miniaturize(_:)
func (w_ Window) Miniaturize(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("miniaturize:"), sender)
}
// Moves the tab to a new containing window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/moveTabToNewWindow(_:)
func (w_ Window) MoveTabToNewWindow(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("moveTabToNewWindow:"), sender)
}
// Returns the next event matching a given mask. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/nextEvent(matching:)
func (w_ Window) NextEventMatchingMask(mask unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("nextEventMatchingMask:"), mask)
	return rv
}
// Forwards the message to the global application object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/nextEvent(matching:until:inMode:dequeue:)
func (w_ Window) NextEventMatchingMaskUntilDateInModeDequeue(mask unsafe.Pointer, expiration unsafe.Pointer, mode unsafe.Pointer, deqFlag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, mode, deqFlag)
	return rv
}
// Repositions the window’s window device in the window server’s screen list. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/order(_:relativeTo:)
func (w_ Window) OrderWindowRelativeTo(place unsafe.Pointer, otherWin int) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("orderWindow:relativeTo:"), place, otherWin)
}
// Moves the window to the back of its level in the screen list, without changing either the key window or the main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/orderBack(_:)
func (w_ Window) OrderBack(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("orderBack:"), sender)
}
// Moves the window to the front of its level in the screen list, without changing either the key window or the main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/orderFront(_:)
func (w_ Window) OrderFront(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("orderFront:"), sender)
}
// Moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/orderFrontRegardless()
func (w_ Window) OrderFrontRegardless() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("orderFrontRegardless"))
}
// Removes the window from the screen list, which hides the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/orderOut(_:)
func (w_ Window) OrderOut(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("orderOut:"), sender)
}
// Simulates the user clicking the close button by momentarily highlighting the button and then closing the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/performClose(_:)
func (w_ Window) PerformClose(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("performClose:"), sender)
}
// Starts a window drag based on the specified mouse-down event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/performDrag(with:)
func (w_ Window) PerformWindowDragWithEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("performWindowDragWithEvent:"), event)
}
// Simulates the user clicking the minimize button by momentarily highlighting the button, then minimizing the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/performMiniaturize(_:)
func (w_ Window) PerformMiniaturize(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("performMiniaturize:"), sender)
}
// This action method simulates the user clicking the zoom box by momentarily highlighting the button and then zooming the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/performZoom(_:)
func (w_ Window) PerformZoom(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("performZoom:"), sender)
}
// Forwards the message to the global application object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/postEvent(_:atStart:)
func (w_ Window) PostEventAtStart(event unsafe.Pointer, flag bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("postEvent:atStart:"), event, flag)
}
// Runs the Print panel, and if the user chooses an option other than canceling, prints the window (its frame view and all subviews). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/printWindow(_:)
func (w_ Window) Print(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("print:"), sender)
}
// Marks the key view loop as “dirty” and in need of recalculation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/recalculateKeyViewLoop()
func (w_ Window) RecalculateKeyViewLoop() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("recalculateKeyViewLoop"))
}
// Registers a set of pasteboard types that the window accepts as the destination of an image-dragging session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/registerForDraggedTypes(_:)
func (w_ Window) RegisterForDraggedTypes(newTypes unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("registerForDraggedTypes:"), newTypes)
}
// Detaches a given child window from the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/removeChildWindow(_:)
func (w_ Window) RemoveChildWindow(childWin unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("removeChildWindow:"), childWin)
}
// Removes the view controller at the specified index from the window’s array of title bar accessory view controllers. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/removeTitlebarAccessoryViewController(at:)
func (w_ Window) RemoveTitlebarAccessoryViewControllerAtIndex(index int) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("removeTitlebarAccessoryViewControllerAtIndex:"), index)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/requestSharingOfWindow(_:completionHandler:)
func (w_ Window) RequestSharingOfWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("requestSharingOfWindow:completionHandler:"), window, completionHandler)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/requestSharingOfWindow(usingPreview:title:completionHandler:)
func (w_ Window) RequestSharingOfWindowUsingPreviewTitleCompletionHandler(image unsafe.Pointer, title unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("requestSharingOfWindowUsingPreview:title:completionHandler:"), image, title, completionHandler)
}
// Clears the window’s cursor rectangles and the cursor rectangles of the   objects in its view hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/resetCursorRects()
func (w_ Window) ResetCursorRects() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("resetCursorRects"))
}
// Resigns the window’s key window status. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/resignKey()
func (w_ Window) ResignKeyWindow() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("resignKeyWindow"))
}
// Resigns the window’s main window status. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/resignMain()
func (w_ Window) ResignMainWindow() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("resignMainWindow"))
}
// Splices the window’s cached image rectangles, if any, back into its raster image (and buffer if it has one), undoing the effect of any drawing performed within those areas since they were established using  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/restoreCachedImage()
func (w_ Window) RestoreCachedImage() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("restoreCachedImage"))
}
// Presents the toolbar customization user interface. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/runToolbarCustomizationPalette(_:)
func (w_ Window) RunToolbarCustomizationPalette(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("runToolbarCustomizationPalette:"), sender)
}
// Saves the window’s frame rectangle in the user defaults system under a given name. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/saveFrame(usingName:)
func (w_ Window) SaveFrameUsingName(name unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("saveFrameUsingName:"), name)
}
// Gives key view status to the view that follows the given view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectKeyView(following:)
func (w_ Window) SelectKeyViewFollowingView(view unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("selectKeyViewFollowingView:"), view)
}
// Gives key view status to the view that precedes the given view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectKeyView(preceding:)
func (w_ Window) SelectKeyViewPrecedingView(view unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("selectKeyViewPrecedingView:"), view)
}
// Searches for a candidate next key view and, if it finds one, tries to make it the first responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectNextKeyView(_:)
func (w_ Window) SelectNextKeyView(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("selectNextKeyView:"), sender)
}
// Selects the next tab in the tab group in the trailing direction. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectNextTab(_:)
func (w_ Window) SelectNextTab(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("selectNextTab:"), sender)
}
// Searches for a candidate previous key view and, if it finds one, tries to make it the first responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectPreviousKeyView(_:)
func (w_ Window) SelectPreviousKeyView(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("selectPreviousKeyView:"), sender)
}
// Selects the previous tab in the tab group in the leading direction. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/selectPreviousTab(_:)
func (w_ Window) SelectPreviousTab(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("selectPreviousTab:"), sender)
}
// This action method dispatches mouse and keyboard events the global application object sends to the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/sendEvent(_:)
func (w_ Window) SendEvent(event unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("sendEvent:"), event)
}
// Sets the part of the window that stays stationary during constraint-based layout. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setAnchorAttribute(_:for:)
func (w_ Window) SetAnchorAttributeForOrientation(attr unsafe.Pointer, orientation unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAnchorAttribute:forOrientation:"), attr, orientation)
}
// Specifies whether the window calculates the thickness of a given border automatically. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setAutorecalculatesContentBorderThickness(_:for:)
func (w_ Window) SetAutorecalculatesContentBorderThicknessForEdge(flag bool, edge unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAutorecalculatesContentBorderThickness:forEdge:"), flag, edge)
}
// Specifies the thickness of a given border of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setContentBorderThickness(_:for:)
func (w_ Window) SetContentBorderThicknessForEdge(thickness float64, edge unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setContentBorderThickness:forEdge:"), thickness, edge)
}
// Sets the size of the window’s content view to a given size, which is expressed in the window’s base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setContentSize(_:)
func (w_ Window) SetContentSize(size unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setContentSize:"), size)
}
// Sets a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setDynamicDepthLimit(_:)
func (w_ Window) SetDynamicDepthLimit(flag bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setDynamicDepthLimit:"), flag)
}
// Sets the origin and size of the window’s frame rectangle according to a given frame rectangle, thereby setting its position and size onscreen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrame(_:display:)
func (w_ Window) SetFrameDisplay(frameRect unsafe.Pointer, flag bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setFrame:display:"), frameRect, flag)
}
// Sets the origin and size of the window’s frame rectangle, with optional animation, according to a given frame rectangle, thereby setting its position and size onscreen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrame(_:display:animate:)
func (w_ Window) SetFrameDisplayAnimate(frameRect unsafe.Pointer, displayFlag bool, animateFlag bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setFrame:display:animate:"), frameRect, displayFlag, animateFlag)
}
// Sets the window’s frame rectangle from a given string representation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrame(from:)
func (w_ Window) SetFrameFromString(string unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setFrameFromString:"), string)
}
// Sets the name AppKit uses to automatically save the window’s frame rectangle data in the defaults system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrameAutosaveName(_:)
func (w_ Window) SetFrameAutosaveName(name unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("setFrameAutosaveName:"), name)
	return rv
}
// Positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrameOrigin(_:)
func (w_ Window) SetFrameOrigin(point unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setFrameOrigin:"), point)
}
// Positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrameTopLeftPoint(_:)
func (w_ Window) SetFrameTopLeftPoint(point unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setFrameTopLeftPoint:"), point)
}
// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrameUsingName(_:)
func (w_ Window) SetFrameUsingName(name unsafe.Pointer) bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("setFrameUsingName:"), name)
	return rv
}
// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system. Can operate on non-resizable windows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setFrameUsingName(_:force:)
func (w_ Window) SetFrameUsingNameForce(name unsafe.Pointer, force bool) bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("setFrameUsingName:force:"), name, force)
	return rv
}
// Sets the window’s miniaturized state to the value you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setIsMiniaturized(_:)
func (w_ Window) SetIsMiniaturized(flag bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setIsMiniaturized:"), flag)
}
// Sets the window’s visible state to the value you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setIsVisible(_:)
func (w_ Window) SetIsVisible(flag bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setIsVisible:"), flag)
}
// Sets the window’s zoomed state to the value you specify. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setIsZoomed(_:)
func (w_ Window) SetIsZoomed(flag bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setIsZoomed:"), flag)
}
// Sets a given path as the window’s title, formatting it as a file-system path, and records this path as the window’s associated file. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/setTitleWithRepresentedFilename(_:)
func (w_ Window) SetTitleWithRepresentedFilename(filename unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setTitleWithRepresentedFilename:"), filename)
}
// Returns the window button of a given window button kind in the window’s view hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/standardWindowButton(_:)
func (w_ Window) StandardWindowButton(b unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("standardWindowButton:"), b)
	return rv
}
// Takes the window into or out of fullscreen mode, [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toggleFullScreen(_:)
func (w_ Window) ToggleFullScreen(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("toggleFullScreen:"), sender)
}
// Shows or hides the tab bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toggleTabBar(_:)
func (w_ Window) ToggleTabBar(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("toggleTabBar:"), sender)
}
// Shows or hides the tab overview. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toggleTabOverview(_:)
func (w_ Window) ToggleTabOverview(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("toggleTabOverview:"), sender)
}
// Toggles the visibility of the window’s toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toggleToolbarShown(_:)
func (w_ Window) ToggleToolbarShown(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("toggleToolbarShown:"), sender)
}
// Tracks events that match the specified mask using the specified tracking handler until the tracking handler explicitly terminates tracking. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/trackEvents(matching:timeout:mode:handler:)
func (w_ Window) TrackEventsMatchingMaskTimeoutModeHandler(mask unsafe.Pointer, timeout unsafe.Pointer, mode unsafe.Pointer, trackingHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("trackEventsMatchingMask:timeout:mode:handler:"), mask, timeout, mode, trackingHandler)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/transferWindowSharing(to:completionHandler:)
func (w_ Window) TransferWindowSharingToWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("transferWindowSharingToWindow:completionHandler:"), window, completionHandler)
}
// Dispatches action messages with a given argument. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/tryToPerform(_:with:)
func (w_ Window) TryToPerformWith(action objc.SEL, object objc.ID) bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("tryToPerform:with:"), action, object)
	return rv
}
// Unregisters the window as a possible destination for dragging operations. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/unregisterDraggedTypes()
func (w_ Window) UnregisterDraggedTypes() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("unregisterDraggedTypes"))
}
// Updates the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/update()
func (w_ Window) Update() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("update"))
}
// Updates the constraints based on changes to views in the window since the last layout. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/updateConstraintsIfNeeded()
func (w_ Window) UpdateConstraintsIfNeeded() {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("updateConstraintsIfNeeded"))
}
// Specifies whether the window is to optimize focusing and drawing when displaying its views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/useOptimizedDrawing(_:)
func (w_ Window) UseOptimizedDrawing(flag bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("useOptimizedDrawing:"), flag)
}
// Returns the scale factor applied to the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/userSpaceScaleFactor
func (w_ Window) UserSpaceScaleFactor() float64 {
	rv := objc.Send[float64](w_.ID(), objc.RegisterName("userSpaceScaleFactor"))
	return rv
}
// Searches for an object that responds to a Services request. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/validRequestor(forSendType:returnType:)
func (w_ Window) ValidRequestorForSendTypeReturnType(sendType unsafe.Pointer, returnType unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](w_.ID(), objc.RegisterName("validRequestorForSendType:returnType:"), sendType, returnType)
	return rv
}
// Displays a visual representation of the supplied constraints in the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/visualizeConstraints(_:)
func (w_ Window) VisualizeConstraints(constraints unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("visualizeConstraints:"), constraints)
}
// Toggles the size and location of the window between its standard state (which the application provides as the best size to display the window’s data) and its user state (a new size and location the user may have set by moving or resizing the window). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/zoom(_:)
func (w_ Window) Zoom(sender objc.ID) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("zoom:"), sender)
}
// A Boolean value that indicates whether the window accepts mouse-moved events. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/acceptsMouseMovedEvents
func (w_ Window) AcceptsMouseMovedEvents() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("acceptsMouseMovedEvents"))
	return rv
}
// SetAcceptsMouseMovedEvents sets the value of the acceptsMouseMovedEvents property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/acceptsMouseMovedEvents
func (w_ Window) SetAcceptsMouseMovedEvents(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAcceptsMouseMovedEvents:"), value)
}
// A Boolean value that indicates whether the window allows multithreaded view drawing. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/allowsConcurrentViewDrawing
func (w_ Window) AllowsConcurrentViewDrawing() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("allowsConcurrentViewDrawing"))
	return rv
}
// SetAllowsConcurrentViewDrawing sets the value of the allowsConcurrentViewDrawing property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/allowsConcurrentViewDrawing
func (w_ Window) SetAllowsConcurrentViewDrawing(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAllowsConcurrentViewDrawing:"), value)
}
// A Boolean value that indicates whether the window can display tooltips even when the application is in the background. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/allowsToolTipsWhenApplicationIsInactive
func (w_ Window) AllowsToolTipsWhenApplicationIsInactive() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("allowsToolTipsWhenApplicationIsInactive"))
	return rv
}
// SetAllowsToolTipsWhenApplicationIsInactive sets the value of the allowsToolTipsWhenApplicationIsInactive property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/allowsToolTipsWhenApplicationIsInactive
func (w_ Window) SetAllowsToolTipsWhenApplicationIsInactive(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAllowsToolTipsWhenApplicationIsInactive:"), value)
}
// The window’s alpha value. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/alphaValue
func (w_ Window) AlphaValue() float64 {
	rv := objc.Send[float64](w_.ID(), objc.RegisterName("alphaValue"))
	return rv
}
// SetAlphaValue sets the value of the alphaValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/alphaValue
func (w_ Window) SetAlphaValue(value float64) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAlphaValue:"), value)
}
// The window’s automatic animation behavior. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/animationBehavior-swift.property
func (w_ Window) AnimationBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("animationBehavior"))
	return rv
}
// SetAnimationBehavior sets the value of the animationBehavior property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/animationBehavior-swift.property
func (w_ Window) SetAnimationBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAnimationBehavior:"), value)
}
// An object that the window inherits its appearance from. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/appearanceSource
func (w_ Window) AppearanceSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("appearanceSource"))
	return rv
}
// SetAppearanceSource sets the value of the appearanceSource property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/appearanceSource
func (w_ Window) SetAppearanceSource(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAppearanceSource:"), value)
}
// A Boolean value that indicates whether the window’s cursor rectangles are enabled. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/areCursorRectsEnabled
func (w_ Window) AreCursorRectsEnabled() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("areCursorRectsEnabled"))
	return rv
}
// The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/aspectRatio
func (w_ Window) AspectRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("aspectRatio"))
	return rv
}
// SetAspectRatio sets the value of the aspectRatio property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/aspectRatio
func (w_ Window) SetAspectRatio(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAspectRatio:"), value)
}
// The sheet attached to the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/attachedSheet
func (w_ Window) AttachedSheet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("attachedSheet"))
	return rv
}
// A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/autorecalculatesKeyViewLoop
func (w_ Window) AutorecalculatesKeyViewLoop() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("autorecalculatesKeyViewLoop"))
	return rv
}
// SetAutorecalculatesKeyViewLoop sets the value of the autorecalculatesKeyViewLoop property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/autorecalculatesKeyViewLoop
func (w_ Window) SetAutorecalculatesKeyViewLoop(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAutorecalculatesKeyViewLoop:"), value)
}
// The color of the window’s background. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/backgroundColor
func (w_ Window) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("backgroundColor"))
	return rv
}
// SetBackgroundColor sets the value of the backgroundColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/backgroundColor
func (w_ Window) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setBackgroundColor:"), value)
}
// The location of the window’s backing store. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/backingLocation-swift.property
func (w_ Window) BackingLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("backingLocation"))
	return rv
}
// The backing scale factor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/backingScaleFactor
func (w_ Window) BackingScaleFactor() float64 {
	rv := objc.Send[float64](w_.ID(), objc.RegisterName("backingScaleFactor"))
	return rv
}
// The window’s backing store type. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/backingType
func (w_ Window) BackingType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("backingType"))
	return rv
}
// SetBackingType sets the value of the backingType property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/backingType
func (w_ Window) SetBackingType(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setBackingType:"), value)
}
// A Boolean value that indicates whether the window can become the key window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/canBecomeKey
func (w_ Window) CanBecomeKeyWindow() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("canBecomeKeyWindow"))
	return rv
}
// A Boolean value that indicates whether the window can become the application’s main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/canBecomeMain
func (w_ Window) CanBecomeMainWindow() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("canBecomeMainWindow"))
	return rv
}
// A Boolean value that indicates whether the window can be displayed at the login window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/canBecomeVisibleWithoutLogin
func (w_ Window) CanBecomeVisibleWithoutLogin() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("canBecomeVisibleWithoutLogin"))
	return rv
}
// SetCanBecomeVisibleWithoutLogin sets the value of the canBecomeVisibleWithoutLogin property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/canBecomeVisibleWithoutLogin
func (w_ Window) SetCanBecomeVisibleWithoutLogin(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setCanBecomeVisibleWithoutLogin:"), value)
}
// A Boolean value that indicates whether the window can hide when its application becomes hidden. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/canHide
func (w_ Window) CanHide() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("canHide"))
	return rv
}
// SetCanHide sets the value of the canHide property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/canHide
func (w_ Window) SetCanHide(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setCanHide:"), value)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/cascadingReferenceFrame
func (w_ Window) CascadingReferenceFrame() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("cascadingReferenceFrame"))
	return rv
}
// An array of the window’s attached child windows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/childWindows
func (w_ Window) ChildWindows() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("childWindows"))
	return rv
}
// A value that identifies the window’s behavior in window collections. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/collectionBehavior-swift.property
func (w_ Window) CollectionBehavior() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("collectionBehavior"))
	return rv
}
// SetCollectionBehavior sets the value of the collectionBehavior property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/collectionBehavior-swift.property
func (w_ Window) SetCollectionBehavior(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setCollectionBehavior:"), value)
}
// The window’s color space. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/colorSpace
func (w_ Window) ColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("colorSpace"))
	return rv
}
// SetColorSpace sets the value of the colorSpace property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/colorSpace
func (w_ Window) SetColorSpace(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setColorSpace:"), value)
}
// The window’s content aspect ratio. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentAspectRatio
func (w_ Window) ContentAspectRatio() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("contentAspectRatio"))
	return rv
}
// SetContentAspectRatio sets the value of the contentAspectRatio property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentAspectRatio
func (w_ Window) SetContentAspectRatio(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setContentAspectRatio:"), value)
}
// A value used by Auto Layout constraints to automatically bind to the value of  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentLayoutGuide
func (w_ Window) ContentLayoutGuide() objc.ID {
	rv := objc.Send[objc.ID](w_.ID(), objc.RegisterName("contentLayoutGuide"))
	return rv
}
// The area inside the window that is for non-obscured content, in window coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentLayoutRect
func (w_ Window) ContentLayoutRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("contentLayoutRect"))
	return rv
}
// The maximum size of the window’s content view in the window’s base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentMaxSize
func (w_ Window) ContentMaxSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("contentMaxSize"))
	return rv
}
// SetContentMaxSize sets the value of the contentMaxSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentMaxSize
func (w_ Window) SetContentMaxSize(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setContentMaxSize:"), value)
}
// The minimum size of the window’s content view in the window’s base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentMinSize
func (w_ Window) ContentMinSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("contentMinSize"))
	return rv
}
// SetContentMinSize sets the value of the contentMinSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentMinSize
func (w_ Window) SetContentMinSize(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setContentMinSize:"), value)
}
// The window’s content-view resizing increments. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentResizeIncrements
func (w_ Window) ContentResizeIncrements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("contentResizeIncrements"))
	return rv
}
// SetContentResizeIncrements sets the value of the contentResizeIncrements property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentResizeIncrements
func (w_ Window) SetContentResizeIncrements(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setContentResizeIncrements:"), value)
}
// The window’s content view, the highest accessible view object in the window’s view hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentView
func (w_ Window) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("contentView"))
	return rv
}
// SetContentView sets the value of the contentView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentView
func (w_ Window) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setContentView:"), value)
}
// The main content view controller for the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentViewController
func (w_ Window) ContentViewController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("contentViewController"))
	return rv
}
// SetContentViewController sets the value of the contentViewController property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/contentViewController
func (w_ Window) SetContentViewController(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setContentViewController:"), value)
}
// The event currently being processed by the application. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/currentEvent
func (w_ Window) CurrentEvent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("currentEvent"))
	return rv
}
// The deepest screen the window is on (it may be split over several screens). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/deepestScreen
func (w_ Window) DeepestScreen() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("deepestScreen"))
	return rv
}
// The button cell that performs as if clicked when the window receives a Return (or Enter) key event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/defaultButtonCell
func (w_ Window) DefaultButtonCell() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("defaultButtonCell"))
	return rv
}
// SetDefaultButtonCell sets the value of the defaultButtonCell property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/defaultButtonCell
func (w_ Window) SetDefaultButtonCell(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setDefaultButtonCell:"), value)
}
// The window’s delegate. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/delegate
func (w_ Window) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("delegate"))
	return rv
}
// SetDelegate sets the value of the delegate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/delegate
func (w_ Window) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setDelegate:"), value)
}
// The depth limit of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/depthLimit
func (w_ Window) DepthLimit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("depthLimit"))
	return rv
}
// SetDepthLimit sets the value of the depthLimit property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/depthLimit
func (w_ Window) SetDepthLimit(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setDepthLimit:"), value)
}
// A dictionary containing information about the window’s resolution, such as color, depth, and so on. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/deviceDescription
func (w_ Window) DeviceDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("deviceDescription"))
	return rv
}
// A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/displaysWhenScreenProfileChanges
func (w_ Window) DisplaysWhenScreenProfileChanges() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("displaysWhenScreenProfileChanges"))
	return rv
}
// SetDisplaysWhenScreenProfileChanges sets the value of the displaysWhenScreenProfileChanges property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/displaysWhenScreenProfileChanges
func (w_ Window) SetDisplaysWhenScreenProfileChanges(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setDisplaysWhenScreenProfileChanges:"), value)
}
// The application’s Dock tile. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/dockTile
func (w_ Window) DockTile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("dockTile"))
	return rv
}
// The collection of drawers associated with the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/drawers
func (w_ Window) Drawers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("drawers"))
	return rv
}
// The window’s first responder. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/firstResponder
func (w_ Window) FirstResponder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("firstResponder"))
	return rv
}
// The window’s frame rectangle in screen coordinates, including the title bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/frame
func (w_ Window) Frame() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("frame"))
	return rv
}
// The name used to automatically save the window’s frame rectangle data in the defaults system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/frameAutosaveName-swift.property
func (w_ Window) FrameAutosaveName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("frameAutosaveName"))
	return rv
}
// A string representation of the window’s frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/frameDescriptor
func (w_ Window) StringWithSavedFrame() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("stringWithSavedFrame"))
	return rv
}
// The graphics context associated with the window for the current thread. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/graphicsContext
func (w_ Window) GraphicsContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("graphicsContext"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/hasActiveWindowSharingSession
func (w_ Window) HasActiveWindowSharingSession() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("hasActiveWindowSharingSession"))
	return rv
}
// A Boolean value that indicates if the window has a close box. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/hasCloseBox
func (w_ Window) HasCloseBox() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("hasCloseBox"))
	return rv
}
// A Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/hasDynamicDepthLimit
func (w_ Window) HasDynamicDepthLimit() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("hasDynamicDepthLimit"))
	return rv
}
// A Boolean value that indicates whether the window has a shadow. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/hasShadow
func (w_ Window) HasShadow() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("hasShadow"))
	return rv
}
// SetHasShadow sets the value of the hasShadow property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/hasShadow
func (w_ Window) SetHasShadow(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setHasShadow:"), value)
}
// A Boolean value that indicates if the window has a title bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/hasTitleBar
func (w_ Window) HasTitleBar() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("hasTitleBar"))
	return rv
}
// A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/hidesOnDeactivate
func (w_ Window) HidesOnDeactivate() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("hidesOnDeactivate"))
	return rv
}
// SetHidesOnDeactivate sets the value of the hidesOnDeactivate property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/hidesOnDeactivate
func (w_ Window) SetHidesOnDeactivate(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setHidesOnDeactivate:"), value)
}
// A Boolean value that indicates whether the window is transparent to mouse events. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/ignoresMouseEvents
func (w_ Window) IgnoresMouseEvents() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("ignoresMouseEvents"))
	return rv
}
// SetIgnoresMouseEvents sets the value of the ignoresMouseEvents property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/ignoresMouseEvents
func (w_ Window) SetIgnoresMouseEvents(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setIgnoresMouseEvents:"), value)
}
// A Boolean value that indicates whether the window is being resized by the user. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/inLiveResize
func (w_ Window) InLiveResize() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("inLiveResize"))
	return rv
}
// The view that’s made first responder (also called the key view) the first time the window is placed onscreen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/initialFirstResponder
func (w_ Window) InitialFirstResponder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("initialFirstResponder"))
	return rv
}
// SetInitialFirstResponder sets the value of the initialFirstResponder property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/initialFirstResponder
func (w_ Window) SetInitialFirstResponder(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setInitialFirstResponder:"), value)
}
// A Boolean value that indicates whether the window automatically displays views that need to be displayed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isAutodisplay
func (w_ Window) Autodisplay() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("autodisplay"))
	return rv
}
// SetAutodisplay sets the value of the autodisplay property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isAutodisplay
func (w_ Window) SetAutodisplay(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setAutodisplay:"), value)
}
// A Boolean value that indicates whether the window’s document has been edited. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isDocumentEdited
func (w_ Window) DocumentEdited() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("documentEdited"))
	return rv
}
// SetDocumentEdited sets the value of the documentEdited property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isDocumentEdited
func (w_ Window) SetDocumentEdited(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setDocumentEdited:"), value)
}
// A Boolean value that indicates whether the window is excluded from the application’s Windows menu. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isExcludedFromWindowsMenu
func (w_ Window) ExcludedFromWindowsMenu() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("excludedFromWindowsMenu"))
	return rv
}
// SetExcludedFromWindowsMenu sets the value of the excludedFromWindowsMenu property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isExcludedFromWindowsMenu
func (w_ Window) SetExcludedFromWindowsMenu(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setExcludedFromWindowsMenu:"), value)
}
// A Boolean value that indicates whether the window is a floating panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isFloatingPanel
func (w_ Window) FloatingPanel() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("floatingPanel"))
	return rv
}
// A Boolean value that indicates whether the window’s flushing ability is disabled. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isFlushWindowDisabled
func (w_ Window) FlushWindowDisabled() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("flushWindowDisabled"))
	return rv
}
// A Boolean value that indicates whether the window is the key window for the application. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isKeyWindow
func (w_ Window) KeyWindow() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("keyWindow"))
	return rv
}
// A Boolean value that indicates whether the window is the application’s main window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isMainWindow
func (w_ Window) MainWindow() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("mainWindow"))
	return rv
}
// A Boolean value that indicates whether the window can minimize. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isMiniaturizable
func (w_ Window) Miniaturizable() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("miniaturizable"))
	return rv
}
// A Boolean value that indicates whether the window is minimized. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isMiniaturized
func (w_ Window) Miniaturized() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("miniaturized"))
	return rv
}
// A Boolean value that indicates whether the window is a modal panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isModalPanel
func (w_ Window) ModalPanel() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("modalPanel"))
	return rv
}
// A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isMovable
func (w_ Window) Movable() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("movable"))
	return rv
}
// SetMovable sets the value of the movable property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isMovable
func (w_ Window) SetMovable(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setMovable:"), value)
}
// A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isMovableByWindowBackground
func (w_ Window) MovableByWindowBackground() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("movableByWindowBackground"))
	return rv
}
// SetMovableByWindowBackground sets the value of the movableByWindowBackground property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isMovableByWindowBackground
func (w_ Window) SetMovableByWindowBackground(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setMovableByWindowBackground:"), value)
}
// A Boolean value that indicates whether the window is on the currently active space. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isOnActiveSpace
func (w_ Window) OnActiveSpace() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("onActiveSpace"))
	return rv
}
// A Boolean value that indicates whether the window device the window manages is freed when it’s removed from the screen list. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isOneShot
func (w_ Window) OneShot() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("oneShot"))
	return rv
}
// SetOneShot sets the value of the oneShot property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isOneShot
func (w_ Window) SetOneShot(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setOneShot:"), value)
}
// A Boolean value that indicates whether the window is opaque. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isOpaque
func (w_ Window) Opaque() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("opaque"))
	return rv
}
// SetOpaque sets the value of the opaque property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isOpaque
func (w_ Window) SetOpaque(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setOpaque:"), value)
}
// A Boolean value that indicates whether the window is released when it receives the   message. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isReleasedWhenClosed
func (w_ Window) ReleasedWhenClosed() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("releasedWhenClosed"))
	return rv
}
// SetReleasedWhenClosed sets the value of the releasedWhenClosed property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isReleasedWhenClosed
func (w_ Window) SetReleasedWhenClosed(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setReleasedWhenClosed:"), value)
}
// A Boolean value that indicates if the user can resize the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isResizable
func (w_ Window) Resizable() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("resizable"))
	return rv
}
// A Boolean value indicating whether the window configuration is preserved between application launches. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isRestorable
func (w_ Window) Restorable() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("restorable"))
	return rv
}
// SetRestorable sets the value of the restorable property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isRestorable
func (w_ Window) SetRestorable(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setRestorable:"), value)
}
// A Boolean value that indicates whether the window has ever run as a modal sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isSheet
func (w_ Window) Sheet() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("sheet"))
	return rv
}
// A Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isVisible
func (w_ Window) Visible() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("visible"))
	return rv
}
// A Boolean value that indicates whether the window allows zooming. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isZoomable
func (w_ Window) Zoomable() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("zoomable"))
	return rv
}
// A Boolean value that indicates whether the window is in a zoomed state. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/isZoomed
func (w_ Window) Zoomed() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("zoomed"))
	return rv
}
// The direction the window is currently using to change the key view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/keyViewSelectionDirection
func (w_ Window) KeyViewSelectionDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("keyViewSelectionDirection"))
	return rv
}
// The window level of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/level-swift.property
func (w_ Window) Level() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("level"))
	return rv
}
// SetLevel sets the value of the level property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/level-swift.property
func (w_ Window) SetLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setLevel:"), value)
}
// A maximum size that is used to determine if a window can fit when it is in full screen in a tile. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/maxFullScreenContentSize
func (w_ Window) MaxFullScreenContentSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("maxFullScreenContentSize"))
	return rv
}
// SetMaxFullScreenContentSize sets the value of the maxFullScreenContentSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/maxFullScreenContentSize
func (w_ Window) SetMaxFullScreenContentSize(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setMaxFullScreenContentSize:"), value)
}
// The maximum size to which the window’s frame (including its title bar) can be sized. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/maxSize
func (w_ Window) MaxSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("maxSize"))
	return rv
}
// SetMaxSize sets the value of the maxSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/maxSize
func (w_ Window) SetMaxSize(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setMaxSize:"), value)
}
// A minimum size that is used to determine if a window can fit when it is in full screen in a tile. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/minFullScreenContentSize
func (w_ Window) MinFullScreenContentSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("minFullScreenContentSize"))
	return rv
}
// SetMinFullScreenContentSize sets the value of the minFullScreenContentSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/minFullScreenContentSize
func (w_ Window) SetMinFullScreenContentSize(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setMinFullScreenContentSize:"), value)
}
// The minimum size to which the window’s frame (including its title bar) can be sized. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/minSize
func (w_ Window) MinSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("minSize"))
	return rv
}
// SetMinSize sets the value of the minSize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/minSize
func (w_ Window) SetMinSize(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setMinSize:"), value)
}
// The custom miniaturized window image of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/miniwindowImage
func (w_ Window) MiniwindowImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("miniwindowImage"))
	return rv
}
// SetMiniwindowImage sets the value of the miniwindowImage property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/miniwindowImage
func (w_ Window) SetMiniwindowImage(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setMiniwindowImage:"), value)
}
// The title displayed in the window’s minimized window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/miniwindowTitle
func (w_ Window) MiniwindowTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("miniwindowTitle"))
	return rv
}
// SetMiniwindowTitle sets the value of the miniwindowTitle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/miniwindowTitle
func (w_ Window) SetMiniwindowTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setMiniwindowTitle:"), value)
}
// The current location of the pointer reckoned in the window’s base coordinate system, regardless of the current event being handled or of any events pending. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/mouseLocationOutsideOfEventStream
func (w_ Window) MouseLocationOutsideOfEventStream() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("mouseLocationOutsideOfEventStream"))
	return rv
}
// The occlusion state of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/occlusionState-swift.property
func (w_ Window) OcclusionState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("occlusionState"))
	return rv
}
// The zero-based position of the window, based on its order from front to back among all visible application windows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/orderedIndex
func (w_ Window) OrderedIndex() int {
	rv := objc.Send[int](w_.ID(), objc.RegisterName("orderedIndex"))
	return rv
}
// SetOrderedIndex sets the value of the orderedIndex property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/orderedIndex
func (w_ Window) SetOrderedIndex(value int) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setOrderedIndex:"), value)
}
// The parent window to which the window is attached as a child. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/parent
func (w_ Window) ParentWindow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("parentWindow"))
	return rv
}
// SetParentWindow sets the value of the parentWindow property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/parent
func (w_ Window) SetParentWindow(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setParentWindow:"), value)
}
// A Boolean value that indicates the preferred location for the window’s backing store. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/preferredBackingLocation
func (w_ Window) PreferredBackingLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("preferredBackingLocation"))
	return rv
}
// SetPreferredBackingLocation sets the value of the preferredBackingLocation property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/preferredBackingLocation
func (w_ Window) SetPreferredBackingLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setPreferredBackingLocation:"), value)
}
// A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/preservesContentDuringLiveResize
func (w_ Window) PreservesContentDuringLiveResize() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("preservesContentDuringLiveResize"))
	return rv
}
// SetPreservesContentDuringLiveResize sets the value of the preservesContentDuringLiveResize property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/preservesContentDuringLiveResize
func (w_ Window) SetPreservesContentDuringLiveResize(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setPreservesContentDuringLiveResize:"), value)
}
// A Boolean value that indicates whether the window prevents application termination when modal. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/preventsApplicationTerminationWhenModal
func (w_ Window) PreventsApplicationTerminationWhenModal() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("preventsApplicationTerminationWhenModal"))
	return rv
}
// SetPreventsApplicationTerminationWhenModal sets the value of the preventsApplicationTerminationWhenModal property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/preventsApplicationTerminationWhenModal
func (w_ Window) SetPreventsApplicationTerminationWhenModal(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setPreventsApplicationTerminationWhenModal:"), value)
}
// The path to the file of the window’s represented file. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/representedFilename
func (w_ Window) RepresentedFilename() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("representedFilename"))
	return rv
}
// SetRepresentedFilename sets the value of the representedFilename property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/representedFilename
func (w_ Window) SetRepresentedFilename(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setRepresentedFilename:"), value)
}
// The URL of the file the window represents. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/representedURL
func (w_ Window) RepresentedURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("representedURL"))
	return rv
}
// SetRepresentedURL sets the value of the representedURL property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/representedURL
func (w_ Window) SetRepresentedURL(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setRepresentedURL:"), value)
}
// The flags field of the event record for the mouse-down event that initiated the resizing session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/resizeFlags
func (w_ Window) ResizeFlags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("resizeFlags"))
	return rv
}
// The window’s resizing increments. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/resizeIncrements
func (w_ Window) ResizeIncrements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("resizeIncrements"))
	return rv
}
// SetResizeIncrements sets the value of the resizeIncrements property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/resizeIncrements
func (w_ Window) SetResizeIncrements(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setResizeIncrements:"), value)
}
// The restoration class associated with the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/restorationClass
func (w_ Window) RestorationClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("restorationClass"))
	return rv
}
// SetRestorationClass sets the value of the restorationClass property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/restorationClass
func (w_ Window) SetRestorationClass(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setRestorationClass:"), value)
}
// The screen the window is on. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/screen
func (w_ Window) Screen() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("screen"))
	return rv
}
// A Boolean value that indicates the level of access other processes have to the window’s content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/sharingType-swift.property
func (w_ Window) SharingType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("sharingType"))
	return rv
}
// SetSharingType sets the value of the sharingType property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/sharingType-swift.property
func (w_ Window) SetSharingType(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setSharingType:"), value)
}
// The window to which the sheet is attached. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/sheetParent
func (w_ Window) SheetParent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("sheetParent"))
	return rv
}
// An array of the sheets currently attached to the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/sheets
func (w_ Window) Sheets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("sheets"))
	return rv
}
// A Boolean value that indicates whether the window’s resize indicator is visible. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/showsResizeIndicator
func (w_ Window) ShowsResizeIndicator() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("showsResizeIndicator"))
	return rv
}
// SetShowsResizeIndicator sets the value of the showsResizeIndicator property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/showsResizeIndicator
func (w_ Window) SetShowsResizeIndicator(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setShowsResizeIndicator:"), value)
}
// A Boolean value that indicates whether the toolbar control button is currently displayed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/showsToolbarButton
func (w_ Window) ShowsToolbarButton() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("showsToolbarButton"))
	return rv
}
// SetShowsToolbarButton sets the value of the showsToolbarButton property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/showsToolbarButton
func (w_ Window) SetShowsToolbarButton(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setShowsToolbarButton:"), value)
}
// Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/styleMask-swift.property
func (w_ Window) StyleMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("styleMask"))
	return rv
}
// SetStyleMask sets the value of the styleMask property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/styleMask-swift.property
func (w_ Window) SetStyleMask(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setStyleMask:"), value)
}
// A secondary line of text that appears in the title bar of the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/subtitle
func (w_ Window) Subtitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("subtitle"))
	return rv
}
// SetSubtitle sets the value of the subtitle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/subtitle
func (w_ Window) SetSubtitle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setSubtitle:"), value)
}
// An object that represents information about a window when it displays as a tab. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/tab
func (w_ Window) Tab() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("tab"))
	return rv
}
// A group of windows that display together as a tab group. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/tabGroup
func (w_ Window) TabGroup() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("tabGroup"))
	return rv
}
// An array of windows that display as tabs. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/tabbedWindows
func (w_ Window) TabbedWindows() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("tabbedWindows"))
	return rv
}
// A value that allows a group of related windows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/tabbingIdentifier-swift.property
func (w_ Window) TabbingIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("tabbingIdentifier"))
	return rv
}
// SetTabbingIdentifier sets the value of the tabbingIdentifier property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/tabbingIdentifier-swift.property
func (w_ Window) SetTabbingIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setTabbingIdentifier:"), value)
}
// A value that indicates when a window displays tabs. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/tabbingMode-swift.property
func (w_ Window) TabbingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("tabbingMode"))
	return rv
}
// SetTabbingMode sets the value of the tabbingMode property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/tabbingMode-swift.property
func (w_ Window) SetTabbingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setTabbingMode:"), value)
}
// The string that appears in the title bar of the window or the path to the represented file. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/title
func (w_ Window) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("title"))
	return rv
}
// SetTitle sets the value of the title property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/title
func (w_ Window) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setTitle:"), value)
}
// A value that indicates the visibility of the window’s title and title bar buttons. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/titleVisibility-swift.property
func (w_ Window) TitleVisibility() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("titleVisibility"))
	return rv
}
// SetTitleVisibility sets the value of the titleVisibility property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/titleVisibility-swift.property
func (w_ Window) SetTitleVisibility(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setTitleVisibility:"), value)
}
// An array of title bar accessory view controllers that are currently added to the window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/titlebarAccessoryViewControllers
func (w_ Window) TitlebarAccessoryViewControllers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("titlebarAccessoryViewControllers"))
	return rv
}
// SetTitlebarAccessoryViewControllers sets the value of the titlebarAccessoryViewControllers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/titlebarAccessoryViewControllers
func (w_ Window) SetTitlebarAccessoryViewControllers(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setTitlebarAccessoryViewControllers:"), value)
}
// A Boolean value that indicates whether the title bar draws its background. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/titlebarAppearsTransparent
func (w_ Window) TitlebarAppearsTransparent() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("titlebarAppearsTransparent"))
	return rv
}
// SetTitlebarAppearsTransparent sets the value of the titlebarAppearsTransparent property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/titlebarAppearsTransparent
func (w_ Window) SetTitlebarAppearsTransparent(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setTitlebarAppearsTransparent:"), value)
}
// The type of separator that the app displays between the title bar and content of a window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/titlebarSeparatorStyle
func (w_ Window) TitlebarSeparatorStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("titlebarSeparatorStyle"))
	return rv
}
// SetTitlebarSeparatorStyle sets the value of the titlebarSeparatorStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/titlebarSeparatorStyle
func (w_ Window) SetTitlebarSeparatorStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setTitlebarSeparatorStyle:"), value)
}
// The window’s toolbar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toolbar
func (w_ Window) Toolbar() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("toolbar"))
	return rv
}
// SetToolbar sets the value of the toolbar property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toolbar
func (w_ Window) SetToolbar(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setToolbar:"), value)
}
// The style that determines the appearance and location of the toolbar in relation to the title bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toolbarStyle-swift.property
func (w_ Window) ToolbarStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("toolbarStyle"))
	return rv
}
// SetToolbarStyle sets the value of the toolbarStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/toolbarStyle-swift.property
func (w_ Window) SetToolbarStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setToolbarStyle:"), value)
}
// A Boolean value that indicates whether any of the window’s views need to be displayed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/viewsNeedDisplay
func (w_ Window) ViewsNeedDisplay() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("viewsNeedDisplay"))
	return rv
}
// SetViewsNeedDisplay sets the value of the viewsNeedDisplay property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/viewsNeedDisplay
func (w_ Window) SetViewsNeedDisplay(value bool) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setViewsNeedDisplay:"), value)
}
// The window’s window controller. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/windowController
func (w_ Window) WindowController() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("windowController"))
	return rv
}
// SetWindowController sets the value of the windowController property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/windowController
func (w_ Window) SetWindowController(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID(), objc.RegisterName("setWindowController:"), value)
}
// The window number of the window’s window device. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/windowNumber
func (w_ Window) WindowNumber() int {
	rv := objc.Send[int](w_.ID(), objc.RegisterName("windowNumber"))
	return rv
}
// The direction the window’s title bar lays text out, either left to right or right to left. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/windowTitlebarLayoutDirection
func (w_ Window) WindowTitlebarLayoutDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID(), objc.RegisterName("windowTitlebarLayoutDirection"))
	return rv
}
// A Boolean value that indicates whether the window is able to receive keyboard and mouse events even when some other window is being run modally. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSWindow/worksWhenModal
func (w_ Window) WorksWhenModal() bool {
	rv := objc.Send[bool](w_.ID(), objc.RegisterName("worksWhenModal"))
	return rv
}
