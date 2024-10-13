# Batman Image Analysis Platform Lite

Batman is an image upload and analysis platform that supports basic image management operations through a RESTful API. This project allows users to upload image metadata, retrieve images, update image details, and delete images.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Project](#running-the-project)
- [API Endpoints](#api-endpoints)
- [Curl Commands](#curl-commands)

## Prerequisites

Make sure you have the following installed:

- Go (1.21.12)
- PostgreSQL (I have used [Neon PostgreSQL](https://neon.tech/) for simplicity)
- Make (optional, for using the Makefile)

## Installation
This will download and install all the necessary Go modules specified in the `go.mod` file.
```sh
go mod tidy
```

## Configuration
For environment variables create a `.env` file in the root directory of your project and add the following variables:
```env
SERVER_PORT=8080
DB_URL=<valid_url>
```

## Running the Project
You can run the project using the `Makefile` or directly with Go.

### Using Makefile
If you have `make` installed, you can use the following command to run the project:
```sh
make run
```
### Using Go
Alternatively, you can run the project directly with Go:
```sh
go run main.go
```

## API Endpoints
Here are the available API endpoints for the Batman project:
These endpoints allow you to manage image metadata and files, ensuring a comprehensive image upload and analysis platform.

### Image Metadata
- **Upload Image Metadata**
    - **Endpoint**: `POST /api/images`
    - **Description**: Upload metadata for a new image.
    - **Request Body**: JSON containing image metadata.

- **Retrieve Image Metadata**
    - **Endpoint**: `GET /api/v1/images/{id}`
    - **Description**: Retrieve metadata for a specific image by ID.
    - **Path Parameter**: `id` - The ID of the image.

- **Update Image Metadata**
    - **Endpoint**: `PUT /api/v1/images/{id}`
    - **Description**: Update metadata for a specific image by ID.
    - **Path Parameter**: `id` - The ID of the image.
    - **Request Body**: JSON containing updated image metadata.

- **Delete Image Metadata**
    - **Endpoint**: `DELETE /api/v1/images/{id}`
    - **Description**: Delete metadata for a specific image by ID.
    - **Path Parameter**: `id` - The ID of the image.

### Image Files
- **Upload Image File**
    - **Endpoint**: `POST /api/v1/images/{id}/file`
    - **Description**: Upload the actual image file for a specific image by ID.
    - **Path Parameter**: `id` - The ID of the image.
    - **Request Body**: Multipart form-data containing the image file.

- **Retrieve Image File**
    - **Endpoint**: `GET /api/v1/images/{id}/file`
    - **Description**: Retrieve the actual image file for a specific image by ID.
    - **Path Parameter**: `id` - The ID of the image.

## Curl Commands
Below are example curl commands to interact with the Batman API endpoints:

### Image Metadata
- **Upload Image Metadata**
    ```sh
    curl -X POST http://localhost:8080/api/v1/images/ \
    -H "Content-Type: application/json" \
    -d '{
        "user_id": 12345,
        "original_filename": "image.jpg",
        "height": 1080,
        "width": 1920,
        "file_size": 204800,
        "file_type": "image/jpeg"
    }'
    ```

- **Retrieve Image Metadata**
    ```sh
    curl -X GET http://localhost:8080/api/v1/images/{id}
    ```

- **Update Image Metadata**
    ```sh
    curl -X PUT http://localhost:8080/api/v1/images/9d0855ff-471e-47e4-be76-46116171d3b6 \
    -H "Content-Type: application/json" \
    -d '{
        "original_filename": "updated_image.jpg"
    }'
    ```

- **Delete Image Metadata**
    ```sh
    curl -X DELETE http://localhost:8080/api/v1/images/{id}
    ```

### Image Files [TODO]
- **Upload Image File**
    ```sh
    curl -X POST http://localhost:8080/api/v1/images/{id}/file -F "file=@/path/to/your/image.jpg"
    ```

- **Retrieve Image File**
    ```sh
    curl -X GET http://localhost:8080/api/v1/images/{id}/file
    ```

### Health/Ping Check
- **Health Check**
    ```sh
    curl -X GET http://localhost:8080/ping
    ```
