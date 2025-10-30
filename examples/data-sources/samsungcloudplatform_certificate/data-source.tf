data "samsungcloudplatform_certificate" "certificate" {
  certificate_id = "CERT-XXXXXXXXXX"
}

output "output_certificate" {
  value = data.samsungcloudplatform_certificate.certificate
}
