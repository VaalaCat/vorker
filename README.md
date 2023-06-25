# Vorker

Vorker is a simple self host cloudflare worker alternative which built with cloudflare's [workerd](https://github.com/cloudflare/workerd).

Fearues and Issues are welcome!

![](./images/arch.png)

## Features

- [x] User authentication
- [x] API control
- [x] Multi worker routing
- [x] Woker CRUD Management
- [x] Web UI & Online Editor
- [x] Multi Node
- [ ] Log
- [ ] Metrics
- [ ] Worker version control
- [ ] Worker Debugging
- [ ] Support KV storage
- [ ] HA support

## Screenshots

- Admin Page

![](./images/worker-admin.png)

- Worker Editor

![](./images/worker-edit.png)

- Worker Config

![](./images/worker-config.png)

- Agent Status

![](./images/status.png)

- Worker Execution

![](https://vaala.cat/images/vorkerexec.png)

## Usage

### Docker

1. Run by docker command or download the docker-compose.yml from repo and execute `docker-compose up -d`.

all envs defined in [env.go](./conf/env.go), you can take a look at it for more details.

```bash
docker run -dit --name=vorker \
	-e WORKER_URL_SUFFIX=.example.com \
	-e COOKIE_DOMAIN=example.com \
	-e ENABLE_REGISTER=false \
	-e COOKIE_NAME=authorization \
	-e JWT_SECRET=xxxxxxx \
	-e JWT_EXPIRETIME=6 \
	-e AGENT_SECRET=xxxxxxx \
	-p 8080:8080 \
	-p 8888:8888 \
	-p 18080:18080 \
	-v /tmp/workerd:/workerd \
	vaalacat/vorker:latest

# this is a example, you can change the env to fit your need
# for this example, you can visit http://localhost:8888/admin to access the web ui
# and the worker URL will be: SCHEME://WORKER_NAME.example.com
```

2. test your workerd, if your vorker is running on localhost, you can use curl to test it.

visit `http://localhost:8888/admin` to control your worker.

```bash
curl localhost:8080 -H "Host: workername" # replace workername with your worker name
```

4. enjoy your untimate self hosted worker!