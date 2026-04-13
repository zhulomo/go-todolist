# go-todolist

基于 Go + Gin 构建的RESTful API，实现了用户认证和任务管理。

## 技术栈
- **框架**: Gin
- **ORM**: GORM
- **数据库**: SQLite
- **认证**: JWT
- **文档**: Swagger

## 功能

- 用户注册 / 登录
- 任务的增删改查
- 基于角色的权限控制（admin/user）
- 统一错误处理中间件

## 环境要求

- Go 1.21+

### 运行

```bash
# 克隆项目
git clone https://github.com/zhulomo/go-todolist.git
cd go-todolist

# 安装依赖
go mod tidy

# 运行
go run main.go
```

### Swagger 文档
http://localhost:8080/swagger/index.html

## API 概览

| 方法 | 路径 | 说明 | 是否需要认证 |
|------|------|------|------------|
| POST | /register | 用户注册 | 否 |
| POST | /login | 用户登录 | 否 |
| POST | /tasks/create | 创建任务 | 是 |
| GET | /tasks/:id | 获取任务 | 是 |
| PUT | /tasks/:id | 更新任务 | 是 |
| DELETE | /tasks/:id | 删除任务 | 是 |
| GET | /tasks/get | 获取所有任务 | 是（admin） |
| GET | /admin/users | 获取所有用户 | 是（admin） |

## 项目结构

```
├── handler/      # 请求处理层
├── service/      # 业务逻辑层
├── repository/   # 数据访问层
├── middleware/   # 中间件
├── dto/          # 数据传输对象
├── response/     # 统一响应格式
├── utils/        # 工具函数
└── router/       # 路由注册
```