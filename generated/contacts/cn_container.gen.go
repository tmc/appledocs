// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNContainer] class.
var (
	CNContainerClass     _CNContainerClass
	CNContainerClassOnce sync.Once
)

func getCNContainerClass() _CNContainerClass {
	CNContainerClassOnce.Do(func() {
		CNContainerClass = _CNContainerClass{objc.GetClass("CNContainer")}
	})
	return CNContainerClass
}

type _CNContainerClass struct {
	class objc.Class
}

// An interface definition for the [CNContainer] class.
type ICNContainer interface {
	objectivec.IObject
	// properties:
	Identifier() string /* primitive/slice/pointer. */
	SetIdentifier(value string /* primitive/slice/pointer. */)
	Name() string /* primitive/slice/pointer. */
	SetName(value string /* primitive/slice/pointer. */)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
	CNContainerIdentifierKey() string /* primitive/slice/pointer. */
	CNContainerNameKey() string /* primitive/slice/pointer. */
	CNContainerTypeKey() string /* primitive/slice/pointer. */
	// methods:
}

// An immutable object that represents a collection of contacts.
//
// A contact can be in only one container. CardDAV accounts usually have only one container whereas Exchange accounts may have multiple containers, where each container represents an Exchange folder. objects are thread-safe, and you may access their properties from any thread of your app.


// An immutable object that represents a collection of contacts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer
type CNContainer struct {
	objectivec.Object
}

// CNContainerFrom constructs a [CNContainer] from an unsafe.Pointer.
//
// An immutable object that represents a collection of contacts.
func CNContainerFrom(ptr unsafe.Pointer) CNContainer {
	return CNContainer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNContainerClass) Alloc() CNContainer {
	rv := objc.Send[CNContainer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNContainerClass) New() CNContainer {
	rv := objc.Send[CNContainer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNContainer) Init() CNContainer {
	rv := objc.Send[CNContainer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNContainer) Autorelease() CNContainer {
	rv := objc.Send[CNContainer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNContainer creates a new CNContainer instance.
func NewCNContainer() CNContainer {
	return getCNContainerClass().New()
}



// The unique identifier for a contacts container on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/identifier
func (c_ CNContainer) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("identifier"))
	return rv
}


// The unique identifier for a contacts container on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/identifier
func (c_ CNContainer) SetIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The name of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/name
func (c_ CNContainer) Name() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("name"))
	return rv
}


// The name of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/name
func (c_ CNContainer) SetName(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), objc.String(value))
}


// The type of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/type
func (c_ CNContainer) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("type"))
	return rv
}


// The type of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/type
func (c_ CNContainer) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setType:"), value)
}


// The identifier key of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontaineridentifierkey
func (c_ CNContainer) CNContainerIdentifierKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNContainerIdentifierKey"))
	return rv
}


// The name of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainernamekey
func (c_ CNContainer) CNContainerNameKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNContainerNameKey"))
	return rv
}


// The type of the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainertypekey
func (c_ CNContainer) CNContainerTypeKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNContainerTypeKey"))
	return rv
}



