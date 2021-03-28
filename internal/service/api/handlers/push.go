package handlers

import (
	"encoding/hex"
	"fmt"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/data"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api/resources"
	"github.com/FixIT-hackathon/meta-transfer-from/pkg/responses"
	"github.com/FixIT-hackathon/meta-transfer-from/pkg/signature"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/jsonapi"
	"net/http"
	"strings"
)

func Push(w http.ResponseWriter, r *http.Request) {
	req, err := resources.NewPushRequest(r)
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}}
		responses.WriteError(w, http.StatusBadRequest, errorsj)
		return
	}

	typedDataHash, _ := req.TypedData.HashStruct(req.TypedData.PrimaryType, req.TypedData.Message)
	domainSeparator, _ := req.TypedData.HashStruct("EIP712Domain", req.TypedData.Domain.Map())
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	challengeHash := crypto.Keccak256Hash(rawData)

	addr, err := signature.RecoverAddress(challengeHash.Bytes(), req.Signature)
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}}
		responses.WriteError(w, http.StatusBadRequest, errorsj)
		return
	}

	if strings.Compare(strings.ToLower(addr.String()), strings.ToLower(req.Sender)) != 0 {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: "Sender and signer is not equal",
		}}
		responses.WriteError(w, http.StatusBadRequest, errorsj)
		return
	}

	sigB, _ := hex.DecodeString(req.Signature)
	param, _ := signature.ParseSignatureParameters(sigB)

	_, err = TransfersQ(r).Create(data.Transfer{
		Sender:   req.Sender,
		Receiver: req.Receiver,
		Fee:      "100",
		Status:   "pending",
		Amount:   req.Amount,
		ERC20:    req.ERC20,

		R: param.R,
		S: param.S,
		V: param.V,
	})
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusInternalServerError),
			Detail: "failed to add to db",
		}}
		responses.WriteError(w, http.StatusInternalServerError, errorsj)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
