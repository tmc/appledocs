// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [InflectionRuleExplicit] class.
var inflectionRuleExplicitClass = _InflectionRuleExplicitClass{objc.GetClass("NSInflectionRuleExplicit")}

type _InflectionRuleExplicitClass struct {
	class objc.Class
}

// An inflection rule that uses a morphology instance to determine how to inflect attribued strings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSInflectionRuleExplicit

type InflectionRuleExplicit struct {
	InflectionRule
}

// InflectionRuleExplicitFrom constructs a [InflectionRuleExplicit] from an unsafe.Pointer.
//
// An inflection rule that uses a morphology instance to determine how to inflect attribued strings.
func InflectionRuleExplicitFrom(ptr unsafe.Pointer) InflectionRuleExplicit {
	return InflectionRuleExplicit{
		InflectionRule: InflectionRuleFrom(ptr),
	}
}



