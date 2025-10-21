// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [HKUnit] class.
var (
	HKUnitClass     _HKUnitClass
	HKUnitClassOnce sync.Once
)

func getHKUnitClass() _HKUnitClass {
	HKUnitClassOnce.Do(func() {
		HKUnitClass = _HKUnitClass{objc.GetClass("HKUnit")}
	})
	return HKUnitClass
}

type _HKUnitClass struct {
	class objc.Class
}

// An interface definition for the [HKUnit] class.
type IHKUnit interface {
	objectivec.IObject
	IsNull() bool
	ReciprocalUnit() unsafe.Pointer
	UnitDividedByUnit(unit unsafe.Pointer) unsafe.Pointer
	UnitMultipliedByUnit(unit unsafe.Pointer) unsafe.Pointer
	UnitRaisedToPower(power int) unsafe.Pointer
}

// A class for managing the units of measure within HealthKit.
//
// The unit class supports most standard SI units (meters, seconds, and grams), SI units with prefixes (centimeters, milliseconds and kilograms) and equivalent non-SI units (feet, minutes, and pounds). HealthKit also supports creating complex units by mathematically combining existing units. You use units when working with HealthKit quantities. Quantities store both the value (as a data type) and its corresponding unit. You can then request the value from the quantity in any compatible units. For more information on working with quantities, see .
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit
type HKUnit struct {
	objectivec.Object
}

// HKUnitFrom constructs a [HKUnit] from an unsafe.Pointer.
//
// A class for managing the units of measure within HealthKit.
func HKUnitFrom(ptr unsafe.Pointer) HKUnit {
	return HKUnit{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HKUnitClass) Alloc() HKUnit {
	rv := objc.Send[HKUnit](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HKUnitClass) New() HKUnit {
	rv := objc.Send[HKUnit](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ HKUnit) Init() HKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ HKUnit) Autorelease() HKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHKUnit creates a new HKUnit instance.
func NewHKUnit() HKUnit {
	return getHKUnitClass().New()
}




// Converts an energy formatter enumeration value into a corresponding HealthKit unit object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-1j1pq
func NewHKUnitFromEnergyFormatterUnit(energyFormatterUnit unsafe.Pointer) HKUnit {
	rv := objc.Send[HKUnit](objc.ID(getHKUnitClass().class), objc.Sel("unitFromEnergyFormatterUnit:"), energyFormatterUnit)
	return rv
}



// Converts a length formatter enumeration value into a corresponding HealthKit object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-55e1u
func NewHKUnitFromLengthFormatterUnit(lengthFormatterUnit unsafe.Pointer) HKUnit {
	rv := objc.Send[HKUnit](objc.ID(getHKUnitClass().class), objc.Sel("unitFromLengthFormatterUnit:"), lengthFormatterUnit)
	return rv
}



// Converts a mass formatter enumeration value into a corresponding HealthKit unit object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-7h2li
func NewHKUnitFromMassFormatterUnit(massFormatterUnit unsafe.Pointer) HKUnit {
	rv := objc.Send[HKUnit](objc.ID(getHKUnitClass().class), objc.Sel("unitFromMassFormatterUnit:"), massFormatterUnit)
	return rv
}



// Returns the unit instance described by the provided string.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-9qont
func NewHKUnitFromString(string_ string) HKUnit {
	rv := objc.Send[HKUnit](objc.ID(getHKUnitClass().class), objc.Sel("unitFromString:"), objc.String(string_))
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/appleEffortScore()
func (hc _HKUnitClass) AppleEffortScoreUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("appleEffortScoreUnit"))
	return rv
}

// Returns a HealthKit unit for measuring pressure in atmospheres.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/atmosphere()
func (hc _HKUnitClass) AtmosphereUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("atmosphereUnit"))
	return rv
}

// Returns a HealthKit unit for measuring energy in calories.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/calorie()
func (hc _HKUnitClass) CalorieUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("calorieUnit"))
	return rv
}

// Returns a HealthKit unit for measuring pressure in centimeters of water.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/centimeterOfWater()
func (hc _HKUnitClass) CentimeterOfWaterUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("centimeterOfWaterUnit"))
	return rv
}

// Returns a HealthKit unit for measuring counts.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/count()
func (hc _HKUnitClass) CountUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("countUnit"))
	return rv
}

// Returns a HealthKit unit for measuring volume in imperial cups.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/cupImperial()
func (hc _HKUnitClass) CupImperialUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("cupImperialUnit"))
	return rv
}

// Returns a HealthKit unit for measuring volume in US cups.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/cupUS()
func (hc _HKUnitClass) CupUSUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("cupUSUnit"))
	return rv
}

// Returns a HealthKit unit for measuring time in days.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/day()
func (hc _HKUnitClass) DayUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("dayUnit"))
	return rv
}

// Returns a HealthKit unit for measuring the difference between the local pressure and the ambient atmospheric pressure caused by sound.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/decibelAWeightedSoundPressureLevel()
func (hc _HKUnitClass) DecibelAWeightedSoundPressureLevelUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("decibelAWeightedSoundPressureLevelUnit"))
	return rv
}

// Returns a HealthKit unit for measuring the intensity of a sound.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/decibelHearingLevel()
func (hc _HKUnitClass) DecibelHearingLevelUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("decibelHearingLevelUnit"))
	return rv
}

// Returns a HealthKit unit for measuring angles using degrees.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/degreeAngle()
func (hc _HKUnitClass) DegreeAngleUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("degreeAngleUnit"))
	return rv
}

// Returns a HealthKit unit for measuring temperature in degrees Celsius.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/degreeCelsius()
func (hc _HKUnitClass) DegreeCelsiusUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("degreeCelsiusUnit"))
	return rv
}

// Returns a HealthKit unit for measuring temperature in degrees Fahrenheit.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/degreeFahrenheit()
func (hc _HKUnitClass) DegreeFahrenheitUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("degreeFahrenheitUnit"))
	return rv
}

// Returns a HealthKit unit for measuring the optical power of a lens using diopter units.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/diopter()
func (hc _HKUnitClass) DiopterUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("diopterUnit"))
	return rv
}

// Converts a HealthKit unit object into a corresponding energy formatter enumeration value.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/energyFormatterUnit(from:)
func (hc _HKUnitClass) EnergyFormatterUnitFromUnit(unit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("energyFormatterUnitFromUnit:"), unit)
	return rv
}

// Returns a HealthKit unit for measuring volume in imperial fluid ounces.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/fluidOunceImperial()
func (hc _HKUnitClass) FluidOunceImperialUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("fluidOunceImperialUnit"))
	return rv
}

// Returns a HealthKit unit for measuring volume in US fluid ounces.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/fluidOunceUS()
func (hc _HKUnitClass) FluidOunceUSUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("fluidOunceUSUnit"))
	return rv
}

// Returns a HealthKit unit for measuring length in feet.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/foot()
func (hc _HKUnitClass) FootUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("footUnit"))
	return rv
}

// Returns a HealthKit unit for measuring mass in grams.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/gram()
func (hc _HKUnitClass) GramUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("gramUnit"))
	return rv
}

// Returns a HealthKit unit for measuring mass, using gram units with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/gramUnit(with:)
func (hc _HKUnitClass) GramUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("gramUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring frequency in hertz.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/hertz()
func (hc _HKUnitClass) HertzUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("hertzUnit"))
	return rv
}

// Returns a HealthKit unit for measuring frequency in hertz with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/hertzUnit(with:)
func (hc _HKUnitClass) HertzUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("hertzUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring time in hours.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/hour()
func (hc _HKUnitClass) HourUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("hourUnit"))
	return rv
}

// Returns a HealthKit unit for measuring length in inches.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/inch()
func (hc _HKUnitClass) InchUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("inchUnit"))
	return rv
}

// Returns a HealthKit unit for measuring pressure in inches of mercury.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/inchesOfMercury()
func (hc _HKUnitClass) InchesOfMercuryUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("inchesOfMercuryUnit"))
	return rv
}

// Converts an energy formatter enumeration value into a corresponding HealthKit unit object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-1j1pq
func (hc _HKUnitClass) UnitFromEnergyFormatterUnit(energyFormatterUnit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("unitFromEnergyFormatterUnit:"), energyFormatterUnit)
	return rv
}

// Converts a length formatter enumeration value into a corresponding HealthKit object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-55e1u
func (hc _HKUnitClass) UnitFromLengthFormatterUnit(lengthFormatterUnit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("unitFromLengthFormatterUnit:"), lengthFormatterUnit)
	return rv
}

// Converts a mass formatter enumeration value into a corresponding HealthKit unit object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-7h2li
func (hc _HKUnitClass) UnitFromMassFormatterUnit(massFormatterUnit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("unitFromMassFormatterUnit:"), massFormatterUnit)
	return rv
}

// Returns the unit instance described by the provided string.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-9qont
func (hc _HKUnitClass) UnitFromString(string_ string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("unitFromString:"), objc.String(string_))
	return rv
}

// Returns a HealthKit unit that measures the amount of a biologically active substance in international units (IU).
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/internationalUnit()
func (hc _HKUnitClass) InternationalUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("internationalUnit"))
	return rv
}

// Returns a HealthKit unit for measuring energy in joules.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/joule()
func (hc _HKUnitClass) JouleUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("jouleUnit"))
	return rv
}

// Returns a HealthKit unit for measuring energy, using joule units with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/jouleUnit(with:)
func (hc _HKUnitClass) JouleUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("jouleUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring temperature in kelvins.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/kelvin()
func (hc _HKUnitClass) KelvinUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("kelvinUnit"))
	return rv
}

// Returns a HealthKit unit for measuring energy in kilocalories.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/kilocalorie()
func (hc _HKUnitClass) KilocalorieUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("kilocalorieUnit"))
	return rv
}

// Returns a HealthKit unit for measuring energy in large calories (Cal).
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/largeCalorie()
func (hc _HKUnitClass) LargeCalorieUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("largeCalorieUnit"))
	return rv
}

// Converts a HealthKit unit object into a corresponding length formatter enumeration value.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/lengthFormatterUnit(from:)
func (hc _HKUnitClass) LengthFormatterUnitFromUnit(unit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("lengthFormatterUnitFromUnit:"), unit)
	return rv
}

// Returns a HealthKit unit for measuring volume in liters.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/liter()
func (hc _HKUnitClass) LiterUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("literUnit"))
	return rv
}

// Returns a HealthKit unit for measuring volume, using liter units with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/literUnit(with:)
func (hc _HKUnitClass) LiterUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("literUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring illuminance in lux.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/lux()
func (hc _HKUnitClass) LuxUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("luxUnit"))
	return rv
}

// Returns a HealthKit unit for measuring illuminance, using lux units with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/luxUnit(with:)
func (hc _HKUnitClass) LuxUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("luxUnitWithMetricPrefix:"), prefix)
	return rv
}

// Converts a HealthKit unit object into a corresponding mass formatter enumeration value.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/massFormatterUnit(from:)
func (hc _HKUnitClass) MassFormatterUnitFromUnit(unit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("massFormatterUnitFromUnit:"), unit)
	return rv
}

// Returns a HealthKit unit for measuring length in meters.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/meter()
func (hc _HKUnitClass) MeterUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("meterUnit"))
	return rv
}

// Returns a HealthKit unit for measuring length, using meter units with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/meterUnit(with:)
func (hc _HKUnitClass) MeterUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("meterUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring length in miles.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/mile()
func (hc _HKUnitClass) MileUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("mileUnit"))
	return rv
}

// Returns a HealthKit unit for measuring pressure in millimeters of mercury.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/millimeterOfMercury()
func (hc _HKUnitClass) MillimeterOfMercuryUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("millimeterOfMercuryUnit"))
	return rv
}

// Returns a HealthKit unit for measuring time in minutes.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/minute()
func (hc _HKUnitClass) MinuteUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("minuteUnit"))
	return rv
}

// Returns a HealthKit unit for measuring mass in moles, with the given prefix and molar mass.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/moleUnit(with:molarMass:)
func (hc _HKUnitClass) MoleUnitWithMetricPrefixMolarMass(prefix unsafe.Pointer, gramsPerMole unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("moleUnitWithMetricPrefix:molarMass:"), prefix, gramsPerMole)
	return rv
}

// Returns a HealthKit unit for measuring mass in moles for a given molar mass.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/moleUnit(withMolarMass:)
func (hc _HKUnitClass) MoleUnitWithMolarMass(gramsPerMole unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("moleUnitWithMolarMass:"), gramsPerMole)
	return rv
}

// Returns a HealthKit unit for measuring mass in ounces.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/ounce()
func (hc _HKUnitClass) OunceUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("ounceUnit"))
	return rv
}

// Returns a HealthKit unit for measuring pressure in pascals.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/pascal()
func (hc _HKUnitClass) PascalUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("pascalUnit"))
	return rv
}

