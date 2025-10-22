// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [MutableComposition] class.
var (
	MutableCompositionClass     _MutableCompositionClass
	MutableCompositionClassOnce sync.Once
)

func getMutableCompositionClass() _MutableCompositionClass {
	MutableCompositionClassOnce.Do(func() {
		MutableCompositionClass = _MutableCompositionClass{objc.GetClass("AVMutableComposition")}
	})
	return MutableCompositionClass
}

type _MutableCompositionClass struct {
	class objc.Class
}

// An interface definition for the [MutableComposition] class.
type IMutableComposition interface {
	IComposition
	InsertTimeRangeOfAssetAtTimeError(timeRange unsafe.Pointer, asset IAVAsset, startTime unsafe.Pointer, outError unsafe.Pointer) bool
	NaturalSize() coregraphics.CGSize
	SetNaturalSize(value coregraphics.CGSize)
	Tracks() unsafe.Pointer
	SetTracks(value unsafe.Pointer)
}

// An object that you use to create a new composition from existing assets.
//
// Use this object to add and remove composition tracks, and add, remove, and scale their time ranges. You can make an immutable snapshot of a mutable composition for playback and inspection as follows:


// An object that you use to create a new composition from existing assets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition

type MutableComposition struct {
	Composition
}

// MutableCompositionFrom constructs a [MutableComposition] from an unsafe.Pointer.
//
// An object that you use to create a new composition from existing assets.
func MutableCompositionFrom(ptr unsafe.Pointer) MutableComposition {
	return MutableComposition{
		Composition: CompositionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableCompositionClass) Alloc() MutableComposition {
	rv := objc.Send[MutableComposition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableCompositionClass) New() MutableComposition {
	rv := objc.Send[MutableComposition](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableComposition) Init() MutableComposition {
	rv := objc.Send[MutableComposition](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableComposition) Autorelease() MutableComposition {
	rv := objc.Send[MutableComposition](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableComposition creates a new MutableComposition instance.
func NewMutableComposition() MutableComposition {
	return getMutableCompositionClass().New()
}




// Inserts all the tracks within a given time range of a specified asset into the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/insertTimeRange(_:of:at:)

func (m_ MutableComposition) InsertTimeRangeOfAssetAtTimeError(timeRange unsafe.Pointer, asset IAVAsset, startTime unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("insertTimeRange:ofAsset:atTime:error:"), timeRange, asset, startTime, outError)
	return rv
}


// The encoded or authored size of the visual portion of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/naturalsize

func (m_ MutableComposition) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](m_.ID, objc.Sel("naturalSize"))
	return rv
}


// The encoded or authored size of the visual portion of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/naturalsize

func (m_ MutableComposition) SetNaturalSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalSize:"), value)
}


// The tracks that a composition contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/tracks

func (m_ MutableComposition) Tracks() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("tracks"))
	return rv
}


// The tracks that a composition contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmutablecomposition/tracks

func (m_ MutableComposition) SetTracks(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTracks:"), value)
}



