resource "aws_ecr_repository" "api" {
  name                 = "${var.name}-api"
}

resource "aws_ecr_repository" "app" {
  name                 = "${var.name}-app"
}