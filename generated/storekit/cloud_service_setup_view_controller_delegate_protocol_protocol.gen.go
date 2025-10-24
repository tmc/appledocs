// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

// PCloudServiceSetupViewControllerDelegate is the SKCloudServiceSetupViewControllerDelegate protocol interface.
//
// A protocol that defines the methods a cloud service setup view controller can use to get the status of the view, including when it is dismissed.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - iOS 10.1+ (Deprecated in 18.0)
//   - iPadOS 10.1+ (Deprecated in 18.0)
//
// See: doc://com.apple.storekit/documentation/StoreKit/SKCloudServiceSetupViewControllerDelegate
type PCloudServiceSetupViewControllerDelegate interface {
	// Optional methods
	CloudServiceSetupViewControllerDidDismiss(cloudServiceSetupViewController ISKCloudServiceSetupViewController)
	HasCloudServiceSetupViewControllerDidDismiss() bool
}

// CloudServiceSetupViewControllerDelegate is a delegate implementation builder for the PCloudServiceSetupViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CloudServiceSetupViewControllerDelegate struct {
	_CloudServiceSetupViewControllerDidDismiss func(cloudServiceSetupViewController ISKCloudServiceSetupViewController)
}

// SetCloudServiceSetupViewControllerDidDismiss sets the handler for the CloudServiceSetupViewControllerDidDismiss delegate method.
//
// Tells the delegate that the cloud service setup view controller was dismissed.
func (d *CloudServiceSetupViewControllerDelegate) SetCloudServiceSetupViewControllerDidDismiss(f func(cloudServiceSetupViewController ISKCloudServiceSetupViewController)) {
	d._CloudServiceSetupViewControllerDidDismiss = f
}

// CloudServiceSetupViewControllerDidDismiss implements the PCloudServiceSetupViewControllerDelegate interface.
func (d *CloudServiceSetupViewControllerDelegate) CloudServiceSetupViewControllerDidDismiss(cloudServiceSetupViewController ISKCloudServiceSetupViewController) {
	if d._CloudServiceSetupViewControllerDidDismiss != nil {
		d._CloudServiceSetupViewControllerDidDismiss(cloudServiceSetupViewController)
	}
}

// HasCloudServiceSetupViewControllerDidDismiss returns true if a handler for CloudServiceSetupViewControllerDidDismiss has been set.
func (d *CloudServiceSetupViewControllerDelegate) HasCloudServiceSetupViewControllerDidDismiss() bool {
	return d._CloudServiceSetupViewControllerDidDismiss != nil
}
