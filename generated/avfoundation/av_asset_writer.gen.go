// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetWriter] class.
var (
	AssetWriterClass     _AssetWriterClass
	AssetWriterClassOnce sync.Once
)

func getAssetWriterClass() _AssetWriterClass {
	AssetWriterClassOnce.Do(func() {
		AssetWriterClass = _AssetWriterClass{objc.GetClass("AVAssetWriter")}
	})
	return AssetWriterClass
}

type _AssetWriterClass struct {
	class objc.Class
}

// An interface definition for the [AssetWriter] class.
type IAssetWriter interface {
	objectivec.IObject
	// properties:
	AvailableMediaTypes() []string /* primitive/slice/pointer */
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	DirectoryForTemporaryFiles() foundation.URL /* not a class type */
	SetDirectoryForTemporaryFiles(value foundation.URL /* not a class type */)
	Error() Error
	InitialMovieFragmentInterval() Time /* not a class type */
	SetInitialMovieFragmentInterval(value Time /* not a class type */)
	InitialMovieFragmentSequenceNumber() int /* primitive/slice/pointer */
	SetInitialMovieFragmentSequenceNumber(value int /* primitive/slice/pointer */)
	InitialSegmentStartTime() Time /* not a class type */
	SetInitialSegmentStartTime(value Time /* not a class type */)
	InputGroups() []AssetWriterInputGroup /* primitive/slice/pointer */
	Inputs() []AssetWriterInput /* primitive/slice/pointer */
	Metadata() []MetadataItem /* primitive/slice/pointer */
	SetMetadata(value []MetadataItem /* primitive/slice/pointer */)
	MovieFragmentInterval() Time /* not a class type */
	SetMovieFragmentInterval(value Time /* not a class type */)
	MovieTimeScale() TimeScale /* not a class type */
	SetMovieTimeScale(value TimeScale /* not a class type */)
	OutputFileType() FileType /* not a class type */
	OutputFileTypeProfile() FileTypeProfile /* not a class type */
	SetOutputFileTypeProfile(value FileTypeProfile /* not a class type */)
	OverallDurationHint() Time /* not a class type */
	SetOverallDurationHint(value Time /* not a class type */)
	PreferredOutputSegmentInterval() Time /* not a class type */
	SetPreferredOutputSegmentInterval(value Time /* not a class type */)
	ProducesCombinableFragments() bool /* primitive/slice/pointer */
	SetProducesCombinableFragments(value bool /* primitive/slice/pointer */)
	ShouldOptimizeForNetworkUse() bool /* primitive/slice/pointer */
	SetShouldOptimizeForNetworkUse(value bool /* primitive/slice/pointer */)
	Status() AssetWriterStatus
	OutputURL() foundation.URL /* not a class type */
	SetOutputURL(value foundation.URL /* not a class type */)
	// methods:
	AddInputGroup(inputGroup IAVAssetWriterInputGroup)
	CanAddInput(input IAVAssetWriterInput) bool /* primitive/slice/pointer */
	CanAddInputGroup(inputGroup IAVAssetWriterInputGroup) bool /* primitive/slice/pointer */
	CanApplyOutputSettingsForMediaType(outputSettings foundation.IDictionary /* already interface */, mediaType MediaType /* not a class type */) bool /* primitive/slice/pointer */
	CancelWriting()
	EndSessionAtSourceTime(endTime Time /* not a class type */)
	FlushSegment()
	StartSessionAtSourceTime(startTime Time /* not a class type */)
}

// An object that writes media data to a container file.
//
// You use an asset writer to write media to file formats such as the QuickTime movie file format and MPEG-4 file format. An asset writer automatically supports interleaving media data from concurrent tracks for efficient playback and storage. It can reencode media samples it writes to the output file, and may also write collections of metadata to the output file.


// An object that writes media data to a container file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter
type AssetWriter struct {
	objectivec.Object
}

