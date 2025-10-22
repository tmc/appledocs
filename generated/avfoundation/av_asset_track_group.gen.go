// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetTrackGroup] class.
var (
	AssetTrackGroupClass     _AssetTrackGroupClass
	AssetTrackGroupClassOnce sync.Once
)

func getAssetTrackGroupClass() _AssetTrackGroupClass {
	AssetTrackGroupClassOnce.Do(func() {
		AssetTrackGroupClass = _AssetTrackGroupClass{objc.GetClass("AVAssetTrackGroup")}
	})
	return AssetTrackGroupClass
}

type _AssetTrackGroupClass struct {
	class objc.Class
}

// An interface definition for the [AssetTrackGroup] class.
type IAssetTrackGroup interface {
	objectivec.IObject
	TrackIDs() foundation.Number
	SetTrackIDs(value foundation.INumber)
}

// A group of related tracks in an asset.
//
// A track group describes a group of related alternative tracks, only one of which should play at a time. Groups of alternative tracks typically contain variations of the same content, like subtitles in multiple translations. You can inspect an asset’s track groups by loading the value of its property.


// A group of related tracks in an asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroup

type AssetTrackGroup struct {
	objectivec.Object
}

// AssetTrackGroupFrom constructs a [AssetTrackGroup] from an unsafe.Pointer.
//
// A group of related tracks in an asset.
func AssetTrackGroupFrom(ptr unsafe.Pointer) AssetTrackGroup {
	return AssetTrackGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetTrackGroupClass) Alloc() AssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetTrackGroupClass) New() AssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetTrackGroup) Init() AssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetTrackGroup) Autorelease() AssetTrackGroup {
	rv := objc.Send[AssetTrackGroup](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetTrackGroup creates a new AssetTrackGroup instance.
func NewAssetTrackGroup() AssetTrackGroup {
	return getAssetTrackGroupClass().New()
}



// The IDs of the tracks in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrackgroup/trackids

func (a_ AssetTrackGroup) TrackIDs() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("trackIDs"))
	return rv
}


// The IDs of the tracks in the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassettrackgroup/trackids

func (a_ AssetTrackGroup) SetTrackIDs(value foundation.INumber) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTrackIDs:"), value)
}



