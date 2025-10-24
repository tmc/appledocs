// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class HKUnit */


/* debug [class_header]: Header for HKUnit */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for HKUnit */
// An interface definition for the [HKUnit] class.
type IHKUnit interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for HKUnit */
	// properties:
	UnitString() objc.IObject /* cross-framework: NSString */
	HKUnitMolarMassBloodGlucose() float64
	SetHKUnitMolarMassBloodGlucose(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for HKUnit */
	// methods:
	IsNull() bool
	ReciprocalUnit() IHKUnit
	UnitDividedByUnit(unit IHKUnit) IHKUnit
	UnitMultipliedByUnit(unit IHKUnit) IHKUnit
	UnitRaisedToPower(power int) IHKUnit
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for HKUnit */
// Alloc allocates a new instance without initialization.
func (hc _HKUnitClass) Alloc() HKUnit {
	rv := objc.Send[HKUnit](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for HKUnit */
// A class for managing the units of measure within HealthKit.
//
// The unit class supports most standard SI units (meters, seconds, and grams), SI units with prefixes (centimeters, milliseconds and kilograms) and equivalent non-SI units (feet, minutes, and pounds). HealthKit also supports creating complex units by mathematically combining existing units. You use units when working with HealthKit quantities. Quantities store both the value (as a data type) and its corresponding unit. You can then request the value from the quantity in any compatible units. For more information on working with quantities, see .


// A class for managing the units of measure within HealthKit.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for HKUnit */

// Converts an energy formatter enumeration value into a corresponding HealthKit unit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-1j1pq
func NewHKUnitFromEnergyFormatterUnit(energyFormatterUnit EnergyFormatterUnit /* not a class type */) HKUnit {
	rv := objc.Send[HKUnit](objc.ID(getHKUnitClass().class), objc.Sel("unitFromEnergyFormatterUnit:"), energyFormatterUnit)
	return rv
}/* debug [class_init_methods/constructor]: NewHKUnitFromEnergyFormatterUnit */


// Converts a length formatter enumeration value into a corresponding HealthKit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-55e1u
func NewHKUnitFromLengthFormatterUnit(lengthFormatterUnit LengthFormatterUnit /* not a class type */) HKUnit {
	rv := objc.Send[HKUnit](objc.ID(getHKUnitClass().class), objc.Sel("unitFromLengthFormatterUnit:"), lengthFormatterUnit)
	return rv
}/* debug [class_init_methods/constructor]: NewHKUnitFromLengthFormatterUnit */


// Converts a mass formatter enumeration value into a corresponding HealthKit unit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-7h2li
func NewHKUnitFromMassFormatterUnit(massFormatterUnit MassFormatterUnit /* not a class type */) HKUnit {
	rv := objc.Send[HKUnit](objc.ID(getHKUnitClass().class), objc.Sel("unitFromMassFormatterUnit:"), massFormatterUnit)
	return rv
}/* debug [class_init_methods/constructor]: NewHKUnitFromMassFormatterUnit */


// Returns the unit instance described by the provided string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-9qont
func NewHKUnitFromString(string_ objc.IObject /* cross-framework: NSString */) HKUnit {
	rv := objc.Send[HKUnit](objc.ID(getHKUnitClass().class), objc.Sel("unitFromString:"), string_)
	return rv
}/* debug [class_init_methods/constructor]: NewHKUnitFromString */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for HKUnit */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/appleEffortScore()
func (hc _HKUnitClass) AppleEffortScoreUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("appleEffortScoreUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AppleEffortScoreUnit) */


// Returns a HealthKit unit for measuring pressure in atmospheres.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/atmosphere()
func (hc _HKUnitClass) AtmosphereUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("atmosphereUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AtmosphereUnit) */


// Returns a HealthKit unit for measuring energy in calories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/calorie()
func (hc _HKUnitClass) CalorieUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("calorieUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CalorieUnit) */


// Returns a HealthKit unit for measuring pressure in centimeters of water.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/centimeterOfWater()
func (hc _HKUnitClass) CentimeterOfWaterUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("centimeterOfWaterUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CentimeterOfWaterUnit) */


// Returns a HealthKit unit for measuring counts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/count()
func (hc _HKUnitClass) CountUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("countUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CountUnit) */


// Returns a HealthKit unit for measuring volume in imperial cups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/cupImperial()
func (hc _HKUnitClass) CupImperialUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("cupImperialUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CupImperialUnit) */


// Returns a HealthKit unit for measuring volume in US cups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/cupUS()
func (hc _HKUnitClass) CupUSUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("cupUSUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CupUSUnit) */


// Returns a HealthKit unit for measuring time in days.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/day()
func (hc _HKUnitClass) DayUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("dayUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DayUnit) */


// Returns a HealthKit unit for measuring the difference between the local pressure and the ambient atmospheric pressure caused by sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/decibelAWeightedSoundPressureLevel()
func (hc _HKUnitClass) DecibelAWeightedSoundPressureLevelUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("decibelAWeightedSoundPressureLevelUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecibelAWeightedSoundPressureLevelUnit) */


// Returns a HealthKit unit for measuring the intensity of a sound.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/decibelHearingLevel()
func (hc _HKUnitClass) DecibelHearingLevelUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("decibelHearingLevelUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DecibelHearingLevelUnit) */


// Returns a HealthKit unit for measuring angles using degrees.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/degreeAngle()
func (hc _HKUnitClass) DegreeAngleUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("degreeAngleUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DegreeAngleUnit) */


// Returns a HealthKit unit for measuring temperature in degrees Celsius.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/degreeCelsius()
func (hc _HKUnitClass) DegreeCelsiusUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("degreeCelsiusUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DegreeCelsiusUnit) */


// Returns a HealthKit unit for measuring temperature in degrees Fahrenheit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/degreeFahrenheit()
func (hc _HKUnitClass) DegreeFahrenheitUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("degreeFahrenheitUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DegreeFahrenheitUnit) */


// Returns a HealthKit unit for measuring the optical power of a lens using diopter units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/diopter()
func (hc _HKUnitClass) DiopterUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("diopterUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DiopterUnit) */


// Converts a HealthKit unit object into a corresponding energy formatter enumeration value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/energyFormatterUnit(from:)
func (hc _HKUnitClass) EnergyFormatterUnitFromUnit(unit IHKUnit) EnergyFormatterUnit /* not a class type */ {
	rv := objc.Send[EnergyFormatterUnit](objc.ID(hc.class), objc.Sel("energyFormatterUnitFromUnit:"), unit)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=EnergyFormatterUnitFromUnit) */


// Returns a HealthKit unit for measuring volume in imperial fluid ounces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/fluidOunceImperial()
func (hc _HKUnitClass) FluidOunceImperialUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("fluidOunceImperialUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FluidOunceImperialUnit) */


// Returns a HealthKit unit for measuring volume in US fluid ounces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/fluidOunceUS()
func (hc _HKUnitClass) FluidOunceUSUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("fluidOunceUSUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FluidOunceUSUnit) */


// Returns a HealthKit unit for measuring length in feet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/foot()
func (hc _HKUnitClass) FootUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("footUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FootUnit) */


// Returns a HealthKit unit for measuring mass in grams.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/gram()
func (hc _HKUnitClass) GramUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("gramUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GramUnit) */


// Returns a HealthKit unit for measuring mass, using gram units with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/gramUnit(with:)
func (hc _HKUnitClass) GramUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("gramUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GramUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring frequency in hertz.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/hertz()
func (hc _HKUnitClass) HertzUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("hertzUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HertzUnit) */


// Returns a HealthKit unit for measuring frequency in hertz with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/hertzUnit(with:)
func (hc _HKUnitClass) HertzUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("hertzUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HertzUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring time in hours.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/hour()
func (hc _HKUnitClass) HourUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("hourUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HourUnit) */


// Returns a HealthKit unit for measuring length in inches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/inch()
func (hc _HKUnitClass) InchUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("inchUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InchUnit) */


// Returns a HealthKit unit for measuring pressure in inches of mercury.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/inchesOfMercury()
func (hc _HKUnitClass) InchesOfMercuryUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("inchesOfMercuryUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InchesOfMercuryUnit) */


// Converts an energy formatter enumeration value into a corresponding HealthKit unit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-1j1pq
func (hc _HKUnitClass) UnitFromEnergyFormatterUnit(energyFormatterUnit EnergyFormatterUnit /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("unitFromEnergyFormatterUnit:"), energyFormatterUnit)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnitFromEnergyFormatterUnit) */


// Converts a length formatter enumeration value into a corresponding HealthKit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-55e1u
func (hc _HKUnitClass) UnitFromLengthFormatterUnit(lengthFormatterUnit LengthFormatterUnit /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("unitFromLengthFormatterUnit:"), lengthFormatterUnit)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnitFromLengthFormatterUnit) */


// Converts a mass formatter enumeration value into a corresponding HealthKit unit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-7h2li
func (hc _HKUnitClass) UnitFromMassFormatterUnit(massFormatterUnit MassFormatterUnit /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("unitFromMassFormatterUnit:"), massFormatterUnit)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnitFromMassFormatterUnit) */


// Returns the unit instance described by the provided string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/init(from:)-9qont
func (hc _HKUnitClass) UnitFromString(string_ objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("unitFromString:"), string_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=UnitFromString) */


// Returns a HealthKit unit that measures the amount of a biologically active substance in international units (IU).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/internationalUnit()
func (hc _HKUnitClass) InternationalUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("internationalUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=InternationalUnit) */


// Returns a HealthKit unit for measuring energy in joules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/joule()
func (hc _HKUnitClass) JouleUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("jouleUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=JouleUnit) */


// Returns a HealthKit unit for measuring energy, using joule units with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/jouleUnit(with:)
func (hc _HKUnitClass) JouleUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("jouleUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=JouleUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring temperature in kelvins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/kelvin()
func (hc _HKUnitClass) KelvinUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("kelvinUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=KelvinUnit) */


// Returns a HealthKit unit for measuring energy in kilocalories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/kilocalorie()
func (hc _HKUnitClass) KilocalorieUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("kilocalorieUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=KilocalorieUnit) */


// Returns a HealthKit unit for measuring energy in large calories (Cal).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/largeCalorie()
func (hc _HKUnitClass) LargeCalorieUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("largeCalorieUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LargeCalorieUnit) */


// Converts a HealthKit unit object into a corresponding length formatter enumeration value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/lengthFormatterUnit(from:)
func (hc _HKUnitClass) LengthFormatterUnitFromUnit(unit IHKUnit) LengthFormatterUnit /* not a class type */ {
	rv := objc.Send[LengthFormatterUnit](objc.ID(hc.class), objc.Sel("lengthFormatterUnitFromUnit:"), unit)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LengthFormatterUnitFromUnit) */


// Returns a HealthKit unit for measuring volume in liters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/liter()
func (hc _HKUnitClass) LiterUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("literUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LiterUnit) */


// Returns a HealthKit unit for measuring volume, using liter units with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/literUnit(with:)
func (hc _HKUnitClass) LiterUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("literUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LiterUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring illuminance in lux.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/lux()
func (hc _HKUnitClass) LuxUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("luxUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LuxUnit) */


// Returns a HealthKit unit for measuring illuminance, using lux units with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/luxUnit(with:)
func (hc _HKUnitClass) LuxUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("luxUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LuxUnitWithMetricPrefix) */


// Converts a HealthKit unit object into a corresponding mass formatter enumeration value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/massFormatterUnit(from:)
func (hc _HKUnitClass) MassFormatterUnitFromUnit(unit IHKUnit) MassFormatterUnit /* not a class type */ {
	rv := objc.Send[MassFormatterUnit](objc.ID(hc.class), objc.Sel("massFormatterUnitFromUnit:"), unit)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MassFormatterUnitFromUnit) */


// Returns a HealthKit unit for measuring length in meters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/meter()
func (hc _HKUnitClass) MeterUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("meterUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MeterUnit) */


// Returns a HealthKit unit for measuring length, using meter units with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/meterUnit(with:)
func (hc _HKUnitClass) MeterUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("meterUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MeterUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring length in miles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/mile()
func (hc _HKUnitClass) MileUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("mileUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MileUnit) */


// Returns a HealthKit unit for measuring pressure in millimeters of mercury.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/millimeterOfMercury()
func (hc _HKUnitClass) MillimeterOfMercuryUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("millimeterOfMercuryUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MillimeterOfMercuryUnit) */


// Returns a HealthKit unit for measuring time in minutes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/minute()
func (hc _HKUnitClass) MinuteUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("minuteUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MinuteUnit) */


// Returns a HealthKit unit for measuring mass in moles, with the given prefix and molar mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/moleUnit(with:molarMass:)
func (hc _HKUnitClass) MoleUnitWithMetricPrefixMolarMass(prefix HKMetricPrefix, gramsPerMole float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("moleUnitWithMetricPrefix:molarMass:"), prefix, gramsPerMole)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MoleUnitWithMetricPrefixMolarMass) */


// Returns a HealthKit unit for measuring mass in moles for a given molar mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/moleUnit(withMolarMass:)
func (hc _HKUnitClass) MoleUnitWithMolarMass(gramsPerMole float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("moleUnitWithMolarMass:"), gramsPerMole)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MoleUnitWithMolarMass) */


// Returns a HealthKit unit for measuring mass in ounces.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/ounce()
func (hc _HKUnitClass) OunceUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("ounceUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=OunceUnit) */


// Returns a HealthKit unit for measuring pressure in pascals.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/pascal()
func (hc _HKUnitClass) PascalUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("pascalUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PascalUnit) */


// Returns a HealthKit unit for measuring pressure, using pascal units with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/pascalUnit(with:)
func (hc _HKUnitClass) PascalUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("pascalUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PascalUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring percentages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/percent()
func (hc _HKUnitClass) PercentUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("percentUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PercentUnit) */


// Returns a HealthKit unit for measuring volume in imperial pints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/pintImperial()
func (hc _HKUnitClass) PintImperialUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("pintImperialUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PintImperialUnit) */


// Returns a HealthKit unit for measuring volume in US pints.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/pintUS()
func (hc _HKUnitClass) PintUSUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("pintUSUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PintUSUnit) */


// Returns a HealthKit unit for measuring mass in pounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/pound()
func (hc _HKUnitClass) PoundUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("poundUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PoundUnit) */


// Returns a HealthKit unit for measuring the prismatic deviation of a lens using prism diopter units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/prismDiopter()
func (hc _HKUnitClass) PrismDiopterUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("prismDiopterUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PrismDiopterUnit) */


// Returns a HealthKit unit for measuring angles using radians.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/radianAngle()
func (hc _HKUnitClass) RadianAngleUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("radianAngleUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RadianAngleUnit) */


// Returns a HealthKit unit for measuring angles, using radian units with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/radianAngleUnit(with:)
func (hc _HKUnitClass) RadianAngleUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("radianAngleUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RadianAngleUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring time in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/second()
func (hc _HKUnitClass) SecondUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("secondUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SecondUnit) */


// Returns a HealthKit unit for measuring time, using second units with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/secondUnit(with:)
func (hc _HKUnitClass) SecondUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("secondUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SecondUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring electrical conductance in siemens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/siemen()
func (hc _HKUnitClass) SiemenUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("siemenUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SiemenUnit) */


// Returns a HealthKit unit for measuring electrical conductance, using siemen units with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/siemenUnit(with:)
func (hc _HKUnitClass) SiemenUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("siemenUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SiemenUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring energy in small calories (cal).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/smallCalorie()
func (hc _HKUnitClass) SmallCalorieUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("smallCalorieUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SmallCalorieUnit) */


// Returns a HealthKit unit for measuring mass in stones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/stone()
func (hc _HKUnitClass) StoneUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("stoneUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=StoneUnit) */


// Returns a HealthKit unit for measuring the difference in electrical potential using volts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/volt()
func (hc _HKUnitClass) VoltUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("voltUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VoltUnit) */


// Returns a HealthKit unit for measuring the electrical potential difference in volts with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/voltUnit(with:)
func (hc _HKUnitClass) VoltUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("voltUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=VoltUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring power in watts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/watt()
func (hc _HKUnitClass) WattUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("wattUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WattUnit) */


// Returns a HealthKit unit for measuring power, using watt units with the provided prefix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/wattUnit(with:)
func (hc _HKUnitClass) WattUnitWithMetricPrefix(prefix HKMetricPrefix) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("wattUnitWithMetricPrefix:"), prefix)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=WattUnitWithMetricPrefix) */


// Returns a HealthKit unit for measuring length in yards.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/yard()
func (hc _HKUnitClass) YardUnit() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(hc.class), objc.Sel("yardUnit"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=YardUnit) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for HKUnit */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for HKUnit */

// Returns a Boolean value indicating whether the unit is null.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/isNull()
func (h_ HKUnit) IsNull() bool {
	rv := objc.Send[bool](h_.ID, objc.Sel("isNull"))
	return rv
}/* debug [instance_methods/method]: IsNull */


// Returns a complex unit representing the unit’s reciprocal.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/reciprocal()
func (h_ HKUnit) ReciprocalUnit() HKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("reciprocalUnit"))
	return rv
}/* debug [instance_methods/method]: ReciprocalUnit */


// Creates a complex unit by dividing the receiving unit by another unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/unitDivided(by:)
func (h_ HKUnit) UnitDividedByUnit(unit IHKUnit) HKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("unitDividedByUnit:"), unit)
	return rv
}/* debug [instance_methods/method]: UnitDividedByUnit */


// Creates a complex unit by multiplying the receiving unit with another unit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/unitMultiplied(by:)
func (h_ HKUnit) UnitMultipliedByUnit(unit IHKUnit) HKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("unitMultipliedByUnit:"), unit)
	return rv
}/* debug [instance_methods/method]: UnitMultipliedByUnit */


// Creates a complex unit by raising the unit to the given power.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/unitRaised(toPower:)
func (h_ HKUnit) UnitRaisedToPower(power int) HKUnit {
	rv := objc.Send[HKUnit](h_.ID, objc.Sel("unitRaisedToPower:"), power)
	return rv
}/* debug [instance_methods/method]: UnitRaisedToPower */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for HKUnit */

// A string representation of the unit object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKUnit/unitString
func (h_ HKUnit) UnitString() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](h_.ID, objc.Sel("unitString"))
	return rv
}/* debug [instance_properties/getter]: unitString */


// The molecular mass of blood glucose, typically used to create mole units for blood glucose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkunitmolarmassbloodglucose
func (h_ HKUnit) HKUnitMolarMassBloodGlucose() float64 {
	rv := objc.Send[float64](h_.ID, objc.Sel("HKUnitMolarMassBloodGlucose"))
	return rv
}/* debug [instance_properties/getter]: HKUnitMolarMassBloodGlucose */


// The molecular mass of blood glucose, typically used to create mole units for blood glucose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/healthkit/hkunitmolarmassbloodglucose
func (h_ HKUnit) SetHKUnitMolarMassBloodGlucose(value float64) {
	objc.Send[objc.ID](h_.ID, objc.Sel("setHKUnitMolarMassBloodGlucose:"), value)
}/* debug [instance_properties/setter]: HKUnitMolarMassBloodGlucose */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class HKUnit */


