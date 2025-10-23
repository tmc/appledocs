// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MetadataSalientObject] class.
type IMetadataSalientObject interface {
	IMetadataObject
	ObjectID() int
	SetObjectID(value int)
}

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

// Alloc allocates a new instance without initialization.
func (mc _MetadataSalientObjectClass) Alloc() MetadataSalientObject {
	rv := objc.Send[MetadataSalientObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An integer value that defines the unique identifier of an object in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatasalientobject/objectid
func (m_ MetadataSalientObject) ObjectID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("objectID"))
	return rv
}


// An integer value that defines the unique identifier of an object in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatasalientobject/objectid
func (m_ MetadataSalientObject) SetObjectID(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setObjectID:"), value)
}



