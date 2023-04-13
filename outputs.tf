output "arn" {
  description = "The ARN of the SQS queue"
  value       = aws_sqs_queue.queue.arn
}

output "name" {
  description = "The name of the SQS queue"
  value       = aws_sqs_queue.queue.name
}

output "url" {
  description = "The URL of the SQS queue"
  value       = aws_sqs_queue.queue.url
}
