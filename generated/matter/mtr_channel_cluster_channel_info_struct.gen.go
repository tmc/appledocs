// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterChannelInfoStruct] class.
var (
	MTRChannelClusterChannelInfoStructClass     _MTRChannelClusterChannelInfoStructClass
	MTRChannelClusterChannelInfoStructClassOnce sync.Once
)

func getMTRChannelClusterChannelInfoStructClass() _MTRChannelClusterChannelInfoStructClass {
	MTRChannelClusterChannelInfoStructClassOnce.Do(func() {
		MTRChannelClusterChannelInfoStructClass = _MTRChannelClusterChannelInfoStructClass{objc.GetClass("MTRChannelClusterChannelInfoStruct")}
	})
	return MTRChannelClusterChannelInfoStructClass
}

type _MTRChannelClusterChannelInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterChannelInfoStruct] class.
type IMTRChannelClusterChannelInfoStruct interface {
	objectivec.IObject
	// properties:
	AffiliateCallSign() objc.IObject /* cross-framework: NSString */
	SetAffiliateCallSign(value objc.IObject /* cross-framework: NSString */)
	CallSign() objc.IObject /* cross-framework: NSString */
	SetCallSign(value objc.IObject /* cross-framework: NSString */)
	Identifier() objc.IObject /* cross-framework: NSString */
	SetIdentifier(value objc.IObject /* cross-framework: NSString */)
	MajorNumber() objc.IObject /* cross-framework: NSNumber */
	SetMajorNumber(value objc.IObject /* cross-framework: NSNumber */)
	MinorNumber() objc.IObject /* cross-framework: NSNumber */
	SetMinorNumber(value objc.IObject /* cross-framework: NSNumber */)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Type() objc.IObject /* cross-framework: NSNumber */
	SetType(value objc.IObject /* cross-framework: NSNumber */)
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterChannelInfoStruct
type MTRChannelClusterChannelInfoStruct struct {
	objectivec.Object
}

// MTRChannelClusterChannelInfoStructFrom constructs a [MTRChannelClusterChannelInfoStruct] from an unsafe.Pointer.
func MTRChannelClusterChannelInfoStructFrom(ptr unsafe.Pointer) MTRChannelClusterChannelInfoStruct {
	return MTRChannelClusterChannelInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterChannelInfoStructClass) Alloc() MTRChannelClusterChannelInfoStruct {
	rv := objc.Send[MTRChannelClusterChannelInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterChannelInfoStructClass) New() MTRChannelClusterChannelInfoStruct {
	rv := objc.Send[MTRChannelClusterChannelInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterChannelInfoStruct) Init() MTRChannelClusterChannelInfoStruct {
	rv := objc.Send[MTRChannelClusterChannelInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterChannelInfoStruct) Autorelease() MTRChannelClusterChannelInfoStruct {
	rv := objc.Send[MTRChannelClusterChannelInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterChannelInfoStruct creates a new MTRChannelClusterChannelInfoStruct instance.
func NewMTRChannelClusterChannelInfoStruct() MTRChannelClusterChannelInfoStruct {
	return getMTRChannelClusterChannelInfoStructClass().New()
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/affiliatecallsign
func (m_ MTRChannelClusterChannelInfoStruct) AffiliateCallSign() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("affiliateCallSign"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/affiliatecallsign
func (m_ MTRChannelClusterChannelInfoStruct) SetAffiliateCallSign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAffiliateCallSign:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/callsign
func (m_ MTRChannelClusterChannelInfoStruct) CallSign() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("callSign"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/callsign
func (m_ MTRChannelClusterChannelInfoStruct) SetCallSign(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCallSign:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/identifier
func (m_ MTRChannelClusterChannelInfoStruct) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("identifier"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/identifier
func (m_ MTRChannelClusterChannelInfoStruct) SetIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/majornumber
func (m_ MTRChannelClusterChannelInfoStruct) MajorNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("majorNumber"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/majornumber
func (m_ MTRChannelClusterChannelInfoStruct) SetMajorNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMajorNumber:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/minornumber
func (m_ MTRChannelClusterChannelInfoStruct) MinorNumber() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("minorNumber"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/minornumber
func (m_ MTRChannelClusterChannelInfoStruct) SetMinorNumber(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinorNumber:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/name
func (m_ MTRChannelClusterChannelInfoStruct) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/name
func (m_ MTRChannelClusterChannelInfoStruct) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/type
func (m_ MTRChannelClusterChannelInfoStruct) Type() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("type"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/type
func (m_ MTRChannelClusterChannelInfoStruct) SetType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}
