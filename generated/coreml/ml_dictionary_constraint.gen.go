// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [DictionaryConstraint] class.
var (
	DictionaryConstraintClass     _DictionaryConstraintClass
	DictionaryConstraintClassOnce sync.Once
)

func getDictionaryConstraintClass() _DictionaryConstraintClass {
	DictionaryConstraintClassOnce.Do(func() {
		DictionaryConstraintClass = _DictionaryConstraintClass{objc.GetClass("MLDictionaryConstraint")}
	})
	return DictionaryConstraintClass
}

type _DictionaryConstraintClass struct {
	class objc.Class
}

// An interface definition for the [DictionaryConstraint] class.
type IDictionaryConstraint interface {
	objectivec.IObject
}

// The constraint on the keys for a dictionary feature.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreML/MLDictionaryConstraint
type DictionaryConstraint struct {
	objectivec.Object
}

// DictionaryConstraintFrom constructs a [DictionaryConstraint] from an unsafe.Pointer.
//
// The constraint on the keys for a dictionary feature.
func DictionaryConstraintFrom(ptr unsafe.Pointer) DictionaryConstraint {
	return DictionaryConstraint{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DictionaryConstraintClass) Alloc() DictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DictionaryConstraintClass) New() DictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DictionaryConstraint) Init() DictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DictionaryConstraint) Autorelease() DictionaryConstraint {
	rv := objc.Send[DictionaryConstraint](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDictionaryConstraint creates a new DictionaryConstraint instance.
func NewDictionaryConstraint() DictionaryConstraint {
	return getDictionaryConstraintClass().New()
}




