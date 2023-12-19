/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	"fmt"

	"github.com/DataTunerX/utility-server/logging"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

func (r *Dataset) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-extension-datatunerx-io-v1beta1-dataset,mutating=true,failurePolicy=fail,sideEffects=None,groups=extension.datatunerx.io,resources=datasets,verbs=create;update,versions=v1beta1,name=mdataset.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Dataset{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Dataset) Default() {
	logging.ZLogger.Infof("Validate default dataset %s/%s", r.Namespace, r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-extension-datatunerx-io-v1beta1-dataset,mutating=false,failurePolicy=fail,sideEffects=None,groups=extension.datatunerx.io,resources=datasets,verbs=create;update,versions=v1beta1,name=vdataset.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Dataset{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Dataset) ValidateCreate() (warnings admission.Warnings, err error) {
	logging.ZLogger.Infof("Validate create dataset %s/%s", r.Namespace, r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Dataset) ValidateUpdate(old runtime.Object) (warnings admission.Warnings, err error) {
	logging.ZLogger.Infof("Validate update dataset %s/%s", r.Namespace, r.Name)
	if r.Status.ReferenceFinetuneName != nil {
		return nil, fmt.Errorf("dataset %s/%s is referenced by finetune, not allow update",
			r.Namespace, r.Name)
	}
	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Dataset) ValidateDelete() (warnings admission.Warnings, err error) {
	logging.ZLogger.Infof("Validate delete dataset %s/%s", r.Namespace, r.Name)
	if r.Status.ReferenceFinetuneName != nil {
		return nil, fmt.Errorf("dataset %s/%s is referenced by finetune, not allow delete",
			r.Namespace, r.Name)
	}
	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
