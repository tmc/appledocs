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

// Discovered functions (854 total):

// CFRelease(CFTypeRef cf);)
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFRetain(CFTypeRef cf);) CFTypeRef
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFAbsoluteTimeAddGregorianUnits(at _, tz :  CFAbsoluteTime,  _, units :  CFTimeZone!,  _, :  CFGregorianUnits) ->  CFAbsoluteTime) func
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


// CFAbsoluteTimeGetCurrent() func

// CFAbsoluteTimeGetDayOfWeek(at _, tz :  CFAbsoluteTime,  _, :  CFTimeZone!) ->  Int32) func
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

// CFAbsoluteTimeGetDayOfYear(at _, tz :  CFAbsoluteTime,  _, :  CFTimeZone!) ->  Int32) func
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


// CFAbsoluteTimeGetDifferenceAsGregorianUnits(at1 _, at2 :  CFAbsoluteTime,  _, tz :  CFAbsoluteTime,  _, unitFlags :  CFTimeZone!,  _, :  CFOptionFlags) ->  CFGregorianUnits) func
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

// CFAbsoluteTimeGetGregorianDate(at _, tz :  CFAbsoluteTime,  _, :  CFTimeZone!) ->  CFGregorianDate) func
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

// CFAbsoluteTimeGetWeekOfYear(at _, tz :  CFAbsoluteTime,  _, :  CFTimeZone!) ->  Int32) func
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


// CFAllocatorAllocate(allocator _, size :  CFAllocator!,  _, hint :  CFIndex,  _, :  CFOptionFlags) ->  UnsafeMutableRawPointer!) func

// CFAllocatorAllocateBytes(allocator _, size :  CFAllocator!,  _, hint :  CFIndex,  _, :  CFOptionFlags) ->  UnsafeMutableRawPointer!) func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+

// CFAllocatorAllocateTyped(allocator _, size :  CFAllocator!,  _, descriptor :  CFIndex,  _, hint :  CFAllocatorTypeID,  _, :  CFOptionFlags) ->  UnsafeMutableRawPointer!) func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+


// CFAllocatorCreate(allocator _, context :  CFAllocator!,  _, :  UnsafeMutablePointer< CFAllocatorContext>!) ->  Unmanaged< CFAllocator>!) func

// CFAllocatorCreateWithZone(allocator CFAllocatorRef, zone ,  struct  _malloc_zone_t *, );) extern   CFAllocatorRef

// CFAllocatorDeallocate(allocator _, ptr :  CFAllocator!,  _, :  UnsafeMutableRawPointer!)) func


// CFAllocatorGetContext(allocator _, context :  CFAllocator!,  _, :  UnsafeMutablePointer< CFAllocatorContext>!)) func

// CFAllocatorGetDefault() func

// CFAllocatorGetPreferredSizeForSize(allocator _, size :  CFAllocator!,  _, hint :  CFIndex,  _, :  CFOptionFlags) ->  CFIndex) func


// CFAllocatorGetTypeID() func

// CFAllocatorReallocate(allocator _, ptr :  CFAllocator!,  _, newsize :  UnsafeMutableRawPointer!,  _, hint :  CFIndex,  _, :  CFOptionFlags) ->  UnsafeMutableRawPointer!) func

// CFAllocatorReallocateBytes(allocator _, ptr :  CFAllocator!,  _, newsize :  UnsafeMutableRawPointer!,  _, hint :  CFIndex,  _, :  CFOptionFlags) ->  UnsafeMutableRawPointer!) func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+


// CFAllocatorReallocateTyped(allocator _, ptr :  CFAllocator!,  _, newsize :  UnsafeMutableRawPointer!,  _, descriptor :  CFIndex,  _, hint :  CFAllocatorTypeID,  _, :  CFOptionFlags) ->  UnsafeMutableRawPointer!) func
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//   - visionOS 2.0+
//   - watchOS 11.0+

// CFAllocatorSetDefault(allocator _, :  CFAllocator!)) func

// CFArrayAppendArray(theArray _, otherArray :  CFMutableArray!,  _, otherRange :  CFArray!,  _, :  CFRange) func


// CFArrayAppendValue(theArray _, value :  CFMutableArray!,  _, :  UnsafeRawPointer!)) func

// CFArrayApplyFunction(theArray _, range :  CFArray!,  _, applier :  CFRange,  _, context : (( UnsafeRawPointer?,  UnsafeMutableRawPointer?) ->  Void)!,  _, :  UnsafeMutableRawPointer!)) func

// CFArrayBSearchValues(theArray _, range :  CFArray!,  _, value :  CFRange,  _, comparator :  UnsafeRawPointer!,  _, context :  CFComparatorFunction!,  _, :  UnsafeMutableRawPointer!) ->  CFIndex) func


// CFArrayContainsValue(theArray _, range :  CFArray!,  _, value :  CFRange,  _, :  UnsafeRawPointer!) ->  Bool) func

// CFArrayCreate(allocator _, values :  CFAllocator!,  _, numValues :  UnsafeMutablePointer< UnsafeRawPointer?>!,  _, callBacks :  CFIndex,  _, :  UnsafePointer< CFArrayCallBacks>!) ->  CFArray!) func

// CFArrayCreateCopy(allocator _, theArray :  CFAllocator!,  _, :  CFArray!) ->  CFArray!) func


// CFArrayCreateMutable(allocator _, capacity :  CFAllocator!,  _, callBacks :  CFIndex,  _, :  UnsafePointer< CFArrayCallBacks>!) ->  CFMutableArray!) func

// CFArrayCreateMutableCopy(allocator _, capacity :  CFAllocator!,  _, theArray :  CFIndex,  _, :  CFArray!) ->  CFMutableArray!) func

// CFArrayExchangeValuesAtIndices(theArray _, idx1 :  CFMutableArray!,  _, idx2 :  CFIndex,  _, :  CFIndex) func


// CFArrayGetCount(theArray _, :  CFArray!) ->  CFIndex) func

// CFArrayGetCountOfValue(theArray _, range :  CFArray!,  _, value :  CFRange,  _, :  UnsafeRawPointer!) ->  CFIndex) func

// CFArrayGetFirstIndexOfValue(theArray _, range :  CFArray!,  _, value :  CFRange,  _, :  UnsafeRawPointer!) ->  CFIndex) func


// CFArrayGetLastIndexOfValue(theArray _, range :  CFArray!,  _, value :  CFRange,  _, :  UnsafeRawPointer!) ->  CFIndex) func

// CFArrayGetTypeID() func

// CFArrayGetValueAtIndex(theArray _, idx :  CFArray!,  _, :  CFIndex) ->  UnsafeRawPointer!) func


// CFArrayGetValues(theArray _, range :  CFArray!,  _, values :  CFRange,  _, :  UnsafeMutablePointer< UnsafeRawPointer?>!)) func

// CFArrayInsertValueAtIndex(theArray _, idx :  CFMutableArray!,  _, value :  CFIndex,  _, :  UnsafeRawPointer!)) func

// CFArrayRemoveAllValues(theArray _, :  CFMutableArray!)) func


// CFArrayRemoveValueAtIndex(theArray _, idx :  CFMutableArray!,  _, :  CFIndex) func

// CFArrayReplaceValues(theArray _, range :  CFMutableArray!,  _, newValues :  CFRange,  _, newCount :  UnsafeMutablePointer< UnsafeRawPointer?>!,  _, :  CFIndex) func

// CFArraySetValueAtIndex(theArray _, idx :  CFMutableArray!,  _, value :  CFIndex,  _, :  UnsafeRawPointer!)) func


// CFArraySortValues(theArray _, range :  CFMutableArray!,  _, comparator :  CFRange,  _, context :  CFComparatorFunction!,  _, :  UnsafeMutableRawPointer!)) func

// CFAttributedStringBeginEditing(aStr _, :  CFMutableAttributedString!)) func

// CFAttributedStringCreate(alloc _, str :  CFAllocator!,  _, attributes :  CFString!,  _, :  CFDictionary!) ->  CFAttributedString!) func


// CFAttributedStringCreateCopy(alloc _, aStr :  CFAllocator!,  _, :  CFAttributedString!) ->  CFAttributedString!) func

// CFAttributedStringCreateMutable(alloc _, maxLength :  CFAllocator!,  _, :  CFIndex) ->  CFMutableAttributedString!) func

// CFAttributedStringCreateMutableCopy(alloc _, maxLength :  CFAllocator!,  _, aStr :  CFIndex,  _, :  CFAttributedString!) ->  CFMutableAttributedString!) func


// CFAttributedStringCreateWithSubstring(alloc _, aStr :  CFAllocator!,  _, range :  CFAttributedString!,  _, :  CFRange) ->  CFAttributedString!) func

// CFAttributedStringEndEditing(aStr _, :  CFMutableAttributedString!)) func

// CFAttributedStringGetAttribute(aStr _, loc :  CFAttributedString!,  _, attrName :  CFIndex,  _, effectiveRange :  CFString!,  _, :  UnsafeMutablePointer< CFRange>!) ->  CFTypeRef!) func


// CFAttributedStringGetAttributeAndLongestEffectiveRange(aStr _, loc :  CFAttributedString!,  _, attrName :  CFIndex,  _, inRange :  CFString!,  _, longestEffectiveRange :  CFRange,  _, :  UnsafeMutablePointer< CFRange>!) ->  CFTypeRef!) func

// CFAttributedStringGetAttributes(aStr _, loc :  CFAttributedString!,  _, effectiveRange :  CFIndex,  _, :  UnsafeMutablePointer< CFRange>!) ->  CFDictionary!) func

// CFAttributedStringGetAttributesAndLongestEffectiveRange(aStr _, loc :  CFAttributedString!,  _, inRange :  CFIndex,  _, longestEffectiveRange :  CFRange,  _, :  UnsafeMutablePointer< CFRange>!) ->  CFDictionary!) func


// CFAttributedStringGetBidiLevelsAndResolvedDirections(attributedString _, range :  CFAttributedString!,  _, baseDirection :  CFRange,  _, bidiLevels :  Int8,  _, baseDirections :  UnsafeMutablePointer< UInt8>!,  _, :  UnsafeMutablePointer< UInt8>!) ->  Bool) func

// CFAttributedStringGetLength(aStr _, :  CFAttributedString!) ->  CFIndex) func

// CFAttributedStringGetMutableString(aStr _, :  CFMutableAttributedString!) ->  CFMutableString!) func


// CFAttributedStringGetStatisticalWritingDirections(attributedString _, range :  CFAttributedString!,  _, baseDirection :  CFRange,  _, bidiLevels :  Int8,  _, baseDirections :  UnsafeMutablePointer< UInt8>!,  _, :  UnsafeMutablePointer< UInt8>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 26.0+
//   - iOS 26.0+
//   - iPadOS 26.0+
//   - macOS 26.0+
//   - tvOS 26.0+
//   - visionOS 26.0+
//   - watchOS 26.0+

// CFAttributedStringGetString(aStr _, :  CFAttributedString!) ->  CFString!) func

// CFAttributedStringGetTypeID() func


// CFAttributedStringRemoveAttribute(aStr _, range :  CFMutableAttributedString!,  _, attrName :  CFRange,  _, :  CFString!)) func

// CFAttributedStringReplaceAttributedString(aStr _, range :  CFMutableAttributedString!,  _, replacement :  CFRange,  _, :  CFAttributedString!)) func

// CFAttributedStringReplaceString(aStr _, range :  CFMutableAttributedString!,  _, replacement :  CFRange,  _, :  CFString!)) func


// CFAttributedStringSetAttribute(aStr _, range :  CFMutableAttributedString!,  _, attrName :  CFRange,  _, value :  CFString!,  _, :  CFTypeRef!)) func

// CFAttributedStringSetAttributes(aStr _, range :  CFMutableAttributedString!,  _, replacement :  CFRange,  _, clearOtherAttributes :  CFDictionary!,  _, :  Bool) func

// CFAutorelease(arg CFTypeRef, );) extern   CFTypeRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFBagAddValue(theBag _, value :  CFMutableBag!,  _, :  UnsafeRawPointer!)) func

// CFBagApplyFunction(theBag _, applier :  CFBag!,  _, context : (( UnsafeRawPointer?,  UnsafeMutableRawPointer?) ->  Void)!,  _, :  UnsafeMutableRawPointer!)) func

// CFBagContainsValue(theBag _, value :  CFBag!,  _, :  UnsafeRawPointer!) ->  Bool) func


// CFBagCreate(allocator _, values :  CFAllocator!,  _, numValues :  UnsafeMutablePointer< UnsafeRawPointer?>!,  _, callBacks :  CFIndex,  _, :  UnsafePointer< CFBagCallBacks>!) ->  CFBag!) func

// CFBagCreateCopy(allocator _, theBag :  CFAllocator!,  _, :  CFBag!) ->  CFBag!) func

// CFBagCreateMutable(allocator _, capacity :  CFAllocator!,  _, callBacks :  CFIndex,  _, :  UnsafePointer< CFBagCallBacks>!) ->  CFMutableBag!) func


// CFBagCreateMutableCopy(allocator _, capacity :  CFAllocator!,  _, theBag :  CFIndex,  _, :  CFBag!) ->  CFMutableBag!) func

// CFBagGetCount(theBag _, :  CFBag!) ->  CFIndex) func

// CFBagGetCountOfValue(theBag _, value :  CFBag!,  _, :  UnsafeRawPointer!) ->  CFIndex) func


// CFBagGetTypeID() func

// CFBagGetValue(theBag _, value :  CFBag!,  _, :  UnsafeRawPointer!) ->  UnsafeRawPointer!) func

// CFBagGetValueIfPresent(theBag _, candidate :  CFBag!,  _, value :  UnsafeRawPointer!,  _, :  UnsafeMutablePointer< UnsafeRawPointer?>!) ->  Bool) func


// CFBagGetValues(theBag _, values :  CFBag!,  _, :  UnsafeMutablePointer< UnsafeRawPointer?>!)) func

// CFBagRemoveAllValues(theBag _, :  CFMutableBag!)) func

// CFBagRemoveValue(theBag _, value :  CFMutableBag!,  _, :  UnsafeRawPointer!)) func


// CFBagReplaceValue(theBag _, value :  CFMutableBag!,  _, :  UnsafeRawPointer!)) func

// CFBagSetValue(theBag _, value :  CFMutableBag!,  _, :  UnsafeRawPointer!)) func

// CFBinaryHeapAddValue(heap _, value :  CFBinaryHeap!,  _, :  UnsafeRawPointer!)) func


// CFBinaryHeapApplyFunction(heap _, applier :  CFBinaryHeap!,  _, context : (( UnsafeRawPointer?,  UnsafeMutableRawPointer?) ->  Void)!,  _, :  UnsafeMutableRawPointer!)) func

// CFBinaryHeapContainsValue(heap _, value :  CFBinaryHeap!,  _, :  UnsafeRawPointer!) ->  Bool) func

// CFBinaryHeapCreate(allocator _, capacity :  CFAllocator!,  _, callBacks :  CFIndex,  _, compareContext :  UnsafePointer< CFBinaryHeapCallBacks>!,  _, :  UnsafePointer< CFBinaryHeapCompareContext>!) ->  CFBinaryHeap!) func


// CFBinaryHeapCreateCopy(allocator _, capacity :  CFAllocator!,  _, heap :  CFIndex,  _, :  CFBinaryHeap!) ->  CFBinaryHeap!) func

// CFBinaryHeapGetCount(heap _, :  CFBinaryHeap!) ->  CFIndex) func

// CFBinaryHeapGetCountOfValue(heap _, value :  CFBinaryHeap!,  _, :  UnsafeRawPointer!) ->  CFIndex) func


// CFBinaryHeapGetMinimum(heap _, :  CFBinaryHeap!) ->  UnsafeRawPointer!) func

// CFBinaryHeapGetMinimumIfPresent(heap _, value :  CFBinaryHeap!,  _, :  UnsafeMutablePointer< UnsafeRawPointer?>!) ->  Bool) func

// CFBinaryHeapGetTypeID() func


// CFBinaryHeapGetValues(heap _, values :  CFBinaryHeap!,  _, :  UnsafeMutablePointer< UnsafeRawPointer?>!)) func

// CFBinaryHeapRemoveAllValues(heap _, :  CFBinaryHeap!)) func

// CFBinaryHeapRemoveMinimumValue(heap _, :  CFBinaryHeap!)) func


// CFBitVectorContainsBit(bv _, range :  CFBitVector!,  _, value :  CFRange,  _, :  CFBit) ->  Bool) func

// CFBitVectorCreate(allocator _, bytes :  CFAllocator!,  _, numBits :  UnsafePointer< UInt8>!,  _, :  CFIndex) ->  CFBitVector!) func

// CFBitVectorCreateCopy(allocator _, bv :  CFAllocator!,  _, :  CFBitVector!) ->  CFBitVector!) func


// CFBitVectorCreateMutable(allocator _, capacity :  CFAllocator!,  _, :  CFIndex) ->  CFMutableBitVector!) func

// CFBitVectorCreateMutableCopy(allocator _, capacity :  CFAllocator!,  _, bv :  CFIndex,  _, :  CFBitVector!) ->  CFMutableBitVector!) func

// CFBitVectorFlipBitAtIndex(bv _, idx :  CFMutableBitVector!,  _, :  CFIndex) func


// CFBitVectorFlipBits(bv _, range :  CFMutableBitVector!,  _, :  CFRange) func

// CFBitVectorGetBitAtIndex(bv _, idx :  CFBitVector!,  _, :  CFIndex) ->  CFBit) func

// CFBitVectorGetBits(bv _, range :  CFBitVector!,  _, bytes :  CFRange,  _, :  UnsafeMutablePointer< UInt8>!)) func


// CFBitVectorGetCount(bv _, :  CFBitVector!) ->  CFIndex) func

// CFBitVectorGetCountOfBit(bv _, range :  CFBitVector!,  _, value :  CFRange,  _, :  CFBit) ->  CFIndex) func

// CFBitVectorGetFirstIndexOfBit(bv _, range :  CFBitVector!,  _, value :  CFRange,  _, :  CFBit) ->  CFIndex) func


// CFBitVectorGetLastIndexOfBit(bv _, range :  CFBitVector!,  _, value :  CFRange,  _, :  CFBit) ->  CFIndex) func

// CFBitVectorGetTypeID() func

// CFBitVectorSetAllBits(bv _, value :  CFMutableBitVector!,  _, :  CFBit) func


// CFBitVectorSetBitAtIndex(bv _, idx :  CFMutableBitVector!,  _, value :  CFIndex,  _, :  CFBit) func

// CFBitVectorSetBits(bv _, range :  CFMutableBitVector!,  _, value :  CFRange,  _, :  CFBit) func

// CFBitVectorSetCount(bv _, count :  CFMutableBitVector!,  _, :  CFIndex) func


// CFBooleanGetTypeID() func

// CFBooleanGetValue(boolean _, :  CFBoolean!) ->  Bool) func

// CFBundleCloseBundleResourceMap(bundle _, refNum :  CFBundle!,  _, :  CFBundleRefNum) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.15)
//
// Deprecated: This function is deprecated.


// CFBundleCopyAuxiliaryExecutableURL(bundle _, executableName :  CFBundle!,  _, :  CFString!) ->  CFURL!) func

// CFBundleCopyBuiltInPlugInsURL(bundle _, :  CFBundle!) ->  CFURL!) func

// CFBundleCopyBundleLocalizations(bundle _, :  CFBundle!) ->  CFArray!) func


// CFBundleCopyBundleURL(bundle _, :  CFBundle!) ->  CFURL!) func

// CFBundleCopyExecutableArchitectures(bundle _, :  CFBundle!) ->  CFArray!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFBundleCopyExecutableArchitecturesForURL(url _, :  CFURL!) ->  CFArray!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFBundleCopyExecutableURL(bundle _, :  CFBundle!) ->  CFURL!) func

// CFBundleCopyInfoDictionaryForURL(url _, :  CFURL!) ->  CFDictionary!) func

// CFBundleCopyInfoDictionaryInDirectory(bundleURL _, :  CFURL!) ->  CFDictionary!) func


// CFBundleCopyLocalizationsForPreferences(locArray _, prefArray :  CFArray!,  _, :  CFArray!) ->  CFArray!) func

// CFBundleCopyLocalizationsForURL(url _, :  CFURL!) ->  CFArray!) func

// CFBundleCopyLocalizedString(bundle _, key :  CFBundle!,  _, value :  CFString!,  _, tableName :  CFString!,  _, :  CFString!) ->  CFString!) func


// CFBundleCopyLocalizedStringForLocalizations(bundle _, key :  CFBundle!,  _, value :  CFString!,  _, tableName :  CFString!,  _, localizations :  CFString!,  _, :  CFArray!) ->  CFString!) func
//
// Availability:
//   - Mac Catalyst 18.4+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 15.4+
//   - tvOS 18.4+
//   - visionOS 2.4+
//   - watchOS 11.4+

// CFBundleCopyPreferredLocalizationsFromArray(locArray _, :  CFArray!) ->  CFArray!) func

// CFBundleCopyPrivateFrameworksURL(bundle _, :  CFBundle!) ->  CFURL!) func


// CFBundleCopyResourceURL(bundle _, resourceName :  CFBundle!,  _, resourceType :  CFString!,  _, subDirName :  CFString!,  _, :  CFString!) ->  CFURL!) func

// CFBundleCopyResourceURLForLocalization(bundle _, resourceName :  CFBundle!,  _, resourceType :  CFString!,  _, subDirName :  CFString!,  _, localizationName :  CFString!,  _, :  CFString!) ->  CFURL!) func

// CFBundleCopyResourceURLInDirectory(bundleURL _, resourceName :  CFURL!,  _, resourceType :  CFString!,  _, subDirName :  CFString!,  _, :  CFString!) ->  CFURL!) func


// CFBundleCopyResourceURLsOfType(bundle _, resourceType :  CFBundle!,  _, subDirName :  CFString!,  _, :  CFString!) ->  CFArray!) func

// CFBundleCopyResourceURLsOfTypeForLocalization(bundle _, resourceType :  CFBundle!,  _, subDirName :  CFString!,  _, localizationName :  CFString!,  _, :  CFString!) ->  CFArray!) func

// CFBundleCopyResourceURLsOfTypeInDirectory(bundleURL _, resourceType :  CFURL!,  _, subDirName :  CFString!,  _, :  CFString!) ->  CFArray!) func


// CFBundleCopyResourcesDirectoryURL(bundle _, :  CFBundle!) ->  CFURL!) func

// CFBundleCopySharedFrameworksURL(bundle _, :  CFBundle!) ->  CFURL!) func

// CFBundleCopySharedSupportURL(bundle _, :  CFBundle!) ->  CFURL!) func


// CFBundleCopySupportFilesDirectoryURL(bundle _, :  CFBundle!) ->  CFURL!) func

// CFBundleCreate(allocator _, bundleURL :  CFAllocator!,  _, :  CFURL!) ->  CFBundle!) func

// CFBundleCreateBundlesFromDirectory(allocator _, directoryURL :  CFAllocator!,  _, bundleType :  CFURL!,  _, :  CFString!) ->  CFArray!) func


// CFBundleGetAllBundles() func

// CFBundleGetBundleWithIdentifier(bundleID _, :  CFString!) ->  CFBundle!) func

// CFBundleGetDataPointerForName(bundle _, symbolName :  CFBundle!,  _, :  CFString!) ->  UnsafeMutableRawPointer!) func


// CFBundleGetDataPointersForNames(bundle _, symbolNames :  CFBundle!,  _, stbl :  CFArray!,  _, :  UnsafeMutablePointer< UnsafeMutableRawPointer?>!)) func

// CFBundleGetDevelopmentRegion(bundle _, :  CFBundle!) ->  CFString!) func

// CFBundleGetFunctionPointerForName(bundle _, functionName :  CFBundle!,  _, :  CFString!) ->  UnsafeMutableRawPointer!) func


// CFBundleGetFunctionPointersForNames(bundle _, functionNames :  CFBundle!,  _, ftbl :  CFArray!,  _, :  UnsafeMutablePointer< UnsafeMutableRawPointer?>!)) func

// CFBundleGetIdentifier(bundle _, :  CFBundle!) ->  CFString!) func

// CFBundleGetInfoDictionary(bundle _, :  CFBundle!) ->  CFDictionary!) func


// CFBundleGetLocalInfoDictionary(bundle _, :  CFBundle!) ->  CFDictionary!) func

// CFBundleGetMainBundle() func

// CFBundleGetPackageInfo(bundle _, packageType :  CFBundle!,  _, packageCreator :  UnsafeMutablePointer< UInt32>!,  _, :  UnsafeMutablePointer< UInt32>!)) func


// CFBundleGetPackageInfoInDirectory(url _, packageType :  CFURL!,  _, packageCreator :  UnsafeMutablePointer< UInt32>!,  _, :  UnsafeMutablePointer< UInt32>!) ->  Bool) func

// CFBundleGetPlugIn(bundle _, :  CFBundle!) ->  CFPlugIn!) func

// CFBundleGetTypeID() func


// CFBundleGetValueForInfoDictionaryKey(bundle _, key :  CFBundle!,  _, :  CFString!) ->  CFTypeRef!) func

// CFBundleGetVersionNumber(bundle _, :  CFBundle!) ->  UInt32) func

// CFBundleIsArchitectureLoadable(arch _, :  cpu_type_t) ->  Bool) func
//
// Availability:
//   - macOS 11.0+


// CFBundleIsExecutableLoadable(bundle _, :  CFBundle!) ->  Bool) func
//
// Availability:
//   - macOS 11.0+

// CFBundleIsExecutableLoadableForURL(url _, :  CFURL!) ->  Bool) func
//
// Availability:
//   - macOS 11.0+

// CFBundleIsExecutableLoaded(bundle _, :  CFBundle!) ->  Bool) func


// CFBundleLoadExecutable(bundle _, :  CFBundle!) ->  Bool) func

// CFBundleLoadExecutableAndReturnError(bundle _, error :  CFBundle!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFBundleOpenBundleResourceFiles(bundle _, refNum :  CFBundle!,  _, localizedRefNum :  UnsafeMutablePointer< CFBundleRefNum>!,  _, :  UnsafeMutablePointer< CFBundleRefNum>!) ->  Int32) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.15)
//
// Deprecated: This function is deprecated.


// CFBundleOpenBundleResourceMap(bundle _, :  CFBundle!) ->  CFBundleRefNum) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.15)
//
// Deprecated: This function is deprecated.

// CFBundlePreflightExecutable(bundle _, error :  CFBundle!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFBundleUnloadExecutable(bundle _, :  CFBundle!)) func


// CFByteOrderGetCurrent() func

// CFCalendarAddComponents(calendar CFCalendarRef, at ,  CFAbsoluteTime *, options ,  CFOptionFlags, componentDesc ,  const  char *, , ...);) extern   Boolean

// CFCalendarComposeAbsoluteTime(calendar CFCalendarRef, at ,  CFAbsoluteTime *, componentDesc ,  const  char *, , ...);) extern   Boolean


// CFCalendarCopyCurrent() func

// CFCalendarCopyLocale(calendar _, :  CFCalendar!) ->  CFLocale!) func

// CFCalendarCopyTimeZone(calendar _, :  CFCalendar!) ->  CFTimeZone!) func


// CFCalendarCreateWithIdentifier(allocator _, identifier :  CFAllocator!,  _, :  CFCalendarIdentifier!) ->  CFCalendar!) func

// CFCalendarDecomposeAbsoluteTime(calendar CFCalendarRef, at ,  CFAbsoluteTime, componentDesc ,  const  char *, , ...);) extern   Boolean

// CFCalendarGetComponentDifference(calendar CFCalendarRef, startingAT ,  CFAbsoluteTime, resultAT ,  CFAbsoluteTime, options ,  CFOptionFlags, componentDesc ,  const  char *, , ...);) extern   Boolean


// CFCalendarGetFirstWeekday(calendar _, :  CFCalendar!) ->  CFIndex) func

// CFCalendarGetIdentifier(calendar _, :  CFCalendar!) ->  CFCalendarIdentifier!) func

// CFCalendarGetMaximumRangeOfUnit(calendar _, unit :  CFCalendar!,  _, :  CFCalendarUnit) ->  CFRange) func


// CFCalendarGetMinimumDaysInFirstWeek(calendar _, :  CFCalendar!) ->  CFIndex) func

// CFCalendarGetMinimumRangeOfUnit(calendar _, unit :  CFCalendar!,  _, :  CFCalendarUnit) ->  CFRange) func

// CFCalendarGetOrdinalityOfUnit(calendar _, smallerUnit :  CFCalendar!,  _, biggerUnit :  CFCalendarUnit,  _, at :  CFCalendarUnit,  _, :  CFAbsoluteTime) ->  CFIndex) func


// CFCalendarGetRangeOfUnit(calendar _, smallerUnit :  CFCalendar!,  _, biggerUnit :  CFCalendarUnit,  _, at :  CFCalendarUnit,  _, :  CFAbsoluteTime) ->  CFRange) func

// CFCalendarGetTimeRangeOfUnit(calendar _, unit :  CFCalendar!,  _, at :  CFCalendarUnit,  _, startp :  CFAbsoluteTime,  _, tip :  UnsafeMutablePointer< CFAbsoluteTime>!,  _, :  UnsafeMutablePointer< CFTimeInterval>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFCalendarGetTypeID() func


// CFCalendarSetFirstWeekday(calendar _, wkdy :  CFCalendar!,  _, :  CFIndex) func

// CFCalendarSetLocale(calendar _, locale :  CFCalendar!,  _, :  CFLocale!)) func

// CFCalendarSetMinimumDaysInFirstWeek(calendar _, mwd :  CFCalendar!,  _, :  CFIndex) func


// CFCalendarSetTimeZone(calendar _, tz :  CFCalendar!,  _, :  CFTimeZone!)) func

// CFCharacterSetAddCharactersInRange(theSet _, theRange :  CFMutableCharacterSet!,  _, :  CFRange) func

// CFCharacterSetAddCharactersInString(theSet _, theString :  CFMutableCharacterSet!,  _, :  CFString!)) func


// CFCharacterSetCreateBitmapRepresentation(alloc _, theSet :  CFAllocator!,  _, :  CFCharacterSet!) ->  CFData!) func

// CFCharacterSetCreateCopy(alloc _, theSet :  CFAllocator!,  _, :  CFCharacterSet!) ->  CFCharacterSet!) func

// CFCharacterSetCreateInvertedSet(alloc _, theSet :  CFAllocator!,  _, :  CFCharacterSet!) ->  CFCharacterSet!) func


// CFCharacterSetCreateMutable(alloc _, :  CFAllocator!) ->  CFMutableCharacterSet!) func

// CFCharacterSetCreateMutableCopy(alloc _, theSet :  CFAllocator!,  _, :  CFCharacterSet!) ->  CFMutableCharacterSet!) func

// CFCharacterSetCreateWithBitmapRepresentation(alloc _, theData :  CFAllocator!,  _, :  CFData!) ->  CFCharacterSet!) func


// CFCharacterSetCreateWithCharactersInRange(alloc _, theRange :  CFAllocator!,  _, :  CFRange) ->  CFCharacterSet!) func

// CFCharacterSetCreateWithCharactersInString(alloc _, theString :  CFAllocator!,  _, :  CFString!) ->  CFCharacterSet!) func

// CFCharacterSetGetPredefined(theSetIdentifier _, :  CFCharacterSetPredefinedSet) ->  CFCharacterSet!) func


// CFCharacterSetGetTypeID() func

// CFCharacterSetHasMemberInPlane(theSet _, thePlane :  CFCharacterSet!,  _, :  CFIndex) ->  Bool) func

// CFCharacterSetIntersect(theSet _, theOtherSet :  CFMutableCharacterSet!,  _, :  CFCharacterSet!)) func


// CFCharacterSetInvert(theSet _, :  CFMutableCharacterSet!)) func

// CFCharacterSetIsCharacterMember(theSet _, theChar :  CFCharacterSet!,  _, :  UniChar) ->  Bool) func

// CFCharacterSetIsLongCharacterMember(theSet _, theChar :  CFCharacterSet!,  _, :  UTF32Char) ->  Bool) func


// CFCharacterSetIsSupersetOfSet(theSet _, theOtherset :  CFCharacterSet!,  _, :  CFCharacterSet!) ->  Bool) func

// CFCharacterSetRemoveCharactersInRange(theSet _, theRange :  CFMutableCharacterSet!,  _, :  CFRange) func

// CFCharacterSetRemoveCharactersInString(theSet _, theString :  CFMutableCharacterSet!,  _, :  CFString!)) func


// CFCharacterSetUnion(theSet _, theOtherSet :  CFMutableCharacterSet!,  _, :  CFCharacterSet!)) func

// CFConvertDoubleHostToSwapped(arg _, :  Double) ->  CFSwappedFloat64) func

// CFConvertDoubleSwappedToHost(arg _, :  CFSwappedFloat64) ->  Double) func


// CFConvertFloat32HostToSwapped(arg _, :  Float32) ->  CFSwappedFloat32) func

// CFConvertFloat32SwappedToHost(arg _, :  CFSwappedFloat32) ->  Float32) func

// CFConvertFloat64HostToSwapped(arg _, :  Float64) ->  CFSwappedFloat64) func


// CFConvertFloat64SwappedToHost(arg _, :  CFSwappedFloat64) ->  Float64) func

// CFConvertFloatHostToSwapped(arg _, :  Float) ->  CFSwappedFloat32) func

// CFConvertFloatSwappedToHost(arg _, :  CFSwappedFloat32) ->  Float) func


// CFCopyDescription(cf _, :  CFTypeRef!) ->  CFString!) func

// CFCopyHomeDirectoryURL() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFDataAppendBytes(theData _, bytes :  CFMutableData!,  _, length :  UnsafePointer< UInt8>!,  _, :  CFIndex) func


// CFDataCreate(allocator _, bytes :  CFAllocator!,  _, length :  UnsafePointer< UInt8>!,  _, :  CFIndex) ->  CFData!) func

// CFDataCreateCopy(allocator _, theData :  CFAllocator!,  _, :  CFData!) ->  CFData!) func

// CFDataCreateMutable(allocator _, capacity :  CFAllocator!,  _, :  CFIndex) ->  CFMutableData!) func


// CFDataCreateMutableCopy(allocator _, capacity :  CFAllocator!,  _, theData :  CFIndex,  _, :  CFData!) ->  CFMutableData!) func

// CFDataCreateWithBytesNoCopy(allocator _, bytes :  CFAllocator!,  _, length :  UnsafePointer< UInt8>!,  _, bytesDeallocator :  CFIndex,  _, :  CFAllocator!) ->  CFData!) func

// CFDataDeleteBytes(theData _, range :  CFMutableData!,  _, :  CFRange) func


// CFDataFind(theData _, dataToFind :  CFData!,  _, searchRange :  CFData!,  _, compareOptions :  CFRange,  _, :  CFDataSearchFlags) ->  CFRange) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFDataGetBytePtr(theData _, :  CFData!) ->  UnsafePointer< UInt8>!) func

// CFDataGetBytes(theData _, range :  CFData!,  _, buffer :  CFRange,  _, :  UnsafeMutablePointer< UInt8>!)) func


// CFDataGetLength(theData _, :  CFData!) ->  CFIndex) func

// CFDataGetMutableBytePtr(theData _, :  CFMutableData!) ->  UnsafeMutablePointer< UInt8>!) func

// CFDataGetTypeID() func


// CFDataIncreaseLength(theData _, extraLength :  CFMutableData!,  _, :  CFIndex) func

// CFDataReplaceBytes(theData _, range :  CFMutableData!,  _, newBytes :  CFRange,  _, newLength :  UnsafePointer< UInt8>!,  _, :  CFIndex) func

// CFDataSetLength(theData _, length :  CFMutableData!,  _, :  CFIndex) func


// CFDateCompare(theDate _, otherDate :  CFDate!,  _, context :  CFDate!,  _, :  UnsafeMutableRawPointer!) ->  CFComparisonResult) func

// CFDateCreate(allocator _, at :  CFAllocator!,  _, :  CFAbsoluteTime) ->  CFDate!) func

// CFDateFormatterCopyProperty(formatter _, key :  CFDateFormatter!,  _, :  CFDateFormatterKey!) ->  CFTypeRef!) func


// CFDateFormatterCreate(allocator _, locale :  CFAllocator!,  _, dateStyle :  CFLocale!,  _, timeStyle :  CFDateFormatterStyle,  _, :  CFDateFormatterStyle) ->  CFDateFormatter!) func

// CFDateFormatterCreateDateFormatFromTemplate(allocator _, tmplate :  CFAllocator!,  _, options :  CFString!,  _, locale :  CFOptionFlags,  _, :  CFLocale!) ->  CFString!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFDateFormatterCreateDateFromString(allocator _, formatter :  CFAllocator!,  _, string :  CFDateFormatter!,  _, rangep :  CFString!,  _, :  UnsafeMutablePointer< CFRange>!) ->  CFDate!) func


// CFDateFormatterCreateISO8601Formatter(allocator _, formatOptions :  CFAllocator!,  _, :  CFISO8601DateFormatOptions) ->  CFDateFormatter!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 10.0+
//   - visionOS 1.0+
//   - watchOS 3.0+

// CFDateFormatterCreateStringWithAbsoluteTime(allocator _, formatter :  CFAllocator!,  _, at :  CFDateFormatter!,  _, :  CFAbsoluteTime) ->  CFString!) func

// CFDateFormatterCreateStringWithDate(allocator _, formatter :  CFAllocator!,  _, date :  CFDateFormatter!,  _, :  CFDate!) ->  CFString!) func


// CFDateFormatterGetAbsoluteTimeFromString(formatter _, string :  CFDateFormatter!,  _, rangep :  CFString!,  _, atp :  UnsafeMutablePointer< CFRange>!,  _, :  UnsafeMutablePointer< CFAbsoluteTime>!) ->  Bool) func

// CFDateFormatterGetDateStyle(formatter _, :  CFDateFormatter!) ->  CFDateFormatterStyle) func

// CFDateFormatterGetFormat(formatter _, :  CFDateFormatter!) ->  CFString!) func


// CFDateFormatterGetLocale(formatter _, :  CFDateFormatter!) ->  CFLocale!) func

// CFDateFormatterGetTimeStyle(formatter _, :  CFDateFormatter!) ->  CFDateFormatterStyle) func

// CFDateFormatterGetTypeID() func


// CFDateFormatterSetFormat(formatter _, formatString :  CFDateFormatter!,  _, :  CFString!)) func

// CFDateFormatterSetProperty(formatter _, key :  CFDateFormatter!,  _, value :  CFString!,  _, :  CFTypeRef!)) func

// CFDateGetAbsoluteTime(theDate _, :  CFDate!) ->  CFAbsoluteTime) func


// CFDateGetTimeIntervalSinceDate(theDate _, otherDate :  CFDate!,  _, :  CFDate!) ->  CFTimeInterval) func

// CFDateGetTypeID() func

// CFDictionaryAddValue(theDict _, key :  CFMutableDictionary!,  _, value :  UnsafeRawPointer!,  _, :  UnsafeRawPointer!)) func


// CFDictionaryApplyFunction(theDict _, applier :  CFDictionary!,  _, context : (( UnsafeRawPointer?,  UnsafeRawPointer?,  UnsafeMutableRawPointer?) ->  Void)!,  _, :  UnsafeMutableRawPointer!)) func

// CFDictionaryContainsKey(theDict _, key :  CFDictionary!,  _, :  UnsafeRawPointer!) ->  Bool) func

// CFDictionaryContainsValue(theDict _, value :  CFDictionary!,  _, :  UnsafeRawPointer!) ->  Bool) func


// CFDictionaryCreate(allocator _, keys :  CFAllocator!,  _, values :  UnsafeMutablePointer< UnsafeRawPointer?>!,  _, numValues :  UnsafeMutablePointer< UnsafeRawPointer?>!,  _, keyCallBacks :  CFIndex,  _, valueCallBacks :  UnsafePointer< CFDictionaryKeyCallBacks>!,  _, :  UnsafePointer< CFDictionaryValueCallBacks>!) ->  CFDictionary!) func

// CFDictionaryCreateCopy(allocator _, theDict :  CFAllocator!,  _, :  CFDictionary!) ->  CFDictionary!) func

// CFDictionaryCreateMutable(allocator _, capacity :  CFAllocator!,  _, keyCallBacks :  CFIndex,  _, valueCallBacks :  UnsafePointer< CFDictionaryKeyCallBacks>!,  _, :  UnsafePointer< CFDictionaryValueCallBacks>!) ->  CFMutableDictionary!) func


// CFDictionaryCreateMutableCopy(allocator _, capacity :  CFAllocator!,  _, theDict :  CFIndex,  _, :  CFDictionary!) ->  CFMutableDictionary!) func

// CFDictionaryGetCount(theDict _, :  CFDictionary!) ->  CFIndex) func

// CFDictionaryGetCountOfKey(theDict _, key :  CFDictionary!,  _, :  UnsafeRawPointer!) ->  CFIndex) func


// CFDictionaryGetCountOfValue(theDict _, value :  CFDictionary!,  _, :  UnsafeRawPointer!) ->  CFIndex) func

// CFDictionaryGetKeysAndValues(theDict _, keys :  CFDictionary!,  _, values :  UnsafeMutablePointer< UnsafeRawPointer?>!,  _, :  UnsafeMutablePointer< UnsafeRawPointer?>!)) func

// CFDictionaryGetTypeID() func


// CFDictionaryGetValue(theDict _, key :  CFDictionary!,  _, :  UnsafeRawPointer!) ->  UnsafeRawPointer!) func

// CFDictionaryGetValueIfPresent(theDict _, key :  CFDictionary!,  _, value :  UnsafeRawPointer!,  _, :  UnsafeMutablePointer< UnsafeRawPointer?>!) ->  Bool) func

// CFDictionaryRemoveAllValues(theDict _, :  CFMutableDictionary!)) func


// CFDictionaryRemoveValue(theDict _, key :  CFMutableDictionary!,  _, :  UnsafeRawPointer!)) func

// CFDictionaryReplaceValue(theDict _, key :  CFMutableDictionary!,  _, value :  UnsafeRawPointer!,  _, :  UnsafeRawPointer!)) func

// CFDictionarySetValue(theDict _, key :  CFMutableDictionary!,  _, value :  UnsafeRawPointer!,  _, :  UnsafeRawPointer!)) func


// CFEqual(cf1 _, cf2 :  CFTypeRef!,  _, :  CFTypeRef!) ->  Bool) func

// CFErrorCopyDescription(err _, :  CFError!) ->  CFString!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFErrorCopyFailureReason(err _, :  CFError!) ->  CFString!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFErrorCopyRecoverySuggestion(err _, :  CFError!) ->  CFString!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFErrorCopyUserInfo(err _, :  CFError!) ->  CFDictionary!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFErrorCreate(allocator _, domain :  CFAllocator!,  _, code :  CFErrorDomain!,  _, userInfo :  CFIndex,  _, :  CFDictionary!) ->  CFError!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFErrorCreateWithUserInfoKeysAndValues(allocator _, domain :  CFAllocator!,  _, code :  CFErrorDomain!,  _, userInfoKeys :  CFIndex,  _, userInfoValues :  UnsafePointer< UnsafeRawPointer?>!,  _, numUserInfoValues :  UnsafePointer< UnsafeRawPointer?>!,  _, :  CFIndex) ->  CFError!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFErrorGetCode(err _, :  CFError!) ->  CFIndex) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFErrorGetDomain(err _, :  CFError!) ->  CFErrorDomain!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFErrorGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorCreate(allocator _, fd :  CFAllocator!,  _, closeOnInvalidate :  CFFileDescriptorNativeDescriptor,  _, callout :  Bool,  _, context :  CFFileDescriptorCallBack!,  _, :  UnsafePointer< CFFileDescriptorContext>!) ->  CFFileDescriptor!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorCreateRunLoopSource(allocator _, f :  CFAllocator!,  _, order :  CFFileDescriptor!,  _, :  CFIndex) ->  CFRunLoopSource!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileDescriptorDisableCallBacks(f _, callBackTypes :  CFFileDescriptor!,  _, :  CFOptionFlags) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorEnableCallBacks(f _, callBackTypes :  CFFileDescriptor!,  _, :  CFOptionFlags) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorGetContext(f _, context :  CFFileDescriptor!,  _, :  UnsafeMutablePointer< CFFileDescriptorContext>!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileDescriptorGetNativeDescriptor(f _, :  CFFileDescriptor!) ->  CFFileDescriptorNativeDescriptor) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileDescriptorInvalidate(f _, :  CFFileDescriptor!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileDescriptorIsValid(f _, :  CFFileDescriptor!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityClearProperties(fileSec _, clearPropertyMask :  CFFileSecurity!,  _, :  CFFileSecurityClearOptions) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityCopyAccessControlList(fileSec _, accessControlList :  CFFileSecurity!,  _, :  UnsafeMutablePointer< acl_t?>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileSecurityCopyGroupUUID(fileSec _, groupUUID :  CFFileSecurity!,  _, :  UnsafeMutablePointer< Unmanaged< CFUUID>?>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityCopyOwnerUUID(fileSec _, ownerUUID :  CFFileSecurity!,  _, :  UnsafeMutablePointer< Unmanaged< CFUUID>?>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityCreate(allocator _, :  CFAllocator!) ->  CFFileSecurity!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileSecurityCreateCopy(allocator _, fileSec :  CFAllocator!,  _, :  CFFileSecurity!) ->  CFFileSecurity!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityGetGroup(fileSec _, group :  CFFileSecurity!,  _, :  UnsafeMutablePointer< gid_t>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityGetMode(fileSec _, mode :  CFFileSecurity!,  _, :  UnsafeMutablePointer< mode_t>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileSecurityGetOwner(fileSec _, owner :  CFFileSecurity!,  _, :  UnsafeMutablePointer< uid_t>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecurityGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecuritySetAccessControlList(fileSec _, accessControlList :  CFFileSecurity!,  _, :  acl_t!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileSecuritySetGroup(fileSec _, group :  CFFileSecurity!,  _, :  gid_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecuritySetGroupUUID(fileSec _, groupUUID :  CFFileSecurity!,  _, :  CFUUID!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecuritySetMode(fileSec _, mode :  CFFileSecurity!,  _, :  mode_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFFileSecuritySetOwner(fileSec _, owner :  CFFileSecurity!,  _, :  uid_t) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFFileSecuritySetOwnerUUID(fileSec _, ownerUUID :  CFFileSecurity!,  _, :  CFUUID!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFGetTypeID(cf _, :  CFTypeRef!) ->  CFTypeID) func


// CFGregorianDateGetAbsoluteTime(gdate _, tz :  CFGregorianDate,  _, :  CFTimeZone!) ->  CFAbsoluteTime) func
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

// CFGregorianDateIsValid(gdate _, unitFlags :  CFGregorianDate,  _, :  CFOptionFlags) ->  Bool) func
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

// CFHash(cf _, :  CFTypeRef!) ->  CFHashCode) func


// CFLocaleCopyAvailableLocaleIdentifiers() func

// CFLocaleCopyCommonISOCurrencyCodes() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFLocaleCopyCurrent() func


// CFLocaleCopyDisplayNameForPropertyValue(displayLocale _, key :  CFLocale!,  _, value :  CFLocaleKey!,  _, :  CFString!) ->  CFString!) func

// CFLocaleCopyISOCountryCodes() func

// CFLocaleCopyISOCurrencyCodes() func


// CFLocaleCopyISOLanguageCodes() func

// CFLocaleCopyPreferredLanguages() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFLocaleCreate(allocator _, localeIdentifier :  CFAllocator!,  _, :  CFLocaleIdentifier!) ->  CFLocale!) func


// CFLocaleCreateCanonicalLanguageIdentifierFromString(allocator _, localeIdentifier :  CFAllocator!,  _, :  CFString!) ->  CFLocaleIdentifier!) func

// CFLocaleCreateCanonicalLocaleIdentifierFromScriptManagerCodes(allocator _, lcode :  CFAllocator!,  _, rcode :  LangCode,  _, :  RegionCode) ->  CFLocaleIdentifier!) func

// CFLocaleCreateCanonicalLocaleIdentifierFromString(allocator _, localeIdentifier :  CFAllocator!,  _, :  CFString!) ->  CFLocaleIdentifier!) func


// CFLocaleCreateComponentsFromLocaleIdentifier(allocator _, localeID :  CFAllocator!,  _, :  CFLocaleIdentifier!) ->  CFDictionary!) func

// CFLocaleCreateCopy(allocator _, locale :  CFAllocator!,  _, :  CFLocale!) ->  CFLocale!) func

// CFLocaleCreateLocaleIdentifierFromComponents(allocator _, dictionary :  CFAllocator!,  _, :  CFDictionary!) ->  CFLocaleIdentifier!) func


// CFLocaleCreateLocaleIdentifierFromWindowsLocaleCode(allocator _, lcid :  CFAllocator!,  _, :  UInt32) ->  CFLocaleIdentifier!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFLocaleGetIdentifier(locale _, :  CFLocale!) ->  CFLocaleIdentifier!) func

// CFLocaleGetLanguageCharacterDirection(isoLangCode _, :  CFString!) ->  CFLocaleLanguageDirection) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFLocaleGetLanguageLineDirection(isoLangCode _, :  CFString!) ->  CFLocaleLanguageDirection) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFLocaleGetSystem() func

// CFLocaleGetTypeID() func


// CFLocaleGetValue(locale _, key :  CFLocale!,  _, :  CFLocaleKey!) ->  CFTypeRef!) func

// CFLocaleGetWindowsLocaleCodeFromLocaleIdentifier(localeIdentifier _, :  CFLocaleIdentifier!) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFMachPortCreate(allocator _, callout :  CFAllocator!,  _, context :  CFMachPortCallBack!,  _, shouldFreeInfo :  UnsafeMutablePointer< CFMachPortContext>!,  _, :  UnsafeMutablePointer< DarwinBoolean>!) ->  CFMachPort!) func


// CFMachPortCreateRunLoopSource(allocator _, port :  CFAllocator!,  _, order :  CFMachPort!,  _, :  CFIndex) ->  CFRunLoopSource!) func

// CFMachPortCreateWithPort(allocator _, portNum :  CFAllocator!,  _, callout :  mach_port_t,  _, context :  CFMachPortCallBack!,  _, shouldFreeInfo :  UnsafeMutablePointer< CFMachPortContext>!,  _, :  UnsafeMutablePointer< DarwinBoolean>!) ->  CFMachPort!) func

// CFMachPortGetContext(port _, context :  CFMachPort!,  _, :  UnsafeMutablePointer< CFMachPortContext>!)) func


// CFMachPortGetInvalidationCallBack(port _, :  CFMachPort!) ->  CFMachPortInvalidationCallBack!) func

// CFMachPortGetPort(port _, :  CFMachPort!) ->  mach_port_t) func

// CFMachPortGetTypeID() func


// CFMachPortInvalidate(port _, :  CFMachPort!)) func

// CFMachPortIsValid(port _, :  CFMachPort!) ->  Bool) func

// CFMachPortSetInvalidationCallBack(port _, callout :  CFMachPort!,  _, :  CFMachPortInvalidationCallBack!)) func


// CFMessagePortCreateLocal(allocator _, name :  CFAllocator!,  _, callout :  CFString!,  _, context :  CFMessagePortCallBack!,  _, shouldFreeInfo :  UnsafeMutablePointer< CFMessagePortContext>!,  _, :  UnsafeMutablePointer< DarwinBoolean>!) ->  CFMessagePort!) func

// CFMessagePortCreateRemote(allocator _, name :  CFAllocator!,  _, :  CFString!) ->  CFMessagePort!) func

// CFMessagePortCreateRunLoopSource(allocator _, local :  CFAllocator!,  _, order :  CFMessagePort!,  _, :  CFIndex) ->  CFRunLoopSource!) func


// CFMessagePortGetContext(ms _, context :  CFMessagePort!,  _, :  UnsafeMutablePointer< CFMessagePortContext>!)) func

// CFMessagePortGetInvalidationCallBack(ms _, :  CFMessagePort!) ->  CFMessagePortInvalidationCallBack!) func

// CFMessagePortGetName(ms _, :  CFMessagePort!) ->  CFString!) func


// CFMessagePortGetTypeID() func

// CFMessagePortInvalidate(ms _, :  CFMessagePort!)) func

// CFMessagePortIsRemote(ms _, :  CFMessagePort!) ->  Bool) func


// CFMessagePortIsValid(ms _, :  CFMessagePort!) ->  Bool) func

// CFMessagePortSendRequest(remote _, msgid :  CFMessagePort!,  _, data :  Int32,  _, sendTimeout :  CFData!,  _, rcvTimeout :  CFTimeInterval,  _, replyMode :  CFTimeInterval,  _, returnData :  CFString!,  _, :  UnsafeMutablePointer< Unmanaged< CFData>?>!) ->  Int32) func

// CFMessagePortSetDispatchQueue(ms _, queue :  CFMessagePort!,  _, :  dispatch_queue_t!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFMessagePortSetInvalidationCallBack(ms _, callout :  CFMessagePort!,  _, :  CFMessagePortInvalidationCallBack!)) func

// CFMessagePortSetName(ms _, newName :  CFMessagePort!,  _, :  CFString!) ->  Bool) func

// CFNotificationCenterAddObserver(center _, observer :  CFNotificationCenter!,  _, callBack :  UnsafeRawPointer!,  _, name :  CFNotificationCallback!,  _, object :  CFString!,  _, suspensionBehavior :  UnsafeRawPointer!,  _, :  CFNotificationSuspensionBehavior) func


// CFNotificationCenterGetDarwinNotifyCenter() func

// CFNotificationCenterGetDistributedCenter() func

// CFNotificationCenterGetLocalCenter() func


// CFNotificationCenterGetTypeID() func

// CFNotificationCenterPostNotification(center _, name :  CFNotificationCenter!,  _, object :  CFNotificationName!,  _, userInfo :  UnsafeRawPointer!,  _, deliverImmediately :  CFDictionary!,  _, :  Bool) func

// CFNotificationCenterPostNotificationWithOptions(center _, name :  CFNotificationCenter!,  _, object :  CFNotificationName!,  _, userInfo :  UnsafeRawPointer!,  _, options :  CFDictionary!,  _, :  CFOptionFlags) func


// CFNotificationCenterRemoveEveryObserver(center _, observer :  CFNotificationCenter!,  _, :  UnsafeRawPointer!)) func

// CFNotificationCenterRemoveObserver(center _, observer :  CFNotificationCenter!,  _, name :  UnsafeRawPointer!,  _, object :  CFNotificationName!,  _, :  UnsafeRawPointer!)) func

// CFNullGetTypeID() func


// CFNumberCompare(number _, otherNumber :  CFNumber!,  _, context :  CFNumber!,  _, :  UnsafeMutableRawPointer!) ->  CFComparisonResult) func

// CFNumberCreate(allocator _, theType :  CFAllocator!,  _, valuePtr :  CFNumberType,  _, :  UnsafeRawPointer!) ->  CFNumber!) func

// CFNumberFormatterCopyProperty(formatter _, key :  CFNumberFormatter!,  _, :  CFNumberFormatterKey!) ->  CFTypeRef!) func


// CFNumberFormatterCreate(allocator _, locale :  CFAllocator!,  _, style :  CFLocale!,  _, :  CFNumberFormatterStyle) ->  CFNumberFormatter!) func

// CFNumberFormatterCreateNumberFromString(allocator _, formatter :  CFAllocator!,  _, string :  CFNumberFormatter!,  _, rangep :  CFString!,  _, options :  UnsafeMutablePointer< CFRange>!,  _, :  CFOptionFlags) ->  CFNumber!) func

// CFNumberFormatterCreateStringWithNumber(allocator _, formatter :  CFAllocator!,  _, number :  CFNumberFormatter!,  _, :  CFNumber!) ->  CFString!) func


// CFNumberFormatterCreateStringWithValue(allocator _, formatter :  CFAllocator!,  _, numberType :  CFNumberFormatter!,  _, valuePtr :  CFNumberType,  _, :  UnsafeRawPointer!) ->  CFString!) func

// CFNumberFormatterGetDecimalInfoForCurrencyCode(currencyCode _, defaultFractionDigits :  CFString!,  _, roundingIncrement :  UnsafeMutablePointer< Int32>!,  _, :  UnsafeMutablePointer< Double>!) ->  Bool) func

// CFNumberFormatterGetFormat(formatter _, :  CFNumberFormatter!) ->  CFString!) func


// CFNumberFormatterGetLocale(formatter _, :  CFNumberFormatter!) ->  CFLocale!) func

// CFNumberFormatterGetStyle(formatter _, :  CFNumberFormatter!) ->  CFNumberFormatterStyle) func

// CFNumberFormatterGetTypeID() func


// CFNumberFormatterGetValueFromString(formatter _, string :  CFNumberFormatter!,  _, rangep :  CFString!,  _, numberType :  UnsafeMutablePointer< CFRange>!,  _, valuePtr :  CFNumberType,  _, :  UnsafeMutableRawPointer!) ->  Bool) func

// CFNumberFormatterSetFormat(formatter _, formatString :  CFNumberFormatter!,  _, :  CFString!)) func

// CFNumberFormatterSetProperty(formatter _, key :  CFNumberFormatter!,  _, value :  CFNumberFormatterKey!,  _, :  CFTypeRef!)) func


// CFNumberGetByteSize(number _, :  CFNumber!) ->  CFIndex) func

// CFNumberGetType(number _, :  CFNumber!) ->  CFNumberType) func

// CFNumberGetTypeID() func


// CFNumberGetValue(number _, theType :  CFNumber!,  _, valuePtr :  CFNumberType,  _, :  UnsafeMutableRawPointer!) ->  Bool) func

// CFNumberIsFloatType(number _, :  CFNumber!) ->  Bool) func

// CFPlugInAddInstanceForFactory(factoryID _, :  CFUUID!)) func


// CFPlugInCreate(allocator _, plugInURL :  CFAllocator!,  _, :  CFURL!) ->  CFPlugIn!) func

// CFPlugInFindFactoriesForPlugInType(typeUUID _, :  CFUUID!) ->  CFArray!) func

// CFPlugInFindFactoriesForPlugInTypeInPlugIn(typeUUID _, plugIn :  CFUUID!,  _, :  CFPlugIn!) ->  CFArray!) func


// CFPlugInGetBundle(plugIn _, :  CFPlugIn!) ->  CFBundle!) func

// CFPlugInGetTypeID() func

// CFPlugInInstanceCreate(allocator _, factoryUUID :  CFAllocator!,  _, typeUUID :  CFUUID!,  _, :  CFUUID!) ->  UnsafeMutableRawPointer!) func


// CFPlugInInstanceCreateWithInstanceDataSize(allocator _, instanceDataSize :  CFAllocator!,  _, deallocateInstanceFunction :  CFIndex,  _, factoryName :  CFPlugInInstanceDeallocateInstanceDataFunction!,  _, getInterfaceFunction :  CFString!,  _, :  CFPlugInInstanceGetInterfaceFunction!) ->  CFPlugInInstance!) func

// CFPlugInInstanceGetFactoryName(instance _, :  CFPlugInInstance!) ->  CFString!) func

// CFPlugInInstanceGetInstanceData(instance _, :  CFPlugInInstance!) ->  UnsafeMutableRawPointer!) func


// CFPlugInInstanceGetInterfaceFunctionTable(instance _, interfaceName :  CFPlugInInstance!,  _, ftbl :  CFString!,  _, :  UnsafeMutablePointer< UnsafeMutableRawPointer?>!) ->  Bool) func

// CFPlugInInstanceGetTypeID() func

// CFPlugInIsLoadOnDemand(plugIn _, :  CFPlugIn!) ->  Bool) func


// CFPlugInRegisterFactoryFunction(factoryUUID _, func :  CFUUID!,  _, :  CFPlugInFactoryFunction!) ->  Bool) func

// CFPlugInRegisterFactoryFunctionByName(factoryUUID _, plugIn :  CFUUID!,  _, functionName :  CFPlugIn!,  _, :  CFString!) ->  Bool) func

// CFPlugInRegisterPlugInType(factoryUUID _, typeUUID :  CFUUID!,  _, :  CFUUID!) ->  Bool) func


// CFPlugInRemoveInstanceForFactory(factoryID _, :  CFUUID!)) func

// CFPlugInSetLoadOnDemand(plugIn _, flag :  CFPlugIn!,  _, :  Bool) func

// CFPlugInUnregisterFactory(factoryUUID _, :  CFUUID!) ->  Bool) func


// CFPlugInUnregisterPlugInType(factoryUUID _, typeUUID :  CFUUID!,  _, :  CFUUID!) ->  Bool) func

// CFPreferencesAddSuitePreferencesToApp(applicationID _, suiteID :  CFString,  _, :  CFString) func

// CFPreferencesAppSynchronize(applicationID _, :  CFString) ->  Bool) func


// CFPreferencesAppValueIsForced(key _, applicationID :  CFString,  _, :  CFString) ->  Bool) func

// CFPreferencesCopyAppValue(key _, applicationID :  CFString,  _, :  CFString) ->  CFPropertyList?) func

// CFPreferencesCopyApplicationList(userName _, hostName :  CFString,  _, :  CFString) ->  CFArray?) func
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.


// CFPreferencesCopyKeyList(applicationID _, userName :  CFString,  _, hostName :  CFString,  _, :  CFString) ->  CFArray?) func

// CFPreferencesCopyMultiple(keysToFetch _, applicationID :  CFArray?,  _, userName :  CFString,  _, hostName :  CFString,  _, :  CFString) ->  CFDictionary) func

// CFPreferencesCopyValue(key _, applicationID :  CFString,  _, userName :  CFString,  _, hostName :  CFString,  _, :  CFString) ->  CFPropertyList?) func


// CFPreferencesGetAppBooleanValue(key _, applicationID :  CFString,  _, keyExistsAndHasValidFormat :  CFString,  _, :  UnsafeMutablePointer< DarwinBoolean>?) ->  Bool) func

// CFPreferencesGetAppIntegerValue(key _, applicationID :  CFString,  _, keyExistsAndHasValidFormat :  CFString,  _, :  UnsafeMutablePointer< DarwinBoolean>?) ->  CFIndex) func

// CFPreferencesRemoveSuitePreferencesFromApp(applicationID _, suiteID :  CFString,  _, :  CFString) func


// CFPreferencesSetAppValue(key _, value :  CFString,  _, applicationID :  CFPropertyList?,  _, :  CFString) func

// CFPreferencesSetMultiple(keysToSet _, keysToRemove :  CFDictionary?,  _, applicationID :  CFArray?,  _, userName :  CFString,  _, hostName :  CFString,  _, :  CFString) func

// CFPreferencesSetValue(key _, value :  CFString,  _, applicationID :  CFPropertyList?,  _, userName :  CFString,  _, hostName :  CFString,  _, :  CFString) func


// CFPreferencesSynchronize(applicationID _, userName :  CFString,  _, hostName :  CFString,  _, :  CFString) ->  Bool) func

// CFPropertyListCreateData(allocator _, propertyList :  CFAllocator!,  _, format :  CFPropertyList!,  _, options :  CFPropertyListFormat,  _, error :  CFOptionFlags,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Unmanaged< CFData>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFPropertyListCreateDeepCopy(allocator _, propertyList :  CFAllocator!,  _, mutabilityOption :  CFPropertyList!,  _, :  CFOptionFlags) ->  CFPropertyList!) func


// CFPropertyListCreateFromStream(allocator _, stream :  CFAllocator!,  _, streamLength :  CFReadStream!,  _, mutabilityOption :  CFIndex,  _, format :  CFOptionFlags,  _, errorString :  UnsafeMutablePointer< CFPropertyListFormat>!,  _, :  UnsafeMutablePointer< Unmanaged< CFString>?>!) ->  Unmanaged< CFPropertyList>!) func
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

// CFPropertyListCreateFromXMLData(allocator _, xmlData :  CFAllocator!,  _, mutabilityOption :  CFData!,  _, errorString :  CFOptionFlags,  _, :  UnsafeMutablePointer< Unmanaged< CFString>?>!) ->  Unmanaged< CFPropertyList>!) func
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

// CFPropertyListCreateWithData(allocator _, data :  CFAllocator!,  _, options :  CFData!,  _, format :  CFOptionFlags,  _, error :  UnsafeMutablePointer< CFPropertyListFormat>!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Unmanaged< CFPropertyList>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFPropertyListCreateWithStream(allocator _, stream :  CFAllocator!,  _, streamLength :  CFReadStream!,  _, options :  CFIndex,  _, format :  CFOptionFlags,  _, error :  UnsafeMutablePointer< CFPropertyListFormat>!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Unmanaged< CFPropertyList>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFPropertyListCreateXMLData(allocator _, propertyList :  CFAllocator!,  _, :  CFPropertyList!) ->  Unmanaged< CFData>!) func
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

// CFPropertyListIsValid(plist _, format :  CFPropertyList!,  _, :  CFPropertyListFormat) ->  Bool) func


// CFPropertyListWrite(propertyList _, stream :  CFPropertyList!,  _, format :  CFWriteStream!,  _, options :  CFPropertyListFormat,  _, error :  CFOptionFlags,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  CFIndex) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFPropertyListWriteToStream(propertyList _, stream :  CFPropertyList!,  _, format :  CFWriteStream!,  _, errorString :  CFPropertyListFormat,  _, :  UnsafeMutablePointer< Unmanaged< CFString>?>!) ->  CFIndex) func
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

// CFRangeMake(loc _, len :  CFIndex,  _, :  CFIndex) ->  CFRange) func


// CFReadStreamClose(stream _, :  CFReadStream!)) func

// CFReadStreamCopyDispatchQueue(stream _, :  CFReadStream!) ->  dispatch_queue_t!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFReadStreamCopyError(stream _, :  CFReadStream!) ->  CFError!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFReadStreamCopyProperty(stream _, propertyName :  CFReadStream!,  _, :  CFStreamPropertyKey!) ->  CFTypeRef!) func

// CFReadStreamCreateWithBytesNoCopy(alloc _, bytes :  CFAllocator!,  _, length :  UnsafePointer< UInt8>!,  _, bytesDeallocator :  CFIndex,  _, :  CFAllocator!) ->  CFReadStream!) func

// CFReadStreamCreateWithFile(alloc _, fileURL :  CFAllocator!,  _, :  CFURL!) ->  CFReadStream!) func


// CFReadStreamGetBuffer(stream _, maxBytesToRead :  CFReadStream!,  _, numBytesRead :  CFIndex,  _, :  UnsafeMutablePointer< CFIndex>!) ->  UnsafePointer< UInt8>!) func

// CFReadStreamGetError(stream _, :  CFReadStream!) ->  CFStreamError) func

// CFReadStreamGetStatus(stream _, :  CFReadStream!) ->  CFStreamStatus) func


// CFReadStreamGetTypeID() func

// CFReadStreamHasBytesAvailable(stream _, :  CFReadStream!) ->  Bool) func

// CFReadStreamOpen(stream _, :  CFReadStream!) ->  Bool) func


// CFReadStreamRead(stream _, buffer :  CFReadStream!,  _, bufferLength :  UnsafeMutablePointer< UInt8>!,  _, :  CFIndex) ->  CFIndex) func

// CFReadStreamScheduleWithRunLoop(stream _, runLoop :  CFReadStream!,  _, runLoopMode :  CFRunLoop!,  _, :  CFRunLoopMode!)) func

// CFReadStreamSetClient(stream _, streamEvents :  CFReadStream!,  _, clientCB :  CFOptionFlags,  _, clientContext :  CFReadStreamClientCallBack!,  _, :  UnsafeMutablePointer< CFStreamClientContext>!) ->  Bool) func


// CFReadStreamSetDispatchQueue(stream _, q :  CFReadStream!,  _, :  dispatch_queue_t!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFReadStreamSetProperty(stream _, propertyName :  CFReadStream!,  _, propertyValue :  CFStreamPropertyKey!,  _, :  CFTypeRef!) ->  Bool) func

// CFReadStreamUnscheduleFromRunLoop(stream _, runLoop :  CFReadStream!,  _, runLoopMode :  CFRunLoop!,  _, :  CFRunLoopMode!)) func


// CFRelease(cf CFTypeRef, );) extern   void

// CFRunLoopAddCommonMode(rl _, mode :  CFRunLoop!,  _, :  CFRunLoopMode!)) func

// CFRunLoopAddObserver(rl _, observer :  CFRunLoop!,  _, mode :  CFRunLoopObserver!,  _, :  CFRunLoopMode!)) func


// CFRunLoopAddSource(rl _, source :  CFRunLoop!,  _, mode :  CFRunLoopSource!,  _, :  CFRunLoopMode!)) func

// CFRunLoopAddTimer(rl _, timer :  CFRunLoop!,  _, mode :  CFRunLoopTimer!,  _, :  CFRunLoopMode!)) func

// CFRunLoopContainsObserver(rl _, observer :  CFRunLoop!,  _, mode :  CFRunLoopObserver!,  _, :  CFRunLoopMode!) ->  Bool) func


// CFRunLoopContainsSource(rl _, source :  CFRunLoop!,  _, mode :  CFRunLoopSource!,  _, :  CFRunLoopMode!) ->  Bool) func

// CFRunLoopContainsTimer(rl _, timer :  CFRunLoop!,  _, mode :  CFRunLoopTimer!,  _, :  CFRunLoopMode!) ->  Bool) func

// CFRunLoopCopyAllModes(rl _, :  CFRunLoop!) ->  CFArray!) func


// CFRunLoopCopyCurrentMode(rl _, :  CFRunLoop!) ->  CFRunLoopMode!) func

// CFRunLoopGetCurrent() func

// CFRunLoopGetMain() func


// CFRunLoopGetNextTimerFireDate(rl _, mode :  CFRunLoop!,  _, :  CFRunLoopMode!) ->  CFAbsoluteTime) func

// CFRunLoopGetTypeID() func

// CFRunLoopIsWaiting(rl _, :  CFRunLoop!) ->  Bool) func


// CFRunLoopObserverCreate(allocator _, activities :  CFAllocator!,  _, repeats :  CFOptionFlags,  _, order :  Bool,  _, callout :  CFIndex,  _, context :  CFRunLoopObserverCallBack!,  _, :  UnsafeMutablePointer< CFRunLoopObserverContext>!) ->  CFRunLoopObserver!) func

// CFRunLoopObserverCreateWithHandler(allocator _, activities :  CFAllocator!,  _, repeats :  CFOptionFlags,  _, order :  Bool,  _, block :  CFIndex,  _, : (( CFRunLoopObserver?,  CFRunLoopActivity) ->  Void)!) ->  CFRunLoopObserver!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFRunLoopObserverDoesRepeat(observer _, :  CFRunLoopObserver!) ->  Bool) func


// CFRunLoopObserverGetActivities(observer _, :  CFRunLoopObserver!) ->  CFOptionFlags) func

// CFRunLoopObserverGetContext(observer _, context :  CFRunLoopObserver!,  _, :  UnsafeMutablePointer< CFRunLoopObserverContext>!)) func

// CFRunLoopObserverGetOrder(observer _, :  CFRunLoopObserver!) ->  CFIndex) func


// CFRunLoopObserverGetTypeID() func

// CFRunLoopObserverInvalidate(observer _, :  CFRunLoopObserver!)) func

// CFRunLoopObserverIsValid(observer _, :  CFRunLoopObserver!) ->  Bool) func


// CFRunLoopPerformBlock(rl _, mode :  CFRunLoop!,  _, block :  CFTypeRef!,  _, : (() ->  Void)!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFRunLoopRemoveObserver(rl _, observer :  CFRunLoop!,  _, mode :  CFRunLoopObserver!,  _, :  CFRunLoopMode!)) func

// CFRunLoopRemoveSource(rl _, source :  CFRunLoop!,  _, mode :  CFRunLoopSource!,  _, :  CFRunLoopMode!)) func


// CFRunLoopRemoveTimer(rl _, timer :  CFRunLoop!,  _, mode :  CFRunLoopTimer!,  _, :  CFRunLoopMode!)) func

// CFRunLoopRun() func

// CFRunLoopRunInMode(mode _, seconds :  CFRunLoopMode!,  _, returnAfterSourceHandled :  CFTimeInterval,  _, :  Bool) ->  CFRunLoopRunResult) func


// CFRunLoopSourceCreate(allocator _, order :  CFAllocator!,  _, context :  CFIndex,  _, :  UnsafeMutablePointer< CFRunLoopSourceContext>!) ->  CFRunLoopSource!) func

// CFRunLoopSourceGetContext(source _, context :  CFRunLoopSource!,  _, :  UnsafeMutablePointer< CFRunLoopSourceContext>!)) func

// CFRunLoopSourceGetOrder(source _, :  CFRunLoopSource!) ->  CFIndex) func


// CFRunLoopSourceGetTypeID() func

// CFRunLoopSourceInvalidate(source _, :  CFRunLoopSource!)) func

// CFRunLoopSourceIsValid(source _, :  CFRunLoopSource!) ->  Bool) func


// CFRunLoopSourceSignal(source _, :  CFRunLoopSource!)) func

// CFRunLoopStop(rl _, :  CFRunLoop!)) func

// CFRunLoopTimerCreate(allocator _, fireDate :  CFAllocator!,  _, interval :  CFAbsoluteTime,  _, flags :  CFTimeInterval,  _, order :  CFOptionFlags,  _, callout :  CFIndex,  _, context :  CFRunLoopTimerCallBack!,  _, :  UnsafeMutablePointer< CFRunLoopTimerContext>!) ->  CFRunLoopTimer!) func


// CFRunLoopTimerCreateWithHandler(allocator _, fireDate :  CFAllocator!,  _, interval :  CFAbsoluteTime,  _, flags :  CFTimeInterval,  _, order :  CFOptionFlags,  _, block :  CFIndex,  _, : (( CFRunLoopTimer?) ->  Void)!) ->  CFRunLoopTimer!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFRunLoopTimerDoesRepeat(timer _, :  CFRunLoopTimer!) ->  Bool) func

// CFRunLoopTimerGetContext(timer _, context :  CFRunLoopTimer!,  _, :  UnsafeMutablePointer< CFRunLoopTimerContext>!)) func


// CFRunLoopTimerGetInterval(timer _, :  CFRunLoopTimer!) ->  CFTimeInterval) func

// CFRunLoopTimerGetNextFireDate(timer _, :  CFRunLoopTimer!) ->  CFAbsoluteTime) func

// CFRunLoopTimerGetOrder(timer _, :  CFRunLoopTimer!) ->  CFIndex) func


// CFRunLoopTimerGetTolerance(timer _, :  CFRunLoopTimer!) ->  CFTimeInterval) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFRunLoopTimerGetTypeID() func

// CFRunLoopTimerInvalidate(timer _, :  CFRunLoopTimer!)) func


// CFRunLoopTimerIsValid(timer _, :  CFRunLoopTimer!) ->  Bool) func

// CFRunLoopTimerSetNextFireDate(timer _, fireDate :  CFRunLoopTimer!,  _, :  CFAbsoluteTime) func

// CFRunLoopTimerSetTolerance(timer _, tolerance :  CFRunLoopTimer!,  _, :  CFTimeInterval) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFRunLoopWakeUp(rl _, :  CFRunLoop!)) func

// CFSetAddValue(theSet _, value :  CFMutableSet!,  _, :  UnsafeRawPointer!)) func

// CFSetApplyFunction(theSet _, applier :  CFSet!,  _, context : (( UnsafeRawPointer?,  UnsafeMutableRawPointer?) ->  Void)!,  _, :  UnsafeMutableRawPointer!)) func


// CFSetContainsValue(theSet _, value :  CFSet!,  _, :  UnsafeRawPointer!) ->  Bool) func

// CFSetCreate(allocator _, values :  CFAllocator!,  _, numValues :  UnsafeMutablePointer< UnsafeRawPointer?>!,  _, callBacks :  CFIndex,  _, :  UnsafePointer< CFSetCallBacks>!) ->  CFSet!) func

// CFSetCreateCopy(allocator _, theSet :  CFAllocator!,  _, :  CFSet!) ->  CFSet!) func


// CFSetCreateMutable(allocator _, capacity :  CFAllocator!,  _, callBacks :  CFIndex,  _, :  UnsafePointer< CFSetCallBacks>!) ->  CFMutableSet!) func

// CFSetCreateMutableCopy(allocator _, capacity :  CFAllocator!,  _, theSet :  CFIndex,  _, :  CFSet!) ->  CFMutableSet!) func

// CFSetGetCount(theSet _, :  CFSet!) ->  CFIndex) func


// CFSetGetCountOfValue(theSet _, value :  CFSet!,  _, :  UnsafeRawPointer!) ->  CFIndex) func

// CFSetGetTypeID() func

// CFSetGetValue(theSet _, value :  CFSet!,  _, :  UnsafeRawPointer!) ->  UnsafeRawPointer!) func


// CFSetGetValueIfPresent(theSet _, candidate :  CFSet!,  _, value :  UnsafeRawPointer!,  _, :  UnsafeMutablePointer< UnsafeRawPointer?>!) ->  Bool) func

// CFSetGetValues(theSet _, values :  CFSet!,  _, :  UnsafeMutablePointer< UnsafeRawPointer?>!)) func

// CFSetRemoveAllValues(theSet _, :  CFMutableSet!)) func


// CFSetRemoveValue(theSet _, value :  CFMutableSet!,  _, :  UnsafeRawPointer!)) func

// CFSetReplaceValue(theSet _, value :  CFMutableSet!,  _, :  UnsafeRawPointer!)) func

// CFSetSetValue(theSet _, value :  CFMutableSet!,  _, :  UnsafeRawPointer!)) func


// CFShowStr(str _, :  CFString!)) func

// CFSocketConnectToAddress(s _, address :  CFSocket!,  _, timeout :  CFData!,  _, :  CFTimeInterval) ->  CFSocketError) func

// CFSocketCopyAddress(s _, :  CFSocket!) ->  CFData!) func


// CFSocketCopyPeerAddress(s _, :  CFSocket!) ->  CFData!) func

// CFSocketCopyRegisteredSocketSignature(nameServerSignature _, timeout :  UnsafePointer< CFSocketSignature>!,  _, name :  CFTimeInterval,  _, signature :  CFString!,  _, nameServerAddress :  UnsafeMutablePointer< CFSocketSignature>!,  _, :  UnsafeMutablePointer< Unmanaged< CFData>?>!) ->  CFSocketError) func

// CFSocketCopyRegisteredValue(nameServerSignature _, timeout :  UnsafePointer< CFSocketSignature>!,  _, name :  CFTimeInterval,  _, value :  CFString!,  _, nameServerAddress :  UnsafeMutablePointer< Unmanaged< CFPropertyList>?>!,  _, :  UnsafeMutablePointer< Unmanaged< CFData>?>!) ->  CFSocketError) func


// CFSocketCreate(allocator _, protocolFamily :  CFAllocator!,  _, socketType :  Int32,  _, protocol :  Int32,  _, callBackTypes :  Int32,  _, callout :  CFOptionFlags,  _, context :  CFSocketCallBack!,  _, :  UnsafePointer< CFSocketContext>!) ->  CFSocket!) func

// CFSocketCreateConnectedToSocketSignature(allocator _, signature :  CFAllocator!,  _, callBackTypes :  UnsafePointer< CFSocketSignature>!,  _, callout :  CFOptionFlags,  _, context :  CFSocketCallBack!,  _, timeout :  UnsafePointer< CFSocketContext>!,  _, :  CFTimeInterval) ->  CFSocket!) func

// CFSocketCreateRunLoopSource(allocator _, s :  CFAllocator!,  _, order :  CFSocket!,  _, :  CFIndex) ->  CFRunLoopSource!) func


// CFSocketCreateWithNative(allocator _, sock :  CFAllocator!,  _, callBackTypes :  CFSocketNativeHandle,  _, callout :  CFOptionFlags,  _, context :  CFSocketCallBack!,  _, :  UnsafePointer< CFSocketContext>!) ->  CFSocket!) func

// CFSocketCreateWithSocketSignature(allocator _, signature :  CFAllocator!,  _, callBackTypes :  UnsafePointer< CFSocketSignature>!,  _, callout :  CFOptionFlags,  _, context :  CFSocketCallBack!,  _, :  UnsafePointer< CFSocketContext>!) ->  CFSocket!) func

// CFSocketDisableCallBacks(s _, callBackTypes :  CFSocket!,  _, :  CFOptionFlags) func


// CFSocketEnableCallBacks(s _, callBackTypes :  CFSocket!,  _, :  CFOptionFlags) func

// CFSocketGetContext(s _, context :  CFSocket!,  _, :  UnsafeMutablePointer< CFSocketContext>!)) func

// CFSocketGetDefaultNameRegistryPortNumber() func


// CFSocketGetNative(s _, :  CFSocket!) ->  CFSocketNativeHandle) func

// CFSocketGetSocketFlags(s _, :  CFSocket!) ->  CFOptionFlags) func

// CFSocketGetTypeID() func


// CFSocketInvalidate(s _, :  CFSocket!)) func

// CFSocketIsValid(s _, :  CFSocket!) ->  Bool) func

// CFSocketRegisterSocketSignature(nameServerSignature _, timeout :  UnsafePointer< CFSocketSignature>!,  _, name :  CFTimeInterval,  _, signature :  CFString!,  _, :  UnsafePointer< CFSocketSignature>!) ->  CFSocketError) func


// CFSocketRegisterValue(nameServerSignature _, timeout :  UnsafePointer< CFSocketSignature>!,  _, name :  CFTimeInterval,  _, value :  CFString!,  _, :  CFPropertyList!) ->  CFSocketError) func

// CFSocketSendData(s _, address :  CFSocket!,  _, data :  CFData!,  _, timeout :  CFData!,  _, :  CFTimeInterval) ->  CFSocketError) func

// CFSocketSetAddress(s _, address :  CFSocket!,  _, :  CFData!) ->  CFSocketError) func


// CFSocketSetDefaultNameRegistryPortNumber(port _, :  UInt16) func

// CFSocketSetSocketFlags(s _, flags :  CFSocket!,  _, :  CFOptionFlags) func

// CFSocketUnregister(nameServerSignature _, timeout :  UnsafePointer< CFSocketSignature>!,  _, name :  CFTimeInterval,  _, :  CFString!) ->  CFSocketError) func


// CFStreamCreateBoundPair(alloc _, readStream :  CFAllocator!,  _, writeStream :  UnsafeMutablePointer< Unmanaged< CFReadStream>?>!,  _, transferBufferSize :  UnsafeMutablePointer< Unmanaged< CFWriteStream>?>!,  _, :  CFIndex) func

// CFStreamCreatePairWithPeerSocketSignature(alloc _, signature :  CFAllocator!,  _, readStream :  UnsafePointer< CFSocketSignature>!,  _, writeStream :  UnsafeMutablePointer< Unmanaged< CFReadStream>?>!,  _, :  UnsafeMutablePointer< Unmanaged< CFWriteStream>?>!)) func
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

// CFStreamCreatePairWithSocket(alloc _, sock :  CFAllocator!,  _, readStream :  CFSocketNativeHandle,  _, writeStream :  UnsafeMutablePointer< Unmanaged< CFReadStream>?>!,  _, :  UnsafeMutablePointer< Unmanaged< CFWriteStream>?>!)) func
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


// CFStreamCreatePairWithSocketToHost(alloc _, host :  CFAllocator!,  _, port :  CFString!,  _, readStream :  UInt32,  _, writeStream :  UnsafeMutablePointer< Unmanaged< CFReadStream>?>!,  _, :  UnsafeMutablePointer< Unmanaged< CFWriteStream>?>!)) func
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

// CFStringAppend(theString _, appendedString :  CFMutableString!,  _, :  CFString!)) func

// CFStringAppendCString(theString _, cStr :  CFMutableString!,  _, encoding :  UnsafePointer< CChar>!,  _, :  CFStringEncoding) func


// CFStringAppendCharacters(theString _, chars :  CFMutableString!,  _, numChars :  UnsafePointer< UniChar>!,  _, :  CFIndex) func

// CFStringAppendFormat(theString CFMutableStringRef, formatOptions ,  CFDictionaryRef, format ,  CFStringRef, , ...);) extern   void

// CFStringAppendFormatAndArguments(theString _, formatOptions :  CFMutableString!,  _, format :  CFDictionary!,  _, arguments :  CFString!,  _, :  CVaListPointer) func


// CFStringAppendPascalString(theString _, pStr :  CFMutableString!,  _, encoding :  ConstStr255Param!,  _, :  CFStringEncoding) func

// CFStringCapitalize(theString _, locale :  CFMutableString!,  _, :  CFLocale!)) func

// CFStringCompare(theString1 _, theString2 :  CFString!,  _, compareOptions :  CFString!,  _, :  CFStringCompareFlags) ->  CFComparisonResult) func


// CFStringCompareWithOptions(theString1 _, theString2 :  CFString!,  _, rangeToCompare :  CFString!,  _, compareOptions :  CFRange,  _, :  CFStringCompareFlags) ->  CFComparisonResult) func

// CFStringCompareWithOptionsAndLocale(theString1 _, theString2 :  CFString!,  _, rangeToCompare :  CFString!,  _, compareOptions :  CFRange,  _, locale :  CFStringCompareFlags,  _, :  CFLocale!) ->  CFComparisonResult) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringConvertEncodingToIANACharSetName(encoding _, :  CFStringEncoding) ->  CFString!) func


// CFStringConvertEncodingToNSStringEncoding(encoding _, :  CFStringEncoding) ->  UInt) func

// CFStringConvertEncodingToWindowsCodepage(encoding _, :  CFStringEncoding) ->  UInt32) func

// CFStringConvertIANACharSetNameToEncoding(theString _, :  CFString!) ->  CFStringEncoding) func


// CFStringConvertNSStringEncodingToEncoding(encoding _, :  UInt) ->  CFStringEncoding) func

// CFStringConvertWindowsCodepageToEncoding(codepage _, :  UInt32) ->  CFStringEncoding) func

// CFStringCreateArrayBySeparatingStrings(alloc _, theString :  CFAllocator!,  _, separatorString :  CFString!,  _, :  CFString!) ->  CFArray!) func


// CFStringCreateArrayWithFindResults(alloc _, theString :  CFAllocator!,  _, stringToFind :  CFString!,  _, rangeToSearch :  CFString!,  _, compareOptions :  CFRange,  _, :  CFStringCompareFlags) ->  CFArray!) func

// CFStringCreateByCombiningStrings(alloc _, theArray :  CFAllocator!,  _, separatorString :  CFArray!,  _, :  CFString!) ->  CFString!) func

// CFStringCreateCopy(alloc _, theString :  CFAllocator!,  _, :  CFString!) ->  CFString!) func


// CFStringCreateExternalRepresentation(alloc _, theString :  CFAllocator!,  _, encoding :  CFString!,  _, lossByte :  CFStringEncoding,  _, :  UInt8) ->  CFData!) func

// CFStringCreateFromExternalRepresentation(alloc _, data :  CFAllocator!,  _, encoding :  CFData!,  _, :  CFStringEncoding) ->  CFString!) func

// CFStringCreateMutable(alloc _, maxLength :  CFAllocator!,  _, :  CFIndex) ->  CFMutableString!) func


// CFStringCreateMutableCopy(alloc _, maxLength :  CFAllocator!,  _, theString :  CFIndex,  _, :  CFString!) ->  CFMutableString!) func

// CFStringCreateMutableWithExternalCharactersNoCopy(alloc _, chars :  CFAllocator!,  _, numChars :  UnsafeMutablePointer< UniChar>!,  _, capacity :  CFIndex,  _, externalCharactersAllocator :  CFIndex,  _, :  CFAllocator!) ->  CFMutableString!) func

// CFStringCreateStringWithValidatedFormat(alloc CFAllocatorRef, formatOptions ,  CFDictionaryRef, validFormatSpecifiers ,  CFStringRef, format ,  CFStringRef, errorPtr ,  CFErrorRef *, , ...);) extern   CFStringRef
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 8.0+
//   - visionOS 1.0+
//   - watchOS 8.0+


// CFStringCreateStringWithValidatedFormatAndArguments(alloc CFAllocatorRef, formatOptions ,  CFDictionaryRef, validFormatSpecifiers ,  CFStringRef, format ,  CFStringRef, arguments ,  va_list, errorPtr ,  CFErrorRef *, );) extern   CFStringRef
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 8.0+
//   - visionOS 1.0+
//   - watchOS 8.0+

// CFStringCreateWithBytes(alloc _, bytes :  CFAllocator!,  _, numBytes :  UnsafePointer< UInt8>!,  _, encoding :  CFIndex,  _, isExternalRepresentation :  CFStringEncoding,  _, :  Bool) ->  CFString!) func

// CFStringCreateWithBytesNoCopy(alloc _, bytes :  CFAllocator!,  _, numBytes :  UnsafePointer< UInt8>!,  _, encoding :  CFIndex,  _, isExternalRepresentation :  CFStringEncoding,  _, contentsDeallocator :  Bool,  _, :  CFAllocator!) ->  CFString!) func


// CFStringCreateWithCString(alloc _, cStr :  CFAllocator!,  _, encoding :  UnsafePointer< CChar>!,  _, :  CFStringEncoding) ->  CFString!) func

// CFStringCreateWithCStringNoCopy(alloc _, cStr :  CFAllocator!,  _, encoding :  UnsafePointer< CChar>!,  _, contentsDeallocator :  CFStringEncoding,  _, :  CFAllocator!) ->  CFString!) func

// CFStringCreateWithCharacters(alloc _, chars :  CFAllocator!,  _, numChars :  UnsafePointer< UniChar>!,  _, :  CFIndex) ->  CFString!) func


// CFStringCreateWithCharactersNoCopy(alloc _, chars :  CFAllocator!,  _, numChars :  UnsafePointer< UniChar>!,  _, contentsDeallocator :  CFIndex,  _, :  CFAllocator!) ->  CFString!) func

// CFStringCreateWithFileSystemRepresentation(alloc _, buffer :  CFAllocator!,  _, :  UnsafePointer< CChar>!) ->  CFString!) func

// CFStringCreateWithFormat(alloc CFAllocatorRef, formatOptions ,  CFDictionaryRef, format ,  CFStringRef, , ...);) extern   CFStringRef


// CFStringCreateWithFormatAndArguments(alloc _, formatOptions :  CFAllocator!,  _, format :  CFDictionary!,  _, arguments :  CFString!,  _, :  CVaListPointer) ->  CFString!) func

// CFStringCreateWithPascalString(alloc _, pStr :  CFAllocator!,  _, encoding :  ConstStr255Param!,  _, :  CFStringEncoding) ->  CFString!) func

// CFStringCreateWithPascalStringNoCopy(alloc _, pStr :  CFAllocator!,  _, encoding :  ConstStr255Param!,  _, contentsDeallocator :  CFStringEncoding,  _, :  CFAllocator!) ->  CFString!) func


// CFStringCreateWithSubstring(alloc _, str :  CFAllocator!,  _, range :  CFString!,  _, :  CFRange) ->  CFString!) func

// CFStringDelete(theString _, range :  CFMutableString!,  _, :  CFRange) func

// CFStringFind(theString _, stringToFind :  CFString!,  _, compareOptions :  CFString!,  _, :  CFStringCompareFlags) ->  CFRange) func


// CFStringFindAndReplace(theString _, stringToFind :  CFMutableString!,  _, replacementString :  CFString!,  _, rangeToSearch :  CFString!,  _, compareOptions :  CFRange,  _, :  CFStringCompareFlags) ->  CFIndex) func

// CFStringFindCharacterFromSet(theString _, theSet :  CFString!,  _, rangeToSearch :  CFCharacterSet!,  _, searchOptions :  CFRange,  _, result :  CFStringCompareFlags,  _, :  UnsafeMutablePointer< CFRange>!) ->  Bool) func

// CFStringFindWithOptions(theString _, stringToFind :  CFString!,  _, rangeToSearch :  CFString!,  _, searchOptions :  CFRange,  _, result :  CFStringCompareFlags,  _, :  UnsafeMutablePointer< CFRange>!) ->  Bool) func


// CFStringFindWithOptionsAndLocale(theString _, stringToFind :  CFString!,  _, rangeToSearch :  CFString!,  _, searchOptions :  CFRange,  _, locale :  CFStringCompareFlags,  _, result :  CFLocale!,  _, :  UnsafeMutablePointer< CFRange>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringFold(theString _, theFlags :  CFMutableString!,  _, theLocale :  CFStringCompareFlags,  _, :  CFLocale!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringGetBytes(theString _, range :  CFString!,  _, encoding :  CFRange,  _, lossByte :  CFStringEncoding,  _, isExternalRepresentation :  UInt8,  _, buffer :  Bool,  _, maxBufLen :  UnsafeMutablePointer< UInt8>!,  _, usedBufLen :  CFIndex,  _, :  UnsafeMutablePointer< CFIndex>!) ->  CFIndex) func


// CFStringGetCString(theString _, buffer :  CFString!,  _, bufferSize :  UnsafeMutablePointer< CChar>!,  _, encoding :  CFIndex,  _, :  CFStringEncoding) ->  Bool) func

// CFStringGetCStringPtr(theString _, encoding :  CFString!,  _, :  CFStringEncoding) ->  UnsafePointer< CChar>!) func

// CFStringGetCharacterAtIndex(theString _, idx :  CFString!,  _, :  CFIndex) ->  UniChar) func


// CFStringGetCharacterFromInlineBuffer(buf _, idx :  UnsafeMutablePointer< CFStringInlineBuffer>!,  _, :  CFIndex) ->  UniChar) func

// CFStringGetCharacters(theString _, range :  CFString!,  _, buffer :  CFRange,  _, :  UnsafeMutablePointer< UniChar>!)) func

// CFStringGetCharactersPtr(theString _, :  CFString!) ->  UnsafePointer< UniChar>!) func


// CFStringGetDoubleValue(str _, :  CFString!) ->  Double) func

// CFStringGetFastestEncoding(theString _, :  CFString!) ->  CFStringEncoding) func

// CFStringGetFileSystemRepresentation(string _, buffer :  CFString!,  _, maxBufLen :  UnsafeMutablePointer< CChar>!,  _, :  CFIndex) ->  Bool) func


// CFStringGetHyphenationLocationBeforeIndex(string _, location :  CFString!,  _, limitRange :  CFIndex,  _, options :  CFRange,  _, locale :  CFOptionFlags,  _, character :  CFLocale!,  _, :  UnsafeMutablePointer< UTF32Char>!) ->  CFIndex) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.2+
//   - iPadOS 4.2+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringGetIntValue(str _, :  CFString!) ->  Int32) func

// CFStringGetLength(theString _, :  CFString!) ->  CFIndex) func


// CFStringGetLineBounds(theString _, range :  CFString!,  _, lineBeginIndex :  CFRange,  _, lineEndIndex :  UnsafeMutablePointer< CFIndex>!,  _, contentsEndIndex :  UnsafeMutablePointer< CFIndex>!,  _, :  UnsafeMutablePointer< CFIndex>!)) func

// CFStringGetListOfAvailableEncodings() func

// CFStringGetLongCharacterForSurrogatePair(surrogateHigh _, surrogateLow :  UniChar,  _, :  UniChar) ->  UTF32Char) func


// CFStringGetMaximumSizeForEncoding(length _, encoding :  CFIndex,  _, :  CFStringEncoding) ->  CFIndex) func

// CFStringGetMaximumSizeOfFileSystemRepresentation(string _, :  CFString!) ->  CFIndex) func

// CFStringGetMostCompatibleMacStringEncoding(encoding _, :  CFStringEncoding) ->  CFStringEncoding) func


// CFStringGetNameOfEncoding(encoding _, :  CFStringEncoding) ->  CFString!) func

// CFStringGetParagraphBounds(string _, range :  CFString!,  _, parBeginIndex :  CFRange,  _, parEndIndex :  UnsafeMutablePointer< CFIndex>!,  _, contentsEndIndex :  UnsafeMutablePointer< CFIndex>!,  _, :  UnsafeMutablePointer< CFIndex>!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringGetPascalString(theString _, buffer :  CFString!,  _, bufferSize :  StringPtr!,  _, encoding :  CFIndex,  _, :  CFStringEncoding) ->  Bool) func


// CFStringGetPascalStringPtr(theString _, encoding :  CFString!,  _, :  CFStringEncoding) ->  ConstStringPtr!) func

// CFStringGetRangeOfComposedCharactersAtIndex(theString _, theIndex :  CFString!,  _, :  CFIndex) ->  CFRange) func

