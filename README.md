## Centrifugo get started project

### Prerequisites
- Docker
- golang

### To run centrifugo server 

First, you need to generate a configuration file.
```bash
docker run --rm -v$PWD:/centrifugo centrifugo/centrifugo:v5 centrifugo genconfig
```

Then you can run the server.
```bash
docker run --rm --ulimit nofile=262144:262144 -v /host/dir/with/config/file:/centrifugo -p 8000:8000 centrifugo/centrifugo:v5 centrifugo -c config.json
```

or just run docker-compose
```bash
docker-compose up
```

then open `http://localhost:8000` in your browser.
password would be `password` if you use the docker-compose, otherwise, you can find it in the config file.

use this jwt token for testing
```bash
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM3MjIiLCJleHAiOjE4NTU0NDgyOTl9.l49QtxVghR2wgg9-FmZycLnlicFqMo_uKd-LZ7qtiqU
```
