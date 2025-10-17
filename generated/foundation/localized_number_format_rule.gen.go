// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LocalizedNumberFormatRule] class.
var localizedNumberFormatRuleClass = _LocalizedNumberFormatRuleClass{objc.GetClass("NSLocalizedNumberFormatRule")}

type _LocalizedNumberFormatRuleClass struct {
	class objc.Class
}

// An interface definition for the [LocalizedNumberFormatRule] class.
type ILocalizedNumberFormatRule interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocalizedNumberFormatRule

type LocalizedNumberFormatRule struct {
	objectivec.Object
}

// LocalizedNumberFormatRuleFrom constructs a [LocalizedNumberFormatRule] from an unsafe.Pointer.
func LocalizedNumberFormatRuleFrom(ptr unsafe.Pointer) LocalizedNumberFormatRule {
	return LocalizedNumberFormatRule{objectivec.Object{objc.ID(ptr)}}
}



