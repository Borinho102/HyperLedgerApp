package chaincode

import (
	"net/http"
	"os"

	"github.com/goledgerdev/ccapi/common"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
)

func Invoke(channelName, ccName, txName string, txArgs [][]byte, transientRequest []byte) (*channel.Response, int, error) {
	// create channel manager
	fabMngr, err := common.NewFabricChClient(channelName, os.Getenv("USER"), os.Getenv("ORG"))
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	// Execute chaincode with channel's client
	rq := channel.Request{ChaincodeID: ccName, Fcn: txName, Args: txArgs}

	if len(transientRequest) != 0 {
		transientMap := make(map[string][]byte)
		transientMap["@request"] = transientRequest
		rq.TransientMap = transientMap
	}

	res, err := fabMngr.Client.Execute(rq, channel.WithRetry(retry.DefaultChannelOpts))

	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	return &res, http.StatusInternalServerError, nil
}
