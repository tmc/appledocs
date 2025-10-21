// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// CoreFoundation Functions (825 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_CFRelease func(unsafe.Pointer) unsafe.Pointer
	_CFMakeCollectable func(unsafe.Pointer) unsafe.Pointer
	_CFRetain func(unsafe.Pointer) unsafe.Pointer
	_CFAbsoluteTimeAddGregorianUnits func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAbsoluteTimeGetCurrent func() unsafe.Pointer
	_CFAbsoluteTimeGetDayOfWeek func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAbsoluteTimeGetDayOfYear func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAbsoluteTimeGetDifferenceAsGregorianUnits func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAbsoluteTimeGetGregorianDate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAbsoluteTimeGetWeekOfYear func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorAllocate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorAllocateBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorAllocateTyped func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorCreateWithZone func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorDeallocate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorGetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorGetDefault func() unsafe.Pointer
	_CFAllocatorGetPreferredSizeForSize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorGetTypeID func() unsafe.Pointer
	_CFAllocatorReallocate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorReallocateBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorReallocateTyped func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorSetDefault func(unsafe.Pointer) unsafe.Pointer
	_CFArrayAppendArray func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayAppendValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayApplyFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayBSearchValues func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayContainsValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayCreateMutable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayExchangeValuesAtIndices func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFArrayGetCountOfValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayGetFirstIndexOfValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayGetLastIndexOfValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayGetTypeID func() unsafe.Pointer
	_CFArrayGetValueAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayGetValues func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayInsertValueAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayRemoveAllValues func(unsafe.Pointer) unsafe.Pointer
	_CFArrayRemoveValueAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayReplaceValues func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArraySetValueAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArraySortValues func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringBeginEditing func(unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreateMutable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreateWithSubstring func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringEndEditing func(unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttributeAndLongestEffectiveRange func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttributesAndLongestEffectiveRange func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetBidiLevelsAndResolvedDirections func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetLength func(unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetMutableString func(unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetStatisticalWritingDirections func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetString func(unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetTypeID func() unsafe.Pointer
	_CFAttributedStringRemoveAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringReplaceAttributedString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringReplaceString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringSetAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringSetAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAutorelease func(unsafe.Pointer) unsafe.Pointer
	_CFBagAddValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagApplyFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagContainsValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagCreateMutable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFBagGetCountOfValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagGetTypeID func() unsafe.Pointer
	_CFBagGetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagGetValueIfPresent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagGetValues func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagRemoveAllValues func(unsafe.Pointer) unsafe.Pointer
	_CFBagRemoveValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagReplaceValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagSetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapAddValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapApplyFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapContainsValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapCreateCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetCountOfValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetMinimum func(unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetMinimumIfPresent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetTypeID func() unsafe.Pointer
	_CFBinaryHeapGetValues func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapRemoveAllValues func(unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapRemoveMinimumValue func(unsafe.Pointer) unsafe.Pointer
	_CFBitVectorContainsBit func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorCreateMutable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorFlipBitAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorFlipBits func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorGetBitAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorGetBits func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFBitVectorGetCountOfBit func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorGetFirstIndexOfBit func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorGetLastIndexOfBit func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorGetTypeID func() unsafe.Pointer
	_CFBitVectorSetAllBits func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorSetBitAtIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorSetBits func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBitVectorSetCount func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBooleanGetTypeID func() unsafe.Pointer
	_CFBooleanGetValue func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCloseBundleResourceMap func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyAuxiliaryExecutableURL func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyBuiltInPlugInsURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyBundleLocalizations func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyBundleURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyExecutableArchitectures func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyExecutableArchitecturesForURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyExecutableURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyInfoDictionaryForURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyInfoDictionaryInDirectory func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyLocalizationsForPreferences func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyLocalizationsForURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyLocalizedString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyLocalizedStringForLocalizations func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyPreferredLocalizationsFromArray func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyPrivateFrameworksURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyResourceURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyResourceURLForLocalization func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyResourceURLInDirectory func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyResourceURLsOfType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyResourceURLsOfTypeForLocalization func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyResourceURLsOfTypeInDirectory func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCopyResourcesDirectoryURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopySharedFrameworksURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopySharedSupportURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCopySupportFilesDirectoryURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleCreateBundlesFromDirectory func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetAllBundles func() unsafe.Pointer
	_CFBundleGetBundleWithIdentifier func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetDataPointerForName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetDataPointersForNames func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetDevelopmentRegion func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetFunctionPointerForName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetFunctionPointersForNames func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetIdentifier func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetInfoDictionary func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetLocalInfoDictionary func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetMainBundle func() unsafe.Pointer
	_CFBundleGetPackageInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetPackageInfoInDirectory func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetPlugIn func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetTypeID func() unsafe.Pointer
	_CFBundleGetValueForInfoDictionaryKey func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetVersionNumber func(unsafe.Pointer) unsafe.Pointer
	_CFBundleIsArchitectureLoadable func(unsafe.Pointer) unsafe.Pointer
	_CFBundleIsExecutableLoadable func(unsafe.Pointer) unsafe.Pointer
	_CFBundleIsExecutableLoadableForURL func(unsafe.Pointer) unsafe.Pointer
	_CFBundleIsExecutableLoaded func(unsafe.Pointer) unsafe.Pointer
	_CFBundleLoadExecutable func(unsafe.Pointer) unsafe.Pointer
	_CFBundleLoadExecutableAndReturnError func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleOpenBundleResourceFiles func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleOpenBundleResourceMap func(unsafe.Pointer) unsafe.Pointer
	_CFBundlePreflightExecutable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleUnloadExecutable func(unsafe.Pointer) unsafe.Pointer
	_CFCalendarAddComponents func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarComposeAbsoluteTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarCopyCurrent func() unsafe.Pointer
	_CFCalendarCopyLocale func(unsafe.Pointer) unsafe.Pointer
	_CFCalendarCopyTimeZone func(unsafe.Pointer) unsafe.Pointer
	_CFCalendarCreateWithIdentifier func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarDecomposeAbsoluteTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetComponentDifference func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetFirstWeekday func(unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetIdentifier func(unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetMaximumRangeOfUnit func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetMinimumDaysInFirstWeek func(unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetMinimumRangeOfUnit func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetOrdinalityOfUnit func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetRangeOfUnit func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetTimeRangeOfUnit func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetTypeID func() unsafe.Pointer
	_CFCalendarSetFirstWeekday func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarSetLocale func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarSetMinimumDaysInFirstWeek func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarSetTimeZone func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetAddCharactersInRange func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetAddCharactersInString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetCreateBitmapRepresentation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetCreateInvertedSet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetCreateMutable func(unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetCreateWithBitmapRepresentation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetCreateWithCharactersInRange func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetCreateWithCharactersInString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetGetPredefined func(unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetGetTypeID func() unsafe.Pointer
	_CFCharacterSetHasMemberInPlane func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetIntersect func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetInvert func(unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetIsCharacterMember func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetIsLongCharacterMember func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetIsSupersetOfSet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetRemoveCharactersInRange func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetRemoveCharactersInString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetUnion func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCopyDescription func(unsafe.Pointer) unsafe.Pointer
	_CFCopyHomeDirectoryURL func() unsafe.Pointer
	_CFCopyTypeIDDescription func(unsafe.Pointer) unsafe.Pointer
	_CFDataAppendBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreateMutable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreateWithBytesNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataDeleteBytes func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataFind func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataGetBytePtr func(unsafe.Pointer) unsafe.Pointer
	_CFDataGetBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataGetLength func(unsafe.Pointer) unsafe.Pointer
	_CFDataGetMutableBytePtr func(unsafe.Pointer) unsafe.Pointer
	_CFDataGetTypeID func() unsafe.Pointer
	_CFDataIncreaseLength func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataReplaceBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataSetLength func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateCompare func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCopyProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreateDateFormatFromTemplate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreateDateFromString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreateISO8601Formatter func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreateStringWithAbsoluteTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreateStringWithDate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetAbsoluteTimeFromString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetDateStyle func(unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetFormat func(unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetLocale func(unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetTimeStyle func(unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetTypeID func() unsafe.Pointer
	_CFDateFormatterSetFormat func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateGetAbsoluteTime func(unsafe.Pointer) unsafe.Pointer
	_CFDateGetTimeIntervalSinceDate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateGetTypeID func() unsafe.Pointer
	_CFDictionaryAddValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryApplyFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryContainsKey func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryContainsValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryCreateMutable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetCountOfKey func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetCountOfValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetKeysAndValues func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetTypeID func() unsafe.Pointer
	_CFDictionaryGetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetValueIfPresent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryRemoveAllValues func(unsafe.Pointer) unsafe.Pointer
	_CFDictionaryRemoveValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryReplaceValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionarySetValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFEqual func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFErrorCopyDescription func(unsafe.Pointer) unsafe.Pointer
	_CFErrorCopyFailureReason func(unsafe.Pointer) unsafe.Pointer
	_CFErrorCopyRecoverySuggestion func(unsafe.Pointer) unsafe.Pointer
	_CFErrorCopyUserInfo func(unsafe.Pointer) unsafe.Pointer
	_CFErrorCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFErrorCreateWithUserInfoKeysAndValues func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFErrorGetCode func(unsafe.Pointer) unsafe.Pointer
	_CFErrorGetDomain func(unsafe.Pointer) unsafe.Pointer
	_CFErrorGetTypeID func() unsafe.Pointer
	_CFFileDescriptorCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileDescriptorCreateRunLoopSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileDescriptorDisableCallBacks func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileDescriptorEnableCallBacks func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileDescriptorGetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileDescriptorGetNativeDescriptor func(unsafe.Pointer) unsafe.Pointer
	_CFFileDescriptorGetTypeID func() unsafe.Pointer
	_CFFileDescriptorInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CFFileDescriptorIsValid func(unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityClearProperties func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityCopyAccessControlList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityCopyGroupUUID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityCopyOwnerUUID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityCreate func(unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityGetGroup func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityGetMode func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityGetOwner func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityGetTypeID func() unsafe.Pointer
	_CFFileSecuritySetAccessControlList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecuritySetGroup func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecuritySetGroupUUID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecuritySetMode func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecuritySetOwner func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileSecuritySetOwnerUUID func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFGetAllocator func(unsafe.Pointer) unsafe.Pointer
	_CFGetRetainCount func(unsafe.Pointer) unsafe.Pointer
	_CFGetTypeID func(unsafe.Pointer) unsafe.Pointer
	_CFGregorianDateGetAbsoluteTime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFGregorianDateIsValid func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFHash func(unsafe.Pointer) unsafe.Pointer
	_CFLocaleCopyAvailableLocaleIdentifiers func() unsafe.Pointer
	_CFLocaleCopyCommonISOCurrencyCodes func() unsafe.Pointer
	_CFLocaleCopyCurrent func() unsafe.Pointer
	_CFLocaleCopyDisplayNameForPropertyValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleCopyISOCountryCodes func() unsafe.Pointer
	_CFLocaleCopyISOCurrencyCodes func() unsafe.Pointer
	_CFLocaleCopyISOLanguageCodes func() unsafe.Pointer
	_CFLocaleCopyPreferredLanguages func() unsafe.Pointer
	_CFLocaleCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleCreateCanonicalLanguageIdentifierFromString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleCreateCanonicalLocaleIdentifierFromString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleCreateComponentsFromLocaleIdentifier func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleCreateLocaleIdentifierFromComponents func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleGetIdentifier func(unsafe.Pointer) unsafe.Pointer
	_CFLocaleGetLanguageCharacterDirection func(unsafe.Pointer) unsafe.Pointer
	_CFLocaleGetLanguageLineDirection func(unsafe.Pointer) unsafe.Pointer
	_CFLocaleGetSystem func() unsafe.Pointer
	_CFLocaleGetTypeID func() unsafe.Pointer
	_CFLocaleGetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier func(unsafe.Pointer) unsafe.Pointer
	_CFMachPortCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMachPortCreateRunLoopSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMachPortCreateWithPort func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMachPortGetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMachPortGetInvalidationCallBack func(unsafe.Pointer) unsafe.Pointer
	_CFMachPortGetPort func(unsafe.Pointer) unsafe.Pointer
	_CFMachPortGetTypeID func() unsafe.Pointer
	_CFMachPortInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CFMachPortIsValid func(unsafe.Pointer) unsafe.Pointer
	_CFMachPortSetInvalidationCallBack func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMessagePortCreateLocal func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMessagePortCreateRemote func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMessagePortCreateRunLoopSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMessagePortGetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMessagePortGetInvalidationCallBack func(unsafe.Pointer) unsafe.Pointer
	_CFMessagePortGetName func(unsafe.Pointer) unsafe.Pointer
	_CFMessagePortGetTypeID func() unsafe.Pointer
	_CFMessagePortInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CFMessagePortIsRemote func(unsafe.Pointer) unsafe.Pointer
	_CFMessagePortIsValid func(unsafe.Pointer) unsafe.Pointer
	_CFMessagePortSendRequest func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMessagePortSetDispatchQueue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMessagePortSetInvalidationCallBack func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFMessagePortSetName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNotificationCenterAddObserver func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNotificationCenterGetDarwinNotifyCenter func() unsafe.Pointer
	_CFNotificationCenterGetDistributedCenter func() unsafe.Pointer
	_CFNotificationCenterGetLocalCenter func() unsafe.Pointer
	_CFNotificationCenterGetTypeID func() unsafe.Pointer
	_CFNotificationCenterPostNotification func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNotificationCenterPostNotificationWithOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNotificationCenterRemoveEveryObserver func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNotificationCenterRemoveObserver func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNullGetTypeID func() unsafe.Pointer
	_CFNumberCompare func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterCopyProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterCreateNumberFromString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterCreateStringWithNumber func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterCreateStringWithValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterGetDecimalInfoForCurrencyCode func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterGetFormat func(unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterGetLocale func(unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterGetStyle func(unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterGetTypeID func() unsafe.Pointer
	_CFNumberFormatterGetValueFromString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterSetFormat func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberGetByteSize func(unsafe.Pointer) unsafe.Pointer
	_CFNumberGetType func(unsafe.Pointer) unsafe.Pointer
	_CFNumberGetTypeID func() unsafe.Pointer
	_CFNumberGetValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFNumberIsFloatType func(unsafe.Pointer) unsafe.Pointer
	_CFPlugInAddInstanceForFactory func(unsafe.Pointer) unsafe.Pointer
	_CFPlugInCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPlugInFindFactoriesForPlugInType func(unsafe.Pointer) unsafe.Pointer
	_CFPlugInFindFactoriesForPlugInTypeInPlugIn func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPlugInGetBundle func(unsafe.Pointer) unsafe.Pointer
	_CFPlugInGetTypeID func() unsafe.Pointer
	_CFPlugInInstanceCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPlugInInstanceCreateWithInstanceDataSize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPlugInInstanceGetFactoryName func(unsafe.Pointer) unsafe.Pointer
	_CFPlugInInstanceGetInstanceData func(unsafe.Pointer) unsafe.Pointer
	_CFPlugInInstanceGetInterfaceFunctionTable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPlugInInstanceGetTypeID func() unsafe.Pointer
	_CFPlugInIsLoadOnDemand func(unsafe.Pointer) unsafe.Pointer
	_CFPlugInRegisterFactoryFunction func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPlugInRegisterFactoryFunctionByName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPlugInRegisterPlugInType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPlugInRemoveInstanceForFactory func(unsafe.Pointer) unsafe.Pointer
	_CFPlugInSetLoadOnDemand func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPlugInUnregisterFactory func(unsafe.Pointer) unsafe.Pointer
	_CFPlugInUnregisterPlugInType func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesAddSuitePreferencesToApp func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesAppSynchronize func(unsafe.Pointer) unsafe.Pointer
	_CFPreferencesAppValueIsForced func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesCopyAppValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesCopyApplicationList func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesCopyKeyList func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesCopyMultiple func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesCopyValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesGetAppBooleanValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesGetAppIntegerValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesRemoveSuitePreferencesFromApp func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesSetAppValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesSetMultiple func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesSetValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesSynchronize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPropertyListCreateData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPropertyListCreateDeepCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPropertyListCreateFromStream func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPropertyListCreateFromXMLData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPropertyListCreateWithData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPropertyListCreateWithStream func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPropertyListCreateXMLData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPropertyListIsValid func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPropertyListWrite func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFPropertyListWriteToStream func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamClose func(unsafe.Pointer) unsafe.Pointer
	_CFReadStreamCopyDispatchQueue func(unsafe.Pointer) unsafe.Pointer
	_CFReadStreamCopyError func(unsafe.Pointer) unsafe.Pointer
	_CFReadStreamCopyProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamCreateWithBytesNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamCreateWithFile func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamGetBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamGetError func(unsafe.Pointer) unsafe.Pointer
	_CFReadStreamGetStatus func(unsafe.Pointer) unsafe.Pointer
	_CFReadStreamGetTypeID func() unsafe.Pointer
	_CFReadStreamHasBytesAvailable func(unsafe.Pointer) unsafe.Pointer
	_CFReadStreamOpen func(unsafe.Pointer) unsafe.Pointer
	_CFReadStreamRead func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamScheduleWithRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamSetClient func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamSetDispatchQueue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamUnscheduleFromRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopAddCommonMode func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopAddObserver func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopAddSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopAddTimer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopContainsObserver func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopContainsSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopContainsTimer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopCopyAllModes func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopCopyCurrentMode func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopGetCurrent func() unsafe.Pointer
	_CFRunLoopGetMain func() unsafe.Pointer
	_CFRunLoopGetNextTimerFireDate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopGetTypeID func() unsafe.Pointer
	_CFRunLoopIsWaiting func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopObserverCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopObserverCreateWithHandler func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopObserverDoesRepeat func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopObserverGetActivities func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopObserverGetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopObserverGetOrder func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopObserverGetTypeID func() unsafe.Pointer
	_CFRunLoopObserverInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopObserverIsValid func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopPerformBlock func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopRemoveObserver func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopRemoveSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopRemoveTimer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopRun func() unsafe.Pointer
	_CFRunLoopRunInMode func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopSourceCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopSourceGetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopSourceGetOrder func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopSourceGetTypeID func() unsafe.Pointer
	_CFRunLoopSourceInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopSourceIsValid func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopSourceSignal func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopStop func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerCreateWithHandler func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerDoesRepeat func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerGetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerGetInterval func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerGetNextFireDate func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerGetOrder func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerGetTolerance func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerGetTypeID func() unsafe.Pointer
	_CFRunLoopTimerInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerIsValid func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerSetNextFireDate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopTimerSetTolerance func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFRunLoopWakeUp func(unsafe.Pointer) unsafe.Pointer
	_CFSetAddValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetApplyFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetContainsValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetCreateMutable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFSetGetCountOfValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetGetTypeID func() unsafe.Pointer
	_CFSetGetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetGetValueIfPresent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetGetValues func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetRemoveAllValues func(unsafe.Pointer) unsafe.Pointer
	_CFSetRemoveValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetReplaceValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetSetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFShow func(unsafe.Pointer) unsafe.Pointer
	_CFShowStr func(unsafe.Pointer) unsafe.Pointer
	_CFSocketConnectToAddress func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketCopyAddress func(unsafe.Pointer) unsafe.Pointer
	_CFSocketCopyPeerAddress func(unsafe.Pointer) unsafe.Pointer
	_CFSocketCopyRegisteredSocketSignature func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketCopyRegisteredValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketCreateConnectedToSocketSignature func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketCreateRunLoopSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketCreateWithNative func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketCreateWithSocketSignature func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketDisableCallBacks func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketEnableCallBacks func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketGetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketGetDefaultNameRegistryPortNumber func() unsafe.Pointer
	_CFSocketGetNative func(unsafe.Pointer) unsafe.Pointer
	_CFSocketGetSocketFlags func(unsafe.Pointer) unsafe.Pointer
	_CFSocketGetTypeID func() unsafe.Pointer
	_CFSocketInvalidate func(unsafe.Pointer) unsafe.Pointer
	_CFSocketIsValid func(unsafe.Pointer) unsafe.Pointer
	_CFSocketRegisterSocketSignature func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketRegisterValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketSendData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketSetAddress func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketSetDefaultNameRegistryPortNumber func(unsafe.Pointer) unsafe.Pointer
	_CFSocketSetSocketFlags func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSocketUnregister func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStreamCreateBoundPair func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStreamCreatePairWithPeerSocketSignature func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStreamCreatePairWithSocket func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStreamCreatePairWithSocketToHost func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringAppend func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringAppendCString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringAppendCharacters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringAppendFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringAppendFormatAndArguments func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringAppendPascalString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCapitalize func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCompare func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCompareWithOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCompareWithOptionsAndLocale func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringConvertEncodingToIANACharSetName func(unsafe.Pointer) unsafe.Pointer
	_CFStringConvertEncodingToNSStringEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringConvertEncodingToWindowsCodepage func(unsafe.Pointer) unsafe.Pointer
	_CFStringConvertIANACharSetNameToEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringConvertNSStringEncodingToEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringConvertWindowsCodepageToEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringCreateArrayBySeparatingStrings func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateArrayWithFindResults func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateByCombiningStrings func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateExternalRepresentation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateFromExternalRepresentation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateMutable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateMutableWithExternalCharactersNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateStringWithValidatedFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateStringWithValidatedFormatAndArguments func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithBytesNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithCString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithCStringNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithCharacters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithCharactersNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithFileSystemRepresentation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithFormat func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithFormatAndArguments func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithPascalString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithPascalStringNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringCreateWithSubstring func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringDelete func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringFind func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringFindAndReplace func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringFindCharacterFromSet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringFindWithOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringFindWithOptionsAndLocale func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringFold func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetCString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetCStringPtr func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetCharacterAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetCharacters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetCharactersPtr func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetDoubleValue func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetFastestEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetFileSystemRepresentation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetHyphenationLocationBeforeIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetIntValue func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetLength func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetLineBounds func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetListOfAvailableEncodings func() unsafe.Pointer
	_CFStringGetMaximumSizeForEncoding func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetMaximumSizeOfFileSystemRepresentation func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetMostCompatibleMacStringEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetNameOfEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetParagraphBounds func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetPascalString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetPascalStringPtr func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetRangeOfComposedCharactersAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetSmallestEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetSystemEncoding func() unsafe.Pointer
	_CFStringGetTypeID func() unsafe.Pointer
	_CFStringHasPrefix func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringHasSuffix func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringInsert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringIsEncodingAvailable func(unsafe.Pointer) unsafe.Pointer
	_CFStringIsHyphenationAvailableForLocale func(unsafe.Pointer) unsafe.Pointer
	_CFStringLowercase func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringNormalize func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringPad func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringReplace func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringReplaceAll func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringSetExternalCharactersNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringTokenizerAdvanceToNextToken func(unsafe.Pointer) unsafe.Pointer
	_CFStringTokenizerCopyBestStringLanguage func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringTokenizerCopyCurrentTokenAttribute func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringTokenizerCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringTokenizerGetCurrentSubTokens func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringTokenizerGetCurrentTokenRange func(unsafe.Pointer) unsafe.Pointer
	_CFStringTokenizerGetTypeID func() unsafe.Pointer
	_CFStringTokenizerGoToTokenAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringTokenizerSetString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringTransform func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringTrim func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringTrimWhitespace func(unsafe.Pointer) unsafe.Pointer
	_CFStringUppercase func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneCopyAbbreviation func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneCopyAbbreviationDictionary func() unsafe.Pointer
	_CFTimeZoneCopyDefault func() unsafe.Pointer
	_CFTimeZoneCopyKnownNames func() unsafe.Pointer
	_CFTimeZoneCopyLocalizedName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneCopySystem func() unsafe.Pointer
	_CFTimeZoneCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneCreateWithName func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneCreateWithTimeIntervalFromGMT func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneGetData func(unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneGetDaylightSavingTimeOffset func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneGetName func(unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneGetNextDaylightSavingTimeTransition func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneGetSecondsFromGMT func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneGetTypeID func() unsafe.Pointer
	_CFTimeZoneIsDaylightSavingTime func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneResetSystem func() unsafe.Pointer
	_CFTimeZoneSetAbbreviationDictionary func(unsafe.Pointer) unsafe.Pointer
	_CFTimeZoneSetDefault func(unsafe.Pointer) unsafe.Pointer
	_CFTreeAppendChild func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTreeApplyFunctionToChildren func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTreeCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTreeFindRoot func(unsafe.Pointer) unsafe.Pointer
	_CFTreeGetChildAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTreeGetChildCount func(unsafe.Pointer) unsafe.Pointer
	_CFTreeGetChildren func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTreeGetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTreeGetFirstChild func(unsafe.Pointer) unsafe.Pointer
	_CFTreeGetNextSibling func(unsafe.Pointer) unsafe.Pointer
	_CFTreeGetParent func(unsafe.Pointer) unsafe.Pointer
	_CFTreeGetTypeID func() unsafe.Pointer
	_CFTreeInsertSibling func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTreePrependChild func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTreeRemove func(unsafe.Pointer) unsafe.Pointer
	_CFTreeRemoveAllChildren func(unsafe.Pointer) unsafe.Pointer
	_CFTreeSetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFTreeSortChildren func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCanBeDecomposed func(unsafe.Pointer) unsafe.Pointer
	_CFURLClearResourcePropertyCache func(unsafe.Pointer) unsafe.Pointer
	_CFURLClearResourcePropertyCacheForKey func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCopyAbsoluteURL func(unsafe.Pointer) unsafe.Pointer
	_CFURLCopyFileSystemPath func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCopyFragment func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCopyHostName func(unsafe.Pointer) unsafe.Pointer
	_CFURLCopyLastPathComponent func(unsafe.Pointer) unsafe.Pointer
	_CFURLCopyNetLocation func(unsafe.Pointer) unsafe.Pointer
	_CFURLCopyParameterString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCopyPassword func(unsafe.Pointer) unsafe.Pointer
	_CFURLCopyPath func(unsafe.Pointer) unsafe.Pointer
	_CFURLCopyPathExtension func(unsafe.Pointer) unsafe.Pointer
	_CFURLCopyQueryString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCopyResourcePropertiesForKeys func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCopyResourcePropertyForKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCopyResourceSpecifier func(unsafe.Pointer) unsafe.Pointer
	_CFURLCopyScheme func(unsafe.Pointer) unsafe.Pointer
	_CFURLCopyStrictPath func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCopyUserName func(unsafe.Pointer) unsafe.Pointer
	_CFURLCreateAbsoluteURLWithBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateBookmarkData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateBookmarkDataFromAliasRecord func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateBookmarkDataFromFile func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateByResolvingBookmarkData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateCopyAppendingPathComponent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateCopyAppendingPathExtension func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateCopyDeletingLastPathComponent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateCopyDeletingPathExtension func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateDataAndPropertiesFromResource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateFilePathURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateFileReferenceURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateFromFSRef func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateFromFileSystemRepresentation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateFromFileSystemRepresentationRelativeToBase func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreatePropertyFromResource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateResourcePropertiesForKeysFromBookmarkData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateResourcePropertyForKeyFromBookmarkData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateStringByAddingPercentEscapes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateStringByReplacingPercentEscapes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateStringByReplacingPercentEscapesUsingEncoding func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateWithBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateWithFileSystemPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateWithFileSystemPathRelativeToBase func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateWithString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLDestroyResource func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLEnumeratorCreateForDirectoryURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLEnumeratorCreateForMountedVolumes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLEnumeratorGetDescendentLevel func(unsafe.Pointer) unsafe.Pointer
	_CFURLEnumeratorGetNextURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLEnumeratorGetSourceDidChange func(unsafe.Pointer) unsafe.Pointer
	_CFURLEnumeratorGetTypeID func() unsafe.Pointer
	_CFURLEnumeratorSkipDescendents func(unsafe.Pointer) unsafe.Pointer
	_CFURLGetBaseURL func(unsafe.Pointer) unsafe.Pointer
	_CFURLGetByteRangeForComponent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLGetBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLGetFSRef func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLGetFileSystemRepresentation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLGetPortNumber func(unsafe.Pointer) unsafe.Pointer
	_CFURLGetString func(unsafe.Pointer) unsafe.Pointer
	_CFURLGetTypeID func() unsafe.Pointer
	_CFURLHasDirectoryPath func(unsafe.Pointer) unsafe.Pointer
	_CFURLIsFileReferenceURL func(unsafe.Pointer) unsafe.Pointer
	_CFURLResourceIsReachable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLSetResourcePropertiesForKeys func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLSetResourcePropertyForKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLSetTemporaryResourcePropertyForKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLStartAccessingSecurityScopedResource func(unsafe.Pointer) unsafe.Pointer
	_CFURLStopAccessingSecurityScopedResource func(unsafe.Pointer) unsafe.Pointer
	_CFURLWriteBookmarkDataToFile func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLWriteDataAndPropertiesToResource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDCreate func(unsafe.Pointer) unsafe.Pointer
	_CFUUIDCreateFromString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDCreateFromUUIDBytes func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDCreateString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDCreateWithBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDGetConstantUUIDWithBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDGetTypeID func() unsafe.Pointer
	_CFUUIDGetUUIDBytes func(unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationCancel func(unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationCreateRunLoopSource func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationDisplayAlert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationDisplayNotice func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationGetResponseDictionary func(unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationGetResponseValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationGetTypeID func() unsafe.Pointer
	_CFUserNotificationReceiveResponse func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationUpdate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCanAcceptBytes func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamClose func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCopyDispatchQueue func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCopyError func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCopyProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCreateWithAllocatedBuffers func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCreateWithBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCreateWithFile func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamGetError func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamGetStatus func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamGetTypeID func() unsafe.Pointer
	_CFWriteStreamOpen func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamScheduleWithRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamSetClient func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamSetDispatchQueue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamUnscheduleFromRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamWrite func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLCreateStringByEscapingEntities func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLCreateStringByUnescapingEntities func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLNodeCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLNodeCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLNodeGetInfoPtr func(unsafe.Pointer) unsafe.Pointer
	_CFXMLNodeGetString func(unsafe.Pointer) unsafe.Pointer
	_CFXMLNodeGetTypeCode func(unsafe.Pointer) unsafe.Pointer
	_CFXMLNodeGetTypeID func() unsafe.Pointer
	_CFXMLNodeGetVersion func(unsafe.Pointer) unsafe.Pointer
	_CFXMLParserAbort func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLParserCopyErrorDescription func(unsafe.Pointer) unsafe.Pointer
	_CFXMLParserCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLParserCreateWithDataFromURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLParserGetCallBacks func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLParserGetContext func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLParserGetDocument func(unsafe.Pointer) unsafe.Pointer
	_CFXMLParserGetLineNumber func(unsafe.Pointer) unsafe.Pointer
	_CFXMLParserGetLocation func(unsafe.Pointer) unsafe.Pointer
	_CFXMLParserGetSourceURL func(unsafe.Pointer) unsafe.Pointer
	_CFXMLParserGetStatusCode func(unsafe.Pointer) unsafe.Pointer
	_CFXMLParserGetTypeID func() unsafe.Pointer
	_CFXMLParserParse func(unsafe.Pointer) unsafe.Pointer
	_CFXMLTreeCreateFromData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLTreeCreateFromDataWithError func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLTreeCreateWithDataFromURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLTreeCreateWithNode func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLTreeCreateXMLData func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFXMLTreeGetNode func(unsafe.Pointer) unsafe.Pointer
	_inset func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_CFRelease, lib, "CFRelease")
	tryRegister(&_CFMakeCollectable, lib, "CFMakeCollectable")
	tryRegister(&_CFRetain, lib, "CFRetain")
	tryRegister(&_CFAbsoluteTimeAddGregorianUnits, lib, "CFAbsoluteTimeAddGregorianUnits")
	tryRegister(&_CFAbsoluteTimeGetCurrent, lib, "CFAbsoluteTimeGetCurrent")
	tryRegister(&_CFAbsoluteTimeGetDayOfWeek, lib, "CFAbsoluteTimeGetDayOfWeek")
	tryRegister(&_CFAbsoluteTimeGetDayOfYear, lib, "CFAbsoluteTimeGetDayOfYear")
	tryRegister(&_CFAbsoluteTimeGetDifferenceAsGregorianUnits, lib, "CFAbsoluteTimeGetDifferenceAsGregorianUnits")
	tryRegister(&_CFAbsoluteTimeGetGregorianDate, lib, "CFAbsoluteTimeGetGregorianDate")
	tryRegister(&_CFAbsoluteTimeGetWeekOfYear, lib, "CFAbsoluteTimeGetWeekOfYear")
	tryRegister(&_CFAllocatorAllocate, lib, "CFAllocatorAllocate")
	tryRegister(&_CFAllocatorAllocateBytes, lib, "CFAllocatorAllocateBytes")
	tryRegister(&_CFAllocatorAllocateTyped, lib, "CFAllocatorAllocateTyped")
	tryRegister(&_CFAllocatorCreate, lib, "CFAllocatorCreate")
	tryRegister(&_CFAllocatorCreateWithZone, lib, "CFAllocatorCreateWithZone")
	tryRegister(&_CFAllocatorDeallocate, lib, "CFAllocatorDeallocate")
	tryRegister(&_CFAllocatorGetContext, lib, "CFAllocatorGetContext")
	tryRegister(&_CFAllocatorGetDefault, lib, "CFAllocatorGetDefault")
	tryRegister(&_CFAllocatorGetPreferredSizeForSize, lib, "CFAllocatorGetPreferredSizeForSize")
	tryRegister(&_CFAllocatorGetTypeID, lib, "CFAllocatorGetTypeID")
	tryRegister(&_CFAllocatorReallocate, lib, "CFAllocatorReallocate")
	tryRegister(&_CFAllocatorReallocateBytes, lib, "CFAllocatorReallocateBytes")
	tryRegister(&_CFAllocatorReallocateTyped, lib, "CFAllocatorReallocateTyped")
	tryRegister(&_CFAllocatorSetDefault, lib, "CFAllocatorSetDefault")
	tryRegister(&_CFArrayAppendArray, lib, "CFArrayAppendArray")
	tryRegister(&_CFArrayAppendValue, lib, "CFArrayAppendValue")
	tryRegister(&_CFArrayApplyFunction, lib, "CFArrayApplyFunction")
	tryRegister(&_CFArrayBSearchValues, lib, "CFArrayBSearchValues")
	tryRegister(&_CFArrayContainsValue, lib, "CFArrayContainsValue")
	tryRegister(&_CFArrayCreate, lib, "CFArrayCreate")
	tryRegister(&_CFArrayCreateCopy, lib, "CFArrayCreateCopy")
	tryRegister(&_CFArrayCreateMutable, lib, "CFArrayCreateMutable")
	tryRegister(&_CFArrayCreateMutableCopy, lib, "CFArrayCreateMutableCopy")
	tryRegister(&_CFArrayExchangeValuesAtIndices, lib, "CFArrayExchangeValuesAtIndices")
	tryRegister(&_CFArrayGetCount, lib, "CFArrayGetCount")
	tryRegister(&_CFArrayGetCountOfValue, lib, "CFArrayGetCountOfValue")
	tryRegister(&_CFArrayGetFirstIndexOfValue, lib, "CFArrayGetFirstIndexOfValue")
	tryRegister(&_CFArrayGetLastIndexOfValue, lib, "CFArrayGetLastIndexOfValue")
	tryRegister(&_CFArrayGetTypeID, lib, "CFArrayGetTypeID")
	tryRegister(&_CFArrayGetValueAtIndex, lib, "CFArrayGetValueAtIndex")
	tryRegister(&_CFArrayGetValues, lib, "CFArrayGetValues")
	tryRegister(&_CFArrayInsertValueAtIndex, lib, "CFArrayInsertValueAtIndex")
	tryRegister(&_CFArrayRemoveAllValues, lib, "CFArrayRemoveAllValues")
	tryRegister(&_CFArrayRemoveValueAtIndex, lib, "CFArrayRemoveValueAtIndex")
	tryRegister(&_CFArrayReplaceValues, lib, "CFArrayReplaceValues")
	tryRegister(&_CFArraySetValueAtIndex, lib, "CFArraySetValueAtIndex")
	tryRegister(&_CFArraySortValues, lib, "CFArraySortValues")
	tryRegister(&_CFAttributedStringBeginEditing, lib, "CFAttributedStringBeginEditing")
	tryRegister(&_CFAttributedStringCreate, lib, "CFAttributedStringCreate")
	tryRegister(&_CFAttributedStringCreateCopy, lib, "CFAttributedStringCreateCopy")
	tryRegister(&_CFAttributedStringCreateMutable, lib, "CFAttributedStringCreateMutable")
	tryRegister(&_CFAttributedStringCreateMutableCopy, lib, "CFAttributedStringCreateMutableCopy")
	tryRegister(&_CFAttributedStringCreateWithSubstring, lib, "CFAttributedStringCreateWithSubstring")
	tryRegister(&_CFAttributedStringEndEditing, lib, "CFAttributedStringEndEditing")
	tryRegister(&_CFAttributedStringGetAttribute, lib, "CFAttributedStringGetAttribute")
	tryRegister(&_CFAttributedStringGetAttributeAndLongestEffectiveRange, lib, "CFAttributedStringGetAttributeAndLongestEffectiveRange")
	tryRegister(&_CFAttributedStringGetAttributes, lib, "CFAttributedStringGetAttributes")
	tryRegister(&_CFAttributedStringGetAttributesAndLongestEffectiveRange, lib, "CFAttributedStringGetAttributesAndLongestEffectiveRange")
	tryRegister(&_CFAttributedStringGetBidiLevelsAndResolvedDirections, lib, "CFAttributedStringGetBidiLevelsAndResolvedDirections")
	tryRegister(&_CFAttributedStringGetLength, lib, "CFAttributedStringGetLength")
	tryRegister(&_CFAttributedStringGetMutableString, lib, "CFAttributedStringGetMutableString")
	tryRegister(&_CFAttributedStringGetStatisticalWritingDirections, lib, "CFAttributedStringGetStatisticalWritingDirections")
	tryRegister(&_CFAttributedStringGetString, lib, "CFAttributedStringGetString")
	tryRegister(&_CFAttributedStringGetTypeID, lib, "CFAttributedStringGetTypeID")
	tryRegister(&_CFAttributedStringRemoveAttribute, lib, "CFAttributedStringRemoveAttribute")
	tryRegister(&_CFAttributedStringReplaceAttributedString, lib, "CFAttributedStringReplaceAttributedString")
	tryRegister(&_CFAttributedStringReplaceString, lib, "CFAttributedStringReplaceString")
	tryRegister(&_CFAttributedStringSetAttribute, lib, "CFAttributedStringSetAttribute")
	tryRegister(&_CFAttributedStringSetAttributes, lib, "CFAttributedStringSetAttributes")
	tryRegister(&_CFAutorelease, lib, "CFAutorelease")
	tryRegister(&_CFBagAddValue, lib, "CFBagAddValue")
	tryRegister(&_CFBagApplyFunction, lib, "CFBagApplyFunction")
	tryRegister(&_CFBagContainsValue, lib, "CFBagContainsValue")
	tryRegister(&_CFBagCreate, lib, "CFBagCreate")
	tryRegister(&_CFBagCreateCopy, lib, "CFBagCreateCopy")
	tryRegister(&_CFBagCreateMutable, lib, "CFBagCreateMutable")
	tryRegister(&_CFBagCreateMutableCopy, lib, "CFBagCreateMutableCopy")
	tryRegister(&_CFBagGetCount, lib, "CFBagGetCount")
	tryRegister(&_CFBagGetCountOfValue, lib, "CFBagGetCountOfValue")
	tryRegister(&_CFBagGetTypeID, lib, "CFBagGetTypeID")
	tryRegister(&_CFBagGetValue, lib, "CFBagGetValue")
	tryRegister(&_CFBagGetValueIfPresent, lib, "CFBagGetValueIfPresent")
	tryRegister(&_CFBagGetValues, lib, "CFBagGetValues")
	tryRegister(&_CFBagRemoveAllValues, lib, "CFBagRemoveAllValues")
	tryRegister(&_CFBagRemoveValue, lib, "CFBagRemoveValue")
	tryRegister(&_CFBagReplaceValue, lib, "CFBagReplaceValue")
	tryRegister(&_CFBagSetValue, lib, "CFBagSetValue")
	tryRegister(&_CFBinaryHeapAddValue, lib, "CFBinaryHeapAddValue")
	tryRegister(&_CFBinaryHeapApplyFunction, lib, "CFBinaryHeapApplyFunction")
	tryRegister(&_CFBinaryHeapContainsValue, lib, "CFBinaryHeapContainsValue")
	tryRegister(&_CFBinaryHeapCreate, lib, "CFBinaryHeapCreate")
	tryRegister(&_CFBinaryHeapCreateCopy, lib, "CFBinaryHeapCreateCopy")
	tryRegister(&_CFBinaryHeapGetCount, lib, "CFBinaryHeapGetCount")
	tryRegister(&_CFBinaryHeapGetCountOfValue, lib, "CFBinaryHeapGetCountOfValue")
	tryRegister(&_CFBinaryHeapGetMinimum, lib, "CFBinaryHeapGetMinimum")
	tryRegister(&_CFBinaryHeapGetMinimumIfPresent, lib, "CFBinaryHeapGetMinimumIfPresent")
	tryRegister(&_CFBinaryHeapGetTypeID, lib, "CFBinaryHeapGetTypeID")
	tryRegister(&_CFBinaryHeapGetValues, lib, "CFBinaryHeapGetValues")
	tryRegister(&_CFBinaryHeapRemoveAllValues, lib, "CFBinaryHeapRemoveAllValues")
	tryRegister(&_CFBinaryHeapRemoveMinimumValue, lib, "CFBinaryHeapRemoveMinimumValue")
	tryRegister(&_CFBitVectorContainsBit, lib, "CFBitVectorContainsBit")
	tryRegister(&_CFBitVectorCreate, lib, "CFBitVectorCreate")
	tryRegister(&_CFBitVectorCreateCopy, lib, "CFBitVectorCreateCopy")
	tryRegister(&_CFBitVectorCreateMutable, lib, "CFBitVectorCreateMutable")
	tryRegister(&_CFBitVectorCreateMutableCopy, lib, "CFBitVectorCreateMutableCopy")
	tryRegister(&_CFBitVectorFlipBitAtIndex, lib, "CFBitVectorFlipBitAtIndex")
	tryRegister(&_CFBitVectorFlipBits, lib, "CFBitVectorFlipBits")
	tryRegister(&_CFBitVectorGetBitAtIndex, lib, "CFBitVectorGetBitAtIndex")
	tryRegister(&_CFBitVectorGetBits, lib, "CFBitVectorGetBits")
	tryRegister(&_CFBitVectorGetCount, lib, "CFBitVectorGetCount")
	tryRegister(&_CFBitVectorGetCountOfBit, lib, "CFBitVectorGetCountOfBit")
	tryRegister(&_CFBitVectorGetFirstIndexOfBit, lib, "CFBitVectorGetFirstIndexOfBit")
	tryRegister(&_CFBitVectorGetLastIndexOfBit, lib, "CFBitVectorGetLastIndexOfBit")
	tryRegister(&_CFBitVectorGetTypeID, lib, "CFBitVectorGetTypeID")
	tryRegister(&_CFBitVectorSetAllBits, lib, "CFBitVectorSetAllBits")
	tryRegister(&_CFBitVectorSetBitAtIndex, lib, "CFBitVectorSetBitAtIndex")
	tryRegister(&_CFBitVectorSetBits, lib, "CFBitVectorSetBits")
	tryRegister(&_CFBitVectorSetCount, lib, "CFBitVectorSetCount")
	tryRegister(&_CFBooleanGetTypeID, lib, "CFBooleanGetTypeID")
	tryRegister(&_CFBooleanGetValue, lib, "CFBooleanGetValue")
	tryRegister(&_CFBundleCloseBundleResourceMap, lib, "CFBundleCloseBundleResourceMap")
	tryRegister(&_CFBundleCopyAuxiliaryExecutableURL, lib, "CFBundleCopyAuxiliaryExecutableURL")
	tryRegister(&_CFBundleCopyBuiltInPlugInsURL, lib, "CFBundleCopyBuiltInPlugInsURL")
	tryRegister(&_CFBundleCopyBundleLocalizations, lib, "CFBundleCopyBundleLocalizations")
	tryRegister(&_CFBundleCopyBundleURL, lib, "CFBundleCopyBundleURL")
	tryRegister(&_CFBundleCopyExecutableArchitectures, lib, "CFBundleCopyExecutableArchitectures")
	tryRegister(&_CFBundleCopyExecutableArchitecturesForURL, lib, "CFBundleCopyExecutableArchitecturesForURL")
	tryRegister(&_CFBundleCopyExecutableURL, lib, "CFBundleCopyExecutableURL")
	tryRegister(&_CFBundleCopyInfoDictionaryForURL, lib, "CFBundleCopyInfoDictionaryForURL")
	tryRegister(&_CFBundleCopyInfoDictionaryInDirectory, lib, "CFBundleCopyInfoDictionaryInDirectory")
	tryRegister(&_CFBundleCopyLocalizationsForPreferences, lib, "CFBundleCopyLocalizationsForPreferences")
	tryRegister(&_CFBundleCopyLocalizationsForURL, lib, "CFBundleCopyLocalizationsForURL")
	tryRegister(&_CFBundleCopyLocalizedString, lib, "CFBundleCopyLocalizedString")
	tryRegister(&_CFBundleCopyLocalizedStringForLocalizations, lib, "CFBundleCopyLocalizedStringForLocalizations")
	tryRegister(&_CFBundleCopyPreferredLocalizationsFromArray, lib, "CFBundleCopyPreferredLocalizationsFromArray")
	tryRegister(&_CFBundleCopyPrivateFrameworksURL, lib, "CFBundleCopyPrivateFrameworksURL")
	tryRegister(&_CFBundleCopyResourceURL, lib, "CFBundleCopyResourceURL")
	tryRegister(&_CFBundleCopyResourceURLForLocalization, lib, "CFBundleCopyResourceURLForLocalization")
	tryRegister(&_CFBundleCopyResourceURLInDirectory, lib, "CFBundleCopyResourceURLInDirectory")
	tryRegister(&_CFBundleCopyResourceURLsOfType, lib, "CFBundleCopyResourceURLsOfType")
	tryRegister(&_CFBundleCopyResourceURLsOfTypeForLocalization, lib, "CFBundleCopyResourceURLsOfTypeForLocalization")
	tryRegister(&_CFBundleCopyResourceURLsOfTypeInDirectory, lib, "CFBundleCopyResourceURLsOfTypeInDirectory")
	tryRegister(&_CFBundleCopyResourcesDirectoryURL, lib, "CFBundleCopyResourcesDirectoryURL")
	tryRegister(&_CFBundleCopySharedFrameworksURL, lib, "CFBundleCopySharedFrameworksURL")
	tryRegister(&_CFBundleCopySharedSupportURL, lib, "CFBundleCopySharedSupportURL")
	tryRegister(&_CFBundleCopySupportFilesDirectoryURL, lib, "CFBundleCopySupportFilesDirectoryURL")
	tryRegister(&_CFBundleCreate, lib, "CFBundleCreate")
	tryRegister(&_CFBundleCreateBundlesFromDirectory, lib, "CFBundleCreateBundlesFromDirectory")
	tryRegister(&_CFBundleGetAllBundles, lib, "CFBundleGetAllBundles")
	tryRegister(&_CFBundleGetBundleWithIdentifier, lib, "CFBundleGetBundleWithIdentifier")
	tryRegister(&_CFBundleGetDataPointerForName, lib, "CFBundleGetDataPointerForName")
	tryRegister(&_CFBundleGetDataPointersForNames, lib, "CFBundleGetDataPointersForNames")
	tryRegister(&_CFBundleGetDevelopmentRegion, lib, "CFBundleGetDevelopmentRegion")
	tryRegister(&_CFBundleGetFunctionPointerForName, lib, "CFBundleGetFunctionPointerForName")
	tryRegister(&_CFBundleGetFunctionPointersForNames, lib, "CFBundleGetFunctionPointersForNames")
	tryRegister(&_CFBundleGetIdentifier, lib, "CFBundleGetIdentifier")
	tryRegister(&_CFBundleGetInfoDictionary, lib, "CFBundleGetInfoDictionary")
	tryRegister(&_CFBundleGetLocalInfoDictionary, lib, "CFBundleGetLocalInfoDictionary")
	tryRegister(&_CFBundleGetMainBundle, lib, "CFBundleGetMainBundle")
	tryRegister(&_CFBundleGetPackageInfo, lib, "CFBundleGetPackageInfo")
	tryRegister(&_CFBundleGetPackageInfoInDirectory, lib, "CFBundleGetPackageInfoInDirectory")
	tryRegister(&_CFBundleGetPlugIn, lib, "CFBundleGetPlugIn")
	tryRegister(&_CFBundleGetTypeID, lib, "CFBundleGetTypeID")
	tryRegister(&_CFBundleGetValueForInfoDictionaryKey, lib, "CFBundleGetValueForInfoDictionaryKey")
	tryRegister(&_CFBundleGetVersionNumber, lib, "CFBundleGetVersionNumber")
	tryRegister(&_CFBundleIsArchitectureLoadable, lib, "CFBundleIsArchitectureLoadable")
	tryRegister(&_CFBundleIsExecutableLoadable, lib, "CFBundleIsExecutableLoadable")
	tryRegister(&_CFBundleIsExecutableLoadableForURL, lib, "CFBundleIsExecutableLoadableForURL")
	tryRegister(&_CFBundleIsExecutableLoaded, lib, "CFBundleIsExecutableLoaded")
	tryRegister(&_CFBundleLoadExecutable, lib, "CFBundleLoadExecutable")
	tryRegister(&_CFBundleLoadExecutableAndReturnError, lib, "CFBundleLoadExecutableAndReturnError")
	tryRegister(&_CFBundleOpenBundleResourceFiles, lib, "CFBundleOpenBundleResourceFiles")
	tryRegister(&_CFBundleOpenBundleResourceMap, lib, "CFBundleOpenBundleResourceMap")
	tryRegister(&_CFBundlePreflightExecutable, lib, "CFBundlePreflightExecutable")
	tryRegister(&_CFBundleUnloadExecutable, lib, "CFBundleUnloadExecutable")
	tryRegister(&_CFCalendarAddComponents, lib, "CFCalendarAddComponents")
	tryRegister(&_CFCalendarComposeAbsoluteTime, lib, "CFCalendarComposeAbsoluteTime")
	tryRegister(&_CFCalendarCopyCurrent, lib, "CFCalendarCopyCurrent")
	tryRegister(&_CFCalendarCopyLocale, lib, "CFCalendarCopyLocale")
	tryRegister(&_CFCalendarCopyTimeZone, lib, "CFCalendarCopyTimeZone")
	tryRegister(&_CFCalendarCreateWithIdentifier, lib, "CFCalendarCreateWithIdentifier")
	tryRegister(&_CFCalendarDecomposeAbsoluteTime, lib, "CFCalendarDecomposeAbsoluteTime")
	tryRegister(&_CFCalendarGetComponentDifference, lib, "CFCalendarGetComponentDifference")
	tryRegister(&_CFCalendarGetFirstWeekday, lib, "CFCalendarGetFirstWeekday")
	tryRegister(&_CFCalendarGetIdentifier, lib, "CFCalendarGetIdentifier")
	tryRegister(&_CFCalendarGetMaximumRangeOfUnit, lib, "CFCalendarGetMaximumRangeOfUnit")
	tryRegister(&_CFCalendarGetMinimumDaysInFirstWeek, lib, "CFCalendarGetMinimumDaysInFirstWeek")
	tryRegister(&_CFCalendarGetMinimumRangeOfUnit, lib, "CFCalendarGetMinimumRangeOfUnit")
	tryRegister(&_CFCalendarGetOrdinalityOfUnit, lib, "CFCalendarGetOrdinalityOfUnit")
	tryRegister(&_CFCalendarGetRangeOfUnit, lib, "CFCalendarGetRangeOfUnit")
	tryRegister(&_CFCalendarGetTimeRangeOfUnit, lib, "CFCalendarGetTimeRangeOfUnit")
	tryRegister(&_CFCalendarGetTypeID, lib, "CFCalendarGetTypeID")
	tryRegister(&_CFCalendarSetFirstWeekday, lib, "CFCalendarSetFirstWeekday")
	tryRegister(&_CFCalendarSetLocale, lib, "CFCalendarSetLocale")
	tryRegister(&_CFCalendarSetMinimumDaysInFirstWeek, lib, "CFCalendarSetMinimumDaysInFirstWeek")
	tryRegister(&_CFCalendarSetTimeZone, lib, "CFCalendarSetTimeZone")
	tryRegister(&_CFCharacterSetAddCharactersInRange, lib, "CFCharacterSetAddCharactersInRange")
	tryRegister(&_CFCharacterSetAddCharactersInString, lib, "CFCharacterSetAddCharactersInString")
	tryRegister(&_CFCharacterSetCreateBitmapRepresentation, lib, "CFCharacterSetCreateBitmapRepresentation")
	tryRegister(&_CFCharacterSetCreateCopy, lib, "CFCharacterSetCreateCopy")
	tryRegister(&_CFCharacterSetCreateInvertedSet, lib, "CFCharacterSetCreateInvertedSet")
	tryRegister(&_CFCharacterSetCreateMutable, lib, "CFCharacterSetCreateMutable")
	tryRegister(&_CFCharacterSetCreateMutableCopy, lib, "CFCharacterSetCreateMutableCopy")
	tryRegister(&_CFCharacterSetCreateWithBitmapRepresentation, lib, "CFCharacterSetCreateWithBitmapRepresentation")
	tryRegister(&_CFCharacterSetCreateWithCharactersInRange, lib, "CFCharacterSetCreateWithCharactersInRange")
	tryRegister(&_CFCharacterSetCreateWithCharactersInString, lib, "CFCharacterSetCreateWithCharactersInString")
	tryRegister(&_CFCharacterSetGetPredefined, lib, "CFCharacterSetGetPredefined")
	tryRegister(&_CFCharacterSetGetTypeID, lib, "CFCharacterSetGetTypeID")
	tryRegister(&_CFCharacterSetHasMemberInPlane, lib, "CFCharacterSetHasMemberInPlane")
	tryRegister(&_CFCharacterSetIntersect, lib, "CFCharacterSetIntersect")
	tryRegister(&_CFCharacterSetInvert, lib, "CFCharacterSetInvert")
	tryRegister(&_CFCharacterSetIsCharacterMember, lib, "CFCharacterSetIsCharacterMember")
	tryRegister(&_CFCharacterSetIsLongCharacterMember, lib, "CFCharacterSetIsLongCharacterMember")
	tryRegister(&_CFCharacterSetIsSupersetOfSet, lib, "CFCharacterSetIsSupersetOfSet")
	tryRegister(&_CFCharacterSetRemoveCharactersInRange, lib, "CFCharacterSetRemoveCharactersInRange")
	tryRegister(&_CFCharacterSetRemoveCharactersInString, lib, "CFCharacterSetRemoveCharactersInString")
	tryRegister(&_CFCharacterSetUnion, lib, "CFCharacterSetUnion")
	tryRegister(&_CFCopyDescription, lib, "CFCopyDescription")
	tryRegister(&_CFCopyHomeDirectoryURL, lib, "CFCopyHomeDirectoryURL")
	tryRegister(&_CFCopyTypeIDDescription, lib, "CFCopyTypeIDDescription")
	tryRegister(&_CFDataAppendBytes, lib, "CFDataAppendBytes")
	tryRegister(&_CFDataCreate, lib, "CFDataCreate")
	tryRegister(&_CFDataCreateCopy, lib, "CFDataCreateCopy")
	tryRegister(&_CFDataCreateMutable, lib, "CFDataCreateMutable")
	tryRegister(&_CFDataCreateMutableCopy, lib, "CFDataCreateMutableCopy")
	tryRegister(&_CFDataCreateWithBytesNoCopy, lib, "CFDataCreateWithBytesNoCopy")
	tryRegister(&_CFDataDeleteBytes, lib, "CFDataDeleteBytes")
	tryRegister(&_CFDataFind, lib, "CFDataFind")
	tryRegister(&_CFDataGetBytePtr, lib, "CFDataGetBytePtr")
	tryRegister(&_CFDataGetBytes, lib, "CFDataGetBytes")
	tryRegister(&_CFDataGetLength, lib, "CFDataGetLength")
	tryRegister(&_CFDataGetMutableBytePtr, lib, "CFDataGetMutableBytePtr")
	tryRegister(&_CFDataGetTypeID, lib, "CFDataGetTypeID")
	tryRegister(&_CFDataIncreaseLength, lib, "CFDataIncreaseLength")
	tryRegister(&_CFDataReplaceBytes, lib, "CFDataReplaceBytes")
	tryRegister(&_CFDataSetLength, lib, "CFDataSetLength")
	tryRegister(&_CFDateCompare, lib, "CFDateCompare")
	tryRegister(&_CFDateCreate, lib, "CFDateCreate")
	tryRegister(&_CFDateFormatterCopyProperty, lib, "CFDateFormatterCopyProperty")
	tryRegister(&_CFDateFormatterCreate, lib, "CFDateFormatterCreate")
	tryRegister(&_CFDateFormatterCreateDateFormatFromTemplate, lib, "CFDateFormatterCreateDateFormatFromTemplate")
	tryRegister(&_CFDateFormatterCreateDateFromString, lib, "CFDateFormatterCreateDateFromString")
	tryRegister(&_CFDateFormatterCreateISO8601Formatter, lib, "CFDateFormatterCreateISO8601Formatter")
	tryRegister(&_CFDateFormatterCreateStringWithAbsoluteTime, lib, "CFDateFormatterCreateStringWithAbsoluteTime")
	tryRegister(&_CFDateFormatterCreateStringWithDate, lib, "CFDateFormatterCreateStringWithDate")
	tryRegister(&_CFDateFormatterGetAbsoluteTimeFromString, lib, "CFDateFormatterGetAbsoluteTimeFromString")
	tryRegister(&_CFDateFormatterGetDateStyle, lib, "CFDateFormatterGetDateStyle")
	tryRegister(&_CFDateFormatterGetFormat, lib, "CFDateFormatterGetFormat")
	tryRegister(&_CFDateFormatterGetLocale, lib, "CFDateFormatterGetLocale")
	tryRegister(&_CFDateFormatterGetTimeStyle, lib, "CFDateFormatterGetTimeStyle")
	tryRegister(&_CFDateFormatterGetTypeID, lib, "CFDateFormatterGetTypeID")
	tryRegister(&_CFDateFormatterSetFormat, lib, "CFDateFormatterSetFormat")
	tryRegister(&_CFDateFormatterSetProperty, lib, "CFDateFormatterSetProperty")
	tryRegister(&_CFDateGetAbsoluteTime, lib, "CFDateGetAbsoluteTime")
	tryRegister(&_CFDateGetTimeIntervalSinceDate, lib, "CFDateGetTimeIntervalSinceDate")
	tryRegister(&_CFDateGetTypeID, lib, "CFDateGetTypeID")
	tryRegister(&_CFDictionaryAddValue, lib, "CFDictionaryAddValue")
	tryRegister(&_CFDictionaryApplyFunction, lib, "CFDictionaryApplyFunction")
	tryRegister(&_CFDictionaryContainsKey, lib, "CFDictionaryContainsKey")
	tryRegister(&_CFDictionaryContainsValue, lib, "CFDictionaryContainsValue")
	tryRegister(&_CFDictionaryCreate, lib, "CFDictionaryCreate")
	tryRegister(&_CFDictionaryCreateCopy, lib, "CFDictionaryCreateCopy")
	tryRegister(&_CFDictionaryCreateMutable, lib, "CFDictionaryCreateMutable")
	tryRegister(&_CFDictionaryCreateMutableCopy, lib, "CFDictionaryCreateMutableCopy")
	tryRegister(&_CFDictionaryGetCount, lib, "CFDictionaryGetCount")
	tryRegister(&_CFDictionaryGetCountOfKey, lib, "CFDictionaryGetCountOfKey")
	tryRegister(&_CFDictionaryGetCountOfValue, lib, "CFDictionaryGetCountOfValue")
	tryRegister(&_CFDictionaryGetKeysAndValues, lib, "CFDictionaryGetKeysAndValues")
	tryRegister(&_CFDictionaryGetTypeID, lib, "CFDictionaryGetTypeID")
	tryRegister(&_CFDictionaryGetValue, lib, "CFDictionaryGetValue")
	tryRegister(&_CFDictionaryGetValueIfPresent, lib, "CFDictionaryGetValueIfPresent")
	tryRegister(&_CFDictionaryRemoveAllValues, lib, "CFDictionaryRemoveAllValues")
	tryRegister(&_CFDictionaryRemoveValue, lib, "CFDictionaryRemoveValue")
	tryRegister(&_CFDictionaryReplaceValue, lib, "CFDictionaryReplaceValue")
	tryRegister(&_CFDictionarySetValue, lib, "CFDictionarySetValue")
	tryRegister(&_CFEqual, lib, "CFEqual")
	tryRegister(&_CFErrorCopyDescription, lib, "CFErrorCopyDescription")
	tryRegister(&_CFErrorCopyFailureReason, lib, "CFErrorCopyFailureReason")
	tryRegister(&_CFErrorCopyRecoverySuggestion, lib, "CFErrorCopyRecoverySuggestion")
	tryRegister(&_CFErrorCopyUserInfo, lib, "CFErrorCopyUserInfo")
	tryRegister(&_CFErrorCreate, lib, "CFErrorCreate")
	tryRegister(&_CFErrorCreateWithUserInfoKeysAndValues, lib, "CFErrorCreateWithUserInfoKeysAndValues")
	tryRegister(&_CFErrorGetCode, lib, "CFErrorGetCode")
	tryRegister(&_CFErrorGetDomain, lib, "CFErrorGetDomain")
	tryRegister(&_CFErrorGetTypeID, lib, "CFErrorGetTypeID")
	tryRegister(&_CFFileDescriptorCreate, lib, "CFFileDescriptorCreate")
	tryRegister(&_CFFileDescriptorCreateRunLoopSource, lib, "CFFileDescriptorCreateRunLoopSource")
	tryRegister(&_CFFileDescriptorDisableCallBacks, lib, "CFFileDescriptorDisableCallBacks")
	tryRegister(&_CFFileDescriptorEnableCallBacks, lib, "CFFileDescriptorEnableCallBacks")
	tryRegister(&_CFFileDescriptorGetContext, lib, "CFFileDescriptorGetContext")
	tryRegister(&_CFFileDescriptorGetNativeDescriptor, lib, "CFFileDescriptorGetNativeDescriptor")
	tryRegister(&_CFFileDescriptorGetTypeID, lib, "CFFileDescriptorGetTypeID")
	tryRegister(&_CFFileDescriptorInvalidate, lib, "CFFileDescriptorInvalidate")
	tryRegister(&_CFFileDescriptorIsValid, lib, "CFFileDescriptorIsValid")
	tryRegister(&_CFFileSecurityClearProperties, lib, "CFFileSecurityClearProperties")
	tryRegister(&_CFFileSecurityCopyAccessControlList, lib, "CFFileSecurityCopyAccessControlList")
	tryRegister(&_CFFileSecurityCopyGroupUUID, lib, "CFFileSecurityCopyGroupUUID")
	tryRegister(&_CFFileSecurityCopyOwnerUUID, lib, "CFFileSecurityCopyOwnerUUID")
	tryRegister(&_CFFileSecurityCreate, lib, "CFFileSecurityCreate")
	tryRegister(&_CFFileSecurityCreateCopy, lib, "CFFileSecurityCreateCopy")
	tryRegister(&_CFFileSecurityGetGroup, lib, "CFFileSecurityGetGroup")
	tryRegister(&_CFFileSecurityGetMode, lib, "CFFileSecurityGetMode")
	tryRegister(&_CFFileSecurityGetOwner, lib, "CFFileSecurityGetOwner")
	tryRegister(&_CFFileSecurityGetTypeID, lib, "CFFileSecurityGetTypeID")
	tryRegister(&_CFFileSecuritySetAccessControlList, lib, "CFFileSecuritySetAccessControlList")
	tryRegister(&_CFFileSecuritySetGroup, lib, "CFFileSecuritySetGroup")
	tryRegister(&_CFFileSecuritySetGroupUUID, lib, "CFFileSecuritySetGroupUUID")
	tryRegister(&_CFFileSecuritySetMode, lib, "CFFileSecuritySetMode")
	tryRegister(&_CFFileSecuritySetOwner, lib, "CFFileSecuritySetOwner")
	tryRegister(&_CFFileSecuritySetOwnerUUID, lib, "CFFileSecuritySetOwnerUUID")
	tryRegister(&_CFGetAllocator, lib, "CFGetAllocator")
	tryRegister(&_CFGetRetainCount, lib, "CFGetRetainCount")
	tryRegister(&_CFGetTypeID, lib, "CFGetTypeID")
	tryRegister(&_CFGregorianDateGetAbsoluteTime, lib, "CFGregorianDateGetAbsoluteTime")
	tryRegister(&_CFGregorianDateIsValid, lib, "CFGregorianDateIsValid")
	tryRegister(&_CFHash, lib, "CFHash")
	tryRegister(&_CFLocaleCopyAvailableLocaleIdentifiers, lib, "CFLocaleCopyAvailableLocaleIdentifiers")
	tryRegister(&_CFLocaleCopyCommonISOCurrencyCodes, lib, "CFLocaleCopyCommonISOCurrencyCodes")
	tryRegister(&_CFLocaleCopyCurrent, lib, "CFLocaleCopyCurrent")
	tryRegister(&_CFLocaleCopyDisplayNameForPropertyValue, lib, "CFLocaleCopyDisplayNameForPropertyValue")
	tryRegister(&_CFLocaleCopyISOCountryCodes, lib, "CFLocaleCopyISOCountryCodes")
	tryRegister(&_CFLocaleCopyISOCurrencyCodes, lib, "CFLocaleCopyISOCurrencyCodes")
	tryRegister(&_CFLocaleCopyISOLanguageCodes, lib, "CFLocaleCopyISOLanguageCodes")
	tryRegister(&_CFLocaleCopyPreferredLanguages, lib, "CFLocaleCopyPreferredLanguages")
	tryRegister(&_CFLocaleCreate, lib, "CFLocaleCreate")
	tryRegister(&_CFLocaleCreateCanonicalLanguageIdentifierFromString, lib, "CFLocaleCreateCanonicalLanguageIdentifierFromString")
	tryRegister(&_CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes, lib, "CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes")
	tryRegister(&_CFLocaleCreateCanonicalLocaleIdentifierFromString, lib, "CFLocaleCreateCanonicalLocaleIdentifierFromString")
	tryRegister(&_CFLocaleCreateComponentsFromLocaleIdentifier, lib, "CFLocaleCreateComponentsFromLocaleIdentifier")
	tryRegister(&_CFLocaleCreateCopy, lib, "CFLocaleCreateCopy")
	tryRegister(&_CFLocaleCreateLocaleIdentifierFromComponents, lib, "CFLocaleCreateLocaleIdentifierFromComponents")
	tryRegister(&_CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode, lib, "CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode")
	tryRegister(&_CFLocaleGetIdentifier, lib, "CFLocaleGetIdentifier")
	tryRegister(&_CFLocaleGetLanguageCharacterDirection, lib, "CFLocaleGetLanguageCharacterDirection")
	tryRegister(&_CFLocaleGetLanguageLineDirection, lib, "CFLocaleGetLanguageLineDirection")
	tryRegister(&_CFLocaleGetSystem, lib, "CFLocaleGetSystem")
	tryRegister(&_CFLocaleGetTypeID, lib, "CFLocaleGetTypeID")
	tryRegister(&_CFLocaleGetValue, lib, "CFLocaleGetValue")
	tryRegister(&_CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier, lib, "CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier")
	tryRegister(&_CFMachPortCreate, lib, "CFMachPortCreate")
	tryRegister(&_CFMachPortCreateRunLoopSource, lib, "CFMachPortCreateRunLoopSource")
	tryRegister(&_CFMachPortCreateWithPort, lib, "CFMachPortCreateWithPort")
	tryRegister(&_CFMachPortGetContext, lib, "CFMachPortGetContext")
	tryRegister(&_CFMachPortGetInvalidationCallBack, lib, "CFMachPortGetInvalidationCallBack")
	tryRegister(&_CFMachPortGetPort, lib, "CFMachPortGetPort")
	tryRegister(&_CFMachPortGetTypeID, lib, "CFMachPortGetTypeID")
	tryRegister(&_CFMachPortInvalidate, lib, "CFMachPortInvalidate")
	tryRegister(&_CFMachPortIsValid, lib, "CFMachPortIsValid")
	tryRegister(&_CFMachPortSetInvalidationCallBack, lib, "CFMachPortSetInvalidationCallBack")
	tryRegister(&_CFMessagePortCreateLocal, lib, "CFMessagePortCreateLocal")
	tryRegister(&_CFMessagePortCreateRemote, lib, "CFMessagePortCreateRemote")
	tryRegister(&_CFMessagePortCreateRunLoopSource, lib, "CFMessagePortCreateRunLoopSource")
	tryRegister(&_CFMessagePortGetContext, lib, "CFMessagePortGetContext")
	tryRegister(&_CFMessagePortGetInvalidationCallBack, lib, "CFMessagePortGetInvalidationCallBack")
	tryRegister(&_CFMessagePortGetName, lib, "CFMessagePortGetName")
	tryRegister(&_CFMessagePortGetTypeID, lib, "CFMessagePortGetTypeID")
	tryRegister(&_CFMessagePortInvalidate, lib, "CFMessagePortInvalidate")
	tryRegister(&_CFMessagePortIsRemote, lib, "CFMessagePortIsRemote")
	tryRegister(&_CFMessagePortIsValid, lib, "CFMessagePortIsValid")
	tryRegister(&_CFMessagePortSendRequest, lib, "CFMessagePortSendRequest")
	tryRegister(&_CFMessagePortSetDispatchQueue, lib, "CFMessagePortSetDispatchQueue")
	tryRegister(&_CFMessagePortSetInvalidationCallBack, lib, "CFMessagePortSetInvalidationCallBack")
	tryRegister(&_CFMessagePortSetName, lib, "CFMessagePortSetName")
	tryRegister(&_CFNotificationCenterAddObserver, lib, "CFNotificationCenterAddObserver")
	tryRegister(&_CFNotificationCenterGetDarwinNotifyCenter, lib, "CFNotificationCenterGetDarwinNotifyCenter")
	tryRegister(&_CFNotificationCenterGetDistributedCenter, lib, "CFNotificationCenterGetDistributedCenter")
	tryRegister(&_CFNotificationCenterGetLocalCenter, lib, "CFNotificationCenterGetLocalCenter")
	tryRegister(&_CFNotificationCenterGetTypeID, lib, "CFNotificationCenterGetTypeID")
	tryRegister(&_CFNotificationCenterPostNotification, lib, "CFNotificationCenterPostNotification")
	tryRegister(&_CFNotificationCenterPostNotificationWithOptions, lib, "CFNotificationCenterPostNotificationWithOptions")
	tryRegister(&_CFNotificationCenterRemoveEveryObserver, lib, "CFNotificationCenterRemoveEveryObserver")
	tryRegister(&_CFNotificationCenterRemoveObserver, lib, "CFNotificationCenterRemoveObserver")
	tryRegister(&_CFNullGetTypeID, lib, "CFNullGetTypeID")
	tryRegister(&_CFNumberCompare, lib, "CFNumberCompare")
	tryRegister(&_CFNumberCreate, lib, "CFNumberCreate")
	tryRegister(&_CFNumberFormatterCopyProperty, lib, "CFNumberFormatterCopyProperty")
	tryRegister(&_CFNumberFormatterCreate, lib, "CFNumberFormatterCreate")
	tryRegister(&_CFNumberFormatterCreateNumberFromString, lib, "CFNumberFormatterCreateNumberFromString")
	tryRegister(&_CFNumberFormatterCreateStringWithNumber, lib, "CFNumberFormatterCreateStringWithNumber")
	tryRegister(&_CFNumberFormatterCreateStringWithValue, lib, "CFNumberFormatterCreateStringWithValue")
	tryRegister(&_CFNumberFormatterGetDecimalInfoForCurrencyCode, lib, "CFNumberFormatterGetDecimalInfoForCurrencyCode")
	tryRegister(&_CFNumberFormatterGetFormat, lib, "CFNumberFormatterGetFormat")
	tryRegister(&_CFNumberFormatterGetLocale, lib, "CFNumberFormatterGetLocale")
	tryRegister(&_CFNumberFormatterGetStyle, lib, "CFNumberFormatterGetStyle")
	tryRegister(&_CFNumberFormatterGetTypeID, lib, "CFNumberFormatterGetTypeID")
	tryRegister(&_CFNumberFormatterGetValueFromString, lib, "CFNumberFormatterGetValueFromString")
	tryRegister(&_CFNumberFormatterSetFormat, lib, "CFNumberFormatterSetFormat")
	tryRegister(&_CFNumberFormatterSetProperty, lib, "CFNumberFormatterSetProperty")
	tryRegister(&_CFNumberGetByteSize, lib, "CFNumberGetByteSize")
	tryRegister(&_CFNumberGetType, lib, "CFNumberGetType")
	tryRegister(&_CFNumberGetTypeID, lib, "CFNumberGetTypeID")
	tryRegister(&_CFNumberGetValue, lib, "CFNumberGetValue")
	tryRegister(&_CFNumberIsFloatType, lib, "CFNumberIsFloatType")
	tryRegister(&_CFPlugInAddInstanceForFactory, lib, "CFPlugInAddInstanceForFactory")
	tryRegister(&_CFPlugInCreate, lib, "CFPlugInCreate")
	tryRegister(&_CFPlugInFindFactoriesForPlugInType, lib, "CFPlugInFindFactoriesForPlugInType")
	tryRegister(&_CFPlugInFindFactoriesForPlugInTypeInPlugIn, lib, "CFPlugInFindFactoriesForPlugInTypeInPlugIn")
	tryRegister(&_CFPlugInGetBundle, lib, "CFPlugInGetBundle")
	tryRegister(&_CFPlugInGetTypeID, lib, "CFPlugInGetTypeID")
	tryRegister(&_CFPlugInInstanceCreate, lib, "CFPlugInInstanceCreate")
	tryRegister(&_CFPlugInInstanceCreateWithInstanceDataSize, lib, "CFPlugInInstanceCreateWithInstanceDataSize")
	tryRegister(&_CFPlugInInstanceGetFactoryName, lib, "CFPlugInInstanceGetFactoryName")
	tryRegister(&_CFPlugInInstanceGetInstanceData, lib, "CFPlugInInstanceGetInstanceData")
	tryRegister(&_CFPlugInInstanceGetInterfaceFunctionTable, lib, "CFPlugInInstanceGetInterfaceFunctionTable")
	tryRegister(&_CFPlugInInstanceGetTypeID, lib, "CFPlugInInstanceGetTypeID")
	tryRegister(&_CFPlugInIsLoadOnDemand, lib, "CFPlugInIsLoadOnDemand")
	tryRegister(&_CFPlugInRegisterFactoryFunction, lib, "CFPlugInRegisterFactoryFunction")
	tryRegister(&_CFPlugInRegisterFactoryFunctionByName, lib, "CFPlugInRegisterFactoryFunctionByName")
	tryRegister(&_CFPlugInRegisterPlugInType, lib, "CFPlugInRegisterPlugInType")
	tryRegister(&_CFPlugInRemoveInstanceForFactory, lib, "CFPlugInRemoveInstanceForFactory")
	tryRegister(&_CFPlugInSetLoadOnDemand, lib, "CFPlugInSetLoadOnDemand")
	tryRegister(&_CFPlugInUnregisterFactory, lib, "CFPlugInUnregisterFactory")
	tryRegister(&_CFPlugInUnregisterPlugInType, lib, "CFPlugInUnregisterPlugInType")
	tryRegister(&_CFPreferencesAddSuitePreferencesToApp, lib, "CFPreferencesAddSuitePreferencesToApp")
	tryRegister(&_CFPreferencesAppSynchronize, lib, "CFPreferencesAppSynchronize")
	tryRegister(&_CFPreferencesAppValueIsForced, lib, "CFPreferencesAppValueIsForced")
	tryRegister(&_CFPreferencesCopyAppValue, lib, "CFPreferencesCopyAppValue")
	tryRegister(&_CFPreferencesCopyApplicationList, lib, "CFPreferencesCopyApplicationList")
	tryRegister(&_CFPreferencesCopyKeyList, lib, "CFPreferencesCopyKeyList")
	tryRegister(&_CFPreferencesCopyMultiple, lib, "CFPreferencesCopyMultiple")
	tryRegister(&_CFPreferencesCopyValue, lib, "CFPreferencesCopyValue")
	tryRegister(&_CFPreferencesGetAppBooleanValue, lib, "CFPreferencesGetAppBooleanValue")
	tryRegister(&_CFPreferencesGetAppIntegerValue, lib, "CFPreferencesGetAppIntegerValue")
	tryRegister(&_CFPreferencesRemoveSuitePreferencesFromApp, lib, "CFPreferencesRemoveSuitePreferencesFromApp")
	tryRegister(&_CFPreferencesSetAppValue, lib, "CFPreferencesSetAppValue")
	tryRegister(&_CFPreferencesSetMultiple, lib, "CFPreferencesSetMultiple")
	tryRegister(&_CFPreferencesSetValue, lib, "CFPreferencesSetValue")
	tryRegister(&_CFPreferencesSynchronize, lib, "CFPreferencesSynchronize")
	tryRegister(&_CFPropertyListCreateData, lib, "CFPropertyListCreateData")
	tryRegister(&_CFPropertyListCreateDeepCopy, lib, "CFPropertyListCreateDeepCopy")
	tryRegister(&_CFPropertyListCreateFromStream, lib, "CFPropertyListCreateFromStream")
	tryRegister(&_CFPropertyListCreateFromXMLData, lib, "CFPropertyListCreateFromXMLData")
	tryRegister(&_CFPropertyListCreateWithData, lib, "CFPropertyListCreateWithData")
	tryRegister(&_CFPropertyListCreateWithStream, lib, "CFPropertyListCreateWithStream")
	tryRegister(&_CFPropertyListCreateXMLData, lib, "CFPropertyListCreateXMLData")
	tryRegister(&_CFPropertyListIsValid, lib, "CFPropertyListIsValid")
	tryRegister(&_CFPropertyListWrite, lib, "CFPropertyListWrite")
	tryRegister(&_CFPropertyListWriteToStream, lib, "CFPropertyListWriteToStream")
	tryRegister(&_CFReadStreamClose, lib, "CFReadStreamClose")
	tryRegister(&_CFReadStreamCopyDispatchQueue, lib, "CFReadStreamCopyDispatchQueue")
	tryRegister(&_CFReadStreamCopyError, lib, "CFReadStreamCopyError")
	tryRegister(&_CFReadStreamCopyProperty, lib, "CFReadStreamCopyProperty")
	tryRegister(&_CFReadStreamCreateWithBytesNoCopy, lib, "CFReadStreamCreateWithBytesNoCopy")
	tryRegister(&_CFReadStreamCreateWithFile, lib, "CFReadStreamCreateWithFile")
	tryRegister(&_CFReadStreamGetBuffer, lib, "CFReadStreamGetBuffer")
	tryRegister(&_CFReadStreamGetError, lib, "CFReadStreamGetError")
	tryRegister(&_CFReadStreamGetStatus, lib, "CFReadStreamGetStatus")
	tryRegister(&_CFReadStreamGetTypeID, lib, "CFReadStreamGetTypeID")
	tryRegister(&_CFReadStreamHasBytesAvailable, lib, "CFReadStreamHasBytesAvailable")
	tryRegister(&_CFReadStreamOpen, lib, "CFReadStreamOpen")
	tryRegister(&_CFReadStreamRead, lib, "CFReadStreamRead")
	tryRegister(&_CFReadStreamScheduleWithRunLoop, lib, "CFReadStreamScheduleWithRunLoop")
	tryRegister(&_CFReadStreamSetClient, lib, "CFReadStreamSetClient")
	tryRegister(&_CFReadStreamSetDispatchQueue, lib, "CFReadStreamSetDispatchQueue")
	tryRegister(&_CFReadStreamSetProperty, lib, "CFReadStreamSetProperty")
	tryRegister(&_CFReadStreamUnscheduleFromRunLoop, lib, "CFReadStreamUnscheduleFromRunLoop")
	tryRegister(&_CFRunLoopAddCommonMode, lib, "CFRunLoopAddCommonMode")
	tryRegister(&_CFRunLoopAddObserver, lib, "CFRunLoopAddObserver")
	tryRegister(&_CFRunLoopAddSource, lib, "CFRunLoopAddSource")
	tryRegister(&_CFRunLoopAddTimer, lib, "CFRunLoopAddTimer")
	tryRegister(&_CFRunLoopContainsObserver, lib, "CFRunLoopContainsObserver")
	tryRegister(&_CFRunLoopContainsSource, lib, "CFRunLoopContainsSource")
	tryRegister(&_CFRunLoopContainsTimer, lib, "CFRunLoopContainsTimer")
	tryRegister(&_CFRunLoopCopyAllModes, lib, "CFRunLoopCopyAllModes")
	tryRegister(&_CFRunLoopCopyCurrentMode, lib, "CFRunLoopCopyCurrentMode")
	tryRegister(&_CFRunLoopGetCurrent, lib, "CFRunLoopGetCurrent")
	tryRegister(&_CFRunLoopGetMain, lib, "CFRunLoopGetMain")
	tryRegister(&_CFRunLoopGetNextTimerFireDate, lib, "CFRunLoopGetNextTimerFireDate")
	tryRegister(&_CFRunLoopGetTypeID, lib, "CFRunLoopGetTypeID")
	tryRegister(&_CFRunLoopIsWaiting, lib, "CFRunLoopIsWaiting")
	tryRegister(&_CFRunLoopObserverCreate, lib, "CFRunLoopObserverCreate")
	tryRegister(&_CFRunLoopObserverCreateWithHandler, lib, "CFRunLoopObserverCreateWithHandler")
	tryRegister(&_CFRunLoopObserverDoesRepeat, lib, "CFRunLoopObserverDoesRepeat")
	tryRegister(&_CFRunLoopObserverGetActivities, lib, "CFRunLoopObserverGetActivities")
	tryRegister(&_CFRunLoopObserverGetContext, lib, "CFRunLoopObserverGetContext")
	tryRegister(&_CFRunLoopObserverGetOrder, lib, "CFRunLoopObserverGetOrder")
	tryRegister(&_CFRunLoopObserverGetTypeID, lib, "CFRunLoopObserverGetTypeID")
	tryRegister(&_CFRunLoopObserverInvalidate, lib, "CFRunLoopObserverInvalidate")
	tryRegister(&_CFRunLoopObserverIsValid, lib, "CFRunLoopObserverIsValid")
	tryRegister(&_CFRunLoopPerformBlock, lib, "CFRunLoopPerformBlock")
	tryRegister(&_CFRunLoopRemoveObserver, lib, "CFRunLoopRemoveObserver")
	tryRegister(&_CFRunLoopRemoveSource, lib, "CFRunLoopRemoveSource")
	tryRegister(&_CFRunLoopRemoveTimer, lib, "CFRunLoopRemoveTimer")
	tryRegister(&_CFRunLoopRun, lib, "CFRunLoopRun")
	tryRegister(&_CFRunLoopRunInMode, lib, "CFRunLoopRunInMode")
	tryRegister(&_CFRunLoopSourceCreate, lib, "CFRunLoopSourceCreate")
	tryRegister(&_CFRunLoopSourceGetContext, lib, "CFRunLoopSourceGetContext")
	tryRegister(&_CFRunLoopSourceGetOrder, lib, "CFRunLoopSourceGetOrder")
	tryRegister(&_CFRunLoopSourceGetTypeID, lib, "CFRunLoopSourceGetTypeID")
	tryRegister(&_CFRunLoopSourceInvalidate, lib, "CFRunLoopSourceInvalidate")
	tryRegister(&_CFRunLoopSourceIsValid, lib, "CFRunLoopSourceIsValid")
	tryRegister(&_CFRunLoopSourceSignal, lib, "CFRunLoopSourceSignal")
	tryRegister(&_CFRunLoopStop, lib, "CFRunLoopStop")
	tryRegister(&_CFRunLoopTimerCreate, lib, "CFRunLoopTimerCreate")
	tryRegister(&_CFRunLoopTimerCreateWithHandler, lib, "CFRunLoopTimerCreateWithHandler")
	tryRegister(&_CFRunLoopTimerDoesRepeat, lib, "CFRunLoopTimerDoesRepeat")
	tryRegister(&_CFRunLoopTimerGetContext, lib, "CFRunLoopTimerGetContext")
	tryRegister(&_CFRunLoopTimerGetInterval, lib, "CFRunLoopTimerGetInterval")
	tryRegister(&_CFRunLoopTimerGetNextFireDate, lib, "CFRunLoopTimerGetNextFireDate")
	tryRegister(&_CFRunLoopTimerGetOrder, lib, "CFRunLoopTimerGetOrder")
	tryRegister(&_CFRunLoopTimerGetTolerance, lib, "CFRunLoopTimerGetTolerance")
	tryRegister(&_CFRunLoopTimerGetTypeID, lib, "CFRunLoopTimerGetTypeID")
	tryRegister(&_CFRunLoopTimerInvalidate, lib, "CFRunLoopTimerInvalidate")
	tryRegister(&_CFRunLoopTimerIsValid, lib, "CFRunLoopTimerIsValid")
	tryRegister(&_CFRunLoopTimerSetNextFireDate, lib, "CFRunLoopTimerSetNextFireDate")
	tryRegister(&_CFRunLoopTimerSetTolerance, lib, "CFRunLoopTimerSetTolerance")
	tryRegister(&_CFRunLoopWakeUp, lib, "CFRunLoopWakeUp")
	tryRegister(&_CFSetAddValue, lib, "CFSetAddValue")
	tryRegister(&_CFSetApplyFunction, lib, "CFSetApplyFunction")
	tryRegister(&_CFSetContainsValue, lib, "CFSetContainsValue")
	tryRegister(&_CFSetCreate, lib, "CFSetCreate")
	tryRegister(&_CFSetCreateCopy, lib, "CFSetCreateCopy")
	tryRegister(&_CFSetCreateMutable, lib, "CFSetCreateMutable")
	tryRegister(&_CFSetCreateMutableCopy, lib, "CFSetCreateMutableCopy")
	tryRegister(&_CFSetGetCount, lib, "CFSetGetCount")
	tryRegister(&_CFSetGetCountOfValue, lib, "CFSetGetCountOfValue")
	tryRegister(&_CFSetGetTypeID, lib, "CFSetGetTypeID")
	tryRegister(&_CFSetGetValue, lib, "CFSetGetValue")
	tryRegister(&_CFSetGetValueIfPresent, lib, "CFSetGetValueIfPresent")
	tryRegister(&_CFSetGetValues, lib, "CFSetGetValues")
	tryRegister(&_CFSetRemoveAllValues, lib, "CFSetRemoveAllValues")
	tryRegister(&_CFSetRemoveValue, lib, "CFSetRemoveValue")
	tryRegister(&_CFSetReplaceValue, lib, "CFSetReplaceValue")
	tryRegister(&_CFSetSetValue, lib, "CFSetSetValue")
	tryRegister(&_CFShow, lib, "CFShow")
	tryRegister(&_CFShowStr, lib, "CFShowStr")
	tryRegister(&_CFSocketConnectToAddress, lib, "CFSocketConnectToAddress")
	tryRegister(&_CFSocketCopyAddress, lib, "CFSocketCopyAddress")
	tryRegister(&_CFSocketCopyPeerAddress, lib, "CFSocketCopyPeerAddress")
	tryRegister(&_CFSocketCopyRegisteredSocketSignature, lib, "CFSocketCopyRegisteredSocketSignature")
	tryRegister(&_CFSocketCopyRegisteredValue, lib, "CFSocketCopyRegisteredValue")
	tryRegister(&_CFSocketCreate, lib, "CFSocketCreate")
	tryRegister(&_CFSocketCreateConnectedToSocketSignature, lib, "CFSocketCreateConnectedToSocketSignature")
	tryRegister(&_CFSocketCreateRunLoopSource, lib, "CFSocketCreateRunLoopSource")
	tryRegister(&_CFSocketCreateWithNative, lib, "CFSocketCreateWithNative")
	tryRegister(&_CFSocketCreateWithSocketSignature, lib, "CFSocketCreateWithSocketSignature")
	tryRegister(&_CFSocketDisableCallBacks, lib, "CFSocketDisableCallBacks")
	tryRegister(&_CFSocketEnableCallBacks, lib, "CFSocketEnableCallBacks")
	tryRegister(&_CFSocketGetContext, lib, "CFSocketGetContext")
	tryRegister(&_CFSocketGetDefaultNameRegistryPortNumber, lib, "CFSocketGetDefaultNameRegistryPortNumber")
	tryRegister(&_CFSocketGetNative, lib, "CFSocketGetNative")
	tryRegister(&_CFSocketGetSocketFlags, lib, "CFSocketGetSocketFlags")
	tryRegister(&_CFSocketGetTypeID, lib, "CFSocketGetTypeID")
	tryRegister(&_CFSocketInvalidate, lib, "CFSocketInvalidate")
	tryRegister(&_CFSocketIsValid, lib, "CFSocketIsValid")
	tryRegister(&_CFSocketRegisterSocketSignature, lib, "CFSocketRegisterSocketSignature")
	tryRegister(&_CFSocketRegisterValue, lib, "CFSocketRegisterValue")
	tryRegister(&_CFSocketSendData, lib, "CFSocketSendData")
	tryRegister(&_CFSocketSetAddress, lib, "CFSocketSetAddress")
	tryRegister(&_CFSocketSetDefaultNameRegistryPortNumber, lib, "CFSocketSetDefaultNameRegistryPortNumber")
	tryRegister(&_CFSocketSetSocketFlags, lib, "CFSocketSetSocketFlags")
	tryRegister(&_CFSocketUnregister, lib, "CFSocketUnregister")
	tryRegister(&_CFStreamCreateBoundPair, lib, "CFStreamCreateBoundPair")
	tryRegister(&_CFStreamCreatePairWithPeerSocketSignature, lib, "CFStreamCreatePairWithPeerSocketSignature")
	tryRegister(&_CFStreamCreatePairWithSocket, lib, "CFStreamCreatePairWithSocket")
	tryRegister(&_CFStreamCreatePairWithSocketToHost, lib, "CFStreamCreatePairWithSocketToHost")
	tryRegister(&_CFStringAppend, lib, "CFStringAppend")
	tryRegister(&_CFStringAppendCString, lib, "CFStringAppendCString")
	tryRegister(&_CFStringAppendCharacters, lib, "CFStringAppendCharacters")
	tryRegister(&_CFStringAppendFormat, lib, "CFStringAppendFormat")
	tryRegister(&_CFStringAppendFormatAndArguments, lib, "CFStringAppendFormatAndArguments")
	tryRegister(&_CFStringAppendPascalString, lib, "CFStringAppendPascalString")
	tryRegister(&_CFStringCapitalize, lib, "CFStringCapitalize")
	tryRegister(&_CFStringCompare, lib, "CFStringCompare")
	tryRegister(&_CFStringCompareWithOptions, lib, "CFStringCompareWithOptions")
	tryRegister(&_CFStringCompareWithOptionsAndLocale, lib, "CFStringCompareWithOptionsAndLocale")
	tryRegister(&_CFStringConvertEncodingToIANACharSetName, lib, "CFStringConvertEncodingToIANACharSetName")
	tryRegister(&_CFStringConvertEncodingToNSStringEncoding, lib, "CFStringConvertEncodingToNSStringEncoding")
	tryRegister(&_CFStringConvertEncodingToWindowsCodepage, lib, "CFStringConvertEncodingToWindowsCodepage")
	tryRegister(&_CFStringConvertIANACharSetNameToEncoding, lib, "CFStringConvertIANACharSetNameToEncoding")
	tryRegister(&_CFStringConvertNSStringEncodingToEncoding, lib, "CFStringConvertNSStringEncodingToEncoding")
	tryRegister(&_CFStringConvertWindowsCodepageToEncoding, lib, "CFStringConvertWindowsCodepageToEncoding")
	tryRegister(&_CFStringCreateArrayBySeparatingStrings, lib, "CFStringCreateArrayBySeparatingStrings")
	tryRegister(&_CFStringCreateArrayWithFindResults, lib, "CFStringCreateArrayWithFindResults")
	tryRegister(&_CFStringCreateByCombiningStrings, lib, "CFStringCreateByCombiningStrings")
	tryRegister(&_CFStringCreateCopy, lib, "CFStringCreateCopy")
	tryRegister(&_CFStringCreateExternalRepresentation, lib, "CFStringCreateExternalRepresentation")
	tryRegister(&_CFStringCreateFromExternalRepresentation, lib, "CFStringCreateFromExternalRepresentation")
	tryRegister(&_CFStringCreateMutable, lib, "CFStringCreateMutable")
	tryRegister(&_CFStringCreateMutableCopy, lib, "CFStringCreateMutableCopy")
	tryRegister(&_CFStringCreateMutableWithExternalCharactersNoCopy, lib, "CFStringCreateMutableWithExternalCharactersNoCopy")
	tryRegister(&_CFStringCreateStringWithValidatedFormat, lib, "CFStringCreateStringWithValidatedFormat")
	tryRegister(&_CFStringCreateStringWithValidatedFormatAndArguments, lib, "CFStringCreateStringWithValidatedFormatAndArguments")
	tryRegister(&_CFStringCreateWithBytes, lib, "CFStringCreateWithBytes")
	tryRegister(&_CFStringCreateWithBytesNoCopy, lib, "CFStringCreateWithBytesNoCopy")
	tryRegister(&_CFStringCreateWithCString, lib, "CFStringCreateWithCString")
	tryRegister(&_CFStringCreateWithCStringNoCopy, lib, "CFStringCreateWithCStringNoCopy")
	tryRegister(&_CFStringCreateWithCharacters, lib, "CFStringCreateWithCharacters")
	tryRegister(&_CFStringCreateWithCharactersNoCopy, lib, "CFStringCreateWithCharactersNoCopy")
	tryRegister(&_CFStringCreateWithFileSystemRepresentation, lib, "CFStringCreateWithFileSystemRepresentation")
	tryRegister(&_CFStringCreateWithFormat, lib, "CFStringCreateWithFormat")
	tryRegister(&_CFStringCreateWithFormatAndArguments, lib, "CFStringCreateWithFormatAndArguments")
	tryRegister(&_CFStringCreateWithPascalString, lib, "CFStringCreateWithPascalString")
	tryRegister(&_CFStringCreateWithPascalStringNoCopy, lib, "CFStringCreateWithPascalStringNoCopy")
	tryRegister(&_CFStringCreateWithSubstring, lib, "CFStringCreateWithSubstring")
	tryRegister(&_CFStringDelete, lib, "CFStringDelete")
	tryRegister(&_CFStringFind, lib, "CFStringFind")
	tryRegister(&_CFStringFindAndReplace, lib, "CFStringFindAndReplace")
	tryRegister(&_CFStringFindCharacterFromSet, lib, "CFStringFindCharacterFromSet")
	tryRegister(&_CFStringFindWithOptions, lib, "CFStringFindWithOptions")
	tryRegister(&_CFStringFindWithOptionsAndLocale, lib, "CFStringFindWithOptionsAndLocale")
	tryRegister(&_CFStringFold, lib, "CFStringFold")
	tryRegister(&_CFStringGetBytes, lib, "CFStringGetBytes")
	tryRegister(&_CFStringGetCString, lib, "CFStringGetCString")
	tryRegister(&_CFStringGetCStringPtr, lib, "CFStringGetCStringPtr")
	tryRegister(&_CFStringGetCharacterAtIndex, lib, "CFStringGetCharacterAtIndex")
	tryRegister(&_CFStringGetCharacters, lib, "CFStringGetCharacters")
	tryRegister(&_CFStringGetCharactersPtr, lib, "CFStringGetCharactersPtr")
	tryRegister(&_CFStringGetDoubleValue, lib, "CFStringGetDoubleValue")
	tryRegister(&_CFStringGetFastestEncoding, lib, "CFStringGetFastestEncoding")
	tryRegister(&_CFStringGetFileSystemRepresentation, lib, "CFStringGetFileSystemRepresentation")
	tryRegister(&_CFStringGetHyphenationLocationBeforeIndex, lib, "CFStringGetHyphenationLocationBeforeIndex")
	tryRegister(&_CFStringGetIntValue, lib, "CFStringGetIntValue")
	tryRegister(&_CFStringGetLength, lib, "CFStringGetLength")
	tryRegister(&_CFStringGetLineBounds, lib, "CFStringGetLineBounds")
	tryRegister(&_CFStringGetListOfAvailableEncodings, lib, "CFStringGetListOfAvailableEncodings")
	tryRegister(&_CFStringGetMaximumSizeForEncoding, lib, "CFStringGetMaximumSizeForEncoding")
	tryRegister(&_CFStringGetMaximumSizeOfFileSystemRepresentation, lib, "CFStringGetMaximumSizeOfFileSystemRepresentation")
	tryRegister(&_CFStringGetMostCompatibleMacStringEncoding, lib, "CFStringGetMostCompatibleMacStringEncoding")
	tryRegister(&_CFStringGetNameOfEncoding, lib, "CFStringGetNameOfEncoding")
	tryRegister(&_CFStringGetParagraphBounds, lib, "CFStringGetParagraphBounds")
	tryRegister(&_CFStringGetPascalString, lib, "CFStringGetPascalString")
	tryRegister(&_CFStringGetPascalStringPtr, lib, "CFStringGetPascalStringPtr")
	tryRegister(&_CFStringGetRangeOfComposedCharactersAtIndex, lib, "CFStringGetRangeOfComposedCharactersAtIndex")
	tryRegister(&_CFStringGetSmallestEncoding, lib, "CFStringGetSmallestEncoding")
	tryRegister(&_CFStringGetSystemEncoding, lib, "CFStringGetSystemEncoding")
	tryRegister(&_CFStringGetTypeID, lib, "CFStringGetTypeID")
	tryRegister(&_CFStringHasPrefix, lib, "CFStringHasPrefix")
	tryRegister(&_CFStringHasSuffix, lib, "CFStringHasSuffix")
	tryRegister(&_CFStringInsert, lib, "CFStringInsert")
	tryRegister(&_CFStringIsEncodingAvailable, lib, "CFStringIsEncodingAvailable")
	tryRegister(&_CFStringIsHyphenationAvailableForLocale, lib, "CFStringIsHyphenationAvailableForLocale")
	tryRegister(&_CFStringLowercase, lib, "CFStringLowercase")
	tryRegister(&_CFStringNormalize, lib, "CFStringNormalize")
	tryRegister(&_CFStringPad, lib, "CFStringPad")
	tryRegister(&_CFStringReplace, lib, "CFStringReplace")
	tryRegister(&_CFStringReplaceAll, lib, "CFStringReplaceAll")
	tryRegister(&_CFStringSetExternalCharactersNoCopy, lib, "CFStringSetExternalCharactersNoCopy")
	tryRegister(&_CFStringTokenizerAdvanceToNextToken, lib, "CFStringTokenizerAdvanceToNextToken")
	tryRegister(&_CFStringTokenizerCopyBestStringLanguage, lib, "CFStringTokenizerCopyBestStringLanguage")
	tryRegister(&_CFStringTokenizerCopyCurrentTokenAttribute, lib, "CFStringTokenizerCopyCurrentTokenAttribute")
	tryRegister(&_CFStringTokenizerCreate, lib, "CFStringTokenizerCreate")
	tryRegister(&_CFStringTokenizerGetCurrentSubTokens, lib, "CFStringTokenizerGetCurrentSubTokens")
	tryRegister(&_CFStringTokenizerGetCurrentTokenRange, lib, "CFStringTokenizerGetCurrentTokenRange")
	tryRegister(&_CFStringTokenizerGetTypeID, lib, "CFStringTokenizerGetTypeID")
	tryRegister(&_CFStringTokenizerGoToTokenAtIndex, lib, "CFStringTokenizerGoToTokenAtIndex")
	tryRegister(&_CFStringTokenizerSetString, lib, "CFStringTokenizerSetString")
	tryRegister(&_CFStringTransform, lib, "CFStringTransform")
	tryRegister(&_CFStringTrim, lib, "CFStringTrim")
	tryRegister(&_CFStringTrimWhitespace, lib, "CFStringTrimWhitespace")
	tryRegister(&_CFStringUppercase, lib, "CFStringUppercase")
	tryRegister(&_CFTimeZoneCopyAbbreviation, lib, "CFTimeZoneCopyAbbreviation")
	tryRegister(&_CFTimeZoneCopyAbbreviationDictionary, lib, "CFTimeZoneCopyAbbreviationDictionary")
	tryRegister(&_CFTimeZoneCopyDefault, lib, "CFTimeZoneCopyDefault")
	tryRegister(&_CFTimeZoneCopyKnownNames, lib, "CFTimeZoneCopyKnownNames")
	tryRegister(&_CFTimeZoneCopyLocalizedName, lib, "CFTimeZoneCopyLocalizedName")
	tryRegister(&_CFTimeZoneCopySystem, lib, "CFTimeZoneCopySystem")
	tryRegister(&_CFTimeZoneCreate, lib, "CFTimeZoneCreate")
	tryRegister(&_CFTimeZoneCreateWithName, lib, "CFTimeZoneCreateWithName")
	tryRegister(&_CFTimeZoneCreateWithTimeIntervalFromGMT, lib, "CFTimeZoneCreateWithTimeIntervalFromGMT")
	tryRegister(&_CFTimeZoneGetData, lib, "CFTimeZoneGetData")
	tryRegister(&_CFTimeZoneGetDaylightSavingTimeOffset, lib, "CFTimeZoneGetDaylightSavingTimeOffset")
	tryRegister(&_CFTimeZoneGetName, lib, "CFTimeZoneGetName")
	tryRegister(&_CFTimeZoneGetNextDaylightSavingTimeTransition, lib, "CFTimeZoneGetNextDaylightSavingTimeTransition")
	tryRegister(&_CFTimeZoneGetSecondsFromGMT, lib, "CFTimeZoneGetSecondsFromGMT")
	tryRegister(&_CFTimeZoneGetTypeID, lib, "CFTimeZoneGetTypeID")
	tryRegister(&_CFTimeZoneIsDaylightSavingTime, lib, "CFTimeZoneIsDaylightSavingTime")
	tryRegister(&_CFTimeZoneResetSystem, lib, "CFTimeZoneResetSystem")
	tryRegister(&_CFTimeZoneSetAbbreviationDictionary, lib, "CFTimeZoneSetAbbreviationDictionary")
	tryRegister(&_CFTimeZoneSetDefault, lib, "CFTimeZoneSetDefault")
	tryRegister(&_CFTreeAppendChild, lib, "CFTreeAppendChild")
	tryRegister(&_CFTreeApplyFunctionToChildren, lib, "CFTreeApplyFunctionToChildren")
	tryRegister(&_CFTreeCreate, lib, "CFTreeCreate")
	tryRegister(&_CFTreeFindRoot, lib, "CFTreeFindRoot")
	tryRegister(&_CFTreeGetChildAtIndex, lib, "CFTreeGetChildAtIndex")
	tryRegister(&_CFTreeGetChildCount, lib, "CFTreeGetChildCount")
	tryRegister(&_CFTreeGetChildren, lib, "CFTreeGetChildren")
	tryRegister(&_CFTreeGetContext, lib, "CFTreeGetContext")
	tryRegister(&_CFTreeGetFirstChild, lib, "CFTreeGetFirstChild")
	tryRegister(&_CFTreeGetNextSibling, lib, "CFTreeGetNextSibling")
	tryRegister(&_CFTreeGetParent, lib, "CFTreeGetParent")
	tryRegister(&_CFTreeGetTypeID, lib, "CFTreeGetTypeID")
	tryRegister(&_CFTreeInsertSibling, lib, "CFTreeInsertSibling")
	tryRegister(&_CFTreePrependChild, lib, "CFTreePrependChild")
	tryRegister(&_CFTreeRemove, lib, "CFTreeRemove")
	tryRegister(&_CFTreeRemoveAllChildren, lib, "CFTreeRemoveAllChildren")
	tryRegister(&_CFTreeSetContext, lib, "CFTreeSetContext")
	tryRegister(&_CFTreeSortChildren, lib, "CFTreeSortChildren")
	tryRegister(&_CFURLCanBeDecomposed, lib, "CFURLCanBeDecomposed")
	tryRegister(&_CFURLClearResourcePropertyCache, lib, "CFURLClearResourcePropertyCache")
	tryRegister(&_CFURLClearResourcePropertyCacheForKey, lib, "CFURLClearResourcePropertyCacheForKey")
	tryRegister(&_CFURLCopyAbsoluteURL, lib, "CFURLCopyAbsoluteURL")
	tryRegister(&_CFURLCopyFileSystemPath, lib, "CFURLCopyFileSystemPath")
	tryRegister(&_CFURLCopyFragment, lib, "CFURLCopyFragment")
	tryRegister(&_CFURLCopyHostName, lib, "CFURLCopyHostName")
	tryRegister(&_CFURLCopyLastPathComponent, lib, "CFURLCopyLastPathComponent")
	tryRegister(&_CFURLCopyNetLocation, lib, "CFURLCopyNetLocation")
	tryRegister(&_CFURLCopyParameterString, lib, "CFURLCopyParameterString")
	tryRegister(&_CFURLCopyPassword, lib, "CFURLCopyPassword")
	tryRegister(&_CFURLCopyPath, lib, "CFURLCopyPath")
	tryRegister(&_CFURLCopyPathExtension, lib, "CFURLCopyPathExtension")
	tryRegister(&_CFURLCopyQueryString, lib, "CFURLCopyQueryString")
	tryRegister(&_CFURLCopyResourcePropertiesForKeys, lib, "CFURLCopyResourcePropertiesForKeys")
	tryRegister(&_CFURLCopyResourcePropertyForKey, lib, "CFURLCopyResourcePropertyForKey")
	tryRegister(&_CFURLCopyResourceSpecifier, lib, "CFURLCopyResourceSpecifier")
	tryRegister(&_CFURLCopyScheme, lib, "CFURLCopyScheme")
	tryRegister(&_CFURLCopyStrictPath, lib, "CFURLCopyStrictPath")
	tryRegister(&_CFURLCopyUserName, lib, "CFURLCopyUserName")
	tryRegister(&_CFURLCreateAbsoluteURLWithBytes, lib, "CFURLCreateAbsoluteURLWithBytes")
	tryRegister(&_CFURLCreateBookmarkData, lib, "CFURLCreateBookmarkData")
	tryRegister(&_CFURLCreateBookmarkDataFromAliasRecord, lib, "CFURLCreateBookmarkDataFromAliasRecord")
	tryRegister(&_CFURLCreateBookmarkDataFromFile, lib, "CFURLCreateBookmarkDataFromFile")
	tryRegister(&_CFURLCreateByResolvingBookmarkData, lib, "CFURLCreateByResolvingBookmarkData")
	tryRegister(&_CFURLCreateCopyAppendingPathComponent, lib, "CFURLCreateCopyAppendingPathComponent")
	tryRegister(&_CFURLCreateCopyAppendingPathExtension, lib, "CFURLCreateCopyAppendingPathExtension")
	tryRegister(&_CFURLCreateCopyDeletingLastPathComponent, lib, "CFURLCreateCopyDeletingLastPathComponent")
	tryRegister(&_CFURLCreateCopyDeletingPathExtension, lib, "CFURLCreateCopyDeletingPathExtension")
	tryRegister(&_CFURLCreateData, lib, "CFURLCreateData")
	tryRegister(&_CFURLCreateDataAndPropertiesFromResource, lib, "CFURLCreateDataAndPropertiesFromResource")
	tryRegister(&_CFURLCreateFilePathURL, lib, "CFURLCreateFilePathURL")
	tryRegister(&_CFURLCreateFileReferenceURL, lib, "CFURLCreateFileReferenceURL")
	tryRegister(&_CFURLCreateFromFSRef, lib, "CFURLCreateFromFSRef")
	tryRegister(&_CFURLCreateFromFileSystemRepresentation, lib, "CFURLCreateFromFileSystemRepresentation")
	tryRegister(&_CFURLCreateFromFileSystemRepresentationRelativeToBase, lib, "CFURLCreateFromFileSystemRepresentationRelativeToBase")
	tryRegister(&_CFURLCreatePropertyFromResource, lib, "CFURLCreatePropertyFromResource")
	tryRegister(&_CFURLCreateResourcePropertiesForKeysFromBookmarkData, lib, "CFURLCreateResourcePropertiesForKeysFromBookmarkData")
	tryRegister(&_CFURLCreateResourcePropertyForKeyFromBookmarkData, lib, "CFURLCreateResourcePropertyForKeyFromBookmarkData")
	tryRegister(&_CFURLCreateStringByAddingPercentEscapes, lib, "CFURLCreateStringByAddingPercentEscapes")
	tryRegister(&_CFURLCreateStringByReplacingPercentEscapes, lib, "CFURLCreateStringByReplacingPercentEscapes")
	tryRegister(&_CFURLCreateStringByReplacingPercentEscapesUsingEncoding, lib, "CFURLCreateStringByReplacingPercentEscapesUsingEncoding")
	tryRegister(&_CFURLCreateWithBytes, lib, "CFURLCreateWithBytes")
	tryRegister(&_CFURLCreateWithFileSystemPath, lib, "CFURLCreateWithFileSystemPath")
	tryRegister(&_CFURLCreateWithFileSystemPathRelativeToBase, lib, "CFURLCreateWithFileSystemPathRelativeToBase")
	tryRegister(&_CFURLCreateWithString, lib, "CFURLCreateWithString")
	tryRegister(&_CFURLDestroyResource, lib, "CFURLDestroyResource")
	tryRegister(&_CFURLEnumeratorCreateForDirectoryURL, lib, "CFURLEnumeratorCreateForDirectoryURL")
	tryRegister(&_CFURLEnumeratorCreateForMountedVolumes, lib, "CFURLEnumeratorCreateForMountedVolumes")
	tryRegister(&_CFURLEnumeratorGetDescendentLevel, lib, "CFURLEnumeratorGetDescendentLevel")
	tryRegister(&_CFURLEnumeratorGetNextURL, lib, "CFURLEnumeratorGetNextURL")
	tryRegister(&_CFURLEnumeratorGetSourceDidChange, lib, "CFURLEnumeratorGetSourceDidChange")
	tryRegister(&_CFURLEnumeratorGetTypeID, lib, "CFURLEnumeratorGetTypeID")
	tryRegister(&_CFURLEnumeratorSkipDescendents, lib, "CFURLEnumeratorSkipDescendents")
	tryRegister(&_CFURLGetBaseURL, lib, "CFURLGetBaseURL")
	tryRegister(&_CFURLGetByteRangeForComponent, lib, "CFURLGetByteRangeForComponent")
	tryRegister(&_CFURLGetBytes, lib, "CFURLGetBytes")
	tryRegister(&_CFURLGetFSRef, lib, "CFURLGetFSRef")
	tryRegister(&_CFURLGetFileSystemRepresentation, lib, "CFURLGetFileSystemRepresentation")
	tryRegister(&_CFURLGetPortNumber, lib, "CFURLGetPortNumber")
	tryRegister(&_CFURLGetString, lib, "CFURLGetString")
	tryRegister(&_CFURLGetTypeID, lib, "CFURLGetTypeID")
	tryRegister(&_CFURLHasDirectoryPath, lib, "CFURLHasDirectoryPath")
	tryRegister(&_CFURLIsFileReferenceURL, lib, "CFURLIsFileReferenceURL")
	tryRegister(&_CFURLResourceIsReachable, lib, "CFURLResourceIsReachable")
	tryRegister(&_CFURLSetResourcePropertiesForKeys, lib, "CFURLSetResourcePropertiesForKeys")
	tryRegister(&_CFURLSetResourcePropertyForKey, lib, "CFURLSetResourcePropertyForKey")
	tryRegister(&_CFURLSetTemporaryResourcePropertyForKey, lib, "CFURLSetTemporaryResourcePropertyForKey")
	tryRegister(&_CFURLStartAccessingSecurityScopedResource, lib, "CFURLStartAccessingSecurityScopedResource")
	tryRegister(&_CFURLStopAccessingSecurityScopedResource, lib, "CFURLStopAccessingSecurityScopedResource")
	tryRegister(&_CFURLWriteBookmarkDataToFile, lib, "CFURLWriteBookmarkDataToFile")
	tryRegister(&_CFURLWriteDataAndPropertiesToResource, lib, "CFURLWriteDataAndPropertiesToResource")
	tryRegister(&_CFUUIDCreate, lib, "CFUUIDCreate")
	tryRegister(&_CFUUIDCreateFromString, lib, "CFUUIDCreateFromString")
	tryRegister(&_CFUUIDCreateFromUUIDBytes, lib, "CFUUIDCreateFromUUIDBytes")
	tryRegister(&_CFUUIDCreateString, lib, "CFUUIDCreateString")
	tryRegister(&_CFUUIDCreateWithBytes, lib, "CFUUIDCreateWithBytes")
	tryRegister(&_CFUUIDGetConstantUUIDWithBytes, lib, "CFUUIDGetConstantUUIDWithBytes")
	tryRegister(&_CFUUIDGetTypeID, lib, "CFUUIDGetTypeID")
	tryRegister(&_CFUUIDGetUUIDBytes, lib, "CFUUIDGetUUIDBytes")
	tryRegister(&_CFUserNotificationCancel, lib, "CFUserNotificationCancel")
	tryRegister(&_CFUserNotificationCreate, lib, "CFUserNotificationCreate")
	tryRegister(&_CFUserNotificationCreateRunLoopSource, lib, "CFUserNotificationCreateRunLoopSource")
	tryRegister(&_CFUserNotificationDisplayAlert, lib, "CFUserNotificationDisplayAlert")
	tryRegister(&_CFUserNotificationDisplayNotice, lib, "CFUserNotificationDisplayNotice")
	tryRegister(&_CFUserNotificationGetResponseDictionary, lib, "CFUserNotificationGetResponseDictionary")
	tryRegister(&_CFUserNotificationGetResponseValue, lib, "CFUserNotificationGetResponseValue")
	tryRegister(&_CFUserNotificationGetTypeID, lib, "CFUserNotificationGetTypeID")
	tryRegister(&_CFUserNotificationReceiveResponse, lib, "CFUserNotificationReceiveResponse")
	tryRegister(&_CFUserNotificationUpdate, lib, "CFUserNotificationUpdate")
	tryRegister(&_CFWriteStreamCanAcceptBytes, lib, "CFWriteStreamCanAcceptBytes")
	tryRegister(&_CFWriteStreamClose, lib, "CFWriteStreamClose")
	tryRegister(&_CFWriteStreamCopyDispatchQueue, lib, "CFWriteStreamCopyDispatchQueue")
	tryRegister(&_CFWriteStreamCopyError, lib, "CFWriteStreamCopyError")
	tryRegister(&_CFWriteStreamCopyProperty, lib, "CFWriteStreamCopyProperty")
	tryRegister(&_CFWriteStreamCreateWithAllocatedBuffers, lib, "CFWriteStreamCreateWithAllocatedBuffers")
	tryRegister(&_CFWriteStreamCreateWithBuffer, lib, "CFWriteStreamCreateWithBuffer")
	tryRegister(&_CFWriteStreamCreateWithFile, lib, "CFWriteStreamCreateWithFile")
	tryRegister(&_CFWriteStreamGetError, lib, "CFWriteStreamGetError")
	tryRegister(&_CFWriteStreamGetStatus, lib, "CFWriteStreamGetStatus")
	tryRegister(&_CFWriteStreamGetTypeID, lib, "CFWriteStreamGetTypeID")
	tryRegister(&_CFWriteStreamOpen, lib, "CFWriteStreamOpen")
	tryRegister(&_CFWriteStreamScheduleWithRunLoop, lib, "CFWriteStreamScheduleWithRunLoop")
	tryRegister(&_CFWriteStreamSetClient, lib, "CFWriteStreamSetClient")
	tryRegister(&_CFWriteStreamSetDispatchQueue, lib, "CFWriteStreamSetDispatchQueue")
	tryRegister(&_CFWriteStreamSetProperty, lib, "CFWriteStreamSetProperty")
	tryRegister(&_CFWriteStreamUnscheduleFromRunLoop, lib, "CFWriteStreamUnscheduleFromRunLoop")
	tryRegister(&_CFWriteStreamWrite, lib, "CFWriteStreamWrite")
	tryRegister(&_CFXMLCreateStringByEscapingEntities, lib, "CFXMLCreateStringByEscapingEntities")
	tryRegister(&_CFXMLCreateStringByUnescapingEntities, lib, "CFXMLCreateStringByUnescapingEntities")
	tryRegister(&_CFXMLNodeCreate, lib, "CFXMLNodeCreate")
	tryRegister(&_CFXMLNodeCreateCopy, lib, "CFXMLNodeCreateCopy")
	tryRegister(&_CFXMLNodeGetInfoPtr, lib, "CFXMLNodeGetInfoPtr")
	tryRegister(&_CFXMLNodeGetString, lib, "CFXMLNodeGetString")
	tryRegister(&_CFXMLNodeGetTypeCode, lib, "CFXMLNodeGetTypeCode")
	tryRegister(&_CFXMLNodeGetTypeID, lib, "CFXMLNodeGetTypeID")
	tryRegister(&_CFXMLNodeGetVersion, lib, "CFXMLNodeGetVersion")
	tryRegister(&_CFXMLParserAbort, lib, "CFXMLParserAbort")
	tryRegister(&_CFXMLParserCopyErrorDescription, lib, "CFXMLParserCopyErrorDescription")
	tryRegister(&_CFXMLParserCreate, lib, "CFXMLParserCreate")
	tryRegister(&_CFXMLParserCreateWithDataFromURL, lib, "CFXMLParserCreateWithDataFromURL")
	tryRegister(&_CFXMLParserGetCallBacks, lib, "CFXMLParserGetCallBacks")
	tryRegister(&_CFXMLParserGetContext, lib, "CFXMLParserGetContext")
	tryRegister(&_CFXMLParserGetDocument, lib, "CFXMLParserGetDocument")
	tryRegister(&_CFXMLParserGetLineNumber, lib, "CFXMLParserGetLineNumber")
	tryRegister(&_CFXMLParserGetLocation, lib, "CFXMLParserGetLocation")
	tryRegister(&_CFXMLParserGetSourceURL, lib, "CFXMLParserGetSourceURL")
	tryRegister(&_CFXMLParserGetStatusCode, lib, "CFXMLParserGetStatusCode")
	tryRegister(&_CFXMLParserGetTypeID, lib, "CFXMLParserGetTypeID")
	tryRegister(&_CFXMLParserParse, lib, "CFXMLParserParse")
	tryRegister(&_CFXMLTreeCreateFromData, lib, "CFXMLTreeCreateFromData")
	tryRegister(&_CFXMLTreeCreateFromDataWithError, lib, "CFXMLTreeCreateFromDataWithError")
	tryRegister(&_CFXMLTreeCreateWithDataFromURL, lib, "CFXMLTreeCreateWithDataFromURL")
	tryRegister(&_CFXMLTreeCreateWithNode, lib, "CFXMLTreeCreateWithNode")
	tryRegister(&_CFXMLTreeCreateXMLData, lib, "CFXMLTreeCreateXMLData")
	tryRegister(&_CFXMLTreeGetNode, lib, "CFXMLTreeGetNode")
	tryRegister(&_inset, lib, "inset")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Releases a Core Foundation object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521153-cfrelease
func CFRelease(p0 unsafe.Pointer) {
	_CFRelease(p0)
	}


// Makes a newly-allocated Core Foundation object eligible for garbage collection. [Full Topic]
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521163-cfmakecollectable
func CFMakeCollectable(p0 unsafe.Pointer) unsafe.Pointer {
	return _CFMakeCollectable(p0)
	}


// Retains a Core Foundation object. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/corefoundation/1521269-cfretain
func CFRetain(p0 unsafe.Pointer) unsafe.Pointer {
	return _CFRetain(p0)
	}


// Adds a time interval, expressed as Gregorian units, to a given absolute time. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeAddGregorianUnits(_:_:_:)
func CFAbsoluteTimeAddGregorianUnits(at unsafe.Pointer, tz unsafe.Pointer, units unsafe.Pointer) unsafe.Pointer {
	return _CFAbsoluteTimeAddGregorianUnits(at, tz, units)
	}


// Returns the current system absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetCurrent()
func CFAbsoluteTimeGetCurrent() unsafe.Pointer {
	return _CFAbsoluteTimeGetCurrent()
	}


// Returns an integer representing the day of the week indicated by the specified absolute time. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetDayOfWeek(_:_:)
func CFAbsoluteTimeGetDayOfWeek(at unsafe.Pointer, tz unsafe.Pointer) unsafe.Pointer {
	return _CFAbsoluteTimeGetDayOfWeek(at, tz)
	}


// Returns an integer representing the day of the year indicated by the specified absolute time. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetDayOfYear(_:_:)
func CFAbsoluteTimeGetDayOfYear(at unsafe.Pointer, tz unsafe.Pointer) unsafe.Pointer {
	return _CFAbsoluteTimeGetDayOfYear(at, tz)
	}


// Computes the time difference between two specified absolute times and returns the result as an interval in Gregorian units. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetDifferenceAsGregorianUnits(_:_:_:_:)
func CFAbsoluteTimeGetDifferenceAsGregorianUnits(at1 unsafe.Pointer, at2 unsafe.Pointer, tz unsafe.Pointer, unitFlags unsafe.Pointer) unsafe.Pointer {
	return _CFAbsoluteTimeGetDifferenceAsGregorianUnits(at1, at2, tz, unitFlags)
	}


// Converts an absolute time value into a Gregorian date. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetGregorianDate(_:_:)
func CFAbsoluteTimeGetGregorianDate(at unsafe.Pointer, tz unsafe.Pointer) unsafe.Pointer {
	return _CFAbsoluteTimeGetGregorianDate(at, tz)
	}


// Returns an integer representing the week of the year indicated by the specified absolute time. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetWeekOfYear(_:_:)
func CFAbsoluteTimeGetWeekOfYear(at unsafe.Pointer, tz unsafe.Pointer) unsafe.Pointer {
	return _CFAbsoluteTimeGetWeekOfYear(at, tz)
	}


// Allocates memory using the specified allocator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocate(_:_:_:)
func CFAllocatorAllocate(allocator unsafe.Pointer, size unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorAllocate(allocator, size, hint)
	}


// CFAllocatorAllocateBytes is a CoreFoundation function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocateBytes(_:_:_:)
func CFAllocatorAllocateBytes(allocator unsafe.Pointer, size unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorAllocateBytes(allocator, size, hint)
	}


// CFAllocatorAllocateTyped is a CoreFoundation function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocateTyped(_:_:_:_:)
func CFAllocatorAllocateTyped(allocator unsafe.Pointer, size unsafe.Pointer, descriptor unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorAllocateTyped(allocator, size, descriptor, hint)
	}


// Creates an allocator object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorCreate(_:_:)
func CFAllocatorCreate(allocator unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorCreate(allocator, context)
	}


// CFAllocatorCreateWithZone is a CoreFoundation function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorCreateWithZone
func CFAllocatorCreateWithZone(allocator unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorCreateWithZone(allocator, zone)
	}


// Deallocates a block of memory with a given allocator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorDeallocate(_:_:)
func CFAllocatorDeallocate(allocator unsafe.Pointer, ptr unsafe.Pointer) {
	_CFAllocatorDeallocate(allocator, ptr)
	}


// Obtains the context of the specified allocator or of the default allocator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetContext(_:_:)
func CFAllocatorGetContext(allocator unsafe.Pointer, context unsafe.Pointer) {
	_CFAllocatorGetContext(allocator, context)
	}


// Gets the default allocator object for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetDefault()
func CFAllocatorGetDefault() unsafe.Pointer {
	return _CFAllocatorGetDefault()
	}


// Obtains the number of bytes likely to be allocated upon a specific request. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetPreferredSizeForSize(_:_:_:)
func CFAllocatorGetPreferredSizeForSize(allocator unsafe.Pointer, size unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorGetPreferredSizeForSize(allocator, size, hint)
	}


// Returns the type identifier for the CFAllocator opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetTypeID()
func CFAllocatorGetTypeID() unsafe.Pointer {
	return _CFAllocatorGetTypeID()
	}


// Reallocates memory using the specified allocator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocate(_:_:_:_:)
func CFAllocatorReallocate(allocator unsafe.Pointer, ptr unsafe.Pointer, newsize unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorReallocate(allocator, ptr, newsize, hint)
	}


// CFAllocatorReallocateBytes is a CoreFoundation function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocateBytes(_:_:_:_:)
func CFAllocatorReallocateBytes(allocator unsafe.Pointer, ptr unsafe.Pointer, newsize unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorReallocateBytes(allocator, ptr, newsize, hint)
	}


// CFAllocatorReallocateTyped is a CoreFoundation function. [Full Topic]
//
// Added in macOS 15.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocateTyped(_:_:_:_:_:)
func CFAllocatorReallocateTyped(allocator unsafe.Pointer, ptr unsafe.Pointer, newsize unsafe.Pointer, descriptor unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorReallocateTyped(allocator, ptr, newsize, descriptor, hint)
	}


// Sets the given allocator as the default for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorSetDefault(_:)
func CFAllocatorSetDefault(allocator unsafe.Pointer) {
	_CFAllocatorSetDefault(allocator)
	}


// Adds the values from one array to another array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayAppendArray(_:_:_:)
func CFArrayAppendArray(theArray unsafe.Pointer, otherArray unsafe.Pointer, otherRange unsafe.Pointer) {
	_CFArrayAppendArray(theArray, otherArray, otherRange)
	}


// Adds a value to an array giving it the new largest index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayAppendValue(_:_:)
func CFArrayAppendValue(theArray unsafe.Pointer, value unsafe.Pointer) {
	_CFArrayAppendValue(theArray, value)
	}


// Calls a function once for each element in range in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayApplyFunction(_:_:_:_:)
func CFArrayApplyFunction(theArray unsafe.Pointer, range_ unsafe.Pointer, applier unsafe.Pointer, context unsafe.Pointer) {
	_CFArrayApplyFunction(theArray, range_, applier, context)
	}


// Searches an array for a value using a binary search algorithm. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayBSearchValues(_:_:_:_:_:)
func CFArrayBSearchValues(theArray unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer, comparator unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFArrayBSearchValues(theArray, range_, value, comparator, context)
	}


// Reports whether or not a value is in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayContainsValue(_:_:_:)
func CFArrayContainsValue(theArray unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFArrayContainsValue(theArray, range_, value)
	}


// Creates a new immutable array with the given values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreate(_:_:_:_:)
func CFArrayCreate(allocator unsafe.Pointer, values unsafe.Pointer, numValues unsafe.Pointer, callBacks unsafe.Pointer) unsafe.Pointer {
	return _CFArrayCreate(allocator, values, numValues, callBacks)
	}


// Creates a new immutable array with the values from another array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateCopy(_:_:)
func CFArrayCreateCopy(allocator unsafe.Pointer, theArray unsafe.Pointer) unsafe.Pointer {
	return _CFArrayCreateCopy(allocator, theArray)
	}


// Creates a new empty mutable array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateMutable(_:_:_:)
func CFArrayCreateMutable(allocator unsafe.Pointer, capacity unsafe.Pointer, callBacks unsafe.Pointer) unsafe.Pointer {
	return _CFArrayCreateMutable(allocator, capacity, callBacks)
	}


// Creates a new mutable array with the values from another array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateMutableCopy(_:_:_:)
func CFArrayCreateMutableCopy(allocator unsafe.Pointer, capacity unsafe.Pointer, theArray unsafe.Pointer) unsafe.Pointer {
	return _CFArrayCreateMutableCopy(allocator, capacity, theArray)
	}


// Exchanges the values at two indices of an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayExchangeValuesAtIndices(_:_:_:)
func CFArrayExchangeValuesAtIndices(theArray unsafe.Pointer, idx1 unsafe.Pointer, idx2 unsafe.Pointer) {
	_CFArrayExchangeValuesAtIndices(theArray, idx1, idx2)
	}


// Returns the number of values currently in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetCount(_:)
func CFArrayGetCount(theArray unsafe.Pointer) unsafe.Pointer {
	return _CFArrayGetCount(theArray)
	}


// Counts the number of times a given value occurs in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetCountOfValue(_:_:_:)
func CFArrayGetCountOfValue(theArray unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFArrayGetCountOfValue(theArray, range_, value)
	}


// Searches an array forward for a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetFirstIndexOfValue(_:_:_:)
func CFArrayGetFirstIndexOfValue(theArray unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFArrayGetFirstIndexOfValue(theArray, range_, value)
	}


// Searches an array backward for a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetLastIndexOfValue(_:_:_:)
func CFArrayGetLastIndexOfValue(theArray unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFArrayGetLastIndexOfValue(theArray, range_, value)
	}


// Returns the type identifier for the CFArray opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetTypeID()
func CFArrayGetTypeID() unsafe.Pointer {
	return _CFArrayGetTypeID()
	}


// Retrieves a value at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetValueAtIndex(_:_:)
func CFArrayGetValueAtIndex(theArray unsafe.Pointer, idx unsafe.Pointer) unsafe.Pointer {
	return _CFArrayGetValueAtIndex(theArray, idx)
	}


// Fills a buffer with values from an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetValues(_:_:_:)
func CFArrayGetValues(theArray unsafe.Pointer, range_ unsafe.Pointer, values unsafe.Pointer) {
	_CFArrayGetValues(theArray, range_, values)
	}


// Inserts a value into an array at a given index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayInsertValueAtIndex(_:_:_:)
func CFArrayInsertValueAtIndex(theArray unsafe.Pointer, idx unsafe.Pointer, value unsafe.Pointer) {
	_CFArrayInsertValueAtIndex(theArray, idx, value)
	}


// Removes all the values from an array, making it empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayRemoveAllValues(_:)
func CFArrayRemoveAllValues(theArray unsafe.Pointer) {
	_CFArrayRemoveAllValues(theArray)
	}


// Removes the value at a given index from an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayRemoveValueAtIndex(_:_:)
func CFArrayRemoveValueAtIndex(theArray unsafe.Pointer, idx unsafe.Pointer) {
	_CFArrayRemoveValueAtIndex(theArray, idx)
	}


// Replaces a range of values in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayReplaceValues(_:_:_:_:)
func CFArrayReplaceValues(theArray unsafe.Pointer, range_ unsafe.Pointer, newValues unsafe.Pointer, newCount unsafe.Pointer) {
	_CFArrayReplaceValues(theArray, range_, newValues, newCount)
	}


// Changes the value at a given index in an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArraySetValueAtIndex(_:_:_:)
func CFArraySetValueAtIndex(theArray unsafe.Pointer, idx unsafe.Pointer, value unsafe.Pointer) {
	_CFArraySetValueAtIndex(theArray, idx, value)
	}


// Sorts the values in an array using a given comparison function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArraySortValues(_:_:_:_:)
func CFArraySortValues(theArray unsafe.Pointer, range_ unsafe.Pointer, comparator unsafe.Pointer, context unsafe.Pointer) {
	_CFArraySortValues(theArray, range_, comparator, context)
	}


// Defers internal consistency-checking and coalescing for a mutable attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringBeginEditing(_:)
func CFAttributedStringBeginEditing(aStr unsafe.Pointer) {
	_CFAttributedStringBeginEditing(aStr)
	}


// Creates an attributed string with specified string and attributes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreate(_:_:_:)
func CFAttributedStringCreate(alloc unsafe.Pointer, str unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreate(alloc, str, attributes)
	}


// Creates an immutable copy of an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateCopy(_:_:)
func CFAttributedStringCreateCopy(alloc unsafe.Pointer, aStr unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreateCopy(alloc, aStr)
	}


// Creates a mutable attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateMutable(_:_:)
func CFAttributedStringCreateMutable(alloc unsafe.Pointer, maxLength unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreateMutable(alloc, maxLength)
	}


// Creates a mutable copy of an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateMutableCopy(_:_:_:)
func CFAttributedStringCreateMutableCopy(alloc unsafe.Pointer, maxLength unsafe.Pointer, aStr unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreateMutableCopy(alloc, maxLength, aStr)
	}


// Creates a sub-attributed string from the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateWithSubstring(_:_:_:)
func CFAttributedStringCreateWithSubstring(alloc unsafe.Pointer, aStr unsafe.Pointer, range_ unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreateWithSubstring(alloc, aStr, range_)
	}


// Re-enables internal consistency-checking and coalescing for a mutable attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringEndEditing(_:)
func CFAttributedStringEndEditing(aStr unsafe.Pointer) {
	_CFAttributedStringEndEditing(aStr)
	}


// Returns the value of a given attribute of an attributed string at a specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttribute(_:_:_:_:)
func CFAttributedStringGetAttribute(aStr unsafe.Pointer, loc unsafe.Pointer, attrName unsafe.Pointer, effectiveRange unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetAttribute(aStr, loc, attrName, effectiveRange)
	}


// Returns the value of a given attribute of an attributed string at a specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributeAndLongestEffectiveRange(_:_:_:_:_:)
func CFAttributedStringGetAttributeAndLongestEffectiveRange(aStr unsafe.Pointer, loc unsafe.Pointer, attrName unsafe.Pointer, inRange unsafe.Pointer, longestEffectiveRange unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetAttributeAndLongestEffectiveRange(aStr, loc, attrName, inRange, longestEffectiveRange)
	}


// Returns the attributes of an attributed string at a specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributes(_:_:_:)
func CFAttributedStringGetAttributes(aStr unsafe.Pointer, loc unsafe.Pointer, effectiveRange unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetAttributes(aStr, loc, effectiveRange)
	}


// Returns the attributes of an attributed string at a specified location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributesAndLongestEffectiveRange(_:_:_:_:)
func CFAttributedStringGetAttributesAndLongestEffectiveRange(aStr unsafe.Pointer, loc unsafe.Pointer, inRange unsafe.Pointer, longestEffectiveRange unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetAttributesAndLongestEffectiveRange(aStr, loc, inRange, longestEffectiveRange)
	}


// CFAttributedStringGetBidiLevelsAndResolvedDirections is a CoreFoundation function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetBidiLevelsAndResolvedDirections(_:_:_:_:_:)
func CFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString unsafe.Pointer, range_ unsafe.Pointer, baseDirection unsafe.Pointer, bidiLevels unsafe.Pointer, baseDirections unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString, range_, baseDirection, bidiLevels, baseDirections)
	}


// Returns the length of the attributed string in characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetLength(_:)
func CFAttributedStringGetLength(aStr unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetLength(aStr)
	}


// Gets as a mutable string the string for an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetMutableString(_:)
func CFAttributedStringGetMutableString(aStr unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetMutableString(aStr)
	}


// CFAttributedStringGetStatisticalWritingDirections is a CoreFoundation function. [Full Topic]
//
// Added in macOS 26.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetStatisticalWritingDirections(_:_:_:_:_:)
func CFAttributedStringGetStatisticalWritingDirections(attributedString unsafe.Pointer, range_ unsafe.Pointer, baseDirection unsafe.Pointer, bidiLevels unsafe.Pointer, baseDirections unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetStatisticalWritingDirections(attributedString, range_, baseDirection, bidiLevels, baseDirections)
	}


// Returns the string for an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetString(_:)
func CFAttributedStringGetString(aStr unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetString(aStr)
	}


// Returns the type identifier for the CFAttributedString opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetTypeID()
func CFAttributedStringGetTypeID() unsafe.Pointer {
	return _CFAttributedStringGetTypeID()
	}


// Removes the value of a single attribute over a specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringRemoveAttribute(_:_:_:)
func CFAttributedStringRemoveAttribute(aStr unsafe.Pointer, range_ unsafe.Pointer, attrName unsafe.Pointer) {
	_CFAttributedStringRemoveAttribute(aStr, range_, attrName)
	}


// Replaces the attributed substring over a range with another attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringReplaceAttributedString(_:_:_:)
func CFAttributedStringReplaceAttributedString(aStr unsafe.Pointer, range_ unsafe.Pointer, replacement unsafe.Pointer) {
	_CFAttributedStringReplaceAttributedString(aStr, range_, replacement)
	}


// Modifies the string of an attributed string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringReplaceString(_:_:_:)
func CFAttributedStringReplaceString(aStr unsafe.Pointer, range_ unsafe.Pointer, replacement unsafe.Pointer) {
	_CFAttributedStringReplaceString(aStr, range_, replacement)
	}


// Sets the value of a single attribute over the specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringSetAttribute(_:_:_:_:)
func CFAttributedStringSetAttribute(aStr unsafe.Pointer, range_ unsafe.Pointer, attrName unsafe.Pointer, value unsafe.Pointer) {
	_CFAttributedStringSetAttribute(aStr, range_, attrName, value)
	}


// Sets the value of attributes of a mutable attributed string over a specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringSetAttributes(_:_:_:_:)
func CFAttributedStringSetAttributes(aStr unsafe.Pointer, range_ unsafe.Pointer, replacement unsafe.Pointer, clearOtherAttributes unsafe.Pointer) {
	_CFAttributedStringSetAttributes(aStr, range_, replacement, clearOtherAttributes)
	}


// CFAutorelease is a CoreFoundation function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAutorelease
func CFAutorelease(arg unsafe.Pointer) unsafe.Pointer {
	return _CFAutorelease(arg)
	}


// Adds a value to a mutable bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagAddValue(_:_:)
func CFBagAddValue(theBag unsafe.Pointer, value unsafe.Pointer) {
	_CFBagAddValue(theBag, value)
	}


// Calls a function once for each value in a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagApplyFunction(_:_:_:)
func CFBagApplyFunction(theBag unsafe.Pointer, applier unsafe.Pointer, context unsafe.Pointer) {
	_CFBagApplyFunction(theBag, applier, context)
	}


// Reports whether or not a value is in a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagContainsValue(_:_:)
func CFBagContainsValue(theBag unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagContainsValue(theBag, value)
	}


// Creates an immutable bag containing specified values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCreate(_:_:_:_:)
func CFBagCreate(allocator unsafe.Pointer, values unsafe.Pointer, numValues unsafe.Pointer, callBacks unsafe.Pointer) unsafe.Pointer {
	return _CFBagCreate(allocator, values, numValues, callBacks)
	}


// Creates an immutable bag with the values of another bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateCopy(_:_:)
func CFBagCreateCopy(allocator unsafe.Pointer, theBag unsafe.Pointer) unsafe.Pointer {
	return _CFBagCreateCopy(allocator, theBag)
	}


// Creates a new empty mutable bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateMutable(_:_:_:)
func CFBagCreateMutable(allocator unsafe.Pointer, capacity unsafe.Pointer, callBacks unsafe.Pointer) unsafe.Pointer {
	return _CFBagCreateMutable(allocator, capacity, callBacks)
	}


// Creates a new mutable bag with the values from another bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateMutableCopy(_:_:_:)
func CFBagCreateMutableCopy(allocator unsafe.Pointer, capacity unsafe.Pointer, theBag unsafe.Pointer) unsafe.Pointer {
	return _CFBagCreateMutableCopy(allocator, capacity, theBag)
	}


// Returns the number of values currently in a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetCount(_:)
func CFBagGetCount(theBag unsafe.Pointer) unsafe.Pointer {
	return _CFBagGetCount(theBag)
	}


// Returns the number of times a value occurs in a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetCountOfValue(_:_:)
func CFBagGetCountOfValue(theBag unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagGetCountOfValue(theBag, value)
	}


// Returns the type identifier for the CFBag opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetTypeID()
func CFBagGetTypeID() unsafe.Pointer {
	return _CFBagGetTypeID()
	}


// Returns a requested value from a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValue(_:_:)
func CFBagGetValue(theBag unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagGetValue(theBag, value)
	}


// Reports whether or not a value is in a bag, and returns that value indirectly if it exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValueIfPresent(_:_:_:)
func CFBagGetValueIfPresent(theBag unsafe.Pointer, candidate unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagGetValueIfPresent(theBag, candidate, value)
	}


// Fills a buffer with values from a bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValues(_:_:)
func CFBagGetValues(theBag unsafe.Pointer, values unsafe.Pointer) {
	_CFBagGetValues(theBag, values)
	}


// Removes all values from a mutable bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagRemoveAllValues(_:)
func CFBagRemoveAllValues(theBag unsafe.Pointer) {
	_CFBagRemoveAllValues(theBag)
	}


// Removes a value from a mutable bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagRemoveValue(_:_:)
func CFBagRemoveValue(theBag unsafe.Pointer, value unsafe.Pointer) {
	_CFBagRemoveValue(theBag, value)
	}


// Replaces a value in a mutable bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagReplaceValue(_:_:)
func CFBagReplaceValue(theBag unsafe.Pointer, value unsafe.Pointer) {
	_CFBagReplaceValue(theBag, value)
	}


// Sets a value in a mutable bag. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagSetValue(_:_:)
func CFBagSetValue(theBag unsafe.Pointer, value unsafe.Pointer) {
	_CFBagSetValue(theBag, value)
	}


// Adds a value to a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapAddValue(_:_:)
func CFBinaryHeapAddValue(heap unsafe.Pointer, value unsafe.Pointer) {
	_CFBinaryHeapAddValue(heap, value)
	}


// Iteratively applies a function to all the values in a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapApplyFunction(_:_:_:)
func CFBinaryHeapApplyFunction(heap unsafe.Pointer, applier unsafe.Pointer, context unsafe.Pointer) {
	_CFBinaryHeapApplyFunction(heap, applier, context)
	}


// Returns whether a given value is in a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapContainsValue(_:_:)
func CFBinaryHeapContainsValue(heap unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapContainsValue(heap, value)
	}


// Creates a new mutable or fixed-mutable binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCreate(_:_:_:_:)
func CFBinaryHeapCreate(allocator unsafe.Pointer, capacity unsafe.Pointer, callBacks unsafe.Pointer, compareContext unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapCreate(allocator, capacity, callBacks, compareContext)
	}


// Creates a new mutable or fixed-mutable binary heap with the values from a pre-existing binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCreateCopy(_:_:_:)
func CFBinaryHeapCreateCopy(allocator unsafe.Pointer, capacity unsafe.Pointer, heap unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapCreateCopy(allocator, capacity, heap)
	}


// Returns the number of values currently in a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetCount(_:)
func CFBinaryHeapGetCount(heap unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapGetCount(heap)
	}


// Counts the number of times a given value occurs in a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetCountOfValue(_:_:)
func CFBinaryHeapGetCountOfValue(heap unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapGetCountOfValue(heap, value)
	}


// Returns the minimum value in a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetMinimum(_:)
func CFBinaryHeapGetMinimum(heap unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapGetMinimum(heap)
	}


// Returns the minimum value in a binary heap, if present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetMinimumIfPresent(_:_:)
func CFBinaryHeapGetMinimumIfPresent(heap unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapGetMinimumIfPresent(heap, value)
	}


// Returns the type identifier of the opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetTypeID()
func CFBinaryHeapGetTypeID() unsafe.Pointer {
	return _CFBinaryHeapGetTypeID()
	}


// Copies all the values from a binary heap into a sorted C array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetValues(_:_:)
func CFBinaryHeapGetValues(heap unsafe.Pointer, values unsafe.Pointer) {
	_CFBinaryHeapGetValues(heap, values)
	}


// Removes all values from a binary heap, making it empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapRemoveAllValues(_:)
func CFBinaryHeapRemoveAllValues(heap unsafe.Pointer) {
	_CFBinaryHeapRemoveAllValues(heap)
	}


// Removes the minimum value from a binary heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapRemoveMinimumValue(_:)
func CFBinaryHeapRemoveMinimumValue(heap unsafe.Pointer) {
	_CFBinaryHeapRemoveMinimumValue(heap)
	}


// Returns whether a bit vector contains a particular bit value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorContainsBit(_:_:_:)
func CFBitVectorContainsBit(bv unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorContainsBit(bv, range_, value)
	}


// Creates an immutable bit vector from a block of memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreate(_:_:_:)
func CFBitVectorCreate(allocator unsafe.Pointer, bytes unsafe.Pointer, numBits unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorCreate(allocator, bytes, numBits)
	}


// Creates an immutable bit vector that is a copy of another bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateCopy(_:_:)
func CFBitVectorCreateCopy(allocator unsafe.Pointer, bv unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorCreateCopy(allocator, bv)
	}


// Creates a mutable bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateMutable(_:_:)
func CFBitVectorCreateMutable(allocator unsafe.Pointer, capacity unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorCreateMutable(allocator, capacity)
	}


// Creates a new mutable bit vector from a pre-existing bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateMutableCopy(_:_:_:)
func CFBitVectorCreateMutableCopy(allocator unsafe.Pointer, capacity unsafe.Pointer, bv unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorCreateMutableCopy(allocator, capacity, bv)
	}


// Flips a bit value in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorFlipBitAtIndex(_:_:)
func CFBitVectorFlipBitAtIndex(bv unsafe.Pointer, idx unsafe.Pointer) {
	_CFBitVectorFlipBitAtIndex(bv, idx)
	}


// Flips a range of bit values in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorFlipBits(_:_:)
func CFBitVectorFlipBits(bv unsafe.Pointer, range_ unsafe.Pointer) {
	_CFBitVectorFlipBits(bv, range_)
	}


// Returns the bit value at a given index in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetBitAtIndex(_:_:)
func CFBitVectorGetBitAtIndex(bv unsafe.Pointer, idx unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorGetBitAtIndex(bv, idx)
	}


// Returns the bit values in a range of indices in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetBits(_:_:_:)
func CFBitVectorGetBits(bv unsafe.Pointer, range_ unsafe.Pointer, bytes unsafe.Pointer) {
	_CFBitVectorGetBits(bv, range_, bytes)
	}


// Returns the number of bit values in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetCount(_:)
func CFBitVectorGetCount(bv unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorGetCount(bv)
	}


// Counts the number of times a certain bit value occurs within a range of bits in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetCountOfBit(_:_:_:)
func CFBitVectorGetCountOfBit(bv unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorGetCountOfBit(bv, range_, value)
	}


// Locates the first occurrence of a certain bit value within a range of bits in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetFirstIndexOfBit(_:_:_:)
func CFBitVectorGetFirstIndexOfBit(bv unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorGetFirstIndexOfBit(bv, range_, value)
	}


// Locates the last occurrence of a certain bit value within a range of bits in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetLastIndexOfBit(_:_:_:)
func CFBitVectorGetLastIndexOfBit(bv unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorGetLastIndexOfBit(bv, range_, value)
	}


// Returns the type identifier for the CFBitVector opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetTypeID()
func CFBitVectorGetTypeID() unsafe.Pointer {
	return _CFBitVectorGetTypeID()
	}


// Sets all bits in a bit vector to a particular value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetAllBits(_:_:)
func CFBitVectorSetAllBits(bv unsafe.Pointer, value unsafe.Pointer) {
	_CFBitVectorSetAllBits(bv, value)
	}


// Sets the value of a particular bit in a bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetBitAtIndex(_:_:_:)
func CFBitVectorSetBitAtIndex(bv unsafe.Pointer, idx unsafe.Pointer, value unsafe.Pointer) {
	_CFBitVectorSetBitAtIndex(bv, idx, value)
	}


// Sets a range of bits in a bit vector to a particular value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetBits(_:_:_:)
func CFBitVectorSetBits(bv unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) {
	_CFBitVectorSetBits(bv, range_, value)
	}


// Changes the size of a mutable bit vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetCount(_:_:)
func CFBitVectorSetCount(bv unsafe.Pointer, count unsafe.Pointer) {
	_CFBitVectorSetCount(bv, count)
	}


// Returns the Core Foundation type identifier for the CFBoolean opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBooleanGetTypeID()
func CFBooleanGetTypeID() unsafe.Pointer {
	return _CFBooleanGetTypeID()
	}


// Returns the value of a CFBoolean object as a standard C type . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBooleanGetValue(_:)
func CFBooleanGetValue(boolean unsafe.Pointer) unsafe.Pointer {
	return _CFBooleanGetValue(boolean)
	}


// Closes an open resource map for a bundle. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCloseBundleResourceMap(_:_:)
func CFBundleCloseBundleResourceMap(bundle unsafe.Pointer, refNum unsafe.Pointer) {
	_CFBundleCloseBundleResourceMap(bundle, refNum)
	}


// Returns the location of a bundle’s auxiliary executable code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyAuxiliaryExecutableURL(_:_:)
func CFBundleCopyAuxiliaryExecutableURL(bundle unsafe.Pointer, executableName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyAuxiliaryExecutableURL(bundle, executableName)
	}


// Returns the location of a bundle’s built in plug-in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBuiltInPlugInsURL(_:)
func CFBundleCopyBuiltInPlugInsURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyBuiltInPlugInsURL(bundle)
	}


// Returns an array containing a bundle’s localizations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBundleLocalizations(_:)
func CFBundleCopyBundleLocalizations(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyBundleLocalizations(bundle)
	}


// Returns the location of a bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBundleURL(_:)
func CFBundleCopyBundleURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyBundleURL(bundle)
	}


// Returns an array of CFNumbers representing the architectures a given bundle provides. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableArchitectures(_:)
func CFBundleCopyExecutableArchitectures(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyExecutableArchitectures(bundle)
	}


// Returns an array of CFNumbers representing the architectures a given URL provides. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableArchitecturesForURL(_:)
func CFBundleCopyExecutableArchitecturesForURL(url unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyExecutableArchitecturesForURL(url)
	}


// Returns the location of a bundle’s main executable code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableURL(_:)
func CFBundleCopyExecutableURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyExecutableURL(bundle)
	}


// Returns the information dictionary for a given URL location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyInfoDictionaryForURL(_:)
func CFBundleCopyInfoDictionaryForURL(url unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyInfoDictionaryForURL(url)
	}


// Returns a bundle’s information dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyInfoDictionaryInDirectory(_:)
func CFBundleCopyInfoDictionaryInDirectory(bundleURL unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyInfoDictionaryInDirectory(bundleURL)
	}


// Given an array of possible localizations and preferred locations, returns the one or more of them that CFBundle would use, without reference to the current application context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizationsForPreferences(_:_:)
func CFBundleCopyLocalizationsForPreferences(locArray unsafe.Pointer, prefArray unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyLocalizationsForPreferences(locArray, prefArray)
	}


// Returns an array containing the localizations for a bundle or executable at a particular location. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizationsForURL(_:)
func CFBundleCopyLocalizationsForURL(url unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyLocalizationsForURL(url)
	}


// Returns a localized string from a bundle’s strings file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizedString(_:_:_:_:)
func CFBundleCopyLocalizedString(bundle unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer, tableName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyLocalizedString(bundle, key, value, tableName)
	}


// Returns a localized string from a bundle’s strings file. [Full Topic]
//
// Added in macOS 15.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizedStringForLocalizations(_:_:_:_:_:)
func CFBundleCopyLocalizedStringForLocalizations(bundle unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer, tableName unsafe.Pointer, localizations unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyLocalizedStringForLocalizations(bundle, key, value, tableName, localizations)
	}


// Given an array of possible localizations, returns the one or more of them that CFBundle would use in the current application context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyPreferredLocalizationsFromArray(_:)
func CFBundleCopyPreferredLocalizationsFromArray(locArray unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyPreferredLocalizationsFromArray(locArray)
	}


// Returns the location of a bundle’s private Frameworks directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyPrivateFrameworksURL(_:)
func CFBundleCopyPrivateFrameworksURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyPrivateFrameworksURL(bundle)
	}


// Returns the location of a resource contained in the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURL(_:_:_:_:)
func CFBundleCopyResourceURL(bundle unsafe.Pointer, resourceName unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURL(bundle, resourceName, resourceType, subDirName)
	}


// Returns the location of a localized resource in a bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLForLocalization(_:_:_:_:_:)
func CFBundleCopyResourceURLForLocalization(bundle unsafe.Pointer, resourceName unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer, localizationName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURLForLocalization(bundle, resourceName, resourceType, subDirName, localizationName)
	}


// Returns the location of a resource contained in the specified bundle directory without requiring the creation of a CFBundle object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLInDirectory(_:_:_:_:)
func CFBundleCopyResourceURLInDirectory(bundleURL unsafe.Pointer, resourceName unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURLInDirectory(bundleURL, resourceName, resourceType, subDirName)
	}


// Assembles an array of URLs specifying all of the resources of the specified type found in a bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfType(_:_:_:)
func CFBundleCopyResourceURLsOfType(bundle unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURLsOfType(bundle, resourceType, subDirName)
	}


// Returns an array containing copies of the URL locations for a specified bundle, resource, and localization name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfTypeForLocalization(_:_:_:_:)
func CFBundleCopyResourceURLsOfTypeForLocalization(bundle unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer, localizationName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURLsOfTypeForLocalization(bundle, resourceType, subDirName, localizationName)
	}


// Returns an array of CFURL objects describing the locations of all resources in a bundle of the specified type without needing to create a CFBundle object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfTypeInDirectory(_:_:_:)
func CFBundleCopyResourceURLsOfTypeInDirectory(bundleURL unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURLsOfTypeInDirectory(bundleURL, resourceType, subDirName)
	}


// Returns the location of a bundle’s Resources directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourcesDirectoryURL(_:)
func CFBundleCopyResourcesDirectoryURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourcesDirectoryURL(bundle)
	}


// Returns the location of a bundle’s shared frameworks directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySharedFrameworksURL(_:)
func CFBundleCopySharedFrameworksURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopySharedFrameworksURL(bundle)
	}


// Returns the location of a bundle’s shared support files directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySharedSupportURL(_:)
func CFBundleCopySharedSupportURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopySharedSupportURL(bundle)
	}


// Returns the location of the bundle’s support files directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySupportFilesDirectoryURL(_:)
func CFBundleCopySupportFilesDirectoryURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopySupportFilesDirectoryURL(bundle)
	}


// Creates a CFBundle object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCreate(_:_:)
func CFBundleCreate(allocator unsafe.Pointer, bundleURL unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCreate(allocator, bundleURL)
	}


// Searches a directory and constructs an array of CFBundle objects from all valid bundles in the specified directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCreateBundlesFromDirectory(_:_:_:)
func CFBundleCreateBundlesFromDirectory(allocator unsafe.Pointer, directoryURL unsafe.Pointer, bundleType unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCreateBundlesFromDirectory(allocator, directoryURL, bundleType)
	}


// Returns an array containing all of the bundles currently open in the application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetAllBundles()
func CFBundleGetAllBundles() unsafe.Pointer {
	return _CFBundleGetAllBundles()
	}


// Locate a bundle given its program-defined identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetBundleWithIdentifier(_:)
func CFBundleGetBundleWithIdentifier(bundleID unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetBundleWithIdentifier(bundleID)
	}


// Returns a data pointer to a symbol of the given name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDataPointerForName(_:_:)
func CFBundleGetDataPointerForName(bundle unsafe.Pointer, symbolName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetDataPointerForName(bundle, symbolName)
	}


// Returns a C array of data pointer to symbols of the given names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDataPointersForNames(_:_:_:)
func CFBundleGetDataPointersForNames(bundle unsafe.Pointer, symbolNames unsafe.Pointer, stbl unsafe.Pointer) {
	_CFBundleGetDataPointersForNames(bundle, symbolNames, stbl)
	}


// Returns the bundle’s development region from the bundle’s information property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDevelopmentRegion(_:)
func CFBundleGetDevelopmentRegion(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetDevelopmentRegion(bundle)
	}


// Returns a pointer to a function in a bundle’s executable code using the function name as the search key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetFunctionPointerForName(_:_:)
func CFBundleGetFunctionPointerForName(bundle unsafe.Pointer, functionName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetFunctionPointerForName(bundle, functionName)
	}


// Constructs a function table containing pointers to all of the functions found in a bundle’s main executable code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetFunctionPointersForNames(_:_:_:)
func CFBundleGetFunctionPointersForNames(bundle unsafe.Pointer, functionNames unsafe.Pointer, ftbl unsafe.Pointer) {
	_CFBundleGetFunctionPointersForNames(bundle, functionNames, ftbl)
	}


// Returns the bundle identifier from a bundle’s information property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetIdentifier(_:)
func CFBundleGetIdentifier(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetIdentifier(bundle)
	}


// Returns a bundle’s information dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetInfoDictionary(_:)
func CFBundleGetInfoDictionary(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetInfoDictionary(bundle)
	}


// Returns a bundle’s localized information dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetLocalInfoDictionary(_:)
func CFBundleGetLocalInfoDictionary(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetLocalInfoDictionary(bundle)
	}


// Returns an application’s main bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetMainBundle()
func CFBundleGetMainBundle() unsafe.Pointer {
	return _CFBundleGetMainBundle()
	}


// Returns a bundle’s package type and creator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPackageInfo(_:_:_:)
func CFBundleGetPackageInfo(bundle unsafe.Pointer, packageType unsafe.Pointer, packageCreator unsafe.Pointer) {
	_CFBundleGetPackageInfo(bundle, packageType, packageCreator)
	}


// Returns a bundle’s package type and creator without having to create a CFBundle object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPackageInfoInDirectory(_:_:_:)
func CFBundleGetPackageInfoInDirectory(url unsafe.Pointer, packageType unsafe.Pointer, packageCreator unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetPackageInfoInDirectory(url, packageType, packageCreator)
	}


// Returns a bundle’s plug-in. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPlugIn(_:)
func CFBundleGetPlugIn(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetPlugIn(bundle)
	}


// Returns the type identifier for the CFBundle opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetTypeID()
func CFBundleGetTypeID() unsafe.Pointer {
	return _CFBundleGetTypeID()
	}


// Returns a value (localized if possible) from a bundle’s information dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetValueForInfoDictionaryKey(_:_:)
func CFBundleGetValueForInfoDictionaryKey(bundle unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetValueForInfoDictionaryKey(bundle, key)
	}


// Returns a bundle’s version number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetVersionNumber(_:)
func CFBundleGetVersionNumber(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetVersionNumber(bundle)
	}


// CFBundleIsArchitectureLoadable is a CoreFoundation function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsArchitectureLoadable(_:)
func CFBundleIsArchitectureLoadable(arch unsafe.Pointer) unsafe.Pointer {
	return _CFBundleIsArchitectureLoadable(arch)
	}


// CFBundleIsExecutableLoadable is a CoreFoundation function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoadable(_:)
func CFBundleIsExecutableLoadable(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleIsExecutableLoadable(bundle)
	}


// CFBundleIsExecutableLoadableForURL is a CoreFoundation function. [Full Topic]
//
// Added in macOS 11.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoadableForURL(_:)
func CFBundleIsExecutableLoadableForURL(url unsafe.Pointer) unsafe.Pointer {
	return _CFBundleIsExecutableLoadableForURL(url)
	}


// Obtains information about the load status for a bundle’s main executable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoaded(_:)
func CFBundleIsExecutableLoaded(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleIsExecutableLoaded(bundle)
	}


// Loads a bundle’s main executable code into memory and dynamically links it into the running application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleLoadExecutable(_:)
func CFBundleLoadExecutable(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleLoadExecutable(bundle)
	}


// Returns a Boolean value that indicates whether a given bundle is loaded, attempting to load it if necessary. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleLoadExecutableAndReturnError(_:_:)
func CFBundleLoadExecutableAndReturnError(bundle unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFBundleLoadExecutableAndReturnError(bundle, error_)
	}


// Opens the non-localized and localized resource files (if any) for a bundle in separate resource maps. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleOpenBundleResourceFiles(_:_:_:)
func CFBundleOpenBundleResourceFiles(bundle unsafe.Pointer, refNum unsafe.Pointer, localizedRefNum unsafe.Pointer) unsafe.Pointer {
	return _CFBundleOpenBundleResourceFiles(bundle, refNum, localizedRefNum)
	}


// Opens the non-localized and localized resource files (if any) for a bundle in a single resource map. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleOpenBundleResourceMap(_:)
func CFBundleOpenBundleResourceMap(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleOpenBundleResourceMap(bundle)
	}


// Returns a Boolean value that indicates whether a given bundle is loaded or appears to be loadable. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundlePreflightExecutable(_:_:)
func CFBundlePreflightExecutable(bundle unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFBundlePreflightExecutable(bundle, error_)
	}


// Unloads the main executable for the specified bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleUnloadExecutable(_:)
func CFBundleUnloadExecutable(bundle unsafe.Pointer) {
	_CFBundleUnloadExecutable(bundle)
	}


// Computes the absolute time when specified components are added to a given absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarAddComponents
func CFCalendarAddComponents(calendar unsafe.Pointer, at unsafe.Pointer, options unsafe.Pointer, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarAddComponents(calendar, at, options, componentDesc)
	}


// Computes the absolute time from components in a description string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarComposeAbsoluteTime
func CFCalendarComposeAbsoluteTime(calendar unsafe.Pointer, at unsafe.Pointer, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarComposeAbsoluteTime(calendar, at, componentDesc)
	}


// Returns a copy of the logical calendar for the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyCurrent()
func CFCalendarCopyCurrent() unsafe.Pointer {
	return _CFCalendarCopyCurrent()
	}


// Returns a locale object for a specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyLocale(_:)
func CFCalendarCopyLocale(calendar unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarCopyLocale(calendar)
	}


// Returns a time zone object for a specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyTimeZone(_:)
func CFCalendarCopyTimeZone(calendar unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarCopyTimeZone(calendar)
	}


// Returns a calendar object for the calendar identified by a calendar identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCreateWithIdentifier(_:_:)
func CFCalendarCreateWithIdentifier(allocator unsafe.Pointer, identifier unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarCreateWithIdentifier(allocator, identifier)
	}


// Computes the components which are indicated by the componentDesc description string for the given absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarDecomposeAbsoluteTime
func CFCalendarDecomposeAbsoluteTime(calendar unsafe.Pointer, at unsafe.Pointer, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarDecomposeAbsoluteTime(calendar, at, componentDesc)
	}


// Computes the difference between the two absolute times, in terms of specified calendrical components. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetComponentDifference
func CFCalendarGetComponentDifference(calendar unsafe.Pointer, startingAT unsafe.Pointer, resultAT unsafe.Pointer, options unsafe.Pointer, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetComponentDifference(calendar, startingAT, resultAT, options, componentDesc)
	}


// Returns the index of first weekday for a specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetFirstWeekday(_:)
func CFCalendarGetFirstWeekday(calendar unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetFirstWeekday(calendar)
	}


// Returns the given calendar’s identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetIdentifier(_:)
func CFCalendarGetIdentifier(calendar unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetIdentifier(calendar)
	}


// Returns the maximum range limits of the values that a specified unit can take on in a given calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMaximumRangeOfUnit(_:_:)
func CFCalendarGetMaximumRangeOfUnit(calendar unsafe.Pointer, unit unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetMaximumRangeOfUnit(calendar, unit)
	}


// Returns the minimum number of days in the first week of a specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMinimumDaysInFirstWeek(_:)
func CFCalendarGetMinimumDaysInFirstWeek(calendar unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetMinimumDaysInFirstWeek(calendar)
	}


// Returns the minimum range limits of the values that a specified unit can take on in a given calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMinimumRangeOfUnit(_:_:)
func CFCalendarGetMinimumRangeOfUnit(calendar unsafe.Pointer, unit unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetMinimumRangeOfUnit(calendar, unit)
	}


// Returns the ordinal number of a calendrical unit within a larger unit at a specified absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetOrdinalityOfUnit(_:_:_:_:)
func CFCalendarGetOrdinalityOfUnit(calendar unsafe.Pointer, smallerUnit unsafe.Pointer, biggerUnit unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetOrdinalityOfUnit(calendar, smallerUnit, biggerUnit, at)
	}


// Returns the range of values that one unit can take on within a larger unit during which a specific absolute time occurs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetRangeOfUnit(_:_:_:_:)
func CFCalendarGetRangeOfUnit(calendar unsafe.Pointer, smallerUnit unsafe.Pointer, biggerUnit unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetRangeOfUnit(calendar, smallerUnit, biggerUnit, at)
	}


// Returns by reference the start time and duration of a given calendar unit that contains a given absolute time. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetTimeRangeOfUnit(_:_:_:_:_:)
func CFCalendarGetTimeRangeOfUnit(calendar unsafe.Pointer, unit unsafe.Pointer, at unsafe.Pointer, startp unsafe.Pointer, tip unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetTimeRangeOfUnit(calendar, unit, at, startp, tip)
	}


// Returns the type identifier for the CFCalendar opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetTypeID()
func CFCalendarGetTypeID() unsafe.Pointer {
	return _CFCalendarGetTypeID()
	}


// Sets the first weekday for a calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetFirstWeekday(_:_:)
func CFCalendarSetFirstWeekday(calendar unsafe.Pointer, wkdy unsafe.Pointer) {
	_CFCalendarSetFirstWeekday(calendar, wkdy)
	}


// Sets the locale for a calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetLocale(_:_:)
func CFCalendarSetLocale(calendar unsafe.Pointer, locale unsafe.Pointer) {
	_CFCalendarSetLocale(calendar, locale)
	}


// Sets the minimum number of days in the first week of a specified calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetMinimumDaysInFirstWeek(_:_:)
func CFCalendarSetMinimumDaysInFirstWeek(calendar unsafe.Pointer, mwd unsafe.Pointer) {
	_CFCalendarSetMinimumDaysInFirstWeek(calendar, mwd)
	}


// Sets the time zone for a calendar. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetTimeZone(_:_:)
func CFCalendarSetTimeZone(calendar unsafe.Pointer, tz unsafe.Pointer) {
	_CFCalendarSetTimeZone(calendar, tz)
	}


// Adds a given range to a character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetAddCharactersInRange(_:_:)
func CFCharacterSetAddCharactersInRange(theSet unsafe.Pointer, theRange unsafe.Pointer) {
	_CFCharacterSetAddCharactersInRange(theSet, theRange)
	}


// Adds the characters in a given string to a character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetAddCharactersInString(_:_:)
func CFCharacterSetAddCharactersInString(theSet unsafe.Pointer, theString unsafe.Pointer) {
	_CFCharacterSetAddCharactersInString(theSet, theString)
	}


// Creates a new immutable data with the bitmap representation from the given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateBitmapRepresentation(_:_:)
func CFCharacterSetCreateBitmapRepresentation(alloc unsafe.Pointer, theSet unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateBitmapRepresentation(alloc, theSet)
	}


// Creates a new character set with the values from a given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateCopy(_:_:)
func CFCharacterSetCreateCopy(alloc unsafe.Pointer, theSet unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateCopy(alloc, theSet)
	}


// Creates a new immutable character set that is the invert of the specified character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateInvertedSet(_:_:)
func CFCharacterSetCreateInvertedSet(alloc unsafe.Pointer, theSet unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateInvertedSet(alloc, theSet)
	}


// Creates a new empty mutable character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateMutable(_:)
func CFCharacterSetCreateMutable(alloc unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateMutable(alloc)
	}


// Creates a new mutable character set with the values from another character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateMutableCopy(_:_:)
func CFCharacterSetCreateMutableCopy(alloc unsafe.Pointer, theSet unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateMutableCopy(alloc, theSet)
	}


// Creates a new immutable character set with the bitmap representation specified by given data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithBitmapRepresentation(_:_:)
func CFCharacterSetCreateWithBitmapRepresentation(alloc unsafe.Pointer, theData unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateWithBitmapRepresentation(alloc, theData)
	}


// Creates a new character set with the values from the given range of Unicode characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithCharactersInRange(_:_:)
func CFCharacterSetCreateWithCharactersInRange(alloc unsafe.Pointer, theRange unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateWithCharactersInRange(alloc, theRange)
	}


// Creates a new character set with the values in the given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithCharactersInString(_:_:)
func CFCharacterSetCreateWithCharactersInString(alloc unsafe.Pointer, theString unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateWithCharactersInString(alloc, theString)
	}


// Returns a predefined character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetGetPredefined(_:)
func CFCharacterSetGetPredefined(theSetIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetGetPredefined(theSetIdentifier)
	}


// Returns the type identifier of the CFCharacterSet opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetGetTypeID()
func CFCharacterSetGetTypeID() unsafe.Pointer {
	return _CFCharacterSetGetTypeID()
	}


// Reports whether or not a character set contains at least one member character in the specified plane. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetHasMemberInPlane(_:_:)
func CFCharacterSetHasMemberInPlane(theSet unsafe.Pointer, thePlane unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetHasMemberInPlane(theSet, thePlane)
	}


// Forms an intersection of two character sets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIntersect(_:_:)
func CFCharacterSetIntersect(theSet unsafe.Pointer, theOtherSet unsafe.Pointer) {
	_CFCharacterSetIntersect(theSet, theOtherSet)
	}


// Inverts the content of a given character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetInvert(_:)
func CFCharacterSetInvert(theSet unsafe.Pointer) {
	_CFCharacterSetInvert(theSet)
	}


// Reports whether or not a given Unicode character is in a character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsCharacterMember(_:_:)
func CFCharacterSetIsCharacterMember(theSet unsafe.Pointer, theChar unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetIsCharacterMember(theSet, theChar)
	}


// Reports whether or not a given UTF-32 character is in a character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsLongCharacterMember(_:_:)
func CFCharacterSetIsLongCharacterMember(theSet unsafe.Pointer, theChar unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetIsLongCharacterMember(theSet, theChar)
	}


// Reports whether or not a character set is a superset of another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsSupersetOfSet(_:_:)
func CFCharacterSetIsSupersetOfSet(theSet unsafe.Pointer, theOtherset unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetIsSupersetOfSet(theSet, theOtherset)
	}


// Removes a given range of Unicode characters from a character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetRemoveCharactersInRange(_:_:)
func CFCharacterSetRemoveCharactersInRange(theSet unsafe.Pointer, theRange unsafe.Pointer) {
	_CFCharacterSetRemoveCharactersInRange(theSet, theRange)
	}


// Removes the characters in a given string from a character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetRemoveCharactersInString(_:_:)
func CFCharacterSetRemoveCharactersInString(theSet unsafe.Pointer, theString unsafe.Pointer) {
	_CFCharacterSetRemoveCharactersInString(theSet, theString)
	}


// Forms the union of two character sets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetUnion(_:_:)
func CFCharacterSetUnion(theSet unsafe.Pointer, theOtherSet unsafe.Pointer) {
	_CFCharacterSetUnion(theSet, theOtherSet)
	}


// Returns a textual description of a Core Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCopyDescription(_:)
func CFCopyDescription(cf unsafe.Pointer) unsafe.Pointer {
	return _CFCopyDescription(cf)
	}


// CFCopyHomeDirectoryURL is a CoreFoundation function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCopyHomeDirectoryURL()
func CFCopyHomeDirectoryURL() unsafe.Pointer {
	return _CFCopyHomeDirectoryURL()
	}


// Returns a textual description of a Core Foundation type, as identified by its type ID, which can be used when debugging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCopyTypeIDDescription(_:)
func CFCopyTypeIDDescription(type_id unsafe.Pointer) unsafe.Pointer {
	return _CFCopyTypeIDDescription(type_id)
	}


// Appends the bytes from a byte buffer to the contents of a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataAppendBytes(_:_:_:)
func CFDataAppendBytes(theData unsafe.Pointer, bytes unsafe.Pointer, length unsafe.Pointer) {
	_CFDataAppendBytes(theData, bytes, length)
	}


// Creates an immutable CFData object using data copied from a specified byte buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreate(_:_:_:)
func CFDataCreate(allocator unsafe.Pointer, bytes unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreate(allocator, bytes, length)
	}


// Creates an immutable copy of a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateCopy(_:_:)
func CFDataCreateCopy(allocator unsafe.Pointer, theData unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreateCopy(allocator, theData)
	}


// Creates an empty CFMutableData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateMutable(_:_:)
func CFDataCreateMutable(allocator unsafe.Pointer, capacity unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreateMutable(allocator, capacity)
	}


// Creates a CFMutableData object by copying another CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateMutableCopy(_:_:_:)
func CFDataCreateMutableCopy(allocator unsafe.Pointer, capacity unsafe.Pointer, theData unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreateMutableCopy(allocator, capacity, theData)
	}


// Creates an immutable CFData object from an external (client-owned) byte buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateWithBytesNoCopy(_:_:_:_:)
func CFDataCreateWithBytesNoCopy(allocator unsafe.Pointer, bytes unsafe.Pointer, length unsafe.Pointer, bytesDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreateWithBytesNoCopy(allocator, bytes, length, bytesDeallocator)
	}


// Deletes the bytes in a CFMutableData object within a specified range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataDeleteBytes(_:_:)
func CFDataDeleteBytes(theData unsafe.Pointer, range_ unsafe.Pointer) {
	_CFDataDeleteBytes(theData, range_)
	}


// Finds and returns the range within a data object of the first occurrence of the given data, within a given range, subject to any given options. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataFind(_:_:_:_:)
func CFDataFind(theData unsafe.Pointer, dataToFind unsafe.Pointer, searchRange unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFDataFind(theData, dataToFind, searchRange, compareOptions)
	}


// Returns a read-only pointer to the bytes of a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetBytePtr(_:)
func CFDataGetBytePtr(theData unsafe.Pointer) unsafe.Pointer {
	return _CFDataGetBytePtr(theData)
	}


// Copies the byte contents of a CFData object to an external buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetBytes(_:_:_:)
func CFDataGetBytes(theData unsafe.Pointer, range_ unsafe.Pointer, buffer unsafe.Pointer) {
	_CFDataGetBytes(theData, range_, buffer)
	}


// Returns the number of bytes contained by a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetLength(_:)
func CFDataGetLength(theData unsafe.Pointer) unsafe.Pointer {
	return _CFDataGetLength(theData)
	}


// Returns a pointer to a mutable byte buffer of a CFMutableData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetMutableBytePtr(_:)
func CFDataGetMutableBytePtr(theData unsafe.Pointer) unsafe.Pointer {
	return _CFDataGetMutableBytePtr(theData)
	}


// Returns the type identifier for the CFData opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetTypeID()
func CFDataGetTypeID() unsafe.Pointer {
	return _CFDataGetTypeID()
	}


// Increases the length of a CFMutableData object’s internal byte buffer, zero-filling the extension to the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataIncreaseLength(_:_:)
func CFDataIncreaseLength(theData unsafe.Pointer, extraLength unsafe.Pointer) {
	_CFDataIncreaseLength(theData, extraLength)
	}


// Replaces those bytes in a CFMutableData object that fall within a specified range with other bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataReplaceBytes(_:_:_:_:)
func CFDataReplaceBytes(theData unsafe.Pointer, range_ unsafe.Pointer, newBytes unsafe.Pointer, newLength unsafe.Pointer) {
	_CFDataReplaceBytes(theData, range_, newBytes, newLength)
	}


// Resets the length of a CFMutableData object’s internal byte buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSetLength(_:_:)
func CFDataSetLength(theData unsafe.Pointer, length unsafe.Pointer) {
	_CFDataSetLength(theData, length)
	}


// Compares two objects and returns a comparison result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateCompare(_:_:_:)
func CFDateCompare(theDate unsafe.Pointer, otherDate unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFDateCompare(theDate, otherDate, context)
	}


// Creates a object given an absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateCreate(_:_:)
func CFDateCreate(allocator unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFDateCreate(allocator, at)
	}


// Returns a copy of a date formatter’s value for a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCopyProperty(_:_:)
func CFDateFormatterCopyProperty(formatter unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCopyProperty(formatter, key)
	}


// Creates a new CFDateFormatter object, localized to the given locale, which will format dates to the given date and time styles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreate(_:_:_:_:)
func CFDateFormatterCreate(allocator unsafe.Pointer, locale unsafe.Pointer, dateStyle unsafe.Pointer, timeStyle unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreate(allocator, locale, dateStyle, timeStyle)
	}


// Returns a localized date format string representing the given date format components arranged appropriately for the specified locale. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateDateFormatFromTemplate(_:_:_:_:)
func CFDateFormatterCreateDateFormatFromTemplate(allocator unsafe.Pointer, tmplate unsafe.Pointer, options unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreateDateFormatFromTemplate(allocator, tmplate, options, locale)
	}


// Returns a date object representing a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateDateFromString(_:_:_:_:)
func CFDateFormatterCreateDateFromString(allocator unsafe.Pointer, formatter unsafe.Pointer, string_ unsafe.Pointer, rangep unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreateDateFromString(allocator, formatter, string_, rangep)
	}


// CFDateFormatterCreateISO8601Formatter is a CoreFoundation function. [Full Topic]
//
// Added in macOS 10.12.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateISO8601Formatter(_:_:)
func CFDateFormatterCreateISO8601Formatter(allocator unsafe.Pointer, formatOptions unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreateISO8601Formatter(allocator, formatOptions)
	}


// Returns a string representation of the given absolute time using the specified date formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateStringWithAbsoluteTime(_:_:_:)
func CFDateFormatterCreateStringWithAbsoluteTime(allocator unsafe.Pointer, formatter unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreateStringWithAbsoluteTime(allocator, formatter, at)
	}


// Returns a string representation of the given date using the specified date formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateStringWithDate(_:_:_:)
func CFDateFormatterCreateStringWithDate(allocator unsafe.Pointer, formatter unsafe.Pointer, date unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreateStringWithDate(allocator, formatter, date)
	}


// Returns an absolute time object representing a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetAbsoluteTimeFromString(_:_:_:_:)
func CFDateFormatterGetAbsoluteTimeFromString(formatter unsafe.Pointer, string_ unsafe.Pointer, rangep unsafe.Pointer, atp unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetAbsoluteTimeFromString(formatter, string_, rangep, atp)
	}


// Returns the date style used to create the given date formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetDateStyle(_:)
func CFDateFormatterGetDateStyle(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetDateStyle(formatter)
	}


// Returns a format string for the given date formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetFormat(_:)
func CFDateFormatterGetFormat(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetFormat(formatter)
	}


// Returns the locale object used to create the given date formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetLocale(_:)
func CFDateFormatterGetLocale(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetLocale(formatter)
	}


// Returns the time style used to create the given date formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetTimeStyle(_:)
func CFDateFormatterGetTimeStyle(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetTimeStyle(formatter)
	}


// Returns the type identifier for CFDateFormatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetTypeID()
func CFDateFormatterGetTypeID() unsafe.Pointer {
	return _CFDateFormatterGetTypeID()
	}


// Sets the format string of the given date formatter to the specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterSetFormat(_:_:)
func CFDateFormatterSetFormat(formatter unsafe.Pointer, formatString unsafe.Pointer) {
	_CFDateFormatterSetFormat(formatter, formatString)
	}


// Sets a date formatter property using a key-value pair. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterSetProperty(_:_:_:)
func CFDateFormatterSetProperty(formatter unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_CFDateFormatterSetProperty(formatter, key, value)
	}


// Returns a object’s absolute time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateGetAbsoluteTime(_:)
func CFDateGetAbsoluteTime(theDate unsafe.Pointer) unsafe.Pointer {
	return _CFDateGetAbsoluteTime(theDate)
	}


// Returns the number of elapsed seconds between the given objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateGetTimeIntervalSinceDate(_:_:)
func CFDateGetTimeIntervalSinceDate(theDate unsafe.Pointer, otherDate unsafe.Pointer) unsafe.Pointer {
	return _CFDateGetTimeIntervalSinceDate(theDate, otherDate)
	}


// Returns the type identifier for the opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateGetTypeID()
func CFDateGetTypeID() unsafe.Pointer {
	return _CFDateGetTypeID()
	}


// Adds a key-value pair to a dictionary if the specified key is not already present. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryAddValue(_:_:_:)
func CFDictionaryAddValue(theDict unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_CFDictionaryAddValue(theDict, key, value)
	}


// Calls a function once for each key-value pair in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryApplyFunction(_:_:_:)
func CFDictionaryApplyFunction(theDict unsafe.Pointer, applier unsafe.Pointer, context unsafe.Pointer) {
	_CFDictionaryApplyFunction(theDict, applier, context)
	}


// Returns a Boolean value that indicates whether a given key is in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryContainsKey(_:_:)
func CFDictionaryContainsKey(theDict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryContainsKey(theDict, key)
	}


// Returns a Boolean value that indicates whether a given value is in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryContainsValue(_:_:)
func CFDictionaryContainsValue(theDict unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryContainsValue(theDict, value)
	}


// Creates an immutable dictionary containing the specified key-value pairs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreate(_:_:_:_:_:_:)
func CFDictionaryCreate(allocator unsafe.Pointer, keys unsafe.Pointer, values unsafe.Pointer, numValues unsafe.Pointer, keyCallBacks unsafe.Pointer, valueCallBacks unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryCreate(allocator, keys, values, numValues, keyCallBacks, valueCallBacks)
	}


// Creates and returns a new immutable dictionary with the key-value pairs of another dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateCopy(_:_:)
func CFDictionaryCreateCopy(allocator unsafe.Pointer, theDict unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryCreateCopy(allocator, theDict)
	}


// Creates a new mutable dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateMutable(_:_:_:_:)
func CFDictionaryCreateMutable(allocator unsafe.Pointer, capacity unsafe.Pointer, keyCallBacks unsafe.Pointer, valueCallBacks unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryCreateMutable(allocator, capacity, keyCallBacks, valueCallBacks)
	}


// Creates a new mutable dictionary with the key-value pairs from another dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateMutableCopy(_:_:_:)
func CFDictionaryCreateMutableCopy(allocator unsafe.Pointer, capacity unsafe.Pointer, theDict unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryCreateMutableCopy(allocator, capacity, theDict)
	}


// Returns the number of key-value pairs in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCount(_:)
func CFDictionaryGetCount(theDict unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetCount(theDict)
	}


// Returns the number of times a key occurs in a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCountOfKey(_:_:)
func CFDictionaryGetCountOfKey(theDict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetCountOfKey(theDict, key)
	}


// Counts the number of times a given value occurs in the dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCountOfValue(_:_:)
func CFDictionaryGetCountOfValue(theDict unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetCountOfValue(theDict, value)
	}


// Fills two buffers with the keys and values from a dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetKeysAndValues(_:_:_:)
func CFDictionaryGetKeysAndValues(theDict unsafe.Pointer, keys unsafe.Pointer, values unsafe.Pointer) {
	_CFDictionaryGetKeysAndValues(theDict, keys, values)
	}


// Returns the type identifier for the CFDictionary opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetTypeID()
func CFDictionaryGetTypeID() unsafe.Pointer {
	return _CFDictionaryGetTypeID()
	}


// Returns the value associated with a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetValue(_:_:)
func CFDictionaryGetValue(theDict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetValue(theDict, key)
	}


// Returns a Boolean value that indicates whether a given value for a given key is in a dictionary, and returns that value indirectly if it exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetValueIfPresent(_:_:_:)
func CFDictionaryGetValueIfPresent(theDict unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetValueIfPresent(theDict, key, value)
	}


// Removes all the key-value pairs from a dictionary, making it empty. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryRemoveAllValues(_:)
func CFDictionaryRemoveAllValues(theDict unsafe.Pointer) {
	_CFDictionaryRemoveAllValues(theDict)
	}


// Removes a key-value pair. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryRemoveValue(_:_:)
func CFDictionaryRemoveValue(theDict unsafe.Pointer, key unsafe.Pointer) {
	_CFDictionaryRemoveValue(theDict, key)
	}


// Replaces a value corresponding to a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryReplaceValue(_:_:_:)
func CFDictionaryReplaceValue(theDict unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_CFDictionaryReplaceValue(theDict, key, value)
	}


// Sets the value corresponding to a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionarySetValue(_:_:_:)
func CFDictionarySetValue(theDict unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_CFDictionarySetValue(theDict, key, value)
	}


// Determines whether two Core Foundation objects are considered equal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFEqual(_:_:)
func CFEqual(cf1 unsafe.Pointer, cf2 unsafe.Pointer) unsafe.Pointer {
	return _CFEqual(cf1, cf2)
	}


// Returns a human-presentable description for a given error. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyDescription(_:)
func CFErrorCopyDescription(err unsafe.Pointer) unsafe.Pointer {
	return _CFErrorCopyDescription(err)
	}


// Returns a human-presentable failure reason for a given error. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyFailureReason(_:)
func CFErrorCopyFailureReason(err unsafe.Pointer) unsafe.Pointer {
	return _CFErrorCopyFailureReason(err)
	}


// Returns a human presentable recovery suggestion for a given error. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyRecoverySuggestion(_:)
func CFErrorCopyRecoverySuggestion(err unsafe.Pointer) unsafe.Pointer {
	return _CFErrorCopyRecoverySuggestion(err)
	}


// Returns the user info dictionary for a given CFError. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyUserInfo(_:)
func CFErrorCopyUserInfo(err unsafe.Pointer) unsafe.Pointer {
	return _CFErrorCopyUserInfo(err)
	}


// Creates a new CFError object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCreate(_:_:_:_:)
func CFErrorCreate(allocator unsafe.Pointer, domain unsafe.Pointer, code unsafe.Pointer, userInfo unsafe.Pointer) unsafe.Pointer {
	return _CFErrorCreate(allocator, domain, code, userInfo)
	}


// Creates a new CFError object using given keys and values to create the user info dictionary. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCreateWithUserInfoKeysAndValues(_:_:_:_:_:_:)
func CFErrorCreateWithUserInfoKeysAndValues(allocator unsafe.Pointer, domain unsafe.Pointer, code unsafe.Pointer, userInfoKeys unsafe.Pointer, userInfoValues unsafe.Pointer, numUserInfoValues unsafe.Pointer) unsafe.Pointer {
	return _CFErrorCreateWithUserInfoKeysAndValues(allocator, domain, code, userInfoKeys, userInfoValues, numUserInfoValues)
	}


// Returns the error code for a given CFError. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetCode(_:)
func CFErrorGetCode(err unsafe.Pointer) unsafe.Pointer {
	return _CFErrorGetCode(err)
	}


// Returns the error domain for a given CFError. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetDomain(_:)
func CFErrorGetDomain(err unsafe.Pointer) unsafe.Pointer {
	return _CFErrorGetDomain(err)
	}


// Returns the type identifier for the CFError opaque type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetTypeID()
func CFErrorGetTypeID() unsafe.Pointer {
	return _CFErrorGetTypeID()
	}


// Creates a new CFFileDescriptor. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorCreate(_:_:_:_:_:)
func CFFileDescriptorCreate(allocator unsafe.Pointer, fd unsafe.Pointer, closeOnInvalidate unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFFileDescriptorCreate(allocator, fd, closeOnInvalidate, callout, context)
	}


// Creates a new runloop source for a given CFFileDescriptor. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorCreateRunLoopSource(_:_:_:)
func CFFileDescriptorCreateRunLoopSource(allocator unsafe.Pointer, f unsafe.Pointer, order unsafe.Pointer) unsafe.Pointer {
	return _CFFileDescriptorCreateRunLoopSource(allocator, f, order)
	}


// Disables callbacks for a given CFFileDescriptor. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorDisableCallBacks(_:_:)
func CFFileDescriptorDisableCallBacks(f unsafe.Pointer, callBackTypes unsafe.Pointer) {
	_CFFileDescriptorDisableCallBacks(f, callBackTypes)
	}


// Enables callbacks for a given CFFileDescriptor. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorEnableCallBacks(_:_:)
func CFFileDescriptorEnableCallBacks(f unsafe.Pointer, callBackTypes unsafe.Pointer) {
	_CFFileDescriptorEnableCallBacks(f, callBackTypes)
	}


// Gets the context for a given CFFileDescriptor. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetContext(_:_:)
func CFFileDescriptorGetContext(f unsafe.Pointer, context unsafe.Pointer) {
	_CFFileDescriptorGetContext(f, context)
	}


// Returns the native file descriptor for a given CFFileDescriptor. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetNativeDescriptor(_:)
func CFFileDescriptorGetNativeDescriptor(f unsafe.Pointer) unsafe.Pointer {
	return _CFFileDescriptorGetNativeDescriptor(f)
	}


// Returns the type identifier for the CFFileDescriptor opaque type. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetTypeID()
func CFFileDescriptorGetTypeID() unsafe.Pointer {
	return _CFFileDescriptorGetTypeID()
	}


// Invalidates a CFFileDescriptor object. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorInvalidate(_:)
func CFFileDescriptorInvalidate(f unsafe.Pointer) {
	_CFFileDescriptorInvalidate(f)
	}


// Returns a Boolean value that indicates whether the native file descriptor for a given CFFileDescriptor is valid. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorIsValid(_:)
func CFFileDescriptorIsValid(f unsafe.Pointer) unsafe.Pointer {
	return _CFFileDescriptorIsValid(f)
	}


// Clears properties from a object. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearProperties(_:_:)
func CFFileSecurityClearProperties(fileSec unsafe.Pointer, clearPropertyMask unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityClearProperties(fileSec, clearPropertyMask)
	}


// Copies the access control list associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyAccessControlList(_:_:)
func CFFileSecurityCopyAccessControlList(fileSec unsafe.Pointer, accessControlList unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityCopyAccessControlList(fileSec, accessControlList)
	}


// Copies the group UUID associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyGroupUUID(_:_:)
func CFFileSecurityCopyGroupUUID(fileSec unsafe.Pointer, groupUUID unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityCopyGroupUUID(fileSec, groupUUID)
	}


// Copies the owner UUID associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyOwnerUUID(_:_:)
func CFFileSecurityCopyOwnerUUID(fileSec unsafe.Pointer, ownerUUID unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityCopyOwnerUUID(fileSec, ownerUUID)
	}


// Creates a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCreate(_:)
func CFFileSecurityCreate(allocator unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityCreate(allocator)
	}


// Creates a copy of a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCreateCopy(_:_:)
func CFFileSecurityCreateCopy(allocator unsafe.Pointer, fileSec unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityCreateCopy(allocator, fileSec)
	}


// Gets the group ID associated with a object [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetGroup(_:_:)
func CFFileSecurityGetGroup(fileSec unsafe.Pointer, group unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityGetGroup(fileSec, group)
	}


// Gets the file mode associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetMode(_:_:)
func CFFileSecurityGetMode(fileSec unsafe.Pointer, mode unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityGetMode(fileSec, mode)
	}


// Gets the owner ID associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetOwner(_:_:)
func CFFileSecurityGetOwner(fileSec unsafe.Pointer, owner unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityGetOwner(fileSec, owner)
	}


// Returns the type identifier for the opaque type. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetTypeID()
func CFFileSecurityGetTypeID() unsafe.Pointer {
	return _CFFileSecurityGetTypeID()
	}


// Sets the access control list associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetAccessControlList(_:_:)
func CFFileSecuritySetAccessControlList(fileSec unsafe.Pointer, accessControlList unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecuritySetAccessControlList(fileSec, accessControlList)
	}


// Sets the group ID associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetGroup(_:_:)
func CFFileSecuritySetGroup(fileSec unsafe.Pointer, group unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecuritySetGroup(fileSec, group)
	}


// Sets the group UUID associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetGroupUUID(_:_:)
func CFFileSecuritySetGroupUUID(fileSec unsafe.Pointer, groupUUID unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecuritySetGroupUUID(fileSec, groupUUID)
	}


// Sets the file mode associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetMode(_:_:)
func CFFileSecuritySetMode(fileSec unsafe.Pointer, mode unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecuritySetMode(fileSec, mode)
	}


// Sets the owner ID associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetOwner(_:_:)
func CFFileSecuritySetOwner(fileSec unsafe.Pointer, owner unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecuritySetOwner(fileSec, owner)
	}


// Sets the owner UUID associated with a object. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetOwnerUUID(_:_:)
func CFFileSecuritySetOwnerUUID(fileSec unsafe.Pointer, ownerUUID unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecuritySetOwnerUUID(fileSec, ownerUUID)
	}


// Returns the allocator used to allocate a Core Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGetAllocator(_:)
func CFGetAllocator(cf unsafe.Pointer) unsafe.Pointer {
	return _CFGetAllocator(cf)
	}


// Returns the reference count of a Core Foundation object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGetRetainCount(_:)
func CFGetRetainCount(cf unsafe.Pointer) unsafe.Pointer {
	return _CFGetRetainCount(cf)
	}


// Returns the unique identifier of an opaque type to which a Core Foundation object belongs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGetTypeID(_:)
func CFGetTypeID(cf unsafe.Pointer) unsafe.Pointer {
	return _CFGetTypeID(cf)
	}


// Converts a Gregorian date value into an absolute time value. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianDateGetAbsoluteTime(_:_:)
func CFGregorianDateGetAbsoluteTime(gdate unsafe.Pointer, tz unsafe.Pointer) unsafe.Pointer {
	return _CFGregorianDateGetAbsoluteTime(gdate, tz)
	}


// Checks the specified fields of a CFGregorianDate structure for valid values. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianDateIsValid(_:_:)
func CFGregorianDateIsValid(gdate unsafe.Pointer, unitFlags unsafe.Pointer) unsafe.Pointer {
	return _CFGregorianDateIsValid(gdate, unitFlags)
	}


// Returns a code that can be used to identify an object in a hashing structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFHash(_:)
func CFHash(cf unsafe.Pointer) unsafe.Pointer {
	return _CFHash(cf)
	}


// Returns an array of CFString objects that represents all locales for which locale data is available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyAvailableLocaleIdentifiers()
func CFLocaleCopyAvailableLocaleIdentifiers() unsafe.Pointer {
	return _CFLocaleCopyAvailableLocaleIdentifiers()
	}


// Returns an array of strings that represents ISO currency codes for currencies in common use. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyCommonISOCurrencyCodes()
func CFLocaleCopyCommonISOCurrencyCodes() unsafe.Pointer {
	return _CFLocaleCopyCommonISOCurrencyCodes()
	}


// Returns a copy of the logical locale for the current user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyCurrent()
func CFLocaleCopyCurrent() unsafe.Pointer {
	return _CFLocaleCopyCurrent()
	}


// Returns the display name for the given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyDisplayNameForPropertyValue(_:_:_:)
func CFLocaleCopyDisplayNameForPropertyValue(displayLocale unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCopyDisplayNameForPropertyValue(displayLocale, key, value)
	}


// Returns an array of CFString objects that represents all known legal ISO country codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOCountryCodes()
func CFLocaleCopyISOCountryCodes() unsafe.Pointer {
	return _CFLocaleCopyISOCountryCodes()
	}


// Returns an array of CFString objects that represents all known legal ISO currency codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOCurrencyCodes()
func CFLocaleCopyISOCurrencyCodes() unsafe.Pointer {
	return _CFLocaleCopyISOCurrencyCodes()
	}


// Returns an array of CFString objects that represents all known legal ISO language codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOLanguageCodes()
func CFLocaleCopyISOLanguageCodes() unsafe.Pointer {
	return _CFLocaleCopyISOLanguageCodes()
	}


// Returns the array of canonicalized language IDs that the user prefers. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyPreferredLanguages()
func CFLocaleCopyPreferredLanguages() unsafe.Pointer {
	return _CFLocaleCopyPreferredLanguages()
	}


// Creates a locale for the given arbitrary locale identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreate(_:_:)
func CFLocaleCreate(allocator unsafe.Pointer, localeIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreate(allocator, localeIdentifier)
	}


// Returns a canonical language identifier by mapping an arbitrary locale identification string to the canonical identifier [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLanguageIdentifierFromString(_:_:)
func CFLocaleCreateCanonicalLanguageIdentifierFromString(allocator unsafe.Pointer, localeIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateCanonicalLanguageIdentifierFromString(allocator, localeIdentifier)
	}


// Returns a canonical locale identifier from given language and region codes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(_:_:_:)
func CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator unsafe.Pointer, lcode unsafe.Pointer, rcode unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator, lcode, rcode)
	}


// Returns a canonical locale identifier by mapping an arbitrary locale identification string to the canonical identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLocaleIdentifierFromString(_:_:)
func CFLocaleCreateCanonicalLocaleIdentifierFromString(allocator unsafe.Pointer, localeIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateCanonicalLocaleIdentifierFromString(allocator, localeIdentifier)
	}


// Returns a dictionary containing the result from parsing a locale ID consisting of language, script, country or region, variant, and keyword/value pairs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateComponentsFromLocaleIdentifier(_:_:)
func CFLocaleCreateComponentsFromLocaleIdentifier(allocator unsafe.Pointer, localeID unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateComponentsFromLocaleIdentifier(allocator, localeID)
	}


// Returns a copy of a locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCopy(_:_:)
func CFLocaleCreateCopy(allocator unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateCopy(allocator, locale)
	}


// Returns a locale identifier consisting of language, script, country or region, variant, and keyword/value pairs derived from a dictionary containing the source information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateLocaleIdentifierFromComponents(_:_:)
func CFLocaleCreateLocaleIdentifierFromComponents(allocator unsafe.Pointer, dictionary unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateLocaleIdentifierFromComponents(allocator, dictionary)
	}


// Returns a locale identifier from a Windows locale code. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(_:_:)
func CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator unsafe.Pointer, lcid unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator, lcid)
	}


// Returns the given locale’s identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetIdentifier(_:)
func CFLocaleGetIdentifier(locale unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleGetIdentifier(locale)
	}


// Returns the character direction for the specified ISO language code. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetLanguageCharacterDirection(_:)
func CFLocaleGetLanguageCharacterDirection(isoLangCode unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleGetLanguageCharacterDirection(isoLangCode)
	}


// Returns the line direction for the specified ISO language code. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetLanguageLineDirection(_:)
func CFLocaleGetLanguageLineDirection(isoLangCode unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleGetLanguageLineDirection(isoLangCode)
	}


// Returns the root, canonical locale. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetSystem()
func CFLocaleGetSystem() unsafe.Pointer {
	return _CFLocaleGetSystem()
	}


// Returns the type identifier for the CFLocale opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetTypeID()
func CFLocaleGetTypeID() unsafe.Pointer {
	return _CFLocaleGetTypeID()
	}


// Returns the corresponding value for the given key of a locale’s key-value pair. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetValue(_:_:)
func CFLocaleGetValue(locale unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleGetValue(locale, key)
	}


// Returns a Windows locale code from the locale identifier. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(_:)
func CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier)
	}


// Creates a CFMachPort object with a new Mach port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreate(_:_:_:_:)
func CFMachPortCreate(allocator unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer, shouldFreeInfo unsafe.Pointer) unsafe.Pointer {
	return _CFMachPortCreate(allocator, callout, context, shouldFreeInfo)
	}


// Creates a CFRunLoopSource object for a CFMachPort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreateRunLoopSource(_:_:_:)
func CFMachPortCreateRunLoopSource(allocator unsafe.Pointer, port unsafe.Pointer, order unsafe.Pointer) unsafe.Pointer {
	return _CFMachPortCreateRunLoopSource(allocator, port, order)
	}


// Creates a CFMachPort object for a pre-existing native Mach port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreateWithPort(_:_:_:_:_:)
func CFMachPortCreateWithPort(allocator unsafe.Pointer, portNum unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer, shouldFreeInfo unsafe.Pointer) unsafe.Pointer {
	return _CFMachPortCreateWithPort(allocator, portNum, callout, context, shouldFreeInfo)
	}


// Returns the context information for a CFMachPort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetContext(_:_:)
func CFMachPortGetContext(port unsafe.Pointer, context unsafe.Pointer) {
	_CFMachPortGetContext(port, context)
	}


// Returns the invalidation callback function for a CFMachPort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetInvalidationCallBack(_:)
func CFMachPortGetInvalidationCallBack(port unsafe.Pointer) unsafe.Pointer {
	return _CFMachPortGetInvalidationCallBack(port)
	}


// Returns the native Mach port represented by a CFMachPort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetPort(_:)
func CFMachPortGetPort(port unsafe.Pointer) unsafe.Pointer {
	return _CFMachPortGetPort(port)
	}


// Returns the type identifier for the CFMachPort opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetTypeID()
func CFMachPortGetTypeID() unsafe.Pointer {
	return _CFMachPortGetTypeID()
	}


// Invalidates a CFMachPort object, stopping it from receiving any more messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortInvalidate(_:)
func CFMachPortInvalidate(port unsafe.Pointer) {
	_CFMachPortInvalidate(port)
	}


// Returns a Boolean value that indicates whether a CFMachPort object is valid and able to receive messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortIsValid(_:)
func CFMachPortIsValid(port unsafe.Pointer) unsafe.Pointer {
	return _CFMachPortIsValid(port)
	}


// Sets the callback function invoked when a CFMachPort object is invalidated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortSetInvalidationCallBack(_:_:)
func CFMachPortSetInvalidationCallBack(port unsafe.Pointer, callout unsafe.Pointer) {
	_CFMachPortSetInvalidationCallBack(port, callout)
	}


// Returns a local CFMessagePort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateLocal(_:_:_:_:_:)
func CFMessagePortCreateLocal(allocator unsafe.Pointer, name unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer, shouldFreeInfo unsafe.Pointer) unsafe.Pointer {
	return _CFMessagePortCreateLocal(allocator, name, callout, context, shouldFreeInfo)
	}


// Returns a CFMessagePort object connected to a remote port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateRemote(_:_:)
func CFMessagePortCreateRemote(allocator unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _CFMessagePortCreateRemote(allocator, name)
	}


// Creates a CFRunLoopSource object for a CFMessagePort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateRunLoopSource(_:_:_:)
func CFMessagePortCreateRunLoopSource(allocator unsafe.Pointer, local unsafe.Pointer, order unsafe.Pointer) unsafe.Pointer {
	return _CFMessagePortCreateRunLoopSource(allocator, local, order)
	}


// Returns the context information for a CFMessagePort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetContext(_:_:)
func CFMessagePortGetContext(ms unsafe.Pointer, context unsafe.Pointer) {
	_CFMessagePortGetContext(ms, context)
	}


// Returns the invalidation callback function for a CFMessagePort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetInvalidationCallBack(_:)
func CFMessagePortGetInvalidationCallBack(ms unsafe.Pointer) unsafe.Pointer {
	return _CFMessagePortGetInvalidationCallBack(ms)
	}


// Returns the name with which a CFMessagePort object is registered. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetName(_:)
func CFMessagePortGetName(ms unsafe.Pointer) unsafe.Pointer {
	return _CFMessagePortGetName(ms)
	}


// Returns the type identifier for the CFMessagePort opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetTypeID()
func CFMessagePortGetTypeID() unsafe.Pointer {
	return _CFMessagePortGetTypeID()
	}


// Invalidates a CFMessagePort object, stopping it from receiving or sending any more messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortInvalidate(_:)
func CFMessagePortInvalidate(ms unsafe.Pointer) {
	_CFMessagePortInvalidate(ms)
	}


// Returns a Boolean value that indicates whether a CFMessagePort object represents a remote port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortIsRemote(_:)
func CFMessagePortIsRemote(ms unsafe.Pointer) unsafe.Pointer {
	return _CFMessagePortIsRemote(ms)
	}


// Returns a Boolean value that indicates whether a CFMessagePort object is valid and able to send or receive messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortIsValid(_:)
func CFMessagePortIsValid(ms unsafe.Pointer) unsafe.Pointer {
	return _CFMessagePortIsValid(ms)
	}


// Sends a message to a remote CFMessagePort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSendRequest(_:_:_:_:_:_:_:)
func CFMessagePortSendRequest(remote unsafe.Pointer, msgid unsafe.Pointer, data unsafe.Pointer, sendTimeout unsafe.Pointer, rcvTimeout unsafe.Pointer, replyMode unsafe.Pointer, returnData unsafe.Pointer) unsafe.Pointer {
	return _CFMessagePortSendRequest(remote, msgid, data, sendTimeout, rcvTimeout, replyMode, returnData)
	}


// Schedules callbacks for the specified message port on the specified dispatch queue. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetDispatchQueue(_:_:)
func CFMessagePortSetDispatchQueue(ms unsafe.Pointer, queue unsafe.Pointer) {
	_CFMessagePortSetDispatchQueue(ms, queue)
	}


// Sets the callback function invoked when a CFMessagePort object is invalidated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetInvalidationCallBack(_:_:)
func CFMessagePortSetInvalidationCallBack(ms unsafe.Pointer, callout unsafe.Pointer) {
	_CFMessagePortSetInvalidationCallBack(ms, callout)
	}


// Sets the name of a local CFMessagePort object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetName(_:_:)
func CFMessagePortSetName(ms unsafe.Pointer, newName unsafe.Pointer) unsafe.Pointer {
	return _CFMessagePortSetName(ms, newName)
	}


// Registers an observer to receive notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterAddObserver(_:_:_:_:_:_:)
func CFNotificationCenterAddObserver(center unsafe.Pointer, observer unsafe.Pointer, callBack unsafe.Pointer, name unsafe.Pointer, object unsafe.Pointer, suspensionBehavior unsafe.Pointer) {
	_CFNotificationCenterAddObserver(center, observer, callBack, name, object, suspensionBehavior)
	}


// Returns the application’s Darwin notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetDarwinNotifyCenter()
func CFNotificationCenterGetDarwinNotifyCenter() unsafe.Pointer {
	return _CFNotificationCenterGetDarwinNotifyCenter()
	}


// Returns the application’s distributed notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetDistributedCenter()
func CFNotificationCenterGetDistributedCenter() unsafe.Pointer {
	return _CFNotificationCenterGetDistributedCenter()
	}


// Returns the application’s local notification center. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetLocalCenter()
func CFNotificationCenterGetLocalCenter() unsafe.Pointer {
	return _CFNotificationCenterGetLocalCenter()
	}


// Returns the type identifier for the CFNotificationCenter opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetTypeID()
func CFNotificationCenterGetTypeID() unsafe.Pointer {
	return _CFNotificationCenterGetTypeID()
	}


// Posts a notification for an object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterPostNotification(_:_:_:_:_:)
func CFNotificationCenterPostNotification(center unsafe.Pointer, name unsafe.Pointer, object unsafe.Pointer, userInfo unsafe.Pointer, deliverImmediately unsafe.Pointer) {
	_CFNotificationCenterPostNotification(center, name, object, userInfo, deliverImmediately)
	}


// Posts a notification for an object using specified options. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterPostNotificationWithOptions(_:_:_:_:_:)
func CFNotificationCenterPostNotificationWithOptions(center unsafe.Pointer, name unsafe.Pointer, object unsafe.Pointer, userInfo unsafe.Pointer, options unsafe.Pointer) {
	_CFNotificationCenterPostNotificationWithOptions(center, name, object, userInfo, options)
	}


// Stops an observer from receiving any notifications from any object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterRemoveEveryObserver(_:_:)
func CFNotificationCenterRemoveEveryObserver(center unsafe.Pointer, observer unsafe.Pointer) {
	_CFNotificationCenterRemoveEveryObserver(center, observer)
	}


// Stops an observer from receiving certain notifications. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterRemoveObserver(_:_:_:_:)
func CFNotificationCenterRemoveObserver(center unsafe.Pointer, observer unsafe.Pointer, name unsafe.Pointer, object unsafe.Pointer) {
	_CFNotificationCenterRemoveObserver(center, observer, name, object)
	}


// Returns the type identifier for the CFNull opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNullGetTypeID()
func CFNullGetTypeID() unsafe.Pointer {
	return _CFNullGetTypeID()
	}


// Compares two CFNumber objects and returns a comparison result. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberCompare(_:_:_:)
func CFNumberCompare(number unsafe.Pointer, otherNumber unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFNumberCompare(number, otherNumber, context)
	}


// Creates a CFNumber object using a specified value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberCreate(_:_:_:)
func CFNumberCreate(allocator unsafe.Pointer, theType unsafe.Pointer, valuePtr unsafe.Pointer) unsafe.Pointer {
	return _CFNumberCreate(allocator, theType, valuePtr)
	}


// Returns a copy of a number formatter’s value for a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCopyProperty(_:_:)
func CFNumberFormatterCopyProperty(formatter unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterCopyProperty(formatter, key)
	}


// Creates a new CFNumberFormatter object, localized to the given locale, which will format numbers to the given style. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreate(_:_:_:)
func CFNumberFormatterCreate(allocator unsafe.Pointer, locale unsafe.Pointer, style unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterCreate(allocator, locale, style)
	}


// Returns a number object representing a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateNumberFromString(_:_:_:_:_:)
func CFNumberFormatterCreateNumberFromString(allocator unsafe.Pointer, formatter unsafe.Pointer, string_ unsafe.Pointer, rangep unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterCreateNumberFromString(allocator, formatter, string_, rangep, options)
	}


// Returns a string representation of the given number using the specified number formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateStringWithNumber(_:_:_:)
func CFNumberFormatterCreateStringWithNumber(allocator unsafe.Pointer, formatter unsafe.Pointer, number unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterCreateStringWithNumber(allocator, formatter, number)
	}


// Returns a string representation of the given number or value using the specified number formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateStringWithValue(_:_:_:_:)
func CFNumberFormatterCreateStringWithValue(allocator unsafe.Pointer, formatter unsafe.Pointer, numberType unsafe.Pointer, valuePtr unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterCreateStringWithValue(allocator, formatter, numberType, valuePtr)
	}


// Returns the number of fraction digits that should be displayed, and the rounding increment, for a given currency. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetDecimalInfoForCurrencyCode(_:_:_:)
func CFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode unsafe.Pointer, defaultFractionDigits unsafe.Pointer, roundingIncrement unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode, defaultFractionDigits, roundingIncrement)
	}


// Returns a format string for the given number formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetFormat(_:)
func CFNumberFormatterGetFormat(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterGetFormat(formatter)
	}


// Returns the locale object used to create the given number formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetLocale(_:)
func CFNumberFormatterGetLocale(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterGetLocale(formatter)
	}


// Returns the number style used to create the given number formatter object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetStyle(_:)
func CFNumberFormatterGetStyle(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterGetStyle(formatter)
	}


// Returns the type identifier for the opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetTypeID()
func CFNumberFormatterGetTypeID() unsafe.Pointer {
	return _CFNumberFormatterGetTypeID()
	}


// Returns a number or value representing a given string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetValueFromString(_:_:_:_:_:)
func CFNumberFormatterGetValueFromString(formatter unsafe.Pointer, string_ unsafe.Pointer, rangep unsafe.Pointer, numberType unsafe.Pointer, valuePtr unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterGetValueFromString(formatter, string_, rangep, numberType, valuePtr)
	}


// Sets the format string of a number formatter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterSetFormat(_:_:)
func CFNumberFormatterSetFormat(formatter unsafe.Pointer, formatString unsafe.Pointer) {
	_CFNumberFormatterSetFormat(formatter, formatString)
	}


// Sets a number formatter property using a key-value pair. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterSetProperty(_:_:_:)
func CFNumberFormatterSetProperty(formatter unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_CFNumberFormatterSetProperty(formatter, key, value)
	}


// Returns the number of bytes used by a CFNumber object to store its value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetByteSize(_:)
func CFNumberGetByteSize(number unsafe.Pointer) unsafe.Pointer {
	return _CFNumberGetByteSize(number)
	}


// Returns the type used by a CFNumber object to store its value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetType(_:)
func CFNumberGetType(number unsafe.Pointer) unsafe.Pointer {
	return _CFNumberGetType(number)
	}


// Returns the type identifier for the CFNumber opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetTypeID()
func CFNumberGetTypeID() unsafe.Pointer {
	return _CFNumberGetTypeID()
	}


// Obtains the value of a CFNumber object cast to a specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetValue(_:_:_:)
func CFNumberGetValue(number unsafe.Pointer, theType unsafe.Pointer, valuePtr unsafe.Pointer) unsafe.Pointer {
	return _CFNumberGetValue(number, theType, valuePtr)
	}


// Determines whether a CFNumber object contains a value stored as one of the defined floating point types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberIsFloatType(_:)
func CFNumberIsFloatType(number unsafe.Pointer) unsafe.Pointer {
	return _CFNumberIsFloatType(number)
	}


// Registers a new instance of a type with . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInAddInstanceForFactory(_:)
func CFPlugInAddInstanceForFactory(factoryID unsafe.Pointer) {
	_CFPlugInAddInstanceForFactory(factoryID)
	}


// Creates a CFPlugIn given its URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInCreate(_:_:)
func CFPlugInCreate(allocator unsafe.Pointer, plugInURL unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInCreate(allocator, plugInURL)
	}


// Searches all registered plug-ins for factory functions capable of creating an instance of the given type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInFindFactoriesForPlugInType(_:)
func CFPlugInFindFactoriesForPlugInType(typeUUID unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInFindFactoriesForPlugInType(typeUUID)
	}


// Searches the given plug-in for factory functions capable of creating an instance of the given type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInFindFactoriesForPlugInTypeInPlugIn(_:_:)
func CFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID unsafe.Pointer, plugIn unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID, plugIn)
	}


// Returns a plug-in’s bundle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInGetBundle(_:)
func CFPlugInGetBundle(plugIn unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInGetBundle(plugIn)
	}


// Returns the type identifier for the opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInGetTypeID()
func CFPlugInGetTypeID() unsafe.Pointer {
	return _CFPlugInGetTypeID()
	}


// Creates a instance of a given type using a given factory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceCreate(_:_:_:)
func CFPlugInInstanceCreate(allocator unsafe.Pointer, factoryUUID unsafe.Pointer, typeUUID unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInInstanceCreate(allocator, factoryUUID, typeUUID)
	}


// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceCreateWithInstanceDataSize(_:_:_:_:_:)
func CFPlugInInstanceCreateWithInstanceDataSize(allocator unsafe.Pointer, instanceDataSize unsafe.Pointer, deallocateInstanceFunction unsafe.Pointer, factoryName unsafe.Pointer, getInterfaceFunction unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInInstanceCreateWithInstanceDataSize(allocator, instanceDataSize, deallocateInstanceFunction, factoryName, getInterfaceFunction)
	}


// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetFactoryName(_:)
func CFPlugInInstanceGetFactoryName(instance unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInInstanceGetFactoryName(instance)
	}


// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetInstanceData(_:)
func CFPlugInInstanceGetInstanceData(instance unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInInstanceGetInstanceData(instance)
	}


// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetInterfaceFunctionTable(_:_:_:)
func CFPlugInInstanceGetInterfaceFunctionTable(instance unsafe.Pointer, interfaceName unsafe.Pointer, ftbl unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInInstanceGetInterfaceFunctionTable(instance, interfaceName, ftbl)
	}


// Not recommended. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetTypeID()
func CFPlugInInstanceGetTypeID() unsafe.Pointer {
	return _CFPlugInInstanceGetTypeID()
	}


// Determines whether or not a plug-in is loaded on demand. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInIsLoadOnDemand(_:)
func CFPlugInIsLoadOnDemand(plugIn unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInIsLoadOnDemand(plugIn)
	}


// Registers a factory function and its UUID with a object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterFactoryFunction(_:_:)
func CFPlugInRegisterFactoryFunction(factoryUUID unsafe.Pointer, func_ unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInRegisterFactoryFunction(factoryUUID, func_)
	}


// Registers a factory function with a object using the function’s name instead of its UUID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterFactoryFunctionByName(_:_:_:)
func CFPlugInRegisterFactoryFunctionByName(factoryUUID unsafe.Pointer, plugIn unsafe.Pointer, functionName unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInRegisterFactoryFunctionByName(factoryUUID, plugIn, functionName)
	}


// Registers a type and its corresponding factory function with a object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterPlugInType(_:_:)
func CFPlugInRegisterPlugInType(factoryUUID unsafe.Pointer, typeUUID unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInRegisterPlugInType(factoryUUID, typeUUID)
	}


// Unregisters an instance of a type with . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRemoveInstanceForFactory(_:)
func CFPlugInRemoveInstanceForFactory(factoryID unsafe.Pointer) {
	_CFPlugInRemoveInstanceForFactory(factoryID)
	}


// Enables or disables load on demand for plug-ins that do dynamic registration (only when a client requests an instance of a supported type). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInSetLoadOnDemand(_:_:)
func CFPlugInSetLoadOnDemand(plugIn unsafe.Pointer, flag unsafe.Pointer) {
	_CFPlugInSetLoadOnDemand(plugIn, flag)
	}


// Removes the given function from a plug-in’s list of registered factory functions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInUnregisterFactory(_:)
func CFPlugInUnregisterFactory(factoryUUID unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInUnregisterFactory(factoryUUID)
	}


// Removes the given type from a plug-in’s list of registered types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInUnregisterPlugInType(_:_:)
func CFPlugInUnregisterPlugInType(factoryUUID unsafe.Pointer, typeUUID unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInUnregisterPlugInType(factoryUUID, typeUUID)
	}


// Adds suite preferences to an application’s preference search chain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAddSuitePreferencesToApp(_:_:)
func CFPreferencesAddSuitePreferencesToApp(applicationID unsafe.Pointer, suiteID unsafe.Pointer) {
	_CFPreferencesAddSuitePreferencesToApp(applicationID, suiteID)
	}


// Writes to permanent storage all pending changes to the preference data for the application, and reads the latest preference data from permanent storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAppSynchronize(_:)
func CFPreferencesAppSynchronize(applicationID unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesAppSynchronize(applicationID)
	}


// Determines whether or not a given key has been imposed on the user. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAppValueIsForced(_:_:)
func CFPreferencesAppValueIsForced(key unsafe.Pointer, applicationID unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesAppValueIsForced(key, applicationID)
	}


// Obtains a preference value for the specified key and application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyAppValue(_:_:)
func CFPreferencesCopyAppValue(key unsafe.Pointer, applicationID unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesCopyAppValue(key, applicationID)
	}


// Constructs and returns the list of all applications that have preferences in the scope of the specified user and host. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyApplicationList(_:_:)
func CFPreferencesCopyApplicationList(userName unsafe.Pointer, hostName unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesCopyApplicationList(userName, hostName)
	}


// Constructs and returns the list of all keys set in the specified domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyKeyList(_:_:_:)
func CFPreferencesCopyKeyList(applicationID unsafe.Pointer, userName unsafe.Pointer, hostName unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesCopyKeyList(applicationID, userName, hostName)
	}


// Returns a dictionary containing preference values for multiple keys. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyMultiple(_:_:_:_:)
func CFPreferencesCopyMultiple(keysToFetch unsafe.Pointer, applicationID unsafe.Pointer, userName unsafe.Pointer, hostName unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesCopyMultiple(keysToFetch, applicationID, userName, hostName)
	}


// Returns a preference value for a given domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyValue(_:_:_:_:)
func CFPreferencesCopyValue(key unsafe.Pointer, applicationID unsafe.Pointer, userName unsafe.Pointer, hostName unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesCopyValue(key, applicationID, userName, hostName)
	}


// Convenience function that directly obtains a Boolean preference value for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesGetAppBooleanValue(_:_:_:)
func CFPreferencesGetAppBooleanValue(key unsafe.Pointer, applicationID unsafe.Pointer, keyExistsAndHasValidFormat unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesGetAppBooleanValue(key, applicationID, keyExistsAndHasValidFormat)
	}


// Convenience function that directly obtains an integer preference value for the specified key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesGetAppIntegerValue(_:_:_:)
func CFPreferencesGetAppIntegerValue(key unsafe.Pointer, applicationID unsafe.Pointer, keyExistsAndHasValidFormat unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesGetAppIntegerValue(key, applicationID, keyExistsAndHasValidFormat)
	}


// Removes suite preferences from an application’s search chain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesRemoveSuitePreferencesFromApp(_:_:)
func CFPreferencesRemoveSuitePreferencesFromApp(applicationID unsafe.Pointer, suiteID unsafe.Pointer) {
	_CFPreferencesRemoveSuitePreferencesFromApp(applicationID, suiteID)
	}


// Adds, modifies, or removes a preference. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetAppValue(_:_:_:)
func CFPreferencesSetAppValue(key unsafe.Pointer, value unsafe.Pointer, applicationID unsafe.Pointer) {
	_CFPreferencesSetAppValue(key, value, applicationID)
	}


// Convenience function that allows you to set and remove multiple preference values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetMultiple(_:_:_:_:_:)
func CFPreferencesSetMultiple(keysToSet unsafe.Pointer, keysToRemove unsafe.Pointer, applicationID unsafe.Pointer, userName unsafe.Pointer, hostName unsafe.Pointer) {
	_CFPreferencesSetMultiple(keysToSet, keysToRemove, applicationID, userName, hostName)
	}


// Adds, modifies, or removes a preference value for the specified domain. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetValue(_:_:_:_:_:)
func CFPreferencesSetValue(key unsafe.Pointer, value unsafe.Pointer, applicationID unsafe.Pointer, userName unsafe.Pointer, hostName unsafe.Pointer) {
	_CFPreferencesSetValue(key, value, applicationID, userName, hostName)
	}


// For the specified domain, writes all pending changes to preference data to permanent storage, and reads latest preference data from permanent storage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSynchronize(_:_:_:)
func CFPreferencesSynchronize(applicationID unsafe.Pointer, userName unsafe.Pointer, hostName unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesSynchronize(applicationID, userName, hostName)
	}


// Returns a CFData object containing a serialized representation of a given property list in a specified format. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateData(_:_:_:_:_:)
func CFPropertyListCreateData(allocator unsafe.Pointer, propertyList unsafe.Pointer, format unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFPropertyListCreateData(allocator, propertyList, format, options, error_)
	}


// Recursively creates a copy of a given property list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateDeepCopy(_:_:_:)
func CFPropertyListCreateDeepCopy(allocator unsafe.Pointer, propertyList unsafe.Pointer, mutabilityOption unsafe.Pointer) unsafe.Pointer {
	return _CFPropertyListCreateDeepCopy(allocator, propertyList, mutabilityOption)
	}


// Creates a property list using data from a stream. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateFromStream(_:_:_:_:_:_:)
func CFPropertyListCreateFromStream(allocator unsafe.Pointer, stream unsafe.Pointer, streamLength unsafe.Pointer, mutabilityOption unsafe.Pointer, format unsafe.Pointer, errorString unsafe.Pointer) unsafe.Pointer {
	return _CFPropertyListCreateFromStream(allocator, stream, streamLength, mutabilityOption, format, errorString)
	}


// Creates a property list using the specified XML or binary property list data. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateFromXMLData(_:_:_:_:)
func CFPropertyListCreateFromXMLData(allocator unsafe.Pointer, xmlData unsafe.Pointer, mutabilityOption unsafe.Pointer, errorString unsafe.Pointer) unsafe.Pointer {
	return _CFPropertyListCreateFromXMLData(allocator, xmlData, mutabilityOption, errorString)
	}


// Creates a property list from a given CFData object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateWithData(_:_:_:_:_:)
func CFPropertyListCreateWithData(allocator unsafe.Pointer, data unsafe.Pointer, options unsafe.Pointer, format unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFPropertyListCreateWithData(allocator, data, options, format, error_)
	}


// Create and return a property list with a CFReadStream input. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateWithStream(_:_:_:_:_:_:)
func CFPropertyListCreateWithStream(allocator unsafe.Pointer, stream unsafe.Pointer, streamLength unsafe.Pointer, options unsafe.Pointer, format unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFPropertyListCreateWithStream(allocator, stream, streamLength, options, format, error_)
	}


// Creates an XML representation of the specified property list. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateXMLData(_:_:)
func CFPropertyListCreateXMLData(allocator unsafe.Pointer, propertyList unsafe.Pointer) unsafe.Pointer {
	return _CFPropertyListCreateXMLData(allocator, propertyList)
	}


// Determines if a property list is valid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListIsValid(_:_:)
func CFPropertyListIsValid(plist unsafe.Pointer, format unsafe.Pointer) unsafe.Pointer {
	return _CFPropertyListIsValid(plist, format)
	}


// Write the bytes of a serialized property list out to a stream. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListWrite(_:_:_:_:_:)
func CFPropertyListWrite(propertyList unsafe.Pointer, stream unsafe.Pointer, format unsafe.Pointer, options unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFPropertyListWrite(propertyList, stream, format, options, error_)
	}


// Writes the bytes of a property list serialization out to a stream. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListWriteToStream(_:_:_:_:)
func CFPropertyListWriteToStream(propertyList unsafe.Pointer, stream unsafe.Pointer, format unsafe.Pointer, errorString unsafe.Pointer) unsafe.Pointer {
	return _CFPropertyListWriteToStream(propertyList, stream, format, errorString)
	}


// Closes a readable stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamClose(_:)
func CFReadStreamClose(stream unsafe.Pointer) {
	_CFReadStreamClose(stream)
	}


// CFReadStreamCopyDispatchQueue is a CoreFoundation function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyDispatchQueue(_:)
func CFReadStreamCopyDispatchQueue(stream unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamCopyDispatchQueue(stream)
	}


// Returns the error associated with a stream. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyError(_:)
func CFReadStreamCopyError(stream unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamCopyError(stream)
	}


// Returns the value of a property for a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyProperty(_:_:)
func CFReadStreamCopyProperty(stream unsafe.Pointer, propertyName unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamCopyProperty(stream, propertyName)
	}


// Creates a readable stream for a block of memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCreateWithBytesNoCopy(_:_:_:_:)
func CFReadStreamCreateWithBytesNoCopy(alloc unsafe.Pointer, bytes unsafe.Pointer, length unsafe.Pointer, bytesDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamCreateWithBytesNoCopy(alloc, bytes, length, bytesDeallocator)
	}


// Creates a readable stream for a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCreateWithFile(_:_:)
func CFReadStreamCreateWithFile(alloc unsafe.Pointer, fileURL unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamCreateWithFile(alloc, fileURL)
	}


// Returns a pointer to a stream’s internal buffer of unread data, if possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetBuffer(_:_:_:)
func CFReadStreamGetBuffer(stream unsafe.Pointer, maxBytesToRead unsafe.Pointer, numBytesRead unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamGetBuffer(stream, maxBytesToRead, numBytesRead)
	}


// Returns the error status of a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetError(_:)
func CFReadStreamGetError(stream unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamGetError(stream)
	}


// Returns the current state of a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetStatus(_:)
func CFReadStreamGetStatus(stream unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamGetStatus(stream)
	}


// Returns the type identifier the opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetTypeID()
func CFReadStreamGetTypeID() unsafe.Pointer {
	return _CFReadStreamGetTypeID()
	}


// Returns a Boolean value that indicates whether a readable stream has data that can be read without blocking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamHasBytesAvailable(_:)
func CFReadStreamHasBytesAvailable(stream unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamHasBytesAvailable(stream)
	}


// Opens a stream for reading. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamOpen(_:)
func CFReadStreamOpen(stream unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamOpen(stream)
	}


// Reads data from a readable stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamRead(_:_:_:)
func CFReadStreamRead(stream unsafe.Pointer, buffer unsafe.Pointer, bufferLength unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamRead(stream, buffer, bufferLength)
	}


// Schedules a stream into a run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamScheduleWithRunLoop(_:_:_:)
func CFReadStreamScheduleWithRunLoop(stream unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_CFReadStreamScheduleWithRunLoop(stream, runLoop, runLoopMode)
	}


// Assigns a client to a stream, which receives callbacks when certain events occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetClient(_:_:_:_:)
func CFReadStreamSetClient(stream unsafe.Pointer, streamEvents unsafe.Pointer, clientCB unsafe.Pointer, clientContext unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamSetClient(stream, streamEvents, clientCB, clientContext)
	}


// CFReadStreamSetDispatchQueue is a CoreFoundation function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetDispatchQueue(_:_:)
func CFReadStreamSetDispatchQueue(stream unsafe.Pointer, q unsafe.Pointer) {
	_CFReadStreamSetDispatchQueue(stream, q)
	}


// Sets the value of a property for a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetProperty(_:_:_:)
func CFReadStreamSetProperty(stream unsafe.Pointer, propertyName unsafe.Pointer, propertyValue unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamSetProperty(stream, propertyName, propertyValue)
	}


// Removes a read stream from a given run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamUnscheduleFromRunLoop(_:_:_:)
func CFReadStreamUnscheduleFromRunLoop(stream unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_CFReadStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode)
	}


// Adds a mode to the set of run loop common modes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddCommonMode(_:_:)
func CFRunLoopAddCommonMode(rl unsafe.Pointer, mode unsafe.Pointer) {
	_CFRunLoopAddCommonMode(rl, mode)
	}


// Adds a CFRunLoopObserver object to a run loop mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddObserver(_:_:_:)
func CFRunLoopAddObserver(rl unsafe.Pointer, observer unsafe.Pointer, mode unsafe.Pointer) {
	_CFRunLoopAddObserver(rl, observer, mode)
	}


// Adds a CFRunLoopSource object to a run loop mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddSource(_:_:_:)
func CFRunLoopAddSource(rl unsafe.Pointer, source unsafe.Pointer, mode unsafe.Pointer) {
	_CFRunLoopAddSource(rl, source, mode)
	}


// Adds a CFRunLoopTimer object to a run loop mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddTimer(_:_:_:)
func CFRunLoopAddTimer(rl unsafe.Pointer, timer unsafe.Pointer, mode unsafe.Pointer) {
	_CFRunLoopAddTimer(rl, timer, mode)
	}


// Returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopObserver object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsObserver(_:_:_:)
func CFRunLoopContainsObserver(rl unsafe.Pointer, observer unsafe.Pointer, mode unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopContainsObserver(rl, observer, mode)
	}


// Returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopSource object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsSource(_:_:_:)
func CFRunLoopContainsSource(rl unsafe.Pointer, source unsafe.Pointer, mode unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopContainsSource(rl, source, mode)
	}


// Returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopTimer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsTimer(_:_:_:)
func CFRunLoopContainsTimer(rl unsafe.Pointer, timer unsafe.Pointer, mode unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopContainsTimer(rl, timer, mode)
	}


// Returns an array that contains all the defined modes for a CFRunLoop object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopCopyAllModes(_:)
func CFRunLoopCopyAllModes(rl unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopCopyAllModes(rl)
	}


// Returns the name of the mode in which a given run loop is currently running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopCopyCurrentMode(_:)
func CFRunLoopCopyCurrentMode(rl unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopCopyCurrentMode(rl)
	}


// Returns the CFRunLoop object for the current thread. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetCurrent()
func CFRunLoopGetCurrent() unsafe.Pointer {
	return _CFRunLoopGetCurrent()
	}


// Returns the main CFRunLoop object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetMain()
func CFRunLoopGetMain() unsafe.Pointer {
	return _CFRunLoopGetMain()
	}


// Returns the time at which the next timer will fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetNextTimerFireDate(_:_:)
func CFRunLoopGetNextTimerFireDate(rl unsafe.Pointer, mode unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopGetNextTimerFireDate(rl, mode)
	}


// Returns the type identifier for the CFRunLoop opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetTypeID()
func CFRunLoopGetTypeID() unsafe.Pointer {
	return _CFRunLoopGetTypeID()
	}


// Returns a Boolean value that indicates whether the run loop is waiting for an event. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopIsWaiting(_:)
func CFRunLoopIsWaiting(rl unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopIsWaiting(rl)
	}


// Creates a CFRunLoopObserver object with a function callback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverCreate(_:_:_:_:_:_:)
func CFRunLoopObserverCreate(allocator unsafe.Pointer, activities unsafe.Pointer, repeats unsafe.Pointer, order unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopObserverCreate(allocator, activities, repeats, order, callout, context)
	}


// Creates a CFRunLoopObserver object with a block-based handler. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverCreateWithHandler(_:_:_:_:_:)
func CFRunLoopObserverCreateWithHandler(allocator unsafe.Pointer, activities unsafe.Pointer, repeats unsafe.Pointer, order unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopObserverCreateWithHandler(allocator, activities, repeats, order)
	}


// Returns a Boolean value that indicates whether a CFRunLoopObserver repeats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverDoesRepeat(_:)
func CFRunLoopObserverDoesRepeat(observer unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopObserverDoesRepeat(observer)
	}


// Returns the run loop stages during which an observer runs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetActivities(_:)
func CFRunLoopObserverGetActivities(observer unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopObserverGetActivities(observer)
	}


// Returns the context information for a CFRunLoopObserver object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetContext(_:_:)
func CFRunLoopObserverGetContext(observer unsafe.Pointer, context unsafe.Pointer) {
	_CFRunLoopObserverGetContext(observer, context)
	}


// Returns the ordering parameter for a CFRunLoopObserver object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetOrder(_:)
func CFRunLoopObserverGetOrder(observer unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopObserverGetOrder(observer)
	}


// Returns the type identifier for the CFRunLoopObserver opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetTypeID()
func CFRunLoopObserverGetTypeID() unsafe.Pointer {
	return _CFRunLoopObserverGetTypeID()
	}


// Invalidates a CFRunLoopObserver object, stopping it from ever firing again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverInvalidate(_:)
func CFRunLoopObserverInvalidate(observer unsafe.Pointer) {
	_CFRunLoopObserverInvalidate(observer)
	}


// Returns a Boolean value that indicates whether a CFRunLoopObserver object is valid and able to fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverIsValid(_:)
func CFRunLoopObserverIsValid(observer unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopObserverIsValid(observer)
	}


// Enqueues a block object on a given runloop to be executed as the runloop cycles in specified modes. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopPerformBlock(_:_:_:)
func CFRunLoopPerformBlock(rl unsafe.Pointer, mode unsafe.Pointer) {
	_CFRunLoopPerformBlock(rl, mode)
	}


// Removes a CFRunLoopObserver object from a run loop mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveObserver(_:_:_:)
func CFRunLoopRemoveObserver(rl unsafe.Pointer, observer unsafe.Pointer, mode unsafe.Pointer) {
	_CFRunLoopRemoveObserver(rl, observer, mode)
	}


// Removes a CFRunLoopSource object from a run loop mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveSource(_:_:_:)
func CFRunLoopRemoveSource(rl unsafe.Pointer, source unsafe.Pointer, mode unsafe.Pointer) {
	_CFRunLoopRemoveSource(rl, source, mode)
	}


// Removes a CFRunLoopTimer object from a run loop mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveTimer(_:_:_:)
func CFRunLoopRemoveTimer(rl unsafe.Pointer, timer unsafe.Pointer, mode unsafe.Pointer) {
	_CFRunLoopRemoveTimer(rl, timer, mode)
	}


// Runs the current thread’s CFRunLoop object in its default mode indefinitely. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRun()
func CFRunLoopRun() {
	_CFRunLoopRun()
	}


// Runs the current thread’s CFRunLoop object in a particular mode. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunInMode(_:_:_:)
func CFRunLoopRunInMode(mode unsafe.Pointer, seconds unsafe.Pointer, returnAfterSourceHandled unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopRunInMode(mode, seconds, returnAfterSourceHandled)
	}


// Creates a CFRunLoopSource object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceCreate(_:_:_:)
func CFRunLoopSourceCreate(allocator unsafe.Pointer, order unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopSourceCreate(allocator, order, context)
	}


// Returns the context information for a CFRunLoopSource object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetContext(_:_:)
func CFRunLoopSourceGetContext(source unsafe.Pointer, context unsafe.Pointer) {
	_CFRunLoopSourceGetContext(source, context)
	}


// Returns the ordering parameter for a CFRunLoopSource object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetOrder(_:)
func CFRunLoopSourceGetOrder(source unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopSourceGetOrder(source)
	}


// Returns the type identifier of the CFRunLoopSource opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetTypeID()
func CFRunLoopSourceGetTypeID() unsafe.Pointer {
	return _CFRunLoopSourceGetTypeID()
	}


// Invalidates a CFRunLoopSource object, stopping it from ever firing again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceInvalidate(_:)
func CFRunLoopSourceInvalidate(source unsafe.Pointer) {
	_CFRunLoopSourceInvalidate(source)
	}


// Returns a Boolean value that indicates whether a CFRunLoopSource object is valid and able to fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceIsValid(_:)
func CFRunLoopSourceIsValid(source unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopSourceIsValid(source)
	}


// Signals a CFRunLoopSource object, marking it as ready to fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceSignal(_:)
func CFRunLoopSourceSignal(source unsafe.Pointer) {
	_CFRunLoopSourceSignal(source)
	}


// Forces a CFRunLoop object to stop running. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopStop(_:)
func CFRunLoopStop(rl unsafe.Pointer) {
	_CFRunLoopStop(rl)
	}


// Creates a new CFRunLoopTimer object with a function callback. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerCreate(_:_:_:_:_:_:_:)
func CFRunLoopTimerCreate(allocator unsafe.Pointer, fireDate unsafe.Pointer, interval unsafe.Pointer, flags unsafe.Pointer, order unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopTimerCreate(allocator, fireDate, interval, flags, order, callout, context)
	}


// Creates a new CFRunLoopTimer object with a block-based handler. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerCreateWithHandler(_:_:_:_:_:_:)
func CFRunLoopTimerCreateWithHandler(allocator unsafe.Pointer, fireDate unsafe.Pointer, interval unsafe.Pointer, flags unsafe.Pointer, order unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopTimerCreateWithHandler(allocator, fireDate, interval, flags, order)
	}


// Returns a Boolean value that indicates whether a CFRunLoopTimer object repeats. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerDoesRepeat(_:)
func CFRunLoopTimerDoesRepeat(timer unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopTimerDoesRepeat(timer)
	}


// Returns the context information for a CFRunLoopTimer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetContext(_:_:)
func CFRunLoopTimerGetContext(timer unsafe.Pointer, context unsafe.Pointer) {
	_CFRunLoopTimerGetContext(timer, context)
	}


// Returns the firing interval of a repeating CFRunLoopTimer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetInterval(_:)
func CFRunLoopTimerGetInterval(timer unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopTimerGetInterval(timer)
	}


// Returns the next firing time for a CFRunLoopTimer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetNextFireDate(_:)
func CFRunLoopTimerGetNextFireDate(timer unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopTimerGetNextFireDate(timer)
	}


// Returns the ordering parameter for a CFRunLoopTimer object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetOrder(_:)
func CFRunLoopTimerGetOrder(timer unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopTimerGetOrder(timer)
	}


// CFRunLoopTimerGetTolerance is a CoreFoundation function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetTolerance(_:)
func CFRunLoopTimerGetTolerance(timer unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopTimerGetTolerance(timer)
	}


// Returns the type identifier of the CFRunLoopTimer opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetTypeID()
func CFRunLoopTimerGetTypeID() unsafe.Pointer {
	return _CFRunLoopTimerGetTypeID()
	}


// Invalidates a CFRunLoopTimer object, stopping it from ever firing again. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerInvalidate(_:)
func CFRunLoopTimerInvalidate(timer unsafe.Pointer) {
	_CFRunLoopTimerInvalidate(timer)
	}


// Returns a Boolean value that indicates whether a CFRunLoopTimer object is valid and able to fire. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerIsValid(_:)
func CFRunLoopTimerIsValid(timer unsafe.Pointer) unsafe.Pointer {
	return _CFRunLoopTimerIsValid(timer)
	}


// Sets the next firing date for a CFRunLoopTimer object . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerSetNextFireDate(_:_:)
func CFRunLoopTimerSetNextFireDate(timer unsafe.Pointer, fireDate unsafe.Pointer) {
	_CFRunLoopTimerSetNextFireDate(timer, fireDate)
	}


// CFRunLoopTimerSetTolerance is a CoreFoundation function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerSetTolerance(_:_:)
func CFRunLoopTimerSetTolerance(timer unsafe.Pointer, tolerance unsafe.Pointer) {
	_CFRunLoopTimerSetTolerance(timer, tolerance)
	}


// Wakes a waiting CFRunLoop object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopWakeUp(_:)
func CFRunLoopWakeUp(rl unsafe.Pointer) {
	_CFRunLoopWakeUp(rl)
	}


// Adds a value to a CFMutableSet object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetAddValue(_:_:)
func CFSetAddValue(theSet unsafe.Pointer, value unsafe.Pointer) {
	_CFSetAddValue(theSet, value)
	}


// Calls a function once for each value in a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetApplyFunction(_:_:_:)
func CFSetApplyFunction(theSet unsafe.Pointer, applier unsafe.Pointer, context unsafe.Pointer) {
	_CFSetApplyFunction(theSet, applier, context)
	}


// Returns a Boolean that indicates whether a set contains a given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetContainsValue(_:_:)
func CFSetContainsValue(theSet unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFSetContainsValue(theSet, value)
	}


// Creates an immutable CFSet object containing supplied values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCreate(_:_:_:_:)
func CFSetCreate(allocator unsafe.Pointer, values unsafe.Pointer, numValues unsafe.Pointer, callBacks unsafe.Pointer) unsafe.Pointer {
	return _CFSetCreate(allocator, values, numValues, callBacks)
	}


// Creates an immutable set containing the values of an existing set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateCopy(_:_:)
func CFSetCreateCopy(allocator unsafe.Pointer, theSet unsafe.Pointer) unsafe.Pointer {
	return _CFSetCreateCopy(allocator, theSet)
	}


// Creates an empty CFMutableSet object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateMutable(_:_:_:)
func CFSetCreateMutable(allocator unsafe.Pointer, capacity unsafe.Pointer, callBacks unsafe.Pointer) unsafe.Pointer {
	return _CFSetCreateMutable(allocator, capacity, callBacks)
	}


// Creates a new mutable set with the values from another set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateMutableCopy(_:_:_:)
func CFSetCreateMutableCopy(allocator unsafe.Pointer, capacity unsafe.Pointer, theSet unsafe.Pointer) unsafe.Pointer {
	return _CFSetCreateMutableCopy(allocator, capacity, theSet)
	}


// Returns the number of values currently in a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetCount(_:)
func CFSetGetCount(theSet unsafe.Pointer) unsafe.Pointer {
	return _CFSetGetCount(theSet)
	}


// Returns the number of values in a set that match a given value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetCountOfValue(_:_:)
func CFSetGetCountOfValue(theSet unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFSetGetCountOfValue(theSet, value)
	}


// Returns the type identifier for the CFSet type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetTypeID()
func CFSetGetTypeID() unsafe.Pointer {
	return _CFSetGetTypeID()
	}


// Obtains a specified value from a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValue(_:_:)
func CFSetGetValue(theSet unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFSetGetValue(theSet, value)
	}


// Reports whether or not a value is in a set, and if it exists returns the value indirectly. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValueIfPresent(_:_:_:)
func CFSetGetValueIfPresent(theSet unsafe.Pointer, candidate unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFSetGetValueIfPresent(theSet, candidate, value)
	}


// Obtains all values in a set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValues(_:_:)
func CFSetGetValues(theSet unsafe.Pointer, values unsafe.Pointer) {
	_CFSetGetValues(theSet, values)
	}


// Removes all values from a CFMutableSet object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetRemoveAllValues(_:)
func CFSetRemoveAllValues(theSet unsafe.Pointer) {
	_CFSetRemoveAllValues(theSet)
	}


// Removes a value from a CFMutableSet object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetRemoveValue(_:_:)
func CFSetRemoveValue(theSet unsafe.Pointer, value unsafe.Pointer) {
	_CFSetRemoveValue(theSet, value)
	}


// Replaces a value in a CFMutableSet object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetReplaceValue(_:_:)
func CFSetReplaceValue(theSet unsafe.Pointer, value unsafe.Pointer) {
	_CFSetReplaceValue(theSet, value)
	}


// Sets a value in a CFMutableSet object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetSetValue(_:_:)
func CFSetSetValue(theSet unsafe.Pointer, value unsafe.Pointer) {
	_CFSetSetValue(theSet, value)
	}


// Prints a description of a Core Foundation object to stderr. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFShow(_:)
func CFShow(obj unsafe.Pointer) {
	_CFShow(obj)
	}


// Prints the attributes of a string during debugging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFShowStr(_:)
func CFShowStr(str unsafe.Pointer) {
	_CFShowStr(str)
	}


// Opens a connection to a remote socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketConnectToAddress(_:_:_:)
func CFSocketConnectToAddress(s unsafe.Pointer, address unsafe.Pointer, timeout unsafe.Pointer) unsafe.Pointer {
	return _CFSocketConnectToAddress(s, address, timeout)
	}


// Returns the local address of a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyAddress(_:)
func CFSocketCopyAddress(s unsafe.Pointer) unsafe.Pointer {
	return _CFSocketCopyAddress(s)
	}


// Returns the remote address to which a CFSocket object is connected. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyPeerAddress(_:)
func CFSocketCopyPeerAddress(s unsafe.Pointer) unsafe.Pointer {
	return _CFSocketCopyPeerAddress(s)
	}


// Returns a socket signature registered with a CFSocket name server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyRegisteredSocketSignature(_:_:_:_:_:)
func CFSocketCopyRegisteredSocketSignature(nameServerSignature unsafe.Pointer, timeout unsafe.Pointer, name unsafe.Pointer, signature unsafe.Pointer, nameServerAddress unsafe.Pointer) unsafe.Pointer {
	return _CFSocketCopyRegisteredSocketSignature(nameServerSignature, timeout, name, signature, nameServerAddress)
	}


// Returns a value registered with a CFSocket name server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyRegisteredValue(_:_:_:_:_:)
func CFSocketCopyRegisteredValue(nameServerSignature unsafe.Pointer, timeout unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer, nameServerAddress unsafe.Pointer) unsafe.Pointer {
	return _CFSocketCopyRegisteredValue(nameServerSignature, timeout, name, value, nameServerAddress)
	}


// Creates a CFSocket object of a specified protocol and type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreate(_:_:_:_:_:_:_:)
func CFSocketCreate(allocator unsafe.Pointer, protocolFamily unsafe.Pointer, socketType unsafe.Pointer, protocol_ unsafe.Pointer, callBackTypes unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFSocketCreate(allocator, protocolFamily, socketType, protocol_, callBackTypes, callout, context)
	}


// Creates a CFSocket object and opens a connection to a remote socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateConnectedToSocketSignature(_:_:_:_:_:_:)
func CFSocketCreateConnectedToSocketSignature(allocator unsafe.Pointer, signature unsafe.Pointer, callBackTypes unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer, timeout unsafe.Pointer) unsafe.Pointer {
	return _CFSocketCreateConnectedToSocketSignature(allocator, signature, callBackTypes, callout, context, timeout)
	}


// Creates a CFRunLoopSource object for a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateRunLoopSource(_:_:_:)
func CFSocketCreateRunLoopSource(allocator unsafe.Pointer, s unsafe.Pointer, order unsafe.Pointer) unsafe.Pointer {
	return _CFSocketCreateRunLoopSource(allocator, s, order)
	}


// Creates a CFSocket object for a pre-existing native socket. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateWithNative(_:_:_:_:_:)
func CFSocketCreateWithNative(allocator unsafe.Pointer, sock unsafe.Pointer, callBackTypes unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFSocketCreateWithNative(allocator, sock, callBackTypes, callout, context)
	}


// Creates a CFSocket object using information from a CFSocketSignature structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateWithSocketSignature(_:_:_:_:_:)
func CFSocketCreateWithSocketSignature(allocator unsafe.Pointer, signature unsafe.Pointer, callBackTypes unsafe.Pointer, callout unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFSocketCreateWithSocketSignature(allocator, signature, callBackTypes, callout, context)
	}


// Disables the callback function of a CFSocket object for certain types of socket activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketDisableCallBacks(_:_:)
func CFSocketDisableCallBacks(s unsafe.Pointer, callBackTypes unsafe.Pointer) {
	_CFSocketDisableCallBacks(s, callBackTypes)
	}


// Enables the callback function of a CFSocket object for certain types of socket activity. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketEnableCallBacks(_:_:)
func CFSocketEnableCallBacks(s unsafe.Pointer, callBackTypes unsafe.Pointer) {
	_CFSocketEnableCallBacks(s, callBackTypes)
	}


// Returns the context information for a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetContext(_:_:)
func CFSocketGetContext(s unsafe.Pointer, context unsafe.Pointer) {
	_CFSocketGetContext(s, context)
	}


// Returns the default port number with which to connect to a CFSocket name server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetDefaultNameRegistryPortNumber()
func CFSocketGetDefaultNameRegistryPortNumber() unsafe.Pointer {
	return _CFSocketGetDefaultNameRegistryPortNumber()
	}


// Returns the native socket associated with a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetNative(_:)
func CFSocketGetNative(s unsafe.Pointer) unsafe.Pointer {
	return _CFSocketGetNative(s)
	}


// Returns flags that control certain behaviors of a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetSocketFlags(_:)
func CFSocketGetSocketFlags(s unsafe.Pointer) unsafe.Pointer {
	return _CFSocketGetSocketFlags(s)
	}


// Returns the type identifier for the CFSocket opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetTypeID()
func CFSocketGetTypeID() unsafe.Pointer {
	return _CFSocketGetTypeID()
	}


// Invalidates a CFSocket object, stopping it from sending or receiving any more messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketInvalidate(_:)
func CFSocketInvalidate(s unsafe.Pointer) {
	_CFSocketInvalidate(s)
	}


// Returns a Boolean value that indicates whether a CFSocket object is valid and able to send or receive messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketIsValid(_:)
func CFSocketIsValid(s unsafe.Pointer) unsafe.Pointer {
	return _CFSocketIsValid(s)
	}


// Registers a socket signature with a CFSocket name server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketRegisterSocketSignature(_:_:_:_:)
func CFSocketRegisterSocketSignature(nameServerSignature unsafe.Pointer, timeout unsafe.Pointer, name unsafe.Pointer, signature unsafe.Pointer) unsafe.Pointer {
	return _CFSocketRegisterSocketSignature(nameServerSignature, timeout, name, signature)
	}


// Registers a property-list value with a CFSocket name server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketRegisterValue(_:_:_:_:)
func CFSocketRegisterValue(nameServerSignature unsafe.Pointer, timeout unsafe.Pointer, name unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFSocketRegisterValue(nameServerSignature, timeout, name, value)
	}


// Sends data over a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketSendData(_:_:_:_:)
func CFSocketSendData(s unsafe.Pointer, address unsafe.Pointer, data unsafe.Pointer, timeout unsafe.Pointer) unsafe.Pointer {
	return _CFSocketSendData(s, address, data, timeout)
	}


// Binds a local address to a CFSocket object and configures it for listening. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetAddress(_:_:)
func CFSocketSetAddress(s unsafe.Pointer, address unsafe.Pointer) unsafe.Pointer {
	return _CFSocketSetAddress(s, address)
	}


// Sets the default port number with which to connect to a CFSocket name server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetDefaultNameRegistryPortNumber(_:)
func CFSocketSetDefaultNameRegistryPortNumber(port unsafe.Pointer) {
	_CFSocketSetDefaultNameRegistryPortNumber(port)
	}


// Sets flags that control certain behaviors of a CFSocket object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetSocketFlags(_:_:)
func CFSocketSetSocketFlags(s unsafe.Pointer, flags unsafe.Pointer) {
	_CFSocketSetSocketFlags(s, flags)
	}


// Unregisters a value or socket signature with a CFSocket name server. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketUnregister(_:_:_:)
func CFSocketUnregister(nameServerSignature unsafe.Pointer, timeout unsafe.Pointer, name unsafe.Pointer) unsafe.Pointer {
	return _CFSocketUnregister(nameServerSignature, timeout, name)
	}


// Creates a bound pair of read and write streams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreateBoundPair(_:_:_:_:)
func CFStreamCreateBoundPair(alloc unsafe.Pointer, readStream unsafe.Pointer, writeStream unsafe.Pointer, transferBufferSize unsafe.Pointer) {
	_CFStreamCreateBoundPair(alloc, readStream, writeStream, transferBufferSize)
	}


// Creates readable and writable streams connected to a socket. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 26.1.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithPeerSocketSignature(_:_:_:_:)
func CFStreamCreatePairWithPeerSocketSignature(alloc unsafe.Pointer, signature unsafe.Pointer, readStream unsafe.Pointer, writeStream unsafe.Pointer) {
	_CFStreamCreatePairWithPeerSocketSignature(alloc, signature, readStream, writeStream)
	}


// Creates readable and writable streams connected to a socket. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 26.1.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithSocket(_:_:_:_:)
func CFStreamCreatePairWithSocket(alloc unsafe.Pointer, sock unsafe.Pointer, readStream unsafe.Pointer, writeStream unsafe.Pointer) {
	_CFStreamCreatePairWithSocket(alloc, sock, readStream, writeStream)
	}


// Creates readable and writable streams connected to a TCP/IP port of a particular host. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 26.1.
//
// Added in macOS 10.1.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithSocketToHost(_:_:_:_:_:)
func CFStreamCreatePairWithSocketToHost(alloc unsafe.Pointer, host unsafe.Pointer, port unsafe.Pointer, readStream unsafe.Pointer, writeStream unsafe.Pointer) {
	_CFStreamCreatePairWithSocketToHost(alloc, host, port, readStream, writeStream)
	}


// Appends the characters of a string to those of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppend(_:_:)
func CFStringAppend(theString unsafe.Pointer, appendedString unsafe.Pointer) {
	_CFStringAppend(theString, appendedString)
	}


// Appends a C string to the character contents of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendCString(_:_:_:)
func CFStringAppendCString(theString unsafe.Pointer, cStr unsafe.Pointer, encoding unsafe.Pointer) {
	_CFStringAppendCString(theString, cStr, encoding)
	}


// Appends a buffer of Unicode characters to the character contents of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendCharacters(_:_:_:)
func CFStringAppendCharacters(theString unsafe.Pointer, chars unsafe.Pointer, numChars unsafe.Pointer) {
	_CFStringAppendCharacters(theString, chars, numChars)
	}


// Appends a formatted string to the character contents of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendFormat
func CFStringAppendFormat(theString unsafe.Pointer, formatOptions unsafe.Pointer, format unsafe.Pointer) {
	_CFStringAppendFormat(theString, formatOptions, format)
	}


// Appends a formatted string to the character contents of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendFormatAndArguments(_:_:_:_:)
func CFStringAppendFormatAndArguments(theString unsafe.Pointer, formatOptions unsafe.Pointer, format unsafe.Pointer, arguments unsafe.Pointer) {
	_CFStringAppendFormatAndArguments(theString, formatOptions, format, arguments)
	}


// Appends a Pascal string to the character contents of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendPascalString(_:_:_:)
func CFStringAppendPascalString(theString unsafe.Pointer, pStr unsafe.Pointer, encoding unsafe.Pointer) {
	_CFStringAppendPascalString(theString, pStr, encoding)
	}


// Changes the first character in each word of a string to uppercase (if it is a lowercase alphabetical character). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCapitalize(_:_:)
func CFStringCapitalize(theString unsafe.Pointer, locale unsafe.Pointer) {
	_CFStringCapitalize(theString, locale)
	}


// Compares one string with another string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompare(_:_:_:)
func CFStringCompare(theString1 unsafe.Pointer, theString2 unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFStringCompare(theString1, theString2, compareOptions)
	}


// Compares a range of the characters in one string with that of another string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareWithOptions(_:_:_:_:)
func CFStringCompareWithOptions(theString1 unsafe.Pointer, theString2 unsafe.Pointer, rangeToCompare unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFStringCompareWithOptions(theString1, theString2, rangeToCompare, compareOptions)
	}


// Compares a range of the characters in one string with another string using a given locale. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareWithOptionsAndLocale(_:_:_:_:_:)
func CFStringCompareWithOptionsAndLocale(theString1 unsafe.Pointer, theString2 unsafe.Pointer, rangeToCompare unsafe.Pointer, compareOptions unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _CFStringCompareWithOptionsAndLocale(theString1, theString2, rangeToCompare, compareOptions, locale)
	}


// Returns the name of the IANA registry “charset” that is the closest mapping to a specified string encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToIANACharSetName(_:)
func CFStringConvertEncodingToIANACharSetName(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertEncodingToIANACharSetName(encoding)
	}


// Returns the Cocoa encoding constant that maps most closely to a given Core Foundation encoding constant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToNSStringEncoding(_:)
func CFStringConvertEncodingToNSStringEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertEncodingToNSStringEncoding(encoding)
	}


// Returns the Windows codepage identifier that maps most closely to a given Core Foundation encoding constant. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToWindowsCodepage(_:)
func CFStringConvertEncodingToWindowsCodepage(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertEncodingToWindowsCodepage(encoding)
	}


// Returns the Core Foundation encoding constant that is the closest mapping to a given IANA registry “charset” name. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertIANACharSetNameToEncoding(_:)
func CFStringConvertIANACharSetNameToEncoding(theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertIANACharSetNameToEncoding(theString)
	}


// Returns the Core Foundation encoding constant that is the closest mapping to a given Cocoa encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertNSStringEncodingToEncoding(_:)
func CFStringConvertNSStringEncodingToEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertNSStringEncodingToEncoding(encoding)
	}


// Returns the Core Foundation encoding constant that is the closest mapping to a given Windows codepage identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertWindowsCodepageToEncoding(_:)
func CFStringConvertWindowsCodepageToEncoding(codepage unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertWindowsCodepageToEncoding(codepage)
	}


// Creates an array of CFString objects from a single CFString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateArrayBySeparatingStrings(_:_:_:)
func CFStringCreateArrayBySeparatingStrings(alloc unsafe.Pointer, theString unsafe.Pointer, separatorString unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateArrayBySeparatingStrings(alloc, theString, separatorString)
	}


// Searches a string for multiple occurrences of a substring and creates an array of ranges identifying the locations of these substrings within the target string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateArrayWithFindResults(_:_:_:_:_:)
func CFStringCreateArrayWithFindResults(alloc unsafe.Pointer, theString unsafe.Pointer, stringToFind unsafe.Pointer, rangeToSearch unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateArrayWithFindResults(alloc, theString, stringToFind, rangeToSearch, compareOptions)
	}


// Creates a single string from the individual CFString objects that comprise the elements of an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateByCombiningStrings(_:_:_:)
func CFStringCreateByCombiningStrings(alloc unsafe.Pointer, theArray unsafe.Pointer, separatorString unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateByCombiningStrings(alloc, theArray, separatorString)
	}


// Creates an immutable copy of a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateCopy(_:_:)
func CFStringCreateCopy(alloc unsafe.Pointer, theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateCopy(alloc, theString)
	}


// Creates an “external representation” of a CFString object, that is, a CFData object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateExternalRepresentation(_:_:_:_:)
func CFStringCreateExternalRepresentation(alloc unsafe.Pointer, theString unsafe.Pointer, encoding unsafe.Pointer, lossByte unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateExternalRepresentation(alloc, theString, encoding, lossByte)
	}


// Creates a string from its “external representation.” [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateFromExternalRepresentation(_:_:_:)
func CFStringCreateFromExternalRepresentation(alloc unsafe.Pointer, data unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateFromExternalRepresentation(alloc, data, encoding)
	}


// Creates an empty CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutable(_:_:)
func CFStringCreateMutable(alloc unsafe.Pointer, maxLength unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateMutable(alloc, maxLength)
	}


// Creates a mutable copy of a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutableCopy(_:_:_:)
func CFStringCreateMutableCopy(alloc unsafe.Pointer, maxLength unsafe.Pointer, theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateMutableCopy(alloc, maxLength, theString)
	}


// Creates a CFMutableString object whose Unicode character buffer is controlled externally. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutableWithExternalCharactersNoCopy(_:_:_:_:_:)
func CFStringCreateMutableWithExternalCharactersNoCopy(alloc unsafe.Pointer, chars unsafe.Pointer, numChars unsafe.Pointer, capacity unsafe.Pointer, externalCharactersAllocator unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateMutableWithExternalCharactersNoCopy(alloc, chars, numChars, capacity, externalCharactersAllocator)
	}


// CFStringCreateStringWithValidatedFormat is a CoreFoundation function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateStringWithValidatedFormat
func CFStringCreateStringWithValidatedFormat(alloc unsafe.Pointer, formatOptions unsafe.Pointer, validFormatSpecifiers unsafe.Pointer, format unsafe.Pointer, errorPtr unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateStringWithValidatedFormat(alloc, formatOptions, validFormatSpecifiers, format, errorPtr)
	}


// CFStringCreateStringWithValidatedFormatAndArguments is a CoreFoundation function. [Full Topic]
//
// Added in macOS 13.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateStringWithValidatedFormatAndArguments
func CFStringCreateStringWithValidatedFormatAndArguments(alloc unsafe.Pointer, formatOptions unsafe.Pointer, validFormatSpecifiers unsafe.Pointer, format unsafe.Pointer, arguments unsafe.Pointer, errorPtr unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateStringWithValidatedFormatAndArguments(alloc, formatOptions, validFormatSpecifiers, format, arguments, errorPtr)
	}


// Creates a string from a buffer containing characters in a specified encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithBytes(_:_:_:_:_:)
func CFStringCreateWithBytes(alloc unsafe.Pointer, bytes unsafe.Pointer, numBytes unsafe.Pointer, encoding unsafe.Pointer, isExternalRepresentation unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithBytes(alloc, bytes, numBytes, encoding, isExternalRepresentation)
	}


// Creates a string from a buffer, containing characters in a specified encoding, that might serve as the backing store for the new string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithBytesNoCopy(_:_:_:_:_:_:)
func CFStringCreateWithBytesNoCopy(alloc unsafe.Pointer, bytes unsafe.Pointer, numBytes unsafe.Pointer, encoding unsafe.Pointer, isExternalRepresentation unsafe.Pointer, contentsDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithBytesNoCopy(alloc, bytes, numBytes, encoding, isExternalRepresentation, contentsDeallocator)
	}


// Creates an immutable string from a C string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCString(_:_:_:)
func CFStringCreateWithCString(alloc unsafe.Pointer, cStr unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithCString(alloc, cStr, encoding)
	}


// Creates a CFString object from an external C string buffer that might serve as the backing store for the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCStringNoCopy(_:_:_:_:)
func CFStringCreateWithCStringNoCopy(alloc unsafe.Pointer, cStr unsafe.Pointer, encoding unsafe.Pointer, contentsDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithCStringNoCopy(alloc, cStr, encoding, contentsDeallocator)
	}


// Creates a string from a buffer of Unicode characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCharacters(_:_:_:)
func CFStringCreateWithCharacters(alloc unsafe.Pointer, chars unsafe.Pointer, numChars unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithCharacters(alloc, chars, numChars)
	}


// Creates a string from a buffer of Unicode characters that might serve as the backing store for the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCharactersNoCopy(_:_:_:_:)
func CFStringCreateWithCharactersNoCopy(alloc unsafe.Pointer, chars unsafe.Pointer, numChars unsafe.Pointer, contentsDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithCharactersNoCopy(alloc, chars, numChars, contentsDeallocator)
	}


// Creates a CFString from a zero-terminated POSIX file system representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFileSystemRepresentation(_:_:)
func CFStringCreateWithFileSystemRepresentation(alloc unsafe.Pointer, buffer unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithFileSystemRepresentation(alloc, buffer)
	}


// Creates an immutable string from a formatted string and a variable number of arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFormat
func CFStringCreateWithFormat(alloc unsafe.Pointer, formatOptions unsafe.Pointer, format unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithFormat(alloc, formatOptions, format)
	}


// Creates an immutable string from a formatted string and a variable number of arguments (specified in a parameter of type ). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFormatAndArguments(_:_:_:_:)
func CFStringCreateWithFormatAndArguments(alloc unsafe.Pointer, formatOptions unsafe.Pointer, format unsafe.Pointer, arguments unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithFormatAndArguments(alloc, formatOptions, format, arguments)
	}


// Creates an immutable CFString object from a Pascal string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithPascalString(_:_:_:)
func CFStringCreateWithPascalString(alloc unsafe.Pointer, pStr unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithPascalString(alloc, pStr, encoding)
	}


// Creates a CFString object from an external Pascal string buffer that might serve as the backing store for the object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithPascalStringNoCopy(_:_:_:_:)
func CFStringCreateWithPascalStringNoCopy(alloc unsafe.Pointer, pStr unsafe.Pointer, encoding unsafe.Pointer, contentsDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithPascalStringNoCopy(alloc, pStr, encoding, contentsDeallocator)
	}


// Creates an immutable string from a segment (substring) of an existing string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithSubstring(_:_:_:)
func CFStringCreateWithSubstring(alloc unsafe.Pointer, str unsafe.Pointer, range_ unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithSubstring(alloc, str, range_)
	}


// Deletes a range of characters in a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringDelete(_:_:)
func CFStringDelete(theString unsafe.Pointer, range_ unsafe.Pointer) {
	_CFStringDelete(theString, range_)
	}


// Searches for a substring within a string and, if it is found, yields the range of the substring within the object’s characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFind(_:_:_:)
func CFStringFind(theString unsafe.Pointer, stringToFind unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFStringFind(theString, stringToFind, compareOptions)
	}


// Replaces all occurrences of a substring within a given range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindAndReplace(_:_:_:_:_:)
func CFStringFindAndReplace(theString unsafe.Pointer, stringToFind unsafe.Pointer, replacementString unsafe.Pointer, rangeToSearch unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFStringFindAndReplace(theString, stringToFind, replacementString, rangeToSearch, compareOptions)
	}


// Query the range of the first character contained in the specified character set. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindCharacterFromSet(_:_:_:_:_:)
func CFStringFindCharacterFromSet(theString unsafe.Pointer, theSet unsafe.Pointer, rangeToSearch unsafe.Pointer, searchOptions unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _CFStringFindCharacterFromSet(theString, theSet, rangeToSearch, searchOptions, result)
	}


// Searches for a substring within a range of the characters represented by a string and, if the substring is found, returns its range within the object’s characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindWithOptions(_:_:_:_:_:)
func CFStringFindWithOptions(theString unsafe.Pointer, stringToFind unsafe.Pointer, rangeToSearch unsafe.Pointer, searchOptions unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _CFStringFindWithOptions(theString, stringToFind, rangeToSearch, searchOptions, result)
	}


// Returns a Boolean value that indicates whether a given string was found in a given source string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindWithOptionsAndLocale(_:_:_:_:_:_:)
func CFStringFindWithOptionsAndLocale(theString unsafe.Pointer, stringToFind unsafe.Pointer, rangeToSearch unsafe.Pointer, searchOptions unsafe.Pointer, locale unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _CFStringFindWithOptionsAndLocale(theString, stringToFind, rangeToSearch, searchOptions, locale, result)
	}


// Folds a given string into the form specified by optional flags. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFold(_:_:_:)
func CFStringFold(theString unsafe.Pointer, theFlags unsafe.Pointer, theLocale unsafe.Pointer) {
	_CFStringFold(theString, theFlags, theLocale)
	}


// Fetches a range of the characters from a string into a byte buffer after converting the characters to a specified encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetBytes(_:_:_:_:_:_:_:_:)
func CFStringGetBytes(theString unsafe.Pointer, range_ unsafe.Pointer, encoding unsafe.Pointer, lossByte unsafe.Pointer, isExternalRepresentation unsafe.Pointer, buffer unsafe.Pointer, maxBufLen unsafe.Pointer, usedBufLen unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetBytes(theString, range_, encoding, lossByte, isExternalRepresentation, buffer, maxBufLen, usedBufLen)
	}


// Copies the character contents of a string to a local C string buffer after converting the characters to a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCString(_:_:_:_:)
func CFStringGetCString(theString unsafe.Pointer, buffer unsafe.Pointer, bufferSize unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetCString(theString, buffer, bufferSize, encoding)
	}


// Quickly obtains a pointer to a C-string buffer containing the characters of a string in a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCStringPtr(_:_:)
func CFStringGetCStringPtr(theString unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetCStringPtr(theString, encoding)
	}


// Returns the Unicode character at a specified location in a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharacterAtIndex(_:_:)
func CFStringGetCharacterAtIndex(theString unsafe.Pointer, idx unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetCharacterAtIndex(theString, idx)
	}


// Copies a range of the Unicode characters from a string to a user-provided buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharacters(_:_:_:)
func CFStringGetCharacters(theString unsafe.Pointer, range_ unsafe.Pointer, buffer unsafe.Pointer) {
	_CFStringGetCharacters(theString, range_, buffer)
	}


// Quickly obtains a pointer to the contents of a string as a buffer of Unicode characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharactersPtr(_:)
func CFStringGetCharactersPtr(theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetCharactersPtr(theString)
	}


// Returns the primary value represented by a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetDoubleValue(_:)
func CFStringGetDoubleValue(str unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetDoubleValue(str)
	}


// Returns for a CFString object the character encoding that requires the least conversion time. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetFastestEncoding(_:)
func CFStringGetFastestEncoding(theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetFastestEncoding(theString)
	}


// Extracts the contents of a string as a -terminated 8-bit string appropriate for passing to POSIX APIs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetFileSystemRepresentation(_:_:_:)
func CFStringGetFileSystemRepresentation(string_ unsafe.Pointer, buffer unsafe.Pointer, maxBufLen unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetFileSystemRepresentation(string_, buffer, maxBufLen)
	}


// Retrieve the first potential hyphenation location found before the specified location. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetHyphenationLocationBeforeIndex(_:_:_:_:_:_:)
func CFStringGetHyphenationLocationBeforeIndex(string_ unsafe.Pointer, location unsafe.Pointer, limitRange unsafe.Pointer, options unsafe.Pointer, locale unsafe.Pointer, character unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetHyphenationLocationBeforeIndex(string_, location, limitRange, options, locale, character)
	}


// Returns the integer value represented by a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetIntValue(_:)
func CFStringGetIntValue(str unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetIntValue(str)
	}


// Returns the number (in terms of UTF-16 code pairs) of Unicode characters in a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetLength(_:)
func CFStringGetLength(theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetLength(theString)
	}


// Given a range of characters in a string, obtains the line bounds—that is, the indexes of the first character and the final characters of the lines containing the range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetLineBounds(_:_:_:_:_:)
func CFStringGetLineBounds(theString unsafe.Pointer, range_ unsafe.Pointer, lineBeginIndex unsafe.Pointer, lineEndIndex unsafe.Pointer, contentsEndIndex unsafe.Pointer) {
	_CFStringGetLineBounds(theString, range_, lineBeginIndex, lineEndIndex, contentsEndIndex)
	}


// Returns a pointer to a list of string encodings supported by the current system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetListOfAvailableEncodings()
func CFStringGetListOfAvailableEncodings() unsafe.Pointer {
	return _CFStringGetListOfAvailableEncodings()
	}


// Returns the maximum number of bytes a string of a specified length (in Unicode characters) will take up if encoded in a specified encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMaximumSizeForEncoding(_:_:)
func CFStringGetMaximumSizeForEncoding(length unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetMaximumSizeForEncoding(length, encoding)
	}


// Determines the upper bound on the number of bytes required to hold the file system representation of the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMaximumSizeOfFileSystemRepresentation(_:)
func CFStringGetMaximumSizeOfFileSystemRepresentation(string_ unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetMaximumSizeOfFileSystemRepresentation(string_)
	}


// Returns the most compatible Mac OS script value for the given input encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMostCompatibleMacStringEncoding(_:)
func CFStringGetMostCompatibleMacStringEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetMostCompatibleMacStringEncoding(encoding)
	}


// Returns the canonical name of a specified string encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetNameOfEncoding(_:)
func CFStringGetNameOfEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetNameOfEncoding(encoding)
	}


// Given a range of characters in a string, obtains the paragraph bounds—that is, the indexes of the first character and the final characters of the paragraph(s) containing the range. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetParagraphBounds(_:_:_:_:_:)
func CFStringGetParagraphBounds(string_ unsafe.Pointer, range_ unsafe.Pointer, parBeginIndex unsafe.Pointer, parEndIndex unsafe.Pointer, contentsEndIndex unsafe.Pointer) {
	_CFStringGetParagraphBounds(string_, range_, parBeginIndex, parEndIndex, contentsEndIndex)
	}


// Copies the character contents of a CFString object to a local Pascal string buffer after converting the characters to a requested encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetPascalString(_:_:_:_:)
func CFStringGetPascalString(theString unsafe.Pointer, buffer unsafe.Pointer, bufferSize unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetPascalString(theString, buffer, bufferSize, encoding)
	}


// Quickly obtains a pointer to a Pascal buffer containing the characters of a string in a given encoding. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetPascalStringPtr(_:_:)
func CFStringGetPascalStringPtr(theString unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetPascalStringPtr(theString, encoding)
	}


// Returns the range of the composed character sequence at a specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetRangeOfComposedCharactersAtIndex(_:_:)
func CFStringGetRangeOfComposedCharactersAtIndex(theString unsafe.Pointer, theIndex unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetRangeOfComposedCharactersAtIndex(theString, theIndex)
	}


// Returns the smallest encoding on the current system for the character contents of a string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetSmallestEncoding(_:)
func CFStringGetSmallestEncoding(theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetSmallestEncoding(theString)
	}


// Returns the default encoding used by the operating system when it creates strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetSystemEncoding()
func CFStringGetSystemEncoding() unsafe.Pointer {
	return _CFStringGetSystemEncoding()
	}


// Returns the type identifier for the CFString opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetTypeID()
func CFStringGetTypeID() unsafe.Pointer {
	return _CFStringGetTypeID()
	}


// Determines if the character data of a string begin with a specified sequence of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringHasPrefix(_:_:)
func CFStringHasPrefix(theString unsafe.Pointer, prefix unsafe.Pointer) unsafe.Pointer {
	return _CFStringHasPrefix(theString, prefix)
	}


// Determines if a string ends with a specified sequence of characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringHasSuffix(_:_:)
func CFStringHasSuffix(theString unsafe.Pointer, suffix unsafe.Pointer) unsafe.Pointer {
	return _CFStringHasSuffix(theString, suffix)
	}


// Inserts a string at a specified location in the character buffer of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringInsert(_:_:_:)
func CFStringInsert(str unsafe.Pointer, idx unsafe.Pointer, insertedStr unsafe.Pointer) {
	_CFStringInsert(str, idx, insertedStr)
	}


// Determines whether a given Core Foundation string encoding is available on the current system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringIsEncodingAvailable(_:)
func CFStringIsEncodingAvailable(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringIsEncodingAvailable(encoding)
	}


// Returns a Boolean value that indicates whether hyphenation data is available. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringIsHyphenationAvailableForLocale(_:)
func CFStringIsHyphenationAvailableForLocale(locale unsafe.Pointer) unsafe.Pointer {
	return _CFStringIsHyphenationAvailableForLocale(locale)
	}


// Changes all uppercase alphabetical characters in a CFMutableString to lowercase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringLowercase(_:_:)
func CFStringLowercase(theString unsafe.Pointer, locale unsafe.Pointer) {
	_CFStringLowercase(theString, locale)
	}


// Normalizes the string into the specified form as described in Unicode Technical Report #15. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalize(_:_:)
func CFStringNormalize(theString unsafe.Pointer, theForm unsafe.Pointer) {
	_CFStringNormalize(theString, theForm)
	}


// Enlarges a string, padding it with specified characters, or truncates the string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringPad(_:_:_:_:)
func CFStringPad(theString unsafe.Pointer, padString unsafe.Pointer, length unsafe.Pointer, indexIntoPad unsafe.Pointer) {
	_CFStringPad(theString, padString, length, indexIntoPad)
	}


// Replaces part of the character contents of a CFMutableString object with another string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringReplace(_:_:_:)
func CFStringReplace(theString unsafe.Pointer, range_ unsafe.Pointer, replacement unsafe.Pointer) {
	_CFStringReplace(theString, range_, replacement)
	}


// Replaces all characters of a CFMutableString object with other characters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringReplaceAll(_:_:)
func CFStringReplaceAll(theString unsafe.Pointer, replacement unsafe.Pointer) {
	_CFStringReplaceAll(theString, replacement)
	}


// Notifies a CFMutableString object that its external backing store of Unicode characters has changed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringSetExternalCharactersNoCopy(_:_:_:_:)
func CFStringSetExternalCharactersNoCopy(theString unsafe.Pointer, chars unsafe.Pointer, length unsafe.Pointer, capacity unsafe.Pointer) {
	_CFStringSetExternalCharactersNoCopy(theString, chars, length, capacity)
	}


// Advances the tokenizer to the next token and sets that as the current token. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerAdvanceToNextToken(_:)
func CFStringTokenizerAdvanceToNextToken(tokenizer unsafe.Pointer) unsafe.Pointer {
	return _CFStringTokenizerAdvanceToNextToken(tokenizer)
	}


// Guesses a language of a given string and returns the guess as a BCP 47 string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCopyBestStringLanguage(_:_:)
func CFStringTokenizerCopyBestStringLanguage(string_ unsafe.Pointer, range_ unsafe.Pointer) unsafe.Pointer {
	return _CFStringTokenizerCopyBestStringLanguage(string_, range_)
	}


// Returns a given attribute of the current token. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCopyCurrentTokenAttribute(_:_:)
func CFStringTokenizerCopyCurrentTokenAttribute(tokenizer unsafe.Pointer, attribute unsafe.Pointer) unsafe.Pointer {
	return _CFStringTokenizerCopyCurrentTokenAttribute(tokenizer, attribute)
	}


// Returns a tokenizer for a given string. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCreate(_:_:_:_:_:)
func CFStringTokenizerCreate(alloc unsafe.Pointer, string_ unsafe.Pointer, range_ unsafe.Pointer, options unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _CFStringTokenizerCreate(alloc, string_, range_, options, locale)
	}


// Retrieves the subtokens or derived subtokens contained in the compound token. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetCurrentSubTokens(_:_:_:_:)
func CFStringTokenizerGetCurrentSubTokens(tokenizer unsafe.Pointer, ranges unsafe.Pointer, maxRangeLength unsafe.Pointer, derivedSubTokens unsafe.Pointer) unsafe.Pointer {
	return _CFStringTokenizerGetCurrentSubTokens(tokenizer, ranges, maxRangeLength, derivedSubTokens)
	}


// Returns the range of the current token. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetCurrentTokenRange(_:)
func CFStringTokenizerGetCurrentTokenRange(tokenizer unsafe.Pointer) unsafe.Pointer {
	return _CFStringTokenizerGetCurrentTokenRange(tokenizer)
	}


// Returns the type ID for CFStringTokenizer. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetTypeID()
func CFStringTokenizerGetTypeID() unsafe.Pointer {
	return _CFStringTokenizerGetTypeID()
	}


// Finds a token that includes the character at a given index, and set it as the current token. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGoToTokenAtIndex(_:_:)
func CFStringTokenizerGoToTokenAtIndex(tokenizer unsafe.Pointer, index unsafe.Pointer) unsafe.Pointer {
	return _CFStringTokenizerGoToTokenAtIndex(tokenizer, index)
	}


// Sets the string for a tokenizer. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerSetString(_:_:_:)
func CFStringTokenizerSetString(tokenizer unsafe.Pointer, string_ unsafe.Pointer, range_ unsafe.Pointer) {
	_CFStringTokenizerSetString(tokenizer, string_, range_)
	}


// Perform in-place transliteration on a mutable string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTransform(_:_:_:_:)
func CFStringTransform(string_ unsafe.Pointer, range_ unsafe.Pointer, transform unsafe.Pointer, reverse unsafe.Pointer) unsafe.Pointer {
	return _CFStringTransform(string_, range_, transform, reverse)
	}


// Trims a specified substring from the beginning and end of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTrim(_:_:)
func CFStringTrim(theString unsafe.Pointer, trimString unsafe.Pointer) {
	_CFStringTrim(theString, trimString)
	}


// Trims whitespace from the beginning and end of a CFMutableString object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTrimWhitespace(_:)
func CFStringTrimWhitespace(theString unsafe.Pointer) {
	_CFStringTrimWhitespace(theString)
	}


// Changes all lowercase alphabetical characters in a CFMutableString object to uppercase. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringUppercase(_:_:)
func CFStringUppercase(theString unsafe.Pointer, locale unsafe.Pointer) {
	_CFStringUppercase(theString, locale)
	}


// Returns the abbreviation of a time zone at a specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyAbbreviation(_:_:)
func CFTimeZoneCopyAbbreviation(tz unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneCopyAbbreviation(tz, at)
	}


// Returns a dictionary holding the mappings of time zone abbreviations to time zone names. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyAbbreviationDictionary()
func CFTimeZoneCopyAbbreviationDictionary() unsafe.Pointer {
	return _CFTimeZoneCopyAbbreviationDictionary()
	}


// Returns the default time zone set for your application. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyDefault()
func CFTimeZoneCopyDefault() unsafe.Pointer {
	return _CFTimeZoneCopyDefault()
	}


// Returns an array of strings containing the names of all the time zones known to the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyKnownNames()
func CFTimeZoneCopyKnownNames() unsafe.Pointer {
	return _CFTimeZoneCopyKnownNames()
	}


// Returns the localized name of a given time zone. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyLocalizedName(_:_:_:)
func CFTimeZoneCopyLocalizedName(tz unsafe.Pointer, style unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneCopyLocalizedName(tz, style, locale)
	}


// Returns the time zone currently used by the system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopySystem()
func CFTimeZoneCopySystem() unsafe.Pointer {
	return _CFTimeZoneCopySystem()
	}


// Creates a time zone with a given name and data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreate(_:_:_:)
func CFTimeZoneCreate(allocator unsafe.Pointer, name unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneCreate(allocator, name, data)
	}


// Returns the time zone object identified by a given name or abbreviation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreateWithName(_:_:_:)
func CFTimeZoneCreateWithName(allocator unsafe.Pointer, name unsafe.Pointer, tryAbbrev unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneCreateWithName(allocator, name, tryAbbrev)
	}


// Returns a time zone object for the specified time interval offset from Greenwich Mean Time (GMT). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreateWithTimeIntervalFromGMT(_:_:)
func CFTimeZoneCreateWithTimeIntervalFromGMT(allocator unsafe.Pointer, ti unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneCreateWithTimeIntervalFromGMT(allocator, ti)
	}


// Returns the data that stores the information used by a time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetData(_:)
func CFTimeZoneGetData(tz unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneGetData(tz)
	}


// Returns the daylight saving time offset for a time zone at a given time. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetDaylightSavingTimeOffset(_:_:)
func CFTimeZoneGetDaylightSavingTimeOffset(tz unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneGetDaylightSavingTimeOffset(tz, at)
	}


// Returns the geopolitical region name that identifies a given time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetName(_:)
func CFTimeZoneGetName(tz unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneGetName(tz)
	}


// Returns the time in a given time zone of the next daylight saving time transition after a given time. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetNextDaylightSavingTimeTransition(_:_:)
func CFTimeZoneGetNextDaylightSavingTimeTransition(tz unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneGetNextDaylightSavingTimeTransition(tz, at)
	}


// Returns the difference in seconds between the receiver and Greenwich Mean Time (GMT) at the specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetSecondsFromGMT(_:_:)
func CFTimeZoneGetSecondsFromGMT(tz unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneGetSecondsFromGMT(tz, at)
	}


// Returns the type identifier for the CFTimeZone opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetTypeID()
func CFTimeZoneGetTypeID() unsafe.Pointer {
	return _CFTimeZoneGetTypeID()
	}


// Returns whether or not a time zone is in daylight savings time at a specified date. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneIsDaylightSavingTime(_:_:)
func CFTimeZoneIsDaylightSavingTime(tz unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneIsDaylightSavingTime(tz, at)
	}


// Clears the previously determined system time zone, if any. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneResetSystem()
func CFTimeZoneResetSystem() {
	_CFTimeZoneResetSystem()
	}


// Sets the abbreviation dictionary to a given dictionary. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneSetAbbreviationDictionary(_:)
func CFTimeZoneSetAbbreviationDictionary(dict unsafe.Pointer) {
	_CFTimeZoneSetAbbreviationDictionary(dict)
	}


// Sets the default time zone for your application the given time zone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneSetDefault(_:)
func CFTimeZoneSetDefault(tz unsafe.Pointer) {
	_CFTimeZoneSetDefault(tz)
	}


// Adds a new child to a tree as the last in its list of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeAppendChild(_:_:)
func CFTreeAppendChild(tree unsafe.Pointer, newChild unsafe.Pointer) {
	_CFTreeAppendChild(tree, newChild)
	}


// Calls a function once for each immediate child of a tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeApplyFunctionToChildren(_:_:_:)
func CFTreeApplyFunctionToChildren(tree unsafe.Pointer, applier unsafe.Pointer, context unsafe.Pointer) {
	_CFTreeApplyFunctionToChildren(tree, applier, context)
	}


// Creates a new CFTree object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeCreate(_:_:)
func CFTreeCreate(allocator unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFTreeCreate(allocator, context)
	}


// Returns the root tree of a given tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeFindRoot(_:)
func CFTreeFindRoot(tree unsafe.Pointer) unsafe.Pointer {
	return _CFTreeFindRoot(tree)
	}


// Returns the child of a tree at the specified index. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildAtIndex(_:_:)
func CFTreeGetChildAtIndex(tree unsafe.Pointer, idx unsafe.Pointer) unsafe.Pointer {
	return _CFTreeGetChildAtIndex(tree, idx)
	}


// Returns the number of children in a tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildCount(_:)
func CFTreeGetChildCount(tree unsafe.Pointer) unsafe.Pointer {
	return _CFTreeGetChildCount(tree)
	}


// Fills a buffer with children from the tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildren(_:_:)
func CFTreeGetChildren(tree unsafe.Pointer, children unsafe.Pointer) {
	_CFTreeGetChildren(tree, children)
	}


// Returns the context of the specified tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetContext(_:_:)
func CFTreeGetContext(tree unsafe.Pointer, context unsafe.Pointer) {
	_CFTreeGetContext(tree, context)
	}


// Returns the first child of a tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetFirstChild(_:)
func CFTreeGetFirstChild(tree unsafe.Pointer) unsafe.Pointer {
	return _CFTreeGetFirstChild(tree)
	}


// Returns the next sibling, adjacent to a given tree, in the parent’s children list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetNextSibling(_:)
func CFTreeGetNextSibling(tree unsafe.Pointer) unsafe.Pointer {
	return _CFTreeGetNextSibling(tree)
	}


// Returns the parent of a given tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetParent(_:)
func CFTreeGetParent(tree unsafe.Pointer) unsafe.Pointer {
	return _CFTreeGetParent(tree)
	}


// Returns the type identifier of the CFTree opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetTypeID()
func CFTreeGetTypeID() unsafe.Pointer {
	return _CFTreeGetTypeID()
	}


// Inserts a new sibling after a given tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeInsertSibling(_:_:)
func CFTreeInsertSibling(tree unsafe.Pointer, newSibling unsafe.Pointer) {
	_CFTreeInsertSibling(tree, newSibling)
	}


// Adds a new child to the specified tree as the first in its list of children. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreePrependChild(_:_:)
func CFTreePrependChild(tree unsafe.Pointer, newChild unsafe.Pointer) {
	_CFTreePrependChild(tree, newChild)
	}


// Removes a tree from its parent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeRemove(_:)
func CFTreeRemove(tree unsafe.Pointer) {
	_CFTreeRemove(tree)
	}


// Removes all the children of a tree. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeRemoveAllChildren(_:)
func CFTreeRemoveAllChildren(tree unsafe.Pointer) {
	_CFTreeRemoveAllChildren(tree)
	}


// Replaces the context of a tree by releasing the old information pointer and retaining the new one. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeSetContext(_:_:)
func CFTreeSetContext(tree unsafe.Pointer, context unsafe.Pointer) {
	_CFTreeSetContext(tree, context)
	}


// Sorts the immediate children of a tree using a specified comparator function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeSortChildren(_:_:_:)
func CFTreeSortChildren(tree unsafe.Pointer, comparator unsafe.Pointer, context unsafe.Pointer) {
	_CFTreeSortChildren(tree, comparator, context)
	}


// Determines if the given URL conforms to RFC 1808 and therefore can be decomposed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCanBeDecomposed(_:)
func CFURLCanBeDecomposed(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCanBeDecomposed(anURL)
	}


// Removes all cached resource values and temporary resource values from the URL object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLClearResourcePropertyCache(_:)
func CFURLClearResourcePropertyCache(url unsafe.Pointer) {
	_CFURLClearResourcePropertyCache(url)
	}


// Removes the cached resource value identified by a given key from the URL object. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLClearResourcePropertyCacheForKey(_:_:)
func CFURLClearResourcePropertyCacheForKey(url unsafe.Pointer, key unsafe.Pointer) {
	_CFURLClearResourcePropertyCacheForKey(url, key)
	}


// Creates a new object by resolving the relative portion of a URL against its base. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyAbsoluteURL(_:)
func CFURLCopyAbsoluteURL(relativeURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyAbsoluteURL(relativeURL)
	}


// Returns the path portion of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyFileSystemPath(_:_:)
func CFURLCopyFileSystemPath(anURL unsafe.Pointer, pathStyle unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyFileSystemPath(anURL, pathStyle)
	}


// Returns the fragment from a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyFragment(_:_:)
func CFURLCopyFragment(anURL unsafe.Pointer, charactersToLeaveEscaped unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyFragment(anURL, charactersToLeaveEscaped)
	}


// Returns the host name of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyHostName(_:)
func CFURLCopyHostName(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyHostName(anURL)
	}


// Returns the last path component of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyLastPathComponent(_:)
func CFURLCopyLastPathComponent(url unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyLastPathComponent(url)
	}


// Returns the net location portion of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyNetLocation(_:)
func CFURLCopyNetLocation(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyNetLocation(anURL)
	}


// Returns the parameter string from a given URL. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyParameterString(_:_:)
func CFURLCopyParameterString(anURL unsafe.Pointer, charactersToLeaveEscaped unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyParameterString(anURL, charactersToLeaveEscaped)
	}


// Returns the password of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPassword(_:)
func CFURLCopyPassword(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyPassword(anURL)
	}


// Returns the path portion of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPath(_:)
func CFURLCopyPath(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyPath(anURL)
	}


// Returns the path extension of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPathExtension(_:)
func CFURLCopyPathExtension(url unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyPathExtension(url)
	}


// Returns the query string of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyQueryString(_:_:)
func CFURLCopyQueryString(anURL unsafe.Pointer, charactersToLeaveEscaped unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyQueryString(anURL, charactersToLeaveEscaped)
	}


// Returns the resource values for the properties identified by specified array of keys. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourcePropertiesForKeys(_:_:_:)
func CFURLCopyResourcePropertiesForKeys(url unsafe.Pointer, keys unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyResourcePropertiesForKeys(url, keys, error_)
	}


// Returns the value of a given resource property of a given URL. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourcePropertyForKey(_:_:_:_:)
func CFURLCopyResourcePropertyForKey(url unsafe.Pointer, key unsafe.Pointer, propertyValueTypeRefPtr unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyResourcePropertyForKey(url, key, propertyValueTypeRefPtr, error_)
	}


// Returns any additional resource specifiers after the path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourceSpecifier(_:)
func CFURLCopyResourceSpecifier(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyResourceSpecifier(anURL)
	}


// Returns the scheme portion of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyScheme(_:)
func CFURLCopyScheme(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyScheme(anURL)
	}


// Returns the path portion of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyStrictPath(_:_:)
func CFURLCopyStrictPath(anURL unsafe.Pointer, isAbsolute unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyStrictPath(anURL, isAbsolute)
	}


// Returns the user name from a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyUserName(_:)
func CFURLCopyUserName(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyUserName(anURL)
	}


// Creates a new object by resolving the relative portion of a URL, specified as bytes, against its given base URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateAbsoluteURLWithBytes(_:_:_:_:_:_:)
func CFURLCreateAbsoluteURLWithBytes(alloc unsafe.Pointer, relativeURLBytes unsafe.Pointer, length unsafe.Pointer, encoding unsafe.Pointer, baseURL unsafe.Pointer, useCompatibilityMode unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateAbsoluteURLWithBytes(alloc, relativeURLBytes, length, encoding, baseURL, useCompatibilityMode)
	}


// Returns bookmark data for a URL, created with specified options and resource values. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkData(_:_:_:_:_:_:)
func CFURLCreateBookmarkData(allocator unsafe.Pointer, url unsafe.Pointer, options unsafe.Pointer, resourcePropertiesToInclude unsafe.Pointer, relativeToURL unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateBookmarkData(allocator, url, options, resourcePropertiesToInclude, relativeToURL, error_)
	}


// Initializes and returns bookmark data derived from an alias record. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkDataFromAliasRecord(_:_:)
func CFURLCreateBookmarkDataFromAliasRecord(allocatorRef unsafe.Pointer, aliasRecordDataRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateBookmarkDataFromAliasRecord(allocatorRef, aliasRecordDataRef)
	}


// Initializes and returns bookmark data derived from a file pointed to by a specified URL. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkDataFromFile(_:_:_:)
func CFURLCreateBookmarkDataFromFile(allocator unsafe.Pointer, fileURL unsafe.Pointer, errorRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateBookmarkDataFromFile(allocator, fileURL, errorRef)
	}


// Returns a new URL made by resolving bookmark data. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateByResolvingBookmarkData(_:_:_:_:_:_:_:)
func CFURLCreateByResolvingBookmarkData(allocator unsafe.Pointer, bookmark unsafe.Pointer, options unsafe.Pointer, relativeToURL unsafe.Pointer, resourcePropertiesToInclude unsafe.Pointer, isStale unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateByResolvingBookmarkData(allocator, bookmark, options, relativeToURL, resourcePropertiesToInclude, isStale, error_)
	}


// Creates a copy of a given URL and appends a path component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyAppendingPathComponent(_:_:_:_:)
func CFURLCreateCopyAppendingPathComponent(allocator unsafe.Pointer, url unsafe.Pointer, pathComponent unsafe.Pointer, isDirectory unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateCopyAppendingPathComponent(allocator, url, pathComponent, isDirectory)
	}


// Creates a copy of a given URL and appends a path extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyAppendingPathExtension(_:_:_:)
func CFURLCreateCopyAppendingPathExtension(allocator unsafe.Pointer, url unsafe.Pointer, extension unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateCopyAppendingPathExtension(allocator, url, extension)
	}


// Creates a copy of a given URL with the last path component deleted. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyDeletingLastPathComponent(_:_:)
func CFURLCreateCopyDeletingLastPathComponent(allocator unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateCopyDeletingLastPathComponent(allocator, url)
	}


// Creates a copy of a given URL with its last path extension removed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyDeletingPathExtension(_:_:)
func CFURLCreateCopyDeletingPathExtension(allocator unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateCopyDeletingPathExtension(allocator, url)
	}


// Creates a object containing the content of a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateData(_:_:_:_:)
func CFURLCreateData(allocator unsafe.Pointer, url unsafe.Pointer, encoding unsafe.Pointer, escapeWhitespace unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateData(allocator, url, encoding, escapeWhitespace)
	}


// Loads the data and properties referred to by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateDataAndPropertiesFromResource(_:_:_:_:_:_:)
func CFURLCreateDataAndPropertiesFromResource(alloc unsafe.Pointer, url unsafe.Pointer, resourceData unsafe.Pointer, properties unsafe.Pointer, desiredProperties unsafe.Pointer, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateDataAndPropertiesFromResource(alloc, url, resourceData, properties, desiredProperties, errorCode)
	}


// Returns a new file path URL that refers to the same resource as a specified URL. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFilePathURL(_:_:_:)
func CFURLCreateFilePathURL(allocator unsafe.Pointer, url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateFilePathURL(allocator, url, error_)
	}


// Returns a new file reference URL that points to the same resource as a specified URL. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFileReferenceURL(_:_:_:)
func CFURLCreateFileReferenceURL(allocator unsafe.Pointer, url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateFileReferenceURL(allocator, url, error_)
	}


// Creates a URL from a given directory or file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFSRef(_:_:)
func CFURLCreateFromFSRef(allocator unsafe.Pointer, fsRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateFromFSRef(allocator, fsRef)
	}


// Creates a new object for a file system entity using the native representation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFileSystemRepresentation(_:_:_:_:)
func CFURLCreateFromFileSystemRepresentation(allocator unsafe.Pointer, buffer unsafe.Pointer, bufLen unsafe.Pointer, isDirectory unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateFromFileSystemRepresentation(allocator, buffer, bufLen, isDirectory)
	}


// Creates a object from a native character string path relative to a base URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFileSystemRepresentationRelativeToBase(_:_:_:_:_:)
func CFURLCreateFromFileSystemRepresentationRelativeToBase(allocator unsafe.Pointer, buffer unsafe.Pointer, bufLen unsafe.Pointer, isDirectory unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateFromFileSystemRepresentationRelativeToBase(allocator, buffer, bufLen, isDirectory, baseURL)
	}


// Returns a given property specified by a given URL and property string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreatePropertyFromResource(_:_:_:_:)
func CFURLCreatePropertyFromResource(alloc unsafe.Pointer, url unsafe.Pointer, property unsafe.Pointer, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreatePropertyFromResource(alloc, url, property, errorCode)
	}


// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateResourcePropertiesForKeysFromBookmarkData(_:_:_:)
func CFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator unsafe.Pointer, resourcePropertiesToReturn unsafe.Pointer, bookmark unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator, resourcePropertiesToReturn, bookmark)
	}


// Returns the value of a resource property from specified bookmark data. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateResourcePropertyForKeyFromBookmarkData(_:_:_:)
func CFURLCreateResourcePropertyForKeyFromBookmarkData(allocator unsafe.Pointer, resourcePropertyKey unsafe.Pointer, bookmark unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateResourcePropertyForKeyFromBookmarkData(allocator, resourcePropertyKey, bookmark)
	}


// Creates a copy of a string, replacing certain characters with the equivalent percent escape sequence based on the specified encoding. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateStringByAddingPercentEscapes(_:_:_:_:_:)
func CFURLCreateStringByAddingPercentEscapes(allocator unsafe.Pointer, originalString unsafe.Pointer, charactersToLeaveUnescaped unsafe.Pointer, legalURLCharactersToBeEscaped unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateStringByAddingPercentEscapes(allocator, originalString, charactersToLeaveUnescaped, legalURLCharactersToBeEscaped, encoding)
	}


// Creates a new string by replacing any percent escape sequences with their character equivalent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateStringByReplacingPercentEscapes(_:_:_:)
func CFURLCreateStringByReplacingPercentEscapes(allocator unsafe.Pointer, originalString unsafe.Pointer, charactersToLeaveEscaped unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateStringByReplacingPercentEscapes(allocator, originalString, charactersToLeaveEscaped)
	}


// Creates a new string by replacing any percent escape sequences with their character equivalent. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateStringByReplacingPercentEscapesUsingEncoding(_:_:_:_:)
func CFURLCreateStringByReplacingPercentEscapesUsingEncoding(allocator unsafe.Pointer, origString unsafe.Pointer, charsToLeaveEscaped unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateStringByReplacingPercentEscapesUsingEncoding(allocator, origString, charsToLeaveEscaped, encoding)
	}


// Creates a object using a given character bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithBytes(_:_:_:_:_:)
func CFURLCreateWithBytes(allocator unsafe.Pointer, URLBytes unsafe.Pointer, length unsafe.Pointer, encoding unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateWithBytes(allocator, URLBytes, length, encoding, baseURL)
	}


// Creates a object using a local file system path string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithFileSystemPath(_:_:_:_:)
func CFURLCreateWithFileSystemPath(allocator unsafe.Pointer, filePath unsafe.Pointer, pathStyle unsafe.Pointer, isDirectory unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateWithFileSystemPath(allocator, filePath, pathStyle, isDirectory)
	}


// Creates a object using a local file system path string relative to a base URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithFileSystemPathRelativeToBase(_:_:_:_:_:)
func CFURLCreateWithFileSystemPathRelativeToBase(allocator unsafe.Pointer, filePath unsafe.Pointer, pathStyle unsafe.Pointer, isDirectory unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateWithFileSystemPathRelativeToBase(allocator, filePath, pathStyle, isDirectory, baseURL)
	}


// Creates a object using a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithString(_:_:_:)
func CFURLCreateWithString(allocator unsafe.Pointer, URLString unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateWithString(allocator, URLString, baseURL)
	}


// Destroys a resource indicated by a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLDestroyResource(_:_:)
func CFURLDestroyResource(url unsafe.Pointer, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLDestroyResource(url, errorCode)
	}


// Creates and returns a directory enumerator with provided enumerator behavior options and properties to be prefetched. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorCreateForDirectoryURL(_:_:_:_:)
func CFURLEnumeratorCreateForDirectoryURL(alloc unsafe.Pointer, directoryURL unsafe.Pointer, option unsafe.Pointer, propertyKeys unsafe.Pointer) unsafe.Pointer {
	return _CFURLEnumeratorCreateForDirectoryURL(alloc, directoryURL, option, propertyKeys)
	}


// Creates and returns a volume enumerator with provided enumerator behavior options and properties to be prefetched. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorCreateForMountedVolumes(_:_:_:)
func CFURLEnumeratorCreateForMountedVolumes(alloc unsafe.Pointer, option unsafe.Pointer, propertyKeys unsafe.Pointer) unsafe.Pointer {
	return _CFURLEnumeratorCreateForMountedVolumes(alloc, option, propertyKeys)
	}


// Returns the number of levels a recursive directory enumerator has descended. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetDescendentLevel(_:)
func CFURLEnumeratorGetDescendentLevel(enumerator unsafe.Pointer) unsafe.Pointer {
	return _CFURLEnumeratorGetDescendentLevel(enumerator)
	}


// Advances an enumerator to the next URL. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetNextURL(_:_:_:)
func CFURLEnumeratorGetNextURL(enumerator unsafe.Pointer, url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLEnumeratorGetNextURL(enumerator, url, error_)
	}


// This function is unimplemented, so it performs no operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetSourceDidChange(_:)
func CFURLEnumeratorGetSourceDidChange(enumerator unsafe.Pointer) unsafe.Pointer {
	return _CFURLEnumeratorGetSourceDidChange(enumerator)
	}


// Returns the opaque type identifier for the CFURLEnumerator opaque type. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetTypeID()
func CFURLEnumeratorGetTypeID() unsafe.Pointer {
	return _CFURLEnumeratorGetTypeID()
	}


// Tells a recursive enumerator not to descend into the directory at the URL that was returned by the most recent call to the function. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorSkipDescendents(_:)
func CFURLEnumeratorSkipDescendents(enumerator unsafe.Pointer) {
	_CFURLEnumeratorSkipDescendents(enumerator)
	}


// Returns the base URL of a given URL if it exists. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetBaseURL(_:)
func CFURLGetBaseURL(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetBaseURL(anURL)
	}


// Returns the range of the specified component in the bytes of a URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetByteRangeForComponent(_:_:_:)
func CFURLGetByteRangeForComponent(url unsafe.Pointer, component unsafe.Pointer, rangeIncludingSeparators unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetByteRangeForComponent(url, component, rangeIncludingSeparators)
	}


// Returns by reference the byte representation of a URL object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetBytes(_:_:_:)
func CFURLGetBytes(url unsafe.Pointer, buffer unsafe.Pointer, bufferLength unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetBytes(url, buffer, bufferLength)
	}


// Converts a given URL to a file or directory object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetFSRef(_:_:)
func CFURLGetFSRef(url unsafe.Pointer, fsRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetFSRef(url, fsRef)
	}


// Fills a buffer with the file system’s native string representation of a given URL’s path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetFileSystemRepresentation(_:_:_:_:)
func CFURLGetFileSystemRepresentation(url unsafe.Pointer, resolveAgainstBase unsafe.Pointer, buffer unsafe.Pointer, maxBufLen unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetFileSystemRepresentation(url, resolveAgainstBase, buffer, maxBufLen)
	}


// Returns the port number from a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetPortNumber(_:)
func CFURLGetPortNumber(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetPortNumber(anURL)
	}


// Returns the URL as a object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetString(_:)
func CFURLGetString(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetString(anURL)
	}


// Returns the type identifier for the opaque type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetTypeID()
func CFURLGetTypeID() unsafe.Pointer {
	return _CFURLGetTypeID()
	}


// Determines if a given URL’s path represents a directory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLHasDirectoryPath(_:)
func CFURLHasDirectoryPath(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLHasDirectoryPath(anURL)
	}


// CFURLIsFileReferenceURL is a CoreFoundation function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLIsFileReferenceURL(_:)
func CFURLIsFileReferenceURL(url unsafe.Pointer) unsafe.Pointer {
	return _CFURLIsFileReferenceURL(url)
	}


// Returns whether the resource pointed to by a file URL can be reached. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLResourceIsReachable(_:_:)
func CFURLResourceIsReachable(url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLResourceIsReachable(url, error_)
	}


// Sets the URL’s resource properties for a given set of keys to a given set of values. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLSetResourcePropertiesForKeys(_:_:_:)
func CFURLSetResourcePropertiesForKeys(url unsafe.Pointer, keyedPropertyValues unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLSetResourcePropertiesForKeys(url, keyedPropertyValues, error_)
	}


// Sets the URL’s resource property for a given key to a given value. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLSetResourcePropertyForKey(_:_:_:_:)
func CFURLSetResourcePropertyForKey(url unsafe.Pointer, key unsafe.Pointer, propertyValue unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLSetResourcePropertyForKey(url, key, propertyValue, error_)
	}


// Sets a temporary resource value on the URL. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLSetTemporaryResourcePropertyForKey(_:_:_:)
func CFURLSetTemporaryResourcePropertyForKey(url unsafe.Pointer, key unsafe.Pointer, propertyValue unsafe.Pointer) {
	_CFURLSetTemporaryResourcePropertyForKey(url, key, propertyValue)
	}


// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLStartAccessingSecurityScopedResource(_:)
func CFURLStartAccessingSecurityScopedResource(url unsafe.Pointer) unsafe.Pointer {
	return _CFURLStartAccessingSecurityScopedResource(url)
	}


// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL. [Full Topic]
//
// Added in macOS 10.7.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLStopAccessingSecurityScopedResource(_:)
func CFURLStopAccessingSecurityScopedResource(url unsafe.Pointer) {
	_CFURLStopAccessingSecurityScopedResource(url)
	}


// Creates an alias file on disk at a specified location with specified bookmark data. [Full Topic]
//
// Added in macOS 10.6.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLWriteBookmarkDataToFile(_:_:_:_:)
func CFURLWriteBookmarkDataToFile(bookmarkRef unsafe.Pointer, fileURL unsafe.Pointer, options unsafe.Pointer, errorRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLWriteBookmarkDataToFile(bookmarkRef, fileURL, options, errorRef)
	}


// Writes the given data and properties to a given URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLWriteDataAndPropertiesToResource(_:_:_:_:)
func CFURLWriteDataAndPropertiesToResource(url unsafe.Pointer, dataToWrite unsafe.Pointer, propertiesToWrite unsafe.Pointer, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLWriteDataAndPropertiesToResource(url, dataToWrite, propertiesToWrite, errorCode)
	}


// Creates a Universally Unique Identifier (UUID) object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreate(_:)
func CFUUIDCreate(alloc unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDCreate(alloc)
	}


// Creates a CFUUID object for a specified string. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateFromString(_:_:)
func CFUUIDCreateFromString(alloc unsafe.Pointer, uuidStr unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDCreateFromString(alloc, uuidStr)
	}


// Creates a CFUUID object from raw UUID bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateFromUUIDBytes(_:_:)
func CFUUIDCreateFromUUIDBytes(alloc unsafe.Pointer, bytes unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDCreateFromUUIDBytes(alloc, bytes)
	}


// Returns the string representation of a specified CFUUID object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateString(_:_:)
func CFUUIDCreateString(alloc unsafe.Pointer, uuid unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDCreateString(alloc, uuid)
	}


// Creates a CFUUID object from raw UUID bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateWithBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CFUUIDCreateWithBytes(alloc unsafe.Pointer, byte0 unsafe.Pointer, byte1 unsafe.Pointer, byte2 unsafe.Pointer, byte3 unsafe.Pointer, byte4 unsafe.Pointer, byte5 unsafe.Pointer, byte6 unsafe.Pointer, byte7 unsafe.Pointer, byte8 unsafe.Pointer, byte9 unsafe.Pointer, byte10 unsafe.Pointer, byte11 unsafe.Pointer, byte12 unsafe.Pointer, byte13 unsafe.Pointer, byte14 unsafe.Pointer, byte15 unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDCreateWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15)
	}


// Returns a CFUUID object from raw UUID bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetConstantUUIDWithBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CFUUIDGetConstantUUIDWithBytes(alloc unsafe.Pointer, byte0 unsafe.Pointer, byte1 unsafe.Pointer, byte2 unsafe.Pointer, byte3 unsafe.Pointer, byte4 unsafe.Pointer, byte5 unsafe.Pointer, byte6 unsafe.Pointer, byte7 unsafe.Pointer, byte8 unsafe.Pointer, byte9 unsafe.Pointer, byte10 unsafe.Pointer, byte11 unsafe.Pointer, byte12 unsafe.Pointer, byte13 unsafe.Pointer, byte14 unsafe.Pointer, byte15 unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDGetConstantUUIDWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15)
	}


// Returns the type identifier for all CFUUID objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetTypeID()
func CFUUIDGetTypeID() unsafe.Pointer {
	return _CFUUIDGetTypeID()
	}


// Returns the value of a UUID object as raw bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetUUIDBytes(_:)
func CFUUIDGetUUIDBytes(uuid unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDGetUUIDBytes(uuid)
	}


// Cancels a user notification dialog. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCancel(_:)
func CFUserNotificationCancel(userNotification unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationCancel(userNotification)
	}


// Creates a CFUserNotification object and displays its notification dialog on screen. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCreate(_:_:_:_:_:)
func CFUserNotificationCreate(allocator unsafe.Pointer, timeout unsafe.Pointer, flags unsafe.Pointer, error_ unsafe.Pointer, dictionary unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationCreate(allocator, timeout, flags, error_, dictionary)
	}


// Creates a run loop source for a user notification. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCreateRunLoopSource(_:_:_:_:)
func CFUserNotificationCreateRunLoopSource(allocator unsafe.Pointer, userNotification unsafe.Pointer, callout unsafe.Pointer, order unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationCreateRunLoopSource(allocator, userNotification, callout, order)
	}


// Displays a user notification dialog and waits for a user response. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationDisplayAlert(_:_:_:_:_:_:_:_:_:_:_:)
func CFUserNotificationDisplayAlert(timeout unsafe.Pointer, flags unsafe.Pointer, iconURL unsafe.Pointer, soundURL unsafe.Pointer, localizationURL unsafe.Pointer, alertHeader unsafe.Pointer, alertMessage unsafe.Pointer, defaultButtonTitle unsafe.Pointer, alternateButtonTitle unsafe.Pointer, otherButtonTitle unsafe.Pointer, responseFlags unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationDisplayAlert(timeout, flags, iconURL, soundURL, localizationURL, alertHeader, alertMessage, defaultButtonTitle, alternateButtonTitle, otherButtonTitle, responseFlags)
	}


// Displays a user notification dialog that does not need a user response. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationDisplayNotice(_:_:_:_:_:_:_:_:)
func CFUserNotificationDisplayNotice(timeout unsafe.Pointer, flags unsafe.Pointer, iconURL unsafe.Pointer, soundURL unsafe.Pointer, localizationURL unsafe.Pointer, alertHeader unsafe.Pointer, alertMessage unsafe.Pointer, defaultButtonTitle unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationDisplayNotice(timeout, flags, iconURL, soundURL, localizationURL, alertHeader, alertMessage, defaultButtonTitle)
	}


// Returns the dictionary containing all the text field values from a dismissed notification dialog. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetResponseDictionary(_:)
func CFUserNotificationGetResponseDictionary(userNotification unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationGetResponseDictionary(userNotification)
	}


// Extracts the values of the text fields from a dismissed notification dialog. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetResponseValue(_:_:_:)
func CFUserNotificationGetResponseValue(userNotification unsafe.Pointer, key unsafe.Pointer, idx unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationGetResponseValue(userNotification, key, idx)
	}


// Returns the type identifier for the opaque type. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetTypeID()
func CFUserNotificationGetTypeID() unsafe.Pointer {
	return _CFUserNotificationGetTypeID()
	}


// Waits for the user to respond to a notification or for the notification to time out. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationReceiveResponse(_:_:_:)
func CFUserNotificationReceiveResponse(userNotification unsafe.Pointer, timeout unsafe.Pointer, responseFlags unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationReceiveResponse(userNotification, timeout, responseFlags)
	}


// Updates a displayed user notification dialog with new user interface information. [Full Topic]
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationUpdate(_:_:_:_:)
func CFUserNotificationUpdate(userNotification unsafe.Pointer, timeout unsafe.Pointer, flags unsafe.Pointer, dictionary unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationUpdate(userNotification, timeout, flags, dictionary)
	}


// Returns whether a writable stream can accept new data without blocking. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCanAcceptBytes(_:)
func CFWriteStreamCanAcceptBytes(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCanAcceptBytes(stream)
	}


// Closes a writable stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamClose(_:)
func CFWriteStreamClose(stream unsafe.Pointer) {
	_CFWriteStreamClose(stream)
	}


// CFWriteStreamCopyDispatchQueue is a CoreFoundation function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyDispatchQueue(_:)
func CFWriteStreamCopyDispatchQueue(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCopyDispatchQueue(stream)
	}


// Returns the error associated with a stream. [Full Topic]
//
// Added in macOS 10.5.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyError(_:)
func CFWriteStreamCopyError(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCopyError(stream)
	}


// Returns the value of a property for a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyProperty(_:_:)
func CFWriteStreamCopyProperty(stream unsafe.Pointer, propertyName unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCopyProperty(stream, propertyName)
	}


// Creates a writable stream for a growable block of memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithAllocatedBuffers(_:_:)
func CFWriteStreamCreateWithAllocatedBuffers(alloc unsafe.Pointer, bufferAllocator unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCreateWithAllocatedBuffers(alloc, bufferAllocator)
	}


// Creates a writable stream for a fixed-size block of memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithBuffer(_:_:_:)
func CFWriteStreamCreateWithBuffer(alloc unsafe.Pointer, buffer unsafe.Pointer, bufferCapacity unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCreateWithBuffer(alloc, buffer, bufferCapacity)
	}


// Creates a writable stream for a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithFile(_:_:)
func CFWriteStreamCreateWithFile(alloc unsafe.Pointer, fileURL unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCreateWithFile(alloc, fileURL)
	}


// Returns the error status of a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetError(_:)
func CFWriteStreamGetError(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamGetError(stream)
	}


// Returns the current state of a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetStatus(_:)
func CFWriteStreamGetStatus(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamGetStatus(stream)
	}


// Returns the type identifier of all CFWriteStream objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetTypeID()
func CFWriteStreamGetTypeID() unsafe.Pointer {
	return _CFWriteStreamGetTypeID()
	}


// Opens a stream for writing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamOpen(_:)
func CFWriteStreamOpen(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamOpen(stream)
	}


// Schedules a stream into a run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamScheduleWithRunLoop(_:_:_:)
func CFWriteStreamScheduleWithRunLoop(stream unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_CFWriteStreamScheduleWithRunLoop(stream, runLoop, runLoopMode)
	}


// Assigns a client to a stream, which receives callbacks when certain events occur. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetClient(_:_:_:_:)
func CFWriteStreamSetClient(stream unsafe.Pointer, streamEvents unsafe.Pointer, clientCB unsafe.Pointer, clientContext unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamSetClient(stream, streamEvents, clientCB, clientContext)
	}


// CFWriteStreamSetDispatchQueue is a CoreFoundation function. [Full Topic]
//
// Added in macOS 10.9.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetDispatchQueue(_:_:)
func CFWriteStreamSetDispatchQueue(stream unsafe.Pointer, q unsafe.Pointer) {
	_CFWriteStreamSetDispatchQueue(stream, q)
	}


// Sets the value of a property for a stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetProperty(_:_:_:)
func CFWriteStreamSetProperty(stream unsafe.Pointer, propertyName unsafe.Pointer, propertyValue unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamSetProperty(stream, propertyName, propertyValue)
	}


// Removes a stream from a particular run loop. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamUnscheduleFromRunLoop(_:_:_:)
func CFWriteStreamUnscheduleFromRunLoop(stream unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_CFWriteStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode)
	}


// Writes data to a writable stream. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamWrite(_:_:_:)
func CFWriteStreamWrite(stream unsafe.Pointer, buffer unsafe.Pointer, bufferLength unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamWrite(stream, buffer, bufferLength)
	}


// Given a CFString object containing XML source with unescaped entities, returns a string with specified XML entities escaped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLCreateStringByEscapingEntities(_:_:_:)
func CFXMLCreateStringByEscapingEntities(allocator unsafe.Pointer, string_ unsafe.Pointer, entitiesDictionary unsafe.Pointer) unsafe.Pointer {
	return _CFXMLCreateStringByEscapingEntities(allocator, string_, entitiesDictionary)
	}


// Given a CFString object containing XML source with escaped entities, returns a string with specified XML entities unescaped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLCreateStringByUnescapingEntities(_:_:_:)
func CFXMLCreateStringByUnescapingEntities(allocator unsafe.Pointer, string_ unsafe.Pointer, entitiesDictionary unsafe.Pointer) unsafe.Pointer {
	return _CFXMLCreateStringByUnescapingEntities(allocator, string_, entitiesDictionary)
	}


// Creates a new CFXMLNode. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeCreate
func CFXMLNodeCreate(alloc unsafe.Pointer, xmlType unsafe.Pointer, dataString unsafe.Pointer, additionalInfoPtr unsafe.Pointer, version unsafe.Pointer) unsafe.Pointer {
	return _CFXMLNodeCreate(alloc, xmlType, dataString, additionalInfoPtr, version)
	}


// Creates a copy of a CFXMLNode object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeCreateCopy
func CFXMLNodeCreateCopy(alloc unsafe.Pointer, origNode unsafe.Pointer) unsafe.Pointer {
	return _CFXMLNodeCreateCopy(alloc, origNode)
	}


// Returns the additional information pointer of a CFXMLNode object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetInfoPtr
func CFXMLNodeGetInfoPtr(node unsafe.Pointer) unsafe.Pointer {
	return _CFXMLNodeGetInfoPtr(node)
	}


// Returns the data string from a CFXMLNode. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetString
func CFXMLNodeGetString(node unsafe.Pointer) unsafe.Pointer {
	return _CFXMLNodeGetString(node)
	}


// Returns the XML structure type code for a CFXMLNode object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetTypeCode
func CFXMLNodeGetTypeCode(node unsafe.Pointer) unsafe.Pointer {
	return _CFXMLNodeGetTypeCode(node)
	}


// Returns the type identifier code for the CFXMLNode opaque type. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetTypeID
func CFXMLNodeGetTypeID() unsafe.Pointer {
	return _CFXMLNodeGetTypeID()
	}


// Returns the version number for a CFXMLNode object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetVersion
func CFXMLNodeGetVersion(node unsafe.Pointer) unsafe.Pointer {
	return _CFXMLNodeGetVersion(node)
	}


// Causes a parser to abort with the given error code and description. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserAbort
func CFXMLParserAbort(parser unsafe.Pointer, errorCode unsafe.Pointer, errorDescription unsafe.Pointer) {
	_CFXMLParserAbort(parser, errorCode, errorDescription)
	}


// Returns the user-readable description of the current error condition. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCopyErrorDescription
func CFXMLParserCopyErrorDescription(parser unsafe.Pointer) unsafe.Pointer {
	return _CFXMLParserCopyErrorDescription(parser)
	}


// Creates a new XML parser for the specified XML data. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCreate
func CFXMLParserCreate(allocator unsafe.Pointer, xmlData unsafe.Pointer, dataSource unsafe.Pointer, parseOptions unsafe.Pointer, versionOfNodes unsafe.Pointer, callBacks unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFXMLParserCreate(allocator, xmlData, dataSource, parseOptions, versionOfNodes, callBacks, context)
	}


// Creates a new XML parser for the specified XML data at the specified URL. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCreateWithDataFromURL
func CFXMLParserCreateWithDataFromURL(allocator unsafe.Pointer, dataSource unsafe.Pointer, parseOptions unsafe.Pointer, versionOfNodes unsafe.Pointer, callBacks unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFXMLParserCreateWithDataFromURL(allocator, dataSource, parseOptions, versionOfNodes, callBacks, context)
	}


// Returns the callbacks associated with an XML parser when it was created. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetCallBacks
func CFXMLParserGetCallBacks(parser unsafe.Pointer, callBacks unsafe.Pointer) {
	_CFXMLParserGetCallBacks(parser, callBacks)
	}


// Returns the context for an XML parser. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetContext
func CFXMLParserGetContext(parser unsafe.Pointer, context unsafe.Pointer) {
	_CFXMLParserGetContext(parser, context)
	}


// Returns the top-most object returned by the create XML structure callback. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetDocument
func CFXMLParserGetDocument(parser unsafe.Pointer) unsafe.Pointer {
	return _CFXMLParserGetDocument(parser)
	}


// Returns the line number of the current parse location. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetLineNumber
func CFXMLParserGetLineNumber(parser unsafe.Pointer) unsafe.Pointer {
	return _CFXMLParserGetLineNumber(parser)
	}


// Returns the character index of the current parse location. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetLocation
func CFXMLParserGetLocation(parser unsafe.Pointer) unsafe.Pointer {
	return _CFXMLParserGetLocation(parser)
	}


// Returns the URL for the XML data being parsed. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetSourceURL
func CFXMLParserGetSourceURL(parser unsafe.Pointer) unsafe.Pointer {
	return _CFXMLParserGetSourceURL(parser)
	}


// Returns a numeric code indicating the current status of the parser. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetStatusCode
func CFXMLParserGetStatusCode(parser unsafe.Pointer) unsafe.Pointer {
	return _CFXMLParserGetStatusCode(parser)
	}


// Returns the type identifier for the CFXMLParser opaque type. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetTypeID
func CFXMLParserGetTypeID() unsafe.Pointer {
	return _CFXMLParserGetTypeID()
	}


// Begins a parse of the XML data that was associated with the parser when it was created. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserParse
func CFXMLParserParse(parser unsafe.Pointer) unsafe.Pointer {
	return _CFXMLParserParse(parser)
	}


// Parses the given XML data and returns the resulting CFXMLTree object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateFromData
func CFXMLTreeCreateFromData(allocator unsafe.Pointer, xmlData unsafe.Pointer, dataSource unsafe.Pointer, parseOptions unsafe.Pointer, versionOfNodes unsafe.Pointer) unsafe.Pointer {
	return _CFXMLTreeCreateFromData(allocator, xmlData, dataSource, parseOptions, versionOfNodes)
	}


// Parses the given XML data and returns the resulting CFXMLTree object and any error information. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateFromDataWithError
func CFXMLTreeCreateFromDataWithError(allocator unsafe.Pointer, xmlData unsafe.Pointer, dataSource unsafe.Pointer, parseOptions unsafe.Pointer, versionOfNodes unsafe.Pointer, errorDict unsafe.Pointer) unsafe.Pointer {
	return _CFXMLTreeCreateFromDataWithError(allocator, xmlData, dataSource, parseOptions, versionOfNodes, errorDict)
	}


// Creates a new CFXMLTree object by loading the data to be parsed directly from a data source. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateWithDataFromURL
func CFXMLTreeCreateWithDataFromURL(allocator unsafe.Pointer, dataSource unsafe.Pointer, parseOptions unsafe.Pointer, versionOfNodes unsafe.Pointer) unsafe.Pointer {
	return _CFXMLTreeCreateWithDataFromURL(allocator, dataSource, parseOptions, versionOfNodes)
	}


// Creates a childless, parentless CFXMLTree object node for a CFXMLNode object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateWithNode
func CFXMLTreeCreateWithNode(allocator unsafe.Pointer, node unsafe.Pointer) unsafe.Pointer {
	return _CFXMLTreeCreateWithNode(allocator, node)
	}


// Generates an XML document from a CFXMLTree object which is ready to be written to permanent storage. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateXMLData
func CFXMLTreeCreateXMLData(allocator unsafe.Pointer, xmlTree unsafe.Pointer) unsafe.Pointer {
	return _CFXMLTreeCreateXMLData(allocator, xmlTree)
	}


// Returns the node of a CFXMLTree object. [Full Topic]
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeGetNode
func CFXMLTreeGetNode(xmlTree unsafe.Pointer) unsafe.Pointer {
	return _CFXMLTreeGetNode(xmlTree)
	}


// inset is a CoreFoundation function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRect/inset(by:)
func inset(insets unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _inset(insets, p1)
	}




