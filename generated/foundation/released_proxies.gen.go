// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class releasedProxies */


/* debug [class_header]: Header for releasedProxies */
// The class instance for the [releasedProxies] class.
var (
	ReleasedProxiesClass     _releasedProxiesClass
	ReleasedProxiesClassOnce sync.Once
)

func getreleasedProxiesClass() _releasedProxiesClass {
	ReleasedProxiesClassOnce.Do(func() {
		ReleasedProxiesClass = _releasedProxiesClass{objc.GetClass("releasedProxies")}
	})
	return ReleasedProxiesClass
}

type _releasedProxiesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for releasedProxies */
// An interface definition for the [releasedProxies] class.
type IreleasedProxies interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for releasedProxies */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for releasedProxies */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for releasedProxies */
// Alloc allocates a new instance without initialization.
func (rc _releasedProxiesClass) Alloc() releasedProxies {
	rv := objc.Send[releasedProxies](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _releasedProxiesClass) New() releasedProxies {
	rv := objc.Send[releasedProxies](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ releasedProxies) Init() releasedProxies {
	rv := objc.Send[releasedProxies](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ releasedProxies) Autorelease() releasedProxies {
	rv := objc.Send[releasedProxies](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewreleasedProxies creates a new releasedProxies instance.
func NewreleasedProxies() releasedProxies {
	return getreleasedProxiesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for releasedProxies */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/releasedProxies
type releasedProxies struct {
	objectivec.Object
}

// releasedProxiesFrom constructs a [releasedProxies] from an unsafe.Pointer.
func releasedProxiesFrom(ptr unsafe.Pointer) releasedProxies {
	return releasedProxies{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for releasedProxies *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for releasedProxies */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for releasedProxies */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for releasedProxies */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for releasedProxies */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class releasedProxies */



