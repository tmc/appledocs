// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INStartPhotoPlaybackIntent] class.
var (
	INStartPhotoPlaybackIntentClass     _INStartPhotoPlaybackIntentClass
	INStartPhotoPlaybackIntentClassOnce sync.Once
)

func getINStartPhotoPlaybackIntentClass() _INStartPhotoPlaybackIntentClass {
	INStartPhotoPlaybackIntentClassOnce.Do(func() {
		INStartPhotoPlaybackIntentClass = _INStartPhotoPlaybackIntentClass{objc.GetClass("INStartPhotoPlaybackIntent")}
	})
	return INStartPhotoPlaybackIntentClass
}

type _INStartPhotoPlaybackIntentClass struct {
	class objc.Class
}

// An interface definition for the [INStartPhotoPlaybackIntent] class.
type IINStartPhotoPlaybackIntent interface {
	IINIntent
}

// A request to search for photos and initiate a slideshow with the results.
//
// The system creates an object when the user asks to start a slideshow of a set of photos. This intent object contains the parameters to use when searching for the photos, including the possible name of a photo album, the people in the photos, or the location of the photos. Use this intent object to perform the search and initiate the slideshow in your app. When performing the search, use only the parameters provided and ignore any that have no values. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results of the search. After a successful search, Siri launches your app so that it can begin the slideshow. For a list of other intents in the photos domain, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INStartPhotoPlaybackIntent
type INStartPhotoPlaybackIntent struct {
	INIntent
}

// INStartPhotoPlaybackIntentFrom constructs a [INStartPhotoPlaybackIntent] from an unsafe.Pointer.
//
// A request to search for photos and initiate a slideshow with the results.
func INStartPhotoPlaybackIntentFrom(ptr unsafe.Pointer) INStartPhotoPlaybackIntent {
	return INStartPhotoPlaybackIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INStartPhotoPlaybackIntentClass) Alloc() INStartPhotoPlaybackIntent {
	rv := objc.Send[INStartPhotoPlaybackIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INStartPhotoPlaybackIntentClass) New() INStartPhotoPlaybackIntent {
	rv := objc.Send[INStartPhotoPlaybackIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INStartPhotoPlaybackIntent) Init() INStartPhotoPlaybackIntent {
	rv := objc.Send[INStartPhotoPlaybackIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INStartPhotoPlaybackIntent) Autorelease() INStartPhotoPlaybackIntent {
	rv := objc.Send[INStartPhotoPlaybackIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINStartPhotoPlaybackIntent creates a new INStartPhotoPlaybackIntent instance.
func NewINStartPhotoPlaybackIntent() INStartPhotoPlaybackIntent {
	return getINStartPhotoPlaybackIntentClass().New()
}


// The people in the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/peopleinphoto
func (i_ INStartPhotoPlaybackIntent) PeopleInPhoto() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("peopleInPhoto"))
	return rv
}


// SetPeopleInPhoto sets the value of the peopleInPhoto property.
// The people in the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/peopleinphoto
func (i_ INStartPhotoPlaybackIntent) SetPeopleInPhoto(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPeopleInPhoto:"), value)
}

// The attributes that must be present in the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/includedattributes
func (i_ INStartPhotoPlaybackIntent) IncludedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("includedAttributes"))
	return rv
}


// SetIncludedAttributes sets the value of the includedAttributes property.
// The attributes that must be present in the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/includedattributes
func (i_ INStartPhotoPlaybackIntent) SetIncludedAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIncludedAttributes:"), value)
}

// The location where someone took the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/locationcreated
func (i_ INStartPhotoPlaybackIntent) LocationCreated() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("locationCreated"))
	return rv
}


// SetLocationCreated sets the value of the locationCreated property.
// The location where someone took the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/locationcreated
func (i_ INStartPhotoPlaybackIntent) SetLocationCreated(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLocationCreated:"), value)
}

// An array of terms to look for in the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/searchterms
func (i_ INStartPhotoPlaybackIntent) SearchTerms() string {
	rv := objc.Send[string](i_.ID, objc.Sel("searchTerms"))
	return rv
}


// SetSearchTerms sets the value of the searchTerms property.
// An array of terms to look for in the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/searchterms
func (i_ INStartPhotoPlaybackIntent) SetSearchTerms(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSearchTerms:"), objc.String(value))
}

// The name of the album that contains the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/albumname
func (i_ INStartPhotoPlaybackIntent) AlbumName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("albumName"))
	return rv
}


// SetAlbumName sets the value of the albumName property.
// The name of the album that contains the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/albumname
func (i_ INStartPhotoPlaybackIntent) SetAlbumName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlbumName:"), objc.String(value))
}

// The operator that defines how to search for people in the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/peopleinphotooperator
func (i_ INStartPhotoPlaybackIntent) PeopleInPhotoOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("peopleInPhotoOperator"))
	return rv
}


// SetPeopleInPhotoOperator sets the value of the peopleInPhotoOperator property.
// The operator that defines how to search for people in the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/peopleinphotooperator
func (i_ INStartPhotoPlaybackIntent) SetPeopleInPhotoOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPeopleInPhotoOperator:"), value)
}

// The operator that defines how to incorporate the search terms when performing the search.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/searchtermsoperator
func (i_ INStartPhotoPlaybackIntent) SearchTermsOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("searchTermsOperator"))
	return rv
}


// SetSearchTermsOperator sets the value of the searchTermsOperator property.
// The operator that defines how to incorporate the search terms when performing the search.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/searchtermsoperator
func (i_ INStartPhotoPlaybackIntent) SetSearchTermsOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSearchTermsOperator:"), value)
}

// The range of dates during which someone took the pictures.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/datecreated
func (i_ INStartPhotoPlaybackIntent) DateCreated() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dateCreated"))
	return rv
}


// SetDateCreated sets the value of the dateCreated property.
// The range of dates during which someone took the pictures.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/datecreated
func (i_ INStartPhotoPlaybackIntent) SetDateCreated(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDateCreated:"), value)
}

// The attributes that must not be present in the photos.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/excludedattributes
func (i_ INStartPhotoPlaybackIntent) ExcludedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("excludedAttributes"))
	return rv
}


// SetExcludedAttributes sets the value of the excludedAttributes property.
// The attributes that must not be present in the photos.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/excludedattributes
func (i_ INStartPhotoPlaybackIntent) SetExcludedAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExcludedAttributes:"), value)
}



