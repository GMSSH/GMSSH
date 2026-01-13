# Welcome to the GMSSH Open Ecosystem 👋
First off, thank you for considering contributing to the GMSSH Ecosystem! It's people like you that make the server management world more efficient and enjoyable.

<img width="2560" height="1314" alt="截屏2026-01-06 21 11 55" src="https://github.com/user-attachments/assets/c454ca21-9aca-41d1-8432-5e355c407b3d" />
<table width="100%">
  <tr>
    <td>
GMSSH is a Desktop AI System High-Perf · Zero Intrusion · AI-Powered tool designed for efficiency and aesthetics. 
While our core engine remains proprietary to ensure stability, we believe in the power of community extension.

## 🚀 Build for GMSSH

We provide a powerful SDK that allows you to extend the capabilities of GMSSH using web technologies (HTML/JS/Vue/React) or Python/Go scripts.

## ❤️Helping Open Source

In addition to creating and maintaining Open Source projects, GMSSH also contributes, sponsors (both financially and via licenses) many Open Source projects. Find out more about how GMSSH helps Open Source.

## 📜Code of Conduct
This code of conduct outlines our expectations for all those who participate in our open source projects and communities (community programs), as well as the consequences for unacceptable behaviour. We invite all those who participate to help us create safe and positive experiences for everyone. Communities mirror the societies in which they exist and positive action is essential to counteract the many forms of inequality and abuses of power that exist in society.

✅How to behave

The following behaviours are expected and requested of all community members:

Participate in an authentic and active way. In doing so, you contribute to the health and longevity of this community.
Exercise consideration, respect and empathy in your speech and actions. Remember, we have all been through different stages of learning when adopting technologies.
Refrain from demeaning, discriminatory, or harassing behaviour and speech.
Disagreements on things are fine, argumentative behaviour or trolling are not.

🚫How not to behave

Do not perform threats of violence or use violent language directed against another person.
Do not make jokes of sexist, racist, homophobic, transphobic, ableist or otherwise discriminatory nature, or use language of this nature.
Do not post or display sexually explicit or violent material.
Do not post or threaten to post other people’s personally identifying information ("doxing").
Do not make personal insults, particularly those related to gender, sexual orientation, race, religion, or disability.
Do not engage in sexual attention. This includes, sexualised comments or jokes and sexual advances.
Do not advocate for, or encourage, any of the above behaviour.
Please take into account that online communities bring together people from many different cultures and backgrounds. It's important to understand that sometimes the combination of cultural differences and online interaction can lead to misunderstandings. That is why having empathy is very important.


Because GMSSH adopts a **"Core Proprietary + Ecosystem Open"** strategy, this guide focuses on how to contribute to the **SDK**, **Official Apps**, and **Documentation**.

## Table of Contents
1. [Code of Conduct](#code-of-conduct)
2. [What Can I Contribute?](#what-can-i-contribute)
3. [Developing a New App (The Workflow)](#developing-a-new-app)
4. [Submitting Bugs & Features](#submitting-bugs--features)
5. [Pull Request Guidelines](#pull-request-guidelines)
6. [App Center Publishing](#app-center-publishing)

---

## Code of Conduct
By participating in this project, you agree to abide by our [Code of Conduct](LINK_TO_YOUR_CODE_OF_CONDUCT). We expect everyone to treat others with respect and empathy.

## What Can I Contribute?

We welcome contributions in several forms:

* **🔌 New Apps/Plugins:** Build a new tool using our SDK to solve a specific operation problem (e.g., Redis GUI, Log Analyzer).
* **🐛 Bug Fixes:** Fix issues in the `gmssh-plugin-sdk` or `gmssh-official-apps`.
* **📝 Documentation:** Improve the SDK documentation or translate it into other languages.
* **💡 Feature Requests:** Suggest new APIs for the SDK (e.g., "We need an API to access the local file system").

> **Note:** The core binary of GMSSH Client is closed source. Please do not reverse engineer or decompile the core logic. Focus on the JavaScript/Python/GO extension layer.

## Developing a New App

**🚀 Fast Track: Jump Straight into Coding**
Want to build an app immediately? If you prefer to skip the architectural theory and start your developer journey right now, click the link below:

[👉 Start Your Developer Journey](https://doc-dev.gmssh.com/en/)

**📖 Understanding GMSSH Design Philosophy**
If you prefer to grasp the underlying architecture before writing code, here is a deep dive into the core concepts:

**1. 🏗️ GMSSH Development Workflow**
To start, one must understand the core architectural philosophy. GMSSH is far more than just a Web server—it is an intricate process orchestration system.

**2. 🧠 The Core Engine: ga_main**
Within GMSSH's topology, ga_main is the cornerstone, serving as the system's "central nervous system" and routing nexus.

Distinct from LAMP: Unlike Apache/Nginx which often process business logic directly, ga_main is architected as a lightweight, high-performance microkernel.

**3. 🛡️ Process Isolation & Stability**
ga_main is fully decoupled from business logic; its sole purpose is lifecycle management and traffic dispatching. All features—from the official File Manager to custom Nginx Managers—run as isolated child processes.

Sandboxed Reliability: If a third-party plugin crashes (e.g., due to memory leaks), it only terminates that single process. ga_main and other services remain unaffected.

Auto-Recovery: The system actively monitors processes and handles restarts automatically.

**4. 🌉 Gateway & Protocol Translation**
ga_main functions as the ingress gateway, aggregating all traffic from client interfaces (Electron/Web).

Responsibilities: It handles Request Validation (AuthN/AuthZ), App ID resolution, and dispatching.

Transparent Bridge: It seamlessly converts external HTTP/WebSocket protocols into internal IPC (Inter-Process Communication).

**5. ⚡ Kernel-level Data Transmission**
While standard TCP/IP incurs overhead (segmentation, checksums, firewall filtering), GMSSH leverages Unix Domain Sockets (UDS) for superior performance. Data is copied directly within Kernel Memory Buffers, bypassing the network stack entirely.

🚀 Zero-Copy Performance: Drastically reduces CPU context switches.

⏱️ Low Latency: Microsecond-level latency (vs. TCP's millisecond-level), perfect for "Real-time Terminals" or "Log Streams."

🔒 Security: Relies on file system permissions. Only ga_main and the App owner can access the socket, fundamentally eliminating port scanning risks.

**6. 🔄 Dual-Mode Startup Support**
Production (Recommended): Use UDS for maximum security and performance.

Development/Debug: GMSSH supports backend services via HTTP ports, facilitating development on Windows or cross-node deployment scenarios where UDS isn't natively supported.

**7. 📜 JSON-RPC 2.0 Protocol**
At the application layer, GMSSH adopts JSON-RPC 2.0—a stateless, lightweight RPC protocol. This makes the backend language-agnostic (Python, Go, Rust, Node.js, etc.).

Request Structure:

jsonrpc: Version

method: Function name

params: Arguments (object/array)

id: Unique identifier

Response Structure:

Returns execution result or an error object.


## Submitting Bugs & Features

We use GitHub Issues to track bugs and features.

* **Bugs:** Please use the **Bug Report** template. Attach screenshots and the `gmssh-sdk` version you are using.
* **Features:** Please use the **Feature Request** template. clearly describe the "Pain Point" and your proposed solution.

## Pull Request Guidelines

Ready to submit your code? Great!

1.  **Sync First:** Ensure your fork is up-to-date with the `main` branch.
2.  **Descriptive Title:** Use a clear title (e.g., `feat: add support for dark mode in Redis plugin`).
3.  **Checklist:**
    * [ ] My code follows the code style of this project (ESLint/Prettier).
    * [ ] I have performed a self-review of my own code.
    * [ ] I have commented on my code, particularly in hard-to-understand areas.
    * [ ] I have verified that the app runs smoothly in GMSSH Client v2.0+.

## App Center Publishing

If you want your app to be listed in the **Official GMSSH App Center** (available to all users):

1.  Submit a Pull Request to the `gmssh-marketplace` registry (or submit a form on our website).
2.  **Security Review:** Our team will review your code to ensure no malicious operations (e.g., stealing SSH keys).
3.  **Approval:** Once approved, your app will be listed, and you will receive the **"GMSSH Contributor"** badge!

---
*Happy Coding!*

*The GMSSH Team*

