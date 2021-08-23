# Gateway Auth


## Hỏi đáp

### Tại sao cần API Gateway?
- Trong hệ thống microservices có rất nhiều service, nếu như client mà gọi trực tiếp đến các service đằng sau thi hệ thống sẽ rất phức tạp, API Gateway sẽ có vai trò là cổng trung gian để nhận request từ client và gửi đến service tương ứng đằng sau.
- Lợi ích: dễ dàng quản lý các request từ client. Có thể thêm các chức năng cân bằng tải, xác thực, phân quyền, thêm cơ chế bảo mật,...
- Nhược điểm là thêm chi phí vận hành, tăng thời gian response, gây nghẽn cổ chai.

### Tại sao cần Single Signon?
- Các hệ thống cùng cơ chế xác thực mà mỗi lần gọi đến lại phải xác thực lại từ đầu thì sẽ phức tạp và mất thời gian, do đó ý tưởng là xác thực 1 lần và có thể thao tác với các hệ thống vệ tinh cùng cơ chế xác thực.

### Khác biệt giữa Session Cookie với JWT Token?
- Session Cookie lưu thông tin người dùng đã đăng nhập thành công ở phía server (file, database). Khi người dùng gửi request, server sẽ get thông tin theo session-id, những thông tin này được lưu trữ lâu dài có thể gây phình to bộ nhớ, và có thể tạo lỗ hổng bảo mật. Cần xóa những session cũ đi và lưu thông tin user ngắn gọn đủ dùng.
- JWT không lưu thông tin người dùng phía server, mà server sẽ giải mã payload trong chuỗi jwt người dùng gửi lên để lấy thông tin người dùng. Khi thông tin được mã hóa quá nhiều trong payload sẽ làm cho size của request tăng lên rất nhiều, ảnh hưởng đến traffic.

### Single Point of Failure là gì?
- Là 1 hệ thống con bị lỗi hoặc 1 lỗi nhỏ ở trong hệ thống bị lỗi nhưng lại gây ra sập toàn bộ hệ thống.

### Proxy khác gì với Reverse Proxy?
- Proxy là cách gọi đơn giản của Forward proxy, nó được dùng ở Client(trình duyệt), chuyển tiếp request của người dùng đến Internet, nó có thể ẩn đi IP hoặc thay đổi IP cho mỗi request. Proxy cũng có thể lọc những request từ clients. Mọi requests ra ngoài hay phản hồi từ Internet đều phải thông qua Proxy. Proxy có thể cache lại để lần sau khi có yêu cầu tải về resource đó, proxy sẽ trả về nội dung được cache.
- Reverse Proxy chủ yếu được dùng ở phía server để cân bằng tải, tính sẵn sàng, xác thực, chặn truy cập. Trong hệ thống sẽ có thể có nhiều service sau Reverse Proxy. Khi nhận request từ client nó sẽ chuyển tiếp đến 1 trong các service tương ứng.

## Traefik
- Là open-source golang, đóng vai trò là 1 reverse-proxy, tích hợp với các công nghệ cluster Docker, Docker Swarm, K8S. Nhận request từ client qua internet và chuyển tiếp đến các service container tương ứng. Có thể định nghĩa các middleware tùy chỉnh.

### Providers trong traefik gồm những loại nào?
Có 3 loại: Orchestrator, Manual, KeyValue
docker
Kubernetes
Consul 
ECS
Marathon
Rancher
File
Consul
Etcd
ZooKeeper
Redis
HTTP

### Docker Provider khác gì với File Provider?
- Docker Provider sử dụng label để define các config trong mỗi container tương ứng với các service như router, rule, middleware, service.
- File Provider sử dụng file YAML TOML để define các config như router, rule, middleware, service.

### Static File Configuration và Dynamic File Configuration có những đặc điểm và tác dụng gì?
- Static File Configuration khai báo những config cần có khi start traefik gateway, nó thường để config những tham số global trong hệ thống.
- Dynamic File Configuration khai báo những config trong khi running, nó thường để config những tham số nhưng routers, services,...

### Làm thể nào để bảo mật được Traefik DashBoard?
- thêm middleware BasicAuth khi router đến Traefik DashBoard

### EntryPoint là gì? và có những kiểu nào và thuộc tính nào đi cùng?
- Xác định điểm vào mà Traefik lắng nghe(port), không thường xuyên thay đổi.
- HTTP, TCP, UDP.

### Ý nghĩa của lệnh này "traefik.http.services.mainsite.loadbalancer.server.port=9001"
- Chỉ định port mà traefik giao tiếp với service

### Trong cấu hình Docker provider sử dụng docker-compose.yml, thì lệnh labels có tác dụng gì?
- labels dùng để cấu hình các tham số của traefik trên container
### Tên của docker container có bắt buộc phải trùng tên với tên của service không? 
- Không

### Ý nghĩa của lệnh này "traefik.http.routers.mainsite.rule=Host(`iris.com`)"
- Xác định đường dẫn hợp lệ để router call để service mainsite
### traefik.yml

- Ý nghĩa của lệnh này api:  insecure: true
Nếu entryPoint traefik không được định cấu hình, nó sẽ tự động được tạo trên cổng 8080.

- Ý nghĩa của accessLog{}: ghi log để biết tầng nào gọi đến tầng nào

### traefik_dynamic.yml

- Trong một cấu hình router có 3 thuộc tính: rule, middlewares và service. Chức năng của từng thuộc tính là gì?
    + rule: đường dẫn hợp lệ để router định tuyến xuống service
    + middlewares: trước khi router gọi xuống service thì cần qua middlewares được chỉ định
- Giải thích rule: "Host(`iris.com`) && !PathPrefix(`/public`)"
    + tất cả đường dẫn có domain là iris.com và path khác /public là hợp lệ. (iris.com/public là ko hợp lệ)

- Giải thích service: mainsite@docker
  + container: mainsite
  + provider: docker

### Một request đến Entry Point trước hay đến Router trước?
- Entry Point

### Ứng với một service có thể định nghĩa nhiều Router được không?
- Có

### docker-compose.yml, dịch vụ auth hãy giải thích 

- "traefik.http.middlewares.auth.forwardauth.address=http://auth:3000/auth"
 + Định nghĩa địa chỉ của forward-auth middleware, nếu trả về status 200 thì đc call đến service, ko sẽ trả về error.

- "traefik.http.middlewares.auth.forwardauth.trustForwardHeader=true"
 + Chấp nhận các header định dạng: X-Forwarded-*

- "traefik.http.middlewares.auth.forwardauth.authResponseHeaders=X-Forwarded-User"
 + Header X-Forwarded-User được copy từ auth vào request

- "traefik.http.services.auth.loadbalancer.server.port=3000"
 + Traefik và auth service giao tiếp qua port 3000
## Project Auth

### main.go

- config.ReadConfig() nếu truyền vào tham số thì tác dụng của tham số này là gì?
 + Truyền vào đường dẫn thư mục chứa file log

- defer logFile.Close() để làm gì?
 + Đóng file log khi chuẩn bị kết thúc hàm.

- tại sao không để defer trong hàm InitSession mà phải để ở main.go
 + Để có thể ghi nhiều log nhất có thể rồi mới close
- Tác dụng của crs := cors.New(cors.Options{
 + Cho phép client gọi đến resouce server từ các nguồn khác nhau
- Khác biệt giữa app.UseRouter và app.Use là gì?
 + app.UseRouter chạy trước UseGlobal
 + app.Use chạy sau UseGlobal

- app.Any khác gì với app.Get và app.Post?
 + app.Any: Tất cả các request method đều hợp lệ với path định nghĩa để đc handle
 + app.Get và app.Post: chỉ method Get và Post mới hợp lệ với path định nghĩa để đc handle
### controller

- Tại sao trong cấu hình traefik luôn trỏ đến /authenticate -> controller.Authenticate()
+ Vì được định danh địa chỉ cho middleware auth là: traefik.http.middlewares.auth.forwardauth.address=http://auth:3000/auth

### config.go

- Khi nào thì chúng ta sẽ đọc config.dev.json và khi nào chúng ta sẽ đọc config.product.json?
+ config.dev.json sửa dụng ở chế độ debug (dựa vào command run)
+ config.product.json sửa dụng ở chế độ production

- Ý nghĩa của lệnh panic(err) là gì?
+ Khi gặp lỗi này thì buộc dừng chương trình 

### model/user.go

- Tại sao chúng ta lại tách định nghĩa user ra khỏi controller và repo?
+ user dùng chung cho cả controller và repo

### session.go

- để đọc một map[string]interface{} vào một struct chúng ta dùng thư viện gì?
+"github.com/mitchellh/mapstructure"

### Session

- Khác biệt giữa Sess.Clear() vs Sess.Delete() và Sess.DeleteFlash() là gì?
+ Sess.Clear() : xoá tất cả entries thuộc session
+ Sess.Delete() : xoá 1 entries bởi key thuộc session
+ Sess.DeleteFlash() : xoá 1 flash msg bởi key thuộc session

## Docker

### Lệnh docker khác gì lệnh docker-compose?
- lệnh docker để build image và run container
- lệnh docker-compose để build và run các container được khai báo là các service trong file docker-compose.yml.

### Portainer để làm gì?
- giao diện dashboard quản lý các thông tin của docker chạy trên máy host

### -p 80:8080 có ý nghĩa gì? 80 là cổng nào và 8080 là cổng nào?
- 80 là cổng máy host, 8080 là cổng của container

### docker build cần có những tham số gì?
- docker build [OPTIONS] PATH | URL | -
- Build docker image từ dockerfile và context nào đó (PATH, URL, GIT Repo)
### Khác biệt giữa docker image và docker container là gì?
- docker image: chỉ là ảnh(mẫu) file chứa source, lib, config,... để chạy ứng dụng.
- docker container: chạy docker image, khi đó các app sẽ được chạy, chia sẻ tài nguyên với máy host, có nhiều container chạy trên máy host.
### Dockerfile khác gì với docker-compose.yml
- Dockerfile chỉ ra cách build image
- docker-compose.yml chỉ ra cách run nhiều container

### Muốn service A trong docker-compose được khởi động sau khi service B đã chạy thì cần thêm cấu hình gì?
- depends_on

### Lệnh tạo một vùng lưu trữ trên host để docker container ghi vào là gì?
- -v app:/app
(volume)

### Ý nghĩa của lệnh này /var/run/docker.sock:/var/run/docker.sock
- Lắng nghe các event từ docker daemon
### docker-compose có lệnh để build lại những docker image cần dùng cho service, lệnh đó là lệnh nào?
  + docker-compose build
