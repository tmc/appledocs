// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/appkit"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/storekit"
)

// PPaymentAuthorizationControllerDelegate is the PKPaymentAuthorizationControllerDelegate protocol interface.
//
// Methods that let you respond to user interactions with your payment authorization controller.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.passkit/documentation/PassKit/PKPaymentAuthorizationControllerDelegate
type PPaymentAuthorizationControllerDelegate interface {
	// Required methods
	PaymentAuthorizationControllerDidFinish(controller PaymentAuthorizationController /* not a class type */)/* debug [protocol_interface/required_method]: PaymentAuthorizationControllerDidFinish */
	// Optional methods
	PaymentAuthorizationControllerDidAuthorizePaymentCompletion(controller PaymentAuthorizationController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)
	HasPaymentAuthorizationControllerDidAuthorizePaymentCompletion() bool
	PaymentAuthorizationControllerDidAuthorizePaymentHandler(controller PaymentAuthorizationController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)
	HasPaymentAuthorizationControllerDidAuthorizePaymentHandler() bool
	PaymentAuthorizationControllerDidChangeCouponCodeHandler(controller PaymentAuthorizationController /* not a class type */, couponCode objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer)
	HasPaymentAuthorizationControllerDidChangeCouponCodeHandler() bool
	PaymentAuthorizationControllerDidRequestMerchantSessionUpdate(controller PaymentAuthorizationController /* not a class type */, handler unsafe.Pointer)
	HasPaymentAuthorizationControllerDidRequestMerchantSessionUpdate() bool
	PaymentAuthorizationControllerDidSelectPaymentMethodCompletion(controller PaymentAuthorizationController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationControllerDidSelectPaymentMethodCompletion() bool
	PaymentAuthorizationControllerDidSelectPaymentMethodHandler(controller PaymentAuthorizationController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationControllerDidSelectPaymentMethodHandler() bool
	PaymentAuthorizationControllerDidSelectShippingContactCompletion(controller PaymentAuthorizationController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationControllerDidSelectShippingContactCompletion() bool
	PaymentAuthorizationControllerDidSelectShippingContactHandler(controller PaymentAuthorizationController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationControllerDidSelectShippingContactHandler() bool
	PaymentAuthorizationControllerDidSelectShippingMethodCompletion(controller PaymentAuthorizationController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationControllerDidSelectShippingMethodCompletion() bool
	PaymentAuthorizationControllerDidSelectShippingMethodHandler(controller PaymentAuthorizationController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationControllerDidSelectShippingMethodHandler() bool
	PaymentAuthorizationControllerWillAuthorizePayment(controller PaymentAuthorizationController /* not a class type */)
	HasPaymentAuthorizationControllerWillAuthorizePayment() bool
	PresentationWindowForPaymentAuthorizationController(controller PaymentAuthorizationController /* not a class type */) appkit.Window
	HasPresentationWindowForPaymentAuthorizationController() bool
}

// PaymentAuthorizationControllerDelegate is a delegate implementation builder for the PPaymentAuthorizationControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PaymentAuthorizationControllerDelegate struct {
	_PaymentAuthorizationControllerDidAuthorizePaymentCompletion func(controller PaymentAuthorizationController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)
	_PaymentAuthorizationControllerDidAuthorizePaymentHandler func(controller PaymentAuthorizationController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)
	_PaymentAuthorizationControllerDidChangeCouponCodeHandler func(controller PaymentAuthorizationController /* not a class type */, couponCode objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer)
	_PaymentAuthorizationControllerDidRequestMerchantSessionUpdate func(controller PaymentAuthorizationController /* not a class type */, handler unsafe.Pointer)
	_PaymentAuthorizationControllerDidSelectPaymentMethodCompletion func(controller PaymentAuthorizationController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationControllerDidSelectPaymentMethodHandler func(controller PaymentAuthorizationController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationControllerDidSelectShippingContactCompletion func(controller PaymentAuthorizationController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationControllerDidSelectShippingContactHandler func(controller PaymentAuthorizationController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationControllerDidSelectShippingMethodCompletion func(controller PaymentAuthorizationController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationControllerDidSelectShippingMethodHandler func(controller PaymentAuthorizationController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationControllerWillAuthorizePayment func(controller PaymentAuthorizationController /* not a class type */)
	_PresentationWindowForPaymentAuthorizationController func(controller PaymentAuthorizationController /* not a class type */) appkit.Window
	_PaymentAuthorizationControllerDidFinish func(controller PaymentAuthorizationController /* not a class type */)
}

// SetPaymentAuthorizationControllerDidAuthorizePaymentCompletion sets the handler for the PaymentAuthorizationControllerDidAuthorizePaymentCompletion delegate method.
//
// Tells the delegate that the user authorized the payment request, and asks for a result.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidAuthorizePaymentCompletion(f func(controller PaymentAuthorizationController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)) {
	d._PaymentAuthorizationControllerDidAuthorizePaymentCompletion = f
}

// SetPaymentAuthorizationControllerDidAuthorizePaymentHandler sets the handler for the PaymentAuthorizationControllerDidAuthorizePaymentHandler delegate method.
//
// Tells the delegate that the user authorized the payment request, and asks for a result.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidAuthorizePaymentHandler(f func(controller PaymentAuthorizationController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)) {
	d._PaymentAuthorizationControllerDidAuthorizePaymentHandler = f
}

// SetPaymentAuthorizationControllerDidChangeCouponCodeHandler sets the handler for the PaymentAuthorizationControllerDidChangeCouponCodeHandler delegate method.
//
// Tells the delegate that the user entered or updated a coupon code.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidChangeCouponCodeHandler(f func(controller PaymentAuthorizationController /* not a class type */, couponCode objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationControllerDidChangeCouponCodeHandler = f
}

// SetPaymentAuthorizationControllerDidRequestMerchantSessionUpdate sets the handler for the PaymentAuthorizationControllerDidRequestMerchantSessionUpdate delegate method.
//
// Requests an object that validates the identity of a merchant for a payment request.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidRequestMerchantSessionUpdate(f func(controller PaymentAuthorizationController /* not a class type */, handler unsafe.Pointer)) {
	d._PaymentAuthorizationControllerDidRequestMerchantSessionUpdate = f
}

// SetPaymentAuthorizationControllerDidSelectPaymentMethodCompletion sets the handler for the PaymentAuthorizationControllerDidSelectPaymentMethodCompletion delegate method.
//
// Tells the delegate that the user changed the payment method, and asks for an updated payment request.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidSelectPaymentMethodCompletion(f func(controller PaymentAuthorizationController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationControllerDidSelectPaymentMethodCompletion = f
}

// SetPaymentAuthorizationControllerDidSelectPaymentMethodHandler sets the handler for the PaymentAuthorizationControllerDidSelectPaymentMethodHandler delegate method.
//
// Tells the delegate that the user changed the payment method, and asks for an updated payment request.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidSelectPaymentMethodHandler(f func(controller PaymentAuthorizationController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationControllerDidSelectPaymentMethodHandler = f
}

// SetPaymentAuthorizationControllerDidSelectShippingContactCompletion sets the handler for the PaymentAuthorizationControllerDidSelectShippingContactCompletion delegate method.
//
// Tells the delegate that the user selected a shipping address.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidSelectShippingContactCompletion(f func(controller PaymentAuthorizationController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationControllerDidSelectShippingContactCompletion = f
}

// SetPaymentAuthorizationControllerDidSelectShippingContactHandler sets the handler for the PaymentAuthorizationControllerDidSelectShippingContactHandler delegate method.
//
// Tells the delegate that the user selected a shipping address.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidSelectShippingContactHandler(f func(controller PaymentAuthorizationController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationControllerDidSelectShippingContactHandler = f
}

// SetPaymentAuthorizationControllerDidSelectShippingMethodCompletion sets the handler for the PaymentAuthorizationControllerDidSelectShippingMethodCompletion delegate method.
//
// Tells the delegate that the user selected a shipping method.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidSelectShippingMethodCompletion(f func(controller PaymentAuthorizationController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationControllerDidSelectShippingMethodCompletion = f
}

// SetPaymentAuthorizationControllerDidSelectShippingMethodHandler sets the handler for the PaymentAuthorizationControllerDidSelectShippingMethodHandler delegate method.
//
// Tells the delegate that the user selected a shipping method.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidSelectShippingMethodHandler(f func(controller PaymentAuthorizationController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationControllerDidSelectShippingMethodHandler = f
}

// SetPaymentAuthorizationControllerWillAuthorizePayment sets the handler for the PaymentAuthorizationControllerWillAuthorizePayment delegate method.
//
// Tells the delegate that the user is authorizing the payment request.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerWillAuthorizePayment(f func(controller PaymentAuthorizationController /* not a class type */)) {
	d._PaymentAuthorizationControllerWillAuthorizePayment = f
}

// SetPresentationWindowForPaymentAuthorizationController sets the handler for the PresentationWindowForPaymentAuthorizationController delegate method.
//
// Returns the window in which to present a payment authorization sheet.
func (d *PaymentAuthorizationControllerDelegate) SetPresentationWindowForPaymentAuthorizationController(f func(controller PaymentAuthorizationController /* not a class type */) appkit.Window) {
	d._PresentationWindowForPaymentAuthorizationController = f
}

// SetPaymentAuthorizationControllerDidFinish sets the handler for the PaymentAuthorizationControllerDidFinish delegate method.
//
// Tells the delegate that payment authorization has completed.
func (d *PaymentAuthorizationControllerDelegate) SetPaymentAuthorizationControllerDidFinish(f func(controller PaymentAuthorizationController /* not a class type */)) {
	d._PaymentAuthorizationControllerDidFinish = f
}

// PaymentAuthorizationControllerDidAuthorizePaymentCompletion implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidAuthorizePaymentCompletion(controller PaymentAuthorizationController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer) {
	if d._PaymentAuthorizationControllerDidAuthorizePaymentCompletion != nil {
		d._PaymentAuthorizationControllerDidAuthorizePaymentCompletion(controller, payment, completion)
	}
}

// HasPaymentAuthorizationControllerDidAuthorizePaymentCompletion returns true if a handler for PaymentAuthorizationControllerDidAuthorizePaymentCompletion has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidAuthorizePaymentCompletion() bool {
	return d._PaymentAuthorizationControllerDidAuthorizePaymentCompletion != nil
}

// PaymentAuthorizationControllerDidAuthorizePaymentHandler implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidAuthorizePaymentHandler(controller PaymentAuthorizationController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer) {
	if d._PaymentAuthorizationControllerDidAuthorizePaymentHandler != nil {
		d._PaymentAuthorizationControllerDidAuthorizePaymentHandler(controller, payment, completion)
	}
}

// HasPaymentAuthorizationControllerDidAuthorizePaymentHandler returns true if a handler for PaymentAuthorizationControllerDidAuthorizePaymentHandler has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidAuthorizePaymentHandler() bool {
	return d._PaymentAuthorizationControllerDidAuthorizePaymentHandler != nil
}

// PaymentAuthorizationControllerDidChangeCouponCodeHandler implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidChangeCouponCodeHandler(controller PaymentAuthorizationController /* not a class type */, couponCode objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationControllerDidChangeCouponCodeHandler != nil {
		d._PaymentAuthorizationControllerDidChangeCouponCodeHandler(controller, couponCode, completion)
	}
}

// HasPaymentAuthorizationControllerDidChangeCouponCodeHandler returns true if a handler for PaymentAuthorizationControllerDidChangeCouponCodeHandler has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidChangeCouponCodeHandler() bool {
	return d._PaymentAuthorizationControllerDidChangeCouponCodeHandler != nil
}

// PaymentAuthorizationControllerDidRequestMerchantSessionUpdate implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidRequestMerchantSessionUpdate(controller PaymentAuthorizationController /* not a class type */, handler unsafe.Pointer) {
	if d._PaymentAuthorizationControllerDidRequestMerchantSessionUpdate != nil {
		d._PaymentAuthorizationControllerDidRequestMerchantSessionUpdate(controller, handler)
	}
}

// HasPaymentAuthorizationControllerDidRequestMerchantSessionUpdate returns true if a handler for PaymentAuthorizationControllerDidRequestMerchantSessionUpdate has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidRequestMerchantSessionUpdate() bool {
	return d._PaymentAuthorizationControllerDidRequestMerchantSessionUpdate != nil
}

// PaymentAuthorizationControllerDidSelectPaymentMethodCompletion implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidSelectPaymentMethodCompletion(controller PaymentAuthorizationController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationControllerDidSelectPaymentMethodCompletion != nil {
		d._PaymentAuthorizationControllerDidSelectPaymentMethodCompletion(controller, paymentMethod, completion)
	}
}

// HasPaymentAuthorizationControllerDidSelectPaymentMethodCompletion returns true if a handler for PaymentAuthorizationControllerDidSelectPaymentMethodCompletion has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidSelectPaymentMethodCompletion() bool {
	return d._PaymentAuthorizationControllerDidSelectPaymentMethodCompletion != nil
}

// PaymentAuthorizationControllerDidSelectPaymentMethodHandler implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidSelectPaymentMethodHandler(controller PaymentAuthorizationController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationControllerDidSelectPaymentMethodHandler != nil {
		d._PaymentAuthorizationControllerDidSelectPaymentMethodHandler(controller, paymentMethod, completion)
	}
}

// HasPaymentAuthorizationControllerDidSelectPaymentMethodHandler returns true if a handler for PaymentAuthorizationControllerDidSelectPaymentMethodHandler has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidSelectPaymentMethodHandler() bool {
	return d._PaymentAuthorizationControllerDidSelectPaymentMethodHandler != nil
}

// PaymentAuthorizationControllerDidSelectShippingContactCompletion implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidSelectShippingContactCompletion(controller PaymentAuthorizationController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationControllerDidSelectShippingContactCompletion != nil {
		d._PaymentAuthorizationControllerDidSelectShippingContactCompletion(controller, contact, completion)
	}
}

// HasPaymentAuthorizationControllerDidSelectShippingContactCompletion returns true if a handler for PaymentAuthorizationControllerDidSelectShippingContactCompletion has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidSelectShippingContactCompletion() bool {
	return d._PaymentAuthorizationControllerDidSelectShippingContactCompletion != nil
}

// PaymentAuthorizationControllerDidSelectShippingContactHandler implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidSelectShippingContactHandler(controller PaymentAuthorizationController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationControllerDidSelectShippingContactHandler != nil {
		d._PaymentAuthorizationControllerDidSelectShippingContactHandler(controller, contact, completion)
	}
}

// HasPaymentAuthorizationControllerDidSelectShippingContactHandler returns true if a handler for PaymentAuthorizationControllerDidSelectShippingContactHandler has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidSelectShippingContactHandler() bool {
	return d._PaymentAuthorizationControllerDidSelectShippingContactHandler != nil
}

// PaymentAuthorizationControllerDidSelectShippingMethodCompletion implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidSelectShippingMethodCompletion(controller PaymentAuthorizationController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationControllerDidSelectShippingMethodCompletion != nil {
		d._PaymentAuthorizationControllerDidSelectShippingMethodCompletion(controller, shippingMethod, completion)
	}
}

// HasPaymentAuthorizationControllerDidSelectShippingMethodCompletion returns true if a handler for PaymentAuthorizationControllerDidSelectShippingMethodCompletion has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidSelectShippingMethodCompletion() bool {
	return d._PaymentAuthorizationControllerDidSelectShippingMethodCompletion != nil
}

// PaymentAuthorizationControllerDidSelectShippingMethodHandler implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidSelectShippingMethodHandler(controller PaymentAuthorizationController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationControllerDidSelectShippingMethodHandler != nil {
		d._PaymentAuthorizationControllerDidSelectShippingMethodHandler(controller, shippingMethod, completion)
	}
}

// HasPaymentAuthorizationControllerDidSelectShippingMethodHandler returns true if a handler for PaymentAuthorizationControllerDidSelectShippingMethodHandler has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidSelectShippingMethodHandler() bool {
	return d._PaymentAuthorizationControllerDidSelectShippingMethodHandler != nil
}

// PaymentAuthorizationControllerWillAuthorizePayment implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerWillAuthorizePayment(controller PaymentAuthorizationController /* not a class type */) {
	if d._PaymentAuthorizationControllerWillAuthorizePayment != nil {
		d._PaymentAuthorizationControllerWillAuthorizePayment(controller)
	}
}

// HasPaymentAuthorizationControllerWillAuthorizePayment returns true if a handler for PaymentAuthorizationControllerWillAuthorizePayment has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerWillAuthorizePayment() bool {
	return d._PaymentAuthorizationControllerWillAuthorizePayment != nil
}

// PresentationWindowForPaymentAuthorizationController implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PresentationWindowForPaymentAuthorizationController(controller PaymentAuthorizationController /* not a class type */) appkit.Window {
	if d._PresentationWindowForPaymentAuthorizationController != nil {
		return d._PresentationWindowForPaymentAuthorizationController(controller)
	}
	var zero appkit.Window
	return zero
}

// HasPresentationWindowForPaymentAuthorizationController returns true if a handler for PresentationWindowForPaymentAuthorizationController has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPresentationWindowForPaymentAuthorizationController() bool {
	return d._PresentationWindowForPaymentAuthorizationController != nil
}

// PaymentAuthorizationControllerDidFinish implements the PPaymentAuthorizationControllerDelegate interface.
func (d *PaymentAuthorizationControllerDelegate) PaymentAuthorizationControllerDidFinish(controller PaymentAuthorizationController /* not a class type */) {
	if d._PaymentAuthorizationControllerDidFinish != nil {
		d._PaymentAuthorizationControllerDidFinish(controller)
	}
}

// HasPaymentAuthorizationControllerDidFinish returns true if a handler for PaymentAuthorizationControllerDidFinish has been set.
func (d *PaymentAuthorizationControllerDelegate) HasPaymentAuthorizationControllerDidFinish() bool {
	return d._PaymentAuthorizationControllerDidFinish != nil
}
