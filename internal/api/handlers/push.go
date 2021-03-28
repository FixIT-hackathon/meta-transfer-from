package handlers

import (
	"encoding/json"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/api/resources"
	"github.com/FixIT-hackathon/meta-transfer-from/pkg/responses"
	"github.com/ethereum/go-ethereum/common/math"
	signer "github.com/ethereum/go-ethereum/signer/core"
	"github.com/google/jsonapi"
	"net/http"
)

func Craft(w http.ResponseWriter, r *http.Request) {
	req, err := resources.NewCraftRequest(r)
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}}
		responses.WriteError(w, http.StatusBadRequest, errorsj)
		return
	}

	chainID := math.HexOrDecimal256(req.ChainID)

}
