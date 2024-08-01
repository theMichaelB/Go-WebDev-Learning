
variable "name-prefix" {
  default = "gh-actions"
}

variable "project_id" {
  default = "test-project-424206"
}

resource "google_service_account" "service_account" {
  account_id   = "${var.name-prefix}-service-account"
  display_name = "${var.name-prefix} Service Account"
}

resource "google_service_account_key" "service_account_key" {
  service_account_id = google_service_account.service_account.name
  public_key_type    = "TYPE_X509_PEM_FILE"
}
resource "local_file" "service_account_key_file" {
  content  = base64decode(google_service_account_key.service_account_key.private_key)
  filename = "service_account_key.json"
}

// grant owner access to project 
resource "google_project_iam_member" "this_project_iam" {
  project = var.project_id
  role    = "roles/owner"
  member  = "serviceAccount:${google_service_account.service_account.email}"
}

// enable iam api 
resource "google_project_service" "iam" {
  project = var.project_id
  service = "iam.googleapis.com"
}

// enable cloud resource manager api

resource "google_project_service" "cloudresourcemanager" {
  project = var.project_id
  service = "cloudresourcemanager.googleapis.com"
}


