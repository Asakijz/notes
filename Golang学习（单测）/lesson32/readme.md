## 通过文件方式输出 Profile

* 灵活性高，适用于特定代码段的分析
* 通过手动调用 runtime/pprof 的 API
* go tool pprof [binary][binary.prof]


## 通过 HTTP 方式输出 Profile

* 简单，适合于持续性运行的应用
* 在应用程序中导入 import _ "net/http/pprof"，并启动 http server 即可
* http://<host>:<port>/debug/pprof/
* go tool pprof _http://<host>:<port>/debug/pprof/profile?seconds=10（默认值为30秒）
* go-torch -seconds 10 http://<host>:<port>/debug/pprof/profile