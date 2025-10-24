// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/quartzcore"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSWindow */


/* debug [class_header]: Header for NSWindow */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Window */
// An interface definition for the [Window] class.
type IWindow interface {
	IResponder
	
/* debug [class_interface_properties]: Properties for Window */
	// properties:
	AcceptsMouseMovedEvents() bool
	SetAcceptsMouseMovedEvents(value bool)
	AllowsConcurrentViewDrawing() bool
	SetAllowsConcurrentViewDrawing(value bool)
	AlphaValue() float64
	SetAlphaValue(value float64)
	AnimationBehavior() WindowAnimationBehavior
	SetAnimationBehavior(value WindowAnimationBehavior)
	AreCursorRectsEnabled() bool
	AspectRatio() Size /* not a class type */
	SetAspectRatio(value Size /* not a class type */)
	AttachedSheet() IWindow
	AutorecalculatesKeyViewLoop() bool
	SetAutorecalculatesKeyViewLoop(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BackingLocation() WindowBackingLocation
	BackingScaleFactor() float64
	BackingType() BackingStoreType
	SetBackingType(value BackingStoreType)
	CanBecomeKeyWindow() bool
	CanBecomeMainWindow() bool
	CanBecomeVisibleWithoutLogin() bool
	SetCanBecomeVisibleWithoutLogin(value bool)
	CanHide() bool
	SetCanHide(value bool)
	CollectionBehavior() WindowCollectionBehavior
	SetCollectionBehavior(value WindowCollectionBehavior)
	ColorSpace() IColorSpace
	SetColorSpace(value IColorSpace)
	ContentAspectRatio() Size /* not a class type */
	SetContentAspectRatio(value Size /* not a class type */)
	ContentLayoutGuide() objc.ID
	ContentLayoutRect() Rect /* not a class type */
	ContentMaxSize() Size /* not a class type */
	SetContentMaxSize(value Size /* not a class type */)
	ContentMinSize() Size /* not a class type */
	SetContentMinSize(value Size /* not a class type */)
	ContentResizeIncrements() Size /* not a class type */
	SetContentResizeIncrements(value Size /* not a class type */)
	ContentView() IView
	SetContentView(value IView)
	ContentViewController() IViewController
	SetContentViewController(value IViewController)
	CurrentEvent() IEvent
	DeepestScreen() IScreen
	DepthLimit() WindowDepth
	SetDepthLimit(value WindowDepth)
	DeviceDescription() foundation.IDictionary
	DisplaysWhenScreenProfileChanges() bool
	SetDisplaysWhenScreenProfileChanges(value bool)
	Drawers() []Drawer
	FirstResponder() IResponder
	Frame() Rect /* not a class type */
	FrameAutosaveName() WindowFrameAutosaveName /* typedef */
	StringWithSavedFrame() WindowPersistableFrameDescriptor /* typedef */
	GraphicsContext() IGraphicsContext
	HasDynamicDepthLimit() bool
	HasShadow() bool
	SetHasShadow(value bool)
	HidesOnDeactivate() bool
	SetHidesOnDeactivate(value bool)
	IgnoresMouseEvents() bool
	SetIgnoresMouseEvents(value bool)
	InitialFirstResponder() IView
	SetInitialFirstResponder(value IView)
	Autodisplay() bool
	SetAutodisplay(value bool)
	DocumentEdited() bool
	SetDocumentEdited(value bool)
	ExcludedFromWindowsMenu() bool
	SetExcludedFromWindowsMenu(value bool)
	FlushWindowDisabled() bool
	KeyWindow() bool
	MainWindow() bool
	Miniaturized() bool
	Movable() bool
	SetMovable(value bool)
	MovableByWindowBackground() bool
	SetMovableByWindowBackground(value bool)
	OnActiveSpace() bool
	OneShot() bool
	SetOneShot(value bool)
	Opaque() bool
	SetOpaque(value bool)
	ReleasedWhenClosed() bool
	SetReleasedWhenClosed(value bool)
	Sheet() bool
	Visible() bool
	Zoomed() bool
	KeyViewSelectionDirection() SelectionDirection
	Level() WindowLevel /* typedef */
	SetLevel(value WindowLevel /* typedef */)
	MaxFullScreenContentSize() Size /* not a class type */
	SetMaxFullScreenContentSize(value Size /* not a class type */)
	MaxSize() Size /* not a class type */
	SetMaxSize(value Size /* not a class type */)
	MinFullScreenContentSize() Size /* not a class type */
	SetMinFullScreenContentSize(value Size /* not a class type */)
	MinSize() Size /* not a class type */
	SetMinSize(value Size /* not a class type */)
	MiniwindowImage() IImage
	SetMiniwindowImage(value IImage)
	MiniwindowTitle() objc.IObject /* cross-framework: NSString */
	SetMiniwindowTitle(value objc.IObject /* cross-framework: NSString */)
	MouseLocationOutsideOfEventStream() vision.Point
	OcclusionState() WindowOcclusionState
	PreferredBackingLocation() WindowBackingLocation
	SetPreferredBackingLocation(value WindowBackingLocation)
	PreservesContentDuringLiveResize() bool
	SetPreservesContentDuringLiveResize(value bool)
	RepresentedFilename() objc.IObject /* cross-framework: NSString */
	SetRepresentedFilename(value objc.IObject /* cross-framework: NSString */)
	RepresentedURL() objc.IObject /* cross-framework: NSURL */
	SetRepresentedURL(value objc.IObject /* cross-framework: NSURL */)
	ResizeFlags() EventModifierFlags
	ResizeIncrements() Size /* not a class type */
	SetResizeIncrements(value Size /* not a class type */)
	RestorationClass() unsafe.Pointer
	SetRestorationClass(value unsafe.Pointer)
	Screen() IScreen
	SharingType() WindowSharingType
	SetSharingType(value WindowSharingType)
	SheetParent() IWindow
	Sheets() []Window
	ShowsResizeIndicator() bool
	SetShowsResizeIndicator(value bool)
	StyleMask() WindowStyleMask
	SetStyleMask(value WindowStyleMask)
	Subtitle() objc.IObject /* cross-framework: NSString */
	SetSubtitle(value objc.IObject /* cross-framework: NSString */)
	Tab() IWindowTab
	TabGroup() IWindowTabGroup
	TabbedWindows() []Window
	TabbingIdentifier() WindowTabbingIdentifier /* typedef */
	SetTabbingIdentifier(value WindowTabbingIdentifier /* typedef */)
	TabbingMode() WindowTabbingMode
	SetTabbingMode(value WindowTabbingMode)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	TitleVisibility() WindowTitleVisibility
	SetTitleVisibility(value WindowTitleVisibility)
	ViewsNeedDisplay() bool
	SetViewsNeedDisplay(value bool)
	WindowNumber() int
	WindowRef() objectivec.IObject
	WorksWhenModal() bool
	NumberOfColorComponents() int
	SetNumberOfColorComponents(value int)
	BitsPerPixel() int
	SetBitsPerPixel(value int)
	BitsPerSample() int
	SetBitsPerSample(value int)
	ColorSpaceName() ColorSpaceName /* typedef */
	SetColorSpaceName(value ColorSpaceName /* typedef */)
	IsPlanar() bool
	SetIsPlanar(value bool)
	AllowsToolTipsWhenApplicationIsInactive() bool
	SetAllowsToolTipsWhenApplicationIsInactive(value bool)
	AppearanceSource() AppearanceCustomization /* not a class type */
	SetAppearanceSource(value AppearanceCustomization /* not a class type */)
	CanBecomeKey() bool
	SetCanBecomeKey(value bool)
	CanBecomeMain() bool
	SetCanBecomeMain(value bool)
	CascadingReferenceFrame() Rect /* not a class type */
	SetCascadingReferenceFrame(value Rect /* not a class type */)
	ChildWindows() IWindow
	SetChildWindows(value IWindow)
	DefaultButtonCell() IButtonCell
	SetDefaultButtonCell(value IButtonCell)
	Delegate() objc.IObject /* cross-framework: WindowDelegate */
	SetDelegate(value objc.IObject /* cross-framework: WindowDelegate */)
	DockTile() IDockTile
	SetDockTile(value IDockTile)
	FrameDescriptor() objectivec.IObject
	SetFrameDescriptor(value objectivec.IObject)
	HasActiveWindowSharingSession() bool
	SetHasActiveWindowSharingSession(value bool)
	HasCloseBox() bool
	SetHasCloseBox(value bool)
	HasTitleBar() bool
	SetHasTitleBar(value bool)
	InLiveResize() bool
	SetInLiveResize(value bool)
	IsDocumentEdited() bool
	SetIsDocumentEdited(value bool)
	IsExcludedFromWindowsMenu() bool
	SetIsExcludedFromWindowsMenu(value bool)
	IsFloatingPanel() bool
	SetIsFloatingPanel(value bool)
	IsKeyWindow() bool
	SetIsKeyWindow(value bool)
	IsMainWindow() bool
	SetIsMainWindow(value bool)
	IsMiniaturizable() bool
	SetIsMiniaturizable(value bool)
	IsMiniaturized() bool
	SetIsMiniaturized(value bool)
	IsModalPanel() bool
	SetIsModalPanel(value bool)
	IsMovable() bool
	SetIsMovable(value bool)
	IsMovableByWindowBackground() bool
	SetIsMovableByWindowBackground(value bool)
	IsOnActiveSpace() bool
	SetIsOnActiveSpace(value bool)
	IsOpaque() bool
	SetIsOpaque(value bool)
	IsReleasedWhenClosed() bool
	SetIsReleasedWhenClosed(value bool)
	IsResizable() bool
	SetIsResizable(value bool)
	IsRestorable() bool
	SetIsRestorable(value bool)
	IsSheet() bool
	SetIsSheet(value bool)
	IsVisible() bool
	SetIsVisible(value bool)
	IsZoomable() bool
	SetIsZoomable(value bool)
	IsZoomed() bool
	SetIsZoomed(value bool)
	OrderedIndex() int
	SetOrderedIndex(value int)
	Parent() IWindow
	SetParent(value IWindow)
	PreventsApplicationTerminationWhenModal() bool
	SetPreventsApplicationTerminationWhenModal(value bool)
	ShowsToolbarButton() bool
	SetShowsToolbarButton(value bool)
	TitlebarAccessoryViewControllers() ITitlebarAccessoryViewController
	SetTitlebarAccessoryViewControllers(value ITitlebarAccessoryViewController)
	TitlebarAppearsTransparent() bool
	SetTitlebarAppearsTransparent(value bool)
	TitlebarSeparatorStyle() TitlebarSeparatorStyle
	SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle)
	Toolbar() IToolbar
	SetToolbar(value IToolbar)
	ToolbarStyle() objectivec.IObject
	SetToolbarStyle(value objectivec.IObject)
	WindowController() IWindowController
	SetWindowController(value IWindowController)
	WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection
	SetWindowTitlebarLayoutDirection(value UserInterfaceLayoutDirection)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Window */
	// methods:
	AddTabbedWindowOrdered(window IWindow, ordered WindowOrderingMode)
	AnimationResizeTime(newFrame Rect /* not a class type */) float64
	AutorecalculatesContentBorderThicknessForEdge(edge RectEdge /* not a class type */) bool
	BackingAlignedRectOptions(rect Rect /* not a class type */, options AlignmentOptions /* not a class type */) Rect /* not a class type */
	BecomeKeyWindow()
	BecomeMainWindow()
	BeginCriticalSheetCompletionHandler(sheetWindow IWindow, handler unsafe.Pointer)
	BeginSheetCompletionHandler(sheetWindow IWindow, handler unsafe.Pointer)
	CanRepresentDisplayGamut(displayGamut DisplayGamut) bool
	CascadeTopLeftFromPoint(topLeftPoint vision.Point) vision.Point
	Center()
	Close()
	ConstrainFrameRectToScreen(frameRect Rect /* not a class type */, screen IScreen) Rect /* not a class type */
	ContentRectForFrameRect(frameRect Rect /* not a class type */) Rect /* not a class type */
	ConvertRectFromBacking(rect Rect /* not a class type */) Rect /* not a class type */
	ConvertRectFromScreen(rect Rect /* not a class type */) Rect /* not a class type */
	ConvertPointFromScreen(point vision.Point) vision.Point
	ConvertPointToScreen(point vision.Point) vision.Point
	ConvertPointFromBacking(point vision.Point) vision.Point
	ConvertPointToBacking(point vision.Point) vision.Point
	ConvertRectToBacking(rect Rect /* not a class type */) Rect /* not a class type */
	ConvertRectToScreen(rect Rect /* not a class type */) Rect /* not a class type */
	Deminiaturize(sender objc.IObject)
	DisableCursorRects()
	DiscardCursorRects()
	DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent)
	Display()
	DisplayIfNeeded()
	DisplayLinkWithTargetSelector(target objc.IObject, selector objc.SEL) quartzcore.DisplayLink
	EnableCursorRects()
	EndEditingFor(object objc.IObject)
	EndSheet(sheetWindow IWindow)
	EndSheetReturnCode(sheetWindow IWindow, returnCode ModalResponse /* typedef */)
	FieldEditorForObject(createFlag bool, object objc.IObject) IText
	FrameRectForContentRect(contentRect Rect /* not a class type */) Rect /* not a class type */
	InvalidateCursorRectsForView(view IView)
	InvalidateShadow()
	MakeFirstResponder(responder IResponder) bool
	MakeKeyWindow()
	MakeKeyAndOrderFront(sender objc.IObject)
	MakeMainWindow()
	MergeAllWindows(sender objc.IObject)
	Miniaturize(sender objc.IObject)
	MoveTabToNewWindow(sender objc.IObject)
	NextEventMatchingMask(mask EventMask) IEvent
	NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration objc.IObject /* cross-framework: NSDate */, mode RunLoopMode /* not a class type */, deqFlag bool) IEvent
	OrderWindowRelativeTo(place WindowOrderingMode, otherWin int)
	OrderBack(sender objc.IObject)
	OrderFront(sender objc.IObject)
	OrderFrontRegardless()
	OrderOut(sender objc.IObject)
	PerformClose(sender objc.IObject)
	PerformWindowDragWithEvent(event IEvent)
	PerformMiniaturize(sender objc.IObject)
	PerformZoom(sender objc.IObject)
	PostEventAtStart(event IEvent, flag bool)
	RecalculateKeyViewLoop()
	RegisterForDraggedTypes(newTypes []string)
	ResetCursorRects()
	ResignKeyWindow()
	ResignMainWindow()
	SaveFrameUsingName(name WindowFrameAutosaveName /* typedef */)
	SelectKeyViewFollowingView(view IView)
	SelectKeyViewPrecedingView(view IView)
	SelectNextKeyView(sender objc.IObject)
	SelectNextTab(sender objc.IObject)
	SelectPreviousKeyView(sender objc.IObject)
	SelectPreviousTab(sender objc.IObject)
	SendEvent(event IEvent)
	SetContentSize(size Size /* not a class type */)
	SetDynamicDepthLimit(flag bool)
	SetFrameDisplay(frameRect Rect /* not a class type */, flag bool)
	SetFrameDisplayAnimate(frameRect Rect /* not a class type */, displayFlag bool, animateFlag bool)
	SetFrameFromString(string_ WindowPersistableFrameDescriptor /* typedef */)
	SetFrameOrigin(point vision.Point)
	SetFrameTopLeftPoint(point vision.Point)
	SetFrameUsingName(name WindowFrameAutosaveName /* typedef */) bool
	SetFrameUsingNameForce(name WindowFrameAutosaveName /* typedef */, force bool) bool
	SetTitleWithRepresentedFilename(filename objc.IObject /* cross-framework: NSString */)
	ToggleFullScreen(sender objc.IObject)
	ToggleTabBar(sender objc.IObject)
	ToggleTabOverview(sender objc.IObject)
	TrackEventsMatchingMaskTimeoutModeHandler(mask EventMask, timeout float64, mode RunLoopMode /* not a class type */, trackingHandler unsafe.Pointer)
	TryToPerformWith(action objc.SEL, object objc.IObject) bool
	UnregisterDraggedTypes()
	VisualizeConstraints(constraints []LayoutConstraint)
	Zoom(sender objc.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Window */
// Alloc allocates a new instance without initialization.
func (wc _WindowClass) Alloc() Window {
	rv := objc.Send[Window](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Window */
// A window that an app displays on the screen.
//
// A single object corresponds to, at most, one on-screen window. Windows perform two principal functions: To place views in a provided area To accept and distribute mouse and keyboard events the user generates to the appropriate views


// A window that an app displays on the screen.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Window */

// Initializes the window with the specified values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:)
func NewWindowWithContentRectStyleMaskBackingDefer(contentRect Rect /* not a class type */, style WindowStyleMask, backingStoreType BackingStoreType, flag bool) Window {
	instance := getWindowClass().Alloc()
	rv := objc.Send[Window](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWindowWithContentRectStyleMaskBackingDefer */


// Initializes an allocated window with the specified values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:screen:)
func NewWindowWithContentRectStyleMaskBackingDeferScreen(contentRect Rect /* not a class type */, style WindowStyleMask, backingStoreType BackingStoreType, flag bool, screen IScreen) Window {
	instance := getWindowClass().Alloc()
	rv := objc.Send[Window](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWindowWithContentRectStyleMaskBackingDeferScreen */


// Creates a titled window that contains the specified content view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentViewController:)
func NewWindowWithContentViewController(contentViewController IViewController) Window {
	rv := objc.Send[Window](objc.ID(getWindowClass().class), objc.Sel("windowWithContentViewController:"), contentViewController)
	return rv
}/* debug [class_init_methods/constructor]: NewWindowWithContentViewController */


// Returns a Cocoa window created from a Carbon window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(windowRef:)
func NewWindowWithWindowRef(windowRef objectivec.IObject) Window {
	instance := getWindowClass().Alloc()
	rv := objc.Send[Window](instance.ID, objc.Sel("initWithWindowRef:"), windowRef)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewWindowWithWindowRef */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Window */

// Returns the content rectangle used by a window with a given frame rectangle and window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentRect(forFrameRect:styleMask:)
func (wc _WindowClass) ContentRectForFrameRectStyleMask(fRect Rect /* not a class type */, style WindowStyleMask) Rect /* not a class type */ {
	rv := objc.Send[Rect](objc.ID(wc.class), objc.Sel("contentRectForFrameRect:styleMask:"), fRect, style)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContentRectForFrameRectStyleMask) */


// Returns the frame rectangle used by a window with a given content rectangle and window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frameRect(forContentRect:styleMask:)
func (wc _WindowClass) FrameRectForContentRectStyleMask(cRect Rect /* not a class type */, style WindowStyleMask) Rect /* not a class type */ {
	rv := objc.Send[Rect](objc.ID(wc.class), objc.Sel("frameRectForContentRect:styleMask:"), cRect, style)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FrameRectForContentRectStyleMask) */


// Creates a titled window that contains the specified content view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentViewController:)
func (wc _WindowClass) WindowWithContentViewController(contentViewController IViewController) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(wc.class), objc.Sel("windowWithContentViewController:"), contentViewController)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WindowWithContentViewController) */


// This method does nothing; it is here for backward compatibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/menuChanged(_:)
func (wc _WindowClass) MenuChanged(menu IMenu) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("menuChanged:"), menu)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MenuChanged) */


// Returns the minimum width a window’s frame rectangle must have for it to display a title, with a given window style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/minFrameWidth(withTitle:styleMask:)
func (wc _WindowClass) MinFrameWidthWithTitleStyleMask(title objc.IObject /* cross-framework: NSString */, style WindowStyleMask) float64 {
	rv := objc.Send[float64](objc.ID(wc.class), objc.Sel("minFrameWidthWithTitle:styleMask:"), title, style)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MinFrameWidthWithTitleStyleMask) */


// Removes the frame data stored under a given name from the application’s user defaults.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/removeFrame(usingName:)
func (wc _WindowClass) RemoveFrameUsingName(name WindowFrameAutosaveName /* typedef */) {
	objc.Send[objc.ID](objc.ID(wc.class), objc.Sel("removeFrameUsingName:"), name)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RemoveFrameUsingName) */


// Returns the number of the frontmost window that would be hit by a mouse-down at the specified screen location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumber(at:belowWindowWithWindowNumber:)
func (wc _WindowClass) WindowNumberAtPointBelowWindowWithWindowNumber(point vision.Point, windowNumber int) int {
	rv := objc.Send[int](objc.ID(wc.class), objc.Sel("windowNumberAtPoint:belowWindowWithWindowNumber:"), point, windowNumber)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WindowNumberAtPointBelowWindowWithWindowNumber) */


// Returns the window numbers for all visible windows satisfying the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumbers(options:)
func (wc _WindowClass) WindowNumbersWithOptions(options WindowNumberListOptions) []foundation.Number {
	rv := objc.Send[[]foundation.Number](objc.ID(wc.class), objc.Sel("windowNumbersWithOptions:"), options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WindowNumbersWithOptions) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Window */

// A Boolean value that indicates whether the app can automatically organize windows into tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsAutomaticWindowTabbing
func (wc _WindowClass) AllowsAutomaticWindowTabbing() bool {
	rv := objc.Send[bool](objc.ID(wc.class), objc.Sel("allowsAutomaticWindowTabbing"))
	return rv
}/* debug [class_properties_class/property]: allowsAutomaticWindowTabbing */

// Returns the default depth limit for instances of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/defaultDepthLimit
func (wc _WindowClass) DefaultDepthLimit() WindowDepth {
	rv := objc.Send[WindowDepth](objc.ID(wc.class), objc.Sel("defaultDepthLimit"))
	return rv
}/* debug [class_properties_class/property]: defaultDepthLimit */

// A value that indicates the user’s preference for window tabbing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/userTabbingPreference-swift.type.property
func (wc _WindowClass) UserTabbingPreference() WindowUserTabbingPreference {
	rv := objc.Send[WindowUserTabbingPreference](objc.ID(wc.class), objc.Sel("userTabbingPreference"))
	return rv
}/* debug [class_properties_class/property]: userTabbingPreference */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Window */

// Adds the provided window as a new tab in a tabbed window using the specified ordering instruction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/addTabbedWindow(_:ordered:)
func (w_ Window) AddTabbedWindowOrdered(window IWindow, ordered WindowOrderingMode) {
	objc.Send[objc.ID](w_.ID, objc.Sel("addTabbedWindow:ordered:"), window, ordered)
}/* debug [instance_methods/method]: AddTabbedWindowOrdered */


// Specifies the duration of a smooth frame-size change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/animationResizeTime(_:)
func (w_ Window) AnimationResizeTime(newFrame Rect /* not a class type */) float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("animationResizeTime:"), newFrame)
	return rv
}/* debug [instance_methods/method]: AnimationResizeTime */


// Indicates whether the window calculates the thickness of a given border automatically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/autorecalculatesContentBorderThickness(for:)
func (w_ Window) AutorecalculatesContentBorderThicknessForEdge(edge RectEdge /* not a class type */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("autorecalculatesContentBorderThicknessForEdge:"), edge)
	return rv
}/* debug [instance_methods/method]: AutorecalculatesContentBorderThicknessForEdge */


// Returns a backing store pixel-aligned rectangle in window coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backingAlignedRect(_:options:)
func (w_ Window) BackingAlignedRectOptions(rect Rect /* not a class type */, options AlignmentOptions /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
	return rv
}/* debug [instance_methods/method]: BackingAlignedRectOptions */


// Informs the window that it has become the key window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/becomeKey()
func (w_ Window) BecomeKeyWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("becomeKeyWindow"))
}/* debug [instance_methods/method]: BecomeKeyWindow */