// AssetWriterFrom constructs a [AssetWriter] from an unsafe.Pointer.
//
// An object that writes media data to a container file.
func AssetWriterFrom(ptr unsafe.Pointer) AssetWriter {
	return AssetWriter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetWriterClass) Alloc() AssetWriter {
	rv := objc.Send[AssetWriter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetWriterClass) New() AssetWriter {
	rv := objc.Send[AssetWriter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetWriter) Init() AssetWriter {
	rv := objc.Send[AssetWriter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetWriter) Autorelease() AssetWriter {
	rv := objc.Send[AssetWriter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetWriter creates a new AssetWriter instance.
func NewAssetWriter() AssetWriter {
	return getAssetWriterClass().New()
}



// Creates an object that outputs segment data in a specified container format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(contentType:)
func NewAssetWriterWithContentType(outputContentType objectivec.IObject) AssetWriter {
	instance := getAssetWriterClass().Alloc()
	rv := objc.Send[AssetWriter](instance.ID, objc.Sel("initWithContentType:"), outputContentType)
	rv.Autorelease()
	return rv
}


// Creates an object that writes media data to a container file at the output URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(outputURL:fileType:)
func NewAssetWriterWithURLFileTypeError(outputURL foundation.URL /* not a class type */, outputFileType FileType /* not a class type */, outError unsafe.Pointer) AssetWriter {
	instance := getAssetWriterClass().Alloc()
	rv := objc.Send[AssetWriter](instance.ID, objc.Sel("initWithURL:fileType:error:"), outputURL, outputFileType, outError)
	rv.Autorelease()
	return rv
}



// Returns a new object that writes media data to a container file at the output URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(url:fileType:)
func (ac _AssetWriterClass) AssetWriterWithURLFileTypeError(outputURL foundation.URL /* not a class type */, outputFileType FileType /* not a class type */, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("assetWriterWithURL:fileType:error:"), outputURL, outputFileType, outError)
	return rv
}


// Adds an input group to an asset writer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/add(_:)-3san4
func (a_ AssetWriter) AddInputGroup(inputGroup IAVAssetWriterInputGroup) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addInputGroup:"), inputGroup)
}


// Determines whether the asset writer supports adding the input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canAdd(_:)-6al7j
func (a_ AssetWriter) CanAddInput(input IAVAssetWriterInput) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddInput:"), input)
	return rv
}


// Determines whether the asset writer supports adding the input group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canAdd(_:)-8s1oh
func (a_ AssetWriter) CanAddInputGroup(inputGroup IAVAssetWriterInputGroup) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddInputGroup:"), inputGroup)
	return rv
}


// Determines whether the output file format supports the output settings for a specific media type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canApply(outputSettings:forMediaType:)
func (a_ AssetWriter) CanApplyOutputSettingsForMediaType(outputSettings foundation.IDictionary /* already interface */, mediaType MediaType /* not a class type */) bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("canApplyOutputSettings:forMediaType:"), outputSettings, mediaType)
	return rv
}


// Cancels the creation of the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/cancelWriting()
func (a_ AssetWriter) CancelWriting() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelWriting"))
}


// Finishes an asset-writing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/endSession(atSourceTime:)
func (a_ AssetWriter) EndSessionAtSourceTime(endTime Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("endSessionAtSourceTime:"), endTime)
}


// Closes the current segment and outputs it to a delegate method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/flushSegment()
func (a_ AssetWriter) FlushSegment() {
	objc.Send[objc.ID](a_.ID, objc.Sel("flushSegment"))
}


// Starts an asset-writing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/startSession(atSourceTime:)
func (a_ AssetWriter) StartSessionAtSourceTime(startTime Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("startSessionAtSourceTime:"), startTime)
}


// The media types the asset writer supports adding as inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/availableMediaTypes
func (a_ AssetWriter) AvailableMediaTypes() []string /* primitive/slice/pointer */ {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableMediaTypes"))
	return rv
}


// A delegate object that responds to asset-writing events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/delegate
func (a_ AssetWriter) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate object that responds to asset-writing events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/delegate
func (a_ AssetWriter) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}


// A directory to contain temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/directoryForTemporaryFiles
func (a_ AssetWriter) DirectoryForTemporaryFiles() foundation.URL /* not a class type */ {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("directoryForTemporaryFiles"))
	return rv
}


// A directory to contain temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/directoryForTemporaryFiles
func (a_ AssetWriter) SetDirectoryForTemporaryFiles(value foundation.URL /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}


// An error object that describes an asset-writing failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/error
func (a_ AssetWriter) Error() Error {
	rv := objc.Send[Error](a_.ID, objc.Sel("error"))
	return rv
}


// The interval at which to write the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialMovieFragmentInterval
func (a_ AssetWriter) InitialMovieFragmentInterval() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("initialMovieFragmentInterval"))
	return rv
}


// The interval at which to write the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialMovieFragmentInterval
func (a_ AssetWriter) SetInitialMovieFragmentInterval(value Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialMovieFragmentInterval:"), value)
}


// The sequence number of the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialMovieFragmentSequenceNumber
func (a_ AssetWriter) InitialMovieFragmentSequenceNumber() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](a_.ID, objc.Sel("initialMovieFragmentSequenceNumber"))
	return rv
}


// The sequence number of the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialMovieFragmentSequenceNumber
func (a_ AssetWriter) SetInitialMovieFragmentSequenceNumber(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialMovieFragmentSequenceNumber:"), value)
}


// The start time of the initial segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialSegmentStartTime
func (a_ AssetWriter) InitialSegmentStartTime() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("initialSegmentStartTime"))
	return rv
}


// The start time of the initial segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialSegmentStartTime
func (a_ AssetWriter) SetInitialSegmentStartTime(value Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialSegmentStartTime:"), value)
}


// The input groups an asset writer contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/inputGroups
func (a_ AssetWriter) InputGroups() []AssetWriterInputGroup /* primitive/slice/pointer */ {
	rv := objc.Send[[]AssetWriterInputGroup](a_.ID, objc.Sel("inputGroups"))
	return rv
}


// The inputs an asset writer contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/inputs
func (a_ AssetWriter) Inputs() []AssetWriterInput /* primitive/slice/pointer */ {
	rv := objc.Send[[]AssetWriterInput](a_.ID, objc.Sel("inputs"))
	return rv
}


// An array of metadata items to write to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/metadata
func (a_ AssetWriter) Metadata() []MetadataItem /* primitive/slice/pointer */ {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// An array of metadata items to write to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/metadata
func (a_ AssetWriter) SetMetadata(value []MetadataItem /* primitive/slice/pointer */) {
	// Convert Go slice to NSArray
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


// The interval at which to write movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieFragmentInterval
func (a_ AssetWriter) MovieFragmentInterval() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("movieFragmentInterval"))
	return rv
}


// The interval at which to write movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieFragmentInterval
func (a_ AssetWriter) SetMovieFragmentInterval(value Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMovieFragmentInterval:"), value)
}


// The time scale of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieTimeScale
func (a_ AssetWriter) MovieTimeScale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](a_.ID, objc.Sel("movieTimeScale"))
	return rv
}


// The time scale of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieTimeScale
func (a_ AssetWriter) SetMovieTimeScale(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMovieTimeScale:"), value)
}


// The type of container file that the writer outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputFileType
func (a_ AssetWriter) OutputFileType() FileType /* not a class type */ {
	rv := objc.Send[FileType](a_.ID, objc.Sel("outputFileType"))
	return rv
}


// A profile for the output file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputFileTypeProfile
func (a_ AssetWriter) OutputFileTypeProfile() FileTypeProfile /* not a class type */ {
	rv := objc.Send[FileTypeProfile](a_.ID, objc.Sel("outputFileTypeProfile"))
	return rv
}


// A profile for the output file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputFileTypeProfile
func (a_ AssetWriter) SetOutputFileTypeProfile(value FileTypeProfile /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFileTypeProfile:"), value)
}


// A hint of the final duration of the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/overallDurationHint
func (a_ AssetWriter) OverallDurationHint() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("overallDurationHint"))
	return rv
}


// A hint of the final duration of the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/overallDurationHint
func (a_ AssetWriter) SetOverallDurationHint(value Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverallDurationHint:"), value)
}


// The interval of output segments that you prefer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/preferredOutputSegmentInterval
func (a_ AssetWriter) PreferredOutputSegmentInterval() Time /* not a class type */ {
	rv := objc.Send[Time](a_.ID, objc.Sel("preferredOutputSegmentInterval"))
	return rv
}


// The interval of output segments that you prefer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/preferredOutputSegmentInterval
func (a_ AssetWriter) SetPreferredOutputSegmentInterval(value Time /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredOutputSegmentInterval:"), value)
}


// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/producesCombinableFragments
func (a_ AssetWriter) ProducesCombinableFragments() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("producesCombinableFragments"))
	return rv
}


// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/producesCombinableFragments
func (a_ AssetWriter) SetProducesCombinableFragments(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProducesCombinableFragments:"), value)
}


// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/shouldOptimizeForNetworkUse
func (a_ AssetWriter) ShouldOptimizeForNetworkUse() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}


// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/shouldOptimizeForNetworkUse
func (a_ AssetWriter) SetShouldOptimizeForNetworkUse(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}


// The status of writing samples to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/status-swift.property
func (a_ AssetWriter) Status() AssetWriterStatus {
	rv := objc.Send[AssetWriterStatus](a_.ID, objc.Sel("status"))
	return rv
}


// The location of the container file that the writer outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/outputurl
func (a_ AssetWriter) OutputURL() foundation.URL /* not a class type */ {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("outputURL"))
	return rv
}


// The location of the container file that the writer outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/outputurl
func (a_ AssetWriter) SetOutputURL(value foundation.URL /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputURL:"), value)
}


