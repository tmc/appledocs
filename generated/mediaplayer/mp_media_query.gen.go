// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MediaQuery] class.
type IMediaQuery interface {
	objectivec.IObject
	AddFilterPredicate(predicate unsafe.Pointer)
	RemoveFilterPredicate(predicate unsafe.Pointer)
}

// A query that specifies a set of media items from the device’s media library using a filter and a grouping type.
//
// Filter and grouping types are both optional; an unqualified query matches the entire library. A query has at most one grouping type. A query’s filter can consist of any number of media property predicates. You build filters using objects, based on property keys described in . After creating and configuring a query, you use it to retrieve media items or media item collections. You can also use a query to retrieve an array of instances, useful for displaying the results of a query in the user interface of your app. See the and properties. This class includes several convenience constructors that each apply a grouping type and, in most cases, match a subset of the library. The following table summarizes the features of these constructors. See for descriptions of the entries in the Filter column. See for descriptions of the entries in the Grouping type column.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MediaQueryClass) Alloc() MediaQuery {
	rv := objc.Send[MediaQuery](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a media query with a set of media property predicates.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/init(filterPredicates:)
func NewMediaQueryWithFilterPredicates(filterPredicates unsafe.Pointer) MediaQuery {
	instance := getMediaQueryClass().Alloc()
	rv := objc.Send[MediaQuery](instance.ID, objc.Sel("initWithFilterPredicates:"), filterPredicates)
	rv.Autorelease()
	return rv
}


// Creates a media query that matches music items and that groups and sorts collections by album name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/albums()
func (mc _MediaQueryClass) AlbumsQuery() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("albumsQuery"))
	return rv
}

// Creates a media query that matches music items and that groups and sorts collections by artist name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/artists()
func (mc _MediaQueryClass) ArtistsQuery() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("artistsQuery"))
	return rv
}

// Creates a media query that matches audio book items and that groups and sorts collections by audio book name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/audiobooks()
func (mc _MediaQueryClass) AudiobooksQuery() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("audiobooksQuery"))
	return rv
}

// Creates a media query that matches compilation items and that groups and sorts collections by album name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/compilations()
func (mc _MediaQueryClass) CompilationsQuery() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("compilationsQuery"))
	return rv
}

// Creates a media query that matches all media items and that groups and sorts collections by composer name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/composers()
func (mc _MediaQueryClass) ComposersQuery() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("composersQuery"))
	return rv
}

// Creates a media query that matches all media items and that groups and sorts collections by genre name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/genres()
func (mc _MediaQueryClass) GenresQuery() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("genresQuery"))
	return rv
}

// Creates a media query that matches the entire library and that groups and sorts collections by playlist name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/playlists()
func (mc _MediaQueryClass) PlaylistsQuery() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("playlistsQuery"))
	return rv
}

// Creates a media query that matches podcast items and that groups and sorts collections by podcast name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/podcasts()
func (mc _MediaQueryClass) PodcastsQuery() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("podcastsQuery"))
	return rv
}

// Creates a media query that matches music items and that groups and sorts collections by song name.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/songs()
func (mc _MediaQueryClass) SongsQuery() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(mc.class), objc.Sel("songsQuery"))
	return rv
}

// Adds a media property predicate to a query.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/addFilterPredicate(_:)
func (m_ MediaQuery) AddFilterPredicate(predicate unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addFilterPredicate:"), predicate)
}

// Removes a filter predicate from a query.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/removeFilterPredicate(_:)
func (m_ MediaQuery) RemoveFilterPredicate(predicate unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeFilterPredicate:"), predicate)
}

// An array representing the section grouping of the query’s specified media item collections.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/collectionSections
func (m_ MediaQuery) CollectionSections() []MediaQuerySection {
	rv := objc.Send[[]MediaQuerySection](m_.ID, objc.Sel("collectionSections"))
	return rv
}

// An array of media item collections whose contained items match the query’s media property predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/collections
func (m_ MediaQuery) Collections() []MediaItemCollection {
	rv := objc.Send[[]MediaItemCollection](m_.ID, objc.Sel("collections"))
	return rv
}

// The media property predicates of the media query.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/filterPredicates
func (m_ MediaQuery) FilterPredicates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("filterPredicates"))
	return rv
}


// SetFilterPredicates sets the value of the filterPredicates property.
// The media property predicates of the media query.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/filterPredicates
func (m_ MediaQuery) SetFilterPredicates(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFilterPredicates:"), value)
}

// The grouping for collections retrieved with the media query.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/groupingType
func (m_ MediaQuery) GroupingType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("groupingType"))
	return rv
}


// SetGroupingType sets the value of the groupingType property.
// The grouping for collections retrieved with the media query.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/groupingType
func (m_ MediaQuery) SetGroupingType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setGroupingType:"), value)
}

// An array representing the section grouping of the query’s specified media items.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/itemSections
func (m_ MediaQuery) ItemSections() []MediaQuerySection {
	rv := objc.Send[[]MediaQuerySection](m_.ID, objc.Sel("itemSections"))
	return rv
}

// An array of media items that match the media query’s predicate.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaPlayer/MPMediaQuery/items
func (m_ MediaQuery) Items() []MediaItem {
	rv := objc.Send[[]MediaItem](m_.ID, objc.Sel("items"))
	return rv
}


