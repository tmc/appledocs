// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
}

//
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


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/affiliatecallsign
func (m_ MTRChannelClusterChannelInfoStruct) AffiliateCallSign() string {
	rv := objc.Send[string](m_.ID, objc.Sel("affiliateCallSign"))
	return rv
}


// SetAffiliateCallSign sets the value of the affiliateCallSign property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/affiliatecallsign
func (m_ MTRChannelClusterChannelInfoStruct) SetAffiliateCallSign(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAffiliateCallSign:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/callsign
func (m_ MTRChannelClusterChannelInfoStruct) CallSign() string {
	rv := objc.Send[string](m_.ID, objc.Sel("callSign"))
	return rv
}


// SetCallSign sets the value of the callSign property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/callsign
func (m_ MTRChannelClusterChannelInfoStruct) SetCallSign(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCallSign:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/identifier
func (m_ MTRChannelClusterChannelInfoStruct) Identifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/identifier
func (m_ MTRChannelClusterChannelInfoStruct) SetIdentifier(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/majornumber
func (m_ MTRChannelClusterChannelInfoStruct) MajorNumber() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("majorNumber"))
	return rv
}


// SetMajorNumber sets the value of the majorNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/majornumber
func (m_ MTRChannelClusterChannelInfoStruct) SetMajorNumber(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMajorNumber:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/minornumber
func (m_ MTRChannelClusterChannelInfoStruct) MinorNumber() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("minorNumber"))
	return rv
}


// SetMinorNumber sets the value of the minorNumber property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/minornumber
func (m_ MTRChannelClusterChannelInfoStruct) SetMinorNumber(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMinorNumber:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/name
func (m_ MTRChannelClusterChannelInfoStruct) Name() string {
	rv := objc.Send[string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/name
func (m_ MTRChannelClusterChannelInfoStruct) SetName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/type
func (m_ MTRChannelClusterChannelInfoStruct) Type() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("type"))
	return rv
}


// SetType sets the value of the type property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterchannelinfostruct/type
func (m_ MTRChannelClusterChannelInfoStruct) SetType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setType:"), value)
}



