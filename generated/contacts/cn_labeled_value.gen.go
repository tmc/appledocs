// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNLabeledValue] class.
var (
	CNLabeledValueClass     _CNLabeledValueClass
	CNLabeledValueClassOnce sync.Once
)

func getCNLabeledValueClass() _CNLabeledValueClass {
	CNLabeledValueClassOnce.Do(func() {
		CNLabeledValueClass = _CNLabeledValueClass{objc.GetClass("CNLabeledValue")}
	})
	return CNLabeledValueClass
}

type _CNLabeledValueClass struct {
	class objc.Class
}

// An interface definition for the [CNLabeledValue] class.
type ICNLabeledValue interface {
	objectivec.IObject
	// properties:
	CNLabelContactRelationAssistant() string /* primitive/slice/pointer. */
	CNLabelContactRelationAunt() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntFathersBrothersWife() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntFathersElderBrothersWife() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntFathersElderSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntFathersSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntFathersYoungerBrothersWife() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntFathersYoungerSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntMothersBrothersWife() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntMothersElderSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntMothersSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntMothersYoungerSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntParentsElderSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntParentsSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationAuntParentsYoungerSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationBoyfriend() string /* primitive/slice/pointer. */
	CNLabelContactRelationBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationBrotherInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationBrotherInLawElderSistersHusband() string /* primitive/slice/pointer. */
	CNLabelContactRelationBrotherInLawHusbandsBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationBrotherInLawHusbandsSistersHusband() string /* primitive/slice/pointer. */
	CNLabelContactRelationBrotherInLawSistersHusband() string /* primitive/slice/pointer. */
	CNLabelContactRelationBrotherInLawSpousesBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationBrotherInLawWifesBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationBrotherInLawWifesSistersHusband() string /* primitive/slice/pointer. */
	CNLabelContactRelationBrotherInLawYoungerSistersHusband() string /* primitive/slice/pointer. */
	CNLabelContactRelationChild() string /* primitive/slice/pointer. */
	CNLabelContactRelationChildInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationCoBrotherInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationCoFatherInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationCoMotherInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationCoParentInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationCoSiblingInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationCoSisterInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationColleague() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousin() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinFathersBrothersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinFathersBrothersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinFathersSistersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinFathersSistersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinGrandparentsSiblingsChild() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinGrandparentsSiblingsDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinGrandparentsSiblingsSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinMothersBrothersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinMothersBrothersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinMothersSistersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinMothersSistersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinOrSiblingsChild() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinParentsSiblingsChild() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinParentsSiblingsDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationCousinParentsSiblingsSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationDaughterInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationDaughterInLawOrSisterInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationDaughterInLawOrStepdaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderBrotherInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousin() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinFathersBrothersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinFathersBrothersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinFathersSistersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinFathersSistersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinMothersBrothersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinMothersBrothersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinMothersSistersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinMothersSistersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinParentsSiblingsDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderCousinParentsSiblingsSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderSiblingInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationElderSisterInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationEldestBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationEldestSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationFather() string /* primitive/slice/pointer. */
	CNLabelContactRelationFatherInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationFatherInLawHusbandsFather() string /* primitive/slice/pointer. */
	CNLabelContactRelationFatherInLawOrStepfather() string /* primitive/slice/pointer. */
	CNLabelContactRelationFatherInLawWifesFather() string /* primitive/slice/pointer. */
	CNLabelContactRelationFemaleCousin() string /* primitive/slice/pointer. */
	CNLabelContactRelationFemaleFriend() string /* primitive/slice/pointer. */
	CNLabelContactRelationFemalePartner() string /* primitive/slice/pointer. */
	CNLabelContactRelationFriend() string /* primitive/slice/pointer. */
	CNLabelContactRelationGirlfriend() string /* primitive/slice/pointer. */
	CNLabelContactRelationGirlfriendOrBoyfriend() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandaunt() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandchild() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandchildOrSiblingsChild() string /* primitive/slice/pointer. */
	CNLabelContactRelationGranddaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationGranddaughterDaughtersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationGranddaughterOrNiece() string /* primitive/slice/pointer. */
	CNLabelContactRelationGranddaughterSonsDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandfather() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandfatherFathersFather() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandfatherMothersFather() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandmother() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandmotherFathersMother() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandmotherMothersMother() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandnephew() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandnephewBrothersGrandson() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandnephewSistersGrandson() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandniece() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandnieceBrothersGranddaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandnieceSistersGranddaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandparent() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandson() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandsonDaughtersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandsonOrNephew() string /* primitive/slice/pointer. */
	CNLabelContactRelationGrandsonSonsSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationGranduncle() string /* primitive/slice/pointer. */
	CNLabelContactRelationGreatGrandchild() string /* primitive/slice/pointer. */
	CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild() string /* primitive/slice/pointer. */
	CNLabelContactRelationGreatGranddaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationGreatGrandfather() string /* primitive/slice/pointer. */
	CNLabelContactRelationGreatGrandmother() string /* primitive/slice/pointer. */
	CNLabelContactRelationGreatGrandparent() string /* primitive/slice/pointer. */
	CNLabelContactRelationGreatGrandson() string /* primitive/slice/pointer. */
	CNLabelContactRelationHusband() string /* primitive/slice/pointer. */
	CNLabelContactRelationMaleCousin() string /* primitive/slice/pointer. */
	CNLabelContactRelationMaleFriend() string /* primitive/slice/pointer. */
	CNLabelContactRelationMalePartner() string /* primitive/slice/pointer. */
	CNLabelContactRelationManager() string /* primitive/slice/pointer. */
	CNLabelContactRelationMother() string /* primitive/slice/pointer. */
	CNLabelContactRelationMotherInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationMotherInLawHusbandsMother() string /* primitive/slice/pointer. */
	CNLabelContactRelationMotherInLawOrStepmother() string /* primitive/slice/pointer. */
	CNLabelContactRelationMotherInLawWifesMother() string /* primitive/slice/pointer. */
	CNLabelContactRelationNephew() string /* primitive/slice/pointer. */
	CNLabelContactRelationNephewBrothersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationNephewOrCousin() string /* primitive/slice/pointer. */
	CNLabelContactRelationNephewSistersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationNiece() string /* primitive/slice/pointer. */
	CNLabelContactRelationNieceBrothersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationNieceOrCousin() string /* primitive/slice/pointer. */
	CNLabelContactRelationNieceSistersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationParent() string /* primitive/slice/pointer. */
	CNLabelContactRelationParentInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationParentsElderSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationParentsSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationParentsSiblingFathersElderSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationParentsSiblingFathersSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationParentsSiblingFathersYoungerSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationParentsSiblingMothersElderSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationParentsSiblingMothersSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationParentsSiblingMothersYoungerSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationParentsYoungerSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationPartner() string /* primitive/slice/pointer. */
	CNLabelContactRelationSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationSiblingInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationSiblingsChild() string /* primitive/slice/pointer. */
	CNLabelContactRelationSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationSisterInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationSisterInLawBrothersWife() string /* primitive/slice/pointer. */
	CNLabelContactRelationSisterInLawElderBrothersWife() string /* primitive/slice/pointer. */
	CNLabelContactRelationSisterInLawHusbandsBrothersWife() string /* primitive/slice/pointer. */
	CNLabelContactRelationSisterInLawHusbandsSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationSisterInLawSpousesSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationSisterInLawWifesBrothersWife() string /* primitive/slice/pointer. */
	CNLabelContactRelationSisterInLawWifesSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationSisterInLawYoungerBrothersWife() string /* primitive/slice/pointer. */
	CNLabelContactRelationSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationSonInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationSonInLawOrBrotherInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationSonInLawOrStepson() string /* primitive/slice/pointer. */
	CNLabelContactRelationSpouse() string /* primitive/slice/pointer. */
	CNLabelContactRelationStepbrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationStepchild() string /* primitive/slice/pointer. */
	CNLabelContactRelationStepdaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationStepfather() string /* primitive/slice/pointer. */
	CNLabelContactRelationStepmother() string /* primitive/slice/pointer. */
	CNLabelContactRelationStepparent() string /* primitive/slice/pointer. */
	CNLabelContactRelationStepsister() string /* primitive/slice/pointer. */
	CNLabelContactRelationStepson() string /* primitive/slice/pointer. */
	CNLabelContactRelationTeacher() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncle() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleFathersBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleFathersElderBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleFathersElderSistersHusband() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleFathersSistersHusband() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleFathersYoungerBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleFathersYoungerSistersHusband() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleMothersBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleMothersElderBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleMothersSistersHusband() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleMothersYoungerBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleParentsBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleParentsElderBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationUncleParentsYoungerBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationWife() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerBrotherInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousin() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinFathersBrothersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinFathersBrothersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinFathersSistersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinFathersSistersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinMothersBrothersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinMothersBrothersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinMothersSistersDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinMothersSistersSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinParentsSiblingsDaughter() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerCousinParentsSiblingsSon() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerSibling() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerSiblingInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerSister() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungerSisterInLaw() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungestBrother() string /* primitive/slice/pointer. */
	CNLabelContactRelationYoungestSister() string /* primitive/slice/pointer. */
	CNLabelDateAnniversary() string /* primitive/slice/pointer. */
	CNLabelEmailiCloud() string /* primitive/slice/pointer. */
	CNLabelHome() string /* primitive/slice/pointer. */
	CNLabelOther() string /* primitive/slice/pointer. */
	CNLabelPhoneNumberAppleWatch() string /* primitive/slice/pointer. */
	CNLabelPhoneNumberHomeFax() string /* primitive/slice/pointer. */
	CNLabelPhoneNumberMain() string /* primitive/slice/pointer. */
	CNLabelPhoneNumberMobile() string /* primitive/slice/pointer. */
	CNLabelPhoneNumberOtherFax() string /* primitive/slice/pointer. */
	CNLabelPhoneNumberPager() string /* primitive/slice/pointer. */
	CNLabelPhoneNumberWorkFax() string /* primitive/slice/pointer. */
	CNLabelPhoneNumberiPhone() string /* primitive/slice/pointer. */
	CNLabelSchool() string /* primitive/slice/pointer. */
	CNLabelURLAddressHomePage() string /* primitive/slice/pointer. */
	CNLabelWork() string /* primitive/slice/pointer. */
	Identifier() string /* primitive/slice/pointer. */
	SetIdentifier(value string /* primitive/slice/pointer. */)
	Label() string /* primitive/slice/pointer. */
	SetLabel(value string /* primitive/slice/pointer. */)
	Value() unsafe.Pointer
	SetValue(value unsafe.Pointer)
	// methods:
}

// An immutable object that combines a contact property value with a label that describes that property.
//
// Labels describe the context for a property. For example, the label for a phone number indicates whether it corresponds to the user’s home, work, or iPhone number. objects are thread-safe, and you can access their properties from any thread of your app.


// An immutable object that combines a contact property value with a label that describes that property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue
type CNLabeledValue struct {
	objectivec.Object
}

// CNLabeledValueFrom constructs a [CNLabeledValue] from an unsafe.Pointer.
//
// An immutable object that combines a contact property value with a label that describes that property.
func CNLabeledValueFrom(ptr unsafe.Pointer) CNLabeledValue {
	return CNLabeledValue{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNLabeledValueClass) Alloc() CNLabeledValue {
	rv := objc.Send[CNLabeledValue](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNLabeledValueClass) New() CNLabeledValue {
	rv := objc.Send[CNLabeledValue](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNLabeledValue) Init() CNLabeledValue {
	rv := objc.Send[CNLabeledValue](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNLabeledValue) Autorelease() CNLabeledValue {
	rv := objc.Send[CNLabeledValue](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNLabeledValue creates a new CNLabeledValue instance.
func NewCNLabeledValue() CNLabeledValue {
	return getCNLabeledValueClass().New()
}



// Returns a localized string for the specified label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/localizedString(forLabel:)
func (cc _CNLabeledValueClass) LocalizedStringForLabel(label string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](objc.ID(cc.class), objc.Sel("localizedStringForLabel:"), objc.String(label))
	return rv
}


// The label for the contact’s assistant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationassistant
func (c_ CNLabeledValue) CNLabelContactRelationAssistant() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAssistant"))
	return rv
}


// The label for the contact’s aunt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationaunt
func (c_ CNLabeledValue) CNLabelContactRelationAunt() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAunt"))
	return rv
}


// The label for the contact’s father’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersBrothersWife() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersBrothersWife"))
	return rv
}


// The label for the contact’s father’s elder brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherselderbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersElderBrothersWife() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersElderBrothersWife"))
	return rv
}


// The label for the contact’s father’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherseldersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersElderSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersElderSister"))
	return rv
}


// The label for the contact’s father’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherssister
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersSister"))
	return rv
}


// The label for the contact’s father’s younger brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersyoungerbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersYoungerBrothersWife() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersYoungerBrothersWife"))
	return rv
}


// The label for the contact’s father’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersYoungerSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersYoungerSister"))
	return rv
}


// The label for the contact’s mother’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmothersbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersBrothersWife() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersBrothersWife"))
	return rv
}


// The label for the contact’s mother’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmotherseldersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersElderSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersElderSister"))
	return rv
}


// The label for the contact’s mother’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmotherssister
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersSister"))
	return rv
}


// The label for the contact’s mother’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmothersyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersYoungerSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersYoungerSister"))
	return rv
}


// The label for the contact’s parent’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentseldersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsElderSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsElderSister"))
	return rv
}


// The label for the contact’s parent’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentssister
func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsSister"))
	return rv
}


// The label for the contact’s parent’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentsyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsYoungerSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsYoungerSister"))
	return rv
}


// The label for the contact’s boyfriend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationboyfriend
func (c_ CNLabeledValue) CNLabelContactRelationBoyfriend() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBoyfriend"))
	return rv
}


// The label for the contact’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrother"))
	return rv
}


// The label for the contact’s brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLaw"))
	return rv
}


// The label for the contact’s elder sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlaweldersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawElderSistersHusband() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawElderSistersHusband"))
	return rv
}


// The label for the contact’s husband’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawhusbandsbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawHusbandsBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawHusbandsBrother"))
	return rv
}


// The label for the contact’s husband’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawhusbandssistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawHusbandsSistersHusband() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawHusbandsSistersHusband"))
	return rv
}


// The label for the contact’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawsistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawSistersHusband() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawSistersHusband"))
	return rv
}


// The label for the contact’s spouse’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawspousesbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawSpousesBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawSpousesBrother"))
	return rv
}


// The label for the contact’s wife’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawwifesbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawWifesBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawWifesBrother"))
	return rv
}


// The label for the contact’s wife’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawwifessistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawWifesSistersHusband() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawWifesSistersHusband"))
	return rv
}


// The label for the contact’s younger sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawyoungersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawYoungerSistersHusband() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawYoungerSistersHusband"))
	return rv
}


// The label for the contact’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationchild
func (c_ CNLabeledValue) CNLabelContactRelationChild() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationChild"))
	return rv
}


// The label for the contact’s child-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationchildinlaw
func (c_ CNLabeledValue) CNLabelContactRelationChildInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationChildInLaw"))
	return rv
}


// The label for the contact’s co-brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcobrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoBrotherInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoBrotherInLaw"))
	return rv
}


// The label for the contact’s co-father-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcofatherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoFatherInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoFatherInLaw"))
	return rv
}


// The label for the contact’s co-mother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcomotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoMotherInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoMotherInLaw"))
	return rv
}


// The label for the contact’s co-parent-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcoparentinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoParentInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoParentInLaw"))
	return rv
}


// The label for the contact’s co-sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcosiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoSiblingInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoSiblingInLaw"))
	return rv
}


// The label for the contact’s co-sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcosisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoSisterInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoSisterInLaw"))
	return rv
}


// The label for the contact’s colleague.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcolleague
func (c_ CNLabeledValue) CNLabelContactRelationColleague() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationColleague"))
	return rv
}


// The label for the contact’s cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousin
func (c_ CNLabeledValue) CNLabelContactRelationCousin() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousin"))
	return rv
}


// The label for the contact’s father’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfathersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersBrothersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersBrothersDaughter"))
	return rv
}


// The label for the contact’s father’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfathersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersBrothersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersBrothersSon"))
	return rv
}


// The label for the contact’s father’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersSistersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersSistersDaughter"))
	return rv
}


// The label for the contact’s father’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersSistersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersSistersSon"))
	return rv
}


// The label for the contact’s grandparent’s sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsChild() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsChild"))
	return rv
}


// The label for the contact’s grandparent’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsDaughter"))
	return rv
}


// The label for the contact’s grandparent’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsSon"))
	return rv
}


// The label for the contact’s mother’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmothersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersBrothersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersBrothersDaughter"))
	return rv
}


// The label for the contact’s mother’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmothersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersBrothersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersBrothersSon"))
	return rv
}


// The label for the contact’s mother’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmotherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersSistersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersSistersDaughter"))
	return rv
}


// The label for the contact’s mother’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmotherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersSistersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersSistersSon"))
	return rv
}


// The label for the contact’s cousin’s or sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinorsiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationCousinOrSiblingsChild() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinOrSiblingsChild"))
	return rv
}


// The label for the contact’s parent’s sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsChild() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsChild"))
	return rv
}


// The label for the contact’s parent’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsDaughter"))
	return rv
}


// The label for the contact’s parent’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsSon"))
	return rv
}


// The label for the contact’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughter
func (c_ CNLabeledValue) CNLabelContactRelationDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationDaughter"))
	return rv
}


// The label for the contact’s daughter-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLaw"))
	return rv
}


// The label for the contact’s daughter-in-law or sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaworsisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLawOrSisterInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLawOrSisterInLaw"))
	return rv
}


// The label for the contact’s daughter-in-law or stepdaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaworstepdaughter
func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLawOrStepdaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLawOrStepdaughter"))
	return rv
}


// The label for the contact’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationelderbrother
func (c_ CNLabeledValue) CNLabelContactRelationElderBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderBrother"))
	return rv
}


// The label for the contact’s elder brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationelderbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationElderBrotherInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderBrotherInLaw"))
	return rv
}


// The label for the contact’s elder cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousin
func (c_ CNLabeledValue) CNLabelContactRelationElderCousin() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousin"))
	return rv
}


// The label for the contact’s father’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfathersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersBrothersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersBrothersDaughter"))
	return rv
}


// The label for the contact’s father’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfathersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersBrothersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersBrothersSon"))
	return rv
}


// The label for the contact’s father’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersSistersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersSistersDaughter"))
	return rv
}


// The label for the contact’s father’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersSistersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersSistersSon"))
	return rv
}


// The label for the contact’s mother’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmothersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersBrothersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersBrothersDaughter"))
	return rv
}


// The label for the contact’s mother’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmothersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersBrothersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersBrothersSon"))
	return rv
}


// The label for the contact’s mother’s sibling’s daughter or father’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssiblingsdaughterorfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter"))
	return rv
}


// The label for the contact’s mother’s sibling’s son or father’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssiblingssonorfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon"))
	return rv
}


// The label for the contact’s mother’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSistersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSistersDaughter"))
	return rv
}


// The label for the contact’s mother’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSistersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSistersSon"))
	return rv
}


// The label for the contact’s parent’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinParentsSiblingsDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinParentsSiblingsDaughter"))
	return rv
}


// The label for the contact’s parent’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinParentsSiblingsSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinParentsSiblingsSon"))
	return rv
}


// The label for the contact’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersibling
func (c_ CNLabeledValue) CNLabelContactRelationElderSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderSibling"))
	return rv
}


// The label for the contact’s elder sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationElderSiblingInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderSiblingInLaw"))
	return rv
}


// The label for the contact’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersister
func (c_ CNLabeledValue) CNLabelContactRelationElderSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderSister"))
	return rv
}


// The label for the contact’s elder sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationElderSisterInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderSisterInLaw"))
	return rv
}


// The label for the contact’s eldest brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldestbrother
func (c_ CNLabeledValue) CNLabelContactRelationEldestBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationEldestBrother"))
	return rv
}


// The label for the contact’s eldest sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldestsister
func (c_ CNLabeledValue) CNLabelContactRelationEldestSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationEldestSister"))
	return rv
}


// The label for the contact’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfather
func (c_ CNLabeledValue) CNLabelContactRelationFather() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFather"))
	return rv
}


