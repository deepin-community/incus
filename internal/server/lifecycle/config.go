package lifecycle

import (
	"github.com/lxc/incus/v6/internal/version"
	"github.com/lxc/incus/v6/shared/api"
)

// ConfigAction represents a lifecycle event action for the server configuration.
type ConfigAction string

// All supported lifecycle events for the server configuration.
const (
	ConfigUpdated = ConfigAction(api.EventLifecycleConfigUpdated)
)

// Event creates the lifecycle event for an action on the server configuration.
func (a ConfigAction) Event(requestor *api.EventLifecycleRequestor, ctx map[string]any) api.EventLifecycle {
	u := api.NewURL().Path(version.APIVersion)

	return api.EventLifecycle{
		Action:    string(a),
		Source:    u.String(),
		Context:   ctx,
		Requestor: requestor,
	}
}
