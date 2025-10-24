// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class PHASEGroupPresetSetting */


/* debug [class_header]: Header for PHASEGroupPresetSetting */
// The class instance for the [PHASEGroupPresetSetting] class.
var (
	PHASEGroupPresetSettingClass     _PHASEGroupPresetSettingClass
	PHASEGroupPresetSettingClassOnce sync.Once
)

func getPHASEGroupPresetSettingClass() _PHASEGroupPresetSettingClass {
	PHASEGroupPresetSettingClassOnce.Do(func() {
		PHASEGroupPresetSettingClass = _PHASEGroupPresetSettingClass{objc.GetClass("PHASEGroupPresetSetting")}
	})
	return PHASEGroupPresetSettingClass
}

type _PHASEGroupPresetSettingClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEGroupPresetSetting */
// An interface definition for the [PHASEGroupPresetSetting] class.
type IPHASEGroupPresetSetting interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PHASEGroupPresetSetting */
	// properties:
	Gain() float64
	GainCurveType() PHASECurveType
	Rate() float64
	RateCurveType() PHASECurveType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEGroupPresetSetting */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEGroupPresetSetting */
// Alloc allocates a new instance without initialization.
func (pc _PHASEGroupPresetSettingClass) Alloc() PHASEGroupPresetSetting {
	rv := objc.Send[PHASEGroupPresetSetting](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEGroupPresetSettingClass) New() PHASEGroupPresetSetting {
	rv := objc.Send[PHASEGroupPresetSetting](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEGroupPresetSetting) Init() PHASEGroupPresetSetting {
	rv := objc.Send[PHASEGroupPresetSetting](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEGroupPresetSetting) Autorelease() PHASEGroupPresetSetting {
	rv := objc.Send[PHASEGroupPresetSetting](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEGroupPresetSetting creates a new PHASEGroupPresetSetting instance.
func NewPHASEGroupPresetSetting() PHASEGroupPresetSetting {
	return getPHASEGroupPresetSettingClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEGroupPresetSetting */
// Settings for group presets.
//
// This class defines playback speed and volume rates of change that an app can apply to groups. To create a group preset setting, instantiate an object of this type and pass it to the parameter of . For an example of preset settings, see .


// Settings for group presets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting
type PHASEGroupPresetSetting struct {
	objectivec.Object
}

// PHASEGroupPresetSettingFrom constructs a [PHASEGroupPresetSetting] from an unsafe.Pointer.
//
// Settings for group presets.
func PHASEGroupPresetSettingFrom(ptr unsafe.Pointer) PHASEGroupPresetSetting {
	return PHASEGroupPresetSetting{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEGroupPresetSetting */

// Creates a group preset setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting/init(gain:rate:gainCurveType:rateCurveType:)
func NewPHASEGroupPresetSettingWithGainRateGainCurveTypeRateCurveType(gain float64, rate float64, gainCurveType PHASECurveType, rateCurveType PHASECurveType) PHASEGroupPresetSetting {
	instance := getPHASEGroupPresetSettingClass().Alloc()
	rv := objc.Send[PHASEGroupPresetSetting](instance.ID, objc.Sel("initWithGain:rate:gainCurveType:rateCurveType:"), gain, rate, gainCurveType, rateCurveType)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEGroupPresetSettingWithGainRateGainCurveTypeRateCurveType */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEGroupPresetSetting */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEGroupPresetSetting */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEGroupPresetSetting */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEGroupPresetSetting */

// The volume of audio playback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting/gain
func (p_ PHASEGroupPresetSetting) Gain() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("gain"))
	return rv
}/* debug [instance_properties/getter]: gain */


// A rate of change for the setting’s volume.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting/gainCurveType
func (p_ PHASEGroupPresetSetting) GainCurveType() PHASECurveType {
	rv := objc.Send[PHASECurveType](p_.ID, objc.Sel("gainCurveType"))
	return rv
}/* debug [instance_properties/getter]: gainCurveType */


// The playback speed for audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting/rate
func (p_ PHASEGroupPresetSetting) Rate() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("rate"))
	return rv
}/* debug [instance_properties/getter]: rate */


// A rate of change for the setting’s playback speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting/rateCurveType
func (p_ PHASEGroupPresetSetting) RateCurveType() PHASECurveType {
	rv := objc.Send[PHASECurveType](p_.ID, objc.Sel("rateCurveType"))
	return rv
}/* debug [instance_properties/getter]: rateCurveType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEGroupPresetSetting */


