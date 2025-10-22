// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/coreimage"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/quartzcore"
)

// The class instance for the [View] class.
var (
	ViewClass     _ViewClass
	ViewClassOnce sync.Once
)

func getViewClass() _ViewClass {
	ViewClassOnce.Do(func() {
		ViewClass = _ViewClass{objc.GetClass("NSView")}
	})
	return ViewClass
}

type _ViewClass struct {
	class objc.Class
}

// An interface definition for the [View] class.
type IView interface {
	IResponder
	AcceptsFirstMouse(event IEvent) bool
	AddConstraint(constraint ILayoutConstraint)
	AddConstraints(constraints []LayoutConstraint)
	AddCursorRectCursor(rect coregraphics.CGRect, object ICursor)
	AddGestureRecognizer(gestureRecognizer IGestureRecognizer)
	AddLayoutGuide(guide ILayoutGuide)
	AddSubview(view IView)
	AddSubviewPositionedRelativeTo(view IView, place WindowOrderingMode, otherView IView)
	AddToolTipRectOwnerUserData(rect coregraphics.CGRect, owner objectivec.IObject, data unsafe.Pointer) ToolTipTag
	AddTrackingArea(trackingArea ITrackingArea)
	AddTrackingRectOwnerUserDataAssumeInside(rect coregraphics.CGRect, owner objectivec.IObject, data unsafe.Pointer, flag bool) TrackingRectTag
	AdjustPageHeightNewTopBottomLimit(newBottom coregraphics.float64, oldTop float64, oldBottom float64, bottomLimit float64)
	AdjustPageWidthNewLeftRightLimit(newRight coregraphics.float64, oldLeft float64, oldRight float64, rightLimit float64)
	AdjustScroll(newVisible coregraphics.CGRect) coregraphics.CGRect
	AlignmentRectForFrame(frame coregraphics.CGRect) coregraphics.CGRect
	AllocateGState()
	AncestorSharedWithView(view IView) View
	Autoscroll(event IEvent) bool
	BackingAlignedRectOptions(rect coregraphics.CGRect, options unsafe.Pointer) coregraphics.CGRect
	BeginDocument()
	BeginDraggingSessionWithItemsEventSource(items []DraggingItem, event IEvent, source objectivec.IObject) DraggingSession
	BeginPageInRectAtPlacement(rect coregraphics.CGRect, location coregraphics.CGPoint)
	BitmapImageRepForCachingDisplayInRect(rect coregraphics.CGRect) BitmapImageRep
	CacheDisplayInRectToBitmapImageRep(rect coregraphics.CGRect, bitmapImageRep IBitmapImageRep)
	CenterScanRect(rect coregraphics.CGRect) coregraphics.CGRect
	ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint
	ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority
	ConvertPointFromView(point coregraphics.CGPoint, view IView) coregraphics.CGPoint
	ConvertSizeFromView(size coregraphics.CGSize, view IView) coregraphics.CGSize
	ConvertRectFromView(rect coregraphics.CGRect, view IView) coregraphics.CGRect
	ConvertRectToView(rect coregraphics.CGRect, view IView) coregraphics.CGRect
	ConvertSizeToView(size coregraphics.CGSize, view IView) coregraphics.CGSize
	ConvertPointToView(point coregraphics.CGPoint, view IView) coregraphics.CGPoint
	ConvertPointFromBacking(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertRectFromBacking(rect coregraphics.CGRect) coregraphics.CGRect
	ConvertSizeFromBacking(size coregraphics.CGSize) coregraphics.CGSize
	ConvertPointFromLayer(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertSizeFromLayer(size coregraphics.CGSize) coregraphics.CGSize
	ConvertRectFromLayer(rect coregraphics.CGRect) coregraphics.CGRect
	ConvertPointFromBase(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertPointToBase(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertRectFromBase(rect coregraphics.CGRect) coregraphics.CGRect
	ConvertRectToBase(rect coregraphics.CGRect) coregraphics.CGRect
	ConvertSizeFromBase(size coregraphics.CGSize) coregraphics.CGSize
	ConvertSizeToBase(size coregraphics.CGSize) coregraphics.CGSize
	ConvertPointToBacking(point coregraphics.CGPoint) coregraphics.CGPoint
	ConvertRectToBacking(rect coregraphics.CGRect) coregraphics.CGRect
	ConvertSizeToBacking(size coregraphics.CGSize) coregraphics.CGSize
	ConvertRectToLayer(rect coregraphics.CGRect) coregraphics.CGRect
	ConvertSizeToLayer(size coregraphics.CGSize) coregraphics.CGSize
	ConvertPointToLayer(point coregraphics.CGPoint) coregraphics.CGPoint
	DataWithEPSInsideRect(rect coregraphics.CGRect) foundation.Data
	DataWithPDFInsideRect(rect coregraphics.CGRect) foundation.Data
	DidAddSubview(subview IView)
	DidCloseMenuWithEvent(menu IMenu, event IEvent)
	DiscardCursorRects()
	Display()
	DisplayRect(rect coregraphics.CGRect)
	DisplayIfNeeded()
	DisplayIfNeededInRect(rect coregraphics.CGRect)
	DisplayIfNeededIgnoringOpacity()
	DisplayIfNeededInRectIgnoringOpacity(rect coregraphics.CGRect)
	DisplayRectIgnoringOpacity(rect coregraphics.CGRect)
	DisplayRectIgnoringOpacityInContext(rect coregraphics.CGRect, context IGraphicsContext)
	DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) quartzcore.DisplayLink
	DragFileFromRectSlideBackEvent(filename string, rect coregraphics.CGRect, flag bool, event IEvent) bool
	DragImageAtOffsetEventPasteboardSourceSlideBack(image IImage, viewLocation coregraphics.CGPoint, initialOffset coregraphics.CGSize, event IEvent, pboard IPasteboard, sourceObj objectivec.IObject, slideFlag bool)
	DragPromisedFilesOfTypesFromRectSourceSlideBackEvent(typeArray []string, rect coregraphics.CGRect, sourceObject objectivec.IObject, flag bool, event IEvent) bool
	DrawRect(dirtyRect coregraphics.CGRect)
	DrawFocusRingMask()
	DrawPageBorderWithSize(borderSize coregraphics.CGSize)
	DrawSheetBorderWithSize(borderSize coregraphics.CGSize)
	EdgeInsetsForLayoutRegion(layoutRegion IViewLayoutRegion) unsafe.Pointer
	EndDocument()
	EndPage()
	EnterFullScreenModeWithOptions(screen IScreen, options unsafe.Pointer) bool
	ExerciseAmbiguityInLayout()
	ExitFullScreenModeWithOptions(options unsafe.Pointer)
	FrameForAlignmentRect(alignmentRect coregraphics.CGRect) coregraphics.CGRect
	GState() int
	GetRectsBeingDrawnCount(rects coregraphics.CGRect, count unsafe.Pointer)
	GetRectsExposedDuringLiveResizeCount(exposedRects unsafe.Pointer, count unsafe.Pointer)
	HitTest(point coregraphics.CGPoint) View
	InvalidateIntrinsicContentSize()
	IsDescendantOf(view IView) bool
	MouseInRect(point coregraphics.CGPoint, rect coregraphics.CGRect) bool
	KnowsPageRange(range_ unsafe.Pointer) bool
	Layout()
	LayoutGuideForLayoutRegion(layoutRegion IViewLayoutRegion) LayoutGuide
	LayoutSubtreeIfNeeded()
	LocationOfPrintRect(rect coregraphics.CGRect) coregraphics.CGPoint
	LockFocus()
	LockFocusIfCanDraw() bool
	LockFocusIfCanDrawInContext(context IGraphicsContext) bool
	MakeBackingLayer() quartzcore.Layer
	MenuForEvent(event IEvent) Menu
	NeedsToDrawRect(rect coregraphics.CGRect) bool
	NoteFocusRingMaskChanged()
	PerformKeyEquivalent(event IEvent) bool
	PerformMnemonic(string_ string) bool
	PrepareContentInRect(rect coregraphics.CGRect)
	PrepareForReuse()
	Print(sender objectivec.IObject)
	RectForLayoutRegion(layoutRegion IViewLayoutRegion) coregraphics.CGRect
	RectForPage(page int) coregraphics.CGRect
	RectForSmartMagnificationAtPointInRect(location coregraphics.CGPoint, visibleRect coregraphics.CGRect) coregraphics.CGRect
	ReflectScrolledClipView(clipView IClipView)
	RegisterForDraggedTypes(newTypes []string)
	ReleaseGState()
	RemoveAllToolTips()
	RemoveConstraint(constraint ILayoutConstraint)
	RemoveConstraints(constraints []LayoutConstraint)
	RemoveCursorRectCursor(rect coregraphics.CGRect, object ICursor)
	RemoveFromSuperview()
	RemoveFromSuperviewWithoutNeedingDisplay()
	RemoveGestureRecognizer(gestureRecognizer IGestureRecognizer)
	RemoveLayoutGuide(guide ILayoutGuide)
	RemoveToolTip(tag IToolTipTag)
	RemoveTrackingArea(trackingArea ITrackingArea)
	RemoveTrackingRect(tag ITrackingRectTag)
	RenewGState()
	ReplaceSubviewWith(oldView IView, newView IView)
	ResetCursorRects()
	ResizeWithOldSuperviewSize(oldSize coregraphics.CGSize)
	ResizeSubviewsWithOldSize(oldSize coregraphics.CGSize)
	RotateByAngle(angle float64)
	RulerViewDidAddMarker(ruler IRulerView, marker IRulerMarker)
	RulerViewDidMoveMarker(ruler IRulerView, marker IRulerMarker)
	RulerViewDidRemoveMarker(ruler IRulerView, marker IRulerMarker)
	RulerViewHandleMouseDown(ruler IRulerView, event IEvent)
	RulerViewLocationForPoint(ruler IRulerView, point coregraphics.CGPoint) float64
	RulerViewPointForLocation(ruler IRulerView, point float64) coregraphics.CGPoint
	RulerViewShouldAddMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerViewShouldMoveMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerViewShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerViewWillAddMarkerAtLocation(ruler IRulerView, marker IRulerMarker, location float64) float64
	RulerViewWillMoveMarkerToLocation(ruler IRulerView, marker IRulerMarker, location float64) float64
	RulerViewWillSetClientView(ruler IRulerView, newClient IView)
	ScaleUnitSquareToSize(newUnitSize coregraphics.CGSize)
	ScrollPoint(point coregraphics.CGPoint)
	ScrollRectBy(rect coregraphics.CGRect, delta coregraphics.CGSize)
	ScrollClipViewToPoint(clipView IClipView, point coregraphics.CGPoint)
	ScrollRectToVisible(rect coregraphics.CGRect) bool
	SetBoundsOrigin(newOrigin coregraphics.CGPoint)
	SetBoundsSize(newSize coregraphics.CGSize)
	SetContentCompressionResistancePriorityForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation)
	SetContentHuggingPriorityForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation)
	SetFrameOrigin(newOrigin coregraphics.CGPoint)
	SetFrameSize(newSize coregraphics.CGSize)
	SetKeyboardFocusRingNeedsDisplayInRect(rect coregraphics.CGRect)
	SetNeedsDisplayInRect(invalidRect coregraphics.CGRect)
	SetUpGState()
	ShouldDelayWindowOrderingForEvent(event IEvent) bool
	ShouldDrawColor() bool
	ShowDefinitionForAttributedStringAtPoint(attrString foundation.IAttributedString, textBaselineOrigin coregraphics.CGPoint)
	ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString foundation.IAttributedString, targetRange foundation.Range, options unsafe.Pointer, originProvider unsafe.Pointer)
	SortSubviewsUsingFunctionContext(compare unsafe.Pointer, context unsafe.Pointer)
	TranslateOriginToPoint(translation coregraphics.CGPoint)
	TranslateRectsNeedingDisplayInRectBy(clipRect coregraphics.CGRect, delta coregraphics.CGSize)
	UnlockFocus()
	UnregisterDraggedTypes()
	UpdateConstraints()
	UpdateConstraintsForSubtreeIfNeeded()
	UpdateLayer()
	UpdateTrackingAreas()
	ViewDidChangeBackingProperties()
	ViewDidChangeEffectiveAppearance()
	ViewDidEndLiveResize()
	ViewDidHide()
	ViewDidMoveToSuperview()
	ViewDidMoveToWindow()
	ViewDidUnhide()
	ViewWillDraw()
	ViewWillMoveToSuperview(newSuperview IView)
	ViewWillMoveToWindow(newWindow IWindow)
	ViewWillStartLiveResize()
	ViewWithTag(tag int) View
	WillOpenMenuWithEvent(menu IMenu, event IEvent)
	WillRemoveSubview(subview IView)
	WriteEPSInsideRectToPasteboard(rect coregraphics.CGRect, pasteboard IPasteboard)
	WritePDFInsideRectToPasteboard(rect coregraphics.CGRect, pasteboard IPasteboard)
	AcceptsTouchEvents() bool
	SetAcceptsTouchEvents(value bool)
	AdditionalSafeAreaInsets() unsafe.Pointer
	SetAdditionalSafeAreaInsets(value unsafe.Pointer)
	AlignmentRectInsets() unsafe.Pointer
	AllowedTouchTypes() TouchTypeMask
	SetAllowedTouchTypes(value TouchTypeMask)
	AllowsVibrancy() bool
	AlphaValue() float64
	SetAlphaValue(value float64)
	AutoresizesSubviews() bool
	SetAutoresizesSubviews(value bool)
	AutoresizingMask() AutoresizingMaskOptions
	SetAutoresizingMask(value AutoresizingMaskOptions)
	BackgroundFilters() []coreimage.Filter
	SetBackgroundFilters(value []coreimage.IFilter)
	BaselineOffsetFromBottom() float64
	BottomAnchor() NSLayoutYAxisAnchor
	Bounds() coregraphics.CGRect
	SetBounds(value coregraphics.CGRect)
	BoundsRotation() float64
	SetBoundsRotation(value float64)
	CanBecomeKeyView() bool
	CanDraw() bool
	CanDrawConcurrently() bool
	SetCanDrawConcurrently(value bool)
	CanDrawSubviewsIntoLayer() bool
	SetCanDrawSubviewsIntoLayer(value bool)
	CandidateListTouchBarItem() NSCandidateListTouchBarItem
	CenterXAnchor() NSLayoutXAxisAnchor
	CenterYAnchor() NSLayoutYAxisAnchor
	ClipsToBounds() bool
	SetClipsToBounds(value bool)
	CompositingFilter() coreimage.Filter
	SetCompositingFilter(value coreimage.IFilter)
	Constraints() []LayoutConstraint
	ContentFilters() []coreimage.Filter
	SetContentFilters(value []coreimage.IFilter)
	EnclosingMenuItem() NSMenuItem
	EnclosingScrollView() NSScrollView
	FirstBaselineAnchor() NSLayoutYAxisAnchor
	FirstBaselineOffsetFromTop() float64
	FittingSize() coregraphics.CGSize
	FocusRingMaskBounds() coregraphics.CGRect
	FocusRingType() FocusRingType
	SetFocusRingType(value FocusRingType)
	Frame() coregraphics.CGRect
	SetFrame(value coregraphics.CGRect)
	FrameCenterRotation() float64
	SetFrameCenterRotation(value float64)
	FrameRotation() float64
	SetFrameRotation(value float64)
	GestureRecognizers() []GestureRecognizer
	SetGestureRecognizers(value []GestureRecognizer)
	HasAmbiguousLayout() bool
	HeightAdjustLimit() float64
	HeightAnchor() NSLayoutDimension
	InLiveResize() bool
	InputContext() NSTextInputContext
	IntrinsicContentSize() coregraphics.CGSize
	DrawingFindIndicator() bool
	Flipped() bool
	Hidden() bool
	SetHidden(value bool)
	HiddenOrHasHiddenAncestor() bool
	HorizontalContentSizeConstraintActive() bool
	SetHorizontalContentSizeConstraintActive(value bool)
	InFullScreenMode() bool
	Opaque() bool
	RotatedFromBase() bool
	RotatedOrScaledFromBase() bool
	VerticalContentSizeConstraintActive() bool
	SetVerticalContentSizeConstraintActive(value bool)
	LastBaselineAnchor() NSLayoutYAxisAnchor
	LastBaselineOffsetFromBottom() float64
	Layer() quartzcore.Layer
	SetLayer(value quartzcore.ILayer)
	LayerContentsPlacement() ViewLayerContentsPlacement
	SetLayerContentsPlacement(value IViewLayerContentsPlacement)
	LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy
	SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy)
	LayerUsesCoreImageFilters() bool
	SetLayerUsesCoreImageFilters(value bool)
	LayoutGuides() []LayoutGuide
	LayoutMarginsGuide() NSLayoutGuide
	LeadingAnchor() NSLayoutXAxisAnchor
	LeftAnchor() NSLayoutXAxisAnchor
	MouseDownCanMoveWindow() bool
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
	NeedsLayout() bool
	SetNeedsLayout(value bool)
	NeedsPanelToBecomeKey() bool
	NeedsUpdateConstraints() bool
	SetNeedsUpdateConstraints(value bool)
	NextKeyView() NSView
	SetNextKeyView(value IView)
	NextValidKeyView() NSView
	OpaqueAncestor() NSView
	PageFooter() foundation.AttributedString
	PageHeader() foundation.AttributedString
	PostsBoundsChangedNotifications() bool
	SetPostsBoundsChangedNotifications(value bool)
	PostsFrameChangedNotifications() bool
	SetPostsFrameChangedNotifications(value bool)
	PrefersCompactControlSizeMetrics() bool
	SetPrefersCompactControlSizeMetrics(value bool)
	PreparedContentRect() coregraphics.CGRect
	SetPreparedContentRect(value coregraphics.CGRect)
	PreservesContentDuringLiveResize() bool
	PressureConfiguration() NSPressureConfiguration
	SetPressureConfiguration(value IPressureConfiguration)
	PreviousKeyView() NSView
	PreviousValidKeyView() NSView
	PrintJobTitle() string
	RectPreservedDuringLiveResize() coregraphics.CGRect
	RegisteredDraggedTypes() []string
	RightAnchor() NSLayoutXAxisAnchor
	SafeAreaInsets() unsafe.Pointer
	SafeAreaLayoutGuide() NSLayoutGuide
	SafeAreaRect() coregraphics.CGRect
	Shadow() NSShadow
	SetShadow(value IShadow)
	Subviews() []View
	SetSubviews(value []View)
	Superview() NSView
	Tag() int
	ToolTip() string
	SetToolTip(value string)
	TopAnchor() NSLayoutYAxisAnchor
	TrackingAreas() []TrackingArea
	TrailingAnchor() NSLayoutXAxisAnchor
	TranslatesAutoresizingMaskIntoConstraints() bool
	SetTranslatesAutoresizingMaskIntoConstraints(value bool)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	VisibleRect() coregraphics.CGRect
	WantsBestResolutionOpenGLSurface() bool
	SetWantsBestResolutionOpenGLSurface(value bool)
	WantsDefaultClipping() bool
	WantsExtendedDynamicRangeOpenGLSurface() bool
	SetWantsExtendedDynamicRangeOpenGLSurface(value bool)
	WantsLayer() bool
	SetWantsLayer(value bool)
	WantsRestingTouches() bool
	SetWantsRestingTouches(value bool)
	WantsUpdateLayer() bool
	WidthAdjustLimit() float64
	WidthAnchor() NSLayoutDimension
	Window() NSWindow
	WritingToolsCoordinator() NSWritingToolsCoordinator
	SetWritingToolsCoordinator(value IWritingToolsCoordinator)
}

