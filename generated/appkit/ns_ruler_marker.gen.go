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

/* debug [class.gen.go]: Generating class NSRulerMarker */


/* debug [class_header]: Header for NSRulerMarker */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RulerMarker */
// An interface definition for the [RulerMarker] class.
type IRulerMarker interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RulerMarker */
	// properties:
	Image() IImage
	SetImage(value IImage)
	ImageOrigin() vision.Point
	SetImageOrigin(value vision.Point)
	ImageRectInRuler() Rect /* not a class type */
	Dragging() bool
	Movable() bool
	SetMovable(value bool)
	Removable() bool
	SetRemovable(value bool)
	MarkerLocation() float64
	SetMarkerLocation(value float64)
	RepresentedObject() unsafe.Pointer
	SetRepresentedObject(value unsafe.Pointer)
	Ruler() IRulerView
	ThicknessRequiredInRuler() float64
	IsDragging() bool
	SetIsDragging(value bool)
	IsMovable() bool
	SetIsMovable(value bool)
	IsRemovable() bool
	SetIsRemovable(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RulerMarker */
	// methods:
	DrawRect(rect Rect /* not a class type */)
	TrackMouseAdding(mouseDownEvent IEvent, isAdding bool) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RulerMarker */
// Alloc allocates a new instance without initialization.
func (rc _RulerMarkerClass) Alloc() RulerMarker {
	rv := objc.Send[RulerMarker](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RulerMarker */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RulerMarker */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/init(coder:)
func NewRulerMarkerWithCoder(coder foundation.Coder) RulerMarker {
	instance := getRulerMarkerClass().Alloc()
	rv := objc.Send[RulerMarker](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRulerMarkerWithCoder */


// Initializes a newly allocated ruler marker, associating it with (but not adding it to) a specified ruler view and assigning the attributes provided.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/init(rulerView:markerLocation:image:imageOrigin:)
func NewRulerMarkerWithRulerViewMarkerLocationImageImageOrigin(ruler IRulerView, location float64, image IImage, imageOrigin vision.Point) RulerMarker {
	instance := getRulerMarkerClass().Alloc()
	rv := objc.Send[RulerMarker](instance.ID, objc.Sel("initWithRulerView:markerLocation:image:imageOrigin:"), ruler, location, image, imageOrigin)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRulerMarkerWithRulerViewMarkerLocationImageImageOrigin */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RulerMarker */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RulerMarker */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RulerMarker */

// Draws the receiver’s image that appears in the supplied rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/draw(_:)
func (r_ RulerMarker) DrawRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("drawRect:"), rect)
}/* debug [instance_methods/method]: DrawRect */


// Handles user manipulation of the receiver in its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/trackMouse(with:adding:)
func (r_ RulerMarker) TrackMouseAdding(mouseDownEvent IEvent, isAdding bool) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("trackMouse:adding:"), mouseDownEvent, isAdding)
	return rv
}/* debug [instance_methods/method]: TrackMouseAdding */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RulerMarker */

// The receiver’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/image
func (r_ RulerMarker) Image() IImage {
	rv := objc.Send[Image](r_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// The receiver’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/image
func (r_ RulerMarker) SetImage(value IImage) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setImage:"), value)
}/* debug [instance_properties/setter]: image */


// The point in the receiver’s image that is positioned at the receiver’s location on the ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/imageOrigin
func (r_ RulerMarker) ImageOrigin() vision.Point {
	rv := objc.Send[vision.Point](r_.ID, objc.Sel("imageOrigin"))
	return rv
}/* debug [instance_properties/getter]: imageOrigin */


// The point in the receiver’s image that is positioned at the receiver’s location on the ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/imageOrigin
func (r_ RulerMarker) SetImageOrigin(value vision.Point) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setImageOrigin:"), value)
}/* debug [instance_properties/setter]: imageOrigin */


// The rectangle occupied by the receiver’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/imageRectInRuler
func (r_ RulerMarker) ImageRectInRuler() Rect /* not a class type */ {
	rv := objc.Send[Rect](r_.ID, objc.Sel("imageRectInRuler"))
	return rv
}/* debug [instance_properties/getter]: imageRectInRuler */


// A Boolean that indicates whether the receiver is being dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isDragging
func (r_ RulerMarker) Dragging() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("dragging"))
	return rv
}/* debug [instance_properties/getter]: dragging */


