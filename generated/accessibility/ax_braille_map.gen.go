// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXBrailleMap */


/* debug [class_header]: Header for AXBrailleMap */
// The class instance for the [AXBrailleMap] class.
var (
	AXBrailleMapClass     _AXBrailleMapClass
	AXBrailleMapClassOnce sync.Once
)

func getAXBrailleMapClass() _AXBrailleMapClass {
	AXBrailleMapClassOnce.Do(func() {
		AXBrailleMapClass = _AXBrailleMapClass{objc.GetClass("AXBrailleMap")}
	})
	return AXBrailleMapClass
}

type _AXBrailleMapClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXBrailleMap */
// An interface definition for the [AXBrailleMap] class.
type IAXBrailleMap interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXBrailleMap */
	// properties:
	Dimensions() corefoundation.CGSize
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXBrailleMap */
	// methods:
	HeightAtPoint(point corefoundation.CGPoint) float32
	PresentImage(image ImageRef /* not a class type */)
	SetHeightAtPoint(status float32, point corefoundation.CGPoint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXBrailleMap */
// Alloc allocates a new instance without initialization.
func (ac _AXBrailleMapClass) Alloc() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AXBrailleMapClass) New() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXBrailleMap) Init() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXBrailleMap) Autorelease() AXBrailleMap {
	rv := objc.Send[AXBrailleMap](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXBrailleMap creates a new AXBrailleMap instance.
func NewAXBrailleMap() AXBrailleMap {
	return getAXBrailleMapClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXBrailleMap */
// A representation of a two-dimensional braille display.
//
// A braille map object represents a two-dimensional braille display that’s connected to the current Apple device. By specifying the dot patterns in the braille map, you can change the content the user experiences. To render the data from the braille map to the display, implement .


// A representation of a two-dimensional braille display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleMap
type AXBrailleMap struct {
	objectivec.Object
}

// AXBrailleMapFrom constructs a [AXBrailleMap] from an unsafe.Pointer.
//
// A representation of a two-dimensional braille display.
func AXBrailleMapFrom(ptr unsafe.Pointer) AXBrailleMap {
	return AXBrailleMap{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXBrailleMap *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXBrailleMap */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXBrailleMap */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXBrailleMap */

// Retrieves the height of an individual pin on the braille display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/height(at:)
func (a_ AXBrailleMap) HeightAtPoint(point corefoundation.CGPoint) float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("heightAtPoint:"), point)
	return rv
}/* debug [instance_methods/method]: HeightAtPoint */


// Converts the data from the image you specify into the braille map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/present(_:)
func (a_ AXBrailleMap) PresentImage(image ImageRef /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("presentImage:"), image)
}/* debug [instance_methods/method]: PresentImage */


// Sets the height of an individual pin on the braille display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/setHeight(_:at:)
func (a_ AXBrailleMap) SetHeightAtPoint(status float32, point corefoundation.CGPoint) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setHeight:atPoint:"), status, point)
}/* debug [instance_methods/method]: SetHeightAtPoint */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXBrailleMap */

// The number of pins in each dimension of the braille display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleMap/dimensions
func (a_ AXBrailleMap) Dimensions() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("dimensions"))
	return rv
}/* debug [instance_properties/getter]: dimensions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXBrailleMap */



