# PBS TF SQS Module

## Installation

### Using the Repo Source

Use this URL for the source of the module. See the usage examples below for more details.

```hcl
github.com/pbs/terraform-aws-sqs-module?ref=x.y.z
```

### Alternative Installation Methods

More information can be found on these install methods and more in [the documentation here](./docs/general/install).

## Usage

Provisions an SQS queue.

Integrate this module like so:

```hcl
module "queue" {
  source = "github.com/pbs/terraform-aws-sqs-module?ref=x.y.z"

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}
```

If you need to integrate a secondary queue as a dead letter queue, this would be a valid configuration:

```hcl
module "queue" {
  source = "github.com/pbs/terraform-aws-sqs-module?ref=x.y.z"

  name = "my-queue"

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}

module "dlq" {
  source = "github.com/pbs/terraform-aws-sqs-module?ref=x.y.z"

  name = "my-queue-dlq"

  redrive_policy = jsonencode({
    deadLetterTargetArn = module.queue.arn
    maxReceiveCount     = 5
  })

  organization = var.organization
  environment  = var.environment
  product      = var.product
  repo         = var.repo
}
```

## Adding This Version of the Module

If this repo is added as a subtree, then the version of the module should be close to the version shown here:

`x.y.z`

Note, however that subtrees can be altered as desired within repositories.

Further documentation on usage can be found [here](./docs).

Below is automatically generated documentation on this Terraform module using [terraform-docs][terraform-docs]

---

[terraform-docs]: https://github.com/terraform-docs/terraform-docs
