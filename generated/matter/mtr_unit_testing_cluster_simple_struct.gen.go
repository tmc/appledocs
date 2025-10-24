// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRUnitTestingClusterSimpleStruct] class.
var (
	MTRUnitTestingClusterSimpleStructClass     _MTRUnitTestingClusterSimpleStructClass
	MTRUnitTestingClusterSimpleStructClassOnce sync.Once
)

func getMTRUnitTestingClusterSimpleStructClass() _MTRUnitTestingClusterSimpleStructClass {
	MTRUnitTestingClusterSimpleStructClassOnce.Do(func() {
		MTRUnitTestingClusterSimpleStructClass = _MTRUnitTestingClusterSimpleStructClass{objc.GetClass("MTRUnitTestingClusterSimpleStruct")}
	})
	return MTRUnitTestingClusterSimpleStructClass
}

type _MTRUnitTestingClusterSimpleStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRUnitTestingClusterSimpleStruct] class.
type IMTRUnitTestingClusterSimpleStruct interface {
	objectivec.IObject
	// properties:
	A() objc.IObject /* cross-framework: NSNumber */
	SetA(value objc.IObject /* cross-framework: NSNumber */)
	B() objc.IObject /* cross-framework: NSNumber */
	SetB(value objc.IObject /* cross-framework: NSNumber */)
	C() objc.IObject /* cross-framework: NSNumber */
	SetC(value objc.IObject /* cross-framework: NSNumber */)
	D() objc.IObject /* cross-framework: Data */
	SetD(value objc.IObject /* cross-framework: Data */)
	E() objc.IObject /* cross-framework: NSString */
	SetE(value objc.IObject /* cross-framework: NSString */)
	F() objc.IObject /* cross-framework: NSNumber */
	SetF(value objc.IObject /* cross-framework: NSNumber */)
	G() objc.IObject /* cross-framework: NSNumber */
	SetG(value objc.IObject /* cross-framework: NSNumber */)
	H() objc.IObject /* cross-framework: NSNumber */
	SetH(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRUnitTestingClusterSimpleStruct
type MTRUnitTestingClusterSimpleStruct struct {
	objectivec.Object
}

// MTRUnitTestingClusterSimpleStructFrom constructs a [MTRUnitTestingClusterSimpleStruct] from an unsafe.Pointer.
func MTRUnitTestingClusterSimpleStructFrom(ptr unsafe.Pointer) MTRUnitTestingClusterSimpleStruct {
	return MTRUnitTestingClusterSimpleStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRUnitTestingClusterSimpleStructClass) Alloc() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRUnitTestingClusterSimpleStructClass) New() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRUnitTestingClusterSimpleStruct) Init() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRUnitTestingClusterSimpleStruct) Autorelease() MTRUnitTestingClusterSimpleStruct {
	rv := objc.Send[MTRUnitTestingClusterSimpleStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRUnitTestingClusterSimpleStruct creates a new MTRUnitTestingClusterSimpleStruct instance.
func NewMTRUnitTestingClusterSimpleStruct() MTRUnitTestingClusterSimpleStruct {
	return getMTRUnitTestingClusterSimpleStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/a
func (m_ MTRUnitTestingClusterSimpleStruct) A() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("a"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/a
func (m_ MTRUnitTestingClusterSimpleStruct) SetA(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/b
func (m_ MTRUnitTestingClusterSimpleStruct) B() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("b"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/b
func (m_ MTRUnitTestingClusterSimpleStruct) SetB(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/c
func (m_ MTRUnitTestingClusterSimpleStruct) C() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("c"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/c
func (m_ MTRUnitTestingClusterSimpleStruct) SetC(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/d
func (m_ MTRUnitTestingClusterSimpleStruct) D() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("d"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/d
func (m_ MTRUnitTestingClusterSimpleStruct) SetD(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setD:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/e
func (m_ MTRUnitTestingClusterSimpleStruct) E() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("e"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/e
func (m_ MTRUnitTestingClusterSimpleStruct) SetE(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setE:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/f
func (m_ MTRUnitTestingClusterSimpleStruct) F() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("f"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/f
func (m_ MTRUnitTestingClusterSimpleStruct) SetF(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setF:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/g
func (m_ MTRUnitTestingClusterSimpleStruct) G() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("g"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/g
func (m_ MTRUnitTestingClusterSimpleStruct) SetG(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setG:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/h
func (m_ MTRUnitTestingClusterSimpleStruct) H() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("h"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrunittestingclustersimplestruct/h
func (m_ MTRUnitTestingClusterSimpleStruct) SetH(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setH:"), value)
}



