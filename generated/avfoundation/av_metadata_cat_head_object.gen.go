// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetadataCatHeadObject */


/* debug [class_header]: Header for AVMetadataCatHeadObject */
// The class instance for the [MetadataCatHeadObject] class.
var (
	MetadataCatHeadObjectClass     _MetadataCatHeadObjectClass
	MetadataCatHeadObjectClassOnce sync.Once
)

func getMetadataCatHeadObjectClass() _MetadataCatHeadObjectClass {
	MetadataCatHeadObjectClassOnce.Do(func() {
		MetadataCatHeadObjectClass = _MetadataCatHeadObjectClass{objc.GetClass("AVMetadataCatHeadObject")}
	})
	return MetadataCatHeadObjectClass
}

type _MetadataCatHeadObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetadataCatHeadObject */
// An interface definition for the [MetadataCatHeadObject] class.
type IMetadataCatHeadObject interface {
	IMetadataObject
	
/* debug [class_interface_properties]: Properties for MetadataCatHeadObject */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetadataCatHeadObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetadataCatHeadObject */
// Alloc allocates a new instance without initialization.
func (mc _MetadataCatHeadObjectClass) Alloc() MetadataCatHeadObject {
	rv := objc.Send[MetadataCatHeadObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataCatHeadObjectClass) New() MetadataCatHeadObject {
	rv := objc.Send[MetadataCatHeadObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataCatHeadObject) Init() MetadataCatHeadObject {
	rv := objc.Send[MetadataCatHeadObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataCatHeadObject) Autorelease() MetadataCatHeadObject {
	rv := objc.Send[MetadataCatHeadObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataCatHeadObject creates a new MetadataCatHeadObject instance.
func NewMetadataCatHeadObject() MetadataCatHeadObject {
	return getMetadataCatHeadObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetadataCatHeadObject */
// A concrete metadata object subclass representing a cat head.
//
// is a concrete subclass of representing a cat head.


// A concrete metadata object subclass representing a cat head.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataCatHeadObject
type MetadataCatHeadObject struct {
	MetadataObject
}

// MetadataCatHeadObjectFrom constructs a [MetadataCatHeadObject] from an unsafe.Pointer.
//
// A concrete metadata object subclass representing a cat head.
func MetadataCatHeadObjectFrom(ptr unsafe.Pointer) MetadataCatHeadObject {
	return MetadataCatHeadObject{
		MetadataObject: MetadataObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetadataCatHeadObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetadataCatHeadObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetadataCatHeadObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetadataCatHeadObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetadataCatHeadObject */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetadataCatHeadObject */



