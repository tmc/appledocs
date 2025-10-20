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
	LabeledValueBySettingLabel(label string) unsafe.Pointer
	LabeledValueBySettingLabelValue(label string, value unsafe.Pointer) unsafe.Pointer
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
func NewCNLabeledValueWithLabelValue(label string, value unsafe.Pointer) CNLabeledValue {
	instance := getCNLabeledValueClass().Alloc()
	rv := objc.Send[CNLabeledValue](instance.ID, objc.Sel("initWithLabel:value:"), objc.String(label), value)
	rv.Autorelease()
	return rv
}


// Returns a new labeled value identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/labeledValueWithLabel:value:
func (cc _CNLabeledValueClass) LabeledValueWithLabelValue(label string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("labeledValueWithLabel:value:"), objc.String(label), value)
	return rv
}

// Returns a localized string for the specified label.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/localizedString(forLabel:)
func (cc _CNLabeledValueClass) LocalizedStringForLabel(label string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("localizedStringForLabel:"), objc.String(label))
	return rv
}

// Returns a labeled value object with an existing value and identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingLabel(_:)
func (c_ CNLabeledValue) LabeledValueBySettingLabel(label string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("labeledValueBySettingLabel:"), objc.String(label))
	return rv
}

// Returns a labeled value object with the specified label and value with the existing identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/settingLabel(_:value:)
func (c_ CNLabeledValue) LabeledValueBySettingLabelValue(label string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("labeledValueBySettingLabel:value:"), objc.String(label), value)
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
func (c_ CNLabeledValue) Identifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("identifier"))
	return rv
}

// The label for a contact property value.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/label
func (c_ CNLabeledValue) Label() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("label"))
	return rv
}

// A contact property value.
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNLabeledValue/value
func (c_ CNLabeledValue) Value() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("value"))
	return rv
}


