package html

type HtmlOption struct {
	label string
	value string
}

func NewOption() Option {
	return &HtmlOption{}
}

func (htmlOption *HtmlOption) SetLabel(label string) {
	htmlOption.label = label
}

func (htmlOption *HtmlOption) GetLabel() string {
	return htmlOption.label
}

func (htmlOption *HtmlOption) SetValue(value string) {
	htmlOption.value = value
}

func (htmlOption *HtmlOption) GetValue() string {
	return htmlOption.value
}
