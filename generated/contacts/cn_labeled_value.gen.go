// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	LabeledValueBySettingLabel(label string) unsafe.Pointer
	LabeledValueBySettingLabelValue(label string, value unsafe.Pointer) unsafe.Pointer
	LabeledValueBySettingValue(value unsafe.Pointer) unsafe.Pointer
	Identifier() string
	Label() string
	Value() unsafe.Pointer
	CNLabelContactRelationAssistant() string
	CNLabelContactRelationAunt() string
	CNLabelContactRelationAuntFathersBrothersWife() string
	CNLabelContactRelationAuntFathersElderBrothersWife() string
	CNLabelContactRelationAuntFathersElderSister() string
	CNLabelContactRelationAuntFathersSister() string
	CNLabelContactRelationAuntFathersYoungerBrothersWife() string
	CNLabelContactRelationAuntFathersYoungerSister() string
	CNLabelContactRelationAuntMothersBrothersWife() string
	CNLabelContactRelationAuntMothersElderSister() string
	CNLabelContactRelationAuntMothersSister() string
	CNLabelContactRelationAuntMothersYoungerSister() string
	CNLabelContactRelationAuntParentsElderSister() string
	CNLabelContactRelationAuntParentsSister() string
	CNLabelContactRelationAuntParentsYoungerSister() string
	CNLabelContactRelationBoyfriend() string
	CNLabelContactRelationBrother() string
	CNLabelContactRelationBrotherInLaw() string
	CNLabelContactRelationBrotherInLawElderSistersHusband() string
	CNLabelContactRelationBrotherInLawHusbandsBrother() string
	CNLabelContactRelationBrotherInLawHusbandsSistersHusband() string
	CNLabelContactRelationBrotherInLawSistersHusband() string
	CNLabelContactRelationBrotherInLawSpousesBrother() string
	CNLabelContactRelationBrotherInLawWifesBrother() string
	CNLabelContactRelationBrotherInLawWifesSistersHusband() string
	CNLabelContactRelationBrotherInLawYoungerSistersHusband() string
	CNLabelContactRelationChild() string
	CNLabelContactRelationChildInLaw() string
	CNLabelContactRelationCoBrotherInLaw() string
	CNLabelContactRelationCoFatherInLaw() string
	CNLabelContactRelationCoMotherInLaw() string
	CNLabelContactRelationCoParentInLaw() string
	CNLabelContactRelationCoSiblingInLaw() string
	CNLabelContactRelationCoSisterInLaw() string
	CNLabelContactRelationColleague() string
	CNLabelContactRelationCousin() string
	CNLabelContactRelationCousinFathersBrothersDaughter() string
	CNLabelContactRelationCousinFathersBrothersSon() string
	CNLabelContactRelationCousinFathersSistersDaughter() string
	CNLabelContactRelationCousinFathersSistersSon() string
	CNLabelContactRelationCousinGrandparentsSiblingsChild() string
	CNLabelContactRelationCousinGrandparentsSiblingsDaughter() string
	CNLabelContactRelationCousinGrandparentsSiblingsSon() string
	CNLabelContactRelationCousinMothersBrothersDaughter() string
	CNLabelContactRelationCousinMothersBrothersSon() string
	CNLabelContactRelationCousinMothersSistersDaughter() string
	CNLabelContactRelationCousinMothersSistersSon() string
	CNLabelContactRelationCousinOrSiblingsChild() string
	CNLabelContactRelationCousinParentsSiblingsChild() string
	CNLabelContactRelationCousinParentsSiblingsDaughter() string
	CNLabelContactRelationCousinParentsSiblingsSon() string
	CNLabelContactRelationDaughter() string
	CNLabelContactRelationDaughterInLaw() string
	CNLabelContactRelationDaughterInLawOrSisterInLaw() string
	CNLabelContactRelationDaughterInLawOrStepdaughter() string
	CNLabelContactRelationElderBrother() string
	CNLabelContactRelationElderBrotherInLaw() string
	CNLabelContactRelationElderCousin() string
	CNLabelContactRelationElderCousinFathersBrothersDaughter() string
	CNLabelContactRelationElderCousinFathersBrothersSon() string
	CNLabelContactRelationElderCousinFathersSistersDaughter() string
	CNLabelContactRelationElderCousinFathersSistersSon() string
	CNLabelContactRelationElderCousinMothersBrothersDaughter() string
	CNLabelContactRelationElderCousinMothersBrothersSon() string
	CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter() string
	CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon() string
	CNLabelContactRelationElderCousinMothersSistersDaughter() string
	CNLabelContactRelationElderCousinMothersSistersSon() string
	CNLabelContactRelationElderCousinParentsSiblingsDaughter() string
	CNLabelContactRelationElderCousinParentsSiblingsSon() string
	CNLabelContactRelationElderSibling() string
	CNLabelContactRelationElderSiblingInLaw() string
	CNLabelContactRelationElderSister() string
	CNLabelContactRelationElderSisterInLaw() string
	CNLabelContactRelationEldestBrother() string
	CNLabelContactRelationEldestSister() string
	CNLabelContactRelationFather() string
	CNLabelContactRelationFatherInLaw() string
	CNLabelContactRelationFatherInLawHusbandsFather() string
	CNLabelContactRelationFatherInLawOrStepfather() string
	CNLabelContactRelationFatherInLawWifesFather() string
	CNLabelContactRelationFemaleCousin() string
	CNLabelContactRelationFemaleFriend() string
	CNLabelContactRelationFemalePartner() string
	CNLabelContactRelationFriend() string
	CNLabelContactRelationGirlfriend() string
	CNLabelContactRelationGirlfriendOrBoyfriend() string
	CNLabelContactRelationGrandaunt() string
	CNLabelContactRelationGrandchild() string
	CNLabelContactRelationGrandchildOrSiblingsChild() string
	CNLabelContactRelationGranddaughter() string
	CNLabelContactRelationGranddaughterDaughtersDaughter() string
	CNLabelContactRelationGranddaughterOrNiece() string
	CNLabelContactRelationGranddaughterSonsDaughter() string
	CNLabelContactRelationGrandfather() string
	CNLabelContactRelationGrandfatherFathersFather() string
	CNLabelContactRelationGrandfatherMothersFather() string
	CNLabelContactRelationGrandmother() string
	CNLabelContactRelationGrandmotherFathersMother() string
	CNLabelContactRelationGrandmotherMothersMother() string
	CNLabelContactRelationGrandnephew() string
	CNLabelContactRelationGrandnephewBrothersGrandson() string
	CNLabelContactRelationGrandnephewSistersGrandson() string
	CNLabelContactRelationGrandniece() string
	CNLabelContactRelationGrandnieceBrothersGranddaughter() string
	CNLabelContactRelationGrandnieceSistersGranddaughter() string
	CNLabelContactRelationGrandparent() string
	CNLabelContactRelationGrandson() string
	CNLabelContactRelationGrandsonDaughtersSon() string
	CNLabelContactRelationGrandsonOrNephew() string
	CNLabelContactRelationGrandsonSonsSon() string
	CNLabelContactRelationGranduncle() string
	CNLabelContactRelationGreatGrandchild() string
	CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild() string
	CNLabelContactRelationGreatGranddaughter() string
	CNLabelContactRelationGreatGrandfather() string
	CNLabelContactRelationGreatGrandmother() string
	CNLabelContactRelationGreatGrandparent() string
	CNLabelContactRelationGreatGrandson() string
	CNLabelContactRelationHusband() string
	CNLabelContactRelationMaleCousin() string
	CNLabelContactRelationMaleFriend() string
	CNLabelContactRelationMalePartner() string
	CNLabelContactRelationManager() string
	CNLabelContactRelationMother() string
	CNLabelContactRelationMotherInLaw() string
	CNLabelContactRelationMotherInLawHusbandsMother() string
	CNLabelContactRelationMotherInLawOrStepmother() string
	CNLabelContactRelationMotherInLawWifesMother() string
	CNLabelContactRelationNephew() string
	CNLabelContactRelationNephewBrothersSon() string
	CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon() string
	CNLabelContactRelationNephewOrCousin() string
	CNLabelContactRelationNephewSistersSon() string
	CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon() string
	CNLabelContactRelationNiece() string
	CNLabelContactRelationNieceBrothersDaughter() string
	CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter() string
	CNLabelContactRelationNieceOrCousin() string
	CNLabelContactRelationNieceSistersDaughter() string
	CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter() string
	CNLabelContactRelationParent() string
	CNLabelContactRelationParentInLaw() string
	CNLabelContactRelationParentsElderSibling() string
	CNLabelContactRelationParentsSibling() string
	CNLabelContactRelationParentsSiblingFathersElderSibling() string
	CNLabelContactRelationParentsSiblingFathersSibling() string
	CNLabelContactRelationParentsSiblingFathersYoungerSibling() string
	CNLabelContactRelationParentsSiblingMothersElderSibling() string
	CNLabelContactRelationParentsSiblingMothersSibling() string
	CNLabelContactRelationParentsSiblingMothersYoungerSibling() string
	CNLabelContactRelationParentsYoungerSibling() string
	CNLabelContactRelationPartner() string
	CNLabelContactRelationSibling() string
	CNLabelContactRelationSiblingInLaw() string
	CNLabelContactRelationSiblingsChild() string
	CNLabelContactRelationSister() string
	CNLabelContactRelationSisterInLaw() string
	CNLabelContactRelationSisterInLawBrothersWife() string
	CNLabelContactRelationSisterInLawElderBrothersWife() string
	CNLabelContactRelationSisterInLawHusbandsBrothersWife() string
	CNLabelContactRelationSisterInLawHusbandsSister() string
	CNLabelContactRelationSisterInLawSpousesSister() string
	CNLabelContactRelationSisterInLawWifesBrothersWife() string
	CNLabelContactRelationSisterInLawWifesSister() string
	CNLabelContactRelationSisterInLawYoungerBrothersWife() string
	CNLabelContactRelationSon() string
	CNLabelContactRelationSonInLaw() string
	CNLabelContactRelationSonInLawOrBrotherInLaw() string
	CNLabelContactRelationSonInLawOrStepson() string
	CNLabelContactRelationSpouse() string
	CNLabelContactRelationStepbrother() string
	CNLabelContactRelationStepchild() string
	CNLabelContactRelationStepdaughter() string
	CNLabelContactRelationStepfather() string
	CNLabelContactRelationStepmother() string
	CNLabelContactRelationStepparent() string
	CNLabelContactRelationStepsister() string
	CNLabelContactRelationStepson() string
	CNLabelContactRelationTeacher() string
	CNLabelContactRelationUncle() string
	CNLabelContactRelationUncleFathersBrother() string
	CNLabelContactRelationUncleFathersElderBrother() string
	CNLabelContactRelationUncleFathersElderSistersHusband() string
	CNLabelContactRelationUncleFathersSistersHusband() string
	CNLabelContactRelationUncleFathersYoungerBrother() string
	CNLabelContactRelationUncleFathersYoungerSistersHusband() string
	CNLabelContactRelationUncleMothersBrother() string
	CNLabelContactRelationUncleMothersElderBrother() string
	CNLabelContactRelationUncleMothersSistersHusband() string
	CNLabelContactRelationUncleMothersYoungerBrother() string
	CNLabelContactRelationUncleParentsBrother() string
	CNLabelContactRelationUncleParentsElderBrother() string
	CNLabelContactRelationUncleParentsYoungerBrother() string
	CNLabelContactRelationWife() string
	CNLabelContactRelationYoungerBrother() string
	CNLabelContactRelationYoungerBrotherInLaw() string
	CNLabelContactRelationYoungerCousin() string
	CNLabelContactRelationYoungerCousinFathersBrothersDaughter() string
	CNLabelContactRelationYoungerCousinFathersBrothersSon() string
	CNLabelContactRelationYoungerCousinFathersSistersDaughter() string
	CNLabelContactRelationYoungerCousinFathersSistersSon() string
	CNLabelContactRelationYoungerCousinMothersBrothersDaughter() string
	CNLabelContactRelationYoungerCousinMothersBrothersSon() string
	CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter() string
	CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon() string
	CNLabelContactRelationYoungerCousinMothersSistersDaughter() string
	CNLabelContactRelationYoungerCousinMothersSistersSon() string
	CNLabelContactRelationYoungerCousinParentsSiblingsDaughter() string
	CNLabelContactRelationYoungerCousinParentsSiblingsSon() string
	CNLabelContactRelationYoungerSibling() string
	CNLabelContactRelationYoungerSiblingInLaw() string
	CNLabelContactRelationYoungerSister() string
	CNLabelContactRelationYoungerSisterInLaw() string
	CNLabelContactRelationYoungestBrother() string
	CNLabelContactRelationYoungestSister() string
	CNLabelDateAnniversary() string
	CNLabelEmailiCloud() string
	CNLabelHome() string
	CNLabelOther() string
	CNLabelPhoneNumberAppleWatch() string
	CNLabelPhoneNumberHomeFax() string
	CNLabelPhoneNumberMain() string
	CNLabelPhoneNumberMobile() string
	CNLabelPhoneNumberOtherFax() string
	CNLabelPhoneNumberPager() string
	CNLabelPhoneNumberWorkFax() string
	CNLabelPhoneNumberiPhone() string
	CNLabelSchool() string
	CNLabelURLAddressHomePage() string
	CNLabelWork() string
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




// Returns a new labeled value identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/init(label:value:)

func NewCNLabeledValueWithLabelValue(label string, value unsafe.Pointer) CNLabeledValue {
	instance := getCNLabeledValueClass().Alloc()
	rv := objc.Send[CNLabeledValue](instance.ID, objc.Sel("initWithLabel:value:"), objc.String(label), value)
	rv.Autorelease()
	return rv
}



// Returns a new labeled value identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/labeledValueWithLabel:value:

func (cc _CNLabeledValueClass) LabeledValueWithLabelValue(label string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("labeledValueWithLabel:value:"), objc.String(label), value)
	return rv
}


// Returns a localized string for the specified label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/localizedString(forLabel:)

func (cc _CNLabeledValueClass) LocalizedStringForLabel(label string) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForLabel:"), objc.String(label))
	return rv
}



// Returns a labeled value object with an existing value and identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingLabel(_:)

func (c_ CNLabeledValue) LabeledValueBySettingLabel(label string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("labeledValueBySettingLabel:"), objc.String(label))
	return rv
}



// Returns a labeled value object with the specified label and value with the existing identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingLabel(_:value:)

func (c_ CNLabeledValue) LabeledValueBySettingLabelValue(label string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("labeledValueBySettingLabel:value:"), objc.String(label), value)
	return rv
}



// Returns a new value for an existing label and identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingValue(_:)

func (c_ CNLabeledValue) LabeledValueBySettingValue(value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("labeledValueBySettingValue:"), value)
	return rv
}


// A unique identifier for the labeled value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/identifier

func (c_ CNLabeledValue) Identifier() string {
	rv := objc.Send[string](c_.ID, objc.Sel("identifier"))
	return rv
}


// The label for a contact property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/label

func (c_ CNLabeledValue) Label() string {
	rv := objc.Send[string](c_.ID, objc.Sel("label"))
	return rv
}


// A contact property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/value

func (c_ CNLabeledValue) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("value"))
	return rv
}


// The label for the contact’s assistant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationassistant

func (c_ CNLabeledValue) CNLabelContactRelationAssistant() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAssistant"))
	return rv
}


// The label for the contact’s aunt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationaunt

func (c_ CNLabeledValue) CNLabelContactRelationAunt() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAunt"))
	return rv
}


// The label for the contact’s father’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersbrotherswife

func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersBrothersWife() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersBrothersWife"))
	return rv
}


// The label for the contact’s father’s elder brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherselderbrotherswife

func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersElderBrothersWife() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersElderBrothersWife"))
	return rv
}


// The label for the contact’s father’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherseldersister

func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersElderSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersElderSister"))
	return rv
}


// The label for the contact’s father’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherssister

func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersSister"))
	return rv
}


// The label for the contact’s father’s younger brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersyoungerbrotherswife

func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersYoungerBrothersWife() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersYoungerBrothersWife"))
	return rv
}


// The label for the contact’s father’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersyoungersister

func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersYoungerSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersYoungerSister"))
	return rv
}


// The label for the contact’s mother’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmothersbrotherswife

func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersBrothersWife() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersBrothersWife"))
	return rv
}


// The label for the contact’s mother’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmotherseldersister

func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersElderSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersElderSister"))
	return rv
}


// The label for the contact’s mother’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmotherssister

func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersSister"))
	return rv
}


// The label for the contact’s mother’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmothersyoungersister

func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersYoungerSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersYoungerSister"))
	return rv
}


// The label for the contact’s parent’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentseldersister

func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsElderSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsElderSister"))
	return rv
}


// The label for the contact’s parent’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentssister

func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsSister"))
	return rv
}


// The label for the contact’s parent’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentsyoungersister

func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsYoungerSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsYoungerSister"))
	return rv
}


// The label for the contact’s boyfriend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationboyfriend

func (c_ CNLabeledValue) CNLabelContactRelationBoyfriend() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBoyfriend"))
	return rv
}


// The label for the contact’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrother

func (c_ CNLabeledValue) CNLabelContactRelationBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrother"))
	return rv
}


// The label for the contact’s brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlaw

func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLaw"))
	return rv
}


// The label for the contact’s elder sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlaweldersistershusband

func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawElderSistersHusband() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawElderSistersHusband"))
	return rv
}


// The label for the contact’s husband’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawhusbandsbrother

func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawHusbandsBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawHusbandsBrother"))
	return rv
}


// The label for the contact’s husband’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawhusbandssistershusband

func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawHusbandsSistersHusband() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawHusbandsSistersHusband"))
	return rv
}


// The label for the contact’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawsistershusband

func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawSistersHusband() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawSistersHusband"))
	return rv
}


// The label for the contact’s spouse’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawspousesbrother

func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawSpousesBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawSpousesBrother"))
	return rv
}


// The label for the contact’s wife’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawwifesbrother

func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawWifesBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawWifesBrother"))
	return rv
}


// The label for the contact’s wife’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawwifessistershusband

func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawWifesSistersHusband() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawWifesSistersHusband"))
	return rv
}


// The label for the contact’s younger sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawyoungersistershusband

func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawYoungerSistersHusband() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawYoungerSistersHusband"))
	return rv
}


// The label for the contact’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationchild

func (c_ CNLabeledValue) CNLabelContactRelationChild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationChild"))
	return rv
}


// The label for the contact’s child-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationchildinlaw

func (c_ CNLabeledValue) CNLabelContactRelationChildInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationChildInLaw"))
	return rv
}


// The label for the contact’s co-brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcobrotherinlaw

func (c_ CNLabeledValue) CNLabelContactRelationCoBrotherInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoBrotherInLaw"))
	return rv
}


// The label for the contact’s co-father-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcofatherinlaw

func (c_ CNLabeledValue) CNLabelContactRelationCoFatherInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoFatherInLaw"))
	return rv
}


// The label for the contact’s co-mother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcomotherinlaw

func (c_ CNLabeledValue) CNLabelContactRelationCoMotherInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoMotherInLaw"))
	return rv
}


// The label for the contact’s co-parent-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcoparentinlaw

func (c_ CNLabeledValue) CNLabelContactRelationCoParentInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoParentInLaw"))
	return rv
}


// The label for the contact’s co-sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcosiblinginlaw

func (c_ CNLabeledValue) CNLabelContactRelationCoSiblingInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoSiblingInLaw"))
	return rv
}


// The label for the contact’s co-sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcosisterinlaw

func (c_ CNLabeledValue) CNLabelContactRelationCoSisterInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCoSisterInLaw"))
	return rv
}


// The label for the contact’s colleague.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcolleague

func (c_ CNLabeledValue) CNLabelContactRelationColleague() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationColleague"))
	return rv
}


// The label for the contact’s cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousin

func (c_ CNLabeledValue) CNLabelContactRelationCousin() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousin"))
	return rv
}


// The label for the contact’s father’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfathersbrothersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersBrothersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersBrothersDaughter"))
	return rv
}


// The label for the contact’s father’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfathersbrothersson

func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersBrothersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersBrothersSon"))
	return rv
}


// The label for the contact’s father’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfatherssistersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersSistersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersSistersDaughter"))
	return rv
}


// The label for the contact’s father’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfatherssistersson

func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersSistersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersSistersSon"))
	return rv
}


// The label for the contact’s grandparent’s sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingschild

func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsChild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsChild"))
	return rv
}


// The label for the contact’s grandparent’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingsdaughter

func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsDaughter"))
	return rv
}


// The label for the contact’s grandparent’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingsson

func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsSon"))
	return rv
}


// The label for the contact’s mother’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmothersbrothersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersBrothersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersBrothersDaughter"))
	return rv
}


// The label for the contact’s mother’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmothersbrothersson

func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersBrothersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersBrothersSon"))
	return rv
}


// The label for the contact’s mother’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmotherssistersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersSistersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersSistersDaughter"))
	return rv
}


// The label for the contact’s mother’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmotherssistersson

func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersSistersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersSistersSon"))
	return rv
}


// The label for the contact’s cousin’s or sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinorsiblingschild

func (c_ CNLabeledValue) CNLabelContactRelationCousinOrSiblingsChild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinOrSiblingsChild"))
	return rv
}


// The label for the contact’s parent’s sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingschild

func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsChild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsChild"))
	return rv
}


// The label for the contact’s parent’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingsdaughter

func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsDaughter"))
	return rv
}


// The label for the contact’s parent’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingsson

func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsSon"))
	return rv
}


// The label for the contact’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughter

func (c_ CNLabeledValue) CNLabelContactRelationDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationDaughter"))
	return rv
}


// The label for the contact’s daughter-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaw

func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLaw"))
	return rv
}


// The label for the contact’s daughter-in-law or sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaworsisterinlaw

func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLawOrSisterInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLawOrSisterInLaw"))
	return rv
}


// The label for the contact’s daughter-in-law or stepdaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaworstepdaughter

func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLawOrStepdaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLawOrStepdaughter"))
	return rv
}


// The label for the contact’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationelderbrother

func (c_ CNLabeledValue) CNLabelContactRelationElderBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderBrother"))
	return rv
}


// The label for the contact’s elder brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationelderbrotherinlaw

func (c_ CNLabeledValue) CNLabelContactRelationElderBrotherInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderBrotherInLaw"))
	return rv
}


// The label for the contact’s elder cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousin

func (c_ CNLabeledValue) CNLabelContactRelationElderCousin() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousin"))
	return rv
}


// The label for the contact’s father’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfathersbrothersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersBrothersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersBrothersDaughter"))
	return rv
}


// The label for the contact’s father’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfathersbrothersson

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersBrothersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersBrothersSon"))
	return rv
}


// The label for the contact’s father’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfatherssistersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersSistersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersSistersDaughter"))
	return rv
}


// The label for the contact’s father’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfatherssistersson

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersSistersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersSistersSon"))
	return rv
}


// The label for the contact’s mother’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmothersbrothersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersBrothersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersBrothersDaughter"))
	return rv
}


// The label for the contact’s mother’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmothersbrothersson

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersBrothersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersBrothersSon"))
	return rv
}


// The label for the contact’s mother’s sibling’s daughter or father’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssiblingsdaughterorfatherssistersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter"))
	return rv
}


// The label for the contact’s mother’s sibling’s son or father’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssiblingssonorfatherssistersson

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon"))
	return rv
}


// The label for the contact’s mother’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssistersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSistersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSistersDaughter"))
	return rv
}


// The label for the contact’s mother’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssistersson

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSistersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSistersSon"))
	return rv
}


// The label for the contact’s parent’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinparentssiblingsdaughter

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinParentsSiblingsDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinParentsSiblingsDaughter"))
	return rv
}


// The label for the contact’s parent’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinparentssiblingsson

func (c_ CNLabeledValue) CNLabelContactRelationElderCousinParentsSiblingsSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinParentsSiblingsSon"))
	return rv
}


// The label for the contact’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersibling

func (c_ CNLabeledValue) CNLabelContactRelationElderSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderSibling"))
	return rv
}


// The label for the contact’s elder sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersiblinginlaw

func (c_ CNLabeledValue) CNLabelContactRelationElderSiblingInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderSiblingInLaw"))
	return rv
}


// The label for the contact’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersister

func (c_ CNLabeledValue) CNLabelContactRelationElderSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderSister"))
	return rv
}


// The label for the contact’s elder sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersisterinlaw

func (c_ CNLabeledValue) CNLabelContactRelationElderSisterInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationElderSisterInLaw"))
	return rv
}


// The label for the contact’s eldest brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldestbrother

func (c_ CNLabeledValue) CNLabelContactRelationEldestBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationEldestBrother"))
	return rv
}


// The label for the contact’s eldest sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldestsister

func (c_ CNLabeledValue) CNLabelContactRelationEldestSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationEldestSister"))
	return rv
}


// The label for the contact’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfather

func (c_ CNLabeledValue) CNLabelContactRelationFather() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFather"))
	return rv
}


// The label for the contact’s father-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlaw

func (c_ CNLabeledValue) CNLabelContactRelationFatherInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLaw"))
	return rv
}


// The label for the contact’s husband’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlawhusbandsfather

func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawHusbandsFather() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawHusbandsFather"))
	return rv
}


// The label for the contact’s father-in-law or stepfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlaworstepfather

func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawOrStepfather() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawOrStepfather"))
	return rv
}


// The label for the contact’s wife’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlawwifesfather

func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawWifesFather() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawWifesFather"))
	return rv
}


// The label for the contact’s female cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalecousin

func (c_ CNLabeledValue) CNLabelContactRelationFemaleCousin() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFemaleCousin"))
	return rv
}


// The label for the contact’s female friend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalefriend

func (c_ CNLabeledValue) CNLabelContactRelationFemaleFriend() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFemaleFriend"))
	return rv
}


// The label for the contact’s female partner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalepartner

func (c_ CNLabeledValue) CNLabelContactRelationFemalePartner() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFemalePartner"))
	return rv
}


// The label for the contact’s friend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfriend

func (c_ CNLabeledValue) CNLabelContactRelationFriend() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationFriend"))
	return rv
}


// The label for the contact’s girlfriend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgirlfriend

func (c_ CNLabeledValue) CNLabelContactRelationGirlfriend() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGirlfriend"))
	return rv
}


// The label for the contact’s girlfriend or boyfriend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgirlfriendorboyfriend

func (c_ CNLabeledValue) CNLabelContactRelationGirlfriendOrBoyfriend() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGirlfriendOrBoyfriend"))
	return rv
}


// The label for the contact’s grandaunt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandaunt

func (c_ CNLabeledValue) CNLabelContactRelationGrandaunt() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandaunt"))
	return rv
}


// The label for the contact’s grandchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandchild

func (c_ CNLabeledValue) CNLabelContactRelationGrandchild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandchild"))
	return rv
}


// The label for the contact’s grandchild or sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandchildorsiblingschild

func (c_ CNLabeledValue) CNLabelContactRelationGrandchildOrSiblingsChild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandchildOrSiblingsChild"))
	return rv
}


// The label for the contact’s granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughter

func (c_ CNLabeledValue) CNLabelContactRelationGranddaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughter"))
	return rv
}


// The label for the contact’s daughter’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughterdaughtersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterDaughtersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterDaughtersDaughter"))
	return rv
}


// The label for the contact’s granddaughter or niece.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughterorniece

func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterOrNiece() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterOrNiece"))
	return rv
}


// The label for the contact’s son’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughtersonsdaughter

func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterSonsDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterSonsDaughter"))
	return rv
}


// The label for the contact’s grandfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfather

func (c_ CNLabeledValue) CNLabelContactRelationGrandfather() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandfather"))
	return rv
}


// The label for the contact’s father’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfatherfathersfather

func (c_ CNLabeledValue) CNLabelContactRelationGrandfatherFathersFather() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandfatherFathersFather"))
	return rv
}


// The label for the contact’s mother’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfathermothersfather

func (c_ CNLabeledValue) CNLabelContactRelationGrandfatherMothersFather() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandfatherMothersFather"))
	return rv
}


// The label for the contact’s grandmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmother

func (c_ CNLabeledValue) CNLabelContactRelationGrandmother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandmother"))
	return rv
}


// The label for the contact’s father’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmotherfathersmother

func (c_ CNLabeledValue) CNLabelContactRelationGrandmotherFathersMother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandmotherFathersMother"))
	return rv
}


// The label for the contact’s mother’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmothermothersmother

func (c_ CNLabeledValue) CNLabelContactRelationGrandmotherMothersMother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandmotherMothersMother"))
	return rv
}


// The label for the contact’s grandnephew.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephew

func (c_ CNLabeledValue) CNLabelContactRelationGrandnephew() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandnephew"))
	return rv
}


// The label for the contact’s brother’s grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephewbrothersgrandson

func (c_ CNLabeledValue) CNLabelContactRelationGrandnephewBrothersGrandson() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandnephewBrothersGrandson"))
	return rv
}


// The label for the contact’s sister’s grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephewsistersgrandson

func (c_ CNLabeledValue) CNLabelContactRelationGrandnephewSistersGrandson() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandnephewSistersGrandson"))
	return rv
}


// The label for the contact’s grandniece.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniece

func (c_ CNLabeledValue) CNLabelContactRelationGrandniece() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandniece"))
	return rv
}


// The label for the contact’s brother’s granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniecebrothersgranddaughter

func (c_ CNLabeledValue) CNLabelContactRelationGrandnieceBrothersGranddaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandnieceBrothersGranddaughter"))
	return rv
}


// The label for the contact’s sister’s granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniecesistersgranddaughter

func (c_ CNLabeledValue) CNLabelContactRelationGrandnieceSistersGranddaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandnieceSistersGranddaughter"))
	return rv
}


// The label for the contact’s grandparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandparent

func (c_ CNLabeledValue) CNLabelContactRelationGrandparent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandparent"))
	return rv
}


// The label for the contact’s grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandson

func (c_ CNLabeledValue) CNLabelContactRelationGrandson() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandson"))
	return rv
}


// The label for the contact’s daughter’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsondaughtersson

func (c_ CNLabeledValue) CNLabelContactRelationGrandsonDaughtersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandsonDaughtersSon"))
	return rv
}


// The label for the contact’s grandson or nephew.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsonornephew

func (c_ CNLabeledValue) CNLabelContactRelationGrandsonOrNephew() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandsonOrNephew"))
	return rv
}


// The label for the contact’s son’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsonsonsson

func (c_ CNLabeledValue) CNLabelContactRelationGrandsonSonsSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGrandsonSonsSon"))
	return rv
}


// The label for the contact’s granduncle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranduncle

func (c_ CNLabeledValue) CNLabelContactRelationGranduncle() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGranduncle"))
	return rv
}


// The label for the contact’s grandchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandchild

func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandchild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandchild"))
	return rv
}


// The label for the contact’s grandchild or sibling’s grandchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandchildorsiblingsgrandchild

func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild"))
	return rv
}


// The label for the contact’s great-granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgranddaughter

func (c_ CNLabeledValue) CNLabelContactRelationGreatGranddaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGranddaughter"))
	return rv
}


// The label for the contact’s great-grandfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandfather

func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandfather() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandfather"))
	return rv
}


// The label for the contact’s great-grandmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandmother

func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandmother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandmother"))
	return rv
}


// The label for the contact’s great-grandparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandparent

func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandparent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandparent"))
	return rv
}


// The label for the contact’s great-grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandson

func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandson() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandson"))
	return rv
}


// The label for the contact’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationhusband

func (c_ CNLabeledValue) CNLabelContactRelationHusband() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationHusband"))
	return rv
}


// The label for the contact’s male cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalecousin

func (c_ CNLabeledValue) CNLabelContactRelationMaleCousin() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMaleCousin"))
	return rv
}


// The label for the contact’s male friend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalefriend

func (c_ CNLabeledValue) CNLabelContactRelationMaleFriend() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMaleFriend"))
	return rv
}


// The label for the contact’s male partner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalepartner

func (c_ CNLabeledValue) CNLabelContactRelationMalePartner() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMalePartner"))
	return rv
}


// The label for the contact’s manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmanager

func (c_ CNLabeledValue) CNLabelContactRelationManager() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationManager"))
	return rv
}


// The label for the contact’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmother

func (c_ CNLabeledValue) CNLabelContactRelationMother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMother"))
	return rv
}


// The label for the contact’s mother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlaw

func (c_ CNLabeledValue) CNLabelContactRelationMotherInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLaw"))
	return rv
}


// The label for the contact’s husband’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlawhusbandsmother

func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawHusbandsMother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawHusbandsMother"))
	return rv
}


// The label for the contact’s mother-in-law or stepmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlaworstepmother

func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawOrStepmother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawOrStepmother"))
	return rv
}


// The label for the contact’s wife’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlawwifesmother

func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawWifesMother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawWifesMother"))
	return rv
}


// The label for the contact’s nephew.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephew

func (c_ CNLabeledValue) CNLabelContactRelationNephew() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephew"))
	return rv
}


// The label for the contact’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewbrothersson

func (c_ CNLabeledValue) CNLabelContactRelationNephewBrothersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephewBrothersSon"))
	return rv
}


// The label for the contact’s brother’s son or husband’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewbrotherssonorhusbandssiblingsson

func (c_ CNLabeledValue) CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon"))
	return rv
}


// The label for the contact’s nephew or cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnepheworcousin

func (c_ CNLabeledValue) CNLabelContactRelationNephewOrCousin() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephewOrCousin"))
	return rv
}


// The label for the contact’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewsistersson

func (c_ CNLabeledValue) CNLabelContactRelationNephewSistersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephewSistersSon"))
	return rv
}


// The label for the contact’s sister’s son or wife’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewsisterssonorwifessiblingsson

func (c_ CNLabeledValue) CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon"))
	return rv
}


// The label for the contact’s niece.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniece

func (c_ CNLabeledValue) CNLabelContactRelationNiece() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNiece"))
	return rv
}


// The label for the contact’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecebrothersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationNieceBrothersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNieceBrothersDaughter"))
	return rv
}


// The label for the contact’s brother’s daughter or husband’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecebrothersdaughterorhusbandssiblingsdaughter

func (c_ CNLabeledValue) CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter"))
	return rv
}


// The label for the contact’s niece or cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnieceorcousin

func (c_ CNLabeledValue) CNLabelContactRelationNieceOrCousin() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNieceOrCousin"))
	return rv
}


// The label for the contact’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecesistersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationNieceSistersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNieceSistersDaughter"))
	return rv
}


// The label for the contact’s sister’s daughter or wife’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecesistersdaughterorwifessiblingsdaughter

func (c_ CNLabeledValue) CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter"))
	return rv
}


// The label for the contact’s parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparent

func (c_ CNLabeledValue) CNLabelContactRelationParent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParent"))
	return rv
}


// The label for the contact’s parent-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentinlaw

func (c_ CNLabeledValue) CNLabelContactRelationParentInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentInLaw"))
	return rv
}


// The label for the contact’s parent’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentseldersibling

func (c_ CNLabeledValue) CNLabelContactRelationParentsElderSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsElderSibling"))
	return rv
}


// The label for the contact’s parent’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssibling

func (c_ CNLabeledValue) CNLabelContactRelationParentsSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSibling"))
	return rv
}


// The label for the contact’s father’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfatherseldersibling

func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersElderSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersElderSibling"))
	return rv
}


// The label for the contact’s father’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfatherssibling

func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersSibling"))
	return rv
}


