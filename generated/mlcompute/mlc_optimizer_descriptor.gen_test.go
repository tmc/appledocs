// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute_test

import (
	"github.com/tmc/appledocs/generated/mlcompute"
)

// Suppress unused import errors
var _ = mlcompute.NewCOptimizerDescriptor

// ExampleNewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale demonstrates how to create a COptimizerDescriptor instance using NewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale.
// Creates a descriptor with the learning rate, gradient rescale, clipping option and values, and regularization type and scale that you specify.
func ExampleNewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale() {
	_ = mlcompute.NewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClipMaxGradientClipMinRegularizationTypeRegularizationScale(
		0.0, // learningRate float32
		1.0, // gradientRescale float32
		false, // appliesGradientClipping bool
		0.0, // gradientClipMax float32
		0.0, // gradientClipMin float32
		mlcompute.CRegularizationType{}, // regularizationType CRegularizationType
		1.0, // regularizationScale float32
	)
	// Output:
}
// ExampleNewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClippingTypeGradientClipMaxGradientClipMinMaximumClippingNormCustomGlobalNormRegularizationTypeRegularizationScale demonstrates how to create a COptimizerDescriptor instance using NewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClippingTypeGradientClipMaxGradientClipMinMaximumClippingNormCustomGlobalNormRegularizationTypeRegularizationScale.
// Creates a descriptor with the learning rate, gradient rescale, clipping option and values, and regularization type and scale that you specify.
func ExampleNewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClippingTypeGradientClipMaxGradientClipMinMaximumClippingNormCustomGlobalNormRegularizationTypeRegularizationScale() {
	_ = mlcompute.NewCOptimizerDescriptorWithLearningRateGradientRescaleAppliesGradientClippingGradientClippingTypeGradientClipMaxGradientClipMinMaximumClippingNormCustomGlobalNormRegularizationTypeRegularizationScale(
		0.0, // learningRate float32
		1.0, // gradientRescale float32
		false, // appliesGradientClipping bool
		mlcompute.CGradientClippingType{}, // gradientClippingType CGradientClippingType
		0.0, // gradientClipMax float32
		0.0, // gradientClipMin float32
		0.0, // maximumClippingNorm float32
		0.0, // customGlobalNorm float32
		mlcompute.CRegularizationType{}, // regularizationType CRegularizationType
		1.0, // regularizationScale float32
	)
	// Output:
}
// ExampleNewCOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale demonstrates how to create a COptimizerDescriptor instance using NewCOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale.
// Creates an optimizer descriptor with the learning rate, gradient rescale, regularization type, and regulation scale that you specify.
func ExampleNewCOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale() {
	_ = mlcompute.NewCOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(
		0.0, // learningRate float32
		1.0, // gradientRescale float32
		mlcompute.CRegularizationType{}, // regularizationType CRegularizationType
		1.0, // regularizationScale float32
	)
	// Output:
}
