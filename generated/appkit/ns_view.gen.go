// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/coreimage"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/quartzcore"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSView */


/* debug [class_header]: Header for NSView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for View */
// An interface definition for the [View] class.
type IView interface {
	IResponder
	
/* debug [class_interface_properties]: Properties for View */
	// properties:
	AcceptsTouchEvents() bool
	SetAcceptsTouchEvents(value bool)
	AdditionalSafeAreaInsets() foundation.EdgeInsets
	SetAdditionalSafeAreaInsets(value foundation.EdgeInsets)
	AlignmentRectInsets() foundation.EdgeInsets
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
	SetBackgroundFilters(value []coreimage.Filter)
	BaselineOffsetFromBottom() float64
	BottomAnchor() ILayoutYAxisAnchor
	Bounds() Rect /* not a class type */
	SetBounds(value Rect /* not a class type */)
	BoundsRotation() float64
	SetBoundsRotation(value float64)
	CanBecomeKeyView() bool
	CanDraw() bool
	CanDrawConcurrently() bool
	SetCanDrawConcurrently(value bool)
	CanDrawSubviewsIntoLayer() bool
	SetCanDrawSubviewsIntoLayer(value bool)
	CandidateListTouchBarItem() ICandidateListTouchBarItem
	CenterXAnchor() ILayoutXAxisAnchor
	CenterYAnchor() ILayoutYAxisAnchor
	ClipsToBounds() bool
	SetClipsToBounds(value bool)
	CompositingFilter() coreimage.Filter
	SetCompositingFilter(value coreimage.Filter)
	Constraints() []LayoutConstraint
	ContentFilters() []coreimage.Filter
	SetContentFilters(value []coreimage.Filter)
	EnclosingMenuItem() IMenuItem
	EnclosingScrollView() IScrollView
	FirstBaselineAnchor() ILayoutYAxisAnchor
	FirstBaselineOffsetFromTop() float64
	FittingSize() Size /* not a class type */
	FocusRingMaskBounds() Rect /* not a class type */
	FocusRingType() FocusRingType
	SetFocusRingType(value FocusRingType)
	Frame() Rect /* not a class type */
	SetFrame(value Rect /* not a class type */)
	FrameCenterRotation() float64
	SetFrameCenterRotation(value float64)
	FrameRotation() float64
	SetFrameRotation(value float64)
	GestureRecognizers() []objc.IObject /* cross-framework: GestureRecognizer */
	SetGestureRecognizers(value []objc.IObject /* cross-framework: GestureRecognizer */)
	HasAmbiguousLayout() bool
	HeightAdjustLimit() float64
	HeightAnchor() ILayoutDimension
	InLiveResize() bool
	InputContext() ITextInputContext
	IntrinsicContentSize() Size /* not a class type */
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
	LastBaselineAnchor() ILayoutYAxisAnchor
	LastBaselineOffsetFromBottom() float64
	Layer() objectivec.IObject
	SetLayer(value objectivec.IObject)
	LayerContentsPlacement() ViewLayerContentsPlacement
	SetLayerContentsPlacement(value ViewLayerContentsPlacement)
	LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy
	SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy)
	LayerUsesCoreImageFilters() bool
	SetLayerUsesCoreImageFilters(value bool)
	LayoutGuides() []LayoutGuide
	LayoutMarginsGuide() ILayoutGuide
	LeadingAnchor() ILayoutXAxisAnchor
	LeftAnchor() ILayoutXAxisAnchor
	MouseDownCanMoveWindow() bool
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
	NeedsLayout() bool
	SetNeedsLayout(value bool)
	NeedsPanelToBecomeKey() bool
	NeedsUpdateConstraints() bool
	SetNeedsUpdateConstraints(value bool)
	NextKeyView() IView
	SetNextKeyView(value IView)
	NextValidKeyView() IView
	OpaqueAncestor() IView
	PageFooter() foundation.AttributedString
	PageHeader() foundation.AttributedString
	PostsBoundsChangedNotifications() bool
	SetPostsBoundsChangedNotifications(value bool)
	PostsFrameChangedNotifications() bool
	SetPostsFrameChangedNotifications(value bool)
	PrefersCompactControlSizeMetrics() bool
	SetPrefersCompactControlSizeMetrics(value bool)
	PreparedContentRect() Rect /* not a class type */
	SetPreparedContentRect(value Rect /* not a class type */)
	PreservesContentDuringLiveResize() bool
	PressureConfiguration() IPressureConfiguration
	SetPressureConfiguration(value IPressureConfiguration)
	PreviousKeyView() IView
	PreviousValidKeyView() IView
	PrintJobTitle() objc.IObject /* cross-framework: NSString */
	RectPreservedDuringLiveResize() Rect /* not a class type */
	RegisteredDraggedTypes() []string
	RightAnchor() ILayoutXAxisAnchor
	SafeAreaInsets() foundation.EdgeInsets
	SafeAreaLayoutGuide() ILayoutGuide
	SafeAreaRect() Rect /* not a class type */
	Shadow() IShadow
	SetShadow(value IShadow)
	Subviews() []View
	SetSubviews(value []View)
	Superview() IView
	Tag() int
	ToolTip() objc.IObject /* cross-framework: NSString */
	SetToolTip(value objc.IObject /* cross-framework: NSString */)
	TopAnchor() ILayoutYAxisAnchor
	TrackingAreas() []TrackingArea
	TrailingAnchor() ILayoutXAxisAnchor
	TranslatesAutoresizingMaskIntoConstraints() bool
	SetTranslatesAutoresizingMaskIntoConstraints(value bool)
	UserInterfaceLayoutDirection() UserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection)
	VisibleRect() Rect /* not a class type */
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
	WidthAnchor() ILayoutDimension
	Window() IWindow
	WritingToolsCoordinator() IWritingToolsCoordinator
	SetWritingToolsCoordinator(value IWritingToolsCoordinator)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for View */
	// methods:
	AcceptsFirstMouse(event IEvent) bool
	AddConstraint(constraint ILayoutConstraint)
	AddConstraints(constraints []LayoutConstraint)
	AddCursorRectCursor(rect Rect /* not a class type */, object ICursor)
	AddGestureRecognizer(gestureRecognizer objc.IObject /* cross-framework: GestureRecognizer */)
	AddLayoutGuide(guide ILayoutGuide)
	AddSubview(view IView)
	AddSubviewPositionedRelativeTo(view IView, place WindowOrderingMode, otherView IView)
	AddToolTipRectOwnerUserData(rect Rect /* not a class type */, owner objc.IObject, data objectivec.IObject) ToolTipTag /* typedef */
	AddTrackingRectOwnerUserDataAssumeInside(rect Rect /* not a class type */, owner objc.IObject, data objectivec.IObject, flag bool) TrackingRectTag /* typedef */
	AddTrackingArea(trackingArea ITrackingArea)
	AdjustPageHeightNewTopBottomLimit(newBottom corefoundation.CGFloat, oldTop float64, oldBottom float64, bottomLimit float64)
	AdjustPageWidthNewLeftRightLimit(newRight corefoundation.CGFloat, oldLeft float64, oldRight float64, rightLimit float64)
	AdjustScroll(newVisible Rect /* not a class type */) Rect /* not a class type */
	AlignmentRectForFrame(frame Rect /* not a class type */) Rect /* not a class type */
	AncestorSharedWithView(view IView) IView
	Autoscroll(event IEvent) bool
	BackingAlignedRectOptions(rect Rect /* not a class type */, options AlignmentOptions /* not a class type */) Rect /* not a class type */
	BeginDocument()
	BeginDraggingSessionWithItemsEventSource(items []DraggingItem, event IEvent, source unsafe.Pointer) IDraggingSession
	BeginPageInRectAtPlacement(rect Rect /* not a class type */, location vision.Point)
	BitmapImageRepForCachingDisplayInRect(rect Rect /* not a class type */) IBitmapImageRep
	CacheDisplayInRectToBitmapImageRep(rect Rect /* not a class type */, bitmapImageRep IBitmapImageRep)
	CenterScanRect(rect Rect /* not a class type */) Rect /* not a class type */
	ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint
	ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority /* typedef */
	ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority /* typedef */
	ConvertPointFromView(point vision.Point, view IView) vision.Point
	ConvertSizeFromView(size Size /* not a class type */, view IView) Size /* not a class type */
	ConvertRectFromView(rect Rect /* not a class type */, view IView) Rect /* not a class type */
	ConvertRectToView(rect Rect /* not a class type */, view IView) Rect /* not a class type */
	ConvertSizeToView(size Size /* not a class type */, view IView) Size /* not a class type */
	ConvertPointToView(point vision.Point, view IView) vision.Point
	ConvertPointFromBacking(point vision.Point) vision.Point
	ConvertRectFromBacking(rect Rect /* not a class type */) Rect /* not a class type */
	ConvertSizeFromBacking(size Size /* not a class type */) Size /* not a class type */
	ConvertPointFromLayer(point vision.Point) vision.Point
	ConvertSizeFromLayer(size Size /* not a class type */) Size /* not a class type */
	ConvertRectFromLayer(rect Rect /* not a class type */) Rect /* not a class type */
	ConvertPointToBacking(point vision.Point) vision.Point
	ConvertRectToBacking(rect Rect /* not a class type */) Rect /* not a class type */
	ConvertSizeToBacking(size Size /* not a class type */) Size /* not a class type */
	ConvertRectToLayer(rect Rect /* not a class type */) Rect /* not a class type */
	ConvertSizeToLayer(size Size /* not a class type */) Size /* not a class type */
	ConvertPointToLayer(point vision.Point) vision.Point
	DataWithEPSInsideRect(rect Rect /* not a class type */) foundation.Data
	DataWithPDFInsideRect(rect Rect /* not a class type */) foundation.Data
	DidAddSubview(subview IView)
	DidCloseMenuWithEvent(menu IMenu, event IEvent)
	DiscardCursorRects()
	Display()
	DisplayRect(rect Rect /* not a class type */)
	DisplayIfNeeded()
	DisplayIfNeededInRect(rect Rect /* not a class type */)
	DisplayIfNeededIgnoringOpacity()
	DisplayIfNeededInRectIgnoringOpacity(rect Rect /* not a class type */)
	DisplayRectIgnoringOpacity(rect Rect /* not a class type */)
	DisplayRectIgnoringOpacityInContext(rect Rect /* not a class type */, context IGraphicsContext)
	DisplayLinkWithTargetSelector(target objc.IObject, selector objc.SEL) quartzcore.DisplayLink
	DrawRect(dirtyRect Rect /* not a class type */)
	DrawFocusRingMask()
	DrawPageBorderWithSize(borderSize Size /* not a class type */)
	EdgeInsetsForLayoutRegion(layoutRegion IViewLayoutRegion) foundation.EdgeInsets
	EndDocument()
	EndPage()
	EnterFullScreenModeWithOptions(screen IScreen, options foundation.IDictionary) bool
	ExerciseAmbiguityInLayout()
	ExitFullScreenModeWithOptions(options foundation.IDictionary)
	FrameForAlignmentRect(alignmentRect Rect /* not a class type */) Rect /* not a class type */
	GetRectsBeingDrawnCount(rects objectivec.IObject, count int)
	GetRectsExposedDuringLiveResizeCount(exposedRects Rect [ 4 ] /* not a class type */, count int)
	HitTest(point vision.Point) IView
	InvalidateIntrinsicContentSize()
	IsDescendantOf(view IView) bool
	MouseInRect(point vision.Point, rect Rect /* not a class type */) bool
	KnowsPageRange(range_ RangePointer /* not a class type */) bool
	Layout()
	LayoutGuideForLayoutRegion(layoutRegion IViewLayoutRegion) ILayoutGuide
	LayoutSubtreeIfNeeded()
	LocationOfPrintRect(rect Rect /* not a class type */) vision.Point
	MakeBackingLayer() objectivec.IObject
	MenuForEvent(event IEvent) IMenu
	NeedsToDrawRect(rect Rect /* not a class type */) bool
	NoteFocusRingMaskChanged()
	PerformKeyEquivalent(event IEvent) bool
	PrepareContentInRect(rect Rect /* not a class type */)
	PrepareForReuse()
	Print(sender objc.IObject)
	RectForLayoutRegion(layoutRegion IViewLayoutRegion) Rect /* not a class type */
	RectForPage(page int) Rect /* not a class type */
	RectForSmartMagnificationAtPointInRect(location vision.Point, visibleRect Rect /* not a class type */) Rect /* not a class type */
	ReflectScrolledClipView(clipView IClipView)
	RegisterForDraggedTypes(newTypes []string)
	RemoveAllToolTips()
	RemoveConstraint(constraint ILayoutConstraint)
	RemoveConstraints(constraints []LayoutConstraint)
	RemoveCursorRectCursor(rect Rect /* not a class type */, object ICursor)
	RemoveFromSuperview()
	RemoveFromSuperviewWithoutNeedingDisplay()
	RemoveGestureRecognizer(gestureRecognizer objc.IObject /* cross-framework: GestureRecognizer */)
	RemoveLayoutGuide(guide ILayoutGuide)
	RemoveToolTip(tag ToolTipTag /* typedef */)
	RemoveTrackingRect(tag TrackingRectTag /* typedef */)
	RemoveTrackingArea(trackingArea ITrackingArea)
	ReplaceSubviewWith(oldView IView, newView IView)
	ResetCursorRects()
	ResizeWithOldSuperviewSize(oldSize Size /* not a class type */)
	ResizeSubviewsWithOldSize(oldSize Size /* not a class type */)
	RotateByAngle(angle float64)
	RulerViewDidAddMarker(ruler IRulerView, marker IRulerMarker)
	RulerViewDidMoveMarker(ruler IRulerView, marker IRulerMarker)
	RulerViewDidRemoveMarker(ruler IRulerView, marker IRulerMarker)
	RulerViewHandleMouseDown(ruler IRulerView, event IEvent)
	RulerViewLocationForPoint(ruler IRulerView, point vision.Point) float64
	RulerViewPointForLocation(ruler IRulerView, point float64) vision.Point
	RulerViewShouldAddMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerViewShouldMoveMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerViewShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool
	RulerViewWillAddMarkerAtLocation(ruler IRulerView, marker IRulerMarker, location float64) float64
	RulerViewWillMoveMarkerToLocation(ruler IRulerView, marker IRulerMarker, location float64) float64
	RulerViewWillSetClientView(ruler IRulerView, newClient IView)
	ScaleUnitSquareToSize(newUnitSize Size /* not a class type */)
	ScrollPoint(point vision.Point)
	ScrollClipViewToPoint(clipView IClipView, point vision.Point)
	ScrollRectToVisible(rect Rect /* not a class type */) bool
	SetBoundsOrigin(newOrigin vision.Point)
	SetBoundsSize(newSize Size /* not a class type */)
	SetContentCompressionResistancePriorityForOrientation(priority LayoutPriority /* typedef */, orientation LayoutConstraintOrientation)
	SetContentHuggingPriorityForOrientation(priority LayoutPriority /* typedef */, orientation LayoutConstraintOrientation)
	SetFrameOrigin(newOrigin vision.Point)
	SetFrameSize(newSize Size /* not a class type */)
	SetKeyboardFocusRingNeedsDisplayInRect(rect Rect /* not a class type */)
	SetNeedsDisplayInRect(invalidRect Rect /* not a class type */)
	ShouldDelayWindowOrderingForEvent(event IEvent) bool
	ShowDefinitionForAttributedStringAtPoint(attrString foundation.AttributedString, textBaselineOrigin vision.Point)
	ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString foundation.AttributedString, targetRange corefoundation.Range, options foundation.IDictionary, originProvider unsafe.Pointer)
	SortSubviewsUsingFunctionContext(compare objectivec.IObject, context objectivec.IObject)
	TranslateOriginToPoint(translation vision.Point)
	TranslateRectsNeedingDisplayInRectBy(clipRect Rect /* not a class type */, delta Size /* not a class type */)
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
	ViewWithTag(tag int) IView
	WillOpenMenuWithEvent(menu IMenu, event IEvent)
	WillRemoveSubview(subview IView)
	WriteEPSInsideRectToPasteboard(rect Rect /* not a class type */, pasteboard IPasteboard)
	WritePDFInsideRectToPasteboard(rect Rect /* not a class type */, pasteboard IPasteboard)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for View */
// Alloc allocates a new instance without initialization.
func (vc _ViewClass) Alloc() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for View */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for View */

// Initializes a view using from data in the specified coder object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewViewWithCoder(coder foundation.Coder) View {
	instance := getViewClass().Alloc()
	rv := objc.Send[View](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewViewWithCoder */


// Initializes and returns a newly allocated object with a specified frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewViewWithFrame(frameRect Rect /* not a class type */) View {
	instance := getViewClass().Alloc()
	rv := objc.Send[View](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewViewWithFrame */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for View */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for View */

// Returns the default focus ring type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/defaultFocusRingType
func (vc _ViewClass) DefaultFocusRingType() FocusRingType {
	rv := objc.Send[FocusRingType](objc.ID(vc.class), objc.Sel("defaultFocusRingType"))
	return rv
}/* debug [class_properties_class/property]: defaultFocusRingType */

// Overridden by subclasses to return the default pop-up menu for instances of the receiving class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/defaultMenu
func (vc _ViewClass) DefaultMenu() IMenu {
	rv := objc.Send[Menu](objc.ID(vc.class), objc.Sel("defaultMenu"))
	return rv
}/* debug [class_properties_class/property]: defaultMenu */

// The currently focused view object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/focusView
func (vc _ViewClass) FocusView() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("focusView"))
	return rv
}/* debug [class_properties_class/property]: focusView */

// A Boolean value that indicates whether views support responsive scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isCompatibleWithResponsiveScrolling
func (vc _ViewClass) CompatibleWithResponsiveScrolling() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("compatibleWithResponsiveScrolling"))
	return rv
}/* debug [class_properties_class/property]: compatibleWithResponsiveScrolling */

// Returns a Boolean value indicating whether the view depends on the constraint-based layout system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/requiresConstraintBasedLayout
func (vc _ViewClass) RequiresConstraintBasedLayout() bool {
	rv := objc.Send[bool](objc.ID(vc.class), objc.Sel("requiresConstraintBasedLayout"))
	return rv
}/* debug [class_properties_class/property]: requiresConstraintBasedLayout */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for View */

// Returns a Boolean value that indicates whether the view accepts the initial mouse-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/acceptsFirstMouse(for:)
func (v_ View) AcceptsFirstMouse(event IEvent) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("acceptsFirstMouse:"), event)
	return rv
}/* debug [instance_methods/method]: AcceptsFirstMouse */


// Adds a constraint on the layout of the receiving view or its subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addConstraint(_:)
func (v_ View) AddConstraint(constraint ILayoutConstraint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addConstraint:"), constraint)
}/* debug [instance_methods/method]: AddConstraint */


// Adds multiple constraints on the layout of the receiving view or its subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addConstraints(_:)
func (v_ View) AddConstraints(constraints []LayoutConstraint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addConstraints:"), constraints)
}/* debug [instance_methods/method]: AddConstraints */


// Establishes the cursor to be used when the mouse pointer lies within a specified region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addCursorRect(_:cursor:)
func (v_ View) AddCursorRectCursor(rect Rect /* not a class type */, object ICursor) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addCursorRect:cursor:"), rect, object)
}/* debug [instance_methods/method]: AddCursorRectCursor */


// Attaches a gesture recognizer to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addGestureRecognizer(_:)
func (v_ View) AddGestureRecognizer(gestureRecognizer objc.IObject /* cross-framework: GestureRecognizer */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addGestureRecognizer:"), gestureRecognizer)
}/* debug [instance_methods/method]: AddGestureRecognizer */


// Adds the provided layout guide to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addLayoutGuide(_:)
func (v_ View) AddLayoutGuide(guide ILayoutGuide) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addLayoutGuide:"), guide)
}/* debug [instance_methods/method]: AddLayoutGuide */


// Adds a view to the view’s subviews so it’s displayed above its siblings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addSubview(_:)
func (v_ View) AddSubview(view IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addSubview:"), view)
}/* debug [instance_methods/method]: AddSubview */


// Inserts a view among the view’s subviews so it’s displayed immediately above or below another view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addSubview(_:positioned:relativeTo:)
func (v_ View) AddSubviewPositionedRelativeTo(view IView, place WindowOrderingMode, otherView IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addSubview:positioned:relativeTo:"), view, place, otherView)
}/* debug [instance_methods/method]: AddSubviewPositionedRelativeTo */


// Creates a tooltip for a defined area in the view and returns a tag that identifies the tooltip rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addToolTip(_:owner:userData:)
func (v_ View) AddToolTipRectOwnerUserData(rect Rect /* not a class type */, owner objc.IObject, data objectivec.IObject) ToolTipTag /* typedef */ {
	rv := objc.Send[int](v_.ID, objc.Sel("addToolTipRect:owner:userData:"), rect, owner, data)
	return rv
}/* debug [instance_methods/method]: AddToolTipRectOwnerUserData */


// Establishes an area for tracking mouse-entered and mouse-exited events within the view and returns a tag that identifies the tracking rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addTrackingRect(_:owner:userData:assumeInside:)
func (v_ View) AddTrackingRectOwnerUserDataAssumeInside(rect Rect /* not a class type */, owner objc.IObject, data objectivec.IObject, flag bool) TrackingRectTag /* typedef */ {
	rv := objc.Send[int](v_.ID, objc.Sel("addTrackingRect:owner:userData:assumeInside:"), rect, owner, data, flag)
	return rv
}/* debug [instance_methods/method]: AddTrackingRectOwnerUserDataAssumeInside */


// Adds a given tracking area to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addTrackingArea(_:)
func (v_ View) AddTrackingArea(trackingArea ITrackingArea) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addTrackingArea:"), trackingArea)
}/* debug [instance_methods/method]: AddTrackingArea */


// Overridden by subclasses to adjust page height during automatic pagination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/adjustPageHeightNew(_:top:bottom:limit:)
func (v_ View) AdjustPageHeightNewTopBottomLimit(newBottom corefoundation.CGFloat, oldTop float64, oldBottom float64, bottomLimit float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("adjustPageHeightNew:top:bottom:limit:"), newBottom, oldTop, oldBottom, bottomLimit)
}/* debug [instance_methods/method]: AdjustPageHeightNewTopBottomLimit */


// Overridden by subclasses to adjust page width during automatic pagination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/adjustPageWidthNew(_:left:right:limit:)
func (v_ View) AdjustPageWidthNewLeftRightLimit(newRight corefoundation.CGFloat, oldLeft float64, oldRight float64, rightLimit float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("adjustPageWidthNew:left:right:limit:"), newRight, oldLeft, oldRight, rightLimit)
}/* debug [instance_methods/method]: AdjustPageWidthNewLeftRightLimit */


// Overridden by subclasses to modify a given rectangle, returning the altered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/adjustScroll(_:)
func (v_ View) AdjustScroll(newVisible Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("adjustScroll:"), newVisible)
	return rv
}/* debug [instance_methods/method]: AdjustScroll */


// Returns the view’s alignment rectangle for a given frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/alignmentRect(forFrame:)
func (v_ View) AlignmentRectForFrame(frame Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("alignmentRectForFrame:"), frame)
	return rv
}/* debug [instance_methods/method]: AlignmentRectForFrame */


// Returns the closest ancestor shared by the view and another specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/ancestorShared(with:)
func (v_ View) AncestorSharedWithView(view IView) IView {
	rv := objc.Send[View](v_.ID, objc.Sel("ancestorSharedWithView:"), view)
	return rv
}/* debug [instance_methods/method]: AncestorSharedWithView */


// Scrolls the view’s closest ancestor object proportionally to the distance of an event that occurs outside of it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoscroll(with:)
func (v_ View) Autoscroll(event IEvent) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("autoscroll:"), event)
	return rv
}/* debug [instance_methods/method]: Autoscroll */


// Returns a backing store pixel-aligned rectangle in local view coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/backingAlignedRect(_:options:)
func (v_ View) BackingAlignedRectOptions(rect Rect /* not a class type */, options AlignmentOptions /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
	return rv
}/* debug [instance_methods/method]: BackingAlignedRectOptions */


// Invoked at the beginning of the printing session, this method sets up the current graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/beginDocument()
func (v_ View) BeginDocument() {
	objc.Send[objc.ID](v_.ID, objc.Sel("beginDocument"))
}/* debug [instance_methods/method]: BeginDocument */


// Initiates a dragging session with a group of dragging items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/beginDraggingSession(with:event:source:)
func (v_ View) BeginDraggingSessionWithItemsEventSource(items []DraggingItem, event IEvent, source unsafe.Pointer) IDraggingSession {
	rv := objc.Send[DraggingSession](v_.ID, objc.Sel("beginDraggingSessionWithItems:event:source:"), items, event, source)
	return rv
}/* debug [instance_methods/method]: BeginDraggingSessionWithItemsEventSource */


// Called at the beginning of each page, this method sets up the coordinate system so that a region inside the view’s bounds is translated to a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/beginPage(in:atPlacement:)
func (v_ View) BeginPageInRectAtPlacement(rect Rect /* not a class type */, location vision.Point) {
	objc.Send[objc.ID](v_.ID, objc.Sel("beginPageInRect:atPlacement:"), rect, location)
}/* debug [instance_methods/method]: BeginPageInRectAtPlacement */


// Returns a bitmap-representation object suitable for caching the specified portion of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/bitmapImageRepForCachingDisplay(in:)
func (v_ View) BitmapImageRepForCachingDisplayInRect(rect Rect /* not a class type */) IBitmapImageRep {
	rv := objc.Send[BitmapImageRep](v_.ID, objc.Sel("bitmapImageRepForCachingDisplayInRect:"), rect)
	return rv
}/* debug [instance_methods/method]: BitmapImageRepForCachingDisplayInRect */


// Draws the specified area of the view, and its descendants, into a provided bitmap-representation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/cacheDisplay(in:to:)
func (v_ View) CacheDisplayInRectToBitmapImageRep(rect Rect /* not a class type */, bitmapImageRep IBitmapImageRep) {
	objc.Send[objc.ID](v_.ID, objc.Sel("cacheDisplayInRect:toBitmapImageRep:"), rect, bitmapImageRep)
}/* debug [instance_methods/method]: CacheDisplayInRectToBitmapImageRep */


