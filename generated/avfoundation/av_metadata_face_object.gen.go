// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [MetadataFaceObject] class.
type IMetadataFaceObject interface {
	IMetadataObject
	

	// properties:
	FaceID() int
	HasRollAngle() bool
	HasYawAngle() bool
	RollAngle() float64
	YawAngle() float64


	

	// methods:


}





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

























// The unique ID for this face metadata object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/faceID
func (m_ MetadataFaceObject) FaceID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("faceID"))
	return rv
}


// A Boolean value indicating whether there is a valid roll angle associated with the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/hasRollAngle
func (m_ MetadataFaceObject) HasRollAngle() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasRollAngle"))
	return rv
}


// A Boolean value indicating whether there is a valid yaw angle associated with the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/hasYawAngle
func (m_ MetadataFaceObject) HasYawAngle() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasYawAngle"))
	return rv
}


// The roll angle of the face specified in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/rollAngle
func (m_ MetadataFaceObject) RollAngle() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("rollAngle"))
	return rv
}


// The yaw angle of the face specified in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataFaceObject/yawAngle
func (m_ MetadataFaceObject) YawAngle() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("yawAngle"))
	return rv
}








