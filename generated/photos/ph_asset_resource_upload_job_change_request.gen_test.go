// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos_test

import (
	"github.com/tmc/appledocs/generated/photos"
)

// Suppress unused import errors
var _ = photos.NewPHAssetResourceUploadJobChangeRequest

// ExampleNewPHAssetResourceUploadJobChangeRequestForUploadJob demonstrates how to create a PHAssetResourceUploadJobChangeRequest instance using NewPHAssetResourceUploadJobChangeRequestForUploadJob.
// Creates a request for modifying the specified upload job.
func ExampleNewPHAssetResourceUploadJobChangeRequestForUploadJob() {
	_ = photos.NewPHAssetResourceUploadJobChangeRequestForUploadJob(
		photos.PHAssetResourceUploadJob{}, // job PHAssetResourceUploadJob
	)
	// Output:
}
