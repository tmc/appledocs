# WebKit Framework Example

This example demonstrates the WebKit framework for embedding web content in native applications using Go.

## What it demonstrates

- WKWebView creation and configuration
- Content loading methods
- JavaScript integration
- Native-to-web messaging
- Navigation and UI delegates
- Cookie and data management
- Content blocking
- Security features

## Running the example

```bash
go run main.go
# or with e2e flag
go run main.go -e2e
```

## Key Concepts

### WebKit Overview

WebKit provides modern web browsing capabilities:
- **WKWebView** - Modern replacement for UIWebView/WebView
- Full HTML5, CSS3, JavaScript ES6+ support
- Hardware-accelerated rendering with Metal
- Multi-process architecture for security
- Native-to-web messaging bridge
- Advanced privacy features (ITP)

### Why Use WebKit?

**For hybrid apps:**
- Embed web content in native UI
- Use web technologies for complex UI
- OAuth and web-based authentication
- Display rich content (HTML, Markdown, PDF)
- Build cross-platform UI with web

**Advantages:**
- Modern web standards
- Excellent performance
- Security sandboxing
- Easy JavaScript integration
- No UIWebView deprecation issues

## Creating a Web View

### Basic Setup

```go
import (
    "github.com/tmc/appledocs/generated/webkit"
    "github.com/tmc/appledocs/generated/foundation"
)

// Create configuration
config := webkit.NewWebViewConfiguration()

// Create web view
frame := CGRect{0, 0, 800, 600}
webView := webkit.NewWebViewWithFrame(frame, config)

// Add to view hierarchy (pseudo-code)
view.AddSubview(webView)
```

### With Preferences

```go
// Configure preferences
preferences := webkit.NewPreferences()
preferences.SetJavaScriptEnabled(true)
preferences.SetJavaScriptCanOpenWindowsAutomatically(false)
preferences.SetMinimumFontSize(12.0)

config := webkit.NewWebViewConfiguration()
config.SetPreferences(preferences)

webView := webkit.NewWebViewWithFrame(frame, config)
```

## Loading Content

### Load URL

```go
// Load URL
url := foundation.NewURLWithString("https://example.com")
request := foundation.NewURLRequestWithURL(url)
webView.LoadRequest(request)
```

### Load HTML String

```go
// Load HTML directly
html := `
<!DOCTYPE html>
<html>
<head><title>Hello</title></head>
<body>
    <h1>Hello from WebKit!</h1>
    <p>This is rendered HTML.</p>
</body>
</html>
`

baseURL := foundation.NewURLWithString("about:blank")
webView.LoadHTMLStringBaseURL(html, baseURL)
```

### Load Local File

```go
// Load local HTML file
fileURL := foundation.NewURLWithString("file:///path/to/index.html")
readAccessURL := foundation.NewURLWithString("file:///path/to/")

webView.LoadFileURLAllowingReadAccessToURL(fileURL, readAccessURL)
```

### Load Data

```go
// Load data with MIME type
data := []byte("<h1>Hello</h1>")
mimeType := "text/html"
encoding := "utf-8"
baseURL := foundation.NewURLWithString("about:blank")

webView.LoadDataMIMETypeCharacterEncodingNameBaseURL(
    data,
    mimeType,
    encoding,
    baseURL,
)
```

## JavaScript Integration

### Execute JavaScript

```go
// Simple execution
javascript := "document.title"
webView.EvaluateJavaScript(javascript) { result, error in
    if error != nil {
        // Handle error
        return
    }

    title := result as? String
    fmt.Printf("Page title: %s\n", title)
}

// Complex execution with async
javascript := `
    fetch('/api/data')
        .then(r => r.json())
        .then(data => data.value)
`
webView.EvaluateJavaScript(javascript) { result, error in
    // Handle result
}
```

### Inject User Scripts

```go
// Create user script
script := `
    console.log('Injected at document start');

    // Modify page
    document.addEventListener('DOMContentLoaded', function() {
        document.body.style.backgroundColor = '#f0f0f0';
    });
`

userScript := webkit.NewUserScript(
    script,
    .atDocumentStart, // or .atDocumentEnd
    false,            // forMainFrameOnly
)

// Add to content controller
contentController := config.UserContentController()
contentController.AddUserScript(userScript)
```

### Native-to-Web Messaging

```go
// In JavaScript, send message to native:
// window.webkit.messageHandlers.nativeHandler.postMessage({
//     action: 'buttonClicked',
//     data: 'some value'
// });

// In native code, handle messages:
type MessageHandler struct{}

func (h *MessageHandler) UserContentControllerDidReceiveScriptMessage(
    controller: WKUserContentController,
    message: WKScriptMessage,
) {
    name := message.Name()
    body := message.Body()

    if name == "nativeHandler" {
        // Parse body (dictionary, string, number, etc.)
        if dict, ok := body.(map[string]interface{}); ok {
            action := dict["action"]
            data := dict["data"]

            // Handle action
            if action == "buttonClicked" {
                fmt.Printf("Button clicked with: %v\n", data)
            }
        }
    }
}

// Register handler
handler := &MessageHandler{}
contentController.AddScriptMessageHandler(handler, "nativeHandler")
```

### Web-to-Native Callback

```go
// Inject callback into page
script := `
    function callNative(data) {
        window.webkit.messageHandlers.callback.postMessage(data);
    }

    // Make available globally
    window.nativeCallback = callNative;
`

userScript := webkit.NewUserScript(script, .atDocumentStart, false)
contentController.AddUserScript(userScript)

// In JavaScript on page:
// window.nativeCallback({status: 'ready'});
```

## Navigation Delegate

### Implementing WKNavigationDelegate

```go
type NavigationDelegate struct{}

// Decide whether to allow navigation
func (d *NavigationDelegate) WebViewDecidePolicyForNavigationAction(
    webView: WKWebView,
    navigationAction: WKNavigationAction,
    decisionHandler: func(WKNavigationActionPolicy),
) {
    url := navigationAction.Request().URL()

    // Allow only HTTPS
    if url.Scheme() == "https" {
        decisionHandler(.allow)
    } else {
        decisionHandler(.cancel)
        // Open in external browser instead
        // NSWorkspace.shared.open(url)
    }
}

// Navigation started
func (d *NavigationDelegate) WebViewDidStartProvisionalNavigation(
    webView: WKWebView,
    navigation: WKNavigation,
) {
    // Show loading indicator
    fmt.Println("Navigation started")
}

// Navigation committed
func (d *NavigationDelegate) WebViewDidCommitNavigation(
    webView: WKWebView,
    navigation: WKNavigation,
) {
    // Page content started loading
    fmt.Println("Navigation committed")
}

// Navigation finished
func (d *NavigationDelegate) WebViewDidFinishNavigation(
    webView: WKWebView,
    navigation: WKNavigation,
) {
    // Hide loading indicator
    fmt.Println("Navigation finished")

    // Get page title
    title := webView.Title()
    fmt.Printf("Loaded: %s\n", title)
}

// Navigation failed
func (d *NavigationDelegate) WebViewDidFailNavigation(
    webView: WKWebView,
    navigation: WKNavigation,
    error: NSError,
) {
    // Handle error
    fmt.Printf("Navigation failed: %v\n", error)
}

// Process crashed
func (d *NavigationDelegate) WebViewWebContentProcessDidTerminate(
    webView: WKWebView,
) {
    // Reload page or show error
    fmt.Println("Web process terminated")
    webView.Reload()
}

// Set delegate
delegate := &NavigationDelegate{}
webView.SetNavigationDelegate(delegate)
```

## UI Delegate

### Implementing WKUIDelegate

```go
type UIDelegate struct{}

// Handle JavaScript alert()
func (d *UIDelegate) WebViewRunJavaScriptAlertPanelWithMessage(
    webView: WKWebView,
    message: string,
    frame: WKFrameInfo,
    completionHandler: func(),
) {
    // Show native alert
    alert := NSAlert()
    alert.SetMessageText(message)
    alert.RunModal()

    completionHandler()
}

// Handle JavaScript confirm()
func (d *UIDelegate) WebViewRunJavaScriptConfirmPanelWithMessage(
    webView: WKWebView,
    message: string,
    frame: WKFrameInfo,
    completionHandler: func(bool),
) {
    // Show confirmation dialog
    alert := NSAlert()
    alert.SetMessageText(message)
    alert.AddButtonWithTitle("OK")
    alert.AddButtonWithTitle("Cancel")

    response := alert.RunModal()
    completionHandler(response == .alertFirstButtonReturn)
}

// Handle JavaScript prompt()
func (d *UIDelegate) WebViewRunJavaScriptTextInputPanelWithPrompt(
    webView: WKWebView,
    prompt: string,
    defaultText: string,
    frame: WKFrameInfo,
    completionHandler: func(string),
) {
    // Show input dialog
    // Get user input
    input := "user input" // Pseudo-code
    completionHandler(input)
}

// Handle window.open()
func (d *UIDelegate) WebViewCreateWebViewWithConfiguration(
    webView: WKWebView,
    configuration: WKWebViewConfiguration,
    navigationAction: WKNavigationAction,
    windowFeatures: WKWindowFeatures,
) -> WKWebView {
    // Create new web view for popup
    // Return new web view or nil to block
    return nil // Block popups
}

// Camera/microphone permission
func (d *UIDelegate) WebViewRequestMediaCapturePermission(
    webView: WKWebView,
    origin: WKSecurityOrigin,
    frame: WKFrameInfo,
    type: WKMediaCaptureType,
    decisionHandler: func(WKPermissionDecision),
) {
    // Request system permission
    // Then allow or deny
    decisionHandler(.grant) // or .deny
}

// Set delegate
delegate := &UIDelegate{}
webView.SetUIDelegate(delegate)
```

## Website Data Management

### Clear Cookies and Cache

```go
// Get data store
dataStore := webkit.WebsiteDataStore.DefaultDataStore()

// Clear all website data
dataTypes := webkit.WebsiteDataStore.AllWebsiteDataTypes()
date := foundation.DistantPast() // All data

dataStore.RemoveDataOfTypesModifiedSince(dataTypes, date) {
    fmt.Println("Data cleared")
}

// Clear specific data types
specificTypes := []string{
    webkit.WebsiteDataTypeCookies,
    webkit.WebsiteDataTypeDiskCache,
    webkit.WebsiteDataTypeMemoryCache,
}

dataStore.RemoveDataOfTypesModifiedSince(specificTypes, date) {
    fmt.Println("Cookies and cache cleared")
}
```

### Fetch Website Data

```go
// Get all stored data
dataStore.FetchDataRecords(dataTypes) { records in
    for record := range records {
        name := record.DisplayName()
        types := record.DataTypes()
        fmt.Printf("Site: %s, Types: %v\n", name, types)
    }
}

// Remove specific site data
dataStore.RemoveDataOfTypesForDataRecords(dataTypes, records) {
    fmt.Println("Site data removed")
}
```

### Custom Data Store

```go
// Non-persistent (private browsing)
dataStore := webkit.NewNonPersistentWebsiteDataStore()

config := webkit.NewWebViewConfiguration()
config.SetWebsiteDataStore(dataStore)

webView := webkit.NewWebViewWithFrame(frame, config)
// Data is cleared when data store is deallocated
```

## Content Blocking

### Create Blocking Rules

```go
// Define rules in JSON
rulesJSON := `
[
    {
        "trigger": {
            "url-filter": ".*",
            "resource-type": ["image"]
        },
        "action": {
            "type": "block"
        }
    },
    {
        "trigger": {
            "url-filter": "evil-tracker\\.com"
        },
        "action": {
            "type": "block"
        }
    },
    {
        "trigger": {
            "url-filter": ".*",
            "if-domain": ["example.com"]
        },
        "action": {
            "type": "css-display-none",
            "selector": ".advertisement"
        }
    }
]
`

// Compile rules
webkit.ContentRuleListStore.DefaultStore().CompileContentRuleList(
    rulesJSON,
    forIdentifier: "ContentBlockingRules",
) { ruleList, error in
    if error != nil {
        fmt.Printf("Error compiling rules: %v\n", error)
        return
    }

    // Add to content controller
    contentController := config.UserContentController()
    contentController.AddContentRuleList(ruleList)
}
```

### Remove Blocking Rules

```go
store := webkit.ContentRuleListStore.DefaultStore()

// Remove specific rule list
store.RemoveContentRuleList(forIdentifier: "ContentBlockingRules") { error in
    if error == nil {
        fmt.Println("Rules removed")
    }
}

// Get available rule lists
store.AvailableIdentifiers { identifiers in
    fmt.Printf("Available rule lists: %v\n", identifiers)
}
```

## Advanced Features

### Custom URL Schemes

