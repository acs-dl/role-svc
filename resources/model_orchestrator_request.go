/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type OrchestratorRequest struct {
	Key
	Attributes    OrchestratorRequestAttributes    `json:"attributes"`
	Relationships OrchestratorRequestRelationships `json:"relationships"`
}
type OrchestratorRequestResponse struct {
	Data     OrchestratorRequest `json:"data"`
	Included Included            `json:"included"`
}

type OrchestratorRequestListResponse struct {
	Data     []OrchestratorRequest `json:"data"`
	Included Included              `json:"included"`
	Links    *Links                `json:"links"`
}

// MustOrchestratorRequest - returns OrchestratorRequest from include collection.
// if entry with specified key does not exist - returns nil
// if entry with specified key exists but type or ID mismatches - panics
func (c *Included) MustOrchestratorRequest(key Key) *OrchestratorRequest {
	var orchestratorRequest OrchestratorRequest
	if c.tryFindEntry(key, &orchestratorRequest) {
		return &orchestratorRequest
	}
	return nil
}
