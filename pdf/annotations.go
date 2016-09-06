/*
 * This file is subject to the terms and conditions defined in
 * file 'LICENSE.md', which is part of this source code package.
 */

package pdf

import (
	"fmt"
	"github.com/unidoc/unidoc/common"
)

//type PdfAnnotation interface{}

type PdfAnnotation struct {
	context      interface{}
	Type         PdfObject // Annot
	Subtype      PdfObject
	Rect         PdfObject
	Contents     PdfObject
	P            PdfObject // Reference to page object.
	NM           PdfObject
	M            PdfObject
	F            PdfObject
	AP           PdfObject
	AS           PdfObject
	Border       PdfObject
	C            PdfObject
	StructParent PdfObject
	OC           PdfObject
}

/*
Subtype
*/

// Subtype: Text
type PdfAnnotationText struct {
	PdfAnnotation
	Open       PdfObject
	Name       PdfObject
	State      PdfObject
	StateModel PdfObject
}

// Subtype: Link
type PdfAnnotationLink struct {
	PdfAnnotation
	A          PdfObject
	Dest       PdfObject
	H          PdfObject
	PA         PdfObject
	QuadPoints PdfObject
	BS         PdfObject
}

// Subtype: FreeText
type PdfAnnotationFreeText struct {
	PdfAnnotation
	DA PdfObject
	Q  PdfObject
	RC PdfObject
	DS PdfObject
	CL PdfObject
	IT PdfObject
	BE PdfObject
	RD PdfObject
	BS PdfObject
	LE PdfObject
}

// Subtype: Line
type PdfAnnotationLine struct {
	PdfAnnotation
	L       PdfObject
	BS      PdfObject
	LE      PdfObject
	IC      PdfObject
	LL      PdfObject
	LLE     PdfObject
	Cap     PdfObject
	IT      PdfObject
	LLO     PdfObject
	CP      PdfObject
	Measure PdfObject
	CO      PdfObject
}

// Subtype: Square
type PdfAnnotationSquare struct {
	PdfAnnotation
	BS PdfObject
	IC PdfObject
	BE PdfObject
	RD PdfObject
}

// Subtype: Circle
type PdfAnnotationCircle struct {
	PdfAnnotation
	BS PdfObject
	IC PdfObject
	BE PdfObject
	RD PdfObject
}

// Subtype: Polygon
type PdfAnnotationPolygon struct {
	PdfAnnotation
	Vertices PdfObject
	LE       PdfObject
	BS       PdfObject
	IC       PdfObject
	BE       PdfObject
	IT       PdfObject
	Measure  PdfObject
}

// Subtype: PolyLine
type PdfAnnotationPolyLine struct {
	PdfAnnotation
	Vertices PdfObject
	LE       PdfObject
	BS       PdfObject
	IC       PdfObject
	BE       PdfObject
	IT       PdfObject
	Measure  PdfObject
}

// Subtype: Highlight
type PdfAnnotationHighlight struct {
	PdfAnnotation
	QuadPoints PdfObject
}

// Subtype: Underline
type PdfAnnotationUnderline struct {
	PdfAnnotation
	QuadPoints PdfObject
}

// Subtype: Squiggly
type PdfAnnotationSquiggly struct {
	PdfAnnotation
	QuadPoints PdfObject
}

// Subtype: StrikeOut
type PdfAnnotationStrikeOut struct {
	PdfAnnotation
	QuadPoints PdfObject
}

// Subtype: Caret
type PdfAnnotationCaret struct {
	PdfAnnotation
	RD PdfObject
	Sy PdfObject
}

// Subtype: Stamp
type PdfAnnotationStamp struct {
	PdfAnnotation
	Name PdfObject
}

// Subtype: Ink
type PdfAnnotationInk struct {
	PdfAnnotation
	InkList PdfObject
	BS      PdfObject
}

// Subtype: Popup
type PdfAnnotationPopup struct {
	PdfAnnotation
	Parent PdfObject
	Open   PdfObject
}

// Subtype: FileAttachment
type PdfAnnotationFileAttachment struct {
	PdfAnnotation
	FS   PdfObject
	Name PdfObject
}

// Subtype: Sound
type PdfAnnotationSound struct {
	PdfAnnotation
	Sound PdfObject
	Name  PdfObject
}

