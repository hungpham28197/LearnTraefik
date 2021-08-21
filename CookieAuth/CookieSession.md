
Hướng dẫn cài đặt Portainer

https://documentation.portainer.io/v2.0/deploy/ceinstalldocker/

```
docker volume create portainer_data
docker run -d -p 8000:8000 -p 9000:9000 --name=portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce
```

```
cd CookieAuth
docker-compose up -d
```
