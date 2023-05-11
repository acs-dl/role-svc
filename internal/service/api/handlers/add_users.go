package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	helpers2 "github.com/acs-dl/role-svc/internal/service/api/helpers"
	"github.com/acs-dl/role-svc/internal/service/api/requests"
	"github.com/acs-dl/role-svc/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func AddUsers(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewCreateRequestRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse create request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	for i, user := range request.Data.Attributes.Users {
		payload, err := helpers2.CreateAddUserPayload(request.Data.Attributes.Link, user, request.Data.Attributes.AccessLevel)
		if err != nil {
			Log(r).WithError(err).Errorf("failed to create payload")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		requestBody, err := createRequestBody(request.Data, i, payload)
		if err != nil {
			Log(r).WithError(err).Errorf("failed to create request body")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		err = helpers2.SendRequestToOrchestrator(LinksParams(r).Orchestrator, requestBody, r.Header.Get("Authorization"))
		if err != nil {
			Log(r).WithError(err).Errorf("failed to send request to orchestrator")
			ape.RenderErr(w, problems.InternalError())
			return
		}
	}

	Log(r).Infof("successfully created adding requests for `%d` users", len(request.Data.Attributes.Users))
	w.WriteHeader(http.StatusAccepted)
	ape.Render(w, newRequestResponse(request.Data))
}

func newRequestResponse(request resources.Request) resources.RequestResponse {
	return resources.RequestResponse{
		Data: request,
	}
}

func createRequestBody(request resources.Request, counter int, payload json.RawMessage) ([]byte, error) {
	return json.Marshal(struct {
		Data resources.OrchestratorRequest `json:"data"`
	}{
		Data: resources.OrchestratorRequest{
			Key: resources.Key{
				ID:   fmt.Sprintf("%d-%s", counter, request.ID),
				Type: resources.REQUESTS,
			},
			Attributes: resources.OrchestratorRequestAttributes{
				FromUser: request.Attributes.FromUser,
				ToUser:   request.Attributes.ToUser,
				Module:   request.Attributes.Module,
				Payload:  payload,
			},
			Relationships: resources.OrchestratorRequestRelationships{
				User: resources.Relation{},
			},
		},
	})
}
