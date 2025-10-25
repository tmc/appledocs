// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMutableCharacterSet */


/* debug [class_header]: Header for NSMutableCharacterSet */
// The class instance for the [MutableCharacterSet] class.
var (
	MutableCharacterSetClass     _MutableCharacterSetClass
	MutableCharacterSetClassOnce sync.Once
)

func getMutableCharacterSetClass() _MutableCharacterSetClass {
	MutableCharacterSetClassOnce.Do(func() {
		MutableCharacterSetClass = _MutableCharacterSetClass{objc.GetClass("NSMutableCharacterSet")}
	})
	return MutableCharacterSetClass
}

type _MutableCharacterSetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableCharacterSet */
// An interface definition for the [MutableCharacterSet] class.
type IMutableCharacterSet interface {
	ICharacterSet
	
/* debug [class_interface_properties]: Properties for MutableCharacterSet */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableCharacterSet */
	// methods:
	AddCharactersInRange(aRange objc.IObject /* cross-framework: Range */)
	AddCharactersInString(aString IString)
	FormIntersectionWithCharacterSet(otherSet ICharacterSet)
	FormUnionWithCharacterSet(otherSet ICharacterSet)
	Invert()
	RemoveCharactersInRange(aRange objc.IObject /* cross-framework: Range */)
	RemoveCharactersInString(aString IString)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableCharacterSet */
// Alloc allocates a new instance without initialization.
func (mc _MutableCharacterSetClass) Alloc() MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableCharacterSetClass) New() MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableCharacterSet) Init() MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableCharacterSet) Autorelease() MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableCharacterSet creates a new MutableCharacterSet instance.
func NewMutableCharacterSet() MutableCharacterSet {
	return getMutableCharacterSetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableCharacterSet */
// An object representing a mutable set of Unicode character values for use in search operations.
//
// In Swift, this object bridges to ; use when you need reference semantics or other Foundation-specific behavior. The class declares the programmatic interface to objects that manage a modifiable set of Unicode characters. You can add or remove characters from a mutable character set as numeric values in structures or as character values in strings, combine character sets by union or intersection, and invert a character set. Mutable character sets are less efficient to use than immutable character sets. If you don’t need to change a character set after creating it, create an immutable copy with and use that. defines no primitive methods. Subclasses must implement all methods declared by this class in addition to the primitives of . They must also implement . is “toll-free bridged” with its Core Foundation counterpart, . See for more information.


// An object representing a mutable set of Unicode character values for use in search operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet
type MutableCharacterSet struct {
	CharacterSet
}

// MutableCharacterSetFrom constructs a [MutableCharacterSet] from an unsafe.Pointer.
//
// An object representing a mutable set of Unicode character values for use in search operations.
func MutableCharacterSetFrom(ptr unsafe.Pointer) MutableCharacterSet {
	return MutableCharacterSet{
		CharacterSet: CharacterSetFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableCharacterSet */

// Returns a character set containing characters determined by a given bitmap representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(bitmapRepresentation:)
func NewMutableCharacterSetWithBitmapRepresentation(data IData) MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(getMutableCharacterSetClass().class), objc.Sel("characterSetWithBitmapRepresentation:"), data)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableCharacterSetWithBitmapRepresentation */


// Returns a character set containing the characters in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(charactersIn:)
func NewMutableCharacterSetWithCharactersInString(aString IString) MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(getMutableCharacterSetClass().class), objc.Sel("characterSetWithCharactersInString:"), aString)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableCharacterSetWithCharactersInString */


// Returns a character set read from the bitmap representation stored in the file a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(contentsOfFile:)
func NewMutableCharacterSetWithContentsOfFile(fName IString) MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(getMutableCharacterSetClass().class), objc.Sel("characterSetWithContentsOfFile:"), fName)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableCharacterSetWithContentsOfFile */


// Returns a character set containing characters with Unicode values in a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(range:)
func NewMutableCharacterSetWithRange(aRange objc.IObject /* cross-framework: Range */) MutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(getMutableCharacterSetClass().class), objc.Sel("characterSetWithRange:"), aRange)
	return rv
}/* debug [class_init_methods/constructor]: NewMutableCharacterSetWithRange */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableCharacterSet */

// Returns a character set containing the characters in Unicode General Categories L*, M*, and N*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/alphanumeric()
func (mc _MutableCharacterSetClass) AlphanumericCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("alphanumericCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AlphanumericCharacterSet) */


// Returns a character set containing the characters in Unicode General Category Lt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/capitalizedLetter()
func (mc _MutableCharacterSetClass) CapitalizedLetterCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("capitalizedLetterCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CapitalizedLetterCharacterSet) */


// Returns a character set containing the characters in Unicode General Category Cc and Cf.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/control()
func (mc _MutableCharacterSetClass) ControlCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("controlCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ControlCharacterSet) */


// Returns a character set containing the characters in the category of decimal numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/decimalDigit()
func (mc _MutableCharacterSetClass) DecimalDigitCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("decimalDigitCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecimalDigitCharacterSet) */


// Returns a character set containing individual Unicode characters that can also be represented as composed character sequences (such as for letters with accents), by the definition of “standard decomposition” in version 3.2 of the Unicode character encoding standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/decomposable()
func (mc _MutableCharacterSetClass) DecomposableCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("decomposableCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecomposableCharacterSet) */


// Returns a character set containing values in the category of Non-Characters or that have not yet been defined in version 3.2 of the Unicode standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/illegal()
func (mc _MutableCharacterSetClass) IllegalCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("illegalCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=IllegalCharacterSet) */


// Returns a character set containing characters determined by a given bitmap representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(bitmapRepresentation:)
func (mc _MutableCharacterSetClass) CharacterSetWithBitmapRepresentation(data IData) IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("characterSetWithBitmapRepresentation:"), data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CharacterSetWithBitmapRepresentation) */


// Returns a character set containing the characters in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(charactersIn:)
func (mc _MutableCharacterSetClass) CharacterSetWithCharactersInString(aString IString) IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("characterSetWithCharactersInString:"), aString)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CharacterSetWithCharactersInString) */


// Returns a character set read from the bitmap representation stored in the file a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(contentsOfFile:)
func (mc _MutableCharacterSetClass) CharacterSetWithContentsOfFile(fName IString) IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("characterSetWithContentsOfFile:"), fName)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CharacterSetWithContentsOfFile) */


// Returns a character set containing characters with Unicode values in a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/init(range:)
func (mc _MutableCharacterSetClass) CharacterSetWithRange(aRange objc.IObject /* cross-framework: Range */) IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("characterSetWithRange:"), aRange)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CharacterSetWithRange) */


// Returns a character set containing the characters in Unicode General Category L* & M*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/letter()
func (mc _MutableCharacterSetClass) LetterCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("letterCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LetterCharacterSet) */


// Returns a character set containing the characters in Unicode General Category Ll.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/lowercaseLetter()
func (mc _MutableCharacterSetClass) LowercaseLetterCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("lowercaseLetterCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LowercaseLetterCharacterSet) */


// Returns a character set containing the newline characters ( ~ , , , and ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/newline()
func (mc _MutableCharacterSetClass) NewlineCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("newlineCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NewlineCharacterSet) */


// Returns a character set containing the characters in Unicode General Category M*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/nonBase()
func (mc _MutableCharacterSetClass) NonBaseCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("nonBaseCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NonBaseCharacterSet) */


// Returns a character set containing the characters in Unicode General Category P*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/punctuation()
func (mc _MutableCharacterSetClass) PunctuationCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("punctuationCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PunctuationCharacterSet) */


// Returns a character set containing the characters in Unicode General Category S*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/symbol()
func (mc _MutableCharacterSetClass) SymbolCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("symbolCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SymbolCharacterSet) */


// Returns a character set containing the characters in Unicode General Category Lu and Lt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/uppercaseLetter()
func (mc _MutableCharacterSetClass) UppercaseLetterCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("uppercaseLetterCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UppercaseLetterCharacterSet) */


// Returns a character set containing the characters in Unicode General Category Zs and ( ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/whitespace()
func (mc _MutableCharacterSetClass) WhitespaceCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("whitespaceCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WhitespaceCharacterSet) */


// Returns a character set containing characters in Unicode General Category Z*, ~ , and .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/whitespaceAndNewline()
func (mc _MutableCharacterSetClass) WhitespaceAndNewlineCharacterSet() IMutableCharacterSet {
	rv := objc.Send[MutableCharacterSet](objc.ID(mc.class), objc.Sel("whitespaceAndNewlineCharacterSet"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WhitespaceAndNewlineCharacterSet) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableCharacterSet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableCharacterSet */

// Adds to the receiver the characters whose Unicode values are in a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/addCharacters(in:)-4ppyw
func (m_ MutableCharacterSet) AddCharactersInRange(aRange objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addCharactersInRange:"), aRange)
}/* debug [instance_methods/method]: AddCharactersInRange */


// Adds to the receiver the characters in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/addCharacters(in:)-7q02
func (m_ MutableCharacterSet) AddCharactersInString(aString IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addCharactersInString:"), aString)
}/* debug [instance_methods/method]: AddCharactersInString */


// Modifies the receiver so it contains only characters that exist in both the receiver and another set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/formIntersection(with:)
func (m_ MutableCharacterSet) FormIntersectionWithCharacterSet(otherSet ICharacterSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("formIntersectionWithCharacterSet:"), otherSet)
}/* debug [instance_methods/method]: FormIntersectionWithCharacterSet */


// Modifies the receiver so it contains all characters that exist in either the receiver or another set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/formUnion(with:)
func (m_ MutableCharacterSet) FormUnionWithCharacterSet(otherSet ICharacterSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("formUnionWithCharacterSet:"), otherSet)
}/* debug [instance_methods/method]: FormUnionWithCharacterSet */


// Replaces all the characters in the receiver with all the characters it didn’t previously contain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/invert()
func (m_ MutableCharacterSet) Invert() {
	objc.Send[objc.ID](m_.ID, objc.Sel("invert"))
}/* debug [instance_methods/method]: Invert */


// Removes from the receiver the characters whose Unicode values are in a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/removeCharacters(in:)-70nqp
func (m_ MutableCharacterSet) RemoveCharactersInRange(aRange objc.IObject /* cross-framework: Range */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeCharactersInRange:"), aRange)
}/* debug [instance_methods/method]: RemoveCharactersInRange */


// Removes from the receiver the characters in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableCharacterSet/removeCharacters(in:)-762gt
func (m_ MutableCharacterSet) RemoveCharactersInString(aString IString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeCharactersInString:"), aString)
}/* debug [instance_methods/method]: RemoveCharactersInString */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableCharacterSet */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMutableCharacterSet */


