// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTimeSynchronizationClusterTimeZoneStruct] class.
var (
	MTRTimeSynchronizationClusterTimeZoneStructClass     _MTRTimeSynchronizationClusterTimeZoneStructClass
	MTRTimeSynchronizationClusterTimeZoneStructClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterTimeZoneStructClass() _MTRTimeSynchronizationClusterTimeZoneStructClass {
	MTRTimeSynchronizationClusterTimeZoneStructClassOnce.Do(func() {
		MTRTimeSynchronizationClusterTimeZoneStructClass = _MTRTimeSynchronizationClusterTimeZoneStructClass{objc.GetClass("MTRTimeSynchronizationClusterTimeZoneStruct")}
	})
	return MTRTimeSynchronizationClusterTimeZoneStructClass
}

type _MTRTimeSynchronizationClusterTimeZoneStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterTimeZoneStruct] class.
type IMTRTimeSynchronizationClusterTimeZoneStruct interface {
	objectivec.IObject
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Offset() objc.IObject /* cross-framework: NSNumber */
	SetOffset(value objc.IObject /* cross-framework: NSNumber */)
	ValidAt() objc.IObject /* cross-framework: NSNumber */
	SetValidAt(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStruct
type MTRTimeSynchronizationClusterTimeZoneStruct struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterTimeZoneStructFrom constructs a [MTRTimeSynchronizationClusterTimeZoneStruct] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterTimeZoneStructFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterTimeZoneStruct {
	return MTRTimeSynchronizationClusterTimeZoneStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterTimeZoneStructClass) Alloc() MTRTimeSynchronizationClusterTimeZoneStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterTimeZoneStructClass) New() MTRTimeSynchronizationClusterTimeZoneStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) Init() MTRTimeSynchronizationClusterTimeZoneStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) Autorelease() MTRTimeSynchronizationClusterTimeZoneStruct {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterTimeZoneStruct creates a new MTRTimeSynchronizationClusterTimeZoneStruct instance.
func NewMTRTimeSynchronizationClusterTimeZoneStruct() MTRTimeSynchronizationClusterTimeZoneStruct {
	return getMTRTimeSynchronizationClusterTimeZoneStructClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonestruct/name
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonestruct/name
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonestruct/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) Offset() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("offset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonestruct/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) SetOffset(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonestruct/validat
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) ValidAt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("validAt"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonestruct/validat
func (m_ MTRTimeSynchronizationClusterTimeZoneStruct) SetValidAt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidAt:"), value)
}



