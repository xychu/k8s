package v1beta1

import (
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

// Install registers the API group and adds types to a scheme
func Install(scheme *runtime.Scheme) {
	log.Info("****************** install been called ******************")
	utilruntime.Must(AddToScheme(scheme))
	utilruntime.Must(scheme.SetVersionPriority(SchemeGroupVersion))
}
