package render

// InfoImpl 구조체는 Info 인터페이스의 내부 구현체
type InfoImpl struct {
	PosX, PosY int32

	Color uint32
}

// GetPos 함수는 좌표 정보를 반환한다.
func (impl InfoImpl) GetPos() (int32, int32) {
	return impl.PosX, impl.PosY
}

// GetColor 함수는 색상 정보를 반환한다.
func (impl *InfoImpl) GetColor() uint32 {
	return impl.Color
}
