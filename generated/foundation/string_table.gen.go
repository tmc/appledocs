// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [stringTable] class.
var (
	StringTableClass     _stringTableClass
	StringTableClassOnce sync.Once
)

func getstringTableClass() _stringTableClass {
	StringTableClassOnce.Do(func() {
		StringTableClass = _stringTableClass{objc.GetClass("stringTable")}
	})
	return StringTableClass
}

type _stringTableClass struct {
	class objc.Class
}





// An interface definition for the [stringTable] class.
type IstringTable interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _stringTableClass) Alloc() stringTable {
	rv := objc.Send[stringTable](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _stringTableClass) New() stringTable {
	rv := objc.Send[stringTable](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ stringTable) Init() stringTable {
	rv := objc.Send[stringTable](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ stringTable) Autorelease() stringTable {
	rv := objc.Send[stringTable](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewstringTable creates a new stringTable instance.
func NewstringTable() stringTable {
	return getstringTableClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/stringTable
type stringTable struct {
	objectivec.Object
}

// stringTableFrom constructs a [stringTable] from an unsafe.Pointer.
func stringTableFrom(ptr unsafe.Pointer) stringTable {
	return stringTable{objectivec.Object{objc.ID(ptr)}}
}































