// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RuleEditor] class.
var RuleEditorClass objc.Class

func init() {
	RuleEditorClass = objc.GetClass("NSRuleEditor")
}

type RuleEditor struct {
	objc.ID
}

func RuleEditorFrom(ptr unsafe.Pointer) RuleEditor {
	return RuleEditor{
		ID: objc.ID(ptr),
	}
}




