// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetadataDogBodyObject */


/* debug [class_header]: Header for AVMetadataDogBodyObject */
// The class instance for the [MetadataDogBodyObject] class.
var (
	MetadataDogBodyObjectClass     _MetadataDogBodyObjectClass
	MetadataDogBodyObjectClassOnce sync.Once
)

func getMetadataDogBodyObjectClass() _MetadataDogBodyObjectClass {
	MetadataDogBodyObjectClassOnce.Do(func() {
		MetadataDogBodyObjectClass = _MetadataDogBodyObjectClass{objc.GetClass("AVMetadataDogBodyObject")}
	})
	return MetadataDogBodyObjectClass
}

type _MetadataDogBodyObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetadataDogBodyObject */
// An interface definition for the [MetadataDogBodyObject] class.
type IMetadataDogBodyObject interface {
	IMetadataBodyObject
	
/* debug [class_interface_properties]: Properties for MetadataDogBodyObject */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetadataDogBodyObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetadataDogBodyObject */
// Alloc allocates a new instance without initialization.
func (mc _MetadataDogBodyObjectClass) Alloc() MetadataDogBodyObject {
	rv := objc.Send[MetadataDogBodyObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataDogBodyObjectClass) New() MetadataDogBodyObject {
	rv := objc.Send[MetadataDogBodyObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataDogBodyObject) Init() MetadataDogBodyObject {
	rv := objc.Send[MetadataDogBodyObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataDogBodyObject) Autorelease() MetadataDogBodyObject {
	rv := objc.Send[MetadataDogBodyObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataDogBodyObject creates a new MetadataDogBodyObject instance.
func NewMetadataDogBodyObject() MetadataDogBodyObject {
	return getMetadataDogBodyObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetadataDogBodyObject */
// An object representing a single detected dog body in a picture.
//
// This object is an immutable type that describes the various features found in the dog body in a picture.


// An object representing a single detected dog body in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataDogBodyObject
type MetadataDogBodyObject struct {
	MetadataBodyObject
}

// MetadataDogBodyObjectFrom constructs a [MetadataDogBodyObject] from an unsafe.Pointer.
//
// An object representing a single detected dog body in a picture.
func MetadataDogBodyObjectFrom(ptr unsafe.Pointer) MetadataDogBodyObject {
	return MetadataDogBodyObject{
		MetadataBodyObject: MetadataBodyObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetadataDogBodyObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetadataDogBodyObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetadataDogBodyObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetadataDogBodyObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetadataDogBodyObject */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetadataDogBodyObject */



