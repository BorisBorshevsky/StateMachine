# StateMachine
State Machine implementation in Golang


### Creating State with rules
```go

var IdleState *State

IdleState = &State{
		Name: "idle",
		Handler: func(ev string) *State {
			switch ev {
			case InsertCardEvent:
				return ActiveState
			case ErrorEvent:
				return OutOfServiceState
			default:
				return nil
			}
		},
	}
	
```


