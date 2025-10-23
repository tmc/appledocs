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
	FaceID() int
	SetFaceID(value int)
	HasRollAngle() bool
	SetHasRollAngle(value bool)
	HasYawAngle() bool
	SetHasYawAngle(value bool)
	RollAngle() float64
	SetRollAngle(value float64)
	YawAngle() float64
	SetYawAngle(value float64)
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

// Alloc allocates a new instance without initialization.
func (mc _MetadataFaceObjectClass) Alloc() MetadataFaceObject {
	rv := objc.Send[MetadataFaceObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The unique ID for this face metadata object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/faceid
func (m_ MetadataFaceObject) FaceID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("faceID"))
	return rv
}


// The unique ID for this face metadata object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/faceid
func (m_ MetadataFaceObject) SetFaceID(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFaceID:"), value)
}


// A Boolean value indicating whether there is a valid roll angle associated with the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/hasrollangle
func (m_ MetadataFaceObject) HasRollAngle() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasRollAngle"))
	return rv
}


// A Boolean value indicating whether there is a valid roll angle associated with the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/hasrollangle
func (m_ MetadataFaceObject) SetHasRollAngle(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasRollAngle:"), value)
}


// A Boolean value indicating whether there is a valid yaw angle associated with the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/hasyawangle
func (m_ MetadataFaceObject) HasYawAngle() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasYawAngle"))
	return rv
}


// A Boolean value indicating whether there is a valid yaw angle associated with the face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/hasyawangle
func (m_ MetadataFaceObject) SetHasYawAngle(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasYawAngle:"), value)
}


// The roll angle of the face specified in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/rollangle
func (m_ MetadataFaceObject) RollAngle() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("rollAngle"))
	return rv
}


// The roll angle of the face specified in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/rollangle
func (m_ MetadataFaceObject) SetRollAngle(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRollAngle:"), value)
}


// The yaw angle of the face specified in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/yawangle
func (m_ MetadataFaceObject) YawAngle() float64 {
	rv := objc.Send[float64](m_.ID, objc.Sel("yawAngle"))
	return rv
}


// The yaw angle of the face specified in degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatafaceobject/yawangle
func (m_ MetadataFaceObject) SetYawAngle(value float64) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setYawAngle:"), value)
}



