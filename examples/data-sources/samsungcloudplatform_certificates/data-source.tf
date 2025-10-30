data "samsungcloudplatform_certificates" "my_certificates" {
}

output "contents" {
  value = data.samsungcloudplatform_certificates.my_certificates.contents
}
