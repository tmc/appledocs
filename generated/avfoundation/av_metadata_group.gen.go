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
	// methods:
}

// A parent class referenced by other AVFoundation classes.


// A parent class referenced by other AVFoundation classes. [Full Topic]
type MetadataGroup struct {
	objectivec.Object
}

// MetadataGroupFrom constructs a [MetadataGroup] from an unsafe.Pointer.
//
// A parent class referenced by other AVFoundation classes.
func MetadataGroupFrom(ptr unsafe.Pointer) MetadataGroup {
	return MetadataGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetadataGroupClass) Alloc() MetadataGroup {
	rv := objc.Send[MetadataGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




