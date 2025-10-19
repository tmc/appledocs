// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FontDescriptor] class.
var (
	fontDescriptorClass     _FontDescriptorClass
	fontDescriptorClassOnce sync.Once
)

func getFontDescriptorClass() _FontDescriptorClass {
	fontDescriptorClassOnce.Do(func() {
		fontDescriptorClass = _FontDescriptorClass{objc.GetClass("NSFontDescriptor")}
	})
	return fontDescriptorClass
}

type _FontDescriptorClass struct {
	class objc.Class
}

// An interface definition for the [FontDescriptor] class.
type IFontDescriptor interface {
	objectivec.IObject
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




