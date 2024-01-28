# Awesome Hetzner Cloud

A curated list of libraries, tools, and integrations for [Hetzner Cloud](https://cloud.hetzner.com/).

Projects formatted in **bold** are official Hetzner Cloud projects.

* [Libraries](#libraries)
* [Tools](#tools)
* [Integrations](#integrations)

Want to have your project listed? Open an Issue!

**Please note that while we try to only include projects in this list which are
functional and of good quality, we cannot provide any guarantee that they actually
work, are complete, nor that they do not cause any harm to your system or your account.**

## Libraries
### .NET
* [Hetzner Cloud API for .NET](https://github.com/lk-code/hetzner-cloud-api-net) — Hetzner Cloud API for .NET is a .NET Standard 2.0 Library

* [Hetzner Cloud DotNet](https://github.com/ljchuello/Hetzner-Cloud-DotNet) — This library is developed in .NET Standard 2.1 and is compatible with all .NET and .NET Core implementations, it can also be used in Console projects, Web API, Class Library and even with Blazor WASM.


### Go
* **[hcloud-go](https://github.com/hetznercloud/hcloud-go) — Package hcloud is a library for the Hetzner Cloud API.**

* [hcdns](https://github.com/acim/hcdns) — Go client library for Hetzner DNS API.


### HCL
* [terraform-hcloud-k3s](https://github.com/identiops/terraform-hcloud-k3s) — Comprehensive module for provisioning a k3s Kubernetes cluster on Hetzner Cloud.
* [terraform-kubernetes-hcloud-csi-driver](https://github.com/colinwilson/terraform-kubernetes-hcloud-csi-driver) — A simple module to provision the Hetzner Container Storage Interface Driver within a Kubernetes cluster running on Hetzner Cloud. See the variables file for the available configuration options. Please note that this module requires Kubernetes 1.15 or newer.


### Java
* [hetznercloud-java](https://github.com/TomSDEVSN/hetznercloud-java) — Simple Java client for the Hetzner Cloud API.


### JavaScript
* [hcloud-nodejs](https://github.com/LeonSpors/hcloud-nodejs) — hcloud-nodejs is a node.js package for Hetzner cloud that can be used to manage your projects.

* [hcloud-js](https://github.com/dennisbruner/hcloud-js) — A Node.js module for the Hetzner Cloud API


### PHP
* [Hetzner Cloud PHP SDK](https://github.com/LKDevelopment/hetzner-cloud-php-sdk) — A PHP SDK for the Hetzner Cloud API: docs.hetzner.cloud


### Python
* **[hcloud-python](https://github.com/hetznercloud/hcloud-python) — hcloud-python is a library for the Hetzner Cloud API.**


### Ruby
* [hcloud-ruby](https://github.com/tonobo/hcloud-ruby) — Native ruby client for HetznerCloud

* [hetznercloud](https://github.com/floriandejonckheere/hcloud) — Modern object-oriented Ruby wrapper for Hetzner Cloud API


### Rust
* [hcloud-rust](https://github.com/HenningHolmDE/hcloud-rust/) — Unofficial Rust crate for accessing the Hetzner Cloud API



## Tools
* **[hcloud cli](https://github.com/hetznercloud/cli) — hcloud is a command-line interface for interacting with Hetzner Cloud.**

* **[setup-hcloud](https://github.com/hetznercloud/setup-hcloud) — GitHub action to install the Hetzner Cloud CLI.**

* [hcloud-failover-keepalived](https://github.com/lehuizi/hcloud-failover-keepalived) — Script for switching Floating IPs in case of keepalived failover

* [hetzner-kube](https://github.com/xetys/hetzner-kube) — This project contains a CLI tool to easily provision kubernetes clusters on Hetzner Cloud.

* [HC Volume Backup](https://gitlab.com/MartinBoehmer/hc-volume-backup) — Bash script to backup Hetzner Cloud Volumes. Automatically creates volumes for backups and maintains a defined number of them.

* [TestFlows GitHub Hetzner Runners](https://github.com/testflows/TestFlows-GitHub-Hetzner-Runners) — Autoscaling GitHub Actions self-hosted runners using Hetzner Cloud

* [KubeOne](https://github.com/kubermatic/kubeone) — Kubermatic KubeOne automates cluster operations on hetzner cloud. KubeOne can install high-available (HA) master clusters as well single master clusters.

* [Hetzner Load Balancer Prometheus Exporter](https://github.com/infraduckture/hetzner-load-balancer-prometheus-exporter) — Exports metrics from Hetzner Load Balancer for consumption by Prometheus

* [Healthzner Bot](https://github.com/raphaelbernhart/healthznerbot) — A discord bot to get periodically updates of the health status of your hetzner cloud machines.

* [hcloud-pricing-exporter](https://github.com/jangraefen/hcloud-pricing-exporter) — A Prometheus exporter that connects to your HCloud account and collects data on your current expenses.

* [Kubernetes Hetzner Keepalived](https://github.com/schemen/kubernetes-hetzner-keepalived) — K8s deployment and image to create a keepalived ip failover with the floating ip feature.

* [Hetzner Cloud Connect](https://github.com/BlueBambooStudios/hcloud-connect) — Handles automatically adding servers to load balancers

* [Hcloud Snapshot-as-Backup](https://github.com/fbrettnich/hcloud-snapshot-as-backup) — Hetzner Cloud - Automatic Snapshots as Backups for more flexibility

* [chaotic](https://github.com/ngine-io/chaotic) — Fault injection your Hetzner Cloud servers to ensure, your services run even with one server stopped.

* [Hetzner rescaler](https://github.com/jonamat/hetzner-rescaler) — Lightweight CLI tool to programmatically rescale your Hetzner Cloud server.

* [cluster-api-provider-hetzner](https://github.com/syself/cluster-api-provider-hetzner) — Kubernetes Cluster API Provider for consistent deployment and day 2 operations of &#34;self-managed&#34; HA Kubernetes clusters on Hetzner.

* [Ansible Role hcloud](https://github.com/ngine-io/ansible-role-hcloud) — Ansible Role for managing hcloud cloud resources.

* [kOps](https://github.com/kubernetes/kops) — Kubernetes Operations (kOps) automates Kubernetes cluster installation, upgrades and management

* [hobby-kube](https://github.com/hobby-kube/guide) — Fully automated cluster setup using Terraform, good balance between resilience and cost, and therefore a great starting point for hobbyists or to build a professional setup with a reasonable foundation.

* [hcloud-selfdestruct](https://github.com/worldworm/hcloud-selfdestruct) — CLI tool to self destruct a hetzner cloud server.

* [kubermatic-kubernetes-platform](https://docs.kubermatic.com/kubermatic) — Kubermatic Multi-Cluster and Mulit-Cloud management tool for self hosted cluster as a service purposes.

* [vitobotta/hetzner-k3s](https://github.com/vitobotta/hetzner-k3s) — A CLI tool written in Crystal to quickly create and manage Kubernetes clusters in Hetzner Cloud

* [terraform-hcloud-kube-hetzner](https://github.com/kube-hetzner/terraform-hcloud-kube-hetzner) — A highly optimized and auto-upgradable, HA-default &amp; Load-Balanced, Kubernetes cluster powered by k3s-on-MicroOS and deployed for peanuts on Hetzner Cloud 🤑 🚀

* [janosmiko/hetzner-k3s](https://github.com/janosmiko/hetzner-k3s) — CLI tool to install and manage production grade lightweight Kubernetes (k3s) clusters in 5 minutes in Hetzner. Features: Hetzner CCM and CSI support, HA, Multiple Worker Pools, Autoscaling etc.

* [terraform-hks](https://github.com/Stupremee/terraform-hcloud-hks) — An opinionated Terraform module for deploying a Hetzner Kubernetes Cluster using RKE2 and Hetzner Cloud.

* [Solidblocks Hetzner Nuke](https://pellepelster.github.io/solidblocks/hetzner/nuke/) — Solidblocks Hetzner Nuke is a tool to delete all resources in a Hetzner account, similar to aws-nuke.

* [1Password Shell Plugin](https://developer.1password.com/docs/cli/shell-plugins/hetzner-cloud/) — The Hetzner Cloud shell plugin allows you to use 1Password to securely authenticate hcloud CLI with your fingerprint, Apple Watch, or system authentication, rather than storing your credentials in plaintext.

* [hcloud assign ip](https://github.com/lehuizi/hcloud_assign_ip) — A little helper to easily assign a floating ip to the current system


## Integrations
* **[Terraform Provider](https://github.com/hetznercloud/terraform-provider-hcloud) — Official Hetzner Cloud Terraform Provider**

* **[hcloud-cloud-controller-manager](https://github.com/hetznercloud/hcloud-cloud-controller-manager) — Kubernetes cloud-controller-manager for Hetzner Cloud**

* **[hcloud-csi](https://github.com/hetznercloud/csi-driver) — Container Storage Interface driver for Hetzner Cloud**

* **[Molecule driver for Hetzner Cloud](https://github.com/ansible-community/molecule-hetznercloud) — A molecule driver allowing you to use on-demand Hetzner Cloud servers for your tests.**

* **[Kubernetes Cluster Autoscaler](https://github.com/kubernetes/autoscaler) — A k8s component that automatically adjusts the size of a Kubernetes Cluster so that all pods have a place to run and there are no unneeded nodes.**

* **[Hetzner Cloud Packer Builder](https://github.com/hetznercloud/packer-plugin-hcloud) — An official Packer.io Builder for Hetzner Cloud**

* **[Prometheus Service Discovery](https://github.com/prometheus/prometheus) — A service discovery build into Prometheus to easily discover Hetzner servers (Cloud &amp; Dedicated)**

* **[Hetzner Cloud Ansible Collection](https://github.com/ansible-collections/hetzner.hcloud) — Ansible Collection for Hetzner Cloud. Part of Ansible Community Distribution (ACD).**

* [Rancher on Hetzner Cloud](https://github.com/alexzimmer96/rancher-hcloud) — Prebuild Terraform templates for deploying a highly available RKE cluster on Hetzner Cloud and installing Rancher into it

* [Paymenter Server Extension](https://github.com/ha1fdan/HetznerCloudExtension) — Simplify Hetzner Cloud server orders with this Paymenter.org extension.

* [docker-machine-driver-hetzner](https://github.com/JonasProgrammer/docker-machine-driver-hetzner) — This library adds the support for creating Docker machines hosted on the Hetzner Cloud.

* [Docker Volume Hetzner](https://github.com/costela/docker-volume-hetzner) — Volume management plugin for Docker (and Swarm)

* [ui-driver-hetzner](https://github.com/mxschmitt/ui-driver-hetzner) — Rancher UI driver for the Hetzner Cloud docker-machine-driver

* [Ploi](https://ploi.io) — Rapidly deploy any site you like on Hetzner Cloud servers.

* [iTop Datacollector for Hetzner Cloud](https://github.com/itomig-de/itomig-hetzner-collector) — This stand-alone application collects information from Hetzner projects in order to automatically synchronize the farms and virtual machines in iTop. iTop is a web based open source tool for IT service management tasks.

* [hcloud IP Floater](https://github.com/costela/hcloud-ip-floater) — Kubernetes controller to attach hcloud floating IPs to services&#39; nodes.

* [Versio-io](https://www.versio.io/import-hetzner-cloud-cmdb-configuration-item.html) — Integration of Hetzner Cloud configuration items (CI) in the full stack Versio.io configuration management database (CMDB).

* [Algo VPN](https://github.com/trailofbits/algo) — Algo VPN is a set of Ansible scripts that simplify the setup of a personal Wireguard and IPSEC VPN.

* [hcloud fip controller](https://github.com/cbeneke/hcloud-fip-controller) — Kubernetes controller to (re-)assign floating IPs to Hetzner Cloud instances.

* [hcloud-freebsd](https://github.com/paulc/hcloud-freebsd) — Hetzner Cloud auto-provisioning for FreeBSD

* [hcloud-cloud-controller-manager-helm-chart](https://gitlab.com/MatthiasLohr/hcloud-cloud-controller-manager-helm-chart) — Hetzner Cloud - Cloud Controller Manager Helm Chart

* [hcloud-csi-driver-helm-chart](https://gitlab.com/MatthiasLohr/hcloud-csi-driver-helm-chart) — Community Helm Chart for Hetzner Cloud CSI Driver for Kubernetes

* [libcluster_hcloud](https://github.com/EightSQ/libcluster_hcloud) — This is a Hetzner Cloud clustering strategy for libcluster

* [Laravel Forge](https://forge.laravel.com/) — Laravel Hosting &amp; Instant PHP Servers on Hetzner Cloud

* [GitLab Hetzner Runner](https://git.shivering-isles.com/shivering-isles/gitlab-hetzner-runner) — A version of the gitlab-runner base image, that allows the usage of Hetzner Cloud instances for GitLab CI.

* [hcloud-tg](https://github.com/Navid2zp/hcloud-tg) — Telegram bot for managing Hetzner cloud servers.

* [terraform-kubernetes-hcloud-controller-manager](https://github.com/colinwilson/terraform-kubernetes-hcloud-controller-manager) — A simple module to provision the Hetzner Cloud Controller Manager (With Network &amp; Load Balancer Support) inside a Kubernetes cluster running on Hetzner Cloud. See the variables file for the available configuration options. Please note that this module requires Kubernetes 1.16 or newer.

* [terraform-hcloud-routeros-router](https://github.com/selfscrum/terraform-hcloud-routeros-router) — Terraform code to seamlessly integrate a RouterOS router into a Hetzner Cloud network.

* [terraform-hcloud-routeros-router-configuration](https://github.com/selfscrum/terraform-hcloud-routeros-router-configuration) — Configures a Hetzner server with an RouterOS Router.

* [ansible-hcloud-inventory](https://github.com/hg8496/ansible-hcloud-inventory) — An dynamic inventory script for hetzner cloud

* [Hetzner Cloud Deploy Server Github Action](https://github.com/TimDaub/hetzner-cloud-deploy-server-action) — Deploy a Hetzner Cloud Server from a GitHub Action.

* [scalr](https://github.com/ngine-io/scalr) — Autoscaling for Clouds - Scale Cloud instances based on policy checks in a configurable interval. With Hetzner Cloud and Prometheus support.

* [Hetzner Cloud Plugin for Jenkins](https://github.com/jenkinsci/hetzner-cloud-plugin) — The Hetzner cloud plugin enables Jenkins CI to schedule builds on dynamically provisioned VMs in Hetzner Cloud.

* [hetzner-machine-provider](https://github.com/bonsai-oss/hetzner-machine-provider) — Gitlab-runner custom driver executing jobs on plain Hetzner Cloud machines like GitHub actions.

* [Solidblocks RDS PostgreSQL](https://registry.terraform.io/modules/pellepelster/solidblocks-rds-postgresql/hcloud/latest) — A containerized PostgreSQL database with all batteries included backup solution powered by pgBackRest

* [hetzner-dyndns](https://github.com/marvinruder/hetzner-dyndns) — A proxy server for updating DNS records on Hetzner DNS using the DynDNS protocol.

* [Coder Template](https://github.com/ntimo/coder-hetzner-cloud-template) — A Terraform template for [Coder](https://github.com/coder/coder) to setup a cloud instance as dev environment with or without VS Code.

* [grafana-hcloud-datasource](https://github.com/apricote/grafana-hcloud-datasource) — Metrics for your Hetzner Cloud Servers and Load Balancers in Grafana.


## License

[Public Domain (CC0)](https://creativecommons.org/publicdomain/zero/1.0/)
