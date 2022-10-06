# avalancheseqlogger

When you work on custom subnets using avalanche-network-runner, you probably have logs all over the place.
You can have all your logs in one place using centralized logging solutions.

This is a tutorial of how to use [Seq](https://datalust.co/).

In this screenshot, you can see the logs from 4 nodes and 2 processes (avalancheGo and SubDotNetVm). You can filter the logs, expand lines to see the structured logs, and there is also a "tail" like mode.

_Centralized logging using Seq_
![seq-running](https://user-images.githubusercontent.com/1549198/194254189-3b6d3bdd-3612-4b0f-a151-eb6aadbd8846.png)


I made very minimal modifications to avalanchego (4 lines of code in log/factory.go).

# How to run 
1) You need to run seq locally.

It's very easy to run seq in your local dev environment using docker.
```shell
docker run \
  --name seqtmp \
  -d \
  -e ACCEPT_EULA=Y \
  -v /tmp/seq:/data \
  -p 18080:80 \
  -p 5341:5341 \
  datalust/seq
```
or if you use podman
```shell
podman run \
  --name seqtmp \
  -d \
  -e ACCEPT_EULA=Y \
  -v /tmp/seq:/data:Z \
  -p 18080:80 \
  -p 5341:5341 \
  datalust/seq
```

2) Clone and build this modified version of [avalanchego](https://github.com/subdotnet/avalanchego)

In this exemple, we will clone it in `/tmp` but you can choose any folder you like.
```shell
cd /tmp
git clone https://github.com/subdotnet/avalanchego
cd avalanchego
./scripts/build.sh
```
The compiled binary path is `/tmp/avalanchego/build/avalanchego`.

3) Install and run avalanche-network-runner

Install using these [instructions](https://github.com/ava-labs/avalanche-network-runner#installation)
You will need two terminals for this : 
Start avalanche-network-runner in one terminal 
```shell
avalanche-network-runner server
```

Start the nodes in another terminal
```shell
avalanche-network-runner control start \
--log-level info \
--number-of-nodes=4 \
--avalanchego-path /tmp/avalanchego/build/avalanchego
```

4) Watch the logs 

Open this url http://localhost:18080/ in your browser


