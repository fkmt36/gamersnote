variable "name" {
  type = string
}

variable "vpc_id" {
  type = string
}

variable "public_subnets" {
  type = list(string)
}

variable "domain" {
  type = string
}

variable "certificate_arn" {
  type = string
}
