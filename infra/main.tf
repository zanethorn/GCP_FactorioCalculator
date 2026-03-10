terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "6.8.0"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_artifact_registry_repository" "recipes" {
  location      = var.region
  repository_id = "recipes"
  format        = "DOCKER"
}


# --- Enable APIs ---

resource "google_project_service" "run" {
  service = "run.googleapis.com"
}

resource "google_project_service" "pubsub" {
  service = "pubsub.googleapis.com"
}

resource "google_project_service" "firestore" {
  service = "firestore.googleapis.com"
}

# --- Service Accounts ---

resource "google_service_account" "read_api" {
  account_id   = "recipes-read-api"
  display_name = "Factorio Recipes Read API"
}

resource "google_service_account" "write_api" {
  account_id   = "recipes-write-api"
  display_name = "Factorio Recipes Write API"
}

resource "google_service_account" "worker" {
  account_id   = "recipes-worker"
  display_name = "Factorio Recipes Worker"
}

# --- Pub/Sub Topic & Subscription ---

resource "google_pubsub_topic" "recipe_writes" {
  name = "recipe-writes"
}

resource "google_pubsub_subscription" "recipe_writes_sub" {
  name  = "recipe-writes-sub"
  topic = google_pubsub_topic.recipe_writes.name

  ack_deadline_seconds = 30
}


# --- IAM Service Bindings ---
# Worker: subscribe to Pub/Sub
resource "google_pubsub_subscription_iam_member" "worker_subscriber" {
  subscription = google_pubsub_subscription.recipe_writes_sub.name
  role         = "roles/pubsub.subscriber"
  member       = "serviceAccount:${google_service_account.worker.email}"
}

# Write API: publish to topic
resource "google_pubsub_topic_iam_member" "write_api_publisher" {
  topic  = google_pubsub_topic.recipe_writes.name
  role   = "roles/pubsub.publisher"
  member = "serviceAccount:${google_service_account.write_api.email}"
}

# Firestore access (Datastore API roles)
resource "google_project_iam_member" "read_api_firestore" {
  project = var.project_id
  role    = "roles/datastore.user"
  member  = "serviceAccount:${google_service_account.read_api.email}"
}

resource "google_project_iam_member" "worker_firestore" {
  project = var.project_id
  role    = "roles/datastore.user"
  member  = "serviceAccount:${google_service_account.worker.email}"
}
