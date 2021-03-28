package handlers

import (
	"encoding/json"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api/resources"
	"github.com/FixIT-hackathon/meta-transfer-from/pkg/responses"
	"github.com/google/jsonapi"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	_, err := resources.NewListRequest(r)
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}}
		responses.WriteError(w, http.StatusBadRequest, errorsj)
		return
	}

	transfers, err := TransfersQ(r).Select()
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: err.Error(),
		}}
		responses.WriteError(w, http.StatusInternalServerError, errorsj)
		return
	}

	resp := make([]resources.Transfer, len(transfers))
	for i, transfer := range transfers {
		resp[i] = *transfer.Resource()
	}

	json.NewEncoder(w).Encode(resp)
}
