// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CNMutableGroup] class.
var (
	CNMutableGroupClass     _CNMutableGroupClass
	CNMutableGroupClassOnce sync.Once
)

func getCNMutableGroupClass() _CNMutableGroupClass {
	CNMutableGroupClassOnce.Do(func() {
		CNMutableGroupClass = _CNMutableGroupClass{objc.GetClass("CNMutableGroup")}
	})
	return CNMutableGroupClass
}

type _CNMutableGroupClass struct {
	class objc.Class
}

// An interface definition for the [CNMutableGroup] class.
type ICNMutableGroup interface {
	ICNGroup
	Name() string
	SetName(value string)
}

// A mutable object that represents a group of contacts.
//
// Contacts may be members of one or more groups, depending upon the accounts they come from. The class is not a thread-safe class.


// A mutable object that represents a group of contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNMutableGroup
type CNMutableGroup struct {
	CNGroup
}

// CNMutableGroupFrom constructs a [CNMutableGroup] from an unsafe.Pointer.
//
// A mutable object that represents a group of contacts.
func CNMutableGroupFrom(ptr unsafe.Pointer) CNMutableGroup {
	return CNMutableGroup{
		CNGroup: CNGroupFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CNMutableGroupClass) Alloc() CNMutableGroup {
	rv := objc.Send[CNMutableGroup](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNMutableGroupClass) New() CNMutableGroup {
	rv := objc.Send[CNMutableGroup](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNMutableGroup) Init() CNMutableGroup {
	rv := objc.Send[CNMutableGroup](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNMutableGroup) Autorelease() CNMutableGroup {
	rv := objc.Send[CNMutableGroup](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNMutableGroup creates a new CNMutableGroup instance.
func NewCNMutableGroup() CNMutableGroup {
	return getCNMutableGroupClass().New()
}



// The name of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablegroup/name
func (c_ CNMutableGroup) Name() string {
	rv := objc.Send[string](c_.ID, objc.Sel("name"))
	return rv
}


// The name of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnmutablegroup/name
func (c_ CNMutableGroup) SetName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), objc.String(value))
}



