package blend

func drawNormalNRGBAToNRGBAFast(dest []byte, src []byte, alpha uint32, d0, s0, y int, sx0 int, sx1 int, sxDelta int, syDelta int, dx0 int, dx1 int, dxDelta int, dyDelta int)

func init() {
	drawNormalNRGBAToNRGBA = drawNormalNRGBAToNRGBAFast
}
