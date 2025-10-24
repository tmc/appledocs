// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTimeSynchronizationClusterTimeZoneType] class.
var (
	MTRTimeSynchronizationClusterTimeZoneTypeClass     _MTRTimeSynchronizationClusterTimeZoneTypeClass
	MTRTimeSynchronizationClusterTimeZoneTypeClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterTimeZoneTypeClass() _MTRTimeSynchronizationClusterTimeZoneTypeClass {
	MTRTimeSynchronizationClusterTimeZoneTypeClassOnce.Do(func() {
		MTRTimeSynchronizationClusterTimeZoneTypeClass = _MTRTimeSynchronizationClusterTimeZoneTypeClass{objc.GetClass("MTRTimeSynchronizationClusterTimeZoneType")}
	})
	return MTRTimeSynchronizationClusterTimeZoneTypeClass
}

type _MTRTimeSynchronizationClusterTimeZoneTypeClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterTimeZoneType] class.
type IMTRTimeSynchronizationClusterTimeZoneType interface {
	IMTRTimeSynchronizationClusterTimeZoneStruct
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
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneType
type MTRTimeSynchronizationClusterTimeZoneType struct {
	MTRTimeSynchronizationClusterTimeZoneStruct
}

// MTRTimeSynchronizationClusterTimeZoneTypeFrom constructs a [MTRTimeSynchronizationClusterTimeZoneType] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterTimeZoneTypeFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterTimeZoneType {
	return MTRTimeSynchronizationClusterTimeZoneType{
		MTRTimeSynchronizationClusterTimeZoneStruct: MTRTimeSynchronizationClusterTimeZoneStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterTimeZoneTypeClass) Alloc() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterTimeZoneTypeClass) New() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Init() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Autorelease() MTRTimeSynchronizationClusterTimeZoneType {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterTimeZoneType creates a new MTRTimeSynchronizationClusterTimeZoneType instance.
func NewMTRTimeSynchronizationClusterTimeZoneType() MTRTimeSynchronizationClusterTimeZoneType {
	return getMTRTimeSynchronizationClusterTimeZoneTypeClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/name
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/name
func (m_ MTRTimeSynchronizationClusterTimeZoneType) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Offset() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("offset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneType) SetOffset(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/validat
func (m_ MTRTimeSynchronizationClusterTimeZoneType) ValidAt() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("validAt"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/validat
func (m_ MTRTimeSynchronizationClusterTimeZoneType) SetValidAt(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidAt:"), value)
}



