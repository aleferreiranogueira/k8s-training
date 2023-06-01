# k8s-training

## Pré Requisitos

- [minikube](https://minikube.sigs.k8s.io/docs/start/)
- [kubectx e kubectl]()

## Intro

Esse repositório é referente ao treinamento de Kubernetes

[Slides do Treinamento](https://docs.google.com/presentation/d/1pJG_QbN4AuYa8FfyBluiyNGnvVpOA8ZIxOiTATHLOlk/edit?usp=sharing),
você irá usá-lo em conjunto com esse repo.

_Certifique-se que você está com o seu cluster minikube ligado e o contexto do
kubectx está apontando para ele_

## Estrutura do Repositório

### 1) app/

Aqui fica a aplicação em container que irá rodar no cluster k8s do minikube, é
escrita em golang [ref](https://docs.docker.com/language/golang/build-images/);
sinta-se a vontade para alterar a app para brincar com o comportamento dos
probes (altamente recomendado caso você seja dev); Publique a imagem em um
dockerhub pessoal seu (já que é uma app simples e não possuí infos da idwall) e
altere nos manifestos

### 2) k8s/

Aqui ficam os recursos apresentados durante a demo (pod, configmap, deployment,
etc) use o kubectl para aplicá-los (ex.: `kubectl apply -f 1-pod.yaml`).

[Consulte a documentação da spec dos componentes do kubernetes caso tenha alguma dúvida](https://kubernetes.io/docs/concepts/)

### 3) challenge

Aqui ficam os recursos utilizados durante o desafio (slide 23), para seguir o
desafio no seu cluster, aplique todos os recursos usando
`kubectl apply -f challenge`, e depois tente resolver os desafios usando o
`kubectl`:

<details><summary>Respostas das Perguntas</summary>
01. s4l4d <br>
02. 3 <br>
03. "App is Up!" <br>
04. "Readiness probe failed: HTTP probe failed with statuscode: 503" <br>
SHOWDOWN. "Liveness probe failed: dial tcp 10.244.1.5:666: connect: connection refused" - O liveness probe está mapeado para a porta 666, porém o container expoẽ a porta 8080 em `.spec.containers.ports, o correto seria alterar o livenessProbe para escutar a porta 8080
</details>
