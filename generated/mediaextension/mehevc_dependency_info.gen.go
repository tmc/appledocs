// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	SyncSampleNALUnitType() int16 /* not a class type */
	SetSyncSampleNALUnitType(value int16 /* not a class type */)
	ConstraintIndicatorFlags() objc.IObject /* cross-framework: Data */
	SetConstraintIndicatorFlags(value objc.IObject /* cross-framework: Data */)
	HasStepwiseTemporalSubLayerAccess() bool
	SetHasStepwiseTemporalSubLayerAccess(value bool)
	HasTemporalSubLayerAccess() bool
	SetHasTemporalSubLayerAccess(value bool)
	LevelIndex() unsafe.Pointer
	SetLevelIndex(value unsafe.Pointer)
	ProfileCompatibilityFlags() objc.IObject /* cross-framework: Data */
	SetProfileCompatibilityFlags(value objc.IObject /* cross-framework: Data */)
	ProfileIndex() unsafe.Pointer
	SetProfileIndex(value unsafe.Pointer)
	ProfileSpace() unsafe.Pointer
	SetProfileSpace(value unsafe.Pointer)
	TemporalLevel() unsafe.Pointer
	SetTemporalLevel(value unsafe.Pointer)
	TierFlag() unsafe.Pointer
	SetTierFlag(value unsafe.Pointer)
	// methods:
}

// An object that provides information about the HEVC dependency attributes of a sample.


// An object that provides information about the HEVC dependency attributes of a sample.
//
// [Full Topic]
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



// The NAL unit type for HEVC sync sample groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/syncSampleNALUnitType
func (m_ MEHEVCDependencyInfo) SyncSampleNALUnitType() int16 /* not a class type */ {
	rv := objc.Send[int16](m_.ID, objc.Sel("syncSampleNALUnitType"))
	return rv
}


// The NAL unit type for HEVC sync sample groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/syncSampleNALUnitType
func (m_ MEHEVCDependencyInfo) SetSyncSampleNALUnitType(value int16 /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSyncSampleNALUnitType:"), value)
}


// The HEVC constraint indicator flags (6 bytes), if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/constraintindicatorflags
func (m_ MEHEVCDependencyInfo) ConstraintIndicatorFlags() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("constraintIndicatorFlags"))
	return rv
}


// The HEVC constraint indicator flags (6 bytes), if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/constraintindicatorflags
func (m_ MEHEVCDependencyInfo) SetConstraintIndicatorFlags(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConstraintIndicatorFlags:"), value)
}


// A Boolean value that indicates if the sample has an HEVC stepwise temporal sublayer access (STSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/hasstepwisetemporalsublayeraccess
func (m_ MEHEVCDependencyInfo) HasStepwiseTemporalSubLayerAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasStepwiseTemporalSubLayerAccess"))
	return rv
}


// A Boolean value that indicates if the sample has an HEVC stepwise temporal sublayer access (STSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/hasstepwisetemporalsublayeraccess
func (m_ MEHEVCDependencyInfo) SetHasStepwiseTemporalSubLayerAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasStepwiseTemporalSubLayerAccess:"), value)
}


// A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/hastemporalsublayeraccess
func (m_ MEHEVCDependencyInfo) HasTemporalSubLayerAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasTemporalSubLayerAccess"))
	return rv
}


// A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/hastemporalsublayeraccess
func (m_ MEHEVCDependencyInfo) SetHasTemporalSubLayerAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasTemporalSubLayerAccess:"), value)
}


// The HEVC level index, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/levelindex
func (m_ MEHEVCDependencyInfo) LevelIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("levelIndex"))
	return rv
}


// The HEVC level index, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/levelindex
func (m_ MEHEVCDependencyInfo) SetLevelIndex(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLevelIndex:"), value)
}


// The HEVC profile compatibility flags (4 bytes), if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/profilecompatibilityflags
func (m_ MEHEVCDependencyInfo) ProfileCompatibilityFlags() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("profileCompatibilityFlags"))
	return rv
}


// The HEVC profile compatibility flags (4 bytes), if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/profilecompatibilityflags
func (m_ MEHEVCDependencyInfo) SetProfileCompatibilityFlags(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileCompatibilityFlags:"), value)
}


// The HEVC profile index, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/profileindex
func (m_ MEHEVCDependencyInfo) ProfileIndex() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("profileIndex"))
	return rv
}


// The HEVC profile index, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/profileindex
func (m_ MEHEVCDependencyInfo) SetProfileIndex(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileIndex:"), value)
}


// The HEVC profile space, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/profilespace
func (m_ MEHEVCDependencyInfo) ProfileSpace() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("profileSpace"))
	return rv
}


// The HEVC profile space, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/profilespace
func (m_ MEHEVCDependencyInfo) SetProfileSpace(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileSpace:"), value)
}


// The HEVC temporal level, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/temporallevel
func (m_ MEHEVCDependencyInfo) TemporalLevel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("temporalLevel"))
	return rv
}


// The HEVC temporal level, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/temporallevel
func (m_ MEHEVCDependencyInfo) SetTemporalLevel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemporalLevel:"), value)
}


// The HEVC tier level flag, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/tierflag
func (m_ MEHEVCDependencyInfo) TierFlag() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("tierFlag"))
	return rv
}


// The HEVC tier level flag, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/tierflag
func (m_ MEHEVCDependencyInfo) SetTierFlag(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTierFlag:"), value)
}



