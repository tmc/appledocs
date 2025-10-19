// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [View] class.
var (
	viewClass     _ViewClass
	viewClassOnce sync.Once
)

func getViewClass() _ViewClass {
	viewClassOnce.Do(func() {
		viewClass = _ViewClass{objc.GetClass("NSView")}
	})
	return viewClass
}

type _ViewClass struct {
	class objc.Class
}

// An interface definition for the [View] class.
type IView interface {
	IResponder
	AcceptsFirstMouse(event unsafe.Pointer) bool
	AddConstraint(constraint unsafe.Pointer)
	AddConstraints(constraints unsafe.Pointer)
	AddCursorRectCursor(rect unsafe.Pointer, object unsafe.Pointer)
	AddGestureRecognizer(gestureRecognizer unsafe.Pointer)
	AddLayoutGuide(guide unsafe.Pointer)
	AddSubview(view unsafe.Pointer)
	AddSubviewPositionedRelativeTo(view unsafe.Pointer, place WindowOrderingMode, otherView unsafe.Pointer)
	AddToolTipRectOwnerUserData(rect unsafe.Pointer, owner objc.ID, data unsafe.Pointer) unsafe.Pointer
	AddTrackingArea(trackingArea unsafe.Pointer)
	AddTrackingRectOwnerUserDataAssumeInside(rect unsafe.Pointer, owner objc.ID, data unsafe.Pointer, flag bool) unsafe.Pointer
	AdjustPageHeightNewTopBottomLimit(newBottom float64, oldTop float64, oldBottom float64, bottomLimit float64)
	AdjustPageWidthNewLeftRightLimit(newRight float64, oldLeft float64, oldRight float64, rightLimit float64)
	AdjustScroll(newVisible unsafe.Pointer) unsafe.Pointer
	AlignmentRectForFrame(frame unsafe.Pointer) unsafe.Pointer
	AllocateGState()
	AncestorSharedWithView(view unsafe.Pointer) unsafe.Pointer
	Autoscroll(event unsafe.Pointer) bool
	BackingAlignedRectOptions(rect unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer
	BeginDocument()
	BeginDraggingSessionWithItemsEventSource(items unsafe.Pointer, event unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer
	BeginPageInRectAtPlacement(rect unsafe.Pointer, location unsafe.Pointer)
	BitmapImageRepForCachingDisplayInRect(rect unsafe.Pointer) unsafe.Pointer
	CacheDisplayInRectToBitmapImageRep(rect unsafe.Pointer, bitmapImageRep unsafe.Pointer)
	CenterScanRect(rect unsafe.Pointer) unsafe.Pointer
	ConstraintsAffectingLayoutForOrientation(orientation unsafe.Pointer) unsafe.Pointer
	ContentCompressionResistancePriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer
	ContentHuggingPriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer
	ConvertPointFromView(point unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertSizeFromView(size unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertRectFromView(rect unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertRectToView(rect unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertSizeToView(size unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertPointToView(point unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertPointFromBacking(point unsafe.Pointer) unsafe.Pointer
	ConvertRectFromBacking(rect unsafe.Pointer) unsafe.Pointer
	ConvertSizeFromBacking(size unsafe.Pointer) unsafe.Pointer
	ConvertPointFromLayer(point unsafe.Pointer) unsafe.Pointer
	ConvertSizeFromLayer(size unsafe.Pointer) unsafe.Pointer
	ConvertRectFromLayer(rect unsafe.Pointer) unsafe.Pointer
	ConvertPointFromBase(point unsafe.Pointer) unsafe.Pointer
	ConvertPointToBase(point unsafe.Pointer) unsafe.Pointer
	ConvertRectFromBase(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectToBase(rect unsafe.Pointer) unsafe.Pointer
	ConvertSizeFromBase(size unsafe.Pointer) unsafe.Pointer
	ConvertSizeToBase(size unsafe.Pointer) unsafe.Pointer
	ConvertPointToBacking(point unsafe.Pointer) unsafe.Pointer
	ConvertRectToBacking(rect unsafe.Pointer) unsafe.Pointer
	ConvertSizeToBacking(size unsafe.Pointer) unsafe.Pointer
	ConvertRectToLayer(rect unsafe.Pointer) unsafe.Pointer
	ConvertSizeToLayer(size unsafe.Pointer) unsafe.Pointer
	ConvertPointToLayer(point unsafe.Pointer) unsafe.Pointer
	DataWithEPSInsideRect(rect unsafe.Pointer) unsafe.Pointer
	DataWithPDFInsideRect(rect unsafe.Pointer) unsafe.Pointer
	DidAddSubview(subview unsafe.Pointer)
	DidCloseMenuWithEvent(menu unsafe.Pointer, event unsafe.Pointer)
	DiscardCursorRects()
	Display()
	DisplayRect(rect unsafe.Pointer)
	DisplayIfNeeded()
	DisplayIfNeededInRect(rect unsafe.Pointer)
	DisplayIfNeededIgnoringOpacity()
	DisplayIfNeededInRectIgnoringOpacity(rect unsafe.Pointer)
	DisplayRectIgnoringOpacity(rect unsafe.Pointer)
	DisplayRectIgnoringOpacityInContext(rect unsafe.Pointer, context unsafe.Pointer)
	DisplayLinkWithTargetSelector(target objc.ID, selector objc.SEL) unsafe.Pointer
	DragFileFromRectSlideBackEvent(filename string, rect unsafe.Pointer, flag bool, event unsafe.Pointer) bool
	DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, viewLocation unsafe.Pointer, initialOffset unsafe.Pointer, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool)
	DragPromisedFilesOfTypesFromRectSourceSlideBackEvent(typeArray unsafe.Pointer, rect unsafe.Pointer, sourceObject objc.ID, flag bool, event unsafe.Pointer) bool
	DrawRect(dirtyRect unsafe.Pointer)
	DrawFocusRingMask()
	DrawPageBorderWithSize(borderSize unsafe.Pointer)
	DrawSheetBorderWithSize(borderSize unsafe.Pointer)
	EdgeInsetsForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer
	EndDocument()
	EndPage()
	EnterFullScreenModeWithOptions(screen unsafe.Pointer, options unsafe.Pointer) bool
	ExerciseAmbiguityInLayout()
	ExitFullScreenModeWithOptions(options unsafe.Pointer)
	FrameForAlignmentRect(alignmentRect unsafe.Pointer) unsafe.Pointer
	GState() int
	GetRectsBeingDrawnCount(rects unsafe.Pointer, count unsafe.Pointer)
	GetRectsExposedDuringLiveResizeCount(exposedRects unsafe.Pointer, count unsafe.Pointer)
	HitTest(point unsafe.Pointer) unsafe.Pointer
	InvalidateIntrinsicContentSize()
	IsDescendantOf(view unsafe.Pointer) bool
	MouseInRect(point unsafe.Pointer, rect unsafe.Pointer) bool
	KnowsPageRange(range_ unsafe.Pointer) bool
	Layout()
	LayoutGuideForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer
	LayoutSubtreeIfNeeded()
	LocationOfPrintRect(rect unsafe.Pointer) unsafe.Pointer
	LockFocus()
	LockFocusIfCanDraw() bool
	LockFocusIfCanDrawInContext(context unsafe.Pointer) bool
	MakeBackingLayer() unsafe.Pointer
	MenuForEvent(event unsafe.Pointer) unsafe.Pointer
	NeedsToDrawRect(rect unsafe.Pointer) bool
	NoteFocusRingMaskChanged()
	PerformKeyEquivalent(event unsafe.Pointer) bool
	PerformMnemonic(string string) bool
	PrepareContentInRect(rect unsafe.Pointer)
	PrepareForReuse()
	Print(sender objc.ID)
	RectForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer
	RectForPage(page int) unsafe.Pointer
	RectForSmartMagnificationAtPointInRect(location unsafe.Pointer, visibleRect unsafe.Pointer) unsafe.Pointer
	ReflectScrolledClipView(clipView unsafe.Pointer)
	RegisterForDraggedTypes(newTypes unsafe.Pointer)
	ReleaseGState()
	RemoveAllToolTips()
	RemoveConstraint(constraint unsafe.Pointer)
	RemoveConstraints(constraints unsafe.Pointer)
	RemoveCursorRectCursor(rect unsafe.Pointer, object unsafe.Pointer)
	RemoveFromSuperview()
	RemoveFromSuperviewWithoutNeedingDisplay()
	RemoveGestureRecognizer(gestureRecognizer unsafe.Pointer)
	RemoveLayoutGuide(guide unsafe.Pointer)
	RemoveToolTip(tag unsafe.Pointer)
	RemoveTrackingArea(trackingArea unsafe.Pointer)
	RemoveTrackingRect(tag unsafe.Pointer)
	RenewGState()
	ReplaceSubviewWith(oldView unsafe.Pointer, newView unsafe.Pointer)
	ResetCursorRects()
	ResizeWithOldSuperviewSize(oldSize unsafe.Pointer)
	ResizeSubviewsWithOldSize(oldSize unsafe.Pointer)
	RotateByAngle(angle float64)
	RulerViewDidAddMarker(ruler unsafe.Pointer, marker unsafe.Pointer)
	RulerViewDidMoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer)
	RulerViewDidRemoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer)
	RulerViewHandleMouseDown(ruler unsafe.Pointer, event unsafe.Pointer)
	RulerViewLocationForPoint(ruler unsafe.Pointer, point unsafe.Pointer) float64
	RulerViewPointForLocation(ruler unsafe.Pointer, point float64) unsafe.Pointer
	RulerViewShouldAddMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool
	RulerViewShouldMoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool
	RulerViewShouldRemoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool
	RulerViewWillAddMarkerAtLocation(ruler unsafe.Pointer, marker unsafe.Pointer, location float64) float64
	RulerViewWillMoveMarkerToLocation(ruler unsafe.Pointer, marker unsafe.Pointer, location float64) float64
	RulerViewWillSetClientView(ruler unsafe.Pointer, newClient unsafe.Pointer)
	ScaleUnitSquareToSize(newUnitSize unsafe.Pointer)
	ScrollPoint(point unsafe.Pointer)
	ScrollRectBy(rect unsafe.Pointer, delta unsafe.Pointer)
	ScrollClipViewToPoint(clipView unsafe.Pointer, point unsafe.Pointer)
	ScrollRectToVisible(rect unsafe.Pointer) bool
	SetBoundsOrigin(newOrigin unsafe.Pointer)
	SetBoundsSize(newSize unsafe.Pointer)
	SetContentCompressionResistancePriorityForOrientation(priority unsafe.Pointer, orientation unsafe.Pointer)
	SetContentHuggingPriorityForOrientation(priority unsafe.Pointer, orientation unsafe.Pointer)
	SetFrameOrigin(newOrigin unsafe.Pointer)
	SetFrameSize(newSize unsafe.Pointer)
	SetKeyboardFocusRingNeedsDisplayInRect(rect unsafe.Pointer)
	SetNeedsDisplayInRect(invalidRect unsafe.Pointer)
	SetUpGState()
	ShouldDelayWindowOrderingForEvent(event unsafe.Pointer) bool
	ShouldDrawColor() bool
	ShowDefinitionForAttributedStringAtPoint(attrString unsafe.Pointer, textBaselineOrigin unsafe.Pointer)
	ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString unsafe.Pointer, targetRange unsafe.Pointer, options unsafe.Pointer, originProvider unsafe.Pointer)
	SortSubviewsUsingFunctionContext(compare unsafe.Pointer, context unsafe.Pointer)
	TranslateOriginToPoint(translation unsafe.Pointer)
	TranslateRectsNeedingDisplayInRectBy(clipRect unsafe.Pointer, delta unsafe.Pointer)
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
	ViewWillMoveToSuperview(newSuperview unsafe.Pointer)
	ViewWillMoveToWindow(newWindow unsafe.Pointer)
	ViewWillStartLiveResize()
	ViewWithTag(tag int) unsafe.Pointer
	WillOpenMenuWithEvent(menu unsafe.Pointer, event unsafe.Pointer)
	WillRemoveSubview(subview unsafe.Pointer)
	WriteEPSInsideRectToPasteboard(rect unsafe.Pointer, pasteboard unsafe.Pointer)
	WritePDFInsideRectToPasteboard(rect unsafe.Pointer, pasteboard unsafe.Pointer)
}

// The infrastructure for drawing, printing, and handling events in an app. [Full Topic]
//
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


// Initializes a view using from data in the specified coder object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewViewWithCoder(coder unsafe.Pointer) View {
	// Instance methods (init*) require Autorelease() to balance the +1 from alloc
	instance := getViewClass().Alloc()
	rv := objc.Send[View](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}
// Initializes and returns a newly allocated object with a specified frame rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewViewWithFrame(frameRect unsafe.Pointer) View {
	// Instance methods (init*) require Autorelease() to balance the +1 from alloc
	instance := getViewClass().Alloc()
	rv := objc.Send[View](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}


// Overridden by subclasses to return if the view should be sent a message for an initial mouse-down event, if not. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/acceptsFirstMouse(for:)
func (v_ View) AcceptsFirstMouse(event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("acceptsFirstMouse:"), event)
	return rv
}
// Adds a constraint on the layout of the receiving view or its subviews. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addConstraint(_:)
func (v_ View) AddConstraint(constraint unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addConstraint:"), constraint)
}
// Adds multiple constraints on the layout of the receiving view or its subviews. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addConstraints(_:)
func (v_ View) AddConstraints(constraints unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addConstraints:"), constraints)
}
// Establishes the cursor to be used when the mouse pointer lies within a specified region. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addCursorRect(_:cursor:)
func (v_ View) AddCursorRectCursor(rect unsafe.Pointer, object unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addCursorRect:cursor:"), rect, object)
}
// Attaches a gesture recognizer to the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addGestureRecognizer(_:)
func (v_ View) AddGestureRecognizer(gestureRecognizer unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addGestureRecognizer:"), gestureRecognizer)
}
// Adds the provided layout guide to the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addLayoutGuide(_:)
func (v_ View) AddLayoutGuide(guide unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addLayoutGuide:"), guide)
}
// Adds a view to the view’s subviews so it’s displayed above its siblings. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addSubview(_:)
func (v_ View) AddSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addSubview:"), view)
}
// Inserts a view among the view’s subviews so it’s displayed immediately above or below another view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addSubview(_:positioned:relativeTo:)
func (v_ View) AddSubviewPositionedRelativeTo(view unsafe.Pointer, place WindowOrderingMode, otherView unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addSubview:positioned:relativeTo:"), view, place, otherView)
}
// Creates a tooltip for a defined area in the view and returns a tag that identifies the tooltip rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addToolTip(_:owner:userData:)
func (v_ View) AddToolTipRectOwnerUserData(rect unsafe.Pointer, owner objc.ID, data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("addToolTipRect:owner:userData:"), rect, owner, data)
	return rv
}
// Adds a given tracking area to the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addTrackingArea(_:)
func (v_ View) AddTrackingArea(trackingArea unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("addTrackingArea:"), trackingArea)
}
// Establishes an area for tracking mouse-entered and mouse-exited events within the view and returns a tag that identifies the tracking rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/addTrackingRect(_:owner:userData:assumeInside:)
func (v_ View) AddTrackingRectOwnerUserDataAssumeInside(rect unsafe.Pointer, owner objc.ID, data unsafe.Pointer, flag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("addTrackingRect:owner:userData:assumeInside:"), rect, owner, data, flag)
	return rv
}
// Overridden by subclasses to adjust page height during automatic pagination. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/adjustPageHeightNew(_:top:bottom:limit:)
func (v_ View) AdjustPageHeightNewTopBottomLimit(newBottom float64, oldTop float64, oldBottom float64, bottomLimit float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("adjustPageHeightNew:top:bottom:limit:"), newBottom, oldTop, oldBottom, bottomLimit)
}
// Overridden by subclasses to adjust page width during automatic pagination. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/adjustPageWidthNew(_:left:right:limit:)
func (v_ View) AdjustPageWidthNewLeftRightLimit(newRight float64, oldLeft float64, oldRight float64, rightLimit float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("adjustPageWidthNew:left:right:limit:"), newRight, oldLeft, oldRight, rightLimit)
}
// Overridden by subclasses to modify a given rectangle, returning the altered rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/adjustScroll(_:)
func (v_ View) AdjustScroll(newVisible unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("adjustScroll:"), newVisible)
	return rv
}
// Returns the view’s alignment rectangle for a given frame. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/alignmentRect(forFrame:)
func (v_ View) AlignmentRectForFrame(frame unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("alignmentRectForFrame:"), frame)
	return rv
}
// Causes the view to maintain a private graphics state object, which encapsulates all parameters of the graphics environment. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/allocateGState()
func (v_ View) AllocateGState() {
	objc.Send[objc.ID](v_.ID, objc.Sel("allocateGState"))
}
// Returns the closest ancestor shared by the view and another specified view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/ancestorShared(with:)
func (v_ View) AncestorSharedWithView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("ancestorSharedWithView:"), view)
	return rv
}
// Scrolls the view’s closest ancestor object proportionally to the distance of an event that occurs outside of it. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/autoscroll(with:)
func (v_ View) Autoscroll(event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("autoscroll:"), event)
	return rv
}
// Returns a backing store pixel-aligned rectangle in local view coordinates. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/backingAlignedRect(_:options:)
func (v_ View) BackingAlignedRectOptions(rect unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
	return rv
}
// Invoked at the beginning of the printing session, this method sets up the current graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/beginDocument()
func (v_ View) BeginDocument() {
	objc.Send[objc.ID](v_.ID, objc.Sel("beginDocument"))
}
// Initiates a dragging session with a group of dragging items. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/beginDraggingSession(with:event:source:)
func (v_ View) BeginDraggingSessionWithItemsEventSource(items unsafe.Pointer, event unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("beginDraggingSessionWithItems:event:source:"), items, event, source)
	return rv
}
// Called at the beginning of each page, this method sets up the coordinate system so that a region inside the view’s bounds is translated to a specified location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/beginPage(in:atPlacement:)
func (v_ View) BeginPageInRectAtPlacement(rect unsafe.Pointer, location unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("beginPageInRect:atPlacement:"), rect, location)
}
// Returns a bitmap-representation object suitable for caching the specified portion of the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/bitmapImageRepForCachingDisplay(in:)
func (v_ View) BitmapImageRepForCachingDisplayInRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("bitmapImageRepForCachingDisplayInRect:"), rect)
	return rv
}
// Draws the specified area of the view, and its descendants, into a provided bitmap-representation object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/cacheDisplay(in:to:)
func (v_ View) CacheDisplayInRectToBitmapImageRep(rect unsafe.Pointer, bitmapImageRep unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("cacheDisplayInRect:toBitmapImageRep:"), rect, bitmapImageRep)
}
// Converts the corners of a specified rectangle to lie on the center of device pixels, which is useful in compensating for rendering overscanning when the coordinate system has been scaled. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/centerScanRect(_:)
func (v_ View) CenterScanRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("centerScanRect:"), rect)
	return rv
}
// Returns the constraints impacting the layout of the view for a given orientation. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/constraintsAffectingLayout(for:)
func (v_ View) ConstraintsAffectingLayoutForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("constraintsAffectingLayoutForOrientation:"), orientation)
	return rv
}
// Returns the priority with which a view resists being made smaller than its intrinsic size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/contentCompressionResistancePriority(for:)
func (v_ View) ContentCompressionResistancePriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("contentCompressionResistancePriorityForOrientation:"), orientation)
	return rv
}
// Returns the priority with which a view resists being made larger than its intrinsic size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/contentHuggingPriority(for:)
func (v_ View) ContentHuggingPriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("contentHuggingPriorityForOrientation:"), orientation)
	return rv
}
// Converts a point from the coordinate system of a given view to that of the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-1dq9l
func (v_ View) ConvertPointFromView(point unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertPoint:fromView:"), point, view)
	return rv
}
// Converts a size from another view’s coordinate system to that of the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-40x0w
func (v_ View) ConvertSizeFromView(size unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertSize:fromView:"), size, view)
	return rv
}
// Converts a rectangle from the coordinate system of another view to that of the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:from:)-7fbb6
func (v_ View) ConvertRectFromView(rect unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertRect:fromView:"), rect, view)
	return rv
}
// Converts a rectangle from the view’s coordinate system to that of another view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-3cqqt
func (v_ View) ConvertRectToView(rect unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertRect:toView:"), rect, view)
	return rv
}
// Converts a size from the view’s coordinate system to that of another view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-5nptx
func (v_ View) ConvertSizeToView(size unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertSize:toView:"), size, view)
	return rv
}
// Converts a point from the view’s coordinate system to that of a given view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convert(_:to:)-6u9ir
func (v_ View) ConvertPointToView(point unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertPoint:toView:"), point, view)
	return rv
}
// Converts a point from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-229ps
func (v_ View) ConvertPointFromBacking(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertPointFromBacking:"), point)
	return rv
}
// Converts a rectangle from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-2njpa
func (v_ View) ConvertRectFromBacking(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertRectFromBacking:"), rect)
	return rv
}
// Converts a size from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromBacking(_:)-4agf9
func (v_ View) ConvertSizeFromBacking(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertSizeFromBacking:"), size)
	return rv
}
// Convert the point from the layer’s interior coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-3nsbu
func (v_ View) ConvertPointFromLayer(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertPointFromLayer:"), point)
	return rv
}
// Convert the size from the layer’s interior coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-3usqp
func (v_ View) ConvertSizeFromLayer(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertSizeFromLayer:"), size)
	return rv
}
// Convert the rectangle from the layer’s interior coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertFromLayer(_:)-8s5bi
func (v_ View) ConvertRectFromLayer(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertRectFromLayer:"), rect)
	return rv
}
// Converts the point from the base coordinate system to the view’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertPointFromBase:
func (v_ View) ConvertPointFromBase(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertPointFromBase:"), point)
	return rv
}
// Converts the point from the view’s coordinate system to the base coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertPointToBase:
func (v_ View) ConvertPointToBase(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertPointToBase:"), point)
	return rv
}
// Converts the rectangle from the base coordinate system to the view’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertRectFromBase:
func (v_ View) ConvertRectFromBase(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertRectFromBase:"), rect)
	return rv
}
// Converts the rectangle from the view’s coordinate system to the base coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertRectToBase:
func (v_ View) ConvertRectToBase(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertRectToBase:"), rect)
	return rv
}
// Converts the size from the base coordinate system to the view’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertSizeFromBase:
func (v_ View) ConvertSizeFromBase(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertSizeFromBase:"), size)
	return rv
}
// Converts the size from the view’s coordinate system to the base coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertSizeToBase:
func (v_ View) ConvertSizeToBase(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertSizeToBase:"), size)
	return rv
}
// Converts a point from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-2xx45
func (v_ View) ConvertPointToBacking(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertPointToBacking:"), point)
	return rv
}
// Converts a rectangle from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-3zors
func (v_ View) ConvertRectToBacking(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertRectToBacking:"), rect)
	return rv
}
// Converts a size from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToBacking(_:)-4ra9y
func (v_ View) ConvertSizeToBacking(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertSizeToBacking:"), size)
	return rv
}
// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-160pw
func (v_ View) ConvertRectToLayer(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertRectToLayer:"), rect)
	return rv
}
// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-2vozx
func (v_ View) ConvertSizeToLayer(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertSizeToLayer:"), size)
	return rv
}
// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/convertToLayer(_:)-44u7d
func (v_ View) ConvertPointToLayer(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("convertPointToLayer:"), point)
	return rv
}
// Returns EPS data that draws the region of the view within a specified rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dataWithEPS(inside:)
func (v_ View) DataWithEPSInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("dataWithEPSInsideRect:"), rect)
	return rv
}
// Returns PDF data that draws the region of the view within a specified rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dataWithPDF(inside:)
func (v_ View) DataWithPDFInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("dataWithPDFInsideRect:"), rect)
	return rv
}
// Overridden by subclasses to perform additional actions when subviews are added to the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/didAddSubview(_:)
func (v_ View) DidAddSubview(subview unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("didAddSubview:"), subview)
}
// Called after a contextual menu that was displayed from the receiving view has been closed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/didCloseMenu(_:with:)
func (v_ View) DidCloseMenuWithEvent(menu unsafe.Pointer, event unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("didCloseMenu:withEvent:"), menu, event)
}
// Invalidates all cursor rectangles set up using . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/discardCursorRects()
func (v_ View) DiscardCursorRects() {
	objc.Send[objc.ID](v_.ID, objc.Sel("discardCursorRects"))
}
// Displays the view and all its subviews if possible, invoking each of the methods , , and as necessary. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/display()
func (v_ View) Display() {
	objc.Send[objc.ID](v_.ID, objc.Sel("display"))
}
// Acts as , but confining drawing to a rectangular region of the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/display(_:)
func (v_ View) DisplayRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayRect:"), rect)
}
// Displays the view and all its subviews if any part of the view has been marked as needing display. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeeded()
func (v_ View) DisplayIfNeeded() {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeeded"))
}
// Acts as , confining drawing to a specified region of the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeeded(_:)
func (v_ View) DisplayIfNeededInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeededInRect:"), rect)
}
// Acts as , except that this method doesn’t back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity()
func (v_ View) DisplayIfNeededIgnoringOpacity() {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeededIgnoringOpacity"))
}
// Acts as , but confining drawing to and not backing up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity(_:)
func (v_ View) DisplayIfNeededInRectIgnoringOpacity(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayIfNeededInRectIgnoringOpacity:"), rect)
}
// Displays the view but confines drawing to a specified region and does not back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIgnoringOpacity(_:)
func (v_ View) DisplayRectIgnoringOpacity(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayRectIgnoringOpacity:"), rect)
}
// Causes the view and its descendants to be redrawn to the specified graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayIgnoringOpacity(_:in:)
func (v_ View) DisplayRectIgnoringOpacityInContext(rect unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("displayRectIgnoringOpacity:inContext:"), rect, context)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/displayLink(target:selector:)
func (v_ View) DisplayLinkWithTargetSelector(target objc.ID, selector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return rv
}
// Initiates a dragging operation from the view, allowing the user to drag a file icon to any application that has window or view objects that accept files. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dragFile(_:from:slideBack:event:)
func (v_ View) DragFileFromRectSlideBackEvent(filename string, rect unsafe.Pointer, flag bool, event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("dragFile:fromRect:slideBack:event:"), objc.String(filename), rect, flag, event)
	return rv
}
// Initiates a dragging operation from the view, allowing the user to drag arbitrary data with a specified icon into any application that has window or view objects that accept dragged data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dragImage:at:offset:event:pasteboard:source:slideBack:
func (v_ View) DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, viewLocation unsafe.Pointer, initialOffset unsafe.Pointer, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("dragImage:at:offset:event:pasteboard:source:slideBack:"), image, viewLocation, initialOffset, event, pboard, sourceObj, slideFlag)
}
// Initiates a dragging operation from the view, allowing the user to drag one or more promised files (or directories) into any application that has window or view objects that accept promised file data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/dragPromisedFiles(ofTypes:from:source:slideBack:event:)
func (v_ View) DragPromisedFilesOfTypesFromRectSourceSlideBackEvent(typeArray unsafe.Pointer, rect unsafe.Pointer, sourceObject objc.ID, flag bool, event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("dragPromisedFilesOfTypes:fromRect:source:slideBack:event:"), typeArray, rect, sourceObject, flag, event)
	return rv
}
// Overridden by subclasses to draw the view’s image within the specified rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/draw(_:)
func (v_ View) DrawRect(dirtyRect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawRect:"), dirtyRect)
}
// Draws the focus ring mask for the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/drawFocusRingMask()
func (v_ View) DrawFocusRingMask() {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawFocusRingMask"))
}
// Allows applications that use the AppKit pagination facility to draw additional marks on each logical page. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/drawPageBorder(with:)
func (v_ View) DrawPageBorderWithSize(borderSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawPageBorderWithSize:"), borderSize)
}
// Allows applications that use the AppKit pagination facility to draw additional marks on each printed sheet. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/drawSheetBorder(with:)
func (v_ View) DrawSheetBorderWithSize(borderSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("drawSheetBorderWithSize:"), borderSize)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/edgeInsetsForLayoutRegion:
func (v_ View) EdgeInsetsForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("edgeInsetsForLayoutRegion:"), layoutRegion)
	return rv
}
// This method is invoked at the end of the printing session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/endDocument()
func (v_ View) EndDocument() {
	objc.Send[objc.ID](v_.ID, objc.Sel("endDocument"))
}
// Writes the end of a conforming page. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/endPage()
func (v_ View) EndPage() {
	objc.Send[objc.ID](v_.ID, objc.Sel("endPage"))
}
// Sets the view to full screen mode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/enterFullScreenMode(_:withOptions:)
func (v_ View) EnterFullScreenModeWithOptions(screen unsafe.Pointer, options unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("enterFullScreenMode:withOptions:"), screen, options)
	return rv
}
// Randomly changes the frame of a view with an ambiguous layout between the different valid values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/exerciseAmbiguityInLayout()
func (v_ View) ExerciseAmbiguityInLayout() {
	objc.Send[objc.ID](v_.ID, objc.Sel("exerciseAmbiguityInLayout"))
}
// Instructs the view to exit full screen mode. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/exitFullScreenMode(options:)
func (v_ View) ExitFullScreenModeWithOptions(options unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("exitFullScreenModeWithOptions:"), options)
}
// Returns the view’s frame for a given alignment rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/frame(forAlignmentRect:)
func (v_ View) FrameForAlignmentRect(alignmentRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("frameForAlignmentRect:"), alignmentRect)
	return rv
}
// Returns the identifier for the view’s graphics state object, or 0 if the view doesn’t have a graphics state object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/gState()
func (v_ View) GState() int {
	rv := objc.Send[int](v_.ID, objc.Sel("gState"))
	return rv
}
// Returns by indirection a list of nonoverlapping rectangles that define the area the view is being asked to draw in . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/getRectsBeingDrawn(_:count:)
func (v_ View) GetRectsBeingDrawnCount(rects unsafe.Pointer, count unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("getRectsBeingDrawn:count:"), rects, count)
}
// Returns a list of rectangles indicating the newly exposed areas of the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/getRectsExposedDuringLiveResize(_:count:)
func (v_ View) GetRectsExposedDuringLiveResizeCount(exposedRects unsafe.Pointer, count unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("getRectsExposedDuringLiveResize:count:"), exposedRects, count)
}
// Returns the farthest descendant of the view in the view hierarchy (including itself) that contains a specified point, or if that point lies completely outside the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/hitTest(_:)
func (v_ View) HitTest(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("hitTest:"), point)
	return rv
}
// Invalidates the view’s intrinsic content size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/invalidateIntrinsicContentSize()
func (v_ View) InvalidateIntrinsicContentSize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("invalidateIntrinsicContentSize"))
}
// Returns a Boolean value that indicates whether the view is a subview of the specified view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isDescendant(of:)
func (v_ View) IsDescendantOf(view unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("isDescendantOf:"), view)
	return rv
}
// Returns whether a region of the view contains a specified point, accounting for whether the view is flipped or not. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/isMousePoint(_:in:)
func (v_ View) MouseInRect(point unsafe.Pointer, rect unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("mouse:inRect:"), point, rect)
	return rv
}
// Returns if the view handles page boundaries, otherwise. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/knowsPageRange(_:)
func (v_ View) KnowsPageRange(range_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("knowsPageRange:"), range_)
	return rv
}
// Perform layout in concert with the constraint-based layout system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layout()
func (v_ View) Layout() {
	objc.Send[objc.ID](v_.ID, objc.Sel("layout"))
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layoutGuideForLayoutRegion:
func (v_ View) LayoutGuideForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("layoutGuideForLayoutRegion:"), layoutRegion)
	return rv
}
// Updates the layout of the receiving view and its subviews based on the current views and constraints. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/layoutSubtreeIfNeeded()
func (v_ View) LayoutSubtreeIfNeeded() {
	objc.Send[objc.ID](v_.ID, objc.Sel("layoutSubtreeIfNeeded"))
}
// Invoked by to determine the location of the region of the view being printed on the physical page. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/locationOfPrintRect(_:)
func (v_ View) LocationOfPrintRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("locationOfPrintRect:"), rect)
	return rv
}
// Locks the focus on the view, so subsequent commands take effect in the view’s window and coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/lockFocus()
func (v_ View) LockFocus() {
	objc.Send[objc.ID](v_.ID, objc.Sel("lockFocus"))
}
// Locks the focus to the view atomically if the method returns and returns the value of . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/lockFocusIfCanDraw()
func (v_ View) LockFocusIfCanDraw() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("lockFocusIfCanDraw"))
	return rv
}
// Locks the focus to the view atomically if drawing can occur in the specified graphics context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/lockFocusIfCanDraw(in:)
func (v_ View) LockFocusIfCanDrawInContext(context unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("lockFocusIfCanDrawInContext:"), context)
	return rv
}
// Creates the view’s backing layer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/makeBackingLayer()
func (v_ View) MakeBackingLayer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("makeBackingLayer"))
	return rv
}
// Overridden by subclasses to return a context-sensitive pop-up menu for a given mouse-down event. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/menu(for:)
func (v_ View) MenuForEvent(event unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("menuForEvent:"), event)
	return rv
}
// Returns a Boolean value indicating whether the specified rectangle intersects any part of the area that the view is being asked to draw. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/needsToDraw(_:)
func (v_ View) NeedsToDrawRect(rect unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("needsToDrawRect:"), rect)
	return rv
}
// Invoked to notify the view that the focus ring mask requires updating. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/noteFocusRingMaskChanged()
func (v_ View) NoteFocusRingMaskChanged() {
	objc.Send[objc.ID](v_.ID, objc.Sel("noteFocusRingMaskChanged"))
}
// Implemented by subclasses to respond to key equivalents (also known as keyboard shortcuts). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/performKeyEquivalent(with:)
func (v_ View) PerformKeyEquivalent(event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("performKeyEquivalent:"), event)
	return rv
}
// Implemented by subclasses to respond to mnemonics. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/performMnemonic:
func (v_ View) PerformMnemonic(string string) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("performMnemonic:"), objc.String(string))
	return rv
}
// Prepares the overdraw region for drawing. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/prepareContent(in:)
func (v_ View) PrepareContentInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("prepareContentInRect:"), rect)
}
// Restores the view to an initial state so that it can be reused. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/prepareForReuse()
func (v_ View) PrepareForReuse() {
	objc.Send[objc.ID](v_.ID, objc.Sel("prepareForReuse"))
}
// This action method opens the Print panel, and if the user chooses an option other than canceling, prints the view and all its subviews to the device specified in the Print panel. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/printView(_:)
func (v_ View) Print(sender objc.ID) {
	objc.Send[objc.ID](v_.ID, objc.Sel("print:"), sender)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectForLayoutRegion:
func (v_ View) RectForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("rectForLayoutRegion:"), layoutRegion)
	return rv
}
// Implemented by subclasses to determine the portion of the view to be printed for the specified page number. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectForPage(_:)
func (v_ View) RectForPage(page int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("rectForPage:"), page)
	return rv
}
// Returns the appropriate rectangle to use when magnifying around the specified point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rectForSmartMagnification(at:in:)
func (v_ View) RectForSmartMagnificationAtPointInRect(location unsafe.Pointer, visibleRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("rectForSmartMagnificationAtPoint:inRect:"), location, visibleRect)
	return rv
}
// Notifies a clip view’s superview that either the clip view’s bounds rectangle or the document view’s frame rectangle has changed, and that any indicators of the scroll position need to be adjusted. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/reflectScrolledClipView(_:)
func (v_ View) ReflectScrolledClipView(clipView unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("reflectScrolledClipView:"), clipView)
}
// Registers the pasteboard types that the view will accept as the destination of an image-dragging session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/registerForDraggedTypes(_:)
func (v_ View) RegisterForDraggedTypes(newTypes unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("registerForDraggedTypes:"), newTypes)
}
// Frees the view’s graphics state object, if it has one. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/releaseGState()
func (v_ View) ReleaseGState() {
	objc.Send[objc.ID](v_.ID, objc.Sel("releaseGState"))
}
// Removes all tooltips assigned to the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeAllToolTips()
func (v_ View) RemoveAllToolTips() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeAllToolTips"))
}
// Removes the specified constraint from the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeConstraint(_:)
func (v_ View) RemoveConstraint(constraint unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeConstraint:"), constraint)
}
// Removes the specified constraints from the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeConstraints(_:)
func (v_ View) RemoveConstraints(constraints unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeConstraints:"), constraints)
}
// Completely removes a cursor rectangle from the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeCursorRect(_:cursor:)
func (v_ View) RemoveCursorRectCursor(rect unsafe.Pointer, object unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeCursorRect:cursor:"), rect, object)
}
// Unlinks the view from its superview and its window, removes it from the responder chain, and invalidates its cursor rectangles. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeFromSuperview()
func (v_ View) RemoveFromSuperview() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeFromSuperview"))
}
// Unlinks the view from its superview and its window and removes it from the responder chain, but does not invalidate its cursor rectangles to cause redrawing. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeFromSuperviewWithoutNeedingDisplay()
func (v_ View) RemoveFromSuperviewWithoutNeedingDisplay() {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeFromSuperviewWithoutNeedingDisplay"))
}
// Detaches a gesture recognizer from the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeGestureRecognizer(_:)
func (v_ View) RemoveGestureRecognizer(gestureRecognizer unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeGestureRecognizer:"), gestureRecognizer)
}
// Removes the provided layout guide from the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeLayoutGuide(_:)
func (v_ View) RemoveLayoutGuide(guide unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeLayoutGuide:"), guide)
}
// Removes the tooltip identified by specified tag. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeToolTip(_:)
func (v_ View) RemoveToolTip(tag unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeToolTip:"), tag)
}
// Removes a given tracking area from the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeTrackingArea(_:)
func (v_ View) RemoveTrackingArea(trackingArea unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeTrackingArea:"), trackingArea)
}
// Removes the tracking rectangle identified by a tag. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/removeTrackingRect(_:)
func (v_ View) RemoveTrackingRect(tag unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("removeTrackingRect:"), tag)
}
// Invalidates the view’s graphics state object, if it has one. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/renewGState()
func (v_ View) RenewGState() {
	objc.Send[objc.ID](v_.ID, objc.Sel("renewGState"))
}
// Replaces one of the view’s subviews with another view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/replaceSubview(_:with:)
func (v_ View) ReplaceSubviewWith(oldView unsafe.Pointer, newView unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("replaceSubview:with:"), oldView, newView)
}
// Overridden by subclasses to define their default cursor rectangles. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/resetCursorRects()
func (v_ View) ResetCursorRects() {
	objc.Send[objc.ID](v_.ID, objc.Sel("resetCursorRects"))
}
// Informs the view that the bounds size of its superview has changed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/resize(withOldSuperviewSize:)
func (v_ View) ResizeWithOldSuperviewSize(oldSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("resizeWithOldSuperviewSize:"), oldSize)
}
// Informs the view’s subviews that the view’s bounds rectangle size has changed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/resizeSubviews(withOldSize:)
func (v_ View) ResizeSubviewsWithOldSize(oldSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("resizeSubviewsWithOldSize:"), oldSize)
}
// Rotates the view’s bounds rectangle by a specified degree value around the origin of the coordinate system, (0.0, 0.0). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rotate(byDegrees:)
func (v_ View) RotateByAngle(angle float64) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rotateByAngle:"), angle)
}
// Informs the client that allowed the user to add . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didAdd:)
func (v_ View) RulerViewDidAddMarker(ruler unsafe.Pointer, marker unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:didAddMarker:"), ruler, marker)
}
// Informs the client that allowed the user to move . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didMove:)
func (v_ View) RulerViewDidMoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:didMoveMarker:"), ruler, marker)
}
// Informs the client that allowed the user to remove . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:didRemove:)
func (v_ View) RulerViewDidRemoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:didRemoveMarker:"), ruler, marker)
}
// Informs the client that the user has pressed the mouse button while the cursor is in the ruler area of . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:handleMouseDownWith:)
func (v_ View) RulerViewHandleMouseDown(ruler unsafe.Pointer, event unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:handleMouseDown:"), ruler, event)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:locationFor:)
func (v_ View) RulerViewLocationForPoint(ruler unsafe.Pointer, point unsafe.Pointer) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("rulerView:locationForPoint:"), ruler, point)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:pointForLocation:)
func (v_ View) RulerViewPointForLocation(ruler unsafe.Pointer, point float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("rulerView:pointForLocation:"), ruler, point)
	return rv
}
// Requests permission for to add , an NSRulerMarker being dragged onto the ruler by the user. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldAdd:)
func (v_ View) RulerViewShouldAddMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rulerView:shouldAddMarker:"), ruler, marker)
	return rv
}
// Requests permission for to move . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldMove:)
func (v_ View) RulerViewShouldMoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rulerView:shouldMoveMarker:"), ruler, marker)
	return rv
}
// Requests permission for to remove . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:shouldRemove:)
func (v_ View) RulerViewShouldRemoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("rulerView:shouldRemoveMarker:"), ruler, marker)
	return rv
}
// Informs the client that will add the new NSRulerMarker, . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willAdd:atLocation:)
func (v_ View) RulerViewWillAddMarkerAtLocation(ruler unsafe.Pointer, marker unsafe.Pointer, location float64) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("rulerView:willAddMarker:atLocation:"), ruler, marker, location)
	return rv
}
// Informs the client that will move , an NSRulerMarker already on the ruler view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willMove:toLocation:)
func (v_ View) RulerViewWillMoveMarkerToLocation(ruler unsafe.Pointer, marker unsafe.Pointer, location float64) float64 {
	rv := objc.Send[float64](v_.ID, objc.Sel("rulerView:willMoveMarker:toLocation:"), ruler, marker, location)
	return rv
}
// Informs the client view that is about to be appropriated by . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/rulerView(_:willSetClientView:)
func (v_ View) RulerViewWillSetClientView(ruler unsafe.Pointer, newClient unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("rulerView:willSetClientView:"), ruler, newClient)
}
// Scales the view’s coordinate system so that the unit square scales to the specified dimensions. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scaleUnitSquare(to:)
func (v_ View) ScaleUnitSquareToSize(newUnitSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scaleUnitSquareToSize:"), newUnitSize)
}
// Scrolls the view’s closest ancestor object so a point in the view lies at the origin of the clip view’s bounds rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:)
func (v_ View) ScrollPoint(point unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scrollPoint:"), point)
}
// Copies the visible portion of the view’s rendered image within a region and lays that portion down again at a specified offset . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:by:)
func (v_ View) ScrollRectBy(rect unsafe.Pointer, delta unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scrollRect:by:"), rect, delta)
}
// Notifies the superview of a clip view that the clip view needs to reset the origin of its bounds rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scroll(_:to:)
func (v_ View) ScrollClipViewToPoint(clipView unsafe.Pointer, point unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("scrollClipView:toPoint:"), clipView, point)
}
// Scrolls the view’s closest ancestor object the minimum distance needed so a specified region of the view becomes visible in the clip view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/scrollToVisible(_:)
func (v_ View) ScrollRectToVisible(rect unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("scrollRectToVisible:"), rect)
	return rv
}
// Sets the origin of the view’s bounds rectangle to a specified point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setBoundsOrigin(_:)
func (v_ View) SetBoundsOrigin(newOrigin unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBoundsOrigin:"), newOrigin)
}
// Sets the size of the view’s bounds rectangle to specified dimensions, inversely scaling its coordinate system relative to its frame rectangle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setBoundsSize(_:)
func (v_ View) SetBoundsSize(newSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setBoundsSize:"), newSize)
}
// Sets the priority with which a view resists being made smaller than its intrinsic size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setContentCompressionResistancePriority(_:for:)
func (v_ View) SetContentCompressionResistancePriorityForOrientation(priority unsafe.Pointer, orientation unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setContentCompressionResistancePriority:forOrientation:"), priority, orientation)
}
// Sets the priority with which a view resists being made larger than its intrinsic size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setContentHuggingPriority(_:for:)
func (v_ View) SetContentHuggingPriorityForOrientation(priority unsafe.Pointer, orientation unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setContentHuggingPriority:forOrientation:"), priority, orientation)
}
// Sets the origin of the view’s frame rectangle to the specified point, effectively repositioning it within its superview. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setFrameOrigin(_:)
func (v_ View) SetFrameOrigin(newOrigin unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameOrigin:"), newOrigin)
}
// Sets the size of the view’s frame rectangle to the specified dimensions, resizing it within its superview without affecting its coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setFrameSize(_:)
func (v_ View) SetFrameSize(newSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameSize:"), newSize)
}
// Invalidates the area around the focus ring. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setKeyboardFocusRingNeedsDisplay(_:)
func (v_ View) SetKeyboardFocusRingNeedsDisplayInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setKeyboardFocusRingNeedsDisplayInRect:"), rect)
}
// Marks the region of the view within the specified rectangle as needing display, increasing the view’s existing invalid region to include it. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setNeedsDisplay(_:)
func (v_ View) SetNeedsDisplayInRect(invalidRect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setNeedsDisplayInRect:"), invalidRect)
}
// Overridden by subclasses to (re)initialize the view’s graphics state object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/setUpGState()
func (v_ View) SetUpGState() {
	objc.Send[objc.ID](v_.ID, objc.Sel("setUpGState"))
}
// Allows the user to drag objects from the view without activating the app or moving the window of the view forward, possibly obscuring the destination. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/shouldDelayWindowOrdering(for:)
func (v_ View) ShouldDelayWindowOrderingForEvent(event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("shouldDelayWindowOrderingForEvent:"), event)
	return rv
}
// Returns a Boolean value indicating whether the view is being drawn to an environment that supports color. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/shouldDrawColor()
func (v_ View) ShouldDrawColor() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("shouldDrawColor"))
	return rv
}
// Shows a window displaying the definition of the attributed string at the specified point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/showDefinition(for:at:)
func (v_ View) ShowDefinitionForAttributedStringAtPoint(attrString unsafe.Pointer, textBaselineOrigin unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("showDefinitionForAttributedString:atPoint:"), attrString, textBaselineOrigin)
}
// Shows a window displaying the definition of the specified range of the attributed string. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/showDefinition(for:range:options:baselineOriginProvider:)
func (v_ View) ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString unsafe.Pointer, targetRange unsafe.Pointer, options unsafe.Pointer, originProvider unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("showDefinitionForAttributedString:range:options:baselineOriginProvider:"), attrString, targetRange, options, originProvider)
}
// Orders the view’s immediate subviews using the specified comparator function. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/sortSubviews(_:context:)
func (v_ View) SortSubviewsUsingFunctionContext(compare unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("sortSubviewsUsingFunction:context:"), compare, context)
}
// Translates the view’s coordinate system so that its origin moves to a new location. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/translateOrigin(to:)
func (v_ View) TranslateOriginToPoint(translation unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("translateOriginToPoint:"), translation)
}
// Translates the display rectangles by the specified delta. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/translateRectsNeedingDisplay(in:by:)
func (v_ View) TranslateRectsNeedingDisplayInRectBy(clipRect unsafe.Pointer, delta unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("translateRectsNeedingDisplayInRect:by:"), clipRect, delta)
}
// Unlocks focus from the current view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/unlockFocus()
func (v_ View) UnlockFocus() {
	objc.Send[objc.ID](v_.ID, objc.Sel("unlockFocus"))
}
// Unregisters the view as a possible destination in a dragging session. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/unregisterDraggedTypes()
func (v_ View) UnregisterDraggedTypes() {
	objc.Send[objc.ID](v_.ID, objc.Sel("unregisterDraggedTypes"))
}
// Update constraints for the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateConstraints()
func (v_ View) UpdateConstraints() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateConstraints"))
}
// Updates the constraints for the receiving view and its subviews. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateConstraintsForSubtreeIfNeeded()
func (v_ View) UpdateConstraintsForSubtreeIfNeeded() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateConstraintsForSubtreeIfNeeded"))
}
// Updates the view’s content by modifying its underlying layer. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateLayer()
func (v_ View) UpdateLayer() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateLayer"))
}
// Invoked automatically when the view’s geometry changes such that its tracking areas need to be recalculated. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/updateTrackingAreas()
func (v_ View) UpdateTrackingAreas() {
	objc.Send[objc.ID](v_.ID, objc.Sel("updateTrackingAreas"))
}
// Responds when the view’s backing store properties change. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidChangeBackingProperties()
func (v_ View) ViewDidChangeBackingProperties() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidChangeBackingProperties"))
}
// Informs the view that its effective appearance changed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidChangeEffectiveAppearance()
func (v_ View) ViewDidChangeEffectiveAppearance() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidChangeEffectiveAppearance"))
}
// Informs the view of the end of a live resize—the user has finished resizing the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidEndLiveResize()
func (v_ View) ViewDidEndLiveResize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidEndLiveResize"))
}
// Invoked when the view is hidden, either directly, or in response to an ancestor being hidden. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidHide()
func (v_ View) ViewDidHide() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidHide"))
}
// Informs the view that its superview has changed (possibly to ). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidMoveToSuperview()
func (v_ View) ViewDidMoveToSuperview() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidMoveToSuperview"))
}
// Informs the view that it has been added to a new view hierarchy. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidMoveToWindow()
func (v_ View) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidMoveToWindow"))
}
// Invoked when the view is unhidden, either directly, or in response to an ancestor being unhidden [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewDidUnhide()
func (v_ View) ViewDidUnhide() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewDidUnhide"))
}
// Informs the view that it’s required to draw content. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillDraw()
func (v_ View) ViewWillDraw() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillDraw"))
}
// Informs the view that its superview is about to change to the specified superview (which may be ). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillMove(toSuperview:)
func (v_ View) ViewWillMoveToSuperview(newSuperview unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillMoveToSuperview:"), newSuperview)
}
// Informs the view that it’s being added to the view hierarchy of the specified window object (which may be ). [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillMove(toWindow:)
func (v_ View) ViewWillMoveToWindow(newWindow unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillMoveToWindow:"), newWindow)
}
// Informs the view of the start of a live resize—the user has started resizing the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWillStartLiveResize()
func (v_ View) ViewWillStartLiveResize() {
	objc.Send[objc.ID](v_.ID, objc.Sel("viewWillStartLiveResize"))
}
// Returns the view’s nearest descendant (including itself) with a specific tag, or if no subview has that tag. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/viewWithTag(_:)
func (v_ View) ViewWithTag(tag int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("viewWithTag:"), tag)
	return rv
}
// Called just before a contextual menu for a view is opened on screen. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/willOpenMenu(_:with:)
func (v_ View) WillOpenMenuWithEvent(menu unsafe.Pointer, event unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("willOpenMenu:withEvent:"), menu, event)
}
// Overridden by subclasses to perform additional actions before subviews are removed from the view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/willRemoveSubview(_:)
func (v_ View) WillRemoveSubview(subview unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("willRemoveSubview:"), subview)
}
// Writes EPS data that draws the region of the view within a specified rectangle onto a pasteboard. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/writeEPS(inside:to:)
func (v_ View) WriteEPSInsideRectToPasteboard(rect unsafe.Pointer, pasteboard unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("writeEPSInsideRect:toPasteboard:"), rect, pasteboard)
}
// Writes PDF data that draws the region of the view within a specified rectangle onto a pasteboard. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/writePDF(inside:to:)
func (v_ View) WritePDFInsideRectToPasteboard(rect unsafe.Pointer, pasteboard unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("writePDFInsideRect:toPasteboard:"), rect, pasteboard)
}

