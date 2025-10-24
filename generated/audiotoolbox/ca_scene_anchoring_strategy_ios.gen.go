//go:build darwin && ios

// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for SceneAnchoringStrategy


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AudioToolbox/CASceneAnchoringStrategy/sceneIdentifier
func (s_ SceneAnchoringStrategy) SceneIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("sceneIdentifier"))
	return rv
}




