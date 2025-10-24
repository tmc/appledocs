// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewSVGFDenoiser

// ExampleSVGFDenoiser_ClearTemporalHistory demonstrates using ClearTemporalHistory on a SVGFDenoiser instance.
func ExampleSVGFDenoiser_ClearTemporalHistory() {
	obj := metalperformanceshaders.NewSVGFDenoiser()
	obj.ClearTemporalHistory()
	// Output:
	}

// ExampleSVGFDenoiser_ReleaseTemporaryTextures demonstrates using ReleaseTemporaryTextures on a SVGFDenoiser instance.
func ExampleSVGFDenoiser_ReleaseTemporaryTextures() {
	obj := metalperformanceshaders.NewSVGFDenoiser()
	obj.ReleaseTemporaryTextures()
	// Output:
	}

// ExampleSVGFDenoiser_Encode demonstrates using Encode on a SVGFDenoiser instance.
func ExampleSVGFDenoiser_Encode() {
	obj := metalperformanceshaders.NewSVGFDenoiser()
	obj.Encode()
	// Output:
	}

