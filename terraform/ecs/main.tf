# Task用のSecurityGroup
resource "aws_security_group" "this" {
  name = "${var.name}-ecs"
  description = "${var.name} ecs"
  vpc_id = var.vpc_id

  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = ["10.0.0.0/16"] # 同一IPのみ
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.name}-ecs"
  }
}

# ECS Cluster
resource "aws_ecs_cluster" "main" {
  name = var.name
}

# # Task Definition
# resource "aws_ecs_task_definition" "main" {
#   family = "${var.name}"

#   requires_compatibilities = ["FARGATE"]

#   cpu    = "256"
#   memory = "512"

#   network_mode = "awsvpc"

#   container_definitions = <<EOL
# [
#   {
#     "name": "nginx",
#     "image": "nginx:1.14",
#     "portMappings": [
#       {
#         "containerPort": 80,
#         "hostPort": 80
#       }
#     ]
#   }
# ]
# EOL
# }

# # ECS Service
# resource "aws_ecs_service" "main" {
#   name = "${var.name}"

#   cluster = aws_ecs_cluster.main.name
#   launch_type = "FARGATE"
#   desired_count = "1"
#   task_definition = aws_ecs_task_definition.main.arn

#   network_configuration {
#     subnets         = var.subnets
#     security_groups = [aws_security_group.this.id]
#   }

#   load_balancer {
#     target_group_arn = var.target_group_arn
#     container_name   = "nginx"
#     container_port   = "80"
#   }
# }
