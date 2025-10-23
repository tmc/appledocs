// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// Foundation Functions (592 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSAllocateObject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSClassFromString func(unsafe.Pointer) unsafe.Pointer
	_NSCountFrames func() unsafe.Pointer
	_NSFrameAddress func(unsafe.Pointer) unsafe.Pointer
	_NSGetSizeAndAlignment func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSHomeDirectory func() unsafe.Pointer
	_NSHomeDirectoryForUser func(unsafe.Pointer) unsafe.Pointer
	_NSOpenStepRootDirectory func() unsafe.Pointer
	_NSSearchPathForDirectoriesInDomains func(unsafe.Pointer, unsafe.Pointer, bool) unsafe.Pointer
	_NSAllHashTableObjects func(unsafe.Pointer) unsafe.Pointer
	_NSAllMapTableKeys func(unsafe.Pointer) unsafe.Pointer
	_NSAllMapTableValues func(unsafe.Pointer) unsafe.Pointer
	_NSCompareHashTables func(unsafe.Pointer, unsafe.Pointer) bool
	_NSCompareMapTables func(unsafe.Pointer, unsafe.Pointer) bool
	_NSContainsRect func(coregraphics.CGRect, coregraphics.CGRect) bool
	_NSCopyHashTableWithZone func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCopyMapTableWithZone func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCountHashTable func(unsafe.Pointer) unsafe.Pointer
	_NSCountMapTable func(unsafe.Pointer) unsafe.Pointer
	_NSCreateHashTable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCreateHashTableWithZone func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCreateMapTable func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCreateMapTableWithZone func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalAdd func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalCompact func(unsafe.Pointer)
	_NSDecimalCompare func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalCopy func(unsafe.Pointer, unsafe.Pointer)
	_NSDecimalDivide func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalMultiply func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalMultiplyByPowerOf10 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalNormalize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalPower func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalRound func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_NSDecimalString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDecimalSubtract func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSDivideRect func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, float64, unsafe.Pointer)
	_NSEdgeInsetsEqual func(unsafe.Pointer, unsafe.Pointer) bool
	_NSEndHashTableEnumeration func(unsafe.Pointer)
	_NSEndMapTableEnumeration func(unsafe.Pointer)
	_NSEnumerateHashTable func(unsafe.Pointer) unsafe.Pointer
	_NSEnumerateMapTable func(unsafe.Pointer) unsafe.Pointer
	_NSEqualPoints func(coregraphics.CGPoint, coregraphics.CGPoint) bool
	_NSEqualRects func(coregraphics.CGRect, coregraphics.CGRect) bool
	_NSEqualSizes func(coregraphics.CGSize, coregraphics.CGSize) bool
	_NSFreeHashTable func(unsafe.Pointer)
	_NSFreeMapTable func(unsafe.Pointer)
	_NSHashGet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSHashInsert func(unsafe.Pointer, unsafe.Pointer)
	_NSHashInsertIfAbsent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSHashInsertKnownAbsent func(unsafe.Pointer, unsafe.Pointer)
	_NSHashRemove func(unsafe.Pointer, unsafe.Pointer)
	_NSInsetRect func(coregraphics.CGRect, float64, float64) coregraphics.CGRect
	_NSIntegralRect func(coregraphics.CGRect) coregraphics.CGRect
	_NSIntegralRectWithOptions func(coregraphics.CGRect, unsafe.Pointer) coregraphics.CGRect
	_NSIntersectionRange func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSIntersectionRect func(coregraphics.CGRect, coregraphics.CGRect) coregraphics.CGRect
	_NSIntersectsRect func(coregraphics.CGRect, coregraphics.CGRect) bool
	_NSIsEmptyRect func(coregraphics.CGRect) bool
	_NSMapGet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSMapInsert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_NSMapInsertIfAbsent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSMapInsertKnownAbsent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_NSMapMember func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_NSMapRemove func(unsafe.Pointer, unsafe.Pointer)
	_NSMouseInRect func(coregraphics.CGPoint, coregraphics.CGRect, bool) bool
	_NSNextHashEnumeratorItem func(unsafe.Pointer) unsafe.Pointer
	_NSNextMapEnumeratorPair func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_NSOffsetRect func(coregraphics.CGRect, float64, float64) coregraphics.CGRect
	_NSPointFromString func(unsafe.Pointer) coregraphics.CGPoint
	_NSPointInRect func(coregraphics.CGPoint, coregraphics.CGRect) bool
	_NSRangeFromString func(unsafe.Pointer) unsafe.Pointer
	_NSRectFromString func(unsafe.Pointer) coregraphics.CGRect
	_NSResetHashTable func(unsafe.Pointer)
	_NSResetMapTable func(unsafe.Pointer)
	_NSSizeFromString func(unsafe.Pointer) coregraphics.CGSize
	_NSStringFromHashTable func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromMapTable func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromPoint func(coregraphics.CGPoint) unsafe.Pointer
	_NSStringFromRange func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromRect func(coregraphics.CGRect) unsafe.Pointer
	_NSStringFromSize func(coregraphics.CGSize) unsafe.Pointer
	_NSUnionRange func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSUnionRect func(coregraphics.CGRect, coregraphics.CGRect) coregraphics.CGRect
	_AXAnimatedImagesEnabled func() bool
	_AXNameFromColor func(coregraphics.ColorRef) unsafe.Pointer
	_AXOpenSettingsFeature func(unsafe.Pointer)
	_AXPrefersActionSliderAlternative func() bool
	_AXPrefersHeadAnchorAlternative func() bool
	_AXPrefersHorizontalTextLayout func() bool
	_AXPrefersNonBlinkingTextInsertionIndicator func() bool
	_AXShowBordersEnabled func() bool
	_AXAssistiveAccessEnabled func() bool
	_NSAccessibilityPostNotification func(unsafe.Pointer, unsafe.Pointer)
	_NSAccessibilityPointInView func(unsafe.Pointer, coregraphics.CGPoint) coregraphics.CGPoint
	_NSAccessibilityFrameInView func(unsafe.Pointer, coregraphics.CGRect) coregraphics.CGRect
	_NSAvailableWindowDepths func() unsafe.Pointer
	_NSBestDepth func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, unsafe.Pointer) unsafe.Pointer
	_NSNumberOfColorComponents func(unsafe.Pointer) unsafe.Pointer
	_NSDrawButton func(coregraphics.CGRect, coregraphics.CGRect)
	_NSDrawDarkBezel func(coregraphics.CGRect, coregraphics.CGRect)
	_NSDrawGrayBezel func(coregraphics.CGRect, coregraphics.CGRect)
	_NSDrawGroove func(coregraphics.CGRect, coregraphics.CGRect)
	_NSDrawLightBezel func(coregraphics.CGRect, coregraphics.CGRect)
	_NSDrawNinePartImage func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, float64, bool)
	_NSDrawThreePartImage func(coregraphics.CGRect, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, bool, unsafe.Pointer, float64, bool)
	_NSDrawTiledRects func(coregraphics.CGRect, coregraphics.CGRect, unsafe.Pointer, []float64, unsafe.Pointer) coregraphics.CGRect
	_NSDrawWhiteBezel func(coregraphics.CGRect, coregraphics.CGRect)
	_NSDrawWindowBackground func(coregraphics.CGRect)
	_NSEraseRect func(coregraphics.CGRect)
	_NSSetFocusRingStyle func(unsafe.Pointer)
	_NSFrameRect func(coregraphics.CGRect)
	_NSFrameRectWithWidth func(coregraphics.CGRect, float64)
	_NSFrameRectWithWidthUsingOperation func(coregraphics.CGRect, float64, unsafe.Pointer)
	_NSHighlightRect func(coregraphics.CGRect)
	_NSCreateFileContentsPboardType func(unsafe.Pointer) unsafe.Pointer
	_NSCreateFilenamePboardType func(unsafe.Pointer) unsafe.Pointer
	_NSGetFileType func(unsafe.Pointer) unsafe.Pointer
	_NSRectClip func(coregraphics.CGRect)
	_NSRectClipList func(unsafe.Pointer, unsafe.Pointer)
	_NSRectFill func(coregraphics.CGRect)
	_NSRectFillList func(unsafe.Pointer, unsafe.Pointer)
	_NSRectFillListUsingOperation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_NSRectFillListWithColors func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_NSRectFillListWithColorsUsingOperation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_NSRectFillListWithGrays func(unsafe.Pointer, []float64, unsafe.Pointer)
	_NSRectFillUsingOperation func(coregraphics.CGRect, unsafe.Pointer)
	_NSBitsPerPixelFromDepth func(unsafe.Pointer) unsafe.Pointer
	_NSBitsPerSampleFromDepth func(unsafe.Pointer) unsafe.Pointer
	_NSColorSpaceFromDepth func(unsafe.Pointer) unsafe.Pointer
	_NSPlanarFromDepth func(unsafe.Pointer) bool
	_CFAllocatorAllocate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorDeallocate func(unsafe.Pointer, unsafe.Pointer)
	_CFAllocatorGetContext func(unsafe.Pointer, unsafe.Pointer)
	_CFAllocatorGetDefault func() unsafe.Pointer
	_CFAllocatorGetPreferredSizeForSize func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorGetTypeID func() unsafe.Pointer
	_CFAllocatorReallocate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAllocatorSetDefault func(unsafe.Pointer)
	_CFArrayApplyFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFArrayBSearchValues func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayContainsValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFArrayGetCountOfValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayGetFirstIndexOfValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayGetLastIndexOfValue func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayGetTypeID func() unsafe.Pointer
	_CFArrayGetValueAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFArrayGetValues func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFAttributedStringBeginEditing func(unsafe.Pointer)
	_CFAttributedStringCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreateMutable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreateWithSubstring func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringEndEditing func(unsafe.Pointer)
	_CFAttributedStringGetAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttributeAndLongestEffectiveRange func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttributesAndLongestEffectiveRange func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetLength func(unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetMutableString func(unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetString func(unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetTypeID func() unsafe.Pointer
	_CFAttributedStringRemoveAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFAttributedStringReplaceAttributedString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFAttributedStringReplaceString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFAttributedStringSetAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFAttributedStringSetAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFBagApplyFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFBagContainsValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFBagGetCountOfValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagGetTypeID func() unsafe.Pointer
	_CFBagGetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagGetValueIfPresent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBagGetValues func(unsafe.Pointer, unsafe.Pointer)
	_CFBinaryHeapAddValue func(unsafe.Pointer, unsafe.Pointer)
	_CFBinaryHeapApplyFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFBinaryHeapContainsValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapCreateCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetCountOfValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetMinimum func(unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetMinimumIfPresent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBinaryHeapGetTypeID func() unsafe.Pointer
	_CFBinaryHeapGetValues func(unsafe.Pointer, unsafe.Pointer)
	_CFBinaryHeapRemoveAllValues func(unsafe.Pointer)
	_CFBinaryHeapRemoveMinimumValue func(unsafe.Pointer)
	_CFBitVectorGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFBooleanGetTypeID func() unsafe.Pointer
	_CFBooleanGetValue func(unsafe.Pointer) unsafe.Pointer
	_CFBundleCloseBundleResourceMap func(unsafe.Pointer, unsafe.Pointer)
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
	_CFBundleGetDataPointersForNames func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFBundleGetDevelopmentRegion func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetFunctionPointerForName func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetFunctionPointersForNames func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFBundleGetIdentifier func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetInfoDictionary func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetLocalInfoDictionary func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetMainBundle func() unsafe.Pointer
	_CFBundleGetPackageInfo func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFBundleGetPackageInfoInDirectory func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetPlugIn func(unsafe.Pointer) unsafe.Pointer
	_CFBundleGetTypeID func() unsafe.Pointer
	_CFBundleGetValueForInfoDictionaryKey func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleGetVersionNumber func(unsafe.Pointer) unsafe.Pointer
	_CFBundleIsExecutableLoaded func(unsafe.Pointer) unsafe.Pointer
	_CFBundleLoadExecutable func(unsafe.Pointer) unsafe.Pointer
	_CFBundleLoadExecutableAndReturnError func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleOpenBundleResourceFiles func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleOpenBundleResourceMap func(unsafe.Pointer) unsafe.Pointer
	_CFBundlePreflightExecutable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFBundleUnloadExecutable func(unsafe.Pointer)
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
	_CFCalendarSetFirstWeekday func(unsafe.Pointer, unsafe.Pointer)
	_CFCalendarSetLocale func(unsafe.Pointer, unsafe.Pointer)
	_CFCalendarSetMinimumDaysInFirstWeek func(unsafe.Pointer, unsafe.Pointer)
	_CFCalendarSetTimeZone func(unsafe.Pointer, unsafe.Pointer)
	_CFCharacterSetAddCharactersInRange func(unsafe.Pointer, unsafe.Pointer)
	_CFCharacterSetAddCharactersInString func(unsafe.Pointer, unsafe.Pointer)
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
	_CFCharacterSetIntersect func(unsafe.Pointer, unsafe.Pointer)
	_CFCharacterSetInvert func(unsafe.Pointer)
	_CFCharacterSetIsCharacterMember func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetIsLongCharacterMember func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetIsSupersetOfSet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFCharacterSetRemoveCharactersInRange func(unsafe.Pointer, unsafe.Pointer)
	_CFCharacterSetRemoveCharactersInString func(unsafe.Pointer, unsafe.Pointer)
	_CFCharacterSetUnion func(unsafe.Pointer, unsafe.Pointer)
	_CFDataAppendBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFDataCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreateMutable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreateMutableCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreateWithBytesNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataDeleteBytes func(unsafe.Pointer, unsafe.Pointer)
	_CFDataFind func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataGetBytePtr func(unsafe.Pointer) unsafe.Pointer
	_CFDataGetBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFDataGetLength func(unsafe.Pointer) unsafe.Pointer
	_CFDataGetMutableBytePtr func(unsafe.Pointer) unsafe.Pointer
	_CFDataGetTypeID func() unsafe.Pointer
	_CFDataIncreaseLength func(unsafe.Pointer, unsafe.Pointer)
	_CFDataReplaceBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFDataSetLength func(unsafe.Pointer, unsafe.Pointer)
	_CFDateCompare func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateCreate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCopyProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreateDateFormatFromTemplate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreateDateFromString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreateStringWithAbsoluteTime func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterCreateStringWithDate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetAbsoluteTimeFromString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetDateStyle func(unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetFormat func(unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetLocale func(unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetTimeStyle func(unsafe.Pointer) unsafe.Pointer
	_CFDateFormatterGetTypeID func() unsafe.Pointer
	_CFDateFormatterSetFormat func(unsafe.Pointer, unsafe.Pointer)
	_CFDateFormatterSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFDateGetAbsoluteTime func(unsafe.Pointer) unsafe.Pointer
	_CFDateGetTimeIntervalSinceDate func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDateGetTypeID func() unsafe.Pointer
	_CFDictionaryApplyFunction func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFDictionaryContainsKey func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryContainsValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetCount func(unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetCountOfKey func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetCountOfValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetKeysAndValues func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFDictionaryGetTypeID func() unsafe.Pointer
	_CFDictionaryGetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDictionaryGetValueIfPresent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFEqual func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFErrorCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFFileDescriptorGetNativeDescriptor func(unsafe.Pointer) unsafe.Pointer
	_CFFileDescriptorIsValid func(unsafe.Pointer) unsafe.Pointer
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
	_CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode func(unsafe.Pointer, uint32) unsafe.Pointer
	_CFLocaleGetIdentifier func(unsafe.Pointer) unsafe.Pointer
	_CFLocaleGetLanguageCharacterDirection func(unsafe.Pointer) unsafe.Pointer
	_CFLocaleGetLanguageLineDirection func(unsafe.Pointer) unsafe.Pointer
	_CFLocaleGetSystem func() unsafe.Pointer
	_CFLocaleGetTypeID func() unsafe.Pointer
	_CFLocaleGetValue func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier func(unsafe.Pointer) uint32
	_CFNullGetTypeID func() unsafe.Pointer
	_CFNumberIsFloatType func(unsafe.Pointer) unsafe.Pointer
	_CFRunLoopStop func(unsafe.Pointer)
	_CFShowStr func(unsafe.Pointer)
	_CFStreamCreatePairWithSocketToHost func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
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
	_CFStringFind func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringFindCharacterFromSet func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringFindWithOptions func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringFindWithOptionsAndLocale func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetCString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetCStringPtr func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetCharacterAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetCharacters func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFStringGetCharactersPtr func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetDoubleValue func(unsafe.Pointer) float64
	_CFStringGetFastestEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetFileSystemRepresentation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetHyphenationLocationBeforeIndex func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetIntValue func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetLength func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetLineBounds func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFStringGetListOfAvailableEncodings func() unsafe.Pointer
	_CFStringGetMaximumSizeForEncoding func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetMaximumSizeOfFileSystemRepresentation func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetMostCompatibleMacStringEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetNameOfEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetParagraphBounds func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFStringGetPascalString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetPascalStringPtr func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetRangeOfComposedCharactersAtIndex func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringGetSmallestEncoding func(unsafe.Pointer) unsafe.Pointer
	_CFStringGetSystemEncoding func() unsafe.Pointer
	_CFStringGetTypeID func() unsafe.Pointer
	_CFStringHasPrefix func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringHasSuffix func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFStringIsEncodingAvailable func(unsafe.Pointer) unsafe.Pointer
	_CFStringIsHyphenationAvailableForLocale func(unsafe.Pointer) unsafe.Pointer
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
	_CFTimeZoneResetSystem func()
	_CFTimeZoneSetAbbreviationDictionary func(unsafe.Pointer)
	_CFTimeZoneSetDefault func(unsafe.Pointer)
	_CFURLCanBeDecomposed func(unsafe.Pointer) unsafe.Pointer
	_CFURLClearResourcePropertyCache func(unsafe.Pointer)
	_CFURLClearResourcePropertyCacheForKey func(unsafe.Pointer, unsafe.Pointer)
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
	_CFURLGetBaseURL func(unsafe.Pointer) unsafe.Pointer
	_CFURLGetByteRangeForComponent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLGetBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLGetFSRef func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLGetFileSystemRepresentation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLGetPortNumber func(unsafe.Pointer) unsafe.Pointer
	_CFURLGetString func(unsafe.Pointer) unsafe.Pointer
	_CFURLGetTypeID func() unsafe.Pointer
	_CFURLHasDirectoryPath func(unsafe.Pointer) unsafe.Pointer
	_CFURLResourceIsReachable func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLSetResourcePropertiesForKeys func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLSetResourcePropertyForKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLSetTemporaryResourcePropertyForKey func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFURLStartAccessingSecurityScopedResource func(unsafe.Pointer) unsafe.Pointer
	_CFURLStopAccessingSecurityScopedResource func(unsafe.Pointer)
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
	_CFWriteStreamCanAcceptBytes func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamClose func(unsafe.Pointer)
	_CFWriteStreamCopyError func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCopyProperty func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCreateWithAllocatedBuffers func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCreateWithBuffer func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamCreateWithFile func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamGetError func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamGetStatus func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamGetTypeID func() unsafe.Pointer
	_CFWriteStreamOpen func(unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamScheduleWithRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFWriteStreamSetClient func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamSetProperty func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFWriteStreamUnscheduleFromRunLoop func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFWriteStreamWrite func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_inset func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSCanURLAcceptURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSOpenFromURLSpec func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSOpenCFURLRef func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSSetDefaultRoleHandlerForContentType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSCopyApplicationURLsForURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSRegisterURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSCopyDefaultApplicationURLForContentType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSSetDefaultHandlerForURLScheme func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSCopyAllRoleHandlersForContentType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSCopyDefaultApplicationURLForURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSCopyApplicationURLsForBundleIdentifier func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_LSCopyDefaultRoleHandlerForContentType func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_floorf func(float32) float32
	_ceill func(unsafe.Pointer) unsafe.Pointer
	_ceilf func(float32) float32
	_ceil func(float64) float64
	_floorl func(unsafe.Pointer) unsafe.Pointer
	_floor func(float64) float64
	_os_proc_available_memory func() uintptr
	_CNCopyCurrentNetworkInfo func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromUIEdgeInsets func(unsafe.Pointer) unsafe.Pointer
	_NSTextAlignmentToCTTextAlignment func(unsafe.Pointer) unsafe.Pointer
	_UIAccessibilityButtonShapesEnabled func() bool
	_UIGuidedAccessConfigureAccessibilityFeatures func(unsafe.Pointer, bool)
	_UIAccessibilityConvertPathToScreenCoordinates func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_UIAccessibilityConvertFrameToScreenCoordinates func(coregraphics.CGRect, unsafe.Pointer) coregraphics.CGRect
	_UIAccessibilityFocusedElement func(unsafe.Pointer) unsafe.Pointer
	_UIAccessibilityHearingDevicePairedEar func() unsafe.Pointer
	_UIAccessibilityIsAssistiveTouchRunning func() bool
	_UIAccessibilityIsBoldTextEnabled func() bool
	_UIAccessibilityIsClosedCaptioningEnabled func() bool
	_UIAccessibilityDarkerSystemColorsEnabled func() bool
	_UIAccessibilityIsGrayscaleEnabled func() bool
	_UIAccessibilityIsGuidedAccessEnabled func() bool
	_UIAccessibilityIsInvertColorsEnabled func() bool
	_UIAccessibilityIsMonoAudioEnabled func() bool
	_UIAccessibilityIsOnOffSwitchLabelsEnabled func() bool
	_UIAccessibilityIsReduceMotionEnabled func() bool
	_UIAccessibilityIsReduceTransparencyEnabled func() bool
	_UIAccessibilityIsShakeToUndoEnabled func() bool
	_UIAccessibilityIsSpeakScreenEnabled func() bool
	_UIAccessibilityIsSpeakSelectionEnabled func() bool
	_UIAccessibilityIsSwitchControlRunning func() bool
	_UIAccessibilityIsVideoAutoplayEnabled func() bool
	_UIAccessibilityIsVoiceOverRunning func() bool
	_UIAccessibilityPostNotification func(unsafe.Pointer, unsafe.Pointer)
	_UIAccessibilityPrefersCrossFadeTransitions func() bool
	_UIAccessibilityRegisterGestureConflictWithZoom func()
	_UIAccessibilityRequestGuidedAccessSession func(bool)
	_UIAccessibilityShouldDifferentiateWithoutColor func() bool
	_UIAccessibilityZoomFocusChanged func(unsafe.Pointer, coregraphics.CGRect, unsafe.Pointer)
	_UIApplicationMain func(int, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
	_UIEdgeInsetsFromString func(unsafe.Pointer) unsafe.Pointer
	_UIImageSymbolWeightForFontWeight func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSAllocateObject, lib, "NSAllocateObject")
	tryRegister(&_NSClassFromString, lib, "NSClassFromString")
	tryRegister(&_NSCountFrames, lib, "NSCountFrames")
	tryRegister(&_NSFrameAddress, lib, "NSFrameAddress")
	tryRegister(&_NSGetSizeAndAlignment, lib, "NSGetSizeAndAlignment")
	tryRegister(&_NSHomeDirectory, lib, "NSHomeDirectory")
	tryRegister(&_NSHomeDirectoryForUser, lib, "NSHomeDirectoryForUser")
	tryRegister(&_NSOpenStepRootDirectory, lib, "NSOpenStepRootDirectory")
	tryRegister(&_NSSearchPathForDirectoriesInDomains, lib, "NSSearchPathForDirectoriesInDomains")
	tryRegister(&_NSAllHashTableObjects, lib, "NSAllHashTableObjects")
	tryRegister(&_NSAllMapTableKeys, lib, "NSAllMapTableKeys")
	tryRegister(&_NSAllMapTableValues, lib, "NSAllMapTableValues")
	tryRegister(&_NSCompareHashTables, lib, "NSCompareHashTables")
	tryRegister(&_NSCompareMapTables, lib, "NSCompareMapTables")
	tryRegister(&_NSContainsRect, lib, "NSContainsRect")
	tryRegister(&_NSCopyHashTableWithZone, lib, "NSCopyHashTableWithZone")
	tryRegister(&_NSCopyMapTableWithZone, lib, "NSCopyMapTableWithZone")
	tryRegister(&_NSCountHashTable, lib, "NSCountHashTable")
	tryRegister(&_NSCountMapTable, lib, "NSCountMapTable")
	tryRegister(&_NSCreateHashTable, lib, "NSCreateHashTable")
	tryRegister(&_NSCreateHashTableWithZone, lib, "NSCreateHashTableWithZone")
	tryRegister(&_NSCreateMapTable, lib, "NSCreateMapTable")
	tryRegister(&_NSCreateMapTableWithZone, lib, "NSCreateMapTableWithZone")
	tryRegister(&_NSDecimalAdd, lib, "NSDecimalAdd")
	tryRegister(&_NSDecimalCompact, lib, "NSDecimalCompact")
	tryRegister(&_NSDecimalCompare, lib, "NSDecimalCompare")
	tryRegister(&_NSDecimalCopy, lib, "NSDecimalCopy")
	tryRegister(&_NSDecimalDivide, lib, "NSDecimalDivide")
	tryRegister(&_NSDecimalMultiply, lib, "NSDecimalMultiply")
	tryRegister(&_NSDecimalMultiplyByPowerOf10, lib, "NSDecimalMultiplyByPowerOf10")
	tryRegister(&_NSDecimalNormalize, lib, "NSDecimalNormalize")
	tryRegister(&_NSDecimalPower, lib, "NSDecimalPower")
	tryRegister(&_NSDecimalRound, lib, "NSDecimalRound")
	tryRegister(&_NSDecimalString, lib, "NSDecimalString")
	tryRegister(&_NSDecimalSubtract, lib, "NSDecimalSubtract")
	tryRegister(&_NSDivideRect, lib, "NSDivideRect")
	tryRegister(&_NSEdgeInsetsEqual, lib, "NSEdgeInsetsEqual")
	tryRegister(&_NSEndHashTableEnumeration, lib, "NSEndHashTableEnumeration")
	tryRegister(&_NSEndMapTableEnumeration, lib, "NSEndMapTableEnumeration")
	tryRegister(&_NSEnumerateHashTable, lib, "NSEnumerateHashTable")
	tryRegister(&_NSEnumerateMapTable, lib, "NSEnumerateMapTable")
	tryRegister(&_NSEqualPoints, lib, "NSEqualPoints")
	tryRegister(&_NSEqualRects, lib, "NSEqualRects")
	tryRegister(&_NSEqualSizes, lib, "NSEqualSizes")
	tryRegister(&_NSFreeHashTable, lib, "NSFreeHashTable")
	tryRegister(&_NSFreeMapTable, lib, "NSFreeMapTable")
	tryRegister(&_NSHashGet, lib, "NSHashGet")
	tryRegister(&_NSHashInsert, lib, "NSHashInsert")
	tryRegister(&_NSHashInsertIfAbsent, lib, "NSHashInsertIfAbsent")
	tryRegister(&_NSHashInsertKnownAbsent, lib, "NSHashInsertKnownAbsent")
	tryRegister(&_NSHashRemove, lib, "NSHashRemove")
	tryRegister(&_NSInsetRect, lib, "NSInsetRect")
	tryRegister(&_NSIntegralRect, lib, "NSIntegralRect")
	tryRegister(&_NSIntegralRectWithOptions, lib, "NSIntegralRectWithOptions")
	tryRegister(&_NSIntersectionRange, lib, "NSIntersectionRange")
	tryRegister(&_NSIntersectionRect, lib, "NSIntersectionRect")
	tryRegister(&_NSIntersectsRect, lib, "NSIntersectsRect")
	tryRegister(&_NSIsEmptyRect, lib, "NSIsEmptyRect")
	tryRegister(&_NSMapGet, lib, "NSMapGet")
	tryRegister(&_NSMapInsert, lib, "NSMapInsert")
	tryRegister(&_NSMapInsertIfAbsent, lib, "NSMapInsertIfAbsent")
	tryRegister(&_NSMapInsertKnownAbsent, lib, "NSMapInsertKnownAbsent")
	tryRegister(&_NSMapMember, lib, "NSMapMember")
	tryRegister(&_NSMapRemove, lib, "NSMapRemove")
	tryRegister(&_NSMouseInRect, lib, "NSMouseInRect")
	tryRegister(&_NSNextHashEnumeratorItem, lib, "NSNextHashEnumeratorItem")
	tryRegister(&_NSNextMapEnumeratorPair, lib, "NSNextMapEnumeratorPair")
	tryRegister(&_NSOffsetRect, lib, "NSOffsetRect")
	tryRegister(&_NSPointFromString, lib, "NSPointFromString")
	tryRegister(&_NSPointInRect, lib, "NSPointInRect")
	tryRegister(&_NSRangeFromString, lib, "NSRangeFromString")
	tryRegister(&_NSRectFromString, lib, "NSRectFromString")
	tryRegister(&_NSResetHashTable, lib, "NSResetHashTable")
	tryRegister(&_NSResetMapTable, lib, "NSResetMapTable")
	tryRegister(&_NSSizeFromString, lib, "NSSizeFromString")
	tryRegister(&_NSStringFromHashTable, lib, "NSStringFromHashTable")
	tryRegister(&_NSStringFromMapTable, lib, "NSStringFromMapTable")
	tryRegister(&_NSStringFromPoint, lib, "NSStringFromPoint")
	tryRegister(&_NSStringFromRange, lib, "NSStringFromRange")
	tryRegister(&_NSStringFromRect, lib, "NSStringFromRect")
	tryRegister(&_NSStringFromSize, lib, "NSStringFromSize")
	tryRegister(&_NSUnionRange, lib, "NSUnionRange")
	tryRegister(&_NSUnionRect, lib, "NSUnionRect")
	tryRegister(&_AXAnimatedImagesEnabled, lib, "AXAnimatedImagesEnabled")
	tryRegister(&_AXNameFromColor, lib, "AXNameFromColor")
	tryRegister(&_AXOpenSettingsFeature, lib, "AXOpenSettingsFeature")
	tryRegister(&_AXPrefersActionSliderAlternative, lib, "AXPrefersActionSliderAlternative")
	tryRegister(&_AXPrefersHeadAnchorAlternative, lib, "AXPrefersHeadAnchorAlternative")
	tryRegister(&_AXPrefersHorizontalTextLayout, lib, "AXPrefersHorizontalTextLayout")
	tryRegister(&_AXPrefersNonBlinkingTextInsertionIndicator, lib, "AXPrefersNonBlinkingTextInsertionIndicator")
	tryRegister(&_AXShowBordersEnabled, lib, "AXShowBordersEnabled")
	tryRegister(&_AXAssistiveAccessEnabled, lib, "AXAssistiveAccessEnabled")
	tryRegister(&_NSAccessibilityPostNotification, lib, "NSAccessibilityPostNotification")
	tryRegister(&_NSAccessibilityPointInView, lib, "NSAccessibilityPointInView")
	tryRegister(&_NSAccessibilityFrameInView, lib, "NSAccessibilityFrameInView")
	tryRegister(&_NSAvailableWindowDepths, lib, "NSAvailableWindowDepths")
	tryRegister(&_NSBestDepth, lib, "NSBestDepth")
	tryRegister(&_NSNumberOfColorComponents, lib, "NSNumberOfColorComponents")
	tryRegister(&_NSDrawButton, lib, "NSDrawButton")
	tryRegister(&_NSDrawDarkBezel, lib, "NSDrawDarkBezel")
	tryRegister(&_NSDrawGrayBezel, lib, "NSDrawGrayBezel")
	tryRegister(&_NSDrawGroove, lib, "NSDrawGroove")
	tryRegister(&_NSDrawLightBezel, lib, "NSDrawLightBezel")
	tryRegister(&_NSDrawNinePartImage, lib, "NSDrawNinePartImage")
	tryRegister(&_NSDrawThreePartImage, lib, "NSDrawThreePartImage")
	tryRegister(&_NSDrawTiledRects, lib, "NSDrawTiledRects")
	tryRegister(&_NSDrawWhiteBezel, lib, "NSDrawWhiteBezel")
	tryRegister(&_NSDrawWindowBackground, lib, "NSDrawWindowBackground")
	tryRegister(&_NSEraseRect, lib, "NSEraseRect")
	tryRegister(&_NSSetFocusRingStyle, lib, "NSSetFocusRingStyle")
	tryRegister(&_NSFrameRect, lib, "NSFrameRect")
	tryRegister(&_NSFrameRectWithWidth, lib, "NSFrameRectWithWidth")
	tryRegister(&_NSFrameRectWithWidthUsingOperation, lib, "NSFrameRectWithWidthUsingOperation")
	tryRegister(&_NSHighlightRect, lib, "NSHighlightRect")
	tryRegister(&_NSCreateFileContentsPboardType, lib, "NSCreateFileContentsPboardType")
	tryRegister(&_NSCreateFilenamePboardType, lib, "NSCreateFilenamePboardType")
	tryRegister(&_NSGetFileType, lib, "NSGetFileType")
	tryRegister(&_NSRectClip, lib, "NSRectClip")
	tryRegister(&_NSRectClipList, lib, "NSRectClipList")
	tryRegister(&_NSRectFill, lib, "NSRectFill")
	tryRegister(&_NSRectFillList, lib, "NSRectFillList")
	tryRegister(&_NSRectFillListUsingOperation, lib, "NSRectFillListUsingOperation")
	tryRegister(&_NSRectFillListWithColors, lib, "NSRectFillListWithColors")
	tryRegister(&_NSRectFillListWithColorsUsingOperation, lib, "NSRectFillListWithColorsUsingOperation")
	tryRegister(&_NSRectFillListWithGrays, lib, "NSRectFillListWithGrays")
	tryRegister(&_NSRectFillUsingOperation, lib, "NSRectFillUsingOperation")
	tryRegister(&_NSBitsPerPixelFromDepth, lib, "NSBitsPerPixelFromDepth")
	tryRegister(&_NSBitsPerSampleFromDepth, lib, "NSBitsPerSampleFromDepth")
	tryRegister(&_NSColorSpaceFromDepth, lib, "NSColorSpaceFromDepth")
	tryRegister(&_NSPlanarFromDepth, lib, "NSPlanarFromDepth")
	tryRegister(&_CFAllocatorAllocate, lib, "CFAllocatorAllocate")
	tryRegister(&_CFAllocatorCreate, lib, "CFAllocatorCreate")
	tryRegister(&_CFAllocatorDeallocate, lib, "CFAllocatorDeallocate")
	tryRegister(&_CFAllocatorGetContext, lib, "CFAllocatorGetContext")
	tryRegister(&_CFAllocatorGetDefault, lib, "CFAllocatorGetDefault")
	tryRegister(&_CFAllocatorGetPreferredSizeForSize, lib, "CFAllocatorGetPreferredSizeForSize")
	tryRegister(&_CFAllocatorGetTypeID, lib, "CFAllocatorGetTypeID")
	tryRegister(&_CFAllocatorReallocate, lib, "CFAllocatorReallocate")
	tryRegister(&_CFAllocatorSetDefault, lib, "CFAllocatorSetDefault")
	tryRegister(&_CFArrayApplyFunction, lib, "CFArrayApplyFunction")
	tryRegister(&_CFArrayBSearchValues, lib, "CFArrayBSearchValues")
	tryRegister(&_CFArrayContainsValue, lib, "CFArrayContainsValue")
	tryRegister(&_CFArrayCreate, lib, "CFArrayCreate")
	tryRegister(&_CFArrayCreateCopy, lib, "CFArrayCreateCopy")
	tryRegister(&_CFArrayGetCount, lib, "CFArrayGetCount")
	tryRegister(&_CFArrayGetCountOfValue, lib, "CFArrayGetCountOfValue")
	tryRegister(&_CFArrayGetFirstIndexOfValue, lib, "CFArrayGetFirstIndexOfValue")
	tryRegister(&_CFArrayGetLastIndexOfValue, lib, "CFArrayGetLastIndexOfValue")
	tryRegister(&_CFArrayGetTypeID, lib, "CFArrayGetTypeID")
	tryRegister(&_CFArrayGetValueAtIndex, lib, "CFArrayGetValueAtIndex")
	tryRegister(&_CFArrayGetValues, lib, "CFArrayGetValues")
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
	tryRegister(&_CFAttributedStringGetLength, lib, "CFAttributedStringGetLength")
	tryRegister(&_CFAttributedStringGetMutableString, lib, "CFAttributedStringGetMutableString")
	tryRegister(&_CFAttributedStringGetString, lib, "CFAttributedStringGetString")
	tryRegister(&_CFAttributedStringGetTypeID, lib, "CFAttributedStringGetTypeID")
	tryRegister(&_CFAttributedStringRemoveAttribute, lib, "CFAttributedStringRemoveAttribute")
	tryRegister(&_CFAttributedStringReplaceAttributedString, lib, "CFAttributedStringReplaceAttributedString")
	tryRegister(&_CFAttributedStringReplaceString, lib, "CFAttributedStringReplaceString")
	tryRegister(&_CFAttributedStringSetAttribute, lib, "CFAttributedStringSetAttribute")
	tryRegister(&_CFAttributedStringSetAttributes, lib, "CFAttributedStringSetAttributes")
	tryRegister(&_CFBagApplyFunction, lib, "CFBagApplyFunction")
	tryRegister(&_CFBagContainsValue, lib, "CFBagContainsValue")
	tryRegister(&_CFBagCreate, lib, "CFBagCreate")
	tryRegister(&_CFBagCreateCopy, lib, "CFBagCreateCopy")
	tryRegister(&_CFBagGetCount, lib, "CFBagGetCount")
	tryRegister(&_CFBagGetCountOfValue, lib, "CFBagGetCountOfValue")
	tryRegister(&_CFBagGetTypeID, lib, "CFBagGetTypeID")
	tryRegister(&_CFBagGetValue, lib, "CFBagGetValue")
	tryRegister(&_CFBagGetValueIfPresent, lib, "CFBagGetValueIfPresent")
	tryRegister(&_CFBagGetValues, lib, "CFBagGetValues")
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
	tryRegister(&_CFBitVectorGetCount, lib, "CFBitVectorGetCount")
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
	tryRegister(&_CFDictionaryApplyFunction, lib, "CFDictionaryApplyFunction")
	tryRegister(&_CFDictionaryContainsKey, lib, "CFDictionaryContainsKey")
	tryRegister(&_CFDictionaryContainsValue, lib, "CFDictionaryContainsValue")
	tryRegister(&_CFDictionaryCreate, lib, "CFDictionaryCreate")
	tryRegister(&_CFDictionaryCreateCopy, lib, "CFDictionaryCreateCopy")
	tryRegister(&_CFDictionaryGetCount, lib, "CFDictionaryGetCount")
	tryRegister(&_CFDictionaryGetCountOfKey, lib, "CFDictionaryGetCountOfKey")
	tryRegister(&_CFDictionaryGetCountOfValue, lib, "CFDictionaryGetCountOfValue")
	tryRegister(&_CFDictionaryGetKeysAndValues, lib, "CFDictionaryGetKeysAndValues")
	tryRegister(&_CFDictionaryGetTypeID, lib, "CFDictionaryGetTypeID")
	tryRegister(&_CFDictionaryGetValue, lib, "CFDictionaryGetValue")
	tryRegister(&_CFDictionaryGetValueIfPresent, lib, "CFDictionaryGetValueIfPresent")
	tryRegister(&_CFEqual, lib, "CFEqual")
	tryRegister(&_CFErrorCreate, lib, "CFErrorCreate")
	tryRegister(&_CFFileDescriptorGetNativeDescriptor, lib, "CFFileDescriptorGetNativeDescriptor")
	tryRegister(&_CFFileDescriptorIsValid, lib, "CFFileDescriptorIsValid")
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
	tryRegister(&_CFNullGetTypeID, lib, "CFNullGetTypeID")
	tryRegister(&_CFNumberIsFloatType, lib, "CFNumberIsFloatType")
	tryRegister(&_CFRunLoopStop, lib, "CFRunLoopStop")
	tryRegister(&_CFShowStr, lib, "CFShowStr")
	tryRegister(&_CFStreamCreatePairWithSocketToHost, lib, "CFStreamCreatePairWithSocketToHost")
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
	tryRegister(&_CFStringFind, lib, "CFStringFind")
	tryRegister(&_CFStringFindCharacterFromSet, lib, "CFStringFindCharacterFromSet")
	tryRegister(&_CFStringFindWithOptions, lib, "CFStringFindWithOptions")
	tryRegister(&_CFStringFindWithOptionsAndLocale, lib, "CFStringFindWithOptionsAndLocale")
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
	tryRegister(&_CFStringIsEncodingAvailable, lib, "CFStringIsEncodingAvailable")
	tryRegister(&_CFStringIsHyphenationAvailableForLocale, lib, "CFStringIsHyphenationAvailableForLocale")
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
	tryRegister(&_CFURLGetBaseURL, lib, "CFURLGetBaseURL")
	tryRegister(&_CFURLGetByteRangeForComponent, lib, "CFURLGetByteRangeForComponent")
	tryRegister(&_CFURLGetBytes, lib, "CFURLGetBytes")
	tryRegister(&_CFURLGetFSRef, lib, "CFURLGetFSRef")
	tryRegister(&_CFURLGetFileSystemRepresentation, lib, "CFURLGetFileSystemRepresentation")
	tryRegister(&_CFURLGetPortNumber, lib, "CFURLGetPortNumber")
	tryRegister(&_CFURLGetString, lib, "CFURLGetString")
	tryRegister(&_CFURLGetTypeID, lib, "CFURLGetTypeID")
	tryRegister(&_CFURLHasDirectoryPath, lib, "CFURLHasDirectoryPath")
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
	tryRegister(&_CFWriteStreamCanAcceptBytes, lib, "CFWriteStreamCanAcceptBytes")
	tryRegister(&_CFWriteStreamClose, lib, "CFWriteStreamClose")
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
	tryRegister(&_CFWriteStreamSetProperty, lib, "CFWriteStreamSetProperty")
	tryRegister(&_CFWriteStreamUnscheduleFromRunLoop, lib, "CFWriteStreamUnscheduleFromRunLoop")
	tryRegister(&_CFWriteStreamWrite, lib, "CFWriteStreamWrite")
	tryRegister(&_inset, lib, "inset")
	tryRegister(&_LSCanURLAcceptURL, lib, "LSCanURLAcceptURL")
	tryRegister(&_LSOpenFromURLSpec, lib, "LSOpenFromURLSpec")
	tryRegister(&_LSOpenCFURLRef, lib, "LSOpenCFURLRef")
	tryRegister(&_LSSetDefaultRoleHandlerForContentType, lib, "LSSetDefaultRoleHandlerForContentType")
	tryRegister(&_LSCopyApplicationURLsForURL, lib, "LSCopyApplicationURLsForURL")
	tryRegister(&_LSRegisterURL, lib, "LSRegisterURL")
	tryRegister(&_LSCopyDefaultApplicationURLForContentType, lib, "LSCopyDefaultApplicationURLForContentType")
	tryRegister(&_LSSetDefaultHandlerForURLScheme, lib, "LSSetDefaultHandlerForURLScheme")
	tryRegister(&_LSCopyAllRoleHandlersForContentType, lib, "LSCopyAllRoleHandlersForContentType")
	tryRegister(&_LSCopyDefaultApplicationURLForURL, lib, "LSCopyDefaultApplicationURLForURL")
	tryRegister(&_LSCopyApplicationURLsForBundleIdentifier, lib, "LSCopyApplicationURLsForBundleIdentifier")
	tryRegister(&_LSCopyDefaultRoleHandlerForContentType, lib, "LSCopyDefaultRoleHandlerForContentType")
	tryRegister(&_floorf, lib, "floorf")
	tryRegister(&_ceill, lib, "ceill")
	tryRegister(&_ceilf, lib, "ceilf")
	tryRegister(&_ceil, lib, "ceil")
	tryRegister(&_floorl, lib, "floorl")
	tryRegister(&_floor, lib, "floor")
	tryRegister(&_os_proc_available_memory, lib, "os_proc_available_memory")
	tryRegister(&_CNCopyCurrentNetworkInfo, lib, "CNCopyCurrentNetworkInfo")
	tryRegister(&_NSStringFromUIEdgeInsets, lib, "NSStringFromUIEdgeInsets")
	tryRegister(&_NSTextAlignmentToCTTextAlignment, lib, "NSTextAlignmentToCTTextAlignment")
	tryRegister(&_UIAccessibilityButtonShapesEnabled, lib, "UIAccessibilityButtonShapesEnabled")
	tryRegister(&_UIGuidedAccessConfigureAccessibilityFeatures, lib, "UIGuidedAccessConfigureAccessibilityFeatures")
	tryRegister(&_UIAccessibilityConvertPathToScreenCoordinates, lib, "UIAccessibilityConvertPathToScreenCoordinates")
	tryRegister(&_UIAccessibilityConvertFrameToScreenCoordinates, lib, "UIAccessibilityConvertFrameToScreenCoordinates")
	tryRegister(&_UIAccessibilityFocusedElement, lib, "UIAccessibilityFocusedElement")
	tryRegister(&_UIAccessibilityHearingDevicePairedEar, lib, "UIAccessibilityHearingDevicePairedEar")
	tryRegister(&_UIAccessibilityIsAssistiveTouchRunning, lib, "UIAccessibilityIsAssistiveTouchRunning")
	tryRegister(&_UIAccessibilityIsBoldTextEnabled, lib, "UIAccessibilityIsBoldTextEnabled")
	tryRegister(&_UIAccessibilityIsClosedCaptioningEnabled, lib, "UIAccessibilityIsClosedCaptioningEnabled")
	tryRegister(&_UIAccessibilityDarkerSystemColorsEnabled, lib, "UIAccessibilityDarkerSystemColorsEnabled")
	tryRegister(&_UIAccessibilityIsGrayscaleEnabled, lib, "UIAccessibilityIsGrayscaleEnabled")
	tryRegister(&_UIAccessibilityIsGuidedAccessEnabled, lib, "UIAccessibilityIsGuidedAccessEnabled")
	tryRegister(&_UIAccessibilityIsInvertColorsEnabled, lib, "UIAccessibilityIsInvertColorsEnabled")
	tryRegister(&_UIAccessibilityIsMonoAudioEnabled, lib, "UIAccessibilityIsMonoAudioEnabled")
	tryRegister(&_UIAccessibilityIsOnOffSwitchLabelsEnabled, lib, "UIAccessibilityIsOnOffSwitchLabelsEnabled")
	tryRegister(&_UIAccessibilityIsReduceMotionEnabled, lib, "UIAccessibilityIsReduceMotionEnabled")
	tryRegister(&_UIAccessibilityIsReduceTransparencyEnabled, lib, "UIAccessibilityIsReduceTransparencyEnabled")
	tryRegister(&_UIAccessibilityIsShakeToUndoEnabled, lib, "UIAccessibilityIsShakeToUndoEnabled")
	tryRegister(&_UIAccessibilityIsSpeakScreenEnabled, lib, "UIAccessibilityIsSpeakScreenEnabled")
	tryRegister(&_UIAccessibilityIsSpeakSelectionEnabled, lib, "UIAccessibilityIsSpeakSelectionEnabled")
	tryRegister(&_UIAccessibilityIsSwitchControlRunning, lib, "UIAccessibilityIsSwitchControlRunning")
	tryRegister(&_UIAccessibilityIsVideoAutoplayEnabled, lib, "UIAccessibilityIsVideoAutoplayEnabled")
	tryRegister(&_UIAccessibilityIsVoiceOverRunning, lib, "UIAccessibilityIsVoiceOverRunning")
	tryRegister(&_UIAccessibilityPostNotification, lib, "UIAccessibilityPostNotification")
	tryRegister(&_UIAccessibilityPrefersCrossFadeTransitions, lib, "UIAccessibilityPrefersCrossFadeTransitions")
	tryRegister(&_UIAccessibilityRegisterGestureConflictWithZoom, lib, "UIAccessibilityRegisterGestureConflictWithZoom")
	tryRegister(&_UIAccessibilityRequestGuidedAccessSession, lib, "UIAccessibilityRequestGuidedAccessSession")
	tryRegister(&_UIAccessibilityShouldDifferentiateWithoutColor, lib, "UIAccessibilityShouldDifferentiateWithoutColor")
	tryRegister(&_UIAccessibilityZoomFocusChanged, lib, "UIAccessibilityZoomFocusChanged")
	tryRegister(&_UIApplicationMain, lib, "UIApplicationMain")
	tryRegister(&_UIEdgeInsetsFromString, lib, "UIEdgeInsetsFromString")
	tryRegister(&_UIImageSymbolWeightForFontWeight, lib, "UIImageSymbolWeightForFontWeight")
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



// Creates and returns a new instance of a given class.
//
// Added in macOS 10.0.
// Creates and returns a new instance of a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllocateObject
func NSAllocateObject(aClass unsafe.Pointer, extraBytes unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NSAllocateObject(aClass, extraBytes, zone)
}

// Obtains a class by name.
//
// Added in macOS 10.0.
// Obtains a class by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSClassFromString(_:)
func NSClassFromString(aClassName unsafe.Pointer) unsafe.Pointer {
	return _NSClassFromString(aClassName)
}

// Returns the number of call frames on the stack.
//
// Added in macOS 10.0.
// Returns the number of call frames on the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountFrames
func NSCountFrames() unsafe.Pointer {
	return _NSCountFrames()
}

// Returns the value of the frame pointer of the specified frame.
//
// Added in macOS 10.0.
// Returns the value of the frame pointer of the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFrameAddress
func NSFrameAddress(frame unsafe.Pointer) unsafe.Pointer {
	return _NSFrameAddress(frame)
}

// Obtains the actual size and the aligned size of an encoded type.
//
// Added in macOS 10.0.
// Obtains the actual size and the aligned size of an encoded type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGetSizeAndAlignment(_:_:_:)
func NSGetSizeAndAlignment(typePtr unsafe.Pointer, sizep unsafe.Pointer, alignp unsafe.Pointer) unsafe.Pointer {
	return _NSGetSizeAndAlignment(typePtr, sizep, alignp)
}

// Returns the path to either the user’s or application’s home directory, depending on the platform.
//
// Added in macOS 10.0.
// Returns the path to either the user’s or application’s home directory, depending on the platform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHomeDirectory()
func NSHomeDirectory() unsafe.Pointer {
	return _NSHomeDirectory()
}

// Returns the path to a given user’s home directory.
//
// Added in macOS 10.0.
// Returns the path to a given user’s home directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHomeDirectoryForUser(_:)
func NSHomeDirectoryForUser(userName unsafe.Pointer) unsafe.Pointer {
	return _NSHomeDirectoryForUser(userName)
}

// Returns the root directory of the user’s system.
//
// Added in macOS 10.0.
// Returns the root directory of the user’s system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOpenStepRootDirectory()
func NSOpenStepRootDirectory() unsafe.Pointer {
	return _NSOpenStepRootDirectory()
}

// Creates a list of directory search paths.
//
// Added in macOS 10.0.
// Creates a list of directory search paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSearchPathForDirectoriesInDomains(_:_:_:)
func NSSearchPathForDirectoriesInDomains(directory unsafe.Pointer, domainMask unsafe.Pointer, expandTilde bool) unsafe.Pointer {
	return _NSSearchPathForDirectoriesInDomains(directory, domainMask, expandTilde)
}

// Returns all of the elements in the specified hash table.
//
// Added in macOS 10.0.
// Returns all of the elements in the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllHashTableObjects(_:)
func NSAllHashTableObjects(table unsafe.Pointer) unsafe.Pointer {
	return _NSAllHashTableObjects(table)
}

// Returns all of the keys in the specified map table.
//
// Added in macOS 10.0.
// Returns all of the keys in the specified map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllMapTableKeys(_:)
func NSAllMapTableKeys(table unsafe.Pointer) unsafe.Pointer {
	return _NSAllMapTableKeys(table)
}

// Returns all of the values in the specified table.
//
// Added in macOS 10.0.
// Returns all of the values in the specified table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllMapTableValues(_:)
func NSAllMapTableValues(table unsafe.Pointer) unsafe.Pointer {
	return _NSAllMapTableValues(table)
}

// Returns a Boolean value that indicates whether the elements of two hash tables are equal.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether the elements of two hash tables are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompareHashTables(_:_:)
func NSCompareHashTables(table1 unsafe.Pointer, table2 unsafe.Pointer) bool {
	return _NSCompareHashTables(table1, table2)
}

// Compares the elements of two map tables for equality.
//
// Added in macOS 10.0.
// Compares the elements of two map tables for equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompareMapTables(_:_:)
func NSCompareMapTables(table1 unsafe.Pointer, table2 unsafe.Pointer) bool {
	return _NSCompareMapTables(table1, table2)
}

// Returns a Boolean value that indicates whether one rectangle completely encloses another.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether one rectangle completely encloses another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSContainsRect(_:_:)
func NSContainsRect(aRect coregraphics.CGRect, bRect coregraphics.CGRect) bool {
	return _NSContainsRect(aRect, bRect)
}

// Performs a shallow copy of the specified hash table.
//
// Added in macOS 10.0.
// Performs a shallow copy of the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCopyHashTableWithZone(_:_:)
func NSCopyHashTableWithZone(table unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NSCopyHashTableWithZone(table, zone)
}

// Performs a shallow copy of the specified map table.
//
// Added in macOS 10.0.
// Performs a shallow copy of the specified map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCopyMapTableWithZone(_:_:)
func NSCopyMapTableWithZone(table unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NSCopyMapTableWithZone(table, zone)
}

// Returns the number of elements in a hash table.
//
// Added in macOS 10.0.
// Returns the number of elements in a hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountHashTable(_:)
func NSCountHashTable(table unsafe.Pointer) unsafe.Pointer {
	return _NSCountHashTable(table)
}

// Returns the number of elements in a map table.
//
// Added in macOS 10.0.
// Returns the number of elements in a map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountMapTable(_:)
func NSCountMapTable(table unsafe.Pointer) unsafe.Pointer {
	return _NSCountMapTable(table)
}

// Creates and returns a new hash table.
//
// Added in macOS 10.0.
// Creates and returns a new hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateHashTable(_:_:)
func NSCreateHashTable(callBacks unsafe.Pointer, capacity unsafe.Pointer) unsafe.Pointer {
	return _NSCreateHashTable(callBacks, capacity)
}

// Creates a new hash table in a given zone.
//
// Added in macOS 10.0.
// Creates a new hash table in a given zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateHashTableWithZone(_:_:_:)
func NSCreateHashTableWithZone(callBacks unsafe.Pointer, capacity unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NSCreateHashTableWithZone(callBacks, capacity, zone)
}

// Creates a new map table in the default zone.
//
// Added in macOS 10.0.
// Creates a new map table in the default zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateMapTable(_:_:_:)
func NSCreateMapTable(keyCallBacks unsafe.Pointer, valueCallBacks unsafe.Pointer, capacity unsafe.Pointer) unsafe.Pointer {
	return _NSCreateMapTable(keyCallBacks, valueCallBacks, capacity)
}

// Creates a new map table in the specified zone.
//
// Added in macOS 10.0.
// Creates a new map table in the specified zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateMapTableWithZone(_:_:_:_:)
func NSCreateMapTableWithZone(keyCallBacks unsafe.Pointer, valueCallBacks unsafe.Pointer, capacity unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NSCreateMapTableWithZone(keyCallBacks, valueCallBacks, capacity, zone)
}

// Adds two decimal values.
//
// Added in macOS 10.0.
// Adds two decimal values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalAdd(_:_:_:_:)
func NSDecimalAdd(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalAdd(result, leftOperand, rightOperand, roundingMode)
}

// Compacts the decimal structure for efficiency.
//
// Added in macOS 10.0.
// Compacts the decimal structure for efficiency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCompact(_:)
func NSDecimalCompact(number unsafe.Pointer) {
	_NSDecimalCompact(number)
}

// Compares two decimal values.
//
// Added in macOS 10.0.
// Compares two decimal values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCompare(_:_:)
func NSDecimalCompare(leftOperand unsafe.Pointer, rightOperand unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalCompare(leftOperand, rightOperand)
}

// Copies the value of a decimal number.
//
// Added in macOS 10.0.
// Copies the value of a decimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCopy(_:_:)
func NSDecimalCopy(destination unsafe.Pointer, source unsafe.Pointer) {
	_NSDecimalCopy(destination, source)
}

// Divides one decimal value by another.
//
// Added in macOS 10.0.
// Divides one decimal value by another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalDivide(_:_:_:_:)
func NSDecimalDivide(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalDivide(result, leftOperand, rightOperand, roundingMode)
}

// Multiplies two decimal numbers together.
//
// Added in macOS 10.0.
// Multiplies two decimal numbers together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalMultiply(_:_:_:_:)
func NSDecimalMultiply(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalMultiply(result, leftOperand, rightOperand, roundingMode)
}

// Multiplies a decimal by the specified power of 10.
//
// Added in macOS 10.0.
// Multiplies a decimal by the specified power of 10.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalMultiplyByPowerOf10(_:_:_:_:)
func NSDecimalMultiplyByPowerOf10(result unsafe.Pointer, number unsafe.Pointer, power unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalMultiplyByPowerOf10(result, number, power, roundingMode)
}

// Normalizes the internal format of two decimal numbers to simplify later operations.
//
// Added in macOS 10.0.
// Normalizes the internal format of two decimal numbers to simplify later operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNormalize(_:_:_:)
func NSDecimalNormalize(number1 unsafe.Pointer, number2 unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalNormalize(number1, number2, roundingMode)
}

// Raises the decimal value to the specified power.
//
// Added in macOS 10.0.
// Raises the decimal value to the specified power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalPower(_:_:_:_:)
func NSDecimalPower(result unsafe.Pointer, number unsafe.Pointer, power unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalPower(result, number, power, roundingMode)
}

// Rounds off the decimal value.
//
// Added in macOS 10.0.
// Rounds off the decimal value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalRound(_:_:_:_:)
func NSDecimalRound(result unsafe.Pointer, number unsafe.Pointer, scale unsafe.Pointer, roundingMode unsafe.Pointer) {
	_NSDecimalRound(result, number, scale, roundingMode)
}

// Returns a string representation of the decimal value appropriate for the specified locale.
//
// Added in macOS 10.0.
// Returns a string representation of the decimal value appropriate for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalString(_:_:)
func NSDecimalString(dcm unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalString(dcm, locale)
}

// Subtracts one decimal value from another.
//
// Added in macOS 10.0.
// Subtracts one decimal value from another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalSubtract(_:_:_:_:)
func NSDecimalSubtract(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode unsafe.Pointer) unsafe.Pointer {
	return _NSDecimalSubtract(result, leftOperand, rightOperand, roundingMode)
}

// Divides a rectangle into two new rectangles.
//
// Added in macOS 10.0.
// Divides a rectangle into two new rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDivideRect(_:_:_:_:_:)
func NSDivideRect(inRect coregraphics.CGRect, slice unsafe.Pointer, rem unsafe.Pointer, amount float64, edge unsafe.Pointer) {
	_NSDivideRect(inRect, slice, rem, amount, edge)
}

// Returns a Boolean value that indicates whether two edge insets structures are equal.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates whether two edge insets structures are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEdgeInsetsEqual(_:_:)
func NSEdgeInsetsEqual(aInsets unsafe.Pointer, bInsets unsafe.Pointer) bool {
	return _NSEdgeInsetsEqual(aInsets, bInsets)
}

// Used when finished with an enumerator.
//
// Added in macOS 10.0.
// Used when finished with an enumerator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEndHashTableEnumeration(_:)
func NSEndHashTableEnumeration(enumerator unsafe.Pointer) {
	_NSEndHashTableEnumeration(enumerator)
}

// Used when finished with an enumerator.
//
// Added in macOS 10.0.
// Used when finished with an enumerator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEndMapTableEnumeration(_:)
func NSEndMapTableEnumeration(enumerator unsafe.Pointer) {
	_NSEndMapTableEnumeration(enumerator)
}

// Creates an enumerator for the specified hash table.
//
// Added in macOS 10.0.
// Creates an enumerator for the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerateHashTable(_:)
func NSEnumerateHashTable(table unsafe.Pointer) unsafe.Pointer {
	return _NSEnumerateHashTable(table)
}

// Creates an enumerator for the specified map table.
//
// Added in macOS 10.0.
// Creates an enumerator for the specified map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerateMapTable(_:)
func NSEnumerateMapTable(table unsafe.Pointer) unsafe.Pointer {
	return _NSEnumerateMapTable(table)
}

// Returns a Boolean value that indicates whether two points are equal.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether two points are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEqualPoints(_:_:)
func NSEqualPoints(aPoint coregraphics.CGPoint, bPoint coregraphics.CGPoint) bool {
	return _NSEqualPoints(aPoint, bPoint)
}

// Returns a Boolean value that indicates whether the two rectangles are equal.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether the two rectangles are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEqualRects(_:_:)
func NSEqualRects(aRect coregraphics.CGRect, bRect coregraphics.CGRect) bool {
	return _NSEqualRects(aRect, bRect)
}

// Returns a Boolean that indicates whether two size values are equal.
//
// Added in macOS 10.0.
// Returns a Boolean that indicates whether two size values are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEqualSizes(_:_:)
func NSEqualSizes(aSize coregraphics.CGSize, bSize coregraphics.CGSize) bool {
	return _NSEqualSizes(aSize, bSize)
}

// Deletes the specified hash table.
//
// Added in macOS 10.0.
// Deletes the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFreeHashTable(_:)
func NSFreeHashTable(table unsafe.Pointer) {
	_NSFreeHashTable(table)
}

// Deletes the specified map table.
//
// Added in macOS 10.0.
// Deletes the specified map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFreeMapTable(_:)
func NSFreeMapTable(table unsafe.Pointer) {
	_NSFreeMapTable(table)
}

// Returns an element of the hash table.
//
// Added in macOS 10.0.
// Returns an element of the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashGet(_:_:)
func NSHashGet(table unsafe.Pointer, pointer unsafe.Pointer) unsafe.Pointer {
	return _NSHashGet(table, pointer)
}

// Adds an element to the specified hash table.
//
// Added in macOS 10.0.
// Adds an element to the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashInsert(_:_:)
func NSHashInsert(table unsafe.Pointer, pointer unsafe.Pointer) {
	_NSHashInsert(table, pointer)
}

// Adds an element to the specified hash table only if the table does not already contain the element.
//
// Added in macOS 10.0.
// Adds an element to the specified hash table only if the table does not already contain the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashInsertIfAbsent(_:_:)
func NSHashInsertIfAbsent(table unsafe.Pointer, pointer unsafe.Pointer) unsafe.Pointer {
	return _NSHashInsertIfAbsent(table, pointer)
}

// Adds an element to the specified hash table.
//
// Added in macOS 10.0.
// Adds an element to the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashInsertKnownAbsent(_:_:)
func NSHashInsertKnownAbsent(table unsafe.Pointer, pointer unsafe.Pointer) {
	_NSHashInsertKnownAbsent(table, pointer)
}

// Removes an element from the specified hash table.
//
// Added in macOS 10.0.
// Removes an element from the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashRemove(_:_:)
func NSHashRemove(table unsafe.Pointer, pointer unsafe.Pointer) {
	_NSHashRemove(table, pointer)
}

// Insets a rectangle by a specified amount.
//
// Added in macOS 10.0.
// Insets a rectangle by a specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInsetRect(_:_:_:)
func NSInsetRect(aRect coregraphics.CGRect, dX float64, dY float64) coregraphics.CGRect {
	return _NSInsetRect(aRect, dX, dY)
}

// Adjusts the sides of a rectangle to integer values.
//
// Added in macOS 10.0.
// Adjusts the sides of a rectangle to integer values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntegralRect(_:)
func NSIntegralRect(aRect coregraphics.CGRect) coregraphics.CGRect {
	return _NSIntegralRect(aRect)
}

// Adjusts the sides of a rectangle to integral values using the specified options.
//
// Added in macOS 10.7.
// Adjusts the sides of a rectangle to integral values using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntegralRectWithOptions(_:_:)
func NSIntegralRectWithOptions(aRect coregraphics.CGRect, opts unsafe.Pointer) coregraphics.CGRect {
	return _NSIntegralRectWithOptions(aRect, opts)
}

// Returns the intersection of the specified ranges.
//
// Added in macOS 10.0.
// Returns the intersection of the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntersectionRange(_:_:)
func NSIntersectionRange(range1 unsafe.Pointer, range2 unsafe.Pointer) unsafe.Pointer {
	return _NSIntersectionRange(range1, range2)
}

// Calculates the intersection of two rectangles.
//
// Added in macOS 10.0.
// Calculates the intersection of two rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntersectionRect(_:_:)
func NSIntersectionRect(aRect coregraphics.CGRect, bRect coregraphics.CGRect) coregraphics.CGRect {
	return _NSIntersectionRect(aRect, bRect)
}

// Returns a Boolean value that indicates whether two rectangles intersect.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether two rectangles intersect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntersectsRect(_:_:)
func NSIntersectsRect(aRect coregraphics.CGRect, bRect coregraphics.CGRect) bool {
	return _NSIntersectsRect(aRect, bRect)
}

// Returns a Boolean value that indicates whether a given rectangle is empty.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether a given rectangle is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIsEmptyRect(_:)
func NSIsEmptyRect(aRect coregraphics.CGRect) bool {
	return _NSIsEmptyRect(aRect)
}

// Returns a map table value for the specified key.
//
// Added in macOS 10.0.
// Returns a map table value for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapGet(_:_:)
func NSMapGet(table unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _NSMapGet(table, key)
}

// Inserts a key-value pair into the specified table.
//
// Added in macOS 10.0.
// Inserts a key-value pair into the specified table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapInsert(_:_:_:)
func NSMapInsert(table unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_NSMapInsert(table, key, value)
}

// Inserts a key-value pair into the specified table.
//
// Added in macOS 10.0.
// Inserts a key-value pair into the specified table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapInsertIfAbsent(_:_:_:)
func NSMapInsertIfAbsent(table unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _NSMapInsertIfAbsent(table, key, value)
}

// Inserts a key-value pair into the specified table if the pair had not been previously added.
//
// Added in macOS 10.0.
// Inserts a key-value pair into the specified table if the pair had not been previously added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapInsertKnownAbsent(_:_:_:)
func NSMapInsertKnownAbsent(table unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_NSMapInsertKnownAbsent(table, key, value)
}

// Indicates whether a given table contains a given key.
//
// Added in macOS 10.0.
// Indicates whether a given table contains a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapMember(_:_:_:_:)
func NSMapMember(table unsafe.Pointer, key unsafe.Pointer, originalKey unsafe.Pointer, value unsafe.Pointer) bool {
	return _NSMapMember(table, key, originalKey, value)
}

// Removes a key and corresponding value from the specified table.
//
// Added in macOS 10.0.
// Removes a key and corresponding value from the specified table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapRemove(_:_:)
func NSMapRemove(table unsafe.Pointer, key unsafe.Pointer) {
	_NSMapRemove(table, key)
}

// Returns a Boolean value that indicates whether the point is in the specified rectangle.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether the point is in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMouseInRect(_:_:_:)
func NSMouseInRect(aPoint coregraphics.CGPoint, aRect coregraphics.CGRect, flipped bool) bool {
	return _NSMouseInRect(aPoint, aRect, flipped)
}

// Returns the next hash-table element in the enumeration.
//
// Added in macOS 10.0.
// Returns the next hash-table element in the enumeration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNextHashEnumeratorItem(_:)
func NSNextHashEnumeratorItem(enumerator unsafe.Pointer) unsafe.Pointer {
	return _NSNextHashEnumeratorItem(enumerator)
}

// Returns a Boolean value that indicates whether the next map-table pair in the enumeration are set.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether the next map-table pair in the enumeration are set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNextMapEnumeratorPair(_:_:_:)
func NSNextMapEnumeratorPair(enumerator unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _NSNextMapEnumeratorPair(enumerator, key, value)
}

// Offsets the rectangle by the specified amount.
//
// Added in macOS 10.0.
// Offsets the rectangle by the specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOffsetRect(_:_:_:)
func NSOffsetRect(aRect coregraphics.CGRect, dX float64, dY float64) coregraphics.CGRect {
	return _NSOffsetRect(aRect, dX, dY)
}

// Returns a point from a text-based representation.
//
// Added in macOS 10.0.
// Returns a point from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointFromString(_:)
func NSPointFromString(aString unsafe.Pointer) coregraphics.CGPoint {
	return _NSPointFromString(aString)
}

// Returns a Boolean value that indicates whether a given point is in a given rectangle.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether a given point is in a given rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointInRect(_:_:)
func NSPointInRect(aPoint coregraphics.CGPoint, aRect coregraphics.CGRect) bool {
	return _NSPointInRect(aPoint, aRect)
}

// Returns a range from a textual representation.
//
// Added in macOS 10.0.
// Returns a range from a textual representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeFromString(_:)
func NSRangeFromString(aString unsafe.Pointer) unsafe.Pointer {
	return _NSRangeFromString(aString)
}

// Returns a rectangle from a text-based representation.
//
// Added in macOS 10.0.
// Returns a rectangle from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectFromString(_:)
func NSRectFromString(aString unsafe.Pointer) coregraphics.CGRect {
	return _NSRectFromString(aString)
}

// Deletes the elements of the specified hash table.
//
// Added in macOS 10.0.
// Deletes the elements of the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSResetHashTable(_:)
func NSResetHashTable(table unsafe.Pointer) {
	_NSResetHashTable(table)
}

// Deletes the elements of the specified map table.
//
// Added in macOS 10.0.
// Deletes the elements of the specified map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSResetMapTable(_:)
func NSResetMapTable(table unsafe.Pointer) {
	_NSResetMapTable(table)
}

// Returns an from a text-based representation.
//
// Added in macOS 10.0.
// Returns an from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSizeFromString(_:)
func NSSizeFromString(aString unsafe.Pointer) coregraphics.CGSize {
	return _NSSizeFromString(aString)
}

// Returns a string describing the hash table’s contents.
//
// Added in macOS 10.0.
// Returns a string describing the hash table’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromHashTable(_:)
func NSStringFromHashTable(table unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromHashTable(table)
}

// Returns a string describing the map table’s contents.
//
// Added in macOS 10.0.
// Returns a string describing the map table’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromMapTable(_:)
func NSStringFromMapTable(table unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromMapTable(table)
}

// Returns a string representation of a point.
//
// Added in macOS 10.0.
// Returns a string representation of a point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromPoint(_:)
func NSStringFromPoint(aPoint coregraphics.CGPoint) unsafe.Pointer {
	return _NSStringFromPoint(aPoint)
}

// Returns a string representation of a range.
//
// Added in macOS 10.0.
// Returns a string representation of a range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromRange(_:)
func NSStringFromRange(range_ unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromRange(range_)
}

// Returns a string representation of a rectangle.
//
// Added in macOS 10.0.
// Returns a string representation of a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromRect(_:)
func NSStringFromRect(aRect coregraphics.CGRect) unsafe.Pointer {
	return _NSStringFromRect(aRect)
}

// Returns a string representation of a size.
//
// Added in macOS 10.0.
// Returns a string representation of a size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromSize(_:)
func NSStringFromSize(aSize coregraphics.CGSize) unsafe.Pointer {
	return _NSStringFromSize(aSize)
}

// Returns the union of the specified ranges.
//
// Added in macOS 10.0.
// Returns the union of the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnionRange(_:_:)
func NSUnionRange(range1 unsafe.Pointer, range2 unsafe.Pointer) unsafe.Pointer {
	return _NSUnionRange(range1, range2)
}

// Calculates the union of two rectangles.
//
// Added in macOS 10.0.
// Calculates the union of two rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnionRect(_:_:)
func NSUnionRect(aRect coregraphics.CGRect, bRect coregraphics.CGRect) coregraphics.CGRect {
	return _NSUnionRect(aRect, bRect)
}

// AXAnimatedImagesEnabled is a Foundation function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXAnimatedImagesEnabled
func AXAnimatedImagesEnabled() bool {
	return _AXAnimatedImagesEnabled()
}

// Returns a localized description of the color to use in accessibility attributes.
//
// Added in macOS 11.0.
// Returns a localized description of the color to use in accessibility attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXNameFromColor(_:)
func AXNameFromColor(color coregraphics.ColorRef) unsafe.Pointer {
	return _AXNameFromColor(color)
}

// AXOpenSettingsFeature is a Foundation function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXOpenSettingsFeature
func AXOpenSettingsFeature(feature unsafe.Pointer) {
	_AXOpenSettingsFeature(feature)
}

// AXPrefersActionSliderAlternative is a Foundation function.
//
// Added in macOS 26.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXPrefersActionSliderAlternative
func AXPrefersActionSliderAlternative() bool {
	return _AXPrefersActionSliderAlternative()
}

// AXPrefersHeadAnchorAlternative is a Foundation function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXPrefersHeadAnchorAlternative
func AXPrefersHeadAnchorAlternative() bool {
	return _AXPrefersHeadAnchorAlternative()
}

// AXPrefersHorizontalTextLayout is a Foundation function.
//
// Added in macOS 14.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXPrefersHorizontalTextLayout
func AXPrefersHorizontalTextLayout() bool {
	return _AXPrefersHorizontalTextLayout()
}

// AXPrefersNonBlinkingTextInsertionIndicator is a Foundation function.
//
// Added in macOS 15.0.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXPrefersNonBlinkingTextInsertionIndicator
func AXPrefersNonBlinkingTextInsertionIndicator() bool {
	return _AXPrefersNonBlinkingTextInsertionIndicator()
}

// AXShowBordersEnabled is a Foundation function.
//
// Added in macOS 26.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXShowBordersEnabled
func AXShowBordersEnabled() bool {
	return _AXShowBordersEnabled()
}

// A Boolean value that indicates whether Assistive Access is running.
//
// Added in macOS 15.0.
// A Boolean value that indicates whether Assistive Access is running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/isAssistiveAccessEnabled
func AXAssistiveAccessEnabled() bool {
	return _AXAssistiveAccessEnabled()
}

// Sends a notification to any observing assistive apps.

// Sends a notification to any observing assistive apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/post(element:notification:)
func NSAccessibilityPostNotification(element unsafe.Pointer, notification unsafe.Pointer) {
	_NSAccessibilityPostNotification(element, notification)
}

// Returns the point in screen coordinates.
//
// Added in macOS 10.10.
// Returns the point in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/screenPoint(fromView:point:)
func NSAccessibilityPointInView(parentView unsafe.Pointer, point coregraphics.CGPoint) coregraphics.CGPoint {
	return _NSAccessibilityPointInView(parentView, point)
}

// Returns the frame in screen coordinates.
//
// Added in macOS 10.10.
// Returns the frame in screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/screenRect(fromView:rect:)
func NSAccessibilityFrameInView(parentView unsafe.Pointer, frame coregraphics.CGRect) coregraphics.CGRect {
	return _NSAccessibilityFrameInView(parentView, frame)
}

// Returns the available window depth values.

// Returns the available window depth values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAvailableWindowDepths
func NSAvailableWindowDepths() unsafe.Pointer {
	return _NSAvailableWindowDepths()
}

// Attempts to return a window depth adequate for the specified parameters.

// Attempts to return a window depth adequate for the specified parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBestDepth
func NSBestDepth(colorSpace unsafe.Pointer, bps unsafe.Pointer, bpp unsafe.Pointer, planar bool, exactMatch unsafe.Pointer) unsafe.Pointer {
	return _NSBestDepth(colorSpace, bps, bpp, planar, exactMatch)
}

// Returns the number of color components in the specified color space.

// Returns the number of color components in the specified color space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/numberOfColorComponents
func NSNumberOfColorComponents(colorSpaceName unsafe.Pointer) unsafe.Pointer {
	return _NSNumberOfColorComponents(colorSpaceName)
}

// Draws a gray-filled rectangle representing a user-interface button.

// Draws a gray-filled rectangle representing a user-interface button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawButton(_:_:)
func NSDrawButton(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawButton(rect, clipRect)
}

// Draws a dark gray-filled rectangle with a bezel border.

// Draws a dark gray-filled rectangle with a bezel border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawDarkBezel(_:_:)
func NSDrawDarkBezel(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawDarkBezel(rect, clipRect)
}

// Draws a gray-filled rectangle with a bezel border.

// Draws a gray-filled rectangle with a bezel border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawGrayBezel(_:_:)
func NSDrawGrayBezel(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawGrayBezel(rect, clipRect)
}

// Draws a gray-filled rectangle with a groove border.

// Draws a gray-filled rectangle with a groove border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawGroove(_:_:)
func NSDrawGroove(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawGroove(rect, clipRect)
}

// Draws a white-filled rectangle with a bezel border.

// Draws a white-filled rectangle with a bezel border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawLightBezel(_:_:)
func NSDrawLightBezel(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawLightBezel(rect, clipRect)
}

// Draws a nine-part tiled image.
//
// Added in macOS 10.5.
// Draws a nine-part tiled image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawNinePartImage(_:_:_:_:_:_:_:_:_:_:_:_:_:)
func NSDrawNinePartImage(frame coregraphics.CGRect, topLeftCorner unsafe.Pointer, topEdgeFill unsafe.Pointer, topRightCorner unsafe.Pointer, leftEdgeFill unsafe.Pointer, centerFill unsafe.Pointer, rightEdgeFill unsafe.Pointer, bottomLeftCorner unsafe.Pointer, bottomEdgeFill unsafe.Pointer, bottomRightCorner unsafe.Pointer, op unsafe.Pointer, alphaFraction float64, flipped bool) {
	_NSDrawNinePartImage(frame, topLeftCorner, topEdgeFill, topRightCorner, leftEdgeFill, centerFill, rightEdgeFill, bottomLeftCorner, bottomEdgeFill, bottomRightCorner, op, alphaFraction, flipped)
}

// Draws a three-part tiled image.
//
// Added in macOS 10.5.
// Draws a three-part tiled image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawThreePartImage(_:_:_:_:_:_:_:_:)
func NSDrawThreePartImage(frame coregraphics.CGRect, startCap unsafe.Pointer, centerFill unsafe.Pointer, endCap unsafe.Pointer, vertical bool, op unsafe.Pointer, alphaFraction float64, flipped bool) {
	_NSDrawThreePartImage(frame, startCap, centerFill, endCap, vertical, op, alphaFraction, flipped)
}

// Draws rectangles with borders.

// Draws rectangles with borders.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawTiledRects(_:_:_:_:_:)
func NSDrawTiledRects(boundsRect coregraphics.CGRect, clipRect coregraphics.CGRect, sides unsafe.Pointer, grays []float64, count unsafe.Pointer) coregraphics.CGRect {
	return _NSDrawTiledRects(boundsRect, clipRect, sides, grays, count)
}

// Draws a white-filled rectangle with a bezel border.

// Draws a white-filled rectangle with a bezel border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawWhiteBezel(_:_:)
func NSDrawWhiteBezel(rect coregraphics.CGRect, clipRect coregraphics.CGRect) {
	_NSDrawWhiteBezel(rect, clipRect)
}

// Draws the window’s default background pattern into the specified rectangle of the currently focused view.

// Draws the window’s default background pattern into the specified rectangle of the currently focused view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawWindowBackground(_:)
func NSDrawWindowBackground(rect coregraphics.CGRect) {
	_NSDrawWindowBackground(rect)
}

// Erases the specified rect by filling it with white.

// Erases the specified rect by filling it with white.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEraseRect(_:)
func NSEraseRect(rect coregraphics.CGRect) {
	_NSEraseRect(rect)
}

// Specifies how the system draws the focus ring.

// Specifies how the system draws the focus ring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/set()
func NSSetFocusRingStyle(placement unsafe.Pointer) {
	_NSSetFocusRingStyle(placement)
}

// Draws a bordered rectangle.

// Draws a bordered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFrameRect
func NSFrameRect(rect coregraphics.CGRect) {
	_NSFrameRect(rect)
}

// Draws a bordered rectangle.

// Draws a bordered rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFrameRectWithWidth
func NSFrameRectWithWidth(rect coregraphics.CGRect, frameWidth float64) {
	_NSFrameRectWithWidth(rect, frameWidth)
}

// Draws a bordered rectangle using the specified compositing operation.

// Draws a bordered rectangle using the specified compositing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFrameRectWithWidthUsingOperation
func NSFrameRectWithWidthUsingOperation(rect coregraphics.CGRect, frameWidth float64, op unsafe.Pointer) {
	_NSFrameRectWithWidthUsingOperation(rect, frameWidth, op)
}

// Highlights the specified rect by filling it with white.
//
// Deprecated: This function was deprecated in macOS 10.0.
//
// Added in macOS 10.0.
// Highlights the specified rect by filling it with white.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHighlightRect
func NSHighlightRect(rect coregraphics.CGRect) {
	_NSHighlightRect(rect)
}

// Returns a pasteboard type based on the passed file type.

// Returns a pasteboard type based on the passed file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileContentsType(forPathExtension:)
func NSCreateFileContentsPboardType(fileType unsafe.Pointer) unsafe.Pointer {
	return _NSCreateFileContentsPboardType(fileType)
}

// Returns a pasteboard type based on the passed file type.

// Returns a pasteboard type based on the passed file type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/fileNameType(forPathExtension:)
func NSCreateFilenamePboardType(fileType unsafe.Pointer) unsafe.Pointer {
	return _NSCreateFilenamePboardType(fileType)
}

// A file type based on the passed pasteboard type.

// A file type based on the passed pasteboard type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/representedPathExtension
func NSGetFileType(pboardType unsafe.Pointer) unsafe.Pointer {
	return _NSGetFileType(pboardType)
}

// Modifies the current clipping path by intersecting it with the passed rect.

// Modifies the current clipping path by intersecting it with the passed rect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectClip
func NSRectClip(rect coregraphics.CGRect) {
	_NSRectClip(rect)
}

// Modifies the current clipping path by intersecting it with the passed rect.

// Modifies the current clipping path by intersecting it with the passed rect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectClipList
func NSRectClipList(rects unsafe.Pointer, count unsafe.Pointer) {
	_NSRectClipList(rects, count)
}

// Fills the passed rectangle with the current color.

// Fills the passed rectangle with the current color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFill
func NSRectFill(rect coregraphics.CGRect) {
	_NSRectFill(rect)
}

// Fills the rectangles in the passed list with the current fill color.

// Fills the rectangles in the passed list with the current fill color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillList
func NSRectFillList(rects unsafe.Pointer, count unsafe.Pointer) {
	_NSRectFillList(rects, count)
}

// Fills the rectangles in a list using the current fill color and specified compositing operation.

// Fills the rectangles in a list using the current fill color and specified compositing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillListUsingOperation
func NSRectFillListUsingOperation(rects unsafe.Pointer, count unsafe.Pointer, op unsafe.Pointer) {
	_NSRectFillListUsingOperation(rects, count, op)
}

// Fills the rectangles in the passed list with the passed list of colors.

// Fills the rectangles in the passed list with the passed list of colors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillListWithColors
func NSRectFillListWithColors(rects unsafe.Pointer, colors unsafe.Pointer, num unsafe.Pointer) {
	_NSRectFillListWithColors(rects, colors, num)
}

// Fills the rectangles in a list using the specified colors and compositing operation.

// Fills the rectangles in a list using the specified colors and compositing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillListWithColorsUsingOperation
func NSRectFillListWithColorsUsingOperation(rects unsafe.Pointer, colors unsafe.Pointer, num unsafe.Pointer, op unsafe.Pointer) {
	_NSRectFillListWithColorsUsingOperation(rects, colors, num, op)
}

// Fills the rectangles in the passed list with the passed list of grays.

// Fills the rectangles in the passed list with the passed list of grays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillListWithGrays
func NSRectFillListWithGrays(rects unsafe.Pointer, grays []float64, num unsafe.Pointer) {
	_NSRectFillListWithGrays(rects, grays, num)
}

// Fills a rectangle using the current fill color and the specified compositing operation.

// Fills a rectangle using the current fill color and the specified compositing operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectFillUsingOperation
func NSRectFillUsingOperation(rect coregraphics.CGRect, op unsafe.Pointer) {
	_NSRectFillUsingOperation(rect, op)
}

// Returns the bits per pixel for the specified window depth.

// Returns the bits per pixel for the specified window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/bitsPerPixel
func NSBitsPerPixelFromDepth(depth unsafe.Pointer) unsafe.Pointer {
	return _NSBitsPerPixelFromDepth(depth)
}

// Returns the bits per sample for the specified window depth.

// Returns the bits per sample for the specified window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/bitsPerSample
func NSBitsPerSampleFromDepth(depth unsafe.Pointer) unsafe.Pointer {
	return _NSBitsPerSampleFromDepth(depth)
}

// Returns the name of the color space corresponding to the passed window depth.

// Returns the name of the color space corresponding to the passed window depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/colorSpaceName
func NSColorSpaceFromDepth(depth unsafe.Pointer) unsafe.Pointer {
	return _NSColorSpaceFromDepth(depth)
}

// Returns whether the specified window depth is planar.

// Returns whether the specified window depth is planar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/isPlanar
func NSPlanarFromDepth(depth unsafe.Pointer) bool {
	return _NSPlanarFromDepth(depth)
}

// Allocates memory using the specified allocator.

// Allocates memory using the specified allocator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocate(_:_:_:)
func CFAllocatorAllocate(allocator unsafe.Pointer, size unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorAllocate(allocator, size, hint)
}

// Creates an allocator object.

// Creates an allocator object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorCreate(_:_:)
func CFAllocatorCreate(allocator unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorCreate(allocator, context)
}

// Deallocates a block of memory with a given allocator.

// Deallocates a block of memory with a given allocator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorDeallocate(_:_:)
func CFAllocatorDeallocate(allocator unsafe.Pointer, ptr unsafe.Pointer) {
	_CFAllocatorDeallocate(allocator, ptr)
}

// Obtains the context of the specified allocator or of the default allocator.

// Obtains the context of the specified allocator or of the default allocator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetContext(_:_:)
func CFAllocatorGetContext(allocator unsafe.Pointer, context unsafe.Pointer) {
	_CFAllocatorGetContext(allocator, context)
}

// Gets the default allocator object for the current thread.

// Gets the default allocator object for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetDefault()
func CFAllocatorGetDefault() unsafe.Pointer {
	return _CFAllocatorGetDefault()
}

// Obtains the number of bytes likely to be allocated upon a specific request.

// Obtains the number of bytes likely to be allocated upon a specific request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetPreferredSizeForSize(_:_:_:)
func CFAllocatorGetPreferredSizeForSize(allocator unsafe.Pointer, size unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorGetPreferredSizeForSize(allocator, size, hint)
}

// Returns the type identifier for the CFAllocator opaque type.

// Returns the type identifier for the CFAllocator opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorGetTypeID()
func CFAllocatorGetTypeID() unsafe.Pointer {
	return _CFAllocatorGetTypeID()
}

// Reallocates memory using the specified allocator.

// Reallocates memory using the specified allocator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocate(_:_:_:_:)
func CFAllocatorReallocate(allocator unsafe.Pointer, ptr unsafe.Pointer, newsize unsafe.Pointer, hint unsafe.Pointer) unsafe.Pointer {
	return _CFAllocatorReallocate(allocator, ptr, newsize, hint)
}

// Sets the given allocator as the default for the current thread.

// Sets the given allocator as the default for the current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorSetDefault(_:)
func CFAllocatorSetDefault(allocator unsafe.Pointer) {
	_CFAllocatorSetDefault(allocator)
}

// Calls a function once for each element in range in an array.

// Calls a function once for each element in range in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayApplyFunction(_:_:_:_:)
func CFArrayApplyFunction(theArray unsafe.Pointer, range_ unsafe.Pointer, applier unsafe.Pointer, context unsafe.Pointer) {
	_CFArrayApplyFunction(theArray, range_, applier, context)
}

// Searches an array for a value using a binary search algorithm.

// Searches an array for a value using a binary search algorithm.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayBSearchValues(_:_:_:_:_:)
func CFArrayBSearchValues(theArray unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer, comparator unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFArrayBSearchValues(theArray, range_, value, comparator, context)
}

// Reports whether or not a value is in an array.

// Reports whether or not a value is in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayContainsValue(_:_:_:)
func CFArrayContainsValue(theArray unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFArrayContainsValue(theArray, range_, value)
}

// Creates a new immutable array with the given values.

// Creates a new immutable array with the given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreate(_:_:_:_:)
func CFArrayCreate(allocator unsafe.Pointer, values unsafe.Pointer, numValues unsafe.Pointer, callBacks unsafe.Pointer) unsafe.Pointer {
	return _CFArrayCreate(allocator, values, numValues, callBacks)
}

// Creates a new immutable array with the values from another array.

// Creates a new immutable array with the values from another array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCreateCopy(_:_:)
func CFArrayCreateCopy(allocator unsafe.Pointer, theArray unsafe.Pointer) unsafe.Pointer {
	return _CFArrayCreateCopy(allocator, theArray)
}

// Returns the number of values currently in an array.

// Returns the number of values currently in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetCount(_:)
func CFArrayGetCount(theArray unsafe.Pointer) unsafe.Pointer {
	return _CFArrayGetCount(theArray)
}

// Counts the number of times a given value occurs in an array.

// Counts the number of times a given value occurs in an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetCountOfValue(_:_:_:)
func CFArrayGetCountOfValue(theArray unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFArrayGetCountOfValue(theArray, range_, value)
}

// Searches an array forward for a value.

// Searches an array forward for a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetFirstIndexOfValue(_:_:_:)
func CFArrayGetFirstIndexOfValue(theArray unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFArrayGetFirstIndexOfValue(theArray, range_, value)
}

// Searches an array backward for a value.

// Searches an array backward for a value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetLastIndexOfValue(_:_:_:)
func CFArrayGetLastIndexOfValue(theArray unsafe.Pointer, range_ unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFArrayGetLastIndexOfValue(theArray, range_, value)
}

// Returns the type identifier for the CFArray opaque type.

// Returns the type identifier for the CFArray opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetTypeID()
func CFArrayGetTypeID() unsafe.Pointer {
	return _CFArrayGetTypeID()
}

// Retrieves a value at a given index.

// Retrieves a value at a given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetValueAtIndex(_:_:)
func CFArrayGetValueAtIndex(theArray unsafe.Pointer, idx unsafe.Pointer) unsafe.Pointer {
	return _CFArrayGetValueAtIndex(theArray, idx)
}

// Fills a buffer with values from an array.

// Fills a buffer with values from an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayGetValues(_:_:_:)
func CFArrayGetValues(theArray unsafe.Pointer, range_ unsafe.Pointer, values unsafe.Pointer) {
	_CFArrayGetValues(theArray, range_, values)
}

// Defers internal consistency-checking and coalescing for a mutable attributed string.

// Defers internal consistency-checking and coalescing for a mutable attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringBeginEditing(_:)
func CFAttributedStringBeginEditing(aStr unsafe.Pointer) {
	_CFAttributedStringBeginEditing(aStr)
}

// Creates an attributed string with specified string and attributes.

// Creates an attributed string with specified string and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreate(_:_:_:)
func CFAttributedStringCreate(alloc unsafe.Pointer, str unsafe.Pointer, attributes unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreate(alloc, str, attributes)
}

// Creates an immutable copy of an attributed string.

// Creates an immutable copy of an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateCopy(_:_:)
func CFAttributedStringCreateCopy(alloc unsafe.Pointer, aStr unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreateCopy(alloc, aStr)
}

// Creates a mutable attributed string.

// Creates a mutable attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateMutable(_:_:)
func CFAttributedStringCreateMutable(alloc unsafe.Pointer, maxLength unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreateMutable(alloc, maxLength)
}

// Creates a mutable copy of an attributed string.

// Creates a mutable copy of an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateMutableCopy(_:_:_:)
func CFAttributedStringCreateMutableCopy(alloc unsafe.Pointer, maxLength unsafe.Pointer, aStr unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreateMutableCopy(alloc, maxLength, aStr)
}

// Creates a sub-attributed string from the specified range.

// Creates a sub-attributed string from the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateWithSubstring(_:_:_:)
func CFAttributedStringCreateWithSubstring(alloc unsafe.Pointer, aStr unsafe.Pointer, range_ unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreateWithSubstring(alloc, aStr, range_)
}

// Re-enables internal consistency-checking and coalescing for a mutable attributed string.

// Re-enables internal consistency-checking and coalescing for a mutable attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringEndEditing(_:)
func CFAttributedStringEndEditing(aStr unsafe.Pointer) {
	_CFAttributedStringEndEditing(aStr)
}

// Returns the value of a given attribute of an attributed string at a specified location.

// Returns the value of a given attribute of an attributed string at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttribute(_:_:_:_:)
func CFAttributedStringGetAttribute(aStr unsafe.Pointer, loc unsafe.Pointer, attrName unsafe.Pointer, effectiveRange unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetAttribute(aStr, loc, attrName, effectiveRange)
}

// Returns the value of a given attribute of an attributed string at a specified location.

// Returns the value of a given attribute of an attributed string at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributeAndLongestEffectiveRange(_:_:_:_:_:)
func CFAttributedStringGetAttributeAndLongestEffectiveRange(aStr unsafe.Pointer, loc unsafe.Pointer, attrName unsafe.Pointer, inRange unsafe.Pointer, longestEffectiveRange unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetAttributeAndLongestEffectiveRange(aStr, loc, attrName, inRange, longestEffectiveRange)
}

// Returns the attributes of an attributed string at a specified location.

// Returns the attributes of an attributed string at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributes(_:_:_:)
func CFAttributedStringGetAttributes(aStr unsafe.Pointer, loc unsafe.Pointer, effectiveRange unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetAttributes(aStr, loc, effectiveRange)
}

// Returns the attributes of an attributed string at a specified location.

// Returns the attributes of an attributed string at a specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetAttributesAndLongestEffectiveRange(_:_:_:_:)
func CFAttributedStringGetAttributesAndLongestEffectiveRange(aStr unsafe.Pointer, loc unsafe.Pointer, inRange unsafe.Pointer, longestEffectiveRange unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetAttributesAndLongestEffectiveRange(aStr, loc, inRange, longestEffectiveRange)
}

// Returns the length of the attributed string in characters.

// Returns the length of the attributed string in characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetLength(_:)
func CFAttributedStringGetLength(aStr unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetLength(aStr)
}

// Gets as a mutable string the string for an attributed string.

// Gets as a mutable string the string for an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetMutableString(_:)
func CFAttributedStringGetMutableString(aStr unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetMutableString(aStr)
}

// Returns the string for an attributed string.

// Returns the string for an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetString(_:)
func CFAttributedStringGetString(aStr unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringGetString(aStr)
}

// Returns the type identifier for the CFAttributedString opaque type.

// Returns the type identifier for the CFAttributedString opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringGetTypeID()
func CFAttributedStringGetTypeID() unsafe.Pointer {
	return _CFAttributedStringGetTypeID()
}

// Removes the value of a single attribute over a specified range.

// Removes the value of a single attribute over a specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringRemoveAttribute(_:_:_:)
func CFAttributedStringRemoveAttribute(aStr unsafe.Pointer, range_ unsafe.Pointer, attrName unsafe.Pointer) {
	_CFAttributedStringRemoveAttribute(aStr, range_, attrName)
}

// Replaces the attributed substring over a range with another attributed string.

// Replaces the attributed substring over a range with another attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringReplaceAttributedString(_:_:_:)
func CFAttributedStringReplaceAttributedString(aStr unsafe.Pointer, range_ unsafe.Pointer, replacement unsafe.Pointer) {
	_CFAttributedStringReplaceAttributedString(aStr, range_, replacement)
}

// Modifies the string of an attributed string.

// Modifies the string of an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringReplaceString(_:_:_:)
func CFAttributedStringReplaceString(aStr unsafe.Pointer, range_ unsafe.Pointer, replacement unsafe.Pointer) {
	_CFAttributedStringReplaceString(aStr, range_, replacement)
}

// Sets the value of a single attribute over the specified range.

// Sets the value of a single attribute over the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringSetAttribute(_:_:_:_:)
func CFAttributedStringSetAttribute(aStr unsafe.Pointer, range_ unsafe.Pointer, attrName unsafe.Pointer, value unsafe.Pointer) {
	_CFAttributedStringSetAttribute(aStr, range_, attrName, value)
}

// Sets the value of attributes of a mutable attributed string over a specified range.

// Sets the value of attributes of a mutable attributed string over a specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringSetAttributes(_:_:_:_:)
func CFAttributedStringSetAttributes(aStr unsafe.Pointer, range_ unsafe.Pointer, replacement unsafe.Pointer, clearOtherAttributes unsafe.Pointer) {
	_CFAttributedStringSetAttributes(aStr, range_, replacement, clearOtherAttributes)
}

// Calls a function once for each value in a bag.

// Calls a function once for each value in a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagApplyFunction(_:_:_:)
func CFBagApplyFunction(theBag unsafe.Pointer, applier unsafe.Pointer, context unsafe.Pointer) {
	_CFBagApplyFunction(theBag, applier, context)
}

// Reports whether or not a value is in a bag.

// Reports whether or not a value is in a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagContainsValue(_:_:)
func CFBagContainsValue(theBag unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagContainsValue(theBag, value)
}

// Creates an immutable bag containing specified values.

// Creates an immutable bag containing specified values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCreate(_:_:_:_:)
func CFBagCreate(allocator unsafe.Pointer, values unsafe.Pointer, numValues unsafe.Pointer, callBacks unsafe.Pointer) unsafe.Pointer {
	return _CFBagCreate(allocator, values, numValues, callBacks)
}

// Creates an immutable bag with the values of another bag.

// Creates an immutable bag with the values of another bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCreateCopy(_:_:)
func CFBagCreateCopy(allocator unsafe.Pointer, theBag unsafe.Pointer) unsafe.Pointer {
	return _CFBagCreateCopy(allocator, theBag)
}

// Returns the number of values currently in a bag.

// Returns the number of values currently in a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetCount(_:)
func CFBagGetCount(theBag unsafe.Pointer) unsafe.Pointer {
	return _CFBagGetCount(theBag)
}

// Returns the number of times a value occurs in a bag.

// Returns the number of times a value occurs in a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetCountOfValue(_:_:)
func CFBagGetCountOfValue(theBag unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagGetCountOfValue(theBag, value)
}

// Returns the type identifier for the CFBag opaque type.

// Returns the type identifier for the CFBag opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetTypeID()
func CFBagGetTypeID() unsafe.Pointer {
	return _CFBagGetTypeID()
}

// Returns a requested value from a bag.

// Returns a requested value from a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValue(_:_:)
func CFBagGetValue(theBag unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagGetValue(theBag, value)
}

// Reports whether or not a value is in a bag, and returns that value indirectly if it exists.

// Reports whether or not a value is in a bag, and returns that value indirectly if it exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValueIfPresent(_:_:_:)
func CFBagGetValueIfPresent(theBag unsafe.Pointer, candidate unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBagGetValueIfPresent(theBag, candidate, value)
}

// Fills a buffer with values from a bag.

// Fills a buffer with values from a bag.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagGetValues(_:_:)
func CFBagGetValues(theBag unsafe.Pointer, values unsafe.Pointer) {
	_CFBagGetValues(theBag, values)
}

// Adds a value to a binary heap.

// Adds a value to a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapAddValue(_:_:)
func CFBinaryHeapAddValue(heap unsafe.Pointer, value unsafe.Pointer) {
	_CFBinaryHeapAddValue(heap, value)
}

// Iteratively applies a function to all the values in a binary heap.

// Iteratively applies a function to all the values in a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapApplyFunction(_:_:_:)
func CFBinaryHeapApplyFunction(heap unsafe.Pointer, applier unsafe.Pointer, context unsafe.Pointer) {
	_CFBinaryHeapApplyFunction(heap, applier, context)
}

// Returns whether a given value is in a binary heap.

// Returns whether a given value is in a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapContainsValue(_:_:)
func CFBinaryHeapContainsValue(heap unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapContainsValue(heap, value)
}

// Creates a new mutable or fixed-mutable binary heap.

// Creates a new mutable or fixed-mutable binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCreate(_:_:_:_:)
func CFBinaryHeapCreate(allocator unsafe.Pointer, capacity unsafe.Pointer, callBacks unsafe.Pointer, compareContext unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapCreate(allocator, capacity, callBacks, compareContext)
}

// Creates a new mutable or fixed-mutable binary heap with the values from a pre-existing binary heap.

// Creates a new mutable or fixed-mutable binary heap with the values from a pre-existing binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapCreateCopy(_:_:_:)
func CFBinaryHeapCreateCopy(allocator unsafe.Pointer, capacity unsafe.Pointer, heap unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapCreateCopy(allocator, capacity, heap)
}

// Returns the number of values currently in a binary heap.

// Returns the number of values currently in a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetCount(_:)
func CFBinaryHeapGetCount(heap unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapGetCount(heap)
}

// Counts the number of times a given value occurs in a binary heap.

// Counts the number of times a given value occurs in a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetCountOfValue(_:_:)
func CFBinaryHeapGetCountOfValue(heap unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapGetCountOfValue(heap, value)
}

// Returns the minimum value in a binary heap.

// Returns the minimum value in a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetMinimum(_:)
func CFBinaryHeapGetMinimum(heap unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapGetMinimum(heap)
}

// Returns the minimum value in a binary heap, if present.

// Returns the minimum value in a binary heap, if present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetMinimumIfPresent(_:_:)
func CFBinaryHeapGetMinimumIfPresent(heap unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFBinaryHeapGetMinimumIfPresent(heap, value)
}

// Returns the type identifier of the opaque type.

// Returns the type identifier of the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetTypeID()
func CFBinaryHeapGetTypeID() unsafe.Pointer {
	return _CFBinaryHeapGetTypeID()
}

// Copies all the values from a binary heap into a sorted C array.

// Copies all the values from a binary heap into a sorted C array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapGetValues(_:_:)
func CFBinaryHeapGetValues(heap unsafe.Pointer, values unsafe.Pointer) {
	_CFBinaryHeapGetValues(heap, values)
}

// Removes all values from a binary heap, making it empty.

// Removes all values from a binary heap, making it empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapRemoveAllValues(_:)
func CFBinaryHeapRemoveAllValues(heap unsafe.Pointer) {
	_CFBinaryHeapRemoveAllValues(heap)
}

// Removes the minimum value from a binary heap.

// Removes the minimum value from a binary heap.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapRemoveMinimumValue(_:)
func CFBinaryHeapRemoveMinimumValue(heap unsafe.Pointer) {
	_CFBinaryHeapRemoveMinimumValue(heap)
}

// Returns the number of bit values in a bit vector.

// Returns the number of bit values in a bit vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVectorGetCount(_:)
func CFBitVectorGetCount(bv unsafe.Pointer) unsafe.Pointer {
	return _CFBitVectorGetCount(bv)
}

// Returns the Core Foundation type identifier for the CFBoolean opaque type.

// Returns the Core Foundation type identifier for the CFBoolean opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBooleanGetTypeID()
func CFBooleanGetTypeID() unsafe.Pointer {
	return _CFBooleanGetTypeID()
}

// Returns the value of a CFBoolean object as a standard C type .

// Returns the value of a CFBoolean object as a standard C type .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBooleanGetValue(_:)
func CFBooleanGetValue(boolean unsafe.Pointer) unsafe.Pointer {
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
func CFBundleCloseBundleResourceMap(bundle unsafe.Pointer, refNum unsafe.Pointer) {
	_CFBundleCloseBundleResourceMap(bundle, refNum)
}

// Returns the location of a bundle’s auxiliary executable code.

// Returns the location of a bundle’s auxiliary executable code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyAuxiliaryExecutableURL(_:_:)
func CFBundleCopyAuxiliaryExecutableURL(bundle unsafe.Pointer, executableName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyAuxiliaryExecutableURL(bundle, executableName)
}

// Returns the location of a bundle’s built in plug-in.

// Returns the location of a bundle’s built in plug-in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBuiltInPlugInsURL(_:)
func CFBundleCopyBuiltInPlugInsURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyBuiltInPlugInsURL(bundle)
}

// Returns an array containing a bundle’s localizations.

// Returns an array containing a bundle’s localizations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBundleLocalizations(_:)
func CFBundleCopyBundleLocalizations(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyBundleLocalizations(bundle)
}

// Returns the location of a bundle.

// Returns the location of a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyBundleURL(_:)
func CFBundleCopyBundleURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyBundleURL(bundle)
}

// Returns an array of CFNumbers representing the architectures a given bundle provides.
//
// Added in macOS 10.5.
// Returns an array of CFNumbers representing the architectures a given bundle provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableArchitectures(_:)
func CFBundleCopyExecutableArchitectures(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyExecutableArchitectures(bundle)
}

// Returns an array of CFNumbers representing the architectures a given URL provides.
//
// Added in macOS 10.5.
// Returns an array of CFNumbers representing the architectures a given URL provides.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableArchitecturesForURL(_:)
func CFBundleCopyExecutableArchitecturesForURL(url unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyExecutableArchitecturesForURL(url)
}

// Returns the location of a bundle’s main executable code.

// Returns the location of a bundle’s main executable code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyExecutableURL(_:)
func CFBundleCopyExecutableURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyExecutableURL(bundle)
}

// Returns the information dictionary for a given URL location.

// Returns the information dictionary for a given URL location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyInfoDictionaryForURL(_:)
func CFBundleCopyInfoDictionaryForURL(url unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyInfoDictionaryForURL(url)
}

// Returns a bundle’s information dictionary.

// Returns a bundle’s information dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyInfoDictionaryInDirectory(_:)
func CFBundleCopyInfoDictionaryInDirectory(bundleURL unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyInfoDictionaryInDirectory(bundleURL)
}

// Given an array of possible localizations and preferred locations, returns the one or more of them that CFBundle would use, without reference to the current application context.

// Given an array of possible localizations and preferred locations, returns the one or more of them that CFBundle would use, without reference to the current application context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizationsForPreferences(_:_:)
func CFBundleCopyLocalizationsForPreferences(locArray unsafe.Pointer, prefArray unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyLocalizationsForPreferences(locArray, prefArray)
}

// Returns an array containing the localizations for a bundle or executable at a particular location.

// Returns an array containing the localizations for a bundle or executable at a particular location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizationsForURL(_:)
func CFBundleCopyLocalizationsForURL(url unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyLocalizationsForURL(url)
}

// Returns a localized string from a bundle’s strings file.

// Returns a localized string from a bundle’s strings file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyLocalizedString(_:_:_:_:)
func CFBundleCopyLocalizedString(bundle unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer, tableName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyLocalizedString(bundle, key, value, tableName)
}

// Given an array of possible localizations, returns the one or more of them that CFBundle would use in the current application context.

// Given an array of possible localizations, returns the one or more of them that CFBundle would use in the current application context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyPreferredLocalizationsFromArray(_:)
func CFBundleCopyPreferredLocalizationsFromArray(locArray unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyPreferredLocalizationsFromArray(locArray)
}

// Returns the location of a bundle’s private Frameworks directory.

// Returns the location of a bundle’s private Frameworks directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyPrivateFrameworksURL(_:)
func CFBundleCopyPrivateFrameworksURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyPrivateFrameworksURL(bundle)
}

// Returns the location of a resource contained in the specified bundle.

// Returns the location of a resource contained in the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURL(_:_:_:_:)
func CFBundleCopyResourceURL(bundle unsafe.Pointer, resourceName unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURL(bundle, resourceName, resourceType, subDirName)
}

// Returns the location of a localized resource in a bundle.

// Returns the location of a localized resource in a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLForLocalization(_:_:_:_:_:)
func CFBundleCopyResourceURLForLocalization(bundle unsafe.Pointer, resourceName unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer, localizationName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURLForLocalization(bundle, resourceName, resourceType, subDirName, localizationName)
}

// Returns the location of a resource contained in the specified bundle directory without requiring the creation of a CFBundle object.

// Returns the location of a resource contained in the specified bundle directory without requiring the creation of a CFBundle object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLInDirectory(_:_:_:_:)
func CFBundleCopyResourceURLInDirectory(bundleURL unsafe.Pointer, resourceName unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURLInDirectory(bundleURL, resourceName, resourceType, subDirName)
}

// Assembles an array of URLs specifying all of the resources of the specified type found in a bundle.

// Assembles an array of URLs specifying all of the resources of the specified type found in a bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfType(_:_:_:)
func CFBundleCopyResourceURLsOfType(bundle unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURLsOfType(bundle, resourceType, subDirName)
}

// Returns an array containing copies of the URL locations for a specified bundle, resource, and localization name.

// Returns an array containing copies of the URL locations for a specified bundle, resource, and localization name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfTypeForLocalization(_:_:_:_:)
func CFBundleCopyResourceURLsOfTypeForLocalization(bundle unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer, localizationName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURLsOfTypeForLocalization(bundle, resourceType, subDirName, localizationName)
}

// Returns an array of CFURL objects describing the locations of all resources in a bundle of the specified type without needing to create a CFBundle object.

// Returns an array of CFURL objects describing the locations of all resources in a bundle of the specified type without needing to create a CFBundle object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourceURLsOfTypeInDirectory(_:_:_:)
func CFBundleCopyResourceURLsOfTypeInDirectory(bundleURL unsafe.Pointer, resourceType unsafe.Pointer, subDirName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourceURLsOfTypeInDirectory(bundleURL, resourceType, subDirName)
}

// Returns the location of a bundle’s Resources directory.

// Returns the location of a bundle’s Resources directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopyResourcesDirectoryURL(_:)
func CFBundleCopyResourcesDirectoryURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopyResourcesDirectoryURL(bundle)
}

// Returns the location of a bundle’s shared frameworks directory.

// Returns the location of a bundle’s shared frameworks directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySharedFrameworksURL(_:)
func CFBundleCopySharedFrameworksURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopySharedFrameworksURL(bundle)
}

// Returns the location of a bundle’s shared support files directory.

// Returns the location of a bundle’s shared support files directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySharedSupportURL(_:)
func CFBundleCopySharedSupportURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopySharedSupportURL(bundle)
}

// Returns the location of the bundle’s support files directory.

// Returns the location of the bundle’s support files directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCopySupportFilesDirectoryURL(_:)
func CFBundleCopySupportFilesDirectoryURL(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCopySupportFilesDirectoryURL(bundle)
}

// Creates a CFBundle object.

// Creates a CFBundle object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCreate(_:_:)
func CFBundleCreate(allocator unsafe.Pointer, bundleURL unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCreate(allocator, bundleURL)
}

// Searches a directory and constructs an array of CFBundle objects from all valid bundles in the specified directory.

// Searches a directory and constructs an array of CFBundle objects from all valid bundles in the specified directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleCreateBundlesFromDirectory(_:_:_:)
func CFBundleCreateBundlesFromDirectory(allocator unsafe.Pointer, directoryURL unsafe.Pointer, bundleType unsafe.Pointer) unsafe.Pointer {
	return _CFBundleCreateBundlesFromDirectory(allocator, directoryURL, bundleType)
}

// Returns an array containing all of the bundles currently open in the application.

// Returns an array containing all of the bundles currently open in the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetAllBundles()
func CFBundleGetAllBundles() unsafe.Pointer {
	return _CFBundleGetAllBundles()
}

// Locate a bundle given its program-defined identifier.

// Locate a bundle given its program-defined identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetBundleWithIdentifier(_:)
func CFBundleGetBundleWithIdentifier(bundleID unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetBundleWithIdentifier(bundleID)
}

// Returns a data pointer to a symbol of the given name.

// Returns a data pointer to a symbol of the given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDataPointerForName(_:_:)
func CFBundleGetDataPointerForName(bundle unsafe.Pointer, symbolName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetDataPointerForName(bundle, symbolName)
}

// Returns a C array of data pointer to symbols of the given names.

// Returns a C array of data pointer to symbols of the given names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDataPointersForNames(_:_:_:)
func CFBundleGetDataPointersForNames(bundle unsafe.Pointer, symbolNames unsafe.Pointer, stbl unsafe.Pointer) {
	_CFBundleGetDataPointersForNames(bundle, symbolNames, stbl)
}

// Returns the bundle’s development region from the bundle’s information property list.

// Returns the bundle’s development region from the bundle’s information property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetDevelopmentRegion(_:)
func CFBundleGetDevelopmentRegion(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetDevelopmentRegion(bundle)
}

// Returns a pointer to a function in a bundle’s executable code using the function name as the search key.

// Returns a pointer to a function in a bundle’s executable code using the function name as the search key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetFunctionPointerForName(_:_:)
func CFBundleGetFunctionPointerForName(bundle unsafe.Pointer, functionName unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetFunctionPointerForName(bundle, functionName)
}

// Constructs a function table containing pointers to all of the functions found in a bundle’s main executable code.

// Constructs a function table containing pointers to all of the functions found in a bundle’s main executable code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetFunctionPointersForNames(_:_:_:)
func CFBundleGetFunctionPointersForNames(bundle unsafe.Pointer, functionNames unsafe.Pointer, ftbl unsafe.Pointer) {
	_CFBundleGetFunctionPointersForNames(bundle, functionNames, ftbl)
}

// Returns the bundle identifier from a bundle’s information property list.

// Returns the bundle identifier from a bundle’s information property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetIdentifier(_:)
func CFBundleGetIdentifier(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetIdentifier(bundle)
}

// Returns a bundle’s information dictionary.

// Returns a bundle’s information dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetInfoDictionary(_:)
func CFBundleGetInfoDictionary(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetInfoDictionary(bundle)
}

// Returns a bundle’s localized information dictionary.

// Returns a bundle’s localized information dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetLocalInfoDictionary(_:)
func CFBundleGetLocalInfoDictionary(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetLocalInfoDictionary(bundle)
}

// Returns an application’s main bundle.

// Returns an application’s main bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetMainBundle()
func CFBundleGetMainBundle() unsafe.Pointer {
	return _CFBundleGetMainBundle()
}

// Returns a bundle’s package type and creator.

// Returns a bundle’s package type and creator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPackageInfo(_:_:_:)
func CFBundleGetPackageInfo(bundle unsafe.Pointer, packageType unsafe.Pointer, packageCreator unsafe.Pointer) {
	_CFBundleGetPackageInfo(bundle, packageType, packageCreator)
}

// Returns a bundle’s package type and creator without having to create a CFBundle object.

// Returns a bundle’s package type and creator without having to create a CFBundle object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPackageInfoInDirectory(_:_:_:)
func CFBundleGetPackageInfoInDirectory(url unsafe.Pointer, packageType unsafe.Pointer, packageCreator unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetPackageInfoInDirectory(url, packageType, packageCreator)
}

// Returns a bundle’s plug-in.

// Returns a bundle’s plug-in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetPlugIn(_:)
func CFBundleGetPlugIn(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetPlugIn(bundle)
}

// Returns the type identifier for the CFBundle opaque type.

// Returns the type identifier for the CFBundle opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetTypeID()
func CFBundleGetTypeID() unsafe.Pointer {
	return _CFBundleGetTypeID()
}

// Returns a value (localized if possible) from a bundle’s information dictionary.

// Returns a value (localized if possible) from a bundle’s information dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetValueForInfoDictionaryKey(_:_:)
func CFBundleGetValueForInfoDictionaryKey(bundle unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetValueForInfoDictionaryKey(bundle, key)
}

// Returns a bundle’s version number.

// Returns a bundle’s version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleGetVersionNumber(_:)
func CFBundleGetVersionNumber(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleGetVersionNumber(bundle)
}

// Obtains information about the load status for a bundle’s main executable.

// Obtains information about the load status for a bundle’s main executable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleIsExecutableLoaded(_:)
func CFBundleIsExecutableLoaded(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleIsExecutableLoaded(bundle)
}

// Loads a bundle’s main executable code into memory and dynamically links it into the running application.

// Loads a bundle’s main executable code into memory and dynamically links it into the running application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleLoadExecutable(_:)
func CFBundleLoadExecutable(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleLoadExecutable(bundle)
}

// Returns a Boolean value that indicates whether a given bundle is loaded, attempting to load it if necessary.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether a given bundle is loaded, attempting to load it if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleLoadExecutableAndReturnError(_:_:)
func CFBundleLoadExecutableAndReturnError(bundle unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
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
func CFBundleOpenBundleResourceFiles(bundle unsafe.Pointer, refNum unsafe.Pointer, localizedRefNum unsafe.Pointer) unsafe.Pointer {
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
func CFBundleOpenBundleResourceMap(bundle unsafe.Pointer) unsafe.Pointer {
	return _CFBundleOpenBundleResourceMap(bundle)
}

// Returns a Boolean value that indicates whether a given bundle is loaded or appears to be loadable.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether a given bundle is loaded or appears to be loadable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundlePreflightExecutable(_:_:)
func CFBundlePreflightExecutable(bundle unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFBundlePreflightExecutable(bundle, error_)
}

// Unloads the main executable for the specified bundle.

// Unloads the main executable for the specified bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleUnloadExecutable(_:)
func CFBundleUnloadExecutable(bundle unsafe.Pointer) {
	_CFBundleUnloadExecutable(bundle)
}

// Computes the absolute time when specified components are added to a given absolute time.

// Computes the absolute time when specified components are added to a given absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarAddComponents
func CFCalendarAddComponents(calendar unsafe.Pointer, at unsafe.Pointer, options unsafe.Pointer, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarAddComponents(calendar, at, options, componentDesc)
}

// Computes the absolute time from components in a description string.

// Computes the absolute time from components in a description string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarComposeAbsoluteTime
func CFCalendarComposeAbsoluteTime(calendar unsafe.Pointer, at unsafe.Pointer, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarComposeAbsoluteTime(calendar, at, componentDesc)
}

// Returns a copy of the logical calendar for the current user.

// Returns a copy of the logical calendar for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyCurrent()
func CFCalendarCopyCurrent() unsafe.Pointer {
	return _CFCalendarCopyCurrent()
}

// Returns a locale object for a specified calendar.

// Returns a locale object for a specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyLocale(_:)
func CFCalendarCopyLocale(calendar unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarCopyLocale(calendar)
}

// Returns a time zone object for a specified calendar.

// Returns a time zone object for a specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCopyTimeZone(_:)
func CFCalendarCopyTimeZone(calendar unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarCopyTimeZone(calendar)
}

// Returns a calendar object for the calendar identified by a calendar identifier.

// Returns a calendar object for the calendar identified by a calendar identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarCreateWithIdentifier(_:_:)
func CFCalendarCreateWithIdentifier(allocator unsafe.Pointer, identifier unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarCreateWithIdentifier(allocator, identifier)
}

// Computes the components which are indicated by the componentDesc description string for the given absolute time.

// Computes the components which are indicated by the componentDesc description string for the given absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarDecomposeAbsoluteTime
func CFCalendarDecomposeAbsoluteTime(calendar unsafe.Pointer, at unsafe.Pointer, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarDecomposeAbsoluteTime(calendar, at, componentDesc)
}

// Computes the difference between the two absolute times, in terms of specified calendrical components.

// Computes the difference between the two absolute times, in terms of specified calendrical components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetComponentDifference
func CFCalendarGetComponentDifference(calendar unsafe.Pointer, startingAT unsafe.Pointer, resultAT unsafe.Pointer, options unsafe.Pointer, componentDesc unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetComponentDifference(calendar, startingAT, resultAT, options, componentDesc)
}

// Returns the index of first weekday for a specified calendar.

// Returns the index of first weekday for a specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetFirstWeekday(_:)
func CFCalendarGetFirstWeekday(calendar unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetFirstWeekday(calendar)
}

// Returns the given calendar’s identifier.

// Returns the given calendar’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetIdentifier(_:)
func CFCalendarGetIdentifier(calendar unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetIdentifier(calendar)
}

// Returns the maximum range limits of the values that a specified unit can take on in a given calendar.

// Returns the maximum range limits of the values that a specified unit can take on in a given calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMaximumRangeOfUnit(_:_:)
func CFCalendarGetMaximumRangeOfUnit(calendar unsafe.Pointer, unit unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetMaximumRangeOfUnit(calendar, unit)
}

// Returns the minimum number of days in the first week of a specified calendar.

// Returns the minimum number of days in the first week of a specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMinimumDaysInFirstWeek(_:)
func CFCalendarGetMinimumDaysInFirstWeek(calendar unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetMinimumDaysInFirstWeek(calendar)
}

// Returns the minimum range limits of the values that a specified unit can take on in a given calendar.

// Returns the minimum range limits of the values that a specified unit can take on in a given calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetMinimumRangeOfUnit(_:_:)
func CFCalendarGetMinimumRangeOfUnit(calendar unsafe.Pointer, unit unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetMinimumRangeOfUnit(calendar, unit)
}

// Returns the ordinal number of a calendrical unit within a larger unit at a specified absolute time.

// Returns the ordinal number of a calendrical unit within a larger unit at a specified absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetOrdinalityOfUnit(_:_:_:_:)
func CFCalendarGetOrdinalityOfUnit(calendar unsafe.Pointer, smallerUnit unsafe.Pointer, biggerUnit unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetOrdinalityOfUnit(calendar, smallerUnit, biggerUnit, at)
}

// Returns the range of values that one unit can take on within a larger unit during which a specific absolute time occurs.

// Returns the range of values that one unit can take on within a larger unit during which a specific absolute time occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetRangeOfUnit(_:_:_:_:)
func CFCalendarGetRangeOfUnit(calendar unsafe.Pointer, smallerUnit unsafe.Pointer, biggerUnit unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetRangeOfUnit(calendar, smallerUnit, biggerUnit, at)
}

// Returns by reference the start time and duration of a given calendar unit that contains a given absolute time.
//
// Added in macOS 10.5.
// Returns by reference the start time and duration of a given calendar unit that contains a given absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetTimeRangeOfUnit(_:_:_:_:_:)
func CFCalendarGetTimeRangeOfUnit(calendar unsafe.Pointer, unit unsafe.Pointer, at unsafe.Pointer, startp unsafe.Pointer, tip unsafe.Pointer) unsafe.Pointer {
	return _CFCalendarGetTimeRangeOfUnit(calendar, unit, at, startp, tip)
}

// Returns the type identifier for the CFCalendar opaque type.

// Returns the type identifier for the CFCalendar opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarGetTypeID()
func CFCalendarGetTypeID() unsafe.Pointer {
	return _CFCalendarGetTypeID()
}

// Sets the first weekday for a calendar.

// Sets the first weekday for a calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetFirstWeekday(_:_:)
func CFCalendarSetFirstWeekday(calendar unsafe.Pointer, wkdy unsafe.Pointer) {
	_CFCalendarSetFirstWeekday(calendar, wkdy)
}

// Sets the locale for a calendar.

// Sets the locale for a calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetLocale(_:_:)
func CFCalendarSetLocale(calendar unsafe.Pointer, locale unsafe.Pointer) {
	_CFCalendarSetLocale(calendar, locale)
}

// Sets the minimum number of days in the first week of a specified calendar.

// Sets the minimum number of days in the first week of a specified calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetMinimumDaysInFirstWeek(_:_:)
func CFCalendarSetMinimumDaysInFirstWeek(calendar unsafe.Pointer, mwd unsafe.Pointer) {
	_CFCalendarSetMinimumDaysInFirstWeek(calendar, mwd)
}

// Sets the time zone for a calendar.

// Sets the time zone for a calendar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarSetTimeZone(_:_:)
func CFCalendarSetTimeZone(calendar unsafe.Pointer, tz unsafe.Pointer) {
	_CFCalendarSetTimeZone(calendar, tz)
}

// Adds a given range to a character set.

// Adds a given range to a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetAddCharactersInRange(_:_:)
func CFCharacterSetAddCharactersInRange(theSet unsafe.Pointer, theRange unsafe.Pointer) {
	_CFCharacterSetAddCharactersInRange(theSet, theRange)
}

// Adds the characters in a given string to a character set.

// Adds the characters in a given string to a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetAddCharactersInString(_:_:)
func CFCharacterSetAddCharactersInString(theSet unsafe.Pointer, theString unsafe.Pointer) {
	_CFCharacterSetAddCharactersInString(theSet, theString)
}

// Creates a new immutable data with the bitmap representation from the given character set.

// Creates a new immutable data with the bitmap representation from the given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateBitmapRepresentation(_:_:)
func CFCharacterSetCreateBitmapRepresentation(alloc unsafe.Pointer, theSet unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateBitmapRepresentation(alloc, theSet)
}

// Creates a new character set with the values from a given character set.

// Creates a new character set with the values from a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateCopy(_:_:)
func CFCharacterSetCreateCopy(alloc unsafe.Pointer, theSet unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateCopy(alloc, theSet)
}

// Creates a new immutable character set that is the invert of the specified character set.

// Creates a new immutable character set that is the invert of the specified character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateInvertedSet(_:_:)
func CFCharacterSetCreateInvertedSet(alloc unsafe.Pointer, theSet unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateInvertedSet(alloc, theSet)
}

// Creates a new empty mutable character set.

// Creates a new empty mutable character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateMutable(_:)
func CFCharacterSetCreateMutable(alloc unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateMutable(alloc)
}

// Creates a new mutable character set with the values from another character set.

// Creates a new mutable character set with the values from another character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateMutableCopy(_:_:)
func CFCharacterSetCreateMutableCopy(alloc unsafe.Pointer, theSet unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateMutableCopy(alloc, theSet)
}

// Creates a new immutable character set with the bitmap representation specified by given data.

// Creates a new immutable character set with the bitmap representation specified by given data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithBitmapRepresentation(_:_:)
func CFCharacterSetCreateWithBitmapRepresentation(alloc unsafe.Pointer, theData unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateWithBitmapRepresentation(alloc, theData)
}

// Creates a new character set with the values from the given range of Unicode characters.

// Creates a new character set with the values from the given range of Unicode characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithCharactersInRange(_:_:)
func CFCharacterSetCreateWithCharactersInRange(alloc unsafe.Pointer, theRange unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateWithCharactersInRange(alloc, theRange)
}

// Creates a new character set with the values in the given string.

// Creates a new character set with the values in the given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetCreateWithCharactersInString(_:_:)
func CFCharacterSetCreateWithCharactersInString(alloc unsafe.Pointer, theString unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetCreateWithCharactersInString(alloc, theString)
}

// Returns a predefined character set.

// Returns a predefined character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetGetPredefined(_:)
func CFCharacterSetGetPredefined(theSetIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetGetPredefined(theSetIdentifier)
}

// Returns the type identifier of the CFCharacterSet opaque type.

// Returns the type identifier of the CFCharacterSet opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetGetTypeID()
func CFCharacterSetGetTypeID() unsafe.Pointer {
	return _CFCharacterSetGetTypeID()
}

// Reports whether or not a character set contains at least one member character in the specified plane.

// Reports whether or not a character set contains at least one member character in the specified plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetHasMemberInPlane(_:_:)
func CFCharacterSetHasMemberInPlane(theSet unsafe.Pointer, thePlane unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetHasMemberInPlane(theSet, thePlane)
}

// Forms an intersection of two character sets.

// Forms an intersection of two character sets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIntersect(_:_:)
func CFCharacterSetIntersect(theSet unsafe.Pointer, theOtherSet unsafe.Pointer) {
	_CFCharacterSetIntersect(theSet, theOtherSet)
}

// Inverts the content of a given character set.

// Inverts the content of a given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetInvert(_:)
func CFCharacterSetInvert(theSet unsafe.Pointer) {
	_CFCharacterSetInvert(theSet)
}

// Reports whether or not a given Unicode character is in a character set.

// Reports whether or not a given Unicode character is in a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsCharacterMember(_:_:)
func CFCharacterSetIsCharacterMember(theSet unsafe.Pointer, theChar unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetIsCharacterMember(theSet, theChar)
}

// Reports whether or not a given UTF-32 character is in a character set.

// Reports whether or not a given UTF-32 character is in a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsLongCharacterMember(_:_:)
func CFCharacterSetIsLongCharacterMember(theSet unsafe.Pointer, theChar unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetIsLongCharacterMember(theSet, theChar)
}

// Reports whether or not a character set is a superset of another set.

// Reports whether or not a character set is a superset of another set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetIsSupersetOfSet(_:_:)
func CFCharacterSetIsSupersetOfSet(theSet unsafe.Pointer, theOtherset unsafe.Pointer) unsafe.Pointer {
	return _CFCharacterSetIsSupersetOfSet(theSet, theOtherset)
}

// Removes a given range of Unicode characters from a character set.

// Removes a given range of Unicode characters from a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetRemoveCharactersInRange(_:_:)
func CFCharacterSetRemoveCharactersInRange(theSet unsafe.Pointer, theRange unsafe.Pointer) {
	_CFCharacterSetRemoveCharactersInRange(theSet, theRange)
}

// Removes the characters in a given string from a character set.

// Removes the characters in a given string from a character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetRemoveCharactersInString(_:_:)
func CFCharacterSetRemoveCharactersInString(theSet unsafe.Pointer, theString unsafe.Pointer) {
	_CFCharacterSetRemoveCharactersInString(theSet, theString)
}

// Forms the union of two character sets.

// Forms the union of two character sets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetUnion(_:_:)
func CFCharacterSetUnion(theSet unsafe.Pointer, theOtherSet unsafe.Pointer) {
	_CFCharacterSetUnion(theSet, theOtherSet)
}

// Appends the bytes from a byte buffer to the contents of a CFData object.

// Appends the bytes from a byte buffer to the contents of a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataAppendBytes(_:_:_:)
func CFDataAppendBytes(theData unsafe.Pointer, bytes unsafe.Pointer, length unsafe.Pointer) {
	_CFDataAppendBytes(theData, bytes, length)
}

// Creates an immutable CFData object using data copied from a specified byte buffer.

// Creates an immutable CFData object using data copied from a specified byte buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreate(_:_:_:)
func CFDataCreate(allocator unsafe.Pointer, bytes unsafe.Pointer, length unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreate(allocator, bytes, length)
}

// Creates an immutable copy of a CFData object.

// Creates an immutable copy of a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateCopy(_:_:)
func CFDataCreateCopy(allocator unsafe.Pointer, theData unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreateCopy(allocator, theData)
}

// Creates an empty CFMutableData object.

// Creates an empty CFMutableData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateMutable(_:_:)
func CFDataCreateMutable(allocator unsafe.Pointer, capacity unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreateMutable(allocator, capacity)
}

// Creates a CFMutableData object by copying another CFData object.

// Creates a CFMutableData object by copying another CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateMutableCopy(_:_:_:)
func CFDataCreateMutableCopy(allocator unsafe.Pointer, capacity unsafe.Pointer, theData unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreateMutableCopy(allocator, capacity, theData)
}

// Creates an immutable CFData object from an external (client-owned) byte buffer.

// Creates an immutable CFData object from an external (client-owned) byte buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateWithBytesNoCopy(_:_:_:_:)
func CFDataCreateWithBytesNoCopy(allocator unsafe.Pointer, bytes unsafe.Pointer, length unsafe.Pointer, bytesDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreateWithBytesNoCopy(allocator, bytes, length, bytesDeallocator)
}

// Deletes the bytes in a CFMutableData object within a specified range.

// Deletes the bytes in a CFMutableData object within a specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataDeleteBytes(_:_:)
func CFDataDeleteBytes(theData unsafe.Pointer, range_ unsafe.Pointer) {
	_CFDataDeleteBytes(theData, range_)
}

// Finds and returns the range within a data object of the first occurrence of the given data, within a given range, subject to any given options.
//
// Added in macOS 10.6.
// Finds and returns the range within a data object of the first occurrence of the given data, within a given range, subject to any given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataFind(_:_:_:_:)
func CFDataFind(theData unsafe.Pointer, dataToFind unsafe.Pointer, searchRange unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFDataFind(theData, dataToFind, searchRange, compareOptions)
}

// Returns a read-only pointer to the bytes of a CFData object.

// Returns a read-only pointer to the bytes of a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetBytePtr(_:)
func CFDataGetBytePtr(theData unsafe.Pointer) unsafe.Pointer {
	return _CFDataGetBytePtr(theData)
}

// Copies the byte contents of a CFData object to an external buffer.

// Copies the byte contents of a CFData object to an external buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetBytes(_:_:_:)
func CFDataGetBytes(theData unsafe.Pointer, range_ unsafe.Pointer, buffer unsafe.Pointer) {
	_CFDataGetBytes(theData, range_, buffer)
}

// Returns the number of bytes contained by a CFData object.

// Returns the number of bytes contained by a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetLength(_:)
func CFDataGetLength(theData unsafe.Pointer) unsafe.Pointer {
	return _CFDataGetLength(theData)
}

// Returns a pointer to a mutable byte buffer of a CFMutableData object.

// Returns a pointer to a mutable byte buffer of a CFMutableData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetMutableBytePtr(_:)
func CFDataGetMutableBytePtr(theData unsafe.Pointer) unsafe.Pointer {
	return _CFDataGetMutableBytePtr(theData)
}

// Returns the type identifier for the CFData opaque type.

// Returns the type identifier for the CFData opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetTypeID()
func CFDataGetTypeID() unsafe.Pointer {
	return _CFDataGetTypeID()
}

// Increases the length of a CFMutableData object’s internal byte buffer, zero-filling the extension to the buffer.

// Increases the length of a CFMutableData object’s internal byte buffer, zero-filling the extension to the buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataIncreaseLength(_:_:)
func CFDataIncreaseLength(theData unsafe.Pointer, extraLength unsafe.Pointer) {
	_CFDataIncreaseLength(theData, extraLength)
}

// Replaces those bytes in a CFMutableData object that fall within a specified range with other bytes.

// Replaces those bytes in a CFMutableData object that fall within a specified range with other bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataReplaceBytes(_:_:_:_:)
func CFDataReplaceBytes(theData unsafe.Pointer, range_ unsafe.Pointer, newBytes unsafe.Pointer, newLength unsafe.Pointer) {
	_CFDataReplaceBytes(theData, range_, newBytes, newLength)
}

// Resets the length of a CFMutableData object’s internal byte buffer.

// Resets the length of a CFMutableData object’s internal byte buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataSetLength(_:_:)
func CFDataSetLength(theData unsafe.Pointer, length unsafe.Pointer) {
	_CFDataSetLength(theData, length)
}

// Compares two objects and returns a comparison result.

// Compares two objects and returns a comparison result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateCompare(_:_:_:)
func CFDateCompare(theDate unsafe.Pointer, otherDate unsafe.Pointer, context unsafe.Pointer) unsafe.Pointer {
	return _CFDateCompare(theDate, otherDate, context)
}

// Creates a object given an absolute time.

// Creates a object given an absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateCreate(_:_:)
func CFDateCreate(allocator unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFDateCreate(allocator, at)
}

// Returns a copy of a date formatter’s value for a given key.

// Returns a copy of a date formatter’s value for a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCopyProperty(_:_:)
func CFDateFormatterCopyProperty(formatter unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCopyProperty(formatter, key)
}

// Creates a new CFDateFormatter object, localized to the given locale, which will format dates to the given date and time styles.

// Creates a new CFDateFormatter object, localized to the given locale, which will format dates to the given date and time styles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreate(_:_:_:_:)
func CFDateFormatterCreate(allocator unsafe.Pointer, locale unsafe.Pointer, dateStyle unsafe.Pointer, timeStyle unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreate(allocator, locale, dateStyle, timeStyle)
}

// Returns a localized date format string representing the given date format components arranged appropriately for the specified locale.
//
// Added in macOS 10.6.
// Returns a localized date format string representing the given date format components arranged appropriately for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateDateFormatFromTemplate(_:_:_:_:)
func CFDateFormatterCreateDateFormatFromTemplate(allocator unsafe.Pointer, tmplate unsafe.Pointer, options unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreateDateFormatFromTemplate(allocator, tmplate, options, locale)
}

// Returns a date object representing a given string.

// Returns a date object representing a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateDateFromString(_:_:_:_:)
func CFDateFormatterCreateDateFromString(allocator unsafe.Pointer, formatter unsafe.Pointer, string_ unsafe.Pointer, rangep unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreateDateFromString(allocator, formatter, string_, rangep)
}

// Returns a string representation of the given absolute time using the specified date formatter.

// Returns a string representation of the given absolute time using the specified date formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateStringWithAbsoluteTime(_:_:_:)
func CFDateFormatterCreateStringWithAbsoluteTime(allocator unsafe.Pointer, formatter unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreateStringWithAbsoluteTime(allocator, formatter, at)
}

// Returns a string representation of the given date using the specified date formatter.

// Returns a string representation of the given date using the specified date formatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterCreateStringWithDate(_:_:_:)
func CFDateFormatterCreateStringWithDate(allocator unsafe.Pointer, formatter unsafe.Pointer, date unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterCreateStringWithDate(allocator, formatter, date)
}

// Returns an absolute time object representing a given string.

// Returns an absolute time object representing a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetAbsoluteTimeFromString(_:_:_:_:)
func CFDateFormatterGetAbsoluteTimeFromString(formatter unsafe.Pointer, string_ unsafe.Pointer, rangep unsafe.Pointer, atp unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetAbsoluteTimeFromString(formatter, string_, rangep, atp)
}

// Returns the date style used to create the given date formatter object.

// Returns the date style used to create the given date formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetDateStyle(_:)
func CFDateFormatterGetDateStyle(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetDateStyle(formatter)
}

// Returns a format string for the given date formatter object.

// Returns a format string for the given date formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetFormat(_:)
func CFDateFormatterGetFormat(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetFormat(formatter)
}

// Returns the locale object used to create the given date formatter object.

// Returns the locale object used to create the given date formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetLocale(_:)
func CFDateFormatterGetLocale(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetLocale(formatter)
}

// Returns the time style used to create the given date formatter object.

// Returns the time style used to create the given date formatter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetTimeStyle(_:)
func CFDateFormatterGetTimeStyle(formatter unsafe.Pointer) unsafe.Pointer {
	return _CFDateFormatterGetTimeStyle(formatter)
}

// Returns the type identifier for CFDateFormatter.

// Returns the type identifier for CFDateFormatter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterGetTypeID()
func CFDateFormatterGetTypeID() unsafe.Pointer {
	return _CFDateFormatterGetTypeID()
}

// Sets the format string of the given date formatter to the specified value.

// Sets the format string of the given date formatter to the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterSetFormat(_:_:)
func CFDateFormatterSetFormat(formatter unsafe.Pointer, formatString unsafe.Pointer) {
	_CFDateFormatterSetFormat(formatter, formatString)
}

// Sets a date formatter property using a key-value pair.

// Sets a date formatter property using a key-value pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterSetProperty(_:_:_:)
func CFDateFormatterSetProperty(formatter unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_CFDateFormatterSetProperty(formatter, key, value)
}

// Returns a object’s absolute time.

// Returns a object’s absolute time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateGetAbsoluteTime(_:)
func CFDateGetAbsoluteTime(theDate unsafe.Pointer) unsafe.Pointer {
	return _CFDateGetAbsoluteTime(theDate)
}

// Returns the number of elapsed seconds between the given objects.

// Returns the number of elapsed seconds between the given objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateGetTimeIntervalSinceDate(_:_:)
func CFDateGetTimeIntervalSinceDate(theDate unsafe.Pointer, otherDate unsafe.Pointer) unsafe.Pointer {
	return _CFDateGetTimeIntervalSinceDate(theDate, otherDate)
}

// Returns the type identifier for the opaque type.

// Returns the type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateGetTypeID()
func CFDateGetTypeID() unsafe.Pointer {
	return _CFDateGetTypeID()
}

// Calls a function once for each key-value pair in a dictionary.

// Calls a function once for each key-value pair in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryApplyFunction(_:_:_:)
func CFDictionaryApplyFunction(theDict unsafe.Pointer, applier unsafe.Pointer, context unsafe.Pointer) {
	_CFDictionaryApplyFunction(theDict, applier, context)
}

// Returns a Boolean value that indicates whether a given key is in a dictionary.

// Returns a Boolean value that indicates whether a given key is in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryContainsKey(_:_:)
func CFDictionaryContainsKey(theDict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryContainsKey(theDict, key)
}

// Returns a Boolean value that indicates whether a given value is in a dictionary.

// Returns a Boolean value that indicates whether a given value is in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryContainsValue(_:_:)
func CFDictionaryContainsValue(theDict unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryContainsValue(theDict, value)
}

// Creates an immutable dictionary containing the specified key-value pairs.

// Creates an immutable dictionary containing the specified key-value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreate(_:_:_:_:_:_:)
func CFDictionaryCreate(allocator unsafe.Pointer, keys unsafe.Pointer, values unsafe.Pointer, numValues unsafe.Pointer, keyCallBacks unsafe.Pointer, valueCallBacks unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryCreate(allocator, keys, values, numValues, keyCallBacks, valueCallBacks)
}

// Creates and returns a new immutable dictionary with the key-value pairs of another dictionary.

// Creates and returns a new immutable dictionary with the key-value pairs of another dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCreateCopy(_:_:)
func CFDictionaryCreateCopy(allocator unsafe.Pointer, theDict unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryCreateCopy(allocator, theDict)
}

// Returns the number of key-value pairs in a dictionary.

// Returns the number of key-value pairs in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCount(_:)
func CFDictionaryGetCount(theDict unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetCount(theDict)
}

// Returns the number of times a key occurs in a dictionary.

// Returns the number of times a key occurs in a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCountOfKey(_:_:)
func CFDictionaryGetCountOfKey(theDict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetCountOfKey(theDict, key)
}

// Counts the number of times a given value occurs in the dictionary.

// Counts the number of times a given value occurs in the dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetCountOfValue(_:_:)
func CFDictionaryGetCountOfValue(theDict unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetCountOfValue(theDict, value)
}

// Fills two buffers with the keys and values from a dictionary.

// Fills two buffers with the keys and values from a dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetKeysAndValues(_:_:_:)
func CFDictionaryGetKeysAndValues(theDict unsafe.Pointer, keys unsafe.Pointer, values unsafe.Pointer) {
	_CFDictionaryGetKeysAndValues(theDict, keys, values)
}

// Returns the type identifier for the CFDictionary opaque type.

// Returns the type identifier for the CFDictionary opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetTypeID()
func CFDictionaryGetTypeID() unsafe.Pointer {
	return _CFDictionaryGetTypeID()
}

// Returns the value associated with a given key.

// Returns the value associated with a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetValue(_:_:)
func CFDictionaryGetValue(theDict unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetValue(theDict, key)
}

// Returns a Boolean value that indicates whether a given value for a given key is in a dictionary, and returns that value indirectly if it exists.

// Returns a Boolean value that indicates whether a given value for a given key is in a dictionary, and returns that value indirectly if it exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryGetValueIfPresent(_:_:_:)
func CFDictionaryGetValueIfPresent(theDict unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFDictionaryGetValueIfPresent(theDict, key, value)
}

// Determines whether two Core Foundation objects are considered equal.

// Determines whether two Core Foundation objects are considered equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFEqual(_:_:)
func CFEqual(cf1 unsafe.Pointer, cf2 unsafe.Pointer) unsafe.Pointer {
	return _CFEqual(cf1, cf2)
}

// Creates a new CFError object.
//
// Added in macOS 10.5.
// Creates a new CFError object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorCreate(_:_:_:_:)
func CFErrorCreate(allocator unsafe.Pointer, domain unsafe.Pointer, code unsafe.Pointer, userInfo unsafe.Pointer) unsafe.Pointer {
	return _CFErrorCreate(allocator, domain, code, userInfo)
}

// Returns the native file descriptor for a given CFFileDescriptor.
//
// Added in macOS 10.5.
// Returns the native file descriptor for a given CFFileDescriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorGetNativeDescriptor(_:)
func CFFileDescriptorGetNativeDescriptor(f unsafe.Pointer) unsafe.Pointer {
	return _CFFileDescriptorGetNativeDescriptor(f)
}

// Returns a Boolean value that indicates whether the native file descriptor for a given CFFileDescriptor is valid.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether the native file descriptor for a given CFFileDescriptor is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorIsValid(_:)
func CFFileDescriptorIsValid(f unsafe.Pointer) unsafe.Pointer {
	return _CFFileDescriptorIsValid(f)
}

// Returns a code that can be used to identify an object in a hashing structure.

// Returns a code that can be used to identify an object in a hashing structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFHash(_:)
func CFHash(cf unsafe.Pointer) unsafe.Pointer {
	return _CFHash(cf)
}

// Returns an array of CFString objects that represents all locales for which locale data is available.

// Returns an array of CFString objects that represents all locales for which locale data is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyAvailableLocaleIdentifiers()
func CFLocaleCopyAvailableLocaleIdentifiers() unsafe.Pointer {
	return _CFLocaleCopyAvailableLocaleIdentifiers()
}

// Returns an array of strings that represents ISO currency codes for currencies in common use.
//
// Added in macOS 10.5.
// Returns an array of strings that represents ISO currency codes for currencies in common use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyCommonISOCurrencyCodes()
func CFLocaleCopyCommonISOCurrencyCodes() unsafe.Pointer {
	return _CFLocaleCopyCommonISOCurrencyCodes()
}

// Returns a copy of the logical locale for the current user.

// Returns a copy of the logical locale for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyCurrent()
func CFLocaleCopyCurrent() unsafe.Pointer {
	return _CFLocaleCopyCurrent()
}

// Returns the display name for the given value.

// Returns the display name for the given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyDisplayNameForPropertyValue(_:_:_:)
func CFLocaleCopyDisplayNameForPropertyValue(displayLocale unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCopyDisplayNameForPropertyValue(displayLocale, key, value)
}

// Returns an array of CFString objects that represents all known legal ISO country codes.

// Returns an array of CFString objects that represents all known legal ISO country codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOCountryCodes()
func CFLocaleCopyISOCountryCodes() unsafe.Pointer {
	return _CFLocaleCopyISOCountryCodes()
}

// Returns an array of CFString objects that represents all known legal ISO currency codes.

// Returns an array of CFString objects that represents all known legal ISO currency codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOCurrencyCodes()
func CFLocaleCopyISOCurrencyCodes() unsafe.Pointer {
	return _CFLocaleCopyISOCurrencyCodes()
}

// Returns an array of CFString objects that represents all known legal ISO language codes.

// Returns an array of CFString objects that represents all known legal ISO language codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyISOLanguageCodes()
func CFLocaleCopyISOLanguageCodes() unsafe.Pointer {
	return _CFLocaleCopyISOLanguageCodes()
}

// Returns the array of canonicalized language IDs that the user prefers.
//
// Added in macOS 10.5.
// Returns the array of canonicalized language IDs that the user prefers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCopyPreferredLanguages()
func CFLocaleCopyPreferredLanguages() unsafe.Pointer {
	return _CFLocaleCopyPreferredLanguages()
}

// Creates a locale for the given arbitrary locale identifier.

// Creates a locale for the given arbitrary locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreate(_:_:)
func CFLocaleCreate(allocator unsafe.Pointer, localeIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreate(allocator, localeIdentifier)
}

// Returns a canonical language identifier by mapping an arbitrary locale identification string to the canonical identifier

// Returns a canonical language identifier by mapping an arbitrary locale identification string to the canonical identifier
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLanguageIdentifierFromString(_:_:)
func CFLocaleCreateCanonicalLanguageIdentifierFromString(allocator unsafe.Pointer, localeIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateCanonicalLanguageIdentifierFromString(allocator, localeIdentifier)
}

// Returns a canonical locale identifier from given language and region codes.

// Returns a canonical locale identifier from given language and region codes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(_:_:_:)
func CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator unsafe.Pointer, lcode unsafe.Pointer, rcode unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator, lcode, rcode)
}

// Returns a canonical locale identifier by mapping an arbitrary locale identification string to the canonical identifier.

// Returns a canonical locale identifier by mapping an arbitrary locale identification string to the canonical identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCanonicalLocaleIdentifierFromString(_:_:)
func CFLocaleCreateCanonicalLocaleIdentifierFromString(allocator unsafe.Pointer, localeIdentifier unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateCanonicalLocaleIdentifierFromString(allocator, localeIdentifier)
}

// Returns a dictionary containing the result from parsing a locale ID consisting of language, script, country or region, variant, and keyword/value pairs.

// Returns a dictionary containing the result from parsing a locale ID consisting of language, script, country or region, variant, and keyword/value pairs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateComponentsFromLocaleIdentifier(_:_:)
func CFLocaleCreateComponentsFromLocaleIdentifier(allocator unsafe.Pointer, localeID unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateComponentsFromLocaleIdentifier(allocator, localeID)
}

// Returns a copy of a locale.

// Returns a copy of a locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateCopy(_:_:)
func CFLocaleCreateCopy(allocator unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateCopy(allocator, locale)
}

// Returns a locale identifier consisting of language, script, country or region, variant, and keyword/value pairs derived from a dictionary containing the source information.

// Returns a locale identifier consisting of language, script, country or region, variant, and keyword/value pairs derived from a dictionary containing the source information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateLocaleIdentifierFromComponents(_:_:)
func CFLocaleCreateLocaleIdentifierFromComponents(allocator unsafe.Pointer, dictionary unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleCreateLocaleIdentifierFromComponents(allocator, dictionary)
}

// Returns a locale identifier from a Windows locale code.
//
// Added in macOS 10.6.
// Returns a locale identifier from a Windows locale code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(_:_:)
func CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator unsafe.Pointer, lcid uint32) unsafe.Pointer {
	return _CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator, lcid)
}

// Returns the given locale’s identifier.

// Returns the given locale’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetIdentifier(_:)
func CFLocaleGetIdentifier(locale unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleGetIdentifier(locale)
}

// Returns the character direction for the specified ISO language code.
//
// Added in macOS 10.6.
// Returns the character direction for the specified ISO language code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetLanguageCharacterDirection(_:)
func CFLocaleGetLanguageCharacterDirection(isoLangCode unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleGetLanguageCharacterDirection(isoLangCode)
}

// Returns the line direction for the specified ISO language code.
//
// Added in macOS 10.6.
// Returns the line direction for the specified ISO language code.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetLanguageLineDirection(_:)
func CFLocaleGetLanguageLineDirection(isoLangCode unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleGetLanguageLineDirection(isoLangCode)
}

// Returns the root, canonical locale.

// Returns the root, canonical locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetSystem()
func CFLocaleGetSystem() unsafe.Pointer {
	return _CFLocaleGetSystem()
}

// Returns the type identifier for the CFLocale opaque type.

// Returns the type identifier for the CFLocale opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetTypeID()
func CFLocaleGetTypeID() unsafe.Pointer {
	return _CFLocaleGetTypeID()
}

// Returns the corresponding value for the given key of a locale’s key-value pair.

// Returns the corresponding value for the given key of a locale’s key-value pair.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetValue(_:_:)
func CFLocaleGetValue(locale unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _CFLocaleGetValue(locale, key)
}

// Returns a Windows locale code from the locale identifier.
//
// Added in macOS 10.6.
// Returns a Windows locale code from the locale identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(_:)
func CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier unsafe.Pointer) uint32 {
	return _CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier)
}

// Returns the type identifier for the CFNull opaque type.

// Returns the type identifier for the CFNull opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNullGetTypeID()
func CFNullGetTypeID() unsafe.Pointer {
	return _CFNullGetTypeID()
}

// Determines whether a CFNumber object contains a value stored as one of the defined floating point types.

// Determines whether a CFNumber object contains a value stored as one of the defined floating point types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberIsFloatType(_:)
func CFNumberIsFloatType(number unsafe.Pointer) unsafe.Pointer {
	return _CFNumberIsFloatType(number)
}

// Forces a CFRunLoop object to stop running.

// Forces a CFRunLoop object to stop running.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopStop(_:)
func CFRunLoopStop(rl unsafe.Pointer) {
	_CFRunLoopStop(rl)
}

// Prints the attributes of a string during debugging.

// Prints the attributes of a string during debugging.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFShowStr(_:)
func CFShowStr(str unsafe.Pointer) {
	_CFShowStr(str)
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
func CFStreamCreatePairWithSocketToHost(alloc unsafe.Pointer, host unsafe.Pointer, port unsafe.Pointer, readStream unsafe.Pointer, writeStream unsafe.Pointer) {
	_CFStreamCreatePairWithSocketToHost(alloc, host, port, readStream, writeStream)
}

// Compares one string with another string.

// Compares one string with another string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompare(_:_:_:)
func CFStringCompare(theString1 unsafe.Pointer, theString2 unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFStringCompare(theString1, theString2, compareOptions)
}

// Compares a range of the characters in one string with that of another string.

// Compares a range of the characters in one string with that of another string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareWithOptions(_:_:_:_:)
func CFStringCompareWithOptions(theString1 unsafe.Pointer, theString2 unsafe.Pointer, rangeToCompare unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFStringCompareWithOptions(theString1, theString2, rangeToCompare, compareOptions)
}

// Compares a range of the characters in one string with another string using a given locale.
//
// Added in macOS 10.5.
// Compares a range of the characters in one string with another string using a given locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareWithOptionsAndLocale(_:_:_:_:_:)
func CFStringCompareWithOptionsAndLocale(theString1 unsafe.Pointer, theString2 unsafe.Pointer, rangeToCompare unsafe.Pointer, compareOptions unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _CFStringCompareWithOptionsAndLocale(theString1, theString2, rangeToCompare, compareOptions, locale)
}

// Returns the name of the IANA registry “charset” that is the closest mapping to a specified string encoding.

// Returns the name of the IANA registry “charset” that is the closest mapping to a specified string encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToIANACharSetName(_:)
func CFStringConvertEncodingToIANACharSetName(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertEncodingToIANACharSetName(encoding)
}

// Returns the Cocoa encoding constant that maps most closely to a given Core Foundation encoding constant.

// Returns the Cocoa encoding constant that maps most closely to a given Core Foundation encoding constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToNSStringEncoding(_:)
func CFStringConvertEncodingToNSStringEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertEncodingToNSStringEncoding(encoding)
}

// Returns the Windows codepage identifier that maps most closely to a given Core Foundation encoding constant.

// Returns the Windows codepage identifier that maps most closely to a given Core Foundation encoding constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertEncodingToWindowsCodepage(_:)
func CFStringConvertEncodingToWindowsCodepage(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertEncodingToWindowsCodepage(encoding)
}

// Returns the Core Foundation encoding constant that is the closest mapping to a given IANA registry “charset” name.

// Returns the Core Foundation encoding constant that is the closest mapping to a given IANA registry “charset” name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertIANACharSetNameToEncoding(_:)
func CFStringConvertIANACharSetNameToEncoding(theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertIANACharSetNameToEncoding(theString)
}

// Returns the Core Foundation encoding constant that is the closest mapping to a given Cocoa encoding.

// Returns the Core Foundation encoding constant that is the closest mapping to a given Cocoa encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertNSStringEncodingToEncoding(_:)
func CFStringConvertNSStringEncodingToEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertNSStringEncodingToEncoding(encoding)
}

// Returns the Core Foundation encoding constant that is the closest mapping to a given Windows codepage identifier.

// Returns the Core Foundation encoding constant that is the closest mapping to a given Windows codepage identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringConvertWindowsCodepageToEncoding(_:)
func CFStringConvertWindowsCodepageToEncoding(codepage unsafe.Pointer) unsafe.Pointer {
	return _CFStringConvertWindowsCodepageToEncoding(codepage)
}

// Creates an array of CFString objects from a single CFString object.

// Creates an array of CFString objects from a single CFString object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateArrayBySeparatingStrings(_:_:_:)
func CFStringCreateArrayBySeparatingStrings(alloc unsafe.Pointer, theString unsafe.Pointer, separatorString unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateArrayBySeparatingStrings(alloc, theString, separatorString)
}

// Searches a string for multiple occurrences of a substring and creates an array of ranges identifying the locations of these substrings within the target string.

// Searches a string for multiple occurrences of a substring and creates an array of ranges identifying the locations of these substrings within the target string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateArrayWithFindResults(_:_:_:_:_:)
func CFStringCreateArrayWithFindResults(alloc unsafe.Pointer, theString unsafe.Pointer, stringToFind unsafe.Pointer, rangeToSearch unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateArrayWithFindResults(alloc, theString, stringToFind, rangeToSearch, compareOptions)
}

// Creates a single string from the individual CFString objects that comprise the elements of an array.

// Creates a single string from the individual CFString objects that comprise the elements of an array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateByCombiningStrings(_:_:_:)
func CFStringCreateByCombiningStrings(alloc unsafe.Pointer, theArray unsafe.Pointer, separatorString unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateByCombiningStrings(alloc, theArray, separatorString)
}

// Creates an immutable copy of a string.

// Creates an immutable copy of a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateCopy(_:_:)
func CFStringCreateCopy(alloc unsafe.Pointer, theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateCopy(alloc, theString)
}

// Creates an “external representation” of a CFString object, that is, a CFData object.

// Creates an “external representation” of a CFString object, that is, a CFData object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateExternalRepresentation(_:_:_:_:)
func CFStringCreateExternalRepresentation(alloc unsafe.Pointer, theString unsafe.Pointer, encoding unsafe.Pointer, lossByte unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateExternalRepresentation(alloc, theString, encoding, lossByte)
}

// Creates a string from its “external representation.”

// Creates a string from its “external representation.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateFromExternalRepresentation(_:_:_:)
func CFStringCreateFromExternalRepresentation(alloc unsafe.Pointer, data unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateFromExternalRepresentation(alloc, data, encoding)
}

// Creates a string from a buffer containing characters in a specified encoding.

// Creates a string from a buffer containing characters in a specified encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithBytes(_:_:_:_:_:)
func CFStringCreateWithBytes(alloc unsafe.Pointer, bytes unsafe.Pointer, numBytes unsafe.Pointer, encoding unsafe.Pointer, isExternalRepresentation unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithBytes(alloc, bytes, numBytes, encoding, isExternalRepresentation)
}

// Creates a string from a buffer, containing characters in a specified encoding, that might serve as the backing store for the new string.

// Creates a string from a buffer, containing characters in a specified encoding, that might serve as the backing store for the new string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithBytesNoCopy(_:_:_:_:_:_:)
func CFStringCreateWithBytesNoCopy(alloc unsafe.Pointer, bytes unsafe.Pointer, numBytes unsafe.Pointer, encoding unsafe.Pointer, isExternalRepresentation unsafe.Pointer, contentsDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithBytesNoCopy(alloc, bytes, numBytes, encoding, isExternalRepresentation, contentsDeallocator)
}

// Creates an immutable string from a C string.

// Creates an immutable string from a C string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCString(_:_:_:)
func CFStringCreateWithCString(alloc unsafe.Pointer, cStr unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithCString(alloc, cStr, encoding)
}

// Creates a CFString object from an external C string buffer that might serve as the backing store for the object.

// Creates a CFString object from an external C string buffer that might serve as the backing store for the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCStringNoCopy(_:_:_:_:)
func CFStringCreateWithCStringNoCopy(alloc unsafe.Pointer, cStr unsafe.Pointer, encoding unsafe.Pointer, contentsDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithCStringNoCopy(alloc, cStr, encoding, contentsDeallocator)
}

// Creates a string from a buffer of Unicode characters.

// Creates a string from a buffer of Unicode characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCharacters(_:_:_:)
func CFStringCreateWithCharacters(alloc unsafe.Pointer, chars unsafe.Pointer, numChars unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithCharacters(alloc, chars, numChars)
}

// Creates a string from a buffer of Unicode characters that might serve as the backing store for the object.

// Creates a string from a buffer of Unicode characters that might serve as the backing store for the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithCharactersNoCopy(_:_:_:_:)
func CFStringCreateWithCharactersNoCopy(alloc unsafe.Pointer, chars unsafe.Pointer, numChars unsafe.Pointer, contentsDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithCharactersNoCopy(alloc, chars, numChars, contentsDeallocator)
}

// Creates a CFString from a zero-terminated POSIX file system representation.

// Creates a CFString from a zero-terminated POSIX file system representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFileSystemRepresentation(_:_:)
func CFStringCreateWithFileSystemRepresentation(alloc unsafe.Pointer, buffer unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithFileSystemRepresentation(alloc, buffer)
}

// Creates an immutable string from a formatted string and a variable number of arguments.

// Creates an immutable string from a formatted string and a variable number of arguments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFormat
func CFStringCreateWithFormat(alloc unsafe.Pointer, formatOptions unsafe.Pointer, format unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithFormat(alloc, formatOptions, format)
}

// Creates an immutable string from a formatted string and a variable number of arguments (specified in a parameter of type ).

// Creates an immutable string from a formatted string and a variable number of arguments (specified in a parameter of type ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithFormatAndArguments(_:_:_:_:)
func CFStringCreateWithFormatAndArguments(alloc unsafe.Pointer, formatOptions unsafe.Pointer, format unsafe.Pointer, arguments unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithFormatAndArguments(alloc, formatOptions, format, arguments)
}

// Creates an immutable CFString object from a Pascal string.

// Creates an immutable CFString object from a Pascal string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithPascalString(_:_:_:)
func CFStringCreateWithPascalString(alloc unsafe.Pointer, pStr unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithPascalString(alloc, pStr, encoding)
}

// Creates a CFString object from an external Pascal string buffer that might serve as the backing store for the object.

// Creates a CFString object from an external Pascal string buffer that might serve as the backing store for the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithPascalStringNoCopy(_:_:_:_:)
func CFStringCreateWithPascalStringNoCopy(alloc unsafe.Pointer, pStr unsafe.Pointer, encoding unsafe.Pointer, contentsDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithPascalStringNoCopy(alloc, pStr, encoding, contentsDeallocator)
}

// Creates an immutable string from a segment (substring) of an existing string.

// Creates an immutable string from a segment (substring) of an existing string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringCreateWithSubstring(_:_:_:)
func CFStringCreateWithSubstring(alloc unsafe.Pointer, str unsafe.Pointer, range_ unsafe.Pointer) unsafe.Pointer {
	return _CFStringCreateWithSubstring(alloc, str, range_)
}

// Searches for a substring within a string and, if it is found, yields the range of the substring within the object’s characters.

// Searches for a substring within a string and, if it is found, yields the range of the substring within the object’s characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFind(_:_:_:)
func CFStringFind(theString unsafe.Pointer, stringToFind unsafe.Pointer, compareOptions unsafe.Pointer) unsafe.Pointer {
	return _CFStringFind(theString, stringToFind, compareOptions)
}

// Query the range of the first character contained in the specified character set.

// Query the range of the first character contained in the specified character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindCharacterFromSet(_:_:_:_:_:)
func CFStringFindCharacterFromSet(theString unsafe.Pointer, theSet unsafe.Pointer, rangeToSearch unsafe.Pointer, searchOptions unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _CFStringFindCharacterFromSet(theString, theSet, rangeToSearch, searchOptions, result)
}

// Searches for a substring within a range of the characters represented by a string and, if the substring is found, returns its range within the object’s characters.

// Searches for a substring within a range of the characters represented by a string and, if the substring is found, returns its range within the object’s characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindWithOptions(_:_:_:_:_:)
func CFStringFindWithOptions(theString unsafe.Pointer, stringToFind unsafe.Pointer, rangeToSearch unsafe.Pointer, searchOptions unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _CFStringFindWithOptions(theString, stringToFind, rangeToSearch, searchOptions, result)
}

// Returns a Boolean value that indicates whether a given string was found in a given source string.
//
// Added in macOS 10.5.
// Returns a Boolean value that indicates whether a given string was found in a given source string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringFindWithOptionsAndLocale(_:_:_:_:_:_:)
func CFStringFindWithOptionsAndLocale(theString unsafe.Pointer, stringToFind unsafe.Pointer, rangeToSearch unsafe.Pointer, searchOptions unsafe.Pointer, locale unsafe.Pointer, result unsafe.Pointer) unsafe.Pointer {
	return _CFStringFindWithOptionsAndLocale(theString, stringToFind, rangeToSearch, searchOptions, locale, result)
}

// Fetches a range of the characters from a string into a byte buffer after converting the characters to a specified encoding.

// Fetches a range of the characters from a string into a byte buffer after converting the characters to a specified encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetBytes(_:_:_:_:_:_:_:_:)
func CFStringGetBytes(theString unsafe.Pointer, range_ unsafe.Pointer, encoding unsafe.Pointer, lossByte unsafe.Pointer, isExternalRepresentation unsafe.Pointer, buffer unsafe.Pointer, maxBufLen unsafe.Pointer, usedBufLen unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetBytes(theString, range_, encoding, lossByte, isExternalRepresentation, buffer, maxBufLen, usedBufLen)
}

// Copies the character contents of a string to a local C string buffer after converting the characters to a given encoding.

// Copies the character contents of a string to a local C string buffer after converting the characters to a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCString(_:_:_:_:)
func CFStringGetCString(theString unsafe.Pointer, buffer unsafe.Pointer, bufferSize unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetCString(theString, buffer, bufferSize, encoding)
}

// Quickly obtains a pointer to a C-string buffer containing the characters of a string in a given encoding.

// Quickly obtains a pointer to a C-string buffer containing the characters of a string in a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCStringPtr(_:_:)
func CFStringGetCStringPtr(theString unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetCStringPtr(theString, encoding)
}

// Returns the Unicode character at a specified location in a string.

// Returns the Unicode character at a specified location in a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharacterAtIndex(_:_:)
func CFStringGetCharacterAtIndex(theString unsafe.Pointer, idx unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetCharacterAtIndex(theString, idx)
}

// Copies a range of the Unicode characters from a string to a user-provided buffer.

// Copies a range of the Unicode characters from a string to a user-provided buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharacters(_:_:_:)
func CFStringGetCharacters(theString unsafe.Pointer, range_ unsafe.Pointer, buffer unsafe.Pointer) {
	_CFStringGetCharacters(theString, range_, buffer)
}

// Quickly obtains a pointer to the contents of a string as a buffer of Unicode characters.

// Quickly obtains a pointer to the contents of a string as a buffer of Unicode characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetCharactersPtr(_:)
func CFStringGetCharactersPtr(theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetCharactersPtr(theString)
}

// Returns the primary value represented by a string.

// Returns the primary value represented by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetDoubleValue(_:)
func CFStringGetDoubleValue(str unsafe.Pointer) float64 {
	return _CFStringGetDoubleValue(str)
}

// Returns for a CFString object the character encoding that requires the least conversion time.

// Returns for a CFString object the character encoding that requires the least conversion time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetFastestEncoding(_:)
func CFStringGetFastestEncoding(theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetFastestEncoding(theString)
}

// Extracts the contents of a string as a -terminated 8-bit string appropriate for passing to POSIX APIs.

// Extracts the contents of a string as a -terminated 8-bit string appropriate for passing to POSIX APIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetFileSystemRepresentation(_:_:_:)
func CFStringGetFileSystemRepresentation(string_ unsafe.Pointer, buffer unsafe.Pointer, maxBufLen unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetFileSystemRepresentation(string_, buffer, maxBufLen)
}

