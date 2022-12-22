# Description
Demo for microservices written in Go.

# Requisites
 - Docker / Docker Compose
 - :)

# How to Run
### 1. hosts file
Make sure you chage your `/etc/hosts/` to have `backend` as `127.0.0.1`:
```
127.0.0.1   localhost backend
```
### 2. Run containers
Go into `project` folder and run:
```
docker compose up -d
```

### 3. Try it out
Access the front-end (see below) and click the buttons to test the communication the microservices.

---

### Closing

Once you're done, you can stop the containers by going inside the project folder again and running:
```
docker compose down
```
## Run with Swarm
TODO

# Useful addresses:
 - Front-end: http://localhost
 - Mailhog: http://localhost:8025

# Things to enhance
 - This readme (swarm; /etc/hosts for Caddy; info about vscode workspace)
 - Move some functions to libraries to be reused across microservices
 - Optimize dockerfiles multistage for caching