// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WritingToolsCoordinatorAnimationParameters] class.
var writingToolsCoordinatorAnimationParametersClass = _WritingToolsCoordinatorAnimationParametersClass{objc.GetClass("NSWritingToolsCoordinatorAnimationParameters")}

type _WritingToolsCoordinatorAnimationParametersClass struct {
	class objc.Class
}

// An object you use to configure additional tasks or animations to run alongside the Writing Tools animations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/AnimationParameters

type WritingToolsCoordinatorAnimationParameters struct {
	objectivec.Object
}

// WritingToolsCoordinatorAnimationParametersFrom constructs a [WritingToolsCoordinatorAnimationParameters] from an unsafe.Pointer.
//
// An object you use to configure additional tasks or animations to run alongside the Writing Tools animations.
func WritingToolsCoordinatorAnimationParametersFrom(ptr unsafe.Pointer) WritingToolsCoordinatorAnimationParameters {
	return WritingToolsCoordinatorAnimationParameters{objectivec.Object{objc.ID(ptr)}}
}



