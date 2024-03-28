terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = var.region
}

module "vpc" {
  source = "../../"

  vpc_name = var.vpc_name

  create_nat_gateway     = true
  create_private_subnets = true
  create_public_subnets  = false
}
