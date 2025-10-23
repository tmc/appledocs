// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	ObjectForKey(attribute unsafe.Pointer) objc.ID
	Matrix() coregraphics.AffineTransform
	SymbolicTraits() FontDescriptorSymbolicTraits
	FontAttributes() unsafe.Pointer
	SetFontAttributes(value unsafe.Pointer)
	PointSize() float64
	SetPointSize(value float64)
	PostscriptName() string
	SetPostscriptName(value string)
	RequiresFontAssetRequest() bool
	SetRequiresFontAssetRequest(value bool)
	NSFontFamilyClassMask() unsafe.Pointer
	SetNSFontFamilyClassMask(value unsafe.Pointer)
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



// Returns a font descriptor with a dictionary of attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/fontDescriptorWithFontAttributes:
func (fc _FontDescriptorClass) FontDescriptorWithFontAttributes(attributes unsafe.Pointer) FontDescriptor {
	rv := objc.Send[FontDescriptor](objc.ID(fc.class), objc.Sel("fontDescriptorWithFontAttributes:"), attributes)
	return rv
}


// Returns the font attribute specified by the given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/object(forKey:)
func (f_ FontDescriptor) ObjectForKey(attribute unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("objectForKey:"), attribute)
	return rv
}


// The current transform matrix of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/matrix
func (f_ FontDescriptor) Matrix() coregraphics.AffineTransform {
	rv := objc.Send[coregraphics.AffineTransform](f_.ID, objc.Sel("matrix"))
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


// The receiver’s dictionary of attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/fontattributes
func (f_ FontDescriptor) FontAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("fontAttributes"))
	return rv
}


// The receiver’s dictionary of attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/fontattributes
func (f_ FontDescriptor) SetFontAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setFontAttributes:"), value)
}


// The point size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/pointsize
func (f_ FontDescriptor) PointSize() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("pointSize"))
	return rv
}


// The point size of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/pointsize
func (f_ FontDescriptor) SetPointSize(value float64) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPointSize:"), value)
}


// The PostScript name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/postscriptname
func (f_ FontDescriptor) PostscriptName() string {
	rv := objc.Send[string](f_.ID, objc.Sel("postscriptName"))
	return rv
}


// The PostScript name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/postscriptname
func (f_ FontDescriptor) SetPostscriptName(value string) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setPostscriptName:"), objc.String(value))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/requiresfontassetrequest
func (f_ FontDescriptor) RequiresFontAssetRequest() bool {
	rv := objc.Send[bool](f_.ID, objc.Sel("requiresFontAssetRequest"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsfontdescriptor/requiresfontassetrequest
func (f_ FontDescriptor) SetRequiresFontAssetRequest(value bool) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setRequiresFontAssetRequest:"), value)
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



