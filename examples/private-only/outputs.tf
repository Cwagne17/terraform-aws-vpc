output "availability_zones" {
  description = "The availability zones of the VPC."
  value       = module.vpc.availability_zones
}

output "nat_gateway_public_ip" {
  description = "The public IP address of the NAT gateways."
  value       = module.vpc.nat_gateway_public_ip
}

output "num_availability_zones" {
  description = "The number of availability zones of the VPC."
  value       = module.vpc.num_availability_zones
}

output "num_nat_gateways" {
  description = "The number of NAT gateways created."
  value       = module.vpc.num_nat_gateways
}

output "private_subnet_cidr_blocks" {
  description = "The CIDR blocks of the private subnets."
  value       = module.vpc.private_subnet_cidr_blocks
}

output "private_subnet_ids" {
  description = "The IDs of the private subnets."
  value       = module.vpc.private_subnet_ids
}

output "private_subnet_route_table_id" {
  description = "The ID of the private subnet route table."
  value       = module.vpc.private_subnet_route_table_id
}

output "private_subnets" {
  description = "A map of all private subnets, with the subnet name as key, and all aws-subnet properties as the value."
  value       = module.vpc.private_subnets
}

output "public_subnet_cidr_blocks" {
  description = "The CIDR blocks of the public subnets."
  value       = module.vpc.public_subnet_cidr_blocks
}

output "public_subnet_ids" {
  description = "The IDs of the public subnets."
  value       = module.vpc.public_subnet_ids
}

output "public_subnet_route_table_id" {
  description = "The ID of the public subnet route table."
  value       = module.vpc.public_subnet_route_table_id
}

output "public_subnets" {
  description = "A map of all public subnets, with the subnet name as key, and all aws-subnet properties as the value."
  value       = module.vpc.public_subnets
}

output "public_subnets_network_acl_id" {
  description = "The ID of the public subnet network ACL."
  value       = module.vpc.public_subnets_network_acl_id
}

output "vpc_cidr_block" {
  description = "The CIDR block of the VPC."
  value       = module.vpc.vpc_cidr_block
}

output "vpc_id" {
  description = "The ID of the VPC."
  value       = module.vpc.vpc_id
}

output "vpc_name" {
  description = "The name of the VPC."
  value       = module.vpc.vpc_name
}
