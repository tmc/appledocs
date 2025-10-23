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
	ContentModificationDate() foundation.NSDate
	SetContentModificationDate(value foundation.NSDate)
	DisplayName() string
	SetDisplayName(value string)
	EncodingApplications() []string
	SetEncodingApplications(value []string)
	EndDate() foundation.NSDate
	SetEndDate(value foundation.NSDate)
	GPSDOP() foundation.Number
	SetGPSDOP(value foundation.Number)
	SupportsNavigation() foundation.Number
	SetSupportsNavigation(value foundation.Number)
	Title() string
	SetTitle(value string)
	CSActionIdentifier() string
	AccountHandles() string
	SetAccountHandles(value string)
	AccountIdentifier() string
	SetAccountIdentifier(value string)
	AcquisitionMake() string
	SetAcquisitionMake(value string)
	AcquisitionModel() string
	SetAcquisitionModel(value string)
	ActionIdentifiers() string
	SetActionIdentifiers(value string)
	AddedDate() foundation.Date
	SetAddedDate(value foundation.Date)
	AdditionalRecipients() ICSPerson
	SetAdditionalRecipients(value ICSPerson)
	Album() string
	SetAlbum(value string)
	AllDay() foundation.Number
	SetAllDay(value foundation.Number)
	AlternateNames() string
	SetAlternateNames(value string)
	Altitude() foundation.Number
	SetAltitude(value foundation.Number)
	Aperture() foundation.Number
	SetAperture(value foundation.Number)
	Artist() string
	SetArtist(value string)
	Audiences() string
	SetAudiences(value string)
	AudioBitRate() foundation.Number
	SetAudioBitRate(value foundation.Number)
	AudioChannelCount() foundation.Number
	SetAudioChannelCount(value foundation.Number)
	AudioEncodingApplication() string
	SetAudioEncodingApplication(value string)
	AudioSampleRate() foundation.Number
	SetAudioSampleRate(value foundation.Number)
	AudioTrackNumber() foundation.Number
	SetAudioTrackNumber(value foundation.Number)
	AuthorAddresses() string
	SetAuthorAddresses(value string)
	AuthorEmailAddresses() string
	SetAuthorEmailAddresses(value string)
	AuthorNames() string
	SetAuthorNames(value string)
	Authors() ICSPerson
	SetAuthors(value ICSPerson)
	BitsPerSample() foundation.Number
	SetBitsPerSample(value foundation.Number)
	CameraOwner() string
	SetCameraOwner(value string)
	City() string
	SetCity(value string)
	Codecs() string
	SetCodecs(value string)
	ColorSpace() string
	SetColorSpace(value string)
	Comment() string
	SetComment(value string)
	CompletionDate() foundation.Date
	SetCompletionDate(value foundation.Date)
	Composer() string
	SetComposer(value string)
	ContactKeywords() string
	SetContactKeywords(value string)
	ContainerDisplayName() string
	SetContainerDisplayName(value string)
	ContainerIdentifier() string
	SetContainerIdentifier(value string)
	ContainerOrder() foundation.Number
	SetContainerOrder(value foundation.Number)
	ContainerTitle() string
	SetContainerTitle(value string)
	ContentCreationDate() foundation.Date
	SetContentCreationDate(value foundation.Date)
	ContentDescription() string
	SetContentDescription(value string)
	ContentRating() foundation.Number
	SetContentRating(value foundation.Number)
	ContentSources() string
	SetContentSources(value string)
	ContentType() string
	SetContentType(value string)
	ContentTypeTree() string
	SetContentTypeTree(value string)
	ContentURL() foundation.URL
	SetContentURL(value foundation.URL)
	Contributors() string
	SetContributors(value string)
	Copyright() string
	SetCopyright(value string)
	Country() string
	SetCountry(value string)
	Coverage() string
	SetCoverage(value string)
	Creator() string
	SetCreator(value string)
	DarkThumbnailURL() foundation.URL
	SetDarkThumbnailURL(value foundation.URL)
	DeliveryType() foundation.Number
	SetDeliveryType(value foundation.Number)
	Director() string
	SetDirector(value string)
	DomainIdentifier() string
	SetDomainIdentifier(value string)
	DownloadedDate() foundation.Date
	SetDownloadedDate(value foundation.Date)
	DueDate() foundation.Date
	SetDueDate(value foundation.Date)
	Duration() foundation.Number
	SetDuration(value foundation.Number)
	Editors() string
	SetEditors(value string)
	EmailAddresses() string
	SetEmailAddresses(value string)
	EmailHeaders() string
	SetEmailHeaders(value string)
	ExifVersion() string
	SetExifVersion(value string)
	ExifgpsVersion() string
	SetExifgpsVersion(value string)
	ExposureMode() foundation.Number
	SetExposureMode(value foundation.Number)
	ExposureProgram() string
	SetExposureProgram(value string)
	ExposureTime() foundation.Number
	SetExposureTime(value foundation.Number)
	ExposureTimeString() string
	SetExposureTimeString(value string)
	FNumber() foundation.Number
	SetFNumber(value foundation.Number)
	FileSize() foundation.Number
	SetFileSize(value foundation.Number)
	FlashOn() foundation.Number
	SetFlashOn(value foundation.Number)
	FocalLength() foundation.Number
	SetFocalLength(value foundation.Number)
	FocalLength35mm() foundation.Number
	SetFocalLength35mm(value foundation.Number)
	FontNames() string
	SetFontNames(value string)
	FullyFormattedAddress() string
	SetFullyFormattedAddress(value string)
	GeneralMIDISequence() foundation.Number
	SetGeneralMIDISequence(value foundation.Number)
	Genre() string
	SetGenre(value string)
	GpsAreaInformation() string
	SetGpsAreaInformation(value string)
	GpsDateStamp() foundation.Date
	SetGpsDateStamp(value foundation.Date)
	GpsDestBearing() foundation.Number
	SetGpsDestBearing(value foundation.Number)
	GpsDestDistance() foundation.Number
	SetGpsDestDistance(value foundation.Number)
	GpsDestLatitude() foundation.Number
	SetGpsDestLatitude(value foundation.Number)
	GpsDestLongitude() foundation.Number
	SetGpsDestLongitude(value foundation.Number)
	GpsDifferental() foundation.Number
	SetGpsDifferental(value foundation.Number)
	GpsMapDatum() string
	SetGpsMapDatum(value string)
	GpsMeasureMode() string
	SetGpsMeasureMode(value string)
	GpsProcessingMethod() string
	SetGpsProcessingMethod(value string)
	GpsStatus() string
	SetGpsStatus(value string)
	GpsTrack() foundation.Number
	SetGpsTrack(value foundation.Number)
	HasAlphaChannel() foundation.Number
	SetHasAlphaChannel(value foundation.Number)
	Headline() string
	SetHeadline(value string)
	HiddenAdditionalRecipients() ICSPerson
	SetHiddenAdditionalRecipients(value ICSPerson)
	HtmlContentData() foundation.Data
	SetHtmlContentData(value foundation.Data)
	Identifier() string
	SetIdentifier(value string)
	ImageDirection() foundation.Number
	SetImageDirection(value foundation.Number)
	ImportantDates() foundation.Date
	SetImportantDates(value foundation.Date)
	Information() string
	SetInformation(value string)
	InstantMessageAddresses() string
	SetInstantMessageAddresses(value string)
	Instructions() string
	SetInstructions(value string)
	IsPriority() foundation.Number
	SetIsPriority(value foundation.Number)
	IsoSpeed() foundation.Number
	SetIsoSpeed(value foundation.Number)
	KeySignature() string
	SetKeySignature(value string)
	Keywords() string
	SetKeywords(value string)
	Kind() string
	SetKind(value string)
	Languages() string
	SetLanguages(value string)
	LastUsedDate() foundation.Date
	SetLastUsedDate(value foundation.Date)
	Latitude() foundation.Number
	SetLatitude(value foundation.Number)
	LayerNames() string
	SetLayerNames(value string)
	LensModel() string
	SetLensModel(value string)
	LikelyJunk() foundation.Number
	SetLikelyJunk(value foundation.Number)
	Local() foundation.Number
	SetLocal(value foundation.Number)
	Longitude() foundation.Number
	SetLongitude(value foundation.Number)
	Lyricist() string
	SetLyricist(value string)
	MailboxIdentifiers() string
	SetMailboxIdentifiers(value string)
	MaxAperture() foundation.Number
	SetMaxAperture(value foundation.Number)
	MediaTypes() string
	SetMediaTypes(value string)
	MetadataModificationDate() foundation.Date
	SetMetadataModificationDate(value foundation.Date)
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
	Organizations() string
	SetOrganizations(value string)
	Orientation() foundation.Number
	SetOrientation(value foundation.Number)
	OriginalFormat() string
	SetOriginalFormat(value string)
	OriginalSource() string
	SetOriginalSource(value string)
	PageCount() foundation.Number
	SetPageCount(value foundation.Number)
	PageHeight() foundation.Number
	SetPageHeight(value foundation.Number)
	PageWidth() foundation.Number
	SetPageWidth(value foundation.Number)
	Participants() string
	SetParticipants(value string)
	Path() string
	SetPath(value string)
	Performers() string
	SetPerformers(value string)
	PhoneNumbers() string
	SetPhoneNumbers(value string)
	PixelCount() foundation.Number
	SetPixelCount(value foundation.Number)
	PixelHeight() foundation.Number
	SetPixelHeight(value foundation.Number)
	PixelWidth() foundation.Number
	SetPixelWidth(value foundation.Number)
	PlayCount() foundation.Number
	SetPlayCount(value foundation.Number)
	PostalCode() string
	SetPostalCode(value string)
	PrimaryRecipients() ICSPerson
	SetPrimaryRecipients(value ICSPerson)
	Producer() string
	SetProducer(value string)
	ProfileName() string
	SetProfileName(value string)
	Projects() string
	SetProjects(value string)
	ProviderDataTypeIdentifiers() string
	SetProviderDataTypeIdentifiers(value string)
	ProviderFileTypeIdentifiers() string
	SetProviderFileTypeIdentifiers(value string)
	ProviderInPlaceFileTypeIdentifiers() string
	SetProviderInPlaceFileTypeIdentifiers(value string)
	Publishers() string
	SetPublishers(value string)
	RankingHint() foundation.Number
	SetRankingHint(value foundation.Number)
	Rating() foundation.Number
	SetRating(value foundation.Number)
	RatingDescription() string
	SetRatingDescription(value string)
	RecipientAddresses() string
	SetRecipientAddresses(value string)
	RecipientEmailAddresses() string
	SetRecipientEmailAddresses(value string)
	RecipientNames() string
	SetRecipientNames(value string)
	RecordingDate() foundation.Date
	SetRecordingDate(value foundation.Date)
	RedEyeOn() foundation.Number
	SetRedEyeOn(value foundation.Number)
	RelatedUniqueIdentifier() string
	SetRelatedUniqueIdentifier(value string)
	ResolutionHeightDPI() foundation.Number
	SetResolutionHeightDPI(value foundation.Number)
	ResolutionWidthDPI() foundation.Number
	SetResolutionWidthDPI(value foundation.Number)
	Rights() string
	SetRights(value string)
	Role() string
	SetRole(value string)
	SecurityMethod() string
	SetSecurityMethod(value string)
	SharedItemContentType() objectivec.IObject
	SetSharedItemContentType(value objectivec.IObject)
	Speed() foundation.Number
	SetSpeed(value foundation.Number)
	StartDate() foundation.Date
	SetStartDate(value foundation.Date)
	StateOrProvince() string
	SetStateOrProvince(value string)
	Streamable() foundation.Number
	SetStreamable(value foundation.Number)
	SubThoroughfare() string
	SetSubThoroughfare(value string)
	Subject() string
	SetSubject(value string)
	SupportsPhoneCall() foundation.Number
	SetSupportsPhoneCall(value foundation.Number)
	Tempo() foundation.Number
	SetTempo(value foundation.Number)
	TextContent() string
	SetTextContent(value string)
	TextContentSummary() string
	SetTextContentSummary(value string)
	Theme() string
	SetTheme(value string)
	Thoroughfare() string
	SetThoroughfare(value string)
	ThumbnailData() foundation.Data
	SetThumbnailData(value foundation.Data)
	ThumbnailURL() foundation.URL
	SetThumbnailURL(value foundation.URL)
	TimeSignature() string
	SetTimeSignature(value string)
	Timestamp() foundation.Date
	SetTimestamp(value foundation.Date)
	TotalBitRate() foundation.Number
	SetTotalBitRate(value foundation.Number)
	TranscribedTextContent() string
	SetTranscribedTextContent(value string)
	Url() foundation.URL
	SetUrl(value foundation.URL)
	UserCreated() foundation.Number
	SetUserCreated(value foundation.Number)
	UserCurated() foundation.Number
	SetUserCurated(value foundation.Number)
	UserOwned() foundation.Number
	SetUserOwned(value foundation.Number)
	Version() string
	SetVersion(value string)
	VideoBitRate() foundation.Number
	SetVideoBitRate(value foundation.Number)
	WeakRelatedUniqueIdentifier() string
	SetWeakRelatedUniqueIdentifier(value string)
	WhiteBalance() foundation.Number
	SetWhiteBalance(value foundation.Number)
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
func (c_ CSSearchableItemAttributeSet) SetContentModificationDate(value foundation.NSDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentModificationDate:"), value)
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
func (c_ CSSearchableItemAttributeSet) SetEndDate(value foundation.NSDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEndDate:"), value)
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
func (c_ CSSearchableItemAttributeSet) SetGPSDOP(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDOP:"), value)
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
func (c_ CSSearchableItemAttributeSet) SetSupportsNavigation(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsNavigation:"), value)
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


// A key that specifies the action’s identifier in a user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csactionidentifier
func (c_ CSSearchableItemAttributeSet) CSActionIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CSActionIdentifier"))
	return rv
}


// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accounthandles
func (c_ CSSearchableItemAttributeSet) AccountHandles() string {
	rv := objc.Send[string](c_.ID, objc.Sel("accountHandles"))
	return rv
}


// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accounthandles
func (c_ CSSearchableItemAttributeSet) SetAccountHandles(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccountHandles:"), objc.String(value))
}


// The unique identifier for the account with which the message is associated, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accountidentifier
func (c_ CSSearchableItemAttributeSet) AccountIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("accountIdentifier"))
	return rv
}


// The unique identifier for the account with which the message is associated, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accountidentifier
func (c_ CSSearchableItemAttributeSet) SetAccountIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccountIdentifier:"), objc.String(value))
}


// The manufacturer of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmake
func (c_ CSSearchableItemAttributeSet) AcquisitionMake() string {
	rv := objc.Send[string](c_.ID, objc.Sel("acquisitionMake"))
	return rv
}


// The manufacturer of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmake
func (c_ CSSearchableItemAttributeSet) SetAcquisitionMake(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionMake:"), objc.String(value))
}


// The model of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmodel
func (c_ CSSearchableItemAttributeSet) AcquisitionModel() string {
	rv := objc.Send[string](c_.ID, objc.Sel("acquisitionModel"))
	return rv
}


// The model of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmodel
func (c_ CSSearchableItemAttributeSet) SetAcquisitionModel(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionModel:"), objc.String(value))
}


// The identifiers that specify custom actions the app supports for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/actionidentifiers
func (c_ CSSearchableItemAttributeSet) ActionIdentifiers() string {
	rv := objc.Send[string](c_.ID, objc.Sel("actionIdentifiers"))
	return rv
}


// The identifiers that specify custom actions the app supports for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/actionidentifiers
func (c_ CSSearchableItemAttributeSet) SetActionIdentifiers(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActionIdentifiers:"), objc.String(value))
}


// The date on which the item was moved into its current location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/addeddate
func (c_ CSSearchableItemAttributeSet) AddedDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("addedDate"))
	return rv
}


// The date on which the item was moved into its current location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/addeddate
func (c_ CSSearchableItemAttributeSet) SetAddedDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAddedDate:"), value)
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/additionalrecipients
func (c_ CSSearchableItemAttributeSet) AdditionalRecipients() ICSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("additionalRecipients"))
	return rv
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/additionalrecipients
func (c_ CSSearchableItemAttributeSet) SetAdditionalRecipients(value ICSPerson) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAdditionalRecipients:"), value)
}


// The title for a collection of audio media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/album
func (c_ CSSearchableItemAttributeSet) Album() string {
	rv := objc.Send[string](c_.ID, objc.Sel("album"))
	return rv
}


// The title for a collection of audio media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/album
func (c_ CSSearchableItemAttributeSet) SetAlbum(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlbum:"), objc.String(value))
}


// A value that indicates if the event covers an entire day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/allday
func (c_ CSSearchableItemAttributeSet) AllDay() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("allDay"))
	return rv
}


// A value that indicates if the event covers an entire day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/allday
func (c_ CSSearchableItemAttributeSet) SetAllDay(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllDay:"), value)
}


// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/alternatenames
func (c_ CSSearchableItemAttributeSet) AlternateNames() string {
	rv := objc.Send[string](c_.ID, objc.Sel("alternateNames"))
	return rv
}


// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/alternatenames
func (c_ CSSearchableItemAttributeSet) SetAlternateNames(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlternateNames:"), objc.String(value))
}


// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/altitude
func (c_ CSSearchableItemAttributeSet) Altitude() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("altitude"))
	return rv
}


// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/altitude
func (c_ CSSearchableItemAttributeSet) SetAltitude(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAltitude:"), value)
}


// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/aperture
func (c_ CSSearchableItemAttributeSet) Aperture() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("aperture"))
	return rv
}


// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/aperture
func (c_ CSSearchableItemAttributeSet) SetAperture(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAperture:"), value)
}


// The artist associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/artist
func (c_ CSSearchableItemAttributeSet) Artist() string {
	rv := objc.Send[string](c_.ID, objc.Sel("artist"))
	return rv
}


// The artist associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/artist
func (c_ CSSearchableItemAttributeSet) SetArtist(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setArtist:"), objc.String(value))
}


// A class of entity for which the item is intended or useful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiences
func (c_ CSSearchableItemAttributeSet) Audiences() string {
	rv := objc.Send[string](c_.ID, objc.Sel("audiences"))
	return rv
}


// A class of entity for which the item is intended or useful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiences
func (c_ CSSearchableItemAttributeSet) SetAudiences(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudiences:"), objc.String(value))
}


// The audio bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiobitrate
func (c_ CSSearchableItemAttributeSet) AudioBitRate() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioBitRate"))
	return rv
}


// The audio bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiobitrate
func (c_ CSSearchableItemAttributeSet) SetAudioBitRate(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioBitRate:"), value)
}


// The number of channels in the audio data that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiochannelcount
func (c_ CSSearchableItemAttributeSet) AudioChannelCount() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioChannelCount"))
	return rv
}


// The number of channels in the audio data that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiochannelcount
func (c_ CSSearchableItemAttributeSet) SetAudioChannelCount(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioChannelCount:"), value)
}


// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audioencodingapplication
func (c_ CSSearchableItemAttributeSet) AudioEncodingApplication() string {
	rv := objc.Send[string](c_.ID, objc.Sel("audioEncodingApplication"))
	return rv
}


// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audioencodingapplication
func (c_ CSSearchableItemAttributeSet) SetAudioEncodingApplication(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioEncodingApplication:"), objc.String(value))
}


// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiosamplerate
func (c_ CSSearchableItemAttributeSet) AudioSampleRate() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioSampleRate"))
	return rv
}


// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiosamplerate
func (c_ CSSearchableItemAttributeSet) SetAudioSampleRate(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSampleRate:"), value)
}


// The track number of a song or audio composition when part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiotracknumber
func (c_ CSSearchableItemAttributeSet) AudioTrackNumber() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioTrackNumber"))
	return rv
}


// The track number of a song or audio composition when part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiotracknumber
func (c_ CSSearchableItemAttributeSet) SetAudioTrackNumber(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioTrackNumber:"), value)
}


// An array of addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoraddresses
func (c_ CSSearchableItemAttributeSet) AuthorAddresses() string {
	rv := objc.Send[string](c_.ID, objc.Sel("authorAddresses"))
	return rv
}


// An array of addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoraddresses
func (c_ CSSearchableItemAttributeSet) SetAuthorAddresses(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorAddresses:"), objc.String(value))
}


// An array of email addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoremailaddresses
func (c_ CSSearchableItemAttributeSet) AuthorEmailAddresses() string {
	rv := objc.Send[string](c_.ID, objc.Sel("authorEmailAddresses"))
	return rv
}


// An array of email addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoremailaddresses
func (c_ CSSearchableItemAttributeSet) SetAuthorEmailAddresses(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorEmailAddresses:"), objc.String(value))
}


// An array of names representing the authors who have worked on the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authornames
func (c_ CSSearchableItemAttributeSet) AuthorNames() string {
	rv := objc.Send[string](c_.ID, objc.Sel("authorNames"))
	return rv
}


// An array of names representing the authors who have worked on the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authornames
func (c_ CSSearchableItemAttributeSet) SetAuthorNames(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorNames:"), objc.String(value))
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authors
func (c_ CSSearchableItemAttributeSet) Authors() ICSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("authors"))
	return rv
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authors
func (c_ CSSearchableItemAttributeSet) SetAuthors(value ICSPerson) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthors:"), value)
}


// The number of bits per sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/bitspersample
func (c_ CSSearchableItemAttributeSet) BitsPerSample() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("bitsPerSample"))
	return rv
}


// The number of bits per sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/bitspersample
func (c_ CSSearchableItemAttributeSet) SetBitsPerSample(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBitsPerSample:"), value)
}


// The owner of the camera that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/cameraowner
func (c_ CSSearchableItemAttributeSet) CameraOwner() string {
	rv := objc.Send[string](c_.ID, objc.Sel("cameraOwner"))
	return rv
}


// The owner of the camera that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/cameraowner
func (c_ CSSearchableItemAttributeSet) SetCameraOwner(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraOwner:"), objc.String(value))
}


// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/city
func (c_ CSSearchableItemAttributeSet) City() string {
	rv := objc.Send[string](c_.ID, objc.Sel("city"))
	return rv
}


// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/city
func (c_ CSSearchableItemAttributeSet) SetCity(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), objc.String(value))
}


// The codecs used to encode/decode the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/codecs
func (c_ CSSearchableItemAttributeSet) Codecs() string {
	rv := objc.Send[string](c_.ID, objc.Sel("codecs"))
	return rv
}


// The codecs used to encode/decode the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/codecs
func (c_ CSSearchableItemAttributeSet) SetCodecs(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCodecs:"), objc.String(value))
}


// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/colorspace
func (c_ CSSearchableItemAttributeSet) ColorSpace() string {
	rv := objc.Send[string](c_.ID, objc.Sel("colorSpace"))
	return rv
}


// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/colorspace
func (c_ CSSearchableItemAttributeSet) SetColorSpace(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSpace:"), objc.String(value))
}


// A comment related to the media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/comment
func (c_ CSSearchableItemAttributeSet) Comment() string {
	rv := objc.Send[string](c_.ID, objc.Sel("comment"))
	return rv
}


// A comment related to the media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/comment
func (c_ CSSearchableItemAttributeSet) SetComment(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComment:"), objc.String(value))
}


// The date on which the item was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/completiondate
func (c_ CSSearchableItemAttributeSet) CompletionDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("completionDate"))
	return rv
}


// The date on which the item was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/completiondate
func (c_ CSSearchableItemAttributeSet) SetCompletionDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionDate:"), value)
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


// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contactkeywords
func (c_ CSSearchableItemAttributeSet) ContactKeywords() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contactKeywords"))
	return rv
}


// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contactkeywords
func (c_ CSSearchableItemAttributeSet) SetContactKeywords(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactKeywords:"), objc.String(value))
}


// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerdisplayname
func (c_ CSSearchableItemAttributeSet) ContainerDisplayName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerDisplayName"))
	return rv
}


// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerdisplayname
func (c_ CSSearchableItemAttributeSet) SetContainerDisplayName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerDisplayName:"), objc.String(value))
}


// The identifier of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containeridentifier
func (c_ CSSearchableItemAttributeSet) ContainerIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}


// The identifier of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containeridentifier
func (c_ CSSearchableItemAttributeSet) SetContainerIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), objc.String(value))
}


// The order of the item within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerorder
func (c_ CSSearchableItemAttributeSet) ContainerOrder() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("containerOrder"))
	return rv
}


// The order of the item within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerorder
func (c_ CSSearchableItemAttributeSet) SetContainerOrder(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerOrder:"), value)
}


// The title of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containertitle
func (c_ CSSearchableItemAttributeSet) ContainerTitle() string {
	rv := objc.Send[string](c_.ID, objc.Sel("containerTitle"))
	return rv
}


// The title of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containertitle
func (c_ CSSearchableItemAttributeSet) SetContainerTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerTitle:"), objc.String(value))
}


// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentcreationdate
func (c_ CSSearchableItemAttributeSet) ContentCreationDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("contentCreationDate"))
	return rv
}


// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentcreationdate
func (c_ CSSearchableItemAttributeSet) SetContentCreationDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentCreationDate:"), value)
}


// A description of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentdescription
func (c_ CSSearchableItemAttributeSet) ContentDescription() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contentDescription"))
	return rv
}


// A description of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentdescription
func (c_ CSSearchableItemAttributeSet) SetContentDescription(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentDescription:"), objc.String(value))
}


// A value that indicates if the media contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentrating
func (c_ CSSearchableItemAttributeSet) ContentRating() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("contentRating"))
	return rv
}


// A value that indicates if the media contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentrating
func (c_ CSSearchableItemAttributeSet) SetContentRating(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentRating:"), value)
}


// An array of sources from which the media was obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentsources
func (c_ CSSearchableItemAttributeSet) ContentSources() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contentSources"))
	return rv
}


// An array of sources from which the media was obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentsources
func (c_ CSSearchableItemAttributeSet) SetContentSources(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentSources:"), objc.String(value))
}


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttype
func (c_ CSSearchableItemAttributeSet) ContentType() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contentType"))
	return rv
}


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttype
func (c_ CSSearchableItemAttributeSet) SetContentType(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentType:"), objc.String(value))
}


// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttypetree
func (c_ CSSearchableItemAttributeSet) ContentTypeTree() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contentTypeTree"))
	return rv
}


// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttypetree
func (c_ CSSearchableItemAttributeSet) SetContentTypeTree(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentTypeTree:"), objc.String(value))
}


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenturl
func (c_ CSSearchableItemAttributeSet) ContentURL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("contentURL"))
	return rv
}


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenturl
func (c_ CSSearchableItemAttributeSet) SetContentURL(value foundation.URL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentURL:"), value)
}


// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contributors
func (c_ CSSearchableItemAttributeSet) Contributors() string {
	rv := objc.Send[string](c_.ID, objc.Sel("contributors"))
	return rv
}


// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contributors
func (c_ CSSearchableItemAttributeSet) SetContributors(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContributors:"), objc.String(value))
}


// The copyright date of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/copyright
func (c_ CSSearchableItemAttributeSet) Copyright() string {
	rv := objc.Send[string](c_.ID, objc.Sel("copyright"))
	return rv
}


// The copyright date of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/copyright
func (c_ CSSearchableItemAttributeSet) SetCopyright(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCopyright:"), objc.String(value))
}


// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/country
func (c_ CSSearchableItemAttributeSet) Country() string {
	rv := objc.Send[string](c_.ID, objc.Sel("country"))
	return rv
}


// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/country
func (c_ CSSearchableItemAttributeSet) SetCountry(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), objc.String(value))
}


// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/coverage
func (c_ CSSearchableItemAttributeSet) Coverage() string {
	rv := objc.Send[string](c_.ID, objc.Sel("coverage"))
	return rv
}


// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/coverage
func (c_ CSSearchableItemAttributeSet) SetCoverage(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCoverage:"), objc.String(value))
}


// The name of the app that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/creator
func (c_ CSSearchableItemAttributeSet) Creator() string {
	rv := objc.Send[string](c_.ID, objc.Sel("creator"))
	return rv
}


// The name of the app that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/creator
func (c_ CSSearchableItemAttributeSet) SetCreator(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreator:"), objc.String(value))
}


// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/darkthumbnailurl
func (c_ CSSearchableItemAttributeSet) DarkThumbnailURL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("darkThumbnailURL"))
	return rv
}


// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/darkthumbnailurl
func (c_ CSSearchableItemAttributeSet) SetDarkThumbnailURL(value foundation.URL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDarkThumbnailURL:"), value)
}


// The delivery type of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/deliverytype
func (c_ CSSearchableItemAttributeSet) DeliveryType() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("deliveryType"))
	return rv
}


// The delivery type of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/deliverytype
func (c_ CSSearchableItemAttributeSet) SetDeliveryType(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeliveryType:"), value)
}


// The name of the director of the media (for example, a movie director).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/director
func (c_ CSSearchableItemAttributeSet) Director() string {
	rv := objc.Send[string](c_.ID, objc.Sel("director"))
	return rv
}


// The name of the director of the media (for example, a movie director).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/director
func (c_ CSSearchableItemAttributeSet) SetDirector(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDirector:"), objc.String(value))
}


// An identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/domainidentifier
func (c_ CSSearchableItemAttributeSet) DomainIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("domainIdentifier"))
	return rv
}


// An identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/domainidentifier
func (c_ CSSearchableItemAttributeSet) SetDomainIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDomainIdentifier:"), objc.String(value))
}


// The most recent date on which the file was downloaded or received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/downloadeddate
func (c_ CSSearchableItemAttributeSet) DownloadedDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("downloadedDate"))
	return rv
}


// The most recent date on which the file was downloaded or received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/downloadeddate
func (c_ CSSearchableItemAttributeSet) SetDownloadedDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDownloadedDate:"), value)
}


// The date on which the item is due.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duedate
func (c_ CSSearchableItemAttributeSet) DueDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("dueDate"))
	return rv
}


// The date on which the item is due.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duedate
func (c_ CSSearchableItemAttributeSet) SetDueDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDueDate:"), value)
}


// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duration
func (c_ CSSearchableItemAttributeSet) Duration() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("duration"))
	return rv
}


// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duration
func (c_ CSSearchableItemAttributeSet) SetDuration(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDuration:"), value)
}


// A list of editors who have worked on the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/editors
func (c_ CSSearchableItemAttributeSet) Editors() string {
	rv := objc.Send[string](c_.ID, objc.Sel("editors"))
	return rv
}


// A list of editors who have worked on the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/editors
func (c_ CSSearchableItemAttributeSet) SetEditors(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEditors:"), objc.String(value))
}


// An array of email addresses associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailaddresses
func (c_ CSSearchableItemAttributeSet) EmailAddresses() string {
	rv := objc.Send[string](c_.ID, objc.Sel("emailAddresses"))
	return rv
}


// An array of email addresses associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailaddresses
func (c_ CSSearchableItemAttributeSet) SetEmailAddresses(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailAddresses:"), objc.String(value))
}


// A dictionary that contains all the headers of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailheaders
func (c_ CSSearchableItemAttributeSet) EmailHeaders() string {
	rv := objc.Send[string](c_.ID, objc.Sel("emailHeaders"))
	return rv
}


// A dictionary that contains all the headers of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailheaders
func (c_ CSSearchableItemAttributeSet) SetEmailHeaders(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailHeaders:"), objc.String(value))
}


// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifversion
func (c_ CSSearchableItemAttributeSet) ExifVersion() string {
	rv := objc.Send[string](c_.ID, objc.Sel("exifVersion"))
	return rv
}


// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifversion
func (c_ CSSearchableItemAttributeSet) SetExifVersion(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExifVersion:"), objc.String(value))
}


// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifgpsversion
func (c_ CSSearchableItemAttributeSet) ExifgpsVersion() string {
	rv := objc.Send[string](c_.ID, objc.Sel("exifgpsVersion"))
	return rv
}


// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifgpsversion
func (c_ CSSearchableItemAttributeSet) SetExifgpsVersion(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExifgpsVersion:"), objc.String(value))
}


// The mode the camera used for the exposure of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuremode
func (c_ CSSearchableItemAttributeSet) ExposureMode() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("exposureMode"))
	return rv
}


// The mode the camera used for the exposure of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuremode
func (c_ CSSearchableItemAttributeSet) SetExposureMode(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureMode:"), value)
}


// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposureprogram
func (c_ CSSearchableItemAttributeSet) ExposureProgram() string {
	rv := objc.Send[string](c_.ID, objc.Sel("exposureProgram"))
	return rv
}


// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposureprogram
func (c_ CSSearchableItemAttributeSet) SetExposureProgram(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureProgram:"), objc.String(value))
}


// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretime
func (c_ CSSearchableItemAttributeSet) ExposureTime() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("exposureTime"))
	return rv
}


// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretime
func (c_ CSSearchableItemAttributeSet) SetExposureTime(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTime:"), value)
}


// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretimestring
func (c_ CSSearchableItemAttributeSet) ExposureTimeString() string {
	rv := objc.Send[string](c_.ID, objc.Sel("exposureTimeString"))
	return rv
}


// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretimestring
func (c_ CSSearchableItemAttributeSet) SetExposureTimeString(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTimeString:"), objc.String(value))
}


// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fnumber
func (c_ CSSearchableItemAttributeSet) FNumber() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("fNumber"))
	return rv
}


// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fnumber
func (c_ CSSearchableItemAttributeSet) SetFNumber(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFNumber:"), value)
}


// The size of the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/filesize
func (c_ CSSearchableItemAttributeSet) FileSize() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("fileSize"))
	return rv
}


// The size of the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/filesize
func (c_ CSSearchableItemAttributeSet) SetFileSize(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFileSize:"), value)
}


// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/flashon
func (c_ CSSearchableItemAttributeSet) FlashOn() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("flashOn"))
	return rv
}


// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/flashon
func (c_ CSSearchableItemAttributeSet) SetFlashOn(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashOn:"), value)
}


// The actual focal length of the lens, in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength
func (c_ CSSearchableItemAttributeSet) FocalLength() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("focalLength"))
	return rv
}


// The actual focal length of the lens, in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength
func (c_ CSSearchableItemAttributeSet) SetFocalLength(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength:"), value)
}


// A value that indicates if the focal length is 35mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength35mm
func (c_ CSSearchableItemAttributeSet) FocalLength35mm() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("focalLength35mm"))
	return rv
}


// A value that indicates if the focal length is 35mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength35mm
func (c_ CSSearchableItemAttributeSet) SetFocalLength35mm(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength35mm:"), value)
}


// An array of font names the document uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fontnames
func (c_ CSSearchableItemAttributeSet) FontNames() string {
	rv := objc.Send[string](c_.ID, objc.Sel("fontNames"))
	return rv
}


// An array of font names the document uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fontnames
func (c_ CSSearchableItemAttributeSet) SetFontNames(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFontNames:"), objc.String(value))
}


// The fully formatted address of the item, received from MapKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fullyformattedaddress
func (c_ CSSearchableItemAttributeSet) FullyFormattedAddress() string {
	rv := objc.Send[string](c_.ID, objc.Sel("fullyFormattedAddress"))
	return rv
}


