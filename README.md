#### nginx日志回放压测
 - 解析日志进行回放压测，模拟后端服务器慢等各种异常情况


#### 方案说明
 - 客户端解析access.log构建请求的host,port,url,body
 - 把后端响应时间，后端响应状态码，后端响应大小放入header头中
 - 后端服务器获取相应的header，进行模拟响应body大小，响应状态码，响应时间


#### 使用方式
  - 拷贝需要测试的access.log的日志到logs文件夹里面
  - 搭建需要测试的nginx服务器，并且配置upstream 指向后端服务器断端口
  - 启动后端服务器实例 server/backserver/main.go
  - 进行压测 bin/wrk -c30 -t1  -s conf/nginx_log.lua http://localhost:8095



### Command Line Options
```
-c, --connections: total number of HTTP connections to keep open with
                   each thread handling N = connections/threads

-d, --duration:    duration of the test, e.g. 2s, 2m, 2h

-t, --threads:     total number of threads to use

-s, --script:      LuaJIT script, see SCRIPTING

-H, --header:      HTTP header to add to request, e.g. "User-Agent: wrk"

    --latency:     print detailed latency statistics

    --timeout:     record a timeout if a response is not received within
                   this amount of time.
```
