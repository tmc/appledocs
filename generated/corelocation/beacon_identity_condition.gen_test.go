// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation_test

import (
	"github.com/tmc/appledocs/generated/corelocation"
)


// ExampleNewBeaconIdentityConditionWithUUID demonstrates how to create a BeaconIdentityCondition instance using NewBeaconIdentityConditionWithUUID.
// Creates a new beacon identity condition with the identifier you specify.
func ExampleNewBeaconIdentityConditionWithUUID() {
	_ = corelocation.NewBeaconIdentityConditionWithUUID(
		nil, // uuid unsafe.Pointer
	)
	// Output:
}

// ExampleNewBeaconIdentityConditionWithUUIDMajor demonstrates how to create a BeaconIdentityCondition instance using NewBeaconIdentityConditionWithUUIDMajor.
// Creates a new beacon identity condition with the identifier and major value you specify.
func ExampleNewBeaconIdentityConditionWithUUIDMajor() {
	_ = corelocation.NewBeaconIdentityConditionWithUUIDMajor(
		nil, // uuid unsafe.Pointer
		nil, // major unsafe.Pointer
	)
	// Output:
}

// ExampleNewBeaconIdentityConditionWithUUIDMajorMinor demonstrates how to create a BeaconIdentityCondition instance using NewBeaconIdentityConditionWithUUIDMajorMinor.
// Creates a new beacon identity condition with the identifier, and major and minor values you specify.
func ExampleNewBeaconIdentityConditionWithUUIDMajorMinor() {
	_ = corelocation.NewBeaconIdentityConditionWithUUIDMajorMinor(
		nil, // uuid unsafe.Pointer
		nil, // major unsafe.Pointer
		nil, // minor unsafe.Pointer
	)
	// Output:
}


