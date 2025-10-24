// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class identifier */


/* debug [class_header]: Header for identifier */
// The class instance for the [identifier] class.
var (
	IdentifierClass     _identifierClass
	IdentifierClassOnce sync.Once
)

func getidentifierClass() _identifierClass {
	IdentifierClassOnce.Do(func() {
		IdentifierClass = _identifierClass{objc.GetClass("identifier")}
	})
	return IdentifierClass
}

type _identifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for identifier */
// An interface definition for the [identifier] class.
type Iidentifier interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for identifier */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for identifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for identifier */
// Alloc allocates a new instance without initialization.
func (ic _identifierClass) Alloc() identifier {
	rv := objc.Send[identifier](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _identifierClass) New() identifier {
	rv := objc.Send[identifier](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ identifier) Init() identifier {
	rv := objc.Send[identifier](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ identifier) Autorelease() identifier {
	rv := objc.Send[identifier](i_.ID, objc.Sel("autorelease"))
	return rv
}

// Newidentifier creates a new identifier instance.
func Newidentifier() identifier {
	return getidentifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for identifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODMappings/identifier-c.ivar
type identifier struct {
	objectivec.Object
}

// identifierFrom constructs a [identifier] from an unsafe.Pointer.
func identifierFrom(ptr unsafe.Pointer) identifier {
	return identifier{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for identifier *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for identifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for identifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for identifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for identifier */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class identifier */



