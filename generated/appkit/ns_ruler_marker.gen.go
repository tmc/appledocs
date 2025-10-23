// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RulerMarker] class.
var (
	RulerMarkerClass     _RulerMarkerClass
	RulerMarkerClassOnce sync.Once
)

func getRulerMarkerClass() _RulerMarkerClass {
	RulerMarkerClassOnce.Do(func() {
		RulerMarkerClass = _RulerMarkerClass{objc.GetClass("NSRulerMarker")}
	})
	return RulerMarkerClass
}

type _RulerMarkerClass struct {
	class objc.Class
}

// An interface definition for the [RulerMarker] class.
type IRulerMarker interface {
	objectivec.IObject
	// properties:
	Image() IImage
	SetImage(value IImage)
	ImageOrigin() coregraphics.CGPoint
	SetImageOrigin(value coregraphics.CGPoint)
	ImageRectInRuler() coregraphics.CGRect
	Dragging() bool /* primitive/slice/pointer. */
	Movable() bool /* primitive/slice/pointer. */
	SetMovable(value bool /* primitive/slice/pointer. */)
	Removable() bool /* primitive/slice/pointer. */
	SetRemovable(value bool /* primitive/slice/pointer. */)
	MarkerLocation() float64 /* primitive/slice/pointer. */
	SetMarkerLocation(value float64 /* primitive/slice/pointer. */)
	RepresentedObject() objc.ID
	SetRepresentedObject(value objc.ID)
	Ruler() IRulerView
	ThicknessRequiredInRuler() float64 /* primitive/slice/pointer. */
	IsDragging() bool /* primitive/slice/pointer. */
	SetIsDragging(value bool /* primitive/slice/pointer. */)
	IsMovable() bool /* primitive/slice/pointer. */
	SetIsMovable(value bool /* primitive/slice/pointer. */)
	IsRemovable() bool /* primitive/slice/pointer. */
	SetIsRemovable(value bool /* primitive/slice/pointer. */)
	// methods:
	DrawRect(rect coregraphics.CGRect)
	TrackMouseAdding(mouseDownEvent IEvent, isAdding bool /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
}

// A symbol on a ruler view, indicating a location for the graphics element it represents in the client of the ruler view.
//
// An example of a marker is the representation of a margin or tab setting, or the edges of a graphic on the page.


// A symbol on a ruler view, indicating a location for the graphics element it represents in the client of the ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker
type RulerMarker struct {
	objectivec.Object
}

// RulerMarkerFrom constructs a [RulerMarker] from an unsafe.Pointer.
//
// A symbol on a ruler view, indicating a location for the graphics element it represents in the client of the ruler view.
func RulerMarkerFrom(ptr unsafe.Pointer) RulerMarker {
	return RulerMarker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RulerMarkerClass) Alloc() RulerMarker {
	rv := objc.Send[RulerMarker](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RulerMarkerClass) New() RulerMarker {
	rv := objc.Send[RulerMarker](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RulerMarker) Init() RulerMarker {
	rv := objc.Send[RulerMarker](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RulerMarker) Autorelease() RulerMarker {
	rv := objc.Send[RulerMarker](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRulerMarker creates a new RulerMarker instance.
func NewRulerMarker() RulerMarker {
	return getRulerMarkerClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/init(coder:)
func NewRulerMarkerWithCoder(coder Coder /* not a class type */) RulerMarker {
	instance := getRulerMarkerClass().Alloc()
	rv := objc.Send[RulerMarker](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated ruler marker, associating it with (but not adding it to) a specified ruler view and assigning the attributes provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/init(rulerView:markerLocation:image:imageOrigin:)
func NewRulerMarkerWithRulerViewMarkerLocationImageImageOrigin(ruler IRulerView, location float64 /* primitive/slice/pointer. */, image IImage, imageOrigin coregraphics.CGPoint) RulerMarker {
	instance := getRulerMarkerClass().Alloc()
	rv := objc.Send[RulerMarker](instance.ID, objc.Sel("initWithRulerView:markerLocation:image:imageOrigin:"), ruler, location, image, imageOrigin)
	rv.Autorelease()
	return rv
}



// Draws the receiver’s image that appears in the supplied rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/draw(_:)
func (r_ RulerMarker) DrawRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](r_.ID, objc.Sel("drawRect:"), rect)
}


// Handles user manipulation of the receiver in its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/trackMouse(with:adding:)
func (r_ RulerMarker) TrackMouseAdding(mouseDownEvent IEvent, isAdding bool /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("trackMouse:adding:"), mouseDownEvent, isAdding)
	return rv
}


// The receiver’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/image
func (r_ RulerMarker) Image() IImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("image"))
	return rv
}


// The receiver’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/image
func (r_ RulerMarker) SetImage(value IImage) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setImage:"), value)
}


// The point in the receiver’s image that is positioned at the receiver’s location on the ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/imageOrigin
func (r_ RulerMarker) ImageOrigin() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](r_.ID, objc.Sel("imageOrigin"))
	return rv
}


// The point in the receiver’s image that is positioned at the receiver’s location on the ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/imageOrigin
func (r_ RulerMarker) SetImageOrigin(value coregraphics.CGPoint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setImageOrigin:"), value)
}


// The rectangle occupied by the receiver’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/imageRectInRuler
func (r_ RulerMarker) ImageRectInRuler() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](r_.ID, objc.Sel("imageRectInRuler"))
	return rv
}


// A Boolean that indicates whether the receiver is being dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isDragging
func (r_ RulerMarker) Dragging() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("dragging"))
	return rv
}


// A Boolean that indicates whether the user can move the receiver in its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isMovable
func (r_ RulerMarker) Movable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("movable"))
	return rv
}


// A Boolean that indicates whether the user can move the receiver in its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isMovable
func (r_ RulerMarker) SetMovable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMovable:"), value)
}


// A Boolean that indicates whether the user can remove the receiver from its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isRemovable
func (r_ RulerMarker) Removable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("removable"))
	return rv
}


// A Boolean that indicates whether the user can remove the receiver from its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isRemovable
func (r_ RulerMarker) SetRemovable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRemovable:"), value)
}


// The location of the receiver in the coordinate system of the ruler view’s client view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/markerLocation
func (r_ RulerMarker) MarkerLocation() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](r_.ID, objc.Sel("markerLocation"))
	return rv
}


// The location of the receiver in the coordinate system of the ruler view’s client view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/markerLocation
func (r_ RulerMarker) SetMarkerLocation(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMarkerLocation:"), value)
}


// The object the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/representedObject
func (r_ RulerMarker) RepresentedObject() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("representedObject"))
	return rv
}


// The object the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/representedObject
func (r_ RulerMarker) SetRepresentedObject(value objc.ID) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRepresentedObject:"), value)
}


// The receiver’s ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/ruler
func (r_ RulerMarker) Ruler() IRulerView {
	rv := objc.Send[RulerView](r_.ID, objc.Sel("ruler"))
	return rv
}


// The amount of the receiver’s image that’s displayed above or to the left of the ruler view’s baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/thicknessRequiredInRuler
func (r_ RulerMarker) ThicknessRequiredInRuler() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](r_.ID, objc.Sel("thicknessRequiredInRuler"))
	return rv
}


// A Boolean that indicates whether the receiver is being dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/isdragging
func (r_ RulerMarker) IsDragging() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isDragging"))
	return rv
}


// A Boolean that indicates whether the receiver is being dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/isdragging
func (r_ RulerMarker) SetIsDragging(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsDragging:"), value)
}


// A Boolean that indicates whether the user can move the receiver in its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/ismovable
func (r_ RulerMarker) IsMovable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isMovable"))
	return rv
}


// A Boolean that indicates whether the user can move the receiver in its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/ismovable
func (r_ RulerMarker) SetIsMovable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsMovable:"), value)
}


// A Boolean that indicates whether the user can remove the receiver from its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/isremovable
func (r_ RulerMarker) IsRemovable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isRemovable"))
	return rv
}


// A Boolean that indicates whether the user can remove the receiver from its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/isremovable
func (r_ RulerMarker) SetIsRemovable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsRemovable:"), value)
}


