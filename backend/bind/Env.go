package bind

import (
	"CodeX/backend/core/env"
)

type EnvBind struct{}

// Node

func (e *EnvBind) GetNodeVersion() string {
	v, err := env.GetNodeVersion()
	if err != nil {
		return "error"
	}
	return v
}

func (e *EnvBind) GetNPMVersion() string {
	v, err := env.GetNPMVersion()
	if err != nil {
		return "error"
	}
	return v
}

func (e *EnvBind) GetYarnVersion() string {
	v, err := env.GetYarnVersion()
	if err != nil {
		return "error"
	}
	return v
}

func (e *EnvBind) GetPNPMVersion() string {
	v, err := env.GetPNPMVersion()
	if err != nil {
		return "error"
	}
	return v
}

// Python

func (e *EnvBind) GetPythonVersion() string {
	v, err := env.GetPythonVersion()
	if err != nil {
		return "error"
	}
	return v
}

// Go

func (e *EnvBind) GetGoVersion() string {
	v, err := env.GetGoVersion()
	if err != nil {
		return "error"
	}
	return v
}

// Rust

func (e *EnvBind) GetRustVersion() string {
	v, err := env.GetRustVersion()
	if err != nil {
		return "error"
	}
	return v
}

func (e *EnvBind) GetCargoVersion() string {
	v, err := env.GetCargoVersion()
	if err != nil {
		return "error"
	}
	return v
}

func (e *EnvBind) GetRustupVersion() string {
	v, err := env.GetRustupVersion()
	if err != nil {
		return "error"
	}
	return v
}
