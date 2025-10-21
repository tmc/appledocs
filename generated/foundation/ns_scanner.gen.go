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
	ScanDouble(result unsafe.Pointer) bool
	ScanFloat(result unsafe.Pointer) bool
	ScanInt(result unsafe.Pointer) bool
}

// A string parser that scans for substrings or characters in a character set, and for numeric values from decimal, hexadecimal, and floating-point representations.
//
// A object interprets and converts the characters of a into number and string values. You assign the scanner’s string when you create the scanner, and the scanner progresses through the characters of that string from beginning to end as you request items. Because of the nature of class clusters, a scanner object isn’t an actual instance of the class, but is one of its private subclasses. Although a scanner object’s class is private, its interface is public, as declared by this abstract superclass, . The objects you create using this class are referred to as scanner objects (and when no confusion will result, merely as scanners). To set a object to ignore a set of characters as it scans the string, use the property. Characters in the skip set are skipped over before the target is scanned. The default set of characters to skip is the whitespace and newline character set. To retrieve the unscanned remainder of the string, use .
//
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


// Returns an object that scans a given string according to the user’s default locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/localizedScanner(with:)
func (sc _ScannerClass) LocalizedScannerWithString(string_ string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("localizedScannerWithString:"), objc.String(string_))
	return rv
}

// Scans for a double value, returning a found value by reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanDouble(_:)
func (s_ Scanner) ScanDouble(result unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanDouble:"), result)
	return rv
}

// Scans for a float value, returning a found value by reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanFloat(_:)
func (s_ Scanner) ScanFloat(result unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanFloat:"), result)
	return rv
}

// Scans for an int value from a decimal representation, returning a found value by reference.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Scanner/scanInt32(_:)
func (s_ Scanner) ScanInt(result unsafe.Pointer) bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scanInt:"), result)
	return rv
}

// A value indicating that a requested item couldn’t be found or doesn’t exist.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsnotfound-4qp9h
func (s_ Scanner) NSNotFound() int {
	rv := objc.Send[int](s_.ID, objc.Sel("NSNotFound"))
	return rv
}

// Flag that indicates whether the receiver distinguishes case in the characters it scans.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/casesensitive
func (s_ Scanner) CaseSensitive() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("caseSensitive"))
	return rv
}


// SetCaseSensitive sets the value of the caseSensitive property.
// Flag that indicates whether the receiver distinguishes case in the characters it scans.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/casesensitive
func (s_ Scanner) SetCaseSensitive(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCaseSensitive:"), value)
}

// Character set containing the characters the scanner ignores when looking for a scannable element.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/characterstobeskipped
func (s_ Scanner) CharactersToBeSkipped() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("charactersToBeSkipped"))
	return rv
}


// SetCharactersToBeSkipped sets the value of the charactersToBeSkipped property.
// Character set containing the characters the scanner ignores when looking for a scannable element.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/characterstobeskipped
func (s_ Scanner) SetCharactersToBeSkipped(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCharactersToBeSkipped:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/currentindex
func (s_ Scanner) CurrentIndex() string {
	rv := objc.Send[string](s_.ID, objc.Sel("currentIndex"))
	return rv
}


// SetCurrentIndex sets the value of the currentIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/currentindex
func (s_ Scanner) SetCurrentIndex(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCurrentIndex:"), objc.String(value))
}

// Flag that indicates whether the receiver has exhausted all significant characters.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/isatend
func (s_ Scanner) IsAtEnd() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isAtEnd"))
	return rv
}


// SetIsAtEnd sets the value of the isAtEnd property.
// Flag that indicates whether the receiver has exhausted all significant characters.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/isatend
func (s_ Scanner) SetIsAtEnd(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsAtEnd:"), value)
}

// The locale to use when scanning.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/locale
func (s_ Scanner) Locale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("locale"))
	return rv
}


// SetLocale sets the value of the locale property.
// The locale to use when scanning.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/locale
func (s_ Scanner) SetLocale(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLocale:"), value)
}

// The character position at which the receiver will begin its next scanning operation.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/scanlocation
func (s_ Scanner) ScanLocation() int {
	rv := objc.Send[int](s_.ID, objc.Sel("scanLocation"))
	return rv
}


// SetScanLocation sets the value of the scanLocation property.
// The character position at which the receiver will begin its next scanning operation.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/scanlocation
func (s_ Scanner) SetScanLocation(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScanLocation:"), value)
}

// The string the scanner will scan.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/string
func (s_ Scanner) String_() string {
	rv := objc.Send[string](s_.ID, objc.Sel("string"))
	return rv
}


// SetString_ sets the value of the string property.
// The string the scanner will scan.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/scanner/string
func (s_ Scanner) SetString_(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setString_:"), objc.String(value))
}



