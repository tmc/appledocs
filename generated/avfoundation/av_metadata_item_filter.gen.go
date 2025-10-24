// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MetadataItemFilter] class.
var (
	MetadataItemFilterClass     _MetadataItemFilterClass
	MetadataItemFilterClassOnce sync.Once
)

func getMetadataItemFilterClass() _MetadataItemFilterClass {
	MetadataItemFilterClassOnce.Do(func() {
		MetadataItemFilterClass = _MetadataItemFilterClass{objc.GetClass("AVMetadataItemFilter")}
	})
	return MetadataItemFilterClass
}

type _MetadataItemFilterClass struct {
	class objc.Class
}





// An interface definition for the [MetadataItemFilter] class.
type IMetadataItemFilter interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetadataItemFilterClass) Alloc() MetadataItemFilter {
	rv := objc.Send[MetadataItemFilter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataItemFilterClass) New() MetadataItemFilter {
	rv := objc.Send[MetadataItemFilter](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataItemFilter) Init() MetadataItemFilter {
	rv := objc.Send[MetadataItemFilter](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataItemFilter) Autorelease() MetadataItemFilter {
	rv := objc.Send[MetadataItemFilter](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataItemFilter creates a new MetadataItemFilter instance.
func NewMetadataItemFilter() MetadataItemFilter {
	return getMetadataItemFilterClass().New()
}





// An object that filters selected information from a metadata item.
//
// Filter instances are opaque, unmodifiable objects, that you create with the class method.


// An object that filters selected information from a metadata item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemFilter
type MetadataItemFilter struct {
	objectivec.Object
}

// MetadataItemFilterFrom constructs a [MetadataItemFilter] from an unsafe.Pointer.
//
// An object that filters selected information from a metadata item.
func MetadataItemFilterFrom(ptr unsafe.Pointer) MetadataItemFilter {
	return MetadataItemFilter{objectivec.Object{objc.ID(ptr)}}
}










// Returns a metadata filter to use for sharing assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItemFilter/forSharing()
func (mc _MetadataItemFilterClass) MetadataItemFilterForSharing() IMetadataItemFilter {
	rv := objc.Send[MetadataItemFilter](objc.ID(mc.class), objc.Sel("metadataItemFilterForSharing"))
	return rv
}























