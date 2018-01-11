package goStateMachine

type Handler func(ev string) *State

type State struct {
	Name    string
	Handler Handler
	OnSet   func()
	Data    interface{}
}

func (s *State) Handle(ev string) *State {
	if s.Handler != nil {
		return s.Handler(ev)
	}
	return s
}

func (s *State) Callback() {
	if s.OnSet != nil {
		s.OnSet()
	}
}
