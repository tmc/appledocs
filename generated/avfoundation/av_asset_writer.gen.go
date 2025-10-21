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
	AddInput(input IAVAssetWriterInput)
	CanAddInput(input IAVAssetWriterInput) bool
	CanAddInputGroup(inputGroup unsafe.Pointer) bool
	CancelWriting()
	FinishWriting() bool
	FinishWritingWithCompletionHandler(handler unsafe.Pointer)
	StartWriting() bool
}

// An object that writes media data to a container file.
//
// You use an asset writer to write media to file formats such as the QuickTime movie file format and MPEG-4 file format. An asset writer automatically supports interleaving media data from concurrent tracks for efficient playback and storage. It can reencode media samples it writes to the output file, and may also write collections of metadata to the output file.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(contentType:)
func NewAssetWriterWithContentType(outputContentType unsafe.Pointer) AssetWriter {
	instance := getAssetWriterClass().Alloc()
	rv := objc.Send[AssetWriter](instance.ID, objc.Sel("initWithContentType:"), outputContentType)
	rv.Autorelease()
	return rv
}


// Adds an input to an asset writer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/add(_:)-4c4d0
func (a_ AssetWriter) AddInput(input IAVAssetWriterInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addInput:"), input)
}

// Determines whether the asset writer supports adding the input.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canAdd(_:)-6al7j
func (a_ AssetWriter) CanAddInput(input IAVAssetWriterInput) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddInput:"), input)
	return rv
}

// Determines whether the asset writer supports adding the input group.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canAdd(_:)-8s1oh
func (a_ AssetWriter) CanAddInputGroup(inputGroup unsafe.Pointer) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddInputGroup:"), inputGroup)
	return rv
}

// Cancels the creation of the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/cancelWriting()
func (a_ AssetWriter) CancelWriting() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelWriting"))
}

// Completes the writing of the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/finishWriting()
func (a_ AssetWriter) FinishWriting() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("finishWriting"))
	return rv
}

// Marks all unfinished inputs as finished and completes the writing of the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/finishWriting(completionHandler:)
func (a_ AssetWriter) FinishWritingWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishWritingWithCompletionHandler:"), handler)
}

// Tells the writer to start writing its output.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/startWriting()
func (a_ AssetWriter) StartWriting() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("startWriting"))
	return rv
}

// The media types the asset writer supports adding as inputs.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/availableMediaTypes
func (a_ AssetWriter) AvailableMediaTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableMediaTypes"))
	return rv
}

// A delegate object that responds to asset-writing events.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/delegate
func (a_ AssetWriter) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate object that responds to asset-writing events.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/delegate
func (a_ AssetWriter) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}

// A directory to contain temporary files that the export process generates.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/directoryForTemporaryFiles
func (a_ AssetWriter) DirectoryForTemporaryFiles() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("directoryForTemporaryFiles"))
	return rv
}


// SetDirectoryForTemporaryFiles sets the value of the directoryForTemporaryFiles property.
// A directory to contain temporary files that the export process generates.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/directoryForTemporaryFiles
func (a_ AssetWriter) SetDirectoryForTemporaryFiles(value foundation.IURL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}

// An error object that describes an asset-writing failure.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/error
func (a_ AssetWriter) Error() Error {
	rv := objc.Send[Error](a_.ID, objc.Sel("error"))
	return rv
}

// The start time of the initial segment.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialSegmentStartTime
func (a_ AssetWriter) InitialSegmentStartTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("initialSegmentStartTime"))
	return rv
}


// SetInitialSegmentStartTime sets the value of the initialSegmentStartTime property.
// The start time of the initial segment.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialSegmentStartTime
func (a_ AssetWriter) SetInitialSegmentStartTime(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialSegmentStartTime:"), value)
}

// An array of metadata items to write to the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/metadata
func (a_ AssetWriter) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// SetMetadata sets the value of the metadata property.
// An array of metadata items to write to the output file.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/metadata
func (a_ AssetWriter) SetMetadata(value []MetadataItem) {
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
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieFragmentInterval
func (a_ AssetWriter) MovieFragmentInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("movieFragmentInterval"))
	return rv
}


// SetMovieFragmentInterval sets the value of the movieFragmentInterval property.
// The interval at which to write movie fragments.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieFragmentInterval
func (a_ AssetWriter) SetMovieFragmentInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMovieFragmentInterval:"), value)
}

// The time scale of the movie.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieTimeScale
func (a_ AssetWriter) MovieTimeScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("movieTimeScale"))
	return rv
}


