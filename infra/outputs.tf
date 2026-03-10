output "read_api_service_account" {
  value = google_service_account.read_api.email
}

output "write_api_service_account" {
  value = google_service_account.write_api.email
}

output "worker_service_account" {
  value = google_service_account.worker.email
}

output "pubsub_topic_recipe_writes" {
  value = google_pubsub_topic.recipe_writes.name
}

output "pubsub_subscription_recipe_writes_sub" {
  value = google_pubsub_subscription.recipe_writes_sub.name
}
