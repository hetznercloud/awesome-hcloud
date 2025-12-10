# Awesome Hetzner Cloud

A curated list of libraries, tools, and integrations for [Hetzner Cloud](https://cloud.hetzner.com/).

Projects formatted in **bold** are official Hetzner Cloud projects.

- [Libraries](#libraries)
- [Tools](#tools)
- [Integrations](#integrations)

Want to have your project listed? [Open an Issue](https://github.com/hetznercloud/awesome-hcloud/issues/new?assignees=&labels=&projects=&template=new-project.yml&title=Add+New+Project%3A+%3Cproject-name%3E)!

**Please note that while we try to only include projects in this list which are
functional and of good quality, we cannot provide any guarantee that they actually
work, are complete, nor that they do not cause any harm to your system or your account.**

## Libraries

### .NET

- [Hetzner Cloud API for .NET](https://github.com/lk-code/hetzner-cloud-api-net) — Hetzner Cloud API for .NET is a .NET Standard 2.0 Library
- [HetznerCloud.API](https://github.com/ljchuello/hetznercloud.api) — This library is developed in .NET Standard 2.0 and is compatible with all .NET and .NET Core implementations, it can also be used in Console projects, Web API, Class Library and even with Blazor WASM.

### Go

- **[hcloud-go](https://github.com/hetznercloud/hcloud-go) — Package hcloud is a library for the Hetzner Cloud API.**
- [hcdns](https://github.com/acim/hcdns) — Go client library for Hetzner DNS API.
- [hetzner-golang-actions](https://github.com/mrsimonemms/hetzner-golang-actions) — Hetzner Golang action completion - the missing part of the Hetzner Golang SDK.

### HCL

- [hcloud-user-data](https://github.com/wszychta/terraform-module.hcloud-user-data) — Terraform Module for ready to use user-data files for Hetzner cloud servers with multiple network managers
- [terraform-hcloud-hetzner-node-pool](https://github.com/hegerdes/terraform-hcloud-hetzner-node-pool) — Terraform module to easily create VMs in the Hetzner Cloud and groups them in pools for easy scaling.
- [terraform-kubernetes-hcloud-csi-driver](https://github.com/colinwilson/terraform-kubernetes-hcloud-csi-driver) — A simple module to provision the Hetzner Container Storage Interface Driver within a Kubernetes cluster running on Hetzner Cloud. See the variables file for the available configuration options. Please note that this module requires Kubernetes 1.15 or newer.

### Haskell

- [hetzner](https://codeberg.org/daniel-casanueva/hetzner) — Haskell library to interact with the Hetzner API.

### Java

- [hcloud-java](https://github.com/cubitsdev/hcloud-java) — hcloud-java is a fully tested and easy to use Java API Integration. It follows an Observer Pattern and focuses on minimal dependency bloat.
- [hetznercloud-java](https://github.com/tomsdevsn/hetznercloud-java) — Simple Java client for the Hetzner Cloud API.

### JavaScript

- [hcloud-js](https://github.com/dennisbruner/hcloud-js) — A Node.js module for the Hetzner Cloud API
- [hcloud-nodejs](https://github.com/leonspors/hcloud-nodejs) — hcloud-nodejs is a node.js package for Hetzner cloud that can be used to manage your projects.

### Kotlin

- [hetznerkloud](https://github.com/sasa-b/hetznerkloud) — Kotlin library for Hetzner Cloud API

### PHP

- [Hetzner Cloud PHP SDK](https://github.com/lkdevelopment/hetzner-cloud-php-sdk) — A PHP SDK for the Hetzner Cloud API: docs.hetzner.cloud

### Python

- **[hcloud-python](https://github.com/hetznercloud/hcloud-python) — hcloud-python is a library for the Hetzner Cloud API.**

### Ruby

- [hcloud-ruby](https://github.com/tonobo/hcloud-ruby) — Native ruby client for HetznerCloud
- [hetznercloud](https://github.com/floriandejonckheere/hcloud) — Modern object-oriented Ruby wrapper for Hetzner Cloud API

### Rust

- [hcloud-rust](https://github.com/henningholmde/hcloud-rust/) — Unofficial Rust crate for accessing the Hetzner Cloud API



## Tools

- **[hcloud cli](https://github.com/hetznercloud/cli) — hcloud is a command-line interface for interacting with Hetzner Cloud.**
- **[setup-hcloud](https://github.com/hetznercloud/setup-hcloud) — GitHub action to install the Hetzner Cloud CLI.**
- [1Password Shell Plugin](https://developer.1password.com/docs/cli/shell-plugins/hetzner-cloud/) — The Hetzner Cloud shell plugin allows you to use 1Password to securely authenticate hcloud CLI with your fingerprint, Apple Watch, or system authentication, rather than storing your credentials in plaintext.
- [Ansible Role hcloud](https://github.com/ngine-io/ansible-role-hcloud) — Ansible Role for managing hcloud cloud resources.
- [HC Volume Backup](https://gitlab.com/martinboehmer/hc-volume-backup) — Bash script to backup Hetzner Cloud Volumes. Automatically creates volumes for backups and maintains a defined number of them.
- [Hcloud Snapshot-as-Backup](https://github.com/fbrettnich/hcloud-snapshot-as-backup) — Hetzner Cloud - Automatic Snapshots as Backups for more flexibility
- [Healthzner Bot](https://github.com/raphaelbernhart/healthznerbot) — A discord bot to get periodically updates of the health status of your hetzner cloud machines.
- [Hetzner Cloud Connect](https://github.com/bluebamboostudios/hcloud-connect) — Handles automatically adding servers to load balancers
- [Hetzner Load Balancer Prometheus Exporter](https://github.com/infraduckture/hetzner-load-balancer-prometheus-exporter) — Exports metrics from Hetzner Load Balancer for consumption by Prometheus
- [Hetzner Nuke](https://github.com/cgroschupp/hetzner-nuke) — Hetzner Nuke is written in go, based on libnuke, which is used by aws-nuke and azure-nuke.
- [Hetzner rescaler](https://github.com/jonamat/hetzner-rescaler) — Lightweight CLI tool to programmatically rescale your Hetzner Cloud server.
- [KubeOne](https://github.com/kubermatic/kubeone) — Kubermatic KubeOne automates cluster operations on hetzner cloud. KubeOne can install high-available (HA) master clusters as well single master clusters.
- [Kubernetes Hetzner Keepalived](https://github.com/schemen/kubernetes-hetzner-keepalived) — K8s deployment and image to create a keepalived ip failover with the floating ip feature.
- [Solidblocks Hetzner Nuke](https://pellepelster.github.io/solidblocks/hetzner/nuke/) — Solidblocks Hetzner Nuke is a tool to delete all resources in a Hetzner account, similar to aws-nuke.
- [TestFlows GitHub Hetzner Runners](https://github.com/testflows/testflows-github-hetzner-runners) — Autoscaling GitHub Actions self-hosted runners using Hetzner Cloud
- [VitoDeploy](https://github.com/vitodeploy/vito) — Vito is a self-hosted web application that helps you to manage your servers and deploy your PHP applications into production servers without a hassle.
- [certmaster](https://github.com/poundifdef/certmaster) — Automatically creates Let's Encrypt certificates and uploads them to Hetzner Cloud Load Balancers. Useful if you do not use Hetzner DNS.
- [chaotic](https://github.com/ngine-io/chaotic) — Fault injection your Hetzner Cloud servers to ensure, your services run even with one server stopped.
- [cluster-api-provider-hetzner](https://github.com/syself/cluster-api-provider-hetzner) — Kubernetes Cluster API Provider for consistent deployment and day 2 operations of "self-managed" HA Kubernetes clusters on Hetzner.
- [hcloud assign ip](https://github.com/lehuizi/hcloud_assign_ip) — A little helper to easily assign a floating ip to the current system
- [hcloud-failover-keepalived](https://github.com/lehuizi/hcloud-failover-keepalived) — Script for switching Floating IPs in case of keepalived failover
- [hcloud-kubernetes](https://github.com/hcloud-k8s/terraform-hcloud-kubernetes) — Terraform Module to Deploy a Highly Available, Production-Ready Talos Kubernetes Cluster on Hetzner Cloud.
- [hcloud-pricing-exporter](https://github.com/jangraefen/hcloud-pricing-exporter) — A Prometheus exporter that connects to your HCloud account and collects data on your current expenses.
- [hcloud-selfdestruct](https://github.com/worldworm/hcloud-selfdestruct) — CLI tool to self destruct a hetzner cloud server.
- [hcloud-upload-image](https://github.com/apricote/hcloud-upload-image) — Quickly upload any raw disk images into your Hetzner Cloud projects!
- [hetzner-kube](https://github.com/xetys/hetzner-kube) — This project contains a CLI tool to easily provision kubernetes clusters on Hetzner Cloud.
- [hicloud](https://github.com/rtulke/hicloud) — aka hetzner interactive cloud is TUI for interacting with Hetzner Cloud by using commands, wizards and tab-completions including resources.
- [hobby-kube](https://github.com/hobby-kube/guide) — Fully automated cluster setup using Terraform, good balance between resilience and cost, and therefore a great starting point for hobbyists or to build a professional setup with a reasonable foundation.
- [janosmiko/hetzner-k3s](https://github.com/janosmiko/hetzner-k3s) — CLI tool to install and manage production grade lightweight Kubernetes (k3s) clusters in 5 minutes in Hetzner. Features: Hetzner CCM and CSI support, HA, Multiple Worker Pools, Autoscaling etc.
- [kOps](https://github.com/kubernetes/kops) — Kubernetes Operations (kOps) automates Kubernetes cluster installation, upgrades and management
- [kubermatic-kubernetes-platform](https://docs.kubermatic.com/kubermatic) — Kubermatic Multi-Cluster and Mulit-Cloud management tool for self hosted cluster as a service purposes.
- [prometheus-storagebox-exporter](https://github.com/crstian19/prometheus-storagebox-exporter) — Modern Prometheus exporter for Hetzner Storage Box with comprehensive metrics.
- [terraform-hcloud-k3s](https://github.com/identiops/terraform-hcloud-k3s) — Comprehensive module for provisioning a k3s Kubernetes cluster on Hetzner Cloud.
- [terraform-hcloud-kube-hetzner](https://github.com/kube-hetzner/terraform-hcloud-kube-hetzner) — A highly optimized and auto-upgradable, HA-default & Load-Balanced, Kubernetes cluster powered by k3s-on-MicroOS and deployed for peanuts on Hetzner Cloud 🤑 🚀
- [terraform-hcloud-talos](https://github.com/hcloud-talos/terraform-hcloud-talos) — Terraform module for creating a Kubernetes cluster with [Talos](https://www.talos.dev/) in the Hetzner Cloud.
- [terraform-hks](https://github.com/stupremee/terraform-hcloud-hks) — An opinionated Terraform module for deploying a Hetzner Kubernetes Cluster using RKE2 and Hetzner Cloud.
- [tg-cli](https://github.com/twingate-labs/tg-cli) — CLI to deploy [Twingate](https://www.twingate.com/) with support for Hetzner Cloud servers
- [vitobotta/hetzner-k3s](https://github.com/vitobotta/hetzner-k3s) — A CLI tool written in Crystal to quickly create and manage Kubernetes clusters in Hetzner Cloud
- [zfs-hetzner-vm](https://github.com/terem42/zfs-hetzner-vm) — Script to install Debian 10, 11, 12 and Ubuntu 18, 20, 22 LTS with ZFS root on Hetzner VPS.


## Integrations

- **[External DNS Hetzner Webhook](https://github.com/hetzner/external-dns-hetzner-webhook) — This webhook provides integration between External DNS and the Hetzner DNS API to manage records automatically.**
- **[Gitlab Fleeting plugin](https://gitlab.com/hetznercloud/fleeting-plugin-hetzner) — A GitLab fleeting plugin for Hetzner Cloud.**
- **[Hetzner Cloud Ansible Collection](https://github.com/ansible-collections/hetzner.hcloud) — Ansible Collection for Hetzner Cloud. Part of Ansible Community Distribution (ACD).**
- **[Hetzner Cloud Packer Builder](https://github.com/hetznercloud/packer-plugin-hcloud) — An official Packer.io Builder for Hetzner Cloud**
- **[Kubernetes Cluster Autoscaler](https://github.com/kubernetes/autoscaler) — A k8s component that automatically adjusts the size of a Kubernetes Cluster so that all pods have a place to run and there are no unneeded nodes.**
- **[Molecule driver for Hetzner Cloud](https://github.com/ansible-community/molecule-hetznercloud) — A molecule driver allowing you to use on-demand Hetzner Cloud servers for your tests.**
- **[Prometheus Service Discovery](https://community.hetzner.com/tutorials/prometheus-discovery) — A service discovery build into [Prometheus](https://github.com/prometheus/prometheus) to easily discover Hetzner servers (Cloud & Dedicated)**
- **[Terraform Provider](https://github.com/hetznercloud/terraform-provider-hcloud) — Official Hetzner Cloud Terraform Provider**
- **[cert-manager Webhook Hetzner](https://github.com/hetzner/cert-manager-webhook-hetzner) — This webhook provides integration between cert-manager and the Hetzner DNS API to handle DNS-01 challenges.**
- **[hcloud-cloud-controller-manager](https://github.com/hetznercloud/hcloud-cloud-controller-manager) — Kubernetes cloud-controller-manager for Hetzner Cloud**
- **[hcloud-csi](https://github.com/hetznercloud/csi-driver) — Container Storage Interface driver for Hetzner Cloud**
- [Algo VPN](https://github.com/trailofbits/algo) — Algo VPN is a set of Ansible scripts that simplify the setup of a personal Wireguard and IPSEC VPN.
- [Cloudfleet](https://cloudfleet.ai) — Managed Kubernetes platform with native Hetzner Cloud support for node-autoprovisioning.
- [Coder Template](https://github.com/ntimo/coder-hetzner-cloud-template) — A Terraform template for [Coder](https://github.com/coder/coder) to setup a cloud instance as dev environment with or without VS Code.
- [DevPod Provider](https://github.com/mrsimonemms/devpod-provider-hetzner) — Run [DevPod](https://github.com/loft-sh/devpod) cloud development environments on Hetzner.
- [Docker Volume Hetzner](https://github.com/costela/docker-volume-hetzner) — Volume management plugin for Docker (and Swarm)
- [GitLab Hetzner Runner](https://git.shivering-isles.com/shivering-isles/gitlab-hetzner-runner) — A version of the gitlab-runner base image, that allows the usage of Hetzner Cloud instances for GitLab CI.
- [Hetzner Cloud Deploy Server Github Action](https://github.com/timdaub/hetzner-cloud-deploy-server-action) — Deploy a Hetzner Cloud Server from a GitHub Action.
- [Hetzner Cloud Plugin for Jenkins](https://github.com/jenkinsci/hetzner-cloud-plugin) — The Hetzner cloud plugin enables Jenkins CI to schedule builds on dynamically provisioned VMs in Hetzner Cloud.
- [Laravel Forge](https://forge.laravel.com/) — Laravel Hosting & Instant PHP Servers on Hetzner Cloud
- [Paymenter Server Extension](https://github.com/ha1fdan/hetznercloudextension) — Simplify Hetzner Cloud server orders with this Paymenter.org extension.
- [Ploi](https://ploi.io) — Rapidly deploy any site you like on Hetzner Cloud servers.
- [Pulumi Hcloud Provider](https://www.pulumi.com/registry/packages/hcloud/) — A Pulumi Provider for setting up your infrastructure on hcloud with Pulumi.
- [Rancher on Hetzner Cloud](https://github.com/alexzimmer96/rancher-hcloud) — Prebuild Terraform templates for deploying a highly available RKE cluster on Hetzner Cloud and installing Rancher into it
- [ServerManagerBot](https://github.com/erfjab/servermanagerbot) — ServerManagerBot is a Telegram bot for managing Hetzner servers. It allows admins to control server actions.
- [Solidblocks RDS PostgreSQL](https://registry.terraform.io/modules/pellepelster/solidblocks-rds-postgresql/hcloud/latest) — A containerized PostgreSQL database with all batteries included backup solution powered by pgBackRest
- [Syself](https://syself.com) — A Kubernetes platform with self-healing, declarative management, GitOps compatibility and support for autoscaling and dedicated servers.
- [Versio-io](https://www.versio.io/import-hetzner-cloud-cmdb-configuration-item.html) — Integration of Hetzner Cloud configuration items (CI) in the full stack Versio.io configuration management database (CMDB).
- [ansible-hcloud-inventory](https://github.com/hg8496/ansible-hcloud-inventory) — An dynamic inventory script for hetzner cloud
- [docker-machine-driver-hetzner](https://github.com/jonasprogrammer/docker-machine-driver-hetzner) — This library adds the support for creating Docker machines hosted on the Hetzner Cloud.
- [gardener-extension-provider-hcloud](https://github.com/23technologies/gardener-extension-provider-hcloud) — Gardener Extension for Hetzner Cloud provider.
- [grafana-hcloud-datasource](https://github.com/apricote/grafana-hcloud-datasource) — Metrics for your Hetzner Cloud Servers and Load Balancers in Grafana.
- [hcloud IP Floater](https://github.com/costela/hcloud-ip-floater) — Kubernetes controller to attach hcloud floating IPs to services' nodes.
- [hcloud fip controller](https://github.com/cbeneke/hcloud-fip-controller) — Kubernetes controller to (re-)assign floating IPs to Hetzner Cloud instances.
- [hcloud-cloud-controller-manager-helm-chart](https://gitlab.com/matthiaslohr/hcloud-cloud-controller-manager-helm-chart) — Hetzner Cloud - Cloud Controller Manager Helm Chart
- [hcloud-csi-driver-helm-chart](https://gitlab.com/matthiaslohr/hcloud-csi-driver-helm-chart) — Community Helm Chart for Hetzner Cloud CSI Driver for Kubernetes
- [hcloud-freebsd](https://github.com/paulc/hcloud-freebsd) — Hetzner Cloud auto-provisioning for FreeBSD
- [hcloud-github-runner](https://github.com/cyclenerd/hcloud-github-runner) — On-demand self-hosted GitHub Actions Runner on Hetzner Cloud
- [hcloud-tg](https://github.com/navid2zp/hcloud-tg) — Telegram bot for managing Hetzner cloud servers.
- [hetzner-bare-metal-ansible](https://github.com/palark/hetzner-bare-metal-ansible) — Ansible playbook for deploying Linux bare-metal servers in Hetzner using Hetzner Robot API.
- [hetzner-dyndns](https://github.com/marvinruder/hetzner-dyndns) — A proxy server for updating DNS records on Hetzner DNS using the DynDNS protocol.
- [hetzner-machine-provider](https://github.com/bonsai-oss/hetzner-machine-provider) — Gitlab-runner custom driver executing jobs on plain Hetzner Cloud machines like GitHub actions.
- [hetzner-metal-kubernetes](https://github.com/cisnerosf/hetzner-metal-kubernetes) — Automates deploying Kubernetes (single-server or HA) on Hetzner dedicated servers with Fedora CoreOS using Ansible. Includes network and security setup.
- [hetzner-server-manager](https://github.com/alfiosalanitri/hetzner-server-manager) — Simple, self-hosted web application to manage multiple Hetzner Cloud servers and projects from a single, clean interface. Built with Python, Flask, and Tailwind CSS, and fully containerized with Docker.
- [iTop Datacollector for Hetzner Cloud](https://github.com/itomig-de/itomig-hetzner-collector) — This stand-alone application collects information from Hetzner projects in order to automatically synchronize the farms and virtual machines in iTop. iTop is a web based open source tool for IT service management tasks.
- [libcluster_hcloud](https://github.com/eightsq/libcluster_hcloud) — This is a Hetzner Cloud clustering strategy for libcluster
- [mcp-hetzner](https://github.com/dkruyt/mcp-hetzner) — A Model Context Protocol (MCP) server for interacting with the Hetzner Cloud API. This server allows language models to manage Hetzner Cloud resources through structured functions.
- [mcp-hetzner-go](https://github.com/MahdadGhasemian/mcp-hetzner-go) — MCP Hetzner Go is a Go-based Model Context Protocol (MCP) server that provides a bridge between AI models and the Hetzner Cloud infrastructure. It enables AI assistants to interact with Hetzner Cloud resources through a standardized protocol, allowing for both read-only and read-write operations across various cloud resources including servers, networks, firewalls, and more.
- [robotlb](https://github.com/intreecom/robotlb) — Cloud Load Balancers for Kubernetes clusters on Robot
- [scalr](https://github.com/ngine-io/scalr) — Autoscaling for Clouds - Scale Cloud instances based on policy checks in a configurable interval. With Hetzner Cloud and Prometheus support.
- [terraform-hcloud-routeros-router](https://github.com/selfscrum/terraform-hcloud-routeros-router) — Terraform code to seamlessly integrate a RouterOS router into a Hetzner Cloud network.
- [terraform-hcloud-routeros-router-configuration](https://github.com/selfscrum/terraform-hcloud-routeros-router-configuration) — Configures a Hetzner server with an RouterOS Router.
- [terraform-kubernetes-hcloud-controller-manager](https://github.com/colinwilson/terraform-kubernetes-hcloud-controller-manager) — A simple module to provision the Hetzner Cloud Controller Manager (With Network & Load Balancer Support) inside a Kubernetes cluster running on Hetzner Cloud. See the variables file for the available configuration options. Please note that this module requires Kubernetes 1.16 or newer.
- [terraform-provider-hetznerdns](https://github.com/germanbrew/terraform-provider-hetznerdns) — A Terraform provider that helps you automate management of DNS zones and records at Hetzner DNS.
- [ui-driver-hetzner](https://github.com/mxschmitt/ui-driver-hetzner) — Rancher UI driver for the Hetzner Cloud docker-machine-driver


## License

[Public Domain (CC0)](https://creativecommons.org/publicdomain/zero/1.0/)
