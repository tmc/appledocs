// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
	coregraphics "github.com/tmc/appledocs/generated/coregraphics"
)


// Foundation Functions (82 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_NSCountFrames func() unsafe.Pointer
	_NSFrameAddress func(unsafe.Pointer) unsafe.Pointer
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
	_CFURLClearResourcePropertyCache func(unsafe.Pointer)
	_CFURLClearResourcePropertyCacheForKey func(unsafe.Pointer, unsafe.Pointer)
	_CFURLCopyNetLocation func(unsafe.Pointer) unsafe.Pointer
	_CFURLStartAccessingSecurityScopedResource func(unsafe.Pointer) unsafe.Pointer
	_CFURLStopAccessingSecurityScopedResource func(unsafe.Pointer)
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_NSCountFrames, lib, "NSCountFrames")
	tryRegister(&_NSFrameAddress, lib, "NSFrameAddress")
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
	tryRegister(&_CFURLClearResourcePropertyCache, lib, "CFURLClearResourcePropertyCache")
	tryRegister(&_CFURLClearResourcePropertyCacheForKey, lib, "CFURLClearResourcePropertyCacheForKey")
	tryRegister(&_CFURLCopyNetLocation, lib, "CFURLCopyNetLocation")
	tryRegister(&_CFURLStartAccessingSecurityScopedResource, lib, "CFURLStartAccessingSecurityScopedResource")
	tryRegister(&_CFURLStopAccessingSecurityScopedResource, lib, "CFURLStopAccessingSecurityScopedResource")
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

// Returns the net location portion of a given URL.

// Returns the net location portion of a given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLCopyNetLocation(_:)
func CFURLCopyNetLocation(anURL unsafe.Pointer) unsafe.Pointer {
	return _CFURLCopyNetLocation(anURL)
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



