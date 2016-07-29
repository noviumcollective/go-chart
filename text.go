package chart

import "strings"

// TextHorizontalAlign is an enum for the horizontal alignment options.
type textHorizontalAlign int

const (
	// TextHorizontalAlignUnset is the unset state for text horizontal alignment.
	TextHorizontalAlignUnset textHorizontalAlign = 0
	// TextHorizontalAlignLeft aligns a string horizontally so that it's left ligature starts at horizontal pixel 0.
	TextHorizontalAlignLeft textHorizontalAlign = 1
	// TextHorizontalAlignCenter left aligns a string horizontally so that there are equal pixels
	// to the left and to the right of a string within a box.
	TextHorizontalAlignCenter textHorizontalAlign = 2
	// TextHorizontalAlignRight right aligns a string horizontally so that the right ligature ends at the right-most pixel
	// of a box.
	TextHorizontalAlignRight textHorizontalAlign = 3
)

// TextWrap is an enum for the word wrap options.
type textWrap int

const (
	// TextWrapUnset is the unset state for text wrap options.
	TextWrapUnset textWrap = 0
	// TextWrapNone will spill text past horizontal boundaries.
	TextWrapNone textWrap = 1
	// TextWrapWord will split a string on words (i.e. spaces) to fit within a horizontal boundary.
	TextWrapWord textWrap = 2
	// TextWrapRune will split a string on a rune (i.e. utf-8 codepage) to fit within a horizontal boundary.
	TextWrapRune textWrap = 3
)

// TextVerticalAlign is an enum for the vertical alignment options.
type textVerticalAlign int

const (
	// TextVerticalAlignUnset is the unset state for vertical alignment options.
	TextVerticalAlignUnset textVerticalAlign = 0
	// TextVerticalAlignBaseline aligns text according to the "baseline" of the string, or where a normal ascender begins.
	TextVerticalAlignBaseline textVerticalAlign = 1
	// TextVerticalAlignBottom aligns the text according to the lowers pixel of any of the ligatures (ex. g or q both extend below the baseline).
	TextVerticalAlignBottom textVerticalAlign = 2
	// TextVerticalAlignMiddle aligns the text so that there is an equal amount of space above and below the top and bottom of the ligatures.
	TextVerticalAlignMiddle textVerticalAlign = 3
	// TextVerticalAlignMiddleBaseline aligns the text veritcally so that there is an equal number of pixels above and below the baseline of the string.
	TextVerticalAlignMiddleBaseline textVerticalAlign = 4
	// TextVerticalAlignTop alignts the text so that the top of the ligatures are at y-pixel 0 in the container.
	TextVerticalAlignTop textVerticalAlign = 5
)

var (
	// Text contains utilities for text.
	Text = &text{}
)

// TextStyle encapsulates text style options.
type TextStyle struct {
	HorizontalAlign textHorizontalAlign
	VerticalAlign   textVerticalAlign
	Wrap            textWrap
}

type text struct{}

func (t text) WrapFit(r Renderer, value string, width int, style Style, wrapOption textWrap) []string {
	valueBox := r.MeasureText(value)
	if valueBox.Width() > width {
		switch wrapOption {
		case TextWrapRune:
			return t.WrapFitRune(r, value, width, style)
		case TextWrapWord:
			return t.WrapFitWord(r, value, width, style)
		}
	}
	return []string{value}
}

func (t text) WrapFitWord(r Renderer, value string, width int, style Style) []string {
	style.WriteToRenderer(r)

	var output []string
	var line string
	var word string

	var textBox Box

	for _, c := range value {
		if c == rune('\n') { // commit the line to output
			output = append(output, t.Trim(line+word))
			line = ""
			word = ""
			continue
		}

		textBox = r.MeasureText(line + word + string(c))

		if textBox.Width() >= width {
			output = append(output, t.Trim(line))
			line = word
			word = string(c)
			continue
		}

		if c == rune(' ') || c == rune('\t') {
			line = line + word + string(c)
			word = ""
			continue
		}
		word = word + string(c)
	}

	return append(output, t.Trim(line+word))
}

func (t text) WrapFitRune(r Renderer, value string, width int, style Style) []string {
	style.WriteToRenderer(r)

	var output []string
	var line string
	var textBox Box
	for _, c := range value {
		if c == rune('\n') {
			output = append(output, line)
			line = ""
			continue
		}

		textBox = r.MeasureText(line + string(c))

		if textBox.Width() >= width {
			output = append(output, line)
			line = string(c)
			continue
		}
		line = line + string(c)
	}
	return t.appendLast(output, line)
}

func (t text) Trim(value string) string {
	return strings.Trim(value, " \t\n\r")
}

func (t text) appendLast(lines []string, text string) []string {
	if len(lines) == 0 {
		return []string{text}
	}
	lastLine := lines[len(lines)-1]
	lines[len(lines)-1] = lastLine + text
	return lines
}
