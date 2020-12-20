variable "name" {}
variable "domain" {}
variable "db_username" {}
variable "db_password" {}
variable "aws_account_id" {}
variable "aws_region" {}
variable "image_tag" {}

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
  source                     = "./rds"
  name                       = var.name
  db_username                = var.db_username
  db_password                = var.db_password
  database_subnet_group_name = module.vpc.database_subnet_group
  vpc_id                     = module.vpc.vpc_id
}

# S3
module "s3" {
  source = "./s3"
  name   = var.name
  domain = var.domain
}

# Cloud Watch
module "cw" {
  source = "./cw"
  name   = var.name
}

# ECS
module "ecs" {
  source           = "./ecs"
  name             = var.name
  subnets          = [module.vpc.private_subnets[0]] # 節約のために一つのサブネットのみ使う
  target_group_arn = module.alb.target_group_arns[0]
  vpc_id           = module.vpc.vpc_id
  aws_account_id   = var.aws_account_id
  aws_region       = var.aws_region
  log_group_name   = module.cw.log_group_name_api
  image_tag        = var.image_tag
}

