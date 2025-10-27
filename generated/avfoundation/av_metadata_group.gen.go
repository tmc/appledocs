// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MetadataGroup] class.
var (
	MetadataGroupClass     _MetadataGroupClass
	MetadataGroupClassOnce sync.Once
)

func getMetadataGroupClass() _MetadataGroupClass {
	MetadataGroupClassOnce.Do(func() {
		MetadataGroupClass = _MetadataGroupClass{objc.GetClass("AVMetadataGroup")}
	})
	return MetadataGroupClass
}

type _MetadataGroupClass struct {
	class objc.Class
}





// An interface definition for the [MetadataGroup] class.
type IMetadataGroup interface {
	objectivec.IObject
	

	// properties:
	ClassifyingLabel() foundation.foundation.INSString
	Items() []MetadataItem
	UniqueID() foundation.foundation.INSString


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MetadataGroupClass) Alloc() MetadataGroup {
	rv := objc.Send[MetadataGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MetadataGroupClass) New() MetadataGroup {
	rv := objc.Send[MetadataGroup](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataGroup) Init() MetadataGroup {
	rv := objc.Send[MetadataGroup](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataGroup) Autorelease() MetadataGroup {
	rv := objc.Send[MetadataGroup](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataGroup creates a new MetadataGroup instance.
func NewMetadataGroup() MetadataGroup {
	return getMetadataGroupClass().New()
}





// A collection of metadata items associated with a timeline segment.


// A collection of metadata items associated with a timeline segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup
type MetadataGroup struct {
	objectivec.Object
}

// MetadataGroupFrom constructs a [MetadataGroup] from an unsafe.Pointer.
//
// A collection of metadata items associated with a timeline segment.
func MetadataGroupFrom(ptr unsafe.Pointer) MetadataGroup {
	return MetadataGroup{objectivec.Object{objc.ID(ptr)}}
}

























// The classifying label associated with the metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup/classifyingLabel
func (m_ MetadataGroup) ClassifyingLabel() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("classifyingLabel"))
	return rv
}


// The array of metadata items associated with the metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup/items
func (m_ MetadataGroup) Items() []MetadataItem {
	rv := objc.Send[[]MetadataItem](m_.ID, objc.Sel("items"))
	return rv
}


// The unique identifier for the metadata group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataGroup/uniqueID
func (m_ MetadataGroup) UniqueID() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("uniqueID"))
	return rv
}








