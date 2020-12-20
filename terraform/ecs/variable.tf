variable "name" {
  type = string
}
variable "subnets" {
  type = list(string)
}
variable "target_group_arn_api" {
  type = string
}
variable "target_group_arn_app" {
  type = string
}
variable "vpc_id" {
  type = string
}
variable "aws_account_id" {
  type = string
}
variable "aws_region" {
  type = string
}
variable "log_group_name_api" {
  type = string
}
variable "log_group_name_app" {
  type = string
}
variable "image_tag" {
  type = string
}
