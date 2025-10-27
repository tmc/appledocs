// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	NaturalSize() corefoundation.CGSize
	SetNaturalSize(value corefoundation.CGSize)
	Tracks() []MutableCompositionTrack


	

	// methods:
	AddTracksForCinematicAssetInfoPreferredStartingTrackID(assetInfo objectivec.IObject, preferredStartingTrackID PersistentTrackID /* not a class type */) objectivec.IObject
	AddMutableTrackWithMediaTypePreferredTrackID(mediaType MediaType, preferredTrackID PersistentTrackID /* not a class type */) IMutableCompositionTrack
	InsertEmptyTimeRange(timeRange objectivec.IObject)
	LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer)
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic, completionHandler unsafe.Pointer)
	LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType, completionHandler unsafe.Pointer)
	MutableTrackCompatibleWithTrack(track IAVAssetTrack) IMutableCompositionTrack
	RemoveTimeRange(timeRange objectivec.IObject)
	RemoveTrack(track IAVCompositionTrack)
	ScaleTimeRangeToDuration(timeRange objectivec.IObject, duration objectivec.IObject)
	TrackWithTrackID(trackID PersistentTrackID /* not a class type */) IMutableCompositionTrack
	TracksWithMediaCharacteristic(mediaCharacteristic MediaCharacteristic) []MutableCompositionTrack
	TracksWithMediaType(mediaType MediaType) []MutableCompositionTrack


}





// Alloc allocates a new instance without initialization.
func (mc _MutableCompositionClass) Alloc() MutableComposition {
	rv := objc.Send[MutableComposition](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates a mutable composition that uses the specified initialization options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/init(urlAssetInitializationOptions:)
func NewMutableCompositionWithURLAssetInitializationOptions(URLAssetInitializationOptions foundation.IDictionary) MutableComposition {
	rv := objc.Send[MutableComposition](objc.ID(getMutableCompositionClass().class), objc.Sel("compositionWithURLAssetInitializationOptions:"), URLAssetInitializationOptions)
	return rv
}







// Returns a new mutable composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/composition
func (mc _MutableCompositionClass) Composition() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("composition"))
	return rv
}


// Creates a mutable composition that uses the specified initialization options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/init(urlAssetInitializationOptions:)
func (mc _MutableCompositionClass) CompositionWithURLAssetInitializationOptions(URLAssetInitializationOptions foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("compositionWithURLAssetInitializationOptions:"), URLAssetInitializationOptions)
	return rv
}












// Adds a group of empty tracks associated with a cinematic asset to a mutable composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/addTracksForCinematicAssetInfo:preferredStartingTrackID:
func (m_ MutableComposition) AddTracksForCinematicAssetInfoPreferredStartingTrackID(assetInfo objectivec.IObject, preferredStartingTrackID PersistentTrackID /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("addTracksForCinematicAssetInfo:preferredStartingTrackID:"), assetInfo, preferredStartingTrackID)
	return rv
}


// Adds an empty track to a composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/addMutableTrack(withMediaType:preferredTrackID:)
func (m_ MutableComposition) AddMutableTrackWithMediaTypePreferredTrackID(mediaType MediaType, preferredTrackID PersistentTrackID /* not a class type */) IMutableCompositionTrack {
	rv := objc.Send[MutableCompositionTrack](m_.ID, objc.Sel("addMutableTrackWithMediaType:preferredTrackID:"), mediaType, preferredTrackID)
	return rv
}


// Adds or extends an empty time range within all tracks of the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/insertEmptyTimeRange(_:)
func (m_ MutableComposition) InsertEmptyTimeRange(timeRange objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertEmptyTimeRange:"), timeRange)
}


// Loads a track that contains the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/loadTrack(withTrackID:completionHandler:)
func (m_ MutableComposition) LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadTrackWithTrackID:completionHandler:"), trackID, completionHandler)
}


// Loads tracks that contain media of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/loadTracks(withMediaCharacteristic:completionHandler:)
func (m_ MutableComposition) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}


// Loads tracks that contain media of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/loadTracks(withMediaType:completionHandler:)
func (m_ MutableComposition) LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadTracksWithMediaType:completionHandler:"), mediaType, completionHandler)
}


// Returns a composition track into which you can insert any time range of the specified asset track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/mutableTrack(compatibleWith:)
func (m_ MutableComposition) MutableTrackCompatibleWithTrack(track IAVAssetTrack) IMutableCompositionTrack {
	rv := objc.Send[MutableCompositionTrack](m_.ID, objc.Sel("mutableTrackCompatibleWithTrack:"), track)
	return rv
}


// Removes a specified time range from all tracks of the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/removeTimeRange(_:)
func (m_ MutableComposition) RemoveTimeRange(timeRange objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTimeRange:"), timeRange)
}


// Removes a specified track from the composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/removeTrack(_:)
func (m_ MutableComposition) RemoveTrack(track IAVCompositionTrack) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeTrack:"), track)
}


// Changes the duration of all tracks in a given time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/scaleTimeRange(_:toDuration:)
func (m_ MutableComposition) ScaleTimeRangeToDuration(timeRange objectivec.IObject, duration objectivec.IObject) {
	objc.Send[objc.ID](m_.ID, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}


// Returns a track that contains the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/track(withTrackID:)
func (m_ MutableComposition) TrackWithTrackID(trackID PersistentTrackID /* not a class type */) IMutableCompositionTrack {
	rv := objc.Send[MutableCompositionTrack](m_.ID, objc.Sel("trackWithTrackID:"), trackID)
	return rv
}


// Returns tracks that contain media of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/tracks(withMediaCharacteristic:)
func (m_ MutableComposition) TracksWithMediaCharacteristic(mediaCharacteristic MediaCharacteristic) []MutableCompositionTrack {
	rv := objc.Send[[]MutableCompositionTrack](m_.ID, objc.Sel("tracksWithMediaCharacteristic:"), mediaCharacteristic)
	return rv
}


// Returns tracks that contain media of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/tracks(withMediaType:)
func (m_ MutableComposition) TracksWithMediaType(mediaType MediaType) []MutableCompositionTrack {
	rv := objc.Send[[]MutableCompositionTrack](m_.ID, objc.Sel("tracksWithMediaType:"), mediaType)
	return rv
}







// The encoded or authored size of the visual portion of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/naturalSize
func (m_ MutableComposition) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m_.ID, objc.Sel("naturalSize"))
	return rv
}


// The encoded or authored size of the visual portion of the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/naturalSize
func (m_ MutableComposition) SetNaturalSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNaturalSize:"), value)
}


// The tracks that a composition contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/tracks
func (m_ MutableComposition) Tracks() []MutableCompositionTrack {
	rv := objc.Send[[]MutableCompositionTrack](m_.ID, objc.Sel("tracks"))
	return rv
}







