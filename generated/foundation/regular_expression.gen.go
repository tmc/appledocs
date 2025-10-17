// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RegularExpression] class.
var RegularExpressionClass objc.Class

func init() {
	RegularExpressionClass = objc.GetClass("NSRegularExpression")
}

type RegularExpression struct {
	objc.ID
}

func RegularExpressionFrom(ptr unsafe.Pointer) RegularExpression {
	return RegularExpression{
		ID: objc.ID(ptr),
	}
}



