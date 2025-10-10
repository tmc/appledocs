// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

// CoreFoundation Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (748 total):

// CFRelease(CFTypeRef  cf)
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFRetain(CFTypeRef  cf) CFTypeRef
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFAbsoluteTimeAddGregorianUnits(at CFAbsoluteTime, tz ,  CFTimeZoneRef, units ,  CFGregorianUnits, ) CFAbsoluteTime
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.4+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.


// CFAbsoluteTimeGetDayOfWeek(at CFAbsoluteTime, tz ,  CFTimeZoneRef, ) SInt32
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.4+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFAbsoluteTimeGetDayOfYear(at CFAbsoluteTime, tz ,  CFTimeZoneRef, ) SInt32
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.4+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFAbsoluteTimeGetDifferenceAsGregorianUnits(at1 CFAbsoluteTime, at2 ,  CFAbsoluteTime, tz ,  CFTimeZoneRef, unitFlags ,  CFOptionFlags, ) CFGregorianUnits
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.4+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.


// CFAbsoluteTimeGetGregorianDate(at CFAbsoluteTime, tz ,  CFTimeZoneRef, ) CFGregorianDate
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.4+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFAbsoluteTimeGetWeekOfYear(at CFAbsoluteTime, tz ,  CFTimeZoneRef, ) SInt32
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.4+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFAllocatorAllocate(allocator CFAllocatorRef, size ,  CFIndex, hint ,  CFOptionFlags, ) void  *


// CFAllocatorAllocateBytes(allocator CFAllocatorRef, size ,  CFIndex, hint ,  CFOptionFlags, ) void  *
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+

// CFAllocatorAllocateTyped(allocator CFAllocatorRef, size ,  CFIndex, descriptor ,  CFAllocatorTypeID, hint ,  CFOptionFlags, ) void  *
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+

// CFAllocatorCreate(allocator CFAllocatorRef, context ,  CFAllocatorContext  *, ) CFAllocatorRef


// CFAllocatorCreateWithZone(allocator CFAllocatorRef, zone ,  struct _malloc_zone_t  *, ) CFAllocatorRef

// CFAllocatorDeallocate(allocator CFAllocatorRef, ptr ,  void  *, )

// CFAllocatorGetContext(allocator CFAllocatorRef, context ,  CFAllocatorContext  *, )


// CFAllocatorGetPreferredSizeForSize(allocator CFAllocatorRef, size ,  CFIndex, hint ,  CFOptionFlags, ) CFIndex

// CFAllocatorReallocate(allocator CFAllocatorRef, ptr ,  void  *, newsize ,  CFIndex, hint ,  CFOptionFlags, ) void  *

// CFAllocatorReallocateBytes(allocator CFAllocatorRef, ptr ,  void  *, newsize ,  CFIndex, hint ,  CFOptionFlags, ) void  *
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+


// CFAllocatorReallocateTyped(allocator CFAllocatorRef, ptr ,  void  *, newsize ,  CFIndex, descriptor ,  CFAllocatorTypeID, hint ,  CFOptionFlags, ) void  *
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+

// CFAllocatorSetDefault(allocator CFAllocatorRef, )

// CFArrayAppendArray(theArray CFMutableArrayRef, otherArray ,  CFArrayRef, otherRange ,  CFRange, )


// CFArrayAppendValue(theArray CFMutableArrayRef, value ,  const void  *, )

// CFArrayApplyFunction(theArray CFArrayRef, range ,  CFRange, applier ,  CFArrayApplierFunction, context ,  void  *, )

// CFArrayBSearchValues(theArray CFArrayRef, range ,  CFRange, value ,  const void  *, comparator ,  CFComparatorFunction, context ,  void  *, ) CFIndex


// CFArrayContainsValue(theArray CFArrayRef, range ,  CFRange, value ,  const void  *, ) Boolean

// CFArrayCreate(allocator CFAllocatorRef, values ,  const void  * *, numValues ,  CFIndex, callBacks ,  const CFArrayCallBacks  *, ) CFArrayRef

// CFArrayCreateCopy(allocator CFAllocatorRef, theArray ,  CFArrayRef, ) CFArrayRef


// CFArrayCreateMutable(allocator CFAllocatorRef, capacity ,  CFIndex, callBacks ,  const CFArrayCallBacks  *, ) CFMutableArrayRef

// CFArrayCreateMutableCopy(allocator CFAllocatorRef, capacity ,  CFIndex, theArray ,  CFArrayRef, ) CFMutableArrayRef

// CFArrayExchangeValuesAtIndices(theArray CFMutableArrayRef, idx1 ,  CFIndex, idx2 ,  CFIndex, )


// CFArrayGetCount(theArray CFArrayRef, ) CFIndex

// CFArrayGetCountOfValue(theArray CFArrayRef, range ,  CFRange, value ,  const void  *, ) CFIndex

// CFArrayGetFirstIndexOfValue(theArray CFArrayRef, range ,  CFRange, value ,  const void  *, ) CFIndex


// CFArrayGetLastIndexOfValue(theArray CFArrayRef, range ,  CFRange, value ,  const void  *, ) CFIndex

// CFArrayGetValueAtIndex(theArray CFArrayRef, idx ,  CFIndex, ) const void  *

// CFArrayGetValues(theArray CFArrayRef, range ,  CFRange, values ,  const void  * *, )


// CFArrayInsertValueAtIndex(theArray CFMutableArrayRef, idx ,  CFIndex, value ,  const void  *, )

// CFArrayRemoveAllValues(theArray CFMutableArrayRef, )

// CFArrayRemoveValueAtIndex(theArray CFMutableArrayRef, idx ,  CFIndex, )


// CFArrayReplaceValues(theArray CFMutableArrayRef, range ,  CFRange, newValues ,  const void  * *, newCount ,  CFIndex, )

// CFArraySetValueAtIndex(theArray CFMutableArrayRef, idx ,  CFIndex, value ,  const void  *, )

// CFArraySortValues(theArray CFMutableArrayRef, range ,  CFRange, comparator ,  CFComparatorFunction, context ,  void  *, )


// CFAttributedStringBeginEditing(aStr CFMutableAttributedStringRef, )

// CFAttributedStringCreate(alloc CFAllocatorRef, str ,  CFStringRef, attributes ,  CFDictionaryRef, ) CFAttributedStringRef

// CFAttributedStringCreateCopy(alloc CFAllocatorRef, aStr ,  CFAttributedStringRef, ) CFAttributedStringRef


// CFAttributedStringCreateMutable(alloc CFAllocatorRef, maxLength ,  CFIndex, ) CFMutableAttributedStringRef

// CFAttributedStringCreateMutableCopy(alloc CFAllocatorRef, maxLength ,  CFIndex, aStr ,  CFAttributedStringRef, ) CFMutableAttributedStringRef

// CFAttributedStringCreateWithSubstring(alloc CFAllocatorRef, aStr ,  CFAttributedStringRef, range ,  CFRange, ) CFAttributedStringRef


// CFAttributedStringEndEditing(aStr CFMutableAttributedStringRef, )

// CFAttributedStringGetAttribute(aStr CFAttributedStringRef, loc ,  CFIndex, attrName ,  CFStringRef, effectiveRange ,  CFRange  *, ) CFTypeRef

// CFAttributedStringGetAttributeAndLongestEffectiveRange(aStr CFAttributedStringRef, loc ,  CFIndex, attrName ,  CFStringRef, inRange ,  CFRange, longestEffectiveRange ,  CFRange  *, ) CFTypeRef


// CFAttributedStringGetAttributes(aStr CFAttributedStringRef, loc ,  CFIndex, effectiveRange ,  CFRange  *, ) CFDictionaryRef

// CFAttributedStringGetAttributesAndLongestEffectiveRange(aStr CFAttributedStringRef, loc ,  CFIndex, inRange ,  CFRange, longestEffectiveRange ,  CFRange  *, ) CFDictionaryRef

// CFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString CFAttributedStringRef, range ,  CFRange, baseDirection ,  int8_t, bidiLevels ,  uint8_t  *, baseDirections ,  uint8_t  *, ) bool


// CFAttributedStringGetLength(aStr CFAttributedStringRef, ) CFIndex

// CFAttributedStringGetMutableString(aStr CFMutableAttributedStringRef, ) CFMutableStringRef

// CFAttributedStringGetStatisticalWritingDirections(attributedString CFAttributedStringRef, range ,  CFRange, baseDirection ,  int8_t, bidiLevels ,  uint8_t  *, baseDirections ,  uint8_t  *, ) bool
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+


// CFAttributedStringGetString(aStr CFAttributedStringRef, ) CFStringRef

// CFAttributedStringRemoveAttribute(aStr CFMutableAttributedStringRef, range ,  CFRange, attrName ,  CFStringRef, )

// CFAttributedStringReplaceAttributedString(aStr CFMutableAttributedStringRef, range ,  CFRange, replacement ,  CFAttributedStringRef, )


// CFAttributedStringReplaceString(aStr CFMutableAttributedStringRef, range ,  CFRange, replacement ,  CFStringRef, )

// CFAttributedStringSetAttribute(aStr CFMutableAttributedStringRef, range ,  CFRange, attrName ,  CFStringRef, value ,  CFTypeRef, )

// CFAttributedStringSetAttributes(aStr CFMutableAttributedStringRef, range ,  CFRange, replacement ,  CFDictionaryRef, clearOtherAttributes ,  Boolean, )


// CFAutorelease(arg CFTypeRef, ) CFTypeRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFBagAddValue(theBag CFMutableBagRef, value ,  const void  *, )

// CFBagApplyFunction(theBag CFBagRef, applier ,  CFBagApplierFunction, context ,  void  *, )


// CFBagContainsValue(theBag CFBagRef, value ,  const void  *, ) Boolean

// CFBagCreate(allocator CFAllocatorRef, values ,  const void  * *, numValues ,  CFIndex, callBacks ,  const CFBagCallBacks  *, ) CFBagRef

// CFBagCreateCopy(allocator CFAllocatorRef, theBag ,  CFBagRef, ) CFBagRef


// CFBagCreateMutable(allocator CFAllocatorRef, capacity ,  CFIndex, callBacks ,  const CFBagCallBacks  *, ) CFMutableBagRef

// CFBagCreateMutableCopy(allocator CFAllocatorRef, capacity ,  CFIndex, theBag ,  CFBagRef, ) CFMutableBagRef

// CFBagGetCount(theBag CFBagRef, ) CFIndex


// CFBagGetCountOfValue(theBag CFBagRef, value ,  const void  *, ) CFIndex

// CFBagGetValue(theBag CFBagRef, value ,  const void  *, ) const void  *

// CFBagGetValueIfPresent(theBag CFBagRef, candidate ,  const void  *, value ,  const void  * *, ) Boolean


// CFBagGetValues(theBag CFBagRef, values ,  const void  * *, )

// CFBagRemoveAllValues(theBag CFMutableBagRef, )

// CFBagRemoveValue(theBag CFMutableBagRef, value ,  const void  *, )


// CFBagReplaceValue(theBag CFMutableBagRef, value ,  const void  *, )

// CFBagSetValue(theBag CFMutableBagRef, value ,  const void  *, )

// CFBinaryHeapAddValue(heap CFBinaryHeapRef, value ,  const void  *, )


// CFBinaryHeapApplyFunction(heap CFBinaryHeapRef, applier ,  CFBinaryHeapApplierFunction, context ,  void  *, )

// CFBinaryHeapContainsValue(heap CFBinaryHeapRef, value ,  const void  *, ) Boolean

// CFBinaryHeapCreate(allocator CFAllocatorRef, capacity ,  CFIndex, callBacks ,  const CFBinaryHeapCallBacks  *, compareContext ,  const CFBinaryHeapCompareContext  *, ) CFBinaryHeapRef


// CFBinaryHeapCreateCopy(allocator CFAllocatorRef, capacity ,  CFIndex, heap ,  CFBinaryHeapRef, ) CFBinaryHeapRef

// CFBinaryHeapGetCount(heap CFBinaryHeapRef, ) CFIndex

// CFBinaryHeapGetCountOfValue(heap CFBinaryHeapRef, value ,  const void  *, ) CFIndex


// CFBinaryHeapGetMinimum(heap CFBinaryHeapRef, ) const void  *

// CFBinaryHeapGetMinimumIfPresent(heap CFBinaryHeapRef, value ,  const void  * *, ) Boolean

// CFBinaryHeapGetValues(heap CFBinaryHeapRef, values ,  const void  * *, )


// CFBinaryHeapRemoveAllValues(heap CFBinaryHeapRef, )

// CFBinaryHeapRemoveMinimumValue(heap CFBinaryHeapRef, )

// CFBitVectorContainsBit(bv CFBitVectorRef, range ,  CFRange, value ,  CFBit, ) Boolean


// CFBitVectorCreate(allocator CFAllocatorRef, bytes ,  const UInt8  *, numBits ,  CFIndex, ) CFBitVectorRef

// CFBitVectorCreateCopy(allocator CFAllocatorRef, bv ,  CFBitVectorRef, ) CFBitVectorRef

// CFBitVectorCreateMutable(allocator CFAllocatorRef, capacity ,  CFIndex, ) CFMutableBitVectorRef


// CFBitVectorCreateMutableCopy(allocator CFAllocatorRef, capacity ,  CFIndex, bv ,  CFBitVectorRef, ) CFMutableBitVectorRef

