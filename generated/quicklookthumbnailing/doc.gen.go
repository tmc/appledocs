// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

// Package quicklookthumbnailing provides Go bindings for the QuickLookThumbnailing framework.
//
// Generate thumbnails for common file types and add a Thumbnail Extension to your app to enable others to create thumbnails of your custom files.
//
// These bindings are generated from Apple's official documentation and
// provide purego-based access to QuickLookThumbnailing without requiring cgo.
//
// See: https://developer.apple.com/documentation/QuickLookThumbnailing
package quicklookthumbnailing

import (
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/QuickLookThumbnailing.framework/QuickLookThumbnailing"
func init() {
	_, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
}

