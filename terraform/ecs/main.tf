# Task用のSecurityGroup
resource "aws_security_group" "this" {
  name        = "${var.name}-ecs"
  description = "${var.name} ecs"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 80
    to_port     = 3000
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

# タスク実行用のRole
resource "aws_iam_role" "ecs_task_ex_role" {
  name = "ecs_task_ex_role"

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

# タスク実行用のポリシー
resource "aws_iam_policy" "ecs_task_ex_policy" {
  name = "ecs_task_ex_policy"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ssm:GetParameters",
        "ecr:GetAuthorizationToken",
        "ecr:BatchCheckLayerAvailability",
        "ecr:GetDownloadUrlForLayer",
        "ecr:BatchGetImage",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "*"
    }
  ]
}
EOF
}

# ロールにポリシーをアタッチ
resource "aws_iam_policy_attachment" "ecs_task_policy_attach" {
  name       = "ecs_task_policy_attach"
  roles      = [aws_iam_role.ecs_task_role.name]
  policy_arn = aws_iam_policy.ecs_task_policy.arn
}
resource "aws_iam_policy_attachment" "ecs_task_ex_policy_attach" {
  name       = "ecs_task_ex_policy_attach"
  roles      = [aws_iam_role.ecs_task_ex_role.name]
  policy_arn = aws_iam_policy.ecs_task_ex_policy.arn
}

# ECS Cluster
resource "aws_ecs_cluster" "main" {
  name = var.name
}

# Task Definition
resource "aws_ecs_task_definition" "api" {
  family                   = "${var.name}-api"
  task_role_arn            = aws_iam_role.ecs_task_role.arn
  execution_role_arn       = aws_iam_role.ecs_task_ex_role.arn
  network_mode             = "awsvpc"
  cpu                      = 256
  memory                   = 512
  requires_compatibilities = ["FARGATE"]
  container_definitions    = <<TASK_DEFINITION
[
  {
    "name": "api",
    "image": "${var.aws_account_id}.dkr.ecr.${var.aws_region}.amazonaws.com/${var.name}-api",
    "portMappings": [
      {
          "containerPort": 3000,
          "hostPort": 3000
      }
    ],
    "logConfiguration": {
      "logDriver": "awslogs",
      "options": {
        "awslogs-region": "${var.aws_region}",
        "awslogs-group": "${var.log_group_name}",
        "awslogs-stream-prefix": "${var.name}-api"
      }
    }
    "secrets": [
      {
        "valueFrom": "gamersnote-postgres-host",
        "name": "POSTGRES_HOST"
      },
      {
        "valueFrom": "gamersnote-postgres-user",
        "name": "POSTGRES_USER"
      },
      {
        "valueFrom": "gamersnote-postgres-password",
        "name": "POSTGRES_PASSWORD"
      },
      {
        "valueFrom": "gamersnote-postgres-dbname",
        "name": "POSTGRES_DBNAME"
      },
      {
        "valueFrom": "gamersnote-postgres-port",
        "name": "POSTGRES_PORT"
      },
      {
        "valueFrom": "gamersnote-email-from-address",
        "name": "EMAIL_FROM_ADDRESS"
      },
      {
        "valueFrom": "gamersnote-aws-access-key-id",
        "name": "AWS_ACCESS_KEY_ID"
      },
      {
        "valueFrom": "gamersnote-aws-secret-key",
        "name": "AWS_SECRET_KEY"
      },
      {
        "valueFrom": "gamersnote-aws-region",
        "name": "AWS_REGION"
      },
      {
        "valueFrom": "gamersnote-base-url",
        "name": "BASE_URL"
      },
      {
        "valueFrom": "gamersnote-s3-bucket",
        "name": "S3_BUCKET"
      },
      {
        "valueFrom": "gamersnote-s3-base-url",
        "name": "S3_BASEURL"
      }
    ]
  }
]
TASK_DEFINITION
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
    container_port   = "3000"
  }
}
