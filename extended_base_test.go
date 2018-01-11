package goStateMachine

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)


var _ = Describe("Extended State Machine", func() {
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
})
