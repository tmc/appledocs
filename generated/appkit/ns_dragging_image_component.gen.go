// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSDraggingImageComponent */


/* debug [class_header]: Header for NSDraggingImageComponent */
// The class instance for the [DraggingImageComponent] class.
var (
	DraggingImageComponentClass     _DraggingImageComponentClass
	DraggingImageComponentClassOnce sync.Once
)

func getDraggingImageComponentClass() _DraggingImageComponentClass {
	DraggingImageComponentClassOnce.Do(func() {
		DraggingImageComponentClass = _DraggingImageComponentClass{objc.GetClass("NSDraggingImageComponent")}
	})
	return DraggingImageComponentClass
}

type _DraggingImageComponentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for DraggingImageComponent */
// An interface definition for the [DraggingImageComponent] class.
type IDraggingImageComponent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for DraggingImageComponent */
	// properties:
	Contents() objc.ID
	SetContents(value objc.ID)
	Frame() Rect /* not a class type */
	SetFrame(value Rect /* not a class type */)
	Key() DraggingImageComponentKey /* typedef */
	SetKey(value DraggingImageComponentKey /* typedef */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for DraggingImageComponent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for DraggingImageComponent */
// Alloc allocates a new instance without initialization.
func (dc _DraggingImageComponentClass) Alloc() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DraggingImageComponentClass) New() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DraggingImageComponent) Init() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DraggingImageComponent) Autorelease() DraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDraggingImageComponent creates a new DraggingImageComponent instance.
func NewDraggingImageComponent() DraggingImageComponent {
	return getDraggingImageComponentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for DraggingImageComponent */
// A single object in a dragging item.
//
// An array of instances are composited together to create the dragging image for an . instances can simply be considered as named images with a location used by an instance.


// A single object in a dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent
type DraggingImageComponent struct {
	objectivec.Object
}

// DraggingImageComponentFrom constructs a [DraggingImageComponent] from an unsafe.Pointer.
//
// A single object in a dragging item.
func DraggingImageComponentFrom(ptr unsafe.Pointer) DraggingImageComponent {
	return DraggingImageComponent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for DraggingImageComponent */

// Initializes and returns a dragging image component with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/init(key:)
func NewDraggingImageComponentWithKey(key DraggingImageComponentKey /* typedef */) DraggingImageComponent {
	instance := getDraggingImageComponentClass().Alloc()
	rv := objc.Send[DraggingImageComponent](instance.ID, objc.Sel("initWithKey:"), key)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewDraggingImageComponentWithKey */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for DraggingImageComponent */

// Creates and returns a dragging image component with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/draggingImageComponentWithKey:
func (dc _DraggingImageComponentClass) DraggingImageComponentWithKey(key DraggingImageComponentKey /* typedef */) IDraggingImageComponent {
	rv := objc.Send[DraggingImageComponent](objc.ID(dc.class), objc.Sel("draggingImageComponentWithKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DraggingImageComponentWithKey) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for DraggingImageComponent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for DraggingImageComponent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for DraggingImageComponent */

// An object providing the image contents of the component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/contents
func (d_ DraggingImageComponent) Contents() objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// An object providing the image contents of the component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/contents
func (d_ DraggingImageComponent) SetContents(value objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setContents:"), value)
}/* debug [instance_properties/setter]: contents */


// The coordinate space is the bounds of the parent dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/frame
func (d_ DraggingImageComponent) Frame() Rect /* not a class type */ {
	rv := objc.Send[Rect](d_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The coordinate space is the bounds of the parent dragging item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/frame
func (d_ DraggingImageComponent) SetFrame(value Rect /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFrame:"), value)
}/* debug [instance_properties/setter]: frame */


// The unique name of this image component instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/key
func (d_ DraggingImageComponent) Key() DraggingImageComponentKey /* typedef */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("key"))
	return rv
}/* debug [instance_properties/getter]: key */


// The unique name of this image component instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingImageComponent/key
func (d_ DraggingImageComponent) SetKey(value DraggingImageComponentKey /* typedef */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setKey:"), value)
}/* debug [instance_properties/setter]: key */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSDraggingImageComponent */


