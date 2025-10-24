// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class maskPattern */


/* debug [class_header]: Header for maskPattern */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for maskPattern */
// An interface definition for the [maskPattern] class.
type ImaskPattern interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for maskPattern */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for maskPattern */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for maskPattern */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for maskPattern */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIQRCodeDescriptor/maskPattern-c.ivar
type maskPattern struct {
	objectivec.Object
}

// maskPatternFrom constructs a [maskPattern] from an unsafe.Pointer.
func maskPatternFrom(ptr unsafe.Pointer) maskPattern {
	return maskPattern{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for maskPattern *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for maskPattern */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for maskPattern */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for maskPattern */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for maskPattern */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class maskPattern */



