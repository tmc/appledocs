// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterNullablesAndOptionalsStruct] class.
var (
	MTRTestClusterClusterNullablesAndOptionalsStructClass     _MTRTestClusterClusterNullablesAndOptionalsStructClass
	MTRTestClusterClusterNullablesAndOptionalsStructClassOnce sync.Once
)

func getMTRTestClusterClusterNullablesAndOptionalsStructClass() _MTRTestClusterClusterNullablesAndOptionalsStructClass {
	MTRTestClusterClusterNullablesAndOptionalsStructClassOnce.Do(func() {
		MTRTestClusterClusterNullablesAndOptionalsStructClass = _MTRTestClusterClusterNullablesAndOptionalsStructClass{objc.GetClass("MTRTestClusterClusterNullablesAndOptionalsStruct")}
	})
	return MTRTestClusterClusterNullablesAndOptionalsStructClass
}

type _MTRTestClusterClusterNullablesAndOptionalsStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterNullablesAndOptionalsStruct] class.
type IMTRTestClusterClusterNullablesAndOptionalsStruct interface {
	IMTRUnitTestingClusterNullablesAndOptionalsStruct
	// properties:
	NullableInt() objc.IObject /* cross-framework: NSNumber */
	SetNullableInt(value objc.IObject /* cross-framework: NSNumber */)
	NullableList() unsafe.Pointer
	SetNullableList(value unsafe.Pointer)
	NullableOptionalInt() objc.IObject /* cross-framework: NSNumber */
	SetNullableOptionalInt(value objc.IObject /* cross-framework: NSNumber */)
	NullableOptionalList() unsafe.Pointer
	SetNullableOptionalList(value unsafe.Pointer)
	NullableOptionalString() objc.IObject /* cross-framework: NSString */
	SetNullableOptionalString(value objc.IObject /* cross-framework: NSString */)
	NullableOptionalStruct() IMTRTestClusterClusterSimpleStruct
	SetNullableOptionalStruct(value IMTRTestClusterClusterSimpleStruct)
	NullableString() objc.IObject /* cross-framework: NSString */
	SetNullableString(value objc.IObject /* cross-framework: NSString */)
	NullableStruct() IMTRTestClusterClusterSimpleStruct
	SetNullableStruct(value IMTRTestClusterClusterSimpleStruct)
	OptionalInt() objc.IObject /* cross-framework: NSNumber */
	SetOptionalInt(value objc.IObject /* cross-framework: NSNumber */)
	OptionalList() unsafe.Pointer
	SetOptionalList(value unsafe.Pointer)
	OptionalString() objc.IObject /* cross-framework: NSString */
	SetOptionalString(value objc.IObject /* cross-framework: NSString */)
	OptionalStruct() IMTRTestClusterClusterSimpleStruct
	SetOptionalStruct(value IMTRTestClusterClusterSimpleStruct)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterNullablesAndOptionalsStruct
type MTRTestClusterClusterNullablesAndOptionalsStruct struct {
	MTRUnitTestingClusterNullablesAndOptionalsStruct
}

// MTRTestClusterClusterNullablesAndOptionalsStructFrom constructs a [MTRTestClusterClusterNullablesAndOptionalsStruct] from an unsafe.Pointer.
func MTRTestClusterClusterNullablesAndOptionalsStructFrom(ptr unsafe.Pointer) MTRTestClusterClusterNullablesAndOptionalsStruct {
	return MTRTestClusterClusterNullablesAndOptionalsStruct{
		MTRUnitTestingClusterNullablesAndOptionalsStruct: MTRUnitTestingClusterNullablesAndOptionalsStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterNullablesAndOptionalsStructClass) Alloc() MTRTestClusterClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRTestClusterClusterNullablesAndOptionalsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterNullablesAndOptionalsStructClass) New() MTRTestClusterClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRTestClusterClusterNullablesAndOptionalsStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) Init() MTRTestClusterClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRTestClusterClusterNullablesAndOptionalsStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) Autorelease() MTRTestClusterClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRTestClusterClusterNullablesAndOptionalsStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterNullablesAndOptionalsStruct creates a new MTRTestClusterClusterNullablesAndOptionalsStruct instance.
func NewMTRTestClusterClusterNullablesAndOptionalsStruct() MTRTestClusterClusterNullablesAndOptionalsStruct {
	return getMTRTestClusterClusterNullablesAndOptionalsStructClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableInt"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableInt:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablelist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableList"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablelist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableList:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableOptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalInt"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalInt:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionallist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableOptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalList"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionallist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalList:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalstring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableOptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableOptionalString"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalstring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalString:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalstruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableOptionalStruct() IMTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStruct"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalstruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableOptionalStruct(value IMTRTestClusterClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStruct:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablestring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableString"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablestring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableString:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablestruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableStruct() IMTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("nullableStruct"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablestruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableStruct(value IMTRTestClusterClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStruct:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) OptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalInt"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalInt:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionallist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) OptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalList"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionallist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalList:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalstring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) OptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("optionalString"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalstring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalString:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalstruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) OptionalStruct() IMTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("optionalStruct"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalstruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetOptionalStruct(value IMTRTestClusterClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStruct:"), value)
}
