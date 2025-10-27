// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	LocalizedTitlesForItem(item objectivec.IObject) []string
	SearchForItemsWithSearchStringResultLimitMatchedItemHandler(searchString foundation.foundation.INSString, resultLimit int, handleMatchedItems unsafe.Pointer)
	// Optional methods
	PerformActionForItem(item objectivec.IObject)
	HasPerformActionForItem() bool
	ShowAllHelpTopicsForSearchString(searchString foundation.foundation.INSString)
	HasShowAllHelpTopicsForSearchString() bool
}
