// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MetadataDogHeadObject] class.
type IMetadataDogHeadObject interface {
	IMetadataObject
}

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

// Alloc allocates a new instance without initialization.
func (mc _MetadataDogHeadObjectClass) Alloc() MetadataDogHeadObject {
	rv := objc.Send[MetadataDogHeadObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




