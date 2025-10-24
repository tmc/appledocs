// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewCommandBuffer

// ExampleCommandBuffer_CommitAndContinue demonstrates using CommitAndContinue on a CommandBuffer instance.
//
// Note: This example is not executed because CommitAndContinue crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleCommandBuffer_CommitAndContinue() {
	obj := metalperformanceshaders.NewCommandBuffer()
	obj.CommitAndContinue()
	}

// ExampleCommandBuffer_PrefetchHeap demonstrates using PrefetchHeap on a CommandBuffer instance.
func ExampleCommandBuffer_PrefetchHeap() {
	obj := metalperformanceshaders.NewCommandBuffer()
	obj.PrefetchHeap()
	// Output:
	}

