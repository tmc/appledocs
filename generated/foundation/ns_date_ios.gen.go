//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Date


// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDate/srAbsoluteTime
func (d_ Date) SrAbsoluteTime() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("srAbsoluteTime"))
	return rv
}




