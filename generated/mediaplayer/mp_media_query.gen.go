// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPMediaQuery */


/* debug [class_header]: Header for MPMediaQuery */
// The class instance for the [MediaQuery] class.
var (
	MediaQueryClass     _MediaQueryClass
	MediaQueryClassOnce sync.Once
)

func getMediaQueryClass() _MediaQueryClass {
	MediaQueryClassOnce.Do(func() {
		MediaQueryClass = _MediaQueryClass{objc.GetClass("MPMediaQuery")}
	})
	return MediaQueryClass
}

type _MediaQueryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaQuery */
// An interface definition for the [MediaQuery] class.
type IMediaQuery interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaQuery */
	// properties:
	MPMediaItemPropertyIsCompilation() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaQuery */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaQuery */
// Alloc allocates a new instance without initialization.
func (mc _MediaQueryClass) Alloc() MediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaQueryClass) New() MediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaQuery) Init() MediaQuery {
	rv := objc.Send[MediaQuery](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaQuery) Autorelease() MediaQuery {
	rv := objc.Send[MediaQuery](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaQuery creates a new MediaQuery instance.
func NewMediaQuery() MediaQuery {
	return getMediaQueryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaQuery */
// A query that specifies a set of media items from the device’s media library using a filter and a grouping type.
//
// Filter and grouping types are both optional; an unqualified query matches the entire library. A query has at most one grouping type. A query’s filter can consist of any number of media property predicates. You build filters using objects, based on property keys described in . After creating and configuring a query, you use it to retrieve media items or media item collections. You can also use a query to retrieve an array of instances, useful for displaying the results of a query in the user interface of your app. See the and properties. This class includes several convenience constructors that each apply a grouping type and, in most cases, match a subset of the library. The following table summarizes the features of these constructors. See for descriptions of the entries in the Filter column. See for descriptions of the entries in the Grouping type column.


// A query that specifies a set of media items from the device’s media library using a filter and a grouping type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery
type MediaQuery struct {
	objectivec.Object
}

// MediaQueryFrom constructs a [MediaQuery] from an unsafe.Pointer.
//
// A query that specifies a set of media items from the device’s media library using a filter and a grouping type.
func MediaQueryFrom(ptr unsafe.Pointer) MediaQuery {
	return MediaQuery{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaQuery */

// Initializes a media query with a set of media property predicates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/init(filterPredicates:)
func NewMediaQueryWithFilterPredicates(filterPredicates unsafe.Pointer) MediaQuery {
	instance := getMediaQueryClass().Alloc()
	rv := objc.Send[MediaQuery](instance.ID, objc.Sel("initWithFilterPredicates:"), filterPredicates)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMediaQueryWithFilterPredicates */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaQuery */

// Creates a media query that matches music items and that groups and sorts collections by album name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/albums()
func (mc _MediaQueryClass) AlbumsQuery() IMediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("albumsQuery"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AlbumsQuery) */


// Creates a media query that matches music items and that groups and sorts collections by artist name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/artists()
func (mc _MediaQueryClass) ArtistsQuery() IMediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("artistsQuery"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ArtistsQuery) */


// Creates a media query that matches audio book items and that groups and sorts collections by audio book name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/audiobooks()
func (mc _MediaQueryClass) AudiobooksQuery() IMediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("audiobooksQuery"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AudiobooksQuery) */


// Creates a media query that matches compilation items and that groups and sorts collections by album name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/compilations()
func (mc _MediaQueryClass) CompilationsQuery() IMediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("compilationsQuery"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CompilationsQuery) */


// Creates a media query that matches all media items and that groups and sorts collections by composer name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/composers()
func (mc _MediaQueryClass) ComposersQuery() IMediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("composersQuery"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ComposersQuery) */


// Creates a media query that matches all media items and that groups and sorts collections by genre name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/genres()
func (mc _MediaQueryClass) GenresQuery() IMediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("genresQuery"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GenresQuery) */


// Creates a media query that matches the entire library and that groups and sorts collections by playlist name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/playlists()
func (mc _MediaQueryClass) PlaylistsQuery() IMediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("playlistsQuery"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PlaylistsQuery) */


// Creates a media query that matches podcast items and that groups and sorts collections by podcast name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/podcasts()
func (mc _MediaQueryClass) PodcastsQuery() IMediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("podcastsQuery"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PodcastsQuery) */


// Creates a media query that matches music items and that groups and sorts collections by song name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/songs()
func (mc _MediaQueryClass) SongsQuery() IMediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("songsQuery"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SongsQuery) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaQuery */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaQuery */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaQuery */

// A Boolean value that indicates whether the media item is part of a compilation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaplayer/mpmediaitempropertyiscompilation
func (m_ MediaQuery) MPMediaItemPropertyIsCompilation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("MPMediaItemPropertyIsCompilation"))
	return rv
}/* debug [instance_properties/getter]: MPMediaItemPropertyIsCompilation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPMediaQuery */


