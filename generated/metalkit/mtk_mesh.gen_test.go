// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit_test

import (
	"github.com/tmc/appledocs/generated/metalkit"
)


// ExampleNewMTKMeshWithMeshDeviceError demonstrates how to create a MTKMesh instance using NewMTKMeshWithMeshDeviceError.
// Initializes a MetalKit mesh and its submeshes from a Model I/O mesh.
func ExampleNewMTKMeshWithMeshDeviceError() {
	_ = metalkit.NewMTKMeshWithMeshDeviceError(
		nil, // mesh unsafe.Pointer
		nil, // device unsafe.Pointer
		nil, // error unsafe.Pointer
	)
	// Output:
}


