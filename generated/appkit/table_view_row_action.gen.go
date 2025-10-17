// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TableViewRowAction] class.
var tableViewRowActionClass = _TableViewRowActionClass{objc.GetClass("NSTableViewRowAction")}

type _TableViewRowActionClass struct {
	class objc.Class
}

// A single action to present when the user swipes horizontally on a table row. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction

type TableViewRowAction struct {
	objectivec.Object
}

// TableViewRowActionFrom constructs a [TableViewRowAction] from an unsafe.Pointer.
//
// A single action to present when the user swipes horizontally on a table row.
func TableViewRowActionFrom(ptr unsafe.Pointer) TableViewRowAction {
	return TableViewRowAction{objectivec.Object{objc.ID(ptr)}}
}



