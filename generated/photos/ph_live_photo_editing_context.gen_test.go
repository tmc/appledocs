// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos_test

import (
	"github.com/tmc/appledocs/generated/photos"
)

// Suppress unused import errors
var _ = photos.NewPHLivePhotoEditingContext

// ExampleNewPHLivePhotoEditingContextWithLivePhotoEditingInput demonstrates how to create a PHLivePhotoEditingContext instance using NewPHLivePhotoEditingContextWithLivePhotoEditingInput.
// Creates a Live Photo editing context for the specified editing input.
func ExampleNewPHLivePhotoEditingContextWithLivePhotoEditingInput() {
	_ = photos.NewPHLivePhotoEditingContextWithLivePhotoEditingInput(
		photos.PHContentEditingInput{}, // livePhotoInput PHContentEditingInput
	)
	// Output:
}
