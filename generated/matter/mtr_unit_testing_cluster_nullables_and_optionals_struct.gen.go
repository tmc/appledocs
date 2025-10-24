// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterNullablesAndOptionalsStruct */


/* debug [class_header]: Header for MTRUnitTestingClusterNullablesAndOptionalsStruct */
// The class instance for the [MTRUnitTestingClusterNullablesAndOptionalsStruct] class.
var (
	MTRUnitTestingClusterNullablesAndOptionalsStructClass     _MTRUnitTestingClusterNullablesAndOptionalsStructClass
	MTRUnitTestingClusterNullablesAndOptionalsStructClassOnce sync.Once
)

func getMTRUnitTestingClusterNullablesAndOptionalsStructClass() _MTRUnitTestingClusterNullablesAndOptionalsStructClass {
	MTRUnitTestingClusterNullablesAndOptionalsStructClassOnce.Do(func() {
		MTRUnitTestingClusterNullablesAndOptionalsStructClass = _MTRUnitTestingClusterNullablesAndOptionalsStructClass{objc.GetClass("MTRUnitTestingClusterNullablesAndOptionalsStruct")}
	})
	return MTRUnitTestingClusterNullablesAndOptionalsStructClass
}

type _MTRUnitTestingClusterNullablesAndOptionalsStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterNullablesAndOptionalsStruct */
// An interface definition for the [MTRUnitTestingClusterNullablesAndOptionalsStruct] class.
type IMTRUnitTestingClusterNullablesAndOptionalsStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterNullablesAndOptionalsStruct */
	// properties:
	NullableInt() objc.IObject /* cross-framework: NSNumber */
	SetNullableInt(value objc.IObject /* cross-framework: NSNumber */)
	NullableList() objc.IObject /* cross-framework: NSArray */
	SetNullableList(value objc.IObject /* cross-framework: NSArray */)
	NullableOptionalInt() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalInt(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalList() objc.IObject /* cross-framework: NSArray */
	SetNullableOptionalList(value objc.IObject /* cross-framework: NSArray */)
	NullableOptionalString() objc.IObject /* cross-framework: NSString */
	SetNullableOptionalString(value objc.IObject /* cross-framework: NSString */)
	NullableOptionalStruct() IMTRUnitTestingClusterSimpleStruct
	SetNullableOptionalStruct(value IMTRUnitTestingClusterSimpleStruct)
	NullableString() objc.IObject /* cross-framework: NSString */
	SetNullableString(value objc.IObject /* cross-framework: NSString */)
	NullableStruct() IMTRUnitTestingClusterSimpleStruct
	SetNullableStruct(value IMTRUnitTestingClusterSimpleStruct)
	OptionalInt() objc.IObject /* cross-framework: NSNumber */
	SetOptionalInt(value objc.IObject /* cross-framework: NSNumber */)
	OptionalList() objc.IObject /* cross-framework: NSArray */
	SetOptionalList(value objc.IObject /* cross-framework: NSArray */)
	OptionalString() objc.IObject /* cross-framework: NSString */
	SetOptionalString(value objc.IObject /* cross-framework: NSString */)
	OptionalStruct() IMTRUnitTestingClusterSimpleStruct
	SetOptionalStruct(value IMTRUnitTestingClusterSimpleStruct)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterNullablesAndOptionalsStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterNullablesAndOptionalsStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterNullablesAndOptionalsStructClass) Alloc() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRUnitTestingClusterNullablesAndOptionalsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterNullablesAndOptionalsStructClass) New() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRUnitTestingClusterNullablesAndOptionalsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) Init() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRUnitTestingClusterNullablesAndOptionalsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) Autorelease() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRUnitTestingClusterNullablesAndOptionalsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterNullablesAndOptionalsStruct creates a new MTRUnitTestingClusterNullablesAndOptionalsStruct instance.
func NewMTRUnitTestingClusterNullablesAndOptionalsStruct() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	return getMTRUnitTestingClusterNullablesAndOptionalsStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterNullablesAndOptionalsStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct
type MTRUnitTestingClusterNullablesAndOptionalsStruct struct {
	objectivec.Object
}

// MTRUnitTestingClusterNullablesAndOptionalsStructFrom constructs a [MTRUnitTestingClusterNullablesAndOptionalsStruct] from an unsafe.Pointer.
func MTRUnitTestingClusterNullablesAndOptionalsStructFrom(ptr unsafe.Pointer) MTRUnitTestingClusterNullablesAndOptionalsStruct {
	return MTRUnitTestingClusterNullablesAndOptionalsStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterNullablesAndOptionalsStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterNullablesAndOptionalsStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterNullablesAndOptionalsStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterNullablesAndOptionalsStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterNullablesAndOptionalsStruct */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableInt
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableInt"))
	return rv
}/* debug [instance_properties/getter]: nullableInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableInt
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableInt:"), value)
}/* debug [instance_properties/setter]: nullableInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableList
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("nullableList"))
	return rv
}/* debug [instance_properties/getter]: nullableList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableList
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableList:"), value)
}/* debug [instance_properties/setter]: nullableList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableOptionalInt
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalInt"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableOptionalInt
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalInt:"), value)
}/* debug [instance_properties/setter]: nullableOptionalInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableOptionalList
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("nullableOptionalList"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableOptionalList
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalList:"), value)
}/* debug [instance_properties/setter]: nullableOptionalList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableOptionalString
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableOptionalString"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableOptionalString
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalString:"), value)
}/* debug [instance_properties/setter]: nullableOptionalString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableOptionalStruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStruct"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableOptionalStruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStruct:"), value)
}/* debug [instance_properties/setter]: nullableOptionalStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableString
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableString"))
	return rv
}/* debug [instance_properties/getter]: nullableString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableString
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableString:"), value)
}/* debug [instance_properties/setter]: nullableString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableStruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStruct"))
	return rv
}/* debug [instance_properties/getter]: nullableStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/nullableStruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStruct:"), value)
}/* debug [instance_properties/setter]: nullableStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/optionalInt
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalInt"))
	return rv
}/* debug [instance_properties/getter]: optionalInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/optionalInt
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalInt:"), value)
}/* debug [instance_properties/setter]: optionalInt */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/optionalList
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("optionalList"))
	return rv
}/* debug [instance_properties/getter]: optionalList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/optionalList
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalList:"), value)
}/* debug [instance_properties/setter]: optionalList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/optionalString
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("optionalString"))
	return rv
}/* debug [instance_properties/getter]: optionalString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/optionalString
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalString:"), value)
}/* debug [instance_properties/setter]: optionalString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/optionalStruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStruct"))
	return rv
}/* debug [instance_properties/getter]: optionalStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct/optionalStruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStruct:"), value)
}/* debug [instance_properties/setter]: optionalStruct */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterNullablesAndOptionalsStruct */



