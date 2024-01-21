package main

// circle trace
var (
	circleTrace     []float32
	showCircleTrace bool
)

func updateCircleTrace(x float32, y float32) {
	circleTraceLen := len(circleTrace)

	if circleTraceLen > 8192 {
		circleTrace = []float32{x, y, x, y}
		circleTraceLen = len(circleTrace)
	}
	circleTrace = append(circleTrace, circleTrace[circleTraceLen-2], circleTrace[circleTraceLen-1])
	circleTrace = append(circleTrace, x, y)
}
