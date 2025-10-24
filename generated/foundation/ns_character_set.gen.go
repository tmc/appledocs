// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCharacterSet */


/* debug [class_header]: Header for NSCharacterSet */
// The class instance for the [CharacterSet] class.
var (
	CharacterSetClass     _CharacterSetClass
	CharacterSetClassOnce sync.Once
)

func getCharacterSetClass() _CharacterSetClass {
	CharacterSetClassOnce.Do(func() {
		CharacterSetClass = _CharacterSetClass{objc.GetClass("NSCharacterSet")}
	})
	return CharacterSetClass
}

type _CharacterSetClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CharacterSet */
// An interface definition for the [CharacterSet] class.
type ICharacterSet interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CharacterSet */
	// properties:
	BitmapRepresentation() IData
	InvertedSet() ICharacterSet
	Inverted() ICharacterSet
	SetInverted(value ICharacterSet)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CharacterSet */
	// methods:
	CharacterIsMember(aCharacter unichar /* typedef */) bool
	HasMemberInPlane(thePlane uint8 /* not a class type */) bool
	IsSupersetOfSet(theOtherSet ICharacterSet) bool
	LongCharacterIsMember(theLongChar objectivec.IObject) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CharacterSet */
// Alloc allocates a new instance without initialization.
func (cc _CharacterSetClass) Alloc() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CharacterSetClass) New() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CharacterSet) Init() CharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CharacterSet) Autorelease() CharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCharacterSet creates a new CharacterSet instance.
func NewCharacterSet() CharacterSet {
	return getCharacterSetClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CharacterSet */
// An object representing a fixed set of Unicode character values for use in search operations.
//
// In Swift, this bridges to a ; use when you need reference semantics or other Foundation-specific behavior. An object represents a set of Unicode-compliant characters. and objects use objects to group characters together for searching operations, so that they can find any of a particular set of characters during a search. The cluster’s two public classes, and , declare the programmatic interface for static and dynamic character sets, respectively. The objects you create using these classes are referred to as character set objects (and when no confusion will result, merely as character sets). Because of the nature of class clusters, character set objects aren’t actual instances of the or classes but of one of their private subclasses. Although a character set object’s class is private, its interface is public, as declared by these abstract superclasses, and . The character set classes adopt the and protocols, making it convenient to convert a character set of one type to the other. The class declares the programmatic interface for an object that manages a set of Unicode characters (see the class cluster specification for information on Unicode). ’s principal primitive method, , provides the basis for all other instance methods in its interface. A subclass of needs only to implement this method, plus , for proper behavior. For optimal performance, a subclass should also override , which otherwise works by invoking for every possible Unicode value. is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.


// An object representing a fixed set of Unicode character values for use in search operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet
type CharacterSet struct {
	objectivec.Object
}

// CharacterSetFrom constructs a [CharacterSet] from an unsafe.Pointer.
//
// An object representing a fixed set of Unicode character values for use in search operations.
func CharacterSetFrom(ptr unsafe.Pointer) CharacterSet {
	return CharacterSet{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CharacterSet */

// Returns a character set containing characters determined by a given bitmap representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(bitmapRepresentation:)
func NewCharacterSetWithBitmapRepresentation(data IData) CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(getCharacterSetClass().class), objc.Sel("characterSetWithBitmapRepresentation:"), data)
	return rv
}/* debug [class_init_methods/constructor]: NewCharacterSetWithBitmapRepresentation */


// Returns a character set containing the characters in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(charactersIn:)
func NewCharacterSetWithCharactersInString(aString IString) CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(getCharacterSetClass().class), objc.Sel("characterSetWithCharactersInString:"), aString)
	return rv
}/* debug [class_init_methods/constructor]: NewCharacterSetWithCharactersInString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(coder:)
func NewCharacterSetWithCoder(coder ICoder) CharacterSet {
	instance := getCharacterSetClass().Alloc()
	rv := objc.Send[CharacterSet](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCharacterSetWithCoder */


// Returns a character set read from the bitmap representation stored in the file a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(contentsOfFile:)
func NewCharacterSetWithContentsOfFile(fName IString) CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(getCharacterSetClass().class), objc.Sel("characterSetWithContentsOfFile:"), fName)
	return rv
}/* debug [class_init_methods/constructor]: NewCharacterSetWithContentsOfFile */


// Returns a character set containing characters with Unicode values in a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(range:)
func NewCharacterSetWithRange(aRange objc.IObject /* cross-framework: Range */) CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(getCharacterSetClass().class), objc.Sel("characterSetWithRange:"), aRange)
	return rv
}/* debug [class_init_methods/constructor]: NewCharacterSetWithRange */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CharacterSet */

// Returns a character set containing characters determined by a given bitmap representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(bitmapRepresentation:)
func (cc _CharacterSetClass) CharacterSetWithBitmapRepresentation(data IData) ICharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("characterSetWithBitmapRepresentation:"), data)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CharacterSetWithBitmapRepresentation) */


// Returns a character set containing the characters in a given string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(charactersIn:)
func (cc _CharacterSetClass) CharacterSetWithCharactersInString(aString IString) ICharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("characterSetWithCharactersInString:"), aString)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CharacterSetWithCharactersInString) */


// Returns a character set read from the bitmap representation stored in the file a given path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(contentsOfFile:)
func (cc _CharacterSetClass) CharacterSetWithContentsOfFile(fName IString) ICharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("characterSetWithContentsOfFile:"), fName)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CharacterSetWithContentsOfFile) */


// Returns a character set containing characters with Unicode values in a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/init(range:)
func (cc _CharacterSetClass) CharacterSetWithRange(aRange objc.IObject /* cross-framework: Range */) ICharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("characterSetWithRange:"), aRange)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CharacterSetWithRange) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CharacterSet */