// Retrieve the first potential hyphenation location found before the specified location.
//
// Added in macOS 10.7.
// Retrieve the first potential hyphenation location found before the specified location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetHyphenationLocationBeforeIndex(_:_:_:_:_:_:)
func CFStringGetHyphenationLocationBeforeIndex(string_ unsafe.Pointer, location unsafe.Pointer, limitRange unsafe.Pointer, options unsafe.Pointer, locale unsafe.Pointer, character unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetHyphenationLocationBeforeIndex(string_, location, limitRange, options, locale, character)
}

// Returns the integer value represented by a string.

// Returns the integer value represented by a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetIntValue(_:)
func CFStringGetIntValue(str unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetIntValue(str)
}

// Returns the number (in terms of UTF-16 code pairs) of Unicode characters in a string.

// Returns the number (in terms of UTF-16 code pairs) of Unicode characters in a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetLength(_:)
func CFStringGetLength(theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetLength(theString)
}

// Given a range of characters in a string, obtains the line bounds—that is, the indexes of the first character and the final characters of the lines containing the range.

// Given a range of characters in a string, obtains the line bounds—that is, the indexes of the first character and the final characters of the lines containing the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetLineBounds(_:_:_:_:_:)
func CFStringGetLineBounds(theString unsafe.Pointer, range_ unsafe.Pointer, lineBeginIndex unsafe.Pointer, lineEndIndex unsafe.Pointer, contentsEndIndex unsafe.Pointer) {
	_CFStringGetLineBounds(theString, range_, lineBeginIndex, lineEndIndex, contentsEndIndex)
}

// Returns a pointer to a list of string encodings supported by the current system.

// Returns a pointer to a list of string encodings supported by the current system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetListOfAvailableEncodings()
func CFStringGetListOfAvailableEncodings() unsafe.Pointer {
	return _CFStringGetListOfAvailableEncodings()
}

// Returns the maximum number of bytes a string of a specified length (in Unicode characters) will take up if encoded in a specified encoding.

// Returns the maximum number of bytes a string of a specified length (in Unicode characters) will take up if encoded in a specified encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMaximumSizeForEncoding(_:_:)
func CFStringGetMaximumSizeForEncoding(length unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetMaximumSizeForEncoding(length, encoding)
}

// Determines the upper bound on the number of bytes required to hold the file system representation of the string.

// Determines the upper bound on the number of bytes required to hold the file system representation of the string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMaximumSizeOfFileSystemRepresentation(_:)
func CFStringGetMaximumSizeOfFileSystemRepresentation(string_ unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetMaximumSizeOfFileSystemRepresentation(string_)
}

// Returns the most compatible Mac OS script value for the given input encoding.

// Returns the most compatible Mac OS script value for the given input encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetMostCompatibleMacStringEncoding(_:)
func CFStringGetMostCompatibleMacStringEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetMostCompatibleMacStringEncoding(encoding)
}

// Returns the canonical name of a specified string encoding.

// Returns the canonical name of a specified string encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetNameOfEncoding(_:)
func CFStringGetNameOfEncoding(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetNameOfEncoding(encoding)
}

// Given a range of characters in a string, obtains the paragraph bounds—that is, the indexes of the first character and the final characters of the paragraph(s) containing the range.
//
// Added in macOS 10.5.
// Given a range of characters in a string, obtains the paragraph bounds—that is, the indexes of the first character and the final characters of the paragraph(s) containing the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetParagraphBounds(_:_:_:_:_:)
func CFStringGetParagraphBounds(string_ unsafe.Pointer, range_ unsafe.Pointer, parBeginIndex unsafe.Pointer, parEndIndex unsafe.Pointer, contentsEndIndex unsafe.Pointer) {
	_CFStringGetParagraphBounds(string_, range_, parBeginIndex, parEndIndex, contentsEndIndex)
}

// Copies the character contents of a CFString object to a local Pascal string buffer after converting the characters to a requested encoding.

// Copies the character contents of a CFString object to a local Pascal string buffer after converting the characters to a requested encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetPascalString(_:_:_:_:)
func CFStringGetPascalString(theString unsafe.Pointer, buffer unsafe.Pointer, bufferSize unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetPascalString(theString, buffer, bufferSize, encoding)
}

// Quickly obtains a pointer to a Pascal buffer containing the characters of a string in a given encoding.

// Quickly obtains a pointer to a Pascal buffer containing the characters of a string in a given encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetPascalStringPtr(_:_:)
func CFStringGetPascalStringPtr(theString unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetPascalStringPtr(theString, encoding)
}

// Returns the range of the composed character sequence at a specified index.

// Returns the range of the composed character sequence at a specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetRangeOfComposedCharactersAtIndex(_:_:)
func CFStringGetRangeOfComposedCharactersAtIndex(theString unsafe.Pointer, theIndex unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetRangeOfComposedCharactersAtIndex(theString, theIndex)
}

// Returns the smallest encoding on the current system for the character contents of a string.

// Returns the smallest encoding on the current system for the character contents of a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetSmallestEncoding(_:)
func CFStringGetSmallestEncoding(theString unsafe.Pointer) unsafe.Pointer {
	return _CFStringGetSmallestEncoding(theString)
}

// Returns the default encoding used by the operating system when it creates strings.

// Returns the default encoding used by the operating system when it creates strings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetSystemEncoding()
func CFStringGetSystemEncoding() unsafe.Pointer {
	return _CFStringGetSystemEncoding()
}

// Returns the type identifier for the CFString opaque type.

// Returns the type identifier for the CFString opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringGetTypeID()
func CFStringGetTypeID() unsafe.Pointer {
	return _CFStringGetTypeID()
}

// Determines if the character data of a string begin with a specified sequence of characters.

// Determines if the character data of a string begin with a specified sequence of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringHasPrefix(_:_:)
func CFStringHasPrefix(theString unsafe.Pointer, prefix unsafe.Pointer) unsafe.Pointer {
	return _CFStringHasPrefix(theString, prefix)
}

// Determines if a string ends with a specified sequence of characters.

// Determines if a string ends with a specified sequence of characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringHasSuffix(_:_:)
func CFStringHasSuffix(theString unsafe.Pointer, suffix unsafe.Pointer) unsafe.Pointer {
	return _CFStringHasSuffix(theString, suffix)
}

// Determines whether a given Core Foundation string encoding is available on the current system.

// Determines whether a given Core Foundation string encoding is available on the current system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringIsEncodingAvailable(_:)
func CFStringIsEncodingAvailable(encoding unsafe.Pointer) unsafe.Pointer {
	return _CFStringIsEncodingAvailable(encoding)
}

// Returns a Boolean value that indicates whether hyphenation data is available.
//
// Added in macOS 10.7.
// Returns a Boolean value that indicates whether hyphenation data is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringIsHyphenationAvailableForLocale(_:)
func CFStringIsHyphenationAvailableForLocale(locale unsafe.Pointer) unsafe.Pointer {
	return _CFStringIsHyphenationAvailableForLocale(locale)
}

// Returns the abbreviation of a time zone at a specified date.

// Returns the abbreviation of a time zone at a specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyAbbreviation(_:_:)
func CFTimeZoneCopyAbbreviation(tz unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneCopyAbbreviation(tz, at)
}

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names.

// Returns a dictionary holding the mappings of time zone abbreviations to time zone names.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyAbbreviationDictionary()
func CFTimeZoneCopyAbbreviationDictionary() unsafe.Pointer {
	return _CFTimeZoneCopyAbbreviationDictionary()
}

// Returns the default time zone set for your application.

// Returns the default time zone set for your application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyDefault()
func CFTimeZoneCopyDefault() unsafe.Pointer {
	return _CFTimeZoneCopyDefault()
}

// Returns an array of strings containing the names of all the time zones known to the system.

// Returns an array of strings containing the names of all the time zones known to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyKnownNames()
func CFTimeZoneCopyKnownNames() unsafe.Pointer {
	return _CFTimeZoneCopyKnownNames()
}

// Returns the localized name of a given time zone.
//
// Added in macOS 10.5.
// Returns the localized name of a given time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopyLocalizedName(_:_:_:)
func CFTimeZoneCopyLocalizedName(tz unsafe.Pointer, style unsafe.Pointer, locale unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneCopyLocalizedName(tz, style, locale)
}

// Returns the time zone currently used by the system.

// Returns the time zone currently used by the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCopySystem()
func CFTimeZoneCopySystem() unsafe.Pointer {
	return _CFTimeZoneCopySystem()
}

// Creates a time zone with a given name and data.

// Creates a time zone with a given name and data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreate(_:_:_:)
func CFTimeZoneCreate(allocator unsafe.Pointer, name unsafe.Pointer, data unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneCreate(allocator, name, data)
}

// Returns the time zone object identified by a given name or abbreviation.

// Returns the time zone object identified by a given name or abbreviation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreateWithName(_:_:_:)
func CFTimeZoneCreateWithName(allocator unsafe.Pointer, name unsafe.Pointer, tryAbbrev unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneCreateWithName(allocator, name, tryAbbrev)
}

// Returns a time zone object for the specified time interval offset from Greenwich Mean Time (GMT).

// Returns a time zone object for the specified time interval offset from Greenwich Mean Time (GMT).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneCreateWithTimeIntervalFromGMT(_:_:)
func CFTimeZoneCreateWithTimeIntervalFromGMT(allocator unsafe.Pointer, ti unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneCreateWithTimeIntervalFromGMT(allocator, ti)
}

// Returns the data that stores the information used by a time zone.

// Returns the data that stores the information used by a time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetData(_:)
func CFTimeZoneGetData(tz unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneGetData(tz)
}

// Returns the daylight saving time offset for a time zone at a given time.
//
// Added in macOS 10.5.
// Returns the daylight saving time offset for a time zone at a given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetDaylightSavingTimeOffset(_:_:)
func CFTimeZoneGetDaylightSavingTimeOffset(tz unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneGetDaylightSavingTimeOffset(tz, at)
}

// Returns the geopolitical region name that identifies a given time zone.

// Returns the geopolitical region name that identifies a given time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetName(_:)
func CFTimeZoneGetName(tz unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneGetName(tz)
}

// Returns the time in a given time zone of the next daylight saving time transition after a given time.
//
// Added in macOS 10.5.
// Returns the time in a given time zone of the next daylight saving time transition after a given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetNextDaylightSavingTimeTransition(_:_:)
func CFTimeZoneGetNextDaylightSavingTimeTransition(tz unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneGetNextDaylightSavingTimeTransition(tz, at)
}

// Returns the difference in seconds between the receiver and Greenwich Mean Time (GMT) at the specified date.

// Returns the difference in seconds between the receiver and Greenwich Mean Time (GMT) at the specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetSecondsFromGMT(_:_:)
func CFTimeZoneGetSecondsFromGMT(tz unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneGetSecondsFromGMT(tz, at)
}

// Returns the type identifier for the CFTimeZone opaque type.

// Returns the type identifier for the CFTimeZone opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneGetTypeID()
func CFTimeZoneGetTypeID() unsafe.Pointer {
	return _CFTimeZoneGetTypeID()
}

// Returns whether or not a time zone is in daylight savings time at a specified date.

// Returns whether or not a time zone is in daylight savings time at a specified date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneIsDaylightSavingTime(_:_:)
func CFTimeZoneIsDaylightSavingTime(tz unsafe.Pointer, at unsafe.Pointer) unsafe.Pointer {
	return _CFTimeZoneIsDaylightSavingTime(tz, at)
}

// Clears the previously determined system time zone, if any.

// Clears the previously determined system time zone, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneResetSystem()
func CFTimeZoneResetSystem() {
	_CFTimeZoneResetSystem()
}

// Sets the abbreviation dictionary to a given dictionary.

// Sets the abbreviation dictionary to a given dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneSetAbbreviationDictionary(_:)
func CFTimeZoneSetAbbreviationDictionary(dict unsafe.Pointer) {
	_CFTimeZoneSetAbbreviationDictionary(dict)
}

// Sets the default time zone for your application the given time zone.

// Sets the default time zone for your application the given time zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneSetDefault(_:)
func CFTimeZoneSetDefault(tz unsafe.Pointer) {
	_CFTimeZoneSetDefault(tz)
}

// Determines if the given URL conforms to RFC 1808 and therefore can be decomposed.

// Determines if the given URL conforms to RFC 1808 and therefore can be decomposed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCanBeDecomposed(_:)
func CFURLCanBeDecomposed(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCanBeDecomposed(anURL)
}

// Removes all cached resource values and temporary resource values from the URL object.
//
// Added in macOS 10.6.
// Removes all cached resource values and temporary resource values from the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLClearResourcePropertyCache(_:)
func CFURLClearResourcePropertyCache(url unsafe.Pointer) {
	_CFURLClearResourcePropertyCache(url)
}

// Removes the cached resource value identified by a given key from the URL object.
//
// Added in macOS 10.6.
// Removes the cached resource value identified by a given key from the URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLClearResourcePropertyCacheForKey(_:_:)
func CFURLClearResourcePropertyCacheForKey(url unsafe.Pointer, key unsafe.Pointer) {
	_CFURLClearResourcePropertyCacheForKey(url, key)
}

// Creates a new object by resolving the relative portion of a URL against its base.

// Creates a new object by resolving the relative portion of a URL against its base.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyAbsoluteURL(_:)
func CFURLCopyAbsoluteURL(relativeURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyAbsoluteURL(relativeURL)
}

// Returns the path portion of a given URL.

// Returns the path portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyFileSystemPath(_:_:)
func CFURLCopyFileSystemPath(anURL unsafe.Pointer, pathStyle unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyFileSystemPath(anURL, pathStyle)
}

// Returns the fragment from a given URL.

// Returns the fragment from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyFragment(_:_:)
func CFURLCopyFragment(anURL unsafe.Pointer, charactersToLeaveEscaped unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyFragment(anURL, charactersToLeaveEscaped)
}

// Returns the host name of a given URL.

// Returns the host name of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyHostName(_:)
func CFURLCopyHostName(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyHostName(anURL)
}

// Returns the last path component of a given URL.

// Returns the last path component of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyLastPathComponent(_:)
func CFURLCopyLastPathComponent(url unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyLastPathComponent(url)
}

// Returns the net location portion of a given URL.

// Returns the net location portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyNetLocation(_:)
func CFURLCopyNetLocation(anURL unsafe.Pointer) unsafe.Pointer {
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
func CFURLCopyParameterString(anURL unsafe.Pointer, charactersToLeaveEscaped unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyParameterString(anURL, charactersToLeaveEscaped)
}

// Returns the password of a given URL.

// Returns the password of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPassword(_:)
func CFURLCopyPassword(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyPassword(anURL)
}

// Returns the path portion of a given URL.

// Returns the path portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPath(_:)
func CFURLCopyPath(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyPath(anURL)
}

// Returns the path extension of a given URL.

// Returns the path extension of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyPathExtension(_:)
func CFURLCopyPathExtension(url unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyPathExtension(url)
}

// Returns the query string of a given URL.

// Returns the query string of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyQueryString(_:_:)
func CFURLCopyQueryString(anURL unsafe.Pointer, charactersToLeaveEscaped unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyQueryString(anURL, charactersToLeaveEscaped)
}

// Returns the resource values for the properties identified by specified array of keys.
//
// Added in macOS 10.6.
// Returns the resource values for the properties identified by specified array of keys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourcePropertiesForKeys(_:_:_:)
func CFURLCopyResourcePropertiesForKeys(url unsafe.Pointer, keys unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyResourcePropertiesForKeys(url, keys, error_)
}

// Returns the value of a given resource property of a given URL.
//
// Added in macOS 10.6.
// Returns the value of a given resource property of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourcePropertyForKey(_:_:_:_:)
func CFURLCopyResourcePropertyForKey(url unsafe.Pointer, key unsafe.Pointer, propertyValueTypeRefPtr unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyResourcePropertyForKey(url, key, propertyValueTypeRefPtr, error_)
}

// Returns any additional resource specifiers after the path.

// Returns any additional resource specifiers after the path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyResourceSpecifier(_:)
func CFURLCopyResourceSpecifier(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyResourceSpecifier(anURL)
}

// Returns the scheme portion of a given URL.

// Returns the scheme portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyScheme(_:)
func CFURLCopyScheme(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyScheme(anURL)
}

// Returns the path portion of a given URL.

// Returns the path portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyStrictPath(_:_:)
func CFURLCopyStrictPath(anURL unsafe.Pointer, isAbsolute unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyStrictPath(anURL, isAbsolute)
}

// Returns the user name from a given URL.

// Returns the user name from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyUserName(_:)
func CFURLCopyUserName(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyUserName(anURL)
}

// Creates a new object by resolving the relative portion of a URL, specified as bytes, against its given base URL.

// Creates a new object by resolving the relative portion of a URL, specified as bytes, against its given base URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateAbsoluteURLWithBytes(_:_:_:_:_:_:)
func CFURLCreateAbsoluteURLWithBytes(alloc unsafe.Pointer, relativeURLBytes unsafe.Pointer, length unsafe.Pointer, encoding unsafe.Pointer, baseURL unsafe.Pointer, useCompatibilityMode unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateAbsoluteURLWithBytes(alloc, relativeURLBytes, length, encoding, baseURL, useCompatibilityMode)
}

// Returns bookmark data for a URL, created with specified options and resource values.
//
// Added in macOS 10.6.
// Returns bookmark data for a URL, created with specified options and resource values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkData(_:_:_:_:_:_:)
func CFURLCreateBookmarkData(allocator unsafe.Pointer, url unsafe.Pointer, options unsafe.Pointer, resourcePropertiesToInclude unsafe.Pointer, relativeToURL unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
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
func CFURLCreateBookmarkDataFromAliasRecord(allocatorRef unsafe.Pointer, aliasRecordDataRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateBookmarkDataFromAliasRecord(allocatorRef, aliasRecordDataRef)
}

// Initializes and returns bookmark data derived from a file pointed to by a specified URL.
//
// Added in macOS 10.6.
// Initializes and returns bookmark data derived from a file pointed to by a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateBookmarkDataFromFile(_:_:_:)
func CFURLCreateBookmarkDataFromFile(allocator unsafe.Pointer, fileURL unsafe.Pointer, errorRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateBookmarkDataFromFile(allocator, fileURL, errorRef)
}

// Returns a new URL made by resolving bookmark data.
//
// Added in macOS 10.6.
// Returns a new URL made by resolving bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateByResolvingBookmarkData(_:_:_:_:_:_:_:)
func CFURLCreateByResolvingBookmarkData(allocator unsafe.Pointer, bookmark unsafe.Pointer, options unsafe.Pointer, relativeToURL unsafe.Pointer, resourcePropertiesToInclude unsafe.Pointer, isStale unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateByResolvingBookmarkData(allocator, bookmark, options, relativeToURL, resourcePropertiesToInclude, isStale, error_)
}

// Creates a copy of a given URL and appends a path component.

// Creates a copy of a given URL and appends a path component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyAppendingPathComponent(_:_:_:_:)
func CFURLCreateCopyAppendingPathComponent(allocator unsafe.Pointer, url unsafe.Pointer, pathComponent unsafe.Pointer, isDirectory unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateCopyAppendingPathComponent(allocator, url, pathComponent, isDirectory)
}

// Creates a copy of a given URL and appends a path extension.

// Creates a copy of a given URL and appends a path extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyAppendingPathExtension(_:_:_:)
func CFURLCreateCopyAppendingPathExtension(allocator unsafe.Pointer, url unsafe.Pointer, extension unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateCopyAppendingPathExtension(allocator, url, extension)
}

// Creates a copy of a given URL with the last path component deleted.

// Creates a copy of a given URL with the last path component deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyDeletingLastPathComponent(_:_:)
func CFURLCreateCopyDeletingLastPathComponent(allocator unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateCopyDeletingLastPathComponent(allocator, url)
}

// Creates a copy of a given URL with its last path extension removed.

// Creates a copy of a given URL with its last path extension removed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateCopyDeletingPathExtension(_:_:)
func CFURLCreateCopyDeletingPathExtension(allocator unsafe.Pointer, url unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateCopyDeletingPathExtension(allocator, url)
}

// Creates a object containing the content of a given URL.

// Creates a object containing the content of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateData(_:_:_:_:)
func CFURLCreateData(allocator unsafe.Pointer, url unsafe.Pointer, encoding unsafe.Pointer, escapeWhitespace unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateData(allocator, url, encoding, escapeWhitespace)
}

// Loads the data and properties referred to by a given URL.

// Loads the data and properties referred to by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateDataAndPropertiesFromResource(_:_:_:_:_:_:)
func CFURLCreateDataAndPropertiesFromResource(alloc unsafe.Pointer, url unsafe.Pointer, resourceData unsafe.Pointer, properties unsafe.Pointer, desiredProperties unsafe.Pointer, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateDataAndPropertiesFromResource(alloc, url, resourceData, properties, desiredProperties, errorCode)
}

// Returns a new file path URL that refers to the same resource as a specified URL.
//
// Added in macOS 10.6.
// Returns a new file path URL that refers to the same resource as a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFilePathURL(_:_:_:)
func CFURLCreateFilePathURL(allocator unsafe.Pointer, url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateFilePathURL(allocator, url, error_)
}

// Returns a new file reference URL that points to the same resource as a specified URL.
//
// Added in macOS 10.6.
// Returns a new file reference URL that points to the same resource as a specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFileReferenceURL(_:_:_:)
func CFURLCreateFileReferenceURL(allocator unsafe.Pointer, url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateFileReferenceURL(allocator, url, error_)
}

// Creates a URL from a given directory or file.

// Creates a URL from a given directory or file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFSRef(_:_:)
func CFURLCreateFromFSRef(allocator unsafe.Pointer, fsRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateFromFSRef(allocator, fsRef)
}

// Creates a new object for a file system entity using the native representation.

// Creates a new object for a file system entity using the native representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFileSystemRepresentation(_:_:_:_:)
func CFURLCreateFromFileSystemRepresentation(allocator unsafe.Pointer, buffer unsafe.Pointer, bufLen unsafe.Pointer, isDirectory unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateFromFileSystemRepresentation(allocator, buffer, bufLen, isDirectory)
}

// Creates a object from a native character string path relative to a base URL.

// Creates a object from a native character string path relative to a base URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateFromFileSystemRepresentationRelativeToBase(_:_:_:_:_:)
func CFURLCreateFromFileSystemRepresentationRelativeToBase(allocator unsafe.Pointer, buffer unsafe.Pointer, bufLen unsafe.Pointer, isDirectory unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateFromFileSystemRepresentationRelativeToBase(allocator, buffer, bufLen, isDirectory, baseURL)
}

// Returns a given property specified by a given URL and property string.

// Returns a given property specified by a given URL and property string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreatePropertyFromResource(_:_:_:_:)
func CFURLCreatePropertyFromResource(alloc unsafe.Pointer, url unsafe.Pointer, property unsafe.Pointer, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreatePropertyFromResource(alloc, url, property, errorCode)
}

// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// Added in macOS 10.6.
// Returns the resource values for properties identified by a specified array of keys contained in specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateResourcePropertiesForKeysFromBookmarkData(_:_:_:)
func CFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator unsafe.Pointer, resourcePropertiesToReturn unsafe.Pointer, bookmark unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator, resourcePropertiesToReturn, bookmark)
}

