package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/kubernetes/pkg/api/legacyscheme"

	appsapiv1 "github.com/pharmer/openshift/apis/apps/v1"
	appsv1 "github.com/pharmer/openshift/apis/apps/v1"
)

func init() {
	Install(legacyscheme.Scheme)
}

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	utilruntime.Must(appsapiv1.Install(scheme))
	utilruntime.Must(scheme.SetVersionPriority(appsv1.GroupVersion))
}
