// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetadataDogHeadObject */


/* debug [class_header]: Header for AVMetadataDogHeadObject */
// The class instance for the [MetadataDogHeadObject] class.
var (
	MetadataDogHeadObjectClass     _MetadataDogHeadObjectClass
	MetadataDogHeadObjectClassOnce sync.Once
)

func getMetadataDogHeadObjectClass() _MetadataDogHeadObjectClass {
	MetadataDogHeadObjectClassOnce.Do(func() {
		MetadataDogHeadObjectClass = _MetadataDogHeadObjectClass{objc.GetClass("AVMetadataDogHeadObject")}
	})
	return MetadataDogHeadObjectClass
}

type _MetadataDogHeadObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetadataDogHeadObject */
// An interface definition for the [MetadataDogHeadObject] class.
type IMetadataDogHeadObject interface {
	IMetadataObject
	
/* debug [class_interface_properties]: Properties for MetadataDogHeadObject */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetadataDogHeadObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetadataDogHeadObject */
// Alloc allocates a new instance without initialization.
func (mc _MetadataDogHeadObjectClass) Alloc() MetadataDogHeadObject {
	rv := objc.Send[MetadataDogHeadObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataDogHeadObjectClass) New() MetadataDogHeadObject {
	rv := objc.Send[MetadataDogHeadObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataDogHeadObject) Init() MetadataDogHeadObject {
	rv := objc.Send[MetadataDogHeadObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataDogHeadObject) Autorelease() MetadataDogHeadObject {
	rv := objc.Send[MetadataDogHeadObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataDogHeadObject creates a new MetadataDogHeadObject instance.
func NewMetadataDogHeadObject() MetadataDogHeadObject {
	return getMetadataDogHeadObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetadataDogHeadObject */
// A concrete metadata object subclass representing a dog head.
//
// is a concrete subclass of representing a dog head.


// A concrete metadata object subclass representing a dog head.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataDogHeadObject
type MetadataDogHeadObject struct {
	MetadataObject
}

// MetadataDogHeadObjectFrom constructs a [MetadataDogHeadObject] from an unsafe.Pointer.
//
// A concrete metadata object subclass representing a dog head.
func MetadataDogHeadObjectFrom(ptr unsafe.Pointer) MetadataDogHeadObject {
	return MetadataDogHeadObject{
		MetadataObject: MetadataObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetadataDogHeadObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetadataDogHeadObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetadataDogHeadObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetadataDogHeadObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetadataDogHeadObject */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetadataDogHeadObject */



