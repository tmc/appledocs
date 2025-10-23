// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MediaSelection] class.
var (
	MediaSelectionClass     _MediaSelectionClass
	MediaSelectionClassOnce sync.Once
)

func getMediaSelectionClass() _MediaSelectionClass {
	MediaSelectionClassOnce.Do(func() {
		MediaSelectionClass = _MediaSelectionClass{objc.GetClass("AVMediaSelection")}
	})
	return MediaSelectionClass
}

type _MediaSelectionClass struct {
	class objc.Class
}

// An interface definition for the [MediaSelection] class.
type IMediaSelection interface {
	objectivec.IObject
	// properties:
	Asset() IAVAsset
	// methods:
	MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) bool /* primitive/slice/pointer. */
	SelectedMediaOptionInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) IMediaSelectionOption
}

// An object that represents a complete rendition of media selection options on an asset.


// An object that represents a complete rendition of media selection options on an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection
type MediaSelection struct {
	objectivec.Object
}

// MediaSelectionFrom constructs a [MediaSelection] from an unsafe.Pointer.
//
// An object that represents a complete rendition of media selection options on an asset.
func MediaSelectionFrom(ptr unsafe.Pointer) MediaSelection {
	return MediaSelection{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MediaSelectionClass) Alloc() MediaSelection {
	rv := objc.Send[MediaSelection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MediaSelectionClass) New() MediaSelection {
	rv := objc.Send[MediaSelection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaSelection) Init() MediaSelection {
	rv := objc.Send[MediaSelection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaSelection) Autorelease() MediaSelection {
	rv := objc.Send[MediaSelection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaSelection creates a new MediaSelection instance.
func NewMediaSelection() MediaSelection {
	return getMediaSelectionClass().New()
}



// Indicates whether the specified media selection group is subject to automatic media selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection/mediaSelectionCriteriaCanBeAppliedAutomatically(to:)
func (m_ MediaSelection) MediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("mediaSelectionCriteriaCanBeAppliedAutomaticallyToMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}


// Returns the media selection option that’s currently selected in the specified group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection/selectedMediaOption(in:)
func (m_ MediaSelection) SelectedMediaOptionInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) IMediaSelectionOption {
	rv := objc.Send[MediaSelectionOption](m_.ID, objc.Sel("selectedMediaOptionInMediaSelectionGroup:"), mediaSelectionGroup)
	return rv
}


// The asset associated with the media selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaSelection/asset
func (m_ MediaSelection) Asset() IAVAsset {
	rv := objc.Send[Asset](m_.ID, objc.Sel("asset"))
	return rv
}



