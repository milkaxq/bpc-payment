package utils

import "golang.org/x/net/html"

func FindRequestId(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "input" {
		// Check if this input element matches the desired criteria
		isHiddenInput := false
		var id, name, value string

		// Extract attributes of the input element
		for _, attr := range n.Attr {
			if attr.Key == "type" && attr.Val == "hidden" {
				isHiddenInput = true
			} else if attr.Key == "id" {
				id = attr.Val
			} else if attr.Key == "name" {
				name = attr.Val
			} else if attr.Key == "value" {
				value = attr.Val
			}
		}

		// Check if it's the desired input element
		if isHiddenInput && id == "request_id" && name == "request_id" {
			return value
		}
	}

	// Recursive traversal of child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if value := FindRequestId(c); value != "" {
			return value
		}
	}

	return ""
}

func FindPaRes(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "input" {
		// Check if this input element matches the desired criteria
		isHiddenInput := false
		var name, value string

		// Extract attributes of the input element
		for _, attr := range n.Attr {
			if attr.Key == "type" && attr.Val == "hidden" {
				isHiddenInput = true
			} else if attr.Key == "name" {
				name = attr.Val
			} else if attr.Key == "value" {
				value = attr.Val
			}
		}

		// Check if it's the desired input element
		if isHiddenInput && name == "PaRes" {
			return value
		}
	}

	// Recursive traversal of child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if value := FindPaRes(c); value != "" {
			return value
		}
	}

	return ""
}
