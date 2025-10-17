// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [InflectionRule] class.
var InflectionRuleClass objc.Class

func init() {
	InflectionRuleClass = objc.GetClass("NSInflectionRule")
}

type InflectionRule struct {
	objc.ID
}

func InflectionRuleFrom(ptr unsafe.Pointer) InflectionRule {
	return InflectionRule{
		ID: objc.ID(ptr),
	}
}




