resource "aws_sqs_queue" "queue" {
  name                       = local.queue_name
  name_prefix                = local.queue_prefix
  delay_seconds              = var.delay_seconds
  max_message_size           = var.max_message_size
  message_retention_seconds  = var.message_retention_seconds
  receive_wait_time_seconds  = var.receive_wait_time_seconds
  redrive_policy             = var.redrive_policy
  visibility_timeout_seconds = var.visibility_timeout_seconds

  sqs_managed_sse_enabled = var.sqs_managed_sse_enabled

  tags = local.tags
}
