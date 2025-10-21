// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INMediaSearch] class.
var (
	INMediaSearchClass     _INMediaSearchClass
	INMediaSearchClassOnce sync.Once
)

func getINMediaSearchClass() _INMediaSearchClass {
	INMediaSearchClassOnce.Do(func() {
		INMediaSearchClass = _INMediaSearchClass{objc.GetClass("INMediaSearch")}
	})
	return INMediaSearchClass
}

type _INMediaSearchClass struct {
	class objc.Class
}

// An interface definition for the [INMediaSearch] class.
type IINMediaSearch interface {
	objectivec.IObject
}

// An object that describes a media type to search for, such as a station name, song name, or album name.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INMediaSearch
type INMediaSearch struct {
	objectivec.Object
}

// INMediaSearchFrom constructs a [INMediaSearch] from an unsafe.Pointer.
//
// An object that describes a media type to search for, such as a station name, song name, or album name.
func INMediaSearchFrom(ptr unsafe.Pointer) INMediaSearch {
	return INMediaSearch{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INMediaSearchClass) Alloc() INMediaSearch {
	rv := objc.Send[INMediaSearch](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INMediaSearchClass) New() INMediaSearch {
	rv := objc.Send[INMediaSearch](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INMediaSearch) Init() INMediaSearch {
	rv := objc.Send[INMediaSearch](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INMediaSearch) Autorelease() INMediaSearch {
	rv := objc.Send[INMediaSearch](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINMediaSearch creates a new INMediaSearch instance.
func NewINMediaSearch() INMediaSearch {
	return getINMediaSearchClass().New()
}


// The sort order for the found media items.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/sortorder
func (i_ INMediaSearch) SortOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("sortOrder"))
	return rv
}


// SetSortOrder sets the value of the sortOrder property.
// The sort order for the found media items.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/sortorder
func (i_ INMediaSearch) SetSortOrder(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSortOrder:"), value)
}

// The media genres to search.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/genrenames
func (i_ INMediaSearch) GenreNames() string {
	rv := objc.Send[string](i_.ID, objc.Sel("genreNames"))
	return rv
}


// SetGenreNames sets the value of the genreNames property.
// The media genres to search.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/genrenames
func (i_ INMediaSearch) SetGenreNames(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setGenreNames:"), objc.String(value))
}

// A reference for the media item to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/reference
func (i_ INMediaSearch) Reference() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("reference"))
	return rv
}


// SetReference sets the value of the reference property.
// A reference for the media item to search for.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/reference
func (i_ INMediaSearch) SetReference(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReference:"), value)
}

// The name of the media to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/medianame
func (i_ INMediaSearch) MediaName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("mediaName"))
	return rv
}


// SetMediaName sets the value of the mediaName property.
// The name of the media to search for.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/medianame
func (i_ INMediaSearch) SetMediaName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaName:"), objc.String(value))
}

// The name of the artist to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/artistname
func (i_ INMediaSearch) ArtistName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("artistName"))
	return rv
}


// SetArtistName sets the value of the artistName property.
// The name of the artist to search for.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/artistname
func (i_ INMediaSearch) SetArtistName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setArtistName:"), objc.String(value))
}

// The activity names to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/activitynames
func (i_ INMediaSearch) ActivityNames() string {
	rv := objc.Send[string](i_.ID, objc.Sel("activityNames"))
	return rv
}


// SetActivityNames sets the value of the activityNames property.
// The activity names to search for.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/activitynames
func (i_ INMediaSearch) SetActivityNames(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setActivityNames:"), objc.String(value))
}

// The moods to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/moodnames
func (i_ INMediaSearch) MoodNames() string {
	rv := objc.Send[string](i_.ID, objc.Sel("moodNames"))
	return rv
}


// SetMoodNames sets the value of the moodNames property.
// The moods to search for.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/moodnames
func (i_ INMediaSearch) SetMoodNames(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMoodNames:"), objc.String(value))
}

// The release date to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/releasedate
func (i_ INMediaSearch) ReleaseDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("releaseDate"))
	return rv
}


// SetReleaseDate sets the value of the releaseDate property.
// The release date to search for.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/releasedate
func (i_ INMediaSearch) SetReleaseDate(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setReleaseDate:"), value)
}

// The name of the album to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/albumname
func (i_ INMediaSearch) AlbumName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("albumName"))
	return rv
}


// SetAlbumName sets the value of the albumName property.
// The name of the album to search for.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/albumname
func (i_ INMediaSearch) SetAlbumName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlbumName:"), objc.String(value))
}

// The unique media identifier to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/mediaidentifier
func (i_ INMediaSearch) MediaIdentifier() string {
	rv := objc.Send[string](i_.ID, objc.Sel("mediaIdentifier"))
	return rv
}


// SetMediaIdentifier sets the value of the mediaIdentifier property.
// The unique media identifier to search for.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/mediaidentifier
func (i_ INMediaSearch) SetMediaIdentifier(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaIdentifier:"), objc.String(value))
}

// The type of media to search for.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/mediatype
func (i_ INMediaSearch) MediaType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("mediaType"))
	return rv
}


// SetMediaType sets the value of the mediaType property.
// The type of media to search for.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inmediasearch/mediatype
func (i_ INMediaSearch) SetMediaType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setMediaType:"), value)
}



