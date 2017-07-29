package html

type HtmlButton struct {
	buttonType string
	id         string
	name       string
	label      string
	classes    []string
}

func NewButton() Button {
	return &HtmlButton{}
}

func (htmlButton *HtmlButton) SetType(buttonType string) {
	htmlButton.buttonType = buttonType
}

func (htmlButton *HtmlButton) GetType() string  {
	return htmlButton.buttonType
}

func (htmlButton *HtmlButton) SetId(id string) {
	htmlButton.id = id
}

func (htmlButton *HtmlButton) GetId() string {
	return htmlButton.id
}

func (htmlButton *HtmlButton) SetName(name string) {
	htmlButton.name = name
}

func (htmlButton *HtmlButton) GetName() string {
	return htmlButton.name
}

func (htmlButton *HtmlButton) SetLabel(label string) {
	htmlButton.label = label
}

func (htmlButton *HtmlButton) GetLabel() string {
	return htmlButton.label
}

func (htmlButton *HtmlButton) AddClass(class string) {
	htmlButton.classes = append(
		htmlButton.classes,
		class,
	)
}

func (htmlButton *HtmlButton) SetClasses(classes []string) {
	htmlButton.classes = classes
}

func (htmlButton *HtmlButton) GetClasses() []string {
	return htmlButton.classes
}
