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
	ContentModificationDate() objc.IObject /* cross-framework: NSDate */
	SetContentModificationDate(value objc.IObject /* cross-framework: NSDate */)
	DisplayName() objc.IObject /* cross-framework: NSString */
	SetDisplayName(value objc.IObject /* cross-framework: NSString */)
	EncodingApplications() []string
	SetEncodingApplications(value []string)
	EndDate() objc.IObject /* cross-framework: NSDate */
	SetEndDate(value objc.IObject /* cross-framework: NSDate */)
	GPSDOP() objc.IObject /* cross-framework: NSNumber */
	SetGPSDOP(value objc.IObject /* cross-framework: NSNumber */)
	SupportsNavigation() objc.IObject /* cross-framework: NSNumber */
	SetSupportsNavigation(value objc.IObject /* cross-framework: NSNumber */)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	CSActionIdentifier() objc.IObject /* cross-framework: NSString */
	AccountHandles() objc.IObject /* cross-framework: NSString */
	SetAccountHandles(value objc.IObject /* cross-framework: NSString */)
	AccountIdentifier() objc.IObject /* cross-framework: NSString */
	SetAccountIdentifier(value objc.IObject /* cross-framework: NSString */)
	AcquisitionMake() objc.IObject /* cross-framework: NSString */
	SetAcquisitionMake(value objc.IObject /* cross-framework: NSString */)
	AcquisitionModel() objc.IObject /* cross-framework: NSString */
	SetAcquisitionModel(value objc.IObject /* cross-framework: NSString */)
	ActionIdentifiers() objc.IObject /* cross-framework: NSString */
	SetActionIdentifiers(value objc.IObject /* cross-framework: NSString */)
	AddedDate() objc.IObject /* cross-framework: Date */
	SetAddedDate(value objc.IObject /* cross-framework: Date */)
	AdditionalRecipients() ICSPerson
	SetAdditionalRecipients(value ICSPerson)
	Album() objc.IObject /* cross-framework: NSString */
	SetAlbum(value objc.IObject /* cross-framework: NSString */)
	AllDay() objc.IObject /* cross-framework: NSNumber */
	SetAllDay(value objc.IObject /* cross-framework: NSNumber */)
	AlternateNames() objc.IObject /* cross-framework: NSString */
	SetAlternateNames(value objc.IObject /* cross-framework: NSString */)
	Altitude() objc.IObject /* cross-framework: NSNumber */
	SetAltitude(value objc.IObject /* cross-framework: NSNumber */)
	Aperture() objc.IObject /* cross-framework: NSNumber */
	SetAperture(value objc.IObject /* cross-framework: NSNumber */)
	Artist() objc.IObject /* cross-framework: NSString */
	SetArtist(value objc.IObject /* cross-framework: NSString */)
	Audiences() objc.IObject /* cross-framework: NSString */
	SetAudiences(value objc.IObject /* cross-framework: NSString */)
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
	AuthorAddresses() objc.IObject /* cross-framework: NSString */
	SetAuthorAddresses(value objc.IObject /* cross-framework: NSString */)
	AuthorEmailAddresses() objc.IObject /* cross-framework: NSString */
	SetAuthorEmailAddresses(value objc.IObject /* cross-framework: NSString */)
	AuthorNames() objc.IObject /* cross-framework: NSString */
	SetAuthorNames(value objc.IObject /* cross-framework: NSString */)
	Authors() ICSPerson
	SetAuthors(value ICSPerson)
	BitsPerSample() objc.IObject /* cross-framework: NSNumber */
	SetBitsPerSample(value objc.IObject /* cross-framework: NSNumber */)
	CameraOwner() objc.IObject /* cross-framework: NSString */
	SetCameraOwner(value objc.IObject /* cross-framework: NSString */)
	City() objc.IObject /* cross-framework: NSString */
	SetCity(value objc.IObject /* cross-framework: NSString */)
	Codecs() objc.IObject /* cross-framework: NSString */
	SetCodecs(value objc.IObject /* cross-framework: NSString */)
	ColorSpace() objc.IObject /* cross-framework: NSString */
	SetColorSpace(value objc.IObject /* cross-framework: NSString */)
	Comment() objc.IObject /* cross-framework: NSString */
	SetComment(value objc.IObject /* cross-framework: NSString */)
	CompletionDate() objc.IObject /* cross-framework: Date */
	SetCompletionDate(value objc.IObject /* cross-framework: Date */)
	Composer() objc.IObject /* cross-framework: NSString */
	SetComposer(value objc.IObject /* cross-framework: NSString */)
	ContactKeywords() objc.IObject /* cross-framework: NSString */
	SetContactKeywords(value objc.IObject /* cross-framework: NSString */)
	ContainerDisplayName() objc.IObject /* cross-framework: NSString */
	SetContainerDisplayName(value objc.IObject /* cross-framework: NSString */)
	ContainerIdentifier() objc.IObject /* cross-framework: NSString */
	SetContainerIdentifier(value objc.IObject /* cross-framework: NSString */)
	ContainerOrder() objc.IObject /* cross-framework: NSNumber */
	SetContainerOrder(value objc.IObject /* cross-framework: NSNumber */)
	ContainerTitle() objc.IObject /* cross-framework: NSString */
	SetContainerTitle(value objc.IObject /* cross-framework: NSString */)
	ContentCreationDate() objc.IObject /* cross-framework: Date */
	SetContentCreationDate(value objc.IObject /* cross-framework: Date */)
	ContentDescription() objc.IObject /* cross-framework: NSString */
	SetContentDescription(value objc.IObject /* cross-framework: NSString */)
	ContentRating() objc.IObject /* cross-framework: NSNumber */
	SetContentRating(value objc.IObject /* cross-framework: NSNumber */)
	ContentSources() objc.IObject /* cross-framework: NSString */
	SetContentSources(value objc.IObject /* cross-framework: NSString */)
	ContentType() objc.IObject /* cross-framework: NSString */
	SetContentType(value objc.IObject /* cross-framework: NSString */)
	ContentTypeTree() objc.IObject /* cross-framework: NSString */
	SetContentTypeTree(value objc.IObject /* cross-framework: NSString */)
	ContentURL() objc.IObject /* cross-framework: URL */
	SetContentURL(value objc.IObject /* cross-framework: URL */)
	Contributors() objc.IObject /* cross-framework: NSString */
	SetContributors(value objc.IObject /* cross-framework: NSString */)
	Copyright() objc.IObject /* cross-framework: NSString */
	SetCopyright(value objc.IObject /* cross-framework: NSString */)
	Country() objc.IObject /* cross-framework: NSString */
	SetCountry(value objc.IObject /* cross-framework: NSString */)
	Coverage() objc.IObject /* cross-framework: NSString */
	SetCoverage(value objc.IObject /* cross-framework: NSString */)
	Creator() objc.IObject /* cross-framework: NSString */
	SetCreator(value objc.IObject /* cross-framework: NSString */)
	DarkThumbnailURL() objc.IObject /* cross-framework: URL */
	SetDarkThumbnailURL(value objc.IObject /* cross-framework: URL */)
	DeliveryType() objc.IObject /* cross-framework: NSNumber */
	SetDeliveryType(value objc.IObject /* cross-framework: NSNumber */)
	Director() objc.IObject /* cross-framework: NSString */
	SetDirector(value objc.IObject /* cross-framework: NSString */)
	DomainIdentifier() objc.IObject /* cross-framework: NSString */
	SetDomainIdentifier(value objc.IObject /* cross-framework: NSString */)
	DownloadedDate() objc.IObject /* cross-framework: Date */
	SetDownloadedDate(value objc.IObject /* cross-framework: Date */)
	DueDate() objc.IObject /* cross-framework: Date */
	SetDueDate(value objc.IObject /* cross-framework: Date */)
	Duration() objc.IObject /* cross-framework: NSNumber */
	SetDuration(value objc.IObject /* cross-framework: NSNumber */)
	Editors() objc.IObject /* cross-framework: NSString */
	SetEditors(value objc.IObject /* cross-framework: NSString */)
	EmailAddresses() objc.IObject /* cross-framework: NSString */
	SetEmailAddresses(value objc.IObject /* cross-framework: NSString */)
	EmailHeaders() objc.IObject /* cross-framework: NSString */
	SetEmailHeaders(value objc.IObject /* cross-framework: NSString */)
	ExifVersion() objc.IObject /* cross-framework: NSString */
	SetExifVersion(value objc.IObject /* cross-framework: NSString */)
	ExifgpsVersion() objc.IObject /* cross-framework: NSString */
	SetExifgpsVersion(value objc.IObject /* cross-framework: NSString */)
	ExposureMode() objc.IObject /* cross-framework: NSNumber */
	SetExposureMode(value objc.IObject /* cross-framework: NSNumber */)
	ExposureProgram() objc.IObject /* cross-framework: NSString */
	SetExposureProgram(value objc.IObject /* cross-framework: NSString */)
	ExposureTime() objc.IObject /* cross-framework: NSNumber */
	SetExposureTime(value objc.IObject /* cross-framework: NSNumber */)
	ExposureTimeString() objc.IObject /* cross-framework: NSString */
	SetExposureTimeString(value objc.IObject /* cross-framework: NSString */)
	FNumber() objc.IObject /* cross-framework: NSNumber */
	SetFNumber(value objc.IObject /* cross-framework: NSNumber */)
	FileSize() objc.IObject /* cross-framework: NSNumber */
	SetFileSize(value objc.IObject /* cross-framework: NSNumber */)
	FlashOn() objc.IObject /* cross-framework: NSNumber */
	SetFlashOn(value objc.IObject /* cross-framework: NSNumber */)
	FocalLength() objc.IObject /* cross-framework: NSNumber */
	SetFocalLength(value objc.IObject /* cross-framework: NSNumber */)
	FocalLength35mm() objc.IObject /* cross-framework: NSNumber */
	SetFocalLength35mm(value objc.IObject /* cross-framework: NSNumber */)
	FontNames() objc.IObject /* cross-framework: NSString */
	SetFontNames(value objc.IObject /* cross-framework: NSString */)
	FullyFormattedAddress() objc.IObject /* cross-framework: NSString */
	SetFullyFormattedAddress(value objc.IObject /* cross-framework: NSString */)
	GeneralMIDISequence() objc.IObject /* cross-framework: NSNumber */
	SetGeneralMIDISequence(value objc.IObject /* cross-framework: NSNumber */)
	Genre() objc.IObject /* cross-framework: NSString */
	SetGenre(value objc.IObject /* cross-framework: NSString */)
	GpsAreaInformation() objc.IObject /* cross-framework: NSString */
	SetGpsAreaInformation(value objc.IObject /* cross-framework: NSString */)
	GpsDateStamp() objc.IObject /* cross-framework: Date */
	SetGpsDateStamp(value objc.IObject /* cross-framework: Date */)
	GpsDestBearing() objc.IObject /* cross-framework: NSNumber */
	SetGpsDestBearing(value objc.IObject /* cross-framework: NSNumber */)
	GpsDestDistance() objc.IObject /* cross-framework: NSNumber */
	SetGpsDestDistance(value objc.IObject /* cross-framework: NSNumber */)
	GpsDestLatitude() objc.IObject /* cross-framework: NSNumber */
	SetGpsDestLatitude(value objc.IObject /* cross-framework: NSNumber */)
	GpsDestLongitude() objc.IObject /* cross-framework: NSNumber */
	SetGpsDestLongitude(value objc.IObject /* cross-framework: NSNumber */)
	GpsDifferental() objc.IObject /* cross-framework: NSNumber */
	SetGpsDifferental(value objc.IObject /* cross-framework: NSNumber */)
	GpsMapDatum() objc.IObject /* cross-framework: NSString */
	SetGpsMapDatum(value objc.IObject /* cross-framework: NSString */)
	GpsMeasureMode() objc.IObject /* cross-framework: NSString */
	SetGpsMeasureMode(value objc.IObject /* cross-framework: NSString */)
	GpsProcessingMethod() objc.IObject /* cross-framework: NSString */
	SetGpsProcessingMethod(value objc.IObject /* cross-framework: NSString */)
	GpsStatus() objc.IObject /* cross-framework: NSString */
	SetGpsStatus(value objc.IObject /* cross-framework: NSString */)
	GpsTrack() objc.IObject /* cross-framework: NSNumber */
	SetGpsTrack(value objc.IObject /* cross-framework: NSNumber */)
	HasAlphaChannel() objc.IObject /* cross-framework: NSNumber */
	SetHasAlphaChannel(value objc.IObject /* cross-framework: NSNumber */)
	Headline() objc.IObject /* cross-framework: NSString */
	SetHeadline(value objc.IObject /* cross-framework: NSString */)
	HiddenAdditionalRecipients() ICSPerson
	SetHiddenAdditionalRecipients(value ICSPerson)
	HtmlContentData() objc.IObject /* cross-framework: Data */
	SetHtmlContentData(value objc.IObject /* cross-framework: Data */)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	ImageDirection() objc.IObject /* cross-framework: NSNumber */
	SetImageDirection(value objc.IObject /* cross-framework: NSNumber */)
	ImportantDates() objc.IObject /* cross-framework: Date */
	SetImportantDates(value objc.IObject /* cross-framework: Date */)
	Information() objc.IObject /* cross-framework: NSString */
	SetInformation(value objc.IObject /* cross-framework: NSString */)
	InstantMessageAddresses() objc.IObject /* cross-framework: NSString */
	SetInstantMessageAddresses(value objc.IObject /* cross-framework: NSString */)
	Instructions() objc.IObject /* cross-framework: NSString */
	SetInstructions(value objc.IObject /* cross-framework: NSString */)
	IsPriority() objc.IObject /* cross-framework: NSNumber */
	SetIsPriority(value objc.IObject /* cross-framework: NSNumber */)
	IsoSpeed() objc.IObject /* cross-framework: NSNumber */
	SetIsoSpeed(value objc.IObject /* cross-framework: NSNumber */)
	KeySignature() objc.IObject /* cross-framework: NSString */
	SetKeySignature(value objc.IObject /* cross-framework: NSString */)
	Keywords() objc.IObject /* cross-framework: NSString */
	SetKeywords(value objc.IObject /* cross-framework: NSString */)
	Kind() objc.IObject /* cross-framework: NSString */
	SetKind(value objc.IObject /* cross-framework: NSString */)
	Languages() objc.IObject /* cross-framework: NSString */
	SetLanguages(value objc.IObject /* cross-framework: NSString */)
	LastUsedDate() objc.IObject /* cross-framework: Date */
	SetLastUsedDate(value objc.IObject /* cross-framework: Date */)
	Latitude() objc.IObject /* cross-framework: NSNumber */
	SetLatitude(value objc.IObject /* cross-framework: NSNumber */)
	LayerNames() objc.IObject /* cross-framework: NSString */
	SetLayerNames(value objc.IObject /* cross-framework: NSString */)
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
	MailboxIdentifiers() objc.IObject /* cross-framework: NSString */
	SetMailboxIdentifiers(value objc.IObject /* cross-framework: NSString */)
	MaxAperture() objc.IObject /* cross-framework: NSNumber */
	SetMaxAperture(value objc.IObject /* cross-framework: NSNumber */)
	MediaTypes() objc.IObject /* cross-framework: NSString */
	SetMediaTypes(value objc.IObject /* cross-framework: NSString */)
	MetadataModificationDate() objc.IObject /* cross-framework: Date */
	SetMetadataModificationDate(value objc.IObject /* cross-framework: Date */)
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
	Organizations() objc.IObject /* cross-framework: NSString */
	SetOrganizations(value objc.IObject /* cross-framework: NSString */)
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
	Participants() objc.IObject /* cross-framework: NSString */
	SetParticipants(value objc.IObject /* cross-framework: NSString */)
	Path() objc.IObject /* cross-framework: NSString */
	SetPath(value objc.IObject /* cross-framework: NSString */)
	Performers() objc.IObject /* cross-framework: NSString */
	SetPerformers(value objc.IObject /* cross-framework: NSString */)
	PhoneNumbers() objc.IObject /* cross-framework: NSString */
	SetPhoneNumbers(value objc.IObject /* cross-framework: NSString */)
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
	PrimaryRecipients() ICSPerson
	SetPrimaryRecipients(value ICSPerson)
	Producer() objc.IObject /* cross-framework: NSString */
	SetProducer(value objc.IObject /* cross-framework: NSString */)
	ProfileName() objc.IObject /* cross-framework: NSString */
	SetProfileName(value objc.IObject /* cross-framework: NSString */)
	Projects() objc.IObject /* cross-framework: NSString */
	SetProjects(value objc.IObject /* cross-framework: NSString */)
	ProviderDataTypeIdentifiers() objc.IObject /* cross-framework: NSString */
	SetProviderDataTypeIdentifiers(value objc.IObject /* cross-framework: NSString */)
	ProviderFileTypeIdentifiers() objc.IObject /* cross-framework: NSString */
	SetProviderFileTypeIdentifiers(value objc.IObject /* cross-framework: NSString */)
	ProviderInPlaceFileTypeIdentifiers() objc.IObject /* cross-framework: NSString */
	SetProviderInPlaceFileTypeIdentifiers(value objc.IObject /* cross-framework: NSString */)
	Publishers() objc.IObject /* cross-framework: NSString */
	SetPublishers(value objc.IObject /* cross-framework: NSString */)
	RankingHint() objc.IObject /* cross-framework: NSNumber */
	SetRankingHint(value objc.IObject /* cross-framework: NSNumber */)
	Rating() objc.IObject /* cross-framework: NSNumber */
	SetRating(value objc.IObject /* cross-framework: NSNumber */)
	RatingDescription() objc.IObject /* cross-framework: NSString */
	SetRatingDescription(value objc.IObject /* cross-framework: NSString */)
	RecipientAddresses() objc.IObject /* cross-framework: NSString */
	SetRecipientAddresses(value objc.IObject /* cross-framework: NSString */)
	RecipientEmailAddresses() objc.IObject /* cross-framework: NSString */
	SetRecipientEmailAddresses(value objc.IObject /* cross-framework: NSString */)
	RecipientNames() objc.IObject /* cross-framework: NSString */
	SetRecipientNames(value objc.IObject /* cross-framework: NSString */)
	RecordingDate() objc.IObject /* cross-framework: Date */
	SetRecordingDate(value objc.IObject /* cross-framework: Date */)
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
	SharedItemContentType() objc.IObject /* cross-framework: UTType */
	SetSharedItemContentType(value objc.IObject /* cross-framework: UTType */)
	Speed() objc.IObject /* cross-framework: NSNumber */
	SetSpeed(value objc.IObject /* cross-framework: NSNumber */)
	StartDate() objc.IObject /* cross-framework: Date */
	SetStartDate(value objc.IObject /* cross-framework: Date */)
	StateOrProvince() objc.IObject /* cross-framework: NSString */
	SetStateOrProvince(value objc.IObject /* cross-framework: NSString */)
	Streamable() objc.IObject /* cross-framework: NSNumber */
	SetStreamable(value objc.IObject /* cross-framework: NSNumber */)
	SubThoroughfare() objc.IObject /* cross-framework: NSString */
	SetSubThoroughfare(value objc.IObject /* cross-framework: NSString */)
	Subject() objc.IObject /* cross-framework: NSString */
	SetSubject(value objc.IObject /* cross-framework: NSString */)
	SupportsPhoneCall() objc.IObject /* cross-framework: NSNumber */
	SetSupportsPhoneCall(value objc.IObject /* cross-framework: NSNumber */)
	Tempo() objc.IObject /* cross-framework: NSNumber */
	SetTempo(value objc.IObject /* cross-framework: NSNumber */)
	TextContent() objc.IObject /* cross-framework: NSString */
	SetTextContent(value objc.IObject /* cross-framework: NSString */)
	TextContentSummary() objc.IObject /* cross-framework: NSString */
	SetTextContentSummary(value objc.IObject /* cross-framework: NSString */)
	Theme() objc.IObject /* cross-framework: NSString */
	SetTheme(value objc.IObject /* cross-framework: NSString */)
	Thoroughfare() objc.IObject /* cross-framework: NSString */
	SetThoroughfare(value objc.IObject /* cross-framework: NSString */)
	ThumbnailData() objc.IObject /* cross-framework: Data */
	SetThumbnailData(value objc.IObject /* cross-framework: Data */)
	ThumbnailURL() objc.IObject /* cross-framework: URL */
	SetThumbnailURL(value objc.IObject /* cross-framework: URL */)
	TimeSignature() objc.IObject /* cross-framework: NSString */
	SetTimeSignature(value objc.IObject /* cross-framework: NSString */)
	Timestamp() objc.IObject /* cross-framework: Date */
	SetTimestamp(value objc.IObject /* cross-framework: Date */)
	TotalBitRate() objc.IObject /* cross-framework: NSNumber */
	SetTotalBitRate(value objc.IObject /* cross-framework: NSNumber */)
	TranscribedTextContent() objc.IObject /* cross-framework: NSString */
	SetTranscribedTextContent(value objc.IObject /* cross-framework: NSString */)
	Url() objc.IObject /* cross-framework: URL */
	SetUrl(value objc.IObject /* cross-framework: URL */)
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
func (c_ CSSearchableItemAttributeSet) ContentModificationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("contentModificationDate"))
	return rv
}


// The date on which the contents of the file was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/contentModificationDate
func (c_ CSSearchableItemAttributeSet) SetContentModificationDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentModificationDate:"), value)
}


// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName
func (c_ CSSearchableItemAttributeSet) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("displayName"))
	return rv
}


// A localized string that contains the name of the item, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/displayName
func (c_ CSSearchableItemAttributeSet) SetDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDisplayName:"), value)
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
func (c_ CSSearchableItemAttributeSet) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("endDate"))
	return rv
}


// The end date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/endDate
func (c_ CSSearchableItemAttributeSet) SetEndDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEndDate:"), value)
}


// The GPS dilution of precision value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop
func (c_ CSSearchableItemAttributeSet) GPSDOP() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("GPSDOP"))
	return rv
}


// The GPS dilution of precision value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/gpsdop
func (c_ CSSearchableItemAttributeSet) SetGPSDOP(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGPSDOP:"), value)
}


// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation
func (c_ CSSearchableItemAttributeSet) SupportsNavigation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("supportsNavigation"))
	return rv
}


// A value that indicates whether the item contains information sufficient to provide navigation to the location it represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/supportsNavigation
func (c_ CSSearchableItemAttributeSet) SetSupportsNavigation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsNavigation:"), value)
}


// The title of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title
func (c_ CSSearchableItemAttributeSet) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("title"))
	return rv
}


