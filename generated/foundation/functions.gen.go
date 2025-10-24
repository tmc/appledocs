// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

/* debug [functions.gen.go]: Generating 123 functions for Foundation */
import (
	"unsafe"

	"github.com/ebitengine/purego"
	objc "github.com/ebitengine/purego/objc"
)


// Foundation Functions (123 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSAllHashTableObjects func(unsafe.Pointer) unsafe.Pointer
	_NSAllMapTableKeys func(unsafe.Pointer) unsafe.Pointer
	_NSAllMapTableValues func(unsafe.Pointer) unsafe.Pointer
	_NSAllocateCollectable func(uint, uint) unsafe.Pointer
	_NSAllocateMemoryPages func(uint) unsafe.Pointer
	_NSAllocateObject func(objc.Class, uint, unsafe.Pointer) objc.ID
	_NSClassFromString func(unsafe.Pointer) objc.Class
	_NSCompareHashTables func(unsafe.Pointer, unsafe.Pointer) bool
	_NSCompareMapTables func(unsafe.Pointer, unsafe.Pointer) bool
	_NSContainsRect func(Rect, Rect) bool
	_NSCopyHashTableWithZone func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCopyMapTableWithZone func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSCopyMemoryPages func(unsafe.Pointer, unsafe.Pointer, uint)
	_NSCopyObject func(objc.ID, uint, unsafe.Pointer) objc.ID
	_NSCountFrames func() uint
	_NSCountHashTable func(unsafe.Pointer) uint
	_NSCountMapTable func(unsafe.Pointer) uint
	_NSCreateHashTable func(HashTableCallBacks, uint) unsafe.Pointer
	_NSCreateHashTableWithZone func(HashTableCallBacks, uint, unsafe.Pointer) unsafe.Pointer
	_NSCreateMapTable func(MapTableKeyCallBacks, MapTableValueCallBacks, uint) unsafe.Pointer
	_NSCreateMapTableWithZone func(MapTableKeyCallBacks, MapTableValueCallBacks, uint, unsafe.Pointer) unsafe.Pointer
	_NSCreateZone func(uint, uint, bool) unsafe.Pointer
	_NSDeallocateMemoryPages func(unsafe.Pointer, uint)
	_NSDeallocateObject func(objc.ID)
	_NSDecimalAdd func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, RoundingMode) CalculationError
	_NSDecimalCompact func(unsafe.Pointer)
	_NSDecimalCompare func(unsafe.Pointer, unsafe.Pointer) ComparisonResult
	_NSDecimalCopy func(unsafe.Pointer, unsafe.Pointer)
	_NSDecimalDivide func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, RoundingMode) CalculationError
	_NSDecimalMultiply func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, RoundingMode) CalculationError
	_NSDecimalMultiplyByPowerOf10 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, RoundingMode) CalculationError
	_NSDecimalNormalize func(unsafe.Pointer, unsafe.Pointer, RoundingMode) CalculationError
	_NSDecimalPower func(unsafe.Pointer, unsafe.Pointer, uint, RoundingMode) CalculationError
	_NSDecimalRound func(unsafe.Pointer, unsafe.Pointer, int, RoundingMode)
	_NSDecimalString func(unsafe.Pointer, objc.ID) unsafe.Pointer
	_NSDecimalSubtract func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, RoundingMode) CalculationError
	_NSDecrementExtraRefCountWasZero func(objc.ID) bool
	_NSDefaultMallocZone func() unsafe.Pointer
	_NSDivideRect func(Rect, unsafe.Pointer, unsafe.Pointer, float64, RectEdge)
	_NSEdgeInsetsEqual func(EdgeInsets, EdgeInsets) bool
	_NSEndHashTableEnumeration func(unsafe.Pointer)
	_NSEndMapTableEnumeration func(unsafe.Pointer)
	_NSEnumerateHashTable func(unsafe.Pointer) HashEnumerator
	_NSEnumerateMapTable func(unsafe.Pointer) MapEnumerator
	_NSEqualPoints func(Point, Point) bool
	_NSEqualRects func(Rect, Rect) bool
	_NSEqualSizes func(Size, Size) bool
	_NSExtraRefCount func(objc.ID) uint
	_NSFrameAddress func(uint) unsafe.Pointer
	_NSFreeHashTable func(unsafe.Pointer)
	_NSFreeMapTable func(unsafe.Pointer)
	_NSFullUserName func() unsafe.Pointer
	_NSGetSizeAndAlignment func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSHashGet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSHashInsert func(unsafe.Pointer, unsafe.Pointer)
	_NSHashInsertIfAbsent func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSHashInsertKnownAbsent func(unsafe.Pointer, unsafe.Pointer)
	_NSHashRemove func(unsafe.Pointer, unsafe.Pointer)
	_NSHomeDirectory func() unsafe.Pointer
	_NSHomeDirectoryForUser func(unsafe.Pointer) unsafe.Pointer
	_NSIncrementExtraRefCount func(objc.ID)
	_NSInsetRect func(Rect, float64, float64) Rect
	_NSIntegralRect func(Rect) Rect
	_NSIntegralRectWithOptions func(Rect, AlignmentOptions) Rect
	_NSIntersectionRange func(Range, Range) Range
	_NSIntersectionRect func(Rect, Rect) Rect
	_NSIntersectsRect func(Rect, Rect) bool
	_NSIsEmptyRect func(Rect) bool
	_NSIsFreedObject func(objc.ID) bool
	_NSLog func(unsafe.Pointer)
	_NSLogPageSize func() uint
	_NSLogv func(unsafe.Pointer, unsafe.Pointer)
	_NSMapGet func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSMapInsert func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_NSMapInsertIfAbsent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_NSMapInsertKnownAbsent func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	_NSMapMember func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_NSMapRemove func(unsafe.Pointer, unsafe.Pointer)
	_NSMouseInRect func(Point, Rect, bool) bool
	_NSNextHashEnumeratorItem func(unsafe.Pointer) unsafe.Pointer
	_NSNextMapEnumeratorPair func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) bool
	_NSOffsetRect func(Rect, float64, float64) Rect
	_NSOpenStepRootDirectory func() unsafe.Pointer
	_NSPageSize func() uint
	_NSPointFromString func(unsafe.Pointer) Point
	_NSPointInRect func(Point, Rect) bool
	_NSProtocolFromString func(unsafe.Pointer) unsafe.Pointer
	_NSRangeFromString func(unsafe.Pointer) Range
	_NSRealMemoryAvailable func() uint
	_NSReallocateCollectable func(unsafe.Pointer, uint, uint) unsafe.Pointer
	_NSRecordAllocationEvent func(int, objc.ID)
	_NSRectFromString func(unsafe.Pointer) Rect
	_NSRecycleZone func(unsafe.Pointer)
	_NSResetHashTable func(unsafe.Pointer)
	_NSResetMapTable func(unsafe.Pointer)
	_NSReturnAddress func(uint) unsafe.Pointer
	_NSRoundDownToMultipleOfPageSize func(uint) uint
	_NSRoundUpToMultipleOfPageSize func(uint) uint
	_NSSearchPathForDirectoriesInDomains func(SearchPathDirectory, SearchPathDomainMask, bool) []unsafe.Pointer
	_NSSelectorFromString func(unsafe.Pointer) objc.SEL
	_NSSetZoneName func(unsafe.Pointer, unsafe.Pointer)
	_NSShouldRetainWithZone func(objc.ID, unsafe.Pointer) bool
	_NSSizeFromString func(unsafe.Pointer) Size
	_NSStringFromClass func(objc.Class) unsafe.Pointer
	_NSStringFromHashTable func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromMapTable func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromPoint func(Point) unsafe.Pointer
	_NSStringFromProtocol func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromRange func(Range) unsafe.Pointer
	_NSStringFromRect func(Rect) unsafe.Pointer
	_NSStringFromSelector func(objc.SEL) unsafe.Pointer
	_NSStringFromSize func(Size) unsafe.Pointer
	_NSTemporaryDirectory func() unsafe.Pointer
	_NSUnionRange func(Range, Range) Range
	_NSUnionRect func(Rect, Rect) Rect
	_NSUserName func() unsafe.Pointer
	_NSZoneCalloc func(unsafe.Pointer, uint, uint) unsafe.Pointer
	_NSZoneFree func(unsafe.Pointer, unsafe.Pointer)
	_NSZoneFromPointer func(unsafe.Pointer) unsafe.Pointer
	_NSZoneMalloc func(unsafe.Pointer, uint) unsafe.Pointer
	_NSZoneName func(unsafe.Pointer) unsafe.Pointer
	_NSZoneRealloc func(unsafe.Pointer, unsafe.Pointer, uint) unsafe.Pointer
	_NXReadNSObjectFromCoder func(unsafe.Pointer) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSAllHashTableObjects, lib, "NSAllHashTableObjects")
	tryRegister(&_NSAllMapTableKeys, lib, "NSAllMapTableKeys")
	tryRegister(&_NSAllMapTableValues, lib, "NSAllMapTableValues")
	tryRegister(&_NSAllocateCollectable, lib, "NSAllocateCollectable")
	tryRegister(&_NSAllocateMemoryPages, lib, "NSAllocateMemoryPages")
	tryRegister(&_NSAllocateObject, lib, "NSAllocateObject")
	tryRegister(&_NSClassFromString, lib, "NSClassFromString")
	tryRegister(&_NSCompareHashTables, lib, "NSCompareHashTables")
	tryRegister(&_NSCompareMapTables, lib, "NSCompareMapTables")
	tryRegister(&_NSContainsRect, lib, "NSContainsRect")
	tryRegister(&_NSCopyHashTableWithZone, lib, "NSCopyHashTableWithZone")
	tryRegister(&_NSCopyMapTableWithZone, lib, "NSCopyMapTableWithZone")
	tryRegister(&_NSCopyMemoryPages, lib, "NSCopyMemoryPages")
	tryRegister(&_NSCopyObject, lib, "NSCopyObject")
	tryRegister(&_NSCountFrames, lib, "NSCountFrames")
	tryRegister(&_NSCountHashTable, lib, "NSCountHashTable")
	tryRegister(&_NSCountMapTable, lib, "NSCountMapTable")
	tryRegister(&_NSCreateHashTable, lib, "NSCreateHashTable")
	tryRegister(&_NSCreateHashTableWithZone, lib, "NSCreateHashTableWithZone")
	tryRegister(&_NSCreateMapTable, lib, "NSCreateMapTable")
	tryRegister(&_NSCreateMapTableWithZone, lib, "NSCreateMapTableWithZone")
	tryRegister(&_NSCreateZone, lib, "NSCreateZone")
	tryRegister(&_NSDeallocateMemoryPages, lib, "NSDeallocateMemoryPages")
	tryRegister(&_NSDeallocateObject, lib, "NSDeallocateObject")
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
	tryRegister(&_NSDecrementExtraRefCountWasZero, lib, "NSDecrementExtraRefCountWasZero")
	tryRegister(&_NSDefaultMallocZone, lib, "NSDefaultMallocZone")
	tryRegister(&_NSDivideRect, lib, "NSDivideRect")
	tryRegister(&_NSEdgeInsetsEqual, lib, "NSEdgeInsetsEqual")
	tryRegister(&_NSEndHashTableEnumeration, lib, "NSEndHashTableEnumeration")
	tryRegister(&_NSEndMapTableEnumeration, lib, "NSEndMapTableEnumeration")
	tryRegister(&_NSEnumerateHashTable, lib, "NSEnumerateHashTable")
	tryRegister(&_NSEnumerateMapTable, lib, "NSEnumerateMapTable")
	tryRegister(&_NSEqualPoints, lib, "NSEqualPoints")
	tryRegister(&_NSEqualRects, lib, "NSEqualRects")
	tryRegister(&_NSEqualSizes, lib, "NSEqualSizes")
	tryRegister(&_NSExtraRefCount, lib, "NSExtraRefCount")
	tryRegister(&_NSFrameAddress, lib, "NSFrameAddress")
	tryRegister(&_NSFreeHashTable, lib, "NSFreeHashTable")
	tryRegister(&_NSFreeMapTable, lib, "NSFreeMapTable")
	tryRegister(&_NSFullUserName, lib, "NSFullUserName")
	tryRegister(&_NSGetSizeAndAlignment, lib, "NSGetSizeAndAlignment")
	tryRegister(&_NSHashGet, lib, "NSHashGet")
	tryRegister(&_NSHashInsert, lib, "NSHashInsert")
	tryRegister(&_NSHashInsertIfAbsent, lib, "NSHashInsertIfAbsent")
	tryRegister(&_NSHashInsertKnownAbsent, lib, "NSHashInsertKnownAbsent")
	tryRegister(&_NSHashRemove, lib, "NSHashRemove")
	tryRegister(&_NSHomeDirectory, lib, "NSHomeDirectory")
	tryRegister(&_NSHomeDirectoryForUser, lib, "NSHomeDirectoryForUser")
	tryRegister(&_NSIncrementExtraRefCount, lib, "NSIncrementExtraRefCount")
	tryRegister(&_NSInsetRect, lib, "NSInsetRect")
	tryRegister(&_NSIntegralRect, lib, "NSIntegralRect")
	tryRegister(&_NSIntegralRectWithOptions, lib, "NSIntegralRectWithOptions")
	tryRegister(&_NSIntersectionRange, lib, "NSIntersectionRange")
	tryRegister(&_NSIntersectionRect, lib, "NSIntersectionRect")
	tryRegister(&_NSIntersectsRect, lib, "NSIntersectsRect")
	tryRegister(&_NSIsEmptyRect, lib, "NSIsEmptyRect")
	tryRegister(&_NSIsFreedObject, lib, "NSIsFreedObject")
	tryRegister(&_NSLog, lib, "NSLog")
	tryRegister(&_NSLogPageSize, lib, "NSLogPageSize")
	tryRegister(&_NSLogv, lib, "NSLogv")
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
	tryRegister(&_NSOpenStepRootDirectory, lib, "NSOpenStepRootDirectory")
	tryRegister(&_NSPageSize, lib, "NSPageSize")
	tryRegister(&_NSPointFromString, lib, "NSPointFromString")
	tryRegister(&_NSPointInRect, lib, "NSPointInRect")
	tryRegister(&_NSProtocolFromString, lib, "NSProtocolFromString")
	tryRegister(&_NSRangeFromString, lib, "NSRangeFromString")
	tryRegister(&_NSRealMemoryAvailable, lib, "NSRealMemoryAvailable")
	tryRegister(&_NSReallocateCollectable, lib, "NSReallocateCollectable")
	tryRegister(&_NSRecordAllocationEvent, lib, "NSRecordAllocationEvent")
	tryRegister(&_NSRectFromString, lib, "NSRectFromString")
	tryRegister(&_NSRecycleZone, lib, "NSRecycleZone")
	tryRegister(&_NSResetHashTable, lib, "NSResetHashTable")
	tryRegister(&_NSResetMapTable, lib, "NSResetMapTable")
	tryRegister(&_NSReturnAddress, lib, "NSReturnAddress")
	tryRegister(&_NSRoundDownToMultipleOfPageSize, lib, "NSRoundDownToMultipleOfPageSize")
	tryRegister(&_NSRoundUpToMultipleOfPageSize, lib, "NSRoundUpToMultipleOfPageSize")
	tryRegister(&_NSSearchPathForDirectoriesInDomains, lib, "NSSearchPathForDirectoriesInDomains")
	tryRegister(&_NSSelectorFromString, lib, "NSSelectorFromString")
	tryRegister(&_NSSetZoneName, lib, "NSSetZoneName")
	tryRegister(&_NSShouldRetainWithZone, lib, "NSShouldRetainWithZone")
	tryRegister(&_NSSizeFromString, lib, "NSSizeFromString")
	tryRegister(&_NSStringFromClass, lib, "NSStringFromClass")
	tryRegister(&_NSStringFromHashTable, lib, "NSStringFromHashTable")
	tryRegister(&_NSStringFromMapTable, lib, "NSStringFromMapTable")
	tryRegister(&_NSStringFromPoint, lib, "NSStringFromPoint")
	tryRegister(&_NSStringFromProtocol, lib, "NSStringFromProtocol")
	tryRegister(&_NSStringFromRange, lib, "NSStringFromRange")
	tryRegister(&_NSStringFromRect, lib, "NSStringFromRect")
	tryRegister(&_NSStringFromSelector, lib, "NSStringFromSelector")
	tryRegister(&_NSStringFromSize, lib, "NSStringFromSize")
	tryRegister(&_NSTemporaryDirectory, lib, "NSTemporaryDirectory")
	tryRegister(&_NSUnionRange, lib, "NSUnionRange")
	tryRegister(&_NSUnionRect, lib, "NSUnionRect")
	tryRegister(&_NSUserName, lib, "NSUserName")
	tryRegister(&_NSZoneCalloc, lib, "NSZoneCalloc")
	tryRegister(&_NSZoneFree, lib, "NSZoneFree")
	tryRegister(&_NSZoneFromPointer, lib, "NSZoneFromPointer")
	tryRegister(&_NSZoneMalloc, lib, "NSZoneMalloc")
	tryRegister(&_NSZoneName, lib, "NSZoneName")
	tryRegister(&_NSZoneRealloc, lib, "NSZoneRealloc")
	tryRegister(&_NXReadNSObjectFromCoder, lib, "NXReadNSObjectFromCoder")
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



// Returns all of the elements in the specified hash table.
//
// Added in macOS 10.0.
// Returns all of the elements in the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllHashTableObjects(_:)
func NSAllHashTableObjects(table unsafe.Pointer) unsafe.Pointer {
	return _NSAllHashTableObjects(table)
}/* debug [functions.gen.go/function]: NSAllHashTableObjects */

// Returns all of the keys in the specified map table.
//
// Added in macOS 10.0.
// Returns all of the keys in the specified map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllMapTableKeys(_:)
func NSAllMapTableKeys(table unsafe.Pointer) unsafe.Pointer {
	return _NSAllMapTableKeys(table)
}/* debug [functions.gen.go/function]: NSAllMapTableKeys */

// Returns all of the values in the specified table.
//
// Added in macOS 10.0.
// Returns all of the values in the specified table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllMapTableValues(_:)
func NSAllMapTableValues(table unsafe.Pointer) unsafe.Pointer {
	return _NSAllMapTableValues(table)
}/* debug [functions.gen.go/function]: NSAllMapTableValues */

// Allocates collectable memory.
//
// Added in macOS 10.0.
// Allocates collectable memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllocateCollectable
func NSAllocateCollectable(size uint, options uint) unsafe.Pointer {
	return _NSAllocateCollectable(size, options)
}/* debug [functions.gen.go/function]: NSAllocateCollectable */

// Allocates a new block of memory.
//
// Added in macOS 10.0.
// Allocates a new block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllocateMemoryPages(_:)
func NSAllocateMemoryPages(bytes uint) unsafe.Pointer {
	return _NSAllocateMemoryPages(bytes)
}/* debug [functions.gen.go/function]: NSAllocateMemoryPages */

// Creates and returns a new instance of a given class.
//
// Added in macOS 10.0.
// Creates and returns a new instance of a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAllocateObject
func NSAllocateObject(aClass objc.Class, extraBytes uint, zone unsafe.Pointer) objc.ID {
	return _NSAllocateObject(aClass, extraBytes, zone)
}/* debug [functions.gen.go/function]: NSAllocateObject */

// Obtains a class by name.
//
// Added in macOS 10.0.
// Obtains a class by name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSClassFromString(_:)
func NSClassFromString(aClassName unsafe.Pointer) objc.Class {
	return _NSClassFromString(aClassName)
}/* debug [functions.gen.go/function]: NSClassFromString */

// Returns a Boolean value that indicates whether the elements of two hash tables are equal.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether the elements of two hash tables are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompareHashTables(_:_:)
func NSCompareHashTables(table1 unsafe.Pointer, table2 unsafe.Pointer) bool {
	return _NSCompareHashTables(table1, table2)
}/* debug [functions.gen.go/function]: NSCompareHashTables */

// Compares the elements of two map tables for equality.
//
// Added in macOS 10.0.
// Compares the elements of two map tables for equality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCompareMapTables(_:_:)
func NSCompareMapTables(table1 unsafe.Pointer, table2 unsafe.Pointer) bool {
	return _NSCompareMapTables(table1, table2)
}/* debug [functions.gen.go/function]: NSCompareMapTables */

// Returns a Boolean value that indicates whether one rectangle completely encloses another.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether one rectangle completely encloses another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSContainsRect(_:_:)
func NSContainsRect(aRect Rect, bRect Rect) bool {
	return _NSContainsRect(aRect, bRect)
}/* debug [functions.gen.go/function]: NSContainsRect */

// Performs a shallow copy of the specified hash table.
//
// Added in macOS 10.0.
// Performs a shallow copy of the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCopyHashTableWithZone(_:_:)
func NSCopyHashTableWithZone(table unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NSCopyHashTableWithZone(table, zone)
}/* debug [functions.gen.go/function]: NSCopyHashTableWithZone */

// Performs a shallow copy of the specified map table.
//
// Added in macOS 10.0.
// Performs a shallow copy of the specified map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCopyMapTableWithZone(_:_:)
func NSCopyMapTableWithZone(table unsafe.Pointer, zone unsafe.Pointer) unsafe.Pointer {
	return _NSCopyMapTableWithZone(table, zone)
}/* debug [functions.gen.go/function]: NSCopyMapTableWithZone */

// Copies a block of memory.
//
// Added in macOS 10.0.
// Copies a block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCopyMemoryPages(_:_:_:)
func NSCopyMemoryPages(source unsafe.Pointer, dest unsafe.Pointer, bytes uint) {
	_NSCopyMemoryPages(source, dest, bytes)
}/* debug [functions.gen.go/function]: NSCopyMemoryPages */

// Creates an exact copy of an object.
//
// Deprecated: This function was deprecated in macOS 10.8.
//
// Added in macOS 10.0.
// Creates an exact copy of an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCopyObject
func NSCopyObject(object objc.ID, extraBytes uint, zone unsafe.Pointer) objc.ID {
	return _NSCopyObject(object, extraBytes, zone)
}/* debug [functions.gen.go/function]: NSCopyObject */

// Returns the number of call frames on the stack.
//
// Added in macOS 10.0.
// Returns the number of call frames on the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountFrames
func NSCountFrames() uint {
	return _NSCountFrames()
}/* debug [functions.gen.go/function]: NSCountFrames */

// Returns the number of elements in a hash table.
//
// Added in macOS 10.0.
// Returns the number of elements in a hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountHashTable(_:)
func NSCountHashTable(table unsafe.Pointer) uint {
	return _NSCountHashTable(table)
}/* debug [functions.gen.go/function]: NSCountHashTable */

// Returns the number of elements in a map table.
//
// Added in macOS 10.0.
// Returns the number of elements in a map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountMapTable(_:)
func NSCountMapTable(table unsafe.Pointer) uint {
	return _NSCountMapTable(table)
}/* debug [functions.gen.go/function]: NSCountMapTable */

// Creates and returns a new hash table.
//
// Added in macOS 10.0.
// Creates and returns a new hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateHashTable(_:_:)
func NSCreateHashTable(callBacks HashTableCallBacks, capacity uint) unsafe.Pointer {
	return _NSCreateHashTable(callBacks, capacity)
}/* debug [functions.gen.go/function]: NSCreateHashTable */

// Creates a new hash table in a given zone.
//
// Added in macOS 10.0.
// Creates a new hash table in a given zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateHashTableWithZone(_:_:_:)
func NSCreateHashTableWithZone(callBacks HashTableCallBacks, capacity uint, zone unsafe.Pointer) unsafe.Pointer {
	return _NSCreateHashTableWithZone(callBacks, capacity, zone)
}/* debug [functions.gen.go/function]: NSCreateHashTableWithZone */

// Creates a new map table in the default zone.
//
// Added in macOS 10.0.
// Creates a new map table in the default zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateMapTable(_:_:_:)
func NSCreateMapTable(keyCallBacks MapTableKeyCallBacks, valueCallBacks MapTableValueCallBacks, capacity uint) unsafe.Pointer {
	return _NSCreateMapTable(keyCallBacks, valueCallBacks, capacity)
}/* debug [functions.gen.go/function]: NSCreateMapTable */

// Creates a new map table in the specified zone.
//
// Added in macOS 10.0.
// Creates a new map table in the specified zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateMapTableWithZone(_:_:_:_:)
func NSCreateMapTableWithZone(keyCallBacks MapTableKeyCallBacks, valueCallBacks MapTableValueCallBacks, capacity uint, zone unsafe.Pointer) unsafe.Pointer {
	return _NSCreateMapTableWithZone(keyCallBacks, valueCallBacks, capacity, zone)
}/* debug [functions.gen.go/function]: NSCreateMapTableWithZone */

// Creates a new zone.
//
// Added in macOS 10.0.
// Creates a new zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCreateZone
func NSCreateZone(startSize uint, granularity uint, canFree bool) unsafe.Pointer {
	return _NSCreateZone(startSize, granularity, canFree)
}/* debug [functions.gen.go/function]: NSCreateZone */

// Deallocates the specified block of memory.
//
// Added in macOS 10.0.
// Deallocates the specified block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDeallocateMemoryPages(_:_:)
func NSDeallocateMemoryPages(ptr unsafe.Pointer, bytes uint) {
	_NSDeallocateMemoryPages(ptr, bytes)
}/* debug [functions.gen.go/function]: NSDeallocateMemoryPages */

// Destroys an existing object.
//
// Added in macOS 10.0.
// Destroys an existing object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDeallocateObject
func NSDeallocateObject(object objc.ID) {
	_NSDeallocateObject(object)
}/* debug [functions.gen.go/function]: NSDeallocateObject */

// Adds two decimal values.
//
// Added in macOS 10.0.
// Adds two decimal values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalAdd(_:_:_:_:)
func NSDecimalAdd(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode RoundingMode) CalculationError {
	return _NSDecimalAdd(result, leftOperand, rightOperand, roundingMode)
}/* debug [functions.gen.go/function]: NSDecimalAdd */

// Compacts the decimal structure for efficiency.
//
// Added in macOS 10.0.
// Compacts the decimal structure for efficiency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCompact(_:)
func NSDecimalCompact(number unsafe.Pointer) {
	_NSDecimalCompact(number)
}/* debug [functions.gen.go/function]: NSDecimalCompact */

// Compares two decimal values.
//
// Added in macOS 10.0.
// Compares two decimal values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCompare(_:_:)
func NSDecimalCompare(leftOperand unsafe.Pointer, rightOperand unsafe.Pointer) ComparisonResult {
	return _NSDecimalCompare(leftOperand, rightOperand)
}/* debug [functions.gen.go/function]: NSDecimalCompare */

// Copies the value of a decimal number.
//
// Added in macOS 10.0.
// Copies the value of a decimal number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalCopy(_:_:)
func NSDecimalCopy(destination unsafe.Pointer, source unsafe.Pointer) {
	_NSDecimalCopy(destination, source)
}/* debug [functions.gen.go/function]: NSDecimalCopy */

// Divides one decimal value by another.
//
// Added in macOS 10.0.
// Divides one decimal value by another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalDivide(_:_:_:_:)
func NSDecimalDivide(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode RoundingMode) CalculationError {
	return _NSDecimalDivide(result, leftOperand, rightOperand, roundingMode)
}/* debug [functions.gen.go/function]: NSDecimalDivide */

// Multiplies two decimal numbers together.
//
// Added in macOS 10.0.
// Multiplies two decimal numbers together.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalMultiply(_:_:_:_:)
func NSDecimalMultiply(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode RoundingMode) CalculationError {
	return _NSDecimalMultiply(result, leftOperand, rightOperand, roundingMode)
}/* debug [functions.gen.go/function]: NSDecimalMultiply */

// Multiplies a decimal by the specified power of 10.
//
// Added in macOS 10.0.
// Multiplies a decimal by the specified power of 10.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalMultiplyByPowerOf10(_:_:_:_:)
func NSDecimalMultiplyByPowerOf10(result unsafe.Pointer, number unsafe.Pointer, power unsafe.Pointer, roundingMode RoundingMode) CalculationError {
	return _NSDecimalMultiplyByPowerOf10(result, number, power, roundingMode)
}/* debug [functions.gen.go/function]: NSDecimalMultiplyByPowerOf10 */

// Normalizes the internal format of two decimal numbers to simplify later operations.
//
// Added in macOS 10.0.
// Normalizes the internal format of two decimal numbers to simplify later operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalNormalize(_:_:_:)
func NSDecimalNormalize(number1 unsafe.Pointer, number2 unsafe.Pointer, roundingMode RoundingMode) CalculationError {
	return _NSDecimalNormalize(number1, number2, roundingMode)
}/* debug [functions.gen.go/function]: NSDecimalNormalize */

// Raises the decimal value to the specified power.
//
// Added in macOS 10.0.
// Raises the decimal value to the specified power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalPower(_:_:_:_:)
func NSDecimalPower(result unsafe.Pointer, number unsafe.Pointer, power uint, roundingMode RoundingMode) CalculationError {
	return _NSDecimalPower(result, number, power, roundingMode)
}/* debug [functions.gen.go/function]: NSDecimalPower */

// Rounds off the decimal value.
//
// Added in macOS 10.0.
// Rounds off the decimal value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalRound(_:_:_:_:)
func NSDecimalRound(result unsafe.Pointer, number unsafe.Pointer, scale int, roundingMode RoundingMode) {
	_NSDecimalRound(result, number, scale, roundingMode)
}/* debug [functions.gen.go/function]: NSDecimalRound */

// Returns a string representation of the decimal value appropriate for the specified locale.
//
// Added in macOS 10.0.
// Returns a string representation of the decimal value appropriate for the specified locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalString(_:_:)
func NSDecimalString(dcm unsafe.Pointer, locale objc.ID) unsafe.Pointer {
	return _NSDecimalString(dcm, locale)
}/* debug [functions.gen.go/function]: NSDecimalString */

// Subtracts one decimal value from another.
//
// Added in macOS 10.0.
// Subtracts one decimal value from another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecimalSubtract(_:_:_:_:)
func NSDecimalSubtract(result unsafe.Pointer, leftOperand unsafe.Pointer, rightOperand unsafe.Pointer, roundingMode RoundingMode) CalculationError {
	return _NSDecimalSubtract(result, leftOperand, rightOperand, roundingMode)
}/* debug [functions.gen.go/function]: NSDecimalSubtract */

// Decrements the specified object’s reference count.
//
// Added in macOS 10.0.
// Decrements the specified object’s reference count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDecrementExtraRefCountWasZero
func NSDecrementExtraRefCountWasZero(object objc.ID) bool {
	return _NSDecrementExtraRefCountWasZero(object)
}/* debug [functions.gen.go/function]: NSDecrementExtraRefCountWasZero */

// Returns the default zone.
//
// Added in macOS 10.0.
// Returns the default zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDefaultMallocZone
func NSDefaultMallocZone() unsafe.Pointer {
	return _NSDefaultMallocZone()
}/* debug [functions.gen.go/function]: NSDefaultMallocZone */

// Divides a rectangle into two new rectangles.
//
// Added in macOS 10.0.
// Divides a rectangle into two new rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDivideRect(_:_:_:_:_:)
func NSDivideRect(inRect Rect, slice unsafe.Pointer, rem unsafe.Pointer, amount float64, edge RectEdge) {
	_NSDivideRect(inRect, slice, rem, amount, edge)
}/* debug [functions.gen.go/function]: NSDivideRect */

// Returns a Boolean value that indicates whether two edge insets structures are equal.
//
// Added in macOS 10.10.
// Returns a Boolean value that indicates whether two edge insets structures are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEdgeInsetsEqual(_:_:)
func NSEdgeInsetsEqual(aInsets EdgeInsets, bInsets EdgeInsets) bool {
	return _NSEdgeInsetsEqual(aInsets, bInsets)
}/* debug [functions.gen.go/function]: NSEdgeInsetsEqual */

// Used when finished with an enumerator.
//
// Added in macOS 10.0.
// Used when finished with an enumerator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEndHashTableEnumeration(_:)
func NSEndHashTableEnumeration(enumerator unsafe.Pointer) {
	_NSEndHashTableEnumeration(enumerator)
}/* debug [functions.gen.go/function]: NSEndHashTableEnumeration */

// Used when finished with an enumerator.
//
// Added in macOS 10.0.
// Used when finished with an enumerator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEndMapTableEnumeration(_:)
func NSEndMapTableEnumeration(enumerator unsafe.Pointer) {
	_NSEndMapTableEnumeration(enumerator)
}/* debug [functions.gen.go/function]: NSEndMapTableEnumeration */

// Creates an enumerator for the specified hash table.
//
// Added in macOS 10.0.
// Creates an enumerator for the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerateHashTable(_:)
func NSEnumerateHashTable(table unsafe.Pointer) HashEnumerator {
	return _NSEnumerateHashTable(table)
}/* debug [functions.gen.go/function]: NSEnumerateHashTable */

// Creates an enumerator for the specified map table.
//
// Added in macOS 10.0.
// Creates an enumerator for the specified map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEnumerateMapTable(_:)
func NSEnumerateMapTable(table unsafe.Pointer) MapEnumerator {
	return _NSEnumerateMapTable(table)
}/* debug [functions.gen.go/function]: NSEnumerateMapTable */

// Returns a Boolean value that indicates whether two points are equal.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether two points are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEqualPoints(_:_:)
func NSEqualPoints(aPoint Point, bPoint Point) bool {
	return _NSEqualPoints(aPoint, bPoint)
}/* debug [functions.gen.go/function]: NSEqualPoints */

// Returns a Boolean value that indicates whether the two rectangles are equal.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether the two rectangles are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEqualRects(_:_:)
func NSEqualRects(aRect Rect, bRect Rect) bool {
	return _NSEqualRects(aRect, bRect)
}/* debug [functions.gen.go/function]: NSEqualRects */

// Returns a Boolean that indicates whether two size values are equal.
//
// Added in macOS 10.0.
// Returns a Boolean that indicates whether two size values are equal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSEqualSizes(_:_:)
func NSEqualSizes(aSize Size, bSize Size) bool {
	return _NSEqualSizes(aSize, bSize)
}/* debug [functions.gen.go/function]: NSEqualSizes */

// Returns the specified object’s reference count.
//
// Added in macOS 10.0.
// Returns the specified object’s reference count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExtraRefCount
func NSExtraRefCount(object objc.ID) uint {
	return _NSExtraRefCount(object)
}/* debug [functions.gen.go/function]: NSExtraRefCount */

// Returns the value of the frame pointer of the specified frame.
//
// Added in macOS 10.0.
// Returns the value of the frame pointer of the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFrameAddress
func NSFrameAddress(frame uint) unsafe.Pointer {
	return _NSFrameAddress(frame)
}/* debug [functions.gen.go/function]: NSFrameAddress */

// Deletes the specified hash table.
//
// Added in macOS 10.0.
// Deletes the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFreeHashTable(_:)
func NSFreeHashTable(table unsafe.Pointer) {
	_NSFreeHashTable(table)
}/* debug [functions.gen.go/function]: NSFreeHashTable */

// Deletes the specified map table.
//
// Added in macOS 10.0.
// Deletes the specified map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFreeMapTable(_:)
func NSFreeMapTable(table unsafe.Pointer) {
	_NSFreeMapTable(table)
}/* debug [functions.gen.go/function]: NSFreeMapTable */

// Returns a string containing the full name of the current user.
//
// Added in macOS 10.0.
// Returns a string containing the full name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSFullUserName()
func NSFullUserName() unsafe.Pointer {
	return _NSFullUserName()
}/* debug [functions.gen.go/function]: NSFullUserName */

// Obtains the actual size and the aligned size of an encoded type.
//
// Added in macOS 10.0.
// Obtains the actual size and the aligned size of an encoded type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSGetSizeAndAlignment(_:_:_:)
func NSGetSizeAndAlignment(typePtr unsafe.Pointer, sizep unsafe.Pointer, alignp unsafe.Pointer) unsafe.Pointer {
	return _NSGetSizeAndAlignment(typePtr, sizep, alignp)
}/* debug [functions.gen.go/function]: NSGetSizeAndAlignment */

// Returns an element of the hash table.
//
// Added in macOS 10.0.
// Returns an element of the hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashGet(_:_:)
func NSHashGet(table unsafe.Pointer, pointer unsafe.Pointer) unsafe.Pointer {
	return _NSHashGet(table, pointer)
}/* debug [functions.gen.go/function]: NSHashGet */

// Adds an element to the specified hash table.
//
// Added in macOS 10.0.
// Adds an element to the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashInsert(_:_:)
func NSHashInsert(table unsafe.Pointer, pointer unsafe.Pointer) {
	_NSHashInsert(table, pointer)
}/* debug [functions.gen.go/function]: NSHashInsert */

// Adds an element to the specified hash table only if the table does not already contain the element.
//
// Added in macOS 10.0.
// Adds an element to the specified hash table only if the table does not already contain the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashInsertIfAbsent(_:_:)
func NSHashInsertIfAbsent(table unsafe.Pointer, pointer unsafe.Pointer) unsafe.Pointer {
	return _NSHashInsertIfAbsent(table, pointer)
}/* debug [functions.gen.go/function]: NSHashInsertIfAbsent */

// Adds an element to the specified hash table.
//
// Added in macOS 10.0.
// Adds an element to the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashInsertKnownAbsent(_:_:)
func NSHashInsertKnownAbsent(table unsafe.Pointer, pointer unsafe.Pointer) {
	_NSHashInsertKnownAbsent(table, pointer)
}/* debug [functions.gen.go/function]: NSHashInsertKnownAbsent */

// Removes an element from the specified hash table.
//
// Added in macOS 10.0.
// Removes an element from the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashRemove(_:_:)
func NSHashRemove(table unsafe.Pointer, pointer unsafe.Pointer) {
	_NSHashRemove(table, pointer)
}/* debug [functions.gen.go/function]: NSHashRemove */

// Returns the path to either the user’s or application’s home directory, depending on the platform.
//
// Added in macOS 10.0.
// Returns the path to either the user’s or application’s home directory, depending on the platform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHomeDirectory()
func NSHomeDirectory() unsafe.Pointer {
	return _NSHomeDirectory()
}/* debug [functions.gen.go/function]: NSHomeDirectory */

// Returns the path to a given user’s home directory.
//
// Added in macOS 10.0.
// Returns the path to a given user’s home directory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHomeDirectoryForUser(_:)
func NSHomeDirectoryForUser(userName unsafe.Pointer) unsafe.Pointer {
	return _NSHomeDirectoryForUser(userName)
}/* debug [functions.gen.go/function]: NSHomeDirectoryForUser */

// Increments the specified object’s reference count.
//
// Added in macOS 10.0.
// Increments the specified object’s reference count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIncrementExtraRefCount
func NSIncrementExtraRefCount(object objc.ID) {
	_NSIncrementExtraRefCount(object)
}/* debug [functions.gen.go/function]: NSIncrementExtraRefCount */

// Insets a rectangle by a specified amount.
//
// Added in macOS 10.0.
// Insets a rectangle by a specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInsetRect(_:_:_:)
func NSInsetRect(aRect Rect, dX float64, dY float64) Rect {
	return _NSInsetRect(aRect, dX, dY)
}/* debug [functions.gen.go/function]: NSInsetRect */

// Adjusts the sides of a rectangle to integer values.
//
// Added in macOS 10.0.
// Adjusts the sides of a rectangle to integer values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntegralRect(_:)
func NSIntegralRect(aRect Rect) Rect {
	return _NSIntegralRect(aRect)
}/* debug [functions.gen.go/function]: NSIntegralRect */

// Adjusts the sides of a rectangle to integral values using the specified options.
//
// Added in macOS 10.7.
// Adjusts the sides of a rectangle to integral values using the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntegralRectWithOptions(_:_:)
func NSIntegralRectWithOptions(aRect Rect, opts AlignmentOptions) Rect {
	return _NSIntegralRectWithOptions(aRect, opts)
}/* debug [functions.gen.go/function]: NSIntegralRectWithOptions */

// Returns the intersection of the specified ranges.
//
// Added in macOS 10.0.
// Returns the intersection of the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntersectionRange(_:_:)
func NSIntersectionRange(range1 Range, range2 Range) Range {
	return _NSIntersectionRange(range1, range2)
}/* debug [functions.gen.go/function]: NSIntersectionRange */

// Calculates the intersection of two rectangles.
//
// Added in macOS 10.0.
// Calculates the intersection of two rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntersectionRect(_:_:)
func NSIntersectionRect(aRect Rect, bRect Rect) Rect {
	return _NSIntersectionRect(aRect, bRect)
}/* debug [functions.gen.go/function]: NSIntersectionRect */

// Returns a Boolean value that indicates whether two rectangles intersect.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether two rectangles intersect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIntersectsRect(_:_:)
func NSIntersectsRect(aRect Rect, bRect Rect) bool {
	return _NSIntersectsRect(aRect, bRect)
}/* debug [functions.gen.go/function]: NSIntersectsRect */

// Returns a Boolean value that indicates whether a given rectangle is empty.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether a given rectangle is empty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIsEmptyRect(_:)
func NSIsEmptyRect(aRect Rect) bool {
	return _NSIsEmptyRect(aRect)
}/* debug [functions.gen.go/function]: NSIsEmptyRect */

// Returns a Boolean indicating whether the specified object has been freed.
//
// Added in macOS 10.0.
// Returns a Boolean indicating whether the specified object has been freed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIsFreedObject
func NSIsFreedObject(anObject objc.ID) bool {
	return _NSIsFreedObject(anObject)
}/* debug [functions.gen.go/function]: NSIsFreedObject */

// Logs an error message to the Apple System Log facility.
//
// Added in macOS 10.0.
// Logs an error message to the Apple System Log facility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLog
func NSLog(format unsafe.Pointer) {
	_NSLog(format)
}/* debug [functions.gen.go/function]: NSLog */

// Returns the binary log of the page size.
//
// Added in macOS 10.0.
// Returns the binary log of the page size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogPageSize()
func NSLogPageSize() uint {
	return _NSLogPageSize()
}/* debug [functions.gen.go/function]: NSLogPageSize */

// Logs an error message to the Apple System Log facility.
//
// Added in macOS 10.0.
// Logs an error message to the Apple System Log facility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLogv(_:_:)
func NSLogv(format unsafe.Pointer, args unsafe.Pointer) {
	_NSLogv(format, args)
}/* debug [functions.gen.go/function]: NSLogv */

// Returns a map table value for the specified key.
//
// Added in macOS 10.0.
// Returns a map table value for the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapGet(_:_:)
func NSMapGet(table unsafe.Pointer, key unsafe.Pointer) unsafe.Pointer {
	return _NSMapGet(table, key)
}/* debug [functions.gen.go/function]: NSMapGet */

// Inserts a key-value pair into the specified table.
//
// Added in macOS 10.0.
// Inserts a key-value pair into the specified table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapInsert(_:_:_:)
func NSMapInsert(table unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_NSMapInsert(table, key, value)
}/* debug [functions.gen.go/function]: NSMapInsert */

// Inserts a key-value pair into the specified table.
//
// Added in macOS 10.0.
// Inserts a key-value pair into the specified table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapInsertIfAbsent(_:_:_:)
func NSMapInsertIfAbsent(table unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) unsafe.Pointer {
	return _NSMapInsertIfAbsent(table, key, value)
}/* debug [functions.gen.go/function]: NSMapInsertIfAbsent */

// Inserts a key-value pair into the specified table if the pair had not been previously added.
//
// Added in macOS 10.0.
// Inserts a key-value pair into the specified table if the pair had not been previously added.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapInsertKnownAbsent(_:_:_:)
func NSMapInsertKnownAbsent(table unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) {
	_NSMapInsertKnownAbsent(table, key, value)
}/* debug [functions.gen.go/function]: NSMapInsertKnownAbsent */

// Indicates whether a given table contains a given key.
//
// Added in macOS 10.0.
// Indicates whether a given table contains a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapMember(_:_:_:_:)
func NSMapMember(table unsafe.Pointer, key unsafe.Pointer, originalKey unsafe.Pointer, value unsafe.Pointer) bool {
	return _NSMapMember(table, key, originalKey, value)
}/* debug [functions.gen.go/function]: NSMapMember */

// Removes a key and corresponding value from the specified table.
//
// Added in macOS 10.0.
// Removes a key and corresponding value from the specified table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapRemove(_:_:)
func NSMapRemove(table unsafe.Pointer, key unsafe.Pointer) {
	_NSMapRemove(table, key)
}/* debug [functions.gen.go/function]: NSMapRemove */

// Returns a Boolean value that indicates whether the point is in the specified rectangle.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether the point is in the specified rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMouseInRect(_:_:_:)
func NSMouseInRect(aPoint Point, aRect Rect, flipped bool) bool {
	return _NSMouseInRect(aPoint, aRect, flipped)
}/* debug [functions.gen.go/function]: NSMouseInRect */

// Returns the next hash-table element in the enumeration.
//
// Added in macOS 10.0.
// Returns the next hash-table element in the enumeration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNextHashEnumeratorItem(_:)
func NSNextHashEnumeratorItem(enumerator unsafe.Pointer) unsafe.Pointer {
	return _NSNextHashEnumeratorItem(enumerator)
}/* debug [functions.gen.go/function]: NSNextHashEnumeratorItem */

// Returns a Boolean value that indicates whether the next map-table pair in the enumeration are set.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether the next map-table pair in the enumeration are set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNextMapEnumeratorPair(_:_:_:)
func NSNextMapEnumeratorPair(enumerator unsafe.Pointer, key unsafe.Pointer, value unsafe.Pointer) bool {
	return _NSNextMapEnumeratorPair(enumerator, key, value)
}/* debug [functions.gen.go/function]: NSNextMapEnumeratorPair */

// Offsets the rectangle by the specified amount.
//
// Added in macOS 10.0.
// Offsets the rectangle by the specified amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOffsetRect(_:_:_:)
func NSOffsetRect(aRect Rect, dX float64, dY float64) Rect {
	return _NSOffsetRect(aRect, dX, dY)
}/* debug [functions.gen.go/function]: NSOffsetRect */

// Returns the root directory of the user’s system.
//
// Added in macOS 10.0.
// Returns the root directory of the user’s system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSOpenStepRootDirectory()
func NSOpenStepRootDirectory() unsafe.Pointer {
	return _NSOpenStepRootDirectory()
}/* debug [functions.gen.go/function]: NSOpenStepRootDirectory */

// Returns the number of bytes in a page.
//
// Added in macOS 10.0.
// Returns the number of bytes in a page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPageSize()
func NSPageSize() uint {
	return _NSPageSize()
}/* debug [functions.gen.go/function]: NSPageSize */

// Returns a point from a text-based representation.
//
// Added in macOS 10.0.
// Returns a point from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointFromString(_:)
func NSPointFromString(aString unsafe.Pointer) Point {
	return _NSPointFromString(aString)
}/* debug [functions.gen.go/function]: NSPointFromString */

// Returns a Boolean value that indicates whether a given point is in a given rectangle.
//
// Added in macOS 10.0.
// Returns a Boolean value that indicates whether a given point is in a given rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointInRect(_:_:)
func NSPointInRect(aPoint Point, aRect Rect) bool {
	return _NSPointInRect(aPoint, aRect)
}/* debug [functions.gen.go/function]: NSPointInRect */

// Returns a the protocol with a given name.
//
// Added in macOS 10.5.
// Returns a the protocol with a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProtocolFromString(_:)
func NSProtocolFromString(namestr unsafe.Pointer) unsafe.Pointer {
	return _NSProtocolFromString(namestr)
}/* debug [functions.gen.go/function]: NSProtocolFromString */

// Returns a range from a textual representation.
//
// Added in macOS 10.0.
// Returns a range from a textual representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeFromString(_:)
func NSRangeFromString(aString unsafe.Pointer) Range {
	return _NSRangeFromString(aString)
}/* debug [functions.gen.go/function]: NSRangeFromString */

// Returns information about the user’s system.

// Returns information about the user’s system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRealMemoryAvailable()
func NSRealMemoryAvailable() uint {
	return _NSRealMemoryAvailable()
}/* debug [functions.gen.go/function]: NSRealMemoryAvailable */

// Reallocates collectable memory.
//
// Added in macOS 10.0.
// Reallocates collectable memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSReallocateCollectable
func NSReallocateCollectable(ptr unsafe.Pointer, size uint, options uint) unsafe.Pointer {
	return _NSReallocateCollectable(ptr, size, options)
}/* debug [functions.gen.go/function]: NSReallocateCollectable */

// Notes an object or zone allocation event and various other statistics, such as the time and current thread.
//
// Added in macOS 10.0.
// Notes an object or zone allocation event and various other statistics, such as the time and current thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecordAllocationEvent
func NSRecordAllocationEvent(eventType int, object objc.ID) {
	_NSRecordAllocationEvent(eventType, object)
}/* debug [functions.gen.go/function]: NSRecordAllocationEvent */

// Returns a rectangle from a text-based representation.
//
// Added in macOS 10.0.
// Returns a rectangle from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectFromString(_:)
func NSRectFromString(aString unsafe.Pointer) Rect {
	return _NSRectFromString(aString)
}/* debug [functions.gen.go/function]: NSRectFromString */

// Frees memory in a zone.
//
// Added in macOS 10.0.
// Frees memory in a zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRecycleZone
func NSRecycleZone(zone unsafe.Pointer) {
	_NSRecycleZone(zone)
}/* debug [functions.gen.go/function]: NSRecycleZone */

// Deletes the elements of the specified hash table.
//
// Added in macOS 10.0.
// Deletes the elements of the specified hash table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSResetHashTable(_:)
func NSResetHashTable(table unsafe.Pointer) {
	_NSResetHashTable(table)
}/* debug [functions.gen.go/function]: NSResetHashTable */

// Deletes the elements of the specified map table.
//
// Added in macOS 10.0.
// Deletes the elements of the specified map table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSResetMapTable(_:)
func NSResetMapTable(table unsafe.Pointer) {
	_NSResetMapTable(table)
}/* debug [functions.gen.go/function]: NSResetMapTable */

// Returns the value of the return address of the specified frame.
//
// Added in macOS 10.0.
// Returns the value of the return address of the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSReturnAddress
func NSReturnAddress(frame uint) unsafe.Pointer {
	return _NSReturnAddress(frame)
}/* debug [functions.gen.go/function]: NSReturnAddress */

// Returns the specified number of bytes rounded down to a multiple of the page size.
//
// Added in macOS 10.0.
// Returns the specified number of bytes rounded down to a multiple of the page size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRoundDownToMultipleOfPageSize(_:)
func NSRoundDownToMultipleOfPageSize(bytes uint) uint {
	return _NSRoundDownToMultipleOfPageSize(bytes)
}/* debug [functions.gen.go/function]: NSRoundDownToMultipleOfPageSize */

// Returns the specified number of bytes rounded up to a multiple of the page size.
//
// Added in macOS 10.0.
// Returns the specified number of bytes rounded up to a multiple of the page size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRoundUpToMultipleOfPageSize(_:)
func NSRoundUpToMultipleOfPageSize(bytes uint) uint {
	return _NSRoundUpToMultipleOfPageSize(bytes)
}/* debug [functions.gen.go/function]: NSRoundUpToMultipleOfPageSize */

// Creates a list of directory search paths.
//
// Added in macOS 10.0.
// Creates a list of directory search paths.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSearchPathForDirectoriesInDomains(_:_:_:)
func NSSearchPathForDirectoriesInDomains(directory SearchPathDirectory, domainMask SearchPathDomainMask, expandTilde bool) []unsafe.Pointer {
	return _NSSearchPathForDirectoriesInDomains(directory, domainMask, expandTilde)
}/* debug [functions.gen.go/function]: NSSearchPathForDirectoriesInDomains */

