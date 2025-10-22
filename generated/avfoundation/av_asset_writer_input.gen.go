// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	AppendSampleBuffer(sampleBuffer unsafe.Pointer) bool
	CanPerformMultiplePasses() bool
	SetCanPerformMultiplePasses(value bool)
	CurrentPassDescription() unsafe.Pointer
	SetCurrentPassDescription(value unsafe.Pointer)
	ExpectsMediaDataInRealTime() bool
	SetExpectsMediaDataInRealTime(value bool)
	ExtendedLanguageTag() string
	SetExtendedLanguageTag(value string)
	IsReadyForMoreMediaData() bool
	SetIsReadyForMoreMediaData(value bool)
	LanguageCode() string
	SetLanguageCode(value string)
	MarksOutputTrackAsEnabled() bool
	SetMarksOutputTrackAsEnabled(value bool)
	MediaDataLocation() unsafe.Pointer
	SetMediaDataLocation(value unsafe.Pointer)
	MediaTimeScale() unsafe.Pointer
	SetMediaTimeScale(value unsafe.Pointer)
	MediaType() MediaType
	SetMediaType(value MediaType)
	Metadata() AVMetadataItem
	SetMetadata(value IAVMetadataItem)
	NaturalSize() coregraphics.CGSize
	SetNaturalSize(value coregraphics.CGSize)
	OutputSettings() string
	SetOutputSettings(value string)
	PerformsMultiPassEncodingIfSupported() bool
	SetPerformsMultiPassEncodingIfSupported(value bool)
	PreferredMediaChunkAlignment() int
	SetPreferredMediaChunkAlignment(value int)
	PreferredMediaChunkDuration() unsafe.Pointer
	SetPreferredMediaChunkDuration(value unsafe.Pointer)
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	SampleReferenceBaseURL() foundation.URL
	SetSampleReferenceBaseURL(value foundation.IURL)
	SourceFormatHint() unsafe.Pointer
	SetSourceFormatHint(value unsafe.Pointer)
	Transform() coregraphics.CGAffineTransform
	SetTransform(value coregraphics.CGAffineTransform)
}

// An object that appends media samples to a track in an asset writer’s output file.
//
// Create an asset writer input to write a single track of media, and optional track-level metadata, to the output file. To write multiple concurrent tracks with ideal interleaving of media data, observe the value of the property of each input. You can use an asset writer input to create tracks in a QuickTime movie file that aren’t self-contained, and instead reference sample data that exists in another file.
//
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




// Creates an input to append sample buffers of the specified type to the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/init(mediaType:outputSettings:)
func NewAssetWriterInputWithMediaTypeOutputSettings(mediaType MediaType, outputSettings unsafe.Pointer) AssetWriterInput {
	instance := getAssetWriterInputClass().Alloc()
	rv := objc.Send[AssetWriterInput](instance.ID, objc.Sel("initWithMediaType:outputSettings:"), mediaType, outputSettings)
	rv.Autorelease()
	return rv
}



// Creates an input that appends sample buffers of the specified type and format hint to the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/init(mediaType:outputSettings:sourceFormatHint:)
func NewAssetWriterInputWithMediaTypeOutputSettingsSourceFormatHint(mediaType MediaType, outputSettings unsafe.Pointer, sourceFormatHint unsafe.Pointer) AssetWriterInput {
	instance := getAssetWriterInputClass().Alloc()
	rv := objc.Send[AssetWriterInput](instance.ID, objc.Sel("initWithMediaType:outputSettings:sourceFormatHint:"), mediaType, outputSettings, sourceFormatHint)
	rv.Autorelease()
	return rv
}


// Appends a sample buffer to an input to write to the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/append(_:)
func (a_ AssetWriterInput) AppendSampleBuffer(sampleBuffer unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("appendSampleBuffer:"), sampleBuffer)
	return rv
}

// A Boolean value that indicates whether the input may perform multiple passes over appended media data.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/canperformmultiplepasses
func (a_ AssetWriterInput) CanPerformMultiplePasses() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformMultiplePasses"))
	return rv
}


// SetCanPerformMultiplePasses sets the value of the canPerformMultiplePasses property.
// A Boolean value that indicates whether the input may perform multiple passes over appended media data.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/canperformmultiplepasses
func (a_ AssetWriterInput) SetCanPerformMultiplePasses(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCanPerformMultiplePasses:"), value)
}

