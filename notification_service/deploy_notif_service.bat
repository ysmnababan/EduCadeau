docker build -t notif_service .
docker tag notif_service gcr.io/test-gcp-427110/notif_service
docker push gcr.io/test-gcp-427110/notif_service
gcloud run deploy notif-service --image gcr.io/test-gcp-427110/notif_service --platform managed --region asia-southeast2 --allow-unauthenticated --project test-gcp-427110 --port 8080