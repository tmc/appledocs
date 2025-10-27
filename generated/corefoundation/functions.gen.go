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
	_CFAbsoluteTimeAddGregorianUnits func(AbsoluteTime, TimeZoneRef, GregorianUnits) AbsoluteTime
	_CFAbsoluteTimeGetCurrent func() AbsoluteTime
	_CFAbsoluteTimeGetDayOfWeek func(AbsoluteTime, TimeZoneRef) unsafe.Pointer
	_CFAbsoluteTimeGetDayOfYear func(AbsoluteTime, TimeZoneRef) unsafe.Pointer
	_CFAbsoluteTimeGetDifferenceAsGregorianUnits func(AbsoluteTime, AbsoluteTime, TimeZoneRef, OptionFlags) GregorianUnits
	_CFAbsoluteTimeGetGregorianDate func(AbsoluteTime, TimeZoneRef) GregorianDate
	_CFAbsoluteTimeGetWeekOfYear func(AbsoluteTime, TimeZoneRef) unsafe.Pointer
	_CFAllocatorAllocate func(AllocatorRef, Index, OptionFlags) unsafe.Pointer
	_CFAllocatorAllocateBytes func(AllocatorRef, Index, OptionFlags) unsafe.Pointer
	_CFAllocatorAllocateTyped func(AllocatorRef, Index, AllocatorTypeID, OptionFlags) unsafe.Pointer
	_CFAllocatorCreate func(AllocatorRef, unsafe.Pointer) AllocatorRef
	_CFAllocatorCreateWithZone func(AllocatorRef, unsafe.Pointer) AllocatorRef
	_CFAllocatorDeallocate func(AllocatorRef, unsafe.Pointer)
	_CFAllocatorGetContext func(AllocatorRef, unsafe.Pointer)
	_CFAllocatorGetDefault func() AllocatorRef
	_CFAllocatorGetPreferredSizeForSize func(AllocatorRef, Index, OptionFlags) Index
	_CFAllocatorGetTypeID func() TypeID
	_CFAllocatorReallocate func(AllocatorRef, unsafe.Pointer, Index, OptionFlags) unsafe.Pointer
	_CFAllocatorReallocateBytes func(AllocatorRef, unsafe.Pointer, Index, OptionFlags) unsafe.Pointer
	_CFAllocatorReallocateTyped func(AllocatorRef, unsafe.Pointer, Index, AllocatorTypeID, OptionFlags) unsafe.Pointer
	_CFAllocatorSetDefault func(AllocatorRef)
	_CFArrayAppendArray func(MutableArrayRef, ArrayRef, Range)
	_CFArrayAppendValue func(MutableArrayRef, unsafe.Pointer)
	_CFArrayApplyFunction func(ArrayRef, Range, ArrayApplierFunction, unsafe.Pointer)
	_CFArrayBSearchValues func(ArrayRef, Range, unsafe.Pointer, ComparatorFunction, unsafe.Pointer) Index
	_CFArrayContainsValue func(ArrayRef, Range, unsafe.Pointer) unsafe.Pointer
	_CFArrayCreate func(AllocatorRef, unsafe.Pointer, Index, unsafe.Pointer) ArrayRef
	_CFArrayCreateCopy func(AllocatorRef, ArrayRef) ArrayRef
	_CFArrayCreateMutable func(AllocatorRef, Index, unsafe.Pointer) MutableArrayRef
	_CFArrayCreateMutableCopy func(AllocatorRef, Index, ArrayRef) MutableArrayRef
	_CFArrayExchangeValuesAtIndices func(MutableArrayRef, Index, Index)
	_CFArrayGetCount func(ArrayRef) Index
	_CFArrayGetCountOfValue func(ArrayRef, Range, unsafe.Pointer) Index
	_CFArrayGetFirstIndexOfValue func(ArrayRef, Range, unsafe.Pointer) Index
	_CFArrayGetLastIndexOfValue func(ArrayRef, Range, unsafe.Pointer) Index
	_CFArrayGetTypeID func() TypeID
	_CFArrayGetValueAtIndex func(ArrayRef, Index) unsafe.Pointer
	_CFArrayGetValues func(ArrayRef, Range, unsafe.Pointer)
	_CFArrayInsertValueAtIndex func(MutableArrayRef, Index, unsafe.Pointer)
	_CFArrayRemoveAllValues func(MutableArrayRef)
	_CFArrayRemoveValueAtIndex func(MutableArrayRef, Index)
	_CFArrayReplaceValues func(MutableArrayRef, Range, unsafe.Pointer, Index)
	_CFArraySetValueAtIndex func(MutableArrayRef, Index, unsafe.Pointer)
	_CFArraySortValues func(MutableArrayRef, Range, ComparatorFunction, unsafe.Pointer)
	_CFAttributedStringBeginEditing func(MutableAttributedStringRef)
	_CFAttributedStringCreate func(AllocatorRef, StringRef, DictionaryRef) AttributedStringRef
	_CFAttributedStringCreateCopy func(AllocatorRef, AttributedStringRef) AttributedStringRef
	_CFAttributedStringCreateMutable func(AllocatorRef, Index) MutableAttributedStringRef
	_CFAttributedStringCreateMutableCopy func(AllocatorRef, Index, AttributedStringRef) MutableAttributedStringRef
	_CFAttributedStringCreateWithSubstring func(AllocatorRef, AttributedStringRef, Range) AttributedStringRef
	_CFAttributedStringEndEditing func(MutableAttributedStringRef)
	_CFAttributedStringGetAttribute func(AttributedStringRef, Index, StringRef, unsafe.Pointer) TypeRef
	_CFAttributedStringGetAttributeAndLongestEffectiveRange func(AttributedStringRef, Index, StringRef, Range, unsafe.Pointer) TypeRef
	_CFAttributedStringGetAttributes func(AttributedStringRef, Index, unsafe.Pointer) DictionaryRef
	_CFAttributedStringGetAttributesAndLongestEffectiveRange func(AttributedStringRef, Index, Range, unsafe.Pointer) DictionaryRef
	_CFAttributedStringGetBidiLevelsAndResolvedDirections func(AttributedStringRef, Range, int8, unsafe.Pointer, unsafe.Pointer) bool
	_CFAttributedStringGetLength func(AttributedStringRef) Index
	_CFAttributedStringGetMutableString func(MutableAttributedStringRef) MutableStringRef
	_CFAttributedStringGetStatisticalWritingDirections func(AttributedStringRef, Range, int8, unsafe.Pointer, unsafe.Pointer) bool
	_CFAttributedStringGetString func(AttributedStringRef) StringRef
	_CFAttributedStringGetTypeID func() TypeID
	_CFAttributedStringRemoveAttribute func(MutableAttributedStringRef, Range, StringRef)
	_CFAttributedStringReplaceAttributedString func(MutableAttributedStringRef, Range, AttributedStringRef)
	_CFAttributedStringReplaceString func(MutableAttributedStringRef, Range, StringRef)
	_CFAttributedStringSetAttribute func(MutableAttributedStringRef, Range, StringRef, TypeRef)
	_CFAttributedStringSetAttributes func(MutableAttributedStringRef, Range, DictionaryRef, unsafe.Pointer)
	_CFAutorelease func(TypeRef) TypeRef
	_CFBagAddValue func(MutableBagRef, unsafe.Pointer)
	_CFBagApplyFunction func(BagRef, BagApplierFunction, unsafe.Pointer)
	_CFBagContainsValue func(BagRef, unsafe.Pointer) unsafe.Pointer
	_CFBagCreate func(AllocatorRef, unsafe.Pointer, Index, unsafe.Pointer) BagRef
	_CFBagCreateCopy func(AllocatorRef, BagRef) BagRef
	_CFBagCreateMutable func(AllocatorRef, Index, unsafe.Pointer) MutableBagRef
	_CFBagCreateMutableCopy func(AllocatorRef, Index, BagRef) MutableBagRef
	_CFBagGetCount func(BagRef) Index
	_CFBagGetCountOfValue func(BagRef, unsafe.Pointer) Index
	_CFBagGetTypeID func() TypeID
	_CFBagGetValue func(BagRef, unsafe.Pointer) unsafe.Pointer
	_CFBagGetValueIfPresent func(BagRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagGetValues func(BagRef, unsafe.Pointer)
	_CFBagRemoveAllValues func(MutableBagRef)
	_CFBagRemoveValue func(MutableBagRef, unsafe.Pointer)
	_CFBagReplaceValue func(MutableBagRef, unsafe.Pointer)
	_CFBagSetValue func(MutableBagRef, unsafe.Pointer)
	_CFBinaryHeapAddValue func(BinaryHeapRef, unsafe.Pointer)
	_CFBinaryHeapApplyFunction func(BinaryHeapRef, BinaryHeapApplierFunction, unsafe.Pointer)
	_CFBinaryHeapContainsValue func(BinaryHeapRef, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapCreate func(AllocatorRef, Index, unsafe.Pointer, unsafe.Pointer) BinaryHeapRef
	_CFBinaryHeapCreateCopy func(AllocatorRef, Index, BinaryHeapRef) BinaryHeapRef
	_CFBinaryHeapGetCount func(BinaryHeapRef) Index
	_CFBinaryHeapGetCountOfValue func(BinaryHeapRef, unsafe.Pointer) Index
	_CFBinaryHeapGetMinimum func(BinaryHeapRef) unsafe.Pointer
	_CFBinaryHeapGetMinimumIfPresent func(BinaryHeapRef, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetTypeID func() TypeID
	_CFBinaryHeapGetValues func(BinaryHeapRef, unsafe.Pointer)
	_CFBinaryHeapRemoveAllValues func(BinaryHeapRef)
	_CFBinaryHeapRemoveMinimumValue func(BinaryHeapRef)
	_CFBitVectorContainsBit func(BitVectorRef, Range, Bit) unsafe.Pointer
	_CFBitVectorCreate func(AllocatorRef, unsafe.Pointer, Index) BitVectorRef
	_CFBitVectorCreateCopy func(AllocatorRef, BitVectorRef) BitVectorRef
	_CFBitVectorCreateMutable func(AllocatorRef, Index) MutableBitVectorRef
	_CFBitVectorCreateMutableCopy func(AllocatorRef, Index, BitVectorRef) MutableBitVectorRef
	_CFBitVectorFlipBitAtIndex func(MutableBitVectorRef, Index)
	_CFBitVectorFlipBits func(MutableBitVectorRef, Range)
	_CFBitVectorGetBitAtIndex func(BitVectorRef, Index) Bit
	_CFBitVectorGetBits func(BitVectorRef, Range, unsafe.Pointer)
	_CFBitVectorGetCount func(BitVectorRef) Index
	_CFBitVectorGetCountOfBit func(BitVectorRef, Range, Bit) Index
	_CFBitVectorGetFirstIndexOfBit func(BitVectorRef, Range, Bit) Index
	_CFBitVectorGetLastIndexOfBit func(BitVectorRef, Range, Bit) Index
	_CFBitVectorGetTypeID func() TypeID
	_CFBitVectorSetAllBits func(MutableBitVectorRef, Bit)
	_CFBitVectorSetBitAtIndex func(MutableBitVectorRef, Index, Bit)
	_CFBitVectorSetBits func(MutableBitVectorRef, Range, Bit)
	_CFBitVectorSetCount func(MutableBitVectorRef, Index)
	_CFBooleanGetTypeID func() TypeID
	_CFBooleanGetValue func(BooleanRef) unsafe.Pointer
	_CFBundleCloseBundleResourceMap func(BundleRef, BundleRefNum)
	_CFBundleCopyAuxiliaryExecutableURL func(BundleRef, StringRef) URLRef
	_CFBundleCopyBuiltInPlugInsURL func(BundleRef) URLRef
	_CFBundleCopyBundleLocalizations func(BundleRef) ArrayRef
	_CFBundleCopyBundleURL func(BundleRef) URLRef
	_CFBundleCopyExecutableArchitectures func(BundleRef) ArrayRef
	_CFBundleCopyExecutableArchitecturesForURL func(URLRef) ArrayRef
	_CFBundleCopyExecutableURL func(BundleRef) URLRef
	_CFBundleCopyInfoDictionaryForURL func(URLRef) DictionaryRef
	_CFBundleCopyInfoDictionaryInDirectory func(URLRef) DictionaryRef
	_CFBundleCopyLocalizationsForPreferences func(ArrayRef, ArrayRef) ArrayRef
	_CFBundleCopyLocalizationsForURL func(URLRef) ArrayRef
	_CFBundleCopyLocalizedString func(BundleRef, StringRef, StringRef, StringRef) StringRef
	_CFBundleCopyLocalizedStringForLocalizations func(BundleRef, StringRef, StringRef, StringRef, ArrayRef) StringRef
	_CFBundleCopyPreferredLocalizationsFromArray func(ArrayRef) ArrayRef
	_CFBundleCopyPrivateFrameworksURL func(BundleRef) URLRef
	_CFBundleCopyResourceURL func(BundleRef, StringRef, StringRef, StringRef) URLRef
	_CFBundleCopyResourceURLForLocalization func(BundleRef, StringRef, StringRef, StringRef, StringRef) URLRef
	_CFBundleCopyResourceURLInDirectory func(URLRef, StringRef, StringRef, StringRef) URLRef
	_CFBundleCopyResourceURLsOfType func(BundleRef, StringRef, StringRef) ArrayRef
	_CFBundleCopyResourceURLsOfTypeForLocalization func(BundleRef, StringRef, StringRef, StringRef) ArrayRef
	_CFBundleCopyResourceURLsOfTypeInDirectory func(URLRef, StringRef, StringRef) ArrayRef
	_CFBundleCopyResourcesDirectoryURL func(BundleRef) URLRef
	_CFBundleCopySharedFrameworksURL func(BundleRef) URLRef
	_CFBundleCopySharedSupportURL func(BundleRef) URLRef
	_CFBundleCopySupportFilesDirectoryURL func(BundleRef) URLRef
	_CFBundleCreate func(AllocatorRef, URLRef) BundleRef
	_CFBundleCreateBundlesFromDirectory func(AllocatorRef, URLRef, StringRef) ArrayRef
	_CFBundleGetAllBundles func() ArrayRef
	_CFBundleGetBundleWithIdentifier func(StringRef) BundleRef
	_CFBundleGetDataPointerForName func(BundleRef, StringRef) unsafe.Pointer
	_CFBundleGetDataPointersForNames func(BundleRef, ArrayRef, unsafe.Pointer)
	_CFBundleGetDevelopmentRegion func(BundleRef) StringRef
	_CFBundleGetFunctionPointerForName func(BundleRef, StringRef) unsafe.Pointer
	_CFBundleGetFunctionPointersForNames func(BundleRef, ArrayRef, unsafe.Pointer)
	_CFBundleGetIdentifier func(BundleRef) StringRef
	_CFBundleGetInfoDictionary func(BundleRef) DictionaryRef
	_CFBundleGetLocalInfoDictionary func(BundleRef) DictionaryRef
	_CFBundleGetMainBundle func() BundleRef
	_CFBundleGetPackageInfo func(BundleRef, unsafe.Pointer, unsafe.Pointer)
	_CFBundleGetPackageInfoInDirectory func(URLRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetPlugIn func(BundleRef) PlugInRef
	_CFBundleGetTypeID func() TypeID
	_CFBundleGetValueForInfoDictionaryKey func(BundleRef, StringRef) TypeRef
	_CFBundleGetVersionNumber func(BundleRef) unsafe.Pointer
	_CFBundleIsArchitectureLoadable func(unsafe.Pointer) unsafe.Pointer
	_CFBundleIsExecutableLoadable func(BundleRef) unsafe.Pointer
	_CFBundleIsExecutableLoadableForURL func(URLRef) unsafe.Pointer
	_CFBundleIsExecutableLoaded func(BundleRef) unsafe.Pointer
	_CFBundleLoadExecutable func(BundleRef) unsafe.Pointer
	_CFBundleLoadExecutableAndReturnError func(BundleRef, unsafe.Pointer) unsafe.Pointer
	_CFBundleOpenBundleResourceFiles func(BundleRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleOpenBundleResourceMap func(BundleRef) BundleRefNum
	_CFBundlePreflightExecutable func(BundleRef, unsafe.Pointer) unsafe.Pointer
	_CFBundleUnloadExecutable func(BundleRef)
	_CFCalendarAddComponents func(CalendarRef, unsafe.Pointer, OptionFlags, unsafe.Pointer) unsafe.Pointer
	_CFCalendarComposeAbsoluteTime func(CalendarRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarCopyCurrent func() CalendarRef
	_CFCalendarCopyLocale func(CalendarRef) LocaleRef
	_CFCalendarCopyTimeZone func(CalendarRef) TimeZoneRef
	_CFCalendarCreateWithIdentifier func(AllocatorRef, CalendarIdentifier) CalendarRef
	_CFCalendarDecomposeAbsoluteTime func(CalendarRef, AbsoluteTime, unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetComponentDifference func(CalendarRef, AbsoluteTime, AbsoluteTime, OptionFlags, unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetFirstWeekday func(CalendarRef) Index
	_CFCalendarGetIdentifier func(CalendarRef) CalendarIdentifier
	_CFCalendarGetMaximumRangeOfUnit func(CalendarRef, CalendarUnit) Range
	_CFCalendarGetMinimumDaysInFirstWeek func(CalendarRef) Index
	_CFCalendarGetMinimumRangeOfUnit func(CalendarRef, CalendarUnit) Range
	_CFCalendarGetOrdinalityOfUnit func(CalendarRef, CalendarUnit, CalendarUnit, AbsoluteTime) Index
	_CFCalendarGetRangeOfUnit func(CalendarRef, CalendarUnit, CalendarUnit, AbsoluteTime) Range
	_CFCalendarGetTimeRangeOfUnit func(CalendarRef, CalendarUnit, AbsoluteTime, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCalendarGetTypeID func() TypeID
	_CFCalendarSetFirstWeekday func(CalendarRef, Index)
	_CFCalendarSetLocale func(CalendarRef, LocaleRef)
	_CFCalendarSetMinimumDaysInFirstWeek func(CalendarRef, Index)
	_CFCalendarSetTimeZone func(CalendarRef, TimeZoneRef)
	_CFCharacterSetAddCharactersInRange func(MutableCharacterSetRef, Range)
	_CFCharacterSetAddCharactersInString func(MutableCharacterSetRef, StringRef)
	_CFCharacterSetCreateBitmapRepresentation func(AllocatorRef, CharacterSetRef) DataRef
	_CFCharacterSetCreateCopy func(AllocatorRef, CharacterSetRef) CharacterSetRef
	_CFCharacterSetCreateInvertedSet func(AllocatorRef, CharacterSetRef) CharacterSetRef
	_CFCharacterSetCreateMutable func(AllocatorRef) MutableCharacterSetRef
	_CFCharacterSetCreateMutableCopy func(AllocatorRef, CharacterSetRef) MutableCharacterSetRef
	_CFCharacterSetCreateWithBitmapRepresentation func(AllocatorRef, DataRef) CharacterSetRef
	_CFCharacterSetCreateWithCharactersInRange func(AllocatorRef, Range) CharacterSetRef
	_CFCharacterSetCreateWithCharactersInString func(AllocatorRef, StringRef) CharacterSetRef
	_CFCharacterSetGetPredefined func(CharacterSetPredefinedSet) CharacterSetRef
	_CFCharacterSetGetTypeID func() TypeID
	_CFCharacterSetHasMemberInPlane func(CharacterSetRef, Index) unsafe.Pointer
	_CFCharacterSetIntersect func(MutableCharacterSetRef, CharacterSetRef)
	_CFCharacterSetInvert func(MutableCharacterSetRef)
	_CFCharacterSetIsCharacterMember func(CharacterSetRef, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetIsLongCharacterMember func(CharacterSetRef, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetIsSupersetOfSet func(CharacterSetRef, CharacterSetRef) unsafe.Pointer
	_CFCharacterSetRemoveCharactersInRange func(MutableCharacterSetRef, Range)
	_CFCharacterSetRemoveCharactersInString func(MutableCharacterSetRef, StringRef)
	_CFCharacterSetUnion func(MutableCharacterSetRef, CharacterSetRef)
	_CFCopyDescription func(TypeRef) StringRef
	_CFCopyHomeDirectoryURL func() URLRef
	_CFCopyTypeIDDescription func(TypeID) StringRef
	_CFDataAppendBytes func(MutableDataRef, unsafe.Pointer, Index)
	_CFDataCreate func(AllocatorRef, unsafe.Pointer, Index) DataRef
	_CFDataCreateCopy func(AllocatorRef, DataRef) DataRef
	_CFDataCreateMutable func(AllocatorRef, Index) MutableDataRef
	_CFDataCreateMutableCopy func(AllocatorRef, Index, DataRef) MutableDataRef
	_CFDataCreateWithBytesNoCopy func(AllocatorRef, unsafe.Pointer, Index, AllocatorRef) DataRef
	_CFDataDeleteBytes func(MutableDataRef, Range)
	_CFDataFind func(DataRef, DataRef, Range, DataSearchFlags) Range
	_CFDataGetBytePtr func(DataRef) unsafe.Pointer
	_CFDataGetBytes func(DataRef, Range, unsafe.Pointer)
	_CFDataGetLength func(DataRef) Index
	_CFDataGetMutableBytePtr func(MutableDataRef) unsafe.Pointer
	_CFDataGetTypeID func() TypeID
	_CFDataIncreaseLength func(MutableDataRef, Index)
	_CFDataReplaceBytes func(MutableDataRef, Range, unsafe.Pointer, Index)
	_CFDataSetLength func(MutableDataRef, Index)
	_CFDateCompare func(DateRef, DateRef, unsafe.Pointer) ComparisonResult
	_CFDateCreate func(AllocatorRef, AbsoluteTime) DateRef
	_CFDateFormatterCopyProperty func(DateFormatterRef, DateFormatterKey) TypeRef
	_CFDateFormatterCreate func(AllocatorRef, LocaleRef, DateFormatterStyle, DateFormatterStyle) DateFormatterRef
	_CFDateFormatterCreateDateFormatFromTemplate func(AllocatorRef, StringRef, OptionFlags, LocaleRef) StringRef
	_CFDateFormatterCreateDateFromString func(AllocatorRef, DateFormatterRef, StringRef, unsafe.Pointer) DateRef
	_CFDateFormatterCreateISO8601Formatter func(AllocatorRef, ISO8601DateFormatOptions) DateFormatterRef
	_CFDateFormatterCreateStringWithAbsoluteTime func(AllocatorRef, DateFormatterRef, AbsoluteTime) StringRef
	_CFDateFormatterCreateStringWithDate func(AllocatorRef, DateFormatterRef, DateRef) StringRef
	_CFDateFormatterGetAbsoluteTimeFromString func(DateFormatterRef, StringRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetDateStyle func(DateFormatterRef) DateFormatterStyle
	_CFDateFormatterGetFormat func(DateFormatterRef) StringRef
	_CFDateFormatterGetLocale func(DateFormatterRef) LocaleRef
	_CFDateFormatterGetTimeStyle func(DateFormatterRef) DateFormatterStyle
	_CFDateFormatterGetTypeID func() TypeID
	_CFDateFormatterSetFormat func(DateFormatterRef, StringRef)
	_CFDateFormatterSetProperty func(DateFormatterRef, StringRef, TypeRef)
	_CFDateGetAbsoluteTime func(DateRef) AbsoluteTime
	_CFDateGetTimeIntervalSinceDate func(DateRef, DateRef) TimeInterval
	_CFDateGetTypeID func() TypeID
	_CFDictionaryAddValue func(MutableDictionaryRef, unsafe.Pointer, unsafe.Pointer)
	_CFDictionaryApplyFunction func(DictionaryRef, DictionaryApplierFunction, unsafe.Pointer)
	_CFDictionaryContainsKey func(DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryContainsValue func(DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryCreate func(AllocatorRef, unsafe.Pointer, unsafe.Pointer, Index, unsafe.Pointer, unsafe.Pointer) DictionaryRef
	_CFDictionaryCreateCopy func(AllocatorRef, DictionaryRef) DictionaryRef
	_CFDictionaryCreateMutable func(AllocatorRef, Index, unsafe.Pointer, unsafe.Pointer) MutableDictionaryRef
	_CFDictionaryCreateMutableCopy func(AllocatorRef, Index, DictionaryRef) MutableDictionaryRef
	_CFDictionaryGetCount func(DictionaryRef) Index
	_CFDictionaryGetCountOfKey func(DictionaryRef, unsafe.Pointer) Index
	_CFDictionaryGetCountOfValue func(DictionaryRef, unsafe.Pointer) Index
	_CFDictionaryGetKeysAndValues func(DictionaryRef, unsafe.Pointer, unsafe.Pointer)
	_CFDictionaryGetTypeID func() TypeID
	_CFDictionaryGetValue func(DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetValueIfPresent func(DictionaryRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryRemoveAllValues func(MutableDictionaryRef)
	_CFDictionaryRemoveValue func(MutableDictionaryRef, unsafe.Pointer)
	_CFDictionaryReplaceValue func(MutableDictionaryRef, unsafe.Pointer, unsafe.Pointer)
	_CFDictionarySetValue func(MutableDictionaryRef, unsafe.Pointer, unsafe.Pointer)
	_CFEqual func(TypeRef, TypeRef) unsafe.Pointer
	_CFErrorCopyDescription func(ErrorRef) StringRef
	_CFErrorCopyFailureReason func(ErrorRef) StringRef
	_CFErrorCopyRecoverySuggestion func(ErrorRef) StringRef
	_CFErrorCopyUserInfo func(ErrorRef) DictionaryRef
	_CFErrorCreate func(AllocatorRef, ErrorDomain, Index, DictionaryRef) ErrorRef
	_CFErrorCreateWithUserInfoKeysAndValues func(AllocatorRef, ErrorDomain, Index, unsafe.Pointer, unsafe.Pointer, Index) ErrorRef
	_CFErrorGetCode func(ErrorRef) Index
	_CFErrorGetDomain func(ErrorRef) ErrorDomain
	_CFErrorGetTypeID func() TypeID
	_CFFileDescriptorCreate func(AllocatorRef, FileDescriptorNativeDescriptor, unsafe.Pointer, FileDescriptorCallBack, unsafe.Pointer) FileDescriptorRef
	_CFFileDescriptorCreateRunLoopSource func(AllocatorRef, FileDescriptorRef, Index) RunLoopSourceRef
	_CFFileDescriptorDisableCallBacks func(FileDescriptorRef, OptionFlags)
	_CFFileDescriptorEnableCallBacks func(FileDescriptorRef, OptionFlags)
	_CFFileDescriptorGetContext func(FileDescriptorRef, unsafe.Pointer)
	_CFFileDescriptorGetNativeDescriptor func(FileDescriptorRef) FileDescriptorNativeDescriptor
	_CFFileDescriptorGetTypeID func() TypeID
	_CFFileDescriptorInvalidate func(FileDescriptorRef)
	_CFFileDescriptorIsValid func(FileDescriptorRef) unsafe.Pointer
	_CFFileSecurityClearProperties func(FileSecurityRef, FileSecurityClearOptions) unsafe.Pointer
	_CFFileSecurityCopyAccessControlList func(FileSecurityRef, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityCopyGroupUUID func(FileSecurityRef, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityCopyOwnerUUID func(FileSecurityRef, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityCreate func(AllocatorRef) FileSecurityRef
	_CFFileSecurityCreateCopy func(AllocatorRef, FileSecurityRef) FileSecurityRef
	_CFFileSecurityGetGroup func(FileSecurityRef, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityGetMode func(FileSecurityRef, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityGetOwner func(FileSecurityRef, unsafe.Pointer) unsafe.Pointer
	_CFFileSecurityGetTypeID func() TypeID
	_CFFileSecuritySetAccessControlList func(FileSecurityRef, unsafe.Pointer) unsafe.Pointer
	_CFFileSecuritySetGroup func(FileSecurityRef, unsafe.Pointer) unsafe.Pointer
	_CFFileSecuritySetGroupUUID func(FileSecurityRef, UUIDRef) unsafe.Pointer
	_CFFileSecuritySetMode func(FileSecurityRef, unsafe.Pointer) unsafe.Pointer
	_CFFileSecuritySetOwner func(FileSecurityRef, unsafe.Pointer) unsafe.Pointer
	_CFFileSecuritySetOwnerUUID func(FileSecurityRef, UUIDRef) unsafe.Pointer
	_CFGetAllocator func(TypeRef) AllocatorRef
	_CFGetRetainCount func(TypeRef) Index
	_CFGetTypeID func(TypeRef) TypeID
	_CFGregorianDateGetAbsoluteTime func(GregorianDate, TimeZoneRef) AbsoluteTime
	_CFGregorianDateIsValid func(GregorianDate, OptionFlags) unsafe.Pointer
	_CFHash func(TypeRef) HashCode
	_CFLocaleCopyAvailableLocaleIdentifiers func() ArrayRef
	_CFLocaleCopyCommonISOCurrencyCodes func() ArrayRef
	_CFLocaleCopyCurrent func() LocaleRef
	_CFLocaleCopyDisplayNameForPropertyValue func(LocaleRef, LocaleKey, StringRef) StringRef
	_CFLocaleCopyISOCountryCodes func() ArrayRef
	_CFLocaleCopyISOCurrencyCodes func() ArrayRef
	_CFLocaleCopyISOLanguageCodes func() ArrayRef
	_CFLocaleCopyPreferredLanguages func() ArrayRef
	_CFLocaleCreate func(AllocatorRef, LocaleIdentifier) LocaleRef
	_CFLocaleCreateCanonicalLanguageIdentifierFromString func(AllocatorRef, StringRef) LocaleIdentifier
	_CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes func(AllocatorRef, unsafe.Pointer, unsafe.Pointer) LocaleIdentifier
	_CFLocaleCreateCanonicalLocaleIdentifierFromString func(AllocatorRef, StringRef) LocaleIdentifier
	_CFLocaleCreateComponentsFromLocaleIdentifier func(AllocatorRef, LocaleIdentifier) DictionaryRef
	_CFLocaleCreateCopy func(AllocatorRef, LocaleRef) LocaleRef
	_CFLocaleCreateLocaleIdentifierFromComponents func(AllocatorRef, DictionaryRef) LocaleIdentifier
	_CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode func(AllocatorRef, uint32) LocaleIdentifier
	_CFLocaleGetIdentifier func(LocaleRef) LocaleIdentifier
	_CFLocaleGetLanguageCharacterDirection func(StringRef) LocaleLanguageDirection
	_CFLocaleGetLanguageLineDirection func(StringRef) LocaleLanguageDirection
	_CFLocaleGetSystem func() LocaleRef
	_CFLocaleGetTypeID func() TypeID
	_CFLocaleGetValue func(LocaleRef, LocaleKey) TypeRef
	_CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier func(LocaleIdentifier) uint32
	_CFMachPortCreate func(AllocatorRef, MachPortCallBack, unsafe.Pointer, unsafe.Pointer) MachPortRef
	_CFMachPortCreateRunLoopSource func(AllocatorRef, MachPortRef, Index) RunLoopSourceRef
	_CFMachPortCreateWithPort func(AllocatorRef, unsafe.Pointer, MachPortCallBack, unsafe.Pointer, unsafe.Pointer) MachPortRef
	_CFMachPortGetContext func(MachPortRef, unsafe.Pointer)
	_CFMachPortGetInvalidationCallBack func(MachPortRef) MachPortInvalidationCallBack
	_CFMachPortGetPort func(MachPortRef) unsafe.Pointer
	_CFMachPortGetTypeID func() TypeID
	_CFMachPortInvalidate func(MachPortRef)
	_CFMachPortIsValid func(MachPortRef) unsafe.Pointer
	_CFMachPortSetInvalidationCallBack func(MachPortRef, MachPortInvalidationCallBack)
	_CFMakeCollectable func(TypeRef) TypeRef
	_CFMessagePortCreateLocal func(AllocatorRef, StringRef, MessagePortCallBack, unsafe.Pointer, unsafe.Pointer) MessagePortRef
	_CFMessagePortCreateRemote func(AllocatorRef, StringRef) MessagePortRef
	_CFMessagePortCreateRunLoopSource func(AllocatorRef, MessagePortRef, Index) RunLoopSourceRef
	_CFMessagePortGetContext func(MessagePortRef, unsafe.Pointer)
	_CFMessagePortGetInvalidationCallBack func(MessagePortRef) MessagePortInvalidationCallBack
	_CFMessagePortGetName func(MessagePortRef) StringRef
	_CFMessagePortGetTypeID func() TypeID
	_CFMessagePortInvalidate func(MessagePortRef)
	_CFMessagePortIsRemote func(MessagePortRef) unsafe.Pointer
	_CFMessagePortIsValid func(MessagePortRef) unsafe.Pointer
	_CFMessagePortSendRequest func(MessagePortRef, unsafe.Pointer, DataRef, TimeInterval, TimeInterval, StringRef, unsafe.Pointer) unsafe.Pointer
	_CFMessagePortSetDispatchQueue func(MessagePortRef, unsafe.Pointer)
	_CFMessagePortSetInvalidationCallBack func(MessagePortRef, MessagePortInvalidationCallBack)
	_CFMessagePortSetName func(MessagePortRef, StringRef) unsafe.Pointer
	_CFNotificationCenterAddObserver func(NotificationCenterRef, unsafe.Pointer, NotificationCallback, StringRef, unsafe.Pointer, NotificationSuspensionBehavior)
	_CFNotificationCenterGetDarwinNotifyCenter func() NotificationCenterRef
	_CFNotificationCenterGetDistributedCenter func() NotificationCenterRef
	_CFNotificationCenterGetLocalCenter func() NotificationCenterRef
	_CFNotificationCenterGetTypeID func() TypeID
	_CFNotificationCenterPostNotification func(NotificationCenterRef, NotificationName, unsafe.Pointer, DictionaryRef, unsafe.Pointer)
	_CFNotificationCenterPostNotificationWithOptions func(NotificationCenterRef, NotificationName, unsafe.Pointer, DictionaryRef, OptionFlags)
	_CFNotificationCenterRemoveEveryObserver func(NotificationCenterRef, unsafe.Pointer)
	_CFNotificationCenterRemoveObserver func(NotificationCenterRef, unsafe.Pointer, NotificationName, unsafe.Pointer)
	_CFNullGetTypeID func() TypeID
	_CFNumberCompare func(NumberRef, NumberRef, unsafe.Pointer) ComparisonResult
	_CFNumberCreate func(AllocatorRef, NumberType, unsafe.Pointer) NumberRef
	_CFNumberFormatterCopyProperty func(NumberFormatterRef, NumberFormatterKey) TypeRef
	_CFNumberFormatterCreate func(AllocatorRef, LocaleRef, NumberFormatterStyle) NumberFormatterRef
	_CFNumberFormatterCreateNumberFromString func(AllocatorRef, NumberFormatterRef, StringRef, unsafe.Pointer, OptionFlags) NumberRef
	_CFNumberFormatterCreateStringWithNumber func(AllocatorRef, NumberFormatterRef, NumberRef) StringRef
	_CFNumberFormatterCreateStringWithValue func(AllocatorRef, NumberFormatterRef, NumberType, unsafe.Pointer) StringRef
	_CFNumberFormatterGetDecimalInfoForCurrencyCode func(StringRef, unsafe.Pointer, []float64) unsafe.Pointer
	_CFNumberFormatterGetFormat func(NumberFormatterRef) StringRef
	_CFNumberFormatterGetLocale func(NumberFormatterRef) LocaleRef
	_CFNumberFormatterGetStyle func(NumberFormatterRef) NumberFormatterStyle
	_CFNumberFormatterGetTypeID func() TypeID
	_CFNumberFormatterGetValueFromString func(NumberFormatterRef, StringRef, unsafe.Pointer, NumberType, unsafe.Pointer) unsafe.Pointer
	_CFNumberFormatterSetFormat func(NumberFormatterRef, StringRef)
	_CFNumberFormatterSetProperty func(NumberFormatterRef, NumberFormatterKey, TypeRef)
	_CFNumberGetByteSize func(NumberRef) Index
	_CFNumberGetType func(NumberRef) NumberType
	_CFNumberGetTypeID func() TypeID
	_CFNumberGetValue func(NumberRef, NumberType, unsafe.Pointer) unsafe.Pointer
	_CFNumberIsFloatType func(NumberRef) unsafe.Pointer
	_CFPlugInAddInstanceForFactory func(UUIDRef)
	_CFPlugInCreate func(AllocatorRef, URLRef) PlugInRef
	_CFPlugInFindFactoriesForPlugInType func(UUIDRef) ArrayRef
	_CFPlugInFindFactoriesForPlugInTypeInPlugIn func(UUIDRef, PlugInRef) ArrayRef
	_CFPlugInGetBundle func(PlugInRef) BundleRef
	_CFPlugInGetTypeID func() TypeID
	_CFPlugInInstanceCreate func(AllocatorRef, UUIDRef, UUIDRef) unsafe.Pointer
	_CFPlugInInstanceCreateWithInstanceDataSize func(AllocatorRef, Index, PlugInInstanceDeallocateInstanceDataFunction, StringRef, PlugInInstanceGetInterfaceFunction) PlugInInstanceRef
	_CFPlugInInstanceGetFactoryName func(PlugInInstanceRef) StringRef
	_CFPlugInInstanceGetInstanceData func(PlugInInstanceRef) unsafe.Pointer
	_CFPlugInInstanceGetInterfaceFunctionTable func(PlugInInstanceRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CFPlugInInstanceGetTypeID func() TypeID
	_CFPlugInIsLoadOnDemand func(PlugInRef) unsafe.Pointer
	_CFPlugInRegisterFactoryFunction func(UUIDRef, PlugInFactoryFunction) unsafe.Pointer
	_CFPlugInRegisterFactoryFunctionByName func(UUIDRef, PlugInRef, StringRef) unsafe.Pointer
	_CFPlugInRegisterPlugInType func(UUIDRef, UUIDRef) unsafe.Pointer
	_CFPlugInRemoveInstanceForFactory func(UUIDRef)
	_CFPlugInSetLoadOnDemand func(PlugInRef, unsafe.Pointer)
	_CFPlugInUnregisterFactory func(UUIDRef) unsafe.Pointer
	_CFPlugInUnregisterPlugInType func(UUIDRef, UUIDRef) unsafe.Pointer
	_CFPreferencesAddSuitePreferencesToApp func(StringRef, StringRef)
	_CFPreferencesAppSynchronize func(StringRef) unsafe.Pointer
	_CFPreferencesAppValueIsForced func(StringRef, StringRef) unsafe.Pointer
	_CFPreferencesCopyAppValue func(StringRef, StringRef) PropertyListRef
	_CFPreferencesCopyApplicationList func(StringRef, StringRef) ArrayRef
	_CFPreferencesCopyKeyList func(StringRef, StringRef, StringRef) ArrayRef
	_CFPreferencesCopyMultiple func(ArrayRef, StringRef, StringRef, StringRef) DictionaryRef
	_CFPreferencesCopyValue func(StringRef, StringRef, StringRef, StringRef) PropertyListRef
	_CFPreferencesGetAppBooleanValue func(StringRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CFPreferencesGetAppIntegerValue func(StringRef, StringRef, unsafe.Pointer) Index
	_CFPreferencesRemoveSuitePreferencesFromApp func(StringRef, StringRef)
	_CFPreferencesSetAppValue func(StringRef, PropertyListRef, StringRef)
	_CFPreferencesSetMultiple func(DictionaryRef, ArrayRef, StringRef, StringRef, StringRef)
	_CFPreferencesSetValue func(StringRef, PropertyListRef, StringRef, StringRef, StringRef)
	_CFPreferencesSynchronize func(StringRef, StringRef, StringRef) unsafe.Pointer
	_CFPropertyListCreateData func(AllocatorRef, PropertyListRef, PropertyListFormat, OptionFlags, unsafe.Pointer) DataRef
	_CFPropertyListCreateDeepCopy func(AllocatorRef, PropertyListRef, OptionFlags) PropertyListRef
	_CFPropertyListCreateFromStream func(AllocatorRef, ReadStreamRef, Index, OptionFlags, unsafe.Pointer, unsafe.Pointer) PropertyListRef
	_CFPropertyListCreateFromXMLData func(AllocatorRef, DataRef, OptionFlags, unsafe.Pointer) PropertyListRef
	_CFPropertyListCreateWithData func(AllocatorRef, DataRef, OptionFlags, unsafe.Pointer, unsafe.Pointer) PropertyListRef
	_CFPropertyListCreateWithStream func(AllocatorRef, ReadStreamRef, Index, OptionFlags, unsafe.Pointer, unsafe.Pointer) PropertyListRef
	_CFPropertyListCreateXMLData func(AllocatorRef, PropertyListRef) DataRef
	_CFPropertyListIsValid func(PropertyListRef, PropertyListFormat) unsafe.Pointer
	_CFPropertyListWrite func(PropertyListRef, WriteStreamRef, PropertyListFormat, OptionFlags, unsafe.Pointer) Index
	_CFPropertyListWriteToStream func(PropertyListRef, WriteStreamRef, PropertyListFormat, unsafe.Pointer) Index
	_CFReadStreamClose func(ReadStreamRef)
	_CFReadStreamCopyDispatchQueue func(ReadStreamRef) unsafe.Pointer
	_CFReadStreamCopyError func(ReadStreamRef) ErrorRef
	_CFReadStreamCopyProperty func(ReadStreamRef, StreamPropertyKey) TypeRef
	_CFReadStreamCreateWithBytesNoCopy func(AllocatorRef, unsafe.Pointer, Index, AllocatorRef) ReadStreamRef
	_CFReadStreamCreateWithFile func(AllocatorRef, URLRef) ReadStreamRef
	_CFReadStreamGetBuffer func(ReadStreamRef, Index, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamGetError func(ReadStreamRef) StreamError
	_CFReadStreamGetStatus func(ReadStreamRef) StreamStatus
	_CFReadStreamGetTypeID func() TypeID
	_CFReadStreamHasBytesAvailable func(ReadStreamRef) unsafe.Pointer
	_CFReadStreamOpen func(ReadStreamRef) unsafe.Pointer
	_CFReadStreamRead func(ReadStreamRef, unsafe.Pointer, Index) Index
	_CFReadStreamScheduleWithRunLoop func(ReadStreamRef, RunLoopRef, RunLoopMode)
	_CFReadStreamSetClient func(ReadStreamRef, OptionFlags, ReadStreamClientCallBack, unsafe.Pointer) unsafe.Pointer
	_CFReadStreamSetDispatchQueue func(ReadStreamRef, unsafe.Pointer)
	_CFReadStreamSetProperty func(ReadStreamRef, StreamPropertyKey, TypeRef) unsafe.Pointer
	_CFReadStreamUnscheduleFromRunLoop func(ReadStreamRef, RunLoopRef, RunLoopMode)
	_CFRelease func(TypeRef)
	_CFRetain func(TypeRef) TypeRef
	_CFRunLoopAddCommonMode func(RunLoopRef, RunLoopMode)
	_CFRunLoopAddObserver func(RunLoopRef, RunLoopObserverRef, RunLoopMode)
	_CFRunLoopAddSource func(RunLoopRef, RunLoopSourceRef, RunLoopMode)
	_CFRunLoopAddTimer func(RunLoopRef, RunLoopTimerRef, RunLoopMode)
	_CFRunLoopContainsObserver func(RunLoopRef, RunLoopObserverRef, RunLoopMode) unsafe.Pointer
	_CFRunLoopContainsSource func(RunLoopRef, RunLoopSourceRef, RunLoopMode) unsafe.Pointer
	_CFRunLoopContainsTimer func(RunLoopRef, RunLoopTimerRef, RunLoopMode) unsafe.Pointer
	_CFRunLoopCopyAllModes func(RunLoopRef) ArrayRef
	_CFRunLoopCopyCurrentMode func(RunLoopRef) RunLoopMode
	_CFRunLoopGetCurrent func() RunLoopRef
	_CFRunLoopGetMain func() RunLoopRef
	_CFRunLoopGetNextTimerFireDate func(RunLoopRef, RunLoopMode) AbsoluteTime
	_CFRunLoopGetTypeID func() TypeID
	_CFRunLoopIsWaiting func(RunLoopRef) unsafe.Pointer
	_CFRunLoopObserverCreate func(AllocatorRef, OptionFlags, unsafe.Pointer, Index, RunLoopObserverCallBack, unsafe.Pointer) RunLoopObserverRef
	_CFRunLoopObserverCreateWithHandler func(AllocatorRef, OptionFlags, unsafe.Pointer, Index) RunLoopObserverRef
	_CFRunLoopObserverDoesRepeat func(RunLoopObserverRef) unsafe.Pointer
	_CFRunLoopObserverGetActivities func(RunLoopObserverRef) OptionFlags
	_CFRunLoopObserverGetContext func(RunLoopObserverRef, unsafe.Pointer)
	_CFRunLoopObserverGetOrder func(RunLoopObserverRef) Index
	_CFRunLoopObserverGetTypeID func() TypeID
	_CFRunLoopObserverInvalidate func(RunLoopObserverRef)
	_CFRunLoopObserverIsValid func(RunLoopObserverRef) unsafe.Pointer
	_CFRunLoopPerformBlock func(RunLoopRef, TypeRef)
	_CFRunLoopRemoveObserver func(RunLoopRef, RunLoopObserverRef, RunLoopMode)
	_CFRunLoopRemoveSource func(RunLoopRef, RunLoopSourceRef, RunLoopMode)
	_CFRunLoopRemoveTimer func(RunLoopRef, RunLoopTimerRef, RunLoopMode)
	_CFRunLoopRun func()
	_CFRunLoopRunInMode func(RunLoopMode, TimeInterval, unsafe.Pointer) RunLoopRunResult
	_CFRunLoopSourceCreate func(AllocatorRef, Index, unsafe.Pointer) RunLoopSourceRef
	_CFRunLoopSourceGetContext func(RunLoopSourceRef, unsafe.Pointer)
	_CFRunLoopSourceGetOrder func(RunLoopSourceRef) Index
	_CFRunLoopSourceGetTypeID func() TypeID
	_CFRunLoopSourceInvalidate func(RunLoopSourceRef)
	_CFRunLoopSourceIsValid func(RunLoopSourceRef) unsafe.Pointer
	_CFRunLoopSourceSignal func(RunLoopSourceRef)
	_CFRunLoopStop func(RunLoopRef)
	_CFRunLoopTimerCreate func(AllocatorRef, AbsoluteTime, TimeInterval, OptionFlags, Index, RunLoopTimerCallBack, unsafe.Pointer) RunLoopTimerRef
	_CFRunLoopTimerCreateWithHandler func(AllocatorRef, AbsoluteTime, TimeInterval, OptionFlags, Index) RunLoopTimerRef
	_CFRunLoopTimerDoesRepeat func(RunLoopTimerRef) unsafe.Pointer
	_CFRunLoopTimerGetContext func(RunLoopTimerRef, unsafe.Pointer)
	_CFRunLoopTimerGetInterval func(RunLoopTimerRef) TimeInterval
	_CFRunLoopTimerGetNextFireDate func(RunLoopTimerRef) AbsoluteTime
	_CFRunLoopTimerGetOrder func(RunLoopTimerRef) Index
	_CFRunLoopTimerGetTolerance func(RunLoopTimerRef) TimeInterval
	_CFRunLoopTimerGetTypeID func() TypeID
	_CFRunLoopTimerInvalidate func(RunLoopTimerRef)
	_CFRunLoopTimerIsValid func(RunLoopTimerRef) unsafe.Pointer
	_CFRunLoopTimerSetNextFireDate func(RunLoopTimerRef, AbsoluteTime)
	_CFRunLoopTimerSetTolerance func(RunLoopTimerRef, TimeInterval)
	_CFRunLoopWakeUp func(RunLoopRef)
	_CFSetAddValue func(MutableSetRef, unsafe.Pointer)
	_CFSetApplyFunction func(SetRef, SetApplierFunction, unsafe.Pointer)
	_CFSetContainsValue func(SetRef, unsafe.Pointer) unsafe.Pointer
	_CFSetCreate func(AllocatorRef, unsafe.Pointer, Index, unsafe.Pointer) SetRef
	_CFSetCreateCopy func(AllocatorRef, SetRef) SetRef
	_CFSetCreateMutable func(AllocatorRef, Index, unsafe.Pointer) MutableSetRef
	_CFSetCreateMutableCopy func(AllocatorRef, Index, SetRef) MutableSetRef
	_CFSetGetCount func(SetRef) Index
	_CFSetGetCountOfValue func(SetRef, unsafe.Pointer) Index
	_CFSetGetTypeID func() TypeID
	_CFSetGetValue func(SetRef, unsafe.Pointer) unsafe.Pointer
	_CFSetGetValueIfPresent func(SetRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFSetGetValues func(SetRef, unsafe.Pointer)
	_CFSetRemoveAllValues func(MutableSetRef)
	_CFSetRemoveValue func(MutableSetRef, unsafe.Pointer)
	_CFSetReplaceValue func(MutableSetRef, unsafe.Pointer)
	_CFSetSetValue func(MutableSetRef, unsafe.Pointer)
	_CFShow func(TypeRef)
	_CFShowStr func(StringRef)
	_CFSocketConnectToAddress func(SocketRef, DataRef, TimeInterval) SocketError
	_CFSocketCopyAddress func(SocketRef) DataRef
	_CFSocketCopyPeerAddress func(SocketRef) DataRef
	_CFSocketCopyRegisteredSocketSignature func(unsafe.Pointer, TimeInterval, StringRef, unsafe.Pointer, unsafe.Pointer) SocketError
	_CFSocketCopyRegisteredValue func(unsafe.Pointer, TimeInterval, StringRef, unsafe.Pointer, unsafe.Pointer) SocketError
	_CFSocketCreate func(AllocatorRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, OptionFlags, SocketCallBack, unsafe.Pointer) SocketRef
	_CFSocketCreateConnectedToSocketSignature func(AllocatorRef, unsafe.Pointer, OptionFlags, SocketCallBack, unsafe.Pointer, TimeInterval) SocketRef
	_CFSocketCreateRunLoopSource func(AllocatorRef, SocketRef, Index) RunLoopSourceRef
	_CFSocketCreateWithNative func(AllocatorRef, SocketNativeHandle, OptionFlags, SocketCallBack, unsafe.Pointer) SocketRef
	_CFSocketCreateWithSocketSignature func(AllocatorRef, unsafe.Pointer, OptionFlags, SocketCallBack, unsafe.Pointer) SocketRef
	_CFSocketDisableCallBacks func(SocketRef, OptionFlags)
	_CFSocketEnableCallBacks func(SocketRef, OptionFlags)
	_CFSocketGetContext func(SocketRef, unsafe.Pointer)
	_CFSocketGetDefaultNameRegistryPortNumber func() unsafe.Pointer
	_CFSocketGetNative func(SocketRef) SocketNativeHandle
	_CFSocketGetSocketFlags func(SocketRef) OptionFlags
	_CFSocketGetTypeID func() TypeID
	_CFSocketInvalidate func(SocketRef)
	_CFSocketIsValid func(SocketRef) unsafe.Pointer
	_CFSocketRegisterSocketSignature func(unsafe.Pointer, TimeInterval, StringRef, unsafe.Pointer) SocketError
	_CFSocketRegisterValue func(unsafe.Pointer, TimeInterval, StringRef, PropertyListRef) SocketError
	_CFSocketSendData func(SocketRef, DataRef, DataRef, TimeInterval) SocketError
	_CFSocketSetAddress func(SocketRef, DataRef) SocketError
	_CFSocketSetDefaultNameRegistryPortNumber func(unsafe.Pointer)
	_CFSocketSetSocketFlags func(SocketRef, OptionFlags)
	_CFSocketUnregister func(unsafe.Pointer, TimeInterval, StringRef) SocketError
	_CFStreamCreateBoundPair func(AllocatorRef, unsafe.Pointer, unsafe.Pointer, Index)
	_CFStreamCreatePairWithPeerSocketSignature func(AllocatorRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFStreamCreatePairWithSocket func(AllocatorRef, SocketNativeHandle, unsafe.Pointer, unsafe.Pointer)
	_CFStreamCreatePairWithSocketToHost func(AllocatorRef, StringRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFStringAppend func(MutableStringRef, StringRef)
	_CFStringAppendCString func(MutableStringRef, unsafe.Pointer, StringEncoding)
	_CFStringAppendCharacters func(MutableStringRef, unsafe.Pointer, Index)
	_CFStringAppendFormat func(MutableStringRef, DictionaryRef, StringRef)
	_CFStringAppendFormatAndArguments func(MutableStringRef, DictionaryRef, StringRef, unsafe.Pointer)
	_CFStringAppendPascalString func(MutableStringRef, unsafe.Pointer, StringEncoding)
	_CFStringCapitalize func(MutableStringRef, LocaleRef)
	_CFStringCompare func(StringRef, StringRef, StringCompareFlags) ComparisonResult
	_CFStringCompareWithOptions func(StringRef, StringRef, Range, StringCompareFlags) ComparisonResult
	_CFStringCompareWithOptionsAndLocale func(StringRef, StringRef, Range, StringCompareFlags, LocaleRef) ComparisonResult
	_CFStringConvertEncodingToIANACharSetName func(StringEncoding) StringRef
	_CFStringConvertEncodingToNSStringEncoding func(StringEncoding) unsafe.Pointer
	_CFStringConvertEncodingToWindowsCodepage func(StringEncoding) unsafe.Pointer
	_CFStringConvertIANACharSetNameToEncoding func(StringRef) StringEncoding
	_CFStringConvertNSStringEncodingToEncoding func(unsafe.Pointer) StringEncoding
	_CFStringConvertWindowsCodepageToEncoding func(unsafe.Pointer) StringEncoding
	_CFStringCreateArrayBySeparatingStrings func(AllocatorRef, StringRef, StringRef) ArrayRef
	_CFStringCreateArrayWithFindResults func(AllocatorRef, StringRef, StringRef, Range, StringCompareFlags) ArrayRef
	_CFStringCreateByCombiningStrings func(AllocatorRef, ArrayRef, StringRef) StringRef
	_CFStringCreateCopy func(AllocatorRef, StringRef) StringRef
	_CFStringCreateExternalRepresentation func(AllocatorRef, StringRef, StringEncoding, unsafe.Pointer) DataRef
	_CFStringCreateFromExternalRepresentation func(AllocatorRef, DataRef, StringEncoding) StringRef
	_CFStringCreateMutable func(AllocatorRef, Index) MutableStringRef
	_CFStringCreateMutableCopy func(AllocatorRef, Index, StringRef) MutableStringRef
	_CFStringCreateMutableWithExternalCharactersNoCopy func(AllocatorRef, unsafe.Pointer, Index, Index, AllocatorRef) MutableStringRef
	_CFStringCreateStringWithValidatedFormat func(AllocatorRef, DictionaryRef, StringRef, StringRef, unsafe.Pointer) StringRef
	_CFStringCreateStringWithValidatedFormatAndArguments func(AllocatorRef, DictionaryRef, StringRef, StringRef, unsafe.Pointer, unsafe.Pointer) StringRef
	_CFStringCreateWithBytes func(AllocatorRef, unsafe.Pointer, Index, StringEncoding, unsafe.Pointer) StringRef
	_CFStringCreateWithBytesNoCopy func(AllocatorRef, unsafe.Pointer, Index, StringEncoding, unsafe.Pointer, AllocatorRef) StringRef
	_CFStringCreateWithCString func(AllocatorRef, unsafe.Pointer, StringEncoding) StringRef
	_CFStringCreateWithCStringNoCopy func(AllocatorRef, unsafe.Pointer, StringEncoding, AllocatorRef) StringRef
	_CFStringCreateWithCharacters func(AllocatorRef, unsafe.Pointer, Index) StringRef
	_CFStringCreateWithCharactersNoCopy func(AllocatorRef, unsafe.Pointer, Index, AllocatorRef) StringRef
	_CFStringCreateWithFileSystemRepresentation func(AllocatorRef, unsafe.Pointer) StringRef
	_CFStringCreateWithFormat func(AllocatorRef, DictionaryRef, StringRef) StringRef
	_CFStringCreateWithFormatAndArguments func(AllocatorRef, DictionaryRef, StringRef, unsafe.Pointer) StringRef
	_CFStringCreateWithPascalString func(AllocatorRef, unsafe.Pointer, StringEncoding) StringRef
	_CFStringCreateWithPascalStringNoCopy func(AllocatorRef, unsafe.Pointer, StringEncoding, AllocatorRef) StringRef
	_CFStringCreateWithSubstring func(AllocatorRef, StringRef, Range) StringRef
	_CFStringDelete func(MutableStringRef, Range)
	_CFStringFind func(StringRef, StringRef, StringCompareFlags) Range
	_CFStringFindAndReplace func(MutableStringRef, StringRef, StringRef, Range, StringCompareFlags) Index
	_CFStringFindCharacterFromSet func(StringRef, CharacterSetRef, Range, StringCompareFlags, unsafe.Pointer) unsafe.Pointer
	_CFStringFindWithOptions func(StringRef, StringRef, Range, StringCompareFlags, unsafe.Pointer) unsafe.Pointer
	_CFStringFindWithOptionsAndLocale func(StringRef, StringRef, Range, StringCompareFlags, LocaleRef, unsafe.Pointer) unsafe.Pointer
	_CFStringFold func(MutableStringRef, StringCompareFlags, LocaleRef)
	_CFStringGetBytes func(StringRef, Range, StringEncoding, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, Index, unsafe.Pointer) Index
	_CFStringGetCString func(StringRef, unsafe.Pointer, Index, StringEncoding) unsafe.Pointer
	_CFStringGetCStringPtr func(StringRef, StringEncoding) unsafe.Pointer
	_CFStringGetCharacterAtIndex func(StringRef, Index) unsafe.Pointer
	_CFStringGetCharacters func(StringRef, Range, unsafe.Pointer)
	_CFStringGetCharactersPtr func(StringRef) unsafe.Pointer
	_CFStringGetDoubleValue func(StringRef) float64
	_CFStringGetFastestEncoding func(StringRef) StringEncoding
	_CFStringGetFileSystemRepresentation func(StringRef, unsafe.Pointer, Index) unsafe.Pointer
	_CFStringGetHyphenationLocationBeforeIndex func(StringRef, Index, Range, OptionFlags, LocaleRef, unsafe.Pointer) Index
	_CFStringGetIntValue func(StringRef) unsafe.Pointer
	_CFStringGetLength func(StringRef) Index
	_CFStringGetLineBounds func(StringRef, Range, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFStringGetListOfAvailableEncodings func() unsafe.Pointer
	_CFStringGetMaximumSizeForEncoding func(Index, StringEncoding) Index
	_CFStringGetMaximumSizeOfFileSystemRepresentation func(StringRef) Index
	_CFStringGetMostCompatibleMacStringEncoding func(StringEncoding) StringEncoding
	_CFStringGetNameOfEncoding func(StringEncoding) StringRef
	_CFStringGetParagraphBounds func(StringRef, Range, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFStringGetPascalString func(StringRef, unsafe.Pointer, Index, StringEncoding) unsafe.Pointer
	_CFStringGetPascalStringPtr func(StringRef, StringEncoding) unsafe.Pointer
	_CFStringGetRangeOfComposedCharactersAtIndex func(StringRef, Index) Range
	_CFStringGetSmallestEncoding func(StringRef) StringEncoding
	_CFStringGetSystemEncoding func() StringEncoding
	_CFStringGetTypeID func() TypeID
	_CFStringHasPrefix func(StringRef, StringRef) unsafe.Pointer
	_CFStringHasSuffix func(StringRef, StringRef) unsafe.Pointer
	_CFStringInsert func(MutableStringRef, Index, StringRef)
	_CFStringIsEncodingAvailable func(StringEncoding) unsafe.Pointer
	_CFStringIsHyphenationAvailableForLocale func(LocaleRef) unsafe.Pointer
	_CFStringLowercase func(MutableStringRef, LocaleRef)
	_CFStringNormalize func(MutableStringRef, StringNormalizationForm)
	_CFStringPad func(MutableStringRef, StringRef, Index, Index)
	_CFStringReplace func(MutableStringRef, Range, StringRef)
	_CFStringReplaceAll func(MutableStringRef, StringRef)
	_CFStringSetExternalCharactersNoCopy func(MutableStringRef, unsafe.Pointer, Index, Index)
	_CFStringTokenizerAdvanceToNextToken func(StringTokenizerRef) StringTokenizerTokenType
	_CFStringTokenizerCopyBestStringLanguage func(StringRef, Range) StringRef
	_CFStringTokenizerCopyCurrentTokenAttribute func(StringTokenizerRef, OptionFlags) TypeRef
	_CFStringTokenizerCreate func(AllocatorRef, StringRef, Range, OptionFlags, LocaleRef) StringTokenizerRef
	_CFStringTokenizerGetCurrentSubTokens func(StringTokenizerRef, unsafe.Pointer, Index, MutableArrayRef) Index
	_CFStringTokenizerGetCurrentTokenRange func(StringTokenizerRef) Range
	_CFStringTokenizerGetTypeID func() TypeID
	_CFStringTokenizerGoToTokenAtIndex func(StringTokenizerRef, Index) StringTokenizerTokenType
	_CFStringTokenizerSetString func(StringTokenizerRef, StringRef, Range)
	_CFStringTransform func(MutableStringRef, unsafe.Pointer, StringRef, unsafe.Pointer) unsafe.Pointer
	_CFStringTrim func(MutableStringRef, StringRef)
	_CFStringTrimWhitespace func(MutableStringRef)
	_CFStringUppercase func(MutableStringRef, LocaleRef)
	_CFTimeZoneCopyAbbreviation func(TimeZoneRef, AbsoluteTime) StringRef
	_CFTimeZoneCopyAbbreviationDictionary func() DictionaryRef
	_CFTimeZoneCopyDefault func() TimeZoneRef
	_CFTimeZoneCopyKnownNames func() ArrayRef
	_CFTimeZoneCopyLocalizedName func(TimeZoneRef, TimeZoneNameStyle, LocaleRef) StringRef
	_CFTimeZoneCopySystem func() TimeZoneRef
	_CFTimeZoneCreate func(AllocatorRef, StringRef, DataRef) TimeZoneRef
	_CFTimeZoneCreateWithName func(AllocatorRef, StringRef, unsafe.Pointer) TimeZoneRef
	_CFTimeZoneCreateWithTimeIntervalFromGMT func(AllocatorRef, TimeInterval) TimeZoneRef
	_CFTimeZoneGetData func(TimeZoneRef) DataRef
	_CFTimeZoneGetDaylightSavingTimeOffset func(TimeZoneRef, AbsoluteTime) TimeInterval
	_CFTimeZoneGetName func(TimeZoneRef) StringRef
	_CFTimeZoneGetNextDaylightSavingTimeTransition func(TimeZoneRef, AbsoluteTime) AbsoluteTime
	_CFTimeZoneGetSecondsFromGMT func(TimeZoneRef, AbsoluteTime) TimeInterval
	_CFTimeZoneGetTypeID func() TypeID
	_CFTimeZoneIsDaylightSavingTime func(TimeZoneRef, AbsoluteTime) unsafe.Pointer
	_CFTimeZoneResetSystem func()
	_CFTimeZoneSetAbbreviationDictionary func(DictionaryRef)
	_CFTimeZoneSetDefault func(TimeZoneRef)
	_CFTreeAppendChild func(TreeRef, TreeRef)
	_CFTreeApplyFunctionToChildren func(TreeRef, TreeApplierFunction, unsafe.Pointer)
	_CFTreeCreate func(AllocatorRef, unsafe.Pointer) TreeRef
	_CFTreeFindRoot func(TreeRef) TreeRef
	_CFTreeGetChildAtIndex func(TreeRef, Index) TreeRef
	_CFTreeGetChildCount func(TreeRef) Index
	_CFTreeGetChildren func(TreeRef, unsafe.Pointer)
	_CFTreeGetContext func(TreeRef, unsafe.Pointer)
	_CFTreeGetFirstChild func(TreeRef) TreeRef
	_CFTreeGetNextSibling func(TreeRef) TreeRef
	_CFTreeGetParent func(TreeRef) TreeRef
	_CFTreeGetTypeID func() TypeID
	_CFTreeInsertSibling func(TreeRef, TreeRef)
	_CFTreePrependChild func(TreeRef, TreeRef)
	_CFTreeRemove func(TreeRef)
	_CFTreeRemoveAllChildren func(TreeRef)
	_CFTreeSetContext func(TreeRef, unsafe.Pointer)
	_CFTreeSortChildren func(TreeRef, ComparatorFunction, unsafe.Pointer)
	_CFURLCanBeDecomposed func(URLRef) unsafe.Pointer
	_CFURLClearResourcePropertyCache func(URLRef)
	_CFURLClearResourcePropertyCacheForKey func(URLRef, StringRef)
	_CFURLCopyAbsoluteURL func(URLRef) URLRef
	_CFURLCopyFileSystemPath func(URLRef, URLPathStyle) StringRef
	_CFURLCopyFragment func(URLRef, StringRef) StringRef
	_CFURLCopyHostName func(URLRef) StringRef
	_CFURLCopyLastPathComponent func(URLRef) StringRef
	_CFURLCopyNetLocation func(URLRef) StringRef
	_CFURLCopyParameterString func(URLRef, StringRef) StringRef
	_CFURLCopyPassword func(URLRef) StringRef
	_CFURLCopyPath func(URLRef) StringRef
	_CFURLCopyPathExtension func(URLRef) StringRef
	_CFURLCopyQueryString func(URLRef, StringRef) StringRef
	_CFURLCopyResourcePropertiesForKeys func(URLRef, ArrayRef, unsafe.Pointer) DictionaryRef
	_CFURLCopyResourcePropertyForKey func(URLRef, StringRef, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCopyResourceSpecifier func(URLRef) StringRef
	_CFURLCopyScheme func(URLRef) StringRef
	_CFURLCopyStrictPath func(URLRef, unsafe.Pointer) StringRef
	_CFURLCopyUserName func(URLRef) StringRef
	_CFURLCreateAbsoluteURLWithBytes func(AllocatorRef, unsafe.Pointer, Index, StringEncoding, URLRef, unsafe.Pointer) URLRef
	_CFURLCreateBookmarkData func(AllocatorRef, URLRef, URLBookmarkCreationOptions, ArrayRef, URLRef, unsafe.Pointer) DataRef
	_CFURLCreateBookmarkDataFromAliasRecord func(AllocatorRef, DataRef) DataRef
	_CFURLCreateBookmarkDataFromFile func(AllocatorRef, URLRef, unsafe.Pointer) DataRef
	_CFURLCreateByResolvingBookmarkData func(AllocatorRef, DataRef, URLBookmarkResolutionOptions, URLRef, ArrayRef, unsafe.Pointer, unsafe.Pointer) URLRef
	_CFURLCreateCopyAppendingPathComponent func(AllocatorRef, URLRef, StringRef, unsafe.Pointer) URLRef
	_CFURLCreateCopyAppendingPathExtension func(AllocatorRef, URLRef, StringRef) URLRef
	_CFURLCreateCopyDeletingLastPathComponent func(AllocatorRef, URLRef) URLRef
	_CFURLCreateCopyDeletingPathExtension func(AllocatorRef, URLRef) URLRef
	_CFURLCreateData func(AllocatorRef, URLRef, StringEncoding, unsafe.Pointer) DataRef
	_CFURLCreateDataAndPropertiesFromResource func(AllocatorRef, URLRef, unsafe.Pointer, unsafe.Pointer, ArrayRef, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateFilePathURL func(AllocatorRef, URLRef, unsafe.Pointer) URLRef
	_CFURLCreateFileReferenceURL func(AllocatorRef, URLRef, unsafe.Pointer) URLRef
	_CFURLCreateFromFSRef func(AllocatorRef, unsafe.Pointer) URLRef
	_CFURLCreateFromFileSystemRepresentation func(AllocatorRef, unsafe.Pointer, Index, unsafe.Pointer) URLRef
	_CFURLCreateFromFileSystemRepresentationRelativeToBase func(AllocatorRef, unsafe.Pointer, Index, unsafe.Pointer, URLRef) URLRef
	_CFURLCreatePropertyFromResource func(AllocatorRef, URLRef, StringRef, unsafe.Pointer) TypeRef
	_CFURLCreateResourcePropertiesForKeysFromBookmarkData func(AllocatorRef, ArrayRef, DataRef) DictionaryRef
	_CFURLCreateResourcePropertyForKeyFromBookmarkData func(AllocatorRef, StringRef, DataRef) TypeRef
	_CFURLCreateStringByAddingPercentEscapes func(AllocatorRef, StringRef, StringRef, StringRef, StringEncoding) StringRef
	_CFURLCreateStringByReplacingPercentEscapes func(AllocatorRef, StringRef, StringRef) StringRef
	_CFURLCreateStringByReplacingPercentEscapesUsingEncoding func(AllocatorRef, StringRef, StringRef, StringEncoding) StringRef
	_CFURLCreateWithBytes func(AllocatorRef, unsafe.Pointer, Index, StringEncoding, URLRef) URLRef
	_CFURLCreateWithFileSystemPath func(AllocatorRef, StringRef, URLPathStyle, unsafe.Pointer) URLRef
	_CFURLCreateWithFileSystemPathRelativeToBase func(AllocatorRef, StringRef, URLPathStyle, unsafe.Pointer, URLRef) URLRef
	_CFURLCreateWithString func(AllocatorRef, StringRef, URLRef) URLRef
	_CFURLDestroyResource func(URLRef, unsafe.Pointer) unsafe.Pointer
	_CFURLEnumeratorCreateForDirectoryURL func(AllocatorRef, URLRef, URLEnumeratorOptions, ArrayRef) URLEnumeratorRef
	_CFURLEnumeratorCreateForMountedVolumes func(AllocatorRef, URLEnumeratorOptions, ArrayRef) URLEnumeratorRef
	_CFURLEnumeratorGetDescendentLevel func(URLEnumeratorRef) Index
	_CFURLEnumeratorGetNextURL func(URLEnumeratorRef, unsafe.Pointer, unsafe.Pointer) URLEnumeratorResult
	_CFURLEnumeratorGetSourceDidChange func(URLEnumeratorRef) unsafe.Pointer
	_CFURLEnumeratorGetTypeID func() TypeID
	_CFURLEnumeratorSkipDescendents func(URLEnumeratorRef)
	_CFURLGetBaseURL func(URLRef) URLRef
	_CFURLGetByteRangeForComponent func(URLRef, URLComponentType, unsafe.Pointer) Range
	_CFURLGetBytes func(URLRef, unsafe.Pointer, Index) Index
	_CFURLGetFSRef func(URLRef, unsafe.Pointer) unsafe.Pointer
	_CFURLGetFileSystemRepresentation func(URLRef, unsafe.Pointer, unsafe.Pointer, Index) unsafe.Pointer
	_CFURLGetPortNumber func(URLRef) unsafe.Pointer
	_CFURLGetString func(URLRef) StringRef
	_CFURLGetTypeID func() TypeID
	_CFURLHasDirectoryPath func(URLRef) unsafe.Pointer
	_CFURLIsFileReferenceURL func(URLRef) unsafe.Pointer
	_CFURLResourceIsReachable func(URLRef, unsafe.Pointer) unsafe.Pointer
	_CFURLSetResourcePropertiesForKeys func(URLRef, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CFURLSetResourcePropertyForKey func(URLRef, StringRef, TypeRef, unsafe.Pointer) unsafe.Pointer
	_CFURLSetTemporaryResourcePropertyForKey func(URLRef, StringRef, TypeRef)
	_CFURLStartAccessingSecurityScopedResource func(URLRef) unsafe.Pointer
	_CFURLStopAccessingSecurityScopedResource func(URLRef)
	_CFURLWriteBookmarkDataToFile func(DataRef, URLRef, URLBookmarkFileCreationOptions, unsafe.Pointer) unsafe.Pointer
	_CFURLWriteDataAndPropertiesToResource func(URLRef, DataRef, DictionaryRef, unsafe.Pointer) unsafe.Pointer
	_CFUUIDCreate func(AllocatorRef) UUIDRef
	_CFUUIDCreateFromString func(AllocatorRef, StringRef) UUIDRef
	_CFUUIDCreateFromUUIDBytes func(AllocatorRef, UUIDBytes) UUIDRef
	_CFUUIDCreateString func(AllocatorRef, UUIDRef) StringRef
	_CFUUIDCreateWithBytes func(AllocatorRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) UUIDRef
	_CFUUIDGetConstantUUIDWithBytes func(AllocatorRef, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) UUIDRef
	_CFUUIDGetTypeID func() TypeID
	_CFUUIDGetUUIDBytes func(UUIDRef) UUIDBytes
	_CFUserNotificationCancel func(UserNotificationRef) unsafe.Pointer
	_CFUserNotificationCreate func(AllocatorRef, TimeInterval, OptionFlags, unsafe.Pointer, DictionaryRef) UserNotificationRef
	_CFUserNotificationCreateRunLoopSource func(AllocatorRef, UserNotificationRef, UserNotificationCallBack, Index) RunLoopSourceRef
	_CFUserNotificationDisplayAlert func(TimeInterval, OptionFlags, URLRef, URLRef, URLRef, StringRef, StringRef, StringRef, StringRef, StringRef, unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationDisplayNotice func(TimeInterval, OptionFlags, URLRef, URLRef, URLRef, StringRef, StringRef, StringRef) unsafe.Pointer
	_CFUserNotificationGetResponseDictionary func(UserNotificationRef) DictionaryRef
	_CFUserNotificationGetResponseValue func(UserNotificationRef, StringRef, Index) StringRef
	_CFUserNotificationGetTypeID func() TypeID
	_CFUserNotificationReceiveResponse func(UserNotificationRef, TimeInterval, unsafe.Pointer) unsafe.Pointer
	_CFUserNotificationUpdate func(UserNotificationRef, TimeInterval, OptionFlags, DictionaryRef) unsafe.Pointer
	_CFWriteStreamCanAcceptBytes func(WriteStreamRef) unsafe.Pointer
	_CFWriteStreamClose func(WriteStreamRef)
	_CFWriteStreamCopyDispatchQueue func(WriteStreamRef) unsafe.Pointer
	_CFWriteStreamCopyError func(WriteStreamRef) ErrorRef
	_CFWriteStreamCopyProperty func(WriteStreamRef, StreamPropertyKey) TypeRef
	_CFWriteStreamCreateWithAllocatedBuffers func(AllocatorRef, AllocatorRef) WriteStreamRef
	_CFWriteStreamCreateWithBuffer func(AllocatorRef, unsafe.Pointer, Index) WriteStreamRef
	_CFWriteStreamCreateWithFile func(AllocatorRef, URLRef) WriteStreamRef
	_CFWriteStreamGetError func(WriteStreamRef) StreamError
	_CFWriteStreamGetStatus func(WriteStreamRef) StreamStatus
	_CFWriteStreamGetTypeID func() TypeID
	_CFWriteStreamOpen func(WriteStreamRef) unsafe.Pointer
	_CFWriteStreamScheduleWithRunLoop func(WriteStreamRef, RunLoopRef, RunLoopMode)
	_CFWriteStreamSetClient func(WriteStreamRef, OptionFlags, WriteStreamClientCallBack, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamSetDispatchQueue func(WriteStreamRef, unsafe.Pointer)
	_CFWriteStreamSetProperty func(WriteStreamRef, StreamPropertyKey, TypeRef) unsafe.Pointer
	_CFWriteStreamUnscheduleFromRunLoop func(WriteStreamRef, RunLoopRef, RunLoopMode)
	_CFWriteStreamWrite func(WriteStreamRef, unsafe.Pointer, Index) Index
	_CFXMLCreateStringByEscapingEntities func(AllocatorRef, StringRef, DictionaryRef) StringRef
	_CFXMLCreateStringByUnescapingEntities func(AllocatorRef, StringRef, DictionaryRef) StringRef
	_CFXMLNodeCreate func(AllocatorRef, XMLNodeTypeCode, StringRef, unsafe.Pointer, Index) XMLNodeRef
	_CFXMLNodeCreateCopy func(AllocatorRef, XMLNodeRef) XMLNodeRef
	_CFXMLNodeGetInfoPtr func(XMLNodeRef) unsafe.Pointer
	_CFXMLNodeGetString func(XMLNodeRef) StringRef
	_CFXMLNodeGetTypeCode func(XMLNodeRef) XMLNodeTypeCode
	_CFXMLNodeGetTypeID func() TypeID
	_CFXMLNodeGetVersion func(XMLNodeRef) Index
	_CFXMLParserAbort func(XMLParserRef, XMLParserStatusCode, StringRef)
	_CFXMLParserCopyErrorDescription func(XMLParserRef) StringRef
	_CFXMLParserCreate func(AllocatorRef, DataRef, URLRef, OptionFlags, Index, unsafe.Pointer, unsafe.Pointer) XMLParserRef
	_CFXMLParserCreateWithDataFromURL func(AllocatorRef, URLRef, OptionFlags, Index, unsafe.Pointer, unsafe.Pointer) XMLParserRef
	_CFXMLParserGetCallBacks func(XMLParserRef, unsafe.Pointer)
	_CFXMLParserGetContext func(XMLParserRef, unsafe.Pointer)
	_CFXMLParserGetDocument func(XMLParserRef) unsafe.Pointer
	_CFXMLParserGetLineNumber func(XMLParserRef) Index
	_CFXMLParserGetLocation func(XMLParserRef) Index
	_CFXMLParserGetSourceURL func(XMLParserRef) URLRef
	_CFXMLParserGetStatusCode func(XMLParserRef) XMLParserStatusCode
	_CFXMLParserGetTypeID func() TypeID
	_CFXMLParserParse func(XMLParserRef) unsafe.Pointer
	_CFXMLTreeCreateFromData func(AllocatorRef, DataRef, URLRef, OptionFlags, Index) XMLTreeRef
	_CFXMLTreeCreateFromDataWithError func(AllocatorRef, DataRef, URLRef, OptionFlags, Index, unsafe.Pointer) XMLTreeRef
	_CFXMLTreeCreateWithDataFromURL func(AllocatorRef, URLRef, OptionFlags, Index) XMLTreeRef
	_CFXMLTreeCreateWithNode func(AllocatorRef, XMLNodeRef) XMLTreeRef
	_CFXMLTreeCreateXMLData func(AllocatorRef, XMLTreeRef) DataRef
	_CFXMLTreeGetNode func(XMLTreeRef) XMLNodeRef
	_inset func(unsafe.Pointer,   UIEdgeInsets) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
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
	tryRegister(&_CFMakeCollectable, lib, "CFMakeCollectable")
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
	tryRegister(&_CFRelease, lib, "CFRelease")
	tryRegister(&_CFRetain, lib, "CFRetain")
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



// Adds a time interval, expressed as Gregorian units, to a given absolute time.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
// Adds a time interval, expressed as Gregorian units, to a given absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeAddGregorianUnits(_:_:_:)
func CFAbsoluteTimeAddGregorianUnits(at AbsoluteTime, tz TimeZoneRef, units GregorianUnits) AbsoluteTime {
	return _CFAbsoluteTimeAddGregorianUnits(at, tz, units)
}

// Returns the current system absolute time.
//
// Added in macOS .
// Returns the current system absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetCurrent()
func CFAbsoluteTimeGetCurrent() AbsoluteTime {
	return _CFAbsoluteTimeGetCurrent()
}

// Returns an integer representing the day of the week indicated by the specified absolute time.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
// Returns an integer representing the day of the week indicated by the specified absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetDayOfWeek(_:_:)
func CFAbsoluteTimeGetDayOfWeek(at AbsoluteTime, tz TimeZoneRef) unsafe.Pointer {
	return _CFAbsoluteTimeGetDayOfWeek(at, tz)
}

// Returns an integer representing the day of the year indicated by the specified absolute time.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
// Returns an integer representing the day of the year indicated by the specified absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetDayOfYear(_:_:)
func CFAbsoluteTimeGetDayOfYear(at AbsoluteTime, tz TimeZoneRef) unsafe.Pointer {
	return _CFAbsoluteTimeGetDayOfYear(at, tz)
}

// Computes the time difference between two specified absolute times and returns the result as an interval in Gregorian units.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
// Computes the time difference between two specified absolute times and returns the result as an interval in Gregorian units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetDifferenceAsGregorianUnits(_:_:_:_:)
func CFAbsoluteTimeGetDifferenceAsGregorianUnits(at1 AbsoluteTime, at2 AbsoluteTime, tz TimeZoneRef, unitFlags OptionFlags) GregorianUnits {
	return _CFAbsoluteTimeGetDifferenceAsGregorianUnits(at1, at2, tz, unitFlags)
}

// Converts an absolute time value into a Gregorian date.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
// Converts an absolute time value into a Gregorian date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetGregorianDate(_:_:)
func CFAbsoluteTimeGetGregorianDate(at AbsoluteTime, tz TimeZoneRef) GregorianDate {
	return _CFAbsoluteTimeGetGregorianDate(at, tz)
}

// Returns an integer representing the week of the year indicated by the specified absolute time.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
// Returns an integer representing the week of the year indicated by the specified absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTimeGetWeekOfYear(_:_:)
func CFAbsoluteTimeGetWeekOfYear(at AbsoluteTime, tz TimeZoneRef) unsafe.Pointer {
	return _CFAbsoluteTimeGetWeekOfYear(at, tz)
}

// Allocates memory using the specified allocator.
//
// Added in macOS .
// Allocates memory using the specified allocator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocate(_:_:_:)
func CFAllocatorAllocate(allocator AllocatorRef, size Index, hint OptionFlags) unsafe.Pointer {
	return _CFAllocatorAllocate(allocator, size, hint)
}

// CFAllocatorAllocateBytes is a CoreFoundation function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocateBytes(_:_:_:)
func CFAllocatorAllocateBytes(allocator AllocatorRef, size Index, hint OptionFlags) unsafe.Pointer {
	return _CFAllocatorAllocateBytes(allocator, size, hint)
}

// CFAllocatorAllocateTyped is a CoreFoundation function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocateTyped(_:_:_:_:)
func CFAllocatorAllocateTyped(allocator AllocatorRef, size Index, descriptor AllocatorTypeID, hint OptionFlags) unsafe.Pointer {
	return _CFAllocatorAllocateTyped(allocator, size, descriptor, hint)
}

// Creates an allocator object.
//
// Added in macOS .
// Creates an allocator object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorCreate(_:_:)
func CFAllocatorCreate(allocator AllocatorRef, context unsafe.Pointer) AllocatorRef {
	return _CFAllocatorCreate(allocator, context)
}

// CFAllocatorCreateWithZone is a CoreFoundation function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorCreateWithZone
func CFAllocatorCreateWithZone(allocator AllocatorRef, zone unsafe.Pointer) AllocatorRef {
	return _CFAllocatorCreateWithZone(allocator, zone)
}

// Deallocates a block of memory with a given allocator.
//
// Added in macOS .
// Deallocates a block of memory with a given allocator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorDeallocate(_:_:)
func CFAllocatorDeallocate(allocator AllocatorRef, ptr unsafe.Pointer) {
	_CFAllocatorDeallocate(allocator, ptr)
}

// Obtains the context of the specified allocator or of the default allocator.
//
// Added in macOS .
// Obtains the context of the specified allocator or of the default allocator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetContext(_:_:)
func CFAllocatorGetContext(allocator AllocatorRef, context unsafe.Pointer) {
	_CFAllocatorGetContext(allocator, context)
}

// Gets the default allocator object for the current thread.
//
// Added in macOS .
// Gets the default allocator object for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetDefault()
func CFAllocatorGetDefault() AllocatorRef {
	return _CFAllocatorGetDefault()
}

// Obtains the number of bytes likely to be allocated upon a specific request.
//
// Added in macOS .
// Obtains the number of bytes likely to be allocated upon a specific request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetPreferredSizeForSize(_:_:_:)
func CFAllocatorGetPreferredSizeForSize(allocator AllocatorRef, size Index, hint OptionFlags) Index {
	return _CFAllocatorGetPreferredSizeForSize(allocator, size, hint)
}

// Returns the type identifier for the CFAllocator opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFAllocator opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetTypeID()
func CFAllocatorGetTypeID() TypeID {
	return _CFAllocatorGetTypeID()
}

// Reallocates memory using the specified allocator.
//
// Added in macOS .
// Reallocates memory using the specified allocator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocate(_:_:_:_:)
func CFAllocatorReallocate(allocator AllocatorRef, ptr unsafe.Pointer, newsize Index, hint OptionFlags) unsafe.Pointer {
	return _CFAllocatorReallocate(allocator, ptr, newsize, hint)
}

// CFAllocatorReallocateBytes is a CoreFoundation function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocateBytes(_:_:_:_:)
func CFAllocatorReallocateBytes(allocator AllocatorRef, ptr unsafe.Pointer, newsize Index, hint OptionFlags) unsafe.Pointer {
	return _CFAllocatorReallocateBytes(allocator, ptr, newsize, hint)
}

// CFAllocatorReallocateTyped is a CoreFoundation function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocateTyped(_:_:_:_:_:)
func CFAllocatorReallocateTyped(allocator AllocatorRef, ptr unsafe.Pointer, newsize Index, descriptor AllocatorTypeID, hint OptionFlags) unsafe.Pointer {
	return _CFAllocatorReallocateTyped(allocator, ptr, newsize, descriptor, hint)
}

// Sets the given allocator as the default for the current thread.
//
// Added in macOS .
// Sets the given allocator as the default for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorSetDefault(_:)
func CFAllocatorSetDefault(allocator AllocatorRef) {
	_CFAllocatorSetDefault(allocator)
}

// Adds the values from one array to another array.
//
// Added in macOS .
// Adds the values from one array to another array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayAppendArray(_:_:_:)
func CFArrayAppendArray(theArray MutableArrayRef, otherArray ArrayRef, otherRange Range) {
	_CFArrayAppendArray(theArray, otherArray, otherRange)
}

// Adds a value to an array giving it the new largest index.
//
// Added in macOS .
// Adds a value to an array giving it the new largest index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayAppendValue(_:_:)
func CFArrayAppendValue(theArray MutableArrayRef, value unsafe.Pointer) {
	_CFArrayAppendValue(theArray, value)
}

// Calls a function once for each element in range in an array.
//
// Added in macOS .
// Calls a function once for each element in range in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayApplyFunction(_:_:_:_:)
func CFArrayApplyFunction(theArray ArrayRef, range_ Range, applier ArrayApplierFunction, context unsafe.Pointer) {
	_CFArrayApplyFunction(theArray, range_, applier, context)
}

// Searches an array for a value using a binary search algorithm.
//
// Added in macOS .
// Searches an array for a value using a binary search algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayBSearchValues(_:_:_:_:_:)
func CFArrayBSearchValues(theArray ArrayRef, range_ Range, value unsafe.Pointer, comparator ComparatorFunction, context unsafe.Pointer) Index {
	return _CFArrayBSearchValues(theArray, range_, value, comparator, context)
}

// Reports whether or not a value is in an array.
//
// Added in macOS .
// Reports whether or not a value is in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayContainsValue(_:_:_:)
func CFArrayContainsValue(theArray ArrayRef, range_ Range, value unsafe.Pointer) unsafe.Pointer {
	return _CFArrayContainsValue(theArray, range_, value)
}

// Creates a new immutable array with the given values.
//
// Added in macOS .
// Creates a new immutable array with the given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreate(_:_:_:_:)
func CFArrayCreate(allocator AllocatorRef, values unsafe.Pointer, numValues Index, callBacks unsafe.Pointer) ArrayRef {
	return _CFArrayCreate(allocator, values, numValues, callBacks)
}

// Creates a new immutable array with the values from another array.
//
// Added in macOS .
// Creates a new immutable array with the values from another array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateCopy(_:_:)
func CFArrayCreateCopy(allocator AllocatorRef, theArray ArrayRef) ArrayRef {
	return _CFArrayCreateCopy(allocator, theArray)
}

// Creates a new empty mutable array.
//
// Added in macOS .
// Creates a new empty mutable array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateMutable(_:_:_:)
func CFArrayCreateMutable(allocator AllocatorRef, capacity Index, callBacks unsafe.Pointer) MutableArrayRef {
	return _CFArrayCreateMutable(allocator, capacity, callBacks)
}

// Creates a new mutable array with the values from another array.
//
// Added in macOS .
// Creates a new mutable array with the values from another array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateMutableCopy(_:_:_:)
func CFArrayCreateMutableCopy(allocator AllocatorRef, capacity Index, theArray ArrayRef) MutableArrayRef {
	return _CFArrayCreateMutableCopy(allocator, capacity, theArray)
}

// Exchanges the values at two indices of an array.
//
// Added in macOS .
// Exchanges the values at two indices of an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayExchangeValuesAtIndices(_:_:_:)
func CFArrayExchangeValuesAtIndices(theArray MutableArrayRef, idx1 Index, idx2 Index) {
	_CFArrayExchangeValuesAtIndices(theArray, idx1, idx2)
}

// Returns the number of values currently in an array.
//
// Added in macOS .
// Returns the number of values currently in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetCount(_:)
func CFArrayGetCount(theArray ArrayRef) Index {
	return _CFArrayGetCount(theArray)
}

// Counts the number of times a given value occurs in an array.
//
// Added in macOS .
// Counts the number of times a given value occurs in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetCountOfValue(_:_:_:)
func CFArrayGetCountOfValue(theArray ArrayRef, range_ Range, value unsafe.Pointer) Index {
	return _CFArrayGetCountOfValue(theArray, range_, value)
}

// Searches an array forward for a value.
//
// Added in macOS .
// Searches an array forward for a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetFirstIndexOfValue(_:_:_:)
func CFArrayGetFirstIndexOfValue(theArray ArrayRef, range_ Range, value unsafe.Pointer) Index {
	return _CFArrayGetFirstIndexOfValue(theArray, range_, value)
}

// Searches an array backward for a value.
//
// Added in macOS .
// Searches an array backward for a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetLastIndexOfValue(_:_:_:)
func CFArrayGetLastIndexOfValue(theArray ArrayRef, range_ Range, value unsafe.Pointer) Index {
	return _CFArrayGetLastIndexOfValue(theArray, range_, value)
}

// Returns the type identifier for the CFArray opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFArray opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetTypeID()
func CFArrayGetTypeID() TypeID {
	return _CFArrayGetTypeID()
}

// Retrieves a value at a given index.
//
// Added in macOS .
// Retrieves a value at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetValueAtIndex(_:_:)
func CFArrayGetValueAtIndex(theArray ArrayRef, idx Index) unsafe.Pointer {
	return _CFArrayGetValueAtIndex(theArray, idx)
}

// Fills a buffer with values from an array.
//
// Added in macOS .
// Fills a buffer with values from an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetValues(_:_:_:)
func CFArrayGetValues(theArray ArrayRef, range_ Range, values unsafe.Pointer) {
	_CFArrayGetValues(theArray, range_, values)
}

// Inserts a value into an array at a given index.
//
// Added in macOS .
// Inserts a value into an array at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayInsertValueAtIndex(_:_:_:)
func CFArrayInsertValueAtIndex(theArray MutableArrayRef, idx Index, value unsafe.Pointer) {
	_CFArrayInsertValueAtIndex(theArray, idx, value)
}

// Removes all the values from an array, making it empty.
//
// Added in macOS .
// Removes all the values from an array, making it empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayRemoveAllValues(_:)
func CFArrayRemoveAllValues(theArray MutableArrayRef) {
	_CFArrayRemoveAllValues(theArray)
}

// Removes the value at a given index from an array.
//
// Added in macOS .
// Removes the value at a given index from an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayRemoveValueAtIndex(_:_:)
func CFArrayRemoveValueAtIndex(theArray MutableArrayRef, idx Index) {
	_CFArrayRemoveValueAtIndex(theArray, idx)
}

// Replaces a range of values in an array.
//
// Added in macOS .
// Replaces a range of values in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayReplaceValues(_:_:_:_:)
func CFArrayReplaceValues(theArray MutableArrayRef, range_ Range, newValues unsafe.Pointer, newCount Index) {
	_CFArrayReplaceValues(theArray, range_, newValues, newCount)
}

// Changes the value at a given index in an array.
//
// Added in macOS .
// Changes the value at a given index in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArraySetValueAtIndex(_:_:_:)
func CFArraySetValueAtIndex(theArray MutableArrayRef, idx Index, value unsafe.Pointer) {
	_CFArraySetValueAtIndex(theArray, idx, value)
}

// Sorts the values in an array using a given comparison function.
//
// Added in macOS .
// Sorts the values in an array using a given comparison function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArraySortValues(_:_:_:_:)
func CFArraySortValues(theArray MutableArrayRef, range_ Range, comparator ComparatorFunction, context unsafe.Pointer) {
	_CFArraySortValues(theArray, range_, comparator, context)
}

// Defers internal consistency-checking and coalescing for a mutable attributed string.
//
// Added in macOS .
// Defers internal consistency-checking and coalescing for a mutable attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringBeginEditing(_:)
func CFAttributedStringBeginEditing(aStr MutableAttributedStringRef) {
	_CFAttributedStringBeginEditing(aStr)
}

// Creates an attributed string with specified string and attributes.
//
// Added in macOS .
// Creates an attributed string with specified string and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreate(_:_:_:)
func CFAttributedStringCreate(alloc AllocatorRef, str StringRef, attributes DictionaryRef) AttributedStringRef {
	return _CFAttributedStringCreate(alloc, str, attributes)
}

// Creates an immutable copy of an attributed string.
//
// Added in macOS .
// Creates an immutable copy of an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateCopy(_:_:)
func CFAttributedStringCreateCopy(alloc AllocatorRef, aStr AttributedStringRef) AttributedStringRef {
	return _CFAttributedStringCreateCopy(alloc, aStr)
}

// Creates a mutable attributed string.
//
// Added in macOS .
// Creates a mutable attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateMutable(_:_:)
func CFAttributedStringCreateMutable(alloc AllocatorRef, maxLength Index) MutableAttributedStringRef {
	return _CFAttributedStringCreateMutable(alloc, maxLength)
}

// Creates a mutable copy of an attributed string.
//
// Added in macOS .
// Creates a mutable copy of an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateMutableCopy(_:_:_:)
func CFAttributedStringCreateMutableCopy(alloc AllocatorRef, maxLength Index, aStr AttributedStringRef) MutableAttributedStringRef {
	return _CFAttributedStringCreateMutableCopy(alloc, maxLength, aStr)
}

// Creates a sub-attributed string from the specified range.
//
// Added in macOS .
// Creates a sub-attributed string from the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateWithSubstring(_:_:_:)
func CFAttributedStringCreateWithSubstring(alloc AllocatorRef, aStr AttributedStringRef, range_ Range) AttributedStringRef {
	return _CFAttributedStringCreateWithSubstring(alloc, aStr, range_)
}

// Re-enables internal consistency-checking and coalescing for a mutable attributed string.
//
// Added in macOS .
// Re-enables internal consistency-checking and coalescing for a mutable attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringEndEditing(_:)
func CFAttributedStringEndEditing(aStr MutableAttributedStringRef) {
	_CFAttributedStringEndEditing(aStr)
}

// Returns the value of a given attribute of an attributed string at a specified location.
//
// Added in macOS .
// Returns the value of a given attribute of an attributed string at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttribute(_:_:_:_:)
func CFAttributedStringGetAttribute(aStr AttributedStringRef, loc Index, attrName StringRef, effectiveRange unsafe.Pointer) TypeRef {
	return _CFAttributedStringGetAttribute(aStr, loc, attrName, effectiveRange)
}

// Returns the value of a given attribute of an attributed string at a specified location.
//
// Added in macOS .
// Returns the value of a given attribute of an attributed string at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributeAndLongestEffectiveRange(_:_:_:_:_:)
func CFAttributedStringGetAttributeAndLongestEffectiveRange(aStr AttributedStringRef, loc Index, attrName StringRef, inRange Range, longestEffectiveRange unsafe.Pointer) TypeRef {
	return _CFAttributedStringGetAttributeAndLongestEffectiveRange(aStr, loc, attrName, inRange, longestEffectiveRange)
}

// Returns the attributes of an attributed string at a specified location.
//
// Added in macOS .
// Returns the attributes of an attributed string at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributes(_:_:_:)
func CFAttributedStringGetAttributes(aStr AttributedStringRef, loc Index, effectiveRange unsafe.Pointer) DictionaryRef {
	return _CFAttributedStringGetAttributes(aStr, loc, effectiveRange)
}

// Returns the attributes of an attributed string at a specified location.
//
// Added in macOS .
// Returns the attributes of an attributed string at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributesAndLongestEffectiveRange(_:_:_:_:)
func CFAttributedStringGetAttributesAndLongestEffectiveRange(aStr AttributedStringRef, loc Index, inRange Range, longestEffectiveRange unsafe.Pointer) DictionaryRef {
	return _CFAttributedStringGetAttributesAndLongestEffectiveRange(aStr, loc, inRange, longestEffectiveRange)
}

// CFAttributedStringGetBidiLevelsAndResolvedDirections is a CoreFoundation function.
//
// Added in macOS .
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetBidiLevelsAndResolvedDirections(_:_:_:_:_:)
func CFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString AttributedStringRef, range_ Range, baseDirection int8, bidiLevels unsafe.Pointer, baseDirections unsafe.Pointer) bool {
	return _CFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString, range_, baseDirection, bidiLevels, baseDirections)
}

// Returns the length of the attributed string in characters.
//
// Added in macOS .
// Returns the length of the attributed string in characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetLength(_:)
func CFAttributedStringGetLength(aStr AttributedStringRef) Index {
	return _CFAttributedStringGetLength(aStr)
}

// Gets as a mutable string the string for an attributed string.
//
// Added in macOS .
// Gets as a mutable string the string for an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetMutableString(_:)
func CFAttributedStringGetMutableString(aStr MutableAttributedStringRef) MutableStringRef {
	return _CFAttributedStringGetMutableString(aStr)
}

// CFAttributedStringGetStatisticalWritingDirections is a CoreFoundation function.
//
// Added in macOS 26.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetStatisticalWritingDirections(_:_:_:_:_:)
func CFAttributedStringGetStatisticalWritingDirections(attributedString AttributedStringRef, range_ Range, baseDirection int8, bidiLevels unsafe.Pointer, baseDirections unsafe.Pointer) bool {
	return _CFAttributedStringGetStatisticalWritingDirections(attributedString, range_, baseDirection, bidiLevels, baseDirections)
}

// Returns the string for an attributed string.
//
// Added in macOS .
// Returns the string for an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetString(_:)
func CFAttributedStringGetString(aStr AttributedStringRef) StringRef {
	return _CFAttributedStringGetString(aStr)
}

// Returns the type identifier for the CFAttributedString opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFAttributedString opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetTypeID()
func CFAttributedStringGetTypeID() TypeID {
	return _CFAttributedStringGetTypeID()
}

// Removes the value of a single attribute over a specified range.
//
// Added in macOS .
// Removes the value of a single attribute over a specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringRemoveAttribute(_:_:_:)
func CFAttributedStringRemoveAttribute(aStr MutableAttributedStringRef, range_ Range, attrName StringRef) {
	_CFAttributedStringRemoveAttribute(aStr, range_, attrName)
}

// Replaces the attributed substring over a range with another attributed string.
//
// Added in macOS .
// Replaces the attributed substring over a range with another attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringReplaceAttributedString(_:_:_:)
func CFAttributedStringReplaceAttributedString(aStr MutableAttributedStringRef, range_ Range, replacement AttributedStringRef) {
	_CFAttributedStringReplaceAttributedString(aStr, range_, replacement)
}

// Modifies the string of an attributed string.
//
// Added in macOS .
// Modifies the string of an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringReplaceString(_:_:_:)
func CFAttributedStringReplaceString(aStr MutableAttributedStringRef, range_ Range, replacement StringRef) {
	_CFAttributedStringReplaceString(aStr, range_, replacement)
}

// Sets the value of a single attribute over the specified range.
//
// Added in macOS .
// Sets the value of a single attribute over the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringSetAttribute(_:_:_:_:)
func CFAttributedStringSetAttribute(aStr MutableAttributedStringRef, range_ Range, attrName StringRef, value TypeRef) {
	_CFAttributedStringSetAttribute(aStr, range_, attrName, value)
}

// Sets the value of attributes of a mutable attributed string over a specified range.
//
// Added in macOS .
// Sets the value of attributes of a mutable attributed string over a specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringSetAttributes(_:_:_:_:)
func CFAttributedStringSetAttributes(aStr MutableAttributedStringRef, range_ Range, replacement DictionaryRef, clearOtherAttributes unsafe.Pointer) {
	_CFAttributedStringSetAttributes(aStr, range_, replacement, clearOtherAttributes)
}

// CFAutorelease is a CoreFoundation function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAutorelease
func CFAutorelease(arg TypeRef) TypeRef {
	return _CFAutorelease(arg)
}

// Adds a value to a mutable bag.
//
// Added in macOS .
// Adds a value to a mutable bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagAddValue(_:_:)
func CFBagAddValue(theBag MutableBagRef, value unsafe.Pointer) {
	_CFBagAddValue(theBag, value)
}

// Calls a function once for each value in a bag.
//
// Added in macOS .
// Calls a function once for each value in a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagApplyFunction(_:_:_:)
func CFBagApplyFunction(theBag BagRef, applier BagApplierFunction, context unsafe.Pointer) {
	_CFBagApplyFunction(theBag, applier, context)
}

// Reports whether or not a value is in a bag.
//
// Added in macOS .
// Reports whether or not a value is in a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagContainsValue(_:_:)
func CFBagContainsValue(theBag BagRef, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagContainsValue(theBag, value)
}

// Creates an immutable bag containing specified values.
//
// Added in macOS .
// Creates an immutable bag containing specified values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCreate(_:_:_:_:)
func CFBagCreate(allocator AllocatorRef, values unsafe.Pointer, numValues Index, callBacks unsafe.Pointer) BagRef {
	return _CFBagCreate(allocator, values, numValues, callBacks)
}

// Creates an immutable bag with the values of another bag.
//
// Added in macOS .
// Creates an immutable bag with the values of another bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateCopy(_:_:)
func CFBagCreateCopy(allocator AllocatorRef, theBag BagRef) BagRef {
	return _CFBagCreateCopy(allocator, theBag)
}

// Creates a new empty mutable bag.
//
// Added in macOS .
// Creates a new empty mutable bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateMutable(_:_:_:)
func CFBagCreateMutable(allocator AllocatorRef, capacity Index, callBacks unsafe.Pointer) MutableBagRef {
	return _CFBagCreateMutable(allocator, capacity, callBacks)
}

// Creates a new mutable bag with the values from another bag.
//
// Added in macOS .
// Creates a new mutable bag with the values from another bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateMutableCopy(_:_:_:)
func CFBagCreateMutableCopy(allocator AllocatorRef, capacity Index, theBag BagRef) MutableBagRef {
	return _CFBagCreateMutableCopy(allocator, capacity, theBag)
}

// Returns the number of values currently in a bag.
//
// Added in macOS .
// Returns the number of values currently in a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetCount(_:)
func CFBagGetCount(theBag BagRef) Index {
	return _CFBagGetCount(theBag)
}

// Returns the number of times a value occurs in a bag.
//
// Added in macOS .
// Returns the number of times a value occurs in a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetCountOfValue(_:_:)
func CFBagGetCountOfValue(theBag BagRef, value unsafe.Pointer) Index {
	return _CFBagGetCountOfValue(theBag, value)
}

// Returns the type identifier for the CFBag opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFBag opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetTypeID()
func CFBagGetTypeID() TypeID {
	return _CFBagGetTypeID()
}

// Returns a requested value from a bag.
//
// Added in macOS .
// Returns a requested value from a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValue(_:_:)
func CFBagGetValue(theBag BagRef, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagGetValue(theBag, value)
}

// Reports whether or not a value is in a bag, and returns that value indirectly if it exists.
//
// Added in macOS .
// Reports whether or not a value is in a bag, and returns that value indirectly if it exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValueIfPresent(_:_:_:)
func CFBagGetValueIfPresent(theBag BagRef, candidate unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagGetValueIfPresent(theBag, candidate, value)
}

// Fills a buffer with values from a bag.
//
// Added in macOS .
// Fills a buffer with values from a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValues(_:_:)
func CFBagGetValues(theBag BagRef, values unsafe.Pointer) {
	_CFBagGetValues(theBag, values)
}

// Removes all values from a mutable bag.
//
// Added in macOS .
// Removes all values from a mutable bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagRemoveAllValues(_:)
func CFBagRemoveAllValues(theBag MutableBagRef) {
	_CFBagRemoveAllValues(theBag)
}

// Removes a value from a mutable bag.
//
// Added in macOS .
// Removes a value from a mutable bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagRemoveValue(_:_:)
func CFBagRemoveValue(theBag MutableBagRef, value unsafe.Pointer) {
	_CFBagRemoveValue(theBag, value)
}

// Replaces a value in a mutable bag.
//
// Added in macOS .
// Replaces a value in a mutable bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagReplaceValue(_:_:)
func CFBagReplaceValue(theBag MutableBagRef, value unsafe.Pointer) {
	_CFBagReplaceValue(theBag, value)
}

// Sets a value in a mutable bag.
//
// Added in macOS .
// Sets a value in a mutable bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagSetValue(_:_:)
func CFBagSetValue(theBag MutableBagRef, value unsafe.Pointer) {
	_CFBagSetValue(theBag, value)
}

// Adds a value to a binary heap.
//
// Added in macOS .
// Adds a value to a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapAddValue(_:_:)
func CFBinaryHeapAddValue(heap BinaryHeapRef, value unsafe.Pointer) {
	_CFBinaryHeapAddValue(heap, value)
}

// Iteratively applies a function to all the values in a binary heap.
//
// Added in macOS .
// Iteratively applies a function to all the values in a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapApplyFunction(_:_:_:)
func CFBinaryHeapApplyFunction(heap BinaryHeapRef, applier BinaryHeapApplierFunction, context unsafe.Pointer) {
	_CFBinaryHeapApplyFunction(heap, applier, context)
}

// Returns whether a given value is in a binary heap.
//
// Added in macOS .
// Returns whether a given value is in a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapContainsValue(_:_:)
func CFBinaryHeapContainsValue(heap BinaryHeapRef, value unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapContainsValue(heap, value)
}

// Creates a new mutable or fixed-mutable binary heap.
//
// Added in macOS .
// Creates a new mutable or fixed-mutable binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCreate(_:_:_:_:)
func CFBinaryHeapCreate(allocator AllocatorRef, capacity Index, callBacks unsafe.Pointer, compareContext unsafe.Pointer) BinaryHeapRef {
	return _CFBinaryHeapCreate(allocator, capacity, callBacks, compareContext)
}

// Creates a new mutable or fixed-mutable binary heap with the values from a pre-existing binary heap.
//
// Added in macOS .
// Creates a new mutable or fixed-mutable binary heap with the values from a pre-existing binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCreateCopy(_:_:_:)
func CFBinaryHeapCreateCopy(allocator AllocatorRef, capacity Index, heap BinaryHeapRef) BinaryHeapRef {
	return _CFBinaryHeapCreateCopy(allocator, capacity, heap)
}

// Returns the number of values currently in a binary heap.
//
// Added in macOS .
// Returns the number of values currently in a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetCount(_:)
func CFBinaryHeapGetCount(heap BinaryHeapRef) Index {
	return _CFBinaryHeapGetCount(heap)
}

// Counts the number of times a given value occurs in a binary heap.
//
// Added in macOS .
// Counts the number of times a given value occurs in a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetCountOfValue(_:_:)
func CFBinaryHeapGetCountOfValue(heap BinaryHeapRef, value unsafe.Pointer) Index {
	return _CFBinaryHeapGetCountOfValue(heap, value)
}

// Returns the minimum value in a binary heap.
//
// Added in macOS .
// Returns the minimum value in a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetMinimum(_:)
func CFBinaryHeapGetMinimum(heap BinaryHeapRef) unsafe.Pointer {
	return _CFBinaryHeapGetMinimum(heap)
}

// Returns the minimum value in a binary heap, if present.
//
// Added in macOS .
// Returns the minimum value in a binary heap, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetMinimumIfPresent(_:_:)
func CFBinaryHeapGetMinimumIfPresent(heap BinaryHeapRef, value unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapGetMinimumIfPresent(heap, value)
}

// Returns the type identifier of the opaque type.
//
// Added in macOS .
// Returns the type identifier of the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetTypeID()
func CFBinaryHeapGetTypeID() TypeID {
	return _CFBinaryHeapGetTypeID()
}

// Copies all the values from a binary heap into a sorted C array.
//
// Added in macOS .
// Copies all the values from a binary heap into a sorted C array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetValues(_:_:)
func CFBinaryHeapGetValues(heap BinaryHeapRef, values unsafe.Pointer) {
	_CFBinaryHeapGetValues(heap, values)
}

// Removes all values from a binary heap, making it empty.
//
// Added in macOS .
// Removes all values from a binary heap, making it empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapRemoveAllValues(_:)
func CFBinaryHeapRemoveAllValues(heap BinaryHeapRef) {
	_CFBinaryHeapRemoveAllValues(heap)
}

// Removes the minimum value from a binary heap.
//
// Added in macOS .
// Removes the minimum value from a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapRemoveMinimumValue(_:)
func CFBinaryHeapRemoveMinimumValue(heap BinaryHeapRef) {
	_CFBinaryHeapRemoveMinimumValue(heap)
}

// Returns whether a bit vector contains a particular bit value.
//
// Added in macOS .
// Returns whether a bit vector contains a particular bit value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorContainsBit(_:_:_:)
func CFBitVectorContainsBit(bv BitVectorRef, range_ Range, value Bit) unsafe.Pointer {
	return _CFBitVectorContainsBit(bv, range_, value)
}

// Creates an immutable bit vector from a block of memory.
//
// Added in macOS .
// Creates an immutable bit vector from a block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreate(_:_:_:)
func CFBitVectorCreate(allocator AllocatorRef, bytes unsafe.Pointer, numBits Index) BitVectorRef {
	return _CFBitVectorCreate(allocator, bytes, numBits)
}

// Creates an immutable bit vector that is a copy of another bit vector.
//
// Added in macOS .
// Creates an immutable bit vector that is a copy of another bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateCopy(_:_:)
func CFBitVectorCreateCopy(allocator AllocatorRef, bv BitVectorRef) BitVectorRef {
	return _CFBitVectorCreateCopy(allocator, bv)
}

// Creates a mutable bit vector.
//
// Added in macOS .
// Creates a mutable bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateMutable(_:_:)
func CFBitVectorCreateMutable(allocator AllocatorRef, capacity Index) MutableBitVectorRef {
	return _CFBitVectorCreateMutable(allocator, capacity)
}

// Creates a new mutable bit vector from a pre-existing bit vector.
//
// Added in macOS .
// Creates a new mutable bit vector from a pre-existing bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorCreateMutableCopy(_:_:_:)
func CFBitVectorCreateMutableCopy(allocator AllocatorRef, capacity Index, bv BitVectorRef) MutableBitVectorRef {
	return _CFBitVectorCreateMutableCopy(allocator, capacity, bv)
}

// Flips a bit value in a bit vector.
//
// Added in macOS .
// Flips a bit value in a bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorFlipBitAtIndex(_:_:)
func CFBitVectorFlipBitAtIndex(bv MutableBitVectorRef, idx Index) {
	_CFBitVectorFlipBitAtIndex(bv, idx)
}

// Flips a range of bit values in a bit vector.
//
// Added in macOS .
// Flips a range of bit values in a bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorFlipBits(_:_:)
func CFBitVectorFlipBits(bv MutableBitVectorRef, range_ Range) {
	_CFBitVectorFlipBits(bv, range_)
}

// Returns the bit value at a given index in a bit vector.
//
// Added in macOS .
// Returns the bit value at a given index in a bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetBitAtIndex(_:_:)
func CFBitVectorGetBitAtIndex(bv BitVectorRef, idx Index) Bit {
	return _CFBitVectorGetBitAtIndex(bv, idx)
}

// Returns the bit values in a range of indices in a bit vector.
//
// Added in macOS .
// Returns the bit values in a range of indices in a bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetBits(_:_:_:)
func CFBitVectorGetBits(bv BitVectorRef, range_ Range, bytes unsafe.Pointer) {
	_CFBitVectorGetBits(bv, range_, bytes)
}

// Returns the number of bit values in a bit vector.
//
// Added in macOS .
// Returns the number of bit values in a bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetCount(_:)
func CFBitVectorGetCount(bv BitVectorRef) Index {
	return _CFBitVectorGetCount(bv)
}

// Counts the number of times a certain bit value occurs within a range of bits in a bit vector.
//
// Added in macOS .
// Counts the number of times a certain bit value occurs within a range of bits in a bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetCountOfBit(_:_:_:)
func CFBitVectorGetCountOfBit(bv BitVectorRef, range_ Range, value Bit) Index {
	return _CFBitVectorGetCountOfBit(bv, range_, value)
}

// Locates the first occurrence of a certain bit value within a range of bits in a bit vector.
//
// Added in macOS .
// Locates the first occurrence of a certain bit value within a range of bits in a bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetFirstIndexOfBit(_:_:_:)
func CFBitVectorGetFirstIndexOfBit(bv BitVectorRef, range_ Range, value Bit) Index {
	return _CFBitVectorGetFirstIndexOfBit(bv, range_, value)
}

// Locates the last occurrence of a certain bit value within a range of bits in a bit vector.
//
// Added in macOS .
// Locates the last occurrence of a certain bit value within a range of bits in a bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetLastIndexOfBit(_:_:_:)
func CFBitVectorGetLastIndexOfBit(bv BitVectorRef, range_ Range, value Bit) Index {
	return _CFBitVectorGetLastIndexOfBit(bv, range_, value)
}

// Returns the type identifier for the CFBitVector opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFBitVector opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetTypeID()
func CFBitVectorGetTypeID() TypeID {
	return _CFBitVectorGetTypeID()
}

// Sets all bits in a bit vector to a particular value.
//
// Added in macOS .
// Sets all bits in a bit vector to a particular value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetAllBits(_:_:)
func CFBitVectorSetAllBits(bv MutableBitVectorRef, value Bit) {
	_CFBitVectorSetAllBits(bv, value)
}

// Sets the value of a particular bit in a bit vector.
//
// Added in macOS .
// Sets the value of a particular bit in a bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetBitAtIndex(_:_:_:)
func CFBitVectorSetBitAtIndex(bv MutableBitVectorRef, idx Index, value Bit) {
	_CFBitVectorSetBitAtIndex(bv, idx, value)
}

// Sets a range of bits in a bit vector to a particular value.
//
// Added in macOS .
// Sets a range of bits in a bit vector to a particular value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetBits(_:_:_:)
func CFBitVectorSetBits(bv MutableBitVectorRef, range_ Range, value Bit) {
	_CFBitVectorSetBits(bv, range_, value)
}

// Changes the size of a mutable bit vector.
//
// Added in macOS .
// Changes the size of a mutable bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorSetCount(_:_:)
func CFBitVectorSetCount(bv MutableBitVectorRef, count Index) {
	_CFBitVectorSetCount(bv, count)
}

// Returns the Core Foundation type identifier for the CFBoolean opaque type.
//
// Added in macOS .
// Returns the Core Foundation type identifier for the CFBoolean opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBooleanGetTypeID()
func CFBooleanGetTypeID() TypeID {
	return _CFBooleanGetTypeID()
}

// Returns the value of a CFBoolean object as a standard C type .
//
// Added in macOS .
// Returns the value of a CFBoolean object as a standard C type .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBooleanGetValue(_:)
func CFBooleanGetValue(boolean BooleanRef) unsafe.Pointer {
	return _CFBooleanGetValue(boolean)
}

// Closes an open resource map for a bundle.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.0.
// Closes an open resource map for a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCloseBundleResourceMap(_:_:)
func CFBundleCloseBundleResourceMap(bundle BundleRef, refNum BundleRefNum) {
	_CFBundleCloseBundleResourceMap(bundle, refNum)
}

// Returns the location of a bundle’s auxiliary executable code.
//
// Added in macOS .
// Returns the location of a bundle’s auxiliary executable code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyAuxiliaryExecutableURL(_:_:)
func CFBundleCopyAuxiliaryExecutableURL(bundle BundleRef, executableName StringRef) URLRef {
	return _CFBundleCopyAuxiliaryExecutableURL(bundle, executableName)
}

// Returns the location of a bundle’s built in plug-in.
//
// Added in macOS .
// Returns the location of a bundle’s built in plug-in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBuiltInPlugInsURL(_:)
func CFBundleCopyBuiltInPlugInsURL(bundle BundleRef) URLRef {
	return _CFBundleCopyBuiltInPlugInsURL(bundle)
}

// Returns an array containing a bundle’s localizations.
//
// Added in macOS .
// Returns an array containing a bundle’s localizations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBundleLocalizations(_:)
func CFBundleCopyBundleLocalizations(bundle BundleRef) ArrayRef {
	return _CFBundleCopyBundleLocalizations(bundle)
}

// Returns the location of a bundle.
//
// Added in macOS .
// Returns the location of a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBundleURL(_:)
func CFBundleCopyBundleURL(bundle BundleRef) URLRef {
	return _CFBundleCopyBundleURL(bundle)
}

// Returns an array of CFNumbers representing the architectures a given bundle provides.
//
// Added in macOS 10.5.
// Returns an array of CFNumbers representing the architectures a given bundle provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableArchitectures(_:)
func CFBundleCopyExecutableArchitectures(bundle BundleRef) ArrayRef {
	return _CFBundleCopyExecutableArchitectures(bundle)
}

// Returns an array of CFNumbers representing the architectures a given URL provides.
//
// Added in macOS 10.5.
// Returns an array of CFNumbers representing the architectures a given URL provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableArchitecturesForURL(_:)
func CFBundleCopyExecutableArchitecturesForURL(url URLRef) ArrayRef {
	return _CFBundleCopyExecutableArchitecturesForURL(url)
}

// Returns the location of a bundle’s main executable code.
//
// Added in macOS .
// Returns the location of a bundle’s main executable code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableURL(_:)
func CFBundleCopyExecutableURL(bundle BundleRef) URLRef {
	return _CFBundleCopyExecutableURL(bundle)
}

// Returns the information dictionary for a given URL location.
//
// Added in macOS .
// Returns the information dictionary for a given URL location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyInfoDictionaryForURL(_:)
func CFBundleCopyInfoDictionaryForURL(url URLRef) DictionaryRef {
	return _CFBundleCopyInfoDictionaryForURL(url)
}

// Returns a bundle’s information dictionary.
//
// Added in macOS .
// Returns a bundle’s information dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyInfoDictionaryInDirectory(_:)
func CFBundleCopyInfoDictionaryInDirectory(bundleURL URLRef) DictionaryRef {
	return _CFBundleCopyInfoDictionaryInDirectory(bundleURL)
}

// Given an array of possible localizations and preferred locations, returns the one or more of them that CFBundle would use, without reference to the current application context.
//
// Added in macOS .
// Given an array of possible localizations and preferred locations, returns the one or more of them that CFBundle would use, without reference to the current application context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizationsForPreferences(_:_:)
func CFBundleCopyLocalizationsForPreferences(locArray ArrayRef, prefArray ArrayRef) ArrayRef {
	return _CFBundleCopyLocalizationsForPreferences(locArray, prefArray)
}

// Returns an array containing the localizations for a bundle or executable at a particular location.
//
// Added in macOS .
// Returns an array containing the localizations for a bundle or executable at a particular location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizationsForURL(_:)
func CFBundleCopyLocalizationsForURL(url URLRef) ArrayRef {
	return _CFBundleCopyLocalizationsForURL(url)
}

// Returns a localized string from a bundle’s strings file.
//
// Added in macOS .
// Returns a localized string from a bundle’s strings file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizedString(_:_:_:_:)
func CFBundleCopyLocalizedString(bundle BundleRef, key StringRef, value StringRef, tableName StringRef) StringRef {
	return _CFBundleCopyLocalizedString(bundle, key, value, tableName)
}

// Returns a localized string from a bundle’s strings file.
//
// Added in macOS 15.4.
// Returns a localized string from a bundle’s strings file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizedStringForLocalizations(_:_:_:_:_:)
func CFBundleCopyLocalizedStringForLocalizations(bundle BundleRef, key StringRef, value StringRef, tableName StringRef, localizations ArrayRef) StringRef {
	return _CFBundleCopyLocalizedStringForLocalizations(bundle, key, value, tableName, localizations)
}

// Given an array of possible localizations, returns the one or more of them that CFBundle would use in the current application context.
//
// Added in macOS .
// Given an array of possible localizations, returns the one or more of them that CFBundle would use in the current application context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyPreferredLocalizationsFromArray(_:)
func CFBundleCopyPreferredLocalizationsFromArray(locArray ArrayRef) ArrayRef {
	return _CFBundleCopyPreferredLocalizationsFromArray(locArray)
}

// Returns the location of a bundle’s private Frameworks directory.
//
// Added in macOS .
// Returns the location of a bundle’s private Frameworks directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyPrivateFrameworksURL(_:)
func CFBundleCopyPrivateFrameworksURL(bundle BundleRef) URLRef {
	return _CFBundleCopyPrivateFrameworksURL(bundle)
}

// Returns the location of a resource contained in the specified bundle.
//
// Added in macOS .
// Returns the location of a resource contained in the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURL(_:_:_:_:)
func CFBundleCopyResourceURL(bundle BundleRef, resourceName StringRef, resourceType StringRef, subDirName StringRef) URLRef {
	return _CFBundleCopyResourceURL(bundle, resourceName, resourceType, subDirName)
}

// Returns the location of a localized resource in a bundle.
//
// Added in macOS .
// Returns the location of a localized resource in a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLForLocalization(_:_:_:_:_:)
func CFBundleCopyResourceURLForLocalization(bundle BundleRef, resourceName StringRef, resourceType StringRef, subDirName StringRef, localizationName StringRef) URLRef {
	return _CFBundleCopyResourceURLForLocalization(bundle, resourceName, resourceType, subDirName, localizationName)
}

// Returns the location of a resource contained in the specified bundle directory without requiring the creation of a CFBundle object.
//
// Added in macOS .
// Returns the location of a resource contained in the specified bundle directory without requiring the creation of a CFBundle object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLInDirectory(_:_:_:_:)
func CFBundleCopyResourceURLInDirectory(bundleURL URLRef, resourceName StringRef, resourceType StringRef, subDirName StringRef) URLRef {
	return _CFBundleCopyResourceURLInDirectory(bundleURL, resourceName, resourceType, subDirName)
}

// Assembles an array of URLs specifying all of the resources of the specified type found in a bundle.
//
// Added in macOS .
// Assembles an array of URLs specifying all of the resources of the specified type found in a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfType(_:_:_:)
func CFBundleCopyResourceURLsOfType(bundle BundleRef, resourceType StringRef, subDirName StringRef) ArrayRef {
	return _CFBundleCopyResourceURLsOfType(bundle, resourceType, subDirName)
}

// Returns an array containing copies of the URL locations for a specified bundle, resource, and localization name.
//
// Added in macOS .
// Returns an array containing copies of the URL locations for a specified bundle, resource, and localization name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfTypeForLocalization(_:_:_:_:)
func CFBundleCopyResourceURLsOfTypeForLocalization(bundle BundleRef, resourceType StringRef, subDirName StringRef, localizationName StringRef) ArrayRef {
	return _CFBundleCopyResourceURLsOfTypeForLocalization(bundle, resourceType, subDirName, localizationName)
}

// Returns an array of CFURL objects describing the locations of all resources in a bundle of the specified type without needing to create a CFBundle object.
//
// Added in macOS .
// Returns an array of CFURL objects describing the locations of all resources in a bundle of the specified type without needing to create a CFBundle object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfTypeInDirectory(_:_:_:)
func CFBundleCopyResourceURLsOfTypeInDirectory(bundleURL URLRef, resourceType StringRef, subDirName StringRef) ArrayRef {
	return _CFBundleCopyResourceURLsOfTypeInDirectory(bundleURL, resourceType, subDirName)
}

// Returns the location of a bundle’s Resources directory.
//
// Added in macOS .
// Returns the location of a bundle’s Resources directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourcesDirectoryURL(_:)
func CFBundleCopyResourcesDirectoryURL(bundle BundleRef) URLRef {
	return _CFBundleCopyResourcesDirectoryURL(bundle)
}

// Returns the location of a bundle’s shared frameworks directory.
//
// Added in macOS .
// Returns the location of a bundle’s shared frameworks directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySharedFrameworksURL(_:)
func CFBundleCopySharedFrameworksURL(bundle BundleRef) URLRef {
	return _CFBundleCopySharedFrameworksURL(bundle)
}

// Returns the location of a bundle’s shared support files directory.
//
// Added in macOS .
// Returns the location of a bundle’s shared support files directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySharedSupportURL(_:)
func CFBundleCopySharedSupportURL(bundle BundleRef) URLRef {
	return _CFBundleCopySharedSupportURL(bundle)
}

// Returns the location of the bundle’s support files directory.
//
// Added in macOS .
// Returns the location of the bundle’s support files directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySupportFilesDirectoryURL(_:)
func CFBundleCopySupportFilesDirectoryURL(bundle BundleRef) URLRef {
	return _CFBundleCopySupportFilesDirectoryURL(bundle)
}

// Creates a CFBundle object.
//
// Added in macOS .
// Creates a CFBundle object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCreate(_:_:)
func CFBundleCreate(allocator AllocatorRef, bundleURL URLRef) BundleRef {
	return _CFBundleCreate(allocator, bundleURL)
}

// Searches a directory and constructs an array of CFBundle objects from all valid bundles in the specified directory.
//
// Added in macOS .
// Searches a directory and constructs an array of CFBundle objects from all valid bundles in the specified directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCreateBundlesFromDirectory(_:_:_:)
func CFBundleCreateBundlesFromDirectory(allocator AllocatorRef, directoryURL URLRef, bundleType StringRef) ArrayRef {
	return _CFBundleCreateBundlesFromDirectory(allocator, directoryURL, bundleType)
}

// Returns an array containing all of the bundles currently open in the application.
//
// Added in macOS .
// Returns an array containing all of the bundles currently open in the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetAllBundles()
func CFBundleGetAllBundles() ArrayRef {
	return _CFBundleGetAllBundles()
}

// Locate a bundle given its program-defined identifier.
//
// Added in macOS .
// Locate a bundle given its program-defined identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetBundleWithIdentifier(_:)
func CFBundleGetBundleWithIdentifier(bundleID StringRef) BundleRef {
	return _CFBundleGetBundleWithIdentifier(bundleID)
}

// Returns a data pointer to a symbol of the given name.
//
// Added in macOS .
// Returns a data pointer to a symbol of the given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDataPointerForName(_:_:)
func CFBundleGetDataPointerForName(bundle BundleRef, symbolName StringRef) unsafe.Pointer {
	return _CFBundleGetDataPointerForName(bundle, symbolName)
}

// Returns a C array of data pointer to symbols of the given names.
//
// Added in macOS .
// Returns a C array of data pointer to symbols of the given names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDataPointersForNames(_:_:_:)
func CFBundleGetDataPointersForNames(bundle BundleRef, symbolNames ArrayRef, stbl unsafe.Pointer) {
	_CFBundleGetDataPointersForNames(bundle, symbolNames, stbl)
}

// Returns the bundle’s development region from the bundle’s information property list.
//
// Added in macOS .
// Returns the bundle’s development region from the bundle’s information property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDevelopmentRegion(_:)
func CFBundleGetDevelopmentRegion(bundle BundleRef) StringRef {
	return _CFBundleGetDevelopmentRegion(bundle)
}

// Returns a pointer to a function in a bundle’s executable code using the function name as the search key.
//
// Added in macOS .
// Returns a pointer to a function in a bundle’s executable code using the function name as the search key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetFunctionPointerForName(_:_:)
func CFBundleGetFunctionPointerForName(bundle BundleRef, functionName StringRef) unsafe.Pointer {
	return _CFBundleGetFunctionPointerForName(bundle, functionName)
}

// Constructs a function table containing pointers to all of the functions found in a bundle’s main executable code.
//
// Added in macOS .
// Constructs a function table containing pointers to all of the functions found in a bundle’s main executable code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetFunctionPointersForNames(_:_:_:)
func CFBundleGetFunctionPointersForNames(bundle BundleRef, functionNames ArrayRef, ftbl unsafe.Pointer) {
	_CFBundleGetFunctionPointersForNames(bundle, functionNames, ftbl)
}

// Returns the bundle identifier from a bundle’s information property list.
//
// Added in macOS .
// Returns the bundle identifier from a bundle’s information property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetIdentifier(_:)
func CFBundleGetIdentifier(bundle BundleRef) StringRef {
	return _CFBundleGetIdentifier(bundle)
}

// Returns a bundle’s information dictionary.
//
// Added in macOS .
// Returns a bundle’s information dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetInfoDictionary(_:)
func CFBundleGetInfoDictionary(bundle BundleRef) DictionaryRef {
	return _CFBundleGetInfoDictionary(bundle)
}

// Returns a bundle’s localized information dictionary.
//
// Added in macOS .
// Returns a bundle’s localized information dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetLocalInfoDictionary(_:)
func CFBundleGetLocalInfoDictionary(bundle BundleRef) DictionaryRef {
	return _CFBundleGetLocalInfoDictionary(bundle)
}

// Returns an application’s main bundle.
//
// Added in macOS .
// Returns an application’s main bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetMainBundle()
func CFBundleGetMainBundle() BundleRef {
	return _CFBundleGetMainBundle()
}

// Returns a bundle’s package type and creator.
//
// Added in macOS .
// Returns a bundle’s package type and creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPackageInfo(_:_:_:)
func CFBundleGetPackageInfo(bundle BundleRef, packageType unsafe.Pointer, packageCreator unsafe.Pointer) {
	_CFBundleGetPackageInfo(bundle, packageType, packageCreator)
}

// Returns a bundle’s package type and creator without having to create a CFBundle object.
//
// Added in macOS .
// Returns a bundle’s package type and creator without having to create a CFBundle object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPackageInfoInDirectory(_:_:_:)
func CFBundleGetPackageInfoInDirectory(url URLRef, packageType unsafe.Pointer, packageCreator unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetPackageInfoInDirectory(url, packageType, packageCreator)
}

// Returns a bundle’s plug-in.
//
// Added in macOS .
// Returns a bundle’s plug-in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPlugIn(_:)
func CFBundleGetPlugIn(bundle BundleRef) PlugInRef {
	return _CFBundleGetPlugIn(bundle)
}

// Returns the type identifier for the CFBundle opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFBundle opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetTypeID()
func CFBundleGetTypeID() TypeID {
	return _CFBundleGetTypeID()
}

// Returns a value (localized if possible) from a bundle’s information dictionary.
//
// Added in macOS .
// Returns a value (localized if possible) from a bundle’s information dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetValueForInfoDictionaryKey(_:_:)
func CFBundleGetValueForInfoDictionaryKey(bundle BundleRef, key StringRef) TypeRef {
	return _CFBundleGetValueForInfoDictionaryKey(bundle, key)
}

// Returns a bundle’s version number.
//
// Added in macOS .
// Returns a bundle’s version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetVersionNumber(_:)
func CFBundleGetVersionNumber(bundle BundleRef) unsafe.Pointer {
	return _CFBundleGetVersionNumber(bundle)
}

// CFBundleIsArchitectureLoadable is a CoreFoundation function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsArchitectureLoadable(_:)
func CFBundleIsArchitectureLoadable(arch unsafe.Pointer) unsafe.Pointer {
	return _CFBundleIsArchitectureLoadable(arch)
}

// CFBundleIsExecutableLoadable is a CoreFoundation function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoadable(_:)
func CFBundleIsExecutableLoadable(bundle BundleRef) unsafe.Pointer {
	return _CFBundleIsExecutableLoadable(bundle)
}

// CFBundleIsExecutableLoadableForURL is a CoreFoundation function.
//
// Added in macOS 11.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoadableForURL(_:)
func CFBundleIsExecutableLoadableForURL(url URLRef) unsafe.Pointer {
	return _CFBundleIsExecutableLoadableForURL(url)
}

// Obtains information about the load status for a bundle’s main executable.
//
// Added in macOS .
// Obtains information about the load status for a bundle’s main executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoaded(_:)
func CFBundleIsExecutableLoaded(bundle BundleRef) unsafe.Pointer {
	return _CFBundleIsExecutableLoaded(bundle)
}

// Loads a bundle’s main executable code into memory and dynamically links it into the running application.
//
// Added in macOS .
// Loads a bundle’s main executable code into memory and dynamically links it into the running application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleLoadExecutable(_:)
func CFBundleLoadExecutable(bundle BundleRef) unsafe.Pointer {
	return _CFBundleLoadExecutable(bundle)
}

// Returns a Boolean value that indicates whether a given bundle is loaded, attempting to load it if necessary.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether a given bundle is loaded, attempting to load it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleLoadExecutableAndReturnError(_:_:)
func CFBundleLoadExecutableAndReturnError(bundle BundleRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFBundleLoadExecutableAndReturnError(bundle, error_)
}

// Opens the non-localized and localized resource files (if any) for a bundle in separate resource maps.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.0.
// Opens the non-localized and localized resource files (if any) for a bundle in separate resource maps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleOpenBundleResourceFiles(_:_:_:)
func CFBundleOpenBundleResourceFiles(bundle BundleRef, refNum unsafe.Pointer, localizedRefNum unsafe.Pointer) unsafe.Pointer {
	return _CFBundleOpenBundleResourceFiles(bundle, refNum, localizedRefNum)
}

// Opens the non-localized and localized resource files (if any) for a bundle in a single resource map.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.0.
// Opens the non-localized and localized resource files (if any) for a bundle in a single resource map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleOpenBundleResourceMap(_:)
func CFBundleOpenBundleResourceMap(bundle BundleRef) BundleRefNum {
	return _CFBundleOpenBundleResourceMap(bundle)
}

// Returns a Boolean value that indicates whether a given bundle is loaded or appears to be loadable.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether a given bundle is loaded or appears to be loadable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundlePreflightExecutable(_:_:)
func CFBundlePreflightExecutable(bundle BundleRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFBundlePreflightExecutable(bundle, error_)
}

// Unloads the main executable for the specified bundle.
//
// Added in macOS .
// Unloads the main executable for the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleUnloadExecutable(_:)
func CFBundleUnloadExecutable(bundle BundleRef) {
	_CFBundleUnloadExecutable(bundle)
}

// Computes the absolute time when specified components are added to a given absolute time.
//
// Added in macOS .
// Computes the absolute time when specified components are added to a given absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarAddComponents
func CFCalendarAddComponents(calendar CalendarRef, at unsafe.Pointer, options OptionFlags, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarAddComponents(calendar, at, options, componentDesc)
}

// Computes the absolute time from components in a description string.
//
// Added in macOS .
// Computes the absolute time from components in a description string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarComposeAbsoluteTime
func CFCalendarComposeAbsoluteTime(calendar CalendarRef, at unsafe.Pointer, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarComposeAbsoluteTime(calendar, at, componentDesc)
}

// Returns a copy of the logical calendar for the current user.
//
// Added in macOS .
// Returns a copy of the logical calendar for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyCurrent()
func CFCalendarCopyCurrent() CalendarRef {
	return _CFCalendarCopyCurrent()
}

// Returns a locale object for a specified calendar.
//
// Added in macOS .
// Returns a locale object for a specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyLocale(_:)
func CFCalendarCopyLocale(calendar CalendarRef) LocaleRef {
	return _CFCalendarCopyLocale(calendar)
}

// Returns a time zone object for a specified calendar.
//
// Added in macOS .
// Returns a time zone object for a specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyTimeZone(_:)
func CFCalendarCopyTimeZone(calendar CalendarRef) TimeZoneRef {
	return _CFCalendarCopyTimeZone(calendar)
}

// Returns a calendar object for the calendar identified by a calendar identifier.
//
// Added in macOS .
// Returns a calendar object for the calendar identified by a calendar identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCreateWithIdentifier(_:_:)
func CFCalendarCreateWithIdentifier(allocator AllocatorRef, identifier CalendarIdentifier) CalendarRef {
	return _CFCalendarCreateWithIdentifier(allocator, identifier)
}

// Computes the components which are indicated by the componentDesc description string for the given absolute time.
//
// Added in macOS .
// Computes the components which are indicated by the componentDesc description string for the given absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarDecomposeAbsoluteTime
func CFCalendarDecomposeAbsoluteTime(calendar CalendarRef, at AbsoluteTime, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarDecomposeAbsoluteTime(calendar, at, componentDesc)
}

// Computes the difference between the two absolute times, in terms of specified calendrical components.
//
// Added in macOS .
// Computes the difference between the two absolute times, in terms of specified calendrical components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetComponentDifference
func CFCalendarGetComponentDifference(calendar CalendarRef, startingAT AbsoluteTime, resultAT AbsoluteTime, options OptionFlags, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetComponentDifference(calendar, startingAT, resultAT, options, componentDesc)
}

// Returns the index of first weekday for a specified calendar.
//
// Added in macOS .
// Returns the index of first weekday for a specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetFirstWeekday(_:)
func CFCalendarGetFirstWeekday(calendar CalendarRef) Index {
	return _CFCalendarGetFirstWeekday(calendar)
}

// Returns the given calendar’s identifier.
//
// Added in macOS .
// Returns the given calendar’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetIdentifier(_:)
func CFCalendarGetIdentifier(calendar CalendarRef) CalendarIdentifier {
	return _CFCalendarGetIdentifier(calendar)
}

// Returns the maximum range limits of the values that a specified unit can take on in a given calendar.
//
// Added in macOS .
// Returns the maximum range limits of the values that a specified unit can take on in a given calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMaximumRangeOfUnit(_:_:)
func CFCalendarGetMaximumRangeOfUnit(calendar CalendarRef, unit CalendarUnit) Range {
	return _CFCalendarGetMaximumRangeOfUnit(calendar, unit)
}

// Returns the minimum number of days in the first week of a specified calendar.
//
// Added in macOS .
// Returns the minimum number of days in the first week of a specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMinimumDaysInFirstWeek(_:)
func CFCalendarGetMinimumDaysInFirstWeek(calendar CalendarRef) Index {
	return _CFCalendarGetMinimumDaysInFirstWeek(calendar)
}

// Returns the minimum range limits of the values that a specified unit can take on in a given calendar.
//
// Added in macOS .
// Returns the minimum range limits of the values that a specified unit can take on in a given calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMinimumRangeOfUnit(_:_:)
func CFCalendarGetMinimumRangeOfUnit(calendar CalendarRef, unit CalendarUnit) Range {
	return _CFCalendarGetMinimumRangeOfUnit(calendar, unit)
}

// Returns the ordinal number of a calendrical unit within a larger unit at a specified absolute time.
//
// Added in macOS .
// Returns the ordinal number of a calendrical unit within a larger unit at a specified absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetOrdinalityOfUnit(_:_:_:_:)
func CFCalendarGetOrdinalityOfUnit(calendar CalendarRef, smallerUnit CalendarUnit, biggerUnit CalendarUnit, at AbsoluteTime) Index {
	return _CFCalendarGetOrdinalityOfUnit(calendar, smallerUnit, biggerUnit, at)
}

// Returns the range of values that one unit can take on within a larger unit during which a specific absolute time occurs.
//
// Added in macOS .
// Returns the range of values that one unit can take on within a larger unit during which a specific absolute time occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetRangeOfUnit(_:_:_:_:)
func CFCalendarGetRangeOfUnit(calendar CalendarRef, smallerUnit CalendarUnit, biggerUnit CalendarUnit, at AbsoluteTime) Range {
	return _CFCalendarGetRangeOfUnit(calendar, smallerUnit, biggerUnit, at)
}

// Returns by reference the start time and duration of a given calendar unit that contains a given absolute time.
//
// Added in macOS 10.5.
// Returns by reference the start time and duration of a given calendar unit that contains a given absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetTimeRangeOfUnit(_:_:_:_:_:)
func CFCalendarGetTimeRangeOfUnit(calendar CalendarRef, unit CalendarUnit, at AbsoluteTime, startp unsafe.Pointer, tip unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetTimeRangeOfUnit(calendar, unit, at, startp, tip)
}

// Returns the type identifier for the CFCalendar opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFCalendar opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetTypeID()
func CFCalendarGetTypeID() TypeID {
	return _CFCalendarGetTypeID()
}

// Sets the first weekday for a calendar.
//
// Added in macOS .
// Sets the first weekday for a calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetFirstWeekday(_:_:)
func CFCalendarSetFirstWeekday(calendar CalendarRef, wkdy Index) {
	_CFCalendarSetFirstWeekday(calendar, wkdy)
}

// Sets the locale for a calendar.
//
// Added in macOS .
// Sets the locale for a calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetLocale(_:_:)
func CFCalendarSetLocale(calendar CalendarRef, locale LocaleRef) {
	_CFCalendarSetLocale(calendar, locale)
}

// Sets the minimum number of days in the first week of a specified calendar.
//
// Added in macOS .
// Sets the minimum number of days in the first week of a specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetMinimumDaysInFirstWeek(_:_:)
func CFCalendarSetMinimumDaysInFirstWeek(calendar CalendarRef, mwd Index) {
	_CFCalendarSetMinimumDaysInFirstWeek(calendar, mwd)
}

// Sets the time zone for a calendar.
//
// Added in macOS .
// Sets the time zone for a calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetTimeZone(_:_:)
func CFCalendarSetTimeZone(calendar CalendarRef, tz TimeZoneRef) {
	_CFCalendarSetTimeZone(calendar, tz)
}

// Adds a given range to a character set.
//
// Added in macOS .
// Adds a given range to a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetAddCharactersInRange(_:_:)
func CFCharacterSetAddCharactersInRange(theSet MutableCharacterSetRef, theRange Range) {
	_CFCharacterSetAddCharactersInRange(theSet, theRange)
}

// Adds the characters in a given string to a character set.
//
// Added in macOS .
// Adds the characters in a given string to a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetAddCharactersInString(_:_:)
func CFCharacterSetAddCharactersInString(theSet MutableCharacterSetRef, theString StringRef) {
	_CFCharacterSetAddCharactersInString(theSet, theString)
}

// Creates a new immutable data with the bitmap representation from the given character set.
//
// Added in macOS .
// Creates a new immutable data with the bitmap representation from the given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateBitmapRepresentation(_:_:)
func CFCharacterSetCreateBitmapRepresentation(alloc AllocatorRef, theSet CharacterSetRef) DataRef {
	return _CFCharacterSetCreateBitmapRepresentation(alloc, theSet)
}

// Creates a new character set with the values from a given character set.
//
// Added in macOS .
// Creates a new character set with the values from a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateCopy(_:_:)
func CFCharacterSetCreateCopy(alloc AllocatorRef, theSet CharacterSetRef) CharacterSetRef {
	return _CFCharacterSetCreateCopy(alloc, theSet)
}

// Creates a new immutable character set that is the invert of the specified character set.
//
// Added in macOS .
// Creates a new immutable character set that is the invert of the specified character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateInvertedSet(_:_:)
func CFCharacterSetCreateInvertedSet(alloc AllocatorRef, theSet CharacterSetRef) CharacterSetRef {
	return _CFCharacterSetCreateInvertedSet(alloc, theSet)
}

// Creates a new empty mutable character set.
//
// Added in macOS .
// Creates a new empty mutable character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateMutable(_:)
func CFCharacterSetCreateMutable(alloc AllocatorRef) MutableCharacterSetRef {
	return _CFCharacterSetCreateMutable(alloc)
}

// Creates a new mutable character set with the values from another character set.
//
// Added in macOS .
// Creates a new mutable character set with the values from another character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateMutableCopy(_:_:)
func CFCharacterSetCreateMutableCopy(alloc AllocatorRef, theSet CharacterSetRef) MutableCharacterSetRef {
	return _CFCharacterSetCreateMutableCopy(alloc, theSet)
}

// Creates a new immutable character set with the bitmap representation specified by given data.
//
// Added in macOS .
// Creates a new immutable character set with the bitmap representation specified by given data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithBitmapRepresentation(_:_:)
func CFCharacterSetCreateWithBitmapRepresentation(alloc AllocatorRef, theData DataRef) CharacterSetRef {
	return _CFCharacterSetCreateWithBitmapRepresentation(alloc, theData)
}

// Creates a new character set with the values from the given range of Unicode characters.
//
// Added in macOS .
// Creates a new character set with the values from the given range of Unicode characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithCharactersInRange(_:_:)
func CFCharacterSetCreateWithCharactersInRange(alloc AllocatorRef, theRange Range) CharacterSetRef {
	return _CFCharacterSetCreateWithCharactersInRange(alloc, theRange)
}

// Creates a new character set with the values in the given string.
//
// Added in macOS .
// Creates a new character set with the values in the given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithCharactersInString(_:_:)
func CFCharacterSetCreateWithCharactersInString(alloc AllocatorRef, theString StringRef) CharacterSetRef {
	return _CFCharacterSetCreateWithCharactersInString(alloc, theString)
}

// Returns a predefined character set.
//
// Added in macOS .
// Returns a predefined character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetGetPredefined(_:)
func CFCharacterSetGetPredefined(theSetIdentifier CharacterSetPredefinedSet) CharacterSetRef {
	return _CFCharacterSetGetPredefined(theSetIdentifier)
}

// Returns the type identifier of the CFCharacterSet opaque type.
//
// Added in macOS .
// Returns the type identifier of the CFCharacterSet opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetGetTypeID()
func CFCharacterSetGetTypeID() TypeID {
	return _CFCharacterSetGetTypeID()
}

// Reports whether or not a character set contains at least one member character in the specified plane.
//
// Added in macOS .
// Reports whether or not a character set contains at least one member character in the specified plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetHasMemberInPlane(_:_:)
func CFCharacterSetHasMemberInPlane(theSet CharacterSetRef, thePlane Index) unsafe.Pointer {
	return _CFCharacterSetHasMemberInPlane(theSet, thePlane)
}

// Forms an intersection of two character sets.
//
// Added in macOS .
// Forms an intersection of two character sets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIntersect(_:_:)
func CFCharacterSetIntersect(theSet MutableCharacterSetRef, theOtherSet CharacterSetRef) {
	_CFCharacterSetIntersect(theSet, theOtherSet)
}

// Inverts the content of a given character set.
//
// Added in macOS .
// Inverts the content of a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetInvert(_:)
func CFCharacterSetInvert(theSet MutableCharacterSetRef) {
	_CFCharacterSetInvert(theSet)
}

// Reports whether or not a given Unicode character is in a character set.
//
// Added in macOS .
// Reports whether or not a given Unicode character is in a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsCharacterMember(_:_:)
func CFCharacterSetIsCharacterMember(theSet CharacterSetRef, theChar unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetIsCharacterMember(theSet, theChar)
}

// Reports whether or not a given UTF-32 character is in a character set.
//
// Added in macOS .
// Reports whether or not a given UTF-32 character is in a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsLongCharacterMember(_:_:)
func CFCharacterSetIsLongCharacterMember(theSet CharacterSetRef, theChar unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetIsLongCharacterMember(theSet, theChar)
}

// Reports whether or not a character set is a superset of another set.
//
// Added in macOS .
// Reports whether or not a character set is a superset of another set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsSupersetOfSet(_:_:)
func CFCharacterSetIsSupersetOfSet(theSet CharacterSetRef, theOtherset CharacterSetRef) unsafe.Pointer {
	return _CFCharacterSetIsSupersetOfSet(theSet, theOtherset)
}

// Removes a given range of Unicode characters from a character set.
//
// Added in macOS .
// Removes a given range of Unicode characters from a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetRemoveCharactersInRange(_:_:)
func CFCharacterSetRemoveCharactersInRange(theSet MutableCharacterSetRef, theRange Range) {
	_CFCharacterSetRemoveCharactersInRange(theSet, theRange)
}

// Removes the characters in a given string from a character set.
//
// Added in macOS .
// Removes the characters in a given string from a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetRemoveCharactersInString(_:_:)
func CFCharacterSetRemoveCharactersInString(theSet MutableCharacterSetRef, theString StringRef) {
	_CFCharacterSetRemoveCharactersInString(theSet, theString)
}

// Forms the union of two character sets.
//
// Added in macOS .
// Forms the union of two character sets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetUnion(_:_:)
func CFCharacterSetUnion(theSet MutableCharacterSetRef, theOtherSet CharacterSetRef) {
	_CFCharacterSetUnion(theSet, theOtherSet)
}

// Returns a textual description of a Core Foundation object.
//
// Added in macOS .
// Returns a textual description of a Core Foundation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCopyDescription(_:)
func CFCopyDescription(cf TypeRef) StringRef {
	return _CFCopyDescription(cf)
}

// CFCopyHomeDirectoryURL is a CoreFoundation function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCopyHomeDirectoryURL()
func CFCopyHomeDirectoryURL() URLRef {
	return _CFCopyHomeDirectoryURL()
}

// Returns a textual description of a Core Foundation type, as identified by its type ID, which can be used when debugging.
//
// Added in macOS .
// Returns a textual description of a Core Foundation type, as identified by its type ID, which can be used when debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCopyTypeIDDescription(_:)
func CFCopyTypeIDDescription(type_id TypeID) StringRef {
	return _CFCopyTypeIDDescription(type_id)
}

// Appends the bytes from a byte buffer to the contents of a CFData object.
//
// Added in macOS .
// Appends the bytes from a byte buffer to the contents of a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataAppendBytes(_:_:_:)
func CFDataAppendBytes(theData MutableDataRef, bytes unsafe.Pointer, length Index) {
	_CFDataAppendBytes(theData, bytes, length)
}

// Creates an immutable CFData object using data copied from a specified byte buffer.
//
// Added in macOS .
// Creates an immutable CFData object using data copied from a specified byte buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreate(_:_:_:)
func CFDataCreate(allocator AllocatorRef, bytes unsafe.Pointer, length Index) DataRef {
	return _CFDataCreate(allocator, bytes, length)
}

// Creates an immutable copy of a CFData object.
//
// Added in macOS .
// Creates an immutable copy of a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateCopy(_:_:)
func CFDataCreateCopy(allocator AllocatorRef, theData DataRef) DataRef {
	return _CFDataCreateCopy(allocator, theData)
}

// Creates an empty CFMutableData object.
//
// Added in macOS .
// Creates an empty CFMutableData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateMutable(_:_:)
func CFDataCreateMutable(allocator AllocatorRef, capacity Index) MutableDataRef {
	return _CFDataCreateMutable(allocator, capacity)
}

// Creates a CFMutableData object by copying another CFData object.
//
// Added in macOS .
// Creates a CFMutableData object by copying another CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateMutableCopy(_:_:_:)
func CFDataCreateMutableCopy(allocator AllocatorRef, capacity Index, theData DataRef) MutableDataRef {
	return _CFDataCreateMutableCopy(allocator, capacity, theData)
}

// Creates an immutable CFData object from an external (client-owned) byte buffer.
//
// Added in macOS .
// Creates an immutable CFData object from an external (client-owned) byte buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateWithBytesNoCopy(_:_:_:_:)
func CFDataCreateWithBytesNoCopy(allocator AllocatorRef, bytes unsafe.Pointer, length Index, bytesDeallocator AllocatorRef) DataRef {
	return _CFDataCreateWithBytesNoCopy(allocator, bytes, length, bytesDeallocator)
}

// Deletes the bytes in a CFMutableData object within a specified range.
//
// Added in macOS .
// Deletes the bytes in a CFMutableData object within a specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataDeleteBytes(_:_:)
func CFDataDeleteBytes(theData MutableDataRef, range_ Range) {
	_CFDataDeleteBytes(theData, range_)
}

// Finds and returns the range within a data object of the first occurrence of the given data, within a given range, subject to any given options.
//
// Added in macOS 10.6.
// Finds and returns the range within a data object of the first occurrence of the given data, within a given range, subject to any given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataFind(_:_:_:_:)
func CFDataFind(theData DataRef, dataToFind DataRef, searchRange Range, compareOptions DataSearchFlags) Range {
	return _CFDataFind(theData, dataToFind, searchRange, compareOptions)
}

// Returns a read-only pointer to the bytes of a CFData object.
//
// Added in macOS .
// Returns a read-only pointer to the bytes of a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetBytePtr(_:)
func CFDataGetBytePtr(theData DataRef) unsafe.Pointer {
	return _CFDataGetBytePtr(theData)
}

// Copies the byte contents of a CFData object to an external buffer.
//
// Added in macOS .
// Copies the byte contents of a CFData object to an external buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetBytes(_:_:_:)
func CFDataGetBytes(theData DataRef, range_ Range, buffer unsafe.Pointer) {
	_CFDataGetBytes(theData, range_, buffer)
}

// Returns the number of bytes contained by a CFData object.
//
// Added in macOS .
// Returns the number of bytes contained by a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetLength(_:)
func CFDataGetLength(theData DataRef) Index {
	return _CFDataGetLength(theData)
}

// Returns a pointer to a mutable byte buffer of a CFMutableData object.
//
// Added in macOS .
// Returns a pointer to a mutable byte buffer of a CFMutableData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetMutableBytePtr(_:)
func CFDataGetMutableBytePtr(theData MutableDataRef) unsafe.Pointer {
	return _CFDataGetMutableBytePtr(theData)
}

// Returns the type identifier for the CFData opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFData opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetTypeID()
func CFDataGetTypeID() TypeID {
	return _CFDataGetTypeID()
}

// Increases the length of a CFMutableData object’s internal byte buffer, zero-filling the extension to the buffer.
//
// Added in macOS .
// Increases the length of a CFMutableData object’s internal byte buffer, zero-filling the extension to the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataIncreaseLength(_:_:)
func CFDataIncreaseLength(theData MutableDataRef, extraLength Index) {
	_CFDataIncreaseLength(theData, extraLength)
}

// Replaces those bytes in a CFMutableData object that fall within a specified range with other bytes.
//
// Added in macOS .
// Replaces those bytes in a CFMutableData object that fall within a specified range with other bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataReplaceBytes(_:_:_:_:)
func CFDataReplaceBytes(theData MutableDataRef, range_ Range, newBytes unsafe.Pointer, newLength Index) {
	_CFDataReplaceBytes(theData, range_, newBytes, newLength)
}

// Resets the length of a CFMutableData object’s internal byte buffer.
//
// Added in macOS .
// Resets the length of a CFMutableData object’s internal byte buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSetLength(_:_:)
func CFDataSetLength(theData MutableDataRef, length Index) {
	_CFDataSetLength(theData, length)
}

// Compares two objects and returns a comparison result.
//
// Added in macOS .
// Compares two objects and returns a comparison result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateCompare(_:_:_:)
func CFDateCompare(theDate DateRef, otherDate DateRef, context unsafe.Pointer) ComparisonResult {
	return _CFDateCompare(theDate, otherDate, context)
}

// Creates a object given an absolute time.
//
// Added in macOS .
// Creates a object given an absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateCreate(_:_:)
func CFDateCreate(allocator AllocatorRef, at AbsoluteTime) DateRef {
	return _CFDateCreate(allocator, at)
}

// Returns a copy of a date formatter’s value for a given key.
//
// Added in macOS .
// Returns a copy of a date formatter’s value for a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCopyProperty(_:_:)
func CFDateFormatterCopyProperty(formatter DateFormatterRef, key DateFormatterKey) TypeRef {
	return _CFDateFormatterCopyProperty(formatter, key)
}

// Creates a new CFDateFormatter object, localized to the given locale, which will format dates to the given date and time styles.
//
// Added in macOS .
// Creates a new CFDateFormatter object, localized to the given locale, which will format dates to the given date and time styles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreate(_:_:_:_:)
func CFDateFormatterCreate(allocator AllocatorRef, locale LocaleRef, dateStyle DateFormatterStyle, timeStyle DateFormatterStyle) DateFormatterRef {
	return _CFDateFormatterCreate(allocator, locale, dateStyle, timeStyle)
}

// Returns a localized date format string representing the given date format components arranged appropriately for the specified locale.
//
// Added in macOS 10.6.
// Returns a localized date format string representing the given date format components arranged appropriately for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateDateFormatFromTemplate(_:_:_:_:)
func CFDateFormatterCreateDateFormatFromTemplate(allocator AllocatorRef, tmplate StringRef, options OptionFlags, locale LocaleRef) StringRef {
	return _CFDateFormatterCreateDateFormatFromTemplate(allocator, tmplate, options, locale)
}

// Returns a date object representing a given string.
//
// Added in macOS .
// Returns a date object representing a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateDateFromString(_:_:_:_:)
func CFDateFormatterCreateDateFromString(allocator AllocatorRef, formatter DateFormatterRef, string_ StringRef, rangep unsafe.Pointer) DateRef {
	return _CFDateFormatterCreateDateFromString(allocator, formatter, string_, rangep)
}

// CFDateFormatterCreateISO8601Formatter is a CoreFoundation function.
//
// Added in macOS 10.12.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateISO8601Formatter(_:_:)
func CFDateFormatterCreateISO8601Formatter(allocator AllocatorRef, formatOptions ISO8601DateFormatOptions) DateFormatterRef {
	return _CFDateFormatterCreateISO8601Formatter(allocator, formatOptions)
}

// Returns a string representation of the given absolute time using the specified date formatter.
//
// Added in macOS .
// Returns a string representation of the given absolute time using the specified date formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateStringWithAbsoluteTime(_:_:_:)
func CFDateFormatterCreateStringWithAbsoluteTime(allocator AllocatorRef, formatter DateFormatterRef, at AbsoluteTime) StringRef {
	return _CFDateFormatterCreateStringWithAbsoluteTime(allocator, formatter, at)
}

// Returns a string representation of the given date using the specified date formatter.
//
// Added in macOS .
// Returns a string representation of the given date using the specified date formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateStringWithDate(_:_:_:)
func CFDateFormatterCreateStringWithDate(allocator AllocatorRef, formatter DateFormatterRef, date DateRef) StringRef {
	return _CFDateFormatterCreateStringWithDate(allocator, formatter, date)
}

// Returns an absolute time object representing a given string.
//
// Added in macOS .
// Returns an absolute time object representing a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetAbsoluteTimeFromString(_:_:_:_:)
func CFDateFormatterGetAbsoluteTimeFromString(formatter DateFormatterRef, string_ StringRef, rangep unsafe.Pointer, atp unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetAbsoluteTimeFromString(formatter, string_, rangep, atp)
}

// Returns the date style used to create the given date formatter object.
//
// Added in macOS .
// Returns the date style used to create the given date formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetDateStyle(_:)
func CFDateFormatterGetDateStyle(formatter DateFormatterRef) DateFormatterStyle {
	return _CFDateFormatterGetDateStyle(formatter)
}

// Returns a format string for the given date formatter object.
//
// Added in macOS .
// Returns a format string for the given date formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetFormat(_:)
func CFDateFormatterGetFormat(formatter DateFormatterRef) StringRef {
	return _CFDateFormatterGetFormat(formatter)
}

// Returns the locale object used to create the given date formatter object.
//
// Added in macOS .
// Returns the locale object used to create the given date formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetLocale(_:)
func CFDateFormatterGetLocale(formatter DateFormatterRef) LocaleRef {
	return _CFDateFormatterGetLocale(formatter)
}

// Returns the time style used to create the given date formatter object.
//
// Added in macOS .
// Returns the time style used to create the given date formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetTimeStyle(_:)
func CFDateFormatterGetTimeStyle(formatter DateFormatterRef) DateFormatterStyle {
	return _CFDateFormatterGetTimeStyle(formatter)
}

// Returns the type identifier for CFDateFormatter.
//
// Added in macOS .
// Returns the type identifier for CFDateFormatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetTypeID()
func CFDateFormatterGetTypeID() TypeID {
	return _CFDateFormatterGetTypeID()
}

// Sets the format string of the given date formatter to the specified value.
//
// Added in macOS .
// Sets the format string of the given date formatter to the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterSetFormat(_:_:)
func CFDateFormatterSetFormat(formatter DateFormatterRef, formatString StringRef) {
	_CFDateFormatterSetFormat(formatter, formatString)
}

// Sets a date formatter property using a key-value pair.
//
// Added in macOS .
// Sets a date formatter property using a key-value pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterSetProperty(_:_:_:)
func CFDateFormatterSetProperty(formatter DateFormatterRef, key StringRef, value TypeRef) {
	_CFDateFormatterSetProperty(formatter, key, value)
}

// Returns a object’s absolute time.
//
// Added in macOS .
// Returns a object’s absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateGetAbsoluteTime(_:)
func CFDateGetAbsoluteTime(theDate DateRef) AbsoluteTime {
	return _CFDateGetAbsoluteTime(theDate)
}

// Returns the number of elapsed seconds between the given objects.
//
// Added in macOS .
// Returns the number of elapsed seconds between the given objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateGetTimeIntervalSinceDate(_:_:)
func CFDateGetTimeIntervalSinceDate(theDate DateRef, otherDate DateRef) TimeInterval {
	return _CFDateGetTimeIntervalSinceDate(theDate, otherDate)
}

// Returns the type identifier for the opaque type.
//
// Added in macOS .
// Returns the type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateGetTypeID()
func CFDateGetTypeID() TypeID {
	return _CFDateGetTypeID()
}

// Adds a key-value pair to a dictionary if the specified key is not already present.
//
// Added in macOS .
// Adds a key-value pair to a dictionary if the specified key is not already present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryAddValue(_:_:_:)
func CFDictionaryAddValue(theDict MutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) {
	_CFDictionaryAddValue(theDict, key, value)
}

// Calls a function once for each key-value pair in a dictionary.
//
// Added in macOS .
// Calls a function once for each key-value pair in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryApplyFunction(_:_:_:)
func CFDictionaryApplyFunction(theDict DictionaryRef, applier DictionaryApplierFunction, context unsafe.Pointer) {
	_CFDictionaryApplyFunction(theDict, applier, context)
}

// Returns a Boolean value that indicates whether a given key is in a dictionary.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a given key is in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryContainsKey(_:_:)
func CFDictionaryContainsKey(theDict DictionaryRef, key unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryContainsKey(theDict, key)
}

// Returns a Boolean value that indicates whether a given value is in a dictionary.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a given value is in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryContainsValue(_:_:)
func CFDictionaryContainsValue(theDict DictionaryRef, value unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryContainsValue(theDict, value)
}

// Creates an immutable dictionary containing the specified key-value pairs.
//
// Added in macOS .
// Creates an immutable dictionary containing the specified key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreate(_:_:_:_:_:_:)
func CFDictionaryCreate(allocator AllocatorRef, keys unsafe.Pointer, values unsafe.Pointer, numValues Index, keyCallBacks unsafe.Pointer, valueCallBacks unsafe.Pointer) DictionaryRef {
	return _CFDictionaryCreate(allocator, keys, values, numValues, keyCallBacks, valueCallBacks)
}

// Creates and returns a new immutable dictionary with the key-value pairs of another dictionary.
//
// Added in macOS .
// Creates and returns a new immutable dictionary with the key-value pairs of another dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateCopy(_:_:)
func CFDictionaryCreateCopy(allocator AllocatorRef, theDict DictionaryRef) DictionaryRef {
	return _CFDictionaryCreateCopy(allocator, theDict)
}

// Creates a new mutable dictionary.
//
// Added in macOS .
// Creates a new mutable dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateMutable(_:_:_:_:)
func CFDictionaryCreateMutable(allocator AllocatorRef, capacity Index, keyCallBacks unsafe.Pointer, valueCallBacks unsafe.Pointer) MutableDictionaryRef {
	return _CFDictionaryCreateMutable(allocator, capacity, keyCallBacks, valueCallBacks)
}

// Creates a new mutable dictionary with the key-value pairs from another dictionary.
//
// Added in macOS .
// Creates a new mutable dictionary with the key-value pairs from another dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateMutableCopy(_:_:_:)
func CFDictionaryCreateMutableCopy(allocator AllocatorRef, capacity Index, theDict DictionaryRef) MutableDictionaryRef {
	return _CFDictionaryCreateMutableCopy(allocator, capacity, theDict)
}

// Returns the number of key-value pairs in a dictionary.
//
// Added in macOS .
// Returns the number of key-value pairs in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCount(_:)
func CFDictionaryGetCount(theDict DictionaryRef) Index {
	return _CFDictionaryGetCount(theDict)
}

// Returns the number of times a key occurs in a dictionary.
//
// Added in macOS .
// Returns the number of times a key occurs in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCountOfKey(_:_:)
func CFDictionaryGetCountOfKey(theDict DictionaryRef, key unsafe.Pointer) Index {
	return _CFDictionaryGetCountOfKey(theDict, key)
}

// Counts the number of times a given value occurs in the dictionary.
//
// Added in macOS .
// Counts the number of times a given value occurs in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCountOfValue(_:_:)
func CFDictionaryGetCountOfValue(theDict DictionaryRef, value unsafe.Pointer) Index {
	return _CFDictionaryGetCountOfValue(theDict, value)
}

// Fills two buffers with the keys and values from a dictionary.
//
// Added in macOS .
// Fills two buffers with the keys and values from a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetKeysAndValues(_:_:_:)
func CFDictionaryGetKeysAndValues(theDict DictionaryRef, keys unsafe.Pointer, values unsafe.Pointer) {
	_CFDictionaryGetKeysAndValues(theDict, keys, values)
}

// Returns the type identifier for the CFDictionary opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFDictionary opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetTypeID()
func CFDictionaryGetTypeID() TypeID {
	return _CFDictionaryGetTypeID()
}

// Returns the value associated with a given key.
//
// Added in macOS .
// Returns the value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetValue(_:_:)
func CFDictionaryGetValue(theDict DictionaryRef, key unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetValue(theDict, key)
}

// Returns a Boolean value that indicates whether a given value for a given key is in a dictionary, and returns that value indirectly if it exists.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a given value for a given key is in a dictionary, and returns that value indirectly if it exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetValueIfPresent(_:_:_:)
func CFDictionaryGetValueIfPresent(theDict DictionaryRef, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetValueIfPresent(theDict, key, value)
}

// Removes all the key-value pairs from a dictionary, making it empty.
//
// Added in macOS .
// Removes all the key-value pairs from a dictionary, making it empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryRemoveAllValues(_:)
func CFDictionaryRemoveAllValues(theDict MutableDictionaryRef) {
	_CFDictionaryRemoveAllValues(theDict)
}

// Removes a key-value pair.
//
// Added in macOS .
// Removes a key-value pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryRemoveValue(_:_:)
func CFDictionaryRemoveValue(theDict MutableDictionaryRef, key unsafe.Pointer) {
	_CFDictionaryRemoveValue(theDict, key)
}

// Replaces a value corresponding to a given key.
//
// Added in macOS .
// Replaces a value corresponding to a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryReplaceValue(_:_:_:)
func CFDictionaryReplaceValue(theDict MutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) {
	_CFDictionaryReplaceValue(theDict, key, value)
}

// Sets the value corresponding to a given key.
//
// Added in macOS .
// Sets the value corresponding to a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionarySetValue(_:_:_:)
func CFDictionarySetValue(theDict MutableDictionaryRef, key unsafe.Pointer, value unsafe.Pointer) {
	_CFDictionarySetValue(theDict, key, value)
}

// Determines whether two Core Foundation objects are considered equal.
//
// Added in macOS .
// Determines whether two Core Foundation objects are considered equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFEqual(_:_:)
func CFEqual(cf1 TypeRef, cf2 TypeRef) unsafe.Pointer {
	return _CFEqual(cf1, cf2)
}

// Returns a human-presentable description for a given error.
//
// Added in macOS 10.5.
// Returns a human-presentable description for a given error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyDescription(_:)
func CFErrorCopyDescription(err ErrorRef) StringRef {
	return _CFErrorCopyDescription(err)
}

// Returns a human-presentable failure reason for a given error.
//
// Added in macOS 10.5.
// Returns a human-presentable failure reason for a given error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyFailureReason(_:)
func CFErrorCopyFailureReason(err ErrorRef) StringRef {
	return _CFErrorCopyFailureReason(err)
}

// Returns a human presentable recovery suggestion for a given error.
//
// Added in macOS 10.5.
// Returns a human presentable recovery suggestion for a given error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyRecoverySuggestion(_:)
func CFErrorCopyRecoverySuggestion(err ErrorRef) StringRef {
	return _CFErrorCopyRecoverySuggestion(err)
}

// Returns the user info dictionary for a given CFError.
//
// Added in macOS 10.5.
// Returns the user info dictionary for a given CFError.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCopyUserInfo(_:)
func CFErrorCopyUserInfo(err ErrorRef) DictionaryRef {
	return _CFErrorCopyUserInfo(err)
}

// Creates a new CFError object.
//
// Added in macOS 10.5.
// Creates a new CFError object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCreate(_:_:_:_:)
func CFErrorCreate(allocator AllocatorRef, domain ErrorDomain, code Index, userInfo DictionaryRef) ErrorRef {
	return _CFErrorCreate(allocator, domain, code, userInfo)
}

// Creates a new CFError object using given keys and values to create the user info dictionary.
//
// Added in macOS 10.5.
// Creates a new CFError object using given keys and values to create the user info dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCreateWithUserInfoKeysAndValues(_:_:_:_:_:_:)
func CFErrorCreateWithUserInfoKeysAndValues(allocator AllocatorRef, domain ErrorDomain, code Index, userInfoKeys unsafe.Pointer, userInfoValues unsafe.Pointer, numUserInfoValues Index) ErrorRef {
	return _CFErrorCreateWithUserInfoKeysAndValues(allocator, domain, code, userInfoKeys, userInfoValues, numUserInfoValues)
}

// Returns the error code for a given CFError.
//
// Added in macOS 10.5.
// Returns the error code for a given CFError.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetCode(_:)
func CFErrorGetCode(err ErrorRef) Index {
	return _CFErrorGetCode(err)
}

// Returns the error domain for a given CFError.
//
// Added in macOS 10.5.
// Returns the error domain for a given CFError.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetDomain(_:)
func CFErrorGetDomain(err ErrorRef) ErrorDomain {
	return _CFErrorGetDomain(err)
}

// Returns the type identifier for the CFError opaque type.
//
// Added in macOS 10.5.
// Returns the type identifier for the CFError opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorGetTypeID()
func CFErrorGetTypeID() TypeID {
	return _CFErrorGetTypeID()
}

// Creates a new CFFileDescriptor.
//
// Added in macOS 10.5.
// Creates a new CFFileDescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorCreate(_:_:_:_:_:)
func CFFileDescriptorCreate(allocator AllocatorRef, fd FileDescriptorNativeDescriptor, closeOnInvalidate unsafe.Pointer, callout FileDescriptorCallBack, context unsafe.Pointer) FileDescriptorRef {
	return _CFFileDescriptorCreate(allocator, fd, closeOnInvalidate, callout, context)
}

// Creates a new runloop source for a given CFFileDescriptor.
//
// Added in macOS 10.5.
// Creates a new runloop source for a given CFFileDescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorCreateRunLoopSource(_:_:_:)
func CFFileDescriptorCreateRunLoopSource(allocator AllocatorRef, f FileDescriptorRef, order Index) RunLoopSourceRef {
	return _CFFileDescriptorCreateRunLoopSource(allocator, f, order)
}

// Disables callbacks for a given CFFileDescriptor.
//
// Added in macOS 10.5.
// Disables callbacks for a given CFFileDescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorDisableCallBacks(_:_:)
func CFFileDescriptorDisableCallBacks(f FileDescriptorRef, callBackTypes OptionFlags) {
	_CFFileDescriptorDisableCallBacks(f, callBackTypes)
}

// Enables callbacks for a given CFFileDescriptor.
//
// Added in macOS 10.5.
// Enables callbacks for a given CFFileDescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorEnableCallBacks(_:_:)
func CFFileDescriptorEnableCallBacks(f FileDescriptorRef, callBackTypes OptionFlags) {
	_CFFileDescriptorEnableCallBacks(f, callBackTypes)
}

// Gets the context for a given CFFileDescriptor.
//
// Added in macOS 10.5.
// Gets the context for a given CFFileDescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetContext(_:_:)
func CFFileDescriptorGetContext(f FileDescriptorRef, context unsafe.Pointer) {
	_CFFileDescriptorGetContext(f, context)
}

// Returns the native file descriptor for a given CFFileDescriptor.
//
// Added in macOS 10.5.
// Returns the native file descriptor for a given CFFileDescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetNativeDescriptor(_:)
func CFFileDescriptorGetNativeDescriptor(f FileDescriptorRef) FileDescriptorNativeDescriptor {
	return _CFFileDescriptorGetNativeDescriptor(f)
}

// Returns the type identifier for the CFFileDescriptor opaque type.
//
// Added in macOS 10.5.
// Returns the type identifier for the CFFileDescriptor opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetTypeID()
func CFFileDescriptorGetTypeID() TypeID {
	return _CFFileDescriptorGetTypeID()
}

// Invalidates a CFFileDescriptor object.
//
// Added in macOS 10.5.
// Invalidates a CFFileDescriptor object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorInvalidate(_:)
func CFFileDescriptorInvalidate(f FileDescriptorRef) {
	_CFFileDescriptorInvalidate(f)
}

// Returns a Boolean value that indicates whether the native file descriptor for a given CFFileDescriptor is valid.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether the native file descriptor for a given CFFileDescriptor is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorIsValid(_:)
func CFFileDescriptorIsValid(f FileDescriptorRef) unsafe.Pointer {
	return _CFFileDescriptorIsValid(f)
}

// Clears properties from a object.
//
// Added in macOS 10.8.
// Clears properties from a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearProperties(_:_:)
func CFFileSecurityClearProperties(fileSec FileSecurityRef, clearPropertyMask FileSecurityClearOptions) unsafe.Pointer {
	return _CFFileSecurityClearProperties(fileSec, clearPropertyMask)
}

// Copies the access control list associated with a object.
//
// Added in macOS 10.7.
// Copies the access control list associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyAccessControlList(_:_:)
func CFFileSecurityCopyAccessControlList(fileSec FileSecurityRef, accessControlList unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityCopyAccessControlList(fileSec, accessControlList)
}

// Copies the group UUID associated with a object.
//
// Added in macOS 10.7.
// Copies the group UUID associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyGroupUUID(_:_:)
func CFFileSecurityCopyGroupUUID(fileSec FileSecurityRef, groupUUID unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityCopyGroupUUID(fileSec, groupUUID)
}

// Copies the owner UUID associated with a object.
//
// Added in macOS 10.7.
// Copies the owner UUID associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCopyOwnerUUID(_:_:)
func CFFileSecurityCopyOwnerUUID(fileSec FileSecurityRef, ownerUUID unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityCopyOwnerUUID(fileSec, ownerUUID)
}

// Creates a object.
//
// Added in macOS 10.7.
// Creates a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCreate(_:)
func CFFileSecurityCreate(allocator AllocatorRef) FileSecurityRef {
	return _CFFileSecurityCreate(allocator)
}

// Creates a copy of a object.
//
// Added in macOS 10.7.
// Creates a copy of a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityCreateCopy(_:_:)
func CFFileSecurityCreateCopy(allocator AllocatorRef, fileSec FileSecurityRef) FileSecurityRef {
	return _CFFileSecurityCreateCopy(allocator, fileSec)
}

// Gets the group ID associated with a object
//
// Added in macOS 10.7.
// Gets the group ID associated with a object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetGroup(_:_:)
func CFFileSecurityGetGroup(fileSec FileSecurityRef, group unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityGetGroup(fileSec, group)
}

// Gets the file mode associated with a object.
//
// Added in macOS 10.7.
// Gets the file mode associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetMode(_:_:)
func CFFileSecurityGetMode(fileSec FileSecurityRef, mode unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityGetMode(fileSec, mode)
}

// Gets the owner ID associated with a object.
//
// Added in macOS 10.7.
// Gets the owner ID associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetOwner(_:_:)
func CFFileSecurityGetOwner(fileSec FileSecurityRef, owner unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecurityGetOwner(fileSec, owner)
}

// Returns the type identifier for the opaque type.
//
// Added in macOS 10.7.
// Returns the type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityGetTypeID()
func CFFileSecurityGetTypeID() TypeID {
	return _CFFileSecurityGetTypeID()
}

// Sets the access control list associated with a object.
//
// Added in macOS 10.7.
// Sets the access control list associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetAccessControlList(_:_:)
func CFFileSecuritySetAccessControlList(fileSec FileSecurityRef, accessControlList unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecuritySetAccessControlList(fileSec, accessControlList)
}

// Sets the group ID associated with a object.
//
// Added in macOS 10.7.
// Sets the group ID associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetGroup(_:_:)
func CFFileSecuritySetGroup(fileSec FileSecurityRef, group unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecuritySetGroup(fileSec, group)
}

// Sets the group UUID associated with a object.
//
// Added in macOS 10.7.
// Sets the group UUID associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetGroupUUID(_:_:)
func CFFileSecuritySetGroupUUID(fileSec FileSecurityRef, groupUUID UUIDRef) unsafe.Pointer {
	return _CFFileSecuritySetGroupUUID(fileSec, groupUUID)
}

// Sets the file mode associated with a object.
//
// Added in macOS 10.7.
// Sets the file mode associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetMode(_:_:)
func CFFileSecuritySetMode(fileSec FileSecurityRef, mode unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecuritySetMode(fileSec, mode)
}

// Sets the owner ID associated with a object.
//
// Added in macOS 10.7.
// Sets the owner ID associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetOwner(_:_:)
func CFFileSecuritySetOwner(fileSec FileSecurityRef, owner unsafe.Pointer) unsafe.Pointer {
	return _CFFileSecuritySetOwner(fileSec, owner)
}

// Sets the owner UUID associated with a object.
//
// Added in macOS 10.7.
// Sets the owner UUID associated with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecuritySetOwnerUUID(_:_:)
func CFFileSecuritySetOwnerUUID(fileSec FileSecurityRef, ownerUUID UUIDRef) unsafe.Pointer {
	return _CFFileSecuritySetOwnerUUID(fileSec, ownerUUID)
}

// Returns the allocator used to allocate a Core Foundation object.
//
// Added in macOS .
// Returns the allocator used to allocate a Core Foundation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGetAllocator(_:)
func CFGetAllocator(cf TypeRef) AllocatorRef {
	return _CFGetAllocator(cf)
}

// Returns the reference count of a Core Foundation object.
//
// Added in macOS .
// Returns the reference count of a Core Foundation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGetRetainCount(_:)
func CFGetRetainCount(cf TypeRef) Index {
	return _CFGetRetainCount(cf)
}

// Returns the unique identifier of an opaque type to which a Core Foundation object belongs.
//
// Added in macOS .
// Returns the unique identifier of an opaque type to which a Core Foundation object belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGetTypeID(_:)
func CFGetTypeID(cf TypeRef) TypeID {
	return _CFGetTypeID(cf)
}

// Converts a Gregorian date value into an absolute time value.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
// Converts a Gregorian date value into an absolute time value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianDateGetAbsoluteTime(_:_:)
func CFGregorianDateGetAbsoluteTime(gdate GregorianDate, tz TimeZoneRef) AbsoluteTime {
	return _CFGregorianDateGetAbsoluteTime(gdate, tz)
}

// Checks the specified fields of a CFGregorianDate structure for valid values.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.4.
// Checks the specified fields of a CFGregorianDate structure for valid values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFGregorianDateIsValid(_:_:)
func CFGregorianDateIsValid(gdate GregorianDate, unitFlags OptionFlags) unsafe.Pointer {
	return _CFGregorianDateIsValid(gdate, unitFlags)
}

// Returns a code that can be used to identify an object in a hashing structure.
//
// Added in macOS .
// Returns a code that can be used to identify an object in a hashing structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFHash(_:)
func CFHash(cf TypeRef) HashCode {
	return _CFHash(cf)
}

// Returns an array of CFString objects that represents all locales for which locale data is available.
//
// Added in macOS .
// Returns an array of CFString objects that represents all locales for which locale data is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyAvailableLocaleIdentifiers()
func CFLocaleCopyAvailableLocaleIdentifiers() ArrayRef {
	return _CFLocaleCopyAvailableLocaleIdentifiers()
}

// Returns an array of strings that represents ISO currency codes for currencies in common use.
//
// Added in macOS 10.5.
// Returns an array of strings that represents ISO currency codes for currencies in common use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyCommonISOCurrencyCodes()
func CFLocaleCopyCommonISOCurrencyCodes() ArrayRef {
	return _CFLocaleCopyCommonISOCurrencyCodes()
}

// Returns a copy of the logical locale for the current user.
//
// Added in macOS .
// Returns a copy of the logical locale for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyCurrent()
func CFLocaleCopyCurrent() LocaleRef {
	return _CFLocaleCopyCurrent()
}

// Returns the display name for the given value.
//
// Added in macOS .
// Returns the display name for the given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyDisplayNameForPropertyValue(_:_:_:)
func CFLocaleCopyDisplayNameForPropertyValue(displayLocale LocaleRef, key LocaleKey, value StringRef) StringRef {
	return _CFLocaleCopyDisplayNameForPropertyValue(displayLocale, key, value)
}

// Returns an array of CFString objects that represents all known legal ISO country codes.
//
// Added in macOS .
// Returns an array of CFString objects that represents all known legal ISO country codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOCountryCodes()
func CFLocaleCopyISOCountryCodes() ArrayRef {
	return _CFLocaleCopyISOCountryCodes()
}

// Returns an array of CFString objects that represents all known legal ISO currency codes.
//
// Added in macOS .
// Returns an array of CFString objects that represents all known legal ISO currency codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOCurrencyCodes()
func CFLocaleCopyISOCurrencyCodes() ArrayRef {
	return _CFLocaleCopyISOCurrencyCodes()
}

// Returns an array of CFString objects that represents all known legal ISO language codes.
//
// Added in macOS .
// Returns an array of CFString objects that represents all known legal ISO language codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOLanguageCodes()
func CFLocaleCopyISOLanguageCodes() ArrayRef {
	return _CFLocaleCopyISOLanguageCodes()
}

// Returns the array of canonicalized language IDs that the user prefers.
//
// Added in macOS 10.5.
// Returns the array of canonicalized language IDs that the user prefers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyPreferredLanguages()
func CFLocaleCopyPreferredLanguages() ArrayRef {
	return _CFLocaleCopyPreferredLanguages()
}

// Creates a locale for the given arbitrary locale identifier.
//
// Added in macOS .
// Creates a locale for the given arbitrary locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreate(_:_:)
func CFLocaleCreate(allocator AllocatorRef, localeIdentifier LocaleIdentifier) LocaleRef {
	return _CFLocaleCreate(allocator, localeIdentifier)
}

// Returns a canonical language identifier by mapping an arbitrary locale identification string to the canonical identifier
//
// Added in macOS .
// Returns a canonical language identifier by mapping an arbitrary locale identification string to the canonical identifier
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLanguageIdentifierFromString(_:_:)
func CFLocaleCreateCanonicalLanguageIdentifierFromString(allocator AllocatorRef, localeIdentifier StringRef) LocaleIdentifier {
	return _CFLocaleCreateCanonicalLanguageIdentifierFromString(allocator, localeIdentifier)
}

// Returns a canonical locale identifier from given language and region codes.
//
// Added in macOS .
// Returns a canonical locale identifier from given language and region codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(_:_:_:)
func CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator AllocatorRef, lcode unsafe.Pointer, rcode unsafe.Pointer) LocaleIdentifier {
	return _CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator, lcode, rcode)
}

// Returns a canonical locale identifier by mapping an arbitrary locale identification string to the canonical identifier.
//
// Added in macOS .
// Returns a canonical locale identifier by mapping an arbitrary locale identification string to the canonical identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLocaleIdentifierFromString(_:_:)
func CFLocaleCreateCanonicalLocaleIdentifierFromString(allocator AllocatorRef, localeIdentifier StringRef) LocaleIdentifier {
	return _CFLocaleCreateCanonicalLocaleIdentifierFromString(allocator, localeIdentifier)
}

// Returns a dictionary containing the result from parsing a locale ID consisting of language, script, country or region, variant, and keyword/value pairs.
//
// Added in macOS .
// Returns a dictionary containing the result from parsing a locale ID consisting of language, script, country or region, variant, and keyword/value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateComponentsFromLocaleIdentifier(_:_:)
func CFLocaleCreateComponentsFromLocaleIdentifier(allocator AllocatorRef, localeID LocaleIdentifier) DictionaryRef {
	return _CFLocaleCreateComponentsFromLocaleIdentifier(allocator, localeID)
}

// Returns a copy of a locale.
//
// Added in macOS .
// Returns a copy of a locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCopy(_:_:)
func CFLocaleCreateCopy(allocator AllocatorRef, locale LocaleRef) LocaleRef {
	return _CFLocaleCreateCopy(allocator, locale)
}

// Returns a locale identifier consisting of language, script, country or region, variant, and keyword/value pairs derived from a dictionary containing the source information.
//
// Added in macOS .
// Returns a locale identifier consisting of language, script, country or region, variant, and keyword/value pairs derived from a dictionary containing the source information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateLocaleIdentifierFromComponents(_:_:)
func CFLocaleCreateLocaleIdentifierFromComponents(allocator AllocatorRef, dictionary DictionaryRef) LocaleIdentifier {
	return _CFLocaleCreateLocaleIdentifierFromComponents(allocator, dictionary)
}

// Returns a locale identifier from a Windows locale code.
//
// Added in macOS 10.6.
// Returns a locale identifier from a Windows locale code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(_:_:)
func CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator AllocatorRef, lcid uint32) LocaleIdentifier {
	return _CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator, lcid)
}

// Returns the given locale’s identifier.
//
// Added in macOS .
// Returns the given locale’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetIdentifier(_:)
func CFLocaleGetIdentifier(locale LocaleRef) LocaleIdentifier {
	return _CFLocaleGetIdentifier(locale)
}

// Returns the character direction for the specified ISO language code.
//
// Added in macOS 10.6.
// Returns the character direction for the specified ISO language code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetLanguageCharacterDirection(_:)
func CFLocaleGetLanguageCharacterDirection(isoLangCode StringRef) LocaleLanguageDirection {
	return _CFLocaleGetLanguageCharacterDirection(isoLangCode)
}

// Returns the line direction for the specified ISO language code.
//
// Added in macOS 10.6.
// Returns the line direction for the specified ISO language code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetLanguageLineDirection(_:)
func CFLocaleGetLanguageLineDirection(isoLangCode StringRef) LocaleLanguageDirection {
	return _CFLocaleGetLanguageLineDirection(isoLangCode)
}

// Returns the root, canonical locale.
//
// Added in macOS .
// Returns the root, canonical locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetSystem()
func CFLocaleGetSystem() LocaleRef {
	return _CFLocaleGetSystem()
}

// Returns the type identifier for the CFLocale opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFLocale opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetTypeID()
func CFLocaleGetTypeID() TypeID {
	return _CFLocaleGetTypeID()
}

// Returns the corresponding value for the given key of a locale’s key-value pair.
//
// Added in macOS .
// Returns the corresponding value for the given key of a locale’s key-value pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetValue(_:_:)
func CFLocaleGetValue(locale LocaleRef, key LocaleKey) TypeRef {
	return _CFLocaleGetValue(locale, key)
}

// Returns a Windows locale code from the locale identifier.
//
// Added in macOS 10.6.
// Returns a Windows locale code from the locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(_:)
func CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier LocaleIdentifier) uint32 {
	return _CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier)
}

// Creates a CFMachPort object with a new Mach port.
//
// Added in macOS .
// Creates a CFMachPort object with a new Mach port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreate(_:_:_:_:)
func CFMachPortCreate(allocator AllocatorRef, callout MachPortCallBack, context unsafe.Pointer, shouldFreeInfo unsafe.Pointer) MachPortRef {
	return _CFMachPortCreate(allocator, callout, context, shouldFreeInfo)
}

// Creates a CFRunLoopSource object for a CFMachPort object.
//
// Added in macOS .
// Creates a CFRunLoopSource object for a CFMachPort object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreateRunLoopSource(_:_:_:)
func CFMachPortCreateRunLoopSource(allocator AllocatorRef, port MachPortRef, order Index) RunLoopSourceRef {
	return _CFMachPortCreateRunLoopSource(allocator, port, order)
}

// Creates a CFMachPort object for a pre-existing native Mach port.
//
// Added in macOS .
// Creates a CFMachPort object for a pre-existing native Mach port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCreateWithPort(_:_:_:_:_:)
func CFMachPortCreateWithPort(allocator AllocatorRef, portNum unsafe.Pointer, callout MachPortCallBack, context unsafe.Pointer, shouldFreeInfo unsafe.Pointer) MachPortRef {
	return _CFMachPortCreateWithPort(allocator, portNum, callout, context, shouldFreeInfo)
}

// Returns the context information for a CFMachPort object.
//
// Added in macOS .
// Returns the context information for a CFMachPort object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetContext(_:_:)
func CFMachPortGetContext(port MachPortRef, context unsafe.Pointer) {
	_CFMachPortGetContext(port, context)
}

// Returns the invalidation callback function for a CFMachPort object.
//
// Added in macOS .
// Returns the invalidation callback function for a CFMachPort object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetInvalidationCallBack(_:)
func CFMachPortGetInvalidationCallBack(port MachPortRef) MachPortInvalidationCallBack {
	return _CFMachPortGetInvalidationCallBack(port)
}

// Returns the native Mach port represented by a CFMachPort object.
//
// Added in macOS .
// Returns the native Mach port represented by a CFMachPort object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetPort(_:)
func CFMachPortGetPort(port MachPortRef) unsafe.Pointer {
	return _CFMachPortGetPort(port)
}

// Returns the type identifier for the CFMachPort opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFMachPort opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortGetTypeID()
func CFMachPortGetTypeID() TypeID {
	return _CFMachPortGetTypeID()
}

// Invalidates a CFMachPort object, stopping it from receiving any more messages.
//
// Added in macOS .
// Invalidates a CFMachPort object, stopping it from receiving any more messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortInvalidate(_:)
func CFMachPortInvalidate(port MachPortRef) {
	_CFMachPortInvalidate(port)
}

// Returns a Boolean value that indicates whether a CFMachPort object is valid and able to receive messages.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a CFMachPort object is valid and able to receive messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortIsValid(_:)
func CFMachPortIsValid(port MachPortRef) unsafe.Pointer {
	return _CFMachPortIsValid(port)
}

// Sets the callback function invoked when a CFMachPort object is invalidated.
//
// Added in macOS .
// Sets the callback function invoked when a CFMachPort object is invalidated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortSetInvalidationCallBack(_:_:)
func CFMachPortSetInvalidationCallBack(port MachPortRef, callout MachPortInvalidationCallBack) {
	_CFMachPortSetInvalidationCallBack(port, callout)
}

// Makes a newly-allocated Core Foundation object eligible for garbage collection.
//
// Added in macOS .
// Makes a newly-allocated Core Foundation object eligible for garbage collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMakeCollectable
func CFMakeCollectable(cf TypeRef) TypeRef {
	return _CFMakeCollectable(cf)
}

// Returns a local CFMessagePort object.
//
// Added in macOS .
// Returns a local CFMessagePort object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateLocal(_:_:_:_:_:)
func CFMessagePortCreateLocal(allocator AllocatorRef, name StringRef, callout MessagePortCallBack, context unsafe.Pointer, shouldFreeInfo unsafe.Pointer) MessagePortRef {
	return _CFMessagePortCreateLocal(allocator, name, callout, context, shouldFreeInfo)
}

// Returns a CFMessagePort object connected to a remote port.
//
// Added in macOS .
// Returns a CFMessagePort object connected to a remote port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateRemote(_:_:)
func CFMessagePortCreateRemote(allocator AllocatorRef, name StringRef) MessagePortRef {
	return _CFMessagePortCreateRemote(allocator, name)
}

// Creates a CFRunLoopSource object for a CFMessagePort object.
//
// Added in macOS .
// Creates a CFRunLoopSource object for a CFMessagePort object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCreateRunLoopSource(_:_:_:)
func CFMessagePortCreateRunLoopSource(allocator AllocatorRef, local MessagePortRef, order Index) RunLoopSourceRef {
	return _CFMessagePortCreateRunLoopSource(allocator, local, order)
}

// Returns the context information for a CFMessagePort object.
//
// Added in macOS .
// Returns the context information for a CFMessagePort object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetContext(_:_:)
func CFMessagePortGetContext(ms MessagePortRef, context unsafe.Pointer) {
	_CFMessagePortGetContext(ms, context)
}

// Returns the invalidation callback function for a CFMessagePort object.
//
// Added in macOS .
// Returns the invalidation callback function for a CFMessagePort object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetInvalidationCallBack(_:)
func CFMessagePortGetInvalidationCallBack(ms MessagePortRef) MessagePortInvalidationCallBack {
	return _CFMessagePortGetInvalidationCallBack(ms)
}

// Returns the name with which a CFMessagePort object is registered.
//
// Added in macOS .
// Returns the name with which a CFMessagePort object is registered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetName(_:)
func CFMessagePortGetName(ms MessagePortRef) StringRef {
	return _CFMessagePortGetName(ms)
}

// Returns the type identifier for the CFMessagePort opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFMessagePort opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortGetTypeID()
func CFMessagePortGetTypeID() TypeID {
	return _CFMessagePortGetTypeID()
}

// Invalidates a CFMessagePort object, stopping it from receiving or sending any more messages.
//
// Added in macOS .
// Invalidates a CFMessagePort object, stopping it from receiving or sending any more messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortInvalidate(_:)
func CFMessagePortInvalidate(ms MessagePortRef) {
	_CFMessagePortInvalidate(ms)
}

// Returns a Boolean value that indicates whether a CFMessagePort object represents a remote port.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a CFMessagePort object represents a remote port.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortIsRemote(_:)
func CFMessagePortIsRemote(ms MessagePortRef) unsafe.Pointer {
	return _CFMessagePortIsRemote(ms)
}

// Returns a Boolean value that indicates whether a CFMessagePort object is valid and able to send or receive messages.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a CFMessagePort object is valid and able to send or receive messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortIsValid(_:)
func CFMessagePortIsValid(ms MessagePortRef) unsafe.Pointer {
	return _CFMessagePortIsValid(ms)
}

// Sends a message to a remote CFMessagePort object.
//
// Added in macOS .
// Sends a message to a remote CFMessagePort object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSendRequest(_:_:_:_:_:_:_:)
func CFMessagePortSendRequest(remote MessagePortRef, msgid unsafe.Pointer, data DataRef, sendTimeout TimeInterval, rcvTimeout TimeInterval, replyMode StringRef, returnData unsafe.Pointer) unsafe.Pointer {
	return _CFMessagePortSendRequest(remote, msgid, data, sendTimeout, rcvTimeout, replyMode, returnData)
}

// Schedules callbacks for the specified message port on the specified dispatch queue.
//
// Added in macOS 10.6.
// Schedules callbacks for the specified message port on the specified dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetDispatchQueue(_:_:)
func CFMessagePortSetDispatchQueue(ms MessagePortRef, queue unsafe.Pointer) {
	_CFMessagePortSetDispatchQueue(ms, queue)
}

// Sets the callback function invoked when a CFMessagePort object is invalidated.
//
// Added in macOS .
// Sets the callback function invoked when a CFMessagePort object is invalidated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetInvalidationCallBack(_:_:)
func CFMessagePortSetInvalidationCallBack(ms MessagePortRef, callout MessagePortInvalidationCallBack) {
	_CFMessagePortSetInvalidationCallBack(ms, callout)
}

// Sets the name of a local CFMessagePort object.
//
// Added in macOS .
// Sets the name of a local CFMessagePort object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortSetName(_:_:)
func CFMessagePortSetName(ms MessagePortRef, newName StringRef) unsafe.Pointer {
	return _CFMessagePortSetName(ms, newName)
}

// Registers an observer to receive notifications.
//
// Added in macOS .
// Registers an observer to receive notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterAddObserver(_:_:_:_:_:_:)
func CFNotificationCenterAddObserver(center NotificationCenterRef, observer unsafe.Pointer, callBack NotificationCallback, name StringRef, object unsafe.Pointer, suspensionBehavior NotificationSuspensionBehavior) {
	_CFNotificationCenterAddObserver(center, observer, callBack, name, object, suspensionBehavior)
}

// Returns the application’s Darwin notification center.
//
// Added in macOS .
// Returns the application’s Darwin notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetDarwinNotifyCenter()
func CFNotificationCenterGetDarwinNotifyCenter() NotificationCenterRef {
	return _CFNotificationCenterGetDarwinNotifyCenter()
}

// Returns the application’s distributed notification center.
//
// Added in macOS .
// Returns the application’s distributed notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetDistributedCenter()
func CFNotificationCenterGetDistributedCenter() NotificationCenterRef {
	return _CFNotificationCenterGetDistributedCenter()
}

// Returns the application’s local notification center.
//
// Added in macOS .
// Returns the application’s local notification center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetLocalCenter()
func CFNotificationCenterGetLocalCenter() NotificationCenterRef {
	return _CFNotificationCenterGetLocalCenter()
}

// Returns the type identifier for the CFNotificationCenter opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFNotificationCenter opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterGetTypeID()
func CFNotificationCenterGetTypeID() TypeID {
	return _CFNotificationCenterGetTypeID()
}

// Posts a notification for an object.
//
// Added in macOS .
// Posts a notification for an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterPostNotification(_:_:_:_:_:)
func CFNotificationCenterPostNotification(center NotificationCenterRef, name NotificationName, object unsafe.Pointer, userInfo DictionaryRef, deliverImmediately unsafe.Pointer) {
	_CFNotificationCenterPostNotification(center, name, object, userInfo, deliverImmediately)
}

// Posts a notification for an object using specified options.
//
// Added in macOS .
// Posts a notification for an object using specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterPostNotificationWithOptions(_:_:_:_:_:)
func CFNotificationCenterPostNotificationWithOptions(center NotificationCenterRef, name NotificationName, object unsafe.Pointer, userInfo DictionaryRef, options OptionFlags) {
	_CFNotificationCenterPostNotificationWithOptions(center, name, object, userInfo, options)
}

// Stops an observer from receiving any notifications from any object.
//
// Added in macOS .
// Stops an observer from receiving any notifications from any object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterRemoveEveryObserver(_:_:)
func CFNotificationCenterRemoveEveryObserver(center NotificationCenterRef, observer unsafe.Pointer) {
	_CFNotificationCenterRemoveEveryObserver(center, observer)
}

// Stops an observer from receiving certain notifications.
//
// Added in macOS .
// Stops an observer from receiving certain notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenterRemoveObserver(_:_:_:_:)
func CFNotificationCenterRemoveObserver(center NotificationCenterRef, observer unsafe.Pointer, name NotificationName, object unsafe.Pointer) {
	_CFNotificationCenterRemoveObserver(center, observer, name, object)
}

// Returns the type identifier for the CFNull opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFNull opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNullGetTypeID()
func CFNullGetTypeID() TypeID {
	return _CFNullGetTypeID()
}

// Compares two CFNumber objects and returns a comparison result.
//
// Added in macOS .
// Compares two CFNumber objects and returns a comparison result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberCompare(_:_:_:)
func CFNumberCompare(number NumberRef, otherNumber NumberRef, context unsafe.Pointer) ComparisonResult {
	return _CFNumberCompare(number, otherNumber, context)
}

// Creates a CFNumber object using a specified value.
//
// Added in macOS .
// Creates a CFNumber object using a specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberCreate(_:_:_:)
func CFNumberCreate(allocator AllocatorRef, theType NumberType, valuePtr unsafe.Pointer) NumberRef {
	return _CFNumberCreate(allocator, theType, valuePtr)
}

// Returns a copy of a number formatter’s value for a given key.
//
// Added in macOS .
// Returns a copy of a number formatter’s value for a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCopyProperty(_:_:)
func CFNumberFormatterCopyProperty(formatter NumberFormatterRef, key NumberFormatterKey) TypeRef {
	return _CFNumberFormatterCopyProperty(formatter, key)
}

// Creates a new CFNumberFormatter object, localized to the given locale, which will format numbers to the given style.
//
// Added in macOS .
// Creates a new CFNumberFormatter object, localized to the given locale, which will format numbers to the given style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreate(_:_:_:)
func CFNumberFormatterCreate(allocator AllocatorRef, locale LocaleRef, style NumberFormatterStyle) NumberFormatterRef {
	return _CFNumberFormatterCreate(allocator, locale, style)
}

// Returns a number object representing a given string.
//
// Added in macOS .
// Returns a number object representing a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateNumberFromString(_:_:_:_:_:)
func CFNumberFormatterCreateNumberFromString(allocator AllocatorRef, formatter NumberFormatterRef, string_ StringRef, rangep unsafe.Pointer, options OptionFlags) NumberRef {
	return _CFNumberFormatterCreateNumberFromString(allocator, formatter, string_, rangep, options)
}

// Returns a string representation of the given number using the specified number formatter.
//
// Added in macOS .
// Returns a string representation of the given number using the specified number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateStringWithNumber(_:_:_:)
func CFNumberFormatterCreateStringWithNumber(allocator AllocatorRef, formatter NumberFormatterRef, number NumberRef) StringRef {
	return _CFNumberFormatterCreateStringWithNumber(allocator, formatter, number)
}

// Returns a string representation of the given number or value using the specified number formatter.
//
// Added in macOS .
// Returns a string representation of the given number or value using the specified number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterCreateStringWithValue(_:_:_:_:)
func CFNumberFormatterCreateStringWithValue(allocator AllocatorRef, formatter NumberFormatterRef, numberType NumberType, valuePtr unsafe.Pointer) StringRef {
	return _CFNumberFormatterCreateStringWithValue(allocator, formatter, numberType, valuePtr)
}

// Returns the number of fraction digits that should be displayed, and the rounding increment, for a given currency.
//
// Added in macOS .
// Returns the number of fraction digits that should be displayed, and the rounding increment, for a given currency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetDecimalInfoForCurrencyCode(_:_:_:)
func CFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode StringRef, defaultFractionDigits unsafe.Pointer, roundingIncrement []float64) unsafe.Pointer {
	return _CFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode, defaultFractionDigits, roundingIncrement)
}

// Returns a format string for the given number formatter object.
//
// Added in macOS .
// Returns a format string for the given number formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetFormat(_:)
func CFNumberFormatterGetFormat(formatter NumberFormatterRef) StringRef {
	return _CFNumberFormatterGetFormat(formatter)
}

// Returns the locale object used to create the given number formatter object.
//
// Added in macOS .
// Returns the locale object used to create the given number formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetLocale(_:)
func CFNumberFormatterGetLocale(formatter NumberFormatterRef) LocaleRef {
	return _CFNumberFormatterGetLocale(formatter)
}

// Returns the number style used to create the given number formatter object.
//
// Added in macOS .
// Returns the number style used to create the given number formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetStyle(_:)
func CFNumberFormatterGetStyle(formatter NumberFormatterRef) NumberFormatterStyle {
	return _CFNumberFormatterGetStyle(formatter)
}

// Returns the type identifier for the opaque type.
//
// Added in macOS .
// Returns the type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetTypeID()
func CFNumberFormatterGetTypeID() TypeID {
	return _CFNumberFormatterGetTypeID()
}

// Returns a number or value representing a given string.
//
// Added in macOS .
// Returns a number or value representing a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterGetValueFromString(_:_:_:_:_:)
func CFNumberFormatterGetValueFromString(formatter NumberFormatterRef, string_ StringRef, rangep unsafe.Pointer, numberType NumberType, valuePtr unsafe.Pointer) unsafe.Pointer {
	return _CFNumberFormatterGetValueFromString(formatter, string_, rangep, numberType, valuePtr)
}

// Sets the format string of a number formatter.
//
// Added in macOS .
// Sets the format string of a number formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterSetFormat(_:_:)
func CFNumberFormatterSetFormat(formatter NumberFormatterRef, formatString StringRef) {
	_CFNumberFormatterSetFormat(formatter, formatString)
}

// Sets a number formatter property using a key-value pair.
//
// Added in macOS .
// Sets a number formatter property using a key-value pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterSetProperty(_:_:_:)
func CFNumberFormatterSetProperty(formatter NumberFormatterRef, key NumberFormatterKey, value TypeRef) {
	_CFNumberFormatterSetProperty(formatter, key, value)
}

// Returns the number of bytes used by a CFNumber object to store its value.
//
// Added in macOS .
// Returns the number of bytes used by a CFNumber object to store its value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetByteSize(_:)
func CFNumberGetByteSize(number NumberRef) Index {
	return _CFNumberGetByteSize(number)
}

// Returns the type used by a CFNumber object to store its value.
//
// Added in macOS .
// Returns the type used by a CFNumber object to store its value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetType(_:)
func CFNumberGetType(number NumberRef) NumberType {
	return _CFNumberGetType(number)
}

// Returns the type identifier for the CFNumber opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFNumber opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetTypeID()
func CFNumberGetTypeID() TypeID {
	return _CFNumberGetTypeID()
}

// Obtains the value of a CFNumber object cast to a specified type.
//
// Added in macOS .
// Obtains the value of a CFNumber object cast to a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberGetValue(_:_:_:)
func CFNumberGetValue(number NumberRef, theType NumberType, valuePtr unsafe.Pointer) unsafe.Pointer {
	return _CFNumberGetValue(number, theType, valuePtr)
}

// Determines whether a CFNumber object contains a value stored as one of the defined floating point types.
//
// Added in macOS .
// Determines whether a CFNumber object contains a value stored as one of the defined floating point types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberIsFloatType(_:)
func CFNumberIsFloatType(number NumberRef) unsafe.Pointer {
	return _CFNumberIsFloatType(number)
}

// Registers a new instance of a type with .
//
// Added in macOS .
// Registers a new instance of a type with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInAddInstanceForFactory(_:)
func CFPlugInAddInstanceForFactory(factoryID UUIDRef) {
	_CFPlugInAddInstanceForFactory(factoryID)
}

// Creates a CFPlugIn given its URL.
//
// Added in macOS .
// Creates a CFPlugIn given its URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInCreate(_:_:)
func CFPlugInCreate(allocator AllocatorRef, plugInURL URLRef) PlugInRef {
	return _CFPlugInCreate(allocator, plugInURL)
}

// Searches all registered plug-ins for factory functions capable of creating an instance of the given type.
//
// Added in macOS .
// Searches all registered plug-ins for factory functions capable of creating an instance of the given type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInFindFactoriesForPlugInType(_:)
func CFPlugInFindFactoriesForPlugInType(typeUUID UUIDRef) ArrayRef {
	return _CFPlugInFindFactoriesForPlugInType(typeUUID)
}

// Searches the given plug-in for factory functions capable of creating an instance of the given type.
//
// Added in macOS .
// Searches the given plug-in for factory functions capable of creating an instance of the given type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInFindFactoriesForPlugInTypeInPlugIn(_:_:)
func CFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID UUIDRef, plugIn PlugInRef) ArrayRef {
	return _CFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID, plugIn)
}

// Returns a plug-in’s bundle.
//
// Added in macOS .
// Returns a plug-in’s bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInGetBundle(_:)
func CFPlugInGetBundle(plugIn PlugInRef) BundleRef {
	return _CFPlugInGetBundle(plugIn)
}

// Returns the type identifier for the opaque type.
//
// Added in macOS .
// Returns the type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInGetTypeID()
func CFPlugInGetTypeID() TypeID {
	return _CFPlugInGetTypeID()
}

// Creates a instance of a given type using a given factory.
//
// Added in macOS .
// Creates a instance of a given type using a given factory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceCreate(_:_:_:)
func CFPlugInInstanceCreate(allocator AllocatorRef, factoryUUID UUIDRef, typeUUID UUIDRef) unsafe.Pointer {
	return _CFPlugInInstanceCreate(allocator, factoryUUID, typeUUID)
}

// Not recommended.
//
// Added in macOS .
// Not recommended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceCreateWithInstanceDataSize(_:_:_:_:_:)
func CFPlugInInstanceCreateWithInstanceDataSize(allocator AllocatorRef, instanceDataSize Index, deallocateInstanceFunction PlugInInstanceDeallocateInstanceDataFunction, factoryName StringRef, getInterfaceFunction PlugInInstanceGetInterfaceFunction) PlugInInstanceRef {
	return _CFPlugInInstanceCreateWithInstanceDataSize(allocator, instanceDataSize, deallocateInstanceFunction, factoryName, getInterfaceFunction)
}

// Not recommended.
//
// Added in macOS .
// Not recommended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetFactoryName(_:)
func CFPlugInInstanceGetFactoryName(instance PlugInInstanceRef) StringRef {
	return _CFPlugInInstanceGetFactoryName(instance)
}

// Not recommended.
//
// Added in macOS .
// Not recommended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetInstanceData(_:)
func CFPlugInInstanceGetInstanceData(instance PlugInInstanceRef) unsafe.Pointer {
	return _CFPlugInInstanceGetInstanceData(instance)
}

// Not recommended.
//
// Added in macOS .
// Not recommended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetInterfaceFunctionTable(_:_:_:)
func CFPlugInInstanceGetInterfaceFunctionTable(instance PlugInInstanceRef, interfaceName StringRef, ftbl unsafe.Pointer) unsafe.Pointer {
	return _CFPlugInInstanceGetInterfaceFunctionTable(instance, interfaceName, ftbl)
}

// Not recommended.
//
// Added in macOS .
// Not recommended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetTypeID()
func CFPlugInInstanceGetTypeID() TypeID {
	return _CFPlugInInstanceGetTypeID()
}

// Determines whether or not a plug-in is loaded on demand.
//
// Added in macOS .
// Determines whether or not a plug-in is loaded on demand.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInIsLoadOnDemand(_:)
func CFPlugInIsLoadOnDemand(plugIn PlugInRef) unsafe.Pointer {
	return _CFPlugInIsLoadOnDemand(plugIn)
}

// Registers a factory function and its UUID with a object.
//
// Added in macOS .
// Registers a factory function and its UUID with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterFactoryFunction(_:_:)
func CFPlugInRegisterFactoryFunction(factoryUUID UUIDRef, func_ PlugInFactoryFunction) unsafe.Pointer {
	return _CFPlugInRegisterFactoryFunction(factoryUUID, func_)
}

// Registers a factory function with a object using the function’s name instead of its UUID.
//
// Added in macOS .
// Registers a factory function with a object using the function’s name instead of its UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterFactoryFunctionByName(_:_:_:)
func CFPlugInRegisterFactoryFunctionByName(factoryUUID UUIDRef, plugIn PlugInRef, functionName StringRef) unsafe.Pointer {
	return _CFPlugInRegisterFactoryFunctionByName(factoryUUID, plugIn, functionName)
}

// Registers a type and its corresponding factory function with a object.
//
// Added in macOS .
// Registers a type and its corresponding factory function with a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRegisterPlugInType(_:_:)
func CFPlugInRegisterPlugInType(factoryUUID UUIDRef, typeUUID UUIDRef) unsafe.Pointer {
	return _CFPlugInRegisterPlugInType(factoryUUID, typeUUID)
}

// Unregisters an instance of a type with .
//
// Added in macOS .
// Unregisters an instance of a type with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInRemoveInstanceForFactory(_:)
func CFPlugInRemoveInstanceForFactory(factoryID UUIDRef) {
	_CFPlugInRemoveInstanceForFactory(factoryID)
}

// Enables or disables load on demand for plug-ins that do dynamic registration (only when a client requests an instance of a supported type).
//
// Added in macOS .
// Enables or disables load on demand for plug-ins that do dynamic registration (only when a client requests an instance of a supported type).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInSetLoadOnDemand(_:_:)
func CFPlugInSetLoadOnDemand(plugIn PlugInRef, flag unsafe.Pointer) {
	_CFPlugInSetLoadOnDemand(plugIn, flag)
}

// Removes the given function from a plug-in’s list of registered factory functions.
//
// Added in macOS .
// Removes the given function from a plug-in’s list of registered factory functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInUnregisterFactory(_:)
func CFPlugInUnregisterFactory(factoryUUID UUIDRef) unsafe.Pointer {
	return _CFPlugInUnregisterFactory(factoryUUID)
}

// Removes the given type from a plug-in’s list of registered types.
//
// Added in macOS .
// Removes the given type from a plug-in’s list of registered types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInUnregisterPlugInType(_:_:)
func CFPlugInUnregisterPlugInType(factoryUUID UUIDRef, typeUUID UUIDRef) unsafe.Pointer {
	return _CFPlugInUnregisterPlugInType(factoryUUID, typeUUID)
}

// Adds suite preferences to an application’s preference search chain.
//
// Added in macOS .
// Adds suite preferences to an application’s preference search chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAddSuitePreferencesToApp(_:_:)
func CFPreferencesAddSuitePreferencesToApp(applicationID StringRef, suiteID StringRef) {
	_CFPreferencesAddSuitePreferencesToApp(applicationID, suiteID)
}

// Writes to permanent storage all pending changes to the preference data for the application, and reads the latest preference data from permanent storage.
//
// Added in macOS .
// Writes to permanent storage all pending changes to the preference data for the application, and reads the latest preference data from permanent storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAppSynchronize(_:)
func CFPreferencesAppSynchronize(applicationID StringRef) unsafe.Pointer {
	return _CFPreferencesAppSynchronize(applicationID)
}

// Determines whether or not a given key has been imposed on the user.
//
// Added in macOS .
// Determines whether or not a given key has been imposed on the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesAppValueIsForced(_:_:)
func CFPreferencesAppValueIsForced(key StringRef, applicationID StringRef) unsafe.Pointer {
	return _CFPreferencesAppValueIsForced(key, applicationID)
}

// Obtains a preference value for the specified key and application.
//
// Added in macOS .
// Obtains a preference value for the specified key and application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyAppValue(_:_:)
func CFPreferencesCopyAppValue(key StringRef, applicationID StringRef) PropertyListRef {
	return _CFPreferencesCopyAppValue(key, applicationID)
}

// Constructs and returns the list of all applications that have preferences in the scope of the specified user and host.

// Constructs and returns the list of all applications that have preferences in the scope of the specified user and host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyApplicationList(_:_:)
func CFPreferencesCopyApplicationList(userName StringRef, hostName StringRef) ArrayRef {
	return _CFPreferencesCopyApplicationList(userName, hostName)
}

// Constructs and returns the list of all keys set in the specified domain.
//
// Added in macOS .
// Constructs and returns the list of all keys set in the specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyKeyList(_:_:_:)
func CFPreferencesCopyKeyList(applicationID StringRef, userName StringRef, hostName StringRef) ArrayRef {
	return _CFPreferencesCopyKeyList(applicationID, userName, hostName)
}

// Returns a dictionary containing preference values for multiple keys.
//
// Added in macOS .
// Returns a dictionary containing preference values for multiple keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyMultiple(_:_:_:_:)
func CFPreferencesCopyMultiple(keysToFetch ArrayRef, applicationID StringRef, userName StringRef, hostName StringRef) DictionaryRef {
	return _CFPreferencesCopyMultiple(keysToFetch, applicationID, userName, hostName)
}

// Returns a preference value for a given domain.
//
// Added in macOS .
// Returns a preference value for a given domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesCopyValue(_:_:_:_:)
func CFPreferencesCopyValue(key StringRef, applicationID StringRef, userName StringRef, hostName StringRef) PropertyListRef {
	return _CFPreferencesCopyValue(key, applicationID, userName, hostName)
}

// Convenience function that directly obtains a Boolean preference value for the specified key.
//
// Added in macOS .
// Convenience function that directly obtains a Boolean preference value for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesGetAppBooleanValue(_:_:_:)
func CFPreferencesGetAppBooleanValue(key StringRef, applicationID StringRef, keyExistsAndHasValidFormat unsafe.Pointer) unsafe.Pointer {
	return _CFPreferencesGetAppBooleanValue(key, applicationID, keyExistsAndHasValidFormat)
}

// Convenience function that directly obtains an integer preference value for the specified key.
//
// Added in macOS .
// Convenience function that directly obtains an integer preference value for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesGetAppIntegerValue(_:_:_:)
func CFPreferencesGetAppIntegerValue(key StringRef, applicationID StringRef, keyExistsAndHasValidFormat unsafe.Pointer) Index {
	return _CFPreferencesGetAppIntegerValue(key, applicationID, keyExistsAndHasValidFormat)
}

// Removes suite preferences from an application’s search chain.
//
// Added in macOS .
// Removes suite preferences from an application’s search chain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesRemoveSuitePreferencesFromApp(_:_:)
func CFPreferencesRemoveSuitePreferencesFromApp(applicationID StringRef, suiteID StringRef) {
	_CFPreferencesRemoveSuitePreferencesFromApp(applicationID, suiteID)
}

// Adds, modifies, or removes a preference.
//
// Added in macOS .
// Adds, modifies, or removes a preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetAppValue(_:_:_:)
func CFPreferencesSetAppValue(key StringRef, value PropertyListRef, applicationID StringRef) {
	_CFPreferencesSetAppValue(key, value, applicationID)
}

// Convenience function that allows you to set and remove multiple preference values.
//
// Added in macOS .
// Convenience function that allows you to set and remove multiple preference values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetMultiple(_:_:_:_:_:)
func CFPreferencesSetMultiple(keysToSet DictionaryRef, keysToRemove ArrayRef, applicationID StringRef, userName StringRef, hostName StringRef) {
	_CFPreferencesSetMultiple(keysToSet, keysToRemove, applicationID, userName, hostName)
}

// Adds, modifies, or removes a preference value for the specified domain.
//
// Added in macOS .
// Adds, modifies, or removes a preference value for the specified domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSetValue(_:_:_:_:_:)
func CFPreferencesSetValue(key StringRef, value PropertyListRef, applicationID StringRef, userName StringRef, hostName StringRef) {
	_CFPreferencesSetValue(key, value, applicationID, userName, hostName)
}

// For the specified domain, writes all pending changes to preference data to permanent storage, and reads latest preference data from permanent storage.
//
// Added in macOS .
// For the specified domain, writes all pending changes to preference data to permanent storage, and reads latest preference data from permanent storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPreferencesSynchronize(_:_:_:)
func CFPreferencesSynchronize(applicationID StringRef, userName StringRef, hostName StringRef) unsafe.Pointer {
	return _CFPreferencesSynchronize(applicationID, userName, hostName)
}

// Returns a CFData object containing a serialized representation of a given property list in a specified format.
//
// Added in macOS 10.6.
// Returns a CFData object containing a serialized representation of a given property list in a specified format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateData(_:_:_:_:_:)
func CFPropertyListCreateData(allocator AllocatorRef, propertyList PropertyListRef, format PropertyListFormat, options OptionFlags, error_ unsafe.Pointer) DataRef {
	return _CFPropertyListCreateData(allocator, propertyList, format, options, error_)
}

// Recursively creates a copy of a given property list.
//
// Added in macOS .
// Recursively creates a copy of a given property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateDeepCopy(_:_:_:)
func CFPropertyListCreateDeepCopy(allocator AllocatorRef, propertyList PropertyListRef, mutabilityOption OptionFlags) PropertyListRef {
	return _CFPropertyListCreateDeepCopy(allocator, propertyList, mutabilityOption)
}

// Creates a property list using data from a stream.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Creates a property list using data from a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateFromStream(_:_:_:_:_:_:)
func CFPropertyListCreateFromStream(allocator AllocatorRef, stream ReadStreamRef, streamLength Index, mutabilityOption OptionFlags, format unsafe.Pointer, errorString unsafe.Pointer) PropertyListRef {
	return _CFPropertyListCreateFromStream(allocator, stream, streamLength, mutabilityOption, format, errorString)
}

// Creates a property list using the specified XML or binary property list data.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Creates a property list using the specified XML or binary property list data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateFromXMLData(_:_:_:_:)
func CFPropertyListCreateFromXMLData(allocator AllocatorRef, xmlData DataRef, mutabilityOption OptionFlags, errorString unsafe.Pointer) PropertyListRef {
	return _CFPropertyListCreateFromXMLData(allocator, xmlData, mutabilityOption, errorString)
}

// Creates a property list from a given CFData object.
//
// Added in macOS 10.6.
// Creates a property list from a given CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateWithData(_:_:_:_:_:)
func CFPropertyListCreateWithData(allocator AllocatorRef, data DataRef, options OptionFlags, format unsafe.Pointer, error_ unsafe.Pointer) PropertyListRef {
	return _CFPropertyListCreateWithData(allocator, data, options, format, error_)
}

// Create and return a property list with a CFReadStream input.
//
// Added in macOS 10.6.
// Create and return a property list with a CFReadStream input.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateWithStream(_:_:_:_:_:_:)
func CFPropertyListCreateWithStream(allocator AllocatorRef, stream ReadStreamRef, streamLength Index, options OptionFlags, format unsafe.Pointer, error_ unsafe.Pointer) PropertyListRef {
	return _CFPropertyListCreateWithStream(allocator, stream, streamLength, options, format, error_)
}

// Creates an XML representation of the specified property list.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.0.
// Creates an XML representation of the specified property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListCreateXMLData(_:_:)
func CFPropertyListCreateXMLData(allocator AllocatorRef, propertyList PropertyListRef) DataRef {
	return _CFPropertyListCreateXMLData(allocator, propertyList)
}

// Determines if a property list is valid.
//
// Added in macOS .
// Determines if a property list is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListIsValid(_:_:)
func CFPropertyListIsValid(plist PropertyListRef, format PropertyListFormat) unsafe.Pointer {
	return _CFPropertyListIsValid(plist, format)
}

// Write the bytes of a serialized property list out to a stream.
//
// Added in macOS 10.6.
// Write the bytes of a serialized property list out to a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListWrite(_:_:_:_:_:)
func CFPropertyListWrite(propertyList PropertyListRef, stream WriteStreamRef, format PropertyListFormat, options OptionFlags, error_ unsafe.Pointer) Index {
	return _CFPropertyListWrite(propertyList, stream, format, options, error_)
}

// Writes the bytes of a property list serialization out to a stream.
//
// Deprecated: This function was deprecated in macOS 10.10.
//
// Added in macOS 10.2.
// Writes the bytes of a property list serialization out to a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListWriteToStream(_:_:_:_:)
func CFPropertyListWriteToStream(propertyList PropertyListRef, stream WriteStreamRef, format PropertyListFormat, errorString unsafe.Pointer) Index {
	return _CFPropertyListWriteToStream(propertyList, stream, format, errorString)
}

// Closes a readable stream.
//
// Added in macOS .
// Closes a readable stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamClose(_:)
func CFReadStreamClose(stream ReadStreamRef) {
	_CFReadStreamClose(stream)
}

// CFReadStreamCopyDispatchQueue is a CoreFoundation function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyDispatchQueue(_:)
func CFReadStreamCopyDispatchQueue(stream ReadStreamRef) unsafe.Pointer {
	return _CFReadStreamCopyDispatchQueue(stream)
}

// Returns the error associated with a stream.
//
// Added in macOS 10.5.
// Returns the error associated with a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyError(_:)
func CFReadStreamCopyError(stream ReadStreamRef) ErrorRef {
	return _CFReadStreamCopyError(stream)
}

// Returns the value of a property for a stream.
//
// Added in macOS .
// Returns the value of a property for a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCopyProperty(_:_:)
func CFReadStreamCopyProperty(stream ReadStreamRef, propertyName StreamPropertyKey) TypeRef {
	return _CFReadStreamCopyProperty(stream, propertyName)
}

// Creates a readable stream for a block of memory.
//
// Added in macOS .
// Creates a readable stream for a block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCreateWithBytesNoCopy(_:_:_:_:)
func CFReadStreamCreateWithBytesNoCopy(alloc AllocatorRef, bytes unsafe.Pointer, length Index, bytesDeallocator AllocatorRef) ReadStreamRef {
	return _CFReadStreamCreateWithBytesNoCopy(alloc, bytes, length, bytesDeallocator)
}

// Creates a readable stream for a file.
//
// Added in macOS .
// Creates a readable stream for a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamCreateWithFile(_:_:)
func CFReadStreamCreateWithFile(alloc AllocatorRef, fileURL URLRef) ReadStreamRef {
	return _CFReadStreamCreateWithFile(alloc, fileURL)
}

// Returns a pointer to a stream’s internal buffer of unread data, if possible.
//
// Added in macOS .
// Returns a pointer to a stream’s internal buffer of unread data, if possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetBuffer(_:_:_:)
func CFReadStreamGetBuffer(stream ReadStreamRef, maxBytesToRead Index, numBytesRead unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamGetBuffer(stream, maxBytesToRead, numBytesRead)
}

// Returns the error status of a stream.
//
// Added in macOS .
// Returns the error status of a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetError(_:)
func CFReadStreamGetError(stream ReadStreamRef) StreamError {
	return _CFReadStreamGetError(stream)
}

// Returns the current state of a stream.
//
// Added in macOS .
// Returns the current state of a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetStatus(_:)
func CFReadStreamGetStatus(stream ReadStreamRef) StreamStatus {
	return _CFReadStreamGetStatus(stream)
}

// Returns the type identifier the opaque type.
//
// Added in macOS .
// Returns the type identifier the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamGetTypeID()
func CFReadStreamGetTypeID() TypeID {
	return _CFReadStreamGetTypeID()
}

// Returns a Boolean value that indicates whether a readable stream has data that can be read without blocking.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a readable stream has data that can be read without blocking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamHasBytesAvailable(_:)
func CFReadStreamHasBytesAvailable(stream ReadStreamRef) unsafe.Pointer {
	return _CFReadStreamHasBytesAvailable(stream)
}

// Opens a stream for reading.
//
// Added in macOS .
// Opens a stream for reading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamOpen(_:)
func CFReadStreamOpen(stream ReadStreamRef) unsafe.Pointer {
	return _CFReadStreamOpen(stream)
}

// Reads data from a readable stream.
//
// Added in macOS .
// Reads data from a readable stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamRead(_:_:_:)
func CFReadStreamRead(stream ReadStreamRef, buffer unsafe.Pointer, bufferLength Index) Index {
	return _CFReadStreamRead(stream, buffer, bufferLength)
}

// Schedules a stream into a run loop.
//
// Added in macOS .
// Schedules a stream into a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamScheduleWithRunLoop(_:_:_:)
func CFReadStreamScheduleWithRunLoop(stream ReadStreamRef, runLoop RunLoopRef, runLoopMode RunLoopMode) {
	_CFReadStreamScheduleWithRunLoop(stream, runLoop, runLoopMode)
}

// Assigns a client to a stream, which receives callbacks when certain events occur.
//
// Added in macOS .
// Assigns a client to a stream, which receives callbacks when certain events occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetClient(_:_:_:_:)
func CFReadStreamSetClient(stream ReadStreamRef, streamEvents OptionFlags, clientCB ReadStreamClientCallBack, clientContext unsafe.Pointer) unsafe.Pointer {
	return _CFReadStreamSetClient(stream, streamEvents, clientCB, clientContext)
}

// CFReadStreamSetDispatchQueue is a CoreFoundation function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetDispatchQueue(_:_:)
func CFReadStreamSetDispatchQueue(stream ReadStreamRef, q unsafe.Pointer) {
	_CFReadStreamSetDispatchQueue(stream, q)
}

// Sets the value of a property for a stream.
//
// Added in macOS .
// Sets the value of a property for a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamSetProperty(_:_:_:)
func CFReadStreamSetProperty(stream ReadStreamRef, propertyName StreamPropertyKey, propertyValue TypeRef) unsafe.Pointer {
	return _CFReadStreamSetProperty(stream, propertyName, propertyValue)
}

// Removes a read stream from a given run loop.
//
// Added in macOS .
// Removes a read stream from a given run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamUnscheduleFromRunLoop(_:_:_:)
func CFReadStreamUnscheduleFromRunLoop(stream ReadStreamRef, runLoop RunLoopRef, runLoopMode RunLoopMode) {
	_CFReadStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode)
}

// Releases a Core Foundation object.
//
// Added in macOS .
// Releases a Core Foundation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRelease
func CFRelease(cf TypeRef) {
	_CFRelease(cf)
}

// Retains a Core Foundation object.
//
// Added in macOS .
// Retains a Core Foundation object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRetain
func CFRetain(cf TypeRef) TypeRef {
	return _CFRetain(cf)
}

// Adds a mode to the set of run loop common modes.
//
// Added in macOS .
// Adds a mode to the set of run loop common modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddCommonMode(_:_:)
func CFRunLoopAddCommonMode(rl RunLoopRef, mode RunLoopMode) {
	_CFRunLoopAddCommonMode(rl, mode)
}

// Adds a CFRunLoopObserver object to a run loop mode.
//
// Added in macOS .
// Adds a CFRunLoopObserver object to a run loop mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddObserver(_:_:_:)
func CFRunLoopAddObserver(rl RunLoopRef, observer RunLoopObserverRef, mode RunLoopMode) {
	_CFRunLoopAddObserver(rl, observer, mode)
}

// Adds a CFRunLoopSource object to a run loop mode.
//
// Added in macOS .
// Adds a CFRunLoopSource object to a run loop mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddSource(_:_:_:)
func CFRunLoopAddSource(rl RunLoopRef, source RunLoopSourceRef, mode RunLoopMode) {
	_CFRunLoopAddSource(rl, source, mode)
}

// Adds a CFRunLoopTimer object to a run loop mode.
//
// Added in macOS .
// Adds a CFRunLoopTimer object to a run loop mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopAddTimer(_:_:_:)
func CFRunLoopAddTimer(rl RunLoopRef, timer RunLoopTimerRef, mode RunLoopMode) {
	_CFRunLoopAddTimer(rl, timer, mode)
}

// Returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopObserver object.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopObserver object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsObserver(_:_:_:)
func CFRunLoopContainsObserver(rl RunLoopRef, observer RunLoopObserverRef, mode RunLoopMode) unsafe.Pointer {
	return _CFRunLoopContainsObserver(rl, observer, mode)
}

// Returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopSource object.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopSource object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsSource(_:_:_:)
func CFRunLoopContainsSource(rl RunLoopRef, source RunLoopSourceRef, mode RunLoopMode) unsafe.Pointer {
	return _CFRunLoopContainsSource(rl, source, mode)
}

// Returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopTimer object.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a run loop mode contains a particular CFRunLoopTimer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopContainsTimer(_:_:_:)
func CFRunLoopContainsTimer(rl RunLoopRef, timer RunLoopTimerRef, mode RunLoopMode) unsafe.Pointer {
	return _CFRunLoopContainsTimer(rl, timer, mode)
}

// Returns an array that contains all the defined modes for a CFRunLoop object.
//
// Added in macOS .
// Returns an array that contains all the defined modes for a CFRunLoop object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopCopyAllModes(_:)
func CFRunLoopCopyAllModes(rl RunLoopRef) ArrayRef {
	return _CFRunLoopCopyAllModes(rl)
}

// Returns the name of the mode in which a given run loop is currently running.
//
// Added in macOS .
// Returns the name of the mode in which a given run loop is currently running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopCopyCurrentMode(_:)
func CFRunLoopCopyCurrentMode(rl RunLoopRef) RunLoopMode {
	return _CFRunLoopCopyCurrentMode(rl)
}

// Returns the CFRunLoop object for the current thread.
//
// Added in macOS .
// Returns the CFRunLoop object for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetCurrent()
func CFRunLoopGetCurrent() RunLoopRef {
	return _CFRunLoopGetCurrent()
}

// Returns the main CFRunLoop object.
//
// Added in macOS .
// Returns the main CFRunLoop object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetMain()
func CFRunLoopGetMain() RunLoopRef {
	return _CFRunLoopGetMain()
}

// Returns the time at which the next timer will fire.
//
// Added in macOS .
// Returns the time at which the next timer will fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetNextTimerFireDate(_:_:)
func CFRunLoopGetNextTimerFireDate(rl RunLoopRef, mode RunLoopMode) AbsoluteTime {
	return _CFRunLoopGetNextTimerFireDate(rl, mode)
}

// Returns the type identifier for the CFRunLoop opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFRunLoop opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopGetTypeID()
func CFRunLoopGetTypeID() TypeID {
	return _CFRunLoopGetTypeID()
}

// Returns a Boolean value that indicates whether the run loop is waiting for an event.
//
// Added in macOS .
// Returns a Boolean value that indicates whether the run loop is waiting for an event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopIsWaiting(_:)
func CFRunLoopIsWaiting(rl RunLoopRef) unsafe.Pointer {
	return _CFRunLoopIsWaiting(rl)
}

// Creates a CFRunLoopObserver object with a function callback.
//
// Added in macOS .
// Creates a CFRunLoopObserver object with a function callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverCreate(_:_:_:_:_:_:)
func CFRunLoopObserverCreate(allocator AllocatorRef, activities OptionFlags, repeats unsafe.Pointer, order Index, callout RunLoopObserverCallBack, context unsafe.Pointer) RunLoopObserverRef {
	return _CFRunLoopObserverCreate(allocator, activities, repeats, order, callout, context)
}

// Creates a CFRunLoopObserver object with a block-based handler.
//
// Added in macOS 10.7.
// Creates a CFRunLoopObserver object with a block-based handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverCreateWithHandler(_:_:_:_:_:)
func CFRunLoopObserverCreateWithHandler(allocator AllocatorRef, activities OptionFlags, repeats unsafe.Pointer, order Index) RunLoopObserverRef {
	return _CFRunLoopObserverCreateWithHandler(allocator, activities, repeats, order)
}

// Returns a Boolean value that indicates whether a CFRunLoopObserver repeats.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a CFRunLoopObserver repeats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverDoesRepeat(_:)
func CFRunLoopObserverDoesRepeat(observer RunLoopObserverRef) unsafe.Pointer {
	return _CFRunLoopObserverDoesRepeat(observer)
}

// Returns the run loop stages during which an observer runs.
//
// Added in macOS .
// Returns the run loop stages during which an observer runs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetActivities(_:)
func CFRunLoopObserverGetActivities(observer RunLoopObserverRef) OptionFlags {
	return _CFRunLoopObserverGetActivities(observer)
}

// Returns the context information for a CFRunLoopObserver object.
//
// Added in macOS .
// Returns the context information for a CFRunLoopObserver object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetContext(_:_:)
func CFRunLoopObserverGetContext(observer RunLoopObserverRef, context unsafe.Pointer) {
	_CFRunLoopObserverGetContext(observer, context)
}

// Returns the ordering parameter for a CFRunLoopObserver object.
//
// Added in macOS .
// Returns the ordering parameter for a CFRunLoopObserver object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetOrder(_:)
func CFRunLoopObserverGetOrder(observer RunLoopObserverRef) Index {
	return _CFRunLoopObserverGetOrder(observer)
}

// Returns the type identifier for the CFRunLoopObserver opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFRunLoopObserver opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverGetTypeID()
func CFRunLoopObserverGetTypeID() TypeID {
	return _CFRunLoopObserverGetTypeID()
}

// Invalidates a CFRunLoopObserver object, stopping it from ever firing again.
//
// Added in macOS .
// Invalidates a CFRunLoopObserver object, stopping it from ever firing again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverInvalidate(_:)
func CFRunLoopObserverInvalidate(observer RunLoopObserverRef) {
	_CFRunLoopObserverInvalidate(observer)
}

// Returns a Boolean value that indicates whether a CFRunLoopObserver object is valid and able to fire.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a CFRunLoopObserver object is valid and able to fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverIsValid(_:)
func CFRunLoopObserverIsValid(observer RunLoopObserverRef) unsafe.Pointer {
	return _CFRunLoopObserverIsValid(observer)
}

// Enqueues a block object on a given runloop to be executed as the runloop cycles in specified modes.
//
// Added in macOS 10.6.
// Enqueues a block object on a given runloop to be executed as the runloop cycles in specified modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopPerformBlock(_:_:_:)
func CFRunLoopPerformBlock(rl RunLoopRef, mode TypeRef) {
	_CFRunLoopPerformBlock(rl, mode)
}

// Removes a CFRunLoopObserver object from a run loop mode.
//
// Added in macOS .
// Removes a CFRunLoopObserver object from a run loop mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveObserver(_:_:_:)
func CFRunLoopRemoveObserver(rl RunLoopRef, observer RunLoopObserverRef, mode RunLoopMode) {
	_CFRunLoopRemoveObserver(rl, observer, mode)
}

// Removes a CFRunLoopSource object from a run loop mode.
//
// Added in macOS .
// Removes a CFRunLoopSource object from a run loop mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveSource(_:_:_:)
func CFRunLoopRemoveSource(rl RunLoopRef, source RunLoopSourceRef, mode RunLoopMode) {
	_CFRunLoopRemoveSource(rl, source, mode)
}

// Removes a CFRunLoopTimer object from a run loop mode.
//
// Added in macOS .
// Removes a CFRunLoopTimer object from a run loop mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRemoveTimer(_:_:_:)
func CFRunLoopRemoveTimer(rl RunLoopRef, timer RunLoopTimerRef, mode RunLoopMode) {
	_CFRunLoopRemoveTimer(rl, timer, mode)
}

// Runs the current thread’s CFRunLoop object in its default mode indefinitely.
//
// Added in macOS .
// Runs the current thread’s CFRunLoop object in its default mode indefinitely.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRun()
func CFRunLoopRun() {
	_CFRunLoopRun()
}

// Runs the current thread’s CFRunLoop object in a particular mode.
//
// Added in macOS .
// Runs the current thread’s CFRunLoop object in a particular mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunInMode(_:_:_:)
func CFRunLoopRunInMode(mode RunLoopMode, seconds TimeInterval, returnAfterSourceHandled unsafe.Pointer) RunLoopRunResult {
	return _CFRunLoopRunInMode(mode, seconds, returnAfterSourceHandled)
}

// Creates a CFRunLoopSource object.
//
// Added in macOS .
// Creates a CFRunLoopSource object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceCreate(_:_:_:)
func CFRunLoopSourceCreate(allocator AllocatorRef, order Index, context unsafe.Pointer) RunLoopSourceRef {
	return _CFRunLoopSourceCreate(allocator, order, context)
}

// Returns the context information for a CFRunLoopSource object.
//
// Added in macOS .
// Returns the context information for a CFRunLoopSource object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetContext(_:_:)
func CFRunLoopSourceGetContext(source RunLoopSourceRef, context unsafe.Pointer) {
	_CFRunLoopSourceGetContext(source, context)
}

// Returns the ordering parameter for a CFRunLoopSource object.
//
// Added in macOS .
// Returns the ordering parameter for a CFRunLoopSource object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetOrder(_:)
func CFRunLoopSourceGetOrder(source RunLoopSourceRef) Index {
	return _CFRunLoopSourceGetOrder(source)
}

// Returns the type identifier of the CFRunLoopSource opaque type.
//
// Added in macOS .
// Returns the type identifier of the CFRunLoopSource opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceGetTypeID()
func CFRunLoopSourceGetTypeID() TypeID {
	return _CFRunLoopSourceGetTypeID()
}

// Invalidates a CFRunLoopSource object, stopping it from ever firing again.
//
// Added in macOS .
// Invalidates a CFRunLoopSource object, stopping it from ever firing again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceInvalidate(_:)
func CFRunLoopSourceInvalidate(source RunLoopSourceRef) {
	_CFRunLoopSourceInvalidate(source)
}

// Returns a Boolean value that indicates whether a CFRunLoopSource object is valid and able to fire.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a CFRunLoopSource object is valid and able to fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceIsValid(_:)
func CFRunLoopSourceIsValid(source RunLoopSourceRef) unsafe.Pointer {
	return _CFRunLoopSourceIsValid(source)
}

// Signals a CFRunLoopSource object, marking it as ready to fire.
//
// Added in macOS .
// Signals a CFRunLoopSource object, marking it as ready to fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSourceSignal(_:)
func CFRunLoopSourceSignal(source RunLoopSourceRef) {
	_CFRunLoopSourceSignal(source)
}

// Forces a CFRunLoop object to stop running.
//
// Added in macOS .
// Forces a CFRunLoop object to stop running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopStop(_:)
func CFRunLoopStop(rl RunLoopRef) {
	_CFRunLoopStop(rl)
}

// Creates a new CFRunLoopTimer object with a function callback.
//
// Added in macOS .
// Creates a new CFRunLoopTimer object with a function callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerCreate(_:_:_:_:_:_:_:)
func CFRunLoopTimerCreate(allocator AllocatorRef, fireDate AbsoluteTime, interval TimeInterval, flags OptionFlags, order Index, callout RunLoopTimerCallBack, context unsafe.Pointer) RunLoopTimerRef {
	return _CFRunLoopTimerCreate(allocator, fireDate, interval, flags, order, callout, context)
}

// Creates a new CFRunLoopTimer object with a block-based handler.
//
// Added in macOS 10.7.
// Creates a new CFRunLoopTimer object with a block-based handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerCreateWithHandler(_:_:_:_:_:_:)
func CFRunLoopTimerCreateWithHandler(allocator AllocatorRef, fireDate AbsoluteTime, interval TimeInterval, flags OptionFlags, order Index) RunLoopTimerRef {
	return _CFRunLoopTimerCreateWithHandler(allocator, fireDate, interval, flags, order)
}

// Returns a Boolean value that indicates whether a CFRunLoopTimer object repeats.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a CFRunLoopTimer object repeats.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerDoesRepeat(_:)
func CFRunLoopTimerDoesRepeat(timer RunLoopTimerRef) unsafe.Pointer {
	return _CFRunLoopTimerDoesRepeat(timer)
}

// Returns the context information for a CFRunLoopTimer object.
//
// Added in macOS .
// Returns the context information for a CFRunLoopTimer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetContext(_:_:)
func CFRunLoopTimerGetContext(timer RunLoopTimerRef, context unsafe.Pointer) {
	_CFRunLoopTimerGetContext(timer, context)
}

// Returns the firing interval of a repeating CFRunLoopTimer object.
//
// Added in macOS .
// Returns the firing interval of a repeating CFRunLoopTimer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetInterval(_:)
func CFRunLoopTimerGetInterval(timer RunLoopTimerRef) TimeInterval {
	return _CFRunLoopTimerGetInterval(timer)
}

// Returns the next firing time for a CFRunLoopTimer object.
//
// Added in macOS .
// Returns the next firing time for a CFRunLoopTimer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetNextFireDate(_:)
func CFRunLoopTimerGetNextFireDate(timer RunLoopTimerRef) AbsoluteTime {
	return _CFRunLoopTimerGetNextFireDate(timer)
}

// Returns the ordering parameter for a CFRunLoopTimer object.
//
// Added in macOS .
// Returns the ordering parameter for a CFRunLoopTimer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetOrder(_:)
func CFRunLoopTimerGetOrder(timer RunLoopTimerRef) Index {
	return _CFRunLoopTimerGetOrder(timer)
}

// CFRunLoopTimerGetTolerance is a CoreFoundation function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetTolerance(_:)
func CFRunLoopTimerGetTolerance(timer RunLoopTimerRef) TimeInterval {
	return _CFRunLoopTimerGetTolerance(timer)
}

// Returns the type identifier of the CFRunLoopTimer opaque type.
//
// Added in macOS .
// Returns the type identifier of the CFRunLoopTimer opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerGetTypeID()
func CFRunLoopTimerGetTypeID() TypeID {
	return _CFRunLoopTimerGetTypeID()
}

// Invalidates a CFRunLoopTimer object, stopping it from ever firing again.
//
// Added in macOS .
// Invalidates a CFRunLoopTimer object, stopping it from ever firing again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerInvalidate(_:)
func CFRunLoopTimerInvalidate(timer RunLoopTimerRef) {
	_CFRunLoopTimerInvalidate(timer)
}

// Returns a Boolean value that indicates whether a CFRunLoopTimer object is valid and able to fire.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a CFRunLoopTimer object is valid and able to fire.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerIsValid(_:)
func CFRunLoopTimerIsValid(timer RunLoopTimerRef) unsafe.Pointer {
	return _CFRunLoopTimerIsValid(timer)
}

// Sets the next firing date for a CFRunLoopTimer object .
//
// Added in macOS .
// Sets the next firing date for a CFRunLoopTimer object .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerSetNextFireDate(_:_:)
func CFRunLoopTimerSetNextFireDate(timer RunLoopTimerRef, fireDate AbsoluteTime) {
	_CFRunLoopTimerSetNextFireDate(timer, fireDate)
}

// CFRunLoopTimerSetTolerance is a CoreFoundation function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerSetTolerance(_:_:)
func CFRunLoopTimerSetTolerance(timer RunLoopTimerRef, tolerance TimeInterval) {
	_CFRunLoopTimerSetTolerance(timer, tolerance)
}

// Wakes a waiting CFRunLoop object.
//
// Added in macOS .
// Wakes a waiting CFRunLoop object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopWakeUp(_:)
func CFRunLoopWakeUp(rl RunLoopRef) {
	_CFRunLoopWakeUp(rl)
}

// Adds a value to a CFMutableSet object.
//
// Added in macOS .
// Adds a value to a CFMutableSet object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetAddValue(_:_:)
func CFSetAddValue(theSet MutableSetRef, value unsafe.Pointer) {
	_CFSetAddValue(theSet, value)
}

// Calls a function once for each value in a set.
//
// Added in macOS .
// Calls a function once for each value in a set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetApplyFunction(_:_:_:)
func CFSetApplyFunction(theSet SetRef, applier SetApplierFunction, context unsafe.Pointer) {
	_CFSetApplyFunction(theSet, applier, context)
}

// Returns a Boolean that indicates whether a set contains a given value.
//
// Added in macOS .
// Returns a Boolean that indicates whether a set contains a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetContainsValue(_:_:)
func CFSetContainsValue(theSet SetRef, value unsafe.Pointer) unsafe.Pointer {
	return _CFSetContainsValue(theSet, value)
}

// Creates an immutable CFSet object containing supplied values.
//
// Added in macOS .
// Creates an immutable CFSet object containing supplied values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCreate(_:_:_:_:)
func CFSetCreate(allocator AllocatorRef, values unsafe.Pointer, numValues Index, callBacks unsafe.Pointer) SetRef {
	return _CFSetCreate(allocator, values, numValues, callBacks)
}

// Creates an immutable set containing the values of an existing set.
//
// Added in macOS .
// Creates an immutable set containing the values of an existing set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateCopy(_:_:)
func CFSetCreateCopy(allocator AllocatorRef, theSet SetRef) SetRef {
	return _CFSetCreateCopy(allocator, theSet)
}

// Creates an empty CFMutableSet object.
//
// Added in macOS .
// Creates an empty CFMutableSet object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateMutable(_:_:_:)
func CFSetCreateMutable(allocator AllocatorRef, capacity Index, callBacks unsafe.Pointer) MutableSetRef {
	return _CFSetCreateMutable(allocator, capacity, callBacks)
}

// Creates a new mutable set with the values from another set.
//
// Added in macOS .
// Creates a new mutable set with the values from another set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCreateMutableCopy(_:_:_:)
func CFSetCreateMutableCopy(allocator AllocatorRef, capacity Index, theSet SetRef) MutableSetRef {
	return _CFSetCreateMutableCopy(allocator, capacity, theSet)
}

// Returns the number of values currently in a set.
//
// Added in macOS .
// Returns the number of values currently in a set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetCount(_:)
func CFSetGetCount(theSet SetRef) Index {
	return _CFSetGetCount(theSet)
}

// Returns the number of values in a set that match a given value.
//
// Added in macOS .
// Returns the number of values in a set that match a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetCountOfValue(_:_:)
func CFSetGetCountOfValue(theSet SetRef, value unsafe.Pointer) Index {
	return _CFSetGetCountOfValue(theSet, value)
}

// Returns the type identifier for the CFSet type.
//
// Added in macOS .
// Returns the type identifier for the CFSet type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetTypeID()
func CFSetGetTypeID() TypeID {
	return _CFSetGetTypeID()
}

// Obtains a specified value from a set.
//
// Added in macOS .
// Obtains a specified value from a set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValue(_:_:)
func CFSetGetValue(theSet SetRef, value unsafe.Pointer) unsafe.Pointer {
	return _CFSetGetValue(theSet, value)
}

// Reports whether or not a value is in a set, and if it exists returns the value indirectly.
//
// Added in macOS .
// Reports whether or not a value is in a set, and if it exists returns the value indirectly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValueIfPresent(_:_:_:)
func CFSetGetValueIfPresent(theSet SetRef, candidate unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFSetGetValueIfPresent(theSet, candidate, value)
}

// Obtains all values in a set.
//
// Added in macOS .
// Obtains all values in a set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetGetValues(_:_:)
func CFSetGetValues(theSet SetRef, values unsafe.Pointer) {
	_CFSetGetValues(theSet, values)
}

// Removes all values from a CFMutableSet object.
//
// Added in macOS .
// Removes all values from a CFMutableSet object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetRemoveAllValues(_:)
func CFSetRemoveAllValues(theSet MutableSetRef) {
	_CFSetRemoveAllValues(theSet)
}

// Removes a value from a CFMutableSet object.
//
// Added in macOS .
// Removes a value from a CFMutableSet object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetRemoveValue(_:_:)
func CFSetRemoveValue(theSet MutableSetRef, value unsafe.Pointer) {
	_CFSetRemoveValue(theSet, value)
}

// Replaces a value in a CFMutableSet object.
//
// Added in macOS .
// Replaces a value in a CFMutableSet object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetReplaceValue(_:_:)
func CFSetReplaceValue(theSet MutableSetRef, value unsafe.Pointer) {
	_CFSetReplaceValue(theSet, value)
}

// Sets a value in a CFMutableSet object.
//
// Added in macOS .
// Sets a value in a CFMutableSet object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetSetValue(_:_:)
func CFSetSetValue(theSet MutableSetRef, value unsafe.Pointer) {
	_CFSetSetValue(theSet, value)
}

// Prints a description of a Core Foundation object to stderr.
//
// Added in macOS .
// Prints a description of a Core Foundation object to stderr.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFShow(_:)
func CFShow(obj TypeRef) {
	_CFShow(obj)
}

// Prints the attributes of a string during debugging.
//
// Added in macOS .
// Prints the attributes of a string during debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFShowStr(_:)
func CFShowStr(str StringRef) {
	_CFShowStr(str)
}

// Opens a connection to a remote socket.
//
// Added in macOS .
// Opens a connection to a remote socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketConnectToAddress(_:_:_:)
func CFSocketConnectToAddress(s SocketRef, address DataRef, timeout TimeInterval) SocketError {
	return _CFSocketConnectToAddress(s, address, timeout)
}

// Returns the local address of a CFSocket object.
//
// Added in macOS .
// Returns the local address of a CFSocket object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyAddress(_:)
func CFSocketCopyAddress(s SocketRef) DataRef {
	return _CFSocketCopyAddress(s)
}

// Returns the remote address to which a CFSocket object is connected.
//
// Added in macOS .
// Returns the remote address to which a CFSocket object is connected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyPeerAddress(_:)
func CFSocketCopyPeerAddress(s SocketRef) DataRef {
	return _CFSocketCopyPeerAddress(s)
}

// Returns a socket signature registered with a CFSocket name server.
//
// Added in macOS .
// Returns a socket signature registered with a CFSocket name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyRegisteredSocketSignature(_:_:_:_:_:)
func CFSocketCopyRegisteredSocketSignature(nameServerSignature unsafe.Pointer, timeout TimeInterval, name StringRef, signature unsafe.Pointer, nameServerAddress unsafe.Pointer) SocketError {
	return _CFSocketCopyRegisteredSocketSignature(nameServerSignature, timeout, name, signature, nameServerAddress)
}

// Returns a value registered with a CFSocket name server.
//
// Added in macOS .
// Returns a value registered with a CFSocket name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCopyRegisteredValue(_:_:_:_:_:)
func CFSocketCopyRegisteredValue(nameServerSignature unsafe.Pointer, timeout TimeInterval, name StringRef, value unsafe.Pointer, nameServerAddress unsafe.Pointer) SocketError {
	return _CFSocketCopyRegisteredValue(nameServerSignature, timeout, name, value, nameServerAddress)
}

// Creates a CFSocket object of a specified protocol and type.
//
// Added in macOS .
// Creates a CFSocket object of a specified protocol and type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreate(_:_:_:_:_:_:_:)
func CFSocketCreate(allocator AllocatorRef, protocolFamily unsafe.Pointer, socketType unsafe.Pointer, protocol_ unsafe.Pointer, callBackTypes OptionFlags, callout SocketCallBack, context unsafe.Pointer) SocketRef {
	return _CFSocketCreate(allocator, protocolFamily, socketType, protocol_, callBackTypes, callout, context)
}

// Creates a CFSocket object and opens a connection to a remote socket.
//
// Added in macOS .
// Creates a CFSocket object and opens a connection to a remote socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateConnectedToSocketSignature(_:_:_:_:_:_:)
func CFSocketCreateConnectedToSocketSignature(allocator AllocatorRef, signature unsafe.Pointer, callBackTypes OptionFlags, callout SocketCallBack, context unsafe.Pointer, timeout TimeInterval) SocketRef {
	return _CFSocketCreateConnectedToSocketSignature(allocator, signature, callBackTypes, callout, context, timeout)
}

// Creates a CFRunLoopSource object for a CFSocket object.
//
// Added in macOS .
// Creates a CFRunLoopSource object for a CFSocket object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateRunLoopSource(_:_:_:)
func CFSocketCreateRunLoopSource(allocator AllocatorRef, s SocketRef, order Index) RunLoopSourceRef {
	return _CFSocketCreateRunLoopSource(allocator, s, order)
}

// Creates a CFSocket object for a pre-existing native socket.
//
// Added in macOS .
// Creates a CFSocket object for a pre-existing native socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateWithNative(_:_:_:_:_:)
func CFSocketCreateWithNative(allocator AllocatorRef, sock SocketNativeHandle, callBackTypes OptionFlags, callout SocketCallBack, context unsafe.Pointer) SocketRef {
	return _CFSocketCreateWithNative(allocator, sock, callBackTypes, callout, context)
}

// Creates a CFSocket object using information from a CFSocketSignature structure.
//
// Added in macOS .
// Creates a CFSocket object using information from a CFSocketSignature structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCreateWithSocketSignature(_:_:_:_:_:)
func CFSocketCreateWithSocketSignature(allocator AllocatorRef, signature unsafe.Pointer, callBackTypes OptionFlags, callout SocketCallBack, context unsafe.Pointer) SocketRef {
	return _CFSocketCreateWithSocketSignature(allocator, signature, callBackTypes, callout, context)
}

// Disables the callback function of a CFSocket object for certain types of socket activity.
//
// Added in macOS .
// Disables the callback function of a CFSocket object for certain types of socket activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketDisableCallBacks(_:_:)
func CFSocketDisableCallBacks(s SocketRef, callBackTypes OptionFlags) {
	_CFSocketDisableCallBacks(s, callBackTypes)
}

// Enables the callback function of a CFSocket object for certain types of socket activity.
//
// Added in macOS .
// Enables the callback function of a CFSocket object for certain types of socket activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketEnableCallBacks(_:_:)
func CFSocketEnableCallBacks(s SocketRef, callBackTypes OptionFlags) {
	_CFSocketEnableCallBacks(s, callBackTypes)
}

// Returns the context information for a CFSocket object.
//
// Added in macOS .
// Returns the context information for a CFSocket object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetContext(_:_:)
func CFSocketGetContext(s SocketRef, context unsafe.Pointer) {
	_CFSocketGetContext(s, context)
}

// Returns the default port number with which to connect to a CFSocket name server.
//
// Added in macOS .
// Returns the default port number with which to connect to a CFSocket name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetDefaultNameRegistryPortNumber()
func CFSocketGetDefaultNameRegistryPortNumber() unsafe.Pointer {
	return _CFSocketGetDefaultNameRegistryPortNumber()
}

// Returns the native socket associated with a CFSocket object.
//
// Added in macOS .
// Returns the native socket associated with a CFSocket object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetNative(_:)
func CFSocketGetNative(s SocketRef) SocketNativeHandle {
	return _CFSocketGetNative(s)
}

// Returns flags that control certain behaviors of a CFSocket object.
//
// Added in macOS .
// Returns flags that control certain behaviors of a CFSocket object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetSocketFlags(_:)
func CFSocketGetSocketFlags(s SocketRef) OptionFlags {
	return _CFSocketGetSocketFlags(s)
}

// Returns the type identifier for the CFSocket opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFSocket opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketGetTypeID()
func CFSocketGetTypeID() TypeID {
	return _CFSocketGetTypeID()
}

// Invalidates a CFSocket object, stopping it from sending or receiving any more messages.
//
// Added in macOS .
// Invalidates a CFSocket object, stopping it from sending or receiving any more messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketInvalidate(_:)
func CFSocketInvalidate(s SocketRef) {
	_CFSocketInvalidate(s)
}

// Returns a Boolean value that indicates whether a CFSocket object is valid and able to send or receive messages.
//
// Added in macOS .
// Returns a Boolean value that indicates whether a CFSocket object is valid and able to send or receive messages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketIsValid(_:)
func CFSocketIsValid(s SocketRef) unsafe.Pointer {
	return _CFSocketIsValid(s)
}

// Registers a socket signature with a CFSocket name server.
//
// Added in macOS .
// Registers a socket signature with a CFSocket name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketRegisterSocketSignature(_:_:_:_:)
func CFSocketRegisterSocketSignature(nameServerSignature unsafe.Pointer, timeout TimeInterval, name StringRef, signature unsafe.Pointer) SocketError {
	return _CFSocketRegisterSocketSignature(nameServerSignature, timeout, name, signature)
}

// Registers a property-list value with a CFSocket name server.
//
// Added in macOS .
// Registers a property-list value with a CFSocket name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketRegisterValue(_:_:_:_:)
func CFSocketRegisterValue(nameServerSignature unsafe.Pointer, timeout TimeInterval, name StringRef, value PropertyListRef) SocketError {
	return _CFSocketRegisterValue(nameServerSignature, timeout, name, value)
}

// Sends data over a CFSocket object.
//
// Added in macOS .
// Sends data over a CFSocket object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketSendData(_:_:_:_:)
func CFSocketSendData(s SocketRef, address DataRef, data DataRef, timeout TimeInterval) SocketError {
	return _CFSocketSendData(s, address, data, timeout)
}

// Binds a local address to a CFSocket object and configures it for listening.
//
// Added in macOS .
// Binds a local address to a CFSocket object and configures it for listening.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetAddress(_:_:)
func CFSocketSetAddress(s SocketRef, address DataRef) SocketError {
	return _CFSocketSetAddress(s, address)
}

// Sets the default port number with which to connect to a CFSocket name server.
//
// Added in macOS .
// Sets the default port number with which to connect to a CFSocket name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetDefaultNameRegistryPortNumber(_:)
func CFSocketSetDefaultNameRegistryPortNumber(port unsafe.Pointer) {
	_CFSocketSetDefaultNameRegistryPortNumber(port)
}

// Sets flags that control certain behaviors of a CFSocket object.
//
// Added in macOS .
// Sets flags that control certain behaviors of a CFSocket object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketSetSocketFlags(_:_:)
func CFSocketSetSocketFlags(s SocketRef, flags OptionFlags) {
	_CFSocketSetSocketFlags(s, flags)
}

// Unregisters a value or socket signature with a CFSocket name server.
//
// Added in macOS .
// Unregisters a value or socket signature with a CFSocket name server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketUnregister(_:_:_:)
func CFSocketUnregister(nameServerSignature unsafe.Pointer, timeout TimeInterval, name StringRef) SocketError {
	return _CFSocketUnregister(nameServerSignature, timeout, name)
}

// Creates a bound pair of read and write streams.
//
// Added in macOS .
// Creates a bound pair of read and write streams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreateBoundPair(_:_:_:_:)
func CFStreamCreateBoundPair(alloc AllocatorRef, readStream unsafe.Pointer, writeStream unsafe.Pointer, transferBufferSize Index) {
	_CFStreamCreateBoundPair(alloc, readStream, writeStream, transferBufferSize)
}

// Creates readable and writable streams connected to a socket.
//
// Deprecated: This function was deprecated in macOS 26.1.
//
// Added in macOS 10.1.
// Creates readable and writable streams connected to a socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithPeerSocketSignature(_:_:_:_:)
func CFStreamCreatePairWithPeerSocketSignature(alloc AllocatorRef, signature unsafe.Pointer, readStream unsafe.Pointer, writeStream unsafe.Pointer) {
	_CFStreamCreatePairWithPeerSocketSignature(alloc, signature, readStream, writeStream)
}

// Creates readable and writable streams connected to a socket.
//
// Deprecated: This function was deprecated in macOS 26.1.
//
// Added in macOS 10.1.
// Creates readable and writable streams connected to a socket.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithSocket(_:_:_:_:)
func CFStreamCreatePairWithSocket(alloc AllocatorRef, sock SocketNativeHandle, readStream unsafe.Pointer, writeStream unsafe.Pointer) {
	_CFStreamCreatePairWithSocket(alloc, sock, readStream, writeStream)
}

// Creates readable and writable streams connected to a TCP/IP port of a particular host.
//
// Deprecated: This function was deprecated in macOS 26.1.
//
// Added in macOS 10.1.
// Creates readable and writable streams connected to a TCP/IP port of a particular host.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreatePairWithSocketToHost(_:_:_:_:_:)
func CFStreamCreatePairWithSocketToHost(alloc AllocatorRef, host StringRef, port unsafe.Pointer, readStream unsafe.Pointer, writeStream unsafe.Pointer) {
	_CFStreamCreatePairWithSocketToHost(alloc, host, port, readStream, writeStream)
}

// Appends the characters of a string to those of a CFMutableString object.
//
// Added in macOS .
// Appends the characters of a string to those of a CFMutableString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppend(_:_:)
func CFStringAppend(theString MutableStringRef, appendedString StringRef) {
	_CFStringAppend(theString, appendedString)
}

// Appends a C string to the character contents of a CFMutableString object.
//
// Added in macOS .
// Appends a C string to the character contents of a CFMutableString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendCString(_:_:_:)
func CFStringAppendCString(theString MutableStringRef, cStr unsafe.Pointer, encoding StringEncoding) {
	_CFStringAppendCString(theString, cStr, encoding)
}

// Appends a buffer of Unicode characters to the character contents of a CFMutableString object.
//
// Added in macOS .
// Appends a buffer of Unicode characters to the character contents of a CFMutableString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendCharacters(_:_:_:)
func CFStringAppendCharacters(theString MutableStringRef, chars unsafe.Pointer, numChars Index) {
	_CFStringAppendCharacters(theString, chars, numChars)
}

// Appends a formatted string to the character contents of a CFMutableString object.
//
// Added in macOS .
// Appends a formatted string to the character contents of a CFMutableString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendFormat
func CFStringAppendFormat(theString MutableStringRef, formatOptions DictionaryRef, format StringRef) {
	_CFStringAppendFormat(theString, formatOptions, format)
}

// Appends a formatted string to the character contents of a CFMutableString object.
//
// Added in macOS .
// Appends a formatted string to the character contents of a CFMutableString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendFormatAndArguments(_:_:_:_:)
func CFStringAppendFormatAndArguments(theString MutableStringRef, formatOptions DictionaryRef, format StringRef, arguments unsafe.Pointer) {
	_CFStringAppendFormatAndArguments(theString, formatOptions, format, arguments)
}

// Appends a Pascal string to the character contents of a CFMutableString object.
//
// Added in macOS .
// Appends a Pascal string to the character contents of a CFMutableString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringAppendPascalString(_:_:_:)
func CFStringAppendPascalString(theString MutableStringRef, pStr unsafe.Pointer, encoding StringEncoding) {
	_CFStringAppendPascalString(theString, pStr, encoding)
}

// Changes the first character in each word of a string to uppercase (if it is a lowercase alphabetical character).
//
// Added in macOS .
// Changes the first character in each word of a string to uppercase (if it is a lowercase alphabetical character).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCapitalize(_:_:)
func CFStringCapitalize(theString MutableStringRef, locale LocaleRef) {
	_CFStringCapitalize(theString, locale)
}

// Compares one string with another string.
//
// Added in macOS .
// Compares one string with another string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompare(_:_:_:)
func CFStringCompare(theString1 StringRef, theString2 StringRef, compareOptions StringCompareFlags) ComparisonResult {
	return _CFStringCompare(theString1, theString2, compareOptions)
}

// Compares a range of the characters in one string with that of another string.
//
// Added in macOS .
// Compares a range of the characters in one string with that of another string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareWithOptions(_:_:_:_:)
func CFStringCompareWithOptions(theString1 StringRef, theString2 StringRef, rangeToCompare Range, compareOptions StringCompareFlags) ComparisonResult {
	return _CFStringCompareWithOptions(theString1, theString2, rangeToCompare, compareOptions)
}

// Compares a range of the characters in one string with another string using a given locale.
//
// Added in macOS 10.5.
// Compares a range of the characters in one string with another string using a given locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareWithOptionsAndLocale(_:_:_:_:_:)
func CFStringCompareWithOptionsAndLocale(theString1 StringRef, theString2 StringRef, rangeToCompare Range, compareOptions StringCompareFlags, locale LocaleRef) ComparisonResult {
	return _CFStringCompareWithOptionsAndLocale(theString1, theString2, rangeToCompare, compareOptions, locale)
}

// Returns the name of the IANA registry “charset” that is the closest mapping to a specified string encoding.
//
// Added in macOS .
// Returns the name of the IANA registry “charset” that is the closest mapping to a specified string encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToIANACharSetName(_:)
func CFStringConvertEncodingToIANACharSetName(encoding StringEncoding) StringRef {
	return _CFStringConvertEncodingToIANACharSetName(encoding)
}

// Returns the Cocoa encoding constant that maps most closely to a given Core Foundation encoding constant.
//
// Added in macOS .
// Returns the Cocoa encoding constant that maps most closely to a given Core Foundation encoding constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToNSStringEncoding(_:)
func CFStringConvertEncodingToNSStringEncoding(encoding StringEncoding) unsafe.Pointer {
	return _CFStringConvertEncodingToNSStringEncoding(encoding)
}

// Returns the Windows codepage identifier that maps most closely to a given Core Foundation encoding constant.
//
// Added in macOS .
// Returns the Windows codepage identifier that maps most closely to a given Core Foundation encoding constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToWindowsCodepage(_:)
func CFStringConvertEncodingToWindowsCodepage(encoding StringEncoding) unsafe.Pointer {
	return _CFStringConvertEncodingToWindowsCodepage(encoding)
}

// Returns the Core Foundation encoding constant that is the closest mapping to a given IANA registry “charset” name.
//
// Added in macOS .
// Returns the Core Foundation encoding constant that is the closest mapping to a given IANA registry “charset” name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertIANACharSetNameToEncoding(_:)
func CFStringConvertIANACharSetNameToEncoding(theString StringRef) StringEncoding {
	return _CFStringConvertIANACharSetNameToEncoding(theString)
}

// Returns the Core Foundation encoding constant that is the closest mapping to a given Cocoa encoding.
//
// Added in macOS .
// Returns the Core Foundation encoding constant that is the closest mapping to a given Cocoa encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertNSStringEncodingToEncoding(_:)
func CFStringConvertNSStringEncodingToEncoding(encoding unsafe.Pointer) StringEncoding {
	return _CFStringConvertNSStringEncodingToEncoding(encoding)
}

// Returns the Core Foundation encoding constant that is the closest mapping to a given Windows codepage identifier.
//
// Added in macOS .
// Returns the Core Foundation encoding constant that is the closest mapping to a given Windows codepage identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertWindowsCodepageToEncoding(_:)
func CFStringConvertWindowsCodepageToEncoding(codepage unsafe.Pointer) StringEncoding {
	return _CFStringConvertWindowsCodepageToEncoding(codepage)
}

// Creates an array of CFString objects from a single CFString object.
//
// Added in macOS .
// Creates an array of CFString objects from a single CFString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateArrayBySeparatingStrings(_:_:_:)
func CFStringCreateArrayBySeparatingStrings(alloc AllocatorRef, theString StringRef, separatorString StringRef) ArrayRef {
	return _CFStringCreateArrayBySeparatingStrings(alloc, theString, separatorString)
}

// Searches a string for multiple occurrences of a substring and creates an array of ranges identifying the locations of these substrings within the target string.
//
// Added in macOS .
// Searches a string for multiple occurrences of a substring and creates an array of ranges identifying the locations of these substrings within the target string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateArrayWithFindResults(_:_:_:_:_:)
func CFStringCreateArrayWithFindResults(alloc AllocatorRef, theString StringRef, stringToFind StringRef, rangeToSearch Range, compareOptions StringCompareFlags) ArrayRef {
	return _CFStringCreateArrayWithFindResults(alloc, theString, stringToFind, rangeToSearch, compareOptions)
}

// Creates a single string from the individual CFString objects that comprise the elements of an array.
//
// Added in macOS .
// Creates a single string from the individual CFString objects that comprise the elements of an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateByCombiningStrings(_:_:_:)
func CFStringCreateByCombiningStrings(alloc AllocatorRef, theArray ArrayRef, separatorString StringRef) StringRef {
	return _CFStringCreateByCombiningStrings(alloc, theArray, separatorString)
}

// Creates an immutable copy of a string.
//
// Added in macOS .
// Creates an immutable copy of a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateCopy(_:_:)
func CFStringCreateCopy(alloc AllocatorRef, theString StringRef) StringRef {
	return _CFStringCreateCopy(alloc, theString)
}

// Creates an “external representation” of a CFString object, that is, a CFData object.
//
// Added in macOS .
// Creates an “external representation” of a CFString object, that is, a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateExternalRepresentation(_:_:_:_:)
func CFStringCreateExternalRepresentation(alloc AllocatorRef, theString StringRef, encoding StringEncoding, lossByte unsafe.Pointer) DataRef {
	return _CFStringCreateExternalRepresentation(alloc, theString, encoding, lossByte)
}

// Creates a string from its “external representation.”
//
// Added in macOS .
// Creates a string from its “external representation.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateFromExternalRepresentation(_:_:_:)
func CFStringCreateFromExternalRepresentation(alloc AllocatorRef, data DataRef, encoding StringEncoding) StringRef {
	return _CFStringCreateFromExternalRepresentation(alloc, data, encoding)
}

// Creates an empty CFMutableString object.
//
// Added in macOS .
// Creates an empty CFMutableString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutable(_:_:)
func CFStringCreateMutable(alloc AllocatorRef, maxLength Index) MutableStringRef {
	return _CFStringCreateMutable(alloc, maxLength)
}

// Creates a mutable copy of a string.
//
// Added in macOS .
// Creates a mutable copy of a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutableCopy(_:_:_:)
func CFStringCreateMutableCopy(alloc AllocatorRef, maxLength Index, theString StringRef) MutableStringRef {
	return _CFStringCreateMutableCopy(alloc, maxLength, theString)
}

// Creates a CFMutableString object whose Unicode character buffer is controlled externally.
//
// Added in macOS .
// Creates a CFMutableString object whose Unicode character buffer is controlled externally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateMutableWithExternalCharactersNoCopy(_:_:_:_:_:)
func CFStringCreateMutableWithExternalCharactersNoCopy(alloc AllocatorRef, chars unsafe.Pointer, numChars Index, capacity Index, externalCharactersAllocator AllocatorRef) MutableStringRef {
	return _CFStringCreateMutableWithExternalCharactersNoCopy(alloc, chars, numChars, capacity, externalCharactersAllocator)
}

// CFStringCreateStringWithValidatedFormat is a CoreFoundation function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateStringWithValidatedFormat
func CFStringCreateStringWithValidatedFormat(alloc AllocatorRef, formatOptions DictionaryRef, validFormatSpecifiers StringRef, format StringRef, errorPtr unsafe.Pointer) StringRef {
	return _CFStringCreateStringWithValidatedFormat(alloc, formatOptions, validFormatSpecifiers, format, errorPtr)
}

// CFStringCreateStringWithValidatedFormatAndArguments is a CoreFoundation function.
//
// Added in macOS 13.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateStringWithValidatedFormatAndArguments
func CFStringCreateStringWithValidatedFormatAndArguments(alloc AllocatorRef, formatOptions DictionaryRef, validFormatSpecifiers StringRef, format StringRef, arguments unsafe.Pointer, errorPtr unsafe.Pointer) StringRef {
	return _CFStringCreateStringWithValidatedFormatAndArguments(alloc, formatOptions, validFormatSpecifiers, format, arguments, errorPtr)
}

// Creates a string from a buffer containing characters in a specified encoding.
//
// Added in macOS .
// Creates a string from a buffer containing characters in a specified encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithBytes(_:_:_:_:_:)
func CFStringCreateWithBytes(alloc AllocatorRef, bytes unsafe.Pointer, numBytes Index, encoding StringEncoding, isExternalRepresentation unsafe.Pointer) StringRef {
	return _CFStringCreateWithBytes(alloc, bytes, numBytes, encoding, isExternalRepresentation)
}

// Creates a string from a buffer, containing characters in a specified encoding, that might serve as the backing store for the new string.
//
// Added in macOS .
// Creates a string from a buffer, containing characters in a specified encoding, that might serve as the backing store for the new string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithBytesNoCopy(_:_:_:_:_:_:)
func CFStringCreateWithBytesNoCopy(alloc AllocatorRef, bytes unsafe.Pointer, numBytes Index, encoding StringEncoding, isExternalRepresentation unsafe.Pointer, contentsDeallocator AllocatorRef) StringRef {
	return _CFStringCreateWithBytesNoCopy(alloc, bytes, numBytes, encoding, isExternalRepresentation, contentsDeallocator)
}

// Creates an immutable string from a C string.
//
// Added in macOS .
// Creates an immutable string from a C string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCString(_:_:_:)
func CFStringCreateWithCString(alloc AllocatorRef, cStr unsafe.Pointer, encoding StringEncoding) StringRef {
	return _CFStringCreateWithCString(alloc, cStr, encoding)
}

// Creates a CFString object from an external C string buffer that might serve as the backing store for the object.
//
// Added in macOS .
// Creates a CFString object from an external C string buffer that might serve as the backing store for the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCStringNoCopy(_:_:_:_:)
func CFStringCreateWithCStringNoCopy(alloc AllocatorRef, cStr unsafe.Pointer, encoding StringEncoding, contentsDeallocator AllocatorRef) StringRef {
	return _CFStringCreateWithCStringNoCopy(alloc, cStr, encoding, contentsDeallocator)
}

// Creates a string from a buffer of Unicode characters.
//
// Added in macOS .
// Creates a string from a buffer of Unicode characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCharacters(_:_:_:)
func CFStringCreateWithCharacters(alloc AllocatorRef, chars unsafe.Pointer, numChars Index) StringRef {
	return _CFStringCreateWithCharacters(alloc, chars, numChars)
}

// Creates a string from a buffer of Unicode characters that might serve as the backing store for the object.
//
// Added in macOS .
// Creates a string from a buffer of Unicode characters that might serve as the backing store for the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCharactersNoCopy(_:_:_:_:)
func CFStringCreateWithCharactersNoCopy(alloc AllocatorRef, chars unsafe.Pointer, numChars Index, contentsDeallocator AllocatorRef) StringRef {
	return _CFStringCreateWithCharactersNoCopy(alloc, chars, numChars, contentsDeallocator)
}

// Creates a CFString from a zero-terminated POSIX file system representation.
//
// Added in macOS .
// Creates a CFString from a zero-terminated POSIX file system representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFileSystemRepresentation(_:_:)
func CFStringCreateWithFileSystemRepresentation(alloc AllocatorRef, buffer unsafe.Pointer) StringRef {
	return _CFStringCreateWithFileSystemRepresentation(alloc, buffer)
}

// Creates an immutable string from a formatted string and a variable number of arguments.
//
// Added in macOS .
// Creates an immutable string from a formatted string and a variable number of arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFormat
func CFStringCreateWithFormat(alloc AllocatorRef, formatOptions DictionaryRef, format StringRef) StringRef {
	return _CFStringCreateWithFormat(alloc, formatOptions, format)
}

// Creates an immutable string from a formatted string and a variable number of arguments (specified in a parameter of type ).
//
// Added in macOS .
// Creates an immutable string from a formatted string and a variable number of arguments (specified in a parameter of type ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFormatAndArguments(_:_:_:_:)
func CFStringCreateWithFormatAndArguments(alloc AllocatorRef, formatOptions DictionaryRef, format StringRef, arguments unsafe.Pointer) StringRef {
	return _CFStringCreateWithFormatAndArguments(alloc, formatOptions, format, arguments)
}

// Creates an immutable CFString object from a Pascal string.
//
// Added in macOS .
// Creates an immutable CFString object from a Pascal string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithPascalString(_:_:_:)
func CFStringCreateWithPascalString(alloc AllocatorRef, pStr unsafe.Pointer, encoding StringEncoding) StringRef {
	return _CFStringCreateWithPascalString(alloc, pStr, encoding)
}

// Creates a CFString object from an external Pascal string buffer that might serve as the backing store for the object.
//
// Added in macOS .
// Creates a CFString object from an external Pascal string buffer that might serve as the backing store for the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithPascalStringNoCopy(_:_:_:_:)
func CFStringCreateWithPascalStringNoCopy(alloc AllocatorRef, pStr unsafe.Pointer, encoding StringEncoding, contentsDeallocator AllocatorRef) StringRef {
	return _CFStringCreateWithPascalStringNoCopy(alloc, pStr, encoding, contentsDeallocator)
}

// Creates an immutable string from a segment (substring) of an existing string.
//
// Added in macOS .
// Creates an immutable string from a segment (substring) of an existing string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithSubstring(_:_:_:)
func CFStringCreateWithSubstring(alloc AllocatorRef, str StringRef, range_ Range) StringRef {
	return _CFStringCreateWithSubstring(alloc, str, range_)
}

// Deletes a range of characters in a string.
//
// Added in macOS .
// Deletes a range of characters in a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringDelete(_:_:)
func CFStringDelete(theString MutableStringRef, range_ Range) {
	_CFStringDelete(theString, range_)
}

// Searches for a substring within a string and, if it is found, yields the range of the substring within the object’s characters.
//
// Added in macOS .
// Searches for a substring within a string and, if it is found, yields the range of the substring within the object’s characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFind(_:_:_:)
func CFStringFind(theString StringRef, stringToFind StringRef, compareOptions StringCompareFlags) Range {
	return _CFStringFind(theString, stringToFind, compareOptions)
}

// Replaces all occurrences of a substring within a given range.
//
// Added in macOS .
// Replaces all occurrences of a substring within a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindAndReplace(_:_:_:_:_:)
func CFStringFindAndReplace(theString MutableStringRef, stringToFind StringRef, replacementString StringRef, rangeToSearch Range, compareOptions StringCompareFlags) Index {
	return _CFStringFindAndReplace(theString, stringToFind, replacementString, rangeToSearch, compareOptions)
}

// Query the range of the first character contained in the specified character set.
//
// Added in macOS .
// Query the range of the first character contained in the specified character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindCharacterFromSet(_:_:_:_:_:)
func CFStringFindCharacterFromSet(theString StringRef, theSet CharacterSetRef, rangeToSearch Range, searchOptions StringCompareFlags, result unsafe.Pointer) unsafe.Pointer {
	return _CFStringFindCharacterFromSet(theString, theSet, rangeToSearch, searchOptions, result)
}

// Searches for a substring within a range of the characters represented by a string and, if the substring is found, returns its range within the object’s characters.
//
// Added in macOS .
// Searches for a substring within a range of the characters represented by a string and, if the substring is found, returns its range within the object’s characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindWithOptions(_:_:_:_:_:)
func CFStringFindWithOptions(theString StringRef, stringToFind StringRef, rangeToSearch Range, searchOptions StringCompareFlags, result unsafe.Pointer) unsafe.Pointer {
	return _CFStringFindWithOptions(theString, stringToFind, rangeToSearch, searchOptions, result)
}

// Returns a Boolean value that indicates whether a given string was found in a given source string.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether a given string was found in a given source string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindWithOptionsAndLocale(_:_:_:_:_:_:)
func CFStringFindWithOptionsAndLocale(theString StringRef, stringToFind StringRef, rangeToSearch Range, searchOptions StringCompareFlags, locale LocaleRef, result unsafe.Pointer) unsafe.Pointer {
	return _CFStringFindWithOptionsAndLocale(theString, stringToFind, rangeToSearch, searchOptions, locale, result)
}

// Folds a given string into the form specified by optional flags.
//
// Added in macOS 10.5.
// Folds a given string into the form specified by optional flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFold(_:_:_:)
func CFStringFold(theString MutableStringRef, theFlags StringCompareFlags, theLocale LocaleRef) {
	_CFStringFold(theString, theFlags, theLocale)
}

// Fetches a range of the characters from a string into a byte buffer after converting the characters to a specified encoding.
//
// Added in macOS .
// Fetches a range of the characters from a string into a byte buffer after converting the characters to a specified encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetBytes(_:_:_:_:_:_:_:_:)
func CFStringGetBytes(theString StringRef, range_ Range, encoding StringEncoding, lossByte unsafe.Pointer, isExternalRepresentation unsafe.Pointer, buffer unsafe.Pointer, maxBufLen Index, usedBufLen unsafe.Pointer) Index {
	return _CFStringGetBytes(theString, range_, encoding, lossByte, isExternalRepresentation, buffer, maxBufLen, usedBufLen)
}

// Copies the character contents of a string to a local C string buffer after converting the characters to a given encoding.
//
// Added in macOS .
// Copies the character contents of a string to a local C string buffer after converting the characters to a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCString(_:_:_:_:)
func CFStringGetCString(theString StringRef, buffer unsafe.Pointer, bufferSize Index, encoding StringEncoding) unsafe.Pointer {
	return _CFStringGetCString(theString, buffer, bufferSize, encoding)
}

// Quickly obtains a pointer to a C-string buffer containing the characters of a string in a given encoding.
//
// Added in macOS .
// Quickly obtains a pointer to a C-string buffer containing the characters of a string in a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCStringPtr(_:_:)
func CFStringGetCStringPtr(theString StringRef, encoding StringEncoding) unsafe.Pointer {
	return _CFStringGetCStringPtr(theString, encoding)
}

// Returns the Unicode character at a specified location in a string.
//
// Added in macOS .
// Returns the Unicode character at a specified location in a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharacterAtIndex(_:_:)
func CFStringGetCharacterAtIndex(theString StringRef, idx Index) unsafe.Pointer {
	return _CFStringGetCharacterAtIndex(theString, idx)
}

// Copies a range of the Unicode characters from a string to a user-provided buffer.
//
// Added in macOS .
// Copies a range of the Unicode characters from a string to a user-provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharacters(_:_:_:)
func CFStringGetCharacters(theString StringRef, range_ Range, buffer unsafe.Pointer) {
	_CFStringGetCharacters(theString, range_, buffer)
}

// Quickly obtains a pointer to the contents of a string as a buffer of Unicode characters.
//
// Added in macOS .
// Quickly obtains a pointer to the contents of a string as a buffer of Unicode characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharactersPtr(_:)
func CFStringGetCharactersPtr(theString StringRef) unsafe.Pointer {
	return _CFStringGetCharactersPtr(theString)
}

// Returns the primary value represented by a string.
//
// Added in macOS .
// Returns the primary value represented by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetDoubleValue(_:)
func CFStringGetDoubleValue(str StringRef) float64 {
	return _CFStringGetDoubleValue(str)
}

// Returns for a CFString object the character encoding that requires the least conversion time.
//
// Added in macOS .
// Returns for a CFString object the character encoding that requires the least conversion time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetFastestEncoding(_:)
func CFStringGetFastestEncoding(theString StringRef) StringEncoding {
	return _CFStringGetFastestEncoding(theString)
}

// Extracts the contents of a string as a -terminated 8-bit string appropriate for passing to POSIX APIs.
//
// Added in macOS .
// Extracts the contents of a string as a -terminated 8-bit string appropriate for passing to POSIX APIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetFileSystemRepresentation(_:_:_:)
func CFStringGetFileSystemRepresentation(string_ StringRef, buffer unsafe.Pointer, maxBufLen Index) unsafe.Pointer {
	return _CFStringGetFileSystemRepresentation(string_, buffer, maxBufLen)
}

// Retrieve the first potential hyphenation location found before the specified location.
//
// Added in macOS 10.7.
// Retrieve the first potential hyphenation location found before the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetHyphenationLocationBeforeIndex(_:_:_:_:_:_:)
func CFStringGetHyphenationLocationBeforeIndex(string_ StringRef, location Index, limitRange Range, options OptionFlags, locale LocaleRef, character unsafe.Pointer) Index {
	return _CFStringGetHyphenationLocationBeforeIndex(string_, location, limitRange, options, locale, character)
}

// Returns the integer value represented by a string.
//
// Added in macOS .
// Returns the integer value represented by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetIntValue(_:)
func CFStringGetIntValue(str StringRef) unsafe.Pointer {
	return _CFStringGetIntValue(str)
}

// Returns the number (in terms of UTF-16 code pairs) of Unicode characters in a string.
//
// Added in macOS .
// Returns the number (in terms of UTF-16 code pairs) of Unicode characters in a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetLength(_:)
func CFStringGetLength(theString StringRef) Index {
	return _CFStringGetLength(theString)
}

// Given a range of characters in a string, obtains the line bounds—that is, the indexes of the first character and the final characters of the lines containing the range.
//
// Added in macOS .
// Given a range of characters in a string, obtains the line bounds—that is, the indexes of the first character and the final characters of the lines containing the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetLineBounds(_:_:_:_:_:)
func CFStringGetLineBounds(theString StringRef, range_ Range, lineBeginIndex unsafe.Pointer, lineEndIndex unsafe.Pointer, contentsEndIndex unsafe.Pointer) {
	_CFStringGetLineBounds(theString, range_, lineBeginIndex, lineEndIndex, contentsEndIndex)
}

// Returns a pointer to a list of string encodings supported by the current system.
//
// Added in macOS .
// Returns a pointer to a list of string encodings supported by the current system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetListOfAvailableEncodings()
func CFStringGetListOfAvailableEncodings() unsafe.Pointer {
	return _CFStringGetListOfAvailableEncodings()
}

// Returns the maximum number of bytes a string of a specified length (in Unicode characters) will take up if encoded in a specified encoding.
//
// Added in macOS .
// Returns the maximum number of bytes a string of a specified length (in Unicode characters) will take up if encoded in a specified encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMaximumSizeForEncoding(_:_:)
func CFStringGetMaximumSizeForEncoding(length Index, encoding StringEncoding) Index {
	return _CFStringGetMaximumSizeForEncoding(length, encoding)
}

// Determines the upper bound on the number of bytes required to hold the file system representation of the string.
//
// Added in macOS .
// Determines the upper bound on the number of bytes required to hold the file system representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMaximumSizeOfFileSystemRepresentation(_:)
func CFStringGetMaximumSizeOfFileSystemRepresentation(string_ StringRef) Index {
	return _CFStringGetMaximumSizeOfFileSystemRepresentation(string_)
}

// Returns the most compatible Mac OS script value for the given input encoding.
//
// Added in macOS .
// Returns the most compatible Mac OS script value for the given input encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMostCompatibleMacStringEncoding(_:)
func CFStringGetMostCompatibleMacStringEncoding(encoding StringEncoding) StringEncoding {
	return _CFStringGetMostCompatibleMacStringEncoding(encoding)
}

// Returns the canonical name of a specified string encoding.
//
// Added in macOS .
// Returns the canonical name of a specified string encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetNameOfEncoding(_:)
func CFStringGetNameOfEncoding(encoding StringEncoding) StringRef {
	return _CFStringGetNameOfEncoding(encoding)
}

// Given a range of characters in a string, obtains the paragraph bounds—that is, the indexes of the first character and the final characters of the paragraph(s) containing the range.
//
// Added in macOS 10.5.
// Given a range of characters in a string, obtains the paragraph bounds—that is, the indexes of the first character and the final characters of the paragraph(s) containing the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetParagraphBounds(_:_:_:_:_:)
func CFStringGetParagraphBounds(string_ StringRef, range_ Range, parBeginIndex unsafe.Pointer, parEndIndex unsafe.Pointer, contentsEndIndex unsafe.Pointer) {
	_CFStringGetParagraphBounds(string_, range_, parBeginIndex, parEndIndex, contentsEndIndex)
}

// Copies the character contents of a CFString object to a local Pascal string buffer after converting the characters to a requested encoding.
//
// Added in macOS .
// Copies the character contents of a CFString object to a local Pascal string buffer after converting the characters to a requested encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetPascalString(_:_:_:_:)
func CFStringGetPascalString(theString StringRef, buffer unsafe.Pointer, bufferSize Index, encoding StringEncoding) unsafe.Pointer {
	return _CFStringGetPascalString(theString, buffer, bufferSize, encoding)
}

// Quickly obtains a pointer to a Pascal buffer containing the characters of a string in a given encoding.
//
// Added in macOS .
// Quickly obtains a pointer to a Pascal buffer containing the characters of a string in a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetPascalStringPtr(_:_:)
func CFStringGetPascalStringPtr(theString StringRef, encoding StringEncoding) unsafe.Pointer {
	return _CFStringGetPascalStringPtr(theString, encoding)
}

// Returns the range of the composed character sequence at a specified index.
//
// Added in macOS .
// Returns the range of the composed character sequence at a specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetRangeOfComposedCharactersAtIndex(_:_:)
func CFStringGetRangeOfComposedCharactersAtIndex(theString StringRef, theIndex Index) Range {
	return _CFStringGetRangeOfComposedCharactersAtIndex(theString, theIndex)
}

// Returns the smallest encoding on the current system for the character contents of a string.
//
// Added in macOS .
// Returns the smallest encoding on the current system for the character contents of a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetSmallestEncoding(_:)
func CFStringGetSmallestEncoding(theString StringRef) StringEncoding {
	return _CFStringGetSmallestEncoding(theString)
}

// Returns the default encoding used by the operating system when it creates strings.
//
// Added in macOS .
// Returns the default encoding used by the operating system when it creates strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetSystemEncoding()
func CFStringGetSystemEncoding() StringEncoding {
	return _CFStringGetSystemEncoding()
}

// Returns the type identifier for the CFString opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFString opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetTypeID()
func CFStringGetTypeID() TypeID {
	return _CFStringGetTypeID()
}

// Determines if the character data of a string begin with a specified sequence of characters.
//
// Added in macOS .
// Determines if the character data of a string begin with a specified sequence of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringHasPrefix(_:_:)
func CFStringHasPrefix(theString StringRef, prefix StringRef) unsafe.Pointer {
	return _CFStringHasPrefix(theString, prefix)
}

// Determines if a string ends with a specified sequence of characters.
//
// Added in macOS .
// Determines if a string ends with a specified sequence of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringHasSuffix(_:_:)
func CFStringHasSuffix(theString StringRef, suffix StringRef) unsafe.Pointer {
	return _CFStringHasSuffix(theString, suffix)
}

// Inserts a string at a specified location in the character buffer of a CFMutableString object.
//
// Added in macOS .
// Inserts a string at a specified location in the character buffer of a CFMutableString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringInsert(_:_:_:)
func CFStringInsert(str MutableStringRef, idx Index, insertedStr StringRef) {
	_CFStringInsert(str, idx, insertedStr)
}

// Determines whether a given Core Foundation string encoding is available on the current system.
//
// Added in macOS .
// Determines whether a given Core Foundation string encoding is available on the current system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringIsEncodingAvailable(_:)
func CFStringIsEncodingAvailable(encoding StringEncoding) unsafe.Pointer {
	return _CFStringIsEncodingAvailable(encoding)
}

// Returns a Boolean value that indicates whether hyphenation data is available.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether hyphenation data is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringIsHyphenationAvailableForLocale(_:)
func CFStringIsHyphenationAvailableForLocale(locale LocaleRef) unsafe.Pointer {
	return _CFStringIsHyphenationAvailableForLocale(locale)
}

// Changes all uppercase alphabetical characters in a CFMutableString to lowercase.
//
// Added in macOS .
// Changes all uppercase alphabetical characters in a CFMutableString to lowercase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringLowercase(_:_:)
func CFStringLowercase(theString MutableStringRef, locale LocaleRef) {
	_CFStringLowercase(theString, locale)
}

// Normalizes the string into the specified form as described in Unicode Technical Report #15.
//
// Added in macOS .
// Normalizes the string into the specified form as described in Unicode Technical Report #15.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalize(_:_:)
func CFStringNormalize(theString MutableStringRef, theForm StringNormalizationForm) {
	_CFStringNormalize(theString, theForm)
}

// Enlarges a string, padding it with specified characters, or truncates the string.
//
// Added in macOS .
// Enlarges a string, padding it with specified characters, or truncates the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringPad(_:_:_:_:)
func CFStringPad(theString MutableStringRef, padString StringRef, length Index, indexIntoPad Index) {
	_CFStringPad(theString, padString, length, indexIntoPad)
}

// Replaces part of the character contents of a CFMutableString object with another string.
//
// Added in macOS .
// Replaces part of the character contents of a CFMutableString object with another string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringReplace(_:_:_:)
func CFStringReplace(theString MutableStringRef, range_ Range, replacement StringRef) {
	_CFStringReplace(theString, range_, replacement)
}

// Replaces all characters of a CFMutableString object with other characters.
//
// Added in macOS .
// Replaces all characters of a CFMutableString object with other characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringReplaceAll(_:_:)
func CFStringReplaceAll(theString MutableStringRef, replacement StringRef) {
	_CFStringReplaceAll(theString, replacement)
}

// Notifies a CFMutableString object that its external backing store of Unicode characters has changed.
//
// Added in macOS .
// Notifies a CFMutableString object that its external backing store of Unicode characters has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringSetExternalCharactersNoCopy(_:_:_:_:)
func CFStringSetExternalCharactersNoCopy(theString MutableStringRef, chars unsafe.Pointer, length Index, capacity Index) {
	_CFStringSetExternalCharactersNoCopy(theString, chars, length, capacity)
}

// Advances the tokenizer to the next token and sets that as the current token.
//
// Added in macOS 10.5.
// Advances the tokenizer to the next token and sets that as the current token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerAdvanceToNextToken(_:)
func CFStringTokenizerAdvanceToNextToken(tokenizer StringTokenizerRef) StringTokenizerTokenType {
	return _CFStringTokenizerAdvanceToNextToken(tokenizer)
}

// Guesses a language of a given string and returns the guess as a BCP 47 string.
//
// Added in macOS 10.5.
// Guesses a language of a given string and returns the guess as a BCP 47 string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCopyBestStringLanguage(_:_:)
func CFStringTokenizerCopyBestStringLanguage(string_ StringRef, range_ Range) StringRef {
	return _CFStringTokenizerCopyBestStringLanguage(string_, range_)
}

// Returns a given attribute of the current token.
//
// Added in macOS 10.5.
// Returns a given attribute of the current token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCopyCurrentTokenAttribute(_:_:)
func CFStringTokenizerCopyCurrentTokenAttribute(tokenizer StringTokenizerRef, attribute OptionFlags) TypeRef {
	return _CFStringTokenizerCopyCurrentTokenAttribute(tokenizer, attribute)
}

// Returns a tokenizer for a given string.
//
// Added in macOS 10.5.
// Returns a tokenizer for a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerCreate(_:_:_:_:_:)
func CFStringTokenizerCreate(alloc AllocatorRef, string_ StringRef, range_ Range, options OptionFlags, locale LocaleRef) StringTokenizerRef {
	return _CFStringTokenizerCreate(alloc, string_, range_, options, locale)
}

// Retrieves the subtokens or derived subtokens contained in the compound token.
//
// Added in macOS 10.5.
// Retrieves the subtokens or derived subtokens contained in the compound token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetCurrentSubTokens(_:_:_:_:)
func CFStringTokenizerGetCurrentSubTokens(tokenizer StringTokenizerRef, ranges unsafe.Pointer, maxRangeLength Index, derivedSubTokens MutableArrayRef) Index {
	return _CFStringTokenizerGetCurrentSubTokens(tokenizer, ranges, maxRangeLength, derivedSubTokens)
}

// Returns the range of the current token.
//
// Added in macOS 10.5.
// Returns the range of the current token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetCurrentTokenRange(_:)
func CFStringTokenizerGetCurrentTokenRange(tokenizer StringTokenizerRef) Range {
	return _CFStringTokenizerGetCurrentTokenRange(tokenizer)
}

// Returns the type ID for CFStringTokenizer.
//
// Added in macOS 10.5.
// Returns the type ID for CFStringTokenizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGetTypeID()
func CFStringTokenizerGetTypeID() TypeID {
	return _CFStringTokenizerGetTypeID()
}

// Finds a token that includes the character at a given index, and set it as the current token.
//
// Added in macOS 10.5.
// Finds a token that includes the character at a given index, and set it as the current token.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerGoToTokenAtIndex(_:_:)
func CFStringTokenizerGoToTokenAtIndex(tokenizer StringTokenizerRef, index Index) StringTokenizerTokenType {
	return _CFStringTokenizerGoToTokenAtIndex(tokenizer, index)
}

// Sets the string for a tokenizer.
//
// Added in macOS 10.5.
// Sets the string for a tokenizer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerSetString(_:_:_:)
func CFStringTokenizerSetString(tokenizer StringTokenizerRef, string_ StringRef, range_ Range) {
	_CFStringTokenizerSetString(tokenizer, string_, range_)
}

// Perform in-place transliteration on a mutable string.
//
// Added in macOS .
// Perform in-place transliteration on a mutable string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTransform(_:_:_:_:)
func CFStringTransform(string_ MutableStringRef, range_ unsafe.Pointer, transform StringRef, reverse unsafe.Pointer) unsafe.Pointer {
	return _CFStringTransform(string_, range_, transform, reverse)
}

// Trims a specified substring from the beginning and end of a CFMutableString object.
//
// Added in macOS .
// Trims a specified substring from the beginning and end of a CFMutableString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTrim(_:_:)
func CFStringTrim(theString MutableStringRef, trimString StringRef) {
	_CFStringTrim(theString, trimString)
}

// Trims whitespace from the beginning and end of a CFMutableString object.
//
// Added in macOS .
// Trims whitespace from the beginning and end of a CFMutableString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTrimWhitespace(_:)
func CFStringTrimWhitespace(theString MutableStringRef) {
	_CFStringTrimWhitespace(theString)
}

// Changes all lowercase alphabetical characters in a CFMutableString object to uppercase.
//
// Added in macOS .
// Changes all lowercase alphabetical characters in a CFMutableString object to uppercase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringUppercase(_:_:)
func CFStringUppercase(theString MutableStringRef, locale LocaleRef) {
	_CFStringUppercase(theString, locale)
}

// Returns the abbreviation of a time zone at a specified date.
//
// Added in macOS .
// Returns the abbreviation of a time zone at a specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyAbbreviation(_:_:)
func CFTimeZoneCopyAbbreviation(tz TimeZoneRef, at AbsoluteTime) StringRef {
	return _CFTimeZoneCopyAbbreviation(tz, at)
}

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names.
//
// Added in macOS .
// Returns a dictionary holding the mappings of time zone abbreviations to time zone names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyAbbreviationDictionary()
func CFTimeZoneCopyAbbreviationDictionary() DictionaryRef {
	return _CFTimeZoneCopyAbbreviationDictionary()
}

// Returns the default time zone set for your application.
//
// Added in macOS .
// Returns the default time zone set for your application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyDefault()
func CFTimeZoneCopyDefault() TimeZoneRef {
	return _CFTimeZoneCopyDefault()
}

// Returns an array of strings containing the names of all the time zones known to the system.
//
// Added in macOS .
// Returns an array of strings containing the names of all the time zones known to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyKnownNames()
func CFTimeZoneCopyKnownNames() ArrayRef {
	return _CFTimeZoneCopyKnownNames()
}

// Returns the localized name of a given time zone.
//
// Added in macOS 10.5.
// Returns the localized name of a given time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyLocalizedName(_:_:_:)
func CFTimeZoneCopyLocalizedName(tz TimeZoneRef, style TimeZoneNameStyle, locale LocaleRef) StringRef {
	return _CFTimeZoneCopyLocalizedName(tz, style, locale)
}

// Returns the time zone currently used by the system.
//
// Added in macOS .
// Returns the time zone currently used by the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopySystem()
func CFTimeZoneCopySystem() TimeZoneRef {
	return _CFTimeZoneCopySystem()
}

// Creates a time zone with a given name and data.
//
// Added in macOS .
// Creates a time zone with a given name and data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreate(_:_:_:)
func CFTimeZoneCreate(allocator AllocatorRef, name StringRef, data DataRef) TimeZoneRef {
	return _CFTimeZoneCreate(allocator, name, data)
}

// Returns the time zone object identified by a given name or abbreviation.
//
// Added in macOS .
// Returns the time zone object identified by a given name or abbreviation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreateWithName(_:_:_:)
func CFTimeZoneCreateWithName(allocator AllocatorRef, name StringRef, tryAbbrev unsafe.Pointer) TimeZoneRef {
	return _CFTimeZoneCreateWithName(allocator, name, tryAbbrev)
}

// Returns a time zone object for the specified time interval offset from Greenwich Mean Time (GMT).
//
// Added in macOS .
// Returns a time zone object for the specified time interval offset from Greenwich Mean Time (GMT).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreateWithTimeIntervalFromGMT(_:_:)
func CFTimeZoneCreateWithTimeIntervalFromGMT(allocator AllocatorRef, ti TimeInterval) TimeZoneRef {
	return _CFTimeZoneCreateWithTimeIntervalFromGMT(allocator, ti)
}

// Returns the data that stores the information used by a time zone.
//
// Added in macOS .
// Returns the data that stores the information used by a time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetData(_:)
func CFTimeZoneGetData(tz TimeZoneRef) DataRef {
	return _CFTimeZoneGetData(tz)
}

// Returns the daylight saving time offset for a time zone at a given time.
//
// Added in macOS 10.5.
// Returns the daylight saving time offset for a time zone at a given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetDaylightSavingTimeOffset(_:_:)
func CFTimeZoneGetDaylightSavingTimeOffset(tz TimeZoneRef, at AbsoluteTime) TimeInterval {
	return _CFTimeZoneGetDaylightSavingTimeOffset(tz, at)
}

// Returns the geopolitical region name that identifies a given time zone.
//
// Added in macOS .
// Returns the geopolitical region name that identifies a given time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetName(_:)
func CFTimeZoneGetName(tz TimeZoneRef) StringRef {
	return _CFTimeZoneGetName(tz)
}

// Returns the time in a given time zone of the next daylight saving time transition after a given time.
//
// Added in macOS 10.5.
// Returns the time in a given time zone of the next daylight saving time transition after a given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetNextDaylightSavingTimeTransition(_:_:)
func CFTimeZoneGetNextDaylightSavingTimeTransition(tz TimeZoneRef, at AbsoluteTime) AbsoluteTime {
	return _CFTimeZoneGetNextDaylightSavingTimeTransition(tz, at)
}

// Returns the difference in seconds between the receiver and Greenwich Mean Time (GMT) at the specified date.
//
// Added in macOS .
// Returns the difference in seconds between the receiver and Greenwich Mean Time (GMT) at the specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetSecondsFromGMT(_:_:)
func CFTimeZoneGetSecondsFromGMT(tz TimeZoneRef, at AbsoluteTime) TimeInterval {
	return _CFTimeZoneGetSecondsFromGMT(tz, at)
}

// Returns the type identifier for the CFTimeZone opaque type.
//
// Added in macOS .
// Returns the type identifier for the CFTimeZone opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetTypeID()
func CFTimeZoneGetTypeID() TypeID {
	return _CFTimeZoneGetTypeID()
}

// Returns whether or not a time zone is in daylight savings time at a specified date.
//
// Added in macOS .
// Returns whether or not a time zone is in daylight savings time at a specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneIsDaylightSavingTime(_:_:)
func CFTimeZoneIsDaylightSavingTime(tz TimeZoneRef, at AbsoluteTime) unsafe.Pointer {
	return _CFTimeZoneIsDaylightSavingTime(tz, at)
}

// Clears the previously determined system time zone, if any.
//
// Added in macOS .
// Clears the previously determined system time zone, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneResetSystem()
func CFTimeZoneResetSystem() {
	_CFTimeZoneResetSystem()
}

// Sets the abbreviation dictionary to a given dictionary.
//
// Added in macOS .
// Sets the abbreviation dictionary to a given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneSetAbbreviationDictionary(_:)
func CFTimeZoneSetAbbreviationDictionary(dict DictionaryRef) {
	_CFTimeZoneSetAbbreviationDictionary(dict)
}

// Sets the default time zone for your application the given time zone.
//
// Added in macOS .
// Sets the default time zone for your application the given time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneSetDefault(_:)
func CFTimeZoneSetDefault(tz TimeZoneRef) {
	_CFTimeZoneSetDefault(tz)
}

// Adds a new child to a tree as the last in its list of children.
//
// Added in macOS .
// Adds a new child to a tree as the last in its list of children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeAppendChild(_:_:)
func CFTreeAppendChild(tree TreeRef, newChild TreeRef) {
	_CFTreeAppendChild(tree, newChild)
}

// Calls a function once for each immediate child of a tree.
//
// Added in macOS .
// Calls a function once for each immediate child of a tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeApplyFunctionToChildren(_:_:_:)
func CFTreeApplyFunctionToChildren(tree TreeRef, applier TreeApplierFunction, context unsafe.Pointer) {
	_CFTreeApplyFunctionToChildren(tree, applier, context)
}

// Creates a new CFTree object.
//
// Added in macOS .
// Creates a new CFTree object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeCreate(_:_:)
func CFTreeCreate(allocator AllocatorRef, context unsafe.Pointer) TreeRef {
	return _CFTreeCreate(allocator, context)
}

// Returns the root tree of a given tree.
//
// Added in macOS .
// Returns the root tree of a given tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeFindRoot(_:)
func CFTreeFindRoot(tree TreeRef) TreeRef {
	return _CFTreeFindRoot(tree)
}

// Returns the child of a tree at the specified index.
//
// Added in macOS .
// Returns the child of a tree at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildAtIndex(_:_:)
func CFTreeGetChildAtIndex(tree TreeRef, idx Index) TreeRef {
	return _CFTreeGetChildAtIndex(tree, idx)
}

// Returns the number of children in a tree.
//
// Added in macOS .
// Returns the number of children in a tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildCount(_:)
func CFTreeGetChildCount(tree TreeRef) Index {
	return _CFTreeGetChildCount(tree)
}

// Fills a buffer with children from the tree.
//
// Added in macOS .
// Fills a buffer with children from the tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetChildren(_:_:)
func CFTreeGetChildren(tree TreeRef, children unsafe.Pointer) {
	_CFTreeGetChildren(tree, children)
}

// Returns the context of the specified tree.
//
// Added in macOS .
// Returns the context of the specified tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetContext(_:_:)
func CFTreeGetContext(tree TreeRef, context unsafe.Pointer) {
	_CFTreeGetContext(tree, context)
}

// Returns the first child of a tree.
//
// Added in macOS .
// Returns the first child of a tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetFirstChild(_:)
func CFTreeGetFirstChild(tree TreeRef) TreeRef {
	return _CFTreeGetFirstChild(tree)
}

// Returns the next sibling, adjacent to a given tree, in the parent’s children list.
//
// Added in macOS .
// Returns the next sibling, adjacent to a given tree, in the parent’s children list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetNextSibling(_:)
func CFTreeGetNextSibling(tree TreeRef) TreeRef {
	return _CFTreeGetNextSibling(tree)
}

// Returns the parent of a given tree.
//
// Added in macOS .
// Returns the parent of a given tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetParent(_:)
func CFTreeGetParent(tree TreeRef) TreeRef {
	return _CFTreeGetParent(tree)
}

// Returns the type identifier of the CFTree opaque type.
//
// Added in macOS .
// Returns the type identifier of the CFTree opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeGetTypeID()
func CFTreeGetTypeID() TypeID {
	return _CFTreeGetTypeID()
}

// Inserts a new sibling after a given tree.
//
// Added in macOS .
// Inserts a new sibling after a given tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeInsertSibling(_:_:)
func CFTreeInsertSibling(tree TreeRef, newSibling TreeRef) {
	_CFTreeInsertSibling(tree, newSibling)
}

// Adds a new child to the specified tree as the first in its list of children.
//
// Added in macOS .
// Adds a new child to the specified tree as the first in its list of children.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreePrependChild(_:_:)
func CFTreePrependChild(tree TreeRef, newChild TreeRef) {
	_CFTreePrependChild(tree, newChild)
}

// Removes a tree from its parent.
//
// Added in macOS .
// Removes a tree from its parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeRemove(_:)
func CFTreeRemove(tree TreeRef) {
	_CFTreeRemove(tree)
}

// Removes all the children of a tree.
//
// Added in macOS .
// Removes all the children of a tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeRemoveAllChildren(_:)
func CFTreeRemoveAllChildren(tree TreeRef) {
	_CFTreeRemoveAllChildren(tree)
}

// Replaces the context of a tree by releasing the old information pointer and retaining the new one.
//
// Added in macOS .
// Replaces the context of a tree by releasing the old information pointer and retaining the new one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeSetContext(_:_:)
func CFTreeSetContext(tree TreeRef, context unsafe.Pointer) {
	_CFTreeSetContext(tree, context)
}

// Sorts the immediate children of a tree using a specified comparator function.
//
// Added in macOS .
// Sorts the immediate children of a tree using a specified comparator function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeSortChildren(_:_:_:)
func CFTreeSortChildren(tree TreeRef, comparator ComparatorFunction, context unsafe.Pointer) {
	_CFTreeSortChildren(tree, comparator, context)
}

// Determines if the given URL conforms to RFC 1808 and therefore can be decomposed.
//
// Added in macOS .
// Determines if the given URL conforms to RFC 1808 and therefore can be decomposed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCanBeDecomposed(_:)
func CFURLCanBeDecomposed(anURL URLRef) unsafe.Pointer {
	return _CFURLCanBeDecomposed(anURL)
}

// Removes all cached resource values and temporary resource values from the URL object.
//
// Added in macOS 10.6.
// Removes all cached resource values and temporary resource values from the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLClearResourcePropertyCache(_:)
func CFURLClearResourcePropertyCache(url URLRef) {
	_CFURLClearResourcePropertyCache(url)
}

// Removes the cached resource value identified by a given key from the URL object.
//
// Added in macOS 10.6.
// Removes the cached resource value identified by a given key from the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLClearResourcePropertyCacheForKey(_:_:)
func CFURLClearResourcePropertyCacheForKey(url URLRef, key StringRef) {
	_CFURLClearResourcePropertyCacheForKey(url, key)
}

// Creates a new object by resolving the relative portion of a URL against its base.
//
// Added in macOS .
// Creates a new object by resolving the relative portion of a URL against its base.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyAbsoluteURL(_:)
func CFURLCopyAbsoluteURL(relativeURL URLRef) URLRef {
	return _CFURLCopyAbsoluteURL(relativeURL)
}

// Returns the path portion of a given URL.
//
// Added in macOS .
// Returns the path portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyFileSystemPath(_:_:)
func CFURLCopyFileSystemPath(anURL URLRef, pathStyle URLPathStyle) StringRef {
	return _CFURLCopyFileSystemPath(anURL, pathStyle)
}

// Returns the fragment from a given URL.
//
// Added in macOS .
// Returns the fragment from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyFragment(_:_:)
func CFURLCopyFragment(anURL URLRef, charactersToLeaveEscaped StringRef) StringRef {
	return _CFURLCopyFragment(anURL, charactersToLeaveEscaped)
}

// Returns the host name of a given URL.
//
// Added in macOS .
// Returns the host name of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyHostName(_:)
func CFURLCopyHostName(anURL URLRef) StringRef {
	return _CFURLCopyHostName(anURL)
}

// Returns the last path component of a given URL.
//
// Added in macOS .
// Returns the last path component of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyLastPathComponent(_:)
func CFURLCopyLastPathComponent(url URLRef) StringRef {
	return _CFURLCopyLastPathComponent(url)
}

// Returns the net location portion of a given URL.
//
// Added in macOS .
// Returns the net location portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyNetLocation(_:)
func CFURLCopyNetLocation(anURL URLRef) StringRef {
	return _CFURLCopyNetLocation(anURL)
}

// Returns the parameter string from a given URL.
//
// Deprecated: This function was deprecated in macOS 10.15.
//
// Added in macOS 10.2.
// Returns the parameter string from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyParameterString(_:_:)
func CFURLCopyParameterString(anURL URLRef, charactersToLeaveEscaped StringRef) StringRef {
	return _CFURLCopyParameterString(anURL, charactersToLeaveEscaped)
}

// Returns the password of a given URL.
//
// Added in macOS .
// Returns the password of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPassword(_:)
func CFURLCopyPassword(anURL URLRef) StringRef {
	return _CFURLCopyPassword(anURL)
}

// Returns the path portion of a given URL.
//
// Added in macOS .
// Returns the path portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPath(_:)
func CFURLCopyPath(anURL URLRef) StringRef {
	return _CFURLCopyPath(anURL)
}

// Returns the path extension of a given URL.
//
// Added in macOS .
// Returns the path extension of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPathExtension(_:)
func CFURLCopyPathExtension(url URLRef) StringRef {
	return _CFURLCopyPathExtension(url)
}

// Returns the query string of a given URL.
//
// Added in macOS .
// Returns the query string of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyQueryString(_:_:)
func CFURLCopyQueryString(anURL URLRef, charactersToLeaveEscaped StringRef) StringRef {
	return _CFURLCopyQueryString(anURL, charactersToLeaveEscaped)
}

// Returns the resource values for the properties identified by specified array of keys.
//
// Added in macOS 10.6.
// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourcePropertiesForKeys(_:_:_:)
func CFURLCopyResourcePropertiesForKeys(url URLRef, keys ArrayRef, error_ unsafe.Pointer) DictionaryRef {
	return _CFURLCopyResourcePropertiesForKeys(url, keys, error_)
}

// Returns the value of a given resource property of a given URL.
//
// Added in macOS 10.6.
// Returns the value of a given resource property of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourcePropertyForKey(_:_:_:_:)
func CFURLCopyResourcePropertyForKey(url URLRef, key StringRef, propertyValueTypeRefPtr unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyResourcePropertyForKey(url, key, propertyValueTypeRefPtr, error_)
}

// Returns any additional resource specifiers after the path.
//
// Added in macOS .
// Returns any additional resource specifiers after the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourceSpecifier(_:)
func CFURLCopyResourceSpecifier(anURL URLRef) StringRef {
	return _CFURLCopyResourceSpecifier(anURL)
}

// Returns the scheme portion of a given URL.
//
// Added in macOS .
// Returns the scheme portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyScheme(_:)
func CFURLCopyScheme(anURL URLRef) StringRef {
	return _CFURLCopyScheme(anURL)
}

// Returns the path portion of a given URL.
//
// Added in macOS .
// Returns the path portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyStrictPath(_:_:)
func CFURLCopyStrictPath(anURL URLRef, isAbsolute unsafe.Pointer) StringRef {
	return _CFURLCopyStrictPath(anURL, isAbsolute)
}

// Returns the user name from a given URL.
//
// Added in macOS .
// Returns the user name from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyUserName(_:)
func CFURLCopyUserName(anURL URLRef) StringRef {
	return _CFURLCopyUserName(anURL)
}

// Creates a new object by resolving the relative portion of a URL, specified as bytes, against its given base URL.
//
// Added in macOS .
// Creates a new object by resolving the relative portion of a URL, specified as bytes, against its given base URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateAbsoluteURLWithBytes(_:_:_:_:_:_:)
func CFURLCreateAbsoluteURLWithBytes(alloc AllocatorRef, relativeURLBytes unsafe.Pointer, length Index, encoding StringEncoding, baseURL URLRef, useCompatibilityMode unsafe.Pointer) URLRef {
	return _CFURLCreateAbsoluteURLWithBytes(alloc, relativeURLBytes, length, encoding, baseURL, useCompatibilityMode)
}

// Returns bookmark data for a URL, created with specified options and resource values.
//
// Added in macOS 10.6.
// Returns bookmark data for a URL, created with specified options and resource values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkData(_:_:_:_:_:_:)
func CFURLCreateBookmarkData(allocator AllocatorRef, url URLRef, options URLBookmarkCreationOptions, resourcePropertiesToInclude ArrayRef, relativeToURL URLRef, error_ unsafe.Pointer) DataRef {
	return _CFURLCreateBookmarkData(allocator, url, options, resourcePropertiesToInclude, relativeToURL, error_)
}

// Initializes and returns bookmark data derived from an alias record.
//
// Deprecated: This function was deprecated in macOS 11.0.
//
// Added in macOS 10.6.
// Initializes and returns bookmark data derived from an alias record.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkDataFromAliasRecord(_:_:)
func CFURLCreateBookmarkDataFromAliasRecord(allocatorRef AllocatorRef, aliasRecordDataRef DataRef) DataRef {
	return _CFURLCreateBookmarkDataFromAliasRecord(allocatorRef, aliasRecordDataRef)
}

// Initializes and returns bookmark data derived from a file pointed to by a specified URL.
//
// Added in macOS 10.6.
// Initializes and returns bookmark data derived from a file pointed to by a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkDataFromFile(_:_:_:)
func CFURLCreateBookmarkDataFromFile(allocator AllocatorRef, fileURL URLRef, errorRef unsafe.Pointer) DataRef {
	return _CFURLCreateBookmarkDataFromFile(allocator, fileURL, errorRef)
}

// Returns a new URL made by resolving bookmark data.
//
// Added in macOS 10.6.
// Returns a new URL made by resolving bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateByResolvingBookmarkData(_:_:_:_:_:_:_:)
func CFURLCreateByResolvingBookmarkData(allocator AllocatorRef, bookmark DataRef, options URLBookmarkResolutionOptions, relativeToURL URLRef, resourcePropertiesToInclude ArrayRef, isStale unsafe.Pointer, error_ unsafe.Pointer) URLRef {
	return _CFURLCreateByResolvingBookmarkData(allocator, bookmark, options, relativeToURL, resourcePropertiesToInclude, isStale, error_)
}

// Creates a copy of a given URL and appends a path component.
//
// Added in macOS .
// Creates a copy of a given URL and appends a path component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyAppendingPathComponent(_:_:_:_:)
func CFURLCreateCopyAppendingPathComponent(allocator AllocatorRef, url URLRef, pathComponent StringRef, isDirectory unsafe.Pointer) URLRef {
	return _CFURLCreateCopyAppendingPathComponent(allocator, url, pathComponent, isDirectory)
}

// Creates a copy of a given URL and appends a path extension.
//
// Added in macOS .
// Creates a copy of a given URL and appends a path extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyAppendingPathExtension(_:_:_:)
func CFURLCreateCopyAppendingPathExtension(allocator AllocatorRef, url URLRef, extension StringRef) URLRef {
	return _CFURLCreateCopyAppendingPathExtension(allocator, url, extension)
}

// Creates a copy of a given URL with the last path component deleted.
//
// Added in macOS .
// Creates a copy of a given URL with the last path component deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyDeletingLastPathComponent(_:_:)
func CFURLCreateCopyDeletingLastPathComponent(allocator AllocatorRef, url URLRef) URLRef {
	return _CFURLCreateCopyDeletingLastPathComponent(allocator, url)
}

// Creates a copy of a given URL with its last path extension removed.
//
// Added in macOS .
// Creates a copy of a given URL with its last path extension removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyDeletingPathExtension(_:_:)
func CFURLCreateCopyDeletingPathExtension(allocator AllocatorRef, url URLRef) URLRef {
	return _CFURLCreateCopyDeletingPathExtension(allocator, url)
}

// Creates a object containing the content of a given URL.
//
// Added in macOS .
// Creates a object containing the content of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateData(_:_:_:_:)
func CFURLCreateData(allocator AllocatorRef, url URLRef, encoding StringEncoding, escapeWhitespace unsafe.Pointer) DataRef {
	return _CFURLCreateData(allocator, url, encoding, escapeWhitespace)
}

// Loads the data and properties referred to by a given URL.

// Loads the data and properties referred to by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateDataAndPropertiesFromResource(_:_:_:_:_:_:)
func CFURLCreateDataAndPropertiesFromResource(alloc AllocatorRef, url URLRef, resourceData unsafe.Pointer, properties unsafe.Pointer, desiredProperties ArrayRef, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateDataAndPropertiesFromResource(alloc, url, resourceData, properties, desiredProperties, errorCode)
}

// Returns a new file path URL that refers to the same resource as a specified URL.
//
// Added in macOS 10.6.
// Returns a new file path URL that refers to the same resource as a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFilePathURL(_:_:_:)
func CFURLCreateFilePathURL(allocator AllocatorRef, url URLRef, error_ unsafe.Pointer) URLRef {
	return _CFURLCreateFilePathURL(allocator, url, error_)
}

// Returns a new file reference URL that points to the same resource as a specified URL.
//
// Added in macOS 10.6.
// Returns a new file reference URL that points to the same resource as a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFileReferenceURL(_:_:_:)
func CFURLCreateFileReferenceURL(allocator AllocatorRef, url URLRef, error_ unsafe.Pointer) URLRef {
	return _CFURLCreateFileReferenceURL(allocator, url, error_)
}

// Creates a URL from a given directory or file.

// Creates a URL from a given directory or file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFSRef(_:_:)
func CFURLCreateFromFSRef(allocator AllocatorRef, fsRef unsafe.Pointer) URLRef {
	return _CFURLCreateFromFSRef(allocator, fsRef)
}

// Creates a new object for a file system entity using the native representation.
//
// Added in macOS .
// Creates a new object for a file system entity using the native representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFileSystemRepresentation(_:_:_:_:)
func CFURLCreateFromFileSystemRepresentation(allocator AllocatorRef, buffer unsafe.Pointer, bufLen Index, isDirectory unsafe.Pointer) URLRef {
	return _CFURLCreateFromFileSystemRepresentation(allocator, buffer, bufLen, isDirectory)
}

// Creates a object from a native character string path relative to a base URL.
//
// Added in macOS .
// Creates a object from a native character string path relative to a base URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFileSystemRepresentationRelativeToBase(_:_:_:_:_:)
func CFURLCreateFromFileSystemRepresentationRelativeToBase(allocator AllocatorRef, buffer unsafe.Pointer, bufLen Index, isDirectory unsafe.Pointer, baseURL URLRef) URLRef {
	return _CFURLCreateFromFileSystemRepresentationRelativeToBase(allocator, buffer, bufLen, isDirectory, baseURL)
}

// Returns a given property specified by a given URL and property string.

// Returns a given property specified by a given URL and property string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreatePropertyFromResource(_:_:_:_:)
func CFURLCreatePropertyFromResource(alloc AllocatorRef, url URLRef, property StringRef, errorCode unsafe.Pointer) TypeRef {
	return _CFURLCreatePropertyFromResource(alloc, url, property, errorCode)
}

// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// Added in macOS 10.6.
// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateResourcePropertiesForKeysFromBookmarkData(_:_:_:)
func CFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator AllocatorRef, resourcePropertiesToReturn ArrayRef, bookmark DataRef) DictionaryRef {
	return _CFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator, resourcePropertiesToReturn, bookmark)
}

// Returns the value of a resource property from specified bookmark data.
//
// Added in macOS 10.6.
// Returns the value of a resource property from specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateResourcePropertyForKeyFromBookmarkData(_:_:_:)
func CFURLCreateResourcePropertyForKeyFromBookmarkData(allocator AllocatorRef, resourcePropertyKey StringRef, bookmark DataRef) TypeRef {
	return _CFURLCreateResourcePropertyForKeyFromBookmarkData(allocator, resourcePropertyKey, bookmark)
}

// Creates a copy of a string, replacing certain characters with the equivalent percent escape sequence based on the specified encoding.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.0.
// Creates a copy of a string, replacing certain characters with the equivalent percent escape sequence based on the specified encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateStringByAddingPercentEscapes(_:_:_:_:_:)
func CFURLCreateStringByAddingPercentEscapes(allocator AllocatorRef, originalString StringRef, charactersToLeaveUnescaped StringRef, legalURLCharactersToBeEscaped StringRef, encoding StringEncoding) StringRef {
	return _CFURLCreateStringByAddingPercentEscapes(allocator, originalString, charactersToLeaveUnescaped, legalURLCharactersToBeEscaped, encoding)
}

// Creates a new string by replacing any percent escape sequences with their character equivalent.
//
// Added in macOS .
// Creates a new string by replacing any percent escape sequences with their character equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateStringByReplacingPercentEscapes(_:_:_:)
func CFURLCreateStringByReplacingPercentEscapes(allocator AllocatorRef, originalString StringRef, charactersToLeaveEscaped StringRef) StringRef {
	return _CFURLCreateStringByReplacingPercentEscapes(allocator, originalString, charactersToLeaveEscaped)
}

// Creates a new string by replacing any percent escape sequences with their character equivalent.
//
// Deprecated: This function was deprecated in macOS 10.11.
//
// Added in macOS 10.0.
// Creates a new string by replacing any percent escape sequences with their character equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateStringByReplacingPercentEscapesUsingEncoding(_:_:_:_:)
func CFURLCreateStringByReplacingPercentEscapesUsingEncoding(allocator AllocatorRef, origString StringRef, charsToLeaveEscaped StringRef, encoding StringEncoding) StringRef {
	return _CFURLCreateStringByReplacingPercentEscapesUsingEncoding(allocator, origString, charsToLeaveEscaped, encoding)
}

// Creates a object using a given character bytes.
//
// Added in macOS .
// Creates a object using a given character bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithBytes(_:_:_:_:_:)
func CFURLCreateWithBytes(allocator AllocatorRef, URLBytes unsafe.Pointer, length Index, encoding StringEncoding, baseURL URLRef) URLRef {
	return _CFURLCreateWithBytes(allocator, URLBytes, length, encoding, baseURL)
}

// Creates a object using a local file system path string.
//
// Added in macOS .
// Creates a object using a local file system path string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithFileSystemPath(_:_:_:_:)
func CFURLCreateWithFileSystemPath(allocator AllocatorRef, filePath StringRef, pathStyle URLPathStyle, isDirectory unsafe.Pointer) URLRef {
	return _CFURLCreateWithFileSystemPath(allocator, filePath, pathStyle, isDirectory)
}

// Creates a object using a local file system path string relative to a base URL.
//
// Added in macOS .
// Creates a object using a local file system path string relative to a base URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithFileSystemPathRelativeToBase(_:_:_:_:_:)
func CFURLCreateWithFileSystemPathRelativeToBase(allocator AllocatorRef, filePath StringRef, pathStyle URLPathStyle, isDirectory unsafe.Pointer, baseURL URLRef) URLRef {
	return _CFURLCreateWithFileSystemPathRelativeToBase(allocator, filePath, pathStyle, isDirectory, baseURL)
}

// Creates a object using a given object.
//
// Added in macOS .
// Creates a object using a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithString(_:_:_:)
func CFURLCreateWithString(allocator AllocatorRef, URLString StringRef, baseURL URLRef) URLRef {
	return _CFURLCreateWithString(allocator, URLString, baseURL)
}

// Destroys a resource indicated by a given URL.

// Destroys a resource indicated by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLDestroyResource(_:_:)
func CFURLDestroyResource(url URLRef, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLDestroyResource(url, errorCode)
}

// Creates and returns a directory enumerator with provided enumerator behavior options and properties to be prefetched.
//
// Added in macOS 10.6.
// Creates and returns a directory enumerator with provided enumerator behavior options and properties to be prefetched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorCreateForDirectoryURL(_:_:_:_:)
func CFURLEnumeratorCreateForDirectoryURL(alloc AllocatorRef, directoryURL URLRef, option URLEnumeratorOptions, propertyKeys ArrayRef) URLEnumeratorRef {
	return _CFURLEnumeratorCreateForDirectoryURL(alloc, directoryURL, option, propertyKeys)
}

// Creates and returns a volume enumerator with provided enumerator behavior options and properties to be prefetched.
//
// Added in macOS 10.6.
// Creates and returns a volume enumerator with provided enumerator behavior options and properties to be prefetched.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorCreateForMountedVolumes(_:_:_:)
func CFURLEnumeratorCreateForMountedVolumes(alloc AllocatorRef, option URLEnumeratorOptions, propertyKeys ArrayRef) URLEnumeratorRef {
	return _CFURLEnumeratorCreateForMountedVolumes(alloc, option, propertyKeys)
}

// Returns the number of levels a recursive directory enumerator has descended.
//
// Added in macOS 10.6.
// Returns the number of levels a recursive directory enumerator has descended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetDescendentLevel(_:)
func CFURLEnumeratorGetDescendentLevel(enumerator URLEnumeratorRef) Index {
	return _CFURLEnumeratorGetDescendentLevel(enumerator)
}

// Advances an enumerator to the next URL.
//
// Added in macOS 10.6.
// Advances an enumerator to the next URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetNextURL(_:_:_:)
func CFURLEnumeratorGetNextURL(enumerator URLEnumeratorRef, url unsafe.Pointer, error_ unsafe.Pointer) URLEnumeratorResult {
	return _CFURLEnumeratorGetNextURL(enumerator, url, error_)
}

// This function is unimplemented, so it performs no operation.

// This function is unimplemented, so it performs no operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetSourceDidChange(_:)
func CFURLEnumeratorGetSourceDidChange(enumerator URLEnumeratorRef) unsafe.Pointer {
	return _CFURLEnumeratorGetSourceDidChange(enumerator)
}

// Returns the opaque type identifier for the CFURLEnumerator opaque type.
//
// Added in macOS 10.6.
// Returns the opaque type identifier for the CFURLEnumerator opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorGetTypeID()
func CFURLEnumeratorGetTypeID() TypeID {
	return _CFURLEnumeratorGetTypeID()
}

// Tells a recursive enumerator not to descend into the directory at the URL that was returned by the most recent call to the function.
//
// Added in macOS 10.6.
// Tells a recursive enumerator not to descend into the directory at the URL that was returned by the most recent call to the function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorSkipDescendents(_:)
func CFURLEnumeratorSkipDescendents(enumerator URLEnumeratorRef) {
	_CFURLEnumeratorSkipDescendents(enumerator)
}

// Returns the base URL of a given URL if it exists.
//
// Added in macOS .
// Returns the base URL of a given URL if it exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetBaseURL(_:)
func CFURLGetBaseURL(anURL URLRef) URLRef {
	return _CFURLGetBaseURL(anURL)
}

// Returns the range of the specified component in the bytes of a URL.
//
// Added in macOS .
// Returns the range of the specified component in the bytes of a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetByteRangeForComponent(_:_:_:)
func CFURLGetByteRangeForComponent(url URLRef, component URLComponentType, rangeIncludingSeparators unsafe.Pointer) Range {
	return _CFURLGetByteRangeForComponent(url, component, rangeIncludingSeparators)
}

// Returns by reference the byte representation of a URL object.
//
// Added in macOS .
// Returns by reference the byte representation of a URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetBytes(_:_:_:)
func CFURLGetBytes(url URLRef, buffer unsafe.Pointer, bufferLength Index) Index {
	return _CFURLGetBytes(url, buffer, bufferLength)
}

// Converts a given URL to a file or directory object.

// Converts a given URL to a file or directory object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetFSRef(_:_:)
func CFURLGetFSRef(url URLRef, fsRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetFSRef(url, fsRef)
}

// Fills a buffer with the file system’s native string representation of a given URL’s path.
//
// Added in macOS .
// Fills a buffer with the file system’s native string representation of a given URL’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetFileSystemRepresentation(_:_:_:_:)
func CFURLGetFileSystemRepresentation(url URLRef, resolveAgainstBase unsafe.Pointer, buffer unsafe.Pointer, maxBufLen Index) unsafe.Pointer {
	return _CFURLGetFileSystemRepresentation(url, resolveAgainstBase, buffer, maxBufLen)
}

// Returns the port number from a given URL.
//
// Added in macOS .
// Returns the port number from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetPortNumber(_:)
func CFURLGetPortNumber(anURL URLRef) unsafe.Pointer {
	return _CFURLGetPortNumber(anURL)
}

// Returns the URL as a object.
//
// Added in macOS .
// Returns the URL as a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetString(_:)
func CFURLGetString(anURL URLRef) StringRef {
	return _CFURLGetString(anURL)
}

// Returns the type identifier for the opaque type.
//
// Added in macOS .
// Returns the type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetTypeID()
func CFURLGetTypeID() TypeID {
	return _CFURLGetTypeID()
}

// Determines if a given URL’s path represents a directory.
//
// Added in macOS .
// Determines if a given URL’s path represents a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLHasDirectoryPath(_:)
func CFURLHasDirectoryPath(anURL URLRef) unsafe.Pointer {
	return _CFURLHasDirectoryPath(anURL)
}

// CFURLIsFileReferenceURL is a CoreFoundation function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLIsFileReferenceURL(_:)
func CFURLIsFileReferenceURL(url URLRef) unsafe.Pointer {
	return _CFURLIsFileReferenceURL(url)
}

// Returns whether the resource pointed to by a file URL can be reached.
//
// Added in macOS 10.6.
// Returns whether the resource pointed to by a file URL can be reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLResourceIsReachable(_:_:)
func CFURLResourceIsReachable(url URLRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLResourceIsReachable(url, error_)
}

// Sets the URL’s resource properties for a given set of keys to a given set of values.
//
// Added in macOS 10.6.
// Sets the URL’s resource properties for a given set of keys to a given set of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLSetResourcePropertiesForKeys(_:_:_:)
func CFURLSetResourcePropertiesForKeys(url URLRef, keyedPropertyValues DictionaryRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLSetResourcePropertiesForKeys(url, keyedPropertyValues, error_)
}

// Sets the URL’s resource property for a given key to a given value.
//
// Added in macOS 10.6.
// Sets the URL’s resource property for a given key to a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLSetResourcePropertyForKey(_:_:_:_:)
func CFURLSetResourcePropertyForKey(url URLRef, key StringRef, propertyValue TypeRef, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLSetResourcePropertyForKey(url, key, propertyValue, error_)
}

// Sets a temporary resource value on the URL.
//
// Added in macOS 10.6.
// Sets a temporary resource value on the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLSetTemporaryResourcePropertyForKey(_:_:_:)
func CFURLSetTemporaryResourcePropertyForKey(url URLRef, key StringRef, propertyValue TypeRef) {
	_CFURLSetTemporaryResourcePropertyForKey(url, key, propertyValue)
}

// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// Added in macOS 10.7.
// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLStartAccessingSecurityScopedResource(_:)
func CFURLStartAccessingSecurityScopedResource(url URLRef) unsafe.Pointer {
	return _CFURLStartAccessingSecurityScopedResource(url)
}

// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// Added in macOS 10.7.
// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLStopAccessingSecurityScopedResource(_:)
func CFURLStopAccessingSecurityScopedResource(url URLRef) {
	_CFURLStopAccessingSecurityScopedResource(url)
}

// Creates an alias file on disk at a specified location with specified bookmark data.
//
// Added in macOS 10.6.
// Creates an alias file on disk at a specified location with specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLWriteBookmarkDataToFile(_:_:_:_:)
func CFURLWriteBookmarkDataToFile(bookmarkRef DataRef, fileURL URLRef, options URLBookmarkFileCreationOptions, errorRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLWriteBookmarkDataToFile(bookmarkRef, fileURL, options, errorRef)
}

// Writes the given data and properties to a given URL.

// Writes the given data and properties to a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLWriteDataAndPropertiesToResource(_:_:_:_:)
func CFURLWriteDataAndPropertiesToResource(url URLRef, dataToWrite DataRef, propertiesToWrite DictionaryRef, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLWriteDataAndPropertiesToResource(url, dataToWrite, propertiesToWrite, errorCode)
}

// Creates a Universally Unique Identifier (UUID) object.
//
// Added in macOS .
// Creates a Universally Unique Identifier (UUID) object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreate(_:)
func CFUUIDCreate(alloc AllocatorRef) UUIDRef {
	return _CFUUIDCreate(alloc)
}

// Creates a CFUUID object for a specified string.
//
// Added in macOS .
// Creates a CFUUID object for a specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateFromString(_:_:)
func CFUUIDCreateFromString(alloc AllocatorRef, uuidStr StringRef) UUIDRef {
	return _CFUUIDCreateFromString(alloc, uuidStr)
}

// Creates a CFUUID object from raw UUID bytes.
//
// Added in macOS .
// Creates a CFUUID object from raw UUID bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateFromUUIDBytes(_:_:)
func CFUUIDCreateFromUUIDBytes(alloc AllocatorRef, bytes UUIDBytes) UUIDRef {
	return _CFUUIDCreateFromUUIDBytes(alloc, bytes)
}

// Returns the string representation of a specified CFUUID object.
//
// Added in macOS .
// Returns the string representation of a specified CFUUID object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateString(_:_:)
func CFUUIDCreateString(alloc AllocatorRef, uuid UUIDRef) StringRef {
	return _CFUUIDCreateString(alloc, uuid)
}

// Creates a CFUUID object from raw UUID bytes.
//
// Added in macOS .
// Creates a CFUUID object from raw UUID bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateWithBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CFUUIDCreateWithBytes(alloc AllocatorRef, byte0 unsafe.Pointer, byte1 unsafe.Pointer, byte2 unsafe.Pointer, byte3 unsafe.Pointer, byte4 unsafe.Pointer, byte5 unsafe.Pointer, byte6 unsafe.Pointer, byte7 unsafe.Pointer, byte8 unsafe.Pointer, byte9 unsafe.Pointer, byte10 unsafe.Pointer, byte11 unsafe.Pointer, byte12 unsafe.Pointer, byte13 unsafe.Pointer, byte14 unsafe.Pointer, byte15 unsafe.Pointer) UUIDRef {
	return _CFUUIDCreateWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15)
}

// Returns a CFUUID object from raw UUID bytes.
//
// Added in macOS .
// Returns a CFUUID object from raw UUID bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetConstantUUIDWithBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CFUUIDGetConstantUUIDWithBytes(alloc AllocatorRef, byte0 unsafe.Pointer, byte1 unsafe.Pointer, byte2 unsafe.Pointer, byte3 unsafe.Pointer, byte4 unsafe.Pointer, byte5 unsafe.Pointer, byte6 unsafe.Pointer, byte7 unsafe.Pointer, byte8 unsafe.Pointer, byte9 unsafe.Pointer, byte10 unsafe.Pointer, byte11 unsafe.Pointer, byte12 unsafe.Pointer, byte13 unsafe.Pointer, byte14 unsafe.Pointer, byte15 unsafe.Pointer) UUIDRef {
	return _CFUUIDGetConstantUUIDWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15)
}

// Returns the type identifier for all CFUUID objects.
//
// Added in macOS .
// Returns the type identifier for all CFUUID objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetTypeID()
func CFUUIDGetTypeID() TypeID {
	return _CFUUIDGetTypeID()
}

// Returns the value of a UUID object as raw bytes.
//
// Added in macOS .
// Returns the value of a UUID object as raw bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetUUIDBytes(_:)
func CFUUIDGetUUIDBytes(uuid UUIDRef) UUIDBytes {
	return _CFUUIDGetUUIDBytes(uuid)
}

// Cancels a user notification dialog.
//
// Added in macOS 10.0.
// Cancels a user notification dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCancel(_:)
func CFUserNotificationCancel(userNotification UserNotificationRef) unsafe.Pointer {
	return _CFUserNotificationCancel(userNotification)
}

// Creates a CFUserNotification object and displays its notification dialog on screen.
//
// Added in macOS 10.0.
// Creates a CFUserNotification object and displays its notification dialog on screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCreate(_:_:_:_:_:)
func CFUserNotificationCreate(allocator AllocatorRef, timeout TimeInterval, flags OptionFlags, error_ unsafe.Pointer, dictionary DictionaryRef) UserNotificationRef {
	return _CFUserNotificationCreate(allocator, timeout, flags, error_, dictionary)
}

// Creates a run loop source for a user notification.
//
// Added in macOS 10.0.
// Creates a run loop source for a user notification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCreateRunLoopSource(_:_:_:_:)
func CFUserNotificationCreateRunLoopSource(allocator AllocatorRef, userNotification UserNotificationRef, callout UserNotificationCallBack, order Index) RunLoopSourceRef {
	return _CFUserNotificationCreateRunLoopSource(allocator, userNotification, callout, order)
}

// Displays a user notification dialog and waits for a user response.
//
// Added in macOS 10.0.
// Displays a user notification dialog and waits for a user response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationDisplayAlert(_:_:_:_:_:_:_:_:_:_:_:)
func CFUserNotificationDisplayAlert(timeout TimeInterval, flags OptionFlags, iconURL URLRef, soundURL URLRef, localizationURL URLRef, alertHeader StringRef, alertMessage StringRef, defaultButtonTitle StringRef, alternateButtonTitle StringRef, otherButtonTitle StringRef, responseFlags unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationDisplayAlert(timeout, flags, iconURL, soundURL, localizationURL, alertHeader, alertMessage, defaultButtonTitle, alternateButtonTitle, otherButtonTitle, responseFlags)
}

// Displays a user notification dialog that does not need a user response.
//
// Added in macOS 10.0.
// Displays a user notification dialog that does not need a user response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationDisplayNotice(_:_:_:_:_:_:_:_:)
func CFUserNotificationDisplayNotice(timeout TimeInterval, flags OptionFlags, iconURL URLRef, soundURL URLRef, localizationURL URLRef, alertHeader StringRef, alertMessage StringRef, defaultButtonTitle StringRef) unsafe.Pointer {
	return _CFUserNotificationDisplayNotice(timeout, flags, iconURL, soundURL, localizationURL, alertHeader, alertMessage, defaultButtonTitle)
}

// Returns the dictionary containing all the text field values from a dismissed notification dialog.
//
// Added in macOS 10.0.
// Returns the dictionary containing all the text field values from a dismissed notification dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetResponseDictionary(_:)
func CFUserNotificationGetResponseDictionary(userNotification UserNotificationRef) DictionaryRef {
	return _CFUserNotificationGetResponseDictionary(userNotification)
}

// Extracts the values of the text fields from a dismissed notification dialog.
//
// Added in macOS 10.0.
// Extracts the values of the text fields from a dismissed notification dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetResponseValue(_:_:_:)
func CFUserNotificationGetResponseValue(userNotification UserNotificationRef, key StringRef, idx Index) StringRef {
	return _CFUserNotificationGetResponseValue(userNotification, key, idx)
}

// Returns the type identifier for the opaque type.
//
// Added in macOS 10.0.
// Returns the type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationGetTypeID()
func CFUserNotificationGetTypeID() TypeID {
	return _CFUserNotificationGetTypeID()
}

// Waits for the user to respond to a notification or for the notification to time out.
//
// Added in macOS 10.0.
// Waits for the user to respond to a notification or for the notification to time out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationReceiveResponse(_:_:_:)
func CFUserNotificationReceiveResponse(userNotification UserNotificationRef, timeout TimeInterval, responseFlags unsafe.Pointer) unsafe.Pointer {
	return _CFUserNotificationReceiveResponse(userNotification, timeout, responseFlags)
}

// Updates a displayed user notification dialog with new user interface information.
//
// Added in macOS 10.0.
// Updates a displayed user notification dialog with new user interface information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationUpdate(_:_:_:_:)
func CFUserNotificationUpdate(userNotification UserNotificationRef, timeout TimeInterval, flags OptionFlags, dictionary DictionaryRef) unsafe.Pointer {
	return _CFUserNotificationUpdate(userNotification, timeout, flags, dictionary)
}

// Returns whether a writable stream can accept new data without blocking.
//
// Added in macOS .
// Returns whether a writable stream can accept new data without blocking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCanAcceptBytes(_:)
func CFWriteStreamCanAcceptBytes(stream WriteStreamRef) unsafe.Pointer {
	return _CFWriteStreamCanAcceptBytes(stream)
}

// Closes a writable stream.
//
// Added in macOS .
// Closes a writable stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamClose(_:)
func CFWriteStreamClose(stream WriteStreamRef) {
	_CFWriteStreamClose(stream)
}

// CFWriteStreamCopyDispatchQueue is a CoreFoundation function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyDispatchQueue(_:)
func CFWriteStreamCopyDispatchQueue(stream WriteStreamRef) unsafe.Pointer {
	return _CFWriteStreamCopyDispatchQueue(stream)
}

// Returns the error associated with a stream.
//
// Added in macOS 10.5.
// Returns the error associated with a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyError(_:)
func CFWriteStreamCopyError(stream WriteStreamRef) ErrorRef {
	return _CFWriteStreamCopyError(stream)
}

// Returns the value of a property for a stream.
//
// Added in macOS .
// Returns the value of a property for a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyProperty(_:_:)
func CFWriteStreamCopyProperty(stream WriteStreamRef, propertyName StreamPropertyKey) TypeRef {
	return _CFWriteStreamCopyProperty(stream, propertyName)
}

// Creates a writable stream for a growable block of memory.
//
// Added in macOS .
// Creates a writable stream for a growable block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithAllocatedBuffers(_:_:)
func CFWriteStreamCreateWithAllocatedBuffers(alloc AllocatorRef, bufferAllocator AllocatorRef) WriteStreamRef {
	return _CFWriteStreamCreateWithAllocatedBuffers(alloc, bufferAllocator)
}

// Creates a writable stream for a fixed-size block of memory.
//
// Added in macOS .
// Creates a writable stream for a fixed-size block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithBuffer(_:_:_:)
func CFWriteStreamCreateWithBuffer(alloc AllocatorRef, buffer unsafe.Pointer, bufferCapacity Index) WriteStreamRef {
	return _CFWriteStreamCreateWithBuffer(alloc, buffer, bufferCapacity)
}

// Creates a writable stream for a file.
//
// Added in macOS .
// Creates a writable stream for a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithFile(_:_:)
func CFWriteStreamCreateWithFile(alloc AllocatorRef, fileURL URLRef) WriteStreamRef {
	return _CFWriteStreamCreateWithFile(alloc, fileURL)
}

// Returns the error status of a stream.
//
// Added in macOS .
// Returns the error status of a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetError(_:)
func CFWriteStreamGetError(stream WriteStreamRef) StreamError {
	return _CFWriteStreamGetError(stream)
}

// Returns the current state of a stream.
//
// Added in macOS .
// Returns the current state of a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetStatus(_:)
func CFWriteStreamGetStatus(stream WriteStreamRef) StreamStatus {
	return _CFWriteStreamGetStatus(stream)
}

// Returns the type identifier of all CFWriteStream objects.
//
// Added in macOS .
// Returns the type identifier of all CFWriteStream objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetTypeID()
func CFWriteStreamGetTypeID() TypeID {
	return _CFWriteStreamGetTypeID()
}

// Opens a stream for writing.
//
// Added in macOS .
// Opens a stream for writing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamOpen(_:)
func CFWriteStreamOpen(stream WriteStreamRef) unsafe.Pointer {
	return _CFWriteStreamOpen(stream)
}

// Schedules a stream into a run loop.
//
// Added in macOS .
// Schedules a stream into a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamScheduleWithRunLoop(_:_:_:)
func CFWriteStreamScheduleWithRunLoop(stream WriteStreamRef, runLoop RunLoopRef, runLoopMode RunLoopMode) {
	_CFWriteStreamScheduleWithRunLoop(stream, runLoop, runLoopMode)
}

// Assigns a client to a stream, which receives callbacks when certain events occur.
//
// Added in macOS .
// Assigns a client to a stream, which receives callbacks when certain events occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetClient(_:_:_:_:)
func CFWriteStreamSetClient(stream WriteStreamRef, streamEvents OptionFlags, clientCB WriteStreamClientCallBack, clientContext unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamSetClient(stream, streamEvents, clientCB, clientContext)
}

// CFWriteStreamSetDispatchQueue is a CoreFoundation function.
//
// Added in macOS 10.9.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetDispatchQueue(_:_:)
func CFWriteStreamSetDispatchQueue(stream WriteStreamRef, q unsafe.Pointer) {
	_CFWriteStreamSetDispatchQueue(stream, q)
}

// Sets the value of a property for a stream.
//
// Added in macOS .
// Sets the value of a property for a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetProperty(_:_:_:)
func CFWriteStreamSetProperty(stream WriteStreamRef, propertyName StreamPropertyKey, propertyValue TypeRef) unsafe.Pointer {
	return _CFWriteStreamSetProperty(stream, propertyName, propertyValue)
}

// Removes a stream from a particular run loop.
//
// Added in macOS .
// Removes a stream from a particular run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamUnscheduleFromRunLoop(_:_:_:)
func CFWriteStreamUnscheduleFromRunLoop(stream WriteStreamRef, runLoop RunLoopRef, runLoopMode RunLoopMode) {
	_CFWriteStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode)
}

// Writes data to a writable stream.
//
// Added in macOS .
// Writes data to a writable stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamWrite(_:_:_:)
func CFWriteStreamWrite(stream WriteStreamRef, buffer unsafe.Pointer, bufferLength Index) Index {
	return _CFWriteStreamWrite(stream, buffer, bufferLength)
}

// Given a CFString object containing XML source with unescaped entities, returns a string with specified XML entities escaped.
//
// Added in macOS .
// Given a CFString object containing XML source with unescaped entities, returns a string with specified XML entities escaped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLCreateStringByEscapingEntities(_:_:_:)
func CFXMLCreateStringByEscapingEntities(allocator AllocatorRef, string_ StringRef, entitiesDictionary DictionaryRef) StringRef {
	return _CFXMLCreateStringByEscapingEntities(allocator, string_, entitiesDictionary)
}

// Given a CFString object containing XML source with escaped entities, returns a string with specified XML entities unescaped.
//
// Added in macOS .
// Given a CFString object containing XML source with escaped entities, returns a string with specified XML entities unescaped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLCreateStringByUnescapingEntities(_:_:_:)
func CFXMLCreateStringByUnescapingEntities(allocator AllocatorRef, string_ StringRef, entitiesDictionary DictionaryRef) StringRef {
	return _CFXMLCreateStringByUnescapingEntities(allocator, string_, entitiesDictionary)
}

// Creates a new CFXMLNode.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Creates a new CFXMLNode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeCreate
func CFXMLNodeCreate(alloc AllocatorRef, xmlType XMLNodeTypeCode, dataString StringRef, additionalInfoPtr unsafe.Pointer, version Index) XMLNodeRef {
	return _CFXMLNodeCreate(alloc, xmlType, dataString, additionalInfoPtr, version)
}

// Creates a copy of a CFXMLNode object.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Creates a copy of a CFXMLNode object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeCreateCopy
func CFXMLNodeCreateCopy(alloc AllocatorRef, origNode XMLNodeRef) XMLNodeRef {
	return _CFXMLNodeCreateCopy(alloc, origNode)
}

// Returns the additional information pointer of a CFXMLNode object.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the additional information pointer of a CFXMLNode object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetInfoPtr
func CFXMLNodeGetInfoPtr(node XMLNodeRef) unsafe.Pointer {
	return _CFXMLNodeGetInfoPtr(node)
}

// Returns the data string from a CFXMLNode.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the data string from a CFXMLNode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetString
func CFXMLNodeGetString(node XMLNodeRef) StringRef {
	return _CFXMLNodeGetString(node)
}

// Returns the XML structure type code for a CFXMLNode object.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the XML structure type code for a CFXMLNode object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetTypeCode
func CFXMLNodeGetTypeCode(node XMLNodeRef) XMLNodeTypeCode {
	return _CFXMLNodeGetTypeCode(node)
}

// Returns the type identifier code for the CFXMLNode opaque type.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the type identifier code for the CFXMLNode opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetTypeID
func CFXMLNodeGetTypeID() TypeID {
	return _CFXMLNodeGetTypeID()
}

// Returns the version number for a CFXMLNode object.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the version number for a CFXMLNode object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNodeGetVersion
func CFXMLNodeGetVersion(node XMLNodeRef) Index {
	return _CFXMLNodeGetVersion(node)
}

// Causes a parser to abort with the given error code and description.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Causes a parser to abort with the given error code and description.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserAbort
func CFXMLParserAbort(parser XMLParserRef, errorCode XMLParserStatusCode, errorDescription StringRef) {
	_CFXMLParserAbort(parser, errorCode, errorDescription)
}

// Returns the user-readable description of the current error condition.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the user-readable description of the current error condition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCopyErrorDescription
func CFXMLParserCopyErrorDescription(parser XMLParserRef) StringRef {
	return _CFXMLParserCopyErrorDescription(parser)
}

// Creates a new XML parser for the specified XML data.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Creates a new XML parser for the specified XML data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCreate
func CFXMLParserCreate(allocator AllocatorRef, xmlData DataRef, dataSource URLRef, parseOptions OptionFlags, versionOfNodes Index, callBacks unsafe.Pointer, context unsafe.Pointer) XMLParserRef {
	return _CFXMLParserCreate(allocator, xmlData, dataSource, parseOptions, versionOfNodes, callBacks, context)
}

// Creates a new XML parser for the specified XML data at the specified URL.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Creates a new XML parser for the specified XML data at the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCreateWithDataFromURL
func CFXMLParserCreateWithDataFromURL(allocator AllocatorRef, dataSource URLRef, parseOptions OptionFlags, versionOfNodes Index, callBacks unsafe.Pointer, context unsafe.Pointer) XMLParserRef {
	return _CFXMLParserCreateWithDataFromURL(allocator, dataSource, parseOptions, versionOfNodes, callBacks, context)
}

// Returns the callbacks associated with an XML parser when it was created.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the callbacks associated with an XML parser when it was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetCallBacks
func CFXMLParserGetCallBacks(parser XMLParserRef, callBacks unsafe.Pointer) {
	_CFXMLParserGetCallBacks(parser, callBacks)
}

// Returns the context for an XML parser.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the context for an XML parser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetContext
func CFXMLParserGetContext(parser XMLParserRef, context unsafe.Pointer) {
	_CFXMLParserGetContext(parser, context)
}

// Returns the top-most object returned by the create XML structure callback.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the top-most object returned by the create XML structure callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetDocument
func CFXMLParserGetDocument(parser XMLParserRef) unsafe.Pointer {
	return _CFXMLParserGetDocument(parser)
}

// Returns the line number of the current parse location.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the line number of the current parse location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetLineNumber
func CFXMLParserGetLineNumber(parser XMLParserRef) Index {
	return _CFXMLParserGetLineNumber(parser)
}

// Returns the character index of the current parse location.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the character index of the current parse location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetLocation
func CFXMLParserGetLocation(parser XMLParserRef) Index {
	return _CFXMLParserGetLocation(parser)
}

// Returns the URL for the XML data being parsed.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the URL for the XML data being parsed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetSourceURL
func CFXMLParserGetSourceURL(parser XMLParserRef) URLRef {
	return _CFXMLParserGetSourceURL(parser)
}

// Returns a numeric code indicating the current status of the parser.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns a numeric code indicating the current status of the parser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetStatusCode
func CFXMLParserGetStatusCode(parser XMLParserRef) XMLParserStatusCode {
	return _CFXMLParserGetStatusCode(parser)
}

// Returns the type identifier for the CFXMLParser opaque type.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the type identifier for the CFXMLParser opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserGetTypeID
func CFXMLParserGetTypeID() TypeID {
	return _CFXMLParserGetTypeID()
}

// Begins a parse of the XML data that was associated with the parser when it was created.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Begins a parse of the XML data that was associated with the parser when it was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserParse
func CFXMLParserParse(parser XMLParserRef) unsafe.Pointer {
	return _CFXMLParserParse(parser)
}

// Parses the given XML data and returns the resulting CFXMLTree object.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Parses the given XML data and returns the resulting CFXMLTree object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateFromData
func CFXMLTreeCreateFromData(allocator AllocatorRef, xmlData DataRef, dataSource URLRef, parseOptions OptionFlags, versionOfNodes Index) XMLTreeRef {
	return _CFXMLTreeCreateFromData(allocator, xmlData, dataSource, parseOptions, versionOfNodes)
}

// Parses the given XML data and returns the resulting CFXMLTree object and any error information.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Parses the given XML data and returns the resulting CFXMLTree object and any error information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateFromDataWithError
func CFXMLTreeCreateFromDataWithError(allocator AllocatorRef, xmlData DataRef, dataSource URLRef, parseOptions OptionFlags, versionOfNodes Index, errorDict unsafe.Pointer) XMLTreeRef {
	return _CFXMLTreeCreateFromDataWithError(allocator, xmlData, dataSource, parseOptions, versionOfNodes, errorDict)
}

// Creates a new CFXMLTree object by loading the data to be parsed directly from a data source.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Creates a new CFXMLTree object by loading the data to be parsed directly from a data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateWithDataFromURL
func CFXMLTreeCreateWithDataFromURL(allocator AllocatorRef, dataSource URLRef, parseOptions OptionFlags, versionOfNodes Index) XMLTreeRef {
	return _CFXMLTreeCreateWithDataFromURL(allocator, dataSource, parseOptions, versionOfNodes)
}

// Creates a childless, parentless CFXMLTree object node for a CFXMLNode object.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Creates a childless, parentless CFXMLTree object node for a CFXMLNode object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateWithNode
func CFXMLTreeCreateWithNode(allocator AllocatorRef, node XMLNodeRef) XMLTreeRef {
	return _CFXMLTreeCreateWithNode(allocator, node)
}

// Generates an XML document from a CFXMLTree object which is ready to be written to permanent storage.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Generates an XML document from a CFXMLTree object which is ready to be written to permanent storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeCreateXMLData
func CFXMLTreeCreateXMLData(allocator AllocatorRef, xmlTree XMLTreeRef) DataRef {
	return _CFXMLTreeCreateXMLData(allocator, xmlTree)
}

// Returns the node of a CFXMLTree object.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Returns the node of a CFXMLTree object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTreeGetNode
func CFXMLTreeGetNode(xmlTree XMLTreeRef) XMLNodeRef {
	return _CFXMLTreeGetNode(xmlTree)
}

// inset is a CoreFoundation function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRect/inset(by:)
func inset(insets unsafe.Pointer, p1   UIEdgeInsets) unsafe.Pointer {
	return _inset(insets, p1)
}




