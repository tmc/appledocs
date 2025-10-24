// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

package findersync

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/appkit"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFIFinderSync is the FIFinderSync protocol interface.
//
// The group of methods to implement for modifying the Finder user interface to express file synchronization status and control.
//
// Availability:
//   - macOS 10.10+
//
// See: doc://com.apple.Finder-Sync/documentation/FinderSync/FIFinderSyncProtocol
type PFIFinderSync interface {
	// Optional methods
	BeginObservingDirectoryAtURL(url objc.IObject /* cross-framework: NSURL */)
	HasBeginObservingDirectoryAtURL() bool
	EndObservingDirectoryAtURL(url objc.IObject /* cross-framework: NSURL */)
	HasEndObservingDirectoryAtURL() bool
	MakeListenerEndpointForServiceNameItemURLAndReturnError(serviceName FileProviderServiceName /* not a class type */, itemURL objc.IObject /* cross-framework: NSURL */, error_ unsafe.Pointer) foundation.XPCListenerEndpoint
	HasMakeListenerEndpointForServiceNameItemURLAndReturnError() bool
	MenuForMenuKind(menu FIMenuKind) appkit.Menu
	HasMenuForMenuKind() bool
	RequestBadgeIdentifierForURL(url objc.IObject /* cross-framework: NSURL */)
	HasRequestBadgeIdentifierForURL() bool
	SupportedServiceNamesForItemWithURL(itemURL objc.IObject /* cross-framework: NSURL */) []string
	HasSupportedServiceNamesForItemWithURL() bool
	ValuesForAttributesForItemWithURLCompletion(attributes []string, itemURL objc.IObject /* cross-framework: NSURL */, completion unsafe.Pointer)
	HasValuesForAttributesForItemWithURLCompletion() bool
}
