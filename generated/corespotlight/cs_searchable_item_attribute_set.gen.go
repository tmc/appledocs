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
	// properties:
	ContentModificationDate() foundation.objc.IObject /* cross-framework: NSDate */
	SetContentModificationDate(value foundation.objc.IObject /* cross-framework: NSDate */)
	DisplayName() string /* primitive/slice/pointer. */
	SetDisplayName(value string /* primitive/slice/pointer. */)
	EncodingApplications() []string /* primitive/slice/pointer. */
	SetEncodingApplications(value []string /* primitive/slice/pointer. */)
	EndDate() foundation.objc.IObject /* cross-framework: NSDate */
	SetEndDate(value foundation.objc.IObject /* cross-framework: NSDate */)
	GPSDOP() foundation.objc.IObject /* cross-framework: Number */
	SetGPSDOP(value foundation.objc.IObject /* cross-framework: Number */)
	SupportsNavigation() foundation.objc.IObject /* cross-framework: Number */
	SetSupportsNavigation(value foundation.objc.IObject /* cross-framework: Number */)
	Title() string /* primitive/slice/pointer. */
	SetTitle(value string /* primitive/slice/pointer. */)
	CSActionIdentifier() string /* primitive/slice/pointer. */
	AccountHandles() string /* primitive/slice/pointer. */
	SetAccountHandles(value string /* primitive/slice/pointer. */)
	AccountIdentifier() string /* primitive/slice/pointer. */
	SetAccountIdentifier(value string /* primitive/slice/pointer. */)
	AcquisitionMake() string /* primitive/slice/pointer. */
	SetAcquisitionMake(value string /* primitive/slice/pointer. */)
	AcquisitionModel() string /* primitive/slice/pointer. */
	SetAcquisitionModel(value string /* primitive/slice/pointer. */)
	ActionIdentifiers() string /* primitive/slice/pointer. */
	SetActionIdentifiers(value string /* primitive/slice/pointer. */)
	AddedDate() foundation.objc.IObject /* cross-framework: Date */
	SetAddedDate(value foundation.objc.IObject /* cross-framework: Date */)
	AdditionalRecipients() ICSPerson
	SetAdditionalRecipients(value ICSPerson)
	Album() string /* primitive/slice/pointer. */
	SetAlbum(value string /* primitive/slice/pointer. */)
	AllDay() foundation.objc.IObject /* cross-framework: Number */
	SetAllDay(value foundation.objc.IObject /* cross-framework: Number */)
	AlternateNames() string /* primitive/slice/pointer. */
	SetAlternateNames(value string /* primitive/slice/pointer. */)
	Altitude() foundation.objc.IObject /* cross-framework: Number */
	SetAltitude(value foundation.objc.IObject /* cross-framework: Number */)
	Aperture() foundation.objc.IObject /* cross-framework: Number */
	SetAperture(value foundation.objc.IObject /* cross-framework: Number */)
	Artist() string /* primitive/slice/pointer. */
	SetArtist(value string /* primitive/slice/pointer. */)
	Audiences() string /* primitive/slice/pointer. */
	SetAudiences(value string /* primitive/slice/pointer. */)
	AudioBitRate() foundation.objc.IObject /* cross-framework: Number */
	SetAudioBitRate(value foundation.objc.IObject /* cross-framework: Number */)
	AudioChannelCount() foundation.objc.IObject /* cross-framework: Number */
	SetAudioChannelCount(value foundation.objc.IObject /* cross-framework: Number */)
	AudioEncodingApplication() string /* primitive/slice/pointer. */
	SetAudioEncodingApplication(value string /* primitive/slice/pointer. */)
	AudioSampleRate() foundation.objc.IObject /* cross-framework: Number */
	SetAudioSampleRate(value foundation.objc.IObject /* cross-framework: Number */)
	AudioTrackNumber() foundation.objc.IObject /* cross-framework: Number */
	SetAudioTrackNumber(value foundation.objc.IObject /* cross-framework: Number */)
	AuthorAddresses() string /* primitive/slice/pointer. */
	SetAuthorAddresses(value string /* primitive/slice/pointer. */)
	AuthorEmailAddresses() string /* primitive/slice/pointer. */
	SetAuthorEmailAddresses(value string /* primitive/slice/pointer. */)
	AuthorNames() string /* primitive/slice/pointer. */
	SetAuthorNames(value string /* primitive/slice/pointer. */)
	Authors() ICSPerson
	SetAuthors(value ICSPerson)
	BitsPerSample() foundation.objc.IObject /* cross-framework: Number */
	SetBitsPerSample(value foundation.objc.IObject /* cross-framework: Number */)
	CameraOwner() string /* primitive/slice/pointer. */
	SetCameraOwner(value string /* primitive/slice/pointer. */)
	City() string /* primitive/slice/pointer. */
	SetCity(value string /* primitive/slice/pointer. */)
	Codecs() string /* primitive/slice/pointer. */
	SetCodecs(value string /* primitive/slice/pointer. */)
	ColorSpace() string /* primitive/slice/pointer. */
	SetColorSpace(value string /* primitive/slice/pointer. */)
	Comment() string /* primitive/slice/pointer. */
	SetComment(value string /* primitive/slice/pointer. */)
	CompletionDate() foundation.objc.IObject /* cross-framework: Date */
	SetCompletionDate(value foundation.objc.IObject /* cross-framework: Date */)
	Composer() string /* primitive/slice/pointer. */
	SetComposer(value string /* primitive/slice/pointer. */)
	ContactKeywords() string /* primitive/slice/pointer. */
	SetContactKeywords(value string /* primitive/slice/pointer. */)
	ContainerDisplayName() string /* primitive/slice/pointer. */
	SetContainerDisplayName(value string /* primitive/slice/pointer. */)
	ContainerIdentifier() string /* primitive/slice/pointer. */
	SetContainerIdentifier(value string /* primitive/slice/pointer. */)
	ContainerOrder() foundation.objc.IObject /* cross-framework: Number */
	SetContainerOrder(value foundation.objc.IObject /* cross-framework: Number */)
	ContainerTitle() string /* primitive/slice/pointer. */
	SetContainerTitle(value string /* primitive/slice/pointer. */)
	ContentCreationDate() foundation.objc.IObject /* cross-framework: Date */
	SetContentCreationDate(value foundation.objc.IObject /* cross-framework: Date */)
	ContentDescription() string /* primitive/slice/pointer. */
	SetContentDescription(value string /* primitive/slice/pointer. */)
	ContentRating() foundation.objc.IObject /* cross-framework: Number */
	SetContentRating(value foundation.objc.IObject /* cross-framework: Number */)
	ContentSources() string /* primitive/slice/pointer. */
	SetContentSources(value string /* primitive/slice/pointer. */)
	ContentType() string /* primitive/slice/pointer. */
	SetContentType(value string /* primitive/slice/pointer. */)
	ContentTypeTree() string /* primitive/slice/pointer. */
	SetContentTypeTree(value string /* primitive/slice/pointer. */)
	ContentURL() foundation.objc.IObject /* cross-framework: URL */
	SetContentURL(value foundation.objc.IObject /* cross-framework: URL */)
	Contributors() string /* primitive/slice/pointer. */
	SetContributors(value string /* primitive/slice/pointer. */)
	Copyright() string /* primitive/slice/pointer. */
	SetCopyright(value string /* primitive/slice/pointer. */)
	Country() string /* primitive/slice/pointer. */
	SetCountry(value string /* primitive/slice/pointer. */)
	Coverage() string /* primitive/slice/pointer. */
	SetCoverage(value string /* primitive/slice/pointer. */)
	Creator() string /* primitive/slice/pointer. */
	SetCreator(value string /* primitive/slice/pointer. */)
	DarkThumbnailURL() foundation.objc.IObject /* cross-framework: URL */
	SetDarkThumbnailURL(value foundation.objc.IObject /* cross-framework: URL */)
	DeliveryType() foundation.objc.IObject /* cross-framework: Number */
	SetDeliveryType(value foundation.objc.IObject /* cross-framework: Number */)
	Director() string /* primitive/slice/pointer. */
	SetDirector(value string /* primitive/slice/pointer. */)
	DomainIdentifier() string /* primitive/slice/pointer. */
	SetDomainIdentifier(value string /* primitive/slice/pointer. */)
	DownloadedDate() foundation.objc.IObject /* cross-framework: Date */
	SetDownloadedDate(value foundation.objc.IObject /* cross-framework: Date */)
	DueDate() foundation.objc.IObject /* cross-framework: Date */
	SetDueDate(value foundation.objc.IObject /* cross-framework: Date */)
	Duration() foundation.objc.IObject /* cross-framework: Number */
	SetDuration(value foundation.objc.IObject /* cross-framework: Number */)
	Editors() string /* primitive/slice/pointer. */
	SetEditors(value string /* primitive/slice/pointer. */)
	EmailAddresses() string /* primitive/slice/pointer. */
	SetEmailAddresses(value string /* primitive/slice/pointer. */)
	EmailHeaders() string /* primitive/slice/pointer. */
	SetEmailHeaders(value string /* primitive/slice/pointer. */)
	ExifVersion() string /* primitive/slice/pointer. */
	SetExifVersion(value string /* primitive/slice/pointer. */)
	ExifgpsVersion() string /* primitive/slice/pointer. */
	SetExifgpsVersion(value string /* primitive/slice/pointer. */)
	ExposureMode() foundation.objc.IObject /* cross-framework: Number */
	SetExposureMode(value foundation.objc.IObject /* cross-framework: Number */)
	ExposureProgram() string /* primitive/slice/pointer. */
	SetExposureProgram(value string /* primitive/slice/pointer. */)
	ExposureTime() foundation.objc.IObject /* cross-framework: Number */
	SetExposureTime(value foundation.objc.IObject /* cross-framework: Number */)
	ExposureTimeString() string /* primitive/slice/pointer. */
	SetExposureTimeString(value string /* primitive/slice/pointer. */)
	FNumber() foundation.objc.IObject /* cross-framework: Number */
	SetFNumber(value foundation.objc.IObject /* cross-framework: Number */)
	FileSize() foundation.objc.IObject /* cross-framework: Number */
	SetFileSize(value foundation.objc.IObject /* cross-framework: Number */)
	FlashOn() foundation.objc.IObject /* cross-framework: Number */
	SetFlashOn(value foundation.objc.IObject /* cross-framework: Number */)
	FocalLength() foundation.objc.IObject /* cross-framework: Number */
	SetFocalLength(value foundation.objc.IObject /* cross-framework: Number */)
	FocalLength35mm() foundation.objc.IObject /* cross-framework: Number */
	SetFocalLength35mm(value foundation.objc.IObject /* cross-framework: Number */)
	FontNames() string /* primitive/slice/pointer. */
	SetFontNames(value string /* primitive/slice/pointer. */)
	FullyFormattedAddress() string /* primitive/slice/pointer. */
	SetFullyFormattedAddress(value string /* primitive/slice/pointer. */)
	GeneralMIDISequence() foundation.objc.IObject /* cross-framework: Number */
	SetGeneralMIDISequence(value foundation.objc.IObject /* cross-framework: Number */)
	Genre() string /* primitive/slice/pointer. */
	SetGenre(value string /* primitive/slice/pointer. */)
	GpsAreaInformation() string /* primitive/slice/pointer. */
	SetGpsAreaInformation(value string /* primitive/slice/pointer. */)
	GpsDateStamp() foundation.objc.IObject /* cross-framework: Date */
	SetGpsDateStamp(value foundation.objc.IObject /* cross-framework: Date */)
	GpsDestBearing() foundation.objc.IObject /* cross-framework: Number */
	SetGpsDestBearing(value foundation.objc.IObject /* cross-framework: Number */)
	GpsDestDistance() foundation.objc.IObject /* cross-framework: Number */
	SetGpsDestDistance(value foundation.objc.IObject /* cross-framework: Number */)
	GpsDestLatitude() foundation.objc.IObject /* cross-framework: Number */
	SetGpsDestLatitude(value foundation.objc.IObject /* cross-framework: Number */)
	GpsDestLongitude() foundation.objc.IObject /* cross-framework: Number */
	SetGpsDestLongitude(value foundation.objc.IObject /* cross-framework: Number */)
	GpsDifferental() foundation.objc.IObject /* cross-framework: Number */
	SetGpsDifferental(value foundation.objc.IObject /* cross-framework: Number */)
	GpsMapDatum() string /* primitive/slice/pointer. */
	SetGpsMapDatum(value string /* primitive/slice/pointer. */)
	GpsMeasureMode() string /* primitive/slice/pointer. */
	SetGpsMeasureMode(value string /* primitive/slice/pointer. */)
	GpsProcessingMethod() string /* primitive/slice/pointer. */
	SetGpsProcessingMethod(value string /* primitive/slice/pointer. */)
	GpsStatus() string /* primitive/slice/pointer. */
	SetGpsStatus(value string /* primitive/slice/pointer. */)
	GpsTrack() foundation.objc.IObject /* cross-framework: Number */
	SetGpsTrack(value foundation.objc.IObject /* cross-framework: Number */)
	HasAlphaChannel() foundation.objc.IObject /* cross-framework: Number */
	SetHasAlphaChannel(value foundation.objc.IObject /* cross-framework: Number */)
	Headline() string /* primitive/slice/pointer. */
	SetHeadline(value string /* primitive/slice/pointer. */)
	HiddenAdditionalRecipients() ICSPerson
	SetHiddenAdditionalRecipients(value ICSPerson)
	HtmlContentData() foundation.objc.IObject /* cross-framework: Data */
	SetHtmlContentData(value foundation.objc.IObject /* cross-framework: Data */)
	Identifier() string /* primitive/slice/pointer. */
	SetIdentifier(value string /* primitive/slice/pointer. */)
	ImageDirection() foundation.objc.IObject /* cross-framework: Number */
	SetImageDirection(value foundation.objc.IObject /* cross-framework: Number */)
	ImportantDates() foundation.objc.IObject /* cross-framework: Date */
	SetImportantDates(value foundation.objc.IObject /* cross-framework: Date */)
	Information() string /* primitive/slice/pointer. */
	SetInformation(value string /* primitive/slice/pointer. */)
	InstantMessageAddresses() string /* primitive/slice/pointer. */
	SetInstantMessageAddresses(value string /* primitive/slice/pointer. */)
	Instructions() string /* primitive/slice/pointer. */
	SetInstructions(value string /* primitive/slice/pointer. */)
	IsPriority() foundation.objc.IObject /* cross-framework: Number */
	SetIsPriority(value foundation.objc.IObject /* cross-framework: Number */)
	IsoSpeed() foundation.objc.IObject /* cross-framework: Number */
	SetIsoSpeed(value foundation.objc.IObject /* cross-framework: Number */)
	KeySignature() string /* primitive/slice/pointer. */
	SetKeySignature(value string /* primitive/slice/pointer. */)
	Keywords() string /* primitive/slice/pointer. */
	SetKeywords(value string /* primitive/slice/pointer. */)
	Kind() string /* primitive/slice/pointer. */
	SetKind(value string /* primitive/slice/pointer. */)
	Languages() string /* primitive/slice/pointer. */
	SetLanguages(value string /* primitive/slice/pointer. */)
	LastUsedDate() foundation.objc.IObject /* cross-framework: Date */
	SetLastUsedDate(value foundation.objc.IObject /* cross-framework: Date */)
	Latitude() foundation.objc.IObject /* cross-framework: Number */
	SetLatitude(value foundation.objc.IObject /* cross-framework: Number */)
	LayerNames() string /* primitive/slice/pointer. */
	SetLayerNames(value string /* primitive/slice/pointer. */)
	LensModel() string /* primitive/slice/pointer. */
	SetLensModel(value string /* primitive/slice/pointer. */)
	LikelyJunk() foundation.objc.IObject /* cross-framework: Number */
	SetLikelyJunk(value foundation.objc.IObject /* cross-framework: Number */)
	Local() foundation.objc.IObject /* cross-framework: Number */
	SetLocal(value foundation.objc.IObject /* cross-framework: Number */)
	Longitude() foundation.objc.IObject /* cross-framework: Number */
	SetLongitude(value foundation.objc.IObject /* cross-framework: Number */)
	Lyricist() string /* primitive/slice/pointer. */
	SetLyricist(value string /* primitive/slice/pointer. */)
	MailboxIdentifiers() string /* primitive/slice/pointer. */
	SetMailboxIdentifiers(value string /* primitive/slice/pointer. */)
	MaxAperture() foundation.objc.IObject /* cross-framework: Number */
	SetMaxAperture(value foundation.objc.IObject /* cross-framework: Number */)
	MediaTypes() string /* primitive/slice/pointer. */
	SetMediaTypes(value string /* primitive/slice/pointer. */)
	MetadataModificationDate() foundation.objc.IObject /* cross-framework: Date */
	SetMetadataModificationDate(value foundation.objc.IObject /* cross-framework: Date */)
	MeteringMode() string /* primitive/slice/pointer. */
	SetMeteringMode(value string /* primitive/slice/pointer. */)
	MusicalGenre() string /* primitive/slice/pointer. */
	SetMusicalGenre(value string /* primitive/slice/pointer. */)
	MusicalInstrumentCategory() string /* primitive/slice/pointer. */
	SetMusicalInstrumentCategory(value string /* primitive/slice/pointer. */)
	MusicalInstrumentName() string /* primitive/slice/pointer. */
	SetMusicalInstrumentName(value string /* primitive/slice/pointer. */)
	NamedLocation() string /* primitive/slice/pointer. */
	SetNamedLocation(value string /* primitive/slice/pointer. */)
	Organizations() string /* primitive/slice/pointer. */
	SetOrganizations(value string /* primitive/slice/pointer. */)
	Orientation() foundation.objc.IObject /* cross-framework: Number */
	SetOrientation(value foundation.objc.IObject /* cross-framework: Number */)
	OriginalFormat() string /* primitive/slice/pointer. */
	SetOriginalFormat(value string /* primitive/slice/pointer. */)
	OriginalSource() string /* primitive/slice/pointer. */
	SetOriginalSource(value string /* primitive/slice/pointer. */)
	PageCount() foundation.objc.IObject /* cross-framework: Number */
	SetPageCount(value foundation.objc.IObject /* cross-framework: Number */)
	PageHeight() foundation.objc.IObject /* cross-framework: Number */
	SetPageHeight(value foundation.objc.IObject /* cross-framework: Number */)
	PageWidth() foundation.objc.IObject /* cross-framework: Number */
	SetPageWidth(value foundation.objc.IObject /* cross-framework: Number */)
	Participants() string /* primitive/slice/pointer. */
	SetParticipants(value string /* primitive/slice/pointer. */)
	Path() string /* primitive/slice/pointer. */
	SetPath(value string /* primitive/slice/pointer. */)
	Performers() string /* primitive/slice/pointer. */
	SetPerformers(value string /* primitive/slice/pointer. */)
	PhoneNumbers() string /* primitive/slice/pointer. */
	SetPhoneNumbers(value string /* primitive/slice/pointer. */)
	PixelCount() foundation.objc.IObject /* cross-framework: Number */
	SetPixelCount(value foundation.objc.IObject /* cross-framework: Number */)
	PixelHeight() foundation.objc.IObject /* cross-framework: Number */
	SetPixelHeight(value foundation.objc.IObject /* cross-framework: Number */)
	PixelWidth() foundation.objc.IObject /* cross-framework: Number */
	SetPixelWidth(value foundation.objc.IObject /* cross-framework: Number */)
	PlayCount() foundation.objc.IObject /* cross-framework: Number */
	SetPlayCount(value foundation.objc.IObject /* cross-framework: Number */)
	PostalCode() string /* primitive/slice/pointer. */
	SetPostalCode(value string /* primitive/slice/pointer. */)
	PrimaryRecipients() ICSPerson
	SetPrimaryRecipients(value ICSPerson)
	Producer() string /* primitive/slice/pointer. */
	SetProducer(value string /* primitive/slice/pointer. */)
	ProfileName() string /* primitive/slice/pointer. */
	SetProfileName(value string /* primitive/slice/pointer. */)
	Projects() string /* primitive/slice/pointer. */
	SetProjects(value string /* primitive/slice/pointer. */)
	ProviderDataTypeIdentifiers() string /* primitive/slice/pointer. */
	SetProviderDataTypeIdentifiers(value string /* primitive/slice/pointer. */)
	ProviderFileTypeIdentifiers() string /* primitive/slice/pointer. */
	SetProviderFileTypeIdentifiers(value string /* primitive/slice/pointer. */)
	ProviderInPlaceFileTypeIdentifiers() string /* primitive/slice/pointer. */
	SetProviderInPlaceFileTypeIdentifiers(value string /* primitive/slice/pointer. */)
	Publishers() string /* primitive/slice/pointer. */
	SetPublishers(value string /* primitive/slice/pointer. */)
	RankingHint() foundation.objc.IObject /* cross-framework: Number */
	SetRankingHint(value foundation.objc.IObject /* cross-framework: Number */)
	Rating() foundation.objc.IObject /* cross-framework: Number */
	SetRating(value foundation.objc.IObject /* cross-framework: Number */)
	RatingDescription() string /* primitive/slice/pointer. */
	SetRatingDescription(value string /* primitive/slice/pointer. */)
	RecipientAddresses() string /* primitive/slice/pointer. */
	SetRecipientAddresses(value string /* primitive/slice/pointer. */)
	RecipientEmailAddresses() string /* primitive/slice/pointer. */
	SetRecipientEmailAddresses(value string /* primitive/slice/pointer. */)
	RecipientNames() string /* primitive/slice/pointer. */
	SetRecipientNames(value string /* primitive/slice/pointer. */)
	RecordingDate() foundation.objc.IObject /* cross-framework: Date */
	SetRecordingDate(value foundation.objc.IObject /* cross-framework: Date */)
	RedEyeOn() foundation.objc.IObject /* cross-framework: Number */
	SetRedEyeOn(value foundation.objc.IObject /* cross-framework: Number */)
	RelatedUniqueIdentifier() string /* primitive/slice/pointer. */
	SetRelatedUniqueIdentifier(value string /* primitive/slice/pointer. */)
	ResolutionHeightDPI() foundation.objc.IObject /* cross-framework: Number */
	SetResolutionHeightDPI(value foundation.objc.IObject /* cross-framework: Number */)
	ResolutionWidthDPI() foundation.objc.IObject /* cross-framework: Number */
	SetResolutionWidthDPI(value foundation.objc.IObject /* cross-framework: Number */)
	Rights() string /* primitive/slice/pointer. */
	SetRights(value string /* primitive/slice/pointer. */)
	Role() string /* primitive/slice/pointer. */
	SetRole(value string /* primitive/slice/pointer. */)
	SecurityMethod() string /* primitive/slice/pointer. */
	SetSecurityMethod(value string /* primitive/slice/pointer. */)
	SharedItemContentType() objectivec.IObject
	SetSharedItemContentType(value objectivec.IObject)
	Speed() foundation.objc.IObject /* cross-framework: Number */
	SetSpeed(value foundation.objc.IObject /* cross-framework: Number */)
	StartDate() foundation.objc.IObject /* cross-framework: Date */
	SetStartDate(value foundation.objc.IObject /* cross-framework: Date */)
	StateOrProvince() string /* primitive/slice/pointer. */
	SetStateOrProvince(value string /* primitive/slice/pointer. */)
	Streamable() foundation.objc.IObject /* cross-framework: Number */
	SetStreamable(value foundation.objc.IObject /* cross-framework: Number */)
	SubThoroughfare() string /* primitive/slice/pointer. */
	SetSubThoroughfare(value string /* primitive/slice/pointer. */)
	Subject() string /* primitive/slice/pointer. */
	SetSubject(value string /* primitive/slice/pointer. */)
	SupportsPhoneCall() foundation.objc.IObject /* cross-framework: Number */
	SetSupportsPhoneCall(value foundation.objc.IObject /* cross-framework: Number */)
	Tempo() foundation.objc.IObject /* cross-framework: Number */
	SetTempo(value foundation.objc.IObject /* cross-framework: Number */)
	TextContent() string /* primitive/slice/pointer. */
	SetTextContent(value string /* primitive/slice/pointer. */)
	TextContentSummary() string /* primitive/slice/pointer. */
	SetTextContentSummary(value string /* primitive/slice/pointer. */)
	Theme() string /* primitive/slice/pointer. */
	SetTheme(value string /* primitive/slice/pointer. */)
	Thoroughfare() string /* primitive/slice/pointer. */
	SetThoroughfare(value string /* primitive/slice/pointer. */)
	ThumbnailData() foundation.objc.IObject /* cross-framework: Data */
	SetThumbnailData(value foundation.objc.IObject /* cross-framework: Data */)
	ThumbnailURL() foundation.objc.IObject /* cross-framework: URL */
	SetThumbnailURL(value foundation.objc.IObject /* cross-framework: URL */)
	TimeSignature() string /* primitive/slice/pointer. */
	SetTimeSignature(value string /* primitive/slice/pointer. */)
	Timestamp() foundation.objc.IObject /* cross-framework: Date */
	SetTimestamp(value foundation.objc.IObject /* cross-framework: Date */)
	TotalBitRate() foundation.objc.IObject /* cross-framework: Number */
	SetTotalBitRate(value foundation.objc.IObject /* cross-framework: Number */)
	TranscribedTextContent() string /* primitive/slice/pointer. */
	SetTranscribedTextContent(value string /* primitive/slice/pointer. */)
	Url() foundation.objc.IObject /* cross-framework: URL */
	SetUrl(value foundation.objc.IObject /* cross-framework: URL */)
	UserCreated() foundation.objc.IObject /* cross-framework: Number */
	SetUserCreated(value foundation.objc.IObject /* cross-framework: Number */)
	UserCurated() foundation.objc.IObject /* cross-framework: Number */
	SetUserCurated(value foundation.objc.IObject /* cross-framework: Number */)
	UserOwned() foundation.objc.IObject /* cross-framework: Number */
	SetUserOwned(value foundation.objc.IObject /* cross-framework: Number */)
	Version() string /* primitive/slice/pointer. */
	SetVersion(value string /* primitive/slice/pointer. */)
	VideoBitRate() foundation.objc.IObject /* cross-framework: Number */
	SetVideoBitRate(value foundation.objc.IObject /* cross-framework: Number */)
	WeakRelatedUniqueIdentifier() string /* primitive/slice/pointer. */
	SetWeakRelatedUniqueIdentifier(value string /* primitive/slice/pointer. */)
	WhiteBalance() foundation.objc.IObject /* cross-framework: Number */
	SetWhiteBalance(value foundation.objc.IObject /* cross-framework: Number */)
	// methods:
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
func (c_ CSSearchableItemAttributeSet) ContentModificationDate() foundation.objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("contentModificationDate"))
	return rv
}


