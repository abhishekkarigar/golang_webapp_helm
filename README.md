# golang_webapp_helm

## development app:

1. first create a github project
2. clone to local machine
3. create project structure
   <github_clone_dir>
     |_MyApp
       |_ goapp-demo ( $ helm create goapp-demo)
       |_ src        ( $ go mod init github.com/abhishekkarigar/golang_webapp_helm) generates go.mod
       |    |_ main.go
            |_ go.mod
       |_ Dockerfile  
       

1. make sure to minikube ip 192.168.99.0/24 in config.json [insecurities_group]
1. minikube start
2. minikube ip
3. minikube docker-env
4. eval $(minikube docker-env)
5. docker build -t myapp .
6. $ docker run -d \
>   --restart=always \
>   -e REGISTRY_HTTP_ADDR=0.0.0.0:5002 \
>   -p 5002:5002 \
>   --name registry \
>    registry:2.6.2

7. docker tag myapp:latest 192.168.99.119:5002/myapp:latest
8. docker push 192.168.99.119:5002/myapp:latest
9. curl -k http://192.168.99.119:5002/v2/_catalog
10. helm install -n demo goapp-demo/

    NAME:   demo
    LAST DEPLOYED: Sat Jan  9 18:54:57 2021
    NAMESPACE: default
    STATUS: DEPLOYED

    RESOURCES:
    ==> v1/Deployment
    NAME             AGE
    demo-goapp-demo  1s

    ==> v1/Service
    NAME             AGE
    demo-goapp-demo  1s

    ==> v1/ServiceAccount
    NAME             AGE
    demo-goapp-demo  1s  


    NOTES:  
    1. Get the application URL by running these commands:  
      export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=goapp-demo,app.kubernetes.io/instance=demo" -o jsonpath="{.items[0].metadata.name}")  
      echo "Visit http://127.0.0.1:8080 to use your application"  
      kubectl port-forward $POD_NAME 8080:80  

11. helm list  
12. kubectl get pods  
13. $ export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=goapp-demo,app.kubernetes.io/instance=demo" -o jsonpath="{.items[0].metadata.name}")  

14. $ nohup kubectl port-forward $POD_NAME 5555:5555 &

15. curl -k 127.0.0.1:5555  
16. kubectl get svc  


