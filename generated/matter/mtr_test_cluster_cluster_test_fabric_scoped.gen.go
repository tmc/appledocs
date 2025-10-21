// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTestClusterClusterTestFabricScoped] class.
var (
	MTRTestClusterClusterTestFabricScopedClass     _MTRTestClusterClusterTestFabricScopedClass
	MTRTestClusterClusterTestFabricScopedClassOnce sync.Once
)

func getMTRTestClusterClusterTestFabricScopedClass() _MTRTestClusterClusterTestFabricScopedClass {
	MTRTestClusterClusterTestFabricScopedClassOnce.Do(func() {
		MTRTestClusterClusterTestFabricScopedClass = _MTRTestClusterClusterTestFabricScopedClass{objc.GetClass("MTRTestClusterClusterTestFabricScoped")}
	})
	return MTRTestClusterClusterTestFabricScopedClass
}

type _MTRTestClusterClusterTestFabricScopedClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterTestFabricScoped] class.
type IMTRTestClusterClusterTestFabricScoped interface {
	IMTRUnitTestingClusterTestFabricScoped
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterTestFabricScoped
type MTRTestClusterClusterTestFabricScoped struct {
	MTRUnitTestingClusterTestFabricScoped
}

// MTRTestClusterClusterTestFabricScopedFrom constructs a [MTRTestClusterClusterTestFabricScoped] from an unsafe.Pointer.
func MTRTestClusterClusterTestFabricScopedFrom(ptr unsafe.Pointer) MTRTestClusterClusterTestFabricScoped {
	return MTRTestClusterClusterTestFabricScoped{
		MTRUnitTestingClusterTestFabricScoped: MTRUnitTestingClusterTestFabricScopedFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterTestFabricScopedClass) Alloc() MTRTestClusterClusterTestFabricScoped {
	rv := objc.Send[MTRTestClusterClusterTestFabricScoped](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterTestFabricScopedClass) New() MTRTestClusterClusterTestFabricScoped {
	rv := objc.Send[MTRTestClusterClusterTestFabricScoped](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterTestFabricScoped) Init() MTRTestClusterClusterTestFabricScoped {
	rv := objc.Send[MTRTestClusterClusterTestFabricScoped](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterTestFabricScoped) Autorelease() MTRTestClusterClusterTestFabricScoped {
	rv := objc.Send[MTRTestClusterClusterTestFabricScoped](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterTestFabricScoped creates a new MTRTestClusterClusterTestFabricScoped instance.
func NewMTRTestClusterClusterTestFabricScoped() MTRTestClusterClusterTestFabricScoped {
	return getMTRTestClusterClusterTestFabricScopedClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/fabricindex
func (m_ MTRTestClusterClusterTestFabricScoped) FabricIndex() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricIndex"))
	return rv
}


// SetFabricIndex sets the value of the fabricIndex property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/fabricindex
func (m_ MTRTestClusterClusterTestFabricScoped) SetFabricIndex(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricIndex:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/fabricsensitivecharstring
func (m_ MTRTestClusterClusterTestFabricScoped) FabricSensitiveCharString() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("fabricSensitiveCharString"))
	return rv
}


// SetFabricSensitiveCharString sets the value of the fabricSensitiveCharString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/fabricsensitivecharstring
func (m_ MTRTestClusterClusterTestFabricScoped) SetFabricSensitiveCharString(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveCharString:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/fabricsensitiveint8u
func (m_ MTRTestClusterClusterTestFabricScoped) FabricSensitiveInt8u() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("fabricSensitiveInt8u"))
	return rv
}


// SetFabricSensitiveInt8u sets the value of the fabricSensitiveInt8u property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/fabricsensitiveint8u
func (m_ MTRTestClusterClusterTestFabricScoped) SetFabricSensitiveInt8u(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveInt8u:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/fabricsensitiveint8ulist
func (m_ MTRTestClusterClusterTestFabricScoped) FabricSensitiveInt8uList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("fabricSensitiveInt8uList"))
	return rv
}


// SetFabricSensitiveInt8uList sets the value of the fabricSensitiveInt8uList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/fabricsensitiveint8ulist
func (m_ MTRTestClusterClusterTestFabricScoped) SetFabricSensitiveInt8uList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveInt8uList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/fabricsensitivestruct
func (m_ MTRTestClusterClusterTestFabricScoped) FabricSensitiveStruct() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("fabricSensitiveStruct"))
	return rv
}


// SetFabricSensitiveStruct sets the value of the fabricSensitiveStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/fabricsensitivestruct
func (m_ MTRTestClusterClusterTestFabricScoped) SetFabricSensitiveStruct(value IMTRTestClusterClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setFabricSensitiveStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/nullablefabricsensitiveint8u
func (m_ MTRTestClusterClusterTestFabricScoped) NullableFabricSensitiveInt8u() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableFabricSensitiveInt8u"))
	return rv
}


// SetNullableFabricSensitiveInt8u sets the value of the nullableFabricSensitiveInt8u property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/nullablefabricsensitiveint8u
func (m_ MTRTestClusterClusterTestFabricScoped) SetNullableFabricSensitiveInt8u(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableFabricSensitiveInt8u:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/nullableoptionalfabricsensitiveint8u
func (m_ MTRTestClusterClusterTestFabricScoped) NullableOptionalFabricSensitiveInt8u() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalFabricSensitiveInt8u"))
	return rv
}


// SetNullableOptionalFabricSensitiveInt8u sets the value of the nullableOptionalFabricSensitiveInt8u property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/nullableoptionalfabricsensitiveint8u
func (m_ MTRTestClusterClusterTestFabricScoped) SetNullableOptionalFabricSensitiveInt8u(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalFabricSensitiveInt8u:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/optionalfabricsensitiveint8u
func (m_ MTRTestClusterClusterTestFabricScoped) OptionalFabricSensitiveInt8u() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalFabricSensitiveInt8u"))
	return rv
}


// SetOptionalFabricSensitiveInt8u sets the value of the optionalFabricSensitiveInt8u property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustertestfabricscoped/optionalfabricsensitiveint8u
func (m_ MTRTestClusterClusterTestFabricScoped) SetOptionalFabricSensitiveInt8u(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalFabricSensitiveInt8u:"), value)
}



