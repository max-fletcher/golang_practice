package main

import "fmt"

// 🧭 1. Switch-based behavior (most common)
type Status int

const (
	Pending Status = iota
	Running
	Completed
	Failed
)

func handleStatus(s Status) {
	switch s {
	case Pending:
		fmt.Println("Waiting to start...")
	case Running:
		fmt.Println("In progress...")
	case Completed:
		fmt.Println("Done!")
	case Failed:
		fmt.Println("Something went wrong")
	default:
		fmt.Println("Unknown status")
	}
}

// 🚦 2. State machine logic
// 	Very common in:

// 	workflows
// 	job processing
// 	pipelines

func nextState(s Status) Status {
	switch s {
	case Pending:
		return Running
	case Running:
		return Completed
	default:
		return Failed
	}
}

// 🔐 3. Permission checks (bit flags — very real-world)
// 👉 This pattern is used in:

// file systems
// auth systems
// feature toggles

type Permission int

const (
	Read Permission = 1 << iota
	Write
	Execute
)

func canWrite(p Permission) bool {
	return p&Write != 0
}

// usage
func checkPermission(){
	p := Read | Write

	if canWrite(p) {
		fmt.Println("User can write")
	}
}

// 🧪 4. Validation logic
func isValidStatus(s Status) bool {
	switch s {
	case Pending, Running, Completed, Failed:
		return true
	default:
		return false
	}
}

// 🔄 5. Mapping enum → string (logging, APIs)

func (s Status) String() string {
	switch s {
	case Pending:
		return "pending"
	case Running:
		return "running"
	case Completed:
		return "completed"
	case Failed:
		return "failed"
	default:
		return "unknown"
	}
}

// Usage:
fmt.Println("Status:", s.String())

// 🔁 6. Parsing string → enum (very common in APIs)

func parseStatus(s string) (Status, error) {
	switch s {
	case "pending":
		return Pending, nil
	case "running":
		return Running, nil
	case "completed":
		return Completed, nil
	case "failed":
		return Failed, nil
	default:
		return 0, fmt.Errorf("invalid status: %s", s)
	}
}

// ⚙️ 7. Controlling behavior (strategy-like logic)
type Operation int

const (
	Add Operation = iota
	Subtract
	Multiply
)

func calculate(op Operation, a, b int) int {
	switch op {
	case Add:
		return a + b
	case Subtract:
		return a - b
	case Multiply:
		return a * b
	default:
		return 0
	}
}

// 🧵 8. Controlling concurrency behavior
type Mode int

const (
	Sequential Mode = iota
	Parallel
)

func process(mode Mode) {
	if mode == Parallel {
		go fmt.Println("Running in parallel")
	} else {
		fmt.Println("Running sequentially")
	}
}

// 🧩 9. Feature flags (bitmask + logic)
type Feature int

const (
	FeatureA Feature = 1 << iota
	FeatureB
	FeatureC
)

func hasFeature(f Feature, flag Feature) bool {
	return f&flag != 0
}

// 🚨 10. Error handling based on enum
type ErrorCode int

const (
	NotFound ErrorCode = iota
	Unauthorized
	ServerError
)

func handleError(code ErrorCode) {
	switch code {
	case NotFound:
		fmt.Println("404 Not Found")
	case Unauthorized:
		fmt.Println("401 Unauthorized")
	case ServerError:
		fmt.Println("500 Internal Error")
	}
}

// 🧠 11. Guard clauses (cleaner logic)
func processStatus(s Status) error {
	if s == Failed {
		return fmt.Errorf("cannot process failed status")
	}
	if s == Completed {
		return fmt.Errorf("already completed")
	}
	return nil
}