// The title of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemAttributeSet/title
func (c_ CSSearchableItemAttributeSet) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), value)
}


// A key that specifies the action’s identifier in a user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/csactionidentifier
func (c_ CSSearchableItemAttributeSet) CSActionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CSActionIdentifier"))
	return rv
}


// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accounthandles
func (c_ CSSearchableItemAttributeSet) AccountHandles() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("accountHandles"))
	return rv
}


// An array of the canonical handles for the account with which the message is associated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accounthandles
func (c_ CSSearchableItemAttributeSet) SetAccountHandles(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccountHandles:"), value)
}


// The unique identifier for the account with which the message is associated, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accountidentifier
func (c_ CSSearchableItemAttributeSet) AccountIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("accountIdentifier"))
	return rv
}


// The unique identifier for the account with which the message is associated, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/accountidentifier
func (c_ CSSearchableItemAttributeSet) SetAccountIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAccountIdentifier:"), value)
}


// The manufacturer of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmake
func (c_ CSSearchableItemAttributeSet) AcquisitionMake() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("acquisitionMake"))
	return rv
}


// The manufacturer of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmake
func (c_ CSSearchableItemAttributeSet) SetAcquisitionMake(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionMake:"), value)
}


// The model of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmodel
func (c_ CSSearchableItemAttributeSet) AcquisitionModel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("acquisitionModel"))
	return rv
}


// The model of the device that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/acquisitionmodel
func (c_ CSSearchableItemAttributeSet) SetAcquisitionModel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAcquisitionModel:"), value)
}


// The identifiers that specify custom actions the app supports for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/actionidentifiers
func (c_ CSSearchableItemAttributeSet) ActionIdentifiers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("actionIdentifiers"))
	return rv
}


// The identifiers that specify custom actions the app supports for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/actionidentifiers
func (c_ CSSearchableItemAttributeSet) SetActionIdentifiers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setActionIdentifiers:"), value)
}


// The date on which the item was moved into its current location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/addeddate
func (c_ CSSearchableItemAttributeSet) AddedDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("addedDate"))
	return rv
}


// The date on which the item was moved into its current location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/addeddate
func (c_ CSSearchableItemAttributeSet) SetAddedDate(value objc.IObject /* cross-framework: Date */) {
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
func (c_ CSSearchableItemAttributeSet) Album() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("album"))
	return rv
}


// The title for a collection of audio media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/album
func (c_ CSSearchableItemAttributeSet) SetAlbum(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlbum:"), value)
}


// A value that indicates if the event covers an entire day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/allday
func (c_ CSSearchableItemAttributeSet) AllDay() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("allDay"))
	return rv
}


// A value that indicates if the event covers an entire day.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/allday
func (c_ CSSearchableItemAttributeSet) SetAllDay(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllDay:"), value)
}


// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/alternatenames
func (c_ CSSearchableItemAttributeSet) AlternateNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("alternateNames"))
	return rv
}


// An array of localized strings that represent alternate display names for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/alternatenames
func (c_ CSSearchableItemAttributeSet) SetAlternateNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAlternateNames:"), value)
}


// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/altitude
func (c_ CSSearchableItemAttributeSet) Altitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("altitude"))
	return rv
}


// The altitude of the item in meters above sea level, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/altitude
func (c_ CSSearchableItemAttributeSet) SetAltitude(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAltitude:"), value)
}


// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/aperture
func (c_ CSSearchableItemAttributeSet) Aperture() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("aperture"))
	return rv
}


// The size of the lens aperture at the time the camera captured the image, as a log-scale APEX value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/aperture
func (c_ CSSearchableItemAttributeSet) SetAperture(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAperture:"), value)
}


// The artist associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/artist
func (c_ CSSearchableItemAttributeSet) Artist() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("artist"))
	return rv
}


// The artist associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/artist
func (c_ CSSearchableItemAttributeSet) SetArtist(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setArtist:"), value)
}


// A class of entity for which the item is intended or useful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiences
func (c_ CSSearchableItemAttributeSet) Audiences() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("audiences"))
	return rv
}


// A class of entity for which the item is intended or useful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiences
func (c_ CSSearchableItemAttributeSet) SetAudiences(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudiences:"), value)
}


// The audio bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiobitrate
func (c_ CSSearchableItemAttributeSet) AudioBitRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("audioBitRate"))
	return rv
}


// The audio bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiobitrate
func (c_ CSSearchableItemAttributeSet) SetAudioBitRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioBitRate:"), value)
}


// The number of channels in the audio data that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiochannelcount
func (c_ CSSearchableItemAttributeSet) AudioChannelCount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("audioChannelCount"))
	return rv
}


// The number of channels in the audio data that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiochannelcount
func (c_ CSSearchableItemAttributeSet) SetAudioChannelCount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioChannelCount:"), value)
}


// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audioencodingapplication
func (c_ CSSearchableItemAttributeSet) AudioEncodingApplication() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("audioEncodingApplication"))
	return rv
}


// The name of the application that encoded the data the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audioencodingapplication
func (c_ CSSearchableItemAttributeSet) SetAudioEncodingApplication(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioEncodingApplication:"), value)
}


// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiosamplerate
func (c_ CSSearchableItemAttributeSet) AudioSampleRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("audioSampleRate"))
	return rv
}


// The sample rate of the audio data the file contains, as a float value representing Hz (audio frames per second), such as 44100.0 or 22254.54.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiosamplerate
func (c_ CSSearchableItemAttributeSet) SetAudioSampleRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioSampleRate:"), value)
}


// The track number of a song or audio composition when part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiotracknumber
func (c_ CSSearchableItemAttributeSet) AudioTrackNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("audioTrackNumber"))
	return rv
}


// The track number of a song or audio composition when part of an album.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/audiotracknumber
func (c_ CSSearchableItemAttributeSet) SetAudioTrackNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAudioTrackNumber:"), value)
}


// An array of addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoraddresses
func (c_ CSSearchableItemAttributeSet) AuthorAddresses() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("authorAddresses"))
	return rv
}


// An array of addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoraddresses
func (c_ CSSearchableItemAttributeSet) SetAuthorAddresses(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorAddresses:"), value)
}


// An array of email addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoremailaddresses
func (c_ CSSearchableItemAttributeSet) AuthorEmailAddresses() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("authorEmailAddresses"))
	return rv
}


// An array of email addresses associated with the author of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authoremailaddresses
func (c_ CSSearchableItemAttributeSet) SetAuthorEmailAddresses(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorEmailAddresses:"), value)
}


// An array of names representing the authors who have worked on the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authornames
func (c_ CSSearchableItemAttributeSet) AuthorNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("authorNames"))
	return rv
}


// An array of names representing the authors who have worked on the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/authornames
func (c_ CSSearchableItemAttributeSet) SetAuthorNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAuthorNames:"), value)
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
func (c_ CSSearchableItemAttributeSet) BitsPerSample() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("bitsPerSample"))
	return rv
}


// The number of bits per sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/bitspersample
func (c_ CSSearchableItemAttributeSet) SetBitsPerSample(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBitsPerSample:"), value)
}


// The owner of the camera that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/cameraowner
func (c_ CSSearchableItemAttributeSet) CameraOwner() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("cameraOwner"))
	return rv
}


// The owner of the camera that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/cameraowner
func (c_ CSSearchableItemAttributeSet) SetCameraOwner(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCameraOwner:"), value)
}


// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/city
func (c_ CSSearchableItemAttributeSet) City() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("city"))
	return rv
}


// The city of the item’s origin according to guidelines that the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/city
func (c_ CSSearchableItemAttributeSet) SetCity(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCity:"), value)
}


// The codecs used to encode/decode the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/codecs
func (c_ CSSearchableItemAttributeSet) Codecs() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("codecs"))
	return rv
}


// The codecs used to encode/decode the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/codecs
func (c_ CSSearchableItemAttributeSet) SetCodecs(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCodecs:"), value)
}


// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/colorspace
func (c_ CSSearchableItemAttributeSet) ColorSpace() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("colorSpace"))
	return rv
}


// The color space model the image uses, such as RGB, CMYK, YUV, or YCbCr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/colorspace
func (c_ CSSearchableItemAttributeSet) SetColorSpace(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setColorSpace:"), value)
}


// A comment related to the media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/comment
func (c_ CSSearchableItemAttributeSet) Comment() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("comment"))
	return rv
}


// A comment related to the media file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/comment
func (c_ CSSearchableItemAttributeSet) SetComment(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComment:"), value)
}


// The date on which the item was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/completiondate
func (c_ CSSearchableItemAttributeSet) CompletionDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("completionDate"))
	return rv
}


// The date on which the item was completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/completiondate
func (c_ CSSearchableItemAttributeSet) SetCompletionDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCompletionDate:"), value)
}


// The composer of the song or audio composition that the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/composer
func (c_ CSSearchableItemAttributeSet) Composer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("composer"))
	return rv
}


// The composer of the song or audio composition that the audio file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/composer
func (c_ CSSearchableItemAttributeSet) SetComposer(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setComposer:"), value)
}


// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contactkeywords
func (c_ CSSearchableItemAttributeSet) ContactKeywords() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contactKeywords"))
	return rv
}


// A list of contacts who are associated with the content in some way, not including the author.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contactkeywords
func (c_ CSSearchableItemAttributeSet) SetContactKeywords(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContactKeywords:"), value)
}


// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerdisplayname
func (c_ CSSearchableItemAttributeSet) ContainerDisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerDisplayName"))
	return rv
}


// A localized string that specifies the name of a container to which the item belongs, suitable to display in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerdisplayname
func (c_ CSSearchableItemAttributeSet) SetContainerDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerDisplayName:"), value)
}


// The identifier of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containeridentifier
func (c_ CSSearchableItemAttributeSet) ContainerIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerIdentifier"))
	return rv
}


// The identifier of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containeridentifier
func (c_ CSSearchableItemAttributeSet) SetContainerIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerIdentifier:"), value)
}


// The order of the item within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerorder
func (c_ CSSearchableItemAttributeSet) ContainerOrder() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("containerOrder"))
	return rv
}


// The order of the item within the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containerorder
func (c_ CSSearchableItemAttributeSet) SetContainerOrder(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerOrder:"), value)
}


// The title of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containertitle
func (c_ CSSearchableItemAttributeSet) ContainerTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("containerTitle"))
	return rv
}


// The title of the container to which the item belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/containertitle
func (c_ CSSearchableItemAttributeSet) SetContainerTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContainerTitle:"), value)
}


// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentcreationdate
func (c_ CSSearchableItemAttributeSet) ContentCreationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("contentCreationDate"))
	return rv
}


// The creation date of an edited or optimized version of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentcreationdate
func (c_ CSSearchableItemAttributeSet) SetContentCreationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentCreationDate:"), value)
}


// A description of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentdescription
func (c_ CSSearchableItemAttributeSet) ContentDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contentDescription"))
	return rv
}


// A description of the item’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentdescription
func (c_ CSSearchableItemAttributeSet) SetContentDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentDescription:"), value)
}


// A value that indicates if the media contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentrating
func (c_ CSSearchableItemAttributeSet) ContentRating() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("contentRating"))
	return rv
}


// A value that indicates if the media contains explicit content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentrating
func (c_ CSSearchableItemAttributeSet) SetContentRating(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentRating:"), value)
}


// An array of sources from which the media was obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentsources
func (c_ CSSearchableItemAttributeSet) ContentSources() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contentSources"))
	return rv
}


// An array of sources from which the media was obtained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contentsources
func (c_ CSSearchableItemAttributeSet) SetContentSources(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentSources:"), value)
}


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttype
func (c_ CSSearchableItemAttributeSet) ContentType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contentType"))
	return rv
}


// The uniform type identifier (UTI) of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttype
func (c_ CSSearchableItemAttributeSet) SetContentType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentType:"), value)
}


// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttypetree
func (c_ CSSearchableItemAttributeSet) ContentTypeTree() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contentTypeTree"))
	return rv
}


// An attribute type that identifies a custom hierarchy of types to describe the attributes of your item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenttypetree
func (c_ CSSearchableItemAttributeSet) SetContentTypeTree(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentTypeTree:"), value)
}


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenturl
func (c_ CSSearchableItemAttributeSet) ContentURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("contentURL"))
	return rv
}


// The file URL of the content to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contenturl
func (c_ CSSearchableItemAttributeSet) SetContentURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContentURL:"), value)
}


// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contributors
func (c_ CSSearchableItemAttributeSet) Contributors() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("contributors"))
	return rv
}


// A list of people, organizations, or services that made contributions to the media content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/contributors
func (c_ CSSearchableItemAttributeSet) SetContributors(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setContributors:"), value)
}


// The copyright date of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/copyright
func (c_ CSSearchableItemAttributeSet) Copyright() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("copyright"))
	return rv
}


// The copyright date of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/copyright
func (c_ CSSearchableItemAttributeSet) SetCopyright(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCopyright:"), value)
}


// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/country
func (c_ CSSearchableItemAttributeSet) Country() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("country"))
	return rv
}


// The full, publishable name of the country or region in which the intellectual property of the item was created, according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/country
func (c_ CSSearchableItemAttributeSet) SetCountry(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCountry:"), value)
}


// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/coverage
func (c_ CSSearchableItemAttributeSet) Coverage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("coverage"))
	return rv
}


// A list of descriptors that specify the extent or scope of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/coverage
func (c_ CSSearchableItemAttributeSet) SetCoverage(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCoverage:"), value)
}


// The name of the app that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/creator
func (c_ CSSearchableItemAttributeSet) Creator() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("creator"))
	return rv
}


// The name of the app that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/creator
func (c_ CSSearchableItemAttributeSet) SetCreator(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCreator:"), value)
}


// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/darkthumbnailurl
func (c_ CSSearchableItemAttributeSet) DarkThumbnailURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("darkThumbnailURL"))
	return rv
}


// The local file URL of the thumbnail image for the item when Dark Mode is active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/darkthumbnailurl
func (c_ CSSearchableItemAttributeSet) SetDarkThumbnailURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDarkThumbnailURL:"), value)
}


// The delivery type of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/deliverytype
func (c_ CSSearchableItemAttributeSet) DeliveryType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("deliveryType"))
	return rv
}


// The delivery type of the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/deliverytype
func (c_ CSSearchableItemAttributeSet) SetDeliveryType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDeliveryType:"), value)
}


// The name of the director of the media (for example, a movie director).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/director
func (c_ CSSearchableItemAttributeSet) Director() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("director"))
	return rv
}


// The name of the director of the media (for example, a movie director).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/director
func (c_ CSSearchableItemAttributeSet) SetDirector(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDirector:"), value)
}


// An identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/domainidentifier
func (c_ CSSearchableItemAttributeSet) DomainIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("domainIdentifier"))
	return rv
}


// An identifier that represents the domain or owner of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/domainidentifier
func (c_ CSSearchableItemAttributeSet) SetDomainIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDomainIdentifier:"), value)
}


// The most recent date on which the file was downloaded or received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/downloadeddate
func (c_ CSSearchableItemAttributeSet) DownloadedDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("downloadedDate"))
	return rv
}


// The most recent date on which the file was downloaded or received.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/downloadeddate
func (c_ CSSearchableItemAttributeSet) SetDownloadedDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDownloadedDate:"), value)
}


// The date on which the item is due.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duedate
func (c_ CSSearchableItemAttributeSet) DueDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("dueDate"))
	return rv
}


// The date on which the item is due.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duedate
func (c_ CSSearchableItemAttributeSet) SetDueDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDueDate:"), value)
}


// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duration
func (c_ CSSearchableItemAttributeSet) Duration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("duration"))
	return rv
}


// The duration (if appropriate) of the content of the file, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/duration
func (c_ CSSearchableItemAttributeSet) SetDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDuration:"), value)
}


// A list of editors who have worked on the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/editors
func (c_ CSSearchableItemAttributeSet) Editors() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("editors"))
	return rv
}


// A list of editors who have worked on the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/editors
func (c_ CSSearchableItemAttributeSet) SetEditors(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEditors:"), value)
}


// An array of email addresses associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailaddresses
func (c_ CSSearchableItemAttributeSet) EmailAddresses() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("emailAddresses"))
	return rv
}


// An array of email addresses associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailaddresses
func (c_ CSSearchableItemAttributeSet) SetEmailAddresses(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailAddresses:"), value)
}


// A dictionary that contains all the headers of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailheaders
func (c_ CSSearchableItemAttributeSet) EmailHeaders() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("emailHeaders"))
	return rv
}


// A dictionary that contains all the headers of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/emailheaders
func (c_ CSSearchableItemAttributeSet) SetEmailHeaders(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setEmailHeaders:"), value)
}


// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifversion
func (c_ CSSearchableItemAttributeSet) ExifVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("exifVersion"))
	return rv
}


// The version of the EXIF header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifversion
func (c_ CSSearchableItemAttributeSet) SetExifVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExifVersion:"), value)
}


// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifgpsversion
func (c_ CSSearchableItemAttributeSet) ExifgpsVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("exifgpsVersion"))
	return rv
}


// The version of GPS Info IFD header that was used to generate the metadata for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exifgpsversion
func (c_ CSSearchableItemAttributeSet) SetExifgpsVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExifgpsVersion:"), value)
}


// The mode the camera used for the exposure of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuremode
func (c_ CSSearchableItemAttributeSet) ExposureMode() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("exposureMode"))
	return rv
}


// The mode the camera used for the exposure of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuremode
func (c_ CSSearchableItemAttributeSet) SetExposureMode(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureMode:"), value)
}


// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposureprogram
func (c_ CSSearchableItemAttributeSet) ExposureProgram() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("exposureProgram"))
	return rv
}


// The class of the program the camera used to set exposure when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposureprogram
func (c_ CSSearchableItemAttributeSet) SetExposureProgram(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureProgram:"), value)
}


// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretime
func (c_ CSSearchableItemAttributeSet) ExposureTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("exposureTime"))
	return rv
}


// The time that the lens was open during exposure, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretime
func (c_ CSSearchableItemAttributeSet) SetExposureTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTime:"), value)
}


// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretimestring
func (c_ CSSearchableItemAttributeSet) ExposureTimeString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("exposureTimeString"))
	return rv
}


// The time that the lens was open during exposure, in a string, such as “1/250 seconds”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/exposuretimestring
func (c_ CSSearchableItemAttributeSet) SetExposureTimeString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setExposureTimeString:"), value)
}


// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fnumber
func (c_ CSSearchableItemAttributeSet) FNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("fNumber"))
	return rv
}


// The focal length of the lens, divided by the diameter of the aperture when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fnumber
func (c_ CSSearchableItemAttributeSet) SetFNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFNumber:"), value)
}


// The size of the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/filesize
func (c_ CSSearchableItemAttributeSet) FileSize() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("fileSize"))
	return rv
}


// The size of the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/filesize
func (c_ CSSearchableItemAttributeSet) SetFileSize(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFileSize:"), value)
}


// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/flashon
func (c_ CSSearchableItemAttributeSet) FlashOn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("flashOn"))
	return rv
}


// A value that indicates if the camera used a flash to capture the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/flashon
func (c_ CSSearchableItemAttributeSet) SetFlashOn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFlashOn:"), value)
}


// The actual focal length of the lens, in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength
func (c_ CSSearchableItemAttributeSet) FocalLength() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("focalLength"))
	return rv
}


// The actual focal length of the lens, in millimeters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength
func (c_ CSSearchableItemAttributeSet) SetFocalLength(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength:"), value)
}


// A value that indicates if the focal length is 35mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength35mm
func (c_ CSSearchableItemAttributeSet) FocalLength35mm() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("focalLength35mm"))
	return rv
}


// A value that indicates if the focal length is 35mm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/focallength35mm
func (c_ CSSearchableItemAttributeSet) SetFocalLength35mm(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFocalLength35mm:"), value)
}


// An array of font names the document uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fontnames
func (c_ CSSearchableItemAttributeSet) FontNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("fontNames"))
	return rv
}


// An array of font names the document uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fontnames
func (c_ CSSearchableItemAttributeSet) SetFontNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFontNames:"), value)
}


// The fully formatted address of the item, received from MapKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fullyformattedaddress
func (c_ CSSearchableItemAttributeSet) FullyFormattedAddress() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("fullyFormattedAddress"))
	return rv
}


// The fully formatted address of the item, received from MapKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/fullyformattedaddress
func (c_ CSSearchableItemAttributeSet) SetFullyFormattedAddress(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFullyFormattedAddress:"), value)
}


// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/generalmidisequence
func (c_ CSSearchableItemAttributeSet) GeneralMIDISequence() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("generalMIDISequence"))
	return rv
}


// A value that indicates whether the MIDI sequence the file contains is set up for use with a general MIDI device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/generalmidisequence
func (c_ CSSearchableItemAttributeSet) SetGeneralMIDISequence(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGeneralMIDISequence:"), value)
}


// The genre of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/genre
func (c_ CSSearchableItemAttributeSet) Genre() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("genre"))
	return rv
}


// The genre of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/genre
func (c_ CSSearchableItemAttributeSet) SetGenre(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGenre:"), value)
}


// Information about the GPS area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsareainformation
func (c_ CSSearchableItemAttributeSet) GpsAreaInformation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("gpsAreaInformation"))
	return rv
}


// Information about the GPS area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsareainformation
func (c_ CSSearchableItemAttributeSet) SetGpsAreaInformation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsAreaInformation:"), value)
}


// The date and time related to the GPS value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdatestamp
func (c_ CSSearchableItemAttributeSet) GpsDateStamp() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("gpsDateStamp"))
	return rv
}


// The date and time related to the GPS value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdatestamp
func (c_ CSSearchableItemAttributeSet) SetGpsDateStamp(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDateStamp:"), value)
}


// The bearing to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestbearing
func (c_ CSSearchableItemAttributeSet) GpsDestBearing() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("gpsDestBearing"))
	return rv
}


// The bearing to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestbearing
func (c_ CSSearchableItemAttributeSet) SetGpsDestBearing(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestBearing:"), value)
}


// The distance to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestdistance
func (c_ CSSearchableItemAttributeSet) GpsDestDistance() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("gpsDestDistance"))
	return rv
}


// The distance to the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestdistance
func (c_ CSSearchableItemAttributeSet) SetGpsDestDistance(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestDistance:"), value)
}


// The latitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlatitude
func (c_ CSSearchableItemAttributeSet) GpsDestLatitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("gpsDestLatitude"))
	return rv
}


// The latitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlatitude
func (c_ CSSearchableItemAttributeSet) SetGpsDestLatitude(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestLatitude:"), value)
}


// The longitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlongitude
func (c_ CSSearchableItemAttributeSet) GpsDestLongitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("gpsDestLongitude"))
	return rv
}


// The longitude of the destination point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdestlongitude
func (c_ CSSearchableItemAttributeSet) SetGpsDestLongitude(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDestLongitude:"), value)
}


// The differential correction applied to the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdifferental
func (c_ CSSearchableItemAttributeSet) GpsDifferental() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("gpsDifferental"))
	return rv
}


// The differential correction applied to the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsdifferental
func (c_ CSSearchableItemAttributeSet) SetGpsDifferental(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsDifferental:"), value)
}


// The geodetic data that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmapdatum
func (c_ CSSearchableItemAttributeSet) GpsMapDatum() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("gpsMapDatum"))
	return rv
}


// The geodetic data that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmapdatum
func (c_ CSSearchableItemAttributeSet) SetGpsMapDatum(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsMapDatum:"), value)
}


// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmeasuremode
func (c_ CSSearchableItemAttributeSet) GpsMeasureMode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("gpsMeasureMode"))
	return rv
}


// The measurement precision mode in use by the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsmeasuremode
func (c_ CSSearchableItemAttributeSet) SetGpsMeasureMode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsMeasureMode:"), value)
}


// The location finding method that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsprocessingmethod
func (c_ CSSearchableItemAttributeSet) GpsProcessingMethod() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("gpsProcessingMethod"))
	return rv
}


// The location finding method that the GPS receiver uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsprocessingmethod
func (c_ CSSearchableItemAttributeSet) SetGpsProcessingMethod(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsProcessingMethod:"), value)
}


// The status of the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsstatus
func (c_ CSSearchableItemAttributeSet) GpsStatus() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("gpsStatus"))
	return rv
}


// The status of the GPS receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpsstatus
func (c_ CSSearchableItemAttributeSet) SetGpsStatus(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsStatus:"), value)
}


// The direction of travel of the item in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpstrack
func (c_ CSSearchableItemAttributeSet) GpsTrack() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("gpsTrack"))
	return rv
}


// The direction of travel of the item in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/gpstrack
func (c_ CSSearchableItemAttributeSet) SetGpsTrack(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setGpsTrack:"), value)
}


// Indicates if the image file has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/hasalphachannel
func (c_ CSSearchableItemAttributeSet) HasAlphaChannel() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("hasAlphaChannel"))
	return rv
}


