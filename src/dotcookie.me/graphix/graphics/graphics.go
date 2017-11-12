package graphics

type graphics struct {
	width,height int
}

func (g graphics) GetWidth() int{
	return g.width
}
func (g graphics) GetHeight() int{
	return g.height
}
