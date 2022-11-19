核心层代码组织结构
cores: 低代码框架核心

    g: 公共变量及对象
    initialize: 初始化
    routers: 路由层: 框架路由和业务路由
    
    api: 对外接口: 参数校验--调用service层
    services: 服务层: 封装业务逻辑行为,调用model层数据
    models: 数据层: 数据实体对象定义及操作
    
    types: 框架业务对象
    config.go: 加载配置文件
    server.go: 服务启动
    zap.go: 日志


