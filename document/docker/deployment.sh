docker network create somenetwork

docker run -d --name elasticsearch \
        --net somenetwork \
        --restart always \
        -p 9200:9200 \
        -p 9300:9300 \
        -e "discovery.type=single-node" \
        elasticsearch:7.9.3

docker run -d --name skywalking-oap \
        --restart always \
        --net somenetwork \
        -e SW_STORAGE=elasticsearch7 \
        -e SW_STORAGE_ES_CLUSTER_NODES=elasticsearch:9200 \
        -p 11800:11800 \
        -p 12800:12800 \
        -e TZ=Asia/Shanghai  \
        apache/skywalking-oap-server

docker run -d --name skywalking-oap-ui \
        --link skywalking-oap:skywalking-oap \
        --net somenetwork \
        --restart always \
        -p 8181:8080 \
        -e SW_OAP_ADDRESS=skywalking-oap:12800 \
        -e TZ=Asia/Shanghai \
        apache/skywalking-ui

docker run -d --name kibana \
        --net somenetwork \
        -p 5601:5601 \
        kibana:7.9.3

docker run -d --name etcd-server \
        --net somenetwork \
        --publish 2379:2379 \
        --publish 2380:2380 \
        --env ALLOW_NONE_AUTHENTICATION=yes \
        --env ETCD_ADVERTISE_CLIENT_URLS=http://etcd-server:2379 \
        bitnami/etcd:latest

docker run -d --name jaeger \
        --net somenetwork \
        -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
        -p 5775:5775/udp \
        -p 6831:6831/udp \
        -p 6832:6832/udp \
        -p 5778:5778 \
        -p 16686:16686 \
        -p 14268:14268 \
        -p 14250:14250 \
        -p 9411:9411 \
        jaegertracing/all-in-one:1.21

docker run -d --name zipkin \
        --net somenetwork \
        --restart always \
        -p 9411:9411  \
        openzipkin/zipkin