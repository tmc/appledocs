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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/name
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/name
func (m_ MTRTimeSynchronizationClusterTimeZoneType) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneType) Offset() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("offset"))
	return rv
}


// SetOffset sets the value of the offset property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneType) SetOffset(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/validat
func (m_ MTRTimeSynchronizationClusterTimeZoneType) ValidAt() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("validAt"))
	return rv
}


// SetValidAt sets the value of the validAt property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonetype/validat
func (m_ MTRTimeSynchronizationClusterTimeZoneType) SetValidAt(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidAt:"), value)
}



