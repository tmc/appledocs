// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class localPort */


/* debug [class_header]: Header for localPort */
// The class instance for the [localPort] class.
var (
	LocalPortClass     _localPortClass
	LocalPortClassOnce sync.Once
)

func getlocalPortClass() _localPortClass {
	LocalPortClassOnce.Do(func() {
		LocalPortClass = _localPortClass{objc.GetClass("localPort")}
	})
	return LocalPortClass
}

type _localPortClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for localPort */
// An interface definition for the [localPort] class.
type IlocalPort interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for localPort */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for localPort */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for localPort */
// Alloc allocates a new instance without initialization.
func (lc _localPortClass) Alloc() localPort {
	rv := objc.Send[localPort](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _localPortClass) New() localPort {
	rv := objc.Send[localPort](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ localPort) Init() localPort {
	rv := objc.Send[localPort](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ localPort) Autorelease() localPort {
	rv := objc.Send[localPort](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewlocalPort creates a new localPort instance.
func NewlocalPort() localPort {
	return getlocalPortClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for localPort */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPortMessage/localPort
type localPort struct {
	objectivec.Object
}

// localPortFrom constructs a [localPort] from an unsafe.Pointer.
func localPortFrom(ptr unsafe.Pointer) localPort {
	return localPort{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for localPort *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for localPort */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for localPort */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for localPort */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for localPort */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class localPort */



