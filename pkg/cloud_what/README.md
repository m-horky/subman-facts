# `cloud-what`

The goal of `cloud-what` is to identify cloud providers: AWS, GCP or Azure.


## Code structure

Each cloud is represented by an instance of `Cloud` struct, which roughly defines it.

- `aws_*.go` detects the Amazon Web Services cloud,
- `gcp_*.go` detects the Google Cloud Provider,
- `azure_*.go` detects the Microsoft Azure cloud,
- `any_*.go` provides shared functions used by all clouds.

Additionally,

- `*_detect.go` detects whether the code is running on the cloud provider or not, exposed via `SystemRunsOn()` and `SystemMaybeRunsOn()`,
- `*_gather.go` talks to metadata servers on the public clouds and gathers information about the systems,
- `*_collect.go` are reporting outputs further ingestible by other tooling.


## Tasks of `cloud-what`

### Cloud detection

Each cloud provider has two methods providing detection: strong and heuristic.

The strong cloud detection uses only strong sights and reports its findings as true or false.
Examples of strong detection are BIOS version or vendor or the host type as reported by `virt-what`.
Examples of heuristic (or weak) detection may also be UUIDs or various DMI values.


### Self identification

`cloud-what`, after finding out which cloud provider it is running under, will contact an IMDS (Instance Metadata Service) server running on the cloud provider.
It is used to request a session token, which is used as an identifier to the metadata server.

This workflow may be a bit different between
[AWS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/instance-metadata-v2-how-it-works.html),
[GCP](https://cloud.google.com/compute/docs/metadata/querying-metadata) and
[Azure](https://learn.microsoft.com/en-us/azure/virtual-machines/instance-metadata-service?tabs=linux).


### Fact provider

Once the system collected the metadata about itself (such as account ID or instance ID), it can provide them as facts.
They are used by ~~`subscription-manager`~~ `rhc.next`.
