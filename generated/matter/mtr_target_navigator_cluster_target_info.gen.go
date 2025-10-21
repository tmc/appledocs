// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRTargetNavigatorClusterTargetInfo] class.
var (
	MTRTargetNavigatorClusterTargetInfoClass     _MTRTargetNavigatorClusterTargetInfoClass
	MTRTargetNavigatorClusterTargetInfoClassOnce sync.Once
)

func getMTRTargetNavigatorClusterTargetInfoClass() _MTRTargetNavigatorClusterTargetInfoClass {
	MTRTargetNavigatorClusterTargetInfoClassOnce.Do(func() {
		MTRTargetNavigatorClusterTargetInfoClass = _MTRTargetNavigatorClusterTargetInfoClass{objc.GetClass("MTRTargetNavigatorClusterTargetInfo")}
	})
	return MTRTargetNavigatorClusterTargetInfoClass
}

type _MTRTargetNavigatorClusterTargetInfoClass struct {
	class objc.Class
}

// An interface definition for the [MTRTargetNavigatorClusterTargetInfo] class.
type IMTRTargetNavigatorClusterTargetInfo interface {
	IMTRTargetNavigatorClusterTargetInfoStruct
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetInfo
type MTRTargetNavigatorClusterTargetInfo struct {
	MTRTargetNavigatorClusterTargetInfoStruct
}

// MTRTargetNavigatorClusterTargetInfoFrom constructs a [MTRTargetNavigatorClusterTargetInfo] from an unsafe.Pointer.
func MTRTargetNavigatorClusterTargetInfoFrom(ptr unsafe.Pointer) MTRTargetNavigatorClusterTargetInfo {
	return MTRTargetNavigatorClusterTargetInfo{
		MTRTargetNavigatorClusterTargetInfoStruct: MTRTargetNavigatorClusterTargetInfoStructFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTargetNavigatorClusterTargetInfoClass) Alloc() MTRTargetNavigatorClusterTargetInfo {
	rv := objc.Send[MTRTargetNavigatorClusterTargetInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTargetNavigatorClusterTargetInfoClass) New() MTRTargetNavigatorClusterTargetInfo {
	rv := objc.Send[MTRTargetNavigatorClusterTargetInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTargetNavigatorClusterTargetInfo) Init() MTRTargetNavigatorClusterTargetInfo {
	rv := objc.Send[MTRTargetNavigatorClusterTargetInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTargetNavigatorClusterTargetInfo) Autorelease() MTRTargetNavigatorClusterTargetInfo {
	rv := objc.Send[MTRTargetNavigatorClusterTargetInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTargetNavigatorClusterTargetInfo creates a new MTRTargetNavigatorClusterTargetInfo instance.
func NewMTRTargetNavigatorClusterTargetInfo() MTRTargetNavigatorClusterTargetInfo {
	return getMTRTargetNavigatorClusterTargetInfoClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclustertargetinfo/identifier
func (m_ MTRTargetNavigatorClusterTargetInfo) Identifier() foundation.Number {
	rv := objc.Send[foundation.Number](m_.ID, objc.Sel("identifier"))
	return rv
}


// SetIdentifier sets the value of the identifier property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclustertargetinfo/identifier
func (m_ MTRTargetNavigatorClusterTargetInfo) SetIdentifier(value foundation.INumber) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIdentifier:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclustertargetinfo/name
func (m_ MTRTargetNavigatorClusterTargetInfo) Name() appkit.string {
	rv := objc.Send[appkit.string](m_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
//
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtargetnavigatorclustertargetinfo/name
func (m_ MTRTargetNavigatorClusterTargetInfo) SetName(value appkit.string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}



