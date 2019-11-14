---
layout: "oci"
page_title: "Tagging Resources"
sidebar_current: "docs-oci-guide-tagging_resources"
description: |-
  The Oracle Cloud Infrastructure provider. Tagging
---
## Tagging OCI Resources
When you have many resources (for example, instances, VCNs, load balancers, and block volumes) across multiple compartments in your tenancy, it can become difficult to track resources used for specific purposes, or to aggregate them, report on them, or take bulk actions on them. Tagging allows you to define keys and values and associate them with resources. You can then use the tags to help you organize and list resources based on your business needs. See [Tagging Overview](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/taggingoverview.htm#overview) to familiarize yourself with concept of tagging and features available.

## Managing Tags and Tag Namespaces
* See [tag_namespaces](../r/identity_tag_namespace.html) for guidance on managing lifecycle of tag namespaces.
* See [tags](../d/identity_tag_namespaces.html) for guidance on managing lifecycle of tags.

## How To Manage Tags on OCI Resources
* **Freeform tags:** Freeform tags are simple key value map
* **Defined tags:** Defined tags provide a key/value map and are organized by combining the tag namespaces with tag keys using dot notation. For example, a tag namespace called `HumanResources` could have a key named `CostCenter`. You then associate the namespace and key `HumanResource.CostCenter` and then assign the desired tag, as shown in the following example.

### Propagation of tagging on resources

If your Terraform configuration has a primary resource (e.g. compute instance) with a nested secondary resource(e.g. VNIC for the compute instance) where both these resources support **Freeform tags** or **Defined tags**, OCI Services will propagate all the primary resource's **Freeform tags** and **Defined tags** to secondary resource. This could cause a drift in the Terraform state resulting in a diff after apply. To avoid this, explicitly add all the primary resource's **Freeform tags** and **Defined tags** on the secondary resources as part of the configuration.
The same behavior could be seen while using the **Tag Default** or **Required Tags** feature. This can also be avoided by applying the **Tag Default** or **Required Tags** on all resources(primary and secondary if any) in the tenancy where **Tag Default** or **Required Tags** exist.
   

### Examples:
#### Example 1:
```hcl
//Hand curated way
resource "oci_core_instance" "t" {
    .
    .
    .
    
    freeform_tags =  {
           Environment = "Prod"
           Department = "Ops"
   }
    defined_tags = {
        HumanResources.CostCenter = "42"
        Operations.Project = "Beta"
        HumanResources.Environment = "Production"
    }
}
```
#### Example 2:
```hcl
//Using Locals(available in terraform 0.10.3 or later) & interpolation

locals {
  //Put all common tags here
  common_tags = "${map(
                        "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag1.name}","value",
                        "HumanResources.Environment", "Production",
                        "Operations.Project", "Beta"
                        )}"
}

resource "oci_core_instance" t {
    .
    .
    .
  
    freeform_tags = "${map("key${count.index}", "value${count.index}",
        "domain", "${lookup(data.oci_identity_availability_domains.ADs.availability_domains[var.AD - 1],"name")}"
        )}"
  
    defined_tags = "${merge(
                        local.common_tags,
                        map(
                            "${oci_identity_tag_namespace.tag_namespace1.name}.${oci_identity_tag.tag2.name}", "awesome-tag-example"
                        )
                    )}"
  
}
```
