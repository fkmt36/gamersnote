resource "aws_s3_bucket" "main" {
  bucket = "img.${var.name}.com"
  acl    = "public-read"

  website {
    index_document = "index.html"
    error_document = "error.html"
  }
}

resource "aws_s3_bucket_policy" "main" {
  bucket = aws_s3_bucket.main.id

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "${var.name}-get-object",
      "Action": "s3:GetObject",
      "Principal": "*",
      "Effect": "Allow",
      "Resource": "${aws_s3_bucket.main.arn}/*"
    }
  ]
}
POLICY
}

data "aws_route53_zone" "main" {
  name = var.domain
}

resource "aws_route53_record" "main" {
  zone_id = data.aws_route53_zone.main.id
  name    = "img"
  type    = "A"

  alias {
    name                   = aws_s3_bucket.main.website_domain
    zone_id                = aws_s3_bucket.main.hosted_zone_id
    evaluate_target_health = true
  }
}
