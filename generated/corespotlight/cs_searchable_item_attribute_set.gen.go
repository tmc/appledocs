// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

/* debug [class.gen.go]: Generating class CSSearchableItemAttributeSet */


/* debug [class_header]: Header for CSSearchableItemAttributeSet */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSSearchableItemAttributeSet */
// An interface definition for the [CSSearchableItemAttributeSet] class.
type ICSSearchableItemAttributeSet interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CSSearchableItemAttributeSet */
	// properties:
	Authors() ICSPerson
	SetAuthors(value ICSPerson)
	HiddenAdditionalRecipients() ICSPerson
	SetHiddenAdditionalRecipients(value ICSPerson)
	AdditionalRecipients() ICSPerson
	SetAdditionalRecipients(value ICSPerson)
	PrimaryRecipients() ICSPerson
	SetPrimaryRecipients(value ICSPerson)
	AccountHandles() []string
	SetAccountHandles(value []string)
	AccountIdentifier() objc.IObject /* cross-framework: NSString */
	SetAccountIdentifier(value objc.IObject /* cross-framework: NSString */)
	AcquisitionMake() objc.IObject /* cross-framework: NSString */
	SetAcquisitionMake(value objc.IObject /* cross-framework: NSString */)
	AcquisitionModel() objc.IObject /* cross-framework: NSString */
	SetAcquisitionModel(value objc.IObject /* cross-framework: NSString */)
	AddedDate() objc.IObject /* cross-framework: NSDate */
	SetAddedDate(value objc.IObject /* cross-framework: NSDate */)
	Album() objc.IObject /* cross-framework: NSString */
	SetAlbum(value objc.IObject /* cross-framework: NSString */)
	AllDay() objc.IObject /* cross-framework: NSNumber */
	SetAllDay(value objc.IObject /* cross-framework: NSNumber */)
	AlternateNames() []string
	SetAlternateNames(value []string)
	Altitude() objc.IObject /* cross-framework: NSNumber */
	SetAltitude(value objc.IObject /* cross-framework: NSNumber */)
	Aperture() objc.IObject /* cross-framework: NSNumber */
	SetAperture(value objc.IObject /* cross-framework: NSNumber */)
	Artist() objc.IObject /* cross-framework: NSString */
	SetArtist(value objc.IObject /* cross-framework: NSString */)
	Audiences() []string
	SetAudiences(value []string)
	AudioBitRate() objc.IObject /* cross-framework: NSNumber */
	SetAudioBitRate(value objc.IObject /* cross-framework: NSNumber */)
	AudioChannelCount() objc.IObject /* cross-framework: NSNumber */
	SetAudioChannelCount(value objc.IObject /* cross-framework: NSNumber */)
	AudioEncodingApplication() objc.IObject /* cross-framework: NSString */
	SetAudioEncodingApplication(value objc.IObject /* cross-framework: NSString */)
	AudioSampleRate() objc.IObject /* cross-framework: NSNumber */
	SetAudioSampleRate(value objc.IObject /* cross-framework: NSNumber */)
	AudioTrackNumber() objc.IObject /* cross-framework: NSNumber */
	SetAudioTrackNumber(value objc.IObject /* cross-framework: NSNumber */)
	AuthorAddresses() []string
	SetAuthorAddresses(value []string)
	AuthorEmailAddresses() []string
	SetAuthorEmailAddresses(value []string)
	AuthorNames() []string
	SetAuthorNames(value []string)
	BitsPerSample() objc.IObject /* cross-framework: NSNumber */
	SetBitsPerSample(value objc.IObject /* cross-framework: NSNumber */)
	CameraOwner() objc.IObject /* cross-framework: NSString */
	SetCameraOwner(value objc.IObject /* cross-framework: NSString */)
	City() objc.IObject /* cross-framework: NSString */
	SetCity(value objc.IObject /* cross-framework: NSString */)
	Codecs() []string
	SetCodecs(value []string)
	ColorSpace() objc.IObject /* cross-framework: NSString */
	SetColorSpace(value objc.IObject /* cross-framework: NSString */)
	Comment() objc.IObject /* cross-framework: NSString */
	SetComment(value objc.IObject /* cross-framework: NSString */)
	CompletionDate() objc.IObject /* cross-framework: NSDate */
	SetCompletionDate(value objc.IObject /* cross-framework: NSDate */)
	ContactKeywords() []string
	SetContactKeywords(value []string)
	ContainerDisplayName() objc.IObject /* cross-framework: NSString */
	SetContainerDisplayName(value objc.IObject /* cross-framework: NSString */)
	ContainerIdentifier() objc.IObject /* cross-framework: NSString */
	SetContainerIdentifier(value objc.IObject /* cross-framework: NSString */)
	ContainerOrder() objc.IObject /* cross-framework: NSNumber */
	SetContainerOrder(value objc.IObject /* cross-framework: NSNumber */)
	ContainerTitle() objc.IObject /* cross-framework: NSString */
	SetContainerTitle(value objc.IObject /* cross-framework: NSString */)
	ContentCreationDate() objc.IObject /* cross-framework: NSDate */
	SetContentCreationDate(value objc.IObject /* cross-framework: NSDate */)
	ContentDescription() objc.IObject /* cross-framework: NSString */
	SetContentDescription(value objc.IObject /* cross-framework: NSString */)
	ContentModificationDate() objc.IObject /* cross-framework: NSDate */
	SetContentModificationDate(value objc.IObject /* cross-framework: NSDate */)
	ContentRating() objc.IObject /* cross-framework: NSNumber */
	SetContentRating(value objc.IObject /* cross-framework: NSNumber */)
	ContentSources() []string
	SetContentSources(value []string)
	ContentType() objc.IObject /* cross-framework: NSString */
	SetContentType(value objc.IObject /* cross-framework: NSString */)
	ContentTypeTree() []string
	SetContentTypeTree(value []string)
	ContentURL() objc.IObject /* cross-framework: NSURL */
	SetContentURL(value objc.IObject /* cross-framework: NSURL */)
	Contributors() []string
	SetContributors(value []string)
	Copyright() objc.IObject /* cross-framework: NSString */
	SetCopyright(value objc.IObject /* cross-framework: NSString */)
	Country() objc.IObject /* cross-framework: NSString */
	SetCountry(value objc.IObject /* cross-framework: NSString */)
	Coverage() []string
	SetCoverage(value []string)
	Creator() objc.IObject /* cross-framework: NSString */
	SetCreator(value objc.IObject /* cross-framework: NSString */)
	DarkThumbnailURL() objc.IObject /* cross-framework: NSURL */
	SetDarkThumbnailURL(value objc.IObject /* cross-framework: NSURL */)
	DeliveryType() objc.IObject /* cross-framework: NSNumber */
	SetDeliveryType(value objc.IObject /* cross-framework: NSNumber */)
	Director() objc.IObject /* cross-framework: NSString */
	SetDirector(value objc.IObject /* cross-framework: NSString */)
	DisplayName() objc.IObject /* cross-framework: NSString */
	SetDisplayName(value objc.IObject /* cross-framework: NSString */)
	DomainIdentifier() objc.IObject /* cross-framework: NSString */
	SetDomainIdentifier(value objc.IObject /* cross-framework: NSString */)
	DownloadedDate() objc.IObject /* cross-framework: NSDate */
	SetDownloadedDate(value objc.IObject /* cross-framework: NSDate */)
	DueDate() objc.IObject /* cross-framework: NSDate */
	SetDueDate(value objc.IObject /* cross-framework: NSDate */)
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	Editors() []string
	SetEditors(value []string)
	EmailAddresses() []string
	SetEmailAddresses(value []string)
	EmailHeaders() foundation.IDictionary
	SetEmailHeaders(value foundation.IDictionary)
	EncodingApplications() []string
	SetEncodingApplications(value []string)
	EndDate() objc.IObject /* cross-framework: NSDate */
	SetEndDate(value objc.IObject /* cross-framework: NSDate */)
	EXIFGPSVersion() objc.IObject /* cross-framework: NSString */
	SetEXIFGPSVersion(value objc.IObject /* cross-framework: NSString */)
	EXIFVersion() objc.IObject /* cross-framework: NSString */
	SetEXIFVersion(value objc.IObject /* cross-framework: NSString */)
	ExposureMode() objc.IObject /* cross-framework: NSNumber */
	SetExposureMode(value objc.IObject /* cross-framework: NSNumber */)
	ExposureProgram() objc.IObject /* cross-framework: NSString */
	SetExposureProgram(value objc.IObject /* cross-framework: NSString */)
	ExposureTime() objc.IObject /* cross-framework: NSNumber */
	SetExposureTime(value objc.IObject /* cross-framework: NSNumber */)
	ExposureTimeString() objc.IObject /* cross-framework: NSString */
	SetExposureTimeString(value objc.IObject /* cross-framework: NSString */)
	FileSize() objc.IObject /* cross-framework: NSNumber */
	SetFileSize(value objc.IObject /* cross-framework: NSNumber */)
	FlashOn() objc.IObject /* cross-framework: NSNumber */
	SetFlashOn(value objc.IObject /* cross-framework: NSNumber */)
	FNumber() objc.IObject /* cross-framework: NSNumber */
	SetFNumber(value objc.IObject /* cross-framework: NSNumber */)
	FocalLength() objc.IObject /* cross-framework: NSNumber */
	SetFocalLength(value objc.IObject /* cross-framework: NSNumber */)
	FocalLength35mm() objc.IObject /* cross-framework: NSNumber */
	SetFocalLength35mm(value objc.IObject /* cross-framework: NSNumber */)
	FontNames() []string
	SetFontNames(value []string)
	FullyFormattedAddress() objc.IObject /* cross-framework: NSString */
	SetFullyFormattedAddress(value objc.IObject /* cross-framework: NSString */)
	GeneralMIDISequence() objc.IObject /* cross-framework: NSNumber */
	SetGeneralMIDISequence(value objc.IObject /* cross-framework: NSNumber */)
	Genre() objc.IObject /* cross-framework: NSString */
	SetGenre(value objc.IObject /* cross-framework: NSString */)
	GPSAreaInformation() objc.IObject /* cross-framework: NSString */
	SetGPSAreaInformation(value objc.IObject /* cross-framework: NSString */)
	GPSDateStamp() objc.IObject /* cross-framework: NSDate */
	SetGPSDateStamp(value objc.IObject /* cross-framework: NSDate */)
	GPSDestBearing() objc.IObject /* cross-framework: NSNumber */
	SetGPSDestBearing(value objc.IObject /* cross-framework: NSNumber */)
	GPSDestDistance() objc.IObject /* cross-framework: NSNumber */
	SetGPSDestDistance(value objc.IObject /* cross-framework: NSNumber */)
	GPSDestLatitude() objc.IObject /* cross-framework: NSNumber */
	SetGPSDestLatitude(value objc.IObject /* cross-framework: NSNumber */)
	GPSDestLongitude() objc.IObject /* cross-framework: NSNumber */
	SetGPSDestLongitude(value objc.IObject /* cross-framework: NSNumber */)
	GPSDifferental() objc.IObject /* cross-framework: NSNumber */
	SetGPSDifferental(value objc.IObject /* cross-framework: NSNumber */)
	GPSDOP() objc.IObject /* cross-framework: NSNumber */
	SetGPSDOP(value objc.IObject /* cross-framework: NSNumber */)
	GPSMapDatum() objc.IObject /* cross-framework: NSString */
	SetGPSMapDatum(value objc.IObject /* cross-framework: NSString */)
	GPSMeasureMode() objc.IObject /* cross-framework: NSString */
	SetGPSMeasureMode(value objc.IObject /* cross-framework: NSString */)
	GPSProcessingMethod() objc.IObject /* cross-framework: NSString */
	SetGPSProcessingMethod(value objc.IObject /* cross-framework: NSString */)
	GPSStatus() objc.IObject /* cross-framework: NSString */
	SetGPSStatus(value objc.IObject /* cross-framework: NSString */)
	GPSTrack() objc.IObject /* cross-framework: NSNumber */
	SetGPSTrack(value objc.IObject /* cross-framework: NSNumber */)
	HasAlphaChannel() objc.IObject /* cross-framework: NSNumber */
	SetHasAlphaChannel(value objc.IObject /* cross-framework: NSNumber */)
	Headline() objc.IObject /* cross-framework: NSString */
	SetHeadline(value objc.IObject /* cross-framework: NSString */)
	HTMLContentData() objc.IObject /* cross-framework: NSData */
	SetHTMLContentData(value objc.IObject /* cross-framework: NSData */)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	ImageDirection() objc.IObject /* cross-framework: NSNumber */
	SetImageDirection(value objc.IObject /* cross-framework: NSNumber */)
	ImportantDates() []foundation.Date
	SetImportantDates(value []foundation.Date)
	Information() objc.IObject /* cross-framework: NSString */
	SetInformation(value objc.IObject /* cross-framework: NSString */)
	InstantMessageAddresses() []string
	SetInstantMessageAddresses(value []string)
	Instructions() objc.IObject /* cross-framework: NSString */
	SetInstructions(value objc.IObject /* cross-framework: NSString */)
	ISOSpeed() objc.IObject /* cross-framework: NSNumber */
	SetISOSpeed(value objc.IObject /* cross-framework: NSNumber */)
	IsPriority() objc.IObject /* cross-framework: NSNumber */
	KeySignature() objc.IObject /* cross-framework: NSString */
	SetKeySignature(value objc.IObject /* cross-framework: NSString */)
	Keywords() []string
	SetKeywords(value []string)
	Kind() objc.IObject /* cross-framework: NSString */
	SetKind(value objc.IObject /* cross-framework: NSString */)
	Languages() []string
	SetLanguages(value []string)
	LastUsedDate() objc.IObject /* cross-framework: NSDate */
	SetLastUsedDate(value objc.IObject /* cross-framework: NSDate */)
	Latitude() objc.IObject /* cross-framework: NSNumber */
	SetLatitude(value objc.IObject /* cross-framework: NSNumber */)
	LayerNames() []string
	SetLayerNames(value []string)
	LensModel() objc.IObject /* cross-framework: NSString */
	SetLensModel(value objc.IObject /* cross-framework: NSString */)
	LikelyJunk() objc.IObject /* cross-framework: NSNumber */
	SetLikelyJunk(value objc.IObject /* cross-framework: NSNumber */)
	Local() objc.IObject /* cross-framework: NSNumber */
	SetLocal(value objc.IObject /* cross-framework: NSNumber */)
	Longitude() objc.IObject /* cross-framework: NSNumber */
	SetLongitude(value objc.IObject /* cross-framework: NSNumber */)
	Lyricist() objc.IObject /* cross-framework: NSString */
	SetLyricist(value objc.IObject /* cross-framework: NSString */)
	MailboxIdentifiers() []string
	SetMailboxIdentifiers(value []string)
	MaxAperture() objc.IObject /* cross-framework: NSNumber */
	SetMaxAperture(value objc.IObject /* cross-framework: NSNumber */)
	MediaTypes() []string
	SetMediaTypes(value []string)
	MetadataModificationDate() objc.IObject /* cross-framework: NSDate */
	SetMetadataModificationDate(value objc.IObject /* cross-framework: NSDate */)
	MeteringMode() objc.IObject /* cross-framework: NSString */
	SetMeteringMode(value objc.IObject /* cross-framework: NSString */)
	MusicalGenre() objc.IObject /* cross-framework: NSString */
	SetMusicalGenre(value objc.IObject /* cross-framework: NSString */)
	MusicalInstrumentCategory() objc.IObject /* cross-framework: NSString */
	SetMusicalInstrumentCategory(value objc.IObject /* cross-framework: NSString */)
	MusicalInstrumentName() objc.IObject /* cross-framework: NSString */
	SetMusicalInstrumentName(value objc.IObject /* cross-framework: NSString */)
	NamedLocation() objc.IObject /* cross-framework: NSString */
	SetNamedLocation(value objc.IObject /* cross-framework: NSString */)
	Organizations() []string
	SetOrganizations(value []string)
	Orientation() objc.IObject /* cross-framework: NSNumber */
	SetOrientation(value objc.IObject /* cross-framework: NSNumber */)
	OriginalFormat() objc.IObject /* cross-framework: NSString */
	SetOriginalFormat(value objc.IObject /* cross-framework: NSString */)
	OriginalSource() objc.IObject /* cross-framework: NSString */
	SetOriginalSource(value objc.IObject /* cross-framework: NSString */)
	PageCount() objc.IObject /* cross-framework: NSNumber */
	SetPageCount(value objc.IObject /* cross-framework: NSNumber */)
	PageHeight() objc.IObject /* cross-framework: NSNumber */
	SetPageHeight(value objc.IObject /* cross-framework: NSNumber */)
	PageWidth() objc.IObject /* cross-framework: NSNumber */
	SetPageWidth(value objc.IObject /* cross-framework: NSNumber */)
	Participants() []string
	SetParticipants(value []string)
	Path() objc.IObject /* cross-framework: NSString */
	SetPath(value objc.IObject /* cross-framework: NSString */)
	Performers() []string
	SetPerformers(value []string)
	PhoneNumbers() []string
	SetPhoneNumbers(value []string)
	PixelCount() objc.IObject /* cross-framework: NSNumber */
	SetPixelCount(value objc.IObject /* cross-framework: NSNumber */)
	PixelHeight() objc.IObject /* cross-framework: NSNumber */
	SetPixelHeight(value objc.IObject /* cross-framework: NSNumber */)
	PixelWidth() objc.IObject /* cross-framework: NSNumber */
	SetPixelWidth(value objc.IObject /* cross-framework: NSNumber */)
	PlayCount() objc.IObject /* cross-framework: NSNumber */
	SetPlayCount(value objc.IObject /* cross-framework: NSNumber */)
	PostalCode() objc.IObject /* cross-framework: NSString */
	SetPostalCode(value objc.IObject /* cross-framework: NSString */)
	Producer() objc.IObject /* cross-framework: NSString */
	SetProducer(value objc.IObject /* cross-framework: NSString */)
	ProfileName() objc.IObject /* cross-framework: NSString */
	SetProfileName(value objc.IObject /* cross-framework: NSString */)
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
	RankingHint() objc.IObject /* cross-framework: NSNumber */
	SetRankingHint(value objc.IObject /* cross-framework: NSNumber */)
	Rating() objc.IObject /* cross-framework: NSNumber */
	SetRating(value objc.IObject /* cross-framework: NSNumber */)
	RatingDescription() objc.IObject /* cross-framework: NSString */
	SetRatingDescription(value objc.IObject /* cross-framework: NSString */)
	RecipientAddresses() []string
	SetRecipientAddresses(value []string)
	RecipientEmailAddresses() []string
	SetRecipientEmailAddresses(value []string)
	RecipientNames() []string
	SetRecipientNames(value []string)
	RecordingDate() objc.IObject /* cross-framework: NSDate */
	SetRecordingDate(value objc.IObject /* cross-framework: NSDate */)
	RedEyeOn() objc.IObject /* cross-framework: NSNumber */
	SetRedEyeOn(value objc.IObject /* cross-framework: NSNumber */)
	RelatedUniqueIdentifier() objc.IObject /* cross-framework: NSString */
	SetRelatedUniqueIdentifier(value objc.IObject /* cross-framework: NSString */)
	ResolutionHeightDPI() objc.IObject /* cross-framework: NSNumber */
	SetResolutionHeightDPI(value objc.IObject /* cross-framework: NSNumber */)
	ResolutionWidthDPI() objc.IObject /* cross-framework: NSNumber */
	SetResolutionWidthDPI(value objc.IObject /* cross-framework: NSNumber */)
	Rights() objc.IObject /* cross-framework: NSString */
	SetRights(value objc.IObject /* cross-framework: NSString */)
	Role() objc.IObject /* cross-framework: NSString */
	SetRole(value objc.IObject /* cross-framework: NSString */)
	SecurityMethod() objc.IObject /* cross-framework: NSString */
	SetSecurityMethod(value objc.IObject /* cross-framework: NSString */)
	Speed() objc.IObject /* cross-framework: NSNumber */
	SetSpeed(value objc.IObject /* cross-framework: NSNumber */)
	StartDate() objc.IObject /* cross-framework: NSDate */
	SetStartDate(value objc.IObject /* cross-framework: NSDate */)
	StateOrProvince() objc.IObject /* cross-framework: NSString */
	SetStateOrProvince(value objc.IObject /* cross-framework: NSString */)
	Streamable() objc.IObject /* cross-framework: NSNumber */
	SetStreamable(value objc.IObject /* cross-framework: NSNumber */)
	Subject() objc.IObject /* cross-framework: NSString */
	SetSubject(value objc.IObject /* cross-framework: NSString */)
	SubThoroughfare() objc.IObject /* cross-framework: NSString */
	SetSubThoroughfare(value objc.IObject /* cross-framework: NSString */)
	SupportsNavigation() objc.IObject /* cross-framework: NSNumber */
	SetSupportsNavigation(value objc.IObject /* cross-framework: NSNumber */)
	SupportsPhoneCall() objc.IObject /* cross-framework: NSNumber */
	SetSupportsPhoneCall(value objc.IObject /* cross-framework: NSNumber */)
	Tempo() objc.IObject /* cross-framework: NSNumber */
	SetTempo(value objc.IObject /* cross-framework: NSNumber */)
	TextContent() objc.IObject /* cross-framework: NSString */
	SetTextContent(value objc.IObject /* cross-framework: NSString */)
	TextContentSummary() objc.IObject /* cross-framework: NSString */
	Theme() objc.IObject /* cross-framework: NSString */
	SetTheme(value objc.IObject /* cross-framework: NSString */)
	Thoroughfare() objc.IObject /* cross-framework: NSString */
	SetThoroughfare(value objc.IObject /* cross-framework: NSString */)
	ThumbnailData() objc.IObject /* cross-framework: NSData */
	SetThumbnailData(value objc.IObject /* cross-framework: NSData */)
	ThumbnailURL() objc.IObject /* cross-framework: NSURL */
	SetThumbnailURL(value objc.IObject /* cross-framework: NSURL */)
	TimeSignature() objc.IObject /* cross-framework: NSString */
	SetTimeSignature(value objc.IObject /* cross-framework: NSString */)
	Timestamp() objc.IObject /* cross-framework: NSDate */
	SetTimestamp(value objc.IObject /* cross-framework: NSDate */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	TotalBitRate() objc.IObject /* cross-framework: NSNumber */
	SetTotalBitRate(value objc.IObject /* cross-framework: NSNumber */)
	TranscribedTextContent() objc.IObject /* cross-framework: NSString */
	SetTranscribedTextContent(value objc.IObject /* cross-framework: NSString */)
	URL() objc.IObject /* cross-framework: NSURL */
	SetURL(value objc.IObject /* cross-framework: NSURL */)
	UserCreated() objc.IObject /* cross-framework: NSNumber */
	SetUserCreated(value objc.IObject /* cross-framework: NSNumber */)
	UserCurated() objc.IObject /* cross-framework: NSNumber */
	SetUserCurated(value objc.IObject /* cross-framework: NSNumber */)
	UserOwned() objc.IObject /* cross-framework: NSNumber */
	SetUserOwned(value objc.IObject /* cross-framework: NSNumber */)
	Version() objc.IObject /* cross-framework: NSString */
	SetVersion(value objc.IObject /* cross-framework: NSString */)
	VideoBitRate() objc.IObject /* cross-framework: NSNumber */
	SetVideoBitRate(value objc.IObject /* cross-framework: NSNumber */)
	WeakRelatedUniqueIdentifier() objc.IObject /* cross-framework: NSString */
	SetWeakRelatedUniqueIdentifier(value objc.IObject /* cross-framework: NSString */)
	WhiteBalance() objc.IObject /* cross-framework: NSNumber */
	SetWhiteBalance(value objc.IObject /* cross-framework: NSNumber */)
	CSActionIdentifier() objc.IObject /* cross-framework: NSString */
	Composer() objc.IObject /* cross-framework: NSString */
	SetComposer(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSSearchableItemAttributeSet */
	// methods:
	MoveFrom(sourceAttributeSet ICSSearchableItemAttributeSet)
	SetValueForCustomKey(value unsafe.Pointer, key ICSCustomAttributeKey)
	ValueForCustomKey(key ICSCustomAttributeKey) unsafe.Pointer
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSSearchableItemAttributeSet */
// Alloc allocates a new instance without initialization.
func (cc _CSSearchableItemAttributeSetClass) Alloc() CSSearchableItemAttributeSet {
	rv := objc.Send[CSSearchableItemAttributeSet](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSSearchableItemAttributeSet */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSSearchableItemAttributeSet */

// Creates an attribute set for the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/init(contentType:)
func NewCSSearchableItemAttributeSetWithContentType(contentType uniformtypeidentifiers.UTType) CSSearchableItemAttributeSet {
	instance := getCSSearchableItemAttributeSetClass().Alloc()
	rv := objc.Send[CSSearchableItemAttributeSet](instance.ID, objc.Sel("initWithContentType:"), contentType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSSearchableItemAttributeSetWithContentType */


// Creates an attribute set for the specified content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/init(itemContentType:)
func NewCSSearchableItemAttributeSetWithItemContentType(itemContentType objc.IObject /* cross-framework: NSString */) CSSearchableItemAttributeSet {
	instance := getCSSearchableItemAttributeSetClass().Alloc()
	rv := objc.Send[CSSearchableItemAttributeSet](instance.ID, objc.Sel("initWithItemContentType:"), itemContentType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCSSearchableItemAttributeSetWithItemContentType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSSearchableItemAttributeSet */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSSearchableItemAttributeSet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSSearchableItemAttributeSet */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/move(from:)
func (c_ CSSearchableItemAttributeSet) MoveFrom(sourceAttributeSet ICSSearchableItemAttributeSet) {
	objc.Send[objc.ID](c_.ID, objc.Sel("moveFrom:"), sourceAttributeSet)
}/* debug [instance_methods/method]: MoveFrom */


// Sets the value for a custom attribute key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/setValue(_:forCustomKey:)
func (c_ CSSearchableItemAttributeSet) SetValueForCustomKey(value unsafe.Pointer, key ICSCustomAttributeKey) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setValue:forCustomKey:"), value, key)
}/* debug [instance_methods/method]: SetValueForCustomKey */


// Returns the value associated with the specified custom attribute key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/value(forCustomKey:)
func (c_ CSSearchableItemAttributeSet) ValueForCustomKey(key ICSCustomAttributeKey) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("valueForCustomKey:"), key)
	return rv
}/* debug [instance_methods/method]: ValueForCustomKey */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSSearchableItemAttributeSet */

// An array of objects representing the content of the From: field in an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621608-authors
func (c_ CSSearchableItemAttributeSet) Authors() ICSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("authors"))
	return rv
}/* debug [instance_properties/getter]: authors */


// An array of objects representing the content of the From: field in an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621608-authors
func (c_ CSSearchableItemAttributeSet) SetAuthors(value ICSPerson) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthors:"), value)
}/* debug [instance_properties/setter]: authors */


// An array of objects representing the content of the Bcc: field in an email message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621645-hiddenadditionalrecipients
func (c_ CSSearchableItemAttributeSet) HiddenAdditionalRecipients() ICSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("hiddenAdditionalRecipients"))
	return rv
}/* debug [instance_properties/getter]: hiddenAdditionalRecipients */


// An array of objects representing the content of the Bcc: field in an email message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621645-hiddenadditionalrecipients
func (c_ CSSearchableItemAttributeSet) SetHiddenAdditionalRecipients(value ICSPerson) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHiddenAdditionalRecipients:"), value)
}/* debug [instance_properties/setter]: hiddenAdditionalRecipients */


// An array of objects representing the content of the Cc: field in an email message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621664-additionalrecipients
func (c_ CSSearchableItemAttributeSet) AdditionalRecipients() ICSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("additionalRecipients"))
	return rv
}/* debug [instance_properties/getter]: additionalRecipients */


// An array of objects representing the content of the Cc: field in an email message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621664-additionalrecipients
func (c_ CSSearchableItemAttributeSet) SetAdditionalRecipients(value ICSPerson) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAdditionalRecipients:"), value)
}/* debug [instance_properties/setter]: additionalRecipients */


// An array of objects representing the content of the To: field in an email message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621675-primaryrecipients
func (c_ CSSearchableItemAttributeSet) PrimaryRecipients() ICSPerson {
	rv := objc.Send[CSPerson](c_.ID, objc.Sel("primaryRecipients"))
	return rv
}/* debug [instance_properties/getter]: primaryRecipients */


// An array of objects representing the content of the To: field in an email message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/1621675-primaryrecipients
func (c_ CSSearchableItemAttributeSet) SetPrimaryRecipients(value ICSPerson) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryRecipients:"), value)
}/* debug [instance_properties/setter]: primaryRecipients */


// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountHandles
func (c_ CSSearchableItemAttributeSet) AccountHandles() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("accountHandles"))
	return rv
}/* debug [instance_properties/getter]: accountHandles */


// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountHandles
func (c_ CSSearchableItemAttributeSet) SetAccountHandles(value []string) {
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
}/* debug [instance_properties/setter]: accountHandles */


// The unique identifier for the account with which the message is associated, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountIdentifier
func (c_ CSSearchableItemAttributeSet) AccountIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("accountIdentifier"))
	return rv
}/* debug [instance_properties/getter]: accountIdentifier */


// The unique identifier for the account with which the message is associated, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/accountIdentifier
func (c_ CSSearchableItemAttributeSet) SetAccountIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccountIdentifier:"), value)
}/* debug [instance_properties/setter]: accountIdentifier */


// The manufacturer of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionMake
func (c_ CSSearchableItemAttributeSet) AcquisitionMake() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("acquisitionMake"))
	return rv
}/* debug [instance_properties/getter]: acquisitionMake */


// The manufacturer of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionMake
func (c_ CSSearchableItemAttributeSet) SetAcquisitionMake(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionMake:"), value)
}/* debug [instance_properties/setter]: acquisitionMake */


// The model of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionModel
func (c_ CSSearchableItemAttributeSet) AcquisitionModel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("acquisitionModel"))
	return rv
}/* debug [instance_properties/getter]: acquisitionModel */


// The model of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/acquisitionModel
func (c_ CSSearchableItemAttributeSet) SetAcquisitionModel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionModel:"), value)
}/* debug [instance_properties/setter]: acquisitionModel */


// The date on which the item was moved into its current location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/addedDate
func (c_ CSSearchableItemAttributeSet) AddedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("addedDate"))
	return rv
}/* debug [instance_properties/getter]: addedDate */


// The date on which the item was moved into its current location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/addedDate
func (c_ CSSearchableItemAttributeSet) SetAddedDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAddedDate:"), value)
}/* debug [instance_properties/setter]: addedDate */


// The title for a collection of audio media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/album
func (c_ CSSearchableItemAttributeSet) Album() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("album"))
	return rv
}/* debug [instance_properties/getter]: album */


// The title for a collection of audio media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/album
func (c_ CSSearchableItemAttributeSet) SetAlbum(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlbum:"), value)
}/* debug [instance_properties/setter]: album */


// A value that indicates if the event covers an entire day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/allDay
func (c_ CSSearchableItemAttributeSet) AllDay() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("allDay"))
	return rv
}/* debug [instance_properties/getter]: allDay */


// A value that indicates if the event covers an entire day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/allDay
func (c_ CSSearchableItemAttributeSet) SetAllDay(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllDay:"), value)
}/* debug [instance_properties/setter]: allDay */


// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/alternateNames
func (c_ CSSearchableItemAttributeSet) AlternateNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("alternateNames"))
	return rv
}/* debug [instance_properties/getter]: alternateNames */


// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/alternateNames
func (c_ CSSearchableItemAttributeSet) SetAlternateNames(value []string) {
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
}/* debug [instance_properties/setter]: alternateNames */


// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/altitude
func (c_ CSSearchableItemAttributeSet) Altitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("altitude"))
	return rv
}/* debug [instance_properties/getter]: altitude */


// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/altitude
func (c_ CSSearchableItemAttributeSet) SetAltitude(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAltitude:"), value)
}/* debug [instance_properties/setter]: altitude */


// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/aperture
func (c_ CSSearchableItemAttributeSet) Aperture() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("aperture"))
	return rv
}/* debug [instance_properties/getter]: aperture */


// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/aperture
func (c_ CSSearchableItemAttributeSet) SetAperture(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAperture:"), value)
}/* debug [instance_properties/setter]: aperture */


// The artist associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/artist
func (c_ CSSearchableItemAttributeSet) Artist() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("artist"))
	return rv
}/* debug [instance_properties/getter]: artist */


// The artist associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/artist
func (c_ CSSearchableItemAttributeSet) SetArtist(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setArtist:"), value)
}/* debug [instance_properties/setter]: artist */


// A class of entity for which the item is intended or useful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audiences
func (c_ CSSearchableItemAttributeSet) Audiences() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("audiences"))
	return rv
}/* debug [instance_properties/getter]: audiences */


// A class of entity for which the item is intended or useful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audiences
func (c_ CSSearchableItemAttributeSet) SetAudiences(value []string) {
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
}/* debug [instance_properties/setter]: audiences */


// The audio bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioBitRate
func (c_ CSSearchableItemAttributeSet) AudioBitRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("audioBitRate"))
	return rv
}/* debug [instance_properties/getter]: audioBitRate */


// The audio bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioBitRate
func (c_ CSSearchableItemAttributeSet) SetAudioBitRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioBitRate:"), value)
}/* debug [instance_properties/setter]: audioBitRate */


// The number of channels in the audio data that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioChannelCount
func (c_ CSSearchableItemAttributeSet) AudioChannelCount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("audioChannelCount"))
	return rv
}/* debug [instance_properties/getter]: audioChannelCount */


// The number of channels in the audio data that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioChannelCount
func (c_ CSSearchableItemAttributeSet) SetAudioChannelCount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioChannelCount:"), value)
}/* debug [instance_properties/setter]: audioChannelCount */


// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioEncodingApplication
func (c_ CSSearchableItemAttributeSet) AudioEncodingApplication() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("audioEncodingApplication"))
	return rv
}/* debug [instance_properties/getter]: audioEncodingApplication */


// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioEncodingApplication
func (c_ CSSearchableItemAttributeSet) SetAudioEncodingApplication(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioEncodingApplication:"), value)
}/* debug [instance_properties/setter]: audioEncodingApplication */


// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioSampleRate
func (c_ CSSearchableItemAttributeSet) AudioSampleRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("audioSampleRate"))
	return rv
}/* debug [instance_properties/getter]: audioSampleRate */


// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioSampleRate
func (c_ CSSearchableItemAttributeSet) SetAudioSampleRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSampleRate:"), value)
}/* debug [instance_properties/setter]: audioSampleRate */


// The track number of a song or audio composition when part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioTrackNumber
func (c_ CSSearchableItemAttributeSet) AudioTrackNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("audioTrackNumber"))
	return rv
}/* debug [instance_properties/getter]: audioTrackNumber */


// The track number of a song or audio composition when part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/audioTrackNumber
func (c_ CSSearchableItemAttributeSet) SetAudioTrackNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioTrackNumber:"), value)
}/* debug [instance_properties/setter]: audioTrackNumber */


// An array of addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorAddresses
func (c_ CSSearchableItemAttributeSet) AuthorAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("authorAddresses"))
	return rv
}/* debug [instance_properties/getter]: authorAddresses */


// An array of addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorAddresses
func (c_ CSSearchableItemAttributeSet) SetAuthorAddresses(value []string) {
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
}/* debug [instance_properties/setter]: authorAddresses */


// An array of email addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorEmailAddresses
func (c_ CSSearchableItemAttributeSet) AuthorEmailAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("authorEmailAddresses"))
	return rv
}/* debug [instance_properties/getter]: authorEmailAddresses */


// An array of email addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorEmailAddresses
func (c_ CSSearchableItemAttributeSet) SetAuthorEmailAddresses(value []string) {
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
}/* debug [instance_properties/setter]: authorEmailAddresses */


// An array of names representing the authors who have worked on the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorNames
func (c_ CSSearchableItemAttributeSet) AuthorNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("authorNames"))
	return rv
}/* debug [instance_properties/getter]: authorNames */


// An array of names representing the authors who have worked on the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/authorNames
func (c_ CSSearchableItemAttributeSet) SetAuthorNames(value []string) {
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
}/* debug [instance_properties/setter]: authorNames */


