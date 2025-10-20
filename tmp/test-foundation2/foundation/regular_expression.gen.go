// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var regularExpressionClass _RegularExpressionClass

func init() {
	regularExpressionClass = _RegularExpressionClass{objc.GetClass("NSRegularExpression")}
}

type _RegularExpressionClass struct {
	class objc.Class
}

type RegularExpression struct {
	objc.ID
}

func RegularExpressionFrom(ptr unsafe.Pointer) RegularExpression {
	return RegularExpression{
		ID: objc.ID(ptr),
	}
}




