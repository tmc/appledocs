// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class cq */


/* debug [class_header]: Header for cq */
// The class instance for the [cq] class.
var (
	CqClass     _cqClass
	CqClassOnce sync.Once
)

func getcqClass() _cqClass {
	CqClassOnce.Do(func() {
		CqClass = _cqClass{objc.GetClass("cq")}
	})
	return CqClass
}

type _cqClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for cq */
// An interface definition for the [cq] class.
type Icq interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for cq */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for cq */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for cq */
// Alloc allocates a new instance without initialization.
func (cc _cqClass) Alloc() cq {
	rv := objc.Send[cq](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _cqClass) New() cq {
	rv := objc.Send[cq](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ cq) Init() cq {
	rv := objc.Send[cq](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ cq) Autorelease() cq {
	rv := objc.Send[cq](c_.ID, objc.Sel("autorelease"))
	return rv
}

// Newcq creates a new cq instance.
func Newcq() cq {
	return getcqClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for cq */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Quartz/QCCompositionRepository/cq
type cq struct {
	objectivec.Object
}

// cqFrom constructs a [cq] from an unsafe.Pointer.
func cqFrom(ptr unsafe.Pointer) cq {
	return cq{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for cq *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for cq */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for cq */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for cq */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for cq */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class cq */



