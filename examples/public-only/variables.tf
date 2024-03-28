variable "vpc_name" {
  description = "The name of the VPC."
  default     = "vpc-public-only-example"
  type        = string
}

variable "region" {
  description = "The AWS region to deploy resources."
  default     = "us-east-1"
  type        = string
}
