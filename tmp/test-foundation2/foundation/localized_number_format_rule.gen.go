// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var localizedNumberFormatRuleClass _LocalizedNumberFormatRuleClass

func init() {
	localizedNumberFormatRuleClass = _LocalizedNumberFormatRuleClass{objc.GetClass("NSLocalizedNumberFormatRule")}
}

type _LocalizedNumberFormatRuleClass struct {
	class objc.Class
}

type LocalizedNumberFormatRule struct {
	objc.ID
}

func LocalizedNumberFormatRuleFrom(ptr unsafe.Pointer) LocalizedNumberFormatRule {
	return LocalizedNumberFormatRule{
		ID: objc.ID(ptr),
	}
}




