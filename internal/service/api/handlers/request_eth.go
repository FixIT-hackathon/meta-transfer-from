package handlers

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"github.com/FixIT-hackathon/meta-transfer-from/internal/service/api/resources"
	"github.com/FixIT-hackathon/meta-transfer-from/pkg/responses"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/jsonapi"
	"math/big"
	"net/http"
)

func RequestEthereum(w http.ResponseWriter, r *http.Request) {
	req, err := resources.NewEthRequest(r)
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: err.Error(),
		}}
		responses.WriteError(w, http.StatusBadRequest, errorsj)
	}
	cli := EthClient(r)

	if checkSenderBalance(w, cli, req) {
		return
	}

	hash, done := sendEthereum(w, cli, &ecdsa.PrivateKey{}, req)
	if !done {
		return
	}

	json.NewEncoder(w).Encode(hash.String())
}

func sendEthereum(w http.ResponseWriter, cli *ethclient.Client,
	privKey *ecdsa.PrivateKey, req *resources.EthRequest) (*common.Hash, bool) {
	nonce, err := cli.PendingNonceAt(context.Background(), req.Sender)
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: "Internal server error",
		}}
		responses.WriteError(w, http.StatusInternalServerError, errorsj)
		return nil, false
	}

	value := big.NewInt(1000000000000000000)
	gasLimit := uint64(21000)
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: "Internal server error",
		}}
		responses.WriteError(w, http.StatusInternalServerError, errorsj)
		return nil, false
	}

	toAddress := req.Sender
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID, err := cli.NetworkID(context.Background())
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: "Internal server error",
		}}
		responses.WriteError(w, http.StatusInternalServerError, errorsj)
		return nil, false
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: "Internal server error",
		}}
		responses.WriteError(w, http.StatusInternalServerError, errorsj)
		return nil, false
	}

	err = cli.SendTransaction(context.Background(), signedTx)
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: "Internal server error",
		}}
		responses.WriteError(w, http.StatusInternalServerError, errorsj)
		return nil, false
	}
	hash := signedTx.Hash()

	return &hash, true
}

func checkSenderBalance(w http.ResponseWriter, cli *ethclient.Client, req *resources.EthRequest) bool {
	balance, err := cli.BalanceAt(context.Background(), req.Sender, nil)
	if err != nil {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: "Internal server error",
		}}
		responses.WriteError(w, http.StatusInternalServerError, errorsj)
		return false
	}

	if balance.Cmp(big.NewInt(0)) != 0 {
		errorsj := []*jsonapi.ErrorObject{{
			Title:  http.StatusText(http.StatusBadRequest),
			Detail: "Ethereum balance > 0",
		}}
		responses.WriteError(w, http.StatusBadRequest, errorsj)
		return false
	}

	return true
}
