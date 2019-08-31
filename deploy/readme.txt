常用命令:
curl -L http://127.0.0.1:12379/v2/members

curl -L http://127.0.0.1:12379/v2/keys/foo -XPUT -d value="Hello foo"
curl -L http://127.0.0.1:12379/v2/keys/foo1/foo1 -XPUT -d value="Hello foo1"
curl -L http://127.0.0.1:12379/v2/keys/foo2/foo2 -XPUT -d value="Hello foo2"
curl -L http://127.0.0.1:12379/v2/keys/foo2/foo21/foo21 -XPUT -d value="Hello foo21"

curl -L http://127.0.0.1:12379/v2/keys/foo
curl -L http://127.0.0.1:12379/v2/keys/foo2
curl -L http://127.0.0.1:12379/v2/keys/foo2?recursive=true


说明:
12379是2379映射的端口
