// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	MoveFrom(sourceAttributeSet ICSSearchableItemAttributeSet)
	SetValueForCustomKey(value objectivec.IObject, key CSCustomAttributeKey)
	ValueForCustomKey(key CSCustomAttributeKey) objc.ID
	Authors() CSPerson
	SetAuthors(value ICSPerson)
	AccountHandles() []string
	SetAccountHandles(value []string)
	AccountIdentifier() string
	SetAccountIdentifier(value string)
	AcquisitionMake() string
	SetAcquisitionMake(value string)
	AcquisitionModel() string
	SetAcquisitionModel(value string)
	ActionIdentifiers() []string
	SetActionIdentifiers(value []string)
	AddedDate() foundation.NSDate
	SetAddedDate(value foundation.IDate)
	AdditionalRecipients() []CSPerson
	SetAdditionalRecipients(value []CSPerson)
	Album() string
	SetAlbum(value string)
	AllDay() foundation.Number
	SetAllDay(value foundation.INumber)
	AlternateNames() []string
	SetAlternateNames(value []string)
	Altitude() foundation.Number
	SetAltitude(value foundation.INumber)
	Aperture() foundation.Number
	SetAperture(value foundation.INumber)
	Artist() string
	SetArtist(value string)
	Audiences() []string
	SetAudiences(value []string)
	AudioBitRate() foundation.Number
	SetAudioBitRate(value foundation.INumber)
	AudioChannelCount() foundation.Number
	SetAudioChannelCount(value foundation.INumber)
	AudioEncodingApplication() string
	SetAudioEncodingApplication(value string)
	AudioSampleRate() foundation.Number
	SetAudioSampleRate(value foundation.INumber)
	AudioTrackNumber() foundation.Number
	SetAudioTrackNumber(value foundation.INumber)
	AuthorAddresses() []string
	SetAuthorAddresses(value []string)
	AuthorEmailAddresses() []string
	SetAuthorEmailAddresses(value []string)
	AuthorNames() []string
	SetAuthorNames(value []string)
	BitsPerSample() foundation.Number
	SetBitsPerSample(value foundation.INumber)
	CameraOwner() string
	SetCameraOwner(value string)
	City() string
	SetCity(value string)
	Codecs() []string
	SetCodecs(value []string)
	ColorSpace() string
	SetColorSpace(value string)
	Comment() string
	SetComment(value string)
	CompletionDate() foundation.NSDate
	SetCompletionDate(value foundation.IDate)
	ContactKeywords() []string
	SetContactKeywords(value []string)
	ContainerDisplayName() string
	SetContainerDisplayName(value string)
	ContainerIdentifier() string
	SetContainerIdentifier(value string)
	ContainerOrder() foundation.Number
	SetContainerOrder(value foundation.INumber)
	ContainerTitle() string
	SetContainerTitle(value string)
	ContentCreationDate() foundation.NSDate
	SetContentCreationDate(value foundation.IDate)
	ContentDescription() string
	SetContentDescription(value string)
	ContentModificationDate() foundation.NSDate
	SetContentModificationDate(value foundation.IDate)
	ContentRating() foundation.Number
	SetContentRating(value foundation.INumber)
	ContentSources() []string
	SetContentSources(value []string)
	ContentType() string
	SetContentType(value string)
	ContentTypeTree() []string
	SetContentTypeTree(value []string)
	ContentURL() foundation.URL
	SetContentURL(value foundation.IURL)
	Contributors() []string
	SetContributors(value []string)
	Copyright() string
	SetCopyright(value string)
	Country() string
	SetCountry(value string)
	Coverage() []string
	SetCoverage(value []string)
	Creator() string
	SetCreator(value string)
	DarkThumbnailURL() foundation.URL
	SetDarkThumbnailURL(value foundation.IURL)
	DeliveryType() foundation.Number
	SetDeliveryType(value foundation.INumber)
	Director() string
	SetDirector(value string)
	DisplayName() string
	SetDisplayName(value string)
	DomainIdentifier() string
	SetDomainIdentifier(value string)
	DownloadedDate() foundation.NSDate
	SetDownloadedDate(value foundation.IDate)
	DueDate() foundation.NSDate
	SetDueDate(value foundation.IDate)
	Duration() foundation.Number
	SetDuration(value foundation.INumber)
	Editors() []string
	SetEditors(value []string)
	EmailAddresses() []string
	SetEmailAddresses(value []string)
	EmailHeaders() unsafe.Pointer
	SetEmailHeaders(value unsafe.Pointer)
	EncodingApplications() []string
	SetEncodingApplications(value []string)
	EndDate() foundation.NSDate
	SetEndDate(value foundation.IDate)
	EXIFVersion() string
	SetEXIFVersion(value string)
	EXIFGPSVersion() string
	SetEXIFGPSVersion(value string)
	ExposureMode() foundation.Number
	SetExposureMode(value foundation.INumber)
	ExposureProgram() string
	SetExposureProgram(value string)
	ExposureTime() foundation.Number
	SetExposureTime(value foundation.INumber)
	ExposureTimeString() string
	SetExposureTimeString(value string)
	FNumber() foundation.Number
	SetFNumber(value foundation.INumber)
	FileSize() foundation.Number
	SetFileSize(value foundation.INumber)
	FlashOn() foundation.Number
	SetFlashOn(value foundation.INumber)
	FocalLength() foundation.Number
	SetFocalLength(value foundation.INumber)
	FocalLength35mm() foundation.Number
	SetFocalLength35mm(value foundation.INumber)
	FontNames() []string
	SetFontNames(value []string)
	FullyFormattedAddress() string
	SetFullyFormattedAddress(value string)
	GeneralMIDISequence() foundation.Number
	SetGeneralMIDISequence(value foundation.INumber)
	Genre() string
	SetGenre(value string)
	GPSAreaInformation() string
	SetGPSAreaInformation(value string)
	GPSDateStamp() foundation.NSDate
	SetGPSDateStamp(value foundation.IDate)
	GPSDestBearing() foundation.Number
	SetGPSDestBearing(value foundation.INumber)
	GPSDestDistance() foundation.Number
	SetGPSDestDistance(value foundation.INumber)
	GPSDestLatitude() foundation.Number
	SetGPSDestLatitude(value foundation.INumber)
	GPSDestLongitude() foundation.Number
	SetGPSDestLongitude(value foundation.INumber)
	GPSDifferental() foundation.Number
	SetGPSDifferental(value foundation.INumber)
	GPSMapDatum() string
	SetGPSMapDatum(value string)
	GPSMeasureMode() string
	SetGPSMeasureMode(value string)
	GPSProcessingMethod() string
	SetGPSProcessingMethod(value string)
	GPSStatus() string
	SetGPSStatus(value string)
	GPSTrack() foundation.Number
	SetGPSTrack(value foundation.INumber)
	GPSDOP() foundation.Number
	SetGPSDOP(value foundation.INumber)
	HasAlphaChannel() foundation.Number
	SetHasAlphaChannel(value foundation.INumber)
	Headline() string
	SetHeadline(value string)
	HiddenAdditionalRecipients() []CSPerson
	SetHiddenAdditionalRecipients(value []CSPerson)
	HTMLContentData() foundation.NSData
	SetHTMLContentData(value foundation.IData)
	Identifier() string
	SetIdentifier(value string)
	ImageDirection() foundation.Number
	SetImageDirection(value foundation.INumber)
	ImportantDates() []foundation.Date
	SetImportantDates(value []foundation.IDate)
	Information() string
	SetInformation(value string)
	InstantMessageAddresses() []string
	SetInstantMessageAddresses(value []string)
	Instructions() string
	SetInstructions(value string)
	IsPriority() foundation.Number
	ISOSpeed() foundation.Number
	SetISOSpeed(value foundation.INumber)
	KeySignature() string
	SetKeySignature(value string)
	Keywords() []string
	SetKeywords(value []string)
	Kind() string
	SetKind(value string)
	Languages() []string
	SetLanguages(value []string)
	LastUsedDate() foundation.NSDate
	SetLastUsedDate(value foundation.IDate)
	Latitude() foundation.Number
	SetLatitude(value foundation.INumber)
	LayerNames() []string
	SetLayerNames(value []string)
	LensModel() string
	SetLensModel(value string)
	LikelyJunk() foundation.Number
	SetLikelyJunk(value foundation.INumber)
	Local() foundation.Number
	SetLocal(value foundation.INumber)
	Longitude() foundation.Number
	SetLongitude(value foundation.INumber)
	Lyricist() string
	SetLyricist(value string)
	MailboxIdentifiers() []string
	SetMailboxIdentifiers(value []string)
	MaxAperture() foundation.Number
	SetMaxAperture(value foundation.INumber)
	MediaTypes() []string
	SetMediaTypes(value []string)
	MetadataModificationDate() foundation.NSDate
	SetMetadataModificationDate(value foundation.IDate)
	MeteringMode() string
	SetMeteringMode(value string)
	MusicalGenre() string
	SetMusicalGenre(value string)
	MusicalInstrumentCategory() string
	SetMusicalInstrumentCategory(value string)
	MusicalInstrumentName() string
	SetMusicalInstrumentName(value string)
	NamedLocation() string
	SetNamedLocation(value string)
	Organizations() []string
	SetOrganizations(value []string)
	Orientation() foundation.Number
	SetOrientation(value foundation.INumber)
	OriginalFormat() string
	SetOriginalFormat(value string)
	OriginalSource() string
	SetOriginalSource(value string)
	PageCount() foundation.Number
	SetPageCount(value foundation.INumber)
	PageHeight() foundation.Number
	SetPageHeight(value foundation.INumber)
	PageWidth() foundation.Number
	SetPageWidth(value foundation.INumber)
	Participants() []string
	SetParticipants(value []string)
	Path() string
	SetPath(value string)
	Performers() []string
	SetPerformers(value []string)
	PhoneNumbers() []string
	SetPhoneNumbers(value []string)
	PixelCount() foundation.Number
	SetPixelCount(value foundation.INumber)
	PixelHeight() foundation.Number
	SetPixelHeight(value foundation.INumber)
	PixelWidth() foundation.Number
	SetPixelWidth(value foundation.INumber)
	PlayCount() foundation.Number
	SetPlayCount(value foundation.INumber)
	PostalCode() string
	SetPostalCode(value string)
	PrimaryRecipients() []CSPerson
	SetPrimaryRecipients(value []CSPerson)
	Producer() string
	SetProducer(value string)
	ProfileName() string
	SetProfileName(value string)
	Projects() []string
	SetProjects(value []string)
	ProviderDataTypeIdentifiers() []string
	SetProviderDataTypeIdentifiers(value []string)
	ProviderFileTypeIdentifiers() []string
	SetProviderFileTypeIdentifiers(value []string)
	ProviderInPlaceFileTypeIdentifiers() []string
	SetProviderInPlaceFileTypeIdentifiers(value []string)
	Publishers() []string
	SetPublishers(value []string)
	RankingHint() foundation.Number
	SetRankingHint(value foundation.INumber)
	Rating() foundation.Number
	SetRating(value foundation.INumber)
	RatingDescription() string
	SetRatingDescription(value string)
	RecipientAddresses() []string
	SetRecipientAddresses(value []string)
	RecipientEmailAddresses() []string
	SetRecipientEmailAddresses(value []string)
	RecipientNames() []string
	SetRecipientNames(value []string)
	RecordingDate() foundation.NSDate
	SetRecordingDate(value foundation.IDate)
	RedEyeOn() foundation.Number
	SetRedEyeOn(value foundation.INumber)
	RelatedUniqueIdentifier() string
	SetRelatedUniqueIdentifier(value string)
	ResolutionHeightDPI() foundation.Number
	SetResolutionHeightDPI(value foundation.INumber)
	ResolutionWidthDPI() foundation.Number
	SetResolutionWidthDPI(value foundation.INumber)
	Rights() string
	SetRights(value string)
	Role() string
	SetRole(value string)
	SecurityMethod() string
	SetSecurityMethod(value string)
	SharedItemContentType() unsafe.Pointer
	SetSharedItemContentType(value unsafe.Pointer)
	Speed() foundation.Number
	SetSpeed(value foundation.INumber)
	StartDate() foundation.NSDate
	SetStartDate(value foundation.IDate)
	StateOrProvince() string
	SetStateOrProvince(value string)
	Streamable() foundation.Number
	SetStreamable(value foundation.INumber)
	SubThoroughfare() string
	SetSubThoroughfare(value string)
	Subject() string
	SetSubject(value string)
	SupportsNavigation() foundation.Number
	SetSupportsNavigation(value foundation.INumber)
	SupportsPhoneCall() foundation.Number
	SetSupportsPhoneCall(value foundation.INumber)
	Tempo() foundation.Number
	SetTempo(value foundation.INumber)
	TextContent() string
	SetTextContent(value string)
	TextContentSummary() string
	Theme() string
	SetTheme(value string)
	Thoroughfare() string
	SetThoroughfare(value string)
	ThumbnailData() foundation.NSData
	SetThumbnailData(value foundation.IData)
	ThumbnailURL() foundation.URL
	SetThumbnailURL(value foundation.IURL)
	TimeSignature() string
	SetTimeSignature(value string)
	Timestamp() foundation.NSDate
	SetTimestamp(value foundation.IDate)
	Title() string
	SetTitle(value string)
	TotalBitRate() foundation.Number
	SetTotalBitRate(value foundation.INumber)
	TranscribedTextContent() string
	SetTranscribedTextContent(value string)
	URL() foundation.URL
	SetURL(value foundation.IURL)
	UserCreated() foundation.Number
	SetUserCreated(value foundation.INumber)
	UserCurated() foundation.Number
	SetUserCurated(value foundation.INumber)
	UserOwned() foundation.Number
	SetUserOwned(value foundation.INumber)
	Version() string
	SetVersion(value string)
	VideoBitRate() foundation.Number
	SetVideoBitRate(value foundation.INumber)
	WeakRelatedUniqueIdentifier() string
	SetWeakRelatedUniqueIdentifier(value string)
	WhiteBalance() foundation.Number
	SetWhiteBalance(value foundation.INumber)
	CSActionIdentifier() string
	Composer() string
	SetComposer(value string)
}

