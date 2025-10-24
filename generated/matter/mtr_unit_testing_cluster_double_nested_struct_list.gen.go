// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterDoubleNestedStructList */


/* debug [class_header]: Header for MTRUnitTestingClusterDoubleNestedStructList */
// The class instance for the [MTRUnitTestingClusterDoubleNestedStructList] class.
var (
	MTRUnitTestingClusterDoubleNestedStructListClass     _MTRUnitTestingClusterDoubleNestedStructListClass
	MTRUnitTestingClusterDoubleNestedStructListClassOnce sync.Once
)

func getMTRUnitTestingClusterDoubleNestedStructListClass() _MTRUnitTestingClusterDoubleNestedStructListClass {
	MTRUnitTestingClusterDoubleNestedStructListClassOnce.Do(func() {
		MTRUnitTestingClusterDoubleNestedStructListClass = _MTRUnitTestingClusterDoubleNestedStructListClass{objc.GetClass("MTRUnitTestingClusterDoubleNestedStructList")}
	})
	return MTRUnitTestingClusterDoubleNestedStructListClass
}

type _MTRUnitTestingClusterDoubleNestedStructListClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterDoubleNestedStructList */
// An interface definition for the [MTRUnitTestingClusterDoubleNestedStructList] class.
type IMTRUnitTestingClusterDoubleNestedStructList interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterDoubleNestedStructList */
	// properties:
	A() objc.IObject /* cross-framework: NSArray */
	SetA(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterDoubleNestedStructList */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterDoubleNestedStructList */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterDoubleNestedStructListClass) Alloc() MTRUnitTestingClusterDoubleNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterDoubleNestedStructList](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterDoubleNestedStructListClass) New() MTRUnitTestingClusterDoubleNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterDoubleNestedStructList](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterDoubleNestedStructList) Init() MTRUnitTestingClusterDoubleNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterDoubleNestedStructList](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterDoubleNestedStructList) Autorelease() MTRUnitTestingClusterDoubleNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterDoubleNestedStructList](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterDoubleNestedStructList creates a new MTRUnitTestingClusterDoubleNestedStructList instance.
func NewMTRUnitTestingClusterDoubleNestedStructList() MTRUnitTestingClusterDoubleNestedStructList {
	return getMTRUnitTestingClusterDoubleNestedStructListClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterDoubleNestedStructList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterDoubleNestedStructList
type MTRUnitTestingClusterDoubleNestedStructList struct {
	objectivec.Object
}

// MTRUnitTestingClusterDoubleNestedStructListFrom constructs a [MTRUnitTestingClusterDoubleNestedStructList] from an unsafe.Pointer.
func MTRUnitTestingClusterDoubleNestedStructListFrom(ptr unsafe.Pointer) MTRUnitTestingClusterDoubleNestedStructList {
	return MTRUnitTestingClusterDoubleNestedStructList{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterDoubleNestedStructList *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterDoubleNestedStructList */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterDoubleNestedStructList */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterDoubleNestedStructList */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterDoubleNestedStructList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterDoubleNestedStructList/a
func (m_ MTRUnitTestingClusterDoubleNestedStructList) A() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("a"))
	return rv
}/* debug [instance_properties/getter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterDoubleNestedStructList/a
func (m_ MTRUnitTestingClusterDoubleNestedStructList) SetA(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}/* debug [instance_properties/setter]: a */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterDoubleNestedStructList */



