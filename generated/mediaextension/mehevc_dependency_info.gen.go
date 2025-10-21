// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MEHEVCDependencyInfo] class.
var (
	MEHEVCDependencyInfoClass     _MEHEVCDependencyInfoClass
	MEHEVCDependencyInfoClassOnce sync.Once
)

func getMEHEVCDependencyInfoClass() _MEHEVCDependencyInfoClass {
	MEHEVCDependencyInfoClassOnce.Do(func() {
		MEHEVCDependencyInfoClass = _MEHEVCDependencyInfoClass{objc.GetClass("MEHEVCDependencyInfo")}
	})
	return MEHEVCDependencyInfoClass
}

type _MEHEVCDependencyInfoClass struct {
	class objc.Class
}

// An interface definition for the [MEHEVCDependencyInfo] class.
type IMEHEVCDependencyInfo interface {
	objectivec.IObject
}

// An object that provides information about the HEVC dependency attributes of a sample.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo
type MEHEVCDependencyInfo struct {
	objectivec.Object
}

// MEHEVCDependencyInfoFrom constructs a [MEHEVCDependencyInfo] from an unsafe.Pointer.
//
// An object that provides information about the HEVC dependency attributes of a sample.
func MEHEVCDependencyInfoFrom(ptr unsafe.Pointer) MEHEVCDependencyInfo {
	return MEHEVCDependencyInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEHEVCDependencyInfoClass) Alloc() MEHEVCDependencyInfo {
	rv := objc.Send[MEHEVCDependencyInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEHEVCDependencyInfoClass) New() MEHEVCDependencyInfo {
	rv := objc.Send[MEHEVCDependencyInfo](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEHEVCDependencyInfo) Init() MEHEVCDependencyInfo {
	rv := objc.Send[MEHEVCDependencyInfo](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEHEVCDependencyInfo) Autorelease() MEHEVCDependencyInfo {
	rv := objc.Send[MEHEVCDependencyInfo](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEHEVCDependencyInfo creates a new MEHEVCDependencyInfo instance.
func NewMEHEVCDependencyInfo() MEHEVCDependencyInfo {
	return getMEHEVCDependencyInfoClass().New()
}


// The HEVC constraint indicator flags (6 bytes), if available.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/constraintIndicatorFlags
func (m_ MEHEVCDependencyInfo) ConstraintIndicatorFlags() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("constraintIndicatorFlags"))
	return rv
}


// SetConstraintIndicatorFlags sets the value of the constraintIndicatorFlags property.
// The HEVC constraint indicator flags (6 bytes), if available.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/constraintIndicatorFlags
func (m_ MEHEVCDependencyInfo) SetConstraintIndicatorFlags(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConstraintIndicatorFlags:"), value)
}
// A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/hasTemporalSubLayerAccess
func (m_ MEHEVCDependencyInfo) TemporalSubLayerAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("temporalSubLayerAccess"))
	return rv
}


// SetTemporalSubLayerAccess sets the value of the temporalSubLayerAccess property.
// A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/hasTemporalSubLayerAccess
func (m_ MEHEVCDependencyInfo) SetTemporalSubLayerAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemporalSubLayerAccess:"), value)
}
// The HEVC level index, if available.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/levelIndex
func (m_ MEHEVCDependencyInfo) LevelIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("levelIndex"))
	return rv
}


// SetLevelIndex sets the value of the levelIndex property.
// The HEVC level index, if available.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/levelIndex
func (m_ MEHEVCDependencyInfo) SetLevelIndex(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLevelIndex:"), value)
}
// The HEVC profile space, if available.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileSpace
func (m_ MEHEVCDependencyInfo) ProfileSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("profileSpace"))
	return rv
}


// SetProfileSpace sets the value of the profileSpace property.
// The HEVC profile space, if available.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileSpace
func (m_ MEHEVCDependencyInfo) SetProfileSpace(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileSpace:"), value)
}
// The NAL unit type for HEVC sync sample groups.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/syncSampleNALUnitType
func (m_ MEHEVCDependencyInfo) SyncSampleNALUnitType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("syncSampleNALUnitType"))
	return rv
}


// SetSyncSampleNALUnitType sets the value of the syncSampleNALUnitType property.
// The NAL unit type for HEVC sync sample groups.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/syncSampleNALUnitType
func (m_ MEHEVCDependencyInfo) SetSyncSampleNALUnitType(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSyncSampleNALUnitType:"), value)
}
// The HEVC temporal level, if available.
//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/temporalLevel
func (m_ MEHEVCDependencyInfo) TemporalLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("temporalLevel"))
	return rv
}


// SetTemporalLevel sets the value of the temporalLevel property.
// The HEVC temporal level, if available.

//
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/temporalLevel
func (m_ MEHEVCDependencyInfo) SetTemporalLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemporalLevel:"), value)
}


