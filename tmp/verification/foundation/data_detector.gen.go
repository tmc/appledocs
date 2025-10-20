// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var dataDetectorClass _DataDetectorClass

func init() {
	dataDetectorClass = _DataDetectorClass{objc.GetClass("NSDataDetector")}
}

type _DataDetectorClass struct {
	class objc.Class
}

type DataDetector struct {
	objc.ID
}

func DataDetectorFrom(ptr unsafe.Pointer) DataDetector {
	return DataDetector{
		ID: objc.ID(ptr),
	}
}




