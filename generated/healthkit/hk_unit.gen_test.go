// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit_test

import (
	"github.com/tmc/appledocs/generated/healthkit"
)

// Suppress unused import errors
var _ = healthkit.NewHKUnit

// ExampleNewHKUnitFromEnergyFormatterUnit demonstrates how to create a HKUnit instance using NewHKUnitFromEnergyFormatterUnit.
// Converts an energy formatter enumeration value into a corresponding HealthKit unit object.
func ExampleNewHKUnitFromEnergyFormatterUnit() {
	_ = healthkit.NewHKUnitFromEnergyFormatterUnit(
		healthkit.EnergyFormatterUnit /* not a class type */{}, // energyFormatterUnit EnergyFormatterUnit /* not a class type */
	)
	// Output:
}
// ExampleNewHKUnitFromLengthFormatterUnit demonstrates how to create a HKUnit instance using NewHKUnitFromLengthFormatterUnit.
// Converts a length formatter enumeration value into a corresponding HealthKit object.
func ExampleNewHKUnitFromLengthFormatterUnit() {
	_ = healthkit.NewHKUnitFromLengthFormatterUnit(
		healthkit.LengthFormatterUnit /* not a class type */{}, // lengthFormatterUnit LengthFormatterUnit /* not a class type */
	)
	// Output:
}
// ExampleNewHKUnitFromMassFormatterUnit demonstrates how to create a HKUnit instance using NewHKUnitFromMassFormatterUnit.
// Converts a mass formatter enumeration value into a corresponding HealthKit unit object.
func ExampleNewHKUnitFromMassFormatterUnit() {
	_ = healthkit.NewHKUnitFromMassFormatterUnit(
		healthkit.MassFormatterUnit /* not a class type */{}, // massFormatterUnit MassFormatterUnit /* not a class type */
	)
	// Output:
}
// ExampleHKUnit_IsNull demonstrates using IsNull on a HKUnit instance.
// Returns a Boolean value indicating whether the unit is null.
func ExampleHKUnit_IsNull() {
	obj := healthkit.NewHKUnit()
	_ = obj.IsNull()
	// Output:
	}

// ExampleHKUnit_ReciprocalUnit demonstrates using ReciprocalUnit on a HKUnit instance.
// Returns a complex unit representing the unit’s reciprocal.
func ExampleHKUnit_ReciprocalUnit() {
	obj := healthkit.NewHKUnit()
	_ = obj.ReciprocalUnit()
	// Output:
	}

