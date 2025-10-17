// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PredicateEditor] class.
var PredicateEditorClass objc.Class

func init() {
	PredicateEditorClass = objc.GetClass("NSPredicateEditor")
}

type PredicateEditor struct {
	objc.ID
}

func PredicateEditorFrom(ptr unsafe.Pointer) PredicateEditor {
	return PredicateEditor{
		ID: objc.ID(ptr),
	}
}