// A character set containing the characters in Unicode General Categories L*, M*, and N*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/alphanumerics
func (cc _CharacterSetClass) AlphanumericCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("alphanumericCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: alphanumericCharacterSet */

// A character set containing the characters in Unicode General Category Lt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/capitalizedLetters
func (cc _CharacterSetClass) CapitalizedLetterCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("capitalizedLetterCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: capitalizedLetterCharacterSet */

// A character set containing the characters in Unicode General Category Cc and Cf.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/controlCharacters
func (cc _CharacterSetClass) ControlCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("controlCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: controlCharacterSet */

// A character set containing the characters in the category of Decimal Numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/decimalDigits
func (cc _CharacterSetClass) DecimalDigitCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("decimalDigitCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: decimalDigitCharacterSet */

// A character set containing individual Unicode characters that can also be represented as composed character sequences (such as for letters with accents), by the definition of “standard decomposition” in version 3.2 of the Unicode character encoding standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/decomposables
func (cc _CharacterSetClass) DecomposableCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("decomposableCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: decomposableCharacterSet */

// A character set containing values in the category of Non-Characters or that have not yet been defined in version 3.2 of the Unicode standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/illegalCharacters
func (cc _CharacterSetClass) IllegalCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("illegalCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: illegalCharacterSet */

// A character set containing the characters in Unicode General Category L* & M*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/letters
func (cc _CharacterSetClass) LetterCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("letterCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: letterCharacterSet */

// A character set containing the characters in Unicode General Category Ll.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/lowercaseLetters
func (cc _CharacterSetClass) LowercaseLetterCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("lowercaseLetterCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: lowercaseLetterCharacterSet */

// A character set containing the newline characters ( ~ , , , and ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/newlines
func (cc _CharacterSetClass) NewlineCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("newlineCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: newlineCharacterSet */

// A character set containing the characters in Unicode General Category M*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/nonBaseCharacters
func (cc _CharacterSetClass) NonBaseCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("nonBaseCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: nonBaseCharacterSet */

// A character set containing the characters in Unicode General Category P*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/punctuationCharacters
func (cc _CharacterSetClass) PunctuationCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("punctuationCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: punctuationCharacterSet */

// A character set containing the characters in Unicode General Category S*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/symbols
func (cc _CharacterSetClass) SymbolCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("symbolCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: symbolCharacterSet */

// A character set containing the characters in Unicode General Category Lu and Lt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/uppercaseLetters
func (cc _CharacterSetClass) UppercaseLetterCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("uppercaseLetterCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: uppercaseLetterCharacterSet */

// Returns the character set for characters allowed in a fragment URL component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlFragmentAllowed
func (cc _CharacterSetClass) URLFragmentAllowedCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("URLFragmentAllowedCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: URLFragmentAllowedCharacterSet */

// Returns the character set for characters allowed in a host URL subcomponent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlHostAllowed
func (cc _CharacterSetClass) URLHostAllowedCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("URLHostAllowedCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: URLHostAllowedCharacterSet */

// Returns the character set for characters allowed in a password URL subcomponent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlPasswordAllowed
func (cc _CharacterSetClass) URLPasswordAllowedCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("URLPasswordAllowedCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: URLPasswordAllowedCharacterSet */

// Returns the character set for characters allowed in a path URL component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlPathAllowed
func (cc _CharacterSetClass) URLPathAllowedCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("URLPathAllowedCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: URLPathAllowedCharacterSet */

// Returns the character set for characters allowed in a query URL component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlQueryAllowed
func (cc _CharacterSetClass) URLQueryAllowedCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("URLQueryAllowedCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: URLQueryAllowedCharacterSet */

// Returns the character set for characters allowed in a user URL subcomponent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlUserAllowed
func (cc _CharacterSetClass) URLUserAllowedCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("URLUserAllowedCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: URLUserAllowedCharacterSet */

// A character set containing the characters in Unicode General Category Zs and ( ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/whitespaces
func (cc _CharacterSetClass) WhitespaceCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("whitespaceCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: whitespaceCharacterSet */

// A character set containing characters in Unicode General Category Z*, ~ , and .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/whitespacesAndNewlines
func (cc _CharacterSetClass) WhitespaceAndNewlineCharacterSet() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("whitespaceAndNewlineCharacterSet"))
	return rv
}/* debug [class_properties_class/property]: whitespaceAndNewlineCharacterSet */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CharacterSet */

// Returns a Boolean value that indicates whether a given character is in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/characterIsMember(_:)
func (c_ CharacterSet) CharacterIsMember(aCharacter unichar /* typedef */) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("characterIsMember:"), aCharacter)
	return rv
}/* debug [instance_methods/method]: CharacterIsMember */


// Returns a Boolean value that indicates whether the receiver has at least one member in a given character plane.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/hasMemberInPlane(_:)
func (c_ CharacterSet) HasMemberInPlane(thePlane uint8 /* not a class type */) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasMemberInPlane:"), thePlane)
	return rv
}/* debug [instance_methods/method]: HasMemberInPlane */


// Returns a Boolean value that indicates whether the receiver is a superset of another given character set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/isSuperset(of:)
func (c_ CharacterSet) IsSupersetOfSet(theOtherSet ICharacterSet) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSupersetOfSet:"), theOtherSet)
	return rv
}/* debug [instance_methods/method]: IsSupersetOfSet */


// Returns a Boolean value that indicates whether a given long character is a member of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/longCharacterIsMember(_:)
func (c_ CharacterSet) LongCharacterIsMember(theLongChar objectivec.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("longCharacterIsMember:"), theLongChar)
	return rv
}/* debug [instance_methods/method]: LongCharacterIsMember */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CharacterSet */

// A character set containing the characters in Unicode General Categories L*, M*, and N*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/alphanumerics
func (c_ CharacterSet) AlphanumericCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("alphanumericCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: alphanumericCharacterSet */


// An object encoding the receiver in binary format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/bitmapRepresentation
func (c_ CharacterSet) BitmapRepresentation() IData {
	rv := objc.Send[Data](c_.ID, objc.Sel("bitmapRepresentation"))
	return rv
}/* debug [instance_properties/getter]: bitmapRepresentation */


// A character set containing the characters in Unicode General Category Lt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/capitalizedLetters
func (c_ CharacterSet) CapitalizedLetterCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("capitalizedLetterCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: capitalizedLetterCharacterSet */


// A character set containing the characters in Unicode General Category Cc and Cf.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/controlCharacters
func (c_ CharacterSet) ControlCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("controlCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: controlCharacterSet */


// A character set containing the characters in the category of Decimal Numbers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/decimalDigits
func (c_ CharacterSet) DecimalDigitCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("decimalDigitCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: decimalDigitCharacterSet */


// A character set containing individual Unicode characters that can also be represented as composed character sequences (such as for letters with accents), by the definition of “standard decomposition” in version 3.2 of the Unicode character encoding standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/decomposables
func (c_ CharacterSet) DecomposableCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("decomposableCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: decomposableCharacterSet */


// A character set containing values in the category of Non-Characters or that have not yet been defined in version 3.2 of the Unicode standard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/illegalCharacters
func (c_ CharacterSet) IllegalCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("illegalCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: illegalCharacterSet */


// A character set containing only characters that don’t exist in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/inverted
func (c_ CharacterSet) InvertedSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("invertedSet"))
	return rv
}/* debug [instance_properties/getter]: invertedSet */


// A character set containing the characters in Unicode General Category L* & M*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/letters
func (c_ CharacterSet) LetterCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("letterCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: letterCharacterSet */


// A character set containing the characters in Unicode General Category Ll.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/lowercaseLetters
func (c_ CharacterSet) LowercaseLetterCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("lowercaseLetterCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: lowercaseLetterCharacterSet */


// A character set containing the newline characters ( ~ , , , and ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/newlines
func (c_ CharacterSet) NewlineCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("newlineCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: newlineCharacterSet */


// A character set containing the characters in Unicode General Category M*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/nonBaseCharacters
func (c_ CharacterSet) NonBaseCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("nonBaseCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: nonBaseCharacterSet */


// A character set containing the characters in Unicode General Category P*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/punctuationCharacters
func (c_ CharacterSet) PunctuationCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("punctuationCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: punctuationCharacterSet */


// A character set containing the characters in Unicode General Category S*.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/symbols
func (c_ CharacterSet) SymbolCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("symbolCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: symbolCharacterSet */


// A character set containing the characters in Unicode General Category Lu and Lt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/uppercaseLetters
func (c_ CharacterSet) UppercaseLetterCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("uppercaseLetterCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: uppercaseLetterCharacterSet */


// Returns the character set for characters allowed in a fragment URL component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlFragmentAllowed
func (c_ CharacterSet) URLFragmentAllowedCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("URLFragmentAllowedCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: URLFragmentAllowedCharacterSet */


// Returns the character set for characters allowed in a host URL subcomponent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlHostAllowed
func (c_ CharacterSet) URLHostAllowedCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("URLHostAllowedCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: URLHostAllowedCharacterSet */


// Returns the character set for characters allowed in a password URL subcomponent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlPasswordAllowed
func (c_ CharacterSet) URLPasswordAllowedCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("URLPasswordAllowedCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: URLPasswordAllowedCharacterSet */


// Returns the character set for characters allowed in a path URL component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlPathAllowed
func (c_ CharacterSet) URLPathAllowedCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("URLPathAllowedCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: URLPathAllowedCharacterSet */


// Returns the character set for characters allowed in a query URL component.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlQueryAllowed
func (c_ CharacterSet) URLQueryAllowedCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("URLQueryAllowedCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: URLQueryAllowedCharacterSet */


// Returns the character set for characters allowed in a user URL subcomponent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlUserAllowed
func (c_ CharacterSet) URLUserAllowedCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("URLUserAllowedCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: URLUserAllowedCharacterSet */


// A character set containing the characters in Unicode General Category Zs and ( ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/whitespaces
func (c_ CharacterSet) WhitespaceCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("whitespaceCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: whitespaceCharacterSet */


// A character set containing characters in Unicode General Category Z*, ~ , and .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/whitespacesAndNewlines
func (c_ CharacterSet) WhitespaceAndNewlineCharacterSet() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("whitespaceAndNewlineCharacterSet"))
	return rv
}/* debug [instance_properties/getter]: whitespaceAndNewlineCharacterSet */


// A character set containing only characters that don’t exist in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/inverted
func (c_ CharacterSet) Inverted() ICharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("inverted"))
	return rv
}/* debug [instance_properties/getter]: inverted */


// A character set containing only characters that don’t exist in the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/inverted
func (c_ CharacterSet) SetInverted(value ICharacterSet) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInverted:"), value)
}/* debug [instance_properties/setter]: inverted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCharacterSet */


