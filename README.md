# A simple file 📁 upload server with the ability to upload ⬆️ and download ⬇️ files with pre-signed URLs written in GO 

- The server uses [Minio](https://min.io/) or [S3](https://aws.amazon.com/s3/) as the object persistence layer

- The object references are stored in [PostgreSQL](https://www.postgresql.org/)

- Still work in progress ⚠️⚠️⚠️

## Feature Roadmap
- [x] Pre-Signed File Upload 
- [x] Pre-Signed File Download
- [x] Ability to split a single S3 or (any S3 compatible) bucket into logical buckets using file keys
- [ ] Resumable File uploads using [Tus](https://tus.io/)
- [ ] Upload tracking for Pre singed Urls via a Job queue
- [ ] Basic JWT based auth
- [ ] File Locker for concurrent Resumable File uploads
- [ ] Caching With Redis
- [ ] Multiple Database support
- [ ] Advance Pluggable Auth Layer
- [ ] UI Dashboard

## Limitations
- The server is limited to handle objects of 5GB in size
  - Primary reason is the logical file rename or logical bucket rename happens by copying the folder or file from on source to destination and removing the old destination and multipart operations are not implemented to handle files larger than 5GB
- The server is underpinned by eventual consistency because all S3 compatible providers don't have file upload events so workers are used to check upload status. So there can be stale data window