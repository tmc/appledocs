// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
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
	LabeledValueBySettingLabel(label appkit.string) unsafe.Pointer
	LabeledValueBySettingLabelValue(label appkit.string, value unsafe.Pointer) unsafe.Pointer
	LabeledValueBySettingValue(value unsafe.Pointer) unsafe.Pointer
}

// An immutable object that combines a contact property value with a label that describes that property.
//
// Labels describe the context for a property. For example, the label for a phone number indicates whether it corresponds to the user’s home, work, or iPhone number. objects are thread-safe, and you can access their properties from any thread of your app.
//
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
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/init(label:value:)
func NewCNLabeledValueWithLabelValue(label appkit.string, value unsafe.Pointer) CNLabeledValue {
	instance := getCNLabeledValueClass().Alloc()
	rv := objc.Send[CNLabeledValue](instance.ID, objc.Sel("initWithLabel:value:"), label, value)
	rv.Autorelease()
	return rv
}


// Returns a new labeled value identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/labeledValueWithLabel:value:
func (cc _CNLabeledValueClass) LabeledValueWithLabelValue(label appkit.string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("labeledValueWithLabel:value:"), label, value)
	return rv
}

// Returns a localized string for the specified label.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/localizedString(forLabel:)
func (cc _CNLabeledValueClass) LocalizedStringForLabel(label appkit.string) foundation.String {
	rv := objc.Send[foundation.String](objc.ID(cc.class), objc.Sel("localizedStringForLabel:"), label)
	return rv
}

// Returns a labeled value object with an existing value and identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingLabel(_:)
func (c_ CNLabeledValue) LabeledValueBySettingLabel(label appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("labeledValueBySettingLabel:"), label)
	return rv
}

// Returns a labeled value object with the specified label and value with the existing identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingLabel(_:value:)
func (c_ CNLabeledValue) LabeledValueBySettingLabelValue(label appkit.string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("labeledValueBySettingLabel:value:"), label, value)
	return rv
}

// Returns a new value for an existing label and identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingValue(_:)
func (c_ CNLabeledValue) LabeledValueBySettingValue(value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("labeledValueBySettingValue:"), value)
	return rv
}

// A unique identifier for the labeled value object.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/identifier
func (c_ CNLabeledValue) Identifier() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("identifier"))
	return rv
}

// The label for a contact property value.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/label
func (c_ CNLabeledValue) Label() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("label"))
	return rv
}

// A contact property value.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/value
func (c_ CNLabeledValue) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("value"))
	return rv
}

// The label for the contact’s assistant.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationassistant
func (c_ CNLabeledValue) CNLabelContactRelationAssistant() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAssistant"))
	return rv
}

// The label for the contact’s aunt.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationaunt
func (c_ CNLabeledValue) CNLabelContactRelationAunt() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAunt"))
	return rv
}

// The label for the contact’s father’s brother’s wife.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersBrothersWife() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersBrothersWife"))
	return rv
}

// The label for the contact’s father’s elder brother’s wife.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherselderbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersElderBrothersWife() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersElderBrothersWife"))
	return rv
}

// The label for the contact’s father’s elder sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherseldersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersElderSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersElderSister"))
	return rv
}

// The label for the contact’s father’s sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfatherssister
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersSister"))
	return rv
}

// The label for the contact’s father’s younger brother’s wife.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersyoungerbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersYoungerBrothersWife() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersYoungerBrothersWife"))
	return rv
}

// The label for the contact’s father’s younger sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntfathersyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntFathersYoungerSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntFathersYoungerSister"))
	return rv
}

// The label for the contact’s mother’s brother’s wife.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmothersbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersBrothersWife() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersBrothersWife"))
	return rv
}

// The label for the contact’s mother’s elder sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmotherseldersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersElderSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersElderSister"))
	return rv
}

// The label for the contact’s mother’s sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmotherssister
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersSister"))
	return rv
}

// The label for the contact’s mother’s younger sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntmothersyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntMothersYoungerSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntMothersYoungerSister"))
	return rv
}

// The label for the contact’s parent’s elder sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentseldersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsElderSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsElderSister"))
	return rv
}

// The label for the contact’s parent’s sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentssister
func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsSister"))
	return rv
}

// The label for the contact’s parent’s younger sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationauntparentsyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationAuntParentsYoungerSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationAuntParentsYoungerSister"))
	return rv
}

// The label for the contact’s boyfriend.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationboyfriend
func (c_ CNLabeledValue) CNLabelContactRelationBoyfriend() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBoyfriend"))
	return rv
}

// The label for the contact’s brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBrother"))
	return rv
}

// The label for the contact’s brother-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLaw"))
	return rv
}

// The label for the contact’s elder sister’s husband.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlaweldersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawElderSistersHusband() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawElderSistersHusband"))
	return rv
}

// The label for the contact’s husband’s brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawhusbandsbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawHusbandsBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawHusbandsBrother"))
	return rv
}

// The label for the contact’s husband’s sister’s husband.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawhusbandssistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawHusbandsSistersHusband() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawHusbandsSistersHusband"))
	return rv
}

// The label for the contact’s sister’s husband.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawsistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawSistersHusband() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawSistersHusband"))
	return rv
}

// The label for the contact’s spouse’s brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawspousesbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawSpousesBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawSpousesBrother"))
	return rv
}

// The label for the contact’s wife’s brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawwifesbrother
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawWifesBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawWifesBrother"))
	return rv
}

// The label for the contact’s wife’s sister’s husband.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawwifessistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawWifesSistersHusband() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawWifesSistersHusband"))
	return rv
}

// The label for the contact’s younger sister’s husband.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationbrotherinlawyoungersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationBrotherInLawYoungerSistersHusband() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationBrotherInLawYoungerSistersHusband"))
	return rv
}

// The label for the contact’s child.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationchild
func (c_ CNLabeledValue) CNLabelContactRelationChild() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationChild"))
	return rv
}

// The label for the contact’s child-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationchildinlaw
func (c_ CNLabeledValue) CNLabelContactRelationChildInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationChildInLaw"))
	return rv
}

// The label for the contact’s co-brother-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcobrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoBrotherInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCoBrotherInLaw"))
	return rv
}

// The label for the contact’s co-father-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcofatherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoFatherInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCoFatherInLaw"))
	return rv
}

// The label for the contact’s co-mother-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcomotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoMotherInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCoMotherInLaw"))
	return rv
}

// The label for the contact’s co-parent-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcoparentinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoParentInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCoParentInLaw"))
	return rv
}

// The label for the contact’s co-sibling-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcosiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoSiblingInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCoSiblingInLaw"))
	return rv
}

// The label for the contact’s co-sister-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcosisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationCoSisterInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCoSisterInLaw"))
	return rv
}

// The label for the contact’s colleague.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcolleague
func (c_ CNLabeledValue) CNLabelContactRelationColleague() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationColleague"))
	return rv
}

// The label for the contact’s cousin.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousin
func (c_ CNLabeledValue) CNLabelContactRelationCousin() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousin"))
	return rv
}

// The label for the contact’s father’s brother’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfathersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersBrothersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersBrothersDaughter"))
	return rv
}

// The label for the contact’s father’s brother’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfathersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersBrothersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersBrothersSon"))
	return rv
}

// The label for the contact’s father’s sister’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersSistersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersSistersDaughter"))
	return rv
}

// The label for the contact’s father’s sister’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinFathersSistersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinFathersSistersSon"))
	return rv
}

// The label for the contact’s grandparent’s sibling’s child.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsChild() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsChild"))
	return rv
}

// The label for the contact’s grandparent’s sibling’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsDaughter"))
	return rv
}

// The label for the contact’s grandparent’s sibling’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousingrandparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationCousinGrandparentsSiblingsSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinGrandparentsSiblingsSon"))
	return rv
}

// The label for the contact’s mother’s brother’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmothersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersBrothersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersBrothersDaughter"))
	return rv
}

// The label for the contact’s mother’s brother’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmothersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersBrothersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersBrothersSon"))
	return rv
}

// The label for the contact’s mother’s sister’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmotherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersSistersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersSistersDaughter"))
	return rv
}

// The label for the contact’s mother’s sister’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinmotherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationCousinMothersSistersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinMothersSistersSon"))
	return rv
}

// The label for the contact’s cousin’s or sibling’s child.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinorsiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationCousinOrSiblingsChild() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinOrSiblingsChild"))
	return rv
}

// The label for the contact’s parent’s sibling’s child.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsChild() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsChild"))
	return rv
}

// The label for the contact’s parent’s sibling’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsDaughter"))
	return rv
}

// The label for the contact’s parent’s sibling’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationcousinparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationCousinParentsSiblingsSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationCousinParentsSiblingsSon"))
	return rv
}

// The label for the contact’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughter
func (c_ CNLabeledValue) CNLabelContactRelationDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationDaughter"))
	return rv
}

// The label for the contact’s daughter-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLaw"))
	return rv
}

// The label for the contact’s daughter-in-law or sister-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaworsisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLawOrSisterInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLawOrSisterInLaw"))
	return rv
}

// The label for the contact’s daughter-in-law or stepdaughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationdaughterinlaworstepdaughter
func (c_ CNLabeledValue) CNLabelContactRelationDaughterInLawOrStepdaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationDaughterInLawOrStepdaughter"))
	return rv
}

// The label for the contact’s elder brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationelderbrother
func (c_ CNLabeledValue) CNLabelContactRelationElderBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderBrother"))
	return rv
}

// The label for the contact’s elder brother-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationelderbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationElderBrotherInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderBrotherInLaw"))
	return rv
}

// The label for the contact’s elder cousin.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousin
func (c_ CNLabeledValue) CNLabelContactRelationElderCousin() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousin"))
	return rv
}

// The label for the contact’s father’s brother’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfathersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersBrothersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersBrothersDaughter"))
	return rv
}

// The label for the contact’s father’s brother’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfathersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersBrothersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersBrothersSon"))
	return rv
}

// The label for the contact’s father’s sister’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersSistersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersSistersDaughter"))
	return rv
}

// The label for the contact’s father’s sister’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinFathersSistersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinFathersSistersSon"))
	return rv
}

// The label for the contact’s mother’s brother’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmothersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersBrothersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersBrothersDaughter"))
	return rv
}

// The label for the contact’s mother’s brother’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmothersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersBrothersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersBrothersSon"))
	return rv
}

// The label for the contact’s mother’s sibling’s daughter or father’s sister’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssiblingsdaughterorfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSiblingsDaughterOrFathersSistersDaughter"))
	return rv
}

// The label for the contact’s mother’s sibling’s son or father’s sister’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssiblingssonorfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSiblingsSonOrFathersSistersSon"))
	return rv
}

// The label for the contact’s mother’s sister’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSistersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSistersDaughter"))
	return rv
}

// The label for the contact’s mother’s sister’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinmotherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinMothersSistersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinMothersSistersSon"))
	return rv
}

// The label for the contact’s parent’s sibling’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinParentsSiblingsDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinParentsSiblingsDaughter"))
	return rv
}

// The label for the contact’s parent’s sibling’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldercousinparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationElderCousinParentsSiblingsSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderCousinParentsSiblingsSon"))
	return rv
}

// The label for the contact’s elder sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersibling
func (c_ CNLabeledValue) CNLabelContactRelationElderSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderSibling"))
	return rv
}

// The label for the contact’s elder sibling-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationElderSiblingInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderSiblingInLaw"))
	return rv
}

// The label for the contact’s elder sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersister
func (c_ CNLabeledValue) CNLabelContactRelationElderSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderSister"))
	return rv
}

// The label for the contact’s elder sister-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldersisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationElderSisterInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationElderSisterInLaw"))
	return rv
}

// The label for the contact’s eldest brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldestbrother
func (c_ CNLabeledValue) CNLabelContactRelationEldestBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationEldestBrother"))
	return rv
}

// The label for the contact’s eldest sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationeldestsister
func (c_ CNLabeledValue) CNLabelContactRelationEldestSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationEldestSister"))
	return rv
}

// The label for the contact’s father.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfather
func (c_ CNLabeledValue) CNLabelContactRelationFather() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationFather"))
	return rv
}

// The label for the contact’s father-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLaw"))
	return rv
}

// The label for the contact’s husband’s father.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlawhusbandsfather
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawHusbandsFather() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawHusbandsFather"))
	return rv
}

// The label for the contact’s father-in-law or stepfather.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlaworstepfather
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawOrStepfather() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawOrStepfather"))
	return rv
}

// The label for the contact’s wife’s father.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfatherinlawwifesfather
func (c_ CNLabeledValue) CNLabelContactRelationFatherInLawWifesFather() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationFatherInLawWifesFather"))
	return rv
}

// The label for the contact’s female cousin.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalecousin
func (c_ CNLabeledValue) CNLabelContactRelationFemaleCousin() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationFemaleCousin"))
	return rv
}

// The label for the contact’s female friend.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalefriend
func (c_ CNLabeledValue) CNLabelContactRelationFemaleFriend() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationFemaleFriend"))
	return rv
}

// The label for the contact’s female partner.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfemalepartner
func (c_ CNLabeledValue) CNLabelContactRelationFemalePartner() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationFemalePartner"))
	return rv
}

// The label for the contact’s friend.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationfriend
func (c_ CNLabeledValue) CNLabelContactRelationFriend() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationFriend"))
	return rv
}

// The label for the contact’s girlfriend.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgirlfriend
func (c_ CNLabeledValue) CNLabelContactRelationGirlfriend() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGirlfriend"))
	return rv
}

// The label for the contact’s girlfriend or boyfriend.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgirlfriendorboyfriend
func (c_ CNLabeledValue) CNLabelContactRelationGirlfriendOrBoyfriend() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGirlfriendOrBoyfriend"))
	return rv
}

// The label for the contact’s grandaunt.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandaunt
func (c_ CNLabeledValue) CNLabelContactRelationGrandaunt() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandaunt"))
	return rv
}

// The label for the contact’s grandchild.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandchild
func (c_ CNLabeledValue) CNLabelContactRelationGrandchild() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandchild"))
	return rv
}

// The label for the contact’s grandchild or sibling’s child.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandchildorsiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationGrandchildOrSiblingsChild() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandchildOrSiblingsChild"))
	return rv
}

// The label for the contact’s granddaughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughter"))
	return rv
}

// The label for the contact’s daughter’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughterdaughtersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterDaughtersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterDaughtersDaughter"))
	return rv
}

// The label for the contact’s granddaughter or niece.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughterorniece
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterOrNiece() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterOrNiece"))
	return rv
}

// The label for the contact’s son’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranddaughtersonsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationGranddaughterSonsDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGranddaughterSonsDaughter"))
	return rv
}

// The label for the contact’s grandfather.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfather
func (c_ CNLabeledValue) CNLabelContactRelationGrandfather() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandfather"))
	return rv
}

// The label for the contact’s father’s father.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfatherfathersfather
func (c_ CNLabeledValue) CNLabelContactRelationGrandfatherFathersFather() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandfatherFathersFather"))
	return rv
}

// The label for the contact’s mother’s father.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandfathermothersfather
func (c_ CNLabeledValue) CNLabelContactRelationGrandfatherMothersFather() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandfatherMothersFather"))
	return rv
}

// The label for the contact’s grandmother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmother
func (c_ CNLabeledValue) CNLabelContactRelationGrandmother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandmother"))
	return rv
}

// The label for the contact’s father’s mother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmotherfathersmother
func (c_ CNLabeledValue) CNLabelContactRelationGrandmotherFathersMother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandmotherFathersMother"))
	return rv
}

// The label for the contact’s mother’s mother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandmothermothersmother
func (c_ CNLabeledValue) CNLabelContactRelationGrandmotherMothersMother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandmotherMothersMother"))
	return rv
}

// The label for the contact’s grandnephew.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephew
func (c_ CNLabeledValue) CNLabelContactRelationGrandnephew() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandnephew"))
	return rv
}

// The label for the contact’s brother’s grandson.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephewbrothersgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGrandnephewBrothersGrandson() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandnephewBrothersGrandson"))
	return rv
}

// The label for the contact’s sister’s grandson.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandnephewsistersgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGrandnephewSistersGrandson() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandnephewSistersGrandson"))
	return rv
}

// The label for the contact’s grandniece.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniece
func (c_ CNLabeledValue) CNLabelContactRelationGrandniece() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandniece"))
	return rv
}

// The label for the contact’s brother’s granddaughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniecebrothersgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGrandnieceBrothersGranddaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandnieceBrothersGranddaughter"))
	return rv
}

// The label for the contact’s sister’s granddaughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandniecesistersgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGrandnieceSistersGranddaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandnieceSistersGranddaughter"))
	return rv
}

// The label for the contact’s grandparent.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandparent
func (c_ CNLabeledValue) CNLabelContactRelationGrandparent() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandparent"))
	return rv
}

// The label for the contact’s grandson.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGrandson() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandson"))
	return rv
}

// The label for the contact’s daughter’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsondaughtersson
func (c_ CNLabeledValue) CNLabelContactRelationGrandsonDaughtersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandsonDaughtersSon"))
	return rv
}

// The label for the contact’s grandson or nephew.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsonornephew
func (c_ CNLabeledValue) CNLabelContactRelationGrandsonOrNephew() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandsonOrNephew"))
	return rv
}

// The label for the contact’s son’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgrandsonsonsson
func (c_ CNLabeledValue) CNLabelContactRelationGrandsonSonsSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGrandsonSonsSon"))
	return rv
}

