// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [MetadataDogBodyObject] class.
type IMetadataDogBodyObject interface {
	IMetadataBodyObject
	

	// properties:


	

	// methods:


}





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