// The number of bits per sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/bitsPerSample
func (c_ CSSearchableItemAttributeSet) BitsPerSample() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("bitsPerSample"))
	return rv
}/* debug [instance_properties/getter]: bitsPerSample */


// The number of bits per sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/bitsPerSample
func (c_ CSSearchableItemAttributeSet) SetBitsPerSample(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBitsPerSample:"), value)
}/* debug [instance_properties/setter]: bitsPerSample */


// The owner of the camera that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/cameraOwner
func (c_ CSSearchableItemAttributeSet) CameraOwner() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("cameraOwner"))
	return rv
}/* debug [instance_properties/getter]: cameraOwner */


// The owner of the camera that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/cameraOwner
func (c_ CSSearchableItemAttributeSet) SetCameraOwner(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraOwner:"), value)
}/* debug [instance_properties/setter]: cameraOwner */


// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/city
func (c_ CSSearchableItemAttributeSet) City() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("city"))
	return rv
}/* debug [instance_properties/getter]: city */


// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/city
func (c_ CSSearchableItemAttributeSet) SetCity(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), value)
}/* debug [instance_properties/setter]: city */


// The codecs used to encode/decode the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/codecs
func (c_ CSSearchableItemAttributeSet) Codecs() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("codecs"))
	return rv
}/* debug [instance_properties/getter]: codecs */


// The codecs used to encode/decode the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/codecs
func (c_ CSSearchableItemAttributeSet) SetCodecs(value []string) {
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
}/* debug [instance_properties/setter]: codecs */


// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/colorSpace
func (c_ CSSearchableItemAttributeSet) ColorSpace() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("colorSpace"))
	return rv
}/* debug [instance_properties/getter]: colorSpace */


// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/colorSpace
func (c_ CSSearchableItemAttributeSet) SetColorSpace(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSpace:"), value)
}/* debug [instance_properties/setter]: colorSpace */


// A comment related to the media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/comment
func (c_ CSSearchableItemAttributeSet) Comment() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("comment"))
	return rv
}/* debug [instance_properties/getter]: comment */


// A comment related to the media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/comment
func (c_ CSSearchableItemAttributeSet) SetComment(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComment:"), value)
}/* debug [instance_properties/setter]: comment */


// The date on which the item was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/completionDate
func (c_ CSSearchableItemAttributeSet) CompletionDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("completionDate"))
	return rv
}/* debug [instance_properties/getter]: completionDate */


// The date on which the item was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/completionDate
func (c_ CSSearchableItemAttributeSet) SetCompletionDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionDate:"), value)
}/* debug [instance_properties/setter]: completionDate */


// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contactKeywords
func (c_ CSSearchableItemAttributeSet) ContactKeywords() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contactKeywords"))
	return rv
}/* debug [instance_properties/getter]: contactKeywords */


// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contactKeywords
func (c_ CSSearchableItemAttributeSet) SetContactKeywords(value []string) {
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
}/* debug [instance_properties/setter]: contactKeywords */


// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerDisplayName
func (c_ CSSearchableItemAttributeSet) ContainerDisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerDisplayName"))
	return rv
}/* debug [instance_properties/getter]: containerDisplayName */


// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerDisplayName
func (c_ CSSearchableItemAttributeSet) SetContainerDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerDisplayName:"), value)
}/* debug [instance_properties/setter]: containerDisplayName */


// The identifier of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerIdentifier
func (c_ CSSearchableItemAttributeSet) ContainerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}/* debug [instance_properties/getter]: containerIdentifier */


// The identifier of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerIdentifier
func (c_ CSSearchableItemAttributeSet) SetContainerIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), value)
}/* debug [instance_properties/setter]: containerIdentifier */


// The order of the item within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerOrder
func (c_ CSSearchableItemAttributeSet) ContainerOrder() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("containerOrder"))
	return rv
}/* debug [instance_properties/getter]: containerOrder */


// The order of the item within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerOrder
func (c_ CSSearchableItemAttributeSet) SetContainerOrder(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerOrder:"), value)
}/* debug [instance_properties/setter]: containerOrder */


// The title of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerTitle
func (c_ CSSearchableItemAttributeSet) ContainerTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerTitle"))
	return rv
}/* debug [instance_properties/getter]: containerTitle */


// The title of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/containerTitle
func (c_ CSSearchableItemAttributeSet) SetContainerTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerTitle:"), value)
}/* debug [instance_properties/setter]: containerTitle */


// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentCreationDate
func (c_ CSSearchableItemAttributeSet) ContentCreationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("contentCreationDate"))
	return rv
}/* debug [instance_properties/getter]: contentCreationDate */


// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentCreationDate
func (c_ CSSearchableItemAttributeSet) SetContentCreationDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentCreationDate:"), value)
}/* debug [instance_properties/setter]: contentCreationDate */


// A description of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentDescription
func (c_ CSSearchableItemAttributeSet) ContentDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contentDescription"))
	return rv
}/* debug [instance_properties/getter]: contentDescription */


// A description of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentDescription
func (c_ CSSearchableItemAttributeSet) SetContentDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentDescription:"), value)
}/* debug [instance_properties/setter]: contentDescription */


// The date on which the contents of the file was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentModificationDate
func (c_ CSSearchableItemAttributeSet) ContentModificationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("contentModificationDate"))
	return rv
}/* debug [instance_properties/getter]: contentModificationDate */


// The date on which the contents of the file was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentModificationDate
func (c_ CSSearchableItemAttributeSet) SetContentModificationDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentModificationDate:"), value)
}/* debug [instance_properties/setter]: contentModificationDate */


// A value that indicates if the media contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentRating
func (c_ CSSearchableItemAttributeSet) ContentRating() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("contentRating"))
	return rv
}/* debug [instance_properties/getter]: contentRating */


// A value that indicates if the media contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentRating
func (c_ CSSearchableItemAttributeSet) SetContentRating(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentRating:"), value)
}/* debug [instance_properties/setter]: contentRating */


// An array of sources from which the media was obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentSources
func (c_ CSSearchableItemAttributeSet) ContentSources() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contentSources"))
	return rv
}/* debug [instance_properties/getter]: contentSources */


// An array of sources from which the media was obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentSources
func (c_ CSSearchableItemAttributeSet) SetContentSources(value []string) {
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
}/* debug [instance_properties/setter]: contentSources */


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentType
func (c_ CSSearchableItemAttributeSet) ContentType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contentType"))
	return rv
}/* debug [instance_properties/getter]: contentType */


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentType
func (c_ CSSearchableItemAttributeSet) SetContentType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentType:"), value)
}/* debug [instance_properties/setter]: contentType */


// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentTypeTree
func (c_ CSSearchableItemAttributeSet) ContentTypeTree() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contentTypeTree"))
	return rv
}/* debug [instance_properties/getter]: contentTypeTree */


// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentTypeTree
func (c_ CSSearchableItemAttributeSet) SetContentTypeTree(value []string) {
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
}/* debug [instance_properties/setter]: contentTypeTree */


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentURL
func (c_ CSSearchableItemAttributeSet) ContentURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("contentURL"))
	return rv
}/* debug [instance_properties/getter]: contentURL */


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentURL
func (c_ CSSearchableItemAttributeSet) SetContentURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentURL:"), value)
}/* debug [instance_properties/setter]: contentURL */


// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contributors
func (c_ CSSearchableItemAttributeSet) Contributors() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("contributors"))
	return rv
}/* debug [instance_properties/getter]: contributors */


// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contributors
func (c_ CSSearchableItemAttributeSet) SetContributors(value []string) {
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
}/* debug [instance_properties/setter]: contributors */


// The copyright date of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/copyright
func (c_ CSSearchableItemAttributeSet) Copyright() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("copyright"))
	return rv
}/* debug [instance_properties/getter]: copyright */


// The copyright date of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/copyright
func (c_ CSSearchableItemAttributeSet) SetCopyright(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCopyright:"), value)
}/* debug [instance_properties/setter]: copyright */


// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/country
func (c_ CSSearchableItemAttributeSet) Country() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("country"))
	return rv
}/* debug [instance_properties/getter]: country */


// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/country
func (c_ CSSearchableItemAttributeSet) SetCountry(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), value)
}/* debug [instance_properties/setter]: country */


// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/coverage
func (c_ CSSearchableItemAttributeSet) Coverage() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("coverage"))
	return rv
}/* debug [instance_properties/getter]: coverage */


// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/coverage
func (c_ CSSearchableItemAttributeSet) SetCoverage(value []string) {
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
}/* debug [instance_properties/setter]: coverage */


// The name of the app that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/creator
func (c_ CSSearchableItemAttributeSet) Creator() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("creator"))
	return rv
}/* debug [instance_properties/getter]: creator */


// The name of the app that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/creator
func (c_ CSSearchableItemAttributeSet) SetCreator(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreator:"), value)
}/* debug [instance_properties/setter]: creator */


// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/darkThumbnailURL
func (c_ CSSearchableItemAttributeSet) DarkThumbnailURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("darkThumbnailURL"))
	return rv
}/* debug [instance_properties/getter]: darkThumbnailURL */


// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/darkThumbnailURL
func (c_ CSSearchableItemAttributeSet) SetDarkThumbnailURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDarkThumbnailURL:"), value)
}/* debug [instance_properties/setter]: darkThumbnailURL */


// The delivery type of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/deliveryType
func (c_ CSSearchableItemAttributeSet) DeliveryType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("deliveryType"))
	return rv
}/* debug [instance_properties/getter]: deliveryType */


// The delivery type of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/deliveryType
func (c_ CSSearchableItemAttributeSet) SetDeliveryType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeliveryType:"), value)
}/* debug [instance_properties/setter]: deliveryType */


// The name of the director of the media (for example, a movie director).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/director
func (c_ CSSearchableItemAttributeSet) Director() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("director"))
	return rv
}/* debug [instance_properties/getter]: director */


// The name of the director of the media (for example, a movie director).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/director
func (c_ CSSearchableItemAttributeSet) SetDirector(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDirector:"), value)
}/* debug [instance_properties/setter]: director */


// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName
func (c_ CSSearchableItemAttributeSet) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName
func (c_ CSSearchableItemAttributeSet) SetDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDisplayName:"), value)
}/* debug [instance_properties/setter]: displayName */


// An identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/domainIdentifier
func (c_ CSSearchableItemAttributeSet) DomainIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("domainIdentifier"))
	return rv
}/* debug [instance_properties/getter]: domainIdentifier */


// An identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/domainIdentifier
func (c_ CSSearchableItemAttributeSet) SetDomainIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDomainIdentifier:"), value)
}/* debug [instance_properties/setter]: domainIdentifier */


// The most recent date on which the file was downloaded or received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/downloadedDate
func (c_ CSSearchableItemAttributeSet) DownloadedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("downloadedDate"))
	return rv
}/* debug [instance_properties/getter]: downloadedDate */


// The most recent date on which the file was downloaded or received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/downloadedDate
func (c_ CSSearchableItemAttributeSet) SetDownloadedDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDownloadedDate:"), value)
}/* debug [instance_properties/setter]: downloadedDate */


// The date on which the item is due.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/dueDate
func (c_ CSSearchableItemAttributeSet) DueDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("dueDate"))
	return rv
}/* debug [instance_properties/getter]: dueDate */


// The date on which the item is due.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/dueDate
func (c_ CSSearchableItemAttributeSet) SetDueDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDueDate:"), value)
}/* debug [instance_properties/setter]: dueDate */


// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/duration
func (c_ CSSearchableItemAttributeSet) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/duration
func (c_ CSSearchableItemAttributeSet) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// A list of editors who have worked on the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/editors
func (c_ CSSearchableItemAttributeSet) Editors() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("editors"))
	return rv
}/* debug [instance_properties/getter]: editors */


// A list of editors who have worked on the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/editors
func (c_ CSSearchableItemAttributeSet) SetEditors(value []string) {
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
}/* debug [instance_properties/setter]: editors */


// An array of email addresses associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailAddresses
func (c_ CSSearchableItemAttributeSet) EmailAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("emailAddresses"))
	return rv
}/* debug [instance_properties/getter]: emailAddresses */


// An array of email addresses associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailAddresses
func (c_ CSSearchableItemAttributeSet) SetEmailAddresses(value []string) {
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
}/* debug [instance_properties/setter]: emailAddresses */


// A dictionary that contains all the headers of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailHeaders
func (c_ CSSearchableItemAttributeSet) EmailHeaders() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("emailHeaders"))
	return rv
}/* debug [instance_properties/getter]: emailHeaders */


// A dictionary that contains all the headers of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/emailHeaders
func (c_ CSSearchableItemAttributeSet) SetEmailHeaders(value foundation.IDictionary) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailHeaders:"), value)
}/* debug [instance_properties/setter]: emailHeaders */


// The name of the apps that converted the original content into a PDF stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/encodingApplications
func (c_ CSSearchableItemAttributeSet) EncodingApplications() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("encodingApplications"))
	return rv
}/* debug [instance_properties/getter]: encodingApplications */


// The name of the apps that converted the original content into a PDF stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/encodingApplications
func (c_ CSSearchableItemAttributeSet) SetEncodingApplications(value []string) {
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
}/* debug [instance_properties/setter]: encodingApplications */


// The end date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/endDate
func (c_ CSSearchableItemAttributeSet) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("endDate"))
	return rv
}/* debug [instance_properties/getter]: endDate */


// The end date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/endDate
func (c_ CSSearchableItemAttributeSet) SetEndDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEndDate:"), value)
}/* debug [instance_properties/setter]: endDate */


// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifgpsVersion
func (c_ CSSearchableItemAttributeSet) EXIFGPSVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("EXIFGPSVersion"))
	return rv
}/* debug [instance_properties/getter]: EXIFGPSVersion */


// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifgpsVersion
func (c_ CSSearchableItemAttributeSet) SetEXIFGPSVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEXIFGPSVersion:"), value)
}/* debug [instance_properties/setter]: EXIFGPSVersion */


// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifVersion
func (c_ CSSearchableItemAttributeSet) EXIFVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("EXIFVersion"))
	return rv
}/* debug [instance_properties/getter]: EXIFVersion */


// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exifVersion
func (c_ CSSearchableItemAttributeSet) SetEXIFVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEXIFVersion:"), value)
}/* debug [instance_properties/setter]: EXIFVersion */


// The mode the camera used for the exposure of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureMode
func (c_ CSSearchableItemAttributeSet) ExposureMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("exposureMode"))
	return rv
}/* debug [instance_properties/getter]: exposureMode */


// The mode the camera used for the exposure of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureMode
func (c_ CSSearchableItemAttributeSet) SetExposureMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureMode:"), value)
}/* debug [instance_properties/setter]: exposureMode */


// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureProgram
func (c_ CSSearchableItemAttributeSet) ExposureProgram() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("exposureProgram"))
	return rv
}/* debug [instance_properties/getter]: exposureProgram */


// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureProgram
func (c_ CSSearchableItemAttributeSet) SetExposureProgram(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureProgram:"), value)
}/* debug [instance_properties/setter]: exposureProgram */


// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTime
func (c_ CSSearchableItemAttributeSet) ExposureTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("exposureTime"))
	return rv
}/* debug [instance_properties/getter]: exposureTime */


// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTime
func (c_ CSSearchableItemAttributeSet) SetExposureTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTime:"), value)
}/* debug [instance_properties/setter]: exposureTime */


// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTimeString
func (c_ CSSearchableItemAttributeSet) ExposureTimeString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("exposureTimeString"))
	return rv
}/* debug [instance_properties/getter]: exposureTimeString */


// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/exposureTimeString
func (c_ CSSearchableItemAttributeSet) SetExposureTimeString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTimeString:"), value)
}/* debug [instance_properties/setter]: exposureTimeString */


// The size of the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fileSize
func (c_ CSSearchableItemAttributeSet) FileSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("fileSize"))
	return rv
}/* debug [instance_properties/getter]: fileSize */


// The size of the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fileSize
func (c_ CSSearchableItemAttributeSet) SetFileSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFileSize:"), value)
}/* debug [instance_properties/setter]: fileSize */


// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/flashOn
func (c_ CSSearchableItemAttributeSet) FlashOn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("flashOn"))
	return rv
}/* debug [instance_properties/getter]: flashOn */


// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/flashOn
func (c_ CSSearchableItemAttributeSet) SetFlashOn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashOn:"), value)
}/* debug [instance_properties/setter]: flashOn */


// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fNumber
func (c_ CSSearchableItemAttributeSet) FNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("fNumber"))
	return rv
}/* debug [instance_properties/getter]: fNumber */


// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fNumber
func (c_ CSSearchableItemAttributeSet) SetFNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFNumber:"), value)
}/* debug [instance_properties/setter]: fNumber */


// The actual focal length of the lens, in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength
func (c_ CSSearchableItemAttributeSet) FocalLength() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("focalLength"))
	return rv
}/* debug [instance_properties/getter]: focalLength */


// The actual focal length of the lens, in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength
func (c_ CSSearchableItemAttributeSet) SetFocalLength(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength:"), value)
}/* debug [instance_properties/setter]: focalLength */


// A value that indicates if the focal length is 35mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength35mm
func (c_ CSSearchableItemAttributeSet) FocalLength35mm() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("focalLength35mm"))
	return rv
}/* debug [instance_properties/getter]: focalLength35mm */


// A value that indicates if the focal length is 35mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/focalLength35mm
func (c_ CSSearchableItemAttributeSet) SetFocalLength35mm(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength35mm:"), value)
}/* debug [instance_properties/setter]: focalLength35mm */


// An array of font names the document uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fontNames
func (c_ CSSearchableItemAttributeSet) FontNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("fontNames"))
	return rv
}/* debug [instance_properties/getter]: fontNames */


// An array of font names the document uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fontNames
func (c_ CSSearchableItemAttributeSet) SetFontNames(value []string) {
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
}/* debug [instance_properties/setter]: fontNames */


// The fully formatted address of the item, received from MapKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fullyFormattedAddress
func (c_ CSSearchableItemAttributeSet) FullyFormattedAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("fullyFormattedAddress"))
	return rv
}/* debug [instance_properties/getter]: fullyFormattedAddress */


// The fully formatted address of the item, received from MapKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/fullyFormattedAddress
func (c_ CSSearchableItemAttributeSet) SetFullyFormattedAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFullyFormattedAddress:"), value)
}/* debug [instance_properties/setter]: fullyFormattedAddress */


// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/generalMIDISequence
func (c_ CSSearchableItemAttributeSet) GeneralMIDISequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("generalMIDISequence"))
	return rv
}/* debug [instance_properties/getter]: generalMIDISequence */


// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/generalMIDISequence
func (c_ CSSearchableItemAttributeSet) SetGeneralMIDISequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGeneralMIDISequence:"), value)
}/* debug [instance_properties/setter]: generalMIDISequence */


// The genre of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/genre
func (c_ CSSearchableItemAttributeSet) Genre() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("genre"))
	return rv
}/* debug [instance_properties/getter]: genre */


// The genre of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/genre
func (c_ CSSearchableItemAttributeSet) SetGenre(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGenre:"), value)
}/* debug [instance_properties/setter]: genre */


// Information about the GPS area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsAreaInformation
func (c_ CSSearchableItemAttributeSet) GPSAreaInformation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("GPSAreaInformation"))
	return rv
}/* debug [instance_properties/getter]: GPSAreaInformation */


// Information about the GPS area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsAreaInformation
func (c_ CSSearchableItemAttributeSet) SetGPSAreaInformation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSAreaInformation:"), value)
}/* debug [instance_properties/setter]: GPSAreaInformation */


// The date and time related to the GPS value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDateStamp
func (c_ CSSearchableItemAttributeSet) GPSDateStamp() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("GPSDateStamp"))
	return rv
}/* debug [instance_properties/getter]: GPSDateStamp */


// The date and time related to the GPS value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDateStamp
func (c_ CSSearchableItemAttributeSet) SetGPSDateStamp(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDateStamp:"), value)
}/* debug [instance_properties/setter]: GPSDateStamp */


// The bearing to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestBearing
func (c_ CSSearchableItemAttributeSet) GPSDestBearing() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("GPSDestBearing"))
	return rv
}/* debug [instance_properties/getter]: GPSDestBearing */


// The bearing to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestBearing
func (c_ CSSearchableItemAttributeSet) SetGPSDestBearing(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestBearing:"), value)
}/* debug [instance_properties/setter]: GPSDestBearing */


// The distance to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestDistance
func (c_ CSSearchableItemAttributeSet) GPSDestDistance() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("GPSDestDistance"))
	return rv
}/* debug [instance_properties/getter]: GPSDestDistance */


// The distance to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestDistance
func (c_ CSSearchableItemAttributeSet) SetGPSDestDistance(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestDistance:"), value)
}/* debug [instance_properties/setter]: GPSDestDistance */


// The latitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLatitude
func (c_ CSSearchableItemAttributeSet) GPSDestLatitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("GPSDestLatitude"))
	return rv
}/* debug [instance_properties/getter]: GPSDestLatitude */


// The latitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLatitude
func (c_ CSSearchableItemAttributeSet) SetGPSDestLatitude(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestLatitude:"), value)
}/* debug [instance_properties/setter]: GPSDestLatitude */


// The longitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLongitude
func (c_ CSSearchableItemAttributeSet) GPSDestLongitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("GPSDestLongitude"))
	return rv
}/* debug [instance_properties/getter]: GPSDestLongitude */


// The longitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDestLongitude
func (c_ CSSearchableItemAttributeSet) SetGPSDestLongitude(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDestLongitude:"), value)
}/* debug [instance_properties/setter]: GPSDestLongitude */


// The differential correction applied to the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDifferental
func (c_ CSSearchableItemAttributeSet) GPSDifferental() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("GPSDifferental"))
	return rv
}/* debug [instance_properties/getter]: GPSDifferental */


// The differential correction applied to the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsDifferental
func (c_ CSSearchableItemAttributeSet) SetGPSDifferental(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDifferental:"), value)
}/* debug [instance_properties/setter]: GPSDifferental */


// The GPS dilution of precision value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop
func (c_ CSSearchableItemAttributeSet) GPSDOP() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("GPSDOP"))
	return rv
}/* debug [instance_properties/getter]: GPSDOP */


// The GPS dilution of precision value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop
func (c_ CSSearchableItemAttributeSet) SetGPSDOP(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDOP:"), value)
}/* debug [instance_properties/setter]: GPSDOP */


// The geodetic data that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMapDatum
func (c_ CSSearchableItemAttributeSet) GPSMapDatum() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("GPSMapDatum"))
	return rv
}/* debug [instance_properties/getter]: GPSMapDatum */


// The geodetic data that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMapDatum
func (c_ CSSearchableItemAttributeSet) SetGPSMapDatum(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSMapDatum:"), value)
}/* debug [instance_properties/setter]: GPSMapDatum */


// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMeasureMode
func (c_ CSSearchableItemAttributeSet) GPSMeasureMode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("GPSMeasureMode"))
	return rv
}/* debug [instance_properties/getter]: GPSMeasureMode */


// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsMeasureMode
func (c_ CSSearchableItemAttributeSet) SetGPSMeasureMode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSMeasureMode:"), value)
}/* debug [instance_properties/setter]: GPSMeasureMode */


// The location finding method that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsProcessingMethod
func (c_ CSSearchableItemAttributeSet) GPSProcessingMethod() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("GPSProcessingMethod"))
	return rv
}/* debug [instance_properties/getter]: GPSProcessingMethod */


// The location finding method that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsProcessingMethod
func (c_ CSSearchableItemAttributeSet) SetGPSProcessingMethod(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSProcessingMethod:"), value)
}/* debug [instance_properties/setter]: GPSProcessingMethod */


// The status of the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsStatus
func (c_ CSSearchableItemAttributeSet) GPSStatus() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("GPSStatus"))
	return rv
}/* debug [instance_properties/getter]: GPSStatus */


// The status of the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsStatus
func (c_ CSSearchableItemAttributeSet) SetGPSStatus(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSStatus:"), value)
}/* debug [instance_properties/setter]: GPSStatus */


// The direction of travel of the item in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsTrack
func (c_ CSSearchableItemAttributeSet) GPSTrack() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("GPSTrack"))
	return rv
}/* debug [instance_properties/getter]: GPSTrack */


// The direction of travel of the item in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsTrack
func (c_ CSSearchableItemAttributeSet) SetGPSTrack(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSTrack:"), value)
}/* debug [instance_properties/setter]: GPSTrack */


