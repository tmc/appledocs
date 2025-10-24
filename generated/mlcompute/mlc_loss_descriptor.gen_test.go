// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCLossDescriptor

// ExampleNewCLossDescriptorWithTypeReductionType demonstrates how to create a CLossDescriptor instance using NewCLossDescriptorWithTypeReductionType.
// Creates a loss descriptor with the loss function and reduction type you specify.
func ExampleNewCLossDescriptorWithTypeReductionType() {
	_ = mlcompute.NewCLossDescriptorWithTypeReductionType(
		mlcompute.CLossType{}, // lossType CLossType
		mlcompute.CReductionType{}, // reductionType CReductionType
	)
	// Output:
}
// ExampleNewCLossDescriptorWithTypeReductionTypeWeight demonstrates how to create a CLossDescriptor instance using NewCLossDescriptorWithTypeReductionTypeWeight.
// Creates a loss descriptor with the loss function, reduction type, and weight you specify.
func ExampleNewCLossDescriptorWithTypeReductionTypeWeight() {
	_ = mlcompute.NewCLossDescriptorWithTypeReductionTypeWeight(
		mlcompute.CLossType{}, // lossType CLossType
		mlcompute.CReductionType{}, // reductionType CReductionType
		0.0, // weight float32
	)
	// Output:
}
// ExampleNewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCount demonstrates how to create a CLossDescriptor instance using NewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCount.
// Creates a loss descriptor with the loss function, reduction type, weight, label smoothing, and number of classes you specify.
func ExampleNewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCount() {
	_ = mlcompute.NewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCount(
		mlcompute.CLossType{}, // lossType CLossType
		mlcompute.CReductionType{}, // reductionType CReductionType
		0.0, // weight float32
		0.0, // labelSmoothing float32
		10, // classCount uint
	)
	// Output:
}
// ExampleNewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCountEpsilonDelta demonstrates how to create a CLossDescriptor instance using NewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCountEpsilonDelta.
// Creates a loss descriptor with the loss function, reduction type, weight, label smoothing, and number of classes, epsilon, and delta that you specify.
func ExampleNewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCountEpsilonDelta() {
	_ = mlcompute.NewCLossDescriptorWithTypeReductionTypeWeightLabelSmoothingClassCountEpsilonDelta(
		mlcompute.CLossType{}, // lossType CLossType
		mlcompute.CReductionType{}, // reductionType CReductionType
		0.0, // weight float32
		0.0, // labelSmoothing float32
		10, // classCount uint
		0.0, // epsilon float32
		0.0, // delta float32
	)
	// Output:
}
