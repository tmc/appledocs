// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ComparisonPredicate] class.
var ComparisonPredicateClass objc.Class

func init() {
	ComparisonPredicateClass = objc.GetClass("NSComparisonPredicate")
}

type ComparisonPredicate struct {
	objc.ID
}

func ComparisonPredicateFrom(ptr unsafe.Pointer) ComparisonPredicate {
	return ComparisonPredicate{
		ID: objc.ID(ptr),
	}
}



