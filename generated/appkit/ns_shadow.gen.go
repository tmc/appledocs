// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Shadow] class.
var (
	ShadowClass     _ShadowClass
	ShadowClassOnce sync.Once
)

func getShadowClass() _ShadowClass {
	ShadowClassOnce.Do(func() {
		ShadowClass = _ShadowClass{objc.GetClass("NSShadow")}
	})
	return ShadowClass
}

type _ShadowClass struct {
	class objc.Class
}

// An interface definition for the [Shadow] class.
type IShadow interface {
	objectivec.IObject
	// properties:
	ShadowBlurRadius() float64
	SetShadowBlurRadius(value float64)
	ShadowColor() IColor
	SetShadowColor(value IColor)
	ShadowOffset() objc.IObject /* cross-framework: Size */
	SetShadowOffset(value objc.IObject /* cross-framework: Size */)
	// methods:
	Set()
}

// An object you use to specify attributes to create and style a drop shadow during drawing operations.
//
// When you create shadows, the system draws them in the default user coordinate space, where coordinates are independent from the pixel values of any particular device. Rotations, translations, and other transformations of the current transformation matrix (CTM) don’t affect the shadow or the apparent position of the shadow’s light source. A shadow has two positional parameters: an x-offset and a y-offset. Express these values with a single size data type ( in iOS, in macOS), using the units of the default user coordinate space. Positive values for these offsets extend down and to the right from the user’s perspective. In addition to its positional parameters, a shadow also contains a blur radius, which specifies how much the system blurs a drawn object’s image mask before compositing the image onto the destination. A value of produces no blur. Larger values produce an increasingly large blurred shadow. You can use an object in one of two ways. First, you can set it, like a color or a font, where attributes apply to everything you draw until you apply another shadow or restore a previous graphics state. Second, you can use an instance as the value for the text attribute, so the system applies the shadow to the glyphs corresponding to the characters bearing this attribute.


// An object you use to specify attributes to create and style a drop shadow during drawing operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShadow
type Shadow struct {
	objectivec.Object
}

// ShadowFrom constructs a [Shadow] from an unsafe.Pointer.
//
// An object you use to specify attributes to create and style a drop shadow during drawing operations.
func ShadowFrom(ptr unsafe.Pointer) Shadow {
	return Shadow{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ShadowClass) Alloc() Shadow {
	rv := objc.Send[Shadow](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ShadowClass) New() Shadow {
	rv := objc.Send[Shadow](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Shadow) Init() Shadow {
	rv := objc.Send[Shadow](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Shadow) Autorelease() Shadow {
	rv := objc.Send[Shadow](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewShadow creates a new Shadow instance.
func NewShadow() Shadow {
	return getShadowClass().New()
}




// Sets the shadow of subsequent drawing operations to the current shadow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShadow/set()
func (s_ Shadow) Set() {
	objc.Send[objc.ID](s_.ID, objc.Sel("set"))
}


// The blur radius of the shadow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShadow/shadowBlurRadius
func (s_ Shadow) ShadowBlurRadius() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("shadowBlurRadius"))
	return rv
}


// The blur radius of the shadow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShadow/shadowBlurRadius
func (s_ Shadow) SetShadowBlurRadius(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShadowBlurRadius:"), value)
}


// The color of the shadow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShadow/shadowColor
func (s_ Shadow) ShadowColor() IColor {
	rv := objc.Send[Color](s_.ID, objc.Sel("shadowColor"))
	return rv
}


// The color of the shadow.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShadow/shadowColor
func (s_ Shadow) SetShadowColor(value IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShadowColor:"), value)
}


// The shadow’s relative position, which you specify with horizontal and vertical offset values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShadow/shadowOffset
func (s_ Shadow) ShadowOffset() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](s_.ID, objc.Sel("shadowOffset"))
	return rv
}


// The shadow’s relative position, which you specify with horizontal and vertical offset values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSShadow/shadowOffset
func (s_ Shadow) SetShadowOffset(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setShadowOffset:"), value)
}


