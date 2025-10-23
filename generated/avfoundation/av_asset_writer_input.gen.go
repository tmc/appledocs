// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetWriterInput] class.
var (
	AssetWriterInputClass     _AssetWriterInputClass
	AssetWriterInputClassOnce sync.Once
)

func getAssetWriterInputClass() _AssetWriterInputClass {
	AssetWriterInputClassOnce.Do(func() {
		AssetWriterInputClass = _AssetWriterInputClass{objc.GetClass("AVAssetWriterInput")}
	})
	return AssetWriterInputClass
}

type _AssetWriterInputClass struct {
	class objc.Class
}

// An interface definition for the [AssetWriterInput] class.
type IAssetWriterInput interface {
	objectivec.IObject
	// properties:
	CanPerformMultiplePasses() bool /* primitive/slice/pointer. */
	SetCanPerformMultiplePasses(value bool /* primitive/slice/pointer. */)
	CurrentPassDescription() AssetWriterInputPassDescription /* not a class type */
	SetCurrentPassDescription(value AssetWriterInputPassDescription /* not a class type */)
	ExpectsMediaDataInRealTime() bool /* primitive/slice/pointer. */
	SetExpectsMediaDataInRealTime(value bool /* primitive/slice/pointer. */)
	ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */
	SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */)
	IsReadyForMoreMediaData() bool /* primitive/slice/pointer. */
	SetIsReadyForMoreMediaData(value bool /* primitive/slice/pointer. */)
	LanguageCode() objc.IObject /* cross-framework: NSString */
	SetLanguageCode(value objc.IObject /* cross-framework: NSString */)
	MarksOutputTrackAsEnabled() bool /* primitive/slice/pointer. */
	SetMarksOutputTrackAsEnabled(value bool /* primitive/slice/pointer. */)
	MediaDataLocation() unsafe.Pointer
	SetMediaDataLocation(value unsafe.Pointer)
	MediaTimeScale() TimeScale /* not a class type */
	SetMediaTimeScale(value TimeScale /* not a class type */)
	MediaType() MediaType /* not a class type */
	SetMediaType(value MediaType /* not a class type */)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	NaturalSize() objc.IObject /* cross-framework: Size */
	SetNaturalSize(value objc.IObject /* cross-framework: Size */)
	OutputSettings() objc.IObject /* cross-framework: NSString */
	SetOutputSettings(value objc.IObject /* cross-framework: NSString */)
	PerformsMultiPassEncodingIfSupported() bool /* primitive/slice/pointer. */
	SetPerformsMultiPassEncodingIfSupported(value bool /* primitive/slice/pointer. */)
	PreferredMediaChunkAlignment() int /* primitive/slice/pointer. */
	SetPreferredMediaChunkAlignment(value int /* primitive/slice/pointer. */)
	PreferredMediaChunkDuration() Time /* not a class type */
	SetPreferredMediaChunkDuration(value Time /* not a class type */)
	PreferredVolume() float32 /* primitive/slice/pointer. */
	SetPreferredVolume(value float32 /* primitive/slice/pointer. */)
	SampleReferenceBaseURL() objc.IObject /* cross-framework: URL */
	SetSampleReferenceBaseURL(value objc.IObject /* cross-framework: URL */)
	SourceFormatHint() FormatDescription /* not a class type */
	SetSourceFormatHint(value FormatDescription /* not a class type */)
	Transform() objc.IObject /* cross-framework: AffineTransform */
	SetTransform(value objc.IObject /* cross-framework: AffineTransform */)
	// methods:
}

// An object that appends media samples to a track in an asset writer’s output file.
//
// Create an asset writer input to write a single track of media, and optional track-level metadata, to the output file. To write multiple concurrent tracks with ideal interleaving of media data, observe the value of the property of each input. You can use an asset writer input to create tracks in a QuickTime movie file that aren’t self-contained, and instead reference sample data that exists in another file.


// An object that appends media samples to a track in an asset writer’s output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput
type AssetWriterInput struct {
	objectivec.Object
}

// AssetWriterInputFrom constructs a [AssetWriterInput] from an unsafe.Pointer.
//
// An object that appends media samples to a track in an asset writer’s output file.
func AssetWriterInputFrom(ptr unsafe.Pointer) AssetWriterInput {
	return AssetWriterInput{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputClass) Alloc() AssetWriterInput {
	rv := objc.Send[AssetWriterInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetWriterInputClass) New() AssetWriterInput {
	rv := objc.Send[AssetWriterInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetWriterInput) Init() AssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetWriterInput) Autorelease() AssetWriterInput {
	rv := objc.Send[AssetWriterInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetWriterInput creates a new AssetWriterInput instance.
func NewAssetWriterInput() AssetWriterInput {
	return getAssetWriterInputClass().New()
}



// A Boolean value that indicates whether the input may perform multiple passes over appended media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/canperformmultiplepasses
func (a_ AssetWriterInput) CanPerformMultiplePasses() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformMultiplePasses"))
	return rv
}


// A Boolean value that indicates whether the input may perform multiple passes over appended media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/canperformmultiplepasses
func (a_ AssetWriterInput) SetCanPerformMultiplePasses(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanPerformMultiplePasses:"), value)
}


// An object that describes the requirements for the current pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/currentpassdescription
func (a_ AssetWriterInput) CurrentPassDescription() AssetWriterInputPassDescription /* not a class type */ {
	rv := objc.Send[AssetWriterInputPassDescription](a_.ID, objc.Sel("currentPassDescription"))
	return rv
}


// An object that describes the requirements for the current pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/currentpassdescription
func (a_ AssetWriterInput) SetCurrentPassDescription(value AssetWriterInputPassDescription /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentPassDescription:"), value)
}


// A Boolean value that indicates whether the input tailors its processing for real-time sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/expectsmediadatainrealtime
func (a_ AssetWriterInput) ExpectsMediaDataInRealTime() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("expectsMediaDataInRealTime"))
	return rv
}


// A Boolean value that indicates whether the input tailors its processing for real-time sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/expectsmediadatainrealtime
func (a_ AssetWriterInput) SetExpectsMediaDataInRealTime(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExpectsMediaDataInRealTime:"), value)
}


// The extended language for the input’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/extendedlanguagetag
func (a_ AssetWriterInput) ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The extended language for the input’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/extendedlanguagetag
func (a_ AssetWriterInput) SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}


// A Boolean value that indicates whether the input is ready to accept media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/isreadyformoremediadata
func (a_ AssetWriterInput) IsReadyForMoreMediaData() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReadyForMoreMediaData"))
	return rv
}


// A Boolean value that indicates whether the input is ready to accept media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/isreadyformoremediadata
func (a_ AssetWriterInput) SetIsReadyForMoreMediaData(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsReadyForMoreMediaData:"), value)
}


// The language code of the input’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/languagecode
func (a_ AssetWriterInput) LanguageCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("languageCode"))
	return rv
}


// The language code of the input’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/languagecode
func (a_ AssetWriterInput) SetLanguageCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLanguageCode:"), value)
}


// A Boolean value that indicates whether to enable a track in the output for playback and processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/marksoutputtrackasenabled
func (a_ AssetWriterInput) MarksOutputTrackAsEnabled() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("marksOutputTrackAsEnabled"))
	return rv
}


// A Boolean value that indicates whether to enable a track in the output for playback and processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/marksoutputtrackasenabled
func (a_ AssetWriterInput) SetMarksOutputTrackAsEnabled(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMarksOutputTrackAsEnabled:"), value)
}


// Specifies how the input lays out and interleaves media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediadatalocation-swift.property
func (a_ AssetWriterInput) MediaDataLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("mediaDataLocation"))
	return rv
}


// Specifies how the input lays out and interleaves media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediadatalocation-swift.property
func (a_ AssetWriterInput) SetMediaDataLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaDataLocation:"), value)
}


// The time scale of the track in the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediatimescale
func (a_ AssetWriterInput) MediaTimeScale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](a_.ID, objc.Sel("mediaTimeScale"))
	return rv
}


// The time scale of the track in the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediatimescale
func (a_ AssetWriterInput) SetMediaTimeScale(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaTimeScale:"), value)
}


// The media type of the samples that the input accepts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediatype
func (a_ AssetWriterInput) MediaType() MediaType /* not a class type */ {
	rv := objc.Send[MediaType](a_.ID, objc.Sel("mediaType"))
	return rv
}


// The media type of the samples that the input accepts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediatype
func (a_ AssetWriterInput) SetMediaType(value MediaType /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaType:"), value)
}


// The track-level metadata to write to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/metadata
func (a_ AssetWriterInput) Metadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// The track-level metadata to write to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/metadata
func (a_ AssetWriterInput) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadata:"), value)
}


// The natural display dimensions of the output’s visual media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/naturalsize
func (a_ AssetWriterInput) NaturalSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[Size](a_.ID, objc.Sel("naturalSize"))
	return rv
}


// The natural display dimensions of the output’s visual media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/naturalsize
func (a_ AssetWriterInput) SetNaturalSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNaturalSize:"), value)
}


// The settings to use for encoding media data you append to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/outputsettings
func (a_ AssetWriterInput) OutputSettings() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("outputSettings"))
	return rv
}


// The settings to use for encoding media data you append to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/outputsettings
func (a_ AssetWriterInput) SetOutputSettings(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputSettings:"), value)
}


// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/performsmultipassencodingifsupported
func (a_ AssetWriterInput) PerformsMultiPassEncodingIfSupported() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("performsMultiPassEncodingIfSupported"))
	return rv
}


// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/performsmultipassencodingifsupported
func (a_ AssetWriterInput) SetPerformsMultiPassEncodingIfSupported(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPerformsMultiPassEncodingIfSupported:"), value)
}


// The boundary, in bytes, for aligning media chunks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredmediachunkalignment
func (a_ AssetWriterInput) PreferredMediaChunkAlignment() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](a_.ID, objc.Sel("preferredMediaChunkAlignment"))
	return rv
}


// The boundary, in bytes, for aligning media chunks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredmediachunkalignment
func (a_ AssetWriterInput) SetPreferredMediaChunkAlignment(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredMediaChunkAlignment:"), value)
}


// The duration to use for each chunk of sample data in the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredmediachunkduration
func (a_ AssetWriterInput) PreferredMediaChunkDuration() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("preferredMediaChunkDuration"))
	return rv
}


// The duration to use for each chunk of sample data in the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredmediachunkduration
func (a_ AssetWriterInput) SetPreferredMediaChunkDuration(value Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredMediaChunkDuration:"), value)
}


// The volume to prefer for playback of the output’s audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredvolume
func (a_ AssetWriterInput) PreferredVolume() float32 /* primitive/slice/pointer. */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The volume to prefer for playback of the output’s audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredvolume
func (a_ AssetWriterInput) SetPreferredVolume(value float32 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredVolume:"), value)
}


// The base URL sample references are relative to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/samplereferencebaseurl
func (a_ AssetWriterInput) SampleReferenceBaseURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("sampleReferenceBaseURL"))
	return rv
}


// The base URL sample references are relative to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/samplereferencebaseurl
func (a_ AssetWriterInput) SetSampleReferenceBaseURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleReferenceBaseURL:"), value)
}


// A hint about the format of the sample buffers to append to the input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/sourceformathint
func (a_ AssetWriterInput) SourceFormatHint() FormatDescription /* not a class type */ {
	rv := objc.Send[FormatDescription](a_.ID, objc.Sel("sourceFormatHint"))
	return rv
}


// A hint about the format of the sample buffers to append to the input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/sourceformathint
func (a_ AssetWriterInput) SetSourceFormatHint(value FormatDescription /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSourceFormatHint:"), value)
}


// The transform to use for display of the output’s visual media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/transform
func (a_ AssetWriterInput) Transform() objc.IObject /* cross-framework: AffineTransform */ {
	rv := objc.Send[AffineTransform](a_.ID, objc.Sel("transform"))
	return rv
}


// The transform to use for display of the output’s visual media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/transform
func (a_ AssetWriterInput) SetTransform(value objc.IObject /* cross-framework: AffineTransform */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransform:"), value)
}



