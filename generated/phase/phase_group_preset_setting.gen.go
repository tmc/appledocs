// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [PHASEGroupPresetSetting] class.
type IPHASEGroupPresetSetting interface {
	objectivec.IObject
}

// Settings for group presets.
//
// This class defines playback speed and volume rates of change that an app can apply to groups. To create a group preset setting, instantiate an object of this type and pass it to the parameter of . For an example of preset settings, see .
//
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

// Alloc allocates a new instance without initialization.
func (pc _PHASEGroupPresetSettingClass) Alloc() PHASEGroupPresetSetting {
	rv := objc.Send[PHASEGroupPresetSetting](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a group preset setting.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting/init(gain:rate:gainCurveType:rateCurveType:)
func NewPHASEGroupPresetSettingWithGainRateGainCurveTypeRateCurveType(gain unsafe.Pointer, rate unsafe.Pointer, gainCurveType unsafe.Pointer, rateCurveType unsafe.Pointer) PHASEGroupPresetSetting {
	instance := getPHASEGroupPresetSettingClass().Alloc()
	rv := objc.Send[PHASEGroupPresetSetting](instance.ID, objc.Sel("initWithGain:rate:gainCurveType:rateCurveType:"), gain, rate, gainCurveType, rateCurveType)
	rv.Autorelease()
	return rv
}


// The volume of audio playback.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting/gain
func (p_ PHASEGroupPresetSetting) Gain() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("gain"))
	return rv
}

// A rate of change for the setting’s volume.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting/gainCurveType
func (p_ PHASEGroupPresetSetting) GainCurveType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("gainCurveType"))
	return rv
}

// The playback speed for audio.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting/rate
func (p_ PHASEGroupPresetSetting) Rate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("rate"))
	return rv
}

// A rate of change for the setting’s playback speed.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEGroupPresetSetting/rateCurveType
func (p_ PHASEGroupPresetSetting) RateCurveType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("rateCurveType"))
	return rv
}


