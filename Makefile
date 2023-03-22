#docker run -v ./data:/data ghcr.io/tigerbeetledb/tigerbeetle format --cluster=0 --replica=0 --replica-count=3 /data/0_0.tigerbeetle
#docker run -v ./data:/data ghcr.io/tigerbeetledb/tigerbeetle format --cluster=0 --replica=1 --replica-count=3 /data/0_1.tigerbeetle
#docker run -v ./data:/data ghcr.io/tigerbeetledb/tigerbeetle format --cluster=0 --replica=2 --replica-count=3 /data/0_2.tigerbeetle

initdb:
	tigerbeetle format --cluster=0 --replica=0 --replica-count=3 ./data/0_0.tigerbeetle
	tigerbeetle format --cluster=0 --replica=1 --replica-count=3 ./data/1_0.tigerbeetle
	tigerbeetle format --cluster=0 --replica=2 --replica-count=3 ./data/2_0.tigerbeetle

rundb:
	tigerbeetle start --addresses=0.0.0.0:3001,0.0.0.0:3002,0.0.0.0:3003 ./data/0_0.tigerbeetle > logs/0.log 2>&1 &
	tigerbeetle start --addresses=0.0.0.0:3001,0.0.0.0:3002,0.0.0.0:3003 ./data/1_0.tigerbeetle > logs/1.log 2>&1 &
	tigerbeetle start --addresses=0.0.0.0:3001,0.0.0.0:3002,0.0.0.0:3003 ./data/2_0.tigerbeetle > logs/2.log 2>&1 &

stopdb:
	killall tigerbeetle

debugdb:
	tigerbeetle-debug start --addresses=0.0.0.0:3001,0.0.0.0:3002,0.0.0.0:3003 ./data/0_0.tigerbeetle > logs/0.log 2>&1 &
	tigerbeetle-debug start --addresses=0.0.0.0:3001,0.0.0.0:3002,0.0.0.0:3003 ./data/1_0.tigerbeetle > logs/1.log 2>&1 &
	tigerbeetle-debug start --addresses=0.0.0.0:3001,0.0.0.0:3002,0.0.0.0:3003 ./data/2_0.tigerbeetle > logs/2.log 2>&1 &

stopdebugdb:
	killall tigerbeetle-debug

initaccounts:
	curl http://localhost:4000/account/create/1
	curl http://localhost:4000/account/create/2
