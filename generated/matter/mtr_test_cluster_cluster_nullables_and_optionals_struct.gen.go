// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionallist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableOptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalList"))
	return rv
}


// SetNullableOptionalList sets the value of the nullableOptionalList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionallist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionallist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) OptionalList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalList"))
	return rv
}


// SetOptionalList sets the value of the optionalList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionallist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetOptionalList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableInt"))
	return rv
}


// SetNullableInt sets the value of the nullableInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableInt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) OptionalInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("optionalInt"))
	return rv
}


// SetOptionalInt sets the value of the optionalInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetOptionalInt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableOptionalInt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("nullableOptionalInt"))
	return rv
}


// SetNullableOptionalInt sets the value of the nullableOptionalInt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalint
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableOptionalInt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalInt:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablestring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("nullableString"))
	return rv
}


// SetNullableString sets the value of the nullableString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablestring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalstruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableOptionalStruct() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableOptionalStruct"))
	return rv
}


// SetNullableOptionalStruct sets the value of the nullableOptionalStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalstruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableOptionalStruct(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalstruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) OptionalStruct() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("optionalStruct"))
	return rv
}


// SetOptionalStruct sets the value of the optionalStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalstruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetOptionalStruct(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalstring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableOptionalString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("nullableOptionalString"))
	return rv
}


// SetNullableOptionalString sets the value of the nullableOptionalString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullableoptionalstring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableOptionalString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableOptionalString:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablestruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableStruct() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableStruct"))
	return rv
}


// SetNullableStruct sets the value of the nullableStruct property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablestruct
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableStruct(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableStruct:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablelist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) NullableList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("nullableList"))
	return rv
}


// SetNullableList sets the value of the nullableList property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/nullablelist
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetNullableList(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNullableList:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalstring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) OptionalString() string {
	rv := objc.Send[string](m_.ID, objc.Sel("optionalString"))
	return rv
}


// SetOptionalString sets the value of the optionalString property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternullablesandoptionalsstruct/optionalstring
func (m_ MTRTestClusterClusterNullablesAndOptionalsStruct) SetOptionalString(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOptionalString:"), objc.String(value))
}



