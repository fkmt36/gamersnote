# RDS用のSecurityGroup
resource "aws_security_group" "rds" {
  name        = "${var.name}-rds"
  description = "${var.name} rds"
  vpc_id      = var.vpc_id

  ingress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 5432
    to_port     = 5432
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.name}-rds"
  }
}

resource "aws_db_instance" "main" {
  identifier                = "${var.name}-db"
  allocated_storage         = 20
  storage_type              = "gp2"
  engine                    = "postgres"
  engine_version            = "12.5"
  instance_class            = "db.t2.micro"
  name                      = var.name
  username                  = var.db_username
  password                  = var.db_password
  parameter_group_name      = "default.postgres12"
  db_subnet_group_name      = var.database_subnet_group_name
  skip_final_snapshot       = false
  final_snapshot_identifier = "${var.name}-db-snapshot"
  publicly_accessible       = true
  vpc_security_group_ids    = [aws_security_group.rds.id]
}