// Returns the value of a resource property from specified bookmark data.
//
// Added in macOS 10.6.
// Returns the value of a resource property from specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateResourcePropertyForKeyFromBookmarkData(_:_:_:)
func CFURLCreateResourcePropertyForKeyFromBookmarkData(allocator unsafe.Pointer, resourcePropertyKey unsafe.Pointer, bookmark unsafe.Pointer) unsafe.Pointer {
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
func CFURLCreateStringByAddingPercentEscapes(allocator unsafe.Pointer, originalString unsafe.Pointer, charactersToLeaveUnescaped unsafe.Pointer, legalURLCharactersToBeEscaped unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateStringByAddingPercentEscapes(allocator, originalString, charactersToLeaveUnescaped, legalURLCharactersToBeEscaped, encoding)
}

// Creates a new string by replacing any percent escape sequences with their character equivalent.

// Creates a new string by replacing any percent escape sequences with their character equivalent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateStringByReplacingPercentEscapes(_:_:_:)
func CFURLCreateStringByReplacingPercentEscapes(allocator unsafe.Pointer, originalString unsafe.Pointer, charactersToLeaveEscaped unsafe.Pointer) unsafe.Pointer {
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
func CFURLCreateStringByReplacingPercentEscapesUsingEncoding(allocator unsafe.Pointer, origString unsafe.Pointer, charsToLeaveEscaped unsafe.Pointer, encoding unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateStringByReplacingPercentEscapesUsingEncoding(allocator, origString, charsToLeaveEscaped, encoding)
}

// Creates a object using a given character bytes.

// Creates a object using a given character bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithBytes(_:_:_:_:_:)
func CFURLCreateWithBytes(allocator unsafe.Pointer, URLBytes unsafe.Pointer, length unsafe.Pointer, encoding unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateWithBytes(allocator, URLBytes, length, encoding, baseURL)
}

// Creates a object using a local file system path string.

// Creates a object using a local file system path string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithFileSystemPath(_:_:_:_:)
func CFURLCreateWithFileSystemPath(allocator unsafe.Pointer, filePath unsafe.Pointer, pathStyle unsafe.Pointer, isDirectory unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateWithFileSystemPath(allocator, filePath, pathStyle, isDirectory)
}

// Creates a object using a local file system path string relative to a base URL.

// Creates a object using a local file system path string relative to a base URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithFileSystemPathRelativeToBase(_:_:_:_:_:)
func CFURLCreateWithFileSystemPathRelativeToBase(allocator unsafe.Pointer, filePath unsafe.Pointer, pathStyle unsafe.Pointer, isDirectory unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateWithFileSystemPathRelativeToBase(allocator, filePath, pathStyle, isDirectory, baseURL)
}

// Creates a object using a given object.

// Creates a object using a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCreateWithString(_:_:_:)
func CFURLCreateWithString(allocator unsafe.Pointer, URLString unsafe.Pointer, baseURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCreateWithString(allocator, URLString, baseURL)
}

// Destroys a resource indicated by a given URL.

// Destroys a resource indicated by a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLDestroyResource(_:_:)
func CFURLDestroyResource(url unsafe.Pointer, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLDestroyResource(url, errorCode)
}

// Returns the base URL of a given URL if it exists.

// Returns the base URL of a given URL if it exists.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetBaseURL(_:)
func CFURLGetBaseURL(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetBaseURL(anURL)
}

// Returns the range of the specified component in the bytes of a URL.

// Returns the range of the specified component in the bytes of a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetByteRangeForComponent(_:_:_:)
func CFURLGetByteRangeForComponent(url unsafe.Pointer, component unsafe.Pointer, rangeIncludingSeparators unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetByteRangeForComponent(url, component, rangeIncludingSeparators)
}

// Returns by reference the byte representation of a URL object.

// Returns by reference the byte representation of a URL object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetBytes(_:_:_:)
func CFURLGetBytes(url unsafe.Pointer, buffer unsafe.Pointer, bufferLength unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetBytes(url, buffer, bufferLength)
}

// Converts a given URL to a file or directory object.

// Converts a given URL to a file or directory object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetFSRef(_:_:)
func CFURLGetFSRef(url unsafe.Pointer, fsRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetFSRef(url, fsRef)
}

// Fills a buffer with the file system’s native string representation of a given URL’s path.

// Fills a buffer with the file system’s native string representation of a given URL’s path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetFileSystemRepresentation(_:_:_:_:)
func CFURLGetFileSystemRepresentation(url unsafe.Pointer, resolveAgainstBase unsafe.Pointer, buffer unsafe.Pointer, maxBufLen unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetFileSystemRepresentation(url, resolveAgainstBase, buffer, maxBufLen)
}

// Returns the port number from a given URL.

// Returns the port number from a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetPortNumber(_:)
func CFURLGetPortNumber(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetPortNumber(anURL)
}

// Returns the URL as a object.

// Returns the URL as a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetString(_:)
func CFURLGetString(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLGetString(anURL)
}

// Returns the type identifier for the opaque type.

// Returns the type identifier for the opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLGetTypeID()
func CFURLGetTypeID() unsafe.Pointer {
	return _CFURLGetTypeID()
}

// Determines if a given URL’s path represents a directory.

// Determines if a given URL’s path represents a directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLHasDirectoryPath(_:)
func CFURLHasDirectoryPath(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLHasDirectoryPath(anURL)
}

// Returns whether the resource pointed to by a file URL can be reached.
//
// Added in macOS 10.6.
// Returns whether the resource pointed to by a file URL can be reached.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLResourceIsReachable(_:_:)
func CFURLResourceIsReachable(url unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLResourceIsReachable(url, error_)
}

// Sets the URL’s resource properties for a given set of keys to a given set of values.
//
// Added in macOS 10.6.
// Sets the URL’s resource properties for a given set of keys to a given set of values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLSetResourcePropertiesForKeys(_:_:_:)
func CFURLSetResourcePropertiesForKeys(url unsafe.Pointer, keyedPropertyValues unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLSetResourcePropertiesForKeys(url, keyedPropertyValues, error_)
}

// Sets the URL’s resource property for a given key to a given value.
//
// Added in macOS 10.6.
// Sets the URL’s resource property for a given key to a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLSetResourcePropertyForKey(_:_:_:_:)
func CFURLSetResourcePropertyForKey(url unsafe.Pointer, key unsafe.Pointer, propertyValue unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	return _CFURLSetResourcePropertyForKey(url, key, propertyValue, error_)
}

// Sets a temporary resource value on the URL.
//
// Added in macOS 10.6.
// Sets a temporary resource value on the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLSetTemporaryResourcePropertyForKey(_:_:_:)
func CFURLSetTemporaryResourcePropertyForKey(url unsafe.Pointer, key unsafe.Pointer, propertyValue unsafe.Pointer) {
	_CFURLSetTemporaryResourcePropertyForKey(url, key, propertyValue)
}

// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// Added in macOS 10.7.
// In an app that has adopted App Sandbox, makes the resource pointed to by a security-scoped URL available to the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLStartAccessingSecurityScopedResource(_:)
func CFURLStartAccessingSecurityScopedResource(url unsafe.Pointer) unsafe.Pointer {
	return _CFURLStartAccessingSecurityScopedResource(url)
}

// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// Added in macOS 10.7.
// In an app that adopts App Sandbox, revokes access to the resource pointed to by a security-scoped URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLStopAccessingSecurityScopedResource(_:)
func CFURLStopAccessingSecurityScopedResource(url unsafe.Pointer) {
	_CFURLStopAccessingSecurityScopedResource(url)
}

// Creates an alias file on disk at a specified location with specified bookmark data.
//
// Added in macOS 10.6.
// Creates an alias file on disk at a specified location with specified bookmark data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLWriteBookmarkDataToFile(_:_:_:_:)
func CFURLWriteBookmarkDataToFile(bookmarkRef unsafe.Pointer, fileURL unsafe.Pointer, options unsafe.Pointer, errorRef unsafe.Pointer) unsafe.Pointer {
	return _CFURLWriteBookmarkDataToFile(bookmarkRef, fileURL, options, errorRef)
}

// Writes the given data and properties to a given URL.

// Writes the given data and properties to a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLWriteDataAndPropertiesToResource(_:_:_:_:)
func CFURLWriteDataAndPropertiesToResource(url unsafe.Pointer, dataToWrite unsafe.Pointer, propertiesToWrite unsafe.Pointer, errorCode unsafe.Pointer) unsafe.Pointer {
	return _CFURLWriteDataAndPropertiesToResource(url, dataToWrite, propertiesToWrite, errorCode)
}

// Creates a Universally Unique Identifier (UUID) object.

// Creates a Universally Unique Identifier (UUID) object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreate(_:)
func CFUUIDCreate(alloc unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDCreate(alloc)
}

// Creates a CFUUID object for a specified string.

// Creates a CFUUID object for a specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateFromString(_:_:)
func CFUUIDCreateFromString(alloc unsafe.Pointer, uuidStr unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDCreateFromString(alloc, uuidStr)
}

// Creates a CFUUID object from raw UUID bytes.

// Creates a CFUUID object from raw UUID bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateFromUUIDBytes(_:_:)
func CFUUIDCreateFromUUIDBytes(alloc unsafe.Pointer, bytes unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDCreateFromUUIDBytes(alloc, bytes)
}

// Returns the string representation of a specified CFUUID object.

// Returns the string representation of a specified CFUUID object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateString(_:_:)
func CFUUIDCreateString(alloc unsafe.Pointer, uuid unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDCreateString(alloc, uuid)
}

// Creates a CFUUID object from raw UUID bytes.

// Creates a CFUUID object from raw UUID bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDCreateWithBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CFUUIDCreateWithBytes(alloc unsafe.Pointer, byte0 unsafe.Pointer, byte1 unsafe.Pointer, byte2 unsafe.Pointer, byte3 unsafe.Pointer, byte4 unsafe.Pointer, byte5 unsafe.Pointer, byte6 unsafe.Pointer, byte7 unsafe.Pointer, byte8 unsafe.Pointer, byte9 unsafe.Pointer, byte10 unsafe.Pointer, byte11 unsafe.Pointer, byte12 unsafe.Pointer, byte13 unsafe.Pointer, byte14 unsafe.Pointer, byte15 unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDCreateWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15)
}

// Returns a CFUUID object from raw UUID bytes.

// Returns a CFUUID object from raw UUID bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetConstantUUIDWithBytes(_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:_:)
func CFUUIDGetConstantUUIDWithBytes(alloc unsafe.Pointer, byte0 unsafe.Pointer, byte1 unsafe.Pointer, byte2 unsafe.Pointer, byte3 unsafe.Pointer, byte4 unsafe.Pointer, byte5 unsafe.Pointer, byte6 unsafe.Pointer, byte7 unsafe.Pointer, byte8 unsafe.Pointer, byte9 unsafe.Pointer, byte10 unsafe.Pointer, byte11 unsafe.Pointer, byte12 unsafe.Pointer, byte13 unsafe.Pointer, byte14 unsafe.Pointer, byte15 unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDGetConstantUUIDWithBytes(alloc, byte0, byte1, byte2, byte3, byte4, byte5, byte6, byte7, byte8, byte9, byte10, byte11, byte12, byte13, byte14, byte15)
}

// Returns the type identifier for all CFUUID objects.

// Returns the type identifier for all CFUUID objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetTypeID()
func CFUUIDGetTypeID() unsafe.Pointer {
	return _CFUUIDGetTypeID()
}

// Returns the value of a UUID object as raw bytes.

// Returns the value of a UUID object as raw bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUIDGetUUIDBytes(_:)
func CFUUIDGetUUIDBytes(uuid unsafe.Pointer) unsafe.Pointer {
	return _CFUUIDGetUUIDBytes(uuid)
}

// Returns whether a writable stream can accept new data without blocking.

// Returns whether a writable stream can accept new data without blocking.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCanAcceptBytes(_:)
func CFWriteStreamCanAcceptBytes(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCanAcceptBytes(stream)
}

// Closes a writable stream.

// Closes a writable stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamClose(_:)
func CFWriteStreamClose(stream unsafe.Pointer) {
	_CFWriteStreamClose(stream)
}

// Returns the error associated with a stream.
//
// Added in macOS 10.5.
// Returns the error associated with a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyError(_:)
func CFWriteStreamCopyError(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCopyError(stream)
}

// Returns the value of a property for a stream.

// Returns the value of a property for a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCopyProperty(_:_:)
func CFWriteStreamCopyProperty(stream unsafe.Pointer, propertyName unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCopyProperty(stream, propertyName)
}

// Creates a writable stream for a growable block of memory.

// Creates a writable stream for a growable block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithAllocatedBuffers(_:_:)
func CFWriteStreamCreateWithAllocatedBuffers(alloc unsafe.Pointer, bufferAllocator unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCreateWithAllocatedBuffers(alloc, bufferAllocator)
}

// Creates a writable stream for a fixed-size block of memory.

// Creates a writable stream for a fixed-size block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithBuffer(_:_:_:)
func CFWriteStreamCreateWithBuffer(alloc unsafe.Pointer, buffer unsafe.Pointer, bufferCapacity unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCreateWithBuffer(alloc, buffer, bufferCapacity)
}

// Creates a writable stream for a file.

// Creates a writable stream for a file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamCreateWithFile(_:_:)
func CFWriteStreamCreateWithFile(alloc unsafe.Pointer, fileURL unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamCreateWithFile(alloc, fileURL)
}

// Returns the error status of a stream.

// Returns the error status of a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetError(_:)
func CFWriteStreamGetError(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamGetError(stream)
}

// Returns the current state of a stream.

// Returns the current state of a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetStatus(_:)
func CFWriteStreamGetStatus(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamGetStatus(stream)
}

// Returns the type identifier of all CFWriteStream objects.

// Returns the type identifier of all CFWriteStream objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamGetTypeID()
func CFWriteStreamGetTypeID() unsafe.Pointer {
	return _CFWriteStreamGetTypeID()
}

// Opens a stream for writing.

// Opens a stream for writing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamOpen(_:)
func CFWriteStreamOpen(stream unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamOpen(stream)
}

// Schedules a stream into a run loop.

// Schedules a stream into a run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamScheduleWithRunLoop(_:_:_:)
func CFWriteStreamScheduleWithRunLoop(stream unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_CFWriteStreamScheduleWithRunLoop(stream, runLoop, runLoopMode)
}

// Assigns a client to a stream, which receives callbacks when certain events occur.

// Assigns a client to a stream, which receives callbacks when certain events occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetClient(_:_:_:_:)
func CFWriteStreamSetClient(stream unsafe.Pointer, streamEvents unsafe.Pointer, clientCB unsafe.Pointer, clientContext unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamSetClient(stream, streamEvents, clientCB, clientContext)
}

// Sets the value of a property for a stream.

// Sets the value of a property for a stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamSetProperty(_:_:_:)
func CFWriteStreamSetProperty(stream unsafe.Pointer, propertyName unsafe.Pointer, propertyValue unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamSetProperty(stream, propertyName, propertyValue)
}

// Removes a stream from a particular run loop.

// Removes a stream from a particular run loop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamUnscheduleFromRunLoop(_:_:_:)
func CFWriteStreamUnscheduleFromRunLoop(stream unsafe.Pointer, runLoop unsafe.Pointer, runLoopMode unsafe.Pointer) {
	_CFWriteStreamUnscheduleFromRunLoop(stream, runLoop, runLoopMode)
}

// Writes data to a writable stream.

// Writes data to a writable stream.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamWrite(_:_:_:)
func CFWriteStreamWrite(stream unsafe.Pointer, buffer unsafe.Pointer, bufferLength unsafe.Pointer) unsafe.Pointer {
	return _CFWriteStreamWrite(stream, buffer, bufferLength)
}

// inset is a Foundation function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRect/inset(by:)
func inset(insets unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _inset(insets, p1)
}

// Tests whether an app can accept (open) an item for a URL.
//
// Added in macOS 10.0.
// Tests whether an app can accept (open) an item for a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1441854-lscanurlaccepturl
func LSCanURLAcceptURL(inItemURL unsafe.Pointer, inTargetURL unsafe.Pointer, inRoleMask unsafe.Pointer, inFlags unsafe.Pointer, outAcceptsItem unsafe.Pointer, p5 unsafe.Pointer) unsafe.Pointer {
	return _LSCanURLAcceptURL(inItemURL, inTargetURL, inRoleMask, inFlags, outAcceptsItem, p5)
}

// Opens one or more items for a URL in the preferred apps or a designated app.
//
// Added in macOS 10.0.
// Opens one or more items for a URL in the preferred apps or a designated app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1441986-lsopenfromurlspec
func LSOpenFromURLSpec(inLaunchSpec unsafe.Pointer, outLaunchedURL unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _LSOpenFromURLSpec(inLaunchSpec, outLaunchedURL, p2)
}

// Opens an item for a URL in the default manner in its preferred app.
//
// Added in macOS 10.0.
// Opens an item for a URL in the default manner in its preferred app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1442850-lsopencfurlref
func LSOpenCFURLRef(inURL unsafe.Pointer, outLaunchedURL unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _LSOpenCFURLRef(inURL, outLaunchedURL, p2)
}

// Sets the user’s preferred default handler for the specified content type in the specified roles.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.4.
// Sets the user’s preferred default handler for the specified content type in the specified roles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1444955-lssetdefaultrolehandlerforconten
func LSSetDefaultRoleHandlerForContentType(inContentType unsafe.Pointer, inRole unsafe.Pointer, inHandlerBundleID unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _LSSetDefaultRoleHandlerForContentType(inContentType, inRole, inHandlerBundleID, p3)
}

// Locates all known apps suitable for opening an item for the specified URL.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.3.
// Locates all known apps suitable for opening an item for the specified URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1445148-lscopyapplicationurlsforurl
func LSCopyApplicationURLsForURL(inURL unsafe.Pointer, inRoleMask unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _LSCopyApplicationURLsForURL(inURL, inRoleMask, p2)
}

// Registers an app, using a URL, in the Launch Services database.
//
// Added in macOS 10.3.
// Registers an app, using a URL, in the Launch Services database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1446350-lsregisterurl
func LSRegisterURL(inURL unsafe.Pointer, inUpdate unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _LSRegisterURL(inURL, inUpdate, p2)
}

// Returns the app that opens a content type.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.10.
// Returns the app that opens a content type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1447734-lscopydefaultapplicationurlforco
func LSCopyDefaultApplicationURLForContentType(inContentType unsafe.Pointer, inRoleMask unsafe.Pointer, outError unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _LSCopyDefaultApplicationURLForContentType(inContentType, inRoleMask, outError, p3)
}

// Sets the user’s preferred default handler for the specified URL scheme.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.4.
// Sets the user’s preferred default handler for the specified URL scheme.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1447760-lssetdefaulthandlerforurlscheme
func LSSetDefaultHandlerForURLScheme(inURLScheme unsafe.Pointer, inHandlerBundleID unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _LSSetDefaultHandlerForURLScheme(inURLScheme, inHandlerBundleID, p2)
}

// Locates an array of bundle identifiers for apps capable of handling a specified content type with the specified roles.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.4.
// Locates an array of bundle identifiers for apps capable of handling a specified content type with the specified roles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1448020-lscopyallrolehandlersforcontentt
func LSCopyAllRoleHandlersForContentType(inContentType unsafe.Pointer, inRole unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _LSCopyAllRoleHandlersForContentType(inContentType, inRole, p2)
}

// Returns the app that opens an item.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.10.
// Returns the app that opens an item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1448824-lscopydefaultapplicationurlforur
func LSCopyDefaultApplicationURLForURL(inURL unsafe.Pointer, inRoleMask unsafe.Pointer, outError unsafe.Pointer, p3 unsafe.Pointer) unsafe.Pointer {
	return _LSCopyDefaultApplicationURLForURL(inURL, inRoleMask, outError, p3)
}

// Locates all URLs for apps that correspond to the specified bundle identifier.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.10.
// Locates all URLs for apps that correspond to the specified bundle identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1449290-lscopyapplicationurlsforbundleid
func LSCopyApplicationURLsForBundleIdentifier(inBundleIdentifier unsafe.Pointer, outError unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _LSCopyApplicationURLsForBundleIdentifier(inBundleIdentifier, outError, p2)
}

// Returns the bundle identifier of the user’s preferred default handler for the specified content type with the specified role.
//
// Deprecated: This function was deprecated in macOS 12.0.
//
// Added in macOS 10.4.
// Returns the bundle identifier of the user’s preferred default handler for the specified content type with the specified role.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coreservices/1449868-lscopydefaultrolehandlerforconte
func LSCopyDefaultRoleHandlerForContentType(inContentType unsafe.Pointer, inRole unsafe.Pointer, p2 unsafe.Pointer) unsafe.Pointer {
	return _LSCopyDefaultRoleHandlerForContentType(inContentType, inRole, p2)
}

// floorf is a Foundation function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/1557176-floorf
func floorf(p0 float32) float32 {
	return _floorf(p0)
}

// ceill is a Foundation function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/1557207-ceill
func ceill(p0 unsafe.Pointer) unsafe.Pointer {
	return _ceill(p0)
}

// ceilf is a Foundation function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/1557263-ceilf
func ceilf(p0 float32) float32 {
	return _ceilf(p0)
}

// ceil is a Foundation function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/1557272-ceil
func ceil(p0 float64) float64 {
	return _ceil(p0)
}

// floorl is a Foundation function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/1557330-floorl
func floorl(p0 unsafe.Pointer) unsafe.Pointer {
	return _floorl(p0)
}

// floor is a Foundation function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/1557338-floor
func floor(p0 float64) float64 {
	return _floor(p0)
}

// Determines the amount of memory available to the current app.

// Determines the amount of memory available to the current app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/os/3191911-os_proc_available_memory
func os_proc_available_memory() uintptr {
	return _os_proc_available_memory()
}

// Returns the current network information for a given network interface.

// Returns the current network information for a given network interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/systemconfiguration/1614126-cncopycurrentnetworkinfo
func CNCopyCurrentNetworkInfo(p0 unsafe.Pointer) unsafe.Pointer {
	return _CNCopyCurrentNetworkInfo(p0)
}

// Returns a string formatted to contain the data from an edge insets structure.

// Returns a string formatted to contain the data from an edge insets structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSStringFromUIEdgeInsets
func NSStringFromUIEdgeInsets(insets unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromUIEdgeInsets(insets)
}

// Converts a UIKit text alignment constant value to the matching constant value that Core Text uses.

// Converts a UIKit text alignment constant value to the matching constant value that Core Text uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/NSTextAlignmentToCTTextAlignment
func NSTextAlignmentToCTTextAlignment(nsTextAlignment unsafe.Pointer) unsafe.Pointer {
	return _NSTextAlignmentToCTTextAlignment(nsTextAlignment)
}

// A Boolean value that indicates whether the Button Shapes setting is in an enabled state.

// A Boolean value that indicates whether the Button Shapes setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/buttonShapesEnabled
func UIAccessibilityButtonShapesEnabled() bool {
	return _UIAccessibilityButtonShapesEnabled()
}

// Enables or disables the specified accessibility features while using Guided Access.

// Enables or disables the specified accessibility features while using Guided Access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/configureForGuidedAccess(features:enabled:completionHandler:)
func UIGuidedAccessConfigureAccessibilityFeatures(features unsafe.Pointer, enabled bool) {
	_UIGuidedAccessConfigureAccessibilityFeatures(features, enabled)
}

// Converts the specified path object to screen coordinates and returns a new path object with the results.

// Converts the specified path object to screen coordinates and returns a new path object with the results.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/convertToScreenCoordinates(_:in:)-6dx4a
func UIAccessibilityConvertPathToScreenCoordinates(path unsafe.Pointer, view unsafe.Pointer) unsafe.Pointer {
	return _UIAccessibilityConvertPathToScreenCoordinates(path, view)
}

// Converts the specified rectangle from view coordinates to screen coordinates.

// Converts the specified rectangle from view coordinates to screen coordinates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/convertToScreenCoordinates(_:in:)-9ziiu
func UIAccessibilityConvertFrameToScreenCoordinates(rect coregraphics.CGRect, view unsafe.Pointer) coregraphics.CGRect {
	return _UIAccessibilityConvertFrameToScreenCoordinates(rect, view)
}

// Returns the accessibility element that’s currently in focus by the specified assistive app.

// Returns the accessibility element that’s currently in focus by the specified assistive app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/focusedElement(using:)
func UIAccessibilityFocusedElement(assistiveTechnologyIdentifier unsafe.Pointer) unsafe.Pointer {
	return _UIAccessibilityFocusedElement(assistiveTechnologyIdentifier)
}

// The current pairing status of Made for iPhone hearing devices.

// The current pairing status of Made for iPhone hearing devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/hearingDevicePairedEar
func UIAccessibilityHearingDevicePairedEar() unsafe.Pointer {
	return _UIAccessibilityHearingDevicePairedEar()
}

// A Boolean value that indicates whether AssistiveTouch is in an enabled state.

// A Boolean value that indicates whether AssistiveTouch is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isAssistiveTouchRunning
func UIAccessibilityIsAssistiveTouchRunning() bool {
	return _UIAccessibilityIsAssistiveTouchRunning()
}

// A Boolean value that indicates whether the Bold Text setting is in an enabled state.

// A Boolean value that indicates whether the Bold Text setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isBoldTextEnabled
func UIAccessibilityIsBoldTextEnabled() bool {
	return _UIAccessibilityIsBoldTextEnabled()
}

// A Boolean value that indicates whether the Closed Captions + SDH setting is in an enabled state.

// A Boolean value that indicates whether the Closed Captions + SDH setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isClosedCaptioningEnabled
func UIAccessibilityIsClosedCaptioningEnabled() bool {
	return _UIAccessibilityIsClosedCaptioningEnabled()
}

// A Boolean value that indicates whether the Increase Contrast setting is in an enabled state.

// A Boolean value that indicates whether the Increase Contrast setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isDarkerSystemColorsEnabled
func UIAccessibilityDarkerSystemColorsEnabled() bool {
	return _UIAccessibilityDarkerSystemColorsEnabled()
}

// A Boolean value that indicates whether the Color Filters and the Grayscale settings are in an enabled state.

// A Boolean value that indicates whether the Color Filters and the Grayscale settings are in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isGrayscaleEnabled
func UIAccessibilityIsGrayscaleEnabled() bool {
	return _UIAccessibilityIsGrayscaleEnabled()
}

// A Boolean value that indicates whether the Guided Access setting is in an enabled state.

// A Boolean value that indicates whether the Guided Access setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isGuidedAccessEnabled
func UIAccessibilityIsGuidedAccessEnabled() bool {
	return _UIAccessibilityIsGuidedAccessEnabled()
}

// A Boolean value that indicates whether the Classic Invert setting is in an enabled state.

// A Boolean value that indicates whether the Classic Invert setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isInvertColorsEnabled
func UIAccessibilityIsInvertColorsEnabled() bool {
	return _UIAccessibilityIsInvertColorsEnabled()
}

// A Boolean value that indicates whether the Mono Audio setting is in an enabled state.

// A Boolean value that indicates whether the Mono Audio setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isMonoAudioEnabled
func UIAccessibilityIsMonoAudioEnabled() bool {
	return _UIAccessibilityIsMonoAudioEnabled()
}

// A Boolean value that indicates whether the On/Off Labels setting is in an enabled state.

// A Boolean value that indicates whether the On/Off Labels setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isOnOffSwitchLabelsEnabled
func UIAccessibilityIsOnOffSwitchLabelsEnabled() bool {
	return _UIAccessibilityIsOnOffSwitchLabelsEnabled()
}

// A Boolean value that indicates whether the Reduce Motion setting is in an enabled state.

// A Boolean value that indicates whether the Reduce Motion setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isReduceMotionEnabled
func UIAccessibilityIsReduceMotionEnabled() bool {
	return _UIAccessibilityIsReduceMotionEnabled()
}

// A Boolean value that indicates whether the Reduce Transparency setting is in an enabled state.

// A Boolean value that indicates whether the Reduce Transparency setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isReduceTransparencyEnabled
func UIAccessibilityIsReduceTransparencyEnabled() bool {
	return _UIAccessibilityIsReduceTransparencyEnabled()
}

// A Boolean value that indicates whether the Shake to Undo setting is in an enabled state.

// A Boolean value that indicates whether the Shake to Undo setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isShakeToUndoEnabled
func UIAccessibilityIsShakeToUndoEnabled() bool {
	return _UIAccessibilityIsShakeToUndoEnabled()
}

// A Boolean value that indicates whether the Speak Screen setting is in an enabled state.

// A Boolean value that indicates whether the Speak Screen setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isSpeakScreenEnabled
func UIAccessibilityIsSpeakScreenEnabled() bool {
	return _UIAccessibilityIsSpeakScreenEnabled()
}

// A Boolean value that indicates whether the Speak Selection setting is in an enabled state.

// A Boolean value that indicates whether the Speak Selection setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isSpeakSelectionEnabled
func UIAccessibilityIsSpeakSelectionEnabled() bool {
	return _UIAccessibilityIsSpeakSelectionEnabled()
}

// A Boolean value that indicates whether the Switch Control setting is in an enabled state.

// A Boolean value that indicates whether the Switch Control setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isSwitchControlRunning
func UIAccessibilityIsSwitchControlRunning() bool {
	return _UIAccessibilityIsSwitchControlRunning()
}

// A Boolean value that indicates whether the Auto-Play Video Previews setting is in an enabled state.

// A Boolean value that indicates whether the Auto-Play Video Previews setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isVideoAutoplayEnabled
func UIAccessibilityIsVideoAutoplayEnabled() bool {
	return _UIAccessibilityIsVideoAutoplayEnabled()
}

// A Boolean value that indicates whether VoiceOver is in an enabled state.

// A Boolean value that indicates whether VoiceOver is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/isVoiceOverRunning
func UIAccessibilityIsVoiceOverRunning() bool {
	return _UIAccessibilityIsVoiceOverRunning()
}

// Posts a notification to assistive apps.

// Posts a notification to assistive apps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/post(notification:argument:)
func UIAccessibilityPostNotification(notification unsafe.Pointer, argument unsafe.Pointer) {
	_UIAccessibilityPostNotification(notification, argument)
}

// A Boolean value that indicates whether the Reduce Motion and the Prefer Cross-Fade Transitions settings are in an enabled state.

// A Boolean value that indicates whether the Reduce Motion and the Prefer Cross-Fade Transitions settings are in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/prefersCrossFadeTransitions
func UIAccessibilityPrefersCrossFadeTransitions() bool {
	return _UIAccessibilityPrefersCrossFadeTransitions()
}

// Warns users that app-specific gestures conflict with the system-defined Zoom accessibility gestures.

// Warns users that app-specific gestures conflict with the system-defined Zoom accessibility gestures.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/registerGestureConflictWithZoom()
func UIAccessibilityRegisterGestureConflictWithZoom() {
	_UIAccessibilityRegisterGestureConflictWithZoom()
}

// Transitions the app to or from Single App mode asynchronously.

// Transitions the app to or from Single App mode asynchronously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/requestGuidedAccessSession(enabled:completionHandler:)
func UIAccessibilityRequestGuidedAccessSession(enable bool) {
	_UIAccessibilityRequestGuidedAccessSession(enable)
}

// A Boolean value that indicates whether the Differentiate Without Color setting is in an enabled state.

// A Boolean value that indicates whether the Differentiate Without Color setting is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/shouldDifferentiateWithoutColor
func UIAccessibilityShouldDifferentiateWithoutColor() bool {
	return _UIAccessibilityShouldDifferentiateWithoutColor()
}

// Notifies the system when the app’s focus changes to a new location.

// Notifies the system when the app’s focus changes to a new location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/zoomFocusChanged(zoomType:toFrame:in:)
func UIAccessibilityZoomFocusChanged(type_ unsafe.Pointer, frame coregraphics.CGRect, view unsafe.Pointer) {
	_UIAccessibilityZoomFocusChanged(type_, frame, view)
}

// Creates the application object and the application delegate and sets up the event cycle.

// Creates the application object and the application delegate and sets up the event cycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplicationMain(_:_:_:_:)-1yub7
func UIApplicationMain(argc int, argv unsafe.Pointer, principalClassName unsafe.Pointer, delegateClassName unsafe.Pointer) int {
	return _UIApplicationMain(argc, argv, principalClassName, delegateClassName)
}

// Returns a UIKit edge insets structure based on the data in the specified string.

// Returns a UIKit edge insets structure based on the data in the specified string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIEdgeInsetsFromString
func UIEdgeInsetsFromString(string_ unsafe.Pointer) unsafe.Pointer {
	return _UIEdgeInsetsFromString(string_)
}

// Provides the corresponding symbol weight for this font weight.

// Provides the corresponding symbol weight for this font weight.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIFont/Weight/symbolWeight()
func UIImageSymbolWeightForFontWeight(fontWeight unsafe.Pointer) unsafe.Pointer {
	return _UIImageSymbolWeightForFontWeight(fontWeight)
}



