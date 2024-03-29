<p align="center">
  <img src="https://raw.githubusercontent.com/PKief/vscode-material-icon-theme/ec559a9f6bfd399b82bb44393651661b08aaf7ba/icons/folder-markdown-open.svg" width="100" alt="project-logo">
</p>
<p align="center">
    <h1 align="center">GO-VRF</h1>
</p>
<p align="center">
    <em>Empowering network innovation through intuitive HTTP solutions.</em>
</p>
<p align="center">
	<img src="https://img.shields.io/github/license/chmenegatti/go-vrf.git?style=default&logo=opensourceinitiative&logoColor=white&color=0080ff" alt="license">
	<img src="https://img.shields.io/github/last-commit/chmenegatti/go-vrf.git?style=default&logo=git&logoColor=white&color=0080ff" alt="last-commit">
	<img src="https://img.shields.io/github/languages/top/chmenegatti/go-vrf.git?style=default&color=0080ff" alt="repo-top-language">
	<img src="https://img.shields.io/github/languages/count/chmenegatti/go-vrf.git?style=default&color=0080ff" alt="repo-language-count">
<p>
<p align="center">
	<!-- default option, no dependency badges. -->
</p>

<br><!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary><br>

- [ Overview](#-overview)
- [ Features](#-features)
- [ Repository Structure](#-repository-structure)
- [ Modules](#-modules)
- [ Getting Started](#-getting-started)
  - [ Installation](#-installation)
  - [ Usage](#-usage)
  - [ Tests](#-tests)
- [ Project Roadmap](#-project-roadmap)
- [ Contributing](#-contributing)
- [ License](#-license)
- [ Acknowledgments](#-acknowledgments)
</details>
<hr>

##  Overview

Go-vrf is a robust open-source project that simplifies network service management through HTTP with Fiber v2 and godotenv. It allows users to handle NSX-T integration efficiently by creating organization VRFs, generating keys, and managing edge clusters. The project promotes scalability and security by leveraging Fiber-based API routes, TLS for secure communication, and UUID generation. With a focus on modularity and organized request handling, go-vrf offers a valuable solution for streamlining network configuration tasks.

---

##  Features

|    | Feature           | Description                                                                                                                                                                                                                                                                                                                                                              |
|----|-------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| ⚙️  | **Architecture**    | The project uses Fiber as the main web framework for handling HTTP services on port 4000. It integrates routes efficiently for network service handling. The codebase is organized into packages for NSX-T API operations, routes definition, services, utilities, and configurations.                                                                                             |
| 🔩 | **Code Quality**    | The codebase exhibits clean and well-structured code following Go best practices. It leverages relevant packages for error handling, API request implementations, and secure communication. The codebase emphasizes readability and maintainability, reducing technical debt and enhancing overall quality.                              |
| 📄 | **Documentation**   | The project includes detailed inline documentation within the source code files, aiding in understanding various functionalities. The repository also contains a go.mod file defining dependencies and their versions for consistency. However, there is room for improvement in providing external documentation for users and contributors.                   |
| 🔌 | **Integrations**    | Key integrations include Fiber v2 for web framework capabilities, godotenv for environment variable management, fasthttp for efficient HTTP handling, and UUID generation for unique identifiers. External dependencies like brotli and uniseg are utilized for specific functionalities within the project.                          |
| 🧩 | **Modularity**      | The codebase is structured into separate packages for distinct functionalities like NSX-T API operations, route definition, services, and utilities. This modular design promotes code reusability and maintainability, enabling easy extension and modification of different components without affecting the entire system.|
| 🧪 | **Testing**         | The project utilizes test files within the src directory to conduct unit tests for NSX-T API endpoints and various service functions. The tests include error handling scenarios, ensuring robustness in handling edge cases and maintaining the expected behavior of the codebase.                                            |
| ⚡️  | **Performance**     | The codebase demonstrates efficient resource usage and speed in handling network services via Fiber. API request implementations with retry logic, TLS for secure communication, and UUID generation contribute to improving performance and ensuring reliable connections within the project's architecture.                     |
| 🛡️ | **Security**        | Security measures are implemented through TLS for secure communication, ensuring data protection in API requests. The use of environment variables managed by godotenv enhances access control to sensitive configurations. However, further security enhancements like input validation can be considered for strengthening data security. |
| 📦 | **Dependencies**    | Key external dependencies include go-runewidth, go-colorable, and mod for additional functionalities. Dependencies like sys, bytebufferpool, and compress support various operations within the codebase. The project relies on go modules to manage dependencies effectively and ensure a stable ecosystem.                  |

---

##  Repository Structure

```sh
└── go-vrf/
    ├── go.mod
    ├── go.sum
    ├── main.go
    └── src
        ├── configs
        ├── controller
        ├── model
        ├── nsxt
        ├── objects
        ├── routes
        ├── service
        └── utils
```

---

##  Modules

<details closed><summary>.</summary>

| File                                                                     | Summary                                                                                                                                                                                                                               |
| ---                                                                      | ---                                                                                                                                                                                                                                   |
| [main.go](https://github.com/chmenegatti/go-vrf.git/blob/master/main.go) | Initializes a Fiber app, integrating routes to handle network services via HTTP on port 4000.                                                                                                                                         |
| [go.mod](https://github.com/chmenegatti/go-vrf.git/blob/master/go.mod)   | Defines dependencies for go-vrf project, leveraging Fiber v2 and godotenv. Notable indirect dependencies include brotli, UUID, and fasthttp. Positioned in the root directory, reinforcing project stability and ecosystem coherence. |
| [go.sum](https://github.com/chmenegatti/go-vrf.git/blob/master/go.sum)   | Defines dependencies and their versions for the project using modules. Ensures consistent versions of external packages for smooth functioning of the codebase.                                                                       |

</details>

<details closed><summary>src.nsxt</summary>

| File                                                                                                | Summary                                                                                                                                                                                             |
| ---                                                                                                 | ---                                                                                                                                                                                                 |
| [nsxt_api_test.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/nsxt/nsxt_api_test.go) | Tests NSX-T API endpoints for various resources. Retrieves and prints Edge Clusters, Tier-0 Gateways, Transport Zones, Logical Switches, Segments, and Groups using test cases with error handling. |
| [nsxt_api.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/nsxt/nsxt_api.go)           | GetEdgeCluster`, `GetTier0Gateways`, `GetSegments`, etc.                                                                                                                                            |
| [client.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/nsxt/client.go)               | Implements NSX-T API request with retry logic and TLS for secure communication, ensuring successful connections in the go-vrf repositorys nsxt package.                                             |

</details>

<details closed><summary>src.routes</summary>

| File                                                                                    | Summary                                                                                                                                                              |
| ---                                                                                     | ---                                                                                                                                                                  |
| [routes.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/routes/routes.go) | Defines routes for generating Etcd key and creating organization VRF in the parent repositorys Fiber-based API, facilitating modular and organized request handling. |

</details>

<details closed><summary>src.service</summary>

| File                                                                                                                   | Summary                                                                                                                                                                                                           |
| ---                                                                                                                    | ---                                                                                                                                                                                                               |
| [generateEtcdKey.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/service/generateEtcdKey.go)             | Generates Edge Cluster and Transport Zone IDs from NSX-T based on provided data to streamline configuration handling.                                                                                             |
| [createOrganizationVRF.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/service/createOrganizationVRF.go) | Creates organization Virtual Routing and Forwarding (VRF) based on provided name and edge. Retrieves Tier1 Gateway and Distributed Firewall Policy IDs from the NSX-T backend to associate with the organization. |

</details>

<details closed><summary>src.utils</summary>

| File                                                                                         | Summary                                                                                                                                                          |
| ---                                                                                          | ---                                                                                                                                                              |
| [utilities.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/utils/utilities.go) | Generates UUIDs and reads/writes JSON data to files for the repositorys EdgeCluster model. Handles file I/O operations efficiently with UUID generation support. |

</details>

<details closed><summary>src.objects</summary>

| File                                                                                       | Summary                                                                                                                                                                                                              |
| ---                                                                                        | ---                                                                                                                                                                                                                  |
| [objects.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/objects/objects.go) | Defines data structures for EdgeClusterEtcd and OrganizationVRF in the objects package. These structs model key attributes related to NSX-T edge clusters and organization VRFs within the repositorys architecture. |

</details>

<details closed><summary>src.configs</summary>

| File                                                                               | Summary                                                                                                                                                   |
| ---                                                                                | ---                                                                                                                                                       |
| [env.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/configs/env.go) | Retrieves environment variables from a.env file using godotenv. Critical for accessing sensitive configurations in the open-source projects architecture. |

</details>

<details closed><summary>src.controller</summary>

| File                                                                                                | Summary                                                                                                                                                                                                                                           |
| ---                                                                                                 | ---                                                                                                                                                                                                                                               |
| [controller.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/controller/controller.go) | Generates Etcd keys and creates Organization VRF, saving data to files. Handles request parsing, error responses, and UUID generation for cluster configuration. Enables efficient VRF and organization setup in the larger repository structure. |

</details>

<details closed><summary>src.model</summary>

| File                                                                                   | Summary                                                                                                                                                                                                                                         |
| ---                                                                                    | ---                                                                                                                                                                                                                                             |
| [models.go](https://github.com/chmenegatti/go-vrf.git/blob/master/src/model/models.go) | Defines data structures for Edge Cluster, Organizations, and Networks with JSON annotations for API responses. Captures key attributes for managing NSX-T integration and organizational networking within the parent repositorys architecture. |

</details>

---

##  Getting Started

**System Requirements:**

* **Go**: `version x.y.z`

###  Installation

<h4>From <code>source</code></h4>

> 1. Clone the go-vrf repository:
>
> ```console
> $ git clone https://github.com/chmenegatti/go-vrf.git
> ```
>
> 2. Change to the project directory:
> ```console
> $ cd go-vrf
> ```
>
> 3. Install the dependencies:
> ```console
> $ go build -o myapp
> ```

###  Usage

<h4>From <code>source</code></h4>

> Run go-vrf using the command below:
> ```console
> $ ./myapp
> ```

###  Tests

> Run the test suite using the command below:
> ```console
> $ go test
> ```

---

##  Project Roadmap

- [X] `► INSERT-TASK-1`
- [ ] `► INSERT-TASK-2`
- [ ] `► ...`

---

##  Contributing

Contributions are welcome! Here are several ways you can contribute:

- **[Report Issues](https://github.com/chmenegatti/go-vrf.git/issues)**: Submit bugs found or log feature requests for the `go-vrf` project.
- **[Submit Pull Requests](https://github.com/chmenegatti/go-vrf.git/blob/main/CONTRIBUTING.md)**: Review open PRs, and submit your own PRs.
- **[Join the Discussions](https://github.com/chmenegatti/go-vrf.git/discussions)**: Share your insights, provide feedback, or ask questions.

<details closed>
<summary>Contributing Guidelines</summary>

1. **Fork the Repository**: Start by forking the project repository to your github account.
2. **Clone Locally**: Clone the forked repository to your local machine using a git client.
   ```sh
   git clone https://github.com/chmenegatti/go-vrf.git
   ```
3. **Create a New Branch**: Always work on a new branch, giving it a descriptive name.
   ```sh
   git checkout -b new-feature-x
   ```
4. **Make Your Changes**: Develop and test your changes locally.
5. **Commit Your Changes**: Commit with a clear message describing your updates.
   ```sh
   git commit -m 'Implemented new feature x.'
   ```
6. **Push to github**: Push the changes to your forked repository.
   ```sh
   git push origin new-feature-x
   ```
7. **Submit a Pull Request**: Create a PR against the original project repository. Clearly describe the changes and their motivations.
8. **Review**: Once your PR is reviewed and approved, it will be merged into the main branch. Congratulations on your contribution!
</details>

<details closed>
<summary>Contributor Graph</summary>
<br>
<p align="center">
   <a href="https://github.com{/chmenegatti/go-vrf.git/}graphs/contributors">
      <img src="https://contrib.rocks/image?repo=chmenegatti/go-vrf.git">
   </a>
</p>
</details>

---

##  License

This project is protected under the [SELECT-A-LICENSE](https://choosealicense.com/licenses) License. For more details, refer to the [LICENSE](https://choosealicense.com/licenses/) file.

---

##  Acknowledgments

- List any resources, contributors, inspiration, etc. here.

[**Return**](#-overview)

---
