// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableAttributedString] class.
var (
	MutableAttributedStringClass     _MutableAttributedStringClass
	MutableAttributedStringClassOnce sync.Once
)

func getMutableAttributedStringClass() _MutableAttributedStringClass {
	MutableAttributedStringClassOnce.Do(func() {
		MutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}
	})
	return MutableAttributedStringClass
}

type _MutableAttributedStringClass struct {
	class objc.Class
}

// An interface definition for the [MutableAttributedString] class.
type IMutableAttributedString interface {
	IAttributedString
	AppendAttributedString(attrString IAttributedString)
	AppendLocalizedFormat(format IAttributedString)
	InsertAttributedStringAtIndex(attrString IAttributedString, loc uint)
	ReplaceCharactersInRangeWithAttributedString(range_ IRange, attrString IAttributedString)
	SetAttributedString(attrString IAttributedString)
	MutableString() NSMutableString
	SetMutableString(value IMutableString)
}

// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text.
//
// The class declares additional methods for mutating the content of an attributed string. You can add and remove characters (raw strings) and attributes separately or together as attributed strings. See the class description for for more information about attributed strings. adds two primitive methods to those of . These primitive methods provide the basis for all the other methods in its class. The primitive method replaces a range of characters with those from a string, leaving all attribute information outside that range intact. The primitive method sets attributes and values for a given range of characters, replacing any previous attributes and values for that range. In macOS, AppKit also uses and its subclass to encapsulate the paragraph or ruler attributes used by the classes. Note that the default font for objects is Helvetica 12-point, which may differ from the macOS system font, so you may wish to create the string with non-default attributes suitable for your application using, for example, . is “toll-free bridged” with its Core Foundation counterpart, . See for more information.


// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString

type MutableAttributedString struct {
	AttributedString
}

// MutableAttributedStringFrom constructs a [MutableAttributedString] from an unsafe.Pointer.
//
// A mutable string with associated attributes (such as visual style, hyperlinks, or accessibility data) for portions of its text.
func MutableAttributedStringFrom(ptr unsafe.Pointer) MutableAttributedString {
	return MutableAttributedString{
		AttributedString: AttributedStringFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableAttributedStringClass) Alloc() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableAttributedStringClass) New() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableAttributedString) Init() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableAttributedString) Autorelease() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableAttributedString creates a new MutableAttributedString instance.
func NewMutableAttributedString() MutableAttributedString {
	return getMutableAttributedStringClass().New()
}




// Adds the characters and attributes of a given attributed string to the end of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/append(_:)

func (m_ MutableAttributedString) AppendAttributedString(attrString IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendAttributedString:"), attrString)
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/appendLocalizedFormat:

func (m_ MutableAttributedString) AppendLocalizedFormat(format IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("appendLocalizedFormat:"), format)
}



// Inserts the characters and attributes of the given attributed string into the receiver at the given index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/insert(_:at:)

func (m_ MutableAttributedString) InsertAttributedStringAtIndex(attrString IAttributedString, loc uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("insertAttributedString:atIndex:"), attrString, loc)
}



// Replaces the characters and attributes in a given range with the characters and attributes of the given attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/replaceCharacters(in:with:)-1uaw7

func (m_ MutableAttributedString) ReplaceCharactersInRangeWithAttributedString(range_ IRange, attrString IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceCharactersInRange:withAttributedString:"), range_, attrString)
}



// Replaces the receiver’s entire contents with the characters and attributes of the given attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString/setAttributedString(_:)

func (m_ MutableAttributedString) SetAttributedString(attrString IAttributedString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributedString:"), attrString)
}


// The character contents of the receiver as a mutable string object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/mutablestring

func (m_ MutableAttributedString) MutableString() NSMutableString {
	rv := objc.Send[NSMutableString](m_.ID, objc.Sel("mutableString"))
	return rv
}


// The character contents of the receiver as a mutable string object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmutableattributedstring/mutablestring

func (m_ MutableAttributedString) SetMutableString(value IMutableString) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMutableString:"), value)
}



