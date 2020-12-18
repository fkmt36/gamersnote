variable "name" {}
variable "domain" {}
variable "aws_region" {}
variable "aws_access_key" {}
variable "aws_secret_key" {}

provider "aws" {
  region     = var.aws_region
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}

# ECR
module "ecr" {
  source = "./ecr"
  name   = var.name
}

# https化
module "acm" {
  source = "./acm"
  domain = var.domain
}

# VPC
module "vpc" {
  source = "./vpc"
  name   = var.name
}

# ALB
module "alb" {
  source          = "./alb"
  name            = var.name
  vpc_id          = module.vpc.vpc_id
  public_subnets  = module.vpc.public_subnets
  domain          = var.domain
  certificate_arn = module.acm.certificate_arn
}

# ECS
module "ecs" {
  source           = "./ecs"
  name             = var.name
  subnets          = [module.vpc.private_subnets[0]] # 節約のために一つのサブネットのみ使う
  target_group_arn = module.alb.target_group_arns[0]
  vpc_id           = module.vpc.vpc_id
}
