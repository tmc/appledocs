// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
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
	CanPerformMultiplePasses() bool
	CurrentPassDescription() IAVAssetWriterInputPassDescription
	ExpectsMediaDataInRealTime() bool
	SetExpectsMediaDataInRealTime(value bool)
	ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */
	SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */)
	ReadyForMoreMediaData() bool
	LanguageCode() objc.IObject /* cross-framework: NSString */
	SetLanguageCode(value objc.IObject /* cross-framework: NSString */)
	MarksOutputTrackAsEnabled() bool
	SetMarksOutputTrackAsEnabled(value bool)
	MediaDataLocation() AssetWriterInputMediaDataLocation /* typedef */
	SetMediaDataLocation(value AssetWriterInputMediaDataLocation /* typedef */)
	MediaTimeScale() TimeScale /* not a class type */
	SetMediaTimeScale(value TimeScale /* not a class type */)
	MediaType() MediaType /* typedef */
	Metadata() []MetadataItem
	SetMetadata(value []MetadataItem)
	NaturalSize() corefoundation.CGSize
	SetNaturalSize(value corefoundation.CGSize)
	OutputSettings() foundation.IDictionary
	PerformsMultiPassEncodingIfSupported() bool
	SetPerformsMultiPassEncodingIfSupported(value bool)
	PreferredMediaChunkAlignment() int
	SetPreferredMediaChunkAlignment(value int)
	PreferredMediaChunkDuration() objc.IObject /* cross-framework: Time */
	SetPreferredMediaChunkDuration(value objc.IObject /* cross-framework: Time */)
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	SampleReferenceBaseURL() objc.IObject /* cross-framework: NSURL */
	SetSampleReferenceBaseURL(value objc.IObject /* cross-framework: NSURL */)
	SourceFormatHint() FormatDescriptionRef /* not a class type */
	Transform() corefoundation.CGAffineTransform
	SetTransform(value corefoundation.CGAffineTransform)
	IsReadyForMoreMediaData() bool
	SetIsReadyForMoreMediaData(value bool)


	

	// methods:
	AddTrackAssociationWithTrackOfInputType(input IAVAssetWriterInput, trackAssociationType objc.IObject /* cross-framework: NSString */)
	CanAddTrackAssociationWithTrackOfInputType(input IAVAssetWriterInput, trackAssociationType objc.IObject /* cross-framework: NSString */) bool
	MarkAsFinished()
	MarkCurrentPassAsFinished()
	RespondToEachPassDescriptionOnQueueUsingBlock(queue objectivec.IObject, block objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (ac _AssetWriterInputClass) Alloc() AssetWriterInput {
	rv := objc.Send[AssetWriterInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// Creates an input to append sample buffers of the specified type to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/init(mediaType:outputSettings:)
func NewAssetWriterInputWithMediaTypeOutputSettings(mediaType MediaType /* typedef */, outputSettings foundation.IDictionary) AssetWriterInput {
	instance := getAssetWriterInputClass().Alloc()
	rv := objc.Send[AssetWriterInput](instance.ID, objc.Sel("initWithMediaType:outputSettings:"), mediaType, outputSettings)
	rv.Autorelease()
	return rv
}


// Creates an input that appends sample buffers of the specified type and format hint to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/init(mediaType:outputSettings:sourceFormatHint:)
func NewAssetWriterInputWithMediaTypeOutputSettingsSourceFormatHint(mediaType MediaType /* typedef */, outputSettings foundation.IDictionary, sourceFormatHint FormatDescriptionRef /* not a class type */) AssetWriterInput {
	instance := getAssetWriterInputClass().Alloc()
	rv := objc.Send[AssetWriterInput](instance.ID, objc.Sel("initWithMediaType:outputSettings:sourceFormatHint:"), mediaType, outputSettings, sourceFormatHint)
	rv.Autorelease()
	return rv
}







// Returns a new input to append sample buffers of the specified type to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/assetWriterInputWithMediaType:outputSettings:
func (ac _AssetWriterInputClass) AssetWriterInputWithMediaTypeOutputSettings(mediaType MediaType /* typedef */, outputSettings foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetWriterInputWithMediaType:outputSettings:"), mediaType, outputSettings)
	return rv
}


// Returns a new input that appends sample buffers of the specified type and format hint to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/assetWriterInputWithMediaType:outputSettings:sourceFormatHint:
func (ac _AssetWriterInputClass) AssetWriterInputWithMediaTypeOutputSettingsSourceFormatHint(mediaType MediaType /* typedef */, outputSettings foundation.IDictionary, sourceFormatHint FormatDescriptionRef /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetWriterInputWithMediaType:outputSettings:sourceFormatHint:"), mediaType, outputSettings, sourceFormatHint)
	return rv
}












// Adds an association between input tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/addTrackAssociation(withTrackOf:type:)
func (a_ AssetWriterInput) AddTrackAssociationWithTrackOfInputType(input IAVAssetWriterInput, trackAssociationType objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addTrackAssociationWithTrackOfInput:type:"), input, trackAssociationType)
}


// Determines whether it’s valid to associate another input’s track with this input’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/canAddTrackAssociation(withTrackOf:type:)
func (a_ AssetWriterInput) CanAddTrackAssociationWithTrackOfInputType(input IAVAssetWriterInput, trackAssociationType objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddTrackAssociationWithTrackOfInput:type:"), input, trackAssociationType)
	return rv
}


// Marks the input as finished to indicate that you’re done appending samples to it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/markAsFinished()
func (a_ AssetWriterInput) MarkAsFinished() {
	objc.Send[objc.ID](a_.ID, objc.Sel("markAsFinished"))
}


// Tells the input to analyze the appended media to determine whether it can improve the results by reencoding certain segments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/markCurrentPassAsFinished()
func (a_ AssetWriterInput) MarkCurrentPassAsFinished() {
	objc.Send[objc.ID](a_.ID, objc.Sel("markCurrentPassAsFinished"))
}


// Tells the input to invoke a callback whenever it begins a new pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/respondToEachPassDescription(on:using:)
func (a_ AssetWriterInput) RespondToEachPassDescriptionOnQueueUsingBlock(queue objectivec.IObject, block objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("respondToEachPassDescriptionOnQueue:usingBlock:"), queue, block)
}







// A Boolean value that indicates whether the input may perform multiple passes over appended media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/canPerformMultiplePasses
func (a_ AssetWriterInput) CanPerformMultiplePasses() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canPerformMultiplePasses"))
	return rv
}


// An object that describes the requirements for the current pass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/currentPassDescription
func (a_ AssetWriterInput) CurrentPassDescription() IAVAssetWriterInputPassDescription {
	rv := objc.Send[AssetWriterInputPassDescription](a_.ID, objc.Sel("currentPassDescription"))
	return rv
}


// A Boolean value that indicates whether the input tailors its processing for real-time sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/expectsMediaDataInRealTime
func (a_ AssetWriterInput) ExpectsMediaDataInRealTime() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("expectsMediaDataInRealTime"))
	return rv
}


// A Boolean value that indicates whether the input tailors its processing for real-time sources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/expectsMediaDataInRealTime
func (a_ AssetWriterInput) SetExpectsMediaDataInRealTime(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExpectsMediaDataInRealTime:"), value)
}


// The extended language for the input’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/extendedLanguageTag
func (a_ AssetWriterInput) ExtendedLanguageTag() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("extendedLanguageTag"))
	return rv
}


// The extended language for the input’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/extendedLanguageTag
func (a_ AssetWriterInput) SetExtendedLanguageTag(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setExtendedLanguageTag:"), value)
}


// A Boolean value that indicates whether the input is ready to accept media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/isReadyForMoreMediaData
func (a_ AssetWriterInput) ReadyForMoreMediaData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("readyForMoreMediaData"))
	return rv
}


// The language code of the input’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/languageCode
func (a_ AssetWriterInput) LanguageCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("languageCode"))
	return rv
}


// The language code of the input’s track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/languageCode
func (a_ AssetWriterInput) SetLanguageCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLanguageCode:"), value)
}


// A Boolean value that indicates whether to enable a track in the output for playback and processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/marksOutputTrackAsEnabled
func (a_ AssetWriterInput) MarksOutputTrackAsEnabled() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("marksOutputTrackAsEnabled"))
	return rv
}


// A Boolean value that indicates whether to enable a track in the output for playback and processing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/marksOutputTrackAsEnabled
func (a_ AssetWriterInput) SetMarksOutputTrackAsEnabled(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMarksOutputTrackAsEnabled:"), value)
}


// Specifies how the input lays out and interleaves media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/mediaDataLocation-swift.property
func (a_ AssetWriterInput) MediaDataLocation() AssetWriterInputMediaDataLocation /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("mediaDataLocation"))
	return rv
}


// Specifies how the input lays out and interleaves media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/mediaDataLocation-swift.property
func (a_ AssetWriterInput) SetMediaDataLocation(value AssetWriterInputMediaDataLocation /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaDataLocation:"), value)
}


// The time scale of the track in the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/mediaTimeScale
func (a_ AssetWriterInput) MediaTimeScale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](a_.ID, objc.Sel("mediaTimeScale"))
	return rv
}


// The time scale of the track in the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/mediaTimeScale
func (a_ AssetWriterInput) SetMediaTimeScale(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaTimeScale:"), value)
}


// The media type of the samples that the input accepts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/mediaType
func (a_ AssetWriterInput) MediaType() MediaType /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("mediaType"))
	return rv
}


// The track-level metadata to write to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/metadata
func (a_ AssetWriterInput) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// The track-level metadata to write to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/metadata
func (a_ AssetWriterInput) SetMetadata(value []MetadataItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadata:"), nsArray)
}


// The natural display dimensions of the output’s visual media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/naturalSize
func (a_ AssetWriterInput) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("naturalSize"))
	return rv
}


// The natural display dimensions of the output’s visual media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/naturalSize
func (a_ AssetWriterInput) SetNaturalSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNaturalSize:"), value)
}


// The settings to use for encoding media data you append to the output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/outputSettings
func (a_ AssetWriterInput) OutputSettings() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](a_.ID, objc.Sel("outputSettings"))
	return rv
}


// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/performsMultiPassEncodingIfSupported
func (a_ AssetWriterInput) PerformsMultiPassEncodingIfSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("performsMultiPassEncodingIfSupported"))
	return rv
}


// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/performsMultiPassEncodingIfSupported
func (a_ AssetWriterInput) SetPerformsMultiPassEncodingIfSupported(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPerformsMultiPassEncodingIfSupported:"), value)
}


// The boundary, in bytes, for aligning media chunks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/preferredMediaChunkAlignment
func (a_ AssetWriterInput) PreferredMediaChunkAlignment() int {
	rv := objc.Send[int](a_.ID, objc.Sel("preferredMediaChunkAlignment"))
	return rv
}


// The boundary, in bytes, for aligning media chunks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/preferredMediaChunkAlignment
func (a_ AssetWriterInput) SetPreferredMediaChunkAlignment(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredMediaChunkAlignment:"), value)
}


// The duration to use for each chunk of sample data in the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/preferredMediaChunkDuration
func (a_ AssetWriterInput) PreferredMediaChunkDuration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("preferredMediaChunkDuration"))
	return rv
}


// The duration to use for each chunk of sample data in the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/preferredMediaChunkDuration
func (a_ AssetWriterInput) SetPreferredMediaChunkDuration(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredMediaChunkDuration:"), value)
}


// The volume to prefer for playback of the output’s audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/preferredVolume
func (a_ AssetWriterInput) PreferredVolume() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("preferredVolume"))
	return rv
}


// The volume to prefer for playback of the output’s audio data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/preferredVolume
func (a_ AssetWriterInput) SetPreferredVolume(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredVolume:"), value)
}


// The base URL sample references are relative to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/sampleReferenceBaseURL
func (a_ AssetWriterInput) SampleReferenceBaseURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("sampleReferenceBaseURL"))
	return rv
}


// The base URL sample references are relative to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/sampleReferenceBaseURL
func (a_ AssetWriterInput) SetSampleReferenceBaseURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleReferenceBaseURL:"), value)
}


// A hint about the format of the sample buffers to append to the input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/sourceFormatHint
func (a_ AssetWriterInput) SourceFormatHint() FormatDescriptionRef /* not a class type */ {
	rv := objc.Send[FormatDescriptionRef](a_.ID, objc.Sel("sourceFormatHint"))
	return rv
}


// The transform to use for display of the output’s visual media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/transform
func (a_ AssetWriterInput) Transform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](a_.ID, objc.Sel("transform"))
	return rv
}


// The transform to use for display of the output’s visual media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/transform
func (a_ AssetWriterInput) SetTransform(value corefoundation.CGAffineTransform) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransform:"), value)
}


// A Boolean value that indicates whether the input is ready to accept media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/isreadyformoremediadata
func (a_ AssetWriterInput) IsReadyForMoreMediaData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isReadyForMoreMediaData"))
	return rv
}


// A Boolean value that indicates whether the input is ready to accept media data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriterinput/isreadyformoremediadata
func (a_ AssetWriterInput) SetIsReadyForMoreMediaData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsReadyForMoreMediaData:"), value)
}







