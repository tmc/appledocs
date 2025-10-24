// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterNestedStructList] class.
var (
	MTRTestClusterClusterNestedStructListClass     _MTRTestClusterClusterNestedStructListClass
	MTRTestClusterClusterNestedStructListClassOnce sync.Once
)

func getMTRTestClusterClusterNestedStructListClass() _MTRTestClusterClusterNestedStructListClass {
	MTRTestClusterClusterNestedStructListClassOnce.Do(func() {
		MTRTestClusterClusterNestedStructListClass = _MTRTestClusterClusterNestedStructListClass{objc.GetClass("MTRTestClusterClusterNestedStructList")}
	})
	return MTRTestClusterClusterNestedStructListClass
}

type _MTRTestClusterClusterNestedStructListClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterNestedStructList] class.
type IMTRTestClusterClusterNestedStructList interface {
	IMTRUnitTestingClusterNestedStructList
	// properties:
	A() objc.IObject /* cross-framework: NSNumber */
	SetA(value objc.IObject /* cross-framework: NSNumber */)
	B() objc.IObject /* cross-framework: NSNumber */
	SetB(value objc.IObject /* cross-framework: NSNumber */)
	C() IMTRTestClusterClusterSimpleStruct
	SetC(value IMTRTestClusterClusterSimpleStruct)
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterNestedStructList
type MTRTestClusterClusterNestedStructList struct {
	MTRUnitTestingClusterNestedStructList
}

// MTRTestClusterClusterNestedStructListFrom constructs a [MTRTestClusterClusterNestedStructList] from an unsafe.Pointer.
func MTRTestClusterClusterNestedStructListFrom(ptr unsafe.Pointer) MTRTestClusterClusterNestedStructList {
	return MTRTestClusterClusterNestedStructList{
		MTRUnitTestingClusterNestedStructList: MTRUnitTestingClusterNestedStructListFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterNestedStructListClass) Alloc() MTRTestClusterClusterNestedStructList {
	rv := objc.Send[MTRTestClusterClusterNestedStructList](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterNestedStructListClass) New() MTRTestClusterClusterNestedStructList {
	rv := objc.Send[MTRTestClusterClusterNestedStructList](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterNestedStructList) Init() MTRTestClusterClusterNestedStructList {
	rv := objc.Send[MTRTestClusterClusterNestedStructList](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterNestedStructList) Autorelease() MTRTestClusterClusterNestedStructList {
	rv := objc.Send[MTRTestClusterClusterNestedStructList](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterNestedStructList creates a new MTRTestClusterClusterNestedStructList instance.
func NewMTRTestClusterClusterNestedStructList() MTRTestClusterClusterNestedStructList {
	return getMTRTestClusterClusterNestedStructListClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/a
func (m_ MTRTestClusterClusterNestedStructList) A() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("a"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/a
func (m_ MTRTestClusterClusterNestedStructList) SetA(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/b
func (m_ MTRTestClusterClusterNestedStructList) B() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("b"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/b
func (m_ MTRTestClusterClusterNestedStructList) SetB(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/c
func (m_ MTRTestClusterClusterNestedStructList) C() IMTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("c"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/c
func (m_ MTRTestClusterClusterNestedStructList) SetC(value IMTRTestClusterClusterSimpleStruct) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/d
func (m_ MTRTestClusterClusterNestedStructList) D() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("d"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/d
func (m_ MTRTestClusterClusterNestedStructList) SetD(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setD:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/e
func (m_ MTRTestClusterClusterNestedStructList) E() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("e"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/e
func (m_ MTRTestClusterClusterNestedStructList) SetE(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setE:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/f
func (m_ MTRTestClusterClusterNestedStructList) F() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("f"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/f
func (m_ MTRTestClusterClusterNestedStructList) SetF(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setF:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/g
func (m_ MTRTestClusterClusterNestedStructList) G() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("g"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclusternestedstructlist/g
func (m_ MTRTestClusterClusterNestedStructList) SetG(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setG:"), value)
}
