# API application **``TTT-Online``**
[![Linter](https://github.com/Drozd0f/ttto-go/actions/workflows/linter.yml/badge.svg)](https://github.com/Drozd0f/ttto-go/actions/workflows/linter.yml)

## Practice project

> **Note**
> Dependencies

* Docker
* docker-compose
* Make

### Quick start

> **Note**
> First of all configure ops/.env file with next variables

```
TTTO_DBURI=DB_URI
TTTO_ADDR=APPLICATION_PORT
TTTO_DEBUG=TRUE_OR_FALSE
TTTO_SECRET=SECRET_FOR_JWT_AND_HASHING_PASSWORD
```

---

**Run**
```shell
make
```

> **Note**
> Application will be hosted on http://localhost:4444

---

Remove containers
```shell
make rm
```

---

Generate SQLC code
```shell
make generate
```
