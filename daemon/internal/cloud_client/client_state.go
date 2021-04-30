package cloud_client

import (
	"github.com/akitasoftware/akita-cli/rest"
	"github.com/akitasoftware/akita-libs/akid"
	"github.com/akitasoftware/akita-libs/daemon"
)

// Logging state for a single service.
type serviceInfo struct {
	// The learning client for this service.
	learnClient rest.LearnClient

	// Contains an entry for each trace ID for which we are collecting events.
	traces map[akid.LearnSessionID]*traceInfo

	// Contains channels to clients waiting to hear about new traces.
	responseChannels []chan<- []daemon.LoggingOptions
}

func (client *cloudClient) newServiceInfo(serviceID akid.ServiceID) *serviceInfo {
	return &serviceInfo{
		learnClient:      client.newLearnClient(serviceID),
		traces:           map[akid.LearnSessionID]*traceInfo{},
		responseChannels: []chan<- []daemon.LoggingOptions{},
	}
}

// Logging state for a single trace.
type traceInfo struct {
	// Whether the trace is active. If this is false, then the daemon is just
	// waiting for clients to finish sending their events.
	active bool

	// The names of the clients from which we have received trace events and have
	// not signalled the end of their event stream.
	clientNames map[string]struct{}

	// The trace's logging options.
	loggingOptions daemon.LoggingOptions

	// The channel on which to send trace events to the trace collector.
	traceEventChannel chan<- *TraceEvent

	// Channels to clients waiting to hear about the deactivation of the trace.
	deactivationChannels []chan<- struct{}
}

func newTraceInfo(loggingOptions daemon.LoggingOptions, traceEventChannel chan<- *TraceEvent) *traceInfo {
	return &traceInfo{
		active:               true,
		clientNames:          map[string]struct{}{},
		loggingOptions:       loggingOptions,
		traceEventChannel:    traceEventChannel,
		deactivationChannels: []chan<- struct{}{},
	}
}
