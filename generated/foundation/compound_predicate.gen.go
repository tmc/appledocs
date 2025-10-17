// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [CompoundPredicate] class.
var CompoundPredicateClass objc.Class

func init() {
	CompoundPredicateClass = objc.GetClass("NSCompoundPredicate")
}

type CompoundPredicate struct {
	objc.ID
}

func CompoundPredicateFrom(ptr unsafe.Pointer) CompoundPredicate {
	return CompoundPredicate{
		ID: objc.ID(ptr),
	}
}



