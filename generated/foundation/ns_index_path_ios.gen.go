//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for IndexPath


// iOS-only properties

// An index number identifying a row in a section of a table view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath/row
func (i_ IndexPath) Row() int {
	rv := objc.Send[int](i_.ID, objc.Sel("row"))
	return rv
}




