// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/storekit"
)

// PPaymentAuthorizationViewControllerDelegate is the PKPaymentAuthorizationViewControllerDelegate protocol interface.
//
// Methods that let you respond to user interactions with your payment authorization view controller.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//
// See: doc://com.apple.passkit/documentation/PassKit/PKPaymentAuthorizationViewControllerDelegate
type PPaymentAuthorizationViewControllerDelegate interface {
	// Required methods
	PaymentAuthorizationViewControllerDidFinish(controller PaymentAuthorizationViewController /* not a class type */)/* debug [protocol_interface/required_method]: PaymentAuthorizationViewControllerDidFinish */
	// Optional methods
	PaymentAuthorizationViewControllerDidAuthorizePaymentCompletion(controller PaymentAuthorizationViewController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidAuthorizePaymentCompletion() bool
	PaymentAuthorizationViewControllerDidAuthorizePaymentHandler(controller PaymentAuthorizationViewController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidAuthorizePaymentHandler() bool
	PaymentAuthorizationViewControllerDidChangeCouponCodeHandler(controller PaymentAuthorizationViewController /* not a class type */, couponCode objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidChangeCouponCodeHandler() bool
	PaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate(controller PaymentAuthorizationViewController /* not a class type */, handler unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate() bool
	PaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion(controller PaymentAuthorizationViewController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion() bool
	PaymentAuthorizationViewControllerDidSelectShippingMethodCompletion(controller PaymentAuthorizationViewController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidSelectShippingMethodCompletion() bool
	PaymentAuthorizationViewControllerDidSelectPaymentMethodHandler(controller PaymentAuthorizationViewController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidSelectPaymentMethodHandler() bool
	PaymentAuthorizationViewControllerDidSelectShippingMethodHandler(controller PaymentAuthorizationViewController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidSelectShippingMethodHandler() bool
	PaymentAuthorizationViewControllerDidSelectShippingAddressCompletion(controller PaymentAuthorizationViewController /* not a class type */, address unsafe.Pointer, completion unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidSelectShippingAddressCompletion() bool
	PaymentAuthorizationViewControllerDidSelectShippingContactCompletion(controller PaymentAuthorizationViewController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidSelectShippingContactCompletion() bool
	PaymentAuthorizationViewControllerDidSelectShippingContactHandler(controller PaymentAuthorizationViewController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)
	HasPaymentAuthorizationViewControllerDidSelectShippingContactHandler() bool
	PaymentAuthorizationViewControllerWillAuthorizePayment(controller PaymentAuthorizationViewController /* not a class type */)
	HasPaymentAuthorizationViewControllerWillAuthorizePayment() bool
}

// PaymentAuthorizationViewControllerDelegate is a delegate implementation builder for the PPaymentAuthorizationViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PaymentAuthorizationViewControllerDelegate struct {
	_PaymentAuthorizationViewControllerDidAuthorizePaymentCompletion func(controller PaymentAuthorizationViewController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)
	_PaymentAuthorizationViewControllerDidAuthorizePaymentHandler func(controller PaymentAuthorizationViewController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)
	_PaymentAuthorizationViewControllerDidChangeCouponCodeHandler func(controller PaymentAuthorizationViewController /* not a class type */, couponCode objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer)
	_PaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate func(controller PaymentAuthorizationViewController /* not a class type */, handler unsafe.Pointer)
	_PaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion func(controller PaymentAuthorizationViewController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationViewControllerDidSelectShippingMethodCompletion func(controller PaymentAuthorizationViewController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationViewControllerDidSelectPaymentMethodHandler func(controller PaymentAuthorizationViewController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationViewControllerDidSelectShippingMethodHandler func(controller PaymentAuthorizationViewController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationViewControllerDidSelectShippingAddressCompletion func(controller PaymentAuthorizationViewController /* not a class type */, address unsafe.Pointer, completion unsafe.Pointer)
	_PaymentAuthorizationViewControllerDidSelectShippingContactCompletion func(controller PaymentAuthorizationViewController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationViewControllerDidSelectShippingContactHandler func(controller PaymentAuthorizationViewController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)
	_PaymentAuthorizationViewControllerWillAuthorizePayment func(controller PaymentAuthorizationViewController /* not a class type */)
	_PaymentAuthorizationViewControllerDidFinish func(controller PaymentAuthorizationViewController /* not a class type */)
}

// SetPaymentAuthorizationViewControllerDidAuthorizePaymentCompletion sets the handler for the PaymentAuthorizationViewControllerDidAuthorizePaymentCompletion delegate method.
//
// Tells the delegate that the user authorized the payment request, and asks for a result.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidAuthorizePaymentCompletion(f func(controller PaymentAuthorizationViewController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidAuthorizePaymentCompletion = f
}

// SetPaymentAuthorizationViewControllerDidAuthorizePaymentHandler sets the handler for the PaymentAuthorizationViewControllerDidAuthorizePaymentHandler delegate method.
//
// Tells the delegate that the user authorized the payment request, and asks for a result.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidAuthorizePaymentHandler(f func(controller PaymentAuthorizationViewController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidAuthorizePaymentHandler = f
}

// SetPaymentAuthorizationViewControllerDidChangeCouponCodeHandler sets the handler for the PaymentAuthorizationViewControllerDidChangeCouponCodeHandler delegate method.
//
// Tells the delegate that the user entered or updated a coupon code.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidChangeCouponCodeHandler(f func(controller PaymentAuthorizationViewController /* not a class type */, couponCode objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidChangeCouponCodeHandler = f
}

// SetPaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate sets the handler for the PaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate delegate method.
//
// Requests an object that validates the identity of a merchant for a payment request.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate(f func(controller PaymentAuthorizationViewController /* not a class type */, handler unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate = f
}

// SetPaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion sets the handler for the PaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion delegate method.
//
// Tells the delegate that the user changed the payment method, and asks for an updated payment request.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion(f func(controller PaymentAuthorizationViewController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion = f
}

// SetPaymentAuthorizationViewControllerDidSelectShippingMethodCompletion sets the handler for the PaymentAuthorizationViewControllerDidSelectShippingMethodCompletion delegate method.
//
// Tells the delegate that the user selected a shipping method, and asks for an updated payment request.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidSelectShippingMethodCompletion(f func(controller PaymentAuthorizationViewController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidSelectShippingMethodCompletion = f
}

// SetPaymentAuthorizationViewControllerDidSelectPaymentMethodHandler sets the handler for the PaymentAuthorizationViewControllerDidSelectPaymentMethodHandler delegate method.
//
// Tells the delegate that the user changed the payment method, and asks for an updated payment request.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidSelectPaymentMethodHandler(f func(controller PaymentAuthorizationViewController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidSelectPaymentMethodHandler = f
}

// SetPaymentAuthorizationViewControllerDidSelectShippingMethodHandler sets the handler for the PaymentAuthorizationViewControllerDidSelectShippingMethodHandler delegate method.
//
// Tells the delegate that the user selected a shipping method, and asks for an updated payment request.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidSelectShippingMethodHandler(f func(controller PaymentAuthorizationViewController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidSelectShippingMethodHandler = f
}

// SetPaymentAuthorizationViewControllerDidSelectShippingAddressCompletion sets the handler for the PaymentAuthorizationViewControllerDidSelectShippingAddressCompletion delegate method.
//
// Tells the delegate that the user selected a shipping address.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidSelectShippingAddressCompletion(f func(controller PaymentAuthorizationViewController /* not a class type */, address unsafe.Pointer, completion unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidSelectShippingAddressCompletion = f
}

// SetPaymentAuthorizationViewControllerDidSelectShippingContactCompletion sets the handler for the PaymentAuthorizationViewControllerDidSelectShippingContactCompletion delegate method.
//
// Tells the delegate that the user selected a shipping address, and asks for an updated payment request.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidSelectShippingContactCompletion(f func(controller PaymentAuthorizationViewController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidSelectShippingContactCompletion = f
}

// SetPaymentAuthorizationViewControllerDidSelectShippingContactHandler sets the handler for the PaymentAuthorizationViewControllerDidSelectShippingContactHandler delegate method.
//
// Tells the delegate that the user selected a shipping address, and asks for an updated payment request.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidSelectShippingContactHandler(f func(controller PaymentAuthorizationViewController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer)) {
	d._PaymentAuthorizationViewControllerDidSelectShippingContactHandler = f
}

// SetPaymentAuthorizationViewControllerWillAuthorizePayment sets the handler for the PaymentAuthorizationViewControllerWillAuthorizePayment delegate method.
//
// Tells the delegate that the user is authorizing the payment request.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerWillAuthorizePayment(f func(controller PaymentAuthorizationViewController /* not a class type */)) {
	d._PaymentAuthorizationViewControllerWillAuthorizePayment = f
}

// SetPaymentAuthorizationViewControllerDidFinish sets the handler for the PaymentAuthorizationViewControllerDidFinish delegate method.
//
// Tells the delegate that payment authorization finished.
func (d *PaymentAuthorizationViewControllerDelegate) SetPaymentAuthorizationViewControllerDidFinish(f func(controller PaymentAuthorizationViewController /* not a class type */)) {
	d._PaymentAuthorizationViewControllerDidFinish = f
}

// PaymentAuthorizationViewControllerDidAuthorizePaymentCompletion implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidAuthorizePaymentCompletion(controller PaymentAuthorizationViewController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidAuthorizePaymentCompletion != nil {
		d._PaymentAuthorizationViewControllerDidAuthorizePaymentCompletion(controller, payment, completion)
	}
}

// HasPaymentAuthorizationViewControllerDidAuthorizePaymentCompletion returns true if a handler for PaymentAuthorizationViewControllerDidAuthorizePaymentCompletion has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidAuthorizePaymentCompletion() bool {
	return d._PaymentAuthorizationViewControllerDidAuthorizePaymentCompletion != nil
}

// PaymentAuthorizationViewControllerDidAuthorizePaymentHandler implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidAuthorizePaymentHandler(controller PaymentAuthorizationViewController /* not a class type */, payment storekit.Payment, completion unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidAuthorizePaymentHandler != nil {
		d._PaymentAuthorizationViewControllerDidAuthorizePaymentHandler(controller, payment, completion)
	}
}

// HasPaymentAuthorizationViewControllerDidAuthorizePaymentHandler returns true if a handler for PaymentAuthorizationViewControllerDidAuthorizePaymentHandler has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidAuthorizePaymentHandler() bool {
	return d._PaymentAuthorizationViewControllerDidAuthorizePaymentHandler != nil
}

// PaymentAuthorizationViewControllerDidChangeCouponCodeHandler implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidChangeCouponCodeHandler(controller PaymentAuthorizationViewController /* not a class type */, couponCode objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidChangeCouponCodeHandler != nil {
		d._PaymentAuthorizationViewControllerDidChangeCouponCodeHandler(controller, couponCode, completion)
	}
}

