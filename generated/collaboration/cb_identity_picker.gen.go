// Code generated from Apple documentation for Collaboration. DO NOT EDIT.

package collaboration

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CBIdentityPicker] class.
var (
	CBIdentityPickerClass     _CBIdentityPickerClass
	CBIdentityPickerClassOnce sync.Once
)

func getCBIdentityPickerClass() _CBIdentityPickerClass {
	CBIdentityPickerClassOnce.Do(func() {
		CBIdentityPickerClass = _CBIdentityPickerClass{objc.GetClass("CBIdentityPicker")}
	})
	return CBIdentityPickerClass
}

type _CBIdentityPickerClass struct {
	class objc.Class
}

// An interface definition for the [CBIdentityPicker] class.
type ICBIdentityPicker interface {
	objectivec.IObject
	RunModal() int
	RunModalForWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer)
	RunModalForWindowModalDelegateDidEndSelectorContextInfo(window unsafe.Pointer, delegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer)
}

// A object allows a user to select identities—for example, user or group objects—that it wants one or more services or shared resources to have access to. An identity picker can be displayed either as an application-modal dialog or as a sheet attached to a document window. An identity picker returns the selected records to be added to access control lists using Collaboration. If a selected record is not a user or group identity, then an identity picker prompts the user for additional information—such as a password—to promote that record to a sharing account.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker
type CBIdentityPicker struct {
	objectivec.Object
}

// CBIdentityPickerFrom constructs a [CBIdentityPicker] from an unsafe.Pointer.
//
// A object allows a user to select identities—for example, user or group objects—that it wants one or more services or shared resources to have access to. An identity picker can be displayed either as an application-modal dialog or as a sheet attached to a document window. An identity picker returns the selected records to be added to access control lists using Collaboration. If a selected record is not a user or group identity, then an identity picker prompts the user for additional information—such as a password—to promote that record to a sharing account.
func CBIdentityPickerFrom(ptr unsafe.Pointer) CBIdentityPicker {
	return CBIdentityPicker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CBIdentityPickerClass) Alloc() CBIdentityPicker {
	rv := objc.Send[CBIdentityPicker](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CBIdentityPickerClass) New() CBIdentityPicker {
	rv := objc.Send[CBIdentityPicker](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBIdentityPicker) Init() CBIdentityPicker {
	rv := objc.Send[CBIdentityPicker](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBIdentityPicker) Autorelease() CBIdentityPicker {
	rv := objc.Send[CBIdentityPicker](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBIdentityPicker creates a new CBIdentityPicker instance.
func NewCBIdentityPicker() CBIdentityPicker {
	return getCBIdentityPickerClass().New()
}


// Runs the receiver as an application-modal dialog.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/runModal()
func (c_ CBIdentityPicker) RunModal() int {
	rv := objc.Send[int](c_.ID, objc.Sel("runModal"))
	return rv
}

// Runs the identity picker modally as a sheet attached to a specified window.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/runModal(for:completionHandler:)
func (c_ CBIdentityPicker) RunModalForWindowCompletionHandler(window unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("runModalForWindow:completionHandler:"), window, completionHandler)
}

// Runs the receiver modally as a sheet attached to a specified window.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/runModal(for:modalDelegate:didEnd:contextInfo:)
func (c_ CBIdentityPicker) RunModalForWindowModalDelegateDidEndSelectorContextInfo(window unsafe.Pointer, delegate objc.ID, didEndSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("runModalForWindow:modalDelegate:didEndSelector:contextInfo:"), window, delegate, didEndSelector, contextInfo)
}

// A Boolean value indicating whether the user is allowed to select multiple identities.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/allowsMultipleSelection
func (c_ CBIdentityPicker) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}


// SetAllowsMultipleSelection sets the value of the allowsMultipleSelection property.
// A Boolean value indicating whether the user is allowed to select multiple identities.

//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/allowsMultipleSelection
func (c_ CBIdentityPicker) SetAllowsMultipleSelection(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}

// The array of identities (represented by objects) selected using the identity picker.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/identities
func (c_ CBIdentityPicker) Identities() []CBIdentity {
	rv := objc.Send[[]CBIdentity](c_.ID, objc.Sel("identities"))
	return rv
}

// The title of the identity picker.
//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/title
func (c_ CBIdentityPicker) Title() string {
	rv := objc.Send[string](c_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The title of the identity picker.

//
// [Full Topic]: https://developer.apple.com/documentation/Collaboration/CBIdentityPicker/title
func (c_ CBIdentityPicker) SetTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), objc.String(value))
}



