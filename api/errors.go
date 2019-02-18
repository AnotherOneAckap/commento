package main

import (
	"errors"
)

var errorMalformedTemplate = errors.New("A template is malformed.")
var errorMissingConfig = errors.New("Missing config environment variable.")
var errorCannotSendEmail = errors.New("Email dispatch failed. Please contact support to resolve this issue.")
var errorInternal = errors.New("An internal error has occurred. If you see this repeatedly, please contact support.")
var errorInvalidJSONBody = errors.New("Invalid JSON request. If you think this shouldn't happen, please contact support.")
var errorMissingField = errors.New("One or more field(s) empty.")
var errorEmailExists = errors.New("That email address is already registered. Sign in instead?")
var errorInvalidEmailPassword = errors.New("Invalid email/password combination.")
var errorUnconfirmedEmail = errors.New("Your email address is still unconfirmed. Please confirm your email address before proceeding.")
var errorNoSuchConfirmationToken = errors.New("This email confirmation link has expired.")
var errorNoSuchResetToken = errors.New("This password reset link has expired.")
var errorNotAuthorised = errors.New("You're not authorised to access that.")
var errorEmailAlreadyExists = errors.New("That email address has already been registered.")
var errorNoSuchToken = errors.New("No such session token.")
var errorNoSuchCommenter = errors.New("No such commenter.")
var errorAlreadyUpvoted = errors.New("You have already upvoted that comment.")
var errorNoSuchDomain = errors.New("This domain is not registered with Commento.")
var errorNoSuchComment = errors.New("No such comment.")
var errorInvalidState = errors.New("Invalid state value.")
var errorDomainFrozen = errors.New("Cannot add a new comment because that domain is frozen.")
var errorDomainAlreadyExists = errors.New("That domain has already been registered. Please contact support if you are the true owner.")
var errorUnauthorisedVote = errors.New("You need to be authenticated to vote.")
var errorNoSuchEmail = errors.New("No such email.")
var errorInvalidEmail = errors.New("You do not have an email registered with that account.")
var errorForbiddenEdit = errors.New("You cannot edit someone else's comment.")
var errorMissingSmtpAddress = errors.New("Missing SMTP_FROM_ADDRESS")
var errorSmtpNotConfigured = errors.New("SMTP is not configured.")
var errorOauthMisconfigured = errors.New("OAuth is misconfigured.")
var errorCannotReadResponse = errors.New("Cannot read response.")
var errorNotModerator = errors.New("You need to be a moderator to do that.")
var errorNotADirectory = errors.New("The given path is not a directory.")
var errorGzip = errors.New("Cannot GZip content.")
var errorCannotDownloadDisqus = errors.New("We could not download your Disqus export file.")
var errorSelfVote = errors.New("You cannot vote on your own comment.")
var errorInvalidConfigFile = errors.New("Invalid config file.")
var errorInvalidConfigValue = errors.New("Invalid config value.")
var errorNewOwnerForbidden = errors.New("New user registrations are forbidden and closed.")
var errorThreadLocked = errors.New("This thread is locked. You cannot add new comments.")
var errorDatabaseMigration = errors.New("Encountered error applying database migration.")
