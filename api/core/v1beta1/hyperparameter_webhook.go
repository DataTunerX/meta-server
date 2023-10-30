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
	"github.com/DataTunerX/meta-server/logging"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var hyperparameterlog = logging.Logger.WithName("hyperparameter-resource")

func (r *Hyperparameter) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-core-datatunerx-io-v1beta1-hyperparameter,mutating=true,failurePolicy=fail,sideEffects=None,groups=core.datatunerx.io,resources=hyperparameters,verbs=create;update,versions=v1beta1,name=mhyperparameter.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &Hyperparameter{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *Hyperparameter) Default() {
	// todo Standardised logs
	hyperparameterlog.Info("default", "name", r.Name)

	//Set the default value to the hyperparameter cr

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-core-datatunerx-io-v1beta1-hyperparameter,mutating=false,failurePolicy=fail,sideEffects=None,groups=core.datatunerx.io,resources=hyperparameters,verbs=create;update,versions=v1beta1,name=vhyperparameter.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Hyperparameter{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Hyperparameter) ValidateCreate() error {
	hyperparameterlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Hyperparameter) ValidateUpdate(old runtime.Object) error {
	hyperparameterlog.Info("validate update", "name", r.Name)
	// 禁止更新
	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Hyperparameter) ValidateDelete() error {
	hyperparameterlog.Info("validate delete", "name", r.Name)
	// 只要有引用就不能删除
	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
