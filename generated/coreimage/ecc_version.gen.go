// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class eccVersion */


/* debug [class_header]: Header for eccVersion */
// The class instance for the [eccVersion] class.
var (
	EccVersionClass     _eccVersionClass
	EccVersionClassOnce sync.Once
)

func geteccVersionClass() _eccVersionClass {
	EccVersionClassOnce.Do(func() {
		EccVersionClass = _eccVersionClass{objc.GetClass("eccVersion")}
	})
	return EccVersionClass
}

type _eccVersionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for eccVersion */
// An interface definition for the [eccVersion] class.
type IeccVersion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for eccVersion */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for eccVersion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for eccVersion */
// Alloc allocates a new instance without initialization.
func (ec _eccVersionClass) Alloc() eccVersion {
	rv := objc.Send[eccVersion](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _eccVersionClass) New() eccVersion {
	rv := objc.Send[eccVersion](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ eccVersion) Init() eccVersion {
	rv := objc.Send[eccVersion](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ eccVersion) Autorelease() eccVersion {
	rv := objc.Send[eccVersion](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NeweccVersion creates a new eccVersion instance.
func NeweccVersion() eccVersion {
	return geteccVersionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for eccVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIDataMatrixCodeDescriptor/eccVersion-c.ivar
type eccVersion struct {
	objectivec.Object
}

// eccVersionFrom constructs a [eccVersion] from an unsafe.Pointer.
func eccVersionFrom(ptr unsafe.Pointer) eccVersion {
	return eccVersion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for eccVersion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for eccVersion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for eccVersion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for eccVersion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for eccVersion */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class eccVersion */



