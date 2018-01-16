package secret

import (
	"github.com/jetstack/kube-lego/pkg/kubelego_const"

	k8sApi "k8s.io/api/core/v1"
)

type Secret struct {
	SecretApi *k8sApi.Secret
	exists    bool
	kubelego  kubelego.KubeLego
}