```go
// Handle custom URL scheme: myapp://action

type SchemeHandler struct{}

func (h *SchemeHandler) WebViewStartURLSchemeTask(
    webView: WKWebView,
    urlSchemeTask: WKURLSchemeTask,
) {
    request := urlSchemeTask.Request()
    url := request.URL()

    // Handle request
    if url.Scheme() == "myapp" {
        // Generate response
        html := "<h1>Custom Scheme Response</h1>"
        data := []byte(html)

        response := foundation.NewHTTPURLResponse(
            url,
            200,
            "HTTP/1.1",
            map[string]string{"Content-Type": "text/html"},
        )

        urlSchemeTask.DidReceiveResponse(response)
        urlSchemeTask.DidReceiveData(data)
        urlSchemeTask.DidFinish()
    }
}

func (h *SchemeHandler) WebViewStopURLSchemeTask(
    webView: WKWebView,
    urlSchemeTask: WKURLSchemeTask,
) {
    // Cancel task
}

// Register handler
handler := &SchemeHandler{}
config.SetURLSchemeHandler(handler, forURLScheme: "myapp")
```

### Snapshot/Screenshot

```go
// Take snapshot of web view
config := webkit.NewSnapshotConfiguration()
config.SetRect(webView.Bounds()) // or specific rect

webView.TakeSnapshot(config) { image, error in
    if error == nil {
        // Save or display image
    }
}
```

### PDF Generation

```go
// Create PDF from web view (iOS 14+, macOS 11+)
config := webkit.NewPDFConfiguration()
config.SetRect(CGRect{0, 0, 612, 792}) // Letter size

webView.CreatePDF(config) { data, error in
    if error == nil {
        // Save PDF data
        data.WriteToURL(pdfURL, atomically: true)
    }
}
```

### Find in Page

```go
// Find text in page
configuration := webkit.NewFindConfiguration()
configuration.SetCaseSensitive(false)
configuration.SetBackwards(false)

webView.Find("search term", configuration) { result in
    found := result.MatchFound()
    fmt.Printf("Found: %v\n", found)
}
```

## Best Practices

### Performance

```go
// Reuse process pools
processPool := webkit.NewProcessPool()
config1.SetProcessPool(processPool)
config2.SetProcessPool(processPool)

// Suppress rendering until complete (for static content)
config.SetSuppressesIncrementalRendering(true)

// Preload content
webView.LoadRequest(request)
// Add to view later when needed
```

### Memory Management

```go
// Clean up when done
webView.StopLoading()
webView.SetNavigationDelegate(nil)
webView.SetUIDelegate(nil)
webView.RemoveFromSuperview()

// Clear data if needed
dataStore.RemoveDataOfTypesModifiedSince(allTypes, distantPast) {
    // Cleanup complete
}
```

### Security

```go
// Enforce HTTPS
// Configure App Transport Security in Info.plist

// Validate URLs before loading
func isSafeURL(url: URL) -> Bool {
    scheme := url.Scheme()
    return scheme == "https" || scheme == "file"
}

// Disable JavaScript if not needed
preferences.SetJavaScriptEnabled(false)

// Use non-persistent data store for sensitive content
config.SetWebsiteDataStore(
    webkit.NewNonPersistentWebsiteDataStore(),
)
```

## Common Patterns

### OAuth Login Flow

```go
// Load OAuth URL
authURL := foundation.NewURLWithString(
    "https://provider.com/oauth/authorize?...",
)
webView.LoadRequest(foundation.NewURLRequestWithURL(authURL))

// Intercept redirect
func WebViewDecidePolicyForNavigationAction(...) {
    url := navigationAction.Request().URL()

    if url.Scheme() == "myapp" {
        // Extract authorization code
        code := url.Query()["code"]

        // Exchange for token
        // ...

        decisionHandler(.cancel)
        return
    }

    decisionHandler(.allow)
}
```

### Hybrid App Communication

```go
// Setup bidirectional communication
// Native -> Web via JavaScript injection
// Web -> Native via message handlers

// Native calls web function
javascript := "window.updateData(\(jsonData))"
webView.EvaluateJavaScript(javascript, nil)

// Web calls native
contentController.AddScriptMessageHandler(handler, "nativeAPI")

// In JavaScript:
// window.webkit.messageHandlers.nativeAPI.postMessage({
//     method: 'getData',
//     params: {...}
// });
```

## References

- [WebKit Documentation](https://developer.apple.com/documentation/webkit)
- [WKWebView](https://developer.apple.com/documentation/webkit/wkwebview)
- [WKUserContentController](https://developer.apple.com/documentation/webkit/wkusercontentcontroller)
- [WKNavigationDelegate](https://developer.apple.com/documentation/webkit/wknavigationdelegate)
- [Content Blocking](https://developer.apple.com/documentation/safariservices/creating_a_content_blocker)