// Returns the selector with a given name.
//
// Added in macOS 10.0.
// Returns the selector with a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSelectorFromString(_:)
func NSSelectorFromString(aSelectorName unsafe.Pointer) objc.SEL {
	return _NSSelectorFromString(aSelectorName)
}/* debug [functions.gen.go/function]: NSSelectorFromString */

// Sets the name of the specified zone.
//
// Added in macOS 10.0.
// Sets the name of the specified zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSetZoneName
func NSSetZoneName(zone unsafe.Pointer, name unsafe.Pointer) {
	_NSSetZoneName(zone, name)
}/* debug [functions.gen.go/function]: NSSetZoneName */

// Indicates whether an object should be retained.
//
// Added in macOS 10.0.
// Indicates whether an object should be retained.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSShouldRetainWithZone
func NSShouldRetainWithZone(anObject objc.ID, requestedZone unsafe.Pointer) bool {
	return _NSShouldRetainWithZone(anObject, requestedZone)
}/* debug [functions.gen.go/function]: NSShouldRetainWithZone */

// Returns an from a text-based representation.
//
// Added in macOS 10.0.
// Returns an from a text-based representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSizeFromString(_:)
func NSSizeFromString(aString unsafe.Pointer) Size {
	return _NSSizeFromString(aString)
}/* debug [functions.gen.go/function]: NSSizeFromString */

// Returns the name of a class as a string.
//
// Added in macOS 10.0.
// Returns the name of a class as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromClass(_:)
func NSStringFromClass(aClass objc.Class) unsafe.Pointer {
	return _NSStringFromClass(aClass)
}/* debug [functions.gen.go/function]: NSStringFromClass */

// Returns a string describing the hash table’s contents.
//
// Added in macOS 10.0.
// Returns a string describing the hash table’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromHashTable(_:)
func NSStringFromHashTable(table unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromHashTable(table)
}/* debug [functions.gen.go/function]: NSStringFromHashTable */

// Returns a string describing the map table’s contents.
//
// Added in macOS 10.0.
// Returns a string describing the map table’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromMapTable(_:)
func NSStringFromMapTable(table unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromMapTable(table)
}/* debug [functions.gen.go/function]: NSStringFromMapTable */

// Returns a string representation of a point.
//
// Added in macOS 10.0.
// Returns a string representation of a point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromPoint(_:)
func NSStringFromPoint(aPoint Point) unsafe.Pointer {
	return _NSStringFromPoint(aPoint)
}/* debug [functions.gen.go/function]: NSStringFromPoint */

// Returns the name of a protocol as a string.
//
// Added in macOS 10.5.
// Returns the name of a protocol as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromProtocol(_:)
func NSStringFromProtocol(proto unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromProtocol(proto)
}/* debug [functions.gen.go/function]: NSStringFromProtocol */

// Returns a string representation of a range.
//
// Added in macOS 10.0.
// Returns a string representation of a range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromRange(_:)
func NSStringFromRange(range_ Range) unsafe.Pointer {
	return _NSStringFromRange(range_)
}/* debug [functions.gen.go/function]: NSStringFromRange */

// Returns a string representation of a rectangle.
//
// Added in macOS 10.0.
// Returns a string representation of a rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromRect(_:)
func NSStringFromRect(aRect Rect) unsafe.Pointer {
	return _NSStringFromRect(aRect)
}/* debug [functions.gen.go/function]: NSStringFromRect */

// Returns a string representation of a given selector.
//
// Added in macOS 10.0.
// Returns a string representation of a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromSelector(_:)
func NSStringFromSelector(aSelector objc.SEL) unsafe.Pointer {
	return _NSStringFromSelector(aSelector)
}/* debug [functions.gen.go/function]: NSStringFromSelector */

// Returns a string representation of a size.
//
// Added in macOS 10.0.
// Returns a string representation of a size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSStringFromSize(_:)
func NSStringFromSize(aSize Size) unsafe.Pointer {
	return _NSStringFromSize(aSize)
}/* debug [functions.gen.go/function]: NSStringFromSize */

// Returns the path of the temporary directory for the current user.
//
// Added in macOS 10.0.
// Returns the path of the temporary directory for the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTemporaryDirectory()
func NSTemporaryDirectory() unsafe.Pointer {
	return _NSTemporaryDirectory()
}/* debug [functions.gen.go/function]: NSTemporaryDirectory */

// Returns the union of the specified ranges.
//
// Added in macOS 10.0.
// Returns the union of the specified ranges.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnionRange(_:_:)
func NSUnionRange(range1 Range, range2 Range) Range {
	return _NSUnionRange(range1, range2)
}/* debug [functions.gen.go/function]: NSUnionRange */

// Calculates the union of two rectangles.
//
// Added in macOS 10.0.
// Calculates the union of two rectangles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnionRect(_:_:)
func NSUnionRect(aRect Rect, bRect Rect) Rect {
	return _NSUnionRect(aRect, bRect)
}/* debug [functions.gen.go/function]: NSUnionRect */

// Returns the logon name of the current user.
//
// Added in macOS 10.0.
// Returns the logon name of the current user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUserName()
func NSUserName() unsafe.Pointer {
	return _NSUserName()
}/* debug [functions.gen.go/function]: NSUserName */

// Allocates memory in a zone.
//
// Added in macOS 10.0.
// Allocates memory in a zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneCalloc
func NSZoneCalloc(zone unsafe.Pointer, numElems uint, byteSize uint) unsafe.Pointer {
	return _NSZoneCalloc(zone, numElems, byteSize)
}/* debug [functions.gen.go/function]: NSZoneCalloc */

// Deallocates a block of memory in the specified zone.
//
// Added in macOS 10.0.
// Deallocates a block of memory in the specified zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneFree
func NSZoneFree(zone unsafe.Pointer, ptr unsafe.Pointer) {
	_NSZoneFree(zone, ptr)
}/* debug [functions.gen.go/function]: NSZoneFree */

// Gets the zone for a given block of memory.
//
// Added in macOS 10.0.
// Gets the zone for a given block of memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneFromPointer
func NSZoneFromPointer(ptr unsafe.Pointer) unsafe.Pointer {
	return _NSZoneFromPointer(ptr)
}/* debug [functions.gen.go/function]: NSZoneFromPointer */

// Allocates memory in a zone.
//
// Added in macOS 10.0.
// Allocates memory in a zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneMalloc
func NSZoneMalloc(zone unsafe.Pointer, size uint) unsafe.Pointer {
	return _NSZoneMalloc(zone, size)
}/* debug [functions.gen.go/function]: NSZoneMalloc */

// Returns the name of the specified zone.
//
// Added in macOS 10.0.
// Returns the name of the specified zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneName
func NSZoneName(zone unsafe.Pointer) unsafe.Pointer {
	return _NSZoneName(zone)
}/* debug [functions.gen.go/function]: NSZoneName */

// Allocates memory in a zone.
//
// Added in macOS 10.0.
// Allocates memory in a zone.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSZoneRealloc
func NSZoneRealloc(zone unsafe.Pointer, ptr unsafe.Pointer, size uint) unsafe.Pointer {
	return _NSZoneRealloc(zone, ptr, size)
}/* debug [functions.gen.go/function]: NSZoneRealloc */

// Returns the next object from the coder.
//
// Deprecated: This function was deprecated in macOS 10.5.
//
// Added in macOS 10.0.
// Returns the next object from the coder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NXReadNSObjectFromCoder
func NXReadNSObjectFromCoder(decoder unsafe.Pointer) unsafe.Pointer {
	return _NXReadNSObjectFromCoder(decoder)
}/* debug [functions.gen.go/function]: NXReadNSObjectFromCoder */




