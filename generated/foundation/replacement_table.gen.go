// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [replacementTable] class.
var (
	ReplacementTableClass     _replacementTableClass
	ReplacementTableClassOnce sync.Once
)

func getreplacementTableClass() _replacementTableClass {
	ReplacementTableClassOnce.Do(func() {
		ReplacementTableClass = _replacementTableClass{objc.GetClass("replacementTable")}
	})
	return ReplacementTableClass
}

type _replacementTableClass struct {
	class objc.Class
}





// An interface definition for the [replacementTable] class.
type IreplacementTable interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _replacementTableClass) Alloc() replacementTable {
	rv := objc.Send[replacementTable](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _replacementTableClass) New() replacementTable {
	rv := objc.Send[replacementTable](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ replacementTable) Init() replacementTable {
	rv := objc.Send[replacementTable](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ replacementTable) Autorelease() replacementTable {
	rv := objc.Send[replacementTable](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewreplacementTable creates a new replacementTable instance.
func NewreplacementTable() replacementTable {
	return getreplacementTableClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArchiver/replacementTable
type replacementTable struct {
	objectivec.Object
}

// replacementTableFrom constructs a [replacementTable] from an unsafe.Pointer.
func replacementTableFrom(ptr unsafe.Pointer) replacementTable {
	return replacementTable{objectivec.Object{objc.ID(ptr)}}
}































