// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterNestedStruct */


/* debug [class_header]: Header for MTRUnitTestingClusterNestedStruct */
// The class instance for the [MTRUnitTestingClusterNestedStruct] class.
var (
	MTRUnitTestingClusterNestedStructClass     _MTRUnitTestingClusterNestedStructClass
	MTRUnitTestingClusterNestedStructClassOnce sync.Once
)

func getMTRUnitTestingClusterNestedStructClass() _MTRUnitTestingClusterNestedStructClass {
	MTRUnitTestingClusterNestedStructClassOnce.Do(func() {
		MTRUnitTestingClusterNestedStructClass = _MTRUnitTestingClusterNestedStructClass{objc.GetClass("MTRUnitTestingClusterNestedStruct")}
	})
	return MTRUnitTestingClusterNestedStructClass
}

type _MTRUnitTestingClusterNestedStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterNestedStruct */
// An interface definition for the [MTRUnitTestingClusterNestedStruct] class.
type IMTRUnitTestingClusterNestedStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterNestedStruct */
	// properties:
	A() objc.IObject /* cross-framework: NSNumber */
	SetA(value objc.IObject /* cross-framework: NSNumber */)
	B() objc.IObject /* cross-framework: NSNumber */
	SetB(value objc.IObject /* cross-framework: NSNumber */)
	C() IMTRUnitTestingClusterSimpleStruct
	SetC(value IMTRUnitTestingClusterSimpleStruct)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterNestedStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterNestedStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterNestedStructClass) Alloc() MTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterNestedStructClass) New() MTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterNestedStruct) Init() MTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterNestedStruct) Autorelease() MTRUnitTestingClusterNestedStruct {
	rv := objc.Send[MTRUnitTestingClusterNestedStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterNestedStruct creates a new MTRUnitTestingClusterNestedStruct instance.
func NewMTRUnitTestingClusterNestedStruct() MTRUnitTestingClusterNestedStruct {
	return getMTRUnitTestingClusterNestedStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterNestedStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStruct
type MTRUnitTestingClusterNestedStruct struct {
	objectivec.Object
}

// MTRUnitTestingClusterNestedStructFrom constructs a [MTRUnitTestingClusterNestedStruct] from an unsafe.Pointer.
func MTRUnitTestingClusterNestedStructFrom(ptr unsafe.Pointer) MTRUnitTestingClusterNestedStruct {
	return MTRUnitTestingClusterNestedStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterNestedStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterNestedStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterNestedStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterNestedStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterNestedStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStruct/a
func (m_ MTRUnitTestingClusterNestedStruct) A() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("a"))
	return rv
}/* debug [instance_properties/getter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStruct/a
func (m_ MTRUnitTestingClusterNestedStruct) SetA(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}/* debug [instance_properties/setter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStruct/b
func (m_ MTRUnitTestingClusterNestedStruct) B() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("b"))
	return rv
}/* debug [instance_properties/getter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStruct/b
func (m_ MTRUnitTestingClusterNestedStruct) SetB(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}/* debug [instance_properties/setter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStruct/c
func (m_ MTRUnitTestingClusterNestedStruct) C() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("c"))
	return rv
}/* debug [instance_properties/getter]: c */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStruct/c
func (m_ MTRUnitTestingClusterNestedStruct) SetC(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}/* debug [instance_properties/setter]: c */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterNestedStruct */



