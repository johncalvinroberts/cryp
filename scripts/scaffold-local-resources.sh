echo "👁 scaffolding local aws resources to localstack 👁"
aws --endpoint-url http://localhost:4566 s3 mb s3://cryps
aws --endpoint-url http://localhost:4566 s3 mb s3://cryp-pointers
aws --endpoint-url http://localhost:4566 s3 mb s3://whoami
aws --endpoint-url http://localhost:4566 s3 ls
echo "👁 done. 👁"