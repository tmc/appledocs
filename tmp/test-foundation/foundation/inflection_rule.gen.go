// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var InflectionRuleClass _InflectionRuleClass

func init() {
	InflectionRuleClass = _InflectionRuleClass{objc.GetClass("NSInflectionRule")}
}

type _InflectionRuleClass struct {
	class objc.Class
}

type InflectionRule struct {
	objc.ID
}

func InflectionRuleFrom(ptr unsafe.Pointer) InflectionRule {
	return InflectionRule{
		ID: objc.ID(ptr),
	}
}




