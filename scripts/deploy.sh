docker build -t ranggarifqi/qashir-backend .
curl https://cli-assets.heroku.com/install.sh | sh
heroku auth:token
# heroku container:login
docker login --username=_ --password=$(heroku auth:token) registry.heroku.com
docker tag ranggarifqi/qashir-backend registry.heroku.com/qashir-backend/web
docker push registry.heroku.com/qashir-backend/web
heroku container:release web