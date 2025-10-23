// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNGroup] class.
var (
	CNGroupClass     _CNGroupClass
	CNGroupClassOnce sync.Once
)

func getCNGroupClass() _CNGroupClass {
	CNGroupClassOnce.Do(func() {
		CNGroupClass = _CNGroupClass{objc.GetClass("CNGroup")}
	})
	return CNGroupClass
}

type _CNGroupClass struct {
	class objc.Class
}

// An interface definition for the [CNGroup] class.
type ICNGroup interface {
	objectivec.IObject
	// properties:
	Identifier() string /* primitive/slice/pointer. */
	SetIdentifier(value string /* primitive/slice/pointer. */)
	Name() string /* primitive/slice/pointer. */
	SetName(value string /* primitive/slice/pointer. */)
	CNGroupIdentifierKey() string /* primitive/slice/pointer. */
	CNGroupNameKey() string /* primitive/slice/pointer. */
	// methods:
}

// An immutable object that represents a group of contacts.
//
// Contacts may be members of one or more groups, depending upon their accounts. objects are thread-safe, and you may access their properties from any thread of your app.


// An immutable object that represents a group of contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNGroup
type CNGroup struct {
	objectivec.Object
}

// CNGroupFrom constructs a [CNGroup] from an unsafe.Pointer.
//
// An immutable object that represents a group of contacts.
func CNGroupFrom(ptr unsafe.Pointer) CNGroup {
	return CNGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNGroupClass) Alloc() CNGroup {
	rv := objc.Send[CNGroup](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNGroupClass) New() CNGroup {
	rv := objc.Send[CNGroup](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNGroup) Init() CNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNGroup) Autorelease() CNGroup {
	rv := objc.Send[CNGroup](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNGroup creates a new CNGroup instance.
func NewCNGroup() CNGroup {
	return getCNGroupClass().New()
}



// The unique identifier for a group on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/identifier
func (c_ CNGroup) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("identifier"))
	return rv
}


// The unique identifier for a group on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/identifier
func (c_ CNGroup) SetIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The name of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/name
func (c_ CNGroup) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("name"))
	return rv
}


// The name of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroup/name
func (c_ CNGroup) SetName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), objc.String(value))
}


// The identifier of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroupidentifierkey
func (c_ CNGroup) CNGroupIdentifierKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNGroupIdentifierKey"))
	return rv
}


// The name of the group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cngroupnamekey
func (c_ CNGroup) CNGroupNameKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNGroupNameKey"))
	return rv
}



