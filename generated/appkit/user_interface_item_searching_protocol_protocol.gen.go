// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PUserInterfaceItemSearching is the NSUserInterfaceItemSearching protocol interface.
//
// A set of methods an app can implement to provide Spotlight for Help for its own custom help data.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSUserInterfaceItemSearching
type PUserInterfaceItemSearching interface {
	// Required methods
	LocalizedTitlesForItem(item objc.IObject) []string/* debug [protocol_interface/required_method]: LocalizedTitlesForItem */
	SearchForItemsWithSearchStringResultLimitMatchedItemHandler(searchString objc.IObject /* cross-framework: NSString */, resultLimit int, handleMatchedItems unsafe.Pointer)/* debug [protocol_interface/required_method]: SearchForItemsWithSearchStringResultLimitMatchedItemHandler */
	// Optional methods
	PerformActionForItem(item objc.IObject)
	HasPerformActionForItem() bool
	ShowAllHelpTopicsForSearchString(searchString objc.IObject /* cross-framework: NSString */)
	HasShowAllHelpTopicsForSearchString() bool
}
