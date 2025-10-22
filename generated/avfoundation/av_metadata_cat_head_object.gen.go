// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MetadataCatHeadObject] class.
type IMetadataCatHeadObject interface {
	IMetadataObject
}

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

// Alloc allocates a new instance without initialization.
func (mc _MetadataCatHeadObjectClass) Alloc() MetadataCatHeadObject {
	rv := objc.Send[MetadataCatHeadObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




