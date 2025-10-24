// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	LineupInfoType() objc.IObject /* cross-framework: NSNumber */
	SetLineupInfoType(value objc.IObject /* cross-framework: NSNumber */)
	LineupName() objc.IObject /* cross-framework: NSString */
	SetLineupName(value objc.IObject /* cross-framework: NSString */)
	OperatorName() objc.IObject /* cross-framework: NSString */
	SetOperatorName(value objc.IObject /* cross-framework: NSString */)
	PostalCode() objc.IObject /* cross-framework: NSString */
	SetPostalCode(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// [Full Topic]
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

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/lineupinfotype
func (m_ MTRChannelClusterLineupInfo) LineupInfoType() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("lineupInfoType"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/lineupinfotype
func (m_ MTRChannelClusterLineupInfo) SetLineupInfoType(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineupInfoType:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/lineupname
func (m_ MTRChannelClusterLineupInfo) LineupName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("lineupName"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/lineupname
func (m_ MTRChannelClusterLineupInfo) SetLineupName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLineupName:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/operatorname
func (m_ MTRChannelClusterLineupInfo) OperatorName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("operatorName"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/operatorname
func (m_ MTRChannelClusterLineupInfo) SetOperatorName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOperatorName:"), value)
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/postalcode
func (m_ MTRChannelClusterLineupInfo) PostalCode() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("postalCode"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterlineupinfo/postalcode
func (m_ MTRChannelClusterLineupInfo) SetPostalCode(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPostalCode:"), value)
}