// An object that describes the requirements for the current pass.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/currentpassdescription
func (a_ AssetWriterInput) CurrentPassDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentPassDescription"))
	return rv
}


// SetCurrentPassDescription sets the value of the currentPassDescription property.
// An object that describes the requirements for the current pass.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/currentpassdescription
func (a_ AssetWriterInput) SetCurrentPassDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentPassDescription:"), value)
}

// A Boolean value that indicates whether the input tailors its processing for real-time sources.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/expectsmediadatainrealtime
func (a_ AssetWriterInput) ExpectsMediaDataInRealTime() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("expectsMediaDataInRealTime"))
	return rv
}


// SetExpectsMediaDataInRealTime sets the value of the expectsMediaDataInRealTime property.
// A Boolean value that indicates whether the input tailors its processing for real-time sources.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/expectsmediadatainrealtime
func (a_ AssetWriterInput) SetExpectsMediaDataInRealTime(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExpectsMediaDataInRealTime:"), value)
}

// The extended language for the input’s track.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/extendedlanguagetag
func (a_ AssetWriterInput) ExtendedLanguageTag() string {
	rv := objc.Send[string](a_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// SetExtendedLanguageTag sets the value of the extendedLanguageTag property.
// The extended language for the input’s track.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/extendedlanguagetag
func (a_ AssetWriterInput) SetExtendedLanguageTag(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExtendedLanguageTag:"), objc.String(value))
}

// A Boolean value that indicates whether the input is ready to accept media data.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/isreadyformoremediadata
func (a_ AssetWriterInput) IsReadyForMoreMediaData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReadyForMoreMediaData"))
	return rv
}


// SetIsReadyForMoreMediaData sets the value of the isReadyForMoreMediaData property.
// A Boolean value that indicates whether the input is ready to accept media data.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/isreadyformoremediadata
func (a_ AssetWriterInput) SetIsReadyForMoreMediaData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsReadyForMoreMediaData:"), value)
}

// The language code of the input’s track.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/languagecode
func (a_ AssetWriterInput) LanguageCode() string {
	rv := objc.Send[string](a_.ID, objc.Sel("languageCode"))
	return rv
}


// SetLanguageCode sets the value of the languageCode property.
// The language code of the input’s track.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/languagecode
func (a_ AssetWriterInput) SetLanguageCode(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLanguageCode:"), objc.String(value))
}

// A Boolean value that indicates whether to enable a track in the output for playback and processing.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/marksoutputtrackasenabled
func (a_ AssetWriterInput) MarksOutputTrackAsEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("marksOutputTrackAsEnabled"))
	return rv
}


// SetMarksOutputTrackAsEnabled sets the value of the marksOutputTrackAsEnabled property.
// A Boolean value that indicates whether to enable a track in the output for playback and processing.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/marksoutputtrackasenabled
func (a_ AssetWriterInput) SetMarksOutputTrackAsEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMarksOutputTrackAsEnabled:"), value)
}

// Specifies how the input lays out and interleaves media data.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediadatalocation-swift.property
func (a_ AssetWriterInput) MediaDataLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("mediaDataLocation"))
	return rv
}


// SetMediaDataLocation sets the value of the mediaDataLocation property.
// Specifies how the input lays out and interleaves media data.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediadatalocation-swift.property
func (a_ AssetWriterInput) SetMediaDataLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaDataLocation:"), value)
}

// The time scale of the track in the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediatimescale
func (a_ AssetWriterInput) MediaTimeScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("mediaTimeScale"))
	return rv
}


// SetMediaTimeScale sets the value of the mediaTimeScale property.
// The time scale of the track in the output file.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediatimescale
func (a_ AssetWriterInput) SetMediaTimeScale(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaTimeScale:"), value)
}

// The media type of the samples that the input accepts.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediatype
func (a_ AssetWriterInput) MediaType() MediaType {
	rv := objc.Send[MediaType](a_.ID, objc.Sel("mediaType"))
	return rv
}


// SetMediaType sets the value of the mediaType property.
// The media type of the samples that the input accepts.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/mediatype
func (a_ AssetWriterInput) SetMediaType(value MediaType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaType:"), value)
}

// The track-level metadata to write to the output.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/metadata
func (a_ AssetWriterInput) Metadata() AVMetadataItem {
	rv := objc.Send[AVMetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// The track-level metadata to write to the output.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/metadata
func (a_ AssetWriterInput) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadata:"), value)
}

// The natural display dimensions of the output’s visual media.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/naturalsize
func (a_ AssetWriterInput) NaturalSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](a_.ID, objc.Sel("naturalSize"))
	return rv
}


// SetNaturalSize sets the value of the naturalSize property.
// The natural display dimensions of the output’s visual media.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/naturalsize
func (a_ AssetWriterInput) SetNaturalSize(value coregraphics.CGSize) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNaturalSize:"), value)
}

// The settings to use for encoding media data you append to the output.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/outputsettings
func (a_ AssetWriterInput) OutputSettings() string {
	rv := objc.Send[string](a_.ID, objc.Sel("outputSettings"))
	return rv
}


// SetOutputSettings sets the value of the outputSettings property.
// The settings to use for encoding media data you append to the output.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/outputsettings
func (a_ AssetWriterInput) SetOutputSettings(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputSettings:"), objc.String(value))
}

// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/performsmultipassencodingifsupported
func (a_ AssetWriterInput) PerformsMultiPassEncodingIfSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("performsMultiPassEncodingIfSupported"))
	return rv
}


// SetPerformsMultiPassEncodingIfSupported sets the value of the performsMultiPassEncodingIfSupported property.
// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/performsmultipassencodingifsupported
func (a_ AssetWriterInput) SetPerformsMultiPassEncodingIfSupported(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPerformsMultiPassEncodingIfSupported:"), value)
}

// The boundary, in bytes, for aligning media chunks.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredmediachunkalignment
func (a_ AssetWriterInput) PreferredMediaChunkAlignment() int {
	rv := objc.Send[int](a_.ID, objc.Sel("preferredMediaChunkAlignment"))
	return rv
}


// SetPreferredMediaChunkAlignment sets the value of the preferredMediaChunkAlignment property.
// The boundary, in bytes, for aligning media chunks.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredmediachunkalignment
func (a_ AssetWriterInput) SetPreferredMediaChunkAlignment(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredMediaChunkAlignment:"), value)
}

// The duration to use for each chunk of sample data in the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredmediachunkduration
func (a_ AssetWriterInput) PreferredMediaChunkDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredMediaChunkDuration"))
	return rv
}


// SetPreferredMediaChunkDuration sets the value of the preferredMediaChunkDuration property.
// The duration to use for each chunk of sample data in the output file.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredmediachunkduration
func (a_ AssetWriterInput) SetPreferredMediaChunkDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredMediaChunkDuration:"), value)
}

// The volume to prefer for playback of the output’s audio data.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredvolume
func (a_ AssetWriterInput) PreferredVolume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredVolume"))
	return rv
}


// SetPreferredVolume sets the value of the preferredVolume property.
// The volume to prefer for playback of the output’s audio data.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/preferredvolume
func (a_ AssetWriterInput) SetPreferredVolume(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredVolume:"), value)
}

// The base URL sample references are relative to.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/samplereferencebaseurl
func (a_ AssetWriterInput) SampleReferenceBaseURL() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("sampleReferenceBaseURL"))
	return rv
}


// SetSampleReferenceBaseURL sets the value of the sampleReferenceBaseURL property.
// The base URL sample references are relative to.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/samplereferencebaseurl
func (a_ AssetWriterInput) SetSampleReferenceBaseURL(value foundation.IURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleReferenceBaseURL:"), value)
}

// A hint about the format of the sample buffers to append to the input.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/sourceformathint
func (a_ AssetWriterInput) SourceFormatHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("sourceFormatHint"))
	return rv
}


// SetSourceFormatHint sets the value of the sourceFormatHint property.
// A hint about the format of the sample buffers to append to the input.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/sourceformathint
func (a_ AssetWriterInput) SetSourceFormatHint(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSourceFormatHint:"), value)
}

// The transform to use for display of the output’s visual media.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/transform
func (a_ AssetWriterInput) Transform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](a_.ID, objc.Sel("transform"))
	return rv
}


// SetTransform sets the value of the transform property.
// The transform to use for display of the output’s visual media.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/transform
func (a_ AssetWriterInput) SetTransform(value coregraphics.CGAffineTransform) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransform:"), value)
}


