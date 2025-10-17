// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [bytes] class.
var bytesClass = _bytesClass{objc.GetClass("bytes")}

type _bytesClass struct {
	class objc.Class
}

// An interface definition for the [bytes] class.
type Ibytes interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSimpleCString/bytes

type bytes struct {
	objectivec.Object
}

// bytesFrom constructs a [bytes] from an unsafe.Pointer.
func bytesFrom(ptr unsafe.Pointer) bytes {
	return bytes{objectivec.Object{objc.ID(ptr)}}
}



