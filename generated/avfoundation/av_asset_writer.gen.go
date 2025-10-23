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
	AvailableMediaTypes() unsafe.Pointer
	SetAvailableMediaTypes(value unsafe.Pointer)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DirectoryForTemporaryFiles() foundation.URL
	SetDirectoryForTemporaryFiles(value foundation.URL)
	Error() AVError
	SetError(value AVError)
	InitialMovieFragmentInterval() unsafe.Pointer
	SetInitialMovieFragmentInterval(value unsafe.Pointer)
	InitialMovieFragmentSequenceNumber() int
	SetInitialMovieFragmentSequenceNumber(value int)
	InitialSegmentStartTime() unsafe.Pointer
	SetInitialSegmentStartTime(value unsafe.Pointer)
	InputGroups() unsafe.Pointer
	SetInputGroups(value unsafe.Pointer)
	Inputs() IAVAssetWriterInput
	SetInputs(value IAVAssetWriterInput)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	MovieFragmentInterval() unsafe.Pointer
	SetMovieFragmentInterval(value unsafe.Pointer)
	MovieTimeScale() unsafe.Pointer
	SetMovieTimeScale(value unsafe.Pointer)
	OutputFileType() unsafe.Pointer
	SetOutputFileType(value unsafe.Pointer)
	OutputFileTypeProfile() unsafe.Pointer
	SetOutputFileTypeProfile(value unsafe.Pointer)
	OutputURL() foundation.URL
	SetOutputURL(value foundation.URL)
	OverallDurationHint() unsafe.Pointer
	SetOverallDurationHint(value unsafe.Pointer)
	PreferredOutputSegmentInterval() unsafe.Pointer
	SetPreferredOutputSegmentInterval(value unsafe.Pointer)
	ProducesCombinableFragments() bool
	SetProducesCombinableFragments(value bool)
	ShouldOptimizeForNetworkUse() bool
	SetShouldOptimizeForNetworkUse(value bool)
	Status() unsafe.Pointer
	SetStatus(value unsafe.Pointer)
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



// The media types the asset writer supports adding as inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/availablemediatypes
func (a_ AssetWriter) AvailableMediaTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("availableMediaTypes"))
	return rv
}


// The media types the asset writer supports adding as inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/availablemediatypes
func (a_ AssetWriter) SetAvailableMediaTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAvailableMediaTypes:"), value)
}


// A delegate object that responds to asset-writing events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/delegate
func (a_ AssetWriter) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate object that responds to asset-writing events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/delegate
func (a_ AssetWriter) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}


// A directory to contain temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/directoryfortemporaryfiles
func (a_ AssetWriter) DirectoryForTemporaryFiles() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("directoryForTemporaryFiles"))
	return rv
}


// A directory to contain temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/directoryfortemporaryfiles
func (a_ AssetWriter) SetDirectoryForTemporaryFiles(value foundation.URL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}


// An error object that describes an asset-writing failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/error
func (a_ AssetWriter) Error() AVError {
	rv := objc.Send[Error](a_.ID, objc.Sel("error"))
	return rv
}


// An error object that describes an asset-writing failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/error
func (a_ AssetWriter) SetError(value AVError) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setError:"), value)
}


// The interval at which to write the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/initialmoviefragmentinterval
func (a_ AssetWriter) InitialMovieFragmentInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("initialMovieFragmentInterval"))
	return rv
}


// The interval at which to write the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/initialmoviefragmentinterval
func (a_ AssetWriter) SetInitialMovieFragmentInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialMovieFragmentInterval:"), value)
}


// The sequence number of the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/initialmoviefragmentsequencenumber
func (a_ AssetWriter) InitialMovieFragmentSequenceNumber() int {
	rv := objc.Send[int](a_.ID, objc.Sel("initialMovieFragmentSequenceNumber"))
	return rv
}


// The sequence number of the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/initialmoviefragmentsequencenumber
func (a_ AssetWriter) SetInitialMovieFragmentSequenceNumber(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialMovieFragmentSequenceNumber:"), value)
}


