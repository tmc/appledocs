// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRUnitTestingClusterTestFabricScoped] class.
type IMTRUnitTestingClusterTestFabricScoped interface {
	objectivec.IObject
	// properties:
	FabricIndex() objc.IObject /* cross-framework: NSNumber */
	SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */)
	FabricSensitiveCharString() objc.IObject /* cross-framework: NSString */
	SetFabricSensitiveCharString(value objc.IObject /* cross-framework: NSString */)
	FabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */
	SetFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */)
	FabricSensitiveInt8uList() unsafe.Pointer
	SetFabricSensitiveInt8uList(value unsafe.Pointer)
	FabricSensitiveStruct() IMTRUnitTestingClusterSimpleStruct
	SetFabricSensitiveStruct(value IMTRUnitTestingClusterSimpleStruct)
	NullableFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */
	SetNullableFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */)
	OptionalFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */
	SetOptionalFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterTestFabricScoped
type MTRUnitTestingClusterTestFabricScoped struct {
	objectivec.Object
}

// MTRUnitTestingClusterTestFabricScopedFrom constructs a [MTRUnitTestingClusterTestFabricScoped] from an unsafe.Pointer.
func MTRUnitTestingClusterTestFabricScopedFrom(ptr unsafe.Pointer) MTRUnitTestingClusterTestFabricScoped {
	return MTRUnitTestingClusterTestFabricScoped{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterTestFabricScopedClass) Alloc() MTRUnitTestingClusterTestFabricScoped {
	rv := objc.Send[MTRUnitTestingClusterTestFabricScoped](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricindex
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricindex
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitivecharstring
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveCharString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("fabricSensitiveCharString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitivecharstring
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveCharString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveCharString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("fabricSensitiveInt8u"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveInt8u:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitiveint8ulist
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveInt8uList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fabricSensitiveInt8uList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitiveint8ulist
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveInt8uList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveInt8uList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitivestruct
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("fabricSensitiveStruct"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitivestruct
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveStruct:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/nullablefabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) NullableFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableFabricSensitiveInt8u"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/nullablefabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetNullableFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableFabricSensitiveInt8u:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/nullableoptionalfabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) NullableOptionalFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalFabricSensitiveInt8u"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/nullableoptionalfabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetNullableOptionalFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalFabricSensitiveInt8u:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/optionalfabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) OptionalFabricSensitiveInt8u() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalFabricSensitiveInt8u"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/optionalfabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetOptionalFabricSensitiveInt8u(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalFabricSensitiveInt8u:"), value)
}



