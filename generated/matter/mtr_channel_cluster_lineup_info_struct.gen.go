// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRChannelClusterLineupInfoStruct] class.
var (
	MTRChannelClusterLineupInfoStructClass     _MTRChannelClusterLineupInfoStructClass
	MTRChannelClusterLineupInfoStructClassOnce sync.Once
)

func getMTRChannelClusterLineupInfoStructClass() _MTRChannelClusterLineupInfoStructClass {
	MTRChannelClusterLineupInfoStructClassOnce.Do(func() {
		MTRChannelClusterLineupInfoStructClass = _MTRChannelClusterLineupInfoStructClass{objc.GetClass("MTRChannelClusterLineupInfoStruct")}
	})
	return MTRChannelClusterLineupInfoStructClass
}

type _MTRChannelClusterLineupInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRChannelClusterLineupInfoStruct] class.
type IMTRChannelClusterLineupInfoStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterLineupInfoStruct
type MTRChannelClusterLineupInfoStruct struct {
	objectivec.Object
}

// MTRChannelClusterLineupInfoStructFrom constructs a [MTRChannelClusterLineupInfoStruct] from an unsafe.Pointer.
func MTRChannelClusterLineupInfoStructFrom(ptr unsafe.Pointer) MTRChannelClusterLineupInfoStruct {
	return MTRChannelClusterLineupInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterLineupInfoStructClass) Alloc() MTRChannelClusterLineupInfoStruct {
	rv := objc.Send[MTRChannelClusterLineupInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRChannelClusterLineupInfoStructClass) New() MTRChannelClusterLineupInfoStruct {
	rv := objc.Send[MTRChannelClusterLineupInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterLineupInfoStruct) Init() MTRChannelClusterLineupInfoStruct {
	rv := objc.Send[MTRChannelClusterLineupInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterLineupInfoStruct) Autorelease() MTRChannelClusterLineupInfoStruct {
	rv := objc.Send[MTRChannelClusterLineupInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterLineupInfoStruct creates a new MTRChannelClusterLineupInfoStruct instance.
func NewMTRChannelClusterLineupInfoStruct() MTRChannelClusterLineupInfoStruct {
	return getMTRChannelClusterLineupInfoStructClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfostruct/operatorname
func (m_ MTRChannelClusterLineupInfoStruct) OperatorName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("operatorName"))
	return rv
}


// SetOperatorName sets the value of the operatorName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfostruct/operatorname
func (m_ MTRChannelClusterLineupInfoStruct) SetOperatorName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperatorName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfostruct/lineupname
func (m_ MTRChannelClusterLineupInfoStruct) LineupName() string {
	rv := objc.Send[string](m_.ID, objc.Sel("lineupName"))
	return rv
}


// SetLineupName sets the value of the lineupName property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfostruct/lineupname
func (m_ MTRChannelClusterLineupInfoStruct) SetLineupName(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineupName:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfostruct/lineupinfotype
func (m_ MTRChannelClusterLineupInfoStruct) LineupInfoType() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("lineupInfoType"))
	return rv
}


// SetLineupInfoType sets the value of the lineupInfoType property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfostruct/lineupinfotype
func (m_ MTRChannelClusterLineupInfoStruct) SetLineupInfoType(value foundation.Number) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineupInfoType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfostruct/postalcode
func (m_ MTRChannelClusterLineupInfoStruct) PostalCode() string {
	rv := objc.Send[string](m_.ID, objc.Sel("postalCode"))
	return rv
}


// SetPostalCode sets the value of the postalCode property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfostruct/postalcode
func (m_ MTRChannelClusterLineupInfoStruct) SetPostalCode(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPostalCode:"), objc.String(value))
}



