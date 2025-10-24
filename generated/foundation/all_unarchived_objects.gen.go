// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [allUnarchivedObjects] class.
var (
	AllUnarchivedObjectsClass     _allUnarchivedObjectsClass
	AllUnarchivedObjectsClassOnce sync.Once
)

func getallUnarchivedObjectsClass() _allUnarchivedObjectsClass {
	AllUnarchivedObjectsClassOnce.Do(func() {
		AllUnarchivedObjectsClass = _allUnarchivedObjectsClass{objc.GetClass("allUnarchivedObjects")}
	})
	return AllUnarchivedObjectsClass
}

type _allUnarchivedObjectsClass struct {
	class objc.Class
}





// An interface definition for the [allUnarchivedObjects] class.
type IallUnarchivedObjects interface {
	objectivec.IObject
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _allUnarchivedObjectsClass) Alloc() allUnarchivedObjects {
	rv := objc.Send[allUnarchivedObjects](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _allUnarchivedObjectsClass) New() allUnarchivedObjects {
	rv := objc.Send[allUnarchivedObjects](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ allUnarchivedObjects) Init() allUnarchivedObjects {
	rv := objc.Send[allUnarchivedObjects](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ allUnarchivedObjects) Autorelease() allUnarchivedObjects {
	rv := objc.Send[allUnarchivedObjects](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewallUnarchivedObjects creates a new allUnarchivedObjects instance.
func NewallUnarchivedObjects() allUnarchivedObjects {
	return getallUnarchivedObjectsClass().New()
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUnarchiver/allUnarchivedObjects
type allUnarchivedObjects struct {
	objectivec.Object
}

// allUnarchivedObjectsFrom constructs a [allUnarchivedObjects] from an unsafe.Pointer.
func allUnarchivedObjectsFrom(ptr unsafe.Pointer) allUnarchivedObjects {
	return allUnarchivedObjects{objectivec.Object{objc.ID(ptr)}}
}































