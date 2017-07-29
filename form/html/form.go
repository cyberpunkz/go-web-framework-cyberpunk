package html

type HtmlForm struct {
	title      string
	actionType string
	url        string
	classes    []string
	button     Button
	groups     []Group
}

func NewForm() Form {
	return &HtmlForm{}
}

func (htmlForm *HtmlForm) SetTitle(title string) {
	htmlForm.title = title
}

func (htmlForm *HtmlForm) GetTitle() string {
	return htmlForm.title
}

func (htmlForm *HtmlForm) SetActionType(actionType string) {
	htmlForm.actionType = actionType
}

func (htmlForm *HtmlForm) GetActionType() string {
	return htmlForm.actionType
}

func (htmlForm *HtmlForm) SetActionUrl(url string) {
	htmlForm.url = url
}

func (htmlForm *HtmlForm) GetActionUrl() string {
	return htmlForm.url
}

func (htmlForm *HtmlForm) AddClass(class string) {
	htmlForm.classes = append(
		htmlForm.classes,
		class,
	)
}

func (htmlForm *HtmlForm) SetClasses(classes []string) {
	htmlForm.classes = classes
}

func (htmlForm *HtmlForm) GetClasses() []string {
	return htmlForm.classes
}

func (htmlForm *HtmlForm) SetButton(button Button) {
	htmlForm.button = button
}

func (htmlForm *HtmlForm) GetButton() Button {
	return htmlForm.button
}

func (htmlForm *HtmlForm) AddGroup(group Group) {
	htmlForm.groups = append(
		htmlForm.groups,
		group,
	)
}

func (htmlForm *HtmlForm) SetGroups(groups []Group) {
	htmlForm.groups = groups
}

func (htmlForm *HtmlForm) GetGroups() []Group {
	return htmlForm.groups
}
