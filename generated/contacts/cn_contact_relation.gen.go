// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	CNLabelContactRelationAssistant() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationChild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFather() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFriend() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationManager() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationMother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParent() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationPartner() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSpouse() objc.IObject /* cross-framework: NSString */
	// methods:
}

// An immutable object that represents the relationship between one contact to another.
//
// objects are thread-safe, and you may access their properties from any thread of your app.


// An immutable object that represents the relationship between one contact to another.
//
// [Full Topic]
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



// The name of the related contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactrelation/name
func (c_ CNContactRelation) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("name"))
	return rv
}


// The name of the related contact.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontactrelation/name
func (c_ CNContactRelation) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setName:"), value)
}


// The label for the contact’s assistant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationassistant
func (c_ CNContactRelation) CNLabelContactRelationAssistant() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAssistant"))
	return rv
}


// The label for the contact’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrother
func (c_ CNContactRelation) CNLabelContactRelationBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrother"))
	return rv
}


// The label for the contact’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationchild
func (c_ CNContactRelation) CNLabelContactRelationChild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationChild"))
	return rv
}


// The label for the contact’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughter
func (c_ CNContactRelation) CNLabelContactRelationDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationDaughter"))
	return rv
}


// The label for the contact’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfather
func (c_ CNContactRelation) CNLabelContactRelationFather() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFather"))
	return rv
}


// The label for the contact’s friend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfriend
func (c_ CNContactRelation) CNLabelContactRelationFriend() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFriend"))
	return rv
}


// The label for the contact’s manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmanager
func (c_ CNContactRelation) CNLabelContactRelationManager() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationManager"))
	return rv
}


// The label for the contact’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmother
func (c_ CNContactRelation) CNLabelContactRelationMother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationMother"))
	return rv
}


// The label for the contact’s parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparent
func (c_ CNContactRelation) CNLabelContactRelationParent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParent"))
	return rv
}


// The label for the contact’s partner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationpartner
func (c_ CNContactRelation) CNLabelContactRelationPartner() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationPartner"))
	return rv
}


// The label for the contact’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsister
func (c_ CNContactRelation) CNLabelContactRelationSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSister"))
	return rv
}


// The label for the contact’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationson
func (c_ CNContactRelation) CNLabelContactRelationSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSon"))
	return rv
}


// The label for the contact’s spouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationspouse
func (c_ CNContactRelation) CNLabelContactRelationSpouse() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSpouse"))
	return rv
}



