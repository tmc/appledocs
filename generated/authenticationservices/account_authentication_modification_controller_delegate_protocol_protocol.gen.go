// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PAccountAuthenticationModificationControllerDelegate is the ASAccountAuthenticationModificationControllerDelegate protocol interface.
//
// An interface you implement for receiving success and failure statuses about modification of an account’s authentication properties.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.authenticationservices/documentation/AuthenticationServices/ASAccountAuthenticationModificationControllerDelegate
type PAccountAuthenticationModificationControllerDelegate interface {
	// Optional methods
	AccountAuthenticationModificationControllerDidFailRequestWithError(controller IASAccountAuthenticationModificationController, request IASAccountAuthenticationModificationRequest, error_ objc.IObject /* cross-framework: Error */)
	HasAccountAuthenticationModificationControllerDidFailRequestWithError() bool
	AccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo(controller IASAccountAuthenticationModificationController, request IASAccountAuthenticationModificationRequest, userInfo objc.IObject /* cross-framework: NSDictionary */)
	HasAccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo() bool
}

// AccountAuthenticationModificationControllerDelegate is a delegate implementation builder for the PAccountAuthenticationModificationControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AccountAuthenticationModificationControllerDelegate struct {
	_AccountAuthenticationModificationControllerDidFailRequestWithError func(controller IASAccountAuthenticationModificationController, request IASAccountAuthenticationModificationRequest, error_ objc.IObject /* cross-framework: Error */)
	_AccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo func(controller IASAccountAuthenticationModificationController, request IASAccountAuthenticationModificationRequest, userInfo objc.IObject /* cross-framework: NSDictionary */)
}

// SetAccountAuthenticationModificationControllerDidFailRequestWithError sets the handler for the AccountAuthenticationModificationControllerDidFailRequestWithError delegate method.
//
// Tells the delegate an account modification request failed.
func (d *AccountAuthenticationModificationControllerDelegate) SetAccountAuthenticationModificationControllerDidFailRequestWithError(f func(controller IASAccountAuthenticationModificationController, request IASAccountAuthenticationModificationRequest, error_ objc.IObject /* cross-framework: Error */)) {
	d._AccountAuthenticationModificationControllerDidFailRequestWithError = f
}

// SetAccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo sets the handler for the AccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo delegate method.
//
// Tells the delegate an account modification request completed successfully.
func (d *AccountAuthenticationModificationControllerDelegate) SetAccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo(f func(controller IASAccountAuthenticationModificationController, request IASAccountAuthenticationModificationRequest, userInfo objc.IObject /* cross-framework: NSDictionary */)) {
	d._AccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo = f
}

// AccountAuthenticationModificationControllerDidFailRequestWithError implements the PAccountAuthenticationModificationControllerDelegate interface.
func (d *AccountAuthenticationModificationControllerDelegate) AccountAuthenticationModificationControllerDidFailRequestWithError(controller IASAccountAuthenticationModificationController, request IASAccountAuthenticationModificationRequest, error_ objc.IObject /* cross-framework: Error */) {
	if d._AccountAuthenticationModificationControllerDidFailRequestWithError != nil {
		d._AccountAuthenticationModificationControllerDidFailRequestWithError(controller, request, error_)
	}
}

// HasAccountAuthenticationModificationControllerDidFailRequestWithError returns true if a handler for AccountAuthenticationModificationControllerDidFailRequestWithError has been set.
func (d *AccountAuthenticationModificationControllerDelegate) HasAccountAuthenticationModificationControllerDidFailRequestWithError() bool {
	return d._AccountAuthenticationModificationControllerDidFailRequestWithError != nil
}

// AccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo implements the PAccountAuthenticationModificationControllerDelegate interface.
func (d *AccountAuthenticationModificationControllerDelegate) AccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo(controller IASAccountAuthenticationModificationController, request IASAccountAuthenticationModificationRequest, userInfo objc.IObject /* cross-framework: NSDictionary */) {
	if d._AccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo != nil {
		d._AccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo(controller, request, userInfo)
	}
}

// HasAccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo returns true if a handler for AccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo has been set.
func (d *AccountAuthenticationModificationControllerDelegate) HasAccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo() bool {
	return d._AccountAuthenticationModificationControllerDidSuccessfullyCompleteRequestWithUserInfo != nil
}
