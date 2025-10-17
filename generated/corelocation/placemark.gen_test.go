// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation_test

import (
	"github.com/tmc/appledocs/generated/corelocation"
)


// ExampleNewPlacemark demonstrates how to create a Placemark instance.
func ExampleNewPlacemark() {
	_ = corelocation.NewPlacemark()
	// Output:
}

// ExampleNewPlacemarkWithLocationNamePostalAddress demonstrates how to create a Placemark instance using NewPlacemarkWithLocationNamePostalAddress.
func ExampleNewPlacemarkWithLocationNamePostalAddress() {
	_ = corelocation.NewPlacemarkWithLocationNamePostalAddress(
		nil, // location unsafe.Pointer
		"name", // name string
		nil, // postalAddress unsafe.Pointer
	)
	// Output:
}

// ExampleNewPlacemarkWithPlacemark demonstrates how to create a Placemark instance using NewPlacemarkWithPlacemark.
// Initializes and returns a placemark object from another placemark object.
func ExampleNewPlacemarkWithPlacemark() {
	_ = corelocation.NewPlacemarkWithPlacemark(
		nil, // placemark unsafe.Pointer
	)
	// Output:
}


