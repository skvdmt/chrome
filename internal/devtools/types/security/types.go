package security

// MixedContentType Описание смешанного контента
// (HTTP-ресурсы на страницах HTTPS), как определено в
// https://www.w3.org/TR/mixed-content/#categories
// Допустимые значения: blockable, optionally-blockable, none
type MixedContentType string

// CertificateId Внутренний идентификатор сертификата.
type CertificateId int

// SecurityState Уровень безопасности страницы или ресурса.
// Допустимые значения: unknown, neutral, insecure, secure, info, insecure-broken
type SecurityState string
