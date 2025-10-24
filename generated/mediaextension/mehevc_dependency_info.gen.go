// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEHEVCDependencyInfo */


/* debug [class_header]: Header for MEHEVCDependencyInfo */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEHEVCDependencyInfo */
// An interface definition for the [MEHEVCDependencyInfo] class.
type IMEHEVCDependencyInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEHEVCDependencyInfo */
	// properties:
	ConstraintIndicatorFlags() objc.IObject /* cross-framework: NSData */
	SetConstraintIndicatorFlags(value objc.IObject /* cross-framework: NSData */)
	StepwiseTemporalSubLayerAccess() bool
	SetStepwiseTemporalSubLayerAccess(value bool)
	TemporalSubLayerAccess() bool
	SetTemporalSubLayerAccess(value bool)
	LevelIndex() int16 /* not a class type */
	SetLevelIndex(value int16 /* not a class type */)
	ProfileCompatibilityFlags() objc.IObject /* cross-framework: NSData */
	SetProfileCompatibilityFlags(value objc.IObject /* cross-framework: NSData */)
	ProfileIndex() int16 /* not a class type */
	SetProfileIndex(value int16 /* not a class type */)
	ProfileSpace() int16 /* not a class type */
	SetProfileSpace(value int16 /* not a class type */)
	SyncSampleNALUnitType() int16 /* not a class type */
	SetSyncSampleNALUnitType(value int16 /* not a class type */)
	TemporalLevel() int16 /* not a class type */
	SetTemporalLevel(value int16 /* not a class type */)
	TierFlag() int16 /* not a class type */
	SetTierFlag(value int16 /* not a class type */)
	HasStepwiseTemporalSubLayerAccess() bool
	SetHasStepwiseTemporalSubLayerAccess(value bool)
	HasTemporalSubLayerAccess() bool
	SetHasTemporalSubLayerAccess(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEHEVCDependencyInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEHEVCDependencyInfo */
// Alloc allocates a new instance without initialization.
func (mc _MEHEVCDependencyInfoClass) Alloc() MEHEVCDependencyInfo {
	rv := objc.Send[MEHEVCDependencyInfo](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEHEVCDependencyInfo */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEHEVCDependencyInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEHEVCDependencyInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEHEVCDependencyInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEHEVCDependencyInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEHEVCDependencyInfo */

// The HEVC constraint indicator flags (6 bytes), if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/constraintIndicatorFlags
func (m_ MEHEVCDependencyInfo) ConstraintIndicatorFlags() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("constraintIndicatorFlags"))
	return rv
}/* debug [instance_properties/getter]: constraintIndicatorFlags */


// The HEVC constraint indicator flags (6 bytes), if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/constraintIndicatorFlags
func (m_ MEHEVCDependencyInfo) SetConstraintIndicatorFlags(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setConstraintIndicatorFlags:"), value)
}/* debug [instance_properties/setter]: constraintIndicatorFlags */


// A Boolean value that indicates if the sample has an HEVC stepwise temporal sublayer access (STSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/hasStepwiseTemporalSubLayerAccess
func (m_ MEHEVCDependencyInfo) StepwiseTemporalSubLayerAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("stepwiseTemporalSubLayerAccess"))
	return rv
}/* debug [instance_properties/getter]: stepwiseTemporalSubLayerAccess */


// A Boolean value that indicates if the sample has an HEVC stepwise temporal sublayer access (STSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/hasStepwiseTemporalSubLayerAccess
func (m_ MEHEVCDependencyInfo) SetStepwiseTemporalSubLayerAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStepwiseTemporalSubLayerAccess:"), value)
}/* debug [instance_properties/setter]: stepwiseTemporalSubLayerAccess */


// A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/hasTemporalSubLayerAccess
func (m_ MEHEVCDependencyInfo) TemporalSubLayerAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("temporalSubLayerAccess"))
	return rv
}/* debug [instance_properties/getter]: temporalSubLayerAccess */


// A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/hasTemporalSubLayerAccess
func (m_ MEHEVCDependencyInfo) SetTemporalSubLayerAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemporalSubLayerAccess:"), value)
}/* debug [instance_properties/setter]: temporalSubLayerAccess */


// The HEVC level index, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/levelIndex
func (m_ MEHEVCDependencyInfo) LevelIndex() int16 /* not a class type */ {
	rv := objc.Send[int16](m_.ID, objc.Sel("levelIndex"))
	return rv
}/* debug [instance_properties/getter]: levelIndex */


// The HEVC level index, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/levelIndex
func (m_ MEHEVCDependencyInfo) SetLevelIndex(value int16 /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLevelIndex:"), value)
}/* debug [instance_properties/setter]: levelIndex */


// The HEVC profile compatibility flags (4 bytes), if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileCompatibilityFlags
func (m_ MEHEVCDependencyInfo) ProfileCompatibilityFlags() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("profileCompatibilityFlags"))
	return rv
}/* debug [instance_properties/getter]: profileCompatibilityFlags */


// The HEVC profile compatibility flags (4 bytes), if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileCompatibilityFlags
func (m_ MEHEVCDependencyInfo) SetProfileCompatibilityFlags(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileCompatibilityFlags:"), value)
}/* debug [instance_properties/setter]: profileCompatibilityFlags */


// The HEVC profile index, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileIndex
func (m_ MEHEVCDependencyInfo) ProfileIndex() int16 /* not a class type */ {
	rv := objc.Send[int16](m_.ID, objc.Sel("profileIndex"))
	return rv
}/* debug [instance_properties/getter]: profileIndex */


// The HEVC profile index, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileIndex
func (m_ MEHEVCDependencyInfo) SetProfileIndex(value int16 /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileIndex:"), value)
}/* debug [instance_properties/setter]: profileIndex */


// The HEVC profile space, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileSpace
func (m_ MEHEVCDependencyInfo) ProfileSpace() int16 /* not a class type */ {
	rv := objc.Send[int16](m_.ID, objc.Sel("profileSpace"))
	return rv
}/* debug [instance_properties/getter]: profileSpace */


// The HEVC profile space, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/profileSpace
func (m_ MEHEVCDependencyInfo) SetProfileSpace(value int16 /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProfileSpace:"), value)
}/* debug [instance_properties/setter]: profileSpace */


// The NAL unit type for HEVC sync sample groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/syncSampleNALUnitType
func (m_ MEHEVCDependencyInfo) SyncSampleNALUnitType() int16 /* not a class type */ {
	rv := objc.Send[int16](m_.ID, objc.Sel("syncSampleNALUnitType"))
	return rv
}/* debug [instance_properties/getter]: syncSampleNALUnitType */


// The NAL unit type for HEVC sync sample groups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/syncSampleNALUnitType
func (m_ MEHEVCDependencyInfo) SetSyncSampleNALUnitType(value int16 /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSyncSampleNALUnitType:"), value)
}/* debug [instance_properties/setter]: syncSampleNALUnitType */


// The HEVC temporal level, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/temporalLevel
func (m_ MEHEVCDependencyInfo) TemporalLevel() int16 /* not a class type */ {
	rv := objc.Send[int16](m_.ID, objc.Sel("temporalLevel"))
	return rv
}/* debug [instance_properties/getter]: temporalLevel */


// The HEVC temporal level, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/temporalLevel
func (m_ MEHEVCDependencyInfo) SetTemporalLevel(value int16 /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTemporalLevel:"), value)
}/* debug [instance_properties/setter]: temporalLevel */


// The HEVC tier level flag, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/tierFlag
func (m_ MEHEVCDependencyInfo) TierFlag() int16 /* not a class type */ {
	rv := objc.Send[int16](m_.ID, objc.Sel("tierFlag"))
	return rv
}/* debug [instance_properties/getter]: tierFlag */


// The HEVC tier level flag, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEHEVCDependencyInfo/tierFlag
func (m_ MEHEVCDependencyInfo) SetTierFlag(value int16 /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTierFlag:"), value)
}/* debug [instance_properties/setter]: tierFlag */


// A Boolean value that indicates if the sample has an HEVC stepwise temporal sublayer access (STSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/hasstepwisetemporalsublayeraccess
func (m_ MEHEVCDependencyInfo) HasStepwiseTemporalSubLayerAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasStepwiseTemporalSubLayerAccess"))
	return rv
}/* debug [instance_properties/getter]: hasStepwiseTemporalSubLayerAccess */


// A Boolean value that indicates if the sample has an HEVC stepwise temporal sublayer access (STSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/hasstepwisetemporalsublayeraccess
func (m_ MEHEVCDependencyInfo) SetHasStepwiseTemporalSubLayerAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasStepwiseTemporalSubLayerAccess:"), value)
}/* debug [instance_properties/setter]: hasStepwiseTemporalSubLayerAccess */


// A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/hastemporalsublayeraccess
func (m_ MEHEVCDependencyInfo) HasTemporalSubLayerAccess() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("hasTemporalSubLayerAccess"))
	return rv
}/* debug [instance_properties/getter]: hasTemporalSubLayerAccess */


// A Boolean value that indicates if the sample has an HEVC temporal sublayer access (TSA) picture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mediaextension/mehevcdependencyinfo/hastemporalsublayeraccess
func (m_ MEHEVCDependencyInfo) SetHasTemporalSubLayerAccess(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setHasTemporalSubLayerAccess:"), value)
}/* debug [instance_properties/setter]: hasTemporalSubLayerAccess */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEHEVCDependencyInfo */



