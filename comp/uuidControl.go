package comp

// UUIDControl UUID 功能性组件 文档：https://aisuda.bce.baidu.com/amis/zh-CN/components/form/uuid
//
// @author  slowlyo
// @version 6.7.0
type UUIDControl Schema

// NewUUIDControl 创建一个新的 UUIDControl 实例
func NewUUIDControl() UUIDControl {
	return UUIDControl{}.set("type", "uuid")
}

func (uc UUIDControl) set(key string, value interface{}) UUIDControl {
	uc[key] = value
	return uc
}

// AutoFill 自动填充，当选项被选择的时候，将选项中的其他值同步设置到表单内。
func (uc UUIDControl) AutoFill(value string) UUIDControl {
	return uc.set("autoFill", value)
}

// ClassName 容器 css 类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (uc UUIDControl) ClassName(value string) UUIDControl {
	return uc.set("className", value)
}

// ClearValueOnHidden 表单项隐藏时，是否在当前 Form 中删除掉该表单项值。注意同名的未隐藏的表单项值也会删掉
func (uc UUIDControl) ClearValueOnHidden(value bool) UUIDControl {
	return uc.set("clearValueOnHidden", value)
}

// Desc
func (uc UUIDControl) Desc(value string) UUIDControl {
	return uc.set("desc", value)
}

// Description 描述内容，支持 Html 片段。
func (uc UUIDControl) Description(value string) UUIDControl {
	return uc.set("description", value)
}

// DescriptionClassName 配置描述上的 className (配置描述上的 className)
func (uc UUIDControl) DescriptionClassName(value string) UUIDControl {
	return uc.set("descriptionClassName", value)
}

// Disabled 是否禁用
func (uc UUIDControl) Disabled(value bool) UUIDControl {
	return uc.set("disabled", value)
}

// DisabledOn 是否禁用表达式 (表达式，语法 `data.xxx > 5`。)
func (uc UUIDControl) DisabledOn(value string) UUIDControl {
	return uc.set("disabledOn", value)
}

// EditorSetting 编辑器配置，运行时可以忽略
func (uc UUIDControl) EditorSetting(value string) UUIDControl {
	return uc.set("editorSetting", value)
}

// ExtraName 额外的字段名，当为范围组件时可以用来将另外一个值打平出来
func (uc UUIDControl) ExtraName(value string) UUIDControl {
	return uc.set("extraName", value)
}

// Hidden 是否隐藏
func (uc UUIDControl) Hidden(value bool) UUIDControl {
	return uc.set("hidden", value)
}

// HiddenOn 是否隐藏表达式 (表达式，语法 `data.xxx > 5`。)
func (uc UUIDControl) HiddenOn(value string) UUIDControl {
	return uc.set("hiddenOn", value)
}

// Hint 输入提示，聚焦的时候显示
func (uc UUIDControl) Hint(value string) UUIDControl {
	return uc.set("hint", value)
}

// Horizontal 当配置为水平布局的时候，用来配置具体的左右分配。 (当配置为水平布局的时候，用来配置具体的左右分配。)
func (uc UUIDControl) Horizontal(value string) UUIDControl {
	return uc.set("horizontal", value)
}

// Id 组件唯一 id，主要用于日志采集
func (uc UUIDControl) Id(value string) UUIDControl {
	return uc.set("id", value)
}

// InitAutoFill
func (uc UUIDControl) InitAutoFill(value string) UUIDControl {
	return uc.set("initAutoFill", value)
}

// Inline 表单 control 是否为 inline 模式。
func (uc UUIDControl) Inline(value bool) UUIDControl {
	return uc.set("inline", value)
}

// InputClassName 配置 input className (配置 input className)
func (uc UUIDControl) InputClassName(value string) UUIDControl {
	return uc.set("inputClassName", value)
}

// Label 描述标题
func (uc UUIDControl) Label(value string) UUIDControl {
	return uc.set("label", value)
}

// LabelAlign 描述标题 (描述标题) 可选值: right | left | top | inherit
func (uc UUIDControl) LabelAlign(value string) UUIDControl {
	return uc.set("labelAlign", value)
}

// LabelClassName 配置 label className
func (uc UUIDControl) LabelClassName(value string) UUIDControl {
	return uc.set("labelClassName", value)
}

// LabelRemark 显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起 (显示一个小图标, 鼠标放上去的时候显示提示内容, 这个小图标跟 label 在一起)
func (uc UUIDControl) LabelRemark(value string) UUIDControl {
	return uc.set("labelRemark", value)
}

// LabelWidth label自定义宽度，默认单位为px
func (uc UUIDControl) LabelWidth(value string) UUIDControl {
	return uc.set("labelWidth", value)
}

// Length 长度，默认 uuid 的长度是 36，如果不需要那么长，可以设置这个来缩短
func (uc UUIDControl) Length(value string) UUIDControl {
	return uc.set("length", value)
}

// Mode 配置当前表单项展示模式 可选值: normal | inline | horizontal
func (uc UUIDControl) Mode(value string) UUIDControl {
	return uc.set("mode", value)
}

// Name 字段名，表单提交时的 key，支持多层级，用.连接，如： a.b.c
func (uc UUIDControl) Name(value string) UUIDControl {
	return uc.set("name", value)
}

// OnEvent 事件动作配置
func (uc UUIDControl) OnEvent(value string) UUIDControl {
	return uc.set("onEvent", value)
}

// Placeholder 占位符
func (uc UUIDControl) Placeholder(value string) UUIDControl {
	return uc.set("placeholder", value)
}

// ReadOnly 是否只读
func (uc UUIDControl) ReadOnly(value bool) UUIDControl {
	return uc.set("readOnly", value)
}

// ReadOnlyOn 只读条件
func (uc UUIDControl) ReadOnlyOn(value string) UUIDControl {
	return uc.set("readOnlyOn", value)
}

// Remark 显示一个小图标, 鼠标放上去的时候显示提示内容 (显示一个小图标, 鼠标放上去的时候显示提示内容)
func (uc UUIDControl) Remark(value string) UUIDControl {
	return uc.set("remark", value)
}

// Required 是否为必填
func (uc UUIDControl) Required(value bool) UUIDControl {
	return uc.set("required", value)
}

// Row
func (uc UUIDControl) Row(value string) UUIDControl {
	return uc.set("row", value)
}

// SaveImmediately 是否立即保存(TableColumn中使用)
func (uc UUIDControl) SaveImmediately(value bool) UUIDControl {
	return uc.set("saveImmediately", value)
}

// Size 表单项大小 可选值: xs | sm | md | lg | full
func (uc UUIDControl) Size(value string) UUIDControl {
	return uc.set("size", value)
}

// Static 是否静态展示
func (uc UUIDControl) Static(value bool) UUIDControl {
	return uc.set("static", value)
}

// StaticClassName 静态展示表单项类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (uc UUIDControl) StaticClassName(value string) UUIDControl {
	return uc.set("staticClassName", value)
}

// StaticInputClassName 静态展示表单项Value类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (uc UUIDControl) StaticInputClassName(value string) UUIDControl {
	return uc.set("staticInputClassName", value)
}

// StaticLabelClassName 静态展示表单项Label类名 (css类名，配置字符串，或者对象。className: "red"用对象配置时意味着你能跟表达式一起搭配使用，如：className: { "red": "data.progress > 80", "blue": "data.progress > 60" })
func (uc UUIDControl) StaticLabelClassName(value string) UUIDControl {
	return uc.set("staticLabelClassName", value)
}

// StaticOn 是否静态展示表达式 (表达式，语法 `data.xxx > 5`。)
func (uc UUIDControl) StaticOn(value string) UUIDControl {
	return uc.set("staticOn", value)
}

// StaticPlaceholder 静态展示空值占位
func (uc UUIDControl) StaticPlaceholder(value string) UUIDControl {
	return uc.set("staticPlaceholder", value)
}

// StaticSchema
func (uc UUIDControl) StaticSchema(value string) UUIDControl {
	return uc.set("staticSchema", value)
}

// Style 组件样式
func (uc UUIDControl) Style(value string) UUIDControl {
	return uc.set("style", value)
}

// SubmitOnChange 当修改完的时候是否提交表单。
func (uc UUIDControl) SubmitOnChange(value bool) UUIDControl {
	return uc.set("submitOnChange", value)
}

// TestIdBuilder
func (uc UUIDControl) TestIdBuilder(value string) UUIDControl {
	return uc.set("testIdBuilder", value)
}

// Type 表单项类型
func (uc UUIDControl) Type(value string) UUIDControl {
	return uc.set("type", value)
}

// UseMobileUI 可以组件级别用来关闭移动端样式
func (uc UUIDControl) UseMobileUI(value bool) UUIDControl {
	return uc.set("useMobileUI", value)
}

// ValidateApi 远端校验表单项接口
func (uc UUIDControl) ValidateApi(value string) UUIDControl {
	return uc.set("validateApi", value)
}

// ValidateOnChange 不设置时，当表单提交过后表单项每次修改都会触发重新验证， 如果设置了，则由此配置项来决定要不要每次修改都触发验证。
func (uc UUIDControl) ValidateOnChange(value bool) UUIDControl {
	return uc.set("validateOnChange", value)
}

// ValidationErrors 验证失败的提示信息
func (uc UUIDControl) ValidationErrors(value string) UUIDControl {
	return uc.set("validationErrors", value)
}

// Validations
func (uc UUIDControl) Validations(value string) UUIDControl {
	return uc.set("validations", value)
}

// Value 默认值，切记只能是静态值，不支持取变量，跟数据关联是通过设置 name 属性来实现的。
func (uc UUIDControl) Value(value string) UUIDControl {
	return uc.set("value", value)
}

// Visible 是否显示
func (uc UUIDControl) Visible(value bool) UUIDControl {
	return uc.set("visible", value)
}

// VisibleOn 是否显示表达式 (表达式，语法 `data.xxx > 5`。)
func (uc UUIDControl) VisibleOn(value string) UUIDControl {
	return uc.set("visibleOn", value)
}

// Width 在Table中调整宽度
func (uc UUIDControl) Width(value string) UUIDControl {
	return uc.set("width", value)
}
