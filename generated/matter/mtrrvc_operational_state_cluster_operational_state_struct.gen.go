// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRRVCOperationalStateClusterOperationalStateStruct] class.
var (
	MTRRVCOperationalStateClusterOperationalStateStructClass     _MTRRVCOperationalStateClusterOperationalStateStructClass
	MTRRVCOperationalStateClusterOperationalStateStructClassOnce sync.Once
)

func getMTRRVCOperationalStateClusterOperationalStateStructClass() _MTRRVCOperationalStateClusterOperationalStateStructClass {
	MTRRVCOperationalStateClusterOperationalStateStructClassOnce.Do(func() {
		MTRRVCOperationalStateClusterOperationalStateStructClass = _MTRRVCOperationalStateClusterOperationalStateStructClass{objc.GetClass("MTRRVCOperationalStateClusterOperationalStateStruct")}
	})
	return MTRRVCOperationalStateClusterOperationalStateStructClass
}

type _MTRRVCOperationalStateClusterOperationalStateStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRRVCOperationalStateClusterOperationalStateStruct] class.
type IMTRRVCOperationalStateClusterOperationalStateStruct interface {
	objectivec.IObject
	// properties:
	OperationalStateID() objc.IObject /* cross-framework: NSNumber */
	SetOperationalStateID(value objc.IObject /* cross-framework: NSNumber */)
	OperationalStateLabel() objc.IObject /* cross-framework: NSString */
	SetOperationalStateLabel(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRRVCOperationalStateClusterOperationalStateStruct
type MTRRVCOperationalStateClusterOperationalStateStruct struct {
	objectivec.Object
}

// MTRRVCOperationalStateClusterOperationalStateStructFrom constructs a [MTRRVCOperationalStateClusterOperationalStateStruct] from an unsafe.Pointer.
func MTRRVCOperationalStateClusterOperationalStateStructFrom(ptr unsafe.Pointer) MTRRVCOperationalStateClusterOperationalStateStruct {
	return MTRRVCOperationalStateClusterOperationalStateStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRRVCOperationalStateClusterOperationalStateStructClass) Alloc() MTRRVCOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalStateStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRRVCOperationalStateClusterOperationalStateStructClass) New() MTRRVCOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalStateStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) Init() MTRRVCOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalStateStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) Autorelease() MTRRVCOperationalStateClusterOperationalStateStruct {
	rv := objc.Send[MTRRVCOperationalStateClusterOperationalStateStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRRVCOperationalStateClusterOperationalStateStruct creates a new MTRRVCOperationalStateClusterOperationalStateStruct instance.
func NewMTRRVCOperationalStateClusterOperationalStateStruct() MTRRVCOperationalStateClusterOperationalStateStruct {
	return getMTRRVCOperationalStateClusterOperationalStateStructClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalstatestruct/operationalstateid
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) OperationalStateID() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("operationalStateID"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalstatestruct/operationalstateid
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) SetOperationalStateID(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalStateID:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalstatestruct/operationalstatelabel
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) OperationalStateLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("operationalStateLabel"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrrvcoperationalstateclusteroperationalstatestruct/operationalstatelabel
func (m_ MTRRVCOperationalStateClusterOperationalStateStruct) SetOperationalStateLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperationalStateLabel:"), value)
}
