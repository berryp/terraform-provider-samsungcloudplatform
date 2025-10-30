resource "samsungcloudplatform_certificate" "my_certificate" {
  name = var.name
  key = "{encoded-data}"
  body = "{encoded-data}"
  chain = "{encoded-data}"
  recipients {
  }
  tags =  {
    tag_test = "tag_value"
  }
}
