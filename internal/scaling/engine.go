package scaling

type Engine struct {
	Evaluator *Evaluator
}

func (e *Engine) Decide() string {
	return "STABLE"
}
