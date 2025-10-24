// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAdaptiveImageGlyph */


/* debug [class_header]: Header for NSAdaptiveImageGlyph */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AdaptiveImageGlyph */
// An interface definition for the [AdaptiveImageGlyph] class.
type IAdaptiveImageGlyph interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AdaptiveImageGlyph */
	// properties:
	ContentDescription() objc.IObject /* cross-framework: NSString */
	ContentIdentifier() objc.IObject /* cross-framework: NSString */
	ImageContent() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AdaptiveImageGlyph */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AdaptiveImageGlyph */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AdaptiveImageGlyph */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AdaptiveImageGlyph */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/init(coder:)
func NewAdaptiveImageGlyphWithCoder(coder foundation.Coder) AdaptiveImageGlyph {
	instance := getAdaptiveImageGlyphClass().Alloc()
	rv := objc.Send[AdaptiveImageGlyph](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAdaptiveImageGlyphWithCoder */


// Create an adaptive image glyph from the previously saved data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/init(imageContent:)
func NewAdaptiveImageGlyphWithImageContent(imageContent objc.IObject /* cross-framework: NSData */) AdaptiveImageGlyph {
	instance := getAdaptiveImageGlyphClass().Alloc()
	rv := objc.Send[AdaptiveImageGlyph](instance.ID, objc.Sel("initWithImageContent:"), imageContent)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAdaptiveImageGlyphWithImageContent */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AdaptiveImageGlyph */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AdaptiveImageGlyph */

// The image data format to use for this image type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentType
func (ac _AdaptiveImageGlyphClass) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](objc.ID(ac.class), objc.Sel("contentType"))
	return rv
}/* debug [class_properties_class/property]: contentType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AdaptiveImageGlyph */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AdaptiveImageGlyph */

// An alternate textual description of the image contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentDescription
func (a_ AdaptiveImageGlyph) ContentDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("contentDescription"))
	return rv
}/* debug [instance_properties/getter]: contentDescription */


// A unique identifier for this image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentIdentifier
func (a_ AdaptiveImageGlyph) ContentIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("contentIdentifier"))
	return rv
}/* debug [instance_properties/getter]: contentIdentifier */


// The image data format to use for this image type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/contentType
func (a_ AdaptiveImageGlyph) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](a_.ID, objc.Sel("contentType"))
	return rv
}/* debug [instance_properties/getter]: contentType */


// The raw data for the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAdaptiveImageGlyph/imageContent
func (a_ AdaptiveImageGlyph) ImageContent() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("imageContent"))
	return rv
}/* debug [instance_properties/getter]: imageContent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAdaptiveImageGlyph */


