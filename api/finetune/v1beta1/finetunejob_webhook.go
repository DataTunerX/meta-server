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
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	"github.com/DataTunerX/meta-server/logging"
)

const (
	defaultLimitGPU    = "1"
	defaultRequestGPU  = "1"
	defaultServerImage = "rayproject/ray:2.7.1-py39-cu121-finetunecode-llm"
	defaultPath        = "/root/model-data"
)

// log is for logging in this package.
var finetunejoblog = logging.Logger.WithName("finetunejob-resource")

func (r *FinetuneJob) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

//+kubebuilder:webhook:path=/mutate-finetune-datatunerx-io-v1beta1-finetunejob,mutating=true,failurePolicy=fail,sideEffects=None,groups=finetune.datatunerx.io,resources=finetunejobs,verbs=create;update,versions=v1beta1,name=mfinetunejob.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &FinetuneJob{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *FinetuneJob) Default() {
	finetunejoblog.Info("default", "name", r.Name)
	if r.Spec.FineTune.FinetuneSpec.Resource == nil {
		r.Spec.FineTune.FinetuneSpec.Resource = &Resource{
			Limits:   &ResourceLimits{},
			Requests: &ResourceLimits{},
		}
	}
	if r.Spec.FineTune.FinetuneSpec.Resource.Limits.CPU.IsZero() {
		r.Spec.FineTune.FinetuneSpec.Resource.Limits.CPU = resource.MustParse("2")
	}
	if r.Spec.FineTune.FinetuneSpec.Resource.Limits.Memory.IsZero() {
		r.Spec.FineTune.FinetuneSpec.Resource.Limits.Memory = resource.MustParse("4Gi")
	}
	if r.Spec.FineTune.FinetuneSpec.Resource.Limits.GPU == nil {
		defaultGPU := defaultLimitGPU
		r.Spec.FineTune.FinetuneSpec.Resource.Limits.GPU = &defaultGPU
	}
	if r.Spec.FineTune.FinetuneSpec.Resource.Requests.CPU.IsZero() {
		r.Spec.FineTune.FinetuneSpec.Resource.Requests.CPU = resource.MustParse("2")
	}
	if r.Spec.FineTune.FinetuneSpec.Resource.Requests.Memory.IsZero() {
		r.Spec.FineTune.FinetuneSpec.Resource.Requests.Memory = resource.MustParse("4Gi")
	}
	if r.Spec.FineTune.FinetuneSpec.Resource.Requests.GPU == nil {
		defaultGPU := defaultRequestGPU
		r.Spec.FineTune.FinetuneSpec.Resource.Requests.GPU = &defaultGPU
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-finetune-datatunerx-io-v1beta1-finetunejob,mutating=false,failurePolicy=fail,sideEffects=None,groups=finetune.datatunerx.io,resources=finetunejobs,verbs=create;update,versions=v1beta1,name=vfinetunejob.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &FinetuneJob{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *FinetuneJob) ValidateCreate() (warnings admission.Warnings, err error) {
	finetunejoblog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *FinetuneJob) ValidateUpdate(old runtime.Object) (warnings admission.Warnings, err error) {
	finetunejoblog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *FinetuneJob) ValidateDelete() (warnings admission.Warnings, err error) {
	finetunejoblog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
