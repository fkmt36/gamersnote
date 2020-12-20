resource "aws_cloudwatch_log_group" "api" {
  name = "${var.name}-api"
}