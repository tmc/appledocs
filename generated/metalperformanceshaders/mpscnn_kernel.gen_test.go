// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders_test

import (
	"github.com/tmc/appledocs/generated/metalperformanceshaders"
)

// Suppress unused import errors
var _ = metalperformanceshaders.NewCNNKernel

// ExampleCNNKernel_Encode demonstrates using Encode on a CNNKernel instance.
// Encodes a kernel into a command buffer.  The ensuing operation proceeds out-of-place.
func ExampleCNNKernel_Encode() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.Encode()
	// Output:
	}

// ExampleCNNKernel_EncodeBatch demonstrates using EncodeBatch on a CNNKernel instance.
func ExampleCNNKernel_EncodeBatch() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.EncodeBatch()
	// Output:
	}

// ExampleCNNKernel_DestinationImageDescriptor demonstrates using DestinationImageDescriptor on a CNNKernel instance.
func ExampleCNNKernel_DestinationImageDescriptor() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.DestinationImageDescriptor()
	// Output:
	}

// ExampleCNNKernel_IsResultStateReusedAcrossBatch demonstrates using IsResultStateReusedAcrossBatch on a CNNKernel instance.
func ExampleCNNKernel_IsResultStateReusedAcrossBatch() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.IsResultStateReusedAcrossBatch()
	// Output:
	}

// ExampleCNNKernel_AppendBatchBarrier demonstrates using AppendBatchBarrier on a CNNKernel instance.
func ExampleCNNKernel_AppendBatchBarrier() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.AppendBatchBarrier()
	// Output:
	}

// ExampleCNNKernel_ResultStateBatch demonstrates using ResultStateBatch on a CNNKernel instance.
func ExampleCNNKernel_ResultStateBatch() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.ResultStateBatch()
	// Output:
	}

// ExampleCNNKernel_ResultState demonstrates using ResultState on a CNNKernel instance.
func ExampleCNNKernel_ResultState() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.ResultState()
	// Output:
	}

// ExampleCNNKernel_TemporaryResultState demonstrates using TemporaryResultState on a CNNKernel instance.
func ExampleCNNKernel_TemporaryResultState() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.TemporaryResultState()
	// Output:
	}

// ExampleCNNKernel_TemporaryResultStateBatch demonstrates using TemporaryResultStateBatch on a CNNKernel instance.
func ExampleCNNKernel_TemporaryResultStateBatch() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.TemporaryResultStateBatch()
	// Output:
	}

// ExampleCNNKernel_BatchEncodingStorageSize demonstrates using BatchEncodingStorageSize on a CNNKernel instance.
func ExampleCNNKernel_BatchEncodingStorageSize() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.BatchEncodingStorageSize()
	// Output:
	}

// ExampleCNNKernel_EncodingStorageSize demonstrates using EncodingStorageSize on a CNNKernel instance.
func ExampleCNNKernel_EncodingStorageSize() {
	obj := metalperformanceshaders.NewCNNKernel()
	obj.EncodingStorageSize()
	// Output:
	}