// The date on which the contents of the file was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentModificationDate
func (c_ CSSearchableItemAttributeSet) SetContentModificationDate(value foundation.objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentModificationDate:"), value)
}


// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName
func (c_ CSSearchableItemAttributeSet) DisplayName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("displayName"))
	return rv
}


// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName
func (c_ CSSearchableItemAttributeSet) SetDisplayName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}


// The name of the apps that converted the original content into a PDF stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/encodingApplications
func (c_ CSSearchableItemAttributeSet) EncodingApplications() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](c_.ID, objc.Sel("encodingApplications"))
	return rv
}


// The name of the apps that converted the original content into a PDF stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/encodingApplications
func (c_ CSSearchableItemAttributeSet) SetEncodingApplications(value []string /* primitive/slice/pointer. */) {
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
func (c_ CSSearchableItemAttributeSet) EndDate() foundation.objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("endDate"))
	return rv
}


// The end date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/endDate
func (c_ CSSearchableItemAttributeSet) SetEndDate(value foundation.objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEndDate:"), value)
}


// The GPS dilution of precision value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop
func (c_ CSSearchableItemAttributeSet) GPSDOP() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("GPSDOP"))
	return rv
}


// The GPS dilution of precision value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop
func (c_ CSSearchableItemAttributeSet) SetGPSDOP(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDOP:"), value)
}


// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation
func (c_ CSSearchableItemAttributeSet) SupportsNavigation() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("supportsNavigation"))
	return rv
}


// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation
func (c_ CSSearchableItemAttributeSet) SetSupportsNavigation(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsNavigation:"), value)
}


// The title of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title
func (c_ CSSearchableItemAttributeSet) Title() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("title"))
	return rv
}


// The title of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title
func (c_ CSSearchableItemAttributeSet) SetTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// A key that specifies the action’s identifier in a user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csactionidentifier
func (c_ CSSearchableItemAttributeSet) CSActionIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CSActionIdentifier"))
	return rv
}


// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accounthandles
func (c_ CSSearchableItemAttributeSet) AccountHandles() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("accountHandles"))
	return rv
}


// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accounthandles
func (c_ CSSearchableItemAttributeSet) SetAccountHandles(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccountHandles:"), objc.String(value))
}


// The unique identifier for the account with which the message is associated, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accountidentifier
func (c_ CSSearchableItemAttributeSet) AccountIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("accountIdentifier"))
	return rv
}


// The unique identifier for the account with which the message is associated, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accountidentifier
func (c_ CSSearchableItemAttributeSet) SetAccountIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccountIdentifier:"), objc.String(value))
}


// The manufacturer of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmake
func (c_ CSSearchableItemAttributeSet) AcquisitionMake() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("acquisitionMake"))
	return rv
}


// The manufacturer of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmake
func (c_ CSSearchableItemAttributeSet) SetAcquisitionMake(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionMake:"), objc.String(value))
}


