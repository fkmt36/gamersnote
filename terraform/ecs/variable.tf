variable "name" {
  type = string
}
variable "subnets" {
  type = list(string)
}
variable "target_group_arn" {
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
variable "log_group_name" {
  type = string
}
variable "image_tag" {
  type = string
}