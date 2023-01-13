
cd ../aegis-web || exit

aws s3 cp index.html s3://aegis.z2h.dev/
aws s3 cp sdk-go.html s3://aegis.z2h.dev/
aws s3 cp assets/aegis-banner.png s3://aegis.z2h.dev/assets
aws s3 cp assets/aegis-icon.png s3://aegis.z2h.dev/assets
aws s3 cp assets/capture.png s3://aegis.z2h.dev/assets
aws s3 cp assets/favicon.png s3://aegis.z2h.dev/assets

aws cloudfront create-invalidation \
  --distribution-id E3LC16VB3C88K6 \
  --paths "/*"

echo "Everything is awesome!"
