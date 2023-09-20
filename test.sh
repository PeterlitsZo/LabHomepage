ROOT_DIR=$(cd $(dirname $0); pwd)

if [[ "$1" == "start_mysql" ]]; then
    docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=homePage mysql
elif [[ "$1" == "dev_backend" ]]; then
    cd $ROOT_DIR/backend
    RUN_MODE=dev DB_USER=root DB_PASSWORD=123456 DB_HOST=localhost:3306 DB_NAME=homePage go run .
elif [[ "$1" == "dev_frontend" ]]; then
    cd $ROOT_DIR/frontend
    pnpm run dev
fi