// The label for the contact’s father-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLaw"))
	return rv
}


// The label for the contact’s husband’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlawhusbandsfather
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawHusbandsFather() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawHusbandsFather"))
	return rv
}


// The label for the contact’s father-in-law or stepfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlaworstepfather
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawOrStepfather() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawOrStepfather"))
	return rv
}


// The label for the contact’s wife’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlawwifesfather
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawWifesFather() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawWifesFather"))
	return rv
}


// The label for the contact’s female cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalecousin
func (c_ CNLabeledValue) CNLabelContactRelationFemaleCousin() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFemaleCousin"))
	return rv
}


// The label for the contact’s female friend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalefriend
func (c_ CNLabeledValue) CNLabelContactRelationFemaleFriend() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFemaleFriend"))
	return rv
}


// The label for the contact’s female partner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalepartner
func (c_ CNLabeledValue) CNLabelContactRelationFemalePartner() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFemalePartner"))
	return rv
}


// The label for the contact’s friend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfriend
func (c_ CNLabeledValue) CNLabelContactRelationFriend() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFriend"))
	return rv
}


// The label for the contact’s girlfriend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgirlfriend
func (c_ CNLabeledValue) CNLabelContactRelationGirlfriend() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGirlfriend"))
	return rv
}


// The label for the contact’s girlfriend or boyfriend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgirlfriendorboyfriend
func (c_ CNLabeledValue) CNLabelContactRelationGirlfriendOrBoyfriend() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGirlfriendOrBoyfriend"))
	return rv
}


// The label for the contact’s grandaunt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandaunt
func (c_ CNLabeledValue) CNLabelContactRelationGrandaunt() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandaunt"))
	return rv
}


// The label for the contact’s grandchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandchild
func (c_ CNLabeledValue) CNLabelContactRelationGrandchild() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandchild"))
	return rv
}


// The label for the contact’s grandchild or sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandchildorsiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationGrandchildOrSiblingsChild() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandchildOrSiblingsChild"))
	return rv
}


// The label for the contact’s granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughter"))
	return rv
}


// The label for the contact’s daughter’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughterdaughtersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterDaughtersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterDaughtersDaughter"))
	return rv
}


// The label for the contact’s granddaughter or niece.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughterorniece
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterOrNiece() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterOrNiece"))
	return rv
}


// The label for the contact’s son’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughtersonsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterSonsDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterSonsDaughter"))
	return rv
}


// The label for the contact’s grandfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfather
func (c_ CNLabeledValue) CNLabelContactRelationGrandfather() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandfather"))
	return rv
}


// The label for the contact’s father’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfatherfathersfather
func (c_ CNLabeledValue) CNLabelContactRelationGrandfatherFathersFather() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandfatherFathersFather"))
	return rv
}


// The label for the contact’s mother’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfathermothersfather
func (c_ CNLabeledValue) CNLabelContactRelationGrandfatherMothersFather() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandfatherMothersFather"))
	return rv
}


// The label for the contact’s grandmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmother
func (c_ CNLabeledValue) CNLabelContactRelationGrandmother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandmother"))
	return rv
}


// The label for the contact’s father’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmotherfathersmother
func (c_ CNLabeledValue) CNLabelContactRelationGrandmotherFathersMother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandmotherFathersMother"))
	return rv
}


// The label for the contact’s mother’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmothermothersmother
func (c_ CNLabeledValue) CNLabelContactRelationGrandmotherMothersMother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandmotherMothersMother"))
	return rv
}


// The label for the contact’s grandnephew.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephew
func (c_ CNLabeledValue) CNLabelContactRelationGrandnephew() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandnephew"))
	return rv
}


// The label for the contact’s brother’s grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephewbrothersgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGrandnephewBrothersGrandson() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandnephewBrothersGrandson"))
	return rv
}


// The label for the contact’s sister’s grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephewsistersgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGrandnephewSistersGrandson() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandnephewSistersGrandson"))
	return rv
}


// The label for the contact’s grandniece.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniece
func (c_ CNLabeledValue) CNLabelContactRelationGrandniece() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandniece"))
	return rv
}


// The label for the contact’s brother’s granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniecebrothersgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGrandnieceBrothersGranddaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandnieceBrothersGranddaughter"))
	return rv
}


// The label for the contact’s sister’s granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniecesistersgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGrandnieceSistersGranddaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandnieceSistersGranddaughter"))
	return rv
}


// The label for the contact’s grandparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandparent
func (c_ CNLabeledValue) CNLabelContactRelationGrandparent() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandparent"))
	return rv
}


// The label for the contact’s grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGrandson() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandson"))
	return rv
}


// The label for the contact’s daughter’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsondaughtersson
func (c_ CNLabeledValue) CNLabelContactRelationGrandsonDaughtersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandsonDaughtersSon"))
	return rv
}


// The label for the contact’s grandson or nephew.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsonornephew
func (c_ CNLabeledValue) CNLabelContactRelationGrandsonOrNephew() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandsonOrNephew"))
	return rv
}


// The label for the contact’s son’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsonsonsson
func (c_ CNLabeledValue) CNLabelContactRelationGrandsonSonsSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandsonSonsSon"))
	return rv
}


// The label for the contact’s granduncle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranduncle
func (c_ CNLabeledValue) CNLabelContactRelationGranduncle() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGranduncle"))
	return rv
}


// The label for the contact’s grandchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandchild
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandchild() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandchild"))
	return rv
}


// The label for the contact’s grandchild or sibling’s grandchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandchildorsiblingsgrandchild
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild"))
	return rv
}


// The label for the contact’s great-granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGreatGranddaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGranddaughter"))
	return rv
}


// The label for the contact’s great-grandfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandfather
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandfather() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandfather"))
	return rv
}


// The label for the contact’s great-grandmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandmother
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandmother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandmother"))
	return rv
}


// The label for the contact’s great-grandparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandparent
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandparent() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandparent"))
	return rv
}


// The label for the contact’s great-grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandson() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandson"))
	return rv
}


// The label for the contact’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationhusband
func (c_ CNLabeledValue) CNLabelContactRelationHusband() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationHusband"))
	return rv
}


// The label for the contact’s male cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalecousin
func (c_ CNLabeledValue) CNLabelContactRelationMaleCousin() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMaleCousin"))
	return rv
}


// The label for the contact’s male friend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalefriend
func (c_ CNLabeledValue) CNLabelContactRelationMaleFriend() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMaleFriend"))
	return rv
}


// The label for the contact’s male partner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalepartner
func (c_ CNLabeledValue) CNLabelContactRelationMalePartner() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMalePartner"))
	return rv
}


// The label for the contact’s manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmanager
func (c_ CNLabeledValue) CNLabelContactRelationManager() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationManager"))
	return rv
}


// The label for the contact’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmother
func (c_ CNLabeledValue) CNLabelContactRelationMother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMother"))
	return rv
}


// The label for the contact’s mother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLaw"))
	return rv
}


// The label for the contact’s husband’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlawhusbandsmother
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawHusbandsMother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawHusbandsMother"))
	return rv
}


// The label for the contact’s mother-in-law or stepmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlaworstepmother
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawOrStepmother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawOrStepmother"))
	return rv
}


// The label for the contact’s wife’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlawwifesmother
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawWifesMother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawWifesMother"))
	return rv
}


// The label for the contact’s nephew.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephew
func (c_ CNLabeledValue) CNLabelContactRelationNephew() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephew"))
	return rv
}


// The label for the contact’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationNephewBrothersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephewBrothersSon"))
	return rv
}


// The label for the contact’s brother’s son or husband’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewbrotherssonorhusbandssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon"))
	return rv
}


// The label for the contact’s nephew or cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnepheworcousin
func (c_ CNLabeledValue) CNLabelContactRelationNephewOrCousin() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephewOrCousin"))
	return rv
}


// The label for the contact’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewsistersson
func (c_ CNLabeledValue) CNLabelContactRelationNephewSistersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephewSistersSon"))
	return rv
}


// The label for the contact’s sister’s son or wife’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewsisterssonorwifessiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon"))
	return rv
}


// The label for the contact’s niece.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniece
func (c_ CNLabeledValue) CNLabelContactRelationNiece() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNiece"))
	return rv
}


// The label for the contact’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecebrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceBrothersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNieceBrothersDaughter"))
	return rv
}


// The label for the contact’s brother’s daughter or husband’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecebrothersdaughterorhusbandssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter"))
	return rv
}


// The label for the contact’s niece or cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnieceorcousin
func (c_ CNLabeledValue) CNLabelContactRelationNieceOrCousin() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNieceOrCousin"))
	return rv
}


// The label for the contact’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecesistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceSistersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNieceSistersDaughter"))
	return rv
}


// The label for the contact’s sister’s daughter or wife’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecesistersdaughterorwifessiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter"))
	return rv
}


// The label for the contact’s parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparent
func (c_ CNLabeledValue) CNLabelContactRelationParent() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParent"))
	return rv
}


// The label for the contact’s parent-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentinlaw
func (c_ CNLabeledValue) CNLabelContactRelationParentInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentInLaw"))
	return rv
}


// The label for the contact’s parent’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentseldersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsElderSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsElderSibling"))
	return rv
}


// The label for the contact’s parent’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSibling"))
	return rv
}


// The label for the contact’s father’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfatherseldersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersElderSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersElderSibling"))
	return rv
}


// The label for the contact’s father’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfatherssibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersSibling"))
	return rv
}


// The label for the contact’s father’s youngest sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfathersyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersYoungerSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersYoungerSibling"))
	return rv
}


// The label for the contact’s mother’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmotherseldersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersElderSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersElderSibling"))
	return rv
}


// The label for the contact’s mother’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmotherssibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersSibling"))
	return rv
}


// The label for the contact’s mother’s younger sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmothersyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersYoungerSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersYoungerSibling"))
	return rv
}


// The label for the contact’s parent’s younger sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentsyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsYoungerSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsYoungerSibling"))
	return rv
}


// The label for the contact’s partner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationpartner
func (c_ CNLabeledValue) CNLabelContactRelationPartner() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationPartner"))
	return rv
}


// The label for the contact’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsibling
func (c_ CNLabeledValue) CNLabelContactRelationSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSibling"))
	return rv
}


// The label for the contact’s sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationSiblingInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSiblingInLaw"))
	return rv
}


// The label for the contact’s sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationSiblingsChild() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSiblingsChild"))
	return rv
}


// The label for the contact’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsister
func (c_ CNLabeledValue) CNLabelContactRelationSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSister"))
	return rv
}


// The label for the contact’s sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLaw"))
	return rv
}


// The label for the contact’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawBrothersWife() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawBrothersWife"))
	return rv
}


// The label for the contact’s elder brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawelderbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawElderBrothersWife() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawElderBrothersWife"))
	return rv
}


// The label for the contact’s husband’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawhusbandsbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawHusbandsBrothersWife() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawHusbandsBrothersWife"))
	return rv
}


// The label for the contact’s husband’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawhusbandssister
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawHusbandsSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawHusbandsSister"))
	return rv
}


// The label for the contact’s spouse’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawspousessister
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawSpousesSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawSpousesSister"))
	return rv
}


// The label for the contact’s wife’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawwifesbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawWifesBrothersWife() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawWifesBrothersWife"))
	return rv
}


// The label for the contact’s wife’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawwifessister
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawWifesSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawWifesSister"))
	return rv
}


// The label for the contact’s younger brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawyoungerbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawYoungerBrothersWife() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawYoungerBrothersWife"))
	return rv
}


// The label for the contact’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationson
func (c_ CNLabeledValue) CNLabelContactRelationSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSon"))
	return rv
}


// The label for the contact’s son-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaw
func (c_ CNLabeledValue) CNLabelContactRelationSonInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSonInLaw"))
	return rv
}


// The label for the contact’s son-in-law or brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaworbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationSonInLawOrBrotherInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSonInLawOrBrotherInLaw"))
	return rv
}


// The label for the contact’s son-in-law or stepson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaworstepson
func (c_ CNLabeledValue) CNLabelContactRelationSonInLawOrStepson() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSonInLawOrStepson"))
	return rv
}


// The label for the contact’s spouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationspouse
func (c_ CNLabeledValue) CNLabelContactRelationSpouse() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSpouse"))
	return rv
}


// The label for the contact’s stepbrother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepbrother
func (c_ CNLabeledValue) CNLabelContactRelationStepbrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepbrother"))
	return rv
}


// The label for the contact’s stepchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepchild
func (c_ CNLabeledValue) CNLabelContactRelationStepchild() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepchild"))
	return rv
}


// The label for the contact’s stepdaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepdaughter
func (c_ CNLabeledValue) CNLabelContactRelationStepdaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepdaughter"))
	return rv
}


// The label for the contact’s stepfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepfather
func (c_ CNLabeledValue) CNLabelContactRelationStepfather() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepfather"))
	return rv
}


// The label for the contact’s stepmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepmother
func (c_ CNLabeledValue) CNLabelContactRelationStepmother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepmother"))
	return rv
}


// The label for the contact’s stepparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepparent
func (c_ CNLabeledValue) CNLabelContactRelationStepparent() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepparent"))
	return rv
}


// The label for the contact’s stepsister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepsister
func (c_ CNLabeledValue) CNLabelContactRelationStepsister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepsister"))
	return rv
}


// The label for the contact’s stepson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepson
func (c_ CNLabeledValue) CNLabelContactRelationStepson() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepson"))
	return rv
}


// The label for the contact’s teacher.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationteacher
func (c_ CNLabeledValue) CNLabelContactRelationTeacher() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationTeacher"))
	return rv
}


// The label for the contact’s uncle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncle
func (c_ CNLabeledValue) CNLabelContactRelationUncle() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncle"))
	return rv
}


// The label for the contact’s father’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersBrother"))
	return rv
}


// The label for the contact’s father’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherselderbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersElderBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersElderBrother"))
	return rv
}


// The label for the contact’s elder sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherseldersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersElderSistersHusband() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersElderSistersHusband"))
	return rv
}


// The label for the contact’s father’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherssistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersSistersHusband() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersSistersHusband"))
	return rv
}


// The label for the contact’s father’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersYoungerBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersYoungerBrother"))
	return rv
}


// The label for the contact’s father’s younger sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersyoungersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersYoungerSistersHusband() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersYoungerSistersHusband"))
	return rv
}


// The label for the contact’s mother’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemothersbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersBrother"))
	return rv
}


// The label for the contact’s mother’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemotherselderbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersElderBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersElderBrother"))
	return rv
}


// The label for the contact’s mother’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemotherssistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersSistersHusband() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersSistersHusband"))
	return rv
}


// The label for the contact’s mother’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemothersyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersYoungerBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersYoungerBrother"))
	return rv
}


// The label for the contact’s parent’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentsbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsBrother"))
	return rv
}


// The label for the contact’s parent’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentselderbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsElderBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsElderBrother"))
	return rv
}


// The label for the contact’s parent’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentsyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsYoungerBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsYoungerBrother"))
	return rv
}


// The label for the contact’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationwife
func (c_ CNLabeledValue) CNLabelContactRelationWife() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationWife"))
	return rv
}


// The label for the contact’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationYoungerBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerBrother"))
	return rv
}


// The label for the contact’s younger brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungerbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationYoungerBrotherInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerBrotherInLaw"))
	return rv
}


// The label for the contact’s younger cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousin
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousin() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousin"))
	return rv
}


// The label for the contact’s father’s brother’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfathersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersBrothersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersBrothersDaughter"))
	return rv
}


// The label for the contact’s father’s brother’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfathersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersBrothersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersBrothersSon"))
	return rv
}


// The label for the contact’s father’s sister’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersSistersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersSistersDaughter"))
	return rv
}


// The label for the contact’s father’s sister’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersSistersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersSistersSon"))
	return rv
}


// The label for the contact’s mother’s brother’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmothersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersBrothersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersBrothersDaughter"))
	return rv
}


// The label for the contact’s mother’s brother’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmothersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersBrothersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersBrothersSon"))
	return rv
}


// The label for the contact’s mother’s sibling’s younger daughter or father’s sister’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssiblingsdaughterorfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter"))
	return rv
}


// The label for the contact’s mother’s sibling’s younger son or father’s sister’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssiblingssonorfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon"))
	return rv
}


// The label for the contact’s mother’s sister’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSistersDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSistersDaughter"))
	return rv
}


// The label for the contact’s mother’s sister’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSistersSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSistersSon"))
	return rv
}


// The label for the contact’s parent’s sibling’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinParentsSiblingsDaughter() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinParentsSiblingsDaughter"))
	return rv
}


// The label for the contact’s parent’s sibling’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinParentsSiblingsSon() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinParentsSiblingsSon"))
	return rv
}


// The label for the contact’s younger sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSibling() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSibling"))
	return rv
}


// The label for the contact’s younger sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSiblingInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSiblingInLaw"))
	return rv
}


// The label for the contact’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSister"))
	return rv
}


// The label for the contact’s younger sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSisterInLaw() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSisterInLaw"))
	return rv
}


// The label for the contact’s youngest brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungestbrother
func (c_ CNLabeledValue) CNLabelContactRelationYoungestBrother() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungestBrother"))
	return rv
}


// The label for the contact’s youngest sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungestsister
func (c_ CNLabeledValue) CNLabelContactRelationYoungestSister() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungestSister"))
	return rv
}


// The label for identifying the contact’s anniversary date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeldateanniversary
func (c_ CNLabeledValue) CNLabelDateAnniversary() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelDateAnniversary"))
	return rv
}


// The label for identifying the contact’s iCloud email information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelemailicloud
func (c_ CNLabeledValue) CNLabelEmailiCloud() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelEmailiCloud"))
	return rv
}


// The label for identifying home information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelhome
func (c_ CNLabeledValue) CNLabelHome() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelHome"))
	return rv
}


// The label for identifying other information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelother
func (c_ CNLabeledValue) CNLabelOther() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelOther"))
	return rv
}


// The label for identifying the contact’s Apple Watch phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberapplewatch
func (c_ CNLabeledValue) CNLabelPhoneNumberAppleWatch() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberAppleWatch"))
	return rv
}


// The label for identifying the contact’s home fax number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberhomefax
func (c_ CNLabeledValue) CNLabelPhoneNumberHomeFax() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberHomeFax"))
	return rv
}


// The label for identifying the contact’s main phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumbermain
func (c_ CNLabeledValue) CNLabelPhoneNumberMain() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberMain"))
	return rv
}


// The label for identifying the contact’s mobile phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumbermobile
func (c_ CNLabeledValue) CNLabelPhoneNumberMobile() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberMobile"))
	return rv
}


// The label for identifying another fax number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberotherfax
func (c_ CNLabeledValue) CNLabelPhoneNumberOtherFax() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberOtherFax"))
	return rv
}


// The label for identifying the contact’s pager number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberpager
func (c_ CNLabeledValue) CNLabelPhoneNumberPager() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberPager"))
	return rv
}


// The label for identifying the contact’s work fax number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberworkfax
func (c_ CNLabeledValue) CNLabelPhoneNumberWorkFax() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberWorkFax"))
	return rv
}


// The label for identifying the contact’s iPhone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberiphone
func (c_ CNLabeledValue) CNLabelPhoneNumberiPhone() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberiPhone"))
	return rv
}


// The label for the contact’s school.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelschool
func (c_ CNLabeledValue) CNLabelSchool() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelSchool"))
	return rv
}


// The label for identifying URL information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelurladdresshomepage
func (c_ CNLabeledValue) CNLabelURLAddressHomePage() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelURLAddressHomePage"))
	return rv
}


// The label for identifying work information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelwork
func (c_ CNLabeledValue) CNLabelWork() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelWork"))
	return rv
}


// A unique identifier for the labeled value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/identifier
func (c_ CNLabeledValue) Identifier() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("identifier"))
	return rv
}


// A unique identifier for the labeled value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/identifier
func (c_ CNLabeledValue) SetIdentifier(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// The label for a contact property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/label
func (c_ CNLabeledValue) Label() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](c_.ID, objc.Sel("label"))
	return rv
}


// The label for a contact property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/label
func (c_ CNLabeledValue) SetLabel(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), objc.String(value))
}


// A contact property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/value
func (c_ CNLabeledValue) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("value"))
	return rv
}


// A contact property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeledvalue/value
func (c_ CNLabeledValue) SetValue(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setValue:"), value)
}



