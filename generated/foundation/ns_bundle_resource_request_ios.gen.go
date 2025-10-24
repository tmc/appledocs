//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for BundleResourceRequest


// Requests access to the resources marked with the managed tags. If any of the resources are not on the device, they are requested from the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundleResourceRequest/beginAccessingResources(completionHandler:)
func (b_ BundleResourceRequest) BeginAccessingResourcesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("beginAccessingResourcesWithCompletionHandler:"), completionHandler)
}

// Checks whether the resources marked with the tags managed by the request are already on the device. If all of the resources are on the device, you can begin accessing those resources.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSBundleResourceRequest/conditionallyBeginAccessingResources(completionHandler:)
func (b_ BundleResourceRequest) ConditionallyBeginAccessingResourcesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("conditionallyBeginAccessingResourcesWithCompletionHandler:"), completionHandler)
}

// iOS-only properties





