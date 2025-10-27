// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [MetadataHumanFullBodyObject] class.
var (
	MetadataHumanFullBodyObjectClass     _MetadataHumanFullBodyObjectClass
	MetadataHumanFullBodyObjectClassOnce sync.Once
)

func getMetadataHumanFullBodyObjectClass() _MetadataHumanFullBodyObjectClass {
	MetadataHumanFullBodyObjectClassOnce.Do(func() {
		MetadataHumanFullBodyObjectClass = _MetadataHumanFullBodyObjectClass{objc.GetClass("AVMetadataHumanFullBodyObject")}
	})
	return MetadataHumanFullBodyObjectClass
}

type _MetadataHumanFullBodyObjectClass struct {
	class objc.Class
}





// An interface definition for the [MetadataHumanFullBodyObject] class.
type IMetadataHumanFullBodyObject interface {
	IMetadataBodyObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetadataHumanFullBodyObjectClass) Alloc() MetadataHumanFullBodyObject {
	rv := objc.Send[MetadataHumanFullBodyObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataHumanFullBodyObjectClass) New() MetadataHumanFullBodyObject {
	rv := objc.Send[MetadataHumanFullBodyObject](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataHumanFullBodyObject) Init() MetadataHumanFullBodyObject {
	rv := objc.Send[MetadataHumanFullBodyObject](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataHumanFullBodyObject) Autorelease() MetadataHumanFullBodyObject {
	rv := objc.Send[MetadataHumanFullBodyObject](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataHumanFullBodyObject creates a new MetadataHumanFullBodyObject instance.
func NewMetadataHumanFullBodyObject() MetadataHumanFullBodyObject {
	return getMetadataHumanFullBodyObjectClass().New()
}





// An object that represents a detected human full body in a picture.
//
// On supported platforms, outputs arrays of detected human full body objects.


// An object that represents a detected human full body in a picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataHumanFullBodyObject
type MetadataHumanFullBodyObject struct {
	MetadataBodyObject
}

// MetadataHumanFullBodyObjectFrom constructs a [MetadataHumanFullBodyObject] from an unsafe.Pointer.
//
// An object that represents a detected human full body in a picture.
func MetadataHumanFullBodyObjectFrom(ptr unsafe.Pointer) MetadataHumanFullBodyObject {
	return MetadataHumanFullBodyObject{
		MetadataBodyObject: MetadataBodyObjectFrom(ptr),
	}
}































