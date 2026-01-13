# Deployment Guide

GMSSH Private Deployment offers both online and offline versions.
For the online version, visit the official website:
https://www.gmssh.com/private
Offline version instructions:

## Prerequisites
- **OS**: Linux Server
- **Kernel Version**: 3.10+

## Installation

1. **Download**: Download the latest stable version installation package from an environment with internet access.
2. **Transfer**: Transfer the package to the target offline deployment server.
3. **Unzip**: Navigate to the installation directory and run the following command to unzip the package:
   ```bash
   tar xzf <package_name>.tar.gz -C .
   ```

## Configuration & Management

### Port Configuration
The service defaults to **TCP port 80**. 

If there is a port conflict, you can update the service port and automatically start the service by running:
```bash
bash ./deploy.sh setport <port_number>
```
> [!IMPORTANT]
> When setting the port, you **MUST** provide the port number directly after the command!

### Service Control

- **Start Service**:
  If the service is not running, execute:
  ```bash
  bash ./deploy.sh start
  ```

- **Check Status**:
  To check the status of the service, execute:
  ```bash
  bash ./deploy.sh status
  ```

- **Initialization**:
  In some cases, you may need to initialize the service before starting it. Run:
  ```bash
  bash ./deploy.sh init
  ```
  After initialization, proceed to start the service.

> [!NOTE]
> The above commands may require **root** or **sudo** privileges.

## Network Security

Ensure that your firewall or security group rules allow traffic on the configured TCP port (default is 80, or your custom port).

## Access

Once the status is normal, access the service via your browser:
`http://<Server_IP>:<Port>`
