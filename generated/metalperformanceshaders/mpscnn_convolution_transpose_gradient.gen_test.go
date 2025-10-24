// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewCNNConvolutionTransposeGradient

// ExampleCNNConvolutionTransposeGradient_ReloadWeightsAndBiasesFromDataSource demonstrates using ReloadWeightsAndBiasesFromDataSource on a CNNConvolutionTransposeGradient instance.
func ExampleCNNConvolutionTransposeGradient_ReloadWeightsAndBiasesFromDataSource() {
	obj := metalperformanceshaders.NewCNNConvolutionTransposeGradient()
	obj.ReloadWeightsAndBiasesFromDataSource()
	// Output:
	}

// ExampleCNNConvolutionTransposeGradient_ReloadWeightsAndBiases demonstrates using ReloadWeightsAndBiases on a CNNConvolutionTransposeGradient instance.
func ExampleCNNConvolutionTransposeGradient_ReloadWeightsAndBiases() {
	obj := metalperformanceshaders.NewCNNConvolutionTransposeGradient()
	obj.ReloadWeightsAndBiases()
	// Output:
	}

