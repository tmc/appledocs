// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation_test

import (
	"github.com/tmc/appledocs/generated/foundation"
)

// Suppress unused import errors
var _ = foundation.NewURLSessionStreamTask

// ExampleNewURLSessionStreamTask demonstrates how to create a URLSessionStreamTask instance.
func ExampleNewURLSessionStreamTask() {
	_ = foundation.NewURLSessionStreamTask()
	// Output:
}
// ExampleURLSessionStreamTask_CaptureStreams demonstrates using CaptureStreams on a URLSessionStreamTask instance.
// Completes any already enqueued reads and writes, and then invokes the   delegate message.
func ExampleURLSessionStreamTask_CaptureStreams() {
	obj := foundation.NewURLSessionStreamTask()
	obj.CaptureStreams()
	// Output:
	}

// ExampleURLSessionStreamTask_CloseRead demonstrates using CloseRead on a URLSessionStreamTask instance.
// Completes any enqueued reads and writes, and then closes the read side of the underlying socket.
func ExampleURLSessionStreamTask_CloseRead() {
	obj := foundation.NewURLSessionStreamTask()
	obj.CloseRead()
	// Output:
	}

// ExampleURLSessionStreamTask_CloseWrite demonstrates using CloseWrite on a URLSessionStreamTask instance.
// Completes any enqueued reads and writes, and then closes the write side of the underlying socket.
func ExampleURLSessionStreamTask_CloseWrite() {
	obj := foundation.NewURLSessionStreamTask()
	obj.CloseWrite()
	// Output:
	}

// ExampleURLSessionStreamTask_StartSecureConnection demonstrates using StartSecureConnection on a URLSessionStreamTask instance.
// Completes any enqueued reads and writes, and establishes a secure connection.
func ExampleURLSessionStreamTask_StartSecureConnection() {
	obj := foundation.NewURLSessionStreamTask()
	obj.StartSecureConnection()
	// Output:
	}

