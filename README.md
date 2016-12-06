[![Build Status](https://travis-ci.org/brandoncole/synthetic.svg?branch=master)](https://travis-ci.org/brandoncole/synthetic)
[![](https://godoc.org/github.com/brandoncole/synthetic?status.svg)](http://godoc.org/github.com/brandoncole/synthetic)
[![Apache 2.0](https://img.shields.io/hexpm/l/plug.svg)](LICENSE)

# Background

On multiple occasions as an operator building distributed cloud platforms or tenant of a distributed system I've encountered the need for the following:

* Predictably utilize finite resources of the system (CPU, Memory, Network, Disk, Ports)
* Push at the edges of a cloud orchestration system (i.e. VirtualBox, Docker, Kubernetes, Mesos, ECS, etc.)
* Have a canary within a cloud orchestration system that predictably when resource requests cannot be fulfilled
* Quantitatively benchmark performance characteristics of multiple cloud systems for comparison
* Generate predictable log streams that can be used to verify the capture, rotation and retention of distributed logs

These simple workflows allow cloud operators and tenants to answer some very important questions:

* Cloud Operators
    * What happens if a tenant consumes all of a finite resource?
    * Are the predefined resource limitations for tenants working (i.e. cgroups, Kubernetes Limits)?
    * Will the system automatically scale up or down to accommodate tenants?
    * Is the monitoring accurately reflecting what's happening on the systems?
    * Are alarms properly configured to detect critical scenarios?
    * Is the system fulfilling its SLA for tenants?
* Cloud Tenants
    * How much CPU / Memory / Network / Disk bandwidth is available?
    * What behavior should be expected if too much bandwidth is consumed?

# Current Work

This project is just starting and is in its infancy, so it is not recommended that others use it now.  The roadmap shown is there to give the audience an idea of what features are expected to come online shortly at which point it will be considered ready for general use.

![CPU Simulation](docs/cpu-fast.gif)

# Roadmap

1. Technical Debt - the codebase is mostly a proof of concept right now.  A lot of this code needs to be cleaned up.
2. Dockerize - provide a `Dockerfile` and publicly available versioned image for others to use.
3. Orchestration - provide configurations to deploy it on common orchestration systems like Kubernetes.
4. Builds - configure an automated build for the project via a tool like Travis CI
5. Completeness - flush out the implementation of the network, disk, memory and stdout tests
6. Refine CLI - refine and document the command line interface and provide more examples

# CLI

### $ synthetic load

```
$ synthetic load --help
Simulates process load

Usage:
  synthetic load [flags]

Examples:

# Runs a synthetic CPU load that utilizes 50% of the CPU
synthetic load -c -p flat --profilemax 50

# Runs a synthetic CPU load that utilizes between 0% and 50% of the CPU over 30 seconds
synthetic load -c -p sine --profilemin 0 --profilemax 50 --profileperiod 30s

Flags:
  -c, --cpu                      Enables a synthetic CPU load
  -d, --disk                     Enables a synthetic disk load
      --duration duration        Amount of time to run the load for, or infinite if 0s
  -m, --memory                   Enables a synthetic memory load
  -n, --network                  Enables a synthetic network load
  -p, --profile string           Specifies the load profile [flat, sine] (default "flat")
      --profilemax int           Maximum load as a percentage of available. (default 50)
      --profilemin int           Minimum load as a percentage of available. (default 50)
      --profileperiod duration   Period duration for sine profile in seconds. (default 1m0s)
```

# Design Documentation

Software documentation for the project is available at the link below courtesy of http://godoc.org

* https://godoc.org/github.com/brandoncole/synthetic