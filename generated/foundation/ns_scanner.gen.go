// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Scanner] class.
var (
	ScannerClass     _ScannerClass
	ScannerClassOnce sync.Once
)

func getScannerClass() _ScannerClass {
	ScannerClassOnce.Do(func() {
		ScannerClass = _ScannerClass{objc.GetClass("NSScanner")}
	})
	return ScannerClass
}

type _ScannerClass struct {
	class objc.Class
}

// An interface definition for the [Scanner] class.
type IScanner interface {
	objectivec.IObject
	// properties:
	CaseSensitive() bool /* primitive/slice/pointer. */
	SetCaseSensitive(value bool /* primitive/slice/pointer. */)
	CharactersToBeSkipped() ICharacterSet
	SetCharactersToBeSkipped(value ICharacterSet)
	AtEnd() bool /* primitive/slice/pointer. */
	Locale() objc.ID
	SetLocale(value objc.ID)
	ScanLocation() uint /* primitive/slice/pointer. */
	SetScanLocation(value uint /* primitive/slice/pointer. */)
	String() string /* primitive/slice/pointer. */
	NSNotFound() int /* primitive/slice/pointer. */
	CurrentIndex() unsafe.Pointer
	SetCurrentIndex(value unsafe.Pointer)
	IsAtEnd() bool /* primitive/slice/pointer. */
	SetIsAtEnd(value bool /* primitive/slice/pointer. */)
	// methods:
	ScanHexDouble(result unsafe.Pointer) bool /* primitive/slice/pointer. */
	ScanHexFloat(result unsafe.Pointer) bool /* primitive/slice/pointer. */
	ScanHexLongLong(result unsafe.Pointer) bool /* primitive/slice/pointer. */
	ScanInteger(result Integer /* not a class type */) bool /* primitive/slice/pointer. */
	ScanLongLong(result unsafe.Pointer) bool /* primitive/slice/pointer. */
	ScanUnsignedLongLong(result unsafe.Pointer) bool /* primitive/slice/pointer. */
}

// A string parser that scans for substrings or characters in a character set, and for numeric values from decimal, hexadecimal, and floating-point representations.
//
// A object interprets and converts the characters of a into number and string values. You assign the scanner’s string when you create the scanner, and the scanner progresses through the characters of that string from beginning to end as you request items. Because of the nature of class clusters, a scanner object isn’t an actual instance of the class, but is one of its private subclasses. Although a scanner object’s class is private, its interface is public, as declared by this abstract superclass, . The objects you create using this class are referred to as scanner objects (and when no confusion will result, merely as scanners). To set a object to ignore a set of characters as it scans the string, use the property. Characters in the skip set are skipped over before the target is scanned. The default set of characters to skip is the whitespace and newline character set. To retrieve the unscanned remainder of the string, use .


// A string parser that scans for substrings or characters in a character set, and for numeric values from decimal, hexadecimal, and floating-point representations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner
type Scanner struct {
	objectivec.Object
}

// ScannerFrom constructs a [Scanner] from an unsafe.Pointer.
//
// A string parser that scans for substrings or characters in a character set, and for numeric values from decimal, hexadecimal, and floating-point representations.
func ScannerFrom(ptr unsafe.Pointer) Scanner {
	return Scanner{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScannerClass) Alloc() Scanner {
	rv := objc.Send[Scanner](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScannerClass) New() Scanner {
	rv := objc.Send[Scanner](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Scanner) Init() Scanner {
	rv := objc.Send[Scanner](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Scanner) Autorelease() Scanner {
	rv := objc.Send[Scanner](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScanner creates a new Scanner instance.
func NewScanner() Scanner {
	return getScannerClass().New()
}



// Returns an object initialized to scan a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/init(string:)
func NewScannerWithString(string_ string /* primitive/slice/pointer. */) Scanner {
	instance := getScannerClass().Alloc()
	rv := objc.Send[Scanner](instance.ID, objc.Sel("initWithString:"), objc.String(string_))
	rv.Autorelease()
	return rv
}



// Returns an object that scans a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSScanner/scannerWithString:
func (sc _ScannerClass) ScannerWithString(string_ string /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("scannerWithString:"), objc.String(string_))
	return rv
}


// Returns an object that scans a given string according to the user’s default locale.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/localizedScanner(with:)
func (sc _ScannerClass) LocalizedScannerWithString(string_ string /* primitive/slice/pointer. */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("localizedScannerWithString:"), objc.String(string_))
	return rv
}


// Scans for a double value from a hexadecimal representation, returning a found value by reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanHexDouble(_:)
func (s_ Scanner) ScanHexDouble(result unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanHexDouble:"), result)
	return rv
}


// Scans for a double value from a hexadecimal representation, returning a found value by reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanHexFloat(_:)
func (s_ Scanner) ScanHexFloat(result unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanHexFloat:"), result)
	return rv
}


// Scans for a long long value from a hexadecimal representation, returning a found value by reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanHexInt64(_:)
func (s_ Scanner) ScanHexLongLong(result unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanHexLongLong:"), result)
	return rv
}


// Scans for an NSInteger value from a decimal representation, returning a found value by reference
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanInt(_:)
func (s_ Scanner) ScanInteger(result Integer /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanInteger:"), result)
	return rv
}


// Scans for a long long value from a decimal representation, returning a found value by reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanInt64(_:)
func (s_ Scanner) ScanLongLong(result unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanLongLong:"), result)
	return rv
}


// Scans for an unsigned long long value from a decimal representation, returning a found value by reference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanUnsignedLongLong(_:)
func (s_ Scanner) ScanUnsignedLongLong(result unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanUnsignedLongLong:"), result)
	return rv
}


// Flag that indicates whether the receiver distinguishes case in the characters it scans.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/caseSensitive
func (s_ Scanner) CaseSensitive() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("caseSensitive"))
	return rv
}


// Flag that indicates whether the receiver distinguishes case in the characters it scans.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/caseSensitive
func (s_ Scanner) SetCaseSensitive(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaseSensitive:"), value)
}


// Character set containing the characters the scanner ignores when looking for a scannable element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/charactersToBeSkipped
func (s_ Scanner) CharactersToBeSkipped() ICharacterSet {
	rv := objc.Send[CharacterSet](s_.ID, objc.Sel("charactersToBeSkipped"))
	return rv
}


// Character set containing the characters the scanner ignores when looking for a scannable element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/charactersToBeSkipped
func (s_ Scanner) SetCharactersToBeSkipped(value ICharacterSet) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCharactersToBeSkipped:"), value)
}


// Flag that indicates whether the receiver has exhausted all significant characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/isAtEnd
func (s_ Scanner) AtEnd() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("atEnd"))
	return rv
}


// The locale to use when scanning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/locale
func (s_ Scanner) Locale() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("locale"))
	return rv
}


// The locale to use when scanning.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/locale
func (s_ Scanner) SetLocale(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLocale:"), value)
}


// The character position at which the receiver will begin its next scanning operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanLocation
func (s_ Scanner) ScanLocation() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](s_.ID, objc.Sel("scanLocation"))
	return rv
}


// The character position at which the receiver will begin its next scanning operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanLocation
func (s_ Scanner) SetScanLocation(value uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScanLocation:"), value)
}


// The string the scanner will scan.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/string
func (s_ Scanner) String() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](s_.ID, objc.Sel("string"))
	return rv
}


// A value indicating that a requested item couldn’t be found or doesn’t exist.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (s_ Scanner) NSNotFound() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](s_.ID, objc.Sel("NSNotFound"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/currentindex
func (s_ Scanner) CurrentIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/currentindex
func (s_ Scanner) SetCurrentIndex(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentIndex:"), value)
}


// Flag that indicates whether the receiver has exhausted all significant characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/isatend
func (s_ Scanner) IsAtEnd() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](s_.ID, objc.Sel("isAtEnd"))
	return rv
}


// Flag that indicates whether the receiver has exhausted all significant characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/isatend
func (s_ Scanner) SetIsAtEnd(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsAtEnd:"), value)
}


