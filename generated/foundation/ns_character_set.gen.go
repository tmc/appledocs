// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [CharacterSet] class.
type ICharacterSet interface {
	objectivec.IObject
}

// An object representing a fixed set of Unicode character values for use in search operations.
//
// In Swift, this bridges to a ; use when you need reference semantics or other Foundation-specific behavior. An object represents a set of Unicode-compliant characters. and objects use objects to group characters together for searching operations, so that they can find any of a particular set of characters during a search. The cluster’s two public classes, and , declare the programmatic interface for static and dynamic character sets, respectively. The objects you create using these classes are referred to as character set objects (and when no confusion will result, merely as character sets). Because of the nature of class clusters, character set objects aren’t actual instances of the or classes but of one of their private subclasses. Although a character set object’s class is private, its interface is public, as declared by these abstract superclasses, and . The character set classes adopt the and protocols, making it convenient to convert a character set of one type to the other. The class declares the programmatic interface for an object that manages a set of Unicode characters (see the class cluster specification for information on Unicode). ’s principal primitive method, , provides the basis for all other instance methods in its interface. A subclass of needs only to implement this method, plus , for proper behavior. For optimal performance, a subclass should also override , which otherwise works by invoking for every possible Unicode value. is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CharacterSetClass) Alloc() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A character set containing the characters in Unicode General Category Cc and Cf.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/controlCharacters
func (cc _CharacterSetClass) ControlCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("controlCharacterSet"))
	return rv
}
// A character set containing values in the category of Non-Characters or that have not yet been defined in version 3.2 of the Unicode standard.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/illegalCharacters
func (cc _CharacterSetClass) IllegalCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("illegalCharacterSet"))
	return rv
}
// A character set containing the characters in Unicode General Category M*.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/nonBaseCharacters
func (cc _CharacterSetClass) NonBaseCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("nonBaseCharacterSet"))
	return rv
}
// A character set containing the characters in Unicode General Category S*.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/symbols
func (cc _CharacterSetClass) SymbolCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("symbolCharacterSet"))
	return rv
}
// Returns the character set for characters allowed in a path URL component.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlPathAllowed
func (cc _CharacterSetClass) URLPathAllowedCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("URLPathAllowedCharacterSet"))
	return rv
}
// Returns the character set for characters allowed in a query URL component.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlQueryAllowed
func (cc _CharacterSetClass) URLQueryAllowedCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("URLQueryAllowedCharacterSet"))
	return rv
}
// A character set containing the characters in Unicode General Category Zs and ( ).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/whitespaces
func (cc _CharacterSetClass) WhitespaceCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("whitespaceCharacterSet"))
	return rv
}
// A character set containing characters in Unicode General Category Z*, ~ , and .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/whitespacesAndNewlines
func (cc _CharacterSetClass) WhitespaceAndNewlineCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("whitespaceAndNewlineCharacterSet"))
	return rv
}
// An
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/bitmaprepresentation
func (c_ CharacterSet) BitmapRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("bitmapRepresentation"))
	return rv
}


// SetBitmapRepresentation sets the value of the bitmapRepresentation property.
// An

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/bitmaprepresentation
func (c_ CharacterSet) SetBitmapRepresentation(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBitmapRepresentation:"), value)
}

// A character set containing only characters that don’t exist in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/inverted
func (c_ CharacterSet) Inverted() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("inverted"))
	return rv
}


// SetInverted sets the value of the inverted property.
// A character set containing only characters that don’t exist in the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nscharacterset/inverted
func (c_ CharacterSet) SetInverted(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInverted:"), value)
}

// A character set containing the characters in Unicode General Category Cc and Cf.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/controlCharacters
func (c_ CharacterSet) ControlCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("controlCharacterSet"))
	return rv
}

// A character set containing values in the category of Non-Characters or that have not yet been defined in version 3.2 of the Unicode standard.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/illegalCharacters
func (c_ CharacterSet) IllegalCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("illegalCharacterSet"))
	return rv
}

// A character set containing only characters that don’t exist in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/inverted
func (c_ CharacterSet) InvertedSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("invertedSet"))
	return rv
}

// A character set containing the characters in Unicode General Category M*.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/nonBaseCharacters
func (c_ CharacterSet) NonBaseCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("nonBaseCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category S*.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/symbols
func (c_ CharacterSet) SymbolCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("symbolCharacterSet"))
	return rv
}

// Returns the character set for characters allowed in a path URL component.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlPathAllowed
func (c_ CharacterSet) URLPathAllowedCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("URLPathAllowedCharacterSet"))
	return rv
}

// Returns the character set for characters allowed in a query URL component.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/urlQueryAllowed
func (c_ CharacterSet) URLQueryAllowedCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("URLQueryAllowedCharacterSet"))
	return rv
}

// A character set containing the characters in Unicode General Category Zs and ( ).
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/whitespaces
func (c_ CharacterSet) WhitespaceCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("whitespaceCharacterSet"))
	return rv
}

// A character set containing characters in Unicode General Category Z*, ~ , and .
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCharacterSet/whitespacesAndNewlines
func (c_ CharacterSet) WhitespaceAndNewlineCharacterSet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("whitespaceAndNewlineCharacterSet"))
	return rv
}



