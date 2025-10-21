// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit_test

import (
	"github.com/tmc/appledocs/generated/fskit"
)

// Suppress unused import errors
var _ = fskit.NewFSVolume

// ExampleNewFSVolumeWithVolumeIDVolumeName demonstrates how to create a FSVolume instance using NewFSVolumeWithVolumeIDVolumeName.
// Creates a volume with the given identifier and name.
func ExampleNewFSVolumeWithVolumeIDVolumeName() {
	_ = fskit.NewFSVolumeWithVolumeIDVolumeName(
		fskit.FSVolumeIdentifier{}, // volumeID FSVolumeIdentifier
		fskit.FSFileName{}, // volumeName FSFileName
	)
	// Output:
}