// The fully formatted address of the item, received from MapKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fullyformattedaddress
func (c_ CSSearchableItemAttributeSet) SetFullyFormattedAddress(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFullyFormattedAddress:"), objc.String(value))
}


// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/generalmidisequence
func (c_ CSSearchableItemAttributeSet) GeneralMIDISequence() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("generalMIDISequence"))
	return rv
}


// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/generalmidisequence
func (c_ CSSearchableItemAttributeSet) SetGeneralMIDISequence(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGeneralMIDISequence:"), value)
}


// The genre of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/genre
func (c_ CSSearchableItemAttributeSet) Genre() string {
	rv := objc.Send[string](c_.ID, objc.Sel("genre"))
	return rv
}


// The genre of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/genre
func (c_ CSSearchableItemAttributeSet) SetGenre(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGenre:"), objc.String(value))
}


// Information about the GPS area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsareainformation
func (c_ CSSearchableItemAttributeSet) GpsAreaInformation() string {
	rv := objc.Send[string](c_.ID, objc.Sel("gpsAreaInformation"))
	return rv
}


// Information about the GPS area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsareainformation
func (c_ CSSearchableItemAttributeSet) SetGpsAreaInformation(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsAreaInformation:"), objc.String(value))
}


// The date and time related to the GPS value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdatestamp
func (c_ CSSearchableItemAttributeSet) GpsDateStamp() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("gpsDateStamp"))
	return rv
}


// The date and time related to the GPS value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdatestamp
func (c_ CSSearchableItemAttributeSet) SetGpsDateStamp(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDateStamp:"), value)
}


// The bearing to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestbearing
func (c_ CSSearchableItemAttributeSet) GpsDestBearing() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsDestBearing"))
	return rv
}


// The bearing to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestbearing
func (c_ CSSearchableItemAttributeSet) SetGpsDestBearing(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestBearing:"), value)
}


// The distance to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestdistance
func (c_ CSSearchableItemAttributeSet) GpsDestDistance() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsDestDistance"))
	return rv
}


// The distance to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestdistance
func (c_ CSSearchableItemAttributeSet) SetGpsDestDistance(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestDistance:"), value)
}


// The latitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlatitude
func (c_ CSSearchableItemAttributeSet) GpsDestLatitude() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsDestLatitude"))
	return rv
}


// The latitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlatitude
func (c_ CSSearchableItemAttributeSet) SetGpsDestLatitude(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestLatitude:"), value)
}


// The longitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlongitude
func (c_ CSSearchableItemAttributeSet) GpsDestLongitude() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsDestLongitude"))
	return rv
}


// The longitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlongitude
func (c_ CSSearchableItemAttributeSet) SetGpsDestLongitude(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestLongitude:"), value)
}


// The differential correction applied to the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdifferental
func (c_ CSSearchableItemAttributeSet) GpsDifferental() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsDifferental"))
	return rv
}


// The differential correction applied to the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdifferental
func (c_ CSSearchableItemAttributeSet) SetGpsDifferental(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDifferental:"), value)
}


// The geodetic data that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmapdatum
func (c_ CSSearchableItemAttributeSet) GpsMapDatum() string {
	rv := objc.Send[string](c_.ID, objc.Sel("gpsMapDatum"))
	return rv
}


// The geodetic data that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmapdatum
func (c_ CSSearchableItemAttributeSet) SetGpsMapDatum(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsMapDatum:"), objc.String(value))
}


// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmeasuremode
func (c_ CSSearchableItemAttributeSet) GpsMeasureMode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("gpsMeasureMode"))
	return rv
}


// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmeasuremode
func (c_ CSSearchableItemAttributeSet) SetGpsMeasureMode(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsMeasureMode:"), objc.String(value))
}


// The location finding method that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsprocessingmethod
func (c_ CSSearchableItemAttributeSet) GpsProcessingMethod() string {
	rv := objc.Send[string](c_.ID, objc.Sel("gpsProcessingMethod"))
	return rv
}


// The location finding method that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsprocessingmethod
func (c_ CSSearchableItemAttributeSet) SetGpsProcessingMethod(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsProcessingMethod:"), objc.String(value))
}


// The status of the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsstatus
func (c_ CSSearchableItemAttributeSet) GpsStatus() string {
	rv := objc.Send[string](c_.ID, objc.Sel("gpsStatus"))
	return rv
}


// The status of the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsstatus
func (c_ CSSearchableItemAttributeSet) SetGpsStatus(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsStatus:"), objc.String(value))
}


// The direction of travel of the item in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpstrack
func (c_ CSSearchableItemAttributeSet) GpsTrack() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsTrack"))
	return rv
}


// The direction of travel of the item in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpstrack
func (c_ CSSearchableItemAttributeSet) SetGpsTrack(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsTrack:"), value)
}


// Indicates if the image file has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/hasalphachannel
func (c_ CSSearchableItemAttributeSet) HasAlphaChannel() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("hasAlphaChannel"))
	return rv
}


// Indicates if the image file has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/hasalphachannel
func (c_ CSSearchableItemAttributeSet) SetHasAlphaChannel(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasAlphaChannel:"), value)
}


// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/headline
func (c_ CSSearchableItemAttributeSet) Headline() string {
	rv := objc.Send[string](c_.ID, objc.Sel("headline"))
	return rv
}


// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/headline
func (c_ CSSearchableItemAttributeSet) SetHeadline(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHeadline:"), objc.String(value))
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/hiddenadditionalrecipients
func (c_ CSSearchableItemAttributeSet) HiddenAdditionalRecipients() ICSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("hiddenAdditionalRecipients"))
	return rv
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/hiddenadditionalrecipients
func (c_ CSSearchableItemAttributeSet) SetHiddenAdditionalRecipients(value ICSPerson) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHiddenAdditionalRecipients:"), value)
}


// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/htmlcontentdata
func (c_ CSSearchableItemAttributeSet) HtmlContentData() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("htmlContentData"))
	return rv
}


// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/htmlcontentdata
func (c_ CSSearchableItemAttributeSet) SetHtmlContentData(value foundation.Data) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHtmlContentData:"), value)
}


// A formal identifier that references the document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/identifier
func (c_ CSSearchableItemAttributeSet) Identifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("identifier"))
	return rv
}


// A formal identifier that references the document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/identifier
func (c_ CSSearchableItemAttributeSet) SetIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The direction of the item’s image in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/imagedirection
func (c_ CSSearchableItemAttributeSet) ImageDirection() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("imageDirection"))
	return rv
}


// The direction of the item’s image in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/imagedirection
func (c_ CSSearchableItemAttributeSet) SetImageDirection(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageDirection:"), value)
}


// An array of important dates associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/importantdates
func (c_ CSSearchableItemAttributeSet) ImportantDates() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("importantDates"))
	return rv
}


// An array of important dates associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/importantdates
func (c_ CSSearchableItemAttributeSet) SetImportantDates(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImportantDates:"), value)
}


// Information about the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/information
func (c_ CSSearchableItemAttributeSet) Information() string {
	rv := objc.Send[string](c_.ID, objc.Sel("information"))
	return rv
}


// Information about the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/information
func (c_ CSSearchableItemAttributeSet) SetInformation(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInformation:"), objc.String(value))
}


// An array of instant message addresses for the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instantmessageaddresses
func (c_ CSSearchableItemAttributeSet) InstantMessageAddresses() string {
	rv := objc.Send[string](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}


// An array of instant message addresses for the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instantmessageaddresses
func (c_ CSSearchableItemAttributeSet) SetInstantMessageAddresses(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstantMessageAddresses:"), objc.String(value))
}


// Instructions that concern the use of the item, such as an embargo or warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instructions
func (c_ CSSearchableItemAttributeSet) Instructions() string {
	rv := objc.Send[string](c_.ID, objc.Sel("instructions"))
	return rv
}


// Instructions that concern the use of the item, such as an embargo or warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instructions
func (c_ CSSearchableItemAttributeSet) SetInstructions(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstructions:"), objc.String(value))
}


// A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ispriority
func (c_ CSSearchableItemAttributeSet) IsPriority() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("isPriority"))
	return rv
}


// A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ispriority
func (c_ CSSearchableItemAttributeSet) SetIsPriority(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPriority:"), value)
}


// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/isospeed
func (c_ CSSearchableItemAttributeSet) IsoSpeed() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("isoSpeed"))
	return rv
}


// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/isospeed
func (c_ CSSearchableItemAttributeSet) SetIsoSpeed(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsoSpeed:"), value)
}


// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keysignature
func (c_ CSSearchableItemAttributeSet) KeySignature() string {
	rv := objc.Send[string](c_.ID, objc.Sel("keySignature"))
	return rv
}


// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keysignature
func (c_ CSSearchableItemAttributeSet) SetKeySignature(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeySignature:"), objc.String(value))
}


// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keywords
func (c_ CSSearchableItemAttributeSet) Keywords() string {
	rv := objc.Send[string](c_.ID, objc.Sel("keywords"))
	return rv
}


// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keywords
func (c_ CSSearchableItemAttributeSet) SetKeywords(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeywords:"), objc.String(value))
}


// A description of the kind of document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/kind
func (c_ CSSearchableItemAttributeSet) Kind() string {
	rv := objc.Send[string](c_.ID, objc.Sel("kind"))
	return rv
}


// A description of the kind of document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/kind
func (c_ CSSearchableItemAttributeSet) SetKind(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKind:"), objc.String(value))
}


// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/languages
func (c_ CSSearchableItemAttributeSet) Languages() string {
	rv := objc.Send[string](c_.ID, objc.Sel("languages"))
	return rv
}


// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/languages
func (c_ CSSearchableItemAttributeSet) SetLanguages(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguages:"), objc.String(value))
}


// The date on which the file was last used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lastuseddate
func (c_ CSSearchableItemAttributeSet) LastUsedDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("lastUsedDate"))
	return rv
}


// The date on which the file was last used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lastuseddate
func (c_ CSSearchableItemAttributeSet) SetLastUsedDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLastUsedDate:"), value)
}


// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/latitude
func (c_ CSSearchableItemAttributeSet) Latitude() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("latitude"))
	return rv
}


// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/latitude
func (c_ CSSearchableItemAttributeSet) SetLatitude(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLatitude:"), value)
}


// An array that contains the names of the various layers in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/layernames
func (c_ CSSearchableItemAttributeSet) LayerNames() string {
	rv := objc.Send[string](c_.ID, objc.Sel("layerNames"))
	return rv
}


// An array that contains the names of the various layers in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/layernames
func (c_ CSSearchableItemAttributeSet) SetLayerNames(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLayerNames:"), objc.String(value))
}


// The model of the lens that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lensmodel
func (c_ CSSearchableItemAttributeSet) LensModel() string {
	rv := objc.Send[string](c_.ID, objc.Sel("lensModel"))
	return rv
}


// The model of the lens that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lensmodel
func (c_ CSSearchableItemAttributeSet) SetLensModel(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLensModel:"), objc.String(value))
}


// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/likelyjunk
func (c_ CSSearchableItemAttributeSet) LikelyJunk() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("likelyJunk"))
	return rv
}


// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/likelyjunk
func (c_ CSSearchableItemAttributeSet) SetLikelyJunk(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLikelyJunk:"), value)
}


// A value that indicates if the media is local.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/local
func (c_ CSSearchableItemAttributeSet) Local() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("local"))
	return rv
}


// A value that indicates if the media is local.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/local
func (c_ CSSearchableItemAttributeSet) SetLocal(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocal:"), value)
}


// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/longitude
func (c_ CSSearchableItemAttributeSet) Longitude() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("longitude"))
	return rv
}


// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/longitude
func (c_ CSSearchableItemAttributeSet) SetLongitude(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongitude:"), value)
}


// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lyricist
func (c_ CSSearchableItemAttributeSet) Lyricist() string {
	rv := objc.Send[string](c_.ID, objc.Sel("lyricist"))
	return rv
}


// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lyricist
func (c_ CSSearchableItemAttributeSet) SetLyricist(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLyricist:"), objc.String(value))
}


// An array of mailbox identifiers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mailboxidentifiers
func (c_ CSSearchableItemAttributeSet) MailboxIdentifiers() string {
	rv := objc.Send[string](c_.ID, objc.Sel("mailboxIdentifiers"))
	return rv
}


// An array of mailbox identifiers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mailboxidentifiers
func (c_ CSSearchableItemAttributeSet) SetMailboxIdentifiers(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMailboxIdentifiers:"), objc.String(value))
}


// The smallest F number of the lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/maxaperture
func (c_ CSSearchableItemAttributeSet) MaxAperture() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("maxAperture"))
	return rv
}


// The smallest F number of the lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/maxaperture
func (c_ CSSearchableItemAttributeSet) SetMaxAperture(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxAperture:"), value)
}


// The media types present in the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mediatypes
func (c_ CSSearchableItemAttributeSet) MediaTypes() string {
	rv := objc.Send[string](c_.ID, objc.Sel("mediaTypes"))
	return rv
}


// The media types present in the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mediatypes
func (c_ CSSearchableItemAttributeSet) SetMediaTypes(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMediaTypes:"), objc.String(value))
}


// The date on which the last metadata attribute was changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/metadatamodificationdate
func (c_ CSSearchableItemAttributeSet) MetadataModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("metadataModificationDate"))
	return rv
}


// The date on which the last metadata attribute was changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/metadatamodificationdate
func (c_ CSSearchableItemAttributeSet) SetMetadataModificationDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataModificationDate:"), value)
}


// The metering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/meteringmode
func (c_ CSSearchableItemAttributeSet) MeteringMode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("meteringMode"))
	return rv
}


// The metering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/meteringmode
func (c_ CSSearchableItemAttributeSet) SetMeteringMode(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMeteringMode:"), objc.String(value))
}


// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalgenre
func (c_ CSSearchableItemAttributeSet) MusicalGenre() string {
	rv := objc.Send[string](c_.ID, objc.Sel("musicalGenre"))
	return rv
}


// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalgenre
func (c_ CSSearchableItemAttributeSet) SetMusicalGenre(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalGenre:"), objc.String(value))
}


// The category of the instrument associated with the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentcategory
func (c_ CSSearchableItemAttributeSet) MusicalInstrumentCategory() string {
	rv := objc.Send[string](c_.ID, objc.Sel("musicalInstrumentCategory"))
	return rv
}


// The category of the instrument associated with the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentcategory
func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentCategory(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentCategory:"), objc.String(value))
}


// The name of an instrument within the context of an instrument category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentname
func (c_ CSSearchableItemAttributeSet) MusicalInstrumentName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("musicalInstrumentName"))
	return rv
}


// The name of an instrument within the context of an instrument category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentname
func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentName:"), objc.String(value))
}


// The name of the location or point of interest associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/namedlocation
func (c_ CSSearchableItemAttributeSet) NamedLocation() string {
	rv := objc.Send[string](c_.ID, objc.Sel("namedLocation"))
	return rv
}


// The name of the location or point of interest associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/namedlocation
func (c_ CSSearchableItemAttributeSet) SetNamedLocation(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNamedLocation:"), objc.String(value))
}


// A list of companies or organizations that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/organizations
func (c_ CSSearchableItemAttributeSet) Organizations() string {
	rv := objc.Send[string](c_.ID, objc.Sel("organizations"))
	return rv
}


// A list of companies or organizations that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/organizations
func (c_ CSSearchableItemAttributeSet) SetOrganizations(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrganizations:"), objc.String(value))
}


// The orientation of the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/orientation
func (c_ CSSearchableItemAttributeSet) Orientation() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("orientation"))
	return rv
}


// The orientation of the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/orientation
func (c_ CSSearchableItemAttributeSet) SetOrientation(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrientation:"), value)
}


// The original format of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalformat
func (c_ CSSearchableItemAttributeSet) OriginalFormat() string {
	rv := objc.Send[string](c_.ID, objc.Sel("originalFormat"))
	return rv
}


// The original format of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalformat
func (c_ CSSearchableItemAttributeSet) SetOriginalFormat(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalFormat:"), objc.String(value))
}


// The original source of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalsource
func (c_ CSSearchableItemAttributeSet) OriginalSource() string {
	rv := objc.Send[string](c_.ID, objc.Sel("originalSource"))
	return rv
}


// The original source of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalsource
func (c_ CSSearchableItemAttributeSet) SetOriginalSource(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalSource:"), objc.String(value))
}


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagecount
func (c_ CSSearchableItemAttributeSet) PageCount() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pageCount"))
	return rv
}


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagecount
func (c_ CSSearchableItemAttributeSet) SetPageCount(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageCount:"), value)
}


// The height of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pageheight
func (c_ CSSearchableItemAttributeSet) PageHeight() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pageHeight"))
	return rv
}


// The height of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pageheight
func (c_ CSSearchableItemAttributeSet) SetPageHeight(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageHeight:"), value)
}


// The width of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagewidth
func (c_ CSSearchableItemAttributeSet) PageWidth() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pageWidth"))
	return rv
}


// The width of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagewidth
func (c_ CSSearchableItemAttributeSet) SetPageWidth(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageWidth:"), value)
}


// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/participants
func (c_ CSSearchableItemAttributeSet) Participants() string {
	rv := objc.Send[string](c_.ID, objc.Sel("participants"))
	return rv
}


// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/participants
func (c_ CSSearchableItemAttributeSet) SetParticipants(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParticipants:"), objc.String(value))
}


// The complete path to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/path
func (c_ CSSearchableItemAttributeSet) Path() string {
	rv := objc.Send[string](c_.ID, objc.Sel("path"))
	return rv
}


// The complete path to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/path
func (c_ CSSearchableItemAttributeSet) SetPath(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPath:"), objc.String(value))
}


// A list of performers in the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/performers
func (c_ CSSearchableItemAttributeSet) Performers() string {
	rv := objc.Send[string](c_.ID, objc.Sel("performers"))
	return rv
}


// A list of performers in the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/performers
func (c_ CSSearchableItemAttributeSet) SetPerformers(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerformers:"), objc.String(value))
}


// An array of phone numbers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/phonenumbers
func (c_ CSSearchableItemAttributeSet) PhoneNumbers() string {
	rv := objc.Send[string](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}


// An array of phone numbers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/phonenumbers
func (c_ CSSearchableItemAttributeSet) SetPhoneNumbers(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneNumbers:"), objc.String(value))
}


// The total number of pixels in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelcount
func (c_ CSSearchableItemAttributeSet) PixelCount() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pixelCount"))
	return rv
}


// The total number of pixels in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelcount
func (c_ CSSearchableItemAttributeSet) SetPixelCount(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelCount:"), value)
}


// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelheight
func (c_ CSSearchableItemAttributeSet) PixelHeight() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pixelHeight"))
	return rv
}


// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelheight
func (c_ CSSearchableItemAttributeSet) SetPixelHeight(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelHeight:"), value)
}


// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelwidth
func (c_ CSSearchableItemAttributeSet) PixelWidth() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pixelWidth"))
	return rv
}


// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelwidth
func (c_ CSSearchableItemAttributeSet) SetPixelWidth(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelWidth:"), value)
}


// A user-supplied play count for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/playcount
func (c_ CSSearchableItemAttributeSet) PlayCount() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("playCount"))
	return rv
}


// A user-supplied play count for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/playcount
func (c_ CSSearchableItemAttributeSet) SetPlayCount(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayCount:"), value)
}


// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/postalcode
func (c_ CSSearchableItemAttributeSet) PostalCode() string {
	rv := objc.Send[string](c_.ID, objc.Sel("postalCode"))
	return rv
}


// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/postalcode
func (c_ CSSearchableItemAttributeSet) SetPostalCode(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalCode:"), objc.String(value))
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/primaryrecipients
func (c_ CSSearchableItemAttributeSet) PrimaryRecipients() ICSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("primaryRecipients"))
	return rv
}


// An array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/primaryrecipients
func (c_ CSSearchableItemAttributeSet) SetPrimaryRecipients(value ICSPerson) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryRecipients:"), value)
}


// The producer of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/producer
func (c_ CSSearchableItemAttributeSet) Producer() string {
	rv := objc.Send[string](c_.ID, objc.Sel("producer"))
	return rv
}


// The producer of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/producer
func (c_ CSSearchableItemAttributeSet) SetProducer(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProducer:"), objc.String(value))
}


// The name of the color profile the camera used for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/profilename
func (c_ CSSearchableItemAttributeSet) ProfileName() string {
	rv := objc.Send[string](c_.ID, objc.Sel("profileName"))
	return rv
}


// The name of the color profile the camera used for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/profilename
func (c_ CSSearchableItemAttributeSet) SetProfileName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProfileName:"), objc.String(value))
}


// A list of projects of which this file is a part.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/projects
func (c_ CSSearchableItemAttributeSet) Projects() string {
	rv := objc.Send[string](c_.ID, objc.Sel("projects"))
	return rv
}


// A list of projects of which this file is a part.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/projects
func (c_ CSSearchableItemAttributeSet) SetProjects(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProjects:"), objc.String(value))
}


// An array of identifiers that corresponds to data representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerdatatypeidentifiers
func (c_ CSSearchableItemAttributeSet) ProviderDataTypeIdentifiers() string {
	rv := objc.Send[string](c_.ID, objc.Sel("providerDataTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to data representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerdatatypeidentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderDataTypeIdentifiers(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderDataTypeIdentifiers:"), objc.String(value))
}


// An array of identifiers that corresponds to file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerfiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) ProviderFileTypeIdentifiers() string {
	rv := objc.Send[string](c_.ID, objc.Sel("providerFileTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerfiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderFileTypeIdentifiers(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderFileTypeIdentifiers:"), objc.String(value))
}


// An array of identifiers that corresponds to in-place file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerinplacefiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) ProviderInPlaceFileTypeIdentifiers() string {
	rv := objc.Send[string](c_.ID, objc.Sel("providerInPlaceFileTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to in-place file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerinplacefiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderInPlaceFileTypeIdentifiers(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderInPlaceFileTypeIdentifiers:"), objc.String(value))
}


// A list of people, organizations, services, or other entities responsible for making the media available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/publishers
func (c_ CSSearchableItemAttributeSet) Publishers() string {
	rv := objc.Send[string](c_.ID, objc.Sel("publishers"))
	return rv
}


// A list of people, organizations, services, or other entities responsible for making the media available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/publishers
func (c_ CSSearchableItemAttributeSet) SetPublishers(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPublishers:"), objc.String(value))
}


// A number that indicates the relative importance of the item among other items from the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rankinghint
func (c_ CSSearchableItemAttributeSet) RankingHint() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("rankingHint"))
	return rv
}


// A number that indicates the relative importance of the item among other items from the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rankinghint
func (c_ CSSearchableItemAttributeSet) SetRankingHint(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRankingHint:"), value)
}


// The user-supplied rating of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rating
func (c_ CSSearchableItemAttributeSet) Rating() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("rating"))
	return rv
}


// The user-supplied rating of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rating
func (c_ CSSearchableItemAttributeSet) SetRating(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRating:"), value)
}


// A description of the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ratingdescription
func (c_ CSSearchableItemAttributeSet) RatingDescription() string {
	rv := objc.Send[string](c_.ID, objc.Sel("ratingDescription"))
	return rv
}


// A description of the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ratingdescription
func (c_ CSSearchableItemAttributeSet) SetRatingDescription(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRatingDescription:"), objc.String(value))
}


// An array of addresses associated with the recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientaddresses
func (c_ CSSearchableItemAttributeSet) RecipientAddresses() string {
	rv := objc.Send[string](c_.ID, objc.Sel("recipientAddresses"))
	return rv
}


// An array of addresses associated with the recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientaddresses
func (c_ CSSearchableItemAttributeSet) SetRecipientAddresses(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientAddresses:"), objc.String(value))
}


// An array of email addresses associated with the recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientemailaddresses
func (c_ CSSearchableItemAttributeSet) RecipientEmailAddresses() string {
	rv := objc.Send[string](c_.ID, objc.Sel("recipientEmailAddresses"))
	return rv
}


// An array of email addresses associated with the recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientemailaddresses
func (c_ CSSearchableItemAttributeSet) SetRecipientEmailAddresses(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientEmailAddresses:"), objc.String(value))
}


// An array of names representing the recipients of this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientnames
func (c_ CSSearchableItemAttributeSet) RecipientNames() string {
	rv := objc.Send[string](c_.ID, objc.Sel("recipientNames"))
	return rv
}


// An array of names representing the recipients of this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientnames
func (c_ CSSearchableItemAttributeSet) SetRecipientNames(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientNames:"), objc.String(value))
}


// The recording date of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recordingdate
func (c_ CSSearchableItemAttributeSet) RecordingDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("recordingDate"))
	return rv
}


// The recording date of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recordingdate
func (c_ CSSearchableItemAttributeSet) SetRecordingDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordingDate:"), value)
}


// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/redeyeon
func (c_ CSSearchableItemAttributeSet) RedEyeOn() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("redEyeOn"))
	return rv
}


// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/redeyeon
func (c_ CSSearchableItemAttributeSet) SetRedEyeOn(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRedEyeOn:"), value)
}


// The unique identifier for the item to which the activity is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/relateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) RelatedUniqueIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("relatedUniqueIdentifier"))
	return rv
}


// The unique identifier for the item to which the activity is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/relateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) SetRelatedUniqueIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRelatedUniqueIdentifier:"), objc.String(value))
}


// The resolution height of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionheightdpi
func (c_ CSSearchableItemAttributeSet) ResolutionHeightDPI() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("resolutionHeightDPI"))
	return rv
}


// The resolution height of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionheightdpi
func (c_ CSSearchableItemAttributeSet) SetResolutionHeightDPI(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionHeightDPI:"), value)
}


// The resolution width of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionwidthdpi
func (c_ CSSearchableItemAttributeSet) ResolutionWidthDPI() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("resolutionWidthDPI"))
	return rv
}


// The resolution width of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionwidthdpi
func (c_ CSSearchableItemAttributeSet) SetResolutionWidthDPI(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionWidthDPI:"), value)
}


// A link to information about the rights held in and over the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rights
func (c_ CSSearchableItemAttributeSet) Rights() string {
	rv := objc.Send[string](c_.ID, objc.Sel("rights"))
	return rv
}


// A link to information about the rights held in and over the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rights
func (c_ CSSearchableItemAttributeSet) SetRights(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRights:"), objc.String(value))
}


// Indicates the role of the content creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/role
func (c_ CSSearchableItemAttributeSet) Role() string {
	rv := objc.Send[string](c_.ID, objc.Sel("role"))
	return rv
}


// Indicates the role of the content creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/role
func (c_ CSSearchableItemAttributeSet) SetRole(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRole:"), objc.String(value))
}


// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/securitymethod
func (c_ CSSearchableItemAttributeSet) SecurityMethod() string {
	rv := objc.Send[string](c_.ID, objc.Sel("securityMethod"))
	return rv
}


// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/securitymethod
func (c_ CSSearchableItemAttributeSet) SetSecurityMethod(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecurityMethod:"), objc.String(value))
}


// The file type of the item to enable the user to share items from Spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/shareditemcontenttype
func (c_ CSSearchableItemAttributeSet) SharedItemContentType() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("sharedItemContentType"))
	return rv
}


// The file type of the item to enable the user to share items from Spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/shareditemcontenttype
func (c_ CSSearchableItemAttributeSet) SetSharedItemContentType(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSharedItemContentType:"), value)
}


// The speed of the item, in kilometers per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/speed
func (c_ CSSearchableItemAttributeSet) Speed() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("speed"))
	return rv
}


// The speed of the item, in kilometers per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/speed
func (c_ CSSearchableItemAttributeSet) SetSpeed(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpeed:"), value)
}


// The start date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/startdate
func (c_ CSSearchableItemAttributeSet) StartDate() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("startDate"))
	return rv
}


// The start date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/startdate
func (c_ CSSearchableItemAttributeSet) SetStartDate(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStartDate:"), value)
}


// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/stateorprovince
func (c_ CSSearchableItemAttributeSet) StateOrProvince() string {
	rv := objc.Send[string](c_.ID, objc.Sel("stateOrProvince"))
	return rv
}


// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/stateorprovince
func (c_ CSSearchableItemAttributeSet) SetStateOrProvince(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStateOrProvince:"), objc.String(value))
}


// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/streamable
func (c_ CSSearchableItemAttributeSet) Streamable() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("streamable"))
	return rv
}


// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/streamable
func (c_ CSSearchableItemAttributeSet) SetStreamable(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreamable:"), value)
}


// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subthoroughfare
func (c_ CSSearchableItemAttributeSet) SubThoroughfare() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subThoroughfare"))
	return rv
}


// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subthoroughfare
func (c_ CSSearchableItemAttributeSet) SetSubThoroughfare(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubThoroughfare:"), objc.String(value))
}


// The subject of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subject
func (c_ CSSearchableItemAttributeSet) Subject() string {
	rv := objc.Send[string](c_.ID, objc.Sel("subject"))
	return rv
}


// The subject of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subject
func (c_ CSSearchableItemAttributeSet) SetSubject(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubject:"), objc.String(value))
}


// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/supportsphonecall
func (c_ CSSearchableItemAttributeSet) SupportsPhoneCall() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("supportsPhoneCall"))
	return rv
}


// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/supportsphonecall
func (c_ CSSearchableItemAttributeSet) SetSupportsPhoneCall(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsPhoneCall:"), value)
}


// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/tempo
func (c_ CSSearchableItemAttributeSet) Tempo() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("tempo"))
	return rv
}


// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/tempo
func (c_ CSSearchableItemAttributeSet) SetTempo(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTempo:"), value)
}


// The textual content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontent
func (c_ CSSearchableItemAttributeSet) TextContent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("textContent"))
	return rv
}


// The textual content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontent
func (c_ CSSearchableItemAttributeSet) SetTextContent(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextContent:"), objc.String(value))
}


// A string that presents the Apple Intelligence summarization of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontentsummary
func (c_ CSSearchableItemAttributeSet) TextContentSummary() string {
	rv := objc.Send[string](c_.ID, objc.Sel("textContentSummary"))
	return rv
}


// A string that presents the Apple Intelligence summarization of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontentsummary
func (c_ CSSearchableItemAttributeSet) SetTextContentSummary(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextContentSummary:"), objc.String(value))
}


// The theme of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/theme
func (c_ CSSearchableItemAttributeSet) Theme() string {
	rv := objc.Send[string](c_.ID, objc.Sel("theme"))
	return rv
}


// The theme of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/theme
func (c_ CSSearchableItemAttributeSet) SetTheme(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTheme:"), objc.String(value))
}


// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thoroughfare
func (c_ CSSearchableItemAttributeSet) Thoroughfare() string {
	rv := objc.Send[string](c_.ID, objc.Sel("thoroughfare"))
	return rv
}


// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thoroughfare
func (c_ CSSearchableItemAttributeSet) SetThoroughfare(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThoroughfare:"), objc.String(value))
}


// Image data that represents the thumbnail of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnaildata
func (c_ CSSearchableItemAttributeSet) ThumbnailData() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("thumbnailData"))
	return rv
}


// Image data that represents the thumbnail of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnaildata
func (c_ CSSearchableItemAttributeSet) SetThumbnailData(value foundation.Data) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailData:"), value)
}


// The local file URL of the thumbnail image for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnailurl
func (c_ CSSearchableItemAttributeSet) ThumbnailURL() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("thumbnailURL"))
	return rv
}


// The local file URL of the thumbnail image for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnailurl
func (c_ CSSearchableItemAttributeSet) SetThumbnailURL(value foundation.URL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailURL:"), value)
}


// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timesignature
func (c_ CSSearchableItemAttributeSet) TimeSignature() string {
	rv := objc.Send[string](c_.ID, objc.Sel("timeSignature"))
	return rv
}


// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timesignature
func (c_ CSSearchableItemAttributeSet) SetTimeSignature(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeSignature:"), objc.String(value))
}


// The timestamp on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timestamp
func (c_ CSSearchableItemAttributeSet) Timestamp() foundation.Date {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("timestamp"))
	return rv
}


// The timestamp on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timestamp
func (c_ CSSearchableItemAttributeSet) SetTimestamp(value foundation.Date) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimestamp:"), value)
}


// The total bit rate of the media, combining audio and video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/totalbitrate
func (c_ CSSearchableItemAttributeSet) TotalBitRate() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("totalBitRate"))
	return rv
}


// The total bit rate of the media, combining audio and video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/totalbitrate
func (c_ CSSearchableItemAttributeSet) SetTotalBitRate(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalBitRate:"), value)
}


// A string that represents the text the system transcribed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/transcribedtextcontent
func (c_ CSSearchableItemAttributeSet) TranscribedTextContent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("transcribedTextContent"))
	return rv
}


// A string that represents the text the system transcribed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/transcribedtextcontent
func (c_ CSSearchableItemAttributeSet) SetTranscribedTextContent(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTranscribedTextContent:"), objc.String(value))
}


// The URL associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/url
func (c_ CSSearchableItemAttributeSet) Url() foundation.URL {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("url"))
	return rv
}


// The URL associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/url
func (c_ CSSearchableItemAttributeSet) SetUrl(value foundation.URL) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrl:"), value)
}


// A value that indicates the user created the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercreated
func (c_ CSSearchableItemAttributeSet) UserCreated() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("userCreated"))
	return rv
}


// A value that indicates the user created the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercreated
func (c_ CSSearchableItemAttributeSet) SetUserCreated(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCreated:"), value)
}


// A value that indicates the user selected the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercurated
func (c_ CSSearchableItemAttributeSet) UserCurated() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("userCurated"))
	return rv
}


// A value that indicates the user selected the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercurated
func (c_ CSSearchableItemAttributeSet) SetUserCurated(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCurated:"), value)
}


// A value that indicates the user purchased or owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/userowned
func (c_ CSSearchableItemAttributeSet) UserOwned() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("userOwned"))
	return rv
}


// A value that indicates the user purchased or owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/userowned
func (c_ CSSearchableItemAttributeSet) SetUserOwned(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserOwned:"), value)
}


// A version string associated with the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/version
func (c_ CSSearchableItemAttributeSet) Version() string {
	rv := objc.Send[string](c_.ID, objc.Sel("version"))
	return rv
}


// A version string associated with the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/version
func (c_ CSSearchableItemAttributeSet) SetVersion(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVersion:"), objc.String(value))
}


// The video bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/videobitrate
func (c_ CSSearchableItemAttributeSet) VideoBitRate() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("videoBitRate"))
	return rv
}


// The video bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/videobitrate
func (c_ CSSearchableItemAttributeSet) SetVideoBitRate(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoBitRate:"), value)
}


// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/weakrelateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) WeakRelatedUniqueIdentifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("weakRelatedUniqueIdentifier"))
	return rv
}


// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/weakrelateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) SetWeakRelatedUniqueIdentifier(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeakRelatedUniqueIdentifier:"), objc.String(value))
}


// The white balance setting when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/whitebalance
func (c_ CSSearchableItemAttributeSet) WhiteBalance() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("whiteBalance"))
	return rv
}


// The white balance setting when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/whitebalance
func (c_ CSSearchableItemAttributeSet) SetWhiteBalance(value foundation.Number) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteBalance:"), value)
}



