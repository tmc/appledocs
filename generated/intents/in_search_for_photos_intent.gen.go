// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchForPhotosIntent] class.
var (
	INSearchForPhotosIntentClass     _INSearchForPhotosIntentClass
	INSearchForPhotosIntentClassOnce sync.Once
)

func getINSearchForPhotosIntentClass() _INSearchForPhotosIntentClass {
	INSearchForPhotosIntentClassOnce.Do(func() {
		INSearchForPhotosIntentClass = _INSearchForPhotosIntentClass{objc.GetClass("INSearchForPhotosIntent")}
	})
	return INSearchForPhotosIntentClass
}

type _INSearchForPhotosIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSearchForPhotosIntent] class.
type IINSearchForPhotosIntent interface {
	IINIntent
}

// A request for the list of photos that match the specified criteria.
//
// The system creates an object when the user asks to search for photos in an app. The intent object contains the parameters to use during the search, including the possible name of a photo album, the people in the photos, or the location of the photos. Use this intent object to validate the search parameters and to begin the search process. When performing the search, use only the provided parameters and ignore any that have no values. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results of the search. For successful searches, Siri offers the user a way to launch your app and see the results. For a list of other intents in the photos domain, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForPhotosIntent
type INSearchForPhotosIntent struct {
	INIntent
}

// INSearchForPhotosIntentFrom constructs a [INSearchForPhotosIntent] from an unsafe.Pointer.
//
// A request for the list of photos that match the specified criteria.
func INSearchForPhotosIntentFrom(ptr unsafe.Pointer) INSearchForPhotosIntent {
	return INSearchForPhotosIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchForPhotosIntentClass) Alloc() INSearchForPhotosIntent {
	rv := objc.Send[INSearchForPhotosIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchForPhotosIntentClass) New() INSearchForPhotosIntent {
	rv := objc.Send[INSearchForPhotosIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchForPhotosIntent) Init() INSearchForPhotosIntent {
	rv := objc.Send[INSearchForPhotosIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchForPhotosIntent) Autorelease() INSearchForPhotosIntent {
	rv := objc.Send[INSearchForPhotosIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchForPhotosIntent creates a new INSearchForPhotosIntent instance.
func NewINSearchForPhotosIntent() INSearchForPhotosIntent {
	return getINSearchForPhotosIntentClass().New()
}


// The name of the album that contains the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/albumname
func (i_ INSearchForPhotosIntent) AlbumName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("albumName"))
	return rv
}


// SetAlbumName sets the value of the albumName property.
// The name of the album that contains the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/albumname
func (i_ INSearchForPhotosIntent) SetAlbumName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlbumName:"), objc.String(value))
}

// The range of dates during which someone took the pictures.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/datecreated
func (i_ INSearchForPhotosIntent) DateCreated() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dateCreated"))
	return rv
}


// SetDateCreated sets the value of the dateCreated property.
// The range of dates during which someone took the pictures.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/datecreated
func (i_ INSearchForPhotosIntent) SetDateCreated(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDateCreated:"), value)
}

// The attributes that must not be present in the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/excludedattributes
func (i_ INSearchForPhotosIntent) ExcludedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("excludedAttributes"))
	return rv
}


// SetExcludedAttributes sets the value of the excludedAttributes property.
// The attributes that must not be present in the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/excludedattributes
func (i_ INSearchForPhotosIntent) SetExcludedAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExcludedAttributes:"), value)
}

// The attributes that must be present in the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/includedattributes
func (i_ INSearchForPhotosIntent) IncludedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("includedAttributes"))
	return rv
}


// SetIncludedAttributes sets the value of the includedAttributes property.
// The attributes that must be present in the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/includedattributes
func (i_ INSearchForPhotosIntent) SetIncludedAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIncludedAttributes:"), value)
}

// The location where someone took the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/locationcreated
func (i_ INSearchForPhotosIntent) LocationCreated() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("locationCreated"))
	return rv
}


// SetLocationCreated sets the value of the locationCreated property.
// The location where someone took the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/locationcreated
func (i_ INSearchForPhotosIntent) SetLocationCreated(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLocationCreated:"), value)
}

// The people identified in the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/peopleinphoto
func (i_ INSearchForPhotosIntent) PeopleInPhoto() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("peopleInPhoto"))
	return rv
}


// SetPeopleInPhoto sets the value of the peopleInPhoto property.
// The people identified in the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/peopleinphoto
func (i_ INSearchForPhotosIntent) SetPeopleInPhoto(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPeopleInPhoto:"), value)
}

// The operator that defines how to search for people in the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/peopleinphotooperator
func (i_ INSearchForPhotosIntent) PeopleInPhotoOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("peopleInPhotoOperator"))
	return rv
}


// SetPeopleInPhotoOperator sets the value of the peopleInPhotoOperator property.
// The operator that defines how to search for people in the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/peopleinphotooperator
func (i_ INSearchForPhotosIntent) SetPeopleInPhotoOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPeopleInPhotoOperator:"), value)
}

// An array of terms to look for in the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/searchterms
func (i_ INSearchForPhotosIntent) SearchTerms() string {
	rv := objc.Send[string](i_.ID, objc.Sel("searchTerms"))
	return rv
}


// SetSearchTerms sets the value of the searchTerms property.
// An array of terms to look for in the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/searchterms
func (i_ INSearchForPhotosIntent) SetSearchTerms(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSearchTerms:"), objc.String(value))
}

// The operator that defines how to incorporate the search terms when performing the search.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/searchtermsoperator
func (i_ INSearchForPhotosIntent) SearchTermsOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("searchTermsOperator"))
	return rv
}


// SetSearchTermsOperator sets the value of the searchTermsOperator property.
// The operator that defines how to incorporate the search terms when performing the search.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforphotosintent/searchtermsoperator
func (i_ INSearchForPhotosIntent) SetSearchTermsOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSearchTermsOperator:"), value)
}



