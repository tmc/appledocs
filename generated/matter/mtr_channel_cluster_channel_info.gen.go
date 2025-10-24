// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	AffiliateCallSign() objc.IObject /* cross-framework: NSString */
	SetAffiliateCallSign(value objc.IObject /* cross-framework: NSString */)
	CallSign() objc.IObject /* cross-framework: NSString */
	SetCallSign(value objc.IObject /* cross-framework: NSString */)
	MajorNumber() objc.IObject /* cross-framework: NSNumber */
	SetMajorNumber(value objc.IObject /* cross-framework: NSNumber */)
	MinorNumber() objc.IObject /* cross-framework: NSNumber */
	SetMinorNumber(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/affiliatecallsign
func (m_ MTRChannelClusterChannelInfo) AffiliateCallSign() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("affiliateCallSign"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/affiliatecallsign
func (m_ MTRChannelClusterChannelInfo) SetAffiliateCallSign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAffiliateCallSign:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/callsign
func (m_ MTRChannelClusterChannelInfo) CallSign() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("callSign"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/callsign
func (m_ MTRChannelClusterChannelInfo) SetCallSign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCallSign:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/majornumber
func (m_ MTRChannelClusterChannelInfo) MajorNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("majorNumber"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/majornumber
func (m_ MTRChannelClusterChannelInfo) SetMajorNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMajorNumber:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/minornumber
func (m_ MTRChannelClusterChannelInfo) MinorNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minorNumber"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/minornumber
func (m_ MTRChannelClusterChannelInfo) SetMinorNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinorNumber:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/name
func (m_ MTRChannelClusterChannelInfo) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfo/name
func (m_ MTRChannelClusterChannelInfo) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}
