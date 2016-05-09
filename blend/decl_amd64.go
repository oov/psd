package blend

func drawNormalNRGBAToNRGBAFast(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int)
func drawNormalNRGBAToNRGBAProtectAlphaFast(dest []byte, src []byte, alpha uint32, y int, xMin int, xMax int, dDelta int, sDelta int, xDelta int)

func init() {
	drawNormalNRGBAToNRGBA = drawNormalNRGBAToNRGBAFast
	drawNormalNRGBAToNRGBAProtectAlpha = drawNormalNRGBAToNRGBAProtectAlphaFast
}
