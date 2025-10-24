// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class uuidString */


/* debug [class_header]: Header for uuidString */
// The class instance for the [uuidString] class.
var (
	UuidStringClass     _uuidStringClass
	UuidStringClassOnce sync.Once
)

func getuuidStringClass() _uuidStringClass {
	UuidStringClassOnce.Do(func() {
		UuidStringClass = _uuidStringClass{objc.GetClass("uuidString")}
	})
	return UuidStringClass
}

type _uuidStringClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for uuidString */
// An interface definition for the [uuidString] class.
type IuuidString interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for uuidString */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for uuidString */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for uuidString */
// Alloc allocates a new instance without initialization.
func (uc _uuidStringClass) Alloc() uuidString {
	rv := objc.Send[uuidString](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (uc _uuidStringClass) New() uuidString {
	rv := objc.Send[uuidString](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ uuidString) Init() uuidString {
	rv := objc.Send[uuidString](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ uuidString) Autorelease() uuidString {
	rv := objc.Send[uuidString](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewuuidString creates a new uuidString instance.
func NewuuidString() uuidString {
	return getuuidStringClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for uuidString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/uuidString-c.ivar
type uuidString struct {
	objectivec.Object
}

// uuidStringFrom constructs a [uuidString] from an unsafe.Pointer.
func uuidStringFrom(ptr unsafe.Pointer) uuidString {
	return uuidString{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for uuidString *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for uuidString */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for uuidString */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for uuidString */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for uuidString */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class uuidString */



