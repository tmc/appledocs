// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ValueTransformer] class.
var (
	valueTransformerClass     _ValueTransformerClass
	valueTransformerClassOnce sync.Once
)

func getValueTransformerClass() _ValueTransformerClass {
	valueTransformerClassOnce.Do(func() {
		valueTransformerClass = _ValueTransformerClass{objc.GetClass("NSValueTransformer")}
	})
	return valueTransformerClass
}

type _ValueTransformerClass struct {
	class objc.Class
}

// An interface definition for the [ValueTransformer] class.
type IValueTransformer interface {
	objectivec.IObject
}

// An abstract class used to transform values from one representation to another.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/ValueTransformer
type ValueTransformer struct {
	objectivec.Object
}

// ValueTransformerFrom constructs a [ValueTransformer] from an unsafe.Pointer.
//
// An abstract class used to transform values from one representation to another.
func ValueTransformerFrom(ptr unsafe.Pointer) ValueTransformer {
	return ValueTransformer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _ValueTransformerClass) Alloc() ValueTransformer {
	rv := objc.Send[ValueTransformer](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _ValueTransformerClass) New() ValueTransformer {
	rv := objc.Send[ValueTransformer](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ ValueTransformer) Init() ValueTransformer {
	rv := objc.Send[ValueTransformer](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ ValueTransformer) Autorelease() ValueTransformer {
	rv := objc.Send[ValueTransformer](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewValueTransformer creates a new ValueTransformer instance.
func NewValueTransformer() ValueTransformer {
	return getValueTransformerClass().New()
}




