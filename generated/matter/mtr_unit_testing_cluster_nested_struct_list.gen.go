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
	// properties:
	A() objc.IObject /* cross-framework: NSNumber */
	SetA(value objc.IObject /* cross-framework: NSNumber */)
	B() objc.IObject /* cross-framework: NSNumber */
	SetB(value objc.IObject /* cross-framework: NSNumber */)
	C() IMTRUnitTestingClusterSimpleStruct
	SetC(value IMTRUnitTestingClusterSimpleStruct)
	D() unsafe.Pointer
	SetD(value unsafe.Pointer)
	E() unsafe.Pointer
	SetE(value unsafe.Pointer)
	F() unsafe.Pointer
	SetF(value unsafe.Pointer)
	G() unsafe.Pointer
	SetG(value unsafe.Pointer)
	// methods:
}



// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/a
func (m_ MTRUnitTestingClusterNestedStructList) A() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("a"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/a
func (m_ MTRUnitTestingClusterNestedStructList) SetA(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/b
func (m_ MTRUnitTestingClusterNestedStructList) B() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("b"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/b
func (m_ MTRUnitTestingClusterNestedStructList) SetB(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/c
func (m_ MTRUnitTestingClusterNestedStructList) C() IMTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("c"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/c
func (m_ MTRUnitTestingClusterNestedStructList) SetC(value IMTRUnitTestingClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/d
func (m_ MTRUnitTestingClusterNestedStructList) D() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("d"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/d
func (m_ MTRUnitTestingClusterNestedStructList) SetD(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setD:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/e
func (m_ MTRUnitTestingClusterNestedStructList) E() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("e"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/e
func (m_ MTRUnitTestingClusterNestedStructList) SetE(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setE:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/f
func (m_ MTRUnitTestingClusterNestedStructList) F() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("f"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/f
func (m_ MTRUnitTestingClusterNestedStructList) SetF(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setF:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/g
func (m_ MTRUnitTestingClusterNestedStructList) G() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("g"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclusternestedstructlist/g
func (m_ MTRUnitTestingClusterNestedStructList) SetG(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setG:"), value)
}



