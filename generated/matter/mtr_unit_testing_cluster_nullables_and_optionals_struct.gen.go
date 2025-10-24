// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MTRUnitTestingClusterNullablesAndOptionalsStruct] class.
type IMTRUnitTestingClusterNullablesAndOptionalsStruct interface {
	objectivec.IObject
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
	NullableOptionalStruct() IMTRUnitTestingClusterSimpleStruct
	SetNullableOptionalStruct(value IMTRUnitTestingClusterSimpleStruct)
	NullableString() objc.IObject /* cross-framework: NSString */
	SetNullableString(value objc.IObject /* cross-framework: NSString */)
	NullableStruct() IMTRUnitTestingClusterSimpleStruct
	SetNullableStruct(value IMTRUnitTestingClusterSimpleStruct)
	OptionalInt() objc.IObject /* cross-framework: NSNumber */
	SetOptionalInt(value objc.IObject /* cross-framework: NSNumber */)
	OptionalList() unsafe.Pointer
	SetOptionalList(value unsafe.Pointer)
	OptionalString() objc.IObject /* cross-framework: NSString */
	SetOptionalString(value objc.IObject /* cross-framework: NSString */)
	OptionalStruct() IMTRUnitTestingClusterSimpleStruct
	SetOptionalStruct(value IMTRUnitTestingClusterSimpleStruct)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNullablesAndOptionalsStruct
type MTRUnitTestingClusterNullablesAndOptionalsStruct struct {
	objectivec.Object
}

// MTRUnitTestingClusterNullablesAndOptionalsStructFrom constructs a [MTRUnitTestingClusterNullablesAndOptionalsStruct] from an unsafe.Pointer.
func MTRUnitTestingClusterNullablesAndOptionalsStructFrom(ptr unsafe.Pointer) MTRUnitTestingClusterNullablesAndOptionalsStruct {
	return MTRUnitTestingClusterNullablesAndOptionalsStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterNullablesAndOptionalsStructClass) Alloc() MTRUnitTestingClusterNullablesAndOptionalsStruct {
	rv := objc.Send[MTRUnitTestingClusterNullablesAndOptionalsStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableInt"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableInt:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablelist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablelist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("nullableOptionalInt"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalInt:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionallist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionallist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalstring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableOptionalString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalstring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalstruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableOptionalStruct"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalstruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStruct:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablestring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("nullableString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablestring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablestruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("nullableStruct"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablestruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStruct:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalInt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("optionalInt"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalInt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalInt:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionallist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalList"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionallist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalList:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalstring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("optionalString"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalstring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalString(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalString:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalstruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalStruct() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("optionalStruct"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalstruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalStruct(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStruct:"), value)
}



