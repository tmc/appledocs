// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterSimpleStruct */


/* debug [class_header]: Header for MTRUnitTestingClusterSimpleStruct */
// The class instance for the [MTRUnitTestingClusterSimpleStruct] class.
var (
	MTRUnitTestingClusterSimpleStructClass     _MTRUnitTestingClusterSimpleStructClass
	MTRUnitTestingClusterSimpleStructClassOnce sync.Once
)

func getMTRUnitTestingClusterSimpleStructClass() _MTRUnitTestingClusterSimpleStructClass {
	MTRUnitTestingClusterSimpleStructClassOnce.Do(func() {
		MTRUnitTestingClusterSimpleStructClass = _MTRUnitTestingClusterSimpleStructClass{objc.GetClass("MTRUnitTestingClusterSimpleStruct")}
	})
	return MTRUnitTestingClusterSimpleStructClass
}

type _MTRUnitTestingClusterSimpleStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterSimpleStruct */
// An interface definition for the [MTRUnitTestingClusterSimpleStruct] class.
type IMTRUnitTestingClusterSimpleStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterSimpleStruct */
	// properties:
	A() objc.IObject /* cross-framework: NSNumber */
	SetA(value objc.IObject /* cross-framework: NSNumber */)
	B() objc.IObject /* cross-framework: NSNumber */
	SetB(value objc.IObject /* cross-framework: NSNumber */)
	C() objc.IObject /* cross-framework: NSNumber */
	SetC(value objc.IObject /* cross-framework: NSNumber */)
	D() objc.IObject /* cross-framework: NSData */
	SetD(value objc.IObject /* cross-framework: NSData */)
	E() objc.IObject /* cross-framework: NSString */
	SetE(value objc.IObject /* cross-framework: NSString */)
	F() objc.IObject /* cross-framework: NSNumber */
	SetF(value objc.IObject /* cross-framework: NSNumber */)
	G() objc.IObject /* cross-framework: NSNumber */
	SetG(value objc.IObject /* cross-framework: NSNumber */)
	H() objc.IObject /* cross-framework: NSNumber */
	SetH(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterSimpleStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterSimpleStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterSimpleStructClass) Alloc() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterSimpleStructClass) New() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterSimpleStruct) Init() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterSimpleStruct) Autorelease() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterSimpleStruct creates a new MTRUnitTestingClusterSimpleStruct instance.
func NewMTRUnitTestingClusterSimpleStruct() MTRUnitTestingClusterSimpleStruct {
	return getMTRUnitTestingClusterSimpleStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterSimpleStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct
type MTRUnitTestingClusterSimpleStruct struct {
	objectivec.Object
}

// MTRUnitTestingClusterSimpleStructFrom constructs a [MTRUnitTestingClusterSimpleStruct] from an unsafe.Pointer.
func MTRUnitTestingClusterSimpleStructFrom(ptr unsafe.Pointer) MTRUnitTestingClusterSimpleStruct {
	return MTRUnitTestingClusterSimpleStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterSimpleStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterSimpleStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterSimpleStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterSimpleStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterSimpleStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/a
func (m_ MTRUnitTestingClusterSimpleStruct) A() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("a"))
	return rv
}/* debug [instance_properties/getter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/a
func (m_ MTRUnitTestingClusterSimpleStruct) SetA(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}/* debug [instance_properties/setter]: a */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/b
func (m_ MTRUnitTestingClusterSimpleStruct) B() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("b"))
	return rv
}/* debug [instance_properties/getter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/b
func (m_ MTRUnitTestingClusterSimpleStruct) SetB(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}/* debug [instance_properties/setter]: b */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/c
func (m_ MTRUnitTestingClusterSimpleStruct) C() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("c"))
	return rv
}/* debug [instance_properties/getter]: c */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/c
func (m_ MTRUnitTestingClusterSimpleStruct) SetC(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}/* debug [instance_properties/setter]: c */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/d
func (m_ MTRUnitTestingClusterSimpleStruct) D() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("d"))
	return rv
}/* debug [instance_properties/getter]: d */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/d
func (m_ MTRUnitTestingClusterSimpleStruct) SetD(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setD:"), value)
}/* debug [instance_properties/setter]: d */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/e
func (m_ MTRUnitTestingClusterSimpleStruct) E() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("e"))
	return rv
}/* debug [instance_properties/getter]: e */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/e
func (m_ MTRUnitTestingClusterSimpleStruct) SetE(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setE:"), value)
}/* debug [instance_properties/setter]: e */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/f
func (m_ MTRUnitTestingClusterSimpleStruct) F() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("f"))
	return rv
}/* debug [instance_properties/getter]: f */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/f
func (m_ MTRUnitTestingClusterSimpleStruct) SetF(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setF:"), value)
}/* debug [instance_properties/setter]: f */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/g
func (m_ MTRUnitTestingClusterSimpleStruct) G() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("g"))
	return rv
}/* debug [instance_properties/getter]: g */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/g
func (m_ MTRUnitTestingClusterSimpleStruct) SetG(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setG:"), value)
}/* debug [instance_properties/setter]: g */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/h
func (m_ MTRUnitTestingClusterSimpleStruct) H() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("h"))
	return rv
}/* debug [instance_properties/getter]: h */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct/h
func (m_ MTRUnitTestingClusterSimpleStruct) SetH(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setH:"), value)
}/* debug [instance_properties/setter]: h */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterSimpleStruct */



