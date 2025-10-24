// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRTestClusterClusterSimpleStruct] class.
var (
	MTRTestClusterClusterSimpleStructClass     _MTRTestClusterClusterSimpleStructClass
	MTRTestClusterClusterSimpleStructClassOnce sync.Once
)

func getMTRTestClusterClusterSimpleStructClass() _MTRTestClusterClusterSimpleStructClass {
	MTRTestClusterClusterSimpleStructClassOnce.Do(func() {
		MTRTestClusterClusterSimpleStructClass = _MTRTestClusterClusterSimpleStructClass{objc.GetClass("MTRTestClusterClusterSimpleStruct")}
	})
	return MTRTestClusterClusterSimpleStructClass
}

type _MTRTestClusterClusterSimpleStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTestClusterClusterSimpleStruct] class.
type IMTRTestClusterClusterSimpleStruct interface {
	IMTRUnitTestingClusterSimpleStruct
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTestClusterClusterSimpleStruct
type MTRTestClusterClusterSimpleStruct struct {
	MTRUnitTestingClusterSimpleStruct
}

// MTRTestClusterClusterSimpleStructFrom constructs a [MTRTestClusterClusterSimpleStruct] from an unsafe.Pointer.
func MTRTestClusterClusterSimpleStructFrom(ptr unsafe.Pointer) MTRTestClusterClusterSimpleStruct {
	return MTRTestClusterClusterSimpleStruct{
		MTRUnitTestingClusterSimpleStruct: MTRUnitTestingClusterSimpleStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTestClusterClusterSimpleStructClass) Alloc() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTestClusterClusterSimpleStructClass) New() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTestClusterClusterSimpleStruct) Init() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTestClusterClusterSimpleStruct) Autorelease() MTRTestClusterClusterSimpleStruct {
	rv := objc.Send[MTRTestClusterClusterSimpleStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTestClusterClusterSimpleStruct creates a new MTRTestClusterClusterSimpleStruct instance.
func NewMTRTestClusterClusterSimpleStruct() MTRTestClusterClusterSimpleStruct {
	return getMTRTestClusterClusterSimpleStructClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/a
func (m_ MTRTestClusterClusterSimpleStruct) A() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("a"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/a
func (m_ MTRTestClusterClusterSimpleStruct) SetA(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setA:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/b
func (m_ MTRTestClusterClusterSimpleStruct) B() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("b"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/b
func (m_ MTRTestClusterClusterSimpleStruct) SetB(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setB:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/c
func (m_ MTRTestClusterClusterSimpleStruct) C() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("c"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/c
func (m_ MTRTestClusterClusterSimpleStruct) SetC(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setC:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/d
func (m_ MTRTestClusterClusterSimpleStruct) D() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("d"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/d
func (m_ MTRTestClusterClusterSimpleStruct) SetD(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setD:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/e
func (m_ MTRTestClusterClusterSimpleStruct) E() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("e"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/e
func (m_ MTRTestClusterClusterSimpleStruct) SetE(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setE:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/f
func (m_ MTRTestClusterClusterSimpleStruct) F() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("f"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/f
func (m_ MTRTestClusterClusterSimpleStruct) SetF(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setF:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/g
func (m_ MTRTestClusterClusterSimpleStruct) G() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("g"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/g
func (m_ MTRTestClusterClusterSimpleStruct) SetG(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setG:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/h
func (m_ MTRTestClusterClusterSimpleStruct) H() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("h"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtestclusterclustersimplestruct/h
func (m_ MTRTestClusterClusterSimpleStruct) SetH(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setH:"), value)
}