// The model of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmodel
func (c_ CSSearchableItemAttributeSet) AcquisitionModel() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("acquisitionModel"))
	return rv
}


// The model of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmodel
func (c_ CSSearchableItemAttributeSet) SetAcquisitionModel(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionModel:"), objc.String(value))
}


// The identifiers that specify custom actions the app supports for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/actionidentifiers
func (c_ CSSearchableItemAttributeSet) ActionIdentifiers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("actionIdentifiers"))
	return rv
}


// The identifiers that specify custom actions the app supports for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/actionidentifiers
func (c_ CSSearchableItemAttributeSet) SetActionIdentifiers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActionIdentifiers:"), objc.String(value))
}


// The date on which the item was moved into its current location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/addeddate
func (c_ CSSearchableItemAttributeSet) AddedDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("addedDate"))
	return rv
}


// The date on which the item was moved into its current location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/addeddate
func (c_ CSSearchableItemAttributeSet) SetAddedDate(value foundation.objc.IObject /* cross-framework: Date */) {
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
func (c_ CSSearchableItemAttributeSet) Album() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("album"))
	return rv
}


// The title for a collection of audio media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/album
func (c_ CSSearchableItemAttributeSet) SetAlbum(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlbum:"), objc.String(value))
}


// A value that indicates if the event covers an entire day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/allday
func (c_ CSSearchableItemAttributeSet) AllDay() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("allDay"))
	return rv
}


// A value that indicates if the event covers an entire day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/allday
func (c_ CSSearchableItemAttributeSet) SetAllDay(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllDay:"), value)
}


// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/alternatenames
func (c_ CSSearchableItemAttributeSet) AlternateNames() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("alternateNames"))
	return rv
}


// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/alternatenames
func (c_ CSSearchableItemAttributeSet) SetAlternateNames(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlternateNames:"), objc.String(value))
}


// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/altitude
func (c_ CSSearchableItemAttributeSet) Altitude() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("altitude"))
	return rv
}


// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/altitude
func (c_ CSSearchableItemAttributeSet) SetAltitude(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAltitude:"), value)
}


// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/aperture
func (c_ CSSearchableItemAttributeSet) Aperture() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("aperture"))
	return rv
}


// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/aperture
func (c_ CSSearchableItemAttributeSet) SetAperture(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAperture:"), value)
}


// The artist associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/artist
func (c_ CSSearchableItemAttributeSet) Artist() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("artist"))
	return rv
}


// The artist associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/artist
func (c_ CSSearchableItemAttributeSet) SetArtist(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setArtist:"), objc.String(value))
}


// A class of entity for which the item is intended or useful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiences
func (c_ CSSearchableItemAttributeSet) Audiences() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("audiences"))
	return rv
}


// A class of entity for which the item is intended or useful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiences
func (c_ CSSearchableItemAttributeSet) SetAudiences(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudiences:"), objc.String(value))
}


// The audio bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiobitrate
func (c_ CSSearchableItemAttributeSet) AudioBitRate() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioBitRate"))
	return rv
}


// The audio bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiobitrate
func (c_ CSSearchableItemAttributeSet) SetAudioBitRate(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioBitRate:"), value)
}


// The number of channels in the audio data that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiochannelcount
func (c_ CSSearchableItemAttributeSet) AudioChannelCount() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioChannelCount"))
	return rv
}


// The number of channels in the audio data that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiochannelcount
func (c_ CSSearchableItemAttributeSet) SetAudioChannelCount(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioChannelCount:"), value)
}


// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audioencodingapplication
func (c_ CSSearchableItemAttributeSet) AudioEncodingApplication() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("audioEncodingApplication"))
	return rv
}


// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audioencodingapplication
func (c_ CSSearchableItemAttributeSet) SetAudioEncodingApplication(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioEncodingApplication:"), objc.String(value))
}


// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiosamplerate
func (c_ CSSearchableItemAttributeSet) AudioSampleRate() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioSampleRate"))
	return rv
}


// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiosamplerate
func (c_ CSSearchableItemAttributeSet) SetAudioSampleRate(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSampleRate:"), value)
}


// The track number of a song or audio composition when part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiotracknumber
func (c_ CSSearchableItemAttributeSet) AudioTrackNumber() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("audioTrackNumber"))
	return rv
}


// The track number of a song or audio composition when part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiotracknumber
func (c_ CSSearchableItemAttributeSet) SetAudioTrackNumber(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioTrackNumber:"), value)
}


// An array of addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoraddresses
func (c_ CSSearchableItemAttributeSet) AuthorAddresses() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("authorAddresses"))
	return rv
}


// An array of addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoraddresses
func (c_ CSSearchableItemAttributeSet) SetAuthorAddresses(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorAddresses:"), objc.String(value))
}


// An array of email addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoremailaddresses
func (c_ CSSearchableItemAttributeSet) AuthorEmailAddresses() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("authorEmailAddresses"))
	return rv
}


// An array of email addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoremailaddresses
func (c_ CSSearchableItemAttributeSet) SetAuthorEmailAddresses(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorEmailAddresses:"), objc.String(value))
}


// An array of names representing the authors who have worked on the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authornames
func (c_ CSSearchableItemAttributeSet) AuthorNames() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("authorNames"))
	return rv
}


// An array of names representing the authors who have worked on the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authornames
func (c_ CSSearchableItemAttributeSet) SetAuthorNames(value string /* primitive/slice/pointer. */) {
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
func (c_ CSSearchableItemAttributeSet) BitsPerSample() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("bitsPerSample"))
	return rv
}


// The number of bits per sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/bitspersample
func (c_ CSSearchableItemAttributeSet) SetBitsPerSample(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBitsPerSample:"), value)
}


// The owner of the camera that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/cameraowner
func (c_ CSSearchableItemAttributeSet) CameraOwner() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("cameraOwner"))
	return rv
}


// The owner of the camera that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/cameraowner
func (c_ CSSearchableItemAttributeSet) SetCameraOwner(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraOwner:"), objc.String(value))
}


// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/city
func (c_ CSSearchableItemAttributeSet) City() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("city"))
	return rv
}


// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/city
func (c_ CSSearchableItemAttributeSet) SetCity(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), objc.String(value))
}


// The codecs used to encode/decode the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/codecs
func (c_ CSSearchableItemAttributeSet) Codecs() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("codecs"))
	return rv
}


// The codecs used to encode/decode the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/codecs
func (c_ CSSearchableItemAttributeSet) SetCodecs(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCodecs:"), objc.String(value))
}


// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/colorspace
func (c_ CSSearchableItemAttributeSet) ColorSpace() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("colorSpace"))
	return rv
}


// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/colorspace
func (c_ CSSearchableItemAttributeSet) SetColorSpace(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSpace:"), objc.String(value))
}


// A comment related to the media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/comment
func (c_ CSSearchableItemAttributeSet) Comment() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("comment"))
	return rv
}


// A comment related to the media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/comment
func (c_ CSSearchableItemAttributeSet) SetComment(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComment:"), objc.String(value))
}


// The date on which the item was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/completiondate
func (c_ CSSearchableItemAttributeSet) CompletionDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("completionDate"))
	return rv
}


// The date on which the item was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/completiondate
func (c_ CSSearchableItemAttributeSet) SetCompletionDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionDate:"), value)
}


// The composer of the song or audio composition that the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/composer
func (c_ CSSearchableItemAttributeSet) Composer() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("composer"))
	return rv
}


// The composer of the song or audio composition that the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/composer
func (c_ CSSearchableItemAttributeSet) SetComposer(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComposer:"), objc.String(value))
}


// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contactkeywords
func (c_ CSSearchableItemAttributeSet) ContactKeywords() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("contactKeywords"))
	return rv
}


// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contactkeywords
func (c_ CSSearchableItemAttributeSet) SetContactKeywords(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactKeywords:"), objc.String(value))
}


// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerdisplayname
func (c_ CSSearchableItemAttributeSet) ContainerDisplayName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("containerDisplayName"))
	return rv
}


// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerdisplayname
func (c_ CSSearchableItemAttributeSet) SetContainerDisplayName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerDisplayName:"), objc.String(value))
}


// The identifier of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containeridentifier
func (c_ CSSearchableItemAttributeSet) ContainerIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}


// The identifier of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containeridentifier
func (c_ CSSearchableItemAttributeSet) SetContainerIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), objc.String(value))
}


// The order of the item within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerorder
func (c_ CSSearchableItemAttributeSet) ContainerOrder() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("containerOrder"))
	return rv
}


// The order of the item within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerorder
func (c_ CSSearchableItemAttributeSet) SetContainerOrder(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerOrder:"), value)
}


// The title of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containertitle
func (c_ CSSearchableItemAttributeSet) ContainerTitle() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("containerTitle"))
	return rv
}


// The title of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containertitle
func (c_ CSSearchableItemAttributeSet) SetContainerTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerTitle:"), objc.String(value))
}


// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentcreationdate
func (c_ CSSearchableItemAttributeSet) ContentCreationDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("contentCreationDate"))
	return rv
}


// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentcreationdate
func (c_ CSSearchableItemAttributeSet) SetContentCreationDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentCreationDate:"), value)
}


// A description of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentdescription
func (c_ CSSearchableItemAttributeSet) ContentDescription() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("contentDescription"))
	return rv
}


// A description of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentdescription
func (c_ CSSearchableItemAttributeSet) SetContentDescription(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentDescription:"), objc.String(value))
}


// A value that indicates if the media contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentrating
func (c_ CSSearchableItemAttributeSet) ContentRating() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("contentRating"))
	return rv
}


// A value that indicates if the media contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentrating
func (c_ CSSearchableItemAttributeSet) SetContentRating(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentRating:"), value)
}


// An array of sources from which the media was obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentsources
func (c_ CSSearchableItemAttributeSet) ContentSources() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("contentSources"))
	return rv
}


