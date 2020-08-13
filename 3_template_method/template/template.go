package template

type TemplateInterface interface {
	StartCall()
	Run()
	GoalCall()
}

type Template struct {
	tmp TemplateInterface
}

func (t *Template) Go() {
	t.tmp.StartCall()
	t.tmp.Run()
	t.tmp.GoalCall()
}