// Informs the window that it has become the main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/becomeMain()
func (w_ Window) BecomeMainWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("becomeMainWindow"))
}/* debug [instance_methods/method]: BecomeMainWindow */


// Starts a document-modal session and presents the specified critical sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/beginCriticalSheet(_:completionHandler:)
func (w_ Window) BeginCriticalSheetCompletionHandler(sheetWindow IWindow, handler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("beginCriticalSheet:completionHandler:"), sheetWindow, handler)
}/* debug [instance_methods/method]: BeginCriticalSheetCompletionHandler */


// Starts a document-modal session and presents—or queues for presentation—a sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/beginSheet(_:completionHandler:)
func (w_ Window) BeginSheetCompletionHandler(sheetWindow IWindow, handler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("beginSheet:completionHandler:"), sheetWindow, handler)
}/* debug [instance_methods/method]: BeginSheetCompletionHandler */


// A Boolean value that indicates if the window and its screen use a color space that can represent the specified display gamut.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canRepresent(_:)
func (w_ Window) CanRepresentDisplayGamut(displayGamut DisplayGamut) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canRepresentDisplayGamut:"), displayGamut)
	return rv
}/* debug [instance_methods/method]: CanRepresentDisplayGamut */


// Positions the window’s top-left to a given point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/cascadeTopLeft(from:)
func (w_ Window) CascadeTopLeftFromPoint(topLeftPoint vision.Point) vision.Point {
	rv := objc.Send[vision.Point](w_.ID, objc.Sel("cascadeTopLeftFromPoint:"), topLeftPoint)
	return rv
}/* debug [instance_methods/method]: CascadeTopLeftFromPoint */


// Sets the window’s location to the center of the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/center()
func (w_ Window) Center() {
	objc.Send[objc.ID](w_.ID, objc.Sel("center"))
}/* debug [instance_methods/method]: Center */


// Removes the window from the screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/close()
func (w_ Window) Close() {
	objc.Send[objc.ID](w_.ID, objc.Sel("close"))
}/* debug [instance_methods/method]: Close */


// Modifies and returns a frame rectangle so that its top edge lies on a specific screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/constrainFrameRect(_:to:)
func (w_ Window) ConstrainFrameRectToScreen(frameRect Rect /* not a class type */, screen IScreen) Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("constrainFrameRect:toScreen:"), frameRect, screen)
	return rv
}/* debug [instance_methods/method]: ConstrainFrameRectToScreen */


// Returns the window’s content rectangle with a given frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentRect(forFrameRect:)
func (w_ Window) ContentRectForFrameRect(frameRect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("contentRectForFrameRect:"), frameRect)
	return rv
}/* debug [instance_methods/method]: ContentRectForFrameRect */


// Converts a rectangle from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertFromBacking(_:)
func (w_ Window) ConvertRectFromBacking(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}/* debug [instance_methods/method]: ConvertRectFromBacking */


// Converts a rectangle from the screen coordinate system to the window’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertFromScreen(_:)
func (w_ Window) ConvertRectFromScreen(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("convertRectFromScreen:"), rect)
	return rv
}/* debug [instance_methods/method]: ConvertRectFromScreen */


// Converts a point from the screen coordinate system to the window’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPoint(fromScreen:)
func (w_ Window) ConvertPointFromScreen(point vision.Point) vision.Point {
	rv := objc.Send[vision.Point](w_.ID, objc.Sel("convertPointFromScreen:"), point)
	return rv
}/* debug [instance_methods/method]: ConvertPointFromScreen */


// Converts a point to the screen coordinate system from the window’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPoint(toScreen:)
func (w_ Window) ConvertPointToScreen(point vision.Point) vision.Point {
	rv := objc.Send[vision.Point](w_.ID, objc.Sel("convertPointToScreen:"), point)
	return rv
}/* debug [instance_methods/method]: ConvertPointToScreen */


// Converts a point from its pixel-aligned backing store coordinate system to the window’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPointFromBacking(_:)
func (w_ Window) ConvertPointFromBacking(point vision.Point) vision.Point {
	rv := objc.Send[vision.Point](w_.ID, objc.Sel("convertPointFromBacking:"), point)
	return rv
}/* debug [instance_methods/method]: ConvertPointFromBacking */


// Converts a point from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertPointToBacking(_:)
func (w_ Window) ConvertPointToBacking(point vision.Point) vision.Point {
	rv := objc.Send[vision.Point](w_.ID, objc.Sel("convertPointToBacking:"), point)
	return rv
}/* debug [instance_methods/method]: ConvertPointToBacking */


// Converts a rectangle from the window’s coordinate system to its pixel-aligned backing store coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertToBacking(_:)
func (w_ Window) ConvertRectToBacking(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("convertRectToBacking:"), rect)
	return rv
}/* debug [instance_methods/method]: ConvertRectToBacking */


// Converts a rectangle to the screen coordinate system from the window’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/convertToScreen(_:)
func (w_ Window) ConvertRectToScreen(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("convertRectToScreen:"), rect)
	return rv
}/* debug [instance_methods/method]: ConvertRectToScreen */


// De-minimizes the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/deminiaturize(_:)
func (w_ Window) Deminiaturize(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("deminiaturize:"), sender)
}/* debug [instance_methods/method]: Deminiaturize */


// Disables all cursor rectangle management within the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/disableCursorRects()
func (w_ Window) DisableCursorRects() {
	objc.Send[objc.ID](w_.ID, objc.Sel("disableCursorRects"))
}/* debug [instance_methods/method]: DisableCursorRects */


// Invalidates all cursor rectangles in the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/discardCursorRects()
func (w_ Window) DiscardCursorRects() {
	objc.Send[objc.ID](w_.ID, objc.Sel("discardCursorRects"))
}/* debug [instance_methods/method]: DiscardCursorRects */


// Forwards the message to the global application object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/discardEvents(matching:before:)
func (w_ Window) DiscardEventsMatchingMaskBeforeEvent(mask EventMask, lastEvent IEvent) {
	objc.Send[objc.ID](w_.ID, objc.Sel("discardEventsMatchingMask:beforeEvent:"), mask, lastEvent)
}/* debug [instance_methods/method]: DiscardEventsMatchingMaskBeforeEvent */


// Passes a display message down the window’s view hierarchy, thus redrawing all views within the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/display()
func (w_ Window) Display() {
	objc.Send[objc.ID](w_.ID, objc.Sel("display"))
}/* debug [instance_methods/method]: Display */


// Passes a display message down the window’s view hierarchy, thus redrawing all views that need displaying.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/displayIfNeeded()
func (w_ Window) DisplayIfNeeded() {
	objc.Send[objc.ID](w_.ID, objc.Sel("displayIfNeeded"))
}/* debug [instance_methods/method]: DisplayIfNeeded */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/displayLink(target:selector:)
func (w_ Window) DisplayLinkWithTargetSelector(target objc.IObject, selector objc.SEL) quartzcore.DisplayLink {
	rv := objc.Send[quartzcore.DisplayLink](w_.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return rv
}/* debug [instance_methods/method]: DisplayLinkWithTargetSelector */


// Reenables cursor rectangle management within the window after a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/enableCursorRects()
func (w_ Window) EnableCursorRects() {
	objc.Send[objc.ID](w_.ID, objc.Sel("enableCursorRects"))
}/* debug [instance_methods/method]: EnableCursorRects */


// Forces the field editor to give up its first responder status and prepares it for its next assignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/endEditing(for:)
func (w_ Window) EndEditingFor(object objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("endEditingFor:"), object)
}/* debug [instance_methods/method]: EndEditingFor */


// Ends a document-modal session and dismisses the specified sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/endSheet(_:)-4dmmq
func (w_ Window) EndSheet(sheetWindow IWindow) {
	objc.Send[objc.ID](w_.ID, objc.Sel("endSheet:"), sheetWindow)
}/* debug [instance_methods/method]: EndSheet */


// Ends a document-modal session and dismisses the specified sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/endSheet(_:returnCode:)
func (w_ Window) EndSheetReturnCode(sheetWindow IWindow, returnCode ModalResponse /* typedef */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("endSheet:returnCode:"), sheetWindow, returnCode)
}/* debug [instance_methods/method]: EndSheetReturnCode */


// Returns the window’s field editor, creating it if requested.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/fieldEditor(_:for:)
func (w_ Window) FieldEditorForObject(createFlag bool, object objc.IObject) IText {
	rv := objc.Send[Text](w_.ID, objc.Sel("fieldEditor:forObject:"), createFlag, object)
	return rv
}/* debug [instance_methods/method]: FieldEditorForObject */


// Returns the window’s frame rectangle with a given content rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frameRect(forContentRect:)
func (w_ Window) FrameRectForContentRect(contentRect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("frameRectForContentRect:"), contentRect)
	return rv
}/* debug [instance_methods/method]: FrameRectForContentRect */


// Marks as invalid the cursor rectangles of a given view object in the window, so they’ll be set up again when the window becomes key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/invalidateCursorRects(for:)
func (w_ Window) InvalidateCursorRectsForView(view IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("invalidateCursorRectsForView:"), view)
}/* debug [instance_methods/method]: InvalidateCursorRectsForView */


// Invalidates the window shadow so that it is recomputed based on the current window shape.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/invalidateShadow()
func (w_ Window) InvalidateShadow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("invalidateShadow"))
}/* debug [instance_methods/method]: InvalidateShadow */


// Attempts to make a given responder the first responder for the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/makeFirstResponder(_:)
func (w_ Window) MakeFirstResponder(responder IResponder) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("makeFirstResponder:"), responder)
	return rv
}/* debug [instance_methods/method]: MakeFirstResponder */


// Makes the window the key window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/makeKey()
func (w_ Window) MakeKeyWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("makeKeyWindow"))
}/* debug [instance_methods/method]: MakeKeyWindow */


// Moves the window to the front of the screen list, within its level, and makes it the key window; that is, it shows the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/makeKeyAndOrderFront(_:)
func (w_ Window) MakeKeyAndOrderFront(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("makeKeyAndOrderFront:"), sender)
}/* debug [instance_methods/method]: MakeKeyAndOrderFront */


// Makes the window the main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/makeMain()
func (w_ Window) MakeMainWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("makeMainWindow"))
}/* debug [instance_methods/method]: MakeMainWindow */


// Merges all open windows into a single tabbed window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/mergeAllWindows(_:)
func (w_ Window) MergeAllWindows(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("mergeAllWindows:"), sender)
}/* debug [instance_methods/method]: MergeAllWindows */


// Removes the window from the screen list and displays the minimized window in the Dock.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/miniaturize(_:)
func (w_ Window) Miniaturize(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("miniaturize:"), sender)
}/* debug [instance_methods/method]: Miniaturize */


// Moves the tab to a new containing window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/moveTabToNewWindow(_:)
func (w_ Window) MoveTabToNewWindow(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("moveTabToNewWindow:"), sender)
}/* debug [instance_methods/method]: MoveTabToNewWindow */


// Returns the next event matching a given mask.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/nextEvent(matching:)
func (w_ Window) NextEventMatchingMask(mask EventMask) IEvent {
	rv := objc.Send[Event](w_.ID, objc.Sel("nextEventMatchingMask:"), mask)
	return rv
}/* debug [instance_methods/method]: NextEventMatchingMask */


// Forwards the message to the global application object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/nextEvent(matching:until:inMode:dequeue:)
func (w_ Window) NextEventMatchingMaskUntilDateInModeDequeue(mask EventMask, expiration objc.IObject /* cross-framework: NSDate */, mode RunLoopMode /* not a class type */, deqFlag bool) IEvent {
	rv := objc.Send[Event](w_.ID, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, mode, deqFlag)
	return rv
}/* debug [instance_methods/method]: NextEventMatchingMaskUntilDateInModeDequeue */


// Repositions the window’s window device in the window server’s screen list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/order(_:relativeTo:)
func (w_ Window) OrderWindowRelativeTo(place WindowOrderingMode, otherWin int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("orderWindow:relativeTo:"), place, otherWin)
}/* debug [instance_methods/method]: OrderWindowRelativeTo */


// Moves the window to the back of its level in the screen list, without changing either the key window or the main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/orderBack(_:)
func (w_ Window) OrderBack(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("orderBack:"), sender)
}/* debug [instance_methods/method]: OrderBack */


// Moves the window to the front of its level in the screen list, without changing either the key window or the main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/orderFront(_:)
func (w_ Window) OrderFront(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("orderFront:"), sender)
}/* debug [instance_methods/method]: OrderFront */


// Moves the window to the front of its level, even if its application isn’t active, without changing either the key window or the main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/orderFrontRegardless()
func (w_ Window) OrderFrontRegardless() {
	objc.Send[objc.ID](w_.ID, objc.Sel("orderFrontRegardless"))
}/* debug [instance_methods/method]: OrderFrontRegardless */


// Removes the window from the screen list, which hides the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/orderOut(_:)
func (w_ Window) OrderOut(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("orderOut:"), sender)
}/* debug [instance_methods/method]: OrderOut */


// Simulates the user clicking the close button by momentarily highlighting the button and then closing the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/performClose(_:)
func (w_ Window) PerformClose(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performClose:"), sender)
}/* debug [instance_methods/method]: PerformClose */


// Starts a window drag based on the specified mouse-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/performDrag(with:)
func (w_ Window) PerformWindowDragWithEvent(event IEvent) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performWindowDragWithEvent:"), event)
}/* debug [instance_methods/method]: PerformWindowDragWithEvent */


// Simulates the user clicking the minimize button by momentarily highlighting the button, then minimizing the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/performMiniaturize(_:)
func (w_ Window) PerformMiniaturize(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performMiniaturize:"), sender)
}/* debug [instance_methods/method]: PerformMiniaturize */


// This action method simulates the user clicking the zoom box by momentarily highlighting the button and then zooming the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/performZoom(_:)
func (w_ Window) PerformZoom(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("performZoom:"), sender)
}/* debug [instance_methods/method]: PerformZoom */


// Forwards the message to the global application object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/postEvent(_:atStart:)
func (w_ Window) PostEventAtStart(event IEvent, flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("postEvent:atStart:"), event, flag)
}/* debug [instance_methods/method]: PostEventAtStart */


// Marks the key view loop as “dirty” and in need of recalculation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/recalculateKeyViewLoop()
func (w_ Window) RecalculateKeyViewLoop() {
	objc.Send[objc.ID](w_.ID, objc.Sel("recalculateKeyViewLoop"))
}/* debug [instance_methods/method]: RecalculateKeyViewLoop */


// Registers a set of pasteboard types that the window accepts as the destination of an image-dragging session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/registerForDraggedTypes(_:)
func (w_ Window) RegisterForDraggedTypes(newTypes []string) {
	objc.Send[objc.ID](w_.ID, objc.Sel("registerForDraggedTypes:"), newTypes)
}/* debug [instance_methods/method]: RegisterForDraggedTypes */


// Clears the window’s cursor rectangles and the cursor rectangles of the objects in its view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resetCursorRects()
func (w_ Window) ResetCursorRects() {
	objc.Send[objc.ID](w_.ID, objc.Sel("resetCursorRects"))
}/* debug [instance_methods/method]: ResetCursorRects */


// Resigns the window’s key window status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resignKey()
func (w_ Window) ResignKeyWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("resignKeyWindow"))
}/* debug [instance_methods/method]: ResignKeyWindow */


// Resigns the window’s main window status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resignMain()
func (w_ Window) ResignMainWindow() {
	objc.Send[objc.ID](w_.ID, objc.Sel("resignMainWindow"))
}/* debug [instance_methods/method]: ResignMainWindow */


// Saves the window’s frame rectangle in the user defaults system under a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/saveFrame(usingName:)
func (w_ Window) SaveFrameUsingName(name WindowFrameAutosaveName /* typedef */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("saveFrameUsingName:"), name)
}/* debug [instance_methods/method]: SaveFrameUsingName */


// Gives key view status to the view that follows the given view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectKeyView(following:)
func (w_ Window) SelectKeyViewFollowingView(view IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectKeyViewFollowingView:"), view)
}/* debug [instance_methods/method]: SelectKeyViewFollowingView */


// Gives key view status to the view that precedes the given view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectKeyView(preceding:)
func (w_ Window) SelectKeyViewPrecedingView(view IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectKeyViewPrecedingView:"), view)
}/* debug [instance_methods/method]: SelectKeyViewPrecedingView */


// Searches for a candidate next key view and, if it finds one, tries to make it the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectNextKeyView(_:)
func (w_ Window) SelectNextKeyView(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectNextKeyView:"), sender)
}/* debug [instance_methods/method]: SelectNextKeyView */


// Selects the next tab in the tab group in the trailing direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectNextTab(_:)
func (w_ Window) SelectNextTab(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectNextTab:"), sender)
}/* debug [instance_methods/method]: SelectNextTab */


// Searches for a candidate previous key view and, if it finds one, tries to make it the first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectPreviousKeyView(_:)
func (w_ Window) SelectPreviousKeyView(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectPreviousKeyView:"), sender)
}/* debug [instance_methods/method]: SelectPreviousKeyView */


// Selects the previous tab in the tab group in the leading direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/selectPreviousTab(_:)
func (w_ Window) SelectPreviousTab(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("selectPreviousTab:"), sender)
}/* debug [instance_methods/method]: SelectPreviousTab */


// This action method dispatches mouse and keyboard events the global application object sends to the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/sendEvent(_:)
func (w_ Window) SendEvent(event IEvent) {
	objc.Send[objc.ID](w_.ID, objc.Sel("sendEvent:"), event)
}/* debug [instance_methods/method]: SendEvent */


// Sets the size of the window’s content view to a given size, which is expressed in the window’s base coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setContentSize(_:)
func (w_ Window) SetContentSize(size Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentSize:"), size)
}/* debug [instance_methods/method]: SetContentSize */


// Sets a Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setDynamicDepthLimit(_:)
func (w_ Window) SetDynamicDepthLimit(flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDynamicDepthLimit:"), flag)
}/* debug [instance_methods/method]: SetDynamicDepthLimit */


// Sets the origin and size of the window’s frame rectangle according to a given frame rectangle, thereby setting its position and size onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(_:display:)
func (w_ Window) SetFrameDisplay(frameRect Rect /* not a class type */, flag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrame:display:"), frameRect, flag)
}/* debug [instance_methods/method]: SetFrameDisplay */


// Sets the origin and size of the window’s frame rectangle, with optional animation, according to a given frame rectangle, thereby setting its position and size onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(_:display:animate:)
func (w_ Window) SetFrameDisplayAnimate(frameRect Rect /* not a class type */, displayFlag bool, animateFlag bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrame:display:animate:"), frameRect, displayFlag, animateFlag)
}/* debug [instance_methods/method]: SetFrameDisplayAnimate */


// Sets the window’s frame rectangle from a given string representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrame(from:)
func (w_ Window) SetFrameFromString(string_ WindowPersistableFrameDescriptor /* typedef */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameFromString:"), string_)
}/* debug [instance_methods/method]: SetFrameFromString */


// Positions the bottom-left corner of the window’s frame rectangle at a given point in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameOrigin(_:)
func (w_ Window) SetFrameOrigin(point vision.Point) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameOrigin:"), point)
}/* debug [instance_methods/method]: SetFrameOrigin */


// Positions the top-left corner of the window’s frame rectangle at a given point in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameTopLeftPoint(_:)
func (w_ Window) SetFrameTopLeftPoint(point vision.Point) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameTopLeftPoint:"), point)
}/* debug [instance_methods/method]: SetFrameTopLeftPoint */


// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameUsingName(_:)
func (w_ Window) SetFrameUsingName(name WindowFrameAutosaveName /* typedef */) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("setFrameUsingName:"), name)
	return rv
}/* debug [instance_methods/method]: SetFrameUsingName */


// Sets the window’s frame rectangle by reading the rectangle data stored under a given name from the defaults system. Can operate on non-resizable windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setFrameUsingName(_:force:)
func (w_ Window) SetFrameUsingNameForce(name WindowFrameAutosaveName /* typedef */, force bool) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("setFrameUsingName:force:"), name, force)
	return rv
}/* debug [instance_methods/method]: SetFrameUsingNameForce */


// Sets a given path as the window’s title, formatting it as a file-system path, and records this path as the window’s associated file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/setTitleWithRepresentedFilename(_:)
func (w_ Window) SetTitleWithRepresentedFilename(filename objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitleWithRepresentedFilename:"), filename)
}/* debug [instance_methods/method]: SetTitleWithRepresentedFilename */


// Takes the window into or out of fullscreen mode,
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toggleFullScreen(_:)
func (w_ Window) ToggleFullScreen(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("toggleFullScreen:"), sender)
}/* debug [instance_methods/method]: ToggleFullScreen */


// Shows or hides the tab bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toggleTabBar(_:)
func (w_ Window) ToggleTabBar(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("toggleTabBar:"), sender)
}/* debug [instance_methods/method]: ToggleTabBar */


// Shows or hides the tab overview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/toggleTabOverview(_:)
func (w_ Window) ToggleTabOverview(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("toggleTabOverview:"), sender)
}/* debug [instance_methods/method]: ToggleTabOverview */


// Tracks events that match the specified mask using the specified tracking handler until the tracking handler explicitly terminates tracking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/trackEvents(matching:timeout:mode:handler:)
func (w_ Window) TrackEventsMatchingMaskTimeoutModeHandler(mask EventMask, timeout float64, mode RunLoopMode /* not a class type */, trackingHandler unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("trackEventsMatchingMask:timeout:mode:handler:"), mask, timeout, mode, trackingHandler)
}/* debug [instance_methods/method]: TrackEventsMatchingMaskTimeoutModeHandler */


// Dispatches action messages with a given argument.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tryToPerform(_:with:)
func (w_ Window) TryToPerformWith(action objc.SEL, object objc.IObject) bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("tryToPerform:with:"), action, object)
	return rv
}/* debug [instance_methods/method]: TryToPerformWith */


// Unregisters the window as a possible destination for dragging operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/unregisterDraggedTypes()
func (w_ Window) UnregisterDraggedTypes() {
	objc.Send[objc.ID](w_.ID, objc.Sel("unregisterDraggedTypes"))
}/* debug [instance_methods/method]: UnregisterDraggedTypes */


// Displays a visual representation of the supplied constraints in the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/visualizeConstraints(_:)
func (w_ Window) VisualizeConstraints(constraints []LayoutConstraint) {
	objc.Send[objc.ID](w_.ID, objc.Sel("visualizeConstraints:"), constraints)
}/* debug [instance_methods/method]: VisualizeConstraints */


// Toggles the size and location of the window between its standard state (which the application provides as the best size to display the window’s data) and its user state (a new size and location the user may have set by moving or resizing the window).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/zoom(_:)
func (w_ Window) Zoom(sender objc.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("zoom:"), sender)
}/* debug [instance_methods/method]: Zoom */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Window */

// A Boolean value that indicates whether the window accepts mouse-moved events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/acceptsMouseMovedEvents
func (w_ Window) AcceptsMouseMovedEvents() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("acceptsMouseMovedEvents"))
	return rv
}/* debug [instance_properties/getter]: acceptsMouseMovedEvents */


// A Boolean value that indicates whether the window accepts mouse-moved events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/acceptsMouseMovedEvents
func (w_ Window) SetAcceptsMouseMovedEvents(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAcceptsMouseMovedEvents:"), value)
}/* debug [instance_properties/setter]: acceptsMouseMovedEvents */


// A Boolean value that indicates whether the app can automatically organize windows into tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsAutomaticWindowTabbing
func (w_ Window) AllowsAutomaticWindowTabbing() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsAutomaticWindowTabbing"))
	return rv
}/* debug [instance_properties/getter]: allowsAutomaticWindowTabbing */


// A Boolean value that indicates whether the app can automatically organize windows into tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsAutomaticWindowTabbing
func (w_ Window) SetAllowsAutomaticWindowTabbing(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsAutomaticWindowTabbing:"), value)
}/* debug [instance_properties/setter]: allowsAutomaticWindowTabbing */


// A Boolean value that indicates whether the window allows multithreaded view drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsConcurrentViewDrawing
func (w_ Window) AllowsConcurrentViewDrawing() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsConcurrentViewDrawing"))
	return rv
}/* debug [instance_properties/getter]: allowsConcurrentViewDrawing */


// A Boolean value that indicates whether the window allows multithreaded view drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/allowsConcurrentViewDrawing
func (w_ Window) SetAllowsConcurrentViewDrawing(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsConcurrentViewDrawing:"), value)
}/* debug [instance_properties/setter]: allowsConcurrentViewDrawing */


// The window’s alpha value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/alphaValue
func (w_ Window) AlphaValue() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("alphaValue"))
	return rv
}/* debug [instance_properties/getter]: alphaValue */


// The window’s alpha value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/alphaValue
func (w_ Window) SetAlphaValue(value float64) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAlphaValue:"), value)
}/* debug [instance_properties/setter]: alphaValue */


// The window’s automatic animation behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/animationBehavior-swift.property
func (w_ Window) AnimationBehavior() WindowAnimationBehavior {
	rv := objc.Send[WindowAnimationBehavior](w_.ID, objc.Sel("animationBehavior"))
	return rv
}/* debug [instance_properties/getter]: animationBehavior */


// The window’s automatic animation behavior.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/animationBehavior-swift.property
func (w_ Window) SetAnimationBehavior(value WindowAnimationBehavior) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAnimationBehavior:"), value)
}/* debug [instance_properties/setter]: animationBehavior */


// A Boolean value that indicates whether the window’s cursor rectangles are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/areCursorRectsEnabled
func (w_ Window) AreCursorRectsEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("areCursorRectsEnabled"))
	return rv
}/* debug [instance_properties/getter]: areCursorRectsEnabled */


// The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/aspectRatio
func (w_ Window) AspectRatio() Size /* not a class type */ {
	rv := objc.Send[Size](w_.ID, objc.Sel("aspectRatio"))
	return rv
}/* debug [instance_properties/getter]: aspectRatio */


// The window’s aspect ratio, which constrains the size of its frame rectangle to integral multiples of this ratio when the user resizes it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/aspectRatio
func (w_ Window) SetAspectRatio(value Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAspectRatio:"), value)
}/* debug [instance_properties/setter]: aspectRatio */


// The sheet attached to the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/attachedSheet
func (w_ Window) AttachedSheet() IWindow {
	rv := objc.Send[Window](w_.ID, objc.Sel("attachedSheet"))
	return rv
}/* debug [instance_properties/getter]: attachedSheet */


// A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/autorecalculatesKeyViewLoop
func (w_ Window) AutorecalculatesKeyViewLoop() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("autorecalculatesKeyViewLoop"))
	return rv
}/* debug [instance_properties/getter]: autorecalculatesKeyViewLoop */


// A Boolean value that indicates whether the window automatically recalculates the key view loop when views are added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/autorecalculatesKeyViewLoop
func (w_ Window) SetAutorecalculatesKeyViewLoop(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAutorecalculatesKeyViewLoop:"), value)
}/* debug [instance_properties/setter]: autorecalculatesKeyViewLoop */


// The color of the window’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backgroundColor
func (w_ Window) BackgroundColor() IColor {
	rv := objc.Send[Color](w_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The color of the window’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backgroundColor
func (w_ Window) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// The location of the window’s backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backingLocation-swift.property
func (w_ Window) BackingLocation() WindowBackingLocation {
	rv := objc.Send[WindowBackingLocation](w_.ID, objc.Sel("backingLocation"))
	return rv
}/* debug [instance_properties/getter]: backingLocation */


// The backing scale factor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backingScaleFactor
func (w_ Window) BackingScaleFactor() float64 {
	rv := objc.Send[float64](w_.ID, objc.Sel("backingScaleFactor"))
	return rv
}/* debug [instance_properties/getter]: backingScaleFactor */


// The window’s backing store type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backingType
func (w_ Window) BackingType() BackingStoreType {
	rv := objc.Send[BackingStoreType](w_.ID, objc.Sel("backingType"))
	return rv
}/* debug [instance_properties/getter]: backingType */


// The window’s backing store type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/backingType
func (w_ Window) SetBackingType(value BackingStoreType) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setBackingType:"), value)
}/* debug [instance_properties/setter]: backingType */


// A Boolean value that indicates whether the window can become the key window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeKey
func (w_ Window) CanBecomeKeyWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canBecomeKeyWindow"))
	return rv
}/* debug [instance_properties/getter]: canBecomeKeyWindow */


// A Boolean value that indicates whether the window can become the application’s main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeMain
func (w_ Window) CanBecomeMainWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canBecomeMainWindow"))
	return rv
}/* debug [instance_properties/getter]: canBecomeMainWindow */


// A Boolean value that indicates whether the window can be displayed at the login window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeVisibleWithoutLogin
func (w_ Window) CanBecomeVisibleWithoutLogin() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canBecomeVisibleWithoutLogin"))
	return rv
}/* debug [instance_properties/getter]: canBecomeVisibleWithoutLogin */


// A Boolean value that indicates whether the window can be displayed at the login window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canBecomeVisibleWithoutLogin
func (w_ Window) SetCanBecomeVisibleWithoutLogin(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanBecomeVisibleWithoutLogin:"), value)
}/* debug [instance_properties/setter]: canBecomeVisibleWithoutLogin */


// A Boolean value that indicates whether the window can hide when its application becomes hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canHide
func (w_ Window) CanHide() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canHide"))
	return rv
}/* debug [instance_properties/getter]: canHide */


// A Boolean value that indicates whether the window can hide when its application becomes hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/canHide
func (w_ Window) SetCanHide(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanHide:"), value)
}/* debug [instance_properties/setter]: canHide */


// A value that identifies the window’s behavior in window collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/collectionBehavior-swift.property
func (w_ Window) CollectionBehavior() WindowCollectionBehavior {
	rv := objc.Send[WindowCollectionBehavior](w_.ID, objc.Sel("collectionBehavior"))
	return rv
}/* debug [instance_properties/getter]: collectionBehavior */


// A value that identifies the window’s behavior in window collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/collectionBehavior-swift.property
func (w_ Window) SetCollectionBehavior(value WindowCollectionBehavior) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCollectionBehavior:"), value)
}/* debug [instance_properties/setter]: collectionBehavior */


// The window’s color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/colorSpace
func (w_ Window) ColorSpace() IColorSpace {
	rv := objc.Send[ColorSpace](w_.ID, objc.Sel("colorSpace"))
	return rv
}/* debug [instance_properties/getter]: colorSpace */


// The window’s color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/colorSpace
func (w_ Window) SetColorSpace(value IColorSpace) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setColorSpace:"), value)
}/* debug [instance_properties/setter]: colorSpace */


// The window’s content aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentAspectRatio
func (w_ Window) ContentAspectRatio() Size /* not a class type */ {
	rv := objc.Send[Size](w_.ID, objc.Sel("contentAspectRatio"))
	return rv
}/* debug [instance_properties/getter]: contentAspectRatio */


// The window’s content aspect ratio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentAspectRatio
func (w_ Window) SetContentAspectRatio(value Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentAspectRatio:"), value)
}/* debug [instance_properties/setter]: contentAspectRatio */


// A value used by Auto Layout constraints to automatically bind to the value of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentLayoutGuide
func (w_ Window) ContentLayoutGuide() objc.ID {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("contentLayoutGuide"))
	return rv
}/* debug [instance_properties/getter]: contentLayoutGuide */


// The area inside the window that is for non-obscured content, in window coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentLayoutRect
func (w_ Window) ContentLayoutRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("contentLayoutRect"))
	return rv
}/* debug [instance_properties/getter]: contentLayoutRect */


// The maximum size of the window’s content view in the window’s base coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentMaxSize
func (w_ Window) ContentMaxSize() Size /* not a class type */ {
	rv := objc.Send[Size](w_.ID, objc.Sel("contentMaxSize"))
	return rv
}/* debug [instance_properties/getter]: contentMaxSize */


// The maximum size of the window’s content view in the window’s base coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentMaxSize
func (w_ Window) SetContentMaxSize(value Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentMaxSize:"), value)
}/* debug [instance_properties/setter]: contentMaxSize */


// The minimum size of the window’s content view in the window’s base coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentMinSize
func (w_ Window) ContentMinSize() Size /* not a class type */ {
	rv := objc.Send[Size](w_.ID, objc.Sel("contentMinSize"))
	return rv
}/* debug [instance_properties/getter]: contentMinSize */


// The minimum size of the window’s content view in the window’s base coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentMinSize
func (w_ Window) SetContentMinSize(value Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentMinSize:"), value)
}/* debug [instance_properties/setter]: contentMinSize */


// The window’s content-view resizing increments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentResizeIncrements
func (w_ Window) ContentResizeIncrements() Size /* not a class type */ {
	rv := objc.Send[Size](w_.ID, objc.Sel("contentResizeIncrements"))
	return rv
}/* debug [instance_properties/getter]: contentResizeIncrements */


// The window’s content-view resizing increments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentResizeIncrements
func (w_ Window) SetContentResizeIncrements(value Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentResizeIncrements:"), value)
}/* debug [instance_properties/setter]: contentResizeIncrements */


// The window’s content view, the highest accessible view object in the window’s view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentView
func (w_ Window) ContentView() IView {
	rv := objc.Send[View](w_.ID, objc.Sel("contentView"))
	return rv
}/* debug [instance_properties/getter]: contentView */


// The window’s content view, the highest accessible view object in the window’s view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentView
func (w_ Window) SetContentView(value IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentView:"), value)
}/* debug [instance_properties/setter]: contentView */


// The main content view controller for the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentViewController
func (w_ Window) ContentViewController() IViewController {
	rv := objc.Send[ViewController](w_.ID, objc.Sel("contentViewController"))
	return rv
}/* debug [instance_properties/getter]: contentViewController */


// The main content view controller for the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/contentViewController
func (w_ Window) SetContentViewController(value IViewController) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setContentViewController:"), value)
}/* debug [instance_properties/setter]: contentViewController */


// The event currently being processed by the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/currentEvent
func (w_ Window) CurrentEvent() IEvent {
	rv := objc.Send[Event](w_.ID, objc.Sel("currentEvent"))
	return rv
}/* debug [instance_properties/getter]: currentEvent */


// The deepest screen the window is on (it may be split over several screens).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/deepestScreen
func (w_ Window) DeepestScreen() IScreen {
	rv := objc.Send[Screen](w_.ID, objc.Sel("deepestScreen"))
	return rv
}/* debug [instance_properties/getter]: deepestScreen */


// Returns the default depth limit for instances of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/defaultDepthLimit
func (w_ Window) DefaultDepthLimit() WindowDepth {
	rv := objc.Send[WindowDepth](w_.ID, objc.Sel("defaultDepthLimit"))
	return rv
}/* debug [instance_properties/getter]: defaultDepthLimit */


// The depth limit of the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/depthLimit
func (w_ Window) DepthLimit() WindowDepth {
	rv := objc.Send[WindowDepth](w_.ID, objc.Sel("depthLimit"))
	return rv
}/* debug [instance_properties/getter]: depthLimit */


// The depth limit of the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/depthLimit
func (w_ Window) SetDepthLimit(value WindowDepth) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDepthLimit:"), value)
}/* debug [instance_properties/setter]: depthLimit */


// A dictionary containing information about the window’s resolution, such as color, depth, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/deviceDescription
func (w_ Window) DeviceDescription() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](w_.ID, objc.Sel("deviceDescription"))
	return rv
}/* debug [instance_properties/getter]: deviceDescription */


// A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/displaysWhenScreenProfileChanges
func (w_ Window) DisplaysWhenScreenProfileChanges() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("displaysWhenScreenProfileChanges"))
	return rv
}/* debug [instance_properties/getter]: displaysWhenScreenProfileChanges */


// A Boolean value that indicates whether the window context should be updated when the screen profile changes or when the window moves to a different screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/displaysWhenScreenProfileChanges
func (w_ Window) SetDisplaysWhenScreenProfileChanges(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDisplaysWhenScreenProfileChanges:"), value)
}/* debug [instance_properties/setter]: displaysWhenScreenProfileChanges */


// The collection of drawers associated with the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/drawers
func (w_ Window) Drawers() []Drawer {
	rv := objc.Send[[]Drawer](w_.ID, objc.Sel("drawers"))
	return rv
}/* debug [instance_properties/getter]: drawers */


// The window’s first responder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/firstResponder
func (w_ Window) FirstResponder() IResponder {
	rv := objc.Send[Responder](w_.ID, objc.Sel("firstResponder"))
	return rv
}/* debug [instance_properties/getter]: firstResponder */


// The window’s frame rectangle in screen coordinates, including the title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frame
func (w_ Window) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The name used to automatically save the window’s frame rectangle data in the defaults system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frameAutosaveName-swift.property
func (w_ Window) FrameAutosaveName() WindowFrameAutosaveName /* typedef */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("frameAutosaveName"))
	return rv
}/* debug [instance_properties/getter]: frameAutosaveName */


// A string representation of the window’s frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/frameDescriptor
func (w_ Window) StringWithSavedFrame() WindowPersistableFrameDescriptor /* typedef */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("stringWithSavedFrame"))
	return rv
}/* debug [instance_properties/getter]: stringWithSavedFrame */


// The graphics context associated with the window for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/graphicsContext
func (w_ Window) GraphicsContext() IGraphicsContext {
	rv := objc.Send[GraphicsContext](w_.ID, objc.Sel("graphicsContext"))
	return rv
}/* debug [instance_properties/getter]: graphicsContext */


// A Boolean value that indicates whether the window’s depth limit can change to match the depth of the screen it’s on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hasDynamicDepthLimit
func (w_ Window) HasDynamicDepthLimit() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasDynamicDepthLimit"))
	return rv
}/* debug [instance_properties/getter]: hasDynamicDepthLimit */


// A Boolean value that indicates whether the window has a shadow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hasShadow
func (w_ Window) HasShadow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasShadow"))
	return rv
}/* debug [instance_properties/getter]: hasShadow */


// A Boolean value that indicates whether the window has a shadow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hasShadow
func (w_ Window) SetHasShadow(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasShadow:"), value)
}/* debug [instance_properties/setter]: hasShadow */


// A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hidesOnDeactivate
func (w_ Window) HidesOnDeactivate() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hidesOnDeactivate"))
	return rv
}/* debug [instance_properties/getter]: hidesOnDeactivate */


// A Boolean value that indicates whether the window is removed from the screen when its application becomes inactive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/hidesOnDeactivate
func (w_ Window) SetHidesOnDeactivate(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHidesOnDeactivate:"), value)
}/* debug [instance_properties/setter]: hidesOnDeactivate */


// A Boolean value that indicates whether the window is transparent to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ignoresMouseEvents
func (w_ Window) IgnoresMouseEvents() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("ignoresMouseEvents"))
	return rv
}/* debug [instance_properties/getter]: ignoresMouseEvents */


// A Boolean value that indicates whether the window is transparent to mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ignoresMouseEvents
func (w_ Window) SetIgnoresMouseEvents(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIgnoresMouseEvents:"), value)
}/* debug [instance_properties/setter]: ignoresMouseEvents */


// The view that’s made first responder (also called the key view) the first time the window is placed onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/initialFirstResponder
func (w_ Window) InitialFirstResponder() IView {
	rv := objc.Send[View](w_.ID, objc.Sel("initialFirstResponder"))
	return rv
}/* debug [instance_properties/getter]: initialFirstResponder */


// The view that’s made first responder (also called the key view) the first time the window is placed onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/initialFirstResponder
func (w_ Window) SetInitialFirstResponder(value IView) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInitialFirstResponder:"), value)
}/* debug [instance_properties/setter]: initialFirstResponder */


// A Boolean value that indicates whether the window automatically displays views that need to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isAutodisplay
func (w_ Window) Autodisplay() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("autodisplay"))
	return rv
}/* debug [instance_properties/getter]: autodisplay */


// A Boolean value that indicates whether the window automatically displays views that need to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isAutodisplay
func (w_ Window) SetAutodisplay(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAutodisplay:"), value)
}/* debug [instance_properties/setter]: autodisplay */


// A Boolean value that indicates whether the window’s document has been edited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isDocumentEdited
func (w_ Window) DocumentEdited() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("documentEdited"))
	return rv
}/* debug [instance_properties/getter]: documentEdited */


// A Boolean value that indicates whether the window’s document has been edited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isDocumentEdited
func (w_ Window) SetDocumentEdited(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDocumentEdited:"), value)
}/* debug [instance_properties/setter]: documentEdited */


// A Boolean value that indicates whether the window is excluded from the application’s Windows menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isExcludedFromWindowsMenu
func (w_ Window) ExcludedFromWindowsMenu() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("excludedFromWindowsMenu"))
	return rv
}/* debug [instance_properties/getter]: excludedFromWindowsMenu */


// A Boolean value that indicates whether the window is excluded from the application’s Windows menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isExcludedFromWindowsMenu
func (w_ Window) SetExcludedFromWindowsMenu(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setExcludedFromWindowsMenu:"), value)
}/* debug [instance_properties/setter]: excludedFromWindowsMenu */


// A Boolean value that indicates whether the window’s flushing ability is disabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isFlushWindowDisabled
func (w_ Window) FlushWindowDisabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("flushWindowDisabled"))
	return rv
}/* debug [instance_properties/getter]: flushWindowDisabled */


// A Boolean value that indicates whether the window is the key window for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isKeyWindow
func (w_ Window) KeyWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("keyWindow"))
	return rv
}/* debug [instance_properties/getter]: keyWindow */


// A Boolean value that indicates whether the window is the application’s main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMainWindow
func (w_ Window) MainWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("mainWindow"))
	return rv
}/* debug [instance_properties/getter]: mainWindow */


// A Boolean value that indicates whether the window is minimized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMiniaturized
func (w_ Window) Miniaturized() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("miniaturized"))
	return rv
}/* debug [instance_properties/getter]: miniaturized */


// A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMovable
func (w_ Window) Movable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("movable"))
	return rv
}/* debug [instance_properties/getter]: movable */


// A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMovable
func (w_ Window) SetMovable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMovable:"), value)
}/* debug [instance_properties/setter]: movable */


// A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMovableByWindowBackground
func (w_ Window) MovableByWindowBackground() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("movableByWindowBackground"))
	return rv
}/* debug [instance_properties/getter]: movableByWindowBackground */


// A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isMovableByWindowBackground
func (w_ Window) SetMovableByWindowBackground(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMovableByWindowBackground:"), value)
}/* debug [instance_properties/setter]: movableByWindowBackground */


// A Boolean value that indicates whether the window is on the currently active space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isOnActiveSpace
func (w_ Window) OnActiveSpace() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("onActiveSpace"))
	return rv
}/* debug [instance_properties/getter]: onActiveSpace */


// A Boolean value that indicates whether the window device the window manages is freed when it’s removed from the screen list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isOneShot
func (w_ Window) OneShot() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("oneShot"))
	return rv
}/* debug [instance_properties/getter]: oneShot */


// A Boolean value that indicates whether the window device the window manages is freed when it’s removed from the screen list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isOneShot
func (w_ Window) SetOneShot(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOneShot:"), value)
}/* debug [instance_properties/setter]: oneShot */


// A Boolean value that indicates whether the window is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isOpaque
func (w_ Window) Opaque() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("opaque"))
	return rv
}/* debug [instance_properties/getter]: opaque */


// A Boolean value that indicates whether the window is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isOpaque
func (w_ Window) SetOpaque(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOpaque:"), value)
}/* debug [instance_properties/setter]: opaque */


// A Boolean value that indicates whether the window is released when it receives the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isReleasedWhenClosed
func (w_ Window) ReleasedWhenClosed() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("releasedWhenClosed"))
	return rv
}/* debug [instance_properties/getter]: releasedWhenClosed */


// A Boolean value that indicates whether the window is released when it receives the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isReleasedWhenClosed
func (w_ Window) SetReleasedWhenClosed(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setReleasedWhenClosed:"), value)
}/* debug [instance_properties/setter]: releasedWhenClosed */


// A Boolean value that indicates whether the window has ever run as a modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isSheet
func (w_ Window) Sheet() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("sheet"))
	return rv
}/* debug [instance_properties/getter]: sheet */


// A Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isVisible
func (w_ Window) Visible() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("visible"))
	return rv
}/* debug [instance_properties/getter]: visible */


// A Boolean value that indicates whether the window is in a zoomed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/isZoomed
func (w_ Window) Zoomed() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("zoomed"))
	return rv
}/* debug [instance_properties/getter]: zoomed */


// The direction the window is currently using to change the key view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/keyViewSelectionDirection
func (w_ Window) KeyViewSelectionDirection() SelectionDirection {
	rv := objc.Send[SelectionDirection](w_.ID, objc.Sel("keyViewSelectionDirection"))
	return rv
}/* debug [instance_properties/getter]: keyViewSelectionDirection */


// The window level of the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/level-swift.property
func (w_ Window) Level() WindowLevel /* typedef */ {
	rv := objc.Send[int](w_.ID, objc.Sel("level"))
	return rv
}/* debug [instance_properties/getter]: level */


// The window level of the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/level-swift.property
func (w_ Window) SetLevel(value WindowLevel /* typedef */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setLevel:"), value)
}/* debug [instance_properties/setter]: level */


// A maximum size that is used to determine if a window can fit when it is in full screen in a tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/maxFullScreenContentSize
func (w_ Window) MaxFullScreenContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](w_.ID, objc.Sel("maxFullScreenContentSize"))
	return rv
}/* debug [instance_properties/getter]: maxFullScreenContentSize */


// A maximum size that is used to determine if a window can fit when it is in full screen in a tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/maxFullScreenContentSize
func (w_ Window) SetMaxFullScreenContentSize(value Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMaxFullScreenContentSize:"), value)
}/* debug [instance_properties/setter]: maxFullScreenContentSize */


// The maximum size to which the window’s frame (including its title bar) can be sized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/maxSize
func (w_ Window) MaxSize() Size /* not a class type */ {
	rv := objc.Send[Size](w_.ID, objc.Sel("maxSize"))
	return rv
}/* debug [instance_properties/getter]: maxSize */


// The maximum size to which the window’s frame (including its title bar) can be sized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/maxSize
func (w_ Window) SetMaxSize(value Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMaxSize:"), value)
}/* debug [instance_properties/setter]: maxSize */


// A minimum size that is used to determine if a window can fit when it is in full screen in a tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/minFullScreenContentSize
func (w_ Window) MinFullScreenContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](w_.ID, objc.Sel("minFullScreenContentSize"))
	return rv
}/* debug [instance_properties/getter]: minFullScreenContentSize */


// A minimum size that is used to determine if a window can fit when it is in full screen in a tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/minFullScreenContentSize
func (w_ Window) SetMinFullScreenContentSize(value Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMinFullScreenContentSize:"), value)
}/* debug [instance_properties/setter]: minFullScreenContentSize */


// The minimum size to which the window’s frame (including its title bar) can be sized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/minSize
func (w_ Window) MinSize() Size /* not a class type */ {
	rv := objc.Send[Size](w_.ID, objc.Sel("minSize"))
	return rv
}/* debug [instance_properties/getter]: minSize */


// The minimum size to which the window’s frame (including its title bar) can be sized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/minSize
func (w_ Window) SetMinSize(value Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMinSize:"), value)
}/* debug [instance_properties/setter]: minSize */


// The custom miniaturized window image of the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowImage
func (w_ Window) MiniwindowImage() IImage {
	rv := objc.Send[Image](w_.ID, objc.Sel("miniwindowImage"))
	return rv
}/* debug [instance_properties/getter]: miniwindowImage */


// The custom miniaturized window image of the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowImage
func (w_ Window) SetMiniwindowImage(value IImage) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMiniwindowImage:"), value)
}/* debug [instance_properties/setter]: miniwindowImage */


// The title displayed in the window’s minimized window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowTitle
func (w_ Window) MiniwindowTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("miniwindowTitle"))
	return rv
}/* debug [instance_properties/getter]: miniwindowTitle */


// The title displayed in the window’s minimized window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/miniwindowTitle
func (w_ Window) SetMiniwindowTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setMiniwindowTitle:"), value)
}/* debug [instance_properties/setter]: miniwindowTitle */


// The current location of the pointer reckoned in the window’s base coordinate system, regardless of the current event being handled or of any events pending.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/mouseLocationOutsideOfEventStream
func (w_ Window) MouseLocationOutsideOfEventStream() vision.Point {
	rv := objc.Send[vision.Point](w_.ID, objc.Sel("mouseLocationOutsideOfEventStream"))
	return rv
}/* debug [instance_properties/getter]: mouseLocationOutsideOfEventStream */


// The occlusion state of the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/occlusionState-swift.property
func (w_ Window) OcclusionState() WindowOcclusionState {
	rv := objc.Send[WindowOcclusionState](w_.ID, objc.Sel("occlusionState"))
	return rv
}/* debug [instance_properties/getter]: occlusionState */


// A Boolean value that indicates the preferred location for the window’s backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/preferredBackingLocation
func (w_ Window) PreferredBackingLocation() WindowBackingLocation {
	rv := objc.Send[WindowBackingLocation](w_.ID, objc.Sel("preferredBackingLocation"))
	return rv
}/* debug [instance_properties/getter]: preferredBackingLocation */


// A Boolean value that indicates the preferred location for the window’s backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/preferredBackingLocation
func (w_ Window) SetPreferredBackingLocation(value WindowBackingLocation) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferredBackingLocation:"), value)
}/* debug [instance_properties/setter]: preferredBackingLocation */


// A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/preservesContentDuringLiveResize
func (w_ Window) PreservesContentDuringLiveResize() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("preservesContentDuringLiveResize"))
	return rv
}/* debug [instance_properties/getter]: preservesContentDuringLiveResize */


// A Boolean value that indicates whether the window tries to optimize user-initiated resize operations by preserving the content of views that have not changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/preservesContentDuringLiveResize
func (w_ Window) SetPreservesContentDuringLiveResize(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreservesContentDuringLiveResize:"), value)
}/* debug [instance_properties/setter]: preservesContentDuringLiveResize */


// The path to the file of the window’s represented file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/representedFilename
func (w_ Window) RepresentedFilename() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("representedFilename"))
	return rv
}/* debug [instance_properties/getter]: representedFilename */


// The path to the file of the window’s represented file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/representedFilename
func (w_ Window) SetRepresentedFilename(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRepresentedFilename:"), value)
}/* debug [instance_properties/setter]: representedFilename */


// The URL of the file the window represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/representedURL
func (w_ Window) RepresentedURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](w_.ID, objc.Sel("representedURL"))
	return rv
}/* debug [instance_properties/getter]: representedURL */


// The URL of the file the window represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/representedURL
func (w_ Window) SetRepresentedURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRepresentedURL:"), value)
}/* debug [instance_properties/setter]: representedURL */


// The flags field of the event record for the mouse-down event that initiated the resizing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resizeFlags
func (w_ Window) ResizeFlags() EventModifierFlags {
	rv := objc.Send[EventModifierFlags](w_.ID, objc.Sel("resizeFlags"))
	return rv
}/* debug [instance_properties/getter]: resizeFlags */


// The window’s resizing increments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resizeIncrements
func (w_ Window) ResizeIncrements() Size /* not a class type */ {
	rv := objc.Send[Size](w_.ID, objc.Sel("resizeIncrements"))
	return rv
}/* debug [instance_properties/getter]: resizeIncrements */


// The window’s resizing increments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/resizeIncrements
func (w_ Window) SetResizeIncrements(value Size /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setResizeIncrements:"), value)
}/* debug [instance_properties/setter]: resizeIncrements */


// The restoration class associated with the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/restorationClass
func (w_ Window) RestorationClass() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("restorationClass"))
	return rv
}/* debug [instance_properties/getter]: restorationClass */


// The restoration class associated with the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/restorationClass
func (w_ Window) SetRestorationClass(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setRestorationClass:"), value)
}/* debug [instance_properties/setter]: restorationClass */


// The screen the window is on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/screen
func (w_ Window) Screen() IScreen {
	rv := objc.Send[Screen](w_.ID, objc.Sel("screen"))
	return rv
}/* debug [instance_properties/getter]: screen */


// A Boolean value that indicates the level of access other processes have to the window’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/sharingType-swift.property
func (w_ Window) SharingType() WindowSharingType {
	rv := objc.Send[WindowSharingType](w_.ID, objc.Sel("sharingType"))
	return rv
}/* debug [instance_properties/getter]: sharingType */


// A Boolean value that indicates the level of access other processes have to the window’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/sharingType-swift.property
func (w_ Window) SetSharingType(value WindowSharingType) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSharingType:"), value)
}/* debug [instance_properties/setter]: sharingType */


// The window to which the sheet is attached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/sheetParent
func (w_ Window) SheetParent() IWindow {
	rv := objc.Send[Window](w_.ID, objc.Sel("sheetParent"))
	return rv
}/* debug [instance_properties/getter]: sheetParent */


// An array of the sheets currently attached to the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/sheets
func (w_ Window) Sheets() []Window {
	rv := objc.Send[[]Window](w_.ID, objc.Sel("sheets"))
	return rv
}/* debug [instance_properties/getter]: sheets */


// A Boolean value that indicates whether the window’s resize indicator is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/showsResizeIndicator
func (w_ Window) ShowsResizeIndicator() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("showsResizeIndicator"))
	return rv
}/* debug [instance_properties/getter]: showsResizeIndicator */


// A Boolean value that indicates whether the window’s resize indicator is visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/showsResizeIndicator
func (w_ Window) SetShowsResizeIndicator(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShowsResizeIndicator:"), value)
}/* debug [instance_properties/setter]: showsResizeIndicator */


// Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/styleMask-swift.property
func (w_ Window) StyleMask() WindowStyleMask {
	rv := objc.Send[WindowStyleMask](w_.ID, objc.Sel("styleMask"))
	return rv
}/* debug [instance_properties/getter]: styleMask */


// Flags that describe the window’s current style, such as if it’s resizable or in full-screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/styleMask-swift.property
func (w_ Window) SetStyleMask(value WindowStyleMask) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setStyleMask:"), value)
}/* debug [instance_properties/setter]: styleMask */


// A secondary line of text that appears in the title bar of the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/subtitle
func (w_ Window) Subtitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("subtitle"))
	return rv
}/* debug [instance_properties/getter]: subtitle */


// A secondary line of text that appears in the title bar of the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/subtitle
func (w_ Window) SetSubtitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setSubtitle:"), value)
}/* debug [instance_properties/setter]: subtitle */


// An object that represents information about a window when it displays as a tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tab
func (w_ Window) Tab() IWindowTab {
	rv := objc.Send[WindowTab](w_.ID, objc.Sel("tab"))
	return rv
}/* debug [instance_properties/getter]: tab */


// A group of windows that display together as a tab group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabGroup
func (w_ Window) TabGroup() IWindowTabGroup {
	rv := objc.Send[WindowTabGroup](w_.ID, objc.Sel("tabGroup"))
	return rv
}/* debug [instance_properties/getter]: tabGroup */


// An array of windows that display as tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabbedWindows
func (w_ Window) TabbedWindows() []Window {
	rv := objc.Send[[]Window](w_.ID, objc.Sel("tabbedWindows"))
	return rv
}/* debug [instance_properties/getter]: tabbedWindows */


// A value that allows a group of related windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingIdentifier-swift.property
func (w_ Window) TabbingIdentifier() WindowTabbingIdentifier /* typedef */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("tabbingIdentifier"))
	return rv
}/* debug [instance_properties/getter]: tabbingIdentifier */


// A value that allows a group of related windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingIdentifier-swift.property
func (w_ Window) SetTabbingIdentifier(value WindowTabbingIdentifier /* typedef */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTabbingIdentifier:"), value)
}/* debug [instance_properties/setter]: tabbingIdentifier */


// A value that indicates when a window displays tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingMode-swift.property
func (w_ Window) TabbingMode() WindowTabbingMode {
	rv := objc.Send[WindowTabbingMode](w_.ID, objc.Sel("tabbingMode"))
	return rv
}/* debug [instance_properties/getter]: tabbingMode */


// A value that indicates when a window displays tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/tabbingMode-swift.property
func (w_ Window) SetTabbingMode(value WindowTabbingMode) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTabbingMode:"), value)
}/* debug [instance_properties/setter]: tabbingMode */


// The string that appears in the title bar of the window or the path to the represented file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/title
func (w_ Window) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The string that appears in the title bar of the window or the path to the represented file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/title
func (w_ Window) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// A value that indicates the visibility of the window’s title and title bar buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/titleVisibility-swift.property
func (w_ Window) TitleVisibility() WindowTitleVisibility {
	rv := objc.Send[WindowTitleVisibility](w_.ID, objc.Sel("titleVisibility"))
	return rv
}/* debug [instance_properties/getter]: titleVisibility */


// A value that indicates the visibility of the window’s title and title bar buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/titleVisibility-swift.property
func (w_ Window) SetTitleVisibility(value WindowTitleVisibility) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitleVisibility:"), value)
}/* debug [instance_properties/setter]: titleVisibility */


// A value that indicates the user’s preference for window tabbing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/userTabbingPreference-swift.type.property
func (w_ Window) UserTabbingPreference() WindowUserTabbingPreference {
	rv := objc.Send[WindowUserTabbingPreference](w_.ID, objc.Sel("userTabbingPreference"))
	return rv
}/* debug [instance_properties/getter]: userTabbingPreference */


// A Boolean value that indicates whether any of the window’s views need to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/viewsNeedDisplay
func (w_ Window) ViewsNeedDisplay() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("viewsNeedDisplay"))
	return rv
}/* debug [instance_properties/getter]: viewsNeedDisplay */


// A Boolean value that indicates whether any of the window’s views need to be displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/viewsNeedDisplay
func (w_ Window) SetViewsNeedDisplay(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setViewsNeedDisplay:"), value)
}/* debug [instance_properties/setter]: viewsNeedDisplay */


// The window number of the window’s window device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowNumber
func (w_ Window) WindowNumber() int {
	rv := objc.Send[int](w_.ID, objc.Sel("windowNumber"))
	return rv
}/* debug [instance_properties/getter]: windowNumber */


// The Carbon window reference associated with the window, creating one if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/windowRef
func (w_ Window) WindowRef() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](w_.ID, objc.Sel("windowRef"))
	return rv
}/* debug [instance_properties/getter]: windowRef */


// A Boolean value that indicates whether the window is able to receive keyboard and mouse events even when some other window is being run modally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/worksWhenModal
func (w_ Window) WorksWhenModal() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("worksWhenModal"))
	return rv
}/* debug [instance_properties/getter]: worksWhenModal */


// Returns the number of color components in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspacename/numberofcolorcomponents
func (w_ Window) NumberOfColorComponents() int {
	rv := objc.Send[int](w_.ID, objc.Sel("numberOfColorComponents"))
	return rv
}/* debug [instance_properties/getter]: numberOfColorComponents */


// Returns the number of color components in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscolorspacename/numberofcolorcomponents
func (w_ Window) SetNumberOfColorComponents(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setNumberOfColorComponents:"), value)
}/* debug [instance_properties/setter]: numberOfColorComponents */


// Returns the bits per pixel for the specified window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/depth/bitsperpixel
func (w_ Window) BitsPerPixel() int {
	rv := objc.Send[int](w_.ID, objc.Sel("bitsPerPixel"))
	return rv
}/* debug [instance_properties/getter]: bitsPerPixel */


// Returns the bits per pixel for the specified window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/depth/bitsperpixel
func (w_ Window) SetBitsPerPixel(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setBitsPerPixel:"), value)
}/* debug [instance_properties/setter]: bitsPerPixel */


// Returns the bits per sample for the specified window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/depth/bitspersample
func (w_ Window) BitsPerSample() int {
	rv := objc.Send[int](w_.ID, objc.Sel("bitsPerSample"))
	return rv
}/* debug [instance_properties/getter]: bitsPerSample */


// Returns the bits per sample for the specified window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/depth/bitspersample
func (w_ Window) SetBitsPerSample(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setBitsPerSample:"), value)
}/* debug [instance_properties/setter]: bitsPerSample */


// Returns the name of the color space corresponding to the passed window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/depth/colorspacename
func (w_ Window) ColorSpaceName() ColorSpaceName /* typedef */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("colorSpaceName"))
	return rv
}/* debug [instance_properties/getter]: colorSpaceName */


// Returns the name of the color space corresponding to the passed window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/depth/colorspacename
func (w_ Window) SetColorSpaceName(value ColorSpaceName /* typedef */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setColorSpaceName:"), value)
}/* debug [instance_properties/setter]: colorSpaceName */


// Returns whether the specified window depth is planar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/depth/isplanar
func (w_ Window) IsPlanar() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isPlanar"))
	return rv
}/* debug [instance_properties/getter]: isPlanar */


// Returns whether the specified window depth is planar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/depth/isplanar
func (w_ Window) SetIsPlanar(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsPlanar:"), value)
}/* debug [instance_properties/setter]: isPlanar */


// A Boolean value that indicates whether the window can display tooltips even when the application is in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/allowstooltipswhenapplicationisinactive
func (w_ Window) AllowsToolTipsWhenApplicationIsInactive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsToolTipsWhenApplicationIsInactive"))
	return rv
}/* debug [instance_properties/getter]: allowsToolTipsWhenApplicationIsInactive */


// A Boolean value that indicates whether the window can display tooltips even when the application is in the background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/allowstooltipswhenapplicationisinactive
func (w_ Window) SetAllowsToolTipsWhenApplicationIsInactive(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsToolTipsWhenApplicationIsInactive:"), value)
}/* debug [instance_properties/setter]: allowsToolTipsWhenApplicationIsInactive */


// An object that the window inherits its appearance from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/appearancesource
func (w_ Window) AppearanceSource() AppearanceCustomization /* not a class type */ {
	rv := objc.Send[AppearanceCustomization](w_.ID, objc.Sel("appearanceSource"))
	return rv
}/* debug [instance_properties/getter]: appearanceSource */


// An object that the window inherits its appearance from.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/appearancesource
func (w_ Window) SetAppearanceSource(value AppearanceCustomization /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAppearanceSource:"), value)
}/* debug [instance_properties/setter]: appearanceSource */


// A Boolean value that indicates whether the window can become the key window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/canbecomekey
func (w_ Window) CanBecomeKey() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canBecomeKey"))
	return rv
}/* debug [instance_properties/getter]: canBecomeKey */


// A Boolean value that indicates whether the window can become the key window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/canbecomekey
func (w_ Window) SetCanBecomeKey(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanBecomeKey:"), value)
}/* debug [instance_properties/setter]: canBecomeKey */


// A Boolean value that indicates whether the window can become the application’s main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/canbecomemain
func (w_ Window) CanBecomeMain() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("canBecomeMain"))
	return rv
}/* debug [instance_properties/getter]: canBecomeMain */


// A Boolean value that indicates whether the window can become the application’s main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/canbecomemain
func (w_ Window) SetCanBecomeMain(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCanBecomeMain:"), value)
}/* debug [instance_properties/setter]: canBecomeMain */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/cascadingreferenceframe
func (w_ Window) CascadingReferenceFrame() Rect /* not a class type */ {
	rv := objc.Send[Rect](w_.ID, objc.Sel("cascadingReferenceFrame"))
	return rv
}/* debug [instance_properties/getter]: cascadingReferenceFrame */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/cascadingreferenceframe
func (w_ Window) SetCascadingReferenceFrame(value Rect /* not a class type */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setCascadingReferenceFrame:"), value)
}/* debug [instance_properties/setter]: cascadingReferenceFrame */


// An array of the window’s attached child windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/childwindows
func (w_ Window) ChildWindows() IWindow {
	rv := objc.Send[Window](w_.ID, objc.Sel("childWindows"))
	return rv
}/* debug [instance_properties/getter]: childWindows */


// An array of the window’s attached child windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/childwindows
func (w_ Window) SetChildWindows(value IWindow) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setChildWindows:"), value)
}/* debug [instance_properties/setter]: childWindows */


// The button cell that performs as if clicked when the window receives a Return (or Enter) key event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/defaultbuttoncell
func (w_ Window) DefaultButtonCell() IButtonCell {
	rv := objc.Send[ButtonCell](w_.ID, objc.Sel("defaultButtonCell"))
	return rv
}/* debug [instance_properties/getter]: defaultButtonCell */


// The button cell that performs as if clicked when the window receives a Return (or Enter) key event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/defaultbuttoncell
func (w_ Window) SetDefaultButtonCell(value IButtonCell) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDefaultButtonCell:"), value)
}/* debug [instance_properties/setter]: defaultButtonCell */


// The window’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/delegate
func (w_ Window) Delegate() objc.IObject /* cross-framework: WindowDelegate */ {
	rv := objc.Send[objc.ID](w_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The window’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/delegate
func (w_ Window) SetDelegate(value objc.IObject /* cross-framework: WindowDelegate */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The application’s Dock tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/docktile
func (w_ Window) DockTile() IDockTile {
	rv := objc.Send[DockTile](w_.ID, objc.Sel("dockTile"))
	return rv
}/* debug [instance_properties/getter]: dockTile */


// The application’s Dock tile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/docktile
func (w_ Window) SetDockTile(value IDockTile) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setDockTile:"), value)
}/* debug [instance_properties/setter]: dockTile */


// A string representation of the window’s frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/framedescriptor
func (w_ Window) FrameDescriptor() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](w_.ID, objc.Sel("frameDescriptor"))
	return rv
}/* debug [instance_properties/getter]: frameDescriptor */


// A string representation of the window’s frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/framedescriptor
func (w_ Window) SetFrameDescriptor(value objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setFrameDescriptor:"), value)
}/* debug [instance_properties/setter]: frameDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/hasactivewindowsharingsession
func (w_ Window) HasActiveWindowSharingSession() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasActiveWindowSharingSession"))
	return rv
}/* debug [instance_properties/getter]: hasActiveWindowSharingSession */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/hasactivewindowsharingsession
func (w_ Window) SetHasActiveWindowSharingSession(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasActiveWindowSharingSession:"), value)
}/* debug [instance_properties/setter]: hasActiveWindowSharingSession */


// A Boolean value that indicates if the window has a close box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/hasclosebox
func (w_ Window) HasCloseBox() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasCloseBox"))
	return rv
}/* debug [instance_properties/getter]: hasCloseBox */


// A Boolean value that indicates if the window has a close box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/hasclosebox
func (w_ Window) SetHasCloseBox(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasCloseBox:"), value)
}/* debug [instance_properties/setter]: hasCloseBox */


// A Boolean value that indicates if the window has a title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/hastitlebar
func (w_ Window) HasTitleBar() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasTitleBar"))
	return rv
}/* debug [instance_properties/getter]: hasTitleBar */


// A Boolean value that indicates if the window has a title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/hastitlebar
func (w_ Window) SetHasTitleBar(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasTitleBar:"), value)
}/* debug [instance_properties/setter]: hasTitleBar */


// A Boolean value that indicates whether the window is being resized by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/inliveresize
func (w_ Window) InLiveResize() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("inLiveResize"))
	return rv
}/* debug [instance_properties/getter]: inLiveResize */


// A Boolean value that indicates whether the window is being resized by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/inliveresize
func (w_ Window) SetInLiveResize(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInLiveResize:"), value)
}/* debug [instance_properties/setter]: inLiveResize */


// A Boolean value that indicates whether the window’s document has been edited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isdocumentedited
func (w_ Window) IsDocumentEdited() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isDocumentEdited"))
	return rv
}/* debug [instance_properties/getter]: isDocumentEdited */


// A Boolean value that indicates whether the window’s document has been edited.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isdocumentedited
func (w_ Window) SetIsDocumentEdited(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsDocumentEdited:"), value)
}/* debug [instance_properties/setter]: isDocumentEdited */


// A Boolean value that indicates whether the window is excluded from the application’s Windows menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isexcludedfromwindowsmenu
func (w_ Window) IsExcludedFromWindowsMenu() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isExcludedFromWindowsMenu"))
	return rv
}/* debug [instance_properties/getter]: isExcludedFromWindowsMenu */


// A Boolean value that indicates whether the window is excluded from the application’s Windows menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isexcludedfromwindowsmenu
func (w_ Window) SetIsExcludedFromWindowsMenu(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsExcludedFromWindowsMenu:"), value)
}/* debug [instance_properties/setter]: isExcludedFromWindowsMenu */


// A Boolean value that indicates whether the window is a floating panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isfloatingpanel
func (w_ Window) IsFloatingPanel() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isFloatingPanel"))
	return rv
}/* debug [instance_properties/getter]: isFloatingPanel */


// A Boolean value that indicates whether the window is a floating panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isfloatingpanel
func (w_ Window) SetIsFloatingPanel(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsFloatingPanel:"), value)
}/* debug [instance_properties/setter]: isFloatingPanel */


// A Boolean value that indicates whether the window is the key window for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/iskeywindow
func (w_ Window) IsKeyWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isKeyWindow"))
	return rv
}/* debug [instance_properties/getter]: isKeyWindow */


// A Boolean value that indicates whether the window is the key window for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/iskeywindow
func (w_ Window) SetIsKeyWindow(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsKeyWindow:"), value)
}/* debug [instance_properties/setter]: isKeyWindow */


// A Boolean value that indicates whether the window is the application’s main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/ismainwindow
func (w_ Window) IsMainWindow() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isMainWindow"))
	return rv
}/* debug [instance_properties/getter]: isMainWindow */


// A Boolean value that indicates whether the window is the application’s main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/ismainwindow
func (w_ Window) SetIsMainWindow(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsMainWindow:"), value)
}/* debug [instance_properties/setter]: isMainWindow */


// A Boolean value that indicates whether the window can minimize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isminiaturizable
func (w_ Window) IsMiniaturizable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isMiniaturizable"))
	return rv
}/* debug [instance_properties/getter]: isMiniaturizable */


// A Boolean value that indicates whether the window can minimize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isminiaturizable
func (w_ Window) SetIsMiniaturizable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsMiniaturizable:"), value)
}/* debug [instance_properties/setter]: isMiniaturizable */


// A Boolean value that indicates whether the window is minimized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isminiaturized
func (w_ Window) IsMiniaturized() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isMiniaturized"))
	return rv
}/* debug [instance_properties/getter]: isMiniaturized */


// A Boolean value that indicates whether the window is minimized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isminiaturized
func (w_ Window) SetIsMiniaturized(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsMiniaturized:"), value)
}/* debug [instance_properties/setter]: isMiniaturized */


// A Boolean value that indicates whether the window is a modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/ismodalpanel
func (w_ Window) IsModalPanel() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isModalPanel"))
	return rv
}/* debug [instance_properties/getter]: isModalPanel */


// A Boolean value that indicates whether the window is a modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/ismodalpanel
func (w_ Window) SetIsModalPanel(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsModalPanel:"), value)
}/* debug [instance_properties/setter]: isModalPanel */


// A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/ismovable
func (w_ Window) IsMovable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isMovable"))
	return rv
}/* debug [instance_properties/getter]: isMovable */


// A Boolean value that indicates whether the window can be dragged by clicking in its title bar or background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/ismovable
func (w_ Window) SetIsMovable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsMovable:"), value)
}/* debug [instance_properties/setter]: isMovable */


// A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/ismovablebywindowbackground
func (w_ Window) IsMovableByWindowBackground() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isMovableByWindowBackground"))
	return rv
}/* debug [instance_properties/getter]: isMovableByWindowBackground */


// A Boolean value that indicates whether the window is movable by clicking and dragging anywhere in its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/ismovablebywindowbackground
func (w_ Window) SetIsMovableByWindowBackground(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsMovableByWindowBackground:"), value)
}/* debug [instance_properties/setter]: isMovableByWindowBackground */


// A Boolean value that indicates whether the window is on the currently active space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isonactivespace
func (w_ Window) IsOnActiveSpace() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isOnActiveSpace"))
	return rv
}/* debug [instance_properties/getter]: isOnActiveSpace */


// A Boolean value that indicates whether the window is on the currently active space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isonactivespace
func (w_ Window) SetIsOnActiveSpace(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsOnActiveSpace:"), value)
}/* debug [instance_properties/setter]: isOnActiveSpace */


// A Boolean value that indicates whether the window is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isopaque
func (w_ Window) IsOpaque() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isOpaque"))
	return rv
}/* debug [instance_properties/getter]: isOpaque */


// A Boolean value that indicates whether the window is opaque.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isopaque
func (w_ Window) SetIsOpaque(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsOpaque:"), value)
}/* debug [instance_properties/setter]: isOpaque */


// A Boolean value that indicates whether the window is released when it receives the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isreleasedwhenclosed
func (w_ Window) IsReleasedWhenClosed() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isReleasedWhenClosed"))
	return rv
}/* debug [instance_properties/getter]: isReleasedWhenClosed */


// A Boolean value that indicates whether the window is released when it receives the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isreleasedwhenclosed
func (w_ Window) SetIsReleasedWhenClosed(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsReleasedWhenClosed:"), value)
}/* debug [instance_properties/setter]: isReleasedWhenClosed */


// A Boolean value that indicates if the user can resize the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isresizable
func (w_ Window) IsResizable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isResizable"))
	return rv
}/* debug [instance_properties/getter]: isResizable */


// A Boolean value that indicates if the user can resize the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isresizable
func (w_ Window) SetIsResizable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsResizable:"), value)
}/* debug [instance_properties/setter]: isResizable */


// A Boolean value indicating whether the window configuration is preserved between application launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isrestorable
func (w_ Window) IsRestorable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isRestorable"))
	return rv
}/* debug [instance_properties/getter]: isRestorable */


// A Boolean value indicating whether the window configuration is preserved between application launches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isrestorable
func (w_ Window) SetIsRestorable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsRestorable:"), value)
}/* debug [instance_properties/setter]: isRestorable */


// A Boolean value that indicates whether the window has ever run as a modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/issheet
func (w_ Window) IsSheet() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isSheet"))
	return rv
}/* debug [instance_properties/getter]: isSheet */


// A Boolean value that indicates whether the window has ever run as a modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/issheet
func (w_ Window) SetIsSheet(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsSheet:"), value)
}/* debug [instance_properties/setter]: isSheet */


// A Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isvisible
func (w_ Window) IsVisible() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isVisible"))
	return rv
}/* debug [instance_properties/getter]: isVisible */


// A Boolean value that indicates whether the window is visible onscreen (even when it’s obscured by other windows).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/isvisible
func (w_ Window) SetIsVisible(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsVisible:"), value)
}/* debug [instance_properties/setter]: isVisible */


// A Boolean value that indicates whether the window allows zooming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/iszoomable
func (w_ Window) IsZoomable() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isZoomable"))
	return rv
}/* debug [instance_properties/getter]: isZoomable */


// A Boolean value that indicates whether the window allows zooming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/iszoomable
func (w_ Window) SetIsZoomable(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsZoomable:"), value)
}/* debug [instance_properties/setter]: isZoomable */


// A Boolean value that indicates whether the window is in a zoomed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/iszoomed
func (w_ Window) IsZoomed() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isZoomed"))
	return rv
}/* debug [instance_properties/getter]: isZoomed */


// A Boolean value that indicates whether the window is in a zoomed state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/iszoomed
func (w_ Window) SetIsZoomed(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsZoomed:"), value)
}/* debug [instance_properties/setter]: isZoomed */


// The zero-based position of the window, based on its order from front to back among all visible application windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/orderedindex
func (w_ Window) OrderedIndex() int {
	rv := objc.Send[int](w_.ID, objc.Sel("orderedIndex"))
	return rv
}/* debug [instance_properties/getter]: orderedIndex */


// The zero-based position of the window, based on its order from front to back among all visible application windows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/orderedindex
func (w_ Window) SetOrderedIndex(value int) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setOrderedIndex:"), value)
}/* debug [instance_properties/setter]: orderedIndex */


// The parent window to which the window is attached as a child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/parent
func (w_ Window) Parent() IWindow {
	rv := objc.Send[Window](w_.ID, objc.Sel("parent"))
	return rv
}/* debug [instance_properties/getter]: parent */


// The parent window to which the window is attached as a child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/parent
func (w_ Window) SetParent(value IWindow) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setParent:"), value)
}/* debug [instance_properties/setter]: parent */


// A Boolean value that indicates whether the window prevents application termination when modal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/preventsapplicationterminationwhenmodal
func (w_ Window) PreventsApplicationTerminationWhenModal() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("preventsApplicationTerminationWhenModal"))
	return rv
}/* debug [instance_properties/getter]: preventsApplicationTerminationWhenModal */


// A Boolean value that indicates whether the window prevents application termination when modal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/preventsapplicationterminationwhenmodal
func (w_ Window) SetPreventsApplicationTerminationWhenModal(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreventsApplicationTerminationWhenModal:"), value)
}/* debug [instance_properties/setter]: preventsApplicationTerminationWhenModal */


// A Boolean value that indicates whether the toolbar control button is currently displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/showstoolbarbutton
func (w_ Window) ShowsToolbarButton() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("showsToolbarButton"))
	return rv
}/* debug [instance_properties/getter]: showsToolbarButton */


// A Boolean value that indicates whether the toolbar control button is currently displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/showstoolbarbutton
func (w_ Window) SetShowsToolbarButton(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setShowsToolbarButton:"), value)
}/* debug [instance_properties/setter]: showsToolbarButton */


// An array of title bar accessory view controllers that are currently added to the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/titlebaraccessoryviewcontrollers
func (w_ Window) TitlebarAccessoryViewControllers() ITitlebarAccessoryViewController {
	rv := objc.Send[TitlebarAccessoryViewController](w_.ID, objc.Sel("titlebarAccessoryViewControllers"))
	return rv
}/* debug [instance_properties/getter]: titlebarAccessoryViewControllers */


// An array of title bar accessory view controllers that are currently added to the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/titlebaraccessoryviewcontrollers
func (w_ Window) SetTitlebarAccessoryViewControllers(value ITitlebarAccessoryViewController) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitlebarAccessoryViewControllers:"), value)
}/* debug [instance_properties/setter]: titlebarAccessoryViewControllers */


// A Boolean value that indicates whether the title bar draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/titlebarappearstransparent
func (w_ Window) TitlebarAppearsTransparent() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("titlebarAppearsTransparent"))
	return rv
}/* debug [instance_properties/getter]: titlebarAppearsTransparent */


// A Boolean value that indicates whether the title bar draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/titlebarappearstransparent
func (w_ Window) SetTitlebarAppearsTransparent(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitlebarAppearsTransparent:"), value)
}/* debug [instance_properties/setter]: titlebarAppearsTransparent */


// The type of separator that the app displays between the title bar and content of a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/titlebarseparatorstyle
func (w_ Window) TitlebarSeparatorStyle() TitlebarSeparatorStyle {
	rv := objc.Send[TitlebarSeparatorStyle](w_.ID, objc.Sel("titlebarSeparatorStyle"))
	return rv
}/* debug [instance_properties/getter]: titlebarSeparatorStyle */


// The type of separator that the app displays between the title bar and content of a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/titlebarseparatorstyle
func (w_ Window) SetTitlebarSeparatorStyle(value TitlebarSeparatorStyle) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setTitlebarSeparatorStyle:"), value)
}/* debug [instance_properties/setter]: titlebarSeparatorStyle */


// The window’s toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/toolbar
func (w_ Window) Toolbar() IToolbar {
	rv := objc.Send[Toolbar](w_.ID, objc.Sel("toolbar"))
	return rv
}/* debug [instance_properties/getter]: toolbar */


// The window’s toolbar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/toolbar
func (w_ Window) SetToolbar(value IToolbar) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setToolbar:"), value)
}/* debug [instance_properties/setter]: toolbar */


// The style that determines the appearance and location of the toolbar in relation to the title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/toolbarstyle-swift.property
func (w_ Window) ToolbarStyle() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](w_.ID, objc.Sel("toolbarStyle"))
	return rv
}/* debug [instance_properties/getter]: toolbarStyle */


// The style that determines the appearance and location of the toolbar in relation to the title bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/toolbarstyle-swift.property
func (w_ Window) SetToolbarStyle(value objectivec.IObject) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setToolbarStyle:"), value)
}/* debug [instance_properties/setter]: toolbarStyle */


// The window’s window controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/windowcontroller
func (w_ Window) WindowController() IWindowController {
	rv := objc.Send[WindowController](w_.ID, objc.Sel("windowController"))
	return rv
}/* debug [instance_properties/getter]: windowController */


// The window’s window controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/windowcontroller
func (w_ Window) SetWindowController(value IWindowController) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowController:"), value)
}/* debug [instance_properties/setter]: windowController */


// The direction the window’s title bar lays text out, either left to right or right to left.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/windowtitlebarlayoutdirection
func (w_ Window) WindowTitlebarLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](w_.ID, objc.Sel("windowTitlebarLayoutDirection"))
	return rv
}/* debug [instance_properties/getter]: windowTitlebarLayoutDirection */


// The direction the window’s title bar lays text out, either left to right or right to left.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nswindow/windowtitlebarlayoutdirection
func (w_ Window) SetWindowTitlebarLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setWindowTitlebarLayoutDirection:"), value)
}/* debug [instance_properties/setter]: windowTitlebarLayoutDirection */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSWindow */


