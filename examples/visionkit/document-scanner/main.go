package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/visionkit"
)

var e2e = flag.Bool("e2e", false, "run end-to-end tests")

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	flag.Parse()

	fmt.Println("VisionKit Framework Examples")
	fmt.Println("============================")

	// Example 1: Create VNDocumentCameraViewController
	fmt.Println("\n1. Creating VNDocumentCameraViewController:")

	documentCamera := visionkit.NewDocumentCameraViewController()
	fmt.Printf("   Document camera created: %v\n", documentCamera)

	// Example 2: VisionKit components
	fmt.Println("\n2. VisionKit Components:")

	components := map[string]string{
		"VNDocumentCameraViewController": "Built-in document scanner UI",
		"VNDocumentCameraScan":           "Scanned document result",
		"VNDocumentCameraDelegate":       "Scanner event callbacks",
		"VNImageAnalysisInteraction":     "Visual lookup and text selection",
		"VNImageAnalyzer":                "Analyze images for features",
	}

	for component, desc := range components {
		fmt.Printf("   %-35s: %s\n", component, desc)
	}

	// Example 3: Document scanner features
	fmt.Println("\n3. Document Scanner Features:")

	features := []string{
		"Automatic document detection",
		"Perspective correction",
		"Multi-page scanning",
		"Auto-capture when document detected",
		"Manual capture override",
		"Flash control",
		"Filter selection (color, grayscale, B&W)",
		"Page reordering",
		"Page deletion",
		"Review scanned pages",
		"Built-in UI (no custom implementation needed)",
	}

	for i, feature := range features {
		fmt.Printf("   %2d. %s\n", i+1, feature)
	}

	// Example 4: Document scanning workflow
	fmt.Println("\n4. Document Scanning Workflow:")

	workflow := []string{
		"1. Create VNDocumentCameraViewController",
		"2. Set delegate for callbacks",
		"3. Present view controller modally",
		"4. User positions document in frame",
		"5. Camera auto-detects document edges",
		"6. Auto-captures or user taps shutter",
		"7. User can scan multiple pages",
		"8. User taps Save",
		"9. Delegate receives scanned document",
		"10. Process or save scanned images",
	}

	for _, step := range workflow {
		fmt.Printf("   %s\n", step)
	}

	// Example 5: Scan result properties
	fmt.Println("\n5. Scanned Document Properties:")

	properties := map[string]string{
		"pageCount":   "Number of scanned pages",
		"imageOfPage": "Get UIImage for specific page",
		"title":       "Document title (optional)",
	}

	for property, desc := range properties {
		fmt.Printf("   %-15s: %s\n", property, desc)
	}

	// Example 6: Delegate methods
	fmt.Println("\n6. Document Camera Delegate Methods:")

	delegateMethods := []string{
		"didFinishWithScan - User completed scanning",
		"didFailWithError - Scanning failed",
		"didCancel - User cancelled scanning",
	}

	for i, method := range delegateMethods {
		fmt.Printf("   %2d. %s\n", i+1, method)
	}

	// Example 7: Image analysis features (iOS 16+)
	fmt.Println("\n7. Image Analysis Features (Live Text):")

	analysisFeatures := []string{
		"Text selection and copy",
		"Translate text",
		"Data detector actions (phone, email, address)",
		"QR code detection and actions",
		"Visual lookup (identify objects, landmarks)",
		"Subject lifting (copy subject from photo)",
		"Language translation",
	}

	for i, feature := range analysisFeatures {
		fmt.Printf("   %2d. %s\n", i+1, feature)
	}

	// Example 8: Analysis interaction types
	fmt.Println("\n8. Image Analysis Interaction Types:")

	interactionTypes := []string{
		"Automatic - All interactions enabled",
		"Text selection - Select and copy text only",
		"Data detectors - Phone, email, address links",
		"Visual lookup - Identify objects and scenes",
		"None - Disable all interactions",
	}

	for i, interactionType := range interactionTypes {
		fmt.Printf("   %2d. %s\n", i+1, interactionType)
	}

	// Example 9: Image analyzer analysis types
	fmt.Println("\n9. Analysis Types:")

	analysisTypes := []string{
		"Text - Recognize and select text",
		"Machine readable codes - QR codes, barcodes",
		"Visual lookup - Object and landmark recognition",
		"Subject - Identify primary subject",
	}

	for i, analysisType := range analysisTypes {
		fmt.Printf("   %2d. %s\n", i+1, analysisType)
	}

	// Example 10: Common use cases
	fmt.Println("\n10. Common Use Cases:")

	useCases := map[string]string{
		"Document Scanning":  "Scan receipts, contracts, forms",
		"Note Taking":        "Digitize handwritten notes",
		"Business Cards":     "Scan and extract contact info",
		"Whiteboard Capture": "Save meeting whiteboard content",
		"Book/Magazine":      "Scan pages from books",
		"Receipt Tracking":   "Expense management apps",
		"Document Storage":   "Digital filing system",
		"ID Scanning":        "Capture ID cards, passports",
		"PDF Creation":       "Multi-page PDF documents",
		"OCR Pipeline":       "Extract text from documents",
	}

	for useCase, desc := range useCases {
		fmt.Printf("   %-20s: %s\n", useCase, desc)
	}

	// Example 11: Live Text use cases
	fmt.Println("\n11. Live Text Use Cases:")

	liveTextCases := []string{
		"Copy text from photos",
		"Translate captured text",
		"Call phone numbers in images",
		"Email addresses become links",
		"Detect and open URLs",
		"Add addresses to Maps",
		"Track packages from tracking numbers",
		"Copy email addresses",
		"Identify plants and animals",
		"Identify landmarks and art",
		"Identify food and dishes",
	}

	for i, useCase := range liveTextCases {
		fmt.Printf("   %2d. %s\n", i+1, useCase)
	}

	// Example 12: Image quality tips
	fmt.Println("\n12. Tips for Best Scan Quality:")

	qualityTips := []string{
		"Good lighting - Avoid shadows",
		"Flat surface - Reduce wrinkles",
		"Contrast - Clear background",
		"Hold steady - Reduce blur",
		"Fill frame - Document edges visible",
		"Perpendicular angle - Avoid skew",
		"Multiple shots - Capture backup pages",
		"Use filters - B&W for text documents",
	}

	for i, tip := range qualityTips {
		fmt.Printf("   %2d. %s\n", i+1, tip)
	}

	// Example 13: Supported languages (Live Text)
	fmt.Println("\n13. Live Text Language Support:")

	languages := []string{
		"English, Chinese, French, German",
		"Italian, Japanese, Korean, Portuguese",
		"Russian, Spanish, Ukrainian",
		"And many more...",
	}

	for i, lang := range languages {
		fmt.Printf("   %2d. %s\n", i+1, lang)
	}

	// Example 14: Privacy and permissions
	fmt.Println("\n14. Privacy Considerations:")

	privacyNotes := []string{
		"Camera permission required",
		"User controls all scanning",
		"No automatic cloud upload",
		"Scanned images stored locally",
		"App handles image storage",
		"On-device text recognition",
		"User must explicitly share data",
	}

	for i, note := range privacyNotes {
		fmt.Printf("   %2d. %s\n", i+1, note)
	}

	// Example 15: Integration points
	fmt.Println("\n15. Framework Integration:")

	integrations := map[string]string{
		"Vision Framework":    "OCR with VNRecognizeTextRequest",
		"Core Image":          "Apply filters to scanned images",
		"PDFKit":              "Create PDF from scanned pages",
		"Photos Framework":    "Save to photo library",
		"UIDocumentPickerVC":  "Save/share documents",
		"CloudKit":            "Sync scanned documents",
		"Core ML":             "Custom document classification",
		"Natural Language":    "Process extracted text",
	}

	for framework, desc := range integrations {
		fmt.Printf("   %-20s: %s\n", framework, desc)
	}

	fmt.Println("\n✓ VisionKit framework examples completed!")
	fmt.Println("\nNote: VisionKit provides ready-to-use document scanning:")
	fmt.Println("  - Built-in camera UI with auto-detection")
	fmt.Println("  - Perspective correction and filtering")
	fmt.Println("  - Multi-page document support")
	fmt.Println("  - Live Text for text selection")
	fmt.Println("  - Visual lookup for object recognition")
	fmt.Println("\nReal applications would:")
	fmt.Println("  - Present VNDocumentCameraViewController")
	fmt.Println("  - Implement delegate for scan results")
	fmt.Println("  - Process scanned images")
	fmt.Println("  - Extract text with Vision framework")
	fmt.Println("  - Create PDFs from scanned pages")
	fmt.Println("  - Add Live Text interaction to images")
}
