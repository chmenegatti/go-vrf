<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" alt="project-logo">
</p>
<p align="center">
    <h1 align="center">GO-VRF</h1>
</p>
<p align="center">
    <em>API HTTP para provisionamento de VRFs e redes no NSX-T.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/chmenegatti/go-vrf?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/chmenegatti/go-vrf?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/chmenegatti/go-vrf?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/go-mod/go-version/chmenegatti/go-vrf?style=default&color=0080ff" alt="go-version">
</p>

---

## Sumário

- [Visão geral](#visão-geral)
- [Funcionalidades](#funcionalidades)
- [Arquitetura](#arquitetura)
- [Estrutura do repositório](#estrutura-do-repositório)
- [Pré-requisitos](#pré-requisitos)
- [Configuração](#configuração)
- [Executando](#executando)
- [Endpoints da API](#endpoints-da-api)
- [Fluxo de uso](#fluxo-de-uso)
- [Testes](#testes)
- [Convenções](#convenções)
- [Contribuindo](#contribuindo)
- [Licença](#licença)

---

## Visão geral

`go-vrf` é um serviço HTTP, escrito em Go com o framework [Fiber](https://gofiber.io/),
que automatiza a configuração de **VRFs** (Virtual Routing and Forwarding) e redes
em ambientes **VMware NSX-T**.

O serviço consulta a API do NSX-T para resolver identificadores de recursos
(edge clusters, transport zones, tier-0/tier-1 gateways, security policies,
segments, groups, profiles e logical switches) a partir de nomes de exibição,
monta os registros de domínio correspondentes, persiste o resultado em arquivos
JSON e devolve os comandos SQL de inserção prontos para o sistema downstream.

Suporta múltiplos ambientes NSX-T (ex.: `TESP2`, `TESP3`, `TESP5`, `TESP6`),
selecionáveis por requisição através do campo `Edge`.

---

## Funcionalidades

- **Resolução de recursos NSX-T** por nome de exibição via API REST.
- **Três endpoints** que cobrem o fluxo de provisionamento: chave etcd → VRF de
  organização → redes de produtos.
- **Multi-ambiente**: credenciais e base path por prefixo de ambiente no `.env`.
- **TLS configurável**: verificação de certificado habilitada por padrão, com
  opção de CA bundle customizado ou bypass explícito.
- **Resiliência**: cliente HTTP com retry em erros de rede e respostas `5xx`,
  backoff exponencial e cancelamento via `context.Context`.
- **Geração de SQL** de inserção (`organizations`, `networks`) com escape seguro
  contra SQL injection.
- **Servidor robusto**: middleware de `recover`, `logger` e `request-id`, além de
  graceful shutdown.

---

## Arquitetura

Requisição HTTP → **controller** (parse + resposta) → **service** (regra de
negócio) → **nsxt** (cliente da API NSX-T) → NSX-T.

| Pacote            | Responsabilidade                                                        |
| ----------------- | ----------------------------------------------------------------------- |
| `main`            | Sobe o Fiber, registra middleware e trata o graceful shutdown.          |
| `src/routes`      | Mapeia os endpoints HTTP para os handlers.                              |
| `src/controller`  | Faz parsing do payload, orquestra os services e monta a resposta/SQL.   |
| `src/service`     | Regra de negócio: valida entrada e correlaciona dados do NSX-T.         |
| `src/nsxt`        | Cliente HTTP autenticado + tipos de resposta da API NSX-T.              |
| `src/model`       | Structs de domínio persistidas (EdgeCluster, Organizations, Networks).  |
| `src/objects`     | Structs dos payloads de requisição.                                     |
| `src/utils`       | UUIDs, leitura/escrita de JSON (com validação de nome) e escape de SQL. |
| `src/configs`     | Carregamento de variáveis de ambiente (`.env` carregado uma vez).       |

---

## Estrutura do repositório

```sh
go-vrf/
├── main.go              # bootstrap do servidor Fiber
├── Makefile             # alvos de build/test/lint
├── go.mod / go.sum
├── .env.example         # modelo de configuração
├── .golangci.yml        # configuração do linter
├── .github/workflows/   # CI (build, test, lint)
├── docs/insomnia/       # coleção Insomnia da API
└── src/
    ├── configs/         # variáveis de ambiente
    ├── controller/      # handlers HTTP
    ├── model/           # structs de domínio
    ├── nsxt/            # cliente da API NSX-T
    ├── objects/         # payloads de requisição
    ├── routes/          # registro de rotas
    ├── service/         # regra de negócio
    └── utils/           # utilitários (UUID, JSON, SQL)
```

---

## Pré-requisitos

- **Go** 1.22 ou superior.
- Acesso de rede e credenciais a um ou mais appliances **NSX-T**.
- (Opcional) [`golangci-lint`](https://golangci-lint.run/) v2.10.x para rodar o lint localmente.

---

## Configuração

A configuração é feita via variáveis de ambiente. Copie o modelo e preencha:

```console
$ cp .env.example .env
```

Cada ambiente NSX-T usa um **prefixo** (ex.: `TESP5`) com três variáveis:

```dotenv
TESP5_BASEPATH=https://nsx.exemplo.local
TESP5_USERNAME=apiuser
TESP5_PASSWORD=troque-me
```

Replique o bloco para cada ambiente que precisar. O prefixo é informado em cada
requisição pelo campo `Edge` (ex.: `"Edge": "TESP5"`).

### TLS

| Variável               | Padrão  | Descrição                                                                 |
| ---------------------- | ------- | ------------------------------------------------------------------------- |
| `NSXT_SKIP_TLS_VERIFY` | `false` | Se `true`, ignora a verificação de certificado (**inseguro**, só dev).    |
| `NSXT_CA_BUNDLE`       | (vazio) | Caminho para um PEM com a CA do NSX-T, anexado às raízes do sistema.      |

Para NSX-T com certificado autoassinado, o recomendado é instalar a CA e apontar
`NSXT_CA_BUNDLE`; `NSXT_SKIP_TLS_VERIFY=true` deve ser apenas um paliativo.

> ⚠️ O arquivo `.env` é ignorado pelo Git e **nunca** deve ser commitado.

---

## Executando

O `Makefile` cobre as tarefas comuns:

```console
$ make run               # sobe o servidor (porta :4000)
$ make build             # compila o binário ./go-vrf
$ make test              # testes unitários (exclui integração)
$ make test-integration  # testes contra um NSX-T real (requer credenciais)
$ make lint              # golangci-lint
$ make fmt               # formata o código
$ make check             # suíte completa: fmt + vet + lint + test
$ make help              # lista todos os alvos
```

O servidor escuta em `:4000` e encerra de forma graciosa ao receber
`SIGINT`/`SIGTERM`.

---

## Endpoints da API

Todos os endpoints são `POST` e recebem/retornam JSON. Base URL: `http://localhost:4000`.

### 1. `POST /generate-etcd-key`

Resolve o edge cluster e a transport zone pelo nome, monta o registro de
`EdgeCluster` e o persiste em `<VrfName>.json`.

**Payload** (`objects.EdgeClusterEtcd`):

```json
{
  "Edge": "TESP5",
  "VrfName": "T0-Cluster_4",
  "NsxtEdgeClusterName": "edge-cluster-01",
  "TransportZoneName": "tz-overlay",
  "VirtualFirewall": "fw-virtual-01",
  "FirewallExternalAddress": "10.0.0.1",
  "MaxOrganization": 100,
  "Enable": true,
  "RubrikDatabaseCluster": "rubrik-01"
}
```

```console
$ curl -X POST http://localhost:4000/generate-etcd-key \
    -H 'Content-Type: application/json' \
    -d @payload-etcd.json
```

**Resposta** (`200`): objeto chaveado pelo `VrfName` com o registro gerado, mais o
nome do arquivo salvo:

```json
{
  "T0-Cluster_4": { "nsxt_tier0_id": "T0-Cluster_4", "index": 11, "...": "..." },
  "filename": "T0-Cluster_4.json"
}
```

### 2. `POST /create-organization-vrf`

Lê `<VrfName>.json`, resolve o Tier-1 gateway e a security policy pelo nome
(`NameTier1`), monta a organização e a persiste em `<NameTier1>.json`.

**Payload** (`objects.OrganizationVRF`):

```json
{
  "Edge": "TESP5",
  "VrfName": "T0-Cluster_4",
  "NameTier1": "DB-Shared_10"
}
```

```console
$ curl -X POST http://localhost:4000/create-organization-vrf \
    -H 'Content-Type: application/json' \
    -d @payload-org.json
```

**Resposta** (`200`): o registro da organização e o `INSERT` SQL correspondente:

```json
{
  "data": { "id": "…", "name": "DB-Shared_10", "status": "COMPLETED", "...": "..." },
  "sql": "INSERT INTO organizations (...) VALUES (...);"
}
```

### 3. `POST /create-networks-vrf`

Lê `<NameTier1>.json`, resolve segments, groups, profiles e logical switches para
cada produto informado, monta as redes e persiste em `<NameTier1>_Networks.json`.

**Payload** (`objects.NetworksProductsVRF`):

```json
{
  "Edge": "TESP5",
  "NameTier1": "DB-Shared_10",
  "Products": ["produto-a", "produto-b"]
}
```

```console
$ curl -X POST http://localhost:4000/create-networks-vrf \
    -H 'Content-Type: application/json' \
    -d @payload-net.json
```

**Resposta** (`200`): a lista de redes e os `INSERT` SQL (um por rede):

```json
{
  "data": [ { "id": "…", "name": "produto-a", "status": "COMPLETED", "...": "..." } ],
  "sql": [ "INSERT INTO networks (...) VALUES (...);" ]
}
```

> Em caso de erro de validação ou de comunicação com o NSX-T, a resposta usa
> status `4xx`/`5xx` com um objeto `{ "message": "..." }`.

---

## Fluxo de uso

Os endpoints são encadeados — cada um consome o JSON gerado pelo anterior:

```
generate-etcd-key      →  escreve  <VrfName>.json
        │
create-organization-vrf →  lê <VrfName>.json   →  escreve <NameTier1>.json
        │
create-networks-vrf     →  lê <NameTier1>.json →  escreve <NameTier1>_Networks.json
```

Os arquivos JSON gerados são artefatos de runtime (ignorados pelo Git).

---

## Testes

```console
$ make test               # unitários — sem chamadas externas
$ make test-integration   # integração — exige NSX-T real + credenciais no .env
```

Os testes de integração (`src/nsxt/nsxt_api_test.go`) ficam atrás da build tag
`integration` e batem em um appliance NSX-T real (`TESP5`), por isso são
excluídos da execução padrão.

---

## Convenções

- **Código e comentários em inglês.** Documentação voltada ao usuário (como este
  README) pode ser em português.
- **Antes de abrir um PR**, rode a suíte completa e garanta que está verde:
  ```console
  $ make check   # fmt + vet + lint + test
  ```
- A CI roda `build`, `test` e `lint` (golangci-lint) em cada push e PR.

---

## Contribuindo

1. **Fork** do repositório.
2. **Clone** local:
   ```sh
   git clone https://github.com/chmenegatti/go-vrf.git
   ```
3. **Branch** descritiva:
   ```sh
   git checkout -b fix/minha-correcao
   ```
4. **Implemente e teste** — rode `make check` antes de commitar.
5. **Commit** com mensagem clara do "porquê" da mudança.
6. **Push** e abra um **Pull Request** contra `main`, descrevendo a motivação.

Bugs e sugestões: [issues do projeto](https://github.com/chmenegatti/go-vrf/issues).

---

## Licença

Defina uma licença para o projeto adicionando um arquivo `LICENSE` na raiz.
Veja [choosealicense.com](https://choosealicense.com/) para ajudar na escolha.

[**Voltar ao topo**](#go-vrf)
