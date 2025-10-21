// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNContactRelation] class.
var (
	CNContactRelationClass     _CNContactRelationClass
	CNContactRelationClassOnce sync.Once
)

func getCNContactRelationClass() _CNContactRelationClass {
	CNContactRelationClassOnce.Do(func() {
		CNContactRelationClass = _CNContactRelationClass{objc.GetClass("CNContactRelation")}
	})
	return CNContactRelationClass
}

type _CNContactRelationClass struct {
	class objc.Class
}

// An interface definition for the [CNContactRelation] class.
type ICNContactRelation interface {
	objectivec.IObject
}

// An immutable object that represents the relationship between one contact to another.
//
// objects are thread-safe, and you may access their properties from any thread of your app.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContactRelation
type CNContactRelation struct {
	objectivec.Object
}

// CNContactRelationFrom constructs a [CNContactRelation] from an unsafe.Pointer.
//
// An immutable object that represents the relationship between one contact to another.
func CNContactRelationFrom(ptr unsafe.Pointer) CNContactRelation {
	return CNContactRelation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNContactRelationClass) Alloc() CNContactRelation {
	rv := objc.Send[CNContactRelation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNContactRelationClass) New() CNContactRelation {
	rv := objc.Send[CNContactRelation](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContactRelation) Init() CNContactRelation {
	rv := objc.Send[CNContactRelation](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContactRelation) Autorelease() CNContactRelation {
	rv := objc.Send[CNContactRelation](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContactRelation creates a new CNContactRelation instance.
func NewCNContactRelation() CNContactRelation {
	return getCNContactRelationClass().New()
}




