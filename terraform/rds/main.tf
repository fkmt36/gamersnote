resource "aws_db_subnet_group" "main" {
  name       = "gamersnote-subnet-group"
  subnet_ids = var.public_subnets
}

resource "aws_db_instance" "main" {
  identifier                = "${var.name}-db"
  allocated_storage         = 20
  storage_type              = "gp2"
  engine                    = "postgres"
  engine_version            = "12.4"
  instance_class            = "db.t2.micro"
  name                      = var.name
  username                  = var.db_username
  password                  = var.db_password
  parameter_group_name      = "default.postgres12"
  db_subnet_group_name      = aws_db_subnet_group.main.name
  skip_final_snapshot       = false
  final_snapshot_identifier = "${var.name}-db-snapshot"
  publicly_accessible       = true
}
