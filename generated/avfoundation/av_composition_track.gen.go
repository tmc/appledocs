// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CompositionTrack] class.
var (
	CompositionTrackClass     _CompositionTrackClass
	CompositionTrackClassOnce sync.Once
)

func getCompositionTrackClass() _CompositionTrackClass {
	CompositionTrackClassOnce.Do(func() {
		CompositionTrackClass = _CompositionTrackClass{objc.GetClass("AVCompositionTrack")}
	})
	return CompositionTrackClass
}

type _CompositionTrackClass struct {
	class objc.Class
}

// An interface definition for the [CompositionTrack] class.
type ICompositionTrack interface {
	IAssetTrack
	MetadataForFormat(format unsafe.Pointer) []MetadataItem
}

// A track in a composition that presents media of a uniform type.
//
// This object provides an immutable composition track. The framework also provides a mutable subclass, .
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack
type CompositionTrack struct {
	AssetTrack
}

// CompositionTrackFrom constructs a [CompositionTrack] from an unsafe.Pointer.
//
// A track in a composition that presents media of a uniform type.
func CompositionTrackFrom(ptr unsafe.Pointer) CompositionTrack {
	return CompositionTrack{
		AssetTrack: AssetTrackFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CompositionTrackClass) Alloc() CompositionTrack {
	rv := objc.Send[CompositionTrack](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CompositionTrackClass) New() CompositionTrack {
	rv := objc.Send[CompositionTrack](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CompositionTrack) Init() CompositionTrack {
	rv := objc.Send[CompositionTrack](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CompositionTrack) Autorelease() CompositionTrack {
	rv := objc.Send[CompositionTrack](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCompositionTrack creates a new CompositionTrack instance.
func NewCompositionTrack() CompositionTrack {
	return getCompositionTrackClass().New()
}


// Returns metadata items that a track contains for the specified format.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata(forFormat:)
func (c_ CompositionTrack) MetadataForFormat(format unsafe.Pointer) []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("metadataForFormat:"), format)
	return rv
}



