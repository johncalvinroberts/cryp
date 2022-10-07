# CRYP SERVER

This is a Go backend for Cryp, used for persisting file uploads to AWS's S3 service.

### What does it do?

Basically, it does three things:

* Issues 1-time authentication JWTs
* Allows the user to buy upload credits
* Orchestrates upload of the encrypted file to S3

### Endpoints

* GET `/api/health`
  * Simple health endpoint
* POST `/api/whoami`
  * Initialize auth flow
  * Accepts an email as input
* POST `/api/whoami/redeem`
  * Accepts a one-time token as input
  * Returns a JWT if the authentication succeeds
* GET `/api/whoami`
  * Returns information about the current token holder
  * Protected
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
* GET `/api/uploads/:id:`
  * Get a persisted encrypted file upload.
  * Public, not protected.
* DELETE `/api/uploads/:id:`
  * Delete an upload
  * Protected
  * Soft deletes