// Indicates if the image file has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hasAlphaChannel
func (c_ CSSearchableItemAttributeSet) HasAlphaChannel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("hasAlphaChannel"))
	return rv
}/* debug [instance_properties/getter]: hasAlphaChannel */


// Indicates if the image file has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/hasAlphaChannel
func (c_ CSSearchableItemAttributeSet) SetHasAlphaChannel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasAlphaChannel:"), value)
}/* debug [instance_properties/setter]: hasAlphaChannel */


// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/headline
func (c_ CSSearchableItemAttributeSet) Headline() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("headline"))
	return rv
}/* debug [instance_properties/getter]: headline */


// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/headline
func (c_ CSSearchableItemAttributeSet) SetHeadline(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHeadline:"), value)
}/* debug [instance_properties/setter]: headline */


// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/htmlContentData
func (c_ CSSearchableItemAttributeSet) HTMLContentData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("HTMLContentData"))
	return rv
}/* debug [instance_properties/getter]: HTMLContentData */


// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/htmlContentData
func (c_ CSSearchableItemAttributeSet) SetHTMLContentData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHTMLContentData:"), value)
}/* debug [instance_properties/setter]: HTMLContentData */


// A formal identifier that references the document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/identifier
func (c_ CSSearchableItemAttributeSet) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// A formal identifier that references the document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/identifier
func (c_ CSSearchableItemAttributeSet) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), value)
}/* debug [instance_properties/setter]: identifier */


// The direction of the item’s image in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/imageDirection
func (c_ CSSearchableItemAttributeSet) ImageDirection() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("imageDirection"))
	return rv
}/* debug [instance_properties/getter]: imageDirection */


// The direction of the item’s image in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/imageDirection
func (c_ CSSearchableItemAttributeSet) SetImageDirection(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageDirection:"), value)
}/* debug [instance_properties/setter]: imageDirection */


// An array of important dates associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/importantDates
func (c_ CSSearchableItemAttributeSet) ImportantDates() []foundation.Date {
	rv := objc.Send[[]foundation.Date](c_.ID, objc.Sel("importantDates"))
	return rv
}/* debug [instance_properties/getter]: importantDates */


// An array of important dates associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/importantDates
func (c_ CSSearchableItemAttributeSet) SetImportantDates(value []foundation.Date) {
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
}/* debug [instance_properties/setter]: importantDates */


// Information about the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/information
func (c_ CSSearchableItemAttributeSet) Information() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("information"))
	return rv
}/* debug [instance_properties/getter]: information */


// Information about the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/information
func (c_ CSSearchableItemAttributeSet) SetInformation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInformation:"), value)
}/* debug [instance_properties/setter]: information */


// An array of instant message addresses for the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instantMessageAddresses
func (c_ CSSearchableItemAttributeSet) InstantMessageAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}/* debug [instance_properties/getter]: instantMessageAddresses */


// An array of instant message addresses for the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instantMessageAddresses
func (c_ CSSearchableItemAttributeSet) SetInstantMessageAddresses(value []string) {
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
}/* debug [instance_properties/setter]: instantMessageAddresses */


// Instructions that concern the use of the item, such as an embargo or warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instructions
func (c_ CSSearchableItemAttributeSet) Instructions() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("instructions"))
	return rv
}/* debug [instance_properties/getter]: instructions */


// Instructions that concern the use of the item, such as an embargo or warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/instructions
func (c_ CSSearchableItemAttributeSet) SetInstructions(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstructions:"), value)
}/* debug [instance_properties/setter]: instructions */


// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isoSpeed
func (c_ CSSearchableItemAttributeSet) ISOSpeed() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("ISOSpeed"))
	return rv
}/* debug [instance_properties/getter]: ISOSpeed */


// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isoSpeed
func (c_ CSSearchableItemAttributeSet) SetISOSpeed(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setISOSpeed:"), value)
}/* debug [instance_properties/setter]: ISOSpeed */


// A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/isPriority
func (c_ CSSearchableItemAttributeSet) IsPriority() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("isPriority"))
	return rv
}/* debug [instance_properties/getter]: isPriority */


// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keySignature
func (c_ CSSearchableItemAttributeSet) KeySignature() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("keySignature"))
	return rv
}/* debug [instance_properties/getter]: keySignature */


// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keySignature
func (c_ CSSearchableItemAttributeSet) SetKeySignature(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeySignature:"), value)
}/* debug [instance_properties/setter]: keySignature */


// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keywords
func (c_ CSSearchableItemAttributeSet) Keywords() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("keywords"))
	return rv
}/* debug [instance_properties/getter]: keywords */


// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/keywords
func (c_ CSSearchableItemAttributeSet) SetKeywords(value []string) {
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
}/* debug [instance_properties/setter]: keywords */


// A description of the kind of document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/kind
func (c_ CSSearchableItemAttributeSet) Kind() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("kind"))
	return rv
}/* debug [instance_properties/getter]: kind */


// A description of the kind of document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/kind
func (c_ CSSearchableItemAttributeSet) SetKind(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKind:"), value)
}/* debug [instance_properties/setter]: kind */


// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/languages
func (c_ CSSearchableItemAttributeSet) Languages() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("languages"))
	return rv
}/* debug [instance_properties/getter]: languages */


// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/languages
func (c_ CSSearchableItemAttributeSet) SetLanguages(value []string) {
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
}/* debug [instance_properties/setter]: languages */


// The date on which the file was last used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lastUsedDate
func (c_ CSSearchableItemAttributeSet) LastUsedDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("lastUsedDate"))
	return rv
}/* debug [instance_properties/getter]: lastUsedDate */


// The date on which the file was last used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lastUsedDate
func (c_ CSSearchableItemAttributeSet) SetLastUsedDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLastUsedDate:"), value)
}/* debug [instance_properties/setter]: lastUsedDate */


// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/latitude
func (c_ CSSearchableItemAttributeSet) Latitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("latitude"))
	return rv
}/* debug [instance_properties/getter]: latitude */


// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/latitude
func (c_ CSSearchableItemAttributeSet) SetLatitude(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLatitude:"), value)
}/* debug [instance_properties/setter]: latitude */


// An array that contains the names of the various layers in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/layerNames
func (c_ CSSearchableItemAttributeSet) LayerNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("layerNames"))
	return rv
}/* debug [instance_properties/getter]: layerNames */


// An array that contains the names of the various layers in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/layerNames
func (c_ CSSearchableItemAttributeSet) SetLayerNames(value []string) {
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
}/* debug [instance_properties/setter]: layerNames */


// The model of the lens that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lensModel
func (c_ CSSearchableItemAttributeSet) LensModel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("lensModel"))
	return rv
}/* debug [instance_properties/getter]: lensModel */


// The model of the lens that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lensModel
func (c_ CSSearchableItemAttributeSet) SetLensModel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLensModel:"), value)
}/* debug [instance_properties/setter]: lensModel */


// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/likelyJunk
func (c_ CSSearchableItemAttributeSet) LikelyJunk() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("likelyJunk"))
	return rv
}/* debug [instance_properties/getter]: likelyJunk */


// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/likelyJunk
func (c_ CSSearchableItemAttributeSet) SetLikelyJunk(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLikelyJunk:"), value)
}/* debug [instance_properties/setter]: likelyJunk */


// A value that indicates if the media is local.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/local
func (c_ CSSearchableItemAttributeSet) Local() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("local"))
	return rv
}/* debug [instance_properties/getter]: local */


// A value that indicates if the media is local.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/local
func (c_ CSSearchableItemAttributeSet) SetLocal(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocal:"), value)
}/* debug [instance_properties/setter]: local */


// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/longitude
func (c_ CSSearchableItemAttributeSet) Longitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("longitude"))
	return rv
}/* debug [instance_properties/getter]: longitude */


// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/longitude
func (c_ CSSearchableItemAttributeSet) SetLongitude(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongitude:"), value)
}/* debug [instance_properties/setter]: longitude */


// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lyricist
func (c_ CSSearchableItemAttributeSet) Lyricist() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("lyricist"))
	return rv
}/* debug [instance_properties/getter]: lyricist */


// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/lyricist
func (c_ CSSearchableItemAttributeSet) SetLyricist(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLyricist:"), value)
}/* debug [instance_properties/setter]: lyricist */


// An array of mailbox identifiers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mailboxIdentifiers
func (c_ CSSearchableItemAttributeSet) MailboxIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("mailboxIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: mailboxIdentifiers */


// An array of mailbox identifiers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mailboxIdentifiers
func (c_ CSSearchableItemAttributeSet) SetMailboxIdentifiers(value []string) {
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
}/* debug [instance_properties/setter]: mailboxIdentifiers */


// The smallest F number of the lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/maxAperture
func (c_ CSSearchableItemAttributeSet) MaxAperture() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("maxAperture"))
	return rv
}/* debug [instance_properties/getter]: maxAperture */


// The smallest F number of the lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/maxAperture
func (c_ CSSearchableItemAttributeSet) SetMaxAperture(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxAperture:"), value)
}/* debug [instance_properties/setter]: maxAperture */


// The media types present in the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mediaTypes
func (c_ CSSearchableItemAttributeSet) MediaTypes() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("mediaTypes"))
	return rv
}/* debug [instance_properties/getter]: mediaTypes */


// The media types present in the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/mediaTypes
func (c_ CSSearchableItemAttributeSet) SetMediaTypes(value []string) {
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
}/* debug [instance_properties/setter]: mediaTypes */


// The date on which the last metadata attribute was changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/metadataModificationDate
func (c_ CSSearchableItemAttributeSet) MetadataModificationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("metadataModificationDate"))
	return rv
}/* debug [instance_properties/getter]: metadataModificationDate */


// The date on which the last metadata attribute was changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/metadataModificationDate
func (c_ CSSearchableItemAttributeSet) SetMetadataModificationDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataModificationDate:"), value)
}/* debug [instance_properties/setter]: metadataModificationDate */


// The metering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/meteringMode
func (c_ CSSearchableItemAttributeSet) MeteringMode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("meteringMode"))
	return rv
}/* debug [instance_properties/getter]: meteringMode */


// The metering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/meteringMode
func (c_ CSSearchableItemAttributeSet) SetMeteringMode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMeteringMode:"), value)
}/* debug [instance_properties/setter]: meteringMode */


// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalGenre
func (c_ CSSearchableItemAttributeSet) MusicalGenre() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("musicalGenre"))
	return rv
}/* debug [instance_properties/getter]: musicalGenre */


// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalGenre
func (c_ CSSearchableItemAttributeSet) SetMusicalGenre(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalGenre:"), value)
}/* debug [instance_properties/setter]: musicalGenre */


// The category of the instrument associated with the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentCategory
func (c_ CSSearchableItemAttributeSet) MusicalInstrumentCategory() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("musicalInstrumentCategory"))
	return rv
}/* debug [instance_properties/getter]: musicalInstrumentCategory */


// The category of the instrument associated with the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentCategory
func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentCategory:"), value)
}/* debug [instance_properties/setter]: musicalInstrumentCategory */


// The name of an instrument within the context of an instrument category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentName
func (c_ CSSearchableItemAttributeSet) MusicalInstrumentName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("musicalInstrumentName"))
	return rv
}/* debug [instance_properties/getter]: musicalInstrumentName */


// The name of an instrument within the context of an instrument category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/musicalInstrumentName
func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentName:"), value)
}/* debug [instance_properties/setter]: musicalInstrumentName */


// The name of the location or point of interest associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/namedLocation
func (c_ CSSearchableItemAttributeSet) NamedLocation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("namedLocation"))
	return rv
}/* debug [instance_properties/getter]: namedLocation */


// The name of the location or point of interest associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/namedLocation
func (c_ CSSearchableItemAttributeSet) SetNamedLocation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNamedLocation:"), value)
}/* debug [instance_properties/setter]: namedLocation */


// A list of companies or organizations that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/organizations
func (c_ CSSearchableItemAttributeSet) Organizations() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("organizations"))
	return rv
}/* debug [instance_properties/getter]: organizations */


// A list of companies or organizations that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/organizations
func (c_ CSSearchableItemAttributeSet) SetOrganizations(value []string) {
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
}/* debug [instance_properties/setter]: organizations */


// The orientation of the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/orientation
func (c_ CSSearchableItemAttributeSet) Orientation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("orientation"))
	return rv
}/* debug [instance_properties/getter]: orientation */


// The orientation of the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/orientation
func (c_ CSSearchableItemAttributeSet) SetOrientation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrientation:"), value)
}/* debug [instance_properties/setter]: orientation */


// The original format of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalFormat
func (c_ CSSearchableItemAttributeSet) OriginalFormat() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("originalFormat"))
	return rv
}/* debug [instance_properties/getter]: originalFormat */


// The original format of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalFormat
func (c_ CSSearchableItemAttributeSet) SetOriginalFormat(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalFormat:"), value)
}/* debug [instance_properties/setter]: originalFormat */


// The original source of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalSource
func (c_ CSSearchableItemAttributeSet) OriginalSource() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("originalSource"))
	return rv
}/* debug [instance_properties/getter]: originalSource */


// The original source of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/originalSource
func (c_ CSSearchableItemAttributeSet) SetOriginalSource(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalSource:"), value)
}/* debug [instance_properties/setter]: originalSource */


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageCount
func (c_ CSSearchableItemAttributeSet) PageCount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pageCount"))
	return rv
}/* debug [instance_properties/getter]: pageCount */


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageCount
func (c_ CSSearchableItemAttributeSet) SetPageCount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageCount:"), value)
}/* debug [instance_properties/setter]: pageCount */


// The height of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageHeight
func (c_ CSSearchableItemAttributeSet) PageHeight() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pageHeight"))
	return rv
}/* debug [instance_properties/getter]: pageHeight */


// The height of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageHeight
func (c_ CSSearchableItemAttributeSet) SetPageHeight(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageHeight:"), value)
}/* debug [instance_properties/setter]: pageHeight */


// The width of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageWidth
func (c_ CSSearchableItemAttributeSet) PageWidth() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pageWidth"))
	return rv
}/* debug [instance_properties/getter]: pageWidth */


// The width of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pageWidth
func (c_ CSSearchableItemAttributeSet) SetPageWidth(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageWidth:"), value)
}/* debug [instance_properties/setter]: pageWidth */


// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/participants
func (c_ CSSearchableItemAttributeSet) Participants() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("participants"))
	return rv
}/* debug [instance_properties/getter]: participants */


// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/participants
func (c_ CSSearchableItemAttributeSet) SetParticipants(value []string) {
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
}/* debug [instance_properties/setter]: participants */


// The complete path to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/path
func (c_ CSSearchableItemAttributeSet) Path() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("path"))
	return rv
}/* debug [instance_properties/getter]: path */


// The complete path to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/path
func (c_ CSSearchableItemAttributeSet) SetPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPath:"), value)
}/* debug [instance_properties/setter]: path */


// A list of performers in the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/performers
func (c_ CSSearchableItemAttributeSet) Performers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("performers"))
	return rv
}/* debug [instance_properties/getter]: performers */


// A list of performers in the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/performers
func (c_ CSSearchableItemAttributeSet) SetPerformers(value []string) {
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
}/* debug [instance_properties/setter]: performers */


// An array of phone numbers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/phoneNumbers
func (c_ CSSearchableItemAttributeSet) PhoneNumbers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}/* debug [instance_properties/getter]: phoneNumbers */


// An array of phone numbers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/phoneNumbers
func (c_ CSSearchableItemAttributeSet) SetPhoneNumbers(value []string) {
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
}/* debug [instance_properties/setter]: phoneNumbers */


// The total number of pixels in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelCount
func (c_ CSSearchableItemAttributeSet) PixelCount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pixelCount"))
	return rv
}/* debug [instance_properties/getter]: pixelCount */


// The total number of pixels in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelCount
func (c_ CSSearchableItemAttributeSet) SetPixelCount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelCount:"), value)
}/* debug [instance_properties/setter]: pixelCount */


// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelHeight
func (c_ CSSearchableItemAttributeSet) PixelHeight() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pixelHeight"))
	return rv
}/* debug [instance_properties/getter]: pixelHeight */


// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelHeight
func (c_ CSSearchableItemAttributeSet) SetPixelHeight(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelHeight:"), value)
}/* debug [instance_properties/setter]: pixelHeight */


// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelWidth
func (c_ CSSearchableItemAttributeSet) PixelWidth() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pixelWidth"))
	return rv
}/* debug [instance_properties/getter]: pixelWidth */


// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/pixelWidth
func (c_ CSSearchableItemAttributeSet) SetPixelWidth(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelWidth:"), value)
}/* debug [instance_properties/setter]: pixelWidth */


// A user-supplied play count for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/playCount
func (c_ CSSearchableItemAttributeSet) PlayCount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("playCount"))
	return rv
}/* debug [instance_properties/getter]: playCount */


// A user-supplied play count for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/playCount
func (c_ CSSearchableItemAttributeSet) SetPlayCount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayCount:"), value)
}/* debug [instance_properties/setter]: playCount */


// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/postalCode
func (c_ CSSearchableItemAttributeSet) PostalCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("postalCode"))
	return rv
}/* debug [instance_properties/getter]: postalCode */


// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/postalCode
func (c_ CSSearchableItemAttributeSet) SetPostalCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalCode:"), value)
}/* debug [instance_properties/setter]: postalCode */


// The producer of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/producer
func (c_ CSSearchableItemAttributeSet) Producer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("producer"))
	return rv
}/* debug [instance_properties/getter]: producer */


// The producer of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/producer
func (c_ CSSearchableItemAttributeSet) SetProducer(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProducer:"), value)
}/* debug [instance_properties/setter]: producer */


// The name of the color profile the camera used for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/profileName
func (c_ CSSearchableItemAttributeSet) ProfileName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("profileName"))
	return rv
}/* debug [instance_properties/getter]: profileName */


// The name of the color profile the camera used for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/profileName
func (c_ CSSearchableItemAttributeSet) SetProfileName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProfileName:"), value)
}/* debug [instance_properties/setter]: profileName */


// A list of projects of which this file is a part.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/projects
func (c_ CSSearchableItemAttributeSet) Projects() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("projects"))
	return rv
}/* debug [instance_properties/getter]: projects */


// A list of projects of which this file is a part.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/projects
func (c_ CSSearchableItemAttributeSet) SetProjects(value []string) {
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
}/* debug [instance_properties/setter]: projects */


// An array of identifiers that corresponds to data representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerDataTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) ProviderDataTypeIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("providerDataTypeIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: providerDataTypeIdentifiers */


// An array of identifiers that corresponds to data representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerDataTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderDataTypeIdentifiers(value []string) {
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
}/* debug [instance_properties/setter]: providerDataTypeIdentifiers */


// An array of identifiers that corresponds to file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerFileTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) ProviderFileTypeIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("providerFileTypeIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: providerFileTypeIdentifiers */


// An array of identifiers that corresponds to file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerFileTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderFileTypeIdentifiers(value []string) {
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
}/* debug [instance_properties/setter]: providerFileTypeIdentifiers */


// An array of identifiers that corresponds to in-place file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerInPlaceFileTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) ProviderInPlaceFileTypeIdentifiers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("providerInPlaceFileTypeIdentifiers"))
	return rv
}/* debug [instance_properties/getter]: providerInPlaceFileTypeIdentifiers */


// An array of identifiers that corresponds to in-place file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/providerInPlaceFileTypeIdentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderInPlaceFileTypeIdentifiers(value []string) {
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
}/* debug [instance_properties/setter]: providerInPlaceFileTypeIdentifiers */


// A list of people, organizations, services, or other entities responsible for making the media available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/publishers
func (c_ CSSearchableItemAttributeSet) Publishers() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("publishers"))
	return rv
}/* debug [instance_properties/getter]: publishers */


// A list of people, organizations, services, or other entities responsible for making the media available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/publishers
func (c_ CSSearchableItemAttributeSet) SetPublishers(value []string) {
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
}/* debug [instance_properties/setter]: publishers */


// A number that indicates the relative importance of the item among other items from the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rankingHint
func (c_ CSSearchableItemAttributeSet) RankingHint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("rankingHint"))
	return rv
}/* debug [instance_properties/getter]: rankingHint */


// A number that indicates the relative importance of the item among other items from the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rankingHint
func (c_ CSSearchableItemAttributeSet) SetRankingHint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRankingHint:"), value)
}/* debug [instance_properties/setter]: rankingHint */


// The user-supplied rating of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rating
func (c_ CSSearchableItemAttributeSet) Rating() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("rating"))
	return rv
}/* debug [instance_properties/getter]: rating */


// The user-supplied rating of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rating
func (c_ CSSearchableItemAttributeSet) SetRating(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRating:"), value)
}/* debug [instance_properties/setter]: rating */


// A description of the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/ratingDescription
func (c_ CSSearchableItemAttributeSet) RatingDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ratingDescription"))
	return rv
}/* debug [instance_properties/getter]: ratingDescription */


// A description of the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/ratingDescription
func (c_ CSSearchableItemAttributeSet) SetRatingDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRatingDescription:"), value)
}/* debug [instance_properties/setter]: ratingDescription */


// An array of addresses associated with the recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientAddresses
func (c_ CSSearchableItemAttributeSet) RecipientAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("recipientAddresses"))
	return rv
}/* debug [instance_properties/getter]: recipientAddresses */


// An array of addresses associated with the recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientAddresses
func (c_ CSSearchableItemAttributeSet) SetRecipientAddresses(value []string) {
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
}/* debug [instance_properties/setter]: recipientAddresses */


// An array of email addresses associated with the recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientEmailAddresses
func (c_ CSSearchableItemAttributeSet) RecipientEmailAddresses() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("recipientEmailAddresses"))
	return rv
}/* debug [instance_properties/getter]: recipientEmailAddresses */


// An array of email addresses associated with the recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientEmailAddresses
func (c_ CSSearchableItemAttributeSet) SetRecipientEmailAddresses(value []string) {
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
}/* debug [instance_properties/setter]: recipientEmailAddresses */


// An array of names representing the recipients of this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientNames
func (c_ CSSearchableItemAttributeSet) RecipientNames() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("recipientNames"))
	return rv
}/* debug [instance_properties/getter]: recipientNames */


// An array of names representing the recipients of this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recipientNames
func (c_ CSSearchableItemAttributeSet) SetRecipientNames(value []string) {
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
}/* debug [instance_properties/setter]: recipientNames */


// The recording date of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recordingDate
func (c_ CSSearchableItemAttributeSet) RecordingDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("recordingDate"))
	return rv
}/* debug [instance_properties/getter]: recordingDate */


// The recording date of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/recordingDate
func (c_ CSSearchableItemAttributeSet) SetRecordingDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordingDate:"), value)
}/* debug [instance_properties/setter]: recordingDate */


// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/redEyeOn
func (c_ CSSearchableItemAttributeSet) RedEyeOn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("redEyeOn"))
	return rv
}/* debug [instance_properties/getter]: redEyeOn */


// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/redEyeOn
func (c_ CSSearchableItemAttributeSet) SetRedEyeOn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRedEyeOn:"), value)
}/* debug [instance_properties/setter]: redEyeOn */


// The unique identifier for the item to which the activity is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/relatedUniqueIdentifier
func (c_ CSSearchableItemAttributeSet) RelatedUniqueIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("relatedUniqueIdentifier"))
	return rv
}/* debug [instance_properties/getter]: relatedUniqueIdentifier */


// The unique identifier for the item to which the activity is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/relatedUniqueIdentifier
func (c_ CSSearchableItemAttributeSet) SetRelatedUniqueIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRelatedUniqueIdentifier:"), value)
}/* debug [instance_properties/setter]: relatedUniqueIdentifier */


// The resolution height of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionHeightDPI
func (c_ CSSearchableItemAttributeSet) ResolutionHeightDPI() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("resolutionHeightDPI"))
	return rv
}/* debug [instance_properties/getter]: resolutionHeightDPI */


// The resolution height of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionHeightDPI
func (c_ CSSearchableItemAttributeSet) SetResolutionHeightDPI(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionHeightDPI:"), value)
}/* debug [instance_properties/setter]: resolutionHeightDPI */


// The resolution width of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionWidthDPI
func (c_ CSSearchableItemAttributeSet) ResolutionWidthDPI() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("resolutionWidthDPI"))
	return rv
}/* debug [instance_properties/getter]: resolutionWidthDPI */


