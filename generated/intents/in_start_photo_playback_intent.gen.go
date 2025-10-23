// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
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
	AlbumName() string
	SetAlbumName(value string)
	DateCreated() INDateComponentsRange
	SetDateCreated(value INDateComponentsRange)
	ExcludedAttributes() unsafe.Pointer
	SetExcludedAttributes(value unsafe.Pointer)
	IncludedAttributes() unsafe.Pointer
	SetIncludedAttributes(value unsafe.Pointer)
	LocationCreated() corelocation.Placemark
	SetLocationCreated(value corelocation.IPlacemark)
	PeopleInPhoto() INPerson
	SetPeopleInPhoto(value INPerson)
	PeopleInPhotoOperator() INConditionalOperator
	SetPeopleInPhotoOperator(value INConditionalOperator)
	SearchTerms() string
	SetSearchTerms(value string)
	SearchTermsOperator() INConditionalOperator
	SetSearchTermsOperator(value INConditionalOperator)
}

// A request to search for photos and initiate a slideshow with the results.
//
// The system creates an object when the user asks to start a slideshow of a set of photos. This intent object contains the parameters to use when searching for the photos, including the possible name of a photo album, the people in the photos, or the location of the photos. Use this intent object to perform the search and initiate the slideshow in your app. When performing the search, use only the parameters provided and ignore any that have no values. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results of the search. After a successful search, Siri launches your app so that it can begin the slideshow. For a list of other intents in the photos domain, see .


// A request to search for photos and initiate a slideshow with the results.
//
// [Full Topic]
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



// The name of the album that contains the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/albumname
func (i_ INStartPhotoPlaybackIntent) AlbumName() string {
	rv := objc.Send[string](i_.ID, objc.Sel("albumName"))
	return rv
}


// The name of the album that contains the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/albumname
func (i_ INStartPhotoPlaybackIntent) SetAlbumName(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAlbumName:"), objc.String(value))
}


// The range of dates during which someone took the pictures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/datecreated
func (i_ INStartPhotoPlaybackIntent) DateCreated() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("dateCreated"))
	return rv
}


// The range of dates during which someone took the pictures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/datecreated
func (i_ INStartPhotoPlaybackIntent) SetDateCreated(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDateCreated:"), value)
}


// The attributes that must not be present in the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/excludedattributes
func (i_ INStartPhotoPlaybackIntent) ExcludedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("excludedAttributes"))
	return rv
}


// The attributes that must not be present in the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/excludedattributes
func (i_ INStartPhotoPlaybackIntent) SetExcludedAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setExcludedAttributes:"), value)
}


// The attributes that must be present in the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/includedattributes
func (i_ INStartPhotoPlaybackIntent) IncludedAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("includedAttributes"))
	return rv
}


// The attributes that must be present in the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/includedattributes
func (i_ INStartPhotoPlaybackIntent) SetIncludedAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIncludedAttributes:"), value)
}


// The location where someone took the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/locationcreated
func (i_ INStartPhotoPlaybackIntent) LocationCreated() corelocation.Placemark {
	rv := objc.Send[corelocation.Placemark](i_.ID, objc.Sel("locationCreated"))
	return rv
}


// The location where someone took the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/locationcreated
func (i_ INStartPhotoPlaybackIntent) SetLocationCreated(value corelocation.IPlacemark) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setLocationCreated:"), value)
}


// The people in the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/peopleinphoto
func (i_ INStartPhotoPlaybackIntent) PeopleInPhoto() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("peopleInPhoto"))
	return rv
}


// The people in the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/peopleinphoto
func (i_ INStartPhotoPlaybackIntent) SetPeopleInPhoto(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPeopleInPhoto:"), value)
}


// The operator that defines how to search for people in the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/peopleinphotooperator
func (i_ INStartPhotoPlaybackIntent) PeopleInPhotoOperator() INConditionalOperator {
	rv := objc.Send[INConditionalOperator](i_.ID, objc.Sel("peopleInPhotoOperator"))
	return rv
}


// The operator that defines how to search for people in the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/peopleinphotooperator
func (i_ INStartPhotoPlaybackIntent) SetPeopleInPhotoOperator(value INConditionalOperator) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPeopleInPhotoOperator:"), value)
}


// An array of terms to look for in the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/searchterms
func (i_ INStartPhotoPlaybackIntent) SearchTerms() string {
	rv := objc.Send[string](i_.ID, objc.Sel("searchTerms"))
	return rv
}


// An array of terms to look for in the photos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/searchterms
func (i_ INStartPhotoPlaybackIntent) SetSearchTerms(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSearchTerms:"), objc.String(value))
}


// The operator that defines how to incorporate the search terms when performing the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/searchtermsoperator
func (i_ INStartPhotoPlaybackIntent) SearchTermsOperator() INConditionalOperator {
	rv := objc.Send[INConditionalOperator](i_.ID, objc.Sel("searchTermsOperator"))
	return rv
}


// The operator that defines how to incorporate the search terms when performing the search.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/instartphotoplaybackintent/searchtermsoperator
func (i_ INStartPhotoPlaybackIntent) SetSearchTermsOperator(value INConditionalOperator) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setSearchTermsOperator:"), value)
}



