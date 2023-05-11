# 工具项目开发支持

### fx 做依赖管理
1. logger ```DEBUG``` 输出应用级别日志

### 依赖模块
1. gorm mysql 支持 db 数据查询
2. corba 支持命令行方式执行 
3. stringy 支持word case 转换

### 模块解析
1. cmd 命令工具模块支持不同功能
2. conf 工具配置数据包
3. dest 代码生成, 工具运行时刻配置。
4. util 工具模块
5. common 通用模块层

- 逻辑层次依赖 main -> cmd -> common -> utils
- conf 全局配置文件 conf 运行初始化
- dest 运行时刻配置文件

### 命令工具
1. model 数据表转换Dao 查询模型
2. word 驼峰大小写转换
3. convert 自动化文件转换工具