output "log_group_name_api" {
  value = aws_cloudwatch_log_group.api.name
}

output "log_group_name_app" {
  value = aws_cloudwatch_log_group.app.name
}