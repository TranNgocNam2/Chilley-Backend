# **CHILLEY TO-DO LIST API DEPLOYMENT GUIDE**

## **Environment File (.env)**
```env
APP_HOST=0.0.0.0
APP_PORT=<port number>
CORS_DEV=<localhost:port>
CORS_PROD=<your-domain>
```

---

## **1️⃣ Build Docker Image**
```sh
docker build -t <image-name> .
```

---

## **2️⃣ Run Docker Image**
```sh
docker run --env-file .env -p 3000:3000 <image-name>
```

---

## **3️⃣ Push Docker Image to Docker Hub**
```sh
docker tag <image-name> <docker-hub-username>/<image-name>
docker push <docker-hub-username>/<image-name>
```

---

## **4️⃣ Deploy to AWS EC2**
### 🔹 **Step 1: Create AWS Account & EC2 Instance**
- Sign up for **AWS Free Tier**.
- Launch an **EC2 instance**.

### 🔹 **Step 2: SSH into EC2**
```sh
ssh -i "<your-key.pem>" ec2-user@<your-ec2-ip>
```
- If you get `Permission denied`, run:
```sh
icacls "C:\path\to\your-key.pem" /inheritance:r
icacls "C:\path\to\your-key.pem" /grant:r %username%:R
```

### 🔹 **Step 3: Install & Enable Docker**
```sh
sudo yum update -y
sudo yum install -y docker
sudo systemctl start docker
sudo systemctl enable docker
```

### 🔹 **Step 4: Pull Image from Docker Hub**
```sh
docker pull <docker-hub-username>/<image-name>
```

### 🔹 **Step 5: Create `.env` File**
```sh
nano .env
```
_Add environment variables from your local `.env` file._

### 🔹 **Step 6: Run Image in Detached Mode**
```sh
docker run -d --env-file .env -p 3000:3000 <image-name>
```

### 🔹 **Step 7: Open Port 3000 in AWS Security Group**
- Go to **AWS Console** → **EC2 Dashboard** → **Security Groups**.
- Edit **Inbound Rules** → Add **Custom TCP Rule** for **port 3000**.

### 🔹 **Step 8: Access Application**
```sh
http://<your-ec2-ip>:3000/tasks
```
_Example:_
```sh
http://ec2-54-197-82-8.compute-1.amazonaws.com:3000/tasks
```

---

## **5️⃣ Custom Domain & SSL (HTTPS)**
### 🔹 **Step 1: Create Domain & Update DNS Records**
- Buy a domain from **Namecheap, Route 53, GoDaddy**.
- Add **A Record** → Point to **EC2 Public IP**.
- Add **CNAME Record** (for `www`).

### 🔹 **Step 2: Install Nginx**
```sh
sudo yum install -y nginx                      # Amazon Linux
```

### 🔹 **Step 3: Configure Nginx**
```sh
sudo nano /etc/nginx/nginx.conf
```
_Add the following:_
```nginx
server {
    listen 80;
    server_name <A> <CNAME>;

    location / {
        proxy_pass http://localhost:3000;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```
_Save and exit (`Ctrl + X`, `Y`, `Enter`)._

### 🔹 **Step 4: Restart Nginx**
```sh
sudo systemctl restart nginx
```

### 🔹 **Step 5: Install Certbot & Get SSL Certificate**
```sh
sudo yum install -y certbot python3-certbot-nginx  # Amazon Linux
sudo certbot --nginx -d <A> -d <CNAME>
```
---

---

## **5️⃣ API Documentation**
### 🔹 **Create a Task**
**Endpoint:** `POST /tasks`
```json
{
  "title": "Task title 4",
  "description": "Task description"
}
```

```curl
curl --location 'https://chilley.nam2507.me/tasks' \
--header 'Content-Type: application/json' \
--data '{
  "title": "Task title 4",
  "description": "Task description"
}'
```

_Response:_
```json
{
  "id": 4
}
```

### 🔹 **Get All Tasks**
**Endpoint:** `GET /tasks`

```curl
curl --location 'https://chilley.nam2507.me/tasks' \
--data ''
```

_Response:_
```json
[
  {
    "id": 1,
    "title": "Task title 1",
    "description": "Task description",
    "completed": false
  },
  {
    "id": 2,
    "title": "Task title 2",
    "description": "Task description",
    "completed": false
  },
  {
    "id": 3,
    "title": "Task title 3",
    "description": "Task description",
    "completed": false
  },
  {
    "id": 4,
    "title": "Task title 4",
    "description": "Task description",
    "completed": false
  }
]
```

### 🔹 **Update a Task**
**Endpoint:** `PUT /tasks/:id`
```json
{
  "completed": true
}
```

```curl
curl --location --request PUT 'https://chilley.nam2507.me/tasks/1' \
--header 'Content-Type: application/json' \
--data '{
  "completed": true
}'
```

_Response:_
```json
{
  "message": "Task updated successfully!"
}
```

### 🔹 **Delete a Task**
**Endpoint:** `DELETE /tasks/:id`

```curl
curl --location --request DELETE 'https://chilley.nam2507.me/tasks/1'
```

_Response:_
```json
{
  "message": "Task deleted successfully!"
}
```

---


## **🎉 Deployment Completed!**



