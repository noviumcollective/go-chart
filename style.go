package chart

import (
	"fmt"
	"strings"

	"github.com/golang/freetype/truetype"
	"github.com/wcharczuk/go-chart/drawing"
)

// Style is a simple style set.
type Style struct {
	Show    bool
	Padding Box

	StrokeWidth     float64
	StrokeColor     drawing.Color
	StrokeDashArray []float64

	FillColor drawing.Color
	FontSize  float64
	FontColor drawing.Color
	Font      *truetype.Font

	TextHorizontalAlign textHorizontalAlign
	TextVerticalAlign   textVerticalAlign
	TextWrap            textWrap
}

// IsZero returns if the object is set or not.
func (s Style) IsZero() bool {
	return s.StrokeColor.IsZero() && s.FillColor.IsZero() && s.StrokeWidth == 0 && s.FontColor.IsZero() && s.FontSize == 0 && s.Font == nil
}

// String returns a text representation of the style.
func (s Style) String() string {
	if s.IsZero() {
		return "{}"
	}

	var output []string
	if s.Show {
		output = []string{"\"show\": true"}
	} else {
		output = []string{"\"show\": false"}
	}

	if !s.Padding.IsZero() {
		output = append(output, fmt.Sprintf("\"padding\": %s", s.Padding.String()))
	} else {
		output = append(output, "\"padding\": null")
	}

	if s.StrokeWidth >= 0 {
		output = append(output, fmt.Sprintf("\"stroke_width\": %0.2f", s.StrokeWidth))
	} else {
		output = append(output, "\"stroke_width\": null")
	}

	if !s.StrokeColor.IsZero() {
		output = append(output, fmt.Sprintf("\"stroke_color\": %s", s.StrokeColor.String()))
	} else {
		output = append(output, "\"stroke_color\": null")
	}

	if len(s.StrokeDashArray) > 0 {
		var elements []string
		for _, v := range s.StrokeDashArray {
			elements = append(elements, fmt.Sprintf("%.2f", v))
		}
		dashArray := strings.Join(elements, ", ")
		output = append(output, fmt.Sprintf("\"stroke_dash_array\": [%s]", dashArray))
	} else {
		output = append(output, "\"stroke_dash_array\": null")
	}

	if !s.FillColor.IsZero() {
		output = append(output, fmt.Sprintf("\"fill_color\": %s", s.FillColor.String()))
	} else {
		output = append(output, "\"fill_color\": null")
	}

	if s.FontSize != 0 {
		output = append(output, fmt.Sprintf("\"font_size\": \"%0.2fpt\"", s.FontSize))
	} else {
		output = append(output, "\"font_size\": null")
	}

	if !s.FontColor.IsZero() {
		output = append(output, fmt.Sprintf("\"font_color\": %s", s.FontColor.String()))
	} else {
		output = append(output, "\"font_color\": null")
	}

	if s.Font != nil {
		output = append(output, fmt.Sprintf("\"font\": \"%s\"", s.Font.Name(truetype.NameIDFontFamily)))
	} else {
		output = append(output, "\"font_color\": null")
	}

	return "{" + strings.Join(output, ", ") + "}"
}

// GetStrokeColor returns the stroke color.
func (s Style) GetStrokeColor(defaults ...drawing.Color) drawing.Color {
	if s.StrokeColor.IsZero() {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return drawing.ColorTransparent
	}
	return s.StrokeColor
}

// GetFillColor returns the fill color.
func (s Style) GetFillColor(defaults ...drawing.Color) drawing.Color {
	if s.FillColor.IsZero() {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return drawing.ColorTransparent
	}
	return s.FillColor
}

// GetStrokeWidth returns the stroke width.
func (s Style) GetStrokeWidth(defaults ...float64) float64 {
	if s.StrokeWidth == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultStrokeWidth
	}
	return s.StrokeWidth
}

// GetStrokeDashArray returns the stroke dash array.
func (s Style) GetStrokeDashArray(defaults ...[]float64) []float64 {
	if len(s.StrokeDashArray) == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return nil
	}
	return s.StrokeDashArray
}

// GetFontSize gets the font size.
func (s Style) GetFontSize(defaults ...float64) float64 {
	if s.FontSize == 0 {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return DefaultFontSize
	}
	return s.FontSize
}

// GetFontColor gets the font size.
func (s Style) GetFontColor(defaults ...drawing.Color) drawing.Color {
	if s.FontColor.IsZero() {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return drawing.ColorTransparent
	}
	return s.FontColor
}

// GetFont returns the font face.
func (s Style) GetFont(defaults ...*truetype.Font) *truetype.Font {
	if s.Font == nil {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return nil
	}
	return s.Font
}

// GetPadding returns the padding.
func (s Style) GetPadding(defaults ...Box) Box {
	if s.Padding.IsZero() {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return Box{}
	}
	return s.Padding
}

// GetTextHorizontalAlign returns the horizontal alignment.
func (s Style) GetTextHorizontalAlign(defaults ...textHorizontalAlign) textHorizontalAlign {
	if s.TextHorizontalAlign == TextHorizontalAlignUnset {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return TextHorizontalAlignLeft
	}
	return s.TextHorizontalAlign
}

// GetTextVerticalAlign returns the vertical alignment.
func (s Style) GetTextVerticalAlign(defaults ...textVerticalAlign) textVerticalAlign {
	if s.TextVerticalAlign == TextVerticalAlignUnset {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return TextVerticalAlignBaseline
	}
	return s.TextVerticalAlign
}

// GetTextWrap returns the word wrap.
func (s Style) GetTextWrap(defaults ...textWrap) textWrap {
	if s.TextWrap == TextWrapUnset {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return TextWrapWord
	}
	return s.TextWrap
}

// WriteToRenderer passes the style's options to a renderer.
func (s Style) WriteToRenderer(r Renderer) {
	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeWidth(s.GetStrokeWidth())
	r.SetStrokeDashArray(s.GetStrokeDashArray())
	r.SetFillColor(s.GetFillColor())
	r.SetFont(s.GetFont())
	r.SetFontColor(s.GetFontColor())
	r.SetFontSize(s.GetFontSize())
}

// WriteDrawingOptionsToRenderer passes just the drawing style options to a renderer.
func (s Style) WriteDrawingOptionsToRenderer(r Renderer) {
	r.SetStrokeColor(s.GetStrokeColor())
	r.SetStrokeWidth(s.GetStrokeWidth())
	r.SetStrokeDashArray(s.GetStrokeDashArray())
	r.SetFillColor(s.GetFillColor())
}

// WriteTextOptionsToRenderer passes just the text style options to a renderer.
func (s Style) WriteTextOptionsToRenderer(r Renderer) {
	r.SetFont(s.GetFont())
	r.SetFontColor(s.GetFontColor())
	r.SetFontSize(s.GetFontSize())
}

// InheritFrom coalesces two styles into a new style.
func (s Style) InheritFrom(defaults Style) (final Style) {
	final.StrokeColor = s.GetStrokeColor(defaults.StrokeColor)
	final.StrokeWidth = s.GetStrokeWidth(defaults.StrokeWidth)
	final.StrokeDashArray = s.GetStrokeDashArray(defaults.StrokeDashArray)
	final.FillColor = s.GetFillColor(defaults.FillColor)
	final.FontColor = s.GetFontColor(defaults.FontColor)
	final.FontSize = s.GetFontSize(defaults.FontSize)
	final.Font = s.GetFont(defaults.Font)
	final.Padding = s.GetPadding(defaults.Padding)
	final.TextHorizontalAlign = s.GetTextHorizontalAlign(defaults.TextHorizontalAlign)
	final.TextVerticalAlign = s.GetTextVerticalAlign(defaults.TextVerticalAlign)
	final.TextWrap = s.GetTextWrap(defaults.TextWrap)
	return
}

// GetStrokeOptions returns the stroke components.
func (s Style) GetStrokeOptions() Style {
	return Style{
		StrokeDashArray: s.StrokeDashArray,
		StrokeColor:     s.StrokeColor,
		StrokeWidth:     s.StrokeWidth,
	}
}

// GetFillOptions returns the fill components.
func (s Style) GetFillOptions() Style {
	return Style{
		FillColor: s.FillColor,
	}
}

// GetFillAndStrokeOptions returns the fill and stroke components.
func (s Style) GetFillAndStrokeOptions() Style {
	return Style{
		StrokeDashArray: s.StrokeDashArray,
		FillColor:       s.FillColor,
		StrokeColor:     s.StrokeColor,
		StrokeWidth:     s.StrokeWidth,
	}
}

// GetTextOptions returns just the text components of the style.
func (s Style) GetTextOptions() Style {
	return Style{
		FontColor:           s.FontColor,
		FontSize:            s.FontSize,
		Font:                s.Font,
		TextHorizontalAlign: s.TextHorizontalAlign,
		TextVerticalAlign:   s.TextVerticalAlign,
		TextWrap:            s.TextWrap,
	}
}
