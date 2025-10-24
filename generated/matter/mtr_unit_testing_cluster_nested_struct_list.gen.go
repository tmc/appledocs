// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterNestedStructList */


/* debug [class_header]: Header for MTRUnitTestingClusterNestedStructList */
// The class instance for the [MTRUnitTestingClusterNestedStructList] class.
var (
	MTRUnitTestingClusterNestedStructListClass     _MTRUnitTestingClusterNestedStructListClass
	MTRUnitTestingClusterNestedStructListClassOnce sync.Once
)

func getMTRUnitTestingClusterNestedStructListClass() _MTRUnitTestingClusterNestedStructListClass {
	MTRUnitTestingClusterNestedStructListClassOnce.Do(func() {
		MTRUnitTestingClusterNestedStructListClass = _MTRUnitTestingClusterNestedStructListClass{objc.GetClass("MTRUnitTestingClusterNestedStructList")}
	})
	return MTRUnitTestingClusterNestedStructListClass
}

type _MTRUnitTestingClusterNestedStructListClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterNestedStructList */
// An interface definition for the [MTRUnitTestingClusterNestedStructList] class.
type IMTRUnitTestingClusterNestedStructList interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterNestedStructList */
	// properties:
	A() objc.IObject /* cross-framework: NSNumber */
	SetA(value objc.IObject /* cross-framework: NSNumber */)
	B() objc.IObject /* cross-framework: NSNumber */
	SetB(value objc.IObject /* cross-framework: NSNumber */)
	C() IMTRUnitTestingClusterSimpleStruct
	SetC(value IMTRUnitTestingClusterSimpleStruct)
	D() objc.IObject /* cross-framework: NSArray */
	SetD(value objc.IObject /* cross-framework: NSArray */)
	E() objc.IObject /* cross-framework: NSArray */
	SetE(value objc.IObject /* cross-framework: NSArray */)
	F() objc.IObject /* cross-framework: NSArray */
	SetF(value objc.IObject /* cross-framework: NSArray */)
	G() objc.IObject /* cross-framework: NSArray */
	SetG(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterNestedStructList */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterNestedStructList */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterNestedStructListClass) Alloc() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterNestedStructListClass) New() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterNestedStructList) Init() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterNestedStructList) Autorelease() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterNestedStructList creates a new MTRUnitTestingClusterNestedStructList instance.
func NewMTRUnitTestingClusterNestedStructList() MTRUnitTestingClusterNestedStructList {
	return getMTRUnitTestingClusterNestedStructListClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterNestedStructList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList
type MTRUnitTestingClusterNestedStructList struct {
	objectivec.Object
}

// MTRUnitTestingClusterNestedStructListFrom constructs a [MTRUnitTestingClusterNestedStructList] from an unsafe.Pointer.
func MTRUnitTestingClusterNestedStructListFrom(ptr unsafe.Pointer) MTRUnitTestingClusterNestedStructList {
	return MTRUnitTestingClusterNestedStructList{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterNestedStructList *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterNestedStructList */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterNestedStructList */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterNestedStructList */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterNestedStructList */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/a
func (m_ MTRUnitTestingClusterNestedStructList) A() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("a"))
	return rv
}/* debug [instance_properties/getter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/a
func (m_ MTRUnitTestingClusterNestedStructList) SetA(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}/* debug [instance_properties/setter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/b
func (m_ MTRUnitTestingClusterNestedStructList) B() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("b"))
	return rv
}/* debug [instance_properties/getter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/b
func (m_ MTRUnitTestingClusterNestedStructList) SetB(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}/* debug [instance_properties/setter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/c
func (m_ MTRUnitTestingClusterNestedStructList) C() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("c"))
	return rv
}/* debug [instance_properties/getter]: c */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/c
func (m_ MTRUnitTestingClusterNestedStructList) SetC(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}/* debug [instance_properties/setter]: c */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/d
func (m_ MTRUnitTestingClusterNestedStructList) D() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("d"))
	return rv
}/* debug [instance_properties/getter]: d */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/d
func (m_ MTRUnitTestingClusterNestedStructList) SetD(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setD:"), value)
}/* debug [instance_properties/setter]: d */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/e
func (m_ MTRUnitTestingClusterNestedStructList) E() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("e"))
	return rv
}/* debug [instance_properties/getter]: e */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/e
func (m_ MTRUnitTestingClusterNestedStructList) SetE(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setE:"), value)
}/* debug [instance_properties/setter]: e */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/f
func (m_ MTRUnitTestingClusterNestedStructList) F() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("f"))
	return rv
}/* debug [instance_properties/getter]: f */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/f
func (m_ MTRUnitTestingClusterNestedStructList) SetF(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setF:"), value)
}/* debug [instance_properties/setter]: f */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/g
func (m_ MTRUnitTestingClusterNestedStructList) G() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("g"))
	return rv
}/* debug [instance_properties/getter]: g */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList/g
func (m_ MTRUnitTestingClusterNestedStructList) SetG(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setG:"), value)
}/* debug [instance_properties/setter]: g */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterNestedStructList */



