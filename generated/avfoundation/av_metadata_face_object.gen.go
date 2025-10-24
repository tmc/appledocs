// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class AVMetadataFaceObject */


/* debug [class_header]: Header for AVMetadataFaceObject */
// The class instance for the [MetadataFaceObject] class.
var (
	MetadataFaceObjectClass     _MetadataFaceObjectClass
	MetadataFaceObjectClassOnce sync.Once
)

func getMetadataFaceObjectClass() _MetadataFaceObjectClass {
	MetadataFaceObjectClassOnce.Do(func() {
		MetadataFaceObjectClass = _MetadataFaceObjectClass{objc.GetClass("AVMetadataFaceObject")}
	})
	return MetadataFaceObjectClass
}

type _MetadataFaceObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MetadataFaceObject */
// An interface definition for the [MetadataFaceObject] class.
type IMetadataFaceObject interface {
	IMetadataObject
	
/* debug [class_interface_properties]: Properties for MetadataFaceObject */
	// properties:
	FaceID() int
	HasRollAngle() bool
	HasYawAngle() bool
	RollAngle() float64
	YawAngle() float64
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MetadataFaceObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MetadataFaceObject */
// Alloc allocates a new instance without initialization.
func (mc _MetadataFaceObjectClass) Alloc() MetadataFaceObject {
	rv := objc.Send[MetadataFaceObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataFaceObjectClass) New() MetadataFaceObject {
	rv := objc.Send[MetadataFaceObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataFaceObject) Init() MetadataFaceObject {
	rv := objc.Send[MetadataFaceObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataFaceObject) Autorelease() MetadataFaceObject {
	rv := objc.Send[MetadataFaceObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataFaceObject creates a new MetadataFaceObject instance.
func NewMetadataFaceObject() MetadataFaceObject {
	return getMetadataFaceObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MetadataFaceObject */
// Face information detected by a metadata capture output.
//
// The class is a concrete subclass of that defines the features of a single detected face. You can retrieve instances of this class from the output of an object on devices that support face detection.


// Face information detected by a metadata capture output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject
type MetadataFaceObject struct {
	MetadataObject
}

// MetadataFaceObjectFrom constructs a [MetadataFaceObject] from an unsafe.Pointer.
//
// Face information detected by a metadata capture output.
func MetadataFaceObjectFrom(ptr unsafe.Pointer) MetadataFaceObject {
	return MetadataFaceObject{
		MetadataObject: MetadataObjectFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MetadataFaceObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MetadataFaceObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MetadataFaceObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MetadataFaceObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MetadataFaceObject */

// The unique ID for this face metadata object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/faceID
func (m_ MetadataFaceObject) FaceID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("faceID"))
	return rv
}/* debug [instance_properties/getter]: faceID */


// A Boolean value indicating whether there is a valid roll angle associated with the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/hasRollAngle
func (m_ MetadataFaceObject) HasRollAngle() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasRollAngle"))
	return rv
}/* debug [instance_properties/getter]: hasRollAngle */


// A Boolean value indicating whether there is a valid yaw angle associated with the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/hasYawAngle
func (m_ MetadataFaceObject) HasYawAngle() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasYawAngle"))
	return rv
}/* debug [instance_properties/getter]: hasYawAngle */


// The roll angle of the face specified in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/rollAngle
func (m_ MetadataFaceObject) RollAngle() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("rollAngle"))
	return rv
}/* debug [instance_properties/getter]: rollAngle */


// The yaw angle of the face specified in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/yawAngle
func (m_ MetadataFaceObject) YawAngle() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("yawAngle"))
	return rv
}/* debug [instance_properties/getter]: yawAngle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMetadataFaceObject */



