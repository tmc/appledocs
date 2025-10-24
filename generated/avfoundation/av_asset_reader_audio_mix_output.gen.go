// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetReaderAudioMixOutput */


/* debug [class_header]: Header for AVAssetReaderAudioMixOutput */
// The class instance for the [AssetReaderAudioMixOutput] class.
var (
	AssetReaderAudioMixOutputClass     _AssetReaderAudioMixOutputClass
	AssetReaderAudioMixOutputClassOnce sync.Once
)

func getAssetReaderAudioMixOutputClass() _AssetReaderAudioMixOutputClass {
	AssetReaderAudioMixOutputClassOnce.Do(func() {
		AssetReaderAudioMixOutputClass = _AssetReaderAudioMixOutputClass{objc.GetClass("AVAssetReaderAudioMixOutput")}
	})
	return AssetReaderAudioMixOutputClass
}

type _AssetReaderAudioMixOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetReaderAudioMixOutput */
// An interface definition for the [AssetReaderAudioMixOutput] class.
type IAssetReaderAudioMixOutput interface {
	IAssetReaderOutput
	
/* debug [class_interface_properties]: Properties for AssetReaderAudioMixOutput */
	// properties:
	AudioMix() IAVAudioMix
	SetAudioMix(value IAVAudioMix)
	AudioSettings() foundation.IDictionary
	AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */
	SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* typedef */)
	AudioTracks() []AssetTrack
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetReaderAudioMixOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetReaderAudioMixOutput */
// Alloc allocates a new instance without initialization.
func (ac _AssetReaderAudioMixOutputClass) Alloc() AssetReaderAudioMixOutput {
	rv := objc.Send[AssetReaderAudioMixOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetReaderAudioMixOutputClass) New() AssetReaderAudioMixOutput {
	rv := objc.Send[AssetReaderAudioMixOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReaderAudioMixOutput) Init() AssetReaderAudioMixOutput {
	rv := objc.Send[AssetReaderAudioMixOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReaderAudioMixOutput) Autorelease() AssetReaderAudioMixOutput {
	rv := objc.Send[AssetReaderAudioMixOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReaderAudioMixOutput creates a new AssetReaderAudioMixOutput instance.
func NewAssetReaderAudioMixOutput() AssetReaderAudioMixOutput {
	return getAssetReaderAudioMixOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetReaderAudioMixOutput */
// An object that reads audio samples that result from mixing audio from one or more tracks.
//
// Read audio data that you mix from one or more asset tracks by adding an audio mix output to an asset reader. You can read the samples in their stored format or you can convert them to an alternative format.


// An object that reads audio samples that result from mixing audio from one or more tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput
type AssetReaderAudioMixOutput struct {
	AssetReaderOutput
}

// AssetReaderAudioMixOutputFrom constructs a [AssetReaderAudioMixOutput] from an unsafe.Pointer.
//
// An object that reads audio samples that result from mixing audio from one or more tracks.
func AssetReaderAudioMixOutputFrom(ptr unsafe.Pointer) AssetReaderAudioMixOutput {
	return AssetReaderAudioMixOutput{
		AssetReaderOutput: AssetReaderOutputFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetReaderAudioMixOutput */

// Creates an object that reads mixed audio from the specified audio tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/init(audioTracks:audioSettings:)
func NewAssetReaderAudioMixOutputWithAudioTracksAudioSettings(audioTracks []AssetTrack, audioSettings foundation.IDictionary) AssetReaderAudioMixOutput {
	instance := getAssetReaderAudioMixOutputClass().Alloc()
	rv := objc.Send[AssetReaderAudioMixOutput](instance.ID, objc.Sel("initWithAudioTracks:audioSettings:"), audioTracks, audioSettings)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetReaderAudioMixOutputWithAudioTracksAudioSettings */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetReaderAudioMixOutput */

// Creates an object that reads mixed audio from the specified audio tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/assetReaderAudioMixOutputWithAudioTracks:audioSettings:
func (ac _AssetReaderAudioMixOutputClass) AssetReaderAudioMixOutputWithAudioTracksAudioSettings(audioTracks []AssetTrack, audioSettings foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetReaderAudioMixOutputWithAudioTracks:audioSettings:"), audioTracks, audioSettings)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetReaderAudioMixOutputWithAudioTracksAudioSettings) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetReaderAudioMixOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetReaderAudioMixOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetReaderAudioMixOutput */

// The audio mix to use with this output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/audioMix
func (a_ AssetReaderAudioMixOutput) AudioMix() IAVAudioMix {
	rv := objc.Send[AudioMix](a_.ID, objc.Sel("audioMix"))
	return rv
}/* debug [instance_properties/getter]: audioMix */


// The audio mix to use with this output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/audioMix
func (a_ AssetReaderAudioMixOutput) SetAudioMix(value IAVAudioMix) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioMix:"), value)
}/* debug [instance_properties/setter]: audioMix */


// The audio settings that the output uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/audioSettings
func (a_ AssetReaderAudioMixOutput) AudioSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("audioSettings"))
	return rv
}/* debug [instance_properties/getter]: audioSettings */


// The processing algorithm to use for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/audioTimePitchAlgorithm
func (a_ AssetReaderAudioMixOutput) AudioTimePitchAlgorithm() AudioTimePitchAlgorithm /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("audioTimePitchAlgorithm"))
	return rv
}/* debug [instance_properties/getter]: audioTimePitchAlgorithm */


// The processing algorithm to use for scaled audio edits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/audioTimePitchAlgorithm
func (a_ AssetReaderAudioMixOutput) SetAudioTimePitchAlgorithm(value AudioTimePitchAlgorithm /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioTimePitchAlgorithm:"), value)
}/* debug [instance_properties/setter]: audioTimePitchAlgorithm */


// The tracks from which the output reads audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/audioTracks
func (a_ AssetReaderAudioMixOutput) AudioTracks() []AssetTrack {
	rv := objc.Send[[]AssetTrack](a_.ID, objc.Sel("audioTracks"))
	return rv
}/* debug [instance_properties/getter]: audioTracks */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetReaderAudioMixOutput */


