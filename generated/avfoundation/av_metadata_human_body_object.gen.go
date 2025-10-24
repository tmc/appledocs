// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetadataHumanBodyObject */


/* debug [class_header]: Header for AVMetadataHumanBodyObject */
// The class instance for the [MetadataHumanBodyObject] class.
var (
	MetadataHumanBodyObjectClass     _MetadataHumanBodyObjectClass
	MetadataHumanBodyObjectClassOnce sync.Once
)

func getMetadataHumanBodyObjectClass() _MetadataHumanBodyObjectClass {
	MetadataHumanBodyObjectClassOnce.Do(func() {
		MetadataHumanBodyObjectClass = _MetadataHumanBodyObjectClass{objc.GetClass("AVMetadataHumanBodyObject")}
	})
	return MetadataHumanBodyObjectClass
}

type _MetadataHumanBodyObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetadataHumanBodyObject */
// An interface definition for the [MetadataHumanBodyObject] class.
type IMetadataHumanBodyObject interface {
	IMetadataBodyObject
	
/* debug [class_interface_properties]: Properties for MetadataHumanBodyObject */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetadataHumanBodyObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetadataHumanBodyObject */
// Alloc allocates a new instance without initialization.
func (mc _MetadataHumanBodyObjectClass) Alloc() MetadataHumanBodyObject {
	rv := objc.Send[MetadataHumanBodyObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataHumanBodyObjectClass) New() MetadataHumanBodyObject {
	rv := objc.Send[MetadataHumanBodyObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataHumanBodyObject) Init() MetadataHumanBodyObject {
	rv := objc.Send[MetadataHumanBodyObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataHumanBodyObject) Autorelease() MetadataHumanBodyObject {
	rv := objc.Send[MetadataHumanBodyObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataHumanBodyObject creates a new MetadataHumanBodyObject instance.
func NewMetadataHumanBodyObject() MetadataHumanBodyObject {
	return getMetadataHumanBodyObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetadataHumanBodyObject */
// An object representing a single detected human body in a picture.
//
// This object is an immutable type that describes the various features found in the human body in a picture.


// An object representing a single detected human body in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataHumanBodyObject
type MetadataHumanBodyObject struct {
	MetadataBodyObject
}

// MetadataHumanBodyObjectFrom constructs a [MetadataHumanBodyObject] from an unsafe.Pointer.
//
// An object representing a single detected human body in a picture.
func MetadataHumanBodyObjectFrom(ptr unsafe.Pointer) MetadataHumanBodyObject {
	return MetadataHumanBodyObject{
		MetadataBodyObject: MetadataBodyObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetadataHumanBodyObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetadataHumanBodyObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetadataHumanBodyObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetadataHumanBodyObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetadataHumanBodyObject */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetadataHumanBodyObject */



