// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Cursor] class.
var (
	CursorClass     _CursorClass
	CursorClassOnce sync.Once
)

func getCursorClass() _CursorClass {
	CursorClassOnce.Do(func() {
		CursorClass = _CursorClass{objc.GetClass("NSCursor")}
	})
	return CursorClass
}

type _CursorClass struct {
	class objc.Class
}

// An interface definition for the [Cursor] class.
type ICursor interface {
	objectivec.IObject
	MouseEntered(event IEvent)
	MouseExited(event IEvent)
	Pop()
	Push()
	Set()
	SetOnMouseEntered(flag bool)
	SetOnMouseExited(flag bool)
	HotSpot() coregraphics.CGPoint
	Image() Image
	SetOnMouseEntered() bool
	SetOnMouseExited() bool
	IsSetOnMouseEntered() bool
	SetIsSetOnMouseEntered(value bool)
	IsSetOnMouseExited() bool
	SetIsSetOnMouseExited(value bool)
}

// A pointer (also called a cursor).
//
// The following table shows and describes the system cursors, and indicates the class method for obtaining them: In macOS 10.3 and later, cursor size is no longer limited to 16 by 16 pixels.


// A pointer (also called a cursor).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor
type Cursor struct {
	objectivec.Object
}

// CursorFrom constructs a [Cursor] from an unsafe.Pointer.
//
// A pointer (also called a cursor).
func CursorFrom(ptr unsafe.Pointer) Cursor {
	return Cursor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CursorClass) Alloc() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CursorClass) New() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Cursor) Init() Cursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Cursor) Autorelease() Cursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCursor creates a new Cursor instance.
func NewCursor() Cursor {
	return getCursorClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/init(coder:)
func NewCursorWithCoder(coder foundation.ICoder) Cursor {
	instance := getCursorClass().Alloc()
	rv := objc.Send[Cursor](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes the cursor with the specified image and hot spot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/init(image:foregroundColorHint:backgroundColorHint:hotSpot:)
func NewCursorWithImageForegroundColorHintBackgroundColorHintHotSpot(newImage IImage, fg IColor, bg IColor, hotSpot coregraphics.CGPoint) Cursor {
	instance := getCursorClass().Alloc()
	rv := objc.Send[Cursor](instance.ID, objc.Sel("initWithImage:foregroundColorHint:backgroundColorHint:hotSpot:"), newImage, fg, bg, hotSpot)
	rv.Autorelease()
	return rv
}


// Initializes a cursor with the given image and hot spot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/init(image:hotSpot:)
func NewCursorWithImageHotSpot(newImage IImage, point coregraphics.CGPoint) Cursor {
	instance := getCursorClass().Alloc()
	rv := objc.Send[Cursor](instance.ID, objc.Sel("initWithImage:hotSpot:"), newImage, point)
	rv.Autorelease()
	return rv
}



// Returns the cursor for resizing a column (vertical divider) in the specified directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/columnResizeCursorInDirections:
func (cc _CursorClass) ColumnResizeCursorInDirections(directions IHorizontalDirections) Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("columnResizeCursorInDirections:"), directions)
	return rv
}


// Returns the cursor for resizing a rectangular frame from the specified edge or corner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/frameResizeCursorFromPosition:inDirections:
func (cc _CursorClass) FrameResizeCursorFromPositionInDirections(position ICursorFrameResizePosition, directions ICursorFrameResizeDirections) Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("frameResizeCursorFromPosition:inDirections:"), position, directions)
	return rv
}


// Makes the current cursor invisible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/hide()
func (cc _CursorClass) Hide() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("hide"))
}


// Pops the current cursor off the top of the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/pop()-swift.type.method
func (cc _CursorClass) Pop() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("pop"))
}


// Returns the cursor for resizing a row (horizontal divider) in the specified directions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/rowResizeCursorInDirections:
func (cc _CursorClass) RowResizeCursorInDirections(directions IVerticalDirections) Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("rowResizeCursorInDirections:"), directions)
	return rv
}


// Sets whether the cursor is hidden until the mouse moves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/setHiddenUntilMouseMoves(_:)
func (cc _CursorClass) SetHiddenUntilMouseMoves(flag bool) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("setHiddenUntilMouseMoves:"), flag)
}


// Negates an earlier call to by showing the current cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/unhide()
func (cc _CursorClass) Unhide() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("unhide"))
}


// Returns the default cursor, the arrow cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/arrow
func (cc _CursorClass) ArrowCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("arrowCursor"))
	return rv
}

// Returns the closed-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/closedHand
func (cc _CursorClass) ClosedHandCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("closedHandCursor"))
	return rv
}

// Returns the cursor for resizing a column (vertical divider) in either direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/columnResize
func (cc _CursorClass) ColumnResizeCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("columnResizeCursor"))
	return rv
}

// Returns the contextual menu system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/contextualMenu
func (cc _CursorClass) ContextualMenuCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("contextualMenuCursor"))
	return rv
}

