// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [classInfoImported] class.
var (
	ClassInfoImportedClass     _classInfoImportedClass
	ClassInfoImportedClassOnce sync.Once
)

func getclassInfoImportedClass() _classInfoImportedClass {
	ClassInfoImportedClassOnce.Do(func() {
		ClassInfoImportedClass = _classInfoImportedClass{objc.GetClass("classInfoImported")}
	})
	return ClassInfoImportedClass
}

type _classInfoImportedClass struct {
	class objc.Class
}

// An interface definition for the [classInfoImported] class.
type IclassInfoImported interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/classInfoImported
type classInfoImported struct {
	objectivec.Object
}

// classInfoImportedFrom constructs a [classInfoImported] from an unsafe.Pointer.
func classInfoImportedFrom(ptr unsafe.Pointer) classInfoImported {
	return classInfoImported{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _classInfoImportedClass) Alloc() classInfoImported {
	rv := objc.Send[classInfoImported](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _classInfoImportedClass) New() classInfoImported {
	rv := objc.Send[classInfoImported](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ classInfoImported) Init() classInfoImported {
	rv := objc.Send[classInfoImported](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ classInfoImported) Autorelease() classInfoImported {
	rv := objc.Send[classInfoImported](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewclassInfoImported creates a new classInfoImported instance.
func NewclassInfoImported() classInfoImported {
	return getclassInfoImportedClass().New()
}




