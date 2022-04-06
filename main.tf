
resource "aws_s3_object" "files" {
    for_each = var.files
    bucket = var.domain_name_bucket
    key = each.value.name
    source = each.value.path
    acl = "public-read"
    content_type = "text/html"
}
