// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [attributes] class.
var (
	AttributesClass     _attributesClass
	AttributesClassOnce sync.Once
)

func getattributesClass() _attributesClass {
	AttributesClassOnce.Do(func() {
		AttributesClass = _attributesClass{objc.GetClass("attributes")}
	})
	return AttributesClass
}

type _attributesClass struct {
	class objc.Class
}

// An interface definition for the [attributes] class.
type Iattributes interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODRecordMap/attributes-c.ivar
type attributes struct {
	objectivec.Object
}

// attributesFrom constructs a [attributes] from an unsafe.Pointer.
func attributesFrom(ptr unsafe.Pointer) attributes {
	return attributes{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _attributesClass) Alloc() attributes {
	rv := objc.Send[attributes](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _attributesClass) New() attributes {
	rv := objc.Send[attributes](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ attributes) Init() attributes {
	rv := objc.Send[attributes](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ attributes) Autorelease() attributes {
	rv := objc.Send[attributes](a_.ID, objc.Sel("autorelease"))
	return rv
}

// Newattributes creates a new attributes instance.
func Newattributes() attributes {
	return getattributesClass().New()
}




