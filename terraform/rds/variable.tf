variable "name" {
  type = string
}

variable "db_username" {
  type = string
}

variable "db_password" {
  type = string
}

variable "public_subnets" {
  type = list(string)
}