// The label for the contact’s granduncle.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgranduncle
func (c_ CNLabeledValue) CNLabelContactRelationGranduncle() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGranduncle"))
	return rv
}

// The label for the contact’s grandchild.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandchild
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandchild() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandchild"))
	return rv
}

// The label for the contact’s grandchild or sibling’s grandchild.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandchildorsiblingsgrandchild
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandchildOrSiblingsGrandchild"))
	return rv
}

// The label for the contact’s great-granddaughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgranddaughter
func (c_ CNLabeledValue) CNLabelContactRelationGreatGranddaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGreatGranddaughter"))
	return rv
}

// The label for the contact’s great-grandfather.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandfather
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandfather() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandfather"))
	return rv
}

// The label for the contact’s great-grandmother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandmother
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandmother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandmother"))
	return rv
}

// The label for the contact’s great-grandparent.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandparent
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandparent() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandparent"))
	return rv
}

// The label for the contact’s great-grandson.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationgreatgrandson
func (c_ CNLabeledValue) CNLabelContactRelationGreatGrandson() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationGreatGrandson"))
	return rv
}

// The label for the contact’s husband.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationhusband
func (c_ CNLabeledValue) CNLabelContactRelationHusband() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationHusband"))
	return rv
}

// The label for the contact’s male cousin.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalecousin
func (c_ CNLabeledValue) CNLabelContactRelationMaleCousin() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationMaleCousin"))
	return rv
}

// The label for the contact’s male friend.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalefriend
func (c_ CNLabeledValue) CNLabelContactRelationMaleFriend() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationMaleFriend"))
	return rv
}

// The label for the contact’s male partner.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmalepartner
func (c_ CNLabeledValue) CNLabelContactRelationMalePartner() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationMalePartner"))
	return rv
}

// The label for the contact’s manager.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmanager
func (c_ CNLabeledValue) CNLabelContactRelationManager() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationManager"))
	return rv
}

// The label for the contact’s mother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmother
func (c_ CNLabeledValue) CNLabelContactRelationMother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationMother"))
	return rv
}

// The label for the contact’s mother-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLaw"))
	return rv
}

// The label for the contact’s husband’s mother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlawhusbandsmother
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawHusbandsMother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawHusbandsMother"))
	return rv
}

// The label for the contact’s mother-in-law or stepmother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlaworstepmother
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawOrStepmother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawOrStepmother"))
	return rv
}

// The label for the contact’s wife’s mother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationmotherinlawwifesmother
func (c_ CNLabeledValue) CNLabelContactRelationMotherInLawWifesMother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationMotherInLawWifesMother"))
	return rv
}

// The label for the contact’s nephew.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephew
func (c_ CNLabeledValue) CNLabelContactRelationNephew() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNephew"))
	return rv
}

// The label for the contact’s brother’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationNephewBrothersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNephewBrothersSon"))
	return rv
}

// The label for the contact’s brother’s son or husband’s sibling’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewbrotherssonorhusbandssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNephewBrothersSonOrHusbandsSiblingsSon"))
	return rv
}

// The label for the contact’s nephew or cousin.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnepheworcousin
func (c_ CNLabeledValue) CNLabelContactRelationNephewOrCousin() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNephewOrCousin"))
	return rv
}

// The label for the contact’s sister’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewsistersson
func (c_ CNLabeledValue) CNLabelContactRelationNephewSistersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNephewSistersSon"))
	return rv
}

// The label for the contact’s sister’s son or wife’s sibling’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnephewsisterssonorwifessiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNephewSistersSonOrWifesSiblingsSon"))
	return rv
}

// The label for the contact’s niece.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniece
func (c_ CNLabeledValue) CNLabelContactRelationNiece() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNiece"))
	return rv
}

// The label for the contact’s brother’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecebrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceBrothersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNieceBrothersDaughter"))
	return rv
}

// The label for the contact’s brother’s daughter or husband’s sibling’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecebrothersdaughterorhusbandssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNieceBrothersDaughterOrHusbandsSiblingsDaughter"))
	return rv
}

// The label for the contact’s niece or cousin.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationnieceorcousin
func (c_ CNLabeledValue) CNLabelContactRelationNieceOrCousin() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNieceOrCousin"))
	return rv
}

// The label for the contact’s sister’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecesistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceSistersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNieceSistersDaughter"))
	return rv
}

// The label for the contact’s sister’s daughter or wife’s sibling’s daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationniecesistersdaughterorwifessiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationNieceSistersDaughterOrWifesSiblingsDaughter"))
	return rv
}

// The label for the contact’s parent.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparent
func (c_ CNLabeledValue) CNLabelContactRelationParent() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParent"))
	return rv
}

// The label for the contact’s parent-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentinlaw
func (c_ CNLabeledValue) CNLabelContactRelationParentInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParentInLaw"))
	return rv
}

// The label for the contact’s parent’s elder sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentseldersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsElderSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParentsElderSibling"))
	return rv
}

// The label for the contact’s parent’s sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParentsSibling"))
	return rv
}

// The label for the contact’s father’s elder sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfatherseldersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersElderSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersElderSibling"))
	return rv
}

// The label for the contact’s father’s sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfatherssibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersSibling"))
	return rv
}

// The label for the contact’s father’s youngest sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingfathersyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingFathersYoungerSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingFathersYoungerSibling"))
	return rv
}

// The label for the contact’s mother’s elder sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmotherseldersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersElderSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersElderSibling"))
	return rv
}

// The label for the contact’s mother’s sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmotherssibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersSibling"))
	return rv
}

// The label for the contact’s mother’s younger sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentssiblingmothersyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsSiblingMothersYoungerSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParentsSiblingMothersYoungerSibling"))
	return rv
}

// The label for the contact’s parent’s younger sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationparentsyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationParentsYoungerSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationParentsYoungerSibling"))
	return rv
}

// The label for the contact’s partner.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationpartner
func (c_ CNLabeledValue) CNLabelContactRelationPartner() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationPartner"))
	return rv
}

// The label for the contact’s sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsibling
func (c_ CNLabeledValue) CNLabelContactRelationSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSibling"))
	return rv
}

// The label for the contact’s sibling-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationSiblingInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSiblingInLaw"))
	return rv
}

// The label for the contact’s sibling’s child.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsiblingschild
func (c_ CNLabeledValue) CNLabelContactRelationSiblingsChild() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSiblingsChild"))
	return rv
}

// The label for the contact’s sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsister
func (c_ CNLabeledValue) CNLabelContactRelationSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSister"))
	return rv
}

// The label for the contact’s sister-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLaw"))
	return rv
}

// The label for the contact’s brother’s wife.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawBrothersWife() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawBrothersWife"))
	return rv
}

// The label for the contact’s elder brother’s wife.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawelderbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawElderBrothersWife() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawElderBrothersWife"))
	return rv
}

// The label for the contact’s husband’s brother’s wife.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawhusbandsbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawHusbandsBrothersWife() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawHusbandsBrothersWife"))
	return rv
}

// The label for the contact’s husband’s sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawhusbandssister
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawHusbandsSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawHusbandsSister"))
	return rv
}

// The label for the contact’s spouse’s sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawspousessister
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawSpousesSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawSpousesSister"))
	return rv
}

// The label for the contact’s wife’s brother’s wife.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawwifesbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawWifesBrothersWife() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawWifesBrothersWife"))
	return rv
}

// The label for the contact’s wife’s sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawwifessister
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawWifesSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawWifesSister"))
	return rv
}

// The label for the contact’s younger brother’s wife.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsisterinlawyoungerbrotherswife
func (c_ CNLabeledValue) CNLabelContactRelationSisterInLawYoungerBrothersWife() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSisterInLawYoungerBrothersWife"))
	return rv
}

// The label for the contact’s son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationson
func (c_ CNLabeledValue) CNLabelContactRelationSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSon"))
	return rv
}

// The label for the contact’s son-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaw
func (c_ CNLabeledValue) CNLabelContactRelationSonInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSonInLaw"))
	return rv
}

// The label for the contact’s son-in-law or brother-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaworbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationSonInLawOrBrotherInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSonInLawOrBrotherInLaw"))
	return rv
}

// The label for the contact’s son-in-law or stepson.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationsoninlaworstepson
func (c_ CNLabeledValue) CNLabelContactRelationSonInLawOrStepson() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSonInLawOrStepson"))
	return rv
}

// The label for the contact’s spouse.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationspouse
func (c_ CNLabeledValue) CNLabelContactRelationSpouse() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationSpouse"))
	return rv
}

// The label for the contact’s stepbrother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepbrother
func (c_ CNLabeledValue) CNLabelContactRelationStepbrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationStepbrother"))
	return rv
}

// The label for the contact’s stepchild.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepchild
func (c_ CNLabeledValue) CNLabelContactRelationStepchild() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationStepchild"))
	return rv
}

// The label for the contact’s stepdaughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepdaughter
func (c_ CNLabeledValue) CNLabelContactRelationStepdaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationStepdaughter"))
	return rv
}

// The label for the contact’s stepfather.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepfather
func (c_ CNLabeledValue) CNLabelContactRelationStepfather() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationStepfather"))
	return rv
}

// The label for the contact’s stepmother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepmother
func (c_ CNLabeledValue) CNLabelContactRelationStepmother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationStepmother"))
	return rv
}

// The label for the contact’s stepparent.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepparent
func (c_ CNLabeledValue) CNLabelContactRelationStepparent() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationStepparent"))
	return rv
}

// The label for the contact’s stepsister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepsister
func (c_ CNLabeledValue) CNLabelContactRelationStepsister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationStepsister"))
	return rv
}

// The label for the contact’s stepson.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationstepson
func (c_ CNLabeledValue) CNLabelContactRelationStepson() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationStepson"))
	return rv
}

// The label for the contact’s teacher.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationteacher
func (c_ CNLabeledValue) CNLabelContactRelationTeacher() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationTeacher"))
	return rv
}

// The label for the contact’s uncle.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncle
func (c_ CNLabeledValue) CNLabelContactRelationUncle() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncle"))
	return rv
}

// The label for the contact’s father’s brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersBrother"))
	return rv
}

// The label for the contact’s father’s elder brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherselderbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersElderBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersElderBrother"))
	return rv
}

// The label for the contact’s elder sister’s husband.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherseldersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersElderSistersHusband() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersElderSistersHusband"))
	return rv
}

// The label for the contact’s father’s sister’s husband.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefatherssistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersSistersHusband() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersSistersHusband"))
	return rv
}

// The label for the contact’s father’s younger brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersYoungerBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersYoungerBrother"))
	return rv
}

// The label for the contact’s father’s younger sister’s husband.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclefathersyoungersistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleFathersYoungerSistersHusband() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleFathersYoungerSistersHusband"))
	return rv
}

// The label for the contact’s mother’s brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemothersbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersBrother"))
	return rv
}

// The label for the contact’s mother’s elder brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemotherselderbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersElderBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersElderBrother"))
	return rv
}

// The label for the contact’s mother’s sister’s husband.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemotherssistershusband
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersSistersHusband() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersSistersHusband"))
	return rv
}

// The label for the contact’s mother’s younger brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationunclemothersyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleMothersYoungerBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleMothersYoungerBrother"))
	return rv
}

// The label for the contact’s parent’s brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentsbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsBrother"))
	return rv
}

// The label for the contact’s parent’s elder brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentselderbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsElderBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsElderBrother"))
	return rv
}

// The label for the contact’s parent’s younger brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationuncleparentsyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationUncleParentsYoungerBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationUncleParentsYoungerBrother"))
	return rv
}

// The label for the contact’s wife.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationwife
func (c_ CNLabeledValue) CNLabelContactRelationWife() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationWife"))
	return rv
}

// The label for the contact’s younger brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungerbrother
func (c_ CNLabeledValue) CNLabelContactRelationYoungerBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerBrother"))
	return rv
}

// The label for the contact’s younger brother-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungerbrotherinlaw
func (c_ CNLabeledValue) CNLabelContactRelationYoungerBrotherInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerBrotherInLaw"))
	return rv
}

// The label for the contact’s younger cousin.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousin
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousin() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousin"))
	return rv
}

// The label for the contact’s father’s brother’s younger daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfathersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersBrothersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersBrothersDaughter"))
	return rv
}

// The label for the contact’s father’s brother’s younger son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfathersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersBrothersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersBrothersSon"))
	return rv
}

// The label for the contact’s father’s sister’s younger daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersSistersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersSistersDaughter"))
	return rv
}

// The label for the contact’s father’s sister’s younger son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinFathersSistersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinFathersSistersSon"))
	return rv
}

// The label for the contact’s mother’s brother’s younger daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmothersbrothersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersBrothersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersBrothersDaughter"))
	return rv
}

// The label for the contact’s mother’s brother’s younger son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmothersbrothersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersBrothersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersBrothersSon"))
	return rv
}

// The label for the contact’s mother’s sibling’s younger daughter or father’s sister’s younger daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssiblingsdaughterorfatherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSiblingsDaughterOrFathersSistersDaughter"))
	return rv
}

// The label for the contact’s mother’s sibling’s younger son or father’s sister’s younger son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssiblingssonorfatherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSiblingsSonOrFathersSistersSon"))
	return rv
}

// The label for the contact’s mother’s sister’s younger daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssistersdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSistersDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSistersDaughter"))
	return rv
}

// The label for the contact’s mother’s sister’s younger son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinmotherssistersson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinMothersSistersSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinMothersSistersSon"))
	return rv
}

// The label for the contact’s parent’s sibling’s younger daughter.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinparentssiblingsdaughter
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinParentsSiblingsDaughter() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinParentsSiblingsDaughter"))
	return rv
}

// The label for the contact’s parent’s sibling’s younger son.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungercousinparentssiblingsson
func (c_ CNLabeledValue) CNLabelContactRelationYoungerCousinParentsSiblingsSon() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerCousinParentsSiblingsSon"))
	return rv
}

// The label for the contact’s younger sibling.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersibling
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSibling() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSibling"))
	return rv
}

// The label for the contact’s younger sibling-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersiblinginlaw
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSiblingInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSiblingInLaw"))
	return rv
}

// The label for the contact’s younger sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersister
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSister"))
	return rv
}

// The label for the contact’s younger sister-in-law.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungersisterinlaw
func (c_ CNLabeledValue) CNLabelContactRelationYoungerSisterInLaw() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungerSisterInLaw"))
	return rv
}

// The label for the contact’s youngest brother.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungestbrother
func (c_ CNLabeledValue) CNLabelContactRelationYoungestBrother() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungestBrother"))
	return rv
}

// The label for the contact’s youngest sister.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelcontactrelationyoungestsister
func (c_ CNLabeledValue) CNLabelContactRelationYoungestSister() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelContactRelationYoungestSister"))
	return rv
}

// The label for identifying the contact’s anniversary date.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabeldateanniversary
func (c_ CNLabeledValue) CNLabelDateAnniversary() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelDateAnniversary"))
	return rv
}

// The label for identifying the contact’s iCloud email information.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelemailicloud
func (c_ CNLabeledValue) CNLabelEmailiCloud() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelEmailiCloud"))
	return rv
}

// The label for identifying home information.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelhome
func (c_ CNLabeledValue) CNLabelHome() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelHome"))
	return rv
}

// The label for identifying other information.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelother
func (c_ CNLabeledValue) CNLabelOther() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelOther"))
	return rv
}

// The label for identifying the contact’s Apple Watch phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberapplewatch
func (c_ CNLabeledValue) CNLabelPhoneNumberAppleWatch() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelPhoneNumberAppleWatch"))
	return rv
}

// The label for identifying the contact’s home fax number.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberhomefax
func (c_ CNLabeledValue) CNLabelPhoneNumberHomeFax() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelPhoneNumberHomeFax"))
	return rv
}

// The label for identifying the contact’s main phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumbermain
func (c_ CNLabeledValue) CNLabelPhoneNumberMain() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelPhoneNumberMain"))
	return rv
}

// The label for identifying the contact’s mobile phone number.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumbermobile
func (c_ CNLabeledValue) CNLabelPhoneNumberMobile() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelPhoneNumberMobile"))
	return rv
}

// The label for identifying another fax number.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberotherfax
func (c_ CNLabeledValue) CNLabelPhoneNumberOtherFax() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelPhoneNumberOtherFax"))
	return rv
}

// The label for identifying the contact’s pager number.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberpager
func (c_ CNLabeledValue) CNLabelPhoneNumberPager() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelPhoneNumberPager"))
	return rv
}

// The label for identifying the contact’s work fax number.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberworkfax
func (c_ CNLabeledValue) CNLabelPhoneNumberWorkFax() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelPhoneNumberWorkFax"))
	return rv
}

// The label for identifying the contact’s iPhone number.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelphonenumberiphone
func (c_ CNLabeledValue) CNLabelPhoneNumberiPhone() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelPhoneNumberiPhone"))
	return rv
}

// The label for the contact’s school.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelschool
func (c_ CNLabeledValue) CNLabelSchool() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelSchool"))
	return rv
}

// The label for identifying URL information.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelurladdresshomepage
func (c_ CNLabeledValue) CNLabelURLAddressHomePage() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelURLAddressHomePage"))
	return rv
}

// The label for identifying work information.
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnlabelwork
func (c_ CNLabeledValue) CNLabelWork() appkit.string {
	rv := objc.Send[appkit.string](c_.ID, objc.Sel("CNLabelWork"))
	return rv
}