// The start time of the initial segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/initialsegmentstarttime
func (a_ AssetWriter) InitialSegmentStartTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("initialSegmentStartTime"))
	return rv
}


// The start time of the initial segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/initialsegmentstarttime
func (a_ AssetWriter) SetInitialSegmentStartTime(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialSegmentStartTime:"), value)
}


// The input groups an asset writer contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/inputgroups
func (a_ AssetWriter) InputGroups() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("inputGroups"))
	return rv
}


// The input groups an asset writer contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/inputgroups
func (a_ AssetWriter) SetInputGroups(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputGroups:"), value)
}


// The inputs an asset writer contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/inputs
func (a_ AssetWriter) Inputs() IAVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](a_.ID, objc.Sel("inputs"))
	return rv
}


// The inputs an asset writer contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/inputs
func (a_ AssetWriter) SetInputs(value IAVAssetWriterInput) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInputs:"), value)
}


// An array of metadata items to write to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/metadata
func (a_ AssetWriter) Metadata() IAVMetadataItem {
	rv := objc.Send[AVMetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}


// An array of metadata items to write to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/metadata
func (a_ AssetWriter) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMetadata:"), value)
}


// The interval at which to write movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/moviefragmentinterval
func (a_ AssetWriter) MovieFragmentInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("movieFragmentInterval"))
	return rv
}


// The interval at which to write movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/moviefragmentinterval
func (a_ AssetWriter) SetMovieFragmentInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMovieFragmentInterval:"), value)
}


// The time scale of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/movietimescale
func (a_ AssetWriter) MovieTimeScale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("movieTimeScale"))
	return rv
}


// The time scale of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/movietimescale
func (a_ AssetWriter) SetMovieTimeScale(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMovieTimeScale:"), value)
}


// The type of container file that the writer outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/outputfiletype
func (a_ AssetWriter) OutputFileType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputFileType"))
	return rv
}


// The type of container file that the writer outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/outputfiletype
func (a_ AssetWriter) SetOutputFileType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFileType:"), value)
}


// A profile for the output file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/outputfiletypeprofile
func (a_ AssetWriter) OutputFileTypeProfile() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("outputFileTypeProfile"))
	return rv
}


// A profile for the output file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/outputfiletypeprofile
func (a_ AssetWriter) SetOutputFileTypeProfile(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFileTypeProfile:"), value)
}


// The location of the container file that the writer outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/outputurl
func (a_ AssetWriter) OutputURL() foundation.URL {
	rv := objc.Send[foundation.URL](a_.ID, objc.Sel("outputURL"))
	return rv
}


// The location of the container file that the writer outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/outputurl
func (a_ AssetWriter) SetOutputURL(value foundation.URL) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputURL:"), value)
}


// A hint of the final duration of the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/overalldurationhint
func (a_ AssetWriter) OverallDurationHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("overallDurationHint"))
	return rv
}


// A hint of the final duration of the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/overalldurationhint
func (a_ AssetWriter) SetOverallDurationHint(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverallDurationHint:"), value)
}


// The interval of output segments that you prefer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/preferredoutputsegmentinterval
func (a_ AssetWriter) PreferredOutputSegmentInterval() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredOutputSegmentInterval"))
	return rv
}


// The interval of output segments that you prefer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/preferredoutputsegmentinterval
func (a_ AssetWriter) SetPreferredOutputSegmentInterval(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredOutputSegmentInterval:"), value)
}


// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/producescombinablefragments
func (a_ AssetWriter) ProducesCombinableFragments() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("producesCombinableFragments"))
	return rv
}


// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/producescombinablefragments
func (a_ AssetWriter) SetProducesCombinableFragments(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProducesCombinableFragments:"), value)
}


// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/shouldoptimizefornetworkuse
func (a_ AssetWriter) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}


// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/shouldoptimizefornetworkuse
func (a_ AssetWriter) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}


// The status of writing samples to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/status-swift.property
func (a_ AssetWriter) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("status"))
	return rv
}


// The status of writing samples to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetwriter/status-swift.property
func (a_ AssetWriter) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setStatus:"), value)
}



