// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MetadataBodyObject] class.
type IMetadataBodyObject interface {
	IMetadataObject
	BodyID() int
	SetBodyID(value int)
}

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

// Alloc allocates a new instance without initialization.
func (mc _MetadataBodyObjectClass) Alloc() MetadataBodyObject {
	rv := objc.Send[MetadataBodyObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// An integer value that defines the unique identifier of an object in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatabodyobject/bodyid

func (m_ MetadataBodyObject) BodyID() int {
	rv := objc.Send[int](m_.ID, objc.Sel("bodyID"))
	return rv
}


// An integer value that defines the unique identifier of an object in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmetadatabodyobject/bodyid

func (m_ MetadataBodyObject) SetBodyID(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setBodyID:"), value)
}



