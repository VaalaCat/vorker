package rpc

import (
	"errors"
	"vorker/conf"
	"vorker/defs"
	"vorker/models"

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

	reqResp, err := req.C().R().
		SetBody(&defs.AgentSyncWorkersReq{}).
		SetSuccessResult(&rtype).
		SetHeaders(map[string]string{
			defs.HeaderNodeName:   conf.AppConfigInstance.NodeName,
			defs.HeaderNodeSecret: conf.RPCToken,
		}).
		Post(url)
	resp = &rtype.Data
	logrus.Infof("sync agent length: %d", len(resp.WorkerList.Workers))
	// if req is zero, update all workers
	models.SyncWorkers(resp.WorkerList)
	// TODO: support modify single worker

	if err != nil || reqResp.StatusCode >= 299 {
		return errors.New("error")
	}
	return nil
}
