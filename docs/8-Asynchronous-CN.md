[← 并发](7-Concurrent-CN.md) | 异步调用[(English)](8-Asynchronous-EN.md) | [首页 →](../README_zh.md)
***
## 异步调用

### 发起异步调用
阿里云Go SDK支持两种方式的异步调用：

1. 使用channel作为返回值
    ```go
    responseChannel, errChannel := client.FooWithChan(request)

    // this will block
    response := <-responseChannel
    err = <-errChannel
    ```

2. 使用 callback 控制回调

    ```go
    blocker := client.FooWithCallback(request, func(response *FooResponse, err error) {
        // handle the response and err
    })

    // blocker 为(chan int)，用于控制同步，返回1为成功，0为失败
    // 在<-blocker返回失败时，err依然会被传入的callback处理
    result := <-blocker
    ```

***
[← 并发](7-Concurrent-CN.md) | 异步调用[(English)](8-Asynchronous-EN.md) | [首页 →](../README_zh.md)
