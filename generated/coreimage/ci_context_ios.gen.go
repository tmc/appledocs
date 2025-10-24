//go:build darwin && ios

// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Context


// Returns the maximum size allowed for any image rendered into the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/inputImageMaximumSize()
func (c_ Context) InputImageMaximumSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](c_.ID, objc.Sel("inputImageMaximumSize"))
	return rv
}

// Returns the maximum size allowed for any image created by the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIContext/outputImageMaximumSize()
func (c_ Context) OutputImageMaximumSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](c_.ID, objc.Sel("outputImageMaximumSize"))
	return rv
}

// iOS-only properties




