# go-gce-switcher
GCEのインスタンス起動停止を行う、Google Cloud Functions

## 必要なgcloudコマンド

### topicの作成
`gcloud pubsub topics create [TOPIC_NAME]`

### funcsionsのデプロイ
`gcloud functions deploy [TARGET_MODULE] --runtime go111 --trigger-topic [TARGET_TOPIC_NAME]`

### schedulerの作成
```
gcloud --project [PROJECT_ID] scheduler jobs create pubsub [PUBSUB_NAME] \
	--time-zone "[TIME_ZONE]" \
	--schedule "[CRON_TIMER]" \
	--topic [TOPIC_NAME \
	--message-body '{"switch":"start", "target":"[TARGET_INSTANCE_NAME], "zone":"[ZONE]"}'
```
