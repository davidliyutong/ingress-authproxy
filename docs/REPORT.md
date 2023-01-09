# Final Project of Cloud Native Apps

<p align="center">
<img src=https://raw.githubusercontent.com/davidliyutong/ingress-authproxy/main/docs/img/logo.svg width="300">
</p>

在云原生应用时代，K8S Ingress 作为集群入口，承载了大部分的 Web 应用流量。许多场景下，这些流量需要被合理的认证制度保护起来，以控制特定用户对特定资源的访问。一个方案是在应用端实现权限认证，但是这意味着需要针对每个应用开发一套认证系统，造成了极大的资源浪费。不同系统的用户认证也很难共享。

所幸，K8S 支持对 Ingress 资源施加外部的认证，这是通过 Ingress 资源的特殊标记（Annotation）进行的。本次期末作业本人选择了开发 Ingress AuthProxy 应用为 Kubernetes Ingress Controllers 提供授权服务。它支持基于规则的访问策略，并具有易于用户管理的 WebUI。通过 Ingress AuthProxy，您可以用同一套用户权限系统保护集群所有的 Ingress 资源。

项目地址:[https://github.com/davidliyutong/ingress-authproxy](https://github.com/davidliyutong/ingress-authproxy)

该项目有以下特性：

-   ✅ 控制对 K8S Ingress 资源的访问，精确到特定用户对特定资源的策略
-   🔑 基于 Ladon 的访问策略，支持基于通配符的规则。规则保存在持久化存储中并在内存中国呢还从
-   🔌 由 Vuetify 提供支持使用 Material Design 的管理前端 UI，精美而人性化
-   😀 自助密码重置，用户在创建后可以自助修改密码
-   📦 完全的容器化，提供 docker-compose 和 k8s 配置，部署 AuthProxy 仅需数分钟
-   ♻️ 基于 GitHub Action 的自动镜像编译和自动二进制编译，并推送到 Docker Hub

## Video Demo

访问 [Ingress AuthProxy 演示](https://www.bilibili.com/video/BV1BA411Z7D3/?share_source=copy_web&vd_source=5b66e799bd8260c54040de6f4f52b978) 查看视频演示

## 项目架构

Ingress AuthProxy 是典型的前后端分离项目。包含前端和后端。数据存储方面使用 MySQL 数据库。

Ingress AuthProxy 前端前端是一个 nginx 服务器，服务于以下路由：

| 路由             | 描述                 |
| ---------------- | -------------------- |
| `/`              | 登录页面             |
| `/passwordreset` | 用户自助密码重置     |
| `/admin`         | 仅系统管理员可以登录 |
| `/v1/*`          | 转发到后端服务       |

在部署中，启动前端时需要指定`AUTHPROXY_BACKEND_URL`，以正确配置前端的 Nginx 服务，具体可以参考`manifests/authproxy/docker-compose.yml`的示例。

后端是一个 GO 编写的 apiserver，主要参考了[scaffold](https://github.com/rebirthmonkey/go/tree/master/scaffold/apiserver)，在`/v1/admin/users/`和`/v1/admin/policies`上提供 CRUD 服务，在`/v1/user`上提供用户自助修改密码服务（只支持用户对自己的 GET/UPDATE），在`/v1/admin/server`上提供服务器信息查询功能，在`/v1/jwt/login`，`/v1/jwt/refresh`上提供 JWT Token 签发功能

具体的业务在`/v1/ingress-auth/:name`开展，该路由对应了 Ladon 模型中的`resources:ingress-auth:<name>`。该路径实现了一个 BasicAuth 认证。认证成功后，该路由还会根据数据库中存储的 Policy 对当前尝试登录的用户做权限鉴定，并根据权限判断用户是否可以访问当前资源。

由于该过程涉及到对 Policy 的检索，以及可能会带来高并发，因此对于策略的存储采取以下模式

-   当`/v1/admin/policies`出现 UPDATE 或者 DELETE 操作时，后端服务器会触发一次更新操作，从 MySQL 数据库同步新的 Policy
-   当`/v1/server/sync`收到 POST 请求时，将强制同步 Policy
-   当超时发生后，将自动同步 Policy，默认为 10 分钟
-   使用 RWMutex 保护 Policy 的一致性

后端初始化后会尝试连接 MySQL 数据库，当连接失败后将会退出，此时一个守护脚本将会自动重启后端。后端连接 MySQL 成功后，将检测特定的数据库是否存在 user/policy/secrets 数据表，若不存在则会执行默认的初始化 SQL 脚本初始化数据库并根据启动时传入的`AUTHPROXY_INIT_USERNAME`和`AUTHPROXY_INIT_PASSWORD`创建默认用户

## 部署应用

该应用支持本地（虚拟机）、Docker、K8S 集群等多种功能部署

首先，从 GitHub 下载源码

```shell
git clone https://github.com/davidliyutong/ingress-authproxy
cd ingress-authproxy/manifests/authproxy
```

> 您也可以访问 GitHub Release 页面，下载预编译的二进制文件（由 GitHub Actions 自动编译）

### 本地

要在本地构建此应用程序并以独立进程或虚拟机的形式，您必须在目标系统上安装 Golang、NodeJS、Nginx

首先编译后端

```shell
make go.build
```

编译前端，前端依赖 yarn 编译

```shell
cd frontend
yarn install && yarn build
```

根据`manifests/authproxy-frontend/nginx/nginx.conf`中的配置设置 Nginx，并拷贝打包后的前端`frontend/dist`到指定路径，然后使用正确的环境变量启动应用

### Docker

项目提供了开箱即用的`docker-compose.yml`供单机部署。您必须在系统上安装 docker 和 docker-compose

```shell
cd manifests/authproxy
docker-compose up -d
```

**警告 ⚠️**: 默认配置下，docker 容器中的数据将使用临时存储，并随着容器销毁而消失，请指定 MySQL 容器的数据卷以持久保存数据

这将启动 3 个容器：一个 mysql 数据库、一个身份验证后端和一个身份验证前端。您可以从 [http://localhost:8080](http://localhost:8080) 访问前端。将使用密码创建默认管理员用户 admin123456

> 将会从 Docker Hub 拉取 latest 镜像

### K8S 集群部署

要在 K8S 集群中部署，您必须拥有一个 K8S 集群并配置使用此集群的`kubectl`工具

修改并应用`manifests/k8s/deployment.yaml`。您可能需要定制存储卷容量和使用的类别。如果使用外部 MySQL 服务器，则需要删除配置文件自带的数据库配置并将相关环境变量修改

```shell
kubectl apply -f manifests/k8s/deployment.yaml -n <your_namespace>
```

这将在群集的执行中创建以下资源：

-   80端口的服务，集群内 DNS 名称为`ingress-authproxy.<your_namespace>.svc.cluster.local`
-   一个容量为 1GiB 的 PersistenceVolumeClaim，使用默认的 StorageClass
-   一个创建 3 容器 Pod 的 Deployment 资源：mysql 数据库、后端、前端

## 保护集群资源

若要保护目标 K8S Ingress 资源，请将此代码段添加到 Ingress 注释中。

```yaml
annotations:
    nginx.ingress.kubernetes.io/auth-response-headers: X-Forwarded-User
    nginx.ingress.kubernetes.io/auth-url: "http://ingress-authproxy.<your_namespace>.svc.cluster.local/v1/ingress-auth/<resource>"
```

您需要将<your_namespace>替换 AuthProxy 所在的命名空间，`<resource>`为当前资源的名称，可以自定义。该资源将会被抽象为`resources:ingress-auth/<resource>`，对该 Ingress 资源对访问操作被定义为`get`。只有对该资源有`get`权限的用户才能访问。以下是一些例子

1. 允许管理员访问所 Ingress 资源：

    ```json
    {
        "description": "Allow admin access",
        "subjects": ["users:admin"],
        "actions": ["get"],
        "effect": "deny",
        "resources": ["resources:ingress-auth:<.*>"]
    }
    ```

    2.允许 alice 访问被具有`nginx.ingress.kubernetes.io/auth-url: "http://ingress-authproxy.<your_namespace>.svc.cluster.local/v1/ingress-auth/test1"`的资源

    ```json
    {
        "description": "Allow alice access",
        "subjects": ["users:alice"],
        "actions": ["get"],
        "effect": "allow",
        "resources": ["resources:ingress-auth:test1"]
    }
    ```

    3.拒绝 bob 访问被具有`nginx.ingress.kubernetes.io/auth-url: "http://ingress-authproxy.<your_namespace>.svc.cluster.local/v1/ingress-auth/test1"`的资源

    ```json
    {
        "description": "Deny bob access",
        "subjects": ["users:bob"],
        "actions": ["get"],
        "effect": "deny",
        "resources": ["resources:ingress-auth:test1"]
    }
    ```

## 管理界面

如果是本地部署，可以访问localhost:8080端口，如果是集群部署，可以创建以下Ingress资源暴露该服务

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    kubernetes.io/ingress.class: nginx # replace with your ingress class
  name: ingress-authproxy
spec:
  rules:
  - host: authproxy.example.com # replace with your hostname
    http:
      paths:
      - backend:
          service:
            name: ingress-authproxy # replace with your service name
            port:
              number: 80
        path: /
        pathType: Prefix
  tls:
  - hosts:
    - authproxy.example.com # replace with your hostname
    secretName: authproxy.example.com-certificates # replace with your certificate
```

使用管理员登录后，可以打开界面了解服务器的响应时间和当前版本

![Demo](img/20230109030830.png)

- Users 面板可以管理用户
- Policy 面板可以管理策略
- Profile 可以查看自己的信息，并且修改密码
- Settings 可以操作服务器，例如同步和重启

![CreatAlice](img/20230109031124.png)

普通用户可以在登录界面选择重置自己的密码（`/passwordreset`），只有管理员才能打开这个面板登录并管理服务器。

默认情况下没有任何策略会被添加，因此任何人都不能访问任何资源。必须创建一些策略。

![CreatePolicy](img/20230109032346.png)


## 登录管理面板的时候会发生什么

用户访问`/`时，前端会渲染登录表单。用户输入表单后，前端JS脚本使用axios库向`/v1/jwt/login`发起POST请求

```json
{
    "accessKey": form.username,
    "secretKey": form.password,
}
```

`/v1/jwt/login`调用JWT中间件完成认证，并检查该用户的isAdmin字段，如果是管理员则发放Token，有效期两小时。

如果得到StatusOK的回复，前端就将token存在localStorage

```javascript
localStorage.setItem('token', response.data.token)
localStorage.setItem('username', form.username)
```

前端配置了一个axios拦截器来给所有向后端的请求加上Authorization Header。因此随后到请求都会使用给token核验

当用户登出时，将会清除localStorage中的token。

当用户访问`/passwordreset`并提交表单重置密码的时候，面板将会向`/v1/users`发起请求，这个请求将使用BasicAuth验证身份，并且执行`/v1/users/name`的update操作。每一个用户都只能修改自己的信息

## 访问被保护的Ingress将会发生什么？

假设URL为`test1.example.com`的Ingress资源拥有`nginx.ingress.kubernetes.io/auth-url: "http://ingress-authproxy.default.svc.cluster.local/v1/ingress-auth/test1"`标记

当用户alice第一次访问该URL时，IngressController就会向AuthProxy的`/v1/ingress-auth/test1`发起Get请求，此时请求中不包括任何Authorization Header，因此后端将会返回401 Unauthorized.

Ingress Controller向浏览器传递该回复，浏览器向用户弹出窗口要求登录。用户输入用户名和密码，浏览器会利用这些信息重新发起请求，这次请求中将会携带BaiscAuth的Authorization Header

Ingress Controller向AuthProxy的接口重新发起带Header的请求，AuthProxy首先根据Header完成鉴权，获取用户名称。然后AuthProxy调用Ladon的PolicyManager验证用户对特定资源的访问权限。

如果用户可以访问资源，则后端返回一个简短的JSON并设置`X-Forwarded-User` Header为用户名。浏览器将会缓存该Authorization Header。

```json
{"authenticated":true,"user":"alice"}
```

> 可以阅读[https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md#external-authentication](https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md#external-authentication) 了解更多例子

## 代码说明

Ingress AuthProxy 包含一个 Vue 前端。和 Golang 后端。以下是一些重要的目录：

| 路径                 | 描述                                      |
| -------------------- | ----------------------------------------- |
| `build/docker`       | 前端和后端的 Dockerfile                   |
| `cmd/authproxy`      | 后端程序入口                              |
| `docs`               | 文档                                      |
| `frontend`           | 前端                                      |
| `internal/apiserver` | 实现了 UserController 和 PolicyController |
| `internal/authproxy` | 初始化和后端启动代码                      |
| `internal/config`    | 后端配置读取和数据库初始化代码            |
| `internal/utils`     | 工具函数                                  |
| `manifests`          | 配置文件                                  |
| `pkg`                | 引用的一些外部包                          |
| `scripts`            | 编译用脚本                                |

项目使用的 Make 系统命令如下

| 命令                 | 描述                            |
| -------------------- | ------------------------------- |
| `build`/`go.build`   | 编译 Go 二进制                  |
| `image`,`image.push` | 编译推送镜像                    |
| `demo`,`demo.stop`   | 使用 `docker-compose` 运行 Demo |

## 待办事项

本项目里完善还有很长一段距离，以下是一些主要有待完成的功能

- **前端和后端数据合法性的检验：** 目前只有很少的前端校验，例如email，phone的合法性，username的合法性等等都没有校验
- **独立的auth-login登录界面：** 访问受保护的Ingress资源时，用户需要输入密码，如果该资源没有配备TLS，则用户名和密码会以明文传输，造成隐患。因此需要一个额外的登录界面用于跳转
- **OpenID Connect支持：** OpenID Connect是流行的现代的鉴权标准，Ingress AuthProxy正在努力支持它
- **Application Password：** 让用户可以自己创建应用的密码，实现多租户模型
- **用户群组：** 实现用户群组模型和基于群组的权限管理

## **许可证**

该项目是根据 MIT 许可证许可的开源软件。

## 致谢

本项目受到了[scaffold](https://github.com/rebirthmonkey/go/tree/master/scaffold/apiserver)项目的很大启发。用户模型，数据库操作，Controller设计都直接源于该项目。

-----
