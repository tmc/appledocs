// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMovie */


/* debug [class_header]: Header for AVMovie */
// The class instance for the [Movie] class.
var (
	MovieClass     _MovieClass
	MovieClassOnce sync.Once
)

func getMovieClass() _MovieClass {
	MovieClassOnce.Do(func() {
		MovieClass = _MovieClass{objc.GetClass("AVMovie")}
	})
	return MovieClass
}

type _MovieClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Movie */
// An interface definition for the [Movie] class.
type IMovie interface {
	IAsset
	
/* debug [class_interface_properties]: Properties for Movie */
	// properties:
	CanContainMovieFragments() bool
	ContainsMovieFragments() bool
	Data() objc.IObject /* cross-framework: NSData */
	DefaultMediaDataStorage() IAVMediaDataStorage
	Tracks() []MovieTrack
	URL() objc.IObject /* cross-framework: NSURL */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Movie */
	// methods:
	IsCompatibleWithFileType(fileType FileType /* typedef */) bool
	LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer)
	LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic /* typedef */, completionHandler unsafe.Pointer)
	LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType /* typedef */, completionHandler unsafe.Pointer)
	MovieHeaderWithFileTypeError(fileType FileType /* typedef */, outError objectivec.IObject) foundation.Data
	WriteMovieHeaderToURLFileTypeOptionsError(URL objc.IObject /* cross-framework: NSURL */, fileType FileType /* typedef */, options MovieWritingOptions, outError objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Movie */
// Alloc allocates a new instance without initialization.
func (mc _MovieClass) Alloc() Movie {
	rv := objc.Send[Movie](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MovieClass) New() Movie {
	rv := objc.Send[Movie](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Movie) Init() Movie {
	rv := objc.Send[Movie](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Movie) Autorelease() Movie {
	rv := objc.Send[Movie](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMovie creates a new Movie instance.
func NewMovie() Movie {
	return getMovieClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Movie */
// An object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
//
// supports operations involving the format-specific portions of the QuickTime movie model that doesn’t support. For instance, retrieving the movie header from an existing QuickTime movie file. You can also use to write a movie header into a new file, thereby creating a reference movie.


// An object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie
type Movie struct {
	Asset
}

// MovieFrom constructs a [Movie] from an unsafe.Pointer.
//
// An object that represents an audiovisual container that conforms to the QuickTime movie file format or a related format like MPEG-4.
func MovieFrom(ptr unsafe.Pointer) Movie {
	return Movie{
		Asset: AssetFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Movie */

// Creates a movie object from a movie file’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/init(data:options:)
func NewMovieWithDataOptions(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary) Movie {
	instance := getMovieClass().Alloc()
	rv := objc.Send[Movie](instance.ID, objc.Sel("initWithData:options:"), data, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMovieWithDataOptions */


// Creates a movie object from a movie header stored in a QuickTime movie file of ISO base media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/init(url:options:)
func NewMovieWithURLOptions(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary) Movie {
	instance := getMovieClass().Alloc()
	rv := objc.Send[Movie](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMovieWithURLOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Movie */

// Returns a new movie object from a movie file’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/movieWithData:options:
func (mc _MovieClass) MovieWithDataOptions(data objc.IObject /* cross-framework: NSData */, options foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("movieWithData:options:"), data, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MovieWithDataOptions) */


// Returns a new movie object from a movie header stored in a QuickTime movie file of ISO base media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/movieWithURL:options:
func (mc _MovieClass) MovieWithURLOptions(URL objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(mc.class), objc.Sel("movieWithURL:options:"), URL, options)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MovieWithURLOptions) */


// Returns the file types that a movie supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/movieTypes()
func (mc _MovieClass) MovieTypes() []string {
	rv := objc.Send[[]string](objc.ID(mc.class), objc.Sel("movieTypes"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MovieTypes) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Movie */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Movie */

// Returns a Boolean value that indicates whether the system can create a movie header of the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/is(compatibleWithFileType:)
func (m_ Movie) IsCompatibleWithFileType(fileType FileType /* typedef */) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isCompatibleWithFileType:"), fileType)
	return rv
}/* debug [instance_methods/method]: IsCompatibleWithFileType */


// Loads a track that contains the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/loadTrack(withTrackID:completionHandler:)
func (m_ Movie) LoadTrackWithTrackIDCompletionHandler(trackID PersistentTrackID /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadTrackWithTrackID:completionHandler:"), trackID, completionHandler)
}/* debug [instance_methods/method]: LoadTrackWithTrackIDCompletionHandler */


// Loads tracks that contain media of a specified characteristic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/loadTracks(withMediaCharacteristic:completionHandler:)
func (m_ Movie) LoadTracksWithMediaCharacteristicCompletionHandler(mediaCharacteristic MediaCharacteristic /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadTracksWithMediaCharacteristic:completionHandler:"), mediaCharacteristic, completionHandler)
}/* debug [instance_methods/method]: LoadTracksWithMediaCharacteristicCompletionHandler */


// Loads tracks that contain media of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/loadTracks(withMediaType:completionHandler:)
func (m_ Movie) LoadTracksWithMediaTypeCompletionHandler(mediaType MediaType /* typedef */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("loadTracksWithMediaType:completionHandler:"), mediaType, completionHandler)
}/* debug [instance_methods/method]: LoadTracksWithMediaTypeCompletionHandler */


// Creates a header for a movie for the specified file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/makeMovieHeader(fileType:)
func (m_ Movie) MovieHeaderWithFileTypeError(fileType FileType /* typedef */, outError objectivec.IObject) foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("movieHeaderWithFileType:error:"), fileType, outError)
	return rv
}/* debug [instance_methods/method]: MovieHeaderWithFileTypeError */


// Writes the movie header to the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/writeHeader(to:fileType:options:)
func (m_ Movie) WriteMovieHeaderToURLFileTypeOptionsError(URL objc.IObject /* cross-framework: NSURL */, fileType FileType /* typedef */, options MovieWritingOptions, outError objectivec.IObject) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("writeMovieHeaderToURL:fileType:options:error:"), URL, fileType, options, outError)
	return rv
}/* debug [instance_methods/method]: WriteMovieHeaderToURLFileTypeOptionsError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Movie */

// A Boolean value that indicates whether fragments can extend the movie file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/canContainMovieFragments
func (m_ Movie) CanContainMovieFragments() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("canContainMovieFragments"))
	return rv
}/* debug [instance_properties/getter]: canContainMovieFragments */


// A Boolean value that indicates whether at least one movie fragment extends the movie file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/containsMovieFragments
func (m_ Movie) ContainsMovieFragments() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("containsMovieFragments"))
	return rv
}/* debug [instance_properties/getter]: containsMovieFragments */


// A data object that contains the movie file’s data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/data
func (m_ Movie) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The default storage container for media data added to a movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/defaultMediaDataStorage
func (m_ Movie) DefaultMediaDataStorage() IAVMediaDataStorage {
	rv := objc.Send[MediaDataStorage](m_.ID, objc.Sel("defaultMediaDataStorage"))
	return rv
}/* debug [instance_properties/getter]: defaultMediaDataStorage */


// The tracks that a movie contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/tracks
func (m_ Movie) Tracks() []MovieTrack {
	rv := objc.Send[[]MovieTrack](m_.ID, objc.Sel("tracks"))
	return rv
}/* debug [instance_properties/getter]: tracks */


// A URL to a QuickTime or ISO base media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMovie/url
func (m_ Movie) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMovie */


