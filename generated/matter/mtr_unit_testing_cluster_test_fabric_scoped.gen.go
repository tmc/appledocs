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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricindex
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricindex
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitivecharstring
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveCharString() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("fabricSensitiveCharString"))
	return rv
}


// SetFabricSensitiveCharString sets the value of the fabricSensitiveCharString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitivecharstring
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveCharString(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveCharString:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveInt8u() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricSensitiveInt8u"))
	return rv
}


// SetFabricSensitiveInt8u sets the value of the fabricSensitiveInt8u property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveInt8u(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveInt8u:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitiveint8ulist
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveInt8uList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fabricSensitiveInt8uList"))
	return rv
}


// SetFabricSensitiveInt8uList sets the value of the fabricSensitiveInt8uList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitiveint8ulist
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveInt8uList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveInt8uList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitivestruct
func (m_ MTRUnitTestingClusterTestFabricScoped) FabricSensitiveStruct() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("fabricSensitiveStruct"))
	return rv
}


// SetFabricSensitiveStruct sets the value of the fabricSensitiveStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/fabricsensitivestruct
func (m_ MTRUnitTestingClusterTestFabricScoped) SetFabricSensitiveStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/nullablefabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) NullableFabricSensitiveInt8u() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableFabricSensitiveInt8u"))
	return rv
}


// SetNullableFabricSensitiveInt8u sets the value of the nullableFabricSensitiveInt8u property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/nullablefabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetNullableFabricSensitiveInt8u(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableFabricSensitiveInt8u:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/nullableoptionalfabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) NullableOptionalFabricSensitiveInt8u() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalFabricSensitiveInt8u"))
	return rv
}


// SetNullableOptionalFabricSensitiveInt8u sets the value of the nullableOptionalFabricSensitiveInt8u property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/nullableoptionalfabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetNullableOptionalFabricSensitiveInt8u(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalFabricSensitiveInt8u:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/optionalfabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) OptionalFabricSensitiveInt8u() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalFabricSensitiveInt8u"))
	return rv
}


// SetOptionalFabricSensitiveInt8u sets the value of the optionalFabricSensitiveInt8u property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustertestfabricscoped/optionalfabricsensitiveint8u
func (m_ MTRUnitTestingClusterTestFabricScoped) SetOptionalFabricSensitiveInt8u(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalFabricSensitiveInt8u:"), value)
}



