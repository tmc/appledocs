// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

// Package imagecapturecore provides Go bindings for the ImageCaptureCore framework.
//
// Browse for media devices and control them programmatically from your app. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ImageCaptureCore without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageCaptureCore
package imagecapturecore

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ImageCaptureCore.framework/ImageCaptureCore"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

