package adapters

import (
	"github.com/cloudevents/sdk-go/v2/event"
	"github.com/nebula-aac/kubernexus-mono/api/protos/meshes"
)

func (h *Adapter) StreamErr(e *meshes.CloudEventsResponse, err error) {
	h.Log.Error(err)
	ce := convertToCloudEvent(e)
	go func() {
		h.CloudEventsStreamer.Publish(ce)
		h.Log.Info("Event stored and sent successfully")
	}()
}

func (h *Adapter) StreamInfo(e *meshes.CloudEventsResponse) {
	h.Log.Info("Sending event")
	ce := convertToCloudEvent(e)
	go func() {
		h.CloudEventsStreamer.Publish(ce)
		h.Log.Info("Event stored and sent successfully")
	}()
}

func convertToCloudEvent(e *meshes.CloudEventsResponse) event.Event {
	ce := event.New()
	ce.SetType(e.EventType.String())
	ce.SetSource("meshery-cloudevents")
	ce.SetExtension("summary", e.Summary)
	// Set other desired CloudEvent fields
	return ce
}
