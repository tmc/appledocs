
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [View] class.
var ViewClass _ViewClass

func init() {
	ViewClass = _ViewClass{objc.GetClass("NSView")}
}

type _ViewClass struct {
	objc.Class
}

// An interface definition for the [View] class.
type IView interface {
	ID() objc.ID
	AcceptsFirstMouse(event unsafe.Pointer) bool
	AddConstraint(constraint unsafe.Pointer)
	AddConstraints(constraints unsafe.Pointer)
	AddCursorRectCursor(rect unsafe.Pointer, object unsafe.Pointer)
	AddGestureRecognizer(gestureRecognizer unsafe.Pointer)
	AddLayoutGuide(guide unsafe.Pointer)
	AddSubview(view unsafe.Pointer)
	AddSubviewPositionedRelativeTo(view unsafe.Pointer, place unsafe.Pointer, otherView unsafe.Pointer)
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
	ConvertPointFromBacking(point unsafe.Pointer) unsafe.Pointer
	ConvertPointFromBase(point unsafe.Pointer) unsafe.Pointer
	ConvertPointFromLayer(point unsafe.Pointer) unsafe.Pointer
	ConvertPointFromView(point unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertPointToBacking(point unsafe.Pointer) unsafe.Pointer
	ConvertPointToBase(point unsafe.Pointer) unsafe.Pointer
	ConvertPointToLayer(point unsafe.Pointer) unsafe.Pointer
	ConvertPointToView(point unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertRectFromBacking(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectFromBase(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectFromLayer(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectFromView(rect unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertRectToBacking(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectToBase(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectToLayer(rect unsafe.Pointer) unsafe.Pointer
	ConvertRectToView(rect unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertSizeFromBacking(size unsafe.Pointer) unsafe.Pointer
	ConvertSizeFromBase(size unsafe.Pointer) unsafe.Pointer
	ConvertSizeFromLayer(size unsafe.Pointer) unsafe.Pointer
	ConvertSizeFromView(size unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	ConvertSizeToBacking(size unsafe.Pointer) unsafe.Pointer
	ConvertSizeToBase(size unsafe.Pointer) unsafe.Pointer
	ConvertSizeToLayer(size unsafe.Pointer) unsafe.Pointer
	ConvertSizeToView(size unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer
	DataWithEPSInsideRect(rect unsafe.Pointer) unsafe.Pointer
	DataWithPDFInsideRect(rect unsafe.Pointer) unsafe.Pointer
	DidAddSubview(subview unsafe.Pointer)
	DidCloseMenuWithEvent(menu unsafe.Pointer, event unsafe.Pointer)
	DiscardCursorRects()
	Display()
	DisplayIfNeeded()
	DisplayIfNeededIgnoringOpacity()
	DisplayIfNeededInRect(rect unsafe.Pointer)
	DisplayIfNeededInRectIgnoringOpacity(rect unsafe.Pointer)
	DisplayLinkWithTargetSelector(target objc.ID, selector objc.SEL) unsafe.Pointer
	DisplayRect(rect unsafe.Pointer)
	DisplayRectIgnoringOpacity(rect unsafe.Pointer)
	DisplayRectIgnoringOpacityInContext(rect unsafe.Pointer, context unsafe.Pointer)
	DragFileFromRectSlideBackEvent(filename unsafe.Pointer, rect unsafe.Pointer, flag bool, event unsafe.Pointer) bool
	DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, viewLocation unsafe.Pointer, initialOffset unsafe.Pointer, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool)
	DragPromisedFilesOfTypesFromRectSourceSlideBackEvent(typeArray unsafe.Pointer, rect unsafe.Pointer, sourceObject objc.ID, flag bool, event unsafe.Pointer) bool
	DrawFocusRingMask()
	DrawPageBorderWithSize(borderSize unsafe.Pointer)
	DrawRect(dirtyRect unsafe.Pointer)
	DrawSheetBorderWithSize(borderSize unsafe.Pointer)
	EdgeInsetsForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer
	EndDocument()
	EndPage()
	EnterFullScreenModeWithOptions(screen unsafe.Pointer, options unsafe.Pointer) bool
	ExerciseAmbiguityInLayout()
	ExitFullScreenModeWithOptions(options unsafe.Pointer)
	FrameForAlignmentRect(alignmentRect unsafe.Pointer) unsafe.Pointer
	GState() int
	GetRectsBeingDrawnCount(rects unsafe.Pointer, count int)
	GetRectsExposedDuringLiveResizeCount(exposedRects unsafe.Pointer, count int)
	HitTest(point unsafe.Pointer) unsafe.Pointer
	InitWithCoder(coder unsafe.Pointer) unsafe.Pointer
	InitWithFrame(frameRect unsafe.Pointer) unsafe.Pointer
	InvalidateIntrinsicContentSize()
	IsDescendantOf(view unsafe.Pointer) bool
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
	MouseInRect(point unsafe.Pointer, rect unsafe.Pointer) bool
	NeedsToDrawRect(rect unsafe.Pointer) bool
	NoteFocusRingMaskChanged()
	PerformKeyEquivalent(event unsafe.Pointer) bool
	PerformMnemonic(string unsafe.Pointer) bool
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
	ResizeSubviewsWithOldSize(oldSize unsafe.Pointer)
	ResizeWithOldSuperviewSize(oldSize unsafe.Pointer)
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
	ScrollClipViewToPoint(clipView unsafe.Pointer, point unsafe.Pointer)
	ScrollPoint(point unsafe.Pointer)
	ScrollRectBy(rect unsafe.Pointer, delta unsafe.Pointer)
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

type View struct {
	id objc.ID
}

func ViewFrom(ptr unsafe.Pointer) View {
	return View{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ View) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _ViewClass) Alloc() View {
	rv := objc.Send[View](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _ViewClass) New() View {
	rv := objc.Send[View](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewView creates and returns a new initialized instance.
func NewView() View {
	return ViewClass.New()
}

// Init initializes the instance.
func (v_ View) Init() View {
	rv := objc.Send[View](v_.ID(), selInit)
	return rv
}
// Overridden by subclasses to return   if the view should be sent a   message for an initial mouse-down event,   if not. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/acceptsFirstMouse(for:)
func (v_ View) AcceptsFirstMouse(event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("acceptsFirstMouse:"), event)
	return rv
}
// Adds a constraint on the layout of the receiving view or its subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addConstraint(_:)
func (v_ View) AddConstraint(constraint unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("addConstraint:"), constraint)
}
// Adds multiple constraints on the layout of the receiving view or its subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addConstraints(_:)
func (v_ View) AddConstraints(constraints unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("addConstraints:"), constraints)
}
// Establishes  the cursor to be used when the mouse pointer lies within a specified region. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addCursorRect(_:cursor:)
func (v_ View) AddCursorRectCursor(rect unsafe.Pointer, object unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("addCursorRect:cursor:"), rect, object)
}
// Attaches a gesture recognizer to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addGestureRecognizer(_:)
func (v_ View) AddGestureRecognizer(gestureRecognizer unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("addGestureRecognizer:"), gestureRecognizer)
}
// Adds the provided layout guide to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addLayoutGuide(_:)
func (v_ View) AddLayoutGuide(guide unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("addLayoutGuide:"), guide)
}
// Adds a view to the view’s subviews so it’s displayed above its siblings. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addSubview(_:)
func (v_ View) AddSubview(view unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("addSubview:"), view)
}
// Inserts a view among the view’s subviews so it’s displayed immediately above or below another view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addSubview(_:positioned:relativeTo:)
func (v_ View) AddSubviewPositionedRelativeTo(view unsafe.Pointer, place unsafe.Pointer, otherView unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("addSubview:positioned:relativeTo:"), view, place, otherView)
}
// Creates a tooltip for a defined area in the view and returns a tag that identifies the tooltip rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addToolTip(_:owner:userData:)
func (v_ View) AddToolTipRectOwnerUserData(rect unsafe.Pointer, owner objc.ID, data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("addToolTipRect:owner:userData:"), rect, owner, data)
	return rv
}
// Adds a given tracking area to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addTrackingArea(_:)
func (v_ View) AddTrackingArea(trackingArea unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("addTrackingArea:"), trackingArea)
}
// Establishes  an area for tracking mouse-entered and mouse-exited events within the view and returns a tag that identifies the tracking rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addTrackingRect(_:owner:userData:assumeInside:)
func (v_ View) AddTrackingRectOwnerUserDataAssumeInside(rect unsafe.Pointer, owner objc.ID, data unsafe.Pointer, flag bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("addTrackingRect:owner:userData:assumeInside:"), rect, owner, data, flag)
	return rv
}
// Overridden by subclasses to adjust page height during automatic pagination. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/adjustPageHeightNew(_:top:bottom:limit:)
func (v_ View) AdjustPageHeightNewTopBottomLimit(newBottom float64, oldTop float64, oldBottom float64, bottomLimit float64) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("adjustPageHeightNew:top:bottom:limit:"), newBottom, oldTop, oldBottom, bottomLimit)
}
// Overridden by subclasses to adjust page width during automatic pagination. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/adjustPageWidthNew(_:left:right:limit:)
func (v_ View) AdjustPageWidthNewLeftRightLimit(newRight float64, oldLeft float64, oldRight float64, rightLimit float64) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("adjustPageWidthNew:left:right:limit:"), newRight, oldLeft, oldRight, rightLimit)
}
// Overridden by subclasses to modify a given rectangle, returning the altered rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/adjustScroll(_:)
func (v_ View) AdjustScroll(newVisible unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("adjustScroll:"), newVisible)
	return rv
}
// Returns the view’s alignment rectangle for a given frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/alignmentRect(forFrame:)
func (v_ View) AlignmentRectForFrame(frame unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("alignmentRectForFrame:"), frame)
	return rv
}
// Causes the view to maintain a private graphics state object, which encapsulates all parameters of the graphics environment. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/allocateGState()
func (v_ View) AllocateGState() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("allocateGState"))
}
// Returns the closest ancestor shared by the view and another specified view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/ancestorShared(with:)
func (v_ View) AncestorSharedWithView(view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("ancestorSharedWithView:"), view)
	return rv
}
// Scrolls the view’s closest ancestor   object proportionally to the distance of an event that occurs outside of it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/autoscroll(with:)
func (v_ View) Autoscroll(event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("autoscroll:"), event)
	return rv
}
// Returns a backing store pixel-aligned rectangle in local view coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/backingAlignedRect(_:options:)
func (v_ View) BackingAlignedRectOptions(rect unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("backingAlignedRect:options:"), rect, options)
	return rv
}
// Invoked at the beginning of the printing session, this method sets up the current graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/beginDocument()
func (v_ View) BeginDocument() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("beginDocument"))
}
// Initiates a dragging session with a group of dragging items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/beginDraggingSession(with:event:source:)
func (v_ View) BeginDraggingSessionWithItemsEventSource(items unsafe.Pointer, event unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("beginDraggingSessionWithItems:event:source:"), items, event, source)
	return rv
}
// Called at the beginning of each page, this method sets up the coordinate system so that a region inside the view’s bounds is translated to a specified location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/beginPage(in:atPlacement:)
func (v_ View) BeginPageInRectAtPlacement(rect unsafe.Pointer, location unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("beginPageInRect:atPlacement:"), rect, location)
}
// Returns a bitmap-representation object suitable for caching the specified portion of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/bitmapImageRepForCachingDisplay(in:)
func (v_ View) BitmapImageRepForCachingDisplayInRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("bitmapImageRepForCachingDisplayInRect:"), rect)
	return rv
}
// Draws the specified area of the view, and its descendants, into a provided bitmap-representation object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/cacheDisplay(in:to:)
func (v_ View) CacheDisplayInRectToBitmapImageRep(rect unsafe.Pointer, bitmapImageRep unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("cacheDisplayInRect:toBitmapImageRep:"), rect, bitmapImageRep)
}
// Converts the corners of a specified rectangle to lie on the center of device pixels, which is useful in compensating for rendering overscanning when the coordinate system has been scaled. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/centerScanRect(_:)
func (v_ View) CenterScanRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("centerScanRect:"), rect)
	return rv
}
// Returns the constraints impacting the layout of the view for a given orientation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/constraintsAffectingLayout(for:)
func (v_ View) ConstraintsAffectingLayoutForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("constraintsAffectingLayoutForOrientation:"), orientation)
	return rv
}
// Returns the priority with which a view resists being made smaller than its intrinsic size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/contentCompressionResistancePriority(for:)
func (v_ View) ContentCompressionResistancePriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("contentCompressionResistancePriorityForOrientation:"), orientation)
	return rv
}
// Returns the priority with which a view resists being made larger than its intrinsic size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/contentHuggingPriority(for:)
func (v_ View) ContentHuggingPriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("contentHuggingPriorityForOrientation:"), orientation)
	return rv
}
// Converts a point from the coordinate system of a given view to that of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:from:)-1dq9l
func (v_ View) ConvertPointFromView(point unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertPoint:fromView:"), point, view)
	return rv
}
// Converts a size from another view’s coordinate system to that of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:from:)-40x0w
func (v_ View) ConvertSizeFromView(size unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertSize:fromView:"), size, view)
	return rv
}
// Converts a rectangle from the coordinate system of another view to that of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:from:)-7fbb6
func (v_ View) ConvertRectFromView(rect unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertRect:fromView:"), rect, view)
	return rv
}
// Converts a rectangle from the view’s coordinate system to that of another view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:to:)-3cqqt
func (v_ View) ConvertRectToView(rect unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertRect:toView:"), rect, view)
	return rv
}
// Converts a size from the view’s coordinate system to that of another view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:to:)-5nptx
func (v_ View) ConvertSizeToView(size unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertSize:toView:"), size, view)
	return rv
}
// Converts a point from the view’s coordinate system to that of a given view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:to:)-6u9ir
func (v_ View) ConvertPointToView(point unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertPoint:toView:"), point, view)
	return rv
}
// Converts a point from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromBacking(_:)-229ps
func (v_ View) ConvertPointFromBacking(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertPointFromBacking:"), point)
	return rv
}
// Converts a rectangle from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromBacking(_:)-2njpa
func (v_ View) ConvertRectFromBacking(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertRectFromBacking:"), rect)
	return rv
}
// Converts a size from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromBacking(_:)-4agf9
func (v_ View) ConvertSizeFromBacking(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertSizeFromBacking:"), size)
	return rv
}
// Convert the point from the layer’s interior coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromLayer(_:)-3nsbu
func (v_ View) ConvertPointFromLayer(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertPointFromLayer:"), point)
	return rv
}
// Convert the size from the layer’s interior coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromLayer(_:)-3usqp
func (v_ View) ConvertSizeFromLayer(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertSizeFromLayer:"), size)
	return rv
}
// Convert the rectangle from the layer’s interior coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromLayer(_:)-8s5bi
func (v_ View) ConvertRectFromLayer(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertRectFromLayer:"), rect)
	return rv
}
// Converts the point from the base coordinate system to the view’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertPointFromBase:
func (v_ View) ConvertPointFromBase(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertPointFromBase:"), point)
	return rv
}
// Converts the point from the view’s coordinate system to the base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertPointToBase:
func (v_ View) ConvertPointToBase(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertPointToBase:"), point)
	return rv
}
// Converts the rectangle from the base coordinate system to the view’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertRectFromBase:
func (v_ View) ConvertRectFromBase(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertRectFromBase:"), rect)
	return rv
}
// Converts the rectangle from the view’s coordinate system to the base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertRectToBase:
func (v_ View) ConvertRectToBase(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertRectToBase:"), rect)
	return rv
}
// Converts the size from the base coordinate system to the view’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertSizeFromBase:
func (v_ View) ConvertSizeFromBase(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertSizeFromBase:"), size)
	return rv
}
// Converts the size from the view’s coordinate system to the base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertSizeToBase:
func (v_ View) ConvertSizeToBase(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertSizeToBase:"), size)
	return rv
}
// Converts a point from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToBacking(_:)-2xx45
func (v_ View) ConvertPointToBacking(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertPointToBacking:"), point)
	return rv
}
// Converts a rectangle from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToBacking(_:)-3zors
func (v_ View) ConvertRectToBacking(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertRectToBacking:"), rect)
	return rv
}
// Converts a size from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToBacking(_:)-4ra9y
func (v_ View) ConvertSizeToBacking(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertSizeToBacking:"), size)
	return rv
}
// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToLayer(_:)-160pw
func (v_ View) ConvertRectToLayer(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertRectToLayer:"), rect)
	return rv
}
// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToLayer(_:)-2vozx
func (v_ View) ConvertSizeToLayer(size unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertSizeToLayer:"), size)
	return rv
}
// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToLayer(_:)-44u7d
func (v_ View) ConvertPointToLayer(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("convertPointToLayer:"), point)
	return rv
}
// Returns EPS data that draws the region of the view within a specified rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/dataWithEPS(inside:)
func (v_ View) DataWithEPSInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("dataWithEPSInsideRect:"), rect)
	return rv
}
// Returns PDF data that draws the region of the view within a specified rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/dataWithPDF(inside:)
func (v_ View) DataWithPDFInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("dataWithPDFInsideRect:"), rect)
	return rv
}
// Overridden by subclasses to perform additional actions when subviews are added to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/didAddSubview(_:)
func (v_ View) DidAddSubview(subview unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("didAddSubview:"), subview)
}
// Called after a contextual menu that was displayed from the receiving view has been closed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/didCloseMenu(_:with:)
func (v_ View) DidCloseMenuWithEvent(menu unsafe.Pointer, event unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("didCloseMenu:withEvent:"), menu, event)
}
// Invalidates all cursor rectangles set up using  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/discardCursorRects()
func (v_ View) DiscardCursorRects() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("discardCursorRects"))
}
// Displays the view and all its subviews if possible, invoking each of the   methods  ,  , and   as necessary. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/display()
func (v_ View) Display() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("display"))
}
// Acts as  , but confining drawing to a rectangular region of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/display(_:)
func (v_ View) DisplayRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("displayRect:"), rect)
}
// Displays the view and all its subviews if any part of the view has been marked as needing display. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeeded()
func (v_ View) DisplayIfNeeded() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("displayIfNeeded"))
}
// Acts as  , confining drawing to a specified region of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeeded(_:)
func (v_ View) DisplayIfNeededInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("displayIfNeededInRect:"), rect)
}
// Acts as  , except that this method doesn’t back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity()
func (v_ View) DisplayIfNeededIgnoringOpacity() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("displayIfNeededIgnoringOpacity"))
}
// Acts as  , but confining drawing to   and not backing up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity(_:)
func (v_ View) DisplayIfNeededInRectIgnoringOpacity(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("displayIfNeededInRectIgnoringOpacity:"), rect)
}
// Displays the view but confines drawing to a specified region and does not back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIgnoringOpacity(_:)
func (v_ View) DisplayRectIgnoringOpacity(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("displayRectIgnoringOpacity:"), rect)
}
// Causes the view and its descendants to be redrawn to the specified graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIgnoringOpacity(_:in:)
func (v_ View) DisplayRectIgnoringOpacityInContext(rect unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("displayRectIgnoringOpacity:inContext:"), rect, context)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayLink(target:selector:)
func (v_ View) DisplayLinkWithTargetSelector(target objc.ID, selector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("displayLinkWithTarget:selector:"), target, selector)
	return rv
}
// Initiates a dragging operation from the view, allowing the user to drag a file icon to any application that has window or view objects that accept files. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/dragFile(_:from:slideBack:event:)
func (v_ View) DragFileFromRectSlideBackEvent(filename unsafe.Pointer, rect unsafe.Pointer, flag bool, event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("dragFile:fromRect:slideBack:event:"), filename, rect, flag, event)
	return rv
}
// Initiates a dragging operation from the view, allowing the user to drag arbitrary data with a specified icon into any application that has window or view objects that accept dragged data. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/dragImage:at:offset:event:pasteboard:source:slideBack:
func (v_ View) DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, viewLocation unsafe.Pointer, initialOffset unsafe.Pointer, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("dragImage:at:offset:event:pasteboard:source:slideBack:"), image, viewLocation, initialOffset, event, pboard, sourceObj, slideFlag)
}
// Initiates a dragging operation from the view, allowing the user to drag one or more promised files (or directories) into any application that has window or view objects that accept promised file data. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/dragPromisedFiles(ofTypes:from:source:slideBack:event:)
func (v_ View) DragPromisedFilesOfTypesFromRectSourceSlideBackEvent(typeArray unsafe.Pointer, rect unsafe.Pointer, sourceObject objc.ID, flag bool, event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("dragPromisedFilesOfTypes:fromRect:source:slideBack:event:"), typeArray, rect, sourceObject, flag, event)
	return rv
}
// Overridden by subclasses to draw the view’s image within the specified rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/draw(_:)
func (v_ View) DrawRect(dirtyRect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("drawRect:"), dirtyRect)
}
// Draws the focus ring mask for the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/drawFocusRingMask()
func (v_ View) DrawFocusRingMask() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("drawFocusRingMask"))
}
// Allows applications that use the AppKit pagination facility to draw additional marks on each logical page. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/drawPageBorder(with:)
func (v_ View) DrawPageBorderWithSize(borderSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("drawPageBorderWithSize:"), borderSize)
}
// Allows applications that use the AppKit pagination facility to draw additional marks on each printed sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/drawSheetBorder(with:)
func (v_ View) DrawSheetBorderWithSize(borderSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("drawSheetBorderWithSize:"), borderSize)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/edgeInsetsForLayoutRegion:
func (v_ View) EdgeInsetsForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("edgeInsetsForLayoutRegion:"), layoutRegion)
	return rv
}
// This method is invoked at the end of the printing session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/endDocument()
func (v_ View) EndDocument() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("endDocument"))
}
// Writes the end of a conforming page. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/endPage()
func (v_ View) EndPage() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("endPage"))
}
// Sets the view to full screen mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/enterFullScreenMode(_:withOptions:)
func (v_ View) EnterFullScreenModeWithOptions(screen unsafe.Pointer, options unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("enterFullScreenMode:withOptions:"), screen, options)
	return rv
}
// Randomly changes the frame of a view with an ambiguous layout between the different valid values. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/exerciseAmbiguityInLayout()
func (v_ View) ExerciseAmbiguityInLayout() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("exerciseAmbiguityInLayout"))
}
// Instructs the view to exit full screen mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/exitFullScreenMode(options:)
func (v_ View) ExitFullScreenModeWithOptions(options unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("exitFullScreenModeWithOptions:"), options)
}
// Returns the view’s frame for a given alignment rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/frame(forAlignmentRect:)
func (v_ View) FrameForAlignmentRect(alignmentRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("frameForAlignmentRect:"), alignmentRect)
	return rv
}
// Returns the identifier for the view’s graphics state object, or 0 if the view doesn’t have a graphics state object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/gState()
func (v_ View) GState() int {
	rv := objc.Send[int](v_.ID(), objc.RegisterName("gState"))
	return rv
}
// Returns by indirection a list of nonoverlapping rectangles that define the area the view is being asked to draw in  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/getRectsBeingDrawn(_:count:)
func (v_ View) GetRectsBeingDrawnCount(rects unsafe.Pointer, count int) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("getRectsBeingDrawn:count:"), rects, count)
}
// Returns a list of rectangles indicating the newly exposed areas of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/getRectsExposedDuringLiveResize(_:count:)
func (v_ View) GetRectsExposedDuringLiveResizeCount(exposedRects unsafe.Pointer, count int) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("getRectsExposedDuringLiveResize:count:"), exposedRects, count)
}
// Returns the farthest descendant of the view in the view hierarchy (including itself) that contains a specified point, or   if that point lies completely outside the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/hitTest(_:)
func (v_ View) HitTest(point unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("hitTest:"), point)
	return rv
}
// Initializes a view using from data in the specified coder object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/init(coder:)
func (v_ View) InitWithCoder(coder unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("initWithCoder:"), coder)
	return rv
}
// Initializes and returns a newly allocated   object with a specified frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/init(frame:)
func (v_ View) InitWithFrame(frameRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("initWithFrame:"), frameRect)
	return rv
}
// Invalidates the view’s intrinsic content size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/invalidateIntrinsicContentSize()
func (v_ View) InvalidateIntrinsicContentSize() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("invalidateIntrinsicContentSize"))
}
// Returns a Boolean value that indicates whether the view is a subview of the specified view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isDescendant(of:)
func (v_ View) IsDescendantOf(view unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("isDescendantOf:"), view)
	return rv
}
// Returns whether a region of the view contains a specified point, accounting for whether the view is flipped or not. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isMousePoint(_:in:)
func (v_ View) MouseInRect(point unsafe.Pointer, rect unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("mouse:inRect:"), point, rect)
	return rv
}
// Returns   if the view handles page boundaries,   otherwise. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/knowsPageRange(_:)
func (v_ View) KnowsPageRange(range_ unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("knowsPageRange:"), range_)
	return rv
}
// Perform layout in concert with the constraint-based layout system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layout()
func (v_ View) Layout() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("layout"))
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layoutGuideForLayoutRegion:
func (v_ View) LayoutGuideForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("layoutGuideForLayoutRegion:"), layoutRegion)
	return rv
}
// Updates the layout of the receiving view and its subviews based on the current views and constraints. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layoutSubtreeIfNeeded()
func (v_ View) LayoutSubtreeIfNeeded() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("layoutSubtreeIfNeeded"))
}
// Invoked by   to determine the location of the region of the view being printed on the physical page. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/locationOfPrintRect(_:)
func (v_ View) LocationOfPrintRect(rect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("locationOfPrintRect:"), rect)
	return rv
}
// Locks the focus on the view, so subsequent commands take effect in the view’s window and coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/lockFocus()
func (v_ View) LockFocus() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("lockFocus"))
}
// Locks the focus to the view atomically if the   method returns   and returns the value of  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/lockFocusIfCanDraw()
func (v_ View) LockFocusIfCanDraw() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("lockFocusIfCanDraw"))
	return rv
}
// Locks the focus to the view atomically if drawing can occur in the specified graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/lockFocusIfCanDraw(in:)
func (v_ View) LockFocusIfCanDrawInContext(context unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("lockFocusIfCanDrawInContext:"), context)
	return rv
}
// Creates the view’s backing layer. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/makeBackingLayer()
func (v_ View) MakeBackingLayer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("makeBackingLayer"))
	return rv
}
// Overridden by subclasses to return a context-sensitive pop-up menu for a given mouse-down event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/menu(for:)
func (v_ View) MenuForEvent(event unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("menuForEvent:"), event)
	return rv
}
// Returns a Boolean value indicating whether the specified rectangle intersects any part of the area that the view is being asked to draw. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/needsToDraw(_:)
func (v_ View) NeedsToDrawRect(rect unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("needsToDrawRect:"), rect)
	return rv
}
// Invoked to notify the view that the focus ring mask requires updating. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/noteFocusRingMaskChanged()
func (v_ View) NoteFocusRingMaskChanged() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("noteFocusRingMaskChanged"))
}
// Implemented by subclasses to respond to key equivalents (also known as keyboard shortcuts). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/performKeyEquivalent(with:)
func (v_ View) PerformKeyEquivalent(event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("performKeyEquivalent:"), event)
	return rv
}
// Implemented by subclasses to respond to mnemonics. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/performMnemonic:
func (v_ View) PerformMnemonic(string unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("performMnemonic:"), string)
	return rv
}
// Prepares the overdraw region for drawing. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/prepareContent(in:)
func (v_ View) PrepareContentInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("prepareContentInRect:"), rect)
}
// Restores the view to an initial state so that it can be reused. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/prepareForReuse()
func (v_ View) PrepareForReuse() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("prepareForReuse"))
}
// This action method opens the Print panel, and if the user chooses an option other than canceling, prints the view and all its subviews to the device specified in the Print panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/printView(_:)
func (v_ View) Print(sender objc.ID) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("print:"), sender)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rectForLayoutRegion:
func (v_ View) RectForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("rectForLayoutRegion:"), layoutRegion)
	return rv
}
// Implemented by subclasses to determine the portion of the view to be printed for the specified page number. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rectForPage(_:)
func (v_ View) RectForPage(page int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("rectForPage:"), page)
	return rv
}
// Returns the appropriate rectangle to use when magnifying around the specified point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rectForSmartMagnification(at:in:)
func (v_ View) RectForSmartMagnificationAtPointInRect(location unsafe.Pointer, visibleRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("rectForSmartMagnificationAtPoint:inRect:"), location, visibleRect)
	return rv
}
// Notifies a clip view’s superview that either the clip view’s bounds rectangle or the document view’s frame rectangle has changed, and that any indicators of the scroll position need to be adjusted. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/reflectScrolledClipView(_:)
func (v_ View) ReflectScrolledClipView(clipView unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("reflectScrolledClipView:"), clipView)
}
// Registers the pasteboard types that the view will accept as the destination of an image-dragging session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/registerForDraggedTypes(_:)
func (v_ View) RegisterForDraggedTypes(newTypes unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("registerForDraggedTypes:"), newTypes)
}
// Frees the view’s graphics state object, if it has one. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/releaseGState()
func (v_ View) ReleaseGState() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("releaseGState"))
}
// Removes all tooltips assigned to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeAllToolTips()
func (v_ View) RemoveAllToolTips() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeAllToolTips"))
}
// Removes the specified constraint from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeConstraint(_:)
func (v_ View) RemoveConstraint(constraint unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeConstraint:"), constraint)
}
// Removes the specified constraints from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeConstraints(_:)
func (v_ View) RemoveConstraints(constraints unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeConstraints:"), constraints)
}
// Completely removes a cursor rectangle from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeCursorRect(_:cursor:)
func (v_ View) RemoveCursorRectCursor(rect unsafe.Pointer, object unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeCursorRect:cursor:"), rect, object)
}
// Unlinks the view from its superview and its window, removes it from the responder chain, and invalidates its cursor rectangles. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeFromSuperview()
func (v_ View) RemoveFromSuperview() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeFromSuperview"))
}
// Unlinks the view from its superview and its window and removes it from the responder chain, but does not invalidate its cursor rectangles to cause redrawing. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeFromSuperviewWithoutNeedingDisplay()
func (v_ View) RemoveFromSuperviewWithoutNeedingDisplay() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeFromSuperviewWithoutNeedingDisplay"))
}
// Detaches a gesture recognizer from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeGestureRecognizer(_:)
func (v_ View) RemoveGestureRecognizer(gestureRecognizer unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeGestureRecognizer:"), gestureRecognizer)
}
// Removes the provided layout guide from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeLayoutGuide(_:)
func (v_ View) RemoveLayoutGuide(guide unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeLayoutGuide:"), guide)
}
// Removes the tooltip identified by specified tag. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeToolTip(_:)
func (v_ View) RemoveToolTip(tag unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeToolTip:"), tag)
}
// Removes a given tracking area from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeTrackingArea(_:)
func (v_ View) RemoveTrackingArea(trackingArea unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeTrackingArea:"), trackingArea)
}
// Removes the tracking rectangle identified by a tag. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeTrackingRect(_:)
func (v_ View) RemoveTrackingRect(tag unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("removeTrackingRect:"), tag)
}
// Invalidates the view’s graphics state object, if it has one. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/renewGState()
func (v_ View) RenewGState() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("renewGState"))
}
// Replaces one of the view’s subviews with another view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/replaceSubview(_:with:)
func (v_ View) ReplaceSubviewWith(oldView unsafe.Pointer, newView unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("replaceSubview:with:"), oldView, newView)
}
// Overridden by subclasses to define their default cursor rectangles. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/resetCursorRects()
func (v_ View) ResetCursorRects() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("resetCursorRects"))
}
// Informs the view that the bounds size of its superview has changed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/resize(withOldSuperviewSize:)
func (v_ View) ResizeWithOldSuperviewSize(oldSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("resizeWithOldSuperviewSize:"), oldSize)
}
// Informs the view’s subviews that the view’s bounds rectangle size has changed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/resizeSubviews(withOldSize:)
func (v_ View) ResizeSubviewsWithOldSize(oldSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("resizeSubviewsWithOldSize:"), oldSize)
}
// Rotates the view’s bounds rectangle by a specified degree value around the origin of the coordinate system, (0.0, 0.0). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rotate(byDegrees:)
func (v_ View) RotateByAngle(angle float64) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("rotateByAngle:"), angle)
}
// Informs the client that   allowed the user to add  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:didAdd:)
func (v_ View) RulerViewDidAddMarker(ruler unsafe.Pointer, marker unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("rulerView:didAddMarker:"), ruler, marker)
}
// Informs the client that   allowed the user to move  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:didMove:)
func (v_ View) RulerViewDidMoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("rulerView:didMoveMarker:"), ruler, marker)
}
// Informs the client that   allowed the user to remove  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:didRemove:)
func (v_ View) RulerViewDidRemoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("rulerView:didRemoveMarker:"), ruler, marker)
}
// Informs the client that the user has pressed the mouse button while the cursor is in the ruler area of  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:handleMouseDownWith:)
func (v_ View) RulerViewHandleMouseDown(ruler unsafe.Pointer, event unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("rulerView:handleMouseDown:"), ruler, event)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:locationFor:)
func (v_ View) RulerViewLocationForPoint(ruler unsafe.Pointer, point unsafe.Pointer) float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("rulerView:locationForPoint:"), ruler, point)
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:pointForLocation:)
func (v_ View) RulerViewPointForLocation(ruler unsafe.Pointer, point float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("rulerView:pointForLocation:"), ruler, point)
	return rv
}
// Requests permission for   to add  , an NSRulerMarker being dragged onto the ruler by the user. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:shouldAdd:)
func (v_ View) RulerViewShouldAddMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("rulerView:shouldAddMarker:"), ruler, marker)
	return rv
}
// Requests permission for   to move  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:shouldMove:)
func (v_ View) RulerViewShouldMoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("rulerView:shouldMoveMarker:"), ruler, marker)
	return rv
}
// Requests permission for   to remove  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:shouldRemove:)
func (v_ View) RulerViewShouldRemoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("rulerView:shouldRemoveMarker:"), ruler, marker)
	return rv
}
// Informs the client that   will add the new NSRulerMarker,  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:willAdd:atLocation:)
func (v_ View) RulerViewWillAddMarkerAtLocation(ruler unsafe.Pointer, marker unsafe.Pointer, location float64) float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("rulerView:willAddMarker:atLocation:"), ruler, marker, location)
	return rv
}
// Informs the client that   will move  , an NSRulerMarker already on the ruler view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:willMove:toLocation:)
func (v_ View) RulerViewWillMoveMarkerToLocation(ruler unsafe.Pointer, marker unsafe.Pointer, location float64) float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("rulerView:willMoveMarker:toLocation:"), ruler, marker, location)
	return rv
}
// Informs the client view that   is about to be appropriated by  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:willSetClientView:)
func (v_ View) RulerViewWillSetClientView(ruler unsafe.Pointer, newClient unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("rulerView:willSetClientView:"), ruler, newClient)
}
// Scales the view’s coordinate system so that the unit square scales to the specified dimensions. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/scaleUnitSquare(to:)
func (v_ View) ScaleUnitSquareToSize(newUnitSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("scaleUnitSquareToSize:"), newUnitSize)
}
// Scrolls the view’s closest ancestor   object so a point in the view lies at the origin of the clip view’s bounds rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/scroll(_:)
func (v_ View) ScrollPoint(point unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("scrollPoint:"), point)
}
// Copies the visible portion of the view’s rendered image within a region and lays that portion down again at a specified offset . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/scroll(_:by:)
func (v_ View) ScrollRectBy(rect unsafe.Pointer, delta unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("scrollRect:by:"), rect, delta)
}
// Notifies the superview of a clip view that the clip view needs to reset the origin of its bounds rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/scroll(_:to:)
func (v_ View) ScrollClipViewToPoint(clipView unsafe.Pointer, point unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("scrollClipView:toPoint:"), clipView, point)
}
// Scrolls the view’s closest ancestor   object the minimum distance needed so a specified region of the view becomes visible in the clip view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/scrollToVisible(_:)
func (v_ View) ScrollRectToVisible(rect unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("scrollRectToVisible:"), rect)
	return rv
}
// Sets the origin of the view’s bounds rectangle to a specified point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setBoundsOrigin(_:)
func (v_ View) SetBoundsOrigin(newOrigin unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setBoundsOrigin:"), newOrigin)
}
// Sets the size of the view’s bounds rectangle to specified dimensions, inversely scaling its coordinate system relative to its frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setBoundsSize(_:)
func (v_ View) SetBoundsSize(newSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setBoundsSize:"), newSize)
}
// Sets the priority with which a view resists being made smaller than its intrinsic size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setContentCompressionResistancePriority(_:for:)
func (v_ View) SetContentCompressionResistancePriorityForOrientation(priority unsafe.Pointer, orientation unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setContentCompressionResistancePriority:forOrientation:"), priority, orientation)
}
// Sets the priority with which a view resists being made larger than its intrinsic size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setContentHuggingPriority(_:for:)
func (v_ View) SetContentHuggingPriorityForOrientation(priority unsafe.Pointer, orientation unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setContentHuggingPriority:forOrientation:"), priority, orientation)
}
// Sets the origin of the view’s frame rectangle to the specified point, effectively repositioning it within its superview. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setFrameOrigin(_:)
func (v_ View) SetFrameOrigin(newOrigin unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setFrameOrigin:"), newOrigin)
}
// Sets the size of the view’s frame rectangle to the specified dimensions, resizing it within its superview without affecting its coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setFrameSize(_:)
func (v_ View) SetFrameSize(newSize unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setFrameSize:"), newSize)
}
// Invalidates the area around the focus ring. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setKeyboardFocusRingNeedsDisplay(_:)
func (v_ View) SetKeyboardFocusRingNeedsDisplayInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setKeyboardFocusRingNeedsDisplayInRect:"), rect)
}
// Marks the region of the view within the specified rectangle as needing display, increasing the view’s existing invalid region to include it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setNeedsDisplay(_:)
func (v_ View) SetNeedsDisplayInRect(invalidRect unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setNeedsDisplayInRect:"), invalidRect)
}
// Overridden by subclasses to (re)initialize the view’s graphics state object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setUpGState()
func (v_ View) SetUpGState() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setUpGState"))
}
// Allows the user to drag objects from the view without activating the app or moving the window of the view forward, possibly obscuring the destination. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/shouldDelayWindowOrdering(for:)
func (v_ View) ShouldDelayWindowOrderingForEvent(event unsafe.Pointer) bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("shouldDelayWindowOrderingForEvent:"), event)
	return rv
}
// Returns a Boolean value indicating whether the view is being drawn to an environment that supports color. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/shouldDrawColor()
func (v_ View) ShouldDrawColor() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("shouldDrawColor"))
	return rv
}
// Shows a window displaying the definition of the attributed string at the specified point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/showDefinition(for:at:)
func (v_ View) ShowDefinitionForAttributedStringAtPoint(attrString unsafe.Pointer, textBaselineOrigin unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("showDefinitionForAttributedString:atPoint:"), attrString, textBaselineOrigin)
}
// Shows a window displaying the definition of the specified range of the attributed string. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/showDefinition(for:range:options:baselineOriginProvider:)
func (v_ View) ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString unsafe.Pointer, targetRange unsafe.Pointer, options unsafe.Pointer, originProvider unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("showDefinitionForAttributedString:range:options:baselineOriginProvider:"), attrString, targetRange, options, originProvider)
}
// Orders the view’s immediate subviews using the specified comparator function. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/sortSubviews(_:context:)
func (v_ View) SortSubviewsUsingFunctionContext(compare unsafe.Pointer, context unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("sortSubviewsUsingFunction:context:"), compare, context)
}
// Translates the view’s coordinate system so that its origin moves to a new location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/translateOrigin(to:)
func (v_ View) TranslateOriginToPoint(translation unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("translateOriginToPoint:"), translation)
}
// Translates the display rectangles by the specified delta. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/translateRectsNeedingDisplay(in:by:)
func (v_ View) TranslateRectsNeedingDisplayInRectBy(clipRect unsafe.Pointer, delta unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("translateRectsNeedingDisplayInRect:by:"), clipRect, delta)
}
// Unlocks focus from the current view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/unlockFocus()
func (v_ View) UnlockFocus() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("unlockFocus"))
}
// Unregisters the view as a possible destination in a dragging session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/unregisterDraggedTypes()
func (v_ View) UnregisterDraggedTypes() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("unregisterDraggedTypes"))
}
// Update constraints for the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/updateConstraints()
func (v_ View) UpdateConstraints() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("updateConstraints"))
}
// Updates the constraints for the receiving view and its subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/updateConstraintsForSubtreeIfNeeded()
func (v_ View) UpdateConstraintsForSubtreeIfNeeded() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("updateConstraintsForSubtreeIfNeeded"))
}
// Updates the view’s content by modifying its underlying layer. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/updateLayer()
func (v_ View) UpdateLayer() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("updateLayer"))
}
// Invoked automatically when the view’s geometry changes such that its tracking areas need to be recalculated. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/updateTrackingAreas()
func (v_ View) UpdateTrackingAreas() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("updateTrackingAreas"))
}
// Responds when the view’s backing store properties change. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidChangeBackingProperties()
func (v_ View) ViewDidChangeBackingProperties() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewDidChangeBackingProperties"))
}
// Informs the view that its effective appearance changed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidChangeEffectiveAppearance()
func (v_ View) ViewDidChangeEffectiveAppearance() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewDidChangeEffectiveAppearance"))
}
// Informs the view of the end of a live resize—the user has finished resizing the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidEndLiveResize()
func (v_ View) ViewDidEndLiveResize() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewDidEndLiveResize"))
}
// Invoked when the view is hidden, either directly, or in response to an ancestor being hidden. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidHide()
func (v_ View) ViewDidHide() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewDidHide"))
}
// Informs the view that its superview has changed (possibly to  ). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidMoveToSuperview()
func (v_ View) ViewDidMoveToSuperview() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewDidMoveToSuperview"))
}
// Informs the view that it has been added to a new view hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidMoveToWindow()
func (v_ View) ViewDidMoveToWindow() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewDidMoveToWindow"))
}
// Invoked when the view is unhidden, either directly, or in response to an ancestor being unhidden [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidUnhide()
func (v_ View) ViewDidUnhide() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewDidUnhide"))
}
// Informs the view that it’s required to draw content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewWillDraw()
func (v_ View) ViewWillDraw() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewWillDraw"))
}
// Informs the view that its superview is about to change to the specified superview (which may be  ). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewWillMove(toSuperview:)
func (v_ View) ViewWillMoveToSuperview(newSuperview unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewWillMoveToSuperview:"), newSuperview)
}
// Informs the view that it’s being added to the view hierarchy of the specified window object (which may be  ). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewWillMove(toWindow:)
func (v_ View) ViewWillMoveToWindow(newWindow unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewWillMoveToWindow:"), newWindow)
}
// Informs the view of the start of a live resize—the user has started resizing the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewWillStartLiveResize()
func (v_ View) ViewWillStartLiveResize() {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("viewWillStartLiveResize"))
}
// Returns the view’s nearest descendant (including itself) with a specific tag, or   if no subview has that tag. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewWithTag(_:)
func (v_ View) ViewWithTag(tag int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("viewWithTag:"), tag)
	return rv
}
// Called just before a contextual menu for a view is opened on screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/willOpenMenu(_:with:)
func (v_ View) WillOpenMenuWithEvent(menu unsafe.Pointer, event unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("willOpenMenu:withEvent:"), menu, event)
}
// Overridden by subclasses to perform additional actions before subviews are removed from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/willRemoveSubview(_:)
func (v_ View) WillRemoveSubview(subview unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("willRemoveSubview:"), subview)
}
// Writes EPS data that draws the region of the view within a specified rectangle onto a pasteboard. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/writeEPS(inside:to:)
func (v_ View) WriteEPSInsideRectToPasteboard(rect unsafe.Pointer, pasteboard unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("writeEPSInsideRect:toPasteboard:"), rect, pasteboard)
}
// Writes PDF data that draws the region of the view within a specified rectangle onto a pasteboard. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/writePDF(inside:to:)
func (v_ View) WritePDFInsideRectToPasteboard(rect unsafe.Pointer, pasteboard unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("writePDFInsideRect:toPasteboard:"), rect, pasteboard)
}
// A Boolean value indicating whether the view accepts touch events. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/acceptsTouchEvents
func (v_ View) AcceptsTouchEvents() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("acceptsTouchEvents"))
	return rv
}
// SetAcceptsTouchEvents sets the value of the acceptsTouchEvents property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/acceptsTouchEvents
func (v_ View) SetAcceptsTouchEvents(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setAcceptsTouchEvents:"), value)
}
// Custom insets that you specify to modify your view’s safe area [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/additionalSafeAreaInsets
func (v_ View) AdditionalSafeAreaInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("additionalSafeAreaInsets"))
	return rv
}
// SetAdditionalSafeAreaInsets sets the value of the additionalSafeAreaInsets property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/additionalSafeAreaInsets
func (v_ View) SetAdditionalSafeAreaInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setAdditionalSafeAreaInsets:"), value)
}
// The insets (in points) from the view’s frame that define its content rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/alignmentRectInsets
func (v_ View) AlignmentRectInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("alignmentRectInsets"))
	return rv
}
// The types of touch interactions the view allows. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/allowedTouchTypes
func (v_ View) AllowedTouchTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("allowedTouchTypes"))
	return rv
}
// SetAllowedTouchTypes sets the value of the allowedTouchTypes property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/allowedTouchTypes
func (v_ View) SetAllowedTouchTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setAllowedTouchTypes:"), value)
}
// A Boolean value indicating whether the view ensures it is vibrant on top of other content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/allowsVibrancy
func (v_ View) AllowsVibrancy() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("allowsVibrancy"))
	return rv
}
// The opacity of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/alphaValue
func (v_ View) AlphaValue() float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("alphaValue"))
	return rv
}
// SetAlphaValue sets the value of the alphaValue property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/alphaValue
func (v_ View) SetAlphaValue(value float64) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setAlphaValue:"), value)
}
// A Boolean value indicating whether the view applies the autoresizing behavior to its subviews when its frame size changes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/autoresizesSubviews
func (v_ View) AutoresizesSubviews() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("autoresizesSubviews"))
	return rv
}
// SetAutoresizesSubviews sets the value of the autoresizesSubviews property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/autoresizesSubviews
func (v_ View) SetAutoresizesSubviews(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setAutoresizesSubviews:"), value)
}
// The options that determine how the view is resized relative to its superview. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/autoresizingMask-swift.property
func (v_ View) AutoresizingMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("autoresizingMask"))
	return rv
}
// SetAutoresizingMask sets the value of the autoresizingMask property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/autoresizingMask-swift.property
func (v_ View) SetAutoresizingMask(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setAutoresizingMask:"), value)
}
// An array of Core Image filters to apply to the view’s background. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/backgroundFilters
func (v_ View) BackgroundFilters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("backgroundFilters"))
	return rv
}
// SetBackgroundFilters sets the value of the backgroundFilters property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/backgroundFilters
func (v_ View) SetBackgroundFilters(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setBackgroundFilters:"), value)
}
// The distance (in points) between the bottom of the view’s alignment rectangle and its baseline. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/baselineOffsetFromBottom
func (v_ View) BaselineOffsetFromBottom() float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("baselineOffsetFromBottom"))
	return rv
}
// A layout anchor representing the bottom edge of the view’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/bottomAnchor
func (v_ View) BottomAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("bottomAnchor"))
	return rv
}
// The view’s bounds rectangle, which expresses its location and size in its own coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/bounds
func (v_ View) Bounds() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("bounds"))
	return rv
}
// SetBounds sets the value of the bounds property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/bounds
func (v_ View) SetBounds(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setBounds:"), value)
}
// The angle of rotation, measured in degrees, applied to the view’s bounds rectangle relative to its frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/boundsRotation
func (v_ View) BoundsRotation() float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("boundsRotation"))
	return rv
}
// SetBoundsRotation sets the value of the boundsRotation property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/boundsRotation
func (v_ View) SetBoundsRotation(value float64) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setBoundsRotation:"), value)
}
// A Boolean value indicating whether the view can become key view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/canBecomeKeyView
func (v_ View) CanBecomeKeyView() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("canBecomeKeyView"))
	return rv
}
// A Boolean value indicating whether drawing commands will produce any results. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/canDraw
func (v_ View) CanDraw() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("canDraw"))
	return rv
}
// A Boolean value indicating whether the view can draw its contents on a background thread. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/canDrawConcurrently
func (v_ View) CanDrawConcurrently() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("canDrawConcurrently"))
	return rv
}
// SetCanDrawConcurrently sets the value of the canDrawConcurrently property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/canDrawConcurrently
func (v_ View) SetCanDrawConcurrently(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setCanDrawConcurrently:"), value)
}
// A Boolean value indicating whether the view incorporates content from its subviews into its own layer. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/canDrawSubviewsIntoLayer
func (v_ View) CanDrawSubviewsIntoLayer() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("canDrawSubviewsIntoLayer"))
	return rv
}
// SetCanDrawSubviewsIntoLayer sets the value of the canDrawSubviewsIntoLayer property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/canDrawSubviewsIntoLayer
func (v_ View) SetCanDrawSubviewsIntoLayer(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setCanDrawSubviewsIntoLayer:"), value)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/candidateListTouchBarItem
func (v_ View) CandidateListTouchBarItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("candidateListTouchBarItem"))
	return rv
}
// A layout anchor representing the horizontal center of the view’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/centerXAnchor
func (v_ View) CenterXAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("centerXAnchor"))
	return rv
}
// A layout anchor representing the vertical center of the view’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/centerYAnchor
func (v_ View) CenterYAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("centerYAnchor"))
	return rv
}
// A Boolean value that indicates whether the view, and its subviews, confine their drawing areas to the bounds of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/clipsToBounds
func (v_ View) ClipsToBounds() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("clipsToBounds"))
	return rv
}
// SetClipsToBounds sets the value of the clipsToBounds property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/clipsToBounds
func (v_ View) SetClipsToBounds(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setClipsToBounds:"), value)
}
// The Core Image filter used to composite the view’s contents with its background. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/compositingFilter
func (v_ View) CompositingFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("compositingFilter"))
	return rv
}
// SetCompositingFilter sets the value of the compositingFilter property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/compositingFilter
func (v_ View) SetCompositingFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setCompositingFilter:"), value)
}
// Returns the constraints held by the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/constraints
func (v_ View) Constraints() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("constraints"))
	return rv
}
// An array of Core Image filters to apply to the contents of the view and its sublayers. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/contentFilters
func (v_ View) ContentFilters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("contentFilters"))
	return rv
}
// SetContentFilters sets the value of the contentFilters property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/contentFilters
func (v_ View) SetContentFilters(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setContentFilters:"), value)
}
// The menu item containing the view or any of its superviews in the view hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/enclosingMenuItem
func (v_ View) EnclosingMenuItem() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("enclosingMenuItem"))
	return rv
}
// The nearest ancestor scroll view that contains the current view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/enclosingScrollView
func (v_ View) EnclosingScrollView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("enclosingScrollView"))
	return rv
}
// A layout anchor representing the baseline for the topmost line of text in the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/firstBaselineAnchor
func (v_ View) FirstBaselineAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("firstBaselineAnchor"))
	return rv
}
// The distance (in points) between the top of the view’s alignment rectangle and its topmost baseline. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/firstBaselineOffsetFromTop
func (v_ View) FirstBaselineOffsetFromTop() float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("firstBaselineOffsetFromTop"))
	return rv
}
// The minimum size of the view that satisfies the constraints it holds. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/fittingSize
func (v_ View) FittingSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("fittingSize"))
	return rv
}
// The focus ring mask bounds, specified in the view’s coordinate space. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/focusRingMaskBounds
func (v_ View) FocusRingMaskBounds() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("focusRingMaskBounds"))
	return rv
}
// The type of focus ring drawn around the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/focusRingType
func (v_ View) FocusRingType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("focusRingType"))
	return rv
}
// SetFocusRingType sets the value of the focusRingType property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/focusRingType
func (v_ View) SetFocusRingType(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setFocusRingType:"), value)
}
// The view’s frame rectangle, which defines its position and size in its superview’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/frame
func (v_ View) Frame() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("frame"))
	return rv
}
// SetFrame sets the value of the frame property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/frame
func (v_ View) SetFrame(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setFrame:"), value)
}
// The rotation angle of the view around the center of its layer. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/frameCenterRotation
func (v_ View) FrameCenterRotation() float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("frameCenterRotation"))
	return rv
}
// SetFrameCenterRotation sets the value of the frameCenterRotation property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/frameCenterRotation
func (v_ View) SetFrameCenterRotation(value float64) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setFrameCenterRotation:"), value)
}
// The angle of rotation, measured in degrees, applied to the view’s frame rectangle relative to its superview’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/frameRotation
func (v_ View) FrameRotation() float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("frameRotation"))
	return rv
}
// SetFrameRotation sets the value of the frameRotation property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/frameRotation
func (v_ View) SetFrameRotation(value float64) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setFrameRotation:"), value)
}
// The gesture recognize objects currently attached to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/gestureRecognizers
func (v_ View) GestureRecognizers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("gestureRecognizers"))
	return rv
}
// SetGestureRecognizers sets the value of the gestureRecognizers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/gestureRecognizers
func (v_ View) SetGestureRecognizers(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setGestureRecognizers:"), value)
}
// A Boolean value indicating whether the constraints impacting the layout of the view incompletely specify the location of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/hasAmbiguousLayout
func (v_ View) HasAmbiguousLayout() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("hasAmbiguousLayout"))
	return rv
}
// The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as lines of text from being divided across pages. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/heightAdjustLimit
func (v_ View) HeightAdjustLimit() float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("heightAdjustLimit"))
	return rv
}
// A layout anchor representing the height of the view’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/heightAnchor
func (v_ View) HeightAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("heightAnchor"))
	return rv
}
// A Boolean value indicating whether the view is being rendered as part of a live resizing operation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/inLiveResize
func (v_ View) InLiveResize() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("inLiveResize"))
	return rv
}
// The text input context object for the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/inputContext
func (v_ View) InputContext() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("inputContext"))
	return rv
}
// The natural size for the receiving view, considering only properties of the view itself. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/intrinsicContentSize
func (v_ View) IntrinsicContentSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("intrinsicContentSize"))
	return rv
}
// A Boolean value indicating whether the view or one of its ancestors is being drawn for a find indicator. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isDrawingFindIndicator
func (v_ View) DrawingFindIndicator() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("drawingFindIndicator"))
	return rv
}
// A Boolean value indicating whether the view uses a flipped coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isFlipped
func (v_ View) Flipped() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("flipped"))
	return rv
}
// A Boolean value indicating whether the view is hidden. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isHidden
func (v_ View) Hidden() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("hidden"))
	return rv
}
// SetHidden sets the value of the hidden property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isHidden
func (v_ View) SetHidden(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setHidden:"), value)
}
// A Boolean value indicating whether the view is hidden from sight because it, or one of its ancestors, is marked as hidden. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isHiddenOrHasHiddenAncestor
func (v_ View) HiddenOrHasHiddenAncestor() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("hiddenOrHasHiddenAncestor"))
	return rv
}
// A Boolean value that indicates whether the view’s horizontal size constraints are active. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isHorizontalContentSizeConstraintActive
func (v_ View) HorizontalContentSizeConstraintActive() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("horizontalContentSizeConstraintActive"))
	return rv
}
// SetHorizontalContentSizeConstraintActive sets the value of the horizontalContentSizeConstraintActive property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isHorizontalContentSizeConstraintActive
func (v_ View) SetHorizontalContentSizeConstraintActive(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setHorizontalContentSizeConstraintActive:"), value)
}
// A Boolean value indicating whether the view is in full screen mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isInFullScreenMode
func (v_ View) InFullScreenMode() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("inFullScreenMode"))
	return rv
}
// A Boolean value indicating whether the view fills its frame rectangle with opaque content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isOpaque
func (v_ View) Opaque() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("opaque"))
	return rv
}
// A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isRotatedFromBase
func (v_ View) RotatedFromBase() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("rotatedFromBase"))
	return rv
}
// A Boolean value indicating whether the view or any of its ancestors has ever had a rotation factor applied to its frame or bounds, or has been scaled from the window’s base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isRotatedOrScaledFromBase
func (v_ View) RotatedOrScaledFromBase() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("rotatedOrScaledFromBase"))
	return rv
}
// A Boolean value that indicates whether the view’s vertical size constraints are active. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isVerticalContentSizeConstraintActive
func (v_ View) VerticalContentSizeConstraintActive() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("verticalContentSizeConstraintActive"))
	return rv
}
// SetVerticalContentSizeConstraintActive sets the value of the verticalContentSizeConstraintActive property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isVerticalContentSizeConstraintActive
func (v_ View) SetVerticalContentSizeConstraintActive(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setVerticalContentSizeConstraintActive:"), value)
}
// A layout anchor representing the baseline for the bottommost line of text in the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/lastBaselineAnchor
func (v_ View) LastBaselineAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("lastBaselineAnchor"))
	return rv
}
// The distance (in points) between the bottom of the view’s alignment rectangle and its bottommost baseline. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/lastBaselineOffsetFromBottom
func (v_ View) LastBaselineOffsetFromBottom() float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("lastBaselineOffsetFromBottom"))
	return rv
}
// The Core Animation layer that the view uses as its backing store. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layer
func (v_ View) Layer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("layer"))
	return rv
}
// SetLayer sets the value of the layer property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layer
func (v_ View) SetLayer(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setLayer:"), value)
}
// The current layer contents placement policy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layerContentsPlacement-swift.property
func (v_ View) LayerContentsPlacement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("layerContentsPlacement"))
	return rv
}
// SetLayerContentsPlacement sets the value of the layerContentsPlacement property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layerContentsPlacement-swift.property
func (v_ View) SetLayerContentsPlacement(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setLayerContentsPlacement:"), value)
}
// The contents redraw policy for the view’s layer. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layerContentsRedrawPolicy-swift.property
func (v_ View) LayerContentsRedrawPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("layerContentsRedrawPolicy"))
	return rv
}
// SetLayerContentsRedrawPolicy sets the value of the layerContentsRedrawPolicy property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layerContentsRedrawPolicy-swift.property
func (v_ View) SetLayerContentsRedrawPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setLayerContentsRedrawPolicy:"), value)
}
// A Boolean value indicating whether the view’s layer uses Core Image filters and needs in-process rendering. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layerUsesCoreImageFilters
func (v_ View) LayerUsesCoreImageFilters() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("layerUsesCoreImageFilters"))
	return rv
}
// SetLayerUsesCoreImageFilters sets the value of the layerUsesCoreImageFilters property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layerUsesCoreImageFilters
func (v_ View) SetLayerUsesCoreImageFilters(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setLayerUsesCoreImageFilters:"), value)
}
// The array of layout guide objects owned by this view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layoutGuides
func (v_ View) LayoutGuides() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("layoutGuides"))
	return rv
}
// A layout guide that provides the recommended amount of padding for content inside of a view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layoutMarginsGuide
func (v_ View) LayoutMarginsGuide() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("layoutMarginsGuide"))
	return rv
}
// A layout anchor representing the leading edge of the view’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/leadingAnchor
func (v_ View) LeadingAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("leadingAnchor"))
	return rv
}
// A layout anchor representing the left edge of the view’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/leftAnchor
func (v_ View) LeftAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("leftAnchor"))
	return rv
}
// A Boolean value indicating whether the view can pass mouse down events through to its superviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/mouseDownCanMoveWindow
func (v_ View) MouseDownCanMoveWindow() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("mouseDownCanMoveWindow"))
	return rv
}
// A Boolean value that determines whether the view needs to be redrawn before being displayed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/needsDisplay
func (v_ View) NeedsDisplay() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("needsDisplay"))
	return rv
}
// SetNeedsDisplay sets the value of the needsDisplay property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/needsDisplay
func (v_ View) SetNeedsDisplay(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setNeedsDisplay:"), value)
}
// A Boolean value indicating whether the view needs a layout pass before it can be drawn. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/needsLayout
func (v_ View) NeedsLayout() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("needsLayout"))
	return rv
}
// SetNeedsLayout sets the value of the needsLayout property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/needsLayout
func (v_ View) SetNeedsLayout(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setNeedsLayout:"), value)
}
// A Boolean value indicating whether the view needs its panel to become the key window before it can handle keyboard input and navigation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/needsPanelToBecomeKey
func (v_ View) NeedsPanelToBecomeKey() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("needsPanelToBecomeKey"))
	return rv
}
// A Boolean value indicating whether the view’s constraints need to be updated. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/needsUpdateConstraints
func (v_ View) NeedsUpdateConstraints() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("needsUpdateConstraints"))
	return rv
}
// SetNeedsUpdateConstraints sets the value of the needsUpdateConstraints property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/needsUpdateConstraints
func (v_ View) SetNeedsUpdateConstraints(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setNeedsUpdateConstraints:"), value)
}
// The view object that follows the current view in the key view loop. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/nextKeyView
func (v_ View) NextKeyView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("nextKeyView"))
	return rv
}
// SetNextKeyView sets the value of the nextKeyView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/nextKeyView
func (v_ View) SetNextKeyView(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setNextKeyView:"), value)
}
// The closest view object in the key view loop that follows the current view in the key view loop and accepts first responder status. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/nextValidKeyView
func (v_ View) NextValidKeyView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("nextValidKeyView"))
	return rv
}
// The view’s closest opaque ancestor, which might be the view itself. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/opaqueAncestor
func (v_ View) OpaqueAncestor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("opaqueAncestor"))
	return rv
}
// A default footer string that includes the current page number and page count. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/pageFooter
func (v_ View) PageFooter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("pageFooter"))
	return rv
}
// A default header string that includes the print job title and date. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/pageHeader
func (v_ View) PageHeader() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("pageHeader"))
	return rv
}
// A Boolean value indicating whether the view posts notifications when its bounds rectangle changes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/postsBoundsChangedNotifications
func (v_ View) PostsBoundsChangedNotifications() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("postsBoundsChangedNotifications"))
	return rv
}
// SetPostsBoundsChangedNotifications sets the value of the postsBoundsChangedNotifications property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/postsBoundsChangedNotifications
func (v_ View) SetPostsBoundsChangedNotifications(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setPostsBoundsChangedNotifications:"), value)
}
// A Boolean value indicating whether the view posts notifications when its frame rectangle changes. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/postsFrameChangedNotifications
func (v_ View) PostsFrameChangedNotifications() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("postsFrameChangedNotifications"))
	return rv
}
// SetPostsFrameChangedNotifications sets the value of the postsFrameChangedNotifications property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/postsFrameChangedNotifications
func (v_ View) SetPostsFrameChangedNotifications(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setPostsFrameChangedNotifications:"), value)
}
// When this property is true, any NSControls in the view or its descendants will be sized with compact   metrics compatible with macOS 15 and earlier.   Defaults to false [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/prefersCompactControlSizeMetrics
func (v_ View) PrefersCompactControlSizeMetrics() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("prefersCompactControlSizeMetrics"))
	return rv
}
// SetPrefersCompactControlSizeMetrics sets the value of the prefersCompactControlSizeMetrics property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/prefersCompactControlSizeMetrics
func (v_ View) SetPrefersCompactControlSizeMetrics(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setPrefersCompactControlSizeMetrics:"), value)
}
// The portion of the view that has been rendered and is available for responsive scrolling. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/preparedContentRect
func (v_ View) PreparedContentRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("preparedContentRect"))
	return rv
}
// SetPreparedContentRect sets the value of the preparedContentRect property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/preparedContentRect
func (v_ View) SetPreparedContentRect(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setPreparedContentRect:"), value)
}
// A Boolean value indicating whether the view optimizes live-resize operations by preserving content that has not moved. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/preservesContentDuringLiveResize
func (v_ View) PreservesContentDuringLiveResize() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("preservesContentDuringLiveResize"))
	return rv
}
// Configures the behavior and progression of the Force Touch trackpad when responding to touch input produced by the user when the cursor is over the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/pressureConfiguration
func (v_ View) PressureConfiguration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("pressureConfiguration"))
	return rv
}
// SetPressureConfiguration sets the value of the pressureConfiguration property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/pressureConfiguration
func (v_ View) SetPressureConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setPressureConfiguration:"), value)
}
// The view object preceding the current view in the key view loop. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/previousKeyView
func (v_ View) PreviousKeyView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("previousKeyView"))
	return rv
}
// The closest view object in the key view loop that precedes the current view and accepts first responder status. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/previousValidKeyView
func (v_ View) PreviousValidKeyView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("previousValidKeyView"))
	return rv
}
// The view’s print job title. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/printJobTitle
func (v_ View) PrintJobTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("printJobTitle"))
	return rv
}
// The rectangle identifying the portion of your view that did not change during a live resize operation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rectPreservedDuringLiveResize
func (v_ View) RectPreservedDuringLiveResize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("rectPreservedDuringLiveResize"))
	return rv
}
// The array of pasteboard drag types that the view can accept. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/registeredDraggedTypes
func (v_ View) RegisteredDraggedTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("registeredDraggedTypes"))
	return rv
}
// A layout anchor representing the right edge of the view’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rightAnchor
func (v_ View) RightAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("rightAnchor"))
	return rv
}
// The distances from the edges of your view that define the safe area. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/safeAreaInsets
func (v_ View) SafeAreaInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("safeAreaInsets"))
	return rv
}
// The layout guide you use to position content inside your view’s safe area. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/safeAreaLayoutGuide
func (v_ View) SafeAreaLayoutGuide() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("safeAreaLayoutGuide"))
	return rv
}
// A rectangle in the view’s coordinate system that contains the unobscured portion of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/safeAreaRect
func (v_ View) SafeAreaRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("safeAreaRect"))
	return rv
}
// The shadow displayed underneath the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/shadow
func (v_ View) Shadow() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("shadow"))
	return rv
}
// SetShadow sets the value of the shadow property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/shadow
func (v_ View) SetShadow(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setShadow:"), value)
}
// The array of views embedded in the current view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/subviews
func (v_ View) Subviews() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("subviews"))
	return rv
}
// SetSubviews sets the value of the subviews property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/subviews
func (v_ View) SetSubviews(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setSubviews:"), value)
}
// The view that is the parent of the current view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/superview
func (v_ View) Superview() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("superview"))
	return rv
}
// The view’s tag, which is an integer that you use to identify the view within your app. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/tag
func (v_ View) Tag() int {
	rv := objc.Send[int](v_.ID(), objc.RegisterName("tag"))
	return rv
}
// The text for the view’s tooltip. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/toolTip
func (v_ View) ToolTip() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("toolTip"))
	return rv
}
// SetToolTip sets the value of the toolTip property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/toolTip
func (v_ View) SetToolTip(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setToolTip:"), value)
}
// A layout anchor representing the top edge of the view’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/topAnchor
func (v_ View) TopAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("topAnchor"))
	return rv
}
// An array of the view’s tracking areas. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/trackingAreas
func (v_ View) TrackingAreas() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("trackingAreas"))
	return rv
}
// A layout anchor representing the trailing edge of the view’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/trailingAnchor
func (v_ View) TrailingAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("trailingAnchor"))
	return rv
}
// A Boolean value indicating whether the view’s autoresizing mask is translated into constraints for the constraint-based layout system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/translatesAutoresizingMaskIntoConstraints
func (v_ View) TranslatesAutoresizingMaskIntoConstraints() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("translatesAutoresizingMaskIntoConstraints"))
	return rv
}
// SetTranslatesAutoresizingMaskIntoConstraints sets the value of the translatesAutoresizingMaskIntoConstraints property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/translatesAutoresizingMaskIntoConstraints
func (v_ View) SetTranslatesAutoresizingMaskIntoConstraints(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setTranslatesAutoresizingMaskIntoConstraints:"), value)
}
// The layout direction for content in the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/userInterfaceLayoutDirection
func (v_ View) UserInterfaceLayoutDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("userInterfaceLayoutDirection"))
	return rv
}
// SetUserInterfaceLayoutDirection sets the value of the userInterfaceLayoutDirection property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/userInterfaceLayoutDirection
func (v_ View) SetUserInterfaceLayoutDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setUserInterfaceLayoutDirection:"), value)
}
// The portion of the view that isn’t clipped by its superviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/visibleRect
func (v_ View) VisibleRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("visibleRect"))
	return rv
}
// A Boolean value indicating whether the view wants an OpenGL backing surface with a resolution greater than 1 pixel per point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/wantsBestResolutionOpenGLSurface
func (v_ View) WantsBestResolutionOpenGLSurface() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("wantsBestResolutionOpenGLSurface"))
	return rv
}
// SetWantsBestResolutionOpenGLSurface sets the value of the wantsBestResolutionOpenGLSurface property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/wantsBestResolutionOpenGLSurface
func (v_ View) SetWantsBestResolutionOpenGLSurface(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setWantsBestResolutionOpenGLSurface:"), value)
}
// A Boolean value indicating whether AppKit’s default clipping behavior is in effect. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/wantsDefaultClipping
func (v_ View) WantsDefaultClipping() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("wantsDefaultClipping"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/wantsExtendedDynamicRangeOpenGLSurface
func (v_ View) WantsExtendedDynamicRangeOpenGLSurface() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("wantsExtendedDynamicRangeOpenGLSurface"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/wantsExtendedDynamicRangeOpenGLSurface
func (v_ View) SetWantsExtendedDynamicRangeOpenGLSurface(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setWantsExtendedDynamicRangeOpenGLSurface:"), value)
}
// A Boolean value indicating whether the view uses a layer as its backing store. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/wantsLayer
func (v_ View) WantsLayer() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("wantsLayer"))
	return rv
}
// SetWantsLayer sets the value of the wantsLayer property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/wantsLayer
func (v_ View) SetWantsLayer(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setWantsLayer:"), value)
}
// A Boolean value indicating whether the view wants resting touches. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/wantsRestingTouches
func (v_ View) WantsRestingTouches() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("wantsRestingTouches"))
	return rv
}
// SetWantsRestingTouches sets the value of the wantsRestingTouches property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/wantsRestingTouches
func (v_ View) SetWantsRestingTouches(value bool) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setWantsRestingTouches:"), value)
}
// A Boolean value indicating which drawing path the view takes when updating its contents. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/wantsUpdateLayer
func (v_ View) WantsUpdateLayer() bool {
	rv := objc.Send[bool](v_.ID(), objc.RegisterName("wantsUpdateLayer"))
	return rv
}
// The fraction of the page that can be pushed onto the next page during automatic pagination to prevent items such as small images or text columns from being divided across pages. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/widthAdjustLimit
func (v_ View) WidthAdjustLimit() float64 {
	rv := objc.Send[float64](v_.ID(), objc.RegisterName("widthAdjustLimit"))
	return rv
}
// A layout anchor representing the width of the view’s frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/widthAnchor
func (v_ View) WidthAnchor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("widthAnchor"))
	return rv
}
// The view’s window object, if it is installed in a window. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/window
func (v_ View) Window() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("window"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/writingToolsCoordinator
func (v_ View) WritingToolsCoordinator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID(), objc.RegisterName("writingToolsCoordinator"))
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/writingToolsCoordinator
func (v_ View) SetWritingToolsCoordinator(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID(), objc.RegisterName("setWritingToolsCoordinator:"), value)
}
