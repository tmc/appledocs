// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [InflectionRule] class.
var inflectionRuleClass = _InflectionRuleClass{objc.GetClass("NSInflectionRule")}

type _InflectionRuleClass struct {
	class objc.Class
}

// A rule that affects how an attributed string performs automatic grammatical agreement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRule

type InflectionRule struct {
	objectivec.Object
}

// InflectionRuleFrom constructs a [InflectionRule] from an unsafe.Pointer.
//
// A rule that affects how an attributed string performs automatic grammatical agreement.
func InflectionRuleFrom(ptr unsafe.Pointer) InflectionRule {
	return InflectionRule{objectivec.Object{objc.ID(ptr)}}
}



