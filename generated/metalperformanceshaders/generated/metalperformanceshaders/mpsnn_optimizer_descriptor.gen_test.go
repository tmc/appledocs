// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewOptimizerDescriptor

// ExampleNewOptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale demonstrates how to create a OptimizerDescriptor instance using NewOptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale.
func ExampleNewOptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale() {
	_ = metalperformanceshaders.NewOptimizerDescriptorWithLearningRateGradientRescaleApplyGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(
		0.0, // learningRate float32
		1.0, // gradientRescale float32
		false, // applyGradientClipping bool
		0.0, // gradientClipMax float32
		0.0, // gradientClipMin float32
		metalperformanceshaders.RegularizationType{}, // regularizationType RegularizationType
		1.0, // regularizationScale float32
	)
	// Output:
}
// ExampleNewOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale demonstrates how to create a OptimizerDescriptor instance using NewOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale.
func ExampleNewOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale() {
	_ = metalperformanceshaders.NewOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(
		0.0, // learningRate float32
		1.0, // gradientRescale float32
		metalperformanceshaders.RegularizationType{}, // regularizationType RegularizationType
		1.0, // regularizationScale float32
	)
	// Output:
}
