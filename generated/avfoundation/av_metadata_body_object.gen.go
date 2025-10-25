// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetadataBodyObject */


/* debug [class_header]: Header for AVMetadataBodyObject */
// The class instance for the [MetadataBodyObject] class.
var (
	MetadataBodyObjectClass     _MetadataBodyObjectClass
	MetadataBodyObjectClassOnce sync.Once
)

func getMetadataBodyObjectClass() _MetadataBodyObjectClass {
	MetadataBodyObjectClassOnce.Do(func() {
		MetadataBodyObjectClass = _MetadataBodyObjectClass{objc.GetClass("AVMetadataBodyObject")}
	})
	return MetadataBodyObjectClass
}

type _MetadataBodyObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetadataBodyObject */
// An interface definition for the [MetadataBodyObject] class.
type IMetadataBodyObject interface {
	IMetadataObject
	
/* debug [class_interface_properties]: Properties for MetadataBodyObject */
	// properties:
	BodyID() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetadataBodyObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetadataBodyObject */
// Alloc allocates a new instance without initialization.
func (mc _MetadataBodyObjectClass) Alloc() MetadataBodyObject {
	rv := objc.Send[MetadataBodyObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataBodyObjectClass) New() MetadataBodyObject {
	rv := objc.Send[MetadataBodyObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataBodyObject) Init() MetadataBodyObject {
	rv := objc.Send[MetadataBodyObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataBodyObject) Autorelease() MetadataBodyObject {
	rv := objc.Send[MetadataBodyObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataBodyObject creates a new MetadataBodyObject instance.
func NewMetadataBodyObject() MetadataBodyObject {
	return getMetadataBodyObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetadataBodyObject */
// An abstract class that defines the interface for a metadata body object.
//
// A metadata body object represents a single detected body in a picture. It’s the base object used to represent bodies, for example , , and .


// An abstract class that defines the interface for a metadata body object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataBodyObject
type MetadataBodyObject struct {
	MetadataObject
}

// MetadataBodyObjectFrom constructs a [MetadataBodyObject] from an unsafe.Pointer.
//
// An abstract class that defines the interface for a metadata body object.
func MetadataBodyObjectFrom(ptr unsafe.Pointer) MetadataBodyObject {
	return MetadataBodyObject{
		MetadataObject: MetadataObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetadataBodyObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetadataBodyObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetadataBodyObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetadataBodyObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetadataBodyObject */

// An integer value that defines the unique identifier of an object in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataBodyObject/bodyID
func (m_ MetadataBodyObject) BodyID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("bodyID"))
	return rv
}/* debug [instance_properties/getter]: bodyID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetadataBodyObject */



