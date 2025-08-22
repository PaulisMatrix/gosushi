package statemachine

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/looplab/fsm"
)

type Workflow struct {
	FSM *fsm.FSM
}

func (wf *Workflow) InitiateSession(e *fsm.Event) {
	fmt.Println("initiating the session. Current transition", e.Src, e.Dst)
	wf.FSM.SetMetadata("sessionId", uuid.NewString())
	wf.FSM.SetMetadata("requestId", uuid.NewString())
}

func (wf *Workflow) RunInputRuleGuardRail(e *fsm.Event) {
	fmt.Println("Running input rule guardrail. Current transition", e.Src, e.Dst)
	wf.FSM.SetMetadata("state", "input_rule_guardrail")

	e.Cancel(errors.New("input rule guardrail flagged"))
}

func (wf *Workflow) RunInputLDRuleGuardRail(e *fsm.Event) {
	fmt.Println("Running input ld rule guardrail. Current transition", e.Src, e.Dst)
	_curState, _ := wf.FSM.Metadata("state")
	curState, _ := _curState.(string)
	curState = curState + "input_ld_rule_guardrail"

	wf.FSM.SetMetadata("state", curState)
}

func (wf *Workflow) RunInputModelGuardRail(e *fsm.Event) {
	fmt.Println("Running input model guardrail. Current transition", e.Src, e.Dst)
	_curState, _ := wf.FSM.Metadata("state")
	curState, _ := _curState.(string)
	curState = curState + "input_model_rule_guardrail"

	wf.FSM.SetMetadata("state", curState)
}

func (wf *Workflow) RunMMGetContext(e *fsm.Event) {
	fmt.Println("Running MM Get Context. Current transition", e.Src, e.Dst)
	sessionId, _ := wf.FSM.Metadata("sessionId")
	requestId, _ := wf.FSM.Metadata("requestId")
	fmt.Println("session details", sessionId, requestId)

	_curState, _ := wf.FSM.Metadata("state")
	curState, _ := _curState.(string)
	curState = curState + "context_retrieval_mm_get_context"

	wf.FSM.SetMetadata("state", curState)

	fmt.Println("final state", curState)
}

func RunFSM() {

	wf := &Workflow{}

	wf.FSM = fsm.NewFSM(
		"START",
		fsm.Events{
			{Name: "Init", Src: []string{"START"}, Dst: "InputRuleGuardRail"},
			{Name: "RunInputRuleGuardrail", Src: []string{"InputRuleGuardRail"}, Dst: "InputLDRuleGuardRail"},
			{Name: "RunInputLDRuleGuardrail", Src: []string{"InputLDRuleGuardRail"}, Dst: "InputModelGuardRail"},
			{Name: "RunInputModelGuardRail", Src: []string{"InputModelGuardRail"}, Dst: "MMGetContext"},
			{Name: "RunMMGetContext", Src: []string{"MMGetContext"}, Dst: "END"},
		},
		fsm.Callbacks{
			"before_Init": func(ctx context.Context, e *fsm.Event) {
				wf.InitiateSession(e)
			},
			"leave_InputRuleGuardRail": func(ctx context.Context, e *fsm.Event) {
				wf.RunInputRuleGuardRail(e)
			},
			"leave_InputLDRuleGuardRail": func(ctx context.Context, e *fsm.Event) {
				wf.RunInputLDRuleGuardRail(e)
			},
			"leave_InputModelGuardRail": func(ctx context.Context, e *fsm.Event) {
				wf.RunInputModelGuardRail(e)
			},
			"leave_MMGetContext": func(ctx context.Context, e *fsm.Event) {
				wf.RunMMGetContext(e)
			},
			"after_RunMMGetContext": func(ctx context.Context, e *fsm.Event) {
				fmt.Println("Exiting")
			},
		},
	)

	wf.FSM.AvailableTransitions()

	fmt.Println("current state", wf.FSM.Current())

	err := wf.FSM.Event(context.Background(), "Init")
	if err != nil {
		panic(err)
	}

	err = wf.FSM.Event(context.Background(), "RunInputRuleGuardrail")
	if err != nil && strings.Contains(err.Error(), errors.New("input rule guardrail flagged").Error()) {
		fmt.Println("input rule guardrail flagged, exiting...", err.Error())
		return
	}

	err = wf.FSM.Event(context.Background(), "RunInputLDRuleGuardrail")
	if err != nil {
		panic(err)
	}

	err = wf.FSM.Event(context.Background(), "RunInputModelGuardRail")
	if err != nil {
		panic(err)
	}

	err = wf.FSM.Event(context.Background(), "RunMMGetContext")
	if err != nil {
		panic(err)
	}

	fmt.Println("current state", wf.FSM.Current())

}