// HasPaymentAuthorizationViewControllerDidChangeCouponCodeHandler returns true if a handler for PaymentAuthorizationViewControllerDidChangeCouponCodeHandler has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidChangeCouponCodeHandler() bool {
	return d._PaymentAuthorizationViewControllerDidChangeCouponCodeHandler != nil
}

// PaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate(controller PaymentAuthorizationViewController /* not a class type */, handler unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate != nil {
		d._PaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate(controller, handler)
	}
}

// HasPaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate returns true if a handler for PaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate() bool {
	return d._PaymentAuthorizationViewControllerDidRequestMerchantSessionUpdate != nil
}

// PaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion(controller PaymentAuthorizationViewController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion != nil {
		d._PaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion(controller, paymentMethod, completion)
	}
}

// HasPaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion returns true if a handler for PaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion() bool {
	return d._PaymentAuthorizationViewControllerDidSelectPaymentMethodCompletion != nil
}

// PaymentAuthorizationViewControllerDidSelectShippingMethodCompletion implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidSelectShippingMethodCompletion(controller PaymentAuthorizationViewController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidSelectShippingMethodCompletion != nil {
		d._PaymentAuthorizationViewControllerDidSelectShippingMethodCompletion(controller, shippingMethod, completion)
	}
}

// HasPaymentAuthorizationViewControllerDidSelectShippingMethodCompletion returns true if a handler for PaymentAuthorizationViewControllerDidSelectShippingMethodCompletion has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidSelectShippingMethodCompletion() bool {
	return d._PaymentAuthorizationViewControllerDidSelectShippingMethodCompletion != nil
}

// PaymentAuthorizationViewControllerDidSelectPaymentMethodHandler implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidSelectPaymentMethodHandler(controller PaymentAuthorizationViewController /* not a class type */, paymentMethod PaymentMethod /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidSelectPaymentMethodHandler != nil {
		d._PaymentAuthorizationViewControllerDidSelectPaymentMethodHandler(controller, paymentMethod, completion)
	}
}

// HasPaymentAuthorizationViewControllerDidSelectPaymentMethodHandler returns true if a handler for PaymentAuthorizationViewControllerDidSelectPaymentMethodHandler has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidSelectPaymentMethodHandler() bool {
	return d._PaymentAuthorizationViewControllerDidSelectPaymentMethodHandler != nil
}

// PaymentAuthorizationViewControllerDidSelectShippingMethodHandler implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidSelectShippingMethodHandler(controller PaymentAuthorizationViewController /* not a class type */, shippingMethod ShippingMethod /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidSelectShippingMethodHandler != nil {
		d._PaymentAuthorizationViewControllerDidSelectShippingMethodHandler(controller, shippingMethod, completion)
	}
}

// HasPaymentAuthorizationViewControllerDidSelectShippingMethodHandler returns true if a handler for PaymentAuthorizationViewControllerDidSelectShippingMethodHandler has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidSelectShippingMethodHandler() bool {
	return d._PaymentAuthorizationViewControllerDidSelectShippingMethodHandler != nil
}

// PaymentAuthorizationViewControllerDidSelectShippingAddressCompletion implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidSelectShippingAddressCompletion(controller PaymentAuthorizationViewController /* not a class type */, address unsafe.Pointer, completion unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidSelectShippingAddressCompletion != nil {
		d._PaymentAuthorizationViewControllerDidSelectShippingAddressCompletion(controller, address, completion)
	}
}

// HasPaymentAuthorizationViewControllerDidSelectShippingAddressCompletion returns true if a handler for PaymentAuthorizationViewControllerDidSelectShippingAddressCompletion has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidSelectShippingAddressCompletion() bool {
	return d._PaymentAuthorizationViewControllerDidSelectShippingAddressCompletion != nil
}

// PaymentAuthorizationViewControllerDidSelectShippingContactCompletion implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidSelectShippingContactCompletion(controller PaymentAuthorizationViewController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidSelectShippingContactCompletion != nil {
		d._PaymentAuthorizationViewControllerDidSelectShippingContactCompletion(controller, contact, completion)
	}
}

// HasPaymentAuthorizationViewControllerDidSelectShippingContactCompletion returns true if a handler for PaymentAuthorizationViewControllerDidSelectShippingContactCompletion has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidSelectShippingContactCompletion() bool {
	return d._PaymentAuthorizationViewControllerDidSelectShippingContactCompletion != nil
}

// PaymentAuthorizationViewControllerDidSelectShippingContactHandler implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidSelectShippingContactHandler(controller PaymentAuthorizationViewController /* not a class type */, contact Contact /* not a class type */, completion unsafe.Pointer) {
	if d._PaymentAuthorizationViewControllerDidSelectShippingContactHandler != nil {
		d._PaymentAuthorizationViewControllerDidSelectShippingContactHandler(controller, contact, completion)
	}
}

// HasPaymentAuthorizationViewControllerDidSelectShippingContactHandler returns true if a handler for PaymentAuthorizationViewControllerDidSelectShippingContactHandler has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidSelectShippingContactHandler() bool {
	return d._PaymentAuthorizationViewControllerDidSelectShippingContactHandler != nil
}

// PaymentAuthorizationViewControllerWillAuthorizePayment implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerWillAuthorizePayment(controller PaymentAuthorizationViewController /* not a class type */) {
	if d._PaymentAuthorizationViewControllerWillAuthorizePayment != nil {
		d._PaymentAuthorizationViewControllerWillAuthorizePayment(controller)
	}
}

// HasPaymentAuthorizationViewControllerWillAuthorizePayment returns true if a handler for PaymentAuthorizationViewControllerWillAuthorizePayment has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerWillAuthorizePayment() bool {
	return d._PaymentAuthorizationViewControllerWillAuthorizePayment != nil
}

// PaymentAuthorizationViewControllerDidFinish implements the PPaymentAuthorizationViewControllerDelegate interface.
func (d *PaymentAuthorizationViewControllerDelegate) PaymentAuthorizationViewControllerDidFinish(controller PaymentAuthorizationViewController /* not a class type */) {
	if d._PaymentAuthorizationViewControllerDidFinish != nil {
		d._PaymentAuthorizationViewControllerDidFinish(controller)
	}
}

// HasPaymentAuthorizationViewControllerDidFinish returns true if a handler for PaymentAuthorizationViewControllerDidFinish has been set.
func (d *PaymentAuthorizationViewControllerDelegate) HasPaymentAuthorizationViewControllerDidFinish() bool {
	return d._PaymentAuthorizationViewControllerDidFinish != nil
}
