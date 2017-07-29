package html

type HtmlInput struct {
	buttonType string
	id string
	name string
	label string
	value string
	classes []string
	isRequired bool
	options []Option
	validation func() (errors []error)
}

func NewInput() Input {
	return &HtmlInput{}
}

func (htmlInput *HtmlInput) SetType(buttonType string) {
	htmlInput.buttonType = buttonType
}

func (htmlInput *HtmlInput) GetType() string {
	return htmlInput.buttonType
}

func (htmlInput *HtmlInput) SetId(id string) {
	htmlInput.id = id
}

func (htmlInput *HtmlInput) GetId() string {
	return htmlInput.id
}

func (htmlInput *HtmlInput) SetName(name string) {
	htmlInput.name = name
}

func (htmlInput *HtmlInput) GetName() string {
	return htmlInput.name
}

func (htmlInput *HtmlInput) SetLabel(label string) {
	htmlInput.label = label
}

func (htmlInput *HtmlInput) GetLabel() string {
	return htmlInput.label
}

func (htmlInput *HtmlInput) SetValue(value string) {
	htmlInput.value = value
}

func (htmlInput *HtmlInput) GetValue() string {
	return htmlInput.value
}

func (htmlInput *HtmlInput) AddClass(class string) {
	htmlInput.classes = append(
		htmlInput.classes,
		class,
	)
}

func (htmlInput *HtmlInput) SetClasses(classes []string) {
	htmlInput.classes = classes
}

func (htmlInput *HtmlInput) GetClasses() []string {
	return htmlInput.classes
}

func (htmlInput *HtmlInput) SetIsRequired(required bool) {
	htmlInput.isRequired = required
}

func (htmlInput *HtmlInput) IsRequired() bool {
	return htmlInput.isRequired
}

func (htmlInput *HtmlInput) AddOption(option Option) {
	htmlInput.options = append(
		htmlInput.options,
		option,
	)
}

func (htmlInput *HtmlInput) SetOptions(options []Option) {
	htmlInput.options = options
}

func (htmlInput *HtmlInput) GetGetOptions() []Option {
	return htmlInput.options
}

func (htmlInput *HtmlInput) SetValidation(validation func() (errors []error)) {
	htmlInput.validation = validation
}

func (htmlInput *HtmlInput) Validate() []error {
	return htmlInput.validation()
}
