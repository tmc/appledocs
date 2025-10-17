// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [LocalizedNumberFormatRule] class.
var LocalizedNumberFormatRuleClass objc.Class

func init() {
	LocalizedNumberFormatRuleClass = objc.GetClass("NSLocalizedNumberFormatRule")
}

type LocalizedNumberFormatRule struct {
	objc.ID
}

func LocalizedNumberFormatRuleFrom(ptr unsafe.Pointer) LocalizedNumberFormatRule {
	return LocalizedNumberFormatRule{
		ID: objc.ID(ptr),
	}
}



