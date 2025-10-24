// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRUnitTestingClusterTestFabricScoped */


/* debug [class_header]: Header for MTRUnitTestingClusterTestFabricScoped */
// The class instance for the [MTRUnitTestingClusterTestFabricScoped] class.
var (
	MTRUnitTestingClusterTestFabricScopedClass     _MTRUnitTestingClusterTestFabricScopedClass
	MTRUnitTestingClusterTestFabricScopedClassOnce sync.Once
)

func getMTRUnitTestingClusterTestFabricScopedClass() _MTRUnitTestingClusterTestFabricScopedClass {
	MTRUnitTestingClusterTestFabricScopedClassOnce.Do(func() {
		MTRUnitTestingClusterTestFabricScopedClass = _MTRUnitTestingClusterTestFabricScopedClass{objc.GetClass("MTRUnitTestingClusterTestFabricScoped")}
	})
	return MTRUnitTestingClusterTestFabricScopedClass
}

type _MTRUnitTestingClusterTestFabricScopedClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRUnitTestingClusterTestFabricScoped */
// An interface definition for the [MTRUnitTestingClusterTestFabricScoped] class.
type IMTRUnitTestingClusterTestFabricScoped interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRUnitTestingClusterTestFabricScoped */
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	FabricSensitiveCharString() objc.IObject /* cross-framework: NSString */
	SetFabricSensitiveCharString(value objc.IObject /* cross-framework: NSString */)
	FabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */
	SetFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */)
	FabricSensitiveInt8uList() objc.IObject /* cross-framework: NSArray */
	SetFabricSensitiveInt8uList(value objc.IObject /* cross-framework: NSArray */)
	FabricSensitiveStruct() IMTRUnitTestingClusterSimpleStruct
	SetFabricSensitiveStruct(value IMTRUnitTestingClusterSimpleStruct)
	NullableFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */
	SetNullableFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */)
	OptionalFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */
	SetOptionalFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRUnitTestingClusterTestFabricScoped */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRUnitTestingClusterTestFabricScoped */
// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestFabricScopedClass) Alloc() MTRUnitTestingClusterTestFabricScoped {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScoped](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRUnitTestingClusterTestFabricScopedClass) New() MTRUnitTestingClusterTestFabricScoped {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScoped](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterTestFabricScoped) Init() MTRUnitTestingClusterTestFabricScoped {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScoped](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterTestFabricScoped) Autorelease() MTRUnitTestingClusterTestFabricScoped {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScoped](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterTestFabricScoped creates a new MTRUnitTestingClusterTestFabricScoped instance.
func NewMTRUnitTestingClusterTestFabricScoped() MTRUnitTestingClusterTestFabricScoped {
	return getMTRUnitTestingClusterTestFabricScopedClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRUnitTestingClusterTestFabricScoped */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped
type MTRUnitTestingClusterTestFabricScoped struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestFabricScopedFrom constructs a [MTRUnitTestingClusterTestFabricScoped] from an unsafe.Pointer.
func MTRUnitTestingClusterTestFabricScopedFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestFabricScoped {
	return MTRUnitTestingClusterTestFabricScoped{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRUnitTestingClusterTestFabricScoped *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRUnitTestingClusterTestFabricScoped */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRUnitTestingClusterTestFabricScoped */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRUnitTestingClusterTestFabricScoped */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRUnitTestingClusterTestFabricScoped */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/fabricIndex
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}/* debug [instance_properties/getter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/fabricIndex
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}/* debug [instance_properties/setter]: fabricIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/fabricSensitiveCharString
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveCharString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("fabricSensitiveCharString"))
	return rv
}/* debug [instance_properties/getter]: fabricSensitiveCharString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/fabricSensitiveCharString
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveCharString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveCharString:"), value)
}/* debug [instance_properties/setter]: fabricSensitiveCharString */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/fabricSensitiveInt8u
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricSensitiveInt8u"))
	return rv
}/* debug [instance_properties/getter]: fabricSensitiveInt8u */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/fabricSensitiveInt8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveInt8u:"), value)
}/* debug [instance_properties/setter]: fabricSensitiveInt8u */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/fabricSensitiveInt8uList
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveInt8uList() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("fabricSensitiveInt8uList"))
	return rv
}/* debug [instance_properties/getter]: fabricSensitiveInt8uList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/fabricSensitiveInt8uList
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveInt8uList(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveInt8uList:"), value)
}/* debug [instance_properties/setter]: fabricSensitiveInt8uList */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/fabricSensitiveStruct
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("fabricSensitiveStruct"))
	return rv
}/* debug [instance_properties/getter]: fabricSensitiveStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/fabricSensitiveStruct
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveStruct:"), value)
}/* debug [instance_properties/setter]: fabricSensitiveStruct */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/nullableFabricSensitiveInt8u
func (m_ MTRUnitTestingClusterTestFabricScoped) NullableFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableFabricSensitiveInt8u"))
	return rv
}/* debug [instance_properties/getter]: nullableFabricSensitiveInt8u */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/nullableFabricSensitiveInt8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetNullableFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableFabricSensitiveInt8u:"), value)
}/* debug [instance_properties/setter]: nullableFabricSensitiveInt8u */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/nullableOptionalFabricSensitiveInt8u
func (m_ MTRUnitTestingClusterTestFabricScoped) NullableOptionalFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalFabricSensitiveInt8u"))
	return rv
}/* debug [instance_properties/getter]: nullableOptionalFabricSensitiveInt8u */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/nullableOptionalFabricSensitiveInt8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetNullableOptionalFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalFabricSensitiveInt8u:"), value)
}/* debug [instance_properties/setter]: nullableOptionalFabricSensitiveInt8u */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/optionalFabricSensitiveInt8u
func (m_ MTRUnitTestingClusterTestFabricScoped) OptionalFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalFabricSensitiveInt8u"))
	return rv
}/* debug [instance_properties/getter]: optionalFabricSensitiveInt8u */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped/optionalFabricSensitiveInt8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetOptionalFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalFabricSensitiveInt8u:"), value)
}/* debug [instance_properties/setter]: optionalFabricSensitiveInt8u */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRUnitTestingClusterTestFabricScoped */



