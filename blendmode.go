package psd

// BlendMode represents the blend mode.
type BlendMode string

// These blend modes are defined in this document.
//
// http://www.adobe.com/devnet-apps/photoshop/fileformatashtml/#50577409_13084
const (
	BlendModePassThrough  = BlendMode("pass")
	BlendModeNormal       = BlendMode("norm")
	BlendModeDissolve     = BlendMode("diss")
	BlendModeDarken       = BlendMode("dark")
	BlendModeMultiply     = BlendMode("mul ")
	BlendModeColorBurn    = BlendMode("idiv")
	BlendModeLinearBurn   = BlendMode("lbrn")
	BlendModeDarkerColor  = BlendMode("dkCl")
	BlendModeLighten      = BlendMode("lite")
	BlendModeScreen       = BlendMode("scrn")
	BlendModeColorDodge   = BlendMode("div ")
	BlendModeLinearDodge  = BlendMode("lddg")
	BlendModeLighterColor = BlendMode("lgCl")
	BlendModeOverlay      = BlendMode("over")
	BlendModeSoftLight    = BlendMode("sLit")
	BlendModeHardLight    = BlendMode("hLit")
	BlendModeVividLight   = BlendMode("vLit")
	BlendModeLinearLight  = BlendMode("lLit")
	BlendModePinLight     = BlendMode("pLit")
	BlendModeHardMix      = BlendMode("hMix")
	BlendModeDifference   = BlendMode("diff")
	BlendModeExclusion    = BlendMode("smud")
	BlendModeSubtract     = BlendMode("fsub")
	BlendModeDivide       = BlendMode("fdiv")
	BlendModeHue          = BlendMode("hue ")
	BlendModeSaturation   = BlendMode("sat ")
	BlendModeColor        = BlendMode("colr")
	BlendModeLuminosity   = BlendMode("lum ")
)
