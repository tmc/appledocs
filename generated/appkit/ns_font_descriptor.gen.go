// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
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
}

// A dictionary of attributes that describe a font.
//
// A font descriptor can be used to create or modify an object. The system provides a font matching capability, so that you can partially describe a font by creating a font descriptor with, for example, just a family name. You can then find all the available fonts on the system with a matching family name using . There are several ways to create a new object. You can use and , , , or . to create a font descriptor based on either your custom attributes dictionary or on a specific font’s name and size. Alternatively you can use one of the instance methods (such as ) to create a modified version of an existing descriptor. The latter methods are useful if you have an existing descriptor and simply want to change one aspect. All attributes in the attributes dictionary are optional.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/fontDescriptorWithFontAttributes:
func (fc _FontDescriptorClass) FontDescriptorWithFontAttributes(attributes unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("fontDescriptorWithFontAttributes:"), attributes)
	return rv
}

// Returns the font attribute specified by the given key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/object(forKey:)
func (f_ FontDescriptor) ObjectForKey(attribute unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](f_.ID, objc.Sel("objectForKey:"), attribute)
	return rv
}

// The current transform matrix of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/matrix
func (f_ FontDescriptor) Matrix() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("matrix"))
	return rv
}

// A bit mask that describes the traits of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/symbolicTraits-swift.property
func (f_ FontDescriptor) SymbolicTraits() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("symbolicTraits"))
	return rv
}