// Converts the corners of a specified rectangle to lie on the center of device pixels, which is useful in compensating for rendering overscanning when the coordinate system has been scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/centerScanRect(_:)
func (v_ View) CenterScanRect(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("centerScanRect:"), rect)
	return rv
}/* debug [instance_methods/method]: CenterScanRect */


// Returns the constraints impacting the layout of the view for a given orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/constraintsAffectingLayout(for:)
func (v_ View) ConstraintsAffectingLayoutForOrientation(orientation LayoutConstraintOrientation) []LayoutConstraint {
	rv := objc.Send[[]LayoutConstraint](v_.ID, objc.Sel("constraintsAffectingLayoutForOrientation:"), orientation)
	return rv
}/* debug [instance_methods/method]: ConstraintsAffectingLayoutForOrientation */


// Returns the priority with which a view resists being made smaller than its intrinsic size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/contentCompressionResistancePriority(for:)
func (v_ View) ContentCompressionResistancePriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority /* typedef */ {
	rv := objc.Send[float32](v_.ID, objc.Sel("contentCompressionResistancePriorityForOrientation:"), orientation)
	return rv
}/* debug [instance_methods/method]: ContentCompressionResistancePriorityForOrientation */


// Returns the priority with which a view resists being made larger than its intrinsic size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/contentHuggingPriority(for:)
func (v_ View) ContentHuggingPriorityForOrientation(orientation LayoutConstraintOrientation) LayoutPriority /* typedef */ {
	rv := objc.Send[float32](v_.ID, objc.Sel("contentHuggingPriorityForOrientation:"), orientation)
	return rv
}/* debug [instance_methods/method]: ContentHuggingPriorityForOrientation */


// Converts a point from the coordinate system of a given view to that of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-1dq9l
func (v_ View) ConvertPointFromView(point vision.Point, view IView) vision.Point {
	rv := objc.Send[vision.Point](v_.ID, objc.Sel("convertPoint:fromView:"), point, view)
	return rv
}/* debug [instance_methods/method]: ConvertPointFromView */


// Converts a size from another view’s coordinate system to that of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-40x0w
func (v_ View) ConvertSizeFromView(size Size /* not a class type */, view IView) Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("convertSize:fromView:"), size, view)
	return rv
}/* debug [instance_methods/method]: ConvertSizeFromView */


// Converts a rectangle from the coordinate system of another view to that of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-7fbb6
func (v_ View) ConvertRectFromView(rect Rect /* not a class type */, view IView) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("convertRect:fromView:"), rect, view)
	return rv
}/* debug [instance_methods/method]: ConvertRectFromView */


// Converts a rectangle from the view’s coordinate system to that of another view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-3cqqt
func (v_ View) ConvertRectToView(rect Rect /* not a class type */, view IView) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("convertRect:toView:"), rect, view)
	return rv
}/* debug [instance_methods/method]: ConvertRectToView */


// Converts a size from the view’s coordinate system to that of another view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-5nptx
func (v_ View) ConvertSizeToView(size Size /* not a class type */, view IView) Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("convertSize:toView:"), size, view)
	return rv
}/* debug [instance_methods/method]: ConvertSizeToView */


// Converts a point from the view’s coordinate system to that of a given view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-6u9ir
func (v_ View) ConvertPointToView(point vision.Point, view IView) vision.Point {
	rv := objc.Send[vision.Point](v_.ID, objc.Sel("convertPoint:toView:"), point, view)
	return rv
}/* debug [instance_methods/method]: ConvertPointToView */


// Converts a point from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-229ps
func (v_ View) ConvertPointFromBacking(point vision.Point) vision.Point {
	rv := objc.Send[vision.Point](v_.ID, objc.Sel("convertPointFromBacking:"), point)
	return rv
}/* debug [instance_methods/method]: ConvertPointFromBacking */


// Converts a rectangle from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-2njpa
func (v_ View) ConvertRectFromBacking(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}/* debug [instance_methods/method]: ConvertRectFromBacking */


// Converts a size from its pixel aligned backing store coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-4agf9
func (v_ View) ConvertSizeFromBacking(size Size /* not a class type */) Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("convertSizeFromBacking:"), size)
	return rv
}/* debug [instance_methods/method]: ConvertSizeFromBacking */


// Convert the point from the layer’s interior coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-3nsbu
func (v_ View) ConvertPointFromLayer(point vision.Point) vision.Point {
	rv := objc.Send[vision.Point](v_.ID, objc.Sel("convertPointFromLayer:"), point)
	return rv
}/* debug [instance_methods/method]: ConvertPointFromLayer */


// Convert the size from the layer’s interior coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-3usqp
func (v_ View) ConvertSizeFromLayer(size Size /* not a class type */) Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("convertSizeFromLayer:"), size)
	return rv
}/* debug [instance_methods/method]: ConvertSizeFromLayer */


// Convert the rectangle from the layer’s interior coordinate system to the view’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-8s5bi
func (v_ View) ConvertRectFromLayer(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("convertRectFromLayer:"), rect)
	return rv
}/* debug [instance_methods/method]: ConvertRectFromLayer */


// Converts a point from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-2xx45
func (v_ View) ConvertPointToBacking(point vision.Point) vision.Point {
	rv := objc.Send[vision.Point](v_.ID, objc.Sel("convertPointToBacking:"), point)
	return rv
}/* debug [instance_methods/method]: ConvertPointToBacking */


// Converts a rectangle from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-3zors
func (v_ View) ConvertRectToBacking(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("convertRectToBacking:"), rect)
	return rv
}/* debug [instance_methods/method]: ConvertRectToBacking */


// Converts a size from the view’s interior coordinate system to its pixel aligned backing store coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-4ra9y
func (v_ View) ConvertSizeToBacking(size Size /* not a class type */) Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("convertSizeToBacking:"), size)
	return rv
}/* debug [instance_methods/method]: ConvertSizeToBacking */


// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-160pw
func (v_ View) ConvertRectToLayer(rect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("convertRectToLayer:"), rect)
	return rv
}/* debug [instance_methods/method]: ConvertRectToLayer */


// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-2vozx
func (v_ View) ConvertSizeToLayer(size Size /* not a class type */) Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("convertSizeToLayer:"), size)
	return rv
}/* debug [instance_methods/method]: ConvertSizeToLayer */


// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-44u7d
func (v_ View) ConvertPointToLayer(point vision.Point) vision.Point {
	rv := objc.Send[vision.Point](v_.ID, objc.Sel("convertPointToLayer:"), point)
	return rv
}/* debug [instance_methods/method]: ConvertPointToLayer */


// Returns EPS data that draws the region of the view within a specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dataWithEPS(inside:)
func (v_ View) DataWithEPSInsideRect(rect Rect /* not a class type */) foundation.Data {
	rv := objc.Send[foundation.Data](v_.ID, objc.Sel("dataWithEPSInsideRect:"), rect)
	return rv
}/* debug [instance_methods/method]: DataWithEPSInsideRect */


// Returns PDF data that draws the region of the view within a specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dataWithPDF(inside:)
func (v_ View) DataWithPDFInsideRect(rect Rect /* not a class type */) foundation.Data {
	rv := objc.Send[foundation.Data](v_.ID, objc.Sel("dataWithPDFInsideRect:"), rect)
	return rv
}/* debug [instance_methods/method]: DataWithPDFInsideRect */


// Overridden by subclasses to perform additional actions when subviews are added to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/didAddSubview(_:)
func (v_ View) DidAddSubview(subview IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("didAddSubview:"), subview)
}/* debug [instance_methods/method]: DidAddSubview */


// Called after a contextual menu that was displayed from the receiving view has been closed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/didCloseMenu(_:with:)
func (v_ View) DidCloseMenuWithEvent(menu IMenu, event IEvent) {
	objc.Send[objc.ID](v_.ID, objc.Sel("didCloseMenu:withEvent:"), menu, event)
}/* debug [instance_methods/method]: DidCloseMenuWithEvent */


// Invalidates all cursor rectangles set up using .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/discardCursorRects()
func (v_ View) DiscardCursorRects() {
	objc.Send[objc.ID](v_.ID, objc.Sel("discardCursorRects"))
}/* debug [instance_methods/method]: DiscardCursorRects */


// Displays the view and all its subviews if possible, invoking each of the methods , , and as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/display()
func (v_ View) Display() {
	objc.Send[objc.ID](v_.ID, objc.Sel("display"))
}/* debug [instance_methods/method]: Display */


// Acts as , but confining drawing to a rectangular region of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/display(_:)
func (v_ View) DisplayRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayRect:"), rect)
}/* debug [instance_methods/method]: DisplayRect */


// Displays the view and all its subviews if any part of the view has been marked as needing display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeeded()
func (v_ View) DisplayIfNeeded() {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeeded"))
}/* debug [instance_methods/method]: DisplayIfNeeded */


// Acts as , confining drawing to a specified region of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeeded(_:)
func (v_ View) DisplayIfNeededInRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeededInRect:"), rect)
}/* debug [instance_methods/method]: DisplayIfNeededInRect */


// Acts as , except that this method doesn’t back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity()
func (v_ View) DisplayIfNeededIgnoringOpacity() {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeededIgnoringOpacity"))
}/* debug [instance_methods/method]: DisplayIfNeededIgnoringOpacity */


// Acts as , but confining drawing to and not backing up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity(_:)
func (v_ View) DisplayIfNeededInRectIgnoringOpacity(rect Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeededInRectIgnoringOpacity:"), rect)
}/* debug [instance_methods/method]: DisplayIfNeededInRectIgnoringOpacity */


// Displays the view but confines drawing to a specified region and does not back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIgnoringOpacity(_:)
func (v_ View) DisplayRectIgnoringOpacity(rect Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayRectIgnoringOpacity:"), rect)
}/* debug [instance_methods/method]: DisplayRectIgnoringOpacity */


// Causes the view and its descendants to be redrawn to the specified graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIgnoringOpacity(_:in:)
func (v_ View) DisplayRectIgnoringOpacityInContext(rect Rect /* not a class type */, context IGraphicsContext) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayRectIgnoringOpacity:inContext:"), rect, context)
}/* debug [instance_methods/method]: DisplayRectIgnoringOpacityInContext */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayLink(target:selector:)
func (v_ View) DisplayLinkWithTargetSelector(target objc.IObject, selector objc.SEL) quartzcore.DisplayLink {
	rv := objc.Send[quartzcore.DisplayLink](v_.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return rv
}/* debug [instance_methods/method]: DisplayLinkWithTargetSelector */


// Overridden by subclasses to draw the view’s image within the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/draw(_:)
func (v_ View) DrawRect(dirtyRect Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawRect:"), dirtyRect)
}/* debug [instance_methods/method]: DrawRect */


// Draws the focus ring mask for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/drawFocusRingMask()
func (v_ View) DrawFocusRingMask() {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawFocusRingMask"))
}/* debug [instance_methods/method]: DrawFocusRingMask */


// Allows applications that use the AppKit pagination facility to draw additional marks on each logical page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/drawPageBorder(with:)
func (v_ View) DrawPageBorderWithSize(borderSize Size /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawPageBorderWithSize:"), borderSize)
}/* debug [instance_methods/method]: DrawPageBorderWithSize */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/edgeInsetsForLayoutRegion:
func (v_ View) EdgeInsetsForLayoutRegion(layoutRegion IViewLayoutRegion) foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](v_.ID, objc.Sel("edgeInsetsForLayoutRegion:"), layoutRegion)
	return rv
}/* debug [instance_methods/method]: EdgeInsetsForLayoutRegion */


// This method is invoked at the end of the printing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/endDocument()
func (v_ View) EndDocument() {
	objc.Send[objc.ID](v_.ID, objc.Sel("endDocument"))
}/* debug [instance_methods/method]: EndDocument */


// Writes the end of a conforming page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/endPage()
func (v_ View) EndPage() {
	objc.Send[objc.ID](v_.ID, objc.Sel("endPage"))
}/* debug [instance_methods/method]: EndPage */


// Sets the view to full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/enterFullScreenMode(_:withOptions:)
func (v_ View) EnterFullScreenModeWithOptions(screen IScreen, options foundation.IDictionary) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("enterFullScreenMode:withOptions:"), screen, options)
	return rv
}/* debug [instance_methods/method]: EnterFullScreenModeWithOptions */


// Randomly changes the frame of a view with an ambiguous layout between the different valid values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/exerciseAmbiguityInLayout()
func (v_ View) ExerciseAmbiguityInLayout() {
	objc.Send[objc.ID](v_.ID, objc.Sel("exerciseAmbiguityInLayout"))
}/* debug [instance_methods/method]: ExerciseAmbiguityInLayout */


// Instructs the view to exit full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/exitFullScreenMode(options:)
func (v_ View) ExitFullScreenModeWithOptions(options foundation.IDictionary) {
	objc.Send[objc.ID](v_.ID, objc.Sel("exitFullScreenModeWithOptions:"), options)
}/* debug [instance_methods/method]: ExitFullScreenModeWithOptions */


// Returns the view’s frame for a given alignment rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frame(forAlignmentRect:)
func (v_ View) FrameForAlignmentRect(alignmentRect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("frameForAlignmentRect:"), alignmentRect)
	return rv
}/* debug [instance_methods/method]: FrameForAlignmentRect */


// Returns by indirection a list of nonoverlapping rectangles that define the area the view is being asked to draw in .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/getRectsBeingDrawn(_:count:)
func (v_ View) GetRectsBeingDrawnCount(rects objectivec.IObject, count int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("getRectsBeingDrawn:count:"), rects, count)
}/* debug [instance_methods/method]: GetRectsBeingDrawnCount */