// The detailed metadata for a searchable item.
//
// A contains an extensive set of attributes that describe your app’s content. Attributes include information such as its title and a brief description. They can also refer to who created the item, what kind of data it represents, when someone created it, and more. During the indexing process, you create objects and use a to fill in the attributes for that item. During a search, you can query the index for items with attributes that match specific values. When creating a , it’s important to fill out as much information in the accompanying object as possible. You don’t have to provide values for every attribute. Instead, choose attributes that match the domain of your content. This type divides attributes into groups such as media, documents, events, places, music, images, and more. You can also add custom attributes to describe new types of content. When defining custom attributes, be as specific as possible in your definition, and provide a value for the property so your custom attribute inherits from a known type.


// The detailed metadata for a searchable item.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/init(contentType:)

func NewCSSearchableItemAttributeSetWithContentType(contentType unsafe.Pointer) CSSearchableItemAttributeSet {
	instance := getCSSearchableItemAttributeSetClass().Alloc()
	rv := objc.Send[CSSearchableItemAttributeSet](instance.ID, objc.Sel("initWithContentType:"), contentType)
	rv.Autorelease()
	return rv
}



// Creates an attribute set for the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/init(itemContentType:)

func NewCSSearchableItemAttributeSetWithItemContentType(itemContentType string) CSSearchableItemAttributeSet {
	instance := getCSSearchableItemAttributeSetClass().Alloc()
	rv := objc.Send[CSSearchableItemAttributeSet](instance.ID, objc.Sel("initWithItemContentType:"), objc.String(itemContentType))
	rv.Autorelease()
	return rv
}




// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/move(from:)

func (c_ CSSearchableItemAttributeSet) MoveFrom(sourceAttributeSet ICSSearchableItemAttributeSet) {
	objc.Send[objc.ID](c_.ID, objc.Sel("moveFrom:"), sourceAttributeSet)
}



// Sets the value for a custom attribute key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/setValue(_:forCustomKey:)

func (c_ CSSearchableItemAttributeSet) SetValueForCustomKey(value objectivec.IObject, key CSCustomAttributeKey) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setValue:forCustomKey:"), value, key)
}



// Returns the value associated with the specified custom attribute key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/value(forCustomKey:)

func (c_ CSSearchableItemAttributeSet) ValueForCustomKey(key CSCustomAttributeKey) objc.ID {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("valueForCustomKey:"), key)
	return rv
}


// An array of objects representing the content of the From: field in an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621608-authors

func (c_ CSSearchableItemAttributeSet) Authors() CSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("authors"))
	return rv
}


// An array of objects representing the content of the From: field in an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621608-authors

func (c_ CSSearchableItemAttributeSet) SetAuthors(value ICSPerson) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthors:"), value)
}


// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountHandles

func (c_ CSSearchableItemAttributeSet) AccountHandles() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("accountHandles"))
	return rv
}


// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountIdentifier

func (c_ CSSearchableItemAttributeSet) AccountIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("accountIdentifier"))
	return rv
}


// The unique identifier for the account with which the message is associated, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountIdentifier

func (c_ CSSearchableItemAttributeSet) SetAccountIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccountIdentifier:"), objc.String(value))
}


// The manufacturer of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionMake

func (c_ CSSearchableItemAttributeSet) AcquisitionMake() string {
	rv := objc.Send[string](c_.ID, objc.Sel("acquisitionMake"))
	return rv
}


// The manufacturer of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionMake

func (c_ CSSearchableItemAttributeSet) SetAcquisitionMake(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionMake:"), objc.String(value))
}


// The model of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionModel

func (c_ CSSearchableItemAttributeSet) AcquisitionModel() string {
	rv := objc.Send[string](c_.ID, objc.Sel("acquisitionModel"))
	return rv
}


// The model of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionModel

func (c_ CSSearchableItemAttributeSet) SetAcquisitionModel(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionModel:"), objc.String(value))
}


// The identifiers that specify custom actions the app supports for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/actionIdentifiers

func (c_ CSSearchableItemAttributeSet) ActionIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("actionIdentifiers"))
	return rv
}


// The identifiers that specify custom actions the app supports for the item.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/addedDate

func (c_ CSSearchableItemAttributeSet) AddedDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("addedDate"))
	return rv
}


// The date on which the item was moved into its current location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/addedDate

func (c_ CSSearchableItemAttributeSet) SetAddedDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAddedDate:"), value)
}


// An array of objects representing the content of the Cc: field in an email message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/additionalRecipients

func (c_ CSSearchableItemAttributeSet) AdditionalRecipients() []CSPerson {
	rv := objc.Send[[]CSPerson](c_.ID, objc.Sel("additionalRecipients"))
	return rv
}


// An array of objects representing the content of the Cc: field in an email message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/album

func (c_ CSSearchableItemAttributeSet) Album() string {
	rv := objc.Send[string](c_.ID, objc.Sel("album"))
	return rv
}


// The title for a collection of audio media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/album

func (c_ CSSearchableItemAttributeSet) SetAlbum(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlbum:"), objc.String(value))
}


// A value that indicates if the event covers an entire day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/allDay

func (c_ CSSearchableItemAttributeSet) AllDay() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("allDay"))
	return rv
}


// A value that indicates if the event covers an entire day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/allDay

func (c_ CSSearchableItemAttributeSet) SetAllDay(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllDay:"), value)
}


// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/alternateNames

func (c_ CSSearchableItemAttributeSet) AlternateNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("alternateNames"))
	return rv
}


// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/altitude

func (c_ CSSearchableItemAttributeSet) Altitude() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("altitude"))
	return rv
}


// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/altitude

func (c_ CSSearchableItemAttributeSet) SetAltitude(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAltitude:"), value)
}


// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/aperture

func (c_ CSSearchableItemAttributeSet) Aperture() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("aperture"))
	return rv
}


// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/aperture

func (c_ CSSearchableItemAttributeSet) SetAperture(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAperture:"), value)
}


// The artist associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/artist

func (c_ CSSearchableItemAttributeSet) Artist() string {
	rv := objc.Send[string](c_.ID, objc.Sel("artist"))
	return rv
}


// The artist associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/artist

func (c_ CSSearchableItemAttributeSet) SetArtist(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setArtist:"), objc.String(value))
}


// A class of entity for which the item is intended or useful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audiences

func (c_ CSSearchableItemAttributeSet) Audiences() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("audiences"))
	return rv
}


// A class of entity for which the item is intended or useful.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioBitRate

func (c_ CSSearchableItemAttributeSet) AudioBitRate() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioBitRate"))
	return rv
}


// The audio bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioBitRate

func (c_ CSSearchableItemAttributeSet) SetAudioBitRate(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioBitRate:"), value)
}


// The number of channels in the audio data that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioChannelCount

func (c_ CSSearchableItemAttributeSet) AudioChannelCount() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioChannelCount"))
	return rv
}


// The number of channels in the audio data that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioChannelCount

func (c_ CSSearchableItemAttributeSet) SetAudioChannelCount(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioChannelCount:"), value)
}


// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioEncodingApplication

func (c_ CSSearchableItemAttributeSet) AudioEncodingApplication() string {
	rv := objc.Send[string](c_.ID, objc.Sel("audioEncodingApplication"))
	return rv
}


// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioEncodingApplication

func (c_ CSSearchableItemAttributeSet) SetAudioEncodingApplication(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioEncodingApplication:"), objc.String(value))
}


// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioSampleRate

func (c_ CSSearchableItemAttributeSet) AudioSampleRate() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioSampleRate"))
	return rv
}


// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioSampleRate

func (c_ CSSearchableItemAttributeSet) SetAudioSampleRate(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSampleRate:"), value)
}


// The track number of a song or audio composition when part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioTrackNumber

func (c_ CSSearchableItemAttributeSet) AudioTrackNumber() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioTrackNumber"))
	return rv
}


// The track number of a song or audio composition when part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioTrackNumber

func (c_ CSSearchableItemAttributeSet) SetAudioTrackNumber(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioTrackNumber:"), value)
}


// An array of addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorAddresses

func (c_ CSSearchableItemAttributeSet) AuthorAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("authorAddresses"))
	return rv
}


// An array of addresses associated with the author of the message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorEmailAddresses

func (c_ CSSearchableItemAttributeSet) AuthorEmailAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("authorEmailAddresses"))
	return rv
}


// An array of email addresses associated with the author of the message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorNames

func (c_ CSSearchableItemAttributeSet) AuthorNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("authorNames"))
	return rv
}


// An array of names representing the authors who have worked on the message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/bitsPerSample

func (c_ CSSearchableItemAttributeSet) BitsPerSample() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("bitsPerSample"))
	return rv
}


// The number of bits per sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/bitsPerSample

func (c_ CSSearchableItemAttributeSet) SetBitsPerSample(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBitsPerSample:"), value)
}


// The owner of the camera that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/cameraOwner

func (c_ CSSearchableItemAttributeSet) CameraOwner() string {
	rv := objc.Send[string](c_.ID, objc.Sel("cameraOwner"))
	return rv
}


// The owner of the camera that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/cameraOwner

func (c_ CSSearchableItemAttributeSet) SetCameraOwner(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraOwner:"), objc.String(value))
}


// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/city

func (c_ CSSearchableItemAttributeSet) City() string {
	rv := objc.Send[string](c_.ID, objc.Sel("city"))
	return rv
}


// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/city

func (c_ CSSearchableItemAttributeSet) SetCity(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), objc.String(value))
}


// The codecs used to encode/decode the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/codecs

func (c_ CSSearchableItemAttributeSet) Codecs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("codecs"))
	return rv
}


// The codecs used to encode/decode the media.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/colorSpace

func (c_ CSSearchableItemAttributeSet) ColorSpace() string {
	rv := objc.Send[string](c_.ID, objc.Sel("colorSpace"))
	return rv
}


// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/colorSpace

func (c_ CSSearchableItemAttributeSet) SetColorSpace(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSpace:"), objc.String(value))
}


// A comment related to the media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/comment

func (c_ CSSearchableItemAttributeSet) Comment() string {
	rv := objc.Send[string](c_.ID, objc.Sel("comment"))
	return rv
}


// A comment related to the media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/comment

func (c_ CSSearchableItemAttributeSet) SetComment(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComment:"), objc.String(value))
}


// The date on which the item was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/completionDate

func (c_ CSSearchableItemAttributeSet) CompletionDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("completionDate"))
	return rv
}


// The date on which the item was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/completionDate

func (c_ CSSearchableItemAttributeSet) SetCompletionDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionDate:"), value)
}


// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contactKeywords

func (c_ CSSearchableItemAttributeSet) ContactKeywords() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contactKeywords"))
	return rv
}


// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerDisplayName

func (c_ CSSearchableItemAttributeSet) ContainerDisplayName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerDisplayName"))
	return rv
}


// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerDisplayName

func (c_ CSSearchableItemAttributeSet) SetContainerDisplayName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerDisplayName:"), objc.String(value))
}


// The identifier of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerIdentifier

func (c_ CSSearchableItemAttributeSet) ContainerIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}


// The identifier of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerIdentifier

func (c_ CSSearchableItemAttributeSet) SetContainerIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), objc.String(value))
}


// The order of the item within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerOrder

func (c_ CSSearchableItemAttributeSet) ContainerOrder() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("containerOrder"))
	return rv
}


// The order of the item within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerOrder

func (c_ CSSearchableItemAttributeSet) SetContainerOrder(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerOrder:"), value)
}


// The title of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerTitle

func (c_ CSSearchableItemAttributeSet) ContainerTitle() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerTitle"))
	return rv
}


// The title of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerTitle

func (c_ CSSearchableItemAttributeSet) SetContainerTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerTitle:"), objc.String(value))
}


// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentCreationDate

func (c_ CSSearchableItemAttributeSet) ContentCreationDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("contentCreationDate"))
	return rv
}


// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentCreationDate

func (c_ CSSearchableItemAttributeSet) SetContentCreationDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentCreationDate:"), value)
}


// A description of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentDescription

func (c_ CSSearchableItemAttributeSet) ContentDescription() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contentDescription"))
	return rv
}


// A description of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentDescription

func (c_ CSSearchableItemAttributeSet) SetContentDescription(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentDescription:"), objc.String(value))
}


// The date on which the contents of the file was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentModificationDate

func (c_ CSSearchableItemAttributeSet) ContentModificationDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("contentModificationDate"))
	return rv
}


// The date on which the contents of the file was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentModificationDate

func (c_ CSSearchableItemAttributeSet) SetContentModificationDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentModificationDate:"), value)
}


// A value that indicates if the media contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentRating

func (c_ CSSearchableItemAttributeSet) ContentRating() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("contentRating"))
	return rv
}


// A value that indicates if the media contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentRating

func (c_ CSSearchableItemAttributeSet) SetContentRating(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentRating:"), value)
}


// An array of sources from which the media was obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentSources

func (c_ CSSearchableItemAttributeSet) ContentSources() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contentSources"))
	return rv
}


// An array of sources from which the media was obtained.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentType

func (c_ CSSearchableItemAttributeSet) ContentType() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contentType"))
	return rv
}


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentType

func (c_ CSSearchableItemAttributeSet) SetContentType(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentType:"), objc.String(value))
}


// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentTypeTree

func (c_ CSSearchableItemAttributeSet) ContentTypeTree() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contentTypeTree"))
	return rv
}


// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentURL

func (c_ CSSearchableItemAttributeSet) ContentURL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("contentURL"))
	return rv
}


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentURL

func (c_ CSSearchableItemAttributeSet) SetContentURL(value foundation.IURL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentURL:"), value)
}


// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contributors

func (c_ CSSearchableItemAttributeSet) Contributors() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contributors"))
	return rv
}


// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/copyright

func (c_ CSSearchableItemAttributeSet) Copyright() string {
	rv := objc.Send[string](c_.ID, objc.Sel("copyright"))
	return rv
}


// The copyright date of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/copyright

func (c_ CSSearchableItemAttributeSet) SetCopyright(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCopyright:"), objc.String(value))
}


// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/country

func (c_ CSSearchableItemAttributeSet) Country() string {
	rv := objc.Send[string](c_.ID, objc.Sel("country"))
	return rv
}


// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/country

func (c_ CSSearchableItemAttributeSet) SetCountry(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), objc.String(value))
}


// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/coverage

func (c_ CSSearchableItemAttributeSet) Coverage() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("coverage"))
	return rv
}


// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/creator

func (c_ CSSearchableItemAttributeSet) Creator() string {
	rv := objc.Send[string](c_.ID, objc.Sel("creator"))
	return rv
}


// The name of the app that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/creator

func (c_ CSSearchableItemAttributeSet) SetCreator(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreator:"), objc.String(value))
}


// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/darkThumbnailURL

func (c_ CSSearchableItemAttributeSet) DarkThumbnailURL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("darkThumbnailURL"))
	return rv
}


// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/darkThumbnailURL

func (c_ CSSearchableItemAttributeSet) SetDarkThumbnailURL(value foundation.IURL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDarkThumbnailURL:"), value)
}


// The delivery type of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/deliveryType

func (c_ CSSearchableItemAttributeSet) DeliveryType() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("deliveryType"))
	return rv
}


// The delivery type of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/deliveryType

func (c_ CSSearchableItemAttributeSet) SetDeliveryType(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeliveryType:"), value)
}


// The name of the director of the media (for example, a movie director).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/director

func (c_ CSSearchableItemAttributeSet) Director() string {
	rv := objc.Send[string](c_.ID, objc.Sel("director"))
	return rv
}


// The name of the director of the media (for example, a movie director).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/director

func (c_ CSSearchableItemAttributeSet) SetDirector(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDirector:"), objc.String(value))
}


// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName

func (c_ CSSearchableItemAttributeSet) DisplayName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("displayName"))
	return rv
}


// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName

func (c_ CSSearchableItemAttributeSet) SetDisplayName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}


// An identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/domainIdentifier

func (c_ CSSearchableItemAttributeSet) DomainIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("domainIdentifier"))
	return rv
}


// An identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/domainIdentifier

func (c_ CSSearchableItemAttributeSet) SetDomainIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDomainIdentifier:"), objc.String(value))
}


// The most recent date on which the file was downloaded or received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/downloadedDate

func (c_ CSSearchableItemAttributeSet) DownloadedDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("downloadedDate"))
	return rv
}


// The most recent date on which the file was downloaded or received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/downloadedDate

func (c_ CSSearchableItemAttributeSet) SetDownloadedDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDownloadedDate:"), value)
}


// The date on which the item is due.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/dueDate

func (c_ CSSearchableItemAttributeSet) DueDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("dueDate"))
	return rv
}


// The date on which the item is due.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/dueDate

func (c_ CSSearchableItemAttributeSet) SetDueDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDueDate:"), value)
}


// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/duration

func (c_ CSSearchableItemAttributeSet) Duration() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("duration"))
	return rv
}


// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/duration

func (c_ CSSearchableItemAttributeSet) SetDuration(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDuration:"), value)
}


// A list of editors who have worked on the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/editors

func (c_ CSSearchableItemAttributeSet) Editors() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("editors"))
	return rv
}


// A list of editors who have worked on the file.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailAddresses

func (c_ CSSearchableItemAttributeSet) EmailAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("emailAddresses"))
	return rv
}


// An array of email addresses associated with the message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailHeaders

func (c_ CSSearchableItemAttributeSet) EmailHeaders() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("emailHeaders"))
	return rv
}


// A dictionary that contains all the headers of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailHeaders

func (c_ CSSearchableItemAttributeSet) SetEmailHeaders(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailHeaders:"), value)
}


// The name of the apps that converted the original content into a PDF stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/encodingApplications

func (c_ CSSearchableItemAttributeSet) EncodingApplications() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("encodingApplications"))
	return rv
}


// The name of the apps that converted the original content into a PDF stream.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/endDate

func (c_ CSSearchableItemAttributeSet) EndDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("endDate"))
	return rv
}


// The end date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/endDate

func (c_ CSSearchableItemAttributeSet) SetEndDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEndDate:"), value)
}


// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifVersion

func (c_ CSSearchableItemAttributeSet) EXIFVersion() string {
	rv := objc.Send[string](c_.ID, objc.Sel("EXIFVersion"))
	return rv
}


// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifVersion

func (c_ CSSearchableItemAttributeSet) SetEXIFVersion(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEXIFVersion:"), objc.String(value))
}


// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifgpsVersion

func (c_ CSSearchableItemAttributeSet) EXIFGPSVersion() string {
	rv := objc.Send[string](c_.ID, objc.Sel("EXIFGPSVersion"))
	return rv
}


// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifgpsVersion

func (c_ CSSearchableItemAttributeSet) SetEXIFGPSVersion(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEXIFGPSVersion:"), objc.String(value))
}


// The mode the camera used for the exposure of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureMode

func (c_ CSSearchableItemAttributeSet) ExposureMode() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("exposureMode"))
	return rv
}


// The mode the camera used for the exposure of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureMode

func (c_ CSSearchableItemAttributeSet) SetExposureMode(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureMode:"), value)
}


// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureProgram

func (c_ CSSearchableItemAttributeSet) ExposureProgram() string {
	rv := objc.Send[string](c_.ID, objc.Sel("exposureProgram"))
	return rv
}


// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureProgram

func (c_ CSSearchableItemAttributeSet) SetExposureProgram(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureProgram:"), objc.String(value))
}


// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTime

func (c_ CSSearchableItemAttributeSet) ExposureTime() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("exposureTime"))
	return rv
}


// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTime

func (c_ CSSearchableItemAttributeSet) SetExposureTime(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTime:"), value)
}


// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTimeString

func (c_ CSSearchableItemAttributeSet) ExposureTimeString() string {
	rv := objc.Send[string](c_.ID, objc.Sel("exposureTimeString"))
	return rv
}


// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTimeString

func (c_ CSSearchableItemAttributeSet) SetExposureTimeString(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTimeString:"), objc.String(value))
}


// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fNumber

func (c_ CSSearchableItemAttributeSet) FNumber() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("fNumber"))
	return rv
}


// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fNumber

func (c_ CSSearchableItemAttributeSet) SetFNumber(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFNumber:"), value)
}


// The size of the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fileSize

func (c_ CSSearchableItemAttributeSet) FileSize() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("fileSize"))
	return rv
}


// The size of the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fileSize

func (c_ CSSearchableItemAttributeSet) SetFileSize(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFileSize:"), value)
}


// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/flashOn

func (c_ CSSearchableItemAttributeSet) FlashOn() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("flashOn"))
	return rv
}


// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/flashOn

func (c_ CSSearchableItemAttributeSet) SetFlashOn(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashOn:"), value)
}


// The actual focal length of the lens, in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength

func (c_ CSSearchableItemAttributeSet) FocalLength() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("focalLength"))
	return rv
}


// The actual focal length of the lens, in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength

func (c_ CSSearchableItemAttributeSet) SetFocalLength(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength:"), value)
}


// A value that indicates if the focal length is 35mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength35mm

func (c_ CSSearchableItemAttributeSet) FocalLength35mm() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("focalLength35mm"))
	return rv
}


// A value that indicates if the focal length is 35mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength35mm

func (c_ CSSearchableItemAttributeSet) SetFocalLength35mm(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength35mm:"), value)
}


// An array of font names the document uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fontNames

func (c_ CSSearchableItemAttributeSet) FontNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("fontNames"))
	return rv
}


// An array of font names the document uses.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fullyFormattedAddress

func (c_ CSSearchableItemAttributeSet) FullyFormattedAddress() string {
	rv := objc.Send[string](c_.ID, objc.Sel("fullyFormattedAddress"))
	return rv
}


// The fully formatted address of the item, received from MapKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fullyFormattedAddress

func (c_ CSSearchableItemAttributeSet) SetFullyFormattedAddress(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFullyFormattedAddress:"), objc.String(value))
}


// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/generalMIDISequence

func (c_ CSSearchableItemAttributeSet) GeneralMIDISequence() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("generalMIDISequence"))
	return rv
}


// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/generalMIDISequence

func (c_ CSSearchableItemAttributeSet) SetGeneralMIDISequence(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGeneralMIDISequence:"), value)
}


// The genre of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/genre

func (c_ CSSearchableItemAttributeSet) Genre() string {
	rv := objc.Send[string](c_.ID, objc.Sel("genre"))
	return rv
}


// The genre of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/genre

func (c_ CSSearchableItemAttributeSet) SetGenre(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGenre:"), objc.String(value))
}


// Information about the GPS area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsAreaInformation

func (c_ CSSearchableItemAttributeSet) GPSAreaInformation() string {
	rv := objc.Send[string](c_.ID, objc.Sel("GPSAreaInformation"))
	return rv
}


// Information about the GPS area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsAreaInformation

func (c_ CSSearchableItemAttributeSet) SetGPSAreaInformation(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSAreaInformation:"), objc.String(value))
}


// The date and time related to the GPS value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDateStamp

func (c_ CSSearchableItemAttributeSet) GPSDateStamp() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("GPSDateStamp"))
	return rv
}


// The date and time related to the GPS value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDateStamp

func (c_ CSSearchableItemAttributeSet) SetGPSDateStamp(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDateStamp:"), value)
}


// The bearing to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestBearing

func (c_ CSSearchableItemAttributeSet) GPSDestBearing() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("GPSDestBearing"))
	return rv
}


// The bearing to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestBearing

func (c_ CSSearchableItemAttributeSet) SetGPSDestBearing(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestBearing:"), value)
}


// The distance to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestDistance

func (c_ CSSearchableItemAttributeSet) GPSDestDistance() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("GPSDestDistance"))
	return rv
}


// The distance to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestDistance

func (c_ CSSearchableItemAttributeSet) SetGPSDestDistance(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestDistance:"), value)
}


// The latitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLatitude

func (c_ CSSearchableItemAttributeSet) GPSDestLatitude() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("GPSDestLatitude"))
	return rv
}


// The latitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLatitude

func (c_ CSSearchableItemAttributeSet) SetGPSDestLatitude(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestLatitude:"), value)
}


// The longitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLongitude

func (c_ CSSearchableItemAttributeSet) GPSDestLongitude() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("GPSDestLongitude"))
	return rv
}


// The longitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLongitude

func (c_ CSSearchableItemAttributeSet) SetGPSDestLongitude(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestLongitude:"), value)
}


// The differential correction applied to the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDifferental

func (c_ CSSearchableItemAttributeSet) GPSDifferental() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("GPSDifferental"))
	return rv
}


// The differential correction applied to the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDifferental

func (c_ CSSearchableItemAttributeSet) SetGPSDifferental(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDifferental:"), value)
}


// The geodetic data that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMapDatum

func (c_ CSSearchableItemAttributeSet) GPSMapDatum() string {
	rv := objc.Send[string](c_.ID, objc.Sel("GPSMapDatum"))
	return rv
}


// The geodetic data that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMapDatum

func (c_ CSSearchableItemAttributeSet) SetGPSMapDatum(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSMapDatum:"), objc.String(value))
}


// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMeasureMode

func (c_ CSSearchableItemAttributeSet) GPSMeasureMode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("GPSMeasureMode"))
	return rv
}


// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMeasureMode

func (c_ CSSearchableItemAttributeSet) SetGPSMeasureMode(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSMeasureMode:"), objc.String(value))
}


// The location finding method that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsProcessingMethod

func (c_ CSSearchableItemAttributeSet) GPSProcessingMethod() string {
	rv := objc.Send[string](c_.ID, objc.Sel("GPSProcessingMethod"))
	return rv
}


// The location finding method that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsProcessingMethod

func (c_ CSSearchableItemAttributeSet) SetGPSProcessingMethod(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSProcessingMethod:"), objc.String(value))
}


// The status of the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsStatus

func (c_ CSSearchableItemAttributeSet) GPSStatus() string {
	rv := objc.Send[string](c_.ID, objc.Sel("GPSStatus"))
	return rv
}


// The status of the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsStatus

func (c_ CSSearchableItemAttributeSet) SetGPSStatus(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSStatus:"), objc.String(value))
}


// The direction of travel of the item in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsTrack

func (c_ CSSearchableItemAttributeSet) GPSTrack() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("GPSTrack"))
	return rv
}


// The direction of travel of the item in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsTrack

func (c_ CSSearchableItemAttributeSet) SetGPSTrack(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSTrack:"), value)
}


// The GPS dilution of precision value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop

func (c_ CSSearchableItemAttributeSet) GPSDOP() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("GPSDOP"))
	return rv
}


// The GPS dilution of precision value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop

func (c_ CSSearchableItemAttributeSet) SetGPSDOP(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDOP:"), value)
}


// Indicates if the image file has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hasAlphaChannel

func (c_ CSSearchableItemAttributeSet) HasAlphaChannel() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("hasAlphaChannel"))
	return rv
}


// Indicates if the image file has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hasAlphaChannel

func (c_ CSSearchableItemAttributeSet) SetHasAlphaChannel(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasAlphaChannel:"), value)
}


// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/headline

func (c_ CSSearchableItemAttributeSet) Headline() string {
	rv := objc.Send[string](c_.ID, objc.Sel("headline"))
	return rv
}


// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/headline

func (c_ CSSearchableItemAttributeSet) SetHeadline(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHeadline:"), objc.String(value))
}


// An array of objects representing the content of the Bcc: field in an email message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hiddenAdditionalRecipients

func (c_ CSSearchableItemAttributeSet) HiddenAdditionalRecipients() []CSPerson {
	rv := objc.Send[[]CSPerson](c_.ID, objc.Sel("hiddenAdditionalRecipients"))
	return rv
}


// An array of objects representing the content of the Bcc: field in an email message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/htmlContentData

func (c_ CSSearchableItemAttributeSet) HTMLContentData() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("HTMLContentData"))
	return rv
}


// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/htmlContentData

func (c_ CSSearchableItemAttributeSet) SetHTMLContentData(value foundation.IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHTMLContentData:"), value)
}


// A formal identifier that references the document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/identifier

func (c_ CSSearchableItemAttributeSet) Identifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("identifier"))
	return rv
}


// A formal identifier that references the document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/identifier

func (c_ CSSearchableItemAttributeSet) SetIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The direction of the item’s image in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/imageDirection

func (c_ CSSearchableItemAttributeSet) ImageDirection() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("imageDirection"))
	return rv
}


// The direction of the item’s image in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/imageDirection

func (c_ CSSearchableItemAttributeSet) SetImageDirection(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageDirection:"), value)
}


// An array of important dates associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/importantDates

func (c_ CSSearchableItemAttributeSet) ImportantDates() []foundation.Date {
	rv := objc.Send[[]foundation.Date](c_.ID, objc.Sel("importantDates"))
	return rv
}


// An array of important dates associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/importantDates

func (c_ CSSearchableItemAttributeSet) SetImportantDates(value []foundation.IDate) {
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/information

func (c_ CSSearchableItemAttributeSet) Information() string {
	rv := objc.Send[string](c_.ID, objc.Sel("information"))
	return rv
}


// Information about the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/information

func (c_ CSSearchableItemAttributeSet) SetInformation(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInformation:"), objc.String(value))
}


// An array of instant message addresses for the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instantMessageAddresses

func (c_ CSSearchableItemAttributeSet) InstantMessageAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}


// An array of instant message addresses for the message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instructions

func (c_ CSSearchableItemAttributeSet) Instructions() string {
	rv := objc.Send[string](c_.ID, objc.Sel("instructions"))
	return rv
}


// Instructions that concern the use of the item, such as an embargo or warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instructions

func (c_ CSSearchableItemAttributeSet) SetInstructions(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstructions:"), objc.String(value))
}


// A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isPriority

func (c_ CSSearchableItemAttributeSet) IsPriority() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("isPriority"))
	return rv
}


// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isoSpeed

func (c_ CSSearchableItemAttributeSet) ISOSpeed() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("ISOSpeed"))
	return rv
}


// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isoSpeed

func (c_ CSSearchableItemAttributeSet) SetISOSpeed(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setISOSpeed:"), value)
}


// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keySignature

func (c_ CSSearchableItemAttributeSet) KeySignature() string {
	rv := objc.Send[string](c_.ID, objc.Sel("keySignature"))
	return rv
}


// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keySignature

func (c_ CSSearchableItemAttributeSet) SetKeySignature(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeySignature:"), objc.String(value))
}


// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keywords

func (c_ CSSearchableItemAttributeSet) Keywords() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("keywords"))
	return rv
}


// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/kind

func (c_ CSSearchableItemAttributeSet) Kind() string {
	rv := objc.Send[string](c_.ID, objc.Sel("kind"))
	return rv
}


// A description of the kind of document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/kind

func (c_ CSSearchableItemAttributeSet) SetKind(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKind:"), objc.String(value))
}


// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/languages

func (c_ CSSearchableItemAttributeSet) Languages() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("languages"))
	return rv
}


// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lastUsedDate

func (c_ CSSearchableItemAttributeSet) LastUsedDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("lastUsedDate"))
	return rv
}


// The date on which the file was last used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lastUsedDate

func (c_ CSSearchableItemAttributeSet) SetLastUsedDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLastUsedDate:"), value)
}


// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/latitude

func (c_ CSSearchableItemAttributeSet) Latitude() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("latitude"))
	return rv
}


// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/latitude

func (c_ CSSearchableItemAttributeSet) SetLatitude(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLatitude:"), value)
}


// An array that contains the names of the various layers in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/layerNames

func (c_ CSSearchableItemAttributeSet) LayerNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("layerNames"))
	return rv
}


// An array that contains the names of the various layers in the file.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lensModel

func (c_ CSSearchableItemAttributeSet) LensModel() string {
	rv := objc.Send[string](c_.ID, objc.Sel("lensModel"))
	return rv
}


// The model of the lens that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lensModel

func (c_ CSSearchableItemAttributeSet) SetLensModel(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLensModel:"), objc.String(value))
}


// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/likelyJunk

func (c_ CSSearchableItemAttributeSet) LikelyJunk() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("likelyJunk"))
	return rv
}


// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/likelyJunk

func (c_ CSSearchableItemAttributeSet) SetLikelyJunk(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLikelyJunk:"), value)
}


// A value that indicates if the media is local.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/local

func (c_ CSSearchableItemAttributeSet) Local() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("local"))
	return rv
}


// A value that indicates if the media is local.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/local

func (c_ CSSearchableItemAttributeSet) SetLocal(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocal:"), value)
}


// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/longitude

func (c_ CSSearchableItemAttributeSet) Longitude() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("longitude"))
	return rv
}


// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/longitude

func (c_ CSSearchableItemAttributeSet) SetLongitude(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongitude:"), value)
}


// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lyricist

func (c_ CSSearchableItemAttributeSet) Lyricist() string {
	rv := objc.Send[string](c_.ID, objc.Sel("lyricist"))
	return rv
}


// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lyricist

func (c_ CSSearchableItemAttributeSet) SetLyricist(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLyricist:"), objc.String(value))
}


// An array of mailbox identifiers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mailboxIdentifiers

func (c_ CSSearchableItemAttributeSet) MailboxIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("mailboxIdentifiers"))
	return rv
}


// An array of mailbox identifiers associated with the message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/maxAperture

func (c_ CSSearchableItemAttributeSet) MaxAperture() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("maxAperture"))
	return rv
}


// The smallest F number of the lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/maxAperture

func (c_ CSSearchableItemAttributeSet) SetMaxAperture(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxAperture:"), value)
}


// The media types present in the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mediaTypes

func (c_ CSSearchableItemAttributeSet) MediaTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("mediaTypes"))
	return rv
}


// The media types present in the content.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/metadataModificationDate

func (c_ CSSearchableItemAttributeSet) MetadataModificationDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("metadataModificationDate"))
	return rv
}


// The date on which the last metadata attribute was changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/metadataModificationDate

func (c_ CSSearchableItemAttributeSet) SetMetadataModificationDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataModificationDate:"), value)
}


// The metering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/meteringMode

func (c_ CSSearchableItemAttributeSet) MeteringMode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("meteringMode"))
	return rv
}


// The metering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/meteringMode

func (c_ CSSearchableItemAttributeSet) SetMeteringMode(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMeteringMode:"), objc.String(value))
}


// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalGenre

func (c_ CSSearchableItemAttributeSet) MusicalGenre() string {
	rv := objc.Send[string](c_.ID, objc.Sel("musicalGenre"))
	return rv
}


// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalGenre

func (c_ CSSearchableItemAttributeSet) SetMusicalGenre(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalGenre:"), objc.String(value))
}


// The category of the instrument associated with the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentCategory

func (c_ CSSearchableItemAttributeSet) MusicalInstrumentCategory() string {
	rv := objc.Send[string](c_.ID, objc.Sel("musicalInstrumentCategory"))
	return rv
}


// The category of the instrument associated with the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentCategory

func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentCategory(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentCategory:"), objc.String(value))
}


// The name of an instrument within the context of an instrument category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentName

func (c_ CSSearchableItemAttributeSet) MusicalInstrumentName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("musicalInstrumentName"))
	return rv
}


// The name of an instrument within the context of an instrument category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentName

func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentName:"), objc.String(value))
}


// The name of the location or point of interest associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/namedLocation

func (c_ CSSearchableItemAttributeSet) NamedLocation() string {
	rv := objc.Send[string](c_.ID, objc.Sel("namedLocation"))
	return rv
}


// The name of the location or point of interest associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/namedLocation

func (c_ CSSearchableItemAttributeSet) SetNamedLocation(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNamedLocation:"), objc.String(value))
}


// A list of companies or organizations that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/organizations

func (c_ CSSearchableItemAttributeSet) Organizations() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("organizations"))
	return rv
}


// A list of companies or organizations that created the content.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/orientation

func (c_ CSSearchableItemAttributeSet) Orientation() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("orientation"))
	return rv
}


// The orientation of the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/orientation

func (c_ CSSearchableItemAttributeSet) SetOrientation(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrientation:"), value)
}


// The original format of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalFormat

func (c_ CSSearchableItemAttributeSet) OriginalFormat() string {
	rv := objc.Send[string](c_.ID, objc.Sel("originalFormat"))
	return rv
}


// The original format of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalFormat

func (c_ CSSearchableItemAttributeSet) SetOriginalFormat(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalFormat:"), objc.String(value))
}


// The original source of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalSource

func (c_ CSSearchableItemAttributeSet) OriginalSource() string {
	rv := objc.Send[string](c_.ID, objc.Sel("originalSource"))
	return rv
}


// The original source of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalSource

func (c_ CSSearchableItemAttributeSet) SetOriginalSource(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalSource:"), objc.String(value))
}


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageCount

func (c_ CSSearchableItemAttributeSet) PageCount() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pageCount"))
	return rv
}


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageCount

func (c_ CSSearchableItemAttributeSet) SetPageCount(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageCount:"), value)
}


// The height of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageHeight

func (c_ CSSearchableItemAttributeSet) PageHeight() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pageHeight"))
	return rv
}


// The height of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageHeight

func (c_ CSSearchableItemAttributeSet) SetPageHeight(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageHeight:"), value)
}


// The width of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageWidth

func (c_ CSSearchableItemAttributeSet) PageWidth() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pageWidth"))
	return rv
}


// The width of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageWidth

func (c_ CSSearchableItemAttributeSet) SetPageWidth(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageWidth:"), value)
}


// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/participants

func (c_ CSSearchableItemAttributeSet) Participants() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("participants"))
	return rv
}


// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/path

func (c_ CSSearchableItemAttributeSet) Path() string {
	rv := objc.Send[string](c_.ID, objc.Sel("path"))
	return rv
}


// The complete path to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/path

func (c_ CSSearchableItemAttributeSet) SetPath(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPath:"), objc.String(value))
}


// A list of performers in the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/performers

func (c_ CSSearchableItemAttributeSet) Performers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("performers"))
	return rv
}


// A list of performers in the media.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/phoneNumbers

func (c_ CSSearchableItemAttributeSet) PhoneNumbers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}


// An array of phone numbers associated with the message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelCount

func (c_ CSSearchableItemAttributeSet) PixelCount() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pixelCount"))
	return rv
}


// The total number of pixels in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelCount

func (c_ CSSearchableItemAttributeSet) SetPixelCount(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelCount:"), value)
}


// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelHeight

func (c_ CSSearchableItemAttributeSet) PixelHeight() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pixelHeight"))
	return rv
}


// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelHeight

func (c_ CSSearchableItemAttributeSet) SetPixelHeight(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelHeight:"), value)
}


// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelWidth

func (c_ CSSearchableItemAttributeSet) PixelWidth() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pixelWidth"))
	return rv
}


// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelWidth

func (c_ CSSearchableItemAttributeSet) SetPixelWidth(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelWidth:"), value)
}


// A user-supplied play count for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/playCount

func (c_ CSSearchableItemAttributeSet) PlayCount() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("playCount"))
	return rv
}


// A user-supplied play count for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/playCount

func (c_ CSSearchableItemAttributeSet) SetPlayCount(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayCount:"), value)
}


// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/postalCode

func (c_ CSSearchableItemAttributeSet) PostalCode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("postalCode"))
	return rv
}


// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/postalCode

func (c_ CSSearchableItemAttributeSet) SetPostalCode(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalCode:"), objc.String(value))
}


// An array of objects representing the content of the To: field in an email message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/primaryRecipients

func (c_ CSSearchableItemAttributeSet) PrimaryRecipients() []CSPerson {
	rv := objc.Send[[]CSPerson](c_.ID, objc.Sel("primaryRecipients"))
	return rv
}


// An array of objects representing the content of the To: field in an email message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/producer

func (c_ CSSearchableItemAttributeSet) Producer() string {
	rv := objc.Send[string](c_.ID, objc.Sel("producer"))
	return rv
}


// The producer of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/producer

func (c_ CSSearchableItemAttributeSet) SetProducer(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProducer:"), objc.String(value))
}


// The name of the color profile the camera used for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/profileName

func (c_ CSSearchableItemAttributeSet) ProfileName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("profileName"))
	return rv
}


// The name of the color profile the camera used for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/profileName

func (c_ CSSearchableItemAttributeSet) SetProfileName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProfileName:"), objc.String(value))
}


// A list of projects of which this file is a part.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/projects

func (c_ CSSearchableItemAttributeSet) Projects() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("projects"))
	return rv
}


// A list of projects of which this file is a part.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerDataTypeIdentifiers

func (c_ CSSearchableItemAttributeSet) ProviderDataTypeIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("providerDataTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to data representations the delegate provides.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerFileTypeIdentifiers

func (c_ CSSearchableItemAttributeSet) ProviderFileTypeIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("providerFileTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to file representations the delegate provides.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerInPlaceFileTypeIdentifiers

func (c_ CSSearchableItemAttributeSet) ProviderInPlaceFileTypeIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("providerInPlaceFileTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to in-place file representations the delegate provides.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/publishers

func (c_ CSSearchableItemAttributeSet) Publishers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("publishers"))
	return rv
}


// A list of people, organizations, services, or other entities responsible for making the media available.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rankingHint

func (c_ CSSearchableItemAttributeSet) RankingHint() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("rankingHint"))
	return rv
}


// A number that indicates the relative importance of the item among other items from the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rankingHint

func (c_ CSSearchableItemAttributeSet) SetRankingHint(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRankingHint:"), value)
}


// The user-supplied rating of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rating

func (c_ CSSearchableItemAttributeSet) Rating() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("rating"))
	return rv
}


// The user-supplied rating of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rating

func (c_ CSSearchableItemAttributeSet) SetRating(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRating:"), value)
}


// A description of the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/ratingDescription

func (c_ CSSearchableItemAttributeSet) RatingDescription() string {
	rv := objc.Send[string](c_.ID, objc.Sel("ratingDescription"))
	return rv
}


// A description of the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/ratingDescription

func (c_ CSSearchableItemAttributeSet) SetRatingDescription(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRatingDescription:"), objc.String(value))
}


// An array of addresses associated with the recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientAddresses

func (c_ CSSearchableItemAttributeSet) RecipientAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("recipientAddresses"))
	return rv
}


// An array of addresses associated with the recipients of the message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientEmailAddresses

func (c_ CSSearchableItemAttributeSet) RecipientEmailAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("recipientEmailAddresses"))
	return rv
}


// An array of email addresses associated with the recipient.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientNames

func (c_ CSSearchableItemAttributeSet) RecipientNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("recipientNames"))
	return rv
}


// An array of names representing the recipients of this message.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recordingDate

func (c_ CSSearchableItemAttributeSet) RecordingDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("recordingDate"))
	return rv
}


// The recording date of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recordingDate

func (c_ CSSearchableItemAttributeSet) SetRecordingDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordingDate:"), value)
}


// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/redEyeOn

func (c_ CSSearchableItemAttributeSet) RedEyeOn() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("redEyeOn"))
	return rv
}


// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/redEyeOn

func (c_ CSSearchableItemAttributeSet) SetRedEyeOn(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRedEyeOn:"), value)
}


// The unique identifier for the item to which the activity is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/relatedUniqueIdentifier

func (c_ CSSearchableItemAttributeSet) RelatedUniqueIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("relatedUniqueIdentifier"))
	return rv
}


// The unique identifier for the item to which the activity is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/relatedUniqueIdentifier

func (c_ CSSearchableItemAttributeSet) SetRelatedUniqueIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRelatedUniqueIdentifier:"), objc.String(value))
}


// The resolution height of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionHeightDPI

func (c_ CSSearchableItemAttributeSet) ResolutionHeightDPI() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("resolutionHeightDPI"))
	return rv
}


// The resolution height of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionHeightDPI

func (c_ CSSearchableItemAttributeSet) SetResolutionHeightDPI(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionHeightDPI:"), value)
}


// The resolution width of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionWidthDPI

func (c_ CSSearchableItemAttributeSet) ResolutionWidthDPI() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("resolutionWidthDPI"))
	return rv
}


// The resolution width of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionWidthDPI

func (c_ CSSearchableItemAttributeSet) SetResolutionWidthDPI(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionWidthDPI:"), value)
}


// A link to information about the rights held in and over the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rights

func (c_ CSSearchableItemAttributeSet) Rights() string {
	rv := objc.Send[string](c_.ID, objc.Sel("rights"))
	return rv
}


// A link to information about the rights held in and over the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rights

func (c_ CSSearchableItemAttributeSet) SetRights(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRights:"), objc.String(value))
}


// Indicates the role of the content creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/role

func (c_ CSSearchableItemAttributeSet) Role() string {
	rv := objc.Send[string](c_.ID, objc.Sel("role"))
	return rv
}


// Indicates the role of the content creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/role

func (c_ CSSearchableItemAttributeSet) SetRole(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRole:"), objc.String(value))
}


// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/securityMethod

func (c_ CSSearchableItemAttributeSet) SecurityMethod() string {
	rv := objc.Send[string](c_.ID, objc.Sel("securityMethod"))
	return rv
}


// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/securityMethod

func (c_ CSSearchableItemAttributeSet) SetSecurityMethod(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecurityMethod:"), objc.String(value))
}


// The file type of the item to enable the user to share items from Spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/sharedItemContentType

func (c_ CSSearchableItemAttributeSet) SharedItemContentType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("sharedItemContentType"))
	return rv
}


// The file type of the item to enable the user to share items from Spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/sharedItemContentType

func (c_ CSSearchableItemAttributeSet) SetSharedItemContentType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSharedItemContentType:"), value)
}


// The speed of the item, in kilometers per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/speed

func (c_ CSSearchableItemAttributeSet) Speed() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("speed"))
	return rv
}


// The speed of the item, in kilometers per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/speed

func (c_ CSSearchableItemAttributeSet) SetSpeed(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpeed:"), value)
}


// The start date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/startDate

func (c_ CSSearchableItemAttributeSet) StartDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("startDate"))
	return rv
}


// The start date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/startDate

func (c_ CSSearchableItemAttributeSet) SetStartDate(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStartDate:"), value)
}


// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/stateOrProvince

func (c_ CSSearchableItemAttributeSet) StateOrProvince() string {
	rv := objc.Send[string](c_.ID, objc.Sel("stateOrProvince"))
	return rv
}


// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/stateOrProvince

func (c_ CSSearchableItemAttributeSet) SetStateOrProvince(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStateOrProvince:"), objc.String(value))
}


// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/streamable

func (c_ CSSearchableItemAttributeSet) Streamable() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("streamable"))
	return rv
}


// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/streamable

func (c_ CSSearchableItemAttributeSet) SetStreamable(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreamable:"), value)
}


// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subThoroughfare

func (c_ CSSearchableItemAttributeSet) SubThoroughfare() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subThoroughfare"))
	return rv
}


// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subThoroughfare

func (c_ CSSearchableItemAttributeSet) SetSubThoroughfare(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubThoroughfare:"), objc.String(value))
}


// The subject of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subject

func (c_ CSSearchableItemAttributeSet) Subject() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subject"))
	return rv
}


// The subject of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subject

func (c_ CSSearchableItemAttributeSet) SetSubject(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubject:"), objc.String(value))
}


// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation

func (c_ CSSearchableItemAttributeSet) SupportsNavigation() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("supportsNavigation"))
	return rv
}


// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation

func (c_ CSSearchableItemAttributeSet) SetSupportsNavigation(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsNavigation:"), value)
}


// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsPhoneCall

func (c_ CSSearchableItemAttributeSet) SupportsPhoneCall() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("supportsPhoneCall"))
	return rv
}


// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsPhoneCall

func (c_ CSSearchableItemAttributeSet) SetSupportsPhoneCall(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsPhoneCall:"), value)
}


// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/tempo

func (c_ CSSearchableItemAttributeSet) Tempo() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("tempo"))
	return rv
}


// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/tempo

func (c_ CSSearchableItemAttributeSet) SetTempo(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTempo:"), value)
}


// The textual content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContent

func (c_ CSSearchableItemAttributeSet) TextContent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("textContent"))
	return rv
}


// The textual content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContent

func (c_ CSSearchableItemAttributeSet) SetTextContent(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextContent:"), objc.String(value))
}


// A string that presents the Apple Intelligence summarization of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContentSummary

func (c_ CSSearchableItemAttributeSet) TextContentSummary() string {
	rv := objc.Send[string](c_.ID, objc.Sel("textContentSummary"))
	return rv
}


// The theme of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/theme

func (c_ CSSearchableItemAttributeSet) Theme() string {
	rv := objc.Send[string](c_.ID, objc.Sel("theme"))
	return rv
}


// The theme of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/theme

func (c_ CSSearchableItemAttributeSet) SetTheme(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTheme:"), objc.String(value))
}


// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thoroughfare

func (c_ CSSearchableItemAttributeSet) Thoroughfare() string {
	rv := objc.Send[string](c_.ID, objc.Sel("thoroughfare"))
	return rv
}


// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thoroughfare

func (c_ CSSearchableItemAttributeSet) SetThoroughfare(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThoroughfare:"), objc.String(value))
}


// Image data that represents the thumbnail of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailData

func (c_ CSSearchableItemAttributeSet) ThumbnailData() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("thumbnailData"))
	return rv
}


// Image data that represents the thumbnail of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailData

func (c_ CSSearchableItemAttributeSet) SetThumbnailData(value foundation.IData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailData:"), value)
}


// The local file URL of the thumbnail image for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailURL

func (c_ CSSearchableItemAttributeSet) ThumbnailURL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("thumbnailURL"))
	return rv
}


// The local file URL of the thumbnail image for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailURL

func (c_ CSSearchableItemAttributeSet) SetThumbnailURL(value foundation.IURL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailURL:"), value)
}


// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timeSignature

func (c_ CSSearchableItemAttributeSet) TimeSignature() string {
	rv := objc.Send[string](c_.ID, objc.Sel("timeSignature"))
	return rv
}


// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timeSignature

func (c_ CSSearchableItemAttributeSet) SetTimeSignature(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeSignature:"), objc.String(value))
}


// The timestamp on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timestamp

func (c_ CSSearchableItemAttributeSet) Timestamp() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("timestamp"))
	return rv
}


// The timestamp on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timestamp

func (c_ CSSearchableItemAttributeSet) SetTimestamp(value foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimestamp:"), value)
}


// The title of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title

func (c_ CSSearchableItemAttributeSet) Title() string {
	rv := objc.Send[string](c_.ID, objc.Sel("title"))
	return rv
}


// The title of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title

func (c_ CSSearchableItemAttributeSet) SetTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The total bit rate of the media, combining audio and video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/totalBitRate

func (c_ CSSearchableItemAttributeSet) TotalBitRate() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("totalBitRate"))
	return rv
}


// The total bit rate of the media, combining audio and video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/totalBitRate

func (c_ CSSearchableItemAttributeSet) SetTotalBitRate(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalBitRate:"), value)
}


// A string that represents the text the system transcribed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/transcribedTextContent

func (c_ CSSearchableItemAttributeSet) TranscribedTextContent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("transcribedTextContent"))
	return rv
}


// A string that represents the text the system transcribed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/transcribedTextContent

func (c_ CSSearchableItemAttributeSet) SetTranscribedTextContent(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTranscribedTextContent:"), objc.String(value))
}


// The URL associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/url

func (c_ CSSearchableItemAttributeSet) URL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("URL"))
	return rv
}


// The URL associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/url

func (c_ CSSearchableItemAttributeSet) SetURL(value foundation.IURL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setURL:"), value)
}


// A value that indicates the user created the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCreated

func (c_ CSSearchableItemAttributeSet) UserCreated() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("userCreated"))
	return rv
}


// A value that indicates the user created the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCreated

func (c_ CSSearchableItemAttributeSet) SetUserCreated(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCreated:"), value)
}


// A value that indicates the user selected the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCurated

func (c_ CSSearchableItemAttributeSet) UserCurated() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("userCurated"))
	return rv
}


// A value that indicates the user selected the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCurated

func (c_ CSSearchableItemAttributeSet) SetUserCurated(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCurated:"), value)
}


// A value that indicates the user purchased or owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userOwned

func (c_ CSSearchableItemAttributeSet) UserOwned() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("userOwned"))
	return rv
}


// A value that indicates the user purchased or owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userOwned

func (c_ CSSearchableItemAttributeSet) SetUserOwned(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserOwned:"), value)
}


// A version string associated with the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/version

func (c_ CSSearchableItemAttributeSet) Version() string {
	rv := objc.Send[string](c_.ID, objc.Sel("version"))
	return rv
}


// A version string associated with the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/version

func (c_ CSSearchableItemAttributeSet) SetVersion(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVersion:"), objc.String(value))
}


// The video bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/videoBitRate

func (c_ CSSearchableItemAttributeSet) VideoBitRate() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("videoBitRate"))
	return rv
}


// The video bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/videoBitRate

func (c_ CSSearchableItemAttributeSet) SetVideoBitRate(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoBitRate:"), value)
}


// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/weakRelatedUniqueIdentifier

func (c_ CSSearchableItemAttributeSet) WeakRelatedUniqueIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("weakRelatedUniqueIdentifier"))
	return rv
}


// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/weakRelatedUniqueIdentifier

func (c_ CSSearchableItemAttributeSet) SetWeakRelatedUniqueIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeakRelatedUniqueIdentifier:"), objc.String(value))
}


// The white balance setting when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/whiteBalance

func (c_ CSSearchableItemAttributeSet) WhiteBalance() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("whiteBalance"))
	return rv
}


// The white balance setting when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/whiteBalance

func (c_ CSSearchableItemAttributeSet) SetWhiteBalance(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteBalance:"), value)
}


// A key that specifies the action’s identifier in a user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csactionidentifier

func (c_ CSSearchableItemAttributeSet) CSActionIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CSActionIdentifier"))
	return rv
}


// The composer of the song or audio composition that the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/composer

func (c_ CSSearchableItemAttributeSet) Composer() string {
	rv := objc.Send[string](c_.ID, objc.Sel("composer"))
	return rv
}


// The composer of the song or audio composition that the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/composer

func (c_ CSSearchableItemAttributeSet) SetComposer(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComposer:"), objc.String(value))
}


