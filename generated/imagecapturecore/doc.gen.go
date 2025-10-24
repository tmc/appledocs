
// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

// Package imagecapturecore provides Go bindings for the ImageCaptureCore framework.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ImageCaptureCore without requiring cgo.
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

