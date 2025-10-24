// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCActivationDescriptor

// ExampleNewCActivationDescriptorWithType demonstrates how to create a CActivationDescriptor instance using NewCActivationDescriptorWithType.
// Creates an activation descriptor with the activation type you specify.
func ExampleNewCActivationDescriptorWithType() {
	_ = mlcompute.NewCActivationDescriptorWithType(
		mlcompute.CActivationType{}, // activationType CActivationType
	)
	// Output:
}
// ExampleNewCActivationDescriptorWithTypeA demonstrates how to create a CActivationDescriptor instance using NewCActivationDescriptorWithTypeA.
// Creates an activation descriptor with the activation type and parameter a that you specify.
func ExampleNewCActivationDescriptorWithTypeA() {
	_ = mlcompute.NewCActivationDescriptorWithTypeA(
		mlcompute.CActivationType{}, // activationType CActivationType
		0.0, // a float32
	)
	// Output:
}
// ExampleNewCActivationDescriptorWithTypeAB demonstrates how to create a CActivationDescriptor instance using NewCActivationDescriptorWithTypeAB.
// Creates an activation descriptor with the activation type and parameters a and b that you specify.
func ExampleNewCActivationDescriptorWithTypeAB() {
	_ = mlcompute.NewCActivationDescriptorWithTypeAB(
		mlcompute.CActivationType{}, // activationType CActivationType
		0.0, // a float32
		0.0, // b float32
	)
	// Output:
}
// ExampleNewCActivationDescriptorWithTypeABC demonstrates how to create a CActivationDescriptor instance using NewCActivationDescriptorWithTypeABC.
// Creates an activation descriptor with the activation type and parameters a, b, and c that you specify.
func ExampleNewCActivationDescriptorWithTypeABC() {
	_ = mlcompute.NewCActivationDescriptorWithTypeABC(
		mlcompute.CActivationType{}, // activationType CActivationType
		0.0, // a float32
		0.0, // b float32
		0.0, // c float32
	)
	// Output:
}
