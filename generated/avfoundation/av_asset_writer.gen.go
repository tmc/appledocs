// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetWriter */


/* debug [class_header]: Header for AVAssetWriter */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetWriter */
// An interface definition for the [AssetWriter] class.
type IAssetWriter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetWriter */
	// properties:
	AvailableMediaTypes() []string
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DirectoryForTemporaryFiles() objc.IObject /* cross-framework: NSURL */
	SetDirectoryForTemporaryFiles(value objc.IObject /* cross-framework: NSURL */)
	Error() Error
	InitialMovieFragmentInterval() objc.IObject /* cross-framework: Time */
	SetInitialMovieFragmentInterval(value objc.IObject /* cross-framework: Time */)
	InitialMovieFragmentSequenceNumber() int
	SetInitialMovieFragmentSequenceNumber(value int)
	InitialSegmentStartTime() objc.IObject /* cross-framework: Time */
	SetInitialSegmentStartTime(value objc.IObject /* cross-framework: Time */)
	InputGroups() []AssetWriterInputGroup
	Inputs() []AssetWriterInput
	Metadata() []MetadataItem
	SetMetadata(value []MetadataItem)
	MovieFragmentInterval() objc.IObject /* cross-framework: Time */
	SetMovieFragmentInterval(value objc.IObject /* cross-framework: Time */)
	MovieTimeScale() TimeScale /* not a class type */
	SetMovieTimeScale(value TimeScale /* not a class type */)
	OutputFileType() FileType /* typedef */
	OutputFileTypeProfile() FileTypeProfile /* typedef */
	SetOutputFileTypeProfile(value FileTypeProfile /* typedef */)
	OutputURL() objc.IObject /* cross-framework: NSURL */
	OverallDurationHint() objc.IObject /* cross-framework: Time */
	SetOverallDurationHint(value objc.IObject /* cross-framework: Time */)
	PreferredOutputSegmentInterval() objc.IObject /* cross-framework: Time */
	SetPreferredOutputSegmentInterval(value objc.IObject /* cross-framework: Time */)
	ProducesCombinableFragments() bool
	SetProducesCombinableFragments(value bool)
	ShouldOptimizeForNetworkUse() bool
	SetShouldOptimizeForNetworkUse(value bool)
	Status() AssetWriterStatus
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetWriter */
	// methods:
	AddInputGroup(inputGroup IAVAssetWriterInputGroup)
	CanAddInput(input IAVAssetWriterInput) bool
	CanAddInputGroup(inputGroup IAVAssetWriterInputGroup) bool
	CanApplyOutputSettingsForMediaType(outputSettings foundation.IDictionary, mediaType MediaType /* typedef */) bool
	CancelWriting()
	EndSessionAtSourceTime(endTime objc.IObject /* cross-framework: Time */)
	FinishWritingWithCompletionHandler(handler unsafe.Pointer)
	FlushSegment()
	StartSessionAtSourceTime(startTime objc.IObject /* cross-framework: Time */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetWriter */
// Alloc allocates a new instance without initialization.
func (ac _AssetWriterClass) Alloc() AssetWriter {
	rv := objc.Send[AssetWriter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetWriter */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetWriter */

// Creates an object that outputs segment data in a specified container format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(contentType:)
func NewAssetWriterWithContentType(outputContentType uniformtypeidentifiers.UTType) AssetWriter {
	instance := getAssetWriterClass().Alloc()
	rv := objc.Send[AssetWriter](instance.ID, objc.Sel("initWithContentType:"), outputContentType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetWriterWithContentType */


// Creates an object that writes media data to a container file at the output URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(outputURL:fileType:)
func NewAssetWriterWithURLFileTypeError(outputURL objc.IObject /* cross-framework: NSURL */, outputFileType FileType /* typedef */, outError objectivec.IObject) AssetWriter {
	instance := getAssetWriterClass().Alloc()
	rv := objc.Send[AssetWriter](instance.ID, objc.Sel("initWithURL:fileType:error:"), outputURL, outputFileType, outError)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetWriterWithURLFileTypeError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetWriter */

// Returns a new object that writes media data to a container file at the output URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/init(url:fileType:)
func (ac _AssetWriterClass) AssetWriterWithURLFileTypeError(outputURL objc.IObject /* cross-framework: NSURL */, outputFileType FileType /* typedef */, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetWriterWithURL:fileType:error:"), outputURL, outputFileType, outError)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetWriterWithURLFileTypeError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetWriter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetWriter */

// Adds an input group to an asset writer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/add(_:)-3san4
func (a_ AssetWriter) AddInputGroup(inputGroup IAVAssetWriterInputGroup) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addInputGroup:"), inputGroup)
}/* debug [instance_methods/method]: AddInputGroup */


// Determines whether the asset writer supports adding the input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canAdd(_:)-6al7j
func (a_ AssetWriter) CanAddInput(input IAVAssetWriterInput) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddInput:"), input)
	return rv
}/* debug [instance_methods/method]: CanAddInput */


// Determines whether the asset writer supports adding the input group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canAdd(_:)-8s1oh
func (a_ AssetWriter) CanAddInputGroup(inputGroup IAVAssetWriterInputGroup) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canAddInputGroup:"), inputGroup)
	return rv
}/* debug [instance_methods/method]: CanAddInputGroup */


// Determines whether the output file format supports the output settings for a specific media type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/canApply(outputSettings:forMediaType:)
func (a_ AssetWriter) CanApplyOutputSettingsForMediaType(outputSettings foundation.IDictionary, mediaType MediaType /* typedef */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("canApplyOutputSettings:forMediaType:"), outputSettings, mediaType)
	return rv
}/* debug [instance_methods/method]: CanApplyOutputSettingsForMediaType */


// Cancels the creation of the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/cancelWriting()
func (a_ AssetWriter) CancelWriting() {
	objc.Send[objc.ID](a_.ID, objc.Sel("cancelWriting"))
}/* debug [instance_methods/method]: CancelWriting */


// Finishes an asset-writing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/endSession(atSourceTime:)
func (a_ AssetWriter) EndSessionAtSourceTime(endTime objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("endSessionAtSourceTime:"), endTime)
}/* debug [instance_methods/method]: EndSessionAtSourceTime */


// Marks all unfinished inputs as finished and completes the writing of the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/finishWriting(completionHandler:)
func (a_ AssetWriter) FinishWritingWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("finishWritingWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: FinishWritingWithCompletionHandler */


// Closes the current segment and outputs it to a delegate method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/flushSegment()
func (a_ AssetWriter) FlushSegment() {
	objc.Send[objc.ID](a_.ID, objc.Sel("flushSegment"))
}/* debug [instance_methods/method]: FlushSegment */


// Starts an asset-writing session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/startSession(atSourceTime:)
func (a_ AssetWriter) StartSessionAtSourceTime(startTime objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("startSessionAtSourceTime:"), startTime)
}/* debug [instance_methods/method]: StartSessionAtSourceTime */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetWriter */

// The media types the asset writer supports adding as inputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/availableMediaTypes
func (a_ AssetWriter) AvailableMediaTypes() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("availableMediaTypes"))
	return rv
}/* debug [instance_properties/getter]: availableMediaTypes */


// A delegate object that responds to asset-writing events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/delegate
func (a_ AssetWriter) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// A delegate object that responds to asset-writing events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/delegate
func (a_ AssetWriter) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A directory to contain temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/directoryForTemporaryFiles
func (a_ AssetWriter) DirectoryForTemporaryFiles() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("directoryForTemporaryFiles"))
	return rv
}/* debug [instance_properties/getter]: directoryForTemporaryFiles */


// A directory to contain temporary files that the export process generates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/directoryForTemporaryFiles
func (a_ AssetWriter) SetDirectoryForTemporaryFiles(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDirectoryForTemporaryFiles:"), value)
}/* debug [instance_properties/setter]: directoryForTemporaryFiles */


// An error object that describes an asset-writing failure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/error
func (a_ AssetWriter) Error() Error {
	rv := objc.Send[Error](a_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The interval at which to write the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialMovieFragmentInterval
func (a_ AssetWriter) InitialMovieFragmentInterval() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("initialMovieFragmentInterval"))
	return rv
}/* debug [instance_properties/getter]: initialMovieFragmentInterval */


// The interval at which to write the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialMovieFragmentInterval
func (a_ AssetWriter) SetInitialMovieFragmentInterval(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialMovieFragmentInterval:"), value)
}/* debug [instance_properties/setter]: initialMovieFragmentInterval */


// The sequence number of the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialMovieFragmentSequenceNumber
func (a_ AssetWriter) InitialMovieFragmentSequenceNumber() int {
	rv := objc.Send[int](a_.ID, objc.Sel("initialMovieFragmentSequenceNumber"))
	return rv
}/* debug [instance_properties/getter]: initialMovieFragmentSequenceNumber */


// The sequence number of the initial movie fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialMovieFragmentSequenceNumber
func (a_ AssetWriter) SetInitialMovieFragmentSequenceNumber(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialMovieFragmentSequenceNumber:"), value)
}/* debug [instance_properties/setter]: initialMovieFragmentSequenceNumber */


// The start time of the initial segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialSegmentStartTime
func (a_ AssetWriter) InitialSegmentStartTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("initialSegmentStartTime"))
	return rv
}/* debug [instance_properties/getter]: initialSegmentStartTime */


// The start time of the initial segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/initialSegmentStartTime
func (a_ AssetWriter) SetInitialSegmentStartTime(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setInitialSegmentStartTime:"), value)
}/* debug [instance_properties/setter]: initialSegmentStartTime */


// The input groups an asset writer contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/inputGroups
func (a_ AssetWriter) InputGroups() []AssetWriterInputGroup {
	rv := objc.Send[[]AssetWriterInputGroup](a_.ID, objc.Sel("inputGroups"))
	return rv
}/* debug [instance_properties/getter]: inputGroups */


// The inputs an asset writer contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/inputs
func (a_ AssetWriter) Inputs() []AssetWriterInput {
	rv := objc.Send[[]AssetWriterInput](a_.ID, objc.Sel("inputs"))
	return rv
}/* debug [instance_properties/getter]: inputs */


// An array of metadata items to write to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/metadata
func (a_ AssetWriter) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](a_.ID, objc.Sel("metadata"))
	return rv
}/* debug [instance_properties/getter]: metadata */


// An array of metadata items to write to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/metadata
func (a_ AssetWriter) SetMetadata(value []MetadataItem) {
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
}/* debug [instance_properties/setter]: metadata */


// The interval at which to write movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieFragmentInterval
func (a_ AssetWriter) MovieFragmentInterval() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("movieFragmentInterval"))
	return rv
}/* debug [instance_properties/getter]: movieFragmentInterval */


// The interval at which to write movie fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieFragmentInterval
func (a_ AssetWriter) SetMovieFragmentInterval(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMovieFragmentInterval:"), value)
}/* debug [instance_properties/setter]: movieFragmentInterval */


// The time scale of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieTimeScale
func (a_ AssetWriter) MovieTimeScale() TimeScale /* not a class type */ {
	rv := objc.Send[TimeScale](a_.ID, objc.Sel("movieTimeScale"))
	return rv
}/* debug [instance_properties/getter]: movieTimeScale */


// The time scale of the movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/movieTimeScale
func (a_ AssetWriter) SetMovieTimeScale(value TimeScale /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMovieTimeScale:"), value)
}/* debug [instance_properties/setter]: movieTimeScale */


// The type of container file that the writer outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputFileType
func (a_ AssetWriter) OutputFileType() FileType /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("outputFileType"))
	return rv
}/* debug [instance_properties/getter]: outputFileType */


// A profile for the output file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputFileTypeProfile
func (a_ AssetWriter) OutputFileTypeProfile() FileTypeProfile /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("outputFileTypeProfile"))
	return rv
}/* debug [instance_properties/getter]: outputFileTypeProfile */


// A profile for the output file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputFileTypeProfile
func (a_ AssetWriter) SetOutputFileTypeProfile(value FileTypeProfile /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOutputFileTypeProfile:"), value)
}/* debug [instance_properties/setter]: outputFileTypeProfile */


// The location of the container file that the writer outputs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/outputURL
func (a_ AssetWriter) OutputURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("outputURL"))
	return rv
}/* debug [instance_properties/getter]: outputURL */


// A hint of the final duration of the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/overallDurationHint
func (a_ AssetWriter) OverallDurationHint() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("overallDurationHint"))
	return rv
}/* debug [instance_properties/getter]: overallDurationHint */


// A hint of the final duration of the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/overallDurationHint
func (a_ AssetWriter) SetOverallDurationHint(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setOverallDurationHint:"), value)
}/* debug [instance_properties/setter]: overallDurationHint */


// The interval of output segments that you prefer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/preferredOutputSegmentInterval
func (a_ AssetWriter) PreferredOutputSegmentInterval() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("preferredOutputSegmentInterval"))
	return rv
}/* debug [instance_properties/getter]: preferredOutputSegmentInterval */


// The interval of output segments that you prefer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/preferredOutputSegmentInterval
func (a_ AssetWriter) SetPreferredOutputSegmentInterval(value objc.IObject /* cross-framework: Time */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredOutputSegmentInterval:"), value)
}/* debug [instance_properties/setter]: preferredOutputSegmentInterval */


// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/producesCombinableFragments
func (a_ AssetWriter) ProducesCombinableFragments() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("producesCombinableFragments"))
	return rv
}/* debug [instance_properties/getter]: producesCombinableFragments */


// A Boolean value that indicates whether the asset writer outputs movie fragments suitable for combining with others.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/producesCombinableFragments
func (a_ AssetWriter) SetProducesCombinableFragments(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProducesCombinableFragments:"), value)
}/* debug [instance_properties/setter]: producesCombinableFragments */


// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/shouldOptimizeForNetworkUse
func (a_ AssetWriter) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}/* debug [instance_properties/getter]: shouldOptimizeForNetworkUse */


// A Boolean value that indicates whether to write the output file to make it more suitable for playback over a network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/shouldOptimizeForNetworkUse
func (a_ AssetWriter) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}/* debug [instance_properties/setter]: shouldOptimizeForNetworkUse */


// The status of writing samples to the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/status-swift.property
func (a_ AssetWriter) Status() AssetWriterStatus {
	rv := objc.Send[AssetWriterStatus](a_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetWriter */


