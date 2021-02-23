# ALB用のSecurityGroup
resource "aws_security_group" "alb" {
  name        = "${var.name}-alb"
  description = "${var.name} alb"
  vpc_id      = var.vpc_id

  ingress {
    from_port = 443
    to_port   = 443
    protocol  = "tcp"
    # cidr_blocks = ["0.0.0.0/0"]
    cidr_blocks = ["133.200.211.64/32"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.name}-alb"
  }
}

resource "aws_security_group_rule" "http" {
  type      = "ingress"
  from_port = 80
  to_port   = 80
  protocol  = "tcp"
  # cidr_blocks = ["0.0.0.0/0"]
  cidr_blocks       = ["133.200.211.64/32"]
  security_group_id = aws_security_group.alb.id
}

# ドメインの紐付け
data "aws_route53_zone" "this" {
  name         = var.domain
  private_zone = false
}
resource "aws_route53_record" "this" {
  type = "A"

  name    = var.domain
  zone_id = data.aws_route53_zone.this.id

  alias {
    name                   = module.aws_alb.this_lb_dns_name
    zone_id                = module.aws_alb.this_lb_zone_id
    evaluate_target_health = true
  }
}

# ALB
# https://github.com/terraform-aws-modules/terraform-aws-alb
module "aws_alb" {
  source  = "terraform-aws-modules/alb/aws"
  version = "~> 5.0"

  name = "${var.name}-alb"

  load_balancer_type = "application"

  vpc_id          = var.vpc_id
  subnets         = var.public_subnets
  security_groups = [aws_security_group.alb.id]

  # access_logs = {
  #   bucket = "my-alb-logs"
  # }

  target_groups = [
    {
      name_prefix      = "gn-"
      backend_protocol = "HTTP"
      backend_port     = 80
      target_type      = "ip"
      health_check = {
        path = "/api/v1/status"
      }
    },
    {
      name_prefix      = "gn-"
      backend_protocol = "HTTP"
      backend_port     = 80
      target_type      = "ip"
      health_check = {
        path = "/"
      }
    }
  ]

  https_listeners = [
    {
      port               = 443
      certificate_arn    = var.certificate_arn
      target_group_index = 1
    }
  ]

  https_listener_rules = [
    {
      https_listener_index = 0

      actions = [{
        type = "redirect"
        redirect = {
          host        = "gamersnote.com"
          status_code = "HTTP_301"
        }
        },
        {
          type               = "forward"
          target_group_index = 0
      }]

      conditions = [{
        host_header = ["54.64.220.254"]
        },
        {
          path_patterns = ["/api/*"]
      }]
    }
  ]

  http_tcp_listeners = [
    {
      port        = 80
      protocol    = "HTTP"
      action_type = "redirect"
      redirect = {
        port        = "443"
        protocol    = "HTTPS"
        status_code = "HTTP_301"
      }
    }
  ]
}
