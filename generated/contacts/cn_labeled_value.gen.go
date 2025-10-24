// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNLabeledValue */


/* debug [class_header]: Header for CNLabeledValue */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNLabeledValue */
// An interface definition for the [CNLabeledValue] class.
type ICNLabeledValue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNLabeledValue */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	Label() objc.IObject /* cross-framework: NSString */
	Value() objectivec.IObject
	CNLabelContactRelationAssistant() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAunt() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntFathersBrothersWife() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntFathersElderBrothersWife() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntFathersElderSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntFathersSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntFathersYoungerBrothersWife() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntFathersYoungerSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntMothersBrothersWife() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntMothersElderSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntMothersSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntMothersYoungerSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntParentsElderSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntParentsSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationAuntParentsYoungerSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBoyfriend() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrotherInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrotherInLawElderSistersHusband() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrotherInLawHusbandsBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrotherInLawHusbandsSistersHusband() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrotherInLawSistersHusband() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrotherInLawSpousesBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrotherInLawWifesBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrotherInLawWifesSistersHusband() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationBrotherInLawYoungerSistersHusband() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationChild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationChildInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCoBrotherInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCoFatherInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCoMotherInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCoParentInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCoSiblingInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCoSisterInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationColleague() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousin() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinFathersBrothersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinFathersBrothersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinFathersSistersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinFathersSistersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinGrandparentsSiblingsChild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinGrandparentsSiblingsDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinGrandparentsSiblingsSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinMothersBrothersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinMothersBrothersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinMothersSistersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinMothersSistersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinOrSiblingsChild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinParentsSiblingsChild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinParentsSiblingsDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationCousinParentsSiblingsSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationDaughterInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationDaughterInLawOrSisterInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationDaughterInLawOrStepdaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderBrotherInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousin() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinFathersBrothersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinFathersBrothersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinFathersSistersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinFathersSistersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinMothersBrothersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinMothersBrothersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinMothersSistersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinMothersSistersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinParentsSiblingsDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderCousinParentsSiblingsSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderSiblingInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationElderSisterInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationEldestBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationEldestSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFather() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFatherInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFatherInLawHusbandsFather() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFatherInLawOrStepfather() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFatherInLawWifesFather() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFemaleCousin() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFemaleFriend() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFemalePartner() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationFriend() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGirlfriend() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGirlfriendOrBoyfriend() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandaunt() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandchild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandchildOrSiblingsChild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGranddaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGranddaughterDaughtersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGranddaughterOrNiece() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGranddaughterSonsDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandfather() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandfatherFathersFather() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandfatherMothersFather() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandmother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandmotherFathersMother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandmotherMothersMother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandnephew() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandnephewBrothersGrandson() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandnephewSistersGrandson() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandniece() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandnieceBrothersGranddaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandnieceSistersGranddaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandparent() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandson() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandsonDaughtersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandsonOrNephew() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGrandsonSonsSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGranduncle() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGreatGrandchild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGreatGranddaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGreatGrandfather() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGreatGrandmother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGreatGrandparent() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationGreatGrandson() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationHusband() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationMaleCousin() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationMaleFriend() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationMalePartner() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationManager() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationMother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationMotherInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationMotherInLawHusbandsMother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationMotherInLawOrStepmother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationMotherInLawWifesMother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNephew() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNephewBrothersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNephewOrCousin() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNephewSistersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNiece() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNieceBrothersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNieceOrCousin() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNieceSistersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParent() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParentInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParentsElderSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParentsSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParentsSiblingFathersElderSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParentsSiblingFathersSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParentsSiblingFathersYoungerSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParentsSiblingMothersElderSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParentsSiblingMothersSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParentsSiblingMothersYoungerSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationParentsYoungerSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationPartner() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSiblingInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSiblingsChild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSisterInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSisterInLawBrothersWife() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSisterInLawElderBrothersWife() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSisterInLawHusbandsBrothersWife() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSisterInLawHusbandsSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSisterInLawSpousesSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSisterInLawWifesBrothersWife() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSisterInLawWifesSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSisterInLawYoungerBrothersWife() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSonInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSonInLawOrBrotherInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSonInLawOrStepson() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationSpouse() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationStepbrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationStepchild() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationStepdaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationStepfather() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationStepmother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationStepparent() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationStepsister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationStepson() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationTeacher() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncle() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleFathersBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleFathersElderBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleFathersElderSistersHusband() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleFathersSistersHusband() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleFathersYoungerBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleFathersYoungerSistersHusband() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleMothersBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleMothersElderBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleMothersSistersHusband() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleMothersYoungerBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleParentsBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleParentsElderBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationUncleParentsYoungerBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationWife() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerBrotherInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousin() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinFathersBrothersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinFathersBrothersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinFathersSistersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinFathersSistersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinMothersBrothersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinMothersBrothersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinMothersSistersDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinMothersSistersSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinParentsSiblingsDaughter() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerCousinParentsSiblingsSon() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerSibling() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerSiblingInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerSister() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungerSisterInLaw() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungestBrother() objc.IObject /* cross-framework: NSString */
	CNLabelContactRelationYoungestSister() objc.IObject /* cross-framework: NSString */
	CNLabelDateAnniversary() objc.IObject /* cross-framework: NSString */
	CNLabelEmailiCloud() objc.IObject /* cross-framework: NSString */
	CNLabelHome() objc.IObject /* cross-framework: NSString */
	CNLabelOther() objc.IObject /* cross-framework: NSString */
	CNLabelPhoneNumberAppleWatch() objc.IObject /* cross-framework: NSString */
	CNLabelPhoneNumberHomeFax() objc.IObject /* cross-framework: NSString */
	CNLabelPhoneNumberMain() objc.IObject /* cross-framework: NSString */
	CNLabelPhoneNumberMobile() objc.IObject /* cross-framework: NSString */
	CNLabelPhoneNumberOtherFax() objc.IObject /* cross-framework: NSString */
	CNLabelPhoneNumberPager() objc.IObject /* cross-framework: NSString */
	CNLabelPhoneNumberWorkFax() objc.IObject /* cross-framework: NSString */
	CNLabelPhoneNumberiPhone() objc.IObject /* cross-framework: NSString */
	CNLabelSchool() objc.IObject /* cross-framework: NSString */
	CNLabelURLAddressHomePage() objc.IObject /* cross-framework: NSString */
	CNLabelWork() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNLabeledValue */
	// methods:
	LabeledValueBySettingLabel(label objc.IObject /* cross-framework: NSString */) objectivec.IObject
	LabeledValueBySettingLabelValue(label objc.IObject /* cross-framework: NSString */, value objectivec.IObject) objectivec.IObject
	LabeledValueBySettingValue(value objectivec.IObject) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNLabeledValue */
// Alloc allocates a new instance without initialization.
func (cc _CNLabeledValueClass) Alloc() CNLabeledValue {
	rv := objc.Send[CNLabeledValue](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNLabeledValue */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNLabeledValue */

// Returns a new labeled value identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/init(label:value:)
func NewCNLabeledValueWithLabelValue(label objc.IObject /* cross-framework: NSString */, value objectivec.IObject) CNLabeledValue {
	instance := getCNLabeledValueClass().Alloc()
	rv := objc.Send[CNLabeledValue](instance.ID, objc.Sel("initWithLabel:value:"), label, value)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNLabeledValueWithLabelValue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNLabeledValue */

// Returns a new labeled value identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/labeledValueWithLabel:value:
func (cc _CNLabeledValueClass) LabeledValueWithLabelValue(label objc.IObject /* cross-framework: NSString */, value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("labeledValueWithLabel:value:"), label, value)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LabeledValueWithLabelValue) */


// Returns a localized string for the specified label.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/localizedString(forLabel:)
func (cc _CNLabeledValueClass) LocalizedStringForLabel(label objc.IObject /* cross-framework: NSString */) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForLabel:"), label)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LocalizedStringForLabel) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNLabeledValue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNLabeledValue */

// Returns a labeled value object with an existing value and identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingLabel(_:)
func (c_ CNLabeledValue) LabeledValueBySettingLabel(label objc.IObject /* cross-framework: NSString */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("labeledValueBySettingLabel:"), label)
	return rv
}/* debug [instance_methods/method]: LabeledValueBySettingLabel */


// Returns a labeled value object with the specified label and value with the existing identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingLabel(_:value:)
func (c_ CNLabeledValue) LabeledValueBySettingLabelValue(label objc.IObject /* cross-framework: NSString */, value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("labeledValueBySettingLabel:value:"), label, value)
	return rv
}/* debug [instance_methods/method]: LabeledValueBySettingLabelValue */


// Returns a new value for an existing label and identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingValue(_:)
func (c_ CNLabeledValue) LabeledValueBySettingValue(value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("labeledValueBySettingValue:"), value)
	return rv
}/* debug [instance_methods/method]: LabeledValueBySettingValue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNLabeledValue */

// A unique identifier for the labeled value object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/identifier
func (c_ CNLabeledValue) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The label for a contact property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/label
func (c_ CNLabeledValue) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A contact property value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/value
func (c_ CNLabeledValue) Value() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("value"))
	return rv
}/* debug [instance_properties/getter]: value */


// The label for the contact’s assistant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationassistant
func (c_ CNLabeledValue) CNLabelContactRelationAssistant() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAssistant"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAssistant */


// The label for the contact’s aunt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationaunt
func (c_ CNLabeledValue) CNLabelContactRelationAunt() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAunt"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAunt */


// The label for the contact’s father’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersBrothersWife() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersBrothersWife"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntFathersBrothersWife */


// The label for the contact’s father’s elder brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherselderbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersElderBrothersWife() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersElderBrothersWife"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntFathersElderBrothersWife */


// The label for the contact’s father’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherseldersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersElderSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersElderSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntFathersElderSister */


// The label for the contact’s father’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherssister
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntFathersSister */


// The label for the contact’s father’s younger brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersyoungerbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersYoungerBrothersWife() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersYoungerBrothersWife"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntFathersYoungerBrothersWife */


// The label for the contact’s father’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersYoungerSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersYoungerSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntFathersYoungerSister */


// The label for the contact’s mother’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmothersbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersBrothersWife() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersBrothersWife"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntMothersBrothersWife */


// The label for the contact’s mother’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmotherseldersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersElderSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersElderSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntMothersElderSister */


// The label for the contact’s mother’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmotherssister
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntMothersSister */


// The label for the contact’s mother’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmothersyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersYoungerSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersYoungerSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntMothersYoungerSister */


// The label for the contact’s parent’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentseldersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsElderSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsElderSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntParentsElderSister */


// The label for the contact’s parent’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentssister
func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntParentsSister */


// The label for the contact’s parent’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentsyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsYoungerSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsYoungerSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationAuntParentsYoungerSister */


// The label for the contact’s boyfriend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationboyfriend
func (c_ CNLabeledValue) CNLabelContactRelationBoyfriend() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBoyfriend"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBoyfriend */


// The label for the contact’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBrother */


// The label for the contact’s brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBrotherInLaw */


// The label for the contact’s elder sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlaweldersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawElderSistersHusband() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawElderSistersHusband"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBrotherInLawElderSistersHusband */


// The label for the contact’s husband’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawhusbandsbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawHusbandsBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawHusbandsBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBrotherInLawHusbandsBrother */


// The label for the contact’s husband’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawhusbandssistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawHusbandsSistersHusband() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawHusbandsSistersHusband"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBrotherInLawHusbandsSistersHusband */


// The label for the contact’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawsistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawSistersHusband() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawSistersHusband"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBrotherInLawSistersHusband */


// The label for the contact’s spouse’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawspousesbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawSpousesBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawSpousesBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBrotherInLawSpousesBrother */


// The label for the contact’s wife’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawwifesbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawWifesBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawWifesBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBrotherInLawWifesBrother */


// The label for the contact’s wife’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawwifessistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawWifesSistersHusband() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawWifesSistersHusband"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBrotherInLawWifesSistersHusband */


// The label for the contact’s younger sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawyoungersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawYoungerSistersHusband() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawYoungerSistersHusband"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationBrotherInLawYoungerSistersHusband */


// The label for the contact’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationchild
func (c_ CNLabeledValue) CNLabelContactRelationChild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationChild"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationChild */


// The label for the contact’s child-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationchildinlaw
func (c_ CNLabeledValue) CNLabelContactRelationChildInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationChildInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationChildInLaw */


// The label for the contact’s co-brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcobrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoBrotherInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCoBrotherInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCoBrotherInLaw */


// The label for the contact’s co-father-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcofatherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoFatherInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCoFatherInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCoFatherInLaw */


// The label for the contact’s co-mother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcomotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoMotherInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCoMotherInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCoMotherInLaw */


// The label for the contact’s co-parent-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcoparentinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoParentInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCoParentInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCoParentInLaw */


// The label for the contact’s co-sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcosiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoSiblingInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCoSiblingInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCoSiblingInLaw */


// The label for the contact’s co-sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcosisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoSisterInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCoSisterInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCoSisterInLaw */


// The label for the contact’s colleague.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcolleague
func (c_ CNLabeledValue) CNLabelContactRelationColleague() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationColleague"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationColleague */


// The label for the contact’s cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousin
func (c_ CNLabeledValue) CNLabelContactRelationCousin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousin"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousin */


// The label for the contact’s father’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfathersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersBrothersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersBrothersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinFathersBrothersDaughter */


// The label for the contact’s father’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfathersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersBrothersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersBrothersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinFathersBrothersSon */


// The label for the contact’s father’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersSistersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersSistersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinFathersSistersDaughter */


// The label for the contact’s father’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersSistersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersSistersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinFathersSistersSon */


// The label for the contact’s grandparent’s sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsChild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsChild"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinGrandparentsSiblingsChild */


// The label for the contact’s grandparent’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinGrandparentsSiblingsDaughter */


// The label for the contact’s grandparent’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinGrandparentsSiblingsSon */


// The label for the contact’s mother’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmothersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersBrothersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersBrothersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinMothersBrothersDaughter */


// The label for the contact’s mother’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmothersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersBrothersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersBrothersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinMothersBrothersSon */


// The label for the contact’s mother’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmotherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersSistersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersSistersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinMothersSistersDaughter */


// The label for the contact’s mother’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmotherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersSistersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersSistersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinMothersSistersSon */


// The label for the contact’s cousin’s or sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinorsiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationCousinOrSiblingsChild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinOrSiblingsChild"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinOrSiblingsChild */


// The label for the contact’s parent’s sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsChild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsChild"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinParentsSiblingsChild */


// The label for the contact’s parent’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinParentsSiblingsDaughter */


// The label for the contact’s parent’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationCousinParentsSiblingsSon */


// The label for the contact’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughter
func (c_ CNLabeledValue) CNLabelContactRelationDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationDaughter */


// The label for the contact’s daughter-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationDaughterInLaw */


// The label for the contact’s daughter-in-law or sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaworsisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLawOrSisterInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLawOrSisterInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationDaughterInLawOrSisterInLaw */


// The label for the contact’s daughter-in-law or stepdaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaworstepdaughter
func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLawOrStepdaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLawOrStepdaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationDaughterInLawOrStepdaughter */


// The label for the contact’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationelderbrother
func (c_ CNLabeledValue) CNLabelContactRelationElderBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderBrother */


// The label for the contact’s elder brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationelderbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationElderBrotherInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderBrotherInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderBrotherInLaw */


// The label for the contact’s elder cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousin
func (c_ CNLabeledValue) CNLabelContactRelationElderCousin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousin"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousin */


// The label for the contact’s father’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfathersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersBrothersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersBrothersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinFathersBrothersDaughter */


// The label for the contact’s father’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfathersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersBrothersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersBrothersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinFathersBrothersSon */


// The label for the contact’s father’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersSistersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersSistersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinFathersSistersDaughter */


// The label for the contact’s father’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersSistersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersSistersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinFathersSistersSon */


// The label for the contact’s mother’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmothersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersBrothersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersBrothersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinMothersBrothersDaughter */


// The label for the contact’s mother’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmothersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersBrothersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersBrothersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinMothersBrothersSon */


// The label for the contact’s mother’s sibling’s daughter or father’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssiblingsdaughterorfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter */


// The label for the contact’s mother’s sibling’s son or father’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssiblingssonorfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon */


// The label for the contact’s mother’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSistersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSistersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinMothersSistersDaughter */


// The label for the contact’s mother’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSistersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSistersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinMothersSistersSon */


// The label for the contact’s parent’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinParentsSiblingsDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinParentsSiblingsDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinParentsSiblingsDaughter */


// The label for the contact’s parent’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinParentsSiblingsSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderCousinParentsSiblingsSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderCousinParentsSiblingsSon */


// The label for the contact’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersibling
func (c_ CNLabeledValue) CNLabelContactRelationElderSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderSibling */


// The label for the contact’s elder sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationElderSiblingInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderSiblingInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderSiblingInLaw */


// The label for the contact’s elder sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersister
func (c_ CNLabeledValue) CNLabelContactRelationElderSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderSister */


// The label for the contact’s elder sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationElderSisterInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationElderSisterInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationElderSisterInLaw */


// The label for the contact’s eldest brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldestbrother
func (c_ CNLabeledValue) CNLabelContactRelationEldestBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationEldestBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationEldestBrother */


// The label for the contact’s eldest sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldestsister
func (c_ CNLabeledValue) CNLabelContactRelationEldestSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationEldestSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationEldestSister */


// The label for the contact’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfather
func (c_ CNLabeledValue) CNLabelContactRelationFather() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFather"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationFather */


// The label for the contact’s father-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFatherInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationFatherInLaw */


// The label for the contact’s husband’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlawhusbandsfather
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawHusbandsFather() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawHusbandsFather"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationFatherInLawHusbandsFather */


// The label for the contact’s father-in-law or stepfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlaworstepfather
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawOrStepfather() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawOrStepfather"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationFatherInLawOrStepfather */


// The label for the contact’s wife’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlawwifesfather
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawWifesFather() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawWifesFather"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationFatherInLawWifesFather */


// The label for the contact’s female cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalecousin
func (c_ CNLabeledValue) CNLabelContactRelationFemaleCousin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFemaleCousin"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationFemaleCousin */


// The label for the contact’s female friend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalefriend
func (c_ CNLabeledValue) CNLabelContactRelationFemaleFriend() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFemaleFriend"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationFemaleFriend */


// The label for the contact’s female partner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalepartner
func (c_ CNLabeledValue) CNLabelContactRelationFemalePartner() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFemalePartner"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationFemalePartner */


// The label for the contact’s friend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfriend
func (c_ CNLabeledValue) CNLabelContactRelationFriend() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationFriend"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationFriend */


// The label for the contact’s girlfriend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgirlfriend
func (c_ CNLabeledValue) CNLabelContactRelationGirlfriend() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGirlfriend"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGirlfriend */


// The label for the contact’s girlfriend or boyfriend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgirlfriendorboyfriend
func (c_ CNLabeledValue) CNLabelContactRelationGirlfriendOrBoyfriend() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGirlfriendOrBoyfriend"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGirlfriendOrBoyfriend */


// The label for the contact’s grandaunt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandaunt
func (c_ CNLabeledValue) CNLabelContactRelationGrandaunt() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandaunt"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandaunt */


// The label for the contact’s grandchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandchild
func (c_ CNLabeledValue) CNLabelContactRelationGrandchild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandchild"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandchild */


// The label for the contact’s grandchild or sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandchildorsiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationGrandchildOrSiblingsChild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandchildOrSiblingsChild"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandchildOrSiblingsChild */


// The label for the contact’s granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGranddaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGranddaughter */


// The label for the contact’s daughter’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughterdaughtersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterDaughtersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterDaughtersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGranddaughterDaughtersDaughter */


// The label for the contact’s granddaughter or niece.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughterorniece
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterOrNiece() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterOrNiece"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGranddaughterOrNiece */


// The label for the contact’s son’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughtersonsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterSonsDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterSonsDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGranddaughterSonsDaughter */


// The label for the contact’s grandfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfather
func (c_ CNLabeledValue) CNLabelContactRelationGrandfather() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandfather"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandfather */


// The label for the contact’s father’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfatherfathersfather
func (c_ CNLabeledValue) CNLabelContactRelationGrandfatherFathersFather() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandfatherFathersFather"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandfatherFathersFather */


// The label for the contact’s mother’s father.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfathermothersfather
func (c_ CNLabeledValue) CNLabelContactRelationGrandfatherMothersFather() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandfatherMothersFather"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandfatherMothersFather */


// The label for the contact’s grandmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmother
func (c_ CNLabeledValue) CNLabelContactRelationGrandmother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandmother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandmother */


// The label for the contact’s father’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmotherfathersmother
func (c_ CNLabeledValue) CNLabelContactRelationGrandmotherFathersMother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandmotherFathersMother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandmotherFathersMother */


// The label for the contact’s mother’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmothermothersmother
func (c_ CNLabeledValue) CNLabelContactRelationGrandmotherMothersMother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandmotherMothersMother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandmotherMothersMother */


// The label for the contact’s grandnephew.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephew
func (c_ CNLabeledValue) CNLabelContactRelationGrandnephew() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandnephew"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandnephew */


// The label for the contact’s brother’s grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephewbrothersgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGrandnephewBrothersGrandson() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandnephewBrothersGrandson"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandnephewBrothersGrandson */


// The label for the contact’s sister’s grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephewsistersgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGrandnephewSistersGrandson() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandnephewSistersGrandson"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandnephewSistersGrandson */


// The label for the contact’s grandniece.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniece
func (c_ CNLabeledValue) CNLabelContactRelationGrandniece() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandniece"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandniece */


// The label for the contact’s brother’s granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniecebrothersgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGrandnieceBrothersGranddaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandnieceBrothersGranddaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandnieceBrothersGranddaughter */


// The label for the contact’s sister’s granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniecesistersgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGrandnieceSistersGranddaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandnieceSistersGranddaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandnieceSistersGranddaughter */


// The label for the contact’s grandparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandparent
func (c_ CNLabeledValue) CNLabelContactRelationGrandparent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandparent"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandparent */


// The label for the contact’s grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGrandson() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandson"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandson */


// The label for the contact’s daughter’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsondaughtersson
func (c_ CNLabeledValue) CNLabelContactRelationGrandsonDaughtersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandsonDaughtersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandsonDaughtersSon */


// The label for the contact’s grandson or nephew.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsonornephew
func (c_ CNLabeledValue) CNLabelContactRelationGrandsonOrNephew() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandsonOrNephew"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandsonOrNephew */


// The label for the contact’s son’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsonsonsson
func (c_ CNLabeledValue) CNLabelContactRelationGrandsonSonsSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGrandsonSonsSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGrandsonSonsSon */


// The label for the contact’s granduncle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranduncle
func (c_ CNLabeledValue) CNLabelContactRelationGranduncle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGranduncle"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGranduncle */


// The label for the contact’s grandchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandchild
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandchild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandchild"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGreatGrandchild */


// The label for the contact’s grandchild or sibling’s grandchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandchildorsiblingsgrandchild
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild */


// The label for the contact’s great-granddaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGreatGranddaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGreatGranddaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGreatGranddaughter */


// The label for the contact’s great-grandfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandfather
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandfather() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandfather"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGreatGrandfather */


// The label for the contact’s great-grandmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandmother
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandmother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandmother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGreatGrandmother */


// The label for the contact’s great-grandparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandparent
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandparent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandparent"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGreatGrandparent */


// The label for the contact’s great-grandson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandson() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandson"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationGreatGrandson */


// The label for the contact’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationhusband
func (c_ CNLabeledValue) CNLabelContactRelationHusband() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationHusband"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationHusband */


// The label for the contact’s male cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalecousin
func (c_ CNLabeledValue) CNLabelContactRelationMaleCousin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationMaleCousin"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationMaleCousin */


// The label for the contact’s male friend.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalefriend
func (c_ CNLabeledValue) CNLabelContactRelationMaleFriend() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationMaleFriend"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationMaleFriend */


// The label for the contact’s male partner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalepartner
func (c_ CNLabeledValue) CNLabelContactRelationMalePartner() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationMalePartner"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationMalePartner */


// The label for the contact’s manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmanager
func (c_ CNLabeledValue) CNLabelContactRelationManager() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationManager"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationManager */


// The label for the contact’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmother
func (c_ CNLabeledValue) CNLabelContactRelationMother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationMother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationMother */


// The label for the contact’s mother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationMotherInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationMotherInLaw */


// The label for the contact’s husband’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlawhusbandsmother
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawHusbandsMother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawHusbandsMother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationMotherInLawHusbandsMother */


// The label for the contact’s mother-in-law or stepmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlaworstepmother
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawOrStepmother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawOrStepmother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationMotherInLawOrStepmother */


// The label for the contact’s wife’s mother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlawwifesmother
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawWifesMother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawWifesMother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationMotherInLawWifesMother */


// The label for the contact’s nephew.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephew
func (c_ CNLabeledValue) CNLabelContactRelationNephew() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNephew"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNephew */


// The label for the contact’s brother’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationNephewBrothersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNephewBrothersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNephewBrothersSon */


// The label for the contact’s brother’s son or husband’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewbrotherssonorhusbandssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon */


// The label for the contact’s nephew or cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnepheworcousin
func (c_ CNLabeledValue) CNLabelContactRelationNephewOrCousin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNephewOrCousin"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNephewOrCousin */


// The label for the contact’s sister’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewsistersson
func (c_ CNLabeledValue) CNLabelContactRelationNephewSistersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNephewSistersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNephewSistersSon */


// The label for the contact’s sister’s son or wife’s sibling’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewsisterssonorwifessiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon */


// The label for the contact’s niece.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniece
func (c_ CNLabeledValue) CNLabelContactRelationNiece() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNiece"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNiece */


// The label for the contact’s brother’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecebrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceBrothersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNieceBrothersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNieceBrothersDaughter */


// The label for the contact’s brother’s daughter or husband’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecebrothersdaughterorhusbandssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter */


// The label for the contact’s niece or cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnieceorcousin
func (c_ CNLabeledValue) CNLabelContactRelationNieceOrCousin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNieceOrCousin"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNieceOrCousin */


// The label for the contact’s sister’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecesistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceSistersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNieceSistersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNieceSistersDaughter */


// The label for the contact’s sister’s daughter or wife’s sibling’s daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecesistersdaughterorwifessiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter */


// The label for the contact’s parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparent
func (c_ CNLabeledValue) CNLabelContactRelationParent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParent"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParent */


// The label for the contact’s parent-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentinlaw
func (c_ CNLabeledValue) CNLabelContactRelationParentInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParentInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParentInLaw */


// The label for the contact’s parent’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentseldersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsElderSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParentsElderSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParentsElderSibling */


// The label for the contact’s parent’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParentsSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParentsSibling */


// The label for the contact’s father’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfatherseldersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersElderSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersElderSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParentsSiblingFathersElderSibling */


// The label for the contact’s father’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfatherssibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParentsSiblingFathersSibling */


// The label for the contact’s father’s youngest sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfathersyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersYoungerSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersYoungerSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParentsSiblingFathersYoungerSibling */


// The label for the contact’s mother’s elder sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmotherseldersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersElderSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersElderSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParentsSiblingMothersElderSibling */


// The label for the contact’s mother’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmotherssibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParentsSiblingMothersSibling */


// The label for the contact’s mother’s younger sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmothersyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersYoungerSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersYoungerSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParentsSiblingMothersYoungerSibling */


// The label for the contact’s parent’s younger sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentsyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsYoungerSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationParentsYoungerSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationParentsYoungerSibling */


// The label for the contact’s partner.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationpartner
func (c_ CNLabeledValue) CNLabelContactRelationPartner() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationPartner"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationPartner */


// The label for the contact’s sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsibling
func (c_ CNLabeledValue) CNLabelContactRelationSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSibling */


// The label for the contact’s sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationSiblingInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSiblingInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSiblingInLaw */


// The label for the contact’s sibling’s child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationSiblingsChild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSiblingsChild"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSiblingsChild */


// The label for the contact’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsister
func (c_ CNLabeledValue) CNLabelContactRelationSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSister */


// The label for the contact’s sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSisterInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSisterInLaw */


// The label for the contact’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawBrothersWife() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawBrothersWife"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSisterInLawBrothersWife */


// The label for the contact’s elder brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawelderbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawElderBrothersWife() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawElderBrothersWife"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSisterInLawElderBrothersWife */


// The label for the contact’s husband’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawhusbandsbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawHusbandsBrothersWife() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawHusbandsBrothersWife"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSisterInLawHusbandsBrothersWife */


// The label for the contact’s husband’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawhusbandssister
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawHusbandsSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawHusbandsSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSisterInLawHusbandsSister */


// The label for the contact’s spouse’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawspousessister
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawSpousesSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawSpousesSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSisterInLawSpousesSister */


// The label for the contact’s wife’s brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawwifesbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawWifesBrothersWife() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawWifesBrothersWife"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSisterInLawWifesBrothersWife */


// The label for the contact’s wife’s sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawwifessister
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawWifesSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawWifesSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSisterInLawWifesSister */


// The label for the contact’s younger brother’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawyoungerbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawYoungerBrothersWife() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawYoungerBrothersWife"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSisterInLawYoungerBrothersWife */


// The label for the contact’s son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationson
func (c_ CNLabeledValue) CNLabelContactRelationSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSon */


// The label for the contact’s son-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaw
func (c_ CNLabeledValue) CNLabelContactRelationSonInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSonInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSonInLaw */


// The label for the contact’s son-in-law or brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaworbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationSonInLawOrBrotherInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSonInLawOrBrotherInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSonInLawOrBrotherInLaw */


// The label for the contact’s son-in-law or stepson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaworstepson
func (c_ CNLabeledValue) CNLabelContactRelationSonInLawOrStepson() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSonInLawOrStepson"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSonInLawOrStepson */


// The label for the contact’s spouse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationspouse
func (c_ CNLabeledValue) CNLabelContactRelationSpouse() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationSpouse"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationSpouse */


// The label for the contact’s stepbrother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepbrother
func (c_ CNLabeledValue) CNLabelContactRelationStepbrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationStepbrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationStepbrother */


// The label for the contact’s stepchild.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepchild
func (c_ CNLabeledValue) CNLabelContactRelationStepchild() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationStepchild"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationStepchild */


// The label for the contact’s stepdaughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepdaughter
func (c_ CNLabeledValue) CNLabelContactRelationStepdaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationStepdaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationStepdaughter */


// The label for the contact’s stepfather.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepfather
func (c_ CNLabeledValue) CNLabelContactRelationStepfather() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationStepfather"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationStepfather */


// The label for the contact’s stepmother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepmother
func (c_ CNLabeledValue) CNLabelContactRelationStepmother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationStepmother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationStepmother */


// The label for the contact’s stepparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepparent
func (c_ CNLabeledValue) CNLabelContactRelationStepparent() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationStepparent"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationStepparent */


// The label for the contact’s stepsister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepsister
func (c_ CNLabeledValue) CNLabelContactRelationStepsister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationStepsister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationStepsister */


// The label for the contact’s stepson.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepson
func (c_ CNLabeledValue) CNLabelContactRelationStepson() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationStepson"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationStepson */


// The label for the contact’s teacher.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationteacher
func (c_ CNLabeledValue) CNLabelContactRelationTeacher() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationTeacher"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationTeacher */


// The label for the contact’s uncle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncle
func (c_ CNLabeledValue) CNLabelContactRelationUncle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncle"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncle */


// The label for the contact’s father’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleFathersBrother */


// The label for the contact’s father’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherselderbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersElderBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersElderBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleFathersElderBrother */


// The label for the contact’s elder sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherseldersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersElderSistersHusband() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersElderSistersHusband"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleFathersElderSistersHusband */


// The label for the contact’s father’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherssistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersSistersHusband() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersSistersHusband"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleFathersSistersHusband */


// The label for the contact’s father’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersYoungerBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersYoungerBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleFathersYoungerBrother */


// The label for the contact’s father’s younger sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersyoungersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersYoungerSistersHusband() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersYoungerSistersHusband"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleFathersYoungerSistersHusband */


// The label for the contact’s mother’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemothersbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleMothersBrother */


// The label for the contact’s mother’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemotherselderbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersElderBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersElderBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleMothersElderBrother */


// The label for the contact’s mother’s sister’s husband.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemotherssistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersSistersHusband() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersSistersHusband"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleMothersSistersHusband */


// The label for the contact’s mother’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemothersyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersYoungerBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersYoungerBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleMothersYoungerBrother */


// The label for the contact’s parent’s brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentsbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleParentsBrother */


// The label for the contact’s parent’s elder brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentselderbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsElderBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsElderBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleParentsElderBrother */


// The label for the contact’s parent’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentsyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsYoungerBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsYoungerBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationUncleParentsYoungerBrother */


// The label for the contact’s wife.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationwife
func (c_ CNLabeledValue) CNLabelContactRelationWife() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationWife"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationWife */


// The label for the contact’s younger brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationYoungerBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerBrother */


// The label for the contact’s younger brother-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungerbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationYoungerBrotherInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerBrotherInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerBrotherInLaw */


// The label for the contact’s younger cousin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousin
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousin() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousin"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousin */


// The label for the contact’s father’s brother’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfathersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersBrothersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersBrothersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinFathersBrothersDaughter */


// The label for the contact’s father’s brother’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfathersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersBrothersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersBrothersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinFathersBrothersSon */


// The label for the contact’s father’s sister’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersSistersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersSistersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinFathersSistersDaughter */


// The label for the contact’s father’s sister’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersSistersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersSistersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinFathersSistersSon */


// The label for the contact’s mother’s brother’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmothersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersBrothersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersBrothersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinMothersBrothersDaughter */


// The label for the contact’s mother’s brother’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmothersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersBrothersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersBrothersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinMothersBrothersSon */


// The label for the contact’s mother’s sibling’s younger daughter or father’s sister’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssiblingsdaughterorfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter */


// The label for the contact’s mother’s sibling’s younger son or father’s sister’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssiblingssonorfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon */


// The label for the contact’s mother’s sister’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSistersDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSistersDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinMothersSistersDaughter */


// The label for the contact’s mother’s sister’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSistersSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSistersSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinMothersSistersSon */


// The label for the contact’s parent’s sibling’s younger daughter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinParentsSiblingsDaughter() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinParentsSiblingsDaughter"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinParentsSiblingsDaughter */


// The label for the contact’s parent’s sibling’s younger son.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinParentsSiblingsSon() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinParentsSiblingsSon"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerCousinParentsSiblingsSon */


// The label for the contact’s younger sibling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSibling() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerSibling"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerSibling */


// The label for the contact’s younger sibling-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSiblingInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerSiblingInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerSiblingInLaw */


// The label for the contact’s younger sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerSister */


// The label for the contact’s younger sister-in-law.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSisterInLaw() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungerSisterInLaw"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungerSisterInLaw */


// The label for the contact’s youngest brother.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungestbrother
func (c_ CNLabeledValue) CNLabelContactRelationYoungestBrother() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungestBrother"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungestBrother */


// The label for the contact’s youngest sister.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungestsister
func (c_ CNLabeledValue) CNLabelContactRelationYoungestSister() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelContactRelationYoungestSister"))
	return rv
}/* debug [instance_properties/getter]: CNLabelContactRelationYoungestSister */


// The label for identifying the contact’s anniversary date.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeldateanniversary
func (c_ CNLabeledValue) CNLabelDateAnniversary() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelDateAnniversary"))
	return rv
}/* debug [instance_properties/getter]: CNLabelDateAnniversary */


// The label for identifying the contact’s iCloud email information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelemailicloud
func (c_ CNLabeledValue) CNLabelEmailiCloud() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelEmailiCloud"))
	return rv
}/* debug [instance_properties/getter]: CNLabelEmailiCloud */


// The label for identifying home information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelhome
func (c_ CNLabeledValue) CNLabelHome() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelHome"))
	return rv
}/* debug [instance_properties/getter]: CNLabelHome */


// The label for identifying other information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelother
func (c_ CNLabeledValue) CNLabelOther() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelOther"))
	return rv
}/* debug [instance_properties/getter]: CNLabelOther */


// The label for identifying the contact’s Apple Watch phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberapplewatch
func (c_ CNLabeledValue) CNLabelPhoneNumberAppleWatch() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelPhoneNumberAppleWatch"))
	return rv
}/* debug [instance_properties/getter]: CNLabelPhoneNumberAppleWatch */


// The label for identifying the contact’s home fax number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberhomefax
func (c_ CNLabeledValue) CNLabelPhoneNumberHomeFax() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelPhoneNumberHomeFax"))
	return rv
}/* debug [instance_properties/getter]: CNLabelPhoneNumberHomeFax */


// The label for identifying the contact’s main phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumbermain
func (c_ CNLabeledValue) CNLabelPhoneNumberMain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelPhoneNumberMain"))
	return rv
}/* debug [instance_properties/getter]: CNLabelPhoneNumberMain */


// The label for identifying the contact’s mobile phone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumbermobile
func (c_ CNLabeledValue) CNLabelPhoneNumberMobile() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelPhoneNumberMobile"))
	return rv
}/* debug [instance_properties/getter]: CNLabelPhoneNumberMobile */


// The label for identifying another fax number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberotherfax
func (c_ CNLabeledValue) CNLabelPhoneNumberOtherFax() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelPhoneNumberOtherFax"))
	return rv
}/* debug [instance_properties/getter]: CNLabelPhoneNumberOtherFax */


// The label for identifying the contact’s pager number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberpager
func (c_ CNLabeledValue) CNLabelPhoneNumberPager() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelPhoneNumberPager"))
	return rv
}/* debug [instance_properties/getter]: CNLabelPhoneNumberPager */


// The label for identifying the contact’s work fax number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberworkfax
func (c_ CNLabeledValue) CNLabelPhoneNumberWorkFax() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelPhoneNumberWorkFax"))
	return rv
}/* debug [instance_properties/getter]: CNLabelPhoneNumberWorkFax */


// The label for identifying the contact’s iPhone number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberiphone
func (c_ CNLabeledValue) CNLabelPhoneNumberiPhone() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelPhoneNumberiPhone"))
	return rv
}/* debug [instance_properties/getter]: CNLabelPhoneNumberiPhone */


// The label for the contact’s school.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelschool
func (c_ CNLabeledValue) CNLabelSchool() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelSchool"))
	return rv
}/* debug [instance_properties/getter]: CNLabelSchool */


// The label for identifying URL information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelurladdresshomepage
func (c_ CNLabeledValue) CNLabelURLAddressHomePage() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelURLAddressHomePage"))
	return rv
}/* debug [instance_properties/getter]: CNLabelURLAddressHomePage */


// The label for identifying work information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelwork
func (c_ CNLabeledValue) CNLabelWork() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CNLabelWork"))
	return rv
}/* debug [instance_properties/getter]: CNLabelWork */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNLabeledValue */


