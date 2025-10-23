// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	HotSpot() coregraphics.CGPoint
	SetHotSpot(value coregraphics.CGPoint)
	Image() IImage
	SetImage(value IImage)
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



// Makes the current cursor invisible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/hide()
func (cc _CursorClass) Hide() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("hide"))
}


// Returns the default cursor, the arrow cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/arrow
func (cc _CursorClass) ArrowCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("arrowCursor"))
	return rv
}

// Returns the closed-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/closedHand
func (cc _CursorClass) ClosedHandCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("closedHandCursor"))
	return rv
}

// Returns the contextual menu system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/contextualMenu
func (cc _CursorClass) ContextualMenuCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("contextualMenuCursor"))
	return rv
}

// Returns the cross-hair system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/crosshair
func (cc _CursorClass) CrosshairCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("crosshairCursor"))
	return rv
}

// Returns a cursor indicating that the current operation will result in a disappearing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/disappearingItem
func (cc _CursorClass) DisappearingItemCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("disappearingItemCursor"))
	return rv
}

// Returns a cursor indicating that the current operation will result in a copy action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragCopy
func (cc _CursorClass) DragCopyCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("dragCopyCursor"))
	return rv
}

// Returns a cursor indicating that the current operation will result in a link action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragLink
func (cc _CursorClass) DragLinkCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("dragLinkCursor"))
	return rv
}

// Returns a cursor that looks like a capital I with a tiny crossbeam at its middle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeam
func (cc _CursorClass) IBeamCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("IBeamCursor"))
	return rv
}

// Returns the cursor for editing vertical layout text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeamCursorForVerticalLayout
func (cc _CursorClass) IBeamCursorForVerticalLayout() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("IBeamCursorForVerticalLayout"))
	return rv
}

// Returns the open-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/openHand
func (cc _CursorClass) OpenHandCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("openHandCursor"))
	return rv
}

// Returns the operation not allowed cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/operationNotAllowed
func (cc _CursorClass) OperationNotAllowedCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("operationNotAllowedCursor"))
	return rv
}

// Returns the pointing-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/pointingHand
func (cc _CursorClass) PointingHandCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("pointingHandCursor"))
	return rv
}

// Returns the resize-down system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeDown
func (cc _CursorClass) ResizeDownCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("resizeDownCursor"))
	return rv
}

// Returns the resize-left system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeLeft
func (cc _CursorClass) ResizeLeftCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("resizeLeftCursor"))
	return rv
}

// Returns the resize-left-and-right system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeLeftRight
func (cc _CursorClass) ResizeLeftRightCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("resizeLeftRightCursor"))
	return rv
}

// Returns the resize-right system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeRight
func (cc _CursorClass) ResizeRightCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("resizeRightCursor"))
	return rv
}

// Returns the resize-up system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeUp
func (cc _CursorClass) ResizeUpCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("resizeUpCursor"))
	return rv
}

// Returns the resize-up-and-down system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeUpDown
func (cc _CursorClass) ResizeUpDownCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("resizeUpDownCursor"))
	return rv
}

// Returns the default cursor, the arrow cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/arrow
func (c_ Cursor) ArrowCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("arrowCursor"))
	return rv
}


// Returns the closed-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/closedHand
func (c_ Cursor) ClosedHandCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("closedHandCursor"))
	return rv
}


// Returns the contextual menu system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/contextualMenu
func (c_ Cursor) ContextualMenuCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("contextualMenuCursor"))
	return rv
}


// Returns the cross-hair system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/crosshair
func (c_ Cursor) CrosshairCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("crosshairCursor"))
	return rv
}


// Returns a cursor indicating that the current operation will result in a disappearing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/disappearingItem
func (c_ Cursor) DisappearingItemCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("disappearingItemCursor"))
	return rv
}


// Returns a cursor indicating that the current operation will result in a copy action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragCopy
func (c_ Cursor) DragCopyCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("dragCopyCursor"))
	return rv
}


// Returns a cursor indicating that the current operation will result in a link action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragLink
func (c_ Cursor) DragLinkCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("dragLinkCursor"))
	return rv
}


// Returns a cursor that looks like a capital I with a tiny crossbeam at its middle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeam
func (c_ Cursor) IBeamCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("IBeamCursor"))
	return rv
}


// Returns the cursor for editing vertical layout text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeamCursorForVerticalLayout
func (c_ Cursor) IBeamCursorForVerticalLayout() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("IBeamCursorForVerticalLayout"))
	return rv
}


// Returns the open-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/openHand
func (c_ Cursor) OpenHandCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("openHandCursor"))
	return rv
}


// Returns the operation not allowed cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/operationNotAllowed
func (c_ Cursor) OperationNotAllowedCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("operationNotAllowedCursor"))
	return rv
}


// Returns the pointing-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/pointingHand
func (c_ Cursor) PointingHandCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("pointingHandCursor"))
	return rv
}


// Returns the resize-down system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeDown
func (c_ Cursor) ResizeDownCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("resizeDownCursor"))
	return rv
}


// Returns the resize-left system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeLeft
func (c_ Cursor) ResizeLeftCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("resizeLeftCursor"))
	return rv
}


// Returns the resize-left-and-right system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeLeftRight
func (c_ Cursor) ResizeLeftRightCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("resizeLeftRightCursor"))
	return rv
}


// Returns the resize-right system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeRight
func (c_ Cursor) ResizeRightCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("resizeRightCursor"))
	return rv
}


// Returns the resize-up system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeUp
func (c_ Cursor) ResizeUpCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("resizeUpCursor"))
	return rv
}


// Returns the resize-up-and-down system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/resizeUpDown
func (c_ Cursor) ResizeUpDownCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("resizeUpDownCursor"))
	return rv
}


// The position of the click location within the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/hotspot
func (c_ Cursor) HotSpot() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](c_.ID, objc.Sel("hotSpot"))
	return rv
}


// The position of the click location within the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/hotspot
func (c_ Cursor) SetHotSpot(value coregraphics.CGPoint) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHotSpot:"), value)
}


// The cursor’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/image
func (c_ Cursor) Image() IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("image"))
	return rv
}


// The cursor’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/image
func (c_ Cursor) SetImage(value IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImage:"), value)
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



