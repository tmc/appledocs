// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var inflectionRuleExplicitClass _InflectionRuleExplicitClass

func init() {
	inflectionRuleExplicitClass = _InflectionRuleExplicitClass{objc.GetClass("NSInflectionRuleExplicit")}
}

type _InflectionRuleExplicitClass struct {
	class objc.Class
}

type InflectionRuleExplicit struct {
	objc.ID
}

func InflectionRuleExplicitFrom(ptr unsafe.Pointer) InflectionRuleExplicit {
	return InflectionRuleExplicit{
		ID: objc.ID(ptr),
	}
}




