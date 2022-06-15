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

* [hetzner-cloud-api-for-net](https://github.com/lk-code/hetzner-cloud-api-net) — Hetzner Cloud API for .NET is a .NET Standard 2.0 Library 

### Go

* **[hcloud-go](https://github.com/hetznercloud/hcloud-go) — Package hcloud is a library for the Hetzner Cloud API.** 

### HCL

* [terraform-kubernetes-hcloud-csi-driver](https://github.com/colinwilson/terraform-kubernetes-hcloud-csi-driver) — A simple module to provision the Hetzner Container Storage Interface Driver within a Kubernetes cluster running on Hetzner Cloud. See the variables file for the available configuration options. Please note that this module requires Kubernetes 1.15 or newer. 

### Java

* [hetznercloud-java](https://github.com/TomSDEVSN/hetznercloud-java) — Simple Java client for the Hetzner Cloud API. 

### JavaScript

* [hcloud-js](https://github.com/dennisbruner/hcloud-js) — A Node.js module for the Hetzner Cloud API 
* [hcloud-nodejs](https://github.com/LeonSpors/hcloud-nodejs) — hcloud-nodejs is a node.js package for Hetzner cloud that can be used to manage your projects. 

### PHP

* [hetzner-cloud-php-sdk](https://github.com/LKDevelopment/hetzner-cloud-php-sdk) — A PHP SDK for the Hetzner Cloud API: docs.hetzner.cloud 

### Python

* **[hcloud-python](https://github.com/hetznercloud/hcloud-python) — hcloud-python is a library for the Hetzner Cloud API.** 

### Ruby

* [hcloud-ruby](https://github.com/tonobo/hcloud-ruby) — Native ruby client for HetznerCloud 

### Rust

* [hcloud-rust](https://github.com/HenningHolmDE/hcloud-rust/) — Unofficial Rust crate for accessing the Hetzner Cloud API 


## Tools

* **[hcloud-cli](https://github.com/hetznercloud/cli) — hcloud is a command-line interface for interacting with Hetzner Cloud.**
* [ansible-role-hcloud](https://github.com/ngine-io/ansible-role-hcloud) — Ansible Role for managing hcloud cloud resources. 
* [chaotic](https://github.com/ngine-io/chaotic) — Fault injection your Hetzner Cloud servers to ensure, your services run even with one server stopped. 
* [cluster-api-provider-hetzner](https://github.com/syself/cluster-api-provider-hetzner) — Kubernetes Cluster API Provider for consistent deployment and day 2 operations of &quot;self-managed&quot; HA Kubernetes clusters on Hetzner. 
* [dda-k8s-crate](https://github.com/DomainDrivenArchitecture/dda-k8s-crate) — dda-k8s-crate installs &amp; configures all in one server k8s on a Hetzner Cloud ubuntu system. 
* [hc-volume-backup](https://gitlab.com/MartinBoehmer/hc-volume-backup) — Bash script to backup Hetzner Cloud Volumes. Automatically creates volumes for backups and maintains a defined number of them. 
* [hcloud-assign-ip](https://github.com/lehuizi/hcloud_assign_ip) — A little helper to easily assign a floating ip to the current system 
* [hcloud-failover-keepalived](https://github.com/lehuizi/hcloud-failover-keepalived) — Script for switching Floating IPs in case of keepalived failover 
* [hcloud-pricing-exporter](https://github.com/jangraefen/hcloud-pricing-exporter) — A Prometheus exporter that connects to your HCloud account and collects data on your current expenses. 
* [hcloud-snapshot-as-backup](https://github.com/fbrettnich/hcloud-snapshot-as-backup) — Hetzner Cloud - Automatic Snapshots as Backups for more flexibility 
* [healthzner-bot](https://github.com/raphaelbernhart/healthznerbot) — A discord bot to get periodically updates of the health status of your hetzner cloud machines. 
* [hetzner-cloud-connect](https://github.com/BlueBambooStudios/hcloud-connect) — Handles automatically adding servers to load balancers 
* [hetzner-kube](https://github.com/xetys/hetzner-kube) — This project contains a CLI tool to easily provision kubernetes clusters on Hetzner Cloud. 
* [hetzner-load-balancer-prometheus-exporter](https://github.com/infraduckture/hetzner-load-balancer-prometheus-exporter) — Exports metrics from Hetzner Load Balancer for consumption by Prometheus 
* [hetzner-rescaler](https://github.com/jonamat/hetzner-rescaler) — Lightweight CLI tool to programmatically rescale your Hetzner Cloud server. 
* [hobby-kube](https://github.com/hobby-kube/guide) — Fully automated cluster setup using Terraform, good balance between resilience and cost, and therefore a great starting point for hobbyists or to build a professional setup with a reasonable foundation. 
* [kops](https://github.com/kubernetes/kops) — Kubernetes Operations (kOps) automates Kubernetes cluster installation, upgrades and management 
* [kubeone](https://github.com/kubermatic/kubeone) — Kubermatic KubeOne automates cluster operations on hetzner cloud. KubeOne can install high-available (HA) master clusters as well single master clusters. 
* [kubernetes-hetzner-keepalived](https://github.com/schemen/kubernetes-hetzner-keepalived) — K8s deployment and image to create a keepalived ip failover with the floating ip feature. 

## Integrations

* **[hcloud-cloud-controller-manager](https://github.com/hetznercloud/hcloud-cloud-controller-manager) — Kubernetes cloud-controller-manager for Hetzner Cloud** 
* **[hcloud-csi](https://github.com/hetznercloud/csi-driver) — Container Storage Interface driver for Hetzner Cloud** 
* **[hetzner-cloud-ansible-collection](https://github.com/ansible-collections/hetzner.hcloud) — Ansible Collection for Hetzner Cloud. Part of Ansible Community Distribution (ACD).** 
* **[hetzner-cloud-packer-builder](https://github.com/hashicorp/packer) — An official Packer.io Builder for Hetzner Cloud** 
* **[kubernetes-cluster-autoscaler](https://github.com/kubernetes/autoscaler) — A k8s component that automatically adjusts the size of a Kubernetes Cluster so that all pods have a place to run and there are no unneeded nodes.** 
* **[prometheus-service-discovery](https://github.com/prometheus/prometheus) — A service discovery build into Prometheus to easily discover Hetzner servers (Cloud &amp; Dedicated)** 
* **[terraform-provider](https://github.com/hetznercloud/terraform-provider-hcloud) — Official Hetzner Cloud Terraform Provider** 
* [algo-vpn](https://github.com/trailofbits/algo) — Algo VPN is a set of Ansible scripts that simplify the setup of a personal Wireguard and IPSEC VPN. 
* [ansible-hcloud-inventory](https://github.com/hg8496/ansible-hcloud-inventory) — An dynamic inventory script for hetzner cloud 
* [docker-machine-driver-hetzner](https://github.com/JonasProgrammer/docker-machine-driver-hetzner) — This library adds the support for creating Docker machines hosted on the Hetzner Cloud. 
* [docker-volume-hetzner](https://github.com/costela/docker-volume-hetzner) — Volume management plugin for Docker (and Swarm) 
* [gitlab-hetzner-runner](https://git.shivering-isles.com/shivering-isles/gitlab-hetzner-runner) — A version of the gitlab-runner base image, that allows the usage of Hetzner Cloud instances for GitLab CI. 
* [hcloud-cloud-controller-manager-helm-chart](https://gitlab.com/MatthiasLohr/hcloud-cloud-controller-manager-helm-chart) — Hetzner Cloud - Cloud Controller Manager Helm Chart 
* [hcloud-csi-driver-helm-chart](https://gitlab.com/MatthiasLohr/hcloud-csi-driver-helm-chart) — Community Helm Chart for Hetzner Cloud CSI Driver for Kubernetes 
* [hcloud-fip-controller](https://github.com/cbeneke/hcloud-fip-controller) — Kubernetes controller to (re-)assign floating IPs to Hetzner Cloud instances. 
* [hcloud-freebsd](https://github.com/paulc/hcloud-freebsd) — Hetzner Cloud auto-provisioning for FreeBSD 
* [hcloud-ip-floater](https://github.com/costela/hcloud-ip-floater) — Kubernetes controller to attach hcloud floating IPs to services&#039; nodes. 
* [hcloud-tg](https://github.com/Navid2zp/hcloud-tg) — Telegram bot for managing Hetzner cloud servers. 
* [hetzner-cloud-deploy-server-github-action](https://github.com/TimDaub/hetzner-cloud-deploy-server-action) — Deploy a Hetzner Cloud Server from a GitHub Action. 
* [hetzner-cloud-plugin-for-jenkins](https://github.com/jenkinsci/hetzner-cloud-plugin) — The Hetzner cloud plugin enables Jenkins CI to schedule builds on dynamically provisioned VMs in Hetzner Cloud. 
* [itop-datacollector-for-hetzner-cloud](https://github.com/itomig-de/itomig-hetzner-collector) — This stand-alone application collects information from Hetzner projects in order to automatically synchronize the farms and virtual machines in iTop. iTop is a web based open source tool for IT service management tasks. 
* [laravel-forge](https://forge.laravel.com/) — Laravel Hosting &amp; Instant PHP Servers on Hetzner Cloud 
* [libcluster-hcloud](https://github.com/EightSQ/libcluster_hcloud) — This is a Hetzner Cloud clustering strategy for libcluster 
* [ploi](https://ploi.io) — Rapidly deploy any site you like on Hetzner Cloud servers. 
* [rancher-on-hetzner-cloud](https://github.com/alexzimmer96/rancher-hcloud) — Prebuild Terraform templates for deploying a highly available RKE cluster on Hetzner Cloud and installing Rancher into it 
* [scalechamp](https://www.scalechamp.com) — Managed databases provider with PostgreSQL, MySQL, KeyDB, Redis support across 11 public clouds including Hetzner Cloud. 
* [scalr](https://github.com/ngine-io/scalr) — Autoscaling for Clouds - Scale Cloud instances based on policy checks in a configurable interval. With Hetzner Cloud and Prometheus support. 
* [terraform-hcloud-routeros-router](https://github.com/selfscrum/terraform-hcloud-routeros-router) — Terraform code to seamlessly integrate a RouterOS router into a Hetzner Cloud network. 
* [terraform-hcloud-routeros-router-configuration](https://github.com/selfscrum/terraform-hcloud-routeros-router-configuration) — Configures a Hetzner server with an RouterOS Router. 
* [terraform-kubernetes-hcloud-controller-manager](https://github.com/colinwilson/terraform-kubernetes-hcloud-controller-manager) — A simple module to provision the Hetzner Cloud Controller Manager (With Network &amp; Load Balancer Support) inside a Kubernetes cluster running on Hetzner Cloud. See the variables file for the available configuration options. Please note that this module requires Kubernetes 1.16 or newer. 
* [ui-driver-hetzner](https://github.com/mxschmitt/ui-driver-hetzner) — Rancher UI driver for the Hetzner Cloud docker-machine-driver 
* [versio-io](https://www.versio.io/import-hetzner-cloud-cmdb-configuration-item.html) — Integration of Hetzner Cloud configuration items (CI) in the full stack Versio.io configuration management database (CMDB). 

## License

[Public Domain (CC0)](https://creativecommons.org/publicdomain/zero/1.0/)