// The label for the contact’s father’s youngest sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfathersyoungersibling

func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersYoungerSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersYoungerSibling"))
	return rv
}


// The label for the contact’s mother’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmotherseldersibling

func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersElderSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersElderSibling"))
	return rv
}


// The label for the contact’s mother’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmotherssibling

func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersSibling"))
	return rv
}


// The label for the contact’s mother’s younger sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmothersyoungersibling

func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersYoungerSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersYoungerSibling"))
	return rv
}


// The label for the contact’s parent’s younger sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentsyoungersibling

func (c_ CNLabeledValue) CNLabelContactRelationParentsYoungerSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationParentsYoungerSibling"))
	return rv
}


// The label for the contact’s partner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationpartner

func (c_ CNLabeledValue) CNLabelContactRelationPartner() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationPartner"))
	return rv
}


// The label for the contact’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsibling

func (c_ CNLabeledValue) CNLabelContactRelationSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSibling"))
	return rv
}


// The label for the contact’s sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsiblinginlaw

func (c_ CNLabeledValue) CNLabelContactRelationSiblingInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSiblingInLaw"))
	return rv
}


// The label for the contact’s sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsiblingschild

func (c_ CNLabeledValue) CNLabelContactRelationSiblingsChild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSiblingsChild"))
	return rv
}


// The label for the contact’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsister

func (c_ CNLabeledValue) CNLabelContactRelationSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSister"))
	return rv
}


// The label for the contact’s sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlaw

func (c_ CNLabeledValue) CNLabelContactRelationSisterInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLaw"))
	return rv
}


// The label for the contact’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawbrotherswife

func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawBrothersWife() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawBrothersWife"))
	return rv
}


// The label for the contact’s elder brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawelderbrotherswife

func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawElderBrothersWife() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawElderBrothersWife"))
	return rv
}


// The label for the contact’s husband’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawhusbandsbrotherswife

func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawHusbandsBrothersWife() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawHusbandsBrothersWife"))
	return rv
}


// The label for the contact’s husband’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawhusbandssister

func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawHusbandsSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawHusbandsSister"))
	return rv
}


// The label for the contact’s spouse’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawspousessister

func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawSpousesSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawSpousesSister"))
	return rv
}


// The label for the contact’s wife’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawwifesbrotherswife

func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawWifesBrothersWife() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawWifesBrothersWife"))
	return rv
}


// The label for the contact’s wife’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawwifessister

func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawWifesSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawWifesSister"))
	return rv
}


// The label for the contact’s younger brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawyoungerbrotherswife

func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawYoungerBrothersWife() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawYoungerBrothersWife"))
	return rv
}


// The label for the contact’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationson

func (c_ CNLabeledValue) CNLabelContactRelationSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSon"))
	return rv
}


// The label for the contact’s son-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaw

func (c_ CNLabeledValue) CNLabelContactRelationSonInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSonInLaw"))
	return rv
}


// The label for the contact’s son-in-law or brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaworbrotherinlaw

func (c_ CNLabeledValue) CNLabelContactRelationSonInLawOrBrotherInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSonInLawOrBrotherInLaw"))
	return rv
}


// The label for the contact’s son-in-law or stepson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaworstepson

func (c_ CNLabeledValue) CNLabelContactRelationSonInLawOrStepson() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSonInLawOrStepson"))
	return rv
}


// The label for the contact’s spouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationspouse

func (c_ CNLabeledValue) CNLabelContactRelationSpouse() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationSpouse"))
	return rv
}


// The label for the contact’s stepbrother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepbrother

func (c_ CNLabeledValue) CNLabelContactRelationStepbrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepbrother"))
	return rv
}


// The label for the contact’s stepchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepchild

func (c_ CNLabeledValue) CNLabelContactRelationStepchild() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepchild"))
	return rv
}


// The label for the contact’s stepdaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepdaughter

func (c_ CNLabeledValue) CNLabelContactRelationStepdaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepdaughter"))
	return rv
}


// The label for the contact’s stepfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepfather

func (c_ CNLabeledValue) CNLabelContactRelationStepfather() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepfather"))
	return rv
}


// The label for the contact’s stepmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepmother

func (c_ CNLabeledValue) CNLabelContactRelationStepmother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepmother"))
	return rv
}


// The label for the contact’s stepparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepparent

func (c_ CNLabeledValue) CNLabelContactRelationStepparent() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepparent"))
	return rv
}


// The label for the contact’s stepsister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepsister

func (c_ CNLabeledValue) CNLabelContactRelationStepsister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepsister"))
	return rv
}


// The label for the contact’s stepson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepson

func (c_ CNLabeledValue) CNLabelContactRelationStepson() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationStepson"))
	return rv
}


// The label for the contact’s teacher.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationteacher

func (c_ CNLabeledValue) CNLabelContactRelationTeacher() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationTeacher"))
	return rv
}


// The label for the contact’s uncle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncle

func (c_ CNLabeledValue) CNLabelContactRelationUncle() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncle"))
	return rv
}


// The label for the contact’s father’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersbrother

func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersBrother"))
	return rv
}


// The label for the contact’s father’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherselderbrother

func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersElderBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersElderBrother"))
	return rv
}


// The label for the contact’s elder sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherseldersistershusband

func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersElderSistersHusband() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersElderSistersHusband"))
	return rv
}


// The label for the contact’s father’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherssistershusband

func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersSistersHusband() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersSistersHusband"))
	return rv
}


// The label for the contact’s father’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersyoungerbrother

func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersYoungerBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersYoungerBrother"))
	return rv
}


// The label for the contact’s father’s younger sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersyoungersistershusband

func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersYoungerSistersHusband() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersYoungerSistersHusband"))
	return rv
}


// The label for the contact’s mother’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemothersbrother

func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersBrother"))
	return rv
}


// The label for the contact’s mother’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemotherselderbrother

func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersElderBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersElderBrother"))
	return rv
}


// The label for the contact’s mother’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemotherssistershusband

func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersSistersHusband() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersSistersHusband"))
	return rv
}


// The label for the contact’s mother’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemothersyoungerbrother

func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersYoungerBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersYoungerBrother"))
	return rv
}


// The label for the contact’s parent’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentsbrother

func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsBrother"))
	return rv
}


// The label for the contact’s parent’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentselderbrother

func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsElderBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsElderBrother"))
	return rv
}


// The label for the contact’s parent’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentsyoungerbrother

func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsYoungerBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsYoungerBrother"))
	return rv
}


// The label for the contact’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationwife

func (c_ CNLabeledValue) CNLabelContactRelationWife() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationWife"))
	return rv
}


// The label for the contact’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungerbrother

func (c_ CNLabeledValue) CNLabelContactRelationYoungerBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerBrother"))
	return rv
}


// The label for the contact’s younger brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungerbrotherinlaw

func (c_ CNLabeledValue) CNLabelContactRelationYoungerBrotherInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerBrotherInLaw"))
	return rv
}


// The label for the contact’s younger cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousin

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousin() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousin"))
	return rv
}


// The label for the contact’s father’s brother’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfathersbrothersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersBrothersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersBrothersDaughter"))
	return rv
}


// The label for the contact’s father’s brother’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfathersbrothersson

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersBrothersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersBrothersSon"))
	return rv
}


// The label for the contact’s father’s sister’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfatherssistersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersSistersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersSistersDaughter"))
	return rv
}


// The label for the contact’s father’s sister’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfatherssistersson

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersSistersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersSistersSon"))
	return rv
}


// The label for the contact’s mother’s brother’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmothersbrothersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersBrothersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersBrothersDaughter"))
	return rv
}


// The label for the contact’s mother’s brother’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmothersbrothersson

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersBrothersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersBrothersSon"))
	return rv
}


// The label for the contact’s mother’s sibling’s younger daughter or father’s sister’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssiblingsdaughterorfatherssistersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter"))
	return rv
}


// The label for the contact’s mother’s sibling’s younger son or father’s sister’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssiblingssonorfatherssistersson

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon"))
	return rv
}


// The label for the contact’s mother’s sister’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssistersdaughter

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSistersDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSistersDaughter"))
	return rv
}


// The label for the contact’s mother’s sister’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssistersson

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSistersSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSistersSon"))
	return rv
}


// The label for the contact’s parent’s sibling’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinparentssiblingsdaughter

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinParentsSiblingsDaughter() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinParentsSiblingsDaughter"))
	return rv
}


// The label for the contact’s parent’s sibling’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinparentssiblingsson

func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinParentsSiblingsSon() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinParentsSiblingsSon"))
	return rv
}


// The label for the contact’s younger sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersibling

func (c_ CNLabeledValue) CNLabelContactRelationYoungerSibling() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSibling"))
	return rv
}


// The label for the contact’s younger sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersiblinginlaw

func (c_ CNLabeledValue) CNLabelContactRelationYoungerSiblingInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSiblingInLaw"))
	return rv
}


// The label for the contact’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersister

func (c_ CNLabeledValue) CNLabelContactRelationYoungerSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSister"))
	return rv
}


// The label for the contact’s younger sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersisterinlaw

func (c_ CNLabeledValue) CNLabelContactRelationYoungerSisterInLaw() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSisterInLaw"))
	return rv
}


// The label for the contact’s youngest brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungestbrother

func (c_ CNLabeledValue) CNLabelContactRelationYoungestBrother() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungestBrother"))
	return rv
}


// The label for the contact’s youngest sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungestsister

func (c_ CNLabeledValue) CNLabelContactRelationYoungestSister() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelContactRelationYoungestSister"))
	return rv
}


// The label for identifying the contact’s anniversary date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeldateanniversary

func (c_ CNLabeledValue) CNLabelDateAnniversary() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelDateAnniversary"))
	return rv
}


// The label for identifying the contact’s iCloud email information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelemailicloud

func (c_ CNLabeledValue) CNLabelEmailiCloud() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelEmailiCloud"))
	return rv
}


// The label for identifying home information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelhome

func (c_ CNLabeledValue) CNLabelHome() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelHome"))
	return rv
}


// The label for identifying other information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelother

func (c_ CNLabeledValue) CNLabelOther() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelOther"))
	return rv
}


// The label for identifying the contact’s Apple Watch phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberapplewatch

func (c_ CNLabeledValue) CNLabelPhoneNumberAppleWatch() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberAppleWatch"))
	return rv
}


// The label for identifying the contact’s home fax number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberhomefax

func (c_ CNLabeledValue) CNLabelPhoneNumberHomeFax() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberHomeFax"))
	return rv
}


// The label for identifying the contact’s main phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumbermain

func (c_ CNLabeledValue) CNLabelPhoneNumberMain() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberMain"))
	return rv
}


// The label for identifying the contact’s mobile phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumbermobile

func (c_ CNLabeledValue) CNLabelPhoneNumberMobile() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberMobile"))
	return rv
}


// The label for identifying another fax number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberotherfax

func (c_ CNLabeledValue) CNLabelPhoneNumberOtherFax() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberOtherFax"))
	return rv
}


// The label for identifying the contact’s pager number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberpager

func (c_ CNLabeledValue) CNLabelPhoneNumberPager() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberPager"))
	return rv
}


// The label for identifying the contact’s work fax number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberworkfax

func (c_ CNLabeledValue) CNLabelPhoneNumberWorkFax() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberWorkFax"))
	return rv
}


// The label for identifying the contact’s iPhone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberiphone

func (c_ CNLabeledValue) CNLabelPhoneNumberiPhone() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelPhoneNumberiPhone"))
	return rv
}


// The label for the contact’s school.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelschool

func (c_ CNLabeledValue) CNLabelSchool() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelSchool"))
	return rv
}


// The label for identifying URL information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelurladdresshomepage

func (c_ CNLabeledValue) CNLabelURLAddressHomePage() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelURLAddressHomePage"))
	return rv
}


// The label for identifying work information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelwork

func (c_ CNLabeledValue) CNLabelWork() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNLabelWork"))
	return rv
}


