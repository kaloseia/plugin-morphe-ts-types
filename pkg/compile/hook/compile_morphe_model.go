package hook

import (
	"github.com/kaloseia/morphe-go/pkg/yaml"
	"github.com/kaloseia/plugin-morphe-ts-types/pkg/compile/cfg"
	"github.com/kaloseia/plugin-morphe-ts-types/pkg/tsdef"
)

type CompileMorpheModel struct {
	OnCompileMorpheModelStart   OnCompileMorpheModelStartHook
	OnCompileMorpheModelSuccess OnCompileMorpheModelSuccessHook
	OnCompileMorpheModelFailure OnCompileMorpheModelFailureHook
}

type OnCompileMorpheModelStartHook = func(config cfg.MorpheModelsConfig, model yaml.Model) (cfg.MorpheModelsConfig, yaml.Model, error)
type OnCompileMorpheModelSuccessHook = func(allModelTypes []*tsdef.Object) ([]*tsdef.Object, error)
type OnCompileMorpheModelFailureHook = func(config cfg.MorpheModelsConfig, model yaml.Model, compileFailure error) error
