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