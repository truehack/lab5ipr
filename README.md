# Лабораторная работа - Kubernetes

## Описание
Данный проект реализует связку двух микросервисов:
Frontend: Веб-интерфейс на Node.js (Express), работающий на порту 3000.
Backend: REST API на Go (Golang), предоставляющий системную информацию и работающий на порту 5000.

# Запуск проекта
### Сборка Frontend
docker build -t k8s-frontend:1.0 ./frontend

### Сборка Backend
docker build -t k8s-backend:1.0 ./backend

### Применение конфигурации
kubectl apply -f k8s-manifests/

## Настройка доступа (Ingress)
Для работы домена k8s-lab.local необходимо добавить запись в файл hosts.


Windows: C:\Windows\System32\drivers\etc\hosts
Linux/macOS: /etc/hosts



Добавьте строку: 127.0.0.1 k8s-lab.local



Если не установлен ingress: kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.11.0/deploy/static/provider/cloud/deploy.yaml
