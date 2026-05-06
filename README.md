<div align="center">

# GMSSH

### An SSH management and server access platform for modern operations

More efficient than native SSH, lighter and easier to adopt than traditional control panels.

[![Chinese](https://img.shields.io/badge/Lang-%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87-brightgreen)](https://github.com/GMSSH/GMSSH/blob/main/README.cn.md)
[![English](http://img.shields.io/badge/Lang-English-blue)](https://github.com/GMSSH/GMSSH/blob/main/README.md)

<p>
  <a href="#what-is-gmssh">Introduction</a> •
  <a href="#why-choose-gmssh">Why GMSSH</a> •
  <a href="#gmssh-vs-traditional-panels-and-systems">Comparison</a> •
  <a href="#quick-start">Quick Start</a>
</p>

</div>

---

## What is GMSSH?

GMSSH is neither a traditional server control panel nor a heavyweight bastion host system designed for large organizations. It is a **lightweight SSH-native server management platform** built for developers, DevOps engineers, and small to mid-sized teams.

If:

- **Native SSH** provides the most basic connection and command-line capability
- **Traditional panels** provide long-running visual management on the server
- **Bastion host systems** provide centralized access control, auditing, and governance

 **GMSSH focuses on a different high-frequency problem: how to connect to, view, operate, and collaborate on servers more efficiently and intuitively, based on SSH**.

---

## Why GMSSH?

As the number of servers grows, many teams run into the same issues:

- SSH addresses, ports, usernames, and keys are scattered across terminals, scripts, and documents
- As server count increases, switching connections, grouping hosts, and distinguishing environments becomes messy
- Sharing host access information across team members is inconvenient and often insecure
- Traditional panels are good at environment and application management, but not ideal as a unified SSH access entry
- Bastion systems are powerful, but often come with higher deployment, learning, and maintenance costs

GMSSH aims to provide a solution **between native SSH and heavyweight operations systems** — lighter, clearer, and more suitable for daily use.

---

## Product Positioning

GMSSH is best understood as:

- A **unified server access entry**
- An **SSH-centered host management tool**
- A **lightweight operations platform for individual developers and small teams**

Its focus is not to “take over everything on the server,” but to:

- Make connection management more organized
- Make host access more efficient
- Make team collaboration more controllable
- Make remote operations entry more lightweight

> **In one sentence: traditional panels manage server internals, while GMSSH manages the server access and connection system.**

---

## Why choose GMSSH?

### 1. Better for multi-host management
When your infrastructure grows from a few machines to dozens or hundreds, the real challenge is no longer whether you can connect — it is how to organize, switch, manage, and collaborate efficiently.

### 2. Lighter than traditional panels
GMSSH does not follow the “all-in-one control panel” approach. It does not aim to replace every operations task. Instead, it focuses on the SSH workflow itself, reducing invasiveness and resource overhead.

GMSSH connects through SSH and starts lightweight components only when needed, minimizing long-running processes and additional exposure on the server.

### 3. Clearer than native SSH
Native SSH is powerful and reliable. But when the number of hosts, users, keys, environments, and team members grows, connection details and operation records become fragmented.

GMSSH keeps the SSH model while providing clearer host management, grouping, access entry, and visualized operations experience.

### 4. Easier to adopt than heavyweight bastion systems
Lower learning cost, lower deployment cost, and lower maintenance burden — making it better suited for individuals, small teams, and growing organizations.

### 5. Open source, self-hosted, and fully controllable
Users can choose the deployment and usage model that fits their environment. Host information, access methods, and future extensibility remain under your control.

---

## What problems does GMSSH solve?

GMSSH mainly addresses the following common problems:

- **Scattered connection information**: host addresses, ports, keys, and accounts are stored in multiple places and are hard to maintain centrally
- **Multi-host management chaos**: as the number of machines grows, there is no clear organization across environments, projects, and services
- **Inefficient team collaboration**: access details are shared through chat messages, verbal communication, or manual documents, which is inefficient and insecure
- **Misaligned panel capabilities**: many panels are better at managing server environments than managing SSH connection relationships
- **Overly heavy traditional solutions**: enterprise-grade systems are powerful, but often exceed the actual needs of many teams

---

## Who is GMSSH for?

GMSSH is suitable for:

- Individual developers managing multiple Linux servers
- Small teams that need a unified SSH access entry
- Teams that primarily rely on command-line operations but want better host organization
- Users who do not need a heavyweight bastion host yet, but do need connection management and collaboration features
- Scenarios where servers are managed through SSH instead of installing traditional panels on every machine

---

## GMSSH vs traditional panels and systems

| Dimension | Native SSH | Traditional Panels | Modern Panels | Bastion Hosts / Heavy Ops Systems | GMSSH |
|---|---|---|---|---|---|
| Core positioning | Command-line remote access | Server environment and site management | Panel-based service and app management | Centralized ops, security auditing, permission control | SSH access and multi-host management platform |
| Single-host management | High | High | High | Medium | High |
| Multi-host management | Low | Medium | Medium | High | High |
| SSH experience | High | Medium | Medium | High | High |
| Integrated terminal and file management | Low | Medium | Medium | Medium | High |
| Team sharing of host information | Low | Medium | Medium | High | Improving |
| Permissions and auditing | Low | Medium | Medium | High | Improving |
| Deployment complexity | Low | Medium | Medium | High | Low |
| Long-running components on server | None | Yes | Yes | Usually yes | Lightweight on-demand agent |
| Extra management ports required | No | Usually yes | Usually yes | Depends on architecture | No extra management ports required |
| Suitable for individuals / small teams | Medium | High | High | Low to Medium | High |
| Resource usage | Low | Medium | Medium | Medium to High | Low |
| Application environment configuration | Low | High | High | Medium | Extensible |
| Unified access entry | Low | Medium | Medium | High | High |
| Open source / self-hosted | Depends on tool | Depends on product | Depends on product | Depends on product | Supported |

---

## How to choose?

If you need:

- **Website environment setup, database GUI management** → choose a traditional panel
- **Enterprise-level auditing, approval workflows, compliance controls** → choose a bastion host system
- **More efficient SSH management and multi-host collaboration** → choose **GMSSH**

---

## Core capabilities

> The following items can be adjusted based on current product status.

- Host asset management
- SSH connection management
- Grouping and tagging
- Multi-environment host organization
- Unified access entry
- Team collaboration support
- Permission control
- Secure access mechanisms
- Open-source self-hosted deployment

---

## Screenshots

### Client
<img width="2880" height="1800" alt="客户端英文版" src="https://github.com/user-attachments/assets/396513c4-e795-4e69-9234-bae463a3490b" />

### Desktop
<img width="1440" height="775" alt="英文版" src="https://github.com/user-attachments/assets/0090af5f-0659-414b-baef-4b2ab5c64ef7" />

### App Center
<img width="1440" height="777" alt="截屏2026-05-06 18 03 44" src="https://github.com/user-attachments/assets/9591dcbe-5cd4-4316-8bfd-84d671401402" />

### Terminal + AI
<img width="820" height="559" alt="截屏2026-05-06 18 04 34" src="https://github.com/user-attachments/assets/c593af72-7110-431b-a6d1-1023b709fae8" />

### Gemius 2.0
<img width="724" height="717" alt="截屏2026-05-06 18 04 59" src="https://github.com/user-attachments/assets/9db08aba-3d4d-4b67-b7f9-e0891659f142" />

---

## Quick Start

### Option 1: Deploy with Docker

Deployment page: https://www.gm.cn/private

### Option 2: Read the documentation

Documentation: https://doc.gm.cn/zh/guide/

Quick start: [Web Version](https://web.gm.cn/newweb/#/)

---

## Roadmap

GMSSH is evolving from an open-source tool into a more complete server access and management platform.

Future improvements will continue to focus on:

- Better host organization capabilities
- More fine-grained permission control
- Clearer team collaboration mechanisms
- Richer operation auditing features
- More flexible automation and extensibility

---

## Why is GMSSH worth watching?

Because it fills a long-standing gap between several traditional solutions:

- More organized than native SSH
- Lighter and better suited for connection management than traditional panels
- Simpler, lower-cost, and easier to adopt than bastion host systems

The value of GMSSH is not in “taking over everything on the server,” but in this:

> Making server access itself simpler, clearer, and more professional.

---

## Community and Support

Project homepage: https://github.com/GMSSH/GMSSH

Issue reporting: Issues

Feature requests: Issues / Discussions

Documentation: [Product Docs](https://doc.gm.cn/zh/guide/)

If this project helps you, feel free to give it a Star.

---

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=GMSSH/GMSSH&type=date&legend=top-left)](https://www.star-history.com/#GMSSH/GMSSH&type=date&legend=top-left)

---

*Happy coding!*

*GMSSH Team*
