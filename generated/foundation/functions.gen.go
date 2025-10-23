// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// Foundation Functions (172 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSAllocateObject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSClassFromString func(unsafe.Pointer) unsafe.Pointer
	_NSCountFrames func() unsafe.Pointer
	_NSFrameAddress func(unsafe.Pointer) unsafe.Pointer
	_NSGetSizeAndAlignment func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
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
	_AXPrefersActionSliderAlternative func() bool
	_AXShowBordersEnabled func() bool
	_CFAttributedStringCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringCreateWithSubstring func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttribute func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttributeAndLongestEffectiveRange func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttributes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetAttributesAndLongestEffectiveRange func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetLength func(unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetString func(unsafe.Pointer) unsafe.Pointer
	_CFAttributedStringGetTypeID func() unsafe.Pointer
	_CFDataCreate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreateCopy func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataCreateWithBytesNoCopy func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataFind func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFDataGetBytePtr func(unsafe.Pointer) unsafe.Pointer
	_CFDataGetBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_CFDataGetLength func(unsafe.Pointer) unsafe.Pointer
	_CFDataGetTypeID func() unsafe.Pointer
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
	_CFURLCreateFilePathURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateFileReferenceURL func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateFromFSRef func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateFromFileSystemRepresentation func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateFromFileSystemRepresentationRelativeToBase func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateResourcePropertiesForKeysFromBookmarkData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateResourcePropertyForKeyFromBookmarkData func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateStringByAddingPercentEscapes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateStringByReplacingPercentEscapes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateStringByReplacingPercentEscapesUsingEncoding func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateWithBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateWithFileSystemPath func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateWithFileSystemPathRelativeToBase func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFURLCreateWithString func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
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
	_CFUUIDCreate func(unsafe.Pointer) unsafe.Pointer
	_CFUUIDCreateFromString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDCreateFromUUIDBytes func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDCreateString func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDCreateWithBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDGetConstantUUIDWithBytes func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_CFUUIDGetTypeID func() unsafe.Pointer
	_CFUUIDGetUUIDBytes func(unsafe.Pointer) unsafe.Pointer
	_inset func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_ceil func(float64) float64
	_UIApplicationMain func(int, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) int
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
	tryRegister(&_AXPrefersActionSliderAlternative, lib, "AXPrefersActionSliderAlternative")
	tryRegister(&_AXShowBordersEnabled, lib, "AXShowBordersEnabled")
	tryRegister(&_CFAttributedStringCreate, lib, "CFAttributedStringCreate")
	tryRegister(&_CFAttributedStringCreateCopy, lib, "CFAttributedStringCreateCopy")
	tryRegister(&_CFAttributedStringCreateWithSubstring, lib, "CFAttributedStringCreateWithSubstring")
	tryRegister(&_CFAttributedStringGetAttribute, lib, "CFAttributedStringGetAttribute")
	tryRegister(&_CFAttributedStringGetAttributeAndLongestEffectiveRange, lib, "CFAttributedStringGetAttributeAndLongestEffectiveRange")
	tryRegister(&_CFAttributedStringGetAttributes, lib, "CFAttributedStringGetAttributes")
	tryRegister(&_CFAttributedStringGetAttributesAndLongestEffectiveRange, lib, "CFAttributedStringGetAttributesAndLongestEffectiveRange")
	tryRegister(&_CFAttributedStringGetLength, lib, "CFAttributedStringGetLength")
	tryRegister(&_CFAttributedStringGetString, lib, "CFAttributedStringGetString")
	tryRegister(&_CFAttributedStringGetTypeID, lib, "CFAttributedStringGetTypeID")
	tryRegister(&_CFDataCreate, lib, "CFDataCreate")
	tryRegister(&_CFDataCreateCopy, lib, "CFDataCreateCopy")
	tryRegister(&_CFDataCreateWithBytesNoCopy, lib, "CFDataCreateWithBytesNoCopy")
	tryRegister(&_CFDataFind, lib, "CFDataFind")
	tryRegister(&_CFDataGetBytePtr, lib, "CFDataGetBytePtr")
	tryRegister(&_CFDataGetBytes, lib, "CFDataGetBytes")
	tryRegister(&_CFDataGetLength, lib, "CFDataGetLength")
	tryRegister(&_CFDataGetTypeID, lib, "CFDataGetTypeID")
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
	tryRegister(&_CFURLCreateFilePathURL, lib, "CFURLCreateFilePathURL")
	tryRegister(&_CFURLCreateFileReferenceURL, lib, "CFURLCreateFileReferenceURL")
	tryRegister(&_CFURLCreateFromFSRef, lib, "CFURLCreateFromFSRef")
	tryRegister(&_CFURLCreateFromFileSystemRepresentation, lib, "CFURLCreateFromFileSystemRepresentation")
	tryRegister(&_CFURLCreateFromFileSystemRepresentationRelativeToBase, lib, "CFURLCreateFromFileSystemRepresentationRelativeToBase")
	tryRegister(&_CFURLCreateResourcePropertiesForKeysFromBookmarkData, lib, "CFURLCreateResourcePropertiesForKeysFromBookmarkData")
	tryRegister(&_CFURLCreateResourcePropertyForKeyFromBookmarkData, lib, "CFURLCreateResourcePropertyForKeyFromBookmarkData")
	tryRegister(&_CFURLCreateStringByAddingPercentEscapes, lib, "CFURLCreateStringByAddingPercentEscapes")
	tryRegister(&_CFURLCreateStringByReplacingPercentEscapes, lib, "CFURLCreateStringByReplacingPercentEscapes")
	tryRegister(&_CFURLCreateStringByReplacingPercentEscapesUsingEncoding, lib, "CFURLCreateStringByReplacingPercentEscapesUsingEncoding")
	tryRegister(&_CFURLCreateWithBytes, lib, "CFURLCreateWithBytes")
	tryRegister(&_CFURLCreateWithFileSystemPath, lib, "CFURLCreateWithFileSystemPath")
	tryRegister(&_CFURLCreateWithFileSystemPathRelativeToBase, lib, "CFURLCreateWithFileSystemPathRelativeToBase")
	tryRegister(&_CFURLCreateWithString, lib, "CFURLCreateWithString")
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
	tryRegister(&_CFUUIDCreate, lib, "CFUUIDCreate")
	tryRegister(&_CFUUIDCreateFromString, lib, "CFUUIDCreateFromString")
	tryRegister(&_CFUUIDCreateFromUUIDBytes, lib, "CFUUIDCreateFromUUIDBytes")
	tryRegister(&_CFUUIDCreateString, lib, "CFUUIDCreateString")
	tryRegister(&_CFUUIDCreateWithBytes, lib, "CFUUIDCreateWithBytes")
	tryRegister(&_CFUUIDGetConstantUUIDWithBytes, lib, "CFUUIDGetConstantUUIDWithBytes")
	tryRegister(&_CFUUIDGetTypeID, lib, "CFUUIDGetTypeID")
	tryRegister(&_CFUUIDGetUUIDBytes, lib, "CFUUIDGetUUIDBytes")
	tryRegister(&_inset, lib, "inset")
	tryRegister(&_ceil, lib, "ceil")
	tryRegister(&_UIApplicationMain, lib, "UIApplicationMain")
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

// AXPrefersActionSliderAlternative is a Foundation function.
//
// Added in macOS 26.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXPrefersActionSliderAlternative
func AXPrefersActionSliderAlternative() bool {
	return _AXPrefersActionSliderAlternative()
}

// AXShowBordersEnabled is a Foundation function.
//
// Added in macOS 26.1.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXShowBordersEnabled
func AXShowBordersEnabled() bool {
	return _AXShowBordersEnabled()
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

// Creates a sub-attributed string from the specified range.

// Creates a sub-attributed string from the specified range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedStringCreateWithSubstring(_:_:_:)
func CFAttributedStringCreateWithSubstring(alloc unsafe.Pointer, aStr unsafe.Pointer, range_ unsafe.Pointer) unsafe.Pointer {
	return _CFAttributedStringCreateWithSubstring(alloc, aStr, range_)
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

// Creates an immutable CFData object from an external (client-owned) byte buffer.

// Creates an immutable CFData object from an external (client-owned) byte buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataCreateWithBytesNoCopy(_:_:_:_:)
func CFDataCreateWithBytesNoCopy(allocator unsafe.Pointer, bytes unsafe.Pointer, length unsafe.Pointer, bytesDeallocator unsafe.Pointer) unsafe.Pointer {
	return _CFDataCreateWithBytesNoCopy(allocator, bytes, length, bytesDeallocator)
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

// Returns the type identifier for the CFData opaque type.

// Returns the type identifier for the CFData opaque type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDataGetTypeID()
func CFDataGetTypeID() unsafe.Pointer {
	return _CFDataGetTypeID()
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

// inset is a Foundation function.

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CGRect/inset(by:)
func inset(insets unsafe.Pointer, p1 unsafe.Pointer) unsafe.Pointer {
	return _inset(insets, p1)
}

// ceil is a Foundation function.
//
// Added in macOS 10.10.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/kernel/1557272-ceil
func ceil(p0 float64) float64 {
	return _ceil(p0)
}

// Creates the application object and the application delegate and sets up the event cycle.

// Creates the application object and the application delegate and sets up the event cycle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplicationMain(_:_:_:_:)-1yub7
func UIApplicationMain(argc int, argv unsafe.Pointer, principalClassName unsafe.Pointer, delegateClassName unsafe.Pointer) int {
	return _UIApplicationMain(argc, argv, principalClassName, delegateClassName)
}



