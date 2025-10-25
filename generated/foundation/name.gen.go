// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class name */


/* debug [class_header]: Header for name */
// The class instance for the [name] class.
var (
	NameClass     _nameClass
	NameClassOnce sync.Once
)

func getnameClass() _nameClass {
	NameClassOnce.Do(func() {
		NameClass = _nameClass{objc.GetClass("name")}
	})
	return NameClass
}

type _nameClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for name */
// An interface definition for the [name] class.
type Iname interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for name */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for name */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for name */
// Alloc allocates a new instance without initialization.
func (nc _nameClass) Alloc() name {
	rv := objc.Send[name](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _nameClass) New() name {
	rv := objc.Send[name](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ name) Init() name {
	rv := objc.Send[name](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ name) Autorelease() name {
	rv := objc.Send[name](n_.ID, objc.Sel("autorelease"))
	return rv
}

// Newname creates a new name instance.
func Newname() name {
	return getnameClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSException/name-c.ivar
type name struct {
	objectivec.Object
}

// nameFrom constructs a [name] from an unsafe.Pointer.
func nameFrom(ptr unsafe.Pointer) name {
	return name{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for name *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for name */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for name */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for name */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for name */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class name */



