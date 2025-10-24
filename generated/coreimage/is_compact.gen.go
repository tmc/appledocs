// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class isCompact */


/* debug [class_header]: Header for isCompact */
// The class instance for the [isCompact] class.
var (
	IsCompactClass     _isCompactClass
	IsCompactClassOnce sync.Once
)

func getisCompactClass() _isCompactClass {
	IsCompactClassOnce.Do(func() {
		IsCompactClass = _isCompactClass{objc.GetClass("isCompact")}
	})
	return IsCompactClass
}

type _isCompactClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for isCompact */
// An interface definition for the [isCompact] class.
type IisCompact interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for isCompact */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for isCompact */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for isCompact */
// Alloc allocates a new instance without initialization.
func (ic _isCompactClass) Alloc() isCompact {
	rv := objc.Send[isCompact](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _isCompactClass) New() isCompact {
	rv := objc.Send[isCompact](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ isCompact) Init() isCompact {
	rv := objc.Send[isCompact](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ isCompact) Autorelease() isCompact {
	rv := objc.Send[isCompact](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewisCompact creates a new isCompact instance.
func NewisCompact() isCompact {
	return getisCompactClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for isCompact */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/isCompact-c.ivar
type isCompact struct {
	objectivec.Object
}

// isCompactFrom constructs a [isCompact] from an unsafe.Pointer.
func isCompactFrom(ptr unsafe.Pointer) isCompact {
	return isCompact{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for isCompact *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for isCompact */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for isCompact */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for isCompact */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for isCompact */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class isCompact */



