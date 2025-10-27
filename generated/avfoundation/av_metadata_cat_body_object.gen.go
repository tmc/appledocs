// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MetadataCatBodyObject] class.
var (
	MetadataCatBodyObjectClass     _MetadataCatBodyObjectClass
	MetadataCatBodyObjectClassOnce sync.Once
)

func getMetadataCatBodyObjectClass() _MetadataCatBodyObjectClass {
	MetadataCatBodyObjectClassOnce.Do(func() {
		MetadataCatBodyObjectClass = _MetadataCatBodyObjectClass{objc.GetClass("AVMetadataCatBodyObject")}
	})
	return MetadataCatBodyObjectClass
}

type _MetadataCatBodyObjectClass struct {
	class objc.Class
}





// An interface definition for the [MetadataCatBodyObject] class.
type IMetadataCatBodyObject interface {
	IMetadataBodyObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetadataCatBodyObjectClass) Alloc() MetadataCatBodyObject {
	rv := objc.Send[MetadataCatBodyObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataCatBodyObjectClass) New() MetadataCatBodyObject {
	rv := objc.Send[MetadataCatBodyObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataCatBodyObject) Init() MetadataCatBodyObject {
	rv := objc.Send[MetadataCatBodyObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataCatBodyObject) Autorelease() MetadataCatBodyObject {
	rv := objc.Send[MetadataCatBodyObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataCatBodyObject creates a new MetadataCatBodyObject instance.
func NewMetadataCatBodyObject() MetadataCatBodyObject {
	return getMetadataCatBodyObjectClass().New()
}





// An object representing a single detected cat body in a picture.
//
// This object is an immutable type that describes the various features found in the cat body in a picture.


// An object representing a single detected cat body in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataCatBodyObject
type MetadataCatBodyObject struct {
	MetadataBodyObject
}

// MetadataCatBodyObjectFrom constructs a [MetadataCatBodyObject] from an unsafe.Pointer.
//
// An object representing a single detected cat body in a picture.
func MetadataCatBodyObjectFrom(ptr unsafe.Pointer) MetadataCatBodyObject {
	return MetadataCatBodyObject{
		MetadataBodyObject: MetadataBodyObjectFrom(ptr),
	}
}































