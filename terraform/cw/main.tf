resource "aws_cloudwatch_log_group" "api" {
  name = "${var.name}-api"
}

resource "aws_cloudwatch_log_group" "app" {
  name = "${var.name}-app"
}