// A Boolean that indicates whether the user can move the receiver in its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isMovable
func (r_ RulerMarker) Movable() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("movable"))
	return rv
}/* debug [instance_properties/getter]: movable */


// A Boolean that indicates whether the user can move the receiver in its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isMovable
func (r_ RulerMarker) SetMovable(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMovable:"), value)
}/* debug [instance_properties/setter]: movable */


// A Boolean that indicates whether the user can remove the receiver from its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isRemovable
func (r_ RulerMarker) Removable() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("removable"))
	return rv
}/* debug [instance_properties/getter]: removable */


// A Boolean that indicates whether the user can remove the receiver from its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/isRemovable
func (r_ RulerMarker) SetRemovable(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRemovable:"), value)
}/* debug [instance_properties/setter]: removable */


// The location of the receiver in the coordinate system of the ruler view’s client view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/markerLocation
func (r_ RulerMarker) MarkerLocation() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("markerLocation"))
	return rv
}/* debug [instance_properties/getter]: markerLocation */


// The location of the receiver in the coordinate system of the ruler view’s client view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/markerLocation
func (r_ RulerMarker) SetMarkerLocation(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMarkerLocation:"), value)
}/* debug [instance_properties/setter]: markerLocation */


// The object the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/representedObject
func (r_ RulerMarker) RepresentedObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("representedObject"))
	return rv
}/* debug [instance_properties/getter]: representedObject */


// The object the receiver represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/representedObject
func (r_ RulerMarker) SetRepresentedObject(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRepresentedObject:"), value)
}/* debug [instance_properties/setter]: representedObject */


// The receiver’s ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/ruler
func (r_ RulerMarker) Ruler() IRulerView {
	rv := objc.Send[RulerView](r_.ID, objc.Sel("ruler"))
	return rv
}/* debug [instance_properties/getter]: ruler */


// The amount of the receiver’s image that’s displayed above or to the left of the ruler view’s baseline.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerMarker/thicknessRequiredInRuler
func (r_ RulerMarker) ThicknessRequiredInRuler() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("thicknessRequiredInRuler"))
	return rv
}/* debug [instance_properties/getter]: thicknessRequiredInRuler */


// A Boolean that indicates whether the receiver is being dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/isdragging
func (r_ RulerMarker) IsDragging() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isDragging"))
	return rv
}/* debug [instance_properties/getter]: isDragging */


// A Boolean that indicates whether the receiver is being dragged.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/isdragging
func (r_ RulerMarker) SetIsDragging(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsDragging:"), value)
}/* debug [instance_properties/setter]: isDragging */


// A Boolean that indicates whether the user can move the receiver in its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/ismovable
func (r_ RulerMarker) IsMovable() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isMovable"))
	return rv
}/* debug [instance_properties/getter]: isMovable */


// A Boolean that indicates whether the user can move the receiver in its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/ismovable
func (r_ RulerMarker) SetIsMovable(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsMovable:"), value)
}/* debug [instance_properties/setter]: isMovable */


// A Boolean that indicates whether the user can remove the receiver from its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/isremovable
func (r_ RulerMarker) IsRemovable() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isRemovable"))
	return rv
}/* debug [instance_properties/getter]: isRemovable */


// A Boolean that indicates whether the user can remove the receiver from its ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulermarker/isremovable
func (r_ RulerMarker) SetIsRemovable(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsRemovable:"), value)
}/* debug [instance_properties/setter]: isRemovable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSRulerMarker */


