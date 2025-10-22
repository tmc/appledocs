// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MetadataItem] class.
var (
	MetadataItemClass     _MetadataItemClass
	MetadataItemClassOnce sync.Once
)

func getMetadataItemClass() _MetadataItemClass {
	MetadataItemClassOnce.Do(func() {
		MetadataItemClass = _MetadataItemClass{objc.GetClass("NSMetadataItem")}
	})
	return MetadataItemClass
}

type _MetadataItemClass struct {
	class objc.Class
}

// An interface definition for the [MetadataItem] class.
type IMetadataItem interface {
	objectivec.IObject
	Attributes() string
	SetAttributes(value string)
	NSMetadataItemAcquisitionMakeKey() string
	NSMetadataItemAcquisitionModelKey() string
	NSMetadataItemAlbumKey() string
	NSMetadataItemAltitudeKey() string
	NSMetadataItemApertureKey() string
	NSMetadataItemAppleLoopDescriptorsKey() string
	NSMetadataItemAppleLoopsKeyFilterTypeKey() string
	NSMetadataItemAppleLoopsLoopModeKey() string
	NSMetadataItemAppleLoopsRootKeyKey() string
	NSMetadataItemApplicationCategoriesKey() string
	NSMetadataItemAttributeChangeDateKey() string
	NSMetadataItemAudiencesKey() string
	NSMetadataItemAudioBitRateKey() string
	NSMetadataItemAudioChannelCountKey() string
	NSMetadataItemAudioEncodingApplicationKey() string
	NSMetadataItemAudioSampleRateKey() string
	NSMetadataItemAudioTrackNumberKey() string
	NSMetadataItemAuthorAddressesKey() string
	NSMetadataItemAuthorEmailAddressesKey() string
	NSMetadataItemAuthorsKey() string
	NSMetadataItemBitsPerSampleKey() string
	NSMetadataItemCFBundleIdentifierKey() string
	NSMetadataItemCameraOwnerKey() string
	NSMetadataItemCityKey() string
	NSMetadataItemCodecsKey() string
	NSMetadataItemColorSpaceKey() string
	NSMetadataItemCommentKey() string
	NSMetadataItemComposerKey() string
	NSMetadataItemContactKeywordsKey() string
	NSMetadataItemContentCreationDateKey() string
	NSMetadataItemContentModificationDateKey() string
	NSMetadataItemContentTypeKey() string
	NSMetadataItemContentTypeTreeKey() string
	NSMetadataItemContributorsKey() string
	NSMetadataItemCopyrightKey() string
	NSMetadataItemCountryKey() string
	NSMetadataItemCoverageKey() string
	NSMetadataItemCreatorKey() string
	NSMetadataItemDateAddedKey() string
	NSMetadataItemDeliveryTypeKey() string
	NSMetadataItemDescriptionKey() string
	NSMetadataItemDirectorKey() string
	NSMetadataItemDisplayNameKey() string
	NSMetadataItemDownloadedDateKey() string
	NSMetadataItemDueDateKey() string
	NSMetadataItemDurationSecondsKey() string
	NSMetadataItemEXIFGPSVersionKey() string
	NSMetadataItemEXIFVersionKey() string
	NSMetadataItemEditorsKey() string
	NSMetadataItemEmailAddressesKey() string
	NSMetadataItemEncodingApplicationsKey() string
	NSMetadataItemExecutableArchitecturesKey() string
	NSMetadataItemExecutablePlatformKey() string
	NSMetadataItemExposureModeKey() string
	NSMetadataItemExposureProgramKey() string
	NSMetadataItemExposureTimeSecondsKey() string
	NSMetadataItemExposureTimeStringKey() string
	NSMetadataItemFNumberKey() string
	NSMetadataItemFSContentChangeDateKey() string
	NSMetadataItemFSCreationDateKey() string
	NSMetadataItemFSNameKey() string
	NSMetadataItemFSSizeKey() string
	NSMetadataItemFinderCommentKey() string
	NSMetadataItemFlashOnOffKey() string
	NSMetadataItemFocalLength35mmKey() string
	NSMetadataItemFocalLengthKey() string
	NSMetadataItemFontsKey() string
	NSMetadataItemGPSAreaInformationKey() string
	NSMetadataItemGPSDOPKey() string
	NSMetadataItemGPSDateStampKey() string
	NSMetadataItemGPSDestBearingKey() string
	NSMetadataItemGPSDestDistanceKey() string
	NSMetadataItemGPSDestLatitudeKey() string
	NSMetadataItemGPSDestLongitudeKey() string
	NSMetadataItemGPSDifferentalKey() string
	NSMetadataItemGPSMapDatumKey() string
	NSMetadataItemGPSMeasureModeKey() string
	NSMetadataItemGPSProcessingMethodKey() string
	NSMetadataItemGPSStatusKey() string
	NSMetadataItemGPSTrackKey() string
	NSMetadataItemGenreKey() string
	NSMetadataItemHasAlphaChannelKey() string
	NSMetadataItemHeadlineKey() string
	NSMetadataItemISOSpeedKey() string
	NSMetadataItemIdentifierKey() string
	NSMetadataItemImageDirectionKey() string
	NSMetadataItemInformationKey() string
	NSMetadataItemInstantMessageAddressesKey() string
	NSMetadataItemInstructionsKey() string
	NSMetadataItemIsApplicationManagedKey() string
	NSMetadataItemIsGeneralMIDISequenceKey() string
	NSMetadataItemIsLikelyJunkKey() string
	NSMetadataItemIsUbiquitousKey() string
	NSMetadataItemKeySignatureKey() string
	NSMetadataItemKeywordsKey() string
	NSMetadataItemKindKey() string
	NSMetadataItemLanguagesKey() string
	NSMetadataItemLastUsedDateKey() string
	NSMetadataItemLatitudeKey() string
	NSMetadataItemLayerNamesKey() string
	NSMetadataItemLensModelKey() string
	NSMetadataItemLongitudeKey() string
	NSMetadataItemLyricistKey() string
	NSMetadataItemMaxApertureKey() string
	NSMetadataItemMediaTypesKey() string
	NSMetadataItemMeteringModeKey() string
	NSMetadataItemMusicalGenreKey() string
	NSMetadataItemMusicalInstrumentCategoryKey() string
	NSMetadataItemMusicalInstrumentNameKey() string
	NSMetadataItemNamedLocationKey() string
	NSMetadataItemNumberOfPagesKey() string
	NSMetadataItemOrganizationsKey() string
	NSMetadataItemOrientationKey() string
	NSMetadataItemOriginalFormatKey() string
	NSMetadataItemOriginalSourceKey() string
	NSMetadataItemPageHeightKey() string
	NSMetadataItemPageWidthKey() string
	NSMetadataItemParticipantsKey() string
	NSMetadataItemPathKey() string
	NSMetadataItemPerformersKey() string
	NSMetadataItemPhoneNumbersKey() string
	NSMetadataItemPixelCountKey() string
	NSMetadataItemPixelHeightKey() string
	NSMetadataItemPixelWidthKey() string
	NSMetadataItemProducerKey() string
	NSMetadataItemProfileNameKey() string
	NSMetadataItemProjectsKey() string
	NSMetadataItemPublishersKey() string
	NSMetadataItemRecipientAddressesKey() string
	NSMetadataItemRecipientEmailAddressesKey() string
	NSMetadataItemRecipientsKey() string
	NSMetadataItemRecordingDateKey() string
	NSMetadataItemRecordingYearKey() string
	NSMetadataItemRedEyeOnOffKey() string
	NSMetadataItemResolutionHeightDPIKey() string
	NSMetadataItemResolutionWidthDPIKey() string
	NSMetadataItemRightsKey() string
	NSMetadataItemSecurityMethodKey() string
	NSMetadataItemSpeedKey() string
	NSMetadataItemStarRatingKey() string
	NSMetadataItemStateOrProvinceKey() string
	NSMetadataItemStreamableKey() string
	NSMetadataItemSubjectKey() string
	NSMetadataItemTempoKey() string
	NSMetadataItemTextContentKey() string
	NSMetadataItemThemeKey() string
	NSMetadataItemTimeSignatureKey() string
	NSMetadataItemTimestampKey() string
	NSMetadataItemTitleKey() string
	NSMetadataItemTotalBitRateKey() string
	NSMetadataItemURLKey() string
	NSMetadataItemVersionKey() string
	NSMetadataItemVideoBitRateKey() string
	NSMetadataItemWhereFromsKey() string
	NSMetadataItemWhiteBalanceKey() string
	NSMetadataUbiquitousItemContainerDisplayNameKey() string
	NSMetadataUbiquitousItemDownloadRequestedKey() string
	NSMetadataUbiquitousItemDownloadingErrorKey() string
	NSMetadataUbiquitousItemDownloadingStatusCurrent() string
	NSMetadataUbiquitousItemDownloadingStatusDownloaded() string
	NSMetadataUbiquitousItemDownloadingStatusKey() string
	NSMetadataUbiquitousItemDownloadingStatusNotDownloaded() string
	NSMetadataUbiquitousItemHasUnresolvedConflictsKey() string
	NSMetadataUbiquitousItemIsDownloadedKey() string
	NSMetadataUbiquitousItemIsDownloadingKey() string
	NSMetadataUbiquitousItemIsExternalDocumentKey() string
	NSMetadataUbiquitousItemIsSharedKey() string
	NSMetadataUbiquitousItemIsUploadedKey() string
	NSMetadataUbiquitousItemIsUploadingKey() string
	NSMetadataUbiquitousItemPercentDownloadedKey() string
	NSMetadataUbiquitousItemPercentUploadedKey() string
	NSMetadataUbiquitousItemURLInLocalContainerKey() string
	NSMetadataUbiquitousItemUploadingErrorKey() string
	NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey() string
	NSMetadataUbiquitousSharedItemCurrentUserRoleKey() string
	NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey() string
	NSMetadataUbiquitousSharedItemOwnerNameComponentsKey() string
	NSMetadataUbiquitousSharedItemPermissionsReadOnly() string
	NSMetadataUbiquitousSharedItemPermissionsReadWrite() string
	NSMetadataUbiquitousSharedItemRoleOwner() string
	NSMetadataUbiquitousSharedItemRoleParticipant() string
}

// The metadata associated with a file.
//
// Metadata items provide a simple interface to retrieve the available attribute names and values.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMetadataItem
type MetadataItem struct {
	objectivec.Object
}

// MetadataItemFrom constructs a [MetadataItem] from an unsafe.Pointer.
//
// The metadata associated with a file.
func MetadataItemFrom(ptr unsafe.Pointer) MetadataItem {
	return MetadataItem{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MetadataItemClass) Alloc() MetadataItem {
	rv := objc.Send[MetadataItem](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MetadataItemClass) New() MetadataItem {
	rv := objc.Send[MetadataItem](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MetadataItem) Init() MetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MetadataItem) Autorelease() MetadataItem {
	rv := objc.Send[MetadataItem](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetadataItem creates a new MetadataItem instance.
func NewMetadataItem() MetadataItem {
	return getMetadataItemClass().New()
}


// An array containing the attribute keys for the metadata item’s values.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitem/attributes
func (m_ MetadataItem) Attributes() string {
	rv := objc.Send[string](m_.ID, objc.Sel("attributes"))
	return rv
}


// SetAttributes sets the value of the attributes property.
// An array containing the attribute keys for the metadata item’s values.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitem/attributes
func (m_ MetadataItem) SetAttributes(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributes:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemacquisitionmakekey
func (m_ MetadataItem) NSMetadataItemAcquisitionMakeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAcquisitionMakeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemacquisitionmodelkey
func (m_ MetadataItem) NSMetadataItemAcquisitionModelKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAcquisitionModelKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemalbumkey
func (m_ MetadataItem) NSMetadataItemAlbumKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAlbumKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaltitudekey
func (m_ MetadataItem) NSMetadataItemAltitudeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAltitudeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaperturekey
func (m_ MetadataItem) NSMetadataItemApertureKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemApertureKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopdescriptorskey
func (m_ MetadataItem) NSMetadataItemAppleLoopDescriptorsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAppleLoopDescriptorsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopskeyfiltertypekey
func (m_ MetadataItem) NSMetadataItemAppleLoopsKeyFilterTypeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAppleLoopsKeyFilterTypeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopsloopmodekey
func (m_ MetadataItem) NSMetadataItemAppleLoopsLoopModeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAppleLoopsLoopModeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopsrootkeykey
func (m_ MetadataItem) NSMetadataItemAppleLoopsRootKeyKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAppleLoopsRootKeyKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemapplicationcategorieskey
func (m_ MetadataItem) NSMetadataItemApplicationCategoriesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemApplicationCategoriesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemattributechangedatekey
func (m_ MetadataItem) NSMetadataItemAttributeChangeDateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAttributeChangeDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudienceskey
func (m_ MetadataItem) NSMetadataItemAudiencesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAudiencesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiobitratekey
func (m_ MetadataItem) NSMetadataItemAudioBitRateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAudioBitRateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiochannelcountkey
func (m_ MetadataItem) NSMetadataItemAudioChannelCountKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAudioChannelCountKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudioencodingapplicationkey
func (m_ MetadataItem) NSMetadataItemAudioEncodingApplicationKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAudioEncodingApplicationKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiosampleratekey
func (m_ MetadataItem) NSMetadataItemAudioSampleRateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAudioSampleRateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiotracknumberkey
func (m_ MetadataItem) NSMetadataItemAudioTrackNumberKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAudioTrackNumberKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemauthoraddresseskey
func (m_ MetadataItem) NSMetadataItemAuthorAddressesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAuthorAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemauthoremailaddresseskey
func (m_ MetadataItem) NSMetadataItemAuthorEmailAddressesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAuthorEmailAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemauthorskey
func (m_ MetadataItem) NSMetadataItemAuthorsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemAuthorsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitembitspersamplekey
func (m_ MetadataItem) NSMetadataItemBitsPerSampleKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemBitsPerSampleKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcfbundleidentifierkey
func (m_ MetadataItem) NSMetadataItemCFBundleIdentifierKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemCFBundleIdentifierKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcameraownerkey
func (m_ MetadataItem) NSMetadataItemCameraOwnerKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemCameraOwnerKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcitykey
func (m_ MetadataItem) NSMetadataItemCityKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemCityKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcodecskey
func (m_ MetadataItem) NSMetadataItemCodecsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemCodecsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcolorspacekey
func (m_ MetadataItem) NSMetadataItemColorSpaceKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemColorSpaceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcommentkey
func (m_ MetadataItem) NSMetadataItemCommentKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemCommentKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcomposerkey
func (m_ MetadataItem) NSMetadataItemComposerKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemComposerKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontactkeywordskey
func (m_ MetadataItem) NSMetadataItemContactKeywordsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemContactKeywordsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontentcreationdatekey
func (m_ MetadataItem) NSMetadataItemContentCreationDateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemContentCreationDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontentmodificationdatekey
func (m_ MetadataItem) NSMetadataItemContentModificationDateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemContentModificationDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontenttypekey
func (m_ MetadataItem) NSMetadataItemContentTypeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemContentTypeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontenttypetreekey
func (m_ MetadataItem) NSMetadataItemContentTypeTreeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemContentTypeTreeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontributorskey
func (m_ MetadataItem) NSMetadataItemContributorsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemContributorsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcopyrightkey
func (m_ MetadataItem) NSMetadataItemCopyrightKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemCopyrightKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcountrykey
func (m_ MetadataItem) NSMetadataItemCountryKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemCountryKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcoveragekey
func (m_ MetadataItem) NSMetadataItemCoverageKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemCoverageKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcreatorkey
func (m_ MetadataItem) NSMetadataItemCreatorKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemCreatorKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdateaddedkey
func (m_ MetadataItem) NSMetadataItemDateAddedKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemDateAddedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdeliverytypekey
func (m_ MetadataItem) NSMetadataItemDeliveryTypeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemDeliveryTypeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdescriptionkey
func (m_ MetadataItem) NSMetadataItemDescriptionKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemDescriptionKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdirectorkey
func (m_ MetadataItem) NSMetadataItemDirectorKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemDirectorKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdisplaynamekey
func (m_ MetadataItem) NSMetadataItemDisplayNameKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemDisplayNameKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdownloadeddatekey
func (m_ MetadataItem) NSMetadataItemDownloadedDateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemDownloadedDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemduedatekey
func (m_ MetadataItem) NSMetadataItemDueDateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemDueDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdurationsecondskey
func (m_ MetadataItem) NSMetadataItemDurationSecondsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemDurationSecondsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexifgpsversionkey
func (m_ MetadataItem) NSMetadataItemEXIFGPSVersionKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemEXIFGPSVersionKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexifversionkey
func (m_ MetadataItem) NSMetadataItemEXIFVersionKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemEXIFVersionKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemeditorskey
func (m_ MetadataItem) NSMetadataItemEditorsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemEditorsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitememailaddresseskey
func (m_ MetadataItem) NSMetadataItemEmailAddressesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemEmailAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemencodingapplicationskey
func (m_ MetadataItem) NSMetadataItemEncodingApplicationsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemEncodingApplicationsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexecutablearchitectureskey
func (m_ MetadataItem) NSMetadataItemExecutableArchitecturesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemExecutableArchitecturesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexecutableplatformkey
func (m_ MetadataItem) NSMetadataItemExecutablePlatformKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemExecutablePlatformKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexposuremodekey
func (m_ MetadataItem) NSMetadataItemExposureModeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemExposureModeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexposureprogramkey
func (m_ MetadataItem) NSMetadataItemExposureProgramKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemExposureProgramKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexposuretimesecondskey
func (m_ MetadataItem) NSMetadataItemExposureTimeSecondsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemExposureTimeSecondsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexposuretimestringkey
func (m_ MetadataItem) NSMetadataItemExposureTimeStringKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemExposureTimeStringKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfnumberkey
func (m_ MetadataItem) NSMetadataItemFNumberKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemFNumberKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfscontentchangedatekey
func (m_ MetadataItem) NSMetadataItemFSContentChangeDateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemFSContentChangeDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfscreationdatekey
func (m_ MetadataItem) NSMetadataItemFSCreationDateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemFSCreationDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfsnamekey
func (m_ MetadataItem) NSMetadataItemFSNameKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemFSNameKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfssizekey
func (m_ MetadataItem) NSMetadataItemFSSizeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemFSSizeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfindercommentkey
func (m_ MetadataItem) NSMetadataItemFinderCommentKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemFinderCommentKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemflashonoffkey
func (m_ MetadataItem) NSMetadataItemFlashOnOffKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemFlashOnOffKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfocallength35mmkey
func (m_ MetadataItem) NSMetadataItemFocalLength35mmKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemFocalLength35mmKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfocallengthkey
func (m_ MetadataItem) NSMetadataItemFocalLengthKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemFocalLengthKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfontskey
func (m_ MetadataItem) NSMetadataItemFontsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemFontsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsareainformationkey
func (m_ MetadataItem) NSMetadataItemGPSAreaInformationKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSAreaInformationKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdopkey
func (m_ MetadataItem) NSMetadataItemGPSDOPKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSDOPKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdatestampkey
func (m_ MetadataItem) NSMetadataItemGPSDateStampKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSDateStampKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestbearingkey
func (m_ MetadataItem) NSMetadataItemGPSDestBearingKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSDestBearingKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestdistancekey
func (m_ MetadataItem) NSMetadataItemGPSDestDistanceKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSDestDistanceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestlatitudekey
func (m_ MetadataItem) NSMetadataItemGPSDestLatitudeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSDestLatitudeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestlongitudekey
func (m_ MetadataItem) NSMetadataItemGPSDestLongitudeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSDestLongitudeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdifferentalkey
func (m_ MetadataItem) NSMetadataItemGPSDifferentalKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSDifferentalKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsmapdatumkey
func (m_ MetadataItem) NSMetadataItemGPSMapDatumKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSMapDatumKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsmeasuremodekey
func (m_ MetadataItem) NSMetadataItemGPSMeasureModeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSMeasureModeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsprocessingmethodkey
func (m_ MetadataItem) NSMetadataItemGPSProcessingMethodKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSProcessingMethodKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsstatuskey
func (m_ MetadataItem) NSMetadataItemGPSStatusKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSStatusKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpstrackkey
func (m_ MetadataItem) NSMetadataItemGPSTrackKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGPSTrackKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgenrekey
func (m_ MetadataItem) NSMetadataItemGenreKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemGenreKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemhasalphachannelkey
func (m_ MetadataItem) NSMetadataItemHasAlphaChannelKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemHasAlphaChannelKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemheadlinekey
func (m_ MetadataItem) NSMetadataItemHeadlineKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemHeadlineKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemisospeedkey
func (m_ MetadataItem) NSMetadataItemISOSpeedKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemISOSpeedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemidentifierkey
func (m_ MetadataItem) NSMetadataItemIdentifierKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemIdentifierKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemimagedirectionkey
func (m_ MetadataItem) NSMetadataItemImageDirectionKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemImageDirectionKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataiteminformationkey
func (m_ MetadataItem) NSMetadataItemInformationKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemInformationKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataiteminstantmessageaddresseskey
func (m_ MetadataItem) NSMetadataItemInstantMessageAddressesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemInstantMessageAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataiteminstructionskey
func (m_ MetadataItem) NSMetadataItemInstructionsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemInstructionsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemisapplicationmanagedkey
func (m_ MetadataItem) NSMetadataItemIsApplicationManagedKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemIsApplicationManagedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemisgeneralmidisequencekey
func (m_ MetadataItem) NSMetadataItemIsGeneralMIDISequenceKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemIsGeneralMIDISequenceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemislikelyjunkkey
func (m_ MetadataItem) NSMetadataItemIsLikelyJunkKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemIsLikelyJunkKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemisubiquitouskey
func (m_ MetadataItem) NSMetadataItemIsUbiquitousKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemIsUbiquitousKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemkeysignaturekey
func (m_ MetadataItem) NSMetadataItemKeySignatureKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemKeySignatureKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemkeywordskey
func (m_ MetadataItem) NSMetadataItemKeywordsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemKeywordsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemkindkey
func (m_ MetadataItem) NSMetadataItemKindKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemKindKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlanguageskey
func (m_ MetadataItem) NSMetadataItemLanguagesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemLanguagesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlastuseddatekey
func (m_ MetadataItem) NSMetadataItemLastUsedDateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemLastUsedDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlatitudekey
func (m_ MetadataItem) NSMetadataItemLatitudeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemLatitudeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlayernameskey
func (m_ MetadataItem) NSMetadataItemLayerNamesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemLayerNamesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlensmodelkey
func (m_ MetadataItem) NSMetadataItemLensModelKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemLensModelKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlongitudekey
func (m_ MetadataItem) NSMetadataItemLongitudeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemLongitudeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlyricistkey
func (m_ MetadataItem) NSMetadataItemLyricistKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemLyricistKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmaxaperturekey
func (m_ MetadataItem) NSMetadataItemMaxApertureKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemMaxApertureKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmediatypeskey
func (m_ MetadataItem) NSMetadataItemMediaTypesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemMediaTypesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmeteringmodekey
func (m_ MetadataItem) NSMetadataItemMeteringModeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemMeteringModeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmusicalgenrekey
func (m_ MetadataItem) NSMetadataItemMusicalGenreKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemMusicalGenreKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmusicalinstrumentcategorykey
func (m_ MetadataItem) NSMetadataItemMusicalInstrumentCategoryKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemMusicalInstrumentCategoryKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmusicalinstrumentnamekey
func (m_ MetadataItem) NSMetadataItemMusicalInstrumentNameKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemMusicalInstrumentNameKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemnamedlocationkey
func (m_ MetadataItem) NSMetadataItemNamedLocationKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemNamedLocationKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemnumberofpageskey
func (m_ MetadataItem) NSMetadataItemNumberOfPagesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemNumberOfPagesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemorganizationskey
func (m_ MetadataItem) NSMetadataItemOrganizationsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemOrganizationsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemorientationkey
func (m_ MetadataItem) NSMetadataItemOrientationKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemOrientationKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemoriginalformatkey
func (m_ MetadataItem) NSMetadataItemOriginalFormatKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemOriginalFormatKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemoriginalsourcekey
func (m_ MetadataItem) NSMetadataItemOriginalSourceKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemOriginalSourceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempageheightkey
func (m_ MetadataItem) NSMetadataItemPageHeightKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemPageHeightKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempagewidthkey
func (m_ MetadataItem) NSMetadataItemPageWidthKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemPageWidthKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemparticipantskey
func (m_ MetadataItem) NSMetadataItemParticipantsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemParticipantsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempathkey
func (m_ MetadataItem) NSMetadataItemPathKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemPathKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemperformerskey
func (m_ MetadataItem) NSMetadataItemPerformersKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemPerformersKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemphonenumberskey
func (m_ MetadataItem) NSMetadataItemPhoneNumbersKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemPhoneNumbersKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempixelcountkey
func (m_ MetadataItem) NSMetadataItemPixelCountKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemPixelCountKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempixelheightkey
func (m_ MetadataItem) NSMetadataItemPixelHeightKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemPixelHeightKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempixelwidthkey
func (m_ MetadataItem) NSMetadataItemPixelWidthKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemPixelWidthKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemproducerkey
func (m_ MetadataItem) NSMetadataItemProducerKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemProducerKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemprofilenamekey
func (m_ MetadataItem) NSMetadataItemProfileNameKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemProfileNameKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemprojectskey
func (m_ MetadataItem) NSMetadataItemProjectsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemProjectsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempublisherskey
func (m_ MetadataItem) NSMetadataItemPublishersKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemPublishersKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrecipientaddresseskey
func (m_ MetadataItem) NSMetadataItemRecipientAddressesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemRecipientAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrecipientemailaddresseskey
func (m_ MetadataItem) NSMetadataItemRecipientEmailAddressesKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemRecipientEmailAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrecipientskey
func (m_ MetadataItem) NSMetadataItemRecipientsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemRecipientsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrecordingdatekey
func (m_ MetadataItem) NSMetadataItemRecordingDateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemRecordingDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrecordingyearkey
func (m_ MetadataItem) NSMetadataItemRecordingYearKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemRecordingYearKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemredeyeonoffkey
func (m_ MetadataItem) NSMetadataItemRedEyeOnOffKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemRedEyeOnOffKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemresolutionheightdpikey
func (m_ MetadataItem) NSMetadataItemResolutionHeightDPIKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemResolutionHeightDPIKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemresolutionwidthdpikey
func (m_ MetadataItem) NSMetadataItemResolutionWidthDPIKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemResolutionWidthDPIKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrightskey
func (m_ MetadataItem) NSMetadataItemRightsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemRightsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemsecuritymethodkey
func (m_ MetadataItem) NSMetadataItemSecurityMethodKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemSecurityMethodKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemspeedkey
func (m_ MetadataItem) NSMetadataItemSpeedKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemSpeedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemstarratingkey
func (m_ MetadataItem) NSMetadataItemStarRatingKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemStarRatingKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemstateorprovincekey
func (m_ MetadataItem) NSMetadataItemStateOrProvinceKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemStateOrProvinceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemstreamablekey
func (m_ MetadataItem) NSMetadataItemStreamableKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemStreamableKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemsubjectkey
func (m_ MetadataItem) NSMetadataItemSubjectKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemSubjectKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtempokey
func (m_ MetadataItem) NSMetadataItemTempoKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemTempoKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtextcontentkey
func (m_ MetadataItem) NSMetadataItemTextContentKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemTextContentKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemthemekey
func (m_ MetadataItem) NSMetadataItemThemeKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemThemeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtimesignaturekey
func (m_ MetadataItem) NSMetadataItemTimeSignatureKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemTimeSignatureKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtimestampkey
func (m_ MetadataItem) NSMetadataItemTimestampKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemTimestampKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtitlekey
func (m_ MetadataItem) NSMetadataItemTitleKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemTitleKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtotalbitratekey
func (m_ MetadataItem) NSMetadataItemTotalBitRateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemTotalBitRateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemurlkey
func (m_ MetadataItem) NSMetadataItemURLKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemURLKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemversionkey
func (m_ MetadataItem) NSMetadataItemVersionKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemVersionKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemvideobitratekey
func (m_ MetadataItem) NSMetadataItemVideoBitRateKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemVideoBitRateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemwherefromskey
func (m_ MetadataItem) NSMetadataItemWhereFromsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemWhereFromsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemwhitebalancekey
func (m_ MetadataItem) NSMetadataItemWhiteBalanceKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataItemWhiteBalanceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemcontainerdisplaynamekey
func (m_ MetadataItem) NSMetadataUbiquitousItemContainerDisplayNameKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemContainerDisplayNameKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadrequestedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadRequestedKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadRequestedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingerrorkey
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadingErrorKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingErrorKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatuscurrent
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadingStatusCurrent() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusCurrent"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatusdownloaded
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadingStatusDownloaded() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusDownloaded"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatuskey
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadingStatusKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatusnotdownloaded
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadingStatusNotDownloaded() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusNotDownloaded"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemhasunresolvedconflictskey
func (m_ MetadataItem) NSMetadataUbiquitousItemHasUnresolvedConflictsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemHasUnresolvedConflictsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisdownloadedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsDownloadedKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsDownloadedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisdownloadingkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsDownloadingKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsDownloadingKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisexternaldocumentkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsExternalDocumentKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsExternalDocumentKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemissharedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsSharedKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsSharedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisuploadedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsUploadedKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsUploadedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisuploadingkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsUploadingKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsUploadingKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitempercentdownloadedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemPercentDownloadedKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemPercentDownloadedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitempercentuploadedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemPercentUploadedKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemPercentUploadedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemurlinlocalcontainerkey
func (m_ MetadataItem) NSMetadataUbiquitousItemURLInLocalContainerKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemURLInLocalContainerKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemuploadingerrorkey
func (m_ MetadataItem) NSMetadataUbiquitousItemUploadingErrorKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousItemUploadingErrorKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemcurrentuserpermissionskey
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemcurrentuserrolekey
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemCurrentUserRoleKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemCurrentUserRoleKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemmostrecenteditornamecomponentskey
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemownernamecomponentskey
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemOwnerNameComponentsKey() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemOwnerNameComponentsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditempermissionsreadonly
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemPermissionsReadOnly() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemPermissionsReadOnly"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditempermissionsreadwrite
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemPermissionsReadWrite() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemPermissionsReadWrite"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemroleowner
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemRoleOwner() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemRoleOwner"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemroleparticipant
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemRoleParticipant() string {
	rv := objc.Send[string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemRoleParticipant"))
	return rv
}



