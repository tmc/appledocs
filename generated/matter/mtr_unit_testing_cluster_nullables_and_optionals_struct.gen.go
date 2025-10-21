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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionallist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalList"))
	return rv
}


// SetOptionalList sets the value of the optionalList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionallist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableInt"))
	return rv
}


// SetNullableInt sets the value of the nullableInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableInt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionallist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalList"))
	return rv
}


// SetNullableOptionalList sets the value of the nullableOptionalList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionallist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablelist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableList"))
	return rv
}


// SetNullableList sets the value of the nullableList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablelist
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalInt"))
	return rv
}


// SetOptionalInt sets the value of the optionalInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalInt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalstruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalStruct() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalStruct"))
	return rv
}


// SetNullableOptionalStruct sets the value of the nullableOptionalStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalstruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalStruct(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablestruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableStruct() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableStruct"))
	return rv
}


// SetNullableStruct sets the value of the nullableStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablestruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableStruct(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablestring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("nullableString"))
	return rv
}


// SetNullableString sets the value of the nullableString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullablestring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalstring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("nullableOptionalString"))
	return rv
}


// SetNullableOptionalString sets the value of the nullableOptionalString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalstring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalstruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalStruct() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalStruct"))
	return rv
}


// SetOptionalStruct sets the value of the optionalStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalstruct
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalStruct(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) NullableOptionalInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalInt"))
	return rv
}


// SetNullableOptionalInt sets the value of the nullableOptionalInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/nullableoptionalint
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetNullableOptionalInt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalstring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) OptionalString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("optionalString"))
	return rv
}


// SetOptionalString sets the value of the optionalString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternullablesandoptionalsstruct/optionalstring
func (m_ MTRUnitTestingClusterNullablesAndOptionalsStruct) SetOptionalString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalString:"), objc.String(value))
}