// The resolution width of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/resolutionWidthDPI
func (c_ CSSearchableItemAttributeSet) SetResolutionWidthDPI(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionWidthDPI:"), value)
}/* debug [instance_properties/setter]: resolutionWidthDPI */


// A link to information about the rights held in and over the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rights
func (c_ CSSearchableItemAttributeSet) Rights() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("rights"))
	return rv
}/* debug [instance_properties/getter]: rights */


// A link to information about the rights held in and over the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/rights
func (c_ CSSearchableItemAttributeSet) SetRights(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRights:"), value)
}/* debug [instance_properties/setter]: rights */


// Indicates the role of the content creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/role
func (c_ CSSearchableItemAttributeSet) Role() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("role"))
	return rv
}/* debug [instance_properties/getter]: role */


// Indicates the role of the content creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/role
func (c_ CSSearchableItemAttributeSet) SetRole(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRole:"), value)
}/* debug [instance_properties/setter]: role */


// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/securityMethod
func (c_ CSSearchableItemAttributeSet) SecurityMethod() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("securityMethod"))
	return rv
}/* debug [instance_properties/getter]: securityMethod */


// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/securityMethod
func (c_ CSSearchableItemAttributeSet) SetSecurityMethod(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecurityMethod:"), value)
}/* debug [instance_properties/setter]: securityMethod */


// The speed of the item, in kilometers per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/speed
func (c_ CSSearchableItemAttributeSet) Speed() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("speed"))
	return rv
}/* debug [instance_properties/getter]: speed */


// The speed of the item, in kilometers per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/speed
func (c_ CSSearchableItemAttributeSet) SetSpeed(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpeed:"), value)
}/* debug [instance_properties/setter]: speed */


// The start date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/startDate
func (c_ CSSearchableItemAttributeSet) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("startDate"))
	return rv
}/* debug [instance_properties/getter]: startDate */


// The start date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/startDate
func (c_ CSSearchableItemAttributeSet) SetStartDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStartDate:"), value)
}/* debug [instance_properties/setter]: startDate */


// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/stateOrProvince
func (c_ CSSearchableItemAttributeSet) StateOrProvince() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("stateOrProvince"))
	return rv
}/* debug [instance_properties/getter]: stateOrProvince */


// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/stateOrProvince
func (c_ CSSearchableItemAttributeSet) SetStateOrProvince(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStateOrProvince:"), value)
}/* debug [instance_properties/setter]: stateOrProvince */


// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/streamable
func (c_ CSSearchableItemAttributeSet) Streamable() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("streamable"))
	return rv
}/* debug [instance_properties/getter]: streamable */


// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/streamable
func (c_ CSSearchableItemAttributeSet) SetStreamable(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreamable:"), value)
}/* debug [instance_properties/setter]: streamable */


// The subject of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subject
func (c_ CSSearchableItemAttributeSet) Subject() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subject"))
	return rv
}/* debug [instance_properties/getter]: subject */


// The subject of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subject
func (c_ CSSearchableItemAttributeSet) SetSubject(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubject:"), value)
}/* debug [instance_properties/setter]: subject */


// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subThoroughfare
func (c_ CSSearchableItemAttributeSet) SubThoroughfare() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subThoroughfare"))
	return rv
}/* debug [instance_properties/getter]: subThoroughfare */


// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/subThoroughfare
func (c_ CSSearchableItemAttributeSet) SetSubThoroughfare(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubThoroughfare:"), value)
}/* debug [instance_properties/setter]: subThoroughfare */


// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation
func (c_ CSSearchableItemAttributeSet) SupportsNavigation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("supportsNavigation"))
	return rv
}/* debug [instance_properties/getter]: supportsNavigation */


// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation
func (c_ CSSearchableItemAttributeSet) SetSupportsNavigation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsNavigation:"), value)
}/* debug [instance_properties/setter]: supportsNavigation */


// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsPhoneCall
func (c_ CSSearchableItemAttributeSet) SupportsPhoneCall() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("supportsPhoneCall"))
	return rv
}/* debug [instance_properties/getter]: supportsPhoneCall */


// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsPhoneCall
func (c_ CSSearchableItemAttributeSet) SetSupportsPhoneCall(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsPhoneCall:"), value)
}/* debug [instance_properties/setter]: supportsPhoneCall */


// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/tempo
func (c_ CSSearchableItemAttributeSet) Tempo() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("tempo"))
	return rv
}/* debug [instance_properties/getter]: tempo */


// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/tempo
func (c_ CSSearchableItemAttributeSet) SetTempo(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTempo:"), value)
}/* debug [instance_properties/setter]: tempo */


// The textual content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContent
func (c_ CSSearchableItemAttributeSet) TextContent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("textContent"))
	return rv
}/* debug [instance_properties/getter]: textContent */


// The textual content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContent
func (c_ CSSearchableItemAttributeSet) SetTextContent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextContent:"), value)
}/* debug [instance_properties/setter]: textContent */


// A string that presents the Apple Intelligence summarization of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/textContentSummary
func (c_ CSSearchableItemAttributeSet) TextContentSummary() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("textContentSummary"))
	return rv
}/* debug [instance_properties/getter]: textContentSummary */


// The theme of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/theme
func (c_ CSSearchableItemAttributeSet) Theme() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("theme"))
	return rv
}/* debug [instance_properties/getter]: theme */


// The theme of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/theme
func (c_ CSSearchableItemAttributeSet) SetTheme(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTheme:"), value)
}/* debug [instance_properties/setter]: theme */


// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thoroughfare
func (c_ CSSearchableItemAttributeSet) Thoroughfare() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("thoroughfare"))
	return rv
}/* debug [instance_properties/getter]: thoroughfare */


// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thoroughfare
func (c_ CSSearchableItemAttributeSet) SetThoroughfare(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThoroughfare:"), value)
}/* debug [instance_properties/setter]: thoroughfare */


// Image data that represents the thumbnail of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailData
func (c_ CSSearchableItemAttributeSet) ThumbnailData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("thumbnailData"))
	return rv
}/* debug [instance_properties/getter]: thumbnailData */


// Image data that represents the thumbnail of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailData
func (c_ CSSearchableItemAttributeSet) SetThumbnailData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailData:"), value)
}/* debug [instance_properties/setter]: thumbnailData */


// The local file URL of the thumbnail image for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailURL
func (c_ CSSearchableItemAttributeSet) ThumbnailURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("thumbnailURL"))
	return rv
}/* debug [instance_properties/getter]: thumbnailURL */


// The local file URL of the thumbnail image for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/thumbnailURL
func (c_ CSSearchableItemAttributeSet) SetThumbnailURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailURL:"), value)
}/* debug [instance_properties/setter]: thumbnailURL */


// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timeSignature
func (c_ CSSearchableItemAttributeSet) TimeSignature() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("timeSignature"))
	return rv
}/* debug [instance_properties/getter]: timeSignature */


// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timeSignature
func (c_ CSSearchableItemAttributeSet) SetTimeSignature(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeSignature:"), value)
}/* debug [instance_properties/setter]: timeSignature */


// The timestamp on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timestamp
func (c_ CSSearchableItemAttributeSet) Timestamp() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("timestamp"))
	return rv
}/* debug [instance_properties/getter]: timestamp */


// The timestamp on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/timestamp
func (c_ CSSearchableItemAttributeSet) SetTimestamp(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimestamp:"), value)
}/* debug [instance_properties/setter]: timestamp */


// The title of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title
func (c_ CSSearchableItemAttributeSet) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The title of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title
func (c_ CSSearchableItemAttributeSet) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The total bit rate of the media, combining audio and video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/totalBitRate
func (c_ CSSearchableItemAttributeSet) TotalBitRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("totalBitRate"))
	return rv
}/* debug [instance_properties/getter]: totalBitRate */


// The total bit rate of the media, combining audio and video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/totalBitRate
func (c_ CSSearchableItemAttributeSet) SetTotalBitRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalBitRate:"), value)
}/* debug [instance_properties/setter]: totalBitRate */


// A string that represents the text the system transcribed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/transcribedTextContent
func (c_ CSSearchableItemAttributeSet) TranscribedTextContent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("transcribedTextContent"))
	return rv
}/* debug [instance_properties/getter]: transcribedTextContent */


// A string that represents the text the system transcribed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/transcribedTextContent
func (c_ CSSearchableItemAttributeSet) SetTranscribedTextContent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTranscribedTextContent:"), value)
}/* debug [instance_properties/setter]: transcribedTextContent */


// The URL associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/url
func (c_ CSSearchableItemAttributeSet) URL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](c_.ID, objc.Sel("URL"))
	return rv
}/* debug [instance_properties/getter]: URL */


// The URL associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/url
func (c_ CSSearchableItemAttributeSet) SetURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setURL:"), value)
}/* debug [instance_properties/setter]: URL */


// A value that indicates the user created the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCreated
func (c_ CSSearchableItemAttributeSet) UserCreated() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("userCreated"))
	return rv
}/* debug [instance_properties/getter]: userCreated */


// A value that indicates the user created the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCreated
func (c_ CSSearchableItemAttributeSet) SetUserCreated(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCreated:"), value)
}/* debug [instance_properties/setter]: userCreated */


// A value that indicates the user selected the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCurated
func (c_ CSSearchableItemAttributeSet) UserCurated() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("userCurated"))
	return rv
}/* debug [instance_properties/getter]: userCurated */


// A value that indicates the user selected the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userCurated
func (c_ CSSearchableItemAttributeSet) SetUserCurated(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCurated:"), value)
}/* debug [instance_properties/setter]: userCurated */


// A value that indicates the user purchased or owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userOwned
func (c_ CSSearchableItemAttributeSet) UserOwned() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("userOwned"))
	return rv
}/* debug [instance_properties/getter]: userOwned */


// A value that indicates the user purchased or owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/userOwned
func (c_ CSSearchableItemAttributeSet) SetUserOwned(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserOwned:"), value)
}/* debug [instance_properties/setter]: userOwned */


// A version string associated with the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/version
func (c_ CSSearchableItemAttributeSet) Version() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("version"))
	return rv
}/* debug [instance_properties/getter]: version */


// A version string associated with the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/version
func (c_ CSSearchableItemAttributeSet) SetVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVersion:"), value)
}/* debug [instance_properties/setter]: version */


// The video bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/videoBitRate
func (c_ CSSearchableItemAttributeSet) VideoBitRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("videoBitRate"))
	return rv
}/* debug [instance_properties/getter]: videoBitRate */


// The video bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/videoBitRate
func (c_ CSSearchableItemAttributeSet) SetVideoBitRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoBitRate:"), value)
}/* debug [instance_properties/setter]: videoBitRate */


// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/weakRelatedUniqueIdentifier
func (c_ CSSearchableItemAttributeSet) WeakRelatedUniqueIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("weakRelatedUniqueIdentifier"))
	return rv
}/* debug [instance_properties/getter]: weakRelatedUniqueIdentifier */


// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/weakRelatedUniqueIdentifier
func (c_ CSSearchableItemAttributeSet) SetWeakRelatedUniqueIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeakRelatedUniqueIdentifier:"), value)
}/* debug [instance_properties/setter]: weakRelatedUniqueIdentifier */


// The white balance setting when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/whiteBalance
func (c_ CSSearchableItemAttributeSet) WhiteBalance() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("whiteBalance"))
	return rv
}/* debug [instance_properties/getter]: whiteBalance */


// The white balance setting when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/whiteBalance
func (c_ CSSearchableItemAttributeSet) SetWhiteBalance(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteBalance:"), value)
}/* debug [instance_properties/setter]: whiteBalance */


// A key that specifies the action’s identifier in a user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csactionidentifier
func (c_ CSSearchableItemAttributeSet) CSActionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CSActionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: CSActionIdentifier */


// The composer of the song or audio composition that the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/composer
func (c_ CSSearchableItemAttributeSet) Composer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("composer"))
	return rv
}/* debug [instance_properties/getter]: composer */


// The composer of the song or audio composition that the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/composer
func (c_ CSSearchableItemAttributeSet) SetComposer(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComposer:"), value)
}/* debug [instance_properties/setter]: composer */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CSSearchableItemAttributeSet */


