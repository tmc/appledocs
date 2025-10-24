//go:build darwin && ios

// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for OSSystemExtensionsWorkspace


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SystemExtensions/OSSystemExtensionsWorkspace/systemExtensions(forApplicationWithBundleID:)
func (o_ OSSystemExtensionsWorkspace) SystemExtensionsForApplicationWithBundleIDError(bundleID objc.IObject /* cross-framework: NSString */, out_error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("systemExtensionsForApplicationWithBundleID:error:"), bundleID, out_error)
	return rv
}

// iOS-only properties





