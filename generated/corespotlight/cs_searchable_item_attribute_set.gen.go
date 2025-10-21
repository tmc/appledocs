// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// The class instance for the [CSSearchableItemAttributeSet] class.
var (
	CSSearchableItemAttributeSetClass     _CSSearchableItemAttributeSetClass
	CSSearchableItemAttributeSetClassOnce sync.Once
)

func getCSSearchableItemAttributeSetClass() _CSSearchableItemAttributeSetClass {
	CSSearchableItemAttributeSetClassOnce.Do(func() {
		CSSearchableItemAttributeSetClass = _CSSearchableItemAttributeSetClass{objc.GetClass("CSSearchableItemAttributeSet")}
	})
	return CSSearchableItemAttributeSetClass
}

type _CSSearchableItemAttributeSetClass struct {
	class objc.Class
}

// An interface definition for the [CSSearchableItemAttributeSet] class.
type ICSSearchableItemAttributeSet interface {
	objectivec.IObject
	MoveFrom(sourceAttributeSet unsafe.Pointer)
	SetValueForCustomKey(value objc.ID, key unsafe.Pointer)
	ValueForCustomKey(key unsafe.Pointer) objc.ID
}

// The detailed metadata for a searchable item.
//
// A contains an extensive set of attributes that describe your app’s content. Attributes include information such as its title and a brief description. They can also refer to who created the item, what kind of data it represents, when someone created it, and more. During the indexing process, you create objects and use a to fill in the attributes for that item. During a search, you can query the index for items with attributes that match specific values. When creating a , it’s important to fill out as much information in the accompanying object as possible. You don’t have to provide values for every attribute. Instead, choose attributes that match the domain of your content. This type divides attributes into groups such as media, documents, events, places, music, images, and more. You can also add custom attributes to describe new types of content. When defining custom attributes, be as specific as possible in your definition, and provide a value for the property so your custom attribute inherits from a known type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet
type CSSearchableItemAttributeSet struct {
	objectivec.Object
}

// CSSearchableItemAttributeSetFrom constructs a [CSSearchableItemAttributeSet] from an unsafe.Pointer.
//
// The detailed metadata for a searchable item.
func CSSearchableItemAttributeSetFrom(ptr unsafe.Pointer) CSSearchableItemAttributeSet {
	return CSSearchableItemAttributeSet{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CSSearchableItemAttributeSetClass) Alloc() CSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CSSearchableItemAttributeSetClass) New() CSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSSearchableItemAttributeSet) Init() CSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSSearchableItemAttributeSet) Autorelease() CSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSSearchableItemAttributeSet creates a new CSSearchableItemAttributeSet instance.
func NewCSSearchableItemAttributeSet() CSSearchableItemAttributeSet {
	return getCSSearchableItemAttributeSetClass().New()
}


// Creates an attribute set for the specified content type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/init(itemContentType:)
func NewCSSearchableItemAttributeSetWithItemContentType(itemContentType string) CSSearchableItemAttributeSet {
	instance := getCSSearchableItemAttributeSetClass().Alloc()
	rv := objc.Send[CSSearchableItemAttributeSet](instance.ID, objc.Sel("initWithItemContentType:"), objc.String(itemContentType))
	rv.Autorelease()
	return rv
}

// Creates an attribute set for the specified content type.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/init(contentType:)
func NewCSSearchableItemAttributeSetWithContentType(contentType uniformtypeidentifiers.UTType) CSSearchableItemAttributeSet {
	instance := getCSSearchableItemAttributeSetClass().Alloc()
	rv := objc.Send[CSSearchableItemAttributeSet](instance.ID, objc.Sel("initWithContentType:"), contentType)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/move(from:)
func (c_ CSSearchableItemAttributeSet) MoveFrom(sourceAttributeSet unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("moveFrom:"), sourceAttributeSet)
}

// Sets the value for a custom attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/setValue(_:forCustomKey:)
func (c_ CSSearchableItemAttributeSet) SetValueForCustomKey(value objc.ID, key unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setValue:forCustomKey:"), value, key)
}

// Returns the value associated with the specified custom attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/value(forCustomKey:)
func (c_ CSSearchableItemAttributeSet) ValueForCustomKey(key unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("valueForCustomKey:"), key)
	return rv
}

// An array of objects representing the content of the From: field in an item.
//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621608-authors
func (c_ CSSearchableItemAttributeSet) Authors() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("authors"))
	return rv
}


// SetAuthors sets the value of the authors property.
// An array of objects representing the content of the From: field in an item.

//
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621608-authors
func (c_ CSSearchableItemAttributeSet) SetAuthors(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthors:"), value)
}
// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountHandles
func (c_ CSSearchableItemAttributeSet) AccountHandles() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("accountHandles"))
	return rv
}


// SetAccountHandles sets the value of the accountHandles property.
// An array of the canonical handles for the account with which the message is associated.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountHandles
func (c_ CSSearchableItemAttributeSet) SetAccountHandles(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccountHandles:"), nsArray)
}
// The unique identifier for the account with which the message is associated, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountIdentifier
func (c_ CSSearchableItemAttributeSet) AccountIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("accountIdentifier"))
	return rv
}


// SetAccountIdentifier sets the value of the accountIdentifier property.
// The unique identifier for the account with which the message is associated, if any.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountIdentifier
func (c_ CSSearchableItemAttributeSet) SetAccountIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccountIdentifier:"), value)
}
// The manufacturer of the device that captured the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionMake
func (c_ CSSearchableItemAttributeSet) AcquisitionMake() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("acquisitionMake"))
	return rv
}


// SetAcquisitionMake sets the value of the acquisitionMake property.
// The manufacturer of the device that captured the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionMake
func (c_ CSSearchableItemAttributeSet) SetAcquisitionMake(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionMake:"), value)
}
// The model of the device that captured the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionModel
func (c_ CSSearchableItemAttributeSet) AcquisitionModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("acquisitionModel"))
	return rv
}


// SetAcquisitionModel sets the value of the acquisitionModel property.
// The model of the device that captured the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionModel
func (c_ CSSearchableItemAttributeSet) SetAcquisitionModel(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionModel:"), value)
}
// The identifiers that specify custom actions the app supports for the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/actionIdentifiers
func (c_ CSSearchableItemAttributeSet) ActionIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("actionIdentifiers"))
	return rv
}


// SetActionIdentifiers sets the value of the actionIdentifiers property.
// The identifiers that specify custom actions the app supports for the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/actionIdentifiers
func (c_ CSSearchableItemAttributeSet) SetActionIdentifiers(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setActionIdentifiers:"), nsArray)
}
// The date on which the item was moved into its current location.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/addedDate
func (c_ CSSearchableItemAttributeSet) AddedDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("addedDate"))
	return rv
}


// SetAddedDate sets the value of the addedDate property.
// The date on which the item was moved into its current location.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/addedDate
func (c_ CSSearchableItemAttributeSet) SetAddedDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAddedDate:"), value)
}
// An array of objects representing the content of the Cc: field in an email message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/additionalRecipients
func (c_ CSSearchableItemAttributeSet) AdditionalRecipients() []CSPerson {
	rv := objc.Send[[]CSPerson](c_.ID, objc.Sel("additionalRecipients"))
	return rv
}


// SetAdditionalRecipients sets the value of the additionalRecipients property.
// An array of objects representing the content of the Cc: field in an email message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/additionalRecipients
func (c_ CSSearchableItemAttributeSet) SetAdditionalRecipients(value []CSPerson) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setAdditionalRecipients:"), nsArray)
}
// The title for a collection of audio media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/album
func (c_ CSSearchableItemAttributeSet) Album() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("album"))
	return rv
}


// SetAlbum sets the value of the album property.
// The title for a collection of audio media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/album
func (c_ CSSearchableItemAttributeSet) SetAlbum(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlbum:"), value)
}
// A value that indicates if the event covers an entire day.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/allDay
func (c_ CSSearchableItemAttributeSet) AllDay() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("allDay"))
	return rv
}


// SetAllDay sets the value of the allDay property.
// A value that indicates if the event covers an entire day.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/allDay
func (c_ CSSearchableItemAttributeSet) SetAllDay(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllDay:"), value)
}
// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/alternateNames
func (c_ CSSearchableItemAttributeSet) AlternateNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("alternateNames"))
	return rv
}


// SetAlternateNames sets the value of the alternateNames property.
// An array of localized strings that represent alternate display names for the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/alternateNames
func (c_ CSSearchableItemAttributeSet) SetAlternateNames(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlternateNames:"), nsArray)
}
// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/altitude
func (c_ CSSearchableItemAttributeSet) Altitude() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("altitude"))
	return rv
}


// SetAltitude sets the value of the altitude property.
// The altitude of the item in meters above sea level, expressed using the WGS84 datum.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/altitude
func (c_ CSSearchableItemAttributeSet) SetAltitude(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAltitude:"), value)
}
// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/aperture
func (c_ CSSearchableItemAttributeSet) Aperture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("aperture"))
	return rv
}


// SetAperture sets the value of the aperture property.
// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/aperture
func (c_ CSSearchableItemAttributeSet) SetAperture(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAperture:"), value)
}
// The artist associated with the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/artist
func (c_ CSSearchableItemAttributeSet) Artist() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("artist"))
	return rv
}


// SetArtist sets the value of the artist property.
// The artist associated with the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/artist
func (c_ CSSearchableItemAttributeSet) SetArtist(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setArtist:"), value)
}
// A class of entity for which the item is intended or useful.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audiences
func (c_ CSSearchableItemAttributeSet) Audiences() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("audiences"))
	return rv
}


// SetAudiences sets the value of the audiences property.
// A class of entity for which the item is intended or useful.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audiences
func (c_ CSSearchableItemAttributeSet) SetAudiences(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudiences:"), nsArray)
}
// The audio bit rate of the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioBitRate
func (c_ CSSearchableItemAttributeSet) AudioBitRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("audioBitRate"))
	return rv
}


// SetAudioBitRate sets the value of the audioBitRate property.
// The audio bit rate of the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioBitRate
func (c_ CSSearchableItemAttributeSet) SetAudioBitRate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioBitRate:"), value)
}
// The number of channels in the audio data that the file contains.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioChannelCount
func (c_ CSSearchableItemAttributeSet) AudioChannelCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("audioChannelCount"))
	return rv
}


// SetAudioChannelCount sets the value of the audioChannelCount property.
// The number of channels in the audio data that the file contains.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioChannelCount
func (c_ CSSearchableItemAttributeSet) SetAudioChannelCount(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioChannelCount:"), value)
}
// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioEncodingApplication
func (c_ CSSearchableItemAttributeSet) AudioEncodingApplication() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("audioEncodingApplication"))
	return rv
}


// SetAudioEncodingApplication sets the value of the audioEncodingApplication property.
// The name of the application that encoded the data the audio file contains.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioEncodingApplication
func (c_ CSSearchableItemAttributeSet) SetAudioEncodingApplication(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioEncodingApplication:"), value)
}
// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioSampleRate
func (c_ CSSearchableItemAttributeSet) AudioSampleRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("audioSampleRate"))
	return rv
}


// SetAudioSampleRate sets the value of the audioSampleRate property.
// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioSampleRate
func (c_ CSSearchableItemAttributeSet) SetAudioSampleRate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSampleRate:"), value)
}
// The track number of a song or audio composition when part of an album.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioTrackNumber
func (c_ CSSearchableItemAttributeSet) AudioTrackNumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("audioTrackNumber"))
	return rv
}


// SetAudioTrackNumber sets the value of the audioTrackNumber property.
// The track number of a song or audio composition when part of an album.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioTrackNumber
func (c_ CSSearchableItemAttributeSet) SetAudioTrackNumber(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioTrackNumber:"), value)
}
// An array of addresses associated with the author of the message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorAddresses
func (c_ CSSearchableItemAttributeSet) AuthorAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("authorAddresses"))
	return rv
}


// SetAuthorAddresses sets the value of the authorAddresses property.
// An array of addresses associated with the author of the message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorAddresses
func (c_ CSSearchableItemAttributeSet) SetAuthorAddresses(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorAddresses:"), nsArray)
}
// An array of email addresses associated with the author of the message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorEmailAddresses
func (c_ CSSearchableItemAttributeSet) AuthorEmailAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("authorEmailAddresses"))
	return rv
}


// SetAuthorEmailAddresses sets the value of the authorEmailAddresses property.
// An array of email addresses associated with the author of the message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorEmailAddresses
func (c_ CSSearchableItemAttributeSet) SetAuthorEmailAddresses(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorEmailAddresses:"), nsArray)
}
// An array of names representing the authors who have worked on the message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorNames
func (c_ CSSearchableItemAttributeSet) AuthorNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("authorNames"))
	return rv
}


// SetAuthorNames sets the value of the authorNames property.
// An array of names representing the authors who have worked on the message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorNames
func (c_ CSSearchableItemAttributeSet) SetAuthorNames(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorNames:"), nsArray)
}
// The number of bits per sample.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/bitsPerSample
func (c_ CSSearchableItemAttributeSet) BitsPerSample() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("bitsPerSample"))
	return rv
}


// SetBitsPerSample sets the value of the bitsPerSample property.
// The number of bits per sample.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/bitsPerSample
func (c_ CSSearchableItemAttributeSet) SetBitsPerSample(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBitsPerSample:"), value)
}
// The owner of the camera that captured the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/cameraOwner
func (c_ CSSearchableItemAttributeSet) CameraOwner() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("cameraOwner"))
	return rv
}


// SetCameraOwner sets the value of the cameraOwner property.
// The owner of the camera that captured the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/cameraOwner
func (c_ CSSearchableItemAttributeSet) SetCameraOwner(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraOwner:"), value)
}
// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/city
func (c_ CSSearchableItemAttributeSet) City() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("city"))
	return rv
}


// SetCity sets the value of the city property.
// The city of the item’s origin according to guidelines that the provider establishes.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/city
func (c_ CSSearchableItemAttributeSet) SetCity(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), value)
}
// The codecs used to encode/decode the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/codecs
func (c_ CSSearchableItemAttributeSet) Codecs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("codecs"))
	return rv
}


// SetCodecs sets the value of the codecs property.
// The codecs used to encode/decode the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/codecs
func (c_ CSSearchableItemAttributeSet) SetCodecs(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setCodecs:"), nsArray)
}
// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/colorSpace
func (c_ CSSearchableItemAttributeSet) ColorSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("colorSpace"))
	return rv
}


// SetColorSpace sets the value of the colorSpace property.
// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/colorSpace
func (c_ CSSearchableItemAttributeSet) SetColorSpace(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSpace:"), value)
}
// A comment related to the media file.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/comment
func (c_ CSSearchableItemAttributeSet) Comment() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("comment"))
	return rv
}


// SetComment sets the value of the comment property.
// A comment related to the media file.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/comment
func (c_ CSSearchableItemAttributeSet) SetComment(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComment:"), value)
}
// The date on which the item was completed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/completionDate
func (c_ CSSearchableItemAttributeSet) CompletionDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("completionDate"))
	return rv
}


// SetCompletionDate sets the value of the completionDate property.
// The date on which the item was completed.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/completionDate
func (c_ CSSearchableItemAttributeSet) SetCompletionDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionDate:"), value)
}
// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contactKeywords
func (c_ CSSearchableItemAttributeSet) ContactKeywords() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contactKeywords"))
	return rv
}


// SetContactKeywords sets the value of the contactKeywords property.
// A list of contacts who are associated with the content in some way, not including the author.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contactKeywords
func (c_ CSSearchableItemAttributeSet) SetContactKeywords(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactKeywords:"), nsArray)
}
// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerDisplayName
func (c_ CSSearchableItemAttributeSet) ContainerDisplayName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("containerDisplayName"))
	return rv
}


// SetContainerDisplayName sets the value of the containerDisplayName property.
// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerDisplayName
func (c_ CSSearchableItemAttributeSet) SetContainerDisplayName(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerDisplayName:"), value)
}
// The identifier of the container to which the item belongs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerIdentifier
func (c_ CSSearchableItemAttributeSet) ContainerIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}


// SetContainerIdentifier sets the value of the containerIdentifier property.
// The identifier of the container to which the item belongs.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerIdentifier
func (c_ CSSearchableItemAttributeSet) SetContainerIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), value)
}
// The order of the item within the container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerOrder
func (c_ CSSearchableItemAttributeSet) ContainerOrder() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("containerOrder"))
	return rv
}


// SetContainerOrder sets the value of the containerOrder property.
// The order of the item within the container.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerOrder
func (c_ CSSearchableItemAttributeSet) SetContainerOrder(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerOrder:"), value)
}
// The title of the container to which the item belongs.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerTitle
func (c_ CSSearchableItemAttributeSet) ContainerTitle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("containerTitle"))
	return rv
}


// SetContainerTitle sets the value of the containerTitle property.
// The title of the container to which the item belongs.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerTitle
func (c_ CSSearchableItemAttributeSet) SetContainerTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerTitle:"), value)
}
// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentCreationDate
func (c_ CSSearchableItemAttributeSet) ContentCreationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("contentCreationDate"))
	return rv
}


// SetContentCreationDate sets the value of the contentCreationDate property.
// The creation date of an edited or optimized version of the song or composition.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentCreationDate
func (c_ CSSearchableItemAttributeSet) SetContentCreationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentCreationDate:"), value)
}
// A description of the item’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentDescription
func (c_ CSSearchableItemAttributeSet) ContentDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("contentDescription"))
	return rv
}


// SetContentDescription sets the value of the contentDescription property.
// A description of the item’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentDescription
func (c_ CSSearchableItemAttributeSet) SetContentDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentDescription:"), value)
}
// The date on which the contents of the file was last modified.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentModificationDate
func (c_ CSSearchableItemAttributeSet) ContentModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("contentModificationDate"))
	return rv
}


// SetContentModificationDate sets the value of the contentModificationDate property.
// The date on which the contents of the file was last modified.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentModificationDate
func (c_ CSSearchableItemAttributeSet) SetContentModificationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentModificationDate:"), value)
}
// A value that indicates if the media contains explicit content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentRating
func (c_ CSSearchableItemAttributeSet) ContentRating() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("contentRating"))
	return rv
}


// SetContentRating sets the value of the contentRating property.
// A value that indicates if the media contains explicit content.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentRating
func (c_ CSSearchableItemAttributeSet) SetContentRating(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentRating:"), value)
}
// An array of sources from which the media was obtained.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentSources
func (c_ CSSearchableItemAttributeSet) ContentSources() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contentSources"))
	return rv
}


// SetContentSources sets the value of the contentSources property.
// An array of sources from which the media was obtained.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentSources
func (c_ CSSearchableItemAttributeSet) SetContentSources(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentSources:"), nsArray)
}
// The uniform type identifier (UTI) of the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentType
func (c_ CSSearchableItemAttributeSet) ContentType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("contentType"))
	return rv
}


// SetContentType sets the value of the contentType property.
// The uniform type identifier (UTI) of the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentType
func (c_ CSSearchableItemAttributeSet) SetContentType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentType:"), value)
}
// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentTypeTree
func (c_ CSSearchableItemAttributeSet) ContentTypeTree() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contentTypeTree"))
	return rv
}


// SetContentTypeTree sets the value of the contentTypeTree property.
// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentTypeTree
func (c_ CSSearchableItemAttributeSet) SetContentTypeTree(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentTypeTree:"), nsArray)
}
// The file URL of the content to index.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentURL
func (c_ CSSearchableItemAttributeSet) ContentURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("contentURL"))
	return rv
}


// SetContentURL sets the value of the contentURL property.
// The file URL of the content to index.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentURL
func (c_ CSSearchableItemAttributeSet) SetContentURL(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentURL:"), value)
}
// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contributors
func (c_ CSSearchableItemAttributeSet) Contributors() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contributors"))
	return rv
}


// SetContributors sets the value of the contributors property.
// A list of people, organizations, or services that made contributions to the media content.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contributors
func (c_ CSSearchableItemAttributeSet) SetContributors(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setContributors:"), nsArray)
}
// The copyright date of the content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/copyright
func (c_ CSSearchableItemAttributeSet) Copyright() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("copyright"))
	return rv
}


// SetCopyright sets the value of the copyright property.
// The copyright date of the content.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/copyright
func (c_ CSSearchableItemAttributeSet) SetCopyright(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCopyright:"), value)
}
// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/country
func (c_ CSSearchableItemAttributeSet) Country() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("country"))
	return rv
}


// SetCountry sets the value of the country property.
// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/country
func (c_ CSSearchableItemAttributeSet) SetCountry(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), value)
}
// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/coverage
func (c_ CSSearchableItemAttributeSet) Coverage() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("coverage"))
	return rv
}


// SetCoverage sets the value of the coverage property.
// A list of descriptors that specify the extent or scope of the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/coverage
func (c_ CSSearchableItemAttributeSet) SetCoverage(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setCoverage:"), nsArray)
}
// The name of the app that created the content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/creator
func (c_ CSSearchableItemAttributeSet) Creator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("creator"))
	return rv
}


// SetCreator sets the value of the creator property.
// The name of the app that created the content.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/creator
func (c_ CSSearchableItemAttributeSet) SetCreator(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreator:"), value)
}
// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/darkThumbnailURL
func (c_ CSSearchableItemAttributeSet) DarkThumbnailURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("darkThumbnailURL"))
	return rv
}


// SetDarkThumbnailURL sets the value of the darkThumbnailURL property.
// The local file URL of the thumbnail image for the item when Dark Mode is active.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/darkThumbnailURL
func (c_ CSSearchableItemAttributeSet) SetDarkThumbnailURL(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDarkThumbnailURL:"), value)
}
// The delivery type of the file.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/deliveryType
func (c_ CSSearchableItemAttributeSet) DeliveryType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("deliveryType"))
	return rv
}


// SetDeliveryType sets the value of the deliveryType property.
// The delivery type of the file.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/deliveryType
func (c_ CSSearchableItemAttributeSet) SetDeliveryType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeliveryType:"), value)
}
// The name of the director of the media (for example, a movie director).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/director
func (c_ CSSearchableItemAttributeSet) Director() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("director"))
	return rv
}


// SetDirector sets the value of the director property.
// The name of the director of the media (for example, a movie director).

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/director
func (c_ CSSearchableItemAttributeSet) SetDirector(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDirector:"), value)
}
// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName
func (c_ CSSearchableItemAttributeSet) DisplayName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("displayName"))
	return rv
}


// SetDisplayName sets the value of the displayName property.
// A localized string that contains the name of the item, suitable to display in the user interface.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName
func (c_ CSSearchableItemAttributeSet) SetDisplayName(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDisplayName:"), value)
}
// An identifier that represents the domain or owner of the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/domainIdentifier
func (c_ CSSearchableItemAttributeSet) DomainIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("domainIdentifier"))
	return rv
}


// SetDomainIdentifier sets the value of the domainIdentifier property.
// An identifier that represents the domain or owner of the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/domainIdentifier
func (c_ CSSearchableItemAttributeSet) SetDomainIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDomainIdentifier:"), value)
}
// The most recent date on which the file was downloaded or received.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/downloadedDate
func (c_ CSSearchableItemAttributeSet) DownloadedDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("downloadedDate"))
	return rv
}


// SetDownloadedDate sets the value of the downloadedDate property.
// The most recent date on which the file was downloaded or received.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/downloadedDate
func (c_ CSSearchableItemAttributeSet) SetDownloadedDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDownloadedDate:"), value)
}
// The date on which the item is due.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/dueDate
func (c_ CSSearchableItemAttributeSet) DueDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("dueDate"))
	return rv
}


// SetDueDate sets the value of the dueDate property.
// The date on which the item is due.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/dueDate
func (c_ CSSearchableItemAttributeSet) SetDueDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDueDate:"), value)
}
// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/duration
func (c_ CSSearchableItemAttributeSet) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
// The duration (if appropriate) of the content of the file, in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/duration
func (c_ CSSearchableItemAttributeSet) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDuration:"), value)
}
// A list of editors who have worked on the file.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/editors
func (c_ CSSearchableItemAttributeSet) Editors() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("editors"))
	return rv
}


// SetEditors sets the value of the editors property.
// A list of editors who have worked on the file.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/editors
func (c_ CSSearchableItemAttributeSet) SetEditors(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setEditors:"), nsArray)
}
// An array of email addresses associated with the message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailAddresses
func (c_ CSSearchableItemAttributeSet) EmailAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("emailAddresses"))
	return rv
}


// SetEmailAddresses sets the value of the emailAddresses property.
// An array of email addresses associated with the message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailAddresses
func (c_ CSSearchableItemAttributeSet) SetEmailAddresses(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailAddresses:"), nsArray)
}
// A dictionary that contains all the headers of the message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailHeaders
func (c_ CSSearchableItemAttributeSet) EmailHeaders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("emailHeaders"))
	return rv
}


// SetEmailHeaders sets the value of the emailHeaders property.
// A dictionary that contains all the headers of the message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailHeaders
func (c_ CSSearchableItemAttributeSet) SetEmailHeaders(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailHeaders:"), value)
}
// The name of the apps that converted the original content into a PDF stream.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/encodingApplications
func (c_ CSSearchableItemAttributeSet) EncodingApplications() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("encodingApplications"))
	return rv
}


// SetEncodingApplications sets the value of the encodingApplications property.
// The name of the apps that converted the original content into a PDF stream.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/encodingApplications
func (c_ CSSearchableItemAttributeSet) SetEncodingApplications(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setEncodingApplications:"), nsArray)
}
// The end date for the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/endDate
func (c_ CSSearchableItemAttributeSet) EndDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("endDate"))
	return rv
}


// SetEndDate sets the value of the endDate property.
// The end date for the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/endDate
func (c_ CSSearchableItemAttributeSet) SetEndDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEndDate:"), value)
}
// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifVersion
func (c_ CSSearchableItemAttributeSet) EXIFVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("EXIFVersion"))
	return rv
}


// SetEXIFVersion sets the value of the EXIFVersion property.
// The version of the EXIF header that was used to generate the metadata for the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifVersion
func (c_ CSSearchableItemAttributeSet) SetEXIFVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEXIFVersion:"), value)
}
// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifgpsVersion
func (c_ CSSearchableItemAttributeSet) EXIFGPSVersion() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("EXIFGPSVersion"))
	return rv
}


// SetEXIFGPSVersion sets the value of the EXIFGPSVersion property.
// The version of GPS Info IFD header that was used to generate the metadata for the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifgpsVersion
func (c_ CSSearchableItemAttributeSet) SetEXIFGPSVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEXIFGPSVersion:"), value)
}
// The mode the camera used for the exposure of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureMode
func (c_ CSSearchableItemAttributeSet) ExposureMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("exposureMode"))
	return rv
}


// SetExposureMode sets the value of the exposureMode property.
// The mode the camera used for the exposure of the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureMode
func (c_ CSSearchableItemAttributeSet) SetExposureMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureMode:"), value)
}
// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureProgram
func (c_ CSSearchableItemAttributeSet) ExposureProgram() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("exposureProgram"))
	return rv
}


// SetExposureProgram sets the value of the exposureProgram property.
// The class of the program the camera used to set exposure when capturing the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureProgram
func (c_ CSSearchableItemAttributeSet) SetExposureProgram(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureProgram:"), value)
}
// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTime
func (c_ CSSearchableItemAttributeSet) ExposureTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("exposureTime"))
	return rv
}


// SetExposureTime sets the value of the exposureTime property.
// The time that the lens was open during exposure, in seconds.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTime
func (c_ CSSearchableItemAttributeSet) SetExposureTime(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTime:"), value)
}
// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTimeString
func (c_ CSSearchableItemAttributeSet) ExposureTimeString() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("exposureTimeString"))
	return rv
}


// SetExposureTimeString sets the value of the exposureTimeString property.
// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTimeString
func (c_ CSSearchableItemAttributeSet) SetExposureTimeString(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTimeString:"), value)
}
// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fNumber
func (c_ CSSearchableItemAttributeSet) FNumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fNumber"))
	return rv
}


// SetFNumber sets the value of the fNumber property.
// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fNumber
func (c_ CSSearchableItemAttributeSet) SetFNumber(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFNumber:"), value)
}
// The size of the document file.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fileSize
func (c_ CSSearchableItemAttributeSet) FileSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fileSize"))
	return rv
}


// SetFileSize sets the value of the fileSize property.
// The size of the document file.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fileSize
func (c_ CSSearchableItemAttributeSet) SetFileSize(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFileSize:"), value)
}
// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/flashOn
func (c_ CSSearchableItemAttributeSet) FlashOn() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("flashOn"))
	return rv
}


// SetFlashOn sets the value of the flashOn property.
// A value that indicates if the camera used a flash to capture the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/flashOn
func (c_ CSSearchableItemAttributeSet) SetFlashOn(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashOn:"), value)
}
// The actual focal length of the lens, in millimeters.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength
func (c_ CSSearchableItemAttributeSet) FocalLength() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("focalLength"))
	return rv
}


// SetFocalLength sets the value of the focalLength property.
// The actual focal length of the lens, in millimeters.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength
func (c_ CSSearchableItemAttributeSet) SetFocalLength(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength:"), value)
}
// A value that indicates if the focal length is 35mm.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength35mm
func (c_ CSSearchableItemAttributeSet) FocalLength35mm() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("focalLength35mm"))
	return rv
}


// SetFocalLength35mm sets the value of the focalLength35mm property.
// A value that indicates if the focal length is 35mm.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength35mm
func (c_ CSSearchableItemAttributeSet) SetFocalLength35mm(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength35mm:"), value)
}
// An array of font names the document uses.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fontNames
func (c_ CSSearchableItemAttributeSet) FontNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("fontNames"))
	return rv
}


// SetFontNames sets the value of the fontNames property.
// An array of font names the document uses.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fontNames
func (c_ CSSearchableItemAttributeSet) SetFontNames(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setFontNames:"), nsArray)
}
// The fully formatted address of the item, received from MapKit.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fullyFormattedAddress
func (c_ CSSearchableItemAttributeSet) FullyFormattedAddress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("fullyFormattedAddress"))
	return rv
}


// SetFullyFormattedAddress sets the value of the fullyFormattedAddress property.
// The fully formatted address of the item, received from MapKit.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fullyFormattedAddress
func (c_ CSSearchableItemAttributeSet) SetFullyFormattedAddress(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFullyFormattedAddress:"), value)
}
// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/generalMIDISequence
func (c_ CSSearchableItemAttributeSet) GeneralMIDISequence() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("generalMIDISequence"))
	return rv
}


// SetGeneralMIDISequence sets the value of the generalMIDISequence property.
// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/generalMIDISequence
func (c_ CSSearchableItemAttributeSet) SetGeneralMIDISequence(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGeneralMIDISequence:"), value)
}
// The genre of the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/genre
func (c_ CSSearchableItemAttributeSet) Genre() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("genre"))
	return rv
}


// SetGenre sets the value of the genre property.
// The genre of the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/genre
func (c_ CSSearchableItemAttributeSet) SetGenre(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGenre:"), value)
}
// Information about the GPS area.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsAreaInformation
func (c_ CSSearchableItemAttributeSet) GPSAreaInformation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSAreaInformation"))
	return rv
}


// SetGPSAreaInformation sets the value of the GPSAreaInformation property.
// Information about the GPS area.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsAreaInformation
func (c_ CSSearchableItemAttributeSet) SetGPSAreaInformation(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSAreaInformation:"), value)
}
// The date and time related to the GPS value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDateStamp
func (c_ CSSearchableItemAttributeSet) GPSDateStamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSDateStamp"))
	return rv
}


// SetGPSDateStamp sets the value of the GPSDateStamp property.
// The date and time related to the GPS value.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDateStamp
func (c_ CSSearchableItemAttributeSet) SetGPSDateStamp(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDateStamp:"), value)
}
// The bearing to the destination point.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestBearing
func (c_ CSSearchableItemAttributeSet) GPSDestBearing() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSDestBearing"))
	return rv
}


// SetGPSDestBearing sets the value of the GPSDestBearing property.
// The bearing to the destination point.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestBearing
func (c_ CSSearchableItemAttributeSet) SetGPSDestBearing(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestBearing:"), value)
}
// The distance to the destination point.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestDistance
func (c_ CSSearchableItemAttributeSet) GPSDestDistance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSDestDistance"))
	return rv
}


// SetGPSDestDistance sets the value of the GPSDestDistance property.
// The distance to the destination point.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestDistance
func (c_ CSSearchableItemAttributeSet) SetGPSDestDistance(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestDistance:"), value)
}
// The latitude of the destination point.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLatitude
func (c_ CSSearchableItemAttributeSet) GPSDestLatitude() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSDestLatitude"))
	return rv
}


// SetGPSDestLatitude sets the value of the GPSDestLatitude property.
// The latitude of the destination point.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLatitude
func (c_ CSSearchableItemAttributeSet) SetGPSDestLatitude(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestLatitude:"), value)
}
// The longitude of the destination point.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLongitude
func (c_ CSSearchableItemAttributeSet) GPSDestLongitude() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSDestLongitude"))
	return rv
}


// SetGPSDestLongitude sets the value of the GPSDestLongitude property.
// The longitude of the destination point.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLongitude
func (c_ CSSearchableItemAttributeSet) SetGPSDestLongitude(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestLongitude:"), value)
}
// The differential correction applied to the GPS receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDifferental
func (c_ CSSearchableItemAttributeSet) GPSDifferental() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSDifferental"))
	return rv
}


// SetGPSDifferental sets the value of the GPSDifferental property.
// The differential correction applied to the GPS receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDifferental
func (c_ CSSearchableItemAttributeSet) SetGPSDifferental(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDifferental:"), value)
}
// The geodetic data that the GPS receiver uses.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMapDatum
func (c_ CSSearchableItemAttributeSet) GPSMapDatum() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSMapDatum"))
	return rv
}


// SetGPSMapDatum sets the value of the GPSMapDatum property.
// The geodetic data that the GPS receiver uses.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMapDatum
func (c_ CSSearchableItemAttributeSet) SetGPSMapDatum(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSMapDatum:"), value)
}
// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMeasureMode
func (c_ CSSearchableItemAttributeSet) GPSMeasureMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSMeasureMode"))
	return rv
}


// SetGPSMeasureMode sets the value of the GPSMeasureMode property.
// The measurement precision mode in use by the GPS receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMeasureMode
func (c_ CSSearchableItemAttributeSet) SetGPSMeasureMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSMeasureMode:"), value)
}
// The location finding method that the GPS receiver uses.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsProcessingMethod
func (c_ CSSearchableItemAttributeSet) GPSProcessingMethod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSProcessingMethod"))
	return rv
}


// SetGPSProcessingMethod sets the value of the GPSProcessingMethod property.
// The location finding method that the GPS receiver uses.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsProcessingMethod
func (c_ CSSearchableItemAttributeSet) SetGPSProcessingMethod(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSProcessingMethod:"), value)
}
// The status of the GPS receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsStatus
func (c_ CSSearchableItemAttributeSet) GPSStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSStatus"))
	return rv
}


// SetGPSStatus sets the value of the GPSStatus property.
// The status of the GPS receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsStatus
func (c_ CSSearchableItemAttributeSet) SetGPSStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSStatus:"), value)
}
// The direction of travel of the item in degrees from true north.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsTrack
func (c_ CSSearchableItemAttributeSet) GPSTrack() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSTrack"))
	return rv
}


// SetGPSTrack sets the value of the GPSTrack property.
// The direction of travel of the item in degrees from true north.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsTrack
func (c_ CSSearchableItemAttributeSet) SetGPSTrack(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSTrack:"), value)
}
// The GPS dilution of precision value.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop
func (c_ CSSearchableItemAttributeSet) GPSDOP() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("GPSDOP"))
	return rv
}


// SetGPSDOP sets the value of the GPSDOP property.
// The GPS dilution of precision value.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop
func (c_ CSSearchableItemAttributeSet) SetGPSDOP(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDOP:"), value)
}
// Indicates if the image file has an alpha channel.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hasAlphaChannel
func (c_ CSSearchableItemAttributeSet) HasAlphaChannel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("hasAlphaChannel"))
	return rv
}


// SetHasAlphaChannel sets the value of the hasAlphaChannel property.
// Indicates if the image file has an alpha channel.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hasAlphaChannel
func (c_ CSSearchableItemAttributeSet) SetHasAlphaChannel(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasAlphaChannel:"), value)
}
// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/headline
func (c_ CSSearchableItemAttributeSet) Headline() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("headline"))
	return rv
}


// SetHeadline sets the value of the headline property.
// A publishable string that provides a synopsis of the contents of the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/headline
func (c_ CSSearchableItemAttributeSet) SetHeadline(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHeadline:"), value)
}
// An array of objects representing the content of the Bcc: field in an email message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hiddenAdditionalRecipients
func (c_ CSSearchableItemAttributeSet) HiddenAdditionalRecipients() []CSPerson {
	rv := objc.Send[[]CSPerson](c_.ID, objc.Sel("hiddenAdditionalRecipients"))
	return rv
}


// SetHiddenAdditionalRecipients sets the value of the hiddenAdditionalRecipients property.
// An array of objects representing the content of the Bcc: field in an email message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hiddenAdditionalRecipients
func (c_ CSSearchableItemAttributeSet) SetHiddenAdditionalRecipients(value []CSPerson) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setHiddenAdditionalRecipients:"), nsArray)
}
// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/htmlContentData
func (c_ CSSearchableItemAttributeSet) HTMLContentData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("HTMLContentData"))
	return rv
}


// SetHTMLContentData sets the value of the HTMLContentData property.
// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/htmlContentData
func (c_ CSSearchableItemAttributeSet) SetHTMLContentData(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHTMLContentData:"), value)
}
// A formal identifier that references the document the item represents.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/identifier
func (c_ CSSearchableItemAttributeSet) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
// A formal identifier that references the document the item represents.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/identifier
func (c_ CSSearchableItemAttributeSet) SetIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), value)
}
// The direction of the item’s image in degrees from true north.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/imageDirection
func (c_ CSSearchableItemAttributeSet) ImageDirection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("imageDirection"))
	return rv
}


// SetImageDirection sets the value of the imageDirection property.
// The direction of the item’s image in degrees from true north.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/imageDirection
func (c_ CSSearchableItemAttributeSet) SetImageDirection(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageDirection:"), value)
}
// An array of important dates associated with the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/importantDates
func (c_ CSSearchableItemAttributeSet) ImportantDates() []foundation.NSDate {
	rv := objc.Send[[]foundation.NSDate](c_.ID, objc.Sel("importantDates"))
	return rv
}


// SetImportantDates sets the value of the importantDates property.
// An array of important dates associated with the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/importantDates
func (c_ CSSearchableItemAttributeSet) SetImportantDates(value []foundation.NSDate) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setImportantDates:"), nsArray)
}
// Information about the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/information
func (c_ CSSearchableItemAttributeSet) Information() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("information"))
	return rv
}


// SetInformation sets the value of the information property.
// Information about the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/information
func (c_ CSSearchableItemAttributeSet) SetInformation(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInformation:"), value)
}
// An array of instant message addresses for the message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instantMessageAddresses
func (c_ CSSearchableItemAttributeSet) InstantMessageAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}


// SetInstantMessageAddresses sets the value of the instantMessageAddresses property.
// An array of instant message addresses for the message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instantMessageAddresses
func (c_ CSSearchableItemAttributeSet) SetInstantMessageAddresses(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstantMessageAddresses:"), nsArray)
}
// Instructions that concern the use of the item, such as an embargo or warning.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instructions
func (c_ CSSearchableItemAttributeSet) Instructions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("instructions"))
	return rv
}


// SetInstructions sets the value of the instructions property.
// Instructions that concern the use of the item, such as an embargo or warning.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instructions
func (c_ CSSearchableItemAttributeSet) SetInstructions(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstructions:"), value)
}
// A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isPriority
func (c_ CSSearchableItemAttributeSet) IsPriority() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("isPriority"))
	return rv
}

// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isoSpeed
func (c_ CSSearchableItemAttributeSet) ISOSpeed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("ISOSpeed"))
	return rv
}


// SetISOSpeed sets the value of the ISOSpeed property.
// The ISO speed setting at the time the camera captured the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isoSpeed
func (c_ CSSearchableItemAttributeSet) SetISOSpeed(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setISOSpeed:"), value)
}
// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keySignature
func (c_ CSSearchableItemAttributeSet) KeySignature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("keySignature"))
	return rv
}


// SetKeySignature sets the value of the keySignature property.
// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keySignature
func (c_ CSSearchableItemAttributeSet) SetKeySignature(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeySignature:"), value)
}
// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keywords
func (c_ CSSearchableItemAttributeSet) Keywords() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("keywords"))
	return rv
}


// SetKeywords sets the value of the keywords property.
// An array of keywords associated with the item, such as work, birthday, important, and so on.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keywords
func (c_ CSSearchableItemAttributeSet) SetKeywords(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeywords:"), nsArray)
}
// A description of the kind of document the item represents.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/kind
func (c_ CSSearchableItemAttributeSet) Kind() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("kind"))
	return rv
}


// SetKind sets the value of the kind property.
// A description of the kind of document the item represents.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/kind
func (c_ CSSearchableItemAttributeSet) SetKind(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKind:"), value)
}
// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/languages
func (c_ CSSearchableItemAttributeSet) Languages() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("languages"))
	return rv
}


// SetLanguages sets the value of the languages property.
// A list of the included languages for the intellectual content of the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/languages
func (c_ CSSearchableItemAttributeSet) SetLanguages(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguages:"), nsArray)
}
// The date on which the file was last used.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lastUsedDate
func (c_ CSSearchableItemAttributeSet) LastUsedDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("lastUsedDate"))
	return rv
}


// SetLastUsedDate sets the value of the lastUsedDate property.
// The date on which the file was last used.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lastUsedDate
func (c_ CSSearchableItemAttributeSet) SetLastUsedDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLastUsedDate:"), value)
}
// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/latitude
func (c_ CSSearchableItemAttributeSet) Latitude() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("latitude"))
	return rv
}


// SetLatitude sets the value of the latitude property.
// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/latitude
func (c_ CSSearchableItemAttributeSet) SetLatitude(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLatitude:"), value)
}
// An array that contains the names of the various layers in the file.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/layerNames
func (c_ CSSearchableItemAttributeSet) LayerNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("layerNames"))
	return rv
}


// SetLayerNames sets the value of the layerNames property.
// An array that contains the names of the various layers in the file.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/layerNames
func (c_ CSSearchableItemAttributeSet) SetLayerNames(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setLayerNames:"), nsArray)
}
// The model of the lens that captured the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lensModel
func (c_ CSSearchableItemAttributeSet) LensModel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("lensModel"))
	return rv
}


// SetLensModel sets the value of the lensModel property.
// The model of the lens that captured the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lensModel
func (c_ CSSearchableItemAttributeSet) SetLensModel(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLensModel:"), value)
}
// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/likelyJunk
func (c_ CSSearchableItemAttributeSet) LikelyJunk() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("likelyJunk"))
	return rv
}


// SetLikelyJunk sets the value of the likelyJunk property.
// A value that indicates if the message is likely to be considered junk.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/likelyJunk
func (c_ CSSearchableItemAttributeSet) SetLikelyJunk(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLikelyJunk:"), value)
}
// A value that indicates if the media is local.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/local
func (c_ CSSearchableItemAttributeSet) Local() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("local"))
	return rv
}


// SetLocal sets the value of the local property.
// A value that indicates if the media is local.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/local
func (c_ CSSearchableItemAttributeSet) SetLocal(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocal:"), value)
}
// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/longitude
func (c_ CSSearchableItemAttributeSet) Longitude() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("longitude"))
	return rv
}


// SetLongitude sets the value of the longitude property.
// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/longitude
func (c_ CSSearchableItemAttributeSet) SetLongitude(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongitude:"), value)
}
// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lyricist
func (c_ CSSearchableItemAttributeSet) Lyricist() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("lyricist"))
	return rv
}


// SetLyricist sets the value of the lyricist property.
// The lyricist or text writer for the song or audio composition that the file contains.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lyricist
func (c_ CSSearchableItemAttributeSet) SetLyricist(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLyricist:"), value)
}
// An array of mailbox identifiers associated with the message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mailboxIdentifiers
func (c_ CSSearchableItemAttributeSet) MailboxIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("mailboxIdentifiers"))
	return rv
}


// SetMailboxIdentifiers sets the value of the mailboxIdentifiers property.
// An array of mailbox identifiers associated with the message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mailboxIdentifiers
func (c_ CSSearchableItemAttributeSet) SetMailboxIdentifiers(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setMailboxIdentifiers:"), nsArray)
}
// The smallest F number of the lens.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/maxAperture
func (c_ CSSearchableItemAttributeSet) MaxAperture() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("maxAperture"))
	return rv
}


// SetMaxAperture sets the value of the maxAperture property.
// The smallest F number of the lens.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/maxAperture
func (c_ CSSearchableItemAttributeSet) SetMaxAperture(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxAperture:"), value)
}
// The media types present in the content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mediaTypes
func (c_ CSSearchableItemAttributeSet) MediaTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("mediaTypes"))
	return rv
}


// SetMediaTypes sets the value of the mediaTypes property.
// The media types present in the content.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mediaTypes
func (c_ CSSearchableItemAttributeSet) SetMediaTypes(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setMediaTypes:"), nsArray)
}
// The date on which the last metadata attribute was changed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/metadataModificationDate
func (c_ CSSearchableItemAttributeSet) MetadataModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("metadataModificationDate"))
	return rv
}


// SetMetadataModificationDate sets the value of the metadataModificationDate property.
// The date on which the last metadata attribute was changed.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/metadataModificationDate
func (c_ CSSearchableItemAttributeSet) SetMetadataModificationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataModificationDate:"), value)
}
// The metering mode.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/meteringMode
func (c_ CSSearchableItemAttributeSet) MeteringMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("meteringMode"))
	return rv
}


// SetMeteringMode sets the value of the meteringMode property.
// The metering mode.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/meteringMode
func (c_ CSSearchableItemAttributeSet) SetMeteringMode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMeteringMode:"), value)
}
// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalGenre
func (c_ CSSearchableItemAttributeSet) MusicalGenre() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("musicalGenre"))
	return rv
}


// SetMusicalGenre sets the value of the musicalGenre property.
// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalGenre
func (c_ CSSearchableItemAttributeSet) SetMusicalGenre(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalGenre:"), value)
}
// The category of the instrument associated with the audio file.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentCategory
func (c_ CSSearchableItemAttributeSet) MusicalInstrumentCategory() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("musicalInstrumentCategory"))
	return rv
}


// SetMusicalInstrumentCategory sets the value of the musicalInstrumentCategory property.
// The category of the instrument associated with the audio file.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentCategory
func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentCategory(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentCategory:"), value)
}
// The name of an instrument within the context of an instrument category.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentName
func (c_ CSSearchableItemAttributeSet) MusicalInstrumentName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("musicalInstrumentName"))
	return rv
}


// SetMusicalInstrumentName sets the value of the musicalInstrumentName property.
// The name of an instrument within the context of an instrument category.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentName
func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentName(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentName:"), value)
}
// The name of the location or point of interest associated with the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/namedLocation
func (c_ CSSearchableItemAttributeSet) NamedLocation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("namedLocation"))
	return rv
}


// SetNamedLocation sets the value of the namedLocation property.
// The name of the location or point of interest associated with the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/namedLocation
func (c_ CSSearchableItemAttributeSet) SetNamedLocation(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNamedLocation:"), value)
}
// A list of companies or organizations that created the content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/organizations
func (c_ CSSearchableItemAttributeSet) Organizations() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("organizations"))
	return rv
}


// SetOrganizations sets the value of the organizations property.
// A list of companies or organizations that created the content.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/organizations
func (c_ CSSearchableItemAttributeSet) SetOrganizations(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrganizations:"), nsArray)
}
// The orientation of the data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/orientation
func (c_ CSSearchableItemAttributeSet) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("orientation"))
	return rv
}


// SetOrientation sets the value of the orientation property.
// The orientation of the data.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/orientation
func (c_ CSSearchableItemAttributeSet) SetOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrientation:"), value)
}
// The original format of the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalFormat
func (c_ CSSearchableItemAttributeSet) OriginalFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("originalFormat"))
	return rv
}


// SetOriginalFormat sets the value of the originalFormat property.
// The original format of the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalFormat
func (c_ CSSearchableItemAttributeSet) SetOriginalFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalFormat:"), value)
}
// The original source of the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalSource
func (c_ CSSearchableItemAttributeSet) OriginalSource() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("originalSource"))
	return rv
}


// SetOriginalSource sets the value of the originalSource property.
// The original source of the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalSource
func (c_ CSSearchableItemAttributeSet) SetOriginalSource(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalSource:"), value)
}
// The number of pages in the document.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageCount
func (c_ CSSearchableItemAttributeSet) PageCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("pageCount"))
	return rv
}


// SetPageCount sets the value of the pageCount property.
// The number of pages in the document.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageCount
func (c_ CSSearchableItemAttributeSet) SetPageCount(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageCount:"), value)
}
// The height of the document page, in points (72 points per inch).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageHeight
func (c_ CSSearchableItemAttributeSet) PageHeight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("pageHeight"))
	return rv
}


// SetPageHeight sets the value of the pageHeight property.
// The height of the document page, in points (72 points per inch).

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageHeight
func (c_ CSSearchableItemAttributeSet) SetPageHeight(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageHeight:"), value)
}
// The width of the document page, in points (72 points per inch).
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageWidth
func (c_ CSSearchableItemAttributeSet) PageWidth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("pageWidth"))
	return rv
}


// SetPageWidth sets the value of the pageWidth property.
// The width of the document page, in points (72 points per inch).

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageWidth
func (c_ CSSearchableItemAttributeSet) SetPageWidth(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageWidth:"), value)
}
// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/participants
func (c_ CSSearchableItemAttributeSet) Participants() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("participants"))
	return rv
}


// SetParticipants sets the value of the participants property.
// A list of people who are visible in an image or movie or written about in a document.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/participants
func (c_ CSSearchableItemAttributeSet) SetParticipants(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setParticipants:"), nsArray)
}
// The complete path to the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/path
func (c_ CSSearchableItemAttributeSet) Path() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("path"))
	return rv
}


// SetPath sets the value of the path property.
// The complete path to the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/path
func (c_ CSSearchableItemAttributeSet) SetPath(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPath:"), value)
}
// A list of performers in the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/performers
func (c_ CSSearchableItemAttributeSet) Performers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("performers"))
	return rv
}


// SetPerformers sets the value of the performers property.
// A list of performers in the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/performers
func (c_ CSSearchableItemAttributeSet) SetPerformers(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerformers:"), nsArray)
}
// An array of phone numbers associated with the message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/phoneNumbers
func (c_ CSSearchableItemAttributeSet) PhoneNumbers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}


// SetPhoneNumbers sets the value of the phoneNumbers property.
// An array of phone numbers associated with the message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/phoneNumbers
func (c_ CSSearchableItemAttributeSet) SetPhoneNumbers(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneNumbers:"), nsArray)
}
// The total number of pixels in the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelCount
func (c_ CSSearchableItemAttributeSet) PixelCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("pixelCount"))
	return rv
}


// SetPixelCount sets the value of the pixelCount property.
// The total number of pixels in the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelCount
func (c_ CSSearchableItemAttributeSet) SetPixelCount(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelCount:"), value)
}
// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelHeight
func (c_ CSSearchableItemAttributeSet) PixelHeight() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("pixelHeight"))
	return rv
}


// SetPixelHeight sets the value of the pixelHeight property.
// The height of the item, such as image or video frame height, in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelHeight
func (c_ CSSearchableItemAttributeSet) SetPixelHeight(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelHeight:"), value)
}
// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelWidth
func (c_ CSSearchableItemAttributeSet) PixelWidth() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("pixelWidth"))
	return rv
}


// SetPixelWidth sets the value of the pixelWidth property.
// The width of the item, such as image or video frame width, in pixels.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelWidth
func (c_ CSSearchableItemAttributeSet) SetPixelWidth(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelWidth:"), value)
}
// A user-supplied play count for the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/playCount
func (c_ CSSearchableItemAttributeSet) PlayCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("playCount"))
	return rv
}


// SetPlayCount sets the value of the playCount property.
// A user-supplied play count for the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/playCount
func (c_ CSSearchableItemAttributeSet) SetPlayCount(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayCount:"), value)
}
// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/postalCode
func (c_ CSSearchableItemAttributeSet) PostalCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("postalCode"))
	return rv
}


// SetPostalCode sets the value of the postalCode property.
// The postal code for the item according to guidelines the provider establishes.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/postalCode
func (c_ CSSearchableItemAttributeSet) SetPostalCode(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalCode:"), value)
}
// An array of objects representing the content of the To: field in an email message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/primaryRecipients
func (c_ CSSearchableItemAttributeSet) PrimaryRecipients() []CSPerson {
	rv := objc.Send[[]CSPerson](c_.ID, objc.Sel("primaryRecipients"))
	return rv
}


// SetPrimaryRecipients sets the value of the primaryRecipients property.
// An array of objects representing the content of the To: field in an email message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/primaryRecipients
func (c_ CSSearchableItemAttributeSet) SetPrimaryRecipients(value []CSPerson) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryRecipients:"), nsArray)
}
// The producer of the content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/producer
func (c_ CSSearchableItemAttributeSet) Producer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("producer"))
	return rv
}


