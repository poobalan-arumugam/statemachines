package statemachine

// Note: very naive state machine implementation - but hopefully readable and easy enough to debug?

// Source:  test_machine1.sm1

// assembly = None
// basestatemachine = None
// comment = None
// fields = None
// hassubmachines = False
// implementationversion = 0.2
// modelfilename = test_machine1.sm1
// modelguid = 7e46f114-606e-4584-8492-dcf41ee57660
// namespace = MurphyPA.SM
// readonly = False
// statemachine = None
// statemachineversion = 16
// usingnamespaces = None

//note - not using transition.eventtype yet
//package MurphyPA
//{
//package SM
//{

import (
	"errors"
	"fmt"
)

type Model_xstatemachine_t interface {
	a_entry()
	a_exit()
	b_entry()
	b_exit()
	c_entry()
	c_exit()
	d_entry()
	d_exit()
	e_entry()
	e_exit()
	f_entry()
	f_exit()
	g_entry()
	g_exit()
	h_entry()
	h_exit()
	t_entry()
	t_exit()
	test1(source *Source1, ev *x) bool
	test2(source *Source1, ev *x) bool
	action1(source *Source1, ev *x)
	action2(source *Source1, ev *x)
	action3(source *Source1, ev *y)
	action4(source *Source2, ev *z)
	action5(source *Source1, ev *x)
	action6(source *Source1, ev *y)
}

type Ports_xstatemachine_t interface {

	//Source1_port PortBinding<Owner, source1, EventBase>

	//Source2_port PortBinding<Owner, source2, EventBase>

}

type PortsModel_xstatemachine_t interface {
	Model_xstatemachine_t
	Ports_xstatemachine_t
}

type xstatemachine_t struct {
	// statemachine_t<EventSourceBase, EventBase>

	initialised bool
	state       uint64
	instanceid  string
	model       Model_xstatemachine_t
}

func (sm *xstatemachine_t) State() uint64 {
	return sm.state
}

func (sm *xstatemachine_t) SetState(value uint64) {
	sm.state = value
}

func (sm *xstatemachine_t) IsInitialised() bool {
	return sm.initialised
}

func (sm *xstatemachine_t) SetInitialised(value bool) {
	sm.initialised = value
}

func (sm *xstatemachine_t) Initialise() {
	sm.initialise_statemachine()
}

func (sm *xstatemachine_t) UnhandledEvent(source EventSource, ev Event) error {
	fmt.Printf("UNHANDLED_EVENT: %#v %#v\n", source, ev)
	return nil
}

func (sm *xstatemachine_t) Dispatch(eventSource EventSource, ev Event) error {
	switch source := eventSource.(type) {
	case *Source1:
		return sm.dispatch_Source1(source, ev)

	case *Source2:
		return sm.dispatch_Source2(source, ev)

	default:
		return errors.New("Unsupported source")
	}
}

// for use as an event source
func (sm *xstatemachine_t) Send(source EventSource, ev Event) error {
	return sm.Dispatch(source, ev)
}

func (sm *xstatemachine_t) InstanceId() string {
	return "sm.instanceid"
}

func (sm *xstatemachine_t) Model() (model Model_xstatemachine_t) {
	return sm.model
}

func (sm *xstatemachine_t) SetModel(model Model_xstatemachine_t) {
	sm.model = model
}

func (sm *xstatemachine_t) ModelName() string {
	return "xstatemachine_t"
}

func (sm *xstatemachine_t) ModelNamespace() string {
	return "MurphyPA::SM"
}

func (sm *xstatemachine_t) ModelGuid() string {
	return "7e46f114-606e-4584-8492-dcf41ee57660"
}

func (sm *xstatemachine_t) ModelFilename() string {
	return "test_machine1.sm1"
}

func (sm *xstatemachine_t) StatemachineVersion() string {
	return "16"
}

func (sm *xstatemachine_t) Comment() string {
	return "None"
}

func (sm *xstatemachine_t) process_event_Source1_x(source *Source1, ev *x) error {
	model := sm.model
	switch sm.State() {
	case A: // bbd966da-e5dc-4e16-9434-09f19ed47960 3
		{
			if model.test2(source, ev) {
				// FROM_LIST: 1 ['T']
				model.a_exit()            // exit-1
				model.t_exit()            // exit-2
				model.action5(source, ev) // transition
				// TO_LIST: 2 ['D', 'H']
				model.d_entry() // entry-1
				model.h_entry() // entry-1
				sm.SetState(H)
			} else if model.test1(source, ev) {
				// FROM_LIST: 1 ['T']
				model.a_exit()            // exit-1
				model.t_exit()            // exit-2
				model.action1(source, ev) // transition
				// TO_LIST: 2 ['D', 'E']
				model.d_entry() // entry-1
				model.e_entry() // entry-1
				model.f_entry() // entry-2
				sm.SetState(F)
			} else { // None
				// FROM_LIST: 1 ['T']
				model.a_exit()            // exit-1
				model.t_exit()            // exit-2
				model.action2(source, ev) // transition
				// TO_LIST: 1 ['D']
				model.d_entry() // entry-1
				model.g_entry() // entry-2
				sm.SetState(G)
			}
		}
		break
	case B: // e5f5bec1-3279-42a2-95da-15d9805aad79 3
		{
			if model.test2(source, ev) {
				// FROM_LIST: 1 ['T']
				model.b_exit()            // exit-1
				model.a_exit()            // exit-1
				model.t_exit()            // exit-2
				model.action5(source, ev) // transition
				// TO_LIST: 2 ['D', 'H']
				model.d_entry() // entry-1
				model.h_entry() // entry-1
				sm.SetState(H)
			} else if model.test1(source, ev) {
				// FROM_LIST: 1 ['T']
				model.b_exit()            // exit-1
				model.a_exit()            // exit-1
				model.t_exit()            // exit-2
				model.action1(source, ev) // transition
				// TO_LIST: 2 ['D', 'E']
				model.d_entry() // entry-1
				model.e_entry() // entry-1
				model.f_entry() // entry-2
				sm.SetState(F)
			} else { // None
				// FROM_LIST: 1 ['T']
				model.b_exit()            // exit-1
				model.a_exit()            // exit-1
				model.t_exit()            // exit-2
				model.action2(source, ev) // transition
				// TO_LIST: 1 ['D']
				model.d_entry() // entry-1
				model.g_entry() // entry-2
				sm.SetState(G)
			}
		}
		break
	case C: // 0b075041-aa7c-4538-b0b2-81db66fc97b0 3
		{
			if model.test2(source, ev) {
				// FROM_LIST: 1 ['T']
				model.c_exit()            // exit-1
				model.b_exit()            // exit-1
				model.a_exit()            // exit-1
				model.t_exit()            // exit-2
				model.action5(source, ev) // transition
				// TO_LIST: 2 ['D', 'H']
				model.d_entry() // entry-1
				model.h_entry() // entry-1
				sm.SetState(H)
			} else if model.test1(source, ev) {
				// FROM_LIST: 1 ['T']
				model.c_exit()            // exit-1
				model.b_exit()            // exit-1
				model.a_exit()            // exit-1
				model.t_exit()            // exit-2
				model.action1(source, ev) // transition
				// TO_LIST: 2 ['D', 'E']
				model.d_entry() // entry-1
				model.e_entry() // entry-1
				model.f_entry() // entry-2
				sm.SetState(F)
			} else { // None
				// FROM_LIST: 1 ['T']
				model.c_exit()            // exit-1
				model.b_exit()            // exit-1
				model.a_exit()            // exit-1
				model.t_exit()            // exit-2
				model.action2(source, ev) // transition
				// TO_LIST: 1 ['D']
				model.d_entry() // entry-1
				model.g_entry() // entry-2
				sm.SetState(G)
			}
		}
		break
	default:
		{
			sm.UnhandledEvent(source, ev)
		}
		break
	}
	return nil
}

func (sm *xstatemachine_t) process_event_Source1_y(source *Source1, ev *y) error {
	model := sm.model
	switch sm.State() {
	// Drop transition: 1a6ca210-bca6-4a41-9567-f3dae0ff5037 in favour of 296f245d-9676-4201-8764-88ff7cfd92d8
	case D: // 7915e533-f297-4ef2-84eb-4a6341ebdbb1 1
		{
			{
				// FROM_LIST: 0 []
				model.d_exit()            // exit-1
				model.action3(source, ev) // transition
				// TO_LIST: 3 ['T', 'A', 'B']
				model.t_entry() // entry-1
				model.a_entry() // entry-1
				model.b_entry() // entry-1
				sm.SetState(B)
			}
		}
		break
	case E: // 4750b88a-14ac-464e-99e8-369fd59a6708 1
		{
			{
				// FROM_LIST: 0 []
				model.e_exit()            // exit-1
				model.d_exit()            // exit-1
				model.action3(source, ev) // transition
				// TO_LIST: 3 ['T', 'A', 'B']
				model.t_entry() // entry-1
				model.a_entry() // entry-1
				model.b_entry() // entry-1
				sm.SetState(B)
			}
		}
		break
	case F: // dd0657e6-adcd-4e94-bca3-5326be56e1e5 1
		{
			{
				// FROM_LIST: 0 []
				model.f_exit()            // exit-1
				model.e_exit()            // exit-1
				model.d_exit()            // exit-1
				model.action3(source, ev) // transition
				// TO_LIST: 3 ['T', 'A', 'B']
				model.t_entry() // entry-1
				model.a_entry() // entry-1
				model.b_entry() // entry-1
				sm.SetState(B)
			}
		}
		break
	case G: // ed8cd3e7-f252-4a9d-8878-e8621f2ad378 1
		{
			{
				// FROM_LIST: 1 ['D']
				model.g_exit()            // exit-1
				model.d_exit()            // exit-2
				model.action6(source, ev) // transition
				// TO_LIST: 2 ['T', 'A']
				model.t_entry() // entry-1
				model.a_entry() // entry-1
				sm.SetState(A)
			}
		}
		break
	case H: // 34e01365-a6de-4d5d-818e-b336e5f885c5 1
		{
			{
				// FROM_LIST: 0 []
				model.h_exit()            // exit-1
				model.d_exit()            // exit-1
				model.action3(source, ev) // transition
				// TO_LIST: 3 ['T', 'A', 'B']
				model.t_entry() // entry-1
				model.a_entry() // entry-1
				model.b_entry() // entry-1
				sm.SetState(B)
			}
		}
		break
	default:
		{
			sm.UnhandledEvent(source, ev)
		}
		break
	}
	return nil
}

func (sm *xstatemachine_t) process_event_Source2_z(source *Source2, ev *z) error {
	model := sm.model
	switch sm.State() {
	case B: // e5f5bec1-3279-42a2-95da-15d9805aad79 1
		{
			{
				model.action4(source, ev) // internal
			}
		}
		break
	case C: // 0b075041-aa7c-4538-b0b2-81db66fc97b0 1
		{
			{
				model.action4(source, ev) // internal
			}
		}
		break
	default:
		{
			sm.UnhandledEvent(source, ev)
		}
		break
	}
	return nil
}

func (sm *xstatemachine_t) dispatch_Source1(source *Source1, event Event) error {
	if !sm.IsInitialised() {
		return errors.New("Statemachine " + sm.ModelName() + ", " + sm.ModelGuid() + sm.InstanceId() + " is not initialised")
	}
	switch ev := event.(type) {
	case *x:
		return sm.process_event_Source1_x(source, ev)
	case *y:
		return sm.process_event_Source1_y(source, ev)
	default:
		return errors.New("Unsupported event")
	}
}

func (sm *xstatemachine_t) dispatch_Source2(source *Source2, event Event) error {
	if !sm.IsInitialised() {
		return errors.New("Statemachine " + sm.ModelName() + ", " + sm.ModelGuid() + sm.InstanceId() + " is not initialised")
	}
	switch ev := event.(type) {
	case *z:
		return sm.process_event_Source2_z(source, ev)
	default:
		return errors.New("Unsupported event")
	}
}

func (sm *xstatemachine_t) initialise_statemachine() error {
	if sm.IsInitialised() {
		return errors.New("Statemachine " + sm.ModelName() + ", " + sm.ModelGuid() + sm.InstanceId() + " is not initialised")
	}
	sm.SetState(TOPSTATE)
	sm.model.t_entry() // initialise by walking default states
	sm.SetState(T)
	sm.model.a_entry() // initialise by walking default states
	sm.SetState(A)
	sm.SetInitialised(true)
	return nil
}

// gen complete -- class  xstatemachine_t

//} // namespace SM
//} // namespace MurphyPA
