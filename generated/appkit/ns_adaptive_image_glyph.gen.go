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
}

// A data object for an emoji-like image that can appear in attributed text.
//
// An contains an image that automatically adapts to different sizes and resolutions. The text system creates instances of this type to represent custom emojis that people create using the system interfaces. This type manages multiple images, along with metadata describing how to adapt those images correctly to different fonts and font attributes. Typically, you receive new objects only from the text-input system. When someone creates a new emoji and inserts it into their text, TextKit creates an instance of this type to represent it. If your app examines or changes the attributes of attributed strings, preserve the attribute when making any changes. For example, if you filter unknown attributes in a custom text-storage object, update your code to preserve this attribute. The value of the attribute is an containing the emoji data. You can save the image data with the rest of your content and use the data to recreate the type later.
//
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

// Alloc allocates a new instance without initialization.
func (ac _AdaptiveImageGlyphClass) Alloc() AdaptiveImageGlyph {
	rv := objc.Send[AdaptiveImageGlyph](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// An alternate textual description of the image contents.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsadaptiveimageglyph/contentdescription
func (a_ AdaptiveImageGlyph) ContentDescription() string {
	rv := objc.Send[string](a_.ID, objc.Sel("contentDescription"))
	return rv
}


// SetContentDescription sets the value of the contentDescription property.
// An alternate textual description of the image contents.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsadaptiveimageglyph/contentdescription
func (a_ AdaptiveImageGlyph) SetContentDescription(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentDescription:"), objc.String(value))
}

// A unique identifier for this image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsadaptiveimageglyph/contentidentifier
func (a_ AdaptiveImageGlyph) ContentIdentifier() string {
	rv := objc.Send[string](a_.ID, objc.Sel("contentIdentifier"))
	return rv
}


// SetContentIdentifier sets the value of the contentIdentifier property.
// A unique identifier for this image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsadaptiveimageglyph/contentidentifier
func (a_ AdaptiveImageGlyph) SetContentIdentifier(value string) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setContentIdentifier:"), objc.String(value))
}

// The raw data for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsadaptiveimageglyph/imagecontent
func (a_ AdaptiveImageGlyph) ImageContent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("imageContent"))
	return rv
}


// SetImageContent sets the value of the imageContent property.
// The raw data for the image.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsadaptiveimageglyph/imagecontent
func (a_ AdaptiveImageGlyph) SetImageContent(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setImageContent:"), value)
}



