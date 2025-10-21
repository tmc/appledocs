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
func (m_ MetadataItem) Attributes() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("attributes"))
	return rv
}


// SetAttributes sets the value of the attributes property.
// An array containing the attribute keys for the metadata item’s values.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitem/attributes
func (m_ MetadataItem) SetAttributes(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributes:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemacquisitionmakekey
func (m_ MetadataItem) NSMetadataItemAcquisitionMakeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAcquisitionMakeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemacquisitionmodelkey
func (m_ MetadataItem) NSMetadataItemAcquisitionModelKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAcquisitionModelKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemalbumkey
func (m_ MetadataItem) NSMetadataItemAlbumKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAlbumKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaltitudekey
func (m_ MetadataItem) NSMetadataItemAltitudeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAltitudeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaperturekey
func (m_ MetadataItem) NSMetadataItemApertureKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemApertureKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopdescriptorskey
func (m_ MetadataItem) NSMetadataItemAppleLoopDescriptorsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAppleLoopDescriptorsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopskeyfiltertypekey
func (m_ MetadataItem) NSMetadataItemAppleLoopsKeyFilterTypeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAppleLoopsKeyFilterTypeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopsloopmodekey
func (m_ MetadataItem) NSMetadataItemAppleLoopsLoopModeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAppleLoopsLoopModeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemappleloopsrootkeykey
func (m_ MetadataItem) NSMetadataItemAppleLoopsRootKeyKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAppleLoopsRootKeyKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemapplicationcategorieskey
func (m_ MetadataItem) NSMetadataItemApplicationCategoriesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemApplicationCategoriesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemattributechangedatekey
func (m_ MetadataItem) NSMetadataItemAttributeChangeDateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAttributeChangeDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudienceskey
func (m_ MetadataItem) NSMetadataItemAudiencesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAudiencesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiobitratekey
func (m_ MetadataItem) NSMetadataItemAudioBitRateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAudioBitRateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiochannelcountkey
func (m_ MetadataItem) NSMetadataItemAudioChannelCountKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAudioChannelCountKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudioencodingapplicationkey
func (m_ MetadataItem) NSMetadataItemAudioEncodingApplicationKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAudioEncodingApplicationKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiosampleratekey
func (m_ MetadataItem) NSMetadataItemAudioSampleRateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAudioSampleRateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemaudiotracknumberkey
func (m_ MetadataItem) NSMetadataItemAudioTrackNumberKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAudioTrackNumberKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemauthoraddresseskey
func (m_ MetadataItem) NSMetadataItemAuthorAddressesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAuthorAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemauthoremailaddresseskey
func (m_ MetadataItem) NSMetadataItemAuthorEmailAddressesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAuthorEmailAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemauthorskey
func (m_ MetadataItem) NSMetadataItemAuthorsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemAuthorsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitembitspersamplekey
func (m_ MetadataItem) NSMetadataItemBitsPerSampleKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemBitsPerSampleKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcfbundleidentifierkey
func (m_ MetadataItem) NSMetadataItemCFBundleIdentifierKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemCFBundleIdentifierKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcameraownerkey
func (m_ MetadataItem) NSMetadataItemCameraOwnerKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemCameraOwnerKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcitykey
func (m_ MetadataItem) NSMetadataItemCityKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemCityKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcodecskey
func (m_ MetadataItem) NSMetadataItemCodecsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemCodecsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcolorspacekey
func (m_ MetadataItem) NSMetadataItemColorSpaceKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemColorSpaceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcommentkey
func (m_ MetadataItem) NSMetadataItemCommentKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemCommentKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcomposerkey
func (m_ MetadataItem) NSMetadataItemComposerKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemComposerKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontactkeywordskey
func (m_ MetadataItem) NSMetadataItemContactKeywordsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemContactKeywordsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontentcreationdatekey
func (m_ MetadataItem) NSMetadataItemContentCreationDateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemContentCreationDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontentmodificationdatekey
func (m_ MetadataItem) NSMetadataItemContentModificationDateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemContentModificationDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontenttypekey
func (m_ MetadataItem) NSMetadataItemContentTypeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemContentTypeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontenttypetreekey
func (m_ MetadataItem) NSMetadataItemContentTypeTreeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemContentTypeTreeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcontributorskey
func (m_ MetadataItem) NSMetadataItemContributorsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemContributorsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcopyrightkey
func (m_ MetadataItem) NSMetadataItemCopyrightKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemCopyrightKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcountrykey
func (m_ MetadataItem) NSMetadataItemCountryKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemCountryKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcoveragekey
func (m_ MetadataItem) NSMetadataItemCoverageKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemCoverageKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemcreatorkey
func (m_ MetadataItem) NSMetadataItemCreatorKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemCreatorKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdateaddedkey
func (m_ MetadataItem) NSMetadataItemDateAddedKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemDateAddedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdeliverytypekey
func (m_ MetadataItem) NSMetadataItemDeliveryTypeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemDeliveryTypeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdescriptionkey
func (m_ MetadataItem) NSMetadataItemDescriptionKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemDescriptionKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdirectorkey
func (m_ MetadataItem) NSMetadataItemDirectorKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemDirectorKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdisplaynamekey
func (m_ MetadataItem) NSMetadataItemDisplayNameKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemDisplayNameKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdownloadeddatekey
func (m_ MetadataItem) NSMetadataItemDownloadedDateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemDownloadedDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemduedatekey
func (m_ MetadataItem) NSMetadataItemDueDateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemDueDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemdurationsecondskey
func (m_ MetadataItem) NSMetadataItemDurationSecondsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemDurationSecondsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexifgpsversionkey
func (m_ MetadataItem) NSMetadataItemEXIFGPSVersionKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemEXIFGPSVersionKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexifversionkey
func (m_ MetadataItem) NSMetadataItemEXIFVersionKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemEXIFVersionKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemeditorskey
func (m_ MetadataItem) NSMetadataItemEditorsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemEditorsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitememailaddresseskey
func (m_ MetadataItem) NSMetadataItemEmailAddressesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemEmailAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemencodingapplicationskey
func (m_ MetadataItem) NSMetadataItemEncodingApplicationsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemEncodingApplicationsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexecutablearchitectureskey
func (m_ MetadataItem) NSMetadataItemExecutableArchitecturesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemExecutableArchitecturesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexecutableplatformkey
func (m_ MetadataItem) NSMetadataItemExecutablePlatformKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemExecutablePlatformKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexposuremodekey
func (m_ MetadataItem) NSMetadataItemExposureModeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemExposureModeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexposureprogramkey
func (m_ MetadataItem) NSMetadataItemExposureProgramKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemExposureProgramKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexposuretimesecondskey
func (m_ MetadataItem) NSMetadataItemExposureTimeSecondsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemExposureTimeSecondsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemexposuretimestringkey
func (m_ MetadataItem) NSMetadataItemExposureTimeStringKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemExposureTimeStringKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfnumberkey
func (m_ MetadataItem) NSMetadataItemFNumberKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemFNumberKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfscontentchangedatekey
func (m_ MetadataItem) NSMetadataItemFSContentChangeDateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemFSContentChangeDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfscreationdatekey
func (m_ MetadataItem) NSMetadataItemFSCreationDateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemFSCreationDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfsnamekey
func (m_ MetadataItem) NSMetadataItemFSNameKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemFSNameKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfssizekey
func (m_ MetadataItem) NSMetadataItemFSSizeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemFSSizeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfindercommentkey
func (m_ MetadataItem) NSMetadataItemFinderCommentKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemFinderCommentKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemflashonoffkey
func (m_ MetadataItem) NSMetadataItemFlashOnOffKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemFlashOnOffKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfocallength35mmkey
func (m_ MetadataItem) NSMetadataItemFocalLength35mmKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemFocalLength35mmKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfocallengthkey
func (m_ MetadataItem) NSMetadataItemFocalLengthKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemFocalLengthKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemfontskey
func (m_ MetadataItem) NSMetadataItemFontsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemFontsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsareainformationkey
func (m_ MetadataItem) NSMetadataItemGPSAreaInformationKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSAreaInformationKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdopkey
func (m_ MetadataItem) NSMetadataItemGPSDOPKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSDOPKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdatestampkey
func (m_ MetadataItem) NSMetadataItemGPSDateStampKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSDateStampKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestbearingkey
func (m_ MetadataItem) NSMetadataItemGPSDestBearingKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSDestBearingKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestdistancekey
func (m_ MetadataItem) NSMetadataItemGPSDestDistanceKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSDestDistanceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestlatitudekey
func (m_ MetadataItem) NSMetadataItemGPSDestLatitudeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSDestLatitudeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdestlongitudekey
func (m_ MetadataItem) NSMetadataItemGPSDestLongitudeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSDestLongitudeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsdifferentalkey
func (m_ MetadataItem) NSMetadataItemGPSDifferentalKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSDifferentalKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsmapdatumkey
func (m_ MetadataItem) NSMetadataItemGPSMapDatumKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSMapDatumKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsmeasuremodekey
func (m_ MetadataItem) NSMetadataItemGPSMeasureModeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSMeasureModeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsprocessingmethodkey
func (m_ MetadataItem) NSMetadataItemGPSProcessingMethodKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSProcessingMethodKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpsstatuskey
func (m_ MetadataItem) NSMetadataItemGPSStatusKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSStatusKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgpstrackkey
func (m_ MetadataItem) NSMetadataItemGPSTrackKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGPSTrackKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemgenrekey
func (m_ MetadataItem) NSMetadataItemGenreKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemGenreKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemhasalphachannelkey
func (m_ MetadataItem) NSMetadataItemHasAlphaChannelKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemHasAlphaChannelKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemheadlinekey
func (m_ MetadataItem) NSMetadataItemHeadlineKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemHeadlineKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemisospeedkey
func (m_ MetadataItem) NSMetadataItemISOSpeedKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemISOSpeedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemidentifierkey
func (m_ MetadataItem) NSMetadataItemIdentifierKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemIdentifierKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemimagedirectionkey
func (m_ MetadataItem) NSMetadataItemImageDirectionKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemImageDirectionKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataiteminformationkey
func (m_ MetadataItem) NSMetadataItemInformationKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemInformationKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataiteminstantmessageaddresseskey
func (m_ MetadataItem) NSMetadataItemInstantMessageAddressesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemInstantMessageAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataiteminstructionskey
func (m_ MetadataItem) NSMetadataItemInstructionsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemInstructionsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemisapplicationmanagedkey
func (m_ MetadataItem) NSMetadataItemIsApplicationManagedKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemIsApplicationManagedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemisgeneralmidisequencekey
func (m_ MetadataItem) NSMetadataItemIsGeneralMIDISequenceKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemIsGeneralMIDISequenceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemislikelyjunkkey
func (m_ MetadataItem) NSMetadataItemIsLikelyJunkKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemIsLikelyJunkKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemisubiquitouskey
func (m_ MetadataItem) NSMetadataItemIsUbiquitousKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemIsUbiquitousKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemkeysignaturekey
func (m_ MetadataItem) NSMetadataItemKeySignatureKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemKeySignatureKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemkeywordskey
func (m_ MetadataItem) NSMetadataItemKeywordsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemKeywordsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemkindkey
func (m_ MetadataItem) NSMetadataItemKindKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemKindKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlanguageskey
func (m_ MetadataItem) NSMetadataItemLanguagesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemLanguagesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlastuseddatekey
func (m_ MetadataItem) NSMetadataItemLastUsedDateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemLastUsedDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlatitudekey
func (m_ MetadataItem) NSMetadataItemLatitudeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemLatitudeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlayernameskey
func (m_ MetadataItem) NSMetadataItemLayerNamesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemLayerNamesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlensmodelkey
func (m_ MetadataItem) NSMetadataItemLensModelKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemLensModelKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlongitudekey
func (m_ MetadataItem) NSMetadataItemLongitudeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemLongitudeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemlyricistkey
func (m_ MetadataItem) NSMetadataItemLyricistKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemLyricistKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmaxaperturekey
func (m_ MetadataItem) NSMetadataItemMaxApertureKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemMaxApertureKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmediatypeskey
func (m_ MetadataItem) NSMetadataItemMediaTypesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemMediaTypesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmeteringmodekey
func (m_ MetadataItem) NSMetadataItemMeteringModeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemMeteringModeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmusicalgenrekey
func (m_ MetadataItem) NSMetadataItemMusicalGenreKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemMusicalGenreKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmusicalinstrumentcategorykey
func (m_ MetadataItem) NSMetadataItemMusicalInstrumentCategoryKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemMusicalInstrumentCategoryKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemmusicalinstrumentnamekey
func (m_ MetadataItem) NSMetadataItemMusicalInstrumentNameKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemMusicalInstrumentNameKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemnamedlocationkey
func (m_ MetadataItem) NSMetadataItemNamedLocationKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemNamedLocationKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemnumberofpageskey
func (m_ MetadataItem) NSMetadataItemNumberOfPagesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemNumberOfPagesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemorganizationskey
func (m_ MetadataItem) NSMetadataItemOrganizationsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemOrganizationsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemorientationkey
func (m_ MetadataItem) NSMetadataItemOrientationKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemOrientationKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemoriginalformatkey
func (m_ MetadataItem) NSMetadataItemOriginalFormatKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemOriginalFormatKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemoriginalsourcekey
func (m_ MetadataItem) NSMetadataItemOriginalSourceKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemOriginalSourceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempageheightkey
func (m_ MetadataItem) NSMetadataItemPageHeightKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemPageHeightKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempagewidthkey
func (m_ MetadataItem) NSMetadataItemPageWidthKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemPageWidthKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemparticipantskey
func (m_ MetadataItem) NSMetadataItemParticipantsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemParticipantsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempathkey
func (m_ MetadataItem) NSMetadataItemPathKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemPathKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemperformerskey
func (m_ MetadataItem) NSMetadataItemPerformersKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemPerformersKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemphonenumberskey
func (m_ MetadataItem) NSMetadataItemPhoneNumbersKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemPhoneNumbersKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempixelcountkey
func (m_ MetadataItem) NSMetadataItemPixelCountKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemPixelCountKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempixelheightkey
func (m_ MetadataItem) NSMetadataItemPixelHeightKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemPixelHeightKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempixelwidthkey
func (m_ MetadataItem) NSMetadataItemPixelWidthKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemPixelWidthKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemproducerkey
func (m_ MetadataItem) NSMetadataItemProducerKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemProducerKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemprofilenamekey
func (m_ MetadataItem) NSMetadataItemProfileNameKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemProfileNameKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemprojectskey
func (m_ MetadataItem) NSMetadataItemProjectsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemProjectsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitempublisherskey
func (m_ MetadataItem) NSMetadataItemPublishersKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemPublishersKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrecipientaddresseskey
func (m_ MetadataItem) NSMetadataItemRecipientAddressesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemRecipientAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrecipientemailaddresseskey
func (m_ MetadataItem) NSMetadataItemRecipientEmailAddressesKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemRecipientEmailAddressesKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrecipientskey
func (m_ MetadataItem) NSMetadataItemRecipientsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemRecipientsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrecordingdatekey
func (m_ MetadataItem) NSMetadataItemRecordingDateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemRecordingDateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrecordingyearkey
func (m_ MetadataItem) NSMetadataItemRecordingYearKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemRecordingYearKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemredeyeonoffkey
func (m_ MetadataItem) NSMetadataItemRedEyeOnOffKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemRedEyeOnOffKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemresolutionheightdpikey
func (m_ MetadataItem) NSMetadataItemResolutionHeightDPIKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemResolutionHeightDPIKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemresolutionwidthdpikey
func (m_ MetadataItem) NSMetadataItemResolutionWidthDPIKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemResolutionWidthDPIKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemrightskey
func (m_ MetadataItem) NSMetadataItemRightsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemRightsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemsecuritymethodkey
func (m_ MetadataItem) NSMetadataItemSecurityMethodKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemSecurityMethodKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemspeedkey
func (m_ MetadataItem) NSMetadataItemSpeedKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemSpeedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemstarratingkey
func (m_ MetadataItem) NSMetadataItemStarRatingKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemStarRatingKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemstateorprovincekey
func (m_ MetadataItem) NSMetadataItemStateOrProvinceKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemStateOrProvinceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemstreamablekey
func (m_ MetadataItem) NSMetadataItemStreamableKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemStreamableKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemsubjectkey
func (m_ MetadataItem) NSMetadataItemSubjectKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemSubjectKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtempokey
func (m_ MetadataItem) NSMetadataItemTempoKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemTempoKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtextcontentkey
func (m_ MetadataItem) NSMetadataItemTextContentKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemTextContentKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemthemekey
func (m_ MetadataItem) NSMetadataItemThemeKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemThemeKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtimesignaturekey
func (m_ MetadataItem) NSMetadataItemTimeSignatureKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemTimeSignatureKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtimestampkey
func (m_ MetadataItem) NSMetadataItemTimestampKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemTimestampKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtitlekey
func (m_ MetadataItem) NSMetadataItemTitleKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemTitleKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemtotalbitratekey
func (m_ MetadataItem) NSMetadataItemTotalBitRateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemTotalBitRateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemurlkey
func (m_ MetadataItem) NSMetadataItemURLKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemURLKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemversionkey
func (m_ MetadataItem) NSMetadataItemVersionKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemVersionKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemvideobitratekey
func (m_ MetadataItem) NSMetadataItemVideoBitRateKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemVideoBitRateKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemwherefromskey
func (m_ MetadataItem) NSMetadataItemWhereFromsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemWhereFromsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataitemwhitebalancekey
func (m_ MetadataItem) NSMetadataItemWhiteBalanceKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataItemWhiteBalanceKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemcontainerdisplaynamekey
func (m_ MetadataItem) NSMetadataUbiquitousItemContainerDisplayNameKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemContainerDisplayNameKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadrequestedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadRequestedKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadRequestedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingerrorkey
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadingErrorKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingErrorKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatuscurrent
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadingStatusCurrent() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusCurrent"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatusdownloaded
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadingStatusDownloaded() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusDownloaded"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatuskey
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadingStatusKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemdownloadingstatusnotdownloaded
func (m_ MetadataItem) NSMetadataUbiquitousItemDownloadingStatusNotDownloaded() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemDownloadingStatusNotDownloaded"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemhasunresolvedconflictskey
func (m_ MetadataItem) NSMetadataUbiquitousItemHasUnresolvedConflictsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemHasUnresolvedConflictsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisdownloadedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsDownloadedKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsDownloadedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisdownloadingkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsDownloadingKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsDownloadingKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisexternaldocumentkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsExternalDocumentKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsExternalDocumentKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemissharedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsSharedKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsSharedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisuploadedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsUploadedKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsUploadedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemisuploadingkey
func (m_ MetadataItem) NSMetadataUbiquitousItemIsUploadingKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemIsUploadingKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitempercentdownloadedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemPercentDownloadedKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemPercentDownloadedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitempercentuploadedkey
func (m_ MetadataItem) NSMetadataUbiquitousItemPercentUploadedKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemPercentUploadedKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemurlinlocalcontainerkey
func (m_ MetadataItem) NSMetadataUbiquitousItemURLInLocalContainerKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemURLInLocalContainerKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousitemuploadingerrorkey
func (m_ MetadataItem) NSMetadataUbiquitousItemUploadingErrorKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousItemUploadingErrorKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemcurrentuserpermissionskey
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemCurrentUserPermissionsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemcurrentuserrolekey
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemCurrentUserRoleKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemCurrentUserRoleKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemmostrecenteditornamecomponentskey
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemMostRecentEditorNameComponentsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemownernamecomponentskey
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemOwnerNameComponentsKey() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemOwnerNameComponentsKey"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditempermissionsreadonly
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemPermissionsReadOnly() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemPermissionsReadOnly"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditempermissionsreadwrite
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemPermissionsReadWrite() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemPermissionsReadWrite"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemroleowner
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemRoleOwner() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemRoleOwner"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmetadataubiquitousshareditemroleparticipant
func (m_ MetadataItem) NSMetadataUbiquitousSharedItemRoleParticipant() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("NSMetadataUbiquitousSharedItemRoleParticipant"))
	return rv
}



