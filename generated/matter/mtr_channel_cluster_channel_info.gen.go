// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRChannelClusterChannelInfo] class.
var (
	MTRChannelClusterChannelInfoClass     _MTRChannelClusterChannelInfoClass
	MTRChannelClusterChannelInfoClassOnce sync.Once
)

func getMTRChannelClusterChannelInfoClass() _MTRChannelClusterChannelInfoClass {
	MTRChannelClusterChannelInfoClassOnce.Do(func() {
		MTRChannelClusterChannelInfoClass = _MTRChannelClusterChannelInfoClass{objc.GetClass("MTRChannelClusterChannelInfo")}
	})
	return MTRChannelClusterChannelInfoClass
}

type _MTRChannelClusterChannelInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterChannelInfo] class.
type IMTRChannelClusterChannelInfo interface {
	IMTRChannelClusterChannelInfoStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelInfo
type MTRChannelClusterChannelInfo struct {
	MTRChannelClusterChannelInfoStruct
}

// MTRChannelClusterChannelInfoFrom constructs a [MTRChannelClusterChannelInfo] from an unsafe.Pointer.
func MTRChannelClusterChannelInfoFrom(ptr unsafe.Pointer) MTRChannelClusterChannelInfo {
	return MTRChannelClusterChannelInfo{
		MTRChannelClusterChannelInfoStruct: MTRChannelClusterChannelInfoStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterChannelInfoClass) Alloc() MTRChannelClusterChannelInfo {
	rv := objc.Send[MTRChannelClusterChannelInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterChannelInfoClass) New() MTRChannelClusterChannelInfo {
	rv := objc.Send[MTRChannelClusterChannelInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterChannelInfo) Init() MTRChannelClusterChannelInfo {
	rv := objc.Send[MTRChannelClusterChannelInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterChannelInfo) Autorelease() MTRChannelClusterChannelInfo {
	rv := objc.Send[MTRChannelClusterChannelInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterChannelInfo creates a new MTRChannelClusterChannelInfo instance.
func NewMTRChannelClusterChannelInfo() MTRChannelClusterChannelInfo {
	return getMTRChannelClusterChannelInfoClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/affiliatecallsign
func (m_ MTRChannelClusterChannelInfo) AffiliateCallSign() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("affiliateCallSign"))
	return rv
}


// SetAffiliateCallSign sets the value of the affiliateCallSign property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/affiliatecallsign
func (m_ MTRChannelClusterChannelInfo) SetAffiliateCallSign(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAffiliateCallSign:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/callsign
func (m_ MTRChannelClusterChannelInfo) CallSign() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("callSign"))
	return rv
}


// SetCallSign sets the value of the callSign property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/callsign
func (m_ MTRChannelClusterChannelInfo) SetCallSign(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCallSign:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/majornumber
func (m_ MTRChannelClusterChannelInfo) MajorNumber() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("majorNumber"))
	return rv
}


// SetMajorNumber sets the value of the majorNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/majornumber
func (m_ MTRChannelClusterChannelInfo) SetMajorNumber(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMajorNumber:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/minornumber
func (m_ MTRChannelClusterChannelInfo) MinorNumber() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minorNumber"))
	return rv
}


// SetMinorNumber sets the value of the minorNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/minornumber
func (m_ MTRChannelClusterChannelInfo) SetMinorNumber(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinorNumber:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/name
func (m_ MTRChannelClusterChannelInfo) Name() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/name
func (m_ MTRChannelClusterChannelInfo) SetName(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}



