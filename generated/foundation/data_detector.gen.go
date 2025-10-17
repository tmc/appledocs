// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [DataDetector] class.
var dataDetectorClass = _DataDetectorClass{objc.GetClass("NSDataDetector")}

type _DataDetectorClass struct {
	class objc.Class
}

// A specialized regular expression object that matches natural language text for predefined data patterns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSDataDetector

type DataDetector struct {
	RegularExpression
}

// DataDetectorFrom constructs a [DataDetector] from an unsafe.Pointer.
//
// A specialized regular expression object that matches natural language text for predefined data patterns.
func DataDetectorFrom(ptr unsafe.Pointer) DataDetector {
	return DataDetector{
		RegularExpression: RegularExpressionFrom(ptr),
	}
}



