// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CharacterSet] class.
var (
	CharacterSetClass     _CharacterSetClass
	CharacterSetClassOnce sync.Once
)

func getCharacterSetClass() _CharacterSetClass {
	CharacterSetClassOnce.Do(func() {
		CharacterSetClass = _CharacterSetClass{objc.GetClass("NSCharacterSet")}
	})
	return CharacterSetClass
}

type _CharacterSetClass struct {
	class objc.Class
}

// An interface definition for the [CharacterSet] class.
type ICharacterSet interface {
	objectivec.IObject
}

// A parent class referenced by other Foundation classes.


// A parent class referenced by other Foundation classes. [Full Topic]
type CharacterSet struct {
	objectivec.Object
}

// CharacterSetFrom constructs a [CharacterSet] from an unsafe.Pointer.
//
// A parent class referenced by other Foundation classes.
func CharacterSetFrom(ptr unsafe.Pointer) CharacterSet {
	return CharacterSet{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CharacterSetClass) Alloc() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CharacterSetClass) New() CharacterSet {
	rv := objc.Send[CharacterSet](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CharacterSet) Init() CharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CharacterSet) Autorelease() CharacterSet {
	rv := objc.Send[CharacterSet](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCharacterSet creates a new CharacterSet instance.
func NewCharacterSet() CharacterSet {
	return getCharacterSetClass().New()
}




