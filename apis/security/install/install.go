package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/api/legacyscheme"

	securityapiv1 "github.com/pharmer/openshift/apis/security/v1"
	securityv1 "github.com/pharmer/openshift/apis/security/v1"
)

func init() {
	Install(legacyscheme.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(securityapiv1.Install(scheme))
	utilruntime.Must(scheme.SetVersionPriority(securityv1.GroupVersion))
}
