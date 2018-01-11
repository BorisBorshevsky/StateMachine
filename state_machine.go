package goStateMachine

type Stater interface {
	GetState() *State
	SetState(*State)
}

type StateMachine struct {
	Stater
	TransitionCallbacks map[Transition]func(data ...interface{})
}

type Transition struct {
	From  *State
	Event string
	To    *State
}

func (d *StateMachine) RegisterCallback(transition Transition, cb func(...interface{})) {
	d.TransitionCallbacks[transition] = cb
}

func (d *StateMachine) handleTransition(transition Transition, data ...interface{}) {
	if fn, ok := d.TransitionCallbacks[transition]; ok {
		fn(data...)
	}
}

func (d *StateMachine) HandleEvent(event string, data ...interface{}) *Transition {
	t := &Transition{
		From:  d.GetState(),
		Event: event,
	}

	to := t.From.Handle(event)

	if to != nil {
		d.SetState(to)
		to.Callback()
		t.To = to

		d.handleTransition(*t, data...)
		return t
	}

	return nil
}

func NewStateMachine(stater Stater, InitialState *State) *StateMachine {
	stater.SetState(InitialState)
	return &StateMachine{
		Stater:              stater,
		TransitionCallbacks: make(map[Transition]func(data ...interface{})),
	}
}
