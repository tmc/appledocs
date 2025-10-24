// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewCNNConvolutionGradient

// ExampleCNNConvolutionGradient_ReloadWeightsAndBiases demonstrates using ReloadWeightsAndBiases on a CNNConvolutionGradient instance.
func ExampleCNNConvolutionGradient_ReloadWeightsAndBiases() {
	obj := metalperformanceshaders.NewCNNConvolutionGradient()
	obj.ReloadWeightsAndBiases()
	// Output:
	}

// ExampleCNNConvolutionGradient_ReloadWeightsAndBiasesFromDataSource demonstrates using ReloadWeightsAndBiasesFromDataSource on a CNNConvolutionGradient instance.
func ExampleCNNConvolutionGradient_ReloadWeightsAndBiasesFromDataSource() {
	obj := metalperformanceshaders.NewCNNConvolutionGradient()
	obj.ReloadWeightsAndBiasesFromDataSource()
	// Output:
	}

