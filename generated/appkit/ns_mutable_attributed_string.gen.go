// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSMutableAttributedString */


/* debug [class_header]: Header for NSMutableAttributedString */
// The class instance for the [MutableAttributedString] class.
var (
	MutableAttributedStringClass     _MutableAttributedStringClass
	MutableAttributedStringClassOnce sync.Once
)

func getMutableAttributedStringClass() _MutableAttributedStringClass {
	MutableAttributedStringClassOnce.Do(func() {
		MutableAttributedStringClass = _MutableAttributedStringClass{objc.GetClass("NSMutableAttributedString")}
	})
	return MutableAttributedStringClass
}

type _MutableAttributedStringClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableAttributedString */
// An interface definition for the [MutableAttributedString] class.
type IMutableAttributedString interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MutableAttributedString */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableAttributedString */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableAttributedString */
// Alloc allocates a new instance without initialization.
func (mc _MutableAttributedStringClass) Alloc() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableAttributedStringClass) New() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableAttributedString) Init() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableAttributedString) Autorelease() MutableAttributedString {
	rv := objc.Send[MutableAttributedString](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableAttributedString creates a new MutableAttributedString instance.
func NewMutableAttributedString() MutableAttributedString {
	return getMutableAttributedStringClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableAttributedString */
// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type MutableAttributedString struct {
	objectivec.Object
}

// MutableAttributedStringFrom constructs a [MutableAttributedString] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
func MutableAttributedStringFrom(ptr unsafe.Pointer) MutableAttributedString {
	return MutableAttributedString{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableAttributedString *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableAttributedString */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableAttributedString */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableAttributedString */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableAttributedString */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMutableAttributedString */



