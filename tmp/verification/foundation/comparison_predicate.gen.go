// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var comparisonPredicateClass _ComparisonPredicateClass

func init() {
	comparisonPredicateClass = _ComparisonPredicateClass{objc.GetClass("NSComparisonPredicate")}
}

type _ComparisonPredicateClass struct {
	class objc.Class
}

type ComparisonPredicate struct {
	objc.ID
}

func ComparisonPredicateFrom(ptr unsafe.Pointer) ComparisonPredicate {
	return ComparisonPredicate{
		ID: objc.ID(ptr),
	}
}