// An array of sources from which the media was obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentsources
func (c_ CSSearchableItemAttributeSet) SetContentSources(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentSources:"), objc.String(value))
}


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttype
func (c_ CSSearchableItemAttributeSet) ContentType() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("contentType"))
	return rv
}


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttype
func (c_ CSSearchableItemAttributeSet) SetContentType(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentType:"), objc.String(value))
}


// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttypetree
func (c_ CSSearchableItemAttributeSet) ContentTypeTree() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("contentTypeTree"))
	return rv
}


// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttypetree
func (c_ CSSearchableItemAttributeSet) SetContentTypeTree(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentTypeTree:"), objc.String(value))
}


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenturl
func (c_ CSSearchableItemAttributeSet) ContentURL() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("contentURL"))
	return rv
}


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenturl
func (c_ CSSearchableItemAttributeSet) SetContentURL(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentURL:"), value)
}


// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contributors
func (c_ CSSearchableItemAttributeSet) Contributors() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("contributors"))
	return rv
}


// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contributors
func (c_ CSSearchableItemAttributeSet) SetContributors(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContributors:"), objc.String(value))
}


// The copyright date of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/copyright
func (c_ CSSearchableItemAttributeSet) Copyright() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("copyright"))
	return rv
}


// The copyright date of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/copyright
func (c_ CSSearchableItemAttributeSet) SetCopyright(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCopyright:"), objc.String(value))
}


// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/country
func (c_ CSSearchableItemAttributeSet) Country() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("country"))
	return rv
}


// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/country
func (c_ CSSearchableItemAttributeSet) SetCountry(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), objc.String(value))
}


// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/coverage
func (c_ CSSearchableItemAttributeSet) Coverage() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("coverage"))
	return rv
}


// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/coverage
func (c_ CSSearchableItemAttributeSet) SetCoverage(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCoverage:"), objc.String(value))
}


// The name of the app that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/creator
func (c_ CSSearchableItemAttributeSet) Creator() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("creator"))
	return rv
}


// The name of the app that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/creator
func (c_ CSSearchableItemAttributeSet) SetCreator(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreator:"), objc.String(value))
}


// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/darkthumbnailurl
func (c_ CSSearchableItemAttributeSet) DarkThumbnailURL() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("darkThumbnailURL"))
	return rv
}


// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/darkthumbnailurl
func (c_ CSSearchableItemAttributeSet) SetDarkThumbnailURL(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDarkThumbnailURL:"), value)
}


// The delivery type of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/deliverytype
func (c_ CSSearchableItemAttributeSet) DeliveryType() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("deliveryType"))
	return rv
}


// The delivery type of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/deliverytype
func (c_ CSSearchableItemAttributeSet) SetDeliveryType(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeliveryType:"), value)
}


// The name of the director of the media (for example, a movie director).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/director
func (c_ CSSearchableItemAttributeSet) Director() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("director"))
	return rv
}


// The name of the director of the media (for example, a movie director).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/director
func (c_ CSSearchableItemAttributeSet) SetDirector(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDirector:"), objc.String(value))
}


// An identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/domainidentifier
func (c_ CSSearchableItemAttributeSet) DomainIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("domainIdentifier"))
	return rv
}


// An identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/domainidentifier
func (c_ CSSearchableItemAttributeSet) SetDomainIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDomainIdentifier:"), objc.String(value))
}


// The most recent date on which the file was downloaded or received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/downloadeddate
func (c_ CSSearchableItemAttributeSet) DownloadedDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("downloadedDate"))
	return rv
}


// The most recent date on which the file was downloaded or received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/downloadeddate
func (c_ CSSearchableItemAttributeSet) SetDownloadedDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDownloadedDate:"), value)
}


// The date on which the item is due.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duedate
func (c_ CSSearchableItemAttributeSet) DueDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("dueDate"))
	return rv
}


// The date on which the item is due.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duedate
func (c_ CSSearchableItemAttributeSet) SetDueDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDueDate:"), value)
}


// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duration
func (c_ CSSearchableItemAttributeSet) Duration() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("duration"))
	return rv
}


// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duration
func (c_ CSSearchableItemAttributeSet) SetDuration(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDuration:"), value)
}


// A list of editors who have worked on the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/editors
func (c_ CSSearchableItemAttributeSet) Editors() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("editors"))
	return rv
}


// A list of editors who have worked on the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/editors
func (c_ CSSearchableItemAttributeSet) SetEditors(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEditors:"), objc.String(value))
}


// An array of email addresses associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailaddresses
func (c_ CSSearchableItemAttributeSet) EmailAddresses() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("emailAddresses"))
	return rv
}


// An array of email addresses associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailaddresses
func (c_ CSSearchableItemAttributeSet) SetEmailAddresses(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailAddresses:"), objc.String(value))
}


// A dictionary that contains all the headers of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailheaders
func (c_ CSSearchableItemAttributeSet) EmailHeaders() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("emailHeaders"))
	return rv
}


// A dictionary that contains all the headers of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailheaders
func (c_ CSSearchableItemAttributeSet) SetEmailHeaders(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailHeaders:"), objc.String(value))
}


// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifversion
func (c_ CSSearchableItemAttributeSet) ExifVersion() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("exifVersion"))
	return rv
}


// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifversion
func (c_ CSSearchableItemAttributeSet) SetExifVersion(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExifVersion:"), objc.String(value))
}


// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifgpsversion
func (c_ CSSearchableItemAttributeSet) ExifgpsVersion() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("exifgpsVersion"))
	return rv
}


// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifgpsversion
func (c_ CSSearchableItemAttributeSet) SetExifgpsVersion(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExifgpsVersion:"), objc.String(value))
}


// The mode the camera used for the exposure of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuremode
func (c_ CSSearchableItemAttributeSet) ExposureMode() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("exposureMode"))
	return rv
}


// The mode the camera used for the exposure of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuremode
func (c_ CSSearchableItemAttributeSet) SetExposureMode(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureMode:"), value)
}


// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposureprogram
func (c_ CSSearchableItemAttributeSet) ExposureProgram() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("exposureProgram"))
	return rv
}


// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposureprogram
func (c_ CSSearchableItemAttributeSet) SetExposureProgram(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureProgram:"), objc.String(value))
}


// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretime
func (c_ CSSearchableItemAttributeSet) ExposureTime() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("exposureTime"))
	return rv
}


// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretime
func (c_ CSSearchableItemAttributeSet) SetExposureTime(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTime:"), value)
}


// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretimestring
func (c_ CSSearchableItemAttributeSet) ExposureTimeString() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("exposureTimeString"))
	return rv
}


// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretimestring
func (c_ CSSearchableItemAttributeSet) SetExposureTimeString(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTimeString:"), objc.String(value))
}


// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fnumber
func (c_ CSSearchableItemAttributeSet) FNumber() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("fNumber"))
	return rv
}


// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fnumber
func (c_ CSSearchableItemAttributeSet) SetFNumber(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFNumber:"), value)
}


// The size of the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/filesize
func (c_ CSSearchableItemAttributeSet) FileSize() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("fileSize"))
	return rv
}


// The size of the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/filesize
func (c_ CSSearchableItemAttributeSet) SetFileSize(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFileSize:"), value)
}


// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/flashon
func (c_ CSSearchableItemAttributeSet) FlashOn() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("flashOn"))
	return rv
}


// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/flashon
func (c_ CSSearchableItemAttributeSet) SetFlashOn(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashOn:"), value)
}


// The actual focal length of the lens, in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength
func (c_ CSSearchableItemAttributeSet) FocalLength() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("focalLength"))
	return rv
}


// The actual focal length of the lens, in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength
func (c_ CSSearchableItemAttributeSet) SetFocalLength(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength:"), value)
}


// A value that indicates if the focal length is 35mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength35mm
func (c_ CSSearchableItemAttributeSet) FocalLength35mm() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("focalLength35mm"))
	return rv
}


// A value that indicates if the focal length is 35mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength35mm
func (c_ CSSearchableItemAttributeSet) SetFocalLength35mm(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength35mm:"), value)
}


// An array of font names the document uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fontnames
func (c_ CSSearchableItemAttributeSet) FontNames() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("fontNames"))
	return rv
}


// An array of font names the document uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fontnames
func (c_ CSSearchableItemAttributeSet) SetFontNames(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFontNames:"), objc.String(value))
}


// The fully formatted address of the item, received from MapKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fullyformattedaddress
func (c_ CSSearchableItemAttributeSet) FullyFormattedAddress() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("fullyFormattedAddress"))
	return rv
}


// The fully formatted address of the item, received from MapKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fullyformattedaddress
func (c_ CSSearchableItemAttributeSet) SetFullyFormattedAddress(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFullyFormattedAddress:"), objc.String(value))
}


// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/generalmidisequence
func (c_ CSSearchableItemAttributeSet) GeneralMIDISequence() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("generalMIDISequence"))
	return rv
}


// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/generalmidisequence
func (c_ CSSearchableItemAttributeSet) SetGeneralMIDISequence(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGeneralMIDISequence:"), value)
}


// The genre of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/genre
func (c_ CSSearchableItemAttributeSet) Genre() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("genre"))
	return rv
}


// The genre of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/genre
func (c_ CSSearchableItemAttributeSet) SetGenre(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGenre:"), objc.String(value))
}


// Information about the GPS area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsareainformation
func (c_ CSSearchableItemAttributeSet) GpsAreaInformation() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("gpsAreaInformation"))
	return rv
}


// Information about the GPS area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsareainformation
func (c_ CSSearchableItemAttributeSet) SetGpsAreaInformation(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsAreaInformation:"), objc.String(value))
}


// The date and time related to the GPS value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdatestamp
func (c_ CSSearchableItemAttributeSet) GpsDateStamp() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("gpsDateStamp"))
	return rv
}


// The date and time related to the GPS value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdatestamp
func (c_ CSSearchableItemAttributeSet) SetGpsDateStamp(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDateStamp:"), value)
}


