<div align="center">

# GMSSH

### Manage Linux servers like Windows — without installing a desktop environment.

#### An SSH-based desktop workspace that makes Linux servers simple, secure, intelligent, and accessible.

[![Chinese](https://img.shields.io/badge/Lang-%E7%AE%80%E4%BD%93%E4%B8%AD%E6%96%87-brightgreen)](https://github.com/GMSSH/GMSSH/blob/main/README.cn.md)
[![English](http://img.shields.io/badge/Lang-English-blue)](https://github.com/GMSSH/GMSSH/blob/main/README.md)

<p>
  <a href="#what-is-gmssh">Introduction</a> •
  <a href="#why-does-gmssh-exist">Why GMSSH</a> •
  <a href="#gmssh-product-positioning">Product Positioning</a> •
  <a href="#gmssh-and-traditional-server-management-methods">Comparison</a> •
  <a href="#quick-start">Quick Start</a>
</p>

</div>

---

<img width="3600" height="2166" alt="GMSSH overview" src="https://github.com/user-attachments/assets/75ed8016-e8f7-49fe-827e-9c6bf2165917" />

## What is GMSSH?

GMSSH is a desktop-oriented Linux server management platform built on SSH.

It does not require a full desktop environment or an additional server control panel to be installed on the server. Once connected over SSH, users can manage files, use a terminal, access a server browser, run applications, and perform operations with GMAI from a familiar desktop workspace.

GMSSH is designed to make Linux servers easier to understand and use, especially for people who are not familiar with Linux command-line tools. It keeps the flexibility and security of SSH while providing a visual experience similar to using Windows.

GMSSH sits between native SSH, traditional server panels, and enterprise bastion hosts:

- **Native SSH** provides reliable, lightweight remote access and command-line capabilities;
- **Traditional server panels** provide visual management for websites, databases, and runtime environments, but usually require panel services to be installed and kept running on the server;
- **Enterprise bastion hosts** focus on centralized access control, permissions, auditing, and compliance, often with higher deployment and maintenance costs;
- **GMSSH** builds on SSH to provide a simpler, more visual, and more accessible way to connect to and use Linux servers.

GMSSH's core goal is:

> **Make the command line no longer a barrier to using Linux servers.**

---

## Why does GMSSH exist?

When a user manages more than one server, the challenge is no longer simply whether SSH can connect. The real challenge is how to organize, switch between, and use those servers efficiently.

Individual developers and small teams often encounter the following problems:

- SSH addresses, ports, usernames, and keys are scattered across terminals, scripts, and documents;
- As the number of servers grows, switching connections, grouping hosts, and distinguishing environments becomes confusing;
- Sharing host information with team members is inconvenient and can be difficult to control securely;
- Traditional server panels are good at managing websites, databases, and runtime environments, but are not ideal as a unified SSH access point;
- Enterprise bastion hosts provide comprehensive governance, but their deployment, learning, and maintenance costs may be excessive for individuals and small teams.

GMSSH fills the gap between native SSH and heavyweight operations systems:

> **It preserves the lightweight and secure nature of SSH while providing a clearer multi-server workflow and a more intuitive desktop experience.**

Users can connect to, switch between, and manage multiple Linux servers from one client without installing a full desktop environment or an additional management panel on every server.

---

## GMSSH product positioning

GMSSH is best understood as:

- A **desktop workspace for using Linux servers**;
- A **visual server management tool built around SSH**;
- A **simple and accessible entry point for Linux server operations**;
- A **lightweight workbench for individual developers and small to mid-sized teams**.

GMSSH is not intended to take over every server operation. Instead, it focuses on lowering the barrier to using Linux and making common server tasks easier to understand and complete:

- Make Linux servers easier for beginners to get started with;
- Replace repetitive command memorization with visual operations where appropriate;
- Use GMAI to explain commands, analyze problems, and assist with server tasks;
- Organize and switch between multiple servers from one desktop workspace;
- Preserve the flexibility of the terminal for experienced Linux users;
- Make Linux servers practical for more people, not only experienced system administrators.

> **GMSSH lets users manage Linux servers like Windows, while keeping SSH available whenever they need the command line.**

---

## Why choose GMSSH?

### 1. Easier multi-server management

When your infrastructure grows from a few machines to dozens or hundreds, the challenge is not only connecting to a server. It is keeping hosts organized, switching between environments, and completing daily work efficiently.

GMSSH provides a single desktop workspace for adding, grouping, searching, and switching between multiple Linux servers.

### 2. More accessible than native SSH

Native SSH is reliable, powerful, and lightweight, but it requires users to understand commands and maintain connection details across terminals, configuration files, scripts, and documents.

GMSSH keeps SSH as its foundation while adding a familiar visual workspace for files, terminals, browsers, applications, and server status. Beginners can start with visual tools, while experienced users can continue using the terminal when needed.

### 3. Lighter than traditional panels

GMSSH does not follow the all-in-one control-panel model and does not require a full desktop environment or a long-running management panel on every server.

It focuses on the SSH workflow and uses lightweight components on demand, reducing server-side intrusion, resource overhead, and operational complexity.

### 4. More approachable than heavyweight bastion hosts

Enterprise bastion hosts are designed for complex access governance, auditing, and compliance requirements. GMSSH is designed for individuals, small teams, and growing organizations that need a simpler way to use and manage Linux servers without the cost and complexity of a full bastion system.

### 5. Intelligent and extensible

GMAI can help explain commands, generate scripts, analyze logs and configuration, troubleshoot common problems, and assist with server tasks.

GMSSH also supports custom models, AI Skills, knowledge bases, applications, and plugins, allowing users to extend the workspace around their own workflows.

### 6. Open source and self-hostable

GMSSH is open source and supports self-hosted deployment. Users can choose a deployment model that fits their environment and retain control over their server connections, deployment configuration, and extension capabilities.

---

## Who is GMSSH for?

GMSSH is suitable for:

- Individual developers managing multiple Linux servers;
- Users who want to operate Linux more like Windows;
- Beginners who want to use Linux servers without first mastering the command line;
- Small teams and studios that need a unified SSH entry point;
- Operations and development engineers who want visual assistance for command-line work;
- Users who want AI to explain commands, analyze logs, and help troubleshoot server problems;
- Users who do not want to install a traditional control panel on every server;
- Teams that need a simpler alternative before adopting a heavyweight bastion host.

Connect over SSH to enter a server desktop and use files, terminals, browsers, applications, and GMAI from one workspace.

---

## GMSSH and traditional server management methods

| Dimension | Traditional SSH / Command Line | Traditional Server Panels | GMSSH |
|---|---|---|---|
| How it is used | Commands entered in a terminal | Server panel accessed through a web browser | Servers accessed through a desktop client or privately deployed web client |
| Deployment | No additional panel required on the server | Usually requires panel services and related components to be installed | SSH-based; no full desktop environment required |
| User experience | Flexible and efficient, but requires command-line knowledge | Visual, but workflows and entry points are tied to the panel | A Windows-like desktop workspace for files, terminals, browsers, and applications |
| Multi-server management | Connection profiles and switching are usually maintained manually | Commonly centered on one server at a time | Add, group, search, and switch between multiple servers in one workspace |
| Server-side impact | Low resource usage; no additional resident service | Panel services and related components usually run continuously | Lightweight components are used on demand to reduce overhead and intrusion |
| Security and access | SSH is secure, but connection information can become scattered | Often requires a separate panel login and an exposed web management port | Reuses SSH access and authentication without requiring an additional panel management port |
| Intelligent assistance | Depends on the user writing commands and scripts | Primarily depends on built-in features and plugins | GMAI, model configuration, AI Skills, and knowledge bases assist with server operations |
| Applications and extensions | Flexible, but installation and configuration are often manual | Mature application ecosystems and one-click deployment are common strengths | Application Center and an open plugin system, with the ecosystem continuing to grow |
| Best suited for | Users who are comfortable with Linux and the command line | Website, database, and runtime environment management | Individuals and small teams that want a simple, visual, multi-server workspace |

GMSSH does not aim to replace every server management method. It provides another entry point: the security and flexibility of SSH combined with a desktop experience that is easier to learn and use.

---

## Core capabilities

### 1. Server connection and organization

- Native SSH connections
- Unified management for multiple servers
- Host grouping, tagging, and search
- Organization by development, testing, production, and other environments
- A unified server access entry point
- Secure local storage for connection profiles
- Batch command execution across multiple servers

### 2. Desktop-oriented server operations

- No full desktop environment required on the server
- File management and SFTP file transfer
- Multi-window terminal
- Server browser
- Task, process, and service inspection
- Server status and resource monitoring
- A Windows-like way to operate Linux servers

### 3. GMAI-assisted operations

- AI conversations and command explanations
- Command completion and script generation
- Log and configuration analysis
- Server troubleshooting
- AI-assisted server tasks
- Custom models, AI Skills, and knowledge bases

### 4. Applications and extensions

- Application Center
- Docker management
- Website, database, and runtime environment management
- Installation and configuration of common server applications
- Open plugin system
- User-developed and custom plugins
- GMSSH API and third-party tool integrations

### 5. Security and deployment

- Communication over SSH tunnels
- No additional panel management port required
- Low-intrusion, low-resource operation
- Lightweight components started on demand
- Locally encrypted credential storage
- Open-source, self-hostable, and user-controlled deployment

---

## Screenshots

### Client
![GMSSH client](https://github.com/user-attachments/assets/77032159-103f-43ea-a5ed-67b74a4e31c7)

### Desktop
![GMSSH desktop](https://github.com/user-attachments/assets/c2d7af27-5439-41dc-9626-246644581751)

### File Management
![File Management](https://github.com/user-attachments/assets/133623af-d9e3-479d-a4e4-1248a724f277)

### Application Center
![Application Center](https://github.com/user-attachments/assets/b44e125d-2f9e-4b98-8c12-aa38604ac8fe)

### Terminal + AI
![Terminal and AI](https://github.com/user-attachments/assets/4f593feb-30a9-4a93-8f6b-8040ad548e03)

### GMAI
![GMAI](https://github.com/user-attachments/assets/087cbfff-c461-49f9-90ab-b91f987fe349)

### Server Browser
![Server Browser](https://github.com/user-attachments/assets/87c82f8a-a0b4-45eb-a53d-40e2d327ac67)

---

## Quick Start

### Option 1: Web version

Use GMSSH directly in your browser to quickly connect to and manage your servers.

Get started: [Web Version](https://web.gm.cn/newweb/#/)

### Option 2: Desktop client

Install the GMSSH desktop client on your PC.

Get started: [Download the client](https://www.gm.cn/download)

### Option 3: Enterprise private deployment

A private deployment edition is available for enterprise scenarios. Contact the GMSSH team to discuss your environment and requirements.

Learn more: [Enterprise private deployment](https://www.gm.cn/private)

---

## Roadmap

GMSSH is evolving from an SSH server management tool into a desktop access and operations platform for Linux servers.

Future development will focus on:

- **Better host organization** — multi-server grouping, tags, environment separation, and unified search;
- **More complete desktop operations** — file management, terminals, browsers, services, monitoring, and applications;
- **Smarter operations** — stronger GMAI support for command explanations, log analysis, troubleshooting, script generation, and task execution;
- **A more open application ecosystem** — Application Center, Docker management, plugins, AI Skills, and API extensions;
- **Better team collaboration** — server sharing, member management, permissions, and operation auditing;
- **More flexible automation** — workflow orchestration, batch operations, and custom automation tasks.

The long-term goal of GMSSH is:

> **To become the unified desktop entry point for accessing, managing, and using Linux servers.**

---

## Community and Support

Official website: <https://www.gm.cn>

Project repository: <https://github.com/GMSSH/GMSSH>

Bug reports: <https://github.com/GMSSH/GMSSH/issues>

Feature requests and community discussion: <https://hub.gm.cn/>

Documentation: [Product Documentation](https://doc.gm.cn/zh/guide/)

If GMSSH helps you, please consider giving the project a Star ✨.

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=GMSSH/GMSSH&type=date&legend=top-left)](https://www.star-history.com/#GMSSH/GMSSH&type=date&legend=top-left)

---

*GMSSH Team*