// Returns a HealthKit unit for measuring pressure, using pascal units with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/pascalUnit(with:)
func (hc _HKUnitClass) PascalUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("pascalUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring percentages.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/percent()
func (hc _HKUnitClass) PercentUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("percentUnit"))
	return rv
}

// Returns a HealthKit unit for measuring volume in imperial pints.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/pintImperial()
func (hc _HKUnitClass) PintImperialUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("pintImperialUnit"))
	return rv
}

// Returns a HealthKit unit for measuring volume in US pints.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/pintUS()
func (hc _HKUnitClass) PintUSUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("pintUSUnit"))
	return rv
}

// Returns a HealthKit unit for measuring mass in pounds.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/pound()
func (hc _HKUnitClass) PoundUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("poundUnit"))
	return rv
}

// Returns a HealthKit unit for measuring the prismatic deviation of a lens using prism diopter units.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/prismDiopter()
func (hc _HKUnitClass) PrismDiopterUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("prismDiopterUnit"))
	return rv
}

// Returns a HealthKit unit for measuring angles using radians.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/radianAngle()
func (hc _HKUnitClass) RadianAngleUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("radianAngleUnit"))
	return rv
}

// Returns a HealthKit unit for measuring angles, using radian units with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/radianAngleUnit(with:)
func (hc _HKUnitClass) RadianAngleUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("radianAngleUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring time in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/second()
func (hc _HKUnitClass) SecondUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("secondUnit"))
	return rv
}

// Returns a HealthKit unit for measuring time, using second units with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/secondUnit(with:)
func (hc _HKUnitClass) SecondUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("secondUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring electrical conductance in siemens.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/siemen()
func (hc _HKUnitClass) SiemenUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("siemenUnit"))
	return rv
}

// Returns a HealthKit unit for measuring electrical conductance, using siemen units with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/siemenUnit(with:)
func (hc _HKUnitClass) SiemenUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("siemenUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring energy in small calories (cal).
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/smallCalorie()
func (hc _HKUnitClass) SmallCalorieUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("smallCalorieUnit"))
	return rv
}

// Returns a HealthKit unit for measuring mass in stones.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/stone()
func (hc _HKUnitClass) StoneUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("stoneUnit"))
	return rv
}

// Returns a HealthKit unit for measuring the difference in electrical potential using volts.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/volt()
func (hc _HKUnitClass) VoltUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("voltUnit"))
	return rv
}

// Returns a HealthKit unit for measuring the electrical potential difference in volts with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/voltUnit(with:)
func (hc _HKUnitClass) VoltUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("voltUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring power in watts.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/watt()
func (hc _HKUnitClass) WattUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("wattUnit"))
	return rv
}

// Returns a HealthKit unit for measuring power, using watt units with the provided prefix.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/wattUnit(with:)
func (hc _HKUnitClass) WattUnitWithMetricPrefix(prefix unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("wattUnitWithMetricPrefix:"), prefix)
	return rv
}

// Returns a HealthKit unit for measuring length in yards.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/yard()
func (hc _HKUnitClass) YardUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(hc.class), objc.Sel("yardUnit"))
	return rv
}

// Returns a Boolean value indicating whether the unit is null.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/isNull()
func (h_ HKUnit) IsNull() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isNull"))
	return rv
}

// Returns a complex unit representing the unit’s reciprocal.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/reciprocal()
func (h_ HKUnit) ReciprocalUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("reciprocalUnit"))
	return rv
}

// Creates a complex unit by dividing the receiving unit by another unit.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/unitDivided(by:)
func (h_ HKUnit) UnitDividedByUnit(unit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("unitDividedByUnit:"), unit)
	return rv
}

// Creates a complex unit by multiplying the receiving unit with another unit.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/unitMultiplied(by:)
func (h_ HKUnit) UnitMultipliedByUnit(unit unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("unitMultipliedByUnit:"), unit)
	return rv
}

// Creates a complex unit by raising the unit to the given power.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/unitRaised(toPower:)
func (h_ HKUnit) UnitRaisedToPower(power int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("unitRaisedToPower:"), power)
	return rv
}

// A string representation of the unit object.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/unitString
func (h_ HKUnit) UnitString() string {
	rv := objc.Send[string](h_.ID, objc.Sel("unitString"))
	return rv
}


