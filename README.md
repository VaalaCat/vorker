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
- [ ] Log
- [ ] Metrics
- [ ] Worker version control
- [ ] Worker Debugging
- [ ] Support KV storage
- [ ] Multi Node HA support

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

1. Download the latest release workerd from [here](https://github.com/cloudflare/workerd/releases/) and place it in a blank directory.

example:
```bash
mkdir vorker && cd vorker
curl -fSL -O https://github.com/cloudflare/workerd/releases/download/v1.20230518.0/workerd-linux-64.gz
gzip -d workerd-linux-64.gz
chmod +x workerd-linux-64
```

2. Run by docker command or download the docker-compose.yml from repo and execute `docker-compose up -d`.

```bash
docker run -dit --name=vorker \
	-e DB_PATH=/path/to/workerd/db.sqlite \
	-e WORKERD_DIR=/path/to/workerd \
	-e WORKERD_BIN_PATH=/bin/workerd \
	-e DB_TYPE=sqlite \
	-e WORKER_LIMIT=10000 \
	-e WORKER_PORT=8080 \
	-e API_PORT=8888 \
	-e LISTEN_ADDR=0.0.0.0 \
	-e WORKER_URL_SUFFIX=.example.com \ # concat with worker name and scheme
	-e SCHEME=http \ # external scheme
	-e ENABLE_REGISTER=false \
	-e COOKIE_NAME=authorization \
	-e COOKIE_AGE=21600 \
	-e COOKIE_DOMAIN=localhost # change it to your domain \
	-e JWT_SECRET=xxxxxxx \
	-e JWT_EXPIRETIME=6 \
	-p 8080:8080 \
	-p 8888:8888 \
	-v $PWD/workerd:/path/to/workerd \
	-v $PWD/workerd-linux-64:/bin/workerd \
	vaalacat/vorker:latest
# this is a example, you can change the env to fit your need
# for this example, you can visit http://localhost:8888/admin to access the web ui
# and the worker URL will be: SCHEME://WORKER_NAME.example.com
```

3. test your workerd, if your vorker is running on localhost, you can use curl to test it.

visit `http://localhost:8888/admin` to control your worker.

```bash
curl localhost:8080 -H "Host: workername" # replace workername with your worker name
```

4. enjoy your untimate self hosted worker!