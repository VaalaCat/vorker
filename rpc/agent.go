package rpc

import (
	"errors"
	"fmt"
	"vorker/conf"
	"vorker/defs"
	"vorker/entities"
	"vorker/utils"

	"github.com/imroc/req/v3"
	"github.com/sirupsen/logrus"
)

func EventNotify(n *entities.Node, eventName string, extra map[string][]byte) error {
	reqResp, err := RPCWrapper().
		SetHeader(defs.HeaderHost, utils.NodeHost(n.Name, n.UID)).
		SetBody(&entities.NotifyEventRequest{EventName: eventName, Extra: extra}).
		Post(
			fmt.Sprintf("http://%s:%d/api/agent/notify",
				conf.AppConfigInstance.TunnelHost,
				conf.AppConfigInstance.TunnelEntryPort))

	if err != nil || reqResp.StatusCode >= 299 {
		logrus.Errorf("event notify error, err: %+v, resp: %+v", err, reqResp)
		return errors.New("error")
	}
	return nil
}

func SyncAgent(endpoint string) (*entities.WorkerList, error) {
	url := endpoint + "/api/agent/sync"
	resp := &entities.AgentSyncWorkersResp{}
	rtype := struct {
		Code int                           `json:"code"`
		Msg  string                        `json:"msg"`
		Data entities.AgentSyncWorkersResp `json:"data"`
	}{}

	reqResp, err := RPCWrapper().
		SetBody(&entities.AgentSyncWorkersReq{}).
		SetSuccessResult(&rtype).
		Post(url)
	resp = &rtype.Data
	logrus.Infof("sync agent length: %d", len(resp.WorkerList.Workers))

	if err != nil || reqResp.StatusCode >= 299 {
		return nil, errors.New("error")
	}
	return resp.WorkerList, nil
}

func AddNode(endpoint string) error {
	url := endpoint + "/api/agent/add"
	rtype := struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}{}

	reqResp, err := RPCWrapper().
		SetBody(&entities.AgentSyncWorkersReq{}).
		SetSuccessResult(&rtype).
		Post(url)

	if err != nil || reqResp.StatusCode >= 299 {
		return errors.New("error")
	}
	return nil
}

func GetNode(endpoint string) (*entities.Node, error) {
	url := endpoint + "/api/agent/nodeinfo"
	rtype := struct {
		Code int            `json:"code"`
		Msg  string         `json:"msg"`
		Data *entities.Node `json:"data"`
	}{}

	reqResp, err := RPCWrapper().
		SetSuccessResult(&rtype).
		Get(url)

	if err != nil || reqResp.StatusCode >= 299 {
		return nil, errors.New("error")
	}
	return rtype.Data, nil
}

func RPCWrapper() *req.Request {
	return req.C().R().
		SetHeaders(map[string]string{
			defs.HeaderNodeName:   conf.AppConfigInstance.NodeName,
			defs.HeaderNodeSecret: conf.RPCToken,
		})
}
