此处使用 go-kit 步骤
[stingsvc官方教程](https://gokit.io/examples/stringsvc.html)
- 为服务建立接口
- 实现接口
- 建立请求和响应的结构体
- 理解 go-kit 中服务的每个方法都将通过适配器转换为端点，适配器接受 service 返回对应于其中方法的端点
- 使用 transport 将服务暴露给外界
