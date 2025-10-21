// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRChannelClusterLineupInfo] class.
var (
	MTRChannelClusterLineupInfoClass     _MTRChannelClusterLineupInfoClass
	MTRChannelClusterLineupInfoClassOnce sync.Once
)

func getMTRChannelClusterLineupInfoClass() _MTRChannelClusterLineupInfoClass {
	MTRChannelClusterLineupInfoClassOnce.Do(func() {
		MTRChannelClusterLineupInfoClass = _MTRChannelClusterLineupInfoClass{objc.GetClass("MTRChannelClusterLineupInfo")}
	})
	return MTRChannelClusterLineupInfoClass
}

type _MTRChannelClusterLineupInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterLineupInfo] class.
type IMTRChannelClusterLineupInfo interface {
	IMTRChannelClusterLineupInfoStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterLineupInfo
type MTRChannelClusterLineupInfo struct {
	MTRChannelClusterLineupInfoStruct
}

// MTRChannelClusterLineupInfoFrom constructs a [MTRChannelClusterLineupInfo] from an unsafe.Pointer.
func MTRChannelClusterLineupInfoFrom(ptr unsafe.Pointer) MTRChannelClusterLineupInfo {
	return MTRChannelClusterLineupInfo{
		MTRChannelClusterLineupInfoStruct: MTRChannelClusterLineupInfoStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterLineupInfoClass) Alloc() MTRChannelClusterLineupInfo {
	rv := objc.Send[MTRChannelClusterLineupInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterLineupInfoClass) New() MTRChannelClusterLineupInfo {
	rv := objc.Send[MTRChannelClusterLineupInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterLineupInfo) Init() MTRChannelClusterLineupInfo {
	rv := objc.Send[MTRChannelClusterLineupInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterLineupInfo) Autorelease() MTRChannelClusterLineupInfo {
	rv := objc.Send[MTRChannelClusterLineupInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterLineupInfo creates a new MTRChannelClusterLineupInfo instance.
func NewMTRChannelClusterLineupInfo() MTRChannelClusterLineupInfo {
	return getMTRChannelClusterLineupInfoClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/lineupinfotype
func (m_ MTRChannelClusterLineupInfo) LineupInfoType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lineupInfoType"))
	return rv
}


// SetLineupInfoType sets the value of the lineupInfoType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/lineupinfotype
func (m_ MTRChannelClusterLineupInfo) SetLineupInfoType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineupInfoType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/lineupname
func (m_ MTRChannelClusterLineupInfo) LineupName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("lineupName"))
	return rv
}


// SetLineupName sets the value of the lineupName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/lineupname
func (m_ MTRChannelClusterLineupInfo) SetLineupName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineupName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/operatorname
func (m_ MTRChannelClusterLineupInfo) OperatorName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("operatorName"))
	return rv
}


// SetOperatorName sets the value of the operatorName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/operatorname
func (m_ MTRChannelClusterLineupInfo) SetOperatorName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperatorName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/postalcode
func (m_ MTRChannelClusterLineupInfo) PostalCode() string {
	rv := objc.Send[string](m_.ID, objc.Sel("postalCode"))
	return rv
}


// SetPostalCode sets the value of the postalCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/postalcode
func (m_ MTRChannelClusterLineupInfo) SetPostalCode(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPostalCode:"), objc.String(value))
}



