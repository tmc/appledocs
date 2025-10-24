// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class datax */


/* debug [class_header]: Header for datax */
// The class instance for the [datax] class.
var (
	DataxClass     _dataxClass
	DataxClassOnce sync.Once
)

func getdataxClass() _dataxClass {
	DataxClassOnce.Do(func() {
		DataxClass = _dataxClass{objc.GetClass("datax")}
	})
	return DataxClass
}

type _dataxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for datax */
// An interface definition for the [datax] class.
type Idatax interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for datax */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for datax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for datax */
// Alloc allocates a new instance without initialization.
func (dc _dataxClass) Alloc() datax {
	rv := objc.Send[datax](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _dataxClass) New() datax {
	rv := objc.Send[datax](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ datax) Init() datax {
	rv := objc.Send[datax](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ datax) Autorelease() datax {
	rv := objc.Send[datax](d_.ID, objc.Sel("autorelease"))
	return rv
}

// Newdatax creates a new datax instance.
func Newdatax() datax {
	return getdataxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for datax */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/datax
type datax struct {
	objectivec.Object
}

// dataxFrom constructs a [datax] from an unsafe.Pointer.
func dataxFrom(ptr unsafe.Pointer) datax {
	return datax{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for datax *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for datax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for datax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for datax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for datax */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class datax */