// The infrastructure for drawing, printing, and handling events in an app.
//
// You typically don’t use objects directly. Instead, you use objects that descend from or you subclass yourself and override its methods to implement the behavior you need. An instance of the class (or one of its subclasses) is commonly known as a view object, or simply as a view. Views handle the presentation and interaction with your app’s visible content. You arrange one or more views inside an object, which acts as a wrapper for your content. A view object defines a rectangular region for drawing and receiving mouse events. Views handle other chores as well, including the dragging of icons and working with the class to support efficient scrolling. AppKit handles most of your app’s management. Unless you’re implementing a concrete subclass of or working intimately with the content of the view hierarchy at runtime, you don’t need to know much about this class’s interface. For any view, there are many methods that you can use as-is. The following methods are commonly used. returns the location and size of the object. returns the internal origin and size of the object. determines whether the object needs to be redrawn. returns the object that contains the object. draws the object. (All subclasses must implement this method, but it’s rarely invoked explicitly.) An alternative to drawing is to update the layer directly using the method. For more information on how instances handle event and action messages, see . For more information on displaying tooltips and contextual menus, see and .


// The infrastructure for drawing, printing, and handling events in an app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView

type View struct {
	Responder
}

// ViewFrom constructs a [View] from an unsafe.Pointer.
//
// The infrastructure for drawing, printing, and handling events in an app.
func ViewFrom(ptr unsafe.Pointer) View {
	return View{
		Responder: ResponderFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc _ViewClass) Alloc() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _ViewClass) New() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ View) Init() View {
	rv := objc.Send[View](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ View) Autorelease() View {
	rv := objc.Send[View](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewView creates a new View instance.
func NewView() View {
	return getViewClass().New()
}




// Initializes a view using from data in the specified coder object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)

func NewViewWithCoder(coder foundation.ICoder) View {
	instance := getViewClass().Alloc()
	rv := objc.Send[View](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Initializes and returns a newly allocated object with a specified frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)

func NewViewWithFrame(frameRect coregraphics.CGRect) View {
	instance := getViewClass().Alloc()
	rv := objc.Send[View](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}



// Returns the default focus ring type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/defaultFocusRingType

func (vc _ViewClass) DefaultFocusRingType() FocusRingType {
	rv := objc.Send[FocusRingType](objc.ID(vc.class), objc.Sel("defaultFocusRingType"))
	return rv
}

// Overridden by subclasses to return the default pop-up menu for instances of the receiving class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/defaultMenu

func (vc _ViewClass) DefaultMenu() NSMenu {
	rv := objc.Send[NSMenu](objc.ID(vc.class), objc.Sel("defaultMenu"))
	return rv
}

// The currently focused view object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/focusView

func (vc _ViewClass) FocusView() View {
	rv := objc.Send[NSView](objc.ID(vc.class), objc.Sel("focusView"))
	return rv
}

// A Boolean value that indicates whether views support responsive scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isCompatibleWithResponsiveScrolling

func (vc _ViewClass) CompatibleWithResponsiveScrolling() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("compatibleWithResponsiveScrolling"))
	return rv
}

// Returns a Boolean value indicating whether the view depends on the constraint-based layout system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/requiresConstraintBasedLayout

func (vc _ViewClass) RequiresConstraintBasedLayout() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("requiresConstraintBasedLayout"))
	return rv
}


// Overridden by subclasses to return if the view should be sent a message for an initial mouse-down event, if not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/acceptsFirstMouse(for:)

func (v_ View) AcceptsFirstMouse(event IEvent) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("acceptsFirstMouse:"), event)
	return rv
}



// Adds a constraint on the layout of the receiving view or its subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addConstraint(_:)

func (v_ View) AddConstraint(constraint ILayoutConstraint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addConstraint:"), constraint)
}



// Adds multiple constraints on the layout of the receiving view or its subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addConstraints(_:)

func (v_ View) AddConstraints(constraints []LayoutConstraint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addConstraints:"), constraints)
}



// Establishes the cursor to be used when the mouse pointer lies within a specified region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addCursorRect(_:cursor:)

func (v_ View) AddCursorRectCursor(rect coregraphics.CGRect, object ICursor) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addCursorRect:cursor:"), rect, object)
}



// Attaches a gesture recognizer to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addGestureRecognizer(_:)

func (v_ View) AddGestureRecognizer(gestureRecognizer IGestureRecognizer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addGestureRecognizer:"), gestureRecognizer)
}



// Adds the provided layout guide to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addLayoutGuide(_:)

func (v_ View) AddLayoutGuide(guide ILayoutGuide) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addLayoutGuide:"), guide)
}



// Adds a view to the view’s subviews so it’s displayed above its siblings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addSubview(_:)

func (v_ View) AddSubview(view IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addSubview:"), view)
}



// Inserts a view among the view’s subviews so it’s displayed immediately above or below another view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addSubview(_:positioned:relativeTo:)

func (v_ View) AddSubviewPositionedRelativeTo(view IView, place WindowOrderingMode, otherView IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addSubview:positioned:relativeTo:"), view, place, otherView)
}



// Creates a tooltip for a defined area in the view and returns a tag that identifies the tooltip rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addToolTip(_:owner:userData:)

func (v_ View) AddToolTipRectOwnerUserData(rect coregraphics.CGRect, owner objectivec.IObject, data unsafe.Pointer) ToolTipTag {
	rv := objc.Send[ToolTipTag](v_.ID, objc.Sel("addToolTipRect:owner:userData:"), rect, owner, data)
	return rv
}



// Adds a given tracking area to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addTrackingArea(_:)

func (v_ View) AddTrackingArea(trackingArea ITrackingArea) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addTrackingArea:"), trackingArea)
}



// Establishes an area for tracking mouse-entered and mouse-exited events within the view and returns a tag that identifies the tracking rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addTrackingRect(_:owner:userData:assumeInside:)

func (v_ View) AddTrackingRectOwnerUserDataAssumeInside(rect coregraphics.CGRect, owner objectivec.IObject, data unsafe.Pointer, flag bool) TrackingRectTag {
	rv := objc.Send[TrackingRectTag](v_.ID, objc.Sel("addTrackingRect:owner:userData:assumeInside:"), rect, owner, data, flag)
	return rv
}



// Overridden by subclasses to adjust page height during automatic pagination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/adjustPageHeightNew(_:top:bottom:limit:)

func (v_ View) AdjustPageHeightNewTopBottomLimit(newBottom coregraphics.float64, oldTop float64, oldBottom float64, bottomLimit float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("adjustPageHeightNew:top:bottom:limit:"), newBottom, oldTop, oldBottom, bottomLimit)
}



// Overridden by subclasses to adjust page width during automatic pagination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/adjustPageWidthNew(_:left:right:limit:)

func (v_ View) AdjustPageWidthNewLeftRightLimit(newRight coregraphics.float64, oldLeft float64, oldRight float64, rightLimit float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("adjustPageWidthNew:left:right:limit:"), newRight, oldLeft, oldRight, rightLimit)
}



// Overridden by subclasses to modify a given rectangle, returning the altered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/adjustScroll(_:)

func (v_ View) AdjustScroll(newVisible coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("adjustScroll:"), newVisible)
	return rv
}



// Returns the view’s alignment rectangle for a given frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/alignmentRect(forFrame:)

func (v_ View) AlignmentRectForFrame(frame coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("alignmentRectForFrame:"), frame)
	return rv
}



// Causes the view to maintain a private graphics state object, which encapsulates all parameters of the graphics environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/allocateGState()

func (v_ View) AllocateGState() {
	objc.Send[objc.ID](v_.ID, objc.Sel("allocateGState"))
}



// Returns the closest ancestor shared by the view and another specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/ancestorShared(with:)

func (v_ View) AncestorSharedWithView(view IView) View {
	rv := objc.Send[View](v_.ID, objc.Sel("ancestorSharedWithView:"), view)
	return rv
}



// Scrolls the view’s closest ancestor object proportionally to the distance of an event that occurs outside of it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoscroll(with:)

func (v_ View) Autoscroll(event IEvent) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("autoscroll:"), event)
	return rv
}



// Returns a backing store pixel-aligned rectangle in local view coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/backingAlignedRect(_:options:)

func (v_ View) BackingAlignedRectOptions(rect coregraphics.CGRect, options unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
	return rv
}



// Invoked at the beginning of the printing session, this method sets up the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/beginDocument()

func (v_ View) BeginDocument() {
	objc.Send[objc.ID](v_.ID, objc.Sel("beginDocument"))
}



// Initiates a dragging session with a group of dragging items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/beginDraggingSession(with:event:source:)

func (v_ View) BeginDraggingSessionWithItemsEventSource(items []DraggingItem, event IEvent, source objectivec.IObject) DraggingSession {
	rv := objc.Send[DraggingSession](v_.ID, objc.Sel("beginDraggingSessionWithItems:event:source:"), items, event, source)
	return rv
}



// Called at the beginning of each page, this method sets up the coordinate system so that a region inside the view’s bounds is translated to a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/beginPage(in:atPlacement:)

func (v_ View) BeginPageInRectAtPlacement(rect coregraphics.CGRect, location coregraphics.CGPoint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("beginPageInRect:atPlacement:"), rect, location)
}



// Returns a bitmap-representation object suitable for caching the specified portion of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/bitmapImageRepForCachingDisplay(in:)

func (v_ View) BitmapImageRepForCachingDisplayInRect(rect coregraphics.CGRect) BitmapImageRep {
	rv := objc.Send[BitmapImageRep](v_.ID, objc.Sel("bitmapImageRepForCachingDisplayInRect:"), rect)
	return rv
}



// Draws the specified area of the view, and its descendants, into a provided bitmap-representation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/cacheDisplay(in:to:)

func (v_ View) CacheDisplayInRectToBitmapImageRep(rect coregraphics.CGRect, bitmapImageRep IBitmapImageRep) {
	objc.Send[objc.ID](v_.ID, objc.Sel("cacheDisplayInRect:toBitmapImageRep:"), rect, bitmapImageRep)
}



// Converts the corners of a specified rectangle to lie on the center of device pixels, which is useful in compensating for rendering overscanning when the coordinate system has been scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/centerScanRect(_:)

func (v_ View) CenterScanRect(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("centerScanRect:"), rect)
	return rv
}



// Returns the constraints impacting the layout of the view for a given orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/constraintsAffectingLayout(for:)

func (v_ View) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := objc.Send[[]LayoutConstraint](v_.ID, objc.Sel("constraintsAffectingLayoutForOrientation:"), orientation)
	return rv
}



// Returns the priority with which a view resists being made smaller than its intrinsic size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/contentCompressionResistancePriority(for:)

func (v_ View) ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.Send[LayoutPriority](v_.ID, objc.Sel("contentCompressionResistancePriorityForOrientation:"), orientation)
	return rv
}



// Returns the priority with which a view resists being made larger than its intrinsic size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/contentHuggingPriority(for:)

func (v_ View) ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority {
	rv := objc.Send[LayoutPriority](v_.ID, objc.Sel("contentHuggingPriorityForOrientation:"), orientation)
	return rv
}



// Converts a point from the coordinate system of a given view to that of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-1dq9l

func (v_ View) ConvertPointFromView(point coregraphics.CGPoint, view IView) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("convertPoint:fromView:"), point, view)
	return rv
}



// Converts a size from another view’s coordinate system to that of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-40x0w

func (v_ View) ConvertSizeFromView(size coregraphics.CGSize, view IView) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("convertSize:fromView:"), size, view)
	return rv
}



// Converts a rectangle from the coordinate system of another view to that of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-7fbb6

func (v_ View) ConvertRectFromView(rect coregraphics.CGRect, view IView) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("convertRect:fromView:"), rect, view)
	return rv
}



// Converts a rectangle from the view’s coordinate system to that of another view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-3cqqt

func (v_ View) ConvertRectToView(rect coregraphics.CGRect, view IView) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("convertRect:toView:"), rect, view)
	return rv
}



// Converts a size from the view’s coordinate system to that of another view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-5nptx

func (v_ View) ConvertSizeToView(size coregraphics.CGSize, view IView) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("convertSize:toView:"), size, view)
	return rv
}



// Converts a point from the view’s coordinate system to that of a given view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-6u9ir

func (v_ View) ConvertPointToView(point coregraphics.CGPoint, view IView) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("convertPoint:toView:"), point, view)
	return rv
}



// Converts a point from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-229ps

func (v_ View) ConvertPointFromBacking(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("convertPointFromBacking:"), point)
	return rv
}



// Converts a rectangle from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-2njpa

func (v_ View) ConvertRectFromBacking(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}



// Converts a size from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-4agf9

func (v_ View) ConvertSizeFromBacking(size coregraphics.CGSize) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("convertSizeFromBacking:"), size)
	return rv
}



// Convert the point from the layer’s interior coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-3nsbu

func (v_ View) ConvertPointFromLayer(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("convertPointFromLayer:"), point)
	return rv
}



// Convert the size from the layer’s interior coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-3usqp

func (v_ View) ConvertSizeFromLayer(size coregraphics.CGSize) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("convertSizeFromLayer:"), size)
	return rv
}



// Convert the rectangle from the layer’s interior coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-8s5bi

func (v_ View) ConvertRectFromLayer(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("convertRectFromLayer:"), rect)
	return rv
}



// Converts the point from the base coordinate system to the view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertPointFromBase:

func (v_ View) ConvertPointFromBase(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("convertPointFromBase:"), point)
	return rv
}



// Converts the point from the view’s coordinate system to the base coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertPointToBase:

func (v_ View) ConvertPointToBase(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("convertPointToBase:"), point)
	return rv
}



// Converts the rectangle from the base coordinate system to the view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertRectFromBase:

func (v_ View) ConvertRectFromBase(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("convertRectFromBase:"), rect)
	return rv
}



// Converts the rectangle from the view’s coordinate system to the base coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertRectToBase:

func (v_ View) ConvertRectToBase(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("convertRectToBase:"), rect)
	return rv
}



// Converts the size from the base coordinate system to the view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertSizeFromBase:

func (v_ View) ConvertSizeFromBase(size coregraphics.CGSize) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("convertSizeFromBase:"), size)
	return rv
}



// Converts the size from the view’s coordinate system to the base coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertSizeToBase:

func (v_ View) ConvertSizeToBase(size coregraphics.CGSize) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("convertSizeToBase:"), size)
	return rv
}



// Converts a point from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-2xx45

func (v_ View) ConvertPointToBacking(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("convertPointToBacking:"), point)
	return rv
}



// Converts a rectangle from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-3zors

func (v_ View) ConvertRectToBacking(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("convertRectToBacking:"), rect)
	return rv
}



// Converts a size from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-4ra9y

func (v_ View) ConvertSizeToBacking(size coregraphics.CGSize) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("convertSizeToBacking:"), size)
	return rv
}



// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-160pw

func (v_ View) ConvertRectToLayer(rect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("convertRectToLayer:"), rect)
	return rv
}



// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-2vozx

func (v_ View) ConvertSizeToLayer(size coregraphics.CGSize) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("convertSizeToLayer:"), size)
	return rv
}



// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-44u7d

func (v_ View) ConvertPointToLayer(point coregraphics.CGPoint) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("convertPointToLayer:"), point)
	return rv
}



// Returns EPS data that draws the region of the view within a specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dataWithEPS(inside:)

func (v_ View) DataWithEPSInsideRect(rect coregraphics.CGRect) foundation.Data {
	rv := objc.Send[foundation.Data](v_.ID, objc.Sel("dataWithEPSInsideRect:"), rect)
	return rv
}



// Returns PDF data that draws the region of the view within a specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dataWithPDF(inside:)

func (v_ View) DataWithPDFInsideRect(rect coregraphics.CGRect) foundation.Data {
	rv := objc.Send[foundation.Data](v_.ID, objc.Sel("dataWithPDFInsideRect:"), rect)
	return rv
}



// Overridden by subclasses to perform additional actions when subviews are added to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/didAddSubview(_:)

func (v_ View) DidAddSubview(subview IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("didAddSubview:"), subview)
}



// Called after a contextual menu that was displayed from the receiving view has been closed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/didCloseMenu(_:with:)

func (v_ View) DidCloseMenuWithEvent(menu IMenu, event IEvent) {
	objc.Send[objc.ID](v_.ID, objc.Sel("didCloseMenu:withEvent:"), menu, event)
}



// Invalidates all cursor rectangles set up using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/discardCursorRects()

func (v_ View) DiscardCursorRects() {
	objc.Send[objc.ID](v_.ID, objc.Sel("discardCursorRects"))
}



// Displays the view and all its subviews if possible, invoking each of the methods , , and as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/display()

func (v_ View) Display() {
	objc.Send[objc.ID](v_.ID, objc.Sel("display"))
}



// Acts as , but confining drawing to a rectangular region of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/display(_:)

func (v_ View) DisplayRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayRect:"), rect)
}



// Displays the view and all its subviews if any part of the view has been marked as needing display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeeded()

func (v_ View) DisplayIfNeeded() {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeeded"))
}



// Acts as , confining drawing to a specified region of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeeded(_:)

func (v_ View) DisplayIfNeededInRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeededInRect:"), rect)
}



// Acts as , except that this method doesn’t back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity()

func (v_ View) DisplayIfNeededIgnoringOpacity() {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeededIgnoringOpacity"))
}



// Acts as , but confining drawing to and not backing up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity(_:)

func (v_ View) DisplayIfNeededInRectIgnoringOpacity(rect coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeededInRectIgnoringOpacity:"), rect)
}



// Displays the view but confines drawing to a specified region and does not back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIgnoringOpacity(_:)

func (v_ View) DisplayRectIgnoringOpacity(rect coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayRectIgnoringOpacity:"), rect)
}



// Causes the view and its descendants to be redrawn to the specified graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIgnoringOpacity(_:in:)

func (v_ View) DisplayRectIgnoringOpacityInContext(rect coregraphics.CGRect, context IGraphicsContext) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayRectIgnoringOpacity:inContext:"), rect, context)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayLink(target:selector:)

func (v_ View) DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) quartzcore.DisplayLink {
	rv := objc.Send[quartzcore.DisplayLink](v_.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return rv
}



// Initiates a dragging operation from the view, allowing the user to drag a file icon to any application that has window or view objects that accept files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dragFile(_:from:slideBack:event:)

func (v_ View) DragFileFromRectSlideBackEvent(filename string, rect coregraphics.CGRect, flag bool, event IEvent) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("dragFile:fromRect:slideBack:event:"), objc.String(filename), rect, flag, event)
	return rv
}



// Initiates a dragging operation from the view, allowing the user to drag arbitrary data with a specified icon into any application that has window or view objects that accept dragged data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dragImage:at:offset:event:pasteboard:source:slideBack:

func (v_ View) DragImageAtOffsetEventPasteboardSourceSlideBack(image IImage, viewLocation coregraphics.CGPoint, initialOffset coregraphics.CGSize, event IEvent, pboard IPasteboard, sourceObj objectivec.IObject, slideFlag bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("dragImage:at:offset:event:pasteboard:source:slideBack:"), image, viewLocation, initialOffset, event, pboard, sourceObj, slideFlag)
}



// Initiates a dragging operation from the view, allowing the user to drag one or more promised files (or directories) into any application that has window or view objects that accept promised file data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dragPromisedFiles(ofTypes:from:source:slideBack:event:)

func (v_ View) DragPromisedFilesOfTypesFromRectSourceSlideBackEvent(typeArray []string, rect coregraphics.CGRect, sourceObject objectivec.IObject, flag bool, event IEvent) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("dragPromisedFilesOfTypes:fromRect:source:slideBack:event:"), typeArray, rect, sourceObject, flag, event)
	return rv
}



// Overridden by subclasses to draw the view’s image within the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/draw(_:)

func (v_ View) DrawRect(dirtyRect coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawRect:"), dirtyRect)
}



// Draws the focus ring mask for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/drawFocusRingMask()

func (v_ View) DrawFocusRingMask() {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawFocusRingMask"))
}



// Allows applications that use the AppKit pagination facility to draw additional marks on each logical page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/drawPageBorder(with:)

func (v_ View) DrawPageBorderWithSize(borderSize coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawPageBorderWithSize:"), borderSize)
}



// Allows applications that use the AppKit pagination facility to draw additional marks on each printed sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/drawSheetBorder(with:)

func (v_ View) DrawSheetBorderWithSize(borderSize coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawSheetBorderWithSize:"), borderSize)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/edgeInsetsForLayoutRegion:

func (v_ View) EdgeInsetsForLayoutRegion(layoutRegion IViewLayoutRegion) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("edgeInsetsForLayoutRegion:"), layoutRegion)
	return rv
}



// This method is invoked at the end of the printing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/endDocument()

func (v_ View) EndDocument() {
	objc.Send[objc.ID](v_.ID, objc.Sel("endDocument"))
}



// Writes the end of a conforming page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/endPage()

func (v_ View) EndPage() {
	objc.Send[objc.ID](v_.ID, objc.Sel("endPage"))
}



// Sets the view to full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/enterFullScreenMode(_:withOptions:)

func (v_ View) EnterFullScreenModeWithOptions(screen IScreen, options unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("enterFullScreenMode:withOptions:"), screen, options)
	return rv
}



// Randomly changes the frame of a view with an ambiguous layout between the different valid values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/exerciseAmbiguityInLayout()

func (v_ View) ExerciseAmbiguityInLayout() {
	objc.Send[objc.ID](v_.ID, objc.Sel("exerciseAmbiguityInLayout"))
}



// Instructs the view to exit full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/exitFullScreenMode(options:)

func (v_ View) ExitFullScreenModeWithOptions(options unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("exitFullScreenModeWithOptions:"), options)
}



// Returns the view’s frame for a given alignment rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frame(forAlignmentRect:)

func (v_ View) FrameForAlignmentRect(alignmentRect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("frameForAlignmentRect:"), alignmentRect)
	return rv
}



// Returns the identifier for the view’s graphics state object, or 0 if the view doesn’t have a graphics state object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/gState()

func (v_ View) GState() int {
	rv := objc.Send[int](v_.ID, objc.Sel("gState"))
	return rv
}



// Returns by indirection a list of nonoverlapping rectangles that define the area the view is being asked to draw in .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/getRectsBeingDrawn(_:count:)

func (v_ View) GetRectsBeingDrawnCount(rects coregraphics.CGRect, count unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("getRectsBeingDrawn:count:"), rects, count)
}



// Returns a list of rectangles indicating the newly exposed areas of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/getRectsExposedDuringLiveResize(_:count:)

func (v_ View) GetRectsExposedDuringLiveResizeCount(exposedRects unsafe.Pointer, count unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("getRectsExposedDuringLiveResize:count:"), exposedRects, count)
}



// Returns the farthest descendant of the view in the view hierarchy (including itself) that contains a specified point, or if that point lies completely outside the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/hitTest(_:)

func (v_ View) HitTest(point coregraphics.CGPoint) View {
	rv := objc.Send[View](v_.ID, objc.Sel("hitTest:"), point)
	return rv
}



// Invalidates the view’s intrinsic content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/invalidateIntrinsicContentSize()

func (v_ View) InvalidateIntrinsicContentSize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("invalidateIntrinsicContentSize"))
}



// Returns a Boolean value that indicates whether the view is a subview of the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isDescendant(of:)

func (v_ View) IsDescendantOf(view IView) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isDescendantOf:"), view)
	return rv
}



// Returns whether a region of the view contains a specified point, accounting for whether the view is flipped or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isMousePoint(_:in:)

func (v_ View) MouseInRect(point coregraphics.CGPoint, rect coregraphics.CGRect) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("mouse:inRect:"), point, rect)
	return rv
}



// Returns if the view handles page boundaries, otherwise.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/knowsPageRange(_:)

func (v_ View) KnowsPageRange(range_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("knowsPageRange:"), range_)
	return rv
}



// Perform layout in concert with the constraint-based layout system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layout()

func (v_ View) Layout() {
	objc.Send[objc.ID](v_.ID, objc.Sel("layout"))
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layoutGuideForLayoutRegion:

func (v_ View) LayoutGuideForLayoutRegion(layoutRegion IViewLayoutRegion) LayoutGuide {
	rv := objc.Send[LayoutGuide](v_.ID, objc.Sel("layoutGuideForLayoutRegion:"), layoutRegion)
	return rv
}



// Updates the layout of the receiving view and its subviews based on the current views and constraints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layoutSubtreeIfNeeded()

func (v_ View) LayoutSubtreeIfNeeded() {
	objc.Send[objc.ID](v_.ID, objc.Sel("layoutSubtreeIfNeeded"))
}



// Invoked by to determine the location of the region of the view being printed on the physical page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/locationOfPrintRect(_:)

func (v_ View) LocationOfPrintRect(rect coregraphics.CGRect) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("locationOfPrintRect:"), rect)
	return rv
}



// Locks the focus on the view, so subsequent commands take effect in the view’s window and coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/lockFocus()

func (v_ View) LockFocus() {
	objc.Send[objc.ID](v_.ID, objc.Sel("lockFocus"))
}



// Locks the focus to the view atomically if the method returns and returns the value of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/lockFocusIfCanDraw()

func (v_ View) LockFocusIfCanDraw() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("lockFocusIfCanDraw"))
	return rv
}



// Locks the focus to the view atomically if drawing can occur in the specified graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/lockFocusIfCanDraw(in:)

func (v_ View) LockFocusIfCanDrawInContext(context IGraphicsContext) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("lockFocusIfCanDrawInContext:"), context)
	return rv
}



// Creates the view’s backing layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/makeBackingLayer()

func (v_ View) MakeBackingLayer() quartzcore.Layer {
	rv := objc.Send[quartzcore.Layer](v_.ID, objc.Sel("makeBackingLayer"))
	return rv
}



// Overridden by subclasses to return a context-sensitive pop-up menu for a given mouse-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/menu(for:)

func (v_ View) MenuForEvent(event IEvent) Menu {
	rv := objc.Send[Menu](v_.ID, objc.Sel("menuForEvent:"), event)
	return rv
}



// Returns a Boolean value indicating whether the specified rectangle intersects any part of the area that the view is being asked to draw.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsToDraw(_:)

func (v_ View) NeedsToDrawRect(rect coregraphics.CGRect) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsToDrawRect:"), rect)
	return rv
}



// Invoked to notify the view that the focus ring mask requires updating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/noteFocusRingMaskChanged()

func (v_ View) NoteFocusRingMaskChanged() {
	objc.Send[objc.ID](v_.ID, objc.Sel("noteFocusRingMaskChanged"))
}



// Implemented by subclasses to respond to key equivalents (also known as keyboard shortcuts).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/performKeyEquivalent(with:)

func (v_ View) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("performKeyEquivalent:"), event)
	return rv
}



// Implemented by subclasses to respond to mnemonics.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/performMnemonic:

func (v_ View) PerformMnemonic(string_ string) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("performMnemonic:"), objc.String(string_))
	return rv
}



// Prepares the overdraw region for drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/prepareContent(in:)

func (v_ View) PrepareContentInRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("prepareContentInRect:"), rect)
}



// Restores the view to an initial state so that it can be reused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/prepareForReuse()

func (v_ View) PrepareForReuse() {
	objc.Send[objc.ID](v_.ID, objc.Sel("prepareForReuse"))
}



// This action method opens the Print panel, and if the user chooses an option other than canceling, prints the view and all its subviews to the device specified in the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/printView(_:)

func (v_ View) Print(sender objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("print:"), sender)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectForLayoutRegion:

func (v_ View) RectForLayoutRegion(layoutRegion IViewLayoutRegion) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("rectForLayoutRegion:"), layoutRegion)
	return rv
}



// Implemented by subclasses to determine the portion of the view to be printed for the specified page number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectForPage(_:)

func (v_ View) RectForPage(page int) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("rectForPage:"), page)
	return rv
}



// Returns the appropriate rectangle to use when magnifying around the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectForSmartMagnification(at:in:)

func (v_ View) RectForSmartMagnificationAtPointInRect(location coregraphics.CGPoint, visibleRect coregraphics.CGRect) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("rectForSmartMagnificationAtPoint:inRect:"), location, visibleRect)
	return rv
}



// Notifies a clip view’s superview that either the clip view’s bounds rectangle or the document view’s frame rectangle has changed, and that any indicators of the scroll position need to be adjusted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/reflectScrolledClipView(_:)

func (v_ View) ReflectScrolledClipView(clipView IClipView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("reflectScrolledClipView:"), clipView)
}



// Registers the pasteboard types that the view will accept as the destination of an image-dragging session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/registerForDraggedTypes(_:)

func (v_ View) RegisterForDraggedTypes(newTypes []string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("registerForDraggedTypes:"), newTypes)
}



// Frees the view’s graphics state object, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/releaseGState()

func (v_ View) ReleaseGState() {
	objc.Send[objc.ID](v_.ID, objc.Sel("releaseGState"))
}



// Removes all tooltips assigned to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeAllToolTips()

func (v_ View) RemoveAllToolTips() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeAllToolTips"))
}



// Removes the specified constraint from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeConstraint(_:)

func (v_ View) RemoveConstraint(constraint ILayoutConstraint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeConstraint:"), constraint)
}



// Removes the specified constraints from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeConstraints(_:)

func (v_ View) RemoveConstraints(constraints []LayoutConstraint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeConstraints:"), constraints)
}



// Completely removes a cursor rectangle from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeCursorRect(_:cursor:)

func (v_ View) RemoveCursorRectCursor(rect coregraphics.CGRect, object ICursor) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeCursorRect:cursor:"), rect, object)
}



// Unlinks the view from its superview and its window, removes it from the responder chain, and invalidates its cursor rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeFromSuperview()

func (v_ View) RemoveFromSuperview() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeFromSuperview"))
}



// Unlinks the view from its superview and its window and removes it from the responder chain, but does not invalidate its cursor rectangles to cause redrawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeFromSuperviewWithoutNeedingDisplay()

func (v_ View) RemoveFromSuperviewWithoutNeedingDisplay() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeFromSuperviewWithoutNeedingDisplay"))
}



// Detaches a gesture recognizer from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeGestureRecognizer(_:)

func (v_ View) RemoveGestureRecognizer(gestureRecognizer IGestureRecognizer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeGestureRecognizer:"), gestureRecognizer)
}



// Removes the provided layout guide from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeLayoutGuide(_:)

func (v_ View) RemoveLayoutGuide(guide ILayoutGuide) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeLayoutGuide:"), guide)
}



// Removes the tooltip identified by specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeToolTip(_:)

func (v_ View) RemoveToolTip(tag IToolTipTag) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeToolTip:"), tag)
}



// Removes a given tracking area from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeTrackingArea(_:)

func (v_ View) RemoveTrackingArea(trackingArea ITrackingArea) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeTrackingArea:"), trackingArea)
}



// Removes the tracking rectangle identified by a tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeTrackingRect(_:)

func (v_ View) RemoveTrackingRect(tag ITrackingRectTag) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeTrackingRect:"), tag)
}



// Invalidates the view’s graphics state object, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/renewGState()

func (v_ View) RenewGState() {
	objc.Send[objc.ID](v_.ID, objc.Sel("renewGState"))
}



// Replaces one of the view’s subviews with another view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/replaceSubview(_:with:)

func (v_ View) ReplaceSubviewWith(oldView IView, newView IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("replaceSubview:with:"), oldView, newView)
}



// Overridden by subclasses to define their default cursor rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/resetCursorRects()

func (v_ View) ResetCursorRects() {
	objc.Send[objc.ID](v_.ID, objc.Sel("resetCursorRects"))
}



// Informs the view that the bounds size of its superview has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/resize(withOldSuperviewSize:)

func (v_ View) ResizeWithOldSuperviewSize(oldSize coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("resizeWithOldSuperviewSize:"), oldSize)
}



// Informs the view’s subviews that the view’s bounds rectangle size has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/resizeSubviews(withOldSize:)

func (v_ View) ResizeSubviewsWithOldSize(oldSize coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("resizeSubviewsWithOldSize:"), oldSize)
}



// Rotates the view’s bounds rectangle by a specified degree value around the origin of the coordinate system, (0.0, 0.0).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rotate(byDegrees:)

func (v_ View) RotateByAngle(angle float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rotateByAngle:"), angle)
}



// Informs the client that allowed the user to add .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didAdd:)

func (v_ View) RulerViewDidAddMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:didAddMarker:"), ruler, marker)
}



// Informs the client that allowed the user to move .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didMove:)

func (v_ View) RulerViewDidMoveMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:didMoveMarker:"), ruler, marker)
}



// Informs the client that allowed the user to remove .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didRemove:)

func (v_ View) RulerViewDidRemoveMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:didRemoveMarker:"), ruler, marker)
}



// Informs the client that the user has pressed the mouse button while the cursor is in the ruler area of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:handleMouseDownWith:)

func (v_ View) RulerViewHandleMouseDown(ruler IRulerView, event IEvent) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:handleMouseDown:"), ruler, event)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:locationFor:)

func (v_ View) RulerViewLocationForPoint(ruler IRulerView, point coregraphics.CGPoint) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("rulerView:locationForPoint:"), ruler, point)
	return rv
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:pointForLocation:)

func (v_ View) RulerViewPointForLocation(ruler IRulerView, point float64) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("rulerView:pointForLocation:"), ruler, point)
	return rv
}



// Requests permission for to add , an NSRulerMarker being dragged onto the ruler by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldAdd:)

func (v_ View) RulerViewShouldAddMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rulerView:shouldAddMarker:"), ruler, marker)
	return rv
}



// Requests permission for to move .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldMove:)

func (v_ View) RulerViewShouldMoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rulerView:shouldMoveMarker:"), ruler, marker)
	return rv
}



// Requests permission for to remove .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldRemove:)

func (v_ View) RulerViewShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rulerView:shouldRemoveMarker:"), ruler, marker)
	return rv
}



// Informs the client that will add the new NSRulerMarker, .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willAdd:atLocation:)

func (v_ View) RulerViewWillAddMarkerAtLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("rulerView:willAddMarker:atLocation:"), ruler, marker, location)
	return rv
}



// Informs the client that will move , an NSRulerMarker already on the ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willMove:toLocation:)

func (v_ View) RulerViewWillMoveMarkerToLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("rulerView:willMoveMarker:toLocation:"), ruler, marker, location)
	return rv
}



// Informs the client view that is about to be appropriated by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willSetClientView:)

func (v_ View) RulerViewWillSetClientView(ruler IRulerView, newClient IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:willSetClientView:"), ruler, newClient)
}



// Scales the view’s coordinate system so that the unit square scales to the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scaleUnitSquare(to:)

func (v_ View) ScaleUnitSquareToSize(newUnitSize coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scaleUnitSquareToSize:"), newUnitSize)
}



// Scrolls the view’s closest ancestor object so a point in the view lies at the origin of the clip view’s bounds rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:)

func (v_ View) ScrollPoint(point coregraphics.CGPoint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scrollPoint:"), point)
}



// Copies the visible portion of the view’s rendered image within a region and lays that portion down again at a specified offset .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:by:)

func (v_ View) ScrollRectBy(rect coregraphics.CGRect, delta coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scrollRect:by:"), rect, delta)
}



// Notifies the superview of a clip view that the clip view needs to reset the origin of its bounds rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:to:)

func (v_ View) ScrollClipViewToPoint(clipView IClipView, point coregraphics.CGPoint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scrollClipView:toPoint:"), clipView, point)
}



// Scrolls the view’s closest ancestor object the minimum distance needed so a specified region of the view becomes visible in the clip view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scrollToVisible(_:)

func (v_ View) ScrollRectToVisible(rect coregraphics.CGRect) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("scrollRectToVisible:"), rect)
	return rv
}



// Sets the origin of the view’s bounds rectangle to a specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setBoundsOrigin(_:)

func (v_ View) SetBoundsOrigin(newOrigin coregraphics.CGPoint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBoundsOrigin:"), newOrigin)
}



// Sets the size of the view’s bounds rectangle to specified dimensions, inversely scaling its coordinate system relative to its frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setBoundsSize(_:)

func (v_ View) SetBoundsSize(newSize coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBoundsSize:"), newSize)
}



// Sets the priority with which a view resists being made smaller than its intrinsic size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setContentCompressionResistancePriority(_:for:)

func (v_ View) SetContentCompressionResistancePriorityForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setContentCompressionResistancePriority:forOrientation:"), priority, orientation)
}



// Sets the priority with which a view resists being made larger than its intrinsic size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setContentHuggingPriority(_:for:)

func (v_ View) SetContentHuggingPriorityForOrientation(priority LayoutPriority, orientation LayoutConstraintOrientation) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setContentHuggingPriority:forOrientation:"), priority, orientation)
}



// Sets the origin of the view’s frame rectangle to the specified point, effectively repositioning it within its superview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setFrameOrigin(_:)

func (v_ View) SetFrameOrigin(newOrigin coregraphics.CGPoint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameOrigin:"), newOrigin)
}



// Sets the size of the view’s frame rectangle to the specified dimensions, resizing it within its superview without affecting its coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setFrameSize(_:)

func (v_ View) SetFrameSize(newSize coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameSize:"), newSize)
}



// Invalidates the area around the focus ring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setKeyboardFocusRingNeedsDisplay(_:)

func (v_ View) SetKeyboardFocusRingNeedsDisplayInRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setKeyboardFocusRingNeedsDisplayInRect:"), rect)
}



// Marks the region of the view within the specified rectangle as needing display, increasing the view’s existing invalid region to include it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setNeedsDisplay(_:)

func (v_ View) SetNeedsDisplayInRect(invalidRect coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNeedsDisplayInRect:"), invalidRect)
}



// Overridden by subclasses to (re)initialize the view’s graphics state object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setUpGState()

func (v_ View) SetUpGState() {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUpGState"))
}



// Allows the user to drag objects from the view without activating the app or moving the window of the view forward, possibly obscuring the destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/shouldDelayWindowOrdering(for:)

func (v_ View) ShouldDelayWindowOrderingForEvent(event IEvent) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("shouldDelayWindowOrderingForEvent:"), event)
	return rv
}



// Returns a Boolean value indicating whether the view is being drawn to an environment that supports color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/shouldDrawColor()

func (v_ View) ShouldDrawColor() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("shouldDrawColor"))
	return rv
}



// Shows a window displaying the definition of the attributed string at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/showDefinition(for:at:)

func (v_ View) ShowDefinitionForAttributedStringAtPoint(attrString foundation.IAttributedString, textBaselineOrigin coregraphics.CGPoint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("showDefinitionForAttributedString:atPoint:"), attrString, textBaselineOrigin)
}



// Shows a window displaying the definition of the specified range of the attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/showDefinition(for:range:options:baselineOriginProvider:)

func (v_ View) ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString foundation.IAttributedString, targetRange foundation.Range, options unsafe.Pointer, originProvider unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("showDefinitionForAttributedString:range:options:baselineOriginProvider:"), attrString, targetRange, options, originProvider)
}



// Orders the view’s immediate subviews using the specified comparator function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/sortSubviews(_:context:)

func (v_ View) SortSubviewsUsingFunctionContext(compare unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("sortSubviewsUsingFunction:context:"), compare, context)
}



// Translates the view’s coordinate system so that its origin moves to a new location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/translateOrigin(to:)

func (v_ View) TranslateOriginToPoint(translation coregraphics.CGPoint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("translateOriginToPoint:"), translation)
}



// Translates the display rectangles by the specified delta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/translateRectsNeedingDisplay(in:by:)

func (v_ View) TranslateRectsNeedingDisplayInRectBy(clipRect coregraphics.CGRect, delta coregraphics.CGSize) {
	objc.Send[objc.ID](v_.ID, objc.Sel("translateRectsNeedingDisplayInRect:by:"), clipRect, delta)
}



// Unlocks focus from the current view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/unlockFocus()

func (v_ View) UnlockFocus() {
	objc.Send[objc.ID](v_.ID, objc.Sel("unlockFocus"))
}



// Unregisters the view as a possible destination in a dragging session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/unregisterDraggedTypes()

func (v_ View) UnregisterDraggedTypes() {
	objc.Send[objc.ID](v_.ID, objc.Sel("unregisterDraggedTypes"))
}



// Update constraints for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateConstraints()

func (v_ View) UpdateConstraints() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateConstraints"))
}



// Updates the constraints for the receiving view and its subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateConstraintsForSubtreeIfNeeded()

func (v_ View) UpdateConstraintsForSubtreeIfNeeded() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateConstraintsForSubtreeIfNeeded"))
}



// Updates the view’s content by modifying its underlying layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateLayer()

func (v_ View) UpdateLayer() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateLayer"))
}



// Invoked automatically when the view’s geometry changes such that its tracking areas need to be recalculated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateTrackingAreas()

func (v_ View) UpdateTrackingAreas() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateTrackingAreas"))
}



// Responds when the view’s backing store properties change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidChangeBackingProperties()

func (v_ View) ViewDidChangeBackingProperties() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidChangeBackingProperties"))
}



// Informs the view that its effective appearance changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidChangeEffectiveAppearance()

func (v_ View) ViewDidChangeEffectiveAppearance() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidChangeEffectiveAppearance"))
}



// Informs the view of the end of a live resize—the user has finished resizing the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidEndLiveResize()

func (v_ View) ViewDidEndLiveResize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidEndLiveResize"))
}



// Invoked when the view is hidden, either directly, or in response to an ancestor being hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidHide()

func (v_ View) ViewDidHide() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidHide"))
}



// Informs the view that its superview has changed (possibly to ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidMoveToSuperview()

func (v_ View) ViewDidMoveToSuperview() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidMoveToSuperview"))
}



// Informs the view that it has been added to a new view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidMoveToWindow()

func (v_ View) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidMoveToWindow"))
}



// Invoked when the view is unhidden, either directly, or in response to an ancestor being unhidden
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidUnhide()

func (v_ View) ViewDidUnhide() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidUnhide"))
}



// Informs the view that it’s required to draw content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillDraw()

func (v_ View) ViewWillDraw() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillDraw"))
}



// Informs the view that its superview is about to change to the specified superview (which may be ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillMove(toSuperview:)

func (v_ View) ViewWillMoveToSuperview(newSuperview IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillMoveToSuperview:"), newSuperview)
}



// Informs the view that it’s being added to the view hierarchy of the specified window object (which may be ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillMove(toWindow:)

func (v_ View) ViewWillMoveToWindow(newWindow IWindow) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillMoveToWindow:"), newWindow)
}



// Informs the view of the start of a live resize—the user has started resizing the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillStartLiveResize()

func (v_ View) ViewWillStartLiveResize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillStartLiveResize"))
}



// Returns the view’s nearest descendant (including itself) with a specific tag, or if no subview has that tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWithTag(_:)

func (v_ View) ViewWithTag(tag int) View {
	rv := objc.Send[View](v_.ID, objc.Sel("viewWithTag:"), tag)
	return rv
}



// Called just before a contextual menu for a view is opened on screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/willOpenMenu(_:with:)

func (v_ View) WillOpenMenuWithEvent(menu IMenu, event IEvent) {
	objc.Send[objc.ID](v_.ID, objc.Sel("willOpenMenu:withEvent:"), menu, event)
}



// Overridden by subclasses to perform additional actions before subviews are removed from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/willRemoveSubview(_:)

func (v_ View) WillRemoveSubview(subview IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("willRemoveSubview:"), subview)
}



// Writes EPS data that draws the region of the view within a specified rectangle onto a pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/writeEPS(inside:to:)

func (v_ View) WriteEPSInsideRectToPasteboard(rect coregraphics.CGRect, pasteboard IPasteboard) {
	objc.Send[objc.ID](v_.ID, objc.Sel("writeEPSInsideRect:toPasteboard:"), rect, pasteboard)
}



// Writes PDF data that draws the region of the view within a specified rectangle onto a pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/writePDF(inside:to:)

func (v_ View) WritePDFInsideRectToPasteboard(rect coregraphics.CGRect, pasteboard IPasteboard) {
	objc.Send[objc.ID](v_.ID, objc.Sel("writePDFInsideRect:toPasteboard:"), rect, pasteboard)
}


// A Boolean value indicating whether the view accepts touch events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/acceptsTouchEvents

func (v_ View) AcceptsTouchEvents() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("acceptsTouchEvents"))
	return rv
}


// A Boolean value indicating whether the view accepts touch events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/acceptsTouchEvents

func (v_ View) SetAcceptsTouchEvents(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAcceptsTouchEvents:"), value)
}


// Custom insets that you specify to modify your view’s safe area
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/additionalSafeAreaInsets

func (v_ View) AdditionalSafeAreaInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("additionalSafeAreaInsets"))
	return rv
}


// Custom insets that you specify to modify your view’s safe area
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/additionalSafeAreaInsets

func (v_ View) SetAdditionalSafeAreaInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAdditionalSafeAreaInsets:"), value)
}


// The insets (in points) from the view’s frame that define its content rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/alignmentRectInsets

func (v_ View) AlignmentRectInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("alignmentRectInsets"))
	return rv
}


// The types of touch interactions the view allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/allowedTouchTypes

func (v_ View) AllowedTouchTypes() TouchTypeMask {
	rv := objc.Send[TouchTypeMask](v_.ID, objc.Sel("allowedTouchTypes"))
	return rv
}


// The types of touch interactions the view allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/allowedTouchTypes

func (v_ View) SetAllowedTouchTypes(value TouchTypeMask) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAllowedTouchTypes:"), value)
}


// A Boolean value indicating whether the view ensures it is vibrant on top of other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/allowsVibrancy

func (v_ View) AllowsVibrancy() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("allowsVibrancy"))
	return rv
}


// The opacity of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/alphaValue

func (v_ View) AlphaValue() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("alphaValue"))
	return rv
}


// The opacity of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/alphaValue

func (v_ View) SetAlphaValue(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAlphaValue:"), value)
}


// A Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoresizesSubviews

func (v_ View) AutoresizesSubviews() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("autoresizesSubviews"))
	return rv
}


// A Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoresizesSubviews

func (v_ View) SetAutoresizesSubviews(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAutoresizesSubviews:"), value)
}


// The options that determine how the view is resized relative to its superview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoresizingMask-swift.property

func (v_ View) AutoresizingMask() AutoresizingMaskOptions {
	rv := objc.Send[AutoresizingMaskOptions](v_.ID, objc.Sel("autoresizingMask"))
	return rv
}


// The options that determine how the view is resized relative to its superview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoresizingMask-swift.property

func (v_ View) SetAutoresizingMask(value AutoresizingMaskOptions) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAutoresizingMask:"), value)
}


// An array of Core Image filters to apply to the view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/backgroundFilters

func (v_ View) BackgroundFilters() []coreimage.Filter {
	rv := objc.Send[[]coreimage.Filter](v_.ID, objc.Sel("backgroundFilters"))
	return rv
}


// An array of Core Image filters to apply to the view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/backgroundFilters

func (v_ View) SetBackgroundFilters(value []coreimage.IFilter) {
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
	objc.Send[objc.ID](v_.ID, objc.Sel("setBackgroundFilters:"), nsArray)
}


// The distance (in points) between the bottom of the view’s alignment rectangle and its baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/baselineOffsetFromBottom

func (v_ View) BaselineOffsetFromBottom() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("baselineOffsetFromBottom"))
	return rv
}


// A layout anchor representing the bottom edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/bottomAnchor

func (v_ View) BottomAnchor() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](v_.ID, objc.Sel("bottomAnchor"))
	return rv
}


// The view’s bounds rectangle, which expresses its location and size in its own coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/bounds

func (v_ View) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("bounds"))
	return rv
}


// The view’s bounds rectangle, which expresses its location and size in its own coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/bounds

func (v_ View) SetBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBounds:"), value)
}


// The angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/boundsRotation

func (v_ View) BoundsRotation() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("boundsRotation"))
	return rv
}


// The angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/boundsRotation

func (v_ View) SetBoundsRotation(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBoundsRotation:"), value)
}


// A Boolean value indicating whether the view can become key view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canBecomeKeyView

func (v_ View) CanBecomeKeyView() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canBecomeKeyView"))
	return rv
}


// A Boolean value indicating whether drawing commands will produce any results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canDraw

func (v_ View) CanDraw() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canDraw"))
	return rv
}


// A Boolean value indicating whether the view can draw its contents on a background thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canDrawConcurrently

func (v_ View) CanDrawConcurrently() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canDrawConcurrently"))
	return rv
}


// A Boolean value indicating whether the view can draw its contents on a background thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canDrawConcurrently

func (v_ View) SetCanDrawConcurrently(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCanDrawConcurrently:"), value)
}


// A Boolean value indicating whether the view incorporates content from its subviews into its own layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canDrawSubviewsIntoLayer

func (v_ View) CanDrawSubviewsIntoLayer() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canDrawSubviewsIntoLayer"))
	return rv
}


// A Boolean value indicating whether the view incorporates content from its subviews into its own layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canDrawSubviewsIntoLayer

func (v_ View) SetCanDrawSubviewsIntoLayer(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCanDrawSubviewsIntoLayer:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/candidateListTouchBarItem

func (v_ View) CandidateListTouchBarItem() NSCandidateListTouchBarItem {
	rv := objc.Send[NSCandidateListTouchBarItem](v_.ID, objc.Sel("candidateListTouchBarItem"))
	return rv
}


// A layout anchor representing the horizontal center of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/centerXAnchor

func (v_ View) CenterXAnchor() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](v_.ID, objc.Sel("centerXAnchor"))
	return rv
}


// A layout anchor representing the vertical center of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/centerYAnchor

func (v_ View) CenterYAnchor() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](v_.ID, objc.Sel("centerYAnchor"))
	return rv
}


// A Boolean value that indicates whether the view, and its subviews, confine their drawing areas to the bounds of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/clipsToBounds

func (v_ View) ClipsToBounds() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("clipsToBounds"))
	return rv
}


// A Boolean value that indicates whether the view, and its subviews, confine their drawing areas to the bounds of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/clipsToBounds

func (v_ View) SetClipsToBounds(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setClipsToBounds:"), value)
}


// The Core Image filter used to composite the view’s contents with its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/compositingFilter

func (v_ View) CompositingFilter() coreimage.Filter {
	rv := objc.Send[coreimage.Filter](v_.ID, objc.Sel("compositingFilter"))
	return rv
}


// The Core Image filter used to composite the view’s contents with its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/compositingFilter

func (v_ View) SetCompositingFilter(value coreimage.IFilter) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCompositingFilter:"), value)
}


// Returns the constraints held by the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/constraints

func (v_ View) Constraints() []LayoutConstraint {
	rv := objc.Send[[]LayoutConstraint](v_.ID, objc.Sel("constraints"))
	return rv
}


// An array of Core Image filters to apply to the contents of the view and its sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/contentFilters

func (v_ View) ContentFilters() []coreimage.Filter {
	rv := objc.Send[[]coreimage.Filter](v_.ID, objc.Sel("contentFilters"))
	return rv
}


// An array of Core Image filters to apply to the contents of the view and its sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/contentFilters

func (v_ View) SetContentFilters(value []coreimage.IFilter) {
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
	objc.Send[objc.ID](v_.ID, objc.Sel("setContentFilters:"), nsArray)
}


// Returns the default focus ring type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/defaultFocusRingType

func (v_ View) DefaultFocusRingType() FocusRingType {
	rv := objc.Send[FocusRingType](v_.ID, objc.Sel("defaultFocusRingType"))
	return rv
}


// Overridden by subclasses to return the default pop-up menu for instances of the receiving class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/defaultMenu

func (v_ View) DefaultMenu() NSMenu {
	rv := objc.Send[NSMenu](v_.ID, objc.Sel("defaultMenu"))
	return rv
}


// The menu item containing the view or any of its superviews in the view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/enclosingMenuItem

func (v_ View) EnclosingMenuItem() NSMenuItem {
	rv := objc.Send[NSMenuItem](v_.ID, objc.Sel("enclosingMenuItem"))
	return rv
}


// The nearest ancestor scroll view that contains the current view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/enclosingScrollView

func (v_ View) EnclosingScrollView() NSScrollView {
	rv := objc.Send[NSScrollView](v_.ID, objc.Sel("enclosingScrollView"))
	return rv
}


// A layout anchor representing the baseline for the topmost line of text in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/firstBaselineAnchor

func (v_ View) FirstBaselineAnchor() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](v_.ID, objc.Sel("firstBaselineAnchor"))
	return rv
}


// The distance (in points) between the top of the view’s alignment rectangle and its topmost baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/firstBaselineOffsetFromTop

func (v_ View) FirstBaselineOffsetFromTop() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("firstBaselineOffsetFromTop"))
	return rv
}


// The minimum size of the view that satisfies the constraints it holds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/fittingSize

func (v_ View) FittingSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("fittingSize"))
	return rv
}


// The focus ring mask bounds, specified in the view’s coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/focusRingMaskBounds

func (v_ View) FocusRingMaskBounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("focusRingMaskBounds"))
	return rv
}


// The type of focus ring drawn around the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/focusRingType

func (v_ View) FocusRingType() FocusRingType {
	rv := objc.Send[FocusRingType](v_.ID, objc.Sel("focusRingType"))
	return rv
}


// The type of focus ring drawn around the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/focusRingType

func (v_ View) SetFocusRingType(value FocusRingType) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFocusRingType:"), value)
}


// The currently focused view object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/focusView

func (v_ View) FocusView() NSView {
	rv := objc.Send[NSView](v_.ID, objc.Sel("focusView"))
	return rv
}


// The view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frame

func (v_ View) Frame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("frame"))
	return rv
}


// The view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frame

func (v_ View) SetFrame(value coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrame:"), value)
}


// The rotation angle of the view around the center of its layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frameCenterRotation

func (v_ View) FrameCenterRotation() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("frameCenterRotation"))
	return rv
}


// The rotation angle of the view around the center of its layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frameCenterRotation

func (v_ View) SetFrameCenterRotation(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameCenterRotation:"), value)
}


// The angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frameRotation

func (v_ View) FrameRotation() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("frameRotation"))
	return rv
}


// The angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frameRotation

func (v_ View) SetFrameRotation(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameRotation:"), value)
}


// The gesture recognize objects currently attached to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/gestureRecognizers

func (v_ View) GestureRecognizers() []GestureRecognizer {
	rv := objc.Send[[]GestureRecognizer](v_.ID, objc.Sel("gestureRecognizers"))
	return rv
}


// The gesture recognize objects currently attached to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/gestureRecognizers

func (v_ View) SetGestureRecognizers(value []GestureRecognizer) {
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
	objc.Send[objc.ID](v_.ID, objc.Sel("setGestureRecognizers:"), nsArray)
}


// A Boolean value indicating whether the constraints impacting the layout of the view incompletely specify the location of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/hasAmbiguousLayout

func (v_ View) HasAmbiguousLayout() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("hasAmbiguousLayout"))
	return rv
}


// The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as lines of text from being divided across pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/heightAdjustLimit

func (v_ View) HeightAdjustLimit() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("heightAdjustLimit"))
	return rv
}


// A layout anchor representing the height of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/heightAnchor

func (v_ View) HeightAnchor() NSLayoutDimension {
	rv := objc.Send[NSLayoutDimension](v_.ID, objc.Sel("heightAnchor"))
	return rv
}


// A Boolean value indicating whether the view is being rendered as part of a live resizing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/inLiveResize

func (v_ View) InLiveResize() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("inLiveResize"))
	return rv
}


// The text input context object for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/inputContext

func (v_ View) InputContext() NSTextInputContext {
	rv := objc.Send[NSTextInputContext](v_.ID, objc.Sel("inputContext"))
	return rv
}


// The natural size for the receiving view, considering only properties of the view itself.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/intrinsicContentSize

func (v_ View) IntrinsicContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](v_.ID, objc.Sel("intrinsicContentSize"))
	return rv
}


// A Boolean value that indicates whether views support responsive scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isCompatibleWithResponsiveScrolling

func (v_ View) CompatibleWithResponsiveScrolling() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("compatibleWithResponsiveScrolling"))
	return rv
}


// A Boolean value indicating whether the view or one of its ancestors is being drawn for a find indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isDrawingFindIndicator

func (v_ View) DrawingFindIndicator() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("drawingFindIndicator"))
	return rv
}


// A Boolean value indicating whether the view uses a flipped coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isFlipped

func (v_ View) Flipped() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("flipped"))
	return rv
}


// A Boolean value indicating whether the view is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isHidden

func (v_ View) Hidden() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("hidden"))
	return rv
}


// A Boolean value indicating whether the view is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isHidden

func (v_ View) SetHidden(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHidden:"), value)
}


// A Boolean value indicating whether the view is hidden from sight because it, or one of its ancestors, is marked as hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isHiddenOrHasHiddenAncestor

func (v_ View) HiddenOrHasHiddenAncestor() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("hiddenOrHasHiddenAncestor"))
	return rv
}


// A Boolean value that indicates whether the view’s horizontal size constraints are active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isHorizontalContentSizeConstraintActive

func (v_ View) HorizontalContentSizeConstraintActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("horizontalContentSizeConstraintActive"))
	return rv
}


// A Boolean value that indicates whether the view’s horizontal size constraints are active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isHorizontalContentSizeConstraintActive

func (v_ View) SetHorizontalContentSizeConstraintActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHorizontalContentSizeConstraintActive:"), value)
}


// A Boolean value indicating whether the view is in full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isInFullScreenMode

func (v_ View) InFullScreenMode() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("inFullScreenMode"))
	return rv
}


// A Boolean value indicating whether the view fills its frame rectangle with opaque content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isOpaque

func (v_ View) Opaque() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("opaque"))
	return rv
}


// A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isRotatedFromBase

func (v_ View) RotatedFromBase() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rotatedFromBase"))
	return rv
}


// A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds, or has been scaled from the window’s base coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isRotatedOrScaledFromBase

func (v_ View) RotatedOrScaledFromBase() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rotatedOrScaledFromBase"))
	return rv
}


// A Boolean value that indicates whether the view’s vertical size constraints are active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isVerticalContentSizeConstraintActive

func (v_ View) VerticalContentSizeConstraintActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("verticalContentSizeConstraintActive"))
	return rv
}


// A Boolean value that indicates whether the view’s vertical size constraints are active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isVerticalContentSizeConstraintActive

func (v_ View) SetVerticalContentSizeConstraintActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVerticalContentSizeConstraintActive:"), value)
}


// A layout anchor representing the baseline for the bottommost line of text in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/lastBaselineAnchor

func (v_ View) LastBaselineAnchor() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](v_.ID, objc.Sel("lastBaselineAnchor"))
	return rv
}


// The distance (in points) between the bottom of the view’s alignment rectangle and its bottommost baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/lastBaselineOffsetFromBottom

func (v_ View) LastBaselineOffsetFromBottom() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("lastBaselineOffsetFromBottom"))
	return rv
}


// The Core Animation layer that the view uses as its backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layer

func (v_ View) Layer() quartzcore.Layer {
	rv := objc.Send[quartzcore.Layer](v_.ID, objc.Sel("layer"))
	return rv
}


// The Core Animation layer that the view uses as its backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layer

func (v_ View) SetLayer(value quartzcore.ILayer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLayer:"), value)
}


// The current layer contents placement policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerContentsPlacement-swift.property

func (v_ View) LayerContentsPlacement() ViewLayerContentsPlacement {
	rv := objc.Send[ViewLayerContentsPlacement](v_.ID, objc.Sel("layerContentsPlacement"))
	return rv
}


// The current layer contents placement policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerContentsPlacement-swift.property

func (v_ View) SetLayerContentsPlacement(value IViewLayerContentsPlacement) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLayerContentsPlacement:"), value)
}


// The contents redraw policy for the view’s layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerContentsRedrawPolicy-swift.property

func (v_ View) LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy {
	rv := objc.Send[ViewLayerContentsRedrawPolicy](v_.ID, objc.Sel("layerContentsRedrawPolicy"))
	return rv
}


// The contents redraw policy for the view’s layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerContentsRedrawPolicy-swift.property

func (v_ View) SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLayerContentsRedrawPolicy:"), value)
}


// A Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerUsesCoreImageFilters

func (v_ View) LayerUsesCoreImageFilters() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("layerUsesCoreImageFilters"))
	return rv
}


// A Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerUsesCoreImageFilters

func (v_ View) SetLayerUsesCoreImageFilters(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLayerUsesCoreImageFilters:"), value)
}


// The array of layout guide objects owned by this view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layoutGuides

func (v_ View) LayoutGuides() []LayoutGuide {
	rv := objc.Send[[]LayoutGuide](v_.ID, objc.Sel("layoutGuides"))
	return rv
}


// A layout guide that provides the recommended amount of padding for content inside of a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layoutMarginsGuide

func (v_ View) LayoutMarginsGuide() NSLayoutGuide {
	rv := objc.Send[NSLayoutGuide](v_.ID, objc.Sel("layoutMarginsGuide"))
	return rv
}


// A layout anchor representing the leading edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/leadingAnchor

func (v_ View) LeadingAnchor() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](v_.ID, objc.Sel("leadingAnchor"))
	return rv
}


// A layout anchor representing the left edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/leftAnchor

func (v_ View) LeftAnchor() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](v_.ID, objc.Sel("leftAnchor"))
	return rv
}


// A Boolean value indicating whether the view can pass mouse down events through to its superviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/mouseDownCanMoveWindow

func (v_ View) MouseDownCanMoveWindow() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("mouseDownCanMoveWindow"))
	return rv
}


// A Boolean value that determines whether the view needs to be redrawn before being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay

func (v_ View) NeedsDisplay() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsDisplay"))
	return rv
}


// A Boolean value that determines whether the view needs to be redrawn before being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay

func (v_ View) SetNeedsDisplay(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNeedsDisplay:"), value)
}


// A Boolean value indicating whether the view needs a layout pass before it can be drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsLayout

func (v_ View) NeedsLayout() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsLayout"))
	return rv
}


// A Boolean value indicating whether the view needs a layout pass before it can be drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsLayout

func (v_ View) SetNeedsLayout(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNeedsLayout:"), value)
}


// A Boolean value indicating whether the view needs its panel to become the key window before it can handle keyboard input and navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsPanelToBecomeKey

func (v_ View) NeedsPanelToBecomeKey() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsPanelToBecomeKey"))
	return rv
}


// A Boolean value indicating whether the view’s constraints need to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsUpdateConstraints

func (v_ View) NeedsUpdateConstraints() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsUpdateConstraints"))
	return rv
}


// A Boolean value indicating whether the view’s constraints need to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsUpdateConstraints

func (v_ View) SetNeedsUpdateConstraints(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNeedsUpdateConstraints:"), value)
}


// The view object that follows the current view in the key view loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/nextKeyView

func (v_ View) NextKeyView() NSView {
	rv := objc.Send[NSView](v_.ID, objc.Sel("nextKeyView"))
	return rv
}


// The view object that follows the current view in the key view loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/nextKeyView

func (v_ View) SetNextKeyView(value IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNextKeyView:"), value)
}


// The closest view object in the key view loop that follows the current view in the key view loop and accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/nextValidKeyView

func (v_ View) NextValidKeyView() NSView {
	rv := objc.Send[NSView](v_.ID, objc.Sel("nextValidKeyView"))
	return rv
}


// The view’s closest opaque ancestor, which might be the view itself.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/opaqueAncestor

func (v_ View) OpaqueAncestor() NSView {
	rv := objc.Send[NSView](v_.ID, objc.Sel("opaqueAncestor"))
	return rv
}


// A default footer string that includes the current page number and page count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/pageFooter

func (v_ View) PageFooter() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](v_.ID, objc.Sel("pageFooter"))
	return rv
}


// A default header string that includes the print job title and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/pageHeader

func (v_ View) PageHeader() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](v_.ID, objc.Sel("pageHeader"))
	return rv
}


// A Boolean value indicating whether the view posts notifications when its bounds rectangle changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/postsBoundsChangedNotifications

func (v_ View) PostsBoundsChangedNotifications() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("postsBoundsChangedNotifications"))
	return rv
}


// A Boolean value indicating whether the view posts notifications when its bounds rectangle changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/postsBoundsChangedNotifications

func (v_ View) SetPostsBoundsChangedNotifications(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPostsBoundsChangedNotifications:"), value)
}


// A Boolean value indicating whether the view posts notifications when its frame rectangle changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/postsFrameChangedNotifications

func (v_ View) PostsFrameChangedNotifications() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("postsFrameChangedNotifications"))
	return rv
}


// A Boolean value indicating whether the view posts notifications when its frame rectangle changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/postsFrameChangedNotifications

func (v_ View) SetPostsFrameChangedNotifications(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPostsFrameChangedNotifications:"), value)
}


// When this property is true, any NSControls in the view or its descendants will be sized with compact metrics compatible with macOS 15 and earlier. Defaults to false
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/prefersCompactControlSizeMetrics

func (v_ View) PrefersCompactControlSizeMetrics() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("prefersCompactControlSizeMetrics"))
	return rv
}


// When this property is true, any NSControls in the view or its descendants will be sized with compact metrics compatible with macOS 15 and earlier. Defaults to false
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/prefersCompactControlSizeMetrics

func (v_ View) SetPrefersCompactControlSizeMetrics(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPrefersCompactControlSizeMetrics:"), value)
}


// The portion of the view that has been rendered and is available for responsive scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/preparedContentRect

func (v_ View) PreparedContentRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("preparedContentRect"))
	return rv
}


// The portion of the view that has been rendered and is available for responsive scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/preparedContentRect

func (v_ View) SetPreparedContentRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreparedContentRect:"), value)
}


// A Boolean value indicating whether the view optimizes live-resize operations by preserving content that has not moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/preservesContentDuringLiveResize

func (v_ View) PreservesContentDuringLiveResize() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("preservesContentDuringLiveResize"))
	return rv
}


// Configures the behavior and progression of the Force Touch trackpad when responding to touch input produced by the user when the cursor is over the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/pressureConfiguration

func (v_ View) PressureConfiguration() NSPressureConfiguration {
	rv := objc.Send[NSPressureConfiguration](v_.ID, objc.Sel("pressureConfiguration"))
	return rv
}


// Configures the behavior and progression of the Force Touch trackpad when responding to touch input produced by the user when the cursor is over the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/pressureConfiguration

func (v_ View) SetPressureConfiguration(value IPressureConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPressureConfiguration:"), value)
}


// The view object preceding the current view in the key view loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/previousKeyView

func (v_ View) PreviousKeyView() NSView {
	rv := objc.Send[NSView](v_.ID, objc.Sel("previousKeyView"))
	return rv
}


// The closest view object in the key view loop that precedes the current view and accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/previousValidKeyView

func (v_ View) PreviousValidKeyView() NSView {
	rv := objc.Send[NSView](v_.ID, objc.Sel("previousValidKeyView"))
	return rv
}


// The view’s print job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/printJobTitle

func (v_ View) PrintJobTitle() string {
	rv := objc.Send[string](v_.ID, objc.Sel("printJobTitle"))
	return rv
}


// The rectangle identifying the portion of your view that did not change during a live resize operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectPreservedDuringLiveResize

func (v_ View) RectPreservedDuringLiveResize() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("rectPreservedDuringLiveResize"))
	return rv
}


// The array of pasteboard drag types that the view can accept.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/registeredDraggedTypes

func (v_ View) RegisteredDraggedTypes() []string {
	rv := objc.Send[[]string](v_.ID, objc.Sel("registeredDraggedTypes"))
	return rv
}


// Returns a Boolean value indicating whether the view depends on the constraint-based layout system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/requiresConstraintBasedLayout

func (v_ View) RequiresConstraintBasedLayout() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("requiresConstraintBasedLayout"))
	return rv
}


// A layout anchor representing the right edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rightAnchor

func (v_ View) RightAnchor() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](v_.ID, objc.Sel("rightAnchor"))
	return rv
}


// The distances from the edges of your view that define the safe area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/safeAreaInsets

func (v_ View) SafeAreaInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("safeAreaInsets"))
	return rv
}


// The layout guide you use to position content inside your view’s safe area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/safeAreaLayoutGuide

func (v_ View) SafeAreaLayoutGuide() NSLayoutGuide {
	rv := objc.Send[NSLayoutGuide](v_.ID, objc.Sel("safeAreaLayoutGuide"))
	return rv
}


// A rectangle in the view’s coordinate system that contains the unobscured portion of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/safeAreaRect

func (v_ View) SafeAreaRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("safeAreaRect"))
	return rv
}


// The shadow displayed underneath the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/shadow

func (v_ View) Shadow() NSShadow {
	rv := objc.Send[NSShadow](v_.ID, objc.Sel("shadow"))
	return rv
}


// The shadow displayed underneath the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/shadow

func (v_ View) SetShadow(value IShadow) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShadow:"), value)
}


// The array of views embedded in the current view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/subviews

func (v_ View) Subviews() []View {
	rv := objc.Send[[]View](v_.ID, objc.Sel("subviews"))
	return rv
}


// The array of views embedded in the current view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/subviews

func (v_ View) SetSubviews(value []View) {
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
	objc.Send[objc.ID](v_.ID, objc.Sel("setSubviews:"), nsArray)
}


// The view that is the parent of the current view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/superview

func (v_ View) Superview() NSView {
	rv := objc.Send[NSView](v_.ID, objc.Sel("superview"))
	return rv
}


// The view’s tag, which is an integer that you use to identify the view within your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/tag

func (v_ View) Tag() int {
	rv := objc.Send[int](v_.ID, objc.Sel("tag"))
	return rv
}


// The text for the view’s tooltip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/toolTip

func (v_ View) ToolTip() string {
	rv := objc.Send[string](v_.ID, objc.Sel("toolTip"))
	return rv
}


// The text for the view’s tooltip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/toolTip

func (v_ View) SetToolTip(value string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setToolTip:"), objc.String(value))
}


// A layout anchor representing the top edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/topAnchor

func (v_ View) TopAnchor() NSLayoutYAxisAnchor {
	rv := objc.Send[NSLayoutYAxisAnchor](v_.ID, objc.Sel("topAnchor"))
	return rv
}


// An array of the view’s tracking areas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/trackingAreas

func (v_ View) TrackingAreas() []TrackingArea {
	rv := objc.Send[[]TrackingArea](v_.ID, objc.Sel("trackingAreas"))
	return rv
}


// A layout anchor representing the trailing edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/trailingAnchor

func (v_ View) TrailingAnchor() NSLayoutXAxisAnchor {
	rv := objc.Send[NSLayoutXAxisAnchor](v_.ID, objc.Sel("trailingAnchor"))
	return rv
}


// A Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/translatesAutoresizingMaskIntoConstraints

func (v_ View) TranslatesAutoresizingMaskIntoConstraints() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("translatesAutoresizingMaskIntoConstraints"))
	return rv
}


// A Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/translatesAutoresizingMaskIntoConstraints

func (v_ View) SetTranslatesAutoresizingMaskIntoConstraints(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTranslatesAutoresizingMaskIntoConstraints:"), value)
}


// The layout direction for content in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/userInterfaceLayoutDirection

func (v_ View) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](v_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}


// The layout direction for content in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/userInterfaceLayoutDirection

func (v_ View) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}


// The portion of the view that isn’t clipped by its superviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/visibleRect

func (v_ View) VisibleRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](v_.ID, objc.Sel("visibleRect"))
	return rv
}


// A Boolean value indicating whether the view wants an OpenGL backing surface with a resolution greater than 1 pixel per point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsBestResolutionOpenGLSurface

func (v_ View) WantsBestResolutionOpenGLSurface() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsBestResolutionOpenGLSurface"))
	return rv
}


// A Boolean value indicating whether the view wants an OpenGL backing surface with a resolution greater than 1 pixel per point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsBestResolutionOpenGLSurface

func (v_ View) SetWantsBestResolutionOpenGLSurface(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWantsBestResolutionOpenGLSurface:"), value)
}


// A Boolean value indicating whether AppKit’s default clipping behavior is in effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsDefaultClipping

func (v_ View) WantsDefaultClipping() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsDefaultClipping"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsExtendedDynamicRangeOpenGLSurface

func (v_ View) WantsExtendedDynamicRangeOpenGLSurface() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsExtendedDynamicRangeOpenGLSurface"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsExtendedDynamicRangeOpenGLSurface

func (v_ View) SetWantsExtendedDynamicRangeOpenGLSurface(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWantsExtendedDynamicRangeOpenGLSurface:"), value)
}


// A Boolean value indicating whether the view uses a layer as its backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsLayer

func (v_ View) WantsLayer() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsLayer"))
	return rv
}


// A Boolean value indicating whether the view uses a layer as its backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsLayer

func (v_ View) SetWantsLayer(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWantsLayer:"), value)
}


// A Boolean value indicating whether the view wants resting touches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsRestingTouches

func (v_ View) WantsRestingTouches() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsRestingTouches"))
	return rv
}


// A Boolean value indicating whether the view wants resting touches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsRestingTouches

func (v_ View) SetWantsRestingTouches(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWantsRestingTouches:"), value)
}


// A Boolean value indicating which drawing path the view takes when updating its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsUpdateLayer

func (v_ View) WantsUpdateLayer() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsUpdateLayer"))
	return rv
}


// The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as small images or text columns from being divided across pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/widthAdjustLimit

func (v_ View) WidthAdjustLimit() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("widthAdjustLimit"))
	return rv
}


// A layout anchor representing the width of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/widthAnchor

func (v_ View) WidthAnchor() NSLayoutDimension {
	rv := objc.Send[NSLayoutDimension](v_.ID, objc.Sel("widthAnchor"))
	return rv
}


// The view’s window object, if it is installed in a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/window

func (v_ View) Window() NSWindow {
	rv := objc.Send[NSWindow](v_.ID, objc.Sel("window"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/writingToolsCoordinator

func (v_ View) WritingToolsCoordinator() NSWritingToolsCoordinator {
	rv := objc.Send[NSWritingToolsCoordinator](v_.ID, objc.Sel("writingToolsCoordinator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/writingToolsCoordinator

func (v_ View) SetWritingToolsCoordinator(value IWritingToolsCoordinator) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWritingToolsCoordinator:"), value)
}


