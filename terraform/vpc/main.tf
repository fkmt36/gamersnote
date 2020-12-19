# VPC
# https://github.com/terraform-aws-modules/terraform-aws-vpc
module "aws_vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = var.name

  cidr = "10.0.0.0/16"

  azs              = ["ap-northeast-1a", "ap-northeast-1c"]
  private_subnets  = ["10.0.1.0/24", "10.0.2.0/24", ]
  public_subnets   = ["10.0.10.0/24", "10.0.20.0/24"]
  database_subnets = ["10.0.100.0/24", "10.0.200.0/24"]

  create_database_subnet_group           = true
  create_database_subnet_route_table     = true
  create_database_internet_gateway_route = true

  enable_nat_gateway = true
  single_nat_gateway = true

  enable_dns_hostnames = true
  enable_dns_support   = true
}
