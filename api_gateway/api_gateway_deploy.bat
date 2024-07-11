docker build -t api_gateway .
docker tag api_gateway gcr.io/test-gcp-427110/api_gateway
docker push gcr.io/test-gcp-427110/api_gateway
gcloud run deploy api-gateway --image gcr.io/test-gcp-427110/api_gateway --platform managed --region asia-southeast2 --allow-unauthenticated --project test-gcp-427110 --port 8080
pause