// CFBitVectorFlipBitAtIndex(bv CFMutableBitVectorRef, idx ,  CFIndex, )

// CFBitVectorFlipBits(bv CFMutableBitVectorRef, range ,  CFRange, )


// CFBitVectorGetBitAtIndex(bv CFBitVectorRef, idx ,  CFIndex, ) CFBit

// CFBitVectorGetBits(bv CFBitVectorRef, range ,  CFRange, bytes ,  UInt8  *, )

// CFBitVectorGetCount(bv CFBitVectorRef, ) CFIndex


// CFBitVectorGetCountOfBit(bv CFBitVectorRef, range ,  CFRange, value ,  CFBit, ) CFIndex

// CFBitVectorGetFirstIndexOfBit(bv CFBitVectorRef, range ,  CFRange, value ,  CFBit, ) CFIndex

// CFBitVectorGetLastIndexOfBit(bv CFBitVectorRef, range ,  CFRange, value ,  CFBit, ) CFIndex


// CFBitVectorSetAllBits(bv CFMutableBitVectorRef, value ,  CFBit, )

// CFBitVectorSetBitAtIndex(bv CFMutableBitVectorRef, idx ,  CFIndex, value ,  CFBit, )

// CFBitVectorSetBits(bv CFMutableBitVectorRef, range ,  CFRange, value ,  CFBit, )


// CFBitVectorSetCount(bv CFMutableBitVectorRef, count ,  CFIndex, )

// CFBooleanGetValue(boolean CFBooleanRef, ) Boolean

// CFBundleCloseBundleResourceMap(bundle CFBundleRef, refNum ,  CFBundleRefNum, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.15)
//
// Deprecated: This function is deprecated.


// CFBundleCopyAuxiliaryExecutableURL(bundle CFBundleRef, executableName ,  CFStringRef, ) CFURLRef

// CFBundleCopyBuiltInPlugInsURL(bundle CFBundleRef, ) CFURLRef

// CFBundleCopyBundleLocalizations(bundle CFBundleRef, ) CFArrayRef


// CFBundleCopyBundleURL(bundle CFBundleRef, ) CFURLRef

// CFBundleCopyExecutableArchitectures(bundle CFBundleRef, ) CFArrayRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFBundleCopyExecutableArchitecturesForURL(url CFURLRef, ) CFArrayRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFBundleCopyExecutableURL(bundle CFBundleRef, ) CFURLRef

// CFBundleCopyInfoDictionaryForURL(url CFURLRef, ) CFDictionaryRef

// CFBundleCopyInfoDictionaryInDirectory(bundleURL CFURLRef, ) CFDictionaryRef


// CFBundleCopyLocalizationsForPreferences(locArray CFArrayRef, prefArray ,  CFArrayRef, ) CFArrayRef

// CFBundleCopyLocalizationsForURL(url CFURLRef, ) CFArrayRef

// CFBundleCopyLocalizedString(bundle CFBundleRef, key ,  CFStringRef, value ,  CFStringRef, tableName ,  CFStringRef, ) CFStringRef


// CFBundleCopyLocalizedStringForLocalizations(bundle CFBundleRef, key ,  CFStringRef, value ,  CFStringRef, tableName ,  CFStringRef, localizations ,  CFArrayRef, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 18.4+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 15.4+
//   - tvOS 18.4+
//   - visionOS 2.4+
//   - watchOS 11.4+

// CFBundleCopyPreferredLocalizationsFromArray(locArray CFArrayRef, ) CFArrayRef

// CFBundleCopyPrivateFrameworksURL(bundle CFBundleRef, ) CFURLRef


// CFBundleCopyResourceURL(bundle CFBundleRef, resourceName ,  CFStringRef, resourceType ,  CFStringRef, subDirName ,  CFStringRef, ) CFURLRef

// CFBundleCopyResourceURLForLocalization(bundle CFBundleRef, resourceName ,  CFStringRef, resourceType ,  CFStringRef, subDirName ,  CFStringRef, localizationName ,  CFStringRef, ) CFURLRef

// CFBundleCopyResourceURLInDirectory(bundleURL CFURLRef, resourceName ,  CFStringRef, resourceType ,  CFStringRef, subDirName ,  CFStringRef, ) CFURLRef


// CFBundleCopyResourceURLsOfType(bundle CFBundleRef, resourceType ,  CFStringRef, subDirName ,  CFStringRef, ) CFArrayRef

// CFBundleCopyResourceURLsOfTypeForLocalization(bundle CFBundleRef, resourceType ,  CFStringRef, subDirName ,  CFStringRef, localizationName ,  CFStringRef, ) CFArrayRef

// CFBundleCopyResourceURLsOfTypeInDirectory(bundleURL CFURLRef, resourceType ,  CFStringRef, subDirName ,  CFStringRef, ) CFArrayRef


// CFBundleCopyResourcesDirectoryURL(bundle CFBundleRef, ) CFURLRef

// CFBundleCopySharedFrameworksURL(bundle CFBundleRef, ) CFURLRef

// CFBundleCopySharedSupportURL(bundle CFBundleRef, ) CFURLRef


// CFBundleCopySupportFilesDirectoryURL(bundle CFBundleRef, ) CFURLRef

// CFBundleCreate(allocator CFAllocatorRef, bundleURL ,  CFURLRef, ) CFBundleRef

// CFBundleCreateBundlesFromDirectory(allocator CFAllocatorRef, directoryURL ,  CFURLRef, bundleType ,  CFStringRef, ) CFArrayRef


// CFBundleGetBundleWithIdentifier(bundleID CFStringRef, ) CFBundleRef

// CFBundleGetDataPointerForName(bundle CFBundleRef, symbolName ,  CFStringRef, ) void  *

// CFBundleGetDataPointersForNames(bundle CFBundleRef, symbolNames ,  CFArrayRef, stbl ,  void  *, [])


// CFBundleGetDevelopmentRegion(bundle CFBundleRef, ) CFStringRef

// CFBundleGetFunctionPointerForName(bundle CFBundleRef, functionName ,  CFStringRef, ) void  *

// CFBundleGetFunctionPointersForNames(bundle CFBundleRef, functionNames ,  CFArrayRef, ftbl ,  void  *, [])


// CFBundleGetIdentifier(bundle CFBundleRef, ) CFStringRef

// CFBundleGetInfoDictionary(bundle CFBundleRef, ) CFDictionaryRef

// CFBundleGetLocalInfoDictionary(bundle CFBundleRef, ) CFDictionaryRef


// CFBundleGetPackageInfo(bundle CFBundleRef, packageType ,  UInt32  *, packageCreator ,  UInt32  *, )

// CFBundleGetPackageInfoInDirectory(url CFURLRef, packageType ,  UInt32  *, packageCreator ,  UInt32  *, ) Boolean

// CFBundleGetPlugIn(bundle CFBundleRef, ) CFPlugInRef


// CFBundleGetValueForInfoDictionaryKey(bundle CFBundleRef, key ,  CFStringRef, ) CFTypeRef

// CFBundleGetVersionNumber(bundle CFBundleRef, ) UInt32

// CFBundleIsArchitectureLoadable(arch cpu_type_t, ) Boolean
//
// Availability:
//   - macOS 11.0+


// CFBundleIsExecutableLoadable(bundle CFBundleRef, ) Boolean
//
// Availability:
//   - macOS 11.0+

// CFBundleIsExecutableLoadableForURL(url CFURLRef, ) Boolean
//
// Availability:
//   - macOS 11.0+

// CFBundleIsExecutableLoaded(bundle CFBundleRef, ) Boolean


// CFBundleLoadExecutable(bundle CFBundleRef, ) Boolean

// CFBundleLoadExecutableAndReturnError(bundle CFBundleRef, error ,  CFErrorRef  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFBundleOpenBundleResourceFiles(bundle CFBundleRef, refNum ,  CFBundleRefNum  *, localizedRefNum ,  CFBundleRefNum  *, ) SInt32
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.15)
//
// Deprecated: This function is deprecated.


// CFBundleOpenBundleResourceMap(bundle CFBundleRef, ) CFBundleRefNum
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.15)
//
// Deprecated: This function is deprecated.

// CFBundlePreflightExecutable(bundle CFBundleRef, error ,  CFErrorRef  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFBundleUnloadExecutable(bundle CFBundleRef, )


// CFCalendarAddComponents(calendar CFCalendarRef, at ,  CFAbsoluteTime  *, options ,  CFOptionFlags, componentDesc ,  const char  *, , ...) Boolean

// CFCalendarComposeAbsoluteTime(calendar CFCalendarRef, at ,  CFAbsoluteTime  *, componentDesc ,  const char  *, , ...) Boolean

// CFCalendarCopyLocale(calendar CFCalendarRef, ) CFLocaleRef


// CFCalendarCopyTimeZone(calendar CFCalendarRef, ) CFTimeZoneRef

// CFCalendarCreateWithIdentifier(allocator CFAllocatorRef, identifier ,  CFCalendarIdentifier, ) CFCalendarRef

// CFCalendarDecomposeAbsoluteTime(calendar CFCalendarRef, at ,  CFAbsoluteTime, componentDesc ,  const char  *, , ...) Boolean


// CFCalendarGetComponentDifference(calendar CFCalendarRef, startingAT ,  CFAbsoluteTime, resultAT ,  CFAbsoluteTime, options ,  CFOptionFlags, componentDesc ,  const char  *, , ...) Boolean

// CFCalendarGetFirstWeekday(calendar CFCalendarRef, ) CFIndex

// CFCalendarGetIdentifier(calendar CFCalendarRef, ) CFCalendarIdentifier


// CFCalendarGetMaximumRangeOfUnit(calendar CFCalendarRef, unit ,  CFCalendarUnit, ) CFRange

// CFCalendarGetMinimumDaysInFirstWeek(calendar CFCalendarRef, ) CFIndex

// CFCalendarGetMinimumRangeOfUnit(calendar CFCalendarRef, unit ,  CFCalendarUnit, ) CFRange


// CFCalendarGetOrdinalityOfUnit(calendar CFCalendarRef, smallerUnit ,  CFCalendarUnit, biggerUnit ,  CFCalendarUnit, at ,  CFAbsoluteTime, ) CFIndex

// CFCalendarGetRangeOfUnit(calendar CFCalendarRef, smallerUnit ,  CFCalendarUnit, biggerUnit ,  CFCalendarUnit, at ,  CFAbsoluteTime, ) CFRange

// CFCalendarGetTimeRangeOfUnit(calendar CFCalendarRef, unit ,  CFCalendarUnit, at ,  CFAbsoluteTime, startp ,  CFAbsoluteTime  *, tip ,  CFTimeInterval  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFCalendarSetFirstWeekday(calendar CFCalendarRef, wkdy ,  CFIndex, )

// CFCalendarSetLocale(calendar CFCalendarRef, locale ,  CFLocaleRef, )

// CFCalendarSetMinimumDaysInFirstWeek(calendar CFCalendarRef, mwd ,  CFIndex, )


// CFCalendarSetTimeZone(calendar CFCalendarRef, tz ,  CFTimeZoneRef, )

// CFCharacterSetAddCharactersInRange(theSet CFMutableCharacterSetRef, theRange ,  CFRange, )

// CFCharacterSetAddCharactersInString(theSet CFMutableCharacterSetRef, theString ,  CFStringRef, )


// CFCharacterSetCreateBitmapRepresentation(alloc CFAllocatorRef, theSet ,  CFCharacterSetRef, ) CFDataRef

// CFCharacterSetCreateCopy(alloc CFAllocatorRef, theSet ,  CFCharacterSetRef, ) CFCharacterSetRef

// CFCharacterSetCreateInvertedSet(alloc CFAllocatorRef, theSet ,  CFCharacterSetRef, ) CFCharacterSetRef


// CFCharacterSetCreateMutable(alloc CFAllocatorRef, ) CFMutableCharacterSetRef

// CFCharacterSetCreateMutableCopy(alloc CFAllocatorRef, theSet ,  CFCharacterSetRef, ) CFMutableCharacterSetRef

// CFCharacterSetCreateWithBitmapRepresentation(alloc CFAllocatorRef, theData ,  CFDataRef, ) CFCharacterSetRef


// CFCharacterSetCreateWithCharactersInRange(alloc CFAllocatorRef, theRange ,  CFRange, ) CFCharacterSetRef

// CFCharacterSetCreateWithCharactersInString(alloc CFAllocatorRef, theString ,  CFStringRef, ) CFCharacterSetRef

// CFCharacterSetGetPredefined(theSetIdentifier CFCharacterSetPredefinedSet, ) CFCharacterSetRef


// CFCharacterSetHasMemberInPlane(theSet CFCharacterSetRef, thePlane ,  CFIndex, ) Boolean

// CFCharacterSetIntersect(theSet CFMutableCharacterSetRef, theOtherSet ,  CFCharacterSetRef, )

// CFCharacterSetInvert(theSet CFMutableCharacterSetRef, )


// CFCharacterSetIsCharacterMember(theSet CFCharacterSetRef, theChar ,  UniChar, ) Boolean

// CFCharacterSetIsLongCharacterMember(theSet CFCharacterSetRef, theChar ,  UTF32Char, ) Boolean

// CFCharacterSetIsSupersetOfSet(theSet CFCharacterSetRef, theOtherset ,  CFCharacterSetRef, ) Boolean


// CFCharacterSetRemoveCharactersInRange(theSet CFMutableCharacterSetRef, theRange ,  CFRange, )

// CFCharacterSetRemoveCharactersInString(theSet CFMutableCharacterSetRef, theString ,  CFStringRef, )

// CFCharacterSetUnion(theSet CFMutableCharacterSetRef, theOtherSet ,  CFCharacterSetRef, )


// CFCopyDescription(cf CFTypeRef, ) CFStringRef

// CFDataAppendBytes(theData CFMutableDataRef, bytes ,  const UInt8  *, length ,  CFIndex, )

// CFDataCreate(allocator CFAllocatorRef, bytes ,  const UInt8  *, length ,  CFIndex, ) CFDataRef


// CFDataCreateCopy(allocator CFAllocatorRef, theData ,  CFDataRef, ) CFDataRef

// CFDataCreateMutable(allocator CFAllocatorRef, capacity ,  CFIndex, ) CFMutableDataRef

// CFDataCreateMutableCopy(allocator CFAllocatorRef, capacity ,  CFIndex, theData ,  CFDataRef, ) CFMutableDataRef


// CFDataCreateWithBytesNoCopy(allocator CFAllocatorRef, bytes ,  const UInt8  *, length ,  CFIndex, bytesDeallocator ,  CFAllocatorRef, ) CFDataRef

// CFDataDeleteBytes(theData CFMutableDataRef, range ,  CFRange, )

// CFDataFind(theData CFDataRef, dataToFind ,  CFDataRef, searchRange ,  CFRange, compareOptions ,  CFDataSearchFlags, ) CFRange
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFDataGetBytePtr(theData CFDataRef, ) const UInt8  *

// CFDataGetBytes(theData CFDataRef, range ,  CFRange, buffer ,  UInt8  *, )

// CFDataGetLength(theData CFDataRef, ) CFIndex


// CFDataGetMutableBytePtr(theData CFMutableDataRef, ) UInt8  *

// CFDataIncreaseLength(theData CFMutableDataRef, extraLength ,  CFIndex, )

// CFDataReplaceBytes(theData CFMutableDataRef, range ,  CFRange, newBytes ,  const UInt8  *, newLength ,  CFIndex, )


// CFDataSetLength(theData CFMutableDataRef, length ,  CFIndex, )

// CFDateCompare(theDate CFDateRef, otherDate ,  CFDateRef, context ,  void  *, ) CFComparisonResult

// CFDateCreate(allocator CFAllocatorRef, at ,  CFAbsoluteTime, ) CFDateRef


// CFDateFormatterCopyProperty(formatter CFDateFormatterRef, key ,  CFDateFormatterKey, ) CFTypeRef

// CFDateFormatterCreate(allocator CFAllocatorRef, locale ,  CFLocaleRef, dateStyle ,  CFDateFormatterStyle, timeStyle ,  CFDateFormatterStyle, ) CFDateFormatterRef

// CFDateFormatterCreateDateFormatFromTemplate(allocator CFAllocatorRef, tmplate ,  CFStringRef, options ,  CFOptionFlags, locale ,  CFLocaleRef, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFDateFormatterCreateDateFromString(allocator CFAllocatorRef, formatter ,  CFDateFormatterRef, string ,  CFStringRef, rangep ,  CFRange  *, ) CFDateRef

// CFDateFormatterCreateISO8601Formatter(allocator CFAllocatorRef, formatOptions ,  CFISO8601DateFormatOptions, ) CFDateFormatterRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+

// CFDateFormatterCreateStringWithAbsoluteTime(allocator CFAllocatorRef, formatter ,  CFDateFormatterRef, at ,  CFAbsoluteTime, ) CFStringRef


// CFDateFormatterCreateStringWithDate(allocator CFAllocatorRef, formatter ,  CFDateFormatterRef, date ,  CFDateRef, ) CFStringRef

// CFDateFormatterGetAbsoluteTimeFromString(formatter CFDateFormatterRef, string ,  CFStringRef, rangep ,  CFRange  *, atp ,  CFAbsoluteTime  *, ) Boolean

// CFDateFormatterGetDateStyle(formatter CFDateFormatterRef, ) CFDateFormatterStyle


// CFDateFormatterGetFormat(formatter CFDateFormatterRef, ) CFStringRef

// CFDateFormatterGetLocale(formatter CFDateFormatterRef, ) CFLocaleRef

// CFDateFormatterGetTimeStyle(formatter CFDateFormatterRef, ) CFDateFormatterStyle


// CFDateFormatterSetFormat(formatter CFDateFormatterRef, formatString ,  CFStringRef, )

// CFDateFormatterSetProperty(formatter CFDateFormatterRef, key ,  CFStringRef, value ,  CFTypeRef, )

// CFDateGetAbsoluteTime(theDate CFDateRef, ) CFAbsoluteTime


// CFDateGetTimeIntervalSinceDate(theDate CFDateRef, otherDate ,  CFDateRef, ) CFTimeInterval

// CFDictionaryAddValue(theDict CFMutableDictionaryRef, key ,  const void  *, value ,  const void  *, )

// CFDictionaryApplyFunction(theDict CFDictionaryRef, applier ,  CFDictionaryApplierFunction, context ,  void  *, )


// CFDictionaryContainsKey(theDict CFDictionaryRef, key ,  const void  *, ) Boolean

// CFDictionaryContainsValue(theDict CFDictionaryRef, value ,  const void  *, ) Boolean

// CFDictionaryCreate(allocator CFAllocatorRef, keys ,  const void  * *, values ,  const void  * *, numValues ,  CFIndex, keyCallBacks ,  const CFDictionaryKeyCallBacks  *, valueCallBacks ,  const CFDictionaryValueCallBacks  *, ) CFDictionaryRef


// CFDictionaryCreateCopy(allocator CFAllocatorRef, theDict ,  CFDictionaryRef, ) CFDictionaryRef

// CFDictionaryCreateMutable(allocator CFAllocatorRef, capacity ,  CFIndex, keyCallBacks ,  const CFDictionaryKeyCallBacks  *, valueCallBacks ,  const CFDictionaryValueCallBacks  *, ) CFMutableDictionaryRef

// CFDictionaryCreateMutableCopy(allocator CFAllocatorRef, capacity ,  CFIndex, theDict ,  CFDictionaryRef, ) CFMutableDictionaryRef


// CFDictionaryGetCount(theDict CFDictionaryRef, ) CFIndex

// CFDictionaryGetCountOfKey(theDict CFDictionaryRef, key ,  const void  *, ) CFIndex

// CFDictionaryGetCountOfValue(theDict CFDictionaryRef, value ,  const void  *, ) CFIndex


// CFDictionaryGetKeysAndValues(theDict CFDictionaryRef, keys ,  const void  * *, values ,  const void  * *, )

// CFDictionaryGetValue(theDict CFDictionaryRef, key ,  const void  *, ) const void  *

// CFDictionaryGetValueIfPresent(theDict CFDictionaryRef, key ,  const void  *, value ,  const void  * *, ) Boolean


// CFDictionaryRemoveAllValues(theDict CFMutableDictionaryRef, )

// CFDictionaryRemoveValue(theDict CFMutableDictionaryRef, key ,  const void  *, )

// CFDictionaryReplaceValue(theDict CFMutableDictionaryRef, key ,  const void  *, value ,  const void  *, )


// CFDictionarySetValue(theDict CFMutableDictionaryRef, key ,  const void  *, value ,  const void  *, )

// CFEqual(cf1 CFTypeRef, cf2 ,  CFTypeRef, ) Boolean

