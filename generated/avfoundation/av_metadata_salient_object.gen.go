// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetadataSalientObject */


/* debug [class_header]: Header for AVMetadataSalientObject */
// The class instance for the [MetadataSalientObject] class.
var (
	MetadataSalientObjectClass     _MetadataSalientObjectClass
	MetadataSalientObjectClassOnce sync.Once
)

func getMetadataSalientObjectClass() _MetadataSalientObjectClass {
	MetadataSalientObjectClassOnce.Do(func() {
		MetadataSalientObjectClass = _MetadataSalientObjectClass{objc.GetClass("AVMetadataSalientObject")}
	})
	return MetadataSalientObjectClass
}

type _MetadataSalientObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetadataSalientObject */
// An interface definition for the [MetadataSalientObject] class.
type IMetadataSalientObject interface {
	IMetadataObject
	
/* debug [class_interface_properties]: Properties for MetadataSalientObject */
	// properties:
	ObjectID() int
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetadataSalientObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetadataSalientObject */
// Alloc allocates a new instance without initialization.
func (mc _MetadataSalientObjectClass) Alloc() MetadataSalientObject {
	rv := objc.Send[MetadataSalientObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataSalientObjectClass) New() MetadataSalientObject {
	rv := objc.Send[MetadataSalientObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataSalientObject) Init() MetadataSalientObject {
	rv := objc.Send[MetadataSalientObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataSalientObject) Autorelease() MetadataSalientObject {
	rv := objc.Send[MetadataSalientObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataSalientObject creates a new MetadataSalientObject instance.
func NewMetadataSalientObject() MetadataSalientObject {
	return getMetadataSalientObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetadataSalientObject */
// An object representing a single salient area in a picture.
//
// This object is an immutable type that describes the various features of the salient object in a picture.


// An object representing a single salient area in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataSalientObject
type MetadataSalientObject struct {
	MetadataObject
}

// MetadataSalientObjectFrom constructs a [MetadataSalientObject] from an unsafe.Pointer.
//
// An object representing a single salient area in a picture.
func MetadataSalientObjectFrom(ptr unsafe.Pointer) MetadataSalientObject {
	return MetadataSalientObject{
		MetadataObject: MetadataObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetadataSalientObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetadataSalientObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetadataSalientObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetadataSalientObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetadataSalientObject */

// An integer value that defines the unique identifier of an object in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataSalientObject/objectID
func (m_ MetadataSalientObject) ObjectID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("objectID"))
	return rv
}/* debug [instance_properties/getter]: objectID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetadataSalientObject */



