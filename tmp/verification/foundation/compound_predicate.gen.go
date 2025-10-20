// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var compoundPredicateClass _CompoundPredicateClass

func init() {
	compoundPredicateClass = _CompoundPredicateClass{objc.GetClass("NSCompoundPredicate")}
}

type _CompoundPredicateClass struct {
	class objc.Class
}

type CompoundPredicate struct {
	objc.ID
}

func CompoundPredicateFrom(ptr unsafe.Pointer) CompoundPredicate {
	return CompoundPredicate{
		ID: objc.ID(ptr),
	}
}




