// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [View] class.
var ViewClass objc.Class

func init() {
	ViewClass = objc.GetClass("NSView")
}

type View struct {
	objc.ID
}

func ViewFrom(ptr unsafe.Pointer) View {
	return View{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (vc View) Alloc() View {
	ret := objc.ID(ViewClass).Send(objc.RegisterName("alloc"))
	return View{ret}
}

// Init initializes the instance.
func (v_ View) Init() View {
	ret := v_.ID.Send(objc.RegisterName("init"))
	return View{ret}
}
// Initializes a view using from data in the specified coder object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/init(coder:)
func NewViewWithCoder(coder unsafe.Pointer) View {
	instance := View{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = View{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
// Initializes and returns a newly allocated   object with a specified frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/init(frame:)
func NewViewWithFrame(frameRect unsafe.Pointer) View {
	instance := View{}.Alloc()
	sel := objc.RegisterName("initWithFrame:")
	ret := instance.ID.Send(sel, frameRect)
	instance = View{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Overridden by subclasses to return   if the view should be sent a   message for an initial mouse-down event,   if not. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/acceptsFirstMouse(for:)
func (v_ View) AcceptsFirstMouse(event unsafe.Pointer) bool {
	sel := objc.RegisterName("acceptsFirstMouse:")
	ret := v_.ID.Send(sel, event)
	return ret != 0
}
// Adds a constraint on the layout of the receiving view or its subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addConstraint(_:)
func (v_ View) AddConstraint(constraint unsafe.Pointer) {
	sel := objc.RegisterName("addConstraint:")
	v_.ID.Send(sel, constraint)
}
// Adds multiple constraints on the layout of the receiving view or its subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addConstraints(_:)
func (v_ View) AddConstraints(constraints unsafe.Pointer) {
	sel := objc.RegisterName("addConstraints:")
	v_.ID.Send(sel, constraints)
}
// Establishes  the cursor to be used when the mouse pointer lies within a specified region. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addCursorRect(_:cursor:)
func (v_ View) AddCursorRectCursor(rect unsafe.Pointer, object unsafe.Pointer) {
	sel := objc.RegisterName("addCursorRect:cursor:")
	v_.ID.Send(sel, rect, object)
}
// Attaches a gesture recognizer to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addGestureRecognizer(_:)
func (v_ View) AddGestureRecognizer(gestureRecognizer unsafe.Pointer) {
	sel := objc.RegisterName("addGestureRecognizer:")
	v_.ID.Send(sel, gestureRecognizer)
}
// Adds the provided layout guide to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addLayoutGuide(_:)
func (v_ View) AddLayoutGuide(guide unsafe.Pointer) {
	sel := objc.RegisterName("addLayoutGuide:")
	v_.ID.Send(sel, guide)
}
// Adds a view to the view’s subviews so it’s displayed above its siblings. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addSubview(_:)
func (v_ View) AddSubview(view unsafe.Pointer) {
	sel := objc.RegisterName("addSubview:")
	v_.ID.Send(sel, view)
}
// Inserts a view among the view’s subviews so it’s displayed immediately above or below another view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addSubview(_:positioned:relativeTo:)
func (v_ View) AddSubviewPositionedRelativeTo(view unsafe.Pointer, place unsafe.Pointer, otherView unsafe.Pointer) {
	sel := objc.RegisterName("addSubview:positioned:relativeTo:")
	v_.ID.Send(sel, view, place, otherView)
}
// Creates a tooltip for a defined area in the view and returns a tag that identifies the tooltip rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addToolTip(_:owner:userData:)
func (v_ View) AddToolTipRectOwnerUserData(rect unsafe.Pointer, owner objc.ID, data unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("addToolTipRect:owner:userData:")
	ret := v_.ID.Send(sel, rect, owner, data)
	return unsafe.Pointer(ret)
}
// Adds a given tracking area to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addTrackingArea(_:)
func (v_ View) AddTrackingArea(trackingArea unsafe.Pointer) {
	sel := objc.RegisterName("addTrackingArea:")
	v_.ID.Send(sel, trackingArea)
}
// Establishes  an area for tracking mouse-entered and mouse-exited events within the view and returns a tag that identifies the tracking rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/addTrackingRect(_:owner:userData:assumeInside:)
func (v_ View) AddTrackingRectOwnerUserDataAssumeInside(rect unsafe.Pointer, owner objc.ID, data unsafe.Pointer, flag bool) unsafe.Pointer {
	sel := objc.RegisterName("addTrackingRect:owner:userData:assumeInside:")
	ret := v_.ID.Send(sel, rect, owner, data, flag)
	return unsafe.Pointer(ret)
}
// Overridden by subclasses to adjust page height during automatic pagination. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/adjustPageHeightNew(_:top:bottom:limit:)
func (v_ View) AdjustPageHeightNewTopBottomLimit(newBottom float64, oldTop float64, oldBottom float64, bottomLimit float64) {
	sel := objc.RegisterName("adjustPageHeightNew:top:bottom:limit:")
	v_.ID.Send(sel, newBottom, oldTop, oldBottom, bottomLimit)
}
// Overridden by subclasses to adjust page width during automatic pagination. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/adjustPageWidthNew(_:left:right:limit:)
func (v_ View) AdjustPageWidthNewLeftRightLimit(newRight float64, oldLeft float64, oldRight float64, rightLimit float64) {
	sel := objc.RegisterName("adjustPageWidthNew:left:right:limit:")
	v_.ID.Send(sel, newRight, oldLeft, oldRight, rightLimit)
}
// Overridden by subclasses to modify a given rectangle, returning the altered rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/adjustScroll(_:)
func (v_ View) AdjustScroll(newVisible unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("adjustScroll:")
	ret := v_.ID.Send(sel, newVisible)
	return unsafe.Pointer(ret)
}
// Returns the view’s alignment rectangle for a given frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/alignmentRect(forFrame:)
func (v_ View) AlignmentRectForFrame(frame unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("alignmentRectForFrame:")
	ret := v_.ID.Send(sel, frame)
	return unsafe.Pointer(ret)
}
// Causes the view to maintain a private graphics state object, which encapsulates all parameters of the graphics environment. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/allocateGState()
func (v_ View) AllocateGState() {
	sel := objc.RegisterName("allocateGState")
	v_.ID.Send(sel)
}
// Returns the closest ancestor shared by the view and another specified view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/ancestorShared(with:)
func (v_ View) AncestorSharedWithView(view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("ancestorSharedWithView:")
	ret := v_.ID.Send(sel, view)
	return unsafe.Pointer(ret)
}
// Scrolls the view’s closest ancestor   object proportionally to the distance of an event that occurs outside of it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/autoscroll(with:)
func (v_ View) Autoscroll(event unsafe.Pointer) bool {
	sel := objc.RegisterName("autoscroll:")
	ret := v_.ID.Send(sel, event)
	return ret != 0
}
// Returns a backing store pixel-aligned rectangle in local view coordinates. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/backingAlignedRect(_:options:)
func (v_ View) BackingAlignedRectOptions(rect unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("backingAlignedRect:options:")
	ret := v_.ID.Send(sel, rect, options)
	return unsafe.Pointer(ret)
}
// Invoked at the beginning of the printing session, this method sets up the current graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/beginDocument()
func (v_ View) BeginDocument() {
	sel := objc.RegisterName("beginDocument")
	v_.ID.Send(sel)
}
// Initiates a dragging session with a group of dragging items. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/beginDraggingSession(with:event:source:)
func (v_ View) BeginDraggingSessionWithItemsEventSource(items unsafe.Pointer, event unsafe.Pointer, source unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("beginDraggingSessionWithItems:event:source:")
	ret := v_.ID.Send(sel, items, event, source)
	return unsafe.Pointer(ret)
}
// Called at the beginning of each page, this method sets up the coordinate system so that a region inside the view’s bounds is translated to a specified location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/beginPage(in:atPlacement:)
func (v_ View) BeginPageInRectAtPlacement(rect unsafe.Pointer, location unsafe.Pointer) {
	sel := objc.RegisterName("beginPageInRect:atPlacement:")
	v_.ID.Send(sel, rect, location)
}
// Returns a bitmap-representation object suitable for caching the specified portion of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/bitmapImageRepForCachingDisplay(in:)
func (v_ View) BitmapImageRepForCachingDisplayInRect(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("bitmapImageRepForCachingDisplayInRect:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Draws the specified area of the view, and its descendants, into a provided bitmap-representation object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/cacheDisplay(in:to:)
func (v_ View) CacheDisplayInRectToBitmapImageRep(rect unsafe.Pointer, bitmapImageRep unsafe.Pointer) {
	sel := objc.RegisterName("cacheDisplayInRect:toBitmapImageRep:")
	v_.ID.Send(sel, rect, bitmapImageRep)
}
// Converts the corners of a specified rectangle to lie on the center of device pixels, which is useful in compensating for rendering overscanning when the coordinate system has been scaled. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/centerScanRect(_:)
func (v_ View) CenterScanRect(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("centerScanRect:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Returns the constraints impacting the layout of the view for a given orientation. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/constraintsAffectingLayout(for:)
func (v_ View) ConstraintsAffectingLayoutForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("constraintsAffectingLayoutForOrientation:")
	ret := v_.ID.Send(sel, orientation)
	return unsafe.Pointer(ret)
}
// Returns the priority with which a view resists being made smaller than its intrinsic size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/contentCompressionResistancePriority(for:)
func (v_ View) ContentCompressionResistancePriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("contentCompressionResistancePriorityForOrientation:")
	ret := v_.ID.Send(sel, orientation)
	return unsafe.Pointer(ret)
}
// Returns the priority with which a view resists being made larger than its intrinsic size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/contentHuggingPriority(for:)
func (v_ View) ContentHuggingPriorityForOrientation(orientation unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("contentHuggingPriorityForOrientation:")
	ret := v_.ID.Send(sel, orientation)
	return unsafe.Pointer(ret)
}
// Converts a point from the coordinate system of a given view to that of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:from:)-1dq9l
func (v_ View) ConvertPointFromView(point unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPoint:fromView:")
	ret := v_.ID.Send(sel, point, view)
	return unsafe.Pointer(ret)
}
// Converts a size from another view’s coordinate system to that of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:from:)-40x0w
func (v_ View) ConvertSizeFromView(size unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertSize:fromView:")
	ret := v_.ID.Send(sel, size, view)
	return unsafe.Pointer(ret)
}
// Converts a rectangle from the coordinate system of another view to that of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:from:)-7fbb6
func (v_ View) ConvertRectFromView(rect unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRect:fromView:")
	ret := v_.ID.Send(sel, rect, view)
	return unsafe.Pointer(ret)
}
// Converts a rectangle from the view’s coordinate system to that of another view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:to:)-3cqqt
func (v_ View) ConvertRectToView(rect unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRect:toView:")
	ret := v_.ID.Send(sel, rect, view)
	return unsafe.Pointer(ret)
}
// Converts a size from the view’s coordinate system to that of another view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:to:)-5nptx
func (v_ View) ConvertSizeToView(size unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertSize:toView:")
	ret := v_.ID.Send(sel, size, view)
	return unsafe.Pointer(ret)
}
// Converts a point from the view’s coordinate system to that of a given view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convert(_:to:)-6u9ir
func (v_ View) ConvertPointToView(point unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPoint:toView:")
	ret := v_.ID.Send(sel, point, view)
	return unsafe.Pointer(ret)
}
// Converts a point from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromBacking(_:)-229ps
func (v_ View) ConvertPointFromBacking(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPointFromBacking:")
	ret := v_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Converts a rectangle from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromBacking(_:)-2njpa
func (v_ View) ConvertRectFromBacking(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRectFromBacking:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Converts a size from its pixel aligned backing store coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromBacking(_:)-4agf9
func (v_ View) ConvertSizeFromBacking(size unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertSizeFromBacking:")
	ret := v_.ID.Send(sel, size)
	return unsafe.Pointer(ret)
}
// Convert the point from the layer’s interior coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromLayer(_:)-3nsbu
func (v_ View) ConvertPointFromLayer(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPointFromLayer:")
	ret := v_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Convert the size from the layer’s interior coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromLayer(_:)-3usqp
func (v_ View) ConvertSizeFromLayer(size unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertSizeFromLayer:")
	ret := v_.ID.Send(sel, size)
	return unsafe.Pointer(ret)
}
// Convert the rectangle from the layer’s interior coordinate system to the view’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertFromLayer(_:)-8s5bi
func (v_ View) ConvertRectFromLayer(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRectFromLayer:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Converts the point from the base coordinate system to the view’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertPointFromBase:
func (v_ View) ConvertPointFromBase(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPointFromBase:")
	ret := v_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Converts the point from the view’s coordinate system to the base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertPointToBase:
func (v_ View) ConvertPointToBase(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPointToBase:")
	ret := v_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Converts the rectangle from the base coordinate system to the view’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertRectFromBase:
func (v_ View) ConvertRectFromBase(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRectFromBase:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Converts the rectangle from the view’s coordinate system to the base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertRectToBase:
func (v_ View) ConvertRectToBase(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRectToBase:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Converts the size from the base coordinate system to the view’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertSizeFromBase:
func (v_ View) ConvertSizeFromBase(size unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertSizeFromBase:")
	ret := v_.ID.Send(sel, size)
	return unsafe.Pointer(ret)
}
// Converts the size from the view’s coordinate system to the base coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertSizeToBase:
func (v_ View) ConvertSizeToBase(size unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertSizeToBase:")
	ret := v_.ID.Send(sel, size)
	return unsafe.Pointer(ret)
}
// Converts a point from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToBacking(_:)-2xx45
func (v_ View) ConvertPointToBacking(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPointToBacking:")
	ret := v_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Converts a rectangle from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToBacking(_:)-3zors
func (v_ View) ConvertRectToBacking(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRectToBacking:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Converts a size from the view’s interior coordinate system to its pixel aligned backing store coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToBacking(_:)-4ra9y
func (v_ View) ConvertSizeToBacking(size unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertSizeToBacking:")
	ret := v_.ID.Send(sel, size)
	return unsafe.Pointer(ret)
}
// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToLayer(_:)-160pw
func (v_ View) ConvertRectToLayer(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertRectToLayer:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToLayer(_:)-2vozx
func (v_ View) ConvertSizeToLayer(size unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertSizeToLayer:")
	ret := v_.ID.Send(sel, size)
	return unsafe.Pointer(ret)
}
// Convert the size from the view’s interior coordinate system to the layer’s interior coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/convertToLayer(_:)-44u7d
func (v_ View) ConvertPointToLayer(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("convertPointToLayer:")
	ret := v_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Returns EPS data that draws the region of the view within a specified rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/dataWithEPS(inside:)
func (v_ View) DataWithEPSInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataWithEPSInsideRect:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Returns PDF data that draws the region of the view within a specified rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/dataWithPDF(inside:)
func (v_ View) DataWithPDFInsideRect(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("dataWithPDFInsideRect:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Overridden by subclasses to perform additional actions when subviews are added to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/didAddSubview(_:)
func (v_ View) DidAddSubview(subview unsafe.Pointer) {
	sel := objc.RegisterName("didAddSubview:")
	v_.ID.Send(sel, subview)
}
// Called after a contextual menu that was displayed from the receiving view has been closed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/didCloseMenu(_:with:)
func (v_ View) DidCloseMenuWithEvent(menu unsafe.Pointer, event unsafe.Pointer) {
	sel := objc.RegisterName("didCloseMenu:withEvent:")
	v_.ID.Send(sel, menu, event)
}
// Invalidates all cursor rectangles set up using  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/discardCursorRects()
func (v_ View) DiscardCursorRects() {
	sel := objc.RegisterName("discardCursorRects")
	v_.ID.Send(sel)
}
// Displays the view and all its subviews if possible, invoking each of the   methods  ,  , and   as necessary. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/display()
func (v_ View) Display() {
	sel := objc.RegisterName("display")
	v_.ID.Send(sel)
}
// Acts as  , but confining drawing to a rectangular region of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/display(_:)
func (v_ View) DisplayRect(rect unsafe.Pointer) {
	sel := objc.RegisterName("displayRect:")
	v_.ID.Send(sel, rect)
}
// Displays the view and all its subviews if any part of the view has been marked as needing display. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeeded()
func (v_ View) DisplayIfNeeded() {
	sel := objc.RegisterName("displayIfNeeded")
	v_.ID.Send(sel)
}
// Acts as  , confining drawing to a specified region of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeeded(_:)
func (v_ View) DisplayIfNeededInRect(rect unsafe.Pointer) {
	sel := objc.RegisterName("displayIfNeededInRect:")
	v_.ID.Send(sel, rect)
}
// Acts as  , except that this method doesn’t back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity()
func (v_ View) DisplayIfNeededIgnoringOpacity() {
	sel := objc.RegisterName("displayIfNeededIgnoringOpacity")
	v_.ID.Send(sel)
}
// Acts as  , but confining drawing to   and not backing up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIfNeededIgnoringOpacity(_:)
func (v_ View) DisplayIfNeededInRectIgnoringOpacity(rect unsafe.Pointer) {
	sel := objc.RegisterName("displayIfNeededInRectIgnoringOpacity:")
	v_.ID.Send(sel, rect)
}
// Displays the view but confines drawing to a specified region and does not back up to the first opaque ancestor—it simply causes the view and its descendants to execute their drawing code. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIgnoringOpacity(_:)
func (v_ View) DisplayRectIgnoringOpacity(rect unsafe.Pointer) {
	sel := objc.RegisterName("displayRectIgnoringOpacity:")
	v_.ID.Send(sel, rect)
}
// Causes the view and its descendants to be redrawn to the specified graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayIgnoringOpacity(_:in:)
func (v_ View) DisplayRectIgnoringOpacityInContext(rect unsafe.Pointer, context unsafe.Pointer) {
	sel := objc.RegisterName("displayRectIgnoringOpacity:inContext:")
	v_.ID.Send(sel, rect, context)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/displayLink(target:selector:)
func (v_ View) DisplayLinkWithTargetSelector(target objc.ID, selector objc.SEL) unsafe.Pointer {
	sel := objc.RegisterName("displayLinkWithTarget:selector:")
	ret := v_.ID.Send(sel, target, selector)
	return unsafe.Pointer(ret)
}
// Initiates a dragging operation from the view, allowing the user to drag a file icon to any application that has window or view objects that accept files. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/dragFile(_:from:slideBack:event:)
func (v_ View) DragFileFromRectSlideBackEvent(filename unsafe.Pointer, rect unsafe.Pointer, flag bool, event unsafe.Pointer) bool {
	sel := objc.RegisterName("dragFile:fromRect:slideBack:event:")
	ret := v_.ID.Send(sel, filename, rect, flag, event)
	return ret != 0
}
// Initiates a dragging operation from the view, allowing the user to drag arbitrary data with a specified icon into any application that has window or view objects that accept dragged data. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/dragImage:at:offset:event:pasteboard:source:slideBack:
func (v_ View) DragImageAtOffsetEventPasteboardSourceSlideBack(image unsafe.Pointer, viewLocation unsafe.Pointer, initialOffset unsafe.Pointer, event unsafe.Pointer, pboard unsafe.Pointer, sourceObj objc.ID, slideFlag bool) {
	sel := objc.RegisterName("dragImage:at:offset:event:pasteboard:source:slideBack:")
	v_.ID.Send(sel, image, viewLocation, initialOffset, event, pboard, sourceObj, slideFlag)
}
// Initiates a dragging operation from the view, allowing the user to drag one or more promised files (or directories) into any application that has window or view objects that accept promised file data. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/dragPromisedFiles(ofTypes:from:source:slideBack:event:)
func (v_ View) DragPromisedFilesOfTypesFromRectSourceSlideBackEvent(typeArray unsafe.Pointer, rect unsafe.Pointer, sourceObject objc.ID, flag bool, event unsafe.Pointer) bool {
	sel := objc.RegisterName("dragPromisedFilesOfTypes:fromRect:source:slideBack:event:")
	ret := v_.ID.Send(sel, typeArray, rect, sourceObject, flag, event)
	return ret != 0
}
// Overridden by subclasses to draw the view’s image within the specified rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/draw(_:)
func (v_ View) DrawRect(dirtyRect unsafe.Pointer) {
	sel := objc.RegisterName("drawRect:")
	v_.ID.Send(sel, dirtyRect)
}
// Draws the focus ring mask for the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/drawFocusRingMask()
func (v_ View) DrawFocusRingMask() {
	sel := objc.RegisterName("drawFocusRingMask")
	v_.ID.Send(sel)
}
// Allows applications that use the AppKit pagination facility to draw additional marks on each logical page. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/drawPageBorder(with:)
func (v_ View) DrawPageBorderWithSize(borderSize unsafe.Pointer) {
	sel := objc.RegisterName("drawPageBorderWithSize:")
	v_.ID.Send(sel, borderSize)
}
// Allows applications that use the AppKit pagination facility to draw additional marks on each printed sheet. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/drawSheetBorder(with:)
func (v_ View) DrawSheetBorderWithSize(borderSize unsafe.Pointer) {
	sel := objc.RegisterName("drawSheetBorderWithSize:")
	v_.ID.Send(sel, borderSize)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/edgeInsetsForLayoutRegion:
func (v_ View) EdgeInsetsForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("edgeInsetsForLayoutRegion:")
	ret := v_.ID.Send(sel, layoutRegion)
	return unsafe.Pointer(ret)
}
// This method is invoked at the end of the printing session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/endDocument()
func (v_ View) EndDocument() {
	sel := objc.RegisterName("endDocument")
	v_.ID.Send(sel)
}
// Writes the end of a conforming page. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/endPage()
func (v_ View) EndPage() {
	sel := objc.RegisterName("endPage")
	v_.ID.Send(sel)
}
// Sets the view to full screen mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/enterFullScreenMode(_:withOptions:)
func (v_ View) EnterFullScreenModeWithOptions(screen unsafe.Pointer, options unsafe.Pointer) bool {
	sel := objc.RegisterName("enterFullScreenMode:withOptions:")
	ret := v_.ID.Send(sel, screen, options)
	return ret != 0
}
// Randomly changes the frame of a view with an ambiguous layout between the different valid values. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/exerciseAmbiguityInLayout()
func (v_ View) ExerciseAmbiguityInLayout() {
	sel := objc.RegisterName("exerciseAmbiguityInLayout")
	v_.ID.Send(sel)
}
// Instructs the view to exit full screen mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/exitFullScreenMode(options:)
func (v_ View) ExitFullScreenModeWithOptions(options unsafe.Pointer) {
	sel := objc.RegisterName("exitFullScreenModeWithOptions:")
	v_.ID.Send(sel, options)
}
// Returns the view’s frame for a given alignment rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/frame(forAlignmentRect:)
func (v_ View) FrameForAlignmentRect(alignmentRect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("frameForAlignmentRect:")
	ret := v_.ID.Send(sel, alignmentRect)
	return unsafe.Pointer(ret)
}
// Returns the identifier for the view’s graphics state object, or 0 if the view doesn’t have a graphics state object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/gState()
func (v_ View) GState() int {
	sel := objc.RegisterName("gState")
	ret := v_.ID.Send(sel)
	return int(ret)
}
// Returns by indirection a list of nonoverlapping rectangles that define the area the view is being asked to draw in  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/getRectsBeingDrawn(_:count:)
func (v_ View) GetRectsBeingDrawnCount(rects unsafe.Pointer, count int) {
	sel := objc.RegisterName("getRectsBeingDrawn:count:")
	v_.ID.Send(sel, rects, count)
}
// Returns a list of rectangles indicating the newly exposed areas of the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/getRectsExposedDuringLiveResize(_:count:)
func (v_ View) GetRectsExposedDuringLiveResizeCount(exposedRects unsafe.Pointer, count int) {
	sel := objc.RegisterName("getRectsExposedDuringLiveResize:count:")
	v_.ID.Send(sel, exposedRects, count)
}
// Returns the farthest descendant of the view in the view hierarchy (including itself) that contains a specified point, or   if that point lies completely outside the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/hitTest(_:)
func (v_ View) HitTest(point unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("hitTest:")
	ret := v_.ID.Send(sel, point)
	return unsafe.Pointer(ret)
}
// Invalidates the view’s intrinsic content size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/invalidateIntrinsicContentSize()
func (v_ View) InvalidateIntrinsicContentSize() {
	sel := objc.RegisterName("invalidateIntrinsicContentSize")
	v_.ID.Send(sel)
}
// Returns a Boolean value that indicates whether the view is a subview of the specified view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isDescendant(of:)
func (v_ View) IsDescendantOf(view unsafe.Pointer) bool {
	sel := objc.RegisterName("isDescendantOf:")
	ret := v_.ID.Send(sel, view)
	return ret != 0
}
// Returns whether a region of the view contains a specified point, accounting for whether the view is flipped or not. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/isMousePoint(_:in:)
func (v_ View) MouseInRect(point unsafe.Pointer, rect unsafe.Pointer) bool {
	sel := objc.RegisterName("mouse:inRect:")
	ret := v_.ID.Send(sel, point, rect)
	return ret != 0
}
// Returns   if the view handles page boundaries,   otherwise. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/knowsPageRange(_:)
func (v_ View) KnowsPageRange(range_ unsafe.Pointer) bool {
	sel := objc.RegisterName("knowsPageRange:")
	ret := v_.ID.Send(sel, range_)
	return ret != 0
}
// Perform layout in concert with the constraint-based layout system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layout()
func (v_ View) Layout() {
	sel := objc.RegisterName("layout")
	v_.ID.Send(sel)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layoutGuideForLayoutRegion:
func (v_ View) LayoutGuideForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("layoutGuideForLayoutRegion:")
	ret := v_.ID.Send(sel, layoutRegion)
	return unsafe.Pointer(ret)
}
// Updates the layout of the receiving view and its subviews based on the current views and constraints. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/layoutSubtreeIfNeeded()
func (v_ View) LayoutSubtreeIfNeeded() {
	sel := objc.RegisterName("layoutSubtreeIfNeeded")
	v_.ID.Send(sel)
}
// Invoked by   to determine the location of the region of the view being printed on the physical page. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/locationOfPrintRect(_:)
func (v_ View) LocationOfPrintRect(rect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("locationOfPrintRect:")
	ret := v_.ID.Send(sel, rect)
	return unsafe.Pointer(ret)
}
// Locks the focus on the view, so subsequent commands take effect in the view’s window and coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/lockFocus()
func (v_ View) LockFocus() {
	sel := objc.RegisterName("lockFocus")
	v_.ID.Send(sel)
}
// Locks the focus to the view atomically if the   method returns   and returns the value of  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/lockFocusIfCanDraw()
func (v_ View) LockFocusIfCanDraw() bool {
	sel := objc.RegisterName("lockFocusIfCanDraw")
	ret := v_.ID.Send(sel)
	return ret != 0
}
// Locks the focus to the view atomically if drawing can occur in the specified graphics context. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/lockFocusIfCanDraw(in:)
func (v_ View) LockFocusIfCanDrawInContext(context unsafe.Pointer) bool {
	sel := objc.RegisterName("lockFocusIfCanDrawInContext:")
	ret := v_.ID.Send(sel, context)
	return ret != 0
}
// Creates the view’s backing layer. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/makeBackingLayer()
func (v_ View) MakeBackingLayer() unsafe.Pointer {
	sel := objc.RegisterName("makeBackingLayer")
	ret := v_.ID.Send(sel)
	return unsafe.Pointer(ret)
}
// Overridden by subclasses to return a context-sensitive pop-up menu for a given mouse-down event. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/menu(for:)
func (v_ View) MenuForEvent(event unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("menuForEvent:")
	ret := v_.ID.Send(sel, event)
	return unsafe.Pointer(ret)
}
// Returns a Boolean value indicating whether the specified rectangle intersects any part of the area that the view is being asked to draw. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/needsToDraw(_:)
func (v_ View) NeedsToDrawRect(rect unsafe.Pointer) bool {
	sel := objc.RegisterName("needsToDrawRect:")
	ret := v_.ID.Send(sel, rect)
	return ret != 0
}
// Invoked to notify the view that the focus ring mask requires updating. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/noteFocusRingMaskChanged()
func (v_ View) NoteFocusRingMaskChanged() {
	sel := objc.RegisterName("noteFocusRingMaskChanged")
	v_.ID.Send(sel)
}
// Implemented by subclasses to respond to key equivalents (also known as keyboard shortcuts). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/performKeyEquivalent(with:)
func (v_ View) PerformKeyEquivalent(event unsafe.Pointer) bool {
	sel := objc.RegisterName("performKeyEquivalent:")
	ret := v_.ID.Send(sel, event)
	return ret != 0
}
// Implemented by subclasses to respond to mnemonics. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/performMnemonic:
func (v_ View) PerformMnemonic(string unsafe.Pointer) bool {
	sel := objc.RegisterName("performMnemonic:")
	ret := v_.ID.Send(sel, string)
	return ret != 0
}
// Prepares the overdraw region for drawing. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/prepareContent(in:)
func (v_ View) PrepareContentInRect(rect unsafe.Pointer) {
	sel := objc.RegisterName("prepareContentInRect:")
	v_.ID.Send(sel, rect)
}
// Restores the view to an initial state so that it can be reused. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/prepareForReuse()
func (v_ View) PrepareForReuse() {
	sel := objc.RegisterName("prepareForReuse")
	v_.ID.Send(sel)
}
// This action method opens the Print panel, and if the user chooses an option other than canceling, prints the view and all its subviews to the device specified in the Print panel. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/printView(_:)
func (v_ View) Print(sender objc.ID) {
	sel := objc.RegisterName("print:")
	v_.ID.Send(sel, sender)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rectForLayoutRegion:
func (v_ View) RectForLayoutRegion(layoutRegion unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("rectForLayoutRegion:")
	ret := v_.ID.Send(sel, layoutRegion)
	return unsafe.Pointer(ret)
}
// Implemented by subclasses to determine the portion of the view to be printed for the specified page number. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rectForPage(_:)
func (v_ View) RectForPage(page int) unsafe.Pointer {
	sel := objc.RegisterName("rectForPage:")
	ret := v_.ID.Send(sel, page)
	return unsafe.Pointer(ret)
}
// Returns the appropriate rectangle to use when magnifying around the specified point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rectForSmartMagnification(at:in:)
func (v_ View) RectForSmartMagnificationAtPointInRect(location unsafe.Pointer, visibleRect unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("rectForSmartMagnificationAtPoint:inRect:")
	ret := v_.ID.Send(sel, location, visibleRect)
	return unsafe.Pointer(ret)
}
// Notifies a clip view’s superview that either the clip view’s bounds rectangle or the document view’s frame rectangle has changed, and that any indicators of the scroll position need to be adjusted. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/reflectScrolledClipView(_:)
func (v_ View) ReflectScrolledClipView(clipView unsafe.Pointer) {
	sel := objc.RegisterName("reflectScrolledClipView:")
	v_.ID.Send(sel, clipView)
}
// Registers the pasteboard types that the view will accept as the destination of an image-dragging session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/registerForDraggedTypes(_:)
func (v_ View) RegisterForDraggedTypes(newTypes unsafe.Pointer) {
	sel := objc.RegisterName("registerForDraggedTypes:")
	v_.ID.Send(sel, newTypes)
}
// Frees the view’s graphics state object, if it has one. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/releaseGState()
func (v_ View) ReleaseGState() {
	sel := objc.RegisterName("releaseGState")
	v_.ID.Send(sel)
}
// Removes all tooltips assigned to the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeAllToolTips()
func (v_ View) RemoveAllToolTips() {
	sel := objc.RegisterName("removeAllToolTips")
	v_.ID.Send(sel)
}
// Removes the specified constraint from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeConstraint(_:)
func (v_ View) RemoveConstraint(constraint unsafe.Pointer) {
	sel := objc.RegisterName("removeConstraint:")
	v_.ID.Send(sel, constraint)
}
// Removes the specified constraints from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeConstraints(_:)
func (v_ View) RemoveConstraints(constraints unsafe.Pointer) {
	sel := objc.RegisterName("removeConstraints:")
	v_.ID.Send(sel, constraints)
}
// Completely removes a cursor rectangle from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeCursorRect(_:cursor:)
func (v_ View) RemoveCursorRectCursor(rect unsafe.Pointer, object unsafe.Pointer) {
	sel := objc.RegisterName("removeCursorRect:cursor:")
	v_.ID.Send(sel, rect, object)
}
// Unlinks the view from its superview and its window, removes it from the responder chain, and invalidates its cursor rectangles. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeFromSuperview()
func (v_ View) RemoveFromSuperview() {
	sel := objc.RegisterName("removeFromSuperview")
	v_.ID.Send(sel)
}
// Unlinks the view from its superview and its window and removes it from the responder chain, but does not invalidate its cursor rectangles to cause redrawing. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeFromSuperviewWithoutNeedingDisplay()
func (v_ View) RemoveFromSuperviewWithoutNeedingDisplay() {
	sel := objc.RegisterName("removeFromSuperviewWithoutNeedingDisplay")
	v_.ID.Send(sel)
}
// Detaches a gesture recognizer from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeGestureRecognizer(_:)
func (v_ View) RemoveGestureRecognizer(gestureRecognizer unsafe.Pointer) {
	sel := objc.RegisterName("removeGestureRecognizer:")
	v_.ID.Send(sel, gestureRecognizer)
}
// Removes the provided layout guide from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeLayoutGuide(_:)
func (v_ View) RemoveLayoutGuide(guide unsafe.Pointer) {
	sel := objc.RegisterName("removeLayoutGuide:")
	v_.ID.Send(sel, guide)
}
// Removes the tooltip identified by specified tag. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeToolTip(_:)
func (v_ View) RemoveToolTip(tag unsafe.Pointer) {
	sel := objc.RegisterName("removeToolTip:")
	v_.ID.Send(sel, tag)
}
// Removes a given tracking area from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeTrackingArea(_:)
func (v_ View) RemoveTrackingArea(trackingArea unsafe.Pointer) {
	sel := objc.RegisterName("removeTrackingArea:")
	v_.ID.Send(sel, trackingArea)
}
// Removes the tracking rectangle identified by a tag. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/removeTrackingRect(_:)
func (v_ View) RemoveTrackingRect(tag unsafe.Pointer) {
	sel := objc.RegisterName("removeTrackingRect:")
	v_.ID.Send(sel, tag)
}
// Invalidates the view’s graphics state object, if it has one. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/renewGState()
func (v_ View) RenewGState() {
	sel := objc.RegisterName("renewGState")
	v_.ID.Send(sel)
}
// Replaces one of the view’s subviews with another view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/replaceSubview(_:with:)
func (v_ View) ReplaceSubviewWith(oldView unsafe.Pointer, newView unsafe.Pointer) {
	sel := objc.RegisterName("replaceSubview:with:")
	v_.ID.Send(sel, oldView, newView)
}
// Overridden by subclasses to define their default cursor rectangles. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/resetCursorRects()
func (v_ View) ResetCursorRects() {
	sel := objc.RegisterName("resetCursorRects")
	v_.ID.Send(sel)
}
// Informs the view that the bounds size of its superview has changed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/resize(withOldSuperviewSize:)
func (v_ View) ResizeWithOldSuperviewSize(oldSize unsafe.Pointer) {
	sel := objc.RegisterName("resizeWithOldSuperviewSize:")
	v_.ID.Send(sel, oldSize)
}
// Informs the view’s subviews that the view’s bounds rectangle size has changed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/resizeSubviews(withOldSize:)
func (v_ View) ResizeSubviewsWithOldSize(oldSize unsafe.Pointer) {
	sel := objc.RegisterName("resizeSubviewsWithOldSize:")
	v_.ID.Send(sel, oldSize)
}
// Rotates the view’s bounds rectangle by a specified degree value around the origin of the coordinate system, (0.0, 0.0). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rotate(byDegrees:)
func (v_ View) RotateByAngle(angle float64) {
	sel := objc.RegisterName("rotateByAngle:")
	v_.ID.Send(sel, angle)
}
// Informs the client that   allowed the user to add  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:didAdd:)
func (v_ View) RulerViewDidAddMarker(ruler unsafe.Pointer, marker unsafe.Pointer) {
	sel := objc.RegisterName("rulerView:didAddMarker:")
	v_.ID.Send(sel, ruler, marker)
}
// Informs the client that   allowed the user to move  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:didMove:)
func (v_ View) RulerViewDidMoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) {
	sel := objc.RegisterName("rulerView:didMoveMarker:")
	v_.ID.Send(sel, ruler, marker)
}
// Informs the client that   allowed the user to remove  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:didRemove:)
func (v_ View) RulerViewDidRemoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) {
	sel := objc.RegisterName("rulerView:didRemoveMarker:")
	v_.ID.Send(sel, ruler, marker)
}
// Informs the client that the user has pressed the mouse button while the cursor is in the ruler area of  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:handleMouseDownWith:)
func (v_ View) RulerViewHandleMouseDown(ruler unsafe.Pointer, event unsafe.Pointer) {
	sel := objc.RegisterName("rulerView:handleMouseDown:")
	v_.ID.Send(sel, ruler, event)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:locationFor:)
func (v_ View) RulerViewLocationForPoint(ruler unsafe.Pointer, point unsafe.Pointer) float64 {
	sel := objc.RegisterName("rulerView:locationForPoint:")
	ret := v_.ID.Send(sel, ruler, point)
	return float64(ret)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:pointForLocation:)
func (v_ View) RulerViewPointForLocation(ruler unsafe.Pointer, point float64) unsafe.Pointer {
	sel := objc.RegisterName("rulerView:pointForLocation:")
	ret := v_.ID.Send(sel, ruler, point)
	return unsafe.Pointer(ret)
}
// Requests permission for   to add  , an NSRulerMarker being dragged onto the ruler by the user. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:shouldAdd:)
func (v_ View) RulerViewShouldAddMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool {
	sel := objc.RegisterName("rulerView:shouldAddMarker:")
	ret := v_.ID.Send(sel, ruler, marker)
	return ret != 0
}
// Requests permission for   to move  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:shouldMove:)
func (v_ View) RulerViewShouldMoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool {
	sel := objc.RegisterName("rulerView:shouldMoveMarker:")
	ret := v_.ID.Send(sel, ruler, marker)
	return ret != 0
}
// Requests permission for   to remove  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:shouldRemove:)
func (v_ View) RulerViewShouldRemoveMarker(ruler unsafe.Pointer, marker unsafe.Pointer) bool {
	sel := objc.RegisterName("rulerView:shouldRemoveMarker:")
	ret := v_.ID.Send(sel, ruler, marker)
	return ret != 0
}
// Informs the client that   will add the new NSRulerMarker,  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:willAdd:atLocation:)
func (v_ View) RulerViewWillAddMarkerAtLocation(ruler unsafe.Pointer, marker unsafe.Pointer, location float64) float64 {
	sel := objc.RegisterName("rulerView:willAddMarker:atLocation:")
	ret := v_.ID.Send(sel, ruler, marker, location)
	return float64(ret)
}
// Informs the client that   will move  , an NSRulerMarker already on the ruler view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:willMove:toLocation:)
func (v_ View) RulerViewWillMoveMarkerToLocation(ruler unsafe.Pointer, marker unsafe.Pointer, location float64) float64 {
	sel := objc.RegisterName("rulerView:willMoveMarker:toLocation:")
	ret := v_.ID.Send(sel, ruler, marker, location)
	return float64(ret)
}
// Informs the client view that   is about to be appropriated by  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/rulerView(_:willSetClientView:)
func (v_ View) RulerViewWillSetClientView(ruler unsafe.Pointer, newClient unsafe.Pointer) {
	sel := objc.RegisterName("rulerView:willSetClientView:")
	v_.ID.Send(sel, ruler, newClient)
}
// Scales the view’s coordinate system so that the unit square scales to the specified dimensions. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/scaleUnitSquare(to:)
func (v_ View) ScaleUnitSquareToSize(newUnitSize unsafe.Pointer) {
	sel := objc.RegisterName("scaleUnitSquareToSize:")
	v_.ID.Send(sel, newUnitSize)
}
// Scrolls the view’s closest ancestor   object so a point in the view lies at the origin of the clip view’s bounds rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/scroll(_:)
func (v_ View) ScrollPoint(point unsafe.Pointer) {
	sel := objc.RegisterName("scrollPoint:")
	v_.ID.Send(sel, point)
}
// Copies the visible portion of the view’s rendered image within a region and lays that portion down again at a specified offset . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/scroll(_:by:)
func (v_ View) ScrollRectBy(rect unsafe.Pointer, delta unsafe.Pointer) {
	sel := objc.RegisterName("scrollRect:by:")
	v_.ID.Send(sel, rect, delta)
}
// Notifies the superview of a clip view that the clip view needs to reset the origin of its bounds rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/scroll(_:to:)
func (v_ View) ScrollClipViewToPoint(clipView unsafe.Pointer, point unsafe.Pointer) {
	sel := objc.RegisterName("scrollClipView:toPoint:")
	v_.ID.Send(sel, clipView, point)
}
// Scrolls the view’s closest ancestor   object the minimum distance needed so a specified region of the view becomes visible in the clip view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/scrollToVisible(_:)
func (v_ View) ScrollRectToVisible(rect unsafe.Pointer) bool {
	sel := objc.RegisterName("scrollRectToVisible:")
	ret := v_.ID.Send(sel, rect)
	return ret != 0
}
// Sets the origin of the view’s bounds rectangle to a specified point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setBoundsOrigin(_:)
func (v_ View) SetBoundsOrigin(newOrigin unsafe.Pointer) {
	sel := objc.RegisterName("setBoundsOrigin:")
	v_.ID.Send(sel, newOrigin)
}
// Sets the size of the view’s bounds rectangle to specified dimensions, inversely scaling its coordinate system relative to its frame rectangle. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setBoundsSize(_:)
func (v_ View) SetBoundsSize(newSize unsafe.Pointer) {
	sel := objc.RegisterName("setBoundsSize:")
	v_.ID.Send(sel, newSize)
}
// Sets the priority with which a view resists being made smaller than its intrinsic size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setContentCompressionResistancePriority(_:for:)
func (v_ View) SetContentCompressionResistancePriorityForOrientation(priority unsafe.Pointer, orientation unsafe.Pointer) {
	sel := objc.RegisterName("setContentCompressionResistancePriority:forOrientation:")
	v_.ID.Send(sel, priority, orientation)
}
// Sets the priority with which a view resists being made larger than its intrinsic size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setContentHuggingPriority(_:for:)
func (v_ View) SetContentHuggingPriorityForOrientation(priority unsafe.Pointer, orientation unsafe.Pointer) {
	sel := objc.RegisterName("setContentHuggingPriority:forOrientation:")
	v_.ID.Send(sel, priority, orientation)
}
// Sets the origin of the view’s frame rectangle to the specified point, effectively repositioning it within its superview. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setFrameOrigin(_:)
func (v_ View) SetFrameOrigin(newOrigin unsafe.Pointer) {
	sel := objc.RegisterName("setFrameOrigin:")
	v_.ID.Send(sel, newOrigin)
}
// Sets the size of the view’s frame rectangle to the specified dimensions, resizing it within its superview without affecting its coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setFrameSize(_:)
func (v_ View) SetFrameSize(newSize unsafe.Pointer) {
	sel := objc.RegisterName("setFrameSize:")
	v_.ID.Send(sel, newSize)
}
// Invalidates the area around the focus ring. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setKeyboardFocusRingNeedsDisplay(_:)
func (v_ View) SetKeyboardFocusRingNeedsDisplayInRect(rect unsafe.Pointer) {
	sel := objc.RegisterName("setKeyboardFocusRingNeedsDisplayInRect:")
	v_.ID.Send(sel, rect)
}
// Marks the region of the view within the specified rectangle as needing display, increasing the view’s existing invalid region to include it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setNeedsDisplay(_:)
func (v_ View) SetNeedsDisplayInRect(invalidRect unsafe.Pointer) {
	sel := objc.RegisterName("setNeedsDisplayInRect:")
	v_.ID.Send(sel, invalidRect)
}
// Overridden by subclasses to (re)initialize the view’s graphics state object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/setUpGState()
func (v_ View) SetUpGState() {
	sel := objc.RegisterName("setUpGState")
	v_.ID.Send(sel)
}
// Allows the user to drag objects from the view without activating the app or moving the window of the view forward, possibly obscuring the destination. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/shouldDelayWindowOrdering(for:)
func (v_ View) ShouldDelayWindowOrderingForEvent(event unsafe.Pointer) bool {
	sel := objc.RegisterName("shouldDelayWindowOrderingForEvent:")
	ret := v_.ID.Send(sel, event)
	return ret != 0
}
// Returns a Boolean value indicating whether the view is being drawn to an environment that supports color. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/shouldDrawColor()
func (v_ View) ShouldDrawColor() bool {
	sel := objc.RegisterName("shouldDrawColor")
	ret := v_.ID.Send(sel)
	return ret != 0
}
// Shows a window displaying the definition of the attributed string at the specified point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/showDefinition(for:at:)
func (v_ View) ShowDefinitionForAttributedStringAtPoint(attrString unsafe.Pointer, textBaselineOrigin unsafe.Pointer) {
	sel := objc.RegisterName("showDefinitionForAttributedString:atPoint:")
	v_.ID.Send(sel, attrString, textBaselineOrigin)
}
// Shows a window displaying the definition of the specified range of the attributed string. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/showDefinition(for:range:options:baselineOriginProvider:)
func (v_ View) ShowDefinitionForAttributedStringRangeOptionsBaselineOriginProvider(attrString unsafe.Pointer, targetRange unsafe.Pointer, options unsafe.Pointer, originProvider unsafe.Pointer) {
	sel := objc.RegisterName("showDefinitionForAttributedString:range:options:baselineOriginProvider:")
	v_.ID.Send(sel, attrString, targetRange, options, originProvider)
}
// Orders the view’s immediate subviews using the specified comparator function. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/sortSubviews(_:context:)
func (v_ View) SortSubviewsUsingFunctionContext(compare unsafe.Pointer, context unsafe.Pointer) {
	sel := objc.RegisterName("sortSubviewsUsingFunction:context:")
	v_.ID.Send(sel, compare, context)
}
// Translates the view’s coordinate system so that its origin moves to a new location. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/translateOrigin(to:)
func (v_ View) TranslateOriginToPoint(translation unsafe.Pointer) {
	sel := objc.RegisterName("translateOriginToPoint:")
	v_.ID.Send(sel, translation)
}
// Translates the display rectangles by the specified delta. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/translateRectsNeedingDisplay(in:by:)
func (v_ View) TranslateRectsNeedingDisplayInRectBy(clipRect unsafe.Pointer, delta unsafe.Pointer) {
	sel := objc.RegisterName("translateRectsNeedingDisplayInRect:by:")
	v_.ID.Send(sel, clipRect, delta)
}
// Unlocks focus from the current view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/unlockFocus()
func (v_ View) UnlockFocus() {
	sel := objc.RegisterName("unlockFocus")
	v_.ID.Send(sel)
}
// Unregisters the view as a possible destination in a dragging session. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/unregisterDraggedTypes()
func (v_ View) UnregisterDraggedTypes() {
	sel := objc.RegisterName("unregisterDraggedTypes")
	v_.ID.Send(sel)
}
// Update constraints for the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/updateConstraints()
func (v_ View) UpdateConstraints() {
	sel := objc.RegisterName("updateConstraints")
	v_.ID.Send(sel)
}
// Updates the constraints for the receiving view and its subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/updateConstraintsForSubtreeIfNeeded()
func (v_ View) UpdateConstraintsForSubtreeIfNeeded() {
	sel := objc.RegisterName("updateConstraintsForSubtreeIfNeeded")
	v_.ID.Send(sel)
}
// Updates the view’s content by modifying its underlying layer. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/updateLayer()
func (v_ View) UpdateLayer() {
	sel := objc.RegisterName("updateLayer")
	v_.ID.Send(sel)
}
// Invoked automatically when the view’s geometry changes such that its tracking areas need to be recalculated. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/updateTrackingAreas()
func (v_ View) UpdateTrackingAreas() {
	sel := objc.RegisterName("updateTrackingAreas")
	v_.ID.Send(sel)
}
// Responds when the view’s backing store properties change. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidChangeBackingProperties()
func (v_ View) ViewDidChangeBackingProperties() {
	sel := objc.RegisterName("viewDidChangeBackingProperties")
	v_.ID.Send(sel)
}
// Informs the view that its effective appearance changed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidChangeEffectiveAppearance()
func (v_ View) ViewDidChangeEffectiveAppearance() {
	sel := objc.RegisterName("viewDidChangeEffectiveAppearance")
	v_.ID.Send(sel)
}
// Informs the view of the end of a live resize—the user has finished resizing the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidEndLiveResize()
func (v_ View) ViewDidEndLiveResize() {
	sel := objc.RegisterName("viewDidEndLiveResize")
	v_.ID.Send(sel)
}
// Invoked when the view is hidden, either directly, or in response to an ancestor being hidden. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidHide()
func (v_ View) ViewDidHide() {
	sel := objc.RegisterName("viewDidHide")
	v_.ID.Send(sel)
}
// Informs the view that its superview has changed (possibly to  ). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidMoveToSuperview()
func (v_ View) ViewDidMoveToSuperview() {
	sel := objc.RegisterName("viewDidMoveToSuperview")
	v_.ID.Send(sel)
}
// Informs the view that it has been added to a new view hierarchy. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidMoveToWindow()
func (v_ View) ViewDidMoveToWindow() {
	sel := objc.RegisterName("viewDidMoveToWindow")
	v_.ID.Send(sel)
}
// Invoked when the view is unhidden, either directly, or in response to an ancestor being unhidden [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewDidUnhide()
func (v_ View) ViewDidUnhide() {
	sel := objc.RegisterName("viewDidUnhide")
	v_.ID.Send(sel)
}
// Informs the view that it’s required to draw content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewWillDraw()
func (v_ View) ViewWillDraw() {
	sel := objc.RegisterName("viewWillDraw")
	v_.ID.Send(sel)
}
// Informs the view that its superview is about to change to the specified superview (which may be  ). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewWillMove(toSuperview:)
func (v_ View) ViewWillMoveToSuperview(newSuperview unsafe.Pointer) {
	sel := objc.RegisterName("viewWillMoveToSuperview:")
	v_.ID.Send(sel, newSuperview)
}
// Informs the view that it’s being added to the view hierarchy of the specified window object (which may be  ). [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewWillMove(toWindow:)
func (v_ View) ViewWillMoveToWindow(newWindow unsafe.Pointer) {
	sel := objc.RegisterName("viewWillMoveToWindow:")
	v_.ID.Send(sel, newWindow)
}
// Informs the view of the start of a live resize—the user has started resizing the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewWillStartLiveResize()
func (v_ View) ViewWillStartLiveResize() {
	sel := objc.RegisterName("viewWillStartLiveResize")
	v_.ID.Send(sel)
}
// Returns the view’s nearest descendant (including itself) with a specific tag, or   if no subview has that tag. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/viewWithTag(_:)
func (v_ View) ViewWithTag(tag int) unsafe.Pointer {
	sel := objc.RegisterName("viewWithTag:")
	ret := v_.ID.Send(sel, tag)
	return unsafe.Pointer(ret)
}
// Called just before a contextual menu for a view is opened on screen. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/willOpenMenu(_:with:)
func (v_ View) WillOpenMenuWithEvent(menu unsafe.Pointer, event unsafe.Pointer) {
	sel := objc.RegisterName("willOpenMenu:withEvent:")
	v_.ID.Send(sel, menu, event)
}
// Overridden by subclasses to perform additional actions before subviews are removed from the view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/willRemoveSubview(_:)
func (v_ View) WillRemoveSubview(subview unsafe.Pointer) {
	sel := objc.RegisterName("willRemoveSubview:")
	v_.ID.Send(sel, subview)
}
// Writes EPS data that draws the region of the view within a specified rectangle onto a pasteboard. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/writeEPS(inside:to:)
func (v_ View) WriteEPSInsideRectToPasteboard(rect unsafe.Pointer, pasteboard unsafe.Pointer) {
	sel := objc.RegisterName("writeEPSInsideRect:toPasteboard:")
	v_.ID.Send(sel, rect, pasteboard)
}
// Writes PDF data that draws the region of the view within a specified rectangle onto a pasteboard. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSView/writePDF(inside:to:)
func (v_ View) WritePDFInsideRectToPasteboard(rect unsafe.Pointer, pasteboard unsafe.Pointer) {
	sel := objc.RegisterName("writePDFInsideRect:toPasteboard:")
	v_.ID.Send(sel, rect, pasteboard)
}

