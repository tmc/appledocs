// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [InflectionRuleExplicit] class.
var InflectionRuleExplicitClass objc.Class

func init() {
	InflectionRuleExplicitClass = objc.GetClass("NSInflectionRuleExplicit")
}

type InflectionRuleExplicit struct {
	objc.ID
}

func InflectionRuleExplicitFrom(ptr unsafe.Pointer) InflectionRuleExplicit {
	return InflectionRuleExplicit{
		ID: objc.ID(ptr),
	}
}