// Subtype: Movie
type PdfAnnotationMovie struct {
	PdfAnnotation
	T     PdfObject
	Movie PdfObject
	A     PdfObject
}

// Subtype: Screen
type PdfAnnotationScreen struct {
	PdfAnnotation
	T  PdfObject
	MK PdfObject
	A  PdfObject
	AA PdfObject
}

// Subtype: Widget
type PdfAnnotationWidget struct {
	PdfAnnotation
	H      PdfObject
	MK     PdfObject
	A      PdfObject
	AA     PdfObject
	BS     PdfObject
	Parent PdfObject
}

// Subtype: Watermark
type PdfAnnotationWatermark struct {
	PdfAnnotation
	FixedPrint PdfObject
}

// Subtype: PrinterMark
type PdfAnnotationPrinterMark struct {
	PdfAnnotation
	MN PdfObject
}

// Subtype: TrapNet
type PdfAnnotationTrapNet struct {
	PdfAnnotation
}

// Subtype: 3D
type PdfAnnotation3D struct {
	PdfAnnotation
	T3DD PdfObject
	T3DV PdfObject
	T3DA PdfObject
	T3DI PdfObject
	T3DB PdfObject
}

type PdfAnnotationRedact struct {
	PdfAnnotation
	QuadPoints  PdfObject
	IC          PdfObject
	RO          PdfObject
	OverlayText PdfObject
	Repeat      PdfObject
	DA          PdfObject
	Q           PdfObject
}

func (this *PdfAnnotation) ToPdfObject() PdfObject {
	switch t := this.context.(type) {
	case *PdfAnnotationText:
		return t.ToPdfObject()
	case *PdfAnnotationLink:
		return t.ToPdfObject()
	case *PdfAnnotationFreeText:
		return t.ToPdfObject()
	case *PdfAnnotationLine:
		return t.ToPdfObject()
	case *PdfAnnotationSquare:
		return t.ToPdfObject()
	case *PdfAnnotationCircle:
		return t.ToPdfObject()
	case *PdfAnnotationPolygon:
		return t.ToPdfObject()
	case *PdfAnnotationPolyLine:
		return t.ToPdfObject()
	case *PdfAnnotationHighlight:
		return t.ToPdfObject()
	case *PdfAnnotationCaret:
		return t.ToPdfObject()
	case *PdfAnnotationStamp:
		return t.ToPdfObject()
	case *PdfAnnotationInk:
		return t.ToPdfObject()
	case *PdfAnnotationPopup:
		return t.ToPdfObject()
	case *PdfAnnotationFileAttachment:
		return t.ToPdfObject()
	case *PdfAnnotationSound:
		return t.ToPdfObject()
	case *PdfAnnotationMovie:
		return t.ToPdfObject()
	case *PdfAnnotationScreen:
		return t.ToPdfObject()
	case *PdfAnnotationWidget:
		return t.ToPdfObject()
	case *PdfAnnotationPrinterMark:
		return t.ToPdfObject()
	case *PdfAnnotationTrapNet:
		return t.ToPdfObject()
	case *PdfAnnotationWatermark:
		return t.ToPdfObject()
	case *PdfAnnotation3D:
		return t.ToPdfObject()
	case *PdfAnnotationRedact:
		return t.ToPdfObject()
	}
	common.Log.Error("Unknown type %T", this.context)
	return nil
}

func (r *PdfReader) newPdfAnnotationFromDict(d PdfObjectDictionary) (*PdfAnnotation, error) {
	annot := PdfAnnotation{}

	if obj, has := d["Type"]; has {
		str, ok := obj.(*PdfObjectString)
		if !ok {
			return nil, fmt.Errorf("Invalid type of Type (%T)", obj)
		}
		if *str != "Annot" {
			// Log a debug message.
			// Not returning an error on this.
			common.Log.Debug("Unsuspected Type != Annot (%s)", *str)
		}
	}

	if obj, has := d["Rect"]; has {
		annot.Rect = obj
	}

	if obj, has := d["Contents"]; has {
		annot.Contents = obj
	}

	if obj, has := d["P"]; has {
		annot.P = obj
	}

	if obj, has := d["NM"]; has {
		annot.NM = obj
	}

	if obj, has := d["M"]; has {
		annot.M = obj
	}

	if obj, has := d["F"]; has {
		annot.F = obj
	}

	if obj, has := d["AP"]; has {
		annot.AP = obj
	}

	if obj, has := d["AS"]; has {
		annot.AS = obj
	}

	if obj, has := d["Border"]; has {
		annot.Border = obj
	}

	if obj, has := d["C"]; has {
		annot.C = obj
	}

	if obj, has := d["StructParent"]; has {
		annot.StructParent = obj
	}

	if obj, has := d["OC"]; has {
		annot.OC = obj
	}

	subtypeObj, has := d["Subtype"]
	if !has {
		return nil, fmt.Errorf("Missing Subtype")
	}
	subtype, ok := subtypeObj.(*PdfObjectName)
	if !ok {
		return nil, fmt.Errorf("Invalid Subtype object type != name (%T)", subtypeObj)
	}
	switch *subtype {
	case "Text":
		ctx, err := newPdfAnnotationTextFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Link":
		ctx, err := newPdfAnnotationLinkFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "FreeText":
		ctx, err := newPdfAnnotationFreeTextFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Line":
		ctx, err := newPdfAnnotationLineFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Square":
		ctx, err := newPdfAnnotationSquareFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Circle":
		ctx, err := newPdfAnnotationCircleFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Polygon":
		ctx, err := newPdfAnnotationPolygonFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "PolyLine":
		ctx, err := newPdfAnnotationPolyLineFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Highlight":
		ctx, err := newPdfAnnotationHighlightFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Caret":
		ctx, err := newPdfAnnotationCaretFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Stamp":
		ctx, err := newPdfAnnotationStampFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Ink":
		ctx, err := newPdfAnnotationInkFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Popup":
		ctx, err := newPdfAnnotationPopupFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "FileAttachment":
		ctx, err := newPdfAnnotationFileAttachmentFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Sound":
		ctx, err := newPdfAnnotationSoundFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Movie":
		ctx, err := newPdfAnnotationMovieFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Screen":
		ctx, err := newPdfAnnotationScreenFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Widget":
		ctx, err := newPdfAnnotationWidgetFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "PrinterMark":
		ctx, err := newPdfAnnotationPrinterMarkFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "TrapNet":
		ctx, err := newPdfAnnotationTrapNetFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Watermark":
		ctx, err := newPdfAnnotationWatermarkFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "3D":
		ctx, err := newPdfAnnotation3DFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	case "Redact":
		ctx, err := newPdfAnnotationRedactFromDict(d)
		if err != nil {
			return nil, err
		}
		ctx.PdfAnnotation = annot
		annot.context = ctx
		return &annot, nil
	}

	err := fmt.Errorf("Unknown annotation (%s)", *subtype)
	return nil, err
}

func newPdfAnnotationTextFromDict(d PdfObjectDictionary) (*PdfAnnotationText, error) {
	annot := PdfAnnotationText{}

	if obj, has := d["Open"]; has {
		annot.Open = obj
	}

	if obj, has := d["Name"]; has {
		annot.Name = obj
	}

	if obj, has := d["State"]; has {
		annot.State = obj
	}

	if obj, has := d["StateModel"]; has {
		annot.StateModel = obj
	}

	return &annot, nil
}

func newPdfAnnotationLinkFromDict(d PdfObjectDictionary) (*PdfAnnotationLink, error) {
	annot := PdfAnnotationLink{}

	if obj, has := d["A"]; has {
		annot.A = obj
	}
	if obj, has := d["Dest"]; has {
		annot.Dest = obj
	}
	if obj, has := d["H"]; has {
		annot.H = obj
	}
	if obj, has := d["PA"]; has {
		annot.PA = obj
	}
	if obj, has := d["QuadPoints"]; has {
		annot.QuadPoints = obj
	}
	if obj, has := d["BS"]; has {
		annot.BS = obj
	}

	return &annot, nil
}

func newPdfAnnotationFreeTextFromDict(d PdfObjectDictionary) (*PdfAnnotationFreeText, error) {
	annot := PdfAnnotationFreeText{}

	if obj, has := d["DA"]; has {
		annot.DA = obj
	}
	if obj, has := d["Q"]; has {
		annot.Q = obj
	}
	if obj, has := d["RC"]; has {
		annot.RC = obj
	}
	if obj, has := d["DS"]; has {
		annot.DS = obj
	}
	if obj, has := d["CL"]; has {
		annot.CL = obj
	}
	if obj, has := d["IT"]; has {
		annot.IT = obj
	}
	if obj, has := d["BE"]; has {
		annot.BE = obj
	}
	if obj, has := d["RD"]; has {
		annot.RD = obj
	}
	if obj, has := d["BS"]; has {
		annot.BS = obj
	}
	if obj, has := d["LE"]; has {
		annot.LE = obj
	}

	return &annot, nil
}

func newPdfAnnotationLineFromDict(d PdfObjectDictionary) (*PdfAnnotationLine, error) {
	annot := PdfAnnotationLine{}

	if obj, has := d["L"]; has {
		annot.L = obj
	}
	if obj, has := d["BS"]; has {
		annot.BS = obj
	}
	if obj, has := d["LE"]; has {
		annot.LE = obj
	}
	if obj, has := d["IC"]; has {
		annot.IC = obj
	}
	if obj, has := d["LL"]; has {
		annot.LL = obj
	}
	if obj, has := d["LLE"]; has {
		annot.LLE = obj
	}
	if obj, has := d["Cap"]; has {
		annot.Cap = obj
	}
	if obj, has := d["IT"]; has {
		annot.IT = obj
	}
	if obj, has := d["LLO"]; has {
		annot.LLO = obj
	}
	if obj, has := d["CP"]; has {
		annot.CP = obj
	}
	if obj, has := d["Measure"]; has {
		annot.Measure = obj
	}
	if obj, has := d["CO"]; has {
		annot.CO = obj
	}

	return &annot, nil
}

func newPdfAnnotationSquareFromDict(d PdfObjectDictionary) (*PdfAnnotationSquare, error) {
	annot := PdfAnnotationSquare{}

	if obj, has := d["BS"]; has {
		annot.BS = obj
	}
	if obj, has := d["IC"]; has {
		annot.IC = obj
	}
	if obj, has := d["BE"]; has {
		annot.BE = obj
	}
	if obj, has := d["RD"]; has {
		annot.RD = obj
	}

	return &annot, nil
}

func newPdfAnnotationCircleFromDict(d PdfObjectDictionary) (*PdfAnnotationCircle, error) {
	annot := PdfAnnotationCircle{}

	if obj, has := d["BS"]; has {
		annot.BS = obj
	}
	if obj, has := d["IC"]; has {
		annot.IC = obj
	}
	if obj, has := d["BE"]; has {
		annot.BE = obj
	}
	if obj, has := d["RD"]; has {
		annot.RD = obj
	}

	return &annot, nil
}

func newPdfAnnotationPolygonFromDict(d PdfObjectDictionary) (*PdfAnnotationPolygon, error) {
	annot := PdfAnnotationPolygon{}
	if obj, has := d["Vertices"]; has {
		annot.Vertices = obj
	}
	if obj, has := d["LE"]; has {
		annot.LE = obj
	}
	if obj, has := d["BS"]; has {
		annot.BS = obj
	}
	if obj, has := d["IC"]; has {
		annot.IC = obj
	}
	if obj, has := d["BE"]; has {
		annot.BE = obj
	}
	if obj, has := d["IT"]; has {
		annot.IT = obj
	}
	if obj, has := d["Measure"]; has {
		annot.Measure = obj
	}

	return &annot, nil
}

func newPdfAnnotationPolyLineFromDict(d PdfObjectDictionary) (*PdfAnnotationPolyLine, error) {
	annot := PdfAnnotationPolyLine{}
	if obj, has := d["Vertices"]; has {
		annot.Vertices = obj
	}
	if obj, has := d["LE"]; has {
		annot.LE = obj
	}
	if obj, has := d["BS"]; has {
		annot.BS = obj
	}
	if obj, has := d["IC"]; has {
		annot.IC = obj
	}
	if obj, has := d["BE"]; has {
		annot.BE = obj
	}
	if obj, has := d["IT"]; has {
		annot.IT = obj
	}
	if obj, has := d["Measure"]; has {
		annot.Measure = obj
	}

	return &annot, nil
}

func newPdfAnnotationHighlightFromDict(d PdfObjectDictionary) (*PdfAnnotationHighlight, error) {
	annot := PdfAnnotationHighlight{}

	if obj, has := d["QuadPoints"]; has {
		annot.QuadPoints = obj
	}

	return &annot, nil
}

func newPdfAnnotationUnderlineFromDict(d PdfObjectDictionary) (*PdfAnnotationUnderline, error) {
	annot := PdfAnnotationUnderline{}

	if obj, has := d["QuadPoints"]; has {
		annot.QuadPoints = obj
	}

	return &annot, nil
}

func newPdfAnnotationSquigglyFromDict(d PdfObjectDictionary) (*PdfAnnotationSquiggly, error) {
	annot := PdfAnnotationSquiggly{}

	if obj, has := d["QuadPoints"]; has {
		annot.QuadPoints = obj
	}

	return &annot, nil
}

func newPdfAnnotationStrikeOut(d PdfObjectDictionary) (*PdfAnnotationStrikeOut, error) {
	annot := PdfAnnotationStrikeOut{}

	if obj, has := d["QuadPoints"]; has {
		annot.QuadPoints = obj
	}

	return &annot, nil
}

func newPdfAnnotationCaretFromDict(d PdfObjectDictionary) (*PdfAnnotationCaret, error) {
	annot := PdfAnnotationCaret{}
	if obj, has := d["RD"]; has {
		annot.RD = obj
	}

	if obj, has := d["Sy"]; has {
		annot.Sy = obj
	}

	return &annot, nil
}
func newPdfAnnotationStampFromDict(d PdfObjectDictionary) (*PdfAnnotationStamp, error) {
	annot := PdfAnnotationStamp{}

	if obj, has := d["Name"]; has {
		annot.Name = obj
	}

	return &annot, nil
}

func newPdfAnnotationInkFromDict(d PdfObjectDictionary) (*PdfAnnotationInk, error) {
	annot := PdfAnnotationInk{}

	if obj, has := d["InkList"]; has {
		annot.InkList = obj
	}

	if obj, has := d["BS"]; has {
		annot.BS = obj
	}

	return &annot, nil
}

func newPdfAnnotationPopupFromDict(d PdfObjectDictionary) (*PdfAnnotationPopup, error) {
	annot := PdfAnnotationPopup{}

	if obj, has := d["Parent"]; has {
		annot.Parent = obj
	}

	if obj, has := d["Open"]; has {
		annot.Open = obj
	}

	return &annot, nil
}

func newPdfAnnotationFileAttachmentFromDict(d PdfObjectDictionary) (*PdfAnnotationFileAttachment, error) {
	annot := PdfAnnotationFileAttachment{}

	if obj, has := d["FS"]; has {
		annot.FS = obj
	}

	if obj, has := d["Name"]; has {
		annot.Name = obj
	}

	return &annot, nil
}

func newPdfAnnotationSoundFromDict(d PdfObjectDictionary) (*PdfAnnotationSound, error) {
	annot := PdfAnnotationSound{}

	if obj, has := d["Name"]; has {
		annot.Name = obj
	}

	if obj, has := d["Sound"]; has {
		annot.Sound = obj
	}

	return &annot, nil
}

func newPdfAnnotationMovieFromDict(d PdfObjectDictionary) (*PdfAnnotationMovie, error) {
	annot := PdfAnnotationMovie{}

	if obj, has := d["T"]; has {
		annot.T = obj
	}

	if obj, has := d["Movie"]; has {
		annot.Movie = obj
	}

	if obj, has := d["A"]; has {
		annot.A = obj
	}

	return &annot, nil
}

func newPdfAnnotationScreenFromDict(d PdfObjectDictionary) (*PdfAnnotationScreen, error) {
	annot := PdfAnnotationScreen{}

	if obj, has := d["T"]; has {
		annot.T = obj
	}

	if obj, has := d["MK"]; has {
		annot.MK = obj
	}

	if obj, has := d["A"]; has {
		annot.A = obj
	}

	if obj, has := d["AA"]; has {
		annot.AA = obj
	}

	return &annot, nil
}

func newPdfAnnotationWidgetFromDict(d PdfObjectDictionary) (*PdfAnnotationWidget, error) {
	annot := PdfAnnotationWidget{}
	if obj, has := d["H"]; has {
		annot.H = obj
	}

	if obj, has := d["MK"]; has {
		annot.MK = obj
	}

	if obj, has := d["A"]; has {
		annot.A = obj
	}

	if obj, has := d["AA"]; has {
		annot.AA = obj
	}

	if obj, has := d["BS"]; has {
		annot.BS = obj
	}

	if obj, has := d["Parent"]; has {
		annot.Parent = obj
	}

	return &annot, nil
}

func newPdfAnnotationPrinterMarkFromDict(d PdfObjectDictionary) (*PdfAnnotationPrinterMark, error) {
	annot := PdfAnnotationPrinterMark{}

	if obj, has := d["MN"]; has {
		annot.MN = obj
	}

	return &annot, nil
}

func newPdfAnnotationTrapNetFromDict(d PdfObjectDictionary) (*PdfAnnotationTrapNet, error) {
	annot := PdfAnnotationTrapNet{}
	// empty?e
	return &annot, nil
}

func newPdfAnnotationWatermarkFromDict(d PdfObjectDictionary) (*PdfAnnotationWatermark, error) {
	annot := PdfAnnotationWatermark{}

	if obj, has := d["FixedPrint"]; has {
		annot.FixedPrint = obj
	}

	return &annot, nil
}

func newPdfAnnotation3DFromDict(d PdfObjectDictionary) (*PdfAnnotation3D, error) {
	annot := PdfAnnotation3D{}

	if obj, has := d["3DD"]; has {
		annot.T3DD = obj
	}
	if obj, has := d["3DV"]; has {
		annot.T3DV = obj
	}
	if obj, has := d["3DA"]; has {
		annot.T3DA = obj
	}
	if obj, has := d["3DI"]; has {
		annot.T3DI = obj
	}
	if obj, has := d["3DB"]; has {
		annot.T3DB = obj
	}

	return &annot, nil
}

func newPdfAnnotationRedactFromDict(d PdfObjectDictionary) (*PdfAnnotationRedact, error) {
	annot := PdfAnnotationRedact{}

	if obj, has := d["QuadPoints"]; has {
		annot.QuadPoints = obj
	}
	if obj, has := d["IC"]; has {
		annot.IC = obj
	}
	if obj, has := d["RO"]; has {
		annot.RO = obj
	}
	if obj, has := d["OverlayText"]; has {
		annot.OverlayText = obj
	}
	if obj, has := d["Repeat"]; has {
		annot.Repeat = obj
	}
	if obj, has := d["DA"]; has {
		annot.DA = obj
	}
	if obj, has := d["Q"]; has {
		annot.Q = obj
	}

	return &annot, nil
}


func (this *PdfAnnotation) ToPdfObjectRef() PdfObject {
	var container *PdfIndirectObject
	var d *PdfObjectDictionary

	if cachedObj, isCached := PdfObjectConverterCache[this] {
		container = cachedObj
		d = container.PdfObject.(*PdfObjectDictionary)
	} else {
		container := &PdfIndirectObject{}
		d := &PdfObjectDictionary{}
		container.PdfObject = d
	}

	return container
}

func (this *PdfAnnotation) ToPdfObject() PdfObject {
	container := this.PdfAnnotation.ToPdfObject() 
	d := container.PdfObject.(*PdfObjectDictionary)

	d.SetIfNotNil("AP", this.AP)
	d.SetIfNotNil("AS", this.AS)
	d.SetIfNotNil("Border", this.Border)
	d.SetIfNotNil("C", this.C)
	d.SetIfNotNil("Contents", this.Contents)
	d.SetIfNotNil("F", this.F)
	d.SetIfNotNil("M", this.M)

	return container
}


func (this *PdfAnnotationText) ToPdfObject() PdfObject {
	container := this.PdfAnnotation.ToPdfObject() 
	d := container.PdfObject.(*PdfObjectDictionary)

	d.SetIfNotNil("Open", this.Open)
	d.SetIfNotNil("Name", this.Name)
	d.SetIfNotNil("State", this.State)
	d.SetIfNotNil("StateModel", this.StateModel)
	return container
}

func (this *PdfAnnotationLink) ToPdfObject()   PdfObject        {
	container := this.PdfAnnotation.ToPdfObject() 
	d := container.PdfObject.(*PdfObjectDictionary)

	d.SetIfNotNil("A", this.A)
	d.SetIfNotNil("Dest", this.Dest)
	d.SetIfNotNil("H", this.H)
	d.SetIfNotNil("PA", this.PA)
	d.SetIfNotNil("QuadPoints", this.QuadPoints)
	d.SetIfNotNil("BS", this.BS)
	return container
}

func (this *PdfAnnotationFreeText) ToPdfObject() PdfObject      {
	container := this.PdfAnnotation.ToPdfObject() 
	d := container.PdfObject.(*PdfObjectDictionary)


	d.SetIfNotNil("DA", this.DA)
	d.SetIfNotNil("Q", this.Q)
	d.SetIfNotNil("RC", this.RC)
	d.SetIfNotNil("DS", this.DS)
	d.SetIfNotNil("CL", this.CL)
	d.SetIfNotNil("IT", this.IT)
	d.SetIfNotNil("BE", this.BE)
	d.SetIfNotNil("RD", this.RD)
	d.SetIfNotNil("BS", this.BS)
	d.SetIfNotNil("LE", this.LE)

	return container
}
func (this *PdfAnnotationLine) ToPdfObject() PdfObject          {
	container := this.PdfAnnotation.ToPdfObject() 
	d := container.PdfObject.(*PdfObjectDictionary)

	d.SetIfNotNil("L", this.L)
	d.SetIfNotNil("BS", this.BS)
	d.SetIfNotNil("LE", this.LE)
	d.SetIfNotNil("IC", this.IC)
	d.SetIfNotNil("LL", this.LL)
	d.SetIfNotNil("LLE", this.LLE)
	d.SetIfNotNil("Cap", this.Cap)
	d.SetIfNotNil("IT", this.IT)
	d.SetIfNotNil("LLO", this.LLO)
	d.SetIfNotNil("CP", this.CP)
	d.SetIfNotNil("Measure", this.Measure)
	d.SetIfNotNil("CO", this.CO)

	return container
}

func (this *PdfAnnotationSquare) ToPdfObject() PdfObject {
	container := this.PdfAnnotation.ToPdfObject() 
	d := container.PdfObject.(*PdfObjectDictionary)

	d.SetIfNotNil("BS", this.BS)
	d.SetIfNotNil("IC", this.IC)
	d.SetIfNotNil("BE", this.BE)
	d.SetIfNotNil("RD", this.RD)

	return container
}

func (this *PdfAnnotationCircle) ToPdfObject()         {
	container := this.PdfAnnotation.ToPdfObject() 
	d := container.PdfObject.(*PdfObjectDictionary)

	d.SetIfNotNil("BS", this.BS)
	d.SetIfNotNil("IC", this.IC)
	d.SetIfNotNil("BE", this.BE)
	d.SetIfNotNil("RD", this.RD)

	return container
}

func (this *PdfAnnotationPolygon) ToPdfObject()        {
	container := this.PdfAnnotation.ToPdfObject() 
	d := container.PdfObject.(*PdfObjectDictionary)

	d.SetIfNotNil("Vertices", this.Vertices)
	d.SetIfNotNil("LE", this.LE)
	d.SetIfNotNil("BS", this.BS)
	d.SetIfNotNil("IC", this.IC)
	d.SetIfNotNil("BE", this.BE)
	d.SetIfNotNil("IT", this.IT)
	d.SetIfNotNil("Measure", this.Measure)

	return container
}

func (this *PdfAnnotationPolyLine) ToPdfObject()       {
	container := this.PdfAnnotation.ToPdfObject() 
	d := container.PdfObject.(*PdfObjectDictionary)

	d.SetIfNotNil("Vertices", this.Vertices)
	d.SetIfNotNil("LE", this.LE)
	d.SetIfNotNil("BS", this.BS)
	d.SetIfNotNil("IC", this.IC)
	d.SetIfNotNil("BE", this.BE)
	d.SetIfNotNil("IT", this.IT)
	d.SetIfNotNil("Measure", this.Measure)

	return container
}

func (this *PdfAnnotationHighlight) ToPdfObject()      {
	container := this.PdfAnnotation.ToPdfObject() 
	d := container.PdfObject.(*PdfObjectDictionary)

	d.SetIfNotNil("QuadPoints", this.QuadPoints)
	return container
}

func (this *PdfAnnotationUnderline) ToPdfObject()      {
		QuadPoints PdfObject
}
func (this *PdfAnnotationSquiggly) ToPdfObject()      {
		QuadPoints PdfObject
}
func (this *PdfAnnotationStrikeOut) ToPdfObject()      {
		QuadPoints PdfObject
}
func (this *PdfAnnotationCaret) ToPdfObject()          {
		RD PdfObject
	Sy PdfObject
}
func (this *PdfAnnotationStamp) ToPdfObject()          {
Name PdfObject
}
func (this *PdfAnnotationInk) ToPdfObject()            {
		InkList PdfObject
	BS      PdfObject
}
func (this *PdfAnnotationPopup) ToPdfObject()          {
	Parent PdfObject
	Open   PdfObject
}
func (this *PdfAnnotationFileAttachment) ToPdfObject() {
	FS   PdfObject
	Name PdfObject
}
func (this *PdfAnnotationSound) ToPdfObject()          {
Sound PdfObject
	Name  PdfObject
}
func (this *PdfAnnotationMovie) ToPdfObject()          {
	T     PdfObject
	Movie PdfObject
	A     PdfObject
}
func (this *PdfAnnotationScreen) ToPdfObject()         {
T  PdfObject
	MK PdfObject
	A  PdfObject
	AA PdfObject

}
func (this *PdfAnnotationWidget) ToPdfObject()         {
H      PdfObject
	MK     PdfObject
	A      PdfObject
	AA     PdfObject
	BS     PdfObject
	Parent PdfObject

}
func (this *PdfAnnotationPrinterMark) ToPdfObject()    {
	MN PdfObject

}
func (this *PdfAnnotationTrapNet) ToPdfObject()        {
}
func (this *PdfAnnotationWatermark) ToPdfObject()      {
FixedPrint PdfObject
}
func (this *PdfAnnotation3D) ToPdfObject()             {
	T3DD PdfObject
	T3DV PdfObject
	T3DA PdfObject
	T3DI PdfObject
	T3DB PdfObject
}
func (this *PdfAnnotationRedact) ToPdfObject() {
	QuadPoints  PdfObject
	IC          PdfObject
	RO          PdfObject
	OverlayText PdfObject
	Repeat      PdfObject
	DA          PdfObject
	Q           PdfObject
}



/////////

/*
func (r *PdfReader) newPdfAnnotationWidgetFromDict(d PdfObjectDictionary, parent *PdfField) *PdfAnnotationWidget {
	annotation := PdfAnnotationWidget{}

	if obj, has := d["Subtype"]; has {
		annotation.Subtype = obj
	}
	if obj, has := d["H"]; has {
		annotation.H = obj
	}
	if obj, has := d["MK"]; has {
		annotation.MK = obj
	}
	if obj, has := d["A"]; has {
		annotation.A = obj
	}
	if obj, has := d["AA"]; has {
		annotation.AA = obj
	}
	if obj, has := d["BS"]; has {
		annotation.BS = obj
	}

	annotation.Parent = parent
	return &annotation
}

func (this *PdfAnnotationWidget) ToPdfObject(updateIfExists bool) PdfObject {
	var container PdfIndirectObject

	if cachedObj, isCached := PdfObjectConverterCache[this]; isCached {
		if !updateIfExists {
			return cachedObj
		}
		obj := cachedObj.(*PdfIndirectObject)
		container = *obj
	}

	container = PdfIndirectObject{}
	dict := PdfObjectDictionary{}
	container.PdfObject = &dict
	d := PdfObjectDictionary{}

	d.SetIfNotNil("Subtype", this.Subtype)
	d.SetIfNotNil("H", this.H)
	d.SetIfNotNil("MK", this.MK)
	d.SetIfNotNil("A", this.A)
	d.SetIfNotNil("AA", this.AA)
	d.SetIfNotNil("BS", this.BS)
	if this.Parent != nil {
		d["Parent"] = this.Parent.ToPdfObject(false)
	}

	PdfObjectConverterCache[this] = &container
	return &container
}
*/
