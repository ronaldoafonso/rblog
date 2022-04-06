
variable "region" {
    description = "AWS region"
    type = string
    default = "us-east-2"
}

variable "domain_name_bucket" {
    description = "AWS S3 bucket name for the static data web site."
    type = string
}

variable "files" {
    description = "Static HTML files."
    type = map
}
