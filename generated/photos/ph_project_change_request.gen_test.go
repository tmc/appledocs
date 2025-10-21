// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos_test

import (
	"github.com/tmc/appledocs/generated/photos"
)

// Suppress unused import errors
var _ = photos.NewPHProjectChangeRequest

// ExampleNewPHProjectChangeRequestWithProject demonstrates how to create a PHProjectChangeRequest instance using NewPHProjectChangeRequestWithProject.
// Creates a change request around the specified project.
func ExampleNewPHProjectChangeRequestWithProject() {
	_ = photos.NewPHProjectChangeRequestWithProject(
		photos.PHProject{}, // project PHProject
	)
	// Output:
}
