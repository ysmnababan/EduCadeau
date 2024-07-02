docker build -t user_service .
docker tag user_service gcr.io/test-gcp-427110/user_service
docker push gcr.io/test-gcp-427110/user_service
gcloud run deploy user-service --image gcr.io/test-gcp-427110/user_service --platform managed --region asia-southeast2 --allow-unauthenticated --project test-gcp-427110 --port 50001
pause