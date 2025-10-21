// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MediaSelectionOption] class.
var (
	MediaSelectionOptionClass     _MediaSelectionOptionClass
	MediaSelectionOptionClassOnce sync.Once
)

func getMediaSelectionOptionClass() _MediaSelectionOptionClass {
	MediaSelectionOptionClassOnce.Do(func() {
		MediaSelectionOptionClass = _MediaSelectionOptionClass{objc.GetClass("AVMediaSelectionOption")}
	})
	return MediaSelectionOptionClass
}

type _MediaSelectionOptionClass struct {
	class objc.Class
}

// An interface definition for the [MediaSelectionOption] class.
type IMediaSelectionOption interface {
	objectivec.IObject
}

// An object that represents a specific option for the presentation of media within a group of options.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption
type MediaSelectionOption struct {
	objectivec.Object
}

// MediaSelectionOptionFrom constructs a [MediaSelectionOption] from an unsafe.Pointer.
//
// An object that represents a specific option for the presentation of media within a group of options.
func MediaSelectionOptionFrom(ptr unsafe.Pointer) MediaSelectionOption {
	return MediaSelectionOption{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaSelectionOptionClass) Alloc() MediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaSelectionOptionClass) New() MediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaSelectionOption) Init() MediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaSelectionOption) Autorelease() MediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaSelectionOption creates a new MediaSelectionOption instance.
func NewMediaSelectionOption() MediaSelectionOption {
	return getMediaSelectionOptionClass().New()
}




