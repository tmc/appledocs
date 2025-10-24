// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSCursor */


/* debug [class_header]: Header for NSCursor */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Cursor */
// An interface definition for the [Cursor] class.
type ICursor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Cursor */
	// properties:
	HotSpot() vision.Point
	Image() IImage
	SetOnMouseEntered() bool
	SetOnMouseExited() bool
	IsSetOnMouseEntered() bool
	SetIsSetOnMouseEntered(value bool)
	IsSetOnMouseExited() bool
	SetIsSetOnMouseExited(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Cursor */
	// methods:
	Pop()
	Push()
	Set()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Cursor */
// Alloc allocates a new instance without initialization.
func (cc _CursorClass) Alloc() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Cursor */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Cursor */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/init(coder:)
func NewCursorWithCoder(coder foundation.Coder) Cursor {
	instance := getCursorClass().Alloc()
	rv := objc.Send[Cursor](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCursorWithCoder */


// Initializes the cursor with the specified image and hot spot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/init(image:foregroundColorHint:backgroundColorHint:hotSpot:)
func NewCursorWithImageForegroundColorHintBackgroundColorHintHotSpot(newImage IImage, fg IColor, bg IColor, hotSpot vision.Point) Cursor {
	instance := getCursorClass().Alloc()
	rv := objc.Send[Cursor](instance.ID, objc.Sel("initWithImage:foregroundColorHint:backgroundColorHint:hotSpot:"), newImage, fg, bg, hotSpot)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCursorWithImageForegroundColorHintBackgroundColorHintHotSpot */


// Initializes a cursor with the given image and hot spot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/init(image:hotSpot:)
func NewCursorWithImageHotSpot(newImage IImage, point vision.Point) Cursor {
	instance := getCursorClass().Alloc()
	rv := objc.Send[Cursor](instance.ID, objc.Sel("initWithImage:hotSpot:"), newImage, point)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCursorWithImageHotSpot */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Cursor */

// Makes the current cursor invisible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/hide()
func (cc _CursorClass) Hide() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("hide"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Hide) */


// Pops the current cursor off the top of the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/pop()-swift.type.method
func (cc _CursorClass) Pop() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("pop"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Pop) */


// Sets whether the cursor is hidden until the mouse moves.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/setHiddenUntilMouseMoves(_:)
func (cc _CursorClass) SetHiddenUntilMouseMoves(flag bool) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("setHiddenUntilMouseMoves:"), flag)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetHiddenUntilMouseMoves) */


// Negates an earlier call to by showing the current cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/unhide()
func (cc _CursorClass) Unhide() {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("unhide"))
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Unhide) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Cursor */

// Returns the default cursor, the arrow cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/arrow
func (cc _CursorClass) ArrowCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("arrowCursor"))
	return rv
}/* debug [class_properties_class/property]: arrowCursor */

// Returns the closed-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/closedHand
func (cc _CursorClass) ClosedHandCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("closedHandCursor"))
	return rv
}/* debug [class_properties_class/property]: closedHandCursor */

// Returns the contextual menu system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/contextualMenu
func (cc _CursorClass) ContextualMenuCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("contextualMenuCursor"))
	return rv
}/* debug [class_properties_class/property]: contextualMenuCursor */

// Returns the cross-hair system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/crosshair
func (cc _CursorClass) CrosshairCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("crosshairCursor"))
	return rv
}/* debug [class_properties_class/property]: crosshairCursor */

// Returns the application’s current cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/current
func (cc _CursorClass) CurrentCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("currentCursor"))
	return rv
}/* debug [class_properties_class/property]: currentCursor */

// Returns the current system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/currentSystem
func (cc _CursorClass) CurrentSystemCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("currentSystemCursor"))
	return rv
}/* debug [class_properties_class/property]: currentSystemCursor */

// Returns a cursor indicating that the current operation will result in a disappearing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/disappearingItem
func (cc _CursorClass) DisappearingItemCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("disappearingItemCursor"))
	return rv
}/* debug [class_properties_class/property]: disappearingItemCursor */

// Returns a cursor indicating that the current operation will result in a copy action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragCopy
func (cc _CursorClass) DragCopyCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("dragCopyCursor"))
	return rv
}/* debug [class_properties_class/property]: dragCopyCursor */

// Returns a cursor indicating that the current operation will result in a link action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragLink
func (cc _CursorClass) DragLinkCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("dragLinkCursor"))
	return rv
}/* debug [class_properties_class/property]: dragLinkCursor */

// Returns a cursor that looks like a capital I with a tiny crossbeam at its middle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeam
func (cc _CursorClass) IBeamCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("IBeamCursor"))
	return rv
}/* debug [class_properties_class/property]: IBeamCursor */

// Returns the cursor for editing vertical layout text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeamCursorForVerticalLayout
func (cc _CursorClass) IBeamCursorForVerticalLayout() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("IBeamCursorForVerticalLayout"))
	return rv
}/* debug [class_properties_class/property]: IBeamCursorForVerticalLayout */

// Returns the open-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/openHand
func (cc _CursorClass) OpenHandCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("openHandCursor"))
	return rv
}/* debug [class_properties_class/property]: openHandCursor */

// Returns the operation not allowed cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/operationNotAllowed
func (cc _CursorClass) OperationNotAllowedCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("operationNotAllowedCursor"))
	return rv
}/* debug [class_properties_class/property]: operationNotAllowedCursor */

// Returns the pointing-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/pointingHand
func (cc _CursorClass) PointingHandCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("pointingHandCursor"))
	return rv
}/* debug [class_properties_class/property]: pointingHandCursor */

// Returns the zoom-in cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/zoomIn
func (cc _CursorClass) ZoomInCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("zoomInCursor"))
	return rv
}/* debug [class_properties_class/property]: zoomInCursor */

// Returns the zoom-out cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/zoomOut
func (cc _CursorClass) ZoomOutCursor() Cursor {
	rv := objc.Send[Cursor](objc.ID(cc.class), objc.Sel("zoomOutCursor"))
	return rv
}/* debug [class_properties_class/property]: zoomOutCursor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Cursor */

// Sends a message to the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/pop()-swift.method
func (c_ Cursor) Pop() {
	objc.Send[objc.ID](c_.ID, objc.Sel("pop"))
}/* debug [instance_methods/method]: Pop */


// Puts the receiver on top of the cursor stack and makes it the current cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/push()
func (c_ Cursor) Push() {
	objc.Send[objc.ID](c_.ID, objc.Sel("push"))
}/* debug [instance_methods/method]: Push */


// Makes the receiver the current cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/set()
func (c_ Cursor) Set() {
	objc.Send[objc.ID](c_.ID, objc.Sel("set"))
}/* debug [instance_methods/method]: Set */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Cursor */

// Returns the default cursor, the arrow cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/arrow
func (c_ Cursor) ArrowCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("arrowCursor"))
	return rv
}/* debug [instance_properties/getter]: arrowCursor */


// Returns the closed-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/closedHand
func (c_ Cursor) ClosedHandCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("closedHandCursor"))
	return rv
}/* debug [instance_properties/getter]: closedHandCursor */


// Returns the contextual menu system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/contextualMenu
func (c_ Cursor) ContextualMenuCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("contextualMenuCursor"))
	return rv
}/* debug [instance_properties/getter]: contextualMenuCursor */


// Returns the cross-hair system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/crosshair
func (c_ Cursor) CrosshairCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("crosshairCursor"))
	return rv
}/* debug [instance_properties/getter]: crosshairCursor */


// Returns the application’s current cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/current
func (c_ Cursor) CurrentCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("currentCursor"))
	return rv
}/* debug [instance_properties/getter]: currentCursor */


// Returns the current system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/currentSystem
func (c_ Cursor) CurrentSystemCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("currentSystemCursor"))
	return rv
}/* debug [instance_properties/getter]: currentSystemCursor */


// Returns a cursor indicating that the current operation will result in a disappearing item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/disappearingItem
func (c_ Cursor) DisappearingItemCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("disappearingItemCursor"))
	return rv
}/* debug [instance_properties/getter]: disappearingItemCursor */


// Returns a cursor indicating that the current operation will result in a copy action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragCopy
func (c_ Cursor) DragCopyCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("dragCopyCursor"))
	return rv
}/* debug [instance_properties/getter]: dragCopyCursor */


// Returns a cursor indicating that the current operation will result in a link action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/dragLink
func (c_ Cursor) DragLinkCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("dragLinkCursor"))
	return rv
}/* debug [instance_properties/getter]: dragLinkCursor */


// The position of the click location within the cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/hotSpot
func (c_ Cursor) HotSpot() vision.Point {
	rv := objc.Send[vision.Point](c_.ID, objc.Sel("hotSpot"))
	return rv
}/* debug [instance_properties/getter]: hotSpot */


// Returns a cursor that looks like a capital I with a tiny crossbeam at its middle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeam
func (c_ Cursor) IBeamCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("IBeamCursor"))
	return rv
}/* debug [instance_properties/getter]: IBeamCursor */


// Returns the cursor for editing vertical layout text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/iBeamCursorForVerticalLayout
func (c_ Cursor) IBeamCursorForVerticalLayout() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("IBeamCursorForVerticalLayout"))
	return rv
}/* debug [instance_properties/getter]: IBeamCursorForVerticalLayout */


// The cursor’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/image
func (c_ Cursor) Image() IImage {
	rv := objc.Send[Image](c_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// A Boolean value indicating whether the receiver becomes current on receiving a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/isSetOnMouseEntered
func (c_ Cursor) SetOnMouseEntered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setOnMouseEntered"))
	return rv
}/* debug [instance_properties/getter]: setOnMouseEntered */


// A Boolean value indicating whether the receiver becomes current when it receives a message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/isSetOnMouseExited
func (c_ Cursor) SetOnMouseExited() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setOnMouseExited"))
	return rv
}/* debug [instance_properties/getter]: setOnMouseExited */


// Returns the open-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/openHand
func (c_ Cursor) OpenHandCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("openHandCursor"))
	return rv
}/* debug [instance_properties/getter]: openHandCursor */


// Returns the operation not allowed cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/operationNotAllowed
func (c_ Cursor) OperationNotAllowedCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("operationNotAllowedCursor"))
	return rv
}/* debug [instance_properties/getter]: operationNotAllowedCursor */


// Returns the pointing-hand system cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/pointingHand
func (c_ Cursor) PointingHandCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("pointingHandCursor"))
	return rv
}/* debug [instance_properties/getter]: pointingHandCursor */


// Returns the zoom-in cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/zoomIn
func (c_ Cursor) ZoomInCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("zoomInCursor"))
	return rv
}/* debug [instance_properties/getter]: zoomInCursor */


// Returns the zoom-out cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/zoomOut
func (c_ Cursor) ZoomOutCursor() ICursor {
	rv := objc.Send[Cursor](c_.ID, objc.Sel("zoomOutCursor"))
	return rv
}/* debug [instance_properties/getter]: zoomOutCursor */


// A Boolean value indicating whether the receiver becomes current on receiving a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/issetonmouseentered
func (c_ Cursor) IsSetOnMouseEntered() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSetOnMouseEntered"))
	return rv
}/* debug [instance_properties/getter]: isSetOnMouseEntered */


// A Boolean value indicating whether the receiver becomes current on receiving a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/issetonmouseentered
func (c_ Cursor) SetIsSetOnMouseEntered(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSetOnMouseEntered:"), value)
}/* debug [instance_properties/setter]: isSetOnMouseEntered */


// A Boolean value indicating whether the receiver becomes current when it receives a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/issetonmouseexited
func (c_ Cursor) IsSetOnMouseExited() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSetOnMouseExited"))
	return rv
}/* debug [instance_properties/getter]: isSetOnMouseExited */


// A Boolean value indicating whether the receiver becomes current when it receives a
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscursor/issetonmouseexited
func (c_ Cursor) SetIsSetOnMouseExited(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSetOnMouseExited:"), value)
}/* debug [instance_properties/setter]: isSetOnMouseExited */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCursor */


