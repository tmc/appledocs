// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FontDescriptor] class.
var (
	FontDescriptorClass     _FontDescriptorClass
	FontDescriptorClassOnce sync.Once
)

func getFontDescriptorClass() _FontDescriptorClass {
	FontDescriptorClassOnce.Do(func() {
		FontDescriptorClass = _FontDescriptorClass{objc.GetClass("NSFontDescriptor")}
	})
	return FontDescriptorClass
}

type _FontDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [FontDescriptor] class.
type IFontDescriptor interface {
	objectivec.IObject
	// properties:
	FontAttributes() foundation.IDictionary
	Matrix() objc.IObject /* cross-framework: AffineTransform */
	PointSize() float64
	PostscriptName() objc.IObject /* cross-framework: NSString */
	RequiresFontAssetRequest() bool
	SymbolicTraits() FontDescriptorSymbolicTraits
	NSFontFamilyClassMask() unsafe.Pointer
	SetNSFontFamilyClassMask(value unsafe.Pointer)
	// methods:
	FontDescriptorByAddingAttributes(attributes foundation.IDictionary) IFontDescriptor
	MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys unsafe.Pointer) IFontDescriptor
	MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys unsafe.Pointer) []FontDescriptor
	ObjectForKey(attribute objc.IObject /* cross-framework: FontDescriptorAttributeName */) objc.ID
	FontDescriptorWithDesign(design objc.IObject /* cross-framework: FontDescriptorSystemDesign */) unsafe.Pointer
	FontDescriptorWithFace(newFace objc.IObject /* cross-framework: NSString */) IFontDescriptor
	FontDescriptorWithFamily(newFamily objc.IObject /* cross-framework: NSString */) IFontDescriptor
	FontDescriptorWithMatrix(matrix objc.IObject /* cross-framework: AffineTransform */) IFontDescriptor
	FontDescriptorWithSize(newPointSize float64) IFontDescriptor
	FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) IFontDescriptor
}

// A dictionary of attributes that describe a font.
//
// A font descriptor can be used to create or modify an object. The system provides a font matching capability, so that you can partially describe a font by creating a font descriptor with, for example, just a family name. You can then find all the available fonts on the system with a matching family name using . There are several ways to create a new object. You can use and , , , or . to create a font descriptor based on either your custom attributes dictionary or on a specific font’s name and size. Alternatively you can use one of the instance methods (such as ) to create a modified version of an existing descriptor. The latter methods are useful if you have an existing descriptor and simply want to change one aspect. All attributes in the attributes dictionary are optional.


// A dictionary of attributes that describe a font.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor
type FontDescriptor struct {
	objectivec.Object
}

// FontDescriptorFrom constructs a [FontDescriptor] from an unsafe.Pointer.
//
// A dictionary of attributes that describe a font.
func FontDescriptorFrom(ptr unsafe.Pointer) FontDescriptor {
	return FontDescriptor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FontDescriptorClass) Alloc() FontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FontDescriptorClass) New() FontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FontDescriptor) Init() FontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FontDescriptor) Autorelease() FontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFontDescriptor creates a new FontDescriptor instance.
func NewFontDescriptor() FontDescriptor {
	return getFontDescriptorClass().New()
}



// Initializes and returns a new font descriptor with the specified attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/init(fontAttributes:)
func NewFontDescriptorWithFontAttributes(attributes foundation.IDictionary) FontDescriptor {
	instance := getFontDescriptorClass().Alloc()
	rv := objc.Send[FontDescriptor](instance.ID, objc.Sel("initWithFontAttributes:"), attributes)
	rv.Autorelease()
	return rv
}


// Returns a font descriptor with the name and matrix attributes set to the given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/init(name:matrix:)
func NewFontDescriptorWithNameMatrix(fontName objc.IObject /* cross-framework: NSString */, matrix objc.IObject /* cross-framework: AffineTransform */) FontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(getFontDescriptorClass().class), objc.Sel("fontDescriptorWithName:matrix:"), fontName, matrix)
	return rv
}


// Returns a font descriptor with the name and size attributes set to the given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/init(name:size:)
func NewFontDescriptorWithNameSize(fontName objc.IObject /* cross-framework: NSString */, size float64) FontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(getFontDescriptorClass().class), objc.Sel("fontDescriptorWithName:size:"), fontName, size)
	return rv
}



// Returns a font descriptor with a dictionary of attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/fontDescriptorWithFontAttributes:
func (fc _FontDescriptorClass) FontDescriptorWithFontAttributes(attributes foundation.IDictionary) IFontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(fc.class), objc.Sel("fontDescriptorWithFontAttributes:"), attributes)
	return rv
}


// Returns a font descriptor with the name and matrix attributes set to the given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/init(name:matrix:)
func (fc _FontDescriptorClass) FontDescriptorWithNameMatrix(fontName objc.IObject /* cross-framework: NSString */, matrix objc.IObject /* cross-framework: AffineTransform */) IFontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(fc.class), objc.Sel("fontDescriptorWithName:matrix:"), fontName, matrix)
	return rv
}


// Returns a font descriptor with the name and size attributes set to the given values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/init(name:size:)
func (fc _FontDescriptorClass) FontDescriptorWithNameSize(fontName objc.IObject /* cross-framework: NSString */, size float64) IFontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(fc.class), objc.Sel("fontDescriptorWithName:size:"), fontName, size)
	return rv
}


// Returns a font descriptor that contains the text style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/preferredFontDescriptor(forTextStyle:options:)
func (fc _FontDescriptorClass) PreferredFontDescriptorForTextStyleOptions(style objc.IObject /* cross-framework: FontTextStyle */, options foundation.IDictionary) IFontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(fc.class), objc.Sel("preferredFontDescriptorForTextStyle:options:"), style, options)
	return rv
}


// Returns a new font descriptor based on the current object, but with the specified attributes taking precedence over the existing ones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/addingAttributes(_:)
func (f_ FontDescriptor) FontDescriptorByAddingAttributes(attributes foundation.IDictionary) IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("fontDescriptorByAddingAttributes:"), attributes)
	return rv
}


// Returns a normalized font descriptor whose specified attributes match those of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/matchingFontDescriptor(withMandatoryKeys:)
func (f_ FontDescriptor) MatchingFontDescriptorWithMandatoryKeys(mandatoryKeys unsafe.Pointer) IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("matchingFontDescriptorWithMandatoryKeys:"), mandatoryKeys)
	return rv
}


// Returns all the fonts available on the system whose specified attributes match those of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/matchingFontDescriptors(withMandatoryKeys:)
func (f_ FontDescriptor) MatchingFontDescriptorsWithMandatoryKeys(mandatoryKeys unsafe.Pointer) []FontDescriptor {
	rv := objc.Send[[]FontDescriptor](f_.ID, objc.Sel("matchingFontDescriptorsWithMandatoryKeys:"), mandatoryKeys)
	return rv
}


// Returns the font attribute specified by the given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/object(forKey:)
func (f_ FontDescriptor) ObjectForKey(attribute objc.IObject /* cross-framework: FontDescriptorAttributeName */) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("objectForKey:"), attribute)
	return rv
}


// Returns a new font descriptor based on the current object, but with the specified design style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withDesign(_:)
func (f_ FontDescriptor) FontDescriptorWithDesign(design objc.IObject /* cross-framework: FontDescriptorSystemDesign */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fontDescriptorWithDesign:"), design)
	return rv
}


// Returns a new font descriptor based on the current object, but with the specified face.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withFace(_:)
func (f_ FontDescriptor) FontDescriptorWithFace(newFace objc.IObject /* cross-framework: NSString */) IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("fontDescriptorWithFace:"), newFace)
	return rv
}


// Returns a new font descriptor based on the current object, but with the specified font family.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withFamily(_:)
func (f_ FontDescriptor) FontDescriptorWithFamily(newFamily objc.IObject /* cross-framework: NSString */) IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("fontDescriptorWithFamily:"), newFamily)
	return rv
}


// Returns a new font descriptor based on the current object, but with the specified font matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withMatrix(_:)
func (f_ FontDescriptor) FontDescriptorWithMatrix(matrix objc.IObject /* cross-framework: AffineTransform */) IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("fontDescriptorWithMatrix:"), matrix)
	return rv
}


// Returns a new font descriptor based on the current object, but with the specified point size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withSize(_:)
func (f_ FontDescriptor) FontDescriptorWithSize(newPointSize float64) IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("fontDescriptorWithSize:"), newPointSize)
	return rv
}


// Returns a new font descriptor based on the current object, but with the specified symbolic traits taking precedence over the existing ones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/withSymbolicTraits(_:)
func (f_ FontDescriptor) FontDescriptorWithSymbolicTraits(symbolicTraits FontDescriptorSymbolicTraits) IFontDescriptor {
	rv := objc.Send[FontDescriptor](f_.ID, objc.Sel("fontDescriptorWithSymbolicTraits:"), symbolicTraits)
	return rv
}


// The receiver’s dictionary of attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/fontAttributes
func (f_ FontDescriptor) FontAttributes() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](f_.ID, objc.Sel("fontAttributes"))
	return rv
}


// The current transform matrix of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/matrix
func (f_ FontDescriptor) Matrix() objc.IObject /* cross-framework: AffineTransform */ {
	rv := objc.Send[corefoundation.AffineTransform](f_.ID, objc.Sel("matrix"))
	return rv
}


// The point size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/pointSize
func (f_ FontDescriptor) PointSize() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("pointSize"))
	return rv
}


// The PostScript name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/postscriptName
func (f_ FontDescriptor) PostscriptName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](f_.ID, objc.Sel("postscriptName"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/requiresFontAssetRequest
func (f_ FontDescriptor) RequiresFontAssetRequest() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("requiresFontAssetRequest"))
	return rv
}


// A bit mask that describes the traits of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/symbolicTraits-swift.property
func (f_ FontDescriptor) SymbolicTraits() FontDescriptorSymbolicTraits {
	rv := objc.Send[FontDescriptorSymbolicTraits](f_.ID, objc.Sel("symbolicTraits"))
	return rv
}


// Constant you use to access
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontfamilyclassmask
func (f_ FontDescriptor) NSFontFamilyClassMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("NSFontFamilyClassMask"))
	return rv
}


// Constant you use to access
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontfamilyclassmask
func (f_ FontDescriptor) SetNSFontFamilyClassMask(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setNSFontFamilyClassMask:"), value)
}


