// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [INSetClimateSettingsInCarIntent] class.
var (
	INSetClimateSettingsInCarIntentClass     _INSetClimateSettingsInCarIntentClass
	INSetClimateSettingsInCarIntentClassOnce sync.Once
)

func getINSetClimateSettingsInCarIntentClass() _INSetClimateSettingsInCarIntentClass {
	INSetClimateSettingsInCarIntentClassOnce.Do(func() {
		INSetClimateSettingsInCarIntentClass = _INSetClimateSettingsInCarIntentClass{objc.GetClass("INSetClimateSettingsInCarIntent")}
	})
	return INSetClimateSettingsInCarIntentClass
}

type _INSetClimateSettingsInCarIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSetClimateSettingsInCarIntent] class.
type IINSetClimateSettingsInCarIntent interface {
	IINIntent
	AirCirculationMode() unsafe.Pointer
	SetAirCirculationMode(value unsafe.Pointer)
	CarName() INSpeakableString
	SetCarName(value INSpeakableString)
	ClimateZone() unsafe.Pointer
	SetClimateZone(value unsafe.Pointer)
	EnableAirConditioner() bool
	SetEnableAirConditioner(value bool)
	EnableAutoMode() bool
	SetEnableAutoMode(value bool)
	EnableClimateControl() bool
	SetEnableClimateControl(value bool)
	EnableFan() bool
	SetEnableFan(value bool)
	FanSpeedIndex() int
	SetFanSpeedIndex(value int)
	FanSpeedPercentage() float64
	SetFanSpeedPercentage(value float64)
	RelativeFanSpeedSetting() unsafe.Pointer
	SetRelativeFanSpeedSetting(value unsafe.Pointer)
	RelativeTemperatureSetting() unsafe.Pointer
	SetRelativeTemperatureSetting(value unsafe.Pointer)
	Temperature() foundation.UnitTemperature
	SetTemperature(value foundation.IUnitTemperature)
}

// A request to change the climate settings in a CarPlay-enabled vehicle.
//
// Automotive vendors can add support for this intent to an Intents extension that they ship with their automotive apps. When the user asks Siri to change a setting related to the vehicle’s climate control, Siri creates an object and delivers it to the app’s Intents extension. You use the intent to identify which setting the user wants to change. This class contains properties for multiple types of climate settings, but a given instance of this class contains changes for only one system at a time. When resolving and confirming the parameters of this intent, use the specified properties to modify your vehicle’s settings. Assume no changes for other settings. If your vehicle doesn’t support a particular setting, offer a reasonable fallback setting and ask the user for confirmation. The object that handles this intent must adopt the protocol. Use this intent object to identify which setting changed and to create an object indicating the results of changing the climate settings.


// A request to change the climate settings in a CarPlay-enabled vehicle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSetClimateSettingsInCarIntent
type INSetClimateSettingsInCarIntent struct {
	INIntent
}

// INSetClimateSettingsInCarIntentFrom constructs a [INSetClimateSettingsInCarIntent] from an unsafe.Pointer.
//
// A request to change the climate settings in a CarPlay-enabled vehicle.
func INSetClimateSettingsInCarIntentFrom(ptr unsafe.Pointer) INSetClimateSettingsInCarIntent {
	return INSetClimateSettingsInCarIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSetClimateSettingsInCarIntentClass) Alloc() INSetClimateSettingsInCarIntent {
	rv := objc.Send[INSetClimateSettingsInCarIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSetClimateSettingsInCarIntentClass) New() INSetClimateSettingsInCarIntent {
	rv := objc.Send[INSetClimateSettingsInCarIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSetClimateSettingsInCarIntent) Init() INSetClimateSettingsInCarIntent {
	rv := objc.Send[INSetClimateSettingsInCarIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSetClimateSettingsInCarIntent) Autorelease() INSetClimateSettingsInCarIntent {
	rv := objc.Send[INSetClimateSettingsInCarIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSetClimateSettingsInCarIntent creates a new INSetClimateSettingsInCarIntent instance.
func NewINSetClimateSettingsInCarIntent() INSetClimateSettingsInCarIntent {
	return getINSetClimateSettingsInCarIntentClass().New()
}



// The air circulation mode for the climate control system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/aircirculationmode
func (i_ INSetClimateSettingsInCarIntent) AirCirculationMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("airCirculationMode"))
	return rv
}


// The air circulation mode for the climate control system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/aircirculationmode
func (i_ INSetClimateSettingsInCarIntent) SetAirCirculationMode(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setAirCirculationMode:"), value)
}


// A name that identifies the user’s car.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/carname
func (i_ INSetClimateSettingsInCarIntent) CarName() INSpeakableString {
	rv := objc.Send[INSpeakableString](i_.ID, objc.Sel("carName"))
	return rv
}


// A name that identifies the user’s car.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/carname
func (i_ INSetClimateSettingsInCarIntent) SetCarName(value INSpeakableString) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCarName:"), value)
}


// The seat position to have its climate settings modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/climatezone
func (i_ INSetClimateSettingsInCarIntent) ClimateZone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("climateZone"))
	return rv
}


// The seat position to have its climate settings modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/climatezone
func (i_ INSetClimateSettingsInCarIntent) SetClimateZone(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setClimateZone:"), value)
}


// A Boolean value indicating whether to turn on the air conditioner system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/enableairconditioner-9q3dr
func (i_ INSetClimateSettingsInCarIntent) EnableAirConditioner() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("enableAirConditioner"))
	return rv
}


// A Boolean value indicating whether to turn on the air conditioner system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/enableairconditioner-9q3dr
func (i_ INSetClimateSettingsInCarIntent) SetEnableAirConditioner(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEnableAirConditioner:"), value)
}


// A Boolean value indicating whether to enable automatic mode for the climate control system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/enableautomode-31bzq
func (i_ INSetClimateSettingsInCarIntent) EnableAutoMode() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("enableAutoMode"))
	return rv
}


// A Boolean value indicating whether to enable automatic mode for the climate control system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/enableautomode-31bzq
func (i_ INSetClimateSettingsInCarIntent) SetEnableAutoMode(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEnableAutoMode:"), value)
}


// A Boolean value indicating whether to turn on the climate control system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/enableclimatecontrol-4y9iz
func (i_ INSetClimateSettingsInCarIntent) EnableClimateControl() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("enableClimateControl"))
	return rv
}


// A Boolean value indicating whether to turn on the climate control system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/enableclimatecontrol-4y9iz
func (i_ INSetClimateSettingsInCarIntent) SetEnableClimateControl(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEnableClimateControl:"), value)
}


// A Boolean value indicating whether to turn on the cabin fan system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/enablefan-5srs9
func (i_ INSetClimateSettingsInCarIntent) EnableFan() bool {
	rv := objc.Send[bool](i_.ID, objc.Sel("enableFan"))
	return rv
}


// A Boolean value indicating whether to turn on the cabin fan system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/enablefan-5srs9
func (i_ INSetClimateSettingsInCarIntent) SetEnableFan(value bool) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setEnableFan:"), value)
}


// An integer value indicating the desired fan speed position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/fanspeedindex-wz49
func (i_ INSetClimateSettingsInCarIntent) FanSpeedIndex() int {
	rv := objc.Send[int](i_.ID, objc.Sel("fanSpeedIndex"))
	return rv
}


// An integer value indicating the desired fan speed position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/fanspeedindex-wz49
func (i_ INSetClimateSettingsInCarIntent) SetFanSpeedIndex(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFanSpeedIndex:"), value)
}


// A floating-point value indicating the requested fan speed specified as a percentage of the maximum speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/fanspeedpercentage-7i2hq
func (i_ INSetClimateSettingsInCarIntent) FanSpeedPercentage() float64 {
	rv := objc.Send[float64](i_.ID, objc.Sel("fanSpeedPercentage"))
	return rv
}


// A floating-point value indicating the requested fan speed specified as a percentage of the maximum speed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/fanspeedpercentage-7i2hq
func (i_ INSetClimateSettingsInCarIntent) SetFanSpeedPercentage(value float64) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFanSpeedPercentage:"), value)
}


// A relative fan speed setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/relativefanspeedsetting
func (i_ INSetClimateSettingsInCarIntent) RelativeFanSpeedSetting() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("relativeFanSpeedSetting"))
	return rv
}


// A relative fan speed setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/relativefanspeedsetting
func (i_ INSetClimateSettingsInCarIntent) SetRelativeFanSpeedSetting(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRelativeFanSpeedSetting:"), value)
}


// A relative temperature setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/relativetemperaturesetting
func (i_ INSetClimateSettingsInCarIntent) RelativeTemperatureSetting() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("relativeTemperatureSetting"))
	return rv
}


// A relative temperature setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/relativetemperaturesetting
func (i_ INSetClimateSettingsInCarIntent) SetRelativeTemperatureSetting(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setRelativeTemperatureSetting:"), value)
}


// The specific temperature to set for the climate control system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/temperature
func (i_ INSetClimateSettingsInCarIntent) Temperature() foundation.UnitTemperature {
	rv := objc.Send[foundation.UnitTemperature](i_.ID, objc.Sel("temperature"))
	return rv
}


// The specific temperature to set for the climate control system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/intents/insetclimatesettingsincarintent/temperature
func (i_ INSetClimateSettingsInCarIntent) SetTemperature(value foundation.IUnitTemperature) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTemperature:"), value)
}



