# Auth Gateway


## Hỏi đáp

### Tại sao cần API Gateway?

### Tại sao cần Single Signon?

### Khác biệt giữa Session Cookie với JWT Token?

### Single Point of Failure là gì?

### Proxy khác gì với Reverse Proxy?

## Docker

### Lệnh docker khác gì lệnh docker-compose?

### Portainer để làm gì?

### -p 80:8080 có ý nghĩa gì? 80 là cổng nào và 8080 là cổng nào?

### docker build cần có những tham số gì?

### Khác biệt giữa docker image và docker container là gì?

### Dockerfile khác gì với docker-compose.yml

### Muốn service A trong docker-compose được khởi động sau khi service B đã chạy thì cần thêm cấu hình gì?

### Lệnh tạo một vùng lưu trữ trên host để docker container ghi vào là gì?

### Ý nghĩa của lệnh này /var/run/docker.sock:/var/run/docker.sock

### docker-compose có lệnh để build lại những docker image cần dùng cho service, lệnh đó là lệnh nào?

## Traefik

### Providers trong traefik gồm những loại nào?

### Docker Provider khác gì với File Provider?

### Static File Configuration và Dynamic File Configuration có những đặc điểm và tác dụng gì?

### Làm thể nào để bảo mật được Traefik DashBoard?

### EntryPoint là gì? và có những kiểu nào và thuộc tính nào đi cùng?

### Ý nghĩa của lệnh này "traefik.http.services.mainsite.loadbalancer.server.port=9001"

### Trong cấu hình Docker provider sử dụng docker-compose.yml, thì lệnh labels có tác dụng gì?

### Tên của docker container có bắt buộc phải trùng tên với tên của service không?

### Ý nghĩa của lệnh này "traefik.http.routers.mainsite.rule=Host(`iris.com`)"

### traefik.yml

- Ý nghĩa của lệnh này api:  
    insecure: true

- Ý nghĩa của accessLog{}

### traefik_dynamic.yml

- Trong một cấu hình router có 3 thuộc tính: rule, middlewares và service. Chức năng của từng thuộc tính là gì?

- Giải thích rule: "Host(`iris.com`) && !PathPrefix(`/public`)"

- Giải thích service: mainsite@docker

### Một request đến Entry Point trước hay đến Router trước?

### Ứng với một service có thể định nghĩa nhiều Router được không?

### docker-compose.yml, dịch vụ auth hãy giải thích 

- "traefik.http.middlewares.auth.forwardauth.address=http://auth:3000/auth"

- "traefik.http.middlewares.auth.forwardauth.trustForwardHeader=true"

- "traefik.http.middlewares.auth.forwardauth.authResponseHeaders=X-Forwarded-User"

- "traefik.http.services.auth.loadbalancer.server.port=3000"

