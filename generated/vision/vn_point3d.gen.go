// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class VNPoint3D */


/* debug [class_header]: Header for VNPoint3D */
// The class instance for the [Point3D] class.
var (
	Point3DClass     _Point3DClass
	Point3DClassOnce sync.Once
)

func getPoint3DClass() _Point3DClass {
	Point3DClassOnce.Do(func() {
		Point3DClass = _Point3DClass{objc.GetClass("VNPoint3D")}
	})
	return Point3DClass
}

type _Point3DClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Point3D */
// An interface definition for the [Point3D] class.
type IPoint3D interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Point3D */
	// properties:
	Position() objectivec.IObject
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Point3D */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Point3D */
// Alloc allocates a new instance without initialization.
func (pc _Point3DClass) Alloc() Point3D {
	rv := objc.Send[Point3D](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _Point3DClass) New() Point3D {
	rv := objc.Send[Point3D](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Point3D) Init() Point3D {
	rv := objc.Send[Point3D](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Point3D) Autorelease() Point3D {
	rv := objc.Send[Point3D](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPoint3D creates a new Point3D instance.
func NewPoint3D() Point3D {
	return getPoint3DClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Point3D */
// An object that represents a 3D point in an image.


// An object that represents a 3D point in an image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint3D
type Point3D struct {
	objectivec.Object
}

// Point3DFrom constructs a [Point3D] from an unsafe.Pointer.
//
// An object that represents a 3D point in an image.
func Point3DFrom(ptr unsafe.Pointer) Point3D {
	return Point3D{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Point3D */

// Creates a point object with the position you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint3D/init(position:)
func NewPoint3DWithPosition(position objectivec.IObject) Point3D {
	instance := getPoint3DClass().Alloc()
	rv := objc.Send[Point3D](instance.ID, objc.Sel("initWithPosition:"), position)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPoint3DWithPosition */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Point3D */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Point3D */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Point3D */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Point3D */

// The three-dimensional position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Vision/VNPoint3D/position
func (p_ Point3D) Position() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class VNPoint3D */