// Returns a list of rectangles indicating the newly exposed areas of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/getRectsExposedDuringLiveResize(_:count:)
func (v_ View) GetRectsExposedDuringLiveResizeCount(exposedRects Rect [ 4 ] /* not a class type */, count int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("getRectsExposedDuringLiveResize:count:"), exposedRects, count)
}/* debug [instance_methods/method]: GetRectsExposedDuringLiveResizeCount */


// Returns the farthest descendant of the view in the view hierarchy (including itself) that contains a specified point, or if that point lies completely outside the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/hitTest(_:)
func (v_ View) HitTest(point vision.Point) IView {
	rv := objc.Send[View](v_.ID, objc.Sel("hitTest:"), point)
	return rv
}/* debug [instance_methods/method]: HitTest */


// Invalidates the view’s intrinsic content size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/invalidateIntrinsicContentSize()
func (v_ View) InvalidateIntrinsicContentSize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("invalidateIntrinsicContentSize"))
}/* debug [instance_methods/method]: InvalidateIntrinsicContentSize */


// Returns a Boolean value that indicates whether the view is a subview of the specified view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isDescendant(of:)
func (v_ View) IsDescendantOf(view IView) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isDescendantOf:"), view)
	return rv
}/* debug [instance_methods/method]: IsDescendantOf */


// Returns whether a region of the view contains a specified point, accounting for whether the view is flipped or not.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isMousePoint(_:in:)
func (v_ View) MouseInRect(point vision.Point, rect Rect /* not a class type */) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("mouse:inRect:"), point, rect)
	return rv
}/* debug [instance_methods/method]: MouseInRect */


// Returns a Boolean value that indicates whether the view handles page boundaries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/knowsPageRange(_:)
func (v_ View) KnowsPageRange(range_ RangePointer /* not a class type */) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("knowsPageRange:"), range_)
	return rv
}/* debug [instance_methods/method]: KnowsPageRange */


// Perform layout in concert with the constraint-based layout system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layout()
func (v_ View) Layout() {
	objc.Send[objc.ID](v_.ID, objc.Sel("layout"))
}/* debug [instance_methods/method]: Layout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layoutGuideForLayoutRegion:
func (v_ View) LayoutGuideForLayoutRegion(layoutRegion IViewLayoutRegion) ILayoutGuide {
	rv := objc.Send[LayoutGuide](v_.ID, objc.Sel("layoutGuideForLayoutRegion:"), layoutRegion)
	return rv
}/* debug [instance_methods/method]: LayoutGuideForLayoutRegion */


// Updates the layout of the receiving view and its subviews based on the current views and constraints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layoutSubtreeIfNeeded()
func (v_ View) LayoutSubtreeIfNeeded() {
	objc.Send[objc.ID](v_.ID, objc.Sel("layoutSubtreeIfNeeded"))
}/* debug [instance_methods/method]: LayoutSubtreeIfNeeded */


// Invoked by to determine the location of the region of the view being printed on the physical page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/locationOfPrintRect(_:)
func (v_ View) LocationOfPrintRect(rect Rect /* not a class type */) vision.Point {
	rv := objc.Send[vision.Point](v_.ID, objc.Sel("locationOfPrintRect:"), rect)
	return rv
}/* debug [instance_methods/method]: LocationOfPrintRect */


// Creates the view’s backing layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/makeBackingLayer()
func (v_ View) MakeBackingLayer() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("makeBackingLayer"))
	return rv
}/* debug [instance_methods/method]: MakeBackingLayer */


// Overridden by subclasses to return a context-sensitive pop-up menu for a given mouse-down event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/menu(for:)
func (v_ View) MenuForEvent(event IEvent) IMenu {
	rv := objc.Send[Menu](v_.ID, objc.Sel("menuForEvent:"), event)
	return rv
}/* debug [instance_methods/method]: MenuForEvent */


// Returns a Boolean value indicating whether the specified rectangle intersects any part of the area that the view is being asked to draw.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsToDraw(_:)
func (v_ View) NeedsToDrawRect(rect Rect /* not a class type */) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsToDrawRect:"), rect)
	return rv
}/* debug [instance_methods/method]: NeedsToDrawRect */


// Invoked to notify the view that the focus ring mask requires updating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/noteFocusRingMaskChanged()
func (v_ View) NoteFocusRingMaskChanged() {
	objc.Send[objc.ID](v_.ID, objc.Sel("noteFocusRingMaskChanged"))
}/* debug [instance_methods/method]: NoteFocusRingMaskChanged */


// Implemented by subclasses to respond to key equivalents (also known as keyboard shortcuts).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/performKeyEquivalent(with:)
func (v_ View) PerformKeyEquivalent(event IEvent) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("performKeyEquivalent:"), event)
	return rv
}/* debug [instance_methods/method]: PerformKeyEquivalent */


// Prepares the overdraw region for drawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/prepareContent(in:)
func (v_ View) PrepareContentInRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("prepareContentInRect:"), rect)
}/* debug [instance_methods/method]: PrepareContentInRect */


// Restores the view to an initial state so that it can be reused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/prepareForReuse()
func (v_ View) PrepareForReuse() {
	objc.Send[objc.ID](v_.ID, objc.Sel("prepareForReuse"))
}/* debug [instance_methods/method]: PrepareForReuse */


// This action method opens the Print panel, and if the user chooses an option other than canceling, prints the view and all its subviews to the device specified in the Print panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/printView(_:)
func (v_ View) Print(sender objc.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("print:"), sender)
}/* debug [instance_methods/method]: Print */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectForLayoutRegion:
func (v_ View) RectForLayoutRegion(layoutRegion IViewLayoutRegion) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("rectForLayoutRegion:"), layoutRegion)
	return rv
}/* debug [instance_methods/method]: RectForLayoutRegion */


// Implemented by subclasses to determine the portion of the view to be printed for the specified page number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectForPage(_:)
func (v_ View) RectForPage(page int) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("rectForPage:"), page)
	return rv
}/* debug [instance_methods/method]: RectForPage */


// Returns the appropriate rectangle to use when magnifying around the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectForSmartMagnification(at:in:)
func (v_ View) RectForSmartMagnificationAtPointInRect(location vision.Point, visibleRect Rect /* not a class type */) Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("rectForSmartMagnificationAtPoint:inRect:"), location, visibleRect)
	return rv
}/* debug [instance_methods/method]: RectForSmartMagnificationAtPointInRect */


// Notifies a clip view’s superview that either the clip view’s bounds rectangle or the document view’s frame rectangle has changed, and that any indicators of the scroll position need to be adjusted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/reflectScrolledClipView(_:)
func (v_ View) ReflectScrolledClipView(clipView IClipView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("reflectScrolledClipView:"), clipView)
}/* debug [instance_methods/method]: ReflectScrolledClipView */


// Registers the pasteboard types that the view will accept as the destination of an image-dragging session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/registerForDraggedTypes(_:)
func (v_ View) RegisterForDraggedTypes(newTypes []string) {
	objc.Send[objc.ID](v_.ID, objc.Sel("registerForDraggedTypes:"), newTypes)
}/* debug [instance_methods/method]: RegisterForDraggedTypes */


// Removes all tooltips assigned to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeAllToolTips()
func (v_ View) RemoveAllToolTips() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeAllToolTips"))
}/* debug [instance_methods/method]: RemoveAllToolTips */


// Removes the specified constraint from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeConstraint(_:)
func (v_ View) RemoveConstraint(constraint ILayoutConstraint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeConstraint:"), constraint)
}/* debug [instance_methods/method]: RemoveConstraint */


// Removes the specified constraints from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeConstraints(_:)
func (v_ View) RemoveConstraints(constraints []LayoutConstraint) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeConstraints:"), constraints)
}/* debug [instance_methods/method]: RemoveConstraints */


// Completely removes a cursor rectangle from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeCursorRect(_:cursor:)
func (v_ View) RemoveCursorRectCursor(rect Rect /* not a class type */, object ICursor) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeCursorRect:cursor:"), rect, object)
}/* debug [instance_methods/method]: RemoveCursorRectCursor */


// Unlinks the view from its superview and its window, removes it from the responder chain, and invalidates its cursor rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeFromSuperview()
func (v_ View) RemoveFromSuperview() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeFromSuperview"))
}/* debug [instance_methods/method]: RemoveFromSuperview */


// Unlinks the view from its superview and its window and removes it from the responder chain, but does not invalidate its cursor rectangles to cause redrawing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeFromSuperviewWithoutNeedingDisplay()
func (v_ View) RemoveFromSuperviewWithoutNeedingDisplay() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeFromSuperviewWithoutNeedingDisplay"))
}/* debug [instance_methods/method]: RemoveFromSuperviewWithoutNeedingDisplay */


// Detaches a gesture recognizer from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeGestureRecognizer(_:)
func (v_ View) RemoveGestureRecognizer(gestureRecognizer objc.IObject /* cross-framework: GestureRecognizer */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeGestureRecognizer:"), gestureRecognizer)
}/* debug [instance_methods/method]: RemoveGestureRecognizer */


// Removes the provided layout guide from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeLayoutGuide(_:)
func (v_ View) RemoveLayoutGuide(guide ILayoutGuide) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeLayoutGuide:"), guide)
}/* debug [instance_methods/method]: RemoveLayoutGuide */


// Removes the tooltip identified by specified tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeToolTip(_:)
func (v_ View) RemoveToolTip(tag ToolTipTag /* typedef */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeToolTip:"), tag)
}/* debug [instance_methods/method]: RemoveToolTip */


// Removes the tracking rectangle identified by a tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeTrackingRect(_:)
func (v_ View) RemoveTrackingRect(tag TrackingRectTag /* typedef */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeTrackingRect:"), tag)
}/* debug [instance_methods/method]: RemoveTrackingRect */


// Removes a given tracking area from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeTrackingArea(_:)
func (v_ View) RemoveTrackingArea(trackingArea ITrackingArea) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeTrackingArea:"), trackingArea)
}/* debug [instance_methods/method]: RemoveTrackingArea */


// Replaces one of the view’s subviews with another view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/replaceSubview(_:with:)
func (v_ View) ReplaceSubviewWith(oldView IView, newView IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("replaceSubview:with:"), oldView, newView)
}/* debug [instance_methods/method]: ReplaceSubviewWith */


// Overridden by subclasses to define their default cursor rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/resetCursorRects()
func (v_ View) ResetCursorRects() {
	objc.Send[objc.ID](v_.ID, objc.Sel("resetCursorRects"))
}/* debug [instance_methods/method]: ResetCursorRects */


// Informs the view that the bounds size of its superview has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/resize(withOldSuperviewSize:)
func (v_ View) ResizeWithOldSuperviewSize(oldSize Size /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("resizeWithOldSuperviewSize:"), oldSize)
}/* debug [instance_methods/method]: ResizeWithOldSuperviewSize */


// Informs the view’s subviews that the view’s bounds rectangle size has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/resizeSubviews(withOldSize:)
func (v_ View) ResizeSubviewsWithOldSize(oldSize Size /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("resizeSubviewsWithOldSize:"), oldSize)
}/* debug [instance_methods/method]: ResizeSubviewsWithOldSize */


// Rotates the view’s bounds rectangle by a specified degree value around the origin of the coordinate system, (0.0, 0.0).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rotate(byDegrees:)
func (v_ View) RotateByAngle(angle float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rotateByAngle:"), angle)
}/* debug [instance_methods/method]: RotateByAngle */


// Informs the client that allowed the user to add .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didAdd:)
func (v_ View) RulerViewDidAddMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:didAddMarker:"), ruler, marker)
}/* debug [instance_methods/method]: RulerViewDidAddMarker */


// Informs the client that allowed the user to move .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didMove:)
func (v_ View) RulerViewDidMoveMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:didMoveMarker:"), ruler, marker)
}/* debug [instance_methods/method]: RulerViewDidMoveMarker */


// Informs the client that allowed the user to remove .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didRemove:)
func (v_ View) RulerViewDidRemoveMarker(ruler IRulerView, marker IRulerMarker) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:didRemoveMarker:"), ruler, marker)
}/* debug [instance_methods/method]: RulerViewDidRemoveMarker */


// Informs the client that the user has pressed the mouse button while the cursor is in the ruler area of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:handleMouseDownWith:)
func (v_ View) RulerViewHandleMouseDown(ruler IRulerView, event IEvent) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:handleMouseDown:"), ruler, event)
}/* debug [instance_methods/method]: RulerViewHandleMouseDown */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:locationFor:)
func (v_ View) RulerViewLocationForPoint(ruler IRulerView, point vision.Point) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("rulerView:locationForPoint:"), ruler, point)
	return rv
}/* debug [instance_methods/method]: RulerViewLocationForPoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:pointForLocation:)
func (v_ View) RulerViewPointForLocation(ruler IRulerView, point float64) vision.Point {
	rv := objc.Send[vision.Point](v_.ID, objc.Sel("rulerView:pointForLocation:"), ruler, point)
	return rv
}/* debug [instance_methods/method]: RulerViewPointForLocation */


// Requests permission for to add , an NSRulerMarker being dragged onto the ruler by the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldAdd:)
func (v_ View) RulerViewShouldAddMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rulerView:shouldAddMarker:"), ruler, marker)
	return rv
}/* debug [instance_methods/method]: RulerViewShouldAddMarker */


// Requests permission for to move .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldMove:)
func (v_ View) RulerViewShouldMoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rulerView:shouldMoveMarker:"), ruler, marker)
	return rv
}/* debug [instance_methods/method]: RulerViewShouldMoveMarker */


// Requests permission for to remove .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldRemove:)
func (v_ View) RulerViewShouldRemoveMarker(ruler IRulerView, marker IRulerMarker) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rulerView:shouldRemoveMarker:"), ruler, marker)
	return rv
}/* debug [instance_methods/method]: RulerViewShouldRemoveMarker */


// Informs the client that will add the new NSRulerMarker, .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willAdd:atLocation:)
func (v_ View) RulerViewWillAddMarkerAtLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("rulerView:willAddMarker:atLocation:"), ruler, marker, location)
	return rv
}/* debug [instance_methods/method]: RulerViewWillAddMarkerAtLocation */


// Informs the client that will move , an NSRulerMarker already on the ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willMove:toLocation:)
func (v_ View) RulerViewWillMoveMarkerToLocation(ruler IRulerView, marker IRulerMarker, location float64) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("rulerView:willMoveMarker:toLocation:"), ruler, marker, location)
	return rv
}/* debug [instance_methods/method]: RulerViewWillMoveMarkerToLocation */


// Informs the client view that is about to be appropriated by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willSetClientView:)
func (v_ View) RulerViewWillSetClientView(ruler IRulerView, newClient IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:willSetClientView:"), ruler, newClient)
}/* debug [instance_methods/method]: RulerViewWillSetClientView */


// Scales the view’s coordinate system so that the unit square scales to the specified dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scaleUnitSquare(to:)
func (v_ View) ScaleUnitSquareToSize(newUnitSize Size /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scaleUnitSquareToSize:"), newUnitSize)
}/* debug [instance_methods/method]: ScaleUnitSquareToSize */


// Scrolls the view’s closest ancestor object so a point in the view lies at the origin of the clip view’s bounds rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:)
func (v_ View) ScrollPoint(point vision.Point) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scrollPoint:"), point)
}/* debug [instance_methods/method]: ScrollPoint */


// Notifies the superview of a clip view that the clip view needs to reset the origin of its bounds rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:to:)
func (v_ View) ScrollClipViewToPoint(clipView IClipView, point vision.Point) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scrollClipView:toPoint:"), clipView, point)
}/* debug [instance_methods/method]: ScrollClipViewToPoint */


// Scrolls the view’s closest ancestor object the minimum distance needed so a specified region of the view becomes visible in the clip view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scrollToVisible(_:)
func (v_ View) ScrollRectToVisible(rect Rect /* not a class type */) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("scrollRectToVisible:"), rect)
	return rv
}/* debug [instance_methods/method]: ScrollRectToVisible */


// Sets the origin of the view’s bounds rectangle to a specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setBoundsOrigin(_:)
func (v_ View) SetBoundsOrigin(newOrigin vision.Point) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBoundsOrigin:"), newOrigin)
}/* debug [instance_methods/method]: SetBoundsOrigin */


// Sets the size of the view’s bounds rectangle to specified dimensions, inversely scaling its coordinate system relative to its frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setBoundsSize(_:)
func (v_ View) SetBoundsSize(newSize Size /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBoundsSize:"), newSize)
}/* debug [instance_methods/method]: SetBoundsSize */


// Sets the priority with which a view resists being made smaller than its intrinsic size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setContentCompressionResistancePriority(_:for:)
func (v_ View) SetContentCompressionResistancePriorityForOrientation(priority LayoutPriority /* typedef */, orientation LayoutConstraintOrientation) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setContentCompressionResistancePriority:forOrientation:"), priority, orientation)
}/* debug [instance_methods/method]: SetContentCompressionResistancePriorityForOrientation */


// Sets the priority with which a view resists being made larger than its intrinsic size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setContentHuggingPriority(_:for:)
func (v_ View) SetContentHuggingPriorityForOrientation(priority LayoutPriority /* typedef */, orientation LayoutConstraintOrientation) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setContentHuggingPriority:forOrientation:"), priority, orientation)
}/* debug [instance_methods/method]: SetContentHuggingPriorityForOrientation */


// Sets the origin of the view’s frame rectangle to the specified point, effectively repositioning it within its superview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setFrameOrigin(_:)
func (v_ View) SetFrameOrigin(newOrigin vision.Point) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameOrigin:"), newOrigin)
}/* debug [instance_methods/method]: SetFrameOrigin */


// Sets the size of the view’s frame rectangle to the specified dimensions, resizing it within its superview without affecting its coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setFrameSize(_:)
func (v_ View) SetFrameSize(newSize Size /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameSize:"), newSize)
}/* debug [instance_methods/method]: SetFrameSize */


// Invalidates the area around the focus ring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setKeyboardFocusRingNeedsDisplay(_:)
func (v_ View) SetKeyboardFocusRingNeedsDisplayInRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setKeyboardFocusRingNeedsDisplayInRect:"), rect)
}/* debug [instance_methods/method]: SetKeyboardFocusRingNeedsDisplayInRect */


// Marks the region of the view within the specified rectangle as needing display, increasing the view’s existing invalid region to include it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setNeedsDisplay(_:)
func (v_ View) SetNeedsDisplayInRect(invalidRect Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNeedsDisplayInRect:"), invalidRect)
}/* debug [instance_methods/method]: SetNeedsDisplayInRect */


// Allows the user to drag objects from the view without activating the app or moving the window of the view forward, possibly obscuring the destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/shouldDelayWindowOrdering(for:)
func (v_ View) ShouldDelayWindowOrderingForEvent(event IEvent) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("shouldDelayWindowOrderingForEvent:"), event)
	return rv
}/* debug [instance_methods/method]: ShouldDelayWindowOrderingForEvent */


// Shows a window displaying the definition of the attributed string at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/showDefinition(for:at:)
func (v_ View) ShowDefinitionForAttributedStringAtPoint(attrString foundation.AttributedString, textBaselineOrigin vision.Point) {
	objc.Send[objc.ID](v_.ID, objc.Sel("showDefinitionForAttributedString:atPoint:"), attrString, textBaselineOrigin)
}/* debug [instance_methods/method]: ShowDefinitionForAttributedStringAtPoint */


// Shows a window displaying the definition of the specified range of the attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/showDefinition(for:range:options:baselineOriginProvider:)
func (v_ View) ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString foundation.AttributedString, targetRange corefoundation.Range, options foundation.IDictionary, originProvider unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("showDefinitionForAttributedString:range:options:baselineOriginProvider:"), attrString, targetRange, options, originProvider)
}/* debug [instance_methods/method]: ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider */


// Orders the view’s immediate subviews using the specified comparator function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/sortSubviews(_:context:)
func (v_ View) SortSubviewsUsingFunctionContext(compare objectivec.IObject, context objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("sortSubviewsUsingFunction:context:"), compare, context)
}/* debug [instance_methods/method]: SortSubviewsUsingFunctionContext */


// Translates the view’s coordinate system so that its origin moves to a new location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/translateOrigin(to:)
func (v_ View) TranslateOriginToPoint(translation vision.Point) {
	objc.Send[objc.ID](v_.ID, objc.Sel("translateOriginToPoint:"), translation)
}/* debug [instance_methods/method]: TranslateOriginToPoint */


// Translates the display rectangles by the specified delta.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/translateRectsNeedingDisplay(in:by:)
func (v_ View) TranslateRectsNeedingDisplayInRectBy(clipRect Rect /* not a class type */, delta Size /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("translateRectsNeedingDisplayInRect:by:"), clipRect, delta)
}/* debug [instance_methods/method]: TranslateRectsNeedingDisplayInRectBy */


// Unregisters the view as a possible destination in a dragging session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/unregisterDraggedTypes()
func (v_ View) UnregisterDraggedTypes() {
	objc.Send[objc.ID](v_.ID, objc.Sel("unregisterDraggedTypes"))
}/* debug [instance_methods/method]: UnregisterDraggedTypes */


// Update constraints for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateConstraints()
func (v_ View) UpdateConstraints() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateConstraints"))
}/* debug [instance_methods/method]: UpdateConstraints */


// Updates the constraints for the receiving view and its subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateConstraintsForSubtreeIfNeeded()
func (v_ View) UpdateConstraintsForSubtreeIfNeeded() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateConstraintsForSubtreeIfNeeded"))
}/* debug [instance_methods/method]: UpdateConstraintsForSubtreeIfNeeded */


// Updates the view’s content by modifying its underlying layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateLayer()
func (v_ View) UpdateLayer() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateLayer"))
}/* debug [instance_methods/method]: UpdateLayer */


// Invoked automatically when the view’s geometry changes such that its tracking areas need to be recalculated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateTrackingAreas()
func (v_ View) UpdateTrackingAreas() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateTrackingAreas"))
}/* debug [instance_methods/method]: UpdateTrackingAreas */


// Responds when the view’s backing store properties change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidChangeBackingProperties()
func (v_ View) ViewDidChangeBackingProperties() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidChangeBackingProperties"))
}/* debug [instance_methods/method]: ViewDidChangeBackingProperties */


// Informs the view that its effective appearance changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidChangeEffectiveAppearance()
func (v_ View) ViewDidChangeEffectiveAppearance() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidChangeEffectiveAppearance"))
}/* debug [instance_methods/method]: ViewDidChangeEffectiveAppearance */


// Informs the view of the end of a live resize—the user has finished resizing the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidEndLiveResize()
func (v_ View) ViewDidEndLiveResize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidEndLiveResize"))
}/* debug [instance_methods/method]: ViewDidEndLiveResize */


// Invoked when the view is hidden, either directly, or in response to an ancestor being hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidHide()
func (v_ View) ViewDidHide() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidHide"))
}/* debug [instance_methods/method]: ViewDidHide */


// Informs the view that its superview has changed (possibly to ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidMoveToSuperview()
func (v_ View) ViewDidMoveToSuperview() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidMoveToSuperview"))
}/* debug [instance_methods/method]: ViewDidMoveToSuperview */


// Informs the view that it has been added to a new view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidMoveToWindow()
func (v_ View) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidMoveToWindow"))
}/* debug [instance_methods/method]: ViewDidMoveToWindow */


// Invoked when the view is unhidden, either directly, or in response to an ancestor being unhidden
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidUnhide()
func (v_ View) ViewDidUnhide() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidUnhide"))
}/* debug [instance_methods/method]: ViewDidUnhide */


// Informs the view that it’s required to draw content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillDraw()
func (v_ View) ViewWillDraw() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillDraw"))
}/* debug [instance_methods/method]: ViewWillDraw */


// Informs the view that its superview is about to change to the specified superview (which may be ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillMove(toSuperview:)
func (v_ View) ViewWillMoveToSuperview(newSuperview IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillMoveToSuperview:"), newSuperview)
}/* debug [instance_methods/method]: ViewWillMoveToSuperview */


// Informs the view that it’s being added to the view hierarchy of the specified window object (which may be ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillMove(toWindow:)
func (v_ View) ViewWillMoveToWindow(newWindow IWindow) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillMoveToWindow:"), newWindow)
}/* debug [instance_methods/method]: ViewWillMoveToWindow */


// Informs the view of the start of a live resize—the user has started resizing the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillStartLiveResize()
func (v_ View) ViewWillStartLiveResize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillStartLiveResize"))
}/* debug [instance_methods/method]: ViewWillStartLiveResize */


// Returns the view’s nearest descendant (including itself) with a specific tag, or if no subview has that tag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWithTag(_:)
func (v_ View) ViewWithTag(tag int) IView {
	rv := objc.Send[View](v_.ID, objc.Sel("viewWithTag:"), tag)
	return rv
}/* debug [instance_methods/method]: ViewWithTag */


// Called just before a contextual menu for a view is opened on screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/willOpenMenu(_:with:)
func (v_ View) WillOpenMenuWithEvent(menu IMenu, event IEvent) {
	objc.Send[objc.ID](v_.ID, objc.Sel("willOpenMenu:withEvent:"), menu, event)
}/* debug [instance_methods/method]: WillOpenMenuWithEvent */


// Overridden by subclasses to perform additional actions before subviews are removed from the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/willRemoveSubview(_:)
func (v_ View) WillRemoveSubview(subview IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("willRemoveSubview:"), subview)
}/* debug [instance_methods/method]: WillRemoveSubview */


// Writes EPS data that draws the region of the view within a specified rectangle onto a pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/writeEPS(inside:to:)
func (v_ View) WriteEPSInsideRectToPasteboard(rect Rect /* not a class type */, pasteboard IPasteboard) {
	objc.Send[objc.ID](v_.ID, objc.Sel("writeEPSInsideRect:toPasteboard:"), rect, pasteboard)
}/* debug [instance_methods/method]: WriteEPSInsideRectToPasteboard */


// Writes PDF data that draws the region of the view within a specified rectangle onto a pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/writePDF(inside:to:)
func (v_ View) WritePDFInsideRectToPasteboard(rect Rect /* not a class type */, pasteboard IPasteboard) {
	objc.Send[objc.ID](v_.ID, objc.Sel("writePDFInsideRect:toPasteboard:"), rect, pasteboard)
}/* debug [instance_methods/method]: WritePDFInsideRectToPasteboard */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for View */

// A Boolean value indicating whether the view accepts touch events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/acceptsTouchEvents
func (v_ View) AcceptsTouchEvents() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("acceptsTouchEvents"))
	return rv
}/* debug [instance_properties/getter]: acceptsTouchEvents */


// A Boolean value indicating whether the view accepts touch events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/acceptsTouchEvents
func (v_ View) SetAcceptsTouchEvents(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAcceptsTouchEvents:"), value)
}/* debug [instance_properties/setter]: acceptsTouchEvents */


// Custom insets that you specify to modify your view’s safe area
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/additionalSafeAreaInsets
func (v_ View) AdditionalSafeAreaInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](v_.ID, objc.Sel("additionalSafeAreaInsets"))
	return rv
}/* debug [instance_properties/getter]: additionalSafeAreaInsets */


// Custom insets that you specify to modify your view’s safe area
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/additionalSafeAreaInsets
func (v_ View) SetAdditionalSafeAreaInsets(value foundation.EdgeInsets) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAdditionalSafeAreaInsets:"), value)
}/* debug [instance_properties/setter]: additionalSafeAreaInsets */


// The insets (in points) from the view’s frame that define its content rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/alignmentRectInsets
func (v_ View) AlignmentRectInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](v_.ID, objc.Sel("alignmentRectInsets"))
	return rv
}/* debug [instance_properties/getter]: alignmentRectInsets */


// The types of touch interactions the view allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/allowedTouchTypes
func (v_ View) AllowedTouchTypes() TouchTypeMask {
	rv := objc.Send[TouchTypeMask](v_.ID, objc.Sel("allowedTouchTypes"))
	return rv
}/* debug [instance_properties/getter]: allowedTouchTypes */


// The types of touch interactions the view allows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/allowedTouchTypes
func (v_ View) SetAllowedTouchTypes(value TouchTypeMask) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAllowedTouchTypes:"), value)
}/* debug [instance_properties/setter]: allowedTouchTypes */


// A Boolean value indicating whether the view ensures it is vibrant on top of other content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/allowsVibrancy
func (v_ View) AllowsVibrancy() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("allowsVibrancy"))
	return rv
}/* debug [instance_properties/getter]: allowsVibrancy */


// The opacity of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/alphaValue
func (v_ View) AlphaValue() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("alphaValue"))
	return rv
}/* debug [instance_properties/getter]: alphaValue */


// The opacity of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/alphaValue
func (v_ View) SetAlphaValue(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAlphaValue:"), value)
}/* debug [instance_properties/setter]: alphaValue */


// A Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoresizesSubviews
func (v_ View) AutoresizesSubviews() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("autoresizesSubviews"))
	return rv
}/* debug [instance_properties/getter]: autoresizesSubviews */


// A Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoresizesSubviews
func (v_ View) SetAutoresizesSubviews(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAutoresizesSubviews:"), value)
}/* debug [instance_properties/setter]: autoresizesSubviews */


// The options that determine how the view is resized relative to its superview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoresizingMask-swift.property
func (v_ View) AutoresizingMask() AutoresizingMaskOptions {
	rv := objc.Send[AutoresizingMaskOptions](v_.ID, objc.Sel("autoresizingMask"))
	return rv
}/* debug [instance_properties/getter]: autoresizingMask */


// The options that determine how the view is resized relative to its superview.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoresizingMask-swift.property
func (v_ View) SetAutoresizingMask(value AutoresizingMaskOptions) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAutoresizingMask:"), value)
}/* debug [instance_properties/setter]: autoresizingMask */


// An array of Core Image filters to apply to the view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/backgroundFilters
func (v_ View) BackgroundFilters() []coreimage.Filter {
	rv := objc.Send[[]coreimage.Filter](v_.ID, objc.Sel("backgroundFilters"))
	return rv
}/* debug [instance_properties/getter]: backgroundFilters */


// An array of Core Image filters to apply to the view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/backgroundFilters
func (v_ View) SetBackgroundFilters(value []coreimage.Filter) {
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
}/* debug [instance_properties/setter]: backgroundFilters */


// The distance (in points) between the bottom of the view’s alignment rectangle and its baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/baselineOffsetFromBottom
func (v_ View) BaselineOffsetFromBottom() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("baselineOffsetFromBottom"))
	return rv
}/* debug [instance_properties/getter]: baselineOffsetFromBottom */


// A layout anchor representing the bottom edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/bottomAnchor
func (v_ View) BottomAnchor() ILayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](v_.ID, objc.Sel("bottomAnchor"))
	return rv
}/* debug [instance_properties/getter]: bottomAnchor */


// The view’s bounds rectangle, which expresses its location and size in its own coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/bounds
func (v_ View) Bounds() Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// The view’s bounds rectangle, which expresses its location and size in its own coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/bounds
func (v_ View) SetBounds(value Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBounds:"), value)
}/* debug [instance_properties/setter]: bounds */


// The angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/boundsRotation
func (v_ View) BoundsRotation() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("boundsRotation"))
	return rv
}/* debug [instance_properties/getter]: boundsRotation */


// The angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/boundsRotation
func (v_ View) SetBoundsRotation(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBoundsRotation:"), value)
}/* debug [instance_properties/setter]: boundsRotation */


// A Boolean value indicating whether the view can become key view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canBecomeKeyView
func (v_ View) CanBecomeKeyView() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canBecomeKeyView"))
	return rv
}/* debug [instance_properties/getter]: canBecomeKeyView */


// A Boolean value indicating whether drawing commands will produce any results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canDraw
func (v_ View) CanDraw() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canDraw"))
	return rv
}/* debug [instance_properties/getter]: canDraw */


// A Boolean value indicating whether the view can draw its contents on a background thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canDrawConcurrently
func (v_ View) CanDrawConcurrently() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canDrawConcurrently"))
	return rv
}/* debug [instance_properties/getter]: canDrawConcurrently */


// A Boolean value indicating whether the view can draw its contents on a background thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canDrawConcurrently
func (v_ View) SetCanDrawConcurrently(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCanDrawConcurrently:"), value)
}/* debug [instance_properties/setter]: canDrawConcurrently */


// A Boolean value indicating whether the view incorporates content from its subviews into its own layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canDrawSubviewsIntoLayer
func (v_ View) CanDrawSubviewsIntoLayer() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("canDrawSubviewsIntoLayer"))
	return rv
}/* debug [instance_properties/getter]: canDrawSubviewsIntoLayer */


// A Boolean value indicating whether the view incorporates content from its subviews into its own layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/canDrawSubviewsIntoLayer
func (v_ View) SetCanDrawSubviewsIntoLayer(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCanDrawSubviewsIntoLayer:"), value)
}/* debug [instance_properties/setter]: canDrawSubviewsIntoLayer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/candidateListTouchBarItem
func (v_ View) CandidateListTouchBarItem() ICandidateListTouchBarItem {
	rv := objc.Send[CandidateListTouchBarItem](v_.ID, objc.Sel("candidateListTouchBarItem"))
	return rv
}/* debug [instance_properties/getter]: candidateListTouchBarItem */


// A layout anchor representing the horizontal center of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/centerXAnchor
func (v_ View) CenterXAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](v_.ID, objc.Sel("centerXAnchor"))
	return rv
}/* debug [instance_properties/getter]: centerXAnchor */


// A layout anchor representing the vertical center of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/centerYAnchor
func (v_ View) CenterYAnchor() ILayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](v_.ID, objc.Sel("centerYAnchor"))
	return rv
}/* debug [instance_properties/getter]: centerYAnchor */


// A Boolean value that indicates whether the view, and its subviews, confine their drawing areas to the bounds of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/clipsToBounds
func (v_ View) ClipsToBounds() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("clipsToBounds"))
	return rv
}/* debug [instance_properties/getter]: clipsToBounds */


// A Boolean value that indicates whether the view, and its subviews, confine their drawing areas to the bounds of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/clipsToBounds
func (v_ View) SetClipsToBounds(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setClipsToBounds:"), value)
}/* debug [instance_properties/setter]: clipsToBounds */


// The Core Image filter used to composite the view’s contents with its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/compositingFilter
func (v_ View) CompositingFilter() coreimage.Filter {
	rv := objc.Send[coreimage.Filter](v_.ID, objc.Sel("compositingFilter"))
	return rv
}/* debug [instance_properties/getter]: compositingFilter */


// The Core Image filter used to composite the view’s contents with its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/compositingFilter
func (v_ View) SetCompositingFilter(value coreimage.Filter) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setCompositingFilter:"), value)
}/* debug [instance_properties/setter]: compositingFilter */


// Returns the constraints held by the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/constraints
func (v_ View) Constraints() []LayoutConstraint {
	rv := objc.Send[[]LayoutConstraint](v_.ID, objc.Sel("constraints"))
	return rv
}/* debug [instance_properties/getter]: constraints */


// An array of Core Image filters to apply to the contents of the view and its sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/contentFilters
func (v_ View) ContentFilters() []coreimage.Filter {
	rv := objc.Send[[]coreimage.Filter](v_.ID, objc.Sel("contentFilters"))
	return rv
}/* debug [instance_properties/getter]: contentFilters */


// An array of Core Image filters to apply to the contents of the view and its sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/contentFilters
func (v_ View) SetContentFilters(value []coreimage.Filter) {
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
}/* debug [instance_properties/setter]: contentFilters */


// Returns the default focus ring type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/defaultFocusRingType
func (v_ View) DefaultFocusRingType() FocusRingType {
	rv := objc.Send[FocusRingType](v_.ID, objc.Sel("defaultFocusRingType"))
	return rv
}/* debug [instance_properties/getter]: defaultFocusRingType */


// Overridden by subclasses to return the default pop-up menu for instances of the receiving class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/defaultMenu
func (v_ View) DefaultMenu() IMenu {
	rv := objc.Send[Menu](v_.ID, objc.Sel("defaultMenu"))
	return rv
}/* debug [instance_properties/getter]: defaultMenu */


// The menu item containing the view or any of its superviews in the view hierarchy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/enclosingMenuItem
func (v_ View) EnclosingMenuItem() IMenuItem {
	rv := objc.Send[MenuItem](v_.ID, objc.Sel("enclosingMenuItem"))
	return rv
}/* debug [instance_properties/getter]: enclosingMenuItem */


// The nearest ancestor scroll view that contains the current view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/enclosingScrollView
func (v_ View) EnclosingScrollView() IScrollView {
	rv := objc.Send[ScrollView](v_.ID, objc.Sel("enclosingScrollView"))
	return rv
}/* debug [instance_properties/getter]: enclosingScrollView */


// A layout anchor representing the baseline for the topmost line of text in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/firstBaselineAnchor
func (v_ View) FirstBaselineAnchor() ILayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](v_.ID, objc.Sel("firstBaselineAnchor"))
	return rv
}/* debug [instance_properties/getter]: firstBaselineAnchor */


// The distance (in points) between the top of the view’s alignment rectangle and its topmost baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/firstBaselineOffsetFromTop
func (v_ View) FirstBaselineOffsetFromTop() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("firstBaselineOffsetFromTop"))
	return rv
}/* debug [instance_properties/getter]: firstBaselineOffsetFromTop */


// The minimum size of the view that satisfies the constraints it holds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/fittingSize
func (v_ View) FittingSize() Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("fittingSize"))
	return rv
}/* debug [instance_properties/getter]: fittingSize */


// The focus ring mask bounds, specified in the view’s coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/focusRingMaskBounds
func (v_ View) FocusRingMaskBounds() Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("focusRingMaskBounds"))
	return rv
}/* debug [instance_properties/getter]: focusRingMaskBounds */


// The type of focus ring drawn around the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/focusRingType
func (v_ View) FocusRingType() FocusRingType {
	rv := objc.Send[FocusRingType](v_.ID, objc.Sel("focusRingType"))
	return rv
}/* debug [instance_properties/getter]: focusRingType */


// The type of focus ring drawn around the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/focusRingType
func (v_ View) SetFocusRingType(value FocusRingType) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFocusRingType:"), value)
}/* debug [instance_properties/setter]: focusRingType */


// The currently focused view object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/focusView
func (v_ View) FocusView() IView {
	rv := objc.Send[View](v_.ID, objc.Sel("focusView"))
	return rv
}/* debug [instance_properties/getter]: focusView */


// The view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frame
func (v_ View) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The view’s frame rectangle, which defines its position and size in its superview’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frame
func (v_ View) SetFrame(value Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrame:"), value)
}/* debug [instance_properties/setter]: frame */


// The rotation angle of the view around the center of its layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frameCenterRotation
func (v_ View) FrameCenterRotation() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("frameCenterRotation"))
	return rv
}/* debug [instance_properties/getter]: frameCenterRotation */


// The rotation angle of the view around the center of its layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frameCenterRotation
func (v_ View) SetFrameCenterRotation(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameCenterRotation:"), value)
}/* debug [instance_properties/setter]: frameCenterRotation */


// The angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frameRotation
func (v_ View) FrameRotation() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("frameRotation"))
	return rv
}/* debug [instance_properties/getter]: frameRotation */


// The angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frameRotation
func (v_ View) SetFrameRotation(value float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameRotation:"), value)
}/* debug [instance_properties/setter]: frameRotation */


// The gesture recognize objects currently attached to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/gestureRecognizers
func (v_ View) GestureRecognizers() []objc.IObject /* cross-framework: GestureRecognizer */ {
	rv := objc.Send[[]GestureRecognizer](v_.ID, objc.Sel("gestureRecognizers"))
	return rv
}/* debug [instance_properties/getter]: gestureRecognizers */


// The gesture recognize objects currently attached to the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/gestureRecognizers
func (v_ View) SetGestureRecognizers(value []objc.IObject /* cross-framework: GestureRecognizer */) {
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
}/* debug [instance_properties/setter]: gestureRecognizers */


// A Boolean value indicating whether the constraints impacting the layout of the view incompletely specify the location of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/hasAmbiguousLayout
func (v_ View) HasAmbiguousLayout() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("hasAmbiguousLayout"))
	return rv
}/* debug [instance_properties/getter]: hasAmbiguousLayout */


// The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as lines of text from being divided across pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/heightAdjustLimit
func (v_ View) HeightAdjustLimit() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("heightAdjustLimit"))
	return rv
}/* debug [instance_properties/getter]: heightAdjustLimit */


// A layout anchor representing the height of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/heightAnchor
func (v_ View) HeightAnchor() ILayoutDimension {
	rv := objc.Send[LayoutDimension](v_.ID, objc.Sel("heightAnchor"))
	return rv
}/* debug [instance_properties/getter]: heightAnchor */


// A Boolean value indicating whether the view is being rendered as part of a live resizing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/inLiveResize
func (v_ View) InLiveResize() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("inLiveResize"))
	return rv
}/* debug [instance_properties/getter]: inLiveResize */


// The text input context object for the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/inputContext
func (v_ View) InputContext() ITextInputContext {
	rv := objc.Send[TextInputContext](v_.ID, objc.Sel("inputContext"))
	return rv
}/* debug [instance_properties/getter]: inputContext */


// The natural size for the receiving view, considering only properties of the view itself.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/intrinsicContentSize
func (v_ View) IntrinsicContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](v_.ID, objc.Sel("intrinsicContentSize"))
	return rv
}/* debug [instance_properties/getter]: intrinsicContentSize */


// A Boolean value that indicates whether views support responsive scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isCompatibleWithResponsiveScrolling
func (v_ View) CompatibleWithResponsiveScrolling() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("compatibleWithResponsiveScrolling"))
	return rv
}/* debug [instance_properties/getter]: compatibleWithResponsiveScrolling */


// A Boolean value indicating whether the view or one of its ancestors is being drawn for a find indicator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isDrawingFindIndicator
func (v_ View) DrawingFindIndicator() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("drawingFindIndicator"))
	return rv
}/* debug [instance_properties/getter]: drawingFindIndicator */


// A Boolean value indicating whether the view uses a flipped coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isFlipped
func (v_ View) Flipped() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("flipped"))
	return rv
}/* debug [instance_properties/getter]: flipped */


// A Boolean value indicating whether the view is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isHidden
func (v_ View) Hidden() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// A Boolean value indicating whether the view is hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isHidden
func (v_ View) SetHidden(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHidden:"), value)
}/* debug [instance_properties/setter]: hidden */


// A Boolean value indicating whether the view is hidden from sight because it, or one of its ancestors, is marked as hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isHiddenOrHasHiddenAncestor
func (v_ View) HiddenOrHasHiddenAncestor() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("hiddenOrHasHiddenAncestor"))
	return rv
}/* debug [instance_properties/getter]: hiddenOrHasHiddenAncestor */


// A Boolean value that indicates whether the view’s horizontal size constraints are active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isHorizontalContentSizeConstraintActive
func (v_ View) HorizontalContentSizeConstraintActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("horizontalContentSizeConstraintActive"))
	return rv
}/* debug [instance_properties/getter]: horizontalContentSizeConstraintActive */


// A Boolean value that indicates whether the view’s horizontal size constraints are active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isHorizontalContentSizeConstraintActive
func (v_ View) SetHorizontalContentSizeConstraintActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setHorizontalContentSizeConstraintActive:"), value)
}/* debug [instance_properties/setter]: horizontalContentSizeConstraintActive */


// A Boolean value indicating whether the view is in full screen mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isInFullScreenMode
func (v_ View) InFullScreenMode() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("inFullScreenMode"))
	return rv
}/* debug [instance_properties/getter]: inFullScreenMode */


// A Boolean value indicating whether the view fills its frame rectangle with opaque content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isOpaque
func (v_ View) Opaque() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("opaque"))
	return rv
}/* debug [instance_properties/getter]: opaque */


// A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isRotatedFromBase
func (v_ View) RotatedFromBase() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rotatedFromBase"))
	return rv
}/* debug [instance_properties/getter]: rotatedFromBase */


// A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds, or has been scaled from the window’s base coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isRotatedOrScaledFromBase
func (v_ View) RotatedOrScaledFromBase() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rotatedOrScaledFromBase"))
	return rv
}/* debug [instance_properties/getter]: rotatedOrScaledFromBase */


// A Boolean value that indicates whether the view’s vertical size constraints are active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isVerticalContentSizeConstraintActive
func (v_ View) VerticalContentSizeConstraintActive() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("verticalContentSizeConstraintActive"))
	return rv
}/* debug [instance_properties/getter]: verticalContentSizeConstraintActive */


// A Boolean value that indicates whether the view’s vertical size constraints are active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isVerticalContentSizeConstraintActive
func (v_ View) SetVerticalContentSizeConstraintActive(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setVerticalContentSizeConstraintActive:"), value)
}/* debug [instance_properties/setter]: verticalContentSizeConstraintActive */


// A layout anchor representing the baseline for the bottommost line of text in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/lastBaselineAnchor
func (v_ View) LastBaselineAnchor() ILayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](v_.ID, objc.Sel("lastBaselineAnchor"))
	return rv
}/* debug [instance_properties/getter]: lastBaselineAnchor */


// The distance (in points) between the bottom of the view’s alignment rectangle and its bottommost baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/lastBaselineOffsetFromBottom
func (v_ View) LastBaselineOffsetFromBottom() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("lastBaselineOffsetFromBottom"))
	return rv
}/* debug [instance_properties/getter]: lastBaselineOffsetFromBottom */


// The Core Animation layer that the view uses as its backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layer
func (v_ View) Layer() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](v_.ID, objc.Sel("layer"))
	return rv
}/* debug [instance_properties/getter]: layer */


// The Core Animation layer that the view uses as its backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layer
func (v_ View) SetLayer(value objectivec.IObject) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLayer:"), value)
}/* debug [instance_properties/setter]: layer */


// The current layer contents placement policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerContentsPlacement-swift.property
func (v_ View) LayerContentsPlacement() ViewLayerContentsPlacement {
	rv := objc.Send[ViewLayerContentsPlacement](v_.ID, objc.Sel("layerContentsPlacement"))
	return rv
}/* debug [instance_properties/getter]: layerContentsPlacement */


// The current layer contents placement policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerContentsPlacement-swift.property
func (v_ View) SetLayerContentsPlacement(value ViewLayerContentsPlacement) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLayerContentsPlacement:"), value)
}/* debug [instance_properties/setter]: layerContentsPlacement */


// The contents redraw policy for the view’s layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerContentsRedrawPolicy-swift.property
func (v_ View) LayerContentsRedrawPolicy() ViewLayerContentsRedrawPolicy {
	rv := objc.Send[ViewLayerContentsRedrawPolicy](v_.ID, objc.Sel("layerContentsRedrawPolicy"))
	return rv
}/* debug [instance_properties/getter]: layerContentsRedrawPolicy */


// The contents redraw policy for the view’s layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerContentsRedrawPolicy-swift.property
func (v_ View) SetLayerContentsRedrawPolicy(value ViewLayerContentsRedrawPolicy) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLayerContentsRedrawPolicy:"), value)
}/* debug [instance_properties/setter]: layerContentsRedrawPolicy */


// A Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerUsesCoreImageFilters
func (v_ View) LayerUsesCoreImageFilters() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("layerUsesCoreImageFilters"))
	return rv
}/* debug [instance_properties/getter]: layerUsesCoreImageFilters */


// A Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layerUsesCoreImageFilters
func (v_ View) SetLayerUsesCoreImageFilters(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setLayerUsesCoreImageFilters:"), value)
}/* debug [instance_properties/setter]: layerUsesCoreImageFilters */


// The array of layout guide objects owned by this view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layoutGuides
func (v_ View) LayoutGuides() []LayoutGuide {
	rv := objc.Send[[]LayoutGuide](v_.ID, objc.Sel("layoutGuides"))
	return rv
}/* debug [instance_properties/getter]: layoutGuides */


// A layout guide that provides the recommended amount of padding for content inside of a view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layoutMarginsGuide
func (v_ View) LayoutMarginsGuide() ILayoutGuide {
	rv := objc.Send[LayoutGuide](v_.ID, objc.Sel("layoutMarginsGuide"))
	return rv
}/* debug [instance_properties/getter]: layoutMarginsGuide */


// A layout anchor representing the leading edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/leadingAnchor
func (v_ View) LeadingAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](v_.ID, objc.Sel("leadingAnchor"))
	return rv
}/* debug [instance_properties/getter]: leadingAnchor */


// A layout anchor representing the left edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/leftAnchor
func (v_ View) LeftAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](v_.ID, objc.Sel("leftAnchor"))
	return rv
}/* debug [instance_properties/getter]: leftAnchor */


// A Boolean value indicating whether the view can pass mouse down events through to its superviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/mouseDownCanMoveWindow
func (v_ View) MouseDownCanMoveWindow() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("mouseDownCanMoveWindow"))
	return rv
}/* debug [instance_properties/getter]: mouseDownCanMoveWindow */


// A Boolean value that determines whether the view needs to be redrawn before being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
func (v_ View) NeedsDisplay() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsDisplay"))
	return rv
}/* debug [instance_properties/getter]: needsDisplay */


// A Boolean value that determines whether the view needs to be redrawn before being displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsDisplay
func (v_ View) SetNeedsDisplay(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNeedsDisplay:"), value)
}/* debug [instance_properties/setter]: needsDisplay */


// A Boolean value indicating whether the view needs a layout pass before it can be drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsLayout
func (v_ View) NeedsLayout() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsLayout"))
	return rv
}/* debug [instance_properties/getter]: needsLayout */


// A Boolean value indicating whether the view needs a layout pass before it can be drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsLayout
func (v_ View) SetNeedsLayout(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNeedsLayout:"), value)
}/* debug [instance_properties/setter]: needsLayout */


// A Boolean value indicating whether the view needs its panel to become the key window before it can handle keyboard input and navigation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsPanelToBecomeKey
func (v_ View) NeedsPanelToBecomeKey() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsPanelToBecomeKey"))
	return rv
}/* debug [instance_properties/getter]: needsPanelToBecomeKey */


// A Boolean value indicating whether the view’s constraints need to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsUpdateConstraints
func (v_ View) NeedsUpdateConstraints() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsUpdateConstraints"))
	return rv
}/* debug [instance_properties/getter]: needsUpdateConstraints */


// A Boolean value indicating whether the view’s constraints need to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsUpdateConstraints
func (v_ View) SetNeedsUpdateConstraints(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNeedsUpdateConstraints:"), value)
}/* debug [instance_properties/setter]: needsUpdateConstraints */


// The view object that follows the current view in the key view loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/nextKeyView
func (v_ View) NextKeyView() IView {
	rv := objc.Send[View](v_.ID, objc.Sel("nextKeyView"))
	return rv
}/* debug [instance_properties/getter]: nextKeyView */


// The view object that follows the current view in the key view loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/nextKeyView
func (v_ View) SetNextKeyView(value IView) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNextKeyView:"), value)
}/* debug [instance_properties/setter]: nextKeyView */


// The closest view object in the key view loop that follows the current view in the key view loop and accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/nextValidKeyView
func (v_ View) NextValidKeyView() IView {
	rv := objc.Send[View](v_.ID, objc.Sel("nextValidKeyView"))
	return rv
}/* debug [instance_properties/getter]: nextValidKeyView */


// The view’s closest opaque ancestor, which might be the view itself.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/opaqueAncestor
func (v_ View) OpaqueAncestor() IView {
	rv := objc.Send[View](v_.ID, objc.Sel("opaqueAncestor"))
	return rv
}/* debug [instance_properties/getter]: opaqueAncestor */


// A default footer string that includes the current page number and page count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/pageFooter
func (v_ View) PageFooter() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](v_.ID, objc.Sel("pageFooter"))
	return rv
}/* debug [instance_properties/getter]: pageFooter */


// A default header string that includes the print job title and date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/pageHeader
func (v_ View) PageHeader() foundation.AttributedString {
	rv := objc.Send[foundation.AttributedString](v_.ID, objc.Sel("pageHeader"))
	return rv
}/* debug [instance_properties/getter]: pageHeader */


// A Boolean value indicating whether the view posts notifications when its bounds rectangle changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/postsBoundsChangedNotifications
func (v_ View) PostsBoundsChangedNotifications() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("postsBoundsChangedNotifications"))
	return rv
}/* debug [instance_properties/getter]: postsBoundsChangedNotifications */


// A Boolean value indicating whether the view posts notifications when its bounds rectangle changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/postsBoundsChangedNotifications
func (v_ View) SetPostsBoundsChangedNotifications(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPostsBoundsChangedNotifications:"), value)
}/* debug [instance_properties/setter]: postsBoundsChangedNotifications */


// A Boolean value indicating whether the view posts notifications when its frame rectangle changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/postsFrameChangedNotifications
func (v_ View) PostsFrameChangedNotifications() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("postsFrameChangedNotifications"))
	return rv
}/* debug [instance_properties/getter]: postsFrameChangedNotifications */


// A Boolean value indicating whether the view posts notifications when its frame rectangle changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/postsFrameChangedNotifications
func (v_ View) SetPostsFrameChangedNotifications(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPostsFrameChangedNotifications:"), value)
}/* debug [instance_properties/setter]: postsFrameChangedNotifications */


// When this property is true, any NSControls in the view or its descendants will be sized with compact metrics compatible with macOS 15 and earlier. Defaults to false
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/prefersCompactControlSizeMetrics
func (v_ View) PrefersCompactControlSizeMetrics() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("prefersCompactControlSizeMetrics"))
	return rv
}/* debug [instance_properties/getter]: prefersCompactControlSizeMetrics */


// When this property is true, any NSControls in the view or its descendants will be sized with compact metrics compatible with macOS 15 and earlier. Defaults to false
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/prefersCompactControlSizeMetrics
func (v_ View) SetPrefersCompactControlSizeMetrics(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPrefersCompactControlSizeMetrics:"), value)
}/* debug [instance_properties/setter]: prefersCompactControlSizeMetrics */


// The portion of the view that has been rendered and is available for responsive scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/preparedContentRect
func (v_ View) PreparedContentRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("preparedContentRect"))
	return rv
}/* debug [instance_properties/getter]: preparedContentRect */


// The portion of the view that has been rendered and is available for responsive scrolling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/preparedContentRect
func (v_ View) SetPreparedContentRect(value Rect /* not a class type */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreparedContentRect:"), value)
}/* debug [instance_properties/setter]: preparedContentRect */


// A Boolean value indicating whether the view optimizes live-resize operations by preserving content that has not moved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/preservesContentDuringLiveResize
func (v_ View) PreservesContentDuringLiveResize() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("preservesContentDuringLiveResize"))
	return rv
}/* debug [instance_properties/getter]: preservesContentDuringLiveResize */


// Configures the behavior and progression of the Force Touch trackpad when responding to touch input produced by the user when the cursor is over the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/pressureConfiguration
func (v_ View) PressureConfiguration() IPressureConfiguration {
	rv := objc.Send[PressureConfiguration](v_.ID, objc.Sel("pressureConfiguration"))
	return rv
}/* debug [instance_properties/getter]: pressureConfiguration */


// Configures the behavior and progression of the Force Touch trackpad when responding to touch input produced by the user when the cursor is over the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/pressureConfiguration
func (v_ View) SetPressureConfiguration(value IPressureConfiguration) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPressureConfiguration:"), value)
}/* debug [instance_properties/setter]: pressureConfiguration */


// The view object preceding the current view in the key view loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/previousKeyView
func (v_ View) PreviousKeyView() IView {
	rv := objc.Send[View](v_.ID, objc.Sel("previousKeyView"))
	return rv
}/* debug [instance_properties/getter]: previousKeyView */


// The closest view object in the key view loop that precedes the current view and accepts first responder status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/previousValidKeyView
func (v_ View) PreviousValidKeyView() IView {
	rv := objc.Send[View](v_.ID, objc.Sel("previousValidKeyView"))
	return rv
}/* debug [instance_properties/getter]: previousValidKeyView */


// The view’s print job title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/printJobTitle
func (v_ View) PrintJobTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("printJobTitle"))
	return rv
}/* debug [instance_properties/getter]: printJobTitle */


// The rectangle identifying the portion of your view that did not change during a live resize operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectPreservedDuringLiveResize
func (v_ View) RectPreservedDuringLiveResize() Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("rectPreservedDuringLiveResize"))
	return rv
}/* debug [instance_properties/getter]: rectPreservedDuringLiveResize */


// The array of pasteboard drag types that the view can accept.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/registeredDraggedTypes
func (v_ View) RegisteredDraggedTypes() []string {
	rv := objc.Send[[]string](v_.ID, objc.Sel("registeredDraggedTypes"))
	return rv
}/* debug [instance_properties/getter]: registeredDraggedTypes */


// Returns a Boolean value indicating whether the view depends on the constraint-based layout system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/requiresConstraintBasedLayout
func (v_ View) RequiresConstraintBasedLayout() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("requiresConstraintBasedLayout"))
	return rv
}/* debug [instance_properties/getter]: requiresConstraintBasedLayout */


// A layout anchor representing the right edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rightAnchor
func (v_ View) RightAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](v_.ID, objc.Sel("rightAnchor"))
	return rv
}/* debug [instance_properties/getter]: rightAnchor */


// The distances from the edges of your view that define the safe area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/safeAreaInsets
func (v_ View) SafeAreaInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](v_.ID, objc.Sel("safeAreaInsets"))
	return rv
}/* debug [instance_properties/getter]: safeAreaInsets */


// The layout guide you use to position content inside your view’s safe area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/safeAreaLayoutGuide
func (v_ View) SafeAreaLayoutGuide() ILayoutGuide {
	rv := objc.Send[LayoutGuide](v_.ID, objc.Sel("safeAreaLayoutGuide"))
	return rv
}/* debug [instance_properties/getter]: safeAreaLayoutGuide */


// A rectangle in the view’s coordinate system that contains the unobscured portion of the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/safeAreaRect
func (v_ View) SafeAreaRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("safeAreaRect"))
	return rv
}/* debug [instance_properties/getter]: safeAreaRect */


// The shadow displayed underneath the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/shadow
func (v_ View) Shadow() IShadow {
	rv := objc.Send[Shadow](v_.ID, objc.Sel("shadow"))
	return rv
}/* debug [instance_properties/getter]: shadow */


// The shadow displayed underneath the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/shadow
func (v_ View) SetShadow(value IShadow) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShadow:"), value)
}/* debug [instance_properties/setter]: shadow */


// The array of views embedded in the current view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/subviews
func (v_ View) Subviews() []View {
	rv := objc.Send[[]View](v_.ID, objc.Sel("subviews"))
	return rv
}/* debug [instance_properties/getter]: subviews */


// The array of views embedded in the current view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/subviews
func (v_ View) SetSubviews(value []View) {
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
}/* debug [instance_properties/setter]: subviews */


// The view that is the parent of the current view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/superview
func (v_ View) Superview() IView {
	rv := objc.Send[View](v_.ID, objc.Sel("superview"))
	return rv
}/* debug [instance_properties/getter]: superview */


// The view’s tag, which is an integer that you use to identify the view within your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/tag
func (v_ View) Tag() int {
	rv := objc.Send[int](v_.ID, objc.Sel("tag"))
	return rv
}/* debug [instance_properties/getter]: tag */


// The text for the view’s tooltip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/toolTip
func (v_ View) ToolTip() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](v_.ID, objc.Sel("toolTip"))
	return rv
}/* debug [instance_properties/getter]: toolTip */


// The text for the view’s tooltip.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/toolTip
func (v_ View) SetToolTip(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setToolTip:"), value)
}/* debug [instance_properties/setter]: toolTip */


// A layout anchor representing the top edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/topAnchor
func (v_ View) TopAnchor() ILayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](v_.ID, objc.Sel("topAnchor"))
	return rv
}/* debug [instance_properties/getter]: topAnchor */


// An array of the view’s tracking areas.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/trackingAreas
func (v_ View) TrackingAreas() []TrackingArea {
	rv := objc.Send[[]TrackingArea](v_.ID, objc.Sel("trackingAreas"))
	return rv
}/* debug [instance_properties/getter]: trackingAreas */


// A layout anchor representing the trailing edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/trailingAnchor
func (v_ View) TrailingAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](v_.ID, objc.Sel("trailingAnchor"))
	return rv
}/* debug [instance_properties/getter]: trailingAnchor */


// A Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/translatesAutoresizingMaskIntoConstraints
func (v_ View) TranslatesAutoresizingMaskIntoConstraints() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("translatesAutoresizingMaskIntoConstraints"))
	return rv
}/* debug [instance_properties/getter]: translatesAutoresizingMaskIntoConstraints */


// A Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/translatesAutoresizingMaskIntoConstraints
func (v_ View) SetTranslatesAutoresizingMaskIntoConstraints(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setTranslatesAutoresizingMaskIntoConstraints:"), value)
}/* debug [instance_properties/setter]: translatesAutoresizingMaskIntoConstraints */


// The layout direction for content in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/userInterfaceLayoutDirection
func (v_ View) UserInterfaceLayoutDirection() UserInterfaceLayoutDirection {
	rv := objc.Send[UserInterfaceLayoutDirection](v_.ID, objc.Sel("userInterfaceLayoutDirection"))
	return rv
}/* debug [instance_properties/getter]: userInterfaceLayoutDirection */


// The layout direction for content in the view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/userInterfaceLayoutDirection
func (v_ View) SetUserInterfaceLayoutDirection(value UserInterfaceLayoutDirection) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}/* debug [instance_properties/setter]: userInterfaceLayoutDirection */


// The portion of the view that isn’t clipped by its superviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/visibleRect
func (v_ View) VisibleRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](v_.ID, objc.Sel("visibleRect"))
	return rv
}/* debug [instance_properties/getter]: visibleRect */


// A Boolean value indicating whether the view wants an OpenGL backing surface with a resolution greater than 1 pixel per point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsBestResolutionOpenGLSurface
func (v_ View) WantsBestResolutionOpenGLSurface() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsBestResolutionOpenGLSurface"))
	return rv
}/* debug [instance_properties/getter]: wantsBestResolutionOpenGLSurface */


// A Boolean value indicating whether the view wants an OpenGL backing surface with a resolution greater than 1 pixel per point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsBestResolutionOpenGLSurface
func (v_ View) SetWantsBestResolutionOpenGLSurface(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWantsBestResolutionOpenGLSurface:"), value)
}/* debug [instance_properties/setter]: wantsBestResolutionOpenGLSurface */


// A Boolean value indicating whether AppKit’s default clipping behavior is in effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsDefaultClipping
func (v_ View) WantsDefaultClipping() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsDefaultClipping"))
	return rv
}/* debug [instance_properties/getter]: wantsDefaultClipping */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsExtendedDynamicRangeOpenGLSurface
func (v_ View) WantsExtendedDynamicRangeOpenGLSurface() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsExtendedDynamicRangeOpenGLSurface"))
	return rv
}/* debug [instance_properties/getter]: wantsExtendedDynamicRangeOpenGLSurface */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsExtendedDynamicRangeOpenGLSurface
func (v_ View) SetWantsExtendedDynamicRangeOpenGLSurface(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWantsExtendedDynamicRangeOpenGLSurface:"), value)
}/* debug [instance_properties/setter]: wantsExtendedDynamicRangeOpenGLSurface */


// A Boolean value indicating whether the view uses a layer as its backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsLayer
func (v_ View) WantsLayer() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsLayer"))
	return rv
}/* debug [instance_properties/getter]: wantsLayer */


// A Boolean value indicating whether the view uses a layer as its backing store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsLayer
func (v_ View) SetWantsLayer(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWantsLayer:"), value)
}/* debug [instance_properties/setter]: wantsLayer */


// A Boolean value indicating whether the view wants resting touches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsRestingTouches
func (v_ View) WantsRestingTouches() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsRestingTouches"))
	return rv
}/* debug [instance_properties/getter]: wantsRestingTouches */


// A Boolean value indicating whether the view wants resting touches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsRestingTouches
func (v_ View) SetWantsRestingTouches(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWantsRestingTouches:"), value)
}/* debug [instance_properties/setter]: wantsRestingTouches */


// A Boolean value indicating which drawing path the view takes when updating its contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/wantsUpdateLayer
func (v_ View) WantsUpdateLayer() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("wantsUpdateLayer"))
	return rv
}/* debug [instance_properties/getter]: wantsUpdateLayer */


// The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as small images or text columns from being divided across pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/widthAdjustLimit
func (v_ View) WidthAdjustLimit() float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("widthAdjustLimit"))
	return rv
}/* debug [instance_properties/getter]: widthAdjustLimit */


// A layout anchor representing the width of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/widthAnchor
func (v_ View) WidthAnchor() ILayoutDimension {
	rv := objc.Send[LayoutDimension](v_.ID, objc.Sel("widthAnchor"))
	return rv
}/* debug [instance_properties/getter]: widthAnchor */


// The view’s window object, if it is installed in a window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/window
func (v_ View) Window() IWindow {
	rv := objc.Send[Window](v_.ID, objc.Sel("window"))
	return rv
}/* debug [instance_properties/getter]: window */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/writingToolsCoordinator
func (v_ View) WritingToolsCoordinator() IWritingToolsCoordinator {
	rv := objc.Send[WritingToolsCoordinator](v_.ID, objc.Sel("writingToolsCoordinator"))
	return rv
}/* debug [instance_properties/getter]: writingToolsCoordinator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/writingToolsCoordinator
func (v_ View) SetWritingToolsCoordinator(value IWritingToolsCoordinator) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setWritingToolsCoordinator:"), value)
}/* debug [instance_properties/setter]: writingToolsCoordinator */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSView */


