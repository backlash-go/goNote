#!/bin/bash

Timer=`date +%Y-%m-%d\[%H:%M:%S\]`



TestServicesNum=`kubectl get deployments.  -l k8s=test |wc -l`
BetaServicesNum=`kubectl get deployments.  -l k8s=beta -n beta |wc -l`
OnlineServicesNum=`kubectl get deployments.  -l k8s=online |wc -l`


path=./bak/$Timer
mkdir -p ./bak/$Timer/{test,beta,online}
a=2
b=2
c=2




if [[ $TestServicesNum >=3 ]]
then

kubectl  get ingresses. -o yaml > ./bak/$Timer/test/ingress.yaml
kubectl  get secrets  -l env=test -o yaml > ./bak/$Timer/test/secrect.yaml
kubectl  get configmap  -l env=test -o yaml > ./bak/$Timer/test/configmap.yaml

while [[ $a -le $TestServicesNum ]]

do
svc=`kubectl get deployments. -l k8s=test |awk '{print $1}'|sed -n $a\p`

mkdir -p $path/test/$svc

datapath=$path/test/$svc

`kubectl get deployments. $svc --export=true -o yaml > $datapath/deployment.yml`> /dev/null 2>&1
`kubectl get svc $svc --export=true -o yaml > $datapath/svc.yml`> /dev/null 2>&1

a=$(($a+1))
sleep 1
done

fi


if [[ $BetaServicesNum >=3 ]]
then
kubectl  get ingresses. -o yaml -n beta > ./bak/$Timer/beta/ingress.yaml
kubectl  get secrets  -l env=beta -o yaml -n beta > ./bak/$Timer/beta/secrect.yaml
kubectl  get configmap  -l env=beta -o yaml -n beta> ./bak/$Timer/beta/configmap.yaml
while [[ $b -le $BetaServicesNum ]]

do
svc=`kubectl get deployments. -n beta -l k8s=beta |awk '{print $1}'|sed -n $b\p`

mkdir -p $path/beta/$svc

datapath=$path/beta/$svc

`kubectl get deployments. -n beta $svc --export=true -o yaml > $datapath/deployment.yml`> /dev/null 2>&1
`kubectl get svc -n beta $svc --export=true -o yaml > $datapath/svc.yml`> /dev/null 2>&1

b=$(($b+1))
sleep 1
done

fi




if [[ $OnlineServicesNum >=3 ]]
then

kubectl  get ingresses. -o yaml > ./bak/$Timer/online/ingress.yaml
kubectl  get secrets  -l env=online -o yaml > ./bak/$Timer/online/secrect.yaml
kubectl  get configmap  -l env=online -o yaml > ./bak/$Timer/online/configmap.yaml

while [[ $c -le $OnlineServicesNum ]]

do
svc=`kubectl get deployments.  -l k8s=online |awk '{print $1}'|sed -n $c\p`

mkdir -p $path/online/$svc

datapath=$path/online/$svc

`kubectl get deployments. $svc --export=true -o yaml > $datapath/deployment.yml`> /dev/null 2>&1
`kubectl get svc $svc --export=true -o yaml > $datapath/svc.yml`> /dev/null 2>&1

c=$(($c+1))
sleep 1
done

fi
git add ./
git commit -m "$Timer"
git push