// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AdaptiveImageGlyph] class.
var (
	AdaptiveImageGlyphClass     _AdaptiveImageGlyphClass
	AdaptiveImageGlyphClassOnce sync.Once
)

func getAdaptiveImageGlyphClass() _AdaptiveImageGlyphClass {
	AdaptiveImageGlyphClassOnce.Do(func() {
		AdaptiveImageGlyphClass = _AdaptiveImageGlyphClass{objc.GetClass("NSAdaptiveImageGlyph")}
	})
	return AdaptiveImageGlyphClass
}

type _AdaptiveImageGlyphClass struct {
	class objc.Class
}





// An interface definition for the [AdaptiveImageGlyph] class.
type IAdaptiveImageGlyph interface {
	objectivec.IObject
	

	// properties:
	ContentDescription() foundation.foundation.INSString
	ContentIdentifier() foundation.foundation.INSString
	ImageContent() foundation.foundation.INSData


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AdaptiveImageGlyphClass) Alloc() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AdaptiveImageGlyphClass) New() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdaptiveImageGlyph) Init() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdaptiveImageGlyph) Autorelease() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdaptiveImageGlyph creates a new AdaptiveImageGlyph instance.
func NewAdaptiveImageGlyph() AdaptiveImageGlyph {
	return getAdaptiveImageGlyphClass().New()
}





// A data object for an emoji-like image that can appear in attributed text.
//
// An contains an image that automatically adapts to different sizes and resolutions. The text system creates instances of this type to represent custom emojis that people create using the system interfaces. This type manages multiple images, along with metadata describing how to adapt those images correctly to different fonts and font attributes. Typically, you receive new objects only from the text-input system. When someone creates a new emoji and inserts it into their text, TextKit creates an instance of this type to represent it. If your app examines or changes the attributes of attributed strings, preserve the attribute when making any changes. For example, if you filter unknown attributes in a custom text-storage object, update your code to preserve this attribute. The value of the attribute is an containing the emoji data. You can save the image data with the rest of your content and use the data to recreate the type later.


// A data object for an emoji-like image that can appear in attributed text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph
type AdaptiveImageGlyph struct {
	objectivec.Object
}

// AdaptiveImageGlyphFrom constructs a [AdaptiveImageGlyph] from an unsafe.Pointer.
//
// A data object for an emoji-like image that can appear in attributed text.
func AdaptiveImageGlyphFrom(ptr unsafe.Pointer) AdaptiveImageGlyph {
	return AdaptiveImageGlyph{objectivec.Object{objc.ID(ptr)}}
}






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/init(coder:)
func NewAdaptiveImageGlyphWithCoder(coder foundation.foundation.INSCoder) AdaptiveImageGlyph {
	instance := getAdaptiveImageGlyphClass().Alloc()
	rv := objc.Send[AdaptiveImageGlyph](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Create an adaptive image glyph from the previously saved data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/init(imageContent:)
func NewAdaptiveImageGlyphWithImageContent(imageContent foundation.foundation.INSData) AdaptiveImageGlyph {
	instance := getAdaptiveImageGlyphClass().Alloc()
	rv := objc.Send[AdaptiveImageGlyph](instance.ID, objc.Sel("initWithImageContent:"), imageContent)
	rv.Autorelease()
	return rv
}












// The image data format to use for this image type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentType
func (ac _AdaptiveImageGlyphClass) ContentType() uniformtypeidentifiers.uniformtypeidentifiers.IUTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](objc.ID(ac.class), objc.Sel("contentType"))
	return rv
}











// An alternate textual description of the image contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentDescription
func (a_ AdaptiveImageGlyph) ContentDescription() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("contentDescription"))
	return rv
}


// A unique identifier for this image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentIdentifier
func (a_ AdaptiveImageGlyph) ContentIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("contentIdentifier"))
	return rv
}


// The image data format to use for this image type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentType
func (a_ AdaptiveImageGlyph) ContentType() uniformtypeidentifiers.uniformtypeidentifiers.IUTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](a_.ID, objc.Sel("contentType"))
	return rv
}


// The raw data for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/imageContent
func (a_ AdaptiveImageGlyph) ImageContent() foundation.foundation.INSData {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("imageContent"))
	return rv
}







