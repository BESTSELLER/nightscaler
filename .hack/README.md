# Hack

> All commands are to be run from the `.hack` folder and the following environment variable to be present.
>
> ```shell
> export PUBSUB_PROJECT_ID="es-standalone-cb21"
> export PUBSUB_EMULATOR_HOST="localhost:8085"
> export NIGHTSCALER_PROJECT_ID="es-standalone-cb21"
> export NIGHTSCALER_LISTEN_TOPIC="nightscaler-scale-up"
> export NIGHTSCALER_PUBLISH_TOPIC="nightscaler-clusters"
> export NIGHTSCALER_CLUSTERNAME="es02-dev"
> export NIGHTSCALER_DEBUG="true"
> export NIGHTSCALER_JSON_LOGGING="false"
> ```

To start the Pub/Sub emulator, run:

```shell
gcloud beta emulators pubsub start --project=$PUBSUB_PROJECT_ID
```

To create the required topics, run:

```shell
pip install -r ./requirements.txt

python ./publisher.py $PUBSUB_PROJECT_ID create $NIGHTSCALER_LISTEN_TOPIC
python ./publisher.py $PUBSUB_PROJECT_ID create $NIGHTSCALER_PUBLISH_TOPIC
```

To send a "scale up" event, run:

```shell
python publisher.py $PUBSUB_PROJECT_ID publish-with-custom-attributes $NIGHTSCALER_LISTEN_TOPIC es02-dev default
```

If you want to view all the messages sent, you can start the json-server:

```shell
python ./subscriber.py $PUBSUB_PROJECT_ID create-push $NIGHTSCALER_PUBLISH_TOPIC $NIGHTSCALER_PUBLISH_TOPIC http://localhost:3000/messages

cat > ./db.json <<EOF
{
  "messages": []
}
EOF

json-server --port 3000 --watch db.json
```

The data field is a base64 encoded string, so you can decode it with:

```shell
cat db.json | jq -r '.messages[0].message.data' | base64 -d | jq '.[].metadata.name'
```

```shell
cat > ./db.json <<EOF
{
  "messages": []
}
EOF

json-server --port 3000 --watch db.json
```
