package goStateMachine

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

const InsertCardEvent = "insert_card_event"
const CancelEvent = "cancel_event"
const DoneEvent = "done_event"
const ErrorEvent = "error_event"
const FixedEvent = "fixed_event"

var IdleState *State
var ActiveState *State
var OutOfServiceState *State

var _ = BeforeSuite(func() {
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

	ActiveState = &State{
		Name: "active",
		Handler: func(ev string) *State {
			switch ev {
			case CancelEvent:
				return IdleState
			case DoneEvent:
				return IdleState
			case ErrorEvent:
				return OutOfServiceState
			default:
				return nil
			}
		},
	}

	OutOfServiceState = &State{
		Name: "outOfService",
		Handler: func(ev string) *State {
			switch ev {
			case FixedEvent:
				return IdleState
			default:
				return nil
			}
		},
	}
})

type InMemStater struct {
	State *State
}

func (m *InMemStater) GetState() *State {
	return m.State
}

func (m *InMemStater) SetState(s *State) {
	m.State = s
}

func TestStateBase(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "State Machine")
}

type ExtendedStateMachine struct {
	*StateMachine
	ExtraData string
}

func NewExtendedStateMachine(stater Stater, InitialState *State) *ExtendedStateMachine {
	return &ExtendedStateMachine{
		StateMachine: NewStateMachine(stater, InitialState),
	}
}

var _ = Describe("Basic State Machine", func() {
	var (
		stater       Stater
		stateMachine *StateMachine
	)

	BeforeEach(func() {
		stater = &InMemStater{}
		stateMachine = NewStateMachine(stater, IdleState)
	})

	Context("when got a handled event", func() {
		BeforeEach(func() {
			stateMachine.HandleEvent(InsertCardEvent)
		})

		It("got a new state", func() {
			Ω(stateMachine.GetState()).Should(Equal(ActiveState))
		})
	})

	Context("when got an unhandled event", func() {
		BeforeEach(func() {
			stateMachine.HandleEvent("NothingHappenEvent")
		})

		It("got a new state", func() {
			Ω(stateMachine.GetState()).Should(Equal(IdleState))
		})
	})

	Context("when we want to save extra data on the state machine", func() {
		var (
			extendedsm *ExtendedStateMachine
		)
		BeforeEach(func() {
			extendedsm = &ExtendedStateMachine{
				StateMachine: stateMachine,
			}
			extendedsm.RegisterCallback(Transition{
				From:  IdleState,
				To:    ActiveState,
				Event: InsertCardEvent,
			}, func(i ...interface{}) {
				extendedsm.ExtraData = i[0].(string)
			})

		})

		Context("event came with extra data", func() {
			var eventData string
			BeforeEach(func() {
				eventData = "1337-3455-3425-1123"
				extendedsm.HandleEvent(InsertCardEvent, eventData) //card number
			})
			It("should save the extra data", func() {
				Ω(extendedsm.ExtraData).Should(Equal(eventData))
			})

		})

	})

})