// Returns the cross-hair system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/crosshair
func (cc _CursorClass) CrosshairCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("crosshairCursor"))
	return rv
}

// Returns the application’s current cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/current
func (cc _CursorClass) CurrentCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("currentCursor"))
	return rv
}

// Returns the current system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/currentSystem
func (cc _CursorClass) CurrentSystemCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("currentSystemCursor"))
	return rv
}

// Returns a cursor indicating that the current operation will result in a disappearing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/disappearingItem
func (cc _CursorClass) DisappearingItemCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("disappearingItemCursor"))
	return rv
}

// Returns a cursor indicating that the current operation will result in a copy action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragCopy
func (cc _CursorClass) DragCopyCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("dragCopyCursor"))
	return rv
}

// Returns a cursor indicating that the current operation will result in a link action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragLink
func (cc _CursorClass) DragLinkCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("dragLinkCursor"))
	return rv
}

// Returns a cursor that looks like a capital I with a tiny crossbeam at its middle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeam
func (cc _CursorClass) IBeamCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("IBeamCursor"))
	return rv
}

// Returns the cursor for editing vertical layout text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeamCursorForVerticalLayout
func (cc _CursorClass) IBeamCursorForVerticalLayout() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("IBeamCursorForVerticalLayout"))
	return rv
}

// Returns the open-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/openHand
func (cc _CursorClass) OpenHandCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("openHandCursor"))
	return rv
}

// Returns the operation not allowed cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/operationNotAllowed
func (cc _CursorClass) OperationNotAllowedCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("operationNotAllowedCursor"))
	return rv
}

// Returns the pointing-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/pointingHand
func (cc _CursorClass) PointingHandCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("pointingHandCursor"))
	return rv
}

// Returns the resize-down system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeDown
func (cc _CursorClass) ResizeDownCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("resizeDownCursor"))
	return rv
}

// Returns the resize-left system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeLeft
func (cc _CursorClass) ResizeLeftCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("resizeLeftCursor"))
	return rv
}

// Returns the resize-left-and-right system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeLeftRight
func (cc _CursorClass) ResizeLeftRightCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("resizeLeftRightCursor"))
	return rv
}

// Returns the resize-right system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeRight
func (cc _CursorClass) ResizeRightCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("resizeRightCursor"))
	return rv
}

// Returns the resize-up system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeUp
func (cc _CursorClass) ResizeUpCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("resizeUpCursor"))
	return rv
}

// Returns the resize-up-and-down system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeUpDown
func (cc _CursorClass) ResizeUpDownCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("resizeUpDownCursor"))
	return rv
}

// Returns the cursor for resizing a row (horizontal divider) in either direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/rowResize
func (cc _CursorClass) RowResizeCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("rowResizeCursor"))
	return rv
}

// Returns the zoom-in cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/zoomIn
func (cc _CursorClass) ZoomInCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("zoomInCursor"))
	return rv
}

// Returns the zoom-out cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/zoomOut
func (cc _CursorClass) ZoomOutCursor() Cursor {
	rv := objc.Send[NSCursor](objc.ID(cc.class), objc.Sel("zoomOutCursor"))
	return rv
}

// Automatically sent to the receiver when the cursor enters a cursor rectangle owned by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/mouseEntered(with:)
func (c_ Cursor) MouseEntered(event IEvent) {
	objc.Send[objc.ID](c_.ID, objc.Sel("mouseEntered:"), event)
}


// Automatically sent to the receiver when the cursor exits a cursor rectangle owned by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/mouseExited(with:)
func (c_ Cursor) MouseExited(event IEvent) {
	objc.Send[objc.ID](c_.ID, objc.Sel("mouseExited:"), event)
}


// Sends a message to the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/pop()-swift.method
func (c_ Cursor) Pop() {
	objc.Send[objc.ID](c_.ID, objc.Sel("pop"))
}


// Puts the receiver on top of the cursor stack and makes it the current cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/push()
func (c_ Cursor) Push() {
	objc.Send[objc.ID](c_.ID, objc.Sel("push"))
}


// Makes the receiver the current cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/set()
func (c_ Cursor) Set() {
	objc.Send[objc.ID](c_.ID, objc.Sel("set"))
}


// Specifies whether the receiver accepts events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/setOnMouseEntered(_:)
func (c_ Cursor) SetOnMouseEntered(flag bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOnMouseEntered:"), flag)
}


// Sets whether the receiver accepts events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/setOnMouseExited(_:)
func (c_ Cursor) SetOnMouseExited(flag bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOnMouseExited:"), flag)
}


// Returns the default cursor, the arrow cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/arrow
func (c_ Cursor) ArrowCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("arrowCursor"))
	return rv
}


// Returns the closed-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/closedHand
func (c_ Cursor) ClosedHandCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("closedHandCursor"))
	return rv
}


