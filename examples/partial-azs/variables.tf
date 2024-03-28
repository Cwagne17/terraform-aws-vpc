variable "vpc_name" {
  description = "The name of the VPC."
  default     = "vpc-partial-azs-example"
  type        = string
}

variable "region" {
  description = "The AWS region to deploy resources."
  default     = "us-east-1"
  type        = string
}
