# Task用のSecurityGroup
resource "aws_security_group" "this" {
  name        = "${var.name}-ecs"
  description = "${var.name} ecs"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["10.0.0.0/16"] # 同一IPのみ
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.name}-ecs"
  }
}

# タスク用のRole
resource "aws_iam_role" "ecs_task_role" {
  name = "ecs_task_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "ecs-tasks.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
EOF
}

# タスク用のポリシー
resource "aws_iam_policy" "ecs_task_policy" {
  name = "ecs_task_policy"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
          "ssm:GetParameters"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

# ロールにポリシーをアタッチ
resource "aws_iam_policy_attachment" "attach" {
  name       = "attach"
  roles      = [aws_iam_role.ecs_task_role.name]
  policy_arn = aws_iam_policy.ecs_task_policy.arn
}


# ECS Cluster
resource "aws_ecs_cluster" "main" {
  name = var.name
}

# Task Definition
resource "aws_ecs_task_definition" "api" {
  family                   = "${var.name}-api"
  task_role_arn            = aws_iam_role.ecs_task_role.arn
  network_mode             = "awsvpc"
  cpu                      = 256
  memory                   = 512
  requires_compatibilities = ["FARGATE"]
  container_definitions    = file("${path.module}/api-container-definition.json")
}

# ECS Service
resource "aws_ecs_service" "api" {
  name = "${var.name}-api"

  cluster         = aws_ecs_cluster.main.name
  launch_type     = "FARGATE"
  desired_count   = "1"
  task_definition = aws_ecs_task_definition.api.arn

  network_configuration {
    subnets         = var.subnets
    security_groups = [aws_security_group.this.id]
  }

  load_balancer {
    target_group_arn = var.target_group_arn
    container_name   = "api"
    container_port   = "80"
  }
}