// CFErrorCopyDescription(err CFErrorRef, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFErrorCopyFailureReason(err CFErrorRef, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFErrorCopyRecoverySuggestion(err CFErrorRef, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFErrorCopyUserInfo(err CFErrorRef, ) CFDictionaryRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFErrorCreate(allocator CFAllocatorRef, domain ,  CFErrorDomain, code ,  CFIndex, userInfo ,  CFDictionaryRef, ) CFErrorRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFErrorCreateWithUserInfoKeysAndValues(allocator CFAllocatorRef, domain ,  CFErrorDomain, code ,  CFIndex, userInfoKeys ,  const void  *  const  *, userInfoValues ,  const void  *  const  *, numUserInfoValues ,  CFIndex, ) CFErrorRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFErrorGetCode(err CFErrorRef, ) CFIndex
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFErrorGetDomain(err CFErrorRef, ) CFErrorDomain
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorCreate(allocator CFAllocatorRef, fd ,  CFFileDescriptorNativeDescriptor, closeOnInvalidate ,  Boolean, callout ,  CFFileDescriptorCallBack, context ,  const CFFileDescriptorContext  *, ) CFFileDescriptorRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorCreateRunLoopSource(allocator CFAllocatorRef, f ,  CFFileDescriptorRef, order ,  CFIndex, ) CFRunLoopSourceRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileDescriptorDisableCallBacks(f CFFileDescriptorRef, callBackTypes ,  CFOptionFlags, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorEnableCallBacks(f CFFileDescriptorRef, callBackTypes ,  CFOptionFlags, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorGetContext(f CFFileDescriptorRef, context ,  CFFileDescriptorContext  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileDescriptorGetNativeDescriptor(f CFFileDescriptorRef, ) CFFileDescriptorNativeDescriptor
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorInvalidate(f CFFileDescriptorRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorIsValid(f CFFileDescriptorRef, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileSecurityClearProperties(fileSec CFFileSecurityRef, clearPropertyMask ,  CFFileSecurityClearOptions, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityCopyAccessControlList(fileSec CFFileSecurityRef, accessControlList ,  acl_t  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityCopyGroupUUID(fileSec CFFileSecurityRef, groupUUID ,  CFUUIDRef  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileSecurityCopyOwnerUUID(fileSec CFFileSecurityRef, ownerUUID ,  CFUUIDRef  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityCreate(allocator CFAllocatorRef, ) CFFileSecurityRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityCreateCopy(allocator CFAllocatorRef, fileSec ,  CFFileSecurityRef, ) CFFileSecurityRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileSecurityGetGroup(fileSec CFFileSecurityRef, group ,  gid_t  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityGetMode(fileSec CFFileSecurityRef, mode ,  mode_t  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityGetOwner(fileSec CFFileSecurityRef, owner ,  uid_t  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileSecuritySetAccessControlList(fileSec CFFileSecurityRef, accessControlList ,  acl_t, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecuritySetGroup(fileSec CFFileSecurityRef, group ,  gid_t, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecuritySetGroupUUID(fileSec CFFileSecurityRef, groupUUID ,  CFUUIDRef, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileSecuritySetMode(fileSec CFFileSecurityRef, mode ,  mode_t, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecuritySetOwner(fileSec CFFileSecurityRef, owner ,  uid_t, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecuritySetOwnerUUID(fileSec CFFileSecurityRef, ownerUUID ,  CFUUIDRef, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFGetTypeID(cf CFTypeRef, ) CFTypeID

// CFGregorianDateGetAbsoluteTime(gdate CFGregorianDate, tz ,  CFTimeZoneRef, ) CFAbsoluteTime
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.4+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFGregorianDateIsValid(gdate CFGregorianDate, unitFlags ,  CFOptionFlags, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.4+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.


// CFHash(cf CFTypeRef, ) CFHashCode

// CFLocaleCopyDisplayNameForPropertyValue(displayLocale CFLocaleRef, key ,  CFLocaleKey, value ,  CFStringRef, ) CFStringRef

// CFLocaleCreate(allocator CFAllocatorRef, localeIdentifier ,  CFLocaleIdentifier, ) CFLocaleRef


// CFLocaleCreateCanonicalLanguageIdentifierFromString(allocator CFAllocatorRef, localeIdentifier ,  CFStringRef, ) CFLocaleIdentifier

// CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator CFAllocatorRef, lcode ,  LangCode, rcode ,  RegionCode, ) CFLocaleIdentifier

// CFLocaleCreateCanonicalLocaleIdentifierFromString(allocator CFAllocatorRef, localeIdentifier ,  CFStringRef, ) CFLocaleIdentifier


// CFLocaleCreateComponentsFromLocaleIdentifier(allocator CFAllocatorRef, localeID ,  CFLocaleIdentifier, ) CFDictionaryRef

// CFLocaleCreateCopy(allocator CFAllocatorRef, locale ,  CFLocaleRef, ) CFLocaleRef

// CFLocaleCreateLocaleIdentifierFromComponents(allocator CFAllocatorRef, dictionary ,  CFDictionaryRef, ) CFLocaleIdentifier


// CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator CFAllocatorRef, lcid ,  uint32_t, ) CFLocaleIdentifier
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFLocaleGetIdentifier(locale CFLocaleRef, ) CFLocaleIdentifier

// CFLocaleGetLanguageCharacterDirection(isoLangCode CFStringRef, ) CFLocaleLanguageDirection
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFLocaleGetLanguageLineDirection(isoLangCode CFStringRef, ) CFLocaleLanguageDirection
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFLocaleGetValue(locale CFLocaleRef, key ,  CFLocaleKey, ) CFTypeRef

// CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier CFLocaleIdentifier, ) uint32_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFMachPortCreate(allocator CFAllocatorRef, callout ,  CFMachPortCallBack, context ,  CFMachPortContext  *, shouldFreeInfo ,  Boolean  *, ) CFMachPortRef

// CFMachPortCreateRunLoopSource(allocator CFAllocatorRef, port ,  CFMachPortRef, order ,  CFIndex, ) CFRunLoopSourceRef

// CFMachPortCreateWithPort(allocator CFAllocatorRef, portNum ,  mach_port_t, callout ,  CFMachPortCallBack, context ,  CFMachPortContext  *, shouldFreeInfo ,  Boolean  *, ) CFMachPortRef


// CFMachPortGetContext(port CFMachPortRef, context ,  CFMachPortContext  *, )

// CFMachPortGetInvalidationCallBack(port CFMachPortRef, ) CFMachPortInvalidationCallBack

// CFMachPortGetPort(port CFMachPortRef, ) mach_port_t


// CFMachPortInvalidate(port CFMachPortRef, )

// CFMachPortIsValid(port CFMachPortRef, ) Boolean

// CFMachPortSetInvalidationCallBack(port CFMachPortRef, callout ,  CFMachPortInvalidationCallBack, )


// CFMessagePortCreateLocal(allocator CFAllocatorRef, name ,  CFStringRef, callout ,  CFMessagePortCallBack, context ,  CFMessagePortContext  *, shouldFreeInfo ,  Boolean  *, ) CFMessagePortRef

// CFMessagePortCreateRemote(allocator CFAllocatorRef, name ,  CFStringRef, ) CFMessagePortRef

// CFMessagePortCreateRunLoopSource(allocator CFAllocatorRef, local ,  CFMessagePortRef, order ,  CFIndex, ) CFRunLoopSourceRef


// CFMessagePortGetContext(ms CFMessagePortRef, context ,  CFMessagePortContext  *, )

// CFMessagePortGetInvalidationCallBack(ms CFMessagePortRef, ) CFMessagePortInvalidationCallBack

// CFMessagePortGetName(ms CFMessagePortRef, ) CFStringRef


// CFMessagePortInvalidate(ms CFMessagePortRef, )

// CFMessagePortIsRemote(ms CFMessagePortRef, ) Boolean

// CFMessagePortIsValid(ms CFMessagePortRef, ) Boolean


// CFMessagePortSendRequest(remote CFMessagePortRef, msgid ,  SInt32, data ,  CFDataRef, sendTimeout ,  CFTimeInterval, rcvTimeout ,  CFTimeInterval, replyMode ,  CFStringRef, returnData ,  CFDataRef  *, ) SInt32

// CFMessagePortSetDispatchQueue(ms CFMessagePortRef, queue ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFMessagePortSetInvalidationCallBack(ms CFMessagePortRef, callout ,  CFMessagePortInvalidationCallBack, )


// CFMessagePortSetName(ms CFMessagePortRef, newName ,  CFStringRef, ) Boolean

// CFNotificationCenterAddObserver(center CFNotificationCenterRef, observer ,  const void  *, callBack ,  CFNotificationCallback, name ,  CFStringRef, object ,  const void  *, suspensionBehavior ,  CFNotificationSuspensionBehavior, )

// CFNotificationCenterPostNotification(center CFNotificationCenterRef, name ,  CFNotificationName, object ,  const void  *, userInfo ,  CFDictionaryRef, deliverImmediately ,  Boolean, )


// CFNotificationCenterPostNotificationWithOptions(center CFNotificationCenterRef, name ,  CFNotificationName, object ,  const void  *, userInfo ,  CFDictionaryRef, options ,  CFOptionFlags, )

// CFNotificationCenterRemoveEveryObserver(center CFNotificationCenterRef, observer ,  const void  *, )

// CFNotificationCenterRemoveObserver(center CFNotificationCenterRef, observer ,  const void  *, name ,  CFNotificationName, object ,  const void  *, )


// CFNumberCompare(number CFNumberRef, otherNumber ,  CFNumberRef, context ,  void  *, ) CFComparisonResult

// CFNumberCreate(allocator CFAllocatorRef, theType ,  CFNumberType, valuePtr ,  const void  *, ) CFNumberRef

// CFNumberFormatterCopyProperty(formatter CFNumberFormatterRef, key ,  CFNumberFormatterKey, ) CFTypeRef


// CFNumberFormatterCreate(allocator CFAllocatorRef, locale ,  CFLocaleRef, style ,  CFNumberFormatterStyle, ) CFNumberFormatterRef

// CFNumberFormatterCreateNumberFromString(allocator CFAllocatorRef, formatter ,  CFNumberFormatterRef, string ,  CFStringRef, rangep ,  CFRange  *, options ,  CFOptionFlags, ) CFNumberRef

// CFNumberFormatterCreateStringWithNumber(allocator CFAllocatorRef, formatter ,  CFNumberFormatterRef, number ,  CFNumberRef, ) CFStringRef


// CFNumberFormatterCreateStringWithValue(allocator CFAllocatorRef, formatter ,  CFNumberFormatterRef, numberType ,  CFNumberType, valuePtr ,  const void  *, ) CFStringRef

// CFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode CFStringRef, defaultFractionDigits ,  int32_t  *, roundingIncrement ,  double  *, ) Boolean

// CFNumberFormatterGetFormat(formatter CFNumberFormatterRef, ) CFStringRef


// CFNumberFormatterGetLocale(formatter CFNumberFormatterRef, ) CFLocaleRef

// CFNumberFormatterGetStyle(formatter CFNumberFormatterRef, ) CFNumberFormatterStyle

// CFNumberFormatterGetValueFromString(formatter CFNumberFormatterRef, string ,  CFStringRef, rangep ,  CFRange  *, numberType ,  CFNumberType, valuePtr ,  void  *, ) Boolean


// CFNumberFormatterSetFormat(formatter CFNumberFormatterRef, formatString ,  CFStringRef, )

// CFNumberFormatterSetProperty(formatter CFNumberFormatterRef, key ,  CFNumberFormatterKey, value ,  CFTypeRef, )

// CFNumberGetByteSize(number CFNumberRef, ) CFIndex


// CFNumberGetType(number CFNumberRef, ) CFNumberType

// CFNumberGetValue(number CFNumberRef, theType ,  CFNumberType, valuePtr ,  void  *, ) Boolean

// CFNumberIsFloatType(number CFNumberRef, ) Boolean


// CFPlugInAddInstanceForFactory(factoryID CFUUIDRef, )

// CFPlugInCreate(allocator CFAllocatorRef, plugInURL ,  CFURLRef, ) CFPlugInRef

// CFPlugInFindFactoriesForPlugInType(typeUUID CFUUIDRef, ) CFArrayRef


// CFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID CFUUIDRef, plugIn ,  CFPlugInRef, ) CFArrayRef

// CFPlugInGetBundle(plugIn CFPlugInRef, ) CFBundleRef

// CFPlugInInstanceCreate(allocator CFAllocatorRef, factoryUUID ,  CFUUIDRef, typeUUID ,  CFUUIDRef, ) void  *


// CFPlugInInstanceCreateWithInstanceDataSize(allocator CFAllocatorRef, instanceDataSize ,  CFIndex, deallocateInstanceFunction ,  CFPlugInInstanceDeallocateInstanceDataFunction, factoryName ,  CFStringRef, getInterfaceFunction ,  CFPlugInInstanceGetInterfaceFunction, ) CFPlugInInstanceRef

// CFPlugInInstanceGetFactoryName(instance CFPlugInInstanceRef, ) CFStringRef

// CFPlugInInstanceGetInstanceData(instance CFPlugInInstanceRef, ) void  *


// CFPlugInInstanceGetInterfaceFunctionTable(instance CFPlugInInstanceRef, interfaceName ,  CFStringRef, ftbl ,  void  * *, ) Boolean

// CFPlugInIsLoadOnDemand(plugIn CFPlugInRef, ) Boolean

// CFPlugInRegisterFactoryFunction(factoryUUID CFUUIDRef, func ,  CFPlugInFactoryFunction, ) Boolean


// CFPlugInRegisterFactoryFunctionByName(factoryUUID CFUUIDRef, plugIn ,  CFPlugInRef, functionName ,  CFStringRef, ) Boolean

// CFPlugInRegisterPlugInType(factoryUUID CFUUIDRef, typeUUID ,  CFUUIDRef, ) Boolean

// CFPlugInRemoveInstanceForFactory(factoryID CFUUIDRef, )


// CFPlugInSetLoadOnDemand(plugIn CFPlugInRef, flag ,  Boolean, )

// CFPlugInUnregisterFactory(factoryUUID CFUUIDRef, ) Boolean

// CFPlugInUnregisterPlugInType(factoryUUID CFUUIDRef, typeUUID ,  CFUUIDRef, ) Boolean


// CFPreferencesAddSuitePreferencesToApp(applicationID CFStringRef, suiteID ,  CFStringRef, )

// CFPreferencesAppSynchronize(applicationID CFStringRef, ) Boolean

// CFPreferencesAppValueIsForced(key CFStringRef, applicationID ,  CFStringRef, ) Boolean


// CFPreferencesCopyAppValue(key CFStringRef, applicationID ,  CFStringRef, ) CFPropertyListRef

// CFPreferencesCopyApplicationList(userName CFStringRef, hostName ,  CFStringRef, ) CFArrayRef
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFPreferencesCopyKeyList(applicationID CFStringRef, userName ,  CFStringRef, hostName ,  CFStringRef, ) CFArrayRef


// CFPreferencesCopyMultiple(keysToFetch CFArrayRef, applicationID ,  CFStringRef, userName ,  CFStringRef, hostName ,  CFStringRef, ) CFDictionaryRef

// CFPreferencesCopyValue(key CFStringRef, applicationID ,  CFStringRef, userName ,  CFStringRef, hostName ,  CFStringRef, ) CFPropertyListRef

// CFPreferencesGetAppBooleanValue(key CFStringRef, applicationID ,  CFStringRef, keyExistsAndHasValidFormat ,  Boolean  *, ) Boolean


// CFPreferencesGetAppIntegerValue(key CFStringRef, applicationID ,  CFStringRef, keyExistsAndHasValidFormat ,  Boolean  *, ) CFIndex

// CFPreferencesRemoveSuitePreferencesFromApp(applicationID CFStringRef, suiteID ,  CFStringRef, )

// CFPreferencesSetAppValue(key CFStringRef, value ,  CFPropertyListRef, applicationID ,  CFStringRef, )


// CFPreferencesSetMultiple(keysToSet CFDictionaryRef, keysToRemove ,  CFArrayRef, applicationID ,  CFStringRef, userName ,  CFStringRef, hostName ,  CFStringRef, )

// CFPreferencesSetValue(key CFStringRef, value ,  CFPropertyListRef, applicationID ,  CFStringRef, userName ,  CFStringRef, hostName ,  CFStringRef, )

// CFPreferencesSynchronize(applicationID CFStringRef, userName ,  CFStringRef, hostName ,  CFStringRef, ) Boolean


// CFPropertyListCreateData(allocator CFAllocatorRef, propertyList ,  CFPropertyListRef, format ,  CFPropertyListFormat, options ,  CFOptionFlags, error ,  CFErrorRef  *, ) CFDataRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFPropertyListCreateDeepCopy(allocator CFAllocatorRef, propertyList ,  CFPropertyListRef, mutabilityOption ,  CFOptionFlags, ) CFPropertyListRef

// CFPropertyListCreateFromStream(allocator CFAllocatorRef, stream ,  CFReadStreamRef, streamLength ,  CFIndex, mutabilityOption ,  CFOptionFlags, format ,  CFPropertyListFormat  *, errorString ,  CFStringRef  *, ) CFPropertyListRef
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.2+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.


// CFPropertyListCreateFromXMLData(allocator CFAllocatorRef, xmlData ,  CFDataRef, mutabilityOption ,  CFOptionFlags, errorString ,  CFStringRef  *, ) CFPropertyListRef
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.0+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFPropertyListCreateWithData(allocator CFAllocatorRef, data ,  CFDataRef, options ,  CFOptionFlags, format ,  CFPropertyListFormat  *, error ,  CFErrorRef  *, ) CFPropertyListRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFPropertyListCreateWithStream(allocator CFAllocatorRef, stream ,  CFReadStreamRef, streamLength ,  CFIndex, options ,  CFOptionFlags, format ,  CFPropertyListFormat  *, error ,  CFErrorRef  *, ) CFPropertyListRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFPropertyListCreateXMLData(allocator CFAllocatorRef, propertyList ,  CFPropertyListRef, ) CFDataRef
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.0+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFPropertyListIsValid(plist CFPropertyListRef, format ,  CFPropertyListFormat, ) Boolean

// CFPropertyListWrite(propertyList CFPropertyListRef, stream ,  CFWriteStreamRef, format ,  CFPropertyListFormat, options ,  CFOptionFlags, error ,  CFErrorRef  *, ) CFIndex
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFPropertyListWriteToStream(propertyList CFPropertyListRef, stream ,  CFWriteStreamRef, format ,  CFPropertyListFormat, errorString ,  CFStringRef  *, ) CFIndex
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 8.0)
//   - iPadOS 2.0+ (Deprecated in 8.0)
//   - macOS 10.2+ (Deprecated in 10.10)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFReadStreamClose(stream CFReadStreamRef, )

// CFReadStreamCopyDispatchQueue(stream CFReadStreamRef, ) dispatch_queue_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFReadStreamCopyError(stream CFReadStreamRef, ) CFErrorRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFReadStreamCopyProperty(stream CFReadStreamRef, propertyName ,  CFStreamPropertyKey, ) CFTypeRef

// CFReadStreamCreateWithBytesNoCopy(alloc CFAllocatorRef, bytes ,  const UInt8  *, length ,  CFIndex, bytesDeallocator ,  CFAllocatorRef, ) CFReadStreamRef


// CFReadStreamCreateWithFile(alloc CFAllocatorRef, fileURL ,  CFURLRef, ) CFReadStreamRef

// CFReadStreamGetBuffer(stream CFReadStreamRef, maxBytesToRead ,  CFIndex, numBytesRead ,  CFIndex  *, ) const UInt8  *

// CFReadStreamGetError(stream CFReadStreamRef, ) CFStreamError


// CFReadStreamGetStatus(stream CFReadStreamRef, ) CFStreamStatus

// CFReadStreamHasBytesAvailable(stream CFReadStreamRef, ) Boolean

// CFReadStreamOpen(stream CFReadStreamRef, ) Boolean


// CFReadStreamRead(stream CFReadStreamRef, buffer ,  UInt8  *, bufferLength ,  CFIndex, ) CFIndex

// CFReadStreamScheduleWithRunLoop(stream CFReadStreamRef, runLoop ,  CFRunLoopRef, runLoopMode ,  CFRunLoopMode, )

// CFReadStreamSetClient(stream CFReadStreamRef, streamEvents ,  CFOptionFlags, clientCB ,  CFReadStreamClientCallBack, clientContext ,  CFStreamClientContext  *, ) Boolean


// CFReadStreamSetDispatchQueue(stream CFReadStreamRef, q ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFReadStreamSetProperty(stream CFReadStreamRef, propertyName ,  CFStreamPropertyKey, propertyValue ,  CFTypeRef, ) Boolean

// CFReadStreamUnscheduleFromRunLoop(stream CFReadStreamRef, runLoop ,  CFRunLoopRef, runLoopMode ,  CFRunLoopMode, )


// CFRelease(cf CFTypeRef, )

// CFRunLoopAddCommonMode(rl CFRunLoopRef, mode ,  CFRunLoopMode, )

// CFRunLoopAddObserver(rl CFRunLoopRef, observer ,  CFRunLoopObserverRef, mode ,  CFRunLoopMode, )


// CFRunLoopAddSource(rl CFRunLoopRef, source ,  CFRunLoopSourceRef, mode ,  CFRunLoopMode, )

// CFRunLoopAddTimer(rl CFRunLoopRef, timer ,  CFRunLoopTimerRef, mode ,  CFRunLoopMode, )

// CFRunLoopContainsObserver(rl CFRunLoopRef, observer ,  CFRunLoopObserverRef, mode ,  CFRunLoopMode, ) Boolean


// CFRunLoopContainsSource(rl CFRunLoopRef, source ,  CFRunLoopSourceRef, mode ,  CFRunLoopMode, ) Boolean

// CFRunLoopContainsTimer(rl CFRunLoopRef, timer ,  CFRunLoopTimerRef, mode ,  CFRunLoopMode, ) Boolean

// CFRunLoopCopyAllModes(rl CFRunLoopRef, ) CFArrayRef


// CFRunLoopCopyCurrentMode(rl CFRunLoopRef, ) CFRunLoopMode

// CFRunLoopGetNextTimerFireDate(rl CFRunLoopRef, mode ,  CFRunLoopMode, ) CFAbsoluteTime

// CFRunLoopIsWaiting(rl CFRunLoopRef, ) Boolean


// CFRunLoopObserverCreate(allocator CFAllocatorRef, activities ,  CFOptionFlags, repeats ,  Boolean, order ,  CFIndex, callout ,  CFRunLoopObserverCallBack, context ,  CFRunLoopObserverContext  *, ) CFRunLoopObserverRef

// CFRunLoopObserverCreateWithHandler(allocator CFAllocatorRef, activities ,  CFOptionFlags, repeats ,  Boolean, order ,  CFIndex, block ,  void  (^, observer )( CFRunLoopObserverRef, activity ,  CFRunLoopActivity, ) CFRunLoopObserverRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFRunLoopObserverDoesRepeat(observer CFRunLoopObserverRef, ) Boolean


// CFRunLoopObserverGetActivities(observer CFRunLoopObserverRef, ) CFOptionFlags

// CFRunLoopObserverGetContext(observer CFRunLoopObserverRef, context ,  CFRunLoopObserverContext  *, )

// CFRunLoopObserverGetOrder(observer CFRunLoopObserverRef, ) CFIndex


// CFRunLoopObserverInvalidate(observer CFRunLoopObserverRef, )

// CFRunLoopObserverIsValid(observer CFRunLoopObserverRef, ) Boolean

// CFRunLoopPerformBlock(rl CFRunLoopRef, mode ,  CFTypeRef, block ,  void  (^, )()
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFRunLoopRemoveObserver(rl CFRunLoopRef, observer ,  CFRunLoopObserverRef, mode ,  CFRunLoopMode, )

// CFRunLoopRemoveSource(rl CFRunLoopRef, source ,  CFRunLoopSourceRef, mode ,  CFRunLoopMode, )

// CFRunLoopRemoveTimer(rl CFRunLoopRef, timer ,  CFRunLoopTimerRef, mode ,  CFRunLoopMode, )


// CFRunLoopRunInMode(mode CFRunLoopMode, seconds ,  CFTimeInterval, returnAfterSourceHandled ,  Boolean, ) CFRunLoopRunResult

// CFRunLoopSourceCreate(allocator CFAllocatorRef, order ,  CFIndex, context ,  CFRunLoopSourceContext  *, ) CFRunLoopSourceRef

// CFRunLoopSourceGetContext(source CFRunLoopSourceRef, context ,  CFRunLoopSourceContext  *, )


// CFRunLoopSourceGetOrder(source CFRunLoopSourceRef, ) CFIndex

// CFRunLoopSourceInvalidate(source CFRunLoopSourceRef, )

// CFRunLoopSourceIsValid(source CFRunLoopSourceRef, ) Boolean


// CFRunLoopSourceSignal(source CFRunLoopSourceRef, )

// CFRunLoopStop(rl CFRunLoopRef, )

// CFRunLoopTimerCreate(allocator CFAllocatorRef, fireDate ,  CFAbsoluteTime, interval ,  CFTimeInterval, flags ,  CFOptionFlags, order ,  CFIndex, callout ,  CFRunLoopTimerCallBack, context ,  CFRunLoopTimerContext  *, ) CFRunLoopTimerRef


// CFRunLoopTimerCreateWithHandler(allocator CFAllocatorRef, fireDate ,  CFAbsoluteTime, interval ,  CFTimeInterval, flags ,  CFOptionFlags, order ,  CFIndex, block ,  void  (^, timer )( CFRunLoopTimerRef, ) CFRunLoopTimerRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFRunLoopTimerDoesRepeat(timer CFRunLoopTimerRef, ) Boolean

// CFRunLoopTimerGetContext(timer CFRunLoopTimerRef, context ,  CFRunLoopTimerContext  *, )


// CFRunLoopTimerGetInterval(timer CFRunLoopTimerRef, ) CFTimeInterval

// CFRunLoopTimerGetNextFireDate(timer CFRunLoopTimerRef, ) CFAbsoluteTime

// CFRunLoopTimerGetOrder(timer CFRunLoopTimerRef, ) CFIndex


// CFRunLoopTimerGetTolerance(timer CFRunLoopTimerRef, ) CFTimeInterval
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFRunLoopTimerInvalidate(timer CFRunLoopTimerRef, )

// CFRunLoopTimerIsValid(timer CFRunLoopTimerRef, ) Boolean


// CFRunLoopTimerSetNextFireDate(timer CFRunLoopTimerRef, fireDate ,  CFAbsoluteTime, )

// CFRunLoopTimerSetTolerance(timer CFRunLoopTimerRef, tolerance ,  CFTimeInterval, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFRunLoopWakeUp(rl CFRunLoopRef, )


// CFSetAddValue(theSet CFMutableSetRef, value ,  const void  *, )

// CFSetApplyFunction(theSet CFSetRef, applier ,  CFSetApplierFunction, context ,  void  *, )

// CFSetContainsValue(theSet CFSetRef, value ,  const void  *, ) Boolean


// CFSetCreate(allocator CFAllocatorRef, values ,  const void  * *, numValues ,  CFIndex, callBacks ,  const CFSetCallBacks  *, ) CFSetRef

// CFSetCreateCopy(allocator CFAllocatorRef, theSet ,  CFSetRef, ) CFSetRef

// CFSetCreateMutable(allocator CFAllocatorRef, capacity ,  CFIndex, callBacks ,  const CFSetCallBacks  *, ) CFMutableSetRef


// CFSetCreateMutableCopy(allocator CFAllocatorRef, capacity ,  CFIndex, theSet ,  CFSetRef, ) CFMutableSetRef

// CFSetGetCount(theSet CFSetRef, ) CFIndex

// CFSetGetCountOfValue(theSet CFSetRef, value ,  const void  *, ) CFIndex


// CFSetGetValue(theSet CFSetRef, value ,  const void  *, ) const void  *

// CFSetGetValueIfPresent(theSet CFSetRef, candidate ,  const void  *, value ,  const void  * *, ) Boolean

// CFSetGetValues(theSet CFSetRef, values ,  const void  * *, )


// CFSetRemoveAllValues(theSet CFMutableSetRef, )

// CFSetRemoveValue(theSet CFMutableSetRef, value ,  const void  *, )

// CFSetReplaceValue(theSet CFMutableSetRef, value ,  const void  *, )


// CFSetSetValue(theSet CFMutableSetRef, value ,  const void  *, )

// CFShowStr(str CFStringRef, )

// CFSocketConnectToAddress(s CFSocketRef, address ,  CFDataRef, timeout ,  CFTimeInterval, ) CFSocketError


// CFSocketCopyAddress(s CFSocketRef, ) CFDataRef

// CFSocketCopyPeerAddress(s CFSocketRef, ) CFDataRef

// CFSocketCopyRegisteredSocketSignature(nameServerSignature const CFSocketSignature  *, timeout ,  CFTimeInterval, name ,  CFStringRef, signature ,  CFSocketSignature  *, nameServerAddress ,  CFDataRef  *, ) CFSocketError


// CFSocketCopyRegisteredValue(nameServerSignature const CFSocketSignature  *, timeout ,  CFTimeInterval, name ,  CFStringRef, value ,  CFPropertyListRef  *, nameServerAddress ,  CFDataRef  *, ) CFSocketError

// CFSocketCreate(allocator CFAllocatorRef, protocolFamily ,  SInt32, socketType ,  SInt32, protocol ,  SInt32, callBackTypes ,  CFOptionFlags, callout ,  CFSocketCallBack, context ,  const CFSocketContext  *, ) CFSocketRef

// CFSocketCreateConnectedToSocketSignature(allocator CFAllocatorRef, signature ,  const CFSocketSignature  *, callBackTypes ,  CFOptionFlags, callout ,  CFSocketCallBack, context ,  const CFSocketContext  *, timeout ,  CFTimeInterval, ) CFSocketRef


// CFSocketCreateRunLoopSource(allocator CFAllocatorRef, s ,  CFSocketRef, order ,  CFIndex, ) CFRunLoopSourceRef

// CFSocketCreateWithNative(allocator CFAllocatorRef, sock ,  CFSocketNativeHandle, callBackTypes ,  CFOptionFlags, callout ,  CFSocketCallBack, context ,  const CFSocketContext  *, ) CFSocketRef

// CFSocketCreateWithSocketSignature(allocator CFAllocatorRef, signature ,  const CFSocketSignature  *, callBackTypes ,  CFOptionFlags, callout ,  CFSocketCallBack, context ,  const CFSocketContext  *, ) CFSocketRef


// CFSocketDisableCallBacks(s CFSocketRef, callBackTypes ,  CFOptionFlags, )

// CFSocketEnableCallBacks(s CFSocketRef, callBackTypes ,  CFOptionFlags, )

// CFSocketGetContext(s CFSocketRef, context ,  CFSocketContext  *, )


// CFSocketGetNative(s CFSocketRef, ) CFSocketNativeHandle

// CFSocketGetSocketFlags(s CFSocketRef, ) CFOptionFlags

// CFSocketInvalidate(s CFSocketRef, )


// CFSocketIsValid(s CFSocketRef, ) Boolean

// CFSocketRegisterSocketSignature(nameServerSignature const CFSocketSignature  *, timeout ,  CFTimeInterval, name ,  CFStringRef, signature ,  const CFSocketSignature  *, ) CFSocketError

// CFSocketRegisterValue(nameServerSignature const CFSocketSignature  *, timeout ,  CFTimeInterval, name ,  CFStringRef, value ,  CFPropertyListRef, ) CFSocketError


// CFSocketSendData(s CFSocketRef, address ,  CFDataRef, data ,  CFDataRef, timeout ,  CFTimeInterval, ) CFSocketError

// CFSocketSetAddress(s CFSocketRef, address ,  CFDataRef, ) CFSocketError

// CFSocketSetDefaultNameRegistryPortNumber(port UInt16, )


// CFSocketSetSocketFlags(s CFSocketRef, flags ,  CFOptionFlags, )

// CFSocketUnregister(nameServerSignature const CFSocketSignature  *, timeout ,  CFTimeInterval, name ,  CFStringRef, ) CFSocketError

// CFStreamCreateBoundPair(alloc CFAllocatorRef, readStream ,  CFReadStreamRef  *, writeStream ,  CFWriteStreamRef  *, transferBufferSize ,  CFIndex, )


// CFStreamCreatePairWithPeerSocketSignature(alloc CFAllocatorRef, signature ,  const CFSocketSignature  *, readStream ,  CFReadStreamRef  *, writeStream ,  CFWriteStreamRef  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+ (Deprecated in 26.1)
//   - iPadOS 2.0+ (Deprecated in 26.1)
//   - macOS 10.1+ (Deprecated in 26.1)
//   - tvOS 9.0+ (Deprecated in 26.1)
//   - visionOS 1.0+ (Deprecated in 26.1)
//   - watchOS 2.0+ (Deprecated in 26.1)
//
// Deprecated: This function is deprecated.

// CFStreamCreatePairWithSocket(alloc CFAllocatorRef, sock ,  CFSocketNativeHandle, readStream ,  CFReadStreamRef  *, writeStream ,  CFWriteStreamRef  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+ (Deprecated in 26.1)
//   - iPadOS 2.0+ (Deprecated in 26.1)
//   - macOS 10.1+ (Deprecated in 26.1)
//   - tvOS 9.0+ (Deprecated in 26.1)
//   - visionOS 1.0+ (Deprecated in 26.1)
//   - watchOS 2.0+ (Deprecated in 26.1)
//
// Deprecated: This function is deprecated.

// CFStreamCreatePairWithSocketToHost(alloc CFAllocatorRef, host ,  CFStringRef, port ,  UInt32, readStream ,  CFReadStreamRef  *, writeStream ,  CFWriteStreamRef  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+ (Deprecated in 26.1)
//   - iPadOS 2.0+ (Deprecated in 26.1)
//   - macOS 10.1+ (Deprecated in 26.1)
//   - tvOS 9.0+ (Deprecated in 26.1)
//   - visionOS 1.0+ (Deprecated in 26.1)
//   - watchOS 2.0+ (Deprecated in 26.1)
//
// Deprecated: This function is deprecated.


// CFStringAppend(theString CFMutableStringRef, appendedString ,  CFStringRef, )

// CFStringAppendCString(theString CFMutableStringRef, cStr ,  const char  *, encoding ,  CFStringEncoding, )

// CFStringAppendCharacters(theString CFMutableStringRef, chars ,  const UniChar  *, numChars ,  CFIndex, )


// CFStringAppendFormat(theString CFMutableStringRef, formatOptions ,  CFDictionaryRef, format ,  CFStringRef, , ...)

// CFStringAppendFormatAndArguments(theString CFMutableStringRef, formatOptions ,  CFDictionaryRef, format ,  CFStringRef, arguments ,  va_list, )

// CFStringAppendPascalString(theString CFMutableStringRef, pStr ,  ConstStr255Param, encoding ,  CFStringEncoding, )


// CFStringCapitalize(theString CFMutableStringRef, locale ,  CFLocaleRef, )

// CFStringCompare(theString1 CFStringRef, theString2 ,  CFStringRef, compareOptions ,  CFStringCompareFlags, ) CFComparisonResult

// CFStringCompareWithOptions(theString1 CFStringRef, theString2 ,  CFStringRef, rangeToCompare ,  CFRange, compareOptions ,  CFStringCompareFlags, ) CFComparisonResult


// CFStringCompareWithOptionsAndLocale(theString1 CFStringRef, theString2 ,  CFStringRef, rangeToCompare ,  CFRange, compareOptions ,  CFStringCompareFlags, locale ,  CFLocaleRef, ) CFComparisonResult
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringConvertEncodingToIANACharSetName(encoding CFStringEncoding, ) CFStringRef

// CFStringConvertEncodingToNSStringEncoding(encoding CFStringEncoding, ) unsigned long


// CFStringConvertEncodingToWindowsCodepage(encoding CFStringEncoding, ) UInt32

// CFStringConvertIANACharSetNameToEncoding(theString CFStringRef, ) CFStringEncoding

// CFStringConvertNSStringEncodingToEncoding(encoding unsigned long, ) CFStringEncoding


// CFStringConvertWindowsCodepageToEncoding(codepage UInt32, ) CFStringEncoding

// CFStringCreateArrayBySeparatingStrings(alloc CFAllocatorRef, theString ,  CFStringRef, separatorString ,  CFStringRef, ) CFArrayRef

// CFStringCreateArrayWithFindResults(alloc CFAllocatorRef, theString ,  CFStringRef, stringToFind ,  CFStringRef, rangeToSearch ,  CFRange, compareOptions ,  CFStringCompareFlags, ) CFArrayRef


// CFStringCreateByCombiningStrings(alloc CFAllocatorRef, theArray ,  CFArrayRef, separatorString ,  CFStringRef, ) CFStringRef

// CFStringCreateCopy(alloc CFAllocatorRef, theString ,  CFStringRef, ) CFStringRef

// CFStringCreateExternalRepresentation(alloc CFAllocatorRef, theString ,  CFStringRef, encoding ,  CFStringEncoding, lossByte ,  UInt8, ) CFDataRef


// CFStringCreateFromExternalRepresentation(alloc CFAllocatorRef, data ,  CFDataRef, encoding ,  CFStringEncoding, ) CFStringRef

// CFStringCreateMutable(alloc CFAllocatorRef, maxLength ,  CFIndex, ) CFMutableStringRef

// CFStringCreateMutableCopy(alloc CFAllocatorRef, maxLength ,  CFIndex, theString ,  CFStringRef, ) CFMutableStringRef


// CFStringCreateMutableWithExternalCharactersNoCopy(alloc CFAllocatorRef, chars ,  UniChar  *, numChars ,  CFIndex, capacity ,  CFIndex, externalCharactersAllocator ,  CFAllocatorRef, ) CFMutableStringRef

// CFStringCreateStringWithValidatedFormat(alloc CFAllocatorRef, formatOptions ,  CFDictionaryRef, validFormatSpecifiers ,  CFStringRef, format ,  CFStringRef, errorPtr ,  CFErrorRef  *, , ...) CFStringRef
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 8.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// CFStringCreateStringWithValidatedFormatAndArguments(alloc CFAllocatorRef, formatOptions ,  CFDictionaryRef, validFormatSpecifiers ,  CFStringRef, format ,  CFStringRef, arguments ,  va_list, errorPtr ,  CFErrorRef  *, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 8.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// CFStringCreateWithBytes(alloc CFAllocatorRef, bytes ,  const UInt8  *, numBytes ,  CFIndex, encoding ,  CFStringEncoding, isExternalRepresentation ,  Boolean, ) CFStringRef

// CFStringCreateWithBytesNoCopy(alloc CFAllocatorRef, bytes ,  const UInt8  *, numBytes ,  CFIndex, encoding ,  CFStringEncoding, isExternalRepresentation ,  Boolean, contentsDeallocator ,  CFAllocatorRef, ) CFStringRef

// CFStringCreateWithCString(alloc CFAllocatorRef, cStr ,  const char  *, encoding ,  CFStringEncoding, ) CFStringRef


// CFStringCreateWithCStringNoCopy(alloc CFAllocatorRef, cStr ,  const char  *, encoding ,  CFStringEncoding, contentsDeallocator ,  CFAllocatorRef, ) CFStringRef

// CFStringCreateWithCharacters(alloc CFAllocatorRef, chars ,  const UniChar  *, numChars ,  CFIndex, ) CFStringRef

// CFStringCreateWithCharactersNoCopy(alloc CFAllocatorRef, chars ,  const UniChar  *, numChars ,  CFIndex, contentsDeallocator ,  CFAllocatorRef, ) CFStringRef


// CFStringCreateWithFileSystemRepresentation(alloc CFAllocatorRef, buffer ,  const char  *, ) CFStringRef

// CFStringCreateWithFormat(alloc CFAllocatorRef, formatOptions ,  CFDictionaryRef, format ,  CFStringRef, , ...) CFStringRef

// CFStringCreateWithFormatAndArguments(alloc CFAllocatorRef, formatOptions ,  CFDictionaryRef, format ,  CFStringRef, arguments ,  va_list, ) CFStringRef


// CFStringCreateWithPascalString(alloc CFAllocatorRef, pStr ,  ConstStr255Param, encoding ,  CFStringEncoding, ) CFStringRef

// CFStringCreateWithPascalStringNoCopy(alloc CFAllocatorRef, pStr ,  ConstStr255Param, encoding ,  CFStringEncoding, contentsDeallocator ,  CFAllocatorRef, ) CFStringRef

// CFStringCreateWithSubstring(alloc CFAllocatorRef, str ,  CFStringRef, range ,  CFRange, ) CFStringRef


// CFStringDelete(theString CFMutableStringRef, range ,  CFRange, )

// CFStringFind(theString CFStringRef, stringToFind ,  CFStringRef, compareOptions ,  CFStringCompareFlags, ) CFRange

// CFStringFindAndReplace(theString CFMutableStringRef, stringToFind ,  CFStringRef, replacementString ,  CFStringRef, rangeToSearch ,  CFRange, compareOptions ,  CFStringCompareFlags, ) CFIndex


// CFStringFindCharacterFromSet(theString CFStringRef, theSet ,  CFCharacterSetRef, rangeToSearch ,  CFRange, searchOptions ,  CFStringCompareFlags, result ,  CFRange  *, ) Boolean

// CFStringFindWithOptions(theString CFStringRef, stringToFind ,  CFStringRef, rangeToSearch ,  CFRange, searchOptions ,  CFStringCompareFlags, result ,  CFRange  *, ) Boolean

// CFStringFindWithOptionsAndLocale(theString CFStringRef, stringToFind ,  CFStringRef, rangeToSearch ,  CFRange, searchOptions ,  CFStringCompareFlags, locale ,  CFLocaleRef, result ,  CFRange  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFStringFold(theString CFMutableStringRef, theFlags ,  CFStringCompareFlags, theLocale ,  CFLocaleRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringGetBytes(theString CFStringRef, range ,  CFRange, encoding ,  CFStringEncoding, lossByte ,  UInt8, isExternalRepresentation ,  Boolean, buffer ,  UInt8  *, maxBufLen ,  CFIndex, usedBufLen ,  CFIndex  *, ) CFIndex

// CFStringGetCString(theString CFStringRef, buffer ,  char  *, bufferSize ,  CFIndex, encoding ,  CFStringEncoding, ) Boolean


// CFStringGetCStringPtr(theString CFStringRef, encoding ,  CFStringEncoding, ) const char  *

// CFStringGetCharacterAtIndex(theString CFStringRef, idx ,  CFIndex, ) UniChar

// CFStringGetCharacters(theString CFStringRef, range ,  CFRange, buffer ,  UniChar  *, )


// CFStringGetCharactersPtr(theString CFStringRef, ) const UniChar  *

// CFStringGetDoubleValue(str CFStringRef, ) double

// CFStringGetFastestEncoding(theString CFStringRef, ) CFStringEncoding


// CFStringGetFileSystemRepresentation(string CFStringRef, buffer ,  char  *, maxBufLen ,  CFIndex, ) Boolean

// CFStringGetHyphenationLocationBeforeIndex(string CFStringRef, location ,  CFIndex, limitRange ,  CFRange, options ,  CFOptionFlags, locale ,  CFLocaleRef, character ,  UTF32Char  *, ) CFIndex
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.2+
//   - iPadOS 4.2+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringGetIntValue(str CFStringRef, ) SInt32


// CFStringGetLength(theString CFStringRef, ) CFIndex

// CFStringGetLineBounds(theString CFStringRef, range ,  CFRange, lineBeginIndex ,  CFIndex  *, lineEndIndex ,  CFIndex  *, contentsEndIndex ,  CFIndex  *, )

// CFStringGetMaximumSizeForEncoding(length CFIndex, encoding ,  CFStringEncoding, ) CFIndex


// CFStringGetMaximumSizeOfFileSystemRepresentation(string CFStringRef, ) CFIndex

// CFStringGetMostCompatibleMacStringEncoding(encoding CFStringEncoding, ) CFStringEncoding

// CFStringGetNameOfEncoding(encoding CFStringEncoding, ) CFStringRef


// CFStringGetParagraphBounds(string CFStringRef, range ,  CFRange, parBeginIndex ,  CFIndex  *, parEndIndex ,  CFIndex  *, contentsEndIndex ,  CFIndex  *, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringGetPascalString(theString CFStringRef, buffer ,  StringPtr, bufferSize ,  CFIndex, encoding ,  CFStringEncoding, ) Boolean

// CFStringGetPascalStringPtr(theString CFStringRef, encoding ,  CFStringEncoding, ) ConstStringPtr


// CFStringGetRangeOfComposedCharactersAtIndex(theString CFStringRef, theIndex ,  CFIndex, ) CFRange

// CFStringGetSmallestEncoding(theString CFStringRef, ) CFStringEncoding

// CFStringHasPrefix(theString CFStringRef, prefix ,  CFStringRef, ) Boolean


// CFStringHasSuffix(theString CFStringRef, suffix ,  CFStringRef, ) Boolean

// CFStringInsert(str CFMutableStringRef, idx ,  CFIndex, insertedStr ,  CFStringRef, )

// CFStringIsEncodingAvailable(encoding CFStringEncoding, ) Boolean


// CFStringIsHyphenationAvailableForLocale(locale CFLocaleRef, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.3+
//   - iPadOS 4.3+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringLowercase(theString CFMutableStringRef, locale ,  CFLocaleRef, )

// CFStringNormalize(theString CFMutableStringRef, theForm ,  CFStringNormalizationForm, )


// CFStringPad(theString CFMutableStringRef, padString ,  CFStringRef, length ,  CFIndex, indexIntoPad ,  CFIndex, )

// CFStringReplace(theString CFMutableStringRef, range ,  CFRange, replacement ,  CFStringRef, )

// CFStringReplaceAll(theString CFMutableStringRef, replacement ,  CFStringRef, )


// CFStringSetExternalCharactersNoCopy(theString CFMutableStringRef, chars ,  UniChar  *, length ,  CFIndex, capacity ,  CFIndex, )

// CFStringTokenizerAdvanceToNextToken(tokenizer CFStringTokenizerRef, ) CFStringTokenizerTokenType
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTokenizerCopyBestStringLanguage(string CFStringRef, range ,  CFRange, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFStringTokenizerCopyCurrentTokenAttribute(tokenizer CFStringTokenizerRef, attribute ,  CFOptionFlags, ) CFTypeRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTokenizerCreate(alloc CFAllocatorRef, string ,  CFStringRef, range ,  CFRange, options ,  CFOptionFlags, locale ,  CFLocaleRef, ) CFStringTokenizerRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTokenizerGetCurrentSubTokens(tokenizer CFStringTokenizerRef, ranges ,  CFRange  *, maxRangeLength ,  CFIndex, derivedSubTokens ,  CFMutableArrayRef, ) CFIndex
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFStringTokenizerGetCurrentTokenRange(tokenizer CFStringTokenizerRef, ) CFRange
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTokenizerGoToTokenAtIndex(tokenizer CFStringTokenizerRef, index ,  CFIndex, ) CFStringTokenizerTokenType
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTokenizerSetString(tokenizer CFStringTokenizerRef, string ,  CFStringRef, range ,  CFRange, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFStringTransform(string CFMutableStringRef, range ,  CFRange  *, transform ,  CFStringRef, reverse ,  Boolean, ) Boolean

// CFStringTrim(theString CFMutableStringRef, trimString ,  CFStringRef, )

// CFStringTrimWhitespace(theString CFMutableStringRef, )


// CFStringUppercase(theString CFMutableStringRef, locale ,  CFLocaleRef, )

// CFTimeZoneCopyAbbreviation(tz CFTimeZoneRef, at ,  CFAbsoluteTime, ) CFStringRef

// CFTimeZoneCopyLocalizedName(tz CFTimeZoneRef, style ,  CFTimeZoneNameStyle, locale ,  CFLocaleRef, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFTimeZoneCreate(allocator CFAllocatorRef, name ,  CFStringRef, data ,  CFDataRef, ) CFTimeZoneRef

// CFTimeZoneCreateWithName(allocator CFAllocatorRef, name ,  CFStringRef, tryAbbrev ,  Boolean, ) CFTimeZoneRef

// CFTimeZoneCreateWithTimeIntervalFromGMT(allocator CFAllocatorRef, ti ,  CFTimeInterval, ) CFTimeZoneRef


// CFTimeZoneGetData(tz CFTimeZoneRef, ) CFDataRef

// CFTimeZoneGetDaylightSavingTimeOffset(tz CFTimeZoneRef, at ,  CFAbsoluteTime, ) CFTimeInterval
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFTimeZoneGetName(tz CFTimeZoneRef, ) CFStringRef


// CFTimeZoneGetNextDaylightSavingTimeTransition(tz CFTimeZoneRef, at ,  CFAbsoluteTime, ) CFAbsoluteTime
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFTimeZoneGetSecondsFromGMT(tz CFTimeZoneRef, at ,  CFAbsoluteTime, ) CFTimeInterval

// CFTimeZoneIsDaylightSavingTime(tz CFTimeZoneRef, at ,  CFAbsoluteTime, ) Boolean


// CFTimeZoneSetAbbreviationDictionary(dict CFDictionaryRef, )

// CFTimeZoneSetDefault(tz CFTimeZoneRef, )

// CFTreeAppendChild(tree CFTreeRef, newChild ,  CFTreeRef, )


// CFTreeApplyFunctionToChildren(tree CFTreeRef, applier ,  CFTreeApplierFunction, context ,  void  *, )

// CFTreeCreate(allocator CFAllocatorRef, context ,  const CFTreeContext  *, ) CFTreeRef

// CFTreeFindRoot(tree CFTreeRef, ) CFTreeRef


// CFTreeGetChildAtIndex(tree CFTreeRef, idx ,  CFIndex, ) CFTreeRef

// CFTreeGetChildCount(tree CFTreeRef, ) CFIndex

// CFTreeGetChildren(tree CFTreeRef, children ,  CFTreeRef  *, )


// CFTreeGetContext(tree CFTreeRef, context ,  CFTreeContext  *, )

// CFTreeGetFirstChild(tree CFTreeRef, ) CFTreeRef

// CFTreeGetNextSibling(tree CFTreeRef, ) CFTreeRef


// CFTreeGetParent(tree CFTreeRef, ) CFTreeRef

// CFTreeInsertSibling(tree CFTreeRef, newSibling ,  CFTreeRef, )

// CFTreePrependChild(tree CFTreeRef, newChild ,  CFTreeRef, )


// CFTreeRemove(tree CFTreeRef, )

// CFTreeRemoveAllChildren(tree CFTreeRef, )

// CFTreeSetContext(tree CFTreeRef, context ,  const CFTreeContext  *, )


// CFTreeSortChildren(tree CFTreeRef, comparator ,  CFComparatorFunction, context ,  void  *, )

// CFURLCanBeDecomposed(anURL CFURLRef, ) Boolean

// CFURLClearResourcePropertyCache(url CFURLRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLClearResourcePropertyCacheForKey(url CFURLRef, key ,  CFStringRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCopyAbsoluteURL(relativeURL CFURLRef, ) CFURLRef

// CFURLCopyFileSystemPath(anURL CFURLRef, pathStyle ,  CFURLPathStyle, ) CFStringRef


// CFURLCopyFragment(anURL CFURLRef, charactersToLeaveEscaped ,  CFStringRef, ) CFStringRef

// CFURLCopyHostName(anURL CFURLRef, ) CFStringRef

// CFURLCopyLastPathComponent(url CFURLRef, ) CFStringRef


// CFURLCopyNetLocation(anURL CFURLRef, ) CFStringRef

// CFURLCopyParameterString(anURL CFURLRef, charactersToLeaveEscaped ,  CFStringRef, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 13.0)
//   - iPadOS 2.0+ (Deprecated in 13.0)
//   - macOS 10.2+ (Deprecated in 10.15)
//   - tvOS 9.0+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 6.0)
//
// Deprecated: This function is deprecated.

// CFURLCopyPassword(anURL CFURLRef, ) CFStringRef


// CFURLCopyPath(anURL CFURLRef, ) CFStringRef

// CFURLCopyPathExtension(url CFURLRef, ) CFStringRef

// CFURLCopyQueryString(anURL CFURLRef, charactersToLeaveEscaped ,  CFStringRef, ) CFStringRef


// CFURLCopyResourcePropertiesForKeys(url CFURLRef, keys ,  CFArrayRef, error ,  CFErrorRef  *, ) CFDictionaryRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCopyResourcePropertyForKey(url CFURLRef, key ,  CFStringRef, propertyValueTypeRefPtr ,  void  *, error ,  CFErrorRef  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCopyResourceSpecifier(anURL CFURLRef, ) CFStringRef


// CFURLCopyScheme(anURL CFURLRef, ) CFStringRef

// CFURLCopyStrictPath(anURL CFURLRef, isAbsolute ,  Boolean  *, ) CFStringRef

// CFURLCopyUserName(anURL CFURLRef, ) CFStringRef


// CFURLCreateAbsoluteURLWithBytes(alloc CFAllocatorRef, relativeURLBytes ,  const UInt8  *, length ,  CFIndex, encoding ,  CFStringEncoding, baseURL ,  CFURLRef, useCompatibilityMode ,  Boolean, ) CFURLRef

// CFURLCreateBookmarkData(allocator CFAllocatorRef, url ,  CFURLRef, options ,  CFURLBookmarkCreationOptions, resourcePropertiesToInclude ,  CFArrayRef, relativeToURL ,  CFURLRef, error ,  CFErrorRef  *, ) CFDataRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCreateBookmarkDataFromAliasRecord(allocatorRef CFAllocatorRef, aliasRecordDataRef ,  CFDataRef, ) CFDataRef
//
// Availability:
//   - macOS 10.6+ (Deprecated in 11.0)
//
// Deprecated: This function is deprecated.


// CFURLCreateBookmarkDataFromFile(allocator CFAllocatorRef, fileURL ,  CFURLRef, errorRef ,  CFErrorRef  *, ) CFDataRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCreateByResolvingBookmarkData(allocator CFAllocatorRef, bookmark ,  CFDataRef, options ,  CFURLBookmarkResolutionOptions, relativeToURL ,  CFURLRef, resourcePropertiesToInclude ,  CFArrayRef, isStale ,  Boolean  *, error ,  CFErrorRef  *, ) CFURLRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCreateCopyAppendingPathComponent(allocator CFAllocatorRef, url ,  CFURLRef, pathComponent ,  CFStringRef, isDirectory ,  Boolean, ) CFURLRef


// CFURLCreateCopyAppendingPathExtension(allocator CFAllocatorRef, url ,  CFURLRef, extension ,  CFStringRef, ) CFURLRef

// CFURLCreateCopyDeletingLastPathComponent(allocator CFAllocatorRef, url ,  CFURLRef, ) CFURLRef

// CFURLCreateCopyDeletingPathExtension(allocator CFAllocatorRef, url ,  CFURLRef, ) CFURLRef


// CFURLCreateData(allocator CFAllocatorRef, url ,  CFURLRef, encoding ,  CFStringEncoding, escapeWhitespace ,  Boolean, ) CFDataRef

// CFURLCreateDataAndPropertiesFromResource(alloc CFAllocatorRef, url ,  CFURLRef, resourceData ,  CFDataRef  *, properties ,  CFDictionaryRef  *, desiredProperties ,  CFArrayRef, errorCode ,  SInt32  *, ) Boolean
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLCreateFilePathURL(allocator CFAllocatorRef, url ,  CFURLRef, error ,  CFErrorRef  *, ) CFURLRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLCreateFileReferenceURL(allocator CFAllocatorRef, url ,  CFURLRef, error ,  CFErrorRef  *, ) CFURLRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCreateFromFSRef(allocator CFAllocatorRef, fsRef ,  const struct FSRef  *, ) CFURLRef
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLCreateFromFileSystemRepresentation(allocator CFAllocatorRef, buffer ,  const UInt8  *, bufLen ,  CFIndex, isDirectory ,  Boolean, ) CFURLRef


// CFURLCreateFromFileSystemRepresentationRelativeToBase(allocator CFAllocatorRef, buffer ,  const UInt8  *, bufLen ,  CFIndex, isDirectory ,  Boolean, baseURL ,  CFURLRef, ) CFURLRef

// CFURLCreatePropertyFromResource(alloc CFAllocatorRef, url ,  CFURLRef, property ,  CFStringRef, errorCode ,  SInt32  *, ) CFTypeRef
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator CFAllocatorRef, resourcePropertiesToReturn ,  CFArrayRef, bookmark ,  CFDataRef, ) CFDictionaryRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLCreateResourcePropertyForKeyFromBookmarkData(allocator CFAllocatorRef, resourcePropertyKey ,  CFStringRef, bookmark ,  CFDataRef, ) CFTypeRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCreateStringByAddingPercentEscapes(allocator CFAllocatorRef, originalString ,  CFStringRef, charactersToLeaveUnescaped ,  CFStringRef, legalURLCharactersToBeEscaped ,  CFStringRef, encoding ,  CFStringEncoding, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 9.0)
//   - iPadOS 2.0+ (Deprecated in 9.0)
//   - macOS 10.0+ (Deprecated in 10.11)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLCreateStringByReplacingPercentEscapes(allocator CFAllocatorRef, originalString ,  CFStringRef, charactersToLeaveEscaped ,  CFStringRef, ) CFStringRef


// CFURLCreateStringByReplacingPercentEscapesUsingEncoding(allocator CFAllocatorRef, origString ,  CFStringRef, charsToLeaveEscaped ,  CFStringRef, encoding ,  CFStringEncoding, ) CFStringRef
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 2.0+ (Deprecated in 9.0)
//   - iPadOS 2.0+ (Deprecated in 9.0)
//   - macOS 10.0+ (Deprecated in 10.11)
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLCreateWithBytes(allocator CFAllocatorRef, URLBytes ,  const UInt8  *, length ,  CFIndex, encoding ,  CFStringEncoding, baseURL ,  CFURLRef, ) CFURLRef

// CFURLCreateWithFileSystemPath(allocator CFAllocatorRef, filePath ,  CFStringRef, pathStyle ,  CFURLPathStyle, isDirectory ,  Boolean, ) CFURLRef


// CFURLCreateWithFileSystemPathRelativeToBase(allocator CFAllocatorRef, filePath ,  CFStringRef, pathStyle ,  CFURLPathStyle, isDirectory ,  Boolean, baseURL ,  CFURLRef, ) CFURLRef

// CFURLCreateWithString(allocator CFAllocatorRef, URLString ,  CFStringRef, baseURL ,  CFURLRef, ) CFURLRef

// CFURLDestroyResource(url CFURLRef, errorCode ,  SInt32  *, ) Boolean
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.


// CFURLEnumeratorCreateForDirectoryURL(alloc CFAllocatorRef, directoryURL ,  CFURLRef, option ,  CFURLEnumeratorOptions, propertyKeys ,  CFArrayRef, ) CFURLEnumeratorRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLEnumeratorCreateForMountedVolumes(alloc CFAllocatorRef, option ,  CFURLEnumeratorOptions, propertyKeys ,  CFArrayRef, ) CFURLEnumeratorRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLEnumeratorGetDescendentLevel(enumerator CFURLEnumeratorRef, ) CFIndex
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLEnumeratorGetNextURL(enumerator CFURLEnumeratorRef, url ,  CFURLRef  *, error ,  CFErrorRef  *, ) CFURLEnumeratorResult
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLEnumeratorGetSourceDidChange(enumerator CFURLEnumeratorRef, ) Boolean
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLEnumeratorSkipDescendents(enumerator CFURLEnumeratorRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLGetBaseURL(anURL CFURLRef, ) CFURLRef

// CFURLGetByteRangeForComponent(url CFURLRef, component ,  CFURLComponentType, rangeIncludingSeparators ,  CFRange  *, ) CFRange

// CFURLGetBytes(url CFURLRef, buffer ,  UInt8  *, bufferLength ,  CFIndex, ) CFIndex


// CFURLGetFSRef(url CFURLRef, fsRef ,  struct FSRef  *, ) Boolean
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLGetFileSystemRepresentation(url CFURLRef, resolveAgainstBase ,  Boolean, buffer ,  UInt8  *, maxBufLen ,  CFIndex, ) Boolean

// CFURLGetPortNumber(anURL CFURLRef, ) SInt32


// CFURLGetString(anURL CFURLRef, ) CFStringRef

// CFURLHasDirectoryPath(anURL CFURLRef, ) Boolean

// CFURLIsFileReferenceURL(url CFURLRef, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLResourceIsReachable(url CFURLRef, error ,  CFErrorRef  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLSetResourcePropertiesForKeys(url CFURLRef, keyedPropertyValues ,  CFDictionaryRef, error ,  CFErrorRef  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLSetResourcePropertyForKey(url CFURLRef, key ,  CFStringRef, propertyValue ,  CFTypeRef, error ,  CFErrorRef  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLSetTemporaryResourcePropertyForKey(url CFURLRef, key ,  CFStringRef, propertyValue ,  CFTypeRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLStartAccessingSecurityScopedResource(url CFURLRef, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLStopAccessingSecurityScopedResource(url CFURLRef, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLWriteBookmarkDataToFile(bookmarkRef CFDataRef, fileURL ,  CFURLRef, options ,  CFURLBookmarkFileCreationOptions, errorRef ,  CFErrorRef  *, ) Boolean
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLWriteDataAndPropertiesToResource(url CFURLRef, dataToWrite ,  CFDataRef, propertiesToWrite ,  CFDictionaryRef, errorCode ,  SInt32  *, ) Boolean
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFUUIDCreate(alloc CFAllocatorRef, ) CFUUIDRef


// CFUUIDCreateFromString(alloc CFAllocatorRef, uuidStr ,  CFStringRef, ) CFUUIDRef

// CFUUIDCreateFromUUIDBytes(alloc CFAllocatorRef, bytes ,  CFUUIDBytes, ) CFUUIDRef

// CFUUIDCreateString(alloc CFAllocatorRef, uuid ,  CFUUIDRef, ) CFStringRef


// CFUUIDCreateWithBytes(alloc CFAllocatorRef, byte0 ,  UInt8, byte1 ,  UInt8, byte2 ,  UInt8, byte3 ,  UInt8, byte4 ,  UInt8, byte5 ,  UInt8, byte6 ,  UInt8, byte7 ,  UInt8, byte8 ,  UInt8, byte9 ,  UInt8, byte10 ,  UInt8, byte11 ,  UInt8, byte12 ,  UInt8, byte13 ,  UInt8, byte14 ,  UInt8, byte15 ,  UInt8, ) CFUUIDRef

// CFUUIDGetConstantUUIDWithBytes(alloc CFAllocatorRef, byte0 ,  UInt8, byte1 ,  UInt8, byte2 ,  UInt8, byte3 ,  UInt8, byte4 ,  UInt8, byte5 ,  UInt8, byte6 ,  UInt8, byte7 ,  UInt8, byte8 ,  UInt8, byte9 ,  UInt8, byte10 ,  UInt8, byte11 ,  UInt8, byte12 ,  UInt8, byte13 ,  UInt8, byte14 ,  UInt8, byte15 ,  UInt8, ) CFUUIDRef

// CFUUIDGetUUIDBytes(uuid CFUUIDRef, ) CFUUIDBytes


// CFUserNotificationCancel(userNotification CFUserNotificationRef, ) SInt32
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationCreate(allocator CFAllocatorRef, timeout ,  CFTimeInterval, flags ,  CFOptionFlags, error ,  SInt32  *, dictionary ,  CFDictionaryRef, ) CFUserNotificationRef
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationCreateRunLoopSource(allocator CFAllocatorRef, userNotification ,  CFUserNotificationRef, callout ,  CFUserNotificationCallBack, order ,  CFIndex, ) CFRunLoopSourceRef
//
// Availability:
//   - macOS 10.0+


// CFUserNotificationDisplayAlert(timeout CFTimeInterval, flags ,  CFOptionFlags, iconURL ,  CFURLRef, soundURL ,  CFURLRef, localizationURL ,  CFURLRef, alertHeader ,  CFStringRef, alertMessage ,  CFStringRef, defaultButtonTitle ,  CFStringRef, alternateButtonTitle ,  CFStringRef, otherButtonTitle ,  CFStringRef, responseFlags ,  CFOptionFlags  *, ) SInt32
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationDisplayNotice(timeout CFTimeInterval, flags ,  CFOptionFlags, iconURL ,  CFURLRef, soundURL ,  CFURLRef, localizationURL ,  CFURLRef, alertHeader ,  CFStringRef, alertMessage ,  CFStringRef, defaultButtonTitle ,  CFStringRef, ) SInt32
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationGetResponseDictionary(userNotification CFUserNotificationRef, ) CFDictionaryRef
//
// Availability:
//   - macOS 10.0+


// CFUserNotificationGetResponseValue(userNotification CFUserNotificationRef, key ,  CFStringRef, idx ,  CFIndex, ) CFStringRef
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationReceiveResponse(userNotification CFUserNotificationRef, timeout ,  CFTimeInterval, responseFlags ,  CFOptionFlags  *, ) SInt32
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationUpdate(userNotification CFUserNotificationRef, timeout ,  CFTimeInterval, flags ,  CFOptionFlags, dictionary ,  CFDictionaryRef, ) SInt32
//
// Availability:
//   - macOS 10.0+


// CFWriteStreamCanAcceptBytes(stream CFWriteStreamRef, ) Boolean

// CFWriteStreamClose(stream CFWriteStreamRef, )

// CFWriteStreamCopyDispatchQueue(stream CFWriteStreamRef, ) dispatch_queue_t
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFWriteStreamCopyError(stream CFWriteStreamRef, ) CFErrorRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFWriteStreamCopyProperty(stream CFWriteStreamRef, propertyName ,  CFStreamPropertyKey, ) CFTypeRef

// CFWriteStreamCreateWithAllocatedBuffers(alloc CFAllocatorRef, bufferAllocator ,  CFAllocatorRef, ) CFWriteStreamRef


// CFWriteStreamCreateWithBuffer(alloc CFAllocatorRef, buffer ,  UInt8  *, bufferCapacity ,  CFIndex, ) CFWriteStreamRef

// CFWriteStreamCreateWithFile(alloc CFAllocatorRef, fileURL ,  CFURLRef, ) CFWriteStreamRef

// CFWriteStreamGetError(stream CFWriteStreamRef, ) CFStreamError


// CFWriteStreamGetStatus(stream CFWriteStreamRef, ) CFStreamStatus

// CFWriteStreamOpen(stream CFWriteStreamRef, ) Boolean

// CFWriteStreamScheduleWithRunLoop(stream CFWriteStreamRef, runLoop ,  CFRunLoopRef, runLoopMode ,  CFRunLoopMode, )


// CFWriteStreamSetClient(stream CFWriteStreamRef, streamEvents ,  CFOptionFlags, clientCB ,  CFWriteStreamClientCallBack, clientContext ,  CFStreamClientContext  *, ) Boolean

// CFWriteStreamSetDispatchQueue(stream CFWriteStreamRef, q ,  dispatch_queue_t, )
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFWriteStreamSetProperty(stream CFWriteStreamRef, propertyName ,  CFStreamPropertyKey, propertyValue ,  CFTypeRef, ) Boolean


// CFWriteStreamUnscheduleFromRunLoop(stream CFWriteStreamRef, runLoop ,  CFRunLoopRef, runLoopMode ,  CFRunLoopMode, )

// CFWriteStreamWrite(stream CFWriteStreamRef, buffer ,  const UInt8  *, bufferLength ,  CFIndex, ) CFIndex

// CFXMLCreateStringByEscapingEntities(allocator CFAllocatorRef, string ,  CFStringRef, entitiesDictionary ,  CFDictionaryRef, ) CFStringRef


// CFXMLCreateStringByUnescapingEntities(allocator CFAllocatorRef, string ,  CFStringRef, entitiesDictionary ,  CFDictionaryRef, ) CFStringRef

// CFXMLNodeCreate(alloc CFAllocatorRef, xmlType ,  CFXMLNodeTypeCode, dataString ,  CFStringRef, additionalInfoPtr ,  const void  *, version ,  CFIndex, ) CFXMLNodeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLNodeCreateCopy(alloc CFAllocatorRef, origNode ,  CFXMLNodeRef, ) CFXMLNodeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLNodeGetInfoPtr(node CFXMLNodeRef, ) const void  *
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLNodeGetString(node CFXMLNodeRef, ) CFStringRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLNodeGetTypeCode(node CFXMLNodeRef, ) CFXMLNodeTypeCode
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLNodeGetVersion(node CFXMLNodeRef, ) CFIndex
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserAbort(parser CFXMLParserRef, errorCode ,  CFXMLParserStatusCode, errorDescription ,  CFStringRef, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserCopyErrorDescription(parser CFXMLParserRef, ) CFStringRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLParserCreate(allocator CFAllocatorRef, xmlData ,  CFDataRef, dataSource ,  CFURLRef, parseOptions ,  CFOptionFlags, versionOfNodes ,  CFIndex, callBacks ,  CFXMLParserCallBacks  *, context ,  CFXMLParserContext  *, ) CFXMLParserRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserCreateWithDataFromURL(allocator CFAllocatorRef, dataSource ,  CFURLRef, parseOptions ,  CFOptionFlags, versionOfNodes ,  CFIndex, callBacks ,  CFXMLParserCallBacks  *, context ,  CFXMLParserContext  *, ) CFXMLParserRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserGetCallBacks(parser CFXMLParserRef, callBacks ,  CFXMLParserCallBacks  *, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLParserGetContext(parser CFXMLParserRef, context ,  CFXMLParserContext  *, )
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserGetDocument(parser CFXMLParserRef, ) void  *
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserGetLineNumber(parser CFXMLParserRef, ) CFIndex
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLParserGetLocation(parser CFXMLParserRef, ) CFIndex
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserGetSourceURL(parser CFXMLParserRef, ) CFURLRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserGetStatusCode(parser CFXMLParserRef, ) CFXMLParserStatusCode
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLParserParse(parser CFXMLParserRef, ) Boolean
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLTreeCreateFromData(allocator CFAllocatorRef, xmlData ,  CFDataRef, dataSource ,  CFURLRef, parseOptions ,  CFOptionFlags, versionOfNodes ,  CFIndex, ) CFXMLTreeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLTreeCreateFromDataWithError(allocator CFAllocatorRef, xmlData ,  CFDataRef, dataSource ,  CFURLRef, parseOptions ,  CFOptionFlags, versionOfNodes ,  CFIndex, errorDict ,  CFDictionaryRef  *, ) CFXMLTreeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLTreeCreateWithDataFromURL(allocator CFAllocatorRef, dataSource ,  CFURLRef, parseOptions ,  CFOptionFlags, versionOfNodes ,  CFIndex, ) CFXMLTreeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLTreeCreateWithNode(allocator CFAllocatorRef, node ,  CFXMLNodeRef, ) CFXMLTreeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLTreeCreateXMLData(allocator CFAllocatorRef, xmlTree ,  CFXMLTreeRef, ) CFDataRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLTreeGetNode(xmlTree CFXMLTreeRef, ) CFXMLNodeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