// CFStringGetSmallestEncoding(theString _, :  CFString!) ->  CFStringEncoding) func


// CFStringGetSurrogatePairForLongCharacter(character _, surrogates :  UTF32Char,  _, :  UnsafeMutablePointer< UniChar>!) ->  Bool) func

// CFStringGetSystemEncoding() func

// CFStringGetTypeID() func


// CFStringHasPrefix(theString _, prefix :  CFString!,  _, :  CFString!) ->  Bool) func

// CFStringHasSuffix(theString _, suffix :  CFString!,  _, :  CFString!) ->  Bool) func

// CFStringInitInlineBuffer(str _, buf :  CFString!,  _, range :  UnsafeMutablePointer< CFStringInlineBuffer>!,  _, :  CFRange) func


// CFStringInsert(str _, idx :  CFMutableString!,  _, insertedStr :  CFIndex,  _, :  CFString!)) func

// CFStringIsEncodingAvailable(encoding _, :  CFStringEncoding) ->  Bool) func

// CFStringIsHyphenationAvailableForLocale(locale _, :  CFLocale!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.3+
//   - iPadOS 4.3+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFStringIsSurrogateHighCharacter(character _, :  UniChar) ->  Bool) func

// CFStringIsSurrogateLowCharacter(character _, :  UniChar) ->  Bool) func

// CFStringLowercase(theString _, locale :  CFMutableString!,  _, :  CFLocale!)) func


// CFStringNormalize(theString _, theForm :  CFMutableString!,  _, :  CFStringNormalizationForm) func

// CFStringPad(theString _, padString :  CFMutableString!,  _, length :  CFString!,  _, indexIntoPad :  CFIndex,  _, :  CFIndex) func

// CFStringReplace(theString _, range :  CFMutableString!,  _, replacement :  CFRange,  _, :  CFString!)) func


// CFStringReplaceAll(theString _, replacement :  CFMutableString!,  _, :  CFString!)) func

// CFStringSetExternalCharactersNoCopy(theString _, chars :  CFMutableString!,  _, length :  UnsafeMutablePointer< UniChar>!,  _, capacity :  CFIndex,  _, :  CFIndex) func

// CFStringTokenizerAdvanceToNextToken(tokenizer _, :  CFStringTokenizer!) ->  CFStringTokenizerTokenType) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFStringTokenizerCopyBestStringLanguage(string _, range :  CFString!,  _, :  CFRange) ->  CFString!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTokenizerCopyCurrentTokenAttribute(tokenizer _, attribute :  CFStringTokenizer!,  _, :  CFOptionFlags) ->  CFTypeRef!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTokenizerCreate(alloc _, string :  CFAllocator!,  _, range :  CFString!,  _, options :  CFRange,  _, locale :  CFOptionFlags,  _, :  CFLocale!) ->  CFStringTokenizer!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFStringTokenizerGetCurrentSubTokens(tokenizer _, ranges :  CFStringTokenizer!,  _, maxRangeLength :  UnsafeMutablePointer< CFRange>!,  _, derivedSubTokens :  CFIndex,  _, :  CFMutableArray!) ->  CFIndex) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTokenizerGetCurrentTokenRange(tokenizer _, :  CFStringTokenizer!) ->  CFRange) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTokenizerGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFStringTokenizerGoToTokenAtIndex(tokenizer _, index :  CFStringTokenizer!,  _, :  CFIndex) ->  CFStringTokenizerTokenType) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTokenizerSetString(tokenizer _, string :  CFStringTokenizer!,  _, range :  CFString!,  _, :  CFRange) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFStringTransform(string _, range :  CFMutableString!,  _, transform :  UnsafeMutablePointer< CFRange>!,  _, reverse :  CFString!,  _, :  Bool) ->  Bool) func


// CFStringTrim(theString _, trimString :  CFMutableString!,  _, :  CFString!)) func

// CFStringTrimWhitespace(theString _, :  CFMutableString!)) func

// CFStringUppercase(theString _, locale :  CFMutableString!,  _, :  CFLocale!)) func


// CFSwapInt16(arg _, :  UInt16) ->  UInt16) func

// CFSwapInt16BigToHost(arg _, :  UInt16) ->  UInt16) func

// CFSwapInt16HostToBig(arg _, :  UInt16) ->  UInt16) func


// CFSwapInt16HostToLittle(arg _, :  UInt16) ->  UInt16) func

// CFSwapInt16LittleToHost(arg _, :  UInt16) ->  UInt16) func

// CFSwapInt32(arg _, :  UInt32) ->  UInt32) func


// CFSwapInt32BigToHost(arg _, :  UInt32) ->  UInt32) func

// CFSwapInt32HostToBig(arg _, :  UInt32) ->  UInt32) func

// CFSwapInt32HostToLittle(arg _, :  UInt32) ->  UInt32) func


// CFSwapInt32LittleToHost(arg _, :  UInt32) ->  UInt32) func

// CFSwapInt64(arg _, :  UInt64) ->  UInt64) func

// CFSwapInt64BigToHost(arg _, :  UInt64) ->  UInt64) func


// CFSwapInt64HostToBig(arg _, :  UInt64) ->  UInt64) func

// CFSwapInt64HostToLittle(arg _, :  UInt64) ->  UInt64) func

// CFSwapInt64LittleToHost(arg _, :  UInt64) ->  UInt64) func


// CFTimeZoneCopyAbbreviation(tz _, at :  CFTimeZone!,  _, :  CFAbsoluteTime) ->  CFString!) func

// CFTimeZoneCopyAbbreviationDictionary() func

// CFTimeZoneCopyDefault() func


// CFTimeZoneCopyKnownNames() func

// CFTimeZoneCopyLocalizedName(tz _, style :  CFTimeZone!,  _, locale :  CFTimeZoneNameStyle,  _, :  CFLocale!) ->  CFString!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFTimeZoneCopySystem() func


// CFTimeZoneCreate(allocator _, name :  CFAllocator!,  _, data :  CFString!,  _, :  CFData!) ->  CFTimeZone!) func

// CFTimeZoneCreateWithName(allocator _, name :  CFAllocator!,  _, tryAbbrev :  CFString!,  _, :  Bool) ->  CFTimeZone!) func

// CFTimeZoneCreateWithTimeIntervalFromGMT(allocator _, ti :  CFAllocator!,  _, :  CFTimeInterval) ->  CFTimeZone!) func


// CFTimeZoneGetData(tz _, :  CFTimeZone!) ->  CFData!) func

// CFTimeZoneGetDaylightSavingTimeOffset(tz _, at :  CFTimeZone!,  _, :  CFAbsoluteTime) ->  CFTimeInterval) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFTimeZoneGetName(tz _, :  CFTimeZone!) ->  CFString!) func


// CFTimeZoneGetNextDaylightSavingTimeTransition(tz _, at :  CFTimeZone!,  _, :  CFAbsoluteTime) ->  CFAbsoluteTime) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFTimeZoneGetSecondsFromGMT(tz _, at :  CFTimeZone!,  _, :  CFAbsoluteTime) ->  CFTimeInterval) func

// CFTimeZoneGetTypeID() func


// CFTimeZoneIsDaylightSavingTime(tz _, at :  CFTimeZone!,  _, :  CFAbsoluteTime) ->  Bool) func

// CFTimeZoneResetSystem() func

// CFTimeZoneSetAbbreviationDictionary(dict _, :  CFDictionary!)) func


// CFTimeZoneSetDefault(tz _, :  CFTimeZone!)) func

// CFTreeAppendChild(tree _, newChild :  CFTree!,  _, :  CFTree!)) func

// CFTreeApplyFunctionToChildren(tree _, applier :  CFTree!,  _, context : (( UnsafeRawPointer?,  UnsafeMutableRawPointer?) ->  Void)!,  _, :  UnsafeMutableRawPointer!)) func


// CFTreeCreate(allocator _, context :  CFAllocator!,  _, :  UnsafePointer< CFTreeContext>!) ->  CFTree!) func

// CFTreeFindRoot(tree _, :  CFTree!) ->  CFTree!) func

// CFTreeGetChildAtIndex(tree _, idx :  CFTree!,  _, :  CFIndex) ->  CFTree!) func


// CFTreeGetChildCount(tree _, :  CFTree!) ->  CFIndex) func

// CFTreeGetChildren(tree _, children :  CFTree!,  _, :  UnsafeMutablePointer< Unmanaged< CFTree>?>!)) func

// CFTreeGetContext(tree _, context :  CFTree!,  _, :  UnsafeMutablePointer< CFTreeContext>!)) func


// CFTreeGetFirstChild(tree _, :  CFTree!) ->  CFTree!) func

// CFTreeGetNextSibling(tree _, :  CFTree!) ->  CFTree!) func

// CFTreeGetParent(tree _, :  CFTree!) ->  CFTree!) func


// CFTreeGetTypeID() func

// CFTreeInsertSibling(tree _, newSibling :  CFTree!,  _, :  CFTree!)) func

// CFTreePrependChild(tree _, newChild :  CFTree!,  _, :  CFTree!)) func


// CFTreeRemove(tree _, :  CFTree!)) func

// CFTreeRemoveAllChildren(tree _, :  CFTree!)) func

// CFTreeSetContext(tree _, context :  CFTree!,  _, :  UnsafePointer< CFTreeContext>!)) func


// CFTreeSortChildren(tree _, comparator :  CFTree!,  _, context :  CFComparatorFunction!,  _, :  UnsafeMutableRawPointer!)) func

// CFURLCanBeDecomposed(anURL _, :  CFURL!) ->  Bool) func

// CFURLClearResourcePropertyCache(url _, :  CFURL!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLClearResourcePropertyCacheForKey(url _, key :  CFURL!,  _, :  CFString!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCopyAbsoluteURL(relativeURL _, :  CFURL!) ->  CFURL!) func

// CFURLCopyFileSystemPath(anURL _, pathStyle :  CFURL!,  _, :  CFURLPathStyle) ->  CFString!) func


// CFURLCopyFragment(anURL _, charactersToLeaveEscaped :  CFURL!,  _, :  CFString!) ->  CFString!) func

// CFURLCopyHostName(anURL _, :  CFURL!) ->  CFString!) func

// CFURLCopyLastPathComponent(url _, :  CFURL!) ->  CFString!) func


// CFURLCopyNetLocation(anURL _, :  CFURL!) ->  CFString!) func

// CFURLCopyParameterString(anURL _, charactersToLeaveEscaped :  CFURL!,  _, :  CFString!) ->  CFString!) func
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

// CFURLCopyPassword(anURL _, :  CFURL!) ->  CFString!) func


// CFURLCopyPath(anURL _, :  CFURL!) ->  CFString!) func

// CFURLCopyPathExtension(url _, :  CFURL!) ->  CFString!) func

// CFURLCopyQueryString(anURL _, charactersToLeaveEscaped :  CFURL!,  _, :  CFString!) ->  CFString!) func


// CFURLCopyResourcePropertiesForKeys(url _, keys :  CFURL!,  _, error :  CFArray!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Unmanaged< CFDictionary>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCopyResourcePropertyForKey(url _, key :  CFURL!,  _, propertyValueTypeRefPtr :  CFString!,  _, error :  UnsafeMutableRawPointer!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCopyResourceSpecifier(anURL _, :  CFURL!) ->  CFString!) func


// CFURLCopyScheme(anURL _, :  CFURL!) ->  CFString!) func

// CFURLCopyStrictPath(anURL _, isAbsolute :  CFURL!,  _, :  UnsafeMutablePointer< DarwinBoolean>!) ->  CFString!) func

// CFURLCopyUserName(anURL _, :  CFURL!) ->  CFString!) func


// CFURLCreateAbsoluteURLWithBytes(alloc _, relativeURLBytes :  CFAllocator!,  _, length :  UnsafePointer< UInt8>!,  _, encoding :  CFIndex,  _, baseURL :  CFStringEncoding,  _, useCompatibilityMode :  CFURL!,  _, :  Bool) ->  CFURL!) func

// CFURLCreateBookmarkData(allocator _, url :  CFAllocator!,  _, options :  CFURL!,  _, resourcePropertiesToInclude :  CFURLBookmarkCreationOptions,  _, relativeToURL :  CFArray!,  _, error :  CFURL!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Unmanaged< CFData>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCreateBookmarkDataFromAliasRecord(allocatorRef _, aliasRecordDataRef :  CFAllocator!,  _, :  CFData!) ->  Unmanaged< CFData>!) func
//
// Availability:
//   - macOS 10.6+ (Deprecated in 11.0)
//
// Deprecated: This function is deprecated.


// CFURLCreateBookmarkDataFromFile(allocator _, fileURL :  CFAllocator!,  _, errorRef :  CFURL!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Unmanaged< CFData>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCreateByResolvingBookmarkData(allocator _, bookmark :  CFAllocator!,  _, options :  CFData!,  _, relativeToURL :  CFURLBookmarkResolutionOptions,  _, resourcePropertiesToInclude :  CFURL!,  _, isStale :  CFArray!,  _, error :  UnsafeMutablePointer< DarwinBoolean>!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Unmanaged< CFURL>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCreateCopyAppendingPathComponent(allocator _, url :  CFAllocator!,  _, pathComponent :  CFURL!,  _, isDirectory :  CFString!,  _, :  Bool) ->  CFURL!) func


// CFURLCreateCopyAppendingPathExtension(allocator _, url :  CFAllocator!,  _, extension :  CFURL!,  _, :  CFString!) ->  CFURL!) func

// CFURLCreateCopyDeletingLastPathComponent(allocator _, url :  CFAllocator!,  _, :  CFURL!) ->  CFURL!) func

// CFURLCreateCopyDeletingPathExtension(allocator _, url :  CFAllocator!,  _, :  CFURL!) ->  CFURL!) func


// CFURLCreateData(allocator _, url :  CFAllocator!,  _, encoding :  CFURL!,  _, escapeWhitespace :  CFStringEncoding,  _, :  Bool) ->  CFData!) func

// CFURLCreateDataAndPropertiesFromResource(alloc _, url :  CFAllocator!,  _, resourceData :  CFURL!,  _, properties :  UnsafeMutablePointer< Unmanaged< CFData>?>!,  _, desiredProperties :  UnsafeMutablePointer< Unmanaged< CFDictionary>?>!,  _, errorCode :  CFArray!,  _, :  UnsafeMutablePointer< Int32>!) ->  Bool) func
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLCreateFilePathURL(allocator _, url :  CFAllocator!,  _, error :  CFURL!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Unmanaged< CFURL>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLCreateFileReferenceURL(allocator _, url :  CFAllocator!,  _, error :  CFURL!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Unmanaged< CFURL>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCreateFromFSRef(allocator _, fsRef :  CFAllocator!,  _, :  OpaquePointer!) ->  CFURL!) func
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLCreateFromFileSystemRepresentation(allocator _, buffer :  CFAllocator!,  _, bufLen :  UnsafePointer< UInt8>!,  _, isDirectory :  CFIndex,  _, :  Bool) ->  CFURL!) func


// CFURLCreateFromFileSystemRepresentationRelativeToBase(allocator _, buffer :  CFAllocator!,  _, bufLen :  UnsafePointer< UInt8>!,  _, isDirectory :  CFIndex,  _, baseURL :  Bool,  _, :  CFURL!) ->  CFURL!) func

// CFURLCreatePropertyFromResource(alloc _, url :  CFAllocator!,  _, property :  CFURL!,  _, errorCode :  CFString!,  _, :  UnsafeMutablePointer< Int32>!) ->  CFTypeRef!) func
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLCreateResourcePropertiesForKeysFromBookmarkData(allocator _, resourcePropertiesToReturn :  CFAllocator!,  _, bookmark :  CFArray!,  _, :  CFData!) ->  Unmanaged< CFDictionary>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLCreateResourcePropertyForKeyFromBookmarkData(allocator _, resourcePropertyKey :  CFAllocator!,  _, bookmark :  CFString!,  _, :  CFData!) ->  Unmanaged< CFTypeRef>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLCreateStringByAddingPercentEscapes(allocator _, originalString :  CFAllocator!,  _, charactersToLeaveUnescaped :  CFString!,  _, legalURLCharactersToBeEscaped :  CFString!,  _, encoding :  CFString!,  _, :  CFStringEncoding) ->  CFString!) func
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

// CFURLCreateStringByReplacingPercentEscapes(allocator _, originalString :  CFAllocator!,  _, charactersToLeaveEscaped :  CFString!,  _, :  CFString!) ->  CFString!) func


// CFURLCreateStringByReplacingPercentEscapesUsingEncoding(allocator _, origString :  CFAllocator!,  _, charsToLeaveEscaped :  CFString!,  _, encoding :  CFString!,  _, :  CFStringEncoding) ->  CFString!) func
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

// CFURLCreateWithBytes(allocator _, URLBytes :  CFAllocator!,  _, length :  UnsafePointer< UInt8>!,  _, encoding :  CFIndex,  _, baseURL :  CFStringEncoding,  _, :  CFURL!) ->  CFURL!) func

// CFURLCreateWithFileSystemPath(allocator _, filePath :  CFAllocator!,  _, pathStyle :  CFString!,  _, isDirectory :  CFURLPathStyle,  _, :  Bool) ->  CFURL!) func


// CFURLCreateWithFileSystemPathRelativeToBase(allocator _, filePath :  CFAllocator!,  _, pathStyle :  CFString!,  _, isDirectory :  CFURLPathStyle,  _, baseURL :  Bool,  _, :  CFURL!) ->  CFURL!) func

// CFURLCreateWithString(allocator _, URLString :  CFAllocator!,  _, baseURL :  CFString!,  _, :  CFURL!) ->  CFURL!) func

// CFURLDestroyResource(url _, errorCode :  CFURL!,  _, :  UnsafeMutablePointer< Int32>!) ->  Bool) func
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.


// CFURLEnumeratorCreateForDirectoryURL(alloc _, directoryURL :  CFAllocator!,  _, option :  CFURL!,  _, propertyKeys :  CFURLEnumeratorOptions,  _, :  CFArray!) ->  CFURLEnumerator!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLEnumeratorCreateForMountedVolumes(alloc _, option :  CFAllocator!,  _, propertyKeys :  CFURLEnumeratorOptions,  _, :  CFArray!) ->  CFURLEnumerator!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLEnumeratorGetDescendentLevel(enumerator _, :  CFURLEnumerator!) ->  CFIndex) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLEnumeratorGetNextURL(enumerator _, url :  CFURLEnumerator!,  _, error :  UnsafeMutablePointer< Unmanaged< CFURL>?>!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  CFURLEnumeratorResult) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLEnumeratorGetSourceDidChange(enumerator _, :  CFURLEnumerator!) ->  Bool) func
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLEnumeratorGetTypeID() func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLEnumeratorSkipDescendents(enumerator _, :  CFURLEnumerator!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLGetBaseURL(anURL _, :  CFURL!) ->  CFURL!) func

// CFURLGetByteRangeForComponent(url _, component :  CFURL!,  _, rangeIncludingSeparators :  CFURLComponentType,  _, :  UnsafeMutablePointer< CFRange>!) ->  CFRange) func


// CFURLGetBytes(url _, buffer :  CFURL!,  _, bufferLength :  UnsafeMutablePointer< UInt8>!,  _, :  CFIndex) ->  CFIndex) func

// CFURLGetFSRef(url _, fsRef :  CFURL!,  _, :  OpaquePointer!) ->  Bool) func
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFURLGetFileSystemRepresentation(url _, resolveAgainstBase :  CFURL!,  _, buffer :  Bool,  _, maxBufLen :  UnsafeMutablePointer< UInt8>!,  _, :  CFIndex) ->  Bool) func


// CFURLGetPortNumber(anURL _, :  CFURL!) ->  Int32) func

// CFURLGetString(anURL _, :  CFURL!) ->  CFString!) func

// CFURLGetTypeID() func


// CFURLHasDirectoryPath(anURL _, :  CFURL!) ->  Bool) func

// CFURLIsFileReferenceURL(url _, :  CFURL!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLResourceIsReachable(url _, error :  CFURL!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLSetResourcePropertiesForKeys(url _, keyedPropertyValues :  CFURL!,  _, error :  CFDictionary!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLSetResourcePropertyForKey(url _, key :  CFURL!,  _, propertyValue :  CFString!,  _, error :  CFTypeRef!,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLSetTemporaryResourcePropertyForKey(url _, key :  CFURL!,  _, propertyValue :  CFString!,  _, :  CFTypeRef!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLStartAccessingSecurityScopedResource(url _, :  CFURL!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLStopAccessingSecurityScopedResource(url _, :  CFURL!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.7+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFURLWriteBookmarkDataToFile(bookmarkRef _, fileURL :  CFData!,  _, options :  CFURL!,  _, errorRef :  CFURLBookmarkFileCreationOptions,  _, :  UnsafeMutablePointer< Unmanaged< CFError>?>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - macOS 10.6+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+


// CFURLWriteDataAndPropertiesToResource(url _, dataToWrite :  CFURL!,  _, propertiesToWrite :  CFData!,  _, errorCode :  CFDictionary!,  _, :  UnsafeMutablePointer< Int32>!) ->  Bool) func
//
// Availability:
//   - tvOS 9.0+ (Deprecated in 9.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Deprecated: This function is deprecated.

// CFUUIDCreate(alloc _, :  CFAllocator!) ->  CFUUID!) func

// CFUUIDCreateFromString(alloc _, uuidStr :  CFAllocator!,  _, :  CFString!) ->  CFUUID!) func


// CFUUIDCreateFromUUIDBytes(alloc _, bytes :  CFAllocator!,  _, :  CFUUIDBytes) ->  CFUUID!) func

// CFUUIDCreateString(alloc _, uuid :  CFAllocator!,  _, :  CFUUID!) ->  CFString!) func

// CFUUIDCreateWithBytes(alloc _, byte0 :  CFAllocator!,  _, byte1 :  UInt8,  _, byte2 :  UInt8,  _, byte3 :  UInt8,  _, byte4 :  UInt8,  _, byte5 :  UInt8,  _, byte6 :  UInt8,  _, byte7 :  UInt8,  _, byte8 :  UInt8,  _, byte9 :  UInt8,  _, byte10 :  UInt8,  _, byte11 :  UInt8,  _, byte12 :  UInt8,  _, byte13 :  UInt8,  _, byte14 :  UInt8,  _, byte15 :  UInt8,  _, :  UInt8) ->  CFUUID!) func


// CFUUIDGetConstantUUIDWithBytes(alloc _, byte0 :  CFAllocator!,  _, byte1 :  UInt8,  _, byte2 :  UInt8,  _, byte3 :  UInt8,  _, byte4 :  UInt8,  _, byte5 :  UInt8,  _, byte6 :  UInt8,  _, byte7 :  UInt8,  _, byte8 :  UInt8,  _, byte9 :  UInt8,  _, byte10 :  UInt8,  _, byte11 :  UInt8,  _, byte12 :  UInt8,  _, byte13 :  UInt8,  _, byte14 :  UInt8,  _, byte15 :  UInt8,  _, :  UInt8) ->  CFUUID!) func

// CFUUIDGetTypeID() func

// CFUUIDGetUUIDBytes(uuid _, :  CFUUID!) ->  CFUUIDBytes) func


// CFUserNotificationCancel(userNotification _, :  CFUserNotification!) ->  Int32) func
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationCheckBoxChecked(i _, :  CFIndex) ->  CFOptionFlags) func
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationCreate(allocator _, timeout :  CFAllocator!,  _, flags :  CFTimeInterval,  _, error :  CFOptionFlags,  _, dictionary :  UnsafeMutablePointer< Int32>!,  _, :  CFDictionary!) ->  CFUserNotification!) func
//
// Availability:
//   - macOS 10.0+


// CFUserNotificationCreateRunLoopSource(allocator _, userNotification :  CFAllocator!,  _, callout :  CFUserNotification!,  _, order :  CFUserNotificationCallBack!,  _, :  CFIndex) ->  CFRunLoopSource!) func
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationDisplayAlert(timeout _, flags :  CFTimeInterval,  _, iconURL :  CFOptionFlags,  _, soundURL :  CFURL!,  _, localizationURL :  CFURL!,  _, alertHeader :  CFURL!,  _, alertMessage :  CFString!,  _, defaultButtonTitle :  CFString!,  _, alternateButtonTitle :  CFString!,  _, otherButtonTitle :  CFString!,  _, responseFlags :  CFString!,  _, :  UnsafeMutablePointer< CFOptionFlags>!) ->  Int32) func
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationDisplayNotice(timeout _, flags :  CFTimeInterval,  _, iconURL :  CFOptionFlags,  _, soundURL :  CFURL!,  _, localizationURL :  CFURL!,  _, alertHeader :  CFURL!,  _, alertMessage :  CFString!,  _, defaultButtonTitle :  CFString!,  _, :  CFString!) ->  Int32) func
//
// Availability:
//   - macOS 10.0+


// CFUserNotificationGetResponseDictionary(userNotification _, :  CFUserNotification!) ->  CFDictionary!) func
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationGetResponseValue(userNotification _, key :  CFUserNotification!,  _, idx :  CFString!,  _, :  CFIndex) ->  CFString!) func
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationGetTypeID() func
//
// Availability:
//   - macOS 10.0+


// CFUserNotificationPopUpSelection(n _, :  CFIndex) ->  CFOptionFlags) func
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationReceiveResponse(userNotification _, timeout :  CFUserNotification!,  _, responseFlags :  CFTimeInterval,  _, :  UnsafeMutablePointer< CFOptionFlags>!) ->  Int32) func
//
// Availability:
//   - macOS 10.0+

// CFUserNotificationSecureTextField(i _, :  CFIndex) ->  CFOptionFlags) func
//
// Availability:
//   - macOS 10.0+


// CFUserNotificationUpdate(userNotification _, timeout :  CFUserNotification!,  _, flags :  CFTimeInterval,  _, dictionary :  CFOptionFlags,  _, :  CFDictionary!) ->  Int32) func
//
// Availability:
//   - macOS 10.0+

// CFWriteStreamCanAcceptBytes(stream _, :  CFWriteStream!) ->  Bool) func

// CFWriteStreamClose(stream _, :  CFWriteStream!)) func


// CFWriteStreamCopyDispatchQueue(stream _, :  CFWriteStream!) ->  dispatch_queue_t!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFWriteStreamCopyError(stream _, :  CFWriteStream!) ->  CFError!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFWriteStreamCopyProperty(stream _, propertyName :  CFWriteStream!,  _, :  CFStreamPropertyKey!) ->  CFTypeRef!) func


// CFWriteStreamCreateWithAllocatedBuffers(alloc _, bufferAllocator :  CFAllocator!,  _, :  CFAllocator!) ->  CFWriteStream!) func

// CFWriteStreamCreateWithBuffer(alloc _, buffer :  CFAllocator!,  _, bufferCapacity :  UnsafeMutablePointer< UInt8>!,  _, :  CFIndex) ->  CFWriteStream!) func

// CFWriteStreamCreateWithFile(alloc _, fileURL :  CFAllocator!,  _, :  CFURL!) ->  CFWriteStream!) func


// CFWriteStreamGetError(stream _, :  CFWriteStream!) ->  CFStreamError) func

// CFWriteStreamGetStatus(stream _, :  CFWriteStream!) ->  CFStreamStatus) func

// CFWriteStreamGetTypeID() func


// CFWriteStreamOpen(stream _, :  CFWriteStream!) ->  Bool) func

// CFWriteStreamScheduleWithRunLoop(stream _, runLoop :  CFWriteStream!,  _, runLoopMode :  CFRunLoop!,  _, :  CFRunLoopMode!)) func

// CFWriteStreamSetClient(stream _, streamEvents :  CFWriteStream!,  _, clientCB :  CFOptionFlags,  _, clientContext :  CFWriteStreamClientCallBack!,  _, :  UnsafeMutablePointer< CFStreamClientContext>!) ->  Bool) func


// CFWriteStreamSetDispatchQueue(stream _, q :  CFWriteStream!,  _, :  dispatch_queue_t!)) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+

// CFWriteStreamSetProperty(stream _, propertyName :  CFWriteStream!,  _, propertyValue :  CFStreamPropertyKey!,  _, :  CFTypeRef!) ->  Bool) func

// CFWriteStreamUnscheduleFromRunLoop(stream _, runLoop :  CFWriteStream!,  _, runLoopMode :  CFRunLoop!,  _, :  CFRunLoopMode!)) func


// CFWriteStreamWrite(stream _, buffer :  CFWriteStream!,  _, bufferLength :  UnsafePointer< UInt8>!,  _, :  CFIndex) ->  CFIndex) func

// CFXMLCreateStringByEscapingEntities(allocator _, string :  CFAllocator!,  _, entitiesDictionary :  CFString!,  _, :  CFDictionary!) ->  CFString!) func

// CFXMLCreateStringByUnescapingEntities(allocator _, string :  CFAllocator!,  _, entitiesDictionary :  CFString!,  _, :  CFDictionary!) ->  CFString!) func


// CFXMLNodeCreate(alloc CFAllocatorRef, xmlType ,  CFXMLNodeTypeCode, dataString ,  CFStringRef, additionalInfoPtr ,  const  void *, version ,  CFIndex, );) extern   CFXMLNodeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLNodeCreateCopy(alloc CFAllocatorRef, origNode ,  CFXMLNodeRef, );) extern   CFXMLNodeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLNodeGetInfoPtr(node CFXMLNodeRef, );) extern   const   void  *
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLNodeGetString(node CFXMLNodeRef, );) extern   CFStringRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLNodeGetTypeCode(node CFXMLNodeRef, );) extern   CFXMLNodeTypeCode
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLNodeGetTypeID() extern   CFTypeID
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLNodeGetVersion(node CFXMLNodeRef, );) extern   CFIndex
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserAbort(parser CFXMLParserRef, errorCode ,  CFXMLParserStatusCode, errorDescription ,  CFStringRef, );) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserCopyErrorDescription(parser CFXMLParserRef, );) extern   CFStringRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLParserCreate(allocator CFAllocatorRef, xmlData ,  CFDataRef, dataSource ,  CFURLRef, parseOptions ,  CFOptionFlags, versionOfNodes ,  CFIndex, callBacks ,  CFXMLParserCallBacks *, context ,  CFXMLParserContext *, );) extern   CFXMLParserRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserCreateWithDataFromURL(allocator CFAllocatorRef, dataSource ,  CFURLRef, parseOptions ,  CFOptionFlags, versionOfNodes ,  CFIndex, callBacks ,  CFXMLParserCallBacks *, context ,  CFXMLParserContext *, );) extern   CFXMLParserRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserGetCallBacks(parser CFXMLParserRef, callBacks ,  CFXMLParserCallBacks *, );) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLParserGetContext(parser CFXMLParserRef, context ,  CFXMLParserContext *, );) extern   void
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserGetDocument(parser CFXMLParserRef, );) extern   void  *
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserGetLineNumber(parser CFXMLParserRef, );) extern   CFIndex
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLParserGetLocation(parser CFXMLParserRef, );) extern   CFIndex
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserGetSourceURL(parser CFXMLParserRef, );) extern   CFURLRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserGetStatusCode(parser CFXMLParserRef, );) extern   CFXMLParserStatusCode
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLParserGetTypeID() extern   CFTypeID
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLParserParse(parser CFXMLParserRef, );) extern   Boolean
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLTreeCreateFromData(allocator CFAllocatorRef, xmlData ,  CFDataRef, dataSource ,  CFURLRef, parseOptions ,  CFOptionFlags, versionOfNodes ,  CFIndex, );) extern   CFXMLTreeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLTreeCreateFromDataWithError(allocator CFAllocatorRef, xmlData ,  CFDataRef, dataSource ,  CFURLRef, parseOptions ,  CFOptionFlags, versionOfNodes ,  CFIndex, errorDict ,  CFDictionaryRef *, );) extern   CFXMLTreeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLTreeCreateWithDataFromURL(allocator CFAllocatorRef, dataSource ,  CFURLRef, parseOptions ,  CFOptionFlags, versionOfNodes ,  CFIndex, );) extern   CFXMLTreeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLTreeCreateWithNode(allocator CFAllocatorRef, node ,  CFXMLNodeRef, );) extern   CFXMLTreeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// CFXMLTreeCreateXMLData(allocator CFAllocatorRef, xmlTree ,  CFXMLTreeRef, );) extern   CFDataRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// CFXMLTreeGetNode(xmlTree CFXMLTreeRef, );) extern   CFXMLNodeRef
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

