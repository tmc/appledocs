// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTimeSynchronizationClusterDstOffsetType] class.
var (
	MTRTimeSynchronizationClusterDstOffsetTypeClass     _MTRTimeSynchronizationClusterDstOffsetTypeClass
	MTRTimeSynchronizationClusterDstOffsetTypeClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterDstOffsetTypeClass() _MTRTimeSynchronizationClusterDstOffsetTypeClass {
	MTRTimeSynchronizationClusterDstOffsetTypeClassOnce.Do(func() {
		MTRTimeSynchronizationClusterDstOffsetTypeClass = _MTRTimeSynchronizationClusterDstOffsetTypeClass{objc.GetClass("MTRTimeSynchronizationClusterDstOffsetType")}
	})
	return MTRTimeSynchronizationClusterDstOffsetTypeClass
}

type _MTRTimeSynchronizationClusterDstOffsetTypeClass struct {
	class objc.Class
}

// An interface definition for the [MTRTimeSynchronizationClusterDstOffsetType] class.
type IMTRTimeSynchronizationClusterDstOffsetType interface {
	IMTRTimeSynchronizationClusterDSTOffsetStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterDstOffsetType
type MTRTimeSynchronizationClusterDstOffsetType struct {
	MTRTimeSynchronizationClusterDSTOffsetStruct
}

// MTRTimeSynchronizationClusterDstOffsetTypeFrom constructs a [MTRTimeSynchronizationClusterDstOffsetType] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterDstOffsetTypeFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterDstOffsetType {
	return MTRTimeSynchronizationClusterDstOffsetType{
		MTRTimeSynchronizationClusterDSTOffsetStruct: MTRTimeSynchronizationClusterDSTOffsetStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterDstOffsetTypeClass) Alloc() MTRTimeSynchronizationClusterDstOffsetType {
	rv := objc.Send[MTRTimeSynchronizationClusterDstOffsetType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTimeSynchronizationClusterDstOffsetTypeClass) New() MTRTimeSynchronizationClusterDstOffsetType {
	rv := objc.Send[MTRTimeSynchronizationClusterDstOffsetType](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterDstOffsetType) Init() MTRTimeSynchronizationClusterDstOffsetType {
	rv := objc.Send[MTRTimeSynchronizationClusterDstOffsetType](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterDstOffsetType) Autorelease() MTRTimeSynchronizationClusterDstOffsetType {
	rv := objc.Send[MTRTimeSynchronizationClusterDstOffsetType](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterDstOffsetType creates a new MTRTimeSynchronizationClusterDstOffsetType instance.
func NewMTRTimeSynchronizationClusterDstOffsetType() MTRTimeSynchronizationClusterDstOffsetType {
	return getMTRTimeSynchronizationClusterDstOffsetTypeClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclusterdstoffsettype/offset
func (m_ MTRTimeSynchronizationClusterDstOffsetType) Offset() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("offset"))
	return rv
}


// SetOffset sets the value of the offset property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclusterdstoffsettype/offset
func (m_ MTRTimeSynchronizationClusterDstOffsetType) SetOffset(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclusterdstoffsettype/validstarting
func (m_ MTRTimeSynchronizationClusterDstOffsetType) ValidStarting() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("validStarting"))
	return rv
}


// SetValidStarting sets the value of the validStarting property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclusterdstoffsettype/validstarting
func (m_ MTRTimeSynchronizationClusterDstOffsetType) SetValidStarting(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidStarting:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclusterdstoffsettype/validuntil
func (m_ MTRTimeSynchronizationClusterDstOffsetType) ValidUntil() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("validUntil"))
	return rv
}


// SetValidUntil sets the value of the validUntil property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclusterdstoffsettype/validuntil
func (m_ MTRTimeSynchronizationClusterDstOffsetType) SetValidUntil(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setValidUntil:"), value)
}



