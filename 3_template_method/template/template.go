package template

type TemplateInterface interface {
	StartColl()
	Run()
	GoalCall()
}

type Template struct {
	template TemplateInterface
}

func (t *Template) Go() {
	t.template.StartColl()
	t.template.Run()
	t.template.GoalCall()
}
