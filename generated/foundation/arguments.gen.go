// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class arguments */


/* debug [class_header]: Header for arguments */
// The class instance for the [arguments] class.
var (
	ArgumentsClass     _argumentsClass
	ArgumentsClassOnce sync.Once
)

func getargumentsClass() _argumentsClass {
	ArgumentsClassOnce.Do(func() {
		ArgumentsClass = _argumentsClass{objc.GetClass("arguments")}
	})
	return ArgumentsClass
}

type _argumentsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for arguments */
// An interface definition for the [arguments] class.
type Iarguments interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for arguments */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for arguments */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for arguments */
// Alloc allocates a new instance without initialization.
func (ac _argumentsClass) Alloc() arguments {
	rv := objc.Send[arguments](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _argumentsClass) New() arguments {
	rv := objc.Send[arguments](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ arguments) Init() arguments {
	rv := objc.Send[arguments](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ arguments) Autorelease() arguments {
	rv := objc.Send[arguments](a_.ID, objc.Sel("autorelease"))
	return rv
}

// Newarguments creates a new arguments instance.
func Newarguments() arguments {
	return getargumentsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for arguments */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProcessInfo/arguments-c.ivar
type arguments struct {
	objectivec.Object
}

// argumentsFrom constructs a [arguments] from an unsafe.Pointer.
func argumentsFrom(ptr unsafe.Pointer) arguments {
	return arguments{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for arguments *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for arguments */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for arguments */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for arguments */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for arguments */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class arguments */



