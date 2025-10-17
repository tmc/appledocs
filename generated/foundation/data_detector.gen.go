// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [DataDetector] class.
var DataDetectorClass objc.Class

func init() {
	DataDetectorClass = objc.GetClass("NSDataDetector")
}

type DataDetector struct {
	objc.ID
}

func DataDetectorFrom(ptr unsafe.Pointer) DataDetector {
	return DataDetector{
		ID: objc.ID(ptr),
	}
}



