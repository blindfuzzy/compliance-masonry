---
permalink: /NIST-800-53/CM-6/
title: CM-6 - Configuration Settings
parent: NIST-800-53 Documentation
---

## Amazon Machine Images
a. - DevOps and Security Engineers maintain the baseline configuration for VPC, EBS and AMIs.  Best practices, FISMA compliant AMIs, and hardened cloud formation templates are utilized as there are no benchmarks available.
- The organization uses FISMA compliant and hardened AMIs within its AWS infrastructure
 
 
### References

* [Amazon Machine Images](http://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AMIs.html)

--------

## Amazon Elastic Block Store
a. - DevOps and Security Engineers maintain the baseline configuration for VPC, EBS and AMIs.  Best practices, FISMA compliant AMIs are utilized as there are no benchmarks available.
 
 
### References

* [Amazon Elastic Block Store](https://aws.amazon.com/ebs/)

--------

## Nessus
d. Nessus and AlienVault USM Joval scans are performed at least on a quarterly basis in the event that no enhancements or upgrades are performed. Both tools meet NIST’s SCAP 1.2 requirements, satisfying OMB Mandate M-08-22 and associated procurement requirements. SCAP scans are performed weekly and monthly to ensure no unauthorized changes, enhancements or upgrades are performed. 
 
### References

* [Nessus Website](http://www.tenable.com/products/nessus-vulnerability-scanner)

--------

## S3
a. Updates to new BOSH stemcells are located and stored within Amazon S3 http://boshartifacts.cloudfoundry.org/file_collections?type=stemcells 
 
### References

* [Amazon S3](https://aws.amazon.com/s3/)

--------

## BOSH Stemcells
BOSH Stemcells are used for the standard baseline OS images and software vulnerability management updates. Updates to new BOSH stemcells are located and stored within Amazon S3. The specifications of the current release of BOSH stemcells are located on GitHub. DevOps implements Cloud Foundry standard BOSH stemcells for baseline OS configuration.
### References

* [Bosh StemCells](https://bosh.io/stemcells)

* [Bosh updates](https://github.com/cloudfoundry/bosh/blob/master/bosh-stemcell/OS_IMAGES.md)

* [New BOSH stemcells](http://boshartifacts.cloudfoundry.org/file_collections?type=stemcells)

* [Bosh Ubuntu trusty specs](https://github.com/cloudfoundry/bosh/blob/master/bosh-stemcell/spec/stemcells/ubuntu_trusty_spec.rb)

--------

## Manifests
Cloud Foundry configuration settings are documented within the deployment manifest on the 18F GitHub and Cloud Foundry sites. DevOps implements manifest templates written in yml to automate deployment of multiple applications at once and the platform within AWS with consistency and reproducibility.
### References

* [Deploying with Application Manifests](https://docs.cloudfoundry.org/devguide/deploy-apps/manifest.html)

--------