// SetMovieTimeScale sets the value of the movieTimeScale property.
// The time scale of the movie.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieTimeScale
func (a_ AssetWriter) SetMovieTimeScale(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMovieTimeScale:"), value)
}

// The type of container file that the writer outputs.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputFileType
func (a_ AssetWriter) OutputFileType() FileType {
	rv := objc.Send[FileType](a_.ID, objc.Sel("outputFileType"))
	return rv
}

// The location of the container file that the writer outputs.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputURL
func (a_ AssetWriter) OutputURL() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("outputURL"))
	return rv
}

// The interval of output segments that you prefer.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/preferredOutputSegmentInterval
func (a_ AssetWriter) PreferredOutputSegmentInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredOutputSegmentInterval"))
	return rv
}


// SetPreferredOutputSegmentInterval sets the value of the preferredOutputSegmentInterval property.
// The interval of output segments that you prefer.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/preferredOutputSegmentInterval
func (a_ AssetWriter) SetPreferredOutputSegmentInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredOutputSegmentInterval:"), value)
}

// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/producesCombinableFragments
func (a_ AssetWriter) ProducesCombinableFragments() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("producesCombinableFragments"))
	return rv
}


// SetProducesCombinableFragments sets the value of the producesCombinableFragments property.
// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/producesCombinableFragments
func (a_ AssetWriter) SetProducesCombinableFragments(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProducesCombinableFragments:"), value)
}

// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/shouldOptimizeForNetworkUse
func (a_ AssetWriter) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}


// SetShouldOptimizeForNetworkUse sets the value of the shouldOptimizeForNetworkUse property.
// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/shouldOptimizeForNetworkUse
func (a_ AssetWriter) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}

// The status of writing samples to the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/status-swift.property
func (a_ AssetWriter) Status() AssetWriterStatus {
	rv := objc.Send[AssetWriterStatus](a_.ID, objc.Sel("status"))
	return rv
}

// The interval at which to write the initial movie fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/initialmoviefragmentinterval
func (a_ AssetWriter) InitialMovieFragmentInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("initialMovieFragmentInterval"))
	return rv
}


// SetInitialMovieFragmentInterval sets the value of the initialMovieFragmentInterval property.
// The interval at which to write the initial movie fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/initialmoviefragmentinterval
func (a_ AssetWriter) SetInitialMovieFragmentInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialMovieFragmentInterval:"), value)
}

// The sequence number of the initial movie fragment.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/initialmoviefragmentsequencenumber
func (a_ AssetWriter) InitialMovieFragmentSequenceNumber() int {
	rv := objc.Send[int](a_.ID, objc.Sel("initialMovieFragmentSequenceNumber"))
	return rv
}


// SetInitialMovieFragmentSequenceNumber sets the value of the initialMovieFragmentSequenceNumber property.
// The sequence number of the initial movie fragment.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/initialmoviefragmentsequencenumber
func (a_ AssetWriter) SetInitialMovieFragmentSequenceNumber(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialMovieFragmentSequenceNumber:"), value)
}

// The input groups an asset writer contains.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/inputgroups
func (a_ AssetWriter) InputGroups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("inputGroups"))
	return rv
}


// SetInputGroups sets the value of the inputGroups property.
// The input groups an asset writer contains.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/inputgroups
func (a_ AssetWriter) SetInputGroups(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputGroups:"), value)
}

// The inputs an asset writer contains.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/inputs
func (a_ AssetWriter) Inputs() AVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](a_.ID, objc.Sel("inputs"))
	return rv
}


// SetInputs sets the value of the inputs property.
// The inputs an asset writer contains.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/inputs
func (a_ AssetWriter) SetInputs(value IAVAssetWriterInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputs:"), value)
}

// A profile for the output file type.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/outputfiletypeprofile
func (a_ AssetWriter) OutputFileTypeProfile() FileTypeProfile {
	rv := objc.Send[FileTypeProfile](a_.ID, objc.Sel("outputFileTypeProfile"))
	return rv
}


// SetOutputFileTypeProfile sets the value of the outputFileTypeProfile property.
// A profile for the output file type.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/outputfiletypeprofile
func (a_ AssetWriter) SetOutputFileTypeProfile(value IFileTypeProfile) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFileTypeProfile:"), value)
}

// A hint of the final duration of the output file.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/overalldurationhint
func (a_ AssetWriter) OverallDurationHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("overallDurationHint"))
	return rv
}


// SetOverallDurationHint sets the value of the overallDurationHint property.
// A hint of the final duration of the output file.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/overalldurationhint
func (a_ AssetWriter) SetOverallDurationHint(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverallDurationHint:"), value)
}


