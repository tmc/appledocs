// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ids */


/* debug [class_header]: Header for ids */
// The class instance for the [ids] class.
var (
	IdsClass     _idsClass
	IdsClassOnce sync.Once
)

func getidsClass() _idsClass {
	IdsClassOnce.Do(func() {
		IdsClass = _idsClass{objc.GetClass("ids")}
	})
	return IdsClass
}

type _idsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ids */
// An interface definition for the [ids] class.
type Iids interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ids */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ids */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ids */
// Alloc allocates a new instance without initialization.
func (ic _idsClass) Alloc() ids {
	rv := objc.Send[ids](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _idsClass) New() ids {
	rv := objc.Send[ids](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ ids) Init() ids {
	rv := objc.Send[ids](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ ids) Autorelease() ids {
	rv := objc.Send[ids](i_.ID, objc.Sel("autorelease"))
	return rv
}

// Newids creates a new ids instance.
func Newids() ids {
	return getidsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ids */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/ids
type ids struct {
	objectivec.Object
}

// idsFrom constructs a [ids] from an unsafe.Pointer.
func idsFrom(ptr unsafe.Pointer) ids {
	return ids{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ids *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ids */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ids */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ids */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ids */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ids */



