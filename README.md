# GoLibSourceAnalyse
go lib source analyse
提供代码分析仓库模板

#### 新模块创建
```shell
mkdir temp
cd temp
go mod init
# add dep
```

通过 go.work 方式来管理各个子模块。

### 目录
1. gorm ORM基于Golang 实现模型。

2. Leetcode 代码算法测试快。

3. lib 基础类库，语言规范。

4. rate 频率限制先关类库。

5. web 基于gin网页开发。

6. zap 日志库。

7. http2 项目网关方式。

8. project 简单开发个人项目。

9. utils 命令行小工具，简化日常开发过程。

10. tools 相关工具库分析使用

11. gopatch 代码格式化工具 [git](https://github.com/uber-go/gopatch)

12. goday 通过 import分析当前项目模块依赖