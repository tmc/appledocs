// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class virtualSubnodes */


/* debug [class_header]: Header for virtualSubnodes */
// The class instance for the [virtualSubnodes] class.
var (
	VirtualSubnodesClass     _virtualSubnodesClass
	VirtualSubnodesClassOnce sync.Once
)

func getvirtualSubnodesClass() _virtualSubnodesClass {
	VirtualSubnodesClassOnce.Do(func() {
		VirtualSubnodesClass = _virtualSubnodesClass{objc.GetClass("virtualSubnodes")}
	})
	return VirtualSubnodesClass
}

type _virtualSubnodesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for virtualSubnodes */
// An interface definition for the [virtualSubnodes] class.
type IvirtualSubnodes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for virtualSubnodes */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for virtualSubnodes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for virtualSubnodes */
// Alloc allocates a new instance without initialization.
func (vc _virtualSubnodesClass) Alloc() virtualSubnodes {
	rv := objc.Send[virtualSubnodes](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _virtualSubnodesClass) New() virtualSubnodes {
	rv := objc.Send[virtualSubnodes](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ virtualSubnodes) Init() virtualSubnodes {
	rv := objc.Send[virtualSubnodes](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ virtualSubnodes) Autorelease() virtualSubnodes {
	rv := objc.Send[virtualSubnodes](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewvirtualSubnodes creates a new virtualSubnodes instance.
func NewvirtualSubnodes() virtualSubnodes {
	return getvirtualSubnodesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for virtualSubnodes */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/virtualSubnodes-c.ivar
type virtualSubnodes struct {
	objectivec.Object
}

// virtualSubnodesFrom constructs a [virtualSubnodes] from an unsafe.Pointer.
func virtualSubnodesFrom(ptr unsafe.Pointer) virtualSubnodes {
	return virtualSubnodes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for virtualSubnodes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for virtualSubnodes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for virtualSubnodes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for virtualSubnodes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for virtualSubnodes */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class virtualSubnodes */