// SetProducer sets the value of the producer property.
// The producer of the content.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/producer
func (c_ CSSearchableItemAttributeSet) SetProducer(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProducer:"), value)
}
// The name of the color profile the camera used for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/profileName
func (c_ CSSearchableItemAttributeSet) ProfileName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("profileName"))
	return rv
}


// SetProfileName sets the value of the profileName property.
// The name of the color profile the camera used for the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/profileName
func (c_ CSSearchableItemAttributeSet) SetProfileName(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProfileName:"), value)
}
// A list of projects of which this file is a part.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/projects
func (c_ CSSearchableItemAttributeSet) Projects() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("projects"))
	return rv
}


// SetProjects sets the value of the projects property.
// A list of projects of which this file is a part.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/projects
func (c_ CSSearchableItemAttributeSet) SetProjects(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setProjects:"), nsArray)
}
// An array of identifiers that corresponds to data representations the delegate provides.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerDataTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) ProviderDataTypeIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("providerDataTypeIdentifiers"))
	return rv
}


// SetProviderDataTypeIdentifiers sets the value of the providerDataTypeIdentifiers property.
// An array of identifiers that corresponds to data representations the delegate provides.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerDataTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderDataTypeIdentifiers(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderDataTypeIdentifiers:"), nsArray)
}
// An array of identifiers that corresponds to file representations the delegate provides.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerFileTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) ProviderFileTypeIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("providerFileTypeIdentifiers"))
	return rv
}


// SetProviderFileTypeIdentifiers sets the value of the providerFileTypeIdentifiers property.
// An array of identifiers that corresponds to file representations the delegate provides.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerFileTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderFileTypeIdentifiers(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderFileTypeIdentifiers:"), nsArray)
}
// An array of identifiers that corresponds to in-place file representations the delegate provides.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerInPlaceFileTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) ProviderInPlaceFileTypeIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("providerInPlaceFileTypeIdentifiers"))
	return rv
}


// SetProviderInPlaceFileTypeIdentifiers sets the value of the providerInPlaceFileTypeIdentifiers property.
// An array of identifiers that corresponds to in-place file representations the delegate provides.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerInPlaceFileTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderInPlaceFileTypeIdentifiers(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderInPlaceFileTypeIdentifiers:"), nsArray)
}
// A list of people, organizations, services, or other entities responsible for making the media available.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/publishers
func (c_ CSSearchableItemAttributeSet) Publishers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("publishers"))
	return rv
}


// SetPublishers sets the value of the publishers property.
// A list of people, organizations, services, or other entities responsible for making the media available.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/publishers
func (c_ CSSearchableItemAttributeSet) SetPublishers(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setPublishers:"), nsArray)
}
// A number that indicates the relative importance of the item among other items from the app.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rankingHint
func (c_ CSSearchableItemAttributeSet) RankingHint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("rankingHint"))
	return rv
}


// SetRankingHint sets the value of the rankingHint property.
// A number that indicates the relative importance of the item among other items from the app.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rankingHint
func (c_ CSSearchableItemAttributeSet) SetRankingHint(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRankingHint:"), value)
}
// The user-supplied rating of the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rating
func (c_ CSSearchableItemAttributeSet) Rating() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("rating"))
	return rv
}


// SetRating sets the value of the rating property.
// The user-supplied rating of the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rating
func (c_ CSSearchableItemAttributeSet) SetRating(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRating:"), value)
}
// A description of the rating.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/ratingDescription
func (c_ CSSearchableItemAttributeSet) RatingDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("ratingDescription"))
	return rv
}


// SetRatingDescription sets the value of the ratingDescription property.
// A description of the rating.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/ratingDescription
func (c_ CSSearchableItemAttributeSet) SetRatingDescription(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRatingDescription:"), value)
}
// An array of addresses associated with the recipients of the message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientAddresses
func (c_ CSSearchableItemAttributeSet) RecipientAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("recipientAddresses"))
	return rv
}


// SetRecipientAddresses sets the value of the recipientAddresses property.
// An array of addresses associated with the recipients of the message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientAddresses
func (c_ CSSearchableItemAttributeSet) SetRecipientAddresses(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientAddresses:"), nsArray)
}
// An array of email addresses associated with the recipient.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientEmailAddresses
func (c_ CSSearchableItemAttributeSet) RecipientEmailAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("recipientEmailAddresses"))
	return rv
}


// SetRecipientEmailAddresses sets the value of the recipientEmailAddresses property.
// An array of email addresses associated with the recipient.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientEmailAddresses
func (c_ CSSearchableItemAttributeSet) SetRecipientEmailAddresses(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientEmailAddresses:"), nsArray)
}
// An array of names representing the recipients of this message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientNames
func (c_ CSSearchableItemAttributeSet) RecipientNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("recipientNames"))
	return rv
}


// SetRecipientNames sets the value of the recipientNames property.
// An array of names representing the recipients of this message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientNames
func (c_ CSSearchableItemAttributeSet) SetRecipientNames(value []string) {
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
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientNames:"), nsArray)
}
// The recording date of the song or composition.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recordingDate
func (c_ CSSearchableItemAttributeSet) RecordingDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("recordingDate"))
	return rv
}


// SetRecordingDate sets the value of the recordingDate property.
// The recording date of the song or composition.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recordingDate
func (c_ CSSearchableItemAttributeSet) SetRecordingDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordingDate:"), value)
}
// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/redEyeOn
func (c_ CSSearchableItemAttributeSet) RedEyeOn() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("redEyeOn"))
	return rv
}


// SetRedEyeOn sets the value of the redEyeOn property.
// A value that indicates if the camera used red-eye reduction when capturing the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/redEyeOn
func (c_ CSSearchableItemAttributeSet) SetRedEyeOn(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRedEyeOn:"), value)
}
// The unique identifier for the item to which the activity is related.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/relatedUniqueIdentifier
func (c_ CSSearchableItemAttributeSet) RelatedUniqueIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("relatedUniqueIdentifier"))
	return rv
}


// SetRelatedUniqueIdentifier sets the value of the relatedUniqueIdentifier property.
// The unique identifier for the item to which the activity is related.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/relatedUniqueIdentifier
func (c_ CSSearchableItemAttributeSet) SetRelatedUniqueIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRelatedUniqueIdentifier:"), value)
}
// The resolution height of the image, in DPI.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionHeightDPI
func (c_ CSSearchableItemAttributeSet) ResolutionHeightDPI() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("resolutionHeightDPI"))
	return rv
}


// SetResolutionHeightDPI sets the value of the resolutionHeightDPI property.
// The resolution height of the image, in DPI.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionHeightDPI
func (c_ CSSearchableItemAttributeSet) SetResolutionHeightDPI(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionHeightDPI:"), value)
}
// The resolution width of the image, in DPI.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionWidthDPI
func (c_ CSSearchableItemAttributeSet) ResolutionWidthDPI() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("resolutionWidthDPI"))
	return rv
}


// SetResolutionWidthDPI sets the value of the resolutionWidthDPI property.
// The resolution width of the image, in DPI.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionWidthDPI
func (c_ CSSearchableItemAttributeSet) SetResolutionWidthDPI(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionWidthDPI:"), value)
}
// A link to information about the rights held in and over the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rights
func (c_ CSSearchableItemAttributeSet) Rights() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("rights"))
	return rv
}


// SetRights sets the value of the rights property.
// A link to information about the rights held in and over the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rights
func (c_ CSSearchableItemAttributeSet) SetRights(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRights:"), value)
}
// Indicates the role of the content creator.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/role
func (c_ CSSearchableItemAttributeSet) Role() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("role"))
	return rv
}


// SetRole sets the value of the role property.
// Indicates the role of the content creator.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/role
func (c_ CSSearchableItemAttributeSet) SetRole(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRole:"), value)
}
// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/securityMethod
func (c_ CSSearchableItemAttributeSet) SecurityMethod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("securityMethod"))
	return rv
}


// SetSecurityMethod sets the value of the securityMethod property.
// The security method (a type of encryption) that protects the document file.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/securityMethod
func (c_ CSSearchableItemAttributeSet) SetSecurityMethod(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecurityMethod:"), value)
}
// The file type of the item to enable the user to share items from Spotlight.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/sharedItemContentType
func (c_ CSSearchableItemAttributeSet) SharedItemContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](c_.ID, objc.Sel("sharedItemContentType"))
	return rv
}


// SetSharedItemContentType sets the value of the sharedItemContentType property.
// The file type of the item to enable the user to share items from Spotlight.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/sharedItemContentType
func (c_ CSSearchableItemAttributeSet) SetSharedItemContentType(value uniformtypeidentifiers.UTType) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSharedItemContentType:"), value)
}
// The speed of the item, in kilometers per hour.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/speed
func (c_ CSSearchableItemAttributeSet) Speed() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("speed"))
	return rv
}


// SetSpeed sets the value of the speed property.
// The speed of the item, in kilometers per hour.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/speed
func (c_ CSSearchableItemAttributeSet) SetSpeed(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpeed:"), value)
}
// The start date for the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/startDate
func (c_ CSSearchableItemAttributeSet) StartDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("startDate"))
	return rv
}


// SetStartDate sets the value of the startDate property.
// The start date for the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/startDate
func (c_ CSSearchableItemAttributeSet) SetStartDate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStartDate:"), value)
}
// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/stateOrProvince
func (c_ CSSearchableItemAttributeSet) StateOrProvince() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("stateOrProvince"))
	return rv
}


// SetStateOrProvince sets the value of the stateOrProvince property.
// The province or state of origin according to guidelines the provider establishes.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/stateOrProvince
func (c_ CSSearchableItemAttributeSet) SetStateOrProvince(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStateOrProvince:"), value)
}
// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/streamable
func (c_ CSSearchableItemAttributeSet) Streamable() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("streamable"))
	return rv
}


// SetStreamable sets the value of the streamable property.
// A value that indicates if the content is prepared for streaming.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/streamable
func (c_ CSSearchableItemAttributeSet) SetStreamable(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreamable:"), value)
}
// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subThoroughfare
func (c_ CSSearchableItemAttributeSet) SubThoroughfare() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("subThoroughfare"))
	return rv
}


// SetSubThoroughfare sets the value of the subThoroughfare property.
// The sublocation, such as a street number, for the item according to guidelines the provider establishes.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subThoroughfare
func (c_ CSSearchableItemAttributeSet) SetSubThoroughfare(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubThoroughfare:"), value)
}
// The subject of the document.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subject
func (c_ CSSearchableItemAttributeSet) Subject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("subject"))
	return rv
}


// SetSubject sets the value of the subject property.
// The subject of the document.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subject
func (c_ CSSearchableItemAttributeSet) SetSubject(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubject:"), value)
}
// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation
func (c_ CSSearchableItemAttributeSet) SupportsNavigation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supportsNavigation"))
	return rv
}


// SetSupportsNavigation sets the value of the supportsNavigation property.
// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation
func (c_ CSSearchableItemAttributeSet) SetSupportsNavigation(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsNavigation:"), value)
}
// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsPhoneCall
func (c_ CSSearchableItemAttributeSet) SupportsPhoneCall() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("supportsPhoneCall"))
	return rv
}


// SetSupportsPhoneCall sets the value of the supportsPhoneCall property.
// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsPhoneCall
func (c_ CSSearchableItemAttributeSet) SetSupportsPhoneCall(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsPhoneCall:"), value)
}
// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/tempo
func (c_ CSSearchableItemAttributeSet) Tempo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("tempo"))
	return rv
}


// SetTempo sets the value of the tempo property.
// The tempo of the music that the audio file contains, in beats per minute.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/tempo
func (c_ CSSearchableItemAttributeSet) SetTempo(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTempo:"), value)
}
// The textual content of the message.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContent
func (c_ CSSearchableItemAttributeSet) TextContent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("textContent"))
	return rv
}


// SetTextContent sets the value of the textContent property.
// The textual content of the message.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContent
func (c_ CSSearchableItemAttributeSet) SetTextContent(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextContent:"), value)
}
// A string that presents the Apple Intelligence summarization of the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContentSummary
func (c_ CSSearchableItemAttributeSet) TextContentSummary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("textContentSummary"))
	return rv
}

// The theme of the document.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/theme
func (c_ CSSearchableItemAttributeSet) Theme() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("theme"))
	return rv
}


// SetTheme sets the value of the theme property.
// The theme of the document.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/theme
func (c_ CSSearchableItemAttributeSet) SetTheme(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTheme:"), value)
}
// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thoroughfare
func (c_ CSSearchableItemAttributeSet) Thoroughfare() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("thoroughfare"))
	return rv
}


// SetThoroughfare sets the value of the thoroughfare property.
// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thoroughfare
func (c_ CSSearchableItemAttributeSet) SetThoroughfare(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThoroughfare:"), value)
}
// Image data that represents the thumbnail of the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailData
func (c_ CSSearchableItemAttributeSet) ThumbnailData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("thumbnailData"))
	return rv
}


// SetThumbnailData sets the value of the thumbnailData property.
// Image data that represents the thumbnail of the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailData
func (c_ CSSearchableItemAttributeSet) SetThumbnailData(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailData:"), value)
}
// The local file URL of the thumbnail image for the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailURL
func (c_ CSSearchableItemAttributeSet) ThumbnailURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("thumbnailURL"))
	return rv
}


// SetThumbnailURL sets the value of the thumbnailURL property.
// The local file URL of the thumbnail image for the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailURL
func (c_ CSSearchableItemAttributeSet) SetThumbnailURL(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailURL:"), value)
}
// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timeSignature
func (c_ CSSearchableItemAttributeSet) TimeSignature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timeSignature"))
	return rv
}


// SetTimeSignature sets the value of the timeSignature property.
// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timeSignature
func (c_ CSSearchableItemAttributeSet) SetTimeSignature(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeSignature:"), value)
}
// The timestamp on the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timestamp
func (c_ CSSearchableItemAttributeSet) Timestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("timestamp"))
	return rv
}


// SetTimestamp sets the value of the timestamp property.
// The timestamp on the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timestamp
func (c_ CSSearchableItemAttributeSet) SetTimestamp(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimestamp:"), value)
}
// The title of the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title
func (c_ CSSearchableItemAttributeSet) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title
func (c_ CSSearchableItemAttributeSet) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), value)
}
// The total bit rate of the media, combining audio and video.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/totalBitRate
func (c_ CSSearchableItemAttributeSet) TotalBitRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("totalBitRate"))
	return rv
}


// SetTotalBitRate sets the value of the totalBitRate property.
// The total bit rate of the media, combining audio and video.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/totalBitRate
func (c_ CSSearchableItemAttributeSet) SetTotalBitRate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalBitRate:"), value)
}
// A string that represents the text the system transcribed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/transcribedTextContent
func (c_ CSSearchableItemAttributeSet) TranscribedTextContent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("transcribedTextContent"))
	return rv
}


// SetTranscribedTextContent sets the value of the transcribedTextContent property.
// A string that represents the text the system transcribed.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/transcribedTextContent
func (c_ CSSearchableItemAttributeSet) SetTranscribedTextContent(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTranscribedTextContent:"), value)
}
// The URL associated with the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/url
func (c_ CSSearchableItemAttributeSet) URL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("URL"))
	return rv
}


// SetURL sets the value of the URL property.
// The URL associated with the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/url
func (c_ CSSearchableItemAttributeSet) SetURL(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setURL:"), value)
}
// A value that indicates the user created the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCreated
func (c_ CSSearchableItemAttributeSet) UserCreated() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("userCreated"))
	return rv
}


// SetUserCreated sets the value of the userCreated property.
// A value that indicates the user created the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCreated
func (c_ CSSearchableItemAttributeSet) SetUserCreated(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCreated:"), value)
}
// A value that indicates the user selected the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCurated
func (c_ CSSearchableItemAttributeSet) UserCurated() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("userCurated"))
	return rv
}


// SetUserCurated sets the value of the userCurated property.
// A value that indicates the user selected the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCurated
func (c_ CSSearchableItemAttributeSet) SetUserCurated(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCurated:"), value)
}
// A value that indicates the user purchased or owns the item.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userOwned
func (c_ CSSearchableItemAttributeSet) UserOwned() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("userOwned"))
	return rv
}


// SetUserOwned sets the value of the userOwned property.
// A value that indicates the user purchased or owns the item.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userOwned
func (c_ CSSearchableItemAttributeSet) SetUserOwned(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserOwned:"), value)
}
// A version string associated with the file.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/version
func (c_ CSSearchableItemAttributeSet) Version() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("version"))
	return rv
}


// SetVersion sets the value of the version property.
// A version string associated with the file.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/version
func (c_ CSSearchableItemAttributeSet) SetVersion(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVersion:"), value)
}
// The video bit rate of the media.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/videoBitRate
func (c_ CSSearchableItemAttributeSet) VideoBitRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("videoBitRate"))
	return rv
}


// SetVideoBitRate sets the value of the videoBitRate property.
// The video bit rate of the media.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/videoBitRate
func (c_ CSSearchableItemAttributeSet) SetVideoBitRate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoBitRate:"), value)
}
// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/weakRelatedUniqueIdentifier
func (c_ CSSearchableItemAttributeSet) WeakRelatedUniqueIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("weakRelatedUniqueIdentifier"))
	return rv
}


// SetWeakRelatedUniqueIdentifier sets the value of the weakRelatedUniqueIdentifier property.
// The unique identifier for the item to which the activity is related, but not linked.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/weakRelatedUniqueIdentifier
func (c_ CSSearchableItemAttributeSet) SetWeakRelatedUniqueIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeakRelatedUniqueIdentifier:"), value)
}
// The white balance setting when the camera captured the image.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/whiteBalance
func (c_ CSSearchableItemAttributeSet) WhiteBalance() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("whiteBalance"))
	return rv
}


// SetWhiteBalance sets the value of the whiteBalance property.
// The white balance setting when the camera captured the image.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/whiteBalance
func (c_ CSSearchableItemAttributeSet) SetWhiteBalance(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteBalance:"), value)
}

