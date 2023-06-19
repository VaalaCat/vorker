package rpc

import (
	"errors"
	"fmt"
	"voker/conf"
	"voker/defs"
	"voker/utils"

	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

func SyncAgentRequest(nodeName string) error {
	return nil
}

func SyncAgent(endpoint string) error {
	url := endpoint + "/api/agent/sync"
	resp := &defs.AgentSyncWorkersResp{}

	token, err := utils.HashPassword(fmt.Sprintf("%s%s",
		conf.AppConfigInstance.NodeName,
		conf.AppConfigInstance.AgentSecret))
	if err != nil {
		return err
	}

	reqResp, err := req.C().R().
		SetBody(&defs.AgentSyncWorkersReq{}).
		SetSuccessResult(resp).
		SetHeaders(map[string]string{
			defs.HeaderNodeName:   conf.AppConfigInstance.NodeName,
			defs.HeaderNodeSecret: token,
		}).
		Post(url)

	logrus.Infof("SyncAgent: %+v,err: %v", reqResp, err)

	if err != nil || reqResp.StatusCode >= 299 {
		return errors.New("error")
	}
	return nil
}
