// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterNestedStructList] class.
var (
	MTRUnitTestingClusterNestedStructListClass     _MTRUnitTestingClusterNestedStructListClass
	MTRUnitTestingClusterNestedStructListClassOnce sync.Once
)

func getMTRUnitTestingClusterNestedStructListClass() _MTRUnitTestingClusterNestedStructListClass {
	MTRUnitTestingClusterNestedStructListClassOnce.Do(func() {
		MTRUnitTestingClusterNestedStructListClass = _MTRUnitTestingClusterNestedStructListClass{objc.GetClass("MTRUnitTestingClusterNestedStructList")}
	})
	return MTRUnitTestingClusterNestedStructListClass
}

type _MTRUnitTestingClusterNestedStructListClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterNestedStructList] class.
type IMTRUnitTestingClusterNestedStructList interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterNestedStructList
type MTRUnitTestingClusterNestedStructList struct {
	objectivec.Object
}

// MTRUnitTestingClusterNestedStructListFrom constructs a [MTRUnitTestingClusterNestedStructList] from an unsafe.Pointer.
func MTRUnitTestingClusterNestedStructListFrom(ptr unsafe.Pointer) MTRUnitTestingClusterNestedStructList {
	return MTRUnitTestingClusterNestedStructList{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterNestedStructListClass) Alloc() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterNestedStructListClass) New() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterNestedStructList) Init() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterNestedStructList) Autorelease() MTRUnitTestingClusterNestedStructList {
	rv := objc.Send[MTRUnitTestingClusterNestedStructList](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterNestedStructList creates a new MTRUnitTestingClusterNestedStructList instance.
func NewMTRUnitTestingClusterNestedStructList() MTRUnitTestingClusterNestedStructList {
	return getMTRUnitTestingClusterNestedStructListClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/a
func (m_ MTRUnitTestingClusterNestedStructList) A() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("a"))
	return rv
}


// SetA sets the value of the a property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/a
func (m_ MTRUnitTestingClusterNestedStructList) SetA(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/b
func (m_ MTRUnitTestingClusterNestedStructList) B() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("b"))
	return rv
}


// SetB sets the value of the b property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/b
func (m_ MTRUnitTestingClusterNestedStructList) SetB(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/c
func (m_ MTRUnitTestingClusterNestedStructList) C() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("c"))
	return rv
}


// SetC sets the value of the c property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/c
func (m_ MTRUnitTestingClusterNestedStructList) SetC(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/d
func (m_ MTRUnitTestingClusterNestedStructList) D() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("d"))
	return rv
}


// SetD sets the value of the d property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/d
func (m_ MTRUnitTestingClusterNestedStructList) SetD(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setD:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/e
func (m_ MTRUnitTestingClusterNestedStructList) E() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("e"))
	return rv
}


// SetE sets the value of the e property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/e
func (m_ MTRUnitTestingClusterNestedStructList) SetE(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setE:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/f
func (m_ MTRUnitTestingClusterNestedStructList) F() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("f"))
	return rv
}


// SetF sets the value of the f property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/f
func (m_ MTRUnitTestingClusterNestedStructList) SetF(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setF:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/g
func (m_ MTRUnitTestingClusterNestedStructList) G() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("g"))
	return rv
}


// SetG sets the value of the g property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/g
func (m_ MTRUnitTestingClusterNestedStructList) SetG(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setG:"), value)
}



