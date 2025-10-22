// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	Identifier() string
	Name() string
	Type() CNContainerType
	CNContainerIdentifierKey() string
	CNContainerNameKey() string
	CNContainerTypeKey() string
}

// An immutable object that represents a collection of contacts.
//
// A contact can be in only one container. CardDAV accounts usually have only one container whereas Exchange accounts may have multiple containers, where each container represents an Exchange folder. objects are thread-safe, and you may access their properties from any thread of your app.
//
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


// Returns a predicate to find the container of the specified contact.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/predicateForContainerOfContact(withIdentifier:)
func (cc _CNContainerClass) PredicateForContainerOfContactWithIdentifier(contactIdentifier string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContainerOfContactWithIdentifier:"), objc.String(contactIdentifier))
	return rv
}

// Returns a predicate to find the container of the specified group.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/predicateForContainerOfGroup(withIdentifier:)
func (cc _CNContainerClass) PredicateForContainerOfGroupWithIdentifier(groupIdentifier string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContainerOfGroupWithIdentifier:"), objc.String(groupIdentifier))
	return rv
}

// Returns a predicate to find the containers with the specified identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/predicateForContainers(withIdentifiers:)
func (cc _CNContainerClass) PredicateForContainersWithIdentifiers(identifiers []string) foundation.Predicate {
	rv := objc.Send[foundation.Predicate](objc.ID(cc.class), objc.Sel("predicateForContainersWithIdentifiers:"), identifiers)
	return rv
}

// The unique identifier for a contacts container on the device.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/identifier
func (c_ CNContainer) Identifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("identifier"))
	return rv
}

// The name of the container.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/name
func (c_ CNContainer) Name() string {
	rv := objc.Send[string](c_.ID, objc.Sel("name"))
	return rv
}

// The type of the container.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNContainer/type
func (c_ CNContainer) Type() CNContainerType {
	rv := objc.Send[CNContainerType](c_.ID, objc.Sel("type"))
	return rv
}

// The identifier key of the container.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontaineridentifierkey
func (c_ CNContainer) CNContainerIdentifierKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNContainerIdentifierKey"))
	return rv
}

// The name of the container.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainernamekey
func (c_ CNContainer) CNContainerNameKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNContainerNameKey"))
	return rv
}

// The type of the container.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainertypekey
func (c_ CNContainer) CNContainerTypeKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNContainerTypeKey"))
	return rv
}



