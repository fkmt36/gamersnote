output "vpc_id" {
  value = module.aws_vpc.vpc_id
}

output "public_subnets" {
  value = module.aws_vpc.public_subnets
}

output "private_subnets" {
  value = module.aws_vpc.private_subnets
}

output "database_subnets" {
  value = module.aws_vpc.database_subnets
}