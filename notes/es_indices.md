### elasticsearch story

```json
curl 59.175.62.64:29200
{
  "name" : "l4ihuK0",
  "cluster_name" : "elasticsearch",
  "cluster_uuid" : "s5NfCxEXQ-iyP7uMT_faOA",
  "version" : {
    "number" : "5.6.16",
    "build_hash" : "3a740d1",
    "build_date" : "2019-03-13T15:33:36.565Z",
    "build_snapshot" : false,
    "lucene_version" : "6.6.1"
  },
  "tagline" : "You Know, for Search"
}

```

### indices
```json 
在只有一个节点的集群创建索引
PUT /my-index-000001?pretty
{
  "settings": {
    "number_of_shards": 2, // number of primary shard,primary shards can located in same node
    "number_of_replicas": 0 // number of replicas of primary shard,replica shard can not located in the node of relative primary shard
  }
}

// 索引状态维green
GET _cat/indices/my-index-000001?v
health status index           uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green open my-index-000001 AMgociyjS46EOdtTD86omQ 2 0 0 0 324b 324b

```

```json
PUT /my-index-000001?pretty
{
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1,
    "index.write.wait_for_active_shards": "2" // 需要等待2个shard都创建之后才返回,这里会超时(504),但是不代表索引没有创建成功
  }
}
```

#### query

#### index （https://www.elastic.co/guide/en/elasticsearch/reference/8.8/indices.html）


#### segment

#### cluster

#### config

#### ops

#### story
- create indices
- define mappings
- reindexing
- ingest pipelines
- delete by query
- data visualizer

#### data
curl -X PUT "http://192.168.0.90:9200/_bulk" -H 'Content-Type: application/json' --data-binary @testdata.json

用bulk方式导入json数据,json需要符合bulk格式风格