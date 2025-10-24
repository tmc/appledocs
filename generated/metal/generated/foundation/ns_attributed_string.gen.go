// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AttributedString] class.
var (
	AttributedStringClass     _AttributedStringClass
	AttributedStringClassOnce sync.Once
)

func getAttributedStringClass() _AttributedStringClass {
	AttributedStringClassOnce.Do(func() {
		AttributedStringClass = _AttributedStringClass{objc.GetClass("NSAttributedString")}
	})
	return AttributedStringClass
}

type _AttributedStringClass struct {
	class objc.Class
}

// An interface definition for the [AttributedString] class.
type IAttributedString interface {
	objectivec.IObject
	// properties:
	Length() int
	SetLength(value int)
	String() IString
	SetString(value IString)
	// methods:
}

// A string of text that manages data, layout, and stylistic information for ranges of characters to support rendering.
//
// is a type you use to manage strings of stylized Unicode text. In addition to text, an attributed string contains key-value pairs known as that specify additional information to apply to ranges of characters within the string. Attributed strings support many different kinds of attributes, including: Rendering attributes that specify font, color, kern, ligature, and other details Attributes for attachments and adaptive image glyphs Semantic attributes such as link URLs or tool-tip information Language attributes to support automatic gender agreement and text layout Accessibility attributes that provide information for assistive technologies Attributes that summarize details of the Markdown import process Custom attributes you define for your app Use attributed strings anywhere you need styled text, or when you need to associate additional information with your text. Because is an immutable type, you specify all of the text and attributes for it at creation time and can’t change them later. You can create attributed strings directly from a string of characters and a dictionary of attributes. You can also create attributed strings from the contents of a file, including files that contain RTF, RTFD, HTML, Markdown, or other file formats. If you need to modify the contents of an attributed string later, use the type instead. If you create an without any font information, the string’s default font is Helvetica 12-point, which might differ from the default system font for the platform. To change the font, specify a font attribute at creation time.


// A string of text that manages data, layout, and stylistic information for ranges of characters to support rendering.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString
type AttributedString struct {
	objectivec.Object
}

// AttributedStringFrom constructs a [AttributedString] from an unsafe.Pointer.
//
// A string of text that manages data, layout, and stylistic information for ranges of characters to support rendering.
func AttributedStringFrom(ptr unsafe.Pointer) AttributedString {
	return AttributedString{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AttributedStringClass) Alloc() AttributedString {
	rv := objc.Send[AttributedString](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AttributedStringClass) New() AttributedString {
	rv := objc.Send[AttributedString](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AttributedString) Init() AttributedString {
	rv := objc.Send[AttributedString](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AttributedString) Autorelease() AttributedString {
	rv := objc.Send[AttributedString](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAttributedString creates a new AttributedString instance.
func NewAttributedString() AttributedString {
	return getAttributedStringClass().New()
}



// Creates an attributed string with the specified text and attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/init(string:attributes:)
func NewAttributedStringWithStringAttributes(str IString, attrs IDictionary) AttributedString {
	instance := getAttributedStringClass().Alloc()
	rv := objc.Send[AttributedString](instance.ID, objc.Sel("initWithString:attributes:"), str, attrs)
	rv.Autorelease()
	return rv
}



// The length of the attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/length
func (a_ AttributedString) Length() int {
	rv := objc.Send[int](a_.ID, objc.Sel("length"))
	return rv
}


// The length of the attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/length
func (a_ AttributedString) SetLength(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLength:"), value)
}


// The character contents of the attributed string as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/string
func (a_ AttributedString) String() IString {
	rv := objc.Send[String](a_.ID, objc.Sel("string"))
	return rv
}


// The character contents of the attributed string as a string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsattributedstring/string
func (a_ AttributedString) SetString(value IString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setString:"), value)
}


