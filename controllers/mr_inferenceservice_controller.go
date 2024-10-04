package controllers

import (
	"github.com/go-logr/logr"
	infrctrl "github.com/kubeflow/model-registry/pkg/inferenceservice-controller"
	"github.com/opendatahub-io/odh-model-controller/controllers/constants"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewModelRegistryInferenceServiceReconciler(client client.Client, log logr.Logger, skipTLSVerify bool, bearerToken, mrNamespace string) *infrctrl.InferenceServiceController {
	return infrctrl.NewInferenceServiceController(
		client,
		log,
		skipTLSVerify,
		bearerToken,
		constants.ModelRegistryInferenceServiceIdLabel,
		constants.ModelRegistryRegisteredModelIdLabel,
		constants.ModelRegistryModelVersionIdLabel,
		constants.ModelRegistryNamespaceLabel,
		constants.ModelRegistryNameLabel,
		constants.ModelRegistryUrlAnnotation,
		constants.ModelRegistryFinalizer,
		mrNamespace,
	)
}
