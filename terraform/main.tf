variable "name" {}
variable "domain" {}
variable "organization" {}

provider "aws" {}

terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }

  backend "remote" {
    organization = var.organization

    workspaces {
      name = var.name
    }
  }
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
