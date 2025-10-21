// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MediaSelectionGroup] class.
var (
	MediaSelectionGroupClass     _MediaSelectionGroupClass
	MediaSelectionGroupClassOnce sync.Once
)

func getMediaSelectionGroupClass() _MediaSelectionGroupClass {
	MediaSelectionGroupClassOnce.Do(func() {
		MediaSelectionGroupClass = _MediaSelectionGroupClass{objc.GetClass("AVMediaSelectionGroup")}
	})
	return MediaSelectionGroupClass
}

type _MediaSelectionGroupClass struct {
	class objc.Class
}

// An interface definition for the [MediaSelectionGroup] class.
type IMediaSelectionGroup interface {
	objectivec.IObject
}

// An object that represents a collection of mutually exclusive options for the presentation of media within an asset.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup
type MediaSelectionGroup struct {
	objectivec.Object
}

// MediaSelectionGroupFrom constructs a [MediaSelectionGroup] from an unsafe.Pointer.
//
// An object that represents a collection of mutually exclusive options for the presentation of media within an asset.
func MediaSelectionGroupFrom(ptr unsafe.Pointer) MediaSelectionGroup {
	return MediaSelectionGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaSelectionGroupClass) Alloc() MediaSelectionGroup {
	rv := objc.Send[MediaSelectionGroup](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaSelectionGroupClass) New() MediaSelectionGroup {
	rv := objc.Send[MediaSelectionGroup](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaSelectionGroup) Init() MediaSelectionGroup {
	rv := objc.Send[MediaSelectionGroup](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaSelectionGroup) Autorelease() MediaSelectionGroup {
	rv := objc.Send[MediaSelectionGroup](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaSelectionGroup creates a new MediaSelectionGroup instance.
func NewMediaSelectionGroup() MediaSelectionGroup {
	return getMediaSelectionGroupClass().New()
}


// A Boolean value that indicates whether it’s possible to present none of the options in the group when an associated player item is played.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/allowsEmptySelection
func (m_ MediaSelectionGroup) AllowsEmptySelection() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowsEmptySelection"))
	return rv
}



