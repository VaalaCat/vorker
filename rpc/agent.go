package rpc

import (
	"errors"
	"fmt"
	"vorker/conf"
	"vorker/defs"
	"vorker/utils"

	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

func SyncAgentRequest(nodeName string) error {
	return nil
}

func SyncAgent(endpoint string) error {
	url := endpoint + "/api/agent/sync"
	resp := &defs.AgentSyncWorkersResp{}
	rtype := struct {
		Code int                       `json:"code"`
		Msg  string                    `json:"msg"`
		Data defs.AgentSyncWorkersResp `json:"data"`
	}{}

	token, err := utils.HashPassword(fmt.Sprintf("%s%s",
		conf.AppConfigInstance.NodeName,
		conf.AppConfigInstance.AgentSecret))
	if err != nil {
		return err
	}

	reqResp, err := req.C().R().
		SetBody(&defs.AgentSyncWorkersReq{}).
		SetSuccessResult(&rtype).
		SetHeaders(map[string]string{
			defs.HeaderNodeName:   conf.AppConfigInstance.NodeName,
			defs.HeaderNodeSecret: token,
		}).
		Post(url)
	resp = &rtype.Data
	logrus.Infof("sync agent length: %d", len(resp.WorkerList.Workers))

	if err != nil || reqResp.StatusCode >= 299 {
		return errors.New("error")
	}
	return nil
}
