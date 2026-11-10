package enum

type OrderState int

const (
	StateCreated OrderState = iota
	StatePaid
	StatePacked
	StateShipped
	StateDelivered
	StateCancelled
	StateReturned
)

func (s OrderState) String() string {
	stateNames := map[OrderState]string{
		StateCreated:   "created",
		StatePaid:      "paid",
		StatePacked:    "packed",
		StateShipped:   "shipped",
		StateDelivered: "delivered",
		StateCancelled: "cancelled",
		StateReturned:  "returned",
	}

	name, ok := stateNames[s]
	if !ok {
		return "unknown"
	}

	return name
}

func NextState(current OrderState, action string) OrderState {
	// Validate the current state.
	if current < StateCreated || current > StateReturned {
		panic("invalid order state")
	}

	// Define the allowed transitions.
	transitions := map[OrderState]map[string]OrderState{
		StateCreated: {
			"pay":    StatePaid,
			"cancel": StateCancelled,
		},
		StatePaid: {
			"pack":   StatePacked,
			"cancel": StateCancelled,
		},
		StatePacked: {
			"ship": StateShipped,
		},
		StateShipped: {
			"deliver": StateDelivered,
			"return":  StateReturned,
		},
		StateDelivered: {
			"return": StateReturned,
		},
	}

	next, ok := transitions[current][action]
	if !ok {
		return current
	}

	return next
}
