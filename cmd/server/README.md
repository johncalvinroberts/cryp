# CRYP SERVER

This is a Go backend for Cryp, used for persisting file uploads to AWS's S3 service.

### What does it do?

Basically, it does three things:

* Issues 1-time authentication JWTs
* Allows the user to buy upload credits
* Orchestrates read/write of the encrypted file to S3

### Endpoints

* GET `/api/health`
  * Simple health endpoint
* POST `/api/whoami/start`
  * Initialize auth flow
  * Accepts an email as input
* POST `/api/whoami/try`
  * Accepts a one-time token as input
  * Returns a JWT if the authentication succeeds
* POST `/api/whoami/refresh`
  * Requires a valid token
  * Returns a new token
* DELETE `/api/whoami/end`
  * Revoke the current token holder's token
* GET `/api/whoami`
  * Returns information about the current token holder
  * Protected
* DELETE `/api/whoami`
  * Delete all records related to the current token holder (!!)
* POST `/api/creds`
  * Initialize payment flow
  * Protected
* GET `/api/creds`
  * Get creds of current token holder
  * Protected
* POST `/api/uploads`
  * Initialize upload flow
  * Protected
* GET `/api/uploads`
  * Get uploads of current token holder
  * Protected
* GET `/api/uploads/:key:`
  * Get a persisted encrypted file upload.
  * Public, not protected.
* DELETE `/api/uploads/:id:`
  * Delete an upload
  * Protected
  * Soft deletes


