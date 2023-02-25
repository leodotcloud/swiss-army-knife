#!/bin/sh

k3s server &> /dev/null &
K3S_PID=$!

while true
do
  if kubectl get nodes; then
    break
  else
    echo "k3s still not up, will recheck in 5 seconds"
    sleep 5
  fi
done

kubectl apply -f test/build-image.yaml
kubectl wait --for=condition=complete job/build-image --timeout=600s

ctr images import swiss.tar.gz docker.io/leodotcloud/swiss-army-knife:test
kubectl apply -f test/swiss.yaml
kubectl rollout status deployment/swiss

# wait another 5 sec for app to start
sleep 5

kubectl run nginx --image=nginx
kubectl wait --for=condition=ready pods/nginx

kubectl exec nginx -- sh -c "if curl swiss | grep Nato; then echo Passed; exit 0; else echo Failed; exit 1; fi"
if [ $? -eq 0 ]; then
  echo "Test passed"
else
  echo "app deployment error"
  sleep infinity
  exit 1
fi

# TODO: Should we kill k3s?
exit 0