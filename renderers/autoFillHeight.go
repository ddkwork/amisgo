package renderers

// AutoFillHeight 自动填充高度渲染器
type AutoFillHeight struct {
	*BaseRenderer
}

// NewAutoFillHeight 创建一个新的 AutoFillHeight 实例
func NewAutoFillHeight() *AutoFillHeight {
	return &AutoFillHeight{BaseRenderer: NewBaseRenderer()}
}

// Set 设置属性
func (a *AutoFillHeight) set(key string, value interface{}) *AutoFillHeight {
	a.BaseRenderer.set(key, value)
	return a
}

// Height 设置高度
func (a *AutoFillHeight) Height(value string) *AutoFillHeight {
	return a.set("height", value)
}

// MaxHeight 设置最大高度
func (a *AutoFillHeight) MaxHeight(value string) *AutoFillHeight {
	return a.set("maxHeight", value)
}
