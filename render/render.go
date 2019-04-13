package render

type InfoImpl struct {
	PosX, PosY int32

	Color uint32
}

func (impl InfoImpl) GetPos() (int32, int32) {
	return impl.PosX, impl.PosY
}

func (impl *InfoImpl) GetColor() uint32 {
	return impl.Color
}