// Indicates if the image file has an alpha channel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/hasalphachannel
func (c_ CSSearchableItemAttributeSet) SetHasAlphaChannel(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasAlphaChannel:"), value)
}


// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/headline
func (c_ CSSearchableItemAttributeSet) Headline() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("headline"))
	return rv
}


// A publishable string that provides a synopsis of the contents of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/headline
func (c_ CSSearchableItemAttributeSet) SetHeadline(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHeadline:"), value)
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
func (c_ CSSearchableItemAttributeSet) HtmlContentData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("htmlContentData"))
	return rv
}


// The HTML content of the document encoded as an NSData object representing a UTF-8 encoded string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/htmlcontentdata
func (c_ CSSearchableItemAttributeSet) SetHtmlContentData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHtmlContentData:"), value)
}


// A formal identifier that references the document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/identifier
func (c_ CSSearchableItemAttributeSet) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}


// A formal identifier that references the document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/identifier
func (c_ CSSearchableItemAttributeSet) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), value)
}


// The direction of the item’s image in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/imagedirection
func (c_ CSSearchableItemAttributeSet) ImageDirection() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("imageDirection"))
	return rv
}


// The direction of the item’s image in degrees from true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/imagedirection
func (c_ CSSearchableItemAttributeSet) SetImageDirection(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageDirection:"), value)
}


// An array of important dates associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/importantdates
func (c_ CSSearchableItemAttributeSet) ImportantDates() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("importantDates"))
	return rv
}


// An array of important dates associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/importantdates
func (c_ CSSearchableItemAttributeSet) SetImportantDates(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImportantDates:"), value)
}


// Information about the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/information
func (c_ CSSearchableItemAttributeSet) Information() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("information"))
	return rv
}


// Information about the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/information
func (c_ CSSearchableItemAttributeSet) SetInformation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInformation:"), value)
}


// An array of instant message addresses for the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instantmessageaddresses
func (c_ CSSearchableItemAttributeSet) InstantMessageAddresses() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("instantMessageAddresses"))
	return rv
}


// An array of instant message addresses for the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instantmessageaddresses
func (c_ CSSearchableItemAttributeSet) SetInstantMessageAddresses(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstantMessageAddresses:"), value)
}


// Instructions that concern the use of the item, such as an embargo or warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instructions
func (c_ CSSearchableItemAttributeSet) Instructions() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("instructions"))
	return rv
}


// Instructions that concern the use of the item, such as an embargo or warning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/instructions
func (c_ CSSearchableItemAttributeSet) SetInstructions(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInstructions:"), value)
}


// A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ispriority
func (c_ CSSearchableItemAttributeSet) IsPriority() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("isPriority"))
	return rv
}


// A Boolean value that indicates whether the mail or messages content represents a prioritized item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ispriority
func (c_ CSSearchableItemAttributeSet) SetIsPriority(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPriority:"), value)
}


// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/isospeed
func (c_ CSSearchableItemAttributeSet) IsoSpeed() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("isoSpeed"))
	return rv
}


// The ISO speed setting at the time the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/isospeed
func (c_ CSSearchableItemAttributeSet) SetIsoSpeed(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsoSpeed:"), value)
}


// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keysignature
func (c_ CSSearchableItemAttributeSet) KeySignature() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("keySignature"))
	return rv
}


// The musical key of the song or audio composition that the file contains, such as C, Dm, or F#m.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keysignature
func (c_ CSSearchableItemAttributeSet) SetKeySignature(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeySignature:"), value)
}


// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keywords
func (c_ CSSearchableItemAttributeSet) Keywords() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("keywords"))
	return rv
}


// An array of keywords associated with the item, such as work, birthday, important, and so on.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/keywords
func (c_ CSSearchableItemAttributeSet) SetKeywords(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKeywords:"), value)
}


// A description of the kind of document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/kind
func (c_ CSSearchableItemAttributeSet) Kind() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("kind"))
	return rv
}


// A description of the kind of document the item represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/kind
func (c_ CSSearchableItemAttributeSet) SetKind(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setKind:"), value)
}


// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/languages
func (c_ CSSearchableItemAttributeSet) Languages() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("languages"))
	return rv
}


// A list of the included languages for the intellectual content of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/languages
func (c_ CSSearchableItemAttributeSet) SetLanguages(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLanguages:"), value)
}


// The date on which the file was last used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lastuseddate
func (c_ CSSearchableItemAttributeSet) LastUsedDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("lastUsedDate"))
	return rv
}


// The date on which the file was last used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lastuseddate
func (c_ CSSearchableItemAttributeSet) SetLastUsedDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLastUsedDate:"), value)
}


// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/latitude
func (c_ CSSearchableItemAttributeSet) Latitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("latitude"))
	return rv
}


// The latitude of the item, in degrees north of the equator, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/latitude
func (c_ CSSearchableItemAttributeSet) SetLatitude(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLatitude:"), value)
}


// An array that contains the names of the various layers in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/layernames
func (c_ CSSearchableItemAttributeSet) LayerNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("layerNames"))
	return rv
}


// An array that contains the names of the various layers in the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/layernames
func (c_ CSSearchableItemAttributeSet) SetLayerNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLayerNames:"), value)
}


// The model of the lens that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lensmodel
func (c_ CSSearchableItemAttributeSet) LensModel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("lensModel"))
	return rv
}


// The model of the lens that captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lensmodel
func (c_ CSSearchableItemAttributeSet) SetLensModel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLensModel:"), value)
}


// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/likelyjunk
func (c_ CSSearchableItemAttributeSet) LikelyJunk() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("likelyJunk"))
	return rv
}


// A value that indicates if the message is likely to be considered junk.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/likelyjunk
func (c_ CSSearchableItemAttributeSet) SetLikelyJunk(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLikelyJunk:"), value)
}


// A value that indicates if the media is local.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/local
func (c_ CSSearchableItemAttributeSet) Local() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("local"))
	return rv
}


// A value that indicates if the media is local.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/local
func (c_ CSSearchableItemAttributeSet) SetLocal(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocal:"), value)
}


// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/longitude
func (c_ CSSearchableItemAttributeSet) Longitude() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("longitude"))
	return rv
}


// The longitude of the item, in degrees east of the prime meridian, expressed using the WGS84 datum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/longitude
func (c_ CSSearchableItemAttributeSet) SetLongitude(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLongitude:"), value)
}


// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lyricist
func (c_ CSSearchableItemAttributeSet) Lyricist() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("lyricist"))
	return rv
}


// The lyricist or text writer for the song or audio composition that the file contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/lyricist
func (c_ CSSearchableItemAttributeSet) SetLyricist(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLyricist:"), value)
}


// An array of mailbox identifiers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mailboxidentifiers
func (c_ CSSearchableItemAttributeSet) MailboxIdentifiers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("mailboxIdentifiers"))
	return rv
}


// An array of mailbox identifiers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mailboxidentifiers
func (c_ CSSearchableItemAttributeSet) SetMailboxIdentifiers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMailboxIdentifiers:"), value)
}


// The smallest F number of the lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/maxaperture
func (c_ CSSearchableItemAttributeSet) MaxAperture() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("maxAperture"))
	return rv
}


// The smallest F number of the lens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/maxaperture
func (c_ CSSearchableItemAttributeSet) SetMaxAperture(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxAperture:"), value)
}


// The media types present in the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mediatypes
func (c_ CSSearchableItemAttributeSet) MediaTypes() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("mediaTypes"))
	return rv
}


// The media types present in the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/mediatypes
func (c_ CSSearchableItemAttributeSet) SetMediaTypes(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMediaTypes:"), value)
}


// The date on which the last metadata attribute was changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/metadatamodificationdate
func (c_ CSSearchableItemAttributeSet) MetadataModificationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("metadataModificationDate"))
	return rv
}


// The date on which the last metadata attribute was changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/metadatamodificationdate
func (c_ CSSearchableItemAttributeSet) SetMetadataModificationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadataModificationDate:"), value)
}


// The metering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/meteringmode
func (c_ CSSearchableItemAttributeSet) MeteringMode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("meteringMode"))
	return rv
}


// The metering mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/meteringmode
func (c_ CSSearchableItemAttributeSet) SetMeteringMode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMeteringMode:"), value)
}


// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalgenre
func (c_ CSSearchableItemAttributeSet) MusicalGenre() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("musicalGenre"))
	return rv
}


// The musical genre of the song or audio composition that the file contains, such as jazz, pop, rock, or classical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalgenre
func (c_ CSSearchableItemAttributeSet) SetMusicalGenre(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalGenre:"), value)
}


// The category of the instrument associated with the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentcategory
func (c_ CSSearchableItemAttributeSet) MusicalInstrumentCategory() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("musicalInstrumentCategory"))
	return rv
}


// The category of the instrument associated with the audio file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentcategory
func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentCategory:"), value)
}


// The name of an instrument within the context of an instrument category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentname
func (c_ CSSearchableItemAttributeSet) MusicalInstrumentName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("musicalInstrumentName"))
	return rv
}


// The name of an instrument within the context of an instrument category.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/musicalinstrumentname
func (c_ CSSearchableItemAttributeSet) SetMusicalInstrumentName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMusicalInstrumentName:"), value)
}


// The name of the location or point of interest associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/namedlocation
func (c_ CSSearchableItemAttributeSet) NamedLocation() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("namedLocation"))
	return rv
}


// The name of the location or point of interest associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/namedlocation
func (c_ CSSearchableItemAttributeSet) SetNamedLocation(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setNamedLocation:"), value)
}


// A list of companies or organizations that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/organizations
func (c_ CSSearchableItemAttributeSet) Organizations() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("organizations"))
	return rv
}


// A list of companies or organizations that created the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/organizations
func (c_ CSSearchableItemAttributeSet) SetOrganizations(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrganizations:"), value)
}


// The orientation of the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/orientation
func (c_ CSSearchableItemAttributeSet) Orientation() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("orientation"))
	return rv
}


// The orientation of the data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/orientation
func (c_ CSSearchableItemAttributeSet) SetOrientation(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOrientation:"), value)
}


// The original format of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalformat
func (c_ CSSearchableItemAttributeSet) OriginalFormat() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("originalFormat"))
	return rv
}


// The original format of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalformat
func (c_ CSSearchableItemAttributeSet) SetOriginalFormat(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalFormat:"), value)
}


// The original source of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalsource
func (c_ CSSearchableItemAttributeSet) OriginalSource() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("originalSource"))
	return rv
}


// The original source of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/originalsource
func (c_ CSSearchableItemAttributeSet) SetOriginalSource(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOriginalSource:"), value)
}


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagecount
func (c_ CSSearchableItemAttributeSet) PageCount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pageCount"))
	return rv
}


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagecount
func (c_ CSSearchableItemAttributeSet) SetPageCount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageCount:"), value)
}


// The height of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pageheight
func (c_ CSSearchableItemAttributeSet) PageHeight() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pageHeight"))
	return rv
}


// The height of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pageheight
func (c_ CSSearchableItemAttributeSet) SetPageHeight(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageHeight:"), value)
}


// The width of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagewidth
func (c_ CSSearchableItemAttributeSet) PageWidth() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pageWidth"))
	return rv
}


// The width of the document page, in points (72 points per inch).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pagewidth
func (c_ CSSearchableItemAttributeSet) SetPageWidth(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPageWidth:"), value)
}


// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/participants
func (c_ CSSearchableItemAttributeSet) Participants() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("participants"))
	return rv
}


// A list of people who are visible in an image or movie or written about in a document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/participants
func (c_ CSSearchableItemAttributeSet) SetParticipants(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setParticipants:"), value)
}


// The complete path to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/path
func (c_ CSSearchableItemAttributeSet) Path() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("path"))
	return rv
}


// The complete path to the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/path
func (c_ CSSearchableItemAttributeSet) SetPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPath:"), value)
}


// A list of performers in the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/performers
func (c_ CSSearchableItemAttributeSet) Performers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("performers"))
	return rv
}


// A list of performers in the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/performers
func (c_ CSSearchableItemAttributeSet) SetPerformers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPerformers:"), value)
}


// An array of phone numbers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/phonenumbers
func (c_ CSSearchableItemAttributeSet) PhoneNumbers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("phoneNumbers"))
	return rv
}


// An array of phone numbers associated with the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/phonenumbers
func (c_ CSSearchableItemAttributeSet) SetPhoneNumbers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPhoneNumbers:"), value)
}


// The total number of pixels in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelcount
func (c_ CSSearchableItemAttributeSet) PixelCount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pixelCount"))
	return rv
}


// The total number of pixels in the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelcount
func (c_ CSSearchableItemAttributeSet) SetPixelCount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelCount:"), value)
}


// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelheight
func (c_ CSSearchableItemAttributeSet) PixelHeight() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pixelHeight"))
	return rv
}


// The height of the item, such as image or video frame height, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelheight
func (c_ CSSearchableItemAttributeSet) SetPixelHeight(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelHeight:"), value)
}


// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelwidth
func (c_ CSSearchableItemAttributeSet) PixelWidth() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("pixelWidth"))
	return rv
}


// The width of the item, such as image or video frame width, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/pixelwidth
func (c_ CSSearchableItemAttributeSet) SetPixelWidth(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPixelWidth:"), value)
}


// A user-supplied play count for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/playcount
func (c_ CSSearchableItemAttributeSet) PlayCount() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("playCount"))
	return rv
}


// A user-supplied play count for the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/playcount
func (c_ CSSearchableItemAttributeSet) SetPlayCount(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPlayCount:"), value)
}


// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/postalcode
func (c_ CSSearchableItemAttributeSet) PostalCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("postalCode"))
	return rv
}


// The postal code for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/postalcode
func (c_ CSSearchableItemAttributeSet) SetPostalCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPostalCode:"), value)
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
func (c_ CSSearchableItemAttributeSet) Producer() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("producer"))
	return rv
}


// The producer of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/producer
func (c_ CSSearchableItemAttributeSet) SetProducer(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProducer:"), value)
}


// The name of the color profile the camera used for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/profilename
func (c_ CSSearchableItemAttributeSet) ProfileName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("profileName"))
	return rv
}


// The name of the color profile the camera used for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/profilename
func (c_ CSSearchableItemAttributeSet) SetProfileName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProfileName:"), value)
}


// A list of projects of which this file is a part.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/projects
func (c_ CSSearchableItemAttributeSet) Projects() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("projects"))
	return rv
}


// A list of projects of which this file is a part.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/projects
func (c_ CSSearchableItemAttributeSet) SetProjects(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProjects:"), value)
}


// An array of identifiers that corresponds to data representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerdatatypeidentifiers
func (c_ CSSearchableItemAttributeSet) ProviderDataTypeIdentifiers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("providerDataTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to data representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerdatatypeidentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderDataTypeIdentifiers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderDataTypeIdentifiers:"), value)
}


// An array of identifiers that corresponds to file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerfiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) ProviderFileTypeIdentifiers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("providerFileTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerfiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderFileTypeIdentifiers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderFileTypeIdentifiers:"), value)
}


// An array of identifiers that corresponds to in-place file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerinplacefiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) ProviderInPlaceFileTypeIdentifiers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("providerInPlaceFileTypeIdentifiers"))
	return rv
}


// An array of identifiers that corresponds to in-place file representations the delegate provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/providerinplacefiletypeidentifiers
func (c_ CSSearchableItemAttributeSet) SetProviderInPlaceFileTypeIdentifiers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setProviderInPlaceFileTypeIdentifiers:"), value)
}


// A list of people, organizations, services, or other entities responsible for making the media available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/publishers
func (c_ CSSearchableItemAttributeSet) Publishers() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("publishers"))
	return rv
}


// A list of people, organizations, services, or other entities responsible for making the media available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/publishers
func (c_ CSSearchableItemAttributeSet) SetPublishers(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPublishers:"), value)
}


// A number that indicates the relative importance of the item among other items from the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rankinghint
func (c_ CSSearchableItemAttributeSet) RankingHint() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("rankingHint"))
	return rv
}


// A number that indicates the relative importance of the item among other items from the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rankinghint
func (c_ CSSearchableItemAttributeSet) SetRankingHint(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRankingHint:"), value)
}


// The user-supplied rating of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rating
func (c_ CSSearchableItemAttributeSet) Rating() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("rating"))
	return rv
}


// The user-supplied rating of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rating
func (c_ CSSearchableItemAttributeSet) SetRating(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRating:"), value)
}


// A description of the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ratingdescription
func (c_ CSSearchableItemAttributeSet) RatingDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("ratingDescription"))
	return rv
}


// A description of the rating.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/ratingdescription
func (c_ CSSearchableItemAttributeSet) SetRatingDescription(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRatingDescription:"), value)
}


// An array of addresses associated with the recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientaddresses
func (c_ CSSearchableItemAttributeSet) RecipientAddresses() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("recipientAddresses"))
	return rv
}


// An array of addresses associated with the recipients of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientaddresses
func (c_ CSSearchableItemAttributeSet) SetRecipientAddresses(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientAddresses:"), value)
}


// An array of email addresses associated with the recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientemailaddresses
func (c_ CSSearchableItemAttributeSet) RecipientEmailAddresses() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("recipientEmailAddresses"))
	return rv
}


// An array of email addresses associated with the recipient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientemailaddresses
func (c_ CSSearchableItemAttributeSet) SetRecipientEmailAddresses(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientEmailAddresses:"), value)
}


// An array of names representing the recipients of this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientnames
func (c_ CSSearchableItemAttributeSet) RecipientNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("recipientNames"))
	return rv
}


// An array of names representing the recipients of this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recipientnames
func (c_ CSSearchableItemAttributeSet) SetRecipientNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecipientNames:"), value)
}


// The recording date of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recordingdate
func (c_ CSSearchableItemAttributeSet) RecordingDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("recordingDate"))
	return rv
}


// The recording date of the song or composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/recordingdate
func (c_ CSSearchableItemAttributeSet) SetRecordingDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRecordingDate:"), value)
}


// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/redeyeon
func (c_ CSSearchableItemAttributeSet) RedEyeOn() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("redEyeOn"))
	return rv
}


// A value that indicates if the camera used red-eye reduction when capturing the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/redeyeon
func (c_ CSSearchableItemAttributeSet) SetRedEyeOn(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRedEyeOn:"), value)
}


// The unique identifier for the item to which the activity is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/relateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) RelatedUniqueIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("relatedUniqueIdentifier"))
	return rv
}


// The unique identifier for the item to which the activity is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/relateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) SetRelatedUniqueIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRelatedUniqueIdentifier:"), value)
}


// The resolution height of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionheightdpi
func (c_ CSSearchableItemAttributeSet) ResolutionHeightDPI() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("resolutionHeightDPI"))
	return rv
}


// The resolution height of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionheightdpi
func (c_ CSSearchableItemAttributeSet) SetResolutionHeightDPI(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionHeightDPI:"), value)
}


// The resolution width of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionwidthdpi
func (c_ CSSearchableItemAttributeSet) ResolutionWidthDPI() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("resolutionWidthDPI"))
	return rv
}


// The resolution width of the image, in DPI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/resolutionwidthdpi
func (c_ CSSearchableItemAttributeSet) SetResolutionWidthDPI(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setResolutionWidthDPI:"), value)
}


// A link to information about the rights held in and over the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rights
func (c_ CSSearchableItemAttributeSet) Rights() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("rights"))
	return rv
}


// A link to information about the rights held in and over the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/rights
func (c_ CSSearchableItemAttributeSet) SetRights(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRights:"), value)
}


// Indicates the role of the content creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/role
func (c_ CSSearchableItemAttributeSet) Role() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("role"))
	return rv
}


// Indicates the role of the content creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/role
func (c_ CSSearchableItemAttributeSet) SetRole(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setRole:"), value)
}


// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/securitymethod
func (c_ CSSearchableItemAttributeSet) SecurityMethod() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("securityMethod"))
	return rv
}


// The security method (a type of encryption) that protects the document file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/securitymethod
func (c_ CSSearchableItemAttributeSet) SetSecurityMethod(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSecurityMethod:"), value)
}


// The file type of the item to enable the user to share items from Spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/shareditemcontenttype
func (c_ CSSearchableItemAttributeSet) SharedItemContentType() objc.IObject /* cross-framework: UTType */ {
	rv := objc.Send[uniformtypeidentifiers.UTType](c_.ID, objc.Sel("sharedItemContentType"))
	return rv
}


// The file type of the item to enable the user to share items from Spotlight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/shareditemcontenttype
func (c_ CSSearchableItemAttributeSet) SetSharedItemContentType(value objc.IObject /* cross-framework: UTType */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSharedItemContentType:"), value)
}


// The speed of the item, in kilometers per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/speed
func (c_ CSSearchableItemAttributeSet) Speed() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("speed"))
	return rv
}


// The speed of the item, in kilometers per hour.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/speed
func (c_ CSSearchableItemAttributeSet) SetSpeed(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpeed:"), value)
}


// The start date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/startdate
func (c_ CSSearchableItemAttributeSet) StartDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("startDate"))
	return rv
}


// The start date for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/startdate
func (c_ CSSearchableItemAttributeSet) SetStartDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStartDate:"), value)
}


// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/stateorprovince
func (c_ CSSearchableItemAttributeSet) StateOrProvince() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("stateOrProvince"))
	return rv
}


// The province or state of origin according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/stateorprovince
func (c_ CSSearchableItemAttributeSet) SetStateOrProvince(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStateOrProvince:"), value)
}


// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/streamable
func (c_ CSSearchableItemAttributeSet) Streamable() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("streamable"))
	return rv
}


// A value that indicates if the content is prepared for streaming.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/streamable
func (c_ CSSearchableItemAttributeSet) SetStreamable(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStreamable:"), value)
}


// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subthoroughfare
func (c_ CSSearchableItemAttributeSet) SubThoroughfare() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subThoroughfare"))
	return rv
}


// The sublocation, such as a street number, for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subthoroughfare
func (c_ CSSearchableItemAttributeSet) SetSubThoroughfare(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubThoroughfare:"), value)
}


// The subject of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subject
func (c_ CSSearchableItemAttributeSet) Subject() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("subject"))
	return rv
}


// The subject of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/subject
func (c_ CSSearchableItemAttributeSet) SetSubject(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSubject:"), value)
}


// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/supportsphonecall
func (c_ CSSearchableItemAttributeSet) SupportsPhoneCall() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("supportsPhoneCall"))
	return rv
}


// A value that indicates whether the item contains information sufficient to allow a phone call to a number associated with the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/supportsphonecall
func (c_ CSSearchableItemAttributeSet) SetSupportsPhoneCall(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupportsPhoneCall:"), value)
}


// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/tempo
func (c_ CSSearchableItemAttributeSet) Tempo() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("tempo"))
	return rv
}


// The tempo of the music that the audio file contains, in beats per minute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/tempo
func (c_ CSSearchableItemAttributeSet) SetTempo(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTempo:"), value)
}


// The textual content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontent
func (c_ CSSearchableItemAttributeSet) TextContent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("textContent"))
	return rv
}


// The textual content of the message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontent
func (c_ CSSearchableItemAttributeSet) SetTextContent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextContent:"), value)
}


// A string that presents the Apple Intelligence summarization of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontentsummary
func (c_ CSSearchableItemAttributeSet) TextContentSummary() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("textContentSummary"))
	return rv
}


// A string that presents the Apple Intelligence summarization of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/textcontentsummary
func (c_ CSSearchableItemAttributeSet) SetTextContentSummary(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTextContentSummary:"), value)
}


// The theme of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/theme
func (c_ CSSearchableItemAttributeSet) Theme() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("theme"))
	return rv
}


// The theme of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/theme
func (c_ CSSearchableItemAttributeSet) SetTheme(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTheme:"), value)
}


// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thoroughfare
func (c_ CSSearchableItemAttributeSet) Thoroughfare() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("thoroughfare"))
	return rv
}


// The thoroughfare, such as a street name, associated with the location for the item according to guidelines the provider establishes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thoroughfare
func (c_ CSSearchableItemAttributeSet) SetThoroughfare(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThoroughfare:"), value)
}


// Image data that represents the thumbnail of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnaildata
func (c_ CSSearchableItemAttributeSet) ThumbnailData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("thumbnailData"))
	return rv
}


// Image data that represents the thumbnail of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnaildata
func (c_ CSSearchableItemAttributeSet) SetThumbnailData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailData:"), value)
}


// The local file URL of the thumbnail image for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnailurl
func (c_ CSSearchableItemAttributeSet) ThumbnailURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("thumbnailURL"))
	return rv
}


// The local file URL of the thumbnail image for the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/thumbnailurl
func (c_ CSSearchableItemAttributeSet) SetThumbnailURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setThumbnailURL:"), value)
}


// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timesignature
func (c_ CSSearchableItemAttributeSet) TimeSignature() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("timeSignature"))
	return rv
}


// The time signature of the musical composition that the audio or MIDI file contains, in a string, such as “4/4” or “7/8”.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timesignature
func (c_ CSSearchableItemAttributeSet) SetTimeSignature(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimeSignature:"), value)
}


// The timestamp on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timestamp
func (c_ CSSearchableItemAttributeSet) Timestamp() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](c_.ID, objc.Sel("timestamp"))
	return rv
}


// The timestamp on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/timestamp
func (c_ CSSearchableItemAttributeSet) SetTimestamp(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTimestamp:"), value)
}


// The total bit rate of the media, combining audio and video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/totalbitrate
func (c_ CSSearchableItemAttributeSet) TotalBitRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("totalBitRate"))
	return rv
}


// The total bit rate of the media, combining audio and video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/totalbitrate
func (c_ CSSearchableItemAttributeSet) SetTotalBitRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTotalBitRate:"), value)
}


// A string that represents the text the system transcribed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/transcribedtextcontent
func (c_ CSSearchableItemAttributeSet) TranscribedTextContent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("transcribedTextContent"))
	return rv
}


// A string that represents the text the system transcribed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/transcribedtextcontent
func (c_ CSSearchableItemAttributeSet) SetTranscribedTextContent(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTranscribedTextContent:"), value)
}


// The URL associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/url
func (c_ CSSearchableItemAttributeSet) Url() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](c_.ID, objc.Sel("url"))
	return rv
}


// The URL associated with the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/url
func (c_ CSSearchableItemAttributeSet) SetUrl(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUrl:"), value)
}


// A value that indicates the user created the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercreated
func (c_ CSSearchableItemAttributeSet) UserCreated() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("userCreated"))
	return rv
}


// A value that indicates the user created the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercreated
func (c_ CSSearchableItemAttributeSet) SetUserCreated(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCreated:"), value)
}


// A value that indicates the user selected the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercurated
func (c_ CSSearchableItemAttributeSet) UserCurated() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("userCurated"))
	return rv
}


// A value that indicates the user selected the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/usercurated
func (c_ CSSearchableItemAttributeSet) SetUserCurated(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserCurated:"), value)
}


// A value that indicates the user purchased or owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/userowned
func (c_ CSSearchableItemAttributeSet) UserOwned() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("userOwned"))
	return rv
}


// A value that indicates the user purchased or owns the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/userowned
func (c_ CSSearchableItemAttributeSet) SetUserOwned(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setUserOwned:"), value)
}


// A version string associated with the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/version
func (c_ CSSearchableItemAttributeSet) Version() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("version"))
	return rv
}


// A version string associated with the file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/version
func (c_ CSSearchableItemAttributeSet) SetVersion(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVersion:"), value)
}


// The video bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/videobitrate
func (c_ CSSearchableItemAttributeSet) VideoBitRate() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("videoBitRate"))
	return rv
}


// The video bit rate of the media.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/videobitrate
func (c_ CSSearchableItemAttributeSet) SetVideoBitRate(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setVideoBitRate:"), value)
}


// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/weakrelateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) WeakRelatedUniqueIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("weakRelatedUniqueIdentifier"))
	return rv
}


// The unique identifier for the item to which the activity is related, but not linked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/weakrelateduniqueidentifier
func (c_ CSSearchableItemAttributeSet) SetWeakRelatedUniqueIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWeakRelatedUniqueIdentifier:"), value)
}


// The white balance setting when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/whitebalance
func (c_ CSSearchableItemAttributeSet) WhiteBalance() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](c_.ID, objc.Sel("whiteBalance"))
	return rv
}


// The white balance setting when the camera captured the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/corespotlight/cssearchableitemattributeset/whitebalance
func (c_ CSSearchableItemAttributeSet) SetWhiteBalance(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setWhiteBalance:"), value)
}



