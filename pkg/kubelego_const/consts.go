package kubelego

import (
	k8sApi "k8s.io/client-go/pkg/api/v1"
)

const RsaKeySize = 2048
const AcmeRegistration = "acme-registration.json"
const AcmeRegistrationUrl = "acme-registration-url"
const AcmePrivateKey = k8sApi.TLSPrivateKeyKey
const AcmeHttpChallengePath = "/.well-known/acme-challenge"
const AcmeHttpSelfTest = "/.well-known/acme-challenge/_selftest"

const TLSCertKey = k8sApi.TLSCertKey
const TLSPrivateKeyKey = k8sApi.TLSPrivateKeyKey

const AnnotationIngressChallengeEndpoints = "kubernetes.io/tls-acme-challenge-endpoints"
const AnnotationIngressClass = "kubernetes.io/ingress.class"
const AnnotationSslRedirect = "ingress.kubernetes.io/ssl-redirect"
const AnnotationKubeLegoManaged = "kubernetes.io/kube-lego-managed"

var SupportedIngressClasses = []string{"nginx", "gce"}
var AnnotationEnabled = "kubernetes.io/tls-acme"
const ChallengeTypeDNS = "dns"
const ChallengeTypeHTTP = "http"