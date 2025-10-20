package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/contactsui"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("ContactsUI Framework Examples")
	fmt.Println("=============================")

	// Example 1: Create CNContactPickerViewController
	fmt.Println("\n1. Creating CNContactPickerViewController:")

	contactPicker := contactsui.NewCNContactPickerViewController()
	fmt.Printf("   Contact picker created: %v\n", contactPicker)

	// Example 2: ContactsUI components
	fmt.Println("\n2. ContactsUI Components:")

	components := map[string]string{
		"CNContactPickerViewController": "Select contacts from library",
		"CNContactViewController":       "Display/edit single contact",
		"CNContactStore":                "Access contacts (Contacts framework)",
		"CNContact":                     "Contact data model",
	}

	for component, desc := range components {
		fmt.Printf("   %-35s: %s\n", component, desc)
	}

	// Example 3: Contact picker modes
	fmt.Println("\n3. Contact Picker Modes:")

	modes := map[string]string{
		"Single contact":        "User selects one contact",
		"Multiple contacts":     "User selects multiple contacts",
		"Single property":       "Select specific property (email, phone)",
		"Multiple properties":   "Select multiple properties",
	}

	for mode, desc := range modes {
		fmt.Printf("   %-25s: %s\n", mode, desc)
	}

	// Example 4: Contact picker workflow
	fmt.Println("\n4. Contact Picker Workflow:")

	workflow := []string{
		"1. Create CNContactPickerViewController",
		"2. Set delegate for callbacks",
		"3. Optional: Set displayedPropertyKeys filter",
		"4. Optional: Set predicateForEnabling...",
		"5. Present view controller modally",
		"6. User browses and selects contacts",
		"7. Delegate receives selected contacts",
		"8. Dismiss picker automatically",
	}

	for _, step := range workflow {
		fmt.Printf("   %s\n", step)
	}

	// Example 5: Delegate methods
	fmt.Println("\n5. Contact Picker Delegate Methods:")

	delegateMethods := []string{
		"contactPickerDidCancel - User cancelled",
		"didSelectContact - Selected single contact",
		"didSelectContacts - Selected multiple contacts",
		"didSelectContactProperty - Selected property",
		"didSelectContactProperties - Selected multiple properties",
	}

	for i, method := range delegateMethods {
		fmt.Printf("   %2d. %s\n", i+1, method)
	}

	// Example 6: Contact properties
	fmt.Println("\n6. Available Contact Properties:")

	properties := []string{
		"givenName - First name",
		"familyName - Last name",
		"organizationName - Company",
		"jobTitle - Job title",
		"phoneNumbers - Phone numbers array",
		"emailAddresses - Email addresses array",
		"postalAddresses - Mailing addresses",
		"birthday - Date of birth",
		"dates - Other dates (anniversary, etc.)",
		"urlAddresses - Websites",
		"socialProfiles - Social media",
		"instantMessageAddresses - IM accounts",
		"note - Notes field",
		"imageData - Profile photo",
		"thumbnailImageData - Small profile photo",
	}

	for i, property := range properties {
		fmt.Printf("   %2d. %s\n", i+1, property)
	}

	// Example 7: Display properties
	fmt.Println("\n7. Commonly Displayed Properties:")

	displayProps := map[string]string{
		"Phone Numbers":   "CNContactPhoneNumbersKey",
		"Email Addresses": "CNContactEmailAddressesKey",
		"Postal Addresses": "CNContactPostalAddressesKey",
		"URLs":            "CNContactUrlAddressesKey",
		"Birthday":        "CNContactBirthdayKey",
		"Organization":    "CNContactOrganizationNameKey",
	}

	for name, key := range displayProps {
		fmt.Printf("   %-20s: %s\n", name, key)
	}

	// Example 8: Contact view controller modes
	fmt.Println("\n8. CNContactViewController Modes:")

	vcModes := map[string]string{
		"Unknown":  "Initial state",
		"Display":  "Read-only contact display",
		"Edit":     "Edit existing contact",
		"New":      "Create new contact",
	}

	for mode, desc := range vcModes {
		fmt.Printf("   %-10s: %s\n", mode, desc)
	}

	// Example 9: Contact view controller workflow
	fmt.Println("\n9. Contact View Controller Workflow:")

	vcWorkflow := []string{
		"1. Get contact from store or create new",
		"2. Create CNContactViewController with contact",
		"3. Set delegate for callbacks",
		"4. Embed in UINavigationController",
		"5. Present navigation controller",
		"6. User views/edits contact",
		"7. Delegate receives completion callback",
		"8. Save changes to store if needed",
	}

	for _, step := range vcWorkflow {
		fmt.Printf("   %s\n", step)
	}

	// Example 10: Common use cases
	fmt.Println("\n10. Common Use Cases:")

	useCases := map[string]string{
		"Share Content":      "Select contact to share with",
		"Send Message":       "Pick recipient for message",
		"Make Call":          "Select contact to call",
		"Send Email":         "Pick email recipient",
		"Invite Users":       "Select contacts to invite",
		"Add Friend":         "Import contact to app",
		"Group Creation":     "Select group members",
		"Contact Display":    "Show contact details",
		"Contact Edit":       "Edit contact information",
		"Quick Contact":      "Autocomplete contact fields",
	}

	for useCase, desc := range useCases {
		fmt.Printf("   %-20s: %s\n", useCase, desc)
	}

	// Example 11: Privacy and permissions
	fmt.Println("\n11. Privacy Considerations:")

	privacyNotes := []string{
		"Contacts permission required",
		"Request access via CNContactStore",
		"User must grant permission",
		"Limited access option available (iOS 18+)",
		"Picker doesn't require permission",
		"Only selected contacts shared with app",
		"User controls all data sharing",
		"Cannot access all contacts without permission",
	}

	for i, note := range privacyNotes {
		fmt.Printf("   %2d. %s\n", i+1, note)
	}

	// Example 12: Filtering contacts
	fmt.Println("\n12. Contact Filtering Options:")

	filterOptions := []string{
		"predicateForEnablingContact - Enable specific contacts",
		"predicateForSelectionOfContact - Allow selection filter",
		"predicateForSelectionOfProperty - Property selection filter",
		"displayedPropertyKeys - Show only specific properties",
	}

	for i, option := range filterOptions {
		fmt.Printf("   %2d. %s\n", i+1, option)
	}

	// Example 13: Contact actions
	fmt.Println("\n13. Built-in Contact Actions:")

	actions := []string{
		"Send Message - iMessage or SMS",
		"FaceTime Audio - Voice call",
		"FaceTime Video - Video call",
		"Mail - Send email",
		"Pay - Apple Pay transaction",
		"Share Contact - Share via AirDrop/Messages",
		"Add to Favorites - Star contact",
		"Create New Contact - Add to library",
	}

	for i, action := range actions {
		fmt.Printf("   %2d. %s\n", i+1, action)
	}

	// Example 14: Integration with other frameworks
	fmt.Println("\n14. Framework Integration:")

	integrations := map[string]string{
		"Contacts":          "CNContactStore for data access",
		"MessageUI":         "Send email/SMS to contact",
		"EventKit":          "Create events with attendees",
		"MapKit":            "Show contact address on map",
		"Social":            "Share via social networks",
		"CallKit":           "Make phone calls",
		"PassKit":           "Apple Pay to contact",
	}

	for framework, desc := range integrations {
		fmt.Printf("   %-20s: %s\n", framework, desc)
	}

	// Example 15: Best practices
	fmt.Println("\n15. Best Practices:")

	bestPractices := []string{
		"Request permission before showing contact UI",
		"Use picker when possible (no permission needed)",
		"Explain why you need contact access",
		"Only request properties you actually need",
		"Handle permission denial gracefully",
		"Don't cache contact data unnecessarily",
		"Respect user's limited access choices",
		"Update cached contacts when changed",
	}

	for i, practice := range bestPractices {
		fmt.Printf("   %2d. %s\n", i+1, practice)
	}

	fmt.Println("\n✓ ContactsUI framework examples completed!")
	fmt.Println("\nNote: ContactsUI provides ready-to-use contact selection:")
	fmt.Println("  - Built-in picker UI (no permission needed)")
	fmt.Println("  - Contact view/edit controllers")
	fmt.Println("  - Property filtering")
	fmt.Println("  - Privacy-friendly design")
	fmt.Println("  - Seamless system integration")
	fmt.Println("\nReal applications would:")
	fmt.Println("  - Present CNContactPickerViewController")
	fmt.Println("  - Implement delegate for selections")
	fmt.Println("  - Filter displayed properties")
	fmt.Println("  - Handle selected contact data")
	fmt.Println("  - Request contacts permission if needed")
	fmt.Println("  - Use CNContactViewController for display/edit")
}
