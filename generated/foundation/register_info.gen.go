// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class registerInfo */


/* debug [class_header]: Header for registerInfo */
// The class instance for the [registerInfo] class.
var (
	RegisterInfoClass     _registerInfoClass
	RegisterInfoClassOnce sync.Once
)

func getregisterInfoClass() _registerInfoClass {
	RegisterInfoClassOnce.Do(func() {
		RegisterInfoClass = _registerInfoClass{objc.GetClass("registerInfo")}
	})
	return RegisterInfoClass
}

type _registerInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for registerInfo */
// An interface definition for the [registerInfo] class.
type IregisterInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for registerInfo */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for registerInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for registerInfo */
// Alloc allocates a new instance without initialization.
func (rc _registerInfoClass) Alloc() registerInfo {
	rv := objc.Send[registerInfo](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _registerInfoClass) New() registerInfo {
	rv := objc.Send[registerInfo](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ registerInfo) Init() registerInfo {
	rv := objc.Send[registerInfo](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ registerInfo) Autorelease() registerInfo {
	rv := objc.Send[registerInfo](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewregisterInfo creates a new registerInfo instance.
func NewregisterInfo() registerInfo {
	return getregisterInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for registerInfo */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/registerInfo
type registerInfo struct {
	objectivec.Object
}

// registerInfoFrom constructs a [registerInfo] from an unsafe.Pointer.
func registerInfoFrom(ptr unsafe.Pointer) registerInfo {
	return registerInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for registerInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for registerInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for registerInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for registerInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for registerInfo */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class registerInfo */