// Returns the cursor for resizing a column (vertical divider) in either direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/columnResize
func (c_ Cursor) ColumnResizeCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("columnResizeCursor"))
	return rv
}


// Returns the contextual menu system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/contextualMenu
func (c_ Cursor) ContextualMenuCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("contextualMenuCursor"))
	return rv
}


// Returns the cross-hair system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/crosshair
func (c_ Cursor) CrosshairCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("crosshairCursor"))
	return rv
}


// Returns the application’s current cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/current
func (c_ Cursor) CurrentCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("currentCursor"))
	return rv
}


// Returns the current system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/currentSystem
func (c_ Cursor) CurrentSystemCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("currentSystemCursor"))
	return rv
}


// Returns a cursor indicating that the current operation will result in a disappearing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/disappearingItem
func (c_ Cursor) DisappearingItemCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("disappearingItemCursor"))
	return rv
}


// Returns a cursor indicating that the current operation will result in a copy action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragCopy
func (c_ Cursor) DragCopyCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("dragCopyCursor"))
	return rv
}


// Returns a cursor indicating that the current operation will result in a link action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragLink
func (c_ Cursor) DragLinkCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("dragLinkCursor"))
	return rv
}


// The position of the click location within the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/hotSpot
func (c_ Cursor) HotSpot() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](c_.ID, objc.Sel("hotSpot"))
	return rv
}


// Returns a cursor that looks like a capital I with a tiny crossbeam at its middle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeam
func (c_ Cursor) IBeamCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("IBeamCursor"))
	return rv
}


// Returns the cursor for editing vertical layout text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeamCursorForVerticalLayout
func (c_ Cursor) IBeamCursorForVerticalLayout() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("IBeamCursorForVerticalLayout"))
	return rv
}


// The cursor’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/image
func (c_ Cursor) Image() Image {
	rv := objc.Send[Image](c_.ID, objc.Sel("image"))
	return rv
}


// A Boolean value indicating whether the receiver becomes current on receiving a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/isSetOnMouseEntered
func (c_ Cursor) SetOnMouseEntered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setOnMouseEntered"))
	return rv
}


// A Boolean value indicating whether the receiver becomes current when it receives a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/isSetOnMouseExited
func (c_ Cursor) SetOnMouseExited() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setOnMouseExited"))
	return rv
}


// Returns the open-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/openHand
func (c_ Cursor) OpenHandCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("openHandCursor"))
	return rv
}


// Returns the operation not allowed cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/operationNotAllowed
func (c_ Cursor) OperationNotAllowedCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("operationNotAllowedCursor"))
	return rv
}


// Returns the pointing-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/pointingHand
func (c_ Cursor) PointingHandCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("pointingHandCursor"))
	return rv
}


// Returns the resize-down system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeDown
func (c_ Cursor) ResizeDownCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("resizeDownCursor"))
	return rv
}


// Returns the resize-left system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeLeft
func (c_ Cursor) ResizeLeftCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("resizeLeftCursor"))
	return rv
}


// Returns the resize-left-and-right system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeLeftRight
func (c_ Cursor) ResizeLeftRightCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("resizeLeftRightCursor"))
	return rv
}


// Returns the resize-right system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeRight
func (c_ Cursor) ResizeRightCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("resizeRightCursor"))
	return rv
}


// Returns the resize-up system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeUp
func (c_ Cursor) ResizeUpCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("resizeUpCursor"))
	return rv
}


// Returns the resize-up-and-down system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeUpDown
func (c_ Cursor) ResizeUpDownCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("resizeUpDownCursor"))
	return rv
}


// Returns the cursor for resizing a row (horizontal divider) in either direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/rowResize
func (c_ Cursor) RowResizeCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("rowResizeCursor"))
	return rv
}


// Returns the zoom-in cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/zoomIn
func (c_ Cursor) ZoomInCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("zoomInCursor"))
	return rv
}


// Returns the zoom-out cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/zoomOut
func (c_ Cursor) ZoomOutCursor() NSCursor {
	rv := objc.Send[NSCursor](c_.ID, objc.Sel("zoomOutCursor"))
	return rv
}


// A Boolean value indicating whether the receiver becomes current on receiving a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/issetonmouseentered
func (c_ Cursor) IsSetOnMouseEntered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSetOnMouseEntered"))
	return rv
}


// A Boolean value indicating whether the receiver becomes current on receiving a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/issetonmouseentered
func (c_ Cursor) SetIsSetOnMouseEntered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSetOnMouseEntered:"), value)
}


// A Boolean value indicating whether the receiver becomes current when it receives a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/issetonmouseexited
func (c_ Cursor) IsSetOnMouseExited() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSetOnMouseExited"))
	return rv
}


// A Boolean value indicating whether the receiver becomes current when it receives a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/issetonmouseexited
func (c_ Cursor) SetIsSetOnMouseExited(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSetOnMouseExited:"), value)
}


