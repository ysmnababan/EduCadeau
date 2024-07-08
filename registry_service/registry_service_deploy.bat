docker build -t registry_service .
docker tag registry_service gcr.io/test-gcp-427110/registry_service
docker push gcr.io/test-gcp-427110/registry_service
gcloud run deploy registry-service --image gcr.io/test-gcp-427110/registry_service --platform managed --region asia-southeast2 --allow-unauthenticated --project test-gcp-427110 --port 50003
pause