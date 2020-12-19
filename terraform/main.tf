variable "name" {}
variable "domain" {}
variable "db_username" {}
variable "db_password" {}

provider "aws" {}

terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
    }
  }

  backend "remote" {
    organization = "happyfukumoto"

    workspaces {
      name = "gamersnote"
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

# RDS
module "rds" {
  source          = "./rds"
  name            = var.name
  db_username     = var.db_username
  db_password     = var.db_password
  private_subnets = module.vpc.private_subnets
}

# S3
module "s3" {
  source = "./s3"
  name   = var.name
}

# ECS
module "ecs" {
  source           = "./ecs"
  name             = var.name
  subnets          = [module.vpc.private_subnets[0]] # 節約のために一つのサブネットのみ使う
  target_group_arn = module.alb.target_group_arns[0]
  vpc_id           = module.vpc.vpc_id
}

