// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

// Package imageio provides Go bindings for the ImageIO framework.
//
// Read and write most image file formats, and access an image’s metadata. [Full Topic]
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to ImageIO without requiring cgo.
//
// [Full Topic]: https://developer.apple.com/documentation/ImageIO
package imageio

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ImageIO.framework/ImageIO"


func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}


