<p align="center">
<img src=https://raw.githubusercontent.com/davidliyutong/ingress-authproxy/main/docs/img/logo.svg width="300">
</p>

<br/>
<h1 align="center">
Ingress AuthProxy
</h1>
<p align="center">
<i>
An authorization server compatible with kubernetes ingress controllers.
</i>
<p align="center">
<img src="https://img.shields.io/github/license/davidliyutong/ingress-authproxy.svg"/>
<img src="https://img.shields.io/github/repo-size/davidliyutong/ingress-authproxy.svg"/>
<img src="https://img.shields.io/github/last-commit/davidliyutong/ingress-authproxy.svg"/>
<img src="https://img.shields.io/docker/stars/davidliyutong/ingress-authproxy-backend.svg"/>
  <a href="https://github.com/davidliyutong/ingress-authproxy/releases">
    <img src="https://img.shields.io/github/release/davidliyutong/ingress-authproxy.svg" alt="GitHub release">
  </a>
</p>

Ingress AuthProxy provides authorization service for Kubernetes Ingress Controllers. It supports rule based access policies and has an easy-to-user management WebUI.

## Features

-   ✅ Control access to K8S ingress resources
-   🔑 Access policies based on Ladon
-   🔌 A manaagement frontend UI that is powered by Vuetify
-   😀 Self-serviced password reset

## Architecture

The Ingress Authproxy contains a frontend and a backend. The frontend is an nginx server that serves at the following routes:

| Route            | Description                                  |
| ---------------- | -------------------------------------------- |
| `/`              | Login Page                                   |
| `/passwordreset` | Self-serviced password reset page            |
| `/admin`         | Admin page, only accessible to system admins |
| `/v1`            | Proxied to backend server                    |

The frontend webserver is configured so that `/v1/` path is proxied to `$AUTHPROXY_BACKEND_URL/v1/`, where `$AUTHPROXY_BACKEND_URL` is an environment variable passed when launching the frontend container.

The backend is an apiserver capable of performing CRUD operation on users and policies under path `/v1/admin/users/` and `/v1/admin/policies`

For details, see [docs/architecture.md](./docs/architecture.md)

## Get Started

### Prerequisites

To build this application locally, you have to install Golang, NodeJS on your system

To run in demo mode or build images, you have to install `docker` and `docker-compose` on your system

To deploy in K8S clusters, you have to own an K8S cluster and configure `kubectl` to use this cluster

### Deploying the AuthProxy app

To get started, clone this repository and run the application with `docker-compose`

```shell
git clone https://github.com/davidliyutong/ingress-authproxy
cd ingress-authproxy/manifests/authproxy
docker-compose up -d
```

This will launch 3 containers: a mysql database, an auth backend and an auth frontend. You can access the frontend from [http://localhost:8080](http://localhost:8080). A default admin user will be created with password `admin123456`

For deployment in kubernetes clusters, use the `manifests/k8s/deployment.yaml`:

```yaml
kubectl apply -f manifests/k8s/deployment.yaml -n <namespace>
```

This will create the following resources in your cluster's namespace

-   an `ingress-authproxy.<namespace>.svc.cluster.local` service with port `80`
-   a PersistenceVolumeClaim with capacity of 1GiB using default StorageClass
-   a Deployment of 3 containers in one Pod: mysql database, backend, frontend

To protect target K8S ingress resource, add this snippet to its annotations.

```yaml
annotations:
    nginx.ingress.kubernetes.io/auth-response-headers: X-Forwarded-User
    nginx.ingress.kubernetes.io/auth-url: "http://ingress-authproxy.<namespace>.svc.cluster.local/ingress-auth/<resource>"
```

Hint: Replace `<namespace>` and `<resources>` with deployed namespace of ingress authproxy and the name of resource.

The ingress resouce is now represented as `resources:ingress-auth:<resource>` in Ladon policy model. For example, you can configure this policy to allow admin to access all ingress resources

```json
{
    "description": "Allow admin access",
    "subjects": ["users:admin"],
    "actions": ["get"],
    "effect": "deny",
    "resources": ["resources:ingress-auth:<.*>"]
}
```

For details, see [docs/deployment.md](./docs/deployment.md)

## Developping

The Ingress AuthProxy contains an Vue frontend. Here are some important directories:

| Path                 | Description                                                      |
| -------------------- | ---------------------------------------------------------------- |
| `build/docker`       | Dockfiles for frontend and backend                               |
| `cmd/authproxy`      | Backend entry point                                              |
| `docs`               | Documents                                                        |
| `frontend`           | Root directory of frontend project                               |
| `internal`           | Internally used code                                             |
| `internal/apiserver` | Apiserver for user and policy resources                          |
| `internal/authproxy` | Backend service code                                             |
| `internal/config`    | Code to prepare database and read configuration from environment |
| `internal/utils`     | Some utilities                                                   |
| `manifests`          | Files for building docker images or deploying                    |
| `pkg`                | Externally used code                                             |
| `scripts`            | Build scripts and etc.                                           |

This project use `make`. Available make commands are:

| Command                          | Description                                |
| -------------------------------- | ------------------------------------------ |
| `build`/`go.build`   | Build go binaries                          |
| `image`,`image.push` | Build and push docker images to Docker Hub |
| `demo`,`demo.stop`   | Start and stop demo with `docker-compose`  |

For details, see [docs/developement.md](./docs/developement.md)

## Documentation

Currently, the documentation of Ingress Auth Proxy is under development.

## License

This project is open-sourced software licensed under the MIT license.
