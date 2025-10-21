// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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


// The label for the contact’s child.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationchild
func (c_ CNContactRelation) CNLabelContactRelationChild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationChild"))
	return rv
}

// The label for the contact’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughter
func (c_ CNContactRelation) CNLabelContactRelationDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationDaughter"))
	return rv
}

// The label for the contact’s parent.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparent
func (c_ CNContactRelation) CNLabelContactRelationParent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParent"))
	return rv
}

// The label for the contact’s mother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmother
func (c_ CNContactRelation) CNLabelContactRelationMother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMother"))
	return rv
}

// The label for the contact’s spouse.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationspouse
func (c_ CNContactRelation) CNLabelContactRelationSpouse() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSpouse"))
	return rv
}

// The label for the contact’s manager.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmanager
func (c_ CNContactRelation) CNLabelContactRelationManager() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationManager"))
	return rv
}

// The label for the contact’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationson
func (c_ CNContactRelation) CNLabelContactRelationSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSon"))
	return rv
}

// The label for the contact’s father.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfather
func (c_ CNContactRelation) CNLabelContactRelationFather() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFather"))
	return rv
}

// The label for the contact’s friend.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfriend
func (c_ CNContactRelation) CNLabelContactRelationFriend() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFriend"))
	return rv
}

// The label for the contact’s sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsister
func (c_ CNContactRelation) CNLabelContactRelationSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSister"))
	return rv
}

// The label for the contact’s assistant.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationassistant
func (c_ CNContactRelation) CNLabelContactRelationAssistant() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAssistant"))
	return rv
}

// The label for the contact’s brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrother
func (c_ CNContactRelation) CNLabelContactRelationBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrother"))
	return rv
}

// The name of the related contact.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactrelation/name
func (c_ CNContactRelation) Name() string {
	rv := objc.Send[string](c_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the related contact.

//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactrelation/name
func (c_ CNContactRelation) SetName(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), objc.String(value))
}

// The label for the contact’s partner.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationpartner
func (c_ CNContactRelation) CNLabelContactRelationPartner() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationPartner"))
	return rv
}