// The bearing to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestbearing
func (c_ CSSearchableItemAttributeSet) GpsDestBearing() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsDestBearing"))
	return rv
}


// The bearing to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestbearing
func (c_ CSSearchableItemAttributeSet) SetGpsDestBearing(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestBearing:"), value)
}


// The distance to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestdistance
func (c_ CSSearchableItemAttributeSet) GpsDestDistance() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsDestDistance"))
	return rv
}


// The distance to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestdistance
func (c_ CSSearchableItemAttributeSet) SetGpsDestDistance(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestDistance:"), value)
}


// The latitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlatitude
func (c_ CSSearchableItemAttributeSet) GpsDestLatitude() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsDestLatitude"))
	return rv
}


// The latitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlatitude
func (c_ CSSearchableItemAttributeSet) SetGpsDestLatitude(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestLatitude:"), value)
}


// The longitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlongitude
func (c_ CSSearchableItemAttributeSet) GpsDestLongitude() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsDestLongitude"))
	return rv
}


// The longitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlongitude
func (c_ CSSearchableItemAttributeSet) SetGpsDestLongitude(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestLongitude:"), value)
}


// The differential correction applied to the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdifferental
func (c_ CSSearchableItemAttributeSet) GpsDifferental() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsDifferental"))
	return rv
}


// The differential correction applied to the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdifferental
func (c_ CSSearchableItemAttributeSet) SetGpsDifferental(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDifferental:"), value)
}


// The geodetic data that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmapdatum
func (c_ CSSearchableItemAttributeSet) GpsMapDatum() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("gpsMapDatum"))
	return rv
}


// The geodetic data that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmapdatum
func (c_ CSSearchableItemAttributeSet) SetGpsMapDatum(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsMapDatum:"), objc.String(value))
}


// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmeasuremode
func (c_ CSSearchableItemAttributeSet) GpsMeasureMode() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("gpsMeasureMode"))
	return rv
}


// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmeasuremode
func (c_ CSSearchableItemAttributeSet) SetGpsMeasureMode(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsMeasureMode:"), objc.String(value))
}


// The location finding method that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsprocessingmethod
func (c_ CSSearchableItemAttributeSet) GpsProcessingMethod() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("gpsProcessingMethod"))
	return rv
}


// The location finding method that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsprocessingmethod
func (c_ CSSearchableItemAttributeSet) SetGpsProcessingMethod(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsProcessingMethod:"), objc.String(value))
}


// The status of the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsstatus
func (c_ CSSearchableItemAttributeSet) GpsStatus() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("gpsStatus"))
	return rv
}


// The status of the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsstatus
func (c_ CSSearchableItemAttributeSet) SetGpsStatus(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsStatus:"), objc.String(value))
}


// The direction of travel of the item in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpstrack
func (c_ CSSearchableItemAttributeSet) GpsTrack() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("gpsTrack"))
	return rv
}


// The direction of travel of the item in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpstrack
func (c_ CSSearchableItemAttributeSet) SetGpsTrack(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsTrack:"), value)
}


// Indicates if the image file has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/hasalphachannel
func (c_ CSSearchableItemAttributeSet) HasAlphaChannel() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("hasAlphaChannel"))
	return rv
}


// Indicates if the image file has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/hasalphachannel
func (c_ CSSearchableItemAttributeSet) SetHasAlphaChannel(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasAlphaChannel:"), value)
}


// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/headline
func (c_ CSSearchableItemAttributeSet) Headline() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("headline"))
	return rv
}


// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/headline
func (c_ CSSearchableItemAttributeSet) SetHeadline(value string /* primitive/slice/pointer. */) {
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
func (c_ CSSearchableItemAttributeSet) HtmlContentData() foundation.objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("htmlContentData"))
	return rv
}


// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/htmlcontentdata
func (c_ CSSearchableItemAttributeSet) SetHtmlContentData(value foundation.objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHtmlContentData:"), value)
}


// A formal identifier that references the document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/identifier
func (c_ CSSearchableItemAttributeSet) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("identifier"))
	return rv
}


// A formal identifier that references the document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/identifier
func (c_ CSSearchableItemAttributeSet) SetIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The direction of the item’s image in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/imagedirection
func (c_ CSSearchableItemAttributeSet) ImageDirection() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("imageDirection"))
	return rv
}


// The direction of the item’s image in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/imagedirection
func (c_ CSSearchableItemAttributeSet) SetImageDirection(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageDirection:"), value)
}


// An array of important dates associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/importantdates
func (c_ CSSearchableItemAttributeSet) ImportantDates() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("importantDates"))
	return rv
}


// An array of important dates associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/importantdates
func (c_ CSSearchableItemAttributeSet) SetImportantDates(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImportantDates:"), value)
}


// Information about the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/information
func (c_ CSSearchableItemAttributeSet) Information() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("information"))
	return rv
}


// Information about the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/information
func (c_ CSSearchableItemAttributeSet) SetInformation(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInformation:"), objc.String(value))
}


// An array of instant message addresses for the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instantmessageaddresses
func (c_ CSSearchableItemAttributeSet) InstantMessageAddresses() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}


// An array of instant message addresses for the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instantmessageaddresses
func (c_ CSSearchableItemAttributeSet) SetInstantMessageAddresses(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstantMessageAddresses:"), objc.String(value))
}


// Instructions that concern the use of the item, such as an embargo or warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instructions
func (c_ CSSearchableItemAttributeSet) Instructions() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("instructions"))
	return rv
}


// Instructions that concern the use of the item, such as an embargo or warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instructions
func (c_ CSSearchableItemAttributeSet) SetInstructions(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstructions:"), objc.String(value))
}


// A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ispriority
func (c_ CSSearchableItemAttributeSet) IsPriority() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("isPriority"))
	return rv
}


// A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ispriority
func (c_ CSSearchableItemAttributeSet) SetIsPriority(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPriority:"), value)
}


// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/isospeed
func (c_ CSSearchableItemAttributeSet) IsoSpeed() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("isoSpeed"))
	return rv
}


// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/isospeed
func (c_ CSSearchableItemAttributeSet) SetIsoSpeed(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsoSpeed:"), value)
}


// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keysignature
func (c_ CSSearchableItemAttributeSet) KeySignature() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("keySignature"))
	return rv
}


// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keysignature
func (c_ CSSearchableItemAttributeSet) SetKeySignature(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeySignature:"), objc.String(value))
}


// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keywords
func (c_ CSSearchableItemAttributeSet) Keywords() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("keywords"))
	return rv
}


// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keywords
func (c_ CSSearchableItemAttributeSet) SetKeywords(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeywords:"), objc.String(value))
}


// A description of the kind of document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/kind
func (c_ CSSearchableItemAttributeSet) Kind() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("kind"))
	return rv
}


// A description of the kind of document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/kind
func (c_ CSSearchableItemAttributeSet) SetKind(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKind:"), objc.String(value))
}


// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/languages
func (c_ CSSearchableItemAttributeSet) Languages() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("languages"))
	return rv
}


// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/languages
func (c_ CSSearchableItemAttributeSet) SetLanguages(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguages:"), objc.String(value))
}


// The date on which the file was last used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lastuseddate
func (c_ CSSearchableItemAttributeSet) LastUsedDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("lastUsedDate"))
	return rv
}


// The date on which the file was last used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lastuseddate
func (c_ CSSearchableItemAttributeSet) SetLastUsedDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLastUsedDate:"), value)
}


// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/latitude
func (c_ CSSearchableItemAttributeSet) Latitude() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("latitude"))
	return rv
}


// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/latitude
func (c_ CSSearchableItemAttributeSet) SetLatitude(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLatitude:"), value)
}


// An array that contains the names of the various layers in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/layernames
func (c_ CSSearchableItemAttributeSet) LayerNames() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("layerNames"))
	return rv
}


// An array that contains the names of the various layers in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/layernames
func (c_ CSSearchableItemAttributeSet) SetLayerNames(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLayerNames:"), objc.String(value))
}


// The model of the lens that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lensmodel
func (c_ CSSearchableItemAttributeSet) LensModel() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("lensModel"))
	return rv
}


// The model of the lens that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lensmodel
func (c_ CSSearchableItemAttributeSet) SetLensModel(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLensModel:"), objc.String(value))
}


// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/likelyjunk
func (c_ CSSearchableItemAttributeSet) LikelyJunk() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("likelyJunk"))
	return rv
}


// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/likelyjunk
func (c_ CSSearchableItemAttributeSet) SetLikelyJunk(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLikelyJunk:"), value)
}


// A value that indicates if the media is local.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/local
func (c_ CSSearchableItemAttributeSet) Local() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("local"))
	return rv
}


// A value that indicates if the media is local.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/local
func (c_ CSSearchableItemAttributeSet) SetLocal(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocal:"), value)
}


// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/longitude
func (c_ CSSearchableItemAttributeSet) Longitude() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("longitude"))
	return rv
}


// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/longitude
func (c_ CSSearchableItemAttributeSet) SetLongitude(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongitude:"), value)
}


// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lyricist
func (c_ CSSearchableItemAttributeSet) Lyricist() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("lyricist"))
	return rv
}


// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lyricist
func (c_ CSSearchableItemAttributeSet) SetLyricist(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLyricist:"), objc.String(value))
}


// An array of mailbox identifiers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mailboxidentifiers
func (c_ CSSearchableItemAttributeSet) MailboxIdentifiers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("mailboxIdentifiers"))
	return rv
}


// An array of mailbox identifiers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mailboxidentifiers
func (c_ CSSearchableItemAttributeSet) SetMailboxIdentifiers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMailboxIdentifiers:"), objc.String(value))
}


// The smallest F number of the lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/maxaperture
func (c_ CSSearchableItemAttributeSet) MaxAperture() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("maxAperture"))
	return rv
}


// The smallest F number of the lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/maxaperture
func (c_ CSSearchableItemAttributeSet) SetMaxAperture(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxAperture:"), value)
}


// The media types present in the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mediatypes
func (c_ CSSearchableItemAttributeSet) MediaTypes() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("mediaTypes"))
	return rv
}


// The media types present in the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mediatypes
func (c_ CSSearchableItemAttributeSet) SetMediaTypes(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMediaTypes:"), objc.String(value))
}


// The date on which the last metadata attribute was changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/metadatamodificationdate
func (c_ CSSearchableItemAttributeSet) MetadataModificationDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("metadataModificationDate"))
	return rv
}


// The date on which the last metadata attribute was changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/metadatamodificationdate
func (c_ CSSearchableItemAttributeSet) SetMetadataModificationDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataModificationDate:"), value)
}


// The metering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/meteringmode
func (c_ CSSearchableItemAttributeSet) MeteringMode() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("meteringMode"))
	return rv
}


// The metering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/meteringmode
func (c_ CSSearchableItemAttributeSet) SetMeteringMode(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMeteringMode:"), objc.String(value))
}


// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalgenre
func (c_ CSSearchableItemAttributeSet) MusicalGenre() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("musicalGenre"))
	return rv
}


// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalgenre
func (c_ CSSearchableItemAttributeSet) SetMusicalGenre(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalGenre:"), objc.String(value))
}


// The category of the instrument associated with the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentcategory
func (c_ CSSearchableItemAttributeSet) MusicalInstrumentCategory() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("musicalInstrumentCategory"))
	return rv
}


// The category of the instrument associated with the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentcategory
func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentCategory(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentCategory:"), objc.String(value))
}


// The name of an instrument within the context of an instrument category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentname
func (c_ CSSearchableItemAttributeSet) MusicalInstrumentName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("musicalInstrumentName"))
	return rv
}


// The name of an instrument within the context of an instrument category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentname
func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentName:"), objc.String(value))
}


// The name of the location or point of interest associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/namedlocation
func (c_ CSSearchableItemAttributeSet) NamedLocation() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("namedLocation"))
	return rv
}


// The name of the location or point of interest associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/namedlocation
func (c_ CSSearchableItemAttributeSet) SetNamedLocation(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNamedLocation:"), objc.String(value))
}


// A list of companies or organizations that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/organizations
func (c_ CSSearchableItemAttributeSet) Organizations() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("organizations"))
	return rv
}


// A list of companies or organizations that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/organizations
func (c_ CSSearchableItemAttributeSet) SetOrganizations(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrganizations:"), objc.String(value))
}


// The orientation of the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/orientation
func (c_ CSSearchableItemAttributeSet) Orientation() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("orientation"))
	return rv
}


// The orientation of the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/orientation
func (c_ CSSearchableItemAttributeSet) SetOrientation(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrientation:"), value)
}


// The original format of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalformat
func (c_ CSSearchableItemAttributeSet) OriginalFormat() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("originalFormat"))
	return rv
}


// The original format of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalformat
func (c_ CSSearchableItemAttributeSet) SetOriginalFormat(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalFormat:"), objc.String(value))
}


// The original source of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalsource
func (c_ CSSearchableItemAttributeSet) OriginalSource() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("originalSource"))
	return rv
}


// The original source of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalsource
func (c_ CSSearchableItemAttributeSet) SetOriginalSource(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalSource:"), objc.String(value))
}


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagecount
func (c_ CSSearchableItemAttributeSet) PageCount() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pageCount"))
	return rv
}


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagecount
func (c_ CSSearchableItemAttributeSet) SetPageCount(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageCount:"), value)
}


// The height of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pageheight
func (c_ CSSearchableItemAttributeSet) PageHeight() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pageHeight"))
	return rv
}


// The height of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pageheight
func (c_ CSSearchableItemAttributeSet) SetPageHeight(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageHeight:"), value)
}


// The width of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagewidth
func (c_ CSSearchableItemAttributeSet) PageWidth() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pageWidth"))
	return rv
}


// The width of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagewidth
func (c_ CSSearchableItemAttributeSet) SetPageWidth(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageWidth:"), value)
}


// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/participants
func (c_ CSSearchableItemAttributeSet) Participants() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("participants"))
	return rv
}


// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/participants
func (c_ CSSearchableItemAttributeSet) SetParticipants(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParticipants:"), objc.String(value))
}


// The complete path to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/path
func (c_ CSSearchableItemAttributeSet) Path() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("path"))
	return rv
}


// The complete path to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/path
func (c_ CSSearchableItemAttributeSet) SetPath(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPath:"), objc.String(value))
}


// A list of performers in the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/performers
func (c_ CSSearchableItemAttributeSet) Performers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("performers"))
	return rv
}


// A list of performers in the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/performers
func (c_ CSSearchableItemAttributeSet) SetPerformers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerformers:"), objc.String(value))
}


// An array of phone numbers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/phonenumbers
func (c_ CSSearchableItemAttributeSet) PhoneNumbers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}


// An array of phone numbers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/phonenumbers
func (c_ CSSearchableItemAttributeSet) SetPhoneNumbers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneNumbers:"), objc.String(value))
}


// The total number of pixels in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelcount
func (c_ CSSearchableItemAttributeSet) PixelCount() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pixelCount"))
	return rv
}


// The total number of pixels in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelcount
func (c_ CSSearchableItemAttributeSet) SetPixelCount(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelCount:"), value)
}


// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelheight
func (c_ CSSearchableItemAttributeSet) PixelHeight() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pixelHeight"))
	return rv
}


// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelheight
func (c_ CSSearchableItemAttributeSet) SetPixelHeight(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelHeight:"), value)
}


// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelwidth
func (c_ CSSearchableItemAttributeSet) PixelWidth() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("pixelWidth"))
	return rv
}


// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelwidth
func (c_ CSSearchableItemAttributeSet) SetPixelWidth(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelWidth:"), value)
}


// A user-supplied play count for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/playcount
func (c_ CSSearchableItemAttributeSet) PlayCount() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("playCount"))
	return rv
}


// A user-supplied play count for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/playcount
func (c_ CSSearchableItemAttributeSet) SetPlayCount(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayCount:"), value)
}


// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/postalcode
func (c_ CSSearchableItemAttributeSet) PostalCode() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("postalCode"))
	return rv
}


// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/postalcode
func (c_ CSSearchableItemAttributeSet) SetPostalCode(value string /* primitive/slice/pointer. */) {
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
func (c_ CSSearchableItemAttributeSet) Producer() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("producer"))
	return rv
}


// The producer of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/producer
func (c_ CSSearchableItemAttributeSet) SetProducer(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProducer:"), objc.String(value))
}


// The name of the color profile the camera used for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/profilename
func (c_ CSSearchableItemAttributeSet) ProfileName() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("profileName"))
	return rv
}


// The name of the color profile the camera used for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/profilename
func (c_ CSSearchableItemAttributeSet) SetProfileName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProfileName:"), objc.String(value))
}


// A list of projects of which this file is a part.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/projects
func (c_ CSSearchableItemAttributeSet) Projects() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("projects"))
	return rv
}


// A list of projects of which this file is a part.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/projects
func (c_ CSSearchableItemAttributeSet) SetProjects(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProjects:"), objc.String(value))
}


// An array of identifiers that corresponds to data representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerdatatypeidentifiers
func (c_ CSSearchableItemAttributeSet) ProviderDataTypeIdentifiers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("providerDataTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to data representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerdatatypeidentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderDataTypeIdentifiers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderDataTypeIdentifiers:"), objc.String(value))
}


// An array of identifiers that corresponds to file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerfiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) ProviderFileTypeIdentifiers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("providerFileTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerfiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderFileTypeIdentifiers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderFileTypeIdentifiers:"), objc.String(value))
}


// An array of identifiers that corresponds to in-place file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerinplacefiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) ProviderInPlaceFileTypeIdentifiers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("providerInPlaceFileTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to in-place file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerinplacefiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderInPlaceFileTypeIdentifiers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderInPlaceFileTypeIdentifiers:"), objc.String(value))
}


// A list of people, organizations, services, or other entities responsible for making the media available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/publishers
func (c_ CSSearchableItemAttributeSet) Publishers() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("publishers"))
	return rv
}


// A list of people, organizations, services, or other entities responsible for making the media available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/publishers
func (c_ CSSearchableItemAttributeSet) SetPublishers(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPublishers:"), objc.String(value))
}


// A number that indicates the relative importance of the item among other items from the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rankinghint
func (c_ CSSearchableItemAttributeSet) RankingHint() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("rankingHint"))
	return rv
}


// A number that indicates the relative importance of the item among other items from the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rankinghint
func (c_ CSSearchableItemAttributeSet) SetRankingHint(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRankingHint:"), value)
}


// The user-supplied rating of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rating
func (c_ CSSearchableItemAttributeSet) Rating() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("rating"))
	return rv
}


// The user-supplied rating of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rating
func (c_ CSSearchableItemAttributeSet) SetRating(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRating:"), value)
}


// A description of the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ratingdescription
func (c_ CSSearchableItemAttributeSet) RatingDescription() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("ratingDescription"))
	return rv
}


// A description of the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ratingdescription
func (c_ CSSearchableItemAttributeSet) SetRatingDescription(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRatingDescription:"), objc.String(value))
}


// An array of addresses associated with the recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientaddresses
func (c_ CSSearchableItemAttributeSet) RecipientAddresses() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("recipientAddresses"))
	return rv
}


// An array of addresses associated with the recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientaddresses
func (c_ CSSearchableItemAttributeSet) SetRecipientAddresses(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientAddresses:"), objc.String(value))
}


// An array of email addresses associated with the recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientemailaddresses
func (c_ CSSearchableItemAttributeSet) RecipientEmailAddresses() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("recipientEmailAddresses"))
	return rv
}


// An array of email addresses associated with the recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientemailaddresses
func (c_ CSSearchableItemAttributeSet) SetRecipientEmailAddresses(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientEmailAddresses:"), objc.String(value))
}


// An array of names representing the recipients of this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientnames
func (c_ CSSearchableItemAttributeSet) RecipientNames() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("recipientNames"))
	return rv
}


// An array of names representing the recipients of this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientnames
func (c_ CSSearchableItemAttributeSet) SetRecipientNames(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientNames:"), objc.String(value))
}


// The recording date of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recordingdate
func (c_ CSSearchableItemAttributeSet) RecordingDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("recordingDate"))
	return rv
}


// The recording date of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recordingdate
func (c_ CSSearchableItemAttributeSet) SetRecordingDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordingDate:"), value)
}


// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/redeyeon
func (c_ CSSearchableItemAttributeSet) RedEyeOn() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("redEyeOn"))
	return rv
}


// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/redeyeon
func (c_ CSSearchableItemAttributeSet) SetRedEyeOn(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRedEyeOn:"), value)
}


// The unique identifier for the item to which the activity is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/relateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) RelatedUniqueIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("relatedUniqueIdentifier"))
	return rv
}


// The unique identifier for the item to which the activity is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/relateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) SetRelatedUniqueIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRelatedUniqueIdentifier:"), objc.String(value))
}


// The resolution height of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionheightdpi
func (c_ CSSearchableItemAttributeSet) ResolutionHeightDPI() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("resolutionHeightDPI"))
	return rv
}


// The resolution height of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionheightdpi
func (c_ CSSearchableItemAttributeSet) SetResolutionHeightDPI(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionHeightDPI:"), value)
}


// The resolution width of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionwidthdpi
func (c_ CSSearchableItemAttributeSet) ResolutionWidthDPI() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("resolutionWidthDPI"))
	return rv
}


// The resolution width of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionwidthdpi
func (c_ CSSearchableItemAttributeSet) SetResolutionWidthDPI(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionWidthDPI:"), value)
}


// A link to information about the rights held in and over the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rights
func (c_ CSSearchableItemAttributeSet) Rights() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("rights"))
	return rv
}


// A link to information about the rights held in and over the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rights
func (c_ CSSearchableItemAttributeSet) SetRights(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRights:"), objc.String(value))
}


// Indicates the role of the content creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/role
func (c_ CSSearchableItemAttributeSet) Role() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("role"))
	return rv
}


// Indicates the role of the content creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/role
func (c_ CSSearchableItemAttributeSet) SetRole(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRole:"), objc.String(value))
}


// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/securitymethod
func (c_ CSSearchableItemAttributeSet) SecurityMethod() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("securityMethod"))
	return rv
}


// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/securitymethod
func (c_ CSSearchableItemAttributeSet) SetSecurityMethod(value string /* primitive/slice/pointer. */) {
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
func (c_ CSSearchableItemAttributeSet) Speed() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("speed"))
	return rv
}


// The speed of the item, in kilometers per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/speed
func (c_ CSSearchableItemAttributeSet) SetSpeed(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpeed:"), value)
}


// The start date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/startdate
func (c_ CSSearchableItemAttributeSet) StartDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("startDate"))
	return rv
}


// The start date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/startdate
func (c_ CSSearchableItemAttributeSet) SetStartDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStartDate:"), value)
}


// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/stateorprovince
func (c_ CSSearchableItemAttributeSet) StateOrProvince() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("stateOrProvince"))
	return rv
}


// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/stateorprovince
func (c_ CSSearchableItemAttributeSet) SetStateOrProvince(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStateOrProvince:"), objc.String(value))
}


// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/streamable
func (c_ CSSearchableItemAttributeSet) Streamable() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("streamable"))
	return rv
}


// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/streamable
func (c_ CSSearchableItemAttributeSet) SetStreamable(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreamable:"), value)
}


// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subthoroughfare
func (c_ CSSearchableItemAttributeSet) SubThoroughfare() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("subThoroughfare"))
	return rv
}


// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subthoroughfare
func (c_ CSSearchableItemAttributeSet) SetSubThoroughfare(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubThoroughfare:"), objc.String(value))
}


// The subject of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subject
func (c_ CSSearchableItemAttributeSet) Subject() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("subject"))
	return rv
}


// The subject of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subject
func (c_ CSSearchableItemAttributeSet) SetSubject(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubject:"), objc.String(value))
}


// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/supportsphonecall
func (c_ CSSearchableItemAttributeSet) SupportsPhoneCall() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("supportsPhoneCall"))
	return rv
}


// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/supportsphonecall
func (c_ CSSearchableItemAttributeSet) SetSupportsPhoneCall(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsPhoneCall:"), value)
}


// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/tempo
func (c_ CSSearchableItemAttributeSet) Tempo() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("tempo"))
	return rv
}


// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/tempo
func (c_ CSSearchableItemAttributeSet) SetTempo(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTempo:"), value)
}


// The textual content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontent
func (c_ CSSearchableItemAttributeSet) TextContent() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("textContent"))
	return rv
}


// The textual content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontent
func (c_ CSSearchableItemAttributeSet) SetTextContent(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextContent:"), objc.String(value))
}


// A string that presents the Apple Intelligence summarization of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontentsummary
func (c_ CSSearchableItemAttributeSet) TextContentSummary() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("textContentSummary"))
	return rv
}


// A string that presents the Apple Intelligence summarization of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontentsummary
func (c_ CSSearchableItemAttributeSet) SetTextContentSummary(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextContentSummary:"), objc.String(value))
}


// The theme of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/theme
func (c_ CSSearchableItemAttributeSet) Theme() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("theme"))
	return rv
}


// The theme of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/theme
func (c_ CSSearchableItemAttributeSet) SetTheme(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTheme:"), objc.String(value))
}


// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thoroughfare
func (c_ CSSearchableItemAttributeSet) Thoroughfare() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("thoroughfare"))
	return rv
}


// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thoroughfare
func (c_ CSSearchableItemAttributeSet) SetThoroughfare(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThoroughfare:"), objc.String(value))
}


// Image data that represents the thumbnail of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnaildata
func (c_ CSSearchableItemAttributeSet) ThumbnailData() foundation.objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("thumbnailData"))
	return rv
}


// Image data that represents the thumbnail of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnaildata
func (c_ CSSearchableItemAttributeSet) SetThumbnailData(value foundation.objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailData:"), value)
}


// The local file URL of the thumbnail image for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnailurl
func (c_ CSSearchableItemAttributeSet) ThumbnailURL() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("thumbnailURL"))
	return rv
}


// The local file URL of the thumbnail image for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnailurl
func (c_ CSSearchableItemAttributeSet) SetThumbnailURL(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailURL:"), value)
}


// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timesignature
func (c_ CSSearchableItemAttributeSet) TimeSignature() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("timeSignature"))
	return rv
}


// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timesignature
func (c_ CSSearchableItemAttributeSet) SetTimeSignature(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeSignature:"), objc.String(value))
}


// The timestamp on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timestamp
func (c_ CSSearchableItemAttributeSet) Timestamp() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("timestamp"))
	return rv
}


// The timestamp on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timestamp
func (c_ CSSearchableItemAttributeSet) SetTimestamp(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimestamp:"), value)
}


// The total bit rate of the media, combining audio and video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/totalbitrate
func (c_ CSSearchableItemAttributeSet) TotalBitRate() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("totalBitRate"))
	return rv
}


// The total bit rate of the media, combining audio and video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/totalbitrate
func (c_ CSSearchableItemAttributeSet) SetTotalBitRate(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalBitRate:"), value)
}


// A string that represents the text the system transcribed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/transcribedtextcontent
func (c_ CSSearchableItemAttributeSet) TranscribedTextContent() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("transcribedTextContent"))
	return rv
}


// A string that represents the text the system transcribed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/transcribedtextcontent
func (c_ CSSearchableItemAttributeSet) SetTranscribedTextContent(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTranscribedTextContent:"), objc.String(value))
}


// The URL associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/url
func (c_ CSSearchableItemAttributeSet) Url() foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("url"))
	return rv
}


// The URL associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/url
func (c_ CSSearchableItemAttributeSet) SetUrl(value foundation.objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrl:"), value)
}


// A value that indicates the user created the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercreated
func (c_ CSSearchableItemAttributeSet) UserCreated() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("userCreated"))
	return rv
}


// A value that indicates the user created the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercreated
func (c_ CSSearchableItemAttributeSet) SetUserCreated(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCreated:"), value)
}


// A value that indicates the user selected the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercurated
func (c_ CSSearchableItemAttributeSet) UserCurated() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("userCurated"))
	return rv
}


// A value that indicates the user selected the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercurated
func (c_ CSSearchableItemAttributeSet) SetUserCurated(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCurated:"), value)
}


// A value that indicates the user purchased or owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/userowned
func (c_ CSSearchableItemAttributeSet) UserOwned() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("userOwned"))
	return rv
}


// A value that indicates the user purchased or owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/userowned
func (c_ CSSearchableItemAttributeSet) SetUserOwned(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserOwned:"), value)
}


// A version string associated with the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/version
func (c_ CSSearchableItemAttributeSet) Version() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("version"))
	return rv
}


// A version string associated with the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/version
func (c_ CSSearchableItemAttributeSet) SetVersion(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVersion:"), objc.String(value))
}


// The video bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/videobitrate
func (c_ CSSearchableItemAttributeSet) VideoBitRate() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("videoBitRate"))
	return rv
}


// The video bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/videobitrate
func (c_ CSSearchableItemAttributeSet) SetVideoBitRate(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoBitRate:"), value)
}


// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/weakrelateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) WeakRelatedUniqueIdentifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("weakRelatedUniqueIdentifier"))
	return rv
}


// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/weakrelateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) SetWeakRelatedUniqueIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeakRelatedUniqueIdentifier:"), objc.String(value))
}


// The white balance setting when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/whitebalance
func (c_ CSSearchableItemAttributeSet) WhiteBalance() foundation.objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("whiteBalance"))
	return rv
}


// The white balance setting when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/whitebalance
func (c_ CSSearchableItemAttributeSet) SetWhiteBalance(value foundation.objc.IObject /* cross-framework: Number */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteBalance:"), value)
}



