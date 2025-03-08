// Package main implements HTML rendering functionality for the Apple docs mirror
package main

import (
	"html/template"
	"os"
)

// generateHTMLFile creates the HTML file with the file tree and embedded JavaScript
func generateHTMLFile(indexPath string, data struct {
	Root      *TreeNode
	FileCount int
	DirCount  int
	Timestamp string
}) error {
	tmpl, err := template.New("index").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Apple Documentation JSON Mirror</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
            line-height: 1.6;
            color: #333;
            max-width: 1200px;
            margin: 0 auto;
            padding: 20px;
        }
        h1, h2 {
            color: #000;
            border-bottom: 1px solid #eee;
            padding-bottom: 10px;
        }
        .tree {
            margin-top: 20px;
            flex: 0 0 350px;
            overflow-y: auto;
            max-height: calc(100vh - 200px);
        }
        .tree ul {
            padding-left: 20px;
            list-style-type: none;
        }
        .tree > ul {
            padding-left: 0;
        }
        .tree li {
            position: relative;
            padding: 4px 0;
        }
        .caret {
            cursor: pointer;
            user-select: none;
            margin-right: 5px;
            color: #888;
            display: inline-block;
            width: 16px;
            text-align: center;
        }
        .folder {
            color: #0070c9;
            font-weight: 500;
        }
        .file {
            color: #333;
        }
        .nested {
            display: none;
        }
        .active {
            display: block;
        }
        a {
            color: inherit;
            text-decoration: none;
        }
        a:hover {
            text-decoration: underline;
        }
        .search-container {
            margin: 20px 0;
        }
        #search {
            padding: 8px;
            width: 300px;
            border: 1px solid #ddd;
            border-radius: 4px;
            font-size: 14px;
        }
        .stats {
            margin: 15px 0;
            font-size: 14px;
            color: #666;
        }
        .footer {
            margin-top: 50px;
            font-size: 12px;
            color: #999;
            text-align: center;
            border-top: 1px solid #eee;
            padding-top: 20px;
        }
        .container {
            display: flex;
            gap: 30px;
        }
        .content {
            flex: 1;
            margin-top: 20px;
            overflow-y: auto;
            max-height: calc(100vh - 200px);
        }
        #jsonViewer {
            background-color: #f9f9f9;
            border: 1px solid #ddd;
            border-radius: 4px;
            padding: 15px;
            margin-top: 20px;
        }
        pre {
            white-space: pre-wrap;
            word-wrap: break-word;
            font-family: monospace;
            font-size: 14px;
        }
        .btn {
            display: inline-block;
            padding: 6px 12px;
            background-color: #0070c9;
            color: white;
            border-radius: 4px;
            text-decoration: none;
            margin-right: 8px;
            font-size: 14px;
            cursor: pointer;
            border: none;
        }
        .btn:hover {
            background-color: #005ba3;
            text-decoration: none;
        }
        .hidden {
            display: none;
        }
        #breadcrumbs {
            margin-bottom: 15px;
            font-size: 14px;
        }
        .breadcrumb-item {
            display: inline-block;
            margin-right: 5px;
        }
        .breadcrumb-item:after {
            content: '>';
            margin-left: 5px;
            color: #999;
        }
        .breadcrumb-item:last-child:after {
            content: '';
        }
        .notification {
            position: fixed;
            top: 20px;
            right: 20px;
            padding: 10px 20px;
            background-color: #4CAF50;
            color: white;
            border-radius: 4px;
            z-index: 100;
            opacity: 0;
            transition: opacity 0.3s ease;
        }
        .show-notification {
            opacity: 1;
        }
    </style>
</head>
<body>
    <h1>Apple Documentation JSON Mirror</h1>
    
    <p>This is a local mirror of Apple's developer documentation JSON files. All content is © Apple Inc.</p>
    
    <div class="stats">
        <p>Found {{.FileCount}} JSON files across {{.DirCount}} directories.</p>
    </div>

    <div class="search-container">
        <input type="text" id="search" placeholder="Search documentation...">
    </div>
    
    <div class="container">
        <div class="tree">
            {{template "tree" .Root}}
        </div>
        
        <div class="content">
            <div id="breadcrumbs"></div>
            <div id="jsonContent" class="hidden">
                <h2 id="jsonTitle">Documentation Viewer</h2>
                <div>
                    <a href="#" id="rawJsonLink" class="btn" target="_blank">View Raw JSON</a>
                    <button id="copyJsonBtn" class="btn">Copy JSON</button>
                </div>
                <div id="jsonViewer">
                    <pre id="jsonContent"></pre>
                </div>
            </div>
        </div>
    </div>
    
    <div class="notification" id="notification"></div>
    
    <div class="footer">
        <p>Generated by appledocs on {{.Timestamp}}.</p>
        <p>This is an unofficial mirror for offline use.</p>
    </div>

    <script>
        // Tree toggle functionality
        document.addEventListener('DOMContentLoaded', function() {
            var toggler = document.getElementsByClassName("caret");
            for (var i = 0; i < toggler.length; i++) {
                toggler[i].addEventListener("click", function() {
                    this.parentElement.querySelector(".nested").classList.toggle("active");
                    this.textContent = this.textContent === "▶" ? "▼" : "▶";
                });
            }

            // Expand first level by default
            var rootItems = document.querySelectorAll(".tree > ul > li > .nested");
            for (var i = 0; i < rootItems.length; i++) {
                rootItems[i].classList.add("active");
                var caret = rootItems[i].parentElement.querySelector(".caret");
                if (caret) {
                    caret.textContent = "▼";
                }
            }

            // Search functionality
            document.getElementById('search').addEventListener('input', function() {
                var searchTerm = this.value.toLowerCase();
                var fileItems = document.querySelectorAll('.file');
                var visibleItems = [];

                if (searchTerm.length < 2) {
                    // Reset visibility when search term is too short
                    resetTreeVisibility();
                    return;
                }

                // Hide all nodes first
                resetTreeVisibility(false);

                // Show matching items and their parent folders
                for (var i = 0; i < fileItems.length; i++) {
                    var item = fileItems[i];
                    var text = item.textContent.toLowerCase();
                    
                    if (text.includes(searchTerm)) {
                        // Make this item visible
                        item.style.display = 'block';
                        
                        // Make all parent ULs visible
                        var parent = item.parentElement;
                        while (parent) {
                            if (parent.classList.contains('nested')) {
                                parent.classList.add('active');
                                var caret = parent.parentElement.querySelector('.caret');
                                if (caret) caret.textContent = "▼";
                            }
                            parent = parent.parentElement;
                        }
                    }
                }
            });

            function resetTreeVisibility(visible = true) {
                // Reset all folders
                var folders = document.querySelectorAll('.folder');
                for (var i = 0; i < folders.length; i++) {
                    folders[i].style.display = visible ? 'block' : 'none';
                }
                
                // Reset all files
                var files = document.querySelectorAll('.file');
                for (var i = 0; i < files.length; i++) {
                    files[i].style.display = visible ? 'block' : 'none';
                }
                
                // Reset nested visibility
                var nestedLists = document.querySelectorAll('.nested');
                for (var i = 0; i < nestedLists.length; i++) {
                    if (visible) {
                        if (nestedLists[i].parentElement.parentElement.classList.contains('tree')) {
                            nestedLists[i].classList.add('active');
                        } else {
                            nestedLists[i].classList.remove('active');
                        }
                    } else {
                        nestedLists[i].classList.add('active');
                    }
                }
                
                // Reset carets
                var carets = document.querySelectorAll('.caret');
                for (var i = 0; i < carets.length; i++) {
                    carets[i].textContent = visible ? "▶" : "▼";
                }
                
                // Keep top level expanded
                if (visible) {
                    var rootItems = document.querySelectorAll(".tree > ul > li > .nested");
                    for (var i = 0; i < rootItems.length; i++) {
                        rootItems[i].classList.add("active");
                        var caret = rootItems[i].parentElement.querySelector(".caret");
                        if (caret) {
                            caret.textContent = "▼";
                        }
                    }
                }
            }
            
            // JSON Viewer functionality
            const jsonContent = document.getElementById('jsonContent');
            const jsonViewer = document.getElementById('jsonViewer').querySelector('pre');
            const jsonTitle = document.getElementById('jsonTitle');
            const rawJsonLink = document.getElementById('rawJsonLink');
            const copyJsonBtn = document.getElementById('copyJsonBtn');
            const breadcrumbs = document.getElementById('breadcrumbs');
            const notification = document.getElementById('notification');
            
            // Add click event to JSON links
            document.querySelectorAll('.json-link').forEach(link => {
                link.addEventListener('click', function(e) {
                    e.preventDefault();
                    const path = this.getAttribute('data-path');
                    loadJSON(path);
                });
            });
            
            // Copy JSON button
            copyJsonBtn.addEventListener('click', function() {
                const jsonText = jsonViewer.textContent;
                navigator.clipboard.writeText(jsonText)
                    .then(() => {
                        showNotification('JSON copied to clipboard!');
                    })
                    .catch(err => {
                        console.error('Failed to copy: ', err);
                        showNotification('Failed to copy to clipboard', true);
                    });
            });
            
            // Show notification
            function showNotification(message, isError = false) {
                notification.textContent = message;
                notification.style.backgroundColor = isError ? '#f44336' : '#4CAF50';
                notification.classList.add('show-notification');
                setTimeout(() => {
                    notification.classList.remove('show-notification');
                }, 2000);
            }
            
            // Load and display a JSON file
            function loadJSON(path) {
                fetch(path)
                    .then(response => {
                        if (!response.ok) {
                            throw new Error('Network response was not ok');
                        }
                        return response.json();
                    })
                    .then(data => {
                        // Show the JSON viewer
                        jsonContent.classList.remove('hidden');
                        
                        // Format JSON for display
                        const formattedJSON = JSON.stringify(data, null, 2);
                        jsonViewer.textContent = formattedJSON;
                        
                        // Update title and link
                        const pathParts = path.split('/');
                        const fileName = pathParts[pathParts.length - 1];
                        jsonTitle.textContent = fileName;
                        rawJsonLink.href = path;
                        
                        // Update breadcrumbs
                        updateBreadcrumbs(path);
                    })
                    .catch(error => {
                        console.error('Error fetching JSON:', error);
                        jsonViewer.textContent = 'Error loading JSON: ' + error.message;
                        jsonContent.classList.remove('hidden');
                        jsonTitle.textContent = 'Error';
                    });
            }
            
            // Update breadcrumbs navigation
            function updateBreadcrumbs(path) {
                const parts = path.split('/');
                breadcrumbs.innerHTML = '';
                
                // Create Home link
                const homeLink = document.createElement('span');
                homeLink.className = 'breadcrumb-item';
                homeLink.innerHTML = '<a href="javascript:void(0)">Home</a>';
                homeLink.querySelector('a').addEventListener('click', function() {
                    jsonContent.classList.add('hidden');
                });
                breadcrumbs.appendChild(homeLink);
                
                // Create path links
                let currentPath = '';
                parts.forEach((part, index) => {
                    if (!part) return;
                    
                    currentPath += '/' + part;
                    const item = document.createElement('span');
                    item.className = 'breadcrumb-item';
                    
                    if (index === parts.length - 1) {
                        // Last item (current file)
                        item.textContent = part;
                    } else {
                        item.textContent = part;
                    }
                    
                    breadcrumbs.appendChild(item);
                });
            }
        });
    </script>
</body>
</html>

{{define "tree"}}
<ul class="{{if ne .Name "root"}}nested{{end}}">
    {{range .Children}}
        <li>
            {{if .IsDir}}
                <span class="caret">▶</span>
                <span class="folder">{{.Name}}</span>
                {{template "tree" .}}
            {{else}}
                <span class="file">
                    <a href="javascript:void(0)" class="json-link" data-path="{{.Path}}">{{.Name}}</a>
                </span>
            {{end}}
        </li>
    {{end}}
</ul>
{{end}}`)

	if err != nil {
		return err
	}

	file, err := os.Create(indexPath)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := tmpl.Execute(file, data); err != nil {
		return err
	}

	return nil
}