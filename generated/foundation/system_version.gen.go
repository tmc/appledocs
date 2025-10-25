// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class systemVersion */


/* debug [class_header]: Header for systemVersion */
// The class instance for the [systemVersion] class.
var (
	SystemVersionClass     _systemVersionClass
	SystemVersionClassOnce sync.Once
)

func getsystemVersionClass() _systemVersionClass {
	SystemVersionClassOnce.Do(func() {
		SystemVersionClass = _systemVersionClass{objc.GetClass("systemVersion")}
	})
	return SystemVersionClass
}

type _systemVersionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for systemVersion */
// An interface definition for the [systemVersion] class.
type IsystemVersion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for systemVersion */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for systemVersion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for systemVersion */
// Alloc allocates a new instance without initialization.
func (sc _systemVersionClass) Alloc() systemVersion {
	rv := objc.Send[systemVersion](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _systemVersionClass) New() systemVersion {
	rv := objc.Send[systemVersion](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ systemVersion) Init() systemVersion {
	rv := objc.Send[systemVersion](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ systemVersion) Autorelease() systemVersion {
	rv := objc.Send[systemVersion](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewsystemVersion creates a new systemVersion instance.
func NewsystemVersion() systemVersion {
	return getsystemVersionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for systemVersion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/systemVersion-c.ivar
type systemVersion struct {
	objectivec.Object
}

// systemVersionFrom constructs a [systemVersion] from an unsafe.Pointer.
func systemVersionFrom(ptr unsafe.Pointer) systemVersion {
	return systemVersion{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for systemVersion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for systemVersion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for systemVersion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for systemVersion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for systemVersion */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class systemVersion */



