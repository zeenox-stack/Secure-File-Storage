# 📂 Secure File Storage API

_A simple file storage service with authentication and basic file operations._

## 📌 Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Environment Variables](#environment-variables)
- [API Endpoints](#api-endpoints)
- [Middleware](#middleware)
- [Security Considerations](#security-considerations)
- [License](#license)

---

## Introduction

Secure File Storage is a **REST API** built with Go that allows authenticated users to **upload, retrieve, download, and delete files**. It ensures security through **API key authentication** and basic access control.

---

## ✨ Features

✅ Upload files with size restrictions.  
✅ Retrieve a list of stored files.  
✅ Download stored files securely.  
✅ Delete files with authorization.  
✅ API authentication using **API keys**.

---

## 🛠 Installation

### **1. Clone the repository**

```sh
git clone https://github.com/yourusername/Secure-File-Storage.git
cd Secure-File-Storage
```

### **2. Install dependencies**

```sh
go mod tidy
```

### **3. Run the server**

```sh
go run main.go
```

By default, the server runs on **http://localhost:8000**.

---

## 🔑 Environment Variables

Create a `.env` file in the project root and define:

```sh
KEY=your-secret-api-key
```

This key will be used for authentication in API requests.

---

## 📡 API Endpoints

### **1️⃣ Upload a File**

- **Endpoint:** `POST /upload`
- **Headers:** `X-Api-Key: your-secret-api-key`
- **Body:** Multipart Form Data (file upload)
- **Response:**
  - `200 OK` → File uploaded successfully
  - `400 Bad Request` → No file provided or limit reached

```sh
curl -X POST http://localhost:8000/upload \
     -H "X-Api-Key: your-secret-api-key" \
     -F "file=@path/to/your/file.txt"
```

---

### **2️⃣ Get List of Files**

- **Endpoint:** `GET /get`
- **Headers:** `X-Api-Key: your-secret-api-key`
- **Response:**
  - `200 OK` → Returns a list of files
  - `500 Internal Server Error` → Issue retrieving files

```sh
curl -X GET http://localhost:8000/get -H "X-Api-Key: your-secret-api-key"
```

---

### **3️⃣ Download a File**

- **Endpoint:** `GET /download?file=filename.txt`
- **Headers:** `X-Api-Key: your-secret-api-key`
- **Response:**
  - `200 OK` → File served
  - `400 Bad Request` → No file specified
  - `404 Not Found` → File does not exist

```sh
curl -X GET "http://localhost:8000/download?file=example.txt" -H "X-Api-Key: your-secret-api-key"
```

---

### **4️⃣ Delete a File**

- **Endpoint:** `DELETE /delete?file=filename.txt`
- **Headers:** `X-Api-Key: your-secret-api-key`
- **Response:**
  - `200 OK` → File deleted
  - `400 Bad Request` → No file specified
  - `404 Not Found` → File does not exist

```sh
curl -X DELETE "http://localhost:8000/delete?file=example.txt" -H "X-Api-Key: your-secret-api-key"
```

---

## 🛡 Middleware (Authentication)

Every request must include an **API key** in the `X-Api-Key` header. The middleware validates this key before processing requests.

**Example Implementation:**

```go
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if key := req.Header.Get("X-Api-Key"); key != os.Getenv("KEY") {
			http.Error(res, "Forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(res, req)
	})
}
```

---

## 🔒 Security Considerations

- **API Key Protection:** Ensure your `.env` file is not exposed in public repositories.
- **File Restrictions:** Limit file size and file types to prevent malicious uploads.
- **Path Sanitization:** Prevent directory traversal (`../../etc/passwd`) by validating filenames.
- **Use HTTPS:** Deploy behind an HTTPS proxy for secure transmission.

---

## 📜 License

This project is licensed under the **MIT License**.

---
