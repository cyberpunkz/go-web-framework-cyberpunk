package html

type Form interface {
	SetTitle(title string)
	GetTitle() string

	SetActionType(actionType string)
	GetActionType() string

	SetActionUrl(url string)
	GetActionUrl() string

	AddClass(class string)
	SetClasses(classes []string)
	GetClasses() []string

	SetButton(button Button)
	GetButton() Button

	AddGroup(group Group)
	SetGroups(groups []Group)
	GetGroups() []Group
}

type Button interface {
	SetType(buttonType string)
	GetType() string

	SetId(id string)
	GetId() string

	SetName(name string)
	GetName() string

	SetLabel(label string)
	GetLabel() string

	AddClass(class string)
	SetClasses(classes []string)
	GetClasses() []string
}

type Group interface {
	SetTitle(title string)
	GetTitle() string

	AddClass(class string)
	SetClasses(classes []string)
	GetClasses() []string

	AddInput(input Input)
	SetInputs(inputs []Input)
	GetInputs() []Input
}

type Input interface {
	SetType(buttonType string)
	GetType() string

	SetId(id string)
	GetId() string

	SetName(name string)
	GetName() string

	SetLabel(label string)
	GetLabel() string

	SetValue(value string)
	GetValue() string

	AddClass(class string)
	SetClasses(classes []string)
	GetClasses() []string

	SetIsRequired(required bool)
	IsRequired() bool

	AddOption(option Option)
	SetOptions(options []Option)
	GetGetOptions() []Option
	
	SetValidation(validation func() (errors []error))
	Validate() []error
}

type Option interface {
	SetLabel(label string)
	GetLabel() string

	SetValue(value string)
	GetValue() string
}
