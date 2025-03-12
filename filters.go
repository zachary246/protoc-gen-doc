package gendoc

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

var (
	paraPattern         = regexp.MustCompile(`(\n|\r|\r\n)\s*`)
	spacePattern        = regexp.MustCompile("( )+")
	multiNewlinePattern = regexp.MustCompile(`(\r\n|\r|\n){2,}`)
	specialCharsPattern = regexp.MustCompile(`[^a-zA-Z0-9_-]`)
)

// PFilter splits the content by new lines and wraps each one in a <p> tag.
func PFilter(content string) template.HTML {
	paragraphs := paraPattern.Split(content, -1)

	return template.HTML(fmt.Sprintf("<p>%s</p>", strings.Join(paragraphs, "</p><p>")))
}

// ParaFilter splits the content by new lines and wraps each one in a <para> tag.
func ParaFilter(content string) string {
	paragraphs := paraPattern.Split(content, -1)
	return fmt.Sprintf("<para>%s</para>", strings.Join(paragraphs, "</para><para>"))
}

// NoBrFilter processes line breaks in content.
// It preserves both single and multiple line breaks while normalizing whitespace.
func NoBrFilter(content string) string {
	// First normalize all line endings to \n
	normalized := strings.Replace(content, "\r\n", "\n", -1)
	normalized = strings.Replace(normalized, "\r", "\n", -1)
	
	// Split content into lines, preserving all line breaks
	lines := strings.Split(normalized, "\n")
	
	// Process each line to clean up extra spaces
	for i, line := range lines {
		// Trim spaces and replace multiple spaces with single space
		lines[i] = spacePattern.ReplaceAllString(strings.TrimSpace(line), " ")
	}
	
	// Join lines back together, preserving all line breaks
	return strings.Join(lines, "\n")
}

// BrFilter converts line breaks to HTML br tags while preserving all line breaks
func BrFilter(content string) template.HTML {
	processed := NoBrFilter(content)
	// Replace all line breaks with <br/> tags
	return template.HTML(strings.Replace(processed, "\n", "<br/>", -1))
}

// AnchorFilter replaces all special characters with URL friendly dashes
func AnchorFilter(str string) string {
	return specialCharsPattern.ReplaceAllString(strings.ReplaceAll(str, "/", "_"), "-")
}
