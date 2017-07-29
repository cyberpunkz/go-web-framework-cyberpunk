package html

type HtmlGroup struct {
	title string
	classes []string
	inputs []Input
}

func NewGroup() Group {
	return &HtmlGroup{}
}

func (htmlGroup *HtmlGroup) SetTitle(title string) {
	htmlGroup.title = title
}

func (htmlGroup *HtmlGroup) GetTitle() string {
	return htmlGroup.title
}

func (htmlGroup *HtmlGroup) AddClass(class string) {
	htmlGroup.classes = append(
		htmlGroup.classes,
		class,
	)
}

func (htmlGroup *HtmlGroup) SetClasses(classes []string) {
	htmlGroup.classes = classes
}

func (htmlGroup *HtmlGroup) GetClasses() []string {
	return htmlGroup.classes
}

func (htmlGroup *HtmlGroup) AddInput(input Input) {
	htmlGroup.inputs = append(
		htmlGroup.inputs,
		input,
	)
}

func (htmlGroup *HtmlGroup) SetInputs(inputs []Input) {
	htmlGroup.inputs = inputs
}

func (htmlGroup *HtmlGroup) GetInputs() []Input {
	return htmlGroup.inputs
}
