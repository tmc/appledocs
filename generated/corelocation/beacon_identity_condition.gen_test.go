// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation_test

import (
	"github.com/tmc/appledocs/generated/corelocation"
)

// Suppress unused import errors
var _ = corelocation.NewBeaconIdentityCondition

// ExampleNewBeaconIdentityConditionWithUUID demonstrates how to create a BeaconIdentityCondition instance using NewBeaconIdentityConditionWithUUID.
// Creates a new beacon identity condition with the identifier you specify.
func ExampleNewBeaconIdentityConditionWithUUID() {
	_ = corelocation.NewBeaconIdentityConditionWithUUID(
		corelocation.UUID{}, // uuid UUID
	)
	// Output:
}

// ExampleNewBeaconIdentityConditionWithUUIDMajor demonstrates how to create a BeaconIdentityCondition instance using NewBeaconIdentityConditionWithUUIDMajor.
// Creates a new beacon identity condition with the identifier and major value you specify.
func ExampleNewBeaconIdentityConditionWithUUIDMajor() {
	_ = corelocation.NewBeaconIdentityConditionWithUUIDMajor(
		corelocation.UUID{},             // uuid UUID
		corelocation.BeaconMajorValue{}, // major BeaconMajorValue
	)
	// Output:
}

// ExampleNewBeaconIdentityConditionWithUUIDMajorMinor demonstrates how to create a BeaconIdentityCondition instance using NewBeaconIdentityConditionWithUUIDMajorMinor.
// Creates a new beacon identity condition with the identifier, and major and minor values you specify.
func ExampleNewBeaconIdentityConditionWithUUIDMajorMinor() {
	_ = corelocation.NewBeaconIdentityConditionWithUUIDMajorMinor(
		corelocation.UUID{},             // uuid UUID
		corelocation.BeaconMajorValue{}, // major BeaconMajorValue
		corelocation.BeaconMinorValue{}, // minor BeaconMinorValue
	)
	// Output:
}
