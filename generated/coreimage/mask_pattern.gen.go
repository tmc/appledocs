// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [maskPattern] class.
var (
	MaskPatternClass     _maskPatternClass
	MaskPatternClassOnce sync.Once
)

func getmaskPatternClass() _maskPatternClass {
	MaskPatternClassOnce.Do(func() {
		MaskPatternClass = _maskPatternClass{objc.GetClass("maskPattern")}
	})
	return MaskPatternClass
}

type _maskPatternClass struct {
	class objc.Class
}





// An interface definition for the [maskPattern] class.
type ImaskPattern interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _maskPatternClass) Alloc() maskPattern {
	rv := objc.Send[maskPattern](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _maskPatternClass) New() maskPattern {
	rv := objc.Send[maskPattern](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ maskPattern) Init() maskPattern {
	rv := objc.Send[maskPattern](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ maskPattern) Autorelease() maskPattern {
	rv := objc.Send[maskPattern](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmaskPattern creates a new maskPattern instance.
func NewmaskPattern() maskPattern {
	return getmaskPatternClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/maskPattern-c.ivar
type maskPattern struct {
	objectivec.Object
}

// maskPatternFrom constructs a [maskPattern] from an unsafe.Pointer.
func maskPatternFrom(ptr unsafe.Pointer) maskPattern {
	return maskPattern{objectivec.Object{objc.ID(ptr)